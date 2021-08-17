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

package v20200309

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

type AssociateDDoSEipAddressRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS instance ID (only Anti-DDoS Advanced). For example, `bgpip-0000011x`.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// EIP of the Anti-DDoS instance ID
	Eip *string `json:"Eip,omitempty" name:"Eip"`

	// Instance ID to bind. For example, `ins-11112222`. It can be queried in the console or obtained from `InstanceId` returned by `DescribeInstances`.
	CvmInstanceID *string `json:"CvmInstanceID,omitempty" name:"CvmInstanceID"`

	// Region of the CVM instance. For example, `ap-hongkong`.
	CvmRegion *string `json:"CvmRegion,omitempty" name:"CvmRegion"`
}

func (r *AssociateDDoSEipAddressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateDDoSEipAddressRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Eip")
	delete(f, "CvmInstanceID")
	delete(f, "CvmRegion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AssociateDDoSEipAddressRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type AssociateDDoSEipAddressResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AssociateDDoSEipAddressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateDDoSEipAddressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BGPIPInstance struct {

	// Anti-DDoS instance details
	InstanceDetail *InstanceRelation `json:"InstanceDetail,omitempty" name:"InstanceDetail"`

	// Anti-DDoS instance specifications
	SpecificationLimit *BGPIPInstanceSpecification `json:"SpecificationLimit,omitempty" name:"SpecificationLimit"`

	// Anti-DDoS instance usage statistics
	Usage *BGPIPInstanceUsages `json:"Usage,omitempty" name:"Usage"`

	// Region of the Anti-DDoS instance
	Region *RegionInfo `json:"Region,omitempty" name:"Region"`

	// Status of the Anti-DDoS instance. Valid values:
	// `idle`: running
	// `attacking`: under attacks
	// `blocking`: blocked
	// `creating`: creating
	// `deblocking`: unblocking
	// `isolate`: reprocessed and isolated
	Status *string `json:"Status,omitempty" name:"Status"`

	// Purchase time
	ExpiredTime *string `json:"ExpiredTime,omitempty" name:"ExpiredTime"`

	// Expired At
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// Name of the Anti-DDoS instance
	Name *string `json:"Name,omitempty" name:"Name"`

	// Package details of the Anti-DDoS instance.
	// Note: This field is `null` for an Anti-DDoS instance without using a package.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	PackInfo *PackInfo `json:"PackInfo,omitempty" name:"PackInfo"`

	// Non-BGP package details of the Anti-DDoS instance.
	// Note: This field is `null` for an Anti-DDoS instance without using a non-BGP package.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	StaticPackRelation *StaticPackRelation `json:"StaticPackRelation,omitempty" name:"StaticPackRelation"`

	// Used to differentiate Anti-DDoS Advanced lines outside the Chinese mainland
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// Used to differentiate clusters
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Tgw *uint64 `json:"Tgw,omitempty" name:"Tgw"`

	// EIP states: `CREATING`, `BINDING`, `BIND`, `UNBINDING`, `UNBIND`, `OFFLINING`, and `BIND_ENI`. The EIP must be bound to an Anti-DDoS Advanced instance.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	EipAddressStatus *string `json:"EipAddressStatus,omitempty" name:"EipAddressStatus"`

	// Whether it is an Anti-DDoS EIP instance. `1`: Yes; `0`: No.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	EipFlag *int64 `json:"EipFlag,omitempty" name:"EipFlag"`

	// EIP package details of the Anti-DDoS Advanced instance.
	// Note: This field is `null` for an Anti-DDoS Advanced instance without using an EIP package.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	EipAddressPackRelation *EipAddressPackRelation `json:"EipAddressPackRelation,omitempty" name:"EipAddressPackRelation"`

	// Details of the Anti-DDoS Advanced instance bound to the EIP.
	// Note: This field is `null` if the EIP is not bound to an Anti-DDoS Advanced instance.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	EipAddressInfo *EipAddressRelation `json:"EipAddressInfo,omitempty" name:"EipAddressInfo"`
}

type BGPIPInstanceSpecification struct {

	// Base protection bandwidth (in Gbps)
	ProtectBandwidth *uint64 `json:"ProtectBandwidth,omitempty" name:"ProtectBandwidth"`

	// CC protection bandwidth (in QPS)
	ProtectCCQPS *uint64 `json:"ProtectCCQPS,omitempty" name:"ProtectCCQPS"`

	// Normal application bandwidth (in Mbps)
	NormalBandwidth *uint64 `json:"NormalBandwidth,omitempty" name:"NormalBandwidth"`

	// Number of forwarding rules
	ForwardRulesLimit *uint64 `json:"ForwardRulesLimit,omitempty" name:"ForwardRulesLimit"`

	// Auto-renewal status. Valid values:
	// `0`: disabled
	// `1`: enabled
	// ]
	AutoRenewFlag *uint64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`

	// Anti-DDoS Advanced line. Valid values:
	// `1`: BGP
	// `2`: China Telecom
	// `3`: China Unicom
	// `4`: China Mobile
	// `99`: third-party line
	// ]
	Line *uint64 `json:"Line,omitempty" name:"Line"`

	// Elastic protection bandwidth (in Gbps)
	ElasticBandwidth *uint64 `json:"ElasticBandwidth,omitempty" name:"ElasticBandwidth"`
}

type BGPIPInstanceUsages struct {

	// Number of used port rules
	PortRulesUsage *uint64 `json:"PortRulesUsage,omitempty" name:"PortRulesUsage"`

	// Number of used domain name rules
	DomainRulesUsage *uint64 `json:"DomainRulesUsage,omitempty" name:"DomainRulesUsage"`

	// Number of attack times in the last 7 days
	Last7DayAttackCount *uint64 `json:"Last7DayAttackCount,omitempty" name:"Last7DayAttackCount"`
}

type BGPInstance struct {

	// Anti-DDoS instance details
	InstanceDetail *InstanceRelation `json:"InstanceDetail,omitempty" name:"InstanceDetail"`

	// Anti-DDoS instance specifications
	SpecificationLimit *BGPInstanceSpecification `json:"SpecificationLimit,omitempty" name:"SpecificationLimit"`

	// Anti-DDoS instance usage statistics
	Usage *BGPInstanceUsages `json:"Usage,omitempty" name:"Usage"`

	// Region of the Anti-DDoS instance
	Region *RegionInfo `json:"Region,omitempty" name:"Region"`

	// Status of the Anti-DDoS instance. Valid values:
	// `idle`: running
	// `attacking`: under attacks
	// `blocking`: blocked
	// `creating`: creating
	// `deblocking`: unblocked
	// `isolate`: isolated
	Status *string `json:"Status,omitempty" name:"Status"`

	// Purchase Time
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// Expiration time
	ExpiredTime *string `json:"ExpiredTime,omitempty" name:"ExpiredTime"`

	// Name of the Anti-DDoS instance
	Name *string `json:"Name,omitempty" name:"Name"`

	// Package details of the Anti-DDoS instance.
	// Note: This field is `null` for an Anti-DDoS instance without using a package.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	PackInfo *PackInfo `json:"PackInfo,omitempty" name:"PackInfo"`

	// Details of the cloud product used by the EIP bound to the Anti-DDoS Pro instance
	EipProductInfos []*EipProductInfo `json:"EipProductInfos,omitempty" name:"EipProductInfos"`

	// Binding status of the Anti-DDoS Pro instance
	// `idle`: the instance is bound.
	//  `bounding`: the instance is in binding.
	// `failed`: the binding failed.
	// ]
	BoundStatus *string `json:"BoundStatus,omitempty" name:"BoundStatus"`

	// Layer-4 protection level
	DDoSLevel *string `json:"DDoSLevel,omitempty" name:"DDoSLevel"`

	// CC protection switch
	CCEnable *int64 `json:"CCEnable,omitempty" name:"CCEnable"`
}

type BGPInstanceSpecification struct {

	// Base protection bandwidth (in Gbps)
	ProtectBandwidth *uint64 `json:"ProtectBandwidth,omitempty" name:"ProtectBandwidth"`

	// Number of protection chances
	ProtectCountLimit *uint64 `json:"ProtectCountLimit,omitempty" name:"ProtectCountLimit"`

	// Number of protection IPs
	ProtectIPNumberLimit *uint64 `json:"ProtectIPNumberLimit,omitempty" name:"ProtectIPNumberLimit"`

	// Auto-renewal status. Valid values:
	// `0`: disabled
	// `1`: enabled
	// ]
	AutoRenewFlag *uint64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`
}

type BGPInstanceUsages struct {

	// Number of used protection chances
	ProtectCountUsage *uint64 `json:"ProtectCountUsage,omitempty" name:"ProtectCountUsage"`

	// Number of protected IPs
	ProtectIPNumberUsage *uint64 `json:"ProtectIPNumberUsage,omitempty" name:"ProtectIPNumberUsage"`

	// Number of attack times in the last 7 days
	Last7DayAttackCount *uint64 `json:"Last7DayAttackCount,omitempty" name:"Last7DayAttackCount"`
}

type BlackWhiteIpRelation struct {

	// IP address
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// IP type. Valid values: `black` (blocklisted IP), `white`(allowlisted IP).
	Type *string `json:"Type,omitempty" name:"Type"`

	// Anti-DDoS instance configured
	InstanceDetailList []*InstanceRelation `json:"InstanceDetailList,omitempty" name:"InstanceDetailList"`

	// IP mask. `0` indicates a 32-bit IP.
	Mask *uint64 `json:"Mask,omitempty" name:"Mask"`
}

type BoundIpInfo struct {

	// IP address
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// Category of product that can be bound. Valid values: `public` (CVM and CLB), `bm` (BM), `eni` (ENI), `vpngw` (VPN gateway), `natgw` (NAT gateway), `waf` (WAF), `fpc` (financial products), `gaap` (GAAP), and `other` (hosted IP).
	BizType *string `json:"BizType,omitempty" name:"BizType"`

	// Anti-DDoS instance ID of the IP. This field is required if the instance ID is bound to a new IP. For example, this field InstanceId will be `eni-*` if the instance ID is bound to an ENI IP; `none` if there is no instance ID to bind to a hosted IP.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Sub-product category. Valid values: `cvm` (CVM), `lb` (Load balancer), `eni` (ENI), `vpngw` (VPN gateway), `natgw` (NAT gateway), `waf` (WAF), `fpc` (financial products), `gaap` (GAAP), `eip` (BM EIP) and `other` (hosted IP).
	DeviceType *string `json:"DeviceType,omitempty" name:"DeviceType"`

	// ISP. Valid values: `0` (China Telecom), `1` (China Unicom), `2` (China Mobile),`5` (BGP).
	IspCode *uint64 `json:"IspCode,omitempty" name:"IspCode"`
}

type CertIdInsL7Rules struct {

	// List of rules configured for certificates
	L7Rules []*InsL7Rules `json:"L7Rules,omitempty" name:"L7Rules"`

	// Certificate ID
	CertId *string `json:"CertId,omitempty" name:"CertId"`
}

type CreateBlackWhiteIpListRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// List of IPs
	IpList []*string `json:"IpList,omitempty" name:"IpList"`

	// IP type. Valid values: `black` (blocklisted IP), `white`(allowlisted IP).
	Type *string `json:"Type,omitempty" name:"Type"`
}

func (r *CreateBlackWhiteIpListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBlackWhiteIpListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "IpList")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBlackWhiteIpListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateBlackWhiteIpListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateBlackWhiteIpListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBlackWhiteIpListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateBoundIPRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS service type. `bgp`: Anti-DDoS Pro (Single IP); `bgp-multip`: Anti-DDoS Pro (Multi-IP)
	Business *string `json:"Business,omitempty" name:"Business"`

	// Anti-DDoS instance ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// Array of IPs to bind to the Anti-DDoS instance. For Anti-DDoS Pro Single IP instance, the array contains only one IP. If there are no IPs to bind, it is empty; however, either `BoundDevList` or `UnBoundDevList` must not be empty.
	BoundDevList []*BoundIpInfo `json:"BoundDevList,omitempty" name:"BoundDevList"`

	// Array of IPs to unbind from the Anti-DDoS instance. For Anti-DDoS Pro Single IP instance, the array contains only one IP; if there are no IPs to unbind, it is empty; however, either `BoundDevList` or `UnBoundDevList` must not be empty.
	UnBoundDevList []*BoundIpInfo `json:"UnBoundDevList,omitempty" name:"UnBoundDevList"`

	// Disused
	CopyPolicy *string `json:"CopyPolicy,omitempty" name:"CopyPolicy"`
}

func (r *CreateBoundIPRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBoundIPRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "BoundDevList")
	delete(f, "UnBoundDevList")
	delete(f, "CopyPolicy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBoundIPRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateBoundIPResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Success code
		Success *SuccessCode `json:"Success,omitempty" name:"Success"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateBoundIPResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBoundIPResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDDoSAIRequest struct {
	*tchttp.BaseRequest

	// List of Anti-DDoS instance IDs
	InstanceIdList []*string `json:"InstanceIdList,omitempty" name:"InstanceIdList"`

	// AI protection switch. Valid values:
	// `on`: enabled
	// `off`: disabled
	// ]
	DDoSAI *string `json:"DDoSAI,omitempty" name:"DDoSAI"`
}

func (r *CreateDDoSAIRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDDoSAIRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIdList")
	delete(f, "DDoSAI")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDDoSAIRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateDDoSAIResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDDoSAIResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDDoSAIResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDDoSGeoIPBlockConfigRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Region blocking configuration. The configuration ID should be cleared when you set this parameter.
	DDoSGeoIPBlockConfig *DDoSGeoIPBlockConfig `json:"DDoSGeoIPBlockConfig,omitempty" name:"DDoSGeoIPBlockConfig"`
}

func (r *CreateDDoSGeoIPBlockConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDDoSGeoIPBlockConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "DDoSGeoIPBlockConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDDoSGeoIPBlockConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateDDoSGeoIPBlockConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDDoSGeoIPBlockConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDDoSGeoIPBlockConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDDoSSpeedLimitConfigRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Access rate limit configuration. The configuration ID should be cleared when you set this parameter.
	DDoSSpeedLimitConfig *DDoSSpeedLimitConfig `json:"DDoSSpeedLimitConfig,omitempty" name:"DDoSSpeedLimitConfig"`
}

func (r *CreateDDoSSpeedLimitConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDDoSSpeedLimitConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "DDoSSpeedLimitConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDDoSSpeedLimitConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateDDoSSpeedLimitConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDDoSSpeedLimitConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDDoSSpeedLimitConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDefaultAlarmThresholdRequest struct {
	*tchttp.BaseRequest

	// Default alarm threshold configuration
	DefaultAlarmConfig *DefaultAlarmThreshold `json:"DefaultAlarmConfig,omitempty" name:"DefaultAlarmConfig"`

	// Product category. Valid values:
	// `bgp`: Anti-DDoS Pro
	// `bgpip`: Anti-DDoS Advanced
	// ]
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
}

func (r *CreateDefaultAlarmThresholdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDefaultAlarmThresholdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DefaultAlarmConfig")
	delete(f, "InstanceType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDefaultAlarmThresholdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateDefaultAlarmThresholdResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDefaultAlarmThresholdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDefaultAlarmThresholdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateIPAlarmThresholdConfigRequest struct {
	*tchttp.BaseRequest

	// List of IP alarm threshold configurations
	IpAlarmThresholdConfigList []*IPAlarmThresholdRelation `json:"IpAlarmThresholdConfigList,omitempty" name:"IpAlarmThresholdConfigList"`
}

func (r *CreateIPAlarmThresholdConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIPAlarmThresholdConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IpAlarmThresholdConfigList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateIPAlarmThresholdConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateIPAlarmThresholdConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateIPAlarmThresholdConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIPAlarmThresholdConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateL7RuleCertsRequest struct {
	*tchttp.BaseRequest

	// SSL certificate ID
	CertId *string `json:"CertId,omitempty" name:"CertId"`

	// List of Layer-7 domain name forwarding rules
	L7Rules []*InsL7Rules `json:"L7Rules,omitempty" name:"L7Rules"`
}

func (r *CreateL7RuleCertsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateL7RuleCertsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertId")
	delete(f, "L7Rules")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateL7RuleCertsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateL7RuleCertsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Success code
		Success *SuccessCode `json:"Success,omitempty" name:"Success"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateL7RuleCertsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateL7RuleCertsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreatePacketFilterConfigRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Feature filtering rules
	PacketFilterConfig *PacketFilterConfig `json:"PacketFilterConfig,omitempty" name:"PacketFilterConfig"`
}

func (r *CreatePacketFilterConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePacketFilterConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "PacketFilterConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePacketFilterConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreatePacketFilterConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreatePacketFilterConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePacketFilterConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateProtocolBlockConfigRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Protocol blocking configuration
	ProtocolBlockConfig *ProtocolBlockConfig `json:"ProtocolBlockConfig,omitempty" name:"ProtocolBlockConfig"`
}

func (r *CreateProtocolBlockConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProtocolBlockConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ProtocolBlockConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateProtocolBlockConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateProtocolBlockConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateProtocolBlockConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProtocolBlockConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSchedulingDomainRequest struct {
	*tchttp.BaseRequest
}

func (r *CreateSchedulingDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSchedulingDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSchedulingDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateSchedulingDomainResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Created domain name
		Domain *string `json:"Domain,omitempty" name:"Domain"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSchedulingDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSchedulingDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateWaterPrintConfigRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Watermark configuration
	WaterPrintConfig *WaterPrintConfig `json:"WaterPrintConfig,omitempty" name:"WaterPrintConfig"`
}

func (r *CreateWaterPrintConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWaterPrintConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "WaterPrintConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateWaterPrintConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateWaterPrintConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateWaterPrintConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWaterPrintConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateWaterPrintKeyRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *CreateWaterPrintKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWaterPrintKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateWaterPrintKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateWaterPrintKeyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateWaterPrintKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWaterPrintKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DDoSAIRelation struct {

	// AI protection switch. Valid values:
	// `on`: enabled
	// `off`: disabled
	// ]
	DDoSAI *string `json:"DDoSAI,omitempty" name:"DDoSAI"`

	// Anti-DDoS instance configured
	InstanceDetailList []*InstanceRelation `json:"InstanceDetailList,omitempty" name:"InstanceDetailList"`
}

type DDoSGeoIPBlockConfig struct {

	// Region type. Valid values:
	// oversea: outside the Chinese mainland
	// `china`: the Chinese mainland
	// `customized`: custom region
	// ]
	RegionType *string `json:"RegionType,omitempty" name:"RegionType"`

	// Blocking action. Valid values:
	// `drop`: the request is blocked.
	// `trans`: the request is allowed.
	// ]
	Action *string `json:"Action,omitempty" name:"Action"`

	// Configuration ID, which is generated after a configuration is added. This field is only required to modify or delete a configuration.
	Id *string `json:"Id,omitempty" name:"Id"`

	// When `RegionType = customized`, AreaList is required and contains up to 128 areas.
	AreaList []*int64 `json:"AreaList,omitempty" name:"AreaList"`
}

type DDoSGeoIPBlockConfigRelation struct {

	// Anti-DDoS region blocking configuration
	GeoIPBlockConfig *DDoSGeoIPBlockConfig `json:"GeoIPBlockConfig,omitempty" name:"GeoIPBlockConfig"`

	// Anti-DDoS instance configured
	InstanceDetailList []*InstanceRelation `json:"InstanceDetailList,omitempty" name:"InstanceDetailList"`
}

type DDoSSpeedLimitConfig struct {

	// Rate limit mode. Valid values:
	// `1`: rate limit based on the real server IP
	// `2`: rate limit based on the destination port
	// ]
	Mode *uint64 `json:"Mode,omitempty" name:"Mode"`

	// Rate limit value. This field contains at least one valid rate limit type. Note that only up to one value of each type is supported.
	SpeedValues []*SpeedValue `json:"SpeedValues,omitempty" name:"SpeedValues"`

	// This field is replaced with a new field DstPortList.
	DstPortScopes []*PortSegment `json:"DstPortScopes,omitempty" name:"DstPortScopes"`

	// 
	Id *string `json:"Id,omitempty" name:"Id"`

	// IP protocol number. Valid values:
	// `ALL`: all protocols
	// `TCP`: TCP protocol
	// `UDP`: UDP protocol
	// `SMP`: SMP protocol
	// `1;2–100`: user-defined protocol with up to 8 ranges
	// ]
	// Note: For custom protocol ranges, only protocol number is supported. Multiple ranges are separated by ";". If the value is `ALL`, any other protocol or protocol number should be excluded.
	ProtocolList *string `json:"ProtocolList,omitempty" name:"ProtocolList"`

	// Port range list, which contains up to 8 ranges. Use ";" to separate multiple ports and "–" to indicate a range of ports, as described in the following formats: `0–65535`, `80;443;1000–2000`.
	DstPortList *string `json:"DstPortList,omitempty" name:"DstPortList"`
}

type DDoSSpeedLimitConfigRelation struct {

	// Anti-DDoS access rate limit configuration
	SpeedLimitConfig *DDoSSpeedLimitConfig `json:"SpeedLimitConfig,omitempty" name:"SpeedLimitConfig"`

	// Anti-DDoS instance configured
	InstanceDetailList []*InstanceRelation `json:"InstanceDetailList,omitempty" name:"InstanceDetailList"`
}

type DefaultAlarmThreshold struct {

	// Alarm threshold type. Valid values:
	// `1`: alarm threshold for inbound traffic
	// `2`: alarm threshold for cleansing attack traffic
	// ]
	AlarmType *uint64 `json:"AlarmType,omitempty" name:"AlarmType"`

	// Alarm threshold (Mbps). The value should be greater than or equal to 0. Note that the alarm threshold configuration will be removed if you pass the parameter for input and set it to 0.
	AlarmThreshold *uint64 `json:"AlarmThreshold,omitempty" name:"AlarmThreshold"`
}

type DeleteBlackWhiteIpListRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// List of IPs
	IpList []*string `json:"IpList,omitempty" name:"IpList"`

	// IP type. Valid values: `black` (blocklisted IP), `white`(allowlisted IP).
	Type *string `json:"Type,omitempty" name:"Type"`
}

func (r *DeleteBlackWhiteIpListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBlackWhiteIpListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "IpList")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteBlackWhiteIpListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteBlackWhiteIpListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteBlackWhiteIpListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBlackWhiteIpListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDDoSGeoIPBlockConfigRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Region blocking configuration. The configuration ID cannot be empty when you set this parameter.
	DDoSGeoIPBlockConfig *DDoSGeoIPBlockConfig `json:"DDoSGeoIPBlockConfig,omitempty" name:"DDoSGeoIPBlockConfig"`
}

func (r *DeleteDDoSGeoIPBlockConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDDoSGeoIPBlockConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "DDoSGeoIPBlockConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDDoSGeoIPBlockConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDDoSGeoIPBlockConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteDDoSGeoIPBlockConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDDoSGeoIPBlockConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDDoSSpeedLimitConfigRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Access rate limit configuration. The configuration ID cannot be empty when you set this parameter.
	DDoSSpeedLimitConfig *DDoSSpeedLimitConfig `json:"DDoSSpeedLimitConfig,omitempty" name:"DDoSSpeedLimitConfig"`
}

func (r *DeleteDDoSSpeedLimitConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDDoSSpeedLimitConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "DDoSSpeedLimitConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDDoSSpeedLimitConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDDoSSpeedLimitConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteDDoSSpeedLimitConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDDoSSpeedLimitConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeletePacketFilterConfigRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Feature filtering configuration
	PacketFilterConfig *PacketFilterConfig `json:"PacketFilterConfig,omitempty" name:"PacketFilterConfig"`
}

func (r *DeletePacketFilterConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePacketFilterConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "PacketFilterConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeletePacketFilterConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeletePacketFilterConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeletePacketFilterConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePacketFilterConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteWaterPrintConfigRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DeleteWaterPrintConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteWaterPrintConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteWaterPrintConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteWaterPrintConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteWaterPrintConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteWaterPrintConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteWaterPrintKeyRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Watermark key ID
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
}

func (r *DeleteWaterPrintKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteWaterPrintKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "KeyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteWaterPrintKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteWaterPrintKeyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteWaterPrintKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteWaterPrintKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBlackWhiteIpListRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeBlackWhiteIpListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBlackWhiteIpListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBlackWhiteIpListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBlackWhiteIpListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// IP blocklist
		BlackIpList []*string `json:"BlackIpList,omitempty" name:"BlackIpList"`

		// IP allowlist
		WhiteIpList []*string `json:"WhiteIpList,omitempty" name:"WhiteIpList"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBlackWhiteIpListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBlackWhiteIpListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDefaultAlarmThresholdRequest struct {
	*tchttp.BaseRequest

	// Product category. Valid values:
	// `bgp`: Anti-DDoS Pro
	// `bgpip`: Anti-DDoS Advanced
	// ]
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// Alarm threshold type filter. Valid values:
	// `1`: alarm threshold for inbound traffic
	// `2`: alarm threshold for cleansing attack traffic
	// ]
	FilterAlarmType *int64 `json:"FilterAlarmType,omitempty" name:"FilterAlarmType"`
}

func (r *DescribeDefaultAlarmThresholdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDefaultAlarmThresholdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceType")
	delete(f, "FilterAlarmType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDefaultAlarmThresholdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDefaultAlarmThresholdResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Default alarm threshold configuration
		DefaultAlarmConfigList []*DefaultAlarmThreshold `json:"DefaultAlarmConfigList,omitempty" name:"DefaultAlarmConfigList"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDefaultAlarmThresholdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDefaultAlarmThresholdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeL7RulesBySSLCertIdRequest struct {
	*tchttp.BaseRequest

	// Domain name status. Valid values: `bindable`, `binded`, `opened`, `closed`, `all` (all states).
	Status *string `json:"Status,omitempty" name:"Status"`

	// List of certificate IDs
	CertIds []*string `json:"CertIds,omitempty" name:"CertIds"`
}

func (r *DescribeL7RulesBySSLCertIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeL7RulesBySSLCertIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Status")
	delete(f, "CertIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeL7RulesBySSLCertIdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeL7RulesBySSLCertIdResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Certificate rule set
		CertSet []*CertIdInsL7Rules `json:"CertSet,omitempty" name:"CertSet"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeL7RulesBySSLCertIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeL7RulesBySSLCertIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeListBGPIPInstancesRequest struct {
	*tchttp.BaseRequest

	// Starting offset of the page. Value: (number of pages – 1) * items per page.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of items per page. The default value is 20 when `Limit = 0`. The maximum value is 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// IP filter
	FilterIp *string `json:"FilterIp,omitempty" name:"FilterIp"`

	// Anti-DDoS instance ID filter. For example, you can filter the Anti-DDoS Advanced instance ID by `bgpip-00000001`.
	FilterInstanceId *string `json:"FilterInstanceId,omitempty" name:"FilterInstanceId"`

	// Anti-DDoS Advanced line filter. Valid values:
	// `1`: BGP line
	// `2`: China Telecom
	// `3`: China Unicom
	// `4`: China Mobile
	// `99`: third-party line
	// ]
	FilterLine *uint64 `json:"FilterLine,omitempty" name:"FilterLine"`

	// Region filter. For example, `ap-guangzhou`.
	FilterRegion *string `json:"FilterRegion,omitempty" name:"FilterRegion"`

	// Name filter
	FilterName *string `json:"FilterName,omitempty" name:"FilterName"`

	// Whether to obtain only Anti-DDoS EIP instances. `1`: Yes; `0`: No.
	FilterEipType *int64 `json:"FilterEipType,omitempty" name:"FilterEipType"`

	// Anti-DDoS Advanced instance binding status filter. Valid values: `BINDING`, `BIND`, `UNBINDING`, `UNBIND`. This filter is only valid when `FilterEipType = 1`.
	FilterEipEipAddressStatus []*string `json:"FilterEipEipAddressStatus,omitempty" name:"FilterEipEipAddressStatus"`
}

func (r *DescribeListBGPIPInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListBGPIPInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "FilterIp")
	delete(f, "FilterInstanceId")
	delete(f, "FilterLine")
	delete(f, "FilterRegion")
	delete(f, "FilterName")
	delete(f, "FilterEipType")
	delete(f, "FilterEipEipAddressStatus")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeListBGPIPInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeListBGPIPInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of lists
		Total *uint64 `json:"Total,omitempty" name:"Total"`

		// List of Anti-DDoS Advanced instances
		InstanceList []*BGPIPInstance `json:"InstanceList,omitempty" name:"InstanceList"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeListBGPIPInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListBGPIPInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeListBGPInstancesRequest struct {
	*tchttp.BaseRequest

	// Starting offset of the page. Value: (number of pages – 1) * items per page.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of items per page. The default value is 20 when `Limit = 0`. The maximum value is 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// IP filter
	FilterIp *string `json:"FilterIp,omitempty" name:"FilterIp"`

	// Anti-DDoS instance ID filter. For example, `bgp-00000001`.
	FilterInstanceId *string `json:"FilterInstanceId,omitempty" name:"FilterInstanceId"`

	// Region filter. For example, `ap-guangzhou`.
	FilterRegion *string `json:"FilterRegion,omitempty" name:"FilterRegion"`

	// Name filter
	FilterName *string `json:"FilterName,omitempty" name:"FilterName"`

	// Line filter. Valid values: 1: BGP; 2: Non-BGP.
	FilterLine *uint64 `json:"FilterLine,omitempty" name:"FilterLine"`
}

func (r *DescribeListBGPInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListBGPInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "FilterIp")
	delete(f, "FilterInstanceId")
	delete(f, "FilterRegion")
	delete(f, "FilterName")
	delete(f, "FilterLine")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeListBGPInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeListBGPInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of lists
		Total *uint64 `json:"Total,omitempty" name:"Total"`

		// List of Anti-DDoS Pro instances
		InstanceList []*BGPInstance `json:"InstanceList,omitempty" name:"InstanceList"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeListBGPInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListBGPInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeListBlackWhiteIpListRequest struct {
	*tchttp.BaseRequest

	// Starting offset of the page. Value: (number of pages – 1) * items per page.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of items per page. The default value is 20 when Limit = 0. The maximum value is 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Anti-DDoS instance ID filter. Anti-DDoS instance prefix wildcard search is supported. For example, you can filter Anti-DDoS Pro instances by `bgp-*`.
	FilterInstanceId *string `json:"FilterInstanceId,omitempty" name:"FilterInstanceId"`

	// IP filter
	FilterIp *string `json:"FilterIp,omitempty" name:"FilterIp"`
}

func (r *DescribeListBlackWhiteIpListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListBlackWhiteIpListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "FilterInstanceId")
	delete(f, "FilterIp")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeListBlackWhiteIpListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeListBlackWhiteIpListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of lists
		Total *int64 `json:"Total,omitempty" name:"Total"`

		// IP blocklist/allowlist
		IpList []*BlackWhiteIpRelation `json:"IpList,omitempty" name:"IpList"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeListBlackWhiteIpListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListBlackWhiteIpListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeListDDoSAIRequest struct {
	*tchttp.BaseRequest

	// Starting offset of the page. Value: (number of pages – 1) * items per page.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of items per page. The default value is 20 when Limit = 0. The maximum value is 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Anti-DDoS instance ID filter. Anti-DDoS instance prefix wildcard search is supported. For example, you can filter Anti-DDoS Pro instances by `bgp-*`.
	FilterInstanceId *string `json:"FilterInstanceId,omitempty" name:"FilterInstanceId"`

	// IP filter
	FilterIp *string `json:"FilterIp,omitempty" name:"FilterIp"`
}

func (r *DescribeListDDoSAIRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListDDoSAIRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "FilterInstanceId")
	delete(f, "FilterIp")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeListDDoSAIRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeListDDoSAIResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of lists
		Total *int64 `json:"Total,omitempty" name:"Total"`

		// List of AI protection switches
		ConfigList []*DDoSAIRelation `json:"ConfigList,omitempty" name:"ConfigList"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeListDDoSAIResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListDDoSAIResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeListDDoSGeoIPBlockConfigRequest struct {
	*tchttp.BaseRequest

	// Starting offset of the page. Value: (number of pages – 1) * items per page.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of items per page. The default value is 20 when Limit = 0. The maximum value is 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Anti-DDoS instance ID filter. Anti-DDoS instance prefix wildcard search is supported. For example, you can filter Anti-DDoS Pro instances by `bgp-*`.
	FilterInstanceId *string `json:"FilterInstanceId,omitempty" name:"FilterInstanceId"`

	// IP filter
	FilterIp *string `json:"FilterIp,omitempty" name:"FilterIp"`
}

func (r *DescribeListDDoSGeoIPBlockConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListDDoSGeoIPBlockConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "FilterInstanceId")
	delete(f, "FilterIp")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeListDDoSGeoIPBlockConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeListDDoSGeoIPBlockConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of lists
		Total *uint64 `json:"Total,omitempty" name:"Total"`

		// List of Anti-DDoS region blocking configurations
		ConfigList []*DDoSGeoIPBlockConfigRelation `json:"ConfigList,omitempty" name:"ConfigList"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeListDDoSGeoIPBlockConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListDDoSGeoIPBlockConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeListDDoSSpeedLimitConfigRequest struct {
	*tchttp.BaseRequest

	// Starting offset of the page. Value: (number of pages – 1) * items per page.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of items per page. The default value is 20 when Limit = 0. The maximum value is 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Anti-DDoS instance ID filter. Anti-DDoS instance prefix wildcard search is supported. For example, you can filter Anti-DDoS Pro instances by `bgp-*`.
	FilterInstanceId *string `json:"FilterInstanceId,omitempty" name:"FilterInstanceId"`

	// IP filter
	FilterIp *string `json:"FilterIp,omitempty" name:"FilterIp"`
}

func (r *DescribeListDDoSSpeedLimitConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListDDoSSpeedLimitConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "FilterInstanceId")
	delete(f, "FilterIp")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeListDDoSSpeedLimitConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeListDDoSSpeedLimitConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of lists
		Total *uint64 `json:"Total,omitempty" name:"Total"`

		// List of access rate limit configurations
		ConfigList []*DDoSSpeedLimitConfigRelation `json:"ConfigList,omitempty" name:"ConfigList"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeListDDoSSpeedLimitConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListDDoSSpeedLimitConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeListIPAlarmConfigRequest struct {
	*tchttp.BaseRequest

	// Starting offset of the page. Value: (number of pages – 1) * items per page.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of items per page. The default value is 20 when Limit = 0. The maximum value is 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Anti-DDoS instance ID filter. Anti-DDoS instance prefix wildcard search is supported. For example, you can filter Anti-DDoS Pro instances by `bgp-*`.
	FilterInstanceId *string `json:"FilterInstanceId,omitempty" name:"FilterInstanceId"`

	// Alarm threshold type filter. Valid values:
	// `1`: alarm threshold for inbound traffic
	// `2`: alarm threshold for cleansing attack traffic
	// ]
	FilterAlarmType *int64 `json:"FilterAlarmType,omitempty" name:"FilterAlarmType"`

	// IP filter
	FilterIp *string `json:"FilterIp,omitempty" name:"FilterIp"`
}

func (r *DescribeListIPAlarmConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListIPAlarmConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "FilterInstanceId")
	delete(f, "FilterAlarmType")
	delete(f, "FilterIp")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeListIPAlarmConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeListIPAlarmConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of lists
		Total *int64 `json:"Total,omitempty" name:"Total"`

		// List of IP alarm threshold configurations
		ConfigList []*IPAlarmThresholdRelation `json:"ConfigList,omitempty" name:"ConfigList"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeListIPAlarmConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListIPAlarmConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeListListenerRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeListListenerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListListenerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeListListenerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeListListenerResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// List of layer-4 forwarding listeners
		Layer4Listeners []*Layer4Rule `json:"Layer4Listeners,omitempty" name:"Layer4Listeners"`

		// List of layer-7 forwarding listeners
		Layer7Listeners []*Layer7Rule `json:"Layer7Listeners,omitempty" name:"Layer7Listeners"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeListListenerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListListenerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeListPacketFilterConfigRequest struct {
	*tchttp.BaseRequest

	// Starting offset of the page. Value: (number of pages – 1) * items per page.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of items per page. The default value is 20 when Limit = 0. The maximum value is 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Anti-DDoS instance ID filter. Anti-DDoS instance prefix wildcard search is supported. For example, you can filter Anti-DDoS Pro instances by `bgp-*`.
	FilterInstanceId *string `json:"FilterInstanceId,omitempty" name:"FilterInstanceId"`

	// IP filter
	FilterIp *string `json:"FilterIp,omitempty" name:"FilterIp"`
}

func (r *DescribeListPacketFilterConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListPacketFilterConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "FilterInstanceId")
	delete(f, "FilterIp")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeListPacketFilterConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeListPacketFilterConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of lists
		Total *int64 `json:"Total,omitempty" name:"Total"`

		// Feature filtering configuration
		ConfigList []*PacketFilterRelation `json:"ConfigList,omitempty" name:"ConfigList"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeListPacketFilterConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListPacketFilterConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeListProtectThresholdConfigRequest struct {
	*tchttp.BaseRequest

	// Starting offset of the page. Value: (number of pages – 1) * items per page.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of items per page. The default value is 20 when `Limit = 0`. The maximum value is 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Anti-DDoS instance ID filter. Anti-DDoS instance prefix wildcard search is supported. For example, you can filter Anti-DDoS Pro instances by `bgp-*`.
	FilterInstanceId *string `json:"FilterInstanceId,omitempty" name:"FilterInstanceId"`

	// IP filter
	FilterIp *string `json:"FilterIp,omitempty" name:"FilterIp"`

	// Domain name filter for querying CC protection thresholds of domain names and protocols
	FilterDomain *string `json:"FilterDomain,omitempty" name:"FilterDomain"`

	// Protocol filter for querying CC protection thresholds of domain names and protocols
	FilterProtocol *string `json:"FilterProtocol,omitempty" name:"FilterProtocol"`
}

func (r *DescribeListProtectThresholdConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListProtectThresholdConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "FilterInstanceId")
	delete(f, "FilterIp")
	delete(f, "FilterDomain")
	delete(f, "FilterProtocol")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeListProtectThresholdConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeListProtectThresholdConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of lists
		Total *uint64 `json:"Total,omitempty" name:"Total"`

		// List of protection threshold configurations
		ConfigList []*ProtectThresholdRelation `json:"ConfigList,omitempty" name:"ConfigList"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeListProtectThresholdConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListProtectThresholdConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeListProtocolBlockConfigRequest struct {
	*tchttp.BaseRequest

	// Starting offset of the page. Value: (number of pages – 1) * items per page.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of items per page. The default value is 20 when `Limit = 0`. The maximum value is 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Anti-DDoS instance ID filter. Anti-DDoS instance prefix wildcard search is supported. For example, you can filter Anti-DDoS Pro instances by `bgp-*`.
	FilterInstanceId *string `json:"FilterInstanceId,omitempty" name:"FilterInstanceId"`

	// IP filter
	FilterIp *string `json:"FilterIp,omitempty" name:"FilterIp"`
}

func (r *DescribeListProtocolBlockConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListProtocolBlockConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "FilterInstanceId")
	delete(f, "FilterIp")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeListProtocolBlockConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeListProtocolBlockConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of lists
		Total *int64 `json:"Total,omitempty" name:"Total"`

		// Protocol blocking configuration
		ConfigList []*ProtocolBlockRelation `json:"ConfigList,omitempty" name:"ConfigList"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeListProtocolBlockConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListProtocolBlockConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeListSchedulingDomainRequest struct {
	*tchttp.BaseRequest

	// Starting offset of the page. Value: (number of pages – 1) * items per page.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of items per page. The default value is 20 when `Limit = 0`. The maximum value is 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Scheduling domain name filter
	FilterDomain *string `json:"FilterDomain,omitempty" name:"FilterDomain"`
}

func (r *DescribeListSchedulingDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListSchedulingDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "FilterDomain")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeListSchedulingDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeListSchedulingDomainResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of lists
		Total *uint64 `json:"Total,omitempty" name:"Total"`

		// List of scheduling domain names
		DomainList []*SchedulingDomainInfo `json:"DomainList,omitempty" name:"DomainList"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeListSchedulingDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListSchedulingDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeListWaterPrintConfigRequest struct {
	*tchttp.BaseRequest

	// Starting offset of the page. Value: (number of pages – 1) * items per page.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of items per page. The default value is 20 when `Limit = 0`. The maximum value is 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Anti-DDoS instance ID filter. Anti-DDoS instance prefix wildcard search is supported. For example, you can filter Anti-DDoS Pro instances by `bgp-*`.
	FilterInstanceId *string `json:"FilterInstanceId,omitempty" name:"FilterInstanceId"`

	// IP filter
	FilterIp *string `json:"FilterIp,omitempty" name:"FilterIp"`
}

func (r *DescribeListWaterPrintConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListWaterPrintConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "FilterInstanceId")
	delete(f, "FilterIp")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeListWaterPrintConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeListWaterPrintConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of lists
		Total *int64 `json:"Total,omitempty" name:"Total"`

		// List of watermark configurations
		ConfigList []*WaterPrintRelation `json:"ConfigList,omitempty" name:"ConfigList"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeListWaterPrintConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListWaterPrintConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DisassociateDDoSEipAddressRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS instance ID (only Anti-DDoS Advanced). For example, `bgpip-0000011x`.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// EIP of the Anti-DDoS instance ID
	Eip *string `json:"Eip,omitempty" name:"Eip"`
}

func (r *DisassociateDDoSEipAddressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisassociateDDoSEipAddressRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Eip")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisassociateDDoSEipAddressRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DisassociateDDoSEipAddressResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DisassociateDDoSEipAddressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisassociateDDoSEipAddressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EipAddressPackRelation struct {

	// Number of package IPs
	IpCount *uint64 `json:"IpCount,omitempty" name:"IpCount"`

	// Auto-renewal flag
	AutoRenewFlag *uint64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`

	// Current expiration time
	CurDeadline *string `json:"CurDeadline,omitempty" name:"CurDeadline"`
}

type EipAddressRelation struct {

	// Region of the Anti-DDoS instance bound to the EIP. For example, hk indicates Hong Kong.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	EipAddressRegion *string `json:"EipAddressRegion,omitempty" name:"EipAddressRegion"`

	// ID of the bound resource. For example, an ID may be bound to an CVM instance.
	// Note: This is field may return `null`, indicating that no valid value can be obtained.
	EipBoundRscIns *string `json:"EipBoundRscIns,omitempty" name:"EipBoundRscIns"`

	// ID of the bound ENI
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	EipBoundRscEni *string `json:"EipBoundRscEni,omitempty" name:"EipBoundRscEni"`

	// Private IP of the bound resource
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	EipBoundRscVip *string `json:"EipBoundRscVip,omitempty" name:"EipBoundRscVip"`

	// Modification time
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
}

type EipProductInfo struct {

	// IP address
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// Cloud product category. Valid values:
	// `public`: CVM
	// `bm`: BM
	// `eni`: ENI
	// `vpngw`: VPN gateway
	//  `natgw`: NAT gateway
	// `waf`: WAF
	// `fpc`: financial products
	// `gaap`: GAAP 
	// `other`: hosted IP
	// ]
	BizType *string `json:"BizType,omitempty" name:"BizType"`

	// Cloud sub-product category. Valid values: `cvm` (CVM), `lb` (Load balancer), `eni` (ENI), `vpngw` (VPN gateway), `natgw` (NAT gateway), `waf` (WAF), `fpc` (financial products), `gaap` (GAAP), `eip` (BM EIP) and `other` (hosted IP).
	DeviceType *string `json:"DeviceType,omitempty" name:"DeviceType"`

	// Cloud instance ID of the IP. This field InstanceId will be `eni-*` if the instance ID is bound to an ENI IP; `none` if there is no instance ID to bind to a hosted IP.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type ForwardListener struct {

	// The starting port for listener forwarding. Value range: 1 to 65535.
	FrontendPort *int64 `json:"FrontendPort,omitempty" name:"FrontendPort"`

	// Forwarding protocol. Valid values:
	// `TCP`
	// `UDP`
	// ]
	ForwardProtocol *string `json:"ForwardProtocol,omitempty" name:"ForwardProtocol"`

	// The ending port for listener forwarding. Value range: 1 to 65535.
	FrontendPortEnd *int64 `json:"FrontendPortEnd,omitempty" name:"FrontendPortEnd"`
}

type IPAlarmThresholdRelation struct {

	// Alarm threshold type. Valid values:
	// `1`: alarm threshold for inbound traffic
	// `2`: alarm threshold for cleansing attack traffic
	// ]
	AlarmType *uint64 `json:"AlarmType,omitempty" name:"AlarmType"`

	// Alarm threshold (Mbps). The value should be greater than or equal to 0. Note that the alarm threshold configuration will be removed if you pass the parameter for input and set it to 0.
	AlarmThreshold *uint64 `json:"AlarmThreshold,omitempty" name:"AlarmThreshold"`

	// Anti-DDoS instance configured
	InstanceDetailList []*InstanceRelation `json:"InstanceDetailList,omitempty" name:"InstanceDetailList"`
}

type IPLineInfo struct {

	// IP line type. Valid values:
	// `bgp`: BGP IP
	// `ctcc`: CTCC IP
	// `cucc`: CUCC IP
	// `cmcc`: CMCC IP
	// `abroad`: IP outside the Chinese mainland
	// ]
	Type *string `json:"Type,omitempty" name:"Type"`

	// 
	Eip *string `json:"Eip,omitempty" name:"Eip"`
}

type InsL7Rules struct {

	// Rule status. Valid values: `0` (the rule is working), `1` (the rule goes into effect), `2` (rule configuration failed), `3` (the rule is being deleted), `5` (rule deletion failed), `6` (waiting to add rules), `7` (waiting to delete rules), `8` (waiting to upload certificates), `9` (resources for the rule not found), `10` (waiting to modify rules), `11` (the rule is being modifying).
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// Domain name
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Protocol
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Instance ID
	InsId *string `json:"InsId,omitempty" name:"InsId"`

	// User App ID
	AppId *string `json:"AppId,omitempty" name:"AppId"`

	// High-defense port
	VirtualPort *string `json:"VirtualPort,omitempty" name:"VirtualPort"`

	// Certificate ID
	SSLId *string `json:"SSLId,omitempty" name:"SSLId"`
}

type InstanceRelation struct {

	// Instance IP
	EipList []*string `json:"EipList,omitempty" name:"EipList"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type Layer4Rule struct {

	// Real server port. Value range: 1–65535.
	BackendPort *uint64 `json:"BackendPort,omitempty" name:"BackendPort"`

	// Forwarding port. Value range: 1–65535.
	FrontendPort *uint64 `json:"FrontendPort,omitempty" name:"FrontendPort"`

	// Forwarding rule. Valid values:
	// TCP
	// UDP
	// ]
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// List of real servers
	RealServers []*SourceServer `json:"RealServers,omitempty" name:"RealServers"`

	// Anti-DDoS instance configured
	InstanceDetails []*InstanceRelation `json:"InstanceDetails,omitempty" name:"InstanceDetails"`
}

type Layer7Rule struct {

	// Domain name
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// List of forwarding types
	ProxyTypeList []*ProxyTypeInfo `json:"ProxyTypeList,omitempty" name:"ProxyTypeList"`

	// List of real servers
	RealServers []*SourceServer `json:"RealServers,omitempty" name:"RealServers"`

	// Anti-DDoS instance configured
	InstanceDetails []*InstanceRelation `json:"InstanceDetails,omitempty" name:"InstanceDetails"`
}

type ListenerCcThreholdConfig struct {

	// Domain name
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Protocol. Value: htttps
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Status. Valid values: `0` (disabled), `1` (enabled).
	CCEnable *int64 `json:"CCEnable,omitempty" name:"CCEnable"`

	// CC protection threshold
	CCThreshold *int64 `json:"CCThreshold,omitempty" name:"CCThreshold"`
}

type ModifyDDoSGeoIPBlockConfigRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Region blocking configuration. The configuration ID cannot be empty when you set this parameter.
	DDoSGeoIPBlockConfig *DDoSGeoIPBlockConfig `json:"DDoSGeoIPBlockConfig,omitempty" name:"DDoSGeoIPBlockConfig"`
}

func (r *ModifyDDoSGeoIPBlockConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDDoSGeoIPBlockConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "DDoSGeoIPBlockConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDDoSGeoIPBlockConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDDoSGeoIPBlockConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDDoSGeoIPBlockConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDDoSGeoIPBlockConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDDoSSpeedLimitConfigRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Access rate limit configuration. The configuration ID cannot be empty when you set this parameter.
	DDoSSpeedLimitConfig *DDoSSpeedLimitConfig `json:"DDoSSpeedLimitConfig,omitempty" name:"DDoSSpeedLimitConfig"`
}

func (r *ModifyDDoSSpeedLimitConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDDoSSpeedLimitConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "DDoSSpeedLimitConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDDoSSpeedLimitConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDDoSSpeedLimitConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDDoSSpeedLimitConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDDoSSpeedLimitConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDomainUsrNameRequest struct {
	*tchttp.BaseRequest

	// User CNAME
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// Domain name
	DomainUserName *string `json:"DomainUserName,omitempty" name:"DomainUserName"`
}

func (r *ModifyDomainUsrNameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDomainUsrNameRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DomainName")
	delete(f, "DomainUserName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDomainUsrNameRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDomainUsrNameResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDomainUsrNameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDomainUsrNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyPacketFilterConfigRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Feature filtering configuration
	PacketFilterConfig *PacketFilterConfig `json:"PacketFilterConfig,omitempty" name:"PacketFilterConfig"`
}

func (r *ModifyPacketFilterConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPacketFilterConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "PacketFilterConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyPacketFilterConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyPacketFilterConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyPacketFilterConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPacketFilterConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PackInfo struct {

	// Package type. Valid values:
	// `staticpack`: non-BGP package
	// `insurance`: guarantee package
	// ]
	PackType *string `json:"PackType,omitempty" name:"PackType"`

	// Package ID
	PackId *string `json:"PackId,omitempty" name:"PackId"`
}

type PacketFilterConfig struct {

	// Protocol. Valid values: `tcp`, `udp`, `icmp`, `all`.
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Start source port. Value range: 0–65535.
	SportStart *int64 `json:"SportStart,omitempty" name:"SportStart"`

	// End source port. Value range: 0–65535. The value also should be greater than or equal to that of the start source port.
	SportEnd *int64 `json:"SportEnd,omitempty" name:"SportEnd"`

	// Start destination port
	DportStart *int64 `json:"DportStart,omitempty" name:"DportStart"`

	// End destination port. Value range: 1–65535. The value also should be greater than or equal to that of the start source port.
	DportEnd *int64 `json:"DportEnd,omitempty" name:"DportEnd"`

	// Minimum message length. Value range: 1–1500.
	PktlenMin *int64 `json:"PktlenMin,omitempty" name:"PktlenMin"`

	// Maximum message length. Value range: 1–1500. The value also should be greater than or equal to that of the minimum message length.
	PktlenMax *int64 `json:"PktlenMax,omitempty" name:"PktlenMax"`

	// Action. Valid values:
	// `drop`: discards the request.
	// `transmit`: allows the request.
	// `drop_black`: discards the request and adds the IP to blocklist.
	// `drop_rst`: blocks the request.
	// `drop_black_rst`: blocks the request and adds the IP to blocklist.
	// `forward`: continues protection.
	// ]
	Action *string `json:"Action,omitempty" name:"Action"`

	// Detection location:
	// `begin_l3`: IP header
	// `begin_l4`: TCP/UDP header
	// `begin_l5`: T load
	// `no_match`: no match
	// ]
	MatchBegin *string `json:"MatchBegin,omitempty" name:"MatchBegin"`

	// Detection type:
	// `sunday`: keyword
	// `pcre`: regular expression
	// ]
	MatchType *string `json:"MatchType,omitempty" name:"MatchType"`

	// Detection value. Should be in key string or regular expression. 
	// When the `MatchType` is `sunday`, enter a string or a string in hexadecimal byte code representation. For example, a string "123" corresponds to the hexadecimal byte code "\x313233".
	// When the `MatchType` is `pcre`, enter a regular expression.
	// ]
	Str *string `json:"Str,omitempty" name:"Str"`

	// Detection depth starting from the detection position. Value range: [0, 1500].
	Depth *int64 `json:"Depth,omitempty" name:"Depth"`

	// Offset starting from the detection position. Value range: [0, Depth].
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Whether the detection value is included:
	// `0`: included
	// `1`: excluded
	// ]
	IsNot *int64 `json:"IsNot,omitempty" name:"IsNot"`

	// Relationship between the first and second detection conditions:
	// `and`: under both the first and second detection conditions
	// `none`: under only the first detection condition
	// ]
	MatchLogic *string `json:"MatchLogic,omitempty" name:"MatchLogic"`

	// The second detection position:
	// `begin_l5`: load
	// `no_match`: no match
	// ]
	MatchBegin2 *string `json:"MatchBegin2,omitempty" name:"MatchBegin2"`

	// The second detection type:
	// `sunday`: keyword
	// `pcre`: regular expression
	// ]
	MatchType2 *string `json:"MatchType2,omitempty" name:"MatchType2"`

	// The second detection value. Should be in key string or regular expression.
	// When the `MatchType` is `sunday`, enter a string or a string in hexadecimal byte code representation. For example, a string "123" corresponds to the hexadecimal byte code "\x313233".
	// When the `MatchType` is `pcre`, enter a regular expression.
	// ]
	Str2 *string `json:"Str2,omitempty" name:"Str2"`

	// Detection depth starting from the second detection position. Value range: [0, 1500].
	Depth2 *int64 `json:"Depth2,omitempty" name:"Depth2"`

	// Offset starting from the second detection position. Value range: [0, Depth2].
	Offset2 *int64 `json:"Offset2,omitempty" name:"Offset2"`

	// Whether the second detection value is included:
	// `0`: included
	// `1`: excluded
	// ]
	IsNot2 *int64 `json:"IsNot2,omitempty" name:"IsNot2"`

	// A rule ID is generated after a feature filtering configuration is added successfully. Leave this field empty when adding a new feature filtering configuration.
	Id *string `json:"Id,omitempty" name:"Id"`
}

type PacketFilterRelation struct {

	// Feature filtering configuration
	PacketFilterConfig *PacketFilterConfig `json:"PacketFilterConfig,omitempty" name:"PacketFilterConfig"`

	// Anti-DDoS instance configured
	InstanceDetailList []*InstanceRelation `json:"InstanceDetailList,omitempty" name:"InstanceDetailList"`
}

type PortSegment struct {

	// Start port. Value range: 1–65535.
	BeginPort *uint64 `json:"BeginPort,omitempty" name:"BeginPort"`

	// End port. The value should be in the range 1–65535 and cannot be less than that of the start port.
	EndPort *uint64 `json:"EndPort,omitempty" name:"EndPort"`
}

type ProtectThresholdRelation struct {

	// DDoS protection level:
	// `low`: loose protection
	// `middle`: medium protection
	// `high`: strict protection
	// ]
	DDoSLevel *string `json:"DDoSLevel,omitempty" name:"DDoSLevel"`

	// DDoS cleansing threshold (in Mbps)
	DDoSThreshold *uint64 `json:"DDoSThreshold,omitempty" name:"DDoSThreshold"`

	// DDoS AI protection switch:
	// `on`: enabled
	// `off`: disabled
	// ]
	DDoSAI *string `json:"DDoSAI,omitempty" name:"DDoSAI"`

	// CC cleansing switch
	// `0`: disabled
	// `1`: enabled
	// ]
	CCEnable *uint64 `json:"CCEnable,omitempty" name:"CCEnable"`

	// CC cleansing threshold (in QPS)
	CCThreshold *uint64 `json:"CCThreshold,omitempty" name:"CCThreshold"`

	// Anti-DDoS instance configured
	InstanceDetailList []*InstanceRelation `json:"InstanceDetailList,omitempty" name:"InstanceDetailList"`

	// Domain name and protocol protection thresholds
	ListenerCcThresholdList []*ListenerCcThreholdConfig `json:"ListenerCcThresholdList,omitempty" name:"ListenerCcThresholdList"`
}

type ProtocolBlockConfig struct {

	// TCP blocking. Valid values: `0` (disabled), `1`(enabled).
	DropTcp *int64 `json:"DropTcp,omitempty" name:"DropTcp"`

	// UDP blocking. Valid values: `0` (disabled), `1`(enabled).
	DropUdp *int64 `json:"DropUdp,omitempty" name:"DropUdp"`

	// ICMP blocking. Valid values: `0` (disabled), `1`(enabled).
	DropIcmp *int64 `json:"DropIcmp,omitempty" name:"DropIcmp"`

	// Other protocol blocking. Valid values: `0` (disabled), `1`(enabled).
	DropOther *int64 `json:"DropOther,omitempty" name:"DropOther"`

	// Null connection protection. Valid values: `0` (disabled), `1` (enabled).
	CheckExceptNullConnect *int64 `json:"CheckExceptNullConnect,omitempty" name:"CheckExceptNullConnect"`
}

type ProtocolBlockRelation struct {

	// Protocol blocking configuration
	ProtocolBlockConfig *ProtocolBlockConfig `json:"ProtocolBlockConfig,omitempty" name:"ProtocolBlockConfig"`

	// Anti-DDoS instance configured
	InstanceDetailList []*InstanceRelation `json:"InstanceDetailList,omitempty" name:"InstanceDetailList"`
}

type ProxyTypeInfo struct {

	// List of forwarding listening ports. Value range: 1–65535.
	ProxyPorts []*int64 `json:"ProxyPorts,omitempty" name:"ProxyPorts"`

	// Forwarding protocol:
	// `http`: HTTP protocol
	// `https`: HTTPS protocol
	// ]
	ProxyType *string `json:"ProxyType,omitempty" name:"ProxyType"`
}

type RegionInfo struct {

	// Region name, such as `ap-guangzhou`
	Region *string `json:"Region,omitempty" name:"Region"`
}

type SchedulingDomainInfo struct {

	// Scheduling domain name
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// List of line IPs
	LineIPList []*IPLineInfo `json:"LineIPList,omitempty" name:"LineIPList"`

	// Scheduling mode. Valid value: `priority` (priority scheduling).
	Method *string `json:"Method,omitempty" name:"Method"`

	// TTL value recorded from the scheduling domain name resolution
	TTL *uint64 `json:"TTL,omitempty" name:"TTL"`

	// Running status. Valid values:
	// `0`: not running
	// `1`: running
	// `2`: abnormal
	// ]
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// Creation time
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// Last modification time
	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`

	// Domain name
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	UsrDomainName *string `json:"UsrDomainName,omitempty" name:"UsrDomainName"`
}

type SourceServer struct {

	// Types of the real server address, such as IP or domain name.
	RealServer *string `json:"RealServer,omitempty" name:"RealServer"`

	// Types of the real server address:
	// `1`: domain name
	// `2`: IP
	// ]
	RsType *int64 `json:"RsType,omitempty" name:"RsType"`

	// Forward weight of the real server. Value range: 1–100.
	Weight *int64 `json:"Weight,omitempty" name:"Weight"`
}

type SpeedValue struct {

	// Rate limit value types:
	// `1`: packets per second (pps)
	// `2`: bits per second (bps)
	// ]
	Type *uint64 `json:"Type,omitempty" name:"Type"`

	// Value
	Value *uint64 `json:"Value,omitempty" name:"Value"`
}

type StaticPackRelation struct {

	// Base protection bandwidth
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	ProtectBandwidth *uint64 `json:"ProtectBandwidth,omitempty" name:"ProtectBandwidth"`

	// Application bandwidth
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	NormalBandwidth *uint64 `json:"NormalBandwidth,omitempty" name:"NormalBandwidth"`

	// Forwarding rules
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	ForwardRulesLimit *uint64 `json:"ForwardRulesLimit,omitempty" name:"ForwardRulesLimit"`

	// Auto-renewal flag
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	AutoRenewFlag *uint64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`

	// Expiration time
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	CurDeadline *string `json:"CurDeadline,omitempty" name:"CurDeadline"`
}

type SuccessCode struct {

	// Description
	Message *string `json:"Message,omitempty" name:"Message"`

	// Success/Error code
	Code *string `json:"Code,omitempty" name:"Code"`
}

type SwitchWaterPrintConfigRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Watermark status. `1`: enabled; `0`: disabled.
	OpenStatus *int64 `json:"OpenStatus,omitempty" name:"OpenStatus"`
}

func (r *SwitchWaterPrintConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchWaterPrintConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "OpenStatus")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SwitchWaterPrintConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type SwitchWaterPrintConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SwitchWaterPrintConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchWaterPrintConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WaterPrintConfig struct {

	// Watermark offset. Value range: [0, 100).
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Start status. Valid values:
	// `0`: manual start
	// `1`: instant start
	// ]
	OpenStatus *int64 `json:"OpenStatus,omitempty" name:"OpenStatus"`

	// List of forwarding listeners configured
	Listeners []*ForwardListener `json:"Listeners,omitempty" name:"Listeners"`

	// A list of watermark keys is generated after a watermark is added successfully. Each watermark can have up to 2 keys. When there is only one valid key, it cannot be deleted.
	Keys []*WaterPrintKey `json:"Keys,omitempty" name:"Keys"`

	// Watermark checking mode, which can be:
	// `checkall`: normal mode
	// `shortfpcheckall`: compact mode
	// ]
	Verify *string `json:"Verify,omitempty" name:"Verify"`
}

type WaterPrintKey struct {

	// Key version
	KeyVersion *string `json:"KeyVersion,omitempty" name:"KeyVersion"`

	// Key content
	KeyContent *string `json:"KeyContent,omitempty" name:"KeyContent"`

	// Key ID
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

	// Key status. Valid value: `1` (enabled).
	KeyOpenStatus *int64 `json:"KeyOpenStatus,omitempty" name:"KeyOpenStatus"`

	// Key creation time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type WaterPrintRelation struct {

	// Watermark configuration
	WaterPrintConfig *WaterPrintConfig `json:"WaterPrintConfig,omitempty" name:"WaterPrintConfig"`

	// Anti-DDoS instance configured
	InstanceDetailList []*InstanceRelation `json:"InstanceDetailList,omitempty" name:"InstanceDetailList"`
}
