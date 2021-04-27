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

package v20180709

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

type BaradData struct {

	// Metric name (connum: number of active TCP connections;
	// new_conn: number of new TCP connections;
	// inactive_conn: number of inactive connections;
	// intraffic: inbound traffic;
	// outtraffic: outbound traffic;
	// alltraffic: sum of inbound and outbound traffic;
	// inpkg: inbound packet rate;
	// outpkg: outbound packet rate;)
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// Value array
	Data []*float64 `json:"Data,omitempty" name:"Data" list`

	// Value array size
	Count *uint64 `json:"Count,omitempty" name:"Count"`
}

type BoundIpInfo struct {

	// IP address
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// Category of product that can be bound. Valid values: public (CVM and CLB), bm (BM), eni (ENI), vpngw (VPN gateway), natgw (NAT gateway), waf (WAF), fpc (financial products), gaap (GAAP), and other (Hosted IP).
	BizType *string `json:"BizType,omitempty" name:"BizType"`

	// Subtype under product type. Valid values: [cvm (CVM), lb (CLB), eni (ENI), vpngw (VPN), natgw (NAT), waf (WAF), fpc (finance), gaap (GAAP), other (hosted IP), eip (BM EIP)]
	DeviceType *string `json:"DeviceType,omitempty" name:"DeviceType"`

	// Resource instance ID of IP. This field is required when binding a new IP. For example, if it is an ENI IP, enter `ID(eni-*)` of the ENI for `InstanceId`; if it is a hosted IP without corresponding resource instance ID, enter "none";
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type CCAlarmThreshold struct {

	// CC alarm threshold
	AlarmThreshold *uint64 `json:"AlarmThreshold,omitempty" name:"AlarmThreshold"`
}

type CCEventRecord struct {

	// Anti-DDoS service type. `bgpip`: Anti-DDoS Advanced; `bgp`: Anti-DDoS Pro (Single IP); `bgp-multip`: Anti-DDoS Pro (Multi-IP); `net`: Anti-DDoS Ultimate; `basic`: Anti-DDoS Basic
	Business *string `json:"Business,omitempty" name:"Business"`

	// Anti-DDoS instance ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// Resource IP
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// Attack start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Attack end time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Total requests peak (QPS)
	ReqQps *uint64 `json:"ReqQps,omitempty" name:"ReqQps"`

	// Attack peak (QPS)
	DropQps *uint64 `json:"DropQps,omitempty" name:"DropQps"`

	// Attack status. Valid values: [0 (ongoing), 1 (ended)]
	AttackStatus *uint64 `json:"AttackStatus,omitempty" name:"AttackStatus"`

	// Resource name
	// Note: this field may return null, indicating that no valid values can be obtained.
	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`

	// Domain name list
	// Note: this field may return null, indicating that no valid values can be obtained.
	DomainList *string `json:"DomainList,omitempty" name:"DomainList"`

	// URI list
	// Note: this field may return null, indicating that no valid values can be obtained.
	UriList *string `json:"UriList,omitempty" name:"UriList"`

	// Attack source list
	// Note: this field may return null, indicating that no valid values can be obtained.
	AttackipList *string `json:"AttackipList,omitempty" name:"AttackipList"`
}

type CCFrequencyRule struct {

	// ID of the access frequency control rule for CC protection
	CCFrequencyRuleId *string `json:"CCFrequencyRuleId,omitempty" name:"CCFrequencyRuleId"`

	// URI string, which must start with `/`, such as `/abc/a.php`. Length limit: 31. If URI is `/`, only prefix match can be selected as the matching mode;
	Uri *string `json:"Uri,omitempty" name:"Uri"`

	// `User-Agent` string. Length limit: 80
	UserAgent *string `json:"UserAgent,omitempty" name:"UserAgent"`

	// Cookie string. Length limit: 40
	Cookie *string `json:"Cookie,omitempty" name:"Cookie"`

	// Matching rule. Valid values: ["include" (prefix match), "equal" (exact match)]
	Mode *string `json:"Mode,omitempty" name:"Mode"`

	// Reference period in seconds. Valid values: [10, 30, 60]
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// Number of access requests. Value range: [1-10000]
	ReqNumber *uint64 `json:"ReqNumber,omitempty" name:"ReqNumber"`

	// Action take. Valid values: ["alg" (CAPTCHA), "drop" (blocking)]
	Act *string `json:"Act,omitempty" name:"Act"`

	// Execution duration in seconds. Valid range: [1-900]
	ExeDuration *uint64 `json:"ExeDuration,omitempty" name:"ExeDuration"`
}

type CCPolicy struct {

	// Policy name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Matching mode. Valid values: [matching (matching mode), speedlimit (speed limiting mode)]
	Smode *string `json:"Smode,omitempty" name:"Smode"`

	// Policy ID
	SetId *string `json:"SetId,omitempty" name:"SetId"`

	// Number of requests allowed per minute
	Frequency *uint64 `json:"Frequency,omitempty" name:"Frequency"`

	// Executed policy mode. Valid values: [alg (verification code), drop (blocking)]
	ExeMode *string `json:"ExeMode,omitempty" name:"ExeMode"`

	// Specifies whether the policy is activated
	Switch *uint64 `json:"Switch,omitempty" name:"Switch"`

	// Creation time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Rule list
	RuleList []*CCRule `json:"RuleList,omitempty" name:"RuleList" list`

	// IP list. If this field is to be left empty, please pass in an empty instead of null;
	IpList []*string `json:"IpList,omitempty" name:"IpList" list`

	// CC protection type. Valid values: [http, https]
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// ID of the forwarding rule corresponding to the HTTPS CC protection domain name
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// HTTPS CC protection domain name
	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

type CCRule struct {

	// Key of the policy. Valid values: `host`, `cgi`, `ua`, `referer`
	Skey *string `json:"Skey,omitempty" name:"Skey"`

	// Rule condition. Valid values: `include`, `not_include`, `equal`
	Operator *string `json:"Operator,omitempty" name:"Operator"`

	// Value of the policy. Length limit: 31 bytes
	Value *string `json:"Value,omitempty" name:"Value"`
}

type CCRuleConfig struct {

	// Reference period in seconds. Valid values: [10, 30, 60]
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// Number of access requests. Value range: [1-10000]
	ReqNumber *uint64 `json:"ReqNumber,omitempty" name:"ReqNumber"`

	// Action take. Valid values: ["alg" (CAPTCHA), "drop" (blocking)]
	Action *string `json:"Action,omitempty" name:"Action"`

	// Execution duration in seconds. Valid range: [1-900]
	ExeDuration *uint64 `json:"ExeDuration,omitempty" name:"ExeDuration"`
}

type CreateBasicDDoSAlarmThresholdRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS service type (`basic`: Anti-DDoS Basic)
	Business *string `json:"Business,omitempty" name:"Business"`

	// `get`: read alarm threshold, `set`: set alarm threshold
	Method *string `json:"Method,omitempty" name:"Method"`

	// Alarm threshold type. 1: inbound traffic, 2: cleansed traffic. This field is required if `Method` is `set`;
	AlarmType *uint64 `json:"AlarmType,omitempty" name:"AlarmType"`

	// Alarm threshold. It is required if `Method` is `set`. If it is set to 0, it means to clear the alarm threshold configuration;
	AlarmThreshold *uint64 `json:"AlarmThreshold,omitempty" name:"AlarmThreshold"`
}

func (r *CreateBasicDDoSAlarmThresholdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBasicDDoSAlarmThresholdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Method")
	delete(f, "AlarmType")
	delete(f, "AlarmThreshold")
	if len(f) > 0 {
		return errors.New("CreateBasicDDoSAlarmThresholdRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateBasicDDoSAlarmThresholdResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// If there is an alarm threshold configuration, the returned alarm threshold will be greater than 0; otherwise, the returned alarm threshold will be 0;
		AlarmThreshold *uint64 `json:"AlarmThreshold,omitempty" name:"AlarmThreshold"`

		// Alarm threshold type. 1: inbound traffic, 2: cleansed traffic. This field is valid if `AlarmThreshold` is above 0;
		AlarmType *uint64 `json:"AlarmType,omitempty" name:"AlarmType"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateBasicDDoSAlarmThresholdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBasicDDoSAlarmThresholdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateBoundIPRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS service type. `bgp`: Anti-DDoS Pro (Single IP); `bgp-multip`: Anti-DDoS Pro (Multi-IP)
	Business *string `json:"Business,omitempty" name:"Business"`

	// Anti-DDoS instance ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// Array of IPs to be bound to the Anti-DDoS instance. For Anti-DDoS Pro Single IP instance, this array can contain only one IP. If there are no IPs to bind, it can be empty; however, either `BoundDevList` or `UnBoundDevList` must not be empty;
	BoundDevList []*BoundIpInfo `json:"BoundDevList,omitempty" name:"BoundDevList" list`

	// Array of IPs to be unbound from Anti-DDoS instance. For Anti-DDoS Pro Single IP instance, this array can contain only one IP; if there are no IPs to unbind, it can be empty; however, either `BoundDevList` or `UnBoundDevList` must not be empty;
	UnBoundDevList []*BoundIpInfo `json:"UnBoundDevList,omitempty" name:"UnBoundDevList" list`

	// [Disused]
	CopyPolicy *string `json:"CopyPolicy,omitempty" name:"CopyPolicy"`
}

func (r *CreateBoundIPRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
		return errors.New("CreateBoundIPRequest has unknown keys!")
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

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBoundIPResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCCFrequencyRulesRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS service type. `bgpip`: Anti-DDoS Advanced; `net`: Anti-DDoS Ultimate
	Business *string `json:"Business,omitempty" name:"Business"`

	// Anti-DDoS instance ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// Layer-7 forwarding rule ID, which can be obtained through the `DescribleL7Rules` API
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// Matching rule. Valid values: ["include" (prefix match), "equal" (exact match)]
	Mode *string `json:"Mode,omitempty" name:"Mode"`

	// Reference period in seconds. Valid values: [10, 30, 60]
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// Number of access requests. Value range: [1-10000]
	ReqNumber *uint64 `json:"ReqNumber,omitempty" name:"ReqNumber"`

	// Action take. Valid values: ["alg" (CAPTCHA), "drop" (blocking)]
	Act *string `json:"Act,omitempty" name:"Act"`

	// Execution duration in seconds. Valid range: [1-900]
	ExeDuration *uint64 `json:"ExeDuration,omitempty" name:"ExeDuration"`

	// URI string, which must start with `/`, such as `/abc/a.php`. Length limit: 31. If URI is `/`, only prefix match can be selected as the matching mode;
	Uri *string `json:"Uri,omitempty" name:"Uri"`

	// `User-Agent` string. Length limit: 80
	UserAgent *string `json:"UserAgent,omitempty" name:"UserAgent"`

	// Cookie string. Length limit: 40
	Cookie *string `json:"Cookie,omitempty" name:"Cookie"`
}

func (r *CreateCCFrequencyRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCCFrequencyRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "RuleId")
	delete(f, "Mode")
	delete(f, "Period")
	delete(f, "ReqNumber")
	delete(f, "Act")
	delete(f, "ExeDuration")
	delete(f, "Uri")
	delete(f, "UserAgent")
	delete(f, "Cookie")
	if len(f) > 0 {
		return errors.New("CreateCCFrequencyRulesRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateCCFrequencyRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Access frequency control rule ID for CC protection
		CCFrequencyRuleId *string `json:"CCFrequencyRuleId,omitempty" name:"CCFrequencyRuleId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCCFrequencyRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCCFrequencyRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCCSelfDefinePolicyRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS service type. `bgpip`: Anti-DDoS Advanced; `bgp`: Anti-DDoS Pro (single IP); `bgp-multip`: Anti-DDoS Pro (multi-IP), `net`: Anti-DDoS Ultimate
	Business *string `json:"Business,omitempty" name:"Business"`

	// Anti-DDoS instance ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// Details of the CC protection policy
	Policy *CCPolicy `json:"Policy,omitempty" name:"Policy"`
}

func (r *CreateCCSelfDefinePolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCCSelfDefinePolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "Policy")
	if len(f) > 0 {
		return errors.New("CreateCCSelfDefinePolicyRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateCCSelfDefinePolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Success code
		Success *SuccessCode `json:"Success,omitempty" name:"Success"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCCSelfDefinePolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCCSelfDefinePolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDDoSPolicyCaseRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS service type. `bgpip`: Anti-DDoS Advanced; `bgp`: Anti-DDoS Pro (single IP); `bgp-multip`: Anti-DDoS Pro (multi-IP), `net`: Anti-DDoS Ultimate
	Business *string `json:"Business,omitempty" name:"Business"`

	// Policy scenario name string. Length limit: 64
	CaseName *string `json:"CaseName,omitempty" name:"CaseName"`

	// Development platform. Valid values: [PC (PC client), MOBILE (mobile device), TV (TV), SERVER (server)]
	PlatformTypes []*string `json:"PlatformTypes,omitempty" name:"PlatformTypes" list`

	// Category. Valid values: [WEB (website), GAME (game), APP (application), OTHER (other)]
	AppType *string `json:"AppType,omitempty" name:"AppType"`

	// Application protocol. Valid values: [tcp (TCP protocol), udp (UDP protocol), icmp (ICMP protocol), all (other protocols)]
	AppProtocols []*string `json:"AppProtocols,omitempty" name:"AppProtocols" list`

	// TCP start port. Value range: (0, 65535]
	TcpSportStart *string `json:"TcpSportStart,omitempty" name:"TcpSportStart"`

	// TCP end port. Value range: (0, 65535). It must be greater than or equal to the TCP start port.
	TcpSportEnd *string `json:"TcpSportEnd,omitempty" name:"TcpSportEnd"`

	// UDP start port. Value range: (0, 65535]
	UdpSportStart *string `json:"UdpSportStart,omitempty" name:"UdpSportStart"`

	// UDP end port. Value range: (0, 65535). It must be greater than or equal to the UDP start port.
	UdpSportEnd *string `json:"UdpSportEnd,omitempty" name:"UdpSportEnd"`

	// Whether there are customers outside China. Valid values: [no, yes]
	HasAbroad *string `json:"HasAbroad,omitempty" name:"HasAbroad"`

	// Whether to actively initiate outbound TCP requests. Valid values: [no, yes]
	HasInitiateTcp *string `json:"HasInitiateTcp,omitempty" name:"HasInitiateTcp"`

	// Whether to actively initiate outbound UDP requests. Valid values: [no, yes]
	HasInitiateUdp *string `json:"HasInitiateUdp,omitempty" name:"HasInitiateUdp"`

	// Port that actively initiates TCP requests. Value range: (0, 65535]
	PeerTcpPort *string `json:"PeerTcpPort,omitempty" name:"PeerTcpPort"`

	// Port that actively initiates UDP requests. Value range: (0, 65535]
	PeerUdpPort *string `json:"PeerUdpPort,omitempty" name:"PeerUdpPort"`

	// Fixed feature code of TCP payload. Max string length: 512
	TcpFootprint *string `json:"TcpFootprint,omitempty" name:"TcpFootprint"`

	// Fixed feature code of UDP payload. Max string length: 512
	UdpFootprint *string `json:"UdpFootprint,omitempty" name:"UdpFootprint"`

	// Web application API URL
	WebApiUrl []*string `json:"WebApiUrl,omitempty" name:"WebApiUrl" list`

	// Minimum length of TCP packet. Value range: (0, 1500)
	MinTcpPackageLen *string `json:"MinTcpPackageLen,omitempty" name:"MinTcpPackageLen"`

	// Maximum length of TCP packet. Value range: (0, 1500). It must be greater than or equal to the minimum length of TCP packet
	MaxTcpPackageLen *string `json:"MaxTcpPackageLen,omitempty" name:"MaxTcpPackageLen"`

	// Minimum length of UDP packet. Value range: (0, 1500)
	MinUdpPackageLen *string `json:"MinUdpPackageLen,omitempty" name:"MinUdpPackageLen"`

	// Maximum length of UDP packet. Value range: (0, 1500). It must be greater than or equal to the minimum length of UDP packet
	MaxUdpPackageLen *string `json:"MaxUdpPackageLen,omitempty" name:"MaxUdpPackageLen"`

	// Whether there are applications using VPN. Valid values: [no, yes]
	HasVPN *string `json:"HasVPN,omitempty" name:"HasVPN"`

	// TCP port list. You can enter a single port, or a port range. Format: 80,443,700-800,53,1000-3000.
	TcpPortList *string `json:"TcpPortList,omitempty" name:"TcpPortList"`

	// UDP port list. You can enter a single port, or a port range. Format: 80,443,700-800,53,1000-3000.
	UdpPortList *string `json:"UdpPortList,omitempty" name:"UdpPortList"`
}

func (r *CreateDDoSPolicyCaseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDDoSPolicyCaseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "CaseName")
	delete(f, "PlatformTypes")
	delete(f, "AppType")
	delete(f, "AppProtocols")
	delete(f, "TcpSportStart")
	delete(f, "TcpSportEnd")
	delete(f, "UdpSportStart")
	delete(f, "UdpSportEnd")
	delete(f, "HasAbroad")
	delete(f, "HasInitiateTcp")
	delete(f, "HasInitiateUdp")
	delete(f, "PeerTcpPort")
	delete(f, "PeerUdpPort")
	delete(f, "TcpFootprint")
	delete(f, "UdpFootprint")
	delete(f, "WebApiUrl")
	delete(f, "MinTcpPackageLen")
	delete(f, "MaxTcpPackageLen")
	delete(f, "MinUdpPackageLen")
	delete(f, "MaxUdpPackageLen")
	delete(f, "HasVPN")
	delete(f, "TcpPortList")
	delete(f, "UdpPortList")
	if len(f) > 0 {
		return errors.New("CreateDDoSPolicyCaseRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateDDoSPolicyCaseResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Policy scenario ID
		SceneId *string `json:"SceneId,omitempty" name:"SceneId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDDoSPolicyCaseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDDoSPolicyCaseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDDoSPolicyRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS service type. `bgpip`: Anti-DDoS Advanced; `bgp`: Anti-DDoS Pro (single IP); `bgp-multip`: Anti-DDoS Pro (multi-IP), `net`: Anti-DDoS Ultimate
	Business *string `json:"Business,omitempty" name:"Business"`

	// Protocol disablement, which must be entered, and the array length must be 1
	DropOptions []*DDoSPolicyDropOption `json:"DropOptions,omitempty" name:"DropOptions" list`

	// Policy name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Ports to be closed. If no ports are to be closed, enter an empty array
	PortLimits []*DDoSPolicyPortLimit `json:"PortLimits,omitempty" name:"PortLimits" list`

	// Request source IP blocklist/allowlist, which should be an empty array if there are no blocked or allowed IPs.
	IpAllowDenys []*IpBlackWhite `json:"IpAllowDenys,omitempty" name:"IpAllowDenys" list`

	// Packet filter. Enter an empty array if there are no packets to filter
	PacketFilters []*DDoSPolicyPacketFilter `json:"PacketFilters,omitempty" name:"PacketFilters" list`

	// Watermarking policy parameters. Enter an empty array if the watermarking feature is not enabled. Only one watermarking policy can be passed in (that is, the size of the array cannot exceed 1)
	WaterPrint []*WaterPrintPolicy `json:"WaterPrint,omitempty" name:"WaterPrint" list`
}

func (r *CreateDDoSPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDDoSPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "DropOptions")
	delete(f, "Name")
	delete(f, "PortLimits")
	delete(f, "IpAllowDenys")
	delete(f, "PacketFilters")
	delete(f, "WaterPrint")
	if len(f) > 0 {
		return errors.New("CreateDDoSPolicyRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateDDoSPolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Policy ID
		PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDDoSPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDDoSPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateInstanceNameRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS service type. `bgpip`: Anti-DDoS Advanced; `bgp`: Anti-DDoS Pro (single IP); `bgp-multip`: Anti-DDoS Pro (multi-IP), `net`: Anti-DDoS Ultimate
	Business *string `json:"Business,omitempty" name:"Business"`

	// Anti-DDoS instance ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// Instance name. Length limit: 32 characters
	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *CreateInstanceNameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInstanceNameRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "Name")
	if len(f) > 0 {
		return errors.New("CreateInstanceNameRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateInstanceNameResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Success code
		Success *SuccessCode `json:"Success,omitempty" name:"Success"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateInstanceNameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInstanceNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateL4HealthConfigRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS service type. `bgpip`: Anti-DDoS Advanced; `net`: Anti-DDoS Ultimate
	Business *string `json:"Business,omitempty" name:"Business"`

	// Anti-DDoS instance ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// Layer-4 health check configuration array
	HealthConfig []*L4HealthConfig `json:"HealthConfig,omitempty" name:"HealthConfig" list`
}

func (r *CreateL4HealthConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateL4HealthConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "HealthConfig")
	if len(f) > 0 {
		return errors.New("CreateL4HealthConfigRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateL4HealthConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Success code
		Success *SuccessCode `json:"Success,omitempty" name:"Success"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateL4HealthConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateL4HealthConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateL4RulesRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS service type. `bgpip`: Anti-DDoS Advanced; `bgp`: Anti-DDoS Pro (single IP); `bgp-multip`: Anti-DDoS Pro (multi-IP), `net`: Anti-DDoS Ultimate
	Business *string `json:"Business,omitempty" name:"Business"`

	// Anti-DDoS instance ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// Rule list
	Rules []*L4RuleEntry `json:"Rules,omitempty" name:"Rules" list`
}

func (r *CreateL4RulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateL4RulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "Rules")
	if len(f) > 0 {
		return errors.New("CreateL4RulesRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateL4RulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Success code
		Success *SuccessCode `json:"Success,omitempty" name:"Success"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateL4RulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateL4RulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateL7CCRuleRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS service type. `bgpip`: Anti-DDoS Advanced; `net`: Anti-DDoS Ultimate
	Business *string `json:"Business,omitempty" name:"Business"`

	// Anti-DDoS instance ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// Operation. Valid values: [query (query), add (add), del (delete)]
	Method *string `json:"Method,omitempty" name:"Method"`

	// Layer-7 forwarding rule ID, such as rule-0000001
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// Custom layer-7 CC protection rule. If the `Method` is `query`, this field can be left empty; if the `Method` is `add` or `del`, it is required, and the array length can only be 1;
	RuleConfig []*CCRuleConfig `json:"RuleConfig,omitempty" name:"RuleConfig" list`
}

func (r *CreateL7CCRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateL7CCRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "Method")
	delete(f, "RuleId")
	delete(f, "RuleConfig")
	if len(f) > 0 {
		return errors.New("CreateL7CCRuleRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateL7CCRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Custom layer-7 CC protection rule parameters. If custom CC protection rule is not enabled, an empty array will be returned.
		RuleConfig []*CCRuleConfig `json:"RuleConfig,omitempty" name:"RuleConfig" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateL7CCRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateL7CCRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateL7HealthConfigRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS service type. `bgpip`: Anti-DDoS Advanced; `net`: Anti-DDoS Ultimate
	Business *string `json:"Business,omitempty" name:"Business"`

	// Anti-DDoS instance ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// Layer-7 health check configuration array
	HealthConfig []*L7HealthConfig `json:"HealthConfig,omitempty" name:"HealthConfig" list`
}

func (r *CreateL7HealthConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateL7HealthConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "HealthConfig")
	if len(f) > 0 {
		return errors.New("CreateL7HealthConfigRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateL7HealthConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Success code
		Success *SuccessCode `json:"Success,omitempty" name:"Success"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateL7HealthConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateL7HealthConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateL7RuleCertRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS service type. `bgpip`: Anti-DDoS Advanced; `net`: Anti-DDoS Ultimate
	Business *string `json:"Business,omitempty" name:"Business"`

	// The resource instance ID, such as the ID of an Anti-DDoS Advanced instance or the ID of an Anti-DDoS Ultimate instance.
	Id *string `json:"Id,omitempty" name:"Id"`

	// Rule ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// Certificate type, which is required if the protocol is HTTPS. Valid value: [2 (Tencent Cloud-hosted certificate)]
	CertType *uint64 `json:"CertType,omitempty" name:"CertType"`

	// If the certificate is a Tencent Cloud-hosted certificate, this field must be entered with the hosted certificate ID.
	SSLId *string `json:"SSLId,omitempty" name:"SSLId"`

	// [Disused] If the certificate is an external certificate, this field must be entered with the certificate content. 
	Cert *string `json:"Cert,omitempty" name:"Cert"`

	// [Disused] If the certificate is an external certificate, this field must be entered with the certificate key. 
	PrivateKey *string `json:"PrivateKey,omitempty" name:"PrivateKey"`
}

func (r *CreateL7RuleCertRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateL7RuleCertRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "RuleId")
	delete(f, "CertType")
	delete(f, "SSLId")
	delete(f, "Cert")
	delete(f, "PrivateKey")
	if len(f) > 0 {
		return errors.New("CreateL7RuleCertRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateL7RuleCertResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Success code
		Success *SuccessCode `json:"Success,omitempty" name:"Success"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateL7RuleCertResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateL7RuleCertResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateL7RulesRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS service type. `bgpip`: Anti-DDoS Advanced; `bgp`: Anti-DDoS Pro (single IP); `bgp-multip`: Anti-DDoS Pro (multi-IP), `net`: Anti-DDoS Ultimate
	Business *string `json:"Business,omitempty" name:"Business"`

	// Anti-DDoS instance ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// Rule list
	Rules []*L7RuleEntry `json:"Rules,omitempty" name:"Rules" list`
}

func (r *CreateL7RulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateL7RulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "Rules")
	if len(f) > 0 {
		return errors.New("CreateL7RulesRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateL7RulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Success code
		Success *SuccessCode `json:"Success,omitempty" name:"Success"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateL7RulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateL7RulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateL7RulesUploadRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS service type. `bgpip`: Anti-DDoS Advanced; `net`: Anti-DDoS Ultimate
	Business *string `json:"Business,omitempty" name:"Business"`

	// Anti-DDoS instance ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// Rule list
	Rules []*L7RuleEntry `json:"Rules,omitempty" name:"Rules" list`
}

func (r *CreateL7RulesUploadRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateL7RulesUploadRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "Rules")
	if len(f) > 0 {
		return errors.New("CreateL7RulesUploadRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateL7RulesUploadResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Success code
		Success *SuccessCode `json:"Success,omitempty" name:"Success"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateL7RulesUploadResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateL7RulesUploadResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateNetReturnRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS service type (`net`: Anti-DDoS Ultimate)
	Business *string `json:"Business,omitempty" name:"Business"`

	// Anti-DDoS instance ID
	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *CreateNetReturnRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNetReturnRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	if len(f) > 0 {
		return errors.New("CreateNetReturnRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateNetReturnResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateNetReturnResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNetReturnResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateNewL7RulesUploadRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS service type (`bgpip`: Anti-DDoS Advanced).
	Business *string `json:"Business,omitempty" name:"Business"`

	// Resource ID list.
	IdList []*string `json:"IdList,omitempty" name:"IdList" list`

	// Resource IP address list.
	VipList []*string `json:"VipList,omitempty" name:"VipList" list`

	// Rule list.
	Rules []*L7RuleEntry `json:"Rules,omitempty" name:"Rules" list`
}

func (r *CreateNewL7RulesUploadRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNewL7RulesUploadRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "IdList")
	delete(f, "VipList")
	delete(f, "Rules")
	if len(f) > 0 {
		return errors.New("CreateNewL7RulesUploadRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateNewL7RulesUploadResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Success code.
		Success *SuccessCode `json:"Success,omitempty" name:"Success"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateNewL7RulesUploadResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNewL7RulesUploadResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateUnblockIpRequest struct {
	*tchttp.BaseRequest

	// IP
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// Type of the unblocking action (`user`: self-service unblocking, `auto`: automatic unblocking, `update`: unblocking by service upgrading, `bind`: unblocking by binding Anti-DDoS Pro instance)
	ActionType *string `json:"ActionType,omitempty" name:"ActionType"`
}

func (r *CreateUnblockIpRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUnblockIpRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ip")
	delete(f, "ActionType")
	if len(f) > 0 {
		return errors.New("CreateUnblockIpRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateUnblockIpResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// IP
		Ip *string `json:"Ip,omitempty" name:"Ip"`

		// Type of the unblocking action (`user`: self-service unblocking, `auto`: automatic unblocking, `update`: unblocking by service upgrading, `bind`: unblocking by binding Anti-DDoS Pro instance)
		ActionType *string `json:"ActionType,omitempty" name:"ActionType"`

		// Unblocked time (estimated)
		UnblockTime *string `json:"UnblockTime,omitempty" name:"UnblockTime"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateUnblockIpResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUnblockIpResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DDoSAlarmThreshold struct {

	// Alarm threshold type. 1: inbound traffic, 2: cleansed traffic
	AlarmType *uint64 `json:"AlarmType,omitempty" name:"AlarmType"`

	// Alarm threshold, which should be greater than 0 (currently scheduled value)
	AlarmThreshold *uint64 `json:"AlarmThreshold,omitempty" name:"AlarmThreshold"`
}

type DDoSAttackSourceRecord struct {

	// Attack source IP
	SrcIp *string `json:"SrcIp,omitempty" name:"SrcIp"`

	// Province (valid for Mainland China)
	Province *string `json:"Province,omitempty" name:"Province"`

	// Country/region
	Nation *string `json:"Nation,omitempty" name:"Nation"`

	// Total number of attack packets
	PacketSum *uint64 `json:"PacketSum,omitempty" name:"PacketSum"`

	// Total attack traffic
	PacketLen *uint64 `json:"PacketLen,omitempty" name:"PacketLen"`
}

type DDoSEventRecord struct {

	// Anti-DDoS service type. `bgpip`: Anti-DDoS Advanced; `bgp`: Anti-DDoS Pro (Single IP); `bgp-multip`: Anti-DDoS Pro (Multi-IP); `net`: Anti-DDoS Ultimate; `basic`: Anti-DDoS Basic
	Business *string `json:"Business,omitempty" name:"Business"`

	// Anti-DDoS instance ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// Resource IP
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// Attack start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Attack end time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Maximum attack bandwidth
	Mbps *uint64 `json:"Mbps,omitempty" name:"Mbps"`

	// Maximum attack packet rate
	Pps *uint64 `json:"Pps,omitempty" name:"Pps"`

	// Attack type
	AttackType *string `json:"AttackType,omitempty" name:"AttackType"`

	// Whether the IP is blocked. Valid values: [1 (yes), 0 (no), 2 (invalid value)]
	BlockFlag *uint64 `json:"BlockFlag,omitempty" name:"BlockFlag"`

	// Whether the elastic protection bandwidth is exceeded. Valid values: [yes (yes), no (no), empty string (unknown value)]
	OverLoad *string `json:"OverLoad,omitempty" name:"OverLoad"`

	// Attack status. Valid values: [0 (ongoing), 1 (ended)]
	AttackStatus *uint64 `json:"AttackStatus,omitempty" name:"AttackStatus"`

	// Resource name
	// Note: this field may return null, indicating that no valid values can be obtained.
	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`

	// Attack event ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	EventId *string `json:"EventId,omitempty" name:"EventId"`
}

type DDoSPolicyDropOption struct {

	// Blocks all TCP traffic. Valid values: [0,1]
	DropTcp *uint64 `json:"DropTcp,omitempty" name:"DropTcp"`

	// Blocks all UDP traffic. Valid values: [0,1]
	DropUdp *uint64 `json:"DropUdp,omitempty" name:"DropUdp"`

	// Blocks all ICMP traffic. Valid values: [0,1]
	DropIcmp *uint64 `json:"DropIcmp,omitempty" name:"DropIcmp"`

	// Blocks traffic of other protocols. Valid values: [0,1]
	DropOther *uint64 `json:"DropOther,omitempty" name:"DropOther"`

	// Rejects traffic from outside Mainland China. Valid values: [0,1]
	DropAbroad *uint64 `json:"DropAbroad,omitempty" name:"DropAbroad"`

	// Null session protection. Valid values: [0,1]
	CheckSyncConn *uint64 `json:"CheckSyncConn,omitempty" name:"CheckSyncConn"`

	// New connection suppression based on source IP and destination IP. Value range: [0,4294967295]
	SdNewLimit *uint64 `json:"SdNewLimit,omitempty" name:"SdNewLimit"`

	// New connection suppression based on destination IP. Value range: [0,4294967295]
	DstNewLimit *uint64 `json:"DstNewLimit,omitempty" name:"DstNewLimit"`

	// Concurrent connection suppression based on source IP and destination IP. Value range: [0,4294967295]
	SdConnLimit *uint64 `json:"SdConnLimit,omitempty" name:"SdConnLimit"`

	// Concurrent connection suppression based on destination IP. Value range: [0,4294967295]
	DstConnLimit *uint64 `json:"DstConnLimit,omitempty" name:"DstConnLimit"`

	// Threshold for triggering connection suppression. Value range: [0,4294967295]
	BadConnThreshold *uint64 `json:"BadConnThreshold,omitempty" name:"BadConnThreshold"`

	// Exceptional connection detection condition: null session protection status. Valid values: [0,1]
	NullConnEnable *uint64 `json:"NullConnEnable,omitempty" name:"NullConnEnable"`

	// Exceptional connection detection condition: connection timeout. Valid values: [0,65535]
	ConnTimeout *uint64 `json:"ConnTimeout,omitempty" name:"ConnTimeout"`

	// Exceptional connection detection condition: percentage of SYN out of ACK. Valid values: [0,100]
	SynRate *uint64 `json:"SynRate,omitempty" name:"SynRate"`

	// Exceptional connection detection condition: SYN threshold. Valid values: [0,100]
	SynLimit *uint64 `json:"SynLimit,omitempty" name:"SynLimit"`

	// TCP speed limit. Value range: [0,4294967295]
	DTcpMbpsLimit *uint64 `json:"DTcpMbpsLimit,omitempty" name:"DTcpMbpsLimit"`

	// UDP speed limit. Value range: [0,4294967295]
	DUdpMbpsLimit *uint64 `json:"DUdpMbpsLimit,omitempty" name:"DUdpMbpsLimit"`

	// ICMP speed limit. Value range: [0,4294967295]
	DIcmpMbpsLimit *uint64 `json:"DIcmpMbpsLimit,omitempty" name:"DIcmpMbpsLimit"`

	// Other protocol speed limit. Value range: [0,4294967295]
	DOtherMbpsLimit *uint64 `json:"DOtherMbpsLimit,omitempty" name:"DOtherMbpsLimit"`
}

type DDoSPolicyPacketFilter struct {

	// Protocol. Valid values: [tcp, udp, icmp, all]
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Start source port. Value range: [0,65535]
	SportStart *uint64 `json:"SportStart,omitempty" name:"SportStart"`

	// End source port. Value range: [0,65535]
	SportEnd *uint64 `json:"SportEnd,omitempty" name:"SportEnd"`

	// Start destination port. Value range: [0,65535]
	DportStart *uint64 `json:"DportStart,omitempty" name:"DportStart"`

	// End destination port. Value range: [0,65535]
	DportEnd *uint64 `json:"DportEnd,omitempty" name:"DportEnd"`

	// Minimum packet length. Value range: [0,1500]
	PktlenMin *uint64 `json:"PktlenMin,omitempty" name:"PktlenMin"`

	// Maximum packet length. Value range: [0,1500]
	PktlenMax *uint64 `json:"PktlenMax,omitempty" name:"PktlenMax"`

	// Whether to detect the payload. Valid values: [
	// begin_l3 (IP header)
	// begin_l4 (TCP header)
	// begin_l5 (payload)
	// no_match (no check)
	// ]
	MatchBegin *string `json:"MatchBegin,omitempty" name:"MatchBegin"`

	// Whether it is a regular expression. Valid values: [sunday (keyword), pcre (regular expression)]
	MatchType *string `json:"MatchType,omitempty" name:"MatchType"`

	// Keyword or regular expression
	Str *string `json:"Str,omitempty" name:"Str"`

	// Detection depth. Value range: [0,1500]
	Depth *uint64 `json:"Depth,omitempty" name:"Depth"`

	// Detection offset. Value range: [0,1500]
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Whether to include. Valid values: [0 (no), 1 (yes)]
	IsNot *uint64 `json:"IsNot,omitempty" name:"IsNot"`

	// Policy action. Valid values: [drop, drop_black, drop_rst, drop_black_rst, transmit]
	Action *string `json:"Action,omitempty" name:"Action"`
}

type DDoSPolicyPortLimit struct {

	// Protocol. Valid values: [tcp, udp, all]
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Start destination port. Value range: [0,65535]
	DPortStart *uint64 `json:"DPortStart,omitempty" name:"DPortStart"`

	// End destination port, which must be greater than or equal to the start destination port. Value range: [0,65535]
	DPortEnd *uint64 `json:"DPortEnd,omitempty" name:"DPortEnd"`

	// Start source port. Value range: [0,65535]
	// Note: this field may return null, indicating that no valid values can be obtained.
	SPortStart *uint64 `json:"SPortStart,omitempty" name:"SPortStart"`

	// End source port, which must be greater than or equal to the start source port. Value range: [0,65535]
	// Note: this field may return null, indicating that no valid values can be obtained.
	SPortEnd *uint64 `json:"SPortEnd,omitempty" name:"SPortEnd"`

	// Action to be executed. Valid values: [drop (discard) , transmit (forward)]
	// Note: this field may return null, indicating that no valid values can be obtained.
	Action *string `json:"Action,omitempty" name:"Action"`

	// Type of port to be disabled. Valid values: [0 (destination port range), 1 (source port range), 2 (destination port range and source port range)]
	// Note: this field may return null, indicating that no valid values can be obtained.
	Kind *uint64 `json:"Kind,omitempty" name:"Kind"`
}

type DDosPolicy struct {

	// Resource bound to policy
	Resources []*ResourceIp `json:"Resources,omitempty" name:"Resources" list`

	// Disabled protocol
	DropOptions *DDoSPolicyDropOption `json:"DropOptions,omitempty" name:"DropOptions"`

	// Disabled port
	PortLimits []*DDoSPolicyPortLimit `json:"PortLimits,omitempty" name:"PortLimits" list`

	// Packet filter
	PacketFilters []*DDoSPolicyPacketFilter `json:"PacketFilters,omitempty" name:"PacketFilters" list`

	// IP blocklist/allowlist
	IpBlackWhiteLists []*IpBlackWhite `json:"IpBlackWhiteLists,omitempty" name:"IpBlackWhiteLists" list`

	// Policy ID
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

	// Policy name
	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`

	// Policy creation time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Watermarking policy parameter. There can be only one policy. If there are no policies, the array is empty
	WaterPrint []*WaterPrintPolicy `json:"WaterPrint,omitempty" name:"WaterPrint" list`

	// Watermark key. There can be up to two keys. If there are no policies, the array is empty
	WaterKey []*WaterPrintKey `json:"WaterKey,omitempty" name:"WaterKey" list`

	// Resource instance bound to policy
	// Note: this field may return null, indicating that no valid values can be obtained.
	BoundResources []*string `json:"BoundResources,omitempty" name:"BoundResources" list`

	// Policy scenario
	// Note: this field may return null, indicating that no valid values can be obtained.
	SceneId *string `json:"SceneId,omitempty" name:"SceneId"`
}

type DeleteCCFrequencyRulesRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS service type. `bgpip`: Anti-DDoS Advanced; `net`: Anti-DDoS Ultimate
	Business *string `json:"Business,omitempty" name:"Business"`

	// Access frequency control rule ID for CC protection
	CCFrequencyRuleId *string `json:"CCFrequencyRuleId,omitempty" name:"CCFrequencyRuleId"`
}

func (r *DeleteCCFrequencyRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCCFrequencyRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "CCFrequencyRuleId")
	if len(f) > 0 {
		return errors.New("DeleteCCFrequencyRulesRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCCFrequencyRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Success code
		Success *SuccessCode `json:"Success,omitempty" name:"Success"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteCCFrequencyRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCCFrequencyRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCCSelfDefinePolicyRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS service type. `bgpip`: Anti-DDoS Advanced; `bgp`: Anti-DDoS Pro (single IP); `bgp-multip`: Anti-DDoS Pro (multi-IP), `net`: Anti-DDoS Ultimate
	Business *string `json:"Business,omitempty" name:"Business"`

	// Anti-DDoS instance ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// Policy ID
	SetId *string `json:"SetId,omitempty" name:"SetId"`
}

func (r *DeleteCCSelfDefinePolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCCSelfDefinePolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "SetId")
	if len(f) > 0 {
		return errors.New("DeleteCCSelfDefinePolicyRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCCSelfDefinePolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Success code
		Success *SuccessCode `json:"Success,omitempty" name:"Success"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteCCSelfDefinePolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCCSelfDefinePolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDDoSPolicyCaseRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS service type. `bgpip`: Anti-DDoS Advanced; `bgp`: Anti-DDoS Pro (single IP); `bgp-multip`: Anti-DDoS Pro (multi-IP), `net`: Anti-DDoS Ultimate
	Business *string `json:"Business,omitempty" name:"Business"`

	// Policy scenario ID
	SceneId *string `json:"SceneId,omitempty" name:"SceneId"`
}

func (r *DeleteDDoSPolicyCaseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDDoSPolicyCaseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "SceneId")
	if len(f) > 0 {
		return errors.New("DeleteDDoSPolicyCaseRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDDoSPolicyCaseResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Success code
		Success *SuccessCode `json:"Success,omitempty" name:"Success"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteDDoSPolicyCaseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDDoSPolicyCaseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDDoSPolicyRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS service type. `bgpip`: Anti-DDoS Advanced; `bgp`: Anti-DDoS Pro (single IP); `bgp-multip`: Anti-DDoS Pro (multi-IP), `net`: Anti-DDoS Ultimate
	Business *string `json:"Business,omitempty" name:"Business"`

	// Policy ID
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`
}

func (r *DeleteDDoSPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDDoSPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "PolicyId")
	if len(f) > 0 {
		return errors.New("DeleteDDoSPolicyRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDDoSPolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Success code
		Success *SuccessCode `json:"Success,omitempty" name:"Success"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteDDoSPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDDoSPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteL4RulesRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS service type. `bgpip`: Anti-DDoS Advanced; `bgp`: Anti-DDoS Pro (single IP); `bgp-multip`: Anti-DDoS Pro (multi-IP), `net`: Anti-DDoS Ultimate
	Business *string `json:"Business,omitempty" name:"Business"`

	// Anti-DDoS instance ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// Rule ID list
	RuleIdList []*string `json:"RuleIdList,omitempty" name:"RuleIdList" list`
}

func (r *DeleteL4RulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteL4RulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "RuleIdList")
	if len(f) > 0 {
		return errors.New("DeleteL4RulesRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteL4RulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Success code
		Success *SuccessCode `json:"Success,omitempty" name:"Success"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteL4RulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteL4RulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteL7RulesRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS service type. `bgpip`: Anti-DDoS Advanced; `bgp`: Anti-DDoS Pro (single IP); `bgp-multip`: Anti-DDoS Pro (multi-IP), `net`: Anti-DDoS Ultimate
	Business *string `json:"Business,omitempty" name:"Business"`

	// Anti-DDoS instance ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// Rule ID list
	RuleIdList []*string `json:"RuleIdList,omitempty" name:"RuleIdList" list`
}

func (r *DeleteL7RulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteL7RulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "RuleIdList")
	if len(f) > 0 {
		return errors.New("DeleteL7RulesRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteL7RulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Success code
		Success *SuccessCode `json:"Success,omitempty" name:"Success"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteL7RulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteL7RulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeActionLogRequest struct {
	*tchttp.BaseRequest

	// Start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Anti-DDoS service type. `bgpip`: Anti-DDoS Advanced; `bgp`: Anti-DDoS Pro (single IP); `bgp-multip`: Anti-DDoS Pro (multi-IP), `net`: Anti-DDoS Ultimate
	Business *string `json:"Business,omitempty" name:"Business"`

	// Search value, which can only be resource ID or user `UIN`
	Filter *string `json:"Filter,omitempty" name:"Filter"`

	// Number of entries per page. A value of 0 means no pagination
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Page start offset, whose value is (page number - 1) * number of entries per page
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeActionLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeActionLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Business")
	delete(f, "Filter")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return errors.New("DescribeActionLogRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeActionLogResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of records
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Record array
		Data []*KeyValueRecord `json:"Data,omitempty" name:"Data" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeActionLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeActionLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBGPIPL7RuleMaxCntRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS service type (`bgpip`: Anti-DDoS Advanced)
	Business *string `json:"Business,omitempty" name:"Business"`

	// Anti-DDoS instance ID
	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeBGPIPL7RuleMaxCntRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBGPIPL7RuleMaxCntRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	if len(f) > 0 {
		return errors.New("DescribeBGPIPL7RuleMaxCntRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBGPIPL7RuleMaxCntResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Maximum number of layer-7 rules that can be added for Anti-DDoS Advanced
		Count *uint64 `json:"Count,omitempty" name:"Count"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBGPIPL7RuleMaxCntResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBGPIPL7RuleMaxCntResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBaradDataRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS service type. `bgpip`: Anti-DDoS Advanced; `net`: Anti-DDoS Ultimate
	Business *string `json:"Business,omitempty" name:"Business"`

	// Anti-DDoS instance ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// Metric name. Valid values:
	// connum: number of active TCP connections;
	// new_conn: number of new TCP connections;
	// inactive_conn: number of inactive connections;
	// intraffic: inbound traffic;
	// outtraffic: outbound traffic;
	// alltraffic: sum of inbound and outbound traffic;
	// inpkg: inbound packet rate;
	// outpkg: outbound packet rate;
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// Statistical time granularity in seconds (300: 5-minute, 3600: hourly, 86400: daily)
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// Statistics start time. The second part is kept at 0, and the minute part is a multiple of 5
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Statistics end time. The second part is kept at 0, and the minute part is a multiple of 5
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Statistical method. Valid values:
	// max: maximum value;
	// min: minimum value;
	// avg: average;
	Statistics *string `json:"Statistics,omitempty" name:"Statistics"`

	// Protocol port array
	ProtocolPort []*ProtocolPort `json:"ProtocolPort,omitempty" name:"ProtocolPort" list`

	// Resource instance IP, which is required only if `Business` is `net` (Anti-DDoS Ultimate), because an Anti-DDoS Ultimate instance has multiple IPs;
	Ip *string `json:"Ip,omitempty" name:"Ip"`
}

func (r *DescribeBaradDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBaradDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "MetricName")
	delete(f, "Period")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Statistics")
	delete(f, "ProtocolPort")
	delete(f, "Ip")
	if len(f) > 0 {
		return errors.New("DescribeBaradDataRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBaradDataResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Returned metric value
		DataList []*BaradData `json:"DataList,omitempty" name:"DataList" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBaradDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBaradDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBasicCCThresholdRequest struct {
	*tchttp.BaseRequest

	// Queried IP address, such as 1.1.1.1
	BasicIp *string `json:"BasicIp,omitempty" name:"BasicIp"`

	// IP region. Valid values: region abbreviations such as gz, bj, sh, and hk
	BasicRegion *string `json:"BasicRegion,omitempty" name:"BasicRegion"`

	// Zone type. Valid values: public (public cloud zone), bm (BM zone), nat (NAT server zone), channel (internet channel).
	BasicBizType *string `json:"BasicBizType,omitempty" name:"BasicBizType"`

	// Device type. Valid values: cvm (CVM), clb (public CLB), lb (BM CLB), nat (NAT server), channel (internet channel).
	BasicDeviceType *string `json:"BasicDeviceType,omitempty" name:"BasicDeviceType"`

	// IPInstance Nat gateway, which is optional. (If the device type to be queried is a NAT server, this parameter is required, which can be obtained through the NAT resource query API)
	BasicIpInstance *string `json:"BasicIpInstance,omitempty" name:"BasicIpInstance"`

	// ISP line, which is optional. (If the device type to be queried is a NAT server, this parameter should be 5)
	BasicIspCode *uint64 `json:"BasicIspCode,omitempty" name:"BasicIspCode"`
}

func (r *DescribeBasicCCThresholdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBasicCCThresholdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BasicIp")
	delete(f, "BasicRegion")
	delete(f, "BasicBizType")
	delete(f, "BasicDeviceType")
	delete(f, "BasicIpInstance")
	delete(f, "BasicIspCode")
	if len(f) > 0 {
		return errors.New("DescribeBasicCCThresholdRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBasicCCThresholdResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// CC status (0: disabled, 1: enabled)
		CCEnable *uint64 `json:"CCEnable,omitempty" name:"CCEnable"`

		// CC protection threshold
		CCThreshold *uint64 `json:"CCThreshold,omitempty" name:"CCThreshold"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBasicCCThresholdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBasicCCThresholdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBasicDeviceThresholdRequest struct {
	*tchttp.BaseRequest

	// Queried IP address, such as 1.1.1.1
	BasicIp *string `json:"BasicIp,omitempty" name:"BasicIp"`

	// IP region. Valid values: region abbreviations such as gz, bj, sh, and hk
	BasicRegion *string `json:"BasicRegion,omitempty" name:"BasicRegion"`

	// Zone type. Valid values: public (public cloud zone), bm (BM zone), nat (NAT server zone), channel (internet channel).
	BasicBizType *string `json:"BasicBizType,omitempty" name:"BasicBizType"`

	// Device type. Valid values: cvm (CVM), clb (public CLB), lb (BM CLB), nat (NAT server), channel (internet channel).
	BasicDeviceType *string `json:"BasicDeviceType,omitempty" name:"BasicDeviceType"`

	// Validity check. Valid value: 1
	BasicCheckFlag *uint64 `json:"BasicCheckFlag,omitempty" name:"BasicCheckFlag"`

	// IPInstance Nat gateway, which is optional. (If the device type to be queried is a NAT server, this parameter is required, which can be obtained through the NAT resource query API)
	BasicIpInstance *string `json:"BasicIpInstance,omitempty" name:"BasicIpInstance"`

	// ISP line, which is optional. (If the device type to be queried is a NAT server, this parameter should be 5)
	BasicIspCode *uint64 `json:"BasicIspCode,omitempty" name:"BasicIspCode"`
}

func (r *DescribeBasicDeviceThresholdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBasicDeviceThresholdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BasicIp")
	delete(f, "BasicRegion")
	delete(f, "BasicBizType")
	delete(f, "BasicDeviceType")
	delete(f, "BasicCheckFlag")
	delete(f, "BasicIpInstance")
	delete(f, "BasicIspCode")
	if len(f) > 0 {
		return errors.New("DescribeBasicDeviceThresholdRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBasicDeviceThresholdResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Blackhole blocking value
		Threshold *uint64 `json:"Threshold,omitempty" name:"Threshold"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBasicDeviceThresholdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBasicDeviceThresholdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBizHttpStatusRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS service type (`bgpip`: Anti-DDoS Advanced)
	Business *string `json:"Business,omitempty" name:"Business"`

	// Resource ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// Statistical period in seconds. Valid values: 300, 1800, 3600, 21600, and 86400.
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// Statistics start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Statistics end time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Statistical mode, which only supports sum.
	Statistics *string `json:"Statistics,omitempty" name:"Statistics"`

	// Protocol and port list, which is valid when the statistical dimension is the number of connections. Valid protocols: TCP, UDP, HTTP, and HTTPS.
	ProtoInfo []*ProtocolPort `json:"ProtoInfo,omitempty" name:"ProtoInfo" list`

	// Specific domain name query
	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *DescribeBizHttpStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBizHttpStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "Period")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Statistics")
	delete(f, "ProtoInfo")
	delete(f, "Domain")
	if len(f) > 0 {
		return errors.New("DescribeBizHttpStatusRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBizHttpStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Statistics on the HTTP status codes of business traffic
		HttpStatusMap *HttpStatusMap `json:"HttpStatusMap,omitempty" name:"HttpStatusMap"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBizHttpStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBizHttpStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCCAlarmThresholdRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS service type (`shield`: Chess Shield, `bgpip`: Anti-DDoS Advanced, `bgp`: Anti-DDoS Pro (single IP), `bgp-multip`: Anti-DDoS Pro (multi-IP), `net`: Anti-DDoS Ultimate)
	Business *string `json:"Business,omitempty" name:"Business"`

	// Anti-DDoS instance ID
	RsId *string `json:"RsId,omitempty" name:"RsId"`
}

func (r *DescribeCCAlarmThresholdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCCAlarmThresholdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "RsId")
	if len(f) > 0 {
		return errors.New("DescribeCCAlarmThresholdRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCCAlarmThresholdResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// CC alarm threshold
		CCAlarmThreshold *CCAlarmThreshold `json:"CCAlarmThreshold,omitempty" name:"CCAlarmThreshold"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCCAlarmThresholdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCCAlarmThresholdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCCEvListRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS service type. `bgpip`: Anti-DDoS Advanced; `bgp`: Anti-DDoS Pro (Single IP); `bgp-multip`: Anti-DDoS Pro (Multi-IP); `net`: Anti-DDoS Ultimate; `basic`: Anti-DDoS Basic
	Business *string `json:"Business,omitempty" name:"Business"`

	// Start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Anti-DDoS instance ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// Resource instance IP. When `business` is not `basic`, if `IpList` is not empty, `Id` must not be empty;
	IpList []*string `json:"IpList,omitempty" name:"IpList" list`

	// Number of entries per page. A value of 0 means no pagination
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Page start offset, whose value is (page number - 1) * number of entries per page
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeCCEvListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCCEvListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Id")
	delete(f, "IpList")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return errors.New("DescribeCCEvListRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCCEvListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Anti-DDoS service type. `shield`: Chess Shield; `bgpip`: Anti-DDoS Advanced; `bgp`: Anti-DDoS Pro (single IP); `bgp-multip`: Anti-DDoS Pro (multi-IP); `net`: Anti-DDoS Ultimate; `basic`: Anti-DDoS Basic
		Business *string `json:"Business,omitempty" name:"Business"`

		// Anti-DDoS instance ID
		Id *string `json:"Id,omitempty" name:"Id"`

		// Instance IP list
	// Note: this field may return null, indicating that no valid values can be obtained.
		IpList []*string `json:"IpList,omitempty" name:"IpList" list`

		// Start time
		StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

		// End time
		EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

		// CC attack event list
		Data []*CCEventRecord `json:"Data,omitempty" name:"Data" list`

		// Total number of records
		Total *uint64 `json:"Total,omitempty" name:"Total"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCCEvListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCCEvListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCCFrequencyRulesRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS service type. `bgpip`: Anti-DDoS Advanced; `net`: Anti-DDoS Ultimate
	Business *string `json:"Business,omitempty" name:"Business"`

	// Anti-DDoS instance ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// Layer-7 forwarding rule ID (which can be obtained from the layer-7 forwarding rule API). If a value is entered, it means to get the access frequency control rule of the forwarding rule;
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
}

func (r *DescribeCCFrequencyRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCCFrequencyRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "RuleId")
	if len(f) > 0 {
		return errors.New("DescribeCCFrequencyRulesRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCCFrequencyRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Access frequency control rule list
		CCFrequencyRuleList []*CCFrequencyRule `json:"CCFrequencyRuleList,omitempty" name:"CCFrequencyRuleList" list`

		// Access frequency control rule status. Valid values: [on, off];
		CCFrequencyRuleStatus *string `json:"CCFrequencyRuleStatus,omitempty" name:"CCFrequencyRuleStatus"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCCFrequencyRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCCFrequencyRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCCIpAllowDenyRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS service type. `bgpip`: Anti-DDoS Advanced; `bgp`: Anti-DDoS Pro (single IP); `bgp-multip`: Anti-DDoS Pro (multi-IP), `net`: Anti-DDoS Ultimate
	Business *string `json:"Business,omitempty" name:"Business"`

	// Anti-DDoS instance ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// Blocklist or allowlist. Valid values: [white (allowlist), black (blocklist)]
	// Note: this array can only have one value. It cannot get the blocklist and allowlist at the same time
	Type []*string `json:"Type,omitempty" name:"Type" list`

	// Pagination parameter
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Pagination parameter
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// HTTP or HTTPS CC protection, which is optional. Valid values: [http (HTTP CC protection), https (HTTPS CC protection)];
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
}

func (r *DescribeCCIpAllowDenyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCCIpAllowDenyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "Type")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Protocol")
	if len(f) > 0 {
		return errors.New("DescribeCCIpAllowDenyRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCCIpAllowDenyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// This field has been replaced by `RecordList` and should not be used
		Data []*KeyValue `json:"Data,omitempty" name:"Data" list`

		// Number of records
		Total *uint64 `json:"Total,omitempty" name:"Total"`

		// Returned Blocklist/allowlist record,
	// If "Key":"ip", "Value": IP;
	// If "Key":"domain", "Value": domain name.
	// If "Key":"type", "Value" can be `white` (allowlist) or `black` (blocklist).
	// If "Key":"protocol", "Value": CC protection protocol (HTTP or HTTPS);
		RecordList []*KeyValueRecord `json:"RecordList,omitempty" name:"RecordList" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCCIpAllowDenyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCCIpAllowDenyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCCSelfDefinePolicyRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS service type. `bgp`: Anti-DDoS Pro, `bgp-multip`: Anti-DDoS Pro (multi-IP)
	Business *string `json:"Business,omitempty" name:"Business"`

	// Anti-DDoS instance ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// Number of entries pulled
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Offset
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeCCSelfDefinePolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCCSelfDefinePolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return errors.New("DescribeCCSelfDefinePolicyRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCCSelfDefinePolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of custom rules
		Total *uint64 `json:"Total,omitempty" name:"Total"`

		// Policy list
		Policys []*CCPolicy `json:"Policys,omitempty" name:"Policys" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCCSelfDefinePolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCCSelfDefinePolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCCTrendRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS service type. `bgpip`: Anti-DDoS Advanced; `bgp`: Anti-DDoS Pro (Single IP); `bgp-multip`: Anti-DDoS Pro (Multi-IP); `net`: Anti-DDoS Ultimate; `basic`: Anti-DDoS Basic
	Business *string `json:"Business,omitempty" name:"Business"`

	// Resource IP
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// Metric. Valid values: [inqps (total requests peak), dropqps (attack requests peak)]
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// Statistical granularity. Valid values: [300 (5-minute), 3600 (hourly), 86400 (daily)]
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// Statistics start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Statistics end time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Resource instance ID. If `Business` is `basic`, this field is not required (because Anti-DDoS Basic has no resource instance)
	Id *string `json:"Id,omitempty" name:"Id"`

	// (Optional) Domain name
	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *DescribeCCTrendRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCCTrendRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Ip")
	delete(f, "MetricName")
	delete(f, "Period")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Id")
	delete(f, "Domain")
	if len(f) > 0 {
		return errors.New("DescribeCCTrendRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCCTrendResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Anti-DDoS service type. `bgpip`: Anti-DDoS Advanced; `bgp`: Anti-DDoS Pro (Single IP); `bgp-multip`: Anti-DDoS Pro (Multi-IP); `net`: Anti-DDoS Ultimate; `basic`: Anti-DDoS Basic
		Business *string `json:"Business,omitempty" name:"Business"`

		// Anti-DDoS instance ID
	// Note: this field may return null, indicating that no valid values can be obtained.
		Id *string `json:"Id,omitempty" name:"Id"`

		// Resource IP
		Ip *string `json:"Ip,omitempty" name:"Ip"`

		// Metric. Valid values: [inqps (total requests peak), dropqps (attack requests peak)]
		MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

		// Statistical granularity. Valid values: [300 (5-minute), 3600 (hourly), 86400 (daily)]
		Period *int64 `json:"Period,omitempty" name:"Period"`

		// Statistics start time
		StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

		// Statistics end time
		EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

		// Value array
		Data []*uint64 `json:"Data,omitempty" name:"Data" list`

		// Number of values
		Count *uint64 `json:"Count,omitempty" name:"Count"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCCTrendResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCCTrendResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCCUrlAllowRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS service type. `bgpip`: Anti-DDoS Advanced; `bgp`: Anti-DDoS Pro (single IP); `bgp-multip`: Anti-DDoS Pro (multi-IP), `net`: Anti-DDoS Ultimate
	Business *string `json:"Business,omitempty" name:"Business"`

	// Anti-DDoS instance ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// Blocklist or allowlist. Valid value: [white (allowlist)]. Currently, only allowlist is supported.
	// Note: this array can only have one value which can only be `white`
	Type []*string `json:"Type,omitempty" name:"Type" list`

	// Pagination parameter
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Pagination parameter
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// HTTP or HTTPS CC protection, which is optional. Valid values: [http (HTTP CC protection), https (HTTPS CC protection)];
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
}

func (r *DescribeCCUrlAllowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCCUrlAllowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "Type")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Protocol")
	if len(f) > 0 {
		return errors.New("DescribeCCUrlAllowRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCCUrlAllowResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// This field has been replaced by `RecordList` and should not be used
		Data []*KeyValue `json:"Data,omitempty" name:"Data" list`

		// Total number of records
		Total *uint64 `json:"Total,omitempty" name:"Total"`

		// Returned Blocklist/allowlist record,
	// If "Key":"url", "Value": URL;
	// If "Key":"domain", "Value": domain name.
	// If "Key":"type", "Value" can be `white` (allowlist) or `black` (blocklist).
	// If "Key":"protocol", "Value": CC protection type (HTTP protection or HTTPS domain name protection);
		RecordList []*KeyValueRecord `json:"RecordList,omitempty" name:"RecordList" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCCUrlAllowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCCUrlAllowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDDoSAlarmThresholdRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS service type (`shield`: Chess Shield, `bgpip`: Anti-DDoS Advanced, `bgp`: Anti-DDoS Pro (single IP), `bgp-multip`: Anti-DDoS Pro (multi-IP), `net`: Anti-DDoS Ultimate)
	Business *string `json:"Business,omitempty" name:"Business"`

	// Anti-DDoS instance ID
	RsId *string `json:"RsId,omitempty" name:"RsId"`
}

func (r *DescribeDDoSAlarmThresholdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSAlarmThresholdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "RsId")
	if len(f) > 0 {
		return errors.New("DescribeDDoSAlarmThresholdRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDDoSAlarmThresholdResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// DDoS alarm threshold
		DDoSAlarmThreshold *DDoSAlarmThreshold `json:"DDoSAlarmThreshold,omitempty" name:"DDoSAlarmThreshold"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDDoSAlarmThresholdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSAlarmThresholdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDDoSAttackIPRegionMapRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS service type (`shield`: Chess Shield, `bgpip`: Anti-DDoS Advanced, `bgp`: Anti-DDoS Pro (single IP), `bgp-multip`: Anti-DDoS Pro (multi-IP), `net`: Anti-DDoS Ultimate)
	Business *string `json:"Business,omitempty" name:"Business"`

	// Anti-DDoS instance ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// Statistics start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Statistics end time. Maximum statistics time range: half a year;
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// IP attack source of specified resource, which is optional
	IpList []*string `json:"IpList,omitempty" name:"IpList" list`
}

func (r *DescribeDDoSAttackIPRegionMapRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSAttackIPRegionMapRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "IpList")
	if len(f) > 0 {
		return errors.New("DescribeDDoSAttackIPRegionMapRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDDoSAttackIPRegionMapResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Global region distribution data
		NationCount []*KeyValueRecord `json:"NationCount,omitempty" name:"NationCount" list`

		// Chinese province distribution data
		ProvinceCount []*KeyValueRecord `json:"ProvinceCount,omitempty" name:"ProvinceCount" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDDoSAttackIPRegionMapResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSAttackIPRegionMapResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDDoSAttackSourceRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS service type. `bgpip`: Anti-DDoS Advanced; `bgp`: Anti-DDoS Pro (single IP); `bgp-multip`: Anti-DDoS Pro (multi-IP), `net`: Anti-DDoS Ultimate
	Business *string `json:"Business,omitempty" name:"Business"`

	// Anti-DDoS instance ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// Start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Number of entries per page. A value of 0 means no pagination
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Page start offset, whose value is (page number - 1) * number of entries per page
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// IP attack source of specified resource, which is optional
	IpList []*string `json:"IpList,omitempty" name:"IpList" list`
}

func (r *DescribeDDoSAttackSourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSAttackSourceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "IpList")
	if len(f) > 0 {
		return errors.New("DescribeDDoSAttackSourceRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDDoSAttackSourceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of attack sources
		Total *uint64 `json:"Total,omitempty" name:"Total"`

		// Attack source list
		AttackSourceList []*DDoSAttackSourceRecord `json:"AttackSourceList,omitempty" name:"AttackSourceList" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDDoSAttackSourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSAttackSourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDDoSCountRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS service type. `bgpip`: Anti-DDoS Advanced; `bgp`: Anti-DDoS Pro (single IP); `bgp-multip`: Anti-DDoS Pro (multi-IP), `net`: Anti-DDoS Ultimate
	Business *string `json:"Business,omitempty" name:"Business"`

	// Anti-DDoS instance ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// Resource IP
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// Statistics start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Statistics end time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Metric. Valid values: [traffic (attack protocol traffic in KB), pkg (number of attack protocol packets), classnum (number of attack events)]
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`
}

func (r *DescribeDDoSCountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSCountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "Ip")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "MetricName")
	if len(f) > 0 {
		return errors.New("DescribeDDoSCountRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDDoSCountResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Anti-DDoS service type. `bgpip`: Anti-DDoS Advanced; `bgp`: Anti-DDoS Pro (single IP); `bgp-multip`: Anti-DDoS Pro (multi-IP), `net`: Anti-DDoS Ultimate
		Business *string `json:"Business,omitempty" name:"Business"`

		// Anti-DDoS instance ID
		Id *string `json:"Id,omitempty" name:"Id"`

		// Resource IP
		Ip *string `json:"Ip,omitempty" name:"Ip"`

		// Statistics start time
		StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

		// Statistics end time
		EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

		// Metric. Valid values: [traffic (attack protocol traffic in KB), pkg (number of attack protocol packets), classnum (number of attack events)]
		MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

		// `Key-Value` array. Valid values of `Key`:
	// If `MetricName` is `traffic`:
	// If `key` is `TcpKBSum`, it indicates the traffic of TCP packets in KB.
	// If `key` is `UdpKBSum`, it indicates the traffic of UDP packets in KB.
	// If `key` is `IcmpKBSum`, it indicates the traffic of ICMP packets in KB.
	// If `key` is `OtherKBSum`, it indicates the traffic of other packets in KB.
	// 
	// If `MetricName` is `pkg`:
	// If `key` is `TcpPacketSum`, it indicates the number of TCP packets.
	// If `key` is `UdpPacketSum`, it indicates the number of UDP packets.
	// If `key` is `IcmpPacketSum`, it indicates the number of ICMP packets.
	// If `key` is `OtherPacketSum`, it indicates the number of other packets.
	// 
	// If `MetricName` is `classnum`:
	// The value of `key` indicates the attack event type. When the `key` is `UNKNOWNFLOOD`, it indicates  unknown attack events.
		Data []*KeyValue `json:"Data,omitempty" name:"Data" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDDoSCountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDDoSDefendStatusRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS service type. `basic`: Anti-DDoS Basic, `bgp`: Anti-DDoS Pro (single IP), `bgp-multip`: Anti-DDoS (multi-IP), `bgpip`: Anti-DDoS Advanced, `net`: Anti-DDoS Ultimate
	Business *string `json:"Business,omitempty" name:"Business"`

	// Instance ID, which is required only if `Business` is not `basic`.
	Id *string `json:"Id,omitempty" name:"Id"`

	// Anti-DDoS Basic IP, which is required only if `Business` is Anti-DDoS Basic;
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// Type of products bound to the anti-DDoS instance, which is required only if `Business` is Anti-DDoS Basic. Valid values: [public (CVM), bm (Bare metal products), eni (elastic network interface), vpngw (VPN Gateway), natgw (NAT Gateway), waf (Web Application Firewall), fpc (Finance products), gaap (GAAP), other (hosted IP)]
	BizType *string `json:"BizType,omitempty" name:"BizType"`

	// Product subtype of IP, which is required only if `Business` is Anti-DDoS Basic. Valid values: [cvm (CVM), lb (CLB), eni (ENI), vpngw (VPN), natgw (NAT), waf (WAF), fpc (finance), gaap (GAAP), other (hosted IP), eip (BM EIP)]
	DeviceType *string `json:"DeviceType,omitempty" name:"DeviceType"`

	// Resource instance ID of IP, which is required only if `Business` is Anti-DDoS Basic. This field is required when binding a new IP. For example, if it is an ENI IP, enter `ID(eni-*)` of the ENI for `InstanceId`;
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// This field is required only if `Business` is Anti-DDoS Basic, indicating the IP region. Valid values:
	// "bj":     North China (Beijing)
	// "cd":     Southwest China (Chengdu)
	// "cq":     Southwest China (Chongqing)
	// "gz":     South China (Guangzhou)
	// "gzopen": South China (Guangzhou Open)
	// "hk":     Hong Kong (China)
	// "kr":     Southeast Asia (Seoul)
	// "sh":     East China (Shanghai)
	// "shjr":   East China (Shanghai Finance)
	// "szjr":   South China (Shenzhen Finance)
	// "sg":     Southeast Asia (Singapore)
	// "th":     Southeast Asia (Thailand)
	// "de":     Europe (Germany)
	// "usw":    West US (Silicon Valley)
	// "ca":     North America (Toronto)
	// "jp":     Japan
	// "hzec":   Hangzhou
	// "in":     India
	// "use":    East US (Virginia)
	// "ru":     Russia
	// "tpe":    Taiwan (China)
	// "nj":     Nanjing
	IPRegion *string `json:"IPRegion,omitempty" name:"IPRegion"`
}

func (r *DescribeDDoSDefendStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSDefendStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "Ip")
	delete(f, "BizType")
	delete(f, "DeviceType")
	delete(f, "InstanceId")
	delete(f, "IPRegion")
	if len(f) > 0 {
		return errors.New("DescribeDDoSDefendStatusRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDDoSDefendStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Protection status. 0: disabled, 1: enabled
	// Note: this field may return null, indicating that no valid values can be obtained.
		DefendStatus *uint64 `json:"DefendStatus,omitempty" name:"DefendStatus"`

		// Expiration time of temporary protection disablement. This field is empty if the protection is in enabled status;
	// Note: this field may return null, indicating that no valid values can be obtained.
		UndefendExpire *string `json:"UndefendExpire,omitempty" name:"UndefendExpire"`

		// Console feature display field. 1: displays console features, 0: hides console features
	// Note: this field may return null, indicating that no valid values can be obtained.
		ShowFlag *uint64 `json:"ShowFlag,omitempty" name:"ShowFlag"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDDoSDefendStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSDefendStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDDoSEvInfoRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS service type. `bgpip`: Anti-DDoS Advanced; `bgp`: Anti-DDoS Pro (single IP); `bgp-multip`: Anti-DDoS Pro (multi-IP), `net`: Anti-DDoS Ultimate
	Business *string `json:"Business,omitempty" name:"Business"`

	// Anti-DDoS instance ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// Resource IP
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// Attack start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Attack end time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeDDoSEvInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSEvInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "Ip")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return errors.New("DescribeDDoSEvInfoRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDDoSEvInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Anti-DDoS service type. `bgpip`: Anti-DDoS Advanced; `bgp`: Anti-DDoS Pro (single IP); `bgp-multip`: Anti-DDoS Pro (multi-IP), `net`: Anti-DDoS Ultimate
		Business *string `json:"Business,omitempty" name:"Business"`

		// Anti-DDoS instance ID
		Id *string `json:"Id,omitempty" name:"Id"`

		// Resource IP
		Ip *string `json:"Ip,omitempty" name:"Ip"`

		// Attack start time
		StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

		// Attack end time
		EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

		// Number of TCP attack packets
		TcpPacketSum *uint64 `json:"TcpPacketSum,omitempty" name:"TcpPacketSum"`

		// Traffic of TCP attacks in KB
		TcpKBSum *uint64 `json:"TcpKBSum,omitempty" name:"TcpKBSum"`

		// Number of UDP attack packets
		UdpPacketSum *uint64 `json:"UdpPacketSum,omitempty" name:"UdpPacketSum"`

		// Traffic of UDP attacks in KB
		UdpKBSum *uint64 `json:"UdpKBSum,omitempty" name:"UdpKBSum"`

		// Number of ICMP attack packets
		IcmpPacketSum *uint64 `json:"IcmpPacketSum,omitempty" name:"IcmpPacketSum"`

		// Traffic of ICMP attacks in KB
		IcmpKBSum *uint64 `json:"IcmpKBSum,omitempty" name:"IcmpKBSum"`

		// Number of other attack packets
		OtherPacketSum *uint64 `json:"OtherPacketSum,omitempty" name:"OtherPacketSum"`

		// Traffic of other attacks in KB
		OtherKBSum *uint64 `json:"OtherKBSum,omitempty" name:"OtherKBSum"`

		// Total attack traffic in KB
		TotalTraffic *uint64 `json:"TotalTraffic,omitempty" name:"TotalTraffic"`

		// Attack traffic bandwidth peak
		Mbps *uint64 `json:"Mbps,omitempty" name:"Mbps"`

		// Attack packet rate peak
		Pps *uint64 `json:"Pps,omitempty" name:"Pps"`

		// PCAP file download link
		PcapUrl []*string `json:"PcapUrl,omitempty" name:"PcapUrl" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDDoSEvInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSEvInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDDoSEvListRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS service type. `bgpip`: Anti-DDoS Advanced; `bgp`: Anti-DDoS Pro (Single IP); `bgp-multip`: Anti-DDoS Pro (Multi-IP); `net`: Anti-DDoS Ultimate; `basic`: Anti-DDoS Basic
	Business *string `json:"Business,omitempty" name:"Business"`

	// Start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Resource instance ID. If `Business` is `basic`, this field is not required (because Anti-DDoS Basic has no resource instance)
	Id *string `json:"Id,omitempty" name:"Id"`

	// Resource IP
	IpList []*string `json:"IpList,omitempty" name:"IpList" list`

	// Whether the elastic protection bandwidth is exceeded. Valid values: [yes, no]. If an empty string is entered, it means no filtering
	OverLoad *string `json:"OverLoad,omitempty" name:"OverLoad"`

	// Number of entries per page. A value of 0 means no pagination
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Page start offset, whose value is (page number - 1) * number of entries per page
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeDDoSEvListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSEvListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Id")
	delete(f, "IpList")
	delete(f, "OverLoad")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return errors.New("DescribeDDoSEvListRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDDoSEvListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Anti-DDoS service type. `bgpip`: Anti-DDoS Advanced; `bgp`: Anti-DDoS Pro (Single IP); `bgp-multip`: Anti-DDoS Pro (Multi-IP); `net`: Anti-DDoS Ultimate; `basic`: Anti-DDoS Basic
		Business *string `json:"Business,omitempty" name:"Business"`

		// Anti-DDoS instance ID
		Id *string `json:"Id,omitempty" name:"Id"`

		// Resource IP
	// Note: this field may return null, indicating that no valid values can be obtained.
		IpList []*string `json:"IpList,omitempty" name:"IpList" list`

		// Start time
		StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

		// End time
		EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

		// DDoS attack event list
		Data []*DDoSEventRecord `json:"Data,omitempty" name:"Data" list`

		// Total number of records
		Total *uint64 `json:"Total,omitempty" name:"Total"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDDoSEvListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSEvListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDDoSIpLogRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS service type (`net`: Anti-DDoS Ultimate)
	Business *string `json:"Business,omitempty" name:"Business"`

	// Anti-DDoS instance ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// Resource IP
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// Attack start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Attack end time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeDDoSIpLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSIpLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "Ip")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return errors.New("DescribeDDoSIpLogRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDDoSIpLogResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Anti-DDoS service type (`net`: Anti-DDoS Ultimate)
		Business *string `json:"Business,omitempty" name:"Business"`

		// Anti-DDoS instance ID
		Id *string `json:"Id,omitempty" name:"Id"`

		// Resource IP
		Ip *string `json:"Ip,omitempty" name:"Ip"`

		// Attack start time
		StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

		// Attack end time
		EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

		// IP attack log, which is a `KeyValue` array. Valid values of `Key-Value`:
	// If `Key` is `LogTime`, `Value` indicates the IP log time
	// If `Key` is `LogMessage`, `Value` indicates the IP log content
		Data []*KeyValueRecord `json:"Data,omitempty" name:"Data" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDDoSIpLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSIpLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDDoSNetCountRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS service type (`net`: Anti-DDoS Ultimate)
	Business *string `json:"Business,omitempty" name:"Business"`

	// Anti-DDoS instance ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// Statistics start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Statistics end time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Metric. Valid values: [traffic (attack protocol traffic in KB), pkg (number of attack protocol packets), classnum (number of attack events)]
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`
}

func (r *DescribeDDoSNetCountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSNetCountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "MetricName")
	if len(f) > 0 {
		return errors.New("DescribeDDoSNetCountRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDDoSNetCountResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Anti-DDoS service type (`net`: Anti-DDoS Ultimate)
		Business *string `json:"Business,omitempty" name:"Business"`

		// Anti-DDoS instance ID
		Id *string `json:"Id,omitempty" name:"Id"`

		// Statistics start time
		StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

		// Statistics end time
		EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

		// Metric. Valid values: [traffic (attack protocol traffic in KB), pkg (number of attack protocol packets), classnum (number of attack events)]
		MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

		// `Key-Value` array. Valid values of `Key`:
	// If `MetricName` is `traffic`:
	// If `key` is `TcpKBSum`, it indicates the traffic of TCP packets in KB.
	// If `key` is `UdpKBSum`, it indicates the traffic of UDP packets in KB.
	// If `key` is `IcmpKBSum`, it indicates the traffic of ICMP packets in KB.
	// If `key` is `OtherKBSum`, it indicates the traffic of other packets in KB.
	// 
	// If `MetricName` is `pkg`:
	// If `key` is `TcpPacketSum`, it indicates the number of TCP packets.
	// If `key` is `UdpPacketSum`, it indicates the number of UDP packets.
	// If `key` is `IcmpPacketSum`, it indicates the number of ICMP packets.
	// If `key` is `OtherPacketSum`, it indicates the number of other packets.
	// 
	// If `MetricName` is `classnum`:
	// The value of `key` indicates the attack event type. When the `key` is `UNKNOWNFLOOD`, it indicates  unknown attack events.
		Data []*KeyValue `json:"Data,omitempty" name:"Data" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDDoSNetCountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSNetCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDDoSNetEvInfoRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS service type (`net`: Anti-DDoS Ultimate)
	Business *string `json:"Business,omitempty" name:"Business"`

	// Anti-DDoS instance ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// Attack start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Attack end time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeDDoSNetEvInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSNetEvInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return errors.New("DescribeDDoSNetEvInfoRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDDoSNetEvInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Anti-DDoS service type (`net`: Anti-DDoS Ultimate)
		Business *string `json:"Business,omitempty" name:"Business"`

		// Anti-DDoS instance ID
		Id *string `json:"Id,omitempty" name:"Id"`

		// Attack start time
		StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

		// Attack end time
		EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

		// Number of TCP attack packets
		TcpPacketSum *uint64 `json:"TcpPacketSum,omitempty" name:"TcpPacketSum"`

		// Traffic of TCP attacks in KB
		TcpKBSum *uint64 `json:"TcpKBSum,omitempty" name:"TcpKBSum"`

		// Number of UDP attack packets
		UdpPacketSum *uint64 `json:"UdpPacketSum,omitempty" name:"UdpPacketSum"`

		// Traffic of UDP attacks in KB
		UdpKBSum *uint64 `json:"UdpKBSum,omitempty" name:"UdpKBSum"`

		// Number of ICMP attack packets
		IcmpPacketSum *uint64 `json:"IcmpPacketSum,omitempty" name:"IcmpPacketSum"`

		// Traffic of ICMP attacks in KB
		IcmpKBSum *uint64 `json:"IcmpKBSum,omitempty" name:"IcmpKBSum"`

		// Number of other attack packets
		OtherPacketSum *uint64 `json:"OtherPacketSum,omitempty" name:"OtherPacketSum"`

		// Traffic of other attacks in KB
		OtherKBSum *uint64 `json:"OtherKBSum,omitempty" name:"OtherKBSum"`

		// Total attack traffic in KB
		TotalTraffic *uint64 `json:"TotalTraffic,omitempty" name:"TotalTraffic"`

		// Attack traffic bandwidth peak
		Mbps *uint64 `json:"Mbps,omitempty" name:"Mbps"`

		// Attack packet rate peak
		Pps *uint64 `json:"Pps,omitempty" name:"Pps"`

		// PCAP file download link
		PcapUrl []*string `json:"PcapUrl,omitempty" name:"PcapUrl" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDDoSNetEvInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSNetEvInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDDoSNetEvListRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS service type (`net`: Anti-DDoS Ultimate)
	Business *string `json:"Business,omitempty" name:"Business"`

	// Anti-DDoS instance ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// Start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Number of entries per page. A value of 0 means no pagination
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Page start offset, whose value is (page number - 1) * number of entries per page
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeDDoSNetEvListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSNetEvListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return errors.New("DescribeDDoSNetEvListRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDDoSNetEvListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Anti-DDoS service type (`net`: Anti-DDoS Ultimate)
		Business *string `json:"Business,omitempty" name:"Business"`

		// Anti-DDoS instance ID
		Id *string `json:"Id,omitempty" name:"Id"`

		// Start time
		StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

		// End time
		EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

		// DDoS attack event list
		Data []*DDoSEventRecord `json:"Data,omitempty" name:"Data" list`

		// Total number of records
		Total *uint64 `json:"Total,omitempty" name:"Total"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDDoSNetEvListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSNetEvListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDDoSNetIpLogRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS service type (`net`: Anti-DDoS Ultimate)
	Business *string `json:"Business,omitempty" name:"Business"`

	// Anti-DDoS instance ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// Attack start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Attack end time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeDDoSNetIpLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSNetIpLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return errors.New("DescribeDDoSNetIpLogRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDDoSNetIpLogResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Anti-DDoS service type (`net`: Anti-DDoS Ultimate)
		Business *string `json:"Business,omitempty" name:"Business"`

		// Anti-DDoS instance ID
		Id *string `json:"Id,omitempty" name:"Id"`

		// Attack start time
		StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

		// Attack end time
		EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

		// IP attack log, which is a `KeyValue` array. Valid values of `Key-Value`:
	// If `Key` is `LogTime`, `Value` indicates the IP log time
	// If `Key` is `LogMessage`, `Value` indicates the IP log content
		Data []*KeyValueRecord `json:"Data,omitempty" name:"Data" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDDoSNetIpLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSNetIpLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDDoSNetTrendRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS service type (`net`: Anti-DDoS Ultimate)
	Business *string `json:"Business,omitempty" name:"Business"`

	// Anti-DDoS instance ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// Metric. Valid values: [bps (attack traffic bandwidth), pps (attack packet rate)]
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// Statistical granularity. Valid values: [300 (5-minute), 3600 (hourly), 86400 (daily)]
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// Statistics start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Statistics end time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeDDoSNetTrendRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSNetTrendRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "MetricName")
	delete(f, "Period")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return errors.New("DescribeDDoSNetTrendRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDDoSNetTrendResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Anti-DDoS service type (`net`: Anti-DDoS Ultimate)
		Business *string `json:"Business,omitempty" name:"Business"`

		// Anti-DDoS instance ID
		Id *string `json:"Id,omitempty" name:"Id"`

		// Metric. Valid values: [bps (attack traffic bandwidth), pps (attack packet rate)]
		MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

		// Statistical granularity. Valid values: [300 (5-minute), 3600 (hourly), 86400 (daily)]
		Period *uint64 `json:"Period,omitempty" name:"Period"`

		// Statistics start time
		StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

		// Statistics end time
		EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

		// Value array
		Data []*uint64 `json:"Data,omitempty" name:"Data" list`

		// Number of values
		Count *uint64 `json:"Count,omitempty" name:"Count"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDDoSNetTrendResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSNetTrendResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDDoSPolicyRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS service type. `bgpip`: Anti-DDoS Advanced; `bgp`: Anti-DDoS Pro (single IP); `bgp-multip`: Anti-DDoS Pro (multi-IP), `net`: Anti-DDoS Ultimate
	Business *string `json:"Business,omitempty" name:"Business"`

	// Resource ID, which is optional. If a value is entered, it indicates the advanced DDoS policy bound to the resource
	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeDDoSPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	if len(f) > 0 {
		return errors.New("DescribeDDoSPolicyRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDDoSPolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Advanced DDoS policy list
		DDosPolicyList []*DDosPolicy `json:"DDosPolicyList,omitempty" name:"DDosPolicyList" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDDoSPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDDoSTrendRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS service type. `bgpip`: Anti-DDoS Advanced; `bgp`: Anti-DDoS Pro (Single IP); `bgp-multip`: Anti-DDoS Pro (Multi-IP); `net`: Anti-DDoS Ultimate; `basic`: Anti-DDoS Basic
	Business *string `json:"Business,omitempty" name:"Business"`

	// Anti-DDoS instance IP
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// Metric. Valid values: [bps (attack traffic bandwidth), pps (attack packet rate)]
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// Statistical granularity. Valid values: [300 (5-minute), 3600 (hourly), 86400 (daily)]
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// Statistics start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Statistics end time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Resource instance ID. If `Business` is `basic`, this field is not required (because Anti-DDoS Basic has no resource instance)
	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeDDoSTrendRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSTrendRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Ip")
	delete(f, "MetricName")
	delete(f, "Period")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Id")
	if len(f) > 0 {
		return errors.New("DescribeDDoSTrendRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDDoSTrendResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Anti-DDoS service type. `bgpip`: Anti-DDoS Advanced; `bgp`: Anti-DDoS Pro (Single IP); `bgp-multip`: Anti-DDoS Pro (Multi-IP); `net`: Anti-DDoS Ultimate; `basic`: Anti-DDoS Basic
		Business *string `json:"Business,omitempty" name:"Business"`

		// Anti-DDoS instance ID
	// Note: this field may return null, indicating that no valid values can be obtained.
		Id *string `json:"Id,omitempty" name:"Id"`

		// Resource IP
		Ip *string `json:"Ip,omitempty" name:"Ip"`

		// Metric. Valid values: [bps (attack traffic bandwidth), pps (attack packet rate)]
		MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

		// Statistical granularity. Valid values: [300 (5-minute), 3600 (hourly), 86400 (daily)]
		Period *int64 `json:"Period,omitempty" name:"Period"`

		// Statistics start time
		StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

		// Statistics end time
		EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

		// Value array. The unit for attack traffic bandwidth is Mbps, and that for the packet rate is pps.
		Data []*uint64 `json:"Data,omitempty" name:"Data" list`

		// Number of values
		Count *uint64 `json:"Count,omitempty" name:"Count"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDDoSTrendResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSTrendResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDDoSUsedStatisRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS service type (`bgpip`: Anti-DDoS Advanced)
	Business *string `json:"Business,omitempty" name:"Business"`
}

func (r *DescribeDDoSUsedStatisRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSUsedStatisRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	if len(f) > 0 {
		return errors.New("DescribeDDoSUsedStatisRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDDoSUsedStatisResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Field value as follows:
	// Days: number of days of Anti-DDoS resource use
	// Attacks: number of DDoS attacks
		Data []*KeyValue `json:"Data,omitempty" name:"Data" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDDoSUsedStatisResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSUsedStatisResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIPProductInfoRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS service type. `bgp`: Anti-DDoS Pro (Single IP); `bgp-multip`: Anti-DDoS Pro (Multi-IP)
	Business *string `json:"Business,omitempty" name:"Business"`

	// IP list
	IpList []*string `json:"IpList,omitempty" name:"IpList" list`
}

func (r *DescribeIPProductInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIPProductInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "IpList")
	if len(f) > 0 {
		return errors.New("DescribeIPProductInfoRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIPProductInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Tencent Cloud product information list. If nothing is found, an empty array will be returned. Valid values:
	// If `Key` is ProductName, `value` indicates the name of a Tencent Cloud product instance;
	// If `Key` is `ProductInstanceId`, `value` indicates the ID of a Tencent Cloud product instance;
	// If `Key` is `ProductType`, `value` indicates the Tencent Cloud product type (cvm: CVM, clb: CLB);
	// If `Key` is `IP`, `value` indicates the IP of a Tencent Cloud product instance;
		Data []*KeyValueRecord `json:"Data,omitempty" name:"Data" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeIPProductInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIPProductInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInsurePacksRequest struct {
	*tchttp.BaseRequest

	// Guarantee package ID, which is optional. If you need to get the guarantee package with a specified ID (such as insure-000000xe), please use this field
	IdList []*string `json:"IdList,omitempty" name:"IdList" list`
}

func (r *DescribeInsurePacksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInsurePacksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IdList")
	if len(f) > 0 {
		return errors.New("DescribeInsurePacksRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInsurePacksResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Guarantee package list
		InsurePacks []*KeyValueRecord `json:"InsurePacks,omitempty" name:"InsurePacks" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInsurePacksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInsurePacksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIpBlockListRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeIpBlockListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIpBlockListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.New("DescribeIpBlockListRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIpBlockListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Blocked IP list
		List []*IpBlockData `json:"List,omitempty" name:"List" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeIpBlockListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIpBlockListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIpUnBlockListRequest struct {
	*tchttp.BaseRequest

	// Start time
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// End time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// IP (if this field is not empty, IP filtering will be performed)
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// Pagination parameter (paginated query will be performed if this field is not empty). This field will be disused in the future. Please use the `Limit` and `Offset` fields instead;
	Paging *Paging `json:"Paging,omitempty" name:"Paging"`

	// Number of entries per page. A value of 0 means no pagination
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Page start offset, whose value is (page number - 1) * number of entries per page
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeIpUnBlockListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIpUnBlockListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BeginTime")
	delete(f, "EndTime")
	delete(f, "Ip")
	delete(f, "Paging")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return errors.New("DescribeIpUnBlockListRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIpUnBlockListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Start time
		BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

		// End time
		EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

		// IP unblocking record
		List []*IpUnBlockData `json:"List,omitempty" name:"List" list`

		// Total number of records
		Total *uint64 `json:"Total,omitempty" name:"Total"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeIpUnBlockListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIpUnBlockListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeL4HealthConfigRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS service type. `bgpip`: Anti-DDoS Advanced; `net`: Anti-DDoS Ultimate
	Business *string `json:"Business,omitempty" name:"Business"`

	// Anti-DDoS instance ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// Rule ID array. To export the health check configurations of all rules, leave this field empty or enter an empty array;
	RuleIdList []*string `json:"RuleIdList,omitempty" name:"RuleIdList" list`
}

func (r *DescribeL4HealthConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeL4HealthConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "RuleIdList")
	if len(f) > 0 {
		return errors.New("DescribeL4HealthConfigRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeL4HealthConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Layer-4 health check configuration array
		HealthConfig []*L4HealthConfig `json:"HealthConfig,omitempty" name:"HealthConfig" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeL4HealthConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeL4HealthConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeL4RulesErrHealthRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS service type. `bgpip`: Anti-DDoS Advanced; `net`: Anti-DDoS Ultimate
	Business *string `json:"Business,omitempty" name:"Business"`

	// Anti-DDoS instance ID
	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeL4RulesErrHealthRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeL4RulesErrHealthRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	if len(f) > 0 {
		return errors.New("DescribeL4RulesErrHealthRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeL4RulesErrHealthResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of exceptional rules
		Total *uint64 `json:"Total,omitempty" name:"Total"`

		// Exceptional rule list. Returned value description: `Key` is the rule ID, while `Value` is the exceptional IP. Multiple IPs are separated by ","
		ErrHealths []*KeyValue `json:"ErrHealths,omitempty" name:"ErrHealths" list`

		// Exceptional rule list (which provides more error-related information). Returned value description:
	// If `key` is `RuleId`, `Value` indicates the rule ID;
	// If `key` is `Protocol`, `Value` indicates the forwarding protocol of a rule;
	// If `key` is `VirtualPort`, `Value` indicates the forwarding port of a rule;
	// If `Key` is `ErrMessage`, `Value` indicates the exception message for health check;
	// Exception message for health check in the format of `"SourceIp:1.1.1.1|SourcePort:1234|AbnormalStatTime:1570689065|AbnormalReason:connection time out|Interval:20|CheckNum:6|FailNum:6"`. Multiple error messages for the source IP should be separated by `,`,
	// SourceIp: real server IP, SourcePort: real server port, AbnormalStatTime: exception time, AbnormalReason: cause of exception, Interval: check frequency, CheckNum: number of checks, FailNum: number of failures;
		ExtErrHealths []*KeyValueRecord `json:"ExtErrHealths,omitempty" name:"ExtErrHealths" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeL4RulesErrHealthResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeL4RulesErrHealthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeL7HealthConfigRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS service type. `bgpip`: Anti-DDoS Advanced; `net`: Anti-DDoS Ultimate
	Business *string `json:"Business,omitempty" name:"Business"`

	// Anti-DDoS instance ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// Rule ID array. To export the health check configurations of all rules, leave this field empty or enter an empty array;
	RuleIdList []*string `json:"RuleIdList,omitempty" name:"RuleIdList" list`
}

func (r *DescribeL7HealthConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeL7HealthConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "RuleIdList")
	if len(f) > 0 {
		return errors.New("DescribeL7HealthConfigRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeL7HealthConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Layer-7 health check configuration array
		HealthConfig []*L7HealthConfig `json:"HealthConfig,omitempty" name:"HealthConfig" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeL7HealthConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeL7HealthConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePackIndexRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS service type. `bgpip`: Anti-DDoS Advanced, `bgp`: Anti-DDoS Pro (single IP), `net`: Anti-DDoS Ultimate
	Business *string `json:"Business,omitempty" name:"Business"`
}

func (r *DescribePackIndexRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePackIndexRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	if len(f) > 0 {
		return errors.New("DescribePackIndexRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribePackIndexResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Field value as follows:
	// TotalPackCount: number of resources
	// AttackPackCount: number of resources being cleansed
	// BlockPackCount: number of blocked resources
	// ExpiredPackCount: number of expired resources
	// ExpireingPackCount: number of expiring resources
	// IsolatePackCount: number of isolated resources
		Data []*KeyValue `json:"Data,omitempty" name:"Data" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePackIndexResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePackIndexResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePcapRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS service type. `bgpip`: Anti-DDoS Advanced; `bgp`: Anti-DDoS Pro (single IP); `bgp-multip`: Anti-DDoS Pro (multi-IP), `net`: Anti-DDoS Ultimate
	Business *string `json:"Business,omitempty" name:"Business"`

	// Anti-DDoS instance ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// Attack event start time in the format of "2018-08-28 07:00:00"
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Attack event end time in the format of "2018-08-28 07:02:00"
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Resource IP, which is required only if `Business` is `net`;
	Ip *string `json:"Ip,omitempty" name:"Ip"`
}

func (r *DescribePcapRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePcapRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Ip")
	if len(f) > 0 {
		return errors.New("DescribePcapRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribePcapResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// PCAP packet download link list, which is an empty array if there are no PCAP packets;
		PcapUrlList []*string `json:"PcapUrlList,omitempty" name:"PcapUrlList" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePcapResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePcapResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePolicyCaseRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS service type. `bgpip`: Anti-DDoS Advanced; `bgp`: Anti-DDoS Pro (single IP); `bgp-multip`: Anti-DDoS Pro (multi-IP), `net`: Anti-DDoS Ultimate
	Business *string `json:"Business,omitempty" name:"Business"`

	// Policy scenario ID
	SceneId *string `json:"SceneId,omitempty" name:"SceneId"`
}

func (r *DescribePolicyCaseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePolicyCaseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "SceneId")
	if len(f) > 0 {
		return errors.New("DescribePolicyCaseRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribePolicyCaseResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Policy scenario list
		CaseList []*KeyValueRecord `json:"CaseList,omitempty" name:"CaseList" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePolicyCaseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePolicyCaseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeResIpListRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS service type. `bgpip`: Anti-DDoS Advanced; `bgp`: Anti-DDoS Pro (single IP); `bgp-multip`: Anti-DDoS Pro (multi-IP), `net`: Anti-DDoS Ultimate
	Business *string `json:"Business,omitempty" name:"Business"`

	// Resource ID. If this field is left empty, it means to get all resources IP of the current user
	IdList []*string `json:"IdList,omitempty" name:"IdList" list`
}

func (r *DescribeResIpListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResIpListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "IdList")
	if len(f) > 0 {
		return errors.New("DescribeResIpListRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeResIpListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Resource IP list
		Resource []*ResourceIp `json:"Resource,omitempty" name:"Resource" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeResIpListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResIpListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeResourceListRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS service type. `bgp`: Anti-DDoS Pro (single IP), `bgp-multip`: Anti-DDoS Pro (multi-IP), `bgpip`: Anti-DDoS Advanced, `net`: Anti-DDoS Ultimate)
	Business *string `json:"Business,omitempty" name:"Business"`

	// Region code search, which is optional. If no regions are to be specified, enter an empty array. If a region is to be specified, enter a region code, such as ["gz", "sh"]
	RegionList []*string `json:"RegionList,omitempty" name:"RegionList" list`

	// Line search. This field can be optionally entered only when getting the list of Anti-DDoS Advanced resources. Valid values: [1 (BGP line), 2 (Nanjing Telecom), 3 (Nanjing Unicom), 99 (third-party partner line)]. Please enter an empty array when getting other products;
	Line []*uint64 `json:"Line,omitempty" name:"Line" list`

	// Resource ID search, which is optional. If this field is not an empty array, it means to get the resource list of a specified resource;
	IdList []*string `json:"IdList,omitempty" name:"IdList" list`

	// Resource name search, which is optional. If this field is not an empty string, it means to search for resources by name;
	Name *string `json:"Name,omitempty" name:"Name"`

	// IP query list, which is optional. Resources will be queried by IP if the list is not empty.
	IpList []*string `json:"IpList,omitempty" name:"IpList" list`

	// Resource status search list, which is optional. Valid values: [0 (running), 1 (cleansing), 2 (blocking)]. No status search will be performed if an empty array is entered;
	Status []*uint64 `json:"Status,omitempty" name:"Status" list`

	// Expiring resource search, which is optional. Valid values: [0 (no search), 1 (search for expiring resources)]
	Expire *uint64 `json:"Expire,omitempty" name:"Expire"`

	// Sort by field, which is optional
	OderBy []*OrderBy `json:"OderBy,omitempty" name:"OderBy" list`

	// Number of entries per page. A value of 0 means no pagination
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Page start offset, whose value is (page number - 1) * number of entries per page
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// CNAME of Anti-DDoS Ultimate resource, which is optional and only valid for the Anti-DDoS Ultimate resource list;
	CName *string `json:"CName,omitempty" name:"CName"`

	// Domain name of Anti-DDoS Ultimate resource, which is optional and only valid for the Anti-DDoS Ultimate resource list;
	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *DescribeResourceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourceListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "RegionList")
	delete(f, "Line")
	delete(f, "IdList")
	delete(f, "Name")
	delete(f, "IpList")
	delete(f, "Status")
	delete(f, "Expire")
	delete(f, "OderBy")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "CName")
	delete(f, "Domain")
	if len(f) > 0 {
		return errors.New("DescribeResourceListRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeResourceListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of records
		Total *uint64 `json:"Total,omitempty" name:"Total"`

		// Resource record list. The description of key values is as follows:
	// "Key": "CreateTime" (Instance purchase time)
	// "Key": "Region" (Instance region)
	// "Key": "BoundIP" (IP bound to the single-IP instance)
	// "Key": "Id" (Instance ID)
	// "Key": "CCEnabled" (CC protection switch status of the instance)
	// "Key": "DDoSThreshold" (Anti-DDoS cleansing threshold of the instance)	
	// "Key": "BoundStatus" (IP binding status of the single-IP/multi-IP instance; binding or bound)
	// "Key": "Type" (Disused field)
	// "Key": "ElasticLimit" (Elastic protection value of the instance)
	// "Key": "DDoSAI" (Anti-DDoS AI protection switch of the instance)
	// "Key": "OverloadCount" (The number of attacks exceeding the elastic protection value to the instance)
	// "Key": "Status" (Instance status; idle: running; attacking: under attacks; blocking: being blocked; isolate: being isolated)
	// "Key": "Lbid" (Disused field)
	// "Key": "ShowFlag" (Disused field)
	// "Key": "Expire" (Instance expiry time)
	// "Key": "CCThreshold" (CC protection trigger value of the instance)
	// "Key": "AutoRenewFlag" (Whether the instance is on auto-renewal)
	// "Key": "IspCode" (Line of the single-IP/multi-IP instance; 0: China Telecom; 1: China Unicom; 2: China Mobile; 5: BGP)
	// "Key": "PackType" (Package type)
	// "Key": "PackId" (Package ID)
	// "Key": "Name" (Instance name)
	// "Key": "Locked" (Disused field)
	// "Key": "IpDDoSLevel" (Protection level of the instance; low: loose; middle: normal; high: strict)
	// "Key": "DefendStatus" (DDoS protection status of the instance; enabled or temporarily disabled)
	// "Key": "UndefendExpire" (End time of the temporary disabling on DDoS protection for the instance)
	// "Key": "Tgw" (Whether it is a new instance)
	// "Key": "Bandwidth" (Base protection value of the Anti-DDoS Pro/Advanced instance)
	// "Key": "DdosMax" (Base protection value of the Anti-DDoS Ultimate instance)
	// "Key": "GFBandwidth" (Base business application bandwidth of the Anti-DDoS Advanced instance)
	// "Key": "ServiceBandwidth" (Base business application bandwidth of the Anti-DDoS Ultimate instance)
		ServicePacks []*KeyValueRecord `json:"ServicePacks,omitempty" name:"ServicePacks" list`

		// Anti-DDoS service type. `bgp`: Anti-DDoS Pro (single IP), `bgp-multip`: Anti-DDoS Pro (multi-IP), `bgpip`: Anti-DDoS Advanced, `net`: Anti-DDoS Ultimate)
		Business *string `json:"Business,omitempty" name:"Business"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeResourceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRuleSetsRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS service type. `bgpip`: Anti-DDoS Advanced; `net`: Anti-DDoS Ultimate
	Business *string `json:"Business,omitempty" name:"Business"`

	// Resource ID list
	IdList []*string `json:"IdList,omitempty" name:"IdList" list`
}

func (r *DescribeRuleSetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRuleSetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "IdList")
	if len(f) > 0 {
		return errors.New("DescribeRuleSetsRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRuleSetsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Rule record array. Valid values:
	// If `Key` is "Id", `Value` indicates the resource ID
	// If `Key` is "RuleIdList", `Value` indicates the resource rule ID. Multiple rule IDs are separated by ","
	// If `Key` is "RuleNameList", `Value` indicates the resource rule name. Multiple rule names are separated by ","
	// If `Key` is "RuleNum", `Value` indicates the number of resource rules
		L4RuleSets []*KeyValueRecord `json:"L4RuleSets,omitempty" name:"L4RuleSets" list`

		// Rule record array. Valid values:
	// If `Key` is "Id", `Value` indicates the resource ID
	// If `Key` is "RuleIdList", `Value` indicates the resource rule ID. Multiple rule IDs are separated by ","
	// If `Key` is "RuleNameList", `Value` indicates the resource rule name. Multiple rule names are separated by ","
	// If `Key` is "RuleNum", `Value` indicates the number of resource rules
		L7RuleSets []*KeyValueRecord `json:"L7RuleSets,omitempty" name:"L7RuleSets" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRuleSetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRuleSetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSchedulingDomainListRequest struct {
	*tchttp.BaseRequest

	// Number of items in a page. Returned results are not paged if you enter '0'.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Starting offset of the page. Value: (number of pages - 1) * items per page
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// (Optional) Filter by specific domain name
	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *DescribeSchedulingDomainListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSchedulingDomainListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Domain")
	if len(f) > 0 {
		return errors.New("DescribeSchedulingDomainListRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSchedulingDomainListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of scheduling domain names
		Total *uint64 `json:"Total,omitempty" name:"Total"`

		// List of scheduling domain names
		DomainList []*SchedulingDomain `json:"DomainList,omitempty" name:"DomainList" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSchedulingDomainListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSchedulingDomainListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecIndexRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeSecIndexRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecIndexRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.New("DescribeSecIndexRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecIndexResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Field value as follows:
	// AttackIpCount: number of attacked IPs
	// AttackCount: number of attacks
	// BlockCount: number of blockings
	// MaxMbps: attack bandwidth peak in Mbps
	// IpNum: IP statistics
		Data []*KeyValue `json:"Data,omitempty" name:"Data" list`

		// Start time of the current month
		BeginDate *string `json:"BeginDate,omitempty" name:"BeginDate"`

		// End time of the current month
		EndDate *string `json:"EndDate,omitempty" name:"EndDate"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSecIndexResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecIndexResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSourceIpSegmentRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS service type. `bgpip`: Anti-DDoS Advanced; `net`: Anti-DDoS Ultimate
	Business *string `json:"Business,omitempty" name:"Business"`

	// Anti-DDoS instance ID
	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeSourceIpSegmentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSourceIpSegmentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	if len(f) > 0 {
		return errors.New("DescribeSourceIpSegmentRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSourceIpSegmentResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Intermediate IP range. Multiple values are separated by ";"
		Data *string `json:"Data,omitempty" name:"Data"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSourceIpSegmentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSourceIpSegmentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTransmitStatisRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS service type. `bgpip`: Anti-DDoS Advanced, `net`: Anti-DDoS Ultimate, `bgp`: Anti-DDoS Pro (single IP), `bgp-multip`: Anti-DDoS Pro (multi-IP)
	Business *string `json:"Business,omitempty" name:"Business"`

	// Anti-DDoS instance ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// Metric name. Valid values:
	// traffic: traffic bandwidth;
	// pkg: packet rate;
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// Statistical time granularity (300: 5-minute, 3600: hourly, 86400: daily)
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// Statistics start time. The second part is kept at 0, and the minute part is a multiple of 5
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Statistics end time. The second part is kept at 0, and the minute part is a multiple of 5
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Resource IP, which is required and only supports one IP if `Business` is `bgp-multip`. If this field is left empty, all IPs of a resource instance will be counted by default. If the resource instance has multiple IPs (such as Anti-DDoS Ultimate), the statistical method is summation;
	IpList []*string `json:"IpList,omitempty" name:"IpList" list`
}

func (r *DescribeTransmitStatisRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTransmitStatisRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "MetricName")
	delete(f, "Period")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "IpList")
	if len(f) > 0 {
		return errors.New("DescribeTransmitStatisRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTransmitStatisResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// If `MetricName` is `traffic`, this field indicates the inbound traffic bandwidth in bps;
	// If `MetricName` is `pkg`, this field indicates the inbound packet rate in pps;
		InDataList []*float64 `json:"InDataList,omitempty" name:"InDataList" list`

		// If `MetricName` is `traffic`, this field indicates the outbound traffic bandwidth in bps;
	// If `MetricName` is `pkg`, this field indicates the outbound packet rate in pps;
		OutDataList []*float64 `json:"OutDataList,omitempty" name:"OutDataList" list`

		// Metric name:
	// traffic: traffic bandwidth;
	// pkg: packet rate;
		MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTransmitStatisResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTransmitStatisResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUnBlockStatisRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeUnBlockStatisRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUnBlockStatisRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.New("DescribeUnBlockStatisRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUnBlockStatisResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of unblocking chances
		Total *uint64 `json:"Total,omitempty" name:"Total"`

		// Number of used chances
		Used *uint64 `json:"Used,omitempty" name:"Used"`

		// Statistics start time
		BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

		// Statistics end time
		EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUnBlockStatisResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUnBlockStatisResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribleL4RulesRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS service type. `bgpip`: Anti-DDoS Advanced; `net`: Anti-DDoS Ultimate
	Business *string `json:"Business,omitempty" name:"Business"`

	// Anti-DDoS instance ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// Rule ID, which is optional. If this field is entered, the specified rule will be obtained
	RuleIdList []*string `json:"RuleIdList,omitempty" name:"RuleIdList" list`

	// Number of entries per page. A value of 0 means no pagination
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Page start offset, whose value is (page number - 1) * number of entries per page
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribleL4RulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribleL4RulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "RuleIdList")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return errors.New("DescribleL4RulesRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribleL4RulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Forwarding rule list
		Rules []*L4RuleEntry `json:"Rules,omitempty" name:"Rules" list`

		// Total number of rules
		Total *uint64 `json:"Total,omitempty" name:"Total"`

		// Health check configuration list
		Healths []*L4RuleHealth `json:"Healths,omitempty" name:"Healths" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribleL4RulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribleL4RulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribleL7RulesRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS service type. `bgpip`: Anti-DDoS Advanced; `net`: Anti-DDoS Ultimate
	Business *string `json:"Business,omitempty" name:"Business"`

	// Anti-DDoS instance ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// Rule ID, which is optional. If this field is entered, the specified rule will be obtained
	RuleIdList []*string `json:"RuleIdList,omitempty" name:"RuleIdList" list`

	// Number of entries per page. A value of 0 means no pagination
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Page start offset, whose value is (page number - 1) * number of entries per page
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Domain name search, which is optional. Enter it if you need to search for domain names
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Forwarding protocol search, which is optional. Valid values: [http, https, http/https]
	ProtocolList []*string `json:"ProtocolList,omitempty" name:"ProtocolList" list`

	// Status search, which is optional. Valid values: [0 (successfully configured rule), 1 (rule configuration taking effect), 2 (rule configuration failed), 3 (rule deletion taking effect), 5 (rule deletion failed), 6 (rule waiting for configuration), 7 (rule waiting for deletion), 8 (rule waiting for certificate configuration)]
	StatusList []*uint64 `json:"StatusList,omitempty" name:"StatusList" list`
}

func (r *DescribleL7RulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribleL7RulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "RuleIdList")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Domain")
	delete(f, "ProtocolList")
	delete(f, "StatusList")
	if len(f) > 0 {
		return errors.New("DescribleL7RulesRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribleL7RulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Forwarding rule list
		Rules []*L7RuleEntry `json:"Rules,omitempty" name:"Rules" list`

		// Total number of rules
		Total *uint64 `json:"Total,omitempty" name:"Total"`

		// Health check configuration list
		Healths []*L7RuleHealth `json:"Healths,omitempty" name:"Healths" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribleL7RulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribleL7RulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribleRegionCountRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS service type. `bgpip`: Anti-DDoS Advanced, `bgp`: Anti-DDoS Pro (single IP), `bgp-multip`: Anti-DDoS Pro (multi-IP)
	Business *string `json:"Business,omitempty" name:"Business"`

	// Line search. Valid values: [1 (BGP line), 2 (Nanjing Telecom), 3 (Nanjing Unicom), 99 (third-party partner line)]. This field is valid only for Anti-DDoS Advanced and should be ignored for other product
	LineList []*uint64 `json:"LineList,omitempty" name:"LineList" list`
}

func (r *DescribleRegionCountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribleRegionCountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "LineList")
	if len(f) > 0 {
		return errors.New("DescribleRegionCountRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribleRegionCountResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of resource instances in region
		RegionList []*RegionInstanceCount `json:"RegionList,omitempty" name:"RegionList" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribleRegionCountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribleRegionCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type HttpStatusMap struct {

	// HTTP 2xx Status code
	Http2xx []*float64 `json:"Http2xx,omitempty" name:"Http2xx" list`

	// HTTP 3xx Status code
	Http3xx []*float64 `json:"Http3xx,omitempty" name:"Http3xx" list`

	// HTTP 404 Status code
	Http404 []*float64 `json:"Http404,omitempty" name:"Http404" list`

	// HTTP 4xx Status code
	Http4xx []*float64 `json:"Http4xx,omitempty" name:"Http4xx" list`

	// HTTP 5xx Status code
	Http5xx []*float64 `json:"Http5xx,omitempty" name:"Http5xx" list`

	// HTTP 2xx Forwarding status code
	SourceHttp2xx []*float64 `json:"SourceHttp2xx,omitempty" name:"SourceHttp2xx" list`

	// HTTP 3xx Forwarding status code
	SourceHttp3xx []*float64 `json:"SourceHttp3xx,omitempty" name:"SourceHttp3xx" list`

	// HTTP 404 Forwarding status code
	SourceHttp404 []*float64 `json:"SourceHttp404,omitempty" name:"SourceHttp404" list`

	// HTTP 4xx Forwarding status code
	SourceHttp4xx []*float64 `json:"SourceHttp4xx,omitempty" name:"SourceHttp4xx" list`

	// HTTP 5xx Forwarding status code
	SourceHttp5xx []*float64 `json:"SourceHttp5xx,omitempty" name:"SourceHttp5xx" list`
}

type IpBlackWhite struct {

	// IP address
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// Blocklist/allowlist type. Valid values: [black, white]
	Type *string `json:"Type,omitempty" name:"Type"`
}

type IpBlockData struct {

	// IP
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// Status (Blocked: blocked, UnBlocking: unblocking, UnBlockFailed: unblocking failed)
	Status *string `json:"Status,omitempty" name:"Status"`

	// Blocked time
	BlockTime *string `json:"BlockTime,omitempty" name:"BlockTime"`

	// Unblocked time (estimated)
	UnBlockTime *string `json:"UnBlockTime,omitempty" name:"UnBlockTime"`

	// Type of the unblocking action (`user`: self-service unblocking, `auto`: automatic unblocking, `update`: unblocking by service upgrading, `bind`: unblocking by binding Anti-DDoS Pro instance)
	ActionType *string `json:"ActionType,omitempty" name:"ActionType"`
}

type IpUnBlockData struct {

	// IP
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// Blocked time
	BlockTime *string `json:"BlockTime,omitempty" name:"BlockTime"`

	// Unblocked time (actual)
	UnBlockTime *string `json:"UnBlockTime,omitempty" name:"UnBlockTime"`

	// Type of the unblocking action (`user`: self-service unblocking, `auto`: automatic unblocking, `update`: unblocking by service upgrading, `bind`: unblocking by binding Anti-DDoS Pro instance)
	ActionType *string `json:"ActionType,omitempty" name:"ActionType"`
}

type KeyValue struct {

	// Field name
	Key *string `json:"Key,omitempty" name:"Key"`

	// Field value
	Value *string `json:"Value,omitempty" name:"Value"`
}

type KeyValueRecord struct {

	// Key-Value array of a record
	Record []*KeyValue `json:"Record,omitempty" name:"Record" list`
}

type L4HealthConfig struct {

	// Forwarding protocol. Valid values: [TCP, UDP]
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Forwarding port
	VirtualPort *uint64 `json:"VirtualPort,omitempty" name:"VirtualPort"`

	// 1: enabled, 0: disabled
	Enable *uint64 `json:"Enable,omitempty" name:"Enable"`

	// Response timeout period in seconds
	TimeOut *uint64 `json:"TimeOut,omitempty" name:"TimeOut"`

	// Detection interval in seconds
	Interval *uint64 `json:"Interval,omitempty" name:"Interval"`

	// Unhealthy threshold in times.
	KickNum *uint64 `json:"KickNum,omitempty" name:"KickNum"`

	// Healthy threshold in times.
	AliveNum *uint64 `json:"AliveNum,omitempty" name:"AliveNum"`

	// Session persistence duration in seconds
	KeepTime *uint64 `json:"KeepTime,omitempty" name:"KeepTime"`
}

type L4RuleEntry struct {

	// Forwarding protocol. Valid values: [TCP, UDP]
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Forwarding port
	VirtualPort *uint64 `json:"VirtualPort,omitempty" name:"VirtualPort"`

	// Real server port
	SourcePort *uint64 `json:"SourcePort,omitempty" name:"SourcePort"`

	// Forwarding method. Valid values: [1 (forwarding via domain name), 2 (forwarding via IP)]
	SourceType *uint64 `json:"SourceType,omitempty" name:"SourceType"`

	// Session persistence duration in seconds
	KeepTime *uint64 `json:"KeepTime,omitempty" name:"KeepTime"`

	// Forward list
	SourceList []*L4RuleSource `json:"SourceList,omitempty" name:"SourceList" list`

	// Load balancing method. Valid values: [1 (weighted round robin), 2 (source IP hash)]
	LbType *uint64 `json:"LbType,omitempty" name:"LbType"`

	// Session persistence status. Valid values: [0 (disabled), 1 (enabled)];
	KeepEnable *uint64 `json:"KeepEnable,omitempty" name:"KeepEnable"`

	// Rule ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// Rule description
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// Watermark removal status. Valid values: [0 (disabled), 1 (enabled)]
	RemoveSwitch *uint64 `json:"RemoveSwitch,omitempty" name:"RemoveSwitch"`
}

type L4RuleHealth struct {

	// Rule ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 1: enabled, 0: disabled
	Enable *uint64 `json:"Enable,omitempty" name:"Enable"`

	// Response timeout period in seconds
	TimeOut *uint64 `json:"TimeOut,omitempty" name:"TimeOut"`

	// Detection interval in seconds, which must be greater than the response timeout period
	Interval *uint64 `json:"Interval,omitempty" name:"Interval"`

	// Unhealthy threshold in times
	KickNum *uint64 `json:"KickNum,omitempty" name:"KickNum"`

	// Healthy threshold in times.
	AliveNum *uint64 `json:"AliveNum,omitempty" name:"AliveNum"`
}

type L4RuleSource struct {

	// Intermediate IP or domain name
	Source *string `json:"Source,omitempty" name:"Source"`

	// Weight value. Value range: [0,100]
	Weight *uint64 `json:"Weight,omitempty" name:"Weight"`
}

type L7HealthConfig struct {

	// Forwarding protocol. Valid values: [http, https, http/https]
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Forwarding domain name
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 1: enabled, 0: disabled
	Enable *uint64 `json:"Enable,omitempty" name:"Enable"`

	// Detection interval in seconds
	Interval *uint64 `json:"Interval,omitempty" name:"Interval"`

	// Number of exceptions in times
	KickNum *uint64 `json:"KickNum,omitempty" name:"KickNum"`

	// Number of health checks in times
	AliveNum *uint64 `json:"AliveNum,omitempty" name:"AliveNum"`

	// Health check detection method. Valid values: HEAD, GET. Default VALUE: HEAD
	Method *string `json:"Method,omitempty" name:"Method"`

	// Healthy status code during health check. xx = 1, 2xx = 2, 3xx = 4, 4xx = 8, 5xx = 16. Multiple status code values are added up
	StatusCode *uint64 `json:"StatusCode,omitempty" name:"StatusCode"`

	// URL of checked directory. Default value: /
	Url *string `json:"Url,omitempty" name:"Url"`
}

type L7RuleEntry struct {

	// Forwarding protocol. Valid values: [http, https]
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Forwarding domain name
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Forwarding method. Valid values: [1 (forwarding via domain name), 2 (forwarding via IP)]
	SourceType *uint64 `json:"SourceType,omitempty" name:"SourceType"`

	// Session persistence duration in seconds
	KeepTime *uint64 `json:"KeepTime,omitempty" name:"KeepTime"`

	// Forward list
	SourceList []*L4RuleSource `json:"SourceList,omitempty" name:"SourceList" list`

	// Load balancing method. Valid value: [1 (weighted round robin)]
	LbType *uint64 `json:"LbType,omitempty" name:"LbType"`

	// Session persistence status. Valid values: [0 (disabled), 1 (enabled)]
	KeepEnable *uint64 `json:"KeepEnable,omitempty" name:"KeepEnable"`

	// Rule ID, which is optional when adding a new rule but required when modifying or deleting a rule;
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// Certificate source, which is required if the forwarding protocol is HTTPS. Valid value: [2 (Tencent Cloud-hosted certificate)]. If the forwarding protocol is HTTP, 0 can be entered for this field
	CertType *uint64 `json:"CertType,omitempty" name:"CertType"`

	// If the certificate is a Tencent Cloud-hosted certificate, this field must be entered with the hosted certificate ID
	SSLId *string `json:"SSLId,omitempty" name:"SSLId"`

	// If the certificate is an external certificate, this field must be entered with the certificate content. (As external certificates are no longer supported, this field has been disused and does not need to be entered)
	Cert *string `json:"Cert,omitempty" name:"Cert"`

	// If the certificate is an external certificate, this field must be entered with the certificate key. (As external certificates are no longer supported, this field has been disused and does not need to be entered)
	PrivateKey *string `json:"PrivateKey,omitempty" name:"PrivateKey"`

	// Rule description
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// Rule status. Valid values: [0 (successfully configured rule), 1 (rule configuration taking effect), 2 (rule configuration failed), 3 (rule deletion taking effect), 5 (rule deletion failed), 6 (rule waiting for configuration), 7 (rule waiting for deletion), 8 (rule waiting for certificate configuration)]
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// CC protection status. Valid values: [0 (disabled), 1 (enabled)]
	CCStatus *uint64 `json:"CCStatus,omitempty" name:"CCStatus"`

	// HTTPS CC protection status. Valid values: [0 (disabled), 1 (enabled)]
	CCEnable *uint64 `json:"CCEnable,omitempty" name:"CCEnable"`

	// HTTPS CC protection threshold
	CCThreshold *uint64 `json:"CCThreshold,omitempty" name:"CCThreshold"`

	// HTTPS CC protection level
	CCLevel *string `json:"CCLevel,omitempty" name:"CCLevel"`

	// Whether to enable **Forward HTTPS requests via HTTP**. Valid values: `0` (disabled) and `1` (enabled). The default value is disabled.
	HttpsToHttpEnable *uint64 `json:"HttpsToHttpEnable,omitempty" name:"HttpsToHttpEnable"`

	// Access port number.
	// Note: this field may return null, indicating that no valid values can be obtained.
	VirtualPort *uint64 `json:"VirtualPort,omitempty" name:"VirtualPort"`
}

type L7RuleHealth struct {

	// Rule ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 1: enabled, 0: disabled
	Enable *uint64 `json:"Enable,omitempty" name:"Enable"`

	// Detection interval in seconds
	Interval *uint64 `json:"Interval,omitempty" name:"Interval"`

	// Unhealthy threshold in times.
	KickNum *uint64 `json:"KickNum,omitempty" name:"KickNum"`

	// Healthy threshold in times.
	AliveNum *uint64 `json:"AliveNum,omitempty" name:"AliveNum"`

	// HTTP request method. Valid values: [HEAD, GET]
	Method *string `json:"Method,omitempty" name:"Method"`

	// Healthy status code during health check. xx = 1, 2xx = 2, 3xx = 4, 4xx = 8, 5xx = 16. Multiple status code values are added up
	StatusCode *uint64 `json:"StatusCode,omitempty" name:"StatusCode"`

	// URL of checked directory. Default value: /
	Url *string `json:"Url,omitempty" name:"Url"`

	// Configuration status. 0: normal, 1: configuring, 2: configuration failed
	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

type ModifyCCAlarmThresholdRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS service type (`shield`: Chess Shield, `bgpip`: Anti-DDoS Advanced, `bgp`: Anti-DDoS Pro (single IP), `bgp-multip`: Anti-DDoS Pro (multi-IP), `net`: Anti-DDoS Ultimate)
	Business *string `json:"Business,omitempty" name:"Business"`

	// Anti-DDoS instance ID
	RsId *string `json:"RsId,omitempty" name:"RsId"`

	// Alarm threshold, which should be greater than 0 (currently scheduled value). It is set to 1000 on the backend by default
	AlarmThreshold *uint64 `json:"AlarmThreshold,omitempty" name:"AlarmThreshold"`

	// List of IPs associated with resource. If no Anti-DDoS Pro instance is bound, pass in an empty array. For Anti-DDoS Ultimate, pass in multiple IPs
	IpList []*string `json:"IpList,omitempty" name:"IpList" list`
}

func (r *ModifyCCAlarmThresholdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCCAlarmThresholdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "RsId")
	delete(f, "AlarmThreshold")
	delete(f, "IpList")
	if len(f) > 0 {
		return errors.New("ModifyCCAlarmThresholdRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCCAlarmThresholdResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Success code
		Success *SuccessCode `json:"Success,omitempty" name:"Success"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyCCAlarmThresholdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCCAlarmThresholdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCCFrequencyRulesRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS service type. `bgpip`: Anti-DDoS Advanced; `net`: Anti-DDoS Ultimate
	Business *string `json:"Business,omitempty" name:"Business"`

	// CC access frequency control rule ID
	CCFrequencyRuleId *string `json:"CCFrequencyRuleId,omitempty" name:"CCFrequencyRuleId"`

	// Matching rule. Valid values: ["include" (prefix match), "equal" (exact match)]
	Mode *string `json:"Mode,omitempty" name:"Mode"`

	// Reference period in seconds. Valid values: [10, 30, 60]
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// Number of access requests. Value range: [1-10000]
	ReqNumber *uint64 `json:"ReqNumber,omitempty" name:"ReqNumber"`

	// Action take. Valid values: ["alg" (CAPTCHA), "drop" (blocking)]
	Act *string `json:"Act,omitempty" name:"Act"`

	// Execution duration in seconds. Valid range: [1-900]
	ExeDuration *uint64 `json:"ExeDuration,omitempty" name:"ExeDuration"`

	// URI string, which must start with `/`, such as `/abc/a.php`. Length limit: 31. If URI is `/`, only prefix match can be selected as the matching mode;
	Uri *string `json:"Uri,omitempty" name:"Uri"`

	// `User-Agent` string. Length limit: 80
	UserAgent *string `json:"UserAgent,omitempty" name:"UserAgent"`

	// Cookie string. Length limit: 40
	Cookie *string `json:"Cookie,omitempty" name:"Cookie"`
}

func (r *ModifyCCFrequencyRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCCFrequencyRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "CCFrequencyRuleId")
	delete(f, "Mode")
	delete(f, "Period")
	delete(f, "ReqNumber")
	delete(f, "Act")
	delete(f, "ExeDuration")
	delete(f, "Uri")
	delete(f, "UserAgent")
	delete(f, "Cookie")
	if len(f) > 0 {
		return errors.New("ModifyCCFrequencyRulesRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCCFrequencyRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Success code
		Success *SuccessCode `json:"Success,omitempty" name:"Success"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyCCFrequencyRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCCFrequencyRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCCFrequencyRulesStatusRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS service type. `bgpip`: Anti-DDoS Advanced; `net`: Anti-DDoS Ultimate
	Business *string `json:"Business,omitempty" name:"Business"`

	// Anti-DDoS instance ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// Layer-7 forwarding rule ID, which can be obtained through the `DescribleL7Rules` API
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// Enables or disable. Valid values: ["on", "off"]
	Method *string `json:"Method,omitempty" name:"Method"`
}

func (r *ModifyCCFrequencyRulesStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCCFrequencyRulesStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "RuleId")
	delete(f, "Method")
	if len(f) > 0 {
		return errors.New("ModifyCCFrequencyRulesStatusRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCCFrequencyRulesStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Success code
		Success *SuccessCode `json:"Success,omitempty" name:"Success"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyCCFrequencyRulesStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCCFrequencyRulesStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCCHostProtectionRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS service type. `bgpip`: Anti-DDoS Advanced; `net`: Anti-DDoS Ultimate
	Business *string `json:"Business,omitempty" name:"Business"`

	// Anti-DDoS instance ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// Rule ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// Enables/disables CC domain name protection. Valid values: [open (enable), close (disable)]
	Method *string `json:"Method,omitempty" name:"Method"`
}

func (r *ModifyCCHostProtectionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCCHostProtectionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "RuleId")
	delete(f, "Method")
	if len(f) > 0 {
		return errors.New("ModifyCCHostProtectionRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCCHostProtectionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Success code
		Success *SuccessCode `json:"Success,omitempty" name:"Success"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyCCHostProtectionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCCHostProtectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCCIpAllowDenyRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS service type. `bgpip`: Anti-DDoS Advanced; `bgp`: Anti-DDoS Pro (single IP); `bgp-multip`: Anti-DDoS Pro (multi-IP), `net`: Anti-DDoS Ultimate
	Business *string `json:"Business,omitempty" name:"Business"`

	// Anti-DDoS instance ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// add: add, delete: delete
	Method *string `json:"Method,omitempty" name:"Method"`

	// Blocklist/allowlist type. Valid values: [white (allowlist), black (blocklist)]
	Type *string `json:"Type,omitempty" name:"Type"`

	// Blocklisted/whitelisted IP array
	IpList []*string `json:"IpList,omitempty" name:"IpList" list`

	// CC protection type, which is optional. Valid values: [http (HTTP CC protection), https (HTTPS CC protection)]; if this field is left empty, HTTPS CC protection will be used by default; if `https` is entered, the `Domain` and `RuleId` fields are required;
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// HTTPS layer-7 forwarding rule domain name (which is optional and can be obtained from the layer-7 forwarding rule API). This field is required only if `Protocol` is `https`;
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// HTTPS layer-7 forwarding rule ID (which is optional and can be obtained from the layer-7 forwarding rule API),
	// If `Method` is `delete`, this field can be left empty;
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
}

func (r *ModifyCCIpAllowDenyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCCIpAllowDenyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "Method")
	delete(f, "Type")
	delete(f, "IpList")
	delete(f, "Protocol")
	delete(f, "Domain")
	delete(f, "RuleId")
	if len(f) > 0 {
		return errors.New("ModifyCCIpAllowDenyRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCCIpAllowDenyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Success code
		Success *SuccessCode `json:"Success,omitempty" name:"Success"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyCCIpAllowDenyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCCIpAllowDenyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCCLevelRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS service type. `bgpip`: Anti-DDoS Advanced; `net`: Anti-DDoS Ultimate
	Business *string `json:"Business,omitempty" name:"Business"`

	// Anti-DDoS instance ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// CC protection level. Valid values: [default (normal), loose (loose), strict (strict)];
	Level *string `json:"Level,omitempty" name:"Level"`

	// CC protection type, which is optional. Valid values: [http (HTTP CC protection), https (HTTPS CC protection)]; if this field is left empty, HTTPS CC protection will be used by default; if `https` is entered, the `RuleId` field is required;
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Layer-7 forwarding rule ID (which can be obtained from the layer-7 forwarding rule API);
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
}

func (r *ModifyCCLevelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCCLevelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "Level")
	delete(f, "Protocol")
	delete(f, "RuleId")
	if len(f) > 0 {
		return errors.New("ModifyCCLevelRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCCLevelResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Success code
		Success *SuccessCode `json:"Success,omitempty" name:"Success"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyCCLevelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCCLevelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCCPolicySwitchRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS service type. `bgpip`: Anti-DDoS Advanced; `bgp`: Anti-DDoS Pro (single IP); `bgp-multip`: Anti-DDoS Pro (multi-IP), `net`: Anti-DDoS Ultimate
	Business *string `json:"Business,omitempty" name:"Business"`

	// Anti-DDoS instance ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// Policy ID
	SetId *string `json:"SetId,omitempty" name:"SetId"`

	// Status
	Switch *uint64 `json:"Switch,omitempty" name:"Switch"`
}

func (r *ModifyCCPolicySwitchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCCPolicySwitchRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "SetId")
	delete(f, "Switch")
	if len(f) > 0 {
		return errors.New("ModifyCCPolicySwitchRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCCPolicySwitchResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Success code
		Success *SuccessCode `json:"Success,omitempty" name:"Success"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyCCPolicySwitchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCCPolicySwitchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCCSelfDefinePolicyRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS service type. `bgpip`: Anti-DDoS Advanced; `bgp`: Anti-DDoS Pro (single IP); `bgp-multip`: Anti-DDoS Pro (multi-IP), `net`: Anti-DDoS Ultimate
	Business *string `json:"Business,omitempty" name:"Business"`

	// Anti-DDoS instance ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// Policy ID
	SetId *string `json:"SetId,omitempty" name:"SetId"`

	// Details of the CC protection policy
	Policy *CCPolicy `json:"Policy,omitempty" name:"Policy"`
}

func (r *ModifyCCSelfDefinePolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCCSelfDefinePolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "SetId")
	delete(f, "Policy")
	if len(f) > 0 {
		return errors.New("ModifyCCSelfDefinePolicyRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCCSelfDefinePolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Success code
		Success *SuccessCode `json:"Success,omitempty" name:"Success"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyCCSelfDefinePolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCCSelfDefinePolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCCThresholdRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS service type. `bgpip`: Anti-DDoS Advanced, `bgp`: Anti-DDoS Pro (single IP), `bgp-multip`: Anti-DDoS Pro (multi-IP), `net`: Anti-DDoS Ultimate, `basic`: Anti-DDoS Basic
	Business *string `json:"Business,omitempty" name:"Business"`

	// CC protection threshold. Valid values: (0 100 150 240 350 480 550 700 850 1000 1500 2000 3000 5000 10000 20000);
	// If `Business` is Anti-DDoS Advanced or Anti-DDoS Ultimate, its maximum CC protection threshold is subject to the base protection bandwidth in the following way:
	//   Base bandwidth: maximum CC protection threshold
	//   10:  20000,
	//   20:  40000,
	//   30:  70000,
	//   40:  100000,
	//   50:  150000,
	//   60:  200000,
	//   80:  250000,
	//   100: 300000,
	Threshold *uint64 `json:"Threshold,omitempty" name:"Threshold"`

	// Anti-DDoS instance ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// CC protection type, which is optional. Valid values: [http (HTTP CC protection), https (HTTPS CC protection)]; if this field is left empty, HTTPS CC protection will be used by default; if `https` is entered, the `RuleId` field is required;
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// HTTPS layer-7 forwarding rule ID (which is optional and can be obtained from the layer-7 forwarding rule API);
	// Required if `Protocol` is `https`;
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// Queried IP address (only provided by Anti-DDoS Basic), such as 1.1.1.1
	BasicIp *string `json:"BasicIp,omitempty" name:"BasicIp"`

	// IP region (only provided for Anti-DDoS Basic). Valid values: region abbreviations such as gz, bj, sh, and hk
	BasicRegion *string `json:"BasicRegion,omitempty" name:"BasicRegion"`

	// Zone type (only provided for Anti-DDoS Basic). Valid values: public (public cloud zone), bm (BM zone), nat (NAT server zone), channel (internet channel).
	BasicBizType *string `json:"BasicBizType,omitempty" name:"BasicBizType"`

	// Device type (only provided for Anti-DDoS Basic). Valid values: cvm (CVM), clb (public CLB), lb (BM CLB), nat (NAT server), channel (internet channel).
	BasicDeviceType *string `json:"BasicDeviceType,omitempty" name:"BasicDeviceType"`

	// IPInstance Nat gateway (only provided for Anti-DDoS Basic), which is optional. (If the device type to be queried is a NAT server, this parameter is required, which can be obtained through the NAT resource query API)
	BasicIpInstance *string `json:"BasicIpInstance,omitempty" name:"BasicIpInstance"`

	// ISP line (only provided for Anti-DDoS Basic), which is optional. (If the device type to be queried is a NAT server, this parameter should be 5)
	BasicIspCode *uint64 `json:"BasicIspCode,omitempty" name:"BasicIspCode"`

	// This optional field must be specified when HTTPS protocol is used.
	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *ModifyCCThresholdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCCThresholdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Threshold")
	delete(f, "Id")
	delete(f, "Protocol")
	delete(f, "RuleId")
	delete(f, "BasicIp")
	delete(f, "BasicRegion")
	delete(f, "BasicBizType")
	delete(f, "BasicDeviceType")
	delete(f, "BasicIpInstance")
	delete(f, "BasicIspCode")
	delete(f, "Domain")
	if len(f) > 0 {
		return errors.New("ModifyCCThresholdRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCCThresholdResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Success code
		Success *SuccessCode `json:"Success,omitempty" name:"Success"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyCCThresholdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCCThresholdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCCUrlAllowRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS service type. `bgpip`: Anti-DDoS Advanced; `bgp`: Anti-DDoS Pro (single IP); `bgp-multip`: Anti-DDoS Pro (multi-IP), `net`: Anti-DDoS Ultimate
	Business *string `json:"Business,omitempty" name:"Business"`

	// Anti-DDoS instance ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// add: add, delete: delete
	Method *string `json:"Method,omitempty" name:"Method"`

	// Blocklist/allowlist type. Valid value: [white (allowlist)]
	Type *string `json:"Type,omitempty" name:"Type"`

	// URL array. URL format:
	// http://domain name/cgi
	// https://domain name/cgi
	UrlList []*string `json:"UrlList,omitempty" name:"UrlList" list`

	// CC protection type, which is optional. Valid values: [http (HTTP CC protection), https (HTTPS CC protection)]; if this field is left empty, HTTPS CC protection will be used by default; if `https` is entered, the `Domain` and `RuleId` fields are required;
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// HTTPS layer-7 forwarding rule domain name (which is optional and can be obtained from the layer-7 forwarding rule API). This field is required only if `Protocol` is `https`;
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// HTTPS layer-7 forwarding rule ID (which is optional and can be obtained from the layer-7 forwarding rule API). This field is required only when adding a rule and `Protocol` is `https`;
	// If `Method` is `delete`, this field can be left empty;
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
}

func (r *ModifyCCUrlAllowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCCUrlAllowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "Method")
	delete(f, "Type")
	delete(f, "UrlList")
	delete(f, "Protocol")
	delete(f, "Domain")
	delete(f, "RuleId")
	if len(f) > 0 {
		return errors.New("ModifyCCUrlAllowRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCCUrlAllowResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Success code
		Success *SuccessCode `json:"Success,omitempty" name:"Success"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyCCUrlAllowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCCUrlAllowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDDoSAIStatusRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS service type. `bgpip`: Anti-DDoS Advanced; `bgp`: Anti-DDoS Pro (single IP); `bgp-multip`: Anti-DDoS Pro (multi-IP), `net`: Anti-DDoS Ultimate
	Business *string `json:"Business,omitempty" name:"Business"`

	// Anti-DDoS instance ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// get (read AI protection status), set (modify AI protection status);
	Method *string `json:"Method,omitempty" name:"Method"`

	// AI protection status, which is required if `Method` is `set`. Valid values: [on, off].
	DDoSAI *string `json:"DDoSAI,omitempty" name:"DDoSAI"`
}

func (r *ModifyDDoSAIStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDDoSAIStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "Method")
	delete(f, "DDoSAI")
	if len(f) > 0 {
		return errors.New("ModifyDDoSAIStatusRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDDoSAIStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// AI protection status. Valid values: [on, off]
		DDoSAI *string `json:"DDoSAI,omitempty" name:"DDoSAI"`

		// Anti-DDoS instance ID
		Id *string `json:"Id,omitempty" name:"Id"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDDoSAIStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDDoSAIStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDDoSAlarmThresholdRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS service type (`shield`: Chess Shield, `bgpip`: Anti-DDoS Advanced, `bgp`: Anti-DDoS Pro (single IP), `bgp-multip`: Anti-DDoS Pro (multi-IP), `net`: Anti-DDoS Ultimate)
	Business *string `json:"Business,omitempty" name:"Business"`

	// Anti-DDoS instance ID
	RsId *string `json:"RsId,omitempty" name:"RsId"`

	// Alarm threshold type. 0: not set, 1: inbound traffic, 2: cleansed traffic
	AlarmType *uint64 `json:"AlarmType,omitempty" name:"AlarmType"`

	// Alarm threshold, which should be greater than 0 (currently scheduled value)
	AlarmThreshold *uint64 `json:"AlarmThreshold,omitempty" name:"AlarmThreshold"`

	// List of IPs associated with resource. If no Anti-DDoS Pro instance is bound, pass in an empty array. For Anti-DDoS Ultimate, pass in multiple IPs
	IpList []*string `json:"IpList,omitempty" name:"IpList" list`
}

func (r *ModifyDDoSAlarmThresholdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDDoSAlarmThresholdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "RsId")
	delete(f, "AlarmType")
	delete(f, "AlarmThreshold")
	delete(f, "IpList")
	if len(f) > 0 {
		return errors.New("ModifyDDoSAlarmThresholdRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDDoSAlarmThresholdResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Success code
		Success *SuccessCode `json:"Success,omitempty" name:"Success"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDDoSAlarmThresholdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDDoSAlarmThresholdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDDoSDefendStatusRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS service type. `bgp`: Anti-DDoS Pro (single IP), `bgp-multip`: Anti-DDoS Pro (multi-IP), `bgpip`: Anti-DDoS Advanced, `net`: Anti-DDoS Ultimate, `basic`: Anti-DDoS Basic
	Business *string `json:"Business,omitempty" name:"Business"`

	// Protection status. Valid values: [0 (disabled), 1 (enabled)]
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// Disablement duration in hours. Valid values: [0, 1, 2, 3, 4, 5, 6]. If `Status` is `0` (indicating to disable), `Hour` must be greater than 0;
	Hour *int64 `json:"Hour,omitempty" name:"Hour"`

	// Resource ID, which is required if `Business` is not Anti-DDoS Basic;
	Id *string `json:"Id,omitempty" name:"Id"`

	// Anti-DDoS Basic IP, which is required only if `Business` is Anti-DDoS Basic;
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// Product type of IP, which is required only if `Business` is Anti-DDoS Basic. Valid values: [public (CVM), bm (BM), eni (ENI), vpngw (VPN Gateway), natgw (NAT Gateway), waf (WAF), fpc (finance product), gaap (GAAP), other (hosted IP)]
	BizType *string `json:"BizType,omitempty" name:"BizType"`

	// Product subtype of IP, which is required only if `Business` is Anti-DDoS Basic. Valid values: [cvm (CVM), lb (CLB), eni (ENI), vpngw (VPN), natgw (NAT), waf (WAF), fpc (finance), gaap (GAAP), other (hosted IP), eip (BM EIP)]
	DeviceType *string `json:"DeviceType,omitempty" name:"DeviceType"`

	// Resource instance ID of IP, which is required only if `Business` is Anti-DDoS Basic. This field is required when binding a new IP. For example, if it is an ENI IP, enter `ID(eni-*)` of the ENI for `InstanceId`;
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// This field is required only if `Business` is Anti-DDoS Basic, indicating the IP region. Valid values:
	// "bj":     North China (Beijing)
	// "cd":     Southwest China (Chengdu)
	// "cq":     Southwest China (Chongqing)
	// "gz":     South China (Guangzhou)
	// "gzopen": South China (Guangzhou Open)
	// "hk":     Hong Kong (China)
	// "kr":     Southeast Asia (Seoul)
	// "sh":     East China (Shanghai)
	// "shjr":   East China (Shanghai Finance)
	// "szjr":   South China (Shenzhen Finance)
	// "sg":     Southeast Asia (Singapore)
	// "th":     Southeast Asia (Thailand)
	// "de":     Europe (Germany)
	// "usw":    West US (Silicon Valley)
	// "ca":     North America (Toronto)
	// "jp":     Japan
	// "hzec":   Hangzhou
	// "in":     India
	// "use":    East US (Virginia)
	// "ru":     Russia
	// "tpe":    Taiwan (China)
	// "nj":     Nanjing
	IPRegion *string `json:"IPRegion,omitempty" name:"IPRegion"`
}

func (r *ModifyDDoSDefendStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDDoSDefendStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Status")
	delete(f, "Hour")
	delete(f, "Id")
	delete(f, "Ip")
	delete(f, "BizType")
	delete(f, "DeviceType")
	delete(f, "InstanceId")
	delete(f, "IPRegion")
	if len(f) > 0 {
		return errors.New("ModifyDDoSDefendStatusRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDDoSDefendStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Success code
		Success *SuccessCode `json:"Success,omitempty" name:"Success"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDDoSDefendStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDDoSDefendStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDDoSLevelRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS service type. `bgpip`: Anti-DDoS Advanced; `bgp`: Anti-DDoS Pro (single IP); `bgp-multip`: Anti-DDoS Pro (multi-IP), `net`: Anti-DDoS Ultimate
	Business *string `json:"Business,omitempty" name:"Business"`

	// Anti-DDoS instance ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// get (read protection level), set (modify protection level);
	Method *string `json:"Method,omitempty" name:"Method"`

	// Protection level, which is required if `Method` is `set`. Valid values: [low,middle,high]
	DDoSLevel *string `json:"DDoSLevel,omitempty" name:"DDoSLevel"`
}

func (r *ModifyDDoSLevelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDDoSLevelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "Method")
	delete(f, "DDoSLevel")
	if len(f) > 0 {
		return errors.New("ModifyDDoSLevelRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDDoSLevelResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Anti-DDoS instance ID
		Id *string `json:"Id,omitempty" name:"Id"`

		// Protection level. Valid values: [low,middle,high]
		DDoSLevel *string `json:"DDoSLevel,omitempty" name:"DDoSLevel"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDDoSLevelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDDoSLevelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDDoSPolicyCaseRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS service type. `bgpip`: Anti-DDoS Advanced; `bgp`: Anti-DDoS Pro (single IP); `bgp-multip`: Anti-DDoS Pro (multi-IP), `net`: Anti-DDoS Ultimate
	Business *string `json:"Business,omitempty" name:"Business"`

	// Policy scenario ID
	SceneId *string `json:"SceneId,omitempty" name:"SceneId"`

	// Development platform. Valid values: [PC (PC client), MOBILE (mobile device), TV (TV), SERVER (server)]
	PlatformTypes []*string `json:"PlatformTypes,omitempty" name:"PlatformTypes" list`

	// Category. Valid values: [WEB (website), GAME (game), APP (application), OTHER (other)]
	AppType *string `json:"AppType,omitempty" name:"AppType"`

	// Application protocol. Valid values: [tcp (TCP protocol), udp (UDP protocol), icmp (ICMP protocol), all (other protocols)]
	AppProtocols []*string `json:"AppProtocols,omitempty" name:"AppProtocols" list`

	// TCP start port. Value range: (0, 65535]
	TcpSportStart *string `json:"TcpSportStart,omitempty" name:"TcpSportStart"`

	// TCP end port. Value range: (0, 65535). It must be greater than or equal to the TCP start port
	TcpSportEnd *string `json:"TcpSportEnd,omitempty" name:"TcpSportEnd"`

	// UDP start port. Value range: (0, 65535]
	UdpSportStart *string `json:"UdpSportStart,omitempty" name:"UdpSportStart"`

	// End UDP business port. Value range: (0, 65535). It must be greater than or equal to the start UDP business port
	UdpSportEnd *string `json:"UdpSportEnd,omitempty" name:"UdpSportEnd"`

	// Whether there are customers outside Mainland China. Valid values: [no, yes]
	HasAbroad *string `json:"HasAbroad,omitempty" name:"HasAbroad"`

	// Whether to actively initiate outbound TCP requests. Valid values: [no, yes]
	HasInitiateTcp *string `json:"HasInitiateTcp,omitempty" name:"HasInitiateTcp"`

	// Whether to actively initiate outbound UDP requests. Valid values: [no, yes]
	HasInitiateUdp *string `json:"HasInitiateUdp,omitempty" name:"HasInitiateUdp"`

	// Port that actively initiates TCP requests. Value range: (0, 65535]
	PeerTcpPort *string `json:"PeerTcpPort,omitempty" name:"PeerTcpPort"`

	// Port that actively initiates UDP requests. Value range: (0, 65535]
	PeerUdpPort *string `json:"PeerUdpPort,omitempty" name:"PeerUdpPort"`

	// Fixed feature code of TCP payload. String length limit: 512
	TcpFootprint *string `json:"TcpFootprint,omitempty" name:"TcpFootprint"`

	// Fixed feature code of UDP payload. String length limit: 512
	UdpFootprint *string `json:"UdpFootprint,omitempty" name:"UdpFootprint"`

	// Web business API URL
	WebApiUrl []*string `json:"WebApiUrl,omitempty" name:"WebApiUrl" list`

	// Minimum length of TCP business packet. Value range: (0, 1500)
	MinTcpPackageLen *string `json:"MinTcpPackageLen,omitempty" name:"MinTcpPackageLen"`

	// Maximum length of TCP business packet. Value range: (0, 1500). It must be greater than or equal to the minimum length of TCP business packet
	MaxTcpPackageLen *string `json:"MaxTcpPackageLen,omitempty" name:"MaxTcpPackageLen"`

	// Minimum length of UDP business packet. Value range: (0, 1500)
	MinUdpPackageLen *string `json:"MinUdpPackageLen,omitempty" name:"MinUdpPackageLen"`

	// Maximum length of UDP business packet. Value range: (0, 1500). It must be greater than or equal to the minimum length of UDP business packet
	MaxUdpPackageLen *string `json:"MaxUdpPackageLen,omitempty" name:"MaxUdpPackageLen"`

	// Whether there are VPN businesses. Valid values: [no, yes]
	HasVPN *string `json:"HasVPN,omitempty" name:"HasVPN"`

	// TCP business port list. Individual ports and port ranges are supported, which should be in string type, such as 80,443,700-800,53,1000-3000
	TcpPortList *string `json:"TcpPortList,omitempty" name:"TcpPortList"`

	// UDP business port list. Individual ports and port ranges are supported, which should be in string type, such as 80,443,700-800,53,1000-3000
	UdpPortList *string `json:"UdpPortList,omitempty" name:"UdpPortList"`
}

func (r *ModifyDDoSPolicyCaseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDDoSPolicyCaseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "SceneId")
	delete(f, "PlatformTypes")
	delete(f, "AppType")
	delete(f, "AppProtocols")
	delete(f, "TcpSportStart")
	delete(f, "TcpSportEnd")
	delete(f, "UdpSportStart")
	delete(f, "UdpSportEnd")
	delete(f, "HasAbroad")
	delete(f, "HasInitiateTcp")
	delete(f, "HasInitiateUdp")
	delete(f, "PeerTcpPort")
	delete(f, "PeerUdpPort")
	delete(f, "TcpFootprint")
	delete(f, "UdpFootprint")
	delete(f, "WebApiUrl")
	delete(f, "MinTcpPackageLen")
	delete(f, "MaxTcpPackageLen")
	delete(f, "MinUdpPackageLen")
	delete(f, "MaxUdpPackageLen")
	delete(f, "HasVPN")
	delete(f, "TcpPortList")
	delete(f, "UdpPortList")
	if len(f) > 0 {
		return errors.New("ModifyDDoSPolicyCaseRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDDoSPolicyCaseResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Success code
		Success *SuccessCode `json:"Success,omitempty" name:"Success"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDDoSPolicyCaseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDDoSPolicyCaseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDDoSPolicyNameRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS service type. `bgpip`: Anti-DDoS Advanced; `bgp`: Anti-DDoS Pro (single IP); `bgp-multip`: Anti-DDoS Pro (multi-IP), `net`: Anti-DDoS Ultimate
	Business *string `json:"Business,omitempty" name:"Business"`

	// Policy ID
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

	// Policy name
	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *ModifyDDoSPolicyNameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDDoSPolicyNameRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "PolicyId")
	delete(f, "Name")
	if len(f) > 0 {
		return errors.New("ModifyDDoSPolicyNameRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDDoSPolicyNameResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Success code
		Success *SuccessCode `json:"Success,omitempty" name:"Success"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDDoSPolicyNameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDDoSPolicyNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDDoSPolicyRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS service type. `bgpip`: Anti-DDoS Advanced; `bgp`: Anti-DDoS Pro (single IP); `bgp-multip`: Anti-DDoS Pro (multi-IP), `net`: Anti-DDoS Ultimate
	Business *string `json:"Business,omitempty" name:"Business"`

	// Policy ID
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

	// Protocol disablement, which must be entered, and the array length must be 1
	DropOptions []*DDoSPolicyDropOption `json:"DropOptions,omitempty" name:"DropOptions" list`

	// Port disablement. If no ports are to be disabled, enter an empty array
	PortLimits []*DDoSPolicyPortLimit `json:"PortLimits,omitempty" name:"PortLimits" list`

	// IP blocklist/allowlist. Enter an empty array if there is no IP blocklist/allowlist
	IpAllowDenys []*IpBlackWhite `json:"IpAllowDenys,omitempty" name:"IpAllowDenys" list`

	// Packet filter. Enter an empty array if there are no packets to filter
	PacketFilters []*DDoSPolicyPacketFilter `json:"PacketFilters,omitempty" name:"PacketFilters" list`

	// Watermarking policy parameter. Enter an empty array if the watermarking feature is not enabled. At most one watermarking policy can be passed in (that is, the size of the array cannot exceed 1)
	WaterPrint []*WaterPrintPolicy `json:"WaterPrint,omitempty" name:"WaterPrint" list`
}

func (r *ModifyDDoSPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDDoSPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "PolicyId")
	delete(f, "DropOptions")
	delete(f, "PortLimits")
	delete(f, "IpAllowDenys")
	delete(f, "PacketFilters")
	delete(f, "WaterPrint")
	if len(f) > 0 {
		return errors.New("ModifyDDoSPolicyRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDDoSPolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Success code
		Success *SuccessCode `json:"Success,omitempty" name:"Success"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDDoSPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDDoSPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDDoSSwitchRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS service type. `basic`: Anti-DDoS Basic
	Business *string `json:"Business,omitempty" name:"Business"`

	// `get`: read DDoS protection status, `set`: modify DDoS protection status
	Method *string `json:"Method,omitempty" name:"Method"`

	// Anti-DDoS Basic IP, which is required only if `Business` is Anti-DDoS Basic;
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// Product type of IP, which is required only if `Business` is Anti-DDoS Basic. Valid values: [public (CVM), bm (BM), eni (ENI), vpngw (VPN Gateway), natgw (NAT Gateway), waf (WAF), fpc (finance product), gaap (GAAP), other (hosted IP)]
	BizType *string `json:"BizType,omitempty" name:"BizType"`

	// Product subtype of IP, which is required only if `Business` is Anti-DDoS Basic. Valid values: [cvm (CVM), lb (CLB), eni (ENI), vpngw (VPN), natgw (NAT), waf (WAF), fpc (finance), gaap (GAAP), other (hosted IP), eip (BM EIP)]
	DeviceType *string `json:"DeviceType,omitempty" name:"DeviceType"`

	// Resource instance ID of IP, which is required only if `Business` is Anti-DDoS Basic. This field is required when binding a new IP. For example, if it is an ENI IP, enter `ID(eni-*)` of the ENI for `InstanceId`;
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// This field is required only if `Business` is Anti-DDoS Basic, indicating the IP region. Valid values:
	// "bj":     North China (Beijing)
	// "cd":     Southwest China (Chengdu)
	// "cq":     Southwest China (Chongqing)
	// "gz":     South China (Guangzhou)
	// "gzopen": South China (Guangzhou Open)
	// "hk":     Hong Kong (China)
	// "kr":     Southeast Asia (Seoul)
	// "sh":     East China (Shanghai)
	// "shjr":   East China (Shanghai Finance)
	// "szjr":   South China (Shenzhen Finance)
	// "sg":     Southeast Asia (Singapore)
	// "th":     Southeast Asia (Thailand)
	// "de":     Europe (Germany)
	// "usw":    West US (Silicon Valley)
	// "ca":     North America (Toronto)
	// "jp":     Japan
	// "hzec":   Hangzhou
	// "in":     India
	// "use":    East US (Virginia)
	// "ru":     Russia
	// "tpe":    Taiwan (China)
	// "nj":     Nanjing
	IPRegion *string `json:"IPRegion,omitempty" name:"IPRegion"`

	// Protection status value, which is optional. Valid values: [0 (disabled), 1 (enabled)]. If `Method` is `get`, this field can be left empty;
	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyDDoSSwitchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDDoSSwitchRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Method")
	delete(f, "Ip")
	delete(f, "BizType")
	delete(f, "DeviceType")
	delete(f, "InstanceId")
	delete(f, "IPRegion")
	delete(f, "Status")
	if len(f) > 0 {
		return errors.New("ModifyDDoSSwitchRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDDoSSwitchResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Current protection status value. Valid values: [0 (disabled), 1 (enabled)]
		Status *uint64 `json:"Status,omitempty" name:"Status"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDDoSSwitchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDDoSSwitchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDDoSThresholdRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS service type. `bgpip`: Anti-DDoS Advanced; `bgp`: Anti-DDoS Pro (single IP); `bgp-multip`: Anti-DDoS Pro (multi-IP), `net`: Anti-DDoS Ultimate
	Business *string `json:"Business,omitempty" name:"Business"`

	// Anti-DDoS instance ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// DDoS cleansing threshold. Valid values: [0, 60, 80, 100, 150, 200, 250, 300, 400, 500, 700, 1000];
	// If the value is set to 0, the default value will be used;
	Threshold *uint64 `json:"Threshold,omitempty" name:"Threshold"`
}

func (r *ModifyDDoSThresholdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDDoSThresholdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "Threshold")
	if len(f) > 0 {
		return errors.New("ModifyDDoSThresholdRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDDoSThresholdResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Success code
		Success *SuccessCode `json:"Success,omitempty" name:"Success"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDDoSThresholdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDDoSThresholdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDDoSWaterKeyRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS service type. `bgpip`: Anti-DDoS Advanced; `bgp`: Anti-DDoS Pro (single IP); `bgp-multip`: Anti-DDoS Pro (multi-IP), `net`: Anti-DDoS Ultimate
	Business *string `json:"Business,omitempty" name:"Business"`

	// Policy ID
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

	// Key operation. Valid values: [add (add), delete (delete), open (open), close (close), get (get key)]
	Method *string `json:"Method,omitempty" name:"Method"`

	// Key ID, which can be left empty or 0 when adding a key but is required for other operations
	KeyId *uint64 `json:"KeyId,omitempty" name:"KeyId"`
}

func (r *ModifyDDoSWaterKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDDoSWaterKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "PolicyId")
	delete(f, "Method")
	delete(f, "KeyId")
	if len(f) > 0 {
		return errors.New("ModifyDDoSWaterKeyRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDDoSWaterKeyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Watermark key list
		KeyList []*WaterPrintKey `json:"KeyList,omitempty" name:"KeyList" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDDoSWaterKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDDoSWaterKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyElasticLimitRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS service type. `bgpip`: Anti-DDoS Advanced; `bgp`: Anti-DDoS Pro (single IP); `bgp-multip`: Anti-DDoS Pro (multi-IP), `net`: Anti-DDoS Ultimate
	Business *string `json:"Business,omitempty" name:"Business"`

	// Anti-DDoS instance ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// Elastic protection threshold. Valid values: [0 10000 20000 30000 40000 50000 60000 70000 80000 90000 100000 120000 150000 200000 250000 300000 400000 600000 800000 220000 310000 110000 270000 610000]
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *ModifyElasticLimitRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyElasticLimitRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "Limit")
	if len(f) > 0 {
		return errors.New("ModifyElasticLimitRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyElasticLimitResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Success code
		Success *SuccessCode `json:"Success,omitempty" name:"Success"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyElasticLimitResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyElasticLimitResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyL4HealthRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS service type. `bgpip`: Anti-DDoS Advanced; `net`: Anti-DDoS Ultimate
	Business *string `json:"Business,omitempty" name:"Business"`

	// Anti-DDoS instance ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// Health check parameter array
	Healths []*L4RuleHealth `json:"Healths,omitempty" name:"Healths" list`
}

func (r *ModifyL4HealthRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyL4HealthRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "Healths")
	if len(f) > 0 {
		return errors.New("ModifyL4HealthRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyL4HealthResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Success code
		Success *SuccessCode `json:"Success,omitempty" name:"Success"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyL4HealthResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyL4HealthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyL4KeepTimeRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS service type. `bgpip`: Anti-DDoS Advanced; `net`: Anti-DDoS Ultimate
	Business *string `json:"Business,omitempty" name:"Business"`

	// Anti-DDoS instance ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// Rule ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// Session persistence status. Valid values: [0 (disabled), 1 (enabled)]
	KeepEnable *uint64 `json:"KeepEnable,omitempty" name:"KeepEnable"`

	// Session persistence duration in seconds
	KeepTime *uint64 `json:"KeepTime,omitempty" name:"KeepTime"`
}

func (r *ModifyL4KeepTimeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyL4KeepTimeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "RuleId")
	delete(f, "KeepEnable")
	delete(f, "KeepTime")
	if len(f) > 0 {
		return errors.New("ModifyL4KeepTimeRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyL4KeepTimeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Success code
		Success *SuccessCode `json:"Success,omitempty" name:"Success"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyL4KeepTimeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyL4KeepTimeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyL4RulesRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS service type. `bgpip`: Anti-DDoS Advanced; `bgp`: Anti-DDoS Pro (single IP); `bgp-multip`: Anti-DDoS Pro (multi-IP), `net`: Anti-DDoS Ultimate
	Business *string `json:"Business,omitempty" name:"Business"`

	// Anti-DDoS instance ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// Rule
	Rule *L4RuleEntry `json:"Rule,omitempty" name:"Rule"`
}

func (r *ModifyL4RulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyL4RulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "Rule")
	if len(f) > 0 {
		return errors.New("ModifyL4RulesRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyL4RulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Success code
		Success *SuccessCode `json:"Success,omitempty" name:"Success"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyL4RulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyL4RulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyL7RulesRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS service type. `bgpip`: Anti-DDoS Advanced; `bgp`: Anti-DDoS Pro (single IP); `bgp-multip`: Anti-DDoS Pro (multi-IP), `net`: Anti-DDoS Ultimate
	Business *string `json:"Business,omitempty" name:"Business"`

	// Anti-DDoS instance ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// Rule
	Rule *L7RuleEntry `json:"Rule,omitempty" name:"Rule"`
}

func (r *ModifyL7RulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyL7RulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "Rule")
	if len(f) > 0 {
		return errors.New("ModifyL7RulesRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyL7RulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Success code
		Success *SuccessCode `json:"Success,omitempty" name:"Success"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyL7RulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyL7RulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNetReturnSwitchRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS service type (`net`: Anti-DDoS Ultimate)
	Business *string `json:"Business,omitempty" name:"Business"`

	// Anti-DDoS instance ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// Switch status. 0: disabled, 1: enabled
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// Switch duration in hours. Valid values: [0,1,2,3,4,5,6;]. If `status` is 1, this field is required and must be greater than 0
	Hour *uint64 `json:"Hour,omitempty" name:"Hour"`
}

func (r *ModifyNetReturnSwitchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNetReturnSwitchRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "Status")
	delete(f, "Hour")
	if len(f) > 0 {
		return errors.New("ModifyNetReturnSwitchRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNetReturnSwitchResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyNetReturnSwitchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNetReturnSwitchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNewDomainRulesRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS service type (`bgpip`: Anti-DDoS Advanced).
	Business *string `json:"Business,omitempty" name:"Business"`

	// Anti-DDoS instance ID.
	Id *string `json:"Id,omitempty" name:"Id"`

	// Domain name forwarding rule.
	Rule *NewL7RuleEntry `json:"Rule,omitempty" name:"Rule"`
}

func (r *ModifyNewDomainRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNewDomainRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "Rule")
	if len(f) > 0 {
		return errors.New("ModifyNewDomainRulesRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNewDomainRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Success code.
		Success *SuccessCode `json:"Success,omitempty" name:"Success"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyNewDomainRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNewDomainRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNewL4RuleRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS service type (`bgpip`: Anti-DDoS Advanced).
	Business *string `json:"Business,omitempty" name:"Business"`

	// Anti-DDoS instance ID.
	Id *string `json:"Id,omitempty" name:"Id"`

	// Forwarding rule.
	Rule *L4RuleEntry `json:"Rule,omitempty" name:"Rule"`
}

func (r *ModifyNewL4RuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNewL4RuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "Rule")
	if len(f) > 0 {
		return errors.New("ModifyNewL4RuleRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNewL4RuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Success code.
		Success *SuccessCode `json:"Success,omitempty" name:"Success"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyNewL4RuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNewL4RuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyResBindDDoSPolicyRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS service type. `bgpip`: Anti-DDoS Advanced; `bgp`: Anti-DDoS Pro (single IP); `bgp-multip`: Anti-DDoS Pro (multi-IP), `net`: Anti-DDoS Ultimate
	Business *string `json:"Business,omitempty" name:"Business"`

	// Anti-DDoS instance ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// Policy ID
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

	// bind: bind policy; unbind: unbind policy
	Method *string `json:"Method,omitempty" name:"Method"`
}

func (r *ModifyResBindDDoSPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyResBindDDoSPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "PolicyId")
	delete(f, "Method")
	if len(f) > 0 {
		return errors.New("ModifyResBindDDoSPolicyRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyResBindDDoSPolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Success code
		Success *SuccessCode `json:"Success,omitempty" name:"Success"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyResBindDDoSPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyResBindDDoSPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyResourceRenewFlagRequest struct {
	*tchttp.BaseRequest

	// Anti-DDoS service type. `bgpip`: Anti-DDoS Advanced, `net`: Anti-DDoS Ultimate, `shield`: Chess Shield, `bgp`: Anti-DDoS Pro (single IP), `bgp-multip`: Anti-DDoS Pro (multi-IP), `insurance`: guarantee package, `staticpack`: non-BGP package
	Business *string `json:"Business,omitempty" name:"Business"`

	// Resource ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// Auto-renewal flag (0: manual renewal, 1: auto-renewal, 2: no renewal upon expiration)
	RenewFlag *uint64 `json:"RenewFlag,omitempty" name:"RenewFlag"`
}

func (r *ModifyResourceRenewFlagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyResourceRenewFlagRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "RenewFlag")
	if len(f) > 0 {
		return errors.New("ModifyResourceRenewFlagRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyResourceRenewFlagResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Success code
		Success *SuccessCode `json:"Success,omitempty" name:"Success"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyResourceRenewFlagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyResourceRenewFlagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NewL7RuleEntry struct {

	// Forwarding protocol. Valid values: `http` and `https`.
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Forwarding domain name.
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Forwarding method. Valid values: `1` (by domain name); `2` (by IP).
	SourceType *uint64 `json:"SourceType,omitempty" name:"SourceType"`

	// Session persistence duration, in seconds.
	KeepTime *uint64 `json:"KeepTime,omitempty" name:"KeepTime"`

	// List of sources
	SourceList []*L4RuleSource `json:"SourceList,omitempty" name:"SourceList" list`

	// Load balancing method. Valid value: `1` (weighed polling).
	LbType *uint64 `json:"LbType,omitempty" name:"LbType"`

	// Whether session persistence is enabled. Valid values: `0` (disabled) and `1` (enabled).
	KeepEnable *uint64 `json:"KeepEnable,omitempty" name:"KeepEnable"`

	// Rule ID. This field is not required for adding a rule, but is required for modifying or deleting a rule.
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// Certificate source. When the forwarding protocol is HTTPS, this field must be set to `2` (Tencent Cloud managed certificate), and for HTTP protocol, it can be set to `0`.
	CertType *uint64 `json:"CertType,omitempty" name:"CertType"`

	// When the certificate source is Tencent Cloud managed certificate, this field must be set to the ID of the managed certificate.
	SSLId *string `json:"SSLId,omitempty" name:"SSLId"`

	// [Disused] When the certificate is an external certificate, the certificate content should be provided here. 
	Cert *string `json:"Cert,omitempty" name:"Cert"`

	// [Disused] When the certificate is an external certificate, the certificate key should be provided here. 
	PrivateKey *string `json:"PrivateKey,omitempty" name:"PrivateKey"`

	// Rule description.
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// Rule status. Valid values: `0` (the rule was successfully configured), `1` (configuring the rule), `2` (rule configuration failed), `3` (deleting the rule), `5` (failed to delete rule), `6` (rule awaiting configuration), `7` (rule awaiting deletion), and `8` (rule awaiting certificate configuration).
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// CC protection status. Valid values: `0` (disabled) and `1` (enabled).
	CCStatus *uint64 `json:"CCStatus,omitempty" name:"CCStatus"`

	// CC protection status based on HTTPS. Valid values: `0` (disabled) and `1` (enabled).
	CCEnable *uint64 `json:"CCEnable,omitempty" name:"CCEnable"`

	// CC protection threshold based on HTTPS.
	CCThreshold *uint64 `json:"CCThreshold,omitempty" name:"CCThreshold"`

	// CC protection level based on HTTPS.
	CCLevel *string `json:"CCLevel,omitempty" name:"CCLevel"`

	// Region code.
	Region *uint64 `json:"Region,omitempty" name:"Region"`

	// Anti-DDoS instance ID.
	Id *string `json:"Id,omitempty" name:"Id"`

	// Anti-DDoS instance IP address.
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// Modification time.
	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`

	// Whether to enable **Forward HTTPS requests via HTTP**. Valid values: `0` (disabled) and `1` (enabled). The default value is disabled.
	HttpsToHttpEnable *uint64 `json:"HttpsToHttpEnable,omitempty" name:"HttpsToHttpEnable"`

	// Access port number.
	// Note: this field may return null, indicating that no valid values can be obtained.
	VirtualPort *uint64 `json:"VirtualPort,omitempty" name:"VirtualPort"`
}

type OrderBy struct {

	// Sort by field name. Valid values: [
	// bandwidth (bandwidth),
	// overloadCount (number of times of exceeding peak value)
	// ]
	Field *string `json:"Field,omitempty" name:"Field"`

	// Sorting order. Valid values: [asc (ascending), desc (descending)]
	Order *string `json:"Order,omitempty" name:"Order"`
}

type Paging struct {

	// Starting position
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Quantity
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type ProtocolPort struct {

	// Protocol (TCP, UDP)
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Port
	Port *uint64 `json:"Port,omitempty" name:"Port"`
}

type RegionInstanceCount struct {

	// Region code
	Region *string `json:"Region,omitempty" name:"Region"`

	// Region code (new specification)
	RegionV3 *string `json:"RegionV3,omitempty" name:"RegionV3"`

	// Number of resource instances
	Count *uint64 `json:"Count,omitempty" name:"Count"`
}

type ResourceIp struct {

	// Anti-DDoS instance ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// Resource IP array
	IpList []*string `json:"IpList,omitempty" name:"IpList" list`
}

type SchedulingDomain struct {

	// Scheduling domain name
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// List of BGP IPs
	BGPIpList []*string `json:"BGPIpList,omitempty" name:"BGPIpList" list`

	// List of CTCC IPs
	CTCCIpList []*string `json:"CTCCIpList,omitempty" name:"CTCCIpList" list`

	// List of CUCC IPs
	CUCCIpList []*string `json:"CUCCIpList,omitempty" name:"CUCCIpList" list`

	// List of CMCC IPs
	CMCCIpList []*string `json:"CMCCIpList,omitempty" name:"CMCCIpList" list`

	// List of IPs outside Mainland China
	OverseaIpList []*string `json:"OverseaIpList,omitempty" name:"OverseaIpList" list`

	// Scheduling method. It only supports `priority` now.
	Method *string `json:"Method,omitempty" name:"Method"`

	// The creation time.
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 
	TTL *uint64 `json:"TTL,omitempty" name:"TTL"`

	// Status
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// Modification time
	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
}

type SuccessCode struct {

	// Success/error code
	Code *string `json:"Code,omitempty" name:"Code"`

	// Description
	Message *string `json:"Message,omitempty" name:"Message"`
}

type WaterPrintKey struct {

	// Watermark key ID
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

	// Watermark key value
	KeyContent *string `json:"KeyContent,omitempty" name:"KeyContent"`

	// Watermark key version number
	KeyVersion *string `json:"KeyVersion,omitempty" name:"KeyVersion"`

	// Whether it is enabled. Valid values: [0 (no), 1 (yes)]
	OpenStatus *uint64 `json:"OpenStatus,omitempty" name:"OpenStatus"`

	// Key generation time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type WaterPrintPolicy struct {

	// TCP port range, such as ["2000-3000","3500-4000"]
	TcpPortList []*string `json:"TcpPortList,omitempty" name:"TcpPortList" list`

	// UDP port range, such as ["2000-3000","3500-4000"]
	UdpPortList []*string `json:"UdpPortList,omitempty" name:"UdpPortList" list`

	// Watermark offset. Value range: [0, 100)
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Whether to automatically remove. Valid values: [0 (no), 1 (yes)]
	RemoveSwitch *uint64 `json:"RemoveSwitch,omitempty" name:"RemoveSwitch"`

	// Whether it is enabled. Valid values: [0 (no), 1 (yes)]
	OpenStatus *uint64 `json:"OpenStatus,omitempty" name:"OpenStatus"`
}
