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
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

// Predefined struct for user
type AcceptAttachCcnInstancesRequestParams struct {
	// The CCN instance ID, such as `ccn-f49l6u0z`.
	CcnId *string `json:"CcnId,omitnil,omitempty" name:"CcnId"`

	// List of associated instances.
	Instances []*CcnInstance `json:"Instances,omitnil,omitempty" name:"Instances"`
}

type AcceptAttachCcnInstancesRequest struct {
	*tchttp.BaseRequest
	
	// The CCN instance ID, such as `ccn-f49l6u0z`.
	CcnId *string `json:"CcnId,omitnil,omitempty" name:"CcnId"`

	// List of associated instances.
	Instances []*CcnInstance `json:"Instances,omitnil,omitempty" name:"Instances"`
}

func (r *AcceptAttachCcnInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AcceptAttachCcnInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CcnId")
	delete(f, "Instances")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AcceptAttachCcnInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AcceptAttachCcnInstancesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AcceptAttachCcnInstancesResponse struct {
	*tchttp.BaseResponse
	Response *AcceptAttachCcnInstancesResponseParams `json:"Response"`
}

func (r *AcceptAttachCcnInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AcceptAttachCcnInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AccountAttribute struct {
	// Attribute name
	AttributeName *string `json:"AttributeName,omitnil,omitempty" name:"AttributeName"`

	// Attribute values
	AttributeValues []*string `json:"AttributeValues,omitnil,omitempty" name:"AttributeValues"`
}

// Predefined struct for user
type AddBandwidthPackageResourcesRequestParams struct {
	// The unique ID of the source, such as 'eip-xxxx' and 'lb-xxxx'. EIP and LB resources are currently supported.
	ResourceIds []*string `json:"ResourceIds,omitnil,omitempty" name:"ResourceIds"`

	// The unique ID of the bandwidth package, such as 'bwp-xxxx'.
	BandwidthPackageId *string `json:"BandwidthPackageId,omitnil,omitempty" name:"BandwidthPackageId"`

	// The network type of the bandwidth package. Valid value: `BGP`, indicating that the internal resource is a BGP IP.
	NetworkType *string `json:"NetworkType,omitnil,omitempty" name:"NetworkType"`

	// The resource type, including `Address` and `LoadBalance`.
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// The protocol type of the bandwidth package. Valid values: `ipv4` and `ipv6`.
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`
}

type AddBandwidthPackageResourcesRequest struct {
	*tchttp.BaseRequest
	
	// The unique ID of the source, such as 'eip-xxxx' and 'lb-xxxx'. EIP and LB resources are currently supported.
	ResourceIds []*string `json:"ResourceIds,omitnil,omitempty" name:"ResourceIds"`

	// The unique ID of the bandwidth package, such as 'bwp-xxxx'.
	BandwidthPackageId *string `json:"BandwidthPackageId,omitnil,omitempty" name:"BandwidthPackageId"`

	// The network type of the bandwidth package. Valid value: `BGP`, indicating that the internal resource is a BGP IP.
	NetworkType *string `json:"NetworkType,omitnil,omitempty" name:"NetworkType"`

	// The resource type, including `Address` and `LoadBalance`.
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// The protocol type of the bandwidth package. Valid values: `ipv4` and `ipv6`.
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`
}

func (r *AddBandwidthPackageResourcesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddBandwidthPackageResourcesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ResourceIds")
	delete(f, "BandwidthPackageId")
	delete(f, "NetworkType")
	delete(f, "ResourceType")
	delete(f, "Protocol")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddBandwidthPackageResourcesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddBandwidthPackageResourcesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AddBandwidthPackageResourcesResponse struct {
	*tchttp.BaseResponse
	Response *AddBandwidthPackageResourcesResponseParams `json:"Response"`
}

func (r *AddBandwidthPackageResourcesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddBandwidthPackageResourcesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Address struct {
	// `EIP` `ID`, the unique ID of the `EIP`.
	AddressId *string `json:"AddressId,omitnil,omitempty" name:"AddressId"`

	// The `EIP` name.
	AddressName *string `json:"AddressName,omitnil,omitempty" name:"AddressName"`

	// Possible `EIP` states are 'CREATING', 'BINDING', 'BIND', 'UNBINDING', 'UNBIND', 'OFFLINING', and 'BIND_ENI'.
	AddressStatus *string `json:"AddressStatus,omitnil,omitempty" name:"AddressStatus"`

	// The public IP address
	AddressIp *string `json:"AddressIp,omitnil,omitempty" name:"AddressIp"`

	// The ID of the bound resource instance. This can be a `CVM` or `NAT`.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// The creation time, which follows the `ISO8601` standard and uses `UTC` time in the format of `YYYY-MM-DDThh:mm:ssZ`.
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// The ID of the bound ENI
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitnil,omitempty" name:"NetworkInterfaceId"`

	// The private IP of the bound resources
	PrivateAddressIp *string `json:"PrivateAddressIp,omitnil,omitempty" name:"PrivateAddressIp"`

	// The isolation status of the resource. `True` indicates the EIP is isolated. `False` indicates that the resource is not isolated.
	IsArrears *bool `json:"IsArrears,omitnil,omitempty" name:"IsArrears"`

	// The block status of the resource. `True` indicates the EIP is blocked. `False` indicates that the EIP is not blocked.
	IsBlocked *bool `json:"IsBlocked,omitnil,omitempty" name:"IsBlocked"`

	// Whether the EIP supports direct connection mode. `True` indicates the EIP supports direct connection. `False` indicates that the resource does not support direct connection.
	IsEipDirectConnection *bool `json:"IsEipDirectConnection,omitnil,omitempty" name:"IsEipDirectConnection"`

	// IP type. Valid values: `CalcIP` (device IP), `WanIP` (public network IP), `EIP` (general elastic IP), `AnycastEIP` (accelerated EIP), and `AntiDDoSEIP` (Anti DDoS EIP).
	AddressType *string `json:"AddressType,omitnil,omitempty" name:"AddressType"`

	// Whether the EIP is automatically released after being unbound. `True` indicates the EIP will be automatically released after being unbound. `False` indicates the EIP will not be automatically released after being unbound.
	CascadeRelease *bool `json:"CascadeRelease,omitnil,omitempty" name:"CascadeRelease"`

	// Type of the protocol used in EIP ALG
	EipAlgType *AlgType `json:"EipAlgType,omitnil,omitempty" name:"EipAlgType"`

	// The ISP of an EIP/Elastic IP, with possible return values currently including "CMCC", "CTCC", "CUCC" and "BGP"
	InternetServiceProvider *string `json:"InternetServiceProvider,omitnil,omitempty" name:"InternetServiceProvider"`

	// Whether the EIP is in a local BGP.
	LocalBgp *bool `json:"LocalBgp,omitnil,omitempty" name:"LocalBgp"`

	// Bandwidth value of EIP. The EIP for the bill-by-CVM account will return `null`.
	// Note: this field may return `null`, indicating that no valid value was found.
	Bandwidth *uint64 `json:"Bandwidth,omitnil,omitempty" name:"Bandwidth"`

	// Network billing mode of EIP. The EIP for the bill-by-CVM account will return `null`.
	// Note: this field may return `null`, indicating that no valid value was found.
	// Including:
	// <li><strong>BANDWIDTH_PREPAID_BY_MONTH</strong></li>
	// <p style="padding-left: 30px;">Prepaid by monthly-subscribed bandwidth.</p>
	// <li><strong>TRAFFIC_POSTPAID_BY_HOUR</strong></li>
	// <p style="padding-left: 30px;">Pay-as-you-go billing by hourly traffic.</p>
	// <li><strong>BANDWIDTH_POSTPAID_BY_HOUR</strong></li>
	// <p style="padding-left: 30px;">Pay-as-you-go billing by hourly bandwidth.</p>
	// <li><strong>BANDWIDTH_PACKAGE</strong></li>
	// <p style="padding-left: 30px;">Bandwidth package.</p>
	// Note: this field may return `null`, indicating that no valid value was found.
	InternetChargeType *string `json:"InternetChargeType,omitnil,omitempty" name:"InternetChargeType"`

	// List of tags associated with the EIP
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	TagSet []*Tag `json:"TagSet,omitnil,omitempty" name:"TagSet"`

	// The expiration time.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	DeadlineDate *string `json:"DeadlineDate,omitnil,omitempty" name:"DeadlineDate"`

	// The type of instance bound with the EIP
	// Note: this field may return `null`, indicating that no valid value was found.
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`


	Egress *string `json:"Egress,omitnil,omitempty" name:"Egress"`

	// ID of the Anti-DDoS service package. It is returned if the EIP is an Anti-DDoS EIP. 
	AntiDDoSPackageId *string `json:"AntiDDoSPackageId,omitnil,omitempty" name:"AntiDDoSPackageId"`


	RenewFlag *string `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`


	BandwidthPackageId *string `json:"BandwidthPackageId,omitnil,omitempty" name:"BandwidthPackageId"`
}

type AddressChargePrepaid struct {
	// Purchased usage period, in month. Valid values: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// Setting of renewal. Valid values: 0: manual renewal; 1: auto-renewal; 2: no renewal after expiration. Default value: 0
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`
}

type AddressInfo struct {
	// IP address
	Address *string `json:"Address,omitnil,omitempty" name:"Address"`

	// Remarks
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type AddressTemplate struct {
	// IP address template name.
	AddressTemplateName *string `json:"AddressTemplateName,omitnil,omitempty" name:"AddressTemplateName"`

	// The unique ID of the IP address template instance.
	AddressTemplateId *string `json:"AddressTemplateId,omitnil,omitempty" name:"AddressTemplateId"`

	// IP address information.
	AddressSet []*string `json:"AddressSet,omitnil,omitempty" name:"AddressSet"`

	// Creation Time.
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// IP address information with remarks
	AddressExtraSet []*AddressInfo `json:"AddressExtraSet,omitnil,omitempty" name:"AddressExtraSet"`
}

type AddressTemplateGroup struct {
	// IP address template group name.
	AddressTemplateGroupName *string `json:"AddressTemplateGroupName,omitnil,omitempty" name:"AddressTemplateGroupName"`

	// IP address template group instance ID, such as `ipmg-dih8xdbq`.
	AddressTemplateGroupId *string `json:"AddressTemplateGroupId,omitnil,omitempty" name:"AddressTemplateGroupId"`

	// IP address template ID.
	AddressTemplateIdSet []*string `json:"AddressTemplateIdSet,omitnil,omitempty" name:"AddressTemplateIdSet"`

	// Creation Time.
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// IP address template instance
	AddressTemplateSet []*AddressTemplateItem `json:"AddressTemplateSet,omitnil,omitempty" name:"AddressTemplateSet"`
}

type AddressTemplateItem struct {
	// ipm-xxxxxxxx
	AddressTemplateId *string `json:"AddressTemplateId,omitnil,omitempty" name:"AddressTemplateId"`

	// IP template name
	AddressTemplateName *string `json:"AddressTemplateName,omitnil,omitempty" name:"AddressTemplateName"`

	// Disused
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// Disused
	To *string `json:"To,omitnil,omitempty" name:"To"`
}

type AddressTemplateSpecification struct {
	// The ID of the IP address, such as `ipm-2uw6ujo6`.
	AddressId *string `json:"AddressId,omitnil,omitempty" name:"AddressId"`

	// The ID of the IP address group, such as `ipmg-2uw6ujo6`.
	AddressGroupId *string `json:"AddressGroupId,omitnil,omitempty" name:"AddressGroupId"`
}

// Predefined struct for user
type AdjustPublicAddressRequestParams struct {
	// The unique ID of the CVM instance, such as `ins-11112222`.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// The unique ID of the EIP, such as `eip-11112222`.
	AddressId *string `json:"AddressId,omitnil,omitempty" name:"AddressId"`
}

type AdjustPublicAddressRequest struct {
	*tchttp.BaseRequest
	
	// The unique ID of the CVM instance, such as `ins-11112222`.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// The unique ID of the EIP, such as `eip-11112222`.
	AddressId *string `json:"AddressId,omitnil,omitempty" name:"AddressId"`
}

func (r *AdjustPublicAddressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AdjustPublicAddressRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "AddressId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AdjustPublicAddressRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AdjustPublicAddressResponseParams struct {
	// The async task ID. You can use the [DescribeTaskResult](https://intl.cloud.tencent.com/document/api/215/36271?from_cn_redirect=1) API to query the task status.
	TaskId *uint64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AdjustPublicAddressResponse struct {
	*tchttp.BaseResponse
	Response *AdjustPublicAddressResponseParams `json:"Response"`
}

func (r *AdjustPublicAddressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AdjustPublicAddressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AlgType struct {
	// Whether FTP ALG is enabled
	Ftp *bool `json:"Ftp,omitnil,omitempty" name:"Ftp"`

	// Whether SIP ALG is enabled
	Sip *bool `json:"Sip,omitnil,omitempty" name:"Sip"`
}

// Predefined struct for user
type AllocateAddressesRequestParams struct {
	// The number of EIPs. Default: 1.
	AddressCount *int64 `json:"AddressCount,omitnil,omitempty" name:"AddressCount"`

	// The EIP line type. Default: BGP.
	// <ul style="margin:0"><li>For a user who has activated the static single-line IP allowlist, possible values are:<ul><li>CMCC: China Mobile</li>
	// <li>CTCC: China Telecom</li>
	// <li>CUCC: China Unicom</li></ul>Note: Only certain regions support static single-line IP addresses.</li></ul>
	InternetServiceProvider *string `json:"InternetServiceProvider,omitnil,omitempty" name:"InternetServiceProvider"`

	// The EIP billing method.
	// <ul style="margin:0"><li>For bill-by-IP account beta users, valid values: <ul><li>BANDWIDTH_PACKAGE: paid by the [bandwidth package](https://intl.cloud.tencent.com/document/product/684/15255?from_cn_redirect=1)(who must also be bandwidth package beta users)</li>
	// <li>BANDWIDTH_POSTPAID_BY_HOUR: billed by hourly bandwidth on a pay-as-you-go basis</li>
	// <li>BANDWIDTH_PREPAID_BY_MONTH: monthly bandwidth subscription</li>
	// <li>TRAFFIC_POSTPAID_BY_HOUR: billed by hourly traffic on a pay-as-you-go basis</li></ul>Default value: TRAFFIC_POSTPAID_BY_HOUR</li>
	// <li>If you are not a bill-by-IP account beta user, the EIP billing is the same as that for the instance bound to the EIP. Therefore, you do not need to pass in this parameter.</li></ul>
	InternetChargeType *string `json:"InternetChargeType,omitnil,omitempty" name:"InternetChargeType"`

	// The EIP outbound bandwidth cap, in Mbps.
	// <ul style="margin:0"><li>For bill-by-IP account beta users, the bandwidth cap range is determined by the EIP billing mode. <ul><li>`BANDWIDTH_PACKAGE`: 1 Mbps to 2000 Mbps</li>
	// <li>`BANDWIDTH_POSTPAID_BY_HOUR`: 1 Mbps to 100 Mbps</li>
	// <li>`BANDWIDTH_PREPAID_BY_MONTH`: 1 Mbps to 200 Mbps</li>
	// <li>`TRAFFIC_POSTPAID_BY_HOUR`: 1 Mbps to 100 Mbps</li></ul>Default value: 1 Mbps </li>
	// <li>If you are not a bill-by-IP account beta user, the EIP outbound bandwidth cap is subject to the bandwidth cap of the instance bound to the EIP. Therefore, you do not need to pass in this parameter. </li></ul>
	InternetMaxBandwidthOut *int64 `json:"InternetMaxBandwidthOut,omitnil,omitempty" name:"InternetMaxBandwidthOut"`

	// A required billing parameter for an EIP billed by monthly bandwidth subscription. For EIPs using other billing modes, it can be ignored.
	AddressChargePrepaid *AddressChargePrepaid `json:"AddressChargePrepaid,omitnil,omitempty" name:"AddressChargePrepaid"`

	// EIP type. Default value: EIP.
	// <ul style="margin:0"><li>For beta users of AIA, the value can be:</li></ul>`AnycastEIP`: an AIA IP address. For more information, see [Anycast Internet Acceleration](https://intl.cloud.tencent.com/document/product/644?from_cn_redirect=1).</li></ul>Note: Anycast EIPs are supported only in partial regions. </li></ul>
	// <ul style="margin:0"><li>For beta users of dedicated IP, the value can be: <ul><li>`HighQualityEIP`: Dedicated IP</li></ul>Note that dedicated IPs are only available in partial regions. </li></ul>
	// </ul>
	// <ul style="margin:0"><li>For beta users of Anti-DDoS IP, the value can be: <ul><li>`AntiDDoSEIP`: Anti-DDoS EIP</li></ul>Note that Anti-DDoS IPs are only available in partial regions. </li></ul>
	AddressType *string `json:"AddressType,omitnil,omitempty" name:"AddressType"`

	// Anycast publishing region
	// <ul style="margin:0"><li>Valid for users who have activated AIA. Values:<ul><li>ANYCAST_ZONE_GLOBAL: global publishing region </li><li>ANYCAST_ZONE_OVERSEAS: overseas publishing region</li><li><b>**[Disused]**</b> ANYCAST_ZONE_A: publishing region A (updated to ANYCAST_ZONE_GLOBAL)</li><li><b>**[Disused]**</b> ANYCAST_ZONE_B: publishing region B (updated to ANYCAST_ZONE_GLOBAL)</li></ul>Default: ANYCAST_ZONE_OVERSEAS.</li></ul>
	AnycastZone *string `json:"AnycastZone,omitnil,omitempty" name:"AnycastZone"`

	// <b>**[Disused]**</b>
	// Whether the Anycast EIP can be bound to CLB instances.
	// <ul style="margin:0"><li>Valid for users who have activated the AIA. Values:<ul><li>TRUE: the Anycast EIP can be bound to CLB instances.</li>
	// <li>FALSE: the Anycast EIP can be bound to CVMs, NAT gateways, and HAVIPs.</li></ul>Default: FALSE.</li></ul>
	ApplicableForCLB *bool `json:"ApplicableForCLB,omitnil,omitempty" name:"ApplicableForCLB"`

	// List of tags to be bound.
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// The unique ID of a BGP bandwidth package. If you configure this parameter and set InternetChargeType as BANDWIDTH_PACKAGE, the new EIP is added to this package and billed by the bandwidth package mode.
	BandwidthPackageId *string `json:"BandwidthPackageId,omitnil,omitempty" name:"BandwidthPackageId"`

	// EIP name, which is the custom EIP name given by the user when applying for the EIP. Default: not named
	AddressName *string `json:"AddressName,omitnil,omitempty" name:"AddressName"`

	// Network egress. It defaults to `center_egress1`.
	Egress *string `json:"Egress,omitnil,omitempty" name:"Egress"`

	// Anti-DDoS service package ID. This is required when you want to request an Anti-DDoS IP.
	AntiDDoSPackageId *string `json:"AntiDDoSPackageId,omitnil,omitempty" name:"AntiDDoSPackageId"`

	// A string used to ensure the idempotency of the request. Generate a value based on your client. This can ensure that the value is unique for different requests. It only supports ASCII characters and can contain up to 64 characters. 
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`
}

type AllocateAddressesRequest struct {
	*tchttp.BaseRequest
	
	// The number of EIPs. Default: 1.
	AddressCount *int64 `json:"AddressCount,omitnil,omitempty" name:"AddressCount"`

	// The EIP line type. Default: BGP.
	// <ul style="margin:0"><li>For a user who has activated the static single-line IP allowlist, possible values are:<ul><li>CMCC: China Mobile</li>
	// <li>CTCC: China Telecom</li>
	// <li>CUCC: China Unicom</li></ul>Note: Only certain regions support static single-line IP addresses.</li></ul>
	InternetServiceProvider *string `json:"InternetServiceProvider,omitnil,omitempty" name:"InternetServiceProvider"`

	// The EIP billing method.
	// <ul style="margin:0"><li>For bill-by-IP account beta users, valid values: <ul><li>BANDWIDTH_PACKAGE: paid by the [bandwidth package](https://intl.cloud.tencent.com/document/product/684/15255?from_cn_redirect=1)(who must also be bandwidth package beta users)</li>
	// <li>BANDWIDTH_POSTPAID_BY_HOUR: billed by hourly bandwidth on a pay-as-you-go basis</li>
	// <li>BANDWIDTH_PREPAID_BY_MONTH: monthly bandwidth subscription</li>
	// <li>TRAFFIC_POSTPAID_BY_HOUR: billed by hourly traffic on a pay-as-you-go basis</li></ul>Default value: TRAFFIC_POSTPAID_BY_HOUR</li>
	// <li>If you are not a bill-by-IP account beta user, the EIP billing is the same as that for the instance bound to the EIP. Therefore, you do not need to pass in this parameter.</li></ul>
	InternetChargeType *string `json:"InternetChargeType,omitnil,omitempty" name:"InternetChargeType"`

	// The EIP outbound bandwidth cap, in Mbps.
	// <ul style="margin:0"><li>For bill-by-IP account beta users, the bandwidth cap range is determined by the EIP billing mode. <ul><li>`BANDWIDTH_PACKAGE`: 1 Mbps to 2000 Mbps</li>
	// <li>`BANDWIDTH_POSTPAID_BY_HOUR`: 1 Mbps to 100 Mbps</li>
	// <li>`BANDWIDTH_PREPAID_BY_MONTH`: 1 Mbps to 200 Mbps</li>
	// <li>`TRAFFIC_POSTPAID_BY_HOUR`: 1 Mbps to 100 Mbps</li></ul>Default value: 1 Mbps </li>
	// <li>If you are not a bill-by-IP account beta user, the EIP outbound bandwidth cap is subject to the bandwidth cap of the instance bound to the EIP. Therefore, you do not need to pass in this parameter. </li></ul>
	InternetMaxBandwidthOut *int64 `json:"InternetMaxBandwidthOut,omitnil,omitempty" name:"InternetMaxBandwidthOut"`

	// A required billing parameter for an EIP billed by monthly bandwidth subscription. For EIPs using other billing modes, it can be ignored.
	AddressChargePrepaid *AddressChargePrepaid `json:"AddressChargePrepaid,omitnil,omitempty" name:"AddressChargePrepaid"`

	// EIP type. Default value: EIP.
	// <ul style="margin:0"><li>For beta users of AIA, the value can be:</li></ul>`AnycastEIP`: an AIA IP address. For more information, see [Anycast Internet Acceleration](https://intl.cloud.tencent.com/document/product/644?from_cn_redirect=1).</li></ul>Note: Anycast EIPs are supported only in partial regions. </li></ul>
	// <ul style="margin:0"><li>For beta users of dedicated IP, the value can be: <ul><li>`HighQualityEIP`: Dedicated IP</li></ul>Note that dedicated IPs are only available in partial regions. </li></ul>
	// </ul>
	// <ul style="margin:0"><li>For beta users of Anti-DDoS IP, the value can be: <ul><li>`AntiDDoSEIP`: Anti-DDoS EIP</li></ul>Note that Anti-DDoS IPs are only available in partial regions. </li></ul>
	AddressType *string `json:"AddressType,omitnil,omitempty" name:"AddressType"`

	// Anycast publishing region
	// <ul style="margin:0"><li>Valid for users who have activated AIA. Values:<ul><li>ANYCAST_ZONE_GLOBAL: global publishing region </li><li>ANYCAST_ZONE_OVERSEAS: overseas publishing region</li><li><b>**[Disused]**</b> ANYCAST_ZONE_A: publishing region A (updated to ANYCAST_ZONE_GLOBAL)</li><li><b>**[Disused]**</b> ANYCAST_ZONE_B: publishing region B (updated to ANYCAST_ZONE_GLOBAL)</li></ul>Default: ANYCAST_ZONE_OVERSEAS.</li></ul>
	AnycastZone *string `json:"AnycastZone,omitnil,omitempty" name:"AnycastZone"`

	// <b>**[Disused]**</b>
	// Whether the Anycast EIP can be bound to CLB instances.
	// <ul style="margin:0"><li>Valid for users who have activated the AIA. Values:<ul><li>TRUE: the Anycast EIP can be bound to CLB instances.</li>
	// <li>FALSE: the Anycast EIP can be bound to CVMs, NAT gateways, and HAVIPs.</li></ul>Default: FALSE.</li></ul>
	ApplicableForCLB *bool `json:"ApplicableForCLB,omitnil,omitempty" name:"ApplicableForCLB"`

	// List of tags to be bound.
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// The unique ID of a BGP bandwidth package. If you configure this parameter and set InternetChargeType as BANDWIDTH_PACKAGE, the new EIP is added to this package and billed by the bandwidth package mode.
	BandwidthPackageId *string `json:"BandwidthPackageId,omitnil,omitempty" name:"BandwidthPackageId"`

	// EIP name, which is the custom EIP name given by the user when applying for the EIP. Default: not named
	AddressName *string `json:"AddressName,omitnil,omitempty" name:"AddressName"`

	// Network egress. It defaults to `center_egress1`.
	Egress *string `json:"Egress,omitnil,omitempty" name:"Egress"`

	// Anti-DDoS service package ID. This is required when you want to request an Anti-DDoS IP.
	AntiDDoSPackageId *string `json:"AntiDDoSPackageId,omitnil,omitempty" name:"AntiDDoSPackageId"`

	// A string used to ensure the idempotency of the request. Generate a value based on your client. This can ensure that the value is unique for different requests. It only supports ASCII characters and can contain up to 64 characters. 
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`
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
	delete(f, "AddressCount")
	delete(f, "InternetServiceProvider")
	delete(f, "InternetChargeType")
	delete(f, "InternetMaxBandwidthOut")
	delete(f, "AddressChargePrepaid")
	delete(f, "AddressType")
	delete(f, "AnycastZone")
	delete(f, "ApplicableForCLB")
	delete(f, "Tags")
	delete(f, "BandwidthPackageId")
	delete(f, "AddressName")
	delete(f, "Egress")
	delete(f, "AntiDDoSPackageId")
	delete(f, "ClientToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AllocateAddressesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AllocateAddressesResponseParams struct {
	// List of the unique IDs of the requested EIPs.
	AddressSet []*string `json:"AddressSet,omitnil,omitempty" name:"AddressSet"`

	// The Async task ID. You can use the [DescribeTaskResult](https://intl.cloud.tencent.com/document/api/215/36271?from_cn_redirect=1) API to query the task status.
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

// Predefined struct for user
type AssignIpv6AddressesRequestParams struct {
	// The `ID` of the ENI instance, such as `eni-m6dyj72l`.
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitnil,omitempty" name:"NetworkInterfaceId"`

	// A list of `IPv6` addresses. You can specify a maximum of 10 at one time. The quota is calculated together with that of `Ipv6AddressCount`, a required input parameter alternative to this one.
	Ipv6Addresses []*Ipv6Address `json:"Ipv6Addresses,omitnil,omitempty" name:"Ipv6Addresses"`

	// The number of automatically assigned `IPv6` addresses. The total number of private IP addresses cannot exceed the quota. The quota is calculated together with that of `Ipv6Addresses`, a required input parameter alternative to this one.
	Ipv6AddressCount *uint64 `json:"Ipv6AddressCount,omitnil,omitempty" name:"Ipv6AddressCount"`
}

type AssignIpv6AddressesRequest struct {
	*tchttp.BaseRequest
	
	// The `ID` of the ENI instance, such as `eni-m6dyj72l`.
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitnil,omitempty" name:"NetworkInterfaceId"`

	// A list of `IPv6` addresses. You can specify a maximum of 10 at one time. The quota is calculated together with that of `Ipv6AddressCount`, a required input parameter alternative to this one.
	Ipv6Addresses []*Ipv6Address `json:"Ipv6Addresses,omitnil,omitempty" name:"Ipv6Addresses"`

	// The number of automatically assigned `IPv6` addresses. The total number of private IP addresses cannot exceed the quota. The quota is calculated together with that of `Ipv6Addresses`, a required input parameter alternative to this one.
	Ipv6AddressCount *uint64 `json:"Ipv6AddressCount,omitnil,omitempty" name:"Ipv6AddressCount"`
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
	delete(f, "NetworkInterfaceId")
	delete(f, "Ipv6Addresses")
	delete(f, "Ipv6AddressCount")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AssignIpv6AddressesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssignIpv6AddressesResponseParams struct {
	// The list of `IPv6` addresses assigned to ENIs.
	Ipv6AddressSet []*Ipv6Address `json:"Ipv6AddressSet,omitnil,omitempty" name:"Ipv6AddressSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type AssignIpv6CidrBlockRequestParams struct {
	// The `ID` of the `VPC`, such as `vpc-f49l6u0z`.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`
}

type AssignIpv6CidrBlockRequest struct {
	*tchttp.BaseRequest
	
	// The `ID` of the `VPC`, such as `vpc-f49l6u0z`.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`
}

func (r *AssignIpv6CidrBlockRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssignIpv6CidrBlockRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AssignIpv6CidrBlockRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssignIpv6CidrBlockResponseParams struct {
	// The assigned `IPv6` IP range, such as `3402:4e00:20:1000::/56`
	Ipv6CidrBlock *string `json:"Ipv6CidrBlock,omitnil,omitempty" name:"Ipv6CidrBlock"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AssignIpv6CidrBlockResponse struct {
	*tchttp.BaseResponse
	Response *AssignIpv6CidrBlockResponseParams `json:"Response"`
}

func (r *AssignIpv6CidrBlockResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssignIpv6CidrBlockResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssignIpv6SubnetCidrBlockRequestParams struct {
	// The `ID` of the VPC where the subnet is located, such as `vpc-f49l6u0z`.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// The assigned `IPv6` subnet IP range list.
	Ipv6SubnetCidrBlocks []*Ipv6SubnetCidrBlock `json:"Ipv6SubnetCidrBlocks,omitnil,omitempty" name:"Ipv6SubnetCidrBlocks"`
}

type AssignIpv6SubnetCidrBlockRequest struct {
	*tchttp.BaseRequest
	
	// The `ID` of the VPC where the subnet is located, such as `vpc-f49l6u0z`.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// The assigned `IPv6` subnet IP range list.
	Ipv6SubnetCidrBlocks []*Ipv6SubnetCidrBlock `json:"Ipv6SubnetCidrBlocks,omitnil,omitempty" name:"Ipv6SubnetCidrBlocks"`
}

func (r *AssignIpv6SubnetCidrBlockRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssignIpv6SubnetCidrBlockRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcId")
	delete(f, "Ipv6SubnetCidrBlocks")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AssignIpv6SubnetCidrBlockRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssignIpv6SubnetCidrBlockResponseParams struct {
	// The assigned `IPv6` subnet IP range list.
	Ipv6SubnetCidrBlockSet []*Ipv6SubnetCidrBlock `json:"Ipv6SubnetCidrBlockSet,omitnil,omitempty" name:"Ipv6SubnetCidrBlockSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AssignIpv6SubnetCidrBlockResponse struct {
	*tchttp.BaseResponse
	Response *AssignIpv6SubnetCidrBlockResponseParams `json:"Response"`
}

func (r *AssignIpv6SubnetCidrBlockResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssignIpv6SubnetCidrBlockResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssignPrivateIpAddressesRequestParams struct {
	// The ID of the ENI instance, such as `eni-m6dyj72l`.
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitnil,omitempty" name:"NetworkInterfaceId"`

	// The information on private IP addresses, of which you can specify a maximum of 10 at a time. You should provide either this parameter or SecondaryPrivateIpAddressCount, or both.
	PrivateIpAddresses []*PrivateIpAddressSpecification `json:"PrivateIpAddresses,omitnil,omitempty" name:"PrivateIpAddresses"`

	// The number of newly-applied private IP addresses. You should provide either this parameter or PrivateIpAddresses, or both. The total number of private IP addresses cannot exceed the quota. For more information, see<a href="/document/product/576/18527">ENI Use Limits</a>.
	SecondaryPrivateIpAddressCount *uint64 `json:"SecondaryPrivateIpAddressCount,omitnil,omitempty" name:"SecondaryPrivateIpAddressCount"`

	// IP u200dservice level. It is used together with `SecondaryPrivateIpAddressCount`. Values: PT` u200d(Gold), `AU` u200d(Silver), `AG `(Bronze) and DEFAULT (Default).
	QosLevel *string `json:"QosLevel,omitnil,omitempty" name:"QosLevel"`
}

type AssignPrivateIpAddressesRequest struct {
	*tchttp.BaseRequest
	
	// The ID of the ENI instance, such as `eni-m6dyj72l`.
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitnil,omitempty" name:"NetworkInterfaceId"`

	// The information on private IP addresses, of which you can specify a maximum of 10 at a time. You should provide either this parameter or SecondaryPrivateIpAddressCount, or both.
	PrivateIpAddresses []*PrivateIpAddressSpecification `json:"PrivateIpAddresses,omitnil,omitempty" name:"PrivateIpAddresses"`

	// The number of newly-applied private IP addresses. You should provide either this parameter or PrivateIpAddresses, or both. The total number of private IP addresses cannot exceed the quota. For more information, see<a href="/document/product/576/18527">ENI Use Limits</a>.
	SecondaryPrivateIpAddressCount *uint64 `json:"SecondaryPrivateIpAddressCount,omitnil,omitempty" name:"SecondaryPrivateIpAddressCount"`

	// IP u200dservice level. It is used together with `SecondaryPrivateIpAddressCount`. Values: PT` u200d(Gold), `AU` u200d(Silver), `AG `(Bronze) and DEFAULT (Default).
	QosLevel *string `json:"QosLevel,omitnil,omitempty" name:"QosLevel"`
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
	delete(f, "PrivateIpAddresses")
	delete(f, "SecondaryPrivateIpAddressCount")
	delete(f, "QosLevel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AssignPrivateIpAddressesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssignPrivateIpAddressesResponseParams struct {
	// The detailed information of the Private IP.
	PrivateIpAddressSet []*PrivateIpAddressSpecification `json:"PrivateIpAddressSet,omitnil,omitempty" name:"PrivateIpAddressSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// The `ID` of a `VPC` instance, such as `vpc-6v2ht8q5`.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// The secondary CIDR, such as `172.16.0.0/16`.
	CidrBlock *string `json:"CidrBlock,omitnil,omitempty" name:"CidrBlock"`

	// The secondary CIDR block type. 0: common secondary CIDR block. 1: container secondary CIDR block. Default: 0.
	AssistantType *int64 `json:"AssistantType,omitnil,omitempty" name:"AssistantType"`

	// Subnets divided by the secondary CIDR.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SubnetSet []*Subnet `json:"SubnetSet,omitnil,omitempty" name:"SubnetSet"`
}

// Predefined struct for user
type AssociateAddressRequestParams struct {
	// The unique ID of the EIP, such as `eip-11112222`.
	AddressId *string `json:"AddressId,omitnil,omitempty" name:"AddressId"`

	// The ID of the instance to be bound, such as `ins-11112222`, `lb-11112222`. You can query the instance ID by logging into the [Console](https://console.cloud.tencent.com/cvm). You can also obtain the parameter value from the `InstanceId` field in the returned result of [DescribeInstances](https://intl.cloud.tencent.com/document/api/213/15728?from_cn_redirect=1) API.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// The ID of the ENI to be bonud, such as `eni-11112222`. `NetworkInterfaceId` and `InstanceId` cannot be specified at the same time. You can query the ENI ID by logging into the [Console](https://console.cloud.tencent.com/vpc/eni). You can also obtain the parameter value from the `networkInterfaceId` field in the returned result of [DescribeNetworkInterfaces](https://intl.cloud.tencent.com/document/api/215/15817?from_cn_redirect=1) API.
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitnil,omitempty" name:"NetworkInterfaceId"`

	// The private IP to be bound. If you specify `NetworkInterfaceId`, then you must also specify `PrivateIpAddress`, indicating the EIP is bound to the specified private IP of the specified ENI. At the same time, you must ensure the specified `PrivateIpAddress` is a private IP on the `NetworkInterfaceId`. You can query the private IP of the specified ENI by logging into the [Console](https://console.cloud.tencent.com/vpc/eni). You can also obtain the parameter value from the `privateIpAddress` field in the returned result of [DescribeNetworkInterfaces](https://intl.cloud.tencent.com/document/api/215/15817?from_cn_redirect=1) API.
	PrivateIpAddress *string `json:"PrivateIpAddress,omitnil,omitempty" name:"PrivateIpAddress"`

	// Specify whether to configure direct access when binding EIPs. For details, see [EIP Direct Access](https://intl.cloud.tencent.com/document/product/213/12540). Valid values: `True` and `False` (default). This parameter can be set to `True` when binding EIPs to a CVM instance or EKS cluster. It is in a beta test. To try it out, please [submit a ticket](https://console.cloud.tencent.com/workorder/category?level1_id=6&level2_id=163&source=0&data_title=%E8%B4%9F%E8%BD%BD%E5%9D%87%E8%A1%A1%20CLB&level3_id=1071&queue=96&scene_code=34639&step=2).
	EipDirectConnection *bool `json:"EipDirectConnection,omitnil,omitempty" name:"EipDirectConnection"`
}

type AssociateAddressRequest struct {
	*tchttp.BaseRequest
	
	// The unique ID of the EIP, such as `eip-11112222`.
	AddressId *string `json:"AddressId,omitnil,omitempty" name:"AddressId"`

	// The ID of the instance to be bound, such as `ins-11112222`, `lb-11112222`. You can query the instance ID by logging into the [Console](https://console.cloud.tencent.com/cvm). You can also obtain the parameter value from the `InstanceId` field in the returned result of [DescribeInstances](https://intl.cloud.tencent.com/document/api/213/15728?from_cn_redirect=1) API.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// The ID of the ENI to be bonud, such as `eni-11112222`. `NetworkInterfaceId` and `InstanceId` cannot be specified at the same time. You can query the ENI ID by logging into the [Console](https://console.cloud.tencent.com/vpc/eni). You can also obtain the parameter value from the `networkInterfaceId` field in the returned result of [DescribeNetworkInterfaces](https://intl.cloud.tencent.com/document/api/215/15817?from_cn_redirect=1) API.
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitnil,omitempty" name:"NetworkInterfaceId"`

	// The private IP to be bound. If you specify `NetworkInterfaceId`, then you must also specify `PrivateIpAddress`, indicating the EIP is bound to the specified private IP of the specified ENI. At the same time, you must ensure the specified `PrivateIpAddress` is a private IP on the `NetworkInterfaceId`. You can query the private IP of the specified ENI by logging into the [Console](https://console.cloud.tencent.com/vpc/eni). You can also obtain the parameter value from the `privateIpAddress` field in the returned result of [DescribeNetworkInterfaces](https://intl.cloud.tencent.com/document/api/215/15817?from_cn_redirect=1) API.
	PrivateIpAddress *string `json:"PrivateIpAddress,omitnil,omitempty" name:"PrivateIpAddress"`

	// Specify whether to configure direct access when binding EIPs. For details, see [EIP Direct Access](https://intl.cloud.tencent.com/document/product/213/12540). Valid values: `True` and `False` (default). This parameter can be set to `True` when binding EIPs to a CVM instance or EKS cluster. It is in a beta test. To try it out, please [submit a ticket](https://console.cloud.tencent.com/workorder/category?level1_id=6&level2_id=163&source=0&data_title=%E8%B4%9F%E8%BD%BD%E5%9D%87%E8%A1%A1%20CLB&level3_id=1071&queue=96&scene_code=34639&step=2).
	EipDirectConnection *bool `json:"EipDirectConnection,omitnil,omitempty" name:"EipDirectConnection"`
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
	delete(f, "AddressId")
	delete(f, "InstanceId")
	delete(f, "NetworkInterfaceId")
	delete(f, "PrivateIpAddress")
	delete(f, "EipDirectConnection")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AssociateAddressRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssociateAddressResponseParams struct {
	// The async task ID. You can use the [DescribeTaskResult](https://intl.cloud.tencent.com/document/api/215/36271?from_cn_redirect=1) API to query the task status.
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type AssociateDirectConnectGatewayNatGatewayRequestParams struct {
	// VPC instance ID. VPC instance ID, which can be obtained from the `VpcId` field in the response of the `DescribeVpcs` API.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// The NAT Gateway ID.
	NatGatewayId *string `json:"NatGatewayId,omitnil,omitempty" name:"NatGatewayId"`

	// The direct connect gateway ID.
	DirectConnectGatewayId *string `json:"DirectConnectGatewayId,omitnil,omitempty" name:"DirectConnectGatewayId"`
}

type AssociateDirectConnectGatewayNatGatewayRequest struct {
	*tchttp.BaseRequest
	
	// VPC instance ID. VPC instance ID, which can be obtained from the `VpcId` field in the response of the `DescribeVpcs` API.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// The NAT Gateway ID.
	NatGatewayId *string `json:"NatGatewayId,omitnil,omitempty" name:"NatGatewayId"`

	// The direct connect gateway ID.
	DirectConnectGatewayId *string `json:"DirectConnectGatewayId,omitnil,omitempty" name:"DirectConnectGatewayId"`
}

func (r *AssociateDirectConnectGatewayNatGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateDirectConnectGatewayNatGatewayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcId")
	delete(f, "NatGatewayId")
	delete(f, "DirectConnectGatewayId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AssociateDirectConnectGatewayNatGatewayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssociateDirectConnectGatewayNatGatewayResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AssociateDirectConnectGatewayNatGatewayResponse struct {
	*tchttp.BaseResponse
	Response *AssociateDirectConnectGatewayNatGatewayResponseParams `json:"Response"`
}

func (r *AssociateDirectConnectGatewayNatGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateDirectConnectGatewayNatGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssociateNatGatewayAddressRequestParams struct {
	// The ID of the NAT gateway, such as `nat-df45454`.
	NatGatewayId *string `json:"NatGatewayId,omitnil,omitempty" name:"NatGatewayId"`

	// The number of EIPs you want to apply for. Either `AddressCount` or `PublicAddresses` must be passed in.
	AddressCount *uint64 `json:"AddressCount,omitnil,omitempty" name:"AddressCount"`

	// Array of the EIPs bound to the NAT gateway. Either `AddressCount` or `PublicAddresses` must be passed in.
	PublicIpAddresses []*string `json:"PublicIpAddresses,omitnil,omitempty" name:"PublicIpAddresses"`

	// The availability zone of the EIP, which is passed in when the EIP is automatically assigned.
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// The bandwidth size (in Mbps) of the EIP bound to the NAT gateway, which defaults to the maximum value applicable for the current user type.
	StockPublicIpAddressesBandwidthOut *uint64 `json:"StockPublicIpAddressesBandwidthOut,omitnil,omitempty" name:"StockPublicIpAddressesBandwidthOut"`

	// The size of the public network IP bandwidth to be applied for (in Mbps), which defaults to the maximum value applicable for the current user type.
	PublicIpAddressesBandwidthOut *uint64 `json:"PublicIpAddressesBandwidthOut,omitnil,omitempty" name:"PublicIpAddressesBandwidthOut"`

	// Whether the public IP and the NAT gateway must be in the same availability zone. Valid values: `true` and `false`. This parameter is valid only when `Zone` is specified.
	PublicIpFromSameZone *bool `json:"PublicIpFromSameZone,omitnil,omitempty" name:"PublicIpFromSameZone"`
}

type AssociateNatGatewayAddressRequest struct {
	*tchttp.BaseRequest
	
	// The ID of the NAT gateway, such as `nat-df45454`.
	NatGatewayId *string `json:"NatGatewayId,omitnil,omitempty" name:"NatGatewayId"`

	// The number of EIPs you want to apply for. Either `AddressCount` or `PublicAddresses` must be passed in.
	AddressCount *uint64 `json:"AddressCount,omitnil,omitempty" name:"AddressCount"`

	// Array of the EIPs bound to the NAT gateway. Either `AddressCount` or `PublicAddresses` must be passed in.
	PublicIpAddresses []*string `json:"PublicIpAddresses,omitnil,omitempty" name:"PublicIpAddresses"`

	// The availability zone of the EIP, which is passed in when the EIP is automatically assigned.
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// The bandwidth size (in Mbps) of the EIP bound to the NAT gateway, which defaults to the maximum value applicable for the current user type.
	StockPublicIpAddressesBandwidthOut *uint64 `json:"StockPublicIpAddressesBandwidthOut,omitnil,omitempty" name:"StockPublicIpAddressesBandwidthOut"`

	// The size of the public network IP bandwidth to be applied for (in Mbps), which defaults to the maximum value applicable for the current user type.
	PublicIpAddressesBandwidthOut *uint64 `json:"PublicIpAddressesBandwidthOut,omitnil,omitempty" name:"PublicIpAddressesBandwidthOut"`

	// Whether the public IP and the NAT gateway must be in the same availability zone. Valid values: `true` and `false`. This parameter is valid only when `Zone` is specified.
	PublicIpFromSameZone *bool `json:"PublicIpFromSameZone,omitnil,omitempty" name:"PublicIpFromSameZone"`
}

func (r *AssociateNatGatewayAddressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateNatGatewayAddressRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NatGatewayId")
	delete(f, "AddressCount")
	delete(f, "PublicIpAddresses")
	delete(f, "Zone")
	delete(f, "StockPublicIpAddressesBandwidthOut")
	delete(f, "PublicIpAddressesBandwidthOut")
	delete(f, "PublicIpFromSameZone")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AssociateNatGatewayAddressRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssociateNatGatewayAddressResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AssociateNatGatewayAddressResponse struct {
	*tchttp.BaseResponse
	Response *AssociateNatGatewayAddressResponseParams `json:"Response"`
}

func (r *AssociateNatGatewayAddressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateNatGatewayAddressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssociateNetworkAclSubnetsRequestParams struct {
	// Network ACL instance ID, such as `acl-12345678`.
	NetworkAclId *string `json:"NetworkAclId,omitnil,omitempty" name:"NetworkAclId"`

	// Array of subnet instance IDs, such as [subnet-12345678]
	SubnetIds []*string `json:"SubnetIds,omitnil,omitempty" name:"SubnetIds"`
}

type AssociateNetworkAclSubnetsRequest struct {
	*tchttp.BaseRequest
	
	// Network ACL instance ID, such as `acl-12345678`.
	NetworkAclId *string `json:"NetworkAclId,omitnil,omitempty" name:"NetworkAclId"`

	// Array of subnet instance IDs, such as [subnet-12345678]
	SubnetIds []*string `json:"SubnetIds,omitnil,omitempty" name:"SubnetIds"`
}

func (r *AssociateNetworkAclSubnetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateNetworkAclSubnetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NetworkAclId")
	delete(f, "SubnetIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AssociateNetworkAclSubnetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssociateNetworkAclSubnetsResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AssociateNetworkAclSubnetsResponse struct {
	*tchttp.BaseResponse
	Response *AssociateNetworkAclSubnetsResponseParams `json:"Response"`
}

func (r *AssociateNetworkAclSubnetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateNetworkAclSubnetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssociateNetworkInterfaceSecurityGroupsRequestParams struct {
	// ENI instance ID, e.g. eni-pxir56ns. You can enter up to 100 instances for each request.
	NetworkInterfaceIds []*string `json:"NetworkInterfaceIds,omitnil,omitempty" name:"NetworkInterfaceIds"`

	// The security group instance ID, such as `sg-33ocnj9n`. It can be obtained through DescribeSecurityGroups. You can enter up to 100 instances for each request.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`
}

type AssociateNetworkInterfaceSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// ENI instance ID, e.g. eni-pxir56ns. You can enter up to 100 instances for each request.
	NetworkInterfaceIds []*string `json:"NetworkInterfaceIds,omitnil,omitempty" name:"NetworkInterfaceIds"`

	// The security group instance ID, such as `sg-33ocnj9n`. It can be obtained through DescribeSecurityGroups. You can enter up to 100 instances for each request.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`
}

func (r *AssociateNetworkInterfaceSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateNetworkInterfaceSecurityGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NetworkInterfaceIds")
	delete(f, "SecurityGroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AssociateNetworkInterfaceSecurityGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssociateNetworkInterfaceSecurityGroupsResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AssociateNetworkInterfaceSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *AssociateNetworkInterfaceSecurityGroupsResponseParams `json:"Response"`
}

func (r *AssociateNetworkInterfaceSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateNetworkInterfaceSecurityGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AttachCcnInstancesRequestParams struct {
	// The CCN instance ID, such as `ccn-f49l6u0z`.
	CcnId *string `json:"CcnId,omitnil,omitempty" name:"CcnId"`

	// List of associated network instances
	Instances []*CcnInstance `json:"Instances,omitnil,omitempty" name:"Instances"`

	// The UIN (root account) of the CCN. By default, the current account belongs to the UIN
	CcnUin *string `json:"CcnUin,omitnil,omitempty" name:"CcnUin"`
}

type AttachCcnInstancesRequest struct {
	*tchttp.BaseRequest
	
	// The CCN instance ID, such as `ccn-f49l6u0z`.
	CcnId *string `json:"CcnId,omitnil,omitempty" name:"CcnId"`

	// List of associated network instances
	Instances []*CcnInstance `json:"Instances,omitnil,omitempty" name:"Instances"`

	// The UIN (root account) of the CCN. By default, the current account belongs to the UIN
	CcnUin *string `json:"CcnUin,omitnil,omitempty" name:"CcnUin"`
}

func (r *AttachCcnInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AttachCcnInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CcnId")
	delete(f, "Instances")
	delete(f, "CcnUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AttachCcnInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AttachCcnInstancesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AttachCcnInstancesResponse struct {
	*tchttp.BaseResponse
	Response *AttachCcnInstancesResponseParams `json:"Response"`
}

func (r *AttachCcnInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AttachCcnInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AttachClassicLinkVpcRequestParams struct {
	// VPC instance ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// CVM Instance ID
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

type AttachClassicLinkVpcRequest struct {
	*tchttp.BaseRequest
	
	// VPC instance ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// CVM Instance ID
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

func (r *AttachClassicLinkVpcRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AttachClassicLinkVpcRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcId")
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AttachClassicLinkVpcRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AttachClassicLinkVpcResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AttachClassicLinkVpcResponse struct {
	*tchttp.BaseResponse
	Response *AttachClassicLinkVpcResponseParams `json:"Response"`
}

func (r *AttachClassicLinkVpcResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AttachClassicLinkVpcResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AttachNetworkInterfaceRequestParams struct {
	// The ID of the ENI instance, such as `eni-m6dyj72l`.
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitnil,omitempty" name:"NetworkInterfaceId"`

	// The ID of the CVM instance, such as `ins-r8hr2upy`.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// ENI mounting type. Valid values: `0` (standard); `1` (extension); default value: `0`
	AttachType *uint64 `json:"AttachType,omitnil,omitempty" name:"AttachType"`
}

type AttachNetworkInterfaceRequest struct {
	*tchttp.BaseRequest
	
	// The ID of the ENI instance, such as `eni-m6dyj72l`.
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitnil,omitempty" name:"NetworkInterfaceId"`

	// The ID of the CVM instance, such as `ins-r8hr2upy`.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// ENI mounting type. Valid values: `0` (standard); `1` (extension); default value: `0`
	AttachType *uint64 `json:"AttachType,omitnil,omitempty" name:"AttachType"`
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
	delete(f, "AttachType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AttachNetworkInterfaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AttachNetworkInterfaceResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

// Predefined struct for user
type AttachSnapshotInstancesRequestParams struct {
	// Snapshot policy IDs
	SnapshotPolicyId *string `json:"SnapshotPolicyId,omitnil,omitempty" name:"SnapshotPolicyId"`

	// Information of associated instances
	Instances []*SnapshotInstance `json:"Instances,omitnil,omitempty" name:"Instances"`
}

type AttachSnapshotInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Snapshot policy IDs
	SnapshotPolicyId *string `json:"SnapshotPolicyId,omitnil,omitempty" name:"SnapshotPolicyId"`

	// Information of associated instances
	Instances []*SnapshotInstance `json:"Instances,omitnil,omitempty" name:"Instances"`
}

func (r *AttachSnapshotInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AttachSnapshotInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SnapshotPolicyId")
	delete(f, "Instances")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AttachSnapshotInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AttachSnapshotInstancesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AttachSnapshotInstancesResponse struct {
	*tchttp.BaseResponse
	Response *AttachSnapshotInstancesResponseParams `json:"Response"`
}

func (r *AttachSnapshotInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AttachSnapshotInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AuditCrossBorderComplianceRequestParams struct {
	// Service provider. Valid values: `UNICOM`.
	ServiceProvider *string `json:"ServiceProvider,omitnil,omitempty" name:"ServiceProvider"`

	// Unique ID of compliance review request.
	ComplianceId *uint64 `json:"ComplianceId,omitnil,omitempty" name:"ComplianceId"`

	// Audit behavior. Valid values: `APPROVED` and `DENY`.
	AuditBehavior *string `json:"AuditBehavior,omitnil,omitempty" name:"AuditBehavior"`
}

type AuditCrossBorderComplianceRequest struct {
	*tchttp.BaseRequest
	
	// Service provider. Valid values: `UNICOM`.
	ServiceProvider *string `json:"ServiceProvider,omitnil,omitempty" name:"ServiceProvider"`

	// Unique ID of compliance review request.
	ComplianceId *uint64 `json:"ComplianceId,omitnil,omitempty" name:"ComplianceId"`

	// Audit behavior. Valid values: `APPROVED` and `DENY`.
	AuditBehavior *string `json:"AuditBehavior,omitnil,omitempty" name:"AuditBehavior"`
}

func (r *AuditCrossBorderComplianceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AuditCrossBorderComplianceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceProvider")
	delete(f, "ComplianceId")
	delete(f, "AuditBehavior")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AuditCrossBorderComplianceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AuditCrossBorderComplianceResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AuditCrossBorderComplianceResponse struct {
	*tchttp.BaseResponse
	Response *AuditCrossBorderComplianceResponseParams `json:"Response"`
}

func (r *AuditCrossBorderComplianceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AuditCrossBorderComplianceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BackupPolicy struct {
	// Scheduled backup day. Values: `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`, `sunday`.
	BackupDay *string `json:"BackupDay,omitnil,omitempty" name:"BackupDay"`

	// Backup point in time. Format: HH:mm:ss.
	BackupTime *string `json:"BackupTime,omitnil,omitempty" name:"BackupTime"`
}

type BandwidthPackage struct {
	// The unique ID of the bandwidth package.
	BandwidthPackageId *string `json:"BandwidthPackageId,omitnil,omitempty" name:"BandwidthPackageId"`

	// Bandwidth package type. Values: `BGP`, `SINGLEISP`, `ANYCAST`, `SINGLEISP_CMCC`, `SINGLEISP_CTCC`, `SINGLEISP_CUCC`
	NetworkType *string `json:"NetworkType,omitnil,omitempty" name:"NetworkType"`

	// The bandwidth package billing mode. Valid values: 'TOP5_POSTPAID_BY_MONTH' and 'PERCENT95_POSTPAID_BY_MONTH'
	ChargeType *string `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`

	// The name of the bandwidth package.
	BandwidthPackageName *string `json:"BandwidthPackageName,omitnil,omitempty" name:"BandwidthPackageName"`

	// The creation time of the bandwidth package, which follows the `ISO8601` standard and uses `UTC` time in the format of `YYYY-MM-DDThh:mm:ssZ`.
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// The status of the bandwidth package. Valid values: 'CREATING', 'CREATED', 'DELETING', and 'DELETED'.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// The resource information of the bandwidth package.
	ResourceSet []*Resource `json:"ResourceSet,omitnil,omitempty" name:"ResourceSet"`

	// The limit of the bandwidth package in Mbps. The value '-1' indicates there is no limit.
	Bandwidth *int64 `json:"Bandwidth,omitnil,omitempty" name:"Bandwidth"`
}

type BandwidthPackageBillBandwidth struct {
	// Current billable usage, in Mbps
	BandwidthUsage *uint64 `json:"BandwidthUsage,omitnil,omitempty" name:"BandwidthUsage"`
}

type BatchModifySnapshotPolicy struct {
	// Snapshot policy IDs
	SnapshotPolicyId *string `json:"SnapshotPolicyId,omitnil,omitempty" name:"SnapshotPolicyId"`

	// Snapshot policy name
	SnapshotPolicyName *string `json:"SnapshotPolicyName,omitnil,omitempty" name:"SnapshotPolicyName"`

	// Backup policy
	BackupPolicies []*BackupPolicy `json:"BackupPolicies,omitnil,omitempty" name:"BackupPolicies"`

	// Snapshot retention period. Range: 1 to 365 days
	KeepTime *uint64 `json:"KeepTime,omitnil,omitempty" name:"KeepTime"`
}

type CCN struct {
	// The unique ID of the CCN
	CcnId *string `json:"CcnId,omitnil,omitempty" name:"CcnId"`

	// The name of the CCN
	CcnName *string `json:"CcnName,omitnil,omitempty" name:"CcnName"`

	// The detailed information of the CCN
	CcnDescription *string `json:"CcnDescription,omitnil,omitempty" name:"CcnDescription"`

	// The number of associated instances
	InstanceCount *uint64 `json:"InstanceCount,omitnil,omitempty" name:"InstanceCount"`

	// The creation time
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// The instance status. 'ISOLATED': Being isolated (instance is in arrears and service is suspended). 'AVAILABLE': Operating.
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// The instance service quality. ’PT’: Platinum , 'AU': Gold, 'AG': Silver.
	QosLevel *string `json:"QosLevel,omitnil,omitempty" name:"QosLevel"`

	// The billing method. POSTPAID indicates postpaid.
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// The limit type. INTER_REGION_LIMIT is the limit between regions. OUTER_REGION_LIMIT is a region egress limit.
	// Note: This field may return null, indicating no valid value.
	BandwidthLimitType *string `json:"BandwidthLimitType,omitnil,omitempty" name:"BandwidthLimitType"`

	// Tag key-value pairs.
	TagSet []*Tag `json:"TagSet,omitnil,omitempty" name:"TagSet"`

	// Whether the CCN route priority feature is supported. Valid values: False: do not support; True: support.
	RoutePriorityFlag *bool `json:"RoutePriorityFlag,omitnil,omitempty" name:"RoutePriorityFlag"`

	// Number of route tables associated with the instance.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	RouteTableCount *uint64 `json:"RouteTableCount,omitnil,omitempty" name:"RouteTableCount"`

	// Whether the multiple route tables feature is enabled for the CCN instance. Valid values: `False`: no; `True`: yes. Default value: `False`.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	RouteTableFlag *bool `json:"RouteTableFlag,omitnil,omitempty" name:"RouteTableFlag"`


	IsSecurityLock *bool `json:"IsSecurityLock,omitnil,omitempty" name:"IsSecurityLock"`

	// Status of CCN route broadcasting policy. Values: `False` (Disabled), `True` (Enabled)
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	RouteBroadcastPolicyFlag *bool `json:"RouteBroadcastPolicyFlag,omitnil,omitempty" name:"RouteBroadcastPolicyFlag"`
}

type CcnAttachedInstance struct {
	// The ID of a CCN instance.
	CcnId *string `json:"CcnId,omitnil,omitempty" name:"CcnId"`

	// The type of associated instances:
	// <li>`VPC`: VPC</li>
	// <li>`DIRECTCONNECT`: Direct Connect</li>
	// <li>`BMVPC`: BM VPC</li>
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// The ID of the associated instance.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// The name of the associated instance.
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// The region to which the associated instance belongs, such as `ap-guangzhou`.
	InstanceRegion *string `json:"InstanceRegion,omitnil,omitempty" name:"InstanceRegion"`

	// The UIN (root account) to which the associated instance belongs.
	InstanceUin *string `json:"InstanceUin,omitnil,omitempty" name:"InstanceUin"`

	// The CIDR of the associated instance.
	CidrBlock []*string `json:"CidrBlock,omitnil,omitempty" name:"CidrBlock"`

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
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// Association Time.
	AttachedTime *string `json:"AttachedTime,omitnil,omitempty" name:"AttachedTime"`

	// The UIN (root account) to which the CCN belongs.
	CcnUin *string `json:"CcnUin,omitnil,omitempty" name:"CcnUin"`

	// General location of the associated instance, such as CHINA_MAINLAND.
	InstanceArea *string `json:"InstanceArea,omitnil,omitempty" name:"InstanceArea"`

	// Description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Route table ID
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	RouteTableId *string `json:"RouteTableId,omitnil,omitempty" name:"RouteTableId"`

	// Route table name
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	RouteTableName *string `json:"RouteTableName,omitnil,omitempty" name:"RouteTableName"`
}

type CcnBandwidthInfo struct {
	// The CCN ID that the bandwidth belongs to.
	// Note: this field may return null, indicating that no valid value was found.
	CcnId *string `json:"CcnId,omitnil,omitempty" name:"CcnId"`

	// The creation time of the instance.
	// Note: this field may return null, indicating that no valid value was found.
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// The expiration time of the instance.
	// Note: this field may return null, indicating that no valid value was found.
	ExpiredTime *string `json:"ExpiredTime,omitnil,omitempty" name:"ExpiredTime"`

	// The unique ID of the bandwidth instance.
	// Note: this field may return null, indicating that no valid value was found.
	RegionFlowControlId *string `json:"RegionFlowControlId,omitnil,omitempty" name:"RegionFlowControlId"`

	// The billing flag.
	// Note: this field may return null, indicating that no valid value was found.
	RenewFlag *string `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// The information of the bandwidth regions and bandwidth caps. The parameter is only returned for the cross-region limit mode, but not for egress limit.
	// Note: this field may return null, indicating that no valid value was found.
	CcnRegionBandwidthLimit *CcnRegionBandwidthLimit `json:"CcnRegionBandwidthLimit,omitnil,omitempty" name:"CcnRegionBandwidthLimit"`

	// Cloud marketplace instance ID.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	MarketId *string `json:"MarketId,omitnil,omitempty" name:"MarketId"`

	// The list of tags to be bound.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	TagSet []*Tag `json:"TagSet,omitnil,omitempty" name:"TagSet"`
}

type CcnInstance struct {
	// The ID of the associated instance.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// The region to which the associated instance ID belongs, such as `ap-guangzhou`.
	InstanceRegion *string `json:"InstanceRegion,omitnil,omitempty" name:"InstanceRegion"`

	// The type of the associated instance. Available values are:
	// <li>`VPC`: VPC</li>
	// <li>`DIRECTCONNECT`: Direct Connect</li>
	// <li>`BMVPC`: BM VPC</li>
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// Description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// The ID of the route table associated with the instance
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	RouteTableId *string `json:"RouteTableId,omitnil,omitempty" name:"RouteTableId"`
}

type CcnRegionBandwidthLimit struct {
	// Region, such as `ap-guangzhou`
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// The outbound bandwidth cap. Units: Mbps
	BandwidthLimit *uint64 `json:"BandwidthLimit,omitnil,omitempty" name:"BandwidthLimit"`

	// Whether it is a BM region. The default is `false`.
	IsBm *bool `json:"IsBm,omitnil,omitempty" name:"IsBm"`

	// The target region, such as `ap-shanghai`
	// Note: This field may return null, indicating no valid value.
	DstRegion *string `json:"DstRegion,omitnil,omitempty" name:"DstRegion"`

	// Whether the target region is a BM region. The default is `false`.
	DstIsBm *bool `json:"DstIsBm,omitnil,omitempty" name:"DstIsBm"`
}

type CcnRoute struct {
	// The ID of the routing policy
	RouteId *string `json:"RouteId,omitnil,omitempty" name:"RouteId"`

	// Destination
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitnil,omitempty" name:"DestinationCidrBlock"`

	// The type of the next hop (associated instance type). Available types: VPC, DIRECTCONNECT
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// The next hop (associated instance)
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// The name of the next hop (associated instance name)
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// The region of the next hop (the region of the associated instance)
	InstanceRegion *string `json:"InstanceRegion,omitnil,omitempty" name:"InstanceRegion"`

	// Update Time
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// Whether the route is enabled
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// The UIN (root account) to which the associated instance belongs
	InstanceUin *string `json:"InstanceUin,omitnil,omitempty" name:"InstanceUin"`

	// Additional status of the route
	ExtraState *string `json:"ExtraState,omitnil,omitempty" name:"ExtraState"`

	// Whether it is a dynamic route
	IsBgp *bool `json:"IsBgp,omitnil,omitempty" name:"IsBgp"`

	// Route priority
	RoutePriority *uint64 `json:"RoutePriority,omitnil,omitempty" name:"RoutePriority"`

	// Next hop port name (associated instance’s port name)
	InstanceExtraName *string `json:"InstanceExtraName,omitnil,omitempty" name:"InstanceExtraName"`
}

// Predefined struct for user
type CheckAssistantCidrRequestParams struct {
	// `VPC` instance `ID`, e.g. `vpc-6v2ht8q5`.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Load CIDR blocks to add. CIDR block set; format: e.g. ["10.0.0.0/16", "172.16.0.0/16"]
	NewCidrBlocks []*string `json:"NewCidrBlocks,omitnil,omitempty" name:"NewCidrBlocks"`

	// Load CIDR blocks to delete. CIDR block set; Format: e.g. ["10.0.0.0/16", "172.16.0.0/16"]
	OldCidrBlocks []*string `json:"OldCidrBlocks,omitnil,omitempty" name:"OldCidrBlocks"`
}

type CheckAssistantCidrRequest struct {
	*tchttp.BaseRequest
	
	// `VPC` instance `ID`, e.g. `vpc-6v2ht8q5`.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Load CIDR blocks to add. CIDR block set; format: e.g. ["10.0.0.0/16", "172.16.0.0/16"]
	NewCidrBlocks []*string `json:"NewCidrBlocks,omitnil,omitempty" name:"NewCidrBlocks"`

	// Load CIDR blocks to delete. CIDR block set; Format: e.g. ["10.0.0.0/16", "172.16.0.0/16"]
	OldCidrBlocks []*string `json:"OldCidrBlocks,omitnil,omitempty" name:"OldCidrBlocks"`
}

func (r *CheckAssistantCidrRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckAssistantCidrRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcId")
	delete(f, "NewCidrBlocks")
	delete(f, "OldCidrBlocks")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckAssistantCidrRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckAssistantCidrResponseParams struct {
	// Array of conflict resources.
	ConflictSourceSet []*ConflictSource `json:"ConflictSourceSet,omitnil,omitempty" name:"ConflictSourceSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CheckAssistantCidrResponse struct {
	*tchttp.BaseResponse
	Response *CheckAssistantCidrResponseParams `json:"Response"`
}

func (r *CheckAssistantCidrResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckAssistantCidrResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckNetDetectStateRequestParams struct {
	// The array of detection destination IPv4 addresses, which contains at most two IP addresses.
	DetectDestinationIp []*string `json:"DetectDestinationIp,omitnil,omitempty" name:"DetectDestinationIp"`

	// The type of the next hop. Currently supported types are:
	// VPN: VPN gateway;
	// DIRECTCONNECT: direct connect gateway;
	// PEERCONNECTION: peering connection;
	// NAT: NAT gateway;
	// NORMAL_CVM: normal CVM.
	NextHopType *string `json:"NextHopType,omitnil,omitempty" name:"NextHopType"`

	// The next-hop destination gateway. The value is related to NextHopType.
	// If NextHopType is set to VPN, the value of this parameter is the VPN gateway ID, such as vpngw-12345678.
	// If NextHopType is set to DIRECTCONNECT, the value of this parameter is the direct connect gateway ID, such as dcg-12345678.
	// If NextHopType is set to PEERCONNECTION, the value of this parameter is the peering connection ID, such as pcx-12345678.
	// If NextHopType is set to NAT, the value of this parameter is the NAT gateway ID, such as nat-12345678.
	// If NextHopType is set to NORMAL_CVM, the value of this parameter is the IPv4 address of the CVM, such as 10.0.0.12.
	NextHopDestination *string `json:"NextHopDestination,omitnil,omitempty" name:"NextHopDestination"`

	// ID of a network inspector instance, e.g. netd-12345678. Enter at least one of this parameter, VpcId, SubnetId, and NetDetectName. Use NetDetectId if it is present.
	NetDetectId *string `json:"NetDetectId,omitnil,omitempty" name:"NetDetectId"`

	// ID of a `VPC` instance, e.g. `vpc-12345678`, which is used together with SubnetId and NetDetectName. You should enter either this parameter or NetDetectId, or both. Use NetDetectId if it is present.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// ID of a subnet instance, e.g. `subnet-12345678`, which is used together with VpcId and NetDetectName. You should enter either this parameter or NetDetectId, or both. Use NetDetectId if it is present.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// The name of a network inspector, up to 60 bytes in length. It is used together with VpcId and NetDetectName. You should enter either this parameter or NetDetectId, or both. Use NetDetectId if it is present.
	NetDetectName *string `json:"NetDetectName,omitnil,omitempty" name:"NetDetectName"`
}

type CheckNetDetectStateRequest struct {
	*tchttp.BaseRequest
	
	// The array of detection destination IPv4 addresses, which contains at most two IP addresses.
	DetectDestinationIp []*string `json:"DetectDestinationIp,omitnil,omitempty" name:"DetectDestinationIp"`

	// The type of the next hop. Currently supported types are:
	// VPN: VPN gateway;
	// DIRECTCONNECT: direct connect gateway;
	// PEERCONNECTION: peering connection;
	// NAT: NAT gateway;
	// NORMAL_CVM: normal CVM.
	NextHopType *string `json:"NextHopType,omitnil,omitempty" name:"NextHopType"`

	// The next-hop destination gateway. The value is related to NextHopType.
	// If NextHopType is set to VPN, the value of this parameter is the VPN gateway ID, such as vpngw-12345678.
	// If NextHopType is set to DIRECTCONNECT, the value of this parameter is the direct connect gateway ID, such as dcg-12345678.
	// If NextHopType is set to PEERCONNECTION, the value of this parameter is the peering connection ID, such as pcx-12345678.
	// If NextHopType is set to NAT, the value of this parameter is the NAT gateway ID, such as nat-12345678.
	// If NextHopType is set to NORMAL_CVM, the value of this parameter is the IPv4 address of the CVM, such as 10.0.0.12.
	NextHopDestination *string `json:"NextHopDestination,omitnil,omitempty" name:"NextHopDestination"`

	// ID of a network inspector instance, e.g. netd-12345678. Enter at least one of this parameter, VpcId, SubnetId, and NetDetectName. Use NetDetectId if it is present.
	NetDetectId *string `json:"NetDetectId,omitnil,omitempty" name:"NetDetectId"`

	// ID of a `VPC` instance, e.g. `vpc-12345678`, which is used together with SubnetId and NetDetectName. You should enter either this parameter or NetDetectId, or both. Use NetDetectId if it is present.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// ID of a subnet instance, e.g. `subnet-12345678`, which is used together with VpcId and NetDetectName. You should enter either this parameter or NetDetectId, or both. Use NetDetectId if it is present.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// The name of a network inspector, up to 60 bytes in length. It is used together with VpcId and NetDetectName. You should enter either this parameter or NetDetectId, or both. Use NetDetectId if it is present.
	NetDetectName *string `json:"NetDetectName,omitnil,omitempty" name:"NetDetectName"`
}

func (r *CheckNetDetectStateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckNetDetectStateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DetectDestinationIp")
	delete(f, "NextHopType")
	delete(f, "NextHopDestination")
	delete(f, "NetDetectId")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "NetDetectName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckNetDetectStateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckNetDetectStateResponseParams struct {
	// The array of network detection verification results.
	NetDetectIpStateSet []*NetDetectIpState `json:"NetDetectIpStateSet,omitnil,omitempty" name:"NetDetectIpStateSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CheckNetDetectStateResponse struct {
	*tchttp.BaseResponse
	Response *CheckNetDetectStateResponseParams `json:"Response"`
}

func (r *CheckNetDetectStateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckNetDetectStateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CidrForCcn struct {
	// Local CIDR block, including subnet CIDR block and secondary CIDR block
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Cidr *string `json:"Cidr,omitnil,omitempty" name:"Cidr"`

	// Whether the routing policy of the VPC subnet is published to CCN.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	PublishedToVbc *bool `json:"PublishedToVbc,omitnil,omitempty" name:"PublishedToVbc"`
}

type ClassicLinkInstance struct {
	// VPC instance ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// The unique ID of the CVM instance
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

// Predefined struct for user
type CloneSecurityGroupRequestParams struct {
	// Security group instance ID, such as `esg-33ocnj9n`, which can be obtained through the <a href="https://intl.cloud.tencent.com/document/product/215/15808?from_cn_redirect=1">DescribeSecurityGroups</a> API.
	SecurityGroupId *string `json:"SecurityGroupId,omitnil,omitempty" name:"SecurityGroupId"`

	// The name of security group clone. You can enter any name within 60 characters. If this parameter is left empty, the security group clone will use the name of the source security group.
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// Description of the security group clone. You can enter up to 100 characters. If this parameter is left empty, the security group clone will use the description of the source security group.
	GroupDescription *string `json:"GroupDescription,omitnil,omitempty" name:"GroupDescription"`

	// Project ID. Default value: 0. You can query it on the <a href="https://console.cloud.tencent.com/project">project management page</a> of the Tencent Cloud console.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// The region of the source security group for a cross-region clone. For example, to clone the security group in Guangzhou to Shanghai, set it to `ap-guangzhou`.
	RemoteRegion *string `json:"RemoteRegion,omitnil,omitempty" name:"RemoteRegion"`
}

type CloneSecurityGroupRequest struct {
	*tchttp.BaseRequest
	
	// Security group instance ID, such as `esg-33ocnj9n`, which can be obtained through the <a href="https://intl.cloud.tencent.com/document/product/215/15808?from_cn_redirect=1">DescribeSecurityGroups</a> API.
	SecurityGroupId *string `json:"SecurityGroupId,omitnil,omitempty" name:"SecurityGroupId"`

	// The name of security group clone. You can enter any name within 60 characters. If this parameter is left empty, the security group clone will use the name of the source security group.
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// Description of the security group clone. You can enter up to 100 characters. If this parameter is left empty, the security group clone will use the description of the source security group.
	GroupDescription *string `json:"GroupDescription,omitnil,omitempty" name:"GroupDescription"`

	// Project ID. Default value: 0. You can query it on the <a href="https://console.cloud.tencent.com/project">project management page</a> of the Tencent Cloud console.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// The region of the source security group for a cross-region clone. For example, to clone the security group in Guangzhou to Shanghai, set it to `ap-guangzhou`.
	RemoteRegion *string `json:"RemoteRegion,omitnil,omitempty" name:"RemoteRegion"`
}

func (r *CloneSecurityGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloneSecurityGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecurityGroupId")
	delete(f, "GroupName")
	delete(f, "GroupDescription")
	delete(f, "ProjectId")
	delete(f, "RemoteRegion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CloneSecurityGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloneSecurityGroupResponseParams struct {
	// Security group object
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	SecurityGroup *SecurityGroup `json:"SecurityGroup,omitnil,omitempty" name:"SecurityGroup"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CloneSecurityGroupResponse struct {
	*tchttp.BaseResponse
	Response *CloneSecurityGroupResponseParams `json:"Response"`
}

func (r *CloneSecurityGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloneSecurityGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConflictItem struct {
	// Conflict resource ID
	ConfilctId *string `json:"ConfilctId,omitnil,omitempty" name:"ConfilctId"`

	// Conflict destination resource
	DestinationItem *string `json:"DestinationItem,omitnil,omitempty" name:"DestinationItem"`
}

type ConflictSource struct {
	// Conflict resource ID
	ConflictSourceId *string `json:"ConflictSourceId,omitnil,omitempty" name:"ConflictSourceId"`

	// Conflict resource
	SourceItem *string `json:"SourceItem,omitnil,omitempty" name:"SourceItem"`

	// Conflict resource items
	ConflictItemSet []*ConflictItem `json:"ConflictItemSet,omitnil,omitempty" name:"ConflictItemSet"`
}

// Predefined struct for user
type CreateAddressTemplateGroupRequestParams struct {
	// Name of the IP address template group
	AddressTemplateGroupName *string `json:"AddressTemplateGroupName,omitnil,omitempty" name:"AddressTemplateGroupName"`

	// IP address template instance ID, such as `ipm-mdunqeb6`.
	AddressTemplateIds []*string `json:"AddressTemplateIds,omitnil,omitempty" name:"AddressTemplateIds"`
}

type CreateAddressTemplateGroupRequest struct {
	*tchttp.BaseRequest
	
	// Name of the IP address template group
	AddressTemplateGroupName *string `json:"AddressTemplateGroupName,omitnil,omitempty" name:"AddressTemplateGroupName"`

	// IP address template instance ID, such as `ipm-mdunqeb6`.
	AddressTemplateIds []*string `json:"AddressTemplateIds,omitnil,omitempty" name:"AddressTemplateIds"`
}

func (r *CreateAddressTemplateGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAddressTemplateGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AddressTemplateGroupName")
	delete(f, "AddressTemplateIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAddressTemplateGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAddressTemplateGroupResponseParams struct {
	// Group object of the IP address template.
	AddressTemplateGroup *AddressTemplateGroup `json:"AddressTemplateGroup,omitnil,omitempty" name:"AddressTemplateGroup"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAddressTemplateGroupResponse struct {
	*tchttp.BaseResponse
	Response *CreateAddressTemplateGroupResponseParams `json:"Response"`
}

func (r *CreateAddressTemplateGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAddressTemplateGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAddressTemplateRequestParams struct {
	// IP address template name.
	AddressTemplateName *string `json:"AddressTemplateName,omitnil,omitempty" name:"AddressTemplateName"`

	// The address information can be presented by the IP, CIDR block or IP address range. Either Addresses or AddressesExtra is required.
	Addresses []*string `json:"Addresses,omitnil,omitempty" name:"Addresses"`

	// The address information can contain remarks and be presented by the IP, CIDR block or IP address range. Either Addresses or AddressesExtra is required.
	AddressesExtra []*AddressInfo `json:"AddressesExtra,omitnil,omitempty" name:"AddressesExtra"`
}

type CreateAddressTemplateRequest struct {
	*tchttp.BaseRequest
	
	// IP address template name.
	AddressTemplateName *string `json:"AddressTemplateName,omitnil,omitempty" name:"AddressTemplateName"`

	// The address information can be presented by the IP, CIDR block or IP address range. Either Addresses or AddressesExtra is required.
	Addresses []*string `json:"Addresses,omitnil,omitempty" name:"Addresses"`

	// The address information can contain remarks and be presented by the IP, CIDR block or IP address range. Either Addresses or AddressesExtra is required.
	AddressesExtra []*AddressInfo `json:"AddressesExtra,omitnil,omitempty" name:"AddressesExtra"`
}

func (r *CreateAddressTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAddressTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AddressTemplateName")
	delete(f, "Addresses")
	delete(f, "AddressesExtra")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAddressTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAddressTemplateResponseParams struct {
	// The template object of the IP address.
	AddressTemplate *AddressTemplate `json:"AddressTemplate,omitnil,omitempty" name:"AddressTemplate"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAddressTemplateResponse struct {
	*tchttp.BaseResponse
	Response *CreateAddressTemplateResponseParams `json:"Response"`
}

func (r *CreateAddressTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAddressTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAndAttachNetworkInterfaceRequestParams struct {
	// The ID of the VPC instance. You can obtain the parameter value from the `VpcId` field in the returned result of the `DescribeVpcs` API.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// The name of the ENI. The maximum length is 60 bytes.
	NetworkInterfaceName *string `json:"NetworkInterfaceName,omitnil,omitempty" name:"NetworkInterfaceName"`

	// The subnet instance ID of the ENI, such as 'subnet-0ap8nwca'.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// CVM instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// The information of the specified private IPs. You can specify a maximum of 10 IPs each time.
	PrivateIpAddresses []*PrivateIpAddressSpecification `json:"PrivateIpAddresses,omitnil,omitempty" name:"PrivateIpAddresses"`

	// The number of private IP addresses you apply for. The total number of private IP addresses cannot exceed the quota.
	SecondaryPrivateIpAddressCount *uint64 `json:"SecondaryPrivateIpAddressCount,omitnil,omitempty" name:"SecondaryPrivateIpAddressCount"`

	// IP u200dservice level. It is used together with `SecondaryPrivateIpAddressCount`. Values: `PT` u200d(Gold), `AU` u200d(Silver), `AG` (Bronze) and `DEFAULT` (Default).
	QosLevel *string `json:"QosLevel,omitnil,omitempty" name:"QosLevel"`

	// The security group to be bound with, such as ['sg-1dd51d'].
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// The ENI description. You can enter any information within 60 characters.
	NetworkInterfaceDescription *string `json:"NetworkInterfaceDescription,omitnil,omitempty" name:"NetworkInterfaceDescription"`

	// Bound tags, such as [{"Key": "city", "Value": "shanghai"}].
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// ENI mounting type. Valid values: `0` (standard); `1` (extension); default value: `0`
	AttachType *uint64 `json:"AttachType,omitnil,omitempty" name:"AttachType"`
}

type CreateAndAttachNetworkInterfaceRequest struct {
	*tchttp.BaseRequest
	
	// The ID of the VPC instance. You can obtain the parameter value from the `VpcId` field in the returned result of the `DescribeVpcs` API.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// The name of the ENI. The maximum length is 60 bytes.
	NetworkInterfaceName *string `json:"NetworkInterfaceName,omitnil,omitempty" name:"NetworkInterfaceName"`

	// The subnet instance ID of the ENI, such as 'subnet-0ap8nwca'.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// CVM instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// The information of the specified private IPs. You can specify a maximum of 10 IPs each time.
	PrivateIpAddresses []*PrivateIpAddressSpecification `json:"PrivateIpAddresses,omitnil,omitempty" name:"PrivateIpAddresses"`

	// The number of private IP addresses you apply for. The total number of private IP addresses cannot exceed the quota.
	SecondaryPrivateIpAddressCount *uint64 `json:"SecondaryPrivateIpAddressCount,omitnil,omitempty" name:"SecondaryPrivateIpAddressCount"`

	// IP u200dservice level. It is used together with `SecondaryPrivateIpAddressCount`. Values: `PT` u200d(Gold), `AU` u200d(Silver), `AG` (Bronze) and `DEFAULT` (Default).
	QosLevel *string `json:"QosLevel,omitnil,omitempty" name:"QosLevel"`

	// The security group to be bound with, such as ['sg-1dd51d'].
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// The ENI description. You can enter any information within 60 characters.
	NetworkInterfaceDescription *string `json:"NetworkInterfaceDescription,omitnil,omitempty" name:"NetworkInterfaceDescription"`

	// Bound tags, such as [{"Key": "city", "Value": "shanghai"}].
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// ENI mounting type. Valid values: `0` (standard); `1` (extension); default value: `0`
	AttachType *uint64 `json:"AttachType,omitnil,omitempty" name:"AttachType"`
}

func (r *CreateAndAttachNetworkInterfaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAndAttachNetworkInterfaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcId")
	delete(f, "NetworkInterfaceName")
	delete(f, "SubnetId")
	delete(f, "InstanceId")
	delete(f, "PrivateIpAddresses")
	delete(f, "SecondaryPrivateIpAddressCount")
	delete(f, "QosLevel")
	delete(f, "SecurityGroupIds")
	delete(f, "NetworkInterfaceDescription")
	delete(f, "Tags")
	delete(f, "AttachType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAndAttachNetworkInterfaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAndAttachNetworkInterfaceResponseParams struct {
	// The ENI instance.
	NetworkInterface *NetworkInterface `json:"NetworkInterface,omitnil,omitempty" name:"NetworkInterface"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAndAttachNetworkInterfaceResponse struct {
	*tchttp.BaseResponse
	Response *CreateAndAttachNetworkInterfaceResponseParams `json:"Response"`
}

func (r *CreateAndAttachNetworkInterfaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAndAttachNetworkInterfaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAssistantCidrRequestParams struct {
	// `VPC` instance `ID`, such as `vpc-6v2ht8q5`.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Array of CIDR blocks, such as ["10.0.0.0/16", "172.16.0.0/16"]
	CidrBlocks []*string `json:"CidrBlocks,omitnil,omitempty" name:"CidrBlocks"`
}

type CreateAssistantCidrRequest struct {
	*tchttp.BaseRequest
	
	// `VPC` instance `ID`, such as `vpc-6v2ht8q5`.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Array of CIDR blocks, such as ["10.0.0.0/16", "172.16.0.0/16"]
	CidrBlocks []*string `json:"CidrBlocks,omitnil,omitempty" name:"CidrBlocks"`
}

func (r *CreateAssistantCidrRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAssistantCidrRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcId")
	delete(f, "CidrBlocks")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAssistantCidrRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAssistantCidrResponseParams struct {
	// Array of secondary CIDR blocks.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AssistantCidrSet []*AssistantCidr `json:"AssistantCidrSet,omitnil,omitempty" name:"AssistantCidrSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAssistantCidrResponse struct {
	*tchttp.BaseResponse
	Response *CreateAssistantCidrResponseParams `json:"Response"`
}

func (r *CreateAssistantCidrResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAssistantCidrResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBandwidthPackageRequestParams struct {
	// The network type of the bandwidth package. Default value: `BGP`. Valid values:
	// `BGP` 
	// `HIGH_QUALITY_BGP`
	NetworkType *string `json:"NetworkType,omitnil,omitempty" name:"NetworkType"`

	// The billing mode of the bandwidth package. Default value: `TOP5_POSTPAID_BY_MONTH`. Valid values:
	// <li>`TOP5_POSTPAID_BY_MONTH`: monthly top 5 </li>
	// <li>`PERCENT95_POSTPAID_BY_MONTH`: monthly 95th percentile</li>
	// <li>`FIXED_PREPAID_BY_MONTH`: monthly subscription</li>
	ChargeType *string `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`

	// The name of the bandwidth package.
	BandwidthPackageName *string `json:"BandwidthPackageName,omitnil,omitempty" name:"BandwidthPackageName"`

	// The number of bandwidth packages to create. Valid range: 1-20. It can only be “1” for bill-by-CVM accounts.
	BandwidthPackageCount *uint64 `json:"BandwidthPackageCount,omitnil,omitempty" name:"BandwidthPackageCount"`

	// The limit of the bandwidth package in Mbps. The value '-1' indicates there is no limit. This feature is currently in beta.
	InternetMaxBandwidth *int64 `json:"InternetMaxBandwidth,omitnil,omitempty" name:"InternetMaxBandwidth"`

	// The list of tags to be bound.
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// The protocol type of the bandwidth package. Valid values: 'ipv4' and 'ipv6'. Default value: 'ipv4'.
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`
}

type CreateBandwidthPackageRequest struct {
	*tchttp.BaseRequest
	
	// The network type of the bandwidth package. Default value: `BGP`. Valid values:
	// `BGP` 
	// `HIGH_QUALITY_BGP`
	NetworkType *string `json:"NetworkType,omitnil,omitempty" name:"NetworkType"`

	// The billing mode of the bandwidth package. Default value: `TOP5_POSTPAID_BY_MONTH`. Valid values:
	// <li>`TOP5_POSTPAID_BY_MONTH`: monthly top 5 </li>
	// <li>`PERCENT95_POSTPAID_BY_MONTH`: monthly 95th percentile</li>
	// <li>`FIXED_PREPAID_BY_MONTH`: monthly subscription</li>
	ChargeType *string `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`

	// The name of the bandwidth package.
	BandwidthPackageName *string `json:"BandwidthPackageName,omitnil,omitempty" name:"BandwidthPackageName"`

	// The number of bandwidth packages to create. Valid range: 1-20. It can only be “1” for bill-by-CVM accounts.
	BandwidthPackageCount *uint64 `json:"BandwidthPackageCount,omitnil,omitempty" name:"BandwidthPackageCount"`

	// The limit of the bandwidth package in Mbps. The value '-1' indicates there is no limit. This feature is currently in beta.
	InternetMaxBandwidth *int64 `json:"InternetMaxBandwidth,omitnil,omitempty" name:"InternetMaxBandwidth"`

	// The list of tags to be bound.
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// The protocol type of the bandwidth package. Valid values: 'ipv4' and 'ipv6'. Default value: 'ipv4'.
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`
}

func (r *CreateBandwidthPackageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBandwidthPackageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NetworkType")
	delete(f, "ChargeType")
	delete(f, "BandwidthPackageName")
	delete(f, "BandwidthPackageCount")
	delete(f, "InternetMaxBandwidth")
	delete(f, "Tags")
	delete(f, "Protocol")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBandwidthPackageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBandwidthPackageResponseParams struct {
	// The unique ID of the bandwidth package.
	BandwidthPackageId *string `json:"BandwidthPackageId,omitnil,omitempty" name:"BandwidthPackageId"`

	// The unique ID list of the bandwidth package (effective only when you apply for more than 1 bandwidth packages).
	BandwidthPackageIds []*string `json:"BandwidthPackageIds,omitnil,omitempty" name:"BandwidthPackageIds"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateBandwidthPackageResponse struct {
	*tchttp.BaseResponse
	Response *CreateBandwidthPackageResponseParams `json:"Response"`
}

func (r *CreateBandwidthPackageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBandwidthPackageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCcnRequestParams struct {
	// The name of the CCN. The maximum length is 60 characters.
	CcnName *string `json:"CcnName,omitnil,omitempty" name:"CcnName"`

	// The description of the CCN. The maximum length is 100 characters.
	CcnDescription *string `json:"CcnDescription,omitnil,omitempty" name:"CcnDescription"`

	// CCN service quality, 'PT': Platinum, 'AU': Gold, 'AG': Silver. The default is 'AU'.
	QosLevel *string `json:"QosLevel,omitnil,omitempty" name:"QosLevel"`

	// The billing method. POSTPAID: postpaid by traffic. Default: POSTPAID.
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// The bandwidth limit type. Valid values: OUTER_REGION_LIMIT: region outbound bandwidth limit; INTER_REGION_LIMIT: inter-region bandwidth limit. Default value: OUTER_REGION_LIMIT. Monthly-subscribed CCN instances only support inter-region bandwidth limit, while pay-as-you-go CCN instances support the both bandwidth limit types.
	BandwidthLimitType *string `json:"BandwidthLimitType,omitnil,omitempty" name:"BandwidthLimitType"`

	// Bound tags, such as [{"Key": "city", "Value": "shanghai"}].
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type CreateCcnRequest struct {
	*tchttp.BaseRequest
	
	// The name of the CCN. The maximum length is 60 characters.
	CcnName *string `json:"CcnName,omitnil,omitempty" name:"CcnName"`

	// The description of the CCN. The maximum length is 100 characters.
	CcnDescription *string `json:"CcnDescription,omitnil,omitempty" name:"CcnDescription"`

	// CCN service quality, 'PT': Platinum, 'AU': Gold, 'AG': Silver. The default is 'AU'.
	QosLevel *string `json:"QosLevel,omitnil,omitempty" name:"QosLevel"`

	// The billing method. POSTPAID: postpaid by traffic. Default: POSTPAID.
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// The bandwidth limit type. Valid values: OUTER_REGION_LIMIT: region outbound bandwidth limit; INTER_REGION_LIMIT: inter-region bandwidth limit. Default value: OUTER_REGION_LIMIT. Monthly-subscribed CCN instances only support inter-region bandwidth limit, while pay-as-you-go CCN instances support the both bandwidth limit types.
	BandwidthLimitType *string `json:"BandwidthLimitType,omitnil,omitempty" name:"BandwidthLimitType"`

	// Bound tags, such as [{"Key": "city", "Value": "shanghai"}].
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *CreateCcnRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCcnRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CcnName")
	delete(f, "CcnDescription")
	delete(f, "QosLevel")
	delete(f, "InstanceChargeType")
	delete(f, "BandwidthLimitType")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCcnRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCcnResponseParams struct {
	// The CCN object.
	Ccn *CCN `json:"Ccn,omitnil,omitempty" name:"Ccn"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCcnResponse struct {
	*tchttp.BaseResponse
	Response *CreateCcnResponseParams `json:"Response"`
}

func (r *CreateCcnResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCcnResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCustomerGatewayRequestParams struct {
	// Customer gateway can be named freely, but the maximum length is 60 characters.
	CustomerGatewayName *string `json:"CustomerGatewayName,omitnil,omitempty" name:"CustomerGatewayName"`

	// Customer gateway public IP.
	IpAddress *string `json:"IpAddress,omitnil,omitempty" name:"IpAddress"`

	// Bound tags, such as [{"Key": "city", "Value": "shanghai"}].
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type CreateCustomerGatewayRequest struct {
	*tchttp.BaseRequest
	
	// Customer gateway can be named freely, but the maximum length is 60 characters.
	CustomerGatewayName *string `json:"CustomerGatewayName,omitnil,omitempty" name:"CustomerGatewayName"`

	// Customer gateway public IP.
	IpAddress *string `json:"IpAddress,omitnil,omitempty" name:"IpAddress"`

	// Bound tags, such as [{"Key": "city", "Value": "shanghai"}].
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *CreateCustomerGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCustomerGatewayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CustomerGatewayName")
	delete(f, "IpAddress")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCustomerGatewayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCustomerGatewayResponseParams struct {
	// Customer gateway object
	CustomerGateway *CustomerGateway `json:"CustomerGateway,omitnil,omitempty" name:"CustomerGateway"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCustomerGatewayResponse struct {
	*tchttp.BaseResponse
	Response *CreateCustomerGatewayResponseParams `json:"Response"`
}

func (r *CreateCustomerGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCustomerGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDefaultVpcRequestParams struct {
	// The ID of the availability zone in which the subnet resides. This parameter can be obtained through the [`DescribeZones`](https://intl.cloud.tencent.com/document/product/213/15707?from_cn_redirect=1) API, such as `ap-guangzhou-1`. If it’s not specified, a random availability zone will be used.
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Whether to forcibly return a default VPC
	Force *bool `json:"Force,omitnil,omitempty" name:"Force"`
}

type CreateDefaultVpcRequest struct {
	*tchttp.BaseRequest
	
	// The ID of the availability zone in which the subnet resides. This parameter can be obtained through the [`DescribeZones`](https://intl.cloud.tencent.com/document/product/213/15707?from_cn_redirect=1) API, such as `ap-guangzhou-1`. If it’s not specified, a random availability zone will be used.
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Whether to forcibly return a default VPC
	Force *bool `json:"Force,omitnil,omitempty" name:"Force"`
}

func (r *CreateDefaultVpcRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDefaultVpcRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Zone")
	delete(f, "Force")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDefaultVpcRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDefaultVpcResponseParams struct {
	// Default VPC and subnet IDs
	Vpc *DefaultVpcSubnet `json:"Vpc,omitnil,omitempty" name:"Vpc"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDefaultVpcResponse struct {
	*tchttp.BaseResponse
	Response *CreateDefaultVpcResponseParams `json:"Response"`
}

func (r *CreateDefaultVpcResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDefaultVpcResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDirectConnectGatewayCcnRoutesRequestParams struct {
	// The ID of the Direct Connect gateway, such as `dcg-prpqlmg1`
	DirectConnectGatewayId *string `json:"DirectConnectGatewayId,omitnil,omitempty" name:"DirectConnectGatewayId"`

	// The list of IDC IP range that must be connected
	Routes []*DirectConnectGatewayCcnRoute `json:"Routes,omitnil,omitempty" name:"Routes"`
}

type CreateDirectConnectGatewayCcnRoutesRequest struct {
	*tchttp.BaseRequest
	
	// The ID of the Direct Connect gateway, such as `dcg-prpqlmg1`
	DirectConnectGatewayId *string `json:"DirectConnectGatewayId,omitnil,omitempty" name:"DirectConnectGatewayId"`

	// The list of IDC IP range that must be connected
	Routes []*DirectConnectGatewayCcnRoute `json:"Routes,omitnil,omitempty" name:"Routes"`
}

func (r *CreateDirectConnectGatewayCcnRoutesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDirectConnectGatewayCcnRoutesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DirectConnectGatewayId")
	delete(f, "Routes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDirectConnectGatewayCcnRoutesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDirectConnectGatewayCcnRoutesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDirectConnectGatewayCcnRoutesResponse struct {
	*tchttp.BaseResponse
	Response *CreateDirectConnectGatewayCcnRoutesResponseParams `json:"Response"`
}

func (r *CreateDirectConnectGatewayCcnRoutesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDirectConnectGatewayCcnRoutesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDirectConnectGatewayRequestParams struct {
	// The name of the direct connect gateway.
	DirectConnectGatewayName *string `json:"DirectConnectGatewayName,omitnil,omitempty" name:"DirectConnectGatewayName"`

	// The type of the associated network. Valid values:
	// <li>VPC</li>
	// <li>CCN</li>
	NetworkType *string `json:"NetworkType,omitnil,omitempty" name:"NetworkType"`

	// <li>When the NetworkType is VPC, this value is the VPC instance ID</li>
	// <li>When the NetworkType is CCN, this value is the CCN instance ID</li>
	NetworkInstanceId *string `json:"NetworkInstanceId,omitnil,omitempty" name:"NetworkInstanceId"`

	// The type of the gateway. Valid values:
	// <li>NORMAL - (Default) Standard type. Note: CCN only supports the standard type</li>
	// <li>NAT - NAT type</li>NAT gateway supports network address translation. The specified type cannot be modified. A VPC can create one NAT direct connect gateway and one non-NAT direct connect gateway
	GatewayType *string `json:"GatewayType,omitnil,omitempty" name:"GatewayType"`

	// CCN route publishing method. Valid values: `standard` and `exquisite`. This parameter is only valid for the CCN direct connect gateway.
	ModeType *string `json:"ModeType,omitnil,omitempty" name:"ModeType"`

	// Availability zone where the direct connect gateway resides.
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// ID of DC highly available placement group
	HaZoneGroupId *string `json:"HaZoneGroupId,omitnil,omitempty" name:"HaZoneGroupId"`
}

type CreateDirectConnectGatewayRequest struct {
	*tchttp.BaseRequest
	
	// The name of the direct connect gateway.
	DirectConnectGatewayName *string `json:"DirectConnectGatewayName,omitnil,omitempty" name:"DirectConnectGatewayName"`

	// The type of the associated network. Valid values:
	// <li>VPC</li>
	// <li>CCN</li>
	NetworkType *string `json:"NetworkType,omitnil,omitempty" name:"NetworkType"`

	// <li>When the NetworkType is VPC, this value is the VPC instance ID</li>
	// <li>When the NetworkType is CCN, this value is the CCN instance ID</li>
	NetworkInstanceId *string `json:"NetworkInstanceId,omitnil,omitempty" name:"NetworkInstanceId"`

	// The type of the gateway. Valid values:
	// <li>NORMAL - (Default) Standard type. Note: CCN only supports the standard type</li>
	// <li>NAT - NAT type</li>NAT gateway supports network address translation. The specified type cannot be modified. A VPC can create one NAT direct connect gateway and one non-NAT direct connect gateway
	GatewayType *string `json:"GatewayType,omitnil,omitempty" name:"GatewayType"`

	// CCN route publishing method. Valid values: `standard` and `exquisite`. This parameter is only valid for the CCN direct connect gateway.
	ModeType *string `json:"ModeType,omitnil,omitempty" name:"ModeType"`

	// Availability zone where the direct connect gateway resides.
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// ID of DC highly available placement group
	HaZoneGroupId *string `json:"HaZoneGroupId,omitnil,omitempty" name:"HaZoneGroupId"`
}

func (r *CreateDirectConnectGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDirectConnectGatewayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DirectConnectGatewayName")
	delete(f, "NetworkType")
	delete(f, "NetworkInstanceId")
	delete(f, "GatewayType")
	delete(f, "ModeType")
	delete(f, "Zone")
	delete(f, "HaZoneGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDirectConnectGatewayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDirectConnectGatewayResponseParams struct {
	// The object of the direct connect gateway.
	DirectConnectGateway *DirectConnectGateway `json:"DirectConnectGateway,omitnil,omitempty" name:"DirectConnectGateway"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDirectConnectGatewayResponse struct {
	*tchttp.BaseResponse
	Response *CreateDirectConnectGatewayResponseParams `json:"Response"`
}

func (r *CreateDirectConnectGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDirectConnectGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFlowLogRequestParams struct {
	// The name of the flow log instance.
	FlowLogName *string `json:"FlowLogName,omitnil,omitempty" name:"FlowLogName"`

	// The type of resource associated with the flow log. Valid values: `VPC`, `SUBNET`, `NETWORKINTERFACE`, `CCN`, `NAT`, and `DCG`.
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// The unique ID of the resource.
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// Type of the flow logs to be collected. Valid values: `ACCEPT`, `REJECT` and `ALL`.
	TrafficType *string `json:"TrafficType,omitnil,omitempty" name:"TrafficType"`

	// The VPC ID or unique ID of the resource. We recommend using the unique ID. This parameter is required unless the `ResourceType` is set to `CCN`.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// The description of the flow log.
	FlowLogDescription *string `json:"FlowLogDescription,omitnil,omitempty" name:"FlowLogDescription"`

	// The storage ID of the flow log.
	CloudLogId *string `json:"CloudLogId,omitnil,omitempty" name:"CloudLogId"`

	// Bound tags, such as [{"Key": "city", "Value": "shanghai"}].
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Consumer types: `cls` and `ckafka`
	StorageType *string `json:"StorageType,omitnil,omitempty" name:"StorageType"`

	// Information of the flow log consumer, which is required when the consumer type is `ckafka`.
	FlowLogStorage *FlowLogStorage `json:"FlowLogStorage,omitnil,omitempty" name:"FlowLogStorage"`

	// The region corresponding to the flow log storage ID. If not passed in, this field defaults to the current region.
	CloudLogRegion *string `json:"CloudLogRegion,omitnil,omitempty" name:"CloudLogRegion"`
}

type CreateFlowLogRequest struct {
	*tchttp.BaseRequest
	
	// The name of the flow log instance.
	FlowLogName *string `json:"FlowLogName,omitnil,omitempty" name:"FlowLogName"`

	// The type of resource associated with the flow log. Valid values: `VPC`, `SUBNET`, `NETWORKINTERFACE`, `CCN`, `NAT`, and `DCG`.
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// The unique ID of the resource.
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// Type of the flow logs to be collected. Valid values: `ACCEPT`, `REJECT` and `ALL`.
	TrafficType *string `json:"TrafficType,omitnil,omitempty" name:"TrafficType"`

	// The VPC ID or unique ID of the resource. We recommend using the unique ID. This parameter is required unless the `ResourceType` is set to `CCN`.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// The description of the flow log.
	FlowLogDescription *string `json:"FlowLogDescription,omitnil,omitempty" name:"FlowLogDescription"`

	// The storage ID of the flow log.
	CloudLogId *string `json:"CloudLogId,omitnil,omitempty" name:"CloudLogId"`

	// Bound tags, such as [{"Key": "city", "Value": "shanghai"}].
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Consumer types: `cls` and `ckafka`
	StorageType *string `json:"StorageType,omitnil,omitempty" name:"StorageType"`

	// Information of the flow log consumer, which is required when the consumer type is `ckafka`.
	FlowLogStorage *FlowLogStorage `json:"FlowLogStorage,omitnil,omitempty" name:"FlowLogStorage"`

	// The region corresponding to the flow log storage ID. If not passed in, this field defaults to the current region.
	CloudLogRegion *string `json:"CloudLogRegion,omitnil,omitempty" name:"CloudLogRegion"`
}

func (r *CreateFlowLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFlowLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FlowLogName")
	delete(f, "ResourceType")
	delete(f, "ResourceId")
	delete(f, "TrafficType")
	delete(f, "VpcId")
	delete(f, "FlowLogDescription")
	delete(f, "CloudLogId")
	delete(f, "Tags")
	delete(f, "StorageType")
	delete(f, "FlowLogStorage")
	delete(f, "CloudLogRegion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateFlowLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFlowLogResponseParams struct {
	// The information of the flow log created.
	FlowLog []*FlowLog `json:"FlowLog,omitnil,omitempty" name:"FlowLog"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateFlowLogResponse struct {
	*tchttp.BaseResponse
	Response *CreateFlowLogResponseParams `json:"Response"`
}

func (r *CreateFlowLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFlowLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateHaVipRequestParams struct {
	// The `ID` of the VPC to which the `HAVIP` belongs.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// The `ID` of the subnet to which the `HAVIP` belongs.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// The name of the `HAVIP`.
	HaVipName *string `json:"HaVipName,omitnil,omitempty" name:"HaVipName"`

	// The specified virtual IP address, which must be within the IP range of the `VPC` and not in use. It will be automatically assigned if not specified.
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// The ID of the ENI associated with the HAVIP.
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitnil,omitempty" name:"NetworkInterfaceId"`
}

type CreateHaVipRequest struct {
	*tchttp.BaseRequest
	
	// The `ID` of the VPC to which the `HAVIP` belongs.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// The `ID` of the subnet to which the `HAVIP` belongs.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// The name of the `HAVIP`.
	HaVipName *string `json:"HaVipName,omitnil,omitempty" name:"HaVipName"`

	// The specified virtual IP address, which must be within the IP range of the `VPC` and not in use. It will be automatically assigned if not specified.
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// The ID of the ENI associated with the HAVIP.
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitnil,omitempty" name:"NetworkInterfaceId"`
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
	delete(f, "NetworkInterfaceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateHaVipRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateHaVipResponseParams struct {
	// `HAVIP` object.
	HaVip *HaVip `json:"HaVip,omitnil,omitempty" name:"HaVip"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type CreateLocalGatewayRequestParams struct {
	// Local gateway name
	LocalGatewayName *string `json:"LocalGatewayName,omitnil,omitempty" name:"LocalGatewayName"`

	// VPC instance ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// CDC instance ID
	CdcId *string `json:"CdcId,omitnil,omitempty" name:"CdcId"`
}

type CreateLocalGatewayRequest struct {
	*tchttp.BaseRequest
	
	// Local gateway name
	LocalGatewayName *string `json:"LocalGatewayName,omitnil,omitempty" name:"LocalGatewayName"`

	// VPC instance ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// CDC instance ID
	CdcId *string `json:"CdcId,omitnil,omitempty" name:"CdcId"`
}

func (r *CreateLocalGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLocalGatewayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LocalGatewayName")
	delete(f, "VpcId")
	delete(f, "CdcId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateLocalGatewayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLocalGatewayResponseParams struct {
	// Local gateway information
	LocalGateway *LocalGateway `json:"LocalGateway,omitnil,omitempty" name:"LocalGateway"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateLocalGatewayResponse struct {
	*tchttp.BaseResponse
	Response *CreateLocalGatewayResponseParams `json:"Response"`
}

func (r *CreateLocalGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLocalGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNatGatewayDestinationIpPortTranslationNatRuleRequestParams struct {
	// The ID of the NAT gateway, such as `nat-df45454`.
	NatGatewayId *string `json:"NatGatewayId,omitnil,omitempty" name:"NatGatewayId"`

	// The port forwarding rules of the NAT gateway.
	DestinationIpPortTranslationNatRules []*DestinationIpPortTranslationNatRule `json:"DestinationIpPortTranslationNatRules,omitnil,omitempty" name:"DestinationIpPortTranslationNatRules"`
}

type CreateNatGatewayDestinationIpPortTranslationNatRuleRequest struct {
	*tchttp.BaseRequest
	
	// The ID of the NAT gateway, such as `nat-df45454`.
	NatGatewayId *string `json:"NatGatewayId,omitnil,omitempty" name:"NatGatewayId"`

	// The port forwarding rules of the NAT gateway.
	DestinationIpPortTranslationNatRules []*DestinationIpPortTranslationNatRule `json:"DestinationIpPortTranslationNatRules,omitnil,omitempty" name:"DestinationIpPortTranslationNatRules"`
}

func (r *CreateNatGatewayDestinationIpPortTranslationNatRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNatGatewayDestinationIpPortTranslationNatRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NatGatewayId")
	delete(f, "DestinationIpPortTranslationNatRules")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateNatGatewayDestinationIpPortTranslationNatRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNatGatewayDestinationIpPortTranslationNatRuleResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateNatGatewayDestinationIpPortTranslationNatRuleResponse struct {
	*tchttp.BaseResponse
	Response *CreateNatGatewayDestinationIpPortTranslationNatRuleResponseParams `json:"Response"`
}

func (r *CreateNatGatewayDestinationIpPortTranslationNatRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNatGatewayDestinationIpPortTranslationNatRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNatGatewayRequestParams struct {
	// NAT gateway name
	NatGatewayName *string `json:"NatGatewayName,omitnil,omitempty" name:"NatGatewayName"`

	// The ID of the VPC instance, which can be obtained from the `VpcId` field in response of the `DescribeVpcs` API.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// The maximum outbound bandwidth of the NAT gateway (unit: Mbps). Supported parameter values: `20, 50, 100, 200, 500, 1000, 2000, 5000`. Default: `100`.
	InternetMaxBandwidthOut *uint64 `json:"InternetMaxBandwidthOut,omitnil,omitempty" name:"InternetMaxBandwidthOut"`

	// The concurrent connection cap of the NAT gateway. Values: `1000000, 3000000, 10000000`. The default value is `1000000`.
	MaxConcurrentConnection *uint64 `json:"MaxConcurrentConnection,omitnil,omitempty" name:"MaxConcurrentConnection"`

	// The number of EIPs that you want to apply for. Either `AddressCount` or `PublicIpAddresses` must be passed in.
	AddressCount *uint64 `json:"AddressCount,omitnil,omitempty" name:"AddressCount"`

	// The EIP array bound to the NAT gateway. Either AddressCount or PublicIpAddresses must be passed in.
	PublicIpAddresses []*string `json:"PublicIpAddresses,omitnil,omitempty" name:"PublicIpAddresses"`

	// The availability zone, such as `ap-guangzhou-1`.
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Bound tags, such as [{"Key": "city", "Value": "shanghai"}]
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Subnet of the NAT gateway
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// The bandwidth size (in Mbps) of the EIP bound to the NAT gateway, which defaults to the maximum value applicable for the current user type.
	StockPublicIpAddressesBandwidthOut *uint64 `json:"StockPublicIpAddressesBandwidthOut,omitnil,omitempty" name:"StockPublicIpAddressesBandwidthOut"`

	// The size of the public network IP bandwidth to be applied for (in Mbps), which defaults to the maximum value applicable for the current user type.
	PublicIpAddressesBandwidthOut *uint64 `json:"PublicIpAddressesBandwidthOut,omitnil,omitempty" name:"PublicIpAddressesBandwidthOut"`

	// Whether the public IP and the NAT gateway must be in the same availability zone. Valid values: `true` and `false`. This parameter is valid only when `Zone` is specified.
	PublicIpFromSameZone *bool `json:"PublicIpFromSameZone,omitnil,omitempty" name:"PublicIpFromSameZone"`


	NatProductVersion *uint64 `json:"NatProductVersion,omitnil,omitempty" name:"NatProductVersion"`
}

type CreateNatGatewayRequest struct {
	*tchttp.BaseRequest
	
	// NAT gateway name
	NatGatewayName *string `json:"NatGatewayName,omitnil,omitempty" name:"NatGatewayName"`

	// The ID of the VPC instance, which can be obtained from the `VpcId` field in response of the `DescribeVpcs` API.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// The maximum outbound bandwidth of the NAT gateway (unit: Mbps). Supported parameter values: `20, 50, 100, 200, 500, 1000, 2000, 5000`. Default: `100`.
	InternetMaxBandwidthOut *uint64 `json:"InternetMaxBandwidthOut,omitnil,omitempty" name:"InternetMaxBandwidthOut"`

	// The concurrent connection cap of the NAT gateway. Values: `1000000, 3000000, 10000000`. The default value is `1000000`.
	MaxConcurrentConnection *uint64 `json:"MaxConcurrentConnection,omitnil,omitempty" name:"MaxConcurrentConnection"`

	// The number of EIPs that you want to apply for. Either `AddressCount` or `PublicIpAddresses` must be passed in.
	AddressCount *uint64 `json:"AddressCount,omitnil,omitempty" name:"AddressCount"`

	// The EIP array bound to the NAT gateway. Either AddressCount or PublicIpAddresses must be passed in.
	PublicIpAddresses []*string `json:"PublicIpAddresses,omitnil,omitempty" name:"PublicIpAddresses"`

	// The availability zone, such as `ap-guangzhou-1`.
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Bound tags, such as [{"Key": "city", "Value": "shanghai"}]
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Subnet of the NAT gateway
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// The bandwidth size (in Mbps) of the EIP bound to the NAT gateway, which defaults to the maximum value applicable for the current user type.
	StockPublicIpAddressesBandwidthOut *uint64 `json:"StockPublicIpAddressesBandwidthOut,omitnil,omitempty" name:"StockPublicIpAddressesBandwidthOut"`

	// The size of the public network IP bandwidth to be applied for (in Mbps), which defaults to the maximum value applicable for the current user type.
	PublicIpAddressesBandwidthOut *uint64 `json:"PublicIpAddressesBandwidthOut,omitnil,omitempty" name:"PublicIpAddressesBandwidthOut"`

	// Whether the public IP and the NAT gateway must be in the same availability zone. Valid values: `true` and `false`. This parameter is valid only when `Zone` is specified.
	PublicIpFromSameZone *bool `json:"PublicIpFromSameZone,omitnil,omitempty" name:"PublicIpFromSameZone"`

	NatProductVersion *uint64 `json:"NatProductVersion,omitnil,omitempty" name:"NatProductVersion"`
}

func (r *CreateNatGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNatGatewayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NatGatewayName")
	delete(f, "VpcId")
	delete(f, "InternetMaxBandwidthOut")
	delete(f, "MaxConcurrentConnection")
	delete(f, "AddressCount")
	delete(f, "PublicIpAddresses")
	delete(f, "Zone")
	delete(f, "Tags")
	delete(f, "SubnetId")
	delete(f, "StockPublicIpAddressesBandwidthOut")
	delete(f, "PublicIpAddressesBandwidthOut")
	delete(f, "PublicIpFromSameZone")
	delete(f, "NatProductVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateNatGatewayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNatGatewayResponseParams struct {
	// NAT gateway object array.
	NatGatewaySet []*NatGateway `json:"NatGatewaySet,omitnil,omitempty" name:"NatGatewaySet"`

	// The number of eligible NAT gateway objects.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateNatGatewayResponse struct {
	*tchttp.BaseResponse
	Response *CreateNatGatewayResponseParams `json:"Response"`
}

func (r *CreateNatGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNatGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNatGatewaySourceIpTranslationNatRuleRequestParams struct {
	// The ID of the NAT Gateway, such as `nat-df45454`
	NatGatewayId *string `json:"NatGatewayId,omitnil,omitempty" name:"NatGatewayId"`

	// The SNAT forwarding rule of the NAT Gateway
	SourceIpTranslationNatRules []*SourceIpTranslationNatRule `json:"SourceIpTranslationNatRules,omitnil,omitempty" name:"SourceIpTranslationNatRules"`
}

type CreateNatGatewaySourceIpTranslationNatRuleRequest struct {
	*tchttp.BaseRequest
	
	// The ID of the NAT Gateway, such as `nat-df45454`
	NatGatewayId *string `json:"NatGatewayId,omitnil,omitempty" name:"NatGatewayId"`

	// The SNAT forwarding rule of the NAT Gateway
	SourceIpTranslationNatRules []*SourceIpTranslationNatRule `json:"SourceIpTranslationNatRules,omitnil,omitempty" name:"SourceIpTranslationNatRules"`
}

func (r *CreateNatGatewaySourceIpTranslationNatRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNatGatewaySourceIpTranslationNatRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NatGatewayId")
	delete(f, "SourceIpTranslationNatRules")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateNatGatewaySourceIpTranslationNatRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNatGatewaySourceIpTranslationNatRuleResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateNatGatewaySourceIpTranslationNatRuleResponse struct {
	*tchttp.BaseResponse
	Response *CreateNatGatewaySourceIpTranslationNatRuleResponseParams `json:"Response"`
}

func (r *CreateNatGatewaySourceIpTranslationNatRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNatGatewaySourceIpTranslationNatRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNetDetectRequestParams struct {
	// The ID of a VPC instance, such as `vpc-12345678`.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// The ID of a subnet instance, such as subnet-12345678.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// The name of a network detection instance. The maximum length is 60 characters.
	NetDetectName *string `json:"NetDetectName,omitnil,omitempty" name:"NetDetectName"`

	// The array of detection destination IPv4 addresses, which contains at most two IP addresses.
	DetectDestinationIp []*string `json:"DetectDestinationIp,omitnil,omitempty" name:"DetectDestinationIp"`

	// Type of the next hop. Valid values:
	// `VPN`: VPN gateway;
	// `DIRECTCONNECT`: Direct connect gateway;
	// `PEERCONNECTION`: Peering connection;
	// `NAT`: NAT gateway;
	// `NORMAL_CVM`: CVM instance;
	// `CCN`: CCN instance;
	// `NONEXTHOP`: No next hop.
	NextHopType *string `json:"NextHopType,omitnil,omitempty" name:"NextHopType"`

	// ID of the next-hop gateway.
	// If NextHopType is set to VPN, the value of this parameter is the VPN gateway ID, such as vpngw-12345678.
	// `NextHopType` = `DIRECTCONNECT`: Direct connect gateway ID, such as `dcg-12345678`.
	// `NextHopType` = `PEERCONNECTION`: Peering connection ID, such as `pcx-12345678`.
	// `NextHopType` = `NAT`: NAT gateway ID, such as `nat-12345678`.
	// If NextHopType is set to NORMAL_CVM, the value of this parameter is the IPv4 address of the CVM, such as 10.0.0.12.
	// `NextHopType` = `CCN`: CCN instance ID, such as `ccn-12345678`.
	// `NextHopType` = `NONEXTHOP`: No next hop.
	NextHopDestination *string `json:"NextHopDestination,omitnil,omitempty" name:"NextHopDestination"`

	// Network detection description.
	NetDetectDescription *string `json:"NetDetectDescription,omitnil,omitempty" name:"NetDetectDescription"`
}

type CreateNetDetectRequest struct {
	*tchttp.BaseRequest
	
	// The ID of a VPC instance, such as `vpc-12345678`.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// The ID of a subnet instance, such as subnet-12345678.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// The name of a network detection instance. The maximum length is 60 characters.
	NetDetectName *string `json:"NetDetectName,omitnil,omitempty" name:"NetDetectName"`

	// The array of detection destination IPv4 addresses, which contains at most two IP addresses.
	DetectDestinationIp []*string `json:"DetectDestinationIp,omitnil,omitempty" name:"DetectDestinationIp"`

	// Type of the next hop. Valid values:
	// `VPN`: VPN gateway;
	// `DIRECTCONNECT`: Direct connect gateway;
	// `PEERCONNECTION`: Peering connection;
	// `NAT`: NAT gateway;
	// `NORMAL_CVM`: CVM instance;
	// `CCN`: CCN instance;
	// `NONEXTHOP`: No next hop.
	NextHopType *string `json:"NextHopType,omitnil,omitempty" name:"NextHopType"`

	// ID of the next-hop gateway.
	// If NextHopType is set to VPN, the value of this parameter is the VPN gateway ID, such as vpngw-12345678.
	// `NextHopType` = `DIRECTCONNECT`: Direct connect gateway ID, such as `dcg-12345678`.
	// `NextHopType` = `PEERCONNECTION`: Peering connection ID, such as `pcx-12345678`.
	// `NextHopType` = `NAT`: NAT gateway ID, such as `nat-12345678`.
	// If NextHopType is set to NORMAL_CVM, the value of this parameter is the IPv4 address of the CVM, such as 10.0.0.12.
	// `NextHopType` = `CCN`: CCN instance ID, such as `ccn-12345678`.
	// `NextHopType` = `NONEXTHOP`: No next hop.
	NextHopDestination *string `json:"NextHopDestination,omitnil,omitempty" name:"NextHopDestination"`

	// Network detection description.
	NetDetectDescription *string `json:"NetDetectDescription,omitnil,omitempty" name:"NetDetectDescription"`
}

func (r *CreateNetDetectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNetDetectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "NetDetectName")
	delete(f, "DetectDestinationIp")
	delete(f, "NextHopType")
	delete(f, "NextHopDestination")
	delete(f, "NetDetectDescription")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateNetDetectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNetDetectResponseParams struct {
	// The network detection (NetDetect) object.
	NetDetect *NetDetect `json:"NetDetect,omitnil,omitempty" name:"NetDetect"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateNetDetectResponse struct {
	*tchttp.BaseResponse
	Response *CreateNetDetectResponseParams `json:"Response"`
}

func (r *CreateNetDetectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNetDetectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNetworkAclQuintupleEntriesRequestParams struct {
	// Network ACL instance ID, such as `acl-12345678`.
	NetworkAclId *string `json:"NetworkAclId,omitnil,omitempty" name:"NetworkAclId"`

	// Network ACL quintuple rule set.
	NetworkAclQuintupleSet *NetworkAclQuintupleEntries `json:"NetworkAclQuintupleSet,omitnil,omitempty" name:"NetworkAclQuintupleSet"`
}

type CreateNetworkAclQuintupleEntriesRequest struct {
	*tchttp.BaseRequest
	
	// Network ACL instance ID, such as `acl-12345678`.
	NetworkAclId *string `json:"NetworkAclId,omitnil,omitempty" name:"NetworkAclId"`

	// Network ACL quintuple rule set.
	NetworkAclQuintupleSet *NetworkAclQuintupleEntries `json:"NetworkAclQuintupleSet,omitnil,omitempty" name:"NetworkAclQuintupleSet"`
}

func (r *CreateNetworkAclQuintupleEntriesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNetworkAclQuintupleEntriesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NetworkAclId")
	delete(f, "NetworkAclQuintupleSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateNetworkAclQuintupleEntriesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNetworkAclQuintupleEntriesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateNetworkAclQuintupleEntriesResponse struct {
	*tchttp.BaseResponse
	Response *CreateNetworkAclQuintupleEntriesResponseParams `json:"Response"`
}

func (r *CreateNetworkAclQuintupleEntriesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNetworkAclQuintupleEntriesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNetworkAclRequestParams struct {
	// VPC instance ID, which can be obtained from the `VpcId` field in the response of the [`DescribeVpcs`](https://intl.cloud.tencent.com/document/product/215/15778?from_cn_redirect=1) API.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Network ACL name, which can contain up to 60 bytes.
	NetworkAclName *string `json:"NetworkAclName,omitnil,omitempty" name:"NetworkAclName"`

	// Network ACL type. Valid values: `TRIPLE` and `QUINTUPLE`.
	NetworkAclType *string `json:"NetworkAclType,omitnil,omitempty" name:"NetworkAclType"`

	// Bound tags, such as [{"Key": "city", "Value": "shanghai"}].
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type CreateNetworkAclRequest struct {
	*tchttp.BaseRequest
	
	// VPC instance ID, which can be obtained from the `VpcId` field in the response of the [`DescribeVpcs`](https://intl.cloud.tencent.com/document/product/215/15778?from_cn_redirect=1) API.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Network ACL name, which can contain up to 60 bytes.
	NetworkAclName *string `json:"NetworkAclName,omitnil,omitempty" name:"NetworkAclName"`

	// Network ACL type. Valid values: `TRIPLE` and `QUINTUPLE`.
	NetworkAclType *string `json:"NetworkAclType,omitnil,omitempty" name:"NetworkAclType"`

	// Bound tags, such as [{"Key": "city", "Value": "shanghai"}].
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *CreateNetworkAclRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNetworkAclRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcId")
	delete(f, "NetworkAclName")
	delete(f, "NetworkAclType")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateNetworkAclRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNetworkAclResponseParams struct {
	// Network ACL instance
	NetworkAcl *NetworkAcl `json:"NetworkAcl,omitnil,omitempty" name:"NetworkAcl"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateNetworkAclResponse struct {
	*tchttp.BaseResponse
	Response *CreateNetworkAclResponseParams `json:"Response"`
}

func (r *CreateNetworkAclResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNetworkAclResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNetworkInterfaceRequestParams struct {
	// The ID of the VPC instance. You can obtain the parameter value from the VpcId field in the returned result of DescribeVpcs API.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// The name of the ENI. The maximum length is 60 characters.
	NetworkInterfaceName *string `json:"NetworkInterfaceName,omitnil,omitempty" name:"NetworkInterfaceName"`

	// The subnet instance ID of the ENI, such as `subnet-0ap8nwca`.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// ENI description can be named freely, but the maximum length is 60 characters.
	NetworkInterfaceDescription *string `json:"NetworkInterfaceDescription,omitnil,omitempty" name:"NetworkInterfaceDescription"`

	// The number of private IP addresses you apply for. The total number of private IP addresses cannot exceed the quota.
	SecondaryPrivateIpAddressCount *uint64 `json:"SecondaryPrivateIpAddressCount,omitnil,omitempty" name:"SecondaryPrivateIpAddressCount"`

	// IP u200dservice level. It is used together with `SecondaryPrivateIpAddressCount`. Values: `PT` u200d(Gold), `AU` u200d(Silver), `AG` (Bronze) and `DEFAULT` (Default).
	QosLevel *string `json:"QosLevel,omitnil,omitempty" name:"QosLevel"`

	// Specifies the security group to be bound with, such as ['sg-1dd51d'].
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// The information of the specified private IPs. You can specify a maximum of 10 each time.
	PrivateIpAddresses []*PrivateIpAddressSpecification `json:"PrivateIpAddresses,omitnil,omitempty" name:"PrivateIpAddresses"`

	// Bound tags, such as [{"Key": "city", "Value": "shanghai"}].
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Configuration of the ENI trunking mode. Valid values: `Enable` and `Disable`. Default value: `Disable`.
	TrunkingFlag *string `json:"TrunkingFlag,omitnil,omitempty" name:"TrunkingFlag"`
}

type CreateNetworkInterfaceRequest struct {
	*tchttp.BaseRequest
	
	// The ID of the VPC instance. You can obtain the parameter value from the VpcId field in the returned result of DescribeVpcs API.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// The name of the ENI. The maximum length is 60 characters.
	NetworkInterfaceName *string `json:"NetworkInterfaceName,omitnil,omitempty" name:"NetworkInterfaceName"`

	// The subnet instance ID of the ENI, such as `subnet-0ap8nwca`.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// ENI description can be named freely, but the maximum length is 60 characters.
	NetworkInterfaceDescription *string `json:"NetworkInterfaceDescription,omitnil,omitempty" name:"NetworkInterfaceDescription"`

	// The number of private IP addresses you apply for. The total number of private IP addresses cannot exceed the quota.
	SecondaryPrivateIpAddressCount *uint64 `json:"SecondaryPrivateIpAddressCount,omitnil,omitempty" name:"SecondaryPrivateIpAddressCount"`

	// IP u200dservice level. It is used together with `SecondaryPrivateIpAddressCount`. Values: `PT` u200d(Gold), `AU` u200d(Silver), `AG` (Bronze) and `DEFAULT` (Default).
	QosLevel *string `json:"QosLevel,omitnil,omitempty" name:"QosLevel"`

	// Specifies the security group to be bound with, such as ['sg-1dd51d'].
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// The information of the specified private IPs. You can specify a maximum of 10 each time.
	PrivateIpAddresses []*PrivateIpAddressSpecification `json:"PrivateIpAddresses,omitnil,omitempty" name:"PrivateIpAddresses"`

	// Bound tags, such as [{"Key": "city", "Value": "shanghai"}].
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Configuration of the ENI trunking mode. Valid values: `Enable` and `Disable`. Default value: `Disable`.
	TrunkingFlag *string `json:"TrunkingFlag,omitnil,omitempty" name:"TrunkingFlag"`
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
	delete(f, "NetworkInterfaceDescription")
	delete(f, "SecondaryPrivateIpAddressCount")
	delete(f, "QosLevel")
	delete(f, "SecurityGroupIds")
	delete(f, "PrivateIpAddresses")
	delete(f, "Tags")
	delete(f, "TrunkingFlag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateNetworkInterfaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNetworkInterfaceResponseParams struct {
	// ENI instance.
	NetworkInterface *NetworkInterface `json:"NetworkInterface,omitnil,omitempty" name:"NetworkInterface"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// The ID of the VPC instance to be operated on. You can obtain the parameter value from the VpcId field in the returned result of DescribeVpcs API.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// The route table name. The maximum length is 60 characters.
	RouteTableName *string `json:"RouteTableName,omitnil,omitempty" name:"RouteTableName"`

	// Bound tags, such as [{"Key": "city", "Value": "shanghai"}].
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type CreateRouteTableRequest struct {
	*tchttp.BaseRequest
	
	// The ID of the VPC instance to be operated on. You can obtain the parameter value from the VpcId field in the returned result of DescribeVpcs API.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// The route table name. The maximum length is 60 characters.
	RouteTableName *string `json:"RouteTableName,omitnil,omitempty" name:"RouteTableName"`

	// Bound tags, such as [{"Key": "city", "Value": "shanghai"}].
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
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
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRouteTableRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRouteTableResponseParams struct {
	// Route table object.
	RouteTable *RouteTable `json:"RouteTable,omitnil,omitempty" name:"RouteTable"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	RouteTableId *string `json:"RouteTableId,omitnil,omitempty" name:"RouteTableId"`

	// Routing policy object.
	Routes []*Route `json:"Routes,omitnil,omitempty" name:"Routes"`
}

type CreateRoutesRequest struct {
	*tchttp.BaseRequest
	
	// Route table instance ID.
	RouteTableId *string `json:"RouteTableId,omitnil,omitempty" name:"RouteTableId"`

	// Routing policy object.
	Routes []*Route `json:"Routes,omitnil,omitempty" name:"Routes"`
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
	// The number of newly added instances.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Route table object.
	RouteTableSet []*RouteTable `json:"RouteTableSet,omitnil,omitempty" name:"RouteTableSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// The security group instance ID, such as `sg-33ocnj9n`. This can be obtained through DescribeSecurityGroups.
	SecurityGroupId *string `json:"SecurityGroupId,omitnil,omitempty" name:"SecurityGroupId"`

	// Security group policy set.
	SecurityGroupPolicySet *SecurityGroupPolicySet `json:"SecurityGroupPolicySet,omitnil,omitempty" name:"SecurityGroupPolicySet"`
}

type CreateSecurityGroupPoliciesRequest struct {
	*tchttp.BaseRequest
	
	// The security group instance ID, such as `sg-33ocnj9n`. This can be obtained through DescribeSecurityGroups.
	SecurityGroupId *string `json:"SecurityGroupId,omitnil,omitempty" name:"SecurityGroupId"`

	// Security group policy set.
	SecurityGroupPolicySet *SecurityGroupPolicySet `json:"SecurityGroupPolicySet,omitnil,omitempty" name:"SecurityGroupPolicySet"`
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
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// Security group can be named freely, but cannot exceed 60 characters.
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// The remarks for the security group. The maximum length is 100 characters.
	GroupDescription *string `json:"GroupDescription,omitnil,omitempty" name:"GroupDescription"`

	// Project ID. The default is 0. You can query it on the project management page of the Qcloud console.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Bound tags, such as [{"Key": "city", "Value": "shanghai"}].
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type CreateSecurityGroupRequest struct {
	*tchttp.BaseRequest
	
	// Security group can be named freely, but cannot exceed 60 characters.
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// The remarks for the security group. The maximum length is 100 characters.
	GroupDescription *string `json:"GroupDescription,omitnil,omitempty" name:"GroupDescription"`

	// Project ID. The default is 0. You can query it on the project management page of the Qcloud console.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Bound tags, such as [{"Key": "city", "Value": "shanghai"}].
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
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
	delete(f, "ProjectId")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSecurityGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSecurityGroupResponseParams struct {
	// Security group object.
	SecurityGroup *SecurityGroup `json:"SecurityGroup,omitnil,omitempty" name:"SecurityGroup"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type CreateSecurityGroupWithPoliciesRequestParams struct {
	// Security group can be named freely, but cannot exceed 60 characters.
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// The remarks for the security group. The maximum length is 100 characters.
	GroupDescription *string `json:"GroupDescription,omitnil,omitempty" name:"GroupDescription"`

	// Project ID. Default value: 0. You can query it on the <a href="https://console.cloud.tencent.com/project">project management page</a> of the Tencent Cloud console.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Security group policy set.
	SecurityGroupPolicySet *SecurityGroupPolicySet `json:"SecurityGroupPolicySet,omitnil,omitempty" name:"SecurityGroupPolicySet"`
}

type CreateSecurityGroupWithPoliciesRequest struct {
	*tchttp.BaseRequest
	
	// Security group can be named freely, but cannot exceed 60 characters.
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// The remarks for the security group. The maximum length is 100 characters.
	GroupDescription *string `json:"GroupDescription,omitnil,omitempty" name:"GroupDescription"`

	// Project ID. Default value: 0. You can query it on the <a href="https://console.cloud.tencent.com/project">project management page</a> of the Tencent Cloud console.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Security group policy set.
	SecurityGroupPolicySet *SecurityGroupPolicySet `json:"SecurityGroupPolicySet,omitnil,omitempty" name:"SecurityGroupPolicySet"`
}

func (r *CreateSecurityGroupWithPoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSecurityGroupWithPoliciesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupName")
	delete(f, "GroupDescription")
	delete(f, "ProjectId")
	delete(f, "SecurityGroupPolicySet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSecurityGroupWithPoliciesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSecurityGroupWithPoliciesResponseParams struct {
	// Security group object.
	SecurityGroup *SecurityGroup `json:"SecurityGroup,omitnil,omitempty" name:"SecurityGroup"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateSecurityGroupWithPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *CreateSecurityGroupWithPoliciesResponseParams `json:"Response"`
}

func (r *CreateSecurityGroupWithPoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSecurityGroupWithPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateServiceTemplateGroupRequestParams struct {
	// Group name of the protocol port template.
	ServiceTemplateGroupName *string `json:"ServiceTemplateGroupName,omitnil,omitempty" name:"ServiceTemplateGroupName"`

	// Instance ID of the protocol port template, such as `ppm-4dw6agho`.
	ServiceTemplateIds []*string `json:"ServiceTemplateIds,omitnil,omitempty" name:"ServiceTemplateIds"`
}

type CreateServiceTemplateGroupRequest struct {
	*tchttp.BaseRequest
	
	// Group name of the protocol port template.
	ServiceTemplateGroupName *string `json:"ServiceTemplateGroupName,omitnil,omitempty" name:"ServiceTemplateGroupName"`

	// Instance ID of the protocol port template, such as `ppm-4dw6agho`.
	ServiceTemplateIds []*string `json:"ServiceTemplateIds,omitnil,omitempty" name:"ServiceTemplateIds"`
}

func (r *CreateServiceTemplateGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateServiceTemplateGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceTemplateGroupName")
	delete(f, "ServiceTemplateIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateServiceTemplateGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateServiceTemplateGroupResponseParams struct {
	// Group object of the protocol port template.
	ServiceTemplateGroup *ServiceTemplateGroup `json:"ServiceTemplateGroup,omitnil,omitempty" name:"ServiceTemplateGroup"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateServiceTemplateGroupResponse struct {
	*tchttp.BaseResponse
	Response *CreateServiceTemplateGroupResponseParams `json:"Response"`
}

func (r *CreateServiceTemplateGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateServiceTemplateGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateServiceTemplateRequestParams struct {
	// Template name of the protocol port
	ServiceTemplateName *string `json:"ServiceTemplateName,omitnil,omitempty" name:"ServiceTemplateName"`

	// Supported ports inlcude single port, multiple ports, consecutive ports and all ports. Supported protocols include TCP, UDP, ICMP and GRE. Either Services or ServicesExtra is required.
	Services []*string `json:"Services,omitnil,omitempty" name:"Services"`

	// You can add remarks. Supported ports include single port, multiple ports, consecutive ports and all ports. Supported protocols include TCP, UDP, ICMP and GRE. Either Services or ServicesExtra is required.
	ServicesExtra []*ServicesInfo `json:"ServicesExtra,omitnil,omitempty" name:"ServicesExtra"`
}

type CreateServiceTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Template name of the protocol port
	ServiceTemplateName *string `json:"ServiceTemplateName,omitnil,omitempty" name:"ServiceTemplateName"`

	// Supported ports inlcude single port, multiple ports, consecutive ports and all ports. Supported protocols include TCP, UDP, ICMP and GRE. Either Services or ServicesExtra is required.
	Services []*string `json:"Services,omitnil,omitempty" name:"Services"`

	// You can add remarks. Supported ports include single port, multiple ports, consecutive ports and all ports. Supported protocols include TCP, UDP, ICMP and GRE. Either Services or ServicesExtra is required.
	ServicesExtra []*ServicesInfo `json:"ServicesExtra,omitnil,omitempty" name:"ServicesExtra"`
}

func (r *CreateServiceTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateServiceTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceTemplateName")
	delete(f, "Services")
	delete(f, "ServicesExtra")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateServiceTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateServiceTemplateResponseParams struct {
	// Protocol port template object.
	ServiceTemplate *ServiceTemplate `json:"ServiceTemplate,omitnil,omitempty" name:"ServiceTemplate"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateServiceTemplateResponse struct {
	*tchttp.BaseResponse
	Response *CreateServiceTemplateResponseParams `json:"Response"`
}

func (r *CreateServiceTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateServiceTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSnapshotPoliciesRequestParams struct {
	// Details of a snapshot policy
	SnapshotPolicies []*SnapshotPolicy `json:"SnapshotPolicies,omitnil,omitempty" name:"SnapshotPolicies"`
}

type CreateSnapshotPoliciesRequest struct {
	*tchttp.BaseRequest
	
	// Details of a snapshot policy
	SnapshotPolicies []*SnapshotPolicy `json:"SnapshotPolicies,omitnil,omitempty" name:"SnapshotPolicies"`
}

func (r *CreateSnapshotPoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSnapshotPoliciesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SnapshotPolicies")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSnapshotPoliciesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSnapshotPoliciesResponseParams struct {
	// Snapshot policies
	SnapshotPolicies []*SnapshotPolicy `json:"SnapshotPolicies,omitnil,omitempty" name:"SnapshotPolicies"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateSnapshotPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *CreateSnapshotPoliciesResponseParams `json:"Response"`
}

func (r *CreateSnapshotPoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSnapshotPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSubnetRequestParams struct {
	// The ID of the VPC instance to be operated on. You can obtain the parameter value from the VpcId field in the returned result of DescribeVpcs API.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// The subnet name. The maximum length is 60 bytes.
	SubnetName *string `json:"SubnetName,omitnil,omitempty" name:"SubnetName"`

	// The subnet IP address range. It must be within the VPC IP address range. Subnet IP address ranges cannot overlap with each other within the same VPC.
	CidrBlock *string `json:"CidrBlock,omitnil,omitempty" name:"CidrBlock"`

	// The ID of the availability zone in which the subnet resides. You can set up disaster recovery across availability zones by choosing different availability zones for different subnets.
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Bound tags, such as [{"Key": "city", "Value": "shanghai"}].
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// CDC instance ID
	CdcId *string `json:"CdcId,omitnil,omitempty" name:"CdcId"`
}

type CreateSubnetRequest struct {
	*tchttp.BaseRequest
	
	// The ID of the VPC instance to be operated on. You can obtain the parameter value from the VpcId field in the returned result of DescribeVpcs API.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// The subnet name. The maximum length is 60 bytes.
	SubnetName *string `json:"SubnetName,omitnil,omitempty" name:"SubnetName"`

	// The subnet IP address range. It must be within the VPC IP address range. Subnet IP address ranges cannot overlap with each other within the same VPC.
	CidrBlock *string `json:"CidrBlock,omitnil,omitempty" name:"CidrBlock"`

	// The ID of the availability zone in which the subnet resides. You can set up disaster recovery across availability zones by choosing different availability zones for different subnets.
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Bound tags, such as [{"Key": "city", "Value": "shanghai"}].
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// CDC instance ID
	CdcId *string `json:"CdcId,omitnil,omitempty" name:"CdcId"`
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
	delete(f, "Tags")
	delete(f, "CdcId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSubnetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSubnetResponseParams struct {
	// Subnet object.
	Subnet *Subnet `json:"Subnet,omitnil,omitempty" name:"Subnet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type CreateSubnetsRequestParams struct {
	// The `ID` of the `VPC` instance, such as `vpc-6v2ht8q5`.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// The subnet object list.
	Subnets []*SubnetInput `json:"Subnets,omitnil,omitempty" name:"Subnets"`

	// Bound tags. Note that the collection of tags here is shared by all subnet objects in the list. You cannot specify tags for each subnet. Example: [{"Key": "city", "Value": "shanghai"}].
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// ID of the CDC instance to which the subnets will be created
	CdcId *string `json:"CdcId,omitnil,omitempty" name:"CdcId"`
}

type CreateSubnetsRequest struct {
	*tchttp.BaseRequest
	
	// The `ID` of the `VPC` instance, such as `vpc-6v2ht8q5`.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// The subnet object list.
	Subnets []*SubnetInput `json:"Subnets,omitnil,omitempty" name:"Subnets"`

	// Bound tags. Note that the collection of tags here is shared by all subnet objects in the list. You cannot specify tags for each subnet. Example: [{"Key": "city", "Value": "shanghai"}].
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// ID of the CDC instance to which the subnets will be created
	CdcId *string `json:"CdcId,omitnil,omitempty" name:"CdcId"`
}

func (r *CreateSubnetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSubnetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcId")
	delete(f, "Subnets")
	delete(f, "Tags")
	delete(f, "CdcId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSubnetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSubnetsResponseParams struct {
	// The list of newly created subnets.
	SubnetSet []*Subnet `json:"SubnetSet,omitnil,omitempty" name:"SubnetSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateSubnetsResponse struct {
	*tchttp.BaseResponse
	Response *CreateSubnetsResponseParams `json:"Response"`
}

func (r *CreateSubnetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSubnetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVpcEndPointRequestParams struct {
	// VPC instance ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Subnet instance ID
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Endpoint name
	EndPointName *string `json:"EndPointName,omitnil,omitempty" name:"EndPointName"`

	// Endpoint service ID
	EndPointServiceId *string `json:"EndPointServiceId,omitnil,omitempty" name:"EndPointServiceId"`

	// Endpoint VIP. You can apply for a specified IP.
	EndPointVip *string `json:"EndPointVip,omitnil,omitempty" name:"EndPointVip"`

	// Security group ID
	SecurityGroupId *string `json:"SecurityGroupId,omitnil,omitempty" name:"SecurityGroupId"`
}

type CreateVpcEndPointRequest struct {
	*tchttp.BaseRequest
	
	// VPC instance ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Subnet instance ID
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Endpoint name
	EndPointName *string `json:"EndPointName,omitnil,omitempty" name:"EndPointName"`

	// Endpoint service ID
	EndPointServiceId *string `json:"EndPointServiceId,omitnil,omitempty" name:"EndPointServiceId"`

	// Endpoint VIP. You can apply for a specified IP.
	EndPointVip *string `json:"EndPointVip,omitnil,omitempty" name:"EndPointVip"`

	// Security group ID
	SecurityGroupId *string `json:"SecurityGroupId,omitnil,omitempty" name:"SecurityGroupId"`
}

func (r *CreateVpcEndPointRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVpcEndPointRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "EndPointName")
	delete(f, "EndPointServiceId")
	delete(f, "EndPointVip")
	delete(f, "SecurityGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateVpcEndPointRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVpcEndPointResponseParams struct {
	// Endpoint details
	EndPoint *EndPoint `json:"EndPoint,omitnil,omitempty" name:"EndPoint"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateVpcEndPointResponse struct {
	*tchttp.BaseResponse
	Response *CreateVpcEndPointResponseParams `json:"Response"`
}

func (r *CreateVpcEndPointResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVpcEndPointResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVpcEndPointServiceRequestParams struct {
	// VPC instance ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Endpoint service name
	EndPointServiceName *string `json:"EndPointServiceName,omitnil,omitempty" name:"EndPointServiceName"`

	// Whether to automatically accept
	AutoAcceptFlag *bool `json:"AutoAcceptFlag,omitnil,omitempty" name:"AutoAcceptFlag"`

	// Real server ID, such as `lb-xxx`.
	ServiceInstanceId *string `json:"ServiceInstanceId,omitnil,omitempty" name:"ServiceInstanceId"`

	// (Disused) Whether it’s a PaaS service
	IsPassService *bool `json:"IsPassService,omitnil,omitempty" name:"IsPassService"`

	// Mounted PaaS service type. Values: `CLB` (default), `CDB`, `CRS`
	ServiceType *string `json:"ServiceType,omitnil,omitempty" name:"ServiceType"`
}

type CreateVpcEndPointServiceRequest struct {
	*tchttp.BaseRequest
	
	// VPC instance ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Endpoint service name
	EndPointServiceName *string `json:"EndPointServiceName,omitnil,omitempty" name:"EndPointServiceName"`

	// Whether to automatically accept
	AutoAcceptFlag *bool `json:"AutoAcceptFlag,omitnil,omitempty" name:"AutoAcceptFlag"`

	// Real server ID, such as `lb-xxx`.
	ServiceInstanceId *string `json:"ServiceInstanceId,omitnil,omitempty" name:"ServiceInstanceId"`

	// (Disused) Whether it’s a PaaS service
	IsPassService *bool `json:"IsPassService,omitnil,omitempty" name:"IsPassService"`

	// Mounted PaaS service type. Values: `CLB` (default), `CDB`, `CRS`
	ServiceType *string `json:"ServiceType,omitnil,omitempty" name:"ServiceType"`
}

func (r *CreateVpcEndPointServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVpcEndPointServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcId")
	delete(f, "EndPointServiceName")
	delete(f, "AutoAcceptFlag")
	delete(f, "ServiceInstanceId")
	delete(f, "IsPassService")
	delete(f, "ServiceType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateVpcEndPointServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVpcEndPointServiceResponseParams struct {
	// Endpoint service details
	EndPointService *EndPointService `json:"EndPointService,omitnil,omitempty" name:"EndPointService"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateVpcEndPointServiceResponse struct {
	*tchttp.BaseResponse
	Response *CreateVpcEndPointServiceResponseParams `json:"Response"`
}

func (r *CreateVpcEndPointServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVpcEndPointServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVpcEndPointServiceWhiteListRequestParams struct {
	// UIN
	UserUin *string `json:"UserUin,omitnil,omitempty" name:"UserUin"`

	// Endpoint service ID
	EndPointServiceId *string `json:"EndPointServiceId,omitnil,omitempty" name:"EndPointServiceId"`

	// Allowlist description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type CreateVpcEndPointServiceWhiteListRequest struct {
	*tchttp.BaseRequest
	
	// UIN
	UserUin *string `json:"UserUin,omitnil,omitempty" name:"UserUin"`

	// Endpoint service ID
	EndPointServiceId *string `json:"EndPointServiceId,omitnil,omitempty" name:"EndPointServiceId"`

	// Allowlist description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *CreateVpcEndPointServiceWhiteListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVpcEndPointServiceWhiteListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserUin")
	delete(f, "EndPointServiceId")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateVpcEndPointServiceWhiteListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVpcEndPointServiceWhiteListResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateVpcEndPointServiceWhiteListResponse struct {
	*tchttp.BaseResponse
	Response *CreateVpcEndPointServiceWhiteListResponseParams `json:"Response"`
}

func (r *CreateVpcEndPointServiceWhiteListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVpcEndPointServiceWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVpcRequestParams struct {
	// The VPC name. The maximum length is 60 bytes.
	VpcName *string `json:"VpcName,omitnil,omitempty" name:"VpcName"`

	// VPC CIDR block, which must fall within the following three private network IP ranges: 10.0.0.0/12, 172.16.0.0/12, and 192.168.0.0/16.
	CidrBlock *string `json:"CidrBlock,omitnil,omitempty" name:"CidrBlock"`

	// Whether multicast is enabled. `true`: Enabled. `false`: Not enabled.
	EnableMulticast *string `json:"EnableMulticast,omitnil,omitempty" name:"EnableMulticast"`

	// DNS address. A maximum of 4 addresses is supported.
	DnsServers []*string `json:"DnsServers,omitnil,omitempty" name:"DnsServers"`

	// Domain name of DHCP
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// Bound tags, such as [{"Key": "city", "Value": "shanghai"}]
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type CreateVpcRequest struct {
	*tchttp.BaseRequest
	
	// The VPC name. The maximum length is 60 bytes.
	VpcName *string `json:"VpcName,omitnil,omitempty" name:"VpcName"`

	// VPC CIDR block, which must fall within the following three private network IP ranges: 10.0.0.0/12, 172.16.0.0/12, and 192.168.0.0/16.
	CidrBlock *string `json:"CidrBlock,omitnil,omitempty" name:"CidrBlock"`

	// Whether multicast is enabled. `true`: Enabled. `false`: Not enabled.
	EnableMulticast *string `json:"EnableMulticast,omitnil,omitempty" name:"EnableMulticast"`

	// DNS address. A maximum of 4 addresses is supported.
	DnsServers []*string `json:"DnsServers,omitnil,omitempty" name:"DnsServers"`

	// Domain name of DHCP
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// Bound tags, such as [{"Key": "city", "Value": "shanghai"}]
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
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
	delete(f, "EnableMulticast")
	delete(f, "DnsServers")
	delete(f, "DomainName")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateVpcRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVpcResponseParams struct {
	// The VPC object.
	Vpc *Vpc `json:"Vpc,omitnil,omitempty" name:"Vpc"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

// Predefined struct for user
type CreateVpnConnectionRequestParams struct {
	// The ID of the VPN gateway instance.
	VpnGatewayId *string `json:"VpnGatewayId,omitnil,omitempty" name:"VpnGatewayId"`

	// The ID of the customer gateway, such as `cgw-2wqq41m9`. You can query the customer gateway by using the [DescribeCustomerGateways](https://intl.cloud.tencent.com/document/product/215/17516?from_cn_redirect=1) API.
	CustomerGatewayId *string `json:"CustomerGatewayId,omitnil,omitempty" name:"CustomerGatewayId"`

	// Gateway can be named freely, but the maximum length is 60 characters.
	VpnConnectionName *string `json:"VpnConnectionName,omitnil,omitempty" name:"VpnConnectionName"`

	// The pre-shared key.
	PreShareKey *string `json:"PreShareKey,omitnil,omitempty" name:"PreShareKey"`

	// VPC instance ID, which can be obtained from the `VpcId` field in the response of the [`DescribeVpcs`](https://intl.cloud.tencent.com/document/product/215/15778?from_cn_redirect=1) API.
	// This parameter is optional for a CCN-based VPN tunnel.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// The SPD policy group, for example: {"10.0.0.5/24":["172.123.10.5/16"]}. 10.0.0.5/24 is the VPC internal IP range, and 172.123.10.5/16 is the IDC IP range. The user specifies the IP range in the VPC that can communicate with the IP range in the IDC.
	SecurityPolicyDatabases []*SecurityPolicyDatabase `json:"SecurityPolicyDatabases,omitnil,omitempty" name:"SecurityPolicyDatabases"`

	// Internet Key Exchange (IKE) configuration. IKE has a self-protection mechanism. The network security protocol is configured by the user.
	IKEOptionsSpecification *IKEOptionsSpecification `json:"IKEOptionsSpecification,omitnil,omitempty" name:"IKEOptionsSpecification"`

	// IPSec configuration. The IPSec secure session configuration is provided by Tencent Cloud.
	IPSECOptionsSpecification *IPSECOptionsSpecification `json:"IPSECOptionsSpecification,omitnil,omitempty" name:"IPSECOptionsSpecification"`

	// Bound tags, such as [{"Key": "city", "Value": "shanghai"}].
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Whether the tunnel health check is supported. The default value is `False`.
	EnableHealthCheck *bool `json:"EnableHealthCheck,omitnil,omitempty" name:"EnableHealthCheck"`

	// Local IP of health check. It defaults to a random IP within 169.254.128.0/17.
	HealthCheckLocalIp *string `json:"HealthCheckLocalIp,omitnil,omitempty" name:"HealthCheckLocalIp"`

	// Peer IP of health check. It defaults to a random IP within 169.254.128.0/17.
	HealthCheckRemoteIp *string `json:"HealthCheckRemoteIp,omitnil,omitempty" name:"HealthCheckRemoteIp"`

	// Tunnel type. Valid values: `STATIC`, `StaticRoute`, and `Policy`.
	RouteType *string `json:"RouteType,omitnil,omitempty" name:"RouteType"`

	// Negotiation type. Valid values: `active` (default value), `passive` and `flowTrigger`.
	NegotiationType *string `json:"NegotiationType,omitnil,omitempty" name:"NegotiationType"`

	// Specifies whether to enable DPD. Valid values: `0` (disable) and `1` (enable)
	DpdEnable *int64 `json:"DpdEnable,omitnil,omitempty" name:"DpdEnable"`

	// DPD timeout period. Default: 30; unit: second. If the request is not responded within this period, the peer end is considered not exists. This parameter is valid when the value of `DpdEnable` is 1. 
	DpdTimeout *string `json:"DpdTimeout,omitnil,omitempty" name:"DpdTimeout"`

	// The action after DPD timeout. Valid values: `clear` (disconnect) and `restart` (try again). It’s valid when `DpdEnable` is `1`. 
	DpdAction *string `json:"DpdAction,omitnil,omitempty" name:"DpdAction"`
}

type CreateVpnConnectionRequest struct {
	*tchttp.BaseRequest
	
	// The ID of the VPN gateway instance.
	VpnGatewayId *string `json:"VpnGatewayId,omitnil,omitempty" name:"VpnGatewayId"`

	// The ID of the customer gateway, such as `cgw-2wqq41m9`. You can query the customer gateway by using the [DescribeCustomerGateways](https://intl.cloud.tencent.com/document/product/215/17516?from_cn_redirect=1) API.
	CustomerGatewayId *string `json:"CustomerGatewayId,omitnil,omitempty" name:"CustomerGatewayId"`

	// Gateway can be named freely, but the maximum length is 60 characters.
	VpnConnectionName *string `json:"VpnConnectionName,omitnil,omitempty" name:"VpnConnectionName"`

	// The pre-shared key.
	PreShareKey *string `json:"PreShareKey,omitnil,omitempty" name:"PreShareKey"`

	// VPC instance ID, which can be obtained from the `VpcId` field in the response of the [`DescribeVpcs`](https://intl.cloud.tencent.com/document/product/215/15778?from_cn_redirect=1) API.
	// This parameter is optional for a CCN-based VPN tunnel.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// The SPD policy group, for example: {"10.0.0.5/24":["172.123.10.5/16"]}. 10.0.0.5/24 is the VPC internal IP range, and 172.123.10.5/16 is the IDC IP range. The user specifies the IP range in the VPC that can communicate with the IP range in the IDC.
	SecurityPolicyDatabases []*SecurityPolicyDatabase `json:"SecurityPolicyDatabases,omitnil,omitempty" name:"SecurityPolicyDatabases"`

	// Internet Key Exchange (IKE) configuration. IKE has a self-protection mechanism. The network security protocol is configured by the user.
	IKEOptionsSpecification *IKEOptionsSpecification `json:"IKEOptionsSpecification,omitnil,omitempty" name:"IKEOptionsSpecification"`

	// IPSec configuration. The IPSec secure session configuration is provided by Tencent Cloud.
	IPSECOptionsSpecification *IPSECOptionsSpecification `json:"IPSECOptionsSpecification,omitnil,omitempty" name:"IPSECOptionsSpecification"`

	// Bound tags, such as [{"Key": "city", "Value": "shanghai"}].
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Whether the tunnel health check is supported. The default value is `False`.
	EnableHealthCheck *bool `json:"EnableHealthCheck,omitnil,omitempty" name:"EnableHealthCheck"`

	// Local IP of health check. It defaults to a random IP within 169.254.128.0/17.
	HealthCheckLocalIp *string `json:"HealthCheckLocalIp,omitnil,omitempty" name:"HealthCheckLocalIp"`

	// Peer IP of health check. It defaults to a random IP within 169.254.128.0/17.
	HealthCheckRemoteIp *string `json:"HealthCheckRemoteIp,omitnil,omitempty" name:"HealthCheckRemoteIp"`

	// Tunnel type. Valid values: `STATIC`, `StaticRoute`, and `Policy`.
	RouteType *string `json:"RouteType,omitnil,omitempty" name:"RouteType"`

	// Negotiation type. Valid values: `active` (default value), `passive` and `flowTrigger`.
	NegotiationType *string `json:"NegotiationType,omitnil,omitempty" name:"NegotiationType"`

	// Specifies whether to enable DPD. Valid values: `0` (disable) and `1` (enable)
	DpdEnable *int64 `json:"DpdEnable,omitnil,omitempty" name:"DpdEnable"`

	// DPD timeout period. Default: 30; unit: second. If the request is not responded within this period, the peer end is considered not exists. This parameter is valid when the value of `DpdEnable` is 1. 
	DpdTimeout *string `json:"DpdTimeout,omitnil,omitempty" name:"DpdTimeout"`

	// The action after DPD timeout. Valid values: `clear` (disconnect) and `restart` (try again). It’s valid when `DpdEnable` is `1`. 
	DpdAction *string `json:"DpdAction,omitnil,omitempty" name:"DpdAction"`
}

func (r *CreateVpnConnectionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVpnConnectionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpnGatewayId")
	delete(f, "CustomerGatewayId")
	delete(f, "VpnConnectionName")
	delete(f, "PreShareKey")
	delete(f, "VpcId")
	delete(f, "SecurityPolicyDatabases")
	delete(f, "IKEOptionsSpecification")
	delete(f, "IPSECOptionsSpecification")
	delete(f, "Tags")
	delete(f, "EnableHealthCheck")
	delete(f, "HealthCheckLocalIp")
	delete(f, "HealthCheckRemoteIp")
	delete(f, "RouteType")
	delete(f, "NegotiationType")
	delete(f, "DpdEnable")
	delete(f, "DpdTimeout")
	delete(f, "DpdAction")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateVpnConnectionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVpnConnectionResponseParams struct {
	// Tunnel instance object.
	VpnConnection *VpnConnection `json:"VpnConnection,omitnil,omitempty" name:"VpnConnection"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateVpnConnectionResponse struct {
	*tchttp.BaseResponse
	Response *CreateVpnConnectionResponseParams `json:"Response"`
}

func (r *CreateVpnConnectionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVpnConnectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVpnGatewayRequestParams struct {
	// VPC instance ID, which can be obtained from the `VpcId` field in the response of the [`DescribeVpcs`](https://intl.cloud.tencent.com/document/product/215/15778?from_cn_redirect=1) API.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// The VPN gateway name. The maximum length is 60 bytes.
	VpnGatewayName *string `json:"VpnGatewayName,omitnil,omitempty" name:"VpnGatewayName"`

	// The public network bandwidth configuration. Available bandwidth specifications: 5, 10, 20, 50, and 100. Unit: Mbps
	InternetMaxBandwidthOut *uint64 `json:"InternetMaxBandwidthOut,omitnil,omitempty" name:"InternetMaxBandwidthOut"`

	// The VPN gateway billing mode. PREPAID: prepaid means monthly subscription. POSTPAID_BY_HOUR: postpaid means pay-as-you-go. Default: POSTPAID_BY_HOUR. If prepaid mode is specified, the `InstanceChargePrepaid` parameter must be entered.
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// Parameter settings for prepaid billing mode, also known as monthly subscription. This parameter can specify the purchase period and other attributes such as auto-renewal. This parameter is mandatory for prepaid instances.
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitnil,omitempty" name:"InstanceChargePrepaid"`

	// The availability zone, such as `ap-guangzhou-2`.
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// VPN gateway type. Values: `CCN` (CCN VPN gateway), `SSL` (SSL VPN gateway)
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Bound tags, such as [{"Key": "city", "Value": "shanghai"}].
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// CDC instance ID
	CdcId *string `json:"CdcId,omitnil,omitempty" name:"CdcId"`

	// Maximum number of connected clients allowed for the SSL VPN gateway. Valid values: [5, 10, 20, 50, 100]. This parameter is only required for SSL VPN gateways.
	MaxConnection *uint64 `json:"MaxConnection,omitnil,omitempty" name:"MaxConnection"`
}

type CreateVpnGatewayRequest struct {
	*tchttp.BaseRequest
	
	// VPC instance ID, which can be obtained from the `VpcId` field in the response of the [`DescribeVpcs`](https://intl.cloud.tencent.com/document/product/215/15778?from_cn_redirect=1) API.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// The VPN gateway name. The maximum length is 60 bytes.
	VpnGatewayName *string `json:"VpnGatewayName,omitnil,omitempty" name:"VpnGatewayName"`

	// The public network bandwidth configuration. Available bandwidth specifications: 5, 10, 20, 50, and 100. Unit: Mbps
	InternetMaxBandwidthOut *uint64 `json:"InternetMaxBandwidthOut,omitnil,omitempty" name:"InternetMaxBandwidthOut"`

	// The VPN gateway billing mode. PREPAID: prepaid means monthly subscription. POSTPAID_BY_HOUR: postpaid means pay-as-you-go. Default: POSTPAID_BY_HOUR. If prepaid mode is specified, the `InstanceChargePrepaid` parameter must be entered.
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// Parameter settings for prepaid billing mode, also known as monthly subscription. This parameter can specify the purchase period and other attributes such as auto-renewal. This parameter is mandatory for prepaid instances.
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitnil,omitempty" name:"InstanceChargePrepaid"`

	// The availability zone, such as `ap-guangzhou-2`.
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// VPN gateway type. Values: `CCN` (CCN VPN gateway), `SSL` (SSL VPN gateway)
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Bound tags, such as [{"Key": "city", "Value": "shanghai"}].
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// CDC instance ID
	CdcId *string `json:"CdcId,omitnil,omitempty" name:"CdcId"`

	// Maximum number of connected clients allowed for the SSL VPN gateway. Valid values: [5, 10, 20, 50, 100]. This parameter is only required for SSL VPN gateways.
	MaxConnection *uint64 `json:"MaxConnection,omitnil,omitempty" name:"MaxConnection"`
}

func (r *CreateVpnGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVpnGatewayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcId")
	delete(f, "VpnGatewayName")
	delete(f, "InternetMaxBandwidthOut")
	delete(f, "InstanceChargeType")
	delete(f, "InstanceChargePrepaid")
	delete(f, "Zone")
	delete(f, "Type")
	delete(f, "Tags")
	delete(f, "CdcId")
	delete(f, "MaxConnection")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateVpnGatewayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVpnGatewayResponseParams struct {
	// VPN gateway object.
	VpnGateway *VpnGateway `json:"VpnGateway,omitnil,omitempty" name:"VpnGateway"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateVpnGatewayResponse struct {
	*tchttp.BaseResponse
	Response *CreateVpnGatewayResponseParams `json:"Response"`
}

func (r *CreateVpnGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVpnGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVpnGatewayRoutesRequestParams struct {
	// VPN gateway ID
	VpnGatewayId *string `json:"VpnGatewayId,omitnil,omitempty" name:"VpnGatewayId"`

	// Destination route list of a VPN gateway
	Routes []*VpnGatewayRoute `json:"Routes,omitnil,omitempty" name:"Routes"`
}

type CreateVpnGatewayRoutesRequest struct {
	*tchttp.BaseRequest
	
	// VPN gateway ID
	VpnGatewayId *string `json:"VpnGatewayId,omitnil,omitempty" name:"VpnGatewayId"`

	// Destination route list of a VPN gateway
	Routes []*VpnGatewayRoute `json:"Routes,omitnil,omitempty" name:"Routes"`
}

func (r *CreateVpnGatewayRoutesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVpnGatewayRoutesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpnGatewayId")
	delete(f, "Routes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateVpnGatewayRoutesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVpnGatewayRoutesResponseParams struct {
	// Destination routes of a VPN gateway
	Routes []*VpnGatewayRoute `json:"Routes,omitnil,omitempty" name:"Routes"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateVpnGatewayRoutesResponse struct {
	*tchttp.BaseResponse
	Response *CreateVpnGatewayRoutesResponseParams `json:"Response"`
}

func (r *CreateVpnGatewayRoutesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVpnGatewayRoutesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CrossBorderCompliance struct {
	// Service provider. Valid values: `UNICOM`.
	ServiceProvider *string `json:"ServiceProvider,omitnil,omitempty" name:"ServiceProvider"`

	// ID of compliance review request.
	ComplianceId *uint64 `json:"ComplianceId,omitnil,omitempty" name:"ComplianceId"`

	// Full company name.
	Company *string `json:"Company,omitnil,omitempty" name:"Company"`

	// Unified Social Credit Code.
	UniformSocialCreditCode *string `json:"UniformSocialCreditCode,omitnil,omitempty" name:"UniformSocialCreditCode"`

	// Legal person.
	LegalPerson *string `json:"LegalPerson,omitnil,omitempty" name:"LegalPerson"`

	// Issuing authority.
	IssuingAuthority *string `json:"IssuingAuthority,omitnil,omitempty" name:"IssuingAuthority"`

	// Business license.
	BusinessLicense *string `json:"BusinessLicense,omitnil,omitempty" name:"BusinessLicense"`

	// Business address.
	BusinessAddress *string `json:"BusinessAddress,omitnil,omitempty" name:"BusinessAddress"`

	// Zip code.
	PostCode *uint64 `json:"PostCode,omitnil,omitempty" name:"PostCode"`

	// Operator.
	Manager *string `json:"Manager,omitnil,omitempty" name:"Manager"`

	// Operator ID card number.
	ManagerId *string `json:"ManagerId,omitnil,omitempty" name:"ManagerId"`

	// Operator ID card.
	ManagerIdCard *string `json:"ManagerIdCard,omitnil,omitempty" name:"ManagerIdCard"`

	// Operator address.
	ManagerAddress *string `json:"ManagerAddress,omitnil,omitempty" name:"ManagerAddress"`

	// Operator phone number.
	ManagerTelephone *string `json:"ManagerTelephone,omitnil,omitempty" name:"ManagerTelephone"`

	// Email.
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// Service handling form.
	ServiceHandlingForm *string `json:"ServiceHandlingForm,omitnil,omitempty" name:"ServiceHandlingForm"`

	// Authorization letter.
	AuthorizationLetter *string `json:"AuthorizationLetter,omitnil,omitempty" name:"AuthorizationLetter"`

	// Information security commitment.
	SafetyCommitment *string `json:"SafetyCommitment,omitnil,omitempty" name:"SafetyCommitment"`

	// Service start date.
	ServiceStartDate *string `json:"ServiceStartDate,omitnil,omitempty" name:"ServiceStartDate"`

	// Service end date.
	ServiceEndDate *string `json:"ServiceEndDate,omitnil,omitempty" name:"ServiceEndDate"`

	// Status. Valid values: `PENDING`, `APPROVED`, and `DENY`.
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// Creation time of the review form.
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`
}

type CustomerGateway struct {
	// The unique ID of the customer gateway
	CustomerGatewayId *string `json:"CustomerGatewayId,omitnil,omitempty" name:"CustomerGatewayId"`

	// Gateway Name
	CustomerGatewayName *string `json:"CustomerGatewayName,omitnil,omitempty" name:"CustomerGatewayName"`

	// Public network address
	IpAddress *string `json:"IpAddress,omitnil,omitempty" name:"IpAddress"`

	// The creation time.
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`
}

type CustomerGatewayVendor struct {
	// Platform.
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// Software version.
	SoftwareVersion *string `json:"SoftwareVersion,omitnil,omitempty" name:"SoftwareVersion"`

	// Vendor name.
	VendorName *string `json:"VendorName,omitnil,omitempty" name:"VendorName"`
}

type CvmInstance struct {
	// VPC instance ID.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Subnet instance ID.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// CVM instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// CVM Name
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// CVM status.
	InstanceState *string `json:"InstanceState,omitnil,omitempty" name:"InstanceState"`

	// Number of CPU cores in an instance (in core).
	CPU *uint64 `json:"CPU,omitnil,omitempty" name:"CPU"`

	// Instance’s memory capacity. Unit: GB.
	Memory *uint64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// The creation time.
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// Instance type.
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// Instance ENI quota (including primary ENIs).
	EniLimit *uint64 `json:"EniLimit,omitnil,omitempty" name:"EniLimit"`

	// Private IP quoata for instance ENIs (including primary ENIs).
	EniIpLimit *uint64 `json:"EniIpLimit,omitnil,omitempty" name:"EniIpLimit"`

	// The number of ENIs (including primary ENIs) bound to a instance.
	InstanceEniCount *uint64 `json:"InstanceEniCount,omitnil,omitempty" name:"InstanceEniCount"`
}

type DefaultVpcSubnet struct {
	// Default VPC ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Default subnet ID
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Default VPC name
	VpcName *string `json:"VpcName,omitnil,omitempty" name:"VpcName"`

	// Default subnet name
	SubnetName *string `json:"SubnetName,omitnil,omitempty" name:"SubnetName"`

	// Default subnet IP range
	CidrBlock *string `json:"CidrBlock,omitnil,omitempty" name:"CidrBlock"`
}

// Predefined struct for user
type DeleteAddressTemplateGroupRequestParams struct {
	// The IP address template group instance ID, such as `ipmg-90cex8mq`.
	AddressTemplateGroupId *string `json:"AddressTemplateGroupId,omitnil,omitempty" name:"AddressTemplateGroupId"`
}

type DeleteAddressTemplateGroupRequest struct {
	*tchttp.BaseRequest
	
	// The IP address template group instance ID, such as `ipmg-90cex8mq`.
	AddressTemplateGroupId *string `json:"AddressTemplateGroupId,omitnil,omitempty" name:"AddressTemplateGroupId"`
}

func (r *DeleteAddressTemplateGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAddressTemplateGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AddressTemplateGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAddressTemplateGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAddressTemplateGroupResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteAddressTemplateGroupResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAddressTemplateGroupResponseParams `json:"Response"`
}

func (r *DeleteAddressTemplateGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAddressTemplateGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAddressTemplateRequestParams struct {
	// The IP address template instance ID, such as `ipm-09o5m8kc`.
	AddressTemplateId *string `json:"AddressTemplateId,omitnil,omitempty" name:"AddressTemplateId"`
}

type DeleteAddressTemplateRequest struct {
	*tchttp.BaseRequest
	
	// The IP address template instance ID, such as `ipm-09o5m8kc`.
	AddressTemplateId *string `json:"AddressTemplateId,omitnil,omitempty" name:"AddressTemplateId"`
}

func (r *DeleteAddressTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAddressTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AddressTemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAddressTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAddressTemplateResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteAddressTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAddressTemplateResponseParams `json:"Response"`
}

func (r *DeleteAddressTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAddressTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAssistantCidrRequestParams struct {
	// `VPC` instance `ID`, such as `vpc-6v2ht8q5`.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Array of CIDR blocks, such as ["10.0.0.0/16", "172.16.0.0/16"]
	CidrBlocks []*string `json:"CidrBlocks,omitnil,omitempty" name:"CidrBlocks"`
}

type DeleteAssistantCidrRequest struct {
	*tchttp.BaseRequest
	
	// `VPC` instance `ID`, such as `vpc-6v2ht8q5`.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Array of CIDR blocks, such as ["10.0.0.0/16", "172.16.0.0/16"]
	CidrBlocks []*string `json:"CidrBlocks,omitnil,omitempty" name:"CidrBlocks"`
}

func (r *DeleteAssistantCidrRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAssistantCidrRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcId")
	delete(f, "CidrBlocks")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAssistantCidrRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAssistantCidrResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteAssistantCidrResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAssistantCidrResponseParams `json:"Response"`
}

func (r *DeleteAssistantCidrResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAssistantCidrResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteBandwidthPackageRequestParams struct {
	// The unique ID of the bandwidth package to be deleted.
	BandwidthPackageId *string `json:"BandwidthPackageId,omitnil,omitempty" name:"BandwidthPackageId"`
}

type DeleteBandwidthPackageRequest struct {
	*tchttp.BaseRequest
	
	// The unique ID of the bandwidth package to be deleted.
	BandwidthPackageId *string `json:"BandwidthPackageId,omitnil,omitempty" name:"BandwidthPackageId"`
}

func (r *DeleteBandwidthPackageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBandwidthPackageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BandwidthPackageId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteBandwidthPackageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteBandwidthPackageResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteBandwidthPackageResponse struct {
	*tchttp.BaseResponse
	Response *DeleteBandwidthPackageResponseParams `json:"Response"`
}

func (r *DeleteBandwidthPackageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBandwidthPackageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCcnRequestParams struct {
	// The CCN instance ID, such as `ccn-f49l6u0z`.
	CcnId *string `json:"CcnId,omitnil,omitempty" name:"CcnId"`
}

type DeleteCcnRequest struct {
	*tchttp.BaseRequest
	
	// The CCN instance ID, such as `ccn-f49l6u0z`.
	CcnId *string `json:"CcnId,omitnil,omitempty" name:"CcnId"`
}

func (r *DeleteCcnRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCcnRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CcnId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCcnRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCcnResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteCcnResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCcnResponseParams `json:"Response"`
}

func (r *DeleteCcnResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCcnResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCustomerGatewayRequestParams struct {
	// The ID of the customer gateway, such as `cgw-2wqq41m9`. You can query the customer gateway by using the [DescribeCustomerGateways](https://intl.cloud.tencent.com/document/api/215/17516?from_cn_redirect=1) API.
	CustomerGatewayId *string `json:"CustomerGatewayId,omitnil,omitempty" name:"CustomerGatewayId"`
}

type DeleteCustomerGatewayRequest struct {
	*tchttp.BaseRequest
	
	// The ID of the customer gateway, such as `cgw-2wqq41m9`. You can query the customer gateway by using the [DescribeCustomerGateways](https://intl.cloud.tencent.com/document/api/215/17516?from_cn_redirect=1) API.
	CustomerGatewayId *string `json:"CustomerGatewayId,omitnil,omitempty" name:"CustomerGatewayId"`
}

func (r *DeleteCustomerGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCustomerGatewayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CustomerGatewayId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCustomerGatewayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCustomerGatewayResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteCustomerGatewayResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCustomerGatewayResponseParams `json:"Response"`
}

func (r *DeleteCustomerGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCustomerGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDirectConnectGatewayCcnRoutesRequestParams struct {
	// The ID of the Direct Connect gateway, such as `dcg-prpqlmg1`
	DirectConnectGatewayId *string `json:"DirectConnectGatewayId,omitnil,omitempty" name:"DirectConnectGatewayId"`

	// The route ID, such as `ccnr-f49l6u0z`.
	RouteIds []*string `json:"RouteIds,omitnil,omitempty" name:"RouteIds"`
}

type DeleteDirectConnectGatewayCcnRoutesRequest struct {
	*tchttp.BaseRequest
	
	// The ID of the Direct Connect gateway, such as `dcg-prpqlmg1`
	DirectConnectGatewayId *string `json:"DirectConnectGatewayId,omitnil,omitempty" name:"DirectConnectGatewayId"`

	// The route ID, such as `ccnr-f49l6u0z`.
	RouteIds []*string `json:"RouteIds,omitnil,omitempty" name:"RouteIds"`
}

func (r *DeleteDirectConnectGatewayCcnRoutesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDirectConnectGatewayCcnRoutesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DirectConnectGatewayId")
	delete(f, "RouteIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDirectConnectGatewayCcnRoutesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDirectConnectGatewayCcnRoutesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteDirectConnectGatewayCcnRoutesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDirectConnectGatewayCcnRoutesResponseParams `json:"Response"`
}

func (r *DeleteDirectConnectGatewayCcnRoutesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDirectConnectGatewayCcnRoutesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDirectConnectGatewayRequestParams struct {
	// The unique `ID` of the direct connect gateway, such as `dcg-9o233uri`.
	DirectConnectGatewayId *string `json:"DirectConnectGatewayId,omitnil,omitempty" name:"DirectConnectGatewayId"`
}

type DeleteDirectConnectGatewayRequest struct {
	*tchttp.BaseRequest
	
	// The unique `ID` of the direct connect gateway, such as `dcg-9o233uri`.
	DirectConnectGatewayId *string `json:"DirectConnectGatewayId,omitnil,omitempty" name:"DirectConnectGatewayId"`
}

func (r *DeleteDirectConnectGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDirectConnectGatewayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DirectConnectGatewayId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDirectConnectGatewayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDirectConnectGatewayResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteDirectConnectGatewayResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDirectConnectGatewayResponseParams `json:"Response"`
}

func (r *DeleteDirectConnectGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDirectConnectGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteFlowLogRequestParams struct {
	// The unique ID of the flow log.
	FlowLogId *string `json:"FlowLogId,omitnil,omitempty" name:"FlowLogId"`

	// The VPC ID or unique ID of the resource. We recommend using the unique ID. This parameter is required unless a CCN flow log is deleted.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`
}

type DeleteFlowLogRequest struct {
	*tchttp.BaseRequest
	
	// The unique ID of the flow log.
	FlowLogId *string `json:"FlowLogId,omitnil,omitempty" name:"FlowLogId"`

	// The VPC ID or unique ID of the resource. We recommend using the unique ID. This parameter is required unless a CCN flow log is deleted.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`
}

func (r *DeleteFlowLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteFlowLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FlowLogId")
	delete(f, "VpcId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteFlowLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteFlowLogResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteFlowLogResponse struct {
	*tchttp.BaseResponse
	Response *DeleteFlowLogResponseParams `json:"Response"`
}

func (r *DeleteFlowLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteFlowLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteHaVipRequestParams struct {
	// The unique `ID` of the `HAVIP`, such as `havip-9o233uri`.
	HaVipId *string `json:"HaVipId,omitnil,omitempty" name:"HaVipId"`
}

type DeleteHaVipRequest struct {
	*tchttp.BaseRequest
	
	// The unique `ID` of the `HAVIP`, such as `havip-9o233uri`.
	HaVipId *string `json:"HaVipId,omitnil,omitempty" name:"HaVipId"`
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
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DeleteLocalGatewayRequestParams struct {
	// Local gateway instance ID
	LocalGatewayId *string `json:"LocalGatewayId,omitnil,omitempty" name:"LocalGatewayId"`

	// CDC instance ID
	CdcId *string `json:"CdcId,omitnil,omitempty" name:"CdcId"`

	// VPC instance ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`
}

type DeleteLocalGatewayRequest struct {
	*tchttp.BaseRequest
	
	// Local gateway instance ID
	LocalGatewayId *string `json:"LocalGatewayId,omitnil,omitempty" name:"LocalGatewayId"`

	// CDC instance ID
	CdcId *string `json:"CdcId,omitnil,omitempty" name:"CdcId"`

	// VPC instance ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`
}

func (r *DeleteLocalGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLocalGatewayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LocalGatewayId")
	delete(f, "CdcId")
	delete(f, "VpcId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteLocalGatewayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLocalGatewayResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteLocalGatewayResponse struct {
	*tchttp.BaseResponse
	Response *DeleteLocalGatewayResponseParams `json:"Response"`
}

func (r *DeleteLocalGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLocalGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNatGatewayDestinationIpPortTranslationNatRuleRequestParams struct {
	// The ID of the NAT gateway, such as `nat-df45454`.
	NatGatewayId *string `json:"NatGatewayId,omitnil,omitempty" name:"NatGatewayId"`

	// The port forwarding rules of the NAT gateway.
	DestinationIpPortTranslationNatRules []*DestinationIpPortTranslationNatRule `json:"DestinationIpPortTranslationNatRules,omitnil,omitempty" name:"DestinationIpPortTranslationNatRules"`
}

type DeleteNatGatewayDestinationIpPortTranslationNatRuleRequest struct {
	*tchttp.BaseRequest
	
	// The ID of the NAT gateway, such as `nat-df45454`.
	NatGatewayId *string `json:"NatGatewayId,omitnil,omitempty" name:"NatGatewayId"`

	// The port forwarding rules of the NAT gateway.
	DestinationIpPortTranslationNatRules []*DestinationIpPortTranslationNatRule `json:"DestinationIpPortTranslationNatRules,omitnil,omitempty" name:"DestinationIpPortTranslationNatRules"`
}

func (r *DeleteNatGatewayDestinationIpPortTranslationNatRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNatGatewayDestinationIpPortTranslationNatRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NatGatewayId")
	delete(f, "DestinationIpPortTranslationNatRules")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteNatGatewayDestinationIpPortTranslationNatRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNatGatewayDestinationIpPortTranslationNatRuleResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteNatGatewayDestinationIpPortTranslationNatRuleResponse struct {
	*tchttp.BaseResponse
	Response *DeleteNatGatewayDestinationIpPortTranslationNatRuleResponseParams `json:"Response"`
}

func (r *DeleteNatGatewayDestinationIpPortTranslationNatRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNatGatewayDestinationIpPortTranslationNatRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNatGatewayRequestParams struct {
	// The ID of the NAT gateway, such as `nat-df45454`.
	NatGatewayId *string `json:"NatGatewayId,omitnil,omitempty" name:"NatGatewayId"`
}

type DeleteNatGatewayRequest struct {
	*tchttp.BaseRequest
	
	// The ID of the NAT gateway, such as `nat-df45454`.
	NatGatewayId *string `json:"NatGatewayId,omitnil,omitempty" name:"NatGatewayId"`
}

func (r *DeleteNatGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNatGatewayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NatGatewayId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteNatGatewayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNatGatewayResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteNatGatewayResponse struct {
	*tchttp.BaseResponse
	Response *DeleteNatGatewayResponseParams `json:"Response"`
}

func (r *DeleteNatGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNatGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNatGatewaySourceIpTranslationNatRuleRequestParams struct {
	// The ID of the NAT Gateway, such as `nat-df45454`
	NatGatewayId *string `json:"NatGatewayId,omitnil,omitempty" name:"NatGatewayId"`

	// The list of SNAT rule IDs of a NAT Gateway, such as `snat-df43254`
	NatGatewaySnatIds []*string `json:"NatGatewaySnatIds,omitnil,omitempty" name:"NatGatewaySnatIds"`
}

type DeleteNatGatewaySourceIpTranslationNatRuleRequest struct {
	*tchttp.BaseRequest
	
	// The ID of the NAT Gateway, such as `nat-df45454`
	NatGatewayId *string `json:"NatGatewayId,omitnil,omitempty" name:"NatGatewayId"`

	// The list of SNAT rule IDs of a NAT Gateway, such as `snat-df43254`
	NatGatewaySnatIds []*string `json:"NatGatewaySnatIds,omitnil,omitempty" name:"NatGatewaySnatIds"`
}

func (r *DeleteNatGatewaySourceIpTranslationNatRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNatGatewaySourceIpTranslationNatRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NatGatewayId")
	delete(f, "NatGatewaySnatIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteNatGatewaySourceIpTranslationNatRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNatGatewaySourceIpTranslationNatRuleResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteNatGatewaySourceIpTranslationNatRuleResponse struct {
	*tchttp.BaseResponse
	Response *DeleteNatGatewaySourceIpTranslationNatRuleResponseParams `json:"Response"`
}

func (r *DeleteNatGatewaySourceIpTranslationNatRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNatGatewaySourceIpTranslationNatRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNetDetectRequestParams struct {
	// ID of a network probe, such as `netd-12345678`.
	NetDetectId *string `json:"NetDetectId,omitnil,omitempty" name:"NetDetectId"`
}

type DeleteNetDetectRequest struct {
	*tchttp.BaseRequest
	
	// ID of a network probe, such as `netd-12345678`.
	NetDetectId *string `json:"NetDetectId,omitnil,omitempty" name:"NetDetectId"`
}

func (r *DeleteNetDetectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNetDetectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NetDetectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteNetDetectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNetDetectResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteNetDetectResponse struct {
	*tchttp.BaseResponse
	Response *DeleteNetDetectResponseParams `json:"Response"`
}

func (r *DeleteNetDetectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNetDetectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNetworkAclQuintupleEntriesRequestParams struct {
	// Network ACL instance ID, such as `acl-12345678`.
	NetworkAclId *string `json:"NetworkAclId,omitnil,omitempty" name:"NetworkAclId"`

	// Network ACL quintuple rule set.
	NetworkAclQuintupleSet *NetworkAclQuintupleEntries `json:"NetworkAclQuintupleSet,omitnil,omitempty" name:"NetworkAclQuintupleSet"`
}

type DeleteNetworkAclQuintupleEntriesRequest struct {
	*tchttp.BaseRequest
	
	// Network ACL instance ID, such as `acl-12345678`.
	NetworkAclId *string `json:"NetworkAclId,omitnil,omitempty" name:"NetworkAclId"`

	// Network ACL quintuple rule set.
	NetworkAclQuintupleSet *NetworkAclQuintupleEntries `json:"NetworkAclQuintupleSet,omitnil,omitempty" name:"NetworkAclQuintupleSet"`
}

func (r *DeleteNetworkAclQuintupleEntriesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNetworkAclQuintupleEntriesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NetworkAclId")
	delete(f, "NetworkAclQuintupleSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteNetworkAclQuintupleEntriesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNetworkAclQuintupleEntriesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteNetworkAclQuintupleEntriesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteNetworkAclQuintupleEntriesResponseParams `json:"Response"`
}

func (r *DeleteNetworkAclQuintupleEntriesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNetworkAclQuintupleEntriesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNetworkAclRequestParams struct {
	// Network ACL instance ID, such as `acl-12345678`.
	NetworkAclId *string `json:"NetworkAclId,omitnil,omitempty" name:"NetworkAclId"`
}

type DeleteNetworkAclRequest struct {
	*tchttp.BaseRequest
	
	// Network ACL instance ID, such as `acl-12345678`.
	NetworkAclId *string `json:"NetworkAclId,omitnil,omitempty" name:"NetworkAclId"`
}

func (r *DeleteNetworkAclRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNetworkAclRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NetworkAclId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteNetworkAclRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNetworkAclResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteNetworkAclResponse struct {
	*tchttp.BaseResponse
	Response *DeleteNetworkAclResponseParams `json:"Response"`
}

func (r *DeleteNetworkAclResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNetworkAclResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNetworkInterfaceRequestParams struct {
	// The ID of the ENI instance, such as `eni-m6dyj72l`.
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitnil,omitempty" name:"NetworkInterfaceId"`
}

type DeleteNetworkInterfaceRequest struct {
	*tchttp.BaseRequest
	
	// The ID of the ENI instance, such as `eni-m6dyj72l`.
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitnil,omitempty" name:"NetworkInterfaceId"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteNetworkInterfaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNetworkInterfaceResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// The route table instance ID, such as `rtb-azd4dt1c`.
	RouteTableId *string `json:"RouteTableId,omitnil,omitempty" name:"RouteTableId"`
}

type DeleteRouteTableRequest struct {
	*tchttp.BaseRequest
	
	// The route table instance ID, such as `rtb-azd4dt1c`.
	RouteTableId *string `json:"RouteTableId,omitnil,omitempty" name:"RouteTableId"`
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
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// Route table instance ID.
	RouteTableId *string `json:"RouteTableId,omitnil,omitempty" name:"RouteTableId"`

	// Routing policy object. Only the `RouteId` field is required when deleting a routing policy.
	Routes []*Route `json:"Routes,omitnil,omitempty" name:"Routes"`
}

type DeleteRoutesRequest struct {
	*tchttp.BaseRequest
	
	// Route table instance ID.
	RouteTableId *string `json:"RouteTableId,omitnil,omitempty" name:"RouteTableId"`

	// Routing policy object. Only the `RouteId` field is required when deleting a routing policy.
	Routes []*Route `json:"Routes,omitnil,omitempty" name:"Routes"`
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
	// Details of the routing policy that has been deleted.
	RouteSet []*Route `json:"RouteSet,omitnil,omitempty" name:"RouteSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// The security group instance ID, such as `sg-33ocnj9n`. This can be obtained through DescribeSecurityGroups.
	SecurityGroupId *string `json:"SecurityGroupId,omitnil,omitempty" name:"SecurityGroupId"`

	// The policy set of the security group. One request can only delete one or more policies in one direction. Both PolicyIndex-matching deletion and security group policy-matching deletion methods are supported. Each request can use only one matching method.
	SecurityGroupPolicySet *SecurityGroupPolicySet `json:"SecurityGroupPolicySet,omitnil,omitempty" name:"SecurityGroupPolicySet"`
}

type DeleteSecurityGroupPoliciesRequest struct {
	*tchttp.BaseRequest
	
	// The security group instance ID, such as `sg-33ocnj9n`. This can be obtained through DescribeSecurityGroups.
	SecurityGroupId *string `json:"SecurityGroupId,omitnil,omitempty" name:"SecurityGroupId"`

	// The policy set of the security group. One request can only delete one or more policies in one direction. Both PolicyIndex-matching deletion and security group policy-matching deletion methods are supported. Each request can use only one matching method.
	SecurityGroupPolicySet *SecurityGroupPolicySet `json:"SecurityGroupPolicySet,omitnil,omitempty" name:"SecurityGroupPolicySet"`
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
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// The security group instance ID, such as `sg-33ocnj9n`. This can be obtained through DescribeSecurityGroups.
	SecurityGroupId *string `json:"SecurityGroupId,omitnil,omitempty" name:"SecurityGroupId"`
}

type DeleteSecurityGroupRequest struct {
	*tchttp.BaseRequest
	
	// The security group instance ID, such as `sg-33ocnj9n`. This can be obtained through DescribeSecurityGroups.
	SecurityGroupId *string `json:"SecurityGroupId,omitnil,omitempty" name:"SecurityGroupId"`
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
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DeleteServiceTemplateGroupRequestParams struct {
	// The protocol port template group instance ID, such as `ppmg-n17uxvve`.
	ServiceTemplateGroupId *string `json:"ServiceTemplateGroupId,omitnil,omitempty" name:"ServiceTemplateGroupId"`
}

type DeleteServiceTemplateGroupRequest struct {
	*tchttp.BaseRequest
	
	// The protocol port template group instance ID, such as `ppmg-n17uxvve`.
	ServiceTemplateGroupId *string `json:"ServiceTemplateGroupId,omitnil,omitempty" name:"ServiceTemplateGroupId"`
}

func (r *DeleteServiceTemplateGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteServiceTemplateGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceTemplateGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteServiceTemplateGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteServiceTemplateGroupResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteServiceTemplateGroupResponse struct {
	*tchttp.BaseResponse
	Response *DeleteServiceTemplateGroupResponseParams `json:"Response"`
}

func (r *DeleteServiceTemplateGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteServiceTemplateGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteServiceTemplateRequestParams struct {
	// Protocol port template instance ID, such as `ppm-e6dy460g`.
	ServiceTemplateId *string `json:"ServiceTemplateId,omitnil,omitempty" name:"ServiceTemplateId"`
}

type DeleteServiceTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Protocol port template instance ID, such as `ppm-e6dy460g`.
	ServiceTemplateId *string `json:"ServiceTemplateId,omitnil,omitempty" name:"ServiceTemplateId"`
}

func (r *DeleteServiceTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteServiceTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceTemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteServiceTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteServiceTemplateResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteServiceTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DeleteServiceTemplateResponseParams `json:"Response"`
}

func (r *DeleteServiceTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteServiceTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSnapshotPoliciesRequestParams struct {
	// Snapshot policy IDs
	SnapshotPolicyIds []*string `json:"SnapshotPolicyIds,omitnil,omitempty" name:"SnapshotPolicyIds"`
}

type DeleteSnapshotPoliciesRequest struct {
	*tchttp.BaseRequest
	
	// Snapshot policy IDs
	SnapshotPolicyIds []*string `json:"SnapshotPolicyIds,omitnil,omitempty" name:"SnapshotPolicyIds"`
}

func (r *DeleteSnapshotPoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSnapshotPoliciesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SnapshotPolicyIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSnapshotPoliciesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSnapshotPoliciesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteSnapshotPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSnapshotPoliciesResponseParams `json:"Response"`
}

func (r *DeleteSnapshotPoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSnapshotPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSubnetRequestParams struct {
	// The ID of the subnet instance. You can obtain the parameter value from the SubnetId field in the returned result of DescribeSubnets API.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`
}

type DeleteSubnetRequest struct {
	*tchttp.BaseRequest
	
	// The ID of the subnet instance. You can obtain the parameter value from the SubnetId field in the returned result of DescribeSubnets API.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSubnetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSubnetResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DeleteTrafficPackagesRequestParams struct {
	// The unique ID array of traffic packages to delete
	TrafficPackageIds []*string `json:"TrafficPackageIds,omitnil,omitempty" name:"TrafficPackageIds"`
}

type DeleteTrafficPackagesRequest struct {
	*tchttp.BaseRequest
	
	// The unique ID array of traffic packages to delete
	TrafficPackageIds []*string `json:"TrafficPackageIds,omitnil,omitempty" name:"TrafficPackageIds"`
}

func (r *DeleteTrafficPackagesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTrafficPackagesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TrafficPackageIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTrafficPackagesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTrafficPackagesResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteTrafficPackagesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteTrafficPackagesResponseParams `json:"Response"`
}

func (r *DeleteTrafficPackagesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTrafficPackagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteVpcEndPointRequestParams struct {
	// Endpoint ID
	EndPointId *string `json:"EndPointId,omitnil,omitempty" name:"EndPointId"`
}

type DeleteVpcEndPointRequest struct {
	*tchttp.BaseRequest
	
	// Endpoint ID
	EndPointId *string `json:"EndPointId,omitnil,omitempty" name:"EndPointId"`
}

func (r *DeleteVpcEndPointRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteVpcEndPointRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EndPointId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteVpcEndPointRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteVpcEndPointResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteVpcEndPointResponse struct {
	*tchttp.BaseResponse
	Response *DeleteVpcEndPointResponseParams `json:"Response"`
}

func (r *DeleteVpcEndPointResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteVpcEndPointResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteVpcEndPointServiceRequestParams struct {
	// Endpoint ID
	EndPointServiceId *string `json:"EndPointServiceId,omitnil,omitempty" name:"EndPointServiceId"`
}

type DeleteVpcEndPointServiceRequest struct {
	*tchttp.BaseRequest
	
	// Endpoint ID
	EndPointServiceId *string `json:"EndPointServiceId,omitnil,omitempty" name:"EndPointServiceId"`
}

func (r *DeleteVpcEndPointServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteVpcEndPointServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EndPointServiceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteVpcEndPointServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteVpcEndPointServiceResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteVpcEndPointServiceResponse struct {
	*tchttp.BaseResponse
	Response *DeleteVpcEndPointServiceResponseParams `json:"Response"`
}

func (r *DeleteVpcEndPointServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteVpcEndPointServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteVpcEndPointServiceWhiteListRequestParams struct {
	// Array of user UINs
	UserUin []*string `json:"UserUin,omitnil,omitempty" name:"UserUin"`

	// Endpoint service ID
	EndPointServiceId *string `json:"EndPointServiceId,omitnil,omitempty" name:"EndPointServiceId"`
}

type DeleteVpcEndPointServiceWhiteListRequest struct {
	*tchttp.BaseRequest
	
	// Array of user UINs
	UserUin []*string `json:"UserUin,omitnil,omitempty" name:"UserUin"`

	// Endpoint service ID
	EndPointServiceId *string `json:"EndPointServiceId,omitnil,omitempty" name:"EndPointServiceId"`
}

func (r *DeleteVpcEndPointServiceWhiteListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteVpcEndPointServiceWhiteListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserUin")
	delete(f, "EndPointServiceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteVpcEndPointServiceWhiteListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteVpcEndPointServiceWhiteListResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteVpcEndPointServiceWhiteListResponse struct {
	*tchttp.BaseResponse
	Response *DeleteVpcEndPointServiceWhiteListResponseParams `json:"Response"`
}

func (r *DeleteVpcEndPointServiceWhiteListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteVpcEndPointServiceWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteVpcRequestParams struct {
	// The ID of the VPC instance. You can obtain the parameter value from the VpcId field in the returned result of DescribeVpcs API.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`
}

type DeleteVpcRequest struct {
	*tchttp.BaseRequest
	
	// The ID of the VPC instance. You can obtain the parameter value from the VpcId field in the returned result of DescribeVpcs API.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteVpcRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteVpcResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DeleteVpnConnectionRequestParams struct {
	// The ID of the VPN gateway instance.
	VpnGatewayId *string `json:"VpnGatewayId,omitnil,omitempty" name:"VpnGatewayId"`

	// The ID of the VPN tunnel instance, such as `vpnx-f49l6u0z`.
	VpnConnectionId *string `json:"VpnConnectionId,omitnil,omitempty" name:"VpnConnectionId"`
}

type DeleteVpnConnectionRequest struct {
	*tchttp.BaseRequest
	
	// The ID of the VPN gateway instance.
	VpnGatewayId *string `json:"VpnGatewayId,omitnil,omitempty" name:"VpnGatewayId"`

	// The ID of the VPN tunnel instance, such as `vpnx-f49l6u0z`.
	VpnConnectionId *string `json:"VpnConnectionId,omitnil,omitempty" name:"VpnConnectionId"`
}

func (r *DeleteVpnConnectionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteVpnConnectionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpnGatewayId")
	delete(f, "VpnConnectionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteVpnConnectionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteVpnConnectionResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteVpnConnectionResponse struct {
	*tchttp.BaseResponse
	Response *DeleteVpnConnectionResponseParams `json:"Response"`
}

func (r *DeleteVpnConnectionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteVpnConnectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteVpnGatewayRequestParams struct {
	// The ID of the VPN gateway instance.
	VpnGatewayId *string `json:"VpnGatewayId,omitnil,omitempty" name:"VpnGatewayId"`
}

type DeleteVpnGatewayRequest struct {
	*tchttp.BaseRequest
	
	// The ID of the VPN gateway instance.
	VpnGatewayId *string `json:"VpnGatewayId,omitnil,omitempty" name:"VpnGatewayId"`
}

func (r *DeleteVpnGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteVpnGatewayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpnGatewayId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteVpnGatewayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteVpnGatewayResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteVpnGatewayResponse struct {
	*tchttp.BaseResponse
	Response *DeleteVpnGatewayResponseParams `json:"Response"`
}

func (r *DeleteVpnGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteVpnGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteVpnGatewayRoutesRequestParams struct {
	// Instance ID of the VPN gateway
	VpnGatewayId *string `json:"VpnGatewayId,omitnil,omitempty" name:"VpnGatewayId"`

	// List of route IDs
	RouteIds []*string `json:"RouteIds,omitnil,omitempty" name:"RouteIds"`
}

type DeleteVpnGatewayRoutesRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID of the VPN gateway
	VpnGatewayId *string `json:"VpnGatewayId,omitnil,omitempty" name:"VpnGatewayId"`

	// List of route IDs
	RouteIds []*string `json:"RouteIds,omitnil,omitempty" name:"RouteIds"`
}

func (r *DeleteVpnGatewayRoutesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteVpnGatewayRoutesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpnGatewayId")
	delete(f, "RouteIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteVpnGatewayRoutesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteVpnGatewayRoutesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteVpnGatewayRoutesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteVpnGatewayRoutesResponseParams `json:"Response"`
}

func (r *DeleteVpnGatewayRoutesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteVpnGatewayRoutesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccountAttributesRequestParams struct {

}

type DescribeAccountAttributesRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeAccountAttributesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccountAttributesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccountAttributesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccountAttributesResponseParams struct {
	// User account attribute object
	AccountAttributeSet []*AccountAttribute `json:"AccountAttributeSet,omitnil,omitempty" name:"AccountAttributeSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAccountAttributesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAccountAttributesResponseParams `json:"Response"`
}

func (r *DescribeAccountAttributesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccountAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAddressQuotaRequestParams struct {

}

type DescribeAddressQuotaRequest struct {
	*tchttp.BaseRequest
	
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
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAddressQuotaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAddressQuotaResponseParams struct {
	// The quota information of EIPs in an account.
	QuotaSet []*Quota `json:"QuotaSet,omitnil,omitempty" name:"QuotaSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeAddressTemplateGroupsRequestParams struct {
	// Filter conditions.
	// <li>address-template-group-name - String - (Filter condition) IP address template group name.</li>
	// <li>address-template-group-id - String - (Filter condition) IP address template group instance ID, such as `ipmg-mdunqeb6`.</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Offset. The default value is 0.
	Offset *string `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of values to be returned. The default value is 20. Maximum is 100.
	Limit *string `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeAddressTemplateGroupsRequest struct {
	*tchttp.BaseRequest
	
	// Filter conditions.
	// <li>address-template-group-name - String - (Filter condition) IP address template group name.</li>
	// <li>address-template-group-id - String - (Filter condition) IP address template group instance ID, such as `ipmg-mdunqeb6`.</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Offset. The default value is 0.
	Offset *string `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of values to be returned. The default value is 20. Maximum is 100.
	Limit *string `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeAddressTemplateGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAddressTemplateGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAddressTemplateGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAddressTemplateGroupsResponseParams struct {
	// The number of instances meeting the filter condition.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// IP address template.
	AddressTemplateGroupSet []*AddressTemplateGroup `json:"AddressTemplateGroupSet,omitnil,omitempty" name:"AddressTemplateGroupSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAddressTemplateGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAddressTemplateGroupsResponseParams `json:"Response"`
}

func (r *DescribeAddressTemplateGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAddressTemplateGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAddressTemplatesRequestParams struct {
	// Filters
	// <li>address-template-name - IP address template name.</li>
	// <li>address-template-id - IP address template ID, such as `ipm-mdunqeb6`.</li>
	// <li>address-ip - IP address.</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Offset. The default value is 0.
	Offset *string `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of values to be returned. The default value is 20. Maximum is 100.
	Limit *string `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeAddressTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// Filters
	// <li>address-template-name - IP address template name.</li>
	// <li>address-template-id - IP address template ID, such as `ipm-mdunqeb6`.</li>
	// <li>address-ip - IP address.</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Offset. The default value is 0.
	Offset *string `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of values to be returned. The default value is 20. Maximum is 100.
	Limit *string `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeAddressTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAddressTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAddressTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAddressTemplatesResponseParams struct {
	// The number of instances meeting the filter condition.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// IP address template.
	AddressTemplateSet []*AddressTemplate `json:"AddressTemplateSet,omitnil,omitempty" name:"AddressTemplateSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAddressTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAddressTemplatesResponseParams `json:"Response"`
}

func (r *DescribeAddressTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAddressTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAddressesRequestParams struct {
	// The list of unique IDs of EIPs in the format of `eip-11112222`. `AddressIds` and `Filters.address-id` cannot be specified at the same time.
	AddressIds []*string `json:"AddressIds,omitnil,omitempty" name:"AddressIds"`

	// Each request can have up to 10 `Filters` and 100 `Filter.Values`. Detailed filter conditions:
	// <li> address-id - String - Optional - Filter by unique EIP ID, such as `eip-11112222`.</li>
	// <li> address-name - String - Optional - Filter by EIP name. Fuzzy filtering is not supported.</li>
	// <li> address-ip - String - Optional - Filter by EIP address.</li>
	// <li> address-status - String - Optional - Filter by EIP status. Valid values: `CREATING`, `BINDING`, `BIND`, `UNBINDING`, `UNBIND`, `OFFLINING`, and `BIND_ENI`.</li>
	// <li> instance-id - String - Optional - Filter by the ID of the instance bound to the EIP, such as `ins-11112222`.</li>
	// <li> private-ip-address - String - Optional - Filter by the private IP address bound to the EIP.</li>
	// <li> network-interface-id - String - Optional - Filter by ID of the ENI bound to the EIP, such as `eni-11112222`.</li>
	// <li> is-arrears - String - Optional - Filter by the fact whether the EIP is overdue (TRUE: the EIP is overdue | FALSE: the billing status of the EIP is normal).</li>
	// <li> address-type - String - Optional - Filter by IP type. Valid values: `WanIP`, `EIP`, `AnycastEIP`, and `HighQualityEIP`. Default value: `EIP`.</li>
	// <li> address-isp - String - Optional - Filter by ISP type. Valid values: `BGP`, `CMCC`, `CUCC`, and `CTCC`.</li>
	// <li> dedicated-cluster-id - String - Optional - Filter by unique CDC ID, such as `cluster-11112222`.</li>
	// <li> tag-key - String - Optional - Filter by tag key.</li>
	// <li> tag-value - String - Optional - Filter by tag value.</li>
	// <li> tag:tag-key - String - Optional - Filter by tag key-value pair. Use a specific tag key to replace `tag-key`.</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// The Offset. The default value is 0. For more information about `Offset`, see the relevant section in the API documentation.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100. For more information on `Limit`, see the relevant section in the API documentation.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeAddressesRequest struct {
	*tchttp.BaseRequest
	
	// The list of unique IDs of EIPs in the format of `eip-11112222`. `AddressIds` and `Filters.address-id` cannot be specified at the same time.
	AddressIds []*string `json:"AddressIds,omitnil,omitempty" name:"AddressIds"`

	// Each request can have up to 10 `Filters` and 100 `Filter.Values`. Detailed filter conditions:
	// <li> address-id - String - Optional - Filter by unique EIP ID, such as `eip-11112222`.</li>
	// <li> address-name - String - Optional - Filter by EIP name. Fuzzy filtering is not supported.</li>
	// <li> address-ip - String - Optional - Filter by EIP address.</li>
	// <li> address-status - String - Optional - Filter by EIP status. Valid values: `CREATING`, `BINDING`, `BIND`, `UNBINDING`, `UNBIND`, `OFFLINING`, and `BIND_ENI`.</li>
	// <li> instance-id - String - Optional - Filter by the ID of the instance bound to the EIP, such as `ins-11112222`.</li>
	// <li> private-ip-address - String - Optional - Filter by the private IP address bound to the EIP.</li>
	// <li> network-interface-id - String - Optional - Filter by ID of the ENI bound to the EIP, such as `eni-11112222`.</li>
	// <li> is-arrears - String - Optional - Filter by the fact whether the EIP is overdue (TRUE: the EIP is overdue | FALSE: the billing status of the EIP is normal).</li>
	// <li> address-type - String - Optional - Filter by IP type. Valid values: `WanIP`, `EIP`, `AnycastEIP`, and `HighQualityEIP`. Default value: `EIP`.</li>
	// <li> address-isp - String - Optional - Filter by ISP type. Valid values: `BGP`, `CMCC`, `CUCC`, and `CTCC`.</li>
	// <li> dedicated-cluster-id - String - Optional - Filter by unique CDC ID, such as `cluster-11112222`.</li>
	// <li> tag-key - String - Optional - Filter by tag key.</li>
	// <li> tag-value - String - Optional - Filter by tag value.</li>
	// <li> tag:tag-key - String - Optional - Filter by tag key-value pair. Use a specific tag key to replace `tag-key`.</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// The Offset. The default value is 0. For more information about `Offset`, see the relevant section in the API documentation.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100. For more information on `Limit`, see the relevant section in the API documentation.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
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
	// Number of EIPs meeting the condition.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List of EIPs details.
	AddressSet []*Address `json:"AddressSet,omitnil,omitempty" name:"AddressSet"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeAssistantCidrRequestParams struct {
	// The ID of a VPC instance set, such as `vpc-6v2ht8q5`.
	VpcIds []*string `json:"VpcIds,omitnil,omitempty" name:"VpcIds"`

	// Filter condition. `VpcIds` and `Filters` cannot be specified at the same time.
	// <li>vpc-id - String - (Filter condition) VPC instance ID, such as `vpc-f49l6u0z`.</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeAssistantCidrRequest struct {
	*tchttp.BaseRequest
	
	// The ID of a VPC instance set, such as `vpc-6v2ht8q5`.
	VpcIds []*string `json:"VpcIds,omitnil,omitempty" name:"VpcIds"`

	// Filter condition. `VpcIds` and `Filters` cannot be specified at the same time.
	// <li>vpc-id - String - (Filter condition) VPC instance ID, such as `vpc-f49l6u0z`.</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeAssistantCidrRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssistantCidrRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAssistantCidrRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssistantCidrResponseParams struct {
	// A set of eligible secondary CIDR blocks
	// Note: This field may return null, indicating that no valid value was found.
	AssistantCidrSet []*AssistantCidr `json:"AssistantCidrSet,omitnil,omitempty" name:"AssistantCidrSet"`

	// Number of eligible instances.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAssistantCidrResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAssistantCidrResponseParams `json:"Response"`
}

func (r *DescribeAssistantCidrResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssistantCidrResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBandwidthPackageBillUsageRequestParams struct {
	// Unique ID of the pay-as-you-go bandwidth package.
	BandwidthPackageId *string `json:"BandwidthPackageId,omitnil,omitempty" name:"BandwidthPackageId"`
}

type DescribeBandwidthPackageBillUsageRequest struct {
	*tchttp.BaseRequest
	
	// Unique ID of the pay-as-you-go bandwidth package.
	BandwidthPackageId *string `json:"BandwidthPackageId,omitnil,omitempty" name:"BandwidthPackageId"`
}

func (r *DescribeBandwidthPackageBillUsageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBandwidthPackageBillUsageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BandwidthPackageId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBandwidthPackageBillUsageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBandwidthPackageBillUsageResponseParams struct {
	// Current billable usage.
	BandwidthPackageBillBandwidthSet []*BandwidthPackageBillBandwidth `json:"BandwidthPackageBillBandwidthSet,omitnil,omitempty" name:"BandwidthPackageBillBandwidthSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBandwidthPackageBillUsageResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBandwidthPackageBillUsageResponseParams `json:"Response"`
}

func (r *DescribeBandwidthPackageBillUsageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBandwidthPackageBillUsageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBandwidthPackageQuotaRequestParams struct {

}

type DescribeBandwidthPackageQuotaRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeBandwidthPackageQuotaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBandwidthPackageQuotaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBandwidthPackageQuotaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBandwidthPackageQuotaResponseParams struct {
	// The quota of the bandwidth package.
	QuotaSet []*Quota `json:"QuotaSet,omitnil,omitempty" name:"QuotaSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBandwidthPackageQuotaResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBandwidthPackageQuotaResponseParams `json:"Response"`
}

func (r *DescribeBandwidthPackageQuotaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBandwidthPackageQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBandwidthPackageResourcesRequestParams struct {
	// Unique ID of the bandwidth package in the form of `bwp-11112222`.
	BandwidthPackageId *string `json:"BandwidthPackageId,omitnil,omitempty" name:"BandwidthPackageId"`

	// Each request can have up to 10 `Filters` and 5 `Filter.Values`. `AddressIds` and `Filters` cannot be specified at the same time. The specific filter conditions are as follows:
	// <li>resource-id - String - Required: no -  (Filter condition) Filters by the unique ID of resources in a bandwidth package, such as `eip-11112222`.</li>
	// <li>resource-type - String - Required: no - (Filter condition) Filters by the type of resources in a bandwidth package. It now supports only EIP (`Address`) and load balancer (`LoadBalance`).</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// The offset. Default value: 0. For more information on `Offset`, see the relevant sections in API [Introduction](https://intl.cloud.tencent.com/document/api/213/11646?from_cn_redirect=1).
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The number of returned results. Default value: 20. Maximum value: 100. For more information on `Limit`, see the relevant sections in API [Introduction](https://intl.cloud.tencent.com/document/api/213/11646?from_cn_redirect=1).
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeBandwidthPackageResourcesRequest struct {
	*tchttp.BaseRequest
	
	// Unique ID of the bandwidth package in the form of `bwp-11112222`.
	BandwidthPackageId *string `json:"BandwidthPackageId,omitnil,omitempty" name:"BandwidthPackageId"`

	// Each request can have up to 10 `Filters` and 5 `Filter.Values`. `AddressIds` and `Filters` cannot be specified at the same time. The specific filter conditions are as follows:
	// <li>resource-id - String - Required: no -  (Filter condition) Filters by the unique ID of resources in a bandwidth package, such as `eip-11112222`.</li>
	// <li>resource-type - String - Required: no - (Filter condition) Filters by the type of resources in a bandwidth package. It now supports only EIP (`Address`) and load balancer (`LoadBalance`).</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// The offset. Default value: 0. For more information on `Offset`, see the relevant sections in API [Introduction](https://intl.cloud.tencent.com/document/api/213/11646?from_cn_redirect=1).
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The number of returned results. Default value: 20. Maximum value: 100. For more information on `Limit`, see the relevant sections in API [Introduction](https://intl.cloud.tencent.com/document/api/213/11646?from_cn_redirect=1).
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeBandwidthPackageResourcesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBandwidthPackageResourcesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BandwidthPackageId")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBandwidthPackageResourcesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBandwidthPackageResourcesResponseParams struct {
	// The number of eligible resources in the bandwidth package.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The list of resources in the bandwidth package.
	ResourceSet []*Resource `json:"ResourceSet,omitnil,omitempty" name:"ResourceSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBandwidthPackageResourcesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBandwidthPackageResourcesResponseParams `json:"Response"`
}

func (r *DescribeBandwidthPackageResourcesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBandwidthPackageResourcesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBandwidthPackagesRequestParams struct {
	// The unique ID list of bandwidth packages.
	BandwidthPackageIds []*string `json:"BandwidthPackageIds,omitnil,omitempty" name:"BandwidthPackageIds"`

	// Each request can have up to 10 `Filters`. `BandwidthPackageIds` and `Filters` cannot be specified at the same time. The specific filter conditions are as follows:
	// <li> bandwidth-package_id - String - Required: No - (Filter condition) Filter by the unique ID of the bandwidth package.</li>
	// <li> bandwidth-package-name - String - Required: No - (Filter condition) Filter by the bandwidth package name. Fuzzy filtering is not supported.</li>
	// <li> network-type - String - Required: No - (Filter condition) Filter by the bandwidth package type. Valid values: `HIGH_QUALITY_BGP`, `BGP`, `SINGLEISP`, and `ANYCAST`.</li>
	// <li> charge-type - String - Required: No - (Filter condition) Filter by the bandwidth package billing mode. Valid values: `TOP5_POSTPAID_BY_MONTH` and `PERCENT95_POSTPAID_BY_MONTH`.</li>
	// <li> resource.resource-type - String - Required: No - (Filter condition) Filter by the bandwidth package resource type. Valid values: `Address` and `LoadBalance`.</li>
	// <li> resource.resource-id - String - Required: No - (Filter condition) Filter by the bandwidth package resource ID, such as `eip-xxxx` and `lb-xxxx`.</li>
	// <li> resource.address-ip - String - Required: No - (Filter condition) Filter by the bandwidth package resource IP.</li>
	// <li> tag-key - String - Required: No - (Filter condition) Filter by tag key.</li>
	// <li> tag-value - String - Required: No - (Filter condition) Filter by tag value.</li>
	// <li> tag:tag-key - String - Required: No - (Filter condition) Filter by tag key-value pair. Use a specific tag key to replace `tag-key`.</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Offset. This value defaults to 0. For more information, see the relevant sections in API [Overview](https://intl.cloud.tencent.com/document/api/213/15688?from_cn_redirect=1).
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of bandwidth packages returned. This value defaults to 20. The maximum is 100. For more information, see the relevant sections in API [Overview](https://intl.cloud.tencent.com/document/api/213/15688?from_cn_redirect=1).
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeBandwidthPackagesRequest struct {
	*tchttp.BaseRequest
	
	// The unique ID list of bandwidth packages.
	BandwidthPackageIds []*string `json:"BandwidthPackageIds,omitnil,omitempty" name:"BandwidthPackageIds"`

	// Each request can have up to 10 `Filters`. `BandwidthPackageIds` and `Filters` cannot be specified at the same time. The specific filter conditions are as follows:
	// <li> bandwidth-package_id - String - Required: No - (Filter condition) Filter by the unique ID of the bandwidth package.</li>
	// <li> bandwidth-package-name - String - Required: No - (Filter condition) Filter by the bandwidth package name. Fuzzy filtering is not supported.</li>
	// <li> network-type - String - Required: No - (Filter condition) Filter by the bandwidth package type. Valid values: `HIGH_QUALITY_BGP`, `BGP`, `SINGLEISP`, and `ANYCAST`.</li>
	// <li> charge-type - String - Required: No - (Filter condition) Filter by the bandwidth package billing mode. Valid values: `TOP5_POSTPAID_BY_MONTH` and `PERCENT95_POSTPAID_BY_MONTH`.</li>
	// <li> resource.resource-type - String - Required: No - (Filter condition) Filter by the bandwidth package resource type. Valid values: `Address` and `LoadBalance`.</li>
	// <li> resource.resource-id - String - Required: No - (Filter condition) Filter by the bandwidth package resource ID, such as `eip-xxxx` and `lb-xxxx`.</li>
	// <li> resource.address-ip - String - Required: No - (Filter condition) Filter by the bandwidth package resource IP.</li>
	// <li> tag-key - String - Required: No - (Filter condition) Filter by tag key.</li>
	// <li> tag-value - String - Required: No - (Filter condition) Filter by tag value.</li>
	// <li> tag:tag-key - String - Required: No - (Filter condition) Filter by tag key-value pair. Use a specific tag key to replace `tag-key`.</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Offset. This value defaults to 0. For more information, see the relevant sections in API [Overview](https://intl.cloud.tencent.com/document/api/213/15688?from_cn_redirect=1).
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of bandwidth packages returned. This value defaults to 20. The maximum is 100. For more information, see the relevant sections in API [Overview](https://intl.cloud.tencent.com/document/api/213/15688?from_cn_redirect=1).
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeBandwidthPackagesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBandwidthPackagesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BandwidthPackageIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBandwidthPackagesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBandwidthPackagesResponseParams struct {
	// The number of eligible bandwidth packages.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Detail information of the bandwidth package.
	BandwidthPackageSet []*BandwidthPackage `json:"BandwidthPackageSet,omitnil,omitempty" name:"BandwidthPackageSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBandwidthPackagesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBandwidthPackagesResponseParams `json:"Response"`
}

func (r *DescribeBandwidthPackagesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBandwidthPackagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCcnAttachedInstancesRequestParams struct {
	// Offset
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The returned quantity
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Filter conditions:
	// <li>`ccn-id` - String - The CCN instance ID. </li>
	// <li>`instance-type` - String - The associated instance type. </li>
	// <li>`instance-region` - String - The associated instance region. </li>
	// <li>`instance-id` - String - The instance ID of the associated instance. </li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// The ID of the CCN instance
	CcnId *string `json:"CcnId,omitnil,omitempty" name:"CcnId"`

	// The order field supports `CcnId`, `InstanceType`, `InstanceId`, `InstanceName`, `InstanceRegion`, `AttachedTime`, and `State`.
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`

	// Order methods. Ascending: `ASC`, Descending: `DESC`.
	OrderDirection *string `json:"OrderDirection,omitnil,omitempty" name:"OrderDirection"`
}

type DescribeCcnAttachedInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Offset
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The returned quantity
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Filter conditions:
	// <li>`ccn-id` - String - The CCN instance ID. </li>
	// <li>`instance-type` - String - The associated instance type. </li>
	// <li>`instance-region` - String - The associated instance region. </li>
	// <li>`instance-id` - String - The instance ID of the associated instance. </li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// The ID of the CCN instance
	CcnId *string `json:"CcnId,omitnil,omitempty" name:"CcnId"`

	// The order field supports `CcnId`, `InstanceType`, `InstanceId`, `InstanceName`, `InstanceRegion`, `AttachedTime`, and `State`.
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`

	// Order methods. Ascending: `ASC`, Descending: `DESC`.
	OrderDirection *string `json:"OrderDirection,omitnil,omitempty" name:"OrderDirection"`
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
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	delete(f, "CcnId")
	delete(f, "OrderField")
	delete(f, "OrderDirection")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCcnAttachedInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCcnAttachedInstancesResponseParams struct {
	// The number of objects meeting the condition.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The list of associated instances.
	InstanceSet []*CcnAttachedInstance `json:"InstanceSet,omitnil,omitempty" name:"InstanceSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeCcnRegionBandwidthLimitsRequestParams struct {
	// The CCN instance ID in the format of `ccn-f49l6u0z`.
	CcnId *string `json:"CcnId,omitnil,omitempty" name:"CcnId"`
}

type DescribeCcnRegionBandwidthLimitsRequest struct {
	*tchttp.BaseRequest
	
	// The CCN instance ID in the format of `ccn-f49l6u0z`.
	CcnId *string `json:"CcnId,omitnil,omitempty" name:"CcnId"`
}

func (r *DescribeCcnRegionBandwidthLimitsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCcnRegionBandwidthLimitsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CcnId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCcnRegionBandwidthLimitsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCcnRegionBandwidthLimitsResponseParams struct {
	// The outbound bandwidth caps of all regions connected with the specified CCN instance
	CcnRegionBandwidthLimitSet []*CcnRegionBandwidthLimit `json:"CcnRegionBandwidthLimitSet,omitnil,omitempty" name:"CcnRegionBandwidthLimitSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCcnRegionBandwidthLimitsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCcnRegionBandwidthLimitsResponseParams `json:"Response"`
}

func (r *DescribeCcnRegionBandwidthLimitsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCcnRegionBandwidthLimitsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCcnRoutesRequestParams struct {
	// The CCN instance ID, such as `ccn-gree226l`.
	CcnId *string `json:"CcnId,omitnil,omitempty" name:"CcnId"`

	// The unique ID of the CCN routing policy, such as `ccnr-f49l6u0z`.
	RouteIds []*string `json:"RouteIds,omitnil,omitempty" name:"RouteIds"`

	// Filter condition. `RouteIds` and `Filters` cannot be specified at the same time.
	// <li>route-id - String - (Filter condition) Routing policy ID.</li>
	// <li>cidr-block - String - (Filter condition) Destination.</li>
	// <li>instance-type - String - (Filter condition) The next hop type.</li>
	// <li>instance-region - String - (Filter condition) The next hop region.</li>
	// <li>instance-id - String - (Filter condition) The instance ID of the next hop.</li>
	// <li>route-table-id - String - (Filter condition) The list of route table IDs in the format of `ccntr-1234edfr`. Filters by the route table ID.</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Offset
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The returned quantity
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeCcnRoutesRequest struct {
	*tchttp.BaseRequest
	
	// The CCN instance ID, such as `ccn-gree226l`.
	CcnId *string `json:"CcnId,omitnil,omitempty" name:"CcnId"`

	// The unique ID of the CCN routing policy, such as `ccnr-f49l6u0z`.
	RouteIds []*string `json:"RouteIds,omitnil,omitempty" name:"RouteIds"`

	// Filter condition. `RouteIds` and `Filters` cannot be specified at the same time.
	// <li>route-id - String - (Filter condition) Routing policy ID.</li>
	// <li>cidr-block - String - (Filter condition) Destination.</li>
	// <li>instance-type - String - (Filter condition) The next hop type.</li>
	// <li>instance-region - String - (Filter condition) The next hop region.</li>
	// <li>instance-id - String - (Filter condition) The instance ID of the next hop.</li>
	// <li>route-table-id - String - (Filter condition) The list of route table IDs in the format of `ccntr-1234edfr`. Filters by the route table ID.</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Offset
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The returned quantity
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeCcnRoutesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCcnRoutesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CcnId")
	delete(f, "RouteIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCcnRoutesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCcnRoutesResponseParams struct {
	// The number of objects meeting the condition.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The CCN routing policy object.
	RouteSet []*CcnRoute `json:"RouteSet,omitnil,omitempty" name:"RouteSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCcnRoutesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCcnRoutesResponseParams `json:"Response"`
}

func (r *DescribeCcnRoutesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCcnRoutesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCcnsRequestParams struct {
	// The CCN instance ID, such as `ccn-f49l6u0z`. Each request can have a maximum of 100 instances. `CcnIds` and `Filters` cannot be specified at the same time
	CcnIds []*string `json:"CcnIds,omitnil,omitempty" name:"CcnIds"`

	// Filter conditions. `CcnIds` and `Filters` cannot be specified at the same time.
	// <li>ccn-id - String - (Filter condition) The unique ID of the CCN, such as `vpc-f49l6u0z`.</li>
	// <li>ccn-name - String - (Filter condition) The CCN name.</li>
	// <li>ccn-description - String - (Filter condition) CCN description.</li>
	// <li>state - String - (Filter condition) The instance status. 'ISOLATED': Isolated (the account is in arrears and the service is suspended.) 'AVAILABLE': Running.</li>
	// <li>tag-key - String - Required: no - (Filter condition) Filters by tag key.</li>
	// <li>`tag:tag-key` - String - Required: no - (Filter condition) Filters by tag key pair. For this parameter, `tag-key` will be replaced with a specific tag key. For more information, see this example: **Querying the list of CCNs bound to tags**.</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Offset
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The returned quantity
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Order fields support `CcnId`, `CcnName`, `CreateTime`, `State`, and `QosLevel`
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`

	// Order methods. Ascending: `ASC`, Descending: `DESC`.
	OrderDirection *string `json:"OrderDirection,omitnil,omitempty" name:"OrderDirection"`
}

type DescribeCcnsRequest struct {
	*tchttp.BaseRequest
	
	// The CCN instance ID, such as `ccn-f49l6u0z`. Each request can have a maximum of 100 instances. `CcnIds` and `Filters` cannot be specified at the same time
	CcnIds []*string `json:"CcnIds,omitnil,omitempty" name:"CcnIds"`

	// Filter conditions. `CcnIds` and `Filters` cannot be specified at the same time.
	// <li>ccn-id - String - (Filter condition) The unique ID of the CCN, such as `vpc-f49l6u0z`.</li>
	// <li>ccn-name - String - (Filter condition) The CCN name.</li>
	// <li>ccn-description - String - (Filter condition) CCN description.</li>
	// <li>state - String - (Filter condition) The instance status. 'ISOLATED': Isolated (the account is in arrears and the service is suspended.) 'AVAILABLE': Running.</li>
	// <li>tag-key - String - Required: no - (Filter condition) Filters by tag key.</li>
	// <li>`tag:tag-key` - String - Required: no - (Filter condition) Filters by tag key pair. For this parameter, `tag-key` will be replaced with a specific tag key. For more information, see this example: **Querying the list of CCNs bound to tags**.</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Offset
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The returned quantity
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Order fields support `CcnId`, `CcnName`, `CreateTime`, `State`, and `QosLevel`
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`

	// Order methods. Ascending: `ASC`, Descending: `DESC`.
	OrderDirection *string `json:"OrderDirection,omitnil,omitempty" name:"OrderDirection"`
}

func (r *DescribeCcnsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCcnsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CcnIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "OrderField")
	delete(f, "OrderDirection")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCcnsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCcnsResponseParams struct {
	// The number of objects meeting the condition.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// CCN object.
	CcnSet []*CCN `json:"CcnSet,omitnil,omitempty" name:"CcnSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCcnsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCcnsResponseParams `json:"Response"`
}

func (r *DescribeCcnsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCcnsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClassicLinkInstancesRequestParams struct {
	// Filter conditions.
	// <li>vpc-id - String - (Filter condition) The VPC instance ID.</li>
	// <li>vm-ip - String - (Filter condition) The IP address of the CVM on the basic network.</li>
	Filters []*FilterObject `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Offset
	Offset *string `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The returned quantity
	Limit *string `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeClassicLinkInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Filter conditions.
	// <li>vpc-id - String - (Filter condition) The VPC instance ID.</li>
	// <li>vm-ip - String - (Filter condition) The IP address of the CVM on the basic network.</li>
	Filters []*FilterObject `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Offset
	Offset *string `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The returned quantity
	Limit *string `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeClassicLinkInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClassicLinkInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClassicLinkInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClassicLinkInstancesResponseParams struct {
	// The number of instances meeting the filter condition.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Classiclink instance.
	ClassicLinkInstanceSet []*ClassicLinkInstance `json:"ClassicLinkInstanceSet,omitnil,omitempty" name:"ClassicLinkInstanceSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeClassicLinkInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClassicLinkInstancesResponseParams `json:"Response"`
}

func (r *DescribeClassicLinkInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClassicLinkInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCrossBorderComplianceRequestParams struct {
	// (Exact match) Service provider. Valid values: `UNICOM`.
	ServiceProvider *string `json:"ServiceProvider,omitnil,omitempty" name:"ServiceProvider"`

	// (Exact match) ID of compliance review request.
	ComplianceId *uint64 `json:"ComplianceId,omitnil,omitempty" name:"ComplianceId"`

	// (Fuzzy match) Company name.
	Company *string `json:"Company,omitnil,omitempty" name:"Company"`

	// (Fuzzy match) Unified Social Credit Code.
	UniformSocialCreditCode *string `json:"UniformSocialCreditCode,omitnil,omitempty" name:"UniformSocialCreditCode"`

	// (Fuzzy match) Legal person.
	LegalPerson *string `json:"LegalPerson,omitnil,omitempty" name:"LegalPerson"`

	// (Fuzzy match) Issuing authority.
	IssuingAuthority *string `json:"IssuingAuthority,omitnil,omitempty" name:"IssuingAuthority"`

	// (Fuzzy match) Business address.
	BusinessAddress *string `json:"BusinessAddress,omitnil,omitempty" name:"BusinessAddress"`

	// (Exact match) Zip code.
	PostCode *uint64 `json:"PostCode,omitnil,omitempty" name:"PostCode"`

	// (Fuzzy match) Operator.
	Manager *string `json:"Manager,omitnil,omitempty" name:"Manager"`

	// (Exact match) Operator ID card number.
	ManagerId *string `json:"ManagerId,omitnil,omitempty" name:"ManagerId"`

	// (Fuzzy match) Operator address.
	ManagerAddress *string `json:"ManagerAddress,omitnil,omitempty" name:"ManagerAddress"`

	// (Exact match) Operator phone number.
	ManagerTelephone *string `json:"ManagerTelephone,omitnil,omitempty" name:"ManagerTelephone"`

	// (Exact match) Email.
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// (Exact match) Service start date, such as `2020-07-28`.
	ServiceStartDate *string `json:"ServiceStartDate,omitnil,omitempty" name:"ServiceStartDate"`

	// (Exact match) Service end date, such as `2020-07-28`.
	ServiceEndDate *string `json:"ServiceEndDate,omitnil,omitempty" name:"ServiceEndDate"`

	// (Exact match) Status. Valid values: `PENDING`, `APPROVED`, and `DENY`.
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// The offset value
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Quantity of returned items
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeCrossBorderComplianceRequest struct {
	*tchttp.BaseRequest
	
	// (Exact match) Service provider. Valid values: `UNICOM`.
	ServiceProvider *string `json:"ServiceProvider,omitnil,omitempty" name:"ServiceProvider"`

	// (Exact match) ID of compliance review request.
	ComplianceId *uint64 `json:"ComplianceId,omitnil,omitempty" name:"ComplianceId"`

	// (Fuzzy match) Company name.
	Company *string `json:"Company,omitnil,omitempty" name:"Company"`

	// (Fuzzy match) Unified Social Credit Code.
	UniformSocialCreditCode *string `json:"UniformSocialCreditCode,omitnil,omitempty" name:"UniformSocialCreditCode"`

	// (Fuzzy match) Legal person.
	LegalPerson *string `json:"LegalPerson,omitnil,omitempty" name:"LegalPerson"`

	// (Fuzzy match) Issuing authority.
	IssuingAuthority *string `json:"IssuingAuthority,omitnil,omitempty" name:"IssuingAuthority"`

	// (Fuzzy match) Business address.
	BusinessAddress *string `json:"BusinessAddress,omitnil,omitempty" name:"BusinessAddress"`

	// (Exact match) Zip code.
	PostCode *uint64 `json:"PostCode,omitnil,omitempty" name:"PostCode"`

	// (Fuzzy match) Operator.
	Manager *string `json:"Manager,omitnil,omitempty" name:"Manager"`

	// (Exact match) Operator ID card number.
	ManagerId *string `json:"ManagerId,omitnil,omitempty" name:"ManagerId"`

	// (Fuzzy match) Operator address.
	ManagerAddress *string `json:"ManagerAddress,omitnil,omitempty" name:"ManagerAddress"`

	// (Exact match) Operator phone number.
	ManagerTelephone *string `json:"ManagerTelephone,omitnil,omitempty" name:"ManagerTelephone"`

	// (Exact match) Email.
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// (Exact match) Service start date, such as `2020-07-28`.
	ServiceStartDate *string `json:"ServiceStartDate,omitnil,omitempty" name:"ServiceStartDate"`

	// (Exact match) Service end date, such as `2020-07-28`.
	ServiceEndDate *string `json:"ServiceEndDate,omitnil,omitempty" name:"ServiceEndDate"`

	// (Exact match) Status. Valid values: `PENDING`, `APPROVED`, and `DENY`.
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// The offset value
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Quantity of returned items
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeCrossBorderComplianceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCrossBorderComplianceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceProvider")
	delete(f, "ComplianceId")
	delete(f, "Company")
	delete(f, "UniformSocialCreditCode")
	delete(f, "LegalPerson")
	delete(f, "IssuingAuthority")
	delete(f, "BusinessAddress")
	delete(f, "PostCode")
	delete(f, "Manager")
	delete(f, "ManagerId")
	delete(f, "ManagerAddress")
	delete(f, "ManagerTelephone")
	delete(f, "Email")
	delete(f, "ServiceStartDate")
	delete(f, "ServiceEndDate")
	delete(f, "State")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCrossBorderComplianceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCrossBorderComplianceResponseParams struct {
	// List of compliance review requests.
	CrossBorderComplianceSet []*CrossBorderCompliance `json:"CrossBorderComplianceSet,omitnil,omitempty" name:"CrossBorderComplianceSet"`

	// Total number of compliance review requests.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCrossBorderComplianceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCrossBorderComplianceResponseParams `json:"Response"`
}

func (r *DescribeCrossBorderComplianceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCrossBorderComplianceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCustomerGatewayVendorsRequestParams struct {

}

type DescribeCustomerGatewayVendorsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeCustomerGatewayVendorsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCustomerGatewayVendorsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCustomerGatewayVendorsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCustomerGatewayVendorsResponseParams struct {
	// Customer gateway vendor information object.
	CustomerGatewayVendorSet []*CustomerGatewayVendor `json:"CustomerGatewayVendorSet,omitnil,omitempty" name:"CustomerGatewayVendorSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCustomerGatewayVendorsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCustomerGatewayVendorsResponseParams `json:"Response"`
}

func (r *DescribeCustomerGatewayVendorsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCustomerGatewayVendorsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCustomerGatewaysRequestParams struct {
	// Customer gateway ID, such as `cgw-2wqq41m9`. Each request can have a maximum of 100 instances. `CustomerGatewayIds` and `Filters` cannot be specified at the same time.
	CustomerGatewayIds []*string `json:"CustomerGatewayIds,omitnil,omitempty" name:"CustomerGatewayIds"`

	// The filter condition. For details, see the Instance Filter Conditions Table. The upper limit for `Filters` in each request is 10 and 5 for `Filter.Values`. `CustomerGatewayIds` and `Filters` cannot be specified at the same time.
	// <li>customer-gateway-id - String - (Filter condition) The unique ID of the user gateway, such as `cgw-mgp33pll`.</li>
	// <li>customer-gateway-name - String - (Filter condition) The name of the user gateway, such as `test-cgw`.</li>
	// <li>ip-address - String - (Filter condition) The public IP address, such as `58.211.1.12`.</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Offset. Default value: 0. For more information on Offset, see the relevant section in the API [Introduction](https://intl.cloud.tencent.com/document/api/213/11646?from_cn_redirect=1).
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeCustomerGatewaysRequest struct {
	*tchttp.BaseRequest
	
	// Customer gateway ID, such as `cgw-2wqq41m9`. Each request can have a maximum of 100 instances. `CustomerGatewayIds` and `Filters` cannot be specified at the same time.
	CustomerGatewayIds []*string `json:"CustomerGatewayIds,omitnil,omitempty" name:"CustomerGatewayIds"`

	// The filter condition. For details, see the Instance Filter Conditions Table. The upper limit for `Filters` in each request is 10 and 5 for `Filter.Values`. `CustomerGatewayIds` and `Filters` cannot be specified at the same time.
	// <li>customer-gateway-id - String - (Filter condition) The unique ID of the user gateway, such as `cgw-mgp33pll`.</li>
	// <li>customer-gateway-name - String - (Filter condition) The name of the user gateway, such as `test-cgw`.</li>
	// <li>ip-address - String - (Filter condition) The public IP address, such as `58.211.1.12`.</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Offset. Default value: 0. For more information on Offset, see the relevant section in the API [Introduction](https://intl.cloud.tencent.com/document/api/213/11646?from_cn_redirect=1).
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeCustomerGatewaysRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCustomerGatewaysRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CustomerGatewayIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCustomerGatewaysRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCustomerGatewaysResponseParams struct {
	// Customer gateway object list
	CustomerGatewaySet []*CustomerGateway `json:"CustomerGatewaySet,omitnil,omitempty" name:"CustomerGatewaySet"`

	// Number of eligible instances
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCustomerGatewaysResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCustomerGatewaysResponseParams `json:"Response"`
}

func (r *DescribeCustomerGatewaysResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCustomerGatewaysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDirectConnectGatewayCcnRoutesRequestParams struct {
	// The ID of the Direct Connect gateway, such as `dcg-prpqlmg1`.
	DirectConnectGatewayId *string `json:"DirectConnectGatewayId,omitnil,omitempty" name:"DirectConnectGatewayId"`

	// The route learning type of the CCN. Available values:
	// <li>`BGP` - Automatic learning.</li>
	// <li>`STATIC` - Static means user-configured. This is the default value.</li>
	CcnRouteType *string `json:"CcnRouteType,omitnil,omitempty" name:"CcnRouteType"`

	// Offset.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The returned quantity.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeDirectConnectGatewayCcnRoutesRequest struct {
	*tchttp.BaseRequest
	
	// The ID of the Direct Connect gateway, such as `dcg-prpqlmg1`.
	DirectConnectGatewayId *string `json:"DirectConnectGatewayId,omitnil,omitempty" name:"DirectConnectGatewayId"`

	// The route learning type of the CCN. Available values:
	// <li>`BGP` - Automatic learning.</li>
	// <li>`STATIC` - Static means user-configured. This is the default value.</li>
	CcnRouteType *string `json:"CcnRouteType,omitnil,omitempty" name:"CcnRouteType"`

	// Offset.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The returned quantity.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeDirectConnectGatewayCcnRoutesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDirectConnectGatewayCcnRoutesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DirectConnectGatewayId")
	delete(f, "CcnRouteType")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDirectConnectGatewayCcnRoutesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDirectConnectGatewayCcnRoutesResponseParams struct {
	// The number of objects meeting the condition.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The CCN route (IDC IP range) list.
	RouteSet []*DirectConnectGatewayCcnRoute `json:"RouteSet,omitnil,omitempty" name:"RouteSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDirectConnectGatewayCcnRoutesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDirectConnectGatewayCcnRoutesResponseParams `json:"Response"`
}

func (r *DescribeDirectConnectGatewayCcnRoutesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDirectConnectGatewayCcnRoutesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDirectConnectGatewaysRequestParams struct {
	// The unique ID of the direct connect gateway, such as `dcg-9o233uri`.
	DirectConnectGatewayIds []*string `json:"DirectConnectGatewayIds,omitnil,omitempty" name:"DirectConnectGatewayIds"`

	// Filter condition. `DirectConnectGatewayIds` and `Filters` cannot be specified at the same time.
	// <li>direct-connect-gateway-id - String - The unique ID of the direct connect gateway, such as `dcg-9o233uri`.</li>
	// <li>direct-connect-gateway-name - String - The name of the direct connect gateway. The default is fuzzy query.</li>
	// <li>direct-connect-gateway-ip - String - The IP of the direct connect gateway.</li>
	// <li>gateway-type - String - The gateway type. Valid values: `NORMAL` (Standard type), `NAT` (NAT type).</li>
	// <li>network-type- String - The network type. Valid values: `VPC` (VPC type), `CCN` (CCN type).</li>
	// <li>ccn-id - String - The ID of the CCN where the direct connect gateway resides.</li>
	// <li>vpc-id - String - The ID of the VPC where the direct connect gateway resides.</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// The offset.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Max number of results returned
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeDirectConnectGatewaysRequest struct {
	*tchttp.BaseRequest
	
	// The unique ID of the direct connect gateway, such as `dcg-9o233uri`.
	DirectConnectGatewayIds []*string `json:"DirectConnectGatewayIds,omitnil,omitempty" name:"DirectConnectGatewayIds"`

	// Filter condition. `DirectConnectGatewayIds` and `Filters` cannot be specified at the same time.
	// <li>direct-connect-gateway-id - String - The unique ID of the direct connect gateway, such as `dcg-9o233uri`.</li>
	// <li>direct-connect-gateway-name - String - The name of the direct connect gateway. The default is fuzzy query.</li>
	// <li>direct-connect-gateway-ip - String - The IP of the direct connect gateway.</li>
	// <li>gateway-type - String - The gateway type. Valid values: `NORMAL` (Standard type), `NAT` (NAT type).</li>
	// <li>network-type- String - The network type. Valid values: `VPC` (VPC type), `CCN` (CCN type).</li>
	// <li>ccn-id - String - The ID of the CCN where the direct connect gateway resides.</li>
	// <li>vpc-id - String - The ID of the VPC where the direct connect gateway resides.</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// The offset.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Max number of results returned
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeDirectConnectGatewaysRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDirectConnectGatewaysRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DirectConnectGatewayIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDirectConnectGatewaysRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDirectConnectGatewaysResponseParams struct {
	// The number of eligible objects.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The object array of the direct connect gateway.
	DirectConnectGatewaySet []*DirectConnectGateway `json:"DirectConnectGatewaySet,omitnil,omitempty" name:"DirectConnectGatewaySet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDirectConnectGatewaysResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDirectConnectGatewaysResponseParams `json:"Response"`
}

func (r *DescribeDirectConnectGatewaysResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDirectConnectGatewaysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFlowLogRequestParams struct {
	// ID of the VPC instance.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// The unique ID of the flow log.
	FlowLogId *string `json:"FlowLogId,omitnil,omitempty" name:"FlowLogId"`
}

type DescribeFlowLogRequest struct {
	*tchttp.BaseRequest
	
	// ID of the VPC instance.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// The unique ID of the flow log.
	FlowLogId *string `json:"FlowLogId,omitnil,omitempty" name:"FlowLogId"`
}

func (r *DescribeFlowLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFlowLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcId")
	delete(f, "FlowLogId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFlowLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFlowLogResponseParams struct {
	// The flow log information.
	FlowLog []*FlowLog `json:"FlowLog,omitnil,omitempty" name:"FlowLog"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeFlowLogResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFlowLogResponseParams `json:"Response"`
}

func (r *DescribeFlowLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFlowLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFlowLogsRequestParams struct {
	// ID of the VPC instance
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// The unique ID of the flow log.
	FlowLogId *string `json:"FlowLogId,omitnil,omitempty" name:"FlowLogId"`

	// The name of the flow log instance.
	FlowLogName *string `json:"FlowLogName,omitnil,omitempty" name:"FlowLogName"`

	// The resource type of the flow log. Valid values: 'VPC', 'SUBNET', and 'NETWORKINTERFACE'.
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// The unique ID of the resource.
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// Type of flow logs to be collected. Valid values: 'ACCEPT', 'REJECT' and 'ALL'.
	TrafficType *string `json:"TrafficType,omitnil,omitempty" name:"TrafficType"`

	// The storage ID of the flow log.
	CloudLogId *string `json:"CloudLogId,omitnil,omitempty" name:"CloudLogId"`

	// The storage ID status of the flow log.
	CloudLogState *string `json:"CloudLogState,omitnil,omitempty" name:"CloudLogState"`

	// Order by field. Valid values: 'flowLogName' and 'createTime'. Default value: 'createTime'.
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`

	// In ascending (`asc`) or descending (`desc`) order. Default value: 'desc'.
	OrderDirection *string `json:"OrderDirection,omitnil,omitempty" name:"OrderDirection"`

	// The offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The number of rows per page. Default value: 10.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Filter condition. `FlowLogId` and `Filters` cannot be specified at the same time.
	// <li> `tag-key` - String - Optional - Filter by the tag key.</li>
	// <li> `tag:tag-key` - String - Optional - Filter by the tag key-value pair. The tag-key should be replaced with a specified tag key.</li>
	Filters *Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// The region corresponding to the flow log storage ID.
	CloudLogRegion *string `json:"CloudLogRegion,omitnil,omitempty" name:"CloudLogRegion"`
}

type DescribeFlowLogsRequest struct {
	*tchttp.BaseRequest
	
	// ID of the VPC instance
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// The unique ID of the flow log.
	FlowLogId *string `json:"FlowLogId,omitnil,omitempty" name:"FlowLogId"`

	// The name of the flow log instance.
	FlowLogName *string `json:"FlowLogName,omitnil,omitempty" name:"FlowLogName"`

	// The resource type of the flow log. Valid values: 'VPC', 'SUBNET', and 'NETWORKINTERFACE'.
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// The unique ID of the resource.
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// Type of flow logs to be collected. Valid values: 'ACCEPT', 'REJECT' and 'ALL'.
	TrafficType *string `json:"TrafficType,omitnil,omitempty" name:"TrafficType"`

	// The storage ID of the flow log.
	CloudLogId *string `json:"CloudLogId,omitnil,omitempty" name:"CloudLogId"`

	// The storage ID status of the flow log.
	CloudLogState *string `json:"CloudLogState,omitnil,omitempty" name:"CloudLogState"`

	// Order by field. Valid values: 'flowLogName' and 'createTime'. Default value: 'createTime'.
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`

	// In ascending (`asc`) or descending (`desc`) order. Default value: 'desc'.
	OrderDirection *string `json:"OrderDirection,omitnil,omitempty" name:"OrderDirection"`

	// The offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The number of rows per page. Default value: 10.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Filter condition. `FlowLogId` and `Filters` cannot be specified at the same time.
	// <li> `tag-key` - String - Optional - Filter by the tag key.</li>
	// <li> `tag:tag-key` - String - Optional - Filter by the tag key-value pair. The tag-key should be replaced with a specified tag key.</li>
	Filters *Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// The region corresponding to the flow log storage ID.
	CloudLogRegion *string `json:"CloudLogRegion,omitnil,omitempty" name:"CloudLogRegion"`
}

func (r *DescribeFlowLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFlowLogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcId")
	delete(f, "FlowLogId")
	delete(f, "FlowLogName")
	delete(f, "ResourceType")
	delete(f, "ResourceId")
	delete(f, "TrafficType")
	delete(f, "CloudLogId")
	delete(f, "CloudLogState")
	delete(f, "OrderField")
	delete(f, "OrderDirection")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	delete(f, "CloudLogRegion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFlowLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFlowLogsResponseParams struct {
	// The instance set of flow logs.
	FlowLog []*FlowLog `json:"FlowLog,omitnil,omitempty" name:"FlowLog"`

	// The total number of flow logs.
	TotalNum *uint64 `json:"TotalNum,omitnil,omitempty" name:"TotalNum"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeFlowLogsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFlowLogsResponseParams `json:"Response"`
}

func (r *DescribeFlowLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFlowLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGatewayFlowMonitorDetailRequestParams struct {
	// The point in time. This indicates details of this minute will be queried. For example, in `2019-02-28 18:15:20`, details at `18:15` will be queried.
	TimePoint *string `json:"TimePoint,omitnil,omitempty" name:"TimePoint"`

	// The instance ID of the VPN gateway, such as `vpn-ltjahce6`.
	VpnId *string `json:"VpnId,omitnil,omitempty" name:"VpnId"`

	// The instance ID of the Direct Connect gateway, such as `dcg-ltjahce6`.
	DirectConnectGatewayId *string `json:"DirectConnectGatewayId,omitnil,omitempty" name:"DirectConnectGatewayId"`

	// The instance ID of the peering connection, such as `pcx-ltjahce6`.
	PeeringConnectionId *string `json:"PeeringConnectionId,omitnil,omitempty" name:"PeeringConnectionId"`

	// The instance ID of the NAT gateway, such as `nat-ltjahce6`.
	NatId *string `json:"NatId,omitnil,omitempty" name:"NatId"`

	// Offset.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The returned quantity.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// The order field supports `InPkg`, `OutPkg`, `InTraffic`, and `OutTraffic`.
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`

	// Order methods. Ascending: `ASC`, Descending: `DESC`.
	OrderDirection *string `json:"OrderDirection,omitnil,omitempty" name:"OrderDirection"`
}

type DescribeGatewayFlowMonitorDetailRequest struct {
	*tchttp.BaseRequest
	
	// The point in time. This indicates details of this minute will be queried. For example, in `2019-02-28 18:15:20`, details at `18:15` will be queried.
	TimePoint *string `json:"TimePoint,omitnil,omitempty" name:"TimePoint"`

	// The instance ID of the VPN gateway, such as `vpn-ltjahce6`.
	VpnId *string `json:"VpnId,omitnil,omitempty" name:"VpnId"`

	// The instance ID of the Direct Connect gateway, such as `dcg-ltjahce6`.
	DirectConnectGatewayId *string `json:"DirectConnectGatewayId,omitnil,omitempty" name:"DirectConnectGatewayId"`

	// The instance ID of the peering connection, such as `pcx-ltjahce6`.
	PeeringConnectionId *string `json:"PeeringConnectionId,omitnil,omitempty" name:"PeeringConnectionId"`

	// The instance ID of the NAT gateway, such as `nat-ltjahce6`.
	NatId *string `json:"NatId,omitnil,omitempty" name:"NatId"`

	// Offset.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The returned quantity.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// The order field supports `InPkg`, `OutPkg`, `InTraffic`, and `OutTraffic`.
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`

	// Order methods. Ascending: `ASC`, Descending: `DESC`.
	OrderDirection *string `json:"OrderDirection,omitnil,omitempty" name:"OrderDirection"`
}

func (r *DescribeGatewayFlowMonitorDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGatewayFlowMonitorDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TimePoint")
	delete(f, "VpnId")
	delete(f, "DirectConnectGatewayId")
	delete(f, "PeeringConnectionId")
	delete(f, "NatId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "OrderField")
	delete(f, "OrderDirection")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGatewayFlowMonitorDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGatewayFlowMonitorDetailResponseParams struct {
	// The number of objects meeting the condition.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The gateway traffic monitoring details.
	GatewayFlowMonitorDetailSet []*GatewayFlowMonitorDetail `json:"GatewayFlowMonitorDetailSet,omitnil,omitempty" name:"GatewayFlowMonitorDetailSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeGatewayFlowMonitorDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGatewayFlowMonitorDetailResponseParams `json:"Response"`
}

func (r *DescribeGatewayFlowMonitorDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGatewayFlowMonitorDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGatewayFlowQosRequestParams struct {
	// Gateway instance ID. Supported types:
	// Direct connect gateway instance, such as `dcg-ltjahce6`;
	// NAT gateway instance, such as `nat-ltjahce6`;
	// VPN gateway instance, such as `vpn-ltjahce6`.
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// CVM private IP addresses with limited bandwidth.
	IpAddresses []*string `json:"IpAddresses,omitnil,omitempty" name:"IpAddresses"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeGatewayFlowQosRequest struct {
	*tchttp.BaseRequest
	
	// Gateway instance ID. Supported types:
	// Direct connect gateway instance, such as `dcg-ltjahce6`;
	// NAT gateway instance, such as `nat-ltjahce6`;
	// VPN gateway instance, such as `vpn-ltjahce6`.
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// CVM private IP addresses with limited bandwidth.
	IpAddresses []*string `json:"IpAddresses,omitnil,omitempty" name:"IpAddresses"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeGatewayFlowQosRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGatewayFlowQosRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "IpAddresses")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGatewayFlowQosRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGatewayFlowQosResponseParams struct {
	// List of instance details
	GatewayQosSet []*GatewayQos `json:"GatewayQosSet,omitnil,omitempty" name:"GatewayQosSet"`

	// Number of eligible instances.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeGatewayFlowQosResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGatewayFlowQosResponseParams `json:"Response"`
}

func (r *DescribeGatewayFlowQosResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGatewayFlowQosResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHaVipsRequestParams struct {
	// The unique `ID` of the `HAVIP`, such as `havip-9o233uri`.
	HaVipIds []*string `json:"HaVipIds,omitnil,omitempty" name:"HaVipIds"`

	// Filter condition. `HaVipIds` and `Filters` cannot be specified at the same time.
	// li>havip-id - String - The unique ID of the HAVIP, such as `havip-9o233uri`.</li>
	// <li>havip-name - String - HAVIP name.</li>
	// <li>vpc-id - String - VPC ID of the HAVIP.</li>
	// <li>subnet-id - String - Subnet ID of the HAVIP.</li>
	// <li>vip - String - Virtual IP address of the HAVIP.</li>
	// <li>address-ip - String - Bound EIP.</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Offset
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The returned quantity
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeHaVipsRequest struct {
	*tchttp.BaseRequest
	
	// The unique `ID` of the `HAVIP`, such as `havip-9o233uri`.
	HaVipIds []*string `json:"HaVipIds,omitnil,omitempty" name:"HaVipIds"`

	// Filter condition. `HaVipIds` and `Filters` cannot be specified at the same time.
	// li>havip-id - String - The unique ID of the HAVIP, such as `havip-9o233uri`.</li>
	// <li>havip-name - String - HAVIP name.</li>
	// <li>vpc-id - String - VPC ID of the HAVIP.</li>
	// <li>subnet-id - String - Subnet ID of the HAVIP.</li>
	// <li>vip - String - Virtual IP address of the HAVIP.</li>
	// <li>address-ip - String - Bound EIP.</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Offset
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The returned quantity
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeHaVipsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHaVipsResponseParams struct {
	// The number of objects meeting the condition.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// `HAVIP` object array.
	HaVipSet []*HaVip `json:"HaVipSet,omitnil,omitempty" name:"HaVipSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeIpGeolocationDatabaseUrlRequestParams struct {
	// Protocol type for an IP location database. Valid value: `ipv4`.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type DescribeIpGeolocationDatabaseUrlRequest struct {
	*tchttp.BaseRequest
	
	// Protocol type for an IP location database. Valid value: `ipv4`.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

func (r *DescribeIpGeolocationDatabaseUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIpGeolocationDatabaseUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIpGeolocationDatabaseUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIpGeolocationDatabaseUrlResponseParams struct {
	// Download link of an IP location database
	DownLoadUrl *string `json:"DownLoadUrl,omitnil,omitempty" name:"DownLoadUrl"`

	// Link expiration time in UTC format following the ISO8601 standard.
	ExpiredAt *string `json:"ExpiredAt,omitnil,omitempty" name:"ExpiredAt"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeIpGeolocationDatabaseUrlResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIpGeolocationDatabaseUrlResponseParams `json:"Response"`
}

func (r *DescribeIpGeolocationDatabaseUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIpGeolocationDatabaseUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIpGeolocationInfosRequestParams struct {
	// The list of IP addresses (only IPv4 addresses are available currently) to be queried; upper limit: 100
	AddressIps []*string `json:"AddressIps,omitnil,omitempty" name:"AddressIps"`

	// Fields of the IP addresses to be queried.
	Fields *IpField `json:"Fields,omitnil,omitempty" name:"Fields"`
}

type DescribeIpGeolocationInfosRequest struct {
	*tchttp.BaseRequest
	
	// The list of IP addresses (only IPv4 addresses are available currently) to be queried; upper limit: 100
	AddressIps []*string `json:"AddressIps,omitnil,omitempty" name:"AddressIps"`

	// Fields of the IP addresses to be queried.
	Fields *IpField `json:"Fields,omitnil,omitempty" name:"Fields"`
}

func (r *DescribeIpGeolocationInfosRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIpGeolocationInfosRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AddressIps")
	delete(f, "Fields")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIpGeolocationInfosRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIpGeolocationInfosResponseParams struct {
	// IP address details
	AddressInfo []*IpGeolocationInfo `json:"AddressInfo,omitnil,omitempty" name:"AddressInfo"`

	// Number of IP addresses
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeIpGeolocationInfosResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIpGeolocationInfosResponseParams `json:"Response"`
}

func (r *DescribeIpGeolocationInfosResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIpGeolocationInfosResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLocalGatewayRequestParams struct {
	// Query criteria:
	// vpc-id: filter by VPC ID; local-gateway-name: filter by local gateway name (fuzzy search is supported); local-gateway-id: filter by local gateway instance ID; cdc-id: filter by CDC instance ID.
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// The offset. Default value: 0. For more information on `Offset`, see the relevant sections in API [Introduction](https://intl.cloud.tencent.com/document/api/213/11646?from_cn_redirect=1).
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The number of returned results. Default value: 20. Maximum value: 100. For more information on `Limit`, see the relevant sections in API [Introduction](https://intl.cloud.tencent.com/document/api/213/11646?from_cn_redirect=1).
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeLocalGatewayRequest struct {
	*tchttp.BaseRequest
	
	// Query criteria:
	// vpc-id: filter by VPC ID; local-gateway-name: filter by local gateway name (fuzzy search is supported); local-gateway-id: filter by local gateway instance ID; cdc-id: filter by CDC instance ID.
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// The offset. Default value: 0. For more information on `Offset`, see the relevant sections in API [Introduction](https://intl.cloud.tencent.com/document/api/213/11646?from_cn_redirect=1).
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The number of returned results. Default value: 20. Maximum value: 100. For more information on `Limit`, see the relevant sections in API [Introduction](https://intl.cloud.tencent.com/document/api/213/11646?from_cn_redirect=1).
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeLocalGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLocalGatewayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLocalGatewayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLocalGatewayResponseParams struct {
	// Information set of local gateways
	LocalGatewaySet []*LocalGateway `json:"LocalGatewaySet,omitnil,omitempty" name:"LocalGatewaySet"`

	// Total number of local gateways
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLocalGatewayResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLocalGatewayResponseParams `json:"Response"`
}

func (r *DescribeLocalGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLocalGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNatGatewayDestinationIpPortTranslationNatRulesRequestParams struct {
	// NAT gateway ID.
	NatGatewayIds []*string `json:"NatGatewayIds,omitnil,omitempty" name:"NatGatewayIds"`

	// Filters:
	// `NatGatewayIds` and `Filters` cannot be specified at the same time.
	// <li> `nat-gateway-id`: The NAT gateway ID, such as `nat-0yi4hekt`.</li>
	// <li> `vpc-id`: The VPC ID, such as `vpc-0yi4hekt`.</li>
	// <li> `public-ip-address`: The EIP, such as `139.199.232.238`.</li>
	// <li>`public-port`: The public network port.</li>
	// <li>`private-ip-address`: The private IP, such as `10.0.0.1`.</li>
	// <li>`private-port`. The private network port.</li>
	// <li>`description`. The rule description.</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeNatGatewayDestinationIpPortTranslationNatRulesRequest struct {
	*tchttp.BaseRequest
	
	// NAT gateway ID.
	NatGatewayIds []*string `json:"NatGatewayIds,omitnil,omitempty" name:"NatGatewayIds"`

	// Filters:
	// `NatGatewayIds` and `Filters` cannot be specified at the same time.
	// <li> `nat-gateway-id`: The NAT gateway ID, such as `nat-0yi4hekt`.</li>
	// <li> `vpc-id`: The VPC ID, such as `vpc-0yi4hekt`.</li>
	// <li> `public-ip-address`: The EIP, such as `139.199.232.238`.</li>
	// <li>`public-port`: The public network port.</li>
	// <li>`private-ip-address`: The private IP, such as `10.0.0.1`.</li>
	// <li>`private-port`. The private network port.</li>
	// <li>`description`. The rule description.</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeNatGatewayDestinationIpPortTranslationNatRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNatGatewayDestinationIpPortTranslationNatRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NatGatewayIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNatGatewayDestinationIpPortTranslationNatRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNatGatewayDestinationIpPortTranslationNatRulesResponseParams struct {
	// The object array of port forwarding rules for the NAT gateway.
	NatGatewayDestinationIpPortTranslationNatRuleSet []*NatGatewayDestinationIpPortTranslationNatRule `json:"NatGatewayDestinationIpPortTranslationNatRuleSet,omitnil,omitempty" name:"NatGatewayDestinationIpPortTranslationNatRuleSet"`

	// The number of eligible object arrays of NAT port forwarding rules.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeNatGatewayDestinationIpPortTranslationNatRulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNatGatewayDestinationIpPortTranslationNatRulesResponseParams `json:"Response"`
}

func (r *DescribeNatGatewayDestinationIpPortTranslationNatRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNatGatewayDestinationIpPortTranslationNatRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNatGatewayDirectConnectGatewayRouteRequestParams struct {
	// Unique ID of the NAT gateway
	NatGatewayId *string `json:"NatGatewayId,omitnil,omitempty" name:"NatGatewayId"`

	// Unique ID of VPC
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Valid range: 0-200
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Greater than 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeNatGatewayDirectConnectGatewayRouteRequest struct {
	*tchttp.BaseRequest
	
	// Unique ID of the NAT gateway
	NatGatewayId *string `json:"NatGatewayId,omitnil,omitempty" name:"NatGatewayId"`

	// Unique ID of VPC
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Valid range: 0-200
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Greater than 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeNatGatewayDirectConnectGatewayRouteRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNatGatewayDirectConnectGatewayRouteRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NatGatewayId")
	delete(f, "VpcId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNatGatewayDirectConnectGatewayRouteRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNatGatewayDirectConnectGatewayRouteResponseParams struct {
	// Route data
	NatDirectConnectGatewayRouteSet []*NatDirectConnectGatewayRoute `json:"NatDirectConnectGatewayRouteSet,omitnil,omitempty" name:"NatDirectConnectGatewayRouteSet"`

	// Total number of routes
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeNatGatewayDirectConnectGatewayRouteResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNatGatewayDirectConnectGatewayRouteResponseParams `json:"Response"`
}

func (r *DescribeNatGatewayDirectConnectGatewayRouteResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNatGatewayDirectConnectGatewayRouteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNatGatewaySourceIpTranslationNatRulesRequestParams struct {
	// The unique ID of the NAT Gateway, such as `nat-123xx454`.
	NatGatewayId *string `json:"NatGatewayId,omitnil,omitempty" name:"NatGatewayId"`

	// Filter:
	// <li>`resource-id`: The subnet ID (such as `subnet-0yi4hekt`) or CVM ID</li>
	// <li>`public-ip-address`: The EIP, such as `139.199.232.238`</li>
	// <li>`description` The rule description</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeNatGatewaySourceIpTranslationNatRulesRequest struct {
	*tchttp.BaseRequest
	
	// The unique ID of the NAT Gateway, such as `nat-123xx454`.
	NatGatewayId *string `json:"NatGatewayId,omitnil,omitempty" name:"NatGatewayId"`

	// Filter:
	// <li>`resource-id`: The subnet ID (such as `subnet-0yi4hekt`) or CVM ID</li>
	// <li>`public-ip-address`: The EIP, such as `139.199.232.238`</li>
	// <li>`description` The rule description</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeNatGatewaySourceIpTranslationNatRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNatGatewaySourceIpTranslationNatRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NatGatewayId")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNatGatewaySourceIpTranslationNatRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNatGatewaySourceIpTranslationNatRulesResponseParams struct {
	// Array of objects of a NAT gateway's SNAT rules.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	SourceIpTranslationNatRuleSet []*SourceIpTranslationNatRule `json:"SourceIpTranslationNatRuleSet,omitnil,omitempty" name:"SourceIpTranslationNatRuleSet"`

	// The number of eligible object arrays of a NAT gateway's forwarding rules.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeNatGatewaySourceIpTranslationNatRulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNatGatewaySourceIpTranslationNatRulesResponseParams `json:"Response"`
}

func (r *DescribeNatGatewaySourceIpTranslationNatRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNatGatewaySourceIpTranslationNatRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNatGatewaysRequestParams struct {
	// The unified ID of the NAT gateways, such as `nat-123xx454`.
	NatGatewayIds []*string `json:"NatGatewayIds,omitnil,omitempty" name:"NatGatewayIds"`

	// Filters. `NatGatewayIds` and `Filters` cannot be specified at the same time.
	// <li>nat-gateway-id - String - (Filter) The ID of the protocol port template instance, such as `nat-123xx454`.</li>
	// <li>vpc-id - String - (Filter) The unique ID of the VPC, such as `vpc-123xx454`.</li>
	// <li>nat-gateway-name - String - (Filter) The ID of the protocol port template instance, such as `test_nat`.</li>
	// <li>tag-key - String - (Filter) The tag key, such as `test-key`.</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeNatGatewaysRequest struct {
	*tchttp.BaseRequest
	
	// The unified ID of the NAT gateways, such as `nat-123xx454`.
	NatGatewayIds []*string `json:"NatGatewayIds,omitnil,omitempty" name:"NatGatewayIds"`

	// Filters. `NatGatewayIds` and `Filters` cannot be specified at the same time.
	// <li>nat-gateway-id - String - (Filter) The ID of the protocol port template instance, such as `nat-123xx454`.</li>
	// <li>vpc-id - String - (Filter) The unique ID of the VPC, such as `vpc-123xx454`.</li>
	// <li>nat-gateway-name - String - (Filter) The ID of the protocol port template instance, such as `test_nat`.</li>
	// <li>tag-key - String - (Filter) The tag key, such as `test-key`.</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeNatGatewaysRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNatGatewaysRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NatGatewayIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNatGatewaysRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNatGatewaysResponseParams struct {
	// NAT gateway object array.
	NatGatewaySet []*NatGateway `json:"NatGatewaySet,omitnil,omitempty" name:"NatGatewaySet"`

	// The number of eligible NAT gateway objects.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeNatGatewaysResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNatGatewaysResponseParams `json:"Response"`
}

func (r *DescribeNatGatewaysResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNatGatewaysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNetDetectStatesRequestParams struct {
	// The array of network probe IDs, such as [`netd-12345678`].
	NetDetectIds []*string `json:"NetDetectIds,omitnil,omitempty" name:"NetDetectIds"`

	// Filter conditions. `NetDetectIds` and `Filters` cannot be specified at the same time.
	// <li>`net-detect-id` - String - The network probe ID, such as `netd-12345678`.</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// The offset. Default: 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The number of returned values. Default: 20. Maximum: 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeNetDetectStatesRequest struct {
	*tchttp.BaseRequest
	
	// The array of network probe IDs, such as [`netd-12345678`].
	NetDetectIds []*string `json:"NetDetectIds,omitnil,omitempty" name:"NetDetectIds"`

	// Filter conditions. `NetDetectIds` and `Filters` cannot be specified at the same time.
	// <li>`net-detect-id` - String - The network probe ID, such as `netd-12345678`.</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// The offset. Default: 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The number of returned values. Default: 20. Maximum: 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeNetDetectStatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNetDetectStatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NetDetectIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNetDetectStatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNetDetectStatesResponseParams struct {
	// The array of network detection verification results that meet requirements.
	// Note: This field may return null, indicating that no valid values can be obtained.
	NetDetectStateSet []*NetDetectState `json:"NetDetectStateSet,omitnil,omitempty" name:"NetDetectStateSet"`

	// The number of network detection verification results that meet requirements.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeNetDetectStatesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNetDetectStatesResponseParams `json:"Response"`
}

func (r *DescribeNetDetectStatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNetDetectStatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNetDetectsRequestParams struct {
	// The array of network probe IDs, such as [`netd-12345678`].
	NetDetectIds []*string `json:"NetDetectIds,omitnil,omitempty" name:"NetDetectIds"`

	// Filter conditions. `NetDetectIds` and `Filters` cannot be specified at the same time.
	// <li>vpc-id - String - (Filter condition) The VPC instance ID, such as vpc-12345678.</li>
	// <li>net-detect-id - String - (Filter condition) The network detection instance ID, such as netd-12345678.</li>
	// <li>subnet-id - String - (Filter condition) The subnet instance ID, such as subnet-12345678.</li>
	// <li>net-detect-name - String - (Filter condition) The network detection name.</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// The offset. Default: 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The number of returned values. Default: 20. Maximum: 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeNetDetectsRequest struct {
	*tchttp.BaseRequest
	
	// The array of network probe IDs, such as [`netd-12345678`].
	NetDetectIds []*string `json:"NetDetectIds,omitnil,omitempty" name:"NetDetectIds"`

	// Filter conditions. `NetDetectIds` and `Filters` cannot be specified at the same time.
	// <li>vpc-id - String - (Filter condition) The VPC instance ID, such as vpc-12345678.</li>
	// <li>net-detect-id - String - (Filter condition) The network detection instance ID, such as netd-12345678.</li>
	// <li>subnet-id - String - (Filter condition) The subnet instance ID, such as subnet-12345678.</li>
	// <li>net-detect-name - String - (Filter condition) The network detection name.</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// The offset. Default: 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The number of returned values. Default: 20. Maximum: 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeNetDetectsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNetDetectsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NetDetectIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNetDetectsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNetDetectsResponseParams struct {
	// The array of network detection objects that meet requirements.
	// Note: This field may return null, indicating that no valid values can be obtained.
	NetDetectSet []*NetDetect `json:"NetDetectSet,omitnil,omitempty" name:"NetDetectSet"`

	// The number of network detection objects that meet requirements.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeNetDetectsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNetDetectsResponseParams `json:"Response"`
}

func (r *DescribeNetDetectsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNetDetectsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNetworkAclQuintupleEntriesRequestParams struct {
	// Network ACL instance ID, such as `acl-12345678`.
	NetworkAclId *string `json:"NetworkAclId,omitnil,omitempty" name:"NetworkAclId"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Returned quantity. Default: 20. Value range: 1-100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Filter condition. `HaVipIds` and `Filters` cannot be specified at the same time.
	// <li>`protocol` - String - Such as `TCP`</li>
	// <li>`description` - String - Description</li>
	// <li>`destination-cidr` - String - Destination CIDR block, such as `192.168.0.0/24`</li>
	// <li>`source-cidr` - String - Source CIDR block, such as `192.168.0.0/24`</li>
	// <li>`action` - String - ·Values: `ACCEPT`, `DROP`</li>
	// <li>`network-acl-quintuple-entry-id` - String - Unique ID of the quintuple, such as `acli45-ahnu4rv5`</li>
	// <li>`network-acl-direction` - String - Direction of the policy. Values: `INGRESS` or `EGRESS`.</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeNetworkAclQuintupleEntriesRequest struct {
	*tchttp.BaseRequest
	
	// Network ACL instance ID, such as `acl-12345678`.
	NetworkAclId *string `json:"NetworkAclId,omitnil,omitempty" name:"NetworkAclId"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Returned quantity. Default: 20. Value range: 1-100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Filter condition. `HaVipIds` and `Filters` cannot be specified at the same time.
	// <li>`protocol` - String - Such as `TCP`</li>
	// <li>`description` - String - Description</li>
	// <li>`destination-cidr` - String - Destination CIDR block, such as `192.168.0.0/24`</li>
	// <li>`source-cidr` - String - Source CIDR block, such as `192.168.0.0/24`</li>
	// <li>`action` - String - ·Values: `ACCEPT`, `DROP`</li>
	// <li>`network-acl-quintuple-entry-id` - String - Unique ID of the quintuple, such as `acli45-ahnu4rv5`</li>
	// <li>`network-acl-direction` - String - Direction of the policy. Values: `INGRESS` or `EGRESS`.</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeNetworkAclQuintupleEntriesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNetworkAclQuintupleEntriesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NetworkAclId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNetworkAclQuintupleEntriesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNetworkAclQuintupleEntriesResponseParams struct {
	// The list of the network ACL quintuple entries
	NetworkAclQuintupleSet []*NetworkAclQuintupleEntry `json:"NetworkAclQuintupleSet,omitnil,omitempty" name:"NetworkAclQuintupleSet"`

	// Number of eligible instances.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeNetworkAclQuintupleEntriesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNetworkAclQuintupleEntriesResponseParams `json:"Response"`
}

func (r *DescribeNetworkAclQuintupleEntriesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNetworkAclQuintupleEntriesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNetworkAclsRequestParams struct {
	// Array of network ACL instance IDs, such as [acl-12345678]. Up to 100 instances are allowed for each request. This parameter does not support specifying `NetworkAclIds` and `Filters` at the same time.
	NetworkAclIds []*string `json:"NetworkAclIds,omitnil,omitempty" name:"NetworkAclIds"`

	// Filter condition. `NetworkAclIds` and `Filters` cannot be specified at the same time.
	// <li>vpc-id - String - (Filter condition) VPC instance ID, such as vpc-12345678.</li>
	// <li>network-acl-id - String - (Filter condition) Network ACL instance ID, such as acl-12345678.</li>
	// <li>network-acl-name - String - (Filter condition) Network ACL instance name.</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Offset. Default: 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Returned quantity. Default: 20. Value range: 1-100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeNetworkAclsRequest struct {
	*tchttp.BaseRequest
	
	// Array of network ACL instance IDs, such as [acl-12345678]. Up to 100 instances are allowed for each request. This parameter does not support specifying `NetworkAclIds` and `Filters` at the same time.
	NetworkAclIds []*string `json:"NetworkAclIds,omitnil,omitempty" name:"NetworkAclIds"`

	// Filter condition. `NetworkAclIds` and `Filters` cannot be specified at the same time.
	// <li>vpc-id - String - (Filter condition) VPC instance ID, such as vpc-12345678.</li>
	// <li>network-acl-id - String - (Filter condition) Network ACL instance ID, such as acl-12345678.</li>
	// <li>network-acl-name - String - (Filter condition) Network ACL instance name.</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Offset. Default: 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Returned quantity. Default: 20. Value range: 1-100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeNetworkAclsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNetworkAclsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NetworkAclIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNetworkAclsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNetworkAclsResponseParams struct {
	// List of instance details.
	NetworkAclSet []*NetworkAcl `json:"NetworkAclSet,omitnil,omitempty" name:"NetworkAclSet"`

	// Number of eligible instances.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeNetworkAclsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNetworkAclsResponseParams `json:"Response"`
}

func (r *DescribeNetworkAclsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNetworkAclsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNetworkInterfaceLimitRequestParams struct {
	// ID of a CVM instance or ENI to query
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeNetworkInterfaceLimitRequest struct {
	*tchttp.BaseRequest
	
	// ID of a CVM instance or ENI to query
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeNetworkInterfaceLimitRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNetworkInterfaceLimitRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNetworkInterfaceLimitRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNetworkInterfaceLimitResponseParams struct {
	// Quota of ENIs mounted to a CVM instance in a standard way
	EniQuantity *int64 `json:"EniQuantity,omitnil,omitempty" name:"EniQuantity"`

	// Quota of IP addresses that can be allocated to each standard-mounted ENI
	EniPrivateIpAddressQuantity *int64 `json:"EniPrivateIpAddressQuantity,omitnil,omitempty" name:"EniPrivateIpAddressQuantity"`

	// Quota of ENIs mounted to a CVM instance as an extension
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ExtendEniQuantity *int64 `json:"ExtendEniQuantity,omitnil,omitempty" name:"ExtendEniQuantity"`

	// Quota of IP addresses that can be allocated to each extension-mounted ENI.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ExtendEniPrivateIpAddressQuantity *int64 `json:"ExtendEniPrivateIpAddressQuantity,omitnil,omitempty" name:"ExtendEniPrivateIpAddressQuantity"`

	// The quota of relayed ENIs
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	SubEniQuantity *int64 `json:"SubEniQuantity,omitnil,omitempty" name:"SubEniQuantity"`

	// The quota of IPs that can be assigned to each relayed ENI.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	SubEniPrivateIpAddressQuantity *int64 `json:"SubEniPrivateIpAddressQuantity,omitnil,omitempty" name:"SubEniPrivateIpAddressQuantity"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeNetworkInterfaceLimitResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNetworkInterfaceLimitResponseParams `json:"Response"`
}

func (r *DescribeNetworkInterfaceLimitResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNetworkInterfaceLimitResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNetworkInterfacesRequestParams struct {
	// Queries the ID of the ENI instance, such as `eni-pxir56ns`. Each request can have a maximum of 100 instances. `NetworkInterfaceIds` and `Filters` cannot be specified at the same time.
	NetworkInterfaceIds []*string `json:"NetworkInterfaceIds,omitnil,omitempty" name:"NetworkInterfaceIds"`

	// Filter. `NetworkInterfaceIds` and `Filters` cannot be specified at the same time.
	// <li>`vpc-id` - String - VPC ID, such as `vpc-f49l6u0z`. </li>
	// <li>`subnet-id` - String - Subnet ID, such as `subnet-f49l6u0z`. </li>
	// <li>`network-interface-id` - String - ENI ID, such as `eni-5k56k7k7`. </li>
	// <li>`attachment.instance-id` - String - ID of the bound CVM instance, such as `ins-3nqpdn3i`. </li>
	// <li>`groups.security-group-id` - String - ID of the bound security group, such as `sg-f9ekbxeq`. </li>
	// <li>`network-interface-name` - String - ENI instance name. </li>
	// <li>`network-interface-description` - String - ENI instance description. </li>
	// <li>`address-ip` - String - Private IPv4 address. A single IP will be fuzzily matched with the suffix, while multiple IPs will be exactly matched. It can be used with `ip-exact-match` to query and exactly match a single IP. </li>
	// <li>`ip-exact-match` - Boolean - Exact match by private IPv4 address. The first value will be returned if multiple values are found. </li>
	// <li>`tag-key` - String - u200dOptional - u200dTag key. See Example 2 to learn more details. </li>
	// <li>`tag:tag-key` - String - Optional - Tag key-value pair. The `tag-key` should be replaced with a specific tag key. See Example 2 to learn more details. </li>
	// <li>`is-primary` - Boolean - Optional - Filter based on whether it is a primary ENI. Values: `true`, `false`. If this parameter is not specified, filter the both. </li>
	// <li>`eni-type` - String - Optional - Filter by ENI type. Values: `0` (Secondary ENI), `1` (Primary ENI), `2` (Relayed ENI) </li>
	// <li>`eni-qos` - String - Optional - Filter by ENI service level. Values: `AG` (Bronze), `AU` (Silver) </li>
	// <li>`address-ipv6` - String - Optional - Filter by private IPv6 address. Multiple IPv6 addresses can be used for query. If this field is used together with `address-ip`, their intersection will be used. </li>
	// <li>`public-address-ip` - String - Public IPv4 address. It supports exact matching. </li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeNetworkInterfacesRequest struct {
	*tchttp.BaseRequest
	
	// Queries the ID of the ENI instance, such as `eni-pxir56ns`. Each request can have a maximum of 100 instances. `NetworkInterfaceIds` and `Filters` cannot be specified at the same time.
	NetworkInterfaceIds []*string `json:"NetworkInterfaceIds,omitnil,omitempty" name:"NetworkInterfaceIds"`

	// Filter. `NetworkInterfaceIds` and `Filters` cannot be specified at the same time.
	// <li>`vpc-id` - String - VPC ID, such as `vpc-f49l6u0z`. </li>
	// <li>`subnet-id` - String - Subnet ID, such as `subnet-f49l6u0z`. </li>
	// <li>`network-interface-id` - String - ENI ID, such as `eni-5k56k7k7`. </li>
	// <li>`attachment.instance-id` - String - ID of the bound CVM instance, such as `ins-3nqpdn3i`. </li>
	// <li>`groups.security-group-id` - String - ID of the bound security group, such as `sg-f9ekbxeq`. </li>
	// <li>`network-interface-name` - String - ENI instance name. </li>
	// <li>`network-interface-description` - String - ENI instance description. </li>
	// <li>`address-ip` - String - Private IPv4 address. A single IP will be fuzzily matched with the suffix, while multiple IPs will be exactly matched. It can be used with `ip-exact-match` to query and exactly match a single IP. </li>
	// <li>`ip-exact-match` - Boolean - Exact match by private IPv4 address. The first value will be returned if multiple values are found. </li>
	// <li>`tag-key` - String - u200dOptional - u200dTag key. See Example 2 to learn more details. </li>
	// <li>`tag:tag-key` - String - Optional - Tag key-value pair. The `tag-key` should be replaced with a specific tag key. See Example 2 to learn more details. </li>
	// <li>`is-primary` - Boolean - Optional - Filter based on whether it is a primary ENI. Values: `true`, `false`. If this parameter is not specified, filter the both. </li>
	// <li>`eni-type` - String - Optional - Filter by ENI type. Values: `0` (Secondary ENI), `1` (Primary ENI), `2` (Relayed ENI) </li>
	// <li>`eni-qos` - String - Optional - Filter by ENI service level. Values: `AG` (Bronze), `AU` (Silver) </li>
	// <li>`address-ipv6` - String - Optional - Filter by private IPv6 address. Multiple IPv6 addresses can be used for query. If this field is used together with `address-ip`, their intersection will be used. </li>
	// <li>`public-address-ip` - String - Public IPv4 address. It supports exact matching. </li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNetworkInterfacesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNetworkInterfacesResponseParams struct {
	// List of instance details.
	NetworkInterfaceSet []*NetworkInterface `json:"NetworkInterfaceSet,omitnil,omitempty" name:"NetworkInterfaceSet"`

	// Number of eligible instances.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeRouteTablesRequestParams struct {
	// Filter condition. `RouteTableIds` and `Filters` cannot be specified at the same time.
	// <li>route-table-id - String - (Filter condition) Route table instance ID.</li>
	// <li>route-table-name - String - (Filter condition) Route table name.</li>
	// <li>vpc-id - String - (Filter condition) VPC instance ID, such as `vpc-f49l6u0z`.</li>
	// <li>association.main - String - (Filter condition) Whether it is the main route table.</li>
	// <li>tag-key - String - Required: no - (Filter condition) Filter by tag key.</li>
	// <li>tag:tag-key - String - Required: no - (Filter condition) Filter by tag key-value pair. Use a specific tag key to replace `tag-key`. See Example 2 for the detailed usage.</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// The route table instance ID, such as `rtb-azd4dt1c`.
	RouteTableIds []*string `json:"RouteTableIds,omitnil,omitempty" name:"RouteTableIds"`

	// Offset.
	Offset *string `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The number of request objects.
	Limit *string `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeRouteTablesRequest struct {
	*tchttp.BaseRequest
	
	// Filter condition. `RouteTableIds` and `Filters` cannot be specified at the same time.
	// <li>route-table-id - String - (Filter condition) Route table instance ID.</li>
	// <li>route-table-name - String - (Filter condition) Route table name.</li>
	// <li>vpc-id - String - (Filter condition) VPC instance ID, such as `vpc-f49l6u0z`.</li>
	// <li>association.main - String - (Filter condition) Whether it is the main route table.</li>
	// <li>tag-key - String - Required: no - (Filter condition) Filter by tag key.</li>
	// <li>tag:tag-key - String - Required: no - (Filter condition) Filter by tag key-value pair. Use a specific tag key to replace `tag-key`. See Example 2 for the detailed usage.</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// The route table instance ID, such as `rtb-azd4dt1c`.
	RouteTableIds []*string `json:"RouteTableIds,omitnil,omitempty" name:"RouteTableIds"`

	// Offset.
	Offset *string `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The number of request objects.
	Limit *string `json:"Limit,omitnil,omitempty" name:"Limit"`
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
	delete(f, "Filters")
	delete(f, "RouteTableIds")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRouteTablesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRouteTablesResponseParams struct {
	// Number of eligible instances.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Route table object.
	RouteTableSet []*RouteTable `json:"RouteTableSet,omitnil,omitempty" name:"RouteTableSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// The Security instance ID, such as `sg-33ocnj9n`. It can be obtained through DescribeSecurityGroups.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`
}

type DescribeSecurityGroupAssociationStatisticsRequest struct {
	*tchttp.BaseRequest
	
	// The Security instance ID, such as `sg-33ocnj9n`. It can be obtained through DescribeSecurityGroups.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`
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
	// Statistics on the instances associated with a security group.
	SecurityGroupAssociationStatisticsSet []*SecurityGroupAssociationStatistics `json:"SecurityGroupAssociationStatisticsSet,omitnil,omitempty" name:"SecurityGroupAssociationStatisticsSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeSecurityGroupPoliciesRequestParams struct {
	// The security group instance ID, such as `sg-33ocnj9n`. It can be obtained through DescribeSecurityGroups.
	SecurityGroupId *string `json:"SecurityGroupId,omitnil,omitempty" name:"SecurityGroupId"`

	// Filters
	// <li>`security-group-id` - String - Security group ID in the rule.</li>
	// <li>`ip` - String - IP. IPV4 and IPV6 fuzzy matching is supported.</li>
	// <li>`address-module` - String - IP address or address group template ID.</li>
	// <li>`service-module` - String - Protocol port or port group template ID.</li>
	// <li>`protocol-type` - String - Protocol supported by the security group policy. Valid values: `TCP`, `UDP`, `ICMP`, `ICMPV6`, `GRE`, `ALL`.</li>
	// <li>`port` - String - Optional - Port. Fuzzy matching is supported. Query all ports when the protocol value is `ALL`.</li>
	// <li>`poly` - String - Policy type. Valid values: `ALL`, `ACCEPT` and `DROP`.</li>
	// <li>`direction` - String - Direction of the rule. Valid values: `ALL`, `INBOUND` and `OUTBOUND`.</li>
	// <li>`description` - String - Policy description. Fuzzy matching is supported.</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeSecurityGroupPoliciesRequest struct {
	*tchttp.BaseRequest
	
	// The security group instance ID, such as `sg-33ocnj9n`. It can be obtained through DescribeSecurityGroups.
	SecurityGroupId *string `json:"SecurityGroupId,omitnil,omitempty" name:"SecurityGroupId"`

	// Filters
	// <li>`security-group-id` - String - Security group ID in the rule.</li>
	// <li>`ip` - String - IP. IPV4 and IPV6 fuzzy matching is supported.</li>
	// <li>`address-module` - String - IP address or address group template ID.</li>
	// <li>`service-module` - String - Protocol port or port group template ID.</li>
	// <li>`protocol-type` - String - Protocol supported by the security group policy. Valid values: `TCP`, `UDP`, `ICMP`, `ICMPV6`, `GRE`, `ALL`.</li>
	// <li>`port` - String - Optional - Port. Fuzzy matching is supported. Query all ports when the protocol value is `ALL`.</li>
	// <li>`poly` - String - Policy type. Valid values: `ALL`, `ACCEPT` and `DROP`.</li>
	// <li>`direction` - String - Direction of the rule. Valid values: `ALL`, `INBOUND` and `OUTBOUND`.</li>
	// <li>`description` - String - Policy description. Fuzzy matching is supported.</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
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
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecurityGroupPoliciesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityGroupPoliciesResponseParams struct {
	// Security group policy set.
	SecurityGroupPolicySet *SecurityGroupPolicySet `json:"SecurityGroupPolicySet,omitnil,omitempty" name:"SecurityGroupPolicySet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeSecurityGroupReferencesRequestParams struct {
	// A set of security group instance IDs, e.g. ['sg-12345678']
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`
}

type DescribeSecurityGroupReferencesRequest struct {
	*tchttp.BaseRequest
	
	// A set of security group instance IDs, e.g. ['sg-12345678']
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`
}

func (r *DescribeSecurityGroupReferencesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityGroupReferencesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecurityGroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecurityGroupReferencesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityGroupReferencesResponseParams struct {
	// Referred security groups.
	ReferredSecurityGroupSet []*ReferredSecurityGroup `json:"ReferredSecurityGroupSet,omitnil,omitempty" name:"ReferredSecurityGroupSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSecurityGroupReferencesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSecurityGroupReferencesResponseParams `json:"Response"`
}

func (r *DescribeSecurityGroupReferencesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityGroupReferencesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityGroupsRequestParams struct {
	// Security group ID, such as `sg-33ocnj9n`. Each request can contain up to 100 instances at a time. `SecurityGroupIds` and `Filters` cannot be specified at the same time.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// Filter conditions. `SecurityGroupIds` and `Filters` cannot be specified at the same time.
	// <li>security-group-id - String - (Filter condition) The security group ID.</li>
	// <li>project-id - Integer - (Filter condition) The project ID.</li>
	// <li>security-group-name - String - (Filter condition) The security group name.</li>
	// <li>tag-key - String - Required: no - (Filter condition) Filters by tag key. For more information, see Example 2.</li>
	// <li> `tag:tag-key` - String - Required: no - (Filter condition) Filters by tag key pair. For this parameter, `tag-key` will be replaced with a specific tag key. For more information, see Example 3.</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Offset. Default value: 0.
	Offset *string `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *string `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Sorting field. Values: `CreatedTime`, `UpdateTime` Note: This field does not have default value.
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`

	// Sorting method Order methods. Ascending: `ASC`, Descending: `DESC`. Default: `ASC`
	OrderDirection *string `json:"OrderDirection,omitnil,omitempty" name:"OrderDirection"`
}

type DescribeSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// Security group ID, such as `sg-33ocnj9n`. Each request can contain up to 100 instances at a time. `SecurityGroupIds` and `Filters` cannot be specified at the same time.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// Filter conditions. `SecurityGroupIds` and `Filters` cannot be specified at the same time.
	// <li>security-group-id - String - (Filter condition) The security group ID.</li>
	// <li>project-id - Integer - (Filter condition) The project ID.</li>
	// <li>security-group-name - String - (Filter condition) The security group name.</li>
	// <li>tag-key - String - Required: no - (Filter condition) Filters by tag key. For more information, see Example 2.</li>
	// <li> `tag:tag-key` - String - Required: no - (Filter condition) Filters by tag key pair. For this parameter, `tag-key` will be replaced with a specific tag key. For more information, see Example 3.</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Offset. Default value: 0.
	Offset *string `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *string `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Sorting field. Values: `CreatedTime`, `UpdateTime` Note: This field does not have default value.
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`

	// Sorting method Order methods. Ascending: `ASC`, Descending: `DESC`. Default: `ASC`
	OrderDirection *string `json:"OrderDirection,omitnil,omitempty" name:"OrderDirection"`
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
	delete(f, "OrderField")
	delete(f, "OrderDirection")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecurityGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityGroupsResponseParams struct {
	// Security group object.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	SecurityGroupSet []*SecurityGroup `json:"SecurityGroupSet,omitnil,omitempty" name:"SecurityGroupSet"`

	// The number of instances meeting the filter condition.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeServiceTemplateGroupsRequestParams struct {
	// Filter conditions.
	// <li>service-template-group-name - String - (Filter condition) Protocol port template group name.</li>
	// <li>service-template-group-id - String - (Filter condition) Protocol port template group instance ID, such as `ppmg-e6dy460g`.</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Offset. The default value is 0.
	Offset *string `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of values to be returned. The default value is 20. Maximum is 100.
	Limit *string `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeServiceTemplateGroupsRequest struct {
	*tchttp.BaseRequest
	
	// Filter conditions.
	// <li>service-template-group-name - String - (Filter condition) Protocol port template group name.</li>
	// <li>service-template-group-id - String - (Filter condition) Protocol port template group instance ID, such as `ppmg-e6dy460g`.</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Offset. The default value is 0.
	Offset *string `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of values to be returned. The default value is 20. Maximum is 100.
	Limit *string `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeServiceTemplateGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServiceTemplateGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeServiceTemplateGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeServiceTemplateGroupsResponseParams struct {
	// The number of instances meeting the filter condition.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Protocol port template group.
	ServiceTemplateGroupSet []*ServiceTemplateGroup `json:"ServiceTemplateGroupSet,omitnil,omitempty" name:"ServiceTemplateGroupSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeServiceTemplateGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeServiceTemplateGroupsResponseParams `json:"Response"`
}

func (r *DescribeServiceTemplateGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServiceTemplateGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeServiceTemplatesRequestParams struct {
	// Filters
	// <li>service-template-name - Protocol port template name.</li>
	// <li>service-template-id - Protocol port template ID, such as `ppm-e6dy460g`.</li>
	// <li>service-port-Protocol port.</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Offset. The default value is 0.
	Offset *string `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of values to be returned. The default value is 20. Maximum is 100.
	Limit *string `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeServiceTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// Filters
	// <li>service-template-name - Protocol port template name.</li>
	// <li>service-template-id - Protocol port template ID, such as `ppm-e6dy460g`.</li>
	// <li>service-port-Protocol port.</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Offset. The default value is 0.
	Offset *string `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of values to be returned. The default value is 20. Maximum is 100.
	Limit *string `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeServiceTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServiceTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeServiceTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeServiceTemplatesResponseParams struct {
	// The number of instances meeting the filter condition.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Protocol port template object.
	ServiceTemplateSet []*ServiceTemplate `json:"ServiceTemplateSet,omitnil,omitempty" name:"ServiceTemplateSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeServiceTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeServiceTemplatesResponseParams `json:"Response"`
}

func (r *DescribeServiceTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServiceTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSgSnapshotFileContentRequestParams struct {
	// Snapshot policy IDs
	SnapshotPolicyId *string `json:"SnapshotPolicyId,omitnil,omitempty" name:"SnapshotPolicyId"`

	// Snapshot file ID
	SnapshotFileId *string `json:"SnapshotFileId,omitnil,omitempty" name:"SnapshotFileId"`

	// Security group ID
	SecurityGroupId *string `json:"SecurityGroupId,omitnil,omitempty" name:"SecurityGroupId"`
}

type DescribeSgSnapshotFileContentRequest struct {
	*tchttp.BaseRequest
	
	// Snapshot policy IDs
	SnapshotPolicyId *string `json:"SnapshotPolicyId,omitnil,omitempty" name:"SnapshotPolicyId"`

	// Snapshot file ID
	SnapshotFileId *string `json:"SnapshotFileId,omitnil,omitempty" name:"SnapshotFileId"`

	// Security group ID
	SecurityGroupId *string `json:"SecurityGroupId,omitnil,omitempty" name:"SecurityGroupId"`
}

func (r *DescribeSgSnapshotFileContentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSgSnapshotFileContentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SnapshotPolicyId")
	delete(f, "SnapshotFileId")
	delete(f, "SecurityGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSgSnapshotFileContentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSgSnapshotFileContentResponseParams struct {
	// Security group ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Snapshot policy IDs
	SnapshotPolicyId *string `json:"SnapshotPolicyId,omitnil,omitempty" name:"SnapshotPolicyId"`

	// Snapshot file ID
	SnapshotFileId *string `json:"SnapshotFileId,omitnil,omitempty" name:"SnapshotFileId"`

	// Backup time
	BackupTime *string `json:"BackupTime,omitnil,omitempty" name:"BackupTime"`

	// Operator
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`

	// Original data
	OriginalData []*SecurityGroupPolicy `json:"OriginalData,omitnil,omitempty" name:"OriginalData"`

	// Backup data
	BackupData []*SecurityGroupPolicy `json:"BackupData,omitnil,omitempty" name:"BackupData"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSgSnapshotFileContentResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSgSnapshotFileContentResponseParams `json:"Response"`
}

func (r *DescribeSgSnapshotFileContentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSgSnapshotFileContentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSnapshotAttachedInstancesRequestParams struct {
	// Snapshot policy IDs
	SnapshotPolicyId *string `json:"SnapshotPolicyId,omitnil,omitempty" name:"SnapshotPolicyId"`

	// Filters
	// Values:
	// <li>`instance-id`: Instance ID</li>
	// <li>`instance-region`: Instance region</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Offset. Default value: `0`.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 200.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeSnapshotAttachedInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Snapshot policy IDs
	SnapshotPolicyId *string `json:"SnapshotPolicyId,omitnil,omitempty" name:"SnapshotPolicyId"`

	// Filters
	// Values:
	// <li>`instance-id`: Instance ID</li>
	// <li>`instance-region`: Instance region</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Offset. Default value: `0`.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 200.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeSnapshotAttachedInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSnapshotAttachedInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SnapshotPolicyId")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSnapshotAttachedInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSnapshotAttachedInstancesResponseParams struct {
	// List of instances
	InstanceSet []*SnapshotInstance `json:"InstanceSet,omitnil,omitempty" name:"InstanceSet"`

	// The number of eligible objects.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSnapshotAttachedInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSnapshotAttachedInstancesResponseParams `json:"Response"`
}

func (r *DescribeSnapshotAttachedInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSnapshotAttachedInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSnapshotFilesRequestParams struct {
	// Type of associated resource. Values: `securitygroup`
	BusinessType *string `json:"BusinessType,omitnil,omitempty" name:"BusinessType"`

	// Service instance ID. It's corresponding to the `BusinessType`. 
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Start time. Format: %Y-%m-%d %H:%M:%S
	StartDate *string `json:"StartDate,omitnil,omitempty" name:"StartDate"`

	// End time. Format: %Y-%m-%d %H:%M:%S
	EndDate *string `json:"EndDate,omitnil,omitempty" name:"EndDate"`

	// Offset. Default value: `0`.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeSnapshotFilesRequest struct {
	*tchttp.BaseRequest
	
	// Type of associated resource. Values: `securitygroup`
	BusinessType *string `json:"BusinessType,omitnil,omitempty" name:"BusinessType"`

	// Service instance ID. It's corresponding to the `BusinessType`. 
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Start time. Format: %Y-%m-%d %H:%M:%S
	StartDate *string `json:"StartDate,omitnil,omitempty" name:"StartDate"`

	// End time. Format: %Y-%m-%d %H:%M:%S
	EndDate *string `json:"EndDate,omitnil,omitempty" name:"EndDate"`

	// Offset. Default value: `0`.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeSnapshotFilesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSnapshotFilesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BusinessType")
	delete(f, "InstanceId")
	delete(f, "StartDate")
	delete(f, "EndDate")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSnapshotFilesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSnapshotFilesResponseParams struct {
	// Snapshot files
	SnapshotFileSet []*SnapshotFileInfo `json:"SnapshotFileSet,omitnil,omitempty" name:"SnapshotFileSet"`

	// The number of eligible objects.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSnapshotFilesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSnapshotFilesResponseParams `json:"Response"`
}

func (r *DescribeSnapshotFilesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSnapshotFilesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSnapshotPoliciesRequestParams struct {
	// Snapshot policy IDs
	SnapshotPolicyIds []*string `json:"SnapshotPolicyIds,omitnil,omitempty" name:"SnapshotPolicyIds"`

	// Filter conditions. `SnapshotPolicyIds` and `Filters` cannot be both specified.
	// <li>`snapshot-policy-id` - String - Snapshot policy ID</li>
	// <li>`snapshot-policy-name` - String - Snapshot policy name</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Offset. Default value: `0`.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 200.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeSnapshotPoliciesRequest struct {
	*tchttp.BaseRequest
	
	// Snapshot policy IDs
	SnapshotPolicyIds []*string `json:"SnapshotPolicyIds,omitnil,omitempty" name:"SnapshotPolicyIds"`

	// Filter conditions. `SnapshotPolicyIds` and `Filters` cannot be both specified.
	// <li>`snapshot-policy-id` - String - Snapshot policy ID</li>
	// <li>`snapshot-policy-name` - String - Snapshot policy name</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Offset. Default value: `0`.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 200.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeSnapshotPoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSnapshotPoliciesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SnapshotPolicyIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSnapshotPoliciesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSnapshotPoliciesResponseParams struct {
	// Snapshot policies
	SnapshotPolicySet []*SnapshotPolicy `json:"SnapshotPolicySet,omitnil,omitempty" name:"SnapshotPolicySet"`

	// The number of eligible objects.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSnapshotPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSnapshotPoliciesResponseParams `json:"Response"`
}

func (r *DescribeSnapshotPoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSnapshotPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSubnetResourceDashboardRequestParams struct {
	// Subnet instance ID, such as `subnet-f1xjkw1b`.
	SubnetIds []*string `json:"SubnetIds,omitnil,omitempty" name:"SubnetIds"`
}

type DescribeSubnetResourceDashboardRequest struct {
	*tchttp.BaseRequest
	
	// Subnet instance ID, such as `subnet-f1xjkw1b`.
	SubnetIds []*string `json:"SubnetIds,omitnil,omitempty" name:"SubnetIds"`
}

func (r *DescribeSubnetResourceDashboardRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSubnetResourceDashboardRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubnetIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSubnetResourceDashboardRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSubnetResourceDashboardResponseParams struct {
	// Information of resources returned
	ResourceStatisticsSet []*ResourceStatistics `json:"ResourceStatisticsSet,omitnil,omitempty" name:"ResourceStatisticsSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSubnetResourceDashboardResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSubnetResourceDashboardResponseParams `json:"Response"`
}

func (r *DescribeSubnetResourceDashboardResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSubnetResourceDashboardResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSubnetsRequestParams struct {
	// Queries the ID of the subnet instance, such as `subnet-pxir56ns`. Each request can have a maximum of 100 instances. `SubnetIds` and `Filters` cannot be specified at the same time.
	SubnetIds []*string `json:"SubnetIds,omitnil,omitempty" name:"SubnetIds"`

	// Filter condition. `SubnetIds` and `Filters` cannot be specified at the same time.
	// <li>subnet-id - String - (Filter condition) Subnet instance name.</li>
	// <li>vpc-id - String - (Filter condition) VPC instance ID, such as `vpc-f49l6u0z`.</li>
	// <li>cidr-block - String - (Filter condition) Subnet IP range, such as `192.168.1.0`.</li>
	// <li>is-default - Boolean - (Filter condition) Whether it is the default subnet.</li>
	// <li>is-remote-vpc-snat - Boolean - (Filter condition) Whether it is a VPC SNAT address pool subnet.</li>
	// <li>subnet-name - String - (Filter condition) Subnet name.</li>
	// <li>zone - String - (Filter condition) Availability zone.</li>
	// <li> tag-key - String - Required: No - (Filter condition) Filter by tag key.</li>
	// <li>tag:tag-key - String - Required: No - (Filter condition) Filter by tag key-value pair. Use a specific tag key to replace `tag-key`. For its usage, see example 2.</li>
	// <li>cdc-id - String - Required: No - (Filter condition) Filter by CDC ID to obtain subnets in the specified CDC.</li>
	// <li>is-cdc-subnet - String - Required: No - (Filter condition) Whether it is a CDC subnet. Valid values: `0` (no); `1` (yes).</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Offset. Default value: 0.
	Offset *string `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *string `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeSubnetsRequest struct {
	*tchttp.BaseRequest
	
	// Queries the ID of the subnet instance, such as `subnet-pxir56ns`. Each request can have a maximum of 100 instances. `SubnetIds` and `Filters` cannot be specified at the same time.
	SubnetIds []*string `json:"SubnetIds,omitnil,omitempty" name:"SubnetIds"`

	// Filter condition. `SubnetIds` and `Filters` cannot be specified at the same time.
	// <li>subnet-id - String - (Filter condition) Subnet instance name.</li>
	// <li>vpc-id - String - (Filter condition) VPC instance ID, such as `vpc-f49l6u0z`.</li>
	// <li>cidr-block - String - (Filter condition) Subnet IP range, such as `192.168.1.0`.</li>
	// <li>is-default - Boolean - (Filter condition) Whether it is the default subnet.</li>
	// <li>is-remote-vpc-snat - Boolean - (Filter condition) Whether it is a VPC SNAT address pool subnet.</li>
	// <li>subnet-name - String - (Filter condition) Subnet name.</li>
	// <li>zone - String - (Filter condition) Availability zone.</li>
	// <li> tag-key - String - Required: No - (Filter condition) Filter by tag key.</li>
	// <li>tag:tag-key - String - Required: No - (Filter condition) Filter by tag key-value pair. Use a specific tag key to replace `tag-key`. For its usage, see example 2.</li>
	// <li>cdc-id - String - Required: No - (Filter condition) Filter by CDC ID to obtain subnets in the specified CDC.</li>
	// <li>is-cdc-subnet - String - Required: No - (Filter condition) Whether it is a CDC subnet. Valid values: `0` (no); `1` (yes).</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Offset. Default value: 0.
	Offset *string `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *string `json:"Limit,omitnil,omitempty" name:"Limit"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSubnetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSubnetsResponseParams struct {
	// The number of instances meeting the filter condition.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Subnet object.
	SubnetSet []*Subnet `json:"SubnetSet,omitnil,omitempty" name:"SubnetSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeTaskResultRequestParams struct {
	// Async task ID. Either TaskId or DealName must be entered.
	TaskId *uint64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Billing order No. Either TaskId or DealName must be entered.
	DealName *string `json:"DealName,omitnil,omitempty" name:"DealName"`
}

type DescribeTaskResultRequest struct {
	*tchttp.BaseRequest
	
	// Async task ID. Either TaskId or DealName must be entered.
	TaskId *uint64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Billing order No. Either TaskId or DealName must be entered.
	DealName *string `json:"DealName,omitnil,omitempty" name:"DealName"`
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
	delete(f, "TaskId")
	delete(f, "DealName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTaskResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskResultResponseParams struct {
	// Job ID
	TaskId *uint64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The execution results, including `SUCCESS`, `FAILED`, and `RUNNING`
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeTrafficPackagesRequestParams struct {
	// Traffic package IDs. Multiple values can be used.
	TrafficPackageIds []*string `json:"TrafficPackageIds,omitnil,omitempty" name:"TrafficPackageIds"`

	// Each request can have up to 10 `Filters`. `TrafficPackageIds` and `Filters` cannot be specified at the same time. The specific filter conditions are as follows:
	// <li> `traffic-package_id` - String - Optional - Filter by the traffic package ID.</li>
	// <li> `traffic-package-name` - String - Optional - Filter by the traffic package name. Fuzzy match is not supported.</li>
	// <li> `status` - String - Optional - Filter by the traffic package status. Values: [AVAILABLE|EXPIRED|EXHAUSTED].</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Pagination parameter
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Pagination parameter
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeTrafficPackagesRequest struct {
	*tchttp.BaseRequest
	
	// Traffic package IDs. Multiple values can be used.
	TrafficPackageIds []*string `json:"TrafficPackageIds,omitnil,omitempty" name:"TrafficPackageIds"`

	// Each request can have up to 10 `Filters`. `TrafficPackageIds` and `Filters` cannot be specified at the same time. The specific filter conditions are as follows:
	// <li> `traffic-package_id` - String - Optional - Filter by the traffic package ID.</li>
	// <li> `traffic-package-name` - String - Optional - Filter by the traffic package name. Fuzzy match is not supported.</li>
	// <li> `status` - String - Optional - Filter by the traffic package status. Values: [AVAILABLE|EXPIRED|EXHAUSTED].</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Pagination parameter
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Pagination parameter
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeTrafficPackagesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTrafficPackagesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TrafficPackageIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTrafficPackagesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTrafficPackagesResponseParams struct {
	// Number of eligible traffic packages
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Traffic package information
	TrafficPackageSet []*TrafficPackage `json:"TrafficPackageSet,omitnil,omitempty" name:"TrafficPackageSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTrafficPackagesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTrafficPackagesResponseParams `json:"Response"`
}

func (r *DescribeTrafficPackagesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTrafficPackagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUsedIpAddressRequestParams struct {
	// VPC instance ID.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Subnet instance ID
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// List of IPs to be queried. The IPs must be within the VPC or subnet. Up to 100 IPs can be queried at a time.
	IpAddresses []*string `json:"IpAddresses,omitnil,omitempty" name:"IpAddresses"`

	// The offset. Default value: 0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The number of returned results. Default value: 20. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeUsedIpAddressRequest struct {
	*tchttp.BaseRequest
	
	// VPC instance ID.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Subnet instance ID
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// List of IPs to be queried. The IPs must be within the VPC or subnet. Up to 100 IPs can be queried at a time.
	IpAddresses []*string `json:"IpAddresses,omitnil,omitempty" name:"IpAddresses"`

	// The offset. Default value: 0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The number of returned results. Default value: 20. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeUsedIpAddressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUsedIpAddressRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "IpAddresses")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUsedIpAddressRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUsedIpAddressResponseParams struct {
	// Information of resources bound with the queried IPs 
	// Note: This parameter may return null, indicating that no valid values can be obtained.
	IpAddressStates []*IpAddressStates `json:"IpAddressStates,omitnil,omitempty" name:"IpAddressStates"`

	// Number of taken IPs 
	// Note: This parameter may return null, indicating that no valid values can be obtained.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeUsedIpAddressResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUsedIpAddressResponseParams `json:"Response"`
}

func (r *DescribeUsedIpAddressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUsedIpAddressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpcEndPointRequestParams struct {
	// Filter condition
	// <li> end-point-service-id - String - (Filter condition) Endpoint service ID.</li>
	// <li>end-point-name - String - (Filter condition) Endpoint instance name.</li>
	// <li> end-point-id - String - (Filter condition) Endpoint instance ID.</li>
	// <li> vpc-id - String - (Filter condition) VPC instance ID.</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of results per page; default value: 20; maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Endpoint ID list
	EndPointId []*string `json:"EndPointId,omitnil,omitempty" name:"EndPointId"`
}

type DescribeVpcEndPointRequest struct {
	*tchttp.BaseRequest
	
	// Filter condition
	// <li> end-point-service-id - String - (Filter condition) Endpoint service ID.</li>
	// <li>end-point-name - String - (Filter condition) Endpoint instance name.</li>
	// <li> end-point-id - String - (Filter condition) Endpoint instance ID.</li>
	// <li> vpc-id - String - (Filter condition) VPC instance ID.</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of results per page; default value: 20; maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Endpoint ID list
	EndPointId []*string `json:"EndPointId,omitnil,omitempty" name:"EndPointId"`
}

func (r *DescribeVpcEndPointRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpcEndPointRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "EndPointId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVpcEndPointRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpcEndPointResponseParams struct {
	// Endpoint
	EndPointSet []*EndPoint `json:"EndPointSet,omitnil,omitempty" name:"EndPointSet"`

	// Number of matched endpoints
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeVpcEndPointResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVpcEndPointResponseParams `json:"Response"`
}

func (r *DescribeVpcEndPointResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpcEndPointResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpcEndPointServiceRequestParams struct {
	// Filters. `EndPointServiceIds` and `Filters` cannot be both passed in. 
	// <li>`service-id` - String - Unique endpoint service ID. </li>
	// <li>`service-name` - String - Endpoint service instance name. </li>
	// <li>`service-instance-id` - String - Unique backend service ID in the format of `lb-xxx`. </li>
	// <li>`service-type` - String - Backend PaaS service type. It can be `CLB`, `CDB` or `CRS`. It defaults to `CLB` if not specified. </li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of results per page; default value: 20; maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Endpoint service ID `EndPointServiceIds` and `Filters` cannot be both passed in. 
	EndPointServiceIds []*string `json:"EndPointServiceIds,omitnil,omitempty" name:"EndPointServiceIds"`


	IsListAuthorizedEndPointService *bool `json:"IsListAuthorizedEndPointService,omitnil,omitempty" name:"IsListAuthorizedEndPointService"`
}

type DescribeVpcEndPointServiceRequest struct {
	*tchttp.BaseRequest
	
	// Filters. `EndPointServiceIds` and `Filters` cannot be both passed in. 
	// <li>`service-id` - String - Unique endpoint service ID. </li>
	// <li>`service-name` - String - Endpoint service instance name. </li>
	// <li>`service-instance-id` - String - Unique backend service ID in the format of `lb-xxx`. </li>
	// <li>`service-type` - String - Backend PaaS service type. It can be `CLB`, `CDB` or `CRS`. It defaults to `CLB` if not specified. </li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of results per page; default value: 20; maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Endpoint service ID `EndPointServiceIds` and `Filters` cannot be both passed in. 
	EndPointServiceIds []*string `json:"EndPointServiceIds,omitnil,omitempty" name:"EndPointServiceIds"`

	IsListAuthorizedEndPointService *bool `json:"IsListAuthorizedEndPointService,omitnil,omitempty" name:"IsListAuthorizedEndPointService"`
}

func (r *DescribeVpcEndPointServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpcEndPointServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "EndPointServiceIds")
	delete(f, "IsListAuthorizedEndPointService")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVpcEndPointServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpcEndPointServiceResponseParams struct {
	// Array of endpoint services
	EndPointServiceSet []*EndPointService `json:"EndPointServiceSet,omitnil,omitempty" name:"EndPointServiceSet"`

	// Number of matched results
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeVpcEndPointServiceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVpcEndPointServiceResponseParams `json:"Response"`
}

func (r *DescribeVpcEndPointServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpcEndPointServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpcEndPointServiceWhiteListRequestParams struct {
	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of results per page; default value: 20; maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Filter condition
	// <li> user-uin - String - (Filter condition) UIN.</li>
	// <li> end-point-service-id - String - (Filter condition) Endpoint service ID.</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeVpcEndPointServiceWhiteListRequest struct {
	*tchttp.BaseRequest
	
	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of results per page; default value: 20; maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Filter condition
	// <li> user-uin - String - (Filter condition) UIN.</li>
	// <li> end-point-service-id - String - (Filter condition) Endpoint service ID.</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeVpcEndPointServiceWhiteListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpcEndPointServiceWhiteListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVpcEndPointServiceWhiteListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpcEndPointServiceWhiteListResponseParams struct {
	// Array of allowed endpoint services
	VpcEndpointServiceUserSet []*VpcEndPointServiceUser `json:"VpcEndpointServiceUserSet,omitnil,omitempty" name:"VpcEndpointServiceUserSet"`

	// Number of matched allowlists
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeVpcEndPointServiceWhiteListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVpcEndPointServiceWhiteListResponseParams `json:"Response"`
}

func (r *DescribeVpcEndPointServiceWhiteListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpcEndPointServiceWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpcInstancesRequestParams struct {
	// Filter condition. `RouteTableIds` and `Filters` cannot be specified at the same time.
	// <li>vpc-id - String - (Filter condition) VPC instance ID, such as `vpc-f49l6u0z`.</li>
	// <li>instance-type - String - (Filter condition) CVM instance ID.</li>
	// <li>instance-name - String - (Filter condition) CVM name.</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Offset.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The number of requested objects.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeVpcInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Filter condition. `RouteTableIds` and `Filters` cannot be specified at the same time.
	// <li>vpc-id - String - (Filter condition) VPC instance ID, such as `vpc-f49l6u0z`.</li>
	// <li>instance-type - String - (Filter condition) CVM instance ID.</li>
	// <li>instance-name - String - (Filter condition) CVM name.</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Offset.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The number of requested objects.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeVpcInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpcInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVpcInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpcInstancesResponseParams struct {
	// List of CVM instances.
	InstanceSet []*CvmInstance `json:"InstanceSet,omitnil,omitempty" name:"InstanceSet"`

	// The number of eligible CVM instances.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeVpcInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVpcInstancesResponseParams `json:"Response"`
}

func (r *DescribeVpcInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpcInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpcIpv6AddressesRequestParams struct {
	// The `ID` of the `VPC`, such as `vpc-f49l6u0z`.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// The `IP` address list. Each request supports a maximum of `10` batch querying.
	Ipv6Addresses []*string `json:"Ipv6Addresses,omitnil,omitempty" name:"Ipv6Addresses"`

	// Offset.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The returned quantity.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeVpcIpv6AddressesRequest struct {
	*tchttp.BaseRequest
	
	// The `ID` of the `VPC`, such as `vpc-f49l6u0z`.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// The `IP` address list. Each request supports a maximum of `10` batch querying.
	Ipv6Addresses []*string `json:"Ipv6Addresses,omitnil,omitempty" name:"Ipv6Addresses"`

	// Offset.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The returned quantity.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeVpcIpv6AddressesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpcIpv6AddressesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcId")
	delete(f, "Ipv6Addresses")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVpcIpv6AddressesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpcIpv6AddressesResponseParams struct {
	// The `IPv6` address list.
	Ipv6AddressSet []*VpcIpv6Address `json:"Ipv6AddressSet,omitnil,omitempty" name:"Ipv6AddressSet"`

	// The total number of `IPv6` addresses.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeVpcIpv6AddressesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVpcIpv6AddressesResponseParams `json:"Response"`
}

func (r *DescribeVpcIpv6AddressesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpcIpv6AddressesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpcPrivateIpAddressesRequestParams struct {
	// The `ID` of the `VPC`, such as `vpc-f49l6u0z`.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// The private `IP` address list. Each request supports a maximum of `10` batch querying.
	PrivateIpAddresses []*string `json:"PrivateIpAddresses,omitnil,omitempty" name:"PrivateIpAddresses"`
}

type DescribeVpcPrivateIpAddressesRequest struct {
	*tchttp.BaseRequest
	
	// The `ID` of the `VPC`, such as `vpc-f49l6u0z`.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// The private `IP` address list. Each request supports a maximum of `10` batch querying.
	PrivateIpAddresses []*string `json:"PrivateIpAddresses,omitnil,omitempty" name:"PrivateIpAddresses"`
}

func (r *DescribeVpcPrivateIpAddressesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpcPrivateIpAddressesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcId")
	delete(f, "PrivateIpAddresses")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVpcPrivateIpAddressesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpcPrivateIpAddressesResponseParams struct {
	// The list of private `IP` address information.
	VpcPrivateIpAddressSet []*VpcPrivateIpAddress `json:"VpcPrivateIpAddressSet,omitnil,omitempty" name:"VpcPrivateIpAddressSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeVpcPrivateIpAddressesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVpcPrivateIpAddressesResponseParams `json:"Response"`
}

func (r *DescribeVpcPrivateIpAddressesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpcPrivateIpAddressesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpcResourceDashboardRequestParams struct {
	// Vpc instance ID, e.g. vpc-f1xjkw1b.
	VpcIds []*string `json:"VpcIds,omitnil,omitempty" name:"VpcIds"`
}

type DescribeVpcResourceDashboardRequest struct {
	*tchttp.BaseRequest
	
	// Vpc instance ID, e.g. vpc-f1xjkw1b.
	VpcIds []*string `json:"VpcIds,omitnil,omitempty" name:"VpcIds"`
}

func (r *DescribeVpcResourceDashboardRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpcResourceDashboardRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVpcResourceDashboardRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpcResourceDashboardResponseParams struct {
	// List of resource objects.
	ResourceDashboardSet []*ResourceDashboard `json:"ResourceDashboardSet,omitnil,omitempty" name:"ResourceDashboardSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeVpcResourceDashboardResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVpcResourceDashboardResponseParams `json:"Response"`
}

func (r *DescribeVpcResourceDashboardResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpcResourceDashboardResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpcTaskResultRequestParams struct {
	// `RequestId` returned by an async task
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type DescribeVpcTaskResultRequest struct {
	*tchttp.BaseRequest
	
	// `RequestId` returned by an async task
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *DescribeVpcTaskResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpcTaskResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVpcTaskResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpcTaskResultResponseParams struct {
	// The execution results of an async task. Valid values: `SUCCESS`(task executed successfully), `FAILED` (task execution failed), and `RUNNING` (task in progress). 
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Output of the async task execution result
	Output *string `json:"Output,omitnil,omitempty" name:"Output"`

	// Detailed result of an async task, such as the result of batch deleting ENIs.
	// Note: this field may return `null`, indicating that no valid value can be obtained.
	Result []*VpcTaskResultDetailInfo `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeVpcTaskResultResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVpcTaskResultResponseParams `json:"Response"`
}

func (r *DescribeVpcTaskResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpcTaskResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpcsRequestParams struct {
	// The VPC instance ID, such as `vpc-f49l6u0z`. Each request supports a maximum of 100 instances. `VpcIds` and `Filters` cannot be specified at the same time.
	VpcIds []*string `json:"VpcIds,omitnil,omitempty" name:"VpcIds"`

	// Filter condition. `VpcIds` and `Filters` cannot be specified at the same time.
	// Valid filters include:
	// <li>`vpc-name`: VPC instance name, which supports fuzzy search.</li>
	// <li>`is-default`: Indicates whether it is the default VPC</li>
	// <li>`vpc-id`: VPC instance ID, such as `vpc-f49l6u0z`</li>
	// <li>`cidr-block`: VPC CIDR block</li>
	// <li>`tag-key`: (Optional) tag key</li>
	// <li>`tag:tag-key`: (Optional) tag key-value pair. Replace the `tag-key` with a specified tag value. For its usage, refer to Example 2.</li>
	//   **Note:** If one filter has multiple values, the logical relationship between these values is `OR`. The logical relationship between filters is `AND`.
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Offset. Default value: 0.
	Offset *string `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *string `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeVpcsRequest struct {
	*tchttp.BaseRequest
	
	// The VPC instance ID, such as `vpc-f49l6u0z`. Each request supports a maximum of 100 instances. `VpcIds` and `Filters` cannot be specified at the same time.
	VpcIds []*string `json:"VpcIds,omitnil,omitempty" name:"VpcIds"`

	// Filter condition. `VpcIds` and `Filters` cannot be specified at the same time.
	// Valid filters include:
	// <li>`vpc-name`: VPC instance name, which supports fuzzy search.</li>
	// <li>`is-default`: Indicates whether it is the default VPC</li>
	// <li>`vpc-id`: VPC instance ID, such as `vpc-f49l6u0z`</li>
	// <li>`cidr-block`: VPC CIDR block</li>
	// <li>`tag-key`: (Optional) tag key</li>
	// <li>`tag:tag-key`: (Optional) tag key-value pair. Replace the `tag-key` with a specified tag value. For its usage, refer to Example 2.</li>
	//   **Note:** If one filter has multiple values, the logical relationship between these values is `OR`. The logical relationship between filters is `AND`.
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Offset. Default value: 0.
	Offset *string `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *string `json:"Limit,omitnil,omitempty" name:"Limit"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVpcsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpcsResponseParams struct {
	// The number of objects meeting the condition.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The VPC object.
	VpcSet []*Vpc `json:"VpcSet,omitnil,omitempty" name:"VpcSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeVpnConnectionsRequestParams struct {
	// The instance ID of the VPN tunnel, such as `vpnx-f49l6u0z`. Each request can have a maximum of 100 instances. `VpnConnectionIds` and `Filters` cannot be specified at the same time.
	VpnConnectionIds []*string `json:"VpnConnectionIds,omitnil,omitempty" name:"VpnConnectionIds"`

	// Filter condition. In each request, the upper limit for `Filters` is 10 and 5 for `Filter.Values`. `VpnConnectionIds` and `Filters` cannot be specified at the same time.
	// <li>vpc-id - String - VPC instance ID, such as `vpc-0a36uwkr`.</li>
	// <li>vpn-gateway-id - String - VPN gateway instance ID, such as `vpngw-p4lmqawn`.</li>
	// <li>customer-gateway-id - String - Customer gateway instance ID, such as `cgw-l4rblw63`.</li>
	// <li>vpn-connection-name - String - Connection name, such as `test-vpn`.</li>
	// <li>vpn-connection-id - String - Connection instance ID, such as `vpnx-5p7vkch8"`.</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// The Offset. The default value is 0. For more information about Offset, see the relevant section in the API Introduction.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of values to be returned. The default value is 20. Maximum is 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeVpnConnectionsRequest struct {
	*tchttp.BaseRequest
	
	// The instance ID of the VPN tunnel, such as `vpnx-f49l6u0z`. Each request can have a maximum of 100 instances. `VpnConnectionIds` and `Filters` cannot be specified at the same time.
	VpnConnectionIds []*string `json:"VpnConnectionIds,omitnil,omitempty" name:"VpnConnectionIds"`

	// Filter condition. In each request, the upper limit for `Filters` is 10 and 5 for `Filter.Values`. `VpnConnectionIds` and `Filters` cannot be specified at the same time.
	// <li>vpc-id - String - VPC instance ID, such as `vpc-0a36uwkr`.</li>
	// <li>vpn-gateway-id - String - VPN gateway instance ID, such as `vpngw-p4lmqawn`.</li>
	// <li>customer-gateway-id - String - Customer gateway instance ID, such as `cgw-l4rblw63`.</li>
	// <li>vpn-connection-name - String - Connection name, such as `test-vpn`.</li>
	// <li>vpn-connection-id - String - Connection instance ID, such as `vpnx-5p7vkch8"`.</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// The Offset. The default value is 0. For more information about Offset, see the relevant section in the API Introduction.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of values to be returned. The default value is 20. Maximum is 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeVpnConnectionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpnConnectionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpnConnectionIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVpnConnectionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpnConnectionsResponseParams struct {
	// The number of instances meeting the filter condition.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// VPN tunnel instance.
	VpnConnectionSet []*VpnConnection `json:"VpnConnectionSet,omitnil,omitempty" name:"VpnConnectionSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeVpnConnectionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVpnConnectionsResponseParams `json:"Response"`
}

func (r *DescribeVpnConnectionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpnConnectionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpnGatewayCcnRoutesRequestParams struct {
	// Instance ID of the VPN gateway
	VpnGatewayId *string `json:"VpnGatewayId,omitnil,omitempty" name:"VpnGatewayId"`

	// Offset.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The returned quantity
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeVpnGatewayCcnRoutesRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID of the VPN gateway
	VpnGatewayId *string `json:"VpnGatewayId,omitnil,omitempty" name:"VpnGatewayId"`

	// Offset.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The returned quantity
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeVpnGatewayCcnRoutesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpnGatewayCcnRoutesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpnGatewayId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVpnGatewayCcnRoutesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpnGatewayCcnRoutesResponseParams struct {
	// The CCN route (IDC IP range) list.
	RouteSet []*VpngwCcnRoutes `json:"RouteSet,omitnil,omitempty" name:"RouteSet"`

	// Number of objects that meet the condition.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeVpnGatewayCcnRoutesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVpnGatewayCcnRoutesResponseParams `json:"Response"`
}

func (r *DescribeVpnGatewayCcnRoutesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpnGatewayCcnRoutesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpnGatewayRoutesRequestParams struct {
	// Instance ID of the VPN gateway
	VpnGatewayId *string `json:"VpnGatewayId,omitnil,omitempty" name:"VpnGatewayId"`

	// Filter condition. Valid values: `DestinationCidr`, `InstanceId`, and `InstanceType`.
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Offset. Default value: 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned results per page. Default value: 20; maximum value: 100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeVpnGatewayRoutesRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID of the VPN gateway
	VpnGatewayId *string `json:"VpnGatewayId,omitnil,omitempty" name:"VpnGatewayId"`

	// Filter condition. Valid values: `DestinationCidr`, `InstanceId`, and `InstanceType`.
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Offset. Default value: 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned results per page. Default value: 20; maximum value: 100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeVpnGatewayRoutesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpnGatewayRoutesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpnGatewayId")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVpnGatewayRoutesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpnGatewayRoutesResponseParams struct {
	// Destination routes of the VPN gateway
	Routes []*VpnGatewayRoute `json:"Routes,omitnil,omitempty" name:"Routes"`

	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeVpnGatewayRoutesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVpnGatewayRoutesResponseParams `json:"Response"`
}

func (r *DescribeVpnGatewayRoutesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpnGatewayRoutesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpnGatewaysRequestParams struct {
	// The VPN gateway instance ID, such as `vpngw-f49l6u0z`. Each request can have a maximum of 100 instances. `VpnGatewayIds` and `Filters` cannot be specified at the same time.
	VpnGatewayIds []*string `json:"VpnGatewayIds,omitnil,omitempty" name:"VpnGatewayIds"`

	// Filter condition. `VpnGatewayIds` and `Filters` cannot be specified at the same time.
	// <li>vpc-id - String - (Filter condition) VPC instance ID, such as `vpc-f49l6u0z`.</li>
	// <li>vpn-gateway-id - String - (Filter condition) VPN instance ID, such as `vpngw-5aluhh9t`.</li>
	// <li>vpn-gateway-name - String - (Filter condition) VPN instance name.</li>
	// <li>type - String - (Filter condition) VPN gateway type: 'IPSEC', 'SSL'.</li>
	// <li>public-ip-address- String - (Filter condition) Public IP.</li>
	// <li>renew-flag - String - (Filter condition) Gateway renewal type. Manual renewal: `NOTIFY_AND_MANUAL_RENEW`, Automatic renewal: `NOTIFY_AND_AUTO_RENEW`.</li>
	// <li>zone - String - (Filter condition) The availability zone where the VPN is located, such as `ap-guangzhou-2`.</li>
	Filters []*FilterObject `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Offset
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The number of request objects.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeVpnGatewaysRequest struct {
	*tchttp.BaseRequest
	
	// The VPN gateway instance ID, such as `vpngw-f49l6u0z`. Each request can have a maximum of 100 instances. `VpnGatewayIds` and `Filters` cannot be specified at the same time.
	VpnGatewayIds []*string `json:"VpnGatewayIds,omitnil,omitempty" name:"VpnGatewayIds"`

	// Filter condition. `VpnGatewayIds` and `Filters` cannot be specified at the same time.
	// <li>vpc-id - String - (Filter condition) VPC instance ID, such as `vpc-f49l6u0z`.</li>
	// <li>vpn-gateway-id - String - (Filter condition) VPN instance ID, such as `vpngw-5aluhh9t`.</li>
	// <li>vpn-gateway-name - String - (Filter condition) VPN instance name.</li>
	// <li>type - String - (Filter condition) VPN gateway type: 'IPSEC', 'SSL'.</li>
	// <li>public-ip-address- String - (Filter condition) Public IP.</li>
	// <li>renew-flag - String - (Filter condition) Gateway renewal type. Manual renewal: `NOTIFY_AND_MANUAL_RENEW`, Automatic renewal: `NOTIFY_AND_AUTO_RENEW`.</li>
	// <li>zone - String - (Filter condition) The availability zone where the VPN is located, such as `ap-guangzhou-2`.</li>
	Filters []*FilterObject `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Offset
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The number of request objects.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeVpnGatewaysRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpnGatewaysRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpnGatewayIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVpnGatewaysRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpnGatewaysResponseParams struct {
	// The number of instances meeting the filter condition.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The list of details of VPN gateway instances.
	VpnGatewaySet []*VpnGateway `json:"VpnGatewaySet,omitnil,omitempty" name:"VpnGatewaySet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeVpnGatewaysResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVpnGatewaysResponseParams `json:"Response"`
}

func (r *DescribeVpnGatewaysResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpnGatewaysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DestinationIpPortTranslationNatRule struct {
	// Network protocol. Valid values: `TCP`, `UDP`.
	IpProtocol *string `json:"IpProtocol,omitnil,omitempty" name:"IpProtocol"`

	// EIP.
	PublicIpAddress *string `json:"PublicIpAddress,omitnil,omitempty" name:"PublicIpAddress"`

	// Public network port.
	PublicPort *uint64 `json:"PublicPort,omitnil,omitempty" name:"PublicPort"`

	// Private network address.
	PrivateIpAddress *string `json:"PrivateIpAddress,omitnil,omitempty" name:"PrivateIpAddress"`

	// Private network port.
	PrivatePort *uint64 `json:"PrivatePort,omitnil,omitempty" name:"PrivatePort"`

	// Description of NAT gateway forwarding rules.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

// Predefined struct for user
type DetachCcnInstancesRequestParams struct {
	// The CCN instance ID, such as `ccn-f49l6u0z`.
	CcnId *string `json:"CcnId,omitnil,omitempty" name:"CcnId"`

	// The list of network instances to be unbound
	Instances []*CcnInstance `json:"Instances,omitnil,omitempty" name:"Instances"`
}

type DetachCcnInstancesRequest struct {
	*tchttp.BaseRequest
	
	// The CCN instance ID, such as `ccn-f49l6u0z`.
	CcnId *string `json:"CcnId,omitnil,omitempty" name:"CcnId"`

	// The list of network instances to be unbound
	Instances []*CcnInstance `json:"Instances,omitnil,omitempty" name:"Instances"`
}

func (r *DetachCcnInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetachCcnInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CcnId")
	delete(f, "Instances")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DetachCcnInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetachCcnInstancesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DetachCcnInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DetachCcnInstancesResponseParams `json:"Response"`
}

func (r *DetachCcnInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetachCcnInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetachClassicLinkVpcRequestParams struct {
	// The ID of the VPC instance. You can obtain the parameter value from the VpcId field in the returned result of DescribeVpcs API.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Queries the ID of the CVM instance, such as `ins-r8hr2upy`.
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

type DetachClassicLinkVpcRequest struct {
	*tchttp.BaseRequest
	
	// The ID of the VPC instance. You can obtain the parameter value from the VpcId field in the returned result of DescribeVpcs API.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Queries the ID of the CVM instance, such as `ins-r8hr2upy`.
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

func (r *DetachClassicLinkVpcRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetachClassicLinkVpcRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcId")
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DetachClassicLinkVpcRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetachClassicLinkVpcResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DetachClassicLinkVpcResponse struct {
	*tchttp.BaseResponse
	Response *DetachClassicLinkVpcResponseParams `json:"Response"`
}

func (r *DetachClassicLinkVpcResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetachClassicLinkVpcResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetachNetworkInterfaceRequestParams struct {
	// The ID of the ENI instance, such as `eni-m6dyj72l`.
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitnil,omitempty" name:"NetworkInterfaceId"`

	// The ID of the CVM instance, such as `ins-r8hr2upy`.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DetachNetworkInterfaceRequest struct {
	*tchttp.BaseRequest
	
	// The ID of the ENI instance, such as `eni-m6dyj72l`.
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitnil,omitempty" name:"NetworkInterfaceId"`

	// The ID of the CVM instance, such as `ins-r8hr2upy`.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DetachNetworkInterfaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetachNetworkInterfaceResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DetachSnapshotInstancesRequestParams struct {
	// Snapshot policy IDs
	SnapshotPolicyId *string `json:"SnapshotPolicyId,omitnil,omitempty" name:"SnapshotPolicyId"`

	// Information of instances
	Instances []*SnapshotInstance `json:"Instances,omitnil,omitempty" name:"Instances"`
}

type DetachSnapshotInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Snapshot policy IDs
	SnapshotPolicyId *string `json:"SnapshotPolicyId,omitnil,omitempty" name:"SnapshotPolicyId"`

	// Information of instances
	Instances []*SnapshotInstance `json:"Instances,omitnil,omitempty" name:"Instances"`
}

func (r *DetachSnapshotInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetachSnapshotInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SnapshotPolicyId")
	delete(f, "Instances")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DetachSnapshotInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetachSnapshotInstancesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DetachSnapshotInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DetachSnapshotInstancesResponseParams `json:"Response"`
}

func (r *DetachSnapshotInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetachSnapshotInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DirectConnectGateway struct {
	// Direct Connect `ID`.
	DirectConnectGatewayId *string `json:"DirectConnectGatewayId,omitnil,omitempty" name:"DirectConnectGatewayId"`

	// Direct Connect gateway name.
	DirectConnectGatewayName *string `json:"DirectConnectGatewayName,omitnil,omitempty" name:"DirectConnectGatewayName"`

	// The `ID` of the `VPC` instance associated with the Direct Connect gateway.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// The associated network type:
	// <li>`VPC` - VPC</li>
	// <li>`CCN` - CCN</li>
	NetworkType *string `json:"NetworkType,omitnil,omitempty" name:"NetworkType"`

	// The `ID` of the associated network instance:
	// <li>When the NetworkType is `VPC`, this value is the VPC instance `ID`</li>
	// <li>When the NetworkType is `CCN`, this value is the CCN instance `ID`</li>
	NetworkInstanceId *string `json:"NetworkInstanceId,omitnil,omitempty" name:"NetworkInstanceId"`

	// Gateway type:
	// <li>NORMAL - Standard type. Note: CCN only supports the standard type</li>
	// <li>NAT - NAT type</li>
	// NAT type supports network address switch configuration. After the type is confirmed, it cannot be modified. A VPC can create one NAT-type Direct Connect gateway and one non-NAT-type Direct Connect gateway
	GatewayType *string `json:"GatewayType,omitnil,omitempty" name:"GatewayType"`

	// Creation Time.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Direct Connect gateway IP.
	DirectConnectGatewayIp *string `json:"DirectConnectGatewayIp,omitnil,omitempty" name:"DirectConnectGatewayIp"`

	// The `ID` of the `CCN` instance associated with the Direct Connect gateway.
	CcnId *string `json:"CcnId,omitnil,omitempty" name:"CcnId"`

	// The route-learning type of the CCN:
	// <li>`BGP` - Automatic learning.</li>
	// <li>`STATIC` - Static, that is, user-configured.</li>
	CcnRouteType *string `json:"CcnRouteType,omitnil,omitempty" name:"CcnRouteType"`

	// Whether BGP is enabled.
	EnableBGP *bool `json:"EnableBGP,omitnil,omitempty" name:"EnableBGP"`

	// Whether to enable BGP's `community` attribute. Valid values: enable, disable
	EnableBGPCommunity *bool `json:"EnableBGPCommunity,omitnil,omitempty" name:"EnableBGPCommunity"`

	// ID of the NAT gateway bound.
	// Note: this field may return `null`, indicating that no valid value was found.
	NatGatewayId *string `json:"NatGatewayId,omitnil,omitempty" name:"NatGatewayId"`

	// Whether the direct connect gateway supports the VXLAN architecture.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	VXLANSupport []*bool `json:"VXLANSupport,omitnil,omitempty" name:"VXLANSupport"`

	// CCN route publishing mode. Valid values: `standard` and `exquisite`.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ModeType *string `json:"ModeType,omitnil,omitempty" name:"ModeType"`

	// Whether the direct connect gateway is for an edge zone.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	LocalZone *bool `json:"LocalZone,omitnil,omitempty" name:"LocalZone"`

	// Availability zone where the direct connect gateway resides.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// The status of gateway traffic monitoring
	// 0: disable
	// 1: enable
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	EnableFlowDetails *uint64 `json:"EnableFlowDetails,omitnil,omitempty" name:"EnableFlowDetails"`

	// The last time when the gateway traffic monitoring is enabled/disabled
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	FlowDetailsUpdateTime *string `json:"FlowDetailsUpdateTime,omitnil,omitempty" name:"FlowDetailsUpdateTime"`

	// Whether gateway traffic monitoring is supported
	// 0: No
	// 1: Yes
	// Note: this field may return `null`, indicating that no valid values can be found.
	NewAfc *uint64 `json:"NewAfc,omitnil,omitempty" name:"NewAfc"`

	// Direct connect gateway access network types:
	// <li>`VXLAN` - VXLAN type.</li>
	// <li>`MPLS` - MPLS type.</li>
	// <li>`Hybrid` - Hybrid type.</li>
	// Note: this field may return `null`, indicating that no valid values can be found.
	AccessNetworkType *string `json:"AccessNetworkType,omitnil,omitempty" name:"AccessNetworkType"`

	// AZ list of direct connect gateway with cross-AZ placement groups
	// Note: this field may return `null`, indicating that no valid values can be found.
	HaZoneList []*string `json:"HaZoneList,omitnil,omitempty" name:"HaZoneList"`
}

type DirectConnectGatewayCcnRoute struct {
	// Route ID.
	RouteId *string `json:"RouteId,omitnil,omitempty" name:"RouteId"`

	// IDC IP range.
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitnil,omitempty" name:"DestinationCidrBlock"`

	// The `AS-Path` attribute of `BGP`.
	ASPath []*string `json:"ASPath,omitnil,omitempty" name:"ASPath"`

	// Remarks
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Last updated time
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type DirectConnectSubnet struct {
	// The direct connect gateway ID.
	DirectConnectGatewayId *string `json:"DirectConnectGatewayId,omitnil,omitempty" name:"DirectConnectGatewayId"`

	// IDC subnet IP range
	CidrBlock *string `json:"CidrBlock,omitnil,omitempty" name:"CidrBlock"`
}

// Predefined struct for user
type DisableCcnRoutesRequestParams struct {
	// The CCN instance ID, such as `ccn-f49l6u0z`.
	CcnId *string `json:"CcnId,omitnil,omitempty" name:"CcnId"`

	// The unique ID of the CCN routing policy, such as `ccnr-f49l6u0z`.
	RouteIds []*string `json:"RouteIds,omitnil,omitempty" name:"RouteIds"`
}

type DisableCcnRoutesRequest struct {
	*tchttp.BaseRequest
	
	// The CCN instance ID, such as `ccn-f49l6u0z`.
	CcnId *string `json:"CcnId,omitnil,omitempty" name:"CcnId"`

	// The unique ID of the CCN routing policy, such as `ccnr-f49l6u0z`.
	RouteIds []*string `json:"RouteIds,omitnil,omitempty" name:"RouteIds"`
}

func (r *DisableCcnRoutesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableCcnRoutesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CcnId")
	delete(f, "RouteIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisableCcnRoutesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableCcnRoutesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DisableCcnRoutesResponse struct {
	*tchttp.BaseResponse
	Response *DisableCcnRoutesResponseParams `json:"Response"`
}

func (r *DisableCcnRoutesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableCcnRoutesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableFlowLogsRequestParams struct {
	// Flow log ID.
	FlowLogIds []*string `json:"FlowLogIds,omitnil,omitempty" name:"FlowLogIds"`
}

type DisableFlowLogsRequest struct {
	*tchttp.BaseRequest
	
	// Flow log ID.
	FlowLogIds []*string `json:"FlowLogIds,omitnil,omitempty" name:"FlowLogIds"`
}

func (r *DisableFlowLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableFlowLogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FlowLogIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisableFlowLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableFlowLogsResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DisableFlowLogsResponse struct {
	*tchttp.BaseResponse
	Response *DisableFlowLogsResponseParams `json:"Response"`
}

func (r *DisableFlowLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableFlowLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableGatewayFlowMonitorRequestParams struct {
	// Gateway instance ID. Supported types:
	// Direct connect gateway instance, such as `dcg-ltjahce6`;
	// NAT gateway instance, such as `nat-ltjahce6`;
	// VPN gateway instance, such as `vpn-ltjahce6`.
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`
}

type DisableGatewayFlowMonitorRequest struct {
	*tchttp.BaseRequest
	
	// Gateway instance ID. Supported types:
	// Direct connect gateway instance, such as `dcg-ltjahce6`;
	// NAT gateway instance, such as `nat-ltjahce6`;
	// VPN gateway instance, such as `vpn-ltjahce6`.
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`
}

func (r *DisableGatewayFlowMonitorRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableGatewayFlowMonitorRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisableGatewayFlowMonitorRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableGatewayFlowMonitorResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DisableGatewayFlowMonitorResponse struct {
	*tchttp.BaseResponse
	Response *DisableGatewayFlowMonitorResponseParams `json:"Response"`
}

func (r *DisableGatewayFlowMonitorResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableGatewayFlowMonitorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableSnapshotPoliciesRequestParams struct {
	// Snapshot policy IDs
	SnapshotPolicyIds []*string `json:"SnapshotPolicyIds,omitnil,omitempty" name:"SnapshotPolicyIds"`
}

type DisableSnapshotPoliciesRequest struct {
	*tchttp.BaseRequest
	
	// Snapshot policy IDs
	SnapshotPolicyIds []*string `json:"SnapshotPolicyIds,omitnil,omitempty" name:"SnapshotPolicyIds"`
}

func (r *DisableSnapshotPoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableSnapshotPoliciesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SnapshotPolicyIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisableSnapshotPoliciesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableSnapshotPoliciesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DisableSnapshotPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *DisableSnapshotPoliciesResponseParams `json:"Response"`
}

func (r *DisableSnapshotPoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableSnapshotPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisassociateAddressRequestParams struct {
	// The unique ID of the EIP, such as `eip-11112222`.
	AddressId *string `json:"AddressId,omitnil,omitempty" name:"AddressId"`

	// Whether a common public IP is assigned after the EIP is unbound. Value range:<br><li>TRUE: Indicates that after the EIP is unbound, a common public IP is assigned.<br><li>FALSE: Indicates that after the EIP is unbound, a common public IP is not assigned.<br>Default value: FALSE.<br><br>The parameter can be specified only under the following conditions:<br><li>It can only be specified when you unbind an EIP from the primary private IP of the primary ENI.<br><li>After an EIP is unbound, you can assign public IPs to an account up to 10 times per day. For more information, use the [DescribeAddressQuota] (https://intl.cloud.tencent.com/document/api/213/1378?from_cn_redirect=1) API.
	ReallocateNormalPublicIp *bool `json:"ReallocateNormalPublicIp,omitnil,omitempty" name:"ReallocateNormalPublicIp"`
}

type DisassociateAddressRequest struct {
	*tchttp.BaseRequest
	
	// The unique ID of the EIP, such as `eip-11112222`.
	AddressId *string `json:"AddressId,omitnil,omitempty" name:"AddressId"`

	// Whether a common public IP is assigned after the EIP is unbound. Value range:<br><li>TRUE: Indicates that after the EIP is unbound, a common public IP is assigned.<br><li>FALSE: Indicates that after the EIP is unbound, a common public IP is not assigned.<br>Default value: FALSE.<br><br>The parameter can be specified only under the following conditions:<br><li>It can only be specified when you unbind an EIP from the primary private IP of the primary ENI.<br><li>After an EIP is unbound, you can assign public IPs to an account up to 10 times per day. For more information, use the [DescribeAddressQuota] (https://intl.cloud.tencent.com/document/api/213/1378?from_cn_redirect=1) API.
	ReallocateNormalPublicIp *bool `json:"ReallocateNormalPublicIp,omitnil,omitempty" name:"ReallocateNormalPublicIp"`
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
	delete(f, "AddressId")
	delete(f, "ReallocateNormalPublicIp")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisassociateAddressRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisassociateAddressResponseParams struct {
	// The async task ID. You can use the [DescribeTaskResult](https://intl.cloud.tencent.com/document/api/215/36271?from_cn_redirect=1) API to query the task status.
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DisassociateDirectConnectGatewayNatGatewayRequestParams struct {
	// The direct connect gateway ID.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// The NAT Gateway ID.
	NatGatewayId *string `json:"NatGatewayId,omitnil,omitempty" name:"NatGatewayId"`

	// The ID of the VPC instance, which can be obtained from the `VpcId` field in response of the `DescribeVpcs` API.
	DirectConnectGatewayId *string `json:"DirectConnectGatewayId,omitnil,omitempty" name:"DirectConnectGatewayId"`
}

type DisassociateDirectConnectGatewayNatGatewayRequest struct {
	*tchttp.BaseRequest
	
	// The direct connect gateway ID.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// The NAT Gateway ID.
	NatGatewayId *string `json:"NatGatewayId,omitnil,omitempty" name:"NatGatewayId"`

	// The ID of the VPC instance, which can be obtained from the `VpcId` field in response of the `DescribeVpcs` API.
	DirectConnectGatewayId *string `json:"DirectConnectGatewayId,omitnil,omitempty" name:"DirectConnectGatewayId"`
}

func (r *DisassociateDirectConnectGatewayNatGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisassociateDirectConnectGatewayNatGatewayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcId")
	delete(f, "NatGatewayId")
	delete(f, "DirectConnectGatewayId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisassociateDirectConnectGatewayNatGatewayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisassociateDirectConnectGatewayNatGatewayResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DisassociateDirectConnectGatewayNatGatewayResponse struct {
	*tchttp.BaseResponse
	Response *DisassociateDirectConnectGatewayNatGatewayResponseParams `json:"Response"`
}

func (r *DisassociateDirectConnectGatewayNatGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisassociateDirectConnectGatewayNatGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisassociateNatGatewayAddressRequestParams struct {
	// The ID of the NAT gateway, such as `nat-df45454`.
	NatGatewayId *string `json:"NatGatewayId,omitnil,omitempty" name:"NatGatewayId"`

	// Array of the EIPs to be unbound from the NAT gateway.
	PublicIpAddresses []*string `json:"PublicIpAddresses,omitnil,omitempty" name:"PublicIpAddresses"`
}

type DisassociateNatGatewayAddressRequest struct {
	*tchttp.BaseRequest
	
	// The ID of the NAT gateway, such as `nat-df45454`.
	NatGatewayId *string `json:"NatGatewayId,omitnil,omitempty" name:"NatGatewayId"`

	// Array of the EIPs to be unbound from the NAT gateway.
	PublicIpAddresses []*string `json:"PublicIpAddresses,omitnil,omitempty" name:"PublicIpAddresses"`
}

func (r *DisassociateNatGatewayAddressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisassociateNatGatewayAddressRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NatGatewayId")
	delete(f, "PublicIpAddresses")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisassociateNatGatewayAddressRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisassociateNatGatewayAddressResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DisassociateNatGatewayAddressResponse struct {
	*tchttp.BaseResponse
	Response *DisassociateNatGatewayAddressResponseParams `json:"Response"`
}

func (r *DisassociateNatGatewayAddressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisassociateNatGatewayAddressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisassociateNetworkAclSubnetsRequestParams struct {
	// Network ACL instance ID, such as `acl-12345678`.
	NetworkAclId *string `json:"NetworkAclId,omitnil,omitempty" name:"NetworkAclId"`

	// Array of subnet instance IDs, such as [subnet-12345678].
	SubnetIds []*string `json:"SubnetIds,omitnil,omitempty" name:"SubnetIds"`
}

type DisassociateNetworkAclSubnetsRequest struct {
	*tchttp.BaseRequest
	
	// Network ACL instance ID, such as `acl-12345678`.
	NetworkAclId *string `json:"NetworkAclId,omitnil,omitempty" name:"NetworkAclId"`

	// Array of subnet instance IDs, such as [subnet-12345678].
	SubnetIds []*string `json:"SubnetIds,omitnil,omitempty" name:"SubnetIds"`
}

func (r *DisassociateNetworkAclSubnetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisassociateNetworkAclSubnetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NetworkAclId")
	delete(f, "SubnetIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisassociateNetworkAclSubnetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisassociateNetworkAclSubnetsResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DisassociateNetworkAclSubnetsResponse struct {
	*tchttp.BaseResponse
	Response *DisassociateNetworkAclSubnetsResponseParams `json:"Response"`
}

func (r *DisassociateNetworkAclSubnetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisassociateNetworkAclSubnetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisassociateNetworkInterfaceSecurityGroupsRequestParams struct {
	// ENI instance ID, e.g. eni-pxir56ns. You can enter up to 100 instances for each request.
	NetworkInterfaceIds []*string `json:"NetworkInterfaceIds,omitnil,omitempty" name:"NetworkInterfaceIds"`

	// The security group instance ID, such as `sg-33ocnj9n`. It can be obtained through DescribeSecurityGroups. You can enter up to 100 instances for each request.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`
}

type DisassociateNetworkInterfaceSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// ENI instance ID, e.g. eni-pxir56ns. You can enter up to 100 instances for each request.
	NetworkInterfaceIds []*string `json:"NetworkInterfaceIds,omitnil,omitempty" name:"NetworkInterfaceIds"`

	// The security group instance ID, such as `sg-33ocnj9n`. It can be obtained through DescribeSecurityGroups. You can enter up to 100 instances for each request.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`
}

func (r *DisassociateNetworkInterfaceSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisassociateNetworkInterfaceSecurityGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NetworkInterfaceIds")
	delete(f, "SecurityGroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisassociateNetworkInterfaceSecurityGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisassociateNetworkInterfaceSecurityGroupsResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DisassociateNetworkInterfaceSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DisassociateNetworkInterfaceSecurityGroupsResponseParams `json:"Response"`
}

func (r *DisassociateNetworkInterfaceSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisassociateNetworkInterfaceSecurityGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisassociateVpcEndPointSecurityGroupsRequestParams struct {
	// Array of security group IDs
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// Endpoint ID
	EndPointId *string `json:"EndPointId,omitnil,omitempty" name:"EndPointId"`
}

type DisassociateVpcEndPointSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// Array of security group IDs
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// Endpoint ID
	EndPointId *string `json:"EndPointId,omitnil,omitempty" name:"EndPointId"`
}

func (r *DisassociateVpcEndPointSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisassociateVpcEndPointSecurityGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecurityGroupIds")
	delete(f, "EndPointId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisassociateVpcEndPointSecurityGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisassociateVpcEndPointSecurityGroupsResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DisassociateVpcEndPointSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DisassociateVpcEndPointSecurityGroupsResponseParams `json:"Response"`
}

func (r *DisassociateVpcEndPointSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisassociateVpcEndPointSecurityGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DownloadCustomerGatewayConfigurationRequestParams struct {
	// The ID of the VPN gateway instance.
	VpnGatewayId *string `json:"VpnGatewayId,omitnil,omitempty" name:"VpnGatewayId"`

	// The ID of the VPN tunnel instance, such as `vpnx-f49l6u0z`.
	VpnConnectionId *string `json:"VpnConnectionId,omitnil,omitempty" name:"VpnConnectionId"`

	// Customer gateway vendor information object, which can be obtained through DescribeCustomerGatewayVendors.
	CustomerGatewayVendor *CustomerGatewayVendor `json:"CustomerGatewayVendor,omitnil,omitempty" name:"CustomerGatewayVendor"`

	// Name of the physical API for tunnel access device.
	InterfaceName *string `json:"InterfaceName,omitnil,omitempty" name:"InterfaceName"`
}

type DownloadCustomerGatewayConfigurationRequest struct {
	*tchttp.BaseRequest
	
	// The ID of the VPN gateway instance.
	VpnGatewayId *string `json:"VpnGatewayId,omitnil,omitempty" name:"VpnGatewayId"`

	// The ID of the VPN tunnel instance, such as `vpnx-f49l6u0z`.
	VpnConnectionId *string `json:"VpnConnectionId,omitnil,omitempty" name:"VpnConnectionId"`

	// Customer gateway vendor information object, which can be obtained through DescribeCustomerGatewayVendors.
	CustomerGatewayVendor *CustomerGatewayVendor `json:"CustomerGatewayVendor,omitnil,omitempty" name:"CustomerGatewayVendor"`

	// Name of the physical API for tunnel access device.
	InterfaceName *string `json:"InterfaceName,omitnil,omitempty" name:"InterfaceName"`
}

func (r *DownloadCustomerGatewayConfigurationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DownloadCustomerGatewayConfigurationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpnGatewayId")
	delete(f, "VpnConnectionId")
	delete(f, "CustomerGatewayVendor")
	delete(f, "InterfaceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DownloadCustomerGatewayConfigurationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DownloadCustomerGatewayConfigurationResponseParams struct {
	// Configuration information in XML format.
	CustomerGatewayConfiguration *string `json:"CustomerGatewayConfiguration,omitnil,omitempty" name:"CustomerGatewayConfiguration"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DownloadCustomerGatewayConfigurationResponse struct {
	*tchttp.BaseResponse
	Response *DownloadCustomerGatewayConfigurationResponseParams `json:"Response"`
}

func (r *DownloadCustomerGatewayConfigurationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DownloadCustomerGatewayConfigurationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableCcnRoutesRequestParams struct {
	// The CCN instance ID, such as `ccn-f49l6u0z`.
	CcnId *string `json:"CcnId,omitnil,omitempty" name:"CcnId"`

	// The unique ID of the CCN routing policy, such as `ccnr-f49l6u0z`.
	RouteIds []*string `json:"RouteIds,omitnil,omitempty" name:"RouteIds"`
}

type EnableCcnRoutesRequest struct {
	*tchttp.BaseRequest
	
	// The CCN instance ID, such as `ccn-f49l6u0z`.
	CcnId *string `json:"CcnId,omitnil,omitempty" name:"CcnId"`

	// The unique ID of the CCN routing policy, such as `ccnr-f49l6u0z`.
	RouteIds []*string `json:"RouteIds,omitnil,omitempty" name:"RouteIds"`
}

func (r *EnableCcnRoutesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableCcnRoutesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CcnId")
	delete(f, "RouteIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EnableCcnRoutesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableCcnRoutesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type EnableCcnRoutesResponse struct {
	*tchttp.BaseResponse
	Response *EnableCcnRoutesResponseParams `json:"Response"`
}

func (r *EnableCcnRoutesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableCcnRoutesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableFlowLogsRequestParams struct {
	// Flow log ID.
	FlowLogIds []*string `json:"FlowLogIds,omitnil,omitempty" name:"FlowLogIds"`
}

type EnableFlowLogsRequest struct {
	*tchttp.BaseRequest
	
	// Flow log ID.
	FlowLogIds []*string `json:"FlowLogIds,omitnil,omitempty" name:"FlowLogIds"`
}

func (r *EnableFlowLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableFlowLogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FlowLogIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EnableFlowLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableFlowLogsResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type EnableFlowLogsResponse struct {
	*tchttp.BaseResponse
	Response *EnableFlowLogsResponseParams `json:"Response"`
}

func (r *EnableFlowLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableFlowLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableGatewayFlowMonitorRequestParams struct {
	// Gateway instance ID. Supported types:
	// Direct connect gateway instance, such as `dcg-ltjahce6`;
	// NAT gateway instance, such as `nat-ltjahce6`;
	// VPN gateway instance, such as `vpn-ltjahce6`.
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`
}

type EnableGatewayFlowMonitorRequest struct {
	*tchttp.BaseRequest
	
	// Gateway instance ID. Supported types:
	// Direct connect gateway instance, such as `dcg-ltjahce6`;
	// NAT gateway instance, such as `nat-ltjahce6`;
	// VPN gateway instance, such as `vpn-ltjahce6`.
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`
}

func (r *EnableGatewayFlowMonitorRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableGatewayFlowMonitorRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EnableGatewayFlowMonitorRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableGatewayFlowMonitorResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type EnableGatewayFlowMonitorResponse struct {
	*tchttp.BaseResponse
	Response *EnableGatewayFlowMonitorResponseParams `json:"Response"`
}

func (r *EnableGatewayFlowMonitorResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableGatewayFlowMonitorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableSnapshotPoliciesRequestParams struct {
	// Snapshot policy IDs
	SnapshotPolicyIds []*string `json:"SnapshotPolicyIds,omitnil,omitempty" name:"SnapshotPolicyIds"`
}

type EnableSnapshotPoliciesRequest struct {
	*tchttp.BaseRequest
	
	// Snapshot policy IDs
	SnapshotPolicyIds []*string `json:"SnapshotPolicyIds,omitnil,omitempty" name:"SnapshotPolicyIds"`
}

func (r *EnableSnapshotPoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableSnapshotPoliciesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SnapshotPolicyIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EnableSnapshotPoliciesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableSnapshotPoliciesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type EnableSnapshotPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *EnableSnapshotPoliciesResponseParams `json:"Response"`
}

func (r *EnableSnapshotPoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableSnapshotPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableVpcEndPointConnectRequestParams struct {
	// Endpoint service ID
	EndPointServiceId *string `json:"EndPointServiceId,omitnil,omitempty" name:"EndPointServiceId"`

	// Endpoint ID
	EndPointId []*string `json:"EndPointId,omitnil,omitempty" name:"EndPointId"`

	// Whether to accept the request of connecting with an endpoint
	AcceptFlag *bool `json:"AcceptFlag,omitnil,omitempty" name:"AcceptFlag"`
}

type EnableVpcEndPointConnectRequest struct {
	*tchttp.BaseRequest
	
	// Endpoint service ID
	EndPointServiceId *string `json:"EndPointServiceId,omitnil,omitempty" name:"EndPointServiceId"`

	// Endpoint ID
	EndPointId []*string `json:"EndPointId,omitnil,omitempty" name:"EndPointId"`

	// Whether to accept the request of connecting with an endpoint
	AcceptFlag *bool `json:"AcceptFlag,omitnil,omitempty" name:"AcceptFlag"`
}

func (r *EnableVpcEndPointConnectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableVpcEndPointConnectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EndPointServiceId")
	delete(f, "EndPointId")
	delete(f, "AcceptFlag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EnableVpcEndPointConnectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableVpcEndPointConnectResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type EnableVpcEndPointConnectResponse struct {
	*tchttp.BaseResponse
	Response *EnableVpcEndPointConnectResponseParams `json:"Response"`
}

func (r *EnableVpcEndPointConnectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableVpcEndPointConnectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EndPoint struct {
	// Endpoint ID
	EndPointId *string `json:"EndPointId,omitnil,omitempty" name:"EndPointId"`

	// VPC ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Subnet ID
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// APP ID
	EndPointOwner *string `json:"EndPointOwner,omitnil,omitempty" name:"EndPointOwner"`

	// Endpoint name
	EndPointName *string `json:"EndPointName,omitnil,omitempty" name:"EndPointName"`

	// Endpoint service VPC ID
	ServiceVpcId *string `json:"ServiceVpcId,omitnil,omitempty" name:"ServiceVpcId"`

	// Endpoint service VIP
	ServiceVip *string `json:"ServiceVip,omitnil,omitempty" name:"ServiceVip"`

	// Endpoint service ID
	EndPointServiceId *string `json:"EndPointServiceId,omitnil,omitempty" name:"EndPointServiceId"`

	// Endpoint VIP
	EndPointVip *string `json:"EndPointVip,omitnil,omitempty" name:"EndPointVip"`

	// Endpoint status. Valid values: `ACTIVE` (available), `PENDING` (to be accepted), `ACCEPTING` (being accepted), `REJECTED` (rejected), and `FAILED` (failed).
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// Creation time
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// ID list of security group instances bound with endpoints
	GroupSet []*string `json:"GroupSet,omitnil,omitempty" name:"GroupSet"`

	// Endpoint service name
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ServiceName *string `json:"ServiceName,omitnil,omitempty" name:"ServiceName"`
}

type EndPointService struct {
	// Endpoint service ID
	EndPointServiceId *string `json:"EndPointServiceId,omitnil,omitempty" name:"EndPointServiceId"`

	// VPC ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// APP ID
	ServiceOwner *string `json:"ServiceOwner,omitnil,omitempty" name:"ServiceOwner"`

	// Endpoint service name
	ServiceName *string `json:"ServiceName,omitnil,omitempty" name:"ServiceName"`

	// Real server VIP
	ServiceVip *string `json:"ServiceVip,omitnil,omitempty" name:"ServiceVip"`

	// Real server ID in the format of `lb-xxx`.
	ServiceInstanceId *string `json:"ServiceInstanceId,omitnil,omitempty" name:"ServiceInstanceId"`

	// Whether to automatically accept
	AutoAcceptFlag *bool `json:"AutoAcceptFlag,omitnil,omitempty" name:"AutoAcceptFlag"`

	// Number of associated endpoints
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	EndPointCount *uint64 `json:"EndPointCount,omitnil,omitempty" name:"EndPointCount"`

	// Array of endpoints
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	EndPointSet []*EndPoint `json:"EndPointSet,omitnil,omitempty" name:"EndPointSet"`

	// Creation time
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Mounted PaaS service type. Values: `CLB`, `CDB`, `CRS`
	ServiceType *string `json:"ServiceType,omitnil,omitempty" name:"ServiceType"`
}

type Filter struct {
	// The attribute name. If more than one Filter exists, the logical relation between these Filters is `AND`.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Attribute value. If multiple values exist in one filter, the logical relationship between these values is `OR`. For a `bool` parameter, the valid values include `TRUE` and `FALSE`.
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type FilterObject struct {
	// The attribute name. If more than one Filter exists, the logical relation between these Filters is `AND`.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// The attribute value. If there are multiple Values for one Filter, the logical relation between these Values under the same Filter is `OR`.
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type FlowLog struct {
	// ID of the VPC instance.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// The unique ID of the flow log.
	FlowLogId *string `json:"FlowLogId,omitnil,omitempty" name:"FlowLogId"`

	// The name of the flow log instance.
	FlowLogName *string `json:"FlowLogName,omitnil,omitempty" name:"FlowLogName"`

	// The type of resource associated with the flow log. Valid values: `VPC`, `SUBNET`, `NETWORKINTERFACE`, `CCN`, `NAT`, and `DCG`.
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// The unique ID of the resource
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// Type of flow logs to be collected. Valid values: `ACCEPT`, `REJECT` and `ALL`.
	TrafficType *string `json:"TrafficType,omitnil,omitempty" name:"TrafficType"`

	// The storage ID of the flow log
	CloudLogId *string `json:"CloudLogId,omitnil,omitempty" name:"CloudLogId"`

	// Flow log storage ID status.
	CloudLogState *string `json:"CloudLogState,omitnil,omitempty" name:"CloudLogState"`

	// The flow log description.
	FlowLogDescription *string `json:"FlowLogDescription,omitnil,omitempty" name:"FlowLogDescription"`

	// The creation time of the flow log.
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// Tag list, such as [{"Key": "city", "Value": "shanghai"}].
	TagSet []*Tag `json:"TagSet,omitnil,omitempty" name:"TagSet"`

	// Whether to enable. `true`: yes; `false`: no.
	Enable *bool `json:"Enable,omitnil,omitempty" name:"Enable"`

	// Consumer end types: cls and ckafka
	// Note: this field may return `null`, indicating that no valid value can be found.
	StorageType *string `json:"StorageType,omitnil,omitempty" name:"StorageType"`

	// Information of the consumer, which is returned when the consumer type is `ckafka`.
	// Note: this field may return `null`, indicating that no valid value can be found.
	FlowLogStorage *FlowLogStorage `json:"FlowLogStorage,omitnil,omitempty" name:"FlowLogStorage"`

	// The region corresponding to the flow log storage ID.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	CloudLogRegion *string `json:"CloudLogRegion,omitnil,omitempty" name:"CloudLogRegion"`
}

type FlowLogStorage struct {
	// Storage instance ID, which is required when `StorageType` is `ckafka`.
	StorageId *string `json:"StorageId,omitnil,omitempty" name:"StorageId"`

	// Topic ID, which is required when `StorageType` is `ckafka`.
	// Note: this field may return `null`, indicating that no valid value can be found.
	StorageTopic *string `json:"StorageTopic,omitnil,omitempty" name:"StorageTopic"`
}

type GatewayFlowMonitorDetail struct {
	// Origin `IP`.
	PrivateIpAddress *string `json:"PrivateIpAddress,omitnil,omitempty" name:"PrivateIpAddress"`

	// Inbound packets.
	InPkg *uint64 `json:"InPkg,omitnil,omitempty" name:"InPkg"`

	// Outbound packets.
	OutPkg *uint64 `json:"OutPkg,omitnil,omitempty" name:"OutPkg"`

	// Inbound traffic, in Byte.
	InTraffic *uint64 `json:"InTraffic,omitnil,omitempty" name:"InTraffic"`

	// Outbound traffic, in Byte.
	OutTraffic *uint64 `json:"OutTraffic,omitnil,omitempty" name:"OutTraffic"`
}

type GatewayQos struct {
	// VPC instance ID.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// CVM Private IP.
	IpAddress *string `json:"IpAddress,omitnil,omitempty" name:"IpAddress"`

	// Bandwidth limit value.
	Bandwidth *int64 `json:"Bandwidth,omitnil,omitempty" name:"Bandwidth"`

	// Creation time.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`
}

// Predefined struct for user
type GenerateVpnConnectionDefaultHealthCheckIpRequestParams struct {
	// VPN gateway ID, such as `vpngw-1w9tue3d`
	VpnGatewayId *string `json:"VpnGatewayId,omitnil,omitempty" name:"VpnGatewayId"`
}

type GenerateVpnConnectionDefaultHealthCheckIpRequest struct {
	*tchttp.BaseRequest
	
	// VPN gateway ID, such as `vpngw-1w9tue3d`
	VpnGatewayId *string `json:"VpnGatewayId,omitnil,omitempty" name:"VpnGatewayId"`
}

func (r *GenerateVpnConnectionDefaultHealthCheckIpRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GenerateVpnConnectionDefaultHealthCheckIpRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpnGatewayId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GenerateVpnConnectionDefaultHealthCheckIpRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GenerateVpnConnectionDefaultHealthCheckIpResponseParams struct {
	// Local IP used for VPN tunnel health check
	HealthCheckLocalIp *string `json:"HealthCheckLocalIp,omitnil,omitempty" name:"HealthCheckLocalIp"`

	// Remote IP used for VPN tunnel health check
	HealthCheckRemoteIp *string `json:"HealthCheckRemoteIp,omitnil,omitempty" name:"HealthCheckRemoteIp"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GenerateVpnConnectionDefaultHealthCheckIpResponse struct {
	*tchttp.BaseResponse
	Response *GenerateVpnConnectionDefaultHealthCheckIpResponseParams `json:"Response"`
}

func (r *GenerateVpnConnectionDefaultHealthCheckIpResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GenerateVpnConnectionDefaultHealthCheckIpResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetCcnRegionBandwidthLimitsRequestParams struct {
	// The CCN instance ID, such as `ccn-f49l6u0z`.
	CcnId *string `json:"CcnId,omitnil,omitempty" name:"CcnId"`

	// The filter condition.
	// <li>sregion - String - (Filter condition) Filter by the source region, such as 'ap-guangzhou'.</li>
	// <li>dregion - String - (Filter condition) Filter by the destination region, such as 'ap-shanghai-bm'.</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// The sorting condition. Valid values: `BandwidthLimit` and `ExpireTime`.
	SortedBy *string `json:"SortedBy,omitnil,omitempty" name:"SortedBy"`

	// Offset
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Quantity of returned items
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// In ascending or descending order. Valid values: 'ASC' and 'DESC'.
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`
}

type GetCcnRegionBandwidthLimitsRequest struct {
	*tchttp.BaseRequest
	
	// The CCN instance ID, such as `ccn-f49l6u0z`.
	CcnId *string `json:"CcnId,omitnil,omitempty" name:"CcnId"`

	// The filter condition.
	// <li>sregion - String - (Filter condition) Filter by the source region, such as 'ap-guangzhou'.</li>
	// <li>dregion - String - (Filter condition) Filter by the destination region, such as 'ap-shanghai-bm'.</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// The sorting condition. Valid values: `BandwidthLimit` and `ExpireTime`.
	SortedBy *string `json:"SortedBy,omitnil,omitempty" name:"SortedBy"`

	// Offset
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Quantity of returned items
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// In ascending or descending order. Valid values: 'ASC' and 'DESC'.
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`
}

func (r *GetCcnRegionBandwidthLimitsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetCcnRegionBandwidthLimitsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CcnId")
	delete(f, "Filters")
	delete(f, "SortedBy")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "OrderBy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetCcnRegionBandwidthLimitsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetCcnRegionBandwidthLimitsResponseParams struct {
	// The outbound bandwidth limits of regions in a CCN instance.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CcnBandwidthSet []*CcnBandwidthInfo `json:"CcnBandwidthSet,omitnil,omitempty" name:"CcnBandwidthSet"`

	// The number of eligible objects.
	// Note: this field may return null, indicating that no valid value was found.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetCcnRegionBandwidthLimitsResponse struct {
	*tchttp.BaseResponse
	Response *GetCcnRegionBandwidthLimitsResponseParams `json:"Response"`
}

func (r *GetCcnRegionBandwidthLimitsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetCcnRegionBandwidthLimitsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type HaVip struct {
	// The `ID` of the `HAVIP`. This is the unique identifier of the `HAVIP`.
	HaVipId *string `json:"HaVipId,omitnil,omitempty" name:"HaVipId"`

	// The name of the `HAVIP`.
	HaVipName *string `json:"HaVipName,omitnil,omitempty" name:"HaVipName"`

	// The virtual IP address.
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// The `ID` of the VPC to which the `HAVIP` belongs.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// The `ID` of the subnet to which the `HAVIP` belongs.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// The `ID` of the ENI associated with the `HAVIP`.
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitnil,omitempty" name:"NetworkInterfaceId"`

	// The `ID` of the bound instance.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Bound `EIP`.
	AddressIp *string `json:"AddressIp,omitnil,omitempty" name:"AddressIp"`

	// Status:
	// <li>`AVAILABLE`: Operating</li>
	// <li>`UNBIND`: Not bound</li>
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// Creation Time.
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// Identifier for businesses that use HAVIP.
	Business *string `json:"Business,omitnil,omitempty" name:"Business"`
}

// Predefined struct for user
type HaVipAssociateAddressIpRequestParams struct {
	// The unique `ID` of the `HAVIP`, such as `havip-9o233uri`. This must be a `HAVIP` that has not been bound to an `EIP`
	HaVipId *string `json:"HaVipId,omitnil,omitempty" name:"HaVipId"`

	// The Elastic `IP`. This must be an `EIP` that has not been bound to a `HAVIP`
	AddressIp *string `json:"AddressIp,omitnil,omitempty" name:"AddressIp"`
}

type HaVipAssociateAddressIpRequest struct {
	*tchttp.BaseRequest
	
	// The unique `ID` of the `HAVIP`, such as `havip-9o233uri`. This must be a `HAVIP` that has not been bound to an `EIP`
	HaVipId *string `json:"HaVipId,omitnil,omitempty" name:"HaVipId"`

	// The Elastic `IP`. This must be an `EIP` that has not been bound to a `HAVIP`
	AddressIp *string `json:"AddressIp,omitnil,omitempty" name:"AddressIp"`
}

func (r *HaVipAssociateAddressIpRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *HaVipAssociateAddressIpRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "HaVipId")
	delete(f, "AddressIp")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "HaVipAssociateAddressIpRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type HaVipAssociateAddressIpResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type HaVipAssociateAddressIpResponse struct {
	*tchttp.BaseResponse
	Response *HaVipAssociateAddressIpResponseParams `json:"Response"`
}

func (r *HaVipAssociateAddressIpResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *HaVipAssociateAddressIpResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type HaVipDisassociateAddressIpRequestParams struct {
	// The unique `ID` of the `HAVIP`, such as `havip-9o233uri`. This must be an `HAVIP` that has been bound to an `EIP`.
	HaVipId *string `json:"HaVipId,omitnil,omitempty" name:"HaVipId"`
}

type HaVipDisassociateAddressIpRequest struct {
	*tchttp.BaseRequest
	
	// The unique `ID` of the `HAVIP`, such as `havip-9o233uri`. This must be an `HAVIP` that has been bound to an `EIP`.
	HaVipId *string `json:"HaVipId,omitnil,omitempty" name:"HaVipId"`
}

func (r *HaVipDisassociateAddressIpRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *HaVipDisassociateAddressIpRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "HaVipId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "HaVipDisassociateAddressIpRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type HaVipDisassociateAddressIpResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type HaVipDisassociateAddressIpResponse struct {
	*tchttp.BaseResponse
	Response *HaVipDisassociateAddressIpResponseParams `json:"Response"`
}

func (r *HaVipDisassociateAddressIpResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *HaVipDisassociateAddressIpResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IKEOptionsSpecification struct {
	// Encryption algorithm. Valid values: `3DES-CBC`, `AES-CBC-128`, `AES-CBS-192`, `AES-CBC-256`, `DES-CBC`, and `SM4`; default value: `3DES-CBC`.
	PropoEncryAlgorithm *string `json:"PropoEncryAlgorithm,omitnil,omitempty" name:"PropoEncryAlgorithm"`

	// Authentication algorithm. Valid values: `MD5`, `SHA1` and `SHA-256`; default value: `MD5`.
	PropoAuthenAlgorithm *string `json:"PropoAuthenAlgorithm,omitnil,omitempty" name:"PropoAuthenAlgorithm"`

	// Negotiation mode. Available values: 'AGGRESSIVE' and 'MAIN'. Default is MAIN.
	ExchangeMode *string `json:"ExchangeMode,omitnil,omitempty" name:"ExchangeMode"`

	// Type of local identity. Available values: 'ADDRESS' and 'FQDN'. Default is ADDRESS.
	LocalIdentity *string `json:"LocalIdentity,omitnil,omitempty" name:"LocalIdentity"`

	// Type of remote identity. Available values: 'ADDRESS' and 'FQDN'. Default is ADDRESS.
	RemoteIdentity *string `json:"RemoteIdentity,omitnil,omitempty" name:"RemoteIdentity"`

	// Local identity. When ADDRESS is selected for LocalIdentity, LocalAddress is required. The default LocalAddress is the public IP of the VPN gateway.
	LocalAddress *string `json:"LocalAddress,omitnil,omitempty" name:"LocalAddress"`

	// Remote identity. When ADDRESS is selected for RemoteIdentity, RemoteAddress is required.
	RemoteAddress *string `json:"RemoteAddress,omitnil,omitempty" name:"RemoteAddress"`

	// Local identity. When FQDN is selected for LocalIdentity, LocalFqdnName is required.
	LocalFqdnName *string `json:"LocalFqdnName,omitnil,omitempty" name:"LocalFqdnName"`

	// Remote identity. When FQDN is selected for RemoteIdentity, RemoteFqdnName is required.
	RemoteFqdnName *string `json:"RemoteFqdnName,omitnil,omitempty" name:"RemoteFqdnName"`

	// DH group. Specify the DH group used for exchanging the key via IKE. Available values: 'GROUP1', 'GROUP2', 'GROUP5', 'GROUP14', and 'GROUP24'.
	DhGroupName *string `json:"DhGroupName,omitnil,omitempty" name:"DhGroupName"`

	// IKE SA lifetime (in sec). Value range: 60-604800
	IKESaLifetimeSeconds *uint64 `json:"IKESaLifetimeSeconds,omitnil,omitempty" name:"IKESaLifetimeSeconds"`

	// IKE version
	IKEVersion *string `json:"IKEVersion,omitnil,omitempty" name:"IKEVersion"`
}

type IPSECOptionsSpecification struct {
	// Encryption algorithm. Valid values: `3DES-CBC`, `AES-CBC-128`, `AES-CBC-192`, `AES-CBC-256`, `DES-CBC`, `SM4`, and `NULL`; default value: `AES-CBC-128`.
	EncryptAlgorithm *string `json:"EncryptAlgorithm,omitnil,omitempty" name:"EncryptAlgorithm"`

	// Authentication algorithm. Valid values: `MD5`, `SHA1` and `SHA-256`; default value: `SHA1`.
	IntegrityAlgorith *string `json:"IntegrityAlgorith,omitnil,omitempty" name:"IntegrityAlgorith"`

	// IPsec SA lifetime (in sec). Value range: 180-604800
	IPSECSaLifetimeSeconds *uint64 `json:"IPSECSaLifetimeSeconds,omitnil,omitempty" name:"IPSECSaLifetimeSeconds"`

	// PFS. Available value: 'NULL', 'DH-GROUP1', 'DH-GROUP2', 'DH-GROUP5', 'DH-GROUP14', and 'DH-GROUP24'. Default is NULL.
	PfsDhGroup *string `json:"PfsDhGroup,omitnil,omitempty" name:"PfsDhGroup"`

	// IPsec SA lifetime (in KB). Value range: 2560-604800
	IPSECSaLifetimeTraffic *uint64 `json:"IPSECSaLifetimeTraffic,omitnil,omitempty" name:"IPSECSaLifetimeTraffic"`
}

// Predefined struct for user
type InquirePriceCreateDirectConnectGatewayRequestParams struct {

}

type InquirePriceCreateDirectConnectGatewayRequest struct {
	*tchttp.BaseRequest
	
}

func (r *InquirePriceCreateDirectConnectGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquirePriceCreateDirectConnectGatewayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquirePriceCreateDirectConnectGatewayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquirePriceCreateDirectConnectGatewayResponseParams struct {
	// Standard access fee for a direct connect gateway
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	TotalCost *int64 `json:"TotalCost,omitnil,omitempty" name:"TotalCost"`

	// Actual access fee for a direct connect gateway
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	RealTotalCost *int64 `json:"RealTotalCost,omitnil,omitempty" name:"RealTotalCost"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type InquirePriceCreateDirectConnectGatewayResponse struct {
	*tchttp.BaseResponse
	Response *InquirePriceCreateDirectConnectGatewayResponseParams `json:"Response"`
}

func (r *InquirePriceCreateDirectConnectGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquirePriceCreateDirectConnectGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceRenewVpnGatewayRequestParams struct {
	// The ID of the VPN gateway instance.
	VpnGatewayId *string `json:"VpnGatewayId,omitnil,omitempty" name:"VpnGatewayId"`

	// Specifies the purchased validity period, whether to enable auto-renewal. This parameter is required for monthly-subscription instances.
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitnil,omitempty" name:"InstanceChargePrepaid"`
}

type InquiryPriceRenewVpnGatewayRequest struct {
	*tchttp.BaseRequest
	
	// The ID of the VPN gateway instance.
	VpnGatewayId *string `json:"VpnGatewayId,omitnil,omitempty" name:"VpnGatewayId"`

	// Specifies the purchased validity period, whether to enable auto-renewal. This parameter is required for monthly-subscription instances.
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitnil,omitempty" name:"InstanceChargePrepaid"`
}

func (r *InquiryPriceRenewVpnGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceRenewVpnGatewayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpnGatewayId")
	delete(f, "InstanceChargePrepaid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquiryPriceRenewVpnGatewayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceRenewVpnGatewayResponseParams struct {
	// Product price.
	Price *Price `json:"Price,omitnil,omitempty" name:"Price"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type InquiryPriceRenewVpnGatewayResponse struct {
	*tchttp.BaseResponse
	Response *InquiryPriceRenewVpnGatewayResponseParams `json:"Response"`
}

func (r *InquiryPriceRenewVpnGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceRenewVpnGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceResetVpnGatewayInternetMaxBandwidthRequestParams struct {
	// The ID of the VPN gateway instance.
	VpnGatewayId *string `json:"VpnGatewayId,omitnil,omitempty" name:"VpnGatewayId"`

	// The public network bandwidth configuration. Available bandwidth specifications: 5, 10, 20, 50, and 100. Unit: Mbps.
	InternetMaxBandwidthOut *uint64 `json:"InternetMaxBandwidthOut,omitnil,omitempty" name:"InternetMaxBandwidthOut"`
}

type InquiryPriceResetVpnGatewayInternetMaxBandwidthRequest struct {
	*tchttp.BaseRequest
	
	// The ID of the VPN gateway instance.
	VpnGatewayId *string `json:"VpnGatewayId,omitnil,omitempty" name:"VpnGatewayId"`

	// The public network bandwidth configuration. Available bandwidth specifications: 5, 10, 20, 50, and 100. Unit: Mbps.
	InternetMaxBandwidthOut *uint64 `json:"InternetMaxBandwidthOut,omitnil,omitempty" name:"InternetMaxBandwidthOut"`
}

func (r *InquiryPriceResetVpnGatewayInternetMaxBandwidthRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceResetVpnGatewayInternetMaxBandwidthRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpnGatewayId")
	delete(f, "InternetMaxBandwidthOut")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquiryPriceResetVpnGatewayInternetMaxBandwidthRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceResetVpnGatewayInternetMaxBandwidthResponseParams struct {
	// Product price.
	Price *Price `json:"Price,omitnil,omitempty" name:"Price"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type InquiryPriceResetVpnGatewayInternetMaxBandwidthResponse struct {
	*tchttp.BaseResponse
	Response *InquiryPriceResetVpnGatewayInternetMaxBandwidthResponseParams `json:"Response"`
}

func (r *InquiryPriceResetVpnGatewayInternetMaxBandwidthResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceResetVpnGatewayInternetMaxBandwidthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstanceChargePrepaid struct {
	// Purchased usage period (in month). Value range: [1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36].
	Period *uint64 `json:"Period,omitnil,omitempty" name:"Period"`

	// Auto-renewal ID. Value range: NOTIFY_AND_AUTO_RENEW: notify expiry and renew automatically, NOTIFY_AND_MANUAL_RENEW: notify expiry but do not renew automatically. The default is NOTIFY_AND_MANUAL_RENEW
	RenewFlag *string `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`
}

type InstanceStatistic struct {
	// Type of instance
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// Number of instances
	InstanceCount *uint64 `json:"InstanceCount,omitnil,omitempty" name:"InstanceCount"`
}

type IpAddressStates struct {
	// VPC instance ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Subnet instance ID
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// IP address
	IpAddress *string `json:"IpAddress,omitnil,omitempty" name:"IpAddress"`

	// Resource type
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// Resource ID
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`
}

type IpField struct {
	// Country/region of the IP
	Country *bool `json:"Country,omitnil,omitempty" name:"Country"`

	// Province/municipality/state of the IP
	Province *bool `json:"Province,omitnil,omitempty" name:"Province"`

	// City of the IP
	City *bool `json:"City,omitnil,omitempty" name:"City"`

	// City district of the IP
	Region *bool `json:"Region,omitnil,omitempty" name:"Region"`

	// Access ISP field
	Isp *bool `json:"Isp,omitnil,omitempty" name:"Isp"`

	// ISP backbone network’s AS field
	AsName *bool `json:"AsName,omitnil,omitempty" name:"AsName"`

	// Backbone AS ID
	AsId *bool `json:"AsId,omitnil,omitempty" name:"AsId"`

	// Comment
	Comment *bool `json:"Comment,omitnil,omitempty" name:"Comment"`
}

type IpGeolocationInfo struct {
	// Country/region
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Country *string `json:"Country,omitnil,omitempty" name:"Country"`

	// Province- or municipality-level administrative region
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Province *string `json:"Province,omitnil,omitempty" name:"Province"`

	// Municipal administrative region
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	City *string `json:"City,omitnil,omitempty" name:"City"`

	// Urban area
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Access ISP
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Isp *string `json:"Isp,omitnil,omitempty" name:"Isp"`

	// ISP backbone network’s AS name
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	AsName *string `json:"AsName,omitnil,omitempty" name:"AsName"`

	// ISP backbone network’s AS ID
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	AsId *string `json:"AsId,omitnil,omitempty" name:"AsId"`

	// Comment. The APN value of mobile users is entered currently. If there is no APN attribute, this is `null`.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// IP address
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	AddressIp *string `json:"AddressIp,omitnil,omitempty" name:"AddressIp"`
}

type Ipv6Address struct {
	// `IPv6` address, such as `3402:4e00:20:100:0:8cd9:2a67:71f3`
	Address *string `json:"Address,omitnil,omitempty" name:"Address"`

	// Whether it is a primary `IP`.
	Primary *bool `json:"Primary,omitnil,omitempty" name:"Primary"`

	// The `ID` of the `EIP` instance, such as `eip-hxlqja90`.
	AddressId *string `json:"AddressId,omitnil,omitempty" name:"AddressId"`

	// Message description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Whether the public IP is blocked.
	IsWanIpBlocked *bool `json:"IsWanIpBlocked,omitnil,omitempty" name:"IsWanIpBlocked"`

	// `IPv6` address status:
	// <li>`PENDING`: Creating</li>
	// <li>`MIGRATING`: Migrating</li>
	// <li>`DELETING`: Deleting</li>
	// <li>`AVAILABLE`: Available</li>
	State *string `json:"State,omitnil,omitempty" name:"State"`
}

type Ipv6SubnetCidrBlock struct {
	// The `ID` of the subnet instance, such as `subnet-pxir56ns`.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// The `IPv6` subnet IP range, such as `3402:4e00:20:1001::/64`
	Ipv6CidrBlock *string `json:"Ipv6CidrBlock,omitnil,omitempty" name:"Ipv6CidrBlock"`
}

type ItemPrice struct {
	// The pay-as-you-go billing method. Unit: CNY.
	UnitPrice *float64 `json:"UnitPrice,omitnil,omitempty" name:"UnitPrice"`

	// Pay-as-you-go billing method. Value Range: HOUR: Indicates billing by the hour. Scenarios using this hourly billing unit include: Instances postpaid on an hourly basis (POSTPAID_BY_HOUR), and bandwidth postpaid on an hourly basis (BANDWIDTH_POSTPAID_BY_HOUR). GB: Indicates billing on a per-GB basis. Scenarios using this billing unit include: Traffic postpaid on an hourly basis (TRAFFIC_POSTPAID_BY_HOUR).
	ChargeUnit *string `json:"ChargeUnit,omitnil,omitempty" name:"ChargeUnit"`

	// Original price of the prepaid product. Unit: CNY.
	OriginalPrice *float64 `json:"OriginalPrice,omitnil,omitempty" name:"OriginalPrice"`

	// Discount price of the prepaid product. Unit: CNY.
	DiscountPrice *float64 `json:"DiscountPrice,omitnil,omitempty" name:"DiscountPrice"`
}

type LocalGateway struct {
	// CDC instance ID
	CdcId *string `json:"CdcId,omitnil,omitempty" name:"CdcId"`

	// VPC instance ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Local gateway instance ID
	UniqLocalGwId *string `json:"UniqLocalGwId,omitnil,omitempty" name:"UniqLocalGwId"`

	// Local gateway name
	LocalGatewayName *string `json:"LocalGatewayName,omitnil,omitempty" name:"LocalGatewayName"`

	// Local gateway IP
	LocalGwIp *string `json:"LocalGwIp,omitnil,omitempty" name:"LocalGwIp"`

	// Creation time of the local gateway
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`
}

// Predefined struct for user
type MigrateNetworkInterfaceRequestParams struct {
	// The ID of the ENI instance, such as `eni-m6dyj72l`.
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitnil,omitempty" name:"NetworkInterfaceId"`

	// The ID of the CVM bound to the ENI, such as `ins-r8hr2upy`.
	SourceInstanceId *string `json:"SourceInstanceId,omitnil,omitempty" name:"SourceInstanceId"`

	// ID of the destination CVM instance to be migrated.
	DestinationInstanceId *string `json:"DestinationInstanceId,omitnil,omitempty" name:"DestinationInstanceId"`

	// ENI mount method. Valid values: 0: standard; 1: extension; default value: 0
	AttachType *uint64 `json:"AttachType,omitnil,omitempty" name:"AttachType"`
}

type MigrateNetworkInterfaceRequest struct {
	*tchttp.BaseRequest
	
	// The ID of the ENI instance, such as `eni-m6dyj72l`.
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitnil,omitempty" name:"NetworkInterfaceId"`

	// The ID of the CVM bound to the ENI, such as `ins-r8hr2upy`.
	SourceInstanceId *string `json:"SourceInstanceId,omitnil,omitempty" name:"SourceInstanceId"`

	// ID of the destination CVM instance to be migrated.
	DestinationInstanceId *string `json:"DestinationInstanceId,omitnil,omitempty" name:"DestinationInstanceId"`

	// ENI mount method. Valid values: 0: standard; 1: extension; default value: 0
	AttachType *uint64 `json:"AttachType,omitnil,omitempty" name:"AttachType"`
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
	delete(f, "NetworkInterfaceId")
	delete(f, "SourceInstanceId")
	delete(f, "DestinationInstanceId")
	delete(f, "AttachType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "MigrateNetworkInterfaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type MigrateNetworkInterfaceResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// ID of the ENI instance bound with the private IP, such as `eni-m6dyj72l`.
	SourceNetworkInterfaceId *string `json:"SourceNetworkInterfaceId,omitnil,omitempty" name:"SourceNetworkInterfaceId"`

	// ID of the destination ENI instance to be migrated.
	DestinationNetworkInterfaceId *string `json:"DestinationNetworkInterfaceId,omitnil,omitempty" name:"DestinationNetworkInterfaceId"`

	// The private IP to be migrated, such as 10.0.0.6.
	PrivateIpAddress *string `json:"PrivateIpAddress,omitnil,omitempty" name:"PrivateIpAddress"`
}

type MigratePrivateIpAddressRequest struct {
	*tchttp.BaseRequest
	
	// ID of the ENI instance bound with the private IP, such as `eni-m6dyj72l`.
	SourceNetworkInterfaceId *string `json:"SourceNetworkInterfaceId,omitnil,omitempty" name:"SourceNetworkInterfaceId"`

	// ID of the destination ENI instance to be migrated.
	DestinationNetworkInterfaceId *string `json:"DestinationNetworkInterfaceId,omitnil,omitempty" name:"DestinationNetworkInterfaceId"`

	// The private IP to be migrated, such as 10.0.0.6.
	PrivateIpAddress *string `json:"PrivateIpAddress,omitnil,omitempty" name:"PrivateIpAddress"`
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
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// The unique ID of the EIP, such as `eip-11112222`.
	AddressId *string `json:"AddressId,omitnil,omitempty" name:"AddressId"`

	// The EIP name after modification. The maximum length is 20 characters.
	AddressName *string `json:"AddressName,omitnil,omitempty" name:"AddressName"`

	// Whether the set EIP is a direct connection EIP. TRUE: yes. FALSE: no. Note that this parameter is available only to users who have activated the EIP direct connection function.
	EipDirectConnection *string `json:"EipDirectConnection,omitnil,omitempty" name:"EipDirectConnection"`
}

type ModifyAddressAttributeRequest struct {
	*tchttp.BaseRequest
	
	// The unique ID of the EIP, such as `eip-11112222`.
	AddressId *string `json:"AddressId,omitnil,omitempty" name:"AddressId"`

	// The EIP name after modification. The maximum length is 20 characters.
	AddressName *string `json:"AddressName,omitnil,omitempty" name:"AddressName"`

	// Whether the set EIP is a direct connection EIP. TRUE: yes. FALSE: no. Note that this parameter is available only to users who have activated the EIP direct connection function.
	EipDirectConnection *string `json:"EipDirectConnection,omitnil,omitempty" name:"EipDirectConnection"`
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
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type ModifyAddressInternetChargeTypeRequestParams struct {
	// Unique EIP ID, such as "eip-xxxx"
	AddressId *string `json:"AddressId,omitnil,omitempty" name:"AddressId"`

	// The target billing method. It can be `BANDWIDTH_PREPAID_BY_MONTH` or `TRAFFIC_POSTPAID_BY_HOUR`
	InternetChargeType *string `json:"InternetChargeType,omitnil,omitempty" name:"InternetChargeType"`

	// The target bandwidth value
	InternetMaxBandwidthOut *uint64 `json:"InternetMaxBandwidthOut,omitnil,omitempty" name:"InternetMaxBandwidthOut"`

	// Billing details of monthly-subscribed network bandwidth. This parameter is required if the target billing method is `BANDWIDTH_PREPAID_BY_MONTH`.
	AddressChargePrepaid *AddressChargePrepaid `json:"AddressChargePrepaid,omitnil,omitempty" name:"AddressChargePrepaid"`
}

type ModifyAddressInternetChargeTypeRequest struct {
	*tchttp.BaseRequest
	
	// Unique EIP ID, such as "eip-xxxx"
	AddressId *string `json:"AddressId,omitnil,omitempty" name:"AddressId"`

	// The target billing method. It can be `BANDWIDTH_PREPAID_BY_MONTH` or `TRAFFIC_POSTPAID_BY_HOUR`
	InternetChargeType *string `json:"InternetChargeType,omitnil,omitempty" name:"InternetChargeType"`

	// The target bandwidth value
	InternetMaxBandwidthOut *uint64 `json:"InternetMaxBandwidthOut,omitnil,omitempty" name:"InternetMaxBandwidthOut"`

	// Billing details of monthly-subscribed network bandwidth. This parameter is required if the target billing method is `BANDWIDTH_PREPAID_BY_MONTH`.
	AddressChargePrepaid *AddressChargePrepaid `json:"AddressChargePrepaid,omitnil,omitempty" name:"AddressChargePrepaid"`
}

func (r *ModifyAddressInternetChargeTypeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAddressInternetChargeTypeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AddressId")
	delete(f, "InternetChargeType")
	delete(f, "InternetMaxBandwidthOut")
	delete(f, "AddressChargePrepaid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAddressInternetChargeTypeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAddressInternetChargeTypeResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAddressInternetChargeTypeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAddressInternetChargeTypeResponseParams `json:"Response"`
}

func (r *ModifyAddressInternetChargeTypeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAddressInternetChargeTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAddressTemplateAttributeRequestParams struct {
	// IP address template instance ID, such as `ipm-mdunqeb6`.
	AddressTemplateId *string `json:"AddressTemplateId,omitnil,omitempty" name:"AddressTemplateId"`

	// IP address template name.
	AddressTemplateName *string `json:"AddressTemplateName,omitnil,omitempty" name:"AddressTemplateName"`

	// Address information, including IP, CIDR and IP address range.
	Addresses []*string `json:"Addresses,omitnil,omitempty" name:"Addresses"`

	// Address information with remarks, including the IP, CIDR block or IP address range.
	AddressesExtra []*AddressInfo `json:"AddressesExtra,omitnil,omitempty" name:"AddressesExtra"`
}

type ModifyAddressTemplateAttributeRequest struct {
	*tchttp.BaseRequest
	
	// IP address template instance ID, such as `ipm-mdunqeb6`.
	AddressTemplateId *string `json:"AddressTemplateId,omitnil,omitempty" name:"AddressTemplateId"`

	// IP address template name.
	AddressTemplateName *string `json:"AddressTemplateName,omitnil,omitempty" name:"AddressTemplateName"`

	// Address information, including IP, CIDR and IP address range.
	Addresses []*string `json:"Addresses,omitnil,omitempty" name:"Addresses"`

	// Address information with remarks, including the IP, CIDR block or IP address range.
	AddressesExtra []*AddressInfo `json:"AddressesExtra,omitnil,omitempty" name:"AddressesExtra"`
}

func (r *ModifyAddressTemplateAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAddressTemplateAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AddressTemplateId")
	delete(f, "AddressTemplateName")
	delete(f, "Addresses")
	delete(f, "AddressesExtra")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAddressTemplateAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAddressTemplateAttributeResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAddressTemplateAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAddressTemplateAttributeResponseParams `json:"Response"`
}

func (r *ModifyAddressTemplateAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAddressTemplateAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAddressTemplateGroupAttributeRequestParams struct {
	// IP address template group instance ID, such as `ipmg-2uw6ujo6`.
	AddressTemplateGroupId *string `json:"AddressTemplateGroupId,omitnil,omitempty" name:"AddressTemplateGroupId"`

	// IP address template group name.
	AddressTemplateGroupName *string `json:"AddressTemplateGroupName,omitnil,omitempty" name:"AddressTemplateGroupName"`

	// IP address template instance ID, such as `ipm-mdunqeb6`.
	AddressTemplateIds []*string `json:"AddressTemplateIds,omitnil,omitempty" name:"AddressTemplateIds"`
}

type ModifyAddressTemplateGroupAttributeRequest struct {
	*tchttp.BaseRequest
	
	// IP address template group instance ID, such as `ipmg-2uw6ujo6`.
	AddressTemplateGroupId *string `json:"AddressTemplateGroupId,omitnil,omitempty" name:"AddressTemplateGroupId"`

	// IP address template group name.
	AddressTemplateGroupName *string `json:"AddressTemplateGroupName,omitnil,omitempty" name:"AddressTemplateGroupName"`

	// IP address template instance ID, such as `ipm-mdunqeb6`.
	AddressTemplateIds []*string `json:"AddressTemplateIds,omitnil,omitempty" name:"AddressTemplateIds"`
}

func (r *ModifyAddressTemplateGroupAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAddressTemplateGroupAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AddressTemplateGroupId")
	delete(f, "AddressTemplateGroupName")
	delete(f, "AddressTemplateIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAddressTemplateGroupAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAddressTemplateGroupAttributeResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAddressTemplateGroupAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAddressTemplateGroupAttributeResponseParams `json:"Response"`
}

func (r *ModifyAddressTemplateGroupAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAddressTemplateGroupAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAddressesBandwidthRequestParams struct {
	// List of EIP IDs, such as “eip-xxxx”.
	AddressIds []*string `json:"AddressIds,omitnil,omitempty" name:"AddressIds"`

	// Target bandwidth value adjustment
	InternetMaxBandwidthOut *int64 `json:"InternetMaxBandwidthOut,omitnil,omitempty" name:"InternetMaxBandwidthOut"`

	// (Disused) The start time of the monthly bandwidth subscription
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// (Disused) The end time of the monthly bandwidth subscription
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type ModifyAddressesBandwidthRequest struct {
	*tchttp.BaseRequest
	
	// List of EIP IDs, such as “eip-xxxx”.
	AddressIds []*string `json:"AddressIds,omitnil,omitempty" name:"AddressIds"`

	// Target bandwidth value adjustment
	InternetMaxBandwidthOut *int64 `json:"InternetMaxBandwidthOut,omitnil,omitempty" name:"InternetMaxBandwidthOut"`

	// (Disused) The start time of the monthly bandwidth subscription
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// (Disused) The end time of the monthly bandwidth subscription
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
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
	delete(f, "AddressIds")
	delete(f, "InternetMaxBandwidthOut")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAddressesBandwidthRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAddressesBandwidthResponseParams struct {
	// The async task ID. You can use the [DescribeTaskResult](https://intl.cloud.tencent.com/document/api/215/36271?from_cn_redirect=1) API to query the task status.
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type ModifyAddressesRenewFlagRequestParams struct {
	// List of unique EIP IDs, for example, eip-xxxx.
	AddressIds []*string `json:"AddressIds,omitnil,omitempty" name:"AddressIds"`

	// Auto-renewal flag.
	// <ul style="margin:0">
	// <li>NOTIFY_AND_AUTO_RENEW: Notify upon expiration and automatically renew.</li>
	// <li>NOTIFY_AND_MANUAL_RENEW: Notify upon expiration but do not automatically renew.</li>
	// <li>DISABLE_NOTIFY_AND_MANUAL_RENEW: Do not notify upon expiration and do not automatically renew.</li>
	// </ul>
	// If this parameter is set to NOTIFY_AND_AUTO_RENEW and the account balance is sufficient, the instance will be automatically renewed on a monthly basis after expiration.
	RenewFlag *string `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`
}

type ModifyAddressesRenewFlagRequest struct {
	*tchttp.BaseRequest
	
	// List of unique EIP IDs, for example, eip-xxxx.
	AddressIds []*string `json:"AddressIds,omitnil,omitempty" name:"AddressIds"`

	// Auto-renewal flag.
	// <ul style="margin:0">
	// <li>NOTIFY_AND_AUTO_RENEW: Notify upon expiration and automatically renew.</li>
	// <li>NOTIFY_AND_MANUAL_RENEW: Notify upon expiration but do not automatically renew.</li>
	// <li>DISABLE_NOTIFY_AND_MANUAL_RENEW: Do not notify upon expiration and do not automatically renew.</li>
	// </ul>
	// If this parameter is set to NOTIFY_AND_AUTO_RENEW and the account balance is sufficient, the instance will be automatically renewed on a monthly basis after expiration.
	RenewFlag *string `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`
}

func (r *ModifyAddressesRenewFlagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAddressesRenewFlagRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AddressIds")
	delete(f, "RenewFlag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAddressesRenewFlagRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAddressesRenewFlagResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAddressesRenewFlagResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAddressesRenewFlagResponseParams `json:"Response"`
}

func (r *ModifyAddressesRenewFlagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAddressesRenewFlagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAssistantCidrRequestParams struct {
	// `VPC` instance `ID`, such as `vpc-6v2ht8q5`.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Array of the secondary CIDR blocks to be added, such as ["10.0.0.0/16", "172.16.0.0/16"]. At least one of `NewCidrBlocks` and `OldCidrBlocks` must be specified.
	NewCidrBlocks []*string `json:"NewCidrBlocks,omitnil,omitempty" name:"NewCidrBlocks"`

	// Array of the secondary CIDR blocks to be deleted, such as ["10.0.0.0/16", "172.16.0.0/16"]. At least one of `NewCidrBlocks` or `OldCidrBlocks` must be specified.
	OldCidrBlocks []*string `json:"OldCidrBlocks,omitnil,omitempty" name:"OldCidrBlocks"`
}

type ModifyAssistantCidrRequest struct {
	*tchttp.BaseRequest
	
	// `VPC` instance `ID`, such as `vpc-6v2ht8q5`.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Array of the secondary CIDR blocks to be added, such as ["10.0.0.0/16", "172.16.0.0/16"]. At least one of `NewCidrBlocks` and `OldCidrBlocks` must be specified.
	NewCidrBlocks []*string `json:"NewCidrBlocks,omitnil,omitempty" name:"NewCidrBlocks"`

	// Array of the secondary CIDR blocks to be deleted, such as ["10.0.0.0/16", "172.16.0.0/16"]. At least one of `NewCidrBlocks` or `OldCidrBlocks` must be specified.
	OldCidrBlocks []*string `json:"OldCidrBlocks,omitnil,omitempty" name:"OldCidrBlocks"`
}

func (r *ModifyAssistantCidrRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAssistantCidrRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcId")
	delete(f, "NewCidrBlocks")
	delete(f, "OldCidrBlocks")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAssistantCidrRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAssistantCidrResponseParams struct {
	// Array of secondary CIDR blocks.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AssistantCidrSet []*AssistantCidr `json:"AssistantCidrSet,omitnil,omitempty" name:"AssistantCidrSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAssistantCidrResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAssistantCidrResponseParams `json:"Response"`
}

func (r *ModifyAssistantCidrResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAssistantCidrResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBandwidthPackageAttributeRequestParams struct {
	// The unique ID of the bandwidth package.
	BandwidthPackageId *string `json:"BandwidthPackageId,omitnil,omitempty" name:"BandwidthPackageId"`

	// The name of the bandwidth package.
	BandwidthPackageName *string `json:"BandwidthPackageName,omitnil,omitempty" name:"BandwidthPackageName"`

	// The billing mode of the bandwidth package. Values: 
	// `TOP5_POSTPAID_BY_MONTH`: Bill by the top 5 bandwidth value of the current month in a postpaid manner
	ChargeType *string `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`
}

type ModifyBandwidthPackageAttributeRequest struct {
	*tchttp.BaseRequest
	
	// The unique ID of the bandwidth package.
	BandwidthPackageId *string `json:"BandwidthPackageId,omitnil,omitempty" name:"BandwidthPackageId"`

	// The name of the bandwidth package.
	BandwidthPackageName *string `json:"BandwidthPackageName,omitnil,omitempty" name:"BandwidthPackageName"`

	// The billing mode of the bandwidth package. Values: 
	// `TOP5_POSTPAID_BY_MONTH`: Bill by the top 5 bandwidth value of the current month in a postpaid manner
	ChargeType *string `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`
}

func (r *ModifyBandwidthPackageAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBandwidthPackageAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BandwidthPackageId")
	delete(f, "BandwidthPackageName")
	delete(f, "ChargeType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyBandwidthPackageAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBandwidthPackageAttributeResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyBandwidthPackageAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyBandwidthPackageAttributeResponseParams `json:"Response"`
}

func (r *ModifyBandwidthPackageAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBandwidthPackageAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCcnAttachedInstancesAttributeRequestParams struct {
	// CCN instance ID in the format of `ccn-f49l6u0z`
	CcnId *string `json:"CcnId,omitnil,omitempty" name:"CcnId"`

	// List of associated network instances
	Instances []*CcnInstance `json:"Instances,omitnil,omitempty" name:"Instances"`
}

type ModifyCcnAttachedInstancesAttributeRequest struct {
	*tchttp.BaseRequest
	
	// CCN instance ID in the format of `ccn-f49l6u0z`
	CcnId *string `json:"CcnId,omitnil,omitempty" name:"CcnId"`

	// List of associated network instances
	Instances []*CcnInstance `json:"Instances,omitnil,omitempty" name:"Instances"`
}

func (r *ModifyCcnAttachedInstancesAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCcnAttachedInstancesAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CcnId")
	delete(f, "Instances")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCcnAttachedInstancesAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCcnAttachedInstancesAttributeResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyCcnAttachedInstancesAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCcnAttachedInstancesAttributeResponseParams `json:"Response"`
}

func (r *ModifyCcnAttachedInstancesAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCcnAttachedInstancesAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCcnAttributeRequestParams struct {
	// The CCN instance ID, such as `ccn-f49l6u0z`.
	CcnId *string `json:"CcnId,omitnil,omitempty" name:"CcnId"`

	// The name of CCN instance. Up to 60 characters allowed. It can contain up to 60 bytes. Either `CcnName` or `CcnDescription` must be specified.
	CcnName *string `json:"CcnName,omitnil,omitempty" name:"CcnName"`

	// The description of CCN instance. Up to 100 characters allowed. It can contain up to 60 bytes. Either `CcnName` or `CcnDescription` must be specified.
	CcnDescription *string `json:"CcnDescription,omitnil,omitempty" name:"CcnDescription"`
}

type ModifyCcnAttributeRequest struct {
	*tchttp.BaseRequest
	
	// The CCN instance ID, such as `ccn-f49l6u0z`.
	CcnId *string `json:"CcnId,omitnil,omitempty" name:"CcnId"`

	// The name of CCN instance. Up to 60 characters allowed. It can contain up to 60 bytes. Either `CcnName` or `CcnDescription` must be specified.
	CcnName *string `json:"CcnName,omitnil,omitempty" name:"CcnName"`

	// The description of CCN instance. Up to 100 characters allowed. It can contain up to 60 bytes. Either `CcnName` or `CcnDescription` must be specified.
	CcnDescription *string `json:"CcnDescription,omitnil,omitempty" name:"CcnDescription"`
}

func (r *ModifyCcnAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCcnAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CcnId")
	delete(f, "CcnName")
	delete(f, "CcnDescription")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCcnAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCcnAttributeResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyCcnAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCcnAttributeResponseParams `json:"Response"`
}

func (r *ModifyCcnAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCcnAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCcnRegionBandwidthLimitsTypeRequestParams struct {
	// CCN instance ID.
	CcnId *string `json:"CcnId,omitnil,omitempty" name:"CcnId"`

	// CCN bandwidth limit type. INTER_REGION_LIMIT: limit between regions. OUTER_REGION_LIMIT: region egress limit.
	BandwidthLimitType *string `json:"BandwidthLimitType,omitnil,omitempty" name:"BandwidthLimitType"`
}

type ModifyCcnRegionBandwidthLimitsTypeRequest struct {
	*tchttp.BaseRequest
	
	// CCN instance ID.
	CcnId *string `json:"CcnId,omitnil,omitempty" name:"CcnId"`

	// CCN bandwidth limit type. INTER_REGION_LIMIT: limit between regions. OUTER_REGION_LIMIT: region egress limit.
	BandwidthLimitType *string `json:"BandwidthLimitType,omitnil,omitempty" name:"BandwidthLimitType"`
}

func (r *ModifyCcnRegionBandwidthLimitsTypeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCcnRegionBandwidthLimitsTypeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CcnId")
	delete(f, "BandwidthLimitType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCcnRegionBandwidthLimitsTypeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCcnRegionBandwidthLimitsTypeResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyCcnRegionBandwidthLimitsTypeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCcnRegionBandwidthLimitsTypeResponseParams `json:"Response"`
}

func (r *ModifyCcnRegionBandwidthLimitsTypeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCcnRegionBandwidthLimitsTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCustomerGatewayAttributeRequestParams struct {
	// The ID of the customer gateway, such as `cgw-2wqq41m9`. You can query the customer gateway by using the [DescribeCustomerGateways](https://intl.cloud.tencent.com/document/api/215/17516?from_cn_redirect=1) API.
	CustomerGatewayId *string `json:"CustomerGatewayId,omitnil,omitempty" name:"CustomerGatewayId"`

	// Customer gateway can be named freely, but the maximum length is 60 characters.
	CustomerGatewayName *string `json:"CustomerGatewayName,omitnil,omitempty" name:"CustomerGatewayName"`
}

type ModifyCustomerGatewayAttributeRequest struct {
	*tchttp.BaseRequest
	
	// The ID of the customer gateway, such as `cgw-2wqq41m9`. You can query the customer gateway by using the [DescribeCustomerGateways](https://intl.cloud.tencent.com/document/api/215/17516?from_cn_redirect=1) API.
	CustomerGatewayId *string `json:"CustomerGatewayId,omitnil,omitempty" name:"CustomerGatewayId"`

	// Customer gateway can be named freely, but the maximum length is 60 characters.
	CustomerGatewayName *string `json:"CustomerGatewayName,omitnil,omitempty" name:"CustomerGatewayName"`
}

func (r *ModifyCustomerGatewayAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCustomerGatewayAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CustomerGatewayId")
	delete(f, "CustomerGatewayName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCustomerGatewayAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCustomerGatewayAttributeResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyCustomerGatewayAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCustomerGatewayAttributeResponseParams `json:"Response"`
}

func (r *ModifyCustomerGatewayAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCustomerGatewayAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDirectConnectGatewayAttributeRequestParams struct {
	// The unique ID of the direct connect gateway, such as `dcg-9o233uri`.
	DirectConnectGatewayId *string `json:"DirectConnectGatewayId,omitnil,omitempty" name:"DirectConnectGatewayId"`

	// The direct connect gateway name. You can enter any name within 60 characters.
	DirectConnectGatewayName *string `json:"DirectConnectGatewayName,omitnil,omitempty" name:"DirectConnectGatewayName"`

	// The CCN route-learning type. Valid values: `BGP` (Automatic learning), `STATIC` (Static, that is, user-configured). You can only modify `CcnRouteType` for a CCN direct connect gateway with BGP enabled.
	CcnRouteType *string `json:"CcnRouteType,omitnil,omitempty" name:"CcnRouteType"`

	// CCN route publishing method. Valid values: `standard` and `exquisite`. You can only modify `ModeType` for a CCN direct connect gateway.
	ModeType *string `json:"ModeType,omitnil,omitempty" name:"ModeType"`
}

type ModifyDirectConnectGatewayAttributeRequest struct {
	*tchttp.BaseRequest
	
	// The unique ID of the direct connect gateway, such as `dcg-9o233uri`.
	DirectConnectGatewayId *string `json:"DirectConnectGatewayId,omitnil,omitempty" name:"DirectConnectGatewayId"`

	// The direct connect gateway name. You can enter any name within 60 characters.
	DirectConnectGatewayName *string `json:"DirectConnectGatewayName,omitnil,omitempty" name:"DirectConnectGatewayName"`

	// The CCN route-learning type. Valid values: `BGP` (Automatic learning), `STATIC` (Static, that is, user-configured). You can only modify `CcnRouteType` for a CCN direct connect gateway with BGP enabled.
	CcnRouteType *string `json:"CcnRouteType,omitnil,omitempty" name:"CcnRouteType"`

	// CCN route publishing method. Valid values: `standard` and `exquisite`. You can only modify `ModeType` for a CCN direct connect gateway.
	ModeType *string `json:"ModeType,omitnil,omitempty" name:"ModeType"`
}

func (r *ModifyDirectConnectGatewayAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDirectConnectGatewayAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DirectConnectGatewayId")
	delete(f, "DirectConnectGatewayName")
	delete(f, "CcnRouteType")
	delete(f, "ModeType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDirectConnectGatewayAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDirectConnectGatewayAttributeResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDirectConnectGatewayAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDirectConnectGatewayAttributeResponseParams `json:"Response"`
}

func (r *ModifyDirectConnectGatewayAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDirectConnectGatewayAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyFlowLogAttributeRequestParams struct {
	// The unique ID of the flow log.
	FlowLogId *string `json:"FlowLogId,omitnil,omitempty" name:"FlowLogId"`

	// The VPC ID or unique ID of the resource. We recommend using the unique ID. This parameter is required unless the attributes of a CCN flow log is modified.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// The name of the flow log instance.
	FlowLogName *string `json:"FlowLogName,omitnil,omitempty" name:"FlowLogName"`

	// The description of the flow log.
	FlowLogDescription *string `json:"FlowLogDescription,omitnil,omitempty" name:"FlowLogDescription"`
}

type ModifyFlowLogAttributeRequest struct {
	*tchttp.BaseRequest
	
	// The unique ID of the flow log.
	FlowLogId *string `json:"FlowLogId,omitnil,omitempty" name:"FlowLogId"`

	// The VPC ID or unique ID of the resource. We recommend using the unique ID. This parameter is required unless the attributes of a CCN flow log is modified.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// The name of the flow log instance.
	FlowLogName *string `json:"FlowLogName,omitnil,omitempty" name:"FlowLogName"`

	// The description of the flow log.
	FlowLogDescription *string `json:"FlowLogDescription,omitnil,omitempty" name:"FlowLogDescription"`
}

func (r *ModifyFlowLogAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyFlowLogAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FlowLogId")
	delete(f, "VpcId")
	delete(f, "FlowLogName")
	delete(f, "FlowLogDescription")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyFlowLogAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyFlowLogAttributeResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyFlowLogAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyFlowLogAttributeResponseParams `json:"Response"`
}

func (r *ModifyFlowLogAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyFlowLogAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyGatewayFlowQosRequestParams struct {
	// Gateway instance ID. Supported types:
	// Direct connect gateway instance, such as `dcg-ltjahce6`;
	// NAT gateway instance, such as `nat-ltjahce6`;
	// VPN gateway instance, such as `vpn-ltjahce6`.
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// Bandwidth limit value in Mbps. Valid values: >0: Set the limit to the specified value. 0: Block all traffic. -1: No bandwidth limit.
	Bandwidth *int64 `json:"Bandwidth,omitnil,omitempty" name:"Bandwidth"`

	// CVM private IP addresses with limited bandwidth.
	IpAddresses []*string `json:"IpAddresses,omitnil,omitempty" name:"IpAddresses"`
}

type ModifyGatewayFlowQosRequest struct {
	*tchttp.BaseRequest
	
	// Gateway instance ID. Supported types:
	// Direct connect gateway instance, such as `dcg-ltjahce6`;
	// NAT gateway instance, such as `nat-ltjahce6`;
	// VPN gateway instance, such as `vpn-ltjahce6`.
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// Bandwidth limit value in Mbps. Valid values: >0: Set the limit to the specified value. 0: Block all traffic. -1: No bandwidth limit.
	Bandwidth *int64 `json:"Bandwidth,omitnil,omitempty" name:"Bandwidth"`

	// CVM private IP addresses with limited bandwidth.
	IpAddresses []*string `json:"IpAddresses,omitnil,omitempty" name:"IpAddresses"`
}

func (r *ModifyGatewayFlowQosRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyGatewayFlowQosRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "Bandwidth")
	delete(f, "IpAddresses")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyGatewayFlowQosRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyGatewayFlowQosResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyGatewayFlowQosResponse struct {
	*tchttp.BaseResponse
	Response *ModifyGatewayFlowQosResponseParams `json:"Response"`
}

func (r *ModifyGatewayFlowQosResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyGatewayFlowQosResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyHaVipAttributeRequestParams struct {
	// The unique `ID` of the `HAVIP`, such as `havip-9o233uri`.
	HaVipId *string `json:"HaVipId,omitnil,omitempty" name:"HaVipId"`

	// `HAVIP` can be named freely, but the maximum length is 60 characters.
	HaVipName *string `json:"HaVipName,omitnil,omitempty" name:"HaVipName"`
}

type ModifyHaVipAttributeRequest struct {
	*tchttp.BaseRequest
	
	// The unique `ID` of the `HAVIP`, such as `havip-9o233uri`.
	HaVipId *string `json:"HaVipId,omitnil,omitempty" name:"HaVipId"`

	// `HAVIP` can be named freely, but the maximum length is 60 characters.
	HaVipName *string `json:"HaVipName,omitnil,omitempty" name:"HaVipName"`
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
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type ModifyIpv6AddressesAttributeRequestParams struct {
	// The `ID` of the ENI instance, such as `eni-m6dyj72l`.
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitnil,omitempty" name:"NetworkInterfaceId"`

	// The information of the specified private `IPv6` addresses.
	Ipv6Addresses []*Ipv6Address `json:"Ipv6Addresses,omitnil,omitempty" name:"Ipv6Addresses"`
}

type ModifyIpv6AddressesAttributeRequest struct {
	*tchttp.BaseRequest
	
	// The `ID` of the ENI instance, such as `eni-m6dyj72l`.
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitnil,omitempty" name:"NetworkInterfaceId"`

	// The information of the specified private `IPv6` addresses.
	Ipv6Addresses []*Ipv6Address `json:"Ipv6Addresses,omitnil,omitempty" name:"Ipv6Addresses"`
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
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type ModifyLocalGatewayRequestParams struct {
	// Local gateway name
	LocalGatewayName *string `json:"LocalGatewayName,omitnil,omitempty" name:"LocalGatewayName"`

	// CDC instance ID
	CdcId *string `json:"CdcId,omitnil,omitempty" name:"CdcId"`

	// Local gateway instance ID
	LocalGatewayId *string `json:"LocalGatewayId,omitnil,omitempty" name:"LocalGatewayId"`

	// VPC instance ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`
}

type ModifyLocalGatewayRequest struct {
	*tchttp.BaseRequest
	
	// Local gateway name
	LocalGatewayName *string `json:"LocalGatewayName,omitnil,omitempty" name:"LocalGatewayName"`

	// CDC instance ID
	CdcId *string `json:"CdcId,omitnil,omitempty" name:"CdcId"`

	// Local gateway instance ID
	LocalGatewayId *string `json:"LocalGatewayId,omitnil,omitempty" name:"LocalGatewayId"`

	// VPC instance ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`
}

func (r *ModifyLocalGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLocalGatewayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LocalGatewayName")
	delete(f, "CdcId")
	delete(f, "LocalGatewayId")
	delete(f, "VpcId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLocalGatewayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLocalGatewayResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyLocalGatewayResponse struct {
	*tchttp.BaseResponse
	Response *ModifyLocalGatewayResponseParams `json:"Response"`
}

func (r *ModifyLocalGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLocalGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNatGatewayAttributeRequestParams struct {
	// The ID of the NAT gateway, such as `nat-df45454`.
	NatGatewayId *string `json:"NatGatewayId,omitnil,omitempty" name:"NatGatewayId"`

	// The NAT gateway name, such as `test_nat`.
	NatGatewayName *string `json:"NatGatewayName,omitnil,omitempty" name:"NatGatewayName"`

	// The maximum outbound bandwidth of the NAT gateway. Unit: Mbps.
	InternetMaxBandwidthOut *uint64 `json:"InternetMaxBandwidthOut,omitnil,omitempty" name:"InternetMaxBandwidthOut"`

	// Whether to modify the security group bound to the NAT gateway
	ModifySecurityGroup *bool `json:"ModifySecurityGroup,omitnil,omitempty" name:"ModifySecurityGroup"`

	// The final security groups bound to the NAT Gateway, such as `['sg-1n232323', 'sg-o4242424']`. An empty list indicates that all the security groups have been deleted.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`
}

type ModifyNatGatewayAttributeRequest struct {
	*tchttp.BaseRequest
	
	// The ID of the NAT gateway, such as `nat-df45454`.
	NatGatewayId *string `json:"NatGatewayId,omitnil,omitempty" name:"NatGatewayId"`

	// The NAT gateway name, such as `test_nat`.
	NatGatewayName *string `json:"NatGatewayName,omitnil,omitempty" name:"NatGatewayName"`

	// The maximum outbound bandwidth of the NAT gateway. Unit: Mbps.
	InternetMaxBandwidthOut *uint64 `json:"InternetMaxBandwidthOut,omitnil,omitempty" name:"InternetMaxBandwidthOut"`

	// Whether to modify the security group bound to the NAT gateway
	ModifySecurityGroup *bool `json:"ModifySecurityGroup,omitnil,omitempty" name:"ModifySecurityGroup"`

	// The final security groups bound to the NAT Gateway, such as `['sg-1n232323', 'sg-o4242424']`. An empty list indicates that all the security groups have been deleted.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`
}

func (r *ModifyNatGatewayAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNatGatewayAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NatGatewayId")
	delete(f, "NatGatewayName")
	delete(f, "InternetMaxBandwidthOut")
	delete(f, "ModifySecurityGroup")
	delete(f, "SecurityGroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyNatGatewayAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNatGatewayAttributeResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyNatGatewayAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyNatGatewayAttributeResponseParams `json:"Response"`
}

func (r *ModifyNatGatewayAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNatGatewayAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNatGatewayDestinationIpPortTranslationNatRuleRequestParams struct {
	// The ID of the NAT gateway, such as `nat-df45454`.
	NatGatewayId *string `json:"NatGatewayId,omitnil,omitempty" name:"NatGatewayId"`

	// The port forwarding rule of the source NAT gateway.
	SourceNatRule *DestinationIpPortTranslationNatRule `json:"SourceNatRule,omitnil,omitempty" name:"SourceNatRule"`

	// The port forwarding rule of the destination NAT gateway.
	DestinationNatRule *DestinationIpPortTranslationNatRule `json:"DestinationNatRule,omitnil,omitempty" name:"DestinationNatRule"`
}

type ModifyNatGatewayDestinationIpPortTranslationNatRuleRequest struct {
	*tchttp.BaseRequest
	
	// The ID of the NAT gateway, such as `nat-df45454`.
	NatGatewayId *string `json:"NatGatewayId,omitnil,omitempty" name:"NatGatewayId"`

	// The port forwarding rule of the source NAT gateway.
	SourceNatRule *DestinationIpPortTranslationNatRule `json:"SourceNatRule,omitnil,omitempty" name:"SourceNatRule"`

	// The port forwarding rule of the destination NAT gateway.
	DestinationNatRule *DestinationIpPortTranslationNatRule `json:"DestinationNatRule,omitnil,omitempty" name:"DestinationNatRule"`
}

func (r *ModifyNatGatewayDestinationIpPortTranslationNatRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNatGatewayDestinationIpPortTranslationNatRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NatGatewayId")
	delete(f, "SourceNatRule")
	delete(f, "DestinationNatRule")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyNatGatewayDestinationIpPortTranslationNatRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNatGatewayDestinationIpPortTranslationNatRuleResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyNatGatewayDestinationIpPortTranslationNatRuleResponse struct {
	*tchttp.BaseResponse
	Response *ModifyNatGatewayDestinationIpPortTranslationNatRuleResponseParams `json:"Response"`
}

func (r *ModifyNatGatewayDestinationIpPortTranslationNatRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNatGatewayDestinationIpPortTranslationNatRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNatGatewaySourceIpTranslationNatRuleRequestParams struct {
	// The ID of the NAT Gateway, such as `nat-df453454`
	NatGatewayId *string `json:"NatGatewayId,omitnil,omitempty" name:"NatGatewayId"`

	// The SNAT forwarding rule of the NAT Gateway
	SourceIpTranslationNatRule *SourceIpTranslationNatRule `json:"SourceIpTranslationNatRule,omitnil,omitempty" name:"SourceIpTranslationNatRule"`
}

type ModifyNatGatewaySourceIpTranslationNatRuleRequest struct {
	*tchttp.BaseRequest
	
	// The ID of the NAT Gateway, such as `nat-df453454`
	NatGatewayId *string `json:"NatGatewayId,omitnil,omitempty" name:"NatGatewayId"`

	// The SNAT forwarding rule of the NAT Gateway
	SourceIpTranslationNatRule *SourceIpTranslationNatRule `json:"SourceIpTranslationNatRule,omitnil,omitempty" name:"SourceIpTranslationNatRule"`
}

func (r *ModifyNatGatewaySourceIpTranslationNatRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNatGatewaySourceIpTranslationNatRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NatGatewayId")
	delete(f, "SourceIpTranslationNatRule")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyNatGatewaySourceIpTranslationNatRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNatGatewaySourceIpTranslationNatRuleResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyNatGatewaySourceIpTranslationNatRuleResponse struct {
	*tchttp.BaseResponse
	Response *ModifyNatGatewaySourceIpTranslationNatRuleResponseParams `json:"Response"`
}

func (r *ModifyNatGatewaySourceIpTranslationNatRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNatGatewaySourceIpTranslationNatRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNetDetectRequestParams struct {
	// The ID of a network detection instance, such as `netd-12345678`.
	NetDetectId *string `json:"NetDetectId,omitnil,omitempty" name:"NetDetectId"`

	// The name of a network detection instance. The maximum length is 60 characters.
	NetDetectName *string `json:"NetDetectName,omitnil,omitempty" name:"NetDetectName"`

	// The array of detection destination IPv4 addresses, which contains at most two IP addresses.
	DetectDestinationIp []*string `json:"DetectDestinationIp,omitnil,omitempty" name:"DetectDestinationIp"`

	// Type of the next hop. Valid values:
	// `VPN`: VPN gateway;
	// `DIRECTCONNECT`: Direct connect gateway;
	// `PEERCONNECTION`: Peering connection;
	// `NAT`: NAT gateway;
	// `NORMAL_CVM`: CVM instance;
	// `CCN`: CCN instance;
	// `NONEXTHOP`: No next hop.
	NextHopType *string `json:"NextHopType,omitnil,omitempty" name:"NextHopType"`

	// ID of the next-hop gateway.
	// If NextHopType is set to VPN, the value of this parameter is the VPN gateway ID, such as vpngw-12345678.
	// `NextHopType` = `DIRECTCONNECT`: Direct connect gateway ID, such as `dcg-12345678`.
	// `NextHopType` = `PEERCONNECTION`: Peering connection ID, such as `pcx-12345678`.
	// `NextHopType` = `NAT`: NAT gateway ID, such as `nat-12345678`.
	// If NextHopType is set to NORMAL_CVM, the value of this parameter is the IPv4 address of the CVM, such as 10.0.0.12.
	// `NextHopType` = `CCN`: CCN instance ID, such as `ccn-12345678`.
	// `NextHopType` = `NONEXTHOP`: No next hop.
	NextHopDestination *string `json:"NextHopDestination,omitnil,omitempty" name:"NextHopDestination"`

	// Network detection description.
	NetDetectDescription *string `json:"NetDetectDescription,omitnil,omitempty" name:"NetDetectDescription"`
}

type ModifyNetDetectRequest struct {
	*tchttp.BaseRequest
	
	// The ID of a network detection instance, such as `netd-12345678`.
	NetDetectId *string `json:"NetDetectId,omitnil,omitempty" name:"NetDetectId"`

	// The name of a network detection instance. The maximum length is 60 characters.
	NetDetectName *string `json:"NetDetectName,omitnil,omitempty" name:"NetDetectName"`

	// The array of detection destination IPv4 addresses, which contains at most two IP addresses.
	DetectDestinationIp []*string `json:"DetectDestinationIp,omitnil,omitempty" name:"DetectDestinationIp"`

	// Type of the next hop. Valid values:
	// `VPN`: VPN gateway;
	// `DIRECTCONNECT`: Direct connect gateway;
	// `PEERCONNECTION`: Peering connection;
	// `NAT`: NAT gateway;
	// `NORMAL_CVM`: CVM instance;
	// `CCN`: CCN instance;
	// `NONEXTHOP`: No next hop.
	NextHopType *string `json:"NextHopType,omitnil,omitempty" name:"NextHopType"`

	// ID of the next-hop gateway.
	// If NextHopType is set to VPN, the value of this parameter is the VPN gateway ID, such as vpngw-12345678.
	// `NextHopType` = `DIRECTCONNECT`: Direct connect gateway ID, such as `dcg-12345678`.
	// `NextHopType` = `PEERCONNECTION`: Peering connection ID, such as `pcx-12345678`.
	// `NextHopType` = `NAT`: NAT gateway ID, such as `nat-12345678`.
	// If NextHopType is set to NORMAL_CVM, the value of this parameter is the IPv4 address of the CVM, such as 10.0.0.12.
	// `NextHopType` = `CCN`: CCN instance ID, such as `ccn-12345678`.
	// `NextHopType` = `NONEXTHOP`: No next hop.
	NextHopDestination *string `json:"NextHopDestination,omitnil,omitempty" name:"NextHopDestination"`

	// Network detection description.
	NetDetectDescription *string `json:"NetDetectDescription,omitnil,omitempty" name:"NetDetectDescription"`
}

func (r *ModifyNetDetectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNetDetectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NetDetectId")
	delete(f, "NetDetectName")
	delete(f, "DetectDestinationIp")
	delete(f, "NextHopType")
	delete(f, "NextHopDestination")
	delete(f, "NetDetectDescription")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyNetDetectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNetDetectResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyNetDetectResponse struct {
	*tchttp.BaseResponse
	Response *ModifyNetDetectResponseParams `json:"Response"`
}

func (r *ModifyNetDetectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNetDetectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNetworkAclAttributeRequestParams struct {
	// Network ACL instance ID, such as `acl-12345678`.
	NetworkAclId *string `json:"NetworkAclId,omitnil,omitempty" name:"NetworkAclId"`

	// Network ACL name, which can contain up to 60 bytes.
	NetworkAclName *string `json:"NetworkAclName,omitnil,omitempty" name:"NetworkAclName"`
}

type ModifyNetworkAclAttributeRequest struct {
	*tchttp.BaseRequest
	
	// Network ACL instance ID, such as `acl-12345678`.
	NetworkAclId *string `json:"NetworkAclId,omitnil,omitempty" name:"NetworkAclId"`

	// Network ACL name, which can contain up to 60 bytes.
	NetworkAclName *string `json:"NetworkAclName,omitnil,omitempty" name:"NetworkAclName"`
}

func (r *ModifyNetworkAclAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNetworkAclAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NetworkAclId")
	delete(f, "NetworkAclName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyNetworkAclAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNetworkAclAttributeResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyNetworkAclAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyNetworkAclAttributeResponseParams `json:"Response"`
}

func (r *ModifyNetworkAclAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNetworkAclAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNetworkAclEntriesRequestParams struct {
	// Network ACL instance ID, such as `acl-12345678`.
	NetworkAclId *string `json:"NetworkAclId,omitnil,omitempty" name:"NetworkAclId"`

	// Network ACL rule set. `NetworkAclEntrySet` and `NetworkAclQuintupleSet` cannot be entered at the same time.
	NetworkAclEntrySet *NetworkAclEntrySet `json:"NetworkAclEntrySet,omitnil,omitempty" name:"NetworkAclEntrySet"`

	// Network ACL quintuple rule set. `NetworkAclEntrySet` and `NetworkAclQuintupleSet` cannot be entered at the same time.
	NetworkAclQuintupleSet *NetworkAclQuintupleEntries `json:"NetworkAclQuintupleSet,omitnil,omitempty" name:"NetworkAclQuintupleSet"`
}

type ModifyNetworkAclEntriesRequest struct {
	*tchttp.BaseRequest
	
	// Network ACL instance ID, such as `acl-12345678`.
	NetworkAclId *string `json:"NetworkAclId,omitnil,omitempty" name:"NetworkAclId"`

	// Network ACL rule set. `NetworkAclEntrySet` and `NetworkAclQuintupleSet` cannot be entered at the same time.
	NetworkAclEntrySet *NetworkAclEntrySet `json:"NetworkAclEntrySet,omitnil,omitempty" name:"NetworkAclEntrySet"`

	// Network ACL quintuple rule set. `NetworkAclEntrySet` and `NetworkAclQuintupleSet` cannot be entered at the same time.
	NetworkAclQuintupleSet *NetworkAclQuintupleEntries `json:"NetworkAclQuintupleSet,omitnil,omitempty" name:"NetworkAclQuintupleSet"`
}

func (r *ModifyNetworkAclEntriesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNetworkAclEntriesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NetworkAclId")
	delete(f, "NetworkAclEntrySet")
	delete(f, "NetworkAclQuintupleSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyNetworkAclEntriesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNetworkAclEntriesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyNetworkAclEntriesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyNetworkAclEntriesResponseParams `json:"Response"`
}

func (r *ModifyNetworkAclEntriesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNetworkAclEntriesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNetworkAclQuintupleEntriesRequestParams struct {
	// Network ACL instance ID, such as `acl-12345678`.
	NetworkAclId *string `json:"NetworkAclId,omitnil,omitempty" name:"NetworkAclId"`

	// Network ACL quintuple rule set.
	NetworkAclQuintupleSet *NetworkAclQuintupleEntries `json:"NetworkAclQuintupleSet,omitnil,omitempty" name:"NetworkAclQuintupleSet"`
}

type ModifyNetworkAclQuintupleEntriesRequest struct {
	*tchttp.BaseRequest
	
	// Network ACL instance ID, such as `acl-12345678`.
	NetworkAclId *string `json:"NetworkAclId,omitnil,omitempty" name:"NetworkAclId"`

	// Network ACL quintuple rule set.
	NetworkAclQuintupleSet *NetworkAclQuintupleEntries `json:"NetworkAclQuintupleSet,omitnil,omitempty" name:"NetworkAclQuintupleSet"`
}

func (r *ModifyNetworkAclQuintupleEntriesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNetworkAclQuintupleEntriesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NetworkAclId")
	delete(f, "NetworkAclQuintupleSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyNetworkAclQuintupleEntriesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNetworkAclQuintupleEntriesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyNetworkAclQuintupleEntriesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyNetworkAclQuintupleEntriesResponseParams `json:"Response"`
}

func (r *ModifyNetworkAclQuintupleEntriesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNetworkAclQuintupleEntriesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNetworkInterfaceAttributeRequestParams struct {
	// The ID of the ENI instance, such as `eni-pxir56ns`.
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitnil,omitempty" name:"NetworkInterfaceId"`

	// The name of the ENI. The maximum length is 60 characters.
	NetworkInterfaceName *string `json:"NetworkInterfaceName,omitnil,omitempty" name:"NetworkInterfaceName"`

	// ENI description can be named freely, but the maximum length is 60 characters.
	NetworkInterfaceDescription *string `json:"NetworkInterfaceDescription,omitnil,omitempty" name:"NetworkInterfaceDescription"`

	// The specified security groups to be bound with, such as ['sg-1dd51d'].
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// Configuration of the ENI trunking mode. Valid values: `Enable` and `Disable`. Default value: `Disable`.
	TrunkingFlag *string `json:"TrunkingFlag,omitnil,omitempty" name:"TrunkingFlag"`
}

type ModifyNetworkInterfaceAttributeRequest struct {
	*tchttp.BaseRequest
	
	// The ID of the ENI instance, such as `eni-pxir56ns`.
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitnil,omitempty" name:"NetworkInterfaceId"`

	// The name of the ENI. The maximum length is 60 characters.
	NetworkInterfaceName *string `json:"NetworkInterfaceName,omitnil,omitempty" name:"NetworkInterfaceName"`

	// ENI description can be named freely, but the maximum length is 60 characters.
	NetworkInterfaceDescription *string `json:"NetworkInterfaceDescription,omitnil,omitempty" name:"NetworkInterfaceDescription"`

	// The specified security groups to be bound with, such as ['sg-1dd51d'].
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// Configuration of the ENI trunking mode. Valid values: `Enable` and `Disable`. Default value: `Disable`.
	TrunkingFlag *string `json:"TrunkingFlag,omitnil,omitempty" name:"TrunkingFlag"`
}

func (r *ModifyNetworkInterfaceAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNetworkInterfaceAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NetworkInterfaceId")
	delete(f, "NetworkInterfaceName")
	delete(f, "NetworkInterfaceDescription")
	delete(f, "SecurityGroupIds")
	delete(f, "TrunkingFlag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyNetworkInterfaceAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNetworkInterfaceAttributeResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyNetworkInterfaceAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyNetworkInterfaceAttributeResponseParams `json:"Response"`
}

func (r *ModifyNetworkInterfaceAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNetworkInterfaceAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPrivateIpAddressesAttributeRequestParams struct {
	// The ID of the ENI instance, such as `eni-m6dyj72l`.
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitnil,omitempty" name:"NetworkInterfaceId"`

	// The specified private IP information.
	PrivateIpAddresses []*PrivateIpAddressSpecification `json:"PrivateIpAddresses,omitnil,omitempty" name:"PrivateIpAddresses"`
}

type ModifyPrivateIpAddressesAttributeRequest struct {
	*tchttp.BaseRequest
	
	// The ID of the ENI instance, such as `eni-m6dyj72l`.
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitnil,omitempty" name:"NetworkInterfaceId"`

	// The specified private IP information.
	PrivateIpAddresses []*PrivateIpAddressSpecification `json:"PrivateIpAddresses,omitnil,omitempty" name:"PrivateIpAddresses"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyPrivateIpAddressesAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPrivateIpAddressesAttributeResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// The route table instance ID, such as `rtb-azd4dt1c`.
	RouteTableId *string `json:"RouteTableId,omitnil,omitempty" name:"RouteTableId"`

	// Route table name.
	RouteTableName *string `json:"RouteTableName,omitnil,omitempty" name:"RouteTableName"`
}

type ModifyRouteTableAttributeRequest struct {
	*tchttp.BaseRequest
	
	// The route table instance ID, such as `rtb-azd4dt1c`.
	RouteTableId *string `json:"RouteTableId,omitnil,omitempty" name:"RouteTableId"`

	// Route table name.
	RouteTableName *string `json:"RouteTableName,omitnil,omitempty" name:"RouteTableName"`
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
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// The security group instance ID, such as `sg-33ocnj9n`. This can be obtained through DescribeSecurityGroups.
	SecurityGroupId *string `json:"SecurityGroupId,omitnil,omitempty" name:"SecurityGroupId"`

	// Security group can be named freely, but cannot exceed 60 characters.
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// The remarks for the security group. The maximum length is 100 characters.
	GroupDescription *string `json:"GroupDescription,omitnil,omitempty" name:"GroupDescription"`
}

type ModifySecurityGroupAttributeRequest struct {
	*tchttp.BaseRequest
	
	// The security group instance ID, such as `sg-33ocnj9n`. This can be obtained through DescribeSecurityGroups.
	SecurityGroupId *string `json:"SecurityGroupId,omitnil,omitempty" name:"SecurityGroupId"`

	// Security group can be named freely, but cannot exceed 60 characters.
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// The remarks for the security group. The maximum length is 100 characters.
	GroupDescription *string `json:"GroupDescription,omitnil,omitempty" name:"GroupDescription"`
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
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// The security group instance ID, such as `sg-33ocnj9n`. This can be obtained through DescribeSecurityGroups.
	SecurityGroupId *string `json:"SecurityGroupId,omitnil,omitempty" name:"SecurityGroupId"`

	// The security group policy set. SecurityGroupPolicySet object must specify new egress and ingress policies at the same time. SecurityGroupPolicy object does not support custom index (PolicyIndex).
	SecurityGroupPolicySet *SecurityGroupPolicySet `json:"SecurityGroupPolicySet,omitnil,omitempty" name:"SecurityGroupPolicySet"`

	// Whether the security group rule is sorted. Default value: False. If it is set to `True`, security group rules will be strictly sorted according to the sequence specified in the `SecurityGroupPolicySet` parameter. Manual entry may cause omission, so we recommend sorting security group rules in the console.
	SortPolicys *bool `json:"SortPolicys,omitnil,omitempty" name:"SortPolicys"`
}

type ModifySecurityGroupPoliciesRequest struct {
	*tchttp.BaseRequest
	
	// The security group instance ID, such as `sg-33ocnj9n`. This can be obtained through DescribeSecurityGroups.
	SecurityGroupId *string `json:"SecurityGroupId,omitnil,omitempty" name:"SecurityGroupId"`

	// The security group policy set. SecurityGroupPolicySet object must specify new egress and ingress policies at the same time. SecurityGroupPolicy object does not support custom index (PolicyIndex).
	SecurityGroupPolicySet *SecurityGroupPolicySet `json:"SecurityGroupPolicySet,omitnil,omitempty" name:"SecurityGroupPolicySet"`

	// Whether the security group rule is sorted. Default value: False. If it is set to `True`, security group rules will be strictly sorted according to the sequence specified in the `SecurityGroupPolicySet` parameter. Manual entry may cause omission, so we recommend sorting security group rules in the console.
	SortPolicys *bool `json:"SortPolicys,omitnil,omitempty" name:"SortPolicys"`
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
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type ModifyServiceTemplateAttributeRequestParams struct {
	// Protocol port template instance ID, such as `ppm-529nwwj8`.
	ServiceTemplateId *string `json:"ServiceTemplateId,omitnil,omitempty" name:"ServiceTemplateId"`

	// Protocol port template name.
	ServiceTemplateName *string `json:"ServiceTemplateName,omitnil,omitempty" name:"ServiceTemplateName"`

	// It supports single port, multiple ports, consecutive ports and all ports. Supported protocols include TCP, UDP, ICMP, and GRE.
	Services []*string `json:"Services,omitnil,omitempty" name:"Services"`

	// Protocol port information with remarks. Supported ports include single port, multiple ports, consecutive ports and other ports. Supported protocols include TCP, UDP, ICMP, and GRE.
	ServicesExtra []*ServicesInfo `json:"ServicesExtra,omitnil,omitempty" name:"ServicesExtra"`
}

type ModifyServiceTemplateAttributeRequest struct {
	*tchttp.BaseRequest
	
	// Protocol port template instance ID, such as `ppm-529nwwj8`.
	ServiceTemplateId *string `json:"ServiceTemplateId,omitnil,omitempty" name:"ServiceTemplateId"`

	// Protocol port template name.
	ServiceTemplateName *string `json:"ServiceTemplateName,omitnil,omitempty" name:"ServiceTemplateName"`

	// It supports single port, multiple ports, consecutive ports and all ports. Supported protocols include TCP, UDP, ICMP, and GRE.
	Services []*string `json:"Services,omitnil,omitempty" name:"Services"`

	// Protocol port information with remarks. Supported ports include single port, multiple ports, consecutive ports and other ports. Supported protocols include TCP, UDP, ICMP, and GRE.
	ServicesExtra []*ServicesInfo `json:"ServicesExtra,omitnil,omitempty" name:"ServicesExtra"`
}

func (r *ModifyServiceTemplateAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyServiceTemplateAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceTemplateId")
	delete(f, "ServiceTemplateName")
	delete(f, "Services")
	delete(f, "ServicesExtra")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyServiceTemplateAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyServiceTemplateAttributeResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyServiceTemplateAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyServiceTemplateAttributeResponseParams `json:"Response"`
}

func (r *ModifyServiceTemplateAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyServiceTemplateAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyServiceTemplateGroupAttributeRequestParams struct {
	// The protocol port template group instance ID, such as `ppmg-ei8hfd9a`.
	ServiceTemplateGroupId *string `json:"ServiceTemplateGroupId,omitnil,omitempty" name:"ServiceTemplateGroupId"`

	// Protocol port template group name.
	ServiceTemplateGroupName *string `json:"ServiceTemplateGroupName,omitnil,omitempty" name:"ServiceTemplateGroupName"`

	// Instance ID of the protocol port template, such as `ppm-4dw6agho`.
	ServiceTemplateIds []*string `json:"ServiceTemplateIds,omitnil,omitempty" name:"ServiceTemplateIds"`
}

type ModifyServiceTemplateGroupAttributeRequest struct {
	*tchttp.BaseRequest
	
	// The protocol port template group instance ID, such as `ppmg-ei8hfd9a`.
	ServiceTemplateGroupId *string `json:"ServiceTemplateGroupId,omitnil,omitempty" name:"ServiceTemplateGroupId"`

	// Protocol port template group name.
	ServiceTemplateGroupName *string `json:"ServiceTemplateGroupName,omitnil,omitempty" name:"ServiceTemplateGroupName"`

	// Instance ID of the protocol port template, such as `ppm-4dw6agho`.
	ServiceTemplateIds []*string `json:"ServiceTemplateIds,omitnil,omitempty" name:"ServiceTemplateIds"`
}

func (r *ModifyServiceTemplateGroupAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyServiceTemplateGroupAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceTemplateGroupId")
	delete(f, "ServiceTemplateGroupName")
	delete(f, "ServiceTemplateIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyServiceTemplateGroupAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyServiceTemplateGroupAttributeResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyServiceTemplateGroupAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyServiceTemplateGroupAttributeResponseParams `json:"Response"`
}

func (r *ModifyServiceTemplateGroupAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyServiceTemplateGroupAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySnapshotPoliciesRequestParams struct {
	// Modify snapshot policies
	SnapshotPoliciesInfo []*BatchModifySnapshotPolicy `json:"SnapshotPoliciesInfo,omitnil,omitempty" name:"SnapshotPoliciesInfo"`
}

type ModifySnapshotPoliciesRequest struct {
	*tchttp.BaseRequest
	
	// Modify snapshot policies
	SnapshotPoliciesInfo []*BatchModifySnapshotPolicy `json:"SnapshotPoliciesInfo,omitnil,omitempty" name:"SnapshotPoliciesInfo"`
}

func (r *ModifySnapshotPoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySnapshotPoliciesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SnapshotPoliciesInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySnapshotPoliciesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySnapshotPoliciesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifySnapshotPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *ModifySnapshotPoliciesResponseParams `json:"Response"`
}

func (r *ModifySnapshotPoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySnapshotPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySubnetAttributeRequestParams struct {
	// Subnet instance ID, such as `subnet-pxir56ns`.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// The subnet name. The maximum length is 60 bytes.
	SubnetName *string `json:"SubnetName,omitnil,omitempty" name:"SubnetName"`

	// Whether the subnet has broadcast enabled.
	EnableBroadcast *string `json:"EnableBroadcast,omitnil,omitempty" name:"EnableBroadcast"`
}

type ModifySubnetAttributeRequest struct {
	*tchttp.BaseRequest
	
	// Subnet instance ID, such as `subnet-pxir56ns`.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// The subnet name. The maximum length is 60 bytes.
	SubnetName *string `json:"SubnetName,omitnil,omitempty" name:"SubnetName"`

	// Whether the subnet has broadcast enabled.
	EnableBroadcast *string `json:"EnableBroadcast,omitnil,omitempty" name:"EnableBroadcast"`
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
	delete(f, "SubnetName")
	delete(f, "EnableBroadcast")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySubnetAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySubnetAttributeResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type ModifyVpcAttributeRequestParams struct {
	// Security group can be named freely, but cannot exceed 60 characters.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// VPC can be named freely, but the maximum length is 60 characters.
	VpcName *string `json:"VpcName,omitnil,omitempty" name:"VpcName"`

	// Whether multicast is enabled. `true`: Enabled. `false`: Off.
	EnableMulticast *string `json:"EnableMulticast,omitnil,omitempty" name:"EnableMulticast"`

	// DNS address. A maximum of 4 addresses is supported. The first one is primary server by default, and the rest are secondary servers.
	DnsServers []*string `json:"DnsServers,omitnil,omitempty" name:"DnsServers"`

	// Domain name
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// Whether to publish the CDC subnet to CCN. `true`: Publish; `false`: u200dDo not publish
	EnableCdcPublish *bool `json:"EnableCdcPublish,omitnil,omitempty" name:"EnableCdcPublish"`
}

type ModifyVpcAttributeRequest struct {
	*tchttp.BaseRequest
	
	// Security group can be named freely, but cannot exceed 60 characters.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// VPC can be named freely, but the maximum length is 60 characters.
	VpcName *string `json:"VpcName,omitnil,omitempty" name:"VpcName"`

	// Whether multicast is enabled. `true`: Enabled. `false`: Off.
	EnableMulticast *string `json:"EnableMulticast,omitnil,omitempty" name:"EnableMulticast"`

	// DNS address. A maximum of 4 addresses is supported. The first one is primary server by default, and the rest are secondary servers.
	DnsServers []*string `json:"DnsServers,omitnil,omitempty" name:"DnsServers"`

	// Domain name
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// Whether to publish the CDC subnet to CCN. `true`: Publish; `false`: u200dDo not publish
	EnableCdcPublish *bool `json:"EnableCdcPublish,omitnil,omitempty" name:"EnableCdcPublish"`
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
	delete(f, "VpcName")
	delete(f, "EnableMulticast")
	delete(f, "DnsServers")
	delete(f, "DomainName")
	delete(f, "EnableCdcPublish")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyVpcAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVpcAttributeResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

// Predefined struct for user
type ModifyVpcEndPointAttributeRequestParams struct {
	// Endpoint ID
	EndPointId *string `json:"EndPointId,omitnil,omitempty" name:"EndPointId"`

	// Endpoint name
	EndPointName *string `json:"EndPointName,omitnil,omitempty" name:"EndPointName"`

	// List of security group IDs
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`
}

type ModifyVpcEndPointAttributeRequest struct {
	*tchttp.BaseRequest
	
	// Endpoint ID
	EndPointId *string `json:"EndPointId,omitnil,omitempty" name:"EndPointId"`

	// Endpoint name
	EndPointName *string `json:"EndPointName,omitnil,omitempty" name:"EndPointName"`

	// List of security group IDs
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`
}

func (r *ModifyVpcEndPointAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVpcEndPointAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EndPointId")
	delete(f, "EndPointName")
	delete(f, "SecurityGroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyVpcEndPointAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVpcEndPointAttributeResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyVpcEndPointAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyVpcEndPointAttributeResponseParams `json:"Response"`
}

func (r *ModifyVpcEndPointAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVpcEndPointAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVpcEndPointServiceAttributeRequestParams struct {
	// Endpoint service ID
	EndPointServiceId *string `json:"EndPointServiceId,omitnil,omitempty" name:"EndPointServiceId"`

	// VPC ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Endpoint service name
	EndPointServiceName *string `json:"EndPointServiceName,omitnil,omitempty" name:"EndPointServiceName"`

	// Whether to automatically accept VPC endpoint connection requests. Valid values: <ui><li>`true`: yes<li>`false`: no</ul>
	AutoAcceptFlag *bool `json:"AutoAcceptFlag,omitnil,omitempty" name:"AutoAcceptFlag"`

	// Real server ID in the format of `lb-xxx`.
	ServiceInstanceId *string `json:"ServiceInstanceId,omitnil,omitempty" name:"ServiceInstanceId"`
}

type ModifyVpcEndPointServiceAttributeRequest struct {
	*tchttp.BaseRequest
	
	// Endpoint service ID
	EndPointServiceId *string `json:"EndPointServiceId,omitnil,omitempty" name:"EndPointServiceId"`

	// VPC ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Endpoint service name
	EndPointServiceName *string `json:"EndPointServiceName,omitnil,omitempty" name:"EndPointServiceName"`

	// Whether to automatically accept VPC endpoint connection requests. Valid values: <ui><li>`true`: yes<li>`false`: no</ul>
	AutoAcceptFlag *bool `json:"AutoAcceptFlag,omitnil,omitempty" name:"AutoAcceptFlag"`

	// Real server ID in the format of `lb-xxx`.
	ServiceInstanceId *string `json:"ServiceInstanceId,omitnil,omitempty" name:"ServiceInstanceId"`
}

func (r *ModifyVpcEndPointServiceAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVpcEndPointServiceAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EndPointServiceId")
	delete(f, "VpcId")
	delete(f, "EndPointServiceName")
	delete(f, "AutoAcceptFlag")
	delete(f, "ServiceInstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyVpcEndPointServiceAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVpcEndPointServiceAttributeResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyVpcEndPointServiceAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyVpcEndPointServiceAttributeResponseParams `json:"Response"`
}

func (r *ModifyVpcEndPointServiceAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVpcEndPointServiceAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVpcEndPointServiceWhiteListRequestParams struct {
	// User UIN
	UserUin *string `json:"UserUin,omitnil,omitempty" name:"UserUin"`

	// Endpoint service ID
	EndPointServiceId *string `json:"EndPointServiceId,omitnil,omitempty" name:"EndPointServiceId"`

	// Allowlist description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type ModifyVpcEndPointServiceWhiteListRequest struct {
	*tchttp.BaseRequest
	
	// User UIN
	UserUin *string `json:"UserUin,omitnil,omitempty" name:"UserUin"`

	// Endpoint service ID
	EndPointServiceId *string `json:"EndPointServiceId,omitnil,omitempty" name:"EndPointServiceId"`

	// Allowlist description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *ModifyVpcEndPointServiceWhiteListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVpcEndPointServiceWhiteListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserUin")
	delete(f, "EndPointServiceId")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyVpcEndPointServiceWhiteListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVpcEndPointServiceWhiteListResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyVpcEndPointServiceWhiteListResponse struct {
	*tchttp.BaseResponse
	Response *ModifyVpcEndPointServiceWhiteListResponseParams `json:"Response"`
}

func (r *ModifyVpcEndPointServiceWhiteListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVpcEndPointServiceWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVpnConnectionAttributeRequestParams struct {
	// The ID of the VPN tunnel instance, such as `vpnx-f49l6u0z`.
	VpnConnectionId *string `json:"VpnConnectionId,omitnil,omitempty" name:"VpnConnectionId"`

	// VPN tunnel can be named freely, but the maximum length is 60 characters.
	VpnConnectionName *string `json:"VpnConnectionName,omitnil,omitempty" name:"VpnConnectionName"`

	// The pre-shared key.
	PreShareKey *string `json:"PreShareKey,omitnil,omitempty" name:"PreShareKey"`

	// SPD policy group. Taking {"10.0.0.5/24":["172.123.10.5/16"]} as an example, 10.0.0.5/24 is the VPC private IP range, and 172.123.10.5/16 is the IDC IP range. The user specifies the IP range in the VPC that can communicate with the IP range in the IDC.
	SecurityPolicyDatabases []*SecurityPolicyDatabase `json:"SecurityPolicyDatabases,omitnil,omitempty" name:"SecurityPolicyDatabases"`

	// IKE (Internet Key Exchange) configuration. IKE comes with a self-protection mechanism. The network security protocol is configured by the user.
	IKEOptionsSpecification *IKEOptionsSpecification `json:"IKEOptionsSpecification,omitnil,omitempty" name:"IKEOptionsSpecification"`

	// IPSec configuration. The IPSec secure session configuration is provided by Tencent Cloud.
	IPSECOptionsSpecification *IPSECOptionsSpecification `json:"IPSECOptionsSpecification,omitnil,omitempty" name:"IPSECOptionsSpecification"`

	// Whether to enable the tunnel health check. The default value is `False`.
	EnableHealthCheck *bool `json:"EnableHealthCheck,omitnil,omitempty" name:"EnableHealthCheck"`

	// Local IP address for the tunnel health check
	HealthCheckLocalIp *string `json:"HealthCheckLocalIp,omitnil,omitempty" name:"HealthCheckLocalIp"`

	// Peer IP address for the tunnel health check
	HealthCheckRemoteIp *string `json:"HealthCheckRemoteIp,omitnil,omitempty" name:"HealthCheckRemoteIp"`

	// Negotiation type. Valid values: `active` (default value), `passive` and `flowTrigger`.
	NegotiationType *string `json:"NegotiationType,omitnil,omitempty" name:"NegotiationType"`

	// Specifies whether to enable DPD. Valid values: `0` (disable) and `1` (enable)
	DpdEnable *int64 `json:"DpdEnable,omitnil,omitempty" name:"DpdEnable"`

	// DPD timeout period. Default: 30; unit: second. If the request is not responded within this period, the peer end is considered not exists. This parameter is valid when the value of `DpdEnable` is 1. 
	DpdTimeout *string `json:"DpdTimeout,omitnil,omitempty" name:"DpdTimeout"`

	// The action after DPD timeout. Valid values: `clear` (disconnect) and `restart` (try again). It’s valid when `DpdEnable` is `1`. 
	DpdAction *string `json:"DpdAction,omitnil,omitempty" name:"DpdAction"`

	// Peer gateway ID. You can update tunnels of V4.0 and later gateways.
	CustomerGatewayId *string `json:"CustomerGatewayId,omitnil,omitempty" name:"CustomerGatewayId"`
}

type ModifyVpnConnectionAttributeRequest struct {
	*tchttp.BaseRequest
	
	// The ID of the VPN tunnel instance, such as `vpnx-f49l6u0z`.
	VpnConnectionId *string `json:"VpnConnectionId,omitnil,omitempty" name:"VpnConnectionId"`

	// VPN tunnel can be named freely, but the maximum length is 60 characters.
	VpnConnectionName *string `json:"VpnConnectionName,omitnil,omitempty" name:"VpnConnectionName"`

	// The pre-shared key.
	PreShareKey *string `json:"PreShareKey,omitnil,omitempty" name:"PreShareKey"`

	// SPD policy group. Taking {"10.0.0.5/24":["172.123.10.5/16"]} as an example, 10.0.0.5/24 is the VPC private IP range, and 172.123.10.5/16 is the IDC IP range. The user specifies the IP range in the VPC that can communicate with the IP range in the IDC.
	SecurityPolicyDatabases []*SecurityPolicyDatabase `json:"SecurityPolicyDatabases,omitnil,omitempty" name:"SecurityPolicyDatabases"`

	// IKE (Internet Key Exchange) configuration. IKE comes with a self-protection mechanism. The network security protocol is configured by the user.
	IKEOptionsSpecification *IKEOptionsSpecification `json:"IKEOptionsSpecification,omitnil,omitempty" name:"IKEOptionsSpecification"`

	// IPSec configuration. The IPSec secure session configuration is provided by Tencent Cloud.
	IPSECOptionsSpecification *IPSECOptionsSpecification `json:"IPSECOptionsSpecification,omitnil,omitempty" name:"IPSECOptionsSpecification"`

	// Whether to enable the tunnel health check. The default value is `False`.
	EnableHealthCheck *bool `json:"EnableHealthCheck,omitnil,omitempty" name:"EnableHealthCheck"`

	// Local IP address for the tunnel health check
	HealthCheckLocalIp *string `json:"HealthCheckLocalIp,omitnil,omitempty" name:"HealthCheckLocalIp"`

	// Peer IP address for the tunnel health check
	HealthCheckRemoteIp *string `json:"HealthCheckRemoteIp,omitnil,omitempty" name:"HealthCheckRemoteIp"`

	// Negotiation type. Valid values: `active` (default value), `passive` and `flowTrigger`.
	NegotiationType *string `json:"NegotiationType,omitnil,omitempty" name:"NegotiationType"`

	// Specifies whether to enable DPD. Valid values: `0` (disable) and `1` (enable)
	DpdEnable *int64 `json:"DpdEnable,omitnil,omitempty" name:"DpdEnable"`

	// DPD timeout period. Default: 30; unit: second. If the request is not responded within this period, the peer end is considered not exists. This parameter is valid when the value of `DpdEnable` is 1. 
	DpdTimeout *string `json:"DpdTimeout,omitnil,omitempty" name:"DpdTimeout"`

	// The action after DPD timeout. Valid values: `clear` (disconnect) and `restart` (try again). It’s valid when `DpdEnable` is `1`. 
	DpdAction *string `json:"DpdAction,omitnil,omitempty" name:"DpdAction"`

	// Peer gateway ID. You can update tunnels of V4.0 and later gateways.
	CustomerGatewayId *string `json:"CustomerGatewayId,omitnil,omitempty" name:"CustomerGatewayId"`
}

func (r *ModifyVpnConnectionAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVpnConnectionAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpnConnectionId")
	delete(f, "VpnConnectionName")
	delete(f, "PreShareKey")
	delete(f, "SecurityPolicyDatabases")
	delete(f, "IKEOptionsSpecification")
	delete(f, "IPSECOptionsSpecification")
	delete(f, "EnableHealthCheck")
	delete(f, "HealthCheckLocalIp")
	delete(f, "HealthCheckRemoteIp")
	delete(f, "NegotiationType")
	delete(f, "DpdEnable")
	delete(f, "DpdTimeout")
	delete(f, "DpdAction")
	delete(f, "CustomerGatewayId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyVpnConnectionAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVpnConnectionAttributeResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyVpnConnectionAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyVpnConnectionAttributeResponseParams `json:"Response"`
}

func (r *ModifyVpnConnectionAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVpnConnectionAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVpnGatewayAttributeRequestParams struct {
	// The ID of the VPN gateway instance.
	VpnGatewayId *string `json:"VpnGatewayId,omitnil,omitempty" name:"VpnGatewayId"`

	// The VPN gateway name. The maximum length is 60 bytes.
	VpnGatewayName *string `json:"VpnGatewayName,omitnil,omitempty" name:"VpnGatewayName"`

	// VPN gateway billing mode. Currently, only the conversion of prepaid (monthly subscription) to postpaid (that is, pay-as-you-go) is supported. That is, the parameters only supports POSTPAID_BY_HOUR.
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`
}

type ModifyVpnGatewayAttributeRequest struct {
	*tchttp.BaseRequest
	
	// The ID of the VPN gateway instance.
	VpnGatewayId *string `json:"VpnGatewayId,omitnil,omitempty" name:"VpnGatewayId"`

	// The VPN gateway name. The maximum length is 60 bytes.
	VpnGatewayName *string `json:"VpnGatewayName,omitnil,omitempty" name:"VpnGatewayName"`

	// VPN gateway billing mode. Currently, only the conversion of prepaid (monthly subscription) to postpaid (that is, pay-as-you-go) is supported. That is, the parameters only supports POSTPAID_BY_HOUR.
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`
}

func (r *ModifyVpnGatewayAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVpnGatewayAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpnGatewayId")
	delete(f, "VpnGatewayName")
	delete(f, "InstanceChargeType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyVpnGatewayAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVpnGatewayAttributeResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyVpnGatewayAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyVpnGatewayAttributeResponseParams `json:"Response"`
}

func (r *ModifyVpnGatewayAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVpnGatewayAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVpnGatewayCcnRoutesRequestParams struct {
	// Instance ID of the VPN gateway
	VpnGatewayId *string `json:"VpnGatewayId,omitnil,omitempty" name:"VpnGatewayId"`

	// The CCN route (IDC IP range) list.
	Routes []*VpngwCcnRoutes `json:"Routes,omitnil,omitempty" name:"Routes"`
}

type ModifyVpnGatewayCcnRoutesRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID of the VPN gateway
	VpnGatewayId *string `json:"VpnGatewayId,omitnil,omitempty" name:"VpnGatewayId"`

	// The CCN route (IDC IP range) list.
	Routes []*VpngwCcnRoutes `json:"Routes,omitnil,omitempty" name:"Routes"`
}

func (r *ModifyVpnGatewayCcnRoutesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVpnGatewayCcnRoutesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpnGatewayId")
	delete(f, "Routes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyVpnGatewayCcnRoutesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVpnGatewayCcnRoutesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyVpnGatewayCcnRoutesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyVpnGatewayCcnRoutesResponseParams `json:"Response"`
}

func (r *ModifyVpnGatewayCcnRoutesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVpnGatewayCcnRoutesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVpnGatewayRoutesRequestParams struct {
	// Instance ID of the VPN gateway
	VpnGatewayId *string `json:"VpnGatewayId,omitnil,omitempty" name:"VpnGatewayId"`

	// Route parameters to modify
	Routes []*VpnGatewayRouteModify `json:"Routes,omitnil,omitempty" name:"Routes"`
}

type ModifyVpnGatewayRoutesRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID of the VPN gateway
	VpnGatewayId *string `json:"VpnGatewayId,omitnil,omitempty" name:"VpnGatewayId"`

	// Route parameters to modify
	Routes []*VpnGatewayRouteModify `json:"Routes,omitnil,omitempty" name:"Routes"`
}

func (r *ModifyVpnGatewayRoutesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVpnGatewayRoutesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpnGatewayId")
	delete(f, "Routes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyVpnGatewayRoutesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVpnGatewayRoutesResponseParams struct {
	// Route information of the VPN gateway
	// Note: this field may return `null`, indicating that no valid value is obtained.
	Routes []*VpnGatewayRoute `json:"Routes,omitnil,omitempty" name:"Routes"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyVpnGatewayRoutesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyVpnGatewayRoutesResponseParams `json:"Response"`
}

func (r *ModifyVpnGatewayRoutesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVpnGatewayRoutesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NatDirectConnectGatewayRoute struct {
	// The `IPv4` `CIDR` of the subnet.
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitnil,omitempty" name:"DestinationCidrBlock"`

	// The type of the next-hop gateway. Supported types:
	// `DIRECTCONNECT`: Direct connect gateway
	GatewayType *string `json:"GatewayType,omitnil,omitempty" name:"GatewayType"`

	// ID of the next-hop gateway
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// The creation time of the route
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// The update time of the route
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type NatGateway struct {
	// NAT gateway ID.
	NatGatewayId *string `json:"NatGatewayId,omitnil,omitempty" name:"NatGatewayId"`

	// NAT gateway name.
	NatGatewayName *string `json:"NatGatewayName,omitnil,omitempty" name:"NatGatewayName"`

	// NAT gateway creation time.
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// The status of the NAT gateway.
	//  `PENDING`: Being created, `DELETING`: Being deleted, `AVAILABLE`: Running, `UPDATING`: Being upgraded,
	// `FAILED`: Failed.
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// The maximum outbound bandwidth of the gateway. Unit: Mbps.
	InternetMaxBandwidthOut *uint64 `json:"InternetMaxBandwidthOut,omitnil,omitempty" name:"InternetMaxBandwidthOut"`

	// The concurrent connections cap of the gateway.
	MaxConcurrentConnection *uint64 `json:"MaxConcurrentConnection,omitnil,omitempty" name:"MaxConcurrentConnection"`

	// The public IP object array of the bound NAT gateway.
	PublicIpAddressSet []*NatGatewayAddress `json:"PublicIpAddressSet,omitnil,omitempty" name:"PublicIpAddressSet"`

	// The NAT gateway status. `AVAILABLE`: Operating, `UNAVAILABLE`: Unavailable, `INSUFFICIENT`: Service suspended due to account overdue.
	NetworkState *string `json:"NetworkState,omitnil,omitempty" name:"NetworkState"`

	// The port forwarding rules of the NAT gateway.
	DestinationIpPortTranslationNatRuleSet []*DestinationIpPortTranslationNatRule `json:"DestinationIpPortTranslationNatRuleSet,omitnil,omitempty" name:"DestinationIpPortTranslationNatRuleSet"`

	// VPC instance ID.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// The availability zone in which the NAT gateway is located.
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// ID of the direct connect gateway bound.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	DirectConnectGatewayIds []*string `json:"DirectConnectGatewayIds,omitnil,omitempty" name:"DirectConnectGatewayIds"`

	// Subnet ID.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Tag key-value pairs.
	TagSet []*Tag `json:"TagSet,omitnil,omitempty" name:"TagSet"`

	// The list of the security groups bound to the NAT Gateway
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	SecurityGroupSet []*string `json:"SecurityGroupSet,omitnil,omitempty" name:"SecurityGroupSet"`

	// SNAT forwarding rule of the NAT gateway.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	SourceIpTranslationNatRuleSet []*SourceIpTranslationNatRule `json:"SourceIpTranslationNatRuleSet,omitnil,omitempty" name:"SourceIpTranslationNatRuleSet"`

	// Whether the NAT gateway is dedicated.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	IsExclusive *bool `json:"IsExclusive,omitnil,omitempty" name:"IsExclusive"`

	// Bandwidth of the gateway cluster where the dedicated NAT Gateway resides. Unit: Mbps. This field does not exist when the `IsExclusive` field is set to `false`.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	ExclusiveGatewayBandwidth *uint64 `json:"ExclusiveGatewayBandwidth,omitnil,omitempty" name:"ExclusiveGatewayBandwidth"`

	// Whether the NAT gateway is blocked. Values: `NORMAL`, `RESTRICTED`
	// Note: This field may return null, indicating that no valid values can be obtained.
	RestrictState *string `json:"RestrictState,omitnil,omitempty" name:"RestrictState"`

	// NAT gateway major version. `1`: Classic, `2`: Standard
	// Note: This field may return null, indicating that no valid values can be obtained.
	NatProductVersion *uint64 `json:"NatProductVersion,omitnil,omitempty" name:"NatProductVersion"`
}

type NatGatewayAddress struct {
	// The unique ID of the Elastic IP (EIP), such as `eip-11112222`.
	AddressId *string `json:"AddressId,omitnil,omitempty" name:"AddressId"`

	// The public IP address, such as `123.121.34.33`.
	PublicIpAddress *string `json:"PublicIpAddress,omitnil,omitempty" name:"PublicIpAddress"`

	// The block status of the resource. `true` indicates the EIP is blocked. `false` indicates that the EIP is not blocked.
	IsBlocked *bool `json:"IsBlocked,omitnil,omitempty" name:"IsBlocked"`
}

type NatGatewayDestinationIpPortTranslationNatRule struct {
	// Network protocol. Available choices: `TCP`, `UDP`.
	IpProtocol *string `json:"IpProtocol,omitnil,omitempty" name:"IpProtocol"`

	// EIP.
	PublicIpAddress *string `json:"PublicIpAddress,omitnil,omitempty" name:"PublicIpAddress"`

	// Public port.
	PublicPort *uint64 `json:"PublicPort,omitnil,omitempty" name:"PublicPort"`

	// Private IP.
	PrivateIpAddress *string `json:"PrivateIpAddress,omitnil,omitempty" name:"PrivateIpAddress"`

	// Private port.
	PrivatePort *uint64 `json:"PrivatePort,omitnil,omitempty" name:"PrivatePort"`

	// NAT gateway forwarding rule description.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// NAT gateway ID.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	NatGatewayId *string `json:"NatGatewayId,omitnil,omitempty" name:"NatGatewayId"`

	// VPC ID.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// The creation time of the NAT gateway forwarding rule.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`
}

type NetDetect struct {
	// The ID of a VPC instance, such as `vpc-12345678`.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// The name of a VPC instance.
	VpcName *string `json:"VpcName,omitnil,omitempty" name:"VpcName"`

	// The ID of a subnet instance, such as subnet-12345678.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// The name of a subnet instance.
	SubnetName *string `json:"SubnetName,omitnil,omitempty" name:"SubnetName"`

	// The ID of a network detection instance, such as netd-12345678.
	NetDetectId *string `json:"NetDetectId,omitnil,omitempty" name:"NetDetectId"`

	// The name of a network detection instance. The maximum length is 60 characters.
	NetDetectName *string `json:"NetDetectName,omitnil,omitempty" name:"NetDetectName"`

	// The array of detection destination IPv4 addresses, which contains at most two IP addresses.
	DetectDestinationIp []*string `json:"DetectDestinationIp,omitnil,omitempty" name:"DetectDestinationIp"`

	// The array of detection source IPv4 addresses automatically allocated by the system. The length is 2.
	DetectSourceIp []*string `json:"DetectSourceIp,omitnil,omitempty" name:"DetectSourceIp"`

	// Type of the next hop. Currently supported types are:
	// VPN: VPN gateway;
	// `DIRECTCONNECT`: Direct connect gateway;
	// `PEERCONNECTION`: Peering connection;
	// `NAT`: NAT gateway;
	// `NORMAL_CVM`: CVM instance;
	// `CCN`: CCN instance;
	// `NONEXTHOP`: No next hop.
	NextHopType *string `json:"NextHopType,omitnil,omitempty" name:"NextHopType"`

	// ID of the next-hop gateway.
	// If NextHopType is set to VPN, the value of this parameter is the VPN gateway ID, such as vpngw-12345678.
	// `NextHopType` = `DIRECTCONNECT`: Direct connect gateway ID, such as `dcg-12345678`.
	// `NextHopType` = `PEERCONNECTION`: Peering connection ID, such as `pcx-12345678`.
	// `NextHopType` = `NAT`: NAT gateway ID, such as `nat-12345678`.
	// If NextHopType is set to NORMAL_CVM, the value of this parameter is the IPv4 address of the CVM, such as 10.0.0.12.
	// `NextHopType` = `CCN`: CCN instance ID, such as `ccn-12345678`.
	// `NextHopType` = `NONEXTHOP`: No next hop.
	NextHopDestination *string `json:"NextHopDestination,omitnil,omitempty" name:"NextHopDestination"`

	// The name of the next-hop gateway.
	// Note: This field may return null, indicating that no valid values can be obtained.
	NextHopName *string `json:"NextHopName,omitnil,omitempty" name:"NextHopName"`

	// Network detection description.
	// Note: This field may return null, indicating that no valid values can be obtained.
	NetDetectDescription *string `json:"NetDetectDescription,omitnil,omitempty" name:"NetDetectDescription"`

	// The creation time.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`
}

type NetDetectIpState struct {
	// The destination IPv4 address of network detection.
	DetectDestinationIp *string `json:"DetectDestinationIp,omitnil,omitempty" name:"DetectDestinationIp"`

	// The detection result.
	// 0: successful;
	// -1: no packet loss occurred during routing;
	// -2: packet loss occurred when outbound traffic is blocked by the ACL;
	// -3: packet loss occurred when inbound traffic is blocked by the ACL;
	// -4: other errors.
	State *int64 `json:"State,omitnil,omitempty" name:"State"`

	// The latency. Unit: ms.
	Delay *uint64 `json:"Delay,omitnil,omitempty" name:"Delay"`

	// The packet loss rate.
	PacketLossRate *uint64 `json:"PacketLossRate,omitnil,omitempty" name:"PacketLossRate"`
}

type NetDetectState struct {
	// The ID of a network detection instance, such as netd-12345678.
	NetDetectId *string `json:"NetDetectId,omitnil,omitempty" name:"NetDetectId"`

	// The array of network detection destination IP verification results.
	NetDetectIpStateSet []*NetDetectIpState `json:"NetDetectIpStateSet,omitnil,omitempty" name:"NetDetectIpStateSet"`
}

type NetworkAcl struct {
	// `ID` of the `VPC` instance.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// `ID` of the network ACL instance.
	NetworkAclId *string `json:"NetworkAclId,omitnil,omitempty" name:"NetworkAclId"`

	// Name of the network ACL. The maximum length is 60 bytes.
	NetworkAclName *string `json:"NetworkAclName,omitnil,omitempty" name:"NetworkAclName"`

	// Creation time.
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// Array of subnets associated with the network ACL.
	SubnetSet []*Subnet `json:"SubnetSet,omitnil,omitempty" name:"SubnetSet"`

	// Inbound rules of the network ACL.
	IngressEntries []*NetworkAclEntry `json:"IngressEntries,omitnil,omitempty" name:"IngressEntries"`

	// Outbound rules of the network ACL.
	EgressEntries []*NetworkAclEntry `json:"EgressEntries,omitnil,omitempty" name:"EgressEntries"`

	// Network ACL type. Valid values: `TRIPLE` and `QUINTUPLE`.
	NetworkAclType *string `json:"NetworkAclType,omitnil,omitempty" name:"NetworkAclType"`

	// Tag key-value pairs
	TagSet []*Tag `json:"TagSet,omitnil,omitempty" name:"TagSet"`
}

type NetworkAclEntry struct {
	// Modification time.
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`

	// Protocol. Valid values: TCP, UDP, ICMP, ALL.
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// Port. Valid values: all, single port, range. When Protocol takes the value `ALL` or `ICMP`, Port cannot be specified.
	Port *string `json:"Port,omitnil,omitempty" name:"Port"`

	// IP range or IP address (mutually exclusive).
	CidrBlock *string `json:"CidrBlock,omitnil,omitempty" name:"CidrBlock"`

	// CIDR block or IPv6 address (mutually exclusive).
	Ipv6CidrBlock *string `json:"Ipv6CidrBlock,omitnil,omitempty" name:"Ipv6CidrBlock"`

	// ACCEPT or DROP.
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`

	// Rule description, which is up to 100 bytes.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type NetworkAclEntrySet struct {
	// Inbound rules.
	Ingress []*NetworkAclEntry `json:"Ingress,omitnil,omitempty" name:"Ingress"`

	// Outbound rules.
	Egress []*NetworkAclEntry `json:"Egress,omitnil,omitempty" name:"Egress"`
}

type NetworkAclQuintupleEntries struct {
	// Network ACL quintuple inbound rule.
	Ingress []*NetworkAclQuintupleEntry `json:"Ingress,omitnil,omitempty" name:"Ingress"`

	// Network ACL quintuple outbound rule.
	Egress []*NetworkAclQuintupleEntry `json:"Egress,omitnil,omitempty" name:"Egress"`
}

type NetworkAclQuintupleEntry struct {
	// Protocol. Valid values: `TCP`, `UDP`, `ICMP`, `ALL`.
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// Description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Source port. Valid values: all, single port, range. When `Protocol` is `ALL` or `ICMP`, `Port` cannot be specified.
	SourcePort *string `json:"SourcePort,omitnil,omitempty" name:"SourcePort"`

	// Source CIDR block.
	SourceCidr *string `json:"SourceCidr,omitnil,omitempty" name:"SourceCidr"`

	// Destination port. Valid values: all, single port, range. When `Protocol` is `ALL` or `ICMP`, `Port` cannot be specified.
	DestinationPort *string `json:"DestinationPort,omitnil,omitempty" name:"DestinationPort"`

	// Destination CIDR block.
	DestinationCidr *string `json:"DestinationCidr,omitnil,omitempty" name:"DestinationCidr"`

	// Action. Valid values: `ACCEPT` and `DROP`.
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`

	// Unique ID of a network ACL entry.
	NetworkAclQuintupleEntryId *string `json:"NetworkAclQuintupleEntryId,omitnil,omitempty" name:"NetworkAclQuintupleEntryId"`

	// Priority. `1` refers to the highest priority.
	Priority *int64 `json:"Priority,omitnil,omitempty" name:"Priority"`

	// Creation time. It’s returned by `DescribeNetworkAclQuintupleEntries`.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Direction of the rule. It’s returned by `DescribeNetworkAclQuintupleEntries`. Valid values: `INGRESS` and `EGRESS`.
	NetworkAclDirection *string `json:"NetworkAclDirection,omitnil,omitempty" name:"NetworkAclDirection"`
}

type NetworkInterface struct {
	// The ID of the ENI instance, such as `eni-f1xjkw1b`.
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitnil,omitempty" name:"NetworkInterfaceId"`

	// ENI Name
	NetworkInterfaceName *string `json:"NetworkInterfaceName,omitnil,omitempty" name:"NetworkInterfaceName"`

	// ENI description.
	NetworkInterfaceDescription *string `json:"NetworkInterfaceDescription,omitnil,omitempty" name:"NetworkInterfaceDescription"`

	// Subnet instance ID.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// VPC instance ID.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Bound security group.
	GroupSet []*string `json:"GroupSet,omitnil,omitempty" name:"GroupSet"`

	// Whether it is the primary ENI.
	Primary *bool `json:"Primary,omitnil,omitempty" name:"Primary"`

	// MAC address
	MacAddress *string `json:"MacAddress,omitnil,omitempty" name:"MacAddress"`

	// ENI status:
	// <li>`PENDING`: Creating</li>
	// <li>`AVAILABLE`: Available</li>
	// <li>`ATTACHING`: Binding</li>
	// <li>`DETACHING`: Unbinding</li>
	// <li>`DELETING`: Deleting</li>
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// Private IP information.
	PrivateIpAddressSet []*PrivateIpAddressSpecification `json:"PrivateIpAddressSet,omitnil,omitempty" name:"PrivateIpAddressSet"`

	// Bound CVM object.
	// Note: This field may return null, indicating no valid value.
	Attachment *NetworkInterfaceAttachment `json:"Attachment,omitnil,omitempty" name:"Attachment"`

	// Availability Zone.
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Creation Time.
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// The `IPv6` address list.
	Ipv6AddressSet []*Ipv6Address `json:"Ipv6AddressSet,omitnil,omitempty" name:"Ipv6AddressSet"`

	// Tag key-value pair.
	TagSet []*Tag `json:"TagSet,omitnil,omitempty" name:"TagSet"`

	// The ENI type. 0: ENI. 1: EVM ENI.
	EniType *uint64 `json:"EniType,omitnil,omitempty" name:"EniType"`

	// Type of the resource bound with an ENI. Valid values: cvm, eks.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Business *string `json:"Business,omitnil,omitempty" name:"Business"`

	// ID of the CDC instance associated with the ENI
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	CdcId *string `json:"CdcId,omitnil,omitempty" name:"CdcId"`

	// ENI type. Valid values: `0` (standard); `1` (extension). Default value: `0`.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	AttachType *uint64 `json:"AttachType,omitnil,omitempty" name:"AttachType"`

	// The ID of resource to retain the ENI primary IP. It’s used as the request parameters for deleting an ENI.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// Service level
	// <li>`DEFAULT`: Default level</lil>
	// <li>`PT`: Gold</li>
	// <li>`AU`: Silver</li>
	// <li>`AG`: Bronze</li>
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	QosLevel *string `json:"QosLevel,omitnil,omitempty" name:"QosLevel"`
}

type NetworkInterfaceAttachment struct {
	// CVM instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// The serial number of ENI in the CVM instance.
	DeviceIndex *uint64 `json:"DeviceIndex,omitnil,omitempty" name:"DeviceIndex"`

	// The account information of the CVM owner.
	InstanceAccountId *string `json:"InstanceAccountId,omitnil,omitempty" name:"InstanceAccountId"`

	// Binding time
	AttachTime *string `json:"AttachTime,omitnil,omitempty" name:"AttachTime"`
}

// Predefined struct for user
type NotifyRoutesRequestParams struct {
	// The unique ID of the route table
	RouteTableId *string `json:"RouteTableId,omitnil,omitempty" name:"RouteTableId"`

	// The unique ID of the route
	RouteItemIds []*string `json:"RouteItemIds,omitnil,omitempty" name:"RouteItemIds"`
}

type NotifyRoutesRequest struct {
	*tchttp.BaseRequest
	
	// The unique ID of the route table
	RouteTableId *string `json:"RouteTableId,omitnil,omitempty" name:"RouteTableId"`

	// The unique ID of the route
	RouteItemIds []*string `json:"RouteItemIds,omitnil,omitempty" name:"RouteItemIds"`
}

func (r *NotifyRoutesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *NotifyRoutesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RouteTableId")
	delete(f, "RouteItemIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "NotifyRoutesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type NotifyRoutesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type NotifyRoutesResponse struct {
	*tchttp.BaseResponse
	Response *NotifyRoutesResponseParams `json:"Response"`
}

func (r *NotifyRoutesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *NotifyRoutesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Price struct {
	// Instance price.
	InstancePrice *ItemPrice `json:"InstancePrice,omitnil,omitempty" name:"InstancePrice"`

	// Bandwidth price
	BandwidthPrice *ItemPrice `json:"BandwidthPrice,omitnil,omitempty" name:"BandwidthPrice"`
}

type PrivateIpAddressSpecification struct {
	// Private IP address.
	PrivateIpAddress *string `json:"PrivateIpAddress,omitnil,omitempty" name:"PrivateIpAddress"`

	// Whether it is a primary IP.
	Primary *bool `json:"Primary,omitnil,omitempty" name:"Primary"`

	// Public IP address.
	PublicIpAddress *string `json:"PublicIpAddress,omitnil,omitempty" name:"PublicIpAddress"`

	// EIP instance ID, such as `eip-11112222`.
	AddressId *string `json:"AddressId,omitnil,omitempty" name:"AddressId"`

	// Private IP description.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Whether the public IP is blocked.
	IsWanIpBlocked *bool `json:"IsWanIpBlocked,omitnil,omitempty" name:"IsWanIpBlocked"`

	// IP status:
	// PENDING: Creating
	// MIGRATING: Migrating
	// DELETING: Deleting
	// AVAILABLE: Available
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// IP u200dservice level. Values: PT` u200d(Gold), `AU` u200d(Silver), `AG `(Bronze) and DEFAULT` (Default).
	QosLevel *string `json:"QosLevel,omitnil,omitempty" name:"QosLevel"`
}

type Quota struct {
	// Quota name. Value range:<br><li>`TOTAL_EIP_QUOTA`:EIP quota under the user's current region<br><li>`DAILY_EIP_APPLY`: Number of EIP applications submitted daily under the user's current region<br><li>`DAILY_PUBLIC_IP_ASSIGN`: Number of public IP reassignments under the user's current region.
	QuotaId *string `json:"QuotaId,omitnil,omitempty" name:"QuotaId"`

	// Current count
	QuotaCurrent *int64 `json:"QuotaCurrent,omitnil,omitempty" name:"QuotaCurrent"`

	// Quota
	QuotaLimit *int64 `json:"QuotaLimit,omitnil,omitempty" name:"QuotaLimit"`
}

type ReferredSecurityGroup struct {
	// Security group instance ID.
	SecurityGroupId *string `json:"SecurityGroupId,omitnil,omitempty" name:"SecurityGroupId"`

	// IDs of all referred security group instances.
	ReferredSecurityGroupIds []*string `json:"ReferredSecurityGroupIds,omitnil,omitempty" name:"ReferredSecurityGroupIds"`
}

// Predefined struct for user
type RefreshDirectConnectGatewayRouteToNatGatewayRequestParams struct {
	// VPC ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// The NAT gateway ID.
	NatGatewayId *string `json:"NatGatewayId,omitnil,omitempty" name:"NatGatewayId"`

	// Whether it is pre-refresh. Valid values: `True` (yes) and `False` (no)
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`
}

type RefreshDirectConnectGatewayRouteToNatGatewayRequest struct {
	*tchttp.BaseRequest
	
	// VPC ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// The NAT gateway ID.
	NatGatewayId *string `json:"NatGatewayId,omitnil,omitempty" name:"NatGatewayId"`

	// Whether it is pre-refresh. Valid values: `True` (yes) and `False` (no)
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`
}

func (r *RefreshDirectConnectGatewayRouteToNatGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RefreshDirectConnectGatewayRouteToNatGatewayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcId")
	delete(f, "NatGatewayId")
	delete(f, "DryRun")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RefreshDirectConnectGatewayRouteToNatGatewayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RefreshDirectConnectGatewayRouteToNatGatewayResponseParams struct {
	// IDC subnet information
	DirectConnectSubnetSet []*DirectConnectSubnet `json:"DirectConnectSubnetSet,omitnil,omitempty" name:"DirectConnectSubnetSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RefreshDirectConnectGatewayRouteToNatGatewayResponse struct {
	*tchttp.BaseResponse
	Response *RefreshDirectConnectGatewayRouteToNatGatewayResponseParams `json:"Response"`
}

func (r *RefreshDirectConnectGatewayRouteToNatGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RefreshDirectConnectGatewayRouteToNatGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RejectAttachCcnInstancesRequestParams struct {
	// The CCN instance ID, such as `ccn-f49l6u0z`.
	CcnId *string `json:"CcnId,omitnil,omitempty" name:"CcnId"`

	// The list of instances whose association is rejected.
	Instances []*CcnInstance `json:"Instances,omitnil,omitempty" name:"Instances"`
}

type RejectAttachCcnInstancesRequest struct {
	*tchttp.BaseRequest
	
	// The CCN instance ID, such as `ccn-f49l6u0z`.
	CcnId *string `json:"CcnId,omitnil,omitempty" name:"CcnId"`

	// The list of instances whose association is rejected.
	Instances []*CcnInstance `json:"Instances,omitnil,omitempty" name:"Instances"`
}

func (r *RejectAttachCcnInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RejectAttachCcnInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CcnId")
	delete(f, "Instances")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RejectAttachCcnInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RejectAttachCcnInstancesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RejectAttachCcnInstancesResponse struct {
	*tchttp.BaseResponse
	Response *RejectAttachCcnInstancesResponseParams `json:"Response"`
}

func (r *RejectAttachCcnInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RejectAttachCcnInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReleaseAddressesRequestParams struct {
	// The unique ID list of the EIP. The unique ID of an EIP is as follows: `eip-11112222`.
	AddressIds []*string `json:"AddressIds,omitnil,omitempty" name:"AddressIds"`
}

type ReleaseAddressesRequest struct {
	*tchttp.BaseRequest
	
	// The unique ID list of the EIP. The unique ID of an EIP is as follows: `eip-11112222`.
	AddressIds []*string `json:"AddressIds,omitnil,omitempty" name:"AddressIds"`
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
	delete(f, "AddressIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReleaseAddressesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReleaseAddressesResponseParams struct {
	// The async task ID. You can use the [DescribeTaskResult](https://intl.cloud.tencent.com/document/api/215/36271?from_cn_redirect=1) API to query the task status.
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type RemoveBandwidthPackageResourcesRequestParams struct {
	// The unique ID of the bandwidth package, such as `bwp-xxxx`.
	BandwidthPackageId *string `json:"BandwidthPackageId,omitnil,omitempty" name:"BandwidthPackageId"`

	// The resource type. Valid values: `Address` and `LoadBalance`.
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// The resource IP, such as `eip-xxxx` and `lb-xxxx`.
	ResourceIds []*string `json:"ResourceIds,omitnil,omitempty" name:"ResourceIds"`
}

type RemoveBandwidthPackageResourcesRequest struct {
	*tchttp.BaseRequest
	
	// The unique ID of the bandwidth package, such as `bwp-xxxx`.
	BandwidthPackageId *string `json:"BandwidthPackageId,omitnil,omitempty" name:"BandwidthPackageId"`

	// The resource type. Valid values: `Address` and `LoadBalance`.
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// The resource IP, such as `eip-xxxx` and `lb-xxxx`.
	ResourceIds []*string `json:"ResourceIds,omitnil,omitempty" name:"ResourceIds"`
}

func (r *RemoveBandwidthPackageResourcesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveBandwidthPackageResourcesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BandwidthPackageId")
	delete(f, "ResourceType")
	delete(f, "ResourceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RemoveBandwidthPackageResourcesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveBandwidthPackageResourcesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RemoveBandwidthPackageResourcesResponse struct {
	*tchttp.BaseResponse
	Response *RemoveBandwidthPackageResourcesResponseParams `json:"Response"`
}

func (r *RemoveBandwidthPackageResourcesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveBandwidthPackageResourcesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RenewVpnGatewayRequestParams struct {
	// The ID of the VPN gateway instance.
	VpnGatewayId *string `json:"VpnGatewayId,omitnil,omitempty" name:"VpnGatewayId"`

	// Billing Methods
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitnil,omitempty" name:"InstanceChargePrepaid"`
}

type RenewVpnGatewayRequest struct {
	*tchttp.BaseRequest
	
	// The ID of the VPN gateway instance.
	VpnGatewayId *string `json:"VpnGatewayId,omitnil,omitempty" name:"VpnGatewayId"`

	// Billing Methods
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitnil,omitempty" name:"InstanceChargePrepaid"`
}

func (r *RenewVpnGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewVpnGatewayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpnGatewayId")
	delete(f, "InstanceChargePrepaid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RenewVpnGatewayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RenewVpnGatewayResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RenewVpnGatewayResponse struct {
	*tchttp.BaseResponse
	Response *RenewVpnGatewayResponseParams `json:"Response"`
}

func (r *RenewVpnGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewVpnGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReplaceDirectConnectGatewayCcnRoutesRequestParams struct {
	// The ID of the Direct Connect gateway, such as `dcg-prpqlmg1`
	DirectConnectGatewayId *string `json:"DirectConnectGatewayId,omitnil,omitempty" name:"DirectConnectGatewayId"`

	// The list of IDC IP range that must be connected
	Routes []*DirectConnectGatewayCcnRoute `json:"Routes,omitnil,omitempty" name:"Routes"`
}

type ReplaceDirectConnectGatewayCcnRoutesRequest struct {
	*tchttp.BaseRequest
	
	// The ID of the Direct Connect gateway, such as `dcg-prpqlmg1`
	DirectConnectGatewayId *string `json:"DirectConnectGatewayId,omitnil,omitempty" name:"DirectConnectGatewayId"`

	// The list of IDC IP range that must be connected
	Routes []*DirectConnectGatewayCcnRoute `json:"Routes,omitnil,omitempty" name:"Routes"`
}

func (r *ReplaceDirectConnectGatewayCcnRoutesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReplaceDirectConnectGatewayCcnRoutesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DirectConnectGatewayId")
	delete(f, "Routes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReplaceDirectConnectGatewayCcnRoutesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReplaceDirectConnectGatewayCcnRoutesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ReplaceDirectConnectGatewayCcnRoutesResponse struct {
	*tchttp.BaseResponse
	Response *ReplaceDirectConnectGatewayCcnRoutesResponseParams `json:"Response"`
}

func (r *ReplaceDirectConnectGatewayCcnRoutesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReplaceDirectConnectGatewayCcnRoutesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReplaceRouteTableAssociationRequestParams struct {
	// Subnet instance ID, such as `subnet-3x5lf5q0`. This can be queried using the DescribeSubnets API.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// The route table instance ID, such as `rtb-azd4dt1c`.
	RouteTableId *string `json:"RouteTableId,omitnil,omitempty" name:"RouteTableId"`
}

type ReplaceRouteTableAssociationRequest struct {
	*tchttp.BaseRequest
	
	// Subnet instance ID, such as `subnet-3x5lf5q0`. This can be queried using the DescribeSubnets API.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// The route table instance ID, such as `rtb-azd4dt1c`.
	RouteTableId *string `json:"RouteTableId,omitnil,omitempty" name:"RouteTableId"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReplaceRouteTableAssociationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReplaceRouteTableAssociationResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// The route table instance ID, such as `rtb-azd4dt1c`.
	RouteTableId *string `json:"RouteTableId,omitnil,omitempty" name:"RouteTableId"`

	// Routing policy object. The routing policy ID (RouteId) must be specified.
	Routes []*Route `json:"Routes,omitnil,omitempty" name:"Routes"`
}

type ReplaceRoutesRequest struct {
	*tchttp.BaseRequest
	
	// The route table instance ID, such as `rtb-azd4dt1c`.
	RouteTableId *string `json:"RouteTableId,omitnil,omitempty" name:"RouteTableId"`

	// Routing policy object. The routing policy ID (RouteId) must be specified.
	Routes []*Route `json:"Routes,omitnil,omitempty" name:"Routes"`
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
	// Old routing policy
	OldRouteSet []*Route `json:"OldRouteSet,omitnil,omitempty" name:"OldRouteSet"`

	// New routing policy
	NewRouteSet []*Route `json:"NewRouteSet,omitnil,omitempty" name:"NewRouteSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type ReplaceSecurityGroupPoliciesRequestParams struct {
	// The security group instance ID, such as `sg-33ocnj9n`. This can be obtained through the `DescribeSecurityGroups` API.
	SecurityGroupId *string `json:"SecurityGroupId,omitnil,omitempty" name:"SecurityGroupId"`

	// Security group policy set object.
	SecurityGroupPolicySet *SecurityGroupPolicySet `json:"SecurityGroupPolicySet,omitnil,omitempty" name:"SecurityGroupPolicySet"`

	// (Optional) The old policy set of the security group, which is used for log records.
	OriginalSecurityGroupPolicySet *SecurityGroupPolicySet `json:"OriginalSecurityGroupPolicySet,omitnil,omitempty" name:"OriginalSecurityGroupPolicySet"`
}

type ReplaceSecurityGroupPoliciesRequest struct {
	*tchttp.BaseRequest
	
	// The security group instance ID, such as `sg-33ocnj9n`. This can be obtained through the `DescribeSecurityGroups` API.
	SecurityGroupId *string `json:"SecurityGroupId,omitnil,omitempty" name:"SecurityGroupId"`

	// Security group policy set object.
	SecurityGroupPolicySet *SecurityGroupPolicySet `json:"SecurityGroupPolicySet,omitnil,omitempty" name:"SecurityGroupPolicySet"`

	// (Optional) The old policy set of the security group, which is used for log records.
	OriginalSecurityGroupPolicySet *SecurityGroupPolicySet `json:"OriginalSecurityGroupPolicySet,omitnil,omitempty" name:"OriginalSecurityGroupPolicySet"`
}

func (r *ReplaceSecurityGroupPoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReplaceSecurityGroupPoliciesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecurityGroupId")
	delete(f, "SecurityGroupPolicySet")
	delete(f, "OriginalSecurityGroupPolicySet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReplaceSecurityGroupPoliciesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReplaceSecurityGroupPoliciesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ReplaceSecurityGroupPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *ReplaceSecurityGroupPoliciesResponseParams `json:"Response"`
}

func (r *ReplaceSecurityGroupPoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReplaceSecurityGroupPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReplaceSecurityGroupPolicyRequestParams struct {
	// The security group instance ID, such as `sg-33ocnj9n`. This can be obtained through DescribeSecurityGroups.
	SecurityGroupId *string `json:"SecurityGroupId,omitnil,omitempty" name:"SecurityGroupId"`

	// Security group policy set object.
	SecurityGroupPolicySet *SecurityGroupPolicySet `json:"SecurityGroupPolicySet,omitnil,omitempty" name:"SecurityGroupPolicySet"`

	// (Optional) The old policy set of the security group, which is used for log records.
	OriginalSecurityGroupPolicySet *SecurityGroupPolicySet `json:"OriginalSecurityGroupPolicySet,omitnil,omitempty" name:"OriginalSecurityGroupPolicySet"`
}

type ReplaceSecurityGroupPolicyRequest struct {
	*tchttp.BaseRequest
	
	// The security group instance ID, such as `sg-33ocnj9n`. This can be obtained through DescribeSecurityGroups.
	SecurityGroupId *string `json:"SecurityGroupId,omitnil,omitempty" name:"SecurityGroupId"`

	// Security group policy set object.
	SecurityGroupPolicySet *SecurityGroupPolicySet `json:"SecurityGroupPolicySet,omitnil,omitempty" name:"SecurityGroupPolicySet"`

	// (Optional) The old policy set of the security group, which is used for log records.
	OriginalSecurityGroupPolicySet *SecurityGroupPolicySet `json:"OriginalSecurityGroupPolicySet,omitnil,omitempty" name:"OriginalSecurityGroupPolicySet"`
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
	delete(f, "OriginalSecurityGroupPolicySet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReplaceSecurityGroupPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReplaceSecurityGroupPolicyResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type ResetAttachCcnInstancesRequestParams struct {
	// The CCN instance ID, such as `ccn-f49l6u0z`.
	CcnId *string `json:"CcnId,omitnil,omitempty" name:"CcnId"`

	// The UIN (root account) to which the CCN belongs.
	CcnUin *string `json:"CcnUin,omitnil,omitempty" name:"CcnUin"`

	// The list of network instances that re-apply for association.
	Instances []*CcnInstance `json:"Instances,omitnil,omitempty" name:"Instances"`
}

type ResetAttachCcnInstancesRequest struct {
	*tchttp.BaseRequest
	
	// The CCN instance ID, such as `ccn-f49l6u0z`.
	CcnId *string `json:"CcnId,omitnil,omitempty" name:"CcnId"`

	// The UIN (root account) to which the CCN belongs.
	CcnUin *string `json:"CcnUin,omitnil,omitempty" name:"CcnUin"`

	// The list of network instances that re-apply for association.
	Instances []*CcnInstance `json:"Instances,omitnil,omitempty" name:"Instances"`
}

func (r *ResetAttachCcnInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetAttachCcnInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CcnId")
	delete(f, "CcnUin")
	delete(f, "Instances")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetAttachCcnInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetAttachCcnInstancesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ResetAttachCcnInstancesResponse struct {
	*tchttp.BaseResponse
	Response *ResetAttachCcnInstancesResponseParams `json:"Response"`
}

func (r *ResetAttachCcnInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetAttachCcnInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetNatGatewayConnectionRequestParams struct {
	// The NAT gateway ID.
	NatGatewayId *string `json:"NatGatewayId,omitnil,omitempty" name:"NatGatewayId"`

	// Concurrent connections cap of the NAT gateway, such as 1000000, 3000000, 10000000.
	MaxConcurrentConnection *uint64 `json:"MaxConcurrentConnection,omitnil,omitempty" name:"MaxConcurrentConnection"`
}

type ResetNatGatewayConnectionRequest struct {
	*tchttp.BaseRequest
	
	// The NAT gateway ID.
	NatGatewayId *string `json:"NatGatewayId,omitnil,omitempty" name:"NatGatewayId"`

	// Concurrent connections cap of the NAT gateway, such as 1000000, 3000000, 10000000.
	MaxConcurrentConnection *uint64 `json:"MaxConcurrentConnection,omitnil,omitempty" name:"MaxConcurrentConnection"`
}

func (r *ResetNatGatewayConnectionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetNatGatewayConnectionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NatGatewayId")
	delete(f, "MaxConcurrentConnection")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetNatGatewayConnectionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetNatGatewayConnectionResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ResetNatGatewayConnectionResponse struct {
	*tchttp.BaseResponse
	Response *ResetNatGatewayConnectionResponseParams `json:"Response"`
}

func (r *ResetNatGatewayConnectionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetNatGatewayConnectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetRoutesRequestParams struct {
	// The route table instance ID, such as `rtb-azd4dt1c`.
	RouteTableId *string `json:"RouteTableId,omitnil,omitempty" name:"RouteTableId"`

	// The route table name. The maximum length is 60 characters.
	RouteTableName *string `json:"RouteTableName,omitnil,omitempty" name:"RouteTableName"`

	// Routing policy.
	Routes []*Route `json:"Routes,omitnil,omitempty" name:"Routes"`
}

type ResetRoutesRequest struct {
	*tchttp.BaseRequest
	
	// The route table instance ID, such as `rtb-azd4dt1c`.
	RouteTableId *string `json:"RouteTableId,omitnil,omitempty" name:"RouteTableId"`

	// The route table name. The maximum length is 60 characters.
	RouteTableName *string `json:"RouteTableName,omitnil,omitempty" name:"RouteTableName"`

	// Routing policy.
	Routes []*Route `json:"Routes,omitnil,omitempty" name:"Routes"`
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
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

// Predefined struct for user
type ResetVpnConnectionRequestParams struct {
	// The ID of the VPN gateway instance.
	VpnGatewayId *string `json:"VpnGatewayId,omitnil,omitempty" name:"VpnGatewayId"`

	// The ID of the VPN tunnel instance, such as `vpnx-f49l6u0z`.
	VpnConnectionId *string `json:"VpnConnectionId,omitnil,omitempty" name:"VpnConnectionId"`
}

type ResetVpnConnectionRequest struct {
	*tchttp.BaseRequest
	
	// The ID of the VPN gateway instance.
	VpnGatewayId *string `json:"VpnGatewayId,omitnil,omitempty" name:"VpnGatewayId"`

	// The ID of the VPN tunnel instance, such as `vpnx-f49l6u0z`.
	VpnConnectionId *string `json:"VpnConnectionId,omitnil,omitempty" name:"VpnConnectionId"`
}

func (r *ResetVpnConnectionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetVpnConnectionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpnGatewayId")
	delete(f, "VpnConnectionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetVpnConnectionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetVpnConnectionResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ResetVpnConnectionResponse struct {
	*tchttp.BaseResponse
	Response *ResetVpnConnectionResponseParams `json:"Response"`
}

func (r *ResetVpnConnectionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetVpnConnectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetVpnGatewayInternetMaxBandwidthRequestParams struct {
	// The ID of the VPN gateway instance.
	VpnGatewayId *string `json:"VpnGatewayId,omitnil,omitempty" name:"VpnGatewayId"`

	// The new bandwidth cap in Mbps. Values: `5`, `10`, `20`, `50`, `100`, `200`, `500` and `1000`. The adjustment of the VPN gateway bandwidth is limited to [5,100] Mbps and [200,1000] Mbps.
	InternetMaxBandwidthOut *uint64 `json:"InternetMaxBandwidthOut,omitnil,omitempty" name:"InternetMaxBandwidthOut"`
}

type ResetVpnGatewayInternetMaxBandwidthRequest struct {
	*tchttp.BaseRequest
	
	// The ID of the VPN gateway instance.
	VpnGatewayId *string `json:"VpnGatewayId,omitnil,omitempty" name:"VpnGatewayId"`

	// The new bandwidth cap in Mbps. Values: `5`, `10`, `20`, `50`, `100`, `200`, `500` and `1000`. The adjustment of the VPN gateway bandwidth is limited to [5,100] Mbps and [200,1000] Mbps.
	InternetMaxBandwidthOut *uint64 `json:"InternetMaxBandwidthOut,omitnil,omitempty" name:"InternetMaxBandwidthOut"`
}

func (r *ResetVpnGatewayInternetMaxBandwidthRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetVpnGatewayInternetMaxBandwidthRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpnGatewayId")
	delete(f, "InternetMaxBandwidthOut")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetVpnGatewayInternetMaxBandwidthRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetVpnGatewayInternetMaxBandwidthResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ResetVpnGatewayInternetMaxBandwidthResponse struct {
	*tchttp.BaseResponse
	Response *ResetVpnGatewayInternetMaxBandwidthResponseParams `json:"Response"`
}

func (r *ResetVpnGatewayInternetMaxBandwidthResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetVpnGatewayInternetMaxBandwidthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Resource struct {
	// The bandwidth package resource type, including `Address`, and `LoadBalance`
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// The bandwidth package ID, such as `eip-xxxx` and `lb-xxxx`.
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// The bandwidth package resource IP.
	AddressIp *string `json:"AddressIp,omitnil,omitempty" name:"AddressIp"`
}

type ResourceDashboard struct {
	// VPC instance ID, such as `vpc-bq4bzxpj`.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Subnet instance ID, such as subnet-bthucmmy.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Classiclink.
	Classiclink *uint64 `json:"Classiclink,omitnil,omitempty" name:"Classiclink"`

	// Direct Connect gateway.
	Dcg *uint64 `json:"Dcg,omitnil,omitempty" name:"Dcg"`

	// Peering connection.
	Pcx *uint64 `json:"Pcx,omitnil,omitempty" name:"Pcx"`

	// Total number of used IPs except for CVM IP, EIP and network probe IP. The three IP types will be independently counted.
	Ip *uint64 `json:"Ip,omitnil,omitempty" name:"Ip"`

	// NAT gateway.
	Nat *uint64 `json:"Nat,omitnil,omitempty" name:"Nat"`

	// VPN gateway.
	Vpngw *uint64 `json:"Vpngw,omitnil,omitempty" name:"Vpngw"`

	// Flow log.
	FlowLog *uint64 `json:"FlowLog,omitnil,omitempty" name:"FlowLog"`

	// Network probing.
	NetworkDetect *uint64 `json:"NetworkDetect,omitnil,omitempty" name:"NetworkDetect"`

	// Network ACL.
	NetworkACL *uint64 `json:"NetworkACL,omitnil,omitempty" name:"NetworkACL"`

	// Cloud Virtual Machine.
	CVM *uint64 `json:"CVM,omitnil,omitempty" name:"CVM"`

	// Load balancer.
	LB *uint64 `json:"LB,omitnil,omitempty" name:"LB"`

	// Relational database.
	CDB *uint64 `json:"CDB,omitnil,omitempty" name:"CDB"`

	// TencentDB for Memcached.
	Cmem *uint64 `json:"Cmem,omitnil,omitempty" name:"Cmem"`

	// Cloud time series database.
	CTSDB *uint64 `json:"CTSDB,omitnil,omitempty" name:"CTSDB"`

	// TencentDB for MariaDB (TDSQL).
	MariaDB *uint64 `json:"MariaDB,omitnil,omitempty" name:"MariaDB"`

	// TencentDB for SQL Server.
	SQLServer *uint64 `json:"SQLServer,omitnil,omitempty" name:"SQLServer"`

	// TencentDB for PostgreSQL.
	Postgres *uint64 `json:"Postgres,omitnil,omitempty" name:"Postgres"`

	// Network attached storage.
	NAS *uint64 `json:"NAS,omitnil,omitempty" name:"NAS"`

	// Snova data warehouse.
	Greenplumn *uint64 `json:"Greenplumn,omitnil,omitempty" name:"Greenplumn"`

	// Cloud Kafka (CKafka).
	Ckafka *uint64 `json:"Ckafka,omitnil,omitempty" name:"Ckafka"`

	// Grocery.
	Grocery *uint64 `json:"Grocery,omitnil,omitempty" name:"Grocery"`

	// Data encryption service.
	HSM *uint64 `json:"HSM,omitnil,omitempty" name:"HSM"`

	// Game storage - Tcaplus.
	Tcaplus *uint64 `json:"Tcaplus,omitnil,omitempty" name:"Tcaplus"`

	// Cnas.
	Cnas *uint64 `json:"Cnas,omitnil,omitempty" name:"Cnas"`

	// HTAP database - TiDB.
	TiDB *uint64 `json:"TiDB,omitnil,omitempty" name:"TiDB"`

	// EMR cluster.
	Emr *uint64 `json:"Emr,omitnil,omitempty" name:"Emr"`

	// SEAL.
	SEAL *uint64 `json:"SEAL,omitnil,omitempty" name:"SEAL"`

	// Cloud file storage - CFS.
	CFS *uint64 `json:"CFS,omitnil,omitempty" name:"CFS"`

	// Oracle.
	Oracle *uint64 `json:"Oracle,omitnil,omitempty" name:"Oracle"`

	// ElasticSearch Service.
	ElasticSearch *uint64 `json:"ElasticSearch,omitnil,omitempty" name:"ElasticSearch"`

	// Blockchain service.
	TBaaS *uint64 `json:"TBaaS,omitnil,omitempty" name:"TBaaS"`

	// Itop.
	Itop *uint64 `json:"Itop,omitnil,omitempty" name:"Itop"`

	// Cloud database audit.
	DBAudit *uint64 `json:"DBAudit,omitnil,omitempty" name:"DBAudit"`

	// Enterprise TencentDB - CynosDB for Postgres.
	CynosDBPostgres *uint64 `json:"CynosDBPostgres,omitnil,omitempty" name:"CynosDBPostgres"`

	// TencentDB for Redis.
	Redis *uint64 `json:"Redis,omitnil,omitempty" name:"Redis"`

	// TencentDB for MongoDB.
	MongoDB *uint64 `json:"MongoDB,omitnil,omitempty" name:"MongoDB"`

	// A distributed cloud database - TencentDB for TDSQL.
	DCDB *uint64 `json:"DCDB,omitnil,omitempty" name:"DCDB"`

	// An enterprise-grade TencentDB - CynosDB for MySQL.
	CynosDBMySQL *uint64 `json:"CynosDBMySQL,omitnil,omitempty" name:"CynosDBMySQL"`

	// Subnets.
	Subnet *uint64 `json:"Subnet,omitnil,omitempty" name:"Subnet"`

	// Route table.
	RouteTable *uint64 `json:"RouteTable,omitnil,omitempty" name:"RouteTable"`
}

type ResourceStatistics struct {
	// VPC instance ID, such as vpc-f1xjkw1b.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Subnet instance ID, such as `subnet-bthucmmy`.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// The total number of used IP addresses.
	Ip *uint64 `json:"Ip,omitnil,omitempty" name:"Ip"`

	// Information of associated resources
	ResourceStatisticsItemSet []*ResourceStatisticsItem `json:"ResourceStatisticsItemSet,omitnil,omitempty" name:"ResourceStatisticsItemSet"`
}

type ResourceStatisticsItem struct {
	// Resource type, such as CVM, ENI
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// Resource name.
	ResourceName *string `json:"ResourceName,omitnil,omitempty" name:"ResourceName"`

	// Number of resources
	ResourceCount *uint64 `json:"ResourceCount,omitnil,omitempty" name:"ResourceCount"`
}

// Predefined struct for user
type ResumeSnapshotInstanceRequestParams struct {
	// Snapshot policy IDs
	SnapshotPolicyId *string `json:"SnapshotPolicyId,omitnil,omitempty" name:"SnapshotPolicyId"`

	// Snapshot file ID
	SnapshotFileId *string `json:"SnapshotFileId,omitnil,omitempty" name:"SnapshotFileId"`

	// ID of the instance.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type ResumeSnapshotInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Snapshot policy IDs
	SnapshotPolicyId *string `json:"SnapshotPolicyId,omitnil,omitempty" name:"SnapshotPolicyId"`

	// Snapshot file ID
	SnapshotFileId *string `json:"SnapshotFileId,omitnil,omitempty" name:"SnapshotFileId"`

	// ID of the instance.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *ResumeSnapshotInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResumeSnapshotInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SnapshotPolicyId")
	delete(f, "SnapshotFileId")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResumeSnapshotInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResumeSnapshotInstanceResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ResumeSnapshotInstanceResponse struct {
	*tchttp.BaseResponse
	Response *ResumeSnapshotInstanceResponseParams `json:"Response"`
}

func (r *ResumeSnapshotInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResumeSnapshotInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReturnNormalAddressesRequestParams struct {
	// EIP addresses. Example: 101.35.139.183
	AddressIps []*string `json:"AddressIps,omitnil,omitempty" name:"AddressIps"`
}

type ReturnNormalAddressesRequest struct {
	*tchttp.BaseRequest
	
	// EIP addresses. Example: 101.35.139.183
	AddressIps []*string `json:"AddressIps,omitnil,omitempty" name:"AddressIps"`
}

func (r *ReturnNormalAddressesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReturnNormalAddressesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AddressIps")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReturnNormalAddressesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReturnNormalAddressesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ReturnNormalAddressesResponse struct {
	*tchttp.BaseResponse
	Response *ReturnNormalAddressesResponseParams `json:"Response"`
}

func (r *ReturnNormalAddressesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReturnNormalAddressesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Route struct {
	// Destination IP range, such as 112.20.51.0/24. Values cannot be in the VPC IP range.
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitnil,omitempty" name:"DestinationCidrBlock"`

	// Type of the next hop. Valid values:
	// `CVM`: public gateway CVM;
	// `VPN`: VPN gateway;
	// `DIRECTCONNECT`: direct connect gateway;
	// `PEERCONNECTION`: peering connection;
	// `HAVIP`: HAVIP;
	// `NAT`: NAT Gateway; 
	// `NORMAL_CVM`: normal CVM;
	// `EIP`: public IP address of the CVM;
	// `LOCAL_GATEWAY`: local gateway.
	GatewayType *string `json:"GatewayType,omitnil,omitempty" name:"GatewayType"`

	// Next hop address. You simply need to specify the gateway ID of a different next hop type, and the system will automatically match the next hop address. 
	// Note: If `GatewayType` is set to `NORMAL_CVM`, `GatewayId` should be the private IP of the instance.
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// Routing policy ID. The IPv4 routing policy will have a meaningful value, while the IPv6 routing policy is always 0. We recommend using the unique ID `RouteItemId` for the routing policy.
	// This field is required when you want to delete a routing policy.
	RouteId *uint64 `json:"RouteId,omitnil,omitempty" name:"RouteId"`

	// The description of the routing policy.
	RouteDescription *string `json:"RouteDescription,omitnil,omitempty" name:"RouteDescription"`

	// Whether it is enabled
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// The route type. Currently, the following types are supported:
	// USER: User route;
	// NETD: Network probe route. When creating a network probe route, the system delivers by default. It cannot be edited or deleted;
	// CCN: CCN route. The system delivers by default. It cannot be edited or deleted.
	// Users can only add and operate USER-type routes.
	RouteType *string `json:"RouteType,omitnil,omitempty" name:"RouteType"`

	// Route table instance ID, such as rtb-azd4dt1c.
	RouteTableId *string `json:"RouteTableId,omitnil,omitempty" name:"RouteTableId"`

	// Destination IPv6 IP range, which cannot be included in VPC IP range, such as 2402:4e00:1000:810b::/64.
	DestinationIpv6CidrBlock *string `json:"DestinationIpv6CidrBlock,omitnil,omitempty" name:"DestinationIpv6CidrBlock"`

	// Unique routing policy ID.
	RouteItemId *string `json:"RouteItemId,omitnil,omitempty" name:"RouteItemId"`

	// Whether the routing policy is published to CCN.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	PublishedToVbc *bool `json:"PublishedToVbc,omitnil,omitempty" name:"PublishedToVbc"`

	// Creation time of the routing policy
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`
}

type RouteTable struct {
	// VPC instance ID.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// The route table instance ID, such as `rtb-azd4dt1c`.
	RouteTableId *string `json:"RouteTableId,omitnil,omitempty" name:"RouteTableId"`

	// Route table name.
	RouteTableName *string `json:"RouteTableName,omitnil,omitempty" name:"RouteTableName"`

	// The association relationships of the route table.
	AssociationSet []*RouteTableAssociation `json:"AssociationSet,omitnil,omitempty" name:"AssociationSet"`

	// IPv4 routing policy set.
	RouteSet []*Route `json:"RouteSet,omitnil,omitempty" name:"RouteSet"`

	// Whether it is the default route table.
	Main *bool `json:"Main,omitnil,omitempty" name:"Main"`

	// Creation Time.
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// Tag key-value pairs.
	TagSet []*Tag `json:"TagSet,omitnil,omitempty" name:"TagSet"`

	// Whether the local route is published to CCN.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	LocalCidrForCcn []*CidrForCcn `json:"LocalCidrForCcn,omitnil,omitempty" name:"LocalCidrForCcn"`
}

type RouteTableAssociation struct {
	// Subnet instance ID.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Route table instance ID.
	RouteTableId *string `json:"RouteTableId,omitnil,omitempty" name:"RouteTableId"`
}

type SecurityGroup struct {
	// The security group instance ID, such as `sg-ohuuioma`.
	SecurityGroupId *string `json:"SecurityGroupId,omitnil,omitempty" name:"SecurityGroupId"`

	// Security group can be named freely, but cannot exceed 60 characters.
	SecurityGroupName *string `json:"SecurityGroupName,omitnil,omitempty" name:"SecurityGroupName"`

	// The remarks for the security group. The maximum length is 100 characters.
	SecurityGroupDesc *string `json:"SecurityGroupDesc,omitnil,omitempty" name:"SecurityGroupDesc"`

	// The project id is 0 by default. You can query this in the project management page of the Qcloud console.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Whether it is the default security group (which cannot be deleted).
	IsDefault *bool `json:"IsDefault,omitnil,omitempty" name:"IsDefault"`

	// Security group creation time.
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// Tag key-value pairs.
	TagSet []*Tag `json:"TagSet,omitnil,omitempty" name:"TagSet"`

	// Security group update time.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type SecurityGroupAssociationStatistics struct {
	// Security group instance ID.
	SecurityGroupId *string `json:"SecurityGroupId,omitnil,omitempty" name:"SecurityGroupId"`

	// Number of CVM instances.
	CVM *uint64 `json:"CVM,omitnil,omitempty" name:"CVM"`

	// Number of TencentDB for MySQL instances
	CDB *uint64 `json:"CDB,omitnil,omitempty" name:"CDB"`

	// Number of ENI instances.
	ENI *uint64 `json:"ENI,omitnil,omitempty" name:"ENI"`

	// Number of times a security group is referenced by other security groups
	SG *uint64 `json:"SG,omitnil,omitempty" name:"SG"`

	// Number of load balancer instances.
	CLB *uint64 `json:"CLB,omitnil,omitempty" name:"CLB"`

	// The binding statistics for all instances.
	InstanceStatistics []*InstanceStatistic `json:"InstanceStatistics,omitnil,omitempty" name:"InstanceStatistics"`

	// Total count of all resources (excluding resources referenced by security groups).
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`
}

type SecurityGroupPolicy struct {
	// The index number of security group rules, which dynamically changes with the rules. This parameter can be obtained via the `DescribeSecurityGroupPolicies` API and used with the `Version` field in the returned value of the API.
	PolicyIndex *int64 `json:"PolicyIndex,omitnil,omitempty" name:"PolicyIndex"`

	// Protocol. Valid values: TCP, UDP, ICMP, ICMPv6, ALL.
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// Port (`all`, a single port, or a port range).
	// Note: If the `Protocol` value is set to `ALL`, the `Port` value also needs to be set to `all`.
	Port *string `json:"Port,omitnil,omitempty" name:"Port"`

	// Protocol port ID or protocol port group ID. ServiceTemplate and Protocol+Port are mutually exclusive.
	ServiceTemplate *ServiceTemplateSpecification `json:"ServiceTemplate,omitnil,omitempty" name:"ServiceTemplate"`

	// Either `CidrBlock` or `Ipv6CidrBlock can be specified. Note that if `0.0.0.0/n` is entered, it is mapped to 0.0.0.0/0.
	CidrBlock *string `json:"CidrBlock,omitnil,omitempty" name:"CidrBlock"`

	// The CIDR block or IPv6 (mutually exclusive).
	Ipv6CidrBlock *string `json:"Ipv6CidrBlock,omitnil,omitempty" name:"Ipv6CidrBlock"`

	// The security group instance ID, such as `sg-ohuuioma`.
	SecurityGroupId *string `json:"SecurityGroupId,omitnil,omitempty" name:"SecurityGroupId"`

	// IP address ID or IP address group ID.
	AddressTemplate *AddressTemplateSpecification `json:"AddressTemplate,omitnil,omitempty" name:"AddressTemplate"`

	// ACCEPT or DROP.
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`

	// Security group policy description.
	PolicyDescription *string `json:"PolicyDescription,omitnil,omitempty" name:"PolicyDescription"`

	// The last modification time of the security group.
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`
}

type SecurityGroupPolicySet struct {
	// The version number of the security group policy, which will automatically increase by one each time you update the security group policy, so as to prevent expiration of the updated routing policies. If it is left empty, any conflicts will be ignored. 
	// Note: This parameter may return null, indicating that no valid values can be obtained.
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// Outbound rule. 
	// Note: This parameter may return null, indicating that no valid values can be obtained.
	Egress []*SecurityGroupPolicy `json:"Egress,omitnil,omitempty" name:"Egress"`

	// Inbound rule. 
	// Note: This parameter may return null, indicating that no valid values can be obtained.
	Ingress []*SecurityGroupPolicy `json:"Ingress,omitnil,omitempty" name:"Ingress"`
}

type SecurityPolicyDatabase struct {
	// Local IP range
	LocalCidrBlock *string `json:"LocalCidrBlock,omitnil,omitempty" name:"LocalCidrBlock"`

	// Opposite IP range
	RemoteCidrBlock []*string `json:"RemoteCidrBlock,omitnil,omitempty" name:"RemoteCidrBlock"`
}

type ServiceTemplate struct {
	// Protocol port instance ID, such as `ppm-f5n1f8da`.
	ServiceTemplateId *string `json:"ServiceTemplateId,omitnil,omitempty" name:"ServiceTemplateId"`

	// Template name.
	ServiceTemplateName *string `json:"ServiceTemplateName,omitnil,omitempty" name:"ServiceTemplateName"`

	// Protocol port information.
	ServiceSet []*string `json:"ServiceSet,omitnil,omitempty" name:"ServiceSet"`

	// Creation Time.
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// Protocol port template information with remarks
	ServiceExtraSet []*ServicesInfo `json:"ServiceExtraSet,omitnil,omitempty" name:"ServiceExtraSet"`
}

type ServiceTemplateGroup struct {
	// Protocol port template group instance ID, such as `ppmg-2klmrefu`.
	ServiceTemplateGroupId *string `json:"ServiceTemplateGroupId,omitnil,omitempty" name:"ServiceTemplateGroupId"`

	// Protocol port template group name.
	ServiceTemplateGroupName *string `json:"ServiceTemplateGroupName,omitnil,omitempty" name:"ServiceTemplateGroupName"`

	// Protocol port template instance ID.
	ServiceTemplateIdSet []*string `json:"ServiceTemplateIdSet,omitnil,omitempty" name:"ServiceTemplateIdSet"`

	// Creation Time.
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// Protocol port template instance information.
	ServiceTemplateSet []*ServiceTemplate `json:"ServiceTemplateSet,omitnil,omitempty" name:"ServiceTemplateSet"`
}

type ServiceTemplateSpecification struct {
	// Protocol port ID, such as `ppm-f5n1f8da`.
	ServiceId *string `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`

	// Protocol port group ID, such as `ppmg-f5n1f8da`.
	ServiceGroupId *string `json:"ServiceGroupId,omitnil,omitempty" name:"ServiceGroupId"`
}

type ServicesInfo struct {
	// Protocol port
	Service *string `json:"Service,omitnil,omitempty" name:"Service"`

	// Remarks
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

// Predefined struct for user
type SetCcnRegionBandwidthLimitsRequestParams struct {
	// The CCN instance ID, such as `ccn-f49l6u0z`.
	CcnId *string `json:"CcnId,omitnil,omitempty" name:"CcnId"`

	// The outbound bandwidth cap of each CCN region.
	CcnRegionBandwidthLimits []*CcnRegionBandwidthLimit `json:"CcnRegionBandwidthLimits,omitnil,omitempty" name:"CcnRegionBandwidthLimits"`

	// Whether to restore the region outbound bandwidth limit or inter-region bandwidth limit to default 1 Gbps. Valid values: `false` (no); `true` (yes). Default value: `false`. When the parameter is set to `true`, the CCN instance created will not be displayed in the console.
	SetDefaultLimitFlag *bool `json:"SetDefaultLimitFlag,omitnil,omitempty" name:"SetDefaultLimitFlag"`
}

type SetCcnRegionBandwidthLimitsRequest struct {
	*tchttp.BaseRequest
	
	// The CCN instance ID, such as `ccn-f49l6u0z`.
	CcnId *string `json:"CcnId,omitnil,omitempty" name:"CcnId"`

	// The outbound bandwidth cap of each CCN region.
	CcnRegionBandwidthLimits []*CcnRegionBandwidthLimit `json:"CcnRegionBandwidthLimits,omitnil,omitempty" name:"CcnRegionBandwidthLimits"`

	// Whether to restore the region outbound bandwidth limit or inter-region bandwidth limit to default 1 Gbps. Valid values: `false` (no); `true` (yes). Default value: `false`. When the parameter is set to `true`, the CCN instance created will not be displayed in the console.
	SetDefaultLimitFlag *bool `json:"SetDefaultLimitFlag,omitnil,omitempty" name:"SetDefaultLimitFlag"`
}

func (r *SetCcnRegionBandwidthLimitsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetCcnRegionBandwidthLimitsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CcnId")
	delete(f, "CcnRegionBandwidthLimits")
	delete(f, "SetDefaultLimitFlag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetCcnRegionBandwidthLimitsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetCcnRegionBandwidthLimitsResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SetCcnRegionBandwidthLimitsResponse struct {
	*tchttp.BaseResponse
	Response *SetCcnRegionBandwidthLimitsResponseParams `json:"Response"`
}

func (r *SetCcnRegionBandwidthLimitsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetCcnRegionBandwidthLimitsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetVpnGatewaysRenewFlagRequestParams struct {
	// VPN gateway IDs
	VpnGatewayIds []*string `json:"VpnGatewayIds,omitnil,omitempty" name:"VpnGatewayIds"`

	// Status of auto-renewal
	// Values: `0` (Follow original), `1` (Enable auto-renewal), `2` (Disable auto-renewal) 
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// VPNGW type: `IPSEC`, `SSL`
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type SetVpnGatewaysRenewFlagRequest struct {
	*tchttp.BaseRequest
	
	// VPN gateway IDs
	VpnGatewayIds []*string `json:"VpnGatewayIds,omitnil,omitempty" name:"VpnGatewayIds"`

	// Status of auto-renewal
	// Values: `0` (Follow original), `1` (Enable auto-renewal), `2` (Disable auto-renewal) 
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// VPNGW type: `IPSEC`, `SSL`
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

func (r *SetVpnGatewaysRenewFlagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetVpnGatewaysRenewFlagRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpnGatewayIds")
	delete(f, "AutoRenewFlag")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetVpnGatewaysRenewFlagRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetVpnGatewaysRenewFlagResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SetVpnGatewaysRenewFlagResponse struct {
	*tchttp.BaseResponse
	Response *SetVpnGatewaysRenewFlagResponseParams `json:"Response"`
}

func (r *SetVpnGatewaysRenewFlagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetVpnGatewaysRenewFlagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SnapshotFileInfo struct {
	// Snapshot policy ID
	SnapshotPolicyId *string `json:"SnapshotPolicyId,omitnil,omitempty" name:"SnapshotPolicyId"`

	// ID of the instance.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Snapshot file ID
	SnapshotFileId *string `json:"SnapshotFileId,omitnil,omitempty" name:"SnapshotFileId"`

	// Backup time
	BackupTime *string `json:"BackupTime,omitnil,omitempty" name:"BackupTime"`

	// Operator UIN
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`
}

type SnapshotInstance struct {
	// ID of the instance.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Type of associated resource. Values: `securitygroup`
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// Instance region
	InstanceRegion *string `json:"InstanceRegion,omitnil,omitempty" name:"InstanceRegion"`

	// Snapshot policy IDs
	SnapshotPolicyId *string `json:"SnapshotPolicyId,omitnil,omitempty" name:"SnapshotPolicyId"`

	// The instance name.
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`
}

type SnapshotPolicy struct {
	// Snapshot policy name
	SnapshotPolicyName *string `json:"SnapshotPolicyName,omitnil,omitempty" name:"SnapshotPolicyName"`

	// Backup policy type. Values: `operate` (Manual backup); `time` (Scheduled backup)
	BackupType *string `json:"BackupType,omitnil,omitempty" name:"BackupType"`

	// Snapshot retention period in days. Range: 1 to 365.
	KeepTime *uint64 `json:"KeepTime,omitnil,omitempty" name:"KeepTime"`

	// Whether to create a new COS bucket. It defaults to `False`.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	CreateNewCos *bool `json:"CreateNewCos,omitnil,omitempty" name:"CreateNewCos"`

	// Region of the COS bucket
	CosRegion *string `json:"CosRegion,omitnil,omitempty" name:"CosRegion"`

	// COS bucket
	CosBucket *string `json:"CosBucket,omitnil,omitempty" name:"CosBucket"`

	// Snapshot policy ID
	SnapshotPolicyId *string `json:"SnapshotPolicyId,omitnil,omitempty" name:"SnapshotPolicyId"`

	// Scheduled backup policies
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	BackupPolicies []*BackupPolicy `json:"BackupPolicies,omitnil,omitempty" name:"BackupPolicies"`

	// Whether to enable the policy. Values: `True` (default), `False`
	Enable *bool `json:"Enable,omitnil,omitempty" name:"Enable"`

	// Creation time
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`
}

type SourceIpTranslationNatRule struct {
	// Resource ID. It can be left empty if `ResourceType` is `USERDEFINED`.
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// Resource type. Valid values: `SUBNET`, `NETWORKINTERFACE`, `USERDEFINED`
	// Note: This field may return null, indicating that no valid values can be obtained.
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// Source IP/IP range
	PrivateIpAddress *string `json:"PrivateIpAddress,omitnil,omitempty" name:"PrivateIpAddress"`

	// Elastic IP address pool
	PublicIpAddresses []*string `json:"PublicIpAddresses,omitnil,omitempty" name:"PublicIpAddresses"`

	// Description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// SNAT rule ID
	NatGatewaySnatId *string `json:"NatGatewaySnatId,omitnil,omitempty" name:"NatGatewaySnatId"`

	// NAT gateway ID.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	NatGatewayId *string `json:"NatGatewayId,omitnil,omitempty" name:"NatGatewayId"`

	// VPC ID.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// The creation time of a NAT gateway's SNAT rule.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`
}

type Subnet struct {
	// The `ID` of the `VPC` instance.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Subnet instance `ID`, such as `subnet-bthucmmy`.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Subnet name.
	SubnetName *string `json:"SubnetName,omitnil,omitempty" name:"SubnetName"`

	// The `IPv4` `CIDR` of the subnet.
	CidrBlock *string `json:"CidrBlock,omitnil,omitempty" name:"CidrBlock"`

	// Whether it is the default subnet.
	IsDefault *bool `json:"IsDefault,omitnil,omitempty" name:"IsDefault"`

	// Whether to enable broadcast.
	EnableBroadcast *bool `json:"EnableBroadcast,omitnil,omitempty" name:"EnableBroadcast"`

	// Availability Zone.
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// The route table instance ID, such as `rtb-l2h8d7c2`.
	RouteTableId *string `json:"RouteTableId,omitnil,omitempty" name:"RouteTableId"`

	// Creation Time.
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// The number of available IPv4 addresses
	AvailableIpAddressCount *uint64 `json:"AvailableIpAddressCount,omitnil,omitempty" name:"AvailableIpAddressCount"`

	// The `IPv6` `CIDR` of the subnet.
	Ipv6CidrBlock *string `json:"Ipv6CidrBlock,omitnil,omitempty" name:"Ipv6CidrBlock"`

	// The associated `ACL`ID
	NetworkAclId *string `json:"NetworkAclId,omitnil,omitempty" name:"NetworkAclId"`

	// Whether it is a `SNAT` address pool subnet.
	IsRemoteVpcSnat *bool `json:"IsRemoteVpcSnat,omitnil,omitempty" name:"IsRemoteVpcSnat"`

	// The total number of IPv4 addresses in the subnet.
	TotalIpAddressCount *uint64 `json:"TotalIpAddressCount,omitnil,omitempty" name:"TotalIpAddressCount"`

	// Tag key-value pairs
	TagSet []*Tag `json:"TagSet,omitnil,omitempty" name:"TagSet"`

	// CDC instance ID
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	CdcId *string `json:"CdcId,omitnil,omitempty" name:"CdcId"`

	// Whether it is a CDC subnet. Valid values: 0: no; 1: yes
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	IsCdcSubnet *int64 `json:"IsCdcSubnet,omitnil,omitempty" name:"IsCdcSubnet"`
}

type SubnetInput struct {
	// The `CIDR` of the subnet.
	CidrBlock *string `json:"CidrBlock,omitnil,omitempty" name:"CidrBlock"`

	// Subnet name.
	SubnetName *string `json:"SubnetName,omitnil,omitempty" name:"SubnetName"`

	// The availability zone, such as `ap-guangzhou-2`.
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// The specified associated route table, such as `rtb-3ryrwzuu`.
	RouteTableId *string `json:"RouteTableId,omitnil,omitempty" name:"RouteTableId"`
}

type Tag struct {
	// Tag key
	// Note: This field may return null, indicating no valid value.
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// Tag value
	// Note: This field may return null, indicating no valid value.
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type TrafficPackage struct {
	// Unique traffic package ID
	TrafficPackageId *string `json:"TrafficPackageId,omitnil,omitempty" name:"TrafficPackageId"`

	// Traffic package name
	// Note: This field may return null, indicating that no valid values can be obtained.
	TrafficPackageName *string `json:"TrafficPackageName,omitnil,omitempty" name:"TrafficPackageName"`

	// Traffic package size in GB
	TotalAmount *float64 `json:"TotalAmount,omitnil,omitempty" name:"TotalAmount"`

	// Traffic package balance in GB
	RemainingAmount *float64 `json:"RemainingAmount,omitnil,omitempty" name:"RemainingAmount"`

	// Traffic package status. Valid values: `AVAILABLE`, `EXPIRED`, `EXHAUSTED`, `REFUNDED`, `DELETED`
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Traffic package creation time
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// Traffic package expiration time
	Deadline *string `json:"Deadline,omitnil,omitempty" name:"Deadline"`

	// Used traffic in GB
	UsedAmount *float64 `json:"UsedAmount,omitnil,omitempty" name:"UsedAmount"`

	// Traffic package tag
	// Note: This field may return null, indicating that no valid values can be obtained.
	TagSet []*Tag `json:"TagSet,omitnil,omitempty" name:"TagSet"`

	// Traffic package type (idle-time or full-time)
	DeductType *string `json:"DeductType,omitnil,omitempty" name:"DeductType"`
}

// Predefined struct for user
type TransformAddressRequestParams struct {
	// The ID of the instance with a common public IP to be operated on, such as `ins-11112222`. You can query the instance ID by logging into the [CVM console](https://console.cloud.tencent.com/cvm). You can also obtain the parameter value from the `InstanceId` field in the returned result of the API [DescribeInstances](https://intl.cloud.tencent.com/document/product/213/33256).
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type TransformAddressRequest struct {
	*tchttp.BaseRequest
	
	// The ID of the instance with a common public IP to be operated on, such as `ins-11112222`. You can query the instance ID by logging into the [CVM console](https://console.cloud.tencent.com/cvm). You can also obtain the parameter value from the `InstanceId` field in the returned result of the API [DescribeInstances](https://intl.cloud.tencent.com/document/product/213/33256).
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *TransformAddressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TransformAddressRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TransformAddressRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TransformAddressResponseParams struct {
	// The async task ID. You can use the [DescribeTaskResult](https://intl.cloud.tencent.com/document/api/215/36271?from_cn_redirect=1) API to query the task status.
	TaskId *uint64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique ID after converting to EIP
	AddressId *string `json:"AddressId,omitnil,omitempty" name:"AddressId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type TransformAddressResponse struct {
	*tchttp.BaseResponse
	Response *TransformAddressResponseParams `json:"Response"`
}

func (r *TransformAddressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TransformAddressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnassignIpv6AddressesRequestParams struct {
	// The `ID` of the ENI instance, such as `eni-m6dyj72l`.
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitnil,omitempty" name:"NetworkInterfaceId"`

	// The list of specified `IPv6` addresses. A maximum of 10 can be specified each time.
	Ipv6Addresses []*Ipv6Address `json:"Ipv6Addresses,omitnil,omitempty" name:"Ipv6Addresses"`
}

type UnassignIpv6AddressesRequest struct {
	*tchttp.BaseRequest
	
	// The `ID` of the ENI instance, such as `eni-m6dyj72l`.
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitnil,omitempty" name:"NetworkInterfaceId"`

	// The list of specified `IPv6` addresses. A maximum of 10 can be specified each time.
	Ipv6Addresses []*Ipv6Address `json:"Ipv6Addresses,omitnil,omitempty" name:"Ipv6Addresses"`
}

func (r *UnassignIpv6AddressesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnassignIpv6AddressesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NetworkInterfaceId")
	delete(f, "Ipv6Addresses")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnassignIpv6AddressesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnassignIpv6AddressesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UnassignIpv6AddressesResponse struct {
	*tchttp.BaseResponse
	Response *UnassignIpv6AddressesResponseParams `json:"Response"`
}

func (r *UnassignIpv6AddressesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnassignIpv6AddressesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnassignIpv6CidrBlockRequestParams struct {
	// The `ID` of the `VPC`, such as `vpc-f49l6u0z`.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// The `IPv6` IP range, such as `3402:4e00:20:1000::/56`
	Ipv6CidrBlock *string `json:"Ipv6CidrBlock,omitnil,omitempty" name:"Ipv6CidrBlock"`
}

type UnassignIpv6CidrBlockRequest struct {
	*tchttp.BaseRequest
	
	// The `ID` of the `VPC`, such as `vpc-f49l6u0z`.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// The `IPv6` IP range, such as `3402:4e00:20:1000::/56`
	Ipv6CidrBlock *string `json:"Ipv6CidrBlock,omitnil,omitempty" name:"Ipv6CidrBlock"`
}

func (r *UnassignIpv6CidrBlockRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnassignIpv6CidrBlockRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcId")
	delete(f, "Ipv6CidrBlock")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnassignIpv6CidrBlockRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnassignIpv6CidrBlockResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UnassignIpv6CidrBlockResponse struct {
	*tchttp.BaseResponse
	Response *UnassignIpv6CidrBlockResponseParams `json:"Response"`
}

func (r *UnassignIpv6CidrBlockResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnassignIpv6CidrBlockResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnassignIpv6SubnetCidrBlockRequestParams struct {
	// The `ID` of the VPC where the subnet is located, such as `vpc-f49l6u0z`.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// The `IPv6` subnet IP range list.
	Ipv6SubnetCidrBlocks []*Ipv6SubnetCidrBlock `json:"Ipv6SubnetCidrBlocks,omitnil,omitempty" name:"Ipv6SubnetCidrBlocks"`
}

type UnassignIpv6SubnetCidrBlockRequest struct {
	*tchttp.BaseRequest
	
	// The `ID` of the VPC where the subnet is located, such as `vpc-f49l6u0z`.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// The `IPv6` subnet IP range list.
	Ipv6SubnetCidrBlocks []*Ipv6SubnetCidrBlock `json:"Ipv6SubnetCidrBlocks,omitnil,omitempty" name:"Ipv6SubnetCidrBlocks"`
}

func (r *UnassignIpv6SubnetCidrBlockRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnassignIpv6SubnetCidrBlockRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcId")
	delete(f, "Ipv6SubnetCidrBlocks")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnassignIpv6SubnetCidrBlockRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnassignIpv6SubnetCidrBlockResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UnassignIpv6SubnetCidrBlockResponse struct {
	*tchttp.BaseResponse
	Response *UnassignIpv6SubnetCidrBlockResponseParams `json:"Response"`
}

func (r *UnassignIpv6SubnetCidrBlockResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnassignIpv6SubnetCidrBlockResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnassignPrivateIpAddressesRequestParams struct {
	// The ID of the ENI instance, such as `eni-m6dyj72l`.
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitnil,omitempty" name:"NetworkInterfaceId"`

	// The information of the specified private IPs. You can specify a maximum of 10 each time.
	PrivateIpAddresses []*PrivateIpAddressSpecification `json:"PrivateIpAddresses,omitnil,omitempty" name:"PrivateIpAddresses"`

	// Instance ID of the server bound with this IP. This parameter is only applicable when you need to return an IP and unbind the related servers.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type UnassignPrivateIpAddressesRequest struct {
	*tchttp.BaseRequest
	
	// The ID of the ENI instance, such as `eni-m6dyj72l`.
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitnil,omitempty" name:"NetworkInterfaceId"`

	// The information of the specified private IPs. You can specify a maximum of 10 each time.
	PrivateIpAddresses []*PrivateIpAddressSpecification `json:"PrivateIpAddresses,omitnil,omitempty" name:"PrivateIpAddresses"`

	// Instance ID of the server bound with this IP. This parameter is only applicable when you need to return an IP and unbind the related servers.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *UnassignPrivateIpAddressesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnassignPrivateIpAddressesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NetworkInterfaceId")
	delete(f, "PrivateIpAddresses")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnassignPrivateIpAddressesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnassignPrivateIpAddressesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UnassignPrivateIpAddressesResponse struct {
	*tchttp.BaseResponse
	Response *UnassignPrivateIpAddressesResponseParams `json:"Response"`
}

func (r *UnassignPrivateIpAddressesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnassignPrivateIpAddressesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Vpc struct {
	// `VPC` name.
	VpcName *string `json:"VpcName,omitnil,omitempty" name:"VpcName"`

	// `VPC` instance `ID`, such as `vpc-azd4dt1c`.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// The `IPv4` `CIDR` of the `VPC`.
	CidrBlock *string `json:"CidrBlock,omitnil,omitempty" name:"CidrBlock"`

	// Whether it is the default `VPC`.
	IsDefault *bool `json:"IsDefault,omitnil,omitempty" name:"IsDefault"`

	// Whether multicast is enabled.
	EnableMulticast *bool `json:"EnableMulticast,omitnil,omitempty" name:"EnableMulticast"`

	// Creation Time.
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// `DNS` list.
	DnsServerSet []*string `json:"DnsServerSet,omitnil,omitempty" name:"DnsServerSet"`

	// DHCP domain name option value.
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// `DHCP` option set `ID`.
	DhcpOptionsId *string `json:"DhcpOptionsId,omitnil,omitempty" name:"DhcpOptionsId"`

	// Whether `DHCP` is enabled.
	EnableDhcp *bool `json:"EnableDhcp,omitnil,omitempty" name:"EnableDhcp"`

	// The `IPv6` `CIDR` of the `VPC`.
	Ipv6CidrBlock *string `json:"Ipv6CidrBlock,omitnil,omitempty" name:"Ipv6CidrBlock"`

	// Tag key-value pair
	TagSet []*Tag `json:"TagSet,omitnil,omitempty" name:"TagSet"`

	// The secondary CIDR block.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AssistantCidrSet []*AssistantCidr `json:"AssistantCidrSet,omitnil,omitempty" name:"AssistantCidrSet"`
}

type VpcEndPointServiceUser struct {
	// APP ID
	Owner *uint64 `json:"Owner,omitnil,omitempty" name:"Owner"`

	// User UIN
	UserUin *string `json:"UserUin,omitnil,omitempty" name:"UserUin"`

	// Description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Creation time
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Endpoint service ID
	EndPointServiceId *string `json:"EndPointServiceId,omitnil,omitempty" name:"EndPointServiceId"`
}

type VpcIpv6Address struct {
	// `VPC` private `IPv6` address
	Ipv6Address *string `json:"Ipv6Address,omitnil,omitempty" name:"Ipv6Address"`

	// The `IPv6` `CIDR` belonging to the subnet.
	CidrBlock *string `json:"CidrBlock,omitnil,omitempty" name:"CidrBlock"`

	// `IPv6` type.
	Ipv6AddressType *string `json:"Ipv6AddressType,omitnil,omitempty" name:"Ipv6AddressType"`

	// `IPv6` application time.
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`
}

type VpcPrivateIpAddress struct {
	// `VPC` private `IP`.
	PrivateIpAddress *string `json:"PrivateIpAddress,omitnil,omitempty" name:"PrivateIpAddress"`

	// The `CIDR` belonging to the subnet.
	CidrBlock *string `json:"CidrBlock,omitnil,omitempty" name:"CidrBlock"`

	// Private `IP` type.
	PrivateIpAddressType *string `json:"PrivateIpAddressType,omitnil,omitempty" name:"PrivateIpAddressType"`

	// `IP` application time.
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`
}

type VpcTaskResultDetailInfo struct {
	// Resource ID 
	// Note: this field may return `null`, indicating that no valid value can be obtained.
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// Status
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

type VpnConnection struct {
	// Tunnel instance ID.
	VpnConnectionId *string `json:"VpnConnectionId,omitnil,omitempty" name:"VpnConnectionId"`

	// Tunnel name.
	VpnConnectionName *string `json:"VpnConnectionName,omitnil,omitempty" name:"VpnConnectionName"`

	// VPC instance ID.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// The ID of the VPN gateway instance.
	VpnGatewayId *string `json:"VpnGatewayId,omitnil,omitempty" name:"VpnGatewayId"`

	// Customer gateway instance ID.
	CustomerGatewayId *string `json:"CustomerGatewayId,omitnil,omitempty" name:"CustomerGatewayId"`

	// The pre-shared key.
	PreShareKey *string `json:"PreShareKey,omitnil,omitempty" name:"PreShareKey"`

	// Tunnel transmission protocol.
	VpnProto *string `json:"VpnProto,omitnil,omitempty" name:"VpnProto"`

	// Tunnel encryption protocol.
	EncryptProto *string `json:"EncryptProto,omitnil,omitempty" name:"EncryptProto"`

	// Route Type.
	RouteType *string `json:"RouteType,omitnil,omitempty" name:"RouteType"`

	// Creation Time.
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// Production status of the tunnel. PENDING: Creating; AVAILABLE: Running; DELETING: Deleting.
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// Connection status of the tunnel. AVAILABLE: Connected.
	NetStatus *string `json:"NetStatus,omitnil,omitempty" name:"NetStatus"`

	// SPD.
	SecurityPolicyDatabaseSet []*SecurityPolicyDatabase `json:"SecurityPolicyDatabaseSet,omitnil,omitempty" name:"SecurityPolicyDatabaseSet"`

	// IKE options.
	IKEOptionsSpecification *IKEOptionsSpecification `json:"IKEOptionsSpecification,omitnil,omitempty" name:"IKEOptionsSpecification"`

	// IPSEC options.
	IPSECOptionsSpecification *IPSECOptionsSpecification `json:"IPSECOptionsSpecification,omitnil,omitempty" name:"IPSECOptionsSpecification"`

	// Whether the health check is supported.
	EnableHealthCheck *bool `json:"EnableHealthCheck,omitnil,omitempty" name:"EnableHealthCheck"`

	// Local IP address for the health check
	HealthCheckLocalIp *string `json:"HealthCheckLocalIp,omitnil,omitempty" name:"HealthCheckLocalIp"`

	// Peer IP address for the health check
	HealthCheckRemoteIp *string `json:"HealthCheckRemoteIp,omitnil,omitempty" name:"HealthCheckRemoteIp"`

	// Tunnel health check status. Valid values: AVAILABLE: healthy; UNAVAILABLE: unhealthy. This parameter will be returned only after health check is enabled.
	HealthCheckStatus *string `json:"HealthCheckStatus,omitnil,omitempty" name:"HealthCheckStatus"`

	// Whether to enable DPD. Values: `0` (Disable) and `1` (Enable)
	// Note: This field may return null, indicating that no valid values can be obtained.
	DpdEnable *int64 `json:"DpdEnable,omitnil,omitempty" name:"DpdEnable"`

	// DPD timeout period. 
	// Note: This field may return null, indicating that no valid values can be obtained.
	DpdTimeout *string `json:"DpdTimeout,omitnil,omitempty" name:"DpdTimeout"`

	// The action to take in case of DPD timeout. Values: `clear` (Disconnect) and `restart` (retry). This parameter only takes effect when `DpdEnable` is set to `1`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DpdAction *string `json:"DpdAction,omitnil,omitempty" name:"DpdAction"`

	// Array of tag key-value pairs
	TagSet []*Tag `json:"TagSet,omitnil,omitempty" name:"TagSet"`

	// Negotiation type
	// Note: This field may return null, indicating that no valid values can be obtained.
	NegotiationType *string `json:"NegotiationType,omitnil,omitempty" name:"NegotiationType"`
}

type VpnGateway struct {
	// Gateway instance ID.
	VpnGatewayId *string `json:"VpnGatewayId,omitnil,omitempty" name:"VpnGatewayId"`

	// VPC instance ID.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Gateway instance name.
	VpnGatewayName *string `json:"VpnGatewayName,omitnil,omitempty" name:"VpnGatewayName"`

	// Gateway instance type. Valid values: 'IPSEC', 'SSL', and 'CCN'.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Gateway instance status. 'PENDING': Creating; 'DELETING': Deleting; 'AVAILABLE': Running.
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// Gateway public IP.
	PublicIpAddress *string `json:"PublicIpAddress,omitnil,omitempty" name:"PublicIpAddress"`

	// Gateway renewal type: 'NOTIFY_AND_MANUAL_RENEW': Manual renewal. 'NOTIFY_AND_AUTO_RENEW': Automatic renewal. 'NOT_NOTIFY_AND_NOT_RENEW': No renewal after expiration.
	RenewFlag *string `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// Gateway billing type: POSTPAID_BY_HOUR: Postpaid by hour; PREPAID: Prepaid.
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// Outbound bandwidth of gateway.
	InternetMaxBandwidthOut *uint64 `json:"InternetMaxBandwidthOut,omitnil,omitempty" name:"InternetMaxBandwidthOut"`

	// Creation Time.
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// Expiration time of the prepaid gateway.
	ExpiredTime *string `json:"ExpiredTime,omitnil,omitempty" name:"ExpiredTime"`

	// Whether the public IP is blocked.
	IsAddressBlocked *bool `json:"IsAddressBlocked,omitnil,omitempty" name:"IsAddressBlocked"`

	// Change of billing method. PREPAID_TO_POSTPAID: Monthly subscription prepaid to postpaid by hour.
	NewPurchasePlan *string `json:"NewPurchasePlan,omitnil,omitempty" name:"NewPurchasePlan"`

	// Gateway billing status. PROTECTIVELY_ISOLATED: Instance is isolated; NORMAL: Normal.
	RestrictState *string `json:"RestrictState,omitnil,omitempty" name:"RestrictState"`

	// The availability zone, such as `ap-guangzhou-2`
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Gateway bandwidth quota information.
	VpnGatewayQuotaSet []*VpnGatewayQuota `json:"VpnGatewayQuotaSet,omitnil,omitempty" name:"VpnGatewayQuotaSet"`

	// Gateway instance version.
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// CCN instance ID when the value of Type is CCN.
	NetworkInstanceId *string `json:"NetworkInstanceId,omitnil,omitempty" name:"NetworkInstanceId"`

	// CDC instance ID
	CdcId *string `json:"CdcId,omitnil,omitempty" name:"CdcId"`

	// Maximum number of connected clients allowed for the SSL VPN gateway.
	MaxConnection *uint64 `json:"MaxConnection,omitnil,omitempty" name:"MaxConnection"`
}

type VpnGatewayQuota struct {
	// The bandwidth quota.
	Bandwidth *uint64 `json:"Bandwidth,omitnil,omitempty" name:"Bandwidth"`

	// The bandwidth quota name in Chinese.
	Cname *string `json:"Cname,omitnil,omitempty" name:"Cname"`

	// The bandwidth quota name in English.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type VpnGatewayRoute struct {
	// Destination IDC IP range
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitnil,omitempty" name:"DestinationCidrBlock"`

	// Next hop type (type of the associated instance). Valid values: `VPNCONN` (VPN tunnel) and `CCN` (CCN instance)
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// Instance ID of the next hop
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Priority. Valid values: `0` and `100`
	Priority *int64 `json:"Priority,omitnil,omitempty" name:"Priority"`

	// Status. Valid values: `ENABLE` and `DISABLE`
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Route ID
	RouteId *string `json:"RouteId,omitnil,omitempty" name:"RouteId"`

	// Route type. Valid values: `VPC`, `CCN` (CCN-propagated route), `Static`, and `BGP`.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// The creation time.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// The update time.
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type VpnGatewayRouteModify struct {
	// Route ID of the VPN gateway
	RouteId *string `json:"RouteId,omitnil,omitempty" name:"RouteId"`

	// Route status of the VPN gateway. Valid values: `ENABLE`, and `DISABLE`.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

type VpngwCcnRoutes struct {
	// Route ID
	RouteId *string `json:"RouteId,omitnil,omitempty" name:"RouteId"`

	// Enable the route or not
	// ENABLE: enable the route
	// DISABLE: do not enable the route
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Route CIDR block
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitnil,omitempty" name:"DestinationCidrBlock"`
}

// Predefined struct for user
type WithdrawNotifyRoutesRequestParams struct {
	// The unique ID of the route table
	RouteTableId *string `json:"RouteTableId,omitnil,omitempty" name:"RouteTableId"`

	// The unique ID of the routing policy
	RouteItemIds []*string `json:"RouteItemIds,omitnil,omitempty" name:"RouteItemIds"`
}

type WithdrawNotifyRoutesRequest struct {
	*tchttp.BaseRequest
	
	// The unique ID of the route table
	RouteTableId *string `json:"RouteTableId,omitnil,omitempty" name:"RouteTableId"`

	// The unique ID of the routing policy
	RouteItemIds []*string `json:"RouteItemIds,omitnil,omitempty" name:"RouteItemIds"`
}

func (r *WithdrawNotifyRoutesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *WithdrawNotifyRoutesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RouteTableId")
	delete(f, "RouteItemIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "WithdrawNotifyRoutesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type WithdrawNotifyRoutesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type WithdrawNotifyRoutesResponse struct {
	*tchttp.BaseResponse
	Response *WithdrawNotifyRoutesResponseParams `json:"Response"`
}

func (r *WithdrawNotifyRoutesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *WithdrawNotifyRoutesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}