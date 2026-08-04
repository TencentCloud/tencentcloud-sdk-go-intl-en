// Copyright (c) 2017-2025 Tencent. All Rights Reserved.
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
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

type AcListsData struct {
	// Rule ID
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// Access source
	// Note: This field may return `null`, indicating that no valid value was found.
	SourceIp *string `json:"SourceIp,omitnil,omitempty" name:"SourceIp"`

	// Access destination
	// Note: This field may return `null`, indicating that no valid value was found.
	TargetIp *string `json:"TargetIp,omitnil,omitempty" name:"TargetIp"`

	// Protocol
	// Note: This field may return `null`, indicating that no valid value was found.
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// Port
	// Note: This field may return `null`, indicating that no valid value was found.
	Port *string `json:"Port,omitnil,omitempty" name:"Port"`

	// Policy
	// Note: This field may return `null`, indicating that no valid value was found.
	Strategy *uint64 `json:"Strategy,omitnil,omitempty" name:"Strategy"`

	// Description
	// Note: This field may return `null`, indicating that no valid value was found.
	Detail *string `json:"Detail,omitnil,omitempty" name:"Detail"`

	// Hit count
	Count *uint64 `json:"Count,omitnil,omitempty" name:"Count"`

	// Priority
	OrderIndex *uint64 `json:"OrderIndex,omitnil,omitempty" name:"OrderIndex"`

	// Alert rule ID
	// Note: This field may return `null`, indicating that no valid value was found.
	LogId *string `json:"LogId,omitnil,omitempty" name:"LogId"`
}

// Predefined struct for user
type AddAcRuleRequestParams struct {
	// -1: lowest priority; 1: highest priority
	OrderIndex *string `json:"OrderIndex,omitnil,omitempty" name:"OrderIndex"`

	// The action that Cloud Firewall performs on the traffic. Valid values:
	// accept: allow
	// drop: deny
	// log: observe
	RuleAction *string `json:"RuleAction,omitnil,omitempty" name:"RuleAction"`

	// The traffic direction for access control rules. Valid values:
	// in: incoming traffic access control
	// out: outgoing traffic access control
	Direction *string `json:"Direction,omitnil,omitempty" name:"Direction"`

	// The description of access control rules.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// The type of source address in access control rules. Valid values:
	// net: source IP or range (IP or CIDR)
	// location: source region
	// template: CFW address template
	// instance: instance ID
	// vendor: Cloud vendor
	SourceType *string `json:"SourceType,omitnil,omitempty" name:"SourceType"`

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
	SourceContent *string `json:"SourceContent,omitnil,omitempty" name:"SourceContent"`

	// The type of destination address in access control rules. Valid values:
	// net: destination IP or range (IP or CIDR)
	// location: source region
	// template: CFW address template
	// instance: instance ID
	// vendor: Cloud vendor
	// domain: Domain name or IP.
	DestType *string `json:"DestType,omitnil,omitempty" name:"DestType"`

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
	DestContent *string `json:"DestContent,omitnil,omitempty" name:"DestContent"`

	// The port to apply access control rules. Valid values:
	// -1/-1: all ports
	// 80,443: 80 or 443
	Port *string `json:"Port,omitnil,omitempty" name:"Port"`

	// The protocol type of traffic in access control rules. Valid value: TCP. Only TCP is supported for edge firewall rules. If this parameter is not specified, it defaults to TCP.
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// The Layer 7 protocol. Valid values:
	// HTTP/HTTPS
	// TLS/SSL
	ApplicationName *string `json:"ApplicationName,omitnil,omitempty" name:"ApplicationName"`

	// Indicates whether to enable the rules. Default: enable. Valid values:
	// true: enable; false: disable
	Enable *string `json:"Enable,omitnil,omitempty" name:"Enable"`
}

type AddAcRuleRequest struct {
	*tchttp.BaseRequest
	
	// -1: lowest priority; 1: highest priority
	OrderIndex *string `json:"OrderIndex,omitnil,omitempty" name:"OrderIndex"`

	// The action that Cloud Firewall performs on the traffic. Valid values:
	// accept: allow
	// drop: deny
	// log: observe
	RuleAction *string `json:"RuleAction,omitnil,omitempty" name:"RuleAction"`

	// The traffic direction for access control rules. Valid values:
	// in: incoming traffic access control
	// out: outgoing traffic access control
	Direction *string `json:"Direction,omitnil,omitempty" name:"Direction"`

	// The description of access control rules.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// The type of source address in access control rules. Valid values:
	// net: source IP or range (IP or CIDR)
	// location: source region
	// template: CFW address template
	// instance: instance ID
	// vendor: Cloud vendor
	SourceType *string `json:"SourceType,omitnil,omitempty" name:"SourceType"`

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
	SourceContent *string `json:"SourceContent,omitnil,omitempty" name:"SourceContent"`

	// The type of destination address in access control rules. Valid values:
	// net: destination IP or range (IP or CIDR)
	// location: source region
	// template: CFW address template
	// instance: instance ID
	// vendor: Cloud vendor
	// domain: Domain name or IP.
	DestType *string `json:"DestType,omitnil,omitempty" name:"DestType"`

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
	DestContent *string `json:"DestContent,omitnil,omitempty" name:"DestContent"`

	// The port to apply access control rules. Valid values:
	// -1/-1: all ports
	// 80,443: 80 or 443
	Port *string `json:"Port,omitnil,omitempty" name:"Port"`

	// The protocol type of traffic in access control rules. Valid value: TCP. Only TCP is supported for edge firewall rules. If this parameter is not specified, it defaults to TCP.
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// The Layer 7 protocol. Valid values:
	// HTTP/HTTPS
	// TLS/SSL
	ApplicationName *string `json:"ApplicationName,omitnil,omitempty" name:"ApplicationName"`

	// Indicates whether to enable the rules. Default: enable. Valid values:
	// true: enable; false: disable
	Enable *string `json:"Enable,omitnil,omitempty" name:"Enable"`
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
	RuleUuid *int64 `json:"RuleUuid,omitnil,omitempty" name:"RuleUuid"`

	// 0: operation successful; -1: operation failed
	ReturnCode *int64 `json:"ReturnCode,omitnil,omitempty" name:"ReturnCode"`

	// success: operation successful; failed: operation failed
	ReturnMsg *string `json:"ReturnMsg,omitnil,omitempty" name:"ReturnMsg"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type AddAclRuleRequestParams struct {
	// The list of Internet boundary rules to be added cannot be empty. Each rule must meet the requirements for direction, source and target, action, scope, protocol port, and template restrictions. The entire request must also comply with rule quota and effective rule count limitations. Account-related values must come from read-only queries: for address templates, call DescribeAddressTemplateList, filter the request with TemplateType=1 or 5, and confirm that the returned Data[].Type is 1 or 5. Write Data[].Uuid (with the mb_ prefix) to the corresponding Content, and do not use Data[].TemplateId (with the ip-/dm- prefix). For protocol port templates, filter the request with TemplateType=6, and write Data[].TemplateId (with the pp- prefix) to ParamTemplateId. For asset instances, call DescribeCfwAssets, parse the returned results, and use assets[].instance_id. For asset groups, call DescribeResourceGroupNew, pass QueryType=resource, GroupId="0", ShowType=all, parse the returned results, and use GroupId. For resource tags, pass QueryType=tag, skip the "all assets" root node, and construct JSON using the first-level node GroupName as the Key and the selected second-level sub-node GroupName as the Value, without writing GroupId. For regions, call DescribeAclRegInfo: for Scope=serial, pass FwType=["SERIAL"]; for Scope=side, pass FwType=["BYPASS"]; for Scope=all, pass both items simultaneously, and use Data[].RegionCode. Do not use display names or manually concatenate values. The range for overwrite import is determined solely by the Direction of the first rule.
	Rules []*CreateRuleItem `json:"Rules,omitnil,omitempty" name:"Rules"`

	// <p>AI operation source</p><p>Enumeration value:</p><ul><li>console: console source value</li><li>wechat: WeChat</li></ul>
	CfwAiAgentOperationSource *string `json:"CfwAiAgentOperationSource,omitnil,omitempty" name:"CfwAiAgentOperationSource"`

	// Add method. Omit or an empty string means ordinary addition; insert_rule means adding new at a specified position; batch_import means non-overwrite batch import; batch_import_cover means overwrite import, which deletes the existing operation rule corresponding to the first rule's Direction and then adds Rules. Deleted rules will not be restored if addition fails, with extremely high risk. The coverage area is only determined by the first rule. The caller must ensure the Directions of the Rules match. Only the above values are supported.
	From *string `json:"From,omitnil,omitempty" name:"From"`
}

type AddAclRuleRequest struct {
	*tchttp.BaseRequest
	
	// The list of Internet boundary rules to be added cannot be empty. Each rule must meet the requirements for direction, source and target, action, scope, protocol port, and template restrictions. The entire request must also comply with rule quota and effective rule count limitations. Account-related values must come from read-only queries: for address templates, call DescribeAddressTemplateList, filter the request with TemplateType=1 or 5, and confirm that the returned Data[].Type is 1 or 5. Write Data[].Uuid (with the mb_ prefix) to the corresponding Content, and do not use Data[].TemplateId (with the ip-/dm- prefix). For protocol port templates, filter the request with TemplateType=6, and write Data[].TemplateId (with the pp- prefix) to ParamTemplateId. For asset instances, call DescribeCfwAssets, parse the returned results, and use assets[].instance_id. For asset groups, call DescribeResourceGroupNew, pass QueryType=resource, GroupId="0", ShowType=all, parse the returned results, and use GroupId. For resource tags, pass QueryType=tag, skip the "all assets" root node, and construct JSON using the first-level node GroupName as the Key and the selected second-level sub-node GroupName as the Value, without writing GroupId. For regions, call DescribeAclRegInfo: for Scope=serial, pass FwType=["SERIAL"]; for Scope=side, pass FwType=["BYPASS"]; for Scope=all, pass both items simultaneously, and use Data[].RegionCode. Do not use display names or manually concatenate values. The range for overwrite import is determined solely by the Direction of the first rule.
	Rules []*CreateRuleItem `json:"Rules,omitnil,omitempty" name:"Rules"`

	// <p>AI operation source</p><p>Enumeration value:</p><ul><li>console: console source value</li><li>wechat: WeChat</li></ul>
	CfwAiAgentOperationSource *string `json:"CfwAiAgentOperationSource,omitnil,omitempty" name:"CfwAiAgentOperationSource"`

	// Add method. Omit or an empty string means ordinary addition; insert_rule means adding new at a specified position; batch_import means non-overwrite batch import; batch_import_cover means overwrite import, which deletes the existing operation rule corresponding to the first rule's Direction and then adds Rules. Deleted rules will not be restored if addition fails, with extremely high risk. The coverage area is only determined by the first rule. The caller must ensure the Directions of the Rules match. Only the above values are supported.
	From *string `json:"From,omitnil,omitempty" name:"From"`
}

func (r *AddAclRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddAclRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Rules")
	delete(f, "CfwAiAgentOperationSource")
	delete(f, "From")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddAclRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddAclRuleResponseParams struct {
	// ID list of the added rules, in the same sequence as Rules.
	RuleUuid []*int64 `json:"RuleUuid,omitnil,omitempty" name:"RuleUuid"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AddAclRuleResponse struct {
	*tchttp.BaseResponse
	Response *AddAclRuleResponseParams `json:"Response"`
}

func (r *AddAclRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddAclRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddEnterpriseSecurityGroupRulesRequestParams struct {
	// Creates rule data
	Data []*SecurityGroupRule `json:"Data,omitnil,omitempty" name:"Data"`

	// Adding type. 0: add to the end; 1: add to the front; 2: insert. Default: 0
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// An identifier to ensure the idempotency of the request. The value of the ClientToken parameter is a unique string that is generated by your client and can contain up to 64 ASCII characters in length.
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// Indicates whether to delay publishing. 1: delay; other values: do not delay
	IsDelay *uint64 `json:"IsDelay,omitnil,omitempty" name:"IsDelay"`
}

type AddEnterpriseSecurityGroupRulesRequest struct {
	*tchttp.BaseRequest
	
	// Creates rule data
	Data []*SecurityGroupRule `json:"Data,omitnil,omitempty" name:"Data"`

	// Adding type. 0: add to the end; 1: add to the front; 2: insert. Default: 0
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// An identifier to ensure the idempotency of the request. The value of the ClientToken parameter is a unique string that is generated by your client and can contain up to 64 ASCII characters in length.
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// Indicates whether to delay publishing. 1: delay; other values: do not delay
	IsDelay *uint64 `json:"IsDelay,omitnil,omitempty" name:"IsDelay"`
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
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Rules []*CreateNatRuleItem `json:"Rules,omitnil,omitempty" name:"Rules"`

	// Source of the rules to be added. Generally, this parameter is not used. The value insert_rule indicates that rules in the specified location are inserted, and the value batch_import indicates that rules are imported in batches. If the parameter is left empty, rules defined in the API request are added.
	From *string `json:"From,omitnil,omitempty" name:"From"`
}

type AddNatAcRuleRequest struct {
	*tchttp.BaseRequest
	
	// NAT access control rules to be added.
	Rules []*CreateNatRuleItem `json:"Rules,omitnil,omitempty" name:"Rules"`

	// Source of the rules to be added. Generally, this parameter is not used. The value insert_rule indicates that rules in the specified location are inserted, and the value batch_import indicates that rules are imported in batches. If the parameter is left empty, rules defined in the API request are added.
	From *string `json:"From,omitnil,omitempty" name:"From"`
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
	RuleUuid []*int64 `json:"RuleUuid,omitnil,omitempty" name:"RuleUuid"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Region
	ZoneEng *string `json:"ZoneEng,omitnil,omitempty" name:"ZoneEng"`
}

type AssociatedInstanceInfo struct {
	// Instance ID
	// Note: This field may return `null`, indicating that no valid value was found.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Instance name
	// Note: This field may return `null`, indicating that no valid value was found.
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Instance type. 3: CVM instance; 4: CLB instance; 5: ENI instance; 6: Cloud database
	// Note: This field may return `null`, indicating that no valid value was found.
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// VPC ID
	// Note: This field may return `null`, indicating that no valid value was found.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// VPC name
	// Note: This field may return `null`, indicating that no valid value was found.
	VpcName *string `json:"VpcName,omitnil,omitempty" name:"VpcName"`

	// Public IP
	// Note: This field may return `null`, indicating that no valid value was found.
	PublicIp *string `json:"PublicIp,omitnil,omitempty" name:"PublicIp"`

	// Private IP
	// Note: This field may return `null`, indicating that no valid value was found.
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// The number of associated security groups
	// Note: This field may return `null`, indicating that no valid value was found.
	SecurityGroupCount *uint64 `json:"SecurityGroupCount,omitnil,omitempty" name:"SecurityGroupCount"`
}

type BetaInfoByACL struct {
	// Task ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Task name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// Last execution time
	// Note: This field may return null, indicating that no valid values can be obtained.
	LastTime *string `json:"LastTime,omitnil,omitempty" name:"LastTime"`
}

type BlockIgnoreRule struct {
	// Domain name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// Rule IP.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Ioc *string `json:"Ioc,omitnil,omitempty" name:"Ioc"`

	// Threat level.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// Source event name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	EventName *string `json:"EventName,omitnil,omitempty" name:"EventName"`

	// Direction. Valid values: 0: outbound; 1: inbound.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Direction *int64 `json:"Direction,omitnil,omitempty" name:"Direction"`

	// Protocol.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// Address.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Address *string `json:"Address,omitnil,omitempty" name:"Address"`

	// Rule type. Valid values: 1: block; 2: allow.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Action *int64 `json:"Action,omitnil,omitempty" name:"Action"`

	// Time when a rule starts to take effect.
	// Note: This field may return null, indicating that no valid values can be obtained.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Time when a rule expires.
	// Note: This field may return null, indicating that no valid values can be obtained.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Reason for ignoring.
	// Note: This field may return null, indicating that no valid values can be obtained.
	IgnoreReason *string `json:"IgnoreReason,omitnil,omitempty" name:"IgnoreReason"`

	// Security event source.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// Rule ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	UniqueId *string `json:"UniqueId,omitnil,omitempty" name:"UniqueId"`

	// Number of rule matching times.
	// Note: This field may return null, indicating that no valid values can be obtained.
	MatchTimes *int64 `json:"MatchTimes,omitnil,omitempty" name:"MatchTimes"`

	// Country.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Country *string `json:"Country,omitnil,omitempty" name:"Country"`

	// Remarks
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`
}

type CfwNatDnatRule struct {
	// Network protocol. Valid values: TCP or UDP.
	IpProtocol *string `json:"IpProtocol,omitnil,omitempty" name:"IpProtocol"`

	// Elastic IP.
	PublicIpAddress *string `json:"PublicIpAddress,omitnil,omitempty" name:"PublicIpAddress"`

	// Public port.
	PublicPort *int64 `json:"PublicPort,omitnil,omitempty" name:"PublicPort"`

	// Private address.
	PrivateIpAddress *string `json:"PrivateIpAddress,omitnil,omitempty" name:"PrivateIpAddress"`

	// Private port.
	PrivatePort *int64 `json:"PrivatePort,omitnil,omitempty" name:"PrivatePort"`

	// The description of NAT firewall forwarding rules.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type CfwStatusMonitorFilter struct {
	// Filter field name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Filter value list, up to 10.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`

	// Operator type, optional; only supported for backend permission types.
	// Note: This field may return null, indicating that no valid values can be obtained.
	OperatorType *int64 `json:"OperatorType,omitnil,omitempty" name:"OperatorType"`
}

type Column struct {
	// Column name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Column attribute
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type CommonFilter struct {
	// Search key.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Search values.
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`

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
	OperatorType *int64 `json:"OperatorType,omitnil,omitempty" name:"OperatorType"`
}

// Predefined struct for user
type CreateAcRulesRequestParams struct {
	// Creates rule data
	Data []*RuleInfoData `json:"Data,omitnil,omitempty" name:"Data"`

	// 0: add (default); 1: insert
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// Edge ID
	EdgeId *string `json:"EdgeId,omitnil,omitempty" name:"EdgeId"`

	// Access control rule status
	Enable *int64 `json:"Enable,omitnil,omitempty" name:"Enable"`

	// 0: add; 1: overwrite
	Overwrite *uint64 `json:"Overwrite,omitnil,omitempty" name:"Overwrite"`

	// NAT instance ID, required when the parameter Area exists
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// portScan: from port scanning; patchImport: from batch import
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// NAT region
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`
}

type CreateAcRulesRequest struct {
	*tchttp.BaseRequest
	
	// Creates rule data
	Data []*RuleInfoData `json:"Data,omitnil,omitempty" name:"Data"`

	// 0: add (default); 1: insert
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// Edge ID
	EdgeId *string `json:"EdgeId,omitnil,omitempty" name:"EdgeId"`

	// Access control rule status
	Enable *int64 `json:"Enable,omitnil,omitempty" name:"Enable"`

	// 0: add; 1: overwrite
	Overwrite *uint64 `json:"Overwrite,omitnil,omitempty" name:"Overwrite"`

	// NAT instance ID, required when the parameter Area exists
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// portScan: from port scanning; patchImport: from batch import
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// NAT region
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`
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
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Returns redundant information
	// Note: This field may return `null`, indicating that no valid value was found.
	Info *string `json:"Info,omitnil,omitempty" name:"Info"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Bandwidth
	Width *int64 `json:"Width,omitnil,omitempty" name:"Width"`

	// Mode. 1: use existing; 0: create new
	Mode *int64 `json:"Mode,omitnil,omitempty" name:"Mode"`

	// Parameter passed for the Create New mode. Either NewModeItems or NatgwList is required.
	NewModeItems *NewModeItems `json:"NewModeItems,omitnil,omitempty" name:"NewModeItems"`

	// NAT gateway list for the Using Existing mode. Either NewModeItems or NatgwList is required.
	NatGwList []*string `json:"NatGwList,omitnil,omitempty" name:"NatGwList"`

	// Primary zone. The default zone is selected if it is empty.
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Secondary zone. The default zone is selected if it is empty.
	ZoneBak *string `json:"ZoneBak,omitnil,omitempty" name:"ZoneBak"`

	// Remote disaster recovery. 1: enable; 0: disable; empty: disable by default
	CrossAZone *int64 `json:"CrossAZone,omitnil,omitempty" name:"CrossAZone"`

	// IP range of the firewall
	FwCidrInfo *FwCidrInfo `json:"FwCidrInfo,omitnil,omitempty" name:"FwCidrInfo"`
}

type CreateNatFwInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Firewall instance name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Bandwidth
	Width *int64 `json:"Width,omitnil,omitempty" name:"Width"`

	// Mode. 1: use existing; 0: create new
	Mode *int64 `json:"Mode,omitnil,omitempty" name:"Mode"`

	// Parameter passed for the Create New mode. Either NewModeItems or NatgwList is required.
	NewModeItems *NewModeItems `json:"NewModeItems,omitnil,omitempty" name:"NewModeItems"`

	// NAT gateway list for the Using Existing mode. Either NewModeItems or NatgwList is required.
	NatGwList []*string `json:"NatGwList,omitnil,omitempty" name:"NatGwList"`

	// Primary zone. The default zone is selected if it is empty.
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Secondary zone. The default zone is selected if it is empty.
	ZoneBak *string `json:"ZoneBak,omitnil,omitempty" name:"ZoneBak"`

	// Remote disaster recovery. 1: enable; 0: disable; empty: disable by default
	CrossAZone *int64 `json:"CrossAZone,omitnil,omitempty" name:"CrossAZone"`

	// IP range of the firewall
	FwCidrInfo *FwCidrInfo `json:"FwCidrInfo,omitnil,omitempty" name:"FwCidrInfo"`
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
	CfwInsId *string `json:"CfwInsId,omitnil,omitempty" name:"CfwInsId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Bandwidth
	Width *int64 `json:"Width,omitnil,omitempty" name:"Width"`

	// Mode. 1: use existing; 0: create new
	Mode *int64 `json:"Mode,omitnil,omitempty" name:"Mode"`

	// Parameter passed for the Create New mode. Either NewModeItems or NatgwList is required.
	NewModeItems *NewModeItems `json:"NewModeItems,omitnil,omitempty" name:"NewModeItems"`

	// NAT gateway list for the Using Existing mode. Either NewModeItems or NatgwList is required.
	NatGwList []*string `json:"NatGwList,omitnil,omitempty" name:"NatGwList"`

	// Primary zone. The default zone is selected if it is empty.
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Secondary zone. The default zone is selected if it is empty.
	ZoneBak *string `json:"ZoneBak,omitnil,omitempty" name:"ZoneBak"`

	// Remote disaster recovery. 1: enable; 0: disable; empty: disable by default
	CrossAZone *int64 `json:"CrossAZone,omitnil,omitempty" name:"CrossAZone"`

	// 0: not create; 1: create
	IsCreateDomain *int64 `json:"IsCreateDomain,omitnil,omitempty" name:"IsCreateDomain"`

	// Required for creating a domain name
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// IP range of the firewall
	FwCidrInfo *FwCidrInfo `json:"FwCidrInfo,omitnil,omitempty" name:"FwCidrInfo"`
}

type CreateNatFwInstanceWithDomainRequest struct {
	*tchttp.BaseRequest
	
	// Firewall instance name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Bandwidth
	Width *int64 `json:"Width,omitnil,omitempty" name:"Width"`

	// Mode. 1: use existing; 0: create new
	Mode *int64 `json:"Mode,omitnil,omitempty" name:"Mode"`

	// Parameter passed for the Create New mode. Either NewModeItems or NatgwList is required.
	NewModeItems *NewModeItems `json:"NewModeItems,omitnil,omitempty" name:"NewModeItems"`

	// NAT gateway list for the Using Existing mode. Either NewModeItems or NatgwList is required.
	NatGwList []*string `json:"NatGwList,omitnil,omitempty" name:"NatGwList"`

	// Primary zone. The default zone is selected if it is empty.
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Secondary zone. The default zone is selected if it is empty.
	ZoneBak *string `json:"ZoneBak,omitnil,omitempty" name:"ZoneBak"`

	// Remote disaster recovery. 1: enable; 0: disable; empty: disable by default
	CrossAZone *int64 `json:"CrossAZone,omitnil,omitempty" name:"CrossAZone"`

	// 0: not create; 1: create
	IsCreateDomain *int64 `json:"IsCreateDomain,omitnil,omitempty" name:"IsCreateDomain"`

	// Required for creating a domain name
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// IP range of the firewall
	FwCidrInfo *FwCidrInfo `json:"FwCidrInfo,omitnil,omitempty" name:"FwCidrInfo"`
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
	CfwInsId *string `json:"CfwInsId,omitnil,omitempty" name:"CfwInsId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	SourceContent *string `json:"SourceContent,omitnil,omitempty" name:"SourceContent"`

	// Access source type. Values for inbound rules: `ip`, `net`, `template`, and `location`. Values for outbound rules: `ip`, `net`, `template`, `instance`, `group`, and `tag`.
	SourceType *string `json:"SourceType,omitnil,omitempty" name:"SourceType"`

	// Access target. Example: `net: IP/CIDR(192.168.0.2); domain: domain name rule, e.g., *.qq.com
	TargetContent *string `json:"TargetContent,omitnil,omitempty" name:"TargetContent"`

	// Access target type. Values for inbound rules: `ip`, `net`, `template`, `instance`, `group`, and `tag`. Values for outbound rules: `ip`, `net`, `domain`, `template`, and `location`.
	TargetType *string `json:"TargetType,omitnil,omitempty" name:"TargetType"`

	// Protocol. Values: `TCP`, `UDP`, `ICMP`, `ANY`, `HTTP`, `HTTPS`, `HTTP/HTTPS`, `SMTP`, `SMTPS`, `SMTP/SMTPS`, `FTP`, and `DNS`.
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// Specify how the CFW instance deals with the traffic hit the access control rule. Values: `accept` (allow), `drop` (reject), and `log` (observe).
	RuleAction *string `json:"RuleAction,omitnil,omitempty" name:"RuleAction"`

	// The port of the access control rule. Values: `-1/-1` (all ports) and `80` (Port 80)
	Port *string `json:"Port,omitnil,omitempty" name:"Port"`

	// Rule direction. Values: `1` (Inbound) and `0` (Outbound)
	Direction *uint64 `json:"Direction,omitnil,omitempty" name:"Direction"`

	// Rule sequence number
	OrderIndex *int64 `json:"OrderIndex,omitnil,omitempty" name:"OrderIndex"`

	// Rule status. `true` (Enabled); `false` (Disabled)
	Enable *string `json:"Enable,omitnil,omitempty" name:"Enable"`

	// The unique ID of the rule, which is not required when you create a rule.
	Uuid *int64 `json:"Uuid,omitnil,omitempty" name:"Uuid"`

	// Description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type CreateRuleItem struct {
	// Rule direction: 1 means inbound, 0 means outbound; other integers or omitted values result in verification failure. Direction also determines the available combinations of SourceType, TargetType, Scope, and Protocol.
	Direction *uint64 `json:"Direction,omitnil,omitempty" name:"Direction"`

	// Rule order, required. Pass -1 to append to the end of the current direction; a positive serial number indicates insertion at the corresponding position and postponement of subsequent rules; treat 0 as 1, other negative numbers and out-of-scope values should not be used. When a new request contains multiple rules, Direction must be the same; pass all -1 for appending, pass consecutive incremental positive serial numbers in request order for insertion. A modify request accepts only one rule.
	OrderIndex *int64 `json:"OrderIndex,omitnil,omitempty" name:"OrderIndex"`

	// Destination port. Ignore this field and set it to an empty string when Protocol is ICMP. For other protocols, you must provide a parse string. You can enter a positive integer single port or a "start/end" range separated with commas. The starting value must not be larger than the end value. -1/-1 indicates all ports. FTP only accepts one positive integer. For domain or domain name template targets within the side or all scope, only -1/-1 or 0/65535 are accepted.
	Port *string `json:"Port,omitnil,omitempty" name:"Port"`

	// Protocol, case-insensitive parsing. Layer-4 values TCP, UDP, ICMP, ICMPV6, ANY are normalized to uppercase. Application-layer values HTTP, HTTPS, HTTP/HTTPS, SMTP, SMTPS, SMTP/SMTPS, FTP, DNS, TLS/SSL and aliases domain, TLS, SSL are normalized to corresponding standard values. ANY means no protocol limitation, not an empty Protocol. It belongs to both parseable Layer-4 protocols and application protocols. domain, TLS, SSL are all normalized to TLS/SSL. The target for domain or domain name template accepts the above application-layer protocols and ANY, but does not accept FTP and other Layer-4 protocols. dnsparse and domainiptwoverify only accept TCP or UDP and only support serial. Other targets in the public cloud environment do not accept application-layer protocols outside of FTP and ANY. Under the side or all scope, inbound only accepts TCP, outbound only accepts TCP, HTTP/HTTPS, or TLS/SSL. When DNS is used for a non-domain target and the target is not *, the destination content must also be a rule list of valid non-IP domain names. When using a protocol port template, each group of protocol and port in the template also executes these integration validations.
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// Traffic processing actions are case-insensitive. accept means allow, drop means deny, and log means observe. isolateinaccept means allow access to allowlisted traffic for isolated assets, isolateindrop means block access to other traffic for isolated assets, isolateoutaccept means allow isolated assets to access allowlisted targets, and isolateoutdrop means block isolated assets from accessing other targets. drop and its deny alias also verify whether the current account has Internet boundary blocking capability.
	RuleAction *string `json:"RuleAction,omitnil,omitempty" name:"RuleAction"`

	// Access the source content. For ip or net, use a valid IP/CIDR list, with a maximum of 10 items in a regular list. For template, use a parseable address template identifier of the current account. When Direction=0, use the corresponding resource identifier for instance, group, and tag. Among them, the instance must be resolvable to a public IP, and the tag must exist with the format {"Key":"tag key","Value":"tag value"}. When Direction=1, use a CSV of region codes for location, which must pass the verification of the new regional rules capability of the current account. For vendor, use a CSV of tencent, aliyun, aws, huawei, azure, or all. location and vendor are converted to region or manufacturer matchmaking information when saved.
	SourceContent *string `json:"SourceContent,omitnil,omitempty" name:"SourceContent"`

	// Access source type, case-insensitive parsing. net and ip both indicate IP/CIDR, template refers to address template, instance refers to asset instance, group refers to asset group, tag refers to resource tag, location refers to region, vendor refers to cloud service provider. Direction=1 accepts ip, net, template, location, vendor; Direction=0 accepts ip, net, template, instance, group, tag. ip and net are handled as the same type.
	SourceType *string `json:"SourceType,omitnil,omitempty" name:"SourceType"`

	// Access the destination content. For ip or net, use a valid IP/CIDR list. For domain, use a valid IP, standard domain name, or wildcard domain name list, and also accept a single *. The regular list supports up to 10 items, and wildcard domain names support up to 5 levels. When domain is used with the DNS protocol, IP is not accepted. For dnsparse, use a single valid domain name, wildcard domain name, or an mb_ domain name template that can be resolved by the current account. For domainiptwoverify, use a single valid domain name without wildcards or such a template. Both do not accept a single *, IP, comma-separated list, or wildcard domain names within the segment. For serial domain segments with wildcards and domainiptwoverify templates, the current environment must support the corresponding capacity. For template, use the address template identifier that can be resolved by the current account. For Direction=1, instance, group, and tag use the corresponding resource identifiers. The instance must be resolvable to a public network IP, and the tag must exist with the format {"Key":"tag key","Value":"tag value"}. For Direction=0, location uses region code CSV, and vendor uses CSV of tencent, aliyun, aws, huawei, azure, or all. The standardized content has a maximum length of 1023.
	TargetContent *string `json:"TargetContent,omitnil,omitempty" name:"TargetContent"`

	// Access purpose type. Case-insensitive parsing. net and ip both indicate IP/CIDR, template indicates address template, instance indicates asset instance, group indicates asset group, tag indicates resource tag, location indicates region, vendor indicates cloud service provider, domain indicates FQDN matching (content can also be IP or *), dnsparse indicates loose matching: Host/SNI matches the domain name, or the destination IP belongs to the IP range of the current DNS resolution result of that domain name, hit if any condition is met; domainiptwoverify indicates strict matching: the above two conditions require simultaneous satisfaction. Direction=1 accepts ip, net, template, domain, instance, group, tag; Direction=0 accepts ip, net, template, domain, dnsparse, domainiptwoverify, location, or vendor.
	TargetType *string `json:"TargetType,omitnil,omitempty" name:"TargetType"`

	// Rule description, no more than 100 characters. When projects are added, save the requested value; when modified, replace fully, do not inherit old values.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Enable status. Non-empty values accept the string true or false in a case-insensitive manner and are normalized to enable or disable. When omitted or an empty string is input, the default enabled configuration for access control of the current account is read. If this configuration is unavailable, it is enabled by default. Existing rules are replaced without inheriting old values.
	Enable *string `json:"Enable,omitnil,omitempty" name:"Enable"`

	// Associated alarm or source event ID. When projects are added, omit or input an empty string to indicate not associated. When modifying, import the rules[].log_id returned by DescribeCfwRules as is. If not returned, omit or input an empty string. The old value is not automatically inherited during replacement. When From=batch_import_cover, a non-empty value is also reused as the string literal identification for the rule after overwrite import.
	LogId *string `json:"LogId,omitnil,omitempty" name:"LogId"`

	// Protocol port template ID. Omit or input an empty string to indicate no template is used. If not empty, it must point to an existing template in the current account with the content format "protocol:port", otherwise the request fails. The protocol and port in the template must meet the integration restrictions of Direction, TargetType, and Scope. Protocol and Port must still comply with their respective field rules, but are not required to be fixed as ANY, -1/-1, or serial.
	ParamTemplateId *string `json:"ParamTemplateId,omitnil,omitempty" name:"ParamTemplateId"`

	// Rule source: 0 means General rule, 2 means isolated asset outgoing access rule. It can be omitted when projects are added, and omitted values are handled as 0. Only 0 or 2 are accepted for explicit input and modification, and the original rule value should be imported during modification.
	RuleSource *int64 `json:"RuleSource,omitnil,omitempty" name:"RuleSource"`

	// Effective scope. Case insensitivity: serial means only Internet boundary serial firewall, side means only Internet boundary bypass firewall, all means acting on both serial and bypass firewalls simultaneously. Omitted, empty string, or other values will result in verification failure. The international site environment will normalize valid user-submitted input to serial. For linkage restrictions on protocol, port, destination type, and protocol port templates, refer to Protocol, Port, and ParamTemplateId.
	Scope *string `json:"Scope,omitnil,omitempty" name:"Scope"`

	// Rule numeric value ID. Ordinary new additions, user-specified location additions, and batch import ignore this field; positive integer ID is usable when From=batch_import_cover; must provide an existing and modifiable positive integer ID of the current account for modification, used for locating and fully replacing the original rule, omitted, non-positive integer, or non-existing IDs cause request failure.
	Uuid *int64 `json:"Uuid,omitnil,omitempty" name:"Uuid"`
}

// Predefined struct for user
type CreateSecurityGroupRulesRequestParams struct {
	// Added enterprise security group rule data
	Data []*SecurityGroupListData `json:"Data,omitnil,omitempty" name:"Data"`

	// Direction. 0: outbound; 1: inbound. 1 by default
	Direction *uint64 `json:"Direction,omitnil,omitempty" name:"Direction"`

	// 0: at the end; 1: at the top; 2: in the middle. 0 by default
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// Indicates whether to enable rules after addition. 0: disable; 1: enable. 1 by default
	Enable *uint64 `json:"Enable,omitnil,omitempty" name:"Enable"`
}

type CreateSecurityGroupRulesRequest struct {
	*tchttp.BaseRequest
	
	// Added enterprise security group rule data
	Data []*SecurityGroupListData `json:"Data,omitnil,omitempty" name:"Data"`

	// Direction. 0: outbound; 1: inbound. 1 by default
	Direction *uint64 `json:"Direction,omitnil,omitempty" name:"Direction"`

	// 0: at the end; 1: at the top; 2: in the middle. 0 by default
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// Indicates whether to enable rules after addition. 0: disable; 1: enable. 1 by default
	Enable *uint64 `json:"Enable,omitnil,omitempty" name:"Enable"`
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
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// Direction. 0: outbound; 1: inbound
	Direction *uint64 `json:"Direction,omitnil,omitempty" name:"Direction"`

	// Edge ID between two VPCs
	EdgeId *string `json:"EdgeId,omitnil,omitempty" name:"EdgeId"`

	// NAT region, e.g. ap-shanghai/ap-guangzhou/ap-chongqing
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`
}

type DeleteAcRuleRequest struct {
	*tchttp.BaseRequest
	
	// The ID of the rule to delete. It can be queried via the DescribeAcLists API.
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// Direction. 0: outbound; 1: inbound
	Direction *uint64 `json:"Direction,omitnil,omitempty" name:"Direction"`

	// Edge ID between two VPCs
	EdgeId *string `json:"EdgeId,omitnil,omitempty" name:"EdgeId"`

	// NAT region, e.g. ap-shanghai/ap-guangzhou/ap-chongqing
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`
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
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Returns redundant information
	// Note: This field may return `null`, indicating that no valid value was found.
	Info *string `json:"Info,omitnil,omitempty" name:"Info"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Direction *uint64 `json:"Direction,omitnil,omitempty" name:"Direction"`

	// Deletes all the access control rules for inter-VPC firewall toggles associated with the EdgeId. It is empty by default. Enter either EdgeId or Area.
	EdgeId *string `json:"EdgeId,omitnil,omitempty" name:"EdgeId"`

	// Deletes all the access control rules for NAT firewalls of this region. It is empty by default. Enter either EdgeId or Area.
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`
}

type DeleteAllAccessControlRuleRequest struct {
	*tchttp.BaseRequest
	
	// Direction. 0: outbound; 1: inbound. 0 by default
	Direction *uint64 `json:"Direction,omitnil,omitempty" name:"Direction"`

	// Deletes all the access control rules for inter-VPC firewall toggles associated with the EdgeId. It is empty by default. Enter either EdgeId or Area.
	EdgeId *string `json:"EdgeId,omitnil,omitempty" name:"EdgeId"`

	// Deletes all the access control rules for NAT firewalls of this region. It is empty by default. Enter either EdgeId or Area.
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`
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
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Number of access control rules deleted.
	// Note: This field may return `null`, indicating that no valid value was found.
	Info *int64 `json:"Info,omitnil,omitempty" name:"Info"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`
}

type DeleteResourceGroupRequest struct {
	*tchttp.BaseRequest
	
	// Group ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`
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
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// Tencent Cloud region (abbreviation)
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// Direction. 0: outbound; 1: inbound
	Direction *uint64 `json:"Direction,omitnil,omitempty" name:"Direction"`

	// Indicates whether to delete the reverse rule. 0: no; 1: yes
	IsDelReverse *uint64 `json:"IsDelReverse,omitnil,omitempty" name:"IsDelReverse"`
}

type DeleteSecurityGroupRuleRequest struct {
	*tchttp.BaseRequest
	
	// ID of the rule to delete
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// Tencent Cloud region (abbreviation)
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// Direction. 0: outbound; 1: inbound
	Direction *uint64 `json:"Direction,omitnil,omitempty" name:"Direction"`

	// Indicates whether to delete the reverse rule. 0: no; 1: yes
	IsDelReverse *uint64 `json:"IsDelReverse,omitnil,omitempty" name:"IsDelReverse"`
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
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Returns redundant information
	// Note: This field may return `null`, indicating that no valid value was found.
	Info *string `json:"Info,omitnil,omitempty" name:"Info"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	SourceContent *string `json:"SourceContent,omitnil,omitempty" name:"SourceContent"`

	// Access destination.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TargetContent *string `json:"TargetContent,omitnil,omitempty" name:"TargetContent"`

	// Protocol.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// Port.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Port *string `json:"Port,omitnil,omitempty" name:"Port"`

	// Action that Cloud Firewall performs on the traffic. Valid values: accept (allow), drop (reject), and log (monitor).
	// Note: This field may return null, indicating that no valid values can be obtained.
	RuleAction *string `json:"RuleAction,omitnil,omitempty" name:"RuleAction"`

	// Description.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Number of rule matching times.
	Count *uint64 `json:"Count,omitnil,omitempty" name:"Count"`

	// Rule sequence number.
	OrderIndex *uint64 `json:"OrderIndex,omitnil,omitempty" name:"OrderIndex"`

	// Access source type. Valid values for an inbound rule: ip, net, template, and location; valid values for an outbound rule: ip, net, template, instance, group, and tag.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SourceType *string `json:"SourceType,omitnil,omitempty" name:"SourceType"`

	// Access destination type. Valid values for an inbound rule: ip, net, template, instance, group, and tag; valid values for an outbound rule: ip, net, domain, template, and location.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TargetType *string `json:"TargetType,omitnil,omitempty" name:"TargetType"`

	// Unique ID of the rule.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Uuid *uint64 `json:"Uuid,omitnil,omitempty" name:"Uuid"`

	// Rule validity.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Invalid *uint64 `json:"Invalid,omitnil,omitempty" name:"Invalid"`

	// Valid values: 0: common rules; 1: regional rules.
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsRegion *uint64 `json:"IsRegion,omitnil,omitempty" name:"IsRegion"`

	// Country ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CountryCode *uint64 `json:"CountryCode,omitnil,omitempty" name:"CountryCode"`

	// City ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CityCode *uint64 `json:"CityCode,omitnil,omitempty" name:"CityCode"`

	// Country name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CountryName *string `json:"CountryName,omitnil,omitempty" name:"CountryName"`

	// City name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CityName *string `json:"CityName,omitnil,omitempty" name:"CityName"`

	// Cloud provider code.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CloudCode *string `json:"CloudCode,omitnil,omitempty" name:"CloudCode"`

	// Valid values: 0: common rules; 1: cloud provider rules.
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsCloud *uint64 `json:"IsCloud,omitnil,omitempty" name:"IsCloud"`

	// Rule status. Valid values: true: enabled; false: disabled.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Enable *string `json:"Enable,omitnil,omitempty" name:"Enable"`

	// Rule direction. Valid values: 1: inbound; 0: outbound.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Direction *uint64 `json:"Direction,omitnil,omitempty" name:"Direction"`

	// Instance name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// UUID for internal use. Generally, this field is not required.
	// Note: This field may return null, indicating that no valid values can be obtained.
	InternalUuid *int64 `json:"InternalUuid,omitnil,omitempty" name:"InternalUuid"`

	// Rule status. This field is valid when you query rule matching details. Valid values: 0: new; 1: deleted; 2: edited and deleted.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Details of associated tasks
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	BetaList []*BetaInfoByACL `json:"BetaList,omitnil,omitempty" name:"BetaList"`
}

// Predefined struct for user
type DescribeAcListsRequestParams struct {
	// Protocol
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// Policy
	Strategy *string `json:"Strategy,omitnil,omitempty" name:"Strategy"`

	// Search value
	SearchValue *string `json:"SearchValue,omitnil,omitempty" name:"SearchValue"`

	// Number of entries per page
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Indicates whether it is outbound or inbound. 1: inbound; 0: outbound
	Direction *uint64 `json:"Direction,omitnil,omitempty" name:"Direction"`

	// EdgeId value
	EdgeId *string `json:"EdgeId,omitnil,omitempty" name:"EdgeId"`

	// Indicates whether the rule is enabled. '0': disabled; '1': enabled. '0' by default
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Region
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeAcListsRequest struct {
	*tchttp.BaseRequest
	
	// Protocol
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// Policy
	Strategy *string `json:"Strategy,omitnil,omitempty" name:"Strategy"`

	// Search value
	SearchValue *string `json:"SearchValue,omitnil,omitempty" name:"SearchValue"`

	// Number of entries per page
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Indicates whether it is outbound or inbound. 1: inbound; 0: outbound
	Direction *uint64 `json:"Direction,omitnil,omitempty" name:"Direction"`

	// EdgeId value
	EdgeId *string `json:"EdgeId,omitnil,omitempty" name:"EdgeId"`

	// Indicates whether the rule is enabled. '0': disabled; '1': enabled. '0' by default
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Region
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
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
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// Access control list data
	Data []*AcListsData `json:"Data,omitnil,omitempty" name:"Data"`

	// Total entries excluding the filtered ones
	AllTotal *uint64 `json:"AllTotal,omitnil,omitempty" name:"AllTotal"`

	// All access control rules enabled/disabled
	// Note: This field may return `null`, indicating that no valid value was found.
	Enable *uint64 `json:"Enable,omitnil,omitempty" name:"Enable"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of records per page
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Region code (e.g. ap-guangzhou). All Tencent Cloud regions are supported.
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// Additional search criteria (JSON string)
	SearchValue *string `json:"SearchValue,omitnil,omitempty" name:"SearchValue"`

	// Sorting field
	By *string `json:"By,omitnil,omitempty" name:"By"`

	// Sort order. asc: ascending; desc: descending
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// Security group ID
	SecurityGroupId *string `json:"SecurityGroupId,omitnil,omitempty" name:"SecurityGroupId"`

	// Instance type. '3': CVM instance; '4': CLB instance; '5': ENI instance; '6': Cloud database
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type DescribeAssociatedInstanceListRequest struct {
	*tchttp.BaseRequest
	
	// List offset
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of records per page
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Region code (e.g. ap-guangzhou). All Tencent Cloud regions are supported.
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// Additional search criteria (JSON string)
	SearchValue *string `json:"SearchValue,omitnil,omitempty" name:"SearchValue"`

	// Sorting field
	By *string `json:"By,omitnil,omitempty" name:"By"`

	// Sort order. asc: ascending; desc: descending
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// Security group ID
	SecurityGroupId *string `json:"SecurityGroupId,omitnil,omitempty" name:"SecurityGroupId"`

	// Instance type. '3': CVM instance; '4': CLB instance; '5': ENI instance; '6': Cloud database
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
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
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// Instance list
	// Note: This field may return `null`, indicating that no valid value was found.
	Data []*AssociatedInstanceInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// End time
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// IP search criteria
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// Start time
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Direction
	Direction *string `json:"Direction,omitnil,omitempty" name:"Direction"`

	// Inter-VPC firewall toggle edge ID
	EdgeId *string `json:"EdgeId,omitnil,omitempty" name:"EdgeId"`

	// Log source. move: inter-VPC firewall
	LogSource *string `json:"LogSource,omitnil,omitempty" name:"LogSource"`

	// Source
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// Region
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`
}

type DescribeBlockByIpTimesListRequest struct {
	*tchttp.BaseRequest
	
	// End time
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// IP search criteria
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// Start time
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Direction
	Direction *string `json:"Direction,omitnil,omitempty" name:"Direction"`

	// Inter-VPC firewall toggle edge ID
	EdgeId *string `json:"EdgeId,omitnil,omitempty" name:"EdgeId"`

	// Log source. move: inter-VPC firewall
	LogSource *string `json:"LogSource,omitnil,omitempty" name:"LogSource"`

	// Source
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// Region
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`
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
	delete(f, "EndTime")
	delete(f, "Ip")
	delete(f, "StartTime")
	delete(f, "Direction")
	delete(f, "EdgeId")
	delete(f, "LogSource")
	delete(f, "Source")
	delete(f, "Zone")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBlockByIpTimesListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBlockByIpTimesListResponseParams struct {
	// Response  data
	Data []*IpStatic `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Page offset.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Direction. Valid values: 1: inbound public access; 0: outbound public access; 3: private network access; empty string: all access.
	Direction *string `json:"Direction,omitnil,omitempty" name:"Direction"`

	// Rule type. Valid values: 1: block; 2: allow.
	RuleType *uint64 `json:"RuleType,omitnil,omitempty" name:"RuleType"`

	// Column by which rules are sorted. Valid values: EndTime: end time; StartTime: start time; MatchTimes: number of matching times.
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// Sort order. Valid values: desc: descending; asc: ascending.
	By *string `json:"By,omitnil,omitempty" name:"By"`

	// Search keys, in a JSON string. Valid values: {}: empty; domain: domain name; level: threat level; ignore_reason: reason for allowing access; rule_source: source of a security event; address: geographical location; common: fuzzy search.
	SearchValue *string `json:"SearchValue,omitnil,omitempty" name:"SearchValue"`
}

type DescribeBlockIgnoreListRequest struct {
	*tchttp.BaseRequest
	
	// Number of entries per page.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Page offset.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Direction. Valid values: 1: inbound public access; 0: outbound public access; 3: private network access; empty string: all access.
	Direction *string `json:"Direction,omitnil,omitempty" name:"Direction"`

	// Rule type. Valid values: 1: block; 2: allow.
	RuleType *uint64 `json:"RuleType,omitnil,omitempty" name:"RuleType"`

	// Column by which rules are sorted. Valid values: EndTime: end time; StartTime: start time; MatchTimes: number of matching times.
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// Sort order. Valid values: desc: descending; asc: ascending.
	By *string `json:"By,omitnil,omitempty" name:"By"`

	// Search keys, in a JSON string. Valid values: {}: empty; domain: domain name; level: threat level; ignore_reason: reason for allowing access; rule_source: source of a security event; address: geographical location; common: fuzzy search.
	SearchValue *string `json:"SearchValue,omitnil,omitempty" name:"SearchValue"`
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
	Data []*BlockIgnoreRule `json:"Data,omitnil,omitempty" name:"Data"`

	// Total number of results, which is used for pagination.
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// Status code. Valid values: 0: successful; others: failed.
	ReturnCode *int64 `json:"ReturnCode,omitnil,omitempty" name:"ReturnCode"`

	// Status message. Valid values: success: successful query; fail: failed query.
	ReturnMsg *string `json:"ReturnMsg,omitnil,omitempty" name:"ReturnMsg"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// List type. Valid values: port, address, or IP
	QueryType *string `json:"QueryType,omitnil,omitempty" name:"QueryType"`

	// Number of top results returned
	Top *int64 `json:"Top,omitnil,omitempty" name:"Top"`

	// Search criteria
	SearchValue *string `json:"SearchValue,omitnil,omitempty" name:"SearchValue"`
}

type DescribeBlockStaticListRequest struct {
	*tchttp.BaseRequest
	
	// Start time
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// List type. Valid values: port, address, or IP
	QueryType *string `json:"QueryType,omitnil,omitempty" name:"QueryType"`

	// Number of top results returned
	Top *int64 `json:"Top,omitnil,omitempty" name:"Top"`

	// Search criteria
	SearchValue *string `json:"SearchValue,omitnil,omitempty" name:"SearchValue"`
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
	Data []*StaticInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeCfwLogsRequestParams struct {
	// Log type. Required for the initial query; cannot be passed when using NextToken for continuation. cfw_netflow_border=Internet boundary traffic, cfw_netflow_vpc=VPC east-west traffic, cfw_netflow_nat=NAT firewall traffic, cfw_netflow_nta=NDR/NTA traffic, cfw_netflow_dns=DNS firewall log, cfw_rule_threatinfo=Intrusion defense/Threat Intelligence Alarm, cfw_rule_acl=Internet Boundary Access Control log, cfw_rule_vpc_acl=VPC access control log, cfw_rule_nat_acl=NAT access control log, cfw_ndr_subject_risk=NDR topic risk, cfw_ndr_dataleak_entry=NDR sensitive data leak, cfw_ndr_ai_audit=NDR AI application identification and Large Model Invocation audit, cfw_feature_collect=Statistical feature and baseline anomaly, cfw_behavior_collect=Beacon/DNS/port/cert/VPC mutual access behavior, operate_log_all=Operation audit log.
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`

	// Log filter expression. Default * means no filtering; for example src_ip:1.1.1.1. Queryable fields vary with LogType. You should preferentially use the field name returned in the corresponding Items. Do not guess non-existing fields. It cannot be passed when using NextToken for continued query.
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// Query start time. Supports RFC3339, YYYY-MM-DD HH:MM:SS, YYYY-MM-DD, or Unix timestamp. Input to query the TimeRange backward from this time. Cannot be imported when using NextToken for continued querying.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Query time range. Default 1h; format is positive integer plus unit m/h/d, such as 5m, 1h, 24h, 7d; cannot be passed when using NextToken for continuation.
	TimeRange *string `json:"TimeRange,omitnil,omitempty" name:"TimeRange"`

	// Return limit. Selectable for initial query, default 100; value 1 to 1000; cannot be passed when using NextToken for continued query.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Previous page opaque continuation token returned by Response.Data. Not required for initial query; only required for continuation query with NextToken. Invalid, tampered, or mismatched tenant will be rejected.
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`
}

type DescribeCfwLogsRequest struct {
	*tchttp.BaseRequest
	
	// Log type. Required for the initial query; cannot be passed when using NextToken for continuation. cfw_netflow_border=Internet boundary traffic, cfw_netflow_vpc=VPC east-west traffic, cfw_netflow_nat=NAT firewall traffic, cfw_netflow_nta=NDR/NTA traffic, cfw_netflow_dns=DNS firewall log, cfw_rule_threatinfo=Intrusion defense/Threat Intelligence Alarm, cfw_rule_acl=Internet Boundary Access Control log, cfw_rule_vpc_acl=VPC access control log, cfw_rule_nat_acl=NAT access control log, cfw_ndr_subject_risk=NDR topic risk, cfw_ndr_dataleak_entry=NDR sensitive data leak, cfw_ndr_ai_audit=NDR AI application identification and Large Model Invocation audit, cfw_feature_collect=Statistical feature and baseline anomaly, cfw_behavior_collect=Beacon/DNS/port/cert/VPC mutual access behavior, operate_log_all=Operation audit log.
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`

	// Log filter expression. Default * means no filtering; for example src_ip:1.1.1.1. Queryable fields vary with LogType. You should preferentially use the field name returned in the corresponding Items. Do not guess non-existing fields. It cannot be passed when using NextToken for continued query.
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// Query start time. Supports RFC3339, YYYY-MM-DD HH:MM:SS, YYYY-MM-DD, or Unix timestamp. Input to query the TimeRange backward from this time. Cannot be imported when using NextToken for continued querying.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Query time range. Default 1h; format is positive integer plus unit m/h/d, such as 5m, 1h, 24h, 7d; cannot be passed when using NextToken for continuation.
	TimeRange *string `json:"TimeRange,omitnil,omitempty" name:"TimeRange"`

	// Return limit. Selectable for initial query, default 100; value 1 to 1000; cannot be passed when using NextToken for continued query.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Previous page opaque continuation token returned by Response.Data. Not required for initial query; only required for continuation query with NextToken. Invalid, tampered, or mismatched tenant will be rejected.
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`
}

func (r *DescribeCfwLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCfwLogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LogType")
	delete(f, "Query")
	delete(f, "StartTime")
	delete(f, "TimeRange")
	delete(f, "Limit")
	delete(f, "NextToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCfwLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCfwLogsResponseParams struct {
	// Query result. UTF-8 JSON object string; the caller must parse Response.Data. Items is the log array of the current page, and fields vary with LogType. TotalCount is the return limit of the current page, Limit is the page size, and LogType and TimeWindow echo the query scope. When HasMore=true, NextToken must be saved and used as-is for continued querying. When HasMore=false, pagination ends.
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCfwLogsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCfwLogsResponseParams `json:"Response"`
}

func (r *DescribeCfwLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCfwLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCfwStatusMonitorRequestParams struct {
	// Operation type. describe_scene means discovery of scenarios and secondary dropdown options; fetch_scene means acquisition of scenario-based snapshots. Required.
	Op *string `json:"Op,omitnil,omitempty" name:"Op"`

	// Firewall scenario type. Supports internet_edge (Internet edge firewall), nat_cluster (NAT border firewall - cluster), nat_ha (NAT border firewall - primary/secondary), vpc_cluster (VPC boundary firewall - cluster), vpc_ha (VPC boundary firewall - primary/secondary). Required.
	FirewallType *string `json:"FirewallType,omitnil,omitempty" name:"FirewallType"`

	// Secondary dropdown option ID. fetch_scene is imported as needed, and the value comes from selection.available_options[].ID returned by describe_scene. internet_edge is the region, NAT is the instance ID, and VPC bandwidth scenario is the firewall group ID. The connections aggregation scenario for VPC_cluster ignores this parameter.
	SelectionId *string `json:"SelectionId,omitnil,omitempty" name:"SelectionId"`

	// Secondary dropdown display name. Can be used as an alternative to SelectionId for matching by name. The value comes from selection.available_options[].name returned by describe_scene.
	SelectionName *string `json:"SelectionName,omitnil,omitempty" name:"SelectionName"`

	// Engine instance ID. Mainly used in vpc ha scenarios where a firewall group corresponds to multiple instances. Preferentially use the selection.available_options[].instance_ID returned by describe_scene. If only instance_ids are available, select a string value from the array.
	SelectionInstanceId *string `json:"SelectionInstanceId,omitnil,omitempty" name:"SelectionInstanceId"`

	// Metrics tab. fetch_scene can be passed; used when not passed, this scenario default value. Support bandwidth, connections.
	Metric *string `json:"Metric,omitnil,omitempty" name:"Metric"`

	// Perspective under the metric. fetch_scene is optional; the default value for this scenario is used when not provided. Supports ip, subnet, session, switch, and vpc. The actual usable composite is subject to the return from describe_scene.
	Perspective *string `json:"Perspective,omitnil,omitempty" name:"Perspective"`

	// NAT primary/secondary number of connections IP perspective range. External means external IP, asset means Asset IP. Only nat_ha + connections + ip is used. Other group input will return InvalidParameter.
	IpScope *string `json:"IpScope,omitnil,omitempty" name:"IpScope"`

	// Preset time range. Default 24h; used by fetch_scene. Supports 5m, 15m, 30m, 1h, 6h, 24h, 3d, 7d, 30d, today, yesterday, day before yesterday, this week, last week, this month.
	TimePreset *string `json:"TimePreset,omitnil,omitempty" name:"TimePreset"`

	// Custom start time. Format YYYY-MM-DD HH:MM:SS; must be specified together with EndTime, maximum span 30 days.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Custom end time. Format YYYY-MM-DD HH:MM:SS; must be consistent with StartTime at the same time, maximum span 30 days.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Page number, starting from 1. Default is 1; used for the fetch_scene list viewing angle.
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`

	// Entries per page. Default 10, value 1 to 100; used for the viewing angle of the fetch_scene list.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Whether to only get overview data. When true, fetch_scene only requests overview, skips table/detail, and is suitable for viewing scenario snapshot summary.
	OverviewOnly *bool `json:"OverviewOnly,omitnil,omitempty" name:"OverviewOnly"`

	// Original offset coverage. Option, overwrites the calculation result of Page after input; value 0 to 10000.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Sorting field. Option. InputMax and OutputMax are supported for the Internet boundary IP and NAT IP/subnet perspective. SwitchName is supported for the VPC switch perspective. FlowMax is supported for the VPC IP/VPC perspective. Do not pass other groups.
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// Sorting order. Default desc; supports asc, desc.
	SortOrder *string `json:"SortOrder,omitnil,omitempty" name:"SortOrder"`

	// Filter condition list. Reserved.
	Filters []*CfwStatusMonitorFilter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeCfwStatusMonitorRequest struct {
	*tchttp.BaseRequest
	
	// Operation type. describe_scene means discovery of scenarios and secondary dropdown options; fetch_scene means acquisition of scenario-based snapshots. Required.
	Op *string `json:"Op,omitnil,omitempty" name:"Op"`

	// Firewall scenario type. Supports internet_edge (Internet edge firewall), nat_cluster (NAT border firewall - cluster), nat_ha (NAT border firewall - primary/secondary), vpc_cluster (VPC boundary firewall - cluster), vpc_ha (VPC boundary firewall - primary/secondary). Required.
	FirewallType *string `json:"FirewallType,omitnil,omitempty" name:"FirewallType"`

	// Secondary dropdown option ID. fetch_scene is imported as needed, and the value comes from selection.available_options[].ID returned by describe_scene. internet_edge is the region, NAT is the instance ID, and VPC bandwidth scenario is the firewall group ID. The connections aggregation scenario for VPC_cluster ignores this parameter.
	SelectionId *string `json:"SelectionId,omitnil,omitempty" name:"SelectionId"`

	// Secondary dropdown display name. Can be used as an alternative to SelectionId for matching by name. The value comes from selection.available_options[].name returned by describe_scene.
	SelectionName *string `json:"SelectionName,omitnil,omitempty" name:"SelectionName"`

	// Engine instance ID. Mainly used in vpc ha scenarios where a firewall group corresponds to multiple instances. Preferentially use the selection.available_options[].instance_ID returned by describe_scene. If only instance_ids are available, select a string value from the array.
	SelectionInstanceId *string `json:"SelectionInstanceId,omitnil,omitempty" name:"SelectionInstanceId"`

	// Metrics tab. fetch_scene can be passed; used when not passed, this scenario default value. Support bandwidth, connections.
	Metric *string `json:"Metric,omitnil,omitempty" name:"Metric"`

	// Perspective under the metric. fetch_scene is optional; the default value for this scenario is used when not provided. Supports ip, subnet, session, switch, and vpc. The actual usable composite is subject to the return from describe_scene.
	Perspective *string `json:"Perspective,omitnil,omitempty" name:"Perspective"`

	// NAT primary/secondary number of connections IP perspective range. External means external IP, asset means Asset IP. Only nat_ha + connections + ip is used. Other group input will return InvalidParameter.
	IpScope *string `json:"IpScope,omitnil,omitempty" name:"IpScope"`

	// Preset time range. Default 24h; used by fetch_scene. Supports 5m, 15m, 30m, 1h, 6h, 24h, 3d, 7d, 30d, today, yesterday, day before yesterday, this week, last week, this month.
	TimePreset *string `json:"TimePreset,omitnil,omitempty" name:"TimePreset"`

	// Custom start time. Format YYYY-MM-DD HH:MM:SS; must be specified together with EndTime, maximum span 30 days.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Custom end time. Format YYYY-MM-DD HH:MM:SS; must be consistent with StartTime at the same time, maximum span 30 days.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Page number, starting from 1. Default is 1; used for the fetch_scene list viewing angle.
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`

	// Entries per page. Default 10, value 1 to 100; used for the viewing angle of the fetch_scene list.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Whether to only get overview data. When true, fetch_scene only requests overview, skips table/detail, and is suitable for viewing scenario snapshot summary.
	OverviewOnly *bool `json:"OverviewOnly,omitnil,omitempty" name:"OverviewOnly"`

	// Original offset coverage. Option, overwrites the calculation result of Page after input; value 0 to 10000.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Sorting field. Option. InputMax and OutputMax are supported for the Internet boundary IP and NAT IP/subnet perspective. SwitchName is supported for the VPC switch perspective. FlowMax is supported for the VPC IP/VPC perspective. Do not pass other groups.
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// Sorting order. Default desc; supports asc, desc.
	SortOrder *string `json:"SortOrder,omitnil,omitempty" name:"SortOrder"`

	// Filter condition list. Reserved.
	Filters []*CfwStatusMonitorFilter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeCfwStatusMonitorRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCfwStatusMonitorRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Op")
	delete(f, "FirewallType")
	delete(f, "SelectionId")
	delete(f, "SelectionName")
	delete(f, "SelectionInstanceId")
	delete(f, "Metric")
	delete(f, "Perspective")
	delete(f, "IpScope")
	delete(f, "TimePreset")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Page")
	delete(f, "Limit")
	delete(f, "OverviewOnly")
	delete(f, "Offset")
	delete(f, "SortBy")
	delete(f, "SortOrder")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCfwStatusMonitorRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCfwStatusMonitorResponseParams struct {
	// Query result. UTF-8 JSON object string; the caller needs to parse Response.Data. The scene returned by describe_scene includes metric_options, perspective_options, default_metric, default_perspective, selection_required_by_metric, selection_kind_by_metric, and time_preset_options; selection.available_options returns options applicable to SelectionId, SelectionName, and SelectionInstanceId. fetch_scene returns a data snapshot of the selected scenario, which may contain overview, table, or detail. The example below is a section of the field structure, and the array only shows representative values.
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCfwStatusMonitorResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCfwStatusMonitorResponseParams `json:"Response"`
}

func (r *DescribeCfwStatusMonitorResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCfwStatusMonitorResponse) FromJsonString(s string) error {
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
	BasicRuleSwitch *int64 `json:"BasicRuleSwitch,omitnil,omitempty" name:"BasicRuleSwitch"`

	// Whether to enable the Security Baseline feature
	BaselineAllSwitch *int64 `json:"BaselineAllSwitch,omitnil,omitempty" name:"BaselineAllSwitch"`

	// Whether to enable the Treat Intelligence feature
	TiSwitch *int64 `json:"TiSwitch,omitnil,omitempty" name:"TiSwitch"`

	// Whether to enable the Virtual Patch feature
	VirtualPatchSwitch *int64 `json:"VirtualPatchSwitch,omitnil,omitempty" name:"VirtualPatchSwitch"`

	// Whether it has been enabled before
	HistoryOpen *int64 `json:"HistoryOpen,omitnil,omitempty" name:"HistoryOpen"`

	// Status code. `0`: Succeeded. Others: Failed
	ReturnCode *int64 `json:"ReturnCode,omitnil,omitempty" name:"ReturnCode"`

	// Status message. `success` and `fail.
	ReturnMsg *string `json:"ReturnMsg,omitnil,omitempty" name:"ReturnMsg"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	PageNo *string `json:"PageNo,omitnil,omitempty" name:"PageNo"`

	// Maximum number of entries per page displayed for query by page number.
	// 
	// Maximum value: 50.
	PageSize *string `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Source example:
	// net: IP/CIDR (192.168.0.2)
	// template: parameter template (ipm-dyodhpby)
	// instance: asset instance (ins-123456)
	// resourcegroup: asset group (/all groups/group 1/subgroup 1)
	// tag: resource tag ({"Key":"tag key","Value":"tag value"})
	// region: region (ap-gaungzhou)
	// Wildcards are supported.
	SourceContent *string `json:"SourceContent,omitnil,omitempty" name:"SourceContent"`

	// Destination example:
	// net: IP/CIDR (192.168.0.2)
	// template: parameter template (ipm-dyodhpby)
	// instance: asset instance (ins-123456)
	// resourcegroup: asset group (/all groups/group 1/subgroup 1)
	// tag: resource tag ({"Key":"tag key","Value":"tag value"})
	// region: region (ap-gaungzhou)
	// Wildcards are supported.
	DestContent *string `json:"DestContent,omitnil,omitempty" name:"DestContent"`

	// Rule description. Wildcards are supported.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// The action that Cloud Firewall performs on the traffic. Valid values:
	// accept: allow
	// drop: deny
	RuleAction *string `json:"RuleAction,omitnil,omitempty" name:"RuleAction"`

	// Indicates whether to enable the rules. Default: enable. Valid values:
	// true: enable; false: disable
	Enable *string `json:"Enable,omitnil,omitempty" name:"Enable"`

	// The port to apply access control rules. Valid values:
	// -1/-1: all ports
	// 80: port 80
	Port *string `json:"Port,omitnil,omitempty" name:"Port"`

	// Protocol. TCP/UDP/ICMP/ANY
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// Parameter template ID of port and protocol type; mutually exclusive with Protocol and Port
	ServiceTemplateId *string `json:"ServiceTemplateId,omitnil,omitempty" name:"ServiceTemplateId"`
}

type DescribeEnterpriseSecurityGroupRuleRequest struct {
	*tchttp.BaseRequest
	
	// Page number of the current page displayed for query by page number.
	// 
	// 1 by default.
	PageNo *string `json:"PageNo,omitnil,omitempty" name:"PageNo"`

	// Maximum number of entries per page displayed for query by page number.
	// 
	// Maximum value: 50.
	PageSize *string `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Source example:
	// net: IP/CIDR (192.168.0.2)
	// template: parameter template (ipm-dyodhpby)
	// instance: asset instance (ins-123456)
	// resourcegroup: asset group (/all groups/group 1/subgroup 1)
	// tag: resource tag ({"Key":"tag key","Value":"tag value"})
	// region: region (ap-gaungzhou)
	// Wildcards are supported.
	SourceContent *string `json:"SourceContent,omitnil,omitempty" name:"SourceContent"`

	// Destination example:
	// net: IP/CIDR (192.168.0.2)
	// template: parameter template (ipm-dyodhpby)
	// instance: asset instance (ins-123456)
	// resourcegroup: asset group (/all groups/group 1/subgroup 1)
	// tag: resource tag ({"Key":"tag key","Value":"tag value"})
	// region: region (ap-gaungzhou)
	// Wildcards are supported.
	DestContent *string `json:"DestContent,omitnil,omitempty" name:"DestContent"`

	// Rule description. Wildcards are supported.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// The action that Cloud Firewall performs on the traffic. Valid values:
	// accept: allow
	// drop: deny
	RuleAction *string `json:"RuleAction,omitnil,omitempty" name:"RuleAction"`

	// Indicates whether to enable the rules. Default: enable. Valid values:
	// true: enable; false: disable
	Enable *string `json:"Enable,omitnil,omitempty" name:"Enable"`

	// The port to apply access control rules. Valid values:
	// -1/-1: all ports
	// 80: port 80
	Port *string `json:"Port,omitnil,omitempty" name:"Port"`

	// Protocol. TCP/UDP/ICMP/ANY
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// Parameter template ID of port and protocol type; mutually exclusive with Protocol and Port
	ServiceTemplateId *string `json:"ServiceTemplateId,omitnil,omitempty" name:"ServiceTemplateId"`
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
	PageNo *string `json:"PageNo,omitnil,omitempty" name:"PageNo"`

	// Maximum number of entries per page displayed for query by page number.
	PageSize *string `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Access control rule list
	Rules []*SecurityGroupRule `json:"Rules,omitnil,omitempty" name:"Rules"`

	// Total number of access control rules
	TotalCount *string `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Data *ScanInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	IPList []*string `json:"IPList,omitnil,omitempty" name:"IPList"`
}

type DescribeIPStatusListRequest struct {
	*tchttp.BaseRequest
	
	// Asset ID
	IPList []*string `json:"IPList,omitnil,omitempty" name:"IPList"`
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
	StatusList []*IPDefendStatus `json:"StatusList,omitnil,omitempty" name:"StatusList"`

	// Status code
	ReturnCode *int64 `json:"ReturnCode,omitnil,omitempty" name:"ReturnCode"`

	// Status information
	ReturnMsg *string `json:"ReturnMsg,omitnil,omitempty" name:"ReturnMsg"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Page offset.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Index to be queried. This parameter is optional, and is used only in specific cases.
	Index *string `json:"Index,omitnil,omitempty" name:"Index"`

	// Filter condition combination.
	Filters []*CommonFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Start time for search. This parameter is optional.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time for search. This parameter is optional.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Valid values: desc: descending; asc: ascending. The returned results are sorted by the value of By. If this parameter is specified, By is also required.
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// Field by which the returned results are sorted.
	By *string `json:"By,omitnil,omitempty" name:"By"`
}

type DescribeNatAcRuleRequest struct {
	*tchttp.BaseRequest
	
	// Number of entries per page.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Page offset.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Index to be queried. This parameter is optional, and is used only in specific cases.
	Index *string `json:"Index,omitnil,omitempty" name:"Index"`

	// Filter condition combination.
	Filters []*CommonFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Start time for search. This parameter is optional.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time for search. This parameter is optional.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Valid values: desc: descending; asc: ascending. The returned results are sorted by the value of By. If this parameter is specified, By is also required.
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// Field by which the returned results are sorted.
	By *string `json:"By,omitnil,omitempty" name:"By"`
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
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// NAT access control list data.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data []*DescAcItem `json:"Data,omitnil,omitempty" name:"Data"`

	// Total number of entries returned without filtering.
	AllTotal *uint64 `json:"AllTotal,omitnil,omitempty" name:"AllTotal"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ReturnMsg *string `json:"ReturnMsg,omitnil,omitempty" name:"ReturnMsg"`

	// Number of NAT instances of the current tenant
	// Note: This field may return `null`, indicating that no valid value was found.
	NatFwInsCount *int64 `json:"NatFwInsCount,omitnil,omitempty" name:"NatFwInsCount"`

	// Number of subnets connected by the current tenant
	// Note: This field may return `null`, indicating that no valid value was found.
	SubnetCount *int64 `json:"SubnetCount,omitnil,omitempty" name:"SubnetCount"`

	// Number of firewalls enabled
	// Note: This field may return `null`, indicating that no valid value was found.
	OpenSwitchCount *int64 `json:"OpenSwitchCount,omitnil,omitempty" name:"OpenSwitchCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	NatinsLst []*NatFwInstance `json:"NatinsLst,omitnil,omitempty" name:"NatinsLst"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	NatinsLst []*NatFwInstance `json:"NatinsLst,omitnil,omitempty" name:"NatinsLst"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Filter []*NatFwFilter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// Page number
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Page length
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeNatFwInstancesInfoRequest struct {
	*tchttp.BaseRequest
	
	// Gets filtering fields of instance list
	Filter []*NatFwFilter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// Page number
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Page length
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
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
	NatinsLst []*NatInstanceInfo `json:"NatinsLst,omitnil,omitempty" name:"NatinsLst"`

	// Number of NAT firewalls
	// Note: This field may return `null`, indicating that no valid value was found.
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	NatFwInsId *string `json:"NatFwInsId,omitnil,omitempty" name:"NatFwInsId"`

	// Content filtered by NAT firewall, separated with ","
	NatInsIdFilter *string `json:"NatInsIdFilter,omitnil,omitempty" name:"NatInsIdFilter"`

	// Number of pages
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Maximum entries per page
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeNatFwVpcDnsLstRequest struct {
	*tchttp.BaseRequest
	
	// NAT firewall instance ID
	NatFwInsId *string `json:"NatFwInsId,omitnil,omitempty" name:"NatFwInsId"`

	// Content filtered by NAT firewall, separated with ","
	NatInsIdFilter *string `json:"NatInsIdFilter,omitnil,omitempty" name:"NatInsIdFilter"`

	// Number of pages
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Maximum entries per page
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
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
	VpcDnsSwitchLst []*VpcDnsInfo `json:"VpcDnsSwitchLst,omitnil,omitempty" name:"VpcDnsSwitchLst"`

	// Response parameter
	// Note: This field may return `null`, indicating that no valid value was found.
	ReturnMsg *string `json:"ReturnMsg,omitnil,omitempty" name:"ReturnMsg"`

	// Total number of toggles
	// Note: This field may return `null`, indicating that no valid value was found.
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	QueryType *string `json:"QueryType,omitnil,omitempty" name:"QueryType"`

	// Asset group ID, 0: all asset group IDs
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// all: all, including subgroups; own: my asset groups only
	ShowType *string `json:"ShowType,omitnil,omitempty" name:"ShowType"`
}

type DescribeResourceGroupNewRequest struct {
	*tchttp.BaseRequest
	
	// Query type. Network–VPC; business recognition–resource; resource tag–tag
	QueryType *string `json:"QueryType,omitnil,omitempty" name:"QueryType"`

	// Asset group ID, 0: all asset group IDs
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// all: all, including subgroups; own: my asset groups only
	ShowType *string `json:"ShowType,omitnil,omitempty" name:"ShowType"`
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
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// Number of uncategorizd instances
	UnResourceNum *int64 `json:"UnResourceNum,omitnil,omitempty" name:"UnResourceNum"`

	// Response message
	ReturnMsg *string `json:"ReturnMsg,omitnil,omitempty" name:"ReturnMsg"`

	// Return code. 0: Request successful
	ReturnCode *int64 `json:"ReturnCode,omitnil,omitempty" name:"ReturnCode"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Direction *uint64 `json:"Direction,omitnil,omitempty" name:"Direction"`
}

type DescribeRuleOverviewRequest struct {
	*tchttp.BaseRequest
	
	// Direction. 0: outbound; 1: inbound
	Direction *uint64 `json:"Direction,omitnil,omitempty" name:"Direction"`
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
	AllTotal *uint64 `json:"AllTotal,omitnil,omitempty" name:"AllTotal"`

	// Number of blocking rules
	// Note: This field may return `null`, indicating that no valid value was found.
	StrategyNum *uint64 `json:"StrategyNum,omitnil,omitempty" name:"StrategyNum"`

	// Number of enabled rules
	// Note: This field may return `null`, indicating that no valid value was found.
	StartRuleNum *uint64 `json:"StartRuleNum,omitnil,omitempty" name:"StartRuleNum"`

	// Number of disabled rules
	// Note: This field may return `null`, indicating that no valid value was found.
	StopRuleNum *uint64 `json:"StopRuleNum,omitnil,omitempty" name:"StopRuleNum"`

	// Remaining quota
	// Note: This field may return `null`, indicating that no valid value was found.
	RemainingNum *int64 `json:"RemainingNum,omitnil,omitempty" name:"RemainingNum"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Direction *uint64 `json:"Direction,omitnil,omitempty" name:"Direction"`

	// Region code (e.g. ap-guangzhou ). All Tencent Cloud regions are supported.
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// Search value
	SearchValue *string `json:"SearchValue,omitnil,omitempty" name:"SearchValue"`

	// Number of entries per page. Default: 10
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset. Default: 0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Status. Null: all; '0': filter disabled rules; '1': filter enabled rules
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 0: not filter; 1: filter out normal rules to retain abnormal rules
	Filter *uint64 `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeSecurityGroupListRequest struct {
	*tchttp.BaseRequest
	
	// 0: outbound rule; 1: inbound rule
	Direction *uint64 `json:"Direction,omitnil,omitempty" name:"Direction"`

	// Region code (e.g. ap-guangzhou ). All Tencent Cloud regions are supported.
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// Search value
	SearchValue *string `json:"SearchValue,omitnil,omitempty" name:"SearchValue"`

	// Number of entries per page. Default: 10
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset. Default: 0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Status. Null: all; '0': filter disabled rules; '1': filter enabled rules
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 0: not filter; 1: filter out normal rules to retain abnormal rules
	Filter *uint64 `json:"Filter,omitnil,omitempty" name:"Filter"`
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
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// Security group rule list data
	Data []*SecurityGroupListData `json:"Data,omitnil,omitempty" name:"Data"`

	// Total entries excluding the filtered ones
	AllTotal *uint64 `json:"AllTotal,omitnil,omitempty" name:"AllTotal"`

	// All access control rules enabled/disabled
	// Note: This field may return `null`, indicating that no valid value was found.
	Enable *uint64 `json:"Enable,omitnil,omitempty" name:"Enable"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	FuzzySearch *string `json:"FuzzySearch,omitnil,omitempty" name:"FuzzySearch"`

	// Asset type. 1: public network; 2: private network
	InsType *string `json:"InsType,omitnil,omitempty" name:"InsType"`

	// If ChooseType is 1, grouped assets are queried; if ChooseType is not 1, non-grouped assets are queried
	ChooseType *string `json:"ChooseType,omitnil,omitempty" name:"ChooseType"`

	// Region
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Maximum number of results returned per page. For example, if it is set to 10, 10 results will be returned at most.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset of search results
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeSourceAssetRequest struct {
	*tchttp.BaseRequest
	
	// Fuzzy search
	FuzzySearch *string `json:"FuzzySearch,omitnil,omitempty" name:"FuzzySearch"`

	// Asset type. 1: public network; 2: private network
	InsType *string `json:"InsType,omitnil,omitempty" name:"InsType"`

	// If ChooseType is 1, grouped assets are queried; if ChooseType is not 1, non-grouped assets are queried
	ChooseType *string `json:"ChooseType,omitnil,omitempty" name:"ChooseType"`

	// Region
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Maximum number of results returned per page. For example, if it is set to 10, 10 results will be returned at most.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset of search results
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
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
	ZoneList []*AssetZone `json:"ZoneList,omitnil,omitempty" name:"ZoneList"`

	// Data
	Data []*InstanceInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// Total number of returned data
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Asset type, e.g. CVM/NAT/VPN/CLB/others
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Region, e.g. Shanghai, Chongqing, Guangzhou, etc.
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// Search value, e.g. "{"common":"106.54.189.45"}"
	SearchValue *string `json:"SearchValue,omitnil,omitempty" name:"SearchValue"`

	// Number of entries. Default: 10
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset. Default: 0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Sort order. desc: descending; asc: ascending
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// Sorting field. PortTimes (number of risky ports)
	By *string `json:"By,omitnil,omitempty" name:"By"`
}

type DescribeSwitchListsRequest struct {
	*tchttp.BaseRequest
	
	// Firewall status. 0: disabled; 1: enabled
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Asset type, e.g. CVM/NAT/VPN/CLB/others
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Region, e.g. Shanghai, Chongqing, Guangzhou, etc.
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// Search value, e.g. "{"common":"106.54.189.45"}"
	SearchValue *string `json:"SearchValue,omitnil,omitempty" name:"SearchValue"`

	// Number of entries. Default: 10
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset. Default: 0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Sort order. desc: descending; asc: ascending
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// Sorting field. PortTimes (number of risky ports)
	By *string `json:"By,omitnil,omitempty" name:"By"`
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
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// List data
	Data []*SwitchListsData `json:"Data,omitnil,omitempty" name:"Data"`

	// Region list
	AreaLists []*string `json:"AreaLists,omitnil,omitempty" name:"AreaLists"`

	// Number of enabled firewalls
	// Note: This field may return `null`, indicating that no valid value was found.
	OnNum *uint64 `json:"OnNum,omitnil,omitempty" name:"OnNum"`

	// Number of disabled firewalls
	// Note: This field may return `null`, indicating that no valid value was found.
	OffNum *uint64 `json:"OffNum,omitnil,omitempty" name:"OffNum"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Type. 1: alert; 2: block
	QueryType *string `json:"QueryType,omitnil,omitempty" name:"QueryType"`

	// Search criteria
	SearchValue *string `json:"SearchValue,omitnil,omitempty" name:"SearchValue"`
}

type DescribeTLogInfoRequest struct {
	*tchttp.BaseRequest
	
	// Start time
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Type. 1: alert; 2: block
	QueryType *string `json:"QueryType,omitnil,omitempty" name:"QueryType"`

	// Search criteria
	SearchValue *string `json:"SearchValue,omitnil,omitempty" name:"SearchValue"`
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
	Data *TLogInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Type. 1: alert; 2: block
	QueryType *string `json:"QueryType,omitnil,omitempty" name:"QueryType"`

	// Number of top results returned
	Top *int64 `json:"Top,omitnil,omitempty" name:"Top"`

	// Search criteria
	SearchValue *string `json:"SearchValue,omitnil,omitempty" name:"SearchValue"`
}

type DescribeTLogIpListRequest struct {
	*tchttp.BaseRequest
	
	// Start time
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Type. 1: alert; 2: block
	QueryType *string `json:"QueryType,omitnil,omitempty" name:"QueryType"`

	// Number of top results returned
	Top *int64 `json:"Top,omitnil,omitempty" name:"Top"`

	// Search criteria
	SearchValue *string `json:"SearchValue,omitnil,omitempty" name:"SearchValue"`
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
	Data []*StaticInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	EdgeId *string `json:"EdgeId,omitnil,omitempty" name:"EdgeId"`

	// Status value. 0: the only default value
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// NAT region, required for NAT
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// Direction. 0: outbound; 1: inbound. 0 by default
	Direction *uint64 `json:"Direction,omitnil,omitempty" name:"Direction"`
}

type DescribeTableStatusRequest struct {
	*tchttp.BaseRequest
	
	// Edge ID between two VPCs, required for VPCs
	EdgeId *string `json:"EdgeId,omitnil,omitempty" name:"EdgeId"`

	// Status value. 0: the only default value
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// NAT region, required for NAT
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// Direction. 0: outbound; 1: inbound. 0 by default
	Direction *uint64 `json:"Direction,omitnil,omitempty" name:"Direction"`
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
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Gets example ID
	AssetID *string `json:"AssetID,omitnil,omitempty" name:"AssetID"`
}

type DescribeUnHandleEventTabListRequest struct {
	*tchttp.BaseRequest
	
	// Start time
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Gets example ID
	AssetID *string `json:"AssetID,omitnil,omitempty" name:"AssetID"`
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
	Data *UnHandleEvent `json:"Data,omitnil,omitempty" name:"Data"`

	// Error code. 0: success; non-0: error
	ReturnCode *int64 `json:"ReturnCode,omitnil,omitempty" name:"ReturnCode"`

	// Return message: success
	ReturnMsg *string `json:"ReturnMsg,omitnil,omitempty" name:"ReturnMsg"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

// Predefined struct for user
type DescribeVpcAcRuleRequestParams struct {
	// Limit
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Index to be queried, used in specific scenarios, can be left blank
	Index *string `json:"Index,omitnil,omitempty" name:"Index"`

	// Filter combinations
	Filters []*CommonFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Search start time
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Search end time
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Order Type:desc,asc
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// Order By FileName
	By *string `json:"By,omitnil,omitempty" name:"By"`
}

type DescribeVpcAcRuleRequest struct {
	*tchttp.BaseRequest
	
	// Limit
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Index to be queried, used in specific scenarios, can be left blank
	Index *string `json:"Index,omitnil,omitempty" name:"Index"`

	// Filter combinations
	Filters []*CommonFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Search start time
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Search end time
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Order Type:desc,asc
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// Order By FileName
	By *string `json:"By,omitnil,omitempty" name:"By"`
}

func (r *DescribeVpcAcRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpcAcRuleRequest) FromJsonString(s string) error {
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
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVpcAcRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpcAcRuleResponseParams struct {
	// Total Data
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// Data list
	Data []*VpcRuleItem `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeVpcAcRuleResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVpcAcRuleResponseParams `json:"Response"`
}

func (r *DescribeVpcAcRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpcAcRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DnsVpcSwitch struct {
	// VPC ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 0: off; 1: on
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

// Predefined struct for user
type ExpandCfwVerticalRequestParams struct {
	// nat: NAT firewall, ew: east-west firewall
	FwType *string `json:"FwType,omitnil,omitempty" name:"FwType"`

	// Bandwidth value
	Width *uint64 `json:"Width,omitnil,omitempty" name:"Width"`

	// Firewall instance ID
	CfwInstance *string `json:"CfwInstance,omitnil,omitempty" name:"CfwInstance"`
}

type ExpandCfwVerticalRequest struct {
	*tchttp.BaseRequest
	
	// nat: NAT firewall, ew: east-west firewall
	FwType *string `json:"FwType,omitnil,omitempty" name:"FwType"`

	// Bandwidth value
	Width *uint64 `json:"Width,omitnil,omitempty" name:"Width"`

	// Firewall instance ID
	CfwInstance *string `json:"CfwInstance,omitnil,omitempty" name:"CfwInstance"`
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
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	FwCidrType *string `json:"FwCidrType,omitnil,omitempty" name:"FwCidrType"`

	// The IP segment assigned for each VPC.
	FwCidrLst []*FwVpcCidr `json:"FwCidrLst,omitnil,omitempty" name:"FwCidrLst"`

	// The IP segment used by other firewalls. Specify this if you want to assign a dedicated segment for the firewall.
	ComFwCidr *string `json:"ComFwCidr,omitnil,omitempty" name:"ComFwCidr"`
}

type FwVpcCidr struct {
	// VPC ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// IP range of the firewall. The mask must be at least /24.
	FwCidr *string `json:"FwCidr,omitnil,omitempty" name:"FwCidr"`
}

type IPDefendStatus struct {
	// IP address
	IP *string `json:"IP,omitnil,omitempty" name:"IP"`

	// Protection status. 1: enabled; -1: incorrect address; others: disabled
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type InstanceInfo struct {
	// App ID
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// Region
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// VPC ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// VPC name
	VPCName *string `json:"VPCName,omitnil,omitempty" name:"VPCName"`

	// Subnet ID
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Asset ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Asset name
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Asset type
	//  3: CVM instance; 4: CLB instance; 5: ENI instance; 6: MySQL; 7: Redis; 8: NAT; 9: VPN; 10: ES; 11: MariaDB; 12: Kafka; 13: NATFW
	InsType *int64 `json:"InsType,omitnil,omitempty" name:"InsType"`

	// Public IP
	PublicIp *string `json:"PublicIp,omitnil,omitempty" name:"PublicIp"`

	// Private IP
	PrivateIp *string `json:"PrivateIp,omitnil,omitempty" name:"PrivateIp"`

	// Number of ports
	PortNum *string `json:"PortNum,omitnil,omitempty" name:"PortNum"`

	// Number of vulnerabilities
	LeakNum *string `json:"LeakNum,omitnil,omitempty" name:"LeakNum"`

	// 1: public network; 2: private network
	InsSource *string `json:"InsSource,omitnil,omitempty" name:"InsSource"`

	// [a,b]
	// Note: This field may return `null`, indicating that no valid value was found.
	ResourcePath []*string `json:"ResourcePath,omitnil,omitempty" name:"ResourcePath"`
}

type IocListData struct {
	// IP address to be handled. Either IP or Domain is required.
	IP *string `json:"IP,omitnil,omitempty" name:"IP"`

	// 0 or 1. 0: outbound; 1: inbound
	Direction *int64 `json:"Direction,omitnil,omitempty" name:"Direction"`

	// Domain name to be handled. Either IP or Domain is required.
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`
}

type IpStatic struct {
	// Value
	Num *int64 `json:"Num,omitnil,omitempty" name:"Num"`

	// Time shown on the x-axis of the line graph
	StatTime *string `json:"StatTime,omitnil,omitempty" name:"StatTime"`
}

type LogInfo struct {
	// Log time, in milliseconds
	Time *int64 `json:"Time,omitnil,omitempty" name:"Time"`

	// Log topic ID
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// Log topic name
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// Log source IP address
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// Log file name
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// ID of Log Upload Request Packet
	PkgId *string `json:"PkgId,omitnil,omitempty" name:"PkgId"`

	// Log ID in Request Packet
	PkgLogId *string `json:"PkgLogId,omitnil,omitempty" name:"PkgLogId"`

	// JSON serialized string of the log content
	// Note: This field may return null, indicating that no valid values can be obtained.
	LogJson *string `json:"LogJson,omitnil,omitempty" name:"LogJson"`

	// Log source host name
	// Note: This field may return null, indicating that no valid values can be obtained.
	HostName *string `json:"HostName,omitnil,omitempty" name:"HostName"`

	// Raw log (only available when there is an error in creating the log index).
	// Note: This field may return null, indicating that no valid values can be obtained.
	RawLog *string `json:"RawLog,omitnil,omitempty" name:"RawLog"`

	// Cause for log index creation exception. It has a value only when a log index creation exception occurs.
	// Note: This field may return null, indicating that no valid values can be obtained.
	IndexStatus *string `json:"IndexStatus,omitnil,omitempty" name:"IndexStatus"`
}

type LogItem struct {
	// Log key
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// Log Value
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type LogItems struct {
	// Key-Value Data Pair returned from analysis results
	Data []*LogItem `json:"Data,omitnil,omitempty" name:"Data"`
}

// Predefined struct for user
type ModifyAcRuleRequestParams struct {
	// Array of rules
	Data []*RuleInfoData `json:"Data,omitnil,omitempty" name:"Data"`

	// EdgeId value
	EdgeId *string `json:"EdgeId,omitnil,omitempty" name:"EdgeId"`

	// Access rule status
	Enable *int64 `json:"Enable,omitnil,omitempty" name:"Enable"`

	// NAT region
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`
}

type ModifyAcRuleRequest struct {
	*tchttp.BaseRequest
	
	// Array of rules
	Data []*RuleInfoData `json:"Data,omitnil,omitempty" name:"Data"`

	// EdgeId value
	EdgeId *string `json:"EdgeId,omitnil,omitempty" name:"EdgeId"`

	// Access rule status
	Enable *int64 `json:"Enable,omitnil,omitempty" name:"Enable"`

	// NAT region
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`
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
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Returns redundant information
	// Note: This field may return `null`, indicating that no valid value was found.
	Info *string `json:"Info,omitnil,omitempty" name:"Info"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// ID of the selected firewall toggle
	FireWallPublicIPs []*string `json:"FireWallPublicIPs,omitnil,omitempty" name:"FireWallPublicIPs"`
}

type ModifyAllPublicIPSwitchStatusRequest struct {
	*tchttp.BaseRequest
	
	// Status. 0: off; 1: on
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// ID of the selected firewall toggle
	FireWallPublicIPs []*string `json:"FireWallPublicIPs,omitnil,omitempty" name:"FireWallPublicIPs"`
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
	ReturnMsg *string `json:"ReturnMsg,omitnil,omitempty" name:"ReturnMsg"`

	// Error code. 0: success; non-0: failed
	// Note: This field may return `null`, indicating that no valid value was found.
	ReturnCode *int64 `json:"ReturnCode,omitnil,omitempty" name:"ReturnCode"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Direction. 0: outbound; 1: inbound
	Direction *uint64 `json:"Direction,omitnil,omitempty" name:"Direction"`

	// Edge ID value
	EdgeId *string `json:"EdgeId,omitnil,omitempty" name:"EdgeId"`

	// NAT region
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`
}

type ModifyAllRuleStatusRequest struct {
	*tchttp.BaseRequest
	
	// Status. 0: all disabled; 1: all enabled
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Direction. 0: outbound; 1: inbound
	Direction *uint64 `json:"Direction,omitnil,omitempty" name:"Direction"`

	// Edge ID value
	EdgeId *string `json:"EdgeId,omitnil,omitempty" name:"EdgeId"`

	// NAT region
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`
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
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// ID of the selected firewall toggle
	FireWallVpcIds []*string `json:"FireWallVpcIds,omitnil,omitempty" name:"FireWallVpcIds"`
}

type ModifyAllVPCSwitchStatusRequest struct {
	*tchttp.BaseRequest
	
	// Status. 0: off; 1: on
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// ID of the selected firewall toggle
	FireWallVpcIds []*string `json:"FireWallVpcIds,omitnil,omitempty" name:"FireWallVpcIds"`
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
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ScanRange *int64 `json:"ScanRange,omitnil,omitempty" name:"ScanRange"`

	// Scan mode: 'heavy', 'medium', 'light'
	ScanDeep *string `json:"ScanDeep,omitnil,omitempty" name:"ScanDeep"`

	// Scan type. 1: scan now; 2: periodic scan
	RangeType *int64 `json:"RangeType,omitnil,omitempty" name:"RangeType"`

	// Scheduled task time, required when RangeType is 2
	ScanPeriod *string `json:"ScanPeriod,omitnil,omitempty" name:"ScanPeriod"`

	// Scans this field now and passes the filtered IPs
	ScanFilterIp []*string `json:"ScanFilterIp,omitnil,omitempty" name:"ScanFilterIp"`

	// 1: all; 2: single
	ScanType *int64 `json:"ScanType,omitnil,omitempty" name:"ScanType"`
}

type ModifyAssetScanRequest struct {
	*tchttp.BaseRequest
	
	// Scan range. 1: port; 2: port + vulnerability scan
	ScanRange *int64 `json:"ScanRange,omitnil,omitempty" name:"ScanRange"`

	// Scan mode: 'heavy', 'medium', 'light'
	ScanDeep *string `json:"ScanDeep,omitnil,omitempty" name:"ScanDeep"`

	// Scan type. 1: scan now; 2: periodic scan
	RangeType *int64 `json:"RangeType,omitnil,omitempty" name:"RangeType"`

	// Scheduled task time, required when RangeType is 2
	ScanPeriod *string `json:"ScanPeriod,omitnil,omitempty" name:"ScanPeriod"`

	// Scans this field now and passes the filtered IPs
	ScanFilterIp []*string `json:"ScanFilterIp,omitnil,omitempty" name:"ScanFilterIp"`

	// 1: all; 2: single
	ScanType *int64 `json:"ScanType,omitnil,omitempty" name:"ScanType"`
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
	ReturnMsg *string `json:"ReturnMsg,omitnil,omitempty" name:"ReturnMsg"`

	// Error code. 0: success; non-0: failed
	// Note: This field may return `null`, indicating that no valid value was found.
	ReturnCode *int64 `json:"ReturnCode,omitnil,omitempty" name:"ReturnCode"`

	// Status value. 0: success; 1: scanning; others: failed
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	RuleType *int64 `json:"RuleType,omitnil,omitempty" name:"RuleType"`

	// Either IP or Domain is required
	IOC []*IocListData `json:"IOC,omitnil,omitempty" name:"IOC"`

	// Optional values: delete, edit, and add
	IocAction *string `json:"IocAction,omitnil,omitempty" name:"IocAction"`

	// Time format: yyyy-MM-dd HH:mm:ss. Required when IocAction is edit or add
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time of the period in the format of yyyy-MM-dd HH:mm:ss. It must be later than both the start time and the current time. It’s required when `IocAction` is `edit` or `add`. 
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type ModifyBlockIgnoreListRequest struct {
	*tchttp.BaseRequest
	
	// Type of the rule. Values: `1` (Blocklist); `2` (Allowlist)
	RuleType *int64 `json:"RuleType,omitnil,omitempty" name:"RuleType"`

	// Either IP or Domain is required
	IOC []*IocListData `json:"IOC,omitnil,omitempty" name:"IOC"`

	// Optional values: delete, edit, and add
	IocAction *string `json:"IocAction,omitnil,omitempty" name:"IocAction"`

	// Time format: yyyy-MM-dd HH:mm:ss. Required when IocAction is edit or add
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time of the period in the format of yyyy-MM-dd HH:mm:ss. It must be later than both the start time and the current time. It’s required when `IocAction` is `edit` or `add`. 
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
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
	ReturnMsg *string `json:"ReturnMsg,omitnil,omitempty" name:"ReturnMsg"`

	// Error code. 0: success; non-0: failed
	ReturnCode *uint64 `json:"ReturnCode,omitnil,omitempty" name:"ReturnCode"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	UniqueId *string `json:"UniqueId,omitnil,omitempty" name:"UniqueId"`

	// Operation type. 1: pin to top; 0: unpin
	OpeType *string `json:"OpeType,omitnil,omitempty" name:"OpeType"`
}

type ModifyBlockTopRequest struct {
	*tchttp.BaseRequest
	
	// Record ID
	UniqueId *string `json:"UniqueId,omitnil,omitempty" name:"UniqueId"`

	// Operation type. 1: pin to top; 0: unpin
	OpeType *string `json:"OpeType,omitnil,omitempty" name:"OpeType"`
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
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type ModifyEnterpriseSecurityDispatchStatusRequest struct {
	*tchttp.BaseRequest
	
	// Status. Values: `0` (Publish now), `1` (Stop publishing)
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`
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
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	RuleUuid *uint64 `json:"RuleUuid,omitnil,omitempty" name:"RuleUuid"`

	// Modification type. Values: `0` (Modify rule content), `1` (Toggle on/off a rule) and `2` (Toggle on/off all rules)
	ModifyType *uint64 `json:"ModifyType,omitnil,omitempty" name:"ModifyType"`

	// The new rule content you want. It’s only required when you want to modify the rule content (`ModifyType=0`)
	Data *SecurityGroupRule `json:"Data,omitnil,omitempty" name:"Data"`

	// `0`: Do not enable; `1`: Enable
	Enable *uint64 `json:"Enable,omitnil,omitempty" name:"Enable"`
}

type ModifyEnterpriseSecurityGroupRuleRequest struct {
	*tchttp.BaseRequest
	
	// UUID of the rule, which can be obtained by querying the rule list
	RuleUuid *uint64 `json:"RuleUuid,omitnil,omitempty" name:"RuleUuid"`

	// Modification type. Values: `0` (Modify rule content), `1` (Toggle on/off a rule) and `2` (Toggle on/off all rules)
	ModifyType *uint64 `json:"ModifyType,omitnil,omitempty" name:"ModifyType"`

	// The new rule content you want. It’s only required when you want to modify the rule content (`ModifyType=0`)
	Data *SecurityGroupRule `json:"Data,omitnil,omitempty" name:"Data"`

	// `0`: Do not enable; `1`: Enable
	Enable *uint64 `json:"Enable,omitnil,omitempty" name:"Enable"`
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
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// ID of new rule generated after the modification
	NewRuleUuid *uint64 `json:"NewRuleUuid,omitnil,omitempty" name:"NewRuleUuid"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Rules []*CreateNatRuleItem `json:"Rules,omitnil,omitempty" name:"Rules"`
}

type ModifyNatAcRuleRequest struct {
	*tchttp.BaseRequest
	
	// Array of rules to be modified.
	Rules []*CreateNatRuleItem `json:"Rules,omitnil,omitempty" name:"Rules"`
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
	RuleUuid []*int64 `json:"RuleUuid,omitnil,omitempty" name:"RuleUuid"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Mode *int64 `json:"Mode,omitnil,omitempty" name:"Mode"`

	// Firewall instance ID
	CfwInstance *string `json:"CfwInstance,omitnil,omitempty" name:"CfwInstance"`

	// List of NAT gateways reconnected for the Using Existing mode. Only one of NatGwList and VpcList can be passed.
	NatGwList []*string `json:"NatGwList,omitnil,omitempty" name:"NatGwList"`

	// List of VPCs reconnected for the Create New mode. Only one of NatGwList and VpcList can be passed.
	VpcList []*string `json:"VpcList,omitnil,omitempty" name:"VpcList"`

	// IP range of the firewall
	FwCidrInfo *FwCidrInfo `json:"FwCidrInfo,omitnil,omitempty" name:"FwCidrInfo"`
}

type ModifyNatFwReSelectRequest struct {
	*tchttp.BaseRequest
	
	// Mode. 1: use existing; 0: create new
	Mode *int64 `json:"Mode,omitnil,omitempty" name:"Mode"`

	// Firewall instance ID
	CfwInstance *string `json:"CfwInstance,omitnil,omitempty" name:"CfwInstance"`

	// List of NAT gateways reconnected for the Using Existing mode. Only one of NatGwList and VpcList can be passed.
	NatGwList []*string `json:"NatGwList,omitnil,omitempty" name:"NatGwList"`

	// List of VPCs reconnected for the Create New mode. Only one of NatGwList and VpcList can be passed.
	VpcList []*string `json:"VpcList,omitnil,omitempty" name:"VpcList"`

	// IP range of the firewall
	FwCidrInfo *FwCidrInfo `json:"FwCidrInfo,omitnil,omitempty" name:"FwCidrInfo"`
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
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Enable *int64 `json:"Enable,omitnil,omitempty" name:"Enable"`

	// List of firewall instance IDs. Only one of CfwInsIdList, SubnetIdList, and RouteTableIdList can be passed.
	CfwInsIdList []*string `json:"CfwInsIdList,omitnil,omitempty" name:"CfwInsIdList"`

	// List of subnet IDs. Only one of CfwInsIdList, SubnetIdList, and RouteTableIdList can be passed.
	SubnetIdList []*string `json:"SubnetIdList,omitnil,omitempty" name:"SubnetIdList"`

	// List of route table IDs. Only one of CfwInsIdList, SubnetIdList, and RouteTableIdList can be passed.
	RouteTableIdList []*string `json:"RouteTableIdList,omitnil,omitempty" name:"RouteTableIdList"`
}

type ModifyNatFwSwitchRequest struct {
	*tchttp.BaseRequest
	
	// Status. 0: off; 1: on
	Enable *int64 `json:"Enable,omitnil,omitempty" name:"Enable"`

	// List of firewall instance IDs. Only one of CfwInsIdList, SubnetIdList, and RouteTableIdList can be passed.
	CfwInsIdList []*string `json:"CfwInsIdList,omitnil,omitempty" name:"CfwInsIdList"`

	// List of subnet IDs. Only one of CfwInsIdList, SubnetIdList, and RouteTableIdList can be passed.
	SubnetIdList []*string `json:"SubnetIdList,omitnil,omitempty" name:"SubnetIdList"`

	// List of route table IDs. Only one of CfwInsIdList, SubnetIdList, and RouteTableIdList can be passed.
	RouteTableIdList []*string `json:"RouteTableIdList,omitnil,omitempty" name:"RouteTableIdList"`
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
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	NatFwInsId *string `json:"NatFwInsId,omitnil,omitempty" name:"NatFwInsId"`

	// DNS toggle list
	DnsVpcSwitchLst []*DnsVpcSwitch `json:"DnsVpcSwitchLst,omitnil,omitempty" name:"DnsVpcSwitchLst"`
}

type ModifyNatFwVpcDnsSwitchRequest struct {
	*tchttp.BaseRequest
	
	// NAT firewall ID
	NatFwInsId *string `json:"NatFwInsId,omitnil,omitempty" name:"NatFwInsId"`

	// DNS toggle list
	DnsVpcSwitchLst []*DnsVpcSwitch `json:"DnsVpcSwitchLst,omitnil,omitempty" name:"DnsVpcSwitchLst"`
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
	ReturnMsg *string `json:"ReturnMsg,omitnil,omitempty" name:"ReturnMsg"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	RuleChangeItems []*RuleChangeItem `json:"RuleChangeItems,omitnil,omitempty" name:"RuleChangeItems"`

	// Rule direction. Values: `1` (Inbound) and `0` (Outbound)
	Direction *uint64 `json:"Direction,omitnil,omitempty" name:"Direction"`
}

type ModifyNatSequenceRulesRequest struct {
	*tchttp.BaseRequest
	
	// Rule sequence number. Values: `OrderIndex` (Original sequence number), `NewOrderIndex` (New sequence number)
	RuleChangeItems []*RuleChangeItem `json:"RuleChangeItems,omitnil,omitempty" name:"RuleChangeItems"`

	// Rule direction. Values: `1` (Inbound) and `0` (Outbound)
	Direction *uint64 `json:"Direction,omitnil,omitempty" name:"Direction"`
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
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	FireWallPublicIP *string `json:"FireWallPublicIP,omitnil,omitempty" name:"FireWallPublicIP"`

	// Status value. 0: off; 1: on
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type ModifyPublicIPSwitchStatusRequest struct {
	*tchttp.BaseRequest
	
	// Public IP
	FireWallPublicIP *string `json:"FireWallPublicIP,omitnil,omitempty" name:"FireWallPublicIP"`

	// Status value. 0: off; 1: on
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
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
	ReturnMsg *string `json:"ReturnMsg,omitnil,omitempty" name:"ReturnMsg"`

	// Error code. 0: success; non-0: failed
	ReturnCode *int64 `json:"ReturnCode,omitnil,omitempty" name:"ReturnCode"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// Group name
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// Parent group ID
	ParentId *string `json:"ParentId,omitnil,omitempty" name:"ParentId"`
}

type ModifyResourceGroupRequest struct {
	*tchttp.BaseRequest
	
	// Group ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// Group name
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// Parent group ID
	ParentId *string `json:"ParentId,omitnil,omitempty" name:"ParentId"`
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
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`
}

type ModifyRunSyncAssetRequest struct {
	*tchttp.BaseRequest
	
	// 0: edge firewall toggle; 1: VPC firewall toggle
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`
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
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Direction *uint64 `json:"Direction,omitnil,omitempty" name:"Direction"`

	// Toggle status. 0: off; 1: on
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Modified priority of enterprise security group rules
	RuleSequence *uint64 `json:"RuleSequence,omitnil,omitempty" name:"RuleSequence"`
}

type ModifySecurityGroupItemRuleStatusRequest struct {
	*tchttp.BaseRequest
	
	// Direction. 0: outbound; 1: inbound. 1 by default
	Direction *uint64 `json:"Direction,omitnil,omitempty" name:"Direction"`

	// Toggle status. 0: off; 1: on
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Modified priority of enterprise security group rules
	RuleSequence *uint64 `json:"RuleSequence,omitnil,omitempty" name:"RuleSequence"`
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
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Direction *uint64 `json:"Direction,omitnil,omitempty" name:"Direction"`

	// Sorting data of enterprise security group rules
	Data []*SecurityGroupOrderIndexData `json:"Data,omitnil,omitempty" name:"Data"`
}

type ModifySecurityGroupSequenceRulesRequest struct {
	*tchttp.BaseRequest
	
	// Direction. 0: outbound; 1: inbound. 1 by default
	Direction *uint64 `json:"Direction,omitnil,omitempty" name:"Direction"`

	// Sorting data of enterprise security group rules
	Data []*SecurityGroupOrderIndexData `json:"Data,omitnil,omitempty" name:"Data"`
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
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	EdgeId *string `json:"EdgeId,omitnil,omitempty" name:"EdgeId"`

	// Modifies data
	Data []*SequenceData `json:"Data,omitnil,omitempty" name:"Data"`

	// NAT region
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// Direction. 0: outbound; 1: inbound
	Direction *uint64 `json:"Direction,omitnil,omitempty" name:"Direction"`
}

type ModifySequenceRulesRequest struct {
	*tchttp.BaseRequest
	
	// Edge ID value
	EdgeId *string `json:"EdgeId,omitnil,omitempty" name:"EdgeId"`

	// Modifies data
	Data []*SequenceData `json:"Data,omitnil,omitempty" name:"Data"`

	// NAT region
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// Direction. 0: outbound; 1: inbound
	Direction *uint64 `json:"Direction,omitnil,omitempty" name:"Direction"`
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
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	EdgeId *string `json:"EdgeId,omitnil,omitempty" name:"EdgeId"`

	// Status value. 1: table locked; 2: table unlocked
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// NAT region
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 0: outbound; 1: inbound
	Direction *uint64 `json:"Direction,omitnil,omitempty" name:"Direction"`
}

type ModifyTableStatusRequest struct {
	*tchttp.BaseRequest
	
	// Edge ID between two VPCs
	EdgeId *string `json:"EdgeId,omitnil,omitempty" name:"EdgeId"`

	// Status value. 1: table locked; 2: table unlocked
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// NAT region
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 0: outbound; 1: inbound
	Direction *uint64 `json:"Direction,omitnil,omitempty" name:"Direction"`
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
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

type MultiTopicSearchInformation struct {
	// ID of the log topic to be searched and analyzed
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// Pass the Context value returned by the last API call to retrieve more subsequent logs. A total of up to 10,000 raw logs can be obtained, with a validity period of 1 hour.
	Context *string `json:"Context,omitnil,omitempty" name:"Context"`
}

type NatFwFilter struct {
	// Filter type, e.g., instance ID
	FilterType *string `json:"FilterType,omitnil,omitempty" name:"FilterType"`

	// Filtered content, separated with ","
	FilterContent *string `json:"FilterContent,omitnil,omitempty" name:"FilterContent"`
}

type NatFwInstance struct {
	// NAT instance ID
	NatinsId *string `json:"NatinsId,omitnil,omitempty" name:"NatinsId"`

	// NAT instance name
	NatinsName *string `json:"NatinsName,omitnil,omitempty" name:"NatinsName"`

	// Instance region
	// Note: This field may return `null`, indicating that no valid value was found.
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 0: create new; 1: use existing
	// Note: This field may return `null`, indicating that no valid value was found.
	FwMode *int64 `json:"FwMode,omitnil,omitempty" name:"FwMode"`

	// 0: normal; 1: creating
	// Note: This field may return `null`, indicating that no valid value was found.
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// NAT public IP
	// Note: This field may return `null`, indicating that no valid value was found.
	NatIp *string `json:"NatIp,omitnil,omitempty" name:"NatIp"`
}

type NatInstanceInfo struct {
	// NAT instance ID
	NatinsId *string `json:"NatinsId,omitnil,omitempty" name:"NatinsId"`

	// NAT instance name
	NatinsName *string `json:"NatinsName,omitnil,omitempty" name:"NatinsName"`

	// Instance region
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 0: create new; 1: use existing
	FwMode *int64 `json:"FwMode,omitnil,omitempty" name:"FwMode"`

	// Instance bandwidth (Mbps)
	BandWidth *int64 `json:"BandWidth,omitnil,omitempty" name:"BandWidth"`

	// Inbound traffic peak bandwidth (bps)
	InFlowMax *int64 `json:"InFlowMax,omitnil,omitempty" name:"InFlowMax"`

	// Outbound traffic peak bandwidth (bps)
	OutFlowMax *uint64 `json:"OutFlowMax,omitnil,omitempty" name:"OutFlowMax"`

	// Chinese region information
	RegionZh *string `json:"RegionZh,omitnil,omitempty" name:"RegionZh"`

	// Public IP array
	// Note: This field may return `null`, indicating that no valid value was found.
	EipAddress []*string `json:"EipAddress,omitnil,omitempty" name:"EipAddress"`

	// Array of internal and external IPs
	// Note: This field may return `null`, indicating that no valid value was found.
	VpcIp []*string `json:"VpcIp,omitnil,omitempty" name:"VpcIp"`

	// Array of subnets associated with an instance
	// Note: This field may return `null`, indicating that no valid value was found.
	Subnets []*string `json:"Subnets,omitnil,omitempty" name:"Subnets"`

	// 0: normal 1: initializing
	// Note: This field may return `null`, indicating that no valid value was found.
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Region information
	// Note: This field may return null, indicating that no valid values can be obtained.
	RegionDetail *string `json:"RegionDetail,omitnil,omitempty" name:"RegionDetail"`

	// Availability zone of the instance
	// Note: This field may return null, indicating that no valid values can be obtained.
	ZoneZh *string `json:"ZoneZh,omitnil,omitempty" name:"ZoneZh"`

	// Availability zone of the instance
	// Note: This field may return null, indicating that no valid values can be obtained.
	ZoneZhBak *string `json:"ZoneZhBak,omitnil,omitempty" name:"ZoneZhBak"`

	// Number of used rules.
	// Note: This field may return null, indicating that no valid values can be obtained.
	RuleUsed *uint64 `json:"RuleUsed,omitnil,omitempty" name:"RuleUsed"`

	// The maximum number of rules allowed in the instance.
	// Note: This field may return null, indicating that no valid values can be obtained.
	RuleMax *uint64 `json:"RuleMax,omitnil,omitempty" name:"RuleMax"`
}

type NewModeItems struct {
	// VPC list for the Create New mode
	VpcList []*string `json:"VpcList,omitnil,omitempty" name:"VpcList"`

	// The list of egress public EIPs bound for the Create New mode. Either Eips or AddCount is required.
	Eips []*string `json:"Eips,omitnil,omitempty" name:"Eips"`

	// The number of egress public EIPs newly bound for the Create New mode. Either Eips or AddCount is required.
	AddCount *int64 `json:"AddCount,omitnil,omitempty" name:"AddCount"`
}

// Predefined struct for user
type RemoveAcRuleRequestParams struct {
	// UUID of the rule, which can be obtained by querying the rule list
	RuleUuid *int64 `json:"RuleUuid,omitnil,omitempty" name:"RuleUuid"`
}

type RemoveAcRuleRequest struct {
	*tchttp.BaseRequest
	
	// UUID of the rule, which can be obtained by querying the rule list
	RuleUuid *int64 `json:"RuleUuid,omitnil,omitempty" name:"RuleUuid"`
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
	RuleUuid *int64 `json:"RuleUuid,omitnil,omitempty" name:"RuleUuid"`

	// 0: operation successful; -1: operation failed
	// Note: This field may return `null`, indicating that no valid value was found.
	ReturnCode *int64 `json:"ReturnCode,omitnil,omitempty" name:"ReturnCode"`

	// success: operation successful; failed: operation failed
	// Note: This field may return `null`, indicating that no valid value was found.
	ReturnMsg *string `json:"ReturnMsg,omitnil,omitempty" name:"ReturnMsg"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	RuleUuid *int64 `json:"RuleUuid,omitnil,omitempty" name:"RuleUuid"`

	// Type of deletion. 0: delete a single entry, and enter ID of the deleted rule for RuleUuid; 1: delete all, and enter 0 for RuleUuid
	RemoveType *int64 `json:"RemoveType,omitnil,omitempty" name:"RemoveType"`
}

type RemoveEnterpriseSecurityGroupRuleRequest struct {
	*tchttp.BaseRequest
	
	// UUID of the rule, which can be obtained by querying the rule list
	RuleUuid *int64 `json:"RuleUuid,omitnil,omitempty" name:"RuleUuid"`

	// Type of deletion. 0: delete a single entry, and enter ID of the deleted rule for RuleUuid; 1: delete all, and enter 0 for RuleUuid
	RemoveType *int64 `json:"RemoveType,omitnil,omitempty" name:"RemoveType"`
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
	RuleUuid *int64 `json:"RuleUuid,omitnil,omitempty" name:"RuleUuid"`

	// 0: operation successful; -1: operation failed
	// Note: This field may return `null`, indicating that no valid value was found.
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	RuleUuid []*int64 `json:"RuleUuid,omitnil,omitempty" name:"RuleUuid"`

	// Rule direction. Valid values: 1: inbound; 0: outbound.
	Direction *uint64 `json:"Direction,omitnil,omitempty" name:"Direction"`
}

type RemoveNatAcRuleRequest struct {
	*tchttp.BaseRequest
	
	// UUIDs of the rules to delete, which can be obtained by querying the rule list. Note: If [-1] is passed in, all rules are deleted.
	RuleUuid []*int64 `json:"RuleUuid,omitnil,omitempty" name:"RuleUuid"`

	// Rule direction. Valid values: 1: inbound; 0: outbound.
	Direction *uint64 `json:"Direction,omitnil,omitempty" name:"Direction"`
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
	RuleUuid []*int64 `json:"RuleUuid,omitnil,omitempty" name:"RuleUuid"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	OrderIndex *int64 `json:"OrderIndex,omitnil,omitempty" name:"OrderIndex"`

	// New sequence number
	NewOrderIndex *int64 `json:"NewOrderIndex,omitnil,omitempty" name:"NewOrderIndex"`
}

type RuleInfoData struct {
	// Priority
	OrderIndex *uint64 `json:"OrderIndex,omitnil,omitempty" name:"OrderIndex"`

	// Access source
	SourceIp *string `json:"SourceIp,omitnil,omitempty" name:"SourceIp"`

	// Access destination
	TargetIp *string `json:"TargetIp,omitnil,omitempty" name:"TargetIp"`

	// Protocol
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// Policy. 0: observe; 1: block; 2: allow
	Strategy *string `json:"Strategy,omitnil,omitempty" name:"Strategy"`

	// Access source type. 1: IP; 3: domain name; 4: IP address template; 5: domain name address template
	SourceType *uint64 `json:"SourceType,omitnil,omitempty" name:"SourceType"`

	// Direction. 0: outbound; 1: inbound
	Direction *uint64 `json:"Direction,omitnil,omitempty" name:"Direction"`

	// Description
	Detail *string `json:"Detail,omitnil,omitempty" name:"Detail"`

	// Access destination type. 1: IP, 3: domain name; 4: IP address template; 5: domain name address template
	TargetType *uint64 `json:"TargetType,omitnil,omitempty" name:"TargetType"`

	// Port
	Port *string `json:"Port,omitnil,omitempty" name:"Port"`

	// ID value
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// Log ID, required when an alert log is created
	LogId *string `json:"LogId,omitnil,omitempty" name:"LogId"`

	// City code
	City *uint64 `json:"City,omitnil,omitempty" name:"City"`

	// Country code
	Country *uint64 `json:"Country,omitnil,omitempty" name:"Country"`

	// Cloud vendor. Multiple vendors are supported and separated with commas. 1: Tencent Cloud (only in Hong Kong, China and overseas); 2: Alibaba Cloud; 3: Amazon Cloud; 4: Huawei Cloud; 5: Microsoft Cloud
	CloudCode *string `json:"CloudCode,omitnil,omitempty" name:"CloudCode"`

	// Indicates whether it is a region
	IsRegion *uint64 `json:"IsRegion,omitnil,omitempty" name:"IsRegion"`

	// City name
	CityName *string `json:"CityName,omitnil,omitempty" name:"CityName"`

	// Country name
	CountryName *string `json:"CountryName,omitnil,omitempty" name:"CountryName"`
}

type ScanInfo struct {
	// Scanning result information
	ScanResultInfo *ScanResultInfo `json:"ScanResultInfo,omitnil,omitempty" name:"ScanResultInfo"`

	// Scanning status. 0: scanning; 1: completed; 2: auto scanning unselected
	ScanStatus *int64 `json:"ScanStatus,omitnil,omitempty" name:"ScanStatus"`

	// Progress
	ScanPercent *float64 `json:"ScanPercent,omitnil,omitempty" name:"ScanPercent"`

	// Estimated completion time
	ScanTime *string `json:"ScanTime,omitnil,omitempty" name:"ScanTime"`
}

type ScanResultInfo struct {
	// Number of vulnerability exploits
	LeakNum *uint64 `json:"LeakNum,omitnil,omitempty" name:"LeakNum"`

	// Number of protected IPs
	IPNum *uint64 `json:"IPNum,omitnil,omitempty" name:"IPNum"`

	// Number of exposed ports
	PortNum *uint64 `json:"PortNum,omitnil,omitempty" name:"PortNum"`

	// Protection status
	IPStatus *bool `json:"IPStatus,omitnil,omitempty" name:"IPStatus"`

	// Attack blocking status
	IdpStatus *bool `json:"IdpStatus,omitnil,omitempty" name:"IdpStatus"`

	// Port blocking status
	BanStatus *bool `json:"BanStatus,omitnil,omitempty" name:"BanStatus"`
}

type SearchLogErrors struct {
	// Log topic ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// Error message
	// Note: This field may return null, indicating that no valid values can be obtained.
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// Error code.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ErrorCodeStr *string `json:"ErrorCodeStr,omitnil,omitempty" name:"ErrorCodeStr"`
}

type SearchLogInfos struct {
	// Log topic ID
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// Log storage lifetime
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// Pass through the Context value returned by this API, which can access more logs later, with an expiration time of 1 hour.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Context *string `json:"Context,omitnil,omitempty" name:"Context"`
}

// Predefined struct for user
type SearchLogRequestParams struct {
	// <p>Start time for logs to be searched and analyzed, which is a Unix timestamp in milliseconds</p>
	From *int64 `json:"From,omitnil,omitempty" name:"From"`

	// <p>End time for logs to be searched and analyzed, which is a Unix timestamp in milliseconds</p>
	To *int64 `json:"To,omitnil,omitempty" name:"To"`

	// <p>The retrieval and analysis statement has a maximum length of 12 KB.<br>The statement consists of <a href="https://www.tencentcloud.com/document/product/614/47044?from_cn_redirect=1" target="_blank">[retrieval condition]</a> | <a href="https://www.tencentcloud.com/document/product/614/44061?from_cn_redirect=1" target="_blank">[SQL statement]</a>. When there is no need to perform statistical analysis on logs, the pipe character <code> | </code> and the SQL statement can be omitted.<br>Use * or an empty string to search all logs.</p>
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// <p>Search syntax rules. Default value is 0. Recommended for use is 1.</p><ul><li>0: Lucene syntax</li><li>1: CQL syntax (dedicated retrieval syntax for CLS, also the default syntax rule used in the console).</li></ul><p>For details, see <a href="https://www.tencentcloud.com/document/product/614/47044?from_cn_redirect=1#RetrievesConditionalRules" target="_blank">Retrieval condition syntax rules</a></p>
	SyntaxRule *uint64 `json:"SyntaxRule,omitnil,omitempty" name:"SyntaxRule"`

	// <ul><li>Log topic ID to be retrieved and analyzed. Only one log topic can be specified.</li><li>If needed, use the Topics parameter to retrieve multiple log topics.</li><li>The TopicId and Topics parameters cannot be used simultaneously. Only one can be selected in a single request.<br>The log topic IDs are as follows:<br>Access control - Internet boundary: cfw_rule_acl<br>Access control - NAT boundary: cfw_rule_nat_acl<br>Access control - VPC boundary: cfw_rule_vpc_acl<br>Access control - DNS switch: cfw_rule_dns_acl<br>Intrusion defense: cfw_rule_threatinfo<br>Full traffic detection and response logs - Traffic analysis: cfw_netflow_nta<br>Full traffic detection and response logs - Traffic alarm: cfw_rule_ndr_threatinfo<br>Zero trust operations and maintenance - Database logon: cfw_operate_db<br>Zero trust operations and maintenance - Server access: operate_remote_om<br>Zero trust operations and maintenance - Web service access: operate_web_access<br>Zero trust operations and maintenance - Behavioral audit: remoteom_commands<br>Traffic log - Internet boundary: cfw_netflow_border<br>Traffic log - NAT boundary: cfw_netflow_nat<br>Traffic log - VPC boundary: cfw_netflow_vpc<br>Traffic log - DNS switch: cfw_netflow_dns<br>Traffic log - Private network traffic: cfw_netflow_fl<br>Operation log: operate_log_all</li></ul>
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// <ul><li>Log topic list for retrieval and analysis, supports a maximum of 50 log topics.</li><li>Use TopicId to retrieve a single log topic.</li><li>TopicId and Topics cannot be used simultaneously. Only select one in a single request.</li></ul>
	Topics []*MultiTopicSearchInformation `json:"Topics,omitnil,omitempty" name:"Topics"`

	// <p>Whether raw logs are returned in time sequence; value range: asc (ascending), desc (descending), default is desc<br>Note:</p><ul><li>Valid only when the search and analysis statement (Query) does not contain SQL</li><li>For SQL result sorting, refer to <a href="https://www.tencentcloud.com/document/product/614/58978?from_cn_redirect=1" target="_blank">SQL ORDER BY syntax</a></li></ul>
	Sort *string `json:"Sort,omitnil,omitempty" name:"Sort"`

	// <p>Number of raw logs returned in a single query. Default value: 100. Maximum value: 1000.<br>Note:</p><ul><li>This parameter is valid only when the search and analysis statement (Query) does not contain SQL.</li><li>For the method for specifying SQL result count, see <a href="https://www.tencentcloud.com/document/product/614/58977?from_cn_redirect=1" target="_blank">SQL LIMIT syntax</a>.</li></ul><p>You can retrieve more logs in two ways:</p><ul><li>Context: Pass the Context value returned by the last API call to retrieve more logs. You can retrieve up to 10,000 entries of raw logs in total.</li><li>Offset: The offset indicates the line number from which to start returning raw logs. There is no log entry limit.</li></ul>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>Query the offset of raw logs, indicating the line number from which to start returning raw logs. Default value is 0.<br>Note:</p><ul><li>Applicable only when the retrieval and analysis statement (Query) does not contain SQL.</li><li>Cannot be used with the Context parameter.</li><li>Applicable only for single log topic retrieval.</li></ul>
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>Pass the Context value returned by the last API call to obtain more logs later. The total number of raw logs that can be obtained is up to 10,000 entries. The expiration time is 1 hour.<br>Note:</p><ul><li>When passing this parameter, do not modify other parameters.</li><li>Applicable only for single log topic retrieval. To retrieve multiple log topics, use the Context in Topics.</li><li>This is valid only when the search and analysis statement (Query) does not contain SQL. For obtaining subsequent results with SQL, refer to <a href="https://www.tencentcloud.com/document/product/614/58977?from_cn_redirect=1" target="_blank">SQL LIMIT syntax</a>.</li></ul>
	Context *string `json:"Context,omitnil,omitempty" name:"Context"`

	// <p>When performing statistical analysis (SQL included in Query), whether to sample raw logs first and then perform statistical analysis.<br>0: Automatic sampling;<br>0–1: Sample at the specified sampling rate, for example 0.02;<br>1: Indicates no sampling, that is, precision analysis.<br>Default value: 1</p>
	SamplingRate *float64 `json:"SamplingRate,omitnil,omitempty" name:"SamplingRate"`

	// <p>true means using the new retrieval result return method, and output parameters AnalysisRecords and Columns are valid.<br>false means using the old retrieval result return method, and output parameters AnalysisResults and ColNames are valid.<br>The two return methods have a slight difference in encoding format. It is recommended to use true.</p>
	UseNewAnalysis *bool `json:"UseNewAnalysis,omitnil,omitempty" name:"UseNewAnalysis"`
}

type SearchLogRequest struct {
	*tchttp.BaseRequest
	
	// <p>Start time for logs to be searched and analyzed, which is a Unix timestamp in milliseconds</p>
	From *int64 `json:"From,omitnil,omitempty" name:"From"`

	// <p>End time for logs to be searched and analyzed, which is a Unix timestamp in milliseconds</p>
	To *int64 `json:"To,omitnil,omitempty" name:"To"`

	// <p>The retrieval and analysis statement has a maximum length of 12 KB.<br>The statement consists of <a href="https://www.tencentcloud.com/document/product/614/47044?from_cn_redirect=1" target="_blank">[retrieval condition]</a> | <a href="https://www.tencentcloud.com/document/product/614/44061?from_cn_redirect=1" target="_blank">[SQL statement]</a>. When there is no need to perform statistical analysis on logs, the pipe character <code> | </code> and the SQL statement can be omitted.<br>Use * or an empty string to search all logs.</p>
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// <p>Search syntax rules. Default value is 0. Recommended for use is 1.</p><ul><li>0: Lucene syntax</li><li>1: CQL syntax (dedicated retrieval syntax for CLS, also the default syntax rule used in the console).</li></ul><p>For details, see <a href="https://www.tencentcloud.com/document/product/614/47044?from_cn_redirect=1#RetrievesConditionalRules" target="_blank">Retrieval condition syntax rules</a></p>
	SyntaxRule *uint64 `json:"SyntaxRule,omitnil,omitempty" name:"SyntaxRule"`

	// <ul><li>Log topic ID to be retrieved and analyzed. Only one log topic can be specified.</li><li>If needed, use the Topics parameter to retrieve multiple log topics.</li><li>The TopicId and Topics parameters cannot be used simultaneously. Only one can be selected in a single request.<br>The log topic IDs are as follows:<br>Access control - Internet boundary: cfw_rule_acl<br>Access control - NAT boundary: cfw_rule_nat_acl<br>Access control - VPC boundary: cfw_rule_vpc_acl<br>Access control - DNS switch: cfw_rule_dns_acl<br>Intrusion defense: cfw_rule_threatinfo<br>Full traffic detection and response logs - Traffic analysis: cfw_netflow_nta<br>Full traffic detection and response logs - Traffic alarm: cfw_rule_ndr_threatinfo<br>Zero trust operations and maintenance - Database logon: cfw_operate_db<br>Zero trust operations and maintenance - Server access: operate_remote_om<br>Zero trust operations and maintenance - Web service access: operate_web_access<br>Zero trust operations and maintenance - Behavioral audit: remoteom_commands<br>Traffic log - Internet boundary: cfw_netflow_border<br>Traffic log - NAT boundary: cfw_netflow_nat<br>Traffic log - VPC boundary: cfw_netflow_vpc<br>Traffic log - DNS switch: cfw_netflow_dns<br>Traffic log - Private network traffic: cfw_netflow_fl<br>Operation log: operate_log_all</li></ul>
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// <ul><li>Log topic list for retrieval and analysis, supports a maximum of 50 log topics.</li><li>Use TopicId to retrieve a single log topic.</li><li>TopicId and Topics cannot be used simultaneously. Only select one in a single request.</li></ul>
	Topics []*MultiTopicSearchInformation `json:"Topics,omitnil,omitempty" name:"Topics"`

	// <p>Whether raw logs are returned in time sequence; value range: asc (ascending), desc (descending), default is desc<br>Note:</p><ul><li>Valid only when the search and analysis statement (Query) does not contain SQL</li><li>For SQL result sorting, refer to <a href="https://www.tencentcloud.com/document/product/614/58978?from_cn_redirect=1" target="_blank">SQL ORDER BY syntax</a></li></ul>
	Sort *string `json:"Sort,omitnil,omitempty" name:"Sort"`

	// <p>Number of raw logs returned in a single query. Default value: 100. Maximum value: 1000.<br>Note:</p><ul><li>This parameter is valid only when the search and analysis statement (Query) does not contain SQL.</li><li>For the method for specifying SQL result count, see <a href="https://www.tencentcloud.com/document/product/614/58977?from_cn_redirect=1" target="_blank">SQL LIMIT syntax</a>.</li></ul><p>You can retrieve more logs in two ways:</p><ul><li>Context: Pass the Context value returned by the last API call to retrieve more logs. You can retrieve up to 10,000 entries of raw logs in total.</li><li>Offset: The offset indicates the line number from which to start returning raw logs. There is no log entry limit.</li></ul>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>Query the offset of raw logs, indicating the line number from which to start returning raw logs. Default value is 0.<br>Note:</p><ul><li>Applicable only when the retrieval and analysis statement (Query) does not contain SQL.</li><li>Cannot be used with the Context parameter.</li><li>Applicable only for single log topic retrieval.</li></ul>
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>Pass the Context value returned by the last API call to obtain more logs later. The total number of raw logs that can be obtained is up to 10,000 entries. The expiration time is 1 hour.<br>Note:</p><ul><li>When passing this parameter, do not modify other parameters.</li><li>Applicable only for single log topic retrieval. To retrieve multiple log topics, use the Context in Topics.</li><li>This is valid only when the search and analysis statement (Query) does not contain SQL. For obtaining subsequent results with SQL, refer to <a href="https://www.tencentcloud.com/document/product/614/58977?from_cn_redirect=1" target="_blank">SQL LIMIT syntax</a>.</li></ul>
	Context *string `json:"Context,omitnil,omitempty" name:"Context"`

	// <p>When performing statistical analysis (SQL included in Query), whether to sample raw logs first and then perform statistical analysis.<br>0: Automatic sampling;<br>0–1: Sample at the specified sampling rate, for example 0.02;<br>1: Indicates no sampling, that is, precision analysis.<br>Default value: 1</p>
	SamplingRate *float64 `json:"SamplingRate,omitnil,omitempty" name:"SamplingRate"`

	// <p>true means using the new retrieval result return method, and output parameters AnalysisRecords and Columns are valid.<br>false means using the old retrieval result return method, and output parameters AnalysisResults and ColNames are valid.<br>The two return methods have a slight difference in encoding format. It is recommended to use true.</p>
	UseNewAnalysis *bool `json:"UseNewAnalysis,omitnil,omitempty" name:"UseNewAnalysis"`
}

func (r *SearchLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "From")
	delete(f, "To")
	delete(f, "Query")
	delete(f, "SyntaxRule")
	delete(f, "TopicId")
	delete(f, "Topics")
	delete(f, "Sort")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Context")
	delete(f, "SamplingRate")
	delete(f, "UseNewAnalysis")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SearchLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SearchLogResponseParams struct {
	// <p>Pass through the Context value returned by this API to obtain more logs later. The expiration time is 1 hour.<br>Note:</p><ul><li>Applicable only for single log topic retrieval. To retrieve multiple log topics, use the Context in Topics.</li></ul>
	Context *string `json:"Context,omitnil,omitempty" name:"Context"`

	// <p>Whether all logs meeting the retrieval criteria have been returned. If not, use Context parameter to retrieve more logs.<br>  <br>Note: This is only valid when the search and analysis statement (Query) does not contain SQL.</p>
	ListOver *bool `json:"ListOver,omitnil,omitempty" name:"ListOver"`

	// <p>Whether the returned data is the SQL analysis result</p>
	Analysis *bool `json:"Analysis,omitnil,omitempty" name:"Analysis"`

	// <p>Raw logs matching the retrieval criteria</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	Results []*LogInfo `json:"Results,omitnil,omitempty" name:"Results"`

	// <p>Column names of log statistics analysis results<br>This parameter is valid only when UseNewAnalysis is false.</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	ColNames []*string `json:"ColNames,omitnil,omitempty" name:"ColNames"`

	// <p>Log statistics and analysis results<br>This parameter is valid only when UseNewAnalysis is false.</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	AnalysisResults []*LogItems `json:"AnalysisResults,omitnil,omitempty" name:"AnalysisResults"`

	// <p>Log statistics and analysis results<br>This parameter is valid only when UseNewAnalysis is true.</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	AnalysisRecords []*string `json:"AnalysisRecords,omitnil,omitempty" name:"AnalysisRecords"`

	// <p>Column attribute of the statistical analysis result<br>This parameter is valid only when UseNewAnalysis is true.</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	Columns []*Column `json:"Columns,omitnil,omitempty" name:"Columns"`

	// <p>Sampling rate used for this statistical analysis</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	SamplingRate *float64 `json:"SamplingRate,omitnil,omitempty" name:"SamplingRate"`

	// <p>When multiple log topics are used for retrieval, basic information of each log topic, such as error message.</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	Topics *SearchLogTopics `json:"Topics,omitnil,omitempty" name:"Topics"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SearchLogResponse struct {
	*tchttp.BaseResponse
	Response *SearchLogResponseParams `json:"Response"`
}

func (r *SearchLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchLogTopics struct {
	// Error information corresponding to multi-log topic retrieval
	// Note: This field may return null, indicating that no valid values can be obtained.
	Errors []*SearchLogErrors `json:"Errors,omitnil,omitempty" name:"Errors"`

	// Information for each log topic during multi-log topic retrieval
	// Note: This field may return null, indicating that no valid values can be obtained.
	Infos []*SearchLogInfos `json:"Infos,omitnil,omitempty" name:"Infos"`
}

type SecurityGroupBothWayInfo struct {
	// Priority
	// Note: This field may return `null`, indicating that no valid value was found.
	OrderIndex *uint64 `json:"OrderIndex,omitnil,omitempty" name:"OrderIndex"`

	// Access source
	// Note: This field may return `null`, indicating that no valid value was found.
	SourceId *string `json:"SourceId,omitnil,omitempty" name:"SourceId"`

	// Access source type. Default: 0. 0: IP; 1: VPC; 2: SUBNET; 3: CVM; 4: CLB; 5: ENI; 6: CDB; 7: Parameter template. 100: asset group
	// Note: This field may return `null`, indicating that no valid value was found.
	SourceType *uint64 `json:"SourceType,omitnil,omitempty" name:"SourceType"`

	// Access destination
	// Note: This field may return `null`, indicating that no valid value was found.
	TargetId *string `json:"TargetId,omitnil,omitempty" name:"TargetId"`

	// Access destination type. Default: 0. 0: IP; 1: VPC; 2: SUBNET; 3: CVM; 4: CLB; 5: ENI; 6: CDB; 7: Parameter template. 100: asset group
	// Note: This field may return `null`, indicating that no valid value was found.
	TargetType *uint64 `json:"TargetType,omitnil,omitempty" name:"TargetType"`

	// Protocol
	// Note: This field may return `null`, indicating that no valid value was found.
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// Destination port
	// Note: This field may return `null`, indicating that no valid value was found.
	Port *string `json:"Port,omitnil,omitempty" name:"Port"`

	// Policy. 1: block; 2: allow
	// Note: This field may return `null`, indicating that no valid value was found.
	Strategy *uint64 `json:"Strategy,omitnil,omitempty" name:"Strategy"`

	// Direction. 0: outbound; 1: inbound. 1 by default
	// Note: This field may return `null`, indicating that no valid value was found.
	Direction *uint64 `json:"Direction,omitnil,omitempty" name:"Direction"`

	// Region
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Description
	// Note: This field may return `null`, indicating that no valid value was found.
	Detail *string `json:"Detail,omitnil,omitempty" name:"Detail"`

	// Toggle status. 0: off; 1: on
	// Note: This field may return `null`, indicating that no valid value was found.
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Indicates whether the rule is normal. 0: normal; 1: abnormal
	// Note: This field may return `null`, indicating that no valid value was found.
	IsNew *uint64 `json:"IsNew,omitnil,omitempty" name:"IsNew"`

	// One-way/two-way. 0: one-way; 1: two-way
	// Note: This field may return `null`, indicating that no valid value was found.
	BothWay *uint64 `json:"BothWay,omitnil,omitempty" name:"BothWay"`

	// VPC ID
	// Note: This field may return `null`, indicating that no valid value was found.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Subnet ID
	// Note: This field may return `null`, indicating that no valid value was found.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Instance name
	// Note: This field may return `null`, indicating that no valid value was found.
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Public IP. Multiple IPs are separated by commas.
	// Note: This field may return `null`, indicating that no valid value was found.
	PublicIp *string `json:"PublicIp,omitnil,omitempty" name:"PublicIp"`

	// Private IP. Multiple IPs are separated by commas.
	// Note: This field may return `null`, indicating that no valid value was found.
	PrivateIp *string `json:"PrivateIp,omitnil,omitempty" name:"PrivateIp"`

	// Masked address. Multiple addresses are separated by commas.
	// Note: This field may return `null`, indicating that no valid value was found.
	Cidr *string `json:"Cidr,omitnil,omitempty" name:"Cidr"`

	// Port protocol template ID
	// Note: This field may return `null`, indicating that no valid value was found.
	ServiceTemplateId *string `json:"ServiceTemplateId,omitnil,omitempty" name:"ServiceTemplateId"`

	// Indicates whether to use the port protocol template. 0: no; 1: yes
	ProtocolPortType *uint64 `json:"ProtocolPortType,omitnil,omitempty" name:"ProtocolPortType"`
}

type SecurityGroupListData struct {
	// Priority
	OrderIndex *uint64 `json:"OrderIndex,omitnil,omitempty" name:"OrderIndex"`

	// Access source
	SourceId *string `json:"SourceId,omitnil,omitempty" name:"SourceId"`

	// Access source type. Default: 0. 1: VPC; 2: SUBNET; 3: CVM; 4: CLB; 5: ENI; 6: CDB; 7: Parameter template. 100: Resource group
	SourceType *uint64 `json:"SourceType,omitnil,omitempty" name:"SourceType"`

	// Access destination
	TargetId *string `json:"TargetId,omitnil,omitempty" name:"TargetId"`

	// Access destination type. Default: 0. 1: VPC; 2: SUBNET; 3: CVM; 4: CLB; 5: ENI; 6: CDB; 7: Parameter template; 100: resource group
	TargetType *uint64 `json:"TargetType,omitnil,omitempty" name:"TargetType"`

	// Protocol
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// Destination port
	Port *string `json:"Port,omitnil,omitempty" name:"Port"`

	// Policy. 1: block; 2: allow
	Strategy *uint64 `json:"Strategy,omitnil,omitempty" name:"Strategy"`

	// Description
	Detail *string `json:"Detail,omitnil,omitempty" name:"Detail"`

	// One-way/two-way. 0: one-way; 1: two-way
	BothWay *uint64 `json:"BothWay,omitnil,omitempty" name:"BothWay"`

	// Rule ID
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// Toggle status. 0: off; 1: on
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Indicates whether the rule is normal. 0: normal; 1: abnormal
	IsNew *uint64 `json:"IsNew,omitnil,omitempty" name:"IsNew"`

	// VPC ID
	// Note: This field may return `null`, indicating that no valid value was found.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Subnet ID
	// Note: This field may return `null`, indicating that no valid value was found.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Instance name
	// Note: This field may return `null`, indicating that no valid value was found.
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Public IP. Multiple IPs are separated by commas.
	// Note: This field may return `null`, indicating that no valid value was found.
	PublicIp *string `json:"PublicIp,omitnil,omitempty" name:"PublicIp"`

	// Private IP. Multiple IPs are separated by commas.
	// Note: This field may return `null`, indicating that no valid value was found.
	PrivateIp *string `json:"PrivateIp,omitnil,omitempty" name:"PrivateIp"`

	// Masked address. Multiple addresses are separated by commas.
	// Note: This field may return `null`, indicating that no valid value was found.
	Cidr *string `json:"Cidr,omitnil,omitempty" name:"Cidr"`

	// Port protocol template ID
	// Note: This field may return `null`, indicating that no valid value was found.
	ServiceTemplateId *string `json:"ServiceTemplateId,omitnil,omitempty" name:"ServiceTemplateId"`

	// Two-way rules
	// Note: This field may return `null`, indicating that no valid value was found.
	BothWayInfo []*SecurityGroupBothWayInfo `json:"BothWayInfo,omitnil,omitempty" name:"BothWayInfo"`

	// Direction. 0: outbound; 1: inbound. 1 by default
	Direction *uint64 `json:"Direction,omitnil,omitempty" name:"Direction"`

	// Indicates whether to use the port protocol template. 0: no; 1: yes
	ProtocolPortType *uint64 `json:"ProtocolPortType,omitnil,omitempty" name:"ProtocolPortType"`
}

type SecurityGroupOrderIndexData struct {
	// Current priority of enterprise security group rules
	OrderIndex *uint64 `json:"OrderIndex,omitnil,omitempty" name:"OrderIndex"`

	// New priority of enterprise security group rules
	NewOrderIndex *uint64 `json:"NewOrderIndex,omitnil,omitempty" name:"NewOrderIndex"`
}

type SecurityGroupRule struct {
	// Source example:
	// net: IP/CIDR (192.168.0.2)
	// template: parameter template (ipm-dyodhpby)
	// instance: asset instance (ins-123456)
	// resourcegroup: asset group (/all groups/group 1/subgroup 1)
	// tag: resource tag ({"Key":"tag key","Value":"tag value"})
	// region: region (ap-gaungzhou)
	SourceContent *string `json:"SourceContent,omitnil,omitempty" name:"SourceContent"`

	// Access source type. Valid values: net|template|instance|resourcegroup|tag|region
	SourceType *string `json:"SourceType,omitnil,omitempty" name:"SourceType"`

	// Destination example:
	// net: IP/CIDR (192.168.0.2)
	// template: parameter template (ipm-dyodhpby)
	// instance: asset instance (ins-123456)
	// resourcegroup: asset group (/all groups/group 1/subgroup 1)
	// tag: resource tag ({"Key":"tag key","Value":"tag value"})
	// region: region (ap-gaungzhou)
	DestContent *string `json:"DestContent,omitnil,omitempty" name:"DestContent"`

	// Access destination type. Valid values: net|template|instance|resourcegroup|tag|region
	DestType *string `json:"DestType,omitnil,omitempty" name:"DestType"`

	// The action that Cloud Firewall performs on the traffic. Valid values:
	// accept: allow
	// drop: deny
	RuleAction *string `json:"RuleAction,omitnil,omitempty" name:"RuleAction"`

	// Description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Rule priority. -1: lowest; 1: highest
	OrderIndex *string `json:"OrderIndex,omitnil,omitempty" name:"OrderIndex"`

	// Protocol. TCP/UDP/ICMP/ANY
	// Note: This field may return `null`, indicating that no valid value was found.
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// The port to apply access control rules. Valid values:
	// -1/-1: all ports
	// 80: port 80
	// Note: This field may return `null`, indicating that no valid value was found.
	Port *string `json:"Port,omitnil,omitempty" name:"Port"`

	// Parameter template ID of port and protocol type; mutually exclusive with Protocol and Port
	// Note: This field may return `null`, indicating that no valid value was found.
	ServiceTemplateId *string `json:"ServiceTemplateId,omitnil,omitempty" name:"ServiceTemplateId"`

	// The unique ID of the rule
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// Rule status. true: enabled; false: disabled
	Enable *string `json:"Enable,omitnil,omitempty" name:"Enable"`
}

type SequenceData struct {
	// Rule ID
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// Rule priority before change
	OrderIndex *uint64 `json:"OrderIndex,omitnil,omitempty" name:"OrderIndex"`

	// Rule priority after change
	NewOrderIndex *uint64 `json:"NewOrderIndex,omitnil,omitempty" name:"NewOrderIndex"`
}

// Predefined struct for user
type SetNatFwDnatRuleRequestParams struct {
	// 0: Create new; 1: Use existing
	Mode *uint64 `json:"Mode,omitnil,omitempty" name:"Mode"`

	// Operation type. Valid values: add, del, and modify.
	OperationType *string `json:"OperationType,omitnil,omitempty" name:"OperationType"`

	// Firewall instance ID. This field is required.
	CfwInstance *string `json:"CfwInstance,omitnil,omitempty" name:"CfwInstance"`

	// List of added/deleted DNAT rules
	AddOrDelDnatRules []*CfwNatDnatRule `json:"AddOrDelDnatRules,omitnil,omitempty" name:"AddOrDelDnatRules"`

	// Original DNAT rule before change
	OriginDnat *CfwNatDnatRule `json:"OriginDnat,omitnil,omitempty" name:"OriginDnat"`

	// New DNAT rule after change
	NewDnat *CfwNatDnatRule `json:"NewDnat,omitnil,omitempty" name:"NewDnat"`
}

type SetNatFwDnatRuleRequest struct {
	*tchttp.BaseRequest
	
	// 0: Create new; 1: Use existing
	Mode *uint64 `json:"Mode,omitnil,omitempty" name:"Mode"`

	// Operation type. Valid values: add, del, and modify.
	OperationType *string `json:"OperationType,omitnil,omitempty" name:"OperationType"`

	// Firewall instance ID. This field is required.
	CfwInstance *string `json:"CfwInstance,omitnil,omitempty" name:"CfwInstance"`

	// List of added/deleted DNAT rules
	AddOrDelDnatRules []*CfwNatDnatRule `json:"AddOrDelDnatRules,omitnil,omitempty" name:"AddOrDelDnatRules"`

	// Original DNAT rule before change
	OriginDnat *CfwNatDnatRule `json:"OriginDnat,omitnil,omitempty" name:"OriginDnat"`

	// New DNAT rule after change
	NewDnat *CfwNatDnatRule `json:"NewDnat,omitnil,omitempty" name:"NewDnat"`
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
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	OperationType *string `json:"OperationType,omitnil,omitempty" name:"OperationType"`

	// Firewall instance ID
	CfwInstance *string `json:"CfwInstance,omitnil,omitempty" name:"CfwInstance"`

	// This field is required when OperationType is "bind" or "unbind".
	EipList []*string `json:"EipList,omitnil,omitempty" name:"EipList"`
}

type SetNatFwEipRequest struct {
	*tchttp.BaseRequest
	
	// bind: bind EIP; unbind: unbind EIP; newAdd: add firewall EIP
	OperationType *string `json:"OperationType,omitnil,omitempty" name:"OperationType"`

	// Firewall instance ID
	CfwInstance *string `json:"CfwInstance,omitnil,omitempty" name:"CfwInstance"`

	// This field is required when OperationType is "bind" or "unbind".
	EipList []*string `json:"EipList,omitnil,omitempty" name:"EipList"`
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
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Num *int64 `json:"Num,omitnil,omitempty" name:"Num"`

	// Port
	Port *string `json:"Port,omitnil,omitempty" name:"Port"`

	// IP
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// Address
	Address *string `json:"Address,omitnil,omitempty" name:"Address"`

	// Asset ID
	InsID *string `json:"InsID,omitnil,omitempty" name:"InsID"`

	// Asset name
	InsName *string `json:"InsName,omitnil,omitempty" name:"InsName"`
}

// Predefined struct for user
type StopSecurityGroupRuleDispatchRequestParams struct {
	// Stops all if set to 1
	StopType *int64 `json:"StopType,omitnil,omitempty" name:"StopType"`
}

type StopSecurityGroupRuleDispatchRequest struct {
	*tchttp.BaseRequest
	
	// Stops all if set to 1
	StopType *int64 `json:"StopType,omitnil,omitempty" name:"StopType"`
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
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	PublicIp *string `json:"PublicIp,omitnil,omitempty" name:"PublicIp"`

	// Private IP
	// Note: This field may return `null`, indicating that no valid value was found.
	IntranetIp *string `json:"IntranetIp,omitnil,omitempty" name:"IntranetIp"`

	// Instance name
	// Note: This field may return `null`, indicating that no valid value was found.
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Instance ID
	// Note: This field may return `null`, indicating that no valid value was found.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Asset type
	AssetType *string `json:"AssetType,omitnil,omitempty" name:"AssetType"`

	// Region
	// Note: This field may return `null`, indicating that no valid value was found.
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// Firewall toggle
	Switch *int64 `json:"Switch,omitnil,omitempty" name:"Switch"`

	// ID value
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// Public IP type
	// Note: This field may return `null`, indicating that no valid value was found.
	PublicIpType *uint64 `json:"PublicIpType,omitnil,omitempty" name:"PublicIpType"`

	// Number of risky ports
	// Note: This field may return `null`, indicating that no valid value was found.
	PortTimes *uint64 `json:"PortTimes,omitnil,omitempty" name:"PortTimes"`

	// Last scan time
	// Note: This field may return `null`, indicating that no valid value was found.
	LastTime *string `json:"LastTime,omitnil,omitempty" name:"LastTime"`

	// Scan mode
	// Note: This field may return `null`, indicating that no valid value was found.
	ScanMode *string `json:"ScanMode,omitnil,omitempty" name:"ScanMode"`

	// Scan status
	// Note: This field may return `null`, indicating that no valid value was found.
	ScanStatus *uint64 `json:"ScanStatus,omitnil,omitempty" name:"ScanStatus"`
}

type TLogInfo struct {
	// Compromised servers
	OutNum *int64 `json:"OutNum,omitnil,omitempty" name:"OutNum"`

	// Unhandled alerts
	HandleNum *int64 `json:"HandleNum,omitnil,omitempty" name:"HandleNum"`

	// Vulnerability attacks
	VulNum *int64 `json:"VulNum,omitnil,omitempty" name:"VulNum"`

	// Detected networks
	NetworkNum *int64 `json:"NetworkNum,omitnil,omitempty" name:"NetworkNum"`

	// Blocklist
	BanNum *int64 `json:"BanNum,omitnil,omitempty" name:"BanNum"`

	// Brute force attacks
	BruteForceNum *int64 `json:"BruteForceNum,omitnil,omitempty" name:"BruteForceNum"`
}

type UnHandleEvent struct {
	// Unhandled event type
	EventTableListStruct []*UnHandleEventDetail `json:"EventTableListStruct,omitnil,omitempty" name:"EventTableListStruct"`

	// 1: yes; 0: no
	BaseLineUser *uint64 `json:"BaseLineUser,omitnil,omitempty" name:"BaseLineUser"`

	// 1: on; 0: off
	BaseLineInSwitch *uint64 `json:"BaseLineInSwitch,omitnil,omitempty" name:"BaseLineInSwitch"`

	// 1: on; 0: off
	BaseLineOutSwitch *uint64 `json:"BaseLineOutSwitch,omitnil,omitempty" name:"BaseLineOutSwitch"`

	// Number of inter-VPC firewall instances
	// Note: This field may return `null`, indicating that no valid value was found.
	VpcFwCount *uint64 `json:"VpcFwCount,omitnil,omitempty" name:"VpcFwCount"`
}

type UnHandleEventDetail struct {
	// Security event name
	EventName *string `json:"EventName,omitnil,omitempty" name:"EventName"`

	// Number of unhandled events
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`
}

type VpcDnsInfo struct {
	// VPC ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// VPC name
	VpcName *string `json:"VpcName,omitnil,omitempty" name:"VpcName"`

	// NAT firewall mode. 0: Create new; 1: Use existing
	FwMode *int64 `json:"FwMode,omitnil,omitempty" name:"FwMode"`

	// VPC IPv4 CIDR block (Classless Inter-Domain Routing)
	VpcIpv4Cidr *string `json:"VpcIpv4Cidr,omitnil,omitempty" name:"VpcIpv4Cidr"`

	// Public EIP, which is the firewall DNS resolution address
	DNSEip *string `json:"DNSEip,omitnil,omitempty" name:"DNSEip"`

	// NAT gateway ID
	// Note: This field may return `null`, indicating that no valid value was found.
	NatInsId *string `json:"NatInsId,omitnil,omitempty" name:"NatInsId"`

	// NAT gateway name
	// Note: This field may return `null`, indicating that no valid value was found.
	NatInsName *string `json:"NatInsName,omitnil,omitempty" name:"NatInsName"`

	// 0: off; 1: on
	SwitchStatus *int64 `json:"SwitchStatus,omitnil,omitempty" name:"SwitchStatus"`
}

type VpcRuleItem struct {
	// Access source example:
	// 
	// net: IP/CIDR (192.168.0.2)
	SourceContent *string `json:"SourceContent,omitnil,omitempty" name:"SourceContent"`

	// Access source type, which can be: net
	SourceType *string `json:"SourceType,omitnil,omitempty" name:"SourceType"`

	// Access destination example:
	// 
	// net: IP/CIDR (192.168.0.2)
	// 
	// domain: domain name rules, for example *.qq.com
	DestContent *string `json:"DestContent,omitnil,omitempty" name:"DestContent"`

	// Access destination type, which can be: net, domain, dnsparse
	DestType *string `json:"DestType,omitnil,omitempty" name:"DestType"`

	// Protocol, optional values:
	// 
	// TCP
	// 
	// UDP
	// 
	// ICMP
	// 
	// ANY
	// 
	// HTTP
	// 
	// HTTPS
	// 
	// HTTP/HTTPS
	// 
	// SMTP
	// 
	// SMTPS
	// 
	// SMTP/SMTPS
	// 
	// FTP
	// 
	// DNS
	// 
	// TLS/SSL
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// How traffic set in the access control policy passes through CFW. Values:
	// 
	// accept: allow
	// 
	// drop: deny
	// 
	// log: observe
	RuleAction *string `json:"RuleAction,omitnil,omitempty" name:"RuleAction"`

	// Access control policy ports. Values:
	// 
	// -1/-1: all ports
	// 
	// 80: port 80
	Port *string `json:"Port,omitnil,omitempty" name:"Port"`

	// Description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Rule order, where -1 indicates the lowest and 1 indicates the highest.
	OrderIndex *int64 `json:"OrderIndex,omitnil,omitempty" name:"OrderIndex"`

	// Rule status, where true indicates enabled and false indicates disabled.
	Enable *string `json:"Enable,omitnil,omitempty" name:"Enable"`

	// The scope of effect for the rule, specifying whether it applies between a specific pair of VPCs or across all VPCs.
	EdgeId *string `json:"EdgeId,omitnil,omitempty" name:"EdgeId"`

	// The unique id corresponding to the rule. This field is ignored when adding a rule; when modifying the rule, the Uuid needs to be filled in. This parameter will be returned in query results.
	Uuid *int64 `json:"Uuid,omitnil,omitempty" name:"Uuid"`

	// The hit count of the rule. This parameter does not need to be passed during CRUD operations and is mainly used for returning query result data.
	DetectedTimes *int64 `json:"DetectedTimes,omitnil,omitempty" name:"DetectedTimes"`

	// Description of the firewall between the pair of VPCs corresponding to EdgeId
	EdgeName *string `json:"EdgeName,omitnil,omitempty" name:"EdgeName"`

	// Internal-use uuid, generally not used
	InternalUuid *int64 `json:"InternalUuid,omitnil,omitempty" name:"InternalUuid"`

	// Rule deletion status: 1 indicates deleted; 0 indicates not deleted
	Deleted *int64 `json:"Deleted,omitnil,omitempty" name:"Deleted"`

	// The firewall instance ID where the rule takes effect
	FwGroupId *string `json:"FwGroupId,omitnil,omitempty" name:"FwGroupId"`

	// Firewall name
	FwGroupName *string `json:"FwGroupName,omitnil,omitempty" name:"FwGroupName"`

	// beta task details
	BetaList []*BetaInfoByACL `json:"BetaList,omitnil,omitempty" name:"BetaList"`

	// Port Protocol Group ID
	ParamTemplateId *string `json:"ParamTemplateId,omitnil,omitempty" name:"ParamTemplateId"`

	// Port Protocol Group Name
	ParamTemplateName *string `json:"ParamTemplateName,omitnil,omitempty" name:"ParamTemplateName"`

	// Access destination name
	TargetName *string `json:"TargetName,omitnil,omitempty" name:"TargetName"`

	// Access source name
	SourceName *string `json:"SourceName,omitnil,omitempty" name:"SourceName"`

	// Ip version, 0: IPv4, 1: IPv6, default is IPv4
	IpVersion *int64 `json:"IpVersion,omitnil,omitempty" name:"IpVersion"`

	// Whether the rule is invalid, where 0 indicates a valid rule and 1 indicates an invalid rule, used in output parameters.
	Invalid *int64 `json:"Invalid,omitnil,omitempty" name:"Invalid"`
}