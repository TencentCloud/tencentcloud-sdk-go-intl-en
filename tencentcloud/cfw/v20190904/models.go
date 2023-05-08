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

package v20190904

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

type AcListsData struct {
	// Rule ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// Access source
	// Note: This field may return `null`, indicating that no valid value was found.
	SourceIp *string `json:"SourceIp,omitempty" name:"SourceIp"`

	// Access destination
	// Note: This field may return `null`, indicating that no valid value was found.
	TargetIp *string `json:"TargetIp,omitempty" name:"TargetIp"`

	// Protocol
	// Note: This field may return `null`, indicating that no valid value was found.
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Port
	// Note: This field may return `null`, indicating that no valid value was found.
	Port *string `json:"Port,omitempty" name:"Port"`

	// Policy
	// Note: This field may return `null`, indicating that no valid value was found.
	Strategy *uint64 `json:"Strategy,omitempty" name:"Strategy"`

	// Description
	// Note: This field may return `null`, indicating that no valid value was found.
	Detail *string `json:"Detail,omitempty" name:"Detail"`

	// Hit count
	Count *uint64 `json:"Count,omitempty" name:"Count"`

	// Priority
	OrderIndex *uint64 `json:"OrderIndex,omitempty" name:"OrderIndex"`

	// Alert rule ID
	// Note: This field may return `null`, indicating that no valid value was found.
	LogId *string `json:"LogId,omitempty" name:"LogId"`
}

// Predefined struct for user
type AddAcRuleRequestParams struct {
	// -1: lowest priority; 1: highest priority
	OrderIndex *string `json:"OrderIndex,omitempty" name:"OrderIndex"`

	// The action that Cloud Firewall performs on the traffic. Valid values:
	// accept: allow
	// drop: deny
	// log: observe
	RuleAction *string `json:"RuleAction,omitempty" name:"RuleAction"`

	// The traffic direction for access control rules. Valid values:
	// in: incoming traffic access control
	// out: outgoing traffic access control
	Direction *string `json:"Direction,omitempty" name:"Direction"`

	// The description of access control rules.
	Description *string `json:"Description,omitempty" name:"Description"`

	// The type of source address in access control rules. Valid values:
	// net: source IP or range (IP or CIDR)
	// location: source region
	// template: CFW address template
	// instance: instance ID
	// vendor: Cloud vendor
	SourceType *string `json:"SourceType,omitempty" name:"SourceType"`

	// The source address in the access control policy. 
	// When `SourceType` is `net`, `SourceContent` is the source IP or CIDR block.
	// For example: 1.1.1.0/24
	// 
	// When `SourceType` is `template`, `SourceContent` must be the source address template ID.
	// 
	// When `SourceType` is `location`, `SourceContent` is the source region. 
	// For example, ["BJ11", "ZB"]
	// 
	// When `SourceType` is `instance`, `SourceContent` is the public IP of the instance.
	// For example, ins-xxxxx
	// 
	// When `SourceType` is `vendor`, `SourceContent` is the cloud service provider.
	// Values: `aws`, `huawei`, `tencent`, `aliyun`, `azure` and `all`. 
	SourceContent *string `json:"SourceContent,omitempty" name:"SourceContent"`

	// The type of destination address in access control rules. Valid values:
	// net: destination IP or range (IP or CIDR)
	// location: source region
	// template: CFW address template
	// instance: instance ID
	// vendor: Cloud vendor
	// domain: Domain name or IP.
	DestType *string `json:"DestType,omitempty" name:"DestType"`

	// The destination address in the access control policy. 
	// When `DestType` is `net`, `DestContent` is the destination IP or CIDR block.
	// For example: 1.1.1.0/24
	// 
	// When `DestType` is `template`, `DestContent` is the destination address template ID.
	// 
	// When `DestType` is `location`, `DestContent` is the destination region. 
	// For example, ["BJ11", "ZB"]
	// 
	// When `DestType` is `instance`, `DestContent` is the public IP of the instance.
	// For example, ins-xxxxx
	// 
	// When `DestType` is `domain`, `DestContent` is the domain name associated with the instance.
	// For example, *.qq.com
	// 
	// When `DestType`, `DestContent` is the selected cloud service provider.
	// Values: `aws`, `huawei`, `tencent`, `aliyun`, `azure` and `all`. 
	DestContent *string `json:"DestContent,omitempty" name:"DestContent"`

	// The port to apply access control rules. Valid values:
	// -1/-1: all ports
	// 80,443: 80 or 443
	Port *string `json:"Port,omitempty" name:"Port"`

	// The protocol type of traffic in access control rules. Valid value: TCP. Only TCP is supported for edge firewall rules. If this parameter is not specified, it defaults to TCP.
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// The Layer 7 protocol. Valid values:
	// HTTP/HTTPS
	// TLS/SSL
	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`

	// Indicates whether to enable the rules. Default: enable. Valid values:
	// true: enable; false: disable
	Enable *string `json:"Enable,omitempty" name:"Enable"`
}

type AddAcRuleRequest struct {
	*tchttp.BaseRequest
	
	// -1: lowest priority; 1: highest priority
	OrderIndex *string `json:"OrderIndex,omitempty" name:"OrderIndex"`

	// The action that Cloud Firewall performs on the traffic. Valid values:
	// accept: allow
	// drop: deny
	// log: observe
	RuleAction *string `json:"RuleAction,omitempty" name:"RuleAction"`

	// The traffic direction for access control rules. Valid values:
	// in: incoming traffic access control
	// out: outgoing traffic access control
	Direction *string `json:"Direction,omitempty" name:"Direction"`

	// The description of access control rules.
	Description *string `json:"Description,omitempty" name:"Description"`

	// The type of source address in access control rules. Valid values:
	// net: source IP or range (IP or CIDR)
	// location: source region
	// template: CFW address template
	// instance: instance ID
	// vendor: Cloud vendor
	SourceType *string `json:"SourceType,omitempty" name:"SourceType"`

	// The source address in the access control policy. 
	// When `SourceType` is `net`, `SourceContent` is the source IP or CIDR block.
	// For example: 1.1.1.0/24
	// 
	// When `SourceType` is `template`, `SourceContent` must be the source address template ID.
	// 
	// When `SourceType` is `location`, `SourceContent` is the source region. 
	// For example, ["BJ11", "ZB"]
	// 
	// When `SourceType` is `instance`, `SourceContent` is the public IP of the instance.
	// For example, ins-xxxxx
	// 
	// When `SourceType` is `vendor`, `SourceContent` is the cloud service provider.
	// Values: `aws`, `huawei`, `tencent`, `aliyun`, `azure` and `all`. 
	SourceContent *string `json:"SourceContent,omitempty" name:"SourceContent"`

	// The type of destination address in access control rules. Valid values:
	// net: destination IP or range (IP or CIDR)
	// location: source region
	// template: CFW address template
	// instance: instance ID
	// vendor: Cloud vendor
	// domain: Domain name or IP.
	DestType *string `json:"DestType,omitempty" name:"DestType"`

	// The destination address in the access control policy. 
	// When `DestType` is `net`, `DestContent` is the destination IP or CIDR block.
	// For example: 1.1.1.0/24
	// 
	// When `DestType` is `template`, `DestContent` is the destination address template ID.
	// 
	// When `DestType` is `location`, `DestContent` is the destination region. 
	// For example, ["BJ11", "ZB"]
	// 
	// When `DestType` is `instance`, `DestContent` is the public IP of the instance.
	// For example, ins-xxxxx
	// 
	// When `DestType` is `domain`, `DestContent` is the domain name associated with the instance.
	// For example, *.qq.com
	// 
	// When `DestType`, `DestContent` is the selected cloud service provider.
	// Values: `aws`, `huawei`, `tencent`, `aliyun`, `azure` and `all`. 
	DestContent *string `json:"DestContent,omitempty" name:"DestContent"`

	// The port to apply access control rules. Valid values:
	// -1/-1: all ports
	// 80,443: 80 or 443
	Port *string `json:"Port,omitempty" name:"Port"`

	// The protocol type of traffic in access control rules. Valid value: TCP. Only TCP is supported for edge firewall rules. If this parameter is not specified, it defaults to TCP.
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// The Layer 7 protocol. Valid values:
	// HTTP/HTTPS
	// TLS/SSL
	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`

	// Indicates whether to enable the rules. Default: enable. Valid values:
	// true: enable; false: disable
	Enable *string `json:"Enable,omitempty" name:"Enable"`
}

func (r *AddAcRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddAcRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OrderIndex")
	delete(f, "RuleAction")
	delete(f, "Direction")
	delete(f, "Description")
	delete(f, "SourceType")
	delete(f, "SourceContent")
	delete(f, "DestType")
	delete(f, "DestContent")
	delete(f, "Port")
	delete(f, "Protocol")
	delete(f, "ApplicationName")
	delete(f, "Enable")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddAcRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddAcRuleResponseParams struct {
	// UUID of the new rule
	RuleUuid *int64 `json:"RuleUuid,omitempty" name:"RuleUuid"`

	// 0: operation successful; -1: operation failed
	ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`

	// success: operation successful; failed: operation failed
	ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AddAcRuleResponse struct {
	*tchttp.BaseResponse
	Response *AddAcRuleResponseParams `json:"Response"`
}

func (r *AddAcRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddAcRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddEnterpriseSecurityGroupRulesRequestParams struct {
	// Creates rule data
	Data []*SecurityGroupRule `json:"Data,omitempty" name:"Data"`

	// Adding type. 0: add to the end; 1: add to the front; 2: insert. Default: 0
	Type *uint64 `json:"Type,omitempty" name:"Type"`

	// An identifier to ensure the idempotency of the request. The value of the ClientToken parameter is a unique string that is generated by your client and can contain up to 64 ASCII characters in length.
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`

	// Indicates whether to delay publishing. 1: delay; other values: do not delay
	IsDelay *uint64 `json:"IsDelay,omitempty" name:"IsDelay"`
}

type AddEnterpriseSecurityGroupRulesRequest struct {
	*tchttp.BaseRequest
	
	// Creates rule data
	Data []*SecurityGroupRule `json:"Data,omitempty" name:"Data"`

	// Adding type. 0: add to the end; 1: add to the front; 2: insert. Default: 0
	Type *uint64 `json:"Type,omitempty" name:"Type"`

	// An identifier to ensure the idempotency of the request. The value of the ClientToken parameter is a unique string that is generated by your client and can contain up to 64 ASCII characters in length.
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`

	// Indicates whether to delay publishing. 1: delay; other values: do not delay
	IsDelay *uint64 `json:"IsDelay,omitempty" name:"IsDelay"`
}

func (r *AddEnterpriseSecurityGroupRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddEnterpriseSecurityGroupRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Data")
	delete(f, "Type")
	delete(f, "ClientToken")
	delete(f, "IsDelay")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddEnterpriseSecurityGroupRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddEnterpriseSecurityGroupRulesResponseParams struct {
	// Status value. 0: added successfully; non-0: failed to add
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AddEnterpriseSecurityGroupRulesResponse struct {
	*tchttp.BaseResponse
	Response *AddEnterpriseSecurityGroupRulesResponseParams `json:"Response"`
}

func (r *AddEnterpriseSecurityGroupRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddEnterpriseSecurityGroupRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddNatAcRuleRequestParams struct {
	// NAT access control rules to be added.
	Rules []*CreateNatRuleItem `json:"Rules,omitempty" name:"Rules"`

	// Source of the rules to be added. Generally, this parameter is not used. The value insert_rule indicates that rules in the specified location are inserted, and the value batch_import indicates that rules are imported in batches. If the parameter is left empty, rules defined in the API request are added.
	From *string `json:"From,omitempty" name:"From"`
}

type AddNatAcRuleRequest struct {
	*tchttp.BaseRequest
	
	// NAT access control rules to be added.
	Rules []*CreateNatRuleItem `json:"Rules,omitempty" name:"Rules"`

	// Source of the rules to be added. Generally, this parameter is not used. The value insert_rule indicates that rules in the specified location are inserted, and the value batch_import indicates that rules are imported in batches. If the parameter is left empty, rules defined in the API request are added.
	From *string `json:"From,omitempty" name:"From"`
}

func (r *AddNatAcRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddNatAcRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Rules")
	delete(f, "From")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddNatAcRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddNatAcRuleResponseParams struct {
	// ID list of new rules.
	RuleUuid []*int64 `json:"RuleUuid,omitempty" name:"RuleUuid"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AddNatAcRuleResponse struct {
	*tchttp.BaseResponse
	Response *AddNatAcRuleResponseParams `json:"Response"`
}

func (r *AddNatAcRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddNatAcRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssetZone struct {
	// Region
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// Region
	ZoneEng *string `json:"ZoneEng,omitempty" name:"ZoneEng"`
}

type AssociatedInstanceInfo struct {
	// Instance ID
	// Note: This field may return `null`, indicating that no valid value was found.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Instance name
	// Note: This field may return `null`, indicating that no valid value was found.
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// Instance type. 3: CVM instance; 4: CLB instance; 5: ENI instance; 6: Cloud database
	// Note: This field may return `null`, indicating that no valid value was found.
	Type *uint64 `json:"Type,omitempty" name:"Type"`

	// VPC ID
	// Note: This field may return `null`, indicating that no valid value was found.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// VPC name
	// Note: This field may return `null`, indicating that no valid value was found.
	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`

	// Public IP
	// Note: This field may return `null`, indicating that no valid value was found.
	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`

	// Private IP
	// Note: This field may return `null`, indicating that no valid value was found.
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// The number of associated security groups
	// Note: This field may return `null`, indicating that no valid value was found.
	SecurityGroupCount *uint64 `json:"SecurityGroupCount,omitempty" name:"SecurityGroupCount"`
}

type BetaInfoByACL struct {
	// Task ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// Task name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// Last execution time
	// Note: This field may return null, indicating that no valid values can be obtained.
	LastTime *string `json:"LastTime,omitempty" name:"LastTime"`
}

type BlockIgnoreRule struct {
	// Domain name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Rule IP.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Ioc *string `json:"Ioc,omitempty" name:"Ioc"`

	// Threat level.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Level *string `json:"Level,omitempty" name:"Level"`

	// Source event name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	EventName *string `json:"EventName,omitempty" name:"EventName"`

	// Direction. Valid values: 0: outbound; 1: inbound.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Direction *int64 `json:"Direction,omitempty" name:"Direction"`

	// Protocol.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Address.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Address *string `json:"Address,omitempty" name:"Address"`

	// Rule type. Valid values: 1: block; 2: allow.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Action *int64 `json:"Action,omitempty" name:"Action"`

	// Time when a rule starts to take effect.
	// Note: This field may return null, indicating that no valid values can be obtained.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Time when a rule expires.
	// Note: This field may return null, indicating that no valid values can be obtained.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Reason for ignoring.
	// Note: This field may return null, indicating that no valid values can be obtained.
	IgnoreReason *string `json:"IgnoreReason,omitempty" name:"IgnoreReason"`

	// Security event source.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Source *string `json:"Source,omitempty" name:"Source"`

	// Rule ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	UniqueId *string `json:"UniqueId,omitempty" name:"UniqueId"`

	// Number of rule matching times.
	// Note: This field may return null, indicating that no valid values can be obtained.
	MatchTimes *int64 `json:"MatchTimes,omitempty" name:"MatchTimes"`

	// Country.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Country *string `json:"Country,omitempty" name:"Country"`


	Comment *string `json:"Comment,omitempty" name:"Comment"`
}

type CfwNatDnatRule struct {
	// Network protocol. Valid values: TCP or UDP.
	IpProtocol *string `json:"IpProtocol,omitempty" name:"IpProtocol"`

	// Elastic IP.
	PublicIpAddress *string `json:"PublicIpAddress,omitempty" name:"PublicIpAddress"`

	// Public port.
	PublicPort *int64 `json:"PublicPort,omitempty" name:"PublicPort"`

	// Private address.
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" name:"PrivateIpAddress"`

	// Private port.
	PrivatePort *int64 `json:"PrivatePort,omitempty" name:"PrivatePort"`

	// The description of NAT firewall forwarding rules.
	Description *string `json:"Description,omitempty" name:"Description"`
}

type CommonFilter struct {
	// Search key.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Search values.
	Values []*string `json:"Values,omitempty" name:"Values"`

	// Enum of integers that represent relations between Name and Values.
	// enum FilterOperatorType {
	//     // Invalid
	//     FILTER_OPERATOR_TYPE_INVALID = 0;
	//     // Equal to
	//     FILTER_OPERATOR_TYPE_EQUAL = 1;
	//     // Greater than
	//     FILTER_OPERATOR_TYPE_GREATER = 2;
	//     // Less than
	//     FILTER_OPERATOR_TYPE_LESS = 3;
	//     // Greater than or equal to
	//     FILTER_OPERATOR_TYPE_GREATER_EQ = 4;
	//     // Less than or equal to
	//     FILTER_OPERATOR_TYPE_LESS_EQ = 5;
	//     // Not equal to
	//     FILTER_OPERATOR_TYPE_NO_EQ = 6;
	//     // In (contained in the array)
	//     FILTER_OPERATOR_TYPE_IN = 7;
	//     // Not in
	//     FILTER_OPERATOR_TYPE_NOT_IN = 8;
	//     // Fuzzily matched
	//     FILTER_OPERATOR_TYPE_FUZZINESS = 9;
	//     // Existing
	//     FILTER_OPERATOR_TYPE_EXIST = 10;
	//     // Not existing
	//     FILTER_OPERATOR_TYPE_NOT_EXIST = 11;
	//     // Regular
	//     FILTER_OPERATOR_TYPE_REGULAR = 12;
	// }
	OperatorType *int64 `json:"OperatorType,omitempty" name:"OperatorType"`
}

// Predefined struct for user
type CreateAcRulesRequestParams struct {
	// Creates rule data
	Data []*RuleInfoData `json:"Data,omitempty" name:"Data"`

	// 0: add (default); 1: insert
	Type *uint64 `json:"Type,omitempty" name:"Type"`

	// Edge ID
	EdgeId *string `json:"EdgeId,omitempty" name:"EdgeId"`

	// Access control rule status
	Enable *int64 `json:"Enable,omitempty" name:"Enable"`

	// 0: add; 1: overwrite
	Overwrite *uint64 `json:"Overwrite,omitempty" name:"Overwrite"`

	// NAT instance ID, required when the parameter Area exists
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// portScan: from port scanning; patchImport: from batch import
	From *string `json:"From,omitempty" name:"From"`

	// NAT region
	Area *string `json:"Area,omitempty" name:"Area"`
}

type CreateAcRulesRequest struct {
	*tchttp.BaseRequest
	
	// Creates rule data
	Data []*RuleInfoData `json:"Data,omitempty" name:"Data"`

	// 0: add (default); 1: insert
	Type *uint64 `json:"Type,omitempty" name:"Type"`

	// Edge ID
	EdgeId *string `json:"EdgeId,omitempty" name:"EdgeId"`

	// Access control rule status
	Enable *int64 `json:"Enable,omitempty" name:"Enable"`

	// 0: add; 1: overwrite
	Overwrite *uint64 `json:"Overwrite,omitempty" name:"Overwrite"`

	// NAT instance ID, required when the parameter Area exists
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// portScan: from port scanning; patchImport: from batch import
	From *string `json:"From,omitempty" name:"From"`

	// NAT region
	Area *string `json:"Area,omitempty" name:"Area"`
}

func (r *CreateAcRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAcRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Data")
	delete(f, "Type")
	delete(f, "EdgeId")
	delete(f, "Enable")
	delete(f, "Overwrite")
	delete(f, "InstanceId")
	delete(f, "From")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAcRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAcRulesResponseParams struct {
	// Status value. 0: operation successful
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// Returns redundant information
	// Note: This field may return `null`, indicating that no valid value was found.
	Info *string `json:"Info,omitempty" name:"Info"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateAcRulesResponse struct {
	*tchttp.BaseResponse
	Response *CreateAcRulesResponseParams `json:"Response"`
}

func (r *CreateAcRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAcRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNatFwInstanceRequestParams struct {
	// Firewall instance name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Bandwidth
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// Mode. 1: use existing; 0: create new
	Mode *int64 `json:"Mode,omitempty" name:"Mode"`

	// Parameter passed for the Create New mode. Either NewModeItems or NatgwList is required.
	NewModeItems *NewModeItems `json:"NewModeItems,omitempty" name:"NewModeItems"`

	// NAT gateway list for the Using Existing mode. Either NewModeItems or NatgwList is required.
	NatGwList []*string `json:"NatGwList,omitempty" name:"NatGwList"`

	// Primary zone. The default zone is selected if it is empty.
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// Secondary zone. The default zone is selected if it is empty.
	ZoneBak *string `json:"ZoneBak,omitempty" name:"ZoneBak"`

	// Remote disaster recovery. 1: enable; 0: disable; empty: disable by default
	CrossAZone *int64 `json:"CrossAZone,omitempty" name:"CrossAZone"`

	// IP range of the firewall
	FwCidrInfo *FwCidrInfo `json:"FwCidrInfo,omitempty" name:"FwCidrInfo"`
}

type CreateNatFwInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Firewall instance name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Bandwidth
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// Mode. 1: use existing; 0: create new
	Mode *int64 `json:"Mode,omitempty" name:"Mode"`

	// Parameter passed for the Create New mode. Either NewModeItems or NatgwList is required.
	NewModeItems *NewModeItems `json:"NewModeItems,omitempty" name:"NewModeItems"`

	// NAT gateway list for the Using Existing mode. Either NewModeItems or NatgwList is required.
	NatGwList []*string `json:"NatGwList,omitempty" name:"NatGwList"`

	// Primary zone. The default zone is selected if it is empty.
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// Secondary zone. The default zone is selected if it is empty.
	ZoneBak *string `json:"ZoneBak,omitempty" name:"ZoneBak"`

	// Remote disaster recovery. 1: enable; 0: disable; empty: disable by default
	CrossAZone *int64 `json:"CrossAZone,omitempty" name:"CrossAZone"`

	// IP range of the firewall
	FwCidrInfo *FwCidrInfo `json:"FwCidrInfo,omitempty" name:"FwCidrInfo"`
}

func (r *CreateNatFwInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNatFwInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Width")
	delete(f, "Mode")
	delete(f, "NewModeItems")
	delete(f, "NatGwList")
	delete(f, "Zone")
	delete(f, "ZoneBak")
	delete(f, "CrossAZone")
	delete(f, "FwCidrInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateNatFwInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNatFwInstanceResponseParams struct {
	// Firewall instance ID
	CfwInsId *string `json:"CfwInsId,omitempty" name:"CfwInsId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateNatFwInstanceResponse struct {
	*tchttp.BaseResponse
	Response *CreateNatFwInstanceResponseParams `json:"Response"`
}

func (r *CreateNatFwInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNatFwInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNatFwInstanceWithDomainRequestParams struct {
	// Firewall instance name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Bandwidth
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// Mode. 1: use existing; 0: create new
	Mode *int64 `json:"Mode,omitempty" name:"Mode"`

	// Parameter passed for the Create New mode. Either NewModeItems or NatgwList is required.
	NewModeItems *NewModeItems `json:"NewModeItems,omitempty" name:"NewModeItems"`

	// NAT gateway list for the Using Existing mode. Either NewModeItems or NatgwList is required.
	NatGwList []*string `json:"NatGwList,omitempty" name:"NatGwList"`

	// Primary zone. The default zone is selected if it is empty.
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// Secondary zone. The default zone is selected if it is empty.
	ZoneBak *string `json:"ZoneBak,omitempty" name:"ZoneBak"`

	// Remote disaster recovery. 1: enable; 0: disable; empty: disable by default
	CrossAZone *int64 `json:"CrossAZone,omitempty" name:"CrossAZone"`

	// 0: not create; 1: create
	IsCreateDomain *int64 `json:"IsCreateDomain,omitempty" name:"IsCreateDomain"`

	// Required for creating a domain name
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// IP range of the firewall
	FwCidrInfo *FwCidrInfo `json:"FwCidrInfo,omitempty" name:"FwCidrInfo"`
}

type CreateNatFwInstanceWithDomainRequest struct {
	*tchttp.BaseRequest
	
	// Firewall instance name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Bandwidth
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// Mode. 1: use existing; 0: create new
	Mode *int64 `json:"Mode,omitempty" name:"Mode"`

	// Parameter passed for the Create New mode. Either NewModeItems or NatgwList is required.
	NewModeItems *NewModeItems `json:"NewModeItems,omitempty" name:"NewModeItems"`

	// NAT gateway list for the Using Existing mode. Either NewModeItems or NatgwList is required.
	NatGwList []*string `json:"NatGwList,omitempty" name:"NatGwList"`

	// Primary zone. The default zone is selected if it is empty.
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// Secondary zone. The default zone is selected if it is empty.
	ZoneBak *string `json:"ZoneBak,omitempty" name:"ZoneBak"`

	// Remote disaster recovery. 1: enable; 0: disable; empty: disable by default
	CrossAZone *int64 `json:"CrossAZone,omitempty" name:"CrossAZone"`

	// 0: not create; 1: create
	IsCreateDomain *int64 `json:"IsCreateDomain,omitempty" name:"IsCreateDomain"`

	// Required for creating a domain name
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// IP range of the firewall
	FwCidrInfo *FwCidrInfo `json:"FwCidrInfo,omitempty" name:"FwCidrInfo"`
}

func (r *CreateNatFwInstanceWithDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNatFwInstanceWithDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Width")
	delete(f, "Mode")
	delete(f, "NewModeItems")
	delete(f, "NatGwList")
	delete(f, "Zone")
	delete(f, "ZoneBak")
	delete(f, "CrossAZone")
	delete(f, "IsCreateDomain")
	delete(f, "Domain")
	delete(f, "FwCidrInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateNatFwInstanceWithDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNatFwInstanceWithDomainResponseParams struct {
	// NAT instance info
	// Note: This field may return `null`, indicating that no valid value was found.
	CfwInsId *string `json:"CfwInsId,omitempty" name:"CfwInsId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateNatFwInstanceWithDomainResponse struct {
	*tchttp.BaseResponse
	Response *CreateNatFwInstanceWithDomainResponseParams `json:"Response"`
}

func (r *CreateNatFwInstanceWithDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNatFwInstanceWithDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateNatRuleItem struct {
	// Access source. Example: `net: IP/CIDR(192.168.0.2)`
	SourceContent *string `json:"SourceContent,omitempty" name:"SourceContent"`

	// Access source type. Values for inbound rules: `ip`, `net`, `template`, and `location`. Values for outbound rules: `ip`, `net`, `template`, `instance`, `group`, and `tag`.
	SourceType *string `json:"SourceType,omitempty" name:"SourceType"`

	// Access target. Example: `net: IP/CIDR(192.168.0.2); domain: domain name rule, e.g., *.qq.com
	TargetContent *string `json:"TargetContent,omitempty" name:"TargetContent"`

	// Access target type. Values for inbound rules: `ip`, `net`, `template`, `instance`, `group`, and `tag`. Values for outbound rules: `ip`, `net`, `domain`, `template`, and `location`.
	TargetType *string `json:"TargetType,omitempty" name:"TargetType"`

	// Protocol. Values: `TCP`, `UDP`, `ICMP`, `ANY`, `HTTP`, `HTTPS`, `HTTP/HTTPS`, `SMTP`, `SMTPS`, `SMTP/SMTPS`, `FTP`, and `DNS`.
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Specify how the CFW instance deals with the traffic hit the access control rule. Values: `accept` (allow), `drop` (reject), and `log` (observe).
	RuleAction *string `json:"RuleAction,omitempty" name:"RuleAction"`

	// The port of the access control rule. Values: `-1/-1` (all ports) and `80` (Port 80)
	Port *string `json:"Port,omitempty" name:"Port"`

	// Rule direction. Values: `1` (Inbound) and `0` (Outbound)
	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`

	// Rule sequence number
	OrderIndex *int64 `json:"OrderIndex,omitempty" name:"OrderIndex"`

	// Rule status. `true` (Enabled); `false` (Disabled)
	Enable *string `json:"Enable,omitempty" name:"Enable"`

	// The unique ID of the rule, which is not required when you create a rule.
	Uuid *int64 `json:"Uuid,omitempty" name:"Uuid"`

	// Description
	Description *string `json:"Description,omitempty" name:"Description"`
}

// Predefined struct for user
type CreateSecurityGroupRulesRequestParams struct {
	// Added enterprise security group rule data
	Data []*SecurityGroupListData `json:"Data,omitempty" name:"Data"`

	// Direction. 0: outbound; 1: inbound. 1 by default
	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`

	// 0: at the end; 1: at the top; 2: in the middle. 0 by default
	Type *uint64 `json:"Type,omitempty" name:"Type"`

	// Indicates whether to enable rules after addition. 0: disable; 1: enable. 1 by default
	Enable *uint64 `json:"Enable,omitempty" name:"Enable"`
}

type CreateSecurityGroupRulesRequest struct {
	*tchttp.BaseRequest
	
	// Added enterprise security group rule data
	Data []*SecurityGroupListData `json:"Data,omitempty" name:"Data"`

	// Direction. 0: outbound; 1: inbound. 1 by default
	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`

	// 0: at the end; 1: at the top; 2: in the middle. 0 by default
	Type *uint64 `json:"Type,omitempty" name:"Type"`

	// Indicates whether to enable rules after addition. 0: disable; 1: enable. 1 by default
	Enable *uint64 `json:"Enable,omitempty" name:"Enable"`
}

func (r *CreateSecurityGroupRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSecurityGroupRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Data")
	delete(f, "Direction")
	delete(f, "Type")
	delete(f, "Enable")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSecurityGroupRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSecurityGroupRulesResponseParams struct {
	// Status value. 0: added successfully; non-0: failed to add
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateSecurityGroupRulesResponse struct {
	*tchttp.BaseResponse
	Response *CreateSecurityGroupRulesResponseParams `json:"Response"`
}

func (r *CreateSecurityGroupRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSecurityGroupRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAcRuleRequestParams struct {
	// The ID of the rule to delete. It can be queried via the DescribeAcLists API.
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// Direction. 0: outbound; 1: inbound
	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`

	// Edge ID between two VPCs
	EdgeId *string `json:"EdgeId,omitempty" name:"EdgeId"`

	// NAT region, e.g. ap-shanghai/ap-guangzhou/ap-chongqing
	Area *string `json:"Area,omitempty" name:"Area"`
}

type DeleteAcRuleRequest struct {
	*tchttp.BaseRequest
	
	// The ID of the rule to delete. It can be queried via the DescribeAcLists API.
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// Direction. 0: outbound; 1: inbound
	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`

	// Edge ID between two VPCs
	EdgeId *string `json:"EdgeId,omitempty" name:"EdgeId"`

	// NAT region, e.g. ap-shanghai/ap-guangzhou/ap-chongqing
	Area *string `json:"Area,omitempty" name:"Area"`
}

func (r *DeleteAcRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAcRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "Direction")
	delete(f, "EdgeId")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAcRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAcRuleResponseParams struct {
	// Status value. 0: deleted successfully; !0: deletion failed
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// Returns redundant information
	// Note: This field may return `null`, indicating that no valid value was found.
	Info *string `json:"Info,omitempty" name:"Info"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteAcRuleResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAcRuleResponseParams `json:"Response"`
}

func (r *DeleteAcRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAcRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAllAccessControlRuleRequestParams struct {
	// Direction. 0: outbound; 1: inbound. 0 by default
	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`

	// Deletes all the access control rules for inter-VPC firewall toggles associated with the EdgeId. It is empty by default. Enter either EdgeId or Area.
	EdgeId *string `json:"EdgeId,omitempty" name:"EdgeId"`

	// Deletes all the access control rules for NAT firewalls of this region. It is empty by default. Enter either EdgeId or Area.
	Area *string `json:"Area,omitempty" name:"Area"`
}

type DeleteAllAccessControlRuleRequest struct {
	*tchttp.BaseRequest
	
	// Direction. 0: outbound; 1: inbound. 0 by default
	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`

	// Deletes all the access control rules for inter-VPC firewall toggles associated with the EdgeId. It is empty by default. Enter either EdgeId or Area.
	EdgeId *string `json:"EdgeId,omitempty" name:"EdgeId"`

	// Deletes all the access control rules for NAT firewalls of this region. It is empty by default. Enter either EdgeId or Area.
	Area *string `json:"Area,omitempty" name:"Area"`
}

func (r *DeleteAllAccessControlRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAllAccessControlRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Direction")
	delete(f, "EdgeId")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAllAccessControlRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAllAccessControlRuleResponseParams struct {
	// Status of the task. `0`: Modified successfully; Others: Modification failed
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// Number of access control rules deleted.
	// Note: This field may return `null`, indicating that no valid value was found.
	Info *int64 `json:"Info,omitempty" name:"Info"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteAllAccessControlRuleResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAllAccessControlRuleResponseParams `json:"Response"`
}

func (r *DeleteAllAccessControlRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAllAccessControlRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteResourceGroupRequestParams struct {
	// Group ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

type DeleteResourceGroupRequest struct {
	*tchttp.BaseRequest
	
	// Group ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DeleteResourceGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteResourceGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteResourceGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteResourceGroupResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteResourceGroupResponse struct {
	*tchttp.BaseResponse
	Response *DeleteResourceGroupResponseParams `json:"Response"`
}

func (r *DeleteResourceGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteResourceGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSecurityGroupRuleRequestParams struct {
	// ID of the rule to delete
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// Tencent Cloud region (abbreviation)
	Area *string `json:"Area,omitempty" name:"Area"`

	// Direction. 0: outbound; 1: inbound
	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`

	// Indicates whether to delete the reverse rule. 0: no; 1: yes
	IsDelReverse *uint64 `json:"IsDelReverse,omitempty" name:"IsDelReverse"`
}

type DeleteSecurityGroupRuleRequest struct {
	*tchttp.BaseRequest
	
	// ID of the rule to delete
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// Tencent Cloud region (abbreviation)
	Area *string `json:"Area,omitempty" name:"Area"`

	// Direction. 0: outbound; 1: inbound
	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`

	// Indicates whether to delete the reverse rule. 0: no; 1: yes
	IsDelReverse *uint64 `json:"IsDelReverse,omitempty" name:"IsDelReverse"`
}

func (r *DeleteSecurityGroupRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSecurityGroupRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "Area")
	delete(f, "Direction")
	delete(f, "IsDelReverse")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSecurityGroupRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSecurityGroupRuleResponseParams struct {
	// Status value. 0: operation successful; non-0: operation failed
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// Returns redundant information
	// Note: This field may return `null`, indicating that no valid value was found.
	Info *string `json:"Info,omitempty" name:"Info"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteSecurityGroupRuleResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSecurityGroupRuleResponseParams `json:"Response"`
}

func (r *DeleteSecurityGroupRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSecurityGroupRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteVpcInstanceRequestParams struct {

}

type DeleteVpcInstanceRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DeleteVpcInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteVpcInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteVpcInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteVpcInstanceResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteVpcInstanceResponse struct {
	*tchttp.BaseResponse
	Response *DeleteVpcInstanceResponseParams `json:"Response"`
}

func (r *DeleteVpcInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteVpcInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescAcItem struct {
	// Access source.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SourceContent *string `json:"SourceContent,omitempty" name:"SourceContent"`

	// Access destination.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TargetContent *string `json:"TargetContent,omitempty" name:"TargetContent"`

	// Protocol.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Port.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Port *string `json:"Port,omitempty" name:"Port"`

	// Action that Cloud Firewall performs on the traffic. Valid values: accept (allow), drop (reject), and log (monitor).
	// Note: This field may return null, indicating that no valid values can be obtained.
	RuleAction *string `json:"RuleAction,omitempty" name:"RuleAction"`

	// Description.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Description *string `json:"Description,omitempty" name:"Description"`

	// Number of rule matching times.
	Count *uint64 `json:"Count,omitempty" name:"Count"`

	// Rule sequence number.
	OrderIndex *uint64 `json:"OrderIndex,omitempty" name:"OrderIndex"`

	// Access source type. Valid values for an inbound rule: ip, net, template, and location; valid values for an outbound rule: ip, net, template, instance, group, and tag.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SourceType *string `json:"SourceType,omitempty" name:"SourceType"`

	// Access destination type. Valid values for an inbound rule: ip, net, template, instance, group, and tag; valid values for an outbound rule: ip, net, domain, template, and location.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TargetType *string `json:"TargetType,omitempty" name:"TargetType"`

	// Unique ID of the rule.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Uuid *uint64 `json:"Uuid,omitempty" name:"Uuid"`

	// Rule validity.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Invalid *uint64 `json:"Invalid,omitempty" name:"Invalid"`

	// Valid values: 0: common rules; 1: regional rules.
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsRegion *uint64 `json:"IsRegion,omitempty" name:"IsRegion"`

	// Country ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CountryCode *uint64 `json:"CountryCode,omitempty" name:"CountryCode"`

	// City ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CityCode *uint64 `json:"CityCode,omitempty" name:"CityCode"`

	// Country name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CountryName *string `json:"CountryName,omitempty" name:"CountryName"`

	// City name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CityName *string `json:"CityName,omitempty" name:"CityName"`

	// Cloud provider code.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CloudCode *string `json:"CloudCode,omitempty" name:"CloudCode"`

	// Valid values: 0: common rules; 1: cloud provider rules.
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsCloud *uint64 `json:"IsCloud,omitempty" name:"IsCloud"`

	// Rule status. Valid values: true: enabled; false: disabled.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Enable *string `json:"Enable,omitempty" name:"Enable"`

	// Rule direction. Valid values: 1: inbound; 0: outbound.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`

	// Instance name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// UUID for internal use. Generally, this field is not required.
	// Note: This field may return null, indicating that no valid values can be obtained.
	InternalUuid *int64 `json:"InternalUuid,omitempty" name:"InternalUuid"`

	// Rule status. This field is valid when you query rule matching details. Valid values: 0: new; 1: deleted; 2: edited and deleted.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// Details of associated tasks
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	BetaList []*BetaInfoByACL `json:"BetaList,omitempty" name:"BetaList"`
}

// Predefined struct for user
type DescribeAcListsRequestParams struct {
	// Protocol
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Policy
	Strategy *string `json:"Strategy,omitempty" name:"Strategy"`

	// Search value
	SearchValue *string `json:"SearchValue,omitempty" name:"SearchValue"`

	// Number of entries per page
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Offset
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Indicates whether it is outbound or inbound. 1: inbound; 0: outbound
	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`

	// EdgeId value
	EdgeId *string `json:"EdgeId,omitempty" name:"EdgeId"`

	// Indicates whether the rule is enabled. '0': disabled; '1': enabled. '0' by default
	Status *string `json:"Status,omitempty" name:"Status"`

	// Region
	Area *string `json:"Area,omitempty" name:"Area"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type DescribeAcListsRequest struct {
	*tchttp.BaseRequest
	
	// Protocol
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Policy
	Strategy *string `json:"Strategy,omitempty" name:"Strategy"`

	// Search value
	SearchValue *string `json:"SearchValue,omitempty" name:"SearchValue"`

	// Number of entries per page
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Offset
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Indicates whether it is outbound or inbound. 1: inbound; 0: outbound
	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`

	// EdgeId value
	EdgeId *string `json:"EdgeId,omitempty" name:"EdgeId"`

	// Indicates whether the rule is enabled. '0': disabled; '1': enabled. '0' by default
	Status *string `json:"Status,omitempty" name:"Status"`

	// Region
	Area *string `json:"Area,omitempty" name:"Area"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeAcListsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAcListsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Protocol")
	delete(f, "Strategy")
	delete(f, "SearchValue")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Direction")
	delete(f, "EdgeId")
	delete(f, "Status")
	delete(f, "Area")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAcListsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAcListsResponseParams struct {
	// Total entries
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// Access control list data
	Data []*AcListsData `json:"Data,omitempty" name:"Data"`

	// Total entries excluding the filtered ones
	AllTotal *uint64 `json:"AllTotal,omitempty" name:"AllTotal"`

	// All access control rules enabled/disabled
	// Note: This field may return `null`, indicating that no valid value was found.
	Enable *uint64 `json:"Enable,omitempty" name:"Enable"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAcListsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAcListsResponseParams `json:"Response"`
}

func (r *DescribeAcListsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAcListsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssociatedInstanceListRequestParams struct {
	// List offset
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of records per page
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Region code (e.g. ap-guangzhou). All Tencent Cloud regions are supported.
	Area *string `json:"Area,omitempty" name:"Area"`

	// Additional search criteria (JSON string)
	SearchValue *string `json:"SearchValue,omitempty" name:"SearchValue"`

	// Sorting field
	By *string `json:"By,omitempty" name:"By"`

	// Sort order. asc: ascending; desc: descending
	Order *string `json:"Order,omitempty" name:"Order"`

	// Security group ID
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`

	// Instance type. '3': CVM instance; '4': CLB instance; '5': ENI instance; '6': Cloud database
	Type *string `json:"Type,omitempty" name:"Type"`
}

type DescribeAssociatedInstanceListRequest struct {
	*tchttp.BaseRequest
	
	// List offset
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of records per page
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Region code (e.g. ap-guangzhou). All Tencent Cloud regions are supported.
	Area *string `json:"Area,omitempty" name:"Area"`

	// Additional search criteria (JSON string)
	SearchValue *string `json:"SearchValue,omitempty" name:"SearchValue"`

	// Sorting field
	By *string `json:"By,omitempty" name:"By"`

	// Sort order. asc: ascending; desc: descending
	Order *string `json:"Order,omitempty" name:"Order"`

	// Security group ID
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`

	// Instance type. '3': CVM instance; '4': CLB instance; '5': ENI instance; '6': Cloud database
	Type *string `json:"Type,omitempty" name:"Type"`
}

func (r *DescribeAssociatedInstanceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssociatedInstanceListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Area")
	delete(f, "SearchValue")
	delete(f, "By")
	delete(f, "Order")
	delete(f, "SecurityGroupId")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAssociatedInstanceListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssociatedInstanceListResponseParams struct {
	// Number of instances
	// Note: This field may return `null`, indicating that no valid value was found.
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// Instance list
	// Note: This field may return `null`, indicating that no valid value was found.
	Data []*AssociatedInstanceInfo `json:"Data,omitempty" name:"Data"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAssociatedInstanceListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAssociatedInstanceListResponseParams `json:"Response"`
}

func (r *DescribeAssociatedInstanceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssociatedInstanceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBlockByIpTimesListRequestParams struct {
	// Start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// IP search criteria
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// Region
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// Direction
	Direction *string `json:"Direction,omitempty" name:"Direction"`

	// Source
	Source *string `json:"Source,omitempty" name:"Source"`

	// Inter-VPC firewall toggle edge ID
	EdgeId *string `json:"EdgeId,omitempty" name:"EdgeId"`

	// Log source. move: inter-VPC firewall
	LogSource *string `json:"LogSource,omitempty" name:"LogSource"`
}

type DescribeBlockByIpTimesListRequest struct {
	*tchttp.BaseRequest
	
	// Start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// IP search criteria
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// Region
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// Direction
	Direction *string `json:"Direction,omitempty" name:"Direction"`

	// Source
	Source *string `json:"Source,omitempty" name:"Source"`

	// Inter-VPC firewall toggle edge ID
	EdgeId *string `json:"EdgeId,omitempty" name:"EdgeId"`

	// Log source. move: inter-VPC firewall
	LogSource *string `json:"LogSource,omitempty" name:"LogSource"`
}

func (r *DescribeBlockByIpTimesListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBlockByIpTimesListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Ip")
	delete(f, "Zone")
	delete(f, "Direction")
	delete(f, "Source")
	delete(f, "EdgeId")
	delete(f, "LogSource")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBlockByIpTimesListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBlockByIpTimesListResponseParams struct {
	// Response data
	Data []*IpStatic `json:"Data,omitempty" name:"Data"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBlockByIpTimesListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBlockByIpTimesListResponseParams `json:"Response"`
}

func (r *DescribeBlockByIpTimesListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBlockByIpTimesListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBlockIgnoreListRequestParams struct {
	// Number of entries per page.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Page offset.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Direction. Valid values: 1: inbound public access; 0: outbound public access; 3: private network access; empty string: all access.
	Direction *string `json:"Direction,omitempty" name:"Direction"`

	// Rule type. Valid values: 1: block; 2: allow.
	RuleType *uint64 `json:"RuleType,omitempty" name:"RuleType"`

	// Column by which rules are sorted. Valid values: EndTime: end time; StartTime: start time; MatchTimes: number of matching times.
	Order *string `json:"Order,omitempty" name:"Order"`

	// Sort order. Valid values: desc: descending; asc: ascending.
	By *string `json:"By,omitempty" name:"By"`

	// Search keys, in a JSON string. Valid values: {}: empty; domain: domain name; level: threat level; ignore_reason: reason for allowing access; rule_source: source of a security event; address: geographical location; common: fuzzy search.
	SearchValue *string `json:"SearchValue,omitempty" name:"SearchValue"`
}

type DescribeBlockIgnoreListRequest struct {
	*tchttp.BaseRequest
	
	// Number of entries per page.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Page offset.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Direction. Valid values: 1: inbound public access; 0: outbound public access; 3: private network access; empty string: all access.
	Direction *string `json:"Direction,omitempty" name:"Direction"`

	// Rule type. Valid values: 1: block; 2: allow.
	RuleType *uint64 `json:"RuleType,omitempty" name:"RuleType"`

	// Column by which rules are sorted. Valid values: EndTime: end time; StartTime: start time; MatchTimes: number of matching times.
	Order *string `json:"Order,omitempty" name:"Order"`

	// Sort order. Valid values: desc: descending; asc: ascending.
	By *string `json:"By,omitempty" name:"By"`

	// Search keys, in a JSON string. Valid values: {}: empty; domain: domain name; level: threat level; ignore_reason: reason for allowing access; rule_source: source of a security event; address: geographical location; common: fuzzy search.
	SearchValue *string `json:"SearchValue,omitempty" name:"SearchValue"`
}

func (r *DescribeBlockIgnoreListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBlockIgnoreListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Direction")
	delete(f, "RuleType")
	delete(f, "Order")
	delete(f, "By")
	delete(f, "SearchValue")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBlockIgnoreListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBlockIgnoreListResponseParams struct {
	// List data.
	Data []*BlockIgnoreRule `json:"Data,omitempty" name:"Data"`

	// Total number of results, which is used for pagination.
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// Status code. Valid values: 0: successful; others: failed.
	ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`

	// Status message. Valid values: success: successful query; fail: failed query.
	ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBlockIgnoreListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBlockIgnoreListResponseParams `json:"Response"`
}

func (r *DescribeBlockIgnoreListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBlockIgnoreListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBlockStaticListRequestParams struct {
	// Start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// List type. Valid values: port, address, or IP
	QueryType *string `json:"QueryType,omitempty" name:"QueryType"`

	// Number of top results returned
	Top *int64 `json:"Top,omitempty" name:"Top"`

	// Search criteria
	SearchValue *string `json:"SearchValue,omitempty" name:"SearchValue"`
}

type DescribeBlockStaticListRequest struct {
	*tchttp.BaseRequest
	
	// Start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// List type. Valid values: port, address, or IP
	QueryType *string `json:"QueryType,omitempty" name:"QueryType"`

	// Number of top results returned
	Top *int64 `json:"Top,omitempty" name:"Top"`

	// Search criteria
	SearchValue *string `json:"SearchValue,omitempty" name:"SearchValue"`
}

func (r *DescribeBlockStaticListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBlockStaticListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "QueryType")
	delete(f, "Top")
	delete(f, "SearchValue")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBlockStaticListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBlockStaticListResponseParams struct {
	// None
	Data []*StaticInfo `json:"Data,omitempty" name:"Data"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBlockStaticListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBlockStaticListResponseParams `json:"Response"`
}

func (r *DescribeBlockStaticListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBlockStaticListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDefenseSwitchRequestParams struct {

}

type DescribeDefenseSwitchRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeDefenseSwitchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDefenseSwitchRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDefenseSwitchRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDefenseSwitchResponseParams struct {
	// Whether to enable the Basic Protection feature
	BasicRuleSwitch *int64 `json:"BasicRuleSwitch,omitempty" name:"BasicRuleSwitch"`

	// Whether to enable the Security Baseline feature
	BaselineAllSwitch *int64 `json:"BaselineAllSwitch,omitempty" name:"BaselineAllSwitch"`

	// Whether to enable the Treat Intelligence feature
	TiSwitch *int64 `json:"TiSwitch,omitempty" name:"TiSwitch"`

	// Whether to enable the Virtual Patch feature
	VirtualPatchSwitch *int64 `json:"VirtualPatchSwitch,omitempty" name:"VirtualPatchSwitch"`

	// Whether it has been enabled before
	HistoryOpen *int64 `json:"HistoryOpen,omitempty" name:"HistoryOpen"`

	// Status code. `0`: Succeeded. Others: Failed
	ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`

	// Status message. `success` and `fail.
	ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDefenseSwitchResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDefenseSwitchResponseParams `json:"Response"`
}

func (r *DescribeDefenseSwitchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDefenseSwitchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEnterpriseSecurityGroupRuleRequestParams struct {
	// Page number of the current page displayed for query by page number.
	// 
	// 1 by default.
	PageNo *string `json:"PageNo,omitempty" name:"PageNo"`

	// Maximum number of entries per page displayed for query by page number.
	// 
	// Maximum value: 50.
	PageSize *string `json:"PageSize,omitempty" name:"PageSize"`

	// Source example:
	// net: IP/CIDR (192.168.0.2)
	// template: parameter template (ipm-dyodhpby)
	// instance: asset instance (ins-123456)
	// resourcegroup: asset group (/all groups/group 1/subgroup 1)
	// tag: resource tag ({"Key":"tag key","Value":"tag value"})
	// region: region (ap-gaungzhou)
	// Wildcards are supported.
	SourceContent *string `json:"SourceContent,omitempty" name:"SourceContent"`

	// Destination example:
	// net: IP/CIDR (192.168.0.2)
	// template: parameter template (ipm-dyodhpby)
	// instance: asset instance (ins-123456)
	// resourcegroup: asset group (/all groups/group 1/subgroup 1)
	// tag: resource tag ({"Key":"tag key","Value":"tag value"})
	// region: region (ap-gaungzhou)
	// Wildcards are supported.
	DestContent *string `json:"DestContent,omitempty" name:"DestContent"`

	// Rule description. Wildcards are supported.
	Description *string `json:"Description,omitempty" name:"Description"`

	// The action that Cloud Firewall performs on the traffic. Valid values:
	// accept: allow
	// drop: deny
	RuleAction *string `json:"RuleAction,omitempty" name:"RuleAction"`

	// Indicates whether to enable the rules. Default: enable. Valid values:
	// true: enable; false: disable
	Enable *string `json:"Enable,omitempty" name:"Enable"`

	// The port to apply access control rules. Valid values:
	// -1/-1: all ports
	// 80: port 80
	Port *string `json:"Port,omitempty" name:"Port"`

	// Protocol. TCP/UDP/ICMP/ANY
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Parameter template ID of port and protocol type; mutually exclusive with Protocol and Port
	ServiceTemplateId *string `json:"ServiceTemplateId,omitempty" name:"ServiceTemplateId"`
}

type DescribeEnterpriseSecurityGroupRuleRequest struct {
	*tchttp.BaseRequest
	
	// Page number of the current page displayed for query by page number.
	// 
	// 1 by default.
	PageNo *string `json:"PageNo,omitempty" name:"PageNo"`

	// Maximum number of entries per page displayed for query by page number.
	// 
	// Maximum value: 50.
	PageSize *string `json:"PageSize,omitempty" name:"PageSize"`

	// Source example:
	// net: IP/CIDR (192.168.0.2)
	// template: parameter template (ipm-dyodhpby)
	// instance: asset instance (ins-123456)
	// resourcegroup: asset group (/all groups/group 1/subgroup 1)
	// tag: resource tag ({"Key":"tag key","Value":"tag value"})
	// region: region (ap-gaungzhou)
	// Wildcards are supported.
	SourceContent *string `json:"SourceContent,omitempty" name:"SourceContent"`

	// Destination example:
	// net: IP/CIDR (192.168.0.2)
	// template: parameter template (ipm-dyodhpby)
	// instance: asset instance (ins-123456)
	// resourcegroup: asset group (/all groups/group 1/subgroup 1)
	// tag: resource tag ({"Key":"tag key","Value":"tag value"})
	// region: region (ap-gaungzhou)
	// Wildcards are supported.
	DestContent *string `json:"DestContent,omitempty" name:"DestContent"`

	// Rule description. Wildcards are supported.
	Description *string `json:"Description,omitempty" name:"Description"`

	// The action that Cloud Firewall performs on the traffic. Valid values:
	// accept: allow
	// drop: deny
	RuleAction *string `json:"RuleAction,omitempty" name:"RuleAction"`

	// Indicates whether to enable the rules. Default: enable. Valid values:
	// true: enable; false: disable
	Enable *string `json:"Enable,omitempty" name:"Enable"`

	// The port to apply access control rules. Valid values:
	// -1/-1: all ports
	// 80: port 80
	Port *string `json:"Port,omitempty" name:"Port"`

	// Protocol. TCP/UDP/ICMP/ANY
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Parameter template ID of port and protocol type; mutually exclusive with Protocol and Port
	ServiceTemplateId *string `json:"ServiceTemplateId,omitempty" name:"ServiceTemplateId"`
}

func (r *DescribeEnterpriseSecurityGroupRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEnterpriseSecurityGroupRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageNo")
	delete(f, "PageSize")
	delete(f, "SourceContent")
	delete(f, "DestContent")
	delete(f, "Description")
	delete(f, "RuleAction")
	delete(f, "Enable")
	delete(f, "Port")
	delete(f, "Protocol")
	delete(f, "ServiceTemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEnterpriseSecurityGroupRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEnterpriseSecurityGroupRuleResponseParams struct {
	// Page number of the current page displayed for query by page number.
	PageNo *string `json:"PageNo,omitempty" name:"PageNo"`

	// Maximum number of entries per page displayed for query by page number.
	PageSize *string `json:"PageSize,omitempty" name:"PageSize"`

	// Access control rule list
	Rules []*SecurityGroupRule `json:"Rules,omitempty" name:"Rules"`

	// Total number of access control rules
	TotalCount *string `json:"TotalCount,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeEnterpriseSecurityGroupRuleResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEnterpriseSecurityGroupRuleResponseParams `json:"Response"`
}

func (r *DescribeEnterpriseSecurityGroupRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEnterpriseSecurityGroupRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGuideScanInfoRequestParams struct {

}

type DescribeGuideScanInfoRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeGuideScanInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGuideScanInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGuideScanInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGuideScanInfoResponseParams struct {
	// Scan information
	Data *ScanInfo `json:"Data,omitempty" name:"Data"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeGuideScanInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGuideScanInfoResponseParams `json:"Response"`
}

func (r *DescribeGuideScanInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGuideScanInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIPStatusListRequestParams struct {
	// Asset ID
	IPList []*string `json:"IPList,omitempty" name:"IPList"`
}

type DescribeIPStatusListRequest struct {
	*tchttp.BaseRequest
	
	// Asset ID
	IPList []*string `json:"IPList,omitempty" name:"IPList"`
}

func (r *DescribeIPStatusListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIPStatusListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IPList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIPStatusListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIPStatusListResponseParams struct {
	// IP status information
	StatusList []*IPDefendStatus `json:"StatusList,omitempty" name:"StatusList"`

	// Status code
	ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`

	// Status information
	ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeIPStatusListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIPStatusListResponseParams `json:"Response"`
}

func (r *DescribeIPStatusListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIPStatusListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNatAcRuleRequestParams struct {
	// Number of entries per page.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Page offset.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Index to be queried. This parameter is optional, and is used only in specific cases.
	Index *string `json:"Index,omitempty" name:"Index"`

	// Filter condition combination.
	Filters []*CommonFilter `json:"Filters,omitempty" name:"Filters"`

	// Start time for search. This parameter is optional.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time for search. This parameter is optional.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Valid values: desc: descending; asc: ascending. The returned results are sorted by the value of By. If this parameter is specified, By is also required.
	Order *string `json:"Order,omitempty" name:"Order"`

	// Field by which the returned results are sorted.
	By *string `json:"By,omitempty" name:"By"`
}

type DescribeNatAcRuleRequest struct {
	*tchttp.BaseRequest
	
	// Number of entries per page.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Page offset.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Index to be queried. This parameter is optional, and is used only in specific cases.
	Index *string `json:"Index,omitempty" name:"Index"`

	// Filter condition combination.
	Filters []*CommonFilter `json:"Filters,omitempty" name:"Filters"`

	// Start time for search. This parameter is optional.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time for search. This parameter is optional.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Valid values: desc: descending; asc: ascending. The returned results are sorted by the value of By. If this parameter is specified, By is also required.
	Order *string `json:"Order,omitempty" name:"Order"`

	// Field by which the returned results are sorted.
	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeNatAcRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNatAcRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Index")
	delete(f, "Filters")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNatAcRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNatAcRuleResponseParams struct {
	// Total number of entries.
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// NAT access control list data.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data []*DescAcItem `json:"Data,omitempty" name:"Data"`

	// Total number of entries returned without filtering.
	AllTotal *uint64 `json:"AllTotal,omitempty" name:"AllTotal"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeNatAcRuleResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNatAcRuleResponseParams `json:"Response"`
}

func (r *DescribeNatAcRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNatAcRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNatFwInfoCountRequestParams struct {

}

type DescribeNatFwInfoCountRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeNatFwInfoCountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNatFwInfoCountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNatFwInfoCountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNatFwInfoCountResponseParams struct {
	// Response parameter
	// Note: This field may return `null`, indicating that no valid value was found.
	ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`

	// Number of NAT instances of the current tenant
	// Note: This field may return `null`, indicating that no valid value was found.
	NatFwInsCount *int64 `json:"NatFwInsCount,omitempty" name:"NatFwInsCount"`

	// Number of subnets connected by the current tenant
	// Note: This field may return `null`, indicating that no valid value was found.
	SubnetCount *int64 `json:"SubnetCount,omitempty" name:"SubnetCount"`

	// Number of firewalls enabled
	// Note: This field may return `null`, indicating that no valid value was found.
	OpenSwitchCount *int64 `json:"OpenSwitchCount,omitempty" name:"OpenSwitchCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeNatFwInfoCountResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNatFwInfoCountResponseParams `json:"Response"`
}

func (r *DescribeNatFwInfoCountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNatFwInfoCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNatFwInstanceRequestParams struct {

}

type DescribeNatFwInstanceRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeNatFwInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNatFwInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNatFwInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNatFwInstanceResponseParams struct {
	// Instance array
	NatinsLst []*NatFwInstance `json:"NatinsLst,omitempty" name:"NatinsLst"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeNatFwInstanceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNatFwInstanceResponseParams `json:"Response"`
}

func (r *DescribeNatFwInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNatFwInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNatFwInstanceWithRegionRequestParams struct {

}

type DescribeNatFwInstanceWithRegionRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeNatFwInstanceWithRegionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNatFwInstanceWithRegionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNatFwInstanceWithRegionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNatFwInstanceWithRegionResponseParams struct {
	// Instance array
	// Note: This field may return `null`, indicating that no valid value was found.
	NatinsLst []*NatFwInstance `json:"NatinsLst,omitempty" name:"NatinsLst"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeNatFwInstanceWithRegionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNatFwInstanceWithRegionResponseParams `json:"Response"`
}

func (r *DescribeNatFwInstanceWithRegionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNatFwInstanceWithRegionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNatFwInstancesInfoRequestParams struct {
	// Gets filtering fields of instance list
	Filter []*NatFwFilter `json:"Filter,omitempty" name:"Filter"`

	// Page number
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Page length
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeNatFwInstancesInfoRequest struct {
	*tchttp.BaseRequest
	
	// Gets filtering fields of instance list
	Filter []*NatFwFilter `json:"Filter,omitempty" name:"Filter"`

	// Page number
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Page length
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeNatFwInstancesInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNatFwInstancesInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filter")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNatFwInstancesInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNatFwInstancesInfoResponseParams struct {
	// Instance card info array
	// Note: This field may return `null`, indicating that no valid value was found.
	NatinsLst []*NatInstanceInfo `json:"NatinsLst,omitempty" name:"NatinsLst"`

	// Number of NAT firewalls
	// Note: This field may return `null`, indicating that no valid value was found.
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeNatFwInstancesInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNatFwInstancesInfoResponseParams `json:"Response"`
}

func (r *DescribeNatFwInstancesInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNatFwInstancesInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNatFwVpcDnsLstRequestParams struct {
	// NAT firewall instance ID
	NatFwInsId *string `json:"NatFwInsId,omitempty" name:"NatFwInsId"`

	// Content filtered by NAT firewall, separated with ","
	NatInsIdFilter *string `json:"NatInsIdFilter,omitempty" name:"NatInsIdFilter"`

	// Number of pages
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Maximum entries per page
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeNatFwVpcDnsLstRequest struct {
	*tchttp.BaseRequest
	
	// NAT firewall instance ID
	NatFwInsId *string `json:"NatFwInsId,omitempty" name:"NatFwInsId"`

	// Content filtered by NAT firewall, separated with ","
	NatInsIdFilter *string `json:"NatInsIdFilter,omitempty" name:"NatInsIdFilter"`

	// Number of pages
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Maximum entries per page
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeNatFwVpcDnsLstRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNatFwVpcDnsLstRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NatFwInsId")
	delete(f, "NatInsIdFilter")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNatFwVpcDnsLstRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNatFwVpcDnsLstResponseParams struct {
	// VPC DNS info array of NAT firewall
	// Note: This field may return `null`, indicating that no valid value was found.
	VpcDnsSwitchLst []*VpcDnsInfo `json:"VpcDnsSwitchLst,omitempty" name:"VpcDnsSwitchLst"`

	// Response parameter
	// Note: This field may return `null`, indicating that no valid value was found.
	ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`

	// Total number of toggles
	// Note: This field may return `null`, indicating that no valid value was found.
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeNatFwVpcDnsLstResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNatFwVpcDnsLstResponseParams `json:"Response"`
}

func (r *DescribeNatFwVpcDnsLstResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNatFwVpcDnsLstResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourceGroupNewRequestParams struct {
	// Query type. Network–VPC; business recognition–resource; resource tag–tag
	QueryType *string `json:"QueryType,omitempty" name:"QueryType"`

	// Asset group ID, 0: all asset group IDs
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// all: all, including subgroups; own: my asset groups only
	ShowType *string `json:"ShowType,omitempty" name:"ShowType"`
}

type DescribeResourceGroupNewRequest struct {
	*tchttp.BaseRequest
	
	// Query type. Network–VPC; business recognition–resource; resource tag–tag
	QueryType *string `json:"QueryType,omitempty" name:"QueryType"`

	// Asset group ID, 0: all asset group IDs
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// all: all, including subgroups; own: my asset groups only
	ShowType *string `json:"ShowType,omitempty" name:"ShowType"`
}

func (r *DescribeResourceGroupNewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourceGroupNewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "QueryType")
	delete(f, "GroupId")
	delete(f, "ShowType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeResourceGroupNewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourceGroupNewResponseParams struct {
	// Returns a tree structure
	Data *string `json:"Data,omitempty" name:"Data"`

	// Number of uncategorizd instances
	UnResourceNum *int64 `json:"UnResourceNum,omitempty" name:"UnResourceNum"`

	// Response message
	ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`

	// Return code. 0: Request successful
	ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeResourceGroupNewResponse struct {
	*tchttp.BaseResponse
	Response *DescribeResourceGroupNewResponseParams `json:"Response"`
}

func (r *DescribeResourceGroupNewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourceGroupNewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRuleOverviewRequestParams struct {
	// Direction. 0: outbound; 1: inbound
	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`
}

type DescribeRuleOverviewRequest struct {
	*tchttp.BaseRequest
	
	// Direction. 0: outbound; 1: inbound
	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`
}

func (r *DescribeRuleOverviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRuleOverviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Direction")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRuleOverviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRuleOverviewResponseParams struct {
	// Total number of rules
	// Note: This field may return `null`, indicating that no valid value was found.
	AllTotal *uint64 `json:"AllTotal,omitempty" name:"AllTotal"`

	// Number of blocking rules
	// Note: This field may return `null`, indicating that no valid value was found.
	StrategyNum *uint64 `json:"StrategyNum,omitempty" name:"StrategyNum"`

	// Number of enabled rules
	// Note: This field may return `null`, indicating that no valid value was found.
	StartRuleNum *uint64 `json:"StartRuleNum,omitempty" name:"StartRuleNum"`

	// Number of disabled rules
	// Note: This field may return `null`, indicating that no valid value was found.
	StopRuleNum *uint64 `json:"StopRuleNum,omitempty" name:"StopRuleNum"`

	// Remaining quota
	// Note: This field may return `null`, indicating that no valid value was found.
	RemainingNum *int64 `json:"RemainingNum,omitempty" name:"RemainingNum"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeRuleOverviewResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRuleOverviewResponseParams `json:"Response"`
}

func (r *DescribeRuleOverviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRuleOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityGroupListRequestParams struct {
	// 0: outbound rule; 1: inbound rule
	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`

	// Region code (e.g. ap-guangzhou ). All Tencent Cloud regions are supported.
	Area *string `json:"Area,omitempty" name:"Area"`

	// Search value
	SearchValue *string `json:"SearchValue,omitempty" name:"SearchValue"`

	// Number of entries per page. Default: 10
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default: 0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Status. Null: all; '0': filter disabled rules; '1': filter enabled rules
	Status *string `json:"Status,omitempty" name:"Status"`

	// 0: not filter; 1: filter out normal rules to retain abnormal rules
	Filter *uint64 `json:"Filter,omitempty" name:"Filter"`
}

type DescribeSecurityGroupListRequest struct {
	*tchttp.BaseRequest
	
	// 0: outbound rule; 1: inbound rule
	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`

	// Region code (e.g. ap-guangzhou ). All Tencent Cloud regions are supported.
	Area *string `json:"Area,omitempty" name:"Area"`

	// Search value
	SearchValue *string `json:"SearchValue,omitempty" name:"SearchValue"`

	// Number of entries per page. Default: 10
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default: 0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Status. Null: all; '0': filter disabled rules; '1': filter enabled rules
	Status *string `json:"Status,omitempty" name:"Status"`

	// 0: not filter; 1: filter out normal rules to retain abnormal rules
	Filter *uint64 `json:"Filter,omitempty" name:"Filter"`
}

func (r *DescribeSecurityGroupListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityGroupListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Direction")
	delete(f, "Area")
	delete(f, "SearchValue")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Status")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecurityGroupListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityGroupListResponseParams struct {
	// Total rules in the list
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// Security group rule list data
	Data []*SecurityGroupListData `json:"Data,omitempty" name:"Data"`

	// Total entries excluding the filtered ones
	AllTotal *uint64 `json:"AllTotal,omitempty" name:"AllTotal"`

	// All access control rules enabled/disabled
	// Note: This field may return `null`, indicating that no valid value was found.
	Enable *uint64 `json:"Enable,omitempty" name:"Enable"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSecurityGroupListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSecurityGroupListResponseParams `json:"Response"`
}

func (r *DescribeSecurityGroupListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityGroupListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSourceAssetRequestParams struct {
	// Fuzzy search
	FuzzySearch *string `json:"FuzzySearch,omitempty" name:"FuzzySearch"`

	// Asset type. 1: public network; 2: private network
	InsType *string `json:"InsType,omitempty" name:"InsType"`

	// If ChooseType is 1, grouped assets are queried; if ChooseType is not 1, non-grouped assets are queried
	ChooseType *string `json:"ChooseType,omitempty" name:"ChooseType"`

	// Region
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// Maximum number of results returned per page. For example, if it is set to 10, 10 results will be returned at most.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Offset of search results
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeSourceAssetRequest struct {
	*tchttp.BaseRequest
	
	// Fuzzy search
	FuzzySearch *string `json:"FuzzySearch,omitempty" name:"FuzzySearch"`

	// Asset type. 1: public network; 2: private network
	InsType *string `json:"InsType,omitempty" name:"InsType"`

	// If ChooseType is 1, grouped assets are queried; if ChooseType is not 1, non-grouped assets are queried
	ChooseType *string `json:"ChooseType,omitempty" name:"ChooseType"`

	// Region
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// Maximum number of results returned per page. For example, if it is set to 10, 10 results will be returned at most.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Offset of search results
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeSourceAssetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSourceAssetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FuzzySearch")
	delete(f, "InsType")
	delete(f, "ChooseType")
	delete(f, "Zone")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSourceAssetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSourceAssetResponseParams struct {
	// Region collection
	ZoneList []*AssetZone `json:"ZoneList,omitempty" name:"ZoneList"`

	// Data
	Data []*InstanceInfo `json:"Data,omitempty" name:"Data"`

	// Total number of returned data
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSourceAssetResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSourceAssetResponseParams `json:"Response"`
}

func (r *DescribeSourceAssetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSourceAssetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSwitchListsRequestParams struct {
	// Firewall status. 0: disabled; 1: enabled
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// Asset type, e.g. CVM/NAT/VPN/CLB/others
	Type *string `json:"Type,omitempty" name:"Type"`

	// Region, e.g. Shanghai, Chongqing, Guangzhou, etc.
	Area *string `json:"Area,omitempty" name:"Area"`

	// Search value, e.g. "{"common":"106.54.189.45"}"
	SearchValue *string `json:"SearchValue,omitempty" name:"SearchValue"`

	// Number of entries. Default: 10
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default: 0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Sort order. desc: descending; asc: ascending
	Order *string `json:"Order,omitempty" name:"Order"`

	// Sorting field. PortTimes (number of risky ports)
	By *string `json:"By,omitempty" name:"By"`
}

type DescribeSwitchListsRequest struct {
	*tchttp.BaseRequest
	
	// Firewall status. 0: disabled; 1: enabled
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// Asset type, e.g. CVM/NAT/VPN/CLB/others
	Type *string `json:"Type,omitempty" name:"Type"`

	// Region, e.g. Shanghai, Chongqing, Guangzhou, etc.
	Area *string `json:"Area,omitempty" name:"Area"`

	// Search value, e.g. "{"common":"106.54.189.45"}"
	SearchValue *string `json:"SearchValue,omitempty" name:"SearchValue"`

	// Number of entries. Default: 10
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default: 0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Sort order. desc: descending; asc: ascending
	Order *string `json:"Order,omitempty" name:"Order"`

	// Sorting field. PortTimes (number of risky ports)
	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeSwitchListsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSwitchListsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Status")
	delete(f, "Type")
	delete(f, "Area")
	delete(f, "SearchValue")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSwitchListsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSwitchListsResponseParams struct {
	// Total entries
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// List data
	Data []*SwitchListsData `json:"Data,omitempty" name:"Data"`

	// Region list
	AreaLists []*string `json:"AreaLists,omitempty" name:"AreaLists"`

	// Number of enabled firewalls
	// Note: This field may return `null`, indicating that no valid value was found.
	OnNum *uint64 `json:"OnNum,omitempty" name:"OnNum"`

	// Number of disabled firewalls
	// Note: This field may return `null`, indicating that no valid value was found.
	OffNum *uint64 `json:"OffNum,omitempty" name:"OffNum"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSwitchListsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSwitchListsResponseParams `json:"Response"`
}

func (r *DescribeSwitchListsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSwitchListsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTLogInfoRequestParams struct {
	// Start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Type. 1: alert; 2: block
	QueryType *string `json:"QueryType,omitempty" name:"QueryType"`

	// Search criteria
	SearchValue *string `json:"SearchValue,omitempty" name:"SearchValue"`
}

type DescribeTLogInfoRequest struct {
	*tchttp.BaseRequest
	
	// Start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Type. 1: alert; 2: block
	QueryType *string `json:"QueryType,omitempty" name:"QueryType"`

	// Search criteria
	SearchValue *string `json:"SearchValue,omitempty" name:"SearchValue"`
}

func (r *DescribeTLogInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTLogInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "QueryType")
	delete(f, "SearchValue")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTLogInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTLogInfoResponseParams struct {
	// `NetworkNum`: Number of detected network scans
	//  `HandleNum`: Number of pending processing events
	// "BanNum": 
	//   `VulNum`: Number of vulnerability exploits
	//   "OutNum`: Number of compromised servers
	// "BruteForceNum": 0
	Data *TLogInfo `json:"Data,omitempty" name:"Data"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTLogInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTLogInfoResponseParams `json:"Response"`
}

func (r *DescribeTLogInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTLogInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTLogIpListRequestParams struct {
	// Start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Type. 1: alert; 2: block
	QueryType *string `json:"QueryType,omitempty" name:"QueryType"`

	// Number of top results returned
	Top *int64 `json:"Top,omitempty" name:"Top"`

	// Search criteria
	SearchValue *string `json:"SearchValue,omitempty" name:"SearchValue"`
}

type DescribeTLogIpListRequest struct {
	*tchttp.BaseRequest
	
	// Start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Type. 1: alert; 2: block
	QueryType *string `json:"QueryType,omitempty" name:"QueryType"`

	// Number of top results returned
	Top *int64 `json:"Top,omitempty" name:"Top"`

	// Search criteria
	SearchValue *string `json:"SearchValue,omitempty" name:"SearchValue"`
}

func (r *DescribeTLogIpListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTLogIpListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "QueryType")
	delete(f, "Top")
	delete(f, "SearchValue")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTLogIpListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTLogIpListResponseParams struct {
	// Data collection
	Data []*StaticInfo `json:"Data,omitempty" name:"Data"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTLogIpListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTLogIpListResponseParams `json:"Response"`
}

func (r *DescribeTLogIpListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTLogIpListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTableStatusRequestParams struct {
	// Edge ID between two VPCs, required for VPCs
	EdgeId *string `json:"EdgeId,omitempty" name:"EdgeId"`

	// Status value. 0: the only default value
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// NAT region, required for NAT
	Area *string `json:"Area,omitempty" name:"Area"`

	// Direction. 0: outbound; 1: inbound. 0 by default
	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`
}

type DescribeTableStatusRequest struct {
	*tchttp.BaseRequest
	
	// Edge ID between two VPCs, required for VPCs
	EdgeId *string `json:"EdgeId,omitempty" name:"EdgeId"`

	// Status value. 0: the only default value
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// NAT region, required for NAT
	Area *string `json:"Area,omitempty" name:"Area"`

	// Direction. 0: outbound; 1: inbound. 0 by default
	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`
}

func (r *DescribeTableStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTableStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EdgeId")
	delete(f, "Status")
	delete(f, "Area")
	delete(f, "Direction")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTableStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTableStatusResponseParams struct {
	// 0: normal; non-0: abnormal
	// Note: This field may return `null`, indicating that no valid value was found.
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTableStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTableStatusResponseParams `json:"Response"`
}

func (r *DescribeTableStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTableStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUnHandleEventTabListRequestParams struct {
	// Start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Gets example ID
	AssetID *string `json:"AssetID,omitempty" name:"AssetID"`
}

type DescribeUnHandleEventTabListRequest struct {
	*tchttp.BaseRequest
	
	// Start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Gets example ID
	AssetID *string `json:"AssetID,omitempty" name:"AssetID"`
}

func (r *DescribeUnHandleEventTabListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUnHandleEventTabListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "AssetID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUnHandleEventTabListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUnHandleEventTabListResponseParams struct {
	// Gets unhandled security events
	// Note: This field may return `null`, indicating that no valid value was found.
	Data *UnHandleEvent `json:"Data,omitempty" name:"Data"`

	// Error code. 0: success; non-0: error
	ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`

	// Return message: success
	ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeUnHandleEventTabListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUnHandleEventTabListResponseParams `json:"Response"`
}

func (r *DescribeUnHandleEventTabListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUnHandleEventTabListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DnsVpcSwitch struct {
	// VPC ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 0: off; 1: on
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

// Predefined struct for user
type ExpandCfwVerticalRequestParams struct {
	// nat: NAT firewall, ew: east-west firewall
	FwType *string `json:"FwType,omitempty" name:"FwType"`

	// Bandwidth value
	Width *uint64 `json:"Width,omitempty" name:"Width"`

	// Firewall instance ID
	CfwInstance *string `json:"CfwInstance,omitempty" name:"CfwInstance"`
}

type ExpandCfwVerticalRequest struct {
	*tchttp.BaseRequest
	
	// nat: NAT firewall, ew: east-west firewall
	FwType *string `json:"FwType,omitempty" name:"FwType"`

	// Bandwidth value
	Width *uint64 `json:"Width,omitempty" name:"Width"`

	// Firewall instance ID
	CfwInstance *string `json:"CfwInstance,omitempty" name:"CfwInstance"`
}

func (r *ExpandCfwVerticalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExpandCfwVerticalRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FwType")
	delete(f, "Width")
	delete(f, "CfwInstance")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExpandCfwVerticalRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExpandCfwVerticalResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ExpandCfwVerticalResponse struct {
	*tchttp.BaseResponse
	Response *ExpandCfwVerticalResponseParams `json:"Response"`
}

func (r *ExpandCfwVerticalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExpandCfwVerticalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FwCidrInfo struct {
	// The IP range type of the firewall. Values: `VpcSelf` (VPC IP range preferred); `Assis` (Secondary IP range preferred); `Custom` (Custom IP range)
	FwCidrType *string `json:"FwCidrType,omitempty" name:"FwCidrType"`

	// The IP segment assigned for each VPC.
	FwCidrLst []*FwVpcCidr `json:"FwCidrLst,omitempty" name:"FwCidrLst"`

	// The IP segment used by other firewalls. Specify this if you want to assign a dedicated segment for the firewall.
	ComFwCidr *string `json:"ComFwCidr,omitempty" name:"ComFwCidr"`
}

type FwVpcCidr struct {
	// VPC ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// IP range of the firewall. The mask must be at least /24.
	FwCidr *string `json:"FwCidr,omitempty" name:"FwCidr"`
}

type IPDefendStatus struct {
	// IP address
	IP *string `json:"IP,omitempty" name:"IP"`

	// Protection status. 1: enabled; -1: incorrect address; others: disabled
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type InstanceInfo struct {
	// App ID
	AppId *string `json:"AppId,omitempty" name:"AppId"`

	// Region
	Region *string `json:"Region,omitempty" name:"Region"`

	// VPC ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// VPC name
	VPCName *string `json:"VPCName,omitempty" name:"VPCName"`

	// Subnet ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// Asset ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Asset name
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// Asset type
	//  3: CVM instance; 4: CLB instance; 5: ENI instance; 6: MySQL; 7: Redis; 8: NAT; 9: VPN; 10: ES; 11: MariaDB; 12: Kafka; 13: NATFW
	InsType *int64 `json:"InsType,omitempty" name:"InsType"`

	// Public IP
	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`

	// Private IP
	PrivateIp *string `json:"PrivateIp,omitempty" name:"PrivateIp"`

	// Number of ports
	PortNum *string `json:"PortNum,omitempty" name:"PortNum"`

	// Number of vulnerabilities
	LeakNum *string `json:"LeakNum,omitempty" name:"LeakNum"`

	// 1: public network; 2: private network
	InsSource *string `json:"InsSource,omitempty" name:"InsSource"`

	// [a,b]
	// Note: This field may return `null`, indicating that no valid value was found.
	ResourcePath []*string `json:"ResourcePath,omitempty" name:"ResourcePath"`
}

type IocListData struct {
	// IP address to be handled. Either IP or Domain is required.
	IP *string `json:"IP,omitempty" name:"IP"`

	// 0 or 1. 0: outbound; 1: inbound
	Direction *int64 `json:"Direction,omitempty" name:"Direction"`

	// Domain name to be handled. Either IP or Domain is required.
	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

type IpStatic struct {
	// Value
	Num *int64 `json:"Num,omitempty" name:"Num"`

	// Time shown on the x-axis of the line graph
	StatTime *string `json:"StatTime,omitempty" name:"StatTime"`
}

// Predefined struct for user
type ModifyAcRuleRequestParams struct {
	// Array of rules
	Data []*RuleInfoData `json:"Data,omitempty" name:"Data"`

	// EdgeId value
	EdgeId *string `json:"EdgeId,omitempty" name:"EdgeId"`

	// Access rule status
	Enable *int64 `json:"Enable,omitempty" name:"Enable"`

	// NAT region
	Area *string `json:"Area,omitempty" name:"Area"`
}

type ModifyAcRuleRequest struct {
	*tchttp.BaseRequest
	
	// Array of rules
	Data []*RuleInfoData `json:"Data,omitempty" name:"Data"`

	// EdgeId value
	EdgeId *string `json:"EdgeId,omitempty" name:"EdgeId"`

	// Access rule status
	Enable *int64 `json:"Enable,omitempty" name:"Enable"`

	// NAT region
	Area *string `json:"Area,omitempty" name:"Area"`
}

func (r *ModifyAcRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAcRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Data")
	delete(f, "EdgeId")
	delete(f, "Enable")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAcRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAcRuleResponseParams struct {
	// Status value. 0: operation successful; non-0: operation failed
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// Returns redundant information
	// Note: This field may return `null`, indicating that no valid value was found.
	Info *string `json:"Info,omitempty" name:"Info"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyAcRuleResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAcRuleResponseParams `json:"Response"`
}

func (r *ModifyAcRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAcRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAllPublicIPSwitchStatusRequestParams struct {
	// Status. 0: off; 1: on
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// ID of the selected firewall toggle
	FireWallPublicIPs []*string `json:"FireWallPublicIPs,omitempty" name:"FireWallPublicIPs"`
}

type ModifyAllPublicIPSwitchStatusRequest struct {
	*tchttp.BaseRequest
	
	// Status. 0: off; 1: on
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// ID of the selected firewall toggle
	FireWallPublicIPs []*string `json:"FireWallPublicIPs,omitempty" name:"FireWallPublicIPs"`
}

func (r *ModifyAllPublicIPSwitchStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAllPublicIPSwitchStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Status")
	delete(f, "FireWallPublicIPs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAllPublicIPSwitchStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAllPublicIPSwitchStatusResponseParams struct {
	// Return message
	// Note: This field may return `null`, indicating that no valid value was found.
	ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`

	// Error code. 0: success; non-0: failed
	// Note: This field may return `null`, indicating that no valid value was found.
	ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyAllPublicIPSwitchStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAllPublicIPSwitchStatusResponseParams `json:"Response"`
}

func (r *ModifyAllPublicIPSwitchStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAllPublicIPSwitchStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAllRuleStatusRequestParams struct {
	// Status. 0: all disabled; 1: all enabled
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// Direction. 0: outbound; 1: inbound
	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`

	// Edge ID value
	EdgeId *string `json:"EdgeId,omitempty" name:"EdgeId"`

	// NAT region
	Area *string `json:"Area,omitempty" name:"Area"`
}

type ModifyAllRuleStatusRequest struct {
	*tchttp.BaseRequest
	
	// Status. 0: all disabled; 1: all enabled
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// Direction. 0: outbound; 1: inbound
	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`

	// Edge ID value
	EdgeId *string `json:"EdgeId,omitempty" name:"EdgeId"`

	// NAT region
	Area *string `json:"Area,omitempty" name:"Area"`
}

func (r *ModifyAllRuleStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAllRuleStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Status")
	delete(f, "Direction")
	delete(f, "EdgeId")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAllRuleStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAllRuleStatusResponseParams struct {
	// 0: modified successfully; non-0: modification failed
	// Note: This field may return `null`, indicating that no valid value was found.
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyAllRuleStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAllRuleStatusResponseParams `json:"Response"`
}

func (r *ModifyAllRuleStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAllRuleStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAllVPCSwitchStatusRequestParams struct {
	// Status. 0: off; 1: on
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// ID of the selected firewall toggle
	FireWallVpcIds []*string `json:"FireWallVpcIds,omitempty" name:"FireWallVpcIds"`
}

type ModifyAllVPCSwitchStatusRequest struct {
	*tchttp.BaseRequest
	
	// Status. 0: off; 1: on
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// ID of the selected firewall toggle
	FireWallVpcIds []*string `json:"FireWallVpcIds,omitempty" name:"FireWallVpcIds"`
}

func (r *ModifyAllVPCSwitchStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAllVPCSwitchStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Status")
	delete(f, "FireWallVpcIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAllVPCSwitchStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAllVPCSwitchStatusResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyAllVPCSwitchStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAllVPCSwitchStatusResponseParams `json:"Response"`
}

func (r *ModifyAllVPCSwitchStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAllVPCSwitchStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAssetScanRequestParams struct {
	// Scan range. 1: port; 2: port + vulnerability scan
	ScanRange *int64 `json:"ScanRange,omitempty" name:"ScanRange"`

	// Scan mode: 'heavy', 'medium', 'light'
	ScanDeep *string `json:"ScanDeep,omitempty" name:"ScanDeep"`

	// Scan type. 1: scan now; 2: periodic scan
	RangeType *int64 `json:"RangeType,omitempty" name:"RangeType"`

	// Scheduled task time, required when RangeType is 2
	ScanPeriod *string `json:"ScanPeriod,omitempty" name:"ScanPeriod"`

	// Scans this field now and passes the filtered IPs
	ScanFilterIp []*string `json:"ScanFilterIp,omitempty" name:"ScanFilterIp"`

	// 1: all; 2: single
	ScanType *int64 `json:"ScanType,omitempty" name:"ScanType"`
}

type ModifyAssetScanRequest struct {
	*tchttp.BaseRequest
	
	// Scan range. 1: port; 2: port + vulnerability scan
	ScanRange *int64 `json:"ScanRange,omitempty" name:"ScanRange"`

	// Scan mode: 'heavy', 'medium', 'light'
	ScanDeep *string `json:"ScanDeep,omitempty" name:"ScanDeep"`

	// Scan type. 1: scan now; 2: periodic scan
	RangeType *int64 `json:"RangeType,omitempty" name:"RangeType"`

	// Scheduled task time, required when RangeType is 2
	ScanPeriod *string `json:"ScanPeriod,omitempty" name:"ScanPeriod"`

	// Scans this field now and passes the filtered IPs
	ScanFilterIp []*string `json:"ScanFilterIp,omitempty" name:"ScanFilterIp"`

	// 1: all; 2: single
	ScanType *int64 `json:"ScanType,omitempty" name:"ScanType"`
}

func (r *ModifyAssetScanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAssetScanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ScanRange")
	delete(f, "ScanDeep")
	delete(f, "RangeType")
	delete(f, "ScanPeriod")
	delete(f, "ScanFilterIp")
	delete(f, "ScanType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAssetScanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAssetScanResponseParams struct {
	// Return message
	// Note: This field may return `null`, indicating that no valid value was found.
	ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`

	// Error code. 0: success; non-0: failed
	// Note: This field may return `null`, indicating that no valid value was found.
	ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`

	// Status value. 0: success; 1: scanning; others: failed
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyAssetScanResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAssetScanResponseParams `json:"Response"`
}

func (r *ModifyAssetScanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAssetScanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBlockIgnoreListRequestParams struct {
	// Type of the rule. Values: `1` (Blocklist); `2` (Allowlist)
	RuleType *int64 `json:"RuleType,omitempty" name:"RuleType"`

	// Either IP or Domain is required
	IOC []*IocListData `json:"IOC,omitempty" name:"IOC"`

	// Optional values: delete, edit, and add
	IocAction *string `json:"IocAction,omitempty" name:"IocAction"`

	// Time format: yyyy-MM-dd HH:mm:ss. Required when IocAction is edit or add
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time of the period in the format of yyyy-MM-dd HH:mm:ss. It must be later than both the start time and the current time. It’s required when `IocAction` is `edit` or `add`. 
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type ModifyBlockIgnoreListRequest struct {
	*tchttp.BaseRequest
	
	// Type of the rule. Values: `1` (Blocklist); `2` (Allowlist)
	RuleType *int64 `json:"RuleType,omitempty" name:"RuleType"`

	// Either IP or Domain is required
	IOC []*IocListData `json:"IOC,omitempty" name:"IOC"`

	// Optional values: delete, edit, and add
	IocAction *string `json:"IocAction,omitempty" name:"IocAction"`

	// Time format: yyyy-MM-dd HH:mm:ss. Required when IocAction is edit or add
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time of the period in the format of yyyy-MM-dd HH:mm:ss. It must be later than both the start time and the current time. It’s required when `IocAction` is `edit` or `add`. 
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *ModifyBlockIgnoreListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBlockIgnoreListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleType")
	delete(f, "IOC")
	delete(f, "IocAction")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyBlockIgnoreListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBlockIgnoreListResponseParams struct {
	// Return message
	ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`

	// Error code. 0: success; non-0: failed
	ReturnCode *uint64 `json:"ReturnCode,omitempty" name:"ReturnCode"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyBlockIgnoreListResponse struct {
	*tchttp.BaseResponse
	Response *ModifyBlockIgnoreListResponseParams `json:"Response"`
}

func (r *ModifyBlockIgnoreListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBlockIgnoreListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBlockTopRequestParams struct {
	// Record ID
	UniqueId *string `json:"UniqueId,omitempty" name:"UniqueId"`

	// Operation type. 1: pin to top; 0: unpin
	OpeType *string `json:"OpeType,omitempty" name:"OpeType"`
}

type ModifyBlockTopRequest struct {
	*tchttp.BaseRequest
	
	// Record ID
	UniqueId *string `json:"UniqueId,omitempty" name:"UniqueId"`

	// Operation type. 1: pin to top; 0: unpin
	OpeType *string `json:"OpeType,omitempty" name:"OpeType"`
}

func (r *ModifyBlockTopRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBlockTopRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UniqueId")
	delete(f, "OpeType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyBlockTopRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBlockTopResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyBlockTopResponse struct {
	*tchttp.BaseResponse
	Response *ModifyBlockTopResponseParams `json:"Response"`
}

func (r *ModifyBlockTopResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBlockTopResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyEnterpriseSecurityDispatchStatusRequestParams struct {
	// Status. Values: `0` (Publish now), `1` (Stop publishing)
	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

type ModifyEnterpriseSecurityDispatchStatusRequest struct {
	*tchttp.BaseRequest
	
	// Status. Values: `0` (Publish now), `1` (Stop publishing)
	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyEnterpriseSecurityDispatchStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyEnterpriseSecurityDispatchStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyEnterpriseSecurityDispatchStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyEnterpriseSecurityDispatchStatusResponseParams struct {
	// `0`: Modified successfully; Others: Modification failed
	// Note: This field may return null, indicating that no valid values can be obtained.
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyEnterpriseSecurityDispatchStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyEnterpriseSecurityDispatchStatusResponseParams `json:"Response"`
}

func (r *ModifyEnterpriseSecurityDispatchStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyEnterpriseSecurityDispatchStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyEnterpriseSecurityGroupRuleRequestParams struct {
	// UUID of the rule, which can be obtained by querying the rule list
	RuleUuid *uint64 `json:"RuleUuid,omitempty" name:"RuleUuid"`

	// Modification type. Values: `0` (Modify rule content), `1` (Toggle on/off a rule) and `2` (Toggle on/off all rules)
	ModifyType *uint64 `json:"ModifyType,omitempty" name:"ModifyType"`

	// The new rule content you want. It’s only required when you want to modify the rule content (`ModifyType=0`)
	Data *SecurityGroupRule `json:"Data,omitempty" name:"Data"`

	// `0`: Do not enable; `1`: Enable
	Enable *uint64 `json:"Enable,omitempty" name:"Enable"`
}

type ModifyEnterpriseSecurityGroupRuleRequest struct {
	*tchttp.BaseRequest
	
	// UUID of the rule, which can be obtained by querying the rule list
	RuleUuid *uint64 `json:"RuleUuid,omitempty" name:"RuleUuid"`

	// Modification type. Values: `0` (Modify rule content), `1` (Toggle on/off a rule) and `2` (Toggle on/off all rules)
	ModifyType *uint64 `json:"ModifyType,omitempty" name:"ModifyType"`

	// The new rule content you want. It’s only required when you want to modify the rule content (`ModifyType=0`)
	Data *SecurityGroupRule `json:"Data,omitempty" name:"Data"`

	// `0`: Do not enable; `1`: Enable
	Enable *uint64 `json:"Enable,omitempty" name:"Enable"`
}

func (r *ModifyEnterpriseSecurityGroupRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyEnterpriseSecurityGroupRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleUuid")
	delete(f, "ModifyType")
	delete(f, "Data")
	delete(f, "Enable")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyEnterpriseSecurityGroupRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyEnterpriseSecurityGroupRuleResponseParams struct {
	// Status value. `0`: Edited successfully; Others: Failed to edit
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// ID of new rule generated after the modification
	NewRuleUuid *uint64 `json:"NewRuleUuid,omitempty" name:"NewRuleUuid"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyEnterpriseSecurityGroupRuleResponse struct {
	*tchttp.BaseResponse
	Response *ModifyEnterpriseSecurityGroupRuleResponseParams `json:"Response"`
}

func (r *ModifyEnterpriseSecurityGroupRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyEnterpriseSecurityGroupRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNatAcRuleRequestParams struct {
	// Array of rules to be modified.
	Rules []*CreateNatRuleItem `json:"Rules,omitempty" name:"Rules"`
}

type ModifyNatAcRuleRequest struct {
	*tchttp.BaseRequest
	
	// Array of rules to be modified.
	Rules []*CreateNatRuleItem `json:"Rules,omitempty" name:"Rules"`
}

func (r *ModifyNatAcRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNatAcRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Rules")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyNatAcRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNatAcRuleResponseParams struct {
	// ID list of new rules that have been successfully modified.
	RuleUuid []*int64 `json:"RuleUuid,omitempty" name:"RuleUuid"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyNatAcRuleResponse struct {
	*tchttp.BaseResponse
	Response *ModifyNatAcRuleResponseParams `json:"Response"`
}

func (r *ModifyNatAcRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNatAcRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNatFwReSelectRequestParams struct {
	// Mode. 1: use existing; 0: create new
	Mode *int64 `json:"Mode,omitempty" name:"Mode"`

	// Firewall instance ID
	CfwInstance *string `json:"CfwInstance,omitempty" name:"CfwInstance"`

	// List of NAT gateways reconnected for the Using Existing mode. Only one of NatGwList and VpcList can be passed.
	NatGwList []*string `json:"NatGwList,omitempty" name:"NatGwList"`

	// List of VPCs reconnected for the Create New mode. Only one of NatGwList and VpcList can be passed.
	VpcList []*string `json:"VpcList,omitempty" name:"VpcList"`

	// IP range of the firewall
	FwCidrInfo *FwCidrInfo `json:"FwCidrInfo,omitempty" name:"FwCidrInfo"`
}

type ModifyNatFwReSelectRequest struct {
	*tchttp.BaseRequest
	
	// Mode. 1: use existing; 0: create new
	Mode *int64 `json:"Mode,omitempty" name:"Mode"`

	// Firewall instance ID
	CfwInstance *string `json:"CfwInstance,omitempty" name:"CfwInstance"`

	// List of NAT gateways reconnected for the Using Existing mode. Only one of NatGwList and VpcList can be passed.
	NatGwList []*string `json:"NatGwList,omitempty" name:"NatGwList"`

	// List of VPCs reconnected for the Create New mode. Only one of NatGwList and VpcList can be passed.
	VpcList []*string `json:"VpcList,omitempty" name:"VpcList"`

	// IP range of the firewall
	FwCidrInfo *FwCidrInfo `json:"FwCidrInfo,omitempty" name:"FwCidrInfo"`
}

func (r *ModifyNatFwReSelectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNatFwReSelectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Mode")
	delete(f, "CfwInstance")
	delete(f, "NatGwList")
	delete(f, "VpcList")
	delete(f, "FwCidrInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyNatFwReSelectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNatFwReSelectResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyNatFwReSelectResponse struct {
	*tchttp.BaseResponse
	Response *ModifyNatFwReSelectResponseParams `json:"Response"`
}

func (r *ModifyNatFwReSelectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNatFwReSelectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNatFwSwitchRequestParams struct {
	// Status. 0: off; 1: on
	Enable *int64 `json:"Enable,omitempty" name:"Enable"`

	// List of firewall instance IDs. Only one of CfwInsIdList, SubnetIdList, and RouteTableIdList can be passed.
	CfwInsIdList []*string `json:"CfwInsIdList,omitempty" name:"CfwInsIdList"`

	// List of subnet IDs. Only one of CfwInsIdList, SubnetIdList, and RouteTableIdList can be passed.
	SubnetIdList []*string `json:"SubnetIdList,omitempty" name:"SubnetIdList"`

	// List of route table IDs. Only one of CfwInsIdList, SubnetIdList, and RouteTableIdList can be passed.
	RouteTableIdList []*string `json:"RouteTableIdList,omitempty" name:"RouteTableIdList"`
}

type ModifyNatFwSwitchRequest struct {
	*tchttp.BaseRequest
	
	// Status. 0: off; 1: on
	Enable *int64 `json:"Enable,omitempty" name:"Enable"`

	// List of firewall instance IDs. Only one of CfwInsIdList, SubnetIdList, and RouteTableIdList can be passed.
	CfwInsIdList []*string `json:"CfwInsIdList,omitempty" name:"CfwInsIdList"`

	// List of subnet IDs. Only one of CfwInsIdList, SubnetIdList, and RouteTableIdList can be passed.
	SubnetIdList []*string `json:"SubnetIdList,omitempty" name:"SubnetIdList"`

	// List of route table IDs. Only one of CfwInsIdList, SubnetIdList, and RouteTableIdList can be passed.
	RouteTableIdList []*string `json:"RouteTableIdList,omitempty" name:"RouteTableIdList"`
}

func (r *ModifyNatFwSwitchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNatFwSwitchRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Enable")
	delete(f, "CfwInsIdList")
	delete(f, "SubnetIdList")
	delete(f, "RouteTableIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyNatFwSwitchRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNatFwSwitchResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyNatFwSwitchResponse struct {
	*tchttp.BaseResponse
	Response *ModifyNatFwSwitchResponseParams `json:"Response"`
}

func (r *ModifyNatFwSwitchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNatFwSwitchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNatFwVpcDnsSwitchRequestParams struct {
	// NAT firewall ID
	NatFwInsId *string `json:"NatFwInsId,omitempty" name:"NatFwInsId"`

	// DNS toggle list
	DnsVpcSwitchLst []*DnsVpcSwitch `json:"DnsVpcSwitchLst,omitempty" name:"DnsVpcSwitchLst"`
}

type ModifyNatFwVpcDnsSwitchRequest struct {
	*tchttp.BaseRequest
	
	// NAT firewall ID
	NatFwInsId *string `json:"NatFwInsId,omitempty" name:"NatFwInsId"`

	// DNS toggle list
	DnsVpcSwitchLst []*DnsVpcSwitch `json:"DnsVpcSwitchLst,omitempty" name:"DnsVpcSwitchLst"`
}

func (r *ModifyNatFwVpcDnsSwitchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNatFwVpcDnsSwitchRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NatFwInsId")
	delete(f, "DnsVpcSwitchLst")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyNatFwVpcDnsSwitchRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNatFwVpcDnsSwitchResponseParams struct {
	// Modified successfully
	// Note: This field may return `null`, indicating that no valid value was found.
	ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyNatFwVpcDnsSwitchResponse struct {
	*tchttp.BaseResponse
	Response *ModifyNatFwVpcDnsSwitchResponseParams `json:"Response"`
}

func (r *ModifyNatFwVpcDnsSwitchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNatFwVpcDnsSwitchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNatSequenceRulesRequestParams struct {
	// Rule sequence number. Values: `OrderIndex` (Original sequence number), `NewOrderIndex` (New sequence number)
	RuleChangeItems []*RuleChangeItem `json:"RuleChangeItems,omitempty" name:"RuleChangeItems"`

	// Rule direction. Values: `1` (Inbound) and `0` (Outbound)
	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`
}

type ModifyNatSequenceRulesRequest struct {
	*tchttp.BaseRequest
	
	// Rule sequence number. Values: `OrderIndex` (Original sequence number), `NewOrderIndex` (New sequence number)
	RuleChangeItems []*RuleChangeItem `json:"RuleChangeItems,omitempty" name:"RuleChangeItems"`

	// Rule direction. Values: `1` (Inbound) and `0` (Outbound)
	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`
}

func (r *ModifyNatSequenceRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNatSequenceRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleChangeItems")
	delete(f, "Direction")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyNatSequenceRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNatSequenceRulesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyNatSequenceRulesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyNatSequenceRulesResponseParams `json:"Response"`
}

func (r *ModifyNatSequenceRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNatSequenceRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPublicIPSwitchStatusRequestParams struct {
	// Public IP
	FireWallPublicIP *string `json:"FireWallPublicIP,omitempty" name:"FireWallPublicIP"`

	// Status value. 0: off; 1: on
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type ModifyPublicIPSwitchStatusRequest struct {
	*tchttp.BaseRequest
	
	// Public IP
	FireWallPublicIP *string `json:"FireWallPublicIP,omitempty" name:"FireWallPublicIP"`

	// Status value. 0: off; 1: on
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyPublicIPSwitchStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPublicIPSwitchStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FireWallPublicIP")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyPublicIPSwitchStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPublicIPSwitchStatusResponseParams struct {
	// Return message
	// Note: This field may return `null`, indicating that no valid value was found.
	ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`

	// Error code. 0: success; non-0: failed
	ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyPublicIPSwitchStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyPublicIPSwitchStatusResponseParams `json:"Response"`
}

func (r *ModifyPublicIPSwitchStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPublicIPSwitchStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyResourceGroupRequestParams struct {
	// Group ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// Group name
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// Parent group ID
	ParentId *string `json:"ParentId,omitempty" name:"ParentId"`
}

type ModifyResourceGroupRequest struct {
	*tchttp.BaseRequest
	
	// Group ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// Group name
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// Parent group ID
	ParentId *string `json:"ParentId,omitempty" name:"ParentId"`
}

func (r *ModifyResourceGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyResourceGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "GroupName")
	delete(f, "ParentId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyResourceGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyResourceGroupResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyResourceGroupResponse struct {
	*tchttp.BaseResponse
	Response *ModifyResourceGroupResponseParams `json:"Response"`
}

func (r *ModifyResourceGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyResourceGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRunSyncAssetRequestParams struct {
	// 0: edge firewall toggle; 1: VPC firewall toggle
	Type *uint64 `json:"Type,omitempty" name:"Type"`
}

type ModifyRunSyncAssetRequest struct {
	*tchttp.BaseRequest
	
	// 0: edge firewall toggle; 1: VPC firewall toggle
	Type *uint64 `json:"Type,omitempty" name:"Type"`
}

func (r *ModifyRunSyncAssetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRunSyncAssetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRunSyncAssetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRunSyncAssetResponseParams struct {
	// 0: synced successfully, 1: updating assets, 2: failed to sync by calling the API at the backend
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyRunSyncAssetResponse struct {
	*tchttp.BaseResponse
	Response *ModifyRunSyncAssetResponseParams `json:"Response"`
}

func (r *ModifyRunSyncAssetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRunSyncAssetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySecurityGroupItemRuleStatusRequestParams struct {
	// Direction. 0: outbound; 1: inbound. 1 by default
	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`

	// Toggle status. 0: off; 1: on
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// Modified priority of enterprise security group rules
	RuleSequence *uint64 `json:"RuleSequence,omitempty" name:"RuleSequence"`
}

type ModifySecurityGroupItemRuleStatusRequest struct {
	*tchttp.BaseRequest
	
	// Direction. 0: outbound; 1: inbound. 1 by default
	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`

	// Toggle status. 0: off; 1: on
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// Modified priority of enterprise security group rules
	RuleSequence *uint64 `json:"RuleSequence,omitempty" name:"RuleSequence"`
}

func (r *ModifySecurityGroupItemRuleStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySecurityGroupItemRuleStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Direction")
	delete(f, "Status")
	delete(f, "RuleSequence")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySecurityGroupItemRuleStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySecurityGroupItemRuleStatusResponseParams struct {
	// Status value. 0: modified successfully; non-0: failed to modify
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifySecurityGroupItemRuleStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifySecurityGroupItemRuleStatusResponseParams `json:"Response"`
}

func (r *ModifySecurityGroupItemRuleStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySecurityGroupItemRuleStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySecurityGroupSequenceRulesRequestParams struct {
	// Direction. 0: outbound; 1: inbound. 1 by default
	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`

	// Sorting data of enterprise security group rules
	Data []*SecurityGroupOrderIndexData `json:"Data,omitempty" name:"Data"`
}

type ModifySecurityGroupSequenceRulesRequest struct {
	*tchttp.BaseRequest
	
	// Direction. 0: outbound; 1: inbound. 1 by default
	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`

	// Sorting data of enterprise security group rules
	Data []*SecurityGroupOrderIndexData `json:"Data,omitempty" name:"Data"`
}

func (r *ModifySecurityGroupSequenceRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySecurityGroupSequenceRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Direction")
	delete(f, "Data")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySecurityGroupSequenceRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySecurityGroupSequenceRulesResponseParams struct {
	// Status value. 0: modified successfully; non-0: failed to modify
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifySecurityGroupSequenceRulesResponse struct {
	*tchttp.BaseResponse
	Response *ModifySecurityGroupSequenceRulesResponseParams `json:"Response"`
}

func (r *ModifySecurityGroupSequenceRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySecurityGroupSequenceRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySequenceRulesRequestParams struct {
	// Edge ID value
	EdgeId *string `json:"EdgeId,omitempty" name:"EdgeId"`

	// Modifies data
	Data []*SequenceData `json:"Data,omitempty" name:"Data"`

	// NAT region
	Area *string `json:"Area,omitempty" name:"Area"`

	// Direction. 0: outbound; 1: inbound
	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`
}

type ModifySequenceRulesRequest struct {
	*tchttp.BaseRequest
	
	// Edge ID value
	EdgeId *string `json:"EdgeId,omitempty" name:"EdgeId"`

	// Modifies data
	Data []*SequenceData `json:"Data,omitempty" name:"Data"`

	// NAT region
	Area *string `json:"Area,omitempty" name:"Area"`

	// Direction. 0: outbound; 1: inbound
	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`
}

func (r *ModifySequenceRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySequenceRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EdgeId")
	delete(f, "Data")
	delete(f, "Area")
	delete(f, "Direction")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySequenceRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySequenceRulesResponseParams struct {
	// 0: modified successfully; non-0: modification failed
	// Note: This field may return `null`, indicating that no valid value was found.
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifySequenceRulesResponse struct {
	*tchttp.BaseResponse
	Response *ModifySequenceRulesResponseParams `json:"Response"`
}

func (r *ModifySequenceRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySequenceRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyStorageSettingRequestParams struct {

}

type ModifyStorageSettingRequest struct {
	*tchttp.BaseRequest
	
}

func (r *ModifyStorageSettingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyStorageSettingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyStorageSettingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyStorageSettingResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyStorageSettingResponse struct {
	*tchttp.BaseResponse
	Response *ModifyStorageSettingResponseParams `json:"Response"`
}

func (r *ModifyStorageSettingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyStorageSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTableStatusRequestParams struct {
	// Edge ID between two VPCs
	EdgeId *string `json:"EdgeId,omitempty" name:"EdgeId"`

	// Status value. 1: table locked; 2: table unlocked
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// NAT region
	Area *string `json:"Area,omitempty" name:"Area"`

	// 0: outbound; 1: inbound
	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`
}

type ModifyTableStatusRequest struct {
	*tchttp.BaseRequest
	
	// Edge ID between two VPCs
	EdgeId *string `json:"EdgeId,omitempty" name:"EdgeId"`

	// Status value. 1: table locked; 2: table unlocked
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// NAT region
	Area *string `json:"Area,omitempty" name:"Area"`

	// 0: outbound; 1: inbound
	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`
}

func (r *ModifyTableStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTableStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EdgeId")
	delete(f, "Status")
	delete(f, "Area")
	delete(f, "Direction")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTableStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTableStatusResponseParams struct {
	// 0: normal; -1: abnormal
	// Note: This field may return `null`, indicating that no valid value was found.
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyTableStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyTableStatusResponseParams `json:"Response"`
}

func (r *ModifyTableStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTableStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NatFwFilter struct {
	// Filter type, e.g., instance ID
	FilterType *string `json:"FilterType,omitempty" name:"FilterType"`

	// Filtered content, separated with ","
	FilterContent *string `json:"FilterContent,omitempty" name:"FilterContent"`
}

type NatFwInstance struct {
	// NAT instance ID
	NatinsId *string `json:"NatinsId,omitempty" name:"NatinsId"`

	// NAT instance name
	NatinsName *string `json:"NatinsName,omitempty" name:"NatinsName"`

	// Instance region
	// Note: This field may return `null`, indicating that no valid value was found.
	Region *string `json:"Region,omitempty" name:"Region"`

	// 0: create new; 1: use existing
	// Note: This field may return `null`, indicating that no valid value was found.
	FwMode *int64 `json:"FwMode,omitempty" name:"FwMode"`

	// 0: normal; 1: creating
	// Note: This field may return `null`, indicating that no valid value was found.
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// NAT public IP
	// Note: This field may return `null`, indicating that no valid value was found.
	NatIp *string `json:"NatIp,omitempty" name:"NatIp"`
}

type NatInstanceInfo struct {
	// NAT instance ID
	NatinsId *string `json:"NatinsId,omitempty" name:"NatinsId"`

	// NAT instance name
	NatinsName *string `json:"NatinsName,omitempty" name:"NatinsName"`

	// Instance region
	Region *string `json:"Region,omitempty" name:"Region"`

	// 0: create new; 1: use existing
	FwMode *int64 `json:"FwMode,omitempty" name:"FwMode"`

	// Instance bandwidth (Mbps)
	BandWidth *int64 `json:"BandWidth,omitempty" name:"BandWidth"`

	// Inbound traffic peak bandwidth (bps)
	InFlowMax *int64 `json:"InFlowMax,omitempty" name:"InFlowMax"`

	// Outbound traffic peak bandwidth (bps)
	OutFlowMax *uint64 `json:"OutFlowMax,omitempty" name:"OutFlowMax"`

	// Chinese region information
	RegionZh *string `json:"RegionZh,omitempty" name:"RegionZh"`

	// Public IP array
	// Note: This field may return `null`, indicating that no valid value was found.
	EipAddress []*string `json:"EipAddress,omitempty" name:"EipAddress"`

	// Array of internal and external IPs
	// Note: This field may return `null`, indicating that no valid value was found.
	VpcIp []*string `json:"VpcIp,omitempty" name:"VpcIp"`

	// Array of subnets associated with an instance
	// Note: This field may return `null`, indicating that no valid value was found.
	Subnets []*string `json:"Subnets,omitempty" name:"Subnets"`

	// 0: normal 1: initializing
	// Note: This field may return `null`, indicating that no valid value was found.
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// Region information
	// Note: This field may return null, indicating that no valid values can be obtained.
	RegionDetail *string `json:"RegionDetail,omitempty" name:"RegionDetail"`

	// Availability zone of the instance
	// Note: This field may return null, indicating that no valid values can be obtained.
	ZoneZh *string `json:"ZoneZh,omitempty" name:"ZoneZh"`

	// Availability zone of the instance
	// Note: This field may return null, indicating that no valid values can be obtained.
	ZoneZhBak *string `json:"ZoneZhBak,omitempty" name:"ZoneZhBak"`

	// Number of used rules.
	// Note: This field may return null, indicating that no valid values can be obtained.
	RuleUsed *uint64 `json:"RuleUsed,omitempty" name:"RuleUsed"`

	// The maximum number of rules allowed in the instance.
	// Note: This field may return null, indicating that no valid values can be obtained.
	RuleMax *uint64 `json:"RuleMax,omitempty" name:"RuleMax"`
}

type NewModeItems struct {
	// VPC list for the Create New mode
	VpcList []*string `json:"VpcList,omitempty" name:"VpcList"`

	// The list of egress public EIPs bound for the Create New mode. Either Eips or AddCount is required.
	Eips []*string `json:"Eips,omitempty" name:"Eips"`

	// The number of egress public EIPs newly bound for the Create New mode. Either Eips or AddCount is required.
	AddCount *int64 `json:"AddCount,omitempty" name:"AddCount"`
}

// Predefined struct for user
type RemoveAcRuleRequestParams struct {
	// UUID of the rule, which can be obtained by querying the rule list
	RuleUuid *int64 `json:"RuleUuid,omitempty" name:"RuleUuid"`
}

type RemoveAcRuleRequest struct {
	*tchttp.BaseRequest
	
	// UUID of the rule, which can be obtained by querying the rule list
	RuleUuid *int64 `json:"RuleUuid,omitempty" name:"RuleUuid"`
}

func (r *RemoveAcRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveAcRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleUuid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RemoveAcRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveAcRuleResponseParams struct {
	// Returns the UUID of the deleted policy after the deletion is successful
	RuleUuid *int64 `json:"RuleUuid,omitempty" name:"RuleUuid"`

	// 0: operation successful; -1: operation failed
	// Note: This field may return `null`, indicating that no valid value was found.
	ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`

	// success: operation successful; failed: operation failed
	// Note: This field may return `null`, indicating that no valid value was found.
	ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RemoveAcRuleResponse struct {
	*tchttp.BaseResponse
	Response *RemoveAcRuleResponseParams `json:"Response"`
}

func (r *RemoveAcRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveAcRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveEnterpriseSecurityGroupRuleRequestParams struct {
	// UUID of the rule, which can be obtained by querying the rule list
	RuleUuid *int64 `json:"RuleUuid,omitempty" name:"RuleUuid"`

	// Type of deletion. 0: delete a single entry, and enter ID of the deleted rule for RuleUuid; 1: delete all, and enter 0 for RuleUuid
	RemoveType *int64 `json:"RemoveType,omitempty" name:"RemoveType"`
}

type RemoveEnterpriseSecurityGroupRuleRequest struct {
	*tchttp.BaseRequest
	
	// UUID of the rule, which can be obtained by querying the rule list
	RuleUuid *int64 `json:"RuleUuid,omitempty" name:"RuleUuid"`

	// Type of deletion. 0: delete a single entry, and enter ID of the deleted rule for RuleUuid; 1: delete all, and enter 0 for RuleUuid
	RemoveType *int64 `json:"RemoveType,omitempty" name:"RemoveType"`
}

func (r *RemoveEnterpriseSecurityGroupRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveEnterpriseSecurityGroupRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleUuid")
	delete(f, "RemoveType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RemoveEnterpriseSecurityGroupRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveEnterpriseSecurityGroupRuleResponseParams struct {
	// Returns the UUID of the deleted policy after the deletion is successful
	RuleUuid *int64 `json:"RuleUuid,omitempty" name:"RuleUuid"`

	// 0: operation successful; -1: operation failed
	// Note: This field may return `null`, indicating that no valid value was found.
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RemoveEnterpriseSecurityGroupRuleResponse struct {
	*tchttp.BaseResponse
	Response *RemoveEnterpriseSecurityGroupRuleResponseParams `json:"Response"`
}

func (r *RemoveEnterpriseSecurityGroupRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveEnterpriseSecurityGroupRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveNatAcRuleRequestParams struct {
	// UUIDs of the rules to delete, which can be obtained by querying the rule list. Note: If [-1] is passed in, all rules are deleted.
	RuleUuid []*int64 `json:"RuleUuid,omitempty" name:"RuleUuid"`

	// Rule direction. Valid values: 1: inbound; 0: outbound.
	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`
}

type RemoveNatAcRuleRequest struct {
	*tchttp.BaseRequest
	
	// UUIDs of the rules to delete, which can be obtained by querying the rule list. Note: If [-1] is passed in, all rules are deleted.
	RuleUuid []*int64 `json:"RuleUuid,omitempty" name:"RuleUuid"`

	// Rule direction. Valid values: 1: inbound; 0: outbound.
	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`
}

func (r *RemoveNatAcRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveNatAcRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleUuid")
	delete(f, "Direction")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RemoveNatAcRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveNatAcRuleResponseParams struct {
	// UUID list of the deleted rules.
	RuleUuid []*int64 `json:"RuleUuid,omitempty" name:"RuleUuid"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RemoveNatAcRuleResponse struct {
	*tchttp.BaseResponse
	Response *RemoveNatAcRuleResponseParams `json:"Response"`
}

func (r *RemoveNatAcRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveNatAcRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RuleChangeItem struct {
	// Original sequence number
	OrderIndex *int64 `json:"OrderIndex,omitempty" name:"OrderIndex"`

	// New sequence number
	NewOrderIndex *int64 `json:"NewOrderIndex,omitempty" name:"NewOrderIndex"`
}

type RuleInfoData struct {
	// Priority
	OrderIndex *uint64 `json:"OrderIndex,omitempty" name:"OrderIndex"`

	// Access source
	SourceIp *string `json:"SourceIp,omitempty" name:"SourceIp"`

	// Access destination
	TargetIp *string `json:"TargetIp,omitempty" name:"TargetIp"`

	// Protocol
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Policy. 0: observe; 1: block; 2: allow
	Strategy *string `json:"Strategy,omitempty" name:"Strategy"`

	// Access source type. 1: IP; 3: domain name; 4: IP address template; 5: domain name address template
	SourceType *uint64 `json:"SourceType,omitempty" name:"SourceType"`

	// Direction. 0: outbound; 1: inbound
	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`

	// Description
	Detail *string `json:"Detail,omitempty" name:"Detail"`

	// Access destination type. 1: IP, 3: domain name; 4: IP address template; 5: domain name address template
	TargetType *uint64 `json:"TargetType,omitempty" name:"TargetType"`

	// Port
	Port *string `json:"Port,omitempty" name:"Port"`

	// ID value
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// Log ID, required when an alert log is created
	LogId *string `json:"LogId,omitempty" name:"LogId"`

	// City code
	City *uint64 `json:"City,omitempty" name:"City"`

	// Country code
	Country *uint64 `json:"Country,omitempty" name:"Country"`

	// Cloud vendor. Multiple vendors are supported and separated with commas. 1: Tencent Cloud (only in Hong Kong, China and overseas); 2: Alibaba Cloud; 3: Amazon Cloud; 4: Huawei Cloud; 5: Microsoft Cloud
	CloudCode *string `json:"CloudCode,omitempty" name:"CloudCode"`

	// Indicates whether it is a region
	IsRegion *uint64 `json:"IsRegion,omitempty" name:"IsRegion"`

	// City name
	CityName *string `json:"CityName,omitempty" name:"CityName"`

	// Country name
	CountryName *string `json:"CountryName,omitempty" name:"CountryName"`
}

type ScanInfo struct {
	// Scanning result information
	ScanResultInfo *ScanResultInfo `json:"ScanResultInfo,omitempty" name:"ScanResultInfo"`

	// Scanning status. 0: scanning; 1: completed; 2: auto scanning unselected
	ScanStatus *int64 `json:"ScanStatus,omitempty" name:"ScanStatus"`

	// Progress
	ScanPercent *float64 `json:"ScanPercent,omitempty" name:"ScanPercent"`

	// Estimated completion time
	ScanTime *string `json:"ScanTime,omitempty" name:"ScanTime"`
}

type ScanResultInfo struct {
	// Number of vulnerability exploits
	LeakNum *uint64 `json:"LeakNum,omitempty" name:"LeakNum"`

	// Number of protected IPs
	IPNum *uint64 `json:"IPNum,omitempty" name:"IPNum"`

	// Number of exposed ports
	PortNum *uint64 `json:"PortNum,omitempty" name:"PortNum"`

	// Protection status
	IPStatus *bool `json:"IPStatus,omitempty" name:"IPStatus"`

	// Attack blocking status
	IdpStatus *bool `json:"IdpStatus,omitempty" name:"IdpStatus"`

	// Port blocking status
	BanStatus *bool `json:"BanStatus,omitempty" name:"BanStatus"`
}

type SecurityGroupBothWayInfo struct {
	// Priority
	// Note: This field may return `null`, indicating that no valid value was found.
	OrderIndex *uint64 `json:"OrderIndex,omitempty" name:"OrderIndex"`

	// Access source
	// Note: This field may return `null`, indicating that no valid value was found.
	SourceId *string `json:"SourceId,omitempty" name:"SourceId"`

	// Access source type. Default: 0. 0: IP; 1: VPC; 2: SUBNET; 3: CVM; 4: CLB; 5: ENI; 6: CDB; 7: Parameter template. 100: asset group
	// Note: This field may return `null`, indicating that no valid value was found.
	SourceType *uint64 `json:"SourceType,omitempty" name:"SourceType"`

	// Access destination
	// Note: This field may return `null`, indicating that no valid value was found.
	TargetId *string `json:"TargetId,omitempty" name:"TargetId"`

	// Access destination type. Default: 0. 0: IP; 1: VPC; 2: SUBNET; 3: CVM; 4: CLB; 5: ENI; 6: CDB; 7: Parameter template. 100: asset group
	// Note: This field may return `null`, indicating that no valid value was found.
	TargetType *uint64 `json:"TargetType,omitempty" name:"TargetType"`

	// Protocol
	// Note: This field may return `null`, indicating that no valid value was found.
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Destination port
	// Note: This field may return `null`, indicating that no valid value was found.
	Port *string `json:"Port,omitempty" name:"Port"`

	// Policy. 1: block; 2: allow
	// Note: This field may return `null`, indicating that no valid value was found.
	Strategy *uint64 `json:"Strategy,omitempty" name:"Strategy"`

	// Direction. 0: outbound; 1: inbound. 1 by default
	// Note: This field may return `null`, indicating that no valid value was found.
	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`

	// Region
	Region *string `json:"Region,omitempty" name:"Region"`

	// Description
	// Note: This field may return `null`, indicating that no valid value was found.
	Detail *string `json:"Detail,omitempty" name:"Detail"`

	// Toggle status. 0: off; 1: on
	// Note: This field may return `null`, indicating that no valid value was found.
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// Indicates whether the rule is normal. 0: normal; 1: abnormal
	// Note: This field may return `null`, indicating that no valid value was found.
	IsNew *uint64 `json:"IsNew,omitempty" name:"IsNew"`

	// One-way/two-way. 0: one-way; 1: two-way
	// Note: This field may return `null`, indicating that no valid value was found.
	BothWay *uint64 `json:"BothWay,omitempty" name:"BothWay"`

	// VPC ID
	// Note: This field may return `null`, indicating that no valid value was found.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Subnet ID
	// Note: This field may return `null`, indicating that no valid value was found.
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// Instance name
	// Note: This field may return `null`, indicating that no valid value was found.
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// Public IP. Multiple IPs are separated by commas.
	// Note: This field may return `null`, indicating that no valid value was found.
	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`

	// Private IP. Multiple IPs are separated by commas.
	// Note: This field may return `null`, indicating that no valid value was found.
	PrivateIp *string `json:"PrivateIp,omitempty" name:"PrivateIp"`

	// Masked address. Multiple addresses are separated by commas.
	// Note: This field may return `null`, indicating that no valid value was found.
	Cidr *string `json:"Cidr,omitempty" name:"Cidr"`

	// Port protocol template ID
	// Note: This field may return `null`, indicating that no valid value was found.
	ServiceTemplateId *string `json:"ServiceTemplateId,omitempty" name:"ServiceTemplateId"`

	// Indicates whether to use the port protocol template. 0: no; 1: yes
	ProtocolPortType *uint64 `json:"ProtocolPortType,omitempty" name:"ProtocolPortType"`
}

type SecurityGroupListData struct {
	// Priority
	OrderIndex *uint64 `json:"OrderIndex,omitempty" name:"OrderIndex"`

	// Access source
	SourceId *string `json:"SourceId,omitempty" name:"SourceId"`

	// Access source type. Default: 0. 1: VPC; 2: SUBNET; 3: CVM; 4: CLB; 5: ENI; 6: CDB; 7: Parameter template. 100: Resource group
	SourceType *uint64 `json:"SourceType,omitempty" name:"SourceType"`

	// Access destination
	TargetId *string `json:"TargetId,omitempty" name:"TargetId"`

	// Access destination type. Default: 0. 1: VPC; 2: SUBNET; 3: CVM; 4: CLB; 5: ENI; 6: CDB; 7: Parameter template; 100: resource group
	TargetType *uint64 `json:"TargetType,omitempty" name:"TargetType"`

	// Protocol
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Destination port
	Port *string `json:"Port,omitempty" name:"Port"`

	// Policy. 1: block; 2: allow
	Strategy *uint64 `json:"Strategy,omitempty" name:"Strategy"`

	// Description
	Detail *string `json:"Detail,omitempty" name:"Detail"`

	// One-way/two-way. 0: one-way; 1: two-way
	BothWay *uint64 `json:"BothWay,omitempty" name:"BothWay"`

	// Rule ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// Toggle status. 0: off; 1: on
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// Indicates whether the rule is normal. 0: normal; 1: abnormal
	IsNew *uint64 `json:"IsNew,omitempty" name:"IsNew"`

	// VPC ID
	// Note: This field may return `null`, indicating that no valid value was found.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Subnet ID
	// Note: This field may return `null`, indicating that no valid value was found.
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// Instance name
	// Note: This field may return `null`, indicating that no valid value was found.
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// Public IP. Multiple IPs are separated by commas.
	// Note: This field may return `null`, indicating that no valid value was found.
	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`

	// Private IP. Multiple IPs are separated by commas.
	// Note: This field may return `null`, indicating that no valid value was found.
	PrivateIp *string `json:"PrivateIp,omitempty" name:"PrivateIp"`

	// Masked address. Multiple addresses are separated by commas.
	// Note: This field may return `null`, indicating that no valid value was found.
	Cidr *string `json:"Cidr,omitempty" name:"Cidr"`

	// Port protocol template ID
	// Note: This field may return `null`, indicating that no valid value was found.
	ServiceTemplateId *string `json:"ServiceTemplateId,omitempty" name:"ServiceTemplateId"`

	// Two-way rules
	// Note: This field may return `null`, indicating that no valid value was found.
	BothWayInfo []*SecurityGroupBothWayInfo `json:"BothWayInfo,omitempty" name:"BothWayInfo"`

	// Direction. 0: outbound; 1: inbound. 1 by default
	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`

	// Indicates whether to use the port protocol template. 0: no; 1: yes
	ProtocolPortType *uint64 `json:"ProtocolPortType,omitempty" name:"ProtocolPortType"`
}

type SecurityGroupOrderIndexData struct {
	// Current priority of enterprise security group rules
	OrderIndex *uint64 `json:"OrderIndex,omitempty" name:"OrderIndex"`

	// New priority of enterprise security group rules
	NewOrderIndex *uint64 `json:"NewOrderIndex,omitempty" name:"NewOrderIndex"`
}

type SecurityGroupRule struct {
	// Source example:
	// net: IP/CIDR (192.168.0.2)
	// template: parameter template (ipm-dyodhpby)
	// instance: asset instance (ins-123456)
	// resourcegroup: asset group (/all groups/group 1/subgroup 1)
	// tag: resource tag ({"Key":"tag key","Value":"tag value"})
	// region: region (ap-gaungzhou)
	SourceContent *string `json:"SourceContent,omitempty" name:"SourceContent"`

	// Access source type. Valid values: net|template|instance|resourcegroup|tag|region
	SourceType *string `json:"SourceType,omitempty" name:"SourceType"`

	// Destination example:
	// net: IP/CIDR (192.168.0.2)
	// template: parameter template (ipm-dyodhpby)
	// instance: asset instance (ins-123456)
	// resourcegroup: asset group (/all groups/group 1/subgroup 1)
	// tag: resource tag ({"Key":"tag key","Value":"tag value"})
	// region: region (ap-gaungzhou)
	DestContent *string `json:"DestContent,omitempty" name:"DestContent"`

	// Access destination type. Valid values: net|template|instance|resourcegroup|tag|region
	DestType *string `json:"DestType,omitempty" name:"DestType"`

	// The action that Cloud Firewall performs on the traffic. Valid values:
	// accept: allow
	// drop: deny
	RuleAction *string `json:"RuleAction,omitempty" name:"RuleAction"`

	// Description
	Description *string `json:"Description,omitempty" name:"Description"`

	// Rule priority. -1: lowest; 1: highest
	OrderIndex *string `json:"OrderIndex,omitempty" name:"OrderIndex"`

	// Protocol. TCP/UDP/ICMP/ANY
	// Note: This field may return `null`, indicating that no valid value was found.
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// The port to apply access control rules. Valid values:
	// -1/-1: all ports
	// 80: port 80
	// Note: This field may return `null`, indicating that no valid value was found.
	Port *string `json:"Port,omitempty" name:"Port"`

	// Parameter template ID of port and protocol type; mutually exclusive with Protocol and Port
	// Note: This field may return `null`, indicating that no valid value was found.
	ServiceTemplateId *string `json:"ServiceTemplateId,omitempty" name:"ServiceTemplateId"`

	// The unique ID of the rule
	Id *string `json:"Id,omitempty" name:"Id"`

	// Rule status. true: enabled; false: disabled
	Enable *string `json:"Enable,omitempty" name:"Enable"`
}

type SequenceData struct {
	// Rule ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// Rule priority before change
	OrderIndex *uint64 `json:"OrderIndex,omitempty" name:"OrderIndex"`

	// Rule priority after change
	NewOrderIndex *uint64 `json:"NewOrderIndex,omitempty" name:"NewOrderIndex"`
}

// Predefined struct for user
type SetNatFwDnatRuleRequestParams struct {
	// 0: Create new; 1: Use existing
	Mode *uint64 `json:"Mode,omitempty" name:"Mode"`

	// Operation type. Valid values: add, del, and modify.
	OperationType *string `json:"OperationType,omitempty" name:"OperationType"`

	// Firewall instance ID. This field is required.
	CfwInstance *string `json:"CfwInstance,omitempty" name:"CfwInstance"`

	// List of added/deleted DNAT rules
	AddOrDelDnatRules []*CfwNatDnatRule `json:"AddOrDelDnatRules,omitempty" name:"AddOrDelDnatRules"`

	// Original DNAT rule before change
	OriginDnat *CfwNatDnatRule `json:"OriginDnat,omitempty" name:"OriginDnat"`

	// New DNAT rule after change
	NewDnat *CfwNatDnatRule `json:"NewDnat,omitempty" name:"NewDnat"`
}

type SetNatFwDnatRuleRequest struct {
	*tchttp.BaseRequest
	
	// 0: Create new; 1: Use existing
	Mode *uint64 `json:"Mode,omitempty" name:"Mode"`

	// Operation type. Valid values: add, del, and modify.
	OperationType *string `json:"OperationType,omitempty" name:"OperationType"`

	// Firewall instance ID. This field is required.
	CfwInstance *string `json:"CfwInstance,omitempty" name:"CfwInstance"`

	// List of added/deleted DNAT rules
	AddOrDelDnatRules []*CfwNatDnatRule `json:"AddOrDelDnatRules,omitempty" name:"AddOrDelDnatRules"`

	// Original DNAT rule before change
	OriginDnat *CfwNatDnatRule `json:"OriginDnat,omitempty" name:"OriginDnat"`

	// New DNAT rule after change
	NewDnat *CfwNatDnatRule `json:"NewDnat,omitempty" name:"NewDnat"`
}

func (r *SetNatFwDnatRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetNatFwDnatRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Mode")
	delete(f, "OperationType")
	delete(f, "CfwInstance")
	delete(f, "AddOrDelDnatRules")
	delete(f, "OriginDnat")
	delete(f, "NewDnat")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetNatFwDnatRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetNatFwDnatRuleResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SetNatFwDnatRuleResponse struct {
	*tchttp.BaseResponse
	Response *SetNatFwDnatRuleResponseParams `json:"Response"`
}

func (r *SetNatFwDnatRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetNatFwDnatRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetNatFwEipRequestParams struct {
	// bind: bind EIP; unbind: unbind EIP; newAdd: add firewall EIP
	OperationType *string `json:"OperationType,omitempty" name:"OperationType"`

	// Firewall instance ID
	CfwInstance *string `json:"CfwInstance,omitempty" name:"CfwInstance"`

	// This field is required when OperationType is "bind" or "unbind".
	EipList []*string `json:"EipList,omitempty" name:"EipList"`
}

type SetNatFwEipRequest struct {
	*tchttp.BaseRequest
	
	// bind: bind EIP; unbind: unbind EIP; newAdd: add firewall EIP
	OperationType *string `json:"OperationType,omitempty" name:"OperationType"`

	// Firewall instance ID
	CfwInstance *string `json:"CfwInstance,omitempty" name:"CfwInstance"`

	// This field is required when OperationType is "bind" or "unbind".
	EipList []*string `json:"EipList,omitempty" name:"EipList"`
}

func (r *SetNatFwEipRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetNatFwEipRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OperationType")
	delete(f, "CfwInstance")
	delete(f, "EipList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetNatFwEipRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetNatFwEipResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SetNatFwEipResponse struct {
	*tchttp.BaseResponse
	Response *SetNatFwEipResponseParams `json:"Response"`
}

func (r *SetNatFwEipResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetNatFwEipResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StaticInfo struct {
	// Number
	Num *int64 `json:"Num,omitempty" name:"Num"`

	// Port
	Port *string `json:"Port,omitempty" name:"Port"`

	// IP
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// Address
	Address *string `json:"Address,omitempty" name:"Address"`

	// Asset ID
	InsID *string `json:"InsID,omitempty" name:"InsID"`

	// Asset name
	InsName *string `json:"InsName,omitempty" name:"InsName"`
}

// Predefined struct for user
type StopSecurityGroupRuleDispatchRequestParams struct {
	// Stops all if set to 1
	StopType *int64 `json:"StopType,omitempty" name:"StopType"`
}

type StopSecurityGroupRuleDispatchRequest struct {
	*tchttp.BaseRequest
	
	// Stops all if set to 1
	StopType *int64 `json:"StopType,omitempty" name:"StopType"`
}

func (r *StopSecurityGroupRuleDispatchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopSecurityGroupRuleDispatchRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StopType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopSecurityGroupRuleDispatchRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopSecurityGroupRuleDispatchResponseParams struct {
	// true: operation successful; false: error
	// Note: This field may return `null`, indicating that no valid value was found.
	Status *bool `json:"Status,omitempty" name:"Status"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type StopSecurityGroupRuleDispatchResponse struct {
	*tchttp.BaseResponse
	Response *StopSecurityGroupRuleDispatchResponseParams `json:"Response"`
}

func (r *StopSecurityGroupRuleDispatchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopSecurityGroupRuleDispatchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SwitchListsData struct {
	// Public IP
	// Note: This field may return `null`, indicating that no valid value was found.
	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`

	// Private IP
	// Note: This field may return `null`, indicating that no valid value was found.
	IntranetIp *string `json:"IntranetIp,omitempty" name:"IntranetIp"`

	// Instance name
	// Note: This field may return `null`, indicating that no valid value was found.
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// Instance ID
	// Note: This field may return `null`, indicating that no valid value was found.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Asset type
	AssetType *string `json:"AssetType,omitempty" name:"AssetType"`

	// Region
	// Note: This field may return `null`, indicating that no valid value was found.
	Area *string `json:"Area,omitempty" name:"Area"`

	// Firewall toggle
	Switch *int64 `json:"Switch,omitempty" name:"Switch"`

	// ID value
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// Public IP type
	// Note: This field may return `null`, indicating that no valid value was found.
	PublicIpType *uint64 `json:"PublicIpType,omitempty" name:"PublicIpType"`

	// Number of risky ports
	// Note: This field may return `null`, indicating that no valid value was found.
	PortTimes *uint64 `json:"PortTimes,omitempty" name:"PortTimes"`

	// Last scan time
	// Note: This field may return `null`, indicating that no valid value was found.
	LastTime *string `json:"LastTime,omitempty" name:"LastTime"`

	// Scan mode
	// Note: This field may return `null`, indicating that no valid value was found.
	ScanMode *string `json:"ScanMode,omitempty" name:"ScanMode"`

	// Scan status
	// Note: This field may return `null`, indicating that no valid value was found.
	ScanStatus *uint64 `json:"ScanStatus,omitempty" name:"ScanStatus"`
}

type TLogInfo struct {
	// Compromised servers
	OutNum *int64 `json:"OutNum,omitempty" name:"OutNum"`

	// Unhandled alerts
	HandleNum *int64 `json:"HandleNum,omitempty" name:"HandleNum"`

	// Vulnerability attacks
	VulNum *int64 `json:"VulNum,omitempty" name:"VulNum"`

	// Detected networks
	NetworkNum *int64 `json:"NetworkNum,omitempty" name:"NetworkNum"`

	// Blocklist
	BanNum *int64 `json:"BanNum,omitempty" name:"BanNum"`

	// Brute force attacks
	BruteForceNum *int64 `json:"BruteForceNum,omitempty" name:"BruteForceNum"`
}

type UnHandleEvent struct {
	// Unhandled event type
	EventTableListStruct []*UnHandleEventDetail `json:"EventTableListStruct,omitempty" name:"EventTableListStruct"`

	// 1: yes; 0: no
	BaseLineUser *uint64 `json:"BaseLineUser,omitempty" name:"BaseLineUser"`

	// 1: on; 0: off
	BaseLineInSwitch *uint64 `json:"BaseLineInSwitch,omitempty" name:"BaseLineInSwitch"`

	// 1: on; 0: off
	BaseLineOutSwitch *uint64 `json:"BaseLineOutSwitch,omitempty" name:"BaseLineOutSwitch"`

	// Number of inter-VPC firewall instances
	// Note: This field may return `null`, indicating that no valid value was found.
	VpcFwCount *uint64 `json:"VpcFwCount,omitempty" name:"VpcFwCount"`
}

type UnHandleEventDetail struct {
	// Security event name
	EventName *string `json:"EventName,omitempty" name:"EventName"`

	// Number of unhandled events
	Total *uint64 `json:"Total,omitempty" name:"Total"`
}

type VpcDnsInfo struct {
	// VPC ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// VPC name
	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`

	// NAT firewall mode. 0: Create new; 1: Use existing
	FwMode *int64 `json:"FwMode,omitempty" name:"FwMode"`

	// VPC IPv4 CIDR block (Classless Inter-Domain Routing)
	VpcIpv4Cidr *string `json:"VpcIpv4Cidr,omitempty" name:"VpcIpv4Cidr"`

	// Public EIP, which is the firewall DNS resolution address
	DNSEip *string `json:"DNSEip,omitempty" name:"DNSEip"`

	// NAT gateway ID
	// Note: This field may return `null`, indicating that no valid value was found.
	NatInsId *string `json:"NatInsId,omitempty" name:"NatInsId"`

	// NAT gateway name
	// Note: This field may return `null`, indicating that no valid value was found.
	NatInsName *string `json:"NatInsName,omitempty" name:"NatInsName"`

	// 0: off; 1: on
	SwitchStatus *int64 `json:"SwitchStatus,omitempty" name:"SwitchStatus"`
}