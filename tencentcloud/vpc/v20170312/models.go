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

type AcceptAttachCcnInstancesRequest struct {
	*tchttp.BaseRequest

	// The CCN instance ID, such as `ccn-f49l6u0z`.
	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`

	// List of associated instances.
	Instances []*CcnInstance `json:"Instances,omitempty" name:"Instances" list`
}

func (r *AcceptAttachCcnInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AcceptAttachCcnInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AcceptAttachCcnInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AcceptAttachCcnInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AcceptAttachCcnInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AccountAttribute struct {

	// Attribute name
	AttributeName *string `json:"AttributeName,omitempty" name:"AttributeName"`

	// Attribute values
	AttributeValues []*string `json:"AttributeValues,omitempty" name:"AttributeValues" list`
}

type AddBandwidthPackageResourcesRequest struct {
	*tchttp.BaseRequest

	// The unique ID of the source, such as 'eip-xxxx' and 'lb-xxxx'. EIP and LB resources are currently supported.
	ResourceIds []*string `json:"ResourceIds,omitempty" name:"ResourceIds" list`

	// The unique ID of the bandwidth package, such as 'bwp-xxxx'.
	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" name:"BandwidthPackageId"`

	// The network type of the bandwidth package. Valid value: `BGP`, indicating that the internal resource is a BGP IP.
	NetworkType *string `json:"NetworkType,omitempty" name:"NetworkType"`

	// The resource type, including `Address` and `LoadBalance`.
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`

	// The protocol type of the bandwidth package. Valid values: `ipv4` and `ipv6`.
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
}

func (r *AddBandwidthPackageResourcesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddBandwidthPackageResourcesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddBandwidthPackageResourcesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddBandwidthPackageResourcesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddBandwidthPackageResourcesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Address struct {

	// `EIP` `ID`, the unique ID of the `EIP`.
	AddressId *string `json:"AddressId,omitempty" name:"AddressId"`

	// The `EIP` name.
	AddressName *string `json:"AddressName,omitempty" name:"AddressName"`

	// Possible `EIP` states are 'CREATING', 'BINDING', 'BIND', 'UNBINDING', 'UNBIND', 'OFFLINING', and 'BIND_ENI'.
	AddressStatus *string `json:"AddressStatus,omitempty" name:"AddressStatus"`

	// The public IP address
	AddressIp *string `json:"AddressIp,omitempty" name:"AddressIp"`

	// The ID of the bound resource instance. This can be a `CVM` or `NAT`.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// The creation time, which follows the `ISO8601` standard and uses `UTC` time in the format of `YYYY-MM-DDThh:mm:ssZ`.
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// The ID of the bound ENI
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`

	// The private IP of the bound resources
	PrivateAddressIp *string `json:"PrivateAddressIp,omitempty" name:"PrivateAddressIp"`

	// The isolation status of the resource. `True` indicates the EIP is isolated. `False` indicates that the resource is not isolated.
	IsArrears *bool `json:"IsArrears,omitempty" name:"IsArrears"`

	// The block status of the resource. `True` indicates the EIP is blocked. `False` indicates that the EIP is not blocked.
	IsBlocked *bool `json:"IsBlocked,omitempty" name:"IsBlocked"`

	// Whether the EIP supports direct connection mode. `True` indicates the EIP supports direct connection. `False` indicates that the resource does not support direct connection.
	IsEipDirectConnection *bool `json:"IsEipDirectConnection,omitempty" name:"IsEipDirectConnection"`

	// The resource type of the EIP. This includes `CalcIP`, `WanIP`, `EIP`, and `AnycastEIP`. Among these, `CalcIP` indicates the device IP, `WanIP` indicates the common public IP, `EIP` indicates Elastic IP, and `AnycastEip` indicates accelerated EIP.
	AddressType *string `json:"AddressType,omitempty" name:"AddressType"`

	// Whether the EIP is automatically released after being unbound. `True` indicates the EIP will be automatically released after being unbound. `False` indicates the EIP will not be automatically released after being unbound.
	CascadeRelease *bool `json:"CascadeRelease,omitempty" name:"CascadeRelease"`

	// Type of the protocol used in EIP ALG
	EipAlgType *AlgType `json:"EipAlgType,omitempty" name:"EipAlgType"`

	// The ISP of an EIP/Elastic IP, with possible return values currently including "CMCC", "CTCC", "CUCC" and "BGP"
	InternetServiceProvider *string `json:"InternetServiceProvider,omitempty" name:"InternetServiceProvider"`

	// Whether the EIP is in a local BGP.
	LocalBgp *bool `json:"LocalBgp,omitempty" name:"LocalBgp"`

	// Bandwidth value of EIP. The EIP for the bill-by-CVM account will return `null`.
	// Note: this field may return `null`, indicating that no valid value was found.
	Bandwidth *uint64 `json:"Bandwidth,omitempty" name:"Bandwidth"`

	// Network billing mode of EIP. The EIP for the bill-by-CVM account will return `null`.
	// Note: this field may return `null`, indicating that no valid value was found.
	InternetChargeType *string `json:"InternetChargeType,omitempty" name:"InternetChargeType"`
}

type AddressChargePrepaid struct {

	// Purchased usage period, in month. Valid values: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// Setting of renewal. Valid values: 0: manual renewal; 1: auto-renewal; 2: no renewal after expiration. Default value: 0
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`
}

type AddressTemplate struct {

	// IP address template name.
	AddressTemplateName *string `json:"AddressTemplateName,omitempty" name:"AddressTemplateName"`

	// The unique ID of the IP address template instance.
	AddressTemplateId *string `json:"AddressTemplateId,omitempty" name:"AddressTemplateId"`

	// IP address information.
	AddressSet []*string `json:"AddressSet,omitempty" name:"AddressSet" list`

	// Creation Time.
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
}

type AddressTemplateGroup struct {

	// IP address template group name.
	AddressTemplateGroupName *string `json:"AddressTemplateGroupName,omitempty" name:"AddressTemplateGroupName"`

	// IP address template group instance ID, such as `ipmg-dih8xdbq`.
	AddressTemplateGroupId *string `json:"AddressTemplateGroupId,omitempty" name:"AddressTemplateGroupId"`

	// IP address template ID.
	AddressTemplateIdSet []*string `json:"AddressTemplateIdSet,omitempty" name:"AddressTemplateIdSet" list`

	// Creation Time.
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// IP address template instance
	AddressTemplateSet []*AddressTemplateItem `json:"AddressTemplateSet,omitempty" name:"AddressTemplateSet" list`
}

type AddressTemplateItem struct {

	// Start address
	From *string `json:"From,omitempty" name:"From"`

	// End address
	To *string `json:"To,omitempty" name:"To"`
}

type AddressTemplateSpecification struct {

	// The ID of the IP address, such as `ipm-2uw6ujo6`.
	AddressId *string `json:"AddressId,omitempty" name:"AddressId"`

	// The ID of the IP address group, such as `ipmg-2uw6ujo6`.
	AddressGroupId *string `json:"AddressGroupId,omitempty" name:"AddressGroupId"`
}

type AlgType struct {

	// Whether FTP ALG is enabled
	Ftp *bool `json:"Ftp,omitempty" name:"Ftp"`

	// Whether SIP ALG is enabled
	Sip *bool `json:"Sip,omitempty" name:"Sip"`
}

type AllocateAddressesRequest struct {
	*tchttp.BaseRequest

	// The number of EIPs. Default: 1.
	AddressCount *int64 `json:"AddressCount,omitempty" name:"AddressCount"`

	// The EIP line type. Default: BGP.
	// <ul style="margin:0"><li>For a user who has activated the static single-line IP allowlist, possible values are:<ul><li>CMCC: China Mobile</li>
	// <li>CTCC: China Telecom</li>
	// <li>CUCC: China Unicom</li></ul>Note: Only certain regions support static single-line IP addresses.</li></ul>
	InternetServiceProvider *string `json:"InternetServiceProvider,omitempty" name:"InternetServiceProvider"`

	// The EIP billing method.
	// <ul style="margin:0"><li>For bill-by-IP account beta users, valid values: <ul><li>BANDWIDTH_PACKAGE: paid by the [bandwidth package](https://intl.cloud.tencent.com/document/product/684/15255?from_cn_redirect=1)(who must also be bandwidth package beta users)</li>
	// <li>BANDWIDTH_POSTPAID_BY_HOUR: billed by hourly bandwidth on a pay-as-you-go basis</li>
	// <li>BANDWIDTH_PREPAID_BY_MONTH: monthly bandwidth subscription</li>
	// <li>TRAFFIC_POSTPAID_BY_HOUR: billed by hourly traffic on a pay-as-you-go basis</li></ul>Default value: TRAFFIC_POSTPAID_BY_HOUR</li>
	// <li>If you are not a bill-by-IP account beta user, the EIP billing is the same as that for the instance bound to the EIP. Therefore, you do not need to pass in this parameter.</li></ul>
	InternetChargeType *string `json:"InternetChargeType,omitempty" name:"InternetChargeType"`

	// The EIP outbound bandwidth cap, in Mbps.
	// <ul style="margin:0"><li>For bill-by-IP account beta users, valid values:<ul><li>BANDWIDTH_PACKAGE: 1 Mbps to 1000 Mbps</li>
	// <li>BANDWIDTH_POSTPAID_BY_HOUR: 1 Mbps to 100 Mbps</li>
	// <li>BANDWIDTH_PREPAID_BY_MONTH: 1 Mbps to 200 Mbps</li>
	// <li>TRAFFIC_POSTPAID_BY_HOUR: 1 Mbps to 100 Mbps</li></ul>Default value: 1 Mbps</li>
	// <li>If you are not a bill-by-IP account beta user, the EIP outbound bandwidth cap is subject to that of the instance bound to the EIP. Therefore, you do not need to pass in this parameter.</li></ul>
	InternetMaxBandwidthOut *int64 `json:"InternetMaxBandwidthOut,omitempty" name:"InternetMaxBandwidthOut"`

	// A required billing parameter for an EIP billed by monthly bandwidth subscription. For EIPs using other billing modes, it can be ignored.
	AddressChargePrepaid *AddressChargePrepaid `json:"AddressChargePrepaid,omitempty" name:"AddressChargePrepaid"`

	// The EIP type. Default: EIP.
	// <ul style="margin:0"><li>For a user who has activated the AIA allowlist, possible values are:<ul><li>AnycastEIP: an Anycast EIP address. For more information, see [Anycast Internet Acceleration](https://intl.cloud.tencent.com/document/product/644?from_cn_redirect=1).</li></ul>Note: Only certain regions support Anycast EIPs.</li></ul>
	AddressType *string `json:"AddressType,omitempty" name:"AddressType"`

	// Anycast publishing region
	// <ul style="margin:0"><li>Valid for users who have activated AIA. Values:<ul><li>ANYCAST_ZONE_GLOBAL: global publishing region </li><li>ANYCAST_ZONE_OVERSEAS: overseas publishing region</li><li><b>**[Disused]**</b> ANYCAST_ZONE_A: publishing region A (updated to ANYCAST_ZONE_GLOBAL)</li><li><b>**[Disused]**</b> ANYCAST_ZONE_B: publishing region B (updated to ANYCAST_ZONE_GLOBAL)</li></ul>Default: ANYCAST_ZONE_OVERSEAS.</li></ul>
	AnycastZone *string `json:"AnycastZone,omitempty" name:"AnycastZone"`

	// <b>**[Disused]**</b>
	// Whether the Anycast EIP can be bound to CLB instances.
	// <ul style="margin:0"><li>Valid for users who have activated the AIA. Values:<ul><li>TRUE: the Anycast EIP can be bound to CLB instances.</li>
	// <li>FALSE: the Anycast EIP can be bound to CVMs, NAT gateways, and HAVIPs.</li></ul>Default: FALSE.</li></ul>
	ApplicableForCLB *bool `json:"ApplicableForCLB,omitempty" name:"ApplicableForCLB"`

	// List of tags to be bound.
	Tags []*Tag `json:"Tags,omitempty" name:"Tags" list`

	// The unique ID of a BGP bandwidth package. If you configure this parameter and set InternetChargeType as BANDWIDTH_PACKAGE, the new EIP is added to this package and billed by the bandwidth package mode.
	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" name:"BandwidthPackageId"`
}

func (r *AllocateAddressesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AllocateAddressesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AllocateAddressesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// List of the unique IDs of the requested EIPs.
		AddressSet []*string `json:"AddressSet,omitempty" name:"AddressSet" list`

		// The Async task ID. You can use the [DescribeTaskResult](https://intl.cloud.tencent.com/document/api/215/36271?from_cn_redirect=1) API to query the task status.
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AllocateAddressesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AllocateAddressesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AssignIpv6AddressesRequest struct {
	*tchttp.BaseRequest

	// The `ID` of the ENI instance, such as `eni-m6dyj72l`.
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`

	// A list of `IPv6` addresses. You can specify a maximum of 10 at one time. The quota is calculated together with that of `Ipv6AddressCount`, a required input parameter alternative to this one.
	Ipv6Addresses []*Ipv6Address `json:"Ipv6Addresses,omitempty" name:"Ipv6Addresses" list`

	// The number of automatically assigned `IPv6` addresses. The total number of private IP addresses cannot exceed the quota. The quota is calculated together with that of `Ipv6Addresses`, a required input parameter alternative to this one.
	Ipv6AddressCount *uint64 `json:"Ipv6AddressCount,omitempty" name:"Ipv6AddressCount"`
}

func (r *AssignIpv6AddressesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AssignIpv6AddressesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AssignIpv6AddressesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The list of `IPv6` addresses assigned to ENIs.
		Ipv6AddressSet []*Ipv6Address `json:"Ipv6AddressSet,omitempty" name:"Ipv6AddressSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AssignIpv6AddressesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AssignIpv6AddressesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AssignIpv6CidrBlockRequest struct {
	*tchttp.BaseRequest

	// The `ID` of the `VPC`, such as `vpc-f49l6u0z`.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
}

func (r *AssignIpv6CidrBlockRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AssignIpv6CidrBlockRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AssignIpv6CidrBlockResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The assigned `IPv6` IP range, such as `3402:4e00:20:1000::/56`
		Ipv6CidrBlock *string `json:"Ipv6CidrBlock,omitempty" name:"Ipv6CidrBlock"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AssignIpv6CidrBlockResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AssignIpv6CidrBlockResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AssignIpv6SubnetCidrBlockRequest struct {
	*tchttp.BaseRequest

	// The `ID` of the VPC where the subnet is located, such as `vpc-f49l6u0z`.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// The assigned `IPv6` subnet IP range list.
	Ipv6SubnetCidrBlocks []*Ipv6SubnetCidrBlock `json:"Ipv6SubnetCidrBlocks,omitempty" name:"Ipv6SubnetCidrBlocks" list`
}

func (r *AssignIpv6SubnetCidrBlockRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AssignIpv6SubnetCidrBlockRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AssignIpv6SubnetCidrBlockResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The assigned `IPv6` subnet IP range list.
		Ipv6SubnetCidrBlockSet []*Ipv6SubnetCidrBlock `json:"Ipv6SubnetCidrBlockSet,omitempty" name:"Ipv6SubnetCidrBlockSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AssignIpv6SubnetCidrBlockResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AssignIpv6SubnetCidrBlockResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AssignPrivateIpAddressesRequest struct {
	*tchttp.BaseRequest

	// The ID of the ENI instance, such as `eni-m6dyj72l`.
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`

	// The information on private IP addresses, of which you can specify a maximum of 10 at a time. You should provide either this parameter or SecondaryPrivateIpAddressCount, or both.
	PrivateIpAddresses []*PrivateIpAddressSpecification `json:"PrivateIpAddresses,omitempty" name:"PrivateIpAddresses" list`

	// The number of newly-applied private IP addresses. You should provide either this parameter or PrivateIpAddresses, or both. The total number of private IP addresses cannot exceed the quota. For more information, see<a href="/document/product/576/18527">ENI Use Limits</a>.
	SecondaryPrivateIpAddressCount *uint64 `json:"SecondaryPrivateIpAddressCount,omitempty" name:"SecondaryPrivateIpAddressCount"`
}

func (r *AssignPrivateIpAddressesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AssignPrivateIpAddressesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AssignPrivateIpAddressesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The detailed information of the Private IP.
		PrivateIpAddressSet []*PrivateIpAddressSpecification `json:"PrivateIpAddressSet,omitempty" name:"PrivateIpAddressSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AssignPrivateIpAddressesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AssignPrivateIpAddressesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AssistantCidr struct {

	// The `ID` of a `VPC` instance, such as `vpc-6v2ht8q5`.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// The secondary CIDR, such as `172.16.0.0/16`.
	CidrBlock *string `json:"CidrBlock,omitempty" name:"CidrBlock"`

	// The secondary CIDR block type. 0: common secondary CIDR block. 1: container secondary CIDR block. Default: 0.
	AssistantType *int64 `json:"AssistantType,omitempty" name:"AssistantType"`

	// Subnets divided by the secondary CIDR.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SubnetSet []*Subnet `json:"SubnetSet,omitempty" name:"SubnetSet" list`
}

type AssociateAddressRequest struct {
	*tchttp.BaseRequest

	// The unique ID of the EIP, such as `eip-11112222`.
	AddressId *string `json:"AddressId,omitempty" name:"AddressId"`

	// The ID of the instance to be bound, such as `ins-11112222`. You can query the instance ID by logging into the [Console](https://console.cloud.tencent.com/cvm). You can also obtain the parameter value from the `InstanceId` field in the returned result of [DescribeInstances](https://intl.cloud.tencent.com/document/api/213/15728?from_cn_redirect=1) API.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// The ID of the ENI to be bonud, such as `eni-11112222`. `NetworkInterfaceId` and `InstanceId` cannot be specified at the same time. You can query the ENI ID by logging into the [Console](https://console.cloud.tencent.com/vpc/eni). You can also obtain the parameter value from the `networkInterfaceId` field in the returned result of [DescribeNetworkInterfaces](https://intl.cloud.tencent.com/document/api/215/15817?from_cn_redirect=1) API.
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`

	// The private IP to be bound. If you specify `NetworkInterfaceId`, then you must also specify `PrivateIpAddress`, indicating the EIP is bound to the specified private IP of the specified ENI. At the same time, you must ensure the specified `PrivateIpAddress` is a private IP on the `NetworkInterfaceId`. You can query the private IP of the specified ENI by logging into the [Console](https://console.cloud.tencent.com/vpc/eni). You can also obtain the parameter value from the `privateIpAddress` field in the returned result of [DescribeNetworkInterfaces](https://intl.cloud.tencent.com/document/api/215/15817?from_cn_redirect=1) API.
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" name:"PrivateIpAddress"`
}

func (r *AssociateAddressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AssociateAddressRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AssociateAddressResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The async task ID. You can use the [DescribeTaskResult](https://intl.cloud.tencent.com/document/api/215/36271?from_cn_redirect=1) API to query the task status.
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AssociateAddressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AssociateAddressResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AssociateDirectConnectGatewayNatGatewayRequest struct {
	*tchttp.BaseRequest

	// The direct connect gateway ID.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// The NAT Gateway ID.
	NatGatewayId *string `json:"NatGatewayId,omitempty" name:"NatGatewayId"`

	// The ID of the VPC instance, which can be obtained from the `VpcId` field in response of the `DescribeVpcs` API.
	DirectConnectGatewayId *string `json:"DirectConnectGatewayId,omitempty" name:"DirectConnectGatewayId"`
}

func (r *AssociateDirectConnectGatewayNatGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AssociateDirectConnectGatewayNatGatewayRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AssociateDirectConnectGatewayNatGatewayResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AssociateDirectConnectGatewayNatGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AssociateDirectConnectGatewayNatGatewayResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AssociateNatGatewayAddressRequest struct {
	*tchttp.BaseRequest

	// The ID of the NAT gateway, such as `nat-df45454`.
	NatGatewayId *string `json:"NatGatewayId,omitempty" name:"NatGatewayId"`

	// The number of EIPs you want to apply for. The system will create the same number of EIPs as you require. Either `AddressCount` or `PublicAddresses` must be passed in.
	AddressCount *uint64 `json:"AddressCount,omitempty" name:"AddressCount"`

	// Array of the EIPs bound to the NAT gateway. Either `AddressCount` or `PublicAddresses` must be passed in.
	PublicIpAddresses []*string `json:"PublicIpAddresses,omitempty" name:"PublicIpAddresses" list`

	// The availability zone of the EIP, which is passed in when the EIP is automatically assigned.
	Zone *string `json:"Zone,omitempty" name:"Zone"`
}

func (r *AssociateNatGatewayAddressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AssociateNatGatewayAddressRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AssociateNatGatewayAddressResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AssociateNatGatewayAddressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AssociateNatGatewayAddressResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AssociateNetworkAclSubnetsRequest struct {
	*tchttp.BaseRequest

	// Network ACL instance ID. Example: acl-12345678.
	NetworkAclId *string `json:"NetworkAclId,omitempty" name:"NetworkAclId"`

	// Array of subnet instance IDs. Example: [subnet-12345678]
	SubnetIds []*string `json:"SubnetIds,omitempty" name:"SubnetIds" list`
}

func (r *AssociateNetworkAclSubnetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AssociateNetworkAclSubnetsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AssociateNetworkAclSubnetsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AssociateNetworkAclSubnetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AssociateNetworkAclSubnetsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AssociateNetworkInterfaceSecurityGroupsRequest struct {
	*tchttp.BaseRequest

	// ENI instance ID, e.g. eni-pxir56ns. You can enter up to 100 instances for each request.
	NetworkInterfaceIds []*string `json:"NetworkInterfaceIds,omitempty" name:"NetworkInterfaceIds" list`

	// The security group instance ID, such as `sg-33ocnj9n`. It can be obtained through DescribeSecurityGroups. You can enter up to 100 instances for each request.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds" list`
}

func (r *AssociateNetworkInterfaceSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AssociateNetworkInterfaceSecurityGroupsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AssociateNetworkInterfaceSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AssociateNetworkInterfaceSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AssociateNetworkInterfaceSecurityGroupsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AttachCcnInstancesRequest struct {
	*tchttp.BaseRequest

	// The CCN instance ID, such as `ccn-f49l6u0z`.
	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`

	// List of associated network instances
	Instances []*CcnInstance `json:"Instances,omitempty" name:"Instances" list`

	// The UIN (root account) of the CCN. By default, the current account belongs to the UIN
	CcnUin *string `json:"CcnUin,omitempty" name:"CcnUin"`
}

func (r *AttachCcnInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AttachCcnInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AttachCcnInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AttachCcnInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AttachCcnInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AttachClassicLinkVpcRequest struct {
	*tchttp.BaseRequest

	// VPC instance ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// CVM Instance ID
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`
}

func (r *AttachClassicLinkVpcRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AttachClassicLinkVpcRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AttachClassicLinkVpcResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AttachClassicLinkVpcResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AttachClassicLinkVpcResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AttachNetworkInterfaceRequest struct {
	*tchttp.BaseRequest

	// The ID of the ENI instance, such as `eni-m6dyj72l`.
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`

	// The ID of the CVM instance, such as `ins-r8hr2upy`.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *AttachNetworkInterfaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AttachNetworkInterfaceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AttachNetworkInterfaceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AttachNetworkInterfaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AttachNetworkInterfaceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AuditCrossBorderComplianceRequest struct {
	*tchttp.BaseRequest

	// Service provider. Valid values: `UNICOM`.
	ServiceProvider *string `json:"ServiceProvider,omitempty" name:"ServiceProvider"`

	// Unique ID of compliance review request.
	ComplianceId *uint64 `json:"ComplianceId,omitempty" name:"ComplianceId"`

	// Audit behavior. Valid values: `APPROVED` and `DENY`.
	AuditBehavior *string `json:"AuditBehavior,omitempty" name:"AuditBehavior"`
}

func (r *AuditCrossBorderComplianceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AuditCrossBorderComplianceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AuditCrossBorderComplianceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AuditCrossBorderComplianceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AuditCrossBorderComplianceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BandwidthPackage struct {

	// The unique ID of the bandwidth package.
	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" name:"BandwidthPackageId"`

	// The bandwidth package type. Valid values: 'BGP', 'SINGLEISP', and 'ANYCAST'
	NetworkType *string `json:"NetworkType,omitempty" name:"NetworkType"`

	// The bandwidth package billing mode. Valid values: 'TOP5_POSTPAID_BY_MONTH' and 'PERCENT95_POSTPAID_BY_MONTH'
	ChargeType *string `json:"ChargeType,omitempty" name:"ChargeType"`

	// The name of the bandwidth package.
	BandwidthPackageName *string `json:"BandwidthPackageName,omitempty" name:"BandwidthPackageName"`

	// The creation time of the bandwidth package, which follows the `ISO8601` standard and uses `UTC` time in the format of `YYYY-MM-DDThh:mm:ssZ`.
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// The status of the bandwidth package. Valid values: 'CREATING', 'CREATED', 'DELETING', and 'DELETED'.
	Status *string `json:"Status,omitempty" name:"Status"`

	// The resource information of the bandwidth package.
	ResourceSet []*Resource `json:"ResourceSet,omitempty" name:"ResourceSet" list`

	// The limit of the bandwidth package in Mbps. The value '-1' indicates there is no limit.
	Bandwidth *int64 `json:"Bandwidth,omitempty" name:"Bandwidth"`
}

type BandwidthPackageBillBandwidth struct {

	// Current billable usage, in Mbps
	BandwidthUsage *uint64 `json:"BandwidthUsage,omitempty" name:"BandwidthUsage"`
}

type CCN struct {

	// The unique ID of the CCN
	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`

	// The name of the CCN
	CcnName *string `json:"CcnName,omitempty" name:"CcnName"`

	// The detailed information of the CCN
	CcnDescription *string `json:"CcnDescription,omitempty" name:"CcnDescription"`

	// The number of associated instances
	InstanceCount *uint64 `json:"InstanceCount,omitempty" name:"InstanceCount"`

	// The creation time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// The instance status. 'ISOLATED': Being isolated (instance is in arrears and service is suspended). 'AVAILABLE': Operating.
	State *string `json:"State,omitempty" name:"State"`

	// The instance service quality. ’PT’: Platinum , 'AU': Gold, 'AG': Silver.
	QosLevel *string `json:"QosLevel,omitempty" name:"QosLevel"`

	// The billing method. POSTPAID indicates postpaid.
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`

	// The limit type. INTER_REGION_LIMIT is the limit between regions. OUTER_REGION_LIMIT is a region egress limit.
	// Note: This field may return null, indicating no valid value.
	BandwidthLimitType *string `json:"BandwidthLimitType,omitempty" name:"BandwidthLimitType"`

	// Tag key-value pairs.
	TagSet []*Tag `json:"TagSet,omitempty" name:"TagSet" list`

	// Whether the CCN route priority feature is supported. Valid values: False: do not support; True: support.
	RoutePriorityFlag *bool `json:"RoutePriorityFlag,omitempty" name:"RoutePriorityFlag"`
}

type CcnAttachedInstance struct {

	// The ID of a CCN instance.
	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`

	// The type of associated instances:
	// <li>`VPC`: VPC</li>
	// <li>`DIRECTCONNECT`: Direct Connect</li>
	// <li>`BMVPC`: BM VPC</li>
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// The ID of the associated instance.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// The name of the associated instance.
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// The region to which the associated instance belongs, such as `ap-guangzhou`.
	InstanceRegion *string `json:"InstanceRegion,omitempty" name:"InstanceRegion"`

	// The UIN (root account) to which the associated instance belongs.
	InstanceUin *string `json:"InstanceUin,omitempty" name:"InstanceUin"`

	// The CIDR of the associated instance.
	CidrBlock []*string `json:"CidrBlock,omitempty" name:"CidrBlock" list`

	// The status of the associated instance:
	// <li>`PENDING`: In application</li>
	// <li>`ACTIVE`: Connected</li>
	// <li>`EXPIRED`: Expired</li>
	// <li>`REJECTED`: Rejected</li>
	// <li>`DELETED`: Deleted</li>
	// <li>`FAILED`: Failed (it will be asynchronously unbound after 2 hours)</li>
	// <li>`ATTACHING`: binding</li>
	// <li>`DETACHING`: Unbinding</li>
	// <li>`DETACHFAILED`: The unbinding failed (it will be asynchronously unbound after 2 hours)</li>
	State *string `json:"State,omitempty" name:"State"`

	// Association Time.
	AttachedTime *string `json:"AttachedTime,omitempty" name:"AttachedTime"`

	// The UIN (root account) to which the CCN belongs.
	CcnUin *string `json:"CcnUin,omitempty" name:"CcnUin"`

	// General location of the associated instance, such as CHINA_MAINLAND.
	InstanceArea *string `json:"InstanceArea,omitempty" name:"InstanceArea"`
}

type CcnBandwidthInfo struct {

	// The CCN ID that the bandwidth belongs to.
	// Note: this field may return null, indicating that no valid value was found.
	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`

	// The creation time of the instance.
	// Note: this field may return null, indicating that no valid value was found.
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// The expiration time of the instance.
	// Note: this field may return null, indicating that no valid value was found.
	ExpiredTime *string `json:"ExpiredTime,omitempty" name:"ExpiredTime"`

	// The unique ID of the bandwidth instance.
	// Note: this field may return null, indicating that no valid value was found.
	RegionFlowControlId *string `json:"RegionFlowControlId,omitempty" name:"RegionFlowControlId"`

	// The billing flag.
	// Note: this field may return null, indicating that no valid value was found.
	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`

	// The information of the bandwidth regions and bandwidth caps. The parameter is only returned for the cross-region limit mode, but not for egress limit.
	// Note: this field may return null, indicating that no valid value was found.
	CcnRegionBandwidthLimit *CcnRegionBandwidthLimit `json:"CcnRegionBandwidthLimit,omitempty" name:"CcnRegionBandwidthLimit"`
}

type CcnInstance struct {

	// The ID of the associated instance.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// The region to which the associated instance ID belongs, such as `ap-guangzhou`.
	InstanceRegion *string `json:"InstanceRegion,omitempty" name:"InstanceRegion"`

	// The type of the associated instance. Available values are:
	// <li>`VPC`: VPC</li>
	// <li>`DIRECTCONNECT`: Direct Connect</li>
	// <li>`BMVPC`: BM VPC</li>
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
}

type CcnRegionBandwidthLimit struct {

	// Region, such as `ap-guangzhou`
	Region *string `json:"Region,omitempty" name:"Region"`

	// The outbound bandwidth cap. Units: Mbps
	BandwidthLimit *uint64 `json:"BandwidthLimit,omitempty" name:"BandwidthLimit"`

	// Whether it is a BM region. The default is `false`.
	IsBm *bool `json:"IsBm,omitempty" name:"IsBm"`

	// The target region, such as `ap-shanghai`
	// Note: This field may return null, indicating no valid value.
	DstRegion *string `json:"DstRegion,omitempty" name:"DstRegion"`

	// Whether the target region is a BM region. The default is `false`.
	DstIsBm *bool `json:"DstIsBm,omitempty" name:"DstIsBm"`
}

type CcnRoute struct {

	// The ID of the routing policy
	RouteId *string `json:"RouteId,omitempty" name:"RouteId"`

	// Destination
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" name:"DestinationCidrBlock"`

	// The type of the next hop (associated instance type). Available types: VPC, DIRECTCONNECT
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// The next hop (associated instance)
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// The name of the next hop (associated instance name)
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// The region of the next hop (the region of the associated instance)
	InstanceRegion *string `json:"InstanceRegion,omitempty" name:"InstanceRegion"`

	// Update Time
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// Whether the route is enabled
	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`

	// The UIN (root account) to which the associated instance belongs
	InstanceUin *string `json:"InstanceUin,omitempty" name:"InstanceUin"`

	// Additional status of the route
	ExtraState *string `json:"ExtraState,omitempty" name:"ExtraState"`

	// Whether it is a dynamic route
	IsBgp *bool `json:"IsBgp,omitempty" name:"IsBgp"`

	// Route priority
	RoutePriority *uint64 `json:"RoutePriority,omitempty" name:"RoutePriority"`

	// Next hop port name (associated instance’s port name)
	InstanceExtraName *string `json:"InstanceExtraName,omitempty" name:"InstanceExtraName"`
}

type CheckAssistantCidrRequest struct {
	*tchttp.BaseRequest

	// `VPC` instance `ID`, e.g. `vpc-6v2ht8q5`.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Load CIDR blocks to add. CIDR block set; format: e.g. ["10.0.0.0/16", "172.16.0.0/16"]
	NewCidrBlocks []*string `json:"NewCidrBlocks,omitempty" name:"NewCidrBlocks" list`

	// Load CIDR blocks to delete. CIDR block set; Format: e.g. ["10.0.0.0/16", "172.16.0.0/16"]
	OldCidrBlocks []*string `json:"OldCidrBlocks,omitempty" name:"OldCidrBlocks" list`
}

func (r *CheckAssistantCidrRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CheckAssistantCidrRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CheckAssistantCidrResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Array of conflict resources.
		ConflictSourceSet []*ConflictSource `json:"ConflictSourceSet,omitempty" name:"ConflictSourceSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckAssistantCidrResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CheckAssistantCidrResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CheckNetDetectStateRequest struct {
	*tchttp.BaseRequest

	// The array of detection destination IPv4 addresses, which contains at most two IP addresses.
	DetectDestinationIp []*string `json:"DetectDestinationIp,omitempty" name:"DetectDestinationIp" list`

	// The type of the next hop. Currently supported types are:
	// VPN: VPN gateway;
	// DIRECTCONNECT: direct connect gateway;
	// PEERCONNECTION: peering connection;
	// NAT: NAT gateway;
	// NORMAL_CVM: normal CVM.
	NextHopType *string `json:"NextHopType,omitempty" name:"NextHopType"`

	// The next-hop destination gateway. The value is related to NextHopType.
	// If NextHopType is set to VPN, the value of this parameter is the VPN gateway ID, such as vpngw-12345678.
	// If NextHopType is set to DIRECTCONNECT, the value of this parameter is the direct connect gateway ID, such as dcg-12345678.
	// If NextHopType is set to PEERCONNECTION, the value of this parameter is the peering connection ID, such as pcx-12345678.
	// If NextHopType is set to NAT, the value of this parameter is the NAT gateway ID, such as nat-12345678.
	// If NextHopType is set to NORMAL_CVM, the value of this parameter is the IPv4 address of the CVM, such as 10.0.0.12.
	NextHopDestination *string `json:"NextHopDestination,omitempty" name:"NextHopDestination"`

	// ID of a network inspector instance, e.g. netd-12345678. Enter at least one of this parameter, VpcId, SubnetId, and NetDetectName. Use NetDetectId if it is present.
	NetDetectId *string `json:"NetDetectId,omitempty" name:"NetDetectId"`

	// ID of a `VPC` instance, e.g. `vpc-12345678`, which is used together with SubnetId and NetDetectName. You should enter either this parameter or NetDetectId, or both. Use NetDetectId if it is present.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// ID of a subnet instance, e.g. `subnet-12345678`, which is used together with VpcId and NetDetectName. You should enter either this parameter or NetDetectId, or both. Use NetDetectId if it is present.
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// The name of a network inspector, up to 60 bytes in length. It is used together with VpcId and NetDetectName. You should enter either this parameter or NetDetectId, or both. Use NetDetectId if it is present.
	NetDetectName *string `json:"NetDetectName,omitempty" name:"NetDetectName"`
}

func (r *CheckNetDetectStateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CheckNetDetectStateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CheckNetDetectStateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The array of network detection verification results.
		NetDetectIpStateSet []*NetDetectIpState `json:"NetDetectIpStateSet,omitempty" name:"NetDetectIpStateSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckNetDetectStateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CheckNetDetectStateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CidrForCcn struct {

	// Local CIDR block, including subnet CIDR block and secondary CIDR block
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Cidr *string `json:"Cidr,omitempty" name:"Cidr"`

	// Whether the routing policy of the VPC subnet is published to CCN.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	PublishedToVbc *bool `json:"PublishedToVbc,omitempty" name:"PublishedToVbc"`
}

type ClassicLinkInstance struct {

	// VPC instance ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// The unique ID of the CVM instance
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type CloneSecurityGroupRequest struct {
	*tchttp.BaseRequest

	// ID of the security group to be cloned, such as `sg-33ocnj9n`. This can be obtained through the `DescribeSecurityGroups` API.
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`

	// The name of security group clone. You can enter any name within 60 characters. If this parameter is left empty, the security group clone will use the name of the source security group.
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// Description of the security group clone. You can enter up to 100 characters. If this parameter is left empty, the security group clone will use the description of the source security group.
	GroupDescription *string `json:"GroupDescription,omitempty" name:"GroupDescription"`

	// Project ID of the security group clone. The default is 0. You can query it on the project management page of the Tencent Cloud console.
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// The region of the source security group for a cross-region clone. For example, to clone the security group in Guangzhou to Shanghai, set it to `ap-guangzhou`.
	RemoteRegion *string `json:"RemoteRegion,omitempty" name:"RemoteRegion"`
}

func (r *CloneSecurityGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CloneSecurityGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CloneSecurityGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Security group object
	// Note: this field may return `null`, indicating that no valid values can be obtained.
		SecurityGroup *SecurityGroup `json:"SecurityGroup,omitempty" name:"SecurityGroup"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CloneSecurityGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CloneSecurityGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ConflictItem struct {

	// Conflict resource ID
	ConfilctId *string `json:"ConfilctId,omitempty" name:"ConfilctId"`

	// Conflict destination resource
	DestinationItem *string `json:"DestinationItem,omitempty" name:"DestinationItem"`
}

type ConflictSource struct {

	// Conflict resource ID
	ConflictSourceId *string `json:"ConflictSourceId,omitempty" name:"ConflictSourceId"`

	// Conflict resource
	SourceItem *string `json:"SourceItem,omitempty" name:"SourceItem"`

	// Conflict resource items
	ConflictItemSet []*ConflictItem `json:"ConflictItemSet,omitempty" name:"ConflictItemSet" list`
}

type CreateAddressTemplateGroupRequest struct {
	*tchttp.BaseRequest

	// The name of the IP address template group.
	AddressTemplateGroupName *string `json:"AddressTemplateGroupName,omitempty" name:"AddressTemplateGroupName"`

	// The instance ID of the IP address template, such as `ipm-mdunqeb6`.
	AddressTemplateIds []*string `json:"AddressTemplateIds,omitempty" name:"AddressTemplateIds" list`
}

func (r *CreateAddressTemplateGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAddressTemplateGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateAddressTemplateGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Group object of the IP address template.
		AddressTemplateGroup *AddressTemplateGroup `json:"AddressTemplateGroup,omitempty" name:"AddressTemplateGroup"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAddressTemplateGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAddressTemplateGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateAddressTemplateRequest struct {
	*tchttp.BaseRequest

	// The name of the IP address template
	AddressTemplateName *string `json:"AddressTemplateName,omitempty" name:"AddressTemplateName"`

	// Address information, including IP, CIDR and IP address range.
	Addresses []*string `json:"Addresses,omitempty" name:"Addresses" list`
}

func (r *CreateAddressTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAddressTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateAddressTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The template object of the IP address.
		AddressTemplate *AddressTemplate `json:"AddressTemplate,omitempty" name:"AddressTemplate"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAddressTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAddressTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateAndAttachNetworkInterfaceRequest struct {
	*tchttp.BaseRequest

	// The ID of the VPC instance. You can obtain the parameter value from the `VpcId` field in the returned result of the `DescribeVpcs` API.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// The name of the ENI. The maximum length is 60 bytes.
	NetworkInterfaceName *string `json:"NetworkInterfaceName,omitempty" name:"NetworkInterfaceName"`

	// The subnet instance ID of the ENI, such as 'subnet-0ap8nwca'.
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// CVM instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// The information of the specified private IPs. You can specify a maximum of 10 IPs each time.
	PrivateIpAddresses []*PrivateIpAddressSpecification `json:"PrivateIpAddresses,omitempty" name:"PrivateIpAddresses" list`

	// The number of private IP addresses you can apply for. The total number of private IP addresses cannot exceed the quota.
	SecondaryPrivateIpAddressCount *uint64 `json:"SecondaryPrivateIpAddressCount,omitempty" name:"SecondaryPrivateIpAddressCount"`

	// The security group to be bound with, such as ['sg-1dd51d'].
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds" list`

	// The ENI description. You can enter any information within 60 characters.
	NetworkInterfaceDescription *string `json:"NetworkInterfaceDescription,omitempty" name:"NetworkInterfaceDescription"`

	// Bound tags, such as [{"Key": "city", "Value": "shanghai"}].
	Tags []*Tag `json:"Tags,omitempty" name:"Tags" list`
}

func (r *CreateAndAttachNetworkInterfaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAndAttachNetworkInterfaceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateAndAttachNetworkInterfaceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The ENI instance.
		NetworkInterface *NetworkInterface `json:"NetworkInterface,omitempty" name:"NetworkInterface"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAndAttachNetworkInterfaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAndAttachNetworkInterfaceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateAssistantCidrRequest struct {
	*tchttp.BaseRequest

	// `VPC` instance `ID`, e.g. `vpc-6v2ht8q5`.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// CIDR set, e.g. ["10.0.0.0/16", "172.16.0.0/16"]
	CidrBlocks []*string `json:"CidrBlocks,omitempty" name:"CidrBlocks" list`
}

func (r *CreateAssistantCidrRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAssistantCidrRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateAssistantCidrResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// A set of secondary CIDR blocks.
	// Note: This field may return null, indicating that no valid value was found.
		AssistantCidrSet []*AssistantCidr `json:"AssistantCidrSet,omitempty" name:"AssistantCidrSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAssistantCidrResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAssistantCidrResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateBandwidthPackageRequest struct {
	*tchttp.BaseRequest

	// The type of the bandwidth package. Valid values: `HIGH_QUALITY_BGP`, `BGP`, `SINGLEISP`, and `ANYCAST`.
	NetworkType *string `json:"NetworkType,omitempty" name:"NetworkType"`

	// The bandwidth package billing mode. Valid values: 'TOP5_POSTPAID_BY_MONTH' and 'PERCENT95_POSTPAID_BY_MONTH'.
	ChargeType *string `json:"ChargeType,omitempty" name:"ChargeType"`

	// The name of the bandwidth package.
	BandwidthPackageName *string `json:"BandwidthPackageName,omitempty" name:"BandwidthPackageName"`

	// The number of bandwidth packages (It can only be “1” for bill-by-CVM accounts)
	BandwidthPackageCount *uint64 `json:"BandwidthPackageCount,omitempty" name:"BandwidthPackageCount"`

	// The limit of the bandwidth package in Mbps. The value '-1' indicates there is no limit. This feature is currently in beta.
	InternetMaxBandwidth *int64 `json:"InternetMaxBandwidth,omitempty" name:"InternetMaxBandwidth"`

	// The list of tags to be bound.
	Tags []*Tag `json:"Tags,omitempty" name:"Tags" list`

	// The protocol type of the bandwidth package. Valid values: 'ipv4' and 'ipv6'. Default value: 'ipv4'.
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
}

func (r *CreateBandwidthPackageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateBandwidthPackageRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateBandwidthPackageResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique ID of the bandwidth package.
		BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" name:"BandwidthPackageId"`

		// The unique ID list of the bandwidth package (effective only when you apply for more than 1 bandwidth packages).
		BandwidthPackageIds []*string `json:"BandwidthPackageIds,omitempty" name:"BandwidthPackageIds" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateBandwidthPackageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateBandwidthPackageResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateCcnRequest struct {
	*tchttp.BaseRequest

	// The name of the CCN. The maximum length is 60 characters.
	CcnName *string `json:"CcnName,omitempty" name:"CcnName"`

	// The description of the CCN. The maximum length is 100 characters.
	CcnDescription *string `json:"CcnDescription,omitempty" name:"CcnDescription"`

	// CCN service quality, 'PT': Platinum, 'AU': Gold, 'AG': Silver. The default is 'AU'.
	QosLevel *string `json:"QosLevel,omitempty" name:"QosLevel"`

	// The billing method. POSTPAID: postpaid by traffic. Default: POSTPAID.
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`

	// The bandwidth limit type. Valid values: OUTER_REGION_LIMIT: region outbound bandwidth limit; INTER_REGION_LIMIT: inter-region bandwidth limit. Default value: OUTER_REGION_LIMIT. Monthly-subscribed CCN instances only support inter-region bandwidth limit, while pay-as-you-go CCN instances support the both bandwidth limit types.
	BandwidthLimitType *string `json:"BandwidthLimitType,omitempty" name:"BandwidthLimitType"`

	// Bound tags, such as [{"Key": "city", "Value": "shanghai"}].
	Tags []*Tag `json:"Tags,omitempty" name:"Tags" list`
}

func (r *CreateCcnRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateCcnRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateCcnResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The CCN object.
		Ccn *CCN `json:"Ccn,omitempty" name:"Ccn"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCcnResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateCcnResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateCustomerGatewayRequest struct {
	*tchttp.BaseRequest

	// Customer gateway can be named freely, but the maximum length is 60 characters.
	CustomerGatewayName *string `json:"CustomerGatewayName,omitempty" name:"CustomerGatewayName"`

	// Customer gateway public IP.
	IpAddress *string `json:"IpAddress,omitempty" name:"IpAddress"`

	// Bound tags, such as [{"Key": "city", "Value": "shanghai"}].
	Tags []*Tag `json:"Tags,omitempty" name:"Tags" list`
}

func (r *CreateCustomerGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateCustomerGatewayRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateCustomerGatewayResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Customer gateway object
		CustomerGateway *CustomerGateway `json:"CustomerGateway,omitempty" name:"CustomerGateway"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCustomerGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateCustomerGatewayResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateDefaultVpcRequest struct {
	*tchttp.BaseRequest

	// The ID of the availability zone in which the subnet resides. This parameter can be obtained through the [`DescribeZones`](https://intl.cloud.tencent.com/document/product/213/15707?from_cn_redirect=1) API, such as `ap-guangzhou-1`. If it’s not specified, a random availability zone will be used.
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// Whether to forcibly return a default VPC
	Force *bool `json:"Force,omitempty" name:"Force"`
}

func (r *CreateDefaultVpcRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateDefaultVpcRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateDefaultVpcResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Default VPC and subnet IDs
		Vpc *DefaultVpcSubnet `json:"Vpc,omitempty" name:"Vpc"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDefaultVpcResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateDefaultVpcResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateDirectConnectGatewayCcnRoutesRequest struct {
	*tchttp.BaseRequest

	// The ID of the Direct Connect gateway, such as `dcg-prpqlmg1`
	DirectConnectGatewayId *string `json:"DirectConnectGatewayId,omitempty" name:"DirectConnectGatewayId"`

	// The list of IDC IP range that must be connected
	Routes []*DirectConnectGatewayCcnRoute `json:"Routes,omitempty" name:"Routes" list`
}

func (r *CreateDirectConnectGatewayCcnRoutesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateDirectConnectGatewayCcnRoutesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateDirectConnectGatewayCcnRoutesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDirectConnectGatewayCcnRoutesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateDirectConnectGatewayCcnRoutesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateDirectConnectGatewayRequest struct {
	*tchttp.BaseRequest

	// The name of the direct connect gateway.
	DirectConnectGatewayName *string `json:"DirectConnectGatewayName,omitempty" name:"DirectConnectGatewayName"`

	// The type of the associated network. Valid values:
	// <li>VPC</li>
	// <li>CCN</li>
	NetworkType *string `json:"NetworkType,omitempty" name:"NetworkType"`

	// <li>When the NetworkType is VPC, this value is the VPC instance ID</li>
	// <li>When the NetworkType is CCN, this value is the CCN instance ID</li>
	NetworkInstanceId *string `json:"NetworkInstanceId,omitempty" name:"NetworkInstanceId"`

	// The type of the gateway. Valid values:
	// <li>NORMAL - (Default) Standard type. Note: CCN only supports the standard type</li>
	// <li>NAT - NAT type</li>NAT gateway supports network address translation. The specified type cannot be modified. A VPC can create one NAT direct connect gateway and one non-NAT direct connect gateway
	GatewayType *string `json:"GatewayType,omitempty" name:"GatewayType"`

	// CCN route publishing method. Valid values: `standard` and `exquisite`. This parameter is only valid for the CCN direct connect gateway.
	ModeType *string `json:"ModeType,omitempty" name:"ModeType"`
}

func (r *CreateDirectConnectGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateDirectConnectGatewayRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateDirectConnectGatewayResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The object of the direct connect gateway.
		DirectConnectGateway *DirectConnectGateway `json:"DirectConnectGateway,omitempty" name:"DirectConnectGateway"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDirectConnectGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateDirectConnectGatewayResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateFlowLogRequest struct {
	*tchttp.BaseRequest

	// ID of the VPC instance
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// The name of the flow log instance.
	FlowLogName *string `json:"FlowLogName,omitempty" name:"FlowLogName"`

	// The type of resources to which the flow log belongs. Valid values: 'VPC', 'SUBNET' and 'NETWORKINTERFACE'.
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`

	// The unique ID of the resource.
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// Type of the flow logs to be collected. Valid values: `ACCEPT`, `REJECT` and `ALL`.
	TrafficType *string `json:"TrafficType,omitempty" name:"TrafficType"`

	// The storage ID of the flow log.
	CloudLogId *string `json:"CloudLogId,omitempty" name:"CloudLogId"`

	// The description of the flow log instance
	FlowLogDescription *string `json:"FlowLogDescription,omitempty" name:"FlowLogDescription"`

	// Bound tags, such as [{"Key": "city", "Value": "shanghai"}]
	Tags []*Tag `json:"Tags,omitempty" name:"Tags" list`
}

func (r *CreateFlowLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateFlowLogRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateFlowLogResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The information of the flow log created.
		FlowLog []*FlowLog `json:"FlowLog,omitempty" name:"FlowLog" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateFlowLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateFlowLogResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateHaVipRequest struct {
	*tchttp.BaseRequest

	// The `ID` of the VPC to which the `HAVIP` belongs.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// The `ID` of the subnet to which the `HAVIP` belongs.
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// The name of the `HAVIP`.
	HaVipName *string `json:"HaVipName,omitempty" name:"HaVipName"`

	// The specified virtual IP address, which must be within the IP range of the `VPC` and not in use. It will be automatically assigned if not specified.
	Vip *string `json:"Vip,omitempty" name:"Vip"`
}

func (r *CreateHaVipRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateHaVipRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateHaVipResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// `HAVIP` object.
		HaVip *HaVip `json:"HaVip,omitempty" name:"HaVip"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateHaVipResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateHaVipResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateNatGatewayDestinationIpPortTranslationNatRuleRequest struct {
	*tchttp.BaseRequest

	// The ID of the NAT gateway, such as `nat-df45454`.
	NatGatewayId *string `json:"NatGatewayId,omitempty" name:"NatGatewayId"`

	// The port forwarding rules of the NAT gateway.
	DestinationIpPortTranslationNatRules []*DestinationIpPortTranslationNatRule `json:"DestinationIpPortTranslationNatRules,omitempty" name:"DestinationIpPortTranslationNatRules" list`
}

func (r *CreateNatGatewayDestinationIpPortTranslationNatRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateNatGatewayDestinationIpPortTranslationNatRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateNatGatewayDestinationIpPortTranslationNatRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateNatGatewayDestinationIpPortTranslationNatRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateNatGatewayDestinationIpPortTranslationNatRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateNatGatewayRequest struct {
	*tchttp.BaseRequest

	// NAT gateway name
	NatGatewayName *string `json:"NatGatewayName,omitempty" name:"NatGatewayName"`

	// The ID of the VPC instance. You can obtain the parameter value from the VpcId field in the returned result of DescribeVpcs API.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// The maximum outbound bandwidth of the NAT gateway (unit: Mbps). Supported parameter values: `20, 50, 100, 200, 500, 1000, 2000, 5000`. Default: `100Mbps`.
	InternetMaxBandwidthOut *uint64 `json:"InternetMaxBandwidthOut,omitempty" name:"InternetMaxBandwidthOut"`

	// The concurrent connection cap of the NAT gateway. Supported parameter values: `1000000, 3000000, 10000000`. The default value is `100000`.
	MaxConcurrentConnection *uint64 `json:"MaxConcurrentConnection,omitempty" name:"MaxConcurrentConnection"`

	// The number of EIPs that needs to be applied for. The system will create N number of EIPs according to your requirements. Either AddressCount or PublicAddresses must be passed in.
	AddressCount *uint64 `json:"AddressCount,omitempty" name:"AddressCount"`

	// The EIP array bound to the NAT gateway. Either AddressCount or PublicAddresses must be passed in.
	PublicIpAddresses []*string `json:"PublicIpAddresses,omitempty" name:"PublicIpAddresses" list`

	// The availability zone, such as `ap-guangzhou-1`.
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// Bound tags, such as [{"Key": "city", "Value": "shanghai"}].
	Tags []*Tag `json:"Tags,omitempty" name:"Tags" list`

	// Subnet of the NAT gateway
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
}

func (r *CreateNatGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateNatGatewayRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateNatGatewayResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// NAT gateway object array.
		NatGatewaySet []*NatGateway `json:"NatGatewaySet,omitempty" name:"NatGatewaySet" list`

		// The number of NAT gateway objects meeting the conditions.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateNatGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateNatGatewayResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateNatGatewaySourceIpTranslationNatRuleRequest struct {
	*tchttp.BaseRequest

	// The ID of the NAT Gateway, such as `nat-df45454`
	NatGatewayId *string `json:"NatGatewayId,omitempty" name:"NatGatewayId"`

	// The SNAT forwarding rule of the NAT Gateway
	SourceIpTranslationNatRules []*SourceIpTranslationNatRule `json:"SourceIpTranslationNatRules,omitempty" name:"SourceIpTranslationNatRules" list`
}

func (r *CreateNatGatewaySourceIpTranslationNatRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateNatGatewaySourceIpTranslationNatRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateNatGatewaySourceIpTranslationNatRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateNatGatewaySourceIpTranslationNatRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateNatGatewaySourceIpTranslationNatRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateNetDetectRequest struct {
	*tchttp.BaseRequest

	// The `ID` of a `VPC` instance, such as `vpc-12345678`.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// The ID of a subnet instance, such as subnet-12345678.
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// The name of a network detection instance. The maximum length is 60 characters.
	NetDetectName *string `json:"NetDetectName,omitempty" name:"NetDetectName"`

	// The array of detection destination IPv4 addresses, which contains at most two IP addresses.
	DetectDestinationIp []*string `json:"DetectDestinationIp,omitempty" name:"DetectDestinationIp" list`

	// The type of the next hop. Currently supported types are:
	// VPN: VPN gateway;
	// DIRECTCONNECT: direct connect gateway;
	// PEERCONNECTION: peering connection;
	// NAT: NAT gateway;
	// NORMAL_CVM: normal CVM.
	NextHopType *string `json:"NextHopType,omitempty" name:"NextHopType"`

	// The next-hop destination gateway. The value is related to NextHopType.
	// If NextHopType is set to VPN, the value of this parameter is the VPN gateway ID, such as vpngw-12345678.
	// If NextHopType is set to DIRECTCONNECT, the value of this parameter is the direct connect gateway ID, such as dcg-12345678.
	// If NextHopType is set to PEERCONNECTION, the value of this parameter is the peering connection ID, such as pcx-12345678.
	// If NextHopType is set to NAT, the value of this parameter is the NAT gateway ID, such as nat-12345678.
	// If NextHopType is set to NORMAL_CVM, the value of this parameter is the IPv4 address of the CVM, such as 10.0.0.12.
	NextHopDestination *string `json:"NextHopDestination,omitempty" name:"NextHopDestination"`

	// Network detection description.
	NetDetectDescription *string `json:"NetDetectDescription,omitempty" name:"NetDetectDescription"`
}

func (r *CreateNetDetectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateNetDetectRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateNetDetectResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The network detection (NetDetect) object.
		NetDetect *NetDetect `json:"NetDetect,omitempty" name:"NetDetect"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateNetDetectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateNetDetectResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateNetworkAclRequest struct {
	*tchttp.BaseRequest

	// ID of the VPC instance. You can obtain the parameter value from the VpcId field in the returned result of the DescribeVpcs API.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Name of the network ACL. The maximum length is 60 bytes.
	NetworkAclName *string `json:"NetworkAclName,omitempty" name:"NetworkAclName"`
}

func (r *CreateNetworkAclRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateNetworkAclRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateNetworkAclResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Network ACL instance
		NetworkAcl *NetworkAcl `json:"NetworkAcl,omitempty" name:"NetworkAcl"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateNetworkAclResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateNetworkAclResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateNetworkInterfaceRequest struct {
	*tchttp.BaseRequest

	// The ID of the VPC instance. You can obtain the parameter value from the VpcId field in the returned result of DescribeVpcs API.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// The name of the ENI. The maximum length is 60 characters.
	NetworkInterfaceName *string `json:"NetworkInterfaceName,omitempty" name:"NetworkInterfaceName"`

	// The subnet instance ID of the ENI, such as `subnet-0ap8nwca`.
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// ENI description can be named freely, but the maximum length is 60 characters.
	NetworkInterfaceDescription *string `json:"NetworkInterfaceDescription,omitempty" name:"NetworkInterfaceDescription"`

	// The number of private IP addresses that is newly applied for. The total number of private IP addresses cannot exceed the quota.
	SecondaryPrivateIpAddressCount *uint64 `json:"SecondaryPrivateIpAddressCount,omitempty" name:"SecondaryPrivateIpAddressCount"`

	// Specifies the security group to be bound with, such as ['sg-1dd51d'].
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds" list`

	// The information of the specified private IPs. You can specify a maximum of 10 each time.
	PrivateIpAddresses []*PrivateIpAddressSpecification `json:"PrivateIpAddresses,omitempty" name:"PrivateIpAddresses" list`

	// Bound tags, such as [{"Key": "city", "Value": "shanghai"}].
	Tags []*Tag `json:"Tags,omitempty" name:"Tags" list`
}

func (r *CreateNetworkInterfaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateNetworkInterfaceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateNetworkInterfaceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// ENI instance.
		NetworkInterface *NetworkInterface `json:"NetworkInterface,omitempty" name:"NetworkInterface"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateNetworkInterfaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateNetworkInterfaceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateRouteTableRequest struct {
	*tchttp.BaseRequest

	// The ID of the VPC instance to be operated on. You can obtain the parameter value from the VpcId field in the returned result of DescribeVpcs API.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// The route table name. The maximum length is 60 characters.
	RouteTableName *string `json:"RouteTableName,omitempty" name:"RouteTableName"`

	// Bound tags, such as [{"Key": "city", "Value": "shanghai"}].
	Tags []*Tag `json:"Tags,omitempty" name:"Tags" list`
}

func (r *CreateRouteTableRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateRouteTableRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateRouteTableResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Route table object.
		RouteTable *RouteTable `json:"RouteTable,omitempty" name:"RouteTable"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateRouteTableResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateRouteTableResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateRoutesRequest struct {
	*tchttp.BaseRequest

	// Route table instance ID.
	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`

	// Routing policy object.
	Routes []*Route `json:"Routes,omitempty" name:"Routes" list`
}

func (r *CreateRoutesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateRoutesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateRoutesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The number of newly added instances.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Route table object.
		RouteTableSet []*RouteTable `json:"RouteTableSet,omitempty" name:"RouteTableSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateRoutesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateRoutesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateSecurityGroupPoliciesRequest struct {
	*tchttp.BaseRequest

	// The security group instance ID, such as `sg-33ocnj9n`. This can be obtained through DescribeSecurityGroups.
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`

	// Security group policy set.
	SecurityGroupPolicySet *SecurityGroupPolicySet `json:"SecurityGroupPolicySet,omitempty" name:"SecurityGroupPolicySet"`
}

func (r *CreateSecurityGroupPoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateSecurityGroupPoliciesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateSecurityGroupPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSecurityGroupPoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateSecurityGroupPoliciesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateSecurityGroupRequest struct {
	*tchttp.BaseRequest

	// Security group can be named freely, but cannot exceed 60 characters.
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// The remarks for the security group. The maximum length is 100 characters.
	GroupDescription *string `json:"GroupDescription,omitempty" name:"GroupDescription"`

	// Project ID. The default is 0. You can query it on the project management page of the Qcloud console.
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// Bound tags, such as [{"Key": "city", "Value": "shanghai"}].
	Tags []*Tag `json:"Tags,omitempty" name:"Tags" list`
}

func (r *CreateSecurityGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateSecurityGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateSecurityGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Security group object.
		SecurityGroup *SecurityGroup `json:"SecurityGroup,omitempty" name:"SecurityGroup"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSecurityGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateSecurityGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateSecurityGroupWithPoliciesRequest struct {
	*tchttp.BaseRequest

	// Security group can be named freely, but cannot exceed 60 characters.
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// The remarks for the security group. The maximum length is 100 characters.
	GroupDescription *string `json:"GroupDescription,omitempty" name:"GroupDescription"`

	// The project id is 0 by default. You can query this in the project management page of the Qcloud console.
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// Security group policy set.
	SecurityGroupPolicySet *SecurityGroupPolicySet `json:"SecurityGroupPolicySet,omitempty" name:"SecurityGroupPolicySet"`
}

func (r *CreateSecurityGroupWithPoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateSecurityGroupWithPoliciesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateSecurityGroupWithPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Security group object.
		SecurityGroup *SecurityGroup `json:"SecurityGroup,omitempty" name:"SecurityGroup"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSecurityGroupWithPoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateSecurityGroupWithPoliciesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateServiceTemplateGroupRequest struct {
	*tchttp.BaseRequest

	// Group name of the protocol port template.
	ServiceTemplateGroupName *string `json:"ServiceTemplateGroupName,omitempty" name:"ServiceTemplateGroupName"`

	// Instance ID of the protocol port template, such as `ppm-4dw6agho`.
	ServiceTemplateIds []*string `json:"ServiceTemplateIds,omitempty" name:"ServiceTemplateIds" list`
}

func (r *CreateServiceTemplateGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateServiceTemplateGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateServiceTemplateGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Group object of the protocol port template.
		ServiceTemplateGroup *ServiceTemplateGroup `json:"ServiceTemplateGroup,omitempty" name:"ServiceTemplateGroup"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateServiceTemplateGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateServiceTemplateGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateServiceTemplateRequest struct {
	*tchttp.BaseRequest

	// Template name of the protocol port
	ServiceTemplateName *string `json:"ServiceTemplateName,omitempty" name:"ServiceTemplateName"`

	// It supports single port, multiple ports, consecutive ports and all ports. Supported protocols include TCP, UDP, ICMP, and GRE.
	Services []*string `json:"Services,omitempty" name:"Services" list`
}

func (r *CreateServiceTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateServiceTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateServiceTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Protocol port template object.
		ServiceTemplate *ServiceTemplate `json:"ServiceTemplate,omitempty" name:"ServiceTemplate"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateServiceTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateServiceTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateSubnetRequest struct {
	*tchttp.BaseRequest

	// The ID of the VPC instance to be operated on. You can obtain the parameter value from the VpcId field in the returned result of DescribeVpcs API.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// The subnet name. The maximum length is 60 bytes.
	SubnetName *string `json:"SubnetName,omitempty" name:"SubnetName"`

	// The subnet IP address range. It must be within the VPC IP address range. Subnet IP address ranges cannot overlap with each other within the same VPC.
	CidrBlock *string `json:"CidrBlock,omitempty" name:"CidrBlock"`

	// The ID of the availability zone in which the subnet resides. You can set up disaster recovery across availability zones by choosing different availability zones for different subnets.
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// Bound tags, such as [{"Key": "city", "Value": "shanghai"}].
	Tags []*Tag `json:"Tags,omitempty" name:"Tags" list`
}

func (r *CreateSubnetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateSubnetRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateSubnetResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Subnet object.
		Subnet *Subnet `json:"Subnet,omitempty" name:"Subnet"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSubnetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateSubnetResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateSubnetsRequest struct {
	*tchttp.BaseRequest

	// The `ID` of the `VPC` instance, such as `vpc-6v2ht8q5`.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// The subnet object list.
	Subnets []*SubnetInput `json:"Subnets,omitempty" name:"Subnets" list`

	// Bound tags. Note that the collection of tags here is shared by all subnet objects in the list. You cannot specify tags for each subnet. Example: [{"Key": "city", "Value": "shanghai"}].
	Tags []*Tag `json:"Tags,omitempty" name:"Tags" list`
}

func (r *CreateSubnetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateSubnetsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateSubnetsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The list of newly created subnets.
		SubnetSet []*Subnet `json:"SubnetSet,omitempty" name:"SubnetSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSubnetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateSubnetsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateVpcRequest struct {
	*tchttp.BaseRequest

	// The VPC name. The maximum length is 60 bytes.
	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`

	// VPC CIDR blocks, which must fall within the following three private network IP ranges: 10.0.0.0/16, 172.16.0.0/16 and 192.168.0.0/16.
	CidrBlock *string `json:"CidrBlock,omitempty" name:"CidrBlock"`

	// Whether multicast is enabled. `true`: Enabled. `false`: Not enabled.
	EnableMulticast *string `json:"EnableMulticast,omitempty" name:"EnableMulticast"`

	// DNS address. A maximum of 4 addresses is supported.
	DnsServers []*string `json:"DnsServers,omitempty" name:"DnsServers" list`

	// Domain name of DHCP
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// Bound tags, such as [{"Key": "city", "Value": "shanghai"}]
	Tags []*Tag `json:"Tags,omitempty" name:"Tags" list`
}

func (r *CreateVpcRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateVpcRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateVpcResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The VPC object.
		Vpc *Vpc `json:"Vpc,omitempty" name:"Vpc"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateVpcResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateVpcResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateVpnConnectionRequest struct {
	*tchttp.BaseRequest

	// VPC instance ID, which can be obtained from the `VpcId` field in the response of the [`DescribeVpcs`](https://intl.cloud.tencent.com/document/product/215/15778?from_cn_redirect=1) API.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// The ID of the VPN gateway instance.
	VpnGatewayId *string `json:"VpnGatewayId,omitempty" name:"VpnGatewayId"`

	// The ID of the customer gateway, such as `cgw-2wqq41m9`. You can query the customer gateway by using the `DescribeCustomerGateways` API.
	CustomerGatewayId *string `json:"CustomerGatewayId,omitempty" name:"CustomerGatewayId"`

	// Gateway can be named freely, but the maximum length is 60 characters.
	VpnConnectionName *string `json:"VpnConnectionName,omitempty" name:"VpnConnectionName"`

	// The pre-shared key.
	PreShareKey *string `json:"PreShareKey,omitempty" name:"PreShareKey"`

	// The SPD policy group, for example: {"10.0.0.5/24":["172.123.10.5/16"]}. 10.0.0.5/24 is the VPC internal IP range, and 172.123.10.5/16 is the IDC IP range. The user specifies the IP range in the VPC that can communicate with the IP range in the IDC.
	SecurityPolicyDatabases []*SecurityPolicyDatabase `json:"SecurityPolicyDatabases,omitempty" name:"SecurityPolicyDatabases" list`

	// Internet Key Exchange (IKE) configuration. IKE has a self-protection mechanism. The network security protocol is configured by the user.
	IKEOptionsSpecification *IKEOptionsSpecification `json:"IKEOptionsSpecification,omitempty" name:"IKEOptionsSpecification"`

	// IPSec configuration. The IPSec secure session configuration is provided by Tencent Cloud.
	IPSECOptionsSpecification *IPSECOptionsSpecification `json:"IPSECOptionsSpecification,omitempty" name:"IPSECOptionsSpecification"`

	// Bound tags, such as [{"Key": "city", "Value": "shanghai"}].
	Tags []*Tag `json:"Tags,omitempty" name:"Tags" list`

	// Whether the tunnel health check is supported.
	EnableHealthCheck *bool `json:"EnableHealthCheck,omitempty" name:"EnableHealthCheck"`

	// Local IP address for the health check
	HealthCheckLocalIp *string `json:"HealthCheckLocalIp,omitempty" name:"HealthCheckLocalIp"`

	// Peer IP address for the health check
	HealthCheckRemoteIp *string `json:"HealthCheckRemoteIp,omitempty" name:"HealthCheckRemoteIp"`
}

func (r *CreateVpnConnectionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateVpnConnectionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateVpnConnectionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Tunnel instance object.
		VpnConnection *VpnConnection `json:"VpnConnection,omitempty" name:"VpnConnection"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateVpnConnectionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateVpnConnectionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateVpnGatewayRequest struct {
	*tchttp.BaseRequest

	// VPC instance ID, which can be obtained from the `VpcId` field in the response of the [`DescribeVpcs`](https://intl.cloud.tencent.com/document/product/215/15778?from_cn_redirect=1) API.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// The VPN gateway name. The maximum length is 60 bytes.
	VpnGatewayName *string `json:"VpnGatewayName,omitempty" name:"VpnGatewayName"`

	// The public network bandwidth configuration. Available bandwidth specifications: 5, 10, 20, 50, and 100. Unit: Mbps
	InternetMaxBandwidthOut *uint64 `json:"InternetMaxBandwidthOut,omitempty" name:"InternetMaxBandwidthOut"`

	// The VPN gateway billing mode. PREPAID: prepaid means monthly subscription. POSTPAID_BY_HOUR: postpaid means pay-as-you-go. Default: POSTPAID_BY_HOUR. If prepaid mode is specified, the `InstanceChargePrepaid` parameter must be entered.
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`

	// Parameter settings for prepaid billing mode, also known as monthly subscription. This parameter can specify the purchase period and other attributes such as auto-renewal. This parameter is mandatory for prepaid instances.
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitempty" name:"InstanceChargePrepaid"`

	// The availability zone, such as `ap-guangzhou-2`.
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// VPN gateway type. Value: `CCN`, indicates CCN-type VPN gateway
	Type *string `json:"Type,omitempty" name:"Type"`

	// Bound tags, such as [{"Key": "city", "Value": "shanghai"}].
	Tags []*Tag `json:"Tags,omitempty" name:"Tags" list`
}

func (r *CreateVpnGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateVpnGatewayRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateVpnGatewayResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// VPN gateway object.
		VpnGateway *VpnGateway `json:"VpnGateway,omitempty" name:"VpnGateway"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateVpnGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateVpnGatewayResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CrossBorderCompliance struct {

	// Service provider. Valid values: `UNICOM`.
	ServiceProvider *string `json:"ServiceProvider,omitempty" name:"ServiceProvider"`

	// ID of compliance review request.
	ComplianceId *uint64 `json:"ComplianceId,omitempty" name:"ComplianceId"`

	// Full company name.
	Company *string `json:"Company,omitempty" name:"Company"`

	// Unified Social Credit Code.
	UniformSocialCreditCode *string `json:"UniformSocialCreditCode,omitempty" name:"UniformSocialCreditCode"`

	// Legal person.
	LegalPerson *string `json:"LegalPerson,omitempty" name:"LegalPerson"`

	// Issuing authority.
	IssuingAuthority *string `json:"IssuingAuthority,omitempty" name:"IssuingAuthority"`

	// Business license.
	BusinessLicense *string `json:"BusinessLicense,omitempty" name:"BusinessLicense"`

	// Business address.
	BusinessAddress *string `json:"BusinessAddress,omitempty" name:"BusinessAddress"`

	// Zip code.
	PostCode *uint64 `json:"PostCode,omitempty" name:"PostCode"`

	// Operator.
	Manager *string `json:"Manager,omitempty" name:"Manager"`

	// Operator ID card number.
	ManagerId *string `json:"ManagerId,omitempty" name:"ManagerId"`

	// Operator ID card.
	ManagerIdCard *string `json:"ManagerIdCard,omitempty" name:"ManagerIdCard"`

	// Operator address.
	ManagerAddress *string `json:"ManagerAddress,omitempty" name:"ManagerAddress"`

	// Operator phone number.
	ManagerTelephone *string `json:"ManagerTelephone,omitempty" name:"ManagerTelephone"`

	// Email.
	Email *string `json:"Email,omitempty" name:"Email"`

	// Service handling form.
	ServiceHandlingForm *string `json:"ServiceHandlingForm,omitempty" name:"ServiceHandlingForm"`

	// Authorization letter.
	AuthorizationLetter *string `json:"AuthorizationLetter,omitempty" name:"AuthorizationLetter"`

	// Information security commitment.
	SafetyCommitment *string `json:"SafetyCommitment,omitempty" name:"SafetyCommitment"`

	// Service start date.
	ServiceStartDate *string `json:"ServiceStartDate,omitempty" name:"ServiceStartDate"`

	// Service end date.
	ServiceEndDate *string `json:"ServiceEndDate,omitempty" name:"ServiceEndDate"`

	// Status. Valid values: `PENDING`, `APPROVED`, and `DENY`.
	State *string `json:"State,omitempty" name:"State"`

	// Creation time of the review form.
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
}

type CustomerGateway struct {

	// The unique ID of the customer gateway
	CustomerGatewayId *string `json:"CustomerGatewayId,omitempty" name:"CustomerGatewayId"`

	// Gateway Name
	CustomerGatewayName *string `json:"CustomerGatewayName,omitempty" name:"CustomerGatewayName"`

	// Public network address
	IpAddress *string `json:"IpAddress,omitempty" name:"IpAddress"`

	// The creation time.
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
}

type CustomerGatewayVendor struct {

	// Platform.
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// Software version.
	SoftwareVersion *string `json:"SoftwareVersion,omitempty" name:"SoftwareVersion"`

	// Vendor name.
	VendorName *string `json:"VendorName,omitempty" name:"VendorName"`
}

type CvmInstance struct {

	// VPC instance ID.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Subnet instance ID.
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// CVM instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// CVM Name
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// CVM status.
	InstanceState *string `json:"InstanceState,omitempty" name:"InstanceState"`

	// Number of CPU cores in an instance (in core).
	CPU *uint64 `json:"CPU,omitempty" name:"CPU"`

	// Instance’s memory capacity. Unit: GB.
	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`

	// The creation time.
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// Instance type.
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// Instance ENI quota (including primary ENIs).
	EniLimit *uint64 `json:"EniLimit,omitempty" name:"EniLimit"`

	// Private IP quoata for instance ENIs (including primary ENIs).
	EniIpLimit *uint64 `json:"EniIpLimit,omitempty" name:"EniIpLimit"`

	// The number of ENIs (including primary ENIs) bound to a instance.
	InstanceEniCount *uint64 `json:"InstanceEniCount,omitempty" name:"InstanceEniCount"`
}

type DefaultVpcSubnet struct {

	// Default VpcId
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Default SubnetId
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
}

type DeleteAddressTemplateGroupRequest struct {
	*tchttp.BaseRequest

	// The IP address template group instance ID, such as `ipmg-90cex8mq`.
	AddressTemplateGroupId *string `json:"AddressTemplateGroupId,omitempty" name:"AddressTemplateGroupId"`
}

func (r *DeleteAddressTemplateGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteAddressTemplateGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteAddressTemplateGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAddressTemplateGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteAddressTemplateGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteAddressTemplateRequest struct {
	*tchttp.BaseRequest

	// The IP address template instance ID, such as `ipm-09o5m8kc`.
	AddressTemplateId *string `json:"AddressTemplateId,omitempty" name:"AddressTemplateId"`
}

func (r *DeleteAddressTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteAddressTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteAddressTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAddressTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteAddressTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteAssistantCidrRequest struct {
	*tchttp.BaseRequest

	// `VPC` instance `ID`, e.g. `vpc-6v2ht8q5`.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// CIDR set, e.g. ["10.0.0.0/16", "172.16.0.0/16"]
	CidrBlocks []*string `json:"CidrBlocks,omitempty" name:"CidrBlocks" list`
}

func (r *DeleteAssistantCidrRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteAssistantCidrRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteAssistantCidrResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAssistantCidrResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteAssistantCidrResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteBandwidthPackageRequest struct {
	*tchttp.BaseRequest

	// The unique ID of the bandwidth package to be deleted.
	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" name:"BandwidthPackageId"`
}

func (r *DeleteBandwidthPackageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteBandwidthPackageRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteBandwidthPackageResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteBandwidthPackageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteBandwidthPackageResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteCcnRequest struct {
	*tchttp.BaseRequest

	// The CCN instance ID, such as `ccn-f49l6u0z`.
	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
}

func (r *DeleteCcnRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteCcnRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteCcnResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteCcnResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteCcnResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteCustomerGatewayRequest struct {
	*tchttp.BaseRequest

	// The ID of the customer gateway, such as `cgw-2wqq41m9`. You can query the customer gateway by using the `DescribeCustomerGateways` API.
	CustomerGatewayId *string `json:"CustomerGatewayId,omitempty" name:"CustomerGatewayId"`
}

func (r *DeleteCustomerGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteCustomerGatewayRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteCustomerGatewayResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteCustomerGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteCustomerGatewayResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteDirectConnectGatewayCcnRoutesRequest struct {
	*tchttp.BaseRequest

	// The ID of the Direct Connect gateway, such as `dcg-prpqlmg1`
	DirectConnectGatewayId *string `json:"DirectConnectGatewayId,omitempty" name:"DirectConnectGatewayId"`

	// The route ID, such as `ccnr-f49l6u0z`.
	RouteIds []*string `json:"RouteIds,omitempty" name:"RouteIds" list`
}

func (r *DeleteDirectConnectGatewayCcnRoutesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteDirectConnectGatewayCcnRoutesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteDirectConnectGatewayCcnRoutesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteDirectConnectGatewayCcnRoutesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteDirectConnectGatewayCcnRoutesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteDirectConnectGatewayRequest struct {
	*tchttp.BaseRequest

	// The unique `ID` of the direct connect gateway, such as `dcg-9o233uri`.
	DirectConnectGatewayId *string `json:"DirectConnectGatewayId,omitempty" name:"DirectConnectGatewayId"`
}

func (r *DeleteDirectConnectGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteDirectConnectGatewayRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteDirectConnectGatewayResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteDirectConnectGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteDirectConnectGatewayResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteFlowLogRequest struct {
	*tchttp.BaseRequest

	// ID of the VPC instance
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// The unique ID of the flow log.
	FlowLogId *string `json:"FlowLogId,omitempty" name:"FlowLogId"`
}

func (r *DeleteFlowLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteFlowLogRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteFlowLogResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteFlowLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteFlowLogResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteHaVipRequest struct {
	*tchttp.BaseRequest

	// The unique `ID` of the `HAVIP`, such as `havip-9o233uri`.
	HaVipId *string `json:"HaVipId,omitempty" name:"HaVipId"`
}

func (r *DeleteHaVipRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteHaVipRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteHaVipResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteHaVipResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteHaVipResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteNatGatewayDestinationIpPortTranslationNatRuleRequest struct {
	*tchttp.BaseRequest

	// The ID of the NAT gateway, such as `nat-df45454`.
	NatGatewayId *string `json:"NatGatewayId,omitempty" name:"NatGatewayId"`

	// The port forwarding rules of the NAT gateway.
	DestinationIpPortTranslationNatRules []*DestinationIpPortTranslationNatRule `json:"DestinationIpPortTranslationNatRules,omitempty" name:"DestinationIpPortTranslationNatRules" list`
}

func (r *DeleteNatGatewayDestinationIpPortTranslationNatRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteNatGatewayDestinationIpPortTranslationNatRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteNatGatewayDestinationIpPortTranslationNatRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteNatGatewayDestinationIpPortTranslationNatRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteNatGatewayDestinationIpPortTranslationNatRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteNatGatewayRequest struct {
	*tchttp.BaseRequest

	// The ID of the NAT gateway, such as `nat-df45454`.
	NatGatewayId *string `json:"NatGatewayId,omitempty" name:"NatGatewayId"`
}

func (r *DeleteNatGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteNatGatewayRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteNatGatewayResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteNatGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteNatGatewayResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteNatGatewaySourceIpTranslationNatRuleRequest struct {
	*tchttp.BaseRequest

	// The ID of the NAT Gateway, such as `nat-df45454`
	NatGatewayId *string `json:"NatGatewayId,omitempty" name:"NatGatewayId"`

	// The list of SNAT rule IDs of a NAT Gateway, such as `snat-df43254`
	NatGatewaySnatIds []*string `json:"NatGatewaySnatIds,omitempty" name:"NatGatewaySnatIds" list`
}

func (r *DeleteNatGatewaySourceIpTranslationNatRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteNatGatewaySourceIpTranslationNatRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteNatGatewaySourceIpTranslationNatRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteNatGatewaySourceIpTranslationNatRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteNatGatewaySourceIpTranslationNatRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteNetDetectRequest struct {
	*tchttp.BaseRequest

	// The `ID` of a network detection instance, such as `netd-12345678`.
	NetDetectId *string `json:"NetDetectId,omitempty" name:"NetDetectId"`
}

func (r *DeleteNetDetectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteNetDetectRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteNetDetectResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteNetDetectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteNetDetectResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteNetworkAclRequest struct {
	*tchttp.BaseRequest

	// Network ACL instance ID. Example: acl-12345678.
	NetworkAclId *string `json:"NetworkAclId,omitempty" name:"NetworkAclId"`
}

func (r *DeleteNetworkAclRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteNetworkAclRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteNetworkAclResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteNetworkAclResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteNetworkAclResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteNetworkInterfaceRequest struct {
	*tchttp.BaseRequest

	// The ID of the ENI instance, such as `eni-m6dyj72l`.
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`
}

func (r *DeleteNetworkInterfaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteNetworkInterfaceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteNetworkInterfaceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteNetworkInterfaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteNetworkInterfaceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteRouteTableRequest struct {
	*tchttp.BaseRequest

	// The route table instance ID, such as `rtb-azd4dt1c`.
	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`
}

func (r *DeleteRouteTableRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteRouteTableRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteRouteTableResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteRouteTableResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteRouteTableResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteRoutesRequest struct {
	*tchttp.BaseRequest

	// Route table instance ID.
	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`

	// Routing policy object. Only the `RouteId` field is required when deleting a routing policy.
	Routes []*Route `json:"Routes,omitempty" name:"Routes" list`
}

func (r *DeleteRoutesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteRoutesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteRoutesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Details of the routing policy that has been deleted.
		RouteSet []*Route `json:"RouteSet,omitempty" name:"RouteSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteRoutesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteRoutesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteSecurityGroupPoliciesRequest struct {
	*tchttp.BaseRequest

	// The security group instance ID, such as `sg-33ocnj9n`. This can be obtained through DescribeSecurityGroups.
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`

	// The policy set of the security group. One request can only delete one or more policies in one direction. Both PolicyIndex-matching deletion and security group policy-matching deletion methods are supported. Each request can use only one matching method.
	SecurityGroupPolicySet *SecurityGroupPolicySet `json:"SecurityGroupPolicySet,omitempty" name:"SecurityGroupPolicySet"`
}

func (r *DeleteSecurityGroupPoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteSecurityGroupPoliciesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteSecurityGroupPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSecurityGroupPoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteSecurityGroupPoliciesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteSecurityGroupRequest struct {
	*tchttp.BaseRequest

	// The security group instance ID, such as `sg-33ocnj9n`. This can be obtained through DescribeSecurityGroups.
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
}

func (r *DeleteSecurityGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteSecurityGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteSecurityGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSecurityGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteSecurityGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteServiceTemplateGroupRequest struct {
	*tchttp.BaseRequest

	// The protocol port template group instance ID, such as `ppmg-n17uxvve`.
	ServiceTemplateGroupId *string `json:"ServiceTemplateGroupId,omitempty" name:"ServiceTemplateGroupId"`
}

func (r *DeleteServiceTemplateGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteServiceTemplateGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteServiceTemplateGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteServiceTemplateGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteServiceTemplateGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteServiceTemplateRequest struct {
	*tchttp.BaseRequest

	// Protocol port template instance ID, such as `ppm-e6dy460g`.
	ServiceTemplateId *string `json:"ServiceTemplateId,omitempty" name:"ServiceTemplateId"`
}

func (r *DeleteServiceTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteServiceTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteServiceTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteServiceTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteServiceTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteSubnetRequest struct {
	*tchttp.BaseRequest

	// The ID of the subnet instance. You can obtain the parameter value from the SubnetId field in the returned result of DescribeSubnets API.
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
}

func (r *DeleteSubnetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteSubnetRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteSubnetResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSubnetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteSubnetResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteVpcRequest struct {
	*tchttp.BaseRequest

	// The ID of the VPC instance. You can obtain the parameter value from the VpcId field in the returned result of DescribeVpcs API.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
}

func (r *DeleteVpcRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteVpcRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteVpcResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteVpcResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteVpcResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteVpnConnectionRequest struct {
	*tchttp.BaseRequest

	// The ID of the VPN gateway instance.
	VpnGatewayId *string `json:"VpnGatewayId,omitempty" name:"VpnGatewayId"`

	// The ID of the VPN tunnel instance, such as `vpnx-f49l6u0z`.
	VpnConnectionId *string `json:"VpnConnectionId,omitempty" name:"VpnConnectionId"`
}

func (r *DeleteVpnConnectionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteVpnConnectionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteVpnConnectionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteVpnConnectionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteVpnConnectionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteVpnGatewayRequest struct {
	*tchttp.BaseRequest

	// The ID of the VPN gateway instance.
	VpnGatewayId *string `json:"VpnGatewayId,omitempty" name:"VpnGatewayId"`
}

func (r *DeleteVpnGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteVpnGatewayRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteVpnGatewayResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteVpnGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteVpnGatewayResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountAttributesRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAccountAttributesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAccountAttributesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountAttributesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// User account attribute object
		AccountAttributeSet []*AccountAttribute `json:"AccountAttributeSet,omitempty" name:"AccountAttributeSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccountAttributesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAccountAttributesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAddressQuotaRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAddressQuotaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAddressQuotaRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAddressQuotaResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The quota information of EIPs in an account.
		QuotaSet []*Quota `json:"QuotaSet,omitempty" name:"QuotaSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAddressQuotaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAddressQuotaResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAddressTemplateGroupsRequest struct {
	*tchttp.BaseRequest

	// Filter conditions.
	// <li>address-template-group-name - String - (Filter condition) IP address template group name.</li>
	// <li>address-template-group-id - String - (Filter condition) IP address template group instance ID, such as `ipmg-mdunqeb6`.</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// Offset. The default value is 0.
	Offset *string `json:"Offset,omitempty" name:"Offset"`

	// Number of values to be returned. The default value is 20. Maximum is 100.
	Limit *string `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeAddressTemplateGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAddressTemplateGroupsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAddressTemplateGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The number of instances meeting the filter condition.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// IP address template.
		AddressTemplateGroupSet []*AddressTemplateGroup `json:"AddressTemplateGroupSet,omitempty" name:"AddressTemplateGroupSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAddressTemplateGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAddressTemplateGroupsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAddressTemplatesRequest struct {
	*tchttp.BaseRequest

	// Filter conditions.
	// <li>address-template-name - String - (Filter condition) IP address template name.</li>
	// <li>address-template-id - String - (Filter condition) IP address template instance ID, such as `ipm-mdunqeb6`.</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// Offset. The default value is 0.
	Offset *string `json:"Offset,omitempty" name:"Offset"`

	// Number of values to be returned. The default value is 20. Maximum is 100.
	Limit *string `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeAddressTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAddressTemplatesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAddressTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The number of instances meeting the filter condition.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// IP address template.
		AddressTemplateSet []*AddressTemplate `json:"AddressTemplateSet,omitempty" name:"AddressTemplateSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAddressTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAddressTemplatesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAddressesRequest struct {
	*tchttp.BaseRequest

	// The list of unique IDs of EIPs, such as `eip-11112222`. `AddressIds` and `Filters` cannot be specified at the same time.
	AddressIds []*string `json:"AddressIds,omitempty" name:"AddressIds" list`

	// The upper limit for `Filters` in each request is 10 and 5 for `Filter.Values`. `AddressIds` and `Filters` cannot be specified at the same time. Detailed filter conditions are as follows:
	// <li> address-id - String - Required: no - (Filter condition) The unique EIP ID, such as `eip-11112222`.</li>
	// <li> address-name - String - Required: no - (Filter condition) The EIP name. Fuzzy filtering is not supported.</li>
	// <li> address-ip - String - Required: no - (Filter condition) Filters by EIP.</li>
	// <li> address-status - String - Required: no - (Filter condition) The EIP state. Possible EIP states are: 'CREATING', 'BINDING', 'BIND', 'UNBINDING', 'UNBIND', 'OFFLINING', and 'BIND_ENI'.</li>
	// <li> instance-id - String - Required: no - (Filter condition) The ID of the instance bound to the EIP, such as `ins-11112222`.</li>
	// <li> private-ip-address - String - Required: no - (Filter condition) The private IP address bound to the EIP.</li>
	// <li> network-interface-id - String - Required: no - (Filter condition) The ID of the ENI bound to the EIP, such as `eni-11112222`.</li>
	// <li> is-arrears - String - Required: no - (Filter condition) Whether the EIP is in arrears. (TRUE: the EIP is in arrears | FALSE: the billing status of the EIP is normal)</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// The Offset. The default value is 0. For more information on `Offset`, see the relevant sections in API [Overview](https://intl.cloud.tencent.com/document/product/11646).
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned results. The default value is 20. The maximum is 100. For more information on `Limit`, see the relevant sections in API [Overview](https://intl.cloud.tencent.com/document/product/11646).
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeAddressesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAddressesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAddressesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of EIPs meeting the condition.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// List of EIPs details.
		AddressSet []*Address `json:"AddressSet,omitempty" name:"AddressSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAddressesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAddressesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAssistantCidrRequest struct {
	*tchttp.BaseRequest

	// The ID of a VPC instance set, such as `vpc-6v2ht8q5`.
	VpcIds []*string `json:"VpcIds,omitempty" name:"VpcIds" list`

	// Filter condition. `NetworkInterfaceIds` and `Filters` cannot be specified at the same time.
	// <li>vpc-id - String - (Filter condition) VPC instance ID, such as `vpc-f49l6u0z`.</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeAssistantCidrRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAssistantCidrRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAssistantCidrResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// A set of eligible secondary CIDR blocks
	// Note: This field may return null, indicating that no valid value was found.
		AssistantCidrSet []*AssistantCidr `json:"AssistantCidrSet,omitempty" name:"AssistantCidrSet" list`

		// Number of eligible instances.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssistantCidrResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAssistantCidrResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBandwidthPackageBillUsageRequest struct {
	*tchttp.BaseRequest

	// Unique ID of the pay-as-you-go bandwidth package.
	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" name:"BandwidthPackageId"`
}

func (r *DescribeBandwidthPackageBillUsageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBandwidthPackageBillUsageRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBandwidthPackageBillUsageResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Current billable usage.
		BandwidthPackageBillBandwidthSet []*BandwidthPackageBillBandwidth `json:"BandwidthPackageBillBandwidthSet,omitempty" name:"BandwidthPackageBillBandwidthSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBandwidthPackageBillUsageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBandwidthPackageBillUsageResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBandwidthPackageQuotaRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeBandwidthPackageQuotaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBandwidthPackageQuotaRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBandwidthPackageQuotaResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The quota of the bandwidth package.
		QuotaSet []*Quota `json:"QuotaSet,omitempty" name:"QuotaSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBandwidthPackageQuotaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBandwidthPackageQuotaResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBandwidthPackageResourcesRequest struct {
	*tchttp.BaseRequest

	// Unique ID of the bandwidth package in the form of `bwp-11112222`.
	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" name:"BandwidthPackageId"`

	// Each request can have up to 10 `Filters` and 5 `Filter.Values`. `AddressIds` and `Filters` cannot be specified at the same time. The specific filter conditions are as follows:
	// <li>resource-id - String - Required: no -  (Filter condition) Filters by the unique ID of resources in a bandwidth package, such as `eip-11112222`.</li>
	// <li>resource-type - String - Required: no - (Filter condition) Filters by the type of resources in a bandwidth package. It now supports only EIP (`Address`) and load balancer (`LoadBalance`).</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// The offset. Default value: 0. For more information on `Offset`, see the relevant sections in API [Introduction](https://intl.cloud.tencent.com/document/api/213/11646?from_cn_redirect=1).
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// The number of returned results. Default value: 20. Maximum value: 100. For more information on `Limit`, see the relevant sections in API [Introduction](https://intl.cloud.tencent.com/document/api/213/11646?from_cn_redirect=1).
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeBandwidthPackageResourcesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBandwidthPackageResourcesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBandwidthPackageResourcesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The number of eligible resources in the bandwidth package.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The list of resources in the bandwidth package.
		ResourceSet []*Resource `json:"ResourceSet,omitempty" name:"ResourceSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBandwidthPackageResourcesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBandwidthPackageResourcesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBandwidthPackagesRequest struct {
	*tchttp.BaseRequest

	// The unique ID list of bandwidth packages.
	BandwidthPackageIds []*string `json:"BandwidthPackageIds,omitempty" name:"BandwidthPackageIds" list`

	// Each request can have up to 10 `Filters`. `BandwidthPackageIds` and `Filters` cannot be specified at the same time. The specific filter conditions are as follows:
	// <li> bandwidth-package_id - String - Required: No - (Filter condition) Filter by the unique ID of the bandwidth package.</li>
	// <li> bandwidth-package-name - String - Required: No - (Filter condition) Filter by the bandwidth package name. Fuzzy filtering is not supported.</li>
	// <li> network-type - String - Required: No - (Filter condition) Filter by the bandwidth package type. Types include 'BGP', 'SINGLEISP', and 'ANYCAST'.</li>
	// <li> charge-type - String - Required: No - (Filter condition) Filter by the bandwidth package billing mode. Billing modes include 'TOP5_POSTPAID_BY_MONTH' and 'PERCENT95_POSTPAID_BY_MONTH'.</li>
	// <li> resource.resource-type - String - Required: No - (Filter condition) Filter by the bandwidth package resource type. Resource types include 'Address' and 'LoadBalance'.</li>
	// <li> resource.resource-id - String - Required: No - (Filter condition) Filter by the bandwidth package resource ID, such as 'eip-xxxx' and 'lb-xxxx'.</li>
	// <li> resource.address-ip - String - Required: No - (Filter condition) Filter by the bandwidth package resource IP.</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// Offset of the query results
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Max number of the bandwidth packages to be returned.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeBandwidthPackagesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBandwidthPackagesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBandwidthPackagesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The number of eligible bandwidth packages.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Detail information of the bandwidth package.
		BandwidthPackageSet []*BandwidthPackage `json:"BandwidthPackageSet,omitempty" name:"BandwidthPackageSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBandwidthPackagesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBandwidthPackagesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCcnAttachedInstancesRequest struct {
	*tchttp.BaseRequest

	// Offset
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// The returned quantity
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Filter conditions:
	// <li>ccn-id - String - (Filter condition) The CCN instance ID.</li>
	// <li>instance-type - String - (Filter condition) The associated instance type.</li>
	// <li>instance-region - String - (Filter condition) The associated instance region.</li>
	// <li>instance-type - String - (Filter condition) The instance ID of the associated instance.</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// The ID of the CCN instance
	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`

	// The order field supports `CcnId`, `InstanceType`, `InstanceId`, `InstanceName`, `InstanceRegion`, `AttachedTime`, and `State`.
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`

	// Order methods. Ascending: `ASC`, Descending: `DESC`.
	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`
}

func (r *DescribeCcnAttachedInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCcnAttachedInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCcnAttachedInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The number of objects meeting the condition.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The list of associated instances.
		InstanceSet []*CcnAttachedInstance `json:"InstanceSet,omitempty" name:"InstanceSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCcnAttachedInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCcnAttachedInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCcnRoutesRequest struct {
	*tchttp.BaseRequest

	// The CCN instance ID, such as `ccn-gree226l`.
	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`

	// The unique ID of the CCN routing policy, such as `ccnr-f49l6u0z`.
	RouteIds []*string `json:"RouteIds,omitempty" name:"RouteIds" list`

	// Filter condition. `RouteIds` and `Filters` cannot be specified at the same time.
	// <li>route-id - String - (Filter condition) Routing policy ID.</li>
	// <li>cidr-block - String - (Filter condition) Destination port.</li>
	// <li>instance-type - String - (Filter condition) The next hop type.</li>
	// <li>instance-region - String - (Filter condition) The next hop region.</li>
	// <li>instance-type - String - (Filter condition) The instance ID of the next hop.</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// Offset
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// The returned quantity
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeCcnRoutesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCcnRoutesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCcnRoutesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The number of objects meeting the condition.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The CCN routing policy object.
		RouteSet []*CcnRoute `json:"RouteSet,omitempty" name:"RouteSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCcnRoutesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCcnRoutesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCcnsRequest struct {
	*tchttp.BaseRequest

	// The CCN instance ID, such as `ccn-f49l6u0z`. Each request can have a maximum of 100 instances. `CcnIds` and `Filters` cannot be specified at the same time
	CcnIds []*string `json:"CcnIds,omitempty" name:"CcnIds" list`

	// Filter conditions. `CcnIds` and `Filters` cannot be specified at the same time.
	// <li>ccn-id - String - (Filter condition) The unique ID of the CCN, such as `vpc-f49l6u0z`.</li>
	// <li>ccn-name - String - (Filter condition) The CCN name.</li>
	// <li>ccn-description - String - (Filter condition) CCN description.</li>
	// <li>state - String - (Filter condition) The instance status. 'ISOLATED': Isolated (the account is in arrears and the service is suspended.) 'AVAILABLE': Running.</li>
	// <li>tag-key - String - Required: no - (Filter condition) Filters by tag key.</li>
	// <li>`tag:tag-key` - String - Required: no - (Filter condition) Filters by tag key pair. For this parameter, `tag-key` will be replaced with a specific tag key. For more information, see this example: **Querying the list of CCNs bound to tags**.</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// Offset
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// The returned quantity
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Order fields support `CcnId`, `CcnName`, `CreateTime`, `State`, and `QosLevel`
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`

	// Order methods. Ascending: `ASC`, Descending: `DESC`.
	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`
}

func (r *DescribeCcnsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCcnsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCcnsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The number of objects meeting the condition.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// CCN object.
		CcnSet []*CCN `json:"CcnSet,omitempty" name:"CcnSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCcnsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCcnsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeClassicLinkInstancesRequest struct {
	*tchttp.BaseRequest

	// Filter conditions.
	// <li>vpc-id - String - (Filter condition) The VPC instance ID.</li>
	// <li>vm-ip - String - (Filter condition) The IP address of the CVM on the basic network.</li>
	Filters []*FilterObject `json:"Filters,omitempty" name:"Filters" list`

	// Offset
	Offset *string `json:"Offset,omitempty" name:"Offset"`

	// The returned quantity
	Limit *string `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeClassicLinkInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeClassicLinkInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeClassicLinkInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The number of instances meeting the filter condition.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Classiclink instance.
		ClassicLinkInstanceSet []*ClassicLinkInstance `json:"ClassicLinkInstanceSet,omitempty" name:"ClassicLinkInstanceSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClassicLinkInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeClassicLinkInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCrossBorderComplianceRequest struct {
	*tchttp.BaseRequest

	// (Exact match) Service provider. Valid values: `UNICOM`.
	ServiceProvider *string `json:"ServiceProvider,omitempty" name:"ServiceProvider"`

	// (Exact match) ID of compliance review request.
	ComplianceId *uint64 `json:"ComplianceId,omitempty" name:"ComplianceId"`

	// (Fuzzy match) Company name.
	Company *string `json:"Company,omitempty" name:"Company"`

	// (Fuzzy match) Unified Social Credit Code.
	UniformSocialCreditCode *string `json:"UniformSocialCreditCode,omitempty" name:"UniformSocialCreditCode"`

	// (Fuzzy match) Legal person.
	LegalPerson *string `json:"LegalPerson,omitempty" name:"LegalPerson"`

	// (Fuzzy match) Issuing authority.
	IssuingAuthority *string `json:"IssuingAuthority,omitempty" name:"IssuingAuthority"`

	// (Fuzzy match) Business address.
	BusinessAddress *string `json:"BusinessAddress,omitempty" name:"BusinessAddress"`

	// (Exact match) Zip code.
	PostCode *uint64 `json:"PostCode,omitempty" name:"PostCode"`

	// (Fuzzy match) Operator.
	Manager *string `json:"Manager,omitempty" name:"Manager"`

	// (Exact match) Operator ID card number.
	ManagerId *string `json:"ManagerId,omitempty" name:"ManagerId"`

	// (Fuzzy match) Operator address.
	ManagerAddress *string `json:"ManagerAddress,omitempty" name:"ManagerAddress"`

	// (Exact match) Operator phone number.
	ManagerTelephone *string `json:"ManagerTelephone,omitempty" name:"ManagerTelephone"`

	// (Exact match) Email.
	Email *string `json:"Email,omitempty" name:"Email"`

	// (Exact match) Service start date, such as `2020-07-28`.
	ServiceStartDate *string `json:"ServiceStartDate,omitempty" name:"ServiceStartDate"`

	// (Exact match) Service end date, such as `2020-07-28`.
	ServiceEndDate *string `json:"ServiceEndDate,omitempty" name:"ServiceEndDate"`

	// (Exact match) Status. Valid values: `PENDING`, `APPROVED`, and `DENY`.
	State *string `json:"State,omitempty" name:"State"`
}

func (r *DescribeCrossBorderComplianceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCrossBorderComplianceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCrossBorderComplianceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// List of compliance review requests.
		CrossBorderComplianceSet []*CrossBorderCompliance `json:"CrossBorderComplianceSet,omitempty" name:"CrossBorderComplianceSet" list`

		// Total number of compliance review requests.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCrossBorderComplianceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCrossBorderComplianceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCustomerGatewayVendorsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeCustomerGatewayVendorsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCustomerGatewayVendorsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCustomerGatewayVendorsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Customer gateway vendor information object.
		CustomerGatewayVendorSet []*CustomerGatewayVendor `json:"CustomerGatewayVendorSet,omitempty" name:"CustomerGatewayVendorSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCustomerGatewayVendorsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCustomerGatewayVendorsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCustomerGatewaysRequest struct {
	*tchttp.BaseRequest

	// Customer gateway ID, such as `cgw-2wqq41m9`. Each request can have a maximum of 100 instances. `CustomerGatewayIds` and `Filters` cannot be specified at the same time.
	CustomerGatewayIds []*string `json:"CustomerGatewayIds,omitempty" name:"CustomerGatewayIds" list`

	// The filter condition. For details, see the Instance Filter Conditions Table. The upper limit for `Filters` in each request is 10 and 5 for `Filter.Values`. `CustomerGatewayIds` and `Filters` cannot be specified at the same time.
	// <li>customer-gateway-id - String - (Filter condition) The unique ID of the user gateway, such as `cgw-mgp33pll`.</li>
	// <li>customer-gateway-name - String - (Filter condition) The name of the user gateway, such as `test-cgw`.</li>
	// <li>ip-address - String - (Filter condition) The public IP address, such as `58.211.1.12`.</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// The Offset. The default value is 0. For more information about Offset, see the relevant section in the API Introduction.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeCustomerGatewaysRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCustomerGatewaysRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCustomerGatewaysResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Customer gateway object list
		CustomerGatewaySet []*CustomerGateway `json:"CustomerGatewaySet,omitempty" name:"CustomerGatewaySet" list`

		// Number of eligible instances
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCustomerGatewaysResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCustomerGatewaysResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDirectConnectGatewayCcnRoutesRequest struct {
	*tchttp.BaseRequest

	// The ID of the Direct Connect gateway, such as `dcg-prpqlmg1`.
	DirectConnectGatewayId *string `json:"DirectConnectGatewayId,omitempty" name:"DirectConnectGatewayId"`

	// The route learning type of the CCN. Available values:
	// <li>`BGP` - Automatic learning.</li>
	// <li>`STATIC` - Static means user-configured. This is the default value.</li>
	CcnRouteType *string `json:"CcnRouteType,omitempty" name:"CcnRouteType"`

	// Offset.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// The returned quantity.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeDirectConnectGatewayCcnRoutesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDirectConnectGatewayCcnRoutesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDirectConnectGatewayCcnRoutesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The number of objects meeting the condition.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The CCN route (IDC IP range) list.
		RouteSet []*DirectConnectGatewayCcnRoute `json:"RouteSet,omitempty" name:"RouteSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDirectConnectGatewayCcnRoutesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDirectConnectGatewayCcnRoutesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDirectConnectGatewaysRequest struct {
	*tchttp.BaseRequest

	// The unique ID of the direct connect gateway, such as `dcg-9o233uri`.
	DirectConnectGatewayIds []*string `json:"DirectConnectGatewayIds,omitempty" name:"DirectConnectGatewayIds" list`

	// Filter condition. `DirectConnectGatewayIds` and `Filters` cannot be specified at the same time.
	// <li>direct-connect-gateway-id - String - The unique ID of the direct connect gateway, such as `dcg-9o233uri`.</li>
	// <li>direct-connect-gateway-name - String - The name of the direct connect gateway. The default is fuzzy query.</li>
	// <li>direct-connect-gateway-ip - String - The IP of the direct connect gateway.</li>
	// <li>gateway-type - String - The gateway type. Valid values: `NORMAL` (Standard type), `NAT` (NAT type).</li>
	// <li>network-type- String - The network type. Valid values: `VPC` (VPC type), `CCN` (CCN type).</li>
	// <li>ccn-id - String - The ID of the CCN where the direct connect gateway resides.</li>
	// <li>vpc-id - String - The ID of the VPC where the direct connect gateway resides.</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// The offset.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Max number of results returned
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeDirectConnectGatewaysRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDirectConnectGatewaysRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDirectConnectGatewaysResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The number of eligible objects.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The object array of the direct connect gateway.
		DirectConnectGatewaySet []*DirectConnectGateway `json:"DirectConnectGatewaySet,omitempty" name:"DirectConnectGatewaySet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDirectConnectGatewaysResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDirectConnectGatewaysResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeFlowLogRequest struct {
	*tchttp.BaseRequest

	// ID of the VPC instance
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// The unique ID of the flow log.
	FlowLogId *string `json:"FlowLogId,omitempty" name:"FlowLogId"`
}

func (r *DescribeFlowLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeFlowLogRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeFlowLogResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The flow log information.
		FlowLog []*FlowLog `json:"FlowLog,omitempty" name:"FlowLog" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFlowLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeFlowLogResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeFlowLogsRequest struct {
	*tchttp.BaseRequest

	// ID of the VPC instance
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// The unique ID of the flow log.
	FlowLogId *string `json:"FlowLogId,omitempty" name:"FlowLogId"`

	// The name of the flow log instance.
	FlowLogName *string `json:"FlowLogName,omitempty" name:"FlowLogName"`

	// The resource type of the flow log. Valid values: 'VPC', 'SUBNET', and 'NETWORKINTERFACE'.
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`

	// The unique ID of the resource.
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// Type of flow logs to be collected. Valid values: 'ACCEPT', 'REJECT' and 'ALL'.
	TrafficType *string `json:"TrafficType,omitempty" name:"TrafficType"`

	// The storage ID of the flow log.
	CloudLogId *string `json:"CloudLogId,omitempty" name:"CloudLogId"`

	// The storage ID status of the flow log.
	CloudLogState *string `json:"CloudLogState,omitempty" name:"CloudLogState"`

	// Order by field. Valid values: 'flowLogName' and 'createTime'. Default value: 'createTime'.
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`

	// In ascending (`asc`) or descending (`desc`) order. Default value: 'desc'.
	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`

	// The offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// The number of rows per page. Default value: 10.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Filter condition. `FlowLogIds` and `Filters` cannot be specified at the same time.
	// <li>tag-key - String - Required: No - (Filter condition) Filter by tag key.</li>
	// <li> tag:tag-key - String - Required: No - (Filter condition) Filter by tag key-value pair. The tag-key should be replaced with a specified tag key.</li>
	Filters *Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeFlowLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeFlowLogsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeFlowLogsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The instance set of flow logs.
		FlowLog []*FlowLog `json:"FlowLog,omitempty" name:"FlowLog" list`

		// The total number of flow logs.
		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFlowLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeFlowLogsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeGatewayFlowMonitorDetailRequest struct {
	*tchttp.BaseRequest

	// The point in time. This indicates details of this minute will be queried. For example, in `2019-02-28 18:15:20`, details at `18:15` will be queried.
	TimePoint *string `json:"TimePoint,omitempty" name:"TimePoint"`

	// The instance ID of the VPN gateway, such as `vpn-ltjahce6`.
	VpnId *string `json:"VpnId,omitempty" name:"VpnId"`

	// The instance ID of the Direct Connect gateway, such as `dcg-ltjahce6`.
	DirectConnectGatewayId *string `json:"DirectConnectGatewayId,omitempty" name:"DirectConnectGatewayId"`

	// The instance ID of the peering connection, such as `pcx-ltjahce6`.
	PeeringConnectionId *string `json:"PeeringConnectionId,omitempty" name:"PeeringConnectionId"`

	// The instance ID of the NAT gateway, such as `nat-ltjahce6`.
	NatId *string `json:"NatId,omitempty" name:"NatId"`

	// Offset.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// The returned quantity.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// The order field supports `InPkg`, `OutPkg`, `InTraffic`, and `OutTraffic`.
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`

	// Order methods. Ascending: `ASC`, Descending: `DESC`.
	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`
}

func (r *DescribeGatewayFlowMonitorDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeGatewayFlowMonitorDetailRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeGatewayFlowMonitorDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The number of objects meeting the condition.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The gateway traffic monitoring details.
		GatewayFlowMonitorDetailSet []*GatewayFlowMonitorDetail `json:"GatewayFlowMonitorDetailSet,omitempty" name:"GatewayFlowMonitorDetailSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGatewayFlowMonitorDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeGatewayFlowMonitorDetailResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeGatewayFlowQosRequest struct {
	*tchttp.BaseRequest

	// Gateway instance ID, which currently supports these types:
	// ID of Direct Connect gateway instance, e.g. `dcg-ltjahce6`;
	// ID of NAT gateway instance, e.g. `nat-ltjahce6`;
	// ID of VPN gateway instance, e.g. `vpn-ltjahce6`.
	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`

	// CVM private IP addresses with limited bandwidth.
	IpAddresses []*string `json:"IpAddresses,omitempty" name:"IpAddresses" list`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeGatewayFlowQosRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeGatewayFlowQosRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeGatewayFlowQosResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// List of instance details.
		GatewayQosSet []*GatewayQos `json:"GatewayQosSet,omitempty" name:"GatewayQosSet" list`

		// Number of eligible instances.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGatewayFlowQosResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeGatewayFlowQosResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeHaVipsRequest struct {
	*tchttp.BaseRequest

	// The unique `ID` of the `HAVIP`, such as `havip-9o233uri`.
	HaVipIds []*string `json:"HaVipIds,omitempty" name:"HaVipIds" list`

	// Filter condition. `HaVipIds` and `Filters` cannot be specified at the same time.
	// li>havip-id - String - The unique ID of the HAVIP, such as `havip-9o233uri`.</li>
	// <li>havip-name - String - HAVIP name.</li>
	// <li>vpc-id - String - VPC ID of the HAVIP.</li>
	// <li>subnet-id - String - Subnet ID of the HAVIP.</li>
	// <li>vip - String - Virtual IP address of the HAVIP.</li>
	// <li>address-ip - String - Bound EIP.</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// Offset
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// The returned quantity
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeHaVipsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeHaVipsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeHaVipsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The number of objects meeting the condition.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// `HAVIP` object array.
		HaVipSet []*HaVip `json:"HaVipSet,omitempty" name:"HaVipSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeHaVipsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeHaVipsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeIpGeolocationDatabaseUrlRequest struct {
	*tchttp.BaseRequest

	// Protocol type for an IP location database. Valid value: `ipv4`.
	Type *string `json:"Type,omitempty" name:"Type"`
}

func (r *DescribeIpGeolocationDatabaseUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeIpGeolocationDatabaseUrlRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeIpGeolocationDatabaseUrlResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Download link of an IP location database
		DownLoadUrl *string `json:"DownLoadUrl,omitempty" name:"DownLoadUrl"`

		// Link expiration time in UTC format following the ISO8601 standard.
		ExpiredAt *string `json:"ExpiredAt,omitempty" name:"ExpiredAt"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeIpGeolocationDatabaseUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeIpGeolocationDatabaseUrlResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeIpGeolocationInfosRequest struct {
	*tchttp.BaseRequest

	// IP addresses to be queried. Both IPv4 and IPv6 addresses are supported.
	AddressIps []*string `json:"AddressIps,omitempty" name:"AddressIps" list`

	// Fields of the IP addresses to be queried, including `Country`, `Province`, `City`, `Region`, `Isp`, `AsName` and `AsId`
	Fields *IpField `json:"Fields,omitempty" name:"Fields"`
}

func (r *DescribeIpGeolocationInfosRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeIpGeolocationInfosRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeIpGeolocationInfosResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// IP address details
		AddressInfo []*IpGeolocationInfo `json:"AddressInfo,omitempty" name:"AddressInfo" list`

		// Number of IP addresses
		Total *int64 `json:"Total,omitempty" name:"Total"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeIpGeolocationInfosResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeIpGeolocationInfosResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeNatGatewayDestinationIpPortTranslationNatRulesRequest struct {
	*tchttp.BaseRequest

	// NAT gateway ID.
	NatGatewayIds []*string `json:"NatGatewayIds,omitempty" name:"NatGatewayIds" list`

	// Filter conditions:
	// `NatGatewayIds` and `Filters` cannot be specified at the same time.
	// <li> nat-gateway-id, the NAT gateway ID, such as `nat-0yi4hekt`.</li>
	// <li> vpc-id, the VPC ID, such as `vpc-0yi4hekt`.</li>
	// <li> public-ip-address, the EIP, such as `139.199.232.238`.</li>
	// <li>public-port, the public network port.</li>
	// <li>private-ip-address, the private IP, such as `10.0.0.1`.</li>
	// <li>private-port, the private network port.</li>
	// <li>description, the rule description.</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// Offset. The default value is 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of values to be returned. The default value is 20. Maximum is 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeNatGatewayDestinationIpPortTranslationNatRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeNatGatewayDestinationIpPortTranslationNatRulesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeNatGatewayDestinationIpPortTranslationNatRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The object array of port forwarding rules for the NAT gateway.
		NatGatewayDestinationIpPortTranslationNatRuleSet []*NatGatewayDestinationIpPortTranslationNatRule `json:"NatGatewayDestinationIpPortTranslationNatRuleSet,omitempty" name:"NatGatewayDestinationIpPortTranslationNatRuleSet" list`

		// The number of object arrays of NAT port forwarding rules meeting the conditions.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNatGatewayDestinationIpPortTranslationNatRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeNatGatewayDestinationIpPortTranslationNatRulesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeNatGatewaySourceIpTranslationNatRulesRequest struct {
	*tchttp.BaseRequest

	// The unique ID of the NAT Gateway, such as `nat-123xx454`.
	NatGatewayId *string `json:"NatGatewayId,omitempty" name:"NatGatewayId"`

	// Filter conditions:
	// <li> resource-id, the subnet ID (such as `subnet-0yi4hekt`) or CVM ID</li>
	// <li> public-ip-address, the EIP, such as `139.199.232.238`</li>
	// <li>description, the rule description</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// Offset. Default is 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeNatGatewaySourceIpTranslationNatRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeNatGatewaySourceIpTranslationNatRulesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeNatGatewaySourceIpTranslationNatRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Object array of the SNAT rule for a NAT Gateway.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
		SourceIpTranslationNatRuleSet []*SourceIpTranslationNatRule `json:"SourceIpTranslationNatRuleSet,omitempty" name:"SourceIpTranslationNatRuleSet" list`

		// The number of object arrays of eligible forwarding rules for a NAT Gateway
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNatGatewaySourceIpTranslationNatRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeNatGatewaySourceIpTranslationNatRulesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeNatGatewaysRequest struct {
	*tchttp.BaseRequest

	// The unified ID of the NAT gateways, such as `nat-123xx454`.
	NatGatewayIds []*string `json:"NatGatewayIds,omitempty" name:"NatGatewayIds" list`

	// Filter condition. `NatGatewayIds` and `Filters` cannot be specified at the same time.
	// <li>nat-gateway-id - String - (Filter condition) The ID of the protocol port template instance, such as `nat-123xx454`.</li>
	// <li>vpc-id - String - (Filter condition) The unique ID of the VPC, such as `vpc-123xx454`.</li>
	// <li>nat-gateway-name - String - (Filter condition) The name of the protocol port template instance, such as `test_nat`.</li>
	// <li>tag-key - String - (Filter condition) The tag key, such as `test-key`.</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// Offset. The default value is 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of values to be returned. The default value is 20. Maximum is 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeNatGatewaysRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeNatGatewaysRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeNatGatewaysResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// NAT gateway object array.
		NatGatewaySet []*NatGateway `json:"NatGatewaySet,omitempty" name:"NatGatewaySet" list`

		// The number of NAT gateway object sets meeting the conditions.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNatGatewaysResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeNatGatewaysResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeNetDetectStatesRequest struct {
	*tchttp.BaseRequest

	// The array of network detection instance `IDs`, such as [`netd-12345678`].
	NetDetectIds []*string `json:"NetDetectIds,omitempty" name:"NetDetectIds" list`

	// Filter conditions. `NetDetectIds` and `Filters` cannot be specified at the same time.
	// <li>net-detect-id - String - (Filter condition) The network detection instance ID, such as netd-12345678.</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// The offset. Default: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// The number of returned values. Default: 20. Maximum: 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeNetDetectStatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeNetDetectStatesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeNetDetectStatesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The array of network detection verification results that meet requirements.
	// Note: This field may return null, indicating that no valid values can be obtained.
		NetDetectStateSet []*NetDetectState `json:"NetDetectStateSet,omitempty" name:"NetDetectStateSet" list`

		// The number of network detection verification results that meet requirements.
	// Note: This field may return null, indicating that no valid values can be obtained.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNetDetectStatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeNetDetectStatesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeNetDetectsRequest struct {
	*tchttp.BaseRequest

	// The array of network detection instance `IDs`, such as [`netd-12345678`].
	NetDetectIds []*string `json:"NetDetectIds,omitempty" name:"NetDetectIds" list`

	// Filter conditions. `NetDetectIds` and `Filters` cannot be specified at the same time.
	// <li>vpc-id - String - (Filter condition) The VPC instance ID, such as vpc-12345678.</li>
	// <li>net-detect-id - String - (Filter condition) The network detection instance ID, such as netd-12345678.</li>
	// <li>subnet-id - String - (Filter condition) The subnet instance ID, such as subnet-12345678.</li>
	// <li>net-detect-name - String - (Filter condition) The network detection name.</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// The offset. Default: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// The number of returned values. Default: 20. Maximum: 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeNetDetectsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeNetDetectsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeNetDetectsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The array of network detection objects that meet requirements.
	// Note: This field may return null, indicating that no valid values can be obtained.
		NetDetectSet []*NetDetect `json:"NetDetectSet,omitempty" name:"NetDetectSet" list`

		// The number of network detection objects that meet requirements.
	// Note: This field may return null, indicating that no valid values can be obtained.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNetDetectsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeNetDetectsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeNetworkAclsRequest struct {
	*tchttp.BaseRequest

	// Array of network ACL instance IDs, such as [acl-12345678]. Up to 100 instances are allowed for each request. This parameter does not support specifying `NetworkAclIds` and `Filters` at the same time.
	NetworkAclIds []*string `json:"NetworkAclIds,omitempty" name:"NetworkAclIds" list`

	// Filter condition. `NetworkAclIds` and `Filters` cannot be specified at the same time.
	// <li>vpc-id - String - (Filter condition) VPC instance ID, such as vpc-12345678.</li>
	// <li>network-acl-id - String - (Filter condition) Network ACL instance ID, such as acl-12345678.</li>
	// <li>network-acl-name - String - (Filter condition) Network ACL instance name.</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// Offset. Default: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Returned quantity. Default: 20. Value range: 1-100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeNetworkAclsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeNetworkAclsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeNetworkAclsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// List of instance details.
		NetworkAclSet []*NetworkAcl `json:"NetworkAclSet,omitempty" name:"NetworkAclSet" list`

		// Number of eligible instances.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNetworkAclsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeNetworkAclsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeNetworkInterfaceLimitRequest struct {
	*tchttp.BaseRequest

	// ID of a CVM instance or ENI to query
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeNetworkInterfaceLimitRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeNetworkInterfaceLimitRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeNetworkInterfaceLimitResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// ENI quota
		EniQuantity *int64 `json:"EniQuantity,omitempty" name:"EniQuantity"`

		// Quota of IP addresses that can be allocated to each ENI.
		EniPrivateIpAddressQuantity *int64 `json:"EniPrivateIpAddressQuantity,omitempty" name:"EniPrivateIpAddressQuantity"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNetworkInterfaceLimitResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeNetworkInterfaceLimitResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeNetworkInterfacesRequest struct {
	*tchttp.BaseRequest

	// Queries the ID of the ENI instance, such as `eni-pxir56ns`. Each request can have a maximum of 100 instances. `NetworkInterfaceIds` and `Filters` cannot be specified at the same time.
	NetworkInterfaceIds []*string `json:"NetworkInterfaceIds,omitempty" name:"NetworkInterfaceIds" list`

	// Filter condition. `NetworkInterfaceIds` and `Filters` cannot be specified at the same time.
	// <li>vpc-id - String - (Filter condition) VPC instance ID, such as `vpc-f49l6u0z`.</li>
	// <li>subnet-id - String - (Filter condition) Subnet instance ID, such as `subnet-f49l6u0z`.</li>
	// <li>network-interface-id - String - (Filter condition) ENI instance ID, such as `eni-5k56k7k7`.</li>
	// <li>attachment.instance-id - String - (Filter condition) CVM instance ID, such as `ins-3nqpdn3i`.</li>
	// <li>groups.security-group-id - String - (Filter condition) Instance ID of the security group, such as `sg-f9ekbxeq`.</li>
	// <li>network-interface-name - String - (Filter condition) ENI instance name.</li>
	// <li>network-interface-description - String - (Filter condition) ENI instance description.</li>
	// <li>address-ip - String - (Filter condition) Private IPv4 address.</li>
	// <li>tag-key - String - Required: no - (Filter condition) Filters by tag key. For more information, see Example 2.</li>
	// <li> `tag:tag-key` - String - Required: no - (Filter condition) Filters by tag key pair. For this parameter, `tag-key` will be replaced with a specific tag key. For more information, see Example 3.</li>
	// <li>is-primary - Boolean - Required: no - (Filter condition) Filters based on whether it is a primary ENI. If the value is ‘true’, filter only the primary ENI. If the value is ‘false’, filter only the secondary ENI. If the secondary filter parameter is provided, filter the both.</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeNetworkInterfacesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeNetworkInterfacesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeNetworkInterfacesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// List of instance details.
		NetworkInterfaceSet []*NetworkInterface `json:"NetworkInterfaceSet,omitempty" name:"NetworkInterfaceSet" list`

		// Number of eligible instances.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNetworkInterfacesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeNetworkInterfacesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRouteTablesRequest struct {
	*tchttp.BaseRequest

	// The route table instance ID, such as `rtb-azd4dt1c`.
	RouteTableIds []*string `json:"RouteTableIds,omitempty" name:"RouteTableIds" list`

	// Filter condition. `RouteTableIds` and `Filters` cannot be speified at the same time.
	// <li>route-table-id - String - (Filter condition) Route table instance ID.</li>
	// <li>route-table-name - String - (Filter condition) Route table name.</li>
	// <li>vpc-id - String - (Filter condition) VPC instance ID, such as `vpc-f49l6u0z`.</li>
	// <li>association.main - String - (Filter condition) Whether it is the main route table.</li>
	// <li>tag-key - String - Required: No - (Filter condition) Filter by tag key.</li>
	// <li>tag:tag-key - String - Required: No - (Filter condition) Filter by tag key-value pair. The tag-key is replaced with the specific tag key. For usage, refer to case 2.</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// Offset.
	Offset *string `json:"Offset,omitempty" name:"Offset"`

	// The number of request objects.
	Limit *string `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeRouteTablesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRouteTablesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRouteTablesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of eligible instances.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Route table object.
		RouteTableSet []*RouteTable `json:"RouteTableSet,omitempty" name:"RouteTableSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRouteTablesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRouteTablesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSecurityGroupAssociationStatisticsRequest struct {
	*tchttp.BaseRequest

	// The Security instance ID, such as `sg-33ocnj9n`. It can be obtained through DescribeSecurityGroups.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds" list`
}

func (r *DescribeSecurityGroupAssociationStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSecurityGroupAssociationStatisticsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSecurityGroupAssociationStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Statistics on the instances associated with a security group.
		SecurityGroupAssociationStatisticsSet []*SecurityGroupAssociationStatistics `json:"SecurityGroupAssociationStatisticsSet,omitempty" name:"SecurityGroupAssociationStatisticsSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSecurityGroupAssociationStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSecurityGroupAssociationStatisticsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSecurityGroupPoliciesRequest struct {
	*tchttp.BaseRequest

	// The security group instance ID, such as `sg-33ocnj9n`. It can be obtained through DescribeSecurityGroups.
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
}

func (r *DescribeSecurityGroupPoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSecurityGroupPoliciesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSecurityGroupPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Security group policy set.
		SecurityGroupPolicySet *SecurityGroupPolicySet `json:"SecurityGroupPolicySet,omitempty" name:"SecurityGroupPolicySet"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSecurityGroupPoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSecurityGroupPoliciesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSecurityGroupReferencesRequest struct {
	*tchttp.BaseRequest

	// A set of security group instance IDs, e.g. ['sg-12345678']
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds" list`
}

func (r *DescribeSecurityGroupReferencesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSecurityGroupReferencesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSecurityGroupReferencesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Referred security groups.
		ReferredSecurityGroupSet []*ReferredSecurityGroup `json:"ReferredSecurityGroupSet,omitempty" name:"ReferredSecurityGroupSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSecurityGroupReferencesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSecurityGroupReferencesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSecurityGroupsRequest struct {
	*tchttp.BaseRequest

	// The security group instance ID, such as `sg-33ocnj9n`. It can be obtained through `DescribeSecurityGroups`. Each request can have a maximum of 100 instances. `SecurityGroupIds` and `Filters` cannot be specified at the same time.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds" list`

	// Filter conditions. `SecurityGroupIds` and `Filters` cannot be specified at the same time.
	// <li>security-group-id - String - (Filter condition) The security group ID.</li>
	// <li>project-id - Integer - (Filter condition) The project ID.</li>
	// <li>security-group-name - String - (Filter condition) The security group name.</li>
	// <li>tag-key - String - Required: no - (Filter condition) Filters by tag key. For more information, see Example 2.</li>
	// <li> `tag:tag-key` - String - Required: no - (Filter condition) Filters by tag key pair. For this parameter, `tag-key` will be replaced with a specific tag key. For more information, see Example 3.</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// Offset. Default value: 0.
	Offset *string `json:"Offset,omitempty" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *string `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSecurityGroupsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Security group object.
		SecurityGroupSet []*SecurityGroup `json:"SecurityGroupSet,omitempty" name:"SecurityGroupSet" list`

		// The number of instances meeting the filter condition.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSecurityGroupsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceTemplateGroupsRequest struct {
	*tchttp.BaseRequest

	// Filter conditions.
	// <li>service-template-group-name - String - (Filter condition) Protocol port template group name.</li>
	// <li>service-template-group-id - String - (Filter condition) Protocol port template group instance ID, such as `ppmg-e6dy460g`.</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// Offset. The default value is 0.
	Offset *string `json:"Offset,omitempty" name:"Offset"`

	// Number of values to be returned. The default value is 20. Maximum is 100.
	Limit *string `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeServiceTemplateGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeServiceTemplateGroupsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceTemplateGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The number of instances meeting the filter condition.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Protocol port template group.
		ServiceTemplateGroupSet []*ServiceTemplateGroup `json:"ServiceTemplateGroupSet,omitempty" name:"ServiceTemplateGroupSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeServiceTemplateGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeServiceTemplateGroupsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceTemplatesRequest struct {
	*tchttp.BaseRequest

	// Filter conditions.
	// <li>service-template-name - String - (Filter condition) Protocol port template name.</li>
	// <li>service-template-id - String - (Filter condition) Protocol port template instance ID, such as `ppm-e6dy460g`.</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// Offset. The default value is 0.
	Offset *string `json:"Offset,omitempty" name:"Offset"`

	// Number of values to be returned. The default value is 20. Maximum is 100.
	Limit *string `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeServiceTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeServiceTemplatesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The number of instances meeting the filter condition.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Protocol port template object.
		ServiceTemplateSet []*ServiceTemplate `json:"ServiceTemplateSet,omitempty" name:"ServiceTemplateSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeServiceTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeServiceTemplatesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSubnetsRequest struct {
	*tchttp.BaseRequest

	// Queries the ID of the subnet instance, such as `subnet-pxir56ns`. Each request can have a maximum of 100 instances. `SubnetIds` and `Filters` cannot be specified at the same time.
	SubnetIds []*string `json:"SubnetIds,omitempty" name:"SubnetIds" list`

	// Filter condition. `SubnetIds` and `Filters` cannot be specified at the same time.
	// <li>subnet-id - String - (Filter condition) Subnet instance name.</li>
	// <li>vpc-id - String - (Filter condition) VPC instance ID, such as `vpc-f49l6u0z`.</li>
	// <li>cidr-block - String - (Filter condition) The subnet IP range, such as 192.168.1.0.</li>
	// <li>is-default - Boolean - (Filter condition) Whether it is the default subnet.</li>
	// <li>is-remote-vpc-snat - Boolean - (Filter condition) Whether it is a VPC SNAT address pool subnet.</li>
	// <li>subnet-name - String - (Filter condition) Subnet name.</li>
	// <li>zone - String - (Filter condition) Availability zone.</li>
	// <li>tag-key - String - Required: No - (Filter condition) Filter by tag key.</li>
	// <li>tag:tag-key - String - Required: No - (Filter condition) Filter by tag key-value pair. The tag-key is replaced with the specific tag key. For usage, refer to case 2.</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// Offset. Default value: 0.
	Offset *string `json:"Offset,omitempty" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *string `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeSubnetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSubnetsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSubnetsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The number of instances meeting the filter condition.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Subnet object.
		SubnetSet []*Subnet `json:"SubnetSet,omitempty" name:"SubnetSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSubnetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSubnetsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskResultRequest struct {
	*tchttp.BaseRequest

	// Async task ID. Either TaskId or DealName must be entered.
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

	// Billing order No. Either TaskId or DealName must be entered.
	DealName *string `json:"DealName,omitempty" name:"DealName"`
}

func (r *DescribeTaskResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTaskResultRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskResultResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Job ID
		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

		// The execution results, including `SUCCESS`, `FAILED`, and `RUNNING`
		Result *string `json:"Result,omitempty" name:"Result"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTaskResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTaskResultResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcInstancesRequest struct {
	*tchttp.BaseRequest

	// Filter condition. `RouteTableIds` and `Filters` cannot be specified at the same time.
	// <li>vpc-id - String - (Filter condition) VPC instance ID, such as `vpc-f49l6u0z`.</li>
	// <li>instance-type - String - (Filter condition) CVM instance ID.</li>
	// <li>instance-name - String - (Filter condition) CVM name.</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// Offset.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// The number of requested objects.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeVpcInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeVpcInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// List of CVM instances.
		InstanceSet []*CvmInstance `json:"InstanceSet,omitempty" name:"InstanceSet" list`

		// The number of eligible CVM instances.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVpcInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeVpcInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcIpv6AddressesRequest struct {
	*tchttp.BaseRequest

	// The `ID` of the `VPC`, such as `vpc-f49l6u0z`.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// The `IP` address list. Each request supports a maximum of `10` batch querying.
	Ipv6Addresses []*string `json:"Ipv6Addresses,omitempty" name:"Ipv6Addresses" list`

	// Offset.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// The returned quantity.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeVpcIpv6AddressesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeVpcIpv6AddressesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcIpv6AddressesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The `IPv6` address list.
		Ipv6AddressSet []*VpcIpv6Address `json:"Ipv6AddressSet,omitempty" name:"Ipv6AddressSet" list`

		// The total number of `IPv6` addresses.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVpcIpv6AddressesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeVpcIpv6AddressesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcPrivateIpAddressesRequest struct {
	*tchttp.BaseRequest

	// The `ID` of the `VPC`, such as `vpc-f49l6u0z`.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// The private `IP` address list. Each request supports a maximum of `10` batch querying.
	PrivateIpAddresses []*string `json:"PrivateIpAddresses,omitempty" name:"PrivateIpAddresses" list`
}

func (r *DescribeVpcPrivateIpAddressesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeVpcPrivateIpAddressesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcPrivateIpAddressesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The list of private `IP` address information.
		VpcPrivateIpAddressSet []*VpcPrivateIpAddress `json:"VpcPrivateIpAddressSet,omitempty" name:"VpcPrivateIpAddressSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVpcPrivateIpAddressesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeVpcPrivateIpAddressesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcResourceDashboardRequest struct {
	*tchttp.BaseRequest

	// Vpc instance ID, e.g. vpc-f1xjkw1b.
	VpcIds []*string `json:"VpcIds,omitempty" name:"VpcIds" list`
}

func (r *DescribeVpcResourceDashboardRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeVpcResourceDashboardRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcResourceDashboardResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// List of resource objects.
		ResourceDashboardSet []*ResourceDashboard `json:"ResourceDashboardSet,omitempty" name:"ResourceDashboardSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVpcResourceDashboardResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeVpcResourceDashboardResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcsRequest struct {
	*tchttp.BaseRequest

	// The VPC instance ID, such as `vpc-f49l6u0z`. Each request supports a maximum of 100 instances. `VpcIds` and `Filters` cannot be specified at the same time.
	VpcIds []*string `json:"VpcIds,omitempty" name:"VpcIds" list`

	// Filter condition. `VpcIds` and `Filters` cannot be specified at the same time.
	// <li>vpc-name - String - (Filter condition) VPC instance name.</li>
	// <li>is-default - String - (Filter condition) Indicates whether it is the default VPC.</li>
	// <li>vpc-id - String - (Filter condition) VPC instance ID, such as `vpc-f49l6u0z`.</li>
	// <li>cidr-block - String - (Filter condition) VPC CIDR.</li>
	// <li>tag-key - String - Required: No - (Filter condition) Filter by tag key.</li>
	// <li>tag:tag-key - String - Required: No - (Filter condition) Filter by tag key-value pair. The tag-key is replaced with the specific tag key. For usage, refer to case 2.</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// Offset. Default value: 0.
	Offset *string `json:"Offset,omitempty" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *string `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeVpcsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeVpcsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The number of objects meeting the condition.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The VPC object.
		VpcSet []*Vpc `json:"VpcSet,omitempty" name:"VpcSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVpcsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeVpcsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeVpnConnectionsRequest struct {
	*tchttp.BaseRequest

	// The instance ID of the VPN tunnel, such as `vpnx-f49l6u0z`. Each request can have a maximum of 100 instances. `VpnConnectionIds` and `Filters` cannot be specified at the same time.
	VpnConnectionIds []*string `json:"VpnConnectionIds,omitempty" name:"VpnConnectionIds" list`

	// Filter condition. In each request, the upper limit for `Filters` is 10 and 5 for `Filter.Values`. `VpnConnectionIds` and `Filters` cannot be specified at the same time.
	// <li>vpc-id - String - VPC instance ID, such as `vpc-0a36uwkr`.</li>
	// <li>vpn-gateway-id - String - VPN gateway instance ID, such as `vpngw-p4lmqawn`.</li>
	// <li>customer-gateway-id - String - Customer gateway instance ID, such as `cgw-l4rblw63`.</li>
	// <li>vpn-connection-name - String - Connection name, such as `test-vpn`.</li>
	// <li>vpn-connection-id - String - Connection instance ID, such as `vpnx-5p7vkch8"`.</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// The Offset. The default value is 0. For more information about Offset, see the relevant section in the API Introduction.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of values to be returned. The default value is 20. Maximum is 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeVpnConnectionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeVpnConnectionsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeVpnConnectionsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The number of instances meeting the filter condition.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// VPN tunnel instance.
		VpnConnectionSet []*VpnConnection `json:"VpnConnectionSet,omitempty" name:"VpnConnectionSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVpnConnectionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeVpnConnectionsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeVpnGatewayCcnRoutesRequest struct {
	*tchttp.BaseRequest

	// The ID of the VPN gateway instance.
	VpnGatewayId *string `json:"VpnGatewayId,omitempty" name:"VpnGatewayId"`

	// Offset.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// The returned quantity
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeVpnGatewayCcnRoutesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeVpnGatewayCcnRoutesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeVpnGatewayCcnRoutesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The CCN route (IDC IP range) list.
		RouteSet []*VpngwCcnRoutes `json:"RouteSet,omitempty" name:"RouteSet" list`

		// Number of objects that meet the condition.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVpnGatewayCcnRoutesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeVpnGatewayCcnRoutesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeVpnGatewaysRequest struct {
	*tchttp.BaseRequest

	// The VPN gateway instance ID, such as `vpngw-f49l6u0z`. Each request can have a maximum of 100 instances. `VpnGatewayIds` and `Filters` cannot be specified at the same time.
	VpnGatewayIds []*string `json:"VpnGatewayIds,omitempty" name:"VpnGatewayIds" list`

	// Filter condition. `VpnGatewayIds` and `Filters` cannot be specified at the same time.
	// <li>vpc-id - String - (Filter condition) VPC instance ID, such as `vpc-f49l6u0z`.</li>
	// <li>vpn-gateway-id - String - (Filter condition) VPN instance ID, such as `vpngw-5aluhh9t`.</li>
	// <li>vpn-gateway-name - String - (Filter condition) VPN instance name.</li>
	// <li>type - String - (Filter condition) VPN gateway type: 'IPSEC', 'SSL'.</li>
	// <li>public-ip-address- String - (Filter condition) Public IP.</li>
	// <li>renew-flag - String - (Filter condition) Gateway renewal type. Manual renewal: `NOTIFY_AND_MANUAL_RENEW`, Automatic renewal: `NOTIFY_AND_AUTO_RENEW`.</li>
	// <li>zone - String - (Filter condition) The availability zone where the VPN is located, such as `ap-guangzhou-2`.</li>
	Filters []*FilterObject `json:"Filters,omitempty" name:"Filters" list`

	// Offset
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// The number of request objects.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeVpnGatewaysRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeVpnGatewaysRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeVpnGatewaysResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The number of instances meeting the filter condition.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The list of details of VPN gateway instances.
		VpnGatewaySet []*VpnGateway `json:"VpnGatewaySet,omitempty" name:"VpnGatewaySet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVpnGatewaysResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeVpnGatewaysResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DestinationIpPortTranslationNatRule struct {

	// Network protocol. Available choices: `TCP`, `UDP`.
	IpProtocol *string `json:"IpProtocol,omitempty" name:"IpProtocol"`

	// EIP.
	PublicIpAddress *string `json:"PublicIpAddress,omitempty" name:"PublicIpAddress"`

	// Public port.
	PublicPort *uint64 `json:"PublicPort,omitempty" name:"PublicPort"`

	// Private network address.
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" name:"PrivateIpAddress"`

	// Private network port.
	PrivatePort *uint64 `json:"PrivatePort,omitempty" name:"PrivatePort"`

	// NAT gateway forwarding rule description.
	Description *string `json:"Description,omitempty" name:"Description"`
}

type DetachCcnInstancesRequest struct {
	*tchttp.BaseRequest

	// The CCN instance ID, such as `ccn-f49l6u0z`.
	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`

	// The list of network instances to be unbound
	Instances []*CcnInstance `json:"Instances,omitempty" name:"Instances" list`
}

func (r *DetachCcnInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DetachCcnInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DetachCcnInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DetachCcnInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DetachCcnInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DetachClassicLinkVpcRequest struct {
	*tchttp.BaseRequest

	// The ID of the VPC instance. You can obtain the parameter value from the VpcId field in the returned result of DescribeVpcs API.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Queries the ID of the CVM instance, such as `ins-r8hr2upy`.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`
}

func (r *DetachClassicLinkVpcRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DetachClassicLinkVpcRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DetachClassicLinkVpcResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DetachClassicLinkVpcResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DetachClassicLinkVpcResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DetachNetworkInterfaceRequest struct {
	*tchttp.BaseRequest

	// The ID of the ENI instance, such as `eni-m6dyj72l`.
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`

	// The ID of the CVM instance, such as `ins-r8hr2upy`.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DetachNetworkInterfaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DetachNetworkInterfaceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DetachNetworkInterfaceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DetachNetworkInterfaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DetachNetworkInterfaceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DirectConnectGateway struct {

	// Direct Connect `ID`.
	DirectConnectGatewayId *string `json:"DirectConnectGatewayId,omitempty" name:"DirectConnectGatewayId"`

	// Direct Connect gateway name.
	DirectConnectGatewayName *string `json:"DirectConnectGatewayName,omitempty" name:"DirectConnectGatewayName"`

	// The `ID` of the `VPC` instance associated with the Direct Connect gateway.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// The associated network type:
	// <li>`VPC` - VPC</li>
	// <li>`CCN` - CCN</li>
	NetworkType *string `json:"NetworkType,omitempty" name:"NetworkType"`

	// The `ID` of the associated network instance:
	// <li>When the NetworkType is `VPC`, this value is the VPC instance `ID`</li>
	// <li>When the NetworkType is `CCN`, this value is the CCN instance `ID`</li>
	NetworkInstanceId *string `json:"NetworkInstanceId,omitempty" name:"NetworkInstanceId"`

	// Gateway type:
	// <li>NORMAL - Standard type. Note: CCN only supports the standard type</li>
	// <li>NAT - NAT type</li>
	// NAT type supports network address switch configuration. After the type is confirmed, it cannot be modified. A VPC can create one NAT-type Direct Connect gateway and one non-NAT-type Direct Connect gateway
	GatewayType *string `json:"GatewayType,omitempty" name:"GatewayType"`

	// Creation Time.
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Direct Connect gateway IP.
	DirectConnectGatewayIp *string `json:"DirectConnectGatewayIp,omitempty" name:"DirectConnectGatewayIp"`

	// The `ID` of the `CCN` instance associated with the Direct Connect gateway.
	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`

	// The route-learning type of the CCN:
	// <li>`BGP` - Automatic learning.</li>
	// <li>`STATIC` - Static, that is, user-configured.</li>
	CcnRouteType *string `json:"CcnRouteType,omitempty" name:"CcnRouteType"`

	// Whether BGP is enabled.
	EnableBGP *bool `json:"EnableBGP,omitempty" name:"EnableBGP"`

	// Whether to enable BGP's `community` attribute. Valid values: enable, disable
	EnableBGPCommunity *bool `json:"EnableBGPCommunity,omitempty" name:"EnableBGPCommunity"`

	// ID of the NAT gateway bound.
	// Note: this field may return `null`, indicating that no valid value was found.
	NatGatewayId *string `json:"NatGatewayId,omitempty" name:"NatGatewayId"`

	// Whether the direct connect gateway supports the VXLAN architecture.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	VXLANSupport []*bool `json:"VXLANSupport,omitempty" name:"VXLANSupport" list`

	// CCN route publishing mode. Valid values: `standard` and `exquisite`.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ModeType *string `json:"ModeType,omitempty" name:"ModeType"`
}

type DirectConnectGatewayCcnRoute struct {

	// Route ID.
	RouteId *string `json:"RouteId,omitempty" name:"RouteId"`

	// IDC IP range.
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" name:"DestinationCidrBlock"`

	// The `AS-Path` attribute of `BGP`.
	ASPath []*string `json:"ASPath,omitempty" name:"ASPath" list`
}

type DisableCcnRoutesRequest struct {
	*tchttp.BaseRequest

	// The CCN instance ID, such as `ccn-f49l6u0z`.
	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`

	// The unique ID of the CCN routing policy, such as `ccnr-f49l6u0z`.
	RouteIds []*string `json:"RouteIds,omitempty" name:"RouteIds" list`
}

func (r *DisableCcnRoutesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DisableCcnRoutesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DisableCcnRoutesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DisableCcnRoutesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DisableCcnRoutesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DisableGatewayFlowMonitorRequest struct {
	*tchttp.BaseRequest

	// Gateway instance ID, which currently supports these types:
	// ID of Direct Connect gateway instance, e.g. `dcg-ltjahce6`;
	// ID of NAT gateway instance, e.g. `nat-ltjahce6`;
	// ID of VPN gateway instance, e.g. `vpn-ltjahce6`.
	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
}

func (r *DisableGatewayFlowMonitorRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DisableGatewayFlowMonitorRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DisableGatewayFlowMonitorResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DisableGatewayFlowMonitorResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DisableGatewayFlowMonitorResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DisassociateAddressRequest struct {
	*tchttp.BaseRequest

	// The unique ID of the EIP, such as `eip-11112222`.
	AddressId *string `json:"AddressId,omitempty" name:"AddressId"`

	// Whether a common public IP is assigned after the EIP is unbound. Value range:<br><li>TRUE: Indicates that after the EIP is unbound, a common public IP is assigned.<br><li>FALSE: Indicates that after the EIP is unbound, a common public IP is not assigned.<br>Default value: FALSE.<br><br>The parameter can be specified only under the following conditions:<br><li>It can only be specified when you unbind an EIP from the primary private IP of the primary ENI.<br><li>After an EIP is unbound, you can assign public IPs to an account up to 10 times per day. For more information, use the [DescribeAddressQuota] (https://intl.cloud.tencent.com/document/api/213/1378?from_cn_redirect=1) API.
	ReallocateNormalPublicIp *bool `json:"ReallocateNormalPublicIp,omitempty" name:"ReallocateNormalPublicIp"`
}

func (r *DisassociateAddressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DisassociateAddressRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DisassociateAddressResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The async task ID. You can use the [DescribeTaskResult](https://intl.cloud.tencent.com/document/api/215/36271?from_cn_redirect=1) API to query the task status.
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DisassociateAddressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DisassociateAddressResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DisassociateDirectConnectGatewayNatGatewayRequest struct {
	*tchttp.BaseRequest

	// The direct connect gateway ID.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// The NAT Gateway ID.
	NatGatewayId *string `json:"NatGatewayId,omitempty" name:"NatGatewayId"`

	// The ID of the VPC instance, which can be obtained from the `VpcId` field in response of the `DescribeVpcs` API.
	DirectConnectGatewayId *string `json:"DirectConnectGatewayId,omitempty" name:"DirectConnectGatewayId"`
}

func (r *DisassociateDirectConnectGatewayNatGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DisassociateDirectConnectGatewayNatGatewayRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DisassociateDirectConnectGatewayNatGatewayResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DisassociateDirectConnectGatewayNatGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DisassociateDirectConnectGatewayNatGatewayResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DisassociateNatGatewayAddressRequest struct {
	*tchttp.BaseRequest

	// The ID of the NAT gateway, such as `nat-df45454`.
	NatGatewayId *string `json:"NatGatewayId,omitempty" name:"NatGatewayId"`

	// Array of the EIPs to be unbound from the NAT gateway.
	PublicIpAddresses []*string `json:"PublicIpAddresses,omitempty" name:"PublicIpAddresses" list`
}

func (r *DisassociateNatGatewayAddressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DisassociateNatGatewayAddressRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DisassociateNatGatewayAddressResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DisassociateNatGatewayAddressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DisassociateNatGatewayAddressResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DisassociateNetworkAclSubnetsRequest struct {
	*tchttp.BaseRequest

	// Network ACL instance ID. Example: acl-12345678.
	NetworkAclId *string `json:"NetworkAclId,omitempty" name:"NetworkAclId"`

	// Array of subnet instance IDs. Example: [subnet-12345678].
	SubnetIds []*string `json:"SubnetIds,omitempty" name:"SubnetIds" list`
}

func (r *DisassociateNetworkAclSubnetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DisassociateNetworkAclSubnetsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DisassociateNetworkAclSubnetsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DisassociateNetworkAclSubnetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DisassociateNetworkAclSubnetsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DisassociateNetworkInterfaceSecurityGroupsRequest struct {
	*tchttp.BaseRequest

	// ENI instance ID, e.g. eni-pxir56ns. You can enter up to 100 instances for each request.
	NetworkInterfaceIds []*string `json:"NetworkInterfaceIds,omitempty" name:"NetworkInterfaceIds" list`

	// The security group instance ID, such as `sg-33ocnj9n`. It can be obtained through DescribeSecurityGroups. You can enter up to 100 instances for each request.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds" list`
}

func (r *DisassociateNetworkInterfaceSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DisassociateNetworkInterfaceSecurityGroupsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DisassociateNetworkInterfaceSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DisassociateNetworkInterfaceSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DisassociateNetworkInterfaceSecurityGroupsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DownloadCustomerGatewayConfigurationRequest struct {
	*tchttp.BaseRequest

	// The ID of the VPN gateway instance.
	VpnGatewayId *string `json:"VpnGatewayId,omitempty" name:"VpnGatewayId"`

	// The ID of the VPN tunnel instance, such as `vpnx-f49l6u0z`.
	VpnConnectionId *string `json:"VpnConnectionId,omitempty" name:"VpnConnectionId"`

	// Customer gateway vendor information object, which can be obtained through DescribeCustomerGatewayVendors.
	CustomerGatewayVendor *CustomerGatewayVendor `json:"CustomerGatewayVendor,omitempty" name:"CustomerGatewayVendor"`

	// Name of the physical API for tunnel access device.
	InterfaceName *string `json:"InterfaceName,omitempty" name:"InterfaceName"`
}

func (r *DownloadCustomerGatewayConfigurationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DownloadCustomerGatewayConfigurationRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DownloadCustomerGatewayConfigurationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Configuration information in XML format.
		CustomerGatewayConfiguration *string `json:"CustomerGatewayConfiguration,omitempty" name:"CustomerGatewayConfiguration"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DownloadCustomerGatewayConfigurationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DownloadCustomerGatewayConfigurationResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EnableCcnRoutesRequest struct {
	*tchttp.BaseRequest

	// The CCN instance ID, such as `ccn-f49l6u0z`.
	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`

	// The unique ID of the CCN routing policy, such as `ccnr-f49l6u0z`.
	RouteIds []*string `json:"RouteIds,omitempty" name:"RouteIds" list`
}

func (r *EnableCcnRoutesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *EnableCcnRoutesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EnableCcnRoutesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EnableCcnRoutesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *EnableCcnRoutesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EnableGatewayFlowMonitorRequest struct {
	*tchttp.BaseRequest

	// Gateway instance ID, which currently supports these types:
	// ID of Direct Connect gateway instance, e.g. `dcg-ltjahce6`;
	// ID of NAT gateway instance, e.g. `nat-ltjahce6`;
	// ID of VPN gateway instance, e.g. `vpn-ltjahce6`.
	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
}

func (r *EnableGatewayFlowMonitorRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *EnableGatewayFlowMonitorRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EnableGatewayFlowMonitorResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EnableGatewayFlowMonitorResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *EnableGatewayFlowMonitorResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Filter struct {

	// The attribute name. If more than one Filter exists, the logical relation between these Filters is `AND`.
	Name *string `json:"Name,omitempty" name:"Name"`

	// The attribute value. If there are multiple Values for one Filter, the logical relation between these Values under the same Filter is `OR`.
	Values []*string `json:"Values,omitempty" name:"Values" list`
}

type FilterObject struct {

	// The attribute name. If more than one Filter exists, the logical relation between these Filters is `AND`.
	Name *string `json:"Name,omitempty" name:"Name"`

	// The attribute value. If there are multiple Values for one Filter, the logical relation between these Values under the same Filter is `OR`.
	Values []*string `json:"Values,omitempty" name:"Values" list`
}

type FlowLog struct {

	// ID of the VPC instance
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// The unique ID of the flow log.
	FlowLogId *string `json:"FlowLogId,omitempty" name:"FlowLogId"`

	// The name of the flow log instance.
	FlowLogName *string `json:"FlowLogName,omitempty" name:"FlowLogName"`

	// The resource type of the flow log. Valid values: 'VPC', 'SUBNET', and 'NETWORKINTERFACE'.
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`

	// The unique ID of the resource.
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// Type of flow logs to be collected. Valid values: `ACCEPT`, `REJECT` and `ALL`.
	TrafficType *string `json:"TrafficType,omitempty" name:"TrafficType"`

	// The storage ID of the flow log.
	CloudLogId *string `json:"CloudLogId,omitempty" name:"CloudLogId"`

	// The storage ID status of the flow log.
	CloudLogState *string `json:"CloudLogState,omitempty" name:"CloudLogState"`

	// The flow log description.
	FlowLogDescription *string `json:"FlowLogDescription,omitempty" name:"FlowLogDescription"`

	// The creation time of the flow log.
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// Tag list, such as [{"Key": "city", "Value": "shanghai"}]
	TagSet []*Tag `json:"TagSet,omitempty" name:"TagSet" list`
}

type GatewayFlowMonitorDetail struct {

	// Origin `IP`.
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" name:"PrivateIpAddress"`

	// Inbound packets.
	InPkg *uint64 `json:"InPkg,omitempty" name:"InPkg"`

	// Outbound packets.
	OutPkg *uint64 `json:"OutPkg,omitempty" name:"OutPkg"`

	// Inbound traffic, in Byte.
	InTraffic *uint64 `json:"InTraffic,omitempty" name:"InTraffic"`

	// Outbound traffic, in Byte.
	OutTraffic *uint64 `json:"OutTraffic,omitempty" name:"OutTraffic"`
}

type GatewayQos struct {

	// VPC instance ID.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// CVM Private IP.
	IpAddress *string `json:"IpAddress,omitempty" name:"IpAddress"`

	// Bandwidth limit value.
	Bandwidth *int64 `json:"Bandwidth,omitempty" name:"Bandwidth"`

	// The creation time.
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type GetCcnRegionBandwidthLimitsRequest struct {
	*tchttp.BaseRequest

	// The CCN instance ID, such as `ccn-f49l6u0z`.
	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`

	// The filter condition.
	// <li>sregion - String - (Filter condition) Filter by the source region, such as 'ap-guangzhou'.</li>
	// <li>dregion - String - (Filter condition) Filter by the destination region, such as 'ap-shanghai-bm'.</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// The sorting condition. Valid values: `BandwidthLimit` and `ExpireTime`.
	SortedBy *string `json:"SortedBy,omitempty" name:"SortedBy"`

	// The offset.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// The returned quantity.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// In ascending or descending order. Valid values: 'ASC' and 'DESC'.
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
}

func (r *GetCcnRegionBandwidthLimitsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetCcnRegionBandwidthLimitsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetCcnRegionBandwidthLimitsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The outbound bandwidth limits of regions in a CCN instance.
	// Note: this field may return null, indicating that no valid value was found.
		CcnBandwidthSet []*CcnBandwidthInfo `json:"CcnBandwidthSet,omitempty" name:"CcnBandwidthSet" list`

		// The number of eligible objects.
	// Note: this field may return null, indicating that no valid value was found.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetCcnRegionBandwidthLimitsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetCcnRegionBandwidthLimitsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type HaVip struct {

	// The `ID` of the `HAVIP`. This is the unique identifier of the `HAVIP`.
	HaVipId *string `json:"HaVipId,omitempty" name:"HaVipId"`

	// The name of the `HAVIP`.
	HaVipName *string `json:"HaVipName,omitempty" name:"HaVipName"`

	// The virtual IP address.
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// The `ID` of the VPC to which the `HAVIP` belongs.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// The `ID` of the subnet to which the `HAVIP` belongs.
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// The `ID` of the ENI associated with the `HAVIP`.
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`

	// The `ID` of the bound instance.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Bound `EIP`.
	AddressIp *string `json:"AddressIp,omitempty" name:"AddressIp"`

	// Status:
	// <li>`AVAILABLE`: Operating</li>
	// <li>`UNBIND`: Not bound</li>
	State *string `json:"State,omitempty" name:"State"`

	// Creation Time.
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// Identifier for businesses that use HAVIP.
	Business *string `json:"Business,omitempty" name:"Business"`
}

type HaVipAssociateAddressIpRequest struct {
	*tchttp.BaseRequest

	// The unique `ID` of the `HAVIP`, such as `havip-9o233uri`. This must be a `HAVIP` that has not been bound to an `EIP`
	HaVipId *string `json:"HaVipId,omitempty" name:"HaVipId"`

	// The Elastic `IP`. This must be an `EIP` that has not been bound to a `HAVIP`
	AddressIp *string `json:"AddressIp,omitempty" name:"AddressIp"`
}

func (r *HaVipAssociateAddressIpRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *HaVipAssociateAddressIpRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type HaVipAssociateAddressIpResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *HaVipAssociateAddressIpResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *HaVipAssociateAddressIpResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type HaVipDisassociateAddressIpRequest struct {
	*tchttp.BaseRequest

	// The unique `ID` of the `HAVIP`, such as `havip-9o233uri`. This must be an `HAVIP` that has been bound to an `EIP`.
	HaVipId *string `json:"HaVipId,omitempty" name:"HaVipId"`
}

func (r *HaVipDisassociateAddressIpRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *HaVipDisassociateAddressIpRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type HaVipDisassociateAddressIpResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *HaVipDisassociateAddressIpResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *HaVipDisassociateAddressIpResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type IKEOptionsSpecification struct {

	// Encryption algorithm. Available values: '3DES-CBC', 'AES-CBC-128', 'AES-CBS-192', 'AES-CBC-256', and 'DES-CBC'. Default is 3DES-CBC.
	PropoEncryAlgorithm *string `json:"PropoEncryAlgorithm,omitempty" name:"PropoEncryAlgorithm"`

	// Verification algorithm. Available value: 'MD5' and 'SHA1'. Default is MD5.
	PropoAuthenAlgorithm *string `json:"PropoAuthenAlgorithm,omitempty" name:"PropoAuthenAlgorithm"`

	// Negotiation mode. Available values: 'AGGRESSIVE' and 'MAIN'. Default is MAIN.
	ExchangeMode *string `json:"ExchangeMode,omitempty" name:"ExchangeMode"`

	// Type of local identity. Available values: 'ADDRESS' and 'FQDN'. Default is ADDRESS.
	LocalIdentity *string `json:"LocalIdentity,omitempty" name:"LocalIdentity"`

	// Type of remote identity. Available values: 'ADDRESS' and 'FQDN'. Default is ADDRESS.
	RemoteIdentity *string `json:"RemoteIdentity,omitempty" name:"RemoteIdentity"`

	// Local identity. When ADDRESS is selected for LocalIdentity, LocalAddress is required. The default LocalAddress is the public IP of the VPN gateway.
	LocalAddress *string `json:"LocalAddress,omitempty" name:"LocalAddress"`

	// Remote identity. When ADDRESS is selected for RemoteIdentity, RemoteAddress is required.
	RemoteAddress *string `json:"RemoteAddress,omitempty" name:"RemoteAddress"`

	// Local identity. When FQDN is selected for LocalIdentity, LocalFqdnName is required.
	LocalFqdnName *string `json:"LocalFqdnName,omitempty" name:"LocalFqdnName"`

	// Remote identity. When FQDN is selected for RemoteIdentity, RemoteFqdnName is required.
	RemoteFqdnName *string `json:"RemoteFqdnName,omitempty" name:"RemoteFqdnName"`

	// DH group. Specify the DH group used for exchanging the key via IKE. Available values: 'GROUP1', 'GROUP2', 'GROUP5', 'GROUP14', and 'GROUP24'.
	DhGroupName *string `json:"DhGroupName,omitempty" name:"DhGroupName"`

	// IKE SA lifetime (in sec). Value range: 60-604800
	IKESaLifetimeSeconds *uint64 `json:"IKESaLifetimeSeconds,omitempty" name:"IKESaLifetimeSeconds"`

	// IKE version
	IKEVersion *string `json:"IKEVersion,omitempty" name:"IKEVersion"`
}

type IPSECOptionsSpecification struct {

	// Encryption algorithm. Available values: '3DES-CBC', 'AES-CBC-128', 'AES-CBC-192', 'AES-CBC-256', 'DES-CBC', and 'NULL'. Default is AES-CBC-128.
	EncryptAlgorithm *string `json:"EncryptAlgorithm,omitempty" name:"EncryptAlgorithm"`

	// Verification algorithm. Available value: 'MD5' and 'SHA1'. Default is:
	IntegrityAlgorith *string `json:"IntegrityAlgorith,omitempty" name:"IntegrityAlgorith"`

	// IPsec SA lifetime (in sec). Value range: 180-604800
	IPSECSaLifetimeSeconds *uint64 `json:"IPSECSaLifetimeSeconds,omitempty" name:"IPSECSaLifetimeSeconds"`

	// PFS. Available value: 'NULL', 'DH-GROUP1', 'DH-GROUP2', 'DH-GROUP5', 'DH-GROUP14', and 'DH-GROUP24'. Default is NULL.
	PfsDhGroup *string `json:"PfsDhGroup,omitempty" name:"PfsDhGroup"`

	// IPsec SA lifetime (in KB). Value range: 2560-604800
	IPSECSaLifetimeTraffic *uint64 `json:"IPSECSaLifetimeTraffic,omitempty" name:"IPSECSaLifetimeTraffic"`
}

type InquirePriceCreateDirectConnectGatewayRequest struct {
	*tchttp.BaseRequest
}

func (r *InquirePriceCreateDirectConnectGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquirePriceCreateDirectConnectGatewayRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquirePriceCreateDirectConnectGatewayResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Standard access fee for a direct connect gateway
	// Note: this field may return `null`, indicating that no valid values can be obtained.
		TotalCost *int64 `json:"TotalCost,omitempty" name:"TotalCost"`

		// Actual access fee for a direct connect gateway
	// Note: this field may return `null`, indicating that no valid values can be obtained.
		RealTotalCost *int64 `json:"RealTotalCost,omitempty" name:"RealTotalCost"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquirePriceCreateDirectConnectGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquirePriceCreateDirectConnectGatewayResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceCreateVpnGatewayRequest struct {
	*tchttp.BaseRequest

	// The public network bandwidth configuration. Available bandwidth specifications: 5, 10, 20, 50, and 100. Unit: Mbps.
	InternetMaxBandwidthOut *uint64 `json:"InternetMaxBandwidthOut,omitempty" name:"InternetMaxBandwidthOut"`

	// The VPN gateway billing mode. PREPAID: prepaid means monthly subscription. POSTPAID_BY_HOUR: postpaid means pay-as-you-go. Default: POSTPAID_BY_HOUR. If prepaid mode is specified, the `InstanceChargePrepaid` parameter must be entered.
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`

	// Parameter settings for prepaid billing mode, also known as monthly subscription. This parameter can specify the purchase period and other attributes such as auto-renewal. This parameter is mandatory for prepaid instances.
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitempty" name:"InstanceChargePrepaid"`
}

func (r *InquiryPriceCreateVpnGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceCreateVpnGatewayRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceCreateVpnGatewayResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Product price.
		Price *Price `json:"Price,omitempty" name:"Price"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceCreateVpnGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceCreateVpnGatewayResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceRenewVpnGatewayRequest struct {
	*tchttp.BaseRequest

	// The ID of the VPN gateway instance.
	VpnGatewayId *string `json:"VpnGatewayId,omitempty" name:"VpnGatewayId"`

	// Specifies the purchased validity period, whether to enable auto-renewal. This parameter is required for monthly-subscription instances.
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitempty" name:"InstanceChargePrepaid"`
}

func (r *InquiryPriceRenewVpnGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceRenewVpnGatewayRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceRenewVpnGatewayResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Product price.
		Price *Price `json:"Price,omitempty" name:"Price"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceRenewVpnGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceRenewVpnGatewayResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceResetVpnGatewayInternetMaxBandwidthRequest struct {
	*tchttp.BaseRequest

	// The ID of the VPN gateway instance.
	VpnGatewayId *string `json:"VpnGatewayId,omitempty" name:"VpnGatewayId"`

	// The public network bandwidth configuration. Available bandwidth specifications: 5, 10, 20, 50, and 100. Unit: Mbps.
	InternetMaxBandwidthOut *uint64 `json:"InternetMaxBandwidthOut,omitempty" name:"InternetMaxBandwidthOut"`
}

func (r *InquiryPriceResetVpnGatewayInternetMaxBandwidthRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceResetVpnGatewayInternetMaxBandwidthRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceResetVpnGatewayInternetMaxBandwidthResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Product price.
		Price *Price `json:"Price,omitempty" name:"Price"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceResetVpnGatewayInternetMaxBandwidthResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceResetVpnGatewayInternetMaxBandwidthResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InstanceChargePrepaid struct {

	// Purchased usage period (in month). Value range: [1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36].
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// Auto-renewal ID. Value range: NOTIFY_AND_AUTO_RENEW: notify expiry and renew automatically, NOTIFY_AND_MANUAL_RENEW: notify expiry but do not renew automatically. The default is NOTIFY_AND_MANUAL_RENEW
	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`
}

type InstanceStatistic struct {

	// Type of instance
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// Number of instances
	InstanceCount *uint64 `json:"InstanceCount,omitempty" name:"InstanceCount"`
}

type IpField struct {

	// Country/region of the IP
	Country *bool `json:"Country,omitempty" name:"Country"`

	// Province/municipality/state of the IP
	Province *bool `json:"Province,omitempty" name:"Province"`

	// City of the IP
	City *bool `json:"City,omitempty" name:"City"`

	// City district of the IP
	Region *bool `json:"Region,omitempty" name:"Region"`

	// Access ISP field
	Isp *bool `json:"Isp,omitempty" name:"Isp"`

	// ISP backbone network’s AS field
	AsName *bool `json:"AsName,omitempty" name:"AsName"`

	// Backbone AS ID
	AsId *bool `json:"AsId,omitempty" name:"AsId"`

	// Comment
	Comment *bool `json:"Comment,omitempty" name:"Comment"`
}

type IpGeolocationInfo struct {

	// Country/region
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Country *string `json:"Country,omitempty" name:"Country"`

	// Province- or municipality-level administrative region
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Province *string `json:"Province,omitempty" name:"Province"`

	// Municipal administrative region
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	City *string `json:"City,omitempty" name:"City"`

	// Urban area
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Region *string `json:"Region,omitempty" name:"Region"`

	// Access ISP
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Isp *string `json:"Isp,omitempty" name:"Isp"`

	// ISP backbone network’s AS name
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	AsName *string `json:"AsName,omitempty" name:"AsName"`

	// ISP backbone network’s AS ID
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	AsId *string `json:"AsId,omitempty" name:"AsId"`

	// Comment. The APN value of mobile users is entered currently. If there is no APN attribute, this is `null`.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// IP address
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	AddressIp *string `json:"AddressIp,omitempty" name:"AddressIp"`
}

type Ipv6Address struct {

	// `IPv6` address, such as `3402:4e00:20:100:0:8cd9:2a67:71f3`
	Address *string `json:"Address,omitempty" name:"Address"`

	// Whether it is a primary `IP`.
	Primary *bool `json:"Primary,omitempty" name:"Primary"`

	// The `ID` of the `EIP` instance, such as `eip-hxlqja90`.
	AddressId *string `json:"AddressId,omitempty" name:"AddressId"`

	// Message description
	Description *string `json:"Description,omitempty" name:"Description"`

	// Whether the public IP is blocked.
	IsWanIpBlocked *bool `json:"IsWanIpBlocked,omitempty" name:"IsWanIpBlocked"`

	// `IPv6` address status:
	// <li>`PENDING`: Creating</li>
	// <li>`MIGRATING`: Migrating</li>
	// <li>`DELETING`: Deleting</li>
	// <li>`AVAILABLE`: Available</li>
	State *string `json:"State,omitempty" name:"State"`
}

type Ipv6SubnetCidrBlock struct {

	// The `ID` of the subnet instance, such as `subnet-pxir56ns`.
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// The `IPv6` subnet IP range, such as `3402:4e00:20:1001::/64`
	Ipv6CidrBlock *string `json:"Ipv6CidrBlock,omitempty" name:"Ipv6CidrBlock"`
}

type ItemPrice struct {

	// The pay-as-you-go billing method. Unit: CNY.
	UnitPrice *float64 `json:"UnitPrice,omitempty" name:"UnitPrice"`

	// Pay-as-you-go billing method. Value Range: HOUR: Indicates billing by the hour. Scenarios using this hourly billing unit include: Instances postpaid on an hourly basis (POSTPAID_BY_HOUR), and bandwidth postpaid on an hourly basis (BANDWIDTH_POSTPAID_BY_HOUR). GB: Indicates billing on a per-GB basis. Scenarios using this billing unit include: Traffic postpaid on an hourly basis (TRAFFIC_POSTPAID_BY_HOUR).
	ChargeUnit *string `json:"ChargeUnit,omitempty" name:"ChargeUnit"`

	// Original price of the prepaid product. Unit: CNY.
	OriginalPrice *float64 `json:"OriginalPrice,omitempty" name:"OriginalPrice"`

	// Discount price of the prepaid product. Unit: CNY.
	DiscountPrice *float64 `json:"DiscountPrice,omitempty" name:"DiscountPrice"`
}

type MigrateNetworkInterfaceRequest struct {
	*tchttp.BaseRequest

	// The ID of the ENI instance, such as `eni-m6dyj72l`.
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`

	// The ID of the CVM bound to the ENI, such as `ins-r8hr2upy`.
	SourceInstanceId *string `json:"SourceInstanceId,omitempty" name:"SourceInstanceId"`

	// ID of the destination CVM instance to be migrated.
	DestinationInstanceId *string `json:"DestinationInstanceId,omitempty" name:"DestinationInstanceId"`
}

func (r *MigrateNetworkInterfaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *MigrateNetworkInterfaceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type MigrateNetworkInterfaceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *MigrateNetworkInterfaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *MigrateNetworkInterfaceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type MigratePrivateIpAddressRequest struct {
	*tchttp.BaseRequest

	// ID of the ENI instance bound with the private IP, such as `eni-m6dyj72l`.
	SourceNetworkInterfaceId *string `json:"SourceNetworkInterfaceId,omitempty" name:"SourceNetworkInterfaceId"`

	// ID of the destination ENI instance to be migrated.
	DestinationNetworkInterfaceId *string `json:"DestinationNetworkInterfaceId,omitempty" name:"DestinationNetworkInterfaceId"`

	// The private IP to be migrated, such as 10.0.0.6.
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" name:"PrivateIpAddress"`
}

func (r *MigratePrivateIpAddressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *MigratePrivateIpAddressRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type MigratePrivateIpAddressResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *MigratePrivateIpAddressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *MigratePrivateIpAddressResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAddressAttributeRequest struct {
	*tchttp.BaseRequest

	// The unique ID of the EIP, such as `eip-11112222`.
	AddressId *string `json:"AddressId,omitempty" name:"AddressId"`

	// The EIP name after modification. The maximum length is 20 characters.
	AddressName *string `json:"AddressName,omitempty" name:"AddressName"`

	// Whether the set EIP is a direct connection EIP. TRUE: yes. FALSE: no. Note that this parameter is available only to users who have activated the EIP direct connection function.
	EipDirectConnection *string `json:"EipDirectConnection,omitempty" name:"EipDirectConnection"`
}

func (r *ModifyAddressAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAddressAttributeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAddressAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAddressAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAddressAttributeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAddressInternetChargeTypeRequest struct {
	*tchttp.BaseRequest

	// Unique EIP ID, such as "eip-xxxx"
	AddressId *string `json:"AddressId,omitempty" name:"AddressId"`

	// The target billing method. It can be `BANDWIDTH_PREPAID_BY_MONTH` or `TRAFFIC_POSTPAID_BY_HOUR`
	InternetChargeType *string `json:"InternetChargeType,omitempty" name:"InternetChargeType"`

	// The target bandwidth value
	InternetMaxBandwidthOut *uint64 `json:"InternetMaxBandwidthOut,omitempty" name:"InternetMaxBandwidthOut"`

	// Billing details of monthly-subscribed network bandwidth. This parameter is required if the target billing method is `BANDWIDTH_PREPAID_BY_MONTH`.
	AddressChargePrepaid *AddressChargePrepaid `json:"AddressChargePrepaid,omitempty" name:"AddressChargePrepaid"`
}

func (r *ModifyAddressInternetChargeTypeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAddressInternetChargeTypeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAddressInternetChargeTypeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAddressInternetChargeTypeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAddressInternetChargeTypeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAddressTemplateAttributeRequest struct {
	*tchttp.BaseRequest

	// IP address template instance ID, such as `ipm-mdunqeb6`.
	AddressTemplateId *string `json:"AddressTemplateId,omitempty" name:"AddressTemplateId"`

	// IP address template name.
	AddressTemplateName *string `json:"AddressTemplateName,omitempty" name:"AddressTemplateName"`

	// Address information, including IP, CIDR and IP address range.
	Addresses []*string `json:"Addresses,omitempty" name:"Addresses" list`
}

func (r *ModifyAddressTemplateAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAddressTemplateAttributeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAddressTemplateAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAddressTemplateAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAddressTemplateAttributeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAddressTemplateGroupAttributeRequest struct {
	*tchttp.BaseRequest

	// IP address template group instance ID, such as `ipmg-2uw6ujo6`.
	AddressTemplateGroupId *string `json:"AddressTemplateGroupId,omitempty" name:"AddressTemplateGroupId"`

	// IP address template group name.
	AddressTemplateGroupName *string `json:"AddressTemplateGroupName,omitempty" name:"AddressTemplateGroupName"`

	// IP address template instance ID, such as `ipm-mdunqeb6`.
	AddressTemplateIds []*string `json:"AddressTemplateIds,omitempty" name:"AddressTemplateIds" list`
}

func (r *ModifyAddressTemplateGroupAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAddressTemplateGroupAttributeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAddressTemplateGroupAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAddressTemplateGroupAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAddressTemplateGroupAttributeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAddressesBandwidthRequest struct {
	*tchttp.BaseRequest

	// List of EIP IDs, such as “eip-xxxx”.
	AddressIds []*string `json:"AddressIds,omitempty" name:"AddressIds" list`

	// Target bandwidth value adjustment
	InternetMaxBandwidthOut *int64 `json:"InternetMaxBandwidthOut,omitempty" name:"InternetMaxBandwidthOut"`

	// (Disused) The start time of the monthly bandwidth subscription
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// (Disused) The end time of the monthly bandwidth subscription
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *ModifyAddressesBandwidthRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAddressesBandwidthRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAddressesBandwidthResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The async task ID. You can use the [DescribeTaskResult](https://intl.cloud.tencent.com/document/api/215/36271?from_cn_redirect=1) API to query the task status.
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAddressesBandwidthResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAddressesBandwidthResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAssistantCidrRequest struct {
	*tchttp.BaseRequest

	// `VPC` instance `ID`, e.g. `vpc-6v2ht8q5`.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Array of the secondary CIDR blocks to be added, such as ["10.0.0.0/16", "172.16.0.0/16"]. Either or both of `NewCidrBlocks` and `OldCidrBlocks` must be specified.
	NewCidrBlocks []*string `json:"NewCidrBlocks,omitempty" name:"NewCidrBlocks" list`

	// Array of the secondary CIDR blocks to be deleted, such as ["10.0.0.0/16", "172.16.0.0/16"]. Either or both of `NewCidrBlocks` and `OldCidrBlocks` must be specified.
	OldCidrBlocks []*string `json:"OldCidrBlocks,omitempty" name:"OldCidrBlocks" list`
}

func (r *ModifyAssistantCidrRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAssistantCidrRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAssistantCidrResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// A set of secondary CIDR blocks.
	// Note: This field may return null, indicating that no valid value was found.
		AssistantCidrSet []*AssistantCidr `json:"AssistantCidrSet,omitempty" name:"AssistantCidrSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAssistantCidrResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAssistantCidrResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyBandwidthPackageAttributeRequest struct {
	*tchttp.BaseRequest

	// The unique ID of the bandwidth package.
	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" name:"BandwidthPackageId"`

	// The name of the bandwidth package.
	BandwidthPackageName *string `json:"BandwidthPackageName,omitempty" name:"BandwidthPackageName"`

	// The billing mode of the bandwidth package.
	ChargeType *string `json:"ChargeType,omitempty" name:"ChargeType"`
}

func (r *ModifyBandwidthPackageAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyBandwidthPackageAttributeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyBandwidthPackageAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyBandwidthPackageAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyBandwidthPackageAttributeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyCcnAttributeRequest struct {
	*tchttp.BaseRequest

	// The CCN instance ID, such as `ccn-f49l6u0z`.
	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`

	// The name of the CCN. The maximum length is 60 characters.
	CcnName *string `json:"CcnName,omitempty" name:"CcnName"`

	// The description of the CCN. The maximum length is 100 characters.
	CcnDescription *string `json:"CcnDescription,omitempty" name:"CcnDescription"`
}

func (r *ModifyCcnAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyCcnAttributeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyCcnAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyCcnAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyCcnAttributeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyCcnRegionBandwidthLimitsTypeRequest struct {
	*tchttp.BaseRequest

	// CCN instance ID.
	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`

	// CCN bandwidth limit type. INTER_REGION_LIMIT: limit between regions. OUTER_REGION_LIMIT: region egress limit.
	BandwidthLimitType *string `json:"BandwidthLimitType,omitempty" name:"BandwidthLimitType"`
}

func (r *ModifyCcnRegionBandwidthLimitsTypeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyCcnRegionBandwidthLimitsTypeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyCcnRegionBandwidthLimitsTypeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyCcnRegionBandwidthLimitsTypeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyCcnRegionBandwidthLimitsTypeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyCustomerGatewayAttributeRequest struct {
	*tchttp.BaseRequest

	// The ID of the customer gateway, such as `cgw-2wqq41m9`. You can query the customer gateway by using the `DescribeCustomerGateways` API.
	CustomerGatewayId *string `json:"CustomerGatewayId,omitempty" name:"CustomerGatewayId"`

	// Customer gateway can be named freely, but the maximum length is 60 characters.
	CustomerGatewayName *string `json:"CustomerGatewayName,omitempty" name:"CustomerGatewayName"`
}

func (r *ModifyCustomerGatewayAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyCustomerGatewayAttributeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyCustomerGatewayAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyCustomerGatewayAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyCustomerGatewayAttributeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyDirectConnectGatewayAttributeRequest struct {
	*tchttp.BaseRequest

	// The unique ID of the direct connect gateway, such as `dcg-9o233uri`.
	DirectConnectGatewayId *string `json:"DirectConnectGatewayId,omitempty" name:"DirectConnectGatewayId"`

	// The direct connect gateway name. You can enter any name within 60 characters.
	DirectConnectGatewayName *string `json:"DirectConnectGatewayName,omitempty" name:"DirectConnectGatewayName"`

	// The CCN route-learning type. Valid values: `BGP` (Automatic learning), `STATIC` (Static, that is, user-configured). You can only modify `CcnRouteType` for a CCN direct connect gateway with BGP enabled.
	CcnRouteType *string `json:"CcnRouteType,omitempty" name:"CcnRouteType"`

	// CCN route publishing method. Valid values: `standard` and `exquisite`. You can only modify `ModeType` for a CCN direct connect gateway.
	ModeType *string `json:"ModeType,omitempty" name:"ModeType"`
}

func (r *ModifyDirectConnectGatewayAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDirectConnectGatewayAttributeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyDirectConnectGatewayAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDirectConnectGatewayAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDirectConnectGatewayAttributeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyFlowLogAttributeRequest struct {
	*tchttp.BaseRequest

	// ID of the VPC instance
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// The unique ID of the flow log.
	FlowLogId *string `json:"FlowLogId,omitempty" name:"FlowLogId"`

	// The name of the flow log.
	FlowLogName *string `json:"FlowLogName,omitempty" name:"FlowLogName"`

	// The description of the flow log.
	FlowLogDescription *string `json:"FlowLogDescription,omitempty" name:"FlowLogDescription"`
}

func (r *ModifyFlowLogAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyFlowLogAttributeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyFlowLogAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyFlowLogAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyFlowLogAttributeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyGatewayFlowQosRequest struct {
	*tchttp.BaseRequest

	// Gateway instance ID, which currently supports these types:
	// ID of Direct Connect gateway instance, e.g. `dcg-ltjahce6`;
	// ID of NAT gateway instance, e.g. `nat-ltjahce6`;
	// ID of VPN gateway instance, e.g. `vpn-ltjahce6`.
	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`

	// Bandwidth limit value.
	Bandwidth *int64 `json:"Bandwidth,omitempty" name:"Bandwidth"`

	// CVM private IP addresses with limited bandwidth.
	IpAddresses []*string `json:"IpAddresses,omitempty" name:"IpAddresses" list`
}

func (r *ModifyGatewayFlowQosRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyGatewayFlowQosRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyGatewayFlowQosResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyGatewayFlowQosResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyGatewayFlowQosResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyHaVipAttributeRequest struct {
	*tchttp.BaseRequest

	// The unique `ID` of the `HAVIP`, such as `havip-9o233uri`.
	HaVipId *string `json:"HaVipId,omitempty" name:"HaVipId"`

	// `HAVIP` can be named freely, but the maximum length is 60 characters.
	HaVipName *string `json:"HaVipName,omitempty" name:"HaVipName"`
}

func (r *ModifyHaVipAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyHaVipAttributeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyHaVipAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyHaVipAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyHaVipAttributeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyIpv6AddressesAttributeRequest struct {
	*tchttp.BaseRequest

	// The `ID` of the ENI instance, such as `eni-m6dyj72l`.
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`

	// The information of the specified private `IPv6` addresses.
	Ipv6Addresses []*Ipv6Address `json:"Ipv6Addresses,omitempty" name:"Ipv6Addresses" list`
}

func (r *ModifyIpv6AddressesAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyIpv6AddressesAttributeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyIpv6AddressesAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyIpv6AddressesAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyIpv6AddressesAttributeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyNatGatewayAttributeRequest struct {
	*tchttp.BaseRequest

	// The ID of the NAT gateway, such as `nat-df45454`.
	NatGatewayId *string `json:"NatGatewayId,omitempty" name:"NatGatewayId"`

	// The NAT gateway name, such as `test_nat`.
	NatGatewayName *string `json:"NatGatewayName,omitempty" name:"NatGatewayName"`

	// The maximum outbound bandwidth of the NAT gateway. Unit: Mbps.
	InternetMaxBandwidthOut *uint64 `json:"InternetMaxBandwidthOut,omitempty" name:"InternetMaxBandwidthOut"`

	// Whether to modify the security group bound to the NAT Gateway
	ModifySecurityGroup *bool `json:"ModifySecurityGroup,omitempty" name:"ModifySecurityGroup"`

	// The final security groups bound to the NAT Gateway, such as `['sg-1n232323', 'sg-o4242424']`. An empty list indicates that all the security groups have been deleted.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds" list`
}

func (r *ModifyNatGatewayAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyNatGatewayAttributeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyNatGatewayAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyNatGatewayAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyNatGatewayAttributeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyNatGatewayDestinationIpPortTranslationNatRuleRequest struct {
	*tchttp.BaseRequest

	// The ID of the NAT gateway, such as `nat-df45454`.
	NatGatewayId *string `json:"NatGatewayId,omitempty" name:"NatGatewayId"`

	// The port forwarding rule of the source NAT gateway.
	SourceNatRule *DestinationIpPortTranslationNatRule `json:"SourceNatRule,omitempty" name:"SourceNatRule"`

	// The port forwarding rule of the destination NAT gateway.
	DestinationNatRule *DestinationIpPortTranslationNatRule `json:"DestinationNatRule,omitempty" name:"DestinationNatRule"`
}

func (r *ModifyNatGatewayDestinationIpPortTranslationNatRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyNatGatewayDestinationIpPortTranslationNatRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyNatGatewayDestinationIpPortTranslationNatRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyNatGatewayDestinationIpPortTranslationNatRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyNatGatewayDestinationIpPortTranslationNatRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyNatGatewaySourceIpTranslationNatRuleRequest struct {
	*tchttp.BaseRequest

	// The ID of the NAT Gateway, such as `nat-df453454`
	NatGatewayId *string `json:"NatGatewayId,omitempty" name:"NatGatewayId"`

	// The SNAT forwarding rule of the NAT Gateway
	SourceIpTranslationNatRule *SourceIpTranslationNatRule `json:"SourceIpTranslationNatRule,omitempty" name:"SourceIpTranslationNatRule"`
}

func (r *ModifyNatGatewaySourceIpTranslationNatRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyNatGatewaySourceIpTranslationNatRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyNatGatewaySourceIpTranslationNatRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyNatGatewaySourceIpTranslationNatRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyNatGatewaySourceIpTranslationNatRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyNetDetectRequest struct {
	*tchttp.BaseRequest

	// The ID of a network detection instance, such as `netd-12345678`.
	NetDetectId *string `json:"NetDetectId,omitempty" name:"NetDetectId"`

	// The name of a network detection instance. The maximum length is 60 characters.
	NetDetectName *string `json:"NetDetectName,omitempty" name:"NetDetectName"`

	// The array of detection destination IPv4 addresses, which contains at most two IP addresses.
	DetectDestinationIp []*string `json:"DetectDestinationIp,omitempty" name:"DetectDestinationIp" list`

	// The type of the next hop. Currently supported types are:
	// VPN: VPN gateway;
	// DIRECTCONNECT: direct connect gateway;
	// PEERCONNECTION: peering connection;
	// NAT: NAT gateway;
	// NORMAL_CVM: normal CVM.
	NextHopType *string `json:"NextHopType,omitempty" name:"NextHopType"`

	// The next-hop destination gateway. The value is related to NextHopType.
	// If NextHopType is set to VPN, the value of this parameter is the VPN gateway ID, such as vpngw-12345678.
	// If NextHopType is set to DIRECTCONNECT, the value of this parameter is the direct connect gateway ID, such as dcg-12345678.
	// If NextHopType is set to PEERCONNECTION, the value of this parameter is the peering connection ID, such as pcx-12345678.
	// If NextHopType is set to NAT, the value of this parameter is the NAT gateway ID, such as nat-12345678.
	// If NextHopType is set to NORMAL_CVM, the value of this parameter is the IPv4 address of the CVM, such as 10.0.0.12.
	NextHopDestination *string `json:"NextHopDestination,omitempty" name:"NextHopDestination"`

	// Network detection description.
	NetDetectDescription *string `json:"NetDetectDescription,omitempty" name:"NetDetectDescription"`
}

func (r *ModifyNetDetectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyNetDetectRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyNetDetectResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyNetDetectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyNetDetectResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyNetworkAclAttributeRequest struct {
	*tchttp.BaseRequest

	// Network ACL instance ID. Example: acl-12345678.
	NetworkAclId *string `json:"NetworkAclId,omitempty" name:"NetworkAclId"`

	// Name of the network ACL. The maximum length is 60 bytes.
	NetworkAclName *string `json:"NetworkAclName,omitempty" name:"NetworkAclName"`
}

func (r *ModifyNetworkAclAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyNetworkAclAttributeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyNetworkAclAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyNetworkAclAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyNetworkAclAttributeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyNetworkAclEntriesRequest struct {
	*tchttp.BaseRequest

	// Network ACL instance ID. Example: acl-12345678.
	NetworkAclId *string `json:"NetworkAclId,omitempty" name:"NetworkAclId"`

	// Network ACL rule set.
	NetworkAclEntrySet *NetworkAclEntrySet `json:"NetworkAclEntrySet,omitempty" name:"NetworkAclEntrySet"`
}

func (r *ModifyNetworkAclEntriesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyNetworkAclEntriesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyNetworkAclEntriesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyNetworkAclEntriesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyNetworkAclEntriesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyNetworkInterfaceAttributeRequest struct {
	*tchttp.BaseRequest

	// The ID of the ENI instance, such as `eni-pxir56ns`.
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`

	// The name of the ENI. The maximum length is 60 characters.
	NetworkInterfaceName *string `json:"NetworkInterfaceName,omitempty" name:"NetworkInterfaceName"`

	// ENI description can be named freely, but the maximum length is 60 characters.
	NetworkInterfaceDescription *string `json:"NetworkInterfaceDescription,omitempty" name:"NetworkInterfaceDescription"`

	// The specified security groups to be bound with, such as ['sg-1dd51d'].
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds" list`
}

func (r *ModifyNetworkInterfaceAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyNetworkInterfaceAttributeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyNetworkInterfaceAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyNetworkInterfaceAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyNetworkInterfaceAttributeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyPrivateIpAddressesAttributeRequest struct {
	*tchttp.BaseRequest

	// The ID of the ENI instance, such as `eni-m6dyj72l`.
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`

	// The specified private IP information.
	PrivateIpAddresses []*PrivateIpAddressSpecification `json:"PrivateIpAddresses,omitempty" name:"PrivateIpAddresses" list`
}

func (r *ModifyPrivateIpAddressesAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyPrivateIpAddressesAttributeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyPrivateIpAddressesAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyPrivateIpAddressesAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyPrivateIpAddressesAttributeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyRouteTableAttributeRequest struct {
	*tchttp.BaseRequest

	// The route table instance ID, such as `rtb-azd4dt1c`.
	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`

	// Route table name.
	RouteTableName *string `json:"RouteTableName,omitempty" name:"RouteTableName"`
}

func (r *ModifyRouteTableAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyRouteTableAttributeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyRouteTableAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyRouteTableAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyRouteTableAttributeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifySecurityGroupAttributeRequest struct {
	*tchttp.BaseRequest

	// The security group instance ID, such as `sg-33ocnj9n`. This can be obtained through DescribeSecurityGroups.
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`

	// Security group can be named freely, but cannot exceed 60 characters.
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// The remarks for the security group. The maximum length is 100 characters.
	GroupDescription *string `json:"GroupDescription,omitempty" name:"GroupDescription"`
}

func (r *ModifySecurityGroupAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifySecurityGroupAttributeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifySecurityGroupAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySecurityGroupAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifySecurityGroupAttributeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifySecurityGroupPoliciesRequest struct {
	*tchttp.BaseRequest

	// The security group instance ID, such as `sg-33ocnj9n`. This can be obtained through DescribeSecurityGroups.
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`

	// The security group policy set. SecurityGroupPolicySet object must specify new egress and ingress policies at the same time. SecurityGroupPolicy object does not support custom index (PolicyIndex).
	SecurityGroupPolicySet *SecurityGroupPolicySet `json:"SecurityGroupPolicySet,omitempty" name:"SecurityGroupPolicySet"`

	// Whether the security group rule is sorted. Default value: False. If it is set to `True`, security group rules will be strictly sorted according to the sequence specified in the `SecurityGroupPolicySet` parameter. Manual entry may cause omission, so we recommend sorting security group rules in the console.
	SortPolicys *bool `json:"SortPolicys,omitempty" name:"SortPolicys"`
}

func (r *ModifySecurityGroupPoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifySecurityGroupPoliciesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifySecurityGroupPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySecurityGroupPoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifySecurityGroupPoliciesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyServiceTemplateAttributeRequest struct {
	*tchttp.BaseRequest

	// Protocol port template instance ID, such as `ppm-529nwwj8`.
	ServiceTemplateId *string `json:"ServiceTemplateId,omitempty" name:"ServiceTemplateId"`

	// Protocol port template name.
	ServiceTemplateName *string `json:"ServiceTemplateName,omitempty" name:"ServiceTemplateName"`

	// It supports single port, multiple ports, consecutive ports and all ports. Supported protocols include TCP, UDP, ICMP, and GRE.
	Services []*string `json:"Services,omitempty" name:"Services" list`
}

func (r *ModifyServiceTemplateAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyServiceTemplateAttributeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyServiceTemplateAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyServiceTemplateAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyServiceTemplateAttributeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyServiceTemplateGroupAttributeRequest struct {
	*tchttp.BaseRequest

	// The protocol port template group instance ID, such as `ppmg-ei8hfd9a`.
	ServiceTemplateGroupId *string `json:"ServiceTemplateGroupId,omitempty" name:"ServiceTemplateGroupId"`

	// Protocol port template group name.
	ServiceTemplateGroupName *string `json:"ServiceTemplateGroupName,omitempty" name:"ServiceTemplateGroupName"`

	// Instance ID of the protocol port template, such as `ppm-4dw6agho`.
	ServiceTemplateIds []*string `json:"ServiceTemplateIds,omitempty" name:"ServiceTemplateIds" list`
}

func (r *ModifyServiceTemplateGroupAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyServiceTemplateGroupAttributeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyServiceTemplateGroupAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyServiceTemplateGroupAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyServiceTemplateGroupAttributeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifySubnetAttributeRequest struct {
	*tchttp.BaseRequest

	// Subnet instance ID, such as `subnet-pxir56ns`.
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// The subnet name. The maximum length is 60 bytes.
	SubnetName *string `json:"SubnetName,omitempty" name:"SubnetName"`

	// Whether the subnet has broadcast enabled.
	EnableBroadcast *string `json:"EnableBroadcast,omitempty" name:"EnableBroadcast"`
}

func (r *ModifySubnetAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifySubnetAttributeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifySubnetAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySubnetAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifySubnetAttributeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyVpcAttributeRequest struct {
	*tchttp.BaseRequest

	// Security group can be named freely, but cannot exceed 60 characters.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// VPC can be named freely, but the maximum length is 60 characters.
	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`

	// Whether multicast is enabled. `true`: Enabled. `false`: Off.
	EnableMulticast *string `json:"EnableMulticast,omitempty" name:"EnableMulticast"`

	// DNS address. A maximum of 4 addresses is supported. The first one is primary server by default, and the rest are secondary servers.
	DnsServers []*string `json:"DnsServers,omitempty" name:"DnsServers" list`

	// Domain name
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`
}

func (r *ModifyVpcAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyVpcAttributeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyVpcAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyVpcAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyVpcAttributeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyVpnConnectionAttributeRequest struct {
	*tchttp.BaseRequest

	// The ID of the VPN tunnel instance, such as `vpnx-f49l6u0z`.
	VpnConnectionId *string `json:"VpnConnectionId,omitempty" name:"VpnConnectionId"`

	// VPN tunnel can be named freely, but the maximum length is 60 characters.
	VpnConnectionName *string `json:"VpnConnectionName,omitempty" name:"VpnConnectionName"`

	// The pre-shared key.
	PreShareKey *string `json:"PreShareKey,omitempty" name:"PreShareKey"`

	// The SPD policy group, for example: {"10.0.0.5/24":["172.123.10.5/16"]}. 10.0.0.5/24 is the VPC internal IP range, and 172.123.10.5/16 is the IDC IP range. The user specifies the IP range in the VPC that can communicate with the IP range in the IDC.
	SecurityPolicyDatabases []*SecurityPolicyDatabase `json:"SecurityPolicyDatabases,omitempty" name:"SecurityPolicyDatabases" list`

	// IKE (Internet Key Exchange) configuration. IKE comes with a self-protection mechanism. The network security protocol is configured by the user.
	IKEOptionsSpecification *IKEOptionsSpecification `json:"IKEOptionsSpecification,omitempty" name:"IKEOptionsSpecification"`

	// IPSec configuration. The IPSec secure session configuration is provided by Tencent Cloud.
	IPSECOptionsSpecification *IPSECOptionsSpecification `json:"IPSECOptionsSpecification,omitempty" name:"IPSECOptionsSpecification"`

	// Whether to enable the tunnel health check.
	EnableHealthCheck *bool `json:"EnableHealthCheck,omitempty" name:"EnableHealthCheck"`

	// Local IP address for the tunnel health check
	HealthCheckLocalIp *string `json:"HealthCheckLocalIp,omitempty" name:"HealthCheckLocalIp"`

	// Peer IP address for the tunnel health check
	HealthCheckRemoteIp *string `json:"HealthCheckRemoteIp,omitempty" name:"HealthCheckRemoteIp"`
}

func (r *ModifyVpnConnectionAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyVpnConnectionAttributeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyVpnConnectionAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyVpnConnectionAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyVpnConnectionAttributeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyVpnGatewayAttributeRequest struct {
	*tchttp.BaseRequest

	// The ID of the VPN gateway instance.
	VpnGatewayId *string `json:"VpnGatewayId,omitempty" name:"VpnGatewayId"`

	// The VPN gateway name. The maximum length is 60 bytes.
	VpnGatewayName *string `json:"VpnGatewayName,omitempty" name:"VpnGatewayName"`

	// VPN gateway billing mode. Currently, only the conversion of prepaid (monthly subscription) to postpaid (that is, pay-as-you-go) is supported. That is, the parameters only supports POSTPAID_BY_HOUR.
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`
}

func (r *ModifyVpnGatewayAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyVpnGatewayAttributeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyVpnGatewayAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyVpnGatewayAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyVpnGatewayAttributeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyVpnGatewayCcnRoutesRequest struct {
	*tchttp.BaseRequest

	// The ID of the VPN gateway instance.
	VpnGatewayId *string `json:"VpnGatewayId,omitempty" name:"VpnGatewayId"`

	// The CCN route (IDC IP range) list.
	Routes []*VpngwCcnRoutes `json:"Routes,omitempty" name:"Routes" list`
}

func (r *ModifyVpnGatewayCcnRoutesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyVpnGatewayCcnRoutesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyVpnGatewayCcnRoutesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyVpnGatewayCcnRoutesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyVpnGatewayCcnRoutesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type NatGateway struct {

	// NAT gateway ID.
	NatGatewayId *string `json:"NatGatewayId,omitempty" name:"NatGatewayId"`

	// NAT gateway name.
	NatGatewayName *string `json:"NatGatewayName,omitempty" name:"NatGatewayName"`

	// NAT gateway creation time.
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// The status of the NAT gateway.
	//  'PENDING': Creating, 'DELETING': Deleting, 'AVAILABLE': Operating, 'UPDATING': Upgrading,
	// ‘FAILED’: Failed.
	State *string `json:"State,omitempty" name:"State"`

	// The maximum outbound bandwidth of the gateway. Unit: Mbps.
	InternetMaxBandwidthOut *uint64 `json:"InternetMaxBandwidthOut,omitempty" name:"InternetMaxBandwidthOut"`

	// The concurrent connections cap of the gateway.
	MaxConcurrentConnection *uint64 `json:"MaxConcurrentConnection,omitempty" name:"MaxConcurrentConnection"`

	// The public IP object array of the bound NAT gateway.
	PublicIpAddressSet []*NatGatewayAddress `json:"PublicIpAddressSet,omitempty" name:"PublicIpAddressSet" list`

	// The NAT gateway status. `AVAILABLE`: Operating, `UNAVAILABLE`: Unavailable, `INSUFFICIENT`: Account is in arrears and the service is suspended.
	NetworkState *string `json:"NetworkState,omitempty" name:"NetworkState"`

	// The port forwarding rules of the NAT gateway.
	DestinationIpPortTranslationNatRuleSet []*DestinationIpPortTranslationNatRule `json:"DestinationIpPortTranslationNatRuleSet,omitempty" name:"DestinationIpPortTranslationNatRuleSet" list`

	// VPC instance ID.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// The availability zone in which the NAT gateway is located.
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// IDs of direct connect gateway associated.
	DirectConnectGatewayIds []*string `json:"DirectConnectGatewayIds,omitempty" name:"DirectConnectGatewayIds" list`

	// Subnet ID.
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// Tag key-value pair.
	TagSet []*Tag `json:"TagSet,omitempty" name:"TagSet" list`

	// The list of the security groups bound to the NAT Gateway
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	SecurityGroupSet []*string `json:"SecurityGroupSet,omitempty" name:"SecurityGroupSet" list`
}

type NatGatewayAddress struct {

	// The unique ID of the Elastic IP (EIP), such as `eip-11112222`.
	AddressId *string `json:"AddressId,omitempty" name:"AddressId"`

	// The public IP address, such as `123.121.34.33`.
	PublicIpAddress *string `json:"PublicIpAddress,omitempty" name:"PublicIpAddress"`

	// The block status of the resource. `true` indicates the EIP is blocked. `false` indicates that the EIP is not blocked.
	IsBlocked *bool `json:"IsBlocked,omitempty" name:"IsBlocked"`
}

type NatGatewayDestinationIpPortTranslationNatRule struct {

	// Network protocol. Available choices: `TCP`, `UDP`.
	IpProtocol *string `json:"IpProtocol,omitempty" name:"IpProtocol"`

	// EIP.
	PublicIpAddress *string `json:"PublicIpAddress,omitempty" name:"PublicIpAddress"`

	// Public port.
	PublicPort *uint64 `json:"PublicPort,omitempty" name:"PublicPort"`

	// Private network address.
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" name:"PrivateIpAddress"`

	// Private network port.
	PrivatePort *uint64 `json:"PrivatePort,omitempty" name:"PrivatePort"`

	// NAT gateway forwarding rule description.
	Description *string `json:"Description,omitempty" name:"Description"`

	// NAT gateway ID.
	// Note: This field may return null, indicating no valid value.
	NatGatewayId *string `json:"NatGatewayId,omitempty" name:"NatGatewayId"`

	// VPC ID.
	// Note: This field may return null, indicating no valid value.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// The creation time of the NAT gateway forwarding rule.
	// Note: This field may return null, indicating no valid value.
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
}

type NetDetect struct {

	// The ID of a VPC instance, such as `vpc-12345678`.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// The name of a VPC instance.
	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`

	// The ID of a subnet instance, such as subnet-12345678.
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// The name of a subnet instance.
	SubnetName *string `json:"SubnetName,omitempty" name:"SubnetName"`

	// The ID of a network detection instance, such as netd-12345678.
	NetDetectId *string `json:"NetDetectId,omitempty" name:"NetDetectId"`

	// The name of a network detection instance. The maximum length is 60 characters.
	NetDetectName *string `json:"NetDetectName,omitempty" name:"NetDetectName"`

	// The array of detection destination IPv4 addresses, which contains at most two IP addresses.
	DetectDestinationIp []*string `json:"DetectDestinationIp,omitempty" name:"DetectDestinationIp" list`

	// The array of detection source IPv4 addresses automatically allocated by the system. The length is 2.
	DetectSourceIp []*string `json:"DetectSourceIp,omitempty" name:"DetectSourceIp" list`

	// Type of the next hop. Currently supported types are:
	// VPN: VPN gateway;
	// DIRECTCONNECT: direct connect gateway;
	// PEERCONNECTION: peering connection;
	// NAT: NAT gateway;
	// NORMAL_CVM: normal CVM.
	NextHopType *string `json:"NextHopType,omitempty" name:"NextHopType"`

	// Next-hop destination gateway. The value is related to NextHopType.
	// If NextHopType is set to VPN, the value of this parameter is the VPN gateway ID, such as vpngw-12345678.
	// If NextHopType is set to DIRECTCONNECT, the value of this parameter is the direct connect gateway ID, such as dcg-12345678.
	// If NextHopType is set to PEERCONNECTION, the value of this parameter is the peering connection ID, such as pcx-12345678.
	// If NextHopType is set to NAT, the value of this parameter is the NAT gateway ID, such as nat-12345678.
	// If NextHopType is set to NORMAL_CVM, the value of this parameter is the IPv4 address of the CVM, such as 10.0.0.12.
	NextHopDestination *string `json:"NextHopDestination,omitempty" name:"NextHopDestination"`

	// The name of the next-hop gateway.
	// Note: This field may return null, indicating that no valid values can be obtained.
	NextHopName *string `json:"NextHopName,omitempty" name:"NextHopName"`

	// Network detection description.
	// Note: This field may return null, indicating that no valid values can be obtained.
	NetDetectDescription *string `json:"NetDetectDescription,omitempty" name:"NetDetectDescription"`

	// The creation time.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type NetDetectIpState struct {

	// The destination IPv4 address of network detection.
	DetectDestinationIp *string `json:"DetectDestinationIp,omitempty" name:"DetectDestinationIp"`

	// The detection result.
	// 0: successful;
	// -1: no packet loss occurred during routing;
	// -2: packet loss occurred when outbound traffic is blocked by the ACL;
	// -3: packet loss occurred when inbound traffic is blocked by the ACL;
	// -4: other errors.
	State *int64 `json:"State,omitempty" name:"State"`

	// The latency. Unit: ms.
	Delay *uint64 `json:"Delay,omitempty" name:"Delay"`

	// The packet loss rate.
	PacketLossRate *uint64 `json:"PacketLossRate,omitempty" name:"PacketLossRate"`
}

type NetDetectState struct {

	// The ID of a network detection instance, such as netd-12345678.
	NetDetectId *string `json:"NetDetectId,omitempty" name:"NetDetectId"`

	// The array of network detection destination IP verification results.
	NetDetectIpStateSet []*NetDetectIpState `json:"NetDetectIpStateSet,omitempty" name:"NetDetectIpStateSet" list`
}

type NetworkAcl struct {

	// `ID` of the `VPC` instance.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// `ID` of the network ACL instance.
	NetworkAclId *string `json:"NetworkAclId,omitempty" name:"NetworkAclId"`

	// Name of the network ACL. The maximum length is 60 bytes.
	NetworkAclName *string `json:"NetworkAclName,omitempty" name:"NetworkAclName"`

	// Creation time.
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// Array of subnets associated with the network ACL.
	SubnetSet []*Subnet `json:"SubnetSet,omitempty" name:"SubnetSet" list`

	// Inbound rules of the network ACL.
	IngressEntries []*NetworkAclEntry `json:"IngressEntries,omitempty" name:"IngressEntries" list`

	// Outbound rules of the network ACL.
	EgressEntries []*NetworkAclEntry `json:"EgressEntries,omitempty" name:"EgressEntries" list`
}

type NetworkAclEntry struct {

	// Modification time.
	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`

	// Protocol. Valid values: TCP, UDP, ICMP, ALL.
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Port. Valid values: all, single port, range. When Protocol takes the value `ALL` or `ICMP`, Port cannot be specified.
	Port *string `json:"Port,omitempty" name:"Port"`

	// IP range or IP address (mutually exclusive).
	CidrBlock *string `json:"CidrBlock,omitempty" name:"CidrBlock"`

	// CIDR block or IPv6 address (mutually exclusive).
	Ipv6CidrBlock *string `json:"Ipv6CidrBlock,omitempty" name:"Ipv6CidrBlock"`

	// ACCEPT or DROP.
	Action *string `json:"Action,omitempty" name:"Action"`

	// Rule description, which is up to 100 bytes.
	Description *string `json:"Description,omitempty" name:"Description"`
}

type NetworkAclEntrySet struct {

	// Inbound rules.
	Ingress []*NetworkAclEntry `json:"Ingress,omitempty" name:"Ingress" list`

	// Outbound rules.
	Egress []*NetworkAclEntry `json:"Egress,omitempty" name:"Egress" list`
}

type NetworkInterface struct {

	// The ID of the ENI instance, such as `eni-f1xjkw1b`.
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`

	// ENI Name
	NetworkInterfaceName *string `json:"NetworkInterfaceName,omitempty" name:"NetworkInterfaceName"`

	// ENI description.
	NetworkInterfaceDescription *string `json:"NetworkInterfaceDescription,omitempty" name:"NetworkInterfaceDescription"`

	// Subnet instance ID.
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// VPC instance ID.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Bound security group.
	GroupSet []*string `json:"GroupSet,omitempty" name:"GroupSet" list`

	// Whether it is the primary ENI.
	Primary *bool `json:"Primary,omitempty" name:"Primary"`

	// MAC address
	MacAddress *string `json:"MacAddress,omitempty" name:"MacAddress"`

	// ENI status:
	// <li>`PENDING`: Creating</li>
	// <li>`AVAILABLE`: Available</li>
	// <li>`ATTACHING`: Binding</li>
	// <li>`DETACHING`: Unbinding</li>
	// <li>`DELETING`: Deleting</li>
	State *string `json:"State,omitempty" name:"State"`

	// Private IP information.
	PrivateIpAddressSet []*PrivateIpAddressSpecification `json:"PrivateIpAddressSet,omitempty" name:"PrivateIpAddressSet" list`

	// Bound CVM object.
	// Note: This field may return null, indicating no valid value.
	Attachment *NetworkInterfaceAttachment `json:"Attachment,omitempty" name:"Attachment"`

	// Availability Zone.
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// Creation Time.
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// The `IPv6` address list.
	Ipv6AddressSet []*Ipv6Address `json:"Ipv6AddressSet,omitempty" name:"Ipv6AddressSet" list`

	// Tag key-value pair.
	TagSet []*Tag `json:"TagSet,omitempty" name:"TagSet" list`

	// The ENI type. 0: ENI. 1: EVM ENI.
	EniType *uint64 `json:"EniType,omitempty" name:"EniType"`

	// Type of the resource bound with an ENI. Valid values: cvm, eks.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Business *string `json:"Business,omitempty" name:"Business"`
}

type NetworkInterfaceAttachment struct {

	// CVM instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// The serial number of ENI in the CVM instance.
	DeviceIndex *uint64 `json:"DeviceIndex,omitempty" name:"DeviceIndex"`

	// The account information of the CVM owner.
	InstanceAccountId *string `json:"InstanceAccountId,omitempty" name:"InstanceAccountId"`

	// Binding time
	AttachTime *string `json:"AttachTime,omitempty" name:"AttachTime"`
}

type NotifyRoutesRequest struct {
	*tchttp.BaseRequest

	// The unique ID of the route table
	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`

	// The unique ID of the routing policy
	RouteItemIds []*string `json:"RouteItemIds,omitempty" name:"RouteItemIds" list`
}

func (r *NotifyRoutesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *NotifyRoutesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type NotifyRoutesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *NotifyRoutesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *NotifyRoutesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Price struct {

	// Instance price.
	InstancePrice *ItemPrice `json:"InstancePrice,omitempty" name:"InstancePrice"`

	// Network price.
	BandwidthPrice *ItemPrice `json:"BandwidthPrice,omitempty" name:"BandwidthPrice"`
}

type PrivateIpAddressSpecification struct {

	// Private IP address.
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" name:"PrivateIpAddress"`

	// Whether it is a primary IP.
	Primary *bool `json:"Primary,omitempty" name:"Primary"`

	// Public IP address.
	PublicIpAddress *string `json:"PublicIpAddress,omitempty" name:"PublicIpAddress"`

	// EIP instance ID, such as `eip-11112222`.
	AddressId *string `json:"AddressId,omitempty" name:"AddressId"`

	// Private IP description.
	Description *string `json:"Description,omitempty" name:"Description"`

	// Whether the public IP is blocked.
	IsWanIpBlocked *bool `json:"IsWanIpBlocked,omitempty" name:"IsWanIpBlocked"`

	// IP status:
	// PENDING: Creating
	// MIGRATING: Migrating
	// DELETING: Deleting
	// AVAILABLE: Available
	State *string `json:"State,omitempty" name:"State"`
}

type Quota struct {

	// Quota name. Value range:<br><li>`TOTAL_EIP_QUOTA`:EIP quota under the user's current region<br><li>`DAILY_EIP_APPLY`: Number of EIP applications submitted daily under the user's current region<br><li>`DAILY_PUBLIC_IP_ASSIGN`: Number of public IP reassignments under the user's current region.
	QuotaId *string `json:"QuotaId,omitempty" name:"QuotaId"`

	// Current count
	QuotaCurrent *int64 `json:"QuotaCurrent,omitempty" name:"QuotaCurrent"`

	// Quota
	QuotaLimit *int64 `json:"QuotaLimit,omitempty" name:"QuotaLimit"`
}

type ReferredSecurityGroup struct {

	// Security group instance ID.
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`

	// IDs of all referred security group instances.
	ReferredSecurityGroupIds []*string `json:"ReferredSecurityGroupIds,omitempty" name:"ReferredSecurityGroupIds" list`
}

type RejectAttachCcnInstancesRequest struct {
	*tchttp.BaseRequest

	// The CCN instance ID, such as `ccn-f49l6u0z`.
	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`

	// The list of instances whose association is rejected.
	Instances []*CcnInstance `json:"Instances,omitempty" name:"Instances" list`
}

func (r *RejectAttachCcnInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RejectAttachCcnInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RejectAttachCcnInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RejectAttachCcnInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RejectAttachCcnInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ReleaseAddressesRequest struct {
	*tchttp.BaseRequest

	// The unique ID list of the EIP. The unique ID of an EIP is as follows: `eip-11112222`.
	AddressIds []*string `json:"AddressIds,omitempty" name:"AddressIds" list`
}

func (r *ReleaseAddressesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ReleaseAddressesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ReleaseAddressesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The async task ID. You can use the [DescribeTaskResult](https://intl.cloud.tencent.com/document/api/215/36271?from_cn_redirect=1) API to query the task status.
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReleaseAddressesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ReleaseAddressesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RemoveBandwidthPackageResourcesRequest struct {
	*tchttp.BaseRequest

	// The unique ID of the bandwidth package, such as `bwp-xxxx`.
	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" name:"BandwidthPackageId"`

	// The resource type. Valid values: `Address` and `LoadBalance`.
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`

	// The resource IP, such as `eip-xxxx` and `lb-xxxx`.
	ResourceIds []*string `json:"ResourceIds,omitempty" name:"ResourceIds" list`
}

func (r *RemoveBandwidthPackageResourcesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RemoveBandwidthPackageResourcesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RemoveBandwidthPackageResourcesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RemoveBandwidthPackageResourcesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RemoveBandwidthPackageResourcesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RenewVpnGatewayRequest struct {
	*tchttp.BaseRequest

	// The ID of the VPN gateway instance.
	VpnGatewayId *string `json:"VpnGatewayId,omitempty" name:"VpnGatewayId"`

	// Billing Methods
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitempty" name:"InstanceChargePrepaid"`
}

func (r *RenewVpnGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RenewVpnGatewayRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RenewVpnGatewayResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RenewVpnGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RenewVpnGatewayResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ReplaceDirectConnectGatewayCcnRoutesRequest struct {
	*tchttp.BaseRequest

	// The ID of the Direct Connect gateway, such as `dcg-prpqlmg1`
	DirectConnectGatewayId *string `json:"DirectConnectGatewayId,omitempty" name:"DirectConnectGatewayId"`

	// The list of IDC IP range that must be connected
	Routes []*DirectConnectGatewayCcnRoute `json:"Routes,omitempty" name:"Routes" list`
}

func (r *ReplaceDirectConnectGatewayCcnRoutesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ReplaceDirectConnectGatewayCcnRoutesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ReplaceDirectConnectGatewayCcnRoutesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReplaceDirectConnectGatewayCcnRoutesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ReplaceDirectConnectGatewayCcnRoutesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ReplaceRouteTableAssociationRequest struct {
	*tchttp.BaseRequest

	// Subnet instance ID, such as `subnet-3x5lf5q0`. This can be queried using the DescribeSubnets API.
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// The route table instance ID, such as `rtb-azd4dt1c`.
	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`
}

func (r *ReplaceRouteTableAssociationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ReplaceRouteTableAssociationRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ReplaceRouteTableAssociationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReplaceRouteTableAssociationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ReplaceRouteTableAssociationResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ReplaceRoutesRequest struct {
	*tchttp.BaseRequest

	// The route table instance ID, such as `rtb-azd4dt1c`.
	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`

	// Routing policy object. The routing policy ID (RouteId) must be specified.
	Routes []*Route `json:"Routes,omitempty" name:"Routes" list`
}

func (r *ReplaceRoutesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ReplaceRoutesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ReplaceRoutesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReplaceRoutesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ReplaceRoutesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ReplaceSecurityGroupPolicyRequest struct {
	*tchttp.BaseRequest

	// The security group instance ID, such as `sg-33ocnj9n`. This can be obtained through DescribeSecurityGroups.
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`

	// Security group policy set object.
	SecurityGroupPolicySet *SecurityGroupPolicySet `json:"SecurityGroupPolicySet,omitempty" name:"SecurityGroupPolicySet"`

	// (Optional) The old policy set of the security group, which is used for log records.
	OriginalSecurityGroupPolicySet *SecurityGroupPolicySet `json:"OriginalSecurityGroupPolicySet,omitempty" name:"OriginalSecurityGroupPolicySet"`
}

func (r *ReplaceSecurityGroupPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ReplaceSecurityGroupPolicyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ReplaceSecurityGroupPolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReplaceSecurityGroupPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ReplaceSecurityGroupPolicyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResetAttachCcnInstancesRequest struct {
	*tchttp.BaseRequest

	// The CCN instance ID, such as `ccn-f49l6u0z`.
	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`

	// The UIN (root account) to which the CCN belongs.
	CcnUin *string `json:"CcnUin,omitempty" name:"CcnUin"`

	// The list of network instances that re-apply for association.
	Instances []*CcnInstance `json:"Instances,omitempty" name:"Instances" list`
}

func (r *ResetAttachCcnInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResetAttachCcnInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResetAttachCcnInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResetAttachCcnInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResetAttachCcnInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResetNatGatewayConnectionRequest struct {
	*tchttp.BaseRequest

	// NAT gateway ID.
	NatGatewayId *string `json:"NatGatewayId,omitempty" name:"NatGatewayId"`

	// Concurrent connections cap of the NAT gateway, such as 1000000, 3000000, 10000000.
	MaxConcurrentConnection *uint64 `json:"MaxConcurrentConnection,omitempty" name:"MaxConcurrentConnection"`
}

func (r *ResetNatGatewayConnectionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResetNatGatewayConnectionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResetNatGatewayConnectionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResetNatGatewayConnectionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResetNatGatewayConnectionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResetRoutesRequest struct {
	*tchttp.BaseRequest

	// The route table instance ID, such as `rtb-azd4dt1c`.
	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`

	// The route table name. The maximum length is 60 characters.
	RouteTableName *string `json:"RouteTableName,omitempty" name:"RouteTableName"`

	// Routing policy.
	Routes []*Route `json:"Routes,omitempty" name:"Routes" list`
}

func (r *ResetRoutesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResetRoutesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResetRoutesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResetRoutesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResetRoutesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResetVpnConnectionRequest struct {
	*tchttp.BaseRequest

	// The ID of the VPN gateway instance.
	VpnGatewayId *string `json:"VpnGatewayId,omitempty" name:"VpnGatewayId"`

	// The ID of the VPN tunnel instance, such as `vpnx-f49l6u0z`.
	VpnConnectionId *string `json:"VpnConnectionId,omitempty" name:"VpnConnectionId"`
}

func (r *ResetVpnConnectionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResetVpnConnectionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResetVpnConnectionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResetVpnConnectionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResetVpnConnectionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResetVpnGatewayInternetMaxBandwidthRequest struct {
	*tchttp.BaseRequest

	// The ID of the VPN gateway instance.
	VpnGatewayId *string `json:"VpnGatewayId,omitempty" name:"VpnGatewayId"`

	// The public network bandwidth configuration. Available bandwidth specifications: 5, 10, 20, 50, and 100. Unit: Mbps.
	InternetMaxBandwidthOut *uint64 `json:"InternetMaxBandwidthOut,omitempty" name:"InternetMaxBandwidthOut"`
}

func (r *ResetVpnGatewayInternetMaxBandwidthRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResetVpnGatewayInternetMaxBandwidthRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResetVpnGatewayInternetMaxBandwidthResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResetVpnGatewayInternetMaxBandwidthResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResetVpnGatewayInternetMaxBandwidthResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Resource struct {

	// The bandwidth package resource type, including `Address`, and `LoadBalance`
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`

	// The bandwidth package ID, such as `eip-xxxx` and `lb-xxxx`.
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// The bandwidth package resource IP.
	AddressIp *string `json:"AddressIp,omitempty" name:"AddressIp"`
}

type ResourceDashboard struct {

	// Vpc instance ID, e.g. vpc-f1xjkw1b.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Subnet instance ID, such as subnet-bthucmmy.
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// Classiclink.
	Classiclink *uint64 `json:"Classiclink,omitempty" name:"Classiclink"`

	// Direct Connect gateway.
	Dcg *uint64 `json:"Dcg,omitempty" name:"Dcg"`

	// Peering connection.
	Pcx *uint64 `json:"Pcx,omitempty" name:"Pcx"`

	// The total number of used IP addresses.
	Ip *uint64 `json:"Ip,omitempty" name:"Ip"`

	// NAT gateway.
	Nat *uint64 `json:"Nat,omitempty" name:"Nat"`

	// VPN gateway.
	Vpngw *uint64 `json:"Vpngw,omitempty" name:"Vpngw"`

	// Flow log.
	FlowLog *uint64 `json:"FlowLog,omitempty" name:"FlowLog"`

	// Network probing.
	NetworkDetect *uint64 `json:"NetworkDetect,omitempty" name:"NetworkDetect"`

	// Network ACL.
	NetworkACL *uint64 `json:"NetworkACL,omitempty" name:"NetworkACL"`

	// Cloud Virtual Machine.
	CVM *uint64 `json:"CVM,omitempty" name:"CVM"`

	// Load balancer.
	LB *uint64 `json:"LB,omitempty" name:"LB"`

	// Relational database.
	CDB *uint64 `json:"CDB,omitempty" name:"CDB"`

	// TencentDB for Memcached.
	Cmem *uint64 `json:"Cmem,omitempty" name:"Cmem"`

	// Cloud time series database.
	CTSDB *uint64 `json:"CTSDB,omitempty" name:"CTSDB"`

	// TencentDB for MariaDB (TDSQL).
	MariaDB *uint64 `json:"MariaDB,omitempty" name:"MariaDB"`

	// TencentDB for SQL Server.
	SQLServer *uint64 `json:"SQLServer,omitempty" name:"SQLServer"`

	// TencentDB for PostgreSQL.
	Postgres *uint64 `json:"Postgres,omitempty" name:"Postgres"`

	// Network attached storage.
	NAS *uint64 `json:"NAS,omitempty" name:"NAS"`

	// Snova data warehouse.
	Greenplumn *uint64 `json:"Greenplumn,omitempty" name:"Greenplumn"`

	// Cloud Kafka (CKafka).
	Ckafka *uint64 `json:"Ckafka,omitempty" name:"Ckafka"`

	// Grocery.
	Grocery *uint64 `json:"Grocery,omitempty" name:"Grocery"`

	// Data encryption service.
	HSM *uint64 `json:"HSM,omitempty" name:"HSM"`

	// Game storage - Tcaplus.
	Tcaplus *uint64 `json:"Tcaplus,omitempty" name:"Tcaplus"`

	// Cnas.
	Cnas *uint64 `json:"Cnas,omitempty" name:"Cnas"`

	// HTAP database - TiDB.
	TiDB *uint64 `json:"TiDB,omitempty" name:"TiDB"`

	// EMR cluster.
	Emr *uint64 `json:"Emr,omitempty" name:"Emr"`

	// SEAL.
	SEAL *uint64 `json:"SEAL,omitempty" name:"SEAL"`

	// Cloud file storage - CFS.
	CFS *uint64 `json:"CFS,omitempty" name:"CFS"`

	// Oracle.
	Oracle *uint64 `json:"Oracle,omitempty" name:"Oracle"`

	// ElasticSearch Service.
	ElasticSearch *uint64 `json:"ElasticSearch,omitempty" name:"ElasticSearch"`

	// Blockchain service.
	TBaaS *uint64 `json:"TBaaS,omitempty" name:"TBaaS"`

	// Itop.
	Itop *uint64 `json:"Itop,omitempty" name:"Itop"`

	// Cloud database audit.
	DBAudit *uint64 `json:"DBAudit,omitempty" name:"DBAudit"`

	// Enterprise TencentDB - CynosDB for Postgres.
	CynosDBPostgres *uint64 `json:"CynosDBPostgres,omitempty" name:"CynosDBPostgres"`

	// TencentDB for Redis.
	Redis *uint64 `json:"Redis,omitempty" name:"Redis"`

	// TencentDB for MongoDB.
	MongoDB *uint64 `json:"MongoDB,omitempty" name:"MongoDB"`

	// A distributed cloud database - TencentDB for TDSQL.
	DCDB *uint64 `json:"DCDB,omitempty" name:"DCDB"`

	// An enterprise-grade TencentDB - CynosDB for MySQL.
	CynosDBMySQL *uint64 `json:"CynosDBMySQL,omitempty" name:"CynosDBMySQL"`

	// Subnets.
	Subnet *uint64 `json:"Subnet,omitempty" name:"Subnet"`

	// Route table.
	RouteTable *uint64 `json:"RouteTable,omitempty" name:"RouteTable"`
}

type Route struct {

	// Destination IP range, such as 112.20.51.0/24. Values cannot be in the VPC IP range.
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" name:"DestinationCidrBlock"`

	// Type of the next hop. Currently supported types are:
	// CVM: CVM of the public gateway type;
	// VPN: VPN gateway;
	// DIRECTCONNECT: direct connect gateway;
	// PEERCONNECTION: peering connection;
	// SSLVPN: sslvpn gateway;
	// NAT: NAT gateway; 
	// NORMAL_CVM: normal CVM;
	// EIP: public IP address of the CVM;
	// CCN: Cloud Connect Network.
	GatewayType *string `json:"GatewayType,omitempty" name:"GatewayType"`

	// Next hop address. You simply need to specify the gateway ID of a different next hop type, and the system will automatically match the next hop address.
	// Important note: When the GatewayType is EIP, the GatewayId has a fixed value `0`
	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`

	// Routing policy ID. The IPv4 routing policy will have a meaningful value, while the IPv6 routing policy is always 0. We recommend using the unique ID `RouteItemId` for the routing policy.
	RouteId *uint64 `json:"RouteId,omitempty" name:"RouteId"`

	// The description of the routing policy.
	RouteDescription *string `json:"RouteDescription,omitempty" name:"RouteDescription"`

	// Whether it is enabled
	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`

	// The route type. Currently, the following types are supported:
	// USER: User route;
	// NETD: Network probe route. When creating a network probe route, the system delivers by default. It cannot be edited or deleted;
	// CCN: CCN route. The system delivers by default. It cannot be edited or deleted.
	// Users can only add and operate USER-type routes.
	RouteType *string `json:"RouteType,omitempty" name:"RouteType"`

	// Route table instance ID, such as rtb-azd4dt1c.
	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`

	// Destination IPv6 IP range, which cannot be included in VPC IP range, such as 2402:4e00:1000:810b::/64.
	DestinationIpv6CidrBlock *string `json:"DestinationIpv6CidrBlock,omitempty" name:"DestinationIpv6CidrBlock"`

	// Unique routing policy ID.
	RouteItemId *string `json:"RouteItemId,omitempty" name:"RouteItemId"`

	// Whether the routing policy is published to CCN.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	PublishedToVbc *bool `json:"PublishedToVbc,omitempty" name:"PublishedToVbc"`
}

type RouteTable struct {

	// VPC instance ID.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// The route table instance ID, such as `rtb-azd4dt1c`.
	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`

	// Route table name.
	RouteTableName *string `json:"RouteTableName,omitempty" name:"RouteTableName"`

	// The association relationships of the route table.
	AssociationSet []*RouteTableAssociation `json:"AssociationSet,omitempty" name:"AssociationSet" list`

	// IPv4 routing policy set.
	RouteSet []*Route `json:"RouteSet,omitempty" name:"RouteSet" list`

	// Whether it is the default route table.
	Main *bool `json:"Main,omitempty" name:"Main"`

	// Creation Time.
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// Tag key-value pairs.
	TagSet []*Tag `json:"TagSet,omitempty" name:"TagSet" list`

	// Whether the local route is published to CCN.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	LocalCidrForCcn []*CidrForCcn `json:"LocalCidrForCcn,omitempty" name:"LocalCidrForCcn" list`
}

type RouteTableAssociation struct {

	// Subnet instance ID.
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// Route table instance ID.
	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`
}

type SecurityGroup struct {

	// The security group instance ID, such as `sg-ohuuioma`.
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`

	// Security group can be named freely, but cannot exceed 60 characters.
	SecurityGroupName *string `json:"SecurityGroupName,omitempty" name:"SecurityGroupName"`

	// The remarks for the security group. The maximum length is 100 characters.
	SecurityGroupDesc *string `json:"SecurityGroupDesc,omitempty" name:"SecurityGroupDesc"`

	// The project id is 0 by default. You can query this in the project management page of the Qcloud console.
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// Whether it is the default security group (which cannot be deleted).
	IsDefault *bool `json:"IsDefault,omitempty" name:"IsDefault"`

	// Security group creation time.
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// Tag key-value pairs.
	TagSet []*Tag `json:"TagSet,omitempty" name:"TagSet" list`
}

type SecurityGroupAssociationStatistics struct {

	// Security group instance ID.
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`

	// Number of CVM instances.
	CVM *uint64 `json:"CVM,omitempty" name:"CVM"`

	// Number of TencentDB for MySQL instances
	CDB *uint64 `json:"CDB,omitempty" name:"CDB"`

	// Number of ENI instances.
	ENI *uint64 `json:"ENI,omitempty" name:"ENI"`

	// Number of times a security group is referenced by other security groups
	SG *uint64 `json:"SG,omitempty" name:"SG"`

	// Number of load balancer instances.
	CLB *uint64 `json:"CLB,omitempty" name:"CLB"`

	// The binding statistics for all instances.
	InstanceStatistics []*InstanceStatistic `json:"InstanceStatistics,omitempty" name:"InstanceStatistics" list`

	// Total count of all resources (excluding resources referenced by security groups).
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

type SecurityGroupPolicy struct {

	// The index number of security group rules, which dynamically changes with the rules. This parameter can be obtained via the `DescribeSecurityGroupPolicies` API and used with the `Version` field in the returned value of the API.
	PolicyIndex *int64 `json:"PolicyIndex,omitempty" name:"PolicyIndex"`

	// Protocol. Valid values: TCP, UDP, ICMP, ICMPv6, ALL.
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Port (all, discrete port, range).
	Port *string `json:"Port,omitempty" name:"Port"`

	// Protocol port ID or protocol port group ID. ServiceTemplate and Protocol+Port are mutually exclusive.
	ServiceTemplate *ServiceTemplateSpecification `json:"ServiceTemplate,omitempty" name:"ServiceTemplate"`

	// IP range or IP (mutually exclusive).
	CidrBlock *string `json:"CidrBlock,omitempty" name:"CidrBlock"`

	// The CIDR block or IPv6 (mutually exclusive).
	Ipv6CidrBlock *string `json:"Ipv6CidrBlock,omitempty" name:"Ipv6CidrBlock"`

	// The security group instance ID, such as `sg-ohuuioma`.
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`

	// IP address ID or IP address group ID.
	AddressTemplate *AddressTemplateSpecification `json:"AddressTemplate,omitempty" name:"AddressTemplate"`

	// ACCEPT or DROP.
	Action *string `json:"Action,omitempty" name:"Action"`

	// Security group policy description.
	PolicyDescription *string `json:"PolicyDescription,omitempty" name:"PolicyDescription"`

	// The last modification time of the security group.
	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
}

type SecurityGroupPolicySet struct {

	// The version of the security group policy. The version number is automatically increased by one each time users update the security policy, to prevent the expiration of updated routing policies. Conflict is ignored if it is left empty.
	Version *string `json:"Version,omitempty" name:"Version"`

	// Outbound policy.
	Egress []*SecurityGroupPolicy `json:"Egress,omitempty" name:"Egress" list`

	// Inbound policy.
	Ingress []*SecurityGroupPolicy `json:"Ingress,omitempty" name:"Ingress" list`
}

type SecurityPolicyDatabase struct {

	// Local IP range
	LocalCidrBlock *string `json:"LocalCidrBlock,omitempty" name:"LocalCidrBlock"`

	// Opposite IP range
	RemoteCidrBlock []*string `json:"RemoteCidrBlock,omitempty" name:"RemoteCidrBlock" list`
}

type ServiceTemplate struct {

	// Protocol port instance ID, such as `ppm-f5n1f8da`.
	ServiceTemplateId *string `json:"ServiceTemplateId,omitempty" name:"ServiceTemplateId"`

	// Template name.
	ServiceTemplateName *string `json:"ServiceTemplateName,omitempty" name:"ServiceTemplateName"`

	// Protocol port information.
	ServiceSet []*string `json:"ServiceSet,omitempty" name:"ServiceSet" list`

	// Creation Time.
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
}

type ServiceTemplateGroup struct {

	// Protocol port template group instance ID, such as `ppmg-2klmrefu`.
	ServiceTemplateGroupId *string `json:"ServiceTemplateGroupId,omitempty" name:"ServiceTemplateGroupId"`

	// Protocol port template group name.
	ServiceTemplateGroupName *string `json:"ServiceTemplateGroupName,omitempty" name:"ServiceTemplateGroupName"`

	// Protocol port template instance ID.
	ServiceTemplateIdSet []*string `json:"ServiceTemplateIdSet,omitempty" name:"ServiceTemplateIdSet" list`

	// Creation Time.
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// Protocol port template instance information.
	ServiceTemplateSet []*ServiceTemplate `json:"ServiceTemplateSet,omitempty" name:"ServiceTemplateSet" list`
}

type ServiceTemplateSpecification struct {

	// Protocol port ID, such as `ppm-f5n1f8da`.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Protocol port group ID, such as `ppmg-f5n1f8da`.
	ServiceGroupId *string `json:"ServiceGroupId,omitempty" name:"ServiceGroupId"`
}

type SetCcnRegionBandwidthLimitsRequest struct {
	*tchttp.BaseRequest

	// The CCN instance ID, such as `ccn-f49l6u0z`.
	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`

	// The outbound bandwidth cap of each CCN region.
	CcnRegionBandwidthLimits []*CcnRegionBandwidthLimit `json:"CcnRegionBandwidthLimits,omitempty" name:"CcnRegionBandwidthLimits" list`
}

func (r *SetCcnRegionBandwidthLimitsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SetCcnRegionBandwidthLimitsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SetCcnRegionBandwidthLimitsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetCcnRegionBandwidthLimitsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SetCcnRegionBandwidthLimitsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SourceIpTranslationNatRule struct {

	// Resource ID
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// Resource type. Valid values: SUBNET, NETWORKINTERFACE
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`

	// Source IP/IP range
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" name:"PrivateIpAddress"`

	// Elastic IP address pool
	PublicIpAddresses []*string `json:"PublicIpAddresses,omitempty" name:"PublicIpAddresses" list`

	// Description
	Description *string `json:"Description,omitempty" name:"Description"`

	// SNAT rule ID
	NatGatewaySnatId *string `json:"NatGatewaySnatId,omitempty" name:"NatGatewaySnatId"`

	// NAT Gateway ID
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	NatGatewayId *string `json:"NatGatewayId,omitempty" name:"NatGatewayId"`

	// VPC ID
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Creation time of a SNAT rule for a NAT Gateway
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
}

type Subnet struct {

	// The `ID` of the `VPC` instance.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Subnet instance `ID`, such as `subnet-bthucmmy`.
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// Subnet name.
	SubnetName *string `json:"SubnetName,omitempty" name:"SubnetName"`

	// The `IPv4` `CIDR` of the subnet.
	CidrBlock *string `json:"CidrBlock,omitempty" name:"CidrBlock"`

	// Whether it is the default subnet.
	IsDefault *bool `json:"IsDefault,omitempty" name:"IsDefault"`

	// Whether to enable broadcast.
	EnableBroadcast *bool `json:"EnableBroadcast,omitempty" name:"EnableBroadcast"`

	// Availability Zone.
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// The route table instance ID, such as `rtb-l2h8d7c2`.
	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`

	// Creation Time.
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// The number of available IPv4 addresses
	AvailableIpAddressCount *uint64 `json:"AvailableIpAddressCount,omitempty" name:"AvailableIpAddressCount"`

	// The `IPv6` `CIDR` of the subnet.
	Ipv6CidrBlock *string `json:"Ipv6CidrBlock,omitempty" name:"Ipv6CidrBlock"`

	// The associated `ACL`ID
	NetworkAclId *string `json:"NetworkAclId,omitempty" name:"NetworkAclId"`

	// Whether it is a `SNAT` address pool subnet.
	IsRemoteVpcSnat *bool `json:"IsRemoteVpcSnat,omitempty" name:"IsRemoteVpcSnat"`

	// The total number of IPv4 addresses in the subnet.
	TotalIpAddressCount *uint64 `json:"TotalIpAddressCount,omitempty" name:"TotalIpAddressCount"`

	// Tag key-value pairs
	TagSet []*Tag `json:"TagSet,omitempty" name:"TagSet" list`
}

type SubnetInput struct {

	// The `CIDR` of the subnet.
	CidrBlock *string `json:"CidrBlock,omitempty" name:"CidrBlock"`

	// Subnet name.
	SubnetName *string `json:"SubnetName,omitempty" name:"SubnetName"`

	// The availability zone, such as `ap-guangzhou-2`.
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// The specified associated route table, such as `rtb-3ryrwzuu`.
	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`
}

type Tag struct {

	// Tag key
	// Note: This field may return null, indicating no valid value.
	Key *string `json:"Key,omitempty" name:"Key"`

	// Tag value
	// Note: This field may return null, indicating no valid value.
	Value *string `json:"Value,omitempty" name:"Value"`
}

type TransformAddressRequest struct {
	*tchttp.BaseRequest

	// The ID of the instance with a common public IP to be operated on, such as `ins-11112222`. You can query the instance ID by logging into the [Console](https://console.cloud.tencent.com/cvm). You can also obtain the parameter value from the `InstanceId` field in the returned result of [DescribeInstances](https://intl.cloud.tencent.com/document/api/213/9389?from_cn_redirect=1) API.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *TransformAddressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *TransformAddressRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type TransformAddressResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TransformAddressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *TransformAddressResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UnassignIpv6AddressesRequest struct {
	*tchttp.BaseRequest

	// The `ID` of the ENI instance, such as `eni-m6dyj72l`.
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`

	// The list of specified `IPv6` addresses. A maximum of 10 can be specified each time.
	Ipv6Addresses []*Ipv6Address `json:"Ipv6Addresses,omitempty" name:"Ipv6Addresses" list`
}

func (r *UnassignIpv6AddressesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UnassignIpv6AddressesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UnassignIpv6AddressesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UnassignIpv6AddressesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UnassignIpv6AddressesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UnassignIpv6CidrBlockRequest struct {
	*tchttp.BaseRequest

	// The `ID` of the `VPC`, such as `vpc-f49l6u0z`.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// The `IPv6` IP range, such as `3402:4e00:20:1000::/56`
	Ipv6CidrBlock *string `json:"Ipv6CidrBlock,omitempty" name:"Ipv6CidrBlock"`
}

func (r *UnassignIpv6CidrBlockRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UnassignIpv6CidrBlockRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UnassignIpv6CidrBlockResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UnassignIpv6CidrBlockResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UnassignIpv6CidrBlockResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UnassignIpv6SubnetCidrBlockRequest struct {
	*tchttp.BaseRequest

	// The `ID` of the VPC where the subnet is located, such as `vpc-f49l6u0z`.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// The `IPv6` subnet IP range list.
	Ipv6SubnetCidrBlocks []*Ipv6SubnetCidrBlock `json:"Ipv6SubnetCidrBlocks,omitempty" name:"Ipv6SubnetCidrBlocks" list`
}

func (r *UnassignIpv6SubnetCidrBlockRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UnassignIpv6SubnetCidrBlockRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UnassignIpv6SubnetCidrBlockResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UnassignIpv6SubnetCidrBlockResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UnassignIpv6SubnetCidrBlockResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UnassignPrivateIpAddressesRequest struct {
	*tchttp.BaseRequest

	// The ID of the ENI instance, such as `eni-m6dyj72l`.
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`

	// The information of the specified private IPs. You can specify a maximum of 10 each time.
	PrivateIpAddresses []*PrivateIpAddressSpecification `json:"PrivateIpAddresses,omitempty" name:"PrivateIpAddresses" list`
}

func (r *UnassignPrivateIpAddressesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UnassignPrivateIpAddressesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UnassignPrivateIpAddressesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UnassignPrivateIpAddressesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UnassignPrivateIpAddressesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Vpc struct {

	// `VPC` name.
	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`

	// `VPC` instance `ID`, such as `vpc-azd4dt1c`.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// The `IPv4` `CIDR` of the `VPC`.
	CidrBlock *string `json:"CidrBlock,omitempty" name:"CidrBlock"`

	// Whether it is the default `VPC`.
	IsDefault *bool `json:"IsDefault,omitempty" name:"IsDefault"`

	// Whether multicast is enabled.
	EnableMulticast *bool `json:"EnableMulticast,omitempty" name:"EnableMulticast"`

	// Creation Time.
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// `DNS` list.
	DnsServerSet []*string `json:"DnsServerSet,omitempty" name:"DnsServerSet" list`

	// DHCP domain name option value.
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// `DHCP` option set `ID`.
	DhcpOptionsId *string `json:"DhcpOptionsId,omitempty" name:"DhcpOptionsId"`

	// Whether `DHCP` is enabled.
	EnableDhcp *bool `json:"EnableDhcp,omitempty" name:"EnableDhcp"`

	// The `IPv6` `CIDR` of the `VPC`.
	Ipv6CidrBlock *string `json:"Ipv6CidrBlock,omitempty" name:"Ipv6CidrBlock"`

	// Tag key-value pair
	TagSet []*Tag `json:"TagSet,omitempty" name:"TagSet" list`

	// The secondary CIDR block.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AssistantCidrSet []*AssistantCidr `json:"AssistantCidrSet,omitempty" name:"AssistantCidrSet" list`
}

type VpcIpv6Address struct {

	// `VPC` private `IPv6` address
	Ipv6Address *string `json:"Ipv6Address,omitempty" name:"Ipv6Address"`

	// The `IPv6` `CIDR` belonging to the subnet.
	CidrBlock *string `json:"CidrBlock,omitempty" name:"CidrBlock"`

	// `IPv6` type.
	Ipv6AddressType *string `json:"Ipv6AddressType,omitempty" name:"Ipv6AddressType"`

	// `IPv6` application time.
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
}

type VpcPrivateIpAddress struct {

	// `VPC` private `IP`.
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" name:"PrivateIpAddress"`

	// The `CIDR` belonging to the subnet.
	CidrBlock *string `json:"CidrBlock,omitempty" name:"CidrBlock"`

	// Private `IP` type.
	PrivateIpAddressType *string `json:"PrivateIpAddressType,omitempty" name:"PrivateIpAddressType"`

	// `IP` application time.
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
}

type VpnConnection struct {

	// Tunnel instance ID.
	VpnConnectionId *string `json:"VpnConnectionId,omitempty" name:"VpnConnectionId"`

	// Tunnel name.
	VpnConnectionName *string `json:"VpnConnectionName,omitempty" name:"VpnConnectionName"`

	// VPC instance ID.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// The ID of the VPN gateway instance.
	VpnGatewayId *string `json:"VpnGatewayId,omitempty" name:"VpnGatewayId"`

	// Customer gateway instance ID.
	CustomerGatewayId *string `json:"CustomerGatewayId,omitempty" name:"CustomerGatewayId"`

	// The pre-shared key.
	PreShareKey *string `json:"PreShareKey,omitempty" name:"PreShareKey"`

	// Tunnel transmission protocol.
	VpnProto *string `json:"VpnProto,omitempty" name:"VpnProto"`

	// Tunnel encryption protocol.
	EncryptProto *string `json:"EncryptProto,omitempty" name:"EncryptProto"`

	// Route Type.
	RouteType *string `json:"RouteType,omitempty" name:"RouteType"`

	// Creation Time.
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// Production status of the tunnel. PENDING: Creating; AVAILABLE: Running; DELETING: Deleting.
	State *string `json:"State,omitempty" name:"State"`

	// Connection status of the tunnel. AVAILABLE: Connected.
	NetStatus *string `json:"NetStatus,omitempty" name:"NetStatus"`

	// SPD.
	SecurityPolicyDatabaseSet []*SecurityPolicyDatabase `json:"SecurityPolicyDatabaseSet,omitempty" name:"SecurityPolicyDatabaseSet" list`

	// IKE options.
	IKEOptionsSpecification *IKEOptionsSpecification `json:"IKEOptionsSpecification,omitempty" name:"IKEOptionsSpecification"`

	// IPSEC options.
	IPSECOptionsSpecification *IPSECOptionsSpecification `json:"IPSECOptionsSpecification,omitempty" name:"IPSECOptionsSpecification"`

	// Whether the health check is supported.
	EnableHealthCheck *bool `json:"EnableHealthCheck,omitempty" name:"EnableHealthCheck"`

	// Local IP address for the health check
	HealthCheckLocalIp *string `json:"HealthCheckLocalIp,omitempty" name:"HealthCheckLocalIp"`

	// Peer IP address for the health check
	HealthCheckRemoteIp *string `json:"HealthCheckRemoteIp,omitempty" name:"HealthCheckRemoteIp"`

	// Tunnel health check status. Valid values: AVAILABLE: healthy; UNAVAILABLE: unhealthy. This parameter will be returned only after health check is enabled.
	HealthCheckStatus *string `json:"HealthCheckStatus,omitempty" name:"HealthCheckStatus"`
}

type VpnGateway struct {

	// Gateway instance ID.
	VpnGatewayId *string `json:"VpnGatewayId,omitempty" name:"VpnGatewayId"`

	// VPC instance ID.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Gateway instance name.
	VpnGatewayName *string `json:"VpnGatewayName,omitempty" name:"VpnGatewayName"`

	// Gateway instance type. Valid values: 'IPSEC', 'SSL', and 'CCN'.
	Type *string `json:"Type,omitempty" name:"Type"`

	// Gateway instance status. 'PENDING': Creating; 'DELETING': Deleting; 'AVAILABLE': Running.
	State *string `json:"State,omitempty" name:"State"`

	// Gateway public IP.
	PublicIpAddress *string `json:"PublicIpAddress,omitempty" name:"PublicIpAddress"`

	// Gateway renewal type: 'NOTIFY_AND_MANUAL_RENEW': Manual renewal. 'NOTIFY_AND_AUTO_RENEW': Automatic renewal. 'NOT_NOTIFY_AND_NOT_RENEW': No renewal after expiration.
	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`

	// Gateway billing type: POSTPAID_BY_HOUR: Postpaid by hour; PREPAID: Prepaid.
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`

	// Outbound bandwidth of gateway.
	InternetMaxBandwidthOut *uint64 `json:"InternetMaxBandwidthOut,omitempty" name:"InternetMaxBandwidthOut"`

	// Creation Time.
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// Expiration time of the prepaid gateway.
	ExpiredTime *string `json:"ExpiredTime,omitempty" name:"ExpiredTime"`

	// Whether the public IP is blocked.
	IsAddressBlocked *bool `json:"IsAddressBlocked,omitempty" name:"IsAddressBlocked"`

	// Change of billing method. PREPAID_TO_POSTPAID: Monthly subscription prepaid to postpaid by hour.
	NewPurchasePlan *string `json:"NewPurchasePlan,omitempty" name:"NewPurchasePlan"`

	// Gateway billing status. PROTECTIVELY_ISOLATED: Instance is isolated; NORMAL: Normal.
	RestrictState *string `json:"RestrictState,omitempty" name:"RestrictState"`

	// The availability zone, such as `ap-guangzhou-2`
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// Gateway bandwidth quota information.
	VpnGatewayQuotaSet []*VpnGatewayQuota `json:"VpnGatewayQuotaSet,omitempty" name:"VpnGatewayQuotaSet" list`

	// Gateway instance version.
	Version *string `json:"Version,omitempty" name:"Version"`

	// CCN instance ID when the value of Type is CCN.
	NetworkInstanceId *string `json:"NetworkInstanceId,omitempty" name:"NetworkInstanceId"`
}

type VpnGatewayQuota struct {

	// The bandwidth quota.
	Bandwidth *uint64 `json:"Bandwidth,omitempty" name:"Bandwidth"`

	// The bandwidth quota name in Chinese.
	Cname *string `json:"Cname,omitempty" name:"Cname"`

	// The bandwidth quota name in English.
	Name *string `json:"Name,omitempty" name:"Name"`
}

type VpngwCcnRoutes struct {

	// Route ID
	RouteId *string `json:"RouteId,omitempty" name:"RouteId"`

	// Enable the route or not
	// ENABLE: enable the route
	// DISABLE: do not enable the route
	Status *string `json:"Status,omitempty" name:"Status"`
}

type WithdrawNotifyRoutesRequest struct {
	*tchttp.BaseRequest

	// The unique ID of the route table
	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`

	// The unique ID of the routing policy
	RouteItemIds []*string `json:"RouteItemIds,omitempty" name:"RouteItemIds" list`
}

func (r *WithdrawNotifyRoutesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *WithdrawNotifyRoutesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type WithdrawNotifyRoutesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *WithdrawNotifyRoutesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *WithdrawNotifyRoutesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}
