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

package v20220802

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

type AggregateResourceInfo struct {
	// Resource type
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// Resource name
	ResourceName *string `json:"ResourceName,omitnil,omitempty" name:"ResourceName"`

	// Resource ID
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// Region
	// 
	// Note: This field may return null, indicating that no valid value is found.
	ResourceRegion *string `json:"ResourceRegion,omitnil,omitempty" name:"ResourceRegion"`

	// Resource Status
	// 
	// Note: This field may return null, indicating that no valid value is found.
	ResourceStatus *string `json:"ResourceStatus,omitnil,omitempty" name:"ResourceStatus"`

	// Whether to delete. 1: Deleted; 0: Not deleted.
	// Note: This field may return null, indicating that no valid value is found.
	ResourceDelete *uint64 `json:"ResourceDelete,omitnil,omitempty" name:"ResourceDelete"`

	// Resource creation time
	// 
	// Note: This field may return null, indicating that no valid value is found.
	ResourceCreateTime *string `json:"ResourceCreateTime,omitnil,omitempty" name:"ResourceCreateTime"`

	// Tag information
	// 
	// Note: This field may return null, indicating that no valid value is found.
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Availability zone
	// 
	// Note: This field may return null, indicating that no valid value is found.
	ResourceZone *string `json:"ResourceZone,omitnil,omitempty" name:"ResourceZone"`

	// Compliance status
	// Note: This field may return null, indicating that no valid value is found.
	ComplianceResult *string `json:"ComplianceResult,omitnil,omitempty" name:"ComplianceResult"`

	// Resource owner uid
	ResourceOwnerId *uint64 `json:"ResourceOwnerId,omitnil,omitempty" name:"ResourceOwnerId"`

	// User nickname
	// Note: This field may return null, indicating that no valid value is found.
	ResourceOwnerName *string `json:"ResourceOwnerName,omitnil,omitempty" name:"ResourceOwnerName"`
}

type Annotation struct {
	// Current actual configuration of the resource. It can contain 0 to 256 characters, which is the non-compliant configuration of the resource.
	// Note: This field may return null, indicating that no valid value is found.
	Configuration *string `json:"Configuration,omitnil,omitempty" name:"Configuration"`

	// Desired configuration of the resource. It can contain 0 to 256 characters, which is the compliant configuration of the resource.
	// Note: This field may return null, indicating that no valid value is found.
	DesiredValue *string `json:"DesiredValue,omitnil,omitempty" name:"DesiredValue"`

	// Comparison operator between current and desired configuration of the resource. Length is 0-16 characters. This field may be empty when custom rule reporting evaluation result.
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`

	// JSON path of current configuration in resource attribute structure. Length is 0-256 characters. This field may be empty when custom rule reporting evaluation result.
	Property *string `json:"Property,omitnil,omitempty" name:"Property"`
}

type ConfigRule struct {
	// Rule identifier
	// Note: This field may return null, indicating that no valid value is found.
	Identifier *string `json:"Identifier,omitnil,omitempty" name:"Identifier"`

	// Name of the rule
	// 
	// Note: This field may return null, indicating that no valid value is found.
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// Rule parameters
	// Note: This field may return null, indicating that no valid value is found.
	InputParameter []*InputParameter `json:"InputParameter,omitnil,omitempty" name:"InputParameter"`

	// Rule trigger condition.
	// 
	// Note: This field may return null, indicating that no valid value is found.
	SourceCondition []*SourceConditionForManage `json:"SourceCondition,omitnil,omitempty" name:"SourceCondition"`

	// Resource types supported by rule. The rule only applies to specified resource types.
	// Note: This field may return null, indicating that no valid value is found.
	ResourceType []*string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// Rule ownership tag
	// Note: This field may return null, indicating that no valid value is found.
	Labels []*string `json:"Labels,omitnil,omitempty" name:"Labels"`

	// Rule risk level
	// 1: Low risk
	// 2: Medium risk
	// 3: High risk
	// Note: This field may return null, indicating that no valid value is found.
	RiskLevel *int64 `json:"RiskLevel,omitnil,omitempty" name:"RiskLevel"`

	// Function corresponding to rule
	// Note: This field may return null, indicating that no valid value is found.
	ServiceFunction *string `json:"ServiceFunction,omitnil,omitempty" name:"ServiceFunction"`

	// Creation time
	// 
	// Format: YYYY-MM-DD h:i:s
	// Note: This field may return null, indicating that no valid value is found.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Rule description
	// 
	// Note: This field may return null, indicating that no valid value is found.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// ACTIVE: Enabled
	// NO_ACTIVE: Disabled
	// Note: This field may return null, indicating that no valid value is found.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Compliance: 'COMPLIANT'
	// 'NON_COMPLIANT'
	// 'NOT_APPLICABLE'
	// Note: This field may return null, indicating that no valid value is found.
	ComplianceResult *string `json:"ComplianceResult,omitnil,omitempty" name:"ComplianceResult"`

	// ["",""]
	// Note: This field may return null, indicating that no valid value is found.
	Annotation *Annotation `json:"Annotation,omitnil,omitempty" name:"Annotation"`

	// Rule evaluation time
	// Format: YYYY-MM-DD h:i:s
	// 
	// Note: This field may return null, indicating that no valid value is found.
	ConfigRuleInvokedTime *string `json:"ConfigRuleInvokedTime,omitnil,omitempty" name:"ConfigRuleInvokedTime"`

	// Rule ID
	// 
	// Note: This field may return null, indicating that no valid value is found.
	ConfigRuleId *string `json:"ConfigRuleId,omitnil,omitempty" name:"ConfigRuleId"`

	// CUSTOMIZE
	// Managed rule
	// Note: This field may return null, indicating that no valid value is found.
	IdentifierType *string `json:"IdentifierType,omitnil,omitempty" name:"IdentifierType"`

	// Compliance package ID
	// Note: This field may return null, indicating that no valid value is found.
	CompliancePackId *string `json:"CompliancePackId,omitnil,omitempty" name:"CompliancePackId"`

	// Trigger Type
	// 
	// Scheduled trigger
	// Triggered by configuration change
	// Note: This field may return null, indicating that no valid value is found.
	TriggerType []*TriggerType `json:"TriggerType,omitnil,omitempty" name:"TriggerType"`

	// Parameter details
	// 
	// Note: This field may return null, indicating that no valid value is found.
	ManageInputParameter []*InputParameterForManage `json:"ManageInputParameter,omitnil,omitempty" name:"ManageInputParameter"`

	// Rule name
	// 
	// Note: This field may return null, indicating that no valid value is found.
	CompliancePackName *string `json:"CompliancePackName,omitnil,omitempty" name:"CompliancePackName"`

	// Associated region
	// Note: This field may return null, indicating that no valid value is found.
	RegionsScope []*string `json:"RegionsScope,omitnil,omitempty" name:"RegionsScope"`

	// Associate Tag
	// 
	// Note: This field may return null, indicating that no valid value is found.
	TagsScope []*Tag `json:"TagsScope,omitnil,omitempty" name:"TagsScope"`

	//  The rule is invalid for the specified resource ID, meaning it does not evaluate the resource.
	// Note: This field may return null, indicating that no valid value is found.
	ExcludeResourceIdsScope []*string `json:"ExcludeResourceIdsScope,omitnil,omitempty" name:"ExcludeResourceIdsScope"`

	// Account group ID
	// 
	// Note: This field may return null, indicating that no valid value is found.
	AccountGroupId *string `json:"AccountGroupId,omitnil,omitempty" name:"AccountGroupId"`

	// Account group name
	// Note: This field may return null, indicating that no valid value is found.
	AccountGroupName *string `json:"AccountGroupName,omitnil,omitempty" name:"AccountGroupName"`

	// Rule owner user ID
	// Note: This field may return null, indicating that no valid value is found.
	RuleOwnerId *uint64 `json:"RuleOwnerId,omitnil,omitempty" name:"RuleOwnerId"`

	// Trigger methods supported by preset rules
	// Scheduled trigger
	// Triggered by configuration change
	ManageTriggerType []*string `json:"ManageTriggerType,omitnil,omitempty" name:"ManageTriggerType"`
}

// Predefined struct for user
type DescribeDiscoveredResourceRequestParams struct {
	// Resource ID
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// Resource type
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// Resource region
	ResourceRegion *string `json:"ResourceRegion,omitnil,omitempty" name:"ResourceRegion"`
}

type DescribeDiscoveredResourceRequest struct {
	*tchttp.BaseRequest
	
	// Resource ID
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// Resource type
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// Resource region
	ResourceRegion *string `json:"ResourceRegion,omitnil,omitempty" name:"ResourceRegion"`
}

func (r *DescribeDiscoveredResourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDiscoveredResourceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ResourceId")
	delete(f, "ResourceType")
	delete(f, "ResourceRegion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDiscoveredResourceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDiscoveredResourceResponseParams struct {
	// Resource ID
	// 
	// Note: This field may return null, indicating that no valid value is found.
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// Resource type
	// 
	// Note: This field may return null, indicating that no valid value is found.
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// Resource Name
	// 
	// Note: This field may return null, indicating that no valid value is found.
	ResourceName *string `json:"ResourceName,omitnil,omitempty" name:"ResourceName"`

	// Resource region
	// 
	// Note: This field may return null, indicating that no valid value is found.
	ResourceRegion *string `json:"ResourceRegion,omitnil,omitempty" name:"ResourceRegion"`

	// Resource availability zone
	// Note: This field may return null, indicating that no valid value is found.
	ResourceZone *string `json:"ResourceZone,omitnil,omitempty" name:"ResourceZone"`

	// Resource configuration
	// 
	// Note: This field may return null, indicating that no valid value is found.
	Configuration *string `json:"Configuration,omitnil,omitempty" name:"Configuration"`

	// Resource creation time
	// 
	// Note: This field may return null, indicating that no valid value is found.
	ResourceCreateTime *string `json:"ResourceCreateTime,omitnil,omitempty" name:"ResourceCreateTime"`

	// Resource tag
	// 
	// Note: This field may return null, indicating that no valid value is found.
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Resource update time
	// Note: This field may return null, indicating that no valid value is found.
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDiscoveredResourceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDiscoveredResourceResponseParams `json:"Response"`
}

func (r *DescribeDiscoveredResourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDiscoveredResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Evaluation struct {
	// Evaluated resource id. It can contain 0 to 256 characters.
	ComplianceResourceId *string `json:"ComplianceResourceId,omitnil,omitempty" name:"ComplianceResourceId"`

	// Evaluated resource type.
	// Supported:
	// QCS::CVM::Instance、 QCS::CBS::Disk、QCS::VPC::Vpc、QCS::VPC::Subnet、QCS::VPC::SecurityGroup、 QCS::CAM::User、QCS::CAM::Group、QCS::CAM::Policy、QCS::CAM::Role、QCS::COS::Bucket
	ComplianceResourceType *string `json:"ComplianceResourceType,omitnil,omitempty" name:"ComplianceResourceType"`

	// Evaluated resource region.
	// It can contain 0 to 32 characters.
	ComplianceRegion *string `json:"ComplianceRegion,omitnil,omitempty" name:"ComplianceRegion"`

	// Compliance type. Valid values:
	// COMPLIANT: Compliant,
	// NON_COMPLIANT: Non-compliant
	ComplianceType *string `json:"ComplianceType,omitnil,omitempty" name:"ComplianceType"`

	// Supplementary information for non-compliant resources.
	Annotation *Annotation `json:"Annotation,omitnil,omitempty" name:"Annotation"`
}

type Filter struct {
	// Query field name Resource name: resourceName Resource ID: resourceId Resource type: resourceType Resource region: resourceRegion Deletion status: resourceDelete 0 not deleted, 1 deleted resourceregionandzone region/az
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Value of the field to query
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type InputParameter struct {
	// Parameter name
	ParameterKey *string `json:"ParameterKey,omitnil,omitempty" name:"ParameterKey"`

	// Parameter type. Required type: Require, optional type: Optional.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Parameter value
	// 
	// Note: This field may return null, indicating that no valid value is found.
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type InputParameterForManage struct {
	// Value type. Integer: Integer, String: String.
	// Note: This field may return null, indicating that no valid value is found.
	ValueType *string `json:"ValueType,omitnil,omitempty" name:"ValueType"`

	// Parameter key
	// Note: This field may return null, indicating that no valid value is found.
	ParameterKey *string `json:"ParameterKey,omitnil,omitempty" name:"ParameterKey"`

	// Parameter type. Required type: Required, Optional type: Optional.
	// Note: This field may return null, indicating that no valid value is found.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Default value
	// 
	// Note: This field may return null, indicating that no valid value is found.
	DefaultValue *string `json:"DefaultValue,omitnil,omitempty" name:"DefaultValue"`

	// Description
	// 
	// Note: This field may return null, indicating that no valid value is found.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

// Predefined struct for user
type ListAggregateConfigRulesRequestParams struct {
	// Specifies the limit per page.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Account group ID
	AccountGroupId *string `json:"AccountGroupId,omitnil,omitempty" name:"AccountGroupId"`

	// Sort type, descending: desc, ascending: asc.
	OrderType *string `json:"OrderType,omitnil,omitempty" name:"OrderType"`

	// Risk level
	// 
	// 1: High risk.
	// 2: Medium risk.
	// 3: Low risk.
	RiskLevel []*uint64 `json:"RiskLevel,omitnil,omitempty" name:"RiskLevel"`

	// Rule status
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// Evaluation result
	ComplianceResult []*string `json:"ComplianceResult,omitnil,omitempty" name:"ComplianceResult"`

	// Name of the rule
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// Rule ownership account ID
	RuleOwnerId *uint64 `json:"RuleOwnerId,omitnil,omitempty" name:"RuleOwnerId"`
}

type ListAggregateConfigRulesRequest struct {
	*tchttp.BaseRequest
	
	// Specifies the limit per page.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Account group ID
	AccountGroupId *string `json:"AccountGroupId,omitnil,omitempty" name:"AccountGroupId"`

	// Sort type, descending: desc, ascending: asc.
	OrderType *string `json:"OrderType,omitnil,omitempty" name:"OrderType"`

	// Risk level
	// 
	// 1: High risk.
	// 2: Medium risk.
	// 3: Low risk.
	RiskLevel []*uint64 `json:"RiskLevel,omitnil,omitempty" name:"RiskLevel"`

	// Rule status
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// Evaluation result
	ComplianceResult []*string `json:"ComplianceResult,omitnil,omitempty" name:"ComplianceResult"`

	// Name of the rule
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// Rule ownership account ID
	RuleOwnerId *uint64 `json:"RuleOwnerId,omitnil,omitempty" name:"RuleOwnerId"`
}

func (r *ListAggregateConfigRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAggregateConfigRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "AccountGroupId")
	delete(f, "OrderType")
	delete(f, "RiskLevel")
	delete(f, "State")
	delete(f, "ComplianceResult")
	delete(f, "RuleName")
	delete(f, "RuleOwnerId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListAggregateConfigRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAggregateConfigRulesResponseParams struct {
	// Total number
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// Details
	Items []*ConfigRule `json:"Items,omitnil,omitempty" name:"Items"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListAggregateConfigRulesResponse struct {
	*tchttp.BaseResponse
	Response *ListAggregateConfigRulesResponseParams `json:"Response"`
}

func (r *ListAggregateConfigRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAggregateConfigRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAggregateDiscoveredResourcesRequestParams struct {
	// Items per Page
	MaxResults *uint64 `json:"MaxResults,omitnil,omitempty" name:"MaxResults"`

	// Account group ID
	AccountGroupId *string `json:"AccountGroupId,omitnil,omitempty" name:"AccountGroupId"`

	// resourceName: Resource name; resourceId: Resource ID; resourceType: Resource type
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <Tag>
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Next page token.
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`

	// Sorting method asc, desc
	OrderType *string `json:"OrderType,omitnil,omitempty" name:"OrderType"`
}

type ListAggregateDiscoveredResourcesRequest struct {
	*tchttp.BaseRequest
	
	// Items per Page
	MaxResults *uint64 `json:"MaxResults,omitnil,omitempty" name:"MaxResults"`

	// Account group ID
	AccountGroupId *string `json:"AccountGroupId,omitnil,omitempty" name:"AccountGroupId"`

	// resourceName: Resource name; resourceId: Resource ID; resourceType: Resource type
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <Tag>
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Next page token.
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`

	// Sorting method asc, desc
	OrderType *string `json:"OrderType,omitnil,omitempty" name:"OrderType"`
}

func (r *ListAggregateDiscoveredResourcesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAggregateDiscoveredResourcesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MaxResults")
	delete(f, "AccountGroupId")
	delete(f, "Filters")
	delete(f, "Tags")
	delete(f, "NextToken")
	delete(f, "OrderType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListAggregateDiscoveredResourcesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAggregateDiscoveredResourcesResponseParams struct {
	// Details.
	Items []*AggregateResourceInfo `json:"Items,omitnil,omitempty" name:"Items"`

	// next page
	// Note: This field may return null, indicating that no valid value is found.
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListAggregateDiscoveredResourcesResponse struct {
	*tchttp.BaseResponse
	Response *ListAggregateDiscoveredResourcesResponseParams `json:"Response"`
}

func (r *ListAggregateDiscoveredResourcesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAggregateDiscoveredResourcesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListConfigRulesRequestParams struct {
	// Page limit
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Sort type. Descending: desc, Ascending: asc.
	OrderType *string `json:"OrderType,omitnil,omitempty" name:"OrderType"`

	// Risk level
	// 
	// 1: High risk.
	// 2: Medium risk.
	// 3: Low risk.
	RiskLevel []*uint64 `json:"RiskLevel,omitnil,omitempty" name:"RiskLevel"`

	// Rule status
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// Evaluation result
	ComplianceResult []*string `json:"ComplianceResult,omitnil,omitempty" name:"ComplianceResult"`

	// Name of the rule
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`
}

type ListConfigRulesRequest struct {
	*tchttp.BaseRequest
	
	// Page limit
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Sort type. Descending: desc, Ascending: asc.
	OrderType *string `json:"OrderType,omitnil,omitempty" name:"OrderType"`

	// Risk level
	// 
	// 1: High risk.
	// 2: Medium risk.
	// 3: Low risk.
	RiskLevel []*uint64 `json:"RiskLevel,omitnil,omitempty" name:"RiskLevel"`

	// Rule status
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// Evaluation result
	ComplianceResult []*string `json:"ComplianceResult,omitnil,omitempty" name:"ComplianceResult"`

	// Name of the rule
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`
}

func (r *ListConfigRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListConfigRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "OrderType")
	delete(f, "RiskLevel")
	delete(f, "State")
	delete(f, "ComplianceResult")
	delete(f, "RuleName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListConfigRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListConfigRulesResponseParams struct {
	// Total number
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// Details
	Items []*ConfigRule `json:"Items,omitnil,omitempty" name:"Items"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListConfigRulesResponse struct {
	*tchttp.BaseResponse
	Response *ListConfigRulesResponseParams `json:"Response"`
}

func (r *ListConfigRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListConfigRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListDiscoveredResourcesRequestParams struct {
	// Items per Page
	MaxResults *uint64 `json:"MaxResults,omitnil,omitempty" name:"MaxResults"`

	// resourceName: Resource name resourceId: Resource ID
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Tag
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Next page token.
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`

	// Sorting method asc, desc
	OrderType *string `json:"OrderType,omitnil,omitempty" name:"OrderType"`
}

type ListDiscoveredResourcesRequest struct {
	*tchttp.BaseRequest
	
	// Items per Page
	MaxResults *uint64 `json:"MaxResults,omitnil,omitempty" name:"MaxResults"`

	// resourceName: Resource name resourceId: Resource ID
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Tag
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Next page token.
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`

	// Sorting method asc, desc
	OrderType *string `json:"OrderType,omitnil,omitempty" name:"OrderType"`
}

func (r *ListDiscoveredResourcesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListDiscoveredResourcesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MaxResults")
	delete(f, "Filters")
	delete(f, "Tags")
	delete(f, "NextToken")
	delete(f, "OrderType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListDiscoveredResourcesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListDiscoveredResourcesResponseParams struct {
	// Details
	Items []*ResourceListInfo `json:"Items,omitnil,omitempty" name:"Items"`

	// Next page
	// Note: This field may return null, indicating that no valid value is found.
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListDiscoveredResourcesResponse struct {
	*tchttp.BaseResponse
	Response *ListDiscoveredResourcesResponseParams `json:"Response"`
}

func (r *ListDiscoveredResourcesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListDiscoveredResourcesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PutEvaluationsRequestParams struct {
	// Callback token. Obtained from the ResultToken value in the Context of the selected Serverless Cloud Function (SCF) for the custom rule.
	ResultToken *string `json:"ResultToken,omitnil,omitempty" name:"ResultToken"`

	// Custom rule evaluation result information.
	Evaluations []*Evaluation `json:"Evaluations,omitnil,omitempty" name:"Evaluations"`
}

type PutEvaluationsRequest struct {
	*tchttp.BaseRequest
	
	// Callback token. Obtained from the ResultToken value in the Context of the selected Serverless Cloud Function (SCF) for the custom rule.
	ResultToken *string `json:"ResultToken,omitnil,omitempty" name:"ResultToken"`

	// Custom rule evaluation result information.
	Evaluations []*Evaluation `json:"Evaluations,omitnil,omitempty" name:"Evaluations"`
}

func (r *PutEvaluationsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PutEvaluationsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ResultToken")
	delete(f, "Evaluations")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PutEvaluationsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PutEvaluationsResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type PutEvaluationsResponse struct {
	*tchttp.BaseResponse
	Response *PutEvaluationsResponseParams `json:"Response"`
}

func (r *PutEvaluationsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PutEvaluationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResourceListInfo struct {
	// Resource type
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// Resource name
	ResourceName *string `json:"ResourceName,omitnil,omitempty" name:"ResourceName"`

	// Resource ID
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// Region
	// 
	// Note: This field may return null, indicating that no valid value is found.
	ResourceRegion *string `json:"ResourceRegion,omitnil,omitempty" name:"ResourceRegion"`

	// Resource Status
	// 
	// Note: This field may return null, indicating that no valid value is found.
	ResourceStatus *string `json:"ResourceStatus,omitnil,omitempty" name:"ResourceStatus"`

	// 1: Deleted. 2: Not deleted.
	// Note: This field may return null, indicating that no valid value is found.
	ResourceDelete *uint64 `json:"ResourceDelete,omitnil,omitempty" name:"ResourceDelete"`

	// Resource creation time
	// 
	// Note: This field may return null, indicating that no valid value is found.
	ResourceCreateTime *string `json:"ResourceCreateTime,omitnil,omitempty" name:"ResourceCreateTime"`

	// Tag information
	// 
	// Note: This field may return null, indicating that no valid value is found.
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Availability zone
	// 
	// Note: This field may return null, indicating that no valid value is found.
	ResourceZone *string `json:"ResourceZone,omitnil,omitempty" name:"ResourceZone"`

	// Compliance status.
	// Note: This field may return null, indicating that no valid value is found.
	ComplianceResult *string `json:"ComplianceResult,omitnil,omitempty" name:"ComplianceResult"`
}

type SourceConditionForManage struct {
	// Condition is empty, Compliant: COMPLIANT, Non-compliant: NON_COMPLIANT, Not applicable: NOT_APPLICABLE.
	// Note: This field may return null, indicating that no valid value is found.
	EmptyAs *string `json:"EmptyAs,omitnil,omitempty" name:"EmptyAs"`

	// Configuration path
	// 
	// Note: This field may return null, indicating that no valid value is found.
	SelectPath *string `json:"SelectPath,omitnil,omitempty" name:"SelectPath"`

	// Operators
	// Note: This field may return null, indicating that no valid value is found.
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`

	// Required or not.
	// 
	// Note: This field may return null, indicating that no valid value is found.
	Required *bool `json:"Required,omitnil,omitempty" name:"Required"`

	// Expected value
	// Note: This field may return null, indicating that no valid value is found.
	DesiredValue *string `json:"DesiredValue,omitnil,omitempty" name:"DesiredValue"`
}

type Tag struct {
	// Tag key
	// 
	// Note: This field may return null, indicating that no valid value is found.
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// Tag value
	// 
	// Note: This field may return null, indicating that no valid value is found.
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

type TriggerType struct {
	// Trigger Type
	MessageType *string `json:"MessageType,omitnil,omitempty" name:"MessageType"`

	// Trigger time period
	// Note: This field may return null, indicating that no valid value is found.
	MaximumExecutionFrequency *string `json:"MaximumExecutionFrequency,omitnil,omitempty" name:"MaximumExecutionFrequency"`
}