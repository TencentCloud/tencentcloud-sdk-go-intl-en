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
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

type Address struct {
	// Unique EIP ID.
	AddressId *string `json:"AddressId,omitempty" name:"AddressId"`

	// EIP name.
	AddressName *string `json:"AddressName,omitempty" name:"AddressName"`

	// EIP status, including 'CREATING' (creating), 'BINDING' (binding), 'BIND' (bound), 'UNBINDING' (unbinding), 'UNBIND' (unbound), 'OFFLINING' (releasing), and 'BIND_ENI' (binding dangling ENI)
	AddressStatus *string `json:"AddressStatus,omitempty" name:"AddressStatus"`

	// Public IP address
	AddressIp *string `json:"AddressIp,omitempty" name:"AddressIp"`

	// ID of the bound resource instance.
	// Note: this field may return null, indicating that no valid values can be obtained.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Creation time in ISO 8601 format (YYYY-MM-DDTHH:mm:ss.sssZ)
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// ID of the bound ENI
	// Note: this field may return null, indicating that no valid values can be obtained.
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`

	// Private IP of the bound resource
	// Note: this field may return null, indicating that no valid values can be obtained.
	PrivateAddressIp *string `json:"PrivateAddressIp,omitempty" name:"PrivateAddressIp"`

	// Isolation status of the resource. true: isolated; false: not isolated.
	IsArrears *bool `json:"IsArrears,omitempty" name:"IsArrears"`

	// Blockage status of the EIP resource. true: blocked; false: not blocked
	IsBlocked *bool `json:"IsBlocked,omitempty" name:"IsBlocked"`

	// Whether the EIP supports direct access mode. true: yes; false: no.
	IsEipDirectConnection *bool `json:"IsEipDirectConnection,omitempty" name:"IsEipDirectConnection"`

	// Resource type of the EIP, including `CalcIP` (device IP), `WanIP` (general public IP), `EIP` (elastic IP), and `AnycastEip` (accelerated EIP).
	AddressType *string `json:"AddressType,omitempty" name:"AddressType"`

	// Whether the EIP is automatically released after being unbound. true: yes; false: no
	CascadeRelease *bool `json:"CascadeRelease,omitempty" name:"CascadeRelease"`

	// ISP. CTCC: China Telecom; CUCC: China Unicom; CMCC: China Mobile
	// Note: this field may return null, indicating that no valid values can be obtained.
	InternetServiceProvider *string `json:"InternetServiceProvider,omitempty" name:"InternetServiceProvider"`

	// Bandwidth cap
	// Note: this field may return null, indicating that no valid values can be obtained.
	Bandwidth *uint64 `json:"Bandwidth,omitempty" name:"Bandwidth"`

	// Billing mode
	// Note: this field may return null, indicating that no valid values can be obtained.
	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`
}

type AddressInfo struct {
	// Public IP information of the instance.
	// Note: this field may return null, indicating that no valid values can be obtained.
	PublicIPAddressInfo *PublicIPAddressInfo `json:"PublicIPAddressInfo,omitempty" name:"PublicIPAddressInfo"`

	// Private IP information of the instance.
	// Note: this field may return null, indicating that no valid values can be obtained.
	PrivateIPAddressInfo *PrivateIPAddressInfo `json:"PrivateIPAddressInfo,omitempty" name:"PrivateIPAddressInfo"`

	// Public IPv6 information of the instance.
	// Note: this field may return null, indicating that no valid values can be obtained.
	PublicIPv6AddressInfo *PublicIPAddressInfo `json:"PublicIPv6AddressInfo,omitempty" name:"PublicIPv6AddressInfo"`
}

type AddressTemplateSpecification struct {
	// IP address ID, such as `eipm-2uw6ujo6`.
	AddressId *string `json:"AddressId,omitempty" name:"AddressId"`

	// IP address group ID, such as `eipmg-2uw6ujo6`.
	AddressGroupId *string `json:"AddressGroupId,omitempty" name:"AddressGroupId"`
}

// Predefined struct for user
type AllocateAddressesRequestParams struct {
	// ECM region
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// Number of EIPs. Default value: 1.
	AddressCount *uint64 `json:"AddressCount,omitempty" name:"AddressCount"`

	// CMCC: China Mobile
	// CTCC: China Telecom
	// CUCC: China Unicom
	InternetServiceProvider *string `json:"InternetServiceProvider,omitempty" name:"InternetServiceProvider"`

	// 1–5000 Mbps. Default value: 1 Mbps.
	InternetMaxBandwidthOut *uint64 `json:"InternetMaxBandwidthOut,omitempty" name:"InternetMaxBandwidthOut"`

	// List of tags to be bound.
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// ID of the instance to be bound.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// ID of the ENI to be bound, such as `eni-11112222`. `NetworkInterfaceId` and `InstanceId` cannot be specified at the same time. The ENI ID can be obtained from the `networkInterfaceId` field in the returned value of the `DescribeNetworkInterfaces` API.
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`

	// Private IP to be bound. If you specify `NetworkInterfaceId`, you must also specify `PrivateIpAddress`, which means to bind the EIP to the specified private IP of the specified ENI. You must also make sure that the specified `PrivateIpAddress` is a private IP of the specified `NetworkInterfaceId`. The private IP of the specified ENI can be obtained from the `privateIpAddress` field in the returned value of the `DescribeNetworkInterfaces` API.
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" name:"PrivateIpAddress"`
}

type AllocateAddressesRequest struct {
	*tchttp.BaseRequest
	
	// ECM region
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// Number of EIPs. Default value: 1.
	AddressCount *uint64 `json:"AddressCount,omitempty" name:"AddressCount"`

	// CMCC: China Mobile
	// CTCC: China Telecom
	// CUCC: China Unicom
	InternetServiceProvider *string `json:"InternetServiceProvider,omitempty" name:"InternetServiceProvider"`

	// 1–5000 Mbps. Default value: 1 Mbps.
	InternetMaxBandwidthOut *uint64 `json:"InternetMaxBandwidthOut,omitempty" name:"InternetMaxBandwidthOut"`

	// List of tags to be bound.
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// ID of the instance to be bound.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// ID of the ENI to be bound, such as `eni-11112222`. `NetworkInterfaceId` and `InstanceId` cannot be specified at the same time. The ENI ID can be obtained from the `networkInterfaceId` field in the returned value of the `DescribeNetworkInterfaces` API.
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`

	// Private IP to be bound. If you specify `NetworkInterfaceId`, you must also specify `PrivateIpAddress`, which means to bind the EIP to the specified private IP of the specified ENI. You must also make sure that the specified `PrivateIpAddress` is a private IP of the specified `NetworkInterfaceId`. The private IP of the specified ENI can be obtained from the `privateIpAddress` field in the returned value of the `DescribeNetworkInterfaces` API.
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" name:"PrivateIpAddress"`
}

func (r *AllocateAddressesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AllocateAddressesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EcmRegion")
	delete(f, "AddressCount")
	delete(f, "InternetServiceProvider")
	delete(f, "InternetMaxBandwidthOut")
	delete(f, "Tags")
	delete(f, "InstanceId")
	delete(f, "NetworkInterfaceId")
	delete(f, "PrivateIpAddress")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AllocateAddressesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AllocateAddressesResponseParams struct {
	// List of unique IDs of the EIPs applied for.
	// Note: this field may return null, indicating that no valid values can be obtained.
	AddressSet []*string `json:"AddressSet,omitempty" name:"AddressSet"`

	// Async task ID (TaskId). You can use the `DescribeTaskResult` API to query the task status.
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AllocateAddressesResponse struct {
	*tchttp.BaseResponse
	Response *AllocateAddressesResponseParams `json:"Response"`
}

func (r *AllocateAddressesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AllocateAddressesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Area struct {
	// Region ID
	AreaId *string `json:"AreaId,omitempty" name:"AreaId"`

	// Region name
	AreaName *string `json:"AreaName,omitempty" name:"AreaName"`
}

// Predefined struct for user
type AssignIpv6AddressesRequestParams struct {
	// ECM region
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// ENI instance ID, such as `eni-1snva0vd`. Currently, only the primary ENI will be assigned the ID.
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`

	// List of specified IPv6 addresses. You can specify up to 10 IPv6 addresses at a time. The quota is calculated together with that of `Ipv6AddressCount`, a required input parameter alternative to this one.
	Ipv6Addresses []*Ipv6Address `json:"Ipv6Addresses,omitempty" name:"Ipv6Addresses"`

	// Number of automatically assigned IPv6 addresses. The total number of private IP addresses cannot exceed the quota. The quota is calculated together with that of `Ipv6Addresses`, a required input parameter alternative to this one.
	Ipv6AddressCount *int64 `json:"Ipv6AddressCount,omitempty" name:"Ipv6AddressCount"`

	// Ipv6 ISP. Valid values:
	// `CTCC`: China Telecom
	// `CUCC`: China Unicom
	// `CMCC`: China Mobile
	Ipv6ISP *string `json:"Ipv6ISP,omitempty" name:"Ipv6ISP"`
}

type AssignIpv6AddressesRequest struct {
	*tchttp.BaseRequest
	
	// ECM region
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// ENI instance ID, such as `eni-1snva0vd`. Currently, only the primary ENI will be assigned the ID.
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`

	// List of specified IPv6 addresses. You can specify up to 10 IPv6 addresses at a time. The quota is calculated together with that of `Ipv6AddressCount`, a required input parameter alternative to this one.
	Ipv6Addresses []*Ipv6Address `json:"Ipv6Addresses,omitempty" name:"Ipv6Addresses"`

	// Number of automatically assigned IPv6 addresses. The total number of private IP addresses cannot exceed the quota. The quota is calculated together with that of `Ipv6Addresses`, a required input parameter alternative to this one.
	Ipv6AddressCount *int64 `json:"Ipv6AddressCount,omitempty" name:"Ipv6AddressCount"`

	// Ipv6 ISP. Valid values:
	// `CTCC`: China Telecom
	// `CUCC`: China Unicom
	// `CMCC`: China Mobile
	Ipv6ISP *string `json:"Ipv6ISP,omitempty" name:"Ipv6ISP"`
}

func (r *AssignIpv6AddressesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssignIpv6AddressesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EcmRegion")
	delete(f, "NetworkInterfaceId")
	delete(f, "Ipv6Addresses")
	delete(f, "Ipv6AddressCount")
	delete(f, "Ipv6ISP")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AssignIpv6AddressesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssignIpv6AddressesResponseParams struct {
	// List of IPv6 addresses assigned to ENIs.
	Ipv6AddressSet []*Ipv6Address `json:"Ipv6AddressSet,omitempty" name:"Ipv6AddressSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AssignIpv6AddressesResponse struct {
	*tchttp.BaseResponse
	Response *AssignIpv6AddressesResponseParams `json:"Response"`
}

func (r *AssignIpv6AddressesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssignIpv6AddressesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssignPrivateIpAddressesRequestParams struct {
	// ENI instance ID, such as `eni-m6dyj72l`.
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`

	// ECM region, such as `ap-xian-ecm`.
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// Information of the specified private IPs. You can specify up to 10 IPs at a time. You must provide either this parameter or `SecondaryPrivateIpAddressCount`, or both.
	PrivateIpAddresses []*PrivateIpAddressSpecification `json:"PrivateIpAddresses,omitempty" name:"PrivateIpAddresses"`

	// Number of private IP addresses applied for. You must provide either this parameter or `PrivateIpAddresses`, or both. The total number of private IP addresses cannot exceed the quota.
	SecondaryPrivateIpAddressCount *uint64 `json:"SecondaryPrivateIpAddressCount,omitempty" name:"SecondaryPrivateIpAddressCount"`
}

type AssignPrivateIpAddressesRequest struct {
	*tchttp.BaseRequest
	
	// ENI instance ID, such as `eni-m6dyj72l`.
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`

	// ECM region, such as `ap-xian-ecm`.
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// Information of the specified private IPs. You can specify up to 10 IPs at a time. You must provide either this parameter or `SecondaryPrivateIpAddressCount`, or both.
	PrivateIpAddresses []*PrivateIpAddressSpecification `json:"PrivateIpAddresses,omitempty" name:"PrivateIpAddresses"`

	// Number of private IP addresses applied for. You must provide either this parameter or `PrivateIpAddresses`, or both. The total number of private IP addresses cannot exceed the quota.
	SecondaryPrivateIpAddressCount *uint64 `json:"SecondaryPrivateIpAddressCount,omitempty" name:"SecondaryPrivateIpAddressCount"`
}

func (r *AssignPrivateIpAddressesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssignPrivateIpAddressesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NetworkInterfaceId")
	delete(f, "EcmRegion")
	delete(f, "PrivateIpAddresses")
	delete(f, "SecondaryPrivateIpAddressCount")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AssignPrivateIpAddressesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssignPrivateIpAddressesResponseParams struct {
	// Private IP details.
	// Note: this field may return null, indicating that no valid values can be obtained.
	PrivateIpAddressSet []*PrivateIpAddressSpecification `json:"PrivateIpAddressSet,omitempty" name:"PrivateIpAddressSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AssignPrivateIpAddressesResponse struct {
	*tchttp.BaseResponse
	Response *AssignPrivateIpAddressesResponseParams `json:"Response"`
}

func (r *AssignPrivateIpAddressesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssignPrivateIpAddressesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssistantCidr struct {
	// VPC instance ID, such as `vpc-6v2ht8q5`
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Secondary CIDR, such as `172.16.0.0/16`
	CidrBlock *string `json:"CidrBlock,omitempty" name:"CidrBlock"`

	// Secondary CIDR block type. 0: general secondary CIDR block; 1: container secondary CIDR block. Default value: 0.
	AssistantType *uint64 `json:"AssistantType,omitempty" name:"AssistantType"`

	// Subnets divided by the secondary CIDR block.
	// Note: this field may return null, indicating that no valid values can be obtained.
	// Note: this field may return null, indicating that no valid values can be obtained.
	SubnetSet []*Subnet `json:"SubnetSet,omitempty" name:"SubnetSet"`
}

// Predefined struct for user
type AssociateAddressRequestParams struct {
	// ECM region
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// Unique EIP ID, such as `eip-11112222`.
	AddressId *string `json:"AddressId,omitempty" name:"AddressId"`

	// ID of the instance to be bound.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// ID of the ENI to be bound, such as `eni-11112222`. `NetworkInterfaceId` and `InstanceId` cannot be specified at the same time. The ENI ID can be obtained from the `networkInterfaceId` field in the returned value of the `DescribeNetworkInterfaces` API.
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`

	// Private IP to be bound. If you specify `NetworkInterfaceId`, you must also specify `PrivateIpAddress`, which means to bind the EIP to the specified private IP of the specified ENI. You must also make sure that the specified `PrivateIpAddress` is a private IP of the specified `NetworkInterfaceId`. The private IP of the specified ENI can be obtained from the `privateIpAddress` field in the returned value of the `DescribeNetworkInterfaces` API.
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" name:"PrivateIpAddress"`
}

type AssociateAddressRequest struct {
	*tchttp.BaseRequest
	
	// ECM region
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// Unique EIP ID, such as `eip-11112222`.
	AddressId *string `json:"AddressId,omitempty" name:"AddressId"`

	// ID of the instance to be bound.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// ID of the ENI to be bound, such as `eni-11112222`. `NetworkInterfaceId` and `InstanceId` cannot be specified at the same time. The ENI ID can be obtained from the `networkInterfaceId` field in the returned value of the `DescribeNetworkInterfaces` API.
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`

	// Private IP to be bound. If you specify `NetworkInterfaceId`, you must also specify `PrivateIpAddress`, which means to bind the EIP to the specified private IP of the specified ENI. You must also make sure that the specified `PrivateIpAddress` is a private IP of the specified `NetworkInterfaceId`. The private IP of the specified ENI can be obtained from the `privateIpAddress` field in the returned value of the `DescribeNetworkInterfaces` API.
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" name:"PrivateIpAddress"`
}

func (r *AssociateAddressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateAddressRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EcmRegion")
	delete(f, "AddressId")
	delete(f, "InstanceId")
	delete(f, "NetworkInterfaceId")
	delete(f, "PrivateIpAddress")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AssociateAddressRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssociateAddressResponseParams struct {
	// Async task ID (TaskId). You can use the `DescribeTaskResult` API to query the task status.
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AssociateAddressResponse struct {
	*tchttp.BaseResponse
	Response *AssociateAddressResponseParams `json:"Response"`
}

func (r *AssociateAddressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateAddressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssociateSecurityGroupsRequestParams struct {
	// ID of the security group to be bound, such as `esg-efil73jd`. You can bind only one security group.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`

	// ID of the bound instance, such as `ein-lesecurk`. You can specify multiple instances and request up to 100 instances at a time.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

type AssociateSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// ID of the security group to be bound, such as `esg-efil73jd`. You can bind only one security group.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`

	// ID of the bound instance, such as `ein-lesecurk`. You can specify multiple instances and request up to 100 instances at a time.
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

// Predefined struct for user
type AttachNetworkInterfaceRequestParams struct {
	// ENI instance ID, such as `eni-m6dyj72l`.
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`

	// Instance ID, such as `ein-r8hr2upy`.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// ECM region, such as `ap-xian-ecm`.
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`
}

type AttachNetworkInterfaceRequest struct {
	*tchttp.BaseRequest
	
	// ENI instance ID, such as `eni-m6dyj72l`.
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`

	// Instance ID, such as `ein-r8hr2upy`.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// ECM region, such as `ap-xian-ecm`.
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`
}

func (r *AttachNetworkInterfaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AttachNetworkInterfaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NetworkInterfaceId")
	delete(f, "InstanceId")
	delete(f, "EcmRegion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AttachNetworkInterfaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AttachNetworkInterfaceResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AttachNetworkInterfaceResponse struct {
	*tchttp.BaseResponse
	Response *AttachNetworkInterfaceResponseParams `json:"Response"`
}

func (r *AttachNetworkInterfaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AttachNetworkInterfaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Backend struct {
	// Unique real server ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Listening port of the real server
	// Note: this field may return null, indicating that no valid values can be obtained.
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// Forwarding weight of the real server. Value range: [0, 100]. Default value: 10.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Weight *int64 `json:"Weight,omitempty" name:"Weight"`

	// Private IP of the real server
	// Note: this field may return null, indicating that no valid values can be obtained.
	PrivateIpAddresses []*string `json:"PrivateIpAddresses,omitempty" name:"PrivateIpAddresses"`

	// Real server binding time
	// Note: this field may return null, indicating that no valid values can be obtained.
	RegisteredTime *string `json:"RegisteredTime,omitempty" name:"RegisteredTime"`

	// Unique ENI ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	EniId *string `json:"EniId,omitempty" name:"EniId"`

	// Public IP of the real server
	// Note: this field may return null, indicating that no valid values can be obtained.
	PublicIpAddresses []*string `json:"PublicIpAddresses,omitempty" name:"PublicIpAddresses"`

	// Real server instance name
	// Note: this field may return null, indicating that no valid values can be obtained.
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
}

// Predefined struct for user
type BatchDeregisterTargetsRequestParams struct {
	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// Unbound targets
	Targets []*BatchTarget `json:"Targets,omitempty" name:"Targets"`
}

type BatchDeregisterTargetsRequest struct {
	*tchttp.BaseRequest
	
	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// Unbound targets
	Targets []*BatchTarget `json:"Targets,omitempty" name:"Targets"`
}

func (r *BatchDeregisterTargetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchDeregisterTargetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "Targets")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchDeregisterTargetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchDeregisterTargetsResponseParams struct {
	// IDs of the listeners failed to be unbound
	// Note: this field may return null, indicating that no valid values can be obtained.
	FailListenerIdSet []*string `json:"FailListenerIdSet,omitempty" name:"FailListenerIdSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type BatchDeregisterTargetsResponse struct {
	*tchttp.BaseResponse
	Response *BatchDeregisterTargetsResponseParams `json:"Response"`
}

func (r *BatchDeregisterTargetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchDeregisterTargetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchModifyTargetWeightRequestParams struct {
	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// List of weights to be modified in batches
	ModifyList []*TargetsWeightRule `json:"ModifyList,omitempty" name:"ModifyList"`
}

type BatchModifyTargetWeightRequest struct {
	*tchttp.BaseRequest
	
	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// List of weights to be modified in batches
	ModifyList []*TargetsWeightRule `json:"ModifyList,omitempty" name:"ModifyList"`
}

func (r *BatchModifyTargetWeightRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchModifyTargetWeightRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "ModifyList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchModifyTargetWeightRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchModifyTargetWeightResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type BatchModifyTargetWeightResponse struct {
	*tchttp.BaseResponse
	Response *BatchModifyTargetWeightResponseParams `json:"Response"`
}

func (r *BatchModifyTargetWeightResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchModifyTargetWeightResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchRegisterTargetsRequestParams struct {
	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// Bound targets
	Targets []*BatchTarget `json:"Targets,omitempty" name:"Targets"`
}

type BatchRegisterTargetsRequest struct {
	*tchttp.BaseRequest
	
	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// Bound targets
	Targets []*BatchTarget `json:"Targets,omitempty" name:"Targets"`
}

func (r *BatchRegisterTargetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchRegisterTargetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "Targets")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchRegisterTargetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchRegisterTargetsResponseParams struct {
	// IDs of the listeners failed to be bound. If this parameter is empty, all listeners have been bound successfully.
	// Note: this field may return null, indicating that no valid values can be obtained.
	FailListenerIdSet []*string `json:"FailListenerIdSet,omitempty" name:"FailListenerIdSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type BatchRegisterTargetsResponse struct {
	*tchttp.BaseResponse
	Response *BatchRegisterTargetsResponseParams `json:"Response"`
}

func (r *BatchRegisterTargetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchRegisterTargetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BatchTarget struct {
	// Listener ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// Bound port
	// Note: this field may return null, indicating that no valid values can be obtained.
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// CVM instance ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// ENI IP
	// Note: this field may return null, indicating that no valid values can be obtained.
	EniIp *string `json:"EniIp,omitempty" name:"EniIp"`

	// Weight of the CVM instance. Value range: [0, 100]. If it is not specified for binding the instance, 10 will be used by default.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Weight *int64 `json:"Weight,omitempty" name:"Weight"`
}

type City struct {
	// City ID
	CityId *string `json:"CityId,omitempty" name:"CityId"`

	// City name
	CityName *string `json:"CityName,omitempty" name:"CityName"`
}

type Country struct {
	// Country/Region ID
	CountryId *string `json:"CountryId,omitempty" name:"CountryId"`

	// Country/Region name
	CountryName *string `json:"CountryName,omitempty" name:"CountryName"`
}

// Predefined struct for user
type CreateHaVipRequestParams struct {
	// VPC ID of the HAVIP.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Subnet ID of the HAVIP.
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// HAVIP name.
	HaVipName *string `json:"HaVipName,omitempty" name:"HaVipName"`

	// The specified virtual IP address, which must be within the IP range of the VPC and not in use. It will be automatically assigned if not specified.
	Vip *string `json:"Vip,omitempty" name:"Vip"`
}

type CreateHaVipRequest struct {
	*tchttp.BaseRequest
	
	// VPC ID of the HAVIP.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Subnet ID of the HAVIP.
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// HAVIP name.
	HaVipName *string `json:"HaVipName,omitempty" name:"HaVipName"`

	// The specified virtual IP address, which must be within the IP range of the VPC and not in use. It will be automatically assigned if not specified.
	Vip *string `json:"Vip,omitempty" name:"Vip"`
}

func (r *CreateHaVipRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateHaVipRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "HaVipName")
	delete(f, "Vip")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateHaVipRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateHaVipResponseParams struct {
	// HAVIP object.
	HaVip *HaVip `json:"HaVip,omitempty" name:"HaVip"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateHaVipResponse struct {
	*tchttp.BaseResponse
	Response *CreateHaVipResponseParams `json:"Response"`
}

func (r *CreateHaVipResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateHaVipResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateImageRequestParams struct {
	// Image name.
	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`

	// ID of the instance for which to make an image.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Image description.
	ImageDescription *string `json:"ImageDescription,omitempty" name:"ImageDescription"`

	// Whether to perform a forced shutdown to make an image. Valid values:
	// TRUE: yes
	// FALSE: no
	// Default value: FALSE.
	ForcePoweroff *string `json:"ForcePoweroff,omitempty" name:"ForcePoweroff"`
}

type CreateImageRequest struct {
	*tchttp.BaseRequest
	
	// Image name.
	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`

	// ID of the instance for which to make an image.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Image description.
	ImageDescription *string `json:"ImageDescription,omitempty" name:"ImageDescription"`

	// Whether to perform a forced shutdown to make an image. Valid values:
	// TRUE: yes
	// FALSE: no
	// Default value: FALSE.
	ForcePoweroff *string `json:"ForcePoweroff,omitempty" name:"ForcePoweroff"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateImageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateImageResponseParams struct {
	// Task ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

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
	// Key pair name, which can contain up to 25 digits, letters, and underscores.
	KeyName *string `json:"KeyName,omitempty" name:"KeyName"`
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
type CreateListenerRequestParams struct {
	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// Specifies for which ports to create listeners. Each port corresponds to a new listener
	Ports []*int64 `json:"Ports,omitempty" name:"Ports"`

	// Listener protocol. Valid values: TCP, UDP
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// List of names of the listeners to be created. The array of names and array of ports are in one-to-one correspondence. If you do not want to name them now, you do not need to provide this parameter.
	ListenerNames []*string `json:"ListenerNames,omitempty" name:"ListenerNames"`

	// Health check parameters
	HealthCheck *HealthCheck `json:"HealthCheck,omitempty" name:"HealthCheck"`

	// Session persistence time in seconds. Value range: 30–3600. The default value is 0, indicating that session persistence is not enabled. This parameter is applicable only to TCP/UDP listeners.
	SessionExpireTime *int64 `json:"SessionExpireTime,omitempty" name:"SessionExpireTime"`

	// Forwarding method of the listener. Valid values: WRR, LEAST_CONN.
	// They represent weighted round robin and least connections, respectively. Default value: WRR.
	Scheduler *string `json:"Scheduler,omitempty" name:"Scheduler"`

	// Session persistence type. Valid values: NORMAL: the default session persistence type; QUIC_CID: session persistence by QUIC connection ID. The `QUIC_CID` value can only be configured in UDP listeners. If this field is not specified, the default session persistence type will be used.
	SessionType *string `json:"SessionType,omitempty" name:"SessionType"`

	// End ports of port ranges, which must be the same as `Ports` in length.
	EndPorts []*int64 `json:"EndPorts,omitempty" name:"EndPorts"`
}

type CreateListenerRequest struct {
	*tchttp.BaseRequest
	
	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// Specifies for which ports to create listeners. Each port corresponds to a new listener
	Ports []*int64 `json:"Ports,omitempty" name:"Ports"`

	// Listener protocol. Valid values: TCP, UDP
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// List of names of the listeners to be created. The array of names and array of ports are in one-to-one correspondence. If you do not want to name them now, you do not need to provide this parameter.
	ListenerNames []*string `json:"ListenerNames,omitempty" name:"ListenerNames"`

	// Health check parameters
	HealthCheck *HealthCheck `json:"HealthCheck,omitempty" name:"HealthCheck"`

	// Session persistence time in seconds. Value range: 30–3600. The default value is 0, indicating that session persistence is not enabled. This parameter is applicable only to TCP/UDP listeners.
	SessionExpireTime *int64 `json:"SessionExpireTime,omitempty" name:"SessionExpireTime"`

	// Forwarding method of the listener. Valid values: WRR, LEAST_CONN.
	// They represent weighted round robin and least connections, respectively. Default value: WRR.
	Scheduler *string `json:"Scheduler,omitempty" name:"Scheduler"`

	// Session persistence type. Valid values: NORMAL: the default session persistence type; QUIC_CID: session persistence by QUIC connection ID. The `QUIC_CID` value can only be configured in UDP listeners. If this field is not specified, the default session persistence type will be used.
	SessionType *string `json:"SessionType,omitempty" name:"SessionType"`

	// End ports of port ranges, which must be the same as `Ports` in length.
	EndPorts []*int64 `json:"EndPorts,omitempty" name:"EndPorts"`
}

func (r *CreateListenerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateListenerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "Ports")
	delete(f, "Protocol")
	delete(f, "ListenerNames")
	delete(f, "HealthCheck")
	delete(f, "SessionExpireTime")
	delete(f, "Scheduler")
	delete(f, "SessionType")
	delete(f, "EndPorts")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateListenerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateListenerResponseParams struct {
	// Array of unique IDs of the created listeners
	// Note: this field may return null, indicating that no valid values can be obtained.
	ListenerIds []*string `json:"ListenerIds,omitempty" name:"ListenerIds"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateListenerResponse struct {
	*tchttp.BaseResponse
	Response *CreateListenerResponseParams `json:"Response"`
}

func (r *CreateListenerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateListenerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLoadBalancerRequestParams struct {
	// ECM region, such as `ap-xian-ecm`.
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// Network type of the CLB instance. Currently, you can pass in only `OPEN`, which indicates public network.
	LoadBalancerType *string `json:"LoadBalancerType,omitempty" name:"LoadBalancerType"`

	// CMCC: China Mobile; CTCC: China Telecom; CUCC: China Unicom.
	VipIsp *string `json:"VipIsp,omitempty" name:"VipIsp"`

	// CLB instance name, which will take effect only when one instance is created. It can contain 1–50 letters, digits, hyphens, and underscores.
	// Note: if the name of the new CLB instance already exists, a default name will be generated automatically.
	LoadBalancerName *string `json:"LoadBalancerName,omitempty" name:"LoadBalancerName"`

	// Network ID of the target device on the CLB backend, such as `vpc-12345678`.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Number of CLB instances to be created. Default value: 1.
	Number *int64 `json:"Number,omitempty" name:"Number"`

	// CLB information such as bandwidth limit.
	InternetAccessible *LoadBalancerInternetAccessible `json:"InternetAccessible,omitempty" name:"InternetAccessible"`

	// Tags.
	Tags []*TagInfo `json:"Tags,omitempty" name:"Tags"`

	// Security groups.
	SecurityGroups []*string `json:"SecurityGroups,omitempty" name:"SecurityGroups"`

	// IP version. Valid values: `IPV4` (default), `IPv6FullChain` (IPv6 version). This parameter is only for public network CLB instances.
	AddressIPVersion *string `json:"AddressIPVersion,omitempty" name:"AddressIPVersion"`

	// Subnet ID. This parameter is required for IPv6 CLB instances.
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
}

type CreateLoadBalancerRequest struct {
	*tchttp.BaseRequest
	
	// ECM region, such as `ap-xian-ecm`.
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// Network type of the CLB instance. Currently, you can pass in only `OPEN`, which indicates public network.
	LoadBalancerType *string `json:"LoadBalancerType,omitempty" name:"LoadBalancerType"`

	// CMCC: China Mobile; CTCC: China Telecom; CUCC: China Unicom.
	VipIsp *string `json:"VipIsp,omitempty" name:"VipIsp"`

	// CLB instance name, which will take effect only when one instance is created. It can contain 1–50 letters, digits, hyphens, and underscores.
	// Note: if the name of the new CLB instance already exists, a default name will be generated automatically.
	LoadBalancerName *string `json:"LoadBalancerName,omitempty" name:"LoadBalancerName"`

	// Network ID of the target device on the CLB backend, such as `vpc-12345678`.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Number of CLB instances to be created. Default value: 1.
	Number *int64 `json:"Number,omitempty" name:"Number"`

	// CLB information such as bandwidth limit.
	InternetAccessible *LoadBalancerInternetAccessible `json:"InternetAccessible,omitempty" name:"InternetAccessible"`

	// Tags.
	Tags []*TagInfo `json:"Tags,omitempty" name:"Tags"`

	// Security groups.
	SecurityGroups []*string `json:"SecurityGroups,omitempty" name:"SecurityGroups"`

	// IP version. Valid values: `IPV4` (default), `IPv6FullChain` (IPv6 version). This parameter is only for public network CLB instances.
	AddressIPVersion *string `json:"AddressIPVersion,omitempty" name:"AddressIPVersion"`

	// Subnet ID. This parameter is required for IPv6 CLB instances.
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
}

func (r *CreateLoadBalancerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLoadBalancerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EcmRegion")
	delete(f, "LoadBalancerType")
	delete(f, "VipIsp")
	delete(f, "LoadBalancerName")
	delete(f, "VpcId")
	delete(f, "Number")
	delete(f, "InternetAccessible")
	delete(f, "Tags")
	delete(f, "SecurityGroups")
	delete(f, "AddressIPVersion")
	delete(f, "SubnetId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateLoadBalancerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLoadBalancerResponseParams struct {
	// Array of CLB instance IDs
	// Note: this field may return null, indicating that no valid values can be obtained.
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitempty" name:"LoadBalancerIds"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateLoadBalancerResponse struct {
	*tchttp.BaseResponse
	Response *CreateLoadBalancerResponseParams `json:"Response"`
}

func (r *CreateLoadBalancerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLoadBalancerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateModuleRequestParams struct {
	// Module name, such as video live streaming module name. It cannot start with a space or exceed 60 characters.
	ModuleName *string `json:"ModuleName,omitempty" name:"ModuleName"`

	// Default bandwidth in Mbps. It cannot exceed the bandwidth range. For more information, see `DescribeConfig`.
	DefaultBandWidth *int64 `json:"DefaultBandWidth,omitempty" name:"DefaultBandWidth"`

	// Default image ID, such as `img-qsdf3ff2`.
	DefaultImageId *string `json:"DefaultImageId,omitempty" name:"DefaultImageId"`

	// Model ID.
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// Default system disk size in GB. It is 50 GB by default and cannot exceed the system disk size range. For more information, see `DescribeConfig`.
	DefaultSystemDiskSize *int64 `json:"DefaultSystemDiskSize,omitempty" name:"DefaultSystemDiskSize"`

	// Default data disk size in GB. It cannot exceed the system disk size range. For more information, see `DescribeConfig`.
	DefaultDataDiskSize *int64 `json:"DefaultDataDiskSize,omitempty" name:"DefaultDataDiskSize"`

	// Whether to disable IP direct access. Valid values:
	// true: yes
	// false: no
	CloseIpDirect *bool `json:"CloseIpDirect,omitempty" name:"CloseIpDirect"`

	// List of tags.
	TagSpecification []*TagSpecification `json:"TagSpecification,omitempty" name:"TagSpecification"`

	// List of default module security groups
	SecurityGroups []*string `json:"SecurityGroups,omitempty" name:"SecurityGroups"`

	// Default inbound bandwidth in Mbps. It cannot exceed the bandwidth range. For more information, see `DescribeConfig`.
	DefaultBandWidthIn *int64 `json:"DefaultBandWidthIn,omitempty" name:"DefaultBandWidthIn"`

	// Whether to prohibit public IP assignment
	DisableWanIp *bool `json:"DisableWanIp,omitempty" name:"DisableWanIp"`

	// System disk information.
	SystemDisk *SystemDisk `json:"SystemDisk,omitempty" name:"SystemDisk"`

	// Data disk information.
	DataDisks []*DataDisk `json:"DataDisks,omitempty" name:"DataDisks"`
}

type CreateModuleRequest struct {
	*tchttp.BaseRequest
	
	// Module name, such as video live streaming module name. It cannot start with a space or exceed 60 characters.
	ModuleName *string `json:"ModuleName,omitempty" name:"ModuleName"`

	// Default bandwidth in Mbps. It cannot exceed the bandwidth range. For more information, see `DescribeConfig`.
	DefaultBandWidth *int64 `json:"DefaultBandWidth,omitempty" name:"DefaultBandWidth"`

	// Default image ID, such as `img-qsdf3ff2`.
	DefaultImageId *string `json:"DefaultImageId,omitempty" name:"DefaultImageId"`

	// Model ID.
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// Default system disk size in GB. It is 50 GB by default and cannot exceed the system disk size range. For more information, see `DescribeConfig`.
	DefaultSystemDiskSize *int64 `json:"DefaultSystemDiskSize,omitempty" name:"DefaultSystemDiskSize"`

	// Default data disk size in GB. It cannot exceed the system disk size range. For more information, see `DescribeConfig`.
	DefaultDataDiskSize *int64 `json:"DefaultDataDiskSize,omitempty" name:"DefaultDataDiskSize"`

	// Whether to disable IP direct access. Valid values:
	// true: yes
	// false: no
	CloseIpDirect *bool `json:"CloseIpDirect,omitempty" name:"CloseIpDirect"`

	// List of tags.
	TagSpecification []*TagSpecification `json:"TagSpecification,omitempty" name:"TagSpecification"`

	// List of default module security groups
	SecurityGroups []*string `json:"SecurityGroups,omitempty" name:"SecurityGroups"`

	// Default inbound bandwidth in Mbps. It cannot exceed the bandwidth range. For more information, see `DescribeConfig`.
	DefaultBandWidthIn *int64 `json:"DefaultBandWidthIn,omitempty" name:"DefaultBandWidthIn"`

	// Whether to prohibit public IP assignment
	DisableWanIp *bool `json:"DisableWanIp,omitempty" name:"DisableWanIp"`

	// System disk information.
	SystemDisk *SystemDisk `json:"SystemDisk,omitempty" name:"SystemDisk"`

	// Data disk information.
	DataDisks []*DataDisk `json:"DataDisks,omitempty" name:"DataDisks"`
}

func (r *CreateModuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateModuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModuleName")
	delete(f, "DefaultBandWidth")
	delete(f, "DefaultImageId")
	delete(f, "InstanceType")
	delete(f, "DefaultSystemDiskSize")
	delete(f, "DefaultDataDiskSize")
	delete(f, "CloseIpDirect")
	delete(f, "TagSpecification")
	delete(f, "SecurityGroups")
	delete(f, "DefaultBandWidthIn")
	delete(f, "DisableWanIp")
	delete(f, "SystemDisk")
	delete(f, "DataDisks")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateModuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateModuleResponseParams struct {
	// Module ID. It is the ID assigned to a module after it is created successfully.
	ModuleId *string `json:"ModuleId,omitempty" name:"ModuleId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateModuleResponse struct {
	*tchttp.BaseResponse
	Response *CreateModuleResponseParams `json:"Response"`
}

func (r *CreateModuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateModuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNetworkInterfaceRequestParams struct {
	// VPC instance ID, which can be obtained from the `VpcId` field in the returned value of the `DescribeVpcs` API.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// ENI name, which can contain up to 60 bytes.
	NetworkInterfaceName *string `json:"NetworkInterfaceName,omitempty" name:"NetworkInterfaceName"`

	// Subnet instance ID of the ENI, such as 'subnet-0ap8nwca'.
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// ECM region, such as `ap-xian-ecm`.
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// ENI description. You can enter any information within 60 characters.
	NetworkInterfaceDescription *string `json:"NetworkInterfaceDescription,omitempty" name:"NetworkInterfaceDescription"`

	// Number of private IP addresses applied for. The total number of private IP addresses cannot exceed the quota.
	SecondaryPrivateIpAddressCount *uint64 `json:"SecondaryPrivateIpAddressCount,omitempty" name:"SecondaryPrivateIpAddressCount"`

	// The security group to be bound with, such as ['sg-1dd51d'].
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`

	// Information of the specified private IPs. You can specify up to 10 IPs at a time.
	PrivateIpAddresses []*PrivateIpAddressSpecification `json:"PrivateIpAddresses,omitempty" name:"PrivateIpAddresses"`

	// List of bound tags, such as [{"Key": "city", "Value": "shanghai"}].
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

type CreateNetworkInterfaceRequest struct {
	*tchttp.BaseRequest
	
	// VPC instance ID, which can be obtained from the `VpcId` field in the returned value of the `DescribeVpcs` API.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// ENI name, which can contain up to 60 bytes.
	NetworkInterfaceName *string `json:"NetworkInterfaceName,omitempty" name:"NetworkInterfaceName"`

	// Subnet instance ID of the ENI, such as 'subnet-0ap8nwca'.
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// ECM region, such as `ap-xian-ecm`.
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// ENI description. You can enter any information within 60 characters.
	NetworkInterfaceDescription *string `json:"NetworkInterfaceDescription,omitempty" name:"NetworkInterfaceDescription"`

	// Number of private IP addresses applied for. The total number of private IP addresses cannot exceed the quota.
	SecondaryPrivateIpAddressCount *uint64 `json:"SecondaryPrivateIpAddressCount,omitempty" name:"SecondaryPrivateIpAddressCount"`

	// The security group to be bound with, such as ['sg-1dd51d'].
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`

	// Information of the specified private IPs. You can specify up to 10 IPs at a time.
	PrivateIpAddresses []*PrivateIpAddressSpecification `json:"PrivateIpAddresses,omitempty" name:"PrivateIpAddresses"`

	// List of bound tags, such as [{"Key": "city", "Value": "shanghai"}].
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

func (r *CreateNetworkInterfaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNetworkInterfaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcId")
	delete(f, "NetworkInterfaceName")
	delete(f, "SubnetId")
	delete(f, "EcmRegion")
	delete(f, "NetworkInterfaceDescription")
	delete(f, "SecondaryPrivateIpAddressCount")
	delete(f, "SecurityGroupIds")
	delete(f, "PrivateIpAddresses")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateNetworkInterfaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNetworkInterfaceResponseParams struct {
	// ENI instance.
	NetworkInterface *NetworkInterface `json:"NetworkInterface,omitempty" name:"NetworkInterface"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateNetworkInterfaceResponse struct {
	*tchttp.BaseResponse
	Response *CreateNetworkInterfaceResponseParams `json:"Response"`
}

func (r *CreateNetworkInterfaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNetworkInterfaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRouteTableRequestParams struct {
	// ID of the VPC instance to be manipulated, which can be obtained from the `VpcId` field in the returned value of the `DescribeVpcs` API.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Route table name, which can contain up to 60 bytes.
	RouteTableName *string `json:"RouteTableName,omitempty" name:"RouteTableName"`

	// ECM region
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`
}

type CreateRouteTableRequest struct {
	*tchttp.BaseRequest
	
	// ID of the VPC instance to be manipulated, which can be obtained from the `VpcId` field in the returned value of the `DescribeVpcs` API.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Route table name, which can contain up to 60 bytes.
	RouteTableName *string `json:"RouteTableName,omitempty" name:"RouteTableName"`

	// ECM region
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`
}

func (r *CreateRouteTableRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRouteTableRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcId")
	delete(f, "RouteTableName")
	delete(f, "EcmRegion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRouteTableRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRouteTableResponseParams struct {
	// Route table object
	RouteTable *RouteTable `json:"RouteTable,omitempty" name:"RouteTable"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateRouteTableResponse struct {
	*tchttp.BaseResponse
	Response *CreateRouteTableResponseParams `json:"Response"`
}

func (r *CreateRouteTableResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRouteTableResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRoutesRequestParams struct {
	// Route table instance ID.
	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`

	// Routing policy object to be created.
	Routes []*Route `json:"Routes,omitempty" name:"Routes"`
}

type CreateRoutesRequest struct {
	*tchttp.BaseRequest
	
	// Route table instance ID.
	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`

	// Routing policy object to be created.
	Routes []*Route `json:"Routes,omitempty" name:"Routes"`
}

func (r *CreateRoutesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRoutesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RouteTableId")
	delete(f, "Routes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRoutesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRoutesResponseParams struct {
	// Number of added instances.
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Route table object.
	// Note: this field may return null, indicating that no valid values can be obtained.
	RouteTableSet []*RouteTable `json:"RouteTableSet,omitempty" name:"RouteTableSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateRoutesResponse struct {
	*tchttp.BaseResponse
	Response *CreateRoutesResponseParams `json:"Response"`
}

func (r *CreateRoutesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRoutesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSecurityGroupPoliciesRequestParams struct {
	// Security group instance ID, such as `esg-33ocnj9n`, which can be obtained through the `DescribeSecurityGroups` API.
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`

	// Security group policy set.
	SecurityGroupPolicySet *SecurityGroupPolicySet `json:"SecurityGroupPolicySet,omitempty" name:"SecurityGroupPolicySet"`
}

type CreateSecurityGroupPoliciesRequest struct {
	*tchttp.BaseRequest
	
	// Security group instance ID, such as `esg-33ocnj9n`, which can be obtained through the `DescribeSecurityGroups` API.
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`

	// Security group policy set.
	SecurityGroupPolicySet *SecurityGroupPolicySet `json:"SecurityGroupPolicySet,omitempty" name:"SecurityGroupPolicySet"`
}

func (r *CreateSecurityGroupPoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSecurityGroupPoliciesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecurityGroupId")
	delete(f, "SecurityGroupPolicySet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSecurityGroupPoliciesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSecurityGroupPoliciesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateSecurityGroupPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *CreateSecurityGroupPoliciesResponseParams `json:"Response"`
}

func (r *CreateSecurityGroupPoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSecurityGroupPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSecurityGroupRequestParams struct {
	// Security group name, which can be customized with up to 60 characters.
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// Security group remarks, which can contain up to 100 characters.
	GroupDescription *string `json:"GroupDescription,omitempty" name:"GroupDescription"`

	// List of bound tags, such as [{"Key": "city", "Value": "shanghai"}].
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

type CreateSecurityGroupRequest struct {
	*tchttp.BaseRequest
	
	// Security group name, which can be customized with up to 60 characters.
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// Security group remarks, which can contain up to 100 characters.
	GroupDescription *string `json:"GroupDescription,omitempty" name:"GroupDescription"`

	// List of bound tags, such as [{"Key": "city", "Value": "shanghai"}].
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

func (r *CreateSecurityGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSecurityGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupName")
	delete(f, "GroupDescription")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSecurityGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSecurityGroupResponseParams struct {
	// Security group object.
	SecurityGroup *SecurityGroup `json:"SecurityGroup,omitempty" name:"SecurityGroup"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateSecurityGroupResponse struct {
	*tchttp.BaseResponse
	Response *CreateSecurityGroupResponseParams `json:"Response"`
}

func (r *CreateSecurityGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSecurityGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSubnetRequestParams struct {
	// ID of the VPC instance to be manipulated, which can be obtained from the `VpcId` field in the returned value of the `DescribeVpcs` API.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Subnet name, which can contain up to 60 bytes.
	SubnetName *string `json:"SubnetName,omitempty" name:"SubnetName"`

	// Subnet IP address range. It must be within the VPC IP address range. Subnet IP address ranges cannot overlap with each other within the same VPC.
	CidrBlock *string `json:"CidrBlock,omitempty" name:"CidrBlock"`

	// AZ ID of the subnet. You can select different AZs for different subnets for cross-AZ disaster recovery.
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// ECM region
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// List of bound tags, such as [{"Key": "city", "Value": "shanghai"}].
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

type CreateSubnetRequest struct {
	*tchttp.BaseRequest
	
	// ID of the VPC instance to be manipulated, which can be obtained from the `VpcId` field in the returned value of the `DescribeVpcs` API.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Subnet name, which can contain up to 60 bytes.
	SubnetName *string `json:"SubnetName,omitempty" name:"SubnetName"`

	// Subnet IP address range. It must be within the VPC IP address range. Subnet IP address ranges cannot overlap with each other within the same VPC.
	CidrBlock *string `json:"CidrBlock,omitempty" name:"CidrBlock"`

	// AZ ID of the subnet. You can select different AZs for different subnets for cross-AZ disaster recovery.
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// ECM region
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// List of bound tags, such as [{"Key": "city", "Value": "shanghai"}].
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

func (r *CreateSubnetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSubnetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcId")
	delete(f, "SubnetName")
	delete(f, "CidrBlock")
	delete(f, "Zone")
	delete(f, "EcmRegion")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSubnetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSubnetResponseParams struct {
	// Subnet object.
	Subnet *Subnet `json:"Subnet,omitempty" name:"Subnet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateSubnetResponse struct {
	*tchttp.BaseResponse
	Response *CreateSubnetResponseParams `json:"Response"`
}

func (r *CreateSubnetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSubnetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVpcRequestParams struct {
	// VPC name, which can contain up to 60 bytes.
	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`

	// VPC CIDR block, which must fall within the following private network IP ranges: 10.*.0.0/16, 172.[16-31].0.0/16, and 192.168.0.0/16.
	CidrBlock *string `json:"CidrBlock,omitempty" name:"CidrBlock"`

	// ECM region
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// Whether multicast is enabled. true: enabled; false: disabled. This parameter is not supported currently
	EnableMulticast *string `json:"EnableMulticast,omitempty" name:"EnableMulticast"`

	// DNS address (not supported currently). Up to 4 addresses can be supported.
	DnsServers []*string `json:"DnsServers,omitempty" name:"DnsServers"`

	// Domain name. This parameter is not supported currently
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// List of bound tags, such as [{"Key": "city", "Value": "shanghai"}].
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// Description
	Description *string `json:"Description,omitempty" name:"Description"`
}

type CreateVpcRequest struct {
	*tchttp.BaseRequest
	
	// VPC name, which can contain up to 60 bytes.
	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`

	// VPC CIDR block, which must fall within the following private network IP ranges: 10.*.0.0/16, 172.[16-31].0.0/16, and 192.168.0.0/16.
	CidrBlock *string `json:"CidrBlock,omitempty" name:"CidrBlock"`

	// ECM region
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// Whether multicast is enabled. true: enabled; false: disabled. This parameter is not supported currently
	EnableMulticast *string `json:"EnableMulticast,omitempty" name:"EnableMulticast"`

	// DNS address (not supported currently). Up to 4 addresses can be supported.
	DnsServers []*string `json:"DnsServers,omitempty" name:"DnsServers"`

	// Domain name. This parameter is not supported currently
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// List of bound tags, such as [{"Key": "city", "Value": "shanghai"}].
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// Description
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *CreateVpcRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVpcRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcName")
	delete(f, "CidrBlock")
	delete(f, "EcmRegion")
	delete(f, "EnableMulticast")
	delete(f, "DnsServers")
	delete(f, "DomainName")
	delete(f, "Tags")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateVpcRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVpcResponseParams struct {
	// VPC object.
	Vpc *VpcInfo `json:"Vpc,omitempty" name:"Vpc"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateVpcResponse struct {
	*tchttp.BaseResponse
	Response *CreateVpcResponseParams `json:"Response"`
}

func (r *CreateVpcResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVpcResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DataDisk struct {
	// Data disk size in GB.
	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`

	// Data disk type. Valid values:
	// - LOCAL_BASIC: local disk
	// - CLOUD_PREMIUM: Premium Cloud Storage
	// 
	// Default value: LOCAL_BASIC.
	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`
}

// Predefined struct for user
type DeleteHaVipRequestParams struct {
	// Unique HAVIP ID, such as `havip-9o233uri`.
	HaVipId *string `json:"HaVipId,omitempty" name:"HaVipId"`
}

type DeleteHaVipRequest struct {
	*tchttp.BaseRequest
	
	// Unique HAVIP ID, such as `havip-9o233uri`.
	HaVipId *string `json:"HaVipId,omitempty" name:"HaVipId"`
}

func (r *DeleteHaVipRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteHaVipRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "HaVipId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteHaVipRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteHaVipResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteHaVipResponse struct {
	*tchttp.BaseResponse
	Response *DeleteHaVipResponseParams `json:"Response"`
}

func (r *DeleteHaVipResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteHaVipResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteImageRequestParams struct {
	// List of image IDs.
	ImageIDSet []*string `json:"ImageIDSet,omitempty" name:"ImageIDSet"`
}

type DeleteImageRequest struct {
	*tchttp.BaseRequest
	
	// List of image IDs.
	ImageIDSet []*string `json:"ImageIDSet,omitempty" name:"ImageIDSet"`
}

func (r *DeleteImageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteImageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageIDSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteImageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteImageResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteImageResponse struct {
	*tchttp.BaseResponse
	Response *DeleteImageResponseParams `json:"Response"`
}

func (r *DeleteImageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteListenerRequestParams struct {
	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// ID of the listener to be deleted
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`
}

type DeleteListenerRequest struct {
	*tchttp.BaseRequest
	
	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// ID of the listener to be deleted
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`
}

func (r *DeleteListenerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteListenerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "ListenerId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteListenerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteListenerResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteListenerResponse struct {
	*tchttp.BaseResponse
	Response *DeleteListenerResponseParams `json:"Response"`
}

func (r *DeleteListenerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteListenerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLoadBalancerListenersRequestParams struct {
	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// Array of IDs of the listeners to be deleted. If this parameter is left empty, all listeners of the CLB instance will be deleted
	ListenerIds []*string `json:"ListenerIds,omitempty" name:"ListenerIds"`
}

type DeleteLoadBalancerListenersRequest struct {
	*tchttp.BaseRequest
	
	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// Array of IDs of the listeners to be deleted. If this parameter is left empty, all listeners of the CLB instance will be deleted
	ListenerIds []*string `json:"ListenerIds,omitempty" name:"ListenerIds"`
}

func (r *DeleteLoadBalancerListenersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLoadBalancerListenersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "ListenerIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteLoadBalancerListenersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLoadBalancerListenersResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteLoadBalancerListenersResponse struct {
	*tchttp.BaseResponse
	Response *DeleteLoadBalancerListenersResponseParams `json:"Response"`
}

func (r *DeleteLoadBalancerListenersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLoadBalancerListenersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLoadBalancerRequestParams struct {
	// Array of IDs of the CLB instances to be deleted. Array length limit: 20
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitempty" name:"LoadBalancerIds"`
}

type DeleteLoadBalancerRequest struct {
	*tchttp.BaseRequest
	
	// Array of IDs of the CLB instances to be deleted. Array length limit: 20
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitempty" name:"LoadBalancerIds"`
}

func (r *DeleteLoadBalancerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLoadBalancerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteLoadBalancerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLoadBalancerResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteLoadBalancerResponse struct {
	*tchttp.BaseResponse
	Response *DeleteLoadBalancerResponseParams `json:"Response"`
}

func (r *DeleteLoadBalancerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLoadBalancerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteModuleRequestParams struct {
	// Module ID, such as `em-qn46snq8`.
	ModuleId *string `json:"ModuleId,omitempty" name:"ModuleId"`
}

type DeleteModuleRequest struct {
	*tchttp.BaseRequest
	
	// Module ID, such as `em-qn46snq8`.
	ModuleId *string `json:"ModuleId,omitempty" name:"ModuleId"`
}

func (r *DeleteModuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteModuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModuleId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteModuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteModuleResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteModuleResponse struct {
	*tchttp.BaseResponse
	Response *DeleteModuleResponseParams `json:"Response"`
}

func (r *DeleteModuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteModuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNetworkInterfaceRequestParams struct {
	// ENI instance ID, such as `eni-m6dyj72l`.
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`

	// ECM region, such as `ap-xian-ecm`.
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`
}

type DeleteNetworkInterfaceRequest struct {
	*tchttp.BaseRequest
	
	// ENI instance ID, such as `eni-m6dyj72l`.
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`

	// ECM region, such as `ap-xian-ecm`.
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`
}

func (r *DeleteNetworkInterfaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNetworkInterfaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NetworkInterfaceId")
	delete(f, "EcmRegion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteNetworkInterfaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNetworkInterfaceResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteNetworkInterfaceResponse struct {
	*tchttp.BaseResponse
	Response *DeleteNetworkInterfaceResponseParams `json:"Response"`
}

func (r *DeleteNetworkInterfaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNetworkInterfaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRouteTableRequestParams struct {
	// Route table instance ID, such as `rtb-azd4dt1c`
	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`
}

type DeleteRouteTableRequest struct {
	*tchttp.BaseRequest
	
	// Route table instance ID, such as `rtb-azd4dt1c`
	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`
}

func (r *DeleteRouteTableRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRouteTableRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RouteTableId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRouteTableRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRouteTableResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteRouteTableResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRouteTableResponseParams `json:"Response"`
}

func (r *DeleteRouteTableResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRouteTableResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRoutesRequestParams struct {
	// Unique route table ID.
	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`

	// Routing policy object.
	Routes []*Route `json:"Routes,omitempty" name:"Routes"`
}

type DeleteRoutesRequest struct {
	*tchttp.BaseRequest
	
	// Unique route table ID.
	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`

	// Routing policy object.
	Routes []*Route `json:"Routes,omitempty" name:"Routes"`
}

func (r *DeleteRoutesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRoutesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RouteTableId")
	delete(f, "Routes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRoutesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRoutesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteRoutesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRoutesResponseParams `json:"Response"`
}

func (r *DeleteRoutesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRoutesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSecurityGroupPoliciesRequestParams struct {
	// Security group instance ID, such as `esg-33ocnj9n`, which can be obtained through the `DescribeSecurityGroups` API.
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`

	// Security group policy set. You can only delete one or more policies in one direction in one request. Both PolicyIndex-matching deletion and security group policy-matching deletion methods are supported. You can use only one matching method in one request.
	SecurityGroupPolicySet *SecurityGroupPolicySet `json:"SecurityGroupPolicySet,omitempty" name:"SecurityGroupPolicySet"`
}

type DeleteSecurityGroupPoliciesRequest struct {
	*tchttp.BaseRequest
	
	// Security group instance ID, such as `esg-33ocnj9n`, which can be obtained through the `DescribeSecurityGroups` API.
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`

	// Security group policy set. You can only delete one or more policies in one direction in one request. Both PolicyIndex-matching deletion and security group policy-matching deletion methods are supported. You can use only one matching method in one request.
	SecurityGroupPolicySet *SecurityGroupPolicySet `json:"SecurityGroupPolicySet,omitempty" name:"SecurityGroupPolicySet"`
}

func (r *DeleteSecurityGroupPoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSecurityGroupPoliciesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecurityGroupId")
	delete(f, "SecurityGroupPolicySet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSecurityGroupPoliciesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSecurityGroupPoliciesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteSecurityGroupPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSecurityGroupPoliciesResponseParams `json:"Response"`
}

func (r *DeleteSecurityGroupPoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSecurityGroupPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSecurityGroupRequestParams struct {
	// Security group instance ID, such as `esg-33ocnj9n`, which can be obtained through the `DescribeSecurityGroups` API.
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
}

type DeleteSecurityGroupRequest struct {
	*tchttp.BaseRequest
	
	// Security group instance ID, such as `esg-33ocnj9n`, which can be obtained through the `DescribeSecurityGroups` API.
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
}

func (r *DeleteSecurityGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSecurityGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecurityGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSecurityGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSecurityGroupResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteSecurityGroupResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSecurityGroupResponseParams `json:"Response"`
}

func (r *DeleteSecurityGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSecurityGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSnapshotsRequestParams struct {
	// List of IDs of the snapshots to be deleted, which can be queried through [DescribeSnapshots](https://intl.cloud.tencent.com/document/product/362/15647?from_cn_redirect=1).
	SnapshotIds []*string `json:"SnapshotIds,omitempty" name:"SnapshotIds"`

	// Whether to force delete the images associated with the snapshot.
	DeleteBindImages *bool `json:"DeleteBindImages,omitempty" name:"DeleteBindImages"`
}

type DeleteSnapshotsRequest struct {
	*tchttp.BaseRequest
	
	// List of IDs of the snapshots to be deleted, which can be queried through [DescribeSnapshots](https://intl.cloud.tencent.com/document/product/362/15647?from_cn_redirect=1).
	SnapshotIds []*string `json:"SnapshotIds,omitempty" name:"SnapshotIds"`

	// Whether to force delete the images associated with the snapshot.
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
type DeleteSubnetRequestParams struct {
	// Subnet instance ID, which can be obtained from the `SubnetId` field in the returned value of the `DescribeSubnets` API.
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// ECM region
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`
}

type DeleteSubnetRequest struct {
	*tchttp.BaseRequest
	
	// Subnet instance ID, which can be obtained from the `SubnetId` field in the returned value of the `DescribeSubnets` API.
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// ECM region
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`
}

func (r *DeleteSubnetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSubnetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubnetId")
	delete(f, "EcmRegion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSubnetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSubnetResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteSubnetResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSubnetResponseParams `json:"Response"`
}

func (r *DeleteSubnetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSubnetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteVpcRequestParams struct {
	// VPC instance ID, which can be obtained from the `VpcId` field in the returned value of the `DescribeVpcs` API.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// ECM region
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`
}

type DeleteVpcRequest struct {
	*tchttp.BaseRequest
	
	// VPC instance ID, which can be obtained from the `VpcId` field in the returned value of the `DescribeVpcs` API.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// ECM region
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`
}

func (r *DeleteVpcRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteVpcRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcId")
	delete(f, "EcmRegion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteVpcRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteVpcResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteVpcResponse struct {
	*tchttp.BaseResponse
	Response *DeleteVpcResponseParams `json:"Response"`
}

func (r *DeleteVpcResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteVpcResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAddressQuotaRequestParams struct {
	// ECM region
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`
}

type DescribeAddressQuotaRequest struct {
	*tchttp.BaseRequest
	
	// ECM region
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`
}

func (r *DescribeAddressQuotaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAddressQuotaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EcmRegion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAddressQuotaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAddressQuotaResponseParams struct {
	// Quota information of EIPs in the account.
	QuotaSet []*EipQuota `json:"QuotaSet,omitempty" name:"QuotaSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAddressQuotaResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAddressQuotaResponseParams `json:"Response"`
}

func (r *DescribeAddressQuotaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAddressQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAddressesRequestParams struct {
	// ECM region
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// List of unique IDs of EIPs, such as `eip-11112222`. `AddressIds` and `Filters` cannot be specified at the same time.
	AddressIds []*string `json:"AddressIds,omitempty" name:"AddressIds"`

	// Each request can contain up to 10 `Filters` and 5 `Filter.Values`. `AddressIds` and `Filters` cannot be specified at the same time. The detailed filters are as follows:
	// address-id - String - Required: no - (Filter) Filter by unique EIP ID, such as `eip-11112222`.
	// address-name - String - Required: no - (Filter) Filter by EIP name. Fuzzy filtering is not supported.
	// address-ip - String - Required: no - (Filter) Filter by EIP IP address.
	// address-status - String - Required: no - (Filter) Filter by EIP status. Value range: see the list of EIP status.
	// instance-id - String - Required: no - (Filter) Filter by the ID of the instance bound to the EIP, such as `ins-11112222`.
	// private-ip-address - String - Required: no - (Filter) Filter by the private IP bound to the EIP.
	// network-interface-id - String - Required: no - (Filter) Filter by ID of the ENI bound to the EIP, such as `eni-11112222`.
	// is-arrears - String - Required: no - (Filter) Filter by whether the EIP is overdue (TRUE: the EIP is overdue | FALSE: the billing status of the EIP is normal)
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of results to be returned. Default value: 20. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeAddressesRequest struct {
	*tchttp.BaseRequest
	
	// ECM region
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// List of unique IDs of EIPs, such as `eip-11112222`. `AddressIds` and `Filters` cannot be specified at the same time.
	AddressIds []*string `json:"AddressIds,omitempty" name:"AddressIds"`

	// Each request can contain up to 10 `Filters` and 5 `Filter.Values`. `AddressIds` and `Filters` cannot be specified at the same time. The detailed filters are as follows:
	// address-id - String - Required: no - (Filter) Filter by unique EIP ID, such as `eip-11112222`.
	// address-name - String - Required: no - (Filter) Filter by EIP name. Fuzzy filtering is not supported.
	// address-ip - String - Required: no - (Filter) Filter by EIP IP address.
	// address-status - String - Required: no - (Filter) Filter by EIP status. Value range: see the list of EIP status.
	// instance-id - String - Required: no - (Filter) Filter by the ID of the instance bound to the EIP, such as `ins-11112222`.
	// private-ip-address - String - Required: no - (Filter) Filter by the private IP bound to the EIP.
	// network-interface-id - String - Required: no - (Filter) Filter by ID of the ENI bound to the EIP, such as `eni-11112222`.
	// is-arrears - String - Required: no - (Filter) Filter by whether the EIP is overdue (TRUE: the EIP is overdue | FALSE: the billing status of the EIP is normal)
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of results to be returned. Default value: 20. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeAddressesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAddressesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EcmRegion")
	delete(f, "AddressIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAddressesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAddressesResponseParams struct {
	// Number of eligible EIPs.
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// List of EIP details.
	AddressSet []*Address `json:"AddressSet,omitempty" name:"AddressSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAddressesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAddressesResponseParams `json:"Response"`
}

func (r *DescribeAddressesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAddressesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBaseOverviewRequestParams struct {

}

type DescribeBaseOverviewRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeBaseOverviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBaseOverviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBaseOverviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBaseOverviewResponseParams struct {
	// Number of modules
	ModuleNum *int64 `json:"ModuleNum,omitempty" name:"ModuleNum"`

	// Number of nodes
	NodeNum *int64 `json:"NodeNum,omitempty" name:"NodeNum"`

	// Number of CPU cores
	VcpuNum *int64 `json:"VcpuNum,omitempty" name:"VcpuNum"`

	// Memory size in GB
	MemoryNum *int64 `json:"MemoryNum,omitempty" name:"MemoryNum"`

	// Disk size in GB
	StorageNum *int64 `json:"StorageNum,omitempty" name:"StorageNum"`

	// Yesterday's network peak in Mbps
	NetworkNum *int64 `json:"NetworkNum,omitempty" name:"NetworkNum"`

	// Number of instances
	InstanceNum *int64 `json:"InstanceNum,omitempty" name:"InstanceNum"`

	// Number of running instances
	RunningNum *int64 `json:"RunningNum,omitempty" name:"RunningNum"`

	// Number of isolated instances
	IsolationNum *int64 `json:"IsolationNum,omitempty" name:"IsolationNum"`

	// Number of expired instances
	ExpiredNum *int64 `json:"ExpiredNum,omitempty" name:"ExpiredNum"`

	// Number of instances about to expire
	WillExpireNum *int64 `json:"WillExpireNum,omitempty" name:"WillExpireNum"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBaseOverviewResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBaseOverviewResponseParams `json:"Response"`
}

func (r *DescribeBaseOverviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBaseOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConfigRequestParams struct {

}

type DescribeConfigRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConfigResponseParams struct {
	// Range of the network bandwidth disk size.
	NetworkStorageRange *NetworkStorageRange `json:"NetworkStorageRange,omitempty" name:"NetworkStorageRange"`

	// Image OS allowlist.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ImageWhiteSet []*string `json:"ImageWhiteSet,omitempty" name:"ImageWhiteSet"`

	// Network quota information.
	// Note: this field may return null, indicating that no valid values can be obtained.
	InstanceNetworkLimitConfigs []*InstanceNetworkLimitConfig `json:"InstanceNetworkLimitConfigs,omitempty" name:"InstanceNetworkLimitConfigs"`

	// Image quota information.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ImageLimits *ImageLimitConfig `json:"ImageLimits,omitempty" name:"ImageLimits"`

	// Default IP direct access, used in scenarios with direct access parameters such as module creation and virtual machine purchase.
	DefaultIPDirect *bool `json:"DefaultIPDirect,omitempty" name:"DefaultIPDirect"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeConfigResponseParams `json:"Response"`
}

func (r *DescribeConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCustomImageTaskRequestParams struct {
	// Supports querying by key and value.
	// task-id: async task ID
	// image-id: image ID
	// image-name: image name
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeCustomImageTaskRequest struct {
	*tchttp.BaseRequest
	
	// Supports querying by key and value.
	// task-id: async task ID
	// image-id: image ID
	// image-name: image name
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeCustomImageTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCustomImageTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCustomImageTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCustomImageTaskResponseParams struct {
	// Import task details
	// Note: this field may return null, indicating that no valid values can be obtained.
	ImageTaskSet []*ImageTask `json:"ImageTaskSet,omitempty" name:"ImageTaskSet"`

	// Total number
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCustomImageTaskResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCustomImageTaskResponseParams `json:"Response"`
}

func (r *DescribeCustomImageTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCustomImageTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDefaultSubnetRequestParams struct {
	// ECM region
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// ECM AZ
	Zone *string `json:"Zone,omitempty" name:"Zone"`
}

type DescribeDefaultSubnetRequest struct {
	*tchttp.BaseRequest
	
	// ECM region
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// ECM AZ
	Zone *string `json:"Zone,omitempty" name:"Zone"`
}

func (r *DescribeDefaultSubnetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDefaultSubnetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EcmRegion")
	delete(f, "Zone")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDefaultSubnetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDefaultSubnetResponseParams struct {
	// Default subnet information. If there is no subnet, this parameter will be empty.
	Subnet *Subnet `json:"Subnet,omitempty" name:"Subnet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDefaultSubnetResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDefaultSubnetResponseParams `json:"Response"`
}

func (r *DescribeDefaultSubnetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDefaultSubnetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHaVipsRequestParams struct {
	// Array of unique HAVIP IDs, such as `havip-9o233uri`.
	HaVipIds []*string `json:"HaVipIds,omitempty" name:"HaVipIds"`

	// Filter. `HaVipIds` and `Filters` cannot be specified at the same time.
	// havip-id - String - Unique HAVIP ID, such as `havip-9o233uri`.
	// havip-name - String - HAVIP name.
	// vpc-id - String - VPC ID of the HAVIP.
	// subnet-id - String - Subnet ID of the HAVIP.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// ECM region. If this parameter is left empty, it will indicate all regions.
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`
}

type DescribeHaVipsRequest struct {
	*tchttp.BaseRequest
	
	// Array of unique HAVIP IDs, such as `havip-9o233uri`.
	HaVipIds []*string `json:"HaVipIds,omitempty" name:"HaVipIds"`

	// Filter. `HaVipIds` and `Filters` cannot be specified at the same time.
	// havip-id - String - Unique HAVIP ID, such as `havip-9o233uri`.
	// havip-name - String - HAVIP name.
	// vpc-id - String - VPC ID of the HAVIP.
	// subnet-id - String - Subnet ID of the HAVIP.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// ECM region. If this parameter is left empty, it will indicate all regions.
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`
}

func (r *DescribeHaVipsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHaVipsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "HaVipIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "EcmRegion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeHaVipsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHaVipsResponseParams struct {
	// Number of eligible objects.
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Array of HAVIP objects.
	// Note: this field may return null, indicating that no valid values can be obtained.
	HaVipSet []*HaVip `json:"HaVipSet,omitempty" name:"HaVipSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeHaVipsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeHaVipsResponseParams `json:"Response"`
}

func (r *DescribeHaVipsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHaVipsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImageRequestParams struct {
	// Filter. Each request can contain up to 10 `Filters`. The detailed filters are as follows:
	// image-id - String - Required: no - (Filter) Filter by image ID.
	// image-type - String - Required: no - (Filter) Filter by image type. Valid values:
	// PRIVATE_IMAGE: private image created by the current account 
	// PUBLIC_IMAGE: public image created by Tencent Cloud
	// instance-type -String - Required: no - (Filter) Filter supported images by model.
	// image-name - String - Required: no - (Filter) Fuzzy match by image name. You can provide only one value.
	// image-os - String - Required: no - (Filter) Fuzzy match by image system name. You can provide only one value.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// Offset. Default value: 0. For more information on `Offset`, see the relevant section of the API overview.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100. For more information on `Limit`, see the relevant section of the API overview.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeImageRequest struct {
	*tchttp.BaseRequest
	
	// Filter. Each request can contain up to 10 `Filters`. The detailed filters are as follows:
	// image-id - String - Required: no - (Filter) Filter by image ID.
	// image-type - String - Required: no - (Filter) Filter by image type. Valid values:
	// PRIVATE_IMAGE: private image created by the current account 
	// PUBLIC_IMAGE: public image created by Tencent Cloud
	// instance-type -String - Required: no - (Filter) Filter supported images by model.
	// image-name - String - Required: no - (Filter) Fuzzy match by image name. You can provide only one value.
	// image-os - String - Required: no - (Filter) Fuzzy match by image system name. You can provide only one value.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// Offset. Default value: 0. For more information on `Offset`, see the relevant section of the API overview.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100. For more information on `Limit`, see the relevant section of the API overview.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeImageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeImageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImageResponseParams struct {
	// Total number of images
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Image array
	// Note: this field may return null, indicating that no valid values can be obtained.
	ImageSet []*Image `json:"ImageSet,omitempty" name:"ImageSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeImageResponse struct {
	*tchttp.BaseResponse
	Response *DescribeImageResponseParams `json:"Response"`
}

func (r *DescribeImageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImageResponse) FromJsonString(s string) error {
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
	// Supported OS types of imported images.
	ImportImageOsListSupported *ImageOsList `json:"ImportImageOsListSupported,omitempty" name:"ImportImageOsListSupported"`

	// Supported OS versions of imported images.
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
type DescribeInstanceTypeConfigRequestParams struct {

}

type DescribeInstanceTypeConfigRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeInstanceTypeConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceTypeConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceTypeConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceTypeConfigResponseParams struct {
	// Total number
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Model configuration information
	// Note: this field may return null, indicating that no valid values can be obtained.
	InstanceTypeConfigSet []*InstanceTypeConfig `json:"InstanceTypeConfigSet,omitempty" name:"InstanceTypeConfigSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeInstanceTypeConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceTypeConfigResponseParams `json:"Response"`
}

func (r *DescribeInstanceTypeConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceTypeConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceVncUrlRequestParams struct {
	// Instance ID, which can be obtained from the `InstanceId` field in the returned value of the `DescribeInstances` API.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type DescribeInstanceVncUrlRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID, which can be obtained from the `InstanceId` field in the returned value of the `DescribeInstances` API.
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

// Predefined struct for user
type DescribeInstanceVncUrlResponseParams struct {
	// Instance VNC URL.
	InstanceVncUrl *string `json:"InstanceVncUrl,omitempty" name:"InstanceVncUrl"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// None
	InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet"`
}

type DescribeInstancesDeniedActionsRequest struct {
	*tchttp.BaseRequest
	
	// None
	InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet"`
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
	delete(f, "InstanceIdSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstancesDeniedActionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesDeniedActionsResponseParams struct {
	// Prohibited operations for the instance
	InstanceOperatorSet []*InstanceOperator `json:"InstanceOperatorSet,omitempty" name:"InstanceOperatorSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeInstancesRequestParams struct {
	// Filter.
	// zone      String      Required: no     (Filter) Filter by AZ abbreviation.
	// zone-name      String      Required: no     (Filter) Filter by AZ name. Fuzzy match is supported.
	// module-id      String      Required: no     (Filter) Filter by module ID.
	// instance-id      String      Required: no      (Filter) Filter by instance ID.
	// instance-name      String      Required: no      (Filter) Filter by instance name. Fuzzy match is supported.
	// ip-address      String      Required: no      (Filter) Filter by the instance's private/public IP.
	// instance-uuid   string Required: no (Filter) Filter instances by `uuid`.
	// instance-state  string  Required: no (Filter) Update instances by instance status.
	// internet-service-provider      String      Required: no      (Filter) Filter by the ISP of the instance's public IP.
	// tag-key      String      Required: no      (Filter) Filter by tag key.
	// tag:tag-key      String      Required: no      (Filter) Filter by tag key-value pair. Replace `tag-key` with the specific tag key.
	// instance-family      String      Required: no      (Filter) Filter by model family.
	// module-name      String      Required: no      (Filter) Filter by module name. Fuzzy match is supported.
	// image-id      String      Required: no      (Filter) Filter by instance image ID.
	// vpc-id String      Required: no      (Filter) Filter by instance VPC ID.
	// subnet-id String      Required: no      (Filter) Filter by instance subnet ID.
	// 
	// If the `Filters` parameter is not specified, the information of all relevant instances will be queried.
	// Each request can contain up to 5 `Filter.Values`.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned results. Default value: 20 (if the number of query results is greater than or equal to 20). Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Specified sort by field. Currently, valid values are as follows:
	// timestamp: sort by instance creation time.
	// Note: you can sort only by creation time currently. More sort criteria may be supported in the future.
	// If this parameter is not specified, instances will be sorted by creation time by default.
	OrderByField *string `json:"OrderByField,omitempty" name:"OrderByField"`

	// Sorting order. 0: descending; 1: ascending. If this parameter is not specified, it will be descending by default.
	OrderDirection *int64 `json:"OrderDirection,omitempty" name:"OrderDirection"`
}

type DescribeInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Filter.
	// zone      String      Required: no     (Filter) Filter by AZ abbreviation.
	// zone-name      String      Required: no     (Filter) Filter by AZ name. Fuzzy match is supported.
	// module-id      String      Required: no     (Filter) Filter by module ID.
	// instance-id      String      Required: no      (Filter) Filter by instance ID.
	// instance-name      String      Required: no      (Filter) Filter by instance name. Fuzzy match is supported.
	// ip-address      String      Required: no      (Filter) Filter by the instance's private/public IP.
	// instance-uuid   string Required: no (Filter) Filter instances by `uuid`.
	// instance-state  string  Required: no (Filter) Update instances by instance status.
	// internet-service-provider      String      Required: no      (Filter) Filter by the ISP of the instance's public IP.
	// tag-key      String      Required: no      (Filter) Filter by tag key.
	// tag:tag-key      String      Required: no      (Filter) Filter by tag key-value pair. Replace `tag-key` with the specific tag key.
	// instance-family      String      Required: no      (Filter) Filter by model family.
	// module-name      String      Required: no      (Filter) Filter by module name. Fuzzy match is supported.
	// image-id      String      Required: no      (Filter) Filter by instance image ID.
	// vpc-id String      Required: no      (Filter) Filter by instance VPC ID.
	// subnet-id String      Required: no      (Filter) Filter by instance subnet ID.
	// 
	// If the `Filters` parameter is not specified, the information of all relevant instances will be queried.
	// Each request can contain up to 5 `Filter.Values`.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned results. Default value: 20 (if the number of query results is greater than or equal to 20). Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Specified sort by field. Currently, valid values are as follows:
	// timestamp: sort by instance creation time.
	// Note: you can sort only by creation time currently. More sort criteria may be supported in the future.
	// If this parameter is not specified, instances will be sorted by creation time by default.
	OrderByField *string `json:"OrderByField,omitempty" name:"OrderByField"`

	// Sorting order. 0: descending; 1: ascending. If this parameter is not specified, it will be descending by default.
	OrderDirection *int64 `json:"OrderDirection,omitempty" name:"OrderDirection"`
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
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "OrderByField")
	delete(f, "OrderDirection")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesResponseParams struct {
	// Length of the list of the returned instance information.
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// List of the returned instance information.
	// Note: this field may return null, indicating that no valid values can be obtained.
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
type DescribeListenersRequestParams struct {
	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// Array of IDs of the CLB listeners to be queried
	ListenerIds []*string `json:"ListenerIds,omitempty" name:"ListenerIds"`

	// Protocol type of the listener to be queried. Valid values: TCP, UDP.
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Port of the listener to be queried
	Port *int64 `json:"Port,omitempty" name:"Port"`
}

type DescribeListenersRequest struct {
	*tchttp.BaseRequest
	
	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// Array of IDs of the CLB listeners to be queried
	ListenerIds []*string `json:"ListenerIds,omitempty" name:"ListenerIds"`

	// Protocol type of the listener to be queried. Valid values: TCP, UDP.
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Port of the listener to be queried
	Port *int64 `json:"Port,omitempty" name:"Port"`
}

func (r *DescribeListenersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListenersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "ListenerIds")
	delete(f, "Protocol")
	delete(f, "Port")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeListenersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeListenersResponseParams struct {
	// List of listeners
	// Note: this field may return null, indicating that no valid values can be obtained.
	Listeners []*Listener `json:"Listeners,omitempty" name:"Listeners"`

	// Total number of listeners
	// Note: this field may return null, indicating that no valid values can be obtained.
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeListenersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeListenersResponseParams `json:"Response"`
}

func (r *DescribeListenersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListenersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLoadBalanceTaskStatusRequestParams struct {
	// Request ID, i.e., the `RequestId` parameter returned by the API
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

type DescribeLoadBalanceTaskStatusRequest struct {
	*tchttp.BaseRequest
	
	// Request ID, i.e., the `RequestId` parameter returned by the API
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeLoadBalanceTaskStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLoadBalanceTaskStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLoadBalanceTaskStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLoadBalanceTaskStatusResponseParams struct {
	// Current task status. 0: succeeded; 1: failed; 2: in progress.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeLoadBalanceTaskStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLoadBalanceTaskStatusResponseParams `json:"Response"`
}

func (r *DescribeLoadBalanceTaskStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLoadBalanceTaskStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLoadBalancersRequestParams struct {
	// Region. If this parameter is not specified, the information of all regions will be queried by default.
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// CLB instance ID.
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitempty" name:"LoadBalancerIds"`

	// CLB instance name.
	LoadBalancerName *string `json:"LoadBalancerName,omitempty" name:"LoadBalancerName"`

	// VIP address of the CLB instance. There can be multiple addresses.
	LoadBalancerVips []*string `json:"LoadBalancerVips,omitempty" name:"LoadBalancerVips"`

	// Private IP of the real server bound to the CLB.
	BackendPrivateIps []*string `json:"BackendPrivateIps,omitempty" name:"BackendPrivateIps"`

	// Data offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned CLB instances. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Whether the CLB instance is bound to a real server. 0: no; 1: yes; -1: query all. 
	// If this parameter is not specified, all will be queried by default.
	WithBackend *int64 `json:"WithBackend,omitempty" name:"WithBackend"`

	// Unique VPC ID of the CLB instance, such as `vpc-bhqkbhdx`.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Each request can contain up to 10 `Filters` and 100 `Filter.Values`. The detailed filters are as follows:
	// tag-key - String - Required: no - (Filter) Filter by tag key.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// Security group.
	SecurityGroup *string `json:"SecurityGroup,omitempty" name:"SecurityGroup"`
}

type DescribeLoadBalancersRequest struct {
	*tchttp.BaseRequest
	
	// Region. If this parameter is not specified, the information of all regions will be queried by default.
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// CLB instance ID.
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitempty" name:"LoadBalancerIds"`

	// CLB instance name.
	LoadBalancerName *string `json:"LoadBalancerName,omitempty" name:"LoadBalancerName"`

	// VIP address of the CLB instance. There can be multiple addresses.
	LoadBalancerVips []*string `json:"LoadBalancerVips,omitempty" name:"LoadBalancerVips"`

	// Private IP of the real server bound to the CLB.
	BackendPrivateIps []*string `json:"BackendPrivateIps,omitempty" name:"BackendPrivateIps"`

	// Data offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned CLB instances. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Whether the CLB instance is bound to a real server. 0: no; 1: yes; -1: query all. 
	// If this parameter is not specified, all will be queried by default.
	WithBackend *int64 `json:"WithBackend,omitempty" name:"WithBackend"`

	// Unique VPC ID of the CLB instance, such as `vpc-bhqkbhdx`.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Each request can contain up to 10 `Filters` and 100 `Filter.Values`. The detailed filters are as follows:
	// tag-key - String - Required: no - (Filter) Filter by tag key.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// Security group.
	SecurityGroup *string `json:"SecurityGroup,omitempty" name:"SecurityGroup"`
}

func (r *DescribeLoadBalancersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLoadBalancersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EcmRegion")
	delete(f, "LoadBalancerIds")
	delete(f, "LoadBalancerName")
	delete(f, "LoadBalancerVips")
	delete(f, "BackendPrivateIps")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "WithBackend")
	delete(f, "VpcId")
	delete(f, "Filters")
	delete(f, "SecurityGroup")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLoadBalancersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLoadBalancersResponseParams struct {
	// Total number of eligible CLB instances. This value is independent of the `Limit` in the input parameter.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Array of returned CLB instances.
	// Note: this field may return null, indicating that no valid values can be obtained.
	LoadBalancerSet []*LoadBalancer `json:"LoadBalancerSet,omitempty" name:"LoadBalancerSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeLoadBalancersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLoadBalancersResponseParams `json:"Response"`
}

func (r *DescribeLoadBalancersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLoadBalancersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModuleDetailRequestParams struct {
	// Module ID, such as `em-qn46snq8`.
	ModuleId *string `json:"ModuleId,omitempty" name:"ModuleId"`
}

type DescribeModuleDetailRequest struct {
	*tchttp.BaseRequest
	
	// Module ID, such as `em-qn46snq8`.
	ModuleId *string `json:"ModuleId,omitempty" name:"ModuleId"`
}

func (r *DescribeModuleDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModuleDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModuleId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeModuleDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModuleDetailResponseParams struct {
	// Module details. For more information, see `ModuleInfo` in the data structure.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Module *Module `json:"Module,omitempty" name:"Module"`

	// Module statistics. For more information, see `ModuleCounterInfo` in the data structure.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ModuleCounter *ModuleCounter `json:"ModuleCounter,omitempty" name:"ModuleCounter"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeModuleDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeModuleDetailResponseParams `json:"Response"`
}

func (r *DescribeModuleDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModuleDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModuleRequestParams struct {
	// Filter.
	// module-name - string - Required: no - (Filter) Filter by module name.
	// module-id - string - Required: no - (Filter) Filter by module ID.
	// image-id      String      Required: no      (Filter) Filter by image ID.
	// instance-family      String      Required: no      (Filter) Filter by model family.
	// security-group-id - string Required: no - (Filter) Filter by ID of the security group bound to the module.
	// Each request can contain up to 10 `Filters` and 5 `Filter.Values`.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// Offset. Default value: 0. For more information on `Offset`, see the relevant section of the API overview.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100. For more information on `Limit`, see the relevant section of the API overview.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Specified sort by field. Currently, valid values are as follows:
	// instance-num: sort by the number of instances.
	// node-num: sort by the number of nodes.
	// timestamp: sort by instance creation time.
	// If this parameter is not specified, instances will be sorted by creation time by default.
	OrderByField *string `json:"OrderByField,omitempty" name:"OrderByField"`

	// Sorting order. 0: descending; 1: ascending. If this parameter is not specified, it will be descending by default.
	OrderDirection *int64 `json:"OrderDirection,omitempty" name:"OrderDirection"`
}

type DescribeModuleRequest struct {
	*tchttp.BaseRequest
	
	// Filter.
	// module-name - string - Required: no - (Filter) Filter by module name.
	// module-id - string - Required: no - (Filter) Filter by module ID.
	// image-id      String      Required: no      (Filter) Filter by image ID.
	// instance-family      String      Required: no      (Filter) Filter by model family.
	// security-group-id - string Required: no - (Filter) Filter by ID of the security group bound to the module.
	// Each request can contain up to 10 `Filters` and 5 `Filter.Values`.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// Offset. Default value: 0. For more information on `Offset`, see the relevant section of the API overview.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100. For more information on `Limit`, see the relevant section of the API overview.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Specified sort by field. Currently, valid values are as follows:
	// instance-num: sort by the number of instances.
	// node-num: sort by the number of nodes.
	// timestamp: sort by instance creation time.
	// If this parameter is not specified, instances will be sorted by creation time by default.
	OrderByField *string `json:"OrderByField,omitempty" name:"OrderByField"`

	// Sorting order. 0: descending; 1: ascending. If this parameter is not specified, it will be descending by default.
	OrderDirection *int64 `json:"OrderDirection,omitempty" name:"OrderDirection"`
}

func (r *DescribeModuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "OrderByField")
	delete(f, "OrderDirection")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeModuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModuleResponseParams struct {
	// Number of eligible modules.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// List of module details.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ModuleItemSet []*ModuleItem `json:"ModuleItemSet,omitempty" name:"ModuleItemSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeModuleResponse struct {
	*tchttp.BaseResponse
	Response *DescribeModuleResponseParams `json:"Response"`
}

func (r *DescribeModuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMonthPeakNetworkRequestParams struct {
	// Month (xxxx-xx), such as `2021-03`. Default value: the last month
	Month *string `json:"Month,omitempty" name:"Month"`

	// Filter
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeMonthPeakNetworkRequest struct {
	*tchttp.BaseRequest
	
	// Month (xxxx-xx), such as `2021-03`. Default value: the last month
	Month *string `json:"Month,omitempty" name:"Month"`

	// Filter
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeMonthPeakNetworkRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMonthPeakNetworkRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Month")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMonthPeakNetworkRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMonthPeakNetworkResponseParams struct {
	// None
	// Note: this field may return null, indicating that no valid values can be obtained.
	MonthNetWorkData []*MonthNetwork `json:"MonthNetWorkData,omitempty" name:"MonthNetWorkData"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeMonthPeakNetworkResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMonthPeakNetworkResponseParams `json:"Response"`
}

func (r *DescribeMonthPeakNetworkResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMonthPeakNetworkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNetworkInterfacesRequestParams struct {
	// Queries the ID of the ENI instance, such as `eni-pxir56ns`. Each request supports a maximum of 100 instances. `NetworkInterfaceIds` and `Filters` cannot be specified at the same time.
	NetworkInterfaceIds []*string `json:"NetworkInterfaceIds,omitempty" name:"NetworkInterfaceIds"`

	// Filter. `NetworkInterfaceIds` and `Filters` cannot be specified at the same time.
	// vpc-id - String - (Filter) VPC instance ID, such as `vpc-f49l6u0z`.
	// subnet-id - String - (Filter) Subnet instance ID, such as `subnet-f49l6u0z`.
	// network-interface-id - String - (Filter) ENI instance ID, such as `eni-5k56k7k7`.
	// attachment.instance-id - String - (Filter) ID of the bound CVM instance, such as `ein-3nqpdn3i`.
	// groups.security-group-id - String - (Filter) ID of the bound security group instance, such as `sg-f9ekbxeq`.
	// network-interface-name - String - (Filter) ENI instance name.
	// network-interface-description - String - (Filter) ENI instance description.
	// address-ip - String - (Filter) Private IPv4 address.
	// tag-key - String - Required: no - (Filter) Filter by tag key. For directions, see Sample 2.
	// tag:tag-key - String - Required: no - (Filter) Filter by tag key-value pair. Replace `tag-key` with the specific tag key. For directions, see Sample 3.
	// is-primary - Boolean - Required: no - (Filter) Filter by whether it is a primary ENI. true: filter only by primary ENI; false: filter only by secondary ENI. If this parameter is not specified, filtering by both primary and secondary ENIs will be used.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of results to be returned. Default value: 20. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// ECM region, such as `ap-xian-ecm`.
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`
}

type DescribeNetworkInterfacesRequest struct {
	*tchttp.BaseRequest
	
	// Queries the ID of the ENI instance, such as `eni-pxir56ns`. Each request supports a maximum of 100 instances. `NetworkInterfaceIds` and `Filters` cannot be specified at the same time.
	NetworkInterfaceIds []*string `json:"NetworkInterfaceIds,omitempty" name:"NetworkInterfaceIds"`

	// Filter. `NetworkInterfaceIds` and `Filters` cannot be specified at the same time.
	// vpc-id - String - (Filter) VPC instance ID, such as `vpc-f49l6u0z`.
	// subnet-id - String - (Filter) Subnet instance ID, such as `subnet-f49l6u0z`.
	// network-interface-id - String - (Filter) ENI instance ID, such as `eni-5k56k7k7`.
	// attachment.instance-id - String - (Filter) ID of the bound CVM instance, such as `ein-3nqpdn3i`.
	// groups.security-group-id - String - (Filter) ID of the bound security group instance, such as `sg-f9ekbxeq`.
	// network-interface-name - String - (Filter) ENI instance name.
	// network-interface-description - String - (Filter) ENI instance description.
	// address-ip - String - (Filter) Private IPv4 address.
	// tag-key - String - Required: no - (Filter) Filter by tag key. For directions, see Sample 2.
	// tag:tag-key - String - Required: no - (Filter) Filter by tag key-value pair. Replace `tag-key` with the specific tag key. For directions, see Sample 3.
	// is-primary - Boolean - Required: no - (Filter) Filter by whether it is a primary ENI. true: filter only by primary ENI; false: filter only by secondary ENI. If this parameter is not specified, filtering by both primary and secondary ENIs will be used.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of results to be returned. Default value: 20. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// ECM region, such as `ap-xian-ecm`.
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`
}

func (r *DescribeNetworkInterfacesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNetworkInterfacesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NetworkInterfaceIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "EcmRegion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNetworkInterfacesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNetworkInterfacesResponseParams struct {
	// Number of eligible instances.
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// List of instance details.
	// Note: this field may return null, indicating that no valid values can be obtained.
	NetworkInterfaceSet []*NetworkInterface `json:"NetworkInterfaceSet,omitempty" name:"NetworkInterfaceSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeNetworkInterfacesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNetworkInterfacesResponseParams `json:"Response"`
}

func (r *DescribeNetworkInterfacesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNetworkInterfacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNodeRequestParams struct {
	// Filter. InstanceFamily: instance family.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeNodeRequest struct {
	*tchttp.BaseRequest
	
	// Filter. InstanceFamily: instance family.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeNodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNodeResponseParams struct {
	// List of node details
	// Note: this field may return null, indicating that no valid values can be obtained.
	NodeSet []*Node `json:"NodeSet,omitempty" name:"NodeSet"`

	// Total number of nodes.
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeNodeResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNodeResponseParams `json:"Response"`
}

func (r *DescribeNodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePackingQuotaGroupRequestParams struct {
	// Filter. Zone: AZ; InstanceType: instance type; DataDiskSize: data disk size.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribePackingQuotaGroupRequest struct {
	*tchttp.BaseRequest
	
	// Filter. Zone: AZ; InstanceType: instance type; DataDiskSize: data disk size.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribePackingQuotaGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePackingQuotaGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePackingQuotaGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePackingQuotaGroupResponseParams struct {
	// Set of packing quotas.
	PackingQuotaSet []*PackingQuotaGroup `json:"PackingQuotaSet,omitempty" name:"PackingQuotaSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePackingQuotaGroupResponse struct {
	*tchttp.BaseResponse
	Response *DescribePackingQuotaGroupResponseParams `json:"Response"`
}

func (r *DescribePackingQuotaGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePackingQuotaGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePeakBaseOverviewRequestParams struct {
	// Start time (xxxx-xx-xx), such as `2019-08-14`. It is 7 days ago by default and should not be more than 90 days ago.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time (xxxx-xx-xx), such as `2019-08-14`. It is yesterday by default and should not be more than 90 days ago. When the time period between the start time and end time is within 30 days, data at the 1-hour granularity will be returned; otherwise, data at the 3-hour granularity will be returned.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type DescribePeakBaseOverviewRequest struct {
	*tchttp.BaseRequest
	
	// Start time (xxxx-xx-xx), such as `2019-08-14`. It is 7 days ago by default and should not be more than 90 days ago.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time (xxxx-xx-xx), such as `2019-08-14`. It is yesterday by default and should not be more than 90 days ago. When the time period between the start time and end time is within 30 days, data at the 1-hour granularity will be returned; otherwise, data at the 3-hour granularity will be returned.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribePeakBaseOverviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePeakBaseOverviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePeakBaseOverviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePeakBaseOverviewResponseParams struct {
	// List of basic peaks.
	// Note: this field may return null, indicating that no valid values can be obtained.
	PeakFamilyInfoSet []*PeakFamilyInfo `json:"PeakFamilyInfoSet,omitempty" name:"PeakFamilyInfoSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePeakBaseOverviewResponse struct {
	*tchttp.BaseResponse
	Response *DescribePeakBaseOverviewResponseParams `json:"Response"`
}

func (r *DescribePeakBaseOverviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePeakBaseOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePeakNetworkOverviewRequestParams struct {
	// Start time (xxxx-xx-xx), such as `2019-08-14`. It is 7 days ago by default and should not be more than 30 days ago.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time (xxxx-xx-xx), such as `2019-08-14`. It is yesterday by default and should not be more than 30 days ago. When the time period between the start time and end time is within 1 day, data at the 1-minute granularity will be returned; when the time period is within 7 days, data at the 5-minute granularity will be returned; otherwise, data at the 1-hour granularity will be returned.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Filter.
	// 
	// region    String      Required: no     (Filter) Filter by region. Fuzzy match is not supported. Note: you need to enter the ECM region to be queried before data can be returned.
	// area       String      Required: no     (Filter) Filter by region. Fuzzy match is not supported. Regions include `china-central`, `china-east`, etc. You can call `DescribeNode` to get the information of all regions. You can also use `ALL_REGION` to indicate all regions.
	// isp         String      Required: no     (Filter) Filter region traffic by ISP. ISPs include CTCC, CUCC, and CMCC. This parameter must be used together with `area`, and you can specify only one ISP at a time.
	// 
	// You can specify either `region` or `area`.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// Statistical period in seconds. Valid values: 60, 300.
	Period *int64 `json:"Period,omitempty" name:"Period"`
}

type DescribePeakNetworkOverviewRequest struct {
	*tchttp.BaseRequest
	
	// Start time (xxxx-xx-xx), such as `2019-08-14`. It is 7 days ago by default and should not be more than 30 days ago.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time (xxxx-xx-xx), such as `2019-08-14`. It is yesterday by default and should not be more than 30 days ago. When the time period between the start time and end time is within 1 day, data at the 1-minute granularity will be returned; when the time period is within 7 days, data at the 5-minute granularity will be returned; otherwise, data at the 1-hour granularity will be returned.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Filter.
	// 
	// region    String      Required: no     (Filter) Filter by region. Fuzzy match is not supported. Note: you need to enter the ECM region to be queried before data can be returned.
	// area       String      Required: no     (Filter) Filter by region. Fuzzy match is not supported. Regions include `china-central`, `china-east`, etc. You can call `DescribeNode` to get the information of all regions. You can also use `ALL_REGION` to indicate all regions.
	// isp         String      Required: no     (Filter) Filter region traffic by ISP. ISPs include CTCC, CUCC, and CMCC. This parameter must be used together with `area`, and you can specify only one ISP at a time.
	// 
	// You can specify either `region` or `area`.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// Statistical period in seconds. Valid values: 60, 300.
	Period *int64 `json:"Period,omitempty" name:"Period"`
}

func (r *DescribePeakNetworkOverviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePeakNetworkOverviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Filters")
	delete(f, "Period")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePeakNetworkOverviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePeakNetworkOverviewResponseParams struct {
	// Array of network peaks.
	// Note: this field may return null, indicating that no valid values can be obtained.
	PeakNetworkRegionSet []*PeakNetworkRegionInfo `json:"PeakNetworkRegionSet,omitempty" name:"PeakNetworkRegionSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePeakNetworkOverviewResponse struct {
	*tchttp.BaseResponse
	Response *DescribePeakNetworkOverviewResponseParams `json:"Response"`
}

func (r *DescribePeakNetworkOverviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePeakNetworkOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePriceRunInstanceRequestParams struct {
	// Instance model information
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// System disk information
	SystemDisk *SystemDisk `json:"SystemDisk,omitempty" name:"SystemDisk"`

	// Number of instances
	InstanceCount *uint64 `json:"InstanceCount,omitempty" name:"InstanceCount"`

	// Data disk information
	DataDisk []*DataDisk `json:"DataDisk,omitempty" name:"DataDisk"`

	// Instance billing type. Valid values:
	// `0`: Bill by daily resource usage peak (CPU, memory, and disk). It applies only to non-GNR models;
	// `1`: Bill by usage hours of an instance. It applies only to GNR models. It’s available to beta users now. To enable it, submit a ticket;
	// `2`: Bill by usage month of an instance. It applies only to GNR models.
	// If this field is left empty, `0` is selected by default for non-GNR models, and `2` is selected by default for GNR models.
	InstanceChargeType *int64 `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`
}

type DescribePriceRunInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Instance model information
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// System disk information
	SystemDisk *SystemDisk `json:"SystemDisk,omitempty" name:"SystemDisk"`

	// Number of instances
	InstanceCount *uint64 `json:"InstanceCount,omitempty" name:"InstanceCount"`

	// Data disk information
	DataDisk []*DataDisk `json:"DataDisk,omitempty" name:"DataDisk"`

	// Instance billing type. Valid values:
	// `0`: Bill by daily resource usage peak (CPU, memory, and disk). It applies only to non-GNR models;
	// `1`: Bill by usage hours of an instance. It applies only to GNR models. It’s available to beta users now. To enable it, submit a ticket;
	// `2`: Bill by usage month of an instance. It applies only to GNR models.
	// If this field is left empty, `0` is selected by default for non-GNR models, and `2` is selected by default for GNR models.
	InstanceChargeType *int64 `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`
}

func (r *DescribePriceRunInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePriceRunInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceType")
	delete(f, "SystemDisk")
	delete(f, "InstanceCount")
	delete(f, "DataDisk")
	delete(f, "InstanceChargeType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePriceRunInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePriceRunInstanceResponseParams struct {
	// Instance price information
	InstancePrice *InstancesPrice `json:"InstancePrice,omitempty" name:"InstancePrice"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePriceRunInstanceResponse struct {
	*tchttp.BaseResponse
	Response *DescribePriceRunInstanceResponseParams `json:"Response"`
}

func (r *DescribePriceRunInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePriceRunInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRouteConflictsRequestParams struct {
	// Route table instance ID, such as `rtb-azd4dt1c`.
	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`

	// List of conflicting destination ports to be checked.
	DestinationCidrBlocks []*string `json:"DestinationCidrBlocks,omitempty" name:"DestinationCidrBlocks"`
}

type DescribeRouteConflictsRequest struct {
	*tchttp.BaseRequest
	
	// Route table instance ID, such as `rtb-azd4dt1c`.
	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`

	// List of conflicting destination ports to be checked.
	DestinationCidrBlocks []*string `json:"DestinationCidrBlocks,omitempty" name:"DestinationCidrBlocks"`
}

func (r *DescribeRouteConflictsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRouteConflictsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RouteTableId")
	delete(f, "DestinationCidrBlocks")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRouteConflictsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRouteConflictsResponseParams struct {
	// List of routing policy conflicts.
	// Note: this field may return null, indicating that no valid values can be obtained.
	RouteConflictSet []*RouteConflict `json:"RouteConflictSet,omitempty" name:"RouteConflictSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeRouteConflictsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRouteConflictsResponseParams `json:"Response"`
}

func (r *DescribeRouteConflictsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRouteConflictsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRouteTablesRequestParams struct {
	// Route table instance ID, such as `rtb-azd4dt1c`.
	RouteTableIds []*string `json:"RouteTableIds,omitempty" name:"RouteTableIds"`

	// Filter. `RouteTableIds` and `Filters` cannot be specified at the same time.
	// route-table-id - String - (Filter) Route table instance ID.
	// route-table-name - String - (Filter) Route table name.
	// vpc-id - String - (Filter) VPC instance ID, such as `vpc-f49l6u0z`.
	// association.main - String - (Filter) Whether it is the main route table.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// Offset
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Limit
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// ECM region. If this parameter is left empty or not specified, it will indicate all regions.
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`
}

type DescribeRouteTablesRequest struct {
	*tchttp.BaseRequest
	
	// Route table instance ID, such as `rtb-azd4dt1c`.
	RouteTableIds []*string `json:"RouteTableIds,omitempty" name:"RouteTableIds"`

	// Filter. `RouteTableIds` and `Filters` cannot be specified at the same time.
	// route-table-id - String - (Filter) Route table instance ID.
	// route-table-name - String - (Filter) Route table name.
	// vpc-id - String - (Filter) VPC instance ID, such as `vpc-f49l6u0z`.
	// association.main - String - (Filter) Whether it is the main route table.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// Offset
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Limit
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// ECM region. If this parameter is left empty or not specified, it will indicate all regions.
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`
}

func (r *DescribeRouteTablesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRouteTablesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RouteTableIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "EcmRegion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRouteTablesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRouteTablesResponseParams struct {
	// Number of eligible instances
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// List of route tables
	// Note: this field may return null, indicating that no valid values can be obtained.
	RouteTableSet []*RouteTable `json:"RouteTableSet,omitempty" name:"RouteTableSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeRouteTablesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRouteTablesResponseParams `json:"Response"`
}

func (r *DescribeRouteTablesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRouteTablesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityGroupAssociationStatisticsRequestParams struct {
	// Security instance ID, such as `esg-33ocnj9n`, which can be obtained through [DescribeSecurityGroups](https://intl.cloud.tencent.com/document/product/1108/47697?from_cn_redirect=1).
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`
}

type DescribeSecurityGroupAssociationStatisticsRequest struct {
	*tchttp.BaseRequest
	
	// Security instance ID, such as `esg-33ocnj9n`, which can be obtained through [DescribeSecurityGroups](https://intl.cloud.tencent.com/document/product/1108/47697?from_cn_redirect=1).
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`
}

func (r *DescribeSecurityGroupAssociationStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityGroupAssociationStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecurityGroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecurityGroupAssociationStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityGroupAssociationStatisticsResponseParams struct {
	// Statistics on the instances associated with the security group.
	SecurityGroupAssociationStatisticsSet []*SecurityGroupAssociationStatistics `json:"SecurityGroupAssociationStatisticsSet,omitempty" name:"SecurityGroupAssociationStatisticsSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSecurityGroupAssociationStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSecurityGroupAssociationStatisticsResponseParams `json:"Response"`
}

func (r *DescribeSecurityGroupAssociationStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityGroupAssociationStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityGroupLimitsRequestParams struct {

}

type DescribeSecurityGroupLimitsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeSecurityGroupLimitsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityGroupLimitsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecurityGroupLimitsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityGroupLimitsResponseParams struct {
	// Security group quota limit.
	SecurityGroupLimitSet *SecurityGroupLimitSet `json:"SecurityGroupLimitSet,omitempty" name:"SecurityGroupLimitSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSecurityGroupLimitsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSecurityGroupLimitsResponseParams `json:"Response"`
}

func (r *DescribeSecurityGroupLimitsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityGroupLimitsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityGroupPoliciesRequestParams struct {
	// Security group instance ID, such as `esg-33ocnj9n`, which can be obtained through [DescribeSecurityGroups](https://intl.cloud.tencent.com/document/product/1108/47697?from_cn_redirect=1).
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
}

type DescribeSecurityGroupPoliciesRequest struct {
	*tchttp.BaseRequest
	
	// Security group instance ID, such as `esg-33ocnj9n`, which can be obtained through [DescribeSecurityGroups](https://intl.cloud.tencent.com/document/product/1108/47697?from_cn_redirect=1).
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
}

func (r *DescribeSecurityGroupPoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityGroupPoliciesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecurityGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecurityGroupPoliciesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityGroupPoliciesResponseParams struct {
	// Security group policy set.
	SecurityGroupPolicySet *SecurityGroupPolicySet `json:"SecurityGroupPolicySet,omitempty" name:"SecurityGroupPolicySet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSecurityGroupPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSecurityGroupPoliciesResponseParams `json:"Response"`
}

func (r *DescribeSecurityGroupPoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityGroupPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityGroupsRequestParams struct {
	// Security group instance ID, such as `esg-33ocnj9n`, which can be obtained through [DescribeSecurityGroups](https://intl.cloud.tencent.com/document/product/1108/47697?from_cn_redirect=1). Each request supports a maximum of 100 instances. `SecurityGroupIds` and `Filters` cannot be specified at the same time.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`

	// Filter. `SecurityGroupIds` and `Filters` cannot be specified at the same time.
	// security-group-id - String - (Filter) Security group ID.
	// security-group-name - String - (Filter) Security group name.
	// tag-key - String - Required: no - (Filter) Filter by tag key.
	// tag:tag-key - String - Required: no - (Filter) Filter by tag key-value pair. Replace `tag-key` with the specific tag key.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of results to be returned. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// Security group instance ID, such as `esg-33ocnj9n`, which can be obtained through [DescribeSecurityGroups](https://intl.cloud.tencent.com/document/product/1108/47697?from_cn_redirect=1). Each request supports a maximum of 100 instances. `SecurityGroupIds` and `Filters` cannot be specified at the same time.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`

	// Filter. `SecurityGroupIds` and `Filters` cannot be specified at the same time.
	// security-group-id - String - (Filter) Security group ID.
	// security-group-name - String - (Filter) Security group name.
	// tag-key - String - Required: no - (Filter) Filter by tag key.
	// tag:tag-key - String - Required: no - (Filter) Filter by tag key-value pair. Replace `tag-key` with the specific tag key.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of results to be returned. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecurityGroupIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecurityGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityGroupsResponseParams struct {
	// Number of eligible instances.
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Security group object.
	SecurityGroupSet []*SecurityGroup `json:"SecurityGroupSet,omitempty" name:"SecurityGroupSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSecurityGroupsResponseParams `json:"Response"`
}

func (r *DescribeSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSnapshotsRequestParams struct {
	// List of IDs of the snapshots to be queried. `SnapshotIds` and `Filters` cannot be specified at the same time.
	SnapshotIds []*string `json:"SnapshotIds,omitempty" name:"SnapshotIds"`

	// Filter. `SnapshotIds` and `Filters` cannot be specified at the same time.<br><li>snapshot-id - Array of String - Required: no - (Filter) Filter by snapshot ID, such as `snap-11112222`.<br><li>snapshot-name - Array of String - Required: no - (Filter) Filter by snapshot name.<br><li>snapshot-state - Array of String - Required: no - (Filter) Filter by snapshot status. NORMAL: normal; CREATING: creating; ROLLBACKING: rolling back.<br><li>disk-usage - Array of String - Required: no - (Filter) Filter by the type of the cloud disk from which a snapshot is created. SYSTEM_DISK: system disk; DATA_DISK: data disk.<br><li>project-id  - Array of String - Required: no - (Filter) Filter by the project ID of the cloud disk.<br><li>disk-id  - Array of String - Required: no - (Filter) Filter by the ID of the cloud disk from which a snapshot is created.<br><li>zone - Array of String - Required: no - (Filter) Filter by [AZ](https://intl.cloud.tencent.com/document/product/213/15753?from_cn_redirect=1#ZoneInfo).<br><li>encrypt - Array of String - Required: no - (Filter) Filter by whether a snapshot is created from an encrypted cloud disk. TRUE: yes; FALSE: no.
	// <li>snapshot-type- Array of String - Required: no - (Filter) Filter by the snapshot type specified in `snapshot-type`.
	// (SHARED_SNAPSHOT: shared snapshot | PRIVATE_SNAPSHOT: private snapshot.)
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// Number of returned results. Default value: 20. Maximum value: 100. For more information on `Limit`, see the relevant section of the API [Overview](https://intl.cloud.tencent.com/document/product/362/15633?from_cn_redirect=1).
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Field by which snapshots are sorted. Valid values:<br><li>CREATE_TIME: sort by snapshot creation time<br>Snapshots are sorted by creation time by default.
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`

	// Offset. Default value: 0. For more information on `Offset`, see the relevant section of the API [Overview](https://intl.cloud.tencent.com/document/product/362/15633?from_cn_redirect=1).
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Sorting order of cloud disks. Valid values:<br><li>ASC: ascending<br><li>DESC: descending.
	Order *string `json:"Order,omitempty" name:"Order"`
}

type DescribeSnapshotsRequest struct {
	*tchttp.BaseRequest
	
	// List of IDs of the snapshots to be queried. `SnapshotIds` and `Filters` cannot be specified at the same time.
	SnapshotIds []*string `json:"SnapshotIds,omitempty" name:"SnapshotIds"`

	// Filter. `SnapshotIds` and `Filters` cannot be specified at the same time.<br><li>snapshot-id - Array of String - Required: no - (Filter) Filter by snapshot ID, such as `snap-11112222`.<br><li>snapshot-name - Array of String - Required: no - (Filter) Filter by snapshot name.<br><li>snapshot-state - Array of String - Required: no - (Filter) Filter by snapshot status. NORMAL: normal; CREATING: creating; ROLLBACKING: rolling back.<br><li>disk-usage - Array of String - Required: no - (Filter) Filter by the type of the cloud disk from which a snapshot is created. SYSTEM_DISK: system disk; DATA_DISK: data disk.<br><li>project-id  - Array of String - Required: no - (Filter) Filter by the project ID of the cloud disk.<br><li>disk-id  - Array of String - Required: no - (Filter) Filter by the ID of the cloud disk from which a snapshot is created.<br><li>zone - Array of String - Required: no - (Filter) Filter by [AZ](https://intl.cloud.tencent.com/document/product/213/15753?from_cn_redirect=1#ZoneInfo).<br><li>encrypt - Array of String - Required: no - (Filter) Filter by whether a snapshot is created from an encrypted cloud disk. TRUE: yes; FALSE: no.
	// <li>snapshot-type- Array of String - Required: no - (Filter) Filter by the snapshot type specified in `snapshot-type`.
	// (SHARED_SNAPSHOT: shared snapshot | PRIVATE_SNAPSHOT: private snapshot.)
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// Number of returned results. Default value: 20. Maximum value: 100. For more information on `Limit`, see the relevant section of the API [Overview](https://intl.cloud.tencent.com/document/product/362/15633?from_cn_redirect=1).
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Field by which snapshots are sorted. Valid values:<br><li>CREATE_TIME: sort by snapshot creation time<br>Snapshots are sorted by creation time by default.
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`

	// Offset. Default value: 0. For more information on `Offset`, see the relevant section of the API [Overview](https://intl.cloud.tencent.com/document/product/362/15633?from_cn_redirect=1).
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Sorting order of cloud disks. Valid values:<br><li>ASC: ascending<br><li>DESC: descending.
	Order *string `json:"Order,omitempty" name:"Order"`
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
	delete(f, "Limit")
	delete(f, "OrderField")
	delete(f, "Offset")
	delete(f, "Order")
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
type DescribeSubnetsRequestParams struct {
	// Subnet instance ID, such as `subnet-pxir56ns`. Each request supports a maximum of 100 instances. `SubnetIds` and `Filters` cannot be specified at the same time.
	SubnetIds []*string `json:"SubnetIds,omitempty" name:"SubnetIds"`

	// Filter. `SubnetIds` and `Filters` cannot be specified at the same time.
	// subnet-id - String - Subnet instance name.
	// subnet-name - String - Subnet name. Only fuzzy query by a single value is supported.
	// cidr-block - String - Subnet IP address range, such as `192.168.1.0`. Only fuzzy query by a single value is supported.
	// vpc-id - String - VPC instance ID, such as `vpc-f49l6u0z`.
	// vpc-cidr-block  - String - VPC IP address range, such as `192.168.1.0`. Only fuzzy query by a single value is supported.
	// region - String - ECM region.
	// zone - String - AZ.
	// tag-key - String - Required: no - Filter by tag key.
	// tag:tag-key - String - Required: no - Filter by tag key-value pair.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// Offset
	Offset *string `json:"Offset,omitempty" name:"Offset"`

	// Number of returned results.
	Limit *string `json:"Limit,omitempty" name:"Limit"`

	// ECM region
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// Sorting method. time: sort in reverse chronological order; default: sort by network planning.
	Sort *string `json:"Sort,omitempty" name:"Sort"`
}

type DescribeSubnetsRequest struct {
	*tchttp.BaseRequest
	
	// Subnet instance ID, such as `subnet-pxir56ns`. Each request supports a maximum of 100 instances. `SubnetIds` and `Filters` cannot be specified at the same time.
	SubnetIds []*string `json:"SubnetIds,omitempty" name:"SubnetIds"`

	// Filter. `SubnetIds` and `Filters` cannot be specified at the same time.
	// subnet-id - String - Subnet instance name.
	// subnet-name - String - Subnet name. Only fuzzy query by a single value is supported.
	// cidr-block - String - Subnet IP address range, such as `192.168.1.0`. Only fuzzy query by a single value is supported.
	// vpc-id - String - VPC instance ID, such as `vpc-f49l6u0z`.
	// vpc-cidr-block  - String - VPC IP address range, such as `192.168.1.0`. Only fuzzy query by a single value is supported.
	// region - String - ECM region.
	// zone - String - AZ.
	// tag-key - String - Required: no - Filter by tag key.
	// tag:tag-key - String - Required: no - Filter by tag key-value pair.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// Offset
	Offset *string `json:"Offset,omitempty" name:"Offset"`

	// Number of returned results.
	Limit *string `json:"Limit,omitempty" name:"Limit"`

	// ECM region
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// Sorting method. time: sort in reverse chronological order; default: sort by network planning.
	Sort *string `json:"Sort,omitempty" name:"Sort"`
}

func (r *DescribeSubnetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSubnetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubnetIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "EcmRegion")
	delete(f, "Sort")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSubnetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSubnetsResponseParams struct {
	// Number of eligible instances.
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Subnet object.
	// Note: this field may return null, indicating that no valid values can be obtained.
	SubnetSet []*Subnet `json:"SubnetSet,omitempty" name:"SubnetSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSubnetsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSubnetsResponseParams `json:"Response"`
}

func (r *DescribeSubnetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSubnetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTargetHealthRequestParams struct {
	// List of IDs of CLB instances to be queried.
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitempty" name:"LoadBalancerIds"`
}

type DescribeTargetHealthRequest struct {
	*tchttp.BaseRequest
	
	// List of IDs of CLB instances to be queried.
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitempty" name:"LoadBalancerIds"`
}

func (r *DescribeTargetHealthRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTargetHealthRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTargetHealthRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTargetHealthResponseParams struct {
	// List of CLB instances.
	// Note: this field may return null, indicating that no valid values can be obtained.
	LoadBalancers []*LoadBalancerHealth `json:"LoadBalancers,omitempty" name:"LoadBalancers"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTargetHealthResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTargetHealthResponseParams `json:"Response"`
}

func (r *DescribeTargetHealthResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTargetHealthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTargetsRequestParams struct {
	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// List of listener IDs.
	ListenerIds []*string `json:"ListenerIds,omitempty" name:"ListenerIds"`

	// Listener protocol type.
	Protocol *int64 `json:"Protocol,omitempty" name:"Protocol"`

	// Listener port.
	Port *int64 `json:"Port,omitempty" name:"Port"`
}

type DescribeTargetsRequest struct {
	*tchttp.BaseRequest
	
	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// List of listener IDs.
	ListenerIds []*string `json:"ListenerIds,omitempty" name:"ListenerIds"`

	// Listener protocol type.
	Protocol *int64 `json:"Protocol,omitempty" name:"Protocol"`

	// Listener port.
	Port *int64 `json:"Port,omitempty" name:"Port"`
}

func (r *DescribeTargetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTargetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "ListenerIds")
	delete(f, "Protocol")
	delete(f, "Port")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTargetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTargetsResponseParams struct {
	// Information of real servers bound to the listener.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Listeners []*ListenerBackend `json:"Listeners,omitempty" name:"Listeners"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTargetsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTargetsResponseParams `json:"Response"`
}

func (r *DescribeTargetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTargetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskResultRequestParams struct {
	// ECM region
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// Async task ID.
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

type DescribeTaskResultRequest struct {
	*tchttp.BaseRequest
	
	// ECM region
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// Async task ID.
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeTaskResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EcmRegion")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTaskResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskResultResponseParams struct {
	// Async task ID.
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// Execution result. Valid values: SUCCESS; FAILED; RUNNING.
	Result *string `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTaskResultResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTaskResultResponseParams `json:"Response"`
}

func (r *DescribeTaskResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskStatusRequestParams struct {
	// Task description.
	TaskSet []*TaskInput `json:"TaskSet,omitempty" name:"TaskSet"`
}

type DescribeTaskStatusRequest struct {
	*tchttp.BaseRequest
	
	// Task description.
	TaskSet []*TaskInput `json:"TaskSet,omitempty" name:"TaskSet"`
}

func (r *DescribeTaskStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTaskStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskStatusResponseParams struct {
	// Task description.
	TaskSet []*TaskOutput `json:"TaskSet,omitempty" name:"TaskSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTaskStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTaskStatusResponseParams `json:"Response"`
}

func (r *DescribeTaskStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpcsRequestParams struct {
	// VPC instance ID, such as `vpc-f49l6u0z`. Each request supports a maximum of 100 instances. `VpcIds` and `Filters` cannot be specified at the same time.
	VpcIds []*string `json:"VpcIds,omitempty" name:"VpcIds"`

	// Filter. `VpcIds` and `Filters` cannot be specified at the same time.
	// vpc-name - String - VPC instance name. Only fuzzy query by a single value is supported.
	// vpc-id - String - VPC instance ID, such as `vpc-f49l6u0z`.
	// cidr-block - String - VPC CIDR. Only fuzzy query by a single value is supported.
	// region - String - VPC region.
	// tag-key - String - Required: no - Filter by tag key.
	// tag:tag-key - String - Required: no - Filter by tag key-value pair.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// Offset
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned results.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Region
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// Sorting method. time: sort in reverse chronological order; default: sort by network planning.
	Sort *string `json:"Sort,omitempty" name:"Sort"`
}

type DescribeVpcsRequest struct {
	*tchttp.BaseRequest
	
	// VPC instance ID, such as `vpc-f49l6u0z`. Each request supports a maximum of 100 instances. `VpcIds` and `Filters` cannot be specified at the same time.
	VpcIds []*string `json:"VpcIds,omitempty" name:"VpcIds"`

	// Filter. `VpcIds` and `Filters` cannot be specified at the same time.
	// vpc-name - String - VPC instance name. Only fuzzy query by a single value is supported.
	// vpc-id - String - VPC instance ID, such as `vpc-f49l6u0z`.
	// cidr-block - String - VPC CIDR. Only fuzzy query by a single value is supported.
	// region - String - VPC region.
	// tag-key - String - Required: no - Filter by tag key.
	// tag:tag-key - String - Required: no - Filter by tag key-value pair.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// Offset
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned results.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Region
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// Sorting method. time: sort in reverse chronological order; default: sort by network planning.
	Sort *string `json:"Sort,omitempty" name:"Sort"`
}

func (r *DescribeVpcsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpcsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "EcmRegion")
	delete(f, "Sort")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVpcsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpcsResponseParams struct {
	// Number of eligible objects.
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// VPC object.
	// Note: this field may return null, indicating that no valid values can be obtained.
	VpcSet []*VpcInfo `json:"VpcSet,omitempty" name:"VpcSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeVpcsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVpcsResponseParams `json:"Response"`
}

func (r *DescribeVpcsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpcsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetachNetworkInterfaceRequestParams struct {
	// ENI instance ID, such as `eni-m6dyj72l`.
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`

	// Instance ID, such as `ein-hcs7jkg4`.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// ECM region, such as `ap-xian-ecm`.
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`
}

type DetachNetworkInterfaceRequest struct {
	*tchttp.BaseRequest
	
	// ENI instance ID, such as `eni-m6dyj72l`.
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`

	// Instance ID, such as `ein-hcs7jkg4`.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// ECM region, such as `ap-xian-ecm`.
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`
}

func (r *DetachNetworkInterfaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetachNetworkInterfaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NetworkInterfaceId")
	delete(f, "InstanceId")
	delete(f, "EcmRegion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DetachNetworkInterfaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetachNetworkInterfaceResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DetachNetworkInterfaceResponse struct {
	*tchttp.BaseResponse
	Response *DetachNetworkInterfaceResponseParams `json:"Response"`
}

func (r *DetachNetworkInterfaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetachNetworkInterfaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableRoutesRequestParams struct {
	// Unique route table ID.
	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`

	// Routing policy ID.
	RouteIds []*uint64 `json:"RouteIds,omitempty" name:"RouteIds"`
}

type DisableRoutesRequest struct {
	*tchttp.BaseRequest
	
	// Unique route table ID.
	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`

	// Routing policy ID.
	RouteIds []*uint64 `json:"RouteIds,omitempty" name:"RouteIds"`
}

func (r *DisableRoutesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableRoutesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RouteTableId")
	delete(f, "RouteIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisableRoutesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableRoutesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DisableRoutesResponse struct {
	*tchttp.BaseResponse
	Response *DisableRoutesResponseParams `json:"Response"`
}

func (r *DisableRoutesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableRoutesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisassociateAddressRequestParams struct {
	// ECM region
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// Unique EIP ID, such as `eip-11112222`.
	AddressId *string `json:"AddressId,omitempty" name:"AddressId"`

	// Indicates whether to assign a general public IP after unbinding an EIP. Valid values:
	// TRUE: yes
	// FALSE: no.
	// Default value: FALSE.
	// 
	// You can specify this parameter only under the following conditions:
	// You can specify this parameter only when you unbind an EIP from the primary private IP of the primary ENI.
	// An account can reassign a general public IP after unbinding an EIP 10 times a day. More information can be obtained through the `DescribeAddressQuota` API.
	ReallocateNormalPublicIp *bool `json:"ReallocateNormalPublicIp,omitempty" name:"ReallocateNormalPublicIp"`
}

type DisassociateAddressRequest struct {
	*tchttp.BaseRequest
	
	// ECM region
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// Unique EIP ID, such as `eip-11112222`.
	AddressId *string `json:"AddressId,omitempty" name:"AddressId"`

	// Indicates whether to assign a general public IP after unbinding an EIP. Valid values:
	// TRUE: yes
	// FALSE: no.
	// Default value: FALSE.
	// 
	// You can specify this parameter only under the following conditions:
	// You can specify this parameter only when you unbind an EIP from the primary private IP of the primary ENI.
	// An account can reassign a general public IP after unbinding an EIP 10 times a day. More information can be obtained through the `DescribeAddressQuota` API.
	ReallocateNormalPublicIp *bool `json:"ReallocateNormalPublicIp,omitempty" name:"ReallocateNormalPublicIp"`
}

func (r *DisassociateAddressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisassociateAddressRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EcmRegion")
	delete(f, "AddressId")
	delete(f, "ReallocateNormalPublicIp")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisassociateAddressRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisassociateAddressResponseParams struct {
	// Async task ID (TaskId). You can use the `DescribeTaskResult` API to query the task status.
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DisassociateAddressResponse struct {
	*tchttp.BaseResponse
	Response *DisassociateAddressResponseParams `json:"Response"`
}

func (r *DisassociateAddressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisassociateAddressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisassociateInstancesKeyPairsRequestParams struct {
	// You can get available instance IDs in the following ways:
	// Log in to the console to query instance IDs.
	// Get the instance IDs from the `InstanceId` field in the information returned by the `DescribeInstances` API.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// List of key pair IDs. Each request can contain up to 100 key pairs. The key pair ID takes the form of `skey-11112222`.
	// 
	// You can get available key IDs in the following ways:
	// Log in to the console to query key IDs.
	// Get the key pair IDs from the `KeyId` field in the information returned by the `DescribeKeyPairs` API.
	KeyIds []*string `json:"KeyIds,omitempty" name:"KeyIds"`

	// Whether to force shut down the running instance. We recommend you manually shut down the running instance before unbinding the key. Valid values:
	// TRUE: yes.
	// FALSE: no.
	// 
	// Default value: FALSE.
	ForceStop *bool `json:"ForceStop,omitempty" name:"ForceStop"`
}

type DisassociateInstancesKeyPairsRequest struct {
	*tchttp.BaseRequest
	
	// You can get available instance IDs in the following ways:
	// Log in to the console to query instance IDs.
	// Get the instance IDs from the `InstanceId` field in the information returned by the `DescribeInstances` API.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// List of key pair IDs. Each request can contain up to 100 key pairs. The key pair ID takes the form of `skey-11112222`.
	// 
	// You can get available key IDs in the following ways:
	// Log in to the console to query key IDs.
	// Get the key pair IDs from the `KeyId` field in the information returned by the `DescribeKeyPairs` API.
	KeyIds []*string `json:"KeyIds,omitempty" name:"KeyIds"`

	// Whether to force shut down the running instance. We recommend you manually shut down the running instance before unbinding the key. Valid values:
	// TRUE: yes.
	// FALSE: no.
	// 
	// Default value: FALSE.
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
	// ID of the security group to be unbound, such as `esg-efil73jd`. You can unbind only one security group at a time.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`

	// ID of the instance to be unbound, such as `ein-lesecurk`. You can specify multiple instances.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

type DisassociateSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// ID of the security group to be unbound, such as `esg-efil73jd`. You can unbind only one security group at a time.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`

	// ID of the instance to be unbound, such as `ein-lesecurk`. You can specify multiple instances.
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

type DiskInfo struct {
	// Disk type: LOCAL_BASIC.
	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`

	// Disk ID
	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`

	// Disk size in GB
	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`
}

type EipQuota struct {
	// Quota name. Valid values:
	// TOTAL_EIP_QUOTA: quota of EIPs in the current region;
	// DAILY_EIP_APPLY: today's number of applications in the current region;
	// DAILY_PUBLIC_IP_ASSIGN: number of public IP reassignments in the current region.
	QuotaId *string `json:"QuotaId,omitempty" name:"QuotaId"`

	// Current quantity
	QuotaCurrent *uint64 `json:"QuotaCurrent,omitempty" name:"QuotaCurrent"`

	// Quota
	QuotaLimit *uint64 `json:"QuotaLimit,omitempty" name:"QuotaLimit"`
}

// Predefined struct for user
type EnableRoutesRequestParams struct {
	// Unique route table ID.
	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`

	// Routing policy ID.
	RouteIds []*uint64 `json:"RouteIds,omitempty" name:"RouteIds"`
}

type EnableRoutesRequest struct {
	*tchttp.BaseRequest
	
	// Unique route table ID.
	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`

	// Routing policy ID.
	RouteIds []*uint64 `json:"RouteIds,omitempty" name:"RouteIds"`
}

func (r *EnableRoutesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableRoutesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RouteTableId")
	delete(f, "RouteIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EnableRoutesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableRoutesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type EnableRoutesResponse struct {
	*tchttp.BaseResponse
	Response *EnableRoutesResponseParams `json:"Response"`
}

func (r *EnableRoutesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableRoutesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnhancedService struct {
	// Whether to enable CWP.
	SecurityService *RunSecurityServiceEnabled `json:"SecurityService,omitempty" name:"SecurityService"`

	// Whether to enable CM.
	MonitorService *RunMonitorServiceEnabled `json:"MonitorService,omitempty" name:"MonitorService"`

	// Whether to enable IP direct access. If this parameter is not specified, IP direct access will be enabled by default for Linux images and is currently not supported for Windows images.
	EIPDirectService *RunEIPDirectServiceEnabled `json:"EIPDirectService,omitempty" name:"EIPDirectService"`
}

type Filter struct {
	// One or more filter values.
	Values []*string `json:"Values,omitempty" name:"Values"`

	// Filter name.
	Name *string `json:"Name,omitempty" name:"Name"`
}

type HaVip struct {
	// Unique HAVIP ID.
	HaVipId *string `json:"HaVipId,omitempty" name:"HaVipId"`

	// HAVIP name.
	HaVipName *string `json:"HaVipName,omitempty" name:"HaVipName"`

	// Virtual IP address.
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// VPC ID of the HAVIP.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Subnet ID of the HAVIP.
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// ID of the ENI associated with the HAVIP.
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`

	// ID of the bound instance.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Bound EIP.
	AddressIp *string `json:"AddressIp,omitempty" name:"AddressIp"`

	// Status:
	// AVAILABLE: running.
	// UNBIND: unbound.
	State *string `json:"State,omitempty" name:"State"`

	// Creation time.
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// ID of businesses that use HAVIP.
	Business *string `json:"Business,omitempty" name:"Business"`
}

type HealthCheck struct {
	// Whether to enable health check. Valid values: 1: enable; 0: disable
	// Note: this field may return null, indicating that no valid values can be obtained.
	HealthSwitch *int64 `json:"HealthSwitch,omitempty" name:"HealthSwitch"`

	// Health check response timeout period in seconds. Value range: 2–60. Default value: 2. The value of this parameter should be smaller than the check interval.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TimeOut *int64 `json:"TimeOut,omitempty" name:"TimeOut"`

	// Health check interval in seconds. Value range: 5–300. Default value: 5.
	// Note: this field may return null, indicating that no valid values can be obtained.
	IntervalTime *int64 `json:"IntervalTime,omitempty" name:"IntervalTime"`

	// Health threshold. Value range: 2–10. Default value: 3, indicating that if a forward is found healthy three consecutive times, it will be considered normal.
	// Note: this field may return null, indicating that no valid values can be obtained.
	HealthNum *int64 `json:"HealthNum,omitempty" name:"HealthNum"`

	// Unhealthy threshold. Value range: 2–10. Default value: 3, indicating that if a forward is found unhealthy three consecutive times, it will be considered exceptional.
	// Note: this field may return null, indicating that no valid values can be obtained.
	UnHealthyNum *int64 `json:"UnHealthyNum,omitempty" name:"UnHealthyNum"`

	// Health check port (a custom check parameter), which is the port of the real server by default. Unless you want to specify a port, we recommend you leave it empty.
	// Note: this field may return null, indicating that no valid values can be obtained.
	CheckPort *int64 `json:"CheckPort,omitempty" name:"CheckPort"`

	// Health check protocol (a custom check parameter), which is required if the value of `CheckType` is `CUSTOM`. This parameter represents the input format of the health check. Valid values: HEX, TEXT. If the value is `HEX`, the characters of `SendContext` and `RecvContext` can only be selected from `0123456789ABCDEF`, and the length must be an even number.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ContextType *string `json:"ContextType,omitempty" name:"ContextType"`

	// Health check protocol (a custom check parameter), which is required if the value of `CheckType` is `CUSTOM`. This parameter represents the content of the request sent by the health check. It can contain up to 500 visible ASCII characters.
	// Note: this field may return null, indicating that no valid values can be obtained.
	SendContext *string `json:"SendContext,omitempty" name:"SendContext"`

	// Health check protocol (a custom check parameter), which is required if the value of `CheckType` is `CUSTOM`. This parameter represents the result returned by the health check. It can contain up to 500 visible ASCII characters.
	// Note: this field may return null, indicating that no valid values can be obtained.
	RecvContext *string `json:"RecvContext,omitempty" name:"RecvContext"`

	// Health check protocol (a custom check parameter). Valid values: TCP, CUSTOM (applicable only to UDP listeners. If custom health check is used, this parameter will be required).
	// Note: this field may return null, indicating that no valid values can be obtained.
	CheckType *string `json:"CheckType,omitempty" name:"CheckType"`
}

type ISP struct {
	// ISP ID
	ISPId *string `json:"ISPId,omitempty" name:"ISPId"`

	// ISP name
	ISPName *string `json:"ISPName,omitempty" name:"ISPName"`
}

type ISPCounter struct {
	// ISP name
	ProviderName *string `json:"ProviderName,omitempty" name:"ProviderName"`

	// Number of nodes
	ProviderNodeNum *int64 `json:"ProviderNodeNum,omitempty" name:"ProviderNodeNum"`

	// Number of instances
	ProvederInstanceNum *int64 `json:"ProvederInstanceNum,omitempty" name:"ProvederInstanceNum"`

	// Zone instance information structure array
	ZoneInstanceInfoSet []*ZoneInstanceInfo `json:"ZoneInstanceInfoSet,omitempty" name:"ZoneInstanceInfoSet"`
}

type Image struct {
	// Image ID
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// Image name
	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`

	// Image status
	ImageState *string `json:"ImageState,omitempty" name:"ImageState"`

	// Image type
	ImageType *string `json:"ImageType,omitempty" name:"ImageType"`

	// OS name
	ImageOsName *string `json:"ImageOsName,omitempty" name:"ImageOsName"`

	// Image description
	ImageDescription *string `json:"ImageDescription,omitempty" name:"ImageDescription"`

	// Image import time
	ImageCreateTime *string `json:"ImageCreateTime,omitempty" name:"ImageCreateTime"`

	// Number of bits of the OS
	Architecture *string `json:"Architecture,omitempty" name:"Architecture"`

	// OS type
	OsType *string `json:"OsType,omitempty" name:"OsType"`

	// OS version
	OsVersion *string `json:"OsVersion,omitempty" name:"OsVersion"`

	// OS platform
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// Image owner
	ImageOwner *int64 `json:"ImageOwner,omitempty" name:"ImageOwner"`

	// Image size in GB
	ImageSize *int64 `json:"ImageSize,omitempty" name:"ImageSize"`

	// Image source information
	SrcImage *SrcImage `json:"SrcImage,omitempty" name:"SrcImage"`

	// Image source type
	ImageSource *string `json:"ImageSource,omitempty" name:"ImageSource"`

	// ID of the task in intermediate or failed status
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// Whether cloud-init is supported
	IsSupportCloudInit *bool `json:"IsSupportCloudInit,omitempty" name:"IsSupportCloudInit"`
}

type ImageLimitConfig struct {
	// Supported maximum image size in GB, including custom image size for import and central cloud image size.
	MaxImageSize *int64 `json:"MaxImageSize,omitempty" name:"MaxImageSize"`
}

type ImageOsList struct {
	// Supported Windows OS
	// Note: this field may return null, indicating that no valid values can be obtained.
	Windows []*string `json:"Windows,omitempty" name:"Windows"`

	// Supported Linux OS
	// Note: this field may return null, indicating that no valid values can be obtained.
	Linux []*string `json:"Linux,omitempty" name:"Linux"`
}

type ImageTask struct {
	// Image import status. Valid values: PENDING, PROCESSING, SUCCESS, FAILED
	State *string `json:"State,omitempty" name:"State"`

	// Cause of import failure (FAILED)
	// Note: this field may return null, indicating that no valid values can be obtained.
	Message *string `json:"Message,omitempty" name:"Message"`

	// Image name
	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`

	// Creation time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

// Predefined struct for user
type ImportImageRequestParams struct {
	// Image ID.
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// Image description.
	ImageDescription *string `json:"ImageDescription,omitempty" name:"ImageDescription"`

	// Source region
	SourceRegion *string `json:"SourceRegion,omitempty" name:"SourceRegion"`
}

type ImportImageRequest struct {
	*tchttp.BaseRequest
	
	// Image ID.
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// Image description.
	ImageDescription *string `json:"ImageDescription,omitempty" name:"ImageDescription"`

	// Source region
	SourceRegion *string `json:"SourceRegion,omitempty" name:"SourceRegion"`
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
	delete(f, "ImageId")
	delete(f, "ImageDescription")
	delete(f, "SourceRegion")
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

type Instance struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Instance name, such as `ens-34241f3s`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// Instance status. Valid values:
	// PENDING: creating
	// LAUNCH_FAILED: failed to create
	// RUNNING: running
	// STOPPED: shut down
	// STARTING: starting
	// STOPPING: shutting down
	// REBOOTING: restarting
	// SHUTDOWN: to be terminated
	// TERMINATING: terminating.
	InstanceState *string `json:"InstanceState,omitempty" name:"InstanceState"`

	// Information of the image currently used by the instance.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Image *Image `json:"Image,omitempty" name:"Image"`

	// Basic information of the current module of the instance.
	// Note: this field may return null, indicating that no valid values can be obtained.
	SimpleModule *SimpleModule `json:"SimpleModule,omitempty" name:"SimpleModule"`

	// Location information of the instance.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Position *Position `json:"Position,omitempty" name:"Position"`

	// Network information of the instance.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Internet *Internet `json:"Internet,omitempty" name:"Internet"`

	// Configuration information of the instance.
	// Note: this field may return null, indicating that no valid values can be obtained.
	InstanceTypeConfig *InstanceTypeConfig `json:"InstanceTypeConfig,omitempty" name:"InstanceTypeConfig"`

	// Instance creation time.
	// Note: this field may return null, indicating that no valid values can be obtained.
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Instance tag information.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TagSet []*Tag `json:"TagSet,omitempty" name:"TagSet"`

	// Last operation on the instance.
	// Note: this field may return null, indicating that no valid values can be obtained.
	LatestOperation *string `json:"LatestOperation,omitempty" name:"LatestOperation"`

	// Result of the last operation on the instance.
	// Note: this field may return null, indicating that no valid values can be obtained.
	LatestOperationState *string `json:"LatestOperationState,omitempty" name:"LatestOperationState"`

	// Instance business status. Valid values:
	// NORMAL: normal
	// EXPIRED: expired
	// PROTECTIVELY_ISOLATED: isolated.
	// Note: this field may return null, indicating that no valid values can be obtained.
	RestrictState *string `json:"RestrictState,omitempty" name:"RestrictState"`

	// System disk size in GB.
	// Note: this field may return null, indicating that no valid values can be obtained.
	SystemDiskSize *int64 `json:"SystemDiskSize,omitempty" name:"SystemDiskSize"`

	// Data disk size in GB.
	// Note: this field may return null, indicating that no valid values can be obtained.
	DataDiskSize *int64 `json:"DataDiskSize,omitempty" name:"DataDiskSize"`

	// Instance UUID
	// Note: this field may return null, indicating that no valid values can be obtained.
	UUID *string `json:"UUID,omitempty" name:"UUID"`

	// Billing mode.
	//     0: postpaid.
	//     1: prepaid.
	// Note: this field may return null, indicating that no valid values can be obtained.
	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`

	// Expiration time in the format of `yyyy-mm-dd HH:mm:ss`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// Isolation time in the format of `yyyy-mm-dd HH:mm:ss`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	IsolatedTime *string `json:"IsolatedTime,omitempty" name:"IsolatedTime"`

	// Auto-Renewal flag.
	//       0: no.
	//       1: yes.
	// Note: this field may return null, indicating that no valid values can be obtained.
	RenewFlag *int64 `json:"RenewFlag,omitempty" name:"RenewFlag"`

	// Expiration status.
	//     NORMAL: normal.
	//     WILL_EXPIRE: about to expire.
	//     EXPIRED: expired.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ExpireState *string `json:"ExpireState,omitempty" name:"ExpireState"`

	// System disk information
	// Note: this field may return null, indicating that no valid values can be obtained.
	SystemDisk *DiskInfo `json:"SystemDisk,omitempty" name:"SystemDisk"`

	// Data disk information
	// Note: this field may return null, indicating that no valid values can be obtained.
	DataDisks []*DiskInfo `json:"DataDisks,omitempty" name:"DataDisks"`

	// New instance flag
	// Note: this field may return null, indicating that no valid values can be obtained.
	NewFlag *int64 `json:"NewFlag,omitempty" name:"NewFlag"`

	// Security group of the instance, which can be obtained from the `sgId` field in the returned value of the `DescribeSecurityGroups` API.
	// Note: this field may return null, indicating that no valid values can be obtained.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`

	// VPC attribute
	// Note: this field may return null, indicating that no valid values can be obtained.
	VirtualPrivateCloud *VirtualPrivateCloud `json:"VirtualPrivateCloud,omitempty" name:"VirtualPrivateCloud"`

	// ISP field of the instance.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ISP *string `json:"ISP,omitempty" name:"ISP"`

	// Physical location information. Note that this field is currently a reserved field and null.
	// Note: this field may return null, indicating that no valid values can be obtained.
	PhysicalPosition *PhysicalPosition `json:"PhysicalPosition,omitempty" name:"PhysicalPosition"`
}

type InstanceFamilyConfig struct {
	// Model name
	InstanceFamilyName *string `json:"InstanceFamilyName,omitempty" name:"InstanceFamilyName"`

	// Model ID
	InstanceFamily *string `json:"InstanceFamily,omitempty" name:"InstanceFamily"`
}

type InstanceFamilyTypeConfig struct {
	// Instance model family type ID
	InstanceFamilyType *string `json:"InstanceFamilyType,omitempty" name:"InstanceFamilyType"`

	// Instance model family type name
	InstanceFamilyTypeName *string `json:"InstanceFamilyTypeName,omitempty" name:"InstanceFamilyTypeName"`
}

type InstanceNetworkInfo struct {
	// Private and public IP information of the instance.
	// Note: this field may return null, indicating that no valid values can be obtained.
	AddressInfoSet []*AddressInfo `json:"AddressInfoSet,omitempty" name:"AddressInfoSet"`

	// ENI ID.
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`

	// ENI name.
	// Note: this field may return null, indicating that no valid values can be obtained.
	NetworkInterfaceName *string `json:"NetworkInterfaceName,omitempty" name:"NetworkInterfaceName"`

	// Primary ENI attribute. Valid values: true: primary ENI; false: secondary ENI.
	Primary *bool `json:"Primary,omitempty" name:"Primary"`
}

type InstanceNetworkLimitConfig struct {
	// Number of CPU cores
	CpuNum *int64 `json:"CpuNum,omitempty" name:"CpuNum"`

	// ENI quantity limit
	NetworkInterfaceLimit *int64 `json:"NetworkInterfaceLimit,omitempty" name:"NetworkInterfaceLimit"`

	// Private IP quantity limit per ENI
	InnerIpPerNetworkInterface *int64 `json:"InnerIpPerNetworkInterface,omitempty" name:"InnerIpPerNetworkInterface"`

	// Public IP limit per instance
	PublicIpPerInstance *int64 `json:"PublicIpPerInstance,omitempty" name:"PublicIpPerInstance"`
}

type InstanceOperator struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Prohibited operations for the instance
	// Note: this field may return null, indicating that no valid values can be obtained.
	DeniedActions []*OperatorAction `json:"DeniedActions,omitempty" name:"DeniedActions"`
}

type InstancePricesPartDetail struct {
	// CPU price information
	CpuPrice *PriceDetail `json:"CpuPrice,omitempty" name:"CpuPrice"`

	// Memory price information
	MemPrice *PriceDetail `json:"MemPrice,omitempty" name:"MemPrice"`

	// Disk price information
	DisksPrice *PriceDetail `json:"DisksPrice,omitempty" name:"DisksPrice"`
}

type InstanceStatistic struct {
	// Instance type
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// Number of instances
	InstanceCount *int64 `json:"InstanceCount,omitempty" name:"InstanceCount"`
}

type InstanceTypeConfig struct {
	// Model family configuration information
	InstanceFamilyConfig *InstanceFamilyConfig `json:"InstanceFamilyConfig,omitempty" name:"InstanceFamilyConfig"`

	// Model
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// Number of CPU cores
	Vcpu *int64 `json:"Vcpu,omitempty" name:"Vcpu"`

	// Memory size
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// Clock rate
	Frequency *string `json:"Frequency,omitempty" name:"Frequency"`

	// CPU model
	CpuModelName *string `json:"CpuModelName,omitempty" name:"CpuModelName"`

	// Instance family type configuration information
	InstanceFamilyTypeConfig *InstanceFamilyTypeConfig `json:"InstanceFamilyTypeConfig,omitempty" name:"InstanceFamilyTypeConfig"`

	// Extra model information, which is a JSON string in the format of `{"dataDiskSize":3200,"systemDiskSize":60, "systemDiskSizeShow":"default system disk size:60 GB","dataDiskSizeShow":"local NVMe SSD: 3200 GB"}`. It indicates a special model if it exists
	// Note: this field may return null, indicating that no valid values can be obtained.
	ExtInfo *string `json:"ExtInfo,omitempty" name:"ExtInfo"`

	// Number of GPU cards
	// Note: this field may return null, indicating that no valid values can be obtained.
	Vgpu *float64 `json:"Vgpu,omitempty" name:"Vgpu"`

	// GPU model
	// Note: this field may return null, indicating that no valid values can be obtained.
	GpuModelName *string `json:"GpuModelName,omitempty" name:"GpuModelName"`
}

type InstancesPrice struct {
	// Instance price details
	InstancePricesPartDetail *InstancePricesPartDetail `json:"InstancePricesPartDetail,omitempty" name:"InstancePricesPartDetail"`

	// Discount on the total instance price
	Discount *uint64 `json:"Discount,omitempty" name:"Discount"`

	// Discounted price
	DiscountPrice *uint64 `json:"DiscountPrice,omitempty" name:"DiscountPrice"`

	// Original cost
	OriginalPrice *uint64 `json:"OriginalPrice,omitempty" name:"OriginalPrice"`
}

type Internet struct {
	// Private network information list of the instance, with the primary ENI followed by secondary ENIs in the order of binding.
	// Note: this field may return null, indicating that no valid values can be obtained.
	PrivateIPAddressSet []*PrivateIPAddressInfo `json:"PrivateIPAddressSet,omitempty" name:"PrivateIPAddressSet"`

	// Public network information list of the instance, with the primary ENI followed by secondary ENIs in the order of binding.
	// Note: this field may return null, indicating that no valid values can be obtained.
	PublicIPAddressSet []*PublicIPAddressInfo `json:"PublicIPAddressSet,omitempty" name:"PublicIPAddressSet"`

	// Network information of the instance.
	// Note: this field may return null, indicating that no valid values can be obtained.
	InstanceNetworkInfoSet []*InstanceNetworkInfo `json:"InstanceNetworkInfoSet,omitempty" name:"InstanceNetworkInfoSet"`
}

type Ipv6Address struct {
	// IPv6 address, such as `3402:4e00:20:100:0:8cd9:2a67:71f3`
	Address *string `json:"Address,omitempty" name:"Address"`

	// Whether it is the primary IP.
	Primary *bool `json:"Primary,omitempty" name:"Primary"`

	// EIP instance ID, such as `eip-hxlqja90`.
	AddressId *string `json:"AddressId,omitempty" name:"AddressId"`

	// Description.
	Description *string `json:"Description,omitempty" name:"Description"`

	// Whether the public IP is blocked.
	IsWanIpBlocked *bool `json:"IsWanIpBlocked,omitempty" name:"IsWanIpBlocked"`

	// IPv6 address status:
	// PENDING: generating
	// MIGRATING: migrating
	// DELETING: deleting
	// AVAILABLE: available
	State *string `json:"State,omitempty" name:"State"`
}

type KeyPair struct {
	// Key pair ID, which is the unique identifier of a key pair.
	// Note: this field may return null, indicating that no valid values can be obtained.
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

	// Key pair name.
	// Note: this field may return null, indicating that no valid values can be obtained.
	KeyName *string `json:"KeyName,omitempty" name:"KeyName"`

	// Project ID of the key pair.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Key pair description.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Description *string `json:"Description,omitempty" name:"Description"`

	// Public key (in plain text) of key pair.
	// Note: this field may return null, indicating that no valid values can be obtained.
	PublicKey *string `json:"PublicKey,omitempty" name:"PublicKey"`

	// Private key (in plaintext) of a key pair. Tencent Cloud do not store private keys. Therefore, keep them secure.
	// Note: this field may return null, indicating that no valid values can be obtained.
	PrivateKey *string `json:"PrivateKey,omitempty" name:"PrivateKey"`

	// List of IDs of the instances associated with the key.
	// Note: this field may return null, indicating that no valid values can be obtained.
	AssociatedInstanceIds []*string `json:"AssociatedInstanceIds,omitempty" name:"AssociatedInstanceIds"`

	// Creation time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.
	// Note: this field may return null, indicating that no valid values can be obtained.
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
}

type Listener struct {
	// CLB listener ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// Listener protocol
	// Note: this field may return null, indicating that no valid values can be obtained.
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Listener port.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// Health check information of the listener
	// Note: this field may return null, indicating that no valid values can be obtained.
	HealthCheck *HealthCheck `json:"HealthCheck,omitempty" name:"HealthCheck"`

	// Request scheduling method
	// Note: this field may return null, indicating that no valid values can be obtained.
	Scheduler *string `json:"Scheduler,omitempty" name:"Scheduler"`

	// Session persistence time
	// Note: this field may return null, indicating that no valid values can be obtained.
	SessionExpireTime *int64 `json:"SessionExpireTime,omitempty" name:"SessionExpireTime"`

	// Listener name
	// Note: this field may return null, indicating that no valid values can be obtained.
	ListenerName *string `json:"ListenerName,omitempty" name:"ListenerName"`

	// Listener creation time
	// Note: this field may return null, indicating that no valid values can be obtained.
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Session type of the listener
	// Note: this field may return null, indicating that no valid values can be obtained.
	SessionType *string `json:"SessionType,omitempty" name:"SessionType"`

	// End port of the port range
	// Note: this field may return null, indicating that no valid values can be obtained.
	EndPort *int64 `json:"EndPort,omitempty" name:"EndPort"`
}

type ListenerBackend struct {
	// Listener ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// Listener protocol
	// Note: this field may return null, indicating that no valid values can be obtained.
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Listener port
	// Note: this field may return null, indicating that no valid values can be obtained.
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// List of real servers bound to the CLB instance
	// Note: this field may return null, indicating that no valid values can be obtained.
	Targets []*Backend `json:"Targets,omitempty" name:"Targets"`
}

type ListenerHealth struct {
	// Listener ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// Listener name
	// Note: this field may return null, indicating that no valid values can be obtained.
	ListenerName *string `json:"ListenerName,omitempty" name:"ListenerName"`

	// Listener protocol
	// Note: this field may return null, indicating that no valid values can be obtained.
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Listener port
	// Note: this field may return null, indicating that no valid values can be obtained.
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// List of forwarding rules of the listener
	// Note: this field may return null, indicating that no valid values can be obtained.
	Rules []*RuleHealth `json:"Rules,omitempty" name:"Rules"`
}

type LoadBalancer struct {
	// Region.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Region *string `json:"Region,omitempty" name:"Region"`

	// Location information.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Position *Position `json:"Position,omitempty" name:"Position"`

	// CLB instance ID.
	// Note: this field may return null, indicating that no valid values can be obtained.
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// CLB instance name.
	// Note: this field may return null, indicating that no valid values can be obtained.
	LoadBalancerName *string `json:"LoadBalancerName,omitempty" name:"LoadBalancerName"`

	// Network type of the CLB instance. Valid values: OPEN: public network
	// Note: this field may return null, indicating that no valid values can be obtained.
	LoadBalancerType *string `json:"LoadBalancerType,omitempty" name:"LoadBalancerType"`

	// List of VIPs of the CLB instance.
	// Note: this field may return null, indicating that no valid values can be obtained.
	LoadBalancerVips []*string `json:"LoadBalancerVips,omitempty" name:"LoadBalancerVips"`

	// CLB instance status. Valid values:
	//  0: creating; 1: running.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// CLB instance creation time.
	// Note: this field may return null, indicating that no valid values can be obtained.
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Last status change time of the CLB instance.
	// Note: this field may return null, indicating that no valid values can be obtained.
	StatusTime *string `json:"StatusTime,omitempty" name:"StatusTime"`

	// VPC ID.
	// Note: this field may return null, indicating that no valid values can be obtained.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// CLB instance tag information.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Tags []*TagInfo `json:"Tags,omitempty" name:"Tags"`

	// ISP of the CLB IP address.
	// Note: this field may return null, indicating that no valid values can be obtained.
	VipIsp *string `json:"VipIsp,omitempty" name:"VipIsp"`

	// Network attribute of the CLB instance.
	// Note: this field may return null, indicating that no valid values can be obtained.
	NetworkAttributes *LoadBalancerInternetAccessible `json:"NetworkAttributes,omitempty" name:"NetworkAttributes"`

	// Security group.
	// Note: this field may return null, indicating that no valid values can be obtained.
	SecureGroups []*string `json:"SecureGroups,omitempty" name:"SecureGroups"`

	// Whether the real server opens the traffic from ELB to the internet.
	// Note: this field may return null, indicating that no valid values can be obtained.
	LoadBalancerPassToTarget *bool `json:"LoadBalancerPassToTarget,omitempty" name:"LoadBalancerPassToTarget"`

	// IPv6 address of a CLB instance
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	AddressIPv6 *string `json:"AddressIPv6,omitempty" name:"AddressIPv6"`
}

type LoadBalancerHealth struct {
	// CLB instance ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// CLB instance name
	// Note: this field may return null, indicating that no valid values can be obtained.
	LoadBalancerName *string `json:"LoadBalancerName,omitempty" name:"LoadBalancerName"`

	// List of listeners
	// Note: this field may return null, indicating that no valid values can be obtained.
	Listeners []*ListenerHealth `json:"Listeners,omitempty" name:"Listeners"`
}

type LoadBalancerInternetAccessible struct {
	// Maximum outbound bandwidth in Mbps. Default value: 10.
	InternetMaxBandwidthOut *int64 `json:"InternetMaxBandwidthOut,omitempty" name:"InternetMaxBandwidthOut"`
}

// Predefined struct for user
type MigrateNetworkInterfaceRequestParams struct {
	// ECM region, such as `ap-xian-ecm`.
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// ENI instance ID, such as `eni-m6dyj72l`.
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`

	// ID of the ECM instance bound to the ENI, such as `ein-r8hr2upy`.
	SourceInstanceId *string `json:"SourceInstanceId,omitempty" name:"SourceInstanceId"`

	// ID of the destination ECM instance to be migrated.
	DestinationInstanceId *string `json:"DestinationInstanceId,omitempty" name:"DestinationInstanceId"`
}

type MigrateNetworkInterfaceRequest struct {
	*tchttp.BaseRequest
	
	// ECM region, such as `ap-xian-ecm`.
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// ENI instance ID, such as `eni-m6dyj72l`.
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`

	// ID of the ECM instance bound to the ENI, such as `ein-r8hr2upy`.
	SourceInstanceId *string `json:"SourceInstanceId,omitempty" name:"SourceInstanceId"`

	// ID of the destination ECM instance to be migrated.
	DestinationInstanceId *string `json:"DestinationInstanceId,omitempty" name:"DestinationInstanceId"`
}

func (r *MigrateNetworkInterfaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *MigrateNetworkInterfaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EcmRegion")
	delete(f, "NetworkInterfaceId")
	delete(f, "SourceInstanceId")
	delete(f, "DestinationInstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "MigrateNetworkInterfaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type MigrateNetworkInterfaceResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type MigrateNetworkInterfaceResponse struct {
	*tchttp.BaseResponse
	Response *MigrateNetworkInterfaceResponseParams `json:"Response"`
}

func (r *MigrateNetworkInterfaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *MigrateNetworkInterfaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type MigratePrivateIpAddressRequestParams struct {
	// ECM region, such as `ap-xian-ecm`.
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// ID of the ENI instance bound to the private IP, such as `eni-11112222`.
	SourceNetworkInterfaceId *string `json:"SourceNetworkInterfaceId,omitempty" name:"SourceNetworkInterfaceId"`

	// ID of the destination ENI instance to be migrated.
	DestinationNetworkInterfaceId *string `json:"DestinationNetworkInterfaceId,omitempty" name:"DestinationNetworkInterfaceId"`

	// Private IP address to be migrated, such as `10.0.0.6`.
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" name:"PrivateIpAddress"`
}

type MigratePrivateIpAddressRequest struct {
	*tchttp.BaseRequest
	
	// ECM region, such as `ap-xian-ecm`.
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// ID of the ENI instance bound to the private IP, such as `eni-11112222`.
	SourceNetworkInterfaceId *string `json:"SourceNetworkInterfaceId,omitempty" name:"SourceNetworkInterfaceId"`

	// ID of the destination ENI instance to be migrated.
	DestinationNetworkInterfaceId *string `json:"DestinationNetworkInterfaceId,omitempty" name:"DestinationNetworkInterfaceId"`

	// Private IP address to be migrated, such as `10.0.0.6`.
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" name:"PrivateIpAddress"`
}

func (r *MigratePrivateIpAddressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *MigratePrivateIpAddressRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EcmRegion")
	delete(f, "SourceNetworkInterfaceId")
	delete(f, "DestinationNetworkInterfaceId")
	delete(f, "PrivateIpAddress")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "MigratePrivateIpAddressRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type MigratePrivateIpAddressResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type MigratePrivateIpAddressResponse struct {
	*tchttp.BaseResponse
	Response *MigratePrivateIpAddressResponseParams `json:"Response"`
}

func (r *MigratePrivateIpAddressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *MigratePrivateIpAddressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAddressAttributeRequestParams struct {
	// ECM region
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// Unique EIP ID, such as `eip-11112222`.
	AddressId *string `json:"AddressId,omitempty" name:"AddressId"`

	// New EIP name, which can contain up to 20 characters.
	AddressName *string `json:"AddressName,omitempty" name:"AddressName"`

	// Whether "direct access" is enabled for the EIP. TRUE: yes. FALSE: no. Note that this parameter is available only to users who have activated the EIP direct access feature.
	EipDirectConnection *string `json:"EipDirectConnection,omitempty" name:"EipDirectConnection"`
}

type ModifyAddressAttributeRequest struct {
	*tchttp.BaseRequest
	
	// ECM region
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// Unique EIP ID, such as `eip-11112222`.
	AddressId *string `json:"AddressId,omitempty" name:"AddressId"`

	// New EIP name, which can contain up to 20 characters.
	AddressName *string `json:"AddressName,omitempty" name:"AddressName"`

	// Whether "direct access" is enabled for the EIP. TRUE: yes. FALSE: no. Note that this parameter is available only to users who have activated the EIP direct access feature.
	EipDirectConnection *string `json:"EipDirectConnection,omitempty" name:"EipDirectConnection"`
}

func (r *ModifyAddressAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAddressAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EcmRegion")
	delete(f, "AddressId")
	delete(f, "AddressName")
	delete(f, "EipDirectConnection")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAddressAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAddressAttributeResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyAddressAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAddressAttributeResponseParams `json:"Response"`
}

func (r *ModifyAddressAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAddressAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAddressesBandwidthRequestParams struct {
	// ECM region
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// Unique EIP ID, such as `eip-xxxxxxx`
	AddressIds []*string `json:"AddressIds,omitempty" name:"AddressIds"`

	// Target bandwidth value
	InternetMaxBandwidthOut *uint64 `json:"InternetMaxBandwidthOut,omitempty" name:"InternetMaxBandwidthOut"`
}

type ModifyAddressesBandwidthRequest struct {
	*tchttp.BaseRequest
	
	// ECM region
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// Unique EIP ID, such as `eip-xxxxxxx`
	AddressIds []*string `json:"AddressIds,omitempty" name:"AddressIds"`

	// Target bandwidth value
	InternetMaxBandwidthOut *uint64 `json:"InternetMaxBandwidthOut,omitempty" name:"InternetMaxBandwidthOut"`
}

func (r *ModifyAddressesBandwidthRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAddressesBandwidthRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EcmRegion")
	delete(f, "AddressIds")
	delete(f, "InternetMaxBandwidthOut")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAddressesBandwidthRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAddressesBandwidthResponseParams struct {
	// Async task ID.
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyAddressesBandwidthResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAddressesBandwidthResponseParams `json:"Response"`
}

func (r *ModifyAddressesBandwidthResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAddressesBandwidthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDefaultSubnetRequestParams struct {
	// ECM region
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// ECM AZ
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// VPC ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Subnet ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
}

type ModifyDefaultSubnetRequest struct {
	*tchttp.BaseRequest
	
	// ECM region
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// ECM AZ
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// VPC ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Subnet ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
}

func (r *ModifyDefaultSubnetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDefaultSubnetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EcmRegion")
	delete(f, "Zone")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDefaultSubnetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDefaultSubnetResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyDefaultSubnetResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDefaultSubnetResponseParams `json:"Response"`
}

func (r *ModifyDefaultSubnetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDefaultSubnetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyHaVipAttributeRequestParams struct {
	// Unique HAVIP ID, such as `havip-9o233uri`.
	HaVipId *string `json:"HaVipId,omitempty" name:"HaVipId"`

	// HAVIP name, which can be customized with up to 60 characters.
	HaVipName *string `json:"HaVipName,omitempty" name:"HaVipName"`
}

type ModifyHaVipAttributeRequest struct {
	*tchttp.BaseRequest
	
	// Unique HAVIP ID, such as `havip-9o233uri`.
	HaVipId *string `json:"HaVipId,omitempty" name:"HaVipId"`

	// HAVIP name, which can be customized with up to 60 characters.
	HaVipName *string `json:"HaVipName,omitempty" name:"HaVipName"`
}

func (r *ModifyHaVipAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyHaVipAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "HaVipId")
	delete(f, "HaVipName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyHaVipAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyHaVipAttributeResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyHaVipAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyHaVipAttributeResponseParams `json:"Response"`
}

func (r *ModifyHaVipAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyHaVipAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyImageAttributeRequestParams struct {
	// Image ID, such as `img-gvbnzy6f`
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// Image name, which must meet the following requirements:
	// It can contain up to 20 characters.
	// - The image name cannot be the same as existing image names.
	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`

	// Image description, which must meet the following requirements:
	// - It cannot exceed 60 characters.
	ImageDescription *string `json:"ImageDescription,omitempty" name:"ImageDescription"`
}

type ModifyImageAttributeRequest struct {
	*tchttp.BaseRequest
	
	// Image ID, such as `img-gvbnzy6f`
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// Image name, which must meet the following requirements:
	// It can contain up to 20 characters.
	// - The image name cannot be the same as existing image names.
	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`

	// Image description, which must meet the following requirements:
	// - It cannot exceed 60 characters.
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
type ModifyInstancesAttributeRequestParams struct {
	// List of IDs of the instances to be modified. You can request up to 100 instances at a time.
	InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet"`

	// Instance name displayed after successful modification, which can contain up to 60 characters. If this parameter is not specified, an empty value will be displayed.
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// List of IDs of the security groups of the instance. The instance will be associated with the specified security groups and will be disassociated from the original security groups. The maximum quantity is 5.
	SecurityGroups []*string `json:"SecurityGroups,omitempty" name:"SecurityGroups"`
}

type ModifyInstancesAttributeRequest struct {
	*tchttp.BaseRequest
	
	// List of IDs of the instances to be modified. You can request up to 100 instances at a time.
	InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet"`

	// Instance name displayed after successful modification, which can contain up to 60 characters. If this parameter is not specified, an empty value will be displayed.
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// List of IDs of the security groups of the instance. The instance will be associated with the specified security groups and will be disassociated from the original security groups. The maximum quantity is 5.
	SecurityGroups []*string `json:"SecurityGroups,omitempty" name:"SecurityGroups"`
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
	delete(f, "InstanceIdSet")
	delete(f, "InstanceName")
	delete(f, "SecurityGroups")
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
type ModifyIpv6AddressesAttributeRequestParams struct {
	// ECM region
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// ENI instance ID, such as `eni-m6dyj72l`.
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`

	// Information of the specified IPv6 addresses.
	Ipv6Addresses []*Ipv6Address `json:"Ipv6Addresses,omitempty" name:"Ipv6Addresses"`
}

type ModifyIpv6AddressesAttributeRequest struct {
	*tchttp.BaseRequest
	
	// ECM region
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// ENI instance ID, such as `eni-m6dyj72l`.
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`

	// Information of the specified IPv6 addresses.
	Ipv6Addresses []*Ipv6Address `json:"Ipv6Addresses,omitempty" name:"Ipv6Addresses"`
}

func (r *ModifyIpv6AddressesAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyIpv6AddressesAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EcmRegion")
	delete(f, "NetworkInterfaceId")
	delete(f, "Ipv6Addresses")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyIpv6AddressesAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyIpv6AddressesAttributeResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyIpv6AddressesAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyIpv6AddressesAttributeResponseParams `json:"Response"`
}

func (r *ModifyIpv6AddressesAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyIpv6AddressesAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyListenerRequestParams struct {
	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// CLB listener ID
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// New listener name
	ListenerName *string `json:"ListenerName,omitempty" name:"ListenerName"`

	// Session persistence time in seconds. Value range: 30–3600. The default value is 0, indicating that session persistence is not enabled. This parameter is applicable only to TCP/UDP listeners.
	SessionExpireTime *int64 `json:"SessionExpireTime,omitempty" name:"SessionExpireTime"`

	// Health check parameters
	HealthCheck *HealthCheck `json:"HealthCheck,omitempty" name:"HealthCheck"`

	// Forwarding method of the listener. Valid values: WRR, LEAST_CONN.
	// They represent weighted round robin and least connections, respectively. Default value: WRR.
	Scheduler *string `json:"Scheduler,omitempty" name:"Scheduler"`
}

type ModifyListenerRequest struct {
	*tchttp.BaseRequest
	
	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// CLB listener ID
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// New listener name
	ListenerName *string `json:"ListenerName,omitempty" name:"ListenerName"`

	// Session persistence time in seconds. Value range: 30–3600. The default value is 0, indicating that session persistence is not enabled. This parameter is applicable only to TCP/UDP listeners.
	SessionExpireTime *int64 `json:"SessionExpireTime,omitempty" name:"SessionExpireTime"`

	// Health check parameters
	HealthCheck *HealthCheck `json:"HealthCheck,omitempty" name:"HealthCheck"`

	// Forwarding method of the listener. Valid values: WRR, LEAST_CONN.
	// They represent weighted round robin and least connections, respectively. Default value: WRR.
	Scheduler *string `json:"Scheduler,omitempty" name:"Scheduler"`
}

func (r *ModifyListenerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyListenerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "ListenerId")
	delete(f, "ListenerName")
	delete(f, "SessionExpireTime")
	delete(f, "HealthCheck")
	delete(f, "Scheduler")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyListenerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyListenerResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyListenerResponse struct {
	*tchttp.BaseResponse
	Response *ModifyListenerResponseParams `json:"Response"`
}

func (r *ModifyListenerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyListenerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLoadBalancerAttributesRequestParams struct {
	// Unique CLB ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// CLB instance name
	LoadBalancerName *string `json:"LoadBalancerName,omitempty" name:"LoadBalancerName"`

	// Network billing and bandwidth parameters
	InternetChargeInfo *LoadBalancerInternetAccessible `json:"InternetChargeInfo,omitempty" name:"InternetChargeInfo"`

	// Whether to allow ELB traffic to the target group. `true`: allows ELB traffic to the target group and verifies security groups only on ELB; `false`: denies ELB traffic to the target group and verifies security groups on both ELB and backend instances.
	LoadBalancerPassToTarget *bool `json:"LoadBalancerPassToTarget,omitempty" name:"LoadBalancerPassToTarget"`
}

type ModifyLoadBalancerAttributesRequest struct {
	*tchttp.BaseRequest
	
	// Unique CLB ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// CLB instance name
	LoadBalancerName *string `json:"LoadBalancerName,omitempty" name:"LoadBalancerName"`

	// Network billing and bandwidth parameters
	InternetChargeInfo *LoadBalancerInternetAccessible `json:"InternetChargeInfo,omitempty" name:"InternetChargeInfo"`

	// Whether to allow ELB traffic to the target group. `true`: allows ELB traffic to the target group and verifies security groups only on ELB; `false`: denies ELB traffic to the target group and verifies security groups on both ELB and backend instances.
	LoadBalancerPassToTarget *bool `json:"LoadBalancerPassToTarget,omitempty" name:"LoadBalancerPassToTarget"`
}

func (r *ModifyLoadBalancerAttributesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLoadBalancerAttributesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "LoadBalancerName")
	delete(f, "InternetChargeInfo")
	delete(f, "LoadBalancerPassToTarget")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLoadBalancerAttributesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLoadBalancerAttributesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyLoadBalancerAttributesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyLoadBalancerAttributesResponseParams `json:"Response"`
}

func (r *ModifyLoadBalancerAttributesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLoadBalancerAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyModuleConfigRequestParams struct {
	// Module ID.
	ModuleId *string `json:"ModuleId,omitempty" name:"ModuleId"`

	// Model ID.
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// Default data disk size in GB. It cannot exceed the system disk size range. For more information, see `DescribeConfig`.
	DefaultDataDiskSize *int64 `json:"DefaultDataDiskSize,omitempty" name:"DefaultDataDiskSize"`

	// Default system disk size in GB. It cannot exceed the system disk size range. For more information, see `DescribeConfig`.
	DefaultSystemDiskSize *int64 `json:"DefaultSystemDiskSize,omitempty" name:"DefaultSystemDiskSize"`

	// System disk
	SystemDisk *SystemDisk `json:"SystemDisk,omitempty" name:"SystemDisk"`

	// Data disk
	DataDisks []*DataDisk `json:"DataDisks,omitempty" name:"DataDisks"`
}

type ModifyModuleConfigRequest struct {
	*tchttp.BaseRequest
	
	// Module ID.
	ModuleId *string `json:"ModuleId,omitempty" name:"ModuleId"`

	// Model ID.
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// Default data disk size in GB. It cannot exceed the system disk size range. For more information, see `DescribeConfig`.
	DefaultDataDiskSize *int64 `json:"DefaultDataDiskSize,omitempty" name:"DefaultDataDiskSize"`

	// Default system disk size in GB. It cannot exceed the system disk size range. For more information, see `DescribeConfig`.
	DefaultSystemDiskSize *int64 `json:"DefaultSystemDiskSize,omitempty" name:"DefaultSystemDiskSize"`

	// System disk
	SystemDisk *SystemDisk `json:"SystemDisk,omitempty" name:"SystemDisk"`

	// Data disk
	DataDisks []*DataDisk `json:"DataDisks,omitempty" name:"DataDisks"`
}

func (r *ModifyModuleConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyModuleConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModuleId")
	delete(f, "InstanceType")
	delete(f, "DefaultDataDiskSize")
	delete(f, "DefaultSystemDiskSize")
	delete(f, "SystemDisk")
	delete(f, "DataDisks")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyModuleConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyModuleConfigResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyModuleConfigResponse struct {
	*tchttp.BaseResponse
	Response *ModifyModuleConfigResponseParams `json:"Response"`
}

func (r *ModifyModuleConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyModuleConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyModuleDisableWanIpRequestParams struct {
	// Module ID
	ModuleId *string `json:"ModuleId,omitempty" name:"ModuleId"`

	// Whether to prohibit public IP assignment. Valid values: true: no; false: yes.
	DisableWanIp *bool `json:"DisableWanIp,omitempty" name:"DisableWanIp"`
}

type ModifyModuleDisableWanIpRequest struct {
	*tchttp.BaseRequest
	
	// Module ID
	ModuleId *string `json:"ModuleId,omitempty" name:"ModuleId"`

	// Whether to prohibit public IP assignment. Valid values: true: no; false: yes.
	DisableWanIp *bool `json:"DisableWanIp,omitempty" name:"DisableWanIp"`
}

func (r *ModifyModuleDisableWanIpRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyModuleDisableWanIpRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModuleId")
	delete(f, "DisableWanIp")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyModuleDisableWanIpRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyModuleDisableWanIpResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyModuleDisableWanIpResponse struct {
	*tchttp.BaseResponse
	Response *ModifyModuleDisableWanIpResponseParams `json:"Response"`
}

func (r *ModifyModuleDisableWanIpResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyModuleDisableWanIpResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyModuleImageRequestParams struct {
	// Default image ID
	DefaultImageId *string `json:"DefaultImageId,omitempty" name:"DefaultImageId"`

	// Module ID
	ModuleId *string `json:"ModuleId,omitempty" name:"ModuleId"`
}

type ModifyModuleImageRequest struct {
	*tchttp.BaseRequest
	
	// Default image ID
	DefaultImageId *string `json:"DefaultImageId,omitempty" name:"DefaultImageId"`

	// Module ID
	ModuleId *string `json:"ModuleId,omitempty" name:"ModuleId"`
}

func (r *ModifyModuleImageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyModuleImageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DefaultImageId")
	delete(f, "ModuleId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyModuleImageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyModuleImageResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyModuleImageResponse struct {
	*tchttp.BaseResponse
	Response *ModifyModuleImageResponseParams `json:"Response"`
}

func (r *ModifyModuleImageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyModuleImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyModuleIpDirectRequestParams struct {
	// Module ID.
	ModuleId *string `json:"ModuleId,omitempty" name:"ModuleId"`

	// Whether to disable IP direct access. Valid values:
	// true: yes
	// false: no
	CloseIpDirect *bool `json:"CloseIpDirect,omitempty" name:"CloseIpDirect"`
}

type ModifyModuleIpDirectRequest struct {
	*tchttp.BaseRequest
	
	// Module ID.
	ModuleId *string `json:"ModuleId,omitempty" name:"ModuleId"`

	// Whether to disable IP direct access. Valid values:
	// true: yes
	// false: no
	CloseIpDirect *bool `json:"CloseIpDirect,omitempty" name:"CloseIpDirect"`
}

func (r *ModifyModuleIpDirectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyModuleIpDirectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModuleId")
	delete(f, "CloseIpDirect")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyModuleIpDirectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyModuleIpDirectResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyModuleIpDirectResponse struct {
	*tchttp.BaseResponse
	Response *ModifyModuleIpDirectResponseParams `json:"Response"`
}

func (r *ModifyModuleIpDirectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyModuleIpDirectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyModuleNameRequestParams struct {
	// Module ID.
	ModuleId *string `json:"ModuleId,omitempty" name:"ModuleId"`

	// Module name.
	ModuleName *string `json:"ModuleName,omitempty" name:"ModuleName"`
}

type ModifyModuleNameRequest struct {
	*tchttp.BaseRequest
	
	// Module ID.
	ModuleId *string `json:"ModuleId,omitempty" name:"ModuleId"`

	// Module name.
	ModuleName *string `json:"ModuleName,omitempty" name:"ModuleName"`
}

func (r *ModifyModuleNameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyModuleNameRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModuleId")
	delete(f, "ModuleName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyModuleNameRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyModuleNameResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyModuleNameResponse struct {
	*tchttp.BaseResponse
	Response *ModifyModuleNameResponseParams `json:"Response"`
}

func (r *ModifyModuleNameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyModuleNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyModuleNetworkRequestParams struct {
	// Module ID
	ModuleId *string `json:"ModuleId,omitempty" name:"ModuleId"`

	// Default outbound bandwidth cap
	DefaultBandwidth *int64 `json:"DefaultBandwidth,omitempty" name:"DefaultBandwidth"`

	// Default inbound bandwidth cap
	DefaultBandwidthIn *int64 `json:"DefaultBandwidthIn,omitempty" name:"DefaultBandwidthIn"`
}

type ModifyModuleNetworkRequest struct {
	*tchttp.BaseRequest
	
	// Module ID
	ModuleId *string `json:"ModuleId,omitempty" name:"ModuleId"`

	// Default outbound bandwidth cap
	DefaultBandwidth *int64 `json:"DefaultBandwidth,omitempty" name:"DefaultBandwidth"`

	// Default inbound bandwidth cap
	DefaultBandwidthIn *int64 `json:"DefaultBandwidthIn,omitempty" name:"DefaultBandwidthIn"`
}

func (r *ModifyModuleNetworkRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyModuleNetworkRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModuleId")
	delete(f, "DefaultBandwidth")
	delete(f, "DefaultBandwidthIn")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyModuleNetworkRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyModuleNetworkResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyModuleNetworkResponse struct {
	*tchttp.BaseResponse
	Response *ModifyModuleNetworkResponseParams `json:"Response"`
}

func (r *ModifyModuleNetworkResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyModuleNetworkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyModuleSecurityGroupsRequestParams struct {
	// List of up to 5 security groups.
	SecurityGroupIdSet []*string `json:"SecurityGroupIdSet,omitempty" name:"SecurityGroupIdSet"`

	// Module ID.
	ModuleId *string `json:"ModuleId,omitempty" name:"ModuleId"`
}

type ModifyModuleSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// List of up to 5 security groups.
	SecurityGroupIdSet []*string `json:"SecurityGroupIdSet,omitempty" name:"SecurityGroupIdSet"`

	// Module ID.
	ModuleId *string `json:"ModuleId,omitempty" name:"ModuleId"`
}

func (r *ModifyModuleSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyModuleSecurityGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecurityGroupIdSet")
	delete(f, "ModuleId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyModuleSecurityGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyModuleSecurityGroupsResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyModuleSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *ModifyModuleSecurityGroupsResponseParams `json:"Response"`
}

func (r *ModifyModuleSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyModuleSecurityGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPrivateIpAddressesAttributeRequestParams struct {
	// ENI instance ID, such as `eni-m6dyj72l`.
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`

	// Information of the specified private IP addresses.
	PrivateIpAddresses []*PrivateIpAddressSpecification `json:"PrivateIpAddresses,omitempty" name:"PrivateIpAddresses"`

	// Region information of the ECM node, such as `ap-xian-ecm`.
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`
}

type ModifyPrivateIpAddressesAttributeRequest struct {
	*tchttp.BaseRequest
	
	// ENI instance ID, such as `eni-m6dyj72l`.
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`

	// Information of the specified private IP addresses.
	PrivateIpAddresses []*PrivateIpAddressSpecification `json:"PrivateIpAddresses,omitempty" name:"PrivateIpAddresses"`

	// Region information of the ECM node, such as `ap-xian-ecm`.
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`
}

func (r *ModifyPrivateIpAddressesAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPrivateIpAddressesAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NetworkInterfaceId")
	delete(f, "PrivateIpAddresses")
	delete(f, "EcmRegion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyPrivateIpAddressesAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPrivateIpAddressesAttributeResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyPrivateIpAddressesAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyPrivateIpAddressesAttributeResponseParams `json:"Response"`
}

func (r *ModifyPrivateIpAddressesAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPrivateIpAddressesAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRouteTableAttributeRequestParams struct {
	// Route table instance ID, such as `rtb-azd4dt1c`
	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`

	// Route table name
	RouteTableName *string `json:"RouteTableName,omitempty" name:"RouteTableName"`
}

type ModifyRouteTableAttributeRequest struct {
	*tchttp.BaseRequest
	
	// Route table instance ID, such as `rtb-azd4dt1c`
	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`

	// Route table name
	RouteTableName *string `json:"RouteTableName,omitempty" name:"RouteTableName"`
}

func (r *ModifyRouteTableAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRouteTableAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RouteTableId")
	delete(f, "RouteTableName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRouteTableAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRouteTableAttributeResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyRouteTableAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyRouteTableAttributeResponseParams `json:"Response"`
}

func (r *ModifyRouteTableAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRouteTableAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySecurityGroupAttributeRequestParams struct {
	// Security group instance ID, such as `esg-33ocnj9n`, which can be obtained through the `DescribeSecurityGroups` API.
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`

	// Security group name, which can be customized with up to 60 characters.
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// Security group remarks, which can contain up to 100 characters.
	GroupDescription *string `json:"GroupDescription,omitempty" name:"GroupDescription"`
}

type ModifySecurityGroupAttributeRequest struct {
	*tchttp.BaseRequest
	
	// Security group instance ID, such as `esg-33ocnj9n`, which can be obtained through the `DescribeSecurityGroups` API.
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`

	// Security group name, which can be customized with up to 60 characters.
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// Security group remarks, which can contain up to 100 characters.
	GroupDescription *string `json:"GroupDescription,omitempty" name:"GroupDescription"`
}

func (r *ModifySecurityGroupAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySecurityGroupAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecurityGroupId")
	delete(f, "GroupName")
	delete(f, "GroupDescription")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySecurityGroupAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySecurityGroupAttributeResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifySecurityGroupAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifySecurityGroupAttributeResponseParams `json:"Response"`
}

func (r *ModifySecurityGroupAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySecurityGroupAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySecurityGroupPoliciesRequestParams struct {
	// Security group instance ID, such as `esg-33ocnj9n`, which can be obtained through the `DescribeSecurityGroups` API.
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`

	// Security group policy set. You must specify both new egress and ingress policies for the `SecurityGroupPolicySet` object. You cannot customize the index (PolicyIndex) of the `SecurityGroupPolicy` object.
	SecurityGroupPolicySet *SecurityGroupPolicySet `json:"SecurityGroupPolicySet,omitempty" name:"SecurityGroupPolicySet"`

	// Whether security group sorting is supported. `True` indicates yes. If `SortPolicys` doesn't exist or is set to `False`, the security group policy can be modified.
	SortPolicys *bool `json:"SortPolicys,omitempty" name:"SortPolicys"`
}

type ModifySecurityGroupPoliciesRequest struct {
	*tchttp.BaseRequest
	
	// Security group instance ID, such as `esg-33ocnj9n`, which can be obtained through the `DescribeSecurityGroups` API.
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`

	// Security group policy set. You must specify both new egress and ingress policies for the `SecurityGroupPolicySet` object. You cannot customize the index (PolicyIndex) of the `SecurityGroupPolicy` object.
	SecurityGroupPolicySet *SecurityGroupPolicySet `json:"SecurityGroupPolicySet,omitempty" name:"SecurityGroupPolicySet"`

	// Whether security group sorting is supported. `True` indicates yes. If `SortPolicys` doesn't exist or is set to `False`, the security group policy can be modified.
	SortPolicys *bool `json:"SortPolicys,omitempty" name:"SortPolicys"`
}

func (r *ModifySecurityGroupPoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySecurityGroupPoliciesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecurityGroupId")
	delete(f, "SecurityGroupPolicySet")
	delete(f, "SortPolicys")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySecurityGroupPoliciesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySecurityGroupPoliciesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifySecurityGroupPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *ModifySecurityGroupPoliciesResponseParams `json:"Response"`
}

func (r *ModifySecurityGroupPoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySecurityGroupPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySubnetAttributeRequestParams struct {
	// Subnet instance ID, such as `subnet-pxir56ns`.
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// ECM region
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// Subnet name, which can contain up to 60 bytes.
	SubnetName *string `json:"SubnetName,omitempty" name:"SubnetName"`

	// Whether to enable broadcast for the subnet.
	EnableBroadcast *string `json:"EnableBroadcast,omitempty" name:"EnableBroadcast"`

	// Tag key value of the subnet
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

type ModifySubnetAttributeRequest struct {
	*tchttp.BaseRequest
	
	// Subnet instance ID, such as `subnet-pxir56ns`.
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// ECM region
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// Subnet name, which can contain up to 60 bytes.
	SubnetName *string `json:"SubnetName,omitempty" name:"SubnetName"`

	// Whether to enable broadcast for the subnet.
	EnableBroadcast *string `json:"EnableBroadcast,omitempty" name:"EnableBroadcast"`

	// Tag key value of the subnet
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

func (r *ModifySubnetAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySubnetAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubnetId")
	delete(f, "EcmRegion")
	delete(f, "SubnetName")
	delete(f, "EnableBroadcast")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySubnetAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySubnetAttributeResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifySubnetAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifySubnetAttributeResponseParams `json:"Response"`
}

func (r *ModifySubnetAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySubnetAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTargetPortRequestParams struct {
	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// CLB listener ID
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// List of real servers for which to modify the ports
	Targets []*Target `json:"Targets,omitempty" name:"Targets"`

	// New port of the real server bound to the listener or forwarding rule
	NewPort *int64 `json:"NewPort,omitempty" name:"NewPort"`
}

type ModifyTargetPortRequest struct {
	*tchttp.BaseRequest
	
	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// CLB listener ID
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// List of real servers for which to modify the ports
	Targets []*Target `json:"Targets,omitempty" name:"Targets"`

	// New port of the real server bound to the listener or forwarding rule
	NewPort *int64 `json:"NewPort,omitempty" name:"NewPort"`
}

func (r *ModifyTargetPortRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTargetPortRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "ListenerId")
	delete(f, "Targets")
	delete(f, "NewPort")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTargetPortRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTargetPortResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyTargetPortResponse struct {
	*tchttp.BaseResponse
	Response *ModifyTargetPortResponseParams `json:"Response"`
}

func (r *ModifyTargetPortResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTargetPortResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTargetWeightRequestParams struct {
	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// CLB listener ID
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// List of real servers for which to modify the weights
	Targets []*Target `json:"Targets,omitempty" name:"Targets"`

	// New forwarding weight of the real server. Value range: 0-100. Default value: 10. This parameter will not take effect if the `Targets.Weight` parameter is set.
	Weight *int64 `json:"Weight,omitempty" name:"Weight"`
}

type ModifyTargetWeightRequest struct {
	*tchttp.BaseRequest
	
	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// CLB listener ID
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// List of real servers for which to modify the weights
	Targets []*Target `json:"Targets,omitempty" name:"Targets"`

	// New forwarding weight of the real server. Value range: 0-100. Default value: 10. This parameter will not take effect if the `Targets.Weight` parameter is set.
	Weight *int64 `json:"Weight,omitempty" name:"Weight"`
}

func (r *ModifyTargetWeightRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTargetWeightRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "ListenerId")
	delete(f, "Targets")
	delete(f, "Weight")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTargetWeightRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTargetWeightResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyTargetWeightResponse struct {
	*tchttp.BaseResponse
	Response *ModifyTargetWeightResponseParams `json:"Response"`
}

func (r *ModifyTargetWeightResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTargetWeightResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVpcAttributeRequestParams struct {
	// VPC instance ID, such as `vpc-f49l6u0z`.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// ECM region
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// VPC name, which can be customized with up to 60 characters.
	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`

	// Tags
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// VPC description
	Description *string `json:"Description,omitempty" name:"Description"`
}

type ModifyVpcAttributeRequest struct {
	*tchttp.BaseRequest
	
	// VPC instance ID, such as `vpc-f49l6u0z`.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// ECM region
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// VPC name, which can be customized with up to 60 characters.
	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`

	// Tags
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// VPC description
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *ModifyVpcAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVpcAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcId")
	delete(f, "EcmRegion")
	delete(f, "VpcName")
	delete(f, "Tags")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyVpcAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVpcAttributeResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyVpcAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyVpcAttributeResponseParams `json:"Response"`
}

func (r *ModifyVpcAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVpcAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Module struct {
	// Module ID.
	ModuleId *string `json:"ModuleId,omitempty" name:"ModuleId"`

	// Module name.
	ModuleName *string `json:"ModuleName,omitempty" name:"ModuleName"`

	// Module status. Valid values:
	// NORMAL: normal.
	// DELETING: deleting 
	// DELETEFAILED: failed to delete.
	ModuleState *string `json:"ModuleState,omitempty" name:"ModuleState"`

	// Default system disk size.
	DefaultSystemDiskSize *int64 `json:"DefaultSystemDiskSize,omitempty" name:"DefaultSystemDiskSize"`

	// Default data disk size.
	DefaultDataDiskSize *int64 `json:"DefaultDataDiskSize,omitempty" name:"DefaultDataDiskSize"`

	// Default model.
	InstanceTypeConfig *InstanceTypeConfig `json:"InstanceTypeConfig,omitempty" name:"InstanceTypeConfig"`

	// Default image.
	DefaultImage *Image `json:"DefaultImage,omitempty" name:"DefaultImage"`

	// Creation time.
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Default outbound bandwidth.
	DefaultBandwidth *int64 `json:"DefaultBandwidth,omitempty" name:"DefaultBandwidth"`

	// Tag set.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TagSet []*Tag `json:"TagSet,omitempty" name:"TagSet"`

	// Whether to disable IP direct access.
	CloseIpDirect *int64 `json:"CloseIpDirect,omitempty" name:"CloseIpDirect"`

	// List of default security group IDs.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`

	// Default inbound bandwidth.
	DefaultBandwidthIn *int64 `json:"DefaultBandwidthIn,omitempty" name:"DefaultBandwidthIn"`

	// Custom script data
	UserData *string `json:"UserData,omitempty" name:"UserData"`

	// System disk information.
	SystemDisk *SystemDisk `json:"SystemDisk,omitempty" name:"SystemDisk"`

	// Data disk information.
	DataDisks []*DataDisk `json:"DataDisks,omitempty" name:"DataDisks"`
}

type ModuleCounter struct {
	// ISP statistics list
	ISPCounterSet []*ISPCounter `json:"ISPCounterSet,omitempty" name:"ISPCounterSet"`

	// Number of provinces/states
	ProvinceNum *int64 `json:"ProvinceNum,omitempty" name:"ProvinceNum"`

	// Number of cities
	CityNum *int64 `json:"CityNum,omitempty" name:"CityNum"`

	// Number of nodes
	NodeNum *int64 `json:"NodeNum,omitempty" name:"NodeNum"`

	// Number of instances
	InstanceNum *int64 `json:"InstanceNum,omitempty" name:"InstanceNum"`
}

type ModuleItem struct {
	// Instance statistics of the node
	NodeInstanceNum *NodeInstanceNum `json:"NodeInstanceNum,omitempty" name:"NodeInstanceNum"`

	// Module information
	Module *Module `json:"Module,omitempty" name:"Module"`
}

type MonthNetwork struct {
	// Zone information of the node
	ZoneInfo *ZoneInfo `json:"ZoneInfo,omitempty" name:"ZoneInfo"`

	// Bandwidth month, such as `202103`
	Month *string `json:"Month,omitempty" name:"Month"`

	// Bandwidth package ID format, such as `bwp-xxxxxxxx`
	BandwidthPkgId *string `json:"BandwidthPkgId,omitempty" name:"BandwidthPkgId"`

	// ISP abbreviation. Valid values: CUCC, CTCC, CMCC
	Isp *string `json:"Isp,omitempty" name:"Isp"`

	// Inbound bandwidth package peak. Value range: 0.0–xxx.xxx
	TrafficMaxIn *float64 `json:"TrafficMaxIn,omitempty" name:"TrafficMaxIn"`

	// Outbound bandwidth package peak. Value range: 0.0–xxx.xxx
	TrafficMaxOut *float64 `json:"TrafficMaxOut,omitempty" name:"TrafficMaxOut"`

	// Billable bandwidth. Value range: 0.0–xxx.xxx
	FeeTraffic *float64 `json:"FeeTraffic,omitempty" name:"FeeTraffic"`

	// Start time of the monthly billable bandwidth in the format of `yyyy-mm-dd HH:mm:ss`
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time of the monthly billable bandwidth in the format of `yyyy-mm-dd HH:mm:ss`
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Number of actual valid days for the monthly billable bandwidth, which must be an integer greater than or equal to 0
	EffectiveDays *int64 `json:"EffectiveDays,omitempty" name:"EffectiveDays"`

	// Actual number of days in the specified month, such as 30
	MonthDays *int64 `json:"MonthDays,omitempty" name:"MonthDays"`

	// Proportion of the number of valid days, accurate to four decimal places, such as `0.2134`
	EffectiveDaysRate *float64 `json:"EffectiveDaysRate,omitempty" name:"EffectiveDaysRate"`

	// Billable bandwidth package type. Valid values: Address, LoadBalance, AddressIpv6
	BandwidthPkgType *string `json:"BandwidthPkgType,omitempty" name:"BandwidthPkgType"`
}

type NetworkInterface struct {
	// ENI instance ID, such as `eni-f1xjkw1b`.
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`

	// ENI name.
	NetworkInterfaceName *string `json:"NetworkInterfaceName,omitempty" name:"NetworkInterfaceName"`

	// ENI description.
	NetworkInterfaceDescription *string `json:"NetworkInterfaceDescription,omitempty" name:"NetworkInterfaceDescription"`

	// Subnet instance ID.
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// VPC instance ID.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Bound security groups.
	// Note: this field may return null, indicating that no valid values can be obtained.
	GroupSet []*string `json:"GroupSet,omitempty" name:"GroupSet"`

	// Whether it is the primary ENI.
	Primary *bool `json:"Primary,omitempty" name:"Primary"`

	// MAC address.
	MacAddress *string `json:"MacAddress,omitempty" name:"MacAddress"`

	// ENI status:
	// PENDING: creating
	// AVAILABLE: available
	// ATTACHING: binding
	// DETACHING: unbinding
	// DELETING: deleting
	State *string `json:"State,omitempty" name:"State"`

	// Private IP information.
	// Note: this field may return null, indicating that no valid values can be obtained.
	PrivateIpAddressSet []*PrivateIpAddressSpecification `json:"PrivateIpAddressSet,omitempty" name:"PrivateIpAddressSet"`

	// Bound CVM object.
	// Note: this field may return null, indicating that no valid values can be obtained.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Attachment *NetworkInterfaceAttachment `json:"Attachment,omitempty" name:"Attachment"`

	// AZ.
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// Creation time.
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// List of IPv6 addresses.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Ipv6AddressSet []*Ipv6Address `json:"Ipv6AddressSet,omitempty" name:"Ipv6AddressSet"`

	// Tag key-value pairs.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TagSet []*Tag `json:"TagSet,omitempty" name:"TagSet"`

	// ENI type. Valid values: 0: ENI; 1: EVM ENI.
	EniType *uint64 `json:"EniType,omitempty" name:"EniType"`

	// ECM region (EcmRegion)
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// Type of the resource bound with an ENI. Valid values: `cvm` and `eks`.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Business *string `json:"Business,omitempty" name:"Business"`
}

type NetworkInterfaceAttachment struct {
	// CVM instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Serial number of the ENI in the CVM instance.
	DeviceIndex *uint64 `json:"DeviceIndex,omitempty" name:"DeviceIndex"`

	// Account information of the CVM instance owner.
	InstanceAccountId *string `json:"InstanceAccountId,omitempty" name:"InstanceAccountId"`

	// Binding time.
	AttachTime *string `json:"AttachTime,omitempty" name:"AttachTime"`
}

type NetworkStorageRange struct {
	// Network bandwidth cap
	MaxBandwidth *int64 `json:"MaxBandwidth,omitempty" name:"MaxBandwidth"`

	// Upper limit of the data disk size
	MaxSystemDiskSize *int64 `json:"MaxSystemDiskSize,omitempty" name:"MaxSystemDiskSize"`

	// Lower limit of the network bandwidth
	MinBandwidth *int64 `json:"MinBandwidth,omitempty" name:"MinBandwidth"`

	// Lower limit of the data disk size
	MinSystemDiskSize *int64 `json:"MinSystemDiskSize,omitempty" name:"MinSystemDiskSize"`

	// Maximum data disk size
	MaxDataDiskSize *int64 `json:"MaxDataDiskSize,omitempty" name:"MaxDataDiskSize"`

	// Minimum data disk size
	MinDataDiskSize *int64 `json:"MinDataDiskSize,omitempty" name:"MinDataDiskSize"`

	// Suggested bandwidth
	SuggestBandwidth *int64 `json:"SuggestBandwidth,omitempty" name:"SuggestBandwidth"`

	// Suggested disk size
	SuggestDataDiskSize *int64 `json:"SuggestDataDiskSize,omitempty" name:"SuggestDataDiskSize"`

	// Suggested system disk size
	SuggestSystemDiskSize *int64 `json:"SuggestSystemDiskSize,omitempty" name:"SuggestSystemDiskSize"`

	// Peak number of CPU cores
	MaxVcpu *int64 `json:"MaxVcpu,omitempty" name:"MaxVcpu"`

	// Minimum number of CPU cores
	MinVcpu *int64 `json:"MinVcpu,omitempty" name:"MinVcpu"`

	// Maximum number of CPU cores per request
	MaxVcpuPerReq *int64 `json:"MaxVcpuPerReq,omitempty" name:"MaxVcpuPerReq"`

	// Bandwidth increment
	PerBandwidth *int64 `json:"PerBandwidth,omitempty" name:"PerBandwidth"`

	// Data disk increment
	PerDataDisk *int64 `json:"PerDataDisk,omitempty" name:"PerDataDisk"`

	// Total number of modules
	MaxModuleNum *int64 `json:"MaxModuleNum,omitempty" name:"MaxModuleNum"`
}

type Node struct {
	// Zone information.
	ZoneInfo *ZoneInfo `json:"ZoneInfo,omitempty" name:"ZoneInfo"`

	// Country/Region information.
	Country *Country `json:"Country,omitempty" name:"Country"`

	// Region information.
	Area *Area `json:"Area,omitempty" name:"Area"`

	// Province/State information.
	Province *Province `json:"Province,omitempty" name:"Province"`

	// City information.
	City *City `json:"City,omitempty" name:"City"`

	// Region information.
	RegionInfo *RegionInfo `json:"RegionInfo,omitempty" name:"RegionInfo"`

	// List of ISPs.
	ISPSet []*ISP `json:"ISPSet,omitempty" name:"ISPSet"`

	// Number of ISPs.
	ISPNum *int64 `json:"ISPNum,omitempty" name:"ISPNum"`
}

type NodeInstanceNum struct {
	// Number of nodes
	NodeNum *int64 `json:"NodeNum,omitempty" name:"NodeNum"`

	// Number of instances
	InstanceNum *int64 `json:"InstanceNum,omitempty" name:"InstanceNum"`
}

type OperatorAction struct {
	// Executable operation
	Action *string `json:"Action,omitempty" name:"Action"`

	// Code
	// Note: this field may return null, indicating that no valid values can be obtained.
	Code *string `json:"Code,omitempty" name:"Code"`

	// Specific information
	// Note: this field may return null, indicating that no valid values can be obtained.
	Message *string `json:"Message,omitempty" name:"Message"`
}

type OsVersion struct {
	// OS type
	OsName *string `json:"OsName,omitempty" name:"OsName"`

	// Supported OS versions
	// Note: this field may return null, indicating that no valid values can be obtained.
	OsVersions []*string `json:"OsVersions,omitempty" name:"OsVersions"`

	// Supported OS architecture
	// Note: this field may return null, indicating that no valid values can be obtained.
	Architecture []*string `json:"Architecture,omitempty" name:"Architecture"`
}

type PackingQuotaGroup struct {
	// AZ
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// AZ ID
	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// ISP id
	ISPId *string `json:"ISPId,omitempty" name:"ISPId"`

	// A set of correlated packing quotas
	PackingQuotaInfos []*PackingQuotaInfo `json:"PackingQuotaInfos,omitempty" name:"PackingQuotaInfos"`
}

type PackingQuotaInfo struct {
	// Instance type
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// Packing quota
	PackingQuota *int64 `json:"PackingQuota,omitempty" name:"PackingQuota"`
}

type PeakBase struct {
	// Peak CPU
	PeakCpuNum *int64 `json:"PeakCpuNum,omitempty" name:"PeakCpuNum"`

	// Peak memory
	PeakMemoryNum *int64 `json:"PeakMemoryNum,omitempty" name:"PeakMemoryNum"`

	// Peak disk
	PeakStorageNum *int64 `json:"PeakStorageNum,omitempty" name:"PeakStorageNum"`

	// Recording time
	RecordTime *string `json:"RecordTime,omitempty" name:"RecordTime"`
}

type PeakFamilyInfo struct {
	// Model type information.
	InstanceFamily *InstanceFamilyTypeConfig `json:"InstanceFamily,omitempty" name:"InstanceFamily"`

	// Peak information of basic data.
	PeakBaseSet []*PeakBase `json:"PeakBaseSet,omitempty" name:"PeakBaseSet"`
}

type PeakNetwork struct {
	// Recording time.
	RecordTime *string `json:"RecordTime,omitempty" name:"RecordTime"`

	// Inbound bandwidth data.
	PeakInNetwork *string `json:"PeakInNetwork,omitempty" name:"PeakInNetwork"`

	// Outbound bandwidth data.
	PeakOutNetwork *string `json:"PeakOutNetwork,omitempty" name:"PeakOutNetwork"`

	// Billable bandwidth in bps
	ChargeNetwork *string `json:"ChargeNetwork,omitempty" name:"ChargeNetwork"`
}

type PeakNetworkRegionInfo struct {
	// Region information
	Region *string `json:"Region,omitempty" name:"Region"`

	// Peak network set
	// Note: this field may return null, indicating that no valid values can be obtained.
	PeakNetworkSet []*PeakNetwork `json:"PeakNetworkSet,omitempty" name:"PeakNetworkSet"`
}

type PhysicalPosition struct {
	// Rack unit
	// Note: this field may return null, indicating that no valid values can be obtained.
	PosId *string `json:"PosId,omitempty" name:"PosId"`

	// Rack
	// Note: this field may return null, indicating that no valid values can be obtained.
	RackId *string `json:"RackId,omitempty" name:"RackId"`

	// Switch
	// Note: this field may return null, indicating that no valid values can be obtained.
	SwitchId *string `json:"SwitchId,omitempty" name:"SwitchId"`
}

type Placement struct {
	// [AZ ID](https://intl.cloud.tencent.com/document/product/213/15753?from_cn_redirect=1#ZoneInfo) of the cloud disk, which can be obtained from the `Zone` field in the returned value of the [DescribeZones](https://intl.cloud.tencent.com/document/product/213/15707?from_cn_redirect=1) API.
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// Cage ID. When it is used as an input parameter, it indicates to manipulate the resources in the cage with the specified `CageId` and can be left empty. When it is used as an output parameter, it represents the cage ID of the resource and can be left empty.
	// Note: this field may return null, indicating that no valid values can be obtained.
	CageId *string `json:"CageId,omitempty" name:"CageId"`

	// Project ID of the instance, which can be obtained from the `projectId` field in the returned value of the [DescribeProject](https://intl.cloud.tencent.com/document/api/378/4400?from_cn_redirect=1) API. If this parameter is not specified, the default project ID will be used.
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Name of the dedicated cluster. When it is used as an input parameter, it is ignored. When it is used as an output parameter, it represents the name of the dedicated cluster to which the cloud disk belongs, and it can be left empty.
	// Note: this field may return null, indicating that no valid values can be obtained.
	CdcName *string `json:"CdcName,omitempty" name:"CdcName"`

	// Dedicated cluster ID of the instance. When it is used as an input parameter, it indicates to manipulate the resources in the dedicated cluster with the specified `CdcId` and can be left empty. When it is used as an output parameter, it represents the dedicated cluster ID of the resource and can be left empty.
	// Note: this field may return null, indicating that no valid values can be obtained.
	CdcId *string `json:"CdcId,omitempty" name:"CdcId"`
}

type Position struct {
	// Zone information of the instance.
	ZoneInfo *ZoneInfo `json:"ZoneInfo,omitempty" name:"ZoneInfo"`

	// Country/Region information of the instance.
	Country *Country `json:"Country,omitempty" name:"Country"`

	// Area information of the instance.
	Area *Area `json:"Area,omitempty" name:"Area"`

	// Province/State information of the instance.
	Province *Province `json:"Province,omitempty" name:"Province"`

	// City information of the instance.
	City *City `json:"City,omitempty" name:"City"`

	// Region information of the instance.
	RegionInfo *RegionInfo `json:"RegionInfo,omitempty" name:"RegionInfo"`
}

type PriceDetail struct {
	// Discount, such as `20`, which represents 80% off
	Discount *uint64 `json:"Discount,omitempty" name:"Discount"`

	// Discounted price in cents
	DiscountPrice *uint64 `json:"DiscountPrice,omitempty" name:"DiscountPrice"`

	// Original price in cents
	OriginalPrice *uint64 `json:"OriginalPrice,omitempty" name:"OriginalPrice"`
}

type PrivateIPAddressInfo struct {
	// Private IP of the instance.
	// Note: this field may return null, indicating that no valid values can be obtained.
	PrivateIPAddress *string `json:"PrivateIPAddress,omitempty" name:"PrivateIPAddress"`
}

type PrivateIpAddressSpecification struct {
	// Private IP address.
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" name:"PrivateIpAddress"`

	// Whether it is the primary IP.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Primary *bool `json:"Primary,omitempty" name:"Primary"`

	// Public IP address.
	PublicIpAddress *string `json:"PublicIpAddress,omitempty" name:"PublicIpAddress"`

	// EIP instance ID, such as `eip-11112222`.
	AddressId *string `json:"AddressId,omitempty" name:"AddressId"`

	// Private IP description.
	Description *string `json:"Description,omitempty" name:"Description"`

	// Whether the public IP is blocked.
	// Note: this field may return null, indicating that no valid values can be obtained.
	IsWanIpBlocked *bool `json:"IsWanIpBlocked,omitempty" name:"IsWanIpBlocked"`

	// IP status:
	// PENDING: generating
	// MIGRATING: migrating
	// DELETING: deleting
	// AVAILABLE: available
	State *string `json:"State,omitempty" name:"State"`
}

type Province struct {
	// Province/State ID
	ProvinceId *string `json:"ProvinceId,omitempty" name:"ProvinceId"`

	// Province/State name
	ProvinceName *string `json:"ProvinceName,omitempty" name:"ProvinceName"`
}

type PublicIPAddressInfo struct {
	// Billing mode.
	ChargeMode *string `json:"ChargeMode,omitempty" name:"ChargeMode"`

	// Public IP of the instance.
	PublicIPAddress *string `json:"PublicIPAddress,omitempty" name:"PublicIPAddress"`

	// Public IP ISP of the instance.
	ISP *ISP `json:"ISP,omitempty" name:"ISP"`

	// Outbound bandwidth cap of the instance in Mbps.
	MaxBandwidthOut *int64 `json:"MaxBandwidthOut,omitempty" name:"MaxBandwidthOut"`

	// Inbound bandwidth cap of the instance in Mbps.
	MaxBandwidthIn *int64 `json:"MaxBandwidthIn,omitempty" name:"MaxBandwidthIn"`
}

// Predefined struct for user
type RebootInstancesRequestParams struct {
	// List of IDs of the instances to be restarted. You can request up to 100 instances in a region at a time.
	InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet"`

	// Whether to force restart the instance after it failed to be restarted normally. Valid values:
	// TRUE: yes;
	// FALSE: no;
	// Default value: FALSE.
	ForceReboot *bool `json:"ForceReboot,omitempty" name:"ForceReboot"`

	// Shutdown type. Valid values:
	// SOFT: soft shutdown
	// HARD: hard shutdown
	// SOFT_FIRST: perform a soft shutdown first; if it fails, perform a hard shutdown
	// 
	// Default value: SOFT.
	StopType *string `json:"StopType,omitempty" name:"StopType"`
}

type RebootInstancesRequest struct {
	*tchttp.BaseRequest
	
	// List of IDs of the instances to be restarted. You can request up to 100 instances in a region at a time.
	InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet"`

	// Whether to force restart the instance after it failed to be restarted normally. Valid values:
	// TRUE: yes;
	// FALSE: no;
	// Default value: FALSE.
	ForceReboot *bool `json:"ForceReboot,omitempty" name:"ForceReboot"`

	// Shutdown type. Valid values:
	// SOFT: soft shutdown
	// HARD: hard shutdown
	// SOFT_FIRST: perform a soft shutdown first; if it fails, perform a hard shutdown
	// 
	// Default value: SOFT.
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
	delete(f, "InstanceIdSet")
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
	// Region
	Region *string `json:"Region,omitempty" name:"Region"`

	// Region name
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`

	// RegionID
	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`
}

// Predefined struct for user
type ReleaseAddressesRequestParams struct {
	// ECM region
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// List of unique IDs of EIPs.
	AddressIds []*string `json:"AddressIds,omitempty" name:"AddressIds"`
}

type ReleaseAddressesRequest struct {
	*tchttp.BaseRequest
	
	// ECM region
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// List of unique IDs of EIPs.
	AddressIds []*string `json:"AddressIds,omitempty" name:"AddressIds"`
}

func (r *ReleaseAddressesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReleaseAddressesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EcmRegion")
	delete(f, "AddressIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReleaseAddressesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReleaseAddressesResponseParams struct {
	// Async task ID (TaskId). You can use the `DescribeTaskResult` API to query the task status.
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ReleaseAddressesResponse struct {
	*tchttp.BaseResponse
	Response *ReleaseAddressesResponseParams `json:"Response"`
}

func (r *ReleaseAddressesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReleaseAddressesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReleaseIpv6AddressesRequestParams struct {
	// ECM region
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// ENI instance ID, such as `eni-m6dyj72l`.
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`

	// List of the specified IPv6 addresses. You can specify up to 10 IPv6 addresses at a time.
	Ipv6Addresses []*Ipv6Address `json:"Ipv6Addresses,omitempty" name:"Ipv6Addresses"`
}

type ReleaseIpv6AddressesRequest struct {
	*tchttp.BaseRequest
	
	// ECM region
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// ENI instance ID, such as `eni-m6dyj72l`.
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`

	// List of the specified IPv6 addresses. You can specify up to 10 IPv6 addresses at a time.
	Ipv6Addresses []*Ipv6Address `json:"Ipv6Addresses,omitempty" name:"Ipv6Addresses"`
}

func (r *ReleaseIpv6AddressesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReleaseIpv6AddressesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EcmRegion")
	delete(f, "NetworkInterfaceId")
	delete(f, "Ipv6Addresses")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReleaseIpv6AddressesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReleaseIpv6AddressesResponseParams struct {
	// Task ID. You can use the `DescribeTaskResult` API to query the task status
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ReleaseIpv6AddressesResponse struct {
	*tchttp.BaseResponse
	Response *ReleaseIpv6AddressesResponseParams `json:"Response"`
}

func (r *ReleaseIpv6AddressesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReleaseIpv6AddressesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemovePrivateIpAddressesRequestParams struct {
	// ECM region, such as `ap-xian-ecm`.
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// ENI instance ID, such as `eni-11112222`.
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`

	// Information of the specified private IPs. You can specify up to 10 IPs at a time.
	PrivateIpAddresses []*PrivateIpAddressSpecification `json:"PrivateIpAddresses,omitempty" name:"PrivateIpAddresses"`
}

type RemovePrivateIpAddressesRequest struct {
	*tchttp.BaseRequest
	
	// ECM region, such as `ap-xian-ecm`.
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// ENI instance ID, such as `eni-11112222`.
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`

	// Information of the specified private IPs. You can specify up to 10 IPs at a time.
	PrivateIpAddresses []*PrivateIpAddressSpecification `json:"PrivateIpAddresses,omitempty" name:"PrivateIpAddresses"`
}

func (r *RemovePrivateIpAddressesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemovePrivateIpAddressesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EcmRegion")
	delete(f, "NetworkInterfaceId")
	delete(f, "PrivateIpAddresses")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RemovePrivateIpAddressesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemovePrivateIpAddressesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RemovePrivateIpAddressesResponse struct {
	*tchttp.BaseResponse
	Response *RemovePrivateIpAddressesResponseParams `json:"Response"`
}

func (r *RemovePrivateIpAddressesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemovePrivateIpAddressesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReplaceRouteTableAssociationRequestParams struct {
	// Subnet instance ID, such as `subnet-3x5lf5q0`, which can be queried through the `DescribeSubnets` API.
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// Route table instance ID, such as `rtb-azd4dt1c`.
	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`

	// ECM region
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`
}

type ReplaceRouteTableAssociationRequest struct {
	*tchttp.BaseRequest
	
	// Subnet instance ID, such as `subnet-3x5lf5q0`, which can be queried through the `DescribeSubnets` API.
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// Route table instance ID, such as `rtb-azd4dt1c`.
	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`

	// ECM region
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`
}

func (r *ReplaceRouteTableAssociationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReplaceRouteTableAssociationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubnetId")
	delete(f, "RouteTableId")
	delete(f, "EcmRegion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReplaceRouteTableAssociationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReplaceRouteTableAssociationResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ReplaceRouteTableAssociationResponse struct {
	*tchttp.BaseResponse
	Response *ReplaceRouteTableAssociationResponseParams `json:"Response"`
}

func (r *ReplaceRouteTableAssociationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReplaceRouteTableAssociationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReplaceRoutesRequestParams struct {
	// Route table instance ID.
	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`

	// Routing policy object.
	Routes []*Route `json:"Routes,omitempty" name:"Routes"`
}

type ReplaceRoutesRequest struct {
	*tchttp.BaseRequest
	
	// Route table instance ID.
	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`

	// Routing policy object.
	Routes []*Route `json:"Routes,omitempty" name:"Routes"`
}

func (r *ReplaceRoutesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReplaceRoutesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RouteTableId")
	delete(f, "Routes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReplaceRoutesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReplaceRoutesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ReplaceRoutesResponse struct {
	*tchttp.BaseResponse
	Response *ReplaceRoutesResponseParams `json:"Response"`
}

func (r *ReplaceRoutesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReplaceRoutesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReplaceSecurityGroupPolicyRequestParams struct {
	// Security group instance ID, such as `esg-33ocnj9n`, which can be obtained through the `DescribeSecurityGroups` API
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`

	// Security group policy set object.
	SecurityGroupPolicySet *SecurityGroupPolicySet `json:"SecurityGroupPolicySet,omitempty" name:"SecurityGroupPolicySet"`
}

type ReplaceSecurityGroupPolicyRequest struct {
	*tchttp.BaseRequest
	
	// Security group instance ID, such as `esg-33ocnj9n`, which can be obtained through the `DescribeSecurityGroups` API
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`

	// Security group policy set object.
	SecurityGroupPolicySet *SecurityGroupPolicySet `json:"SecurityGroupPolicySet,omitempty" name:"SecurityGroupPolicySet"`
}

func (r *ReplaceSecurityGroupPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReplaceSecurityGroupPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecurityGroupId")
	delete(f, "SecurityGroupPolicySet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReplaceSecurityGroupPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReplaceSecurityGroupPolicyResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ReplaceSecurityGroupPolicyResponse struct {
	*tchttp.BaseResponse
	Response *ReplaceSecurityGroupPolicyResponseParams `json:"Response"`
}

func (r *ReplaceSecurityGroupPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReplaceSecurityGroupPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetInstancesMaxBandwidthRequestParams struct {
	// List of IDs of the instances for which to reset the bandwidth cap. You can request up to 100 instances in a region at a time.
	InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet"`

	// Modified outbound bandwidth cap.
	MaxBandwidthOut *int64 `json:"MaxBandwidthOut,omitempty" name:"MaxBandwidthOut"`

	// Modified inbound bandwidth cap.
	MaxBandwidthIn *int64 `json:"MaxBandwidthIn,omitempty" name:"MaxBandwidthIn"`
}

type ResetInstancesMaxBandwidthRequest struct {
	*tchttp.BaseRequest
	
	// List of IDs of the instances for which to reset the bandwidth cap. You can request up to 100 instances in a region at a time.
	InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet"`

	// Modified outbound bandwidth cap.
	MaxBandwidthOut *int64 `json:"MaxBandwidthOut,omitempty" name:"MaxBandwidthOut"`

	// Modified inbound bandwidth cap.
	MaxBandwidthIn *int64 `json:"MaxBandwidthIn,omitempty" name:"MaxBandwidthIn"`
}

func (r *ResetInstancesMaxBandwidthRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetInstancesMaxBandwidthRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIdSet")
	delete(f, "MaxBandwidthOut")
	delete(f, "MaxBandwidthIn")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetInstancesMaxBandwidthRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetInstancesMaxBandwidthResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ResetInstancesMaxBandwidthResponse struct {
	*tchttp.BaseResponse
	Response *ResetInstancesMaxBandwidthResponseParams `json:"Response"`
}

func (r *ResetInstancesMaxBandwidthResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetInstancesMaxBandwidthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetInstancesPasswordRequestParams struct {
	// List of IDs of the instances for which to set the password. You can request up to 100 instances in a region at a time.
	InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet"`

	// New password. The password of a Linux instance must contain 8–16 characters in at least two of the following character types: letters, digits, and special symbols [( ) ~ ~ ! @ # $ % ^ & * - + = _ | { } [ ] : ; ' < > , . ? /] and cannot start with `/`.
	// The password of a Windows instance must contain 12–16 characters in at least three of the following character types: letters, digits, and special symbols [( ) ~ ~ ! @ # $ % ^ & * - + = _ | { } [ ] : ; ' < > , . ? /] and cannot start with `/`.
	// If the instances include both Linux instances and Windows instances, the password complexity limit for Windows instances will apply.
	Password *string `json:"Password,omitempty" name:"Password"`

	// Whether to force shut down. Default value: false.
	ForceStop *bool `json:"ForceStop,omitempty" name:"ForceStop"`

	// Username of the instance for which to reset the password, which can contain up to 64 characters. If this parameter is not specified, the password of the root user will be reset by default for Linux, and the password of the admin will be reset by default for Windows.
	UserName *string `json:"UserName,omitempty" name:"UserName"`
}

type ResetInstancesPasswordRequest struct {
	*tchttp.BaseRequest
	
	// List of IDs of the instances for which to set the password. You can request up to 100 instances in a region at a time.
	InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet"`

	// New password. The password of a Linux instance must contain 8–16 characters in at least two of the following character types: letters, digits, and special symbols [( ) ~ ~ ! @ # $ % ^ & * - + = _ | { } [ ] : ; ' < > , . ? /] and cannot start with `/`.
	// The password of a Windows instance must contain 12–16 characters in at least three of the following character types: letters, digits, and special symbols [( ) ~ ~ ! @ # $ % ^ & * - + = _ | { } [ ] : ; ' < > , . ? /] and cannot start with `/`.
	// If the instances include both Linux instances and Windows instances, the password complexity limit for Windows instances will apply.
	Password *string `json:"Password,omitempty" name:"Password"`

	// Whether to force shut down. Default value: false.
	ForceStop *bool `json:"ForceStop,omitempty" name:"ForceStop"`

	// Username of the instance for which to reset the password, which can contain up to 64 characters. If this parameter is not specified, the password of the root user will be reset by default for Linux, and the password of the admin will be reset by default for Windows.
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
	delete(f, "InstanceIdSet")
	delete(f, "Password")
	delete(f, "ForceStop")
	delete(f, "UserName")
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
type ResetInstancesRequestParams struct {
	// List of IDs of the instances to be reinstalled.
	InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet"`

	// ID of the image from which to install the instance. If this parameter is not specified, the current image of the instance will be used.
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// Password. If this parameter is not specified, the password will be subsequently displayed in the Message Center.
	Password *string `json:"Password,omitempty" name:"Password"`

	// Whether to enable CM and CWP. If this parameter is not specified, they will be enabled by default.
	EnhancedService *EnhancedService `json:"EnhancedService,omitempty" name:"EnhancedService"`

	// Whether to retain the data on the data disk. Valid values: true, false. Default value: true
	KeepData *string `json:"KeepData,omitempty" name:"KeepData"`

	// Whether to keep the original settings for the image. You cannot specify this parameter if `Password` or `KeyIds.N` is specified. You can specify this parameter as `TRUE` only when you create an instance by using a custom image, shared image, or image imported from an external resource. Valid values:
	// TRUE: yes
	// FALSE: no
	// 
	// Default value: FALSE.
	// Note: this field may return null, indicating that no valid values can be obtained.
	KeepImageLogin *string `json:"KeepImageLogin,omitempty" name:"KeepImageLogin"`
}

type ResetInstancesRequest struct {
	*tchttp.BaseRequest
	
	// List of IDs of the instances to be reinstalled.
	InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet"`

	// ID of the image from which to install the instance. If this parameter is not specified, the current image of the instance will be used.
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// Password. If this parameter is not specified, the password will be subsequently displayed in the Message Center.
	Password *string `json:"Password,omitempty" name:"Password"`

	// Whether to enable CM and CWP. If this parameter is not specified, they will be enabled by default.
	EnhancedService *EnhancedService `json:"EnhancedService,omitempty" name:"EnhancedService"`

	// Whether to retain the data on the data disk. Valid values: true, false. Default value: true
	KeepData *string `json:"KeepData,omitempty" name:"KeepData"`

	// Whether to keep the original settings for the image. You cannot specify this parameter if `Password` or `KeyIds.N` is specified. You can specify this parameter as `TRUE` only when you create an instance by using a custom image, shared image, or image imported from an external resource. Valid values:
	// TRUE: yes
	// FALSE: no
	// 
	// Default value: FALSE.
	// Note: this field may return null, indicating that no valid values can be obtained.
	KeepImageLogin *string `json:"KeepImageLogin,omitempty" name:"KeepImageLogin"`
}

func (r *ResetInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIdSet")
	delete(f, "ImageId")
	delete(f, "Password")
	delete(f, "EnhancedService")
	delete(f, "KeepData")
	delete(f, "KeepImageLogin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetInstancesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ResetInstancesResponse struct {
	*tchttp.BaseResponse
	Response *ResetInstancesResponseParams `json:"Response"`
}

func (r *ResetInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetRoutesRequestParams struct {
	// Route table instance ID, such as `rtb-azd4dt1c`.
	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`

	// Route table name, which can contain up to 60 bytes.
	RouteTableName *string `json:"RouteTableName,omitempty" name:"RouteTableName"`

	// Routing policy.
	Routes []*Route `json:"Routes,omitempty" name:"Routes"`
}

type ResetRoutesRequest struct {
	*tchttp.BaseRequest
	
	// Route table instance ID, such as `rtb-azd4dt1c`.
	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`

	// Route table name, which can contain up to 60 bytes.
	RouteTableName *string `json:"RouteTableName,omitempty" name:"RouteTableName"`

	// Routing policy.
	Routes []*Route `json:"Routes,omitempty" name:"Routes"`
}

func (r *ResetRoutesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetRoutesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RouteTableId")
	delete(f, "RouteTableName")
	delete(f, "Routes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetRoutesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetRoutesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ResetRoutesResponse struct {
	*tchttp.BaseResponse
	Response *ResetRoutesResponseParams `json:"Response"`
}

func (r *ResetRoutesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetRoutesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Route struct {
	// Destination IPv4 IP range
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" name:"DestinationCidrBlock"`

	// Next hop type
	// NORMAL_CVM: general CVM;
	GatewayType *string `json:"GatewayType,omitempty" name:"GatewayType"`

	// Next hop address
	// You simply need to specify the gateway ID of a different next hop type, and the system will automatically match the next hop address
	// When `GatewayType` is `EIP`, the value of `GatewayId` will be fixed at `0`
	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`

	// Unique routing policy ID
	RouteItemId *string `json:"RouteItemId,omitempty" name:"RouteItemId"`

	// Routing policy description
	RouteDescription *string `json:"RouteDescription,omitempty" name:"RouteDescription"`

	// Whether to enable
	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`

	// Route type. Valid values:
	// USER: user route;
	// NETD: network probe route, which will be delivered by the system by default when you create a network probe instance and cannot be edited or deleted;
	// CCN: CCN route, which will be delivered by the system by default and cannot be edited or deleted.
	// You can only add and manipulate routes of `USER` type.
	RouteType *string `json:"RouteType,omitempty" name:"RouteType"`

	// Routing policy ID. The IPv4 routing policy will have a meaningful value, while the IPv6 routing policy is always 0. We recommend you use the unique ID `RouteItemId` for the routing policy
	RouteId *uint64 `json:"RouteId,omitempty" name:"RouteId"`
}

type RouteConflict struct {
	// Route table instance ID
	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`

	// The conflicting destination ports to be checked
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" name:"DestinationCidrBlock"`

	// List of conflicting routing policies
	// Note: this field may return null, indicating that no valid values can be obtained.
	ConflictSet []*Route `json:"ConflictSet,omitempty" name:"ConflictSet"`
}

type RouteTable struct {
	// VPC instance ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Route table instance ID
	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`

	// Route table name
	RouteTableName *string `json:"RouteTableName,omitempty" name:"RouteTableName"`

	// Association relationships of the route table
	// Note: this field may return null, indicating that no valid values can be obtained.
	AssociationSet []*RouteTableAssociation `json:"AssociationSet,omitempty" name:"AssociationSet"`

	// IPv4 routing policy set
	// Note: this field may return null, indicating that no valid values can be obtained.
	RouteSet []*Route `json:"RouteSet,omitempty" name:"RouteSet"`

	// Whether it is the default route table
	Main *bool `json:"Main,omitempty" name:"Main"`

	// Creation time
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
}

type RouteTableAssociation struct {
	// Subnet instance ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// Route table instance ID
	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`
}

type RuleHealth struct {
	// Health check status of the real server bound to the rule
	// Note: this field may return null, indicating that no valid values can be obtained.
	Targets []*TargetHealth `json:"Targets,omitempty" name:"Targets"`
}

type RunEIPDirectServiceEnabled struct {
	// Whether to enable IP direct access. Valid values:
	// TRUE: yes
	// FALSE: no
	// Default value: TRUE.
	// Currently, Windows images do not support IP direct access.
	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`
}

// Predefined struct for user
type RunInstancesRequestParams struct {
	// List of AZs in which to create instances, the number of instances to be created, and the ISPs. You can create up to 100 instances in a region at a time.
	ZoneInstanceCountISPSet []*ZoneInstanceCountISP `json:"ZoneInstanceCountISPSet,omitempty" name:"ZoneInstanceCountISPSet"`

	// Instance login password. Different OS types have different limits on password complexity as detailed below:
	// The password of a Linux instance must contain 8–30 characters in at least two of the following character types: letters, digits, and special symbols [( ) ` ~ ! @ # $ % ^ & - + = | { } [ ] : ; ' , . ? / ]. The password of a Windows instance must contain 12–30 characters in at least three of the following character types: letters, digits, and special symbols [( ) ` ~ ! @ # $ % ^ & - + = | { } [ ] : ; ' , . ? /].
	Password *string `json:"Password,omitempty" name:"Password"`

	// Public network outbound bandwidth cap in Mbps.
	// 1. If you don't specify this parameter or specify it as `0`, the default value under the corresponding module will be used.
	// 2. If you don't specify this parameter or specify it as `0` without specifying the module, the value of `InternetMaxBandwidthIn` will be used
	InternetMaxBandwidthOut *int64 `json:"InternetMaxBandwidthOut,omitempty" name:"InternetMaxBandwidthOut"`

	// Module ID. If you don't specify this parameter, you must specify the `ImageId`, `InstanceType`, `DataDiskSize`, and `InternetMaxBandwidthOut` parameters
	ModuleId *string `json:"ModuleId,omitempty" name:"ModuleId"`

	// Image ID. If you don't specify this parameter or specify it as null, the default value under the corresponding module will be used.
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// Instance display name.
	// If this parameter is not specified, `Not named` will be displayed by default.
	// If you purchase multiple instances and specify the pattern string `{R:x}`, display names will be generated based on `[x, x+n-1]`, where `n` is the number of the purchased instances. For example, if you specify `server\_{R:3}` and purchase 1 instance, the display name will be `server\_3`, and if you purchase 2 instances, the display names will be `server\_3` and `server\_4` respectively.
	// You can specify multiple pattern strings `{R:x}`.
	// If you purchase multiple instances and don't specify the pattern string, the instance display names will be suffixed with 1, 2...n, where `n` indicates the number of the purchased instances. For example, if you specify `server_` and purchase 2 instances, the instance display names will be `server\_1` and `server\_2` respectively.
	// If the purchased instances belong to different regions or ISPs, the above rules will apply to each region and ISP independently.
	// It can contain up to 60 characters (including pattern string).
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// Server name
	// `HostName` cannot start or end with a dot or hyphen and cannot contain consecutive dots or hyphens.
	// Windows instance: the name can contain 2–15 letters, digits, and hyphens but not dots or only digits.
	// Other types (such as Linux) of instances: the name should be a combination of 2 to 60 characters, supporting multiple dots. A string between two dots can contain letters, digits, and hyphens.
	HostName *string `json:"HostName,omitempty" name:"HostName"`

	// The string used to ensure the idempotency of the request. Currently, it is a reserved parameter; therefore, do not use it.
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`

	// Enhanced services. You can use this parameter to specify whether to enable services such as Cloud Security and Cloud Monitor. If this parameter is not specified, Cloud Monitor and Cloud Security will be enabled for public images by default.
	EnhancedService *EnhancedService `json:"EnhancedService,omitempty" name:"EnhancedService"`

	// Tag list
	TagSpecification []*TagSpecification `json:"TagSpecification,omitempty" name:"TagSpecification"`

	// The user data provided to the instance, which needs to be Base64-encoded with a maximum size of 16 KB
	UserData *string `json:"UserData,omitempty" name:"UserData"`

	// Model. If you don't specify this parameter or specify it as null, the default value under the corresponding module will be used.
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// Data disk size in GB. If you don't specify this parameter or specify it as `0`, the default value under the corresponding module will be used
	DataDiskSize *int64 `json:"DataDiskSize,omitempty" name:"DataDiskSize"`

	// Security group of the instance, which can be obtained from the `sgId` field in the returned value of the `DescribeSecurityGroups` API. If this parameter is not specified, the default security group will be bound by default.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`

	// System disk size in GB. If you don't specify this parameter or specify it as `0`, the default value under the corresponding module will be used
	SystemDiskSize *uint64 `json:"SystemDiskSize,omitempty" name:"SystemDiskSize"`

	// Public network inbound bandwidth cap in Mbps.
	// 1. If you don't specify this parameter or specify it as `0`, the default value under the corresponding module will be used.
	// 2. If you don't specify this parameter or specify it as `0` without specifying the module, the value of `InternetMaxBandwidthOut` will be used
	InternetMaxBandwidthIn *int64 `json:"InternetMaxBandwidthIn,omitempty" name:"InternetMaxBandwidthIn"`

	// Instance billing type. Valid values:
	// 0: postpaid by resource usage, where the daily peak usage of the CPU, memory, and disk will be calculated. This billing mode applies only to non-GNR models;
	// 1: hourly postpaid at the unit price of xx USD/instance/hour. This billing mode applies only to GNR models. To enable it, submit a ticket for application;
	// 2: monthly postpaid at the unit price of xx USD/instance/month. This billing mode applies only to GNR models;
	// If this field is left empty, `0` will be selected by default for non-GNR models, and `2` will be selected by default for GNR models.
	InstanceChargeType *int64 `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`

	// Key pair.
	KeyIds []*string `json:"KeyIds,omitempty" name:"KeyIds"`

	// Whether to keep the original settings for the image. You cannot specify this parameter if `Password` or `KeyIds.N` is specified. You can specify this parameter as `TRUE` only when you create an instance by using a custom image, shared image, or image imported from an external resource. Valid values:
	// TRUE: yes
	// FALSE: no
	// 
	// Default value: FALSE.
	// Note: this field may return null, indicating that no valid values can be obtained.
	KeepImageLogin *string `json:"KeepImageLogin,omitempty" name:"KeepImageLogin"`

	// System disk information.
	SystemDisk *SystemDisk `json:"SystemDisk,omitempty" name:"SystemDisk"`

	// Data disk information.
	DataDisks []*DataDisk `json:"DataDisks,omitempty" name:"DataDisks"`
}

type RunInstancesRequest struct {
	*tchttp.BaseRequest
	
	// List of AZs in which to create instances, the number of instances to be created, and the ISPs. You can create up to 100 instances in a region at a time.
	ZoneInstanceCountISPSet []*ZoneInstanceCountISP `json:"ZoneInstanceCountISPSet,omitempty" name:"ZoneInstanceCountISPSet"`

	// Instance login password. Different OS types have different limits on password complexity as detailed below:
	// The password of a Linux instance must contain 8–30 characters in at least two of the following character types: letters, digits, and special symbols [( ) ` ~ ! @ # $ % ^ & - + = | { } [ ] : ; ' , . ? / ]. The password of a Windows instance must contain 12–30 characters in at least three of the following character types: letters, digits, and special symbols [( ) ` ~ ! @ # $ % ^ & - + = | { } [ ] : ; ' , . ? /].
	Password *string `json:"Password,omitempty" name:"Password"`

	// Public network outbound bandwidth cap in Mbps.
	// 1. If you don't specify this parameter or specify it as `0`, the default value under the corresponding module will be used.
	// 2. If you don't specify this parameter or specify it as `0` without specifying the module, the value of `InternetMaxBandwidthIn` will be used
	InternetMaxBandwidthOut *int64 `json:"InternetMaxBandwidthOut,omitempty" name:"InternetMaxBandwidthOut"`

	// Module ID. If you don't specify this parameter, you must specify the `ImageId`, `InstanceType`, `DataDiskSize`, and `InternetMaxBandwidthOut` parameters
	ModuleId *string `json:"ModuleId,omitempty" name:"ModuleId"`

	// Image ID. If you don't specify this parameter or specify it as null, the default value under the corresponding module will be used.
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// Instance display name.
	// If this parameter is not specified, `Not named` will be displayed by default.
	// If you purchase multiple instances and specify the pattern string `{R:x}`, display names will be generated based on `[x, x+n-1]`, where `n` is the number of the purchased instances. For example, if you specify `server\_{R:3}` and purchase 1 instance, the display name will be `server\_3`, and if you purchase 2 instances, the display names will be `server\_3` and `server\_4` respectively.
	// You can specify multiple pattern strings `{R:x}`.
	// If you purchase multiple instances and don't specify the pattern string, the instance display names will be suffixed with 1, 2...n, where `n` indicates the number of the purchased instances. For example, if you specify `server_` and purchase 2 instances, the instance display names will be `server\_1` and `server\_2` respectively.
	// If the purchased instances belong to different regions or ISPs, the above rules will apply to each region and ISP independently.
	// It can contain up to 60 characters (including pattern string).
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// Server name
	// `HostName` cannot start or end with a dot or hyphen and cannot contain consecutive dots or hyphens.
	// Windows instance: the name can contain 2–15 letters, digits, and hyphens but not dots or only digits.
	// Other types (such as Linux) of instances: the name should be a combination of 2 to 60 characters, supporting multiple dots. A string between two dots can contain letters, digits, and hyphens.
	HostName *string `json:"HostName,omitempty" name:"HostName"`

	// The string used to ensure the idempotency of the request. Currently, it is a reserved parameter; therefore, do not use it.
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`

	// Enhanced services. You can use this parameter to specify whether to enable services such as Cloud Security and Cloud Monitor. If this parameter is not specified, Cloud Monitor and Cloud Security will be enabled for public images by default.
	EnhancedService *EnhancedService `json:"EnhancedService,omitempty" name:"EnhancedService"`

	// Tag list
	TagSpecification []*TagSpecification `json:"TagSpecification,omitempty" name:"TagSpecification"`

	// The user data provided to the instance, which needs to be Base64-encoded with a maximum size of 16 KB
	UserData *string `json:"UserData,omitempty" name:"UserData"`

	// Model. If you don't specify this parameter or specify it as null, the default value under the corresponding module will be used.
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// Data disk size in GB. If you don't specify this parameter or specify it as `0`, the default value under the corresponding module will be used
	DataDiskSize *int64 `json:"DataDiskSize,omitempty" name:"DataDiskSize"`

	// Security group of the instance, which can be obtained from the `sgId` field in the returned value of the `DescribeSecurityGroups` API. If this parameter is not specified, the default security group will be bound by default.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`

	// System disk size in GB. If you don't specify this parameter or specify it as `0`, the default value under the corresponding module will be used
	SystemDiskSize *uint64 `json:"SystemDiskSize,omitempty" name:"SystemDiskSize"`

	// Public network inbound bandwidth cap in Mbps.
	// 1. If you don't specify this parameter or specify it as `0`, the default value under the corresponding module will be used.
	// 2. If you don't specify this parameter or specify it as `0` without specifying the module, the value of `InternetMaxBandwidthOut` will be used
	InternetMaxBandwidthIn *int64 `json:"InternetMaxBandwidthIn,omitempty" name:"InternetMaxBandwidthIn"`

	// Instance billing type. Valid values:
	// 0: postpaid by resource usage, where the daily peak usage of the CPU, memory, and disk will be calculated. This billing mode applies only to non-GNR models;
	// 1: hourly postpaid at the unit price of xx USD/instance/hour. This billing mode applies only to GNR models. To enable it, submit a ticket for application;
	// 2: monthly postpaid at the unit price of xx USD/instance/month. This billing mode applies only to GNR models;
	// If this field is left empty, `0` will be selected by default for non-GNR models, and `2` will be selected by default for GNR models.
	InstanceChargeType *int64 `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`

	// Key pair.
	KeyIds []*string `json:"KeyIds,omitempty" name:"KeyIds"`

	// Whether to keep the original settings for the image. You cannot specify this parameter if `Password` or `KeyIds.N` is specified. You can specify this parameter as `TRUE` only when you create an instance by using a custom image, shared image, or image imported from an external resource. Valid values:
	// TRUE: yes
	// FALSE: no
	// 
	// Default value: FALSE.
	// Note: this field may return null, indicating that no valid values can be obtained.
	KeepImageLogin *string `json:"KeepImageLogin,omitempty" name:"KeepImageLogin"`

	// System disk information.
	SystemDisk *SystemDisk `json:"SystemDisk,omitempty" name:"SystemDisk"`

	// Data disk information.
	DataDisks []*DataDisk `json:"DataDisks,omitempty" name:"DataDisks"`
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
	delete(f, "ZoneInstanceCountISPSet")
	delete(f, "Password")
	delete(f, "InternetMaxBandwidthOut")
	delete(f, "ModuleId")
	delete(f, "ImageId")
	delete(f, "InstanceName")
	delete(f, "HostName")
	delete(f, "ClientToken")
	delete(f, "EnhancedService")
	delete(f, "TagSpecification")
	delete(f, "UserData")
	delete(f, "InstanceType")
	delete(f, "DataDiskSize")
	delete(f, "SecurityGroupIds")
	delete(f, "SystemDiskSize")
	delete(f, "InternetMaxBandwidthIn")
	delete(f, "InstanceChargeType")
	delete(f, "KeyIds")
	delete(f, "KeepImageLogin")
	delete(f, "SystemDisk")
	delete(f, "DataDisks")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RunInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RunInstancesResponseParams struct {
	// List of IDs of the instances being created
	// Note: this field may return null, indicating that no valid values can be obtained.
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
	// Whether to enable.
	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`
}

type RunSecurityServiceEnabled struct {
	// Whether to enable.
	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`

	// CWP edition. Valid values: 0: Basic Edition; 1: Pro Edition. Currently, only Basic Edition is supported
	Version *int64 `json:"Version,omitempty" name:"Version"`
}

type SecurityGroup struct {
	// Security group instance ID, such as `esg-ohuuioma`.
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`

	// Security group name, which can be customized with up to 60 characters.
	SecurityGroupName *string `json:"SecurityGroupName,omitempty" name:"SecurityGroupName"`

	// Security group remarks, which can contain up to 100 characters.
	SecurityGroupDesc *string `json:"SecurityGroupDesc,omitempty" name:"SecurityGroupDesc"`

	// Whether it is the default security group (which cannot be deleted).
	IsDefault *bool `json:"IsDefault,omitempty" name:"IsDefault"`

	// Security group creation time.
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// Tag key-value pairs.
	TagSet []*Tag `json:"TagSet,omitempty" name:"TagSet"`
}

type SecurityGroupAssociationStatistics struct {
	// Security group instance ID.
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`

	// Number of ECM instances.
	ECM *int64 `json:"ECM,omitempty" name:"ECM"`

	// Number of ECM modules.
	Module *int64 `json:"Module,omitempty" name:"Module"`

	// Number of ENI instances.
	ENI *int64 `json:"ENI,omitempty" name:"ENI"`

	// Number of times the security group is referenced by other security groups.
	SG *int64 `json:"SG,omitempty" name:"SG"`

	// Number of CLB instances.
	CLB *int64 `json:"CLB,omitempty" name:"CLB"`

	// Binding statistics of all instances.
	InstanceStatistics []*InstanceStatistic `json:"InstanceStatistics,omitempty" name:"InstanceStatistics"`

	// Total number of all resources (excluding resources referenced by security groups).
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

type SecurityGroupLimitSet struct {
	// Total number of security groups that can be created
	SecurityGroupLimit *int64 `json:"SecurityGroupLimit,omitempty" name:"SecurityGroupLimit"`

	// Maximum number of rules under the security group
	SecurityGroupPolicyLimit *int64 `json:"SecurityGroupPolicyLimit,omitempty" name:"SecurityGroupPolicyLimit"`

	// Number of nested security group rules under the security group
	ReferedSecurityGroupLimit *int64 `json:"ReferedSecurityGroupLimit,omitempty" name:"ReferedSecurityGroupLimit"`

	// Number of instances associated with the security group
	SecurityGroupInstanceLimit *int64 `json:"SecurityGroupInstanceLimit,omitempty" name:"SecurityGroupInstanceLimit"`

	// Number of security groups associated with the instance
	InstanceSecurityGroupLimit *int64 `json:"InstanceSecurityGroupLimit,omitempty" name:"InstanceSecurityGroupLimit"`

	// Number of modules associated with the security group
	SecurityGroupModuleLimit *int64 `json:"SecurityGroupModuleLimit,omitempty" name:"SecurityGroupModuleLimit"`

	// Number of security groups associated with the module
	ModuleSecurityGroupLimit *int64 `json:"ModuleSecurityGroupLimit,omitempty" name:"ModuleSecurityGroupLimit"`
}

type SecurityGroupPolicy struct {
	// Security group policy index number
	PolicyIndex *int64 `json:"PolicyIndex,omitempty" name:"PolicyIndex"`

	// Protocol. Valid values: TCP, UDP, ICMP.
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Port. Valid values: all, discrete port, range.
	Port *string `json:"Port,omitempty" name:"Port"`

	// Protocol port ID or protocol port group ID. `ServiceTemplate` and `Protocol+Port` are mutually exclusive.
	ServiceTemplate *ServiceTemplateSpecification `json:"ServiceTemplate,omitempty" name:"ServiceTemplate"`

	// IP range or IP address (mutually exclusive).
	CidrBlock *string `json:"CidrBlock,omitempty" name:"CidrBlock"`

	// Security group instance ID, such as `esg-ohuuioma`.
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`

	// IP address ID or IP address group ID.
	AddressTemplate *AddressTemplateSpecification `json:"AddressTemplate,omitempty" name:"AddressTemplate"`

	// `ACCEPT` or `DROP`.
	Action *string `json:"Action,omitempty" name:"Action"`

	// Security group policy description.
	PolicyDescription *string `json:"PolicyDescription,omitempty" name:"PolicyDescription"`

	// Modification time, such as `2020-07-22 19:27:23`
	// Note: this field may return null, indicating that no valid values can be obtained.
	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`

	// IP range or IPv6 address (mutually exclusive).
	// Note: this field may return null, indicating that no valid values can be obtained.
	Ipv6CidrBlock *string `json:"Ipv6CidrBlock,omitempty" name:"Ipv6CidrBlock"`
}

type SecurityGroupPolicySet struct {
	// The version number of the security group policy, which will automatically increase by one each time you update the security group policy, so as to prevent expiration of the updated routing policies. If it is left empty, any conflicts will be ignored.
	Version *string `json:"Version,omitempty" name:"Version"`

	// Outbound rule. You must select either an outbound rule or inbound rule.
	Egress []*SecurityGroupPolicy `json:"Egress,omitempty" name:"Egress"`

	// Inbound rule. You must select either outbound rule or inbound rule.
	Ingress []*SecurityGroupPolicy `json:"Ingress,omitempty" name:"Ingress"`
}

type ServiceTemplateSpecification struct {
	// Protocol port ID, such as `eppm-f5n1f8da`.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Protocol port group ID, such as `eppmg-f5n1f8da`.
	ServiceGroupId *string `json:"ServiceGroupId,omitempty" name:"ServiceGroupId"`
}

// Predefined struct for user
type SetLoadBalancerSecurityGroupsRequestParams struct {
	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// Array of security group IDs. You can bind up to 5 security groups to a CLB instance. If you want to unbind all security groups, leave this parameter empty or pass in an empty array
	SecurityGroups []*string `json:"SecurityGroups,omitempty" name:"SecurityGroups"`
}

type SetLoadBalancerSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// Array of security group IDs. You can bind up to 5 security groups to a CLB instance. If you want to unbind all security groups, leave this parameter empty or pass in an empty array
	SecurityGroups []*string `json:"SecurityGroups,omitempty" name:"SecurityGroups"`
}

func (r *SetLoadBalancerSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetLoadBalancerSecurityGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "SecurityGroups")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetLoadBalancerSecurityGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetLoadBalancerSecurityGroupsResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SetLoadBalancerSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *SetLoadBalancerSecurityGroupsResponseParams `json:"Response"`
}

func (r *SetLoadBalancerSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetLoadBalancerSecurityGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetSecurityGroupForLoadbalancersRequestParams struct {
	// Array of CLB instance IDs
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitempty" name:"LoadBalancerIds"`

	// Security group ID, such as `esg-12345678`
	SecurityGroup *string `json:"SecurityGroup,omitempty" name:"SecurityGroup"`

	// ADD: bind security group;
	// DEL: unbind security group
	OperationType *string `json:"OperationType,omitempty" name:"OperationType"`
}

type SetSecurityGroupForLoadbalancersRequest struct {
	*tchttp.BaseRequest
	
	// Array of CLB instance IDs
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitempty" name:"LoadBalancerIds"`

	// Security group ID, such as `esg-12345678`
	SecurityGroup *string `json:"SecurityGroup,omitempty" name:"SecurityGroup"`

	// ADD: bind security group;
	// DEL: unbind security group
	OperationType *string `json:"OperationType,omitempty" name:"OperationType"`
}

func (r *SetSecurityGroupForLoadbalancersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetSecurityGroupForLoadbalancersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerIds")
	delete(f, "SecurityGroup")
	delete(f, "OperationType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetSecurityGroupForLoadbalancersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetSecurityGroupForLoadbalancersResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SetSecurityGroupForLoadbalancersResponse struct {
	*tchttp.BaseResponse
	Response *SetSecurityGroupForLoadbalancersResponseParams `json:"Response"`
}

func (r *SetSecurityGroupForLoadbalancersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetSecurityGroupForLoadbalancersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SimpleModule struct {
	// Module ID
	ModuleId *string `json:"ModuleId,omitempty" name:"ModuleId"`

	// Module name
	ModuleName *string `json:"ModuleName,omitempty" name:"ModuleName"`
}

type Snapshot struct {
	// Snapshot location.
	Placement *Placement `json:"Placement,omitempty" name:"Placement"`

	// Whether the snapshot is replicated across regions. Valid values:<br><li>true: yes;<br><li>false: no.
	CopyFromRemote *bool `json:"CopyFromRemote,omitempty" name:"CopyFromRemote"`

	// Whether the snapshot is a permanent snapshot. Valid values:<br><li>true: yes<br><li>false: no.
	IsPermanent *bool `json:"IsPermanent,omitempty" name:"IsPermanent"`

	// Snapshot name, i.e., the user-defined snapshot alias. You can call [ModifySnapshotAttribute](https://intl.cloud.tencent.com/document/product/362/15650?from_cn_redirect=1) to modify this field.
	SnapshotName *string `json:"SnapshotName,omitempty" name:"SnapshotName"`

	// Snapshot creation progress in percentage. This field will always be `100` once the snapshot is created successfully.
	Percent *uint64 `json:"Percent,omitempty" name:"Percent"`

	// List of images associated with the snapshot.
	Images []*Image `json:"Images,omitempty" name:"Images"`

	// Number of snapshots currently shared.
	ShareReference *uint64 `json:"ShareReference,omitempty" name:"ShareReference"`

	// Snapshot type. Valid values: PRIVATE_SNAPSHOT, SHARED_SNAPSHOT
	SnapshotType *string `json:"SnapshotType,omitempty" name:"SnapshotType"`

	// Size in GB of the cloud disk for which the snapshot is created.
	DiskSize *uint64 `json:"DiskSize,omitempty" name:"DiskSize"`

	// ID of the cloud disk for which the snapshot is created.
	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`

	// Destination region to which the snapshot is being replicated. Default value: [].
	CopyingToRegions []*string `json:"CopyingToRegions,omitempty" name:"CopyingToRegions"`

	// Snapshot ID.
	SnapshotId *string `json:"SnapshotId,omitempty" name:"SnapshotId"`

	// Type of the cloud disk for which the snapshot is created. Valid values:<br><li>SYSTEM_DISK: system disk<br><li>DATA_DISK: data disk.
	DiskUsage *string `json:"DiskUsage,omitempty" name:"DiskUsage"`

	// Whether the snapshot is created from an encrypted disk. Valid values:<br><li>true: yes<br><li>false: no.
	Encrypt *bool `json:"Encrypt,omitempty" name:"Encrypt"`

	// Snapshot creation time.
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Number of images associated with the snapshot.
	ImageCount *uint64 `json:"ImageCount,omitempty" name:"ImageCount"`

	// Snapshot status. Valid values:<br><li>NORMAL: normal<br><li>CREATING: creating<br><li>ROLLBACKING: rolling back<br><li>COPYING_FROM_REMOTE: cross-region replicating<br><li>CHECKING_COPIED: verifying the cross-region replicated data<br><li>TORECYCLE: to be repossessed.
	SnapshotState *string `json:"SnapshotState,omitempty" name:"SnapshotState"`

	// Snapshot expiration time.
	DeadlineTime *string `json:"DeadlineTime,omitempty" name:"DeadlineTime"`

	// Time when snapshot sharing starts.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TimeStartShare *string `json:"TimeStartShare,omitempty" name:"TimeStartShare"`
}

type SrcImage struct {
	// Image ID
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// Image name
	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`

	// System name
	ImageOsName *string `json:"ImageOsName,omitempty" name:"ImageOsName"`

	// Image description
	ImageDescription *string `json:"ImageDescription,omitempty" name:"ImageDescription"`

	// Region
	Region *string `json:"Region,omitempty" name:"Region"`

	// Region ID
	RegionID *int64 `json:"RegionID,omitempty" name:"RegionID"`

	// Region name
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`

	// Source instance name
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// Source instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Source image type
	ImageType *string `json:"ImageType,omitempty" name:"ImageType"`
}

// Predefined struct for user
type StartInstancesRequestParams struct {
	// List of IDs of the instances to be started. You can request up to 100 instances in a region at a time.
	InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet"`
}

type StartInstancesRequest struct {
	*tchttp.BaseRequest
	
	// List of IDs of the instances to be started. You can request up to 100 instances in a region at a time.
	InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet"`
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
	delete(f, "InstanceIdSet")
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
	// List of IDs of the instances to be shut down. You can request up to 100 instances in a region at a time.
	InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet"`

	// Whether to force shut down the instance after it failed to be shut down normally. Default value: false: no.
	ForceStop *bool `json:"ForceStop,omitempty" name:"ForceStop"`

	// Instance shutdown mode. Valid values:
	// SOFT_FIRST: perform a forced shutdown in case of a failure of the normal shutdown;
	// HARD: forced shutdown;
	// SOFT: Soft shutdown;
	// Default value: SOFT.
	StopType *string `json:"StopType,omitempty" name:"StopType"`
}

type StopInstancesRequest struct {
	*tchttp.BaseRequest
	
	// List of IDs of the instances to be shut down. You can request up to 100 instances in a region at a time.
	InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet"`

	// Whether to force shut down the instance after it failed to be shut down normally. Default value: false: no.
	ForceStop *bool `json:"ForceStop,omitempty" name:"ForceStop"`

	// Instance shutdown mode. Valid values:
	// SOFT_FIRST: perform a forced shutdown in case of a failure of the normal shutdown;
	// HARD: forced shutdown;
	// SOFT: Soft shutdown;
	// Default value: SOFT.
	StopType *string `json:"StopType,omitempty" name:"StopType"`
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
	delete(f, "InstanceIdSet")
	delete(f, "ForceStop")
	delete(f, "StopType")
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

type Subnet struct {
	// VPC instance ID.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Subnet instance ID, such as `subnet-bthucmmy`.
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// Subnet name.
	SubnetName *string `json:"SubnetName,omitempty" name:"SubnetName"`

	// IPv4 CIDR block of the subnet.
	CidrBlock *string `json:"CidrBlock,omitempty" name:"CidrBlock"`

	// Whether it is the default subnet.
	IsDefault *bool `json:"IsDefault,omitempty" name:"IsDefault"`

	// Whether to enable broadcast.
	EnableBroadcast *bool `json:"EnableBroadcast,omitempty" name:"EnableBroadcast"`

	// Route table instance ID, such as `rtb-l2h8d7c2`.
	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`

	// Creation time.
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// Number of available IPs.
	AvailableIpAddressCount *uint64 `json:"AvailableIpAddressCount,omitempty" name:"AvailableIpAddressCount"`

	// IPv6 CIDR block of the subnet.
	Ipv6CidrBlock *string `json:"Ipv6CidrBlock,omitempty" name:"Ipv6CidrBlock"`

	// Associated ACLID
	NetworkAclId *string `json:"NetworkAclId,omitempty" name:"NetworkAclId"`

	// Whether it is an SNAT address pool subnet.
	IsRemoteVpcSnat *bool `json:"IsRemoteVpcSnat,omitempty" name:"IsRemoteVpcSnat"`

	// Tag key-value pairs.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TagSet []*Tag `json:"TagSet,omitempty" name:"TagSet"`

	// Region
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// AZ name
	// Note: this field may return null, indicating that no valid values can be obtained.
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// Number of instances
	// Note: this field may return null, indicating that no valid values can be obtained.
	InstanceCount *uint64 `json:"InstanceCount,omitempty" name:"InstanceCount"`

	// IPv4 CIDR block of the VPC.
	// Note: this field may return null, indicating that no valid values can be obtained.
	VpcCidrBlock *string `json:"VpcCidrBlock,omitempty" name:"VpcCidrBlock"`

	// IPv6 CIDR block of the VPC.
	// Note: this field may return null, indicating that no valid values can be obtained.
	VpcIpv6CidrBlock *string `json:"VpcIpv6CidrBlock,omitempty" name:"VpcIpv6CidrBlock"`

	// Region
	// Note: this field may return null, indicating that no valid values can be obtained.
	Region *string `json:"Region,omitempty" name:"Region"`
}

type SystemDisk struct {
	// Disk type. Valid values:
	// - LOCAL_BASIC: local disk;
	// - CLOUD_PREMIUM: Premium Cloud Storage;
	// Default value: CLOUD_BASIC.
	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`

	// Disk ID. This parameter is temporarily unavailable.
	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`

	// Disk size in GB.
	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`
}

type Tag struct {
	// Tag key.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Key *string `json:"Key,omitempty" name:"Key"`

	// Tag value.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Value *string `json:"Value,omitempty" name:"Value"`
}

type TagInfo struct {
	// Tag key.
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// Tag value.
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type TagSpecification struct {
	// Resource type. Valid values: instance, module
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`

	// Tag list
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

type Target struct {
	// Listening port of the real server
	// Note: this field may return null, indicating that no valid values can be obtained.
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// CVM instance ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Forwarding weight of the real server. Value range: [0, 100]. Default value: 10.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Weight *int64 `json:"Weight,omitempty" name:"Weight"`

	// You need to pass in this parameter when binding an ENI. It represents the IP address of the ENI. You must bind an ENI to a CVM instance first before you can bind it to a CLB instance. Note: you must pass in either `InstanceId` or `EniIp`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	EniIp *string `json:"EniIp,omitempty" name:"EniIp"`
}

type TargetHealth struct {
	// Private IP of the target
	// Note: this field may return null, indicating that no valid values can be obtained.
	IP *string `json:"IP,omitempty" name:"IP"`

	// Port bound to the target
	// Note: this field may return null, indicating that no valid values can be obtained.
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// Current health status. Valid values: true: healthy; false: unhealthy (e.g., check not started, checking, or exceptional status). CLB instance will route traffic to only healthy real servers whose weights are greater than 0.
	// Note: this field may return null, indicating that no valid values can be obtained.
	HealthStatus *bool `json:"HealthStatus,omitempty" name:"HealthStatus"`

	// Instance ID of the target
	// Note: this field may return null, indicating that no valid values can be obtained.
	TargetId *string `json:"TargetId,omitempty" name:"TargetId"`

	// Detailed information of the current health status. Valid values: Alive: healthy; Dead: exceptional; Unknown: check not started/checking/unknown status; Close: health check not configured.
	// Note: this field may return null, indicating that no valid values can be obtained.
	HealthStatusDetail *string `json:"HealthStatusDetail,omitempty" name:"HealthStatusDetail"`
}

type TargetsWeightRule struct {
	// CLB listener ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// List of real servers for which to modify the weights
	// Note: this field may return null, indicating that no valid values can be obtained.
	Targets []*Target `json:"Targets,omitempty" name:"Targets"`

	// New forwarding weight of the real server. Value range: 0–100.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Weight *int64 `json:"Weight,omitempty" name:"Weight"`
}

type TaskInput struct {
	// Operation name, i.e., API name, such as `CreateImage`
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// Task ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

type TaskOutput struct {
	// Task ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// Status description
	Message *string `json:"Message,omitempty" name:"Message"`

	// Status value. Valid values: SUCCESS, FAILED, OPERATING
	Status *string `json:"Status,omitempty" name:"Status"`

	// Task submission time
	AddTime *string `json:"AddTime,omitempty" name:"AddTime"`

	// Task end time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Operation name
	Operation *string `json:"Operation,omitempty" name:"Operation"`
}

// Predefined struct for user
type TerminateInstancesRequestParams struct {
	// List of IDs of the instances to be terminated.
	InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet"`

	// Whether to enable scheduled termination. Default value: no.
	TerminateDelay *bool `json:"TerminateDelay,omitempty" name:"TerminateDelay"`

	// Scheduled termination time, such as `2019-08-05 12:01:30`. If you don't enable scheduled termination, you can ignore this parameter.
	TerminateTime *string `json:"TerminateTime,omitempty" name:"TerminateTime"`

	// Whether to delete the bound ENI and EIP. Default value: true.
	// true: the ENI and EIP will also be deleted;
	// false: only the server will be terminated, while the ENI and EIP will be retained.
	AssociatedResourceDestroy *bool `json:"AssociatedResourceDestroy,omitempty" name:"AssociatedResourceDestroy"`
}

type TerminateInstancesRequest struct {
	*tchttp.BaseRequest
	
	// List of IDs of the instances to be terminated.
	InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet"`

	// Whether to enable scheduled termination. Default value: no.
	TerminateDelay *bool `json:"TerminateDelay,omitempty" name:"TerminateDelay"`

	// Scheduled termination time, such as `2019-08-05 12:01:30`. If you don't enable scheduled termination, you can ignore this parameter.
	TerminateTime *string `json:"TerminateTime,omitempty" name:"TerminateTime"`

	// Whether to delete the bound ENI and EIP. Default value: true.
	// true: the ENI and EIP will also be deleted;
	// false: only the server will be terminated, while the ENI and EIP will be retained.
	AssociatedResourceDestroy *bool `json:"AssociatedResourceDestroy,omitempty" name:"AssociatedResourceDestroy"`
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
	delete(f, "InstanceIdSet")
	delete(f, "TerminateDelay")
	delete(f, "TerminateTime")
	delete(f, "AssociatedResourceDestroy")
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
	// VPC ID, such as `vpc-xxx`.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Subnet ID of the VPC, such as `subnet-xxx`.
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// Whether it is used as a public gateway. The public gateway can be used only when the instance has a public IP and resides in a VPC. Valid values:
	// TRUE: yes
	// FALSE: no
	// 
	// Default value: FALSE.
	AsVpcGateway *bool `json:"AsVpcGateway,omitempty" name:"AsVpcGateway"`

	// Array of VPC subnet IPs. This parameter can be used to create instances or modify the VPC attributes of instances.
	PrivateIpAddresses []*string `json:"PrivateIpAddresses,omitempty" name:"PrivateIpAddresses"`

	// Number of the IPv6 addresses to be randomly generated for the ENI.
	Ipv6AddressCount *int64 `json:"Ipv6AddressCount,omitempty" name:"Ipv6AddressCount"`
}

type VpcInfo struct {
	// VPC name.
	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`

	// VPC instance ID, such as `vpc-azd4dt1c`.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// IPv4 CIDR block of the VPC.
	CidrBlock *string `json:"CidrBlock,omitempty" name:"CidrBlock"`

	// Whether it is the default VPC.
	IsDefault *bool `json:"IsDefault,omitempty" name:"IsDefault"`

	// Whether to enable multicast.
	EnableMulticast *bool `json:"EnableMulticast,omitempty" name:"EnableMulticast"`

	// Creation time.
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// List of DNS servers.
	// Note: this field may return null, indicating that no valid values can be obtained.
	DnsServerSet []*string `json:"DnsServerSet,omitempty" name:"DnsServerSet"`

	// DHCP domain option value.
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// DHCP option set ID.
	DhcpOptionsId *string `json:"DhcpOptionsId,omitempty" name:"DhcpOptionsId"`

	// Whether to enable DHCP.
	EnableDhcp *bool `json:"EnableDhcp,omitempty" name:"EnableDhcp"`

	// IPv6 CIDR block of the VPC.
	Ipv6CidrBlock *string `json:"Ipv6CidrBlock,omitempty" name:"Ipv6CidrBlock"`

	// Tag key-value pair
	// Note: this field may return null, indicating that no valid values can be obtained.
	TagSet []*Tag `json:"TagSet,omitempty" name:"TagSet"`

	// Secondary CIDR block
	// Note: this field may return null, indicating that no valid values can be obtained.
	AssistantCidrSet []*AssistantCidr `json:"AssistantCidrSet,omitempty" name:"AssistantCidrSet"`

	// Region
	Region *string `json:"Region,omitempty" name:"Region"`

	// Description
	Description *string `json:"Description,omitempty" name:"Description"`

	// Region name
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`

	// Number of included subnets
	SubnetCount *uint64 `json:"SubnetCount,omitempty" name:"SubnetCount"`

	// Number of included instances
	InstanceCount *uint64 `json:"InstanceCount,omitempty" name:"InstanceCount"`
}

type ZoneInfo struct {
	// ZoneId
	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// ZoneName
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// Zone
	Zone *string `json:"Zone,omitempty" name:"Zone"`
}

type ZoneInstanceCountISP struct {
	// The AZ in which to create an instance.
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// Number of instances to be created in the current AZ.
	InstanceCount *int64 `json:"InstanceCount,omitempty" name:"InstanceCount"`

	// ISP name. Valid values:
	// CTCC: China Telecom
	// CUCC: China Unicom
	// CMCC: China Mobile
	// If there are multiple ISP names, you need to separate them by semicolons, such as `CMCC;CUCC;CTCC`. To use multiple ISPs, contact Tencent Cloud customer service for assistance.
	ISP *string `json:"ISP,omitempty" name:"ISP"`

	// ID of the specified VPC. You must specify both `SubnetId` and `VpcId` or neither
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// ID of the specified subnet. You must specify both `SubnetId` and `VpcId` or neither
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// Private IP of the specified primary ENI. You must specify both `SubnetId` and `VpcId` at the same time. The number of IP addresses must be the same as `InstanceCount`. You can get the private IP of the secondary ENI of a multi-IP server through DHCP in the same subnet.
	PrivateIpAddresses []*string `json:"PrivateIpAddresses,omitempty" name:"PrivateIpAddresses"`

	// Number of the IPv6 addresses to be randomly generated for the ENI, which cannot be greater than 1.
	Ipv6AddressCount *int64 `json:"Ipv6AddressCount,omitempty" name:"Ipv6AddressCount"`
}

type ZoneInstanceInfo struct {
	// Zone name
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// Number of instances
	InstanceNum *int64 `json:"InstanceNum,omitempty" name:"InstanceNum"`
}