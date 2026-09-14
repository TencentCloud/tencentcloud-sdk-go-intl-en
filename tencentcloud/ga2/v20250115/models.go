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

package v20250115

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

type AcceleratorAreas struct {
	// <p>Acceleration region.</p>
	AccelerateRegion *string `json:"AccelerateRegion,omitnil,omitempty" name:"AccelerateRegion"`

	// <p>Bandwidth.</p>
	Bandwidth *uint64 `json:"Bandwidth,omitnil,omitempty" name:"Bandwidth"`

	// <p>Supports &#39;BGP&#39;, &#39;QUALITY_BGP&#39;, and &#39;STATIC_IP&#39;. Default: BGP.</p><p>Enumeration values:</p><ul><li>BGP: BGP</li><li>STATIC_IP: triple-network</li><li>QUALITY_BGP: dedicated BGP</li></ul>
	IspType *string `json:"IspType,omitnil,omitempty" name:"IspType"`

	// <p>Only IPv4 is supported, and IPv4 is selected by default.</p>
	IpVersion *string `json:"IpVersion,omitnil,omitempty" name:"IpVersion"`

	// <p>Acceleration region ID.</p>
	AcceleratorAreaId *string `json:"AcceleratorAreaId,omitnil,omitempty" name:"AcceleratorAreaId"`

	// <p>IP.</p>
	IpAddress []*string `json:"IpAddress,omitnil,omitempty" name:"IpAddress"`

	// <p>IP information.</p>
	IpAddressInfoSet []*IpAddressInfoSet `json:"IpAddressInfoSet,omitnil,omitempty" name:"IpAddressInfoSet"`
}

type AcceleratorRegionSet struct {
	// <p>Chinese Name of Region.</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Whether available; 0: unavailable, 1: available.</p>
	IsAvailable *int64 `json:"IsAvailable,omitnil,omitempty" name:"IsAvailable"`

	// <p>Regional information.</p>
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// <p>Zone name.</p>
	AreaName *string `json:"AreaName,omitnil,omitempty" name:"AreaName"`

	// <p>Whether it is a China region.</p>
	IsChinaMainland *uint64 `json:"IsChinaMainland,omitnil,omitempty" name:"IsChinaMainland"`

	// <p>Support the IspType type.</p>
	SupportIspType []*string `json:"SupportIspType,omitnil,omitempty" name:"SupportIspType"`

	// <p>Whether it is a Tencent region.</p>
	IsTencentRegion *uint64 `json:"IsTencentRegion,omitnil,omitempty" name:"IsTencentRegion"`
}

type AclEntries struct {
	// <p>Protocol.</p><p>Input limits: supports configuration of 'TCP', 'UDP', 'ALL';</p>
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// <p>Port.</p>
	Port *string `json:"Port,omitnil,omitempty" name:"Port"`

	// <p>IP range.</p>
	SourceCidrBlock *string `json:"SourceCidrBlock,omitnil,omitempty" name:"SourceCidrBlock"`

	// <p>Execute action.</p><p>Input parameter limit: can be configured with 'ACCEPT' and 'DROP';</p>
	Policy *string `json:"Policy,omitnil,omitempty" name:"Policy"`

	// <p>Description. Maximum length cannot exceed 100 bytes.</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

// Predefined struct for user
type CreateAccelerateAreasRequestParams struct {
	// <p>Global acceleration instance ID.</p>
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// <p>Acceleration region info. Up to 10 acceleration region groups can be created at a time.</p>
	AcceleratorAreas []*AcceleratorAreas `json:"AcceleratorAreas,omitnil,omitempty" name:"AcceleratorAreas"`
}

type CreateAccelerateAreasRequest struct {
	*tchttp.BaseRequest
	
	// <p>Global acceleration instance ID.</p>
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// <p>Acceleration region info. Up to 10 acceleration region groups can be created at a time.</p>
	AcceleratorAreas []*AcceleratorAreas `json:"AcceleratorAreas,omitnil,omitempty" name:"AcceleratorAreas"`
}

func (r *CreateAccelerateAreasRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAccelerateAreasRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GlobalAcceleratorId")
	delete(f, "AcceleratorAreas")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAccelerateAreasRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAccelerateAreasResponseParams struct {
	// <p>Asynchronous task ID.</p>
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAccelerateAreasResponse struct {
	*tchttp.BaseResponse
	Response *CreateAccelerateAreasResponseParams `json:"Response"`
}

func (r *CreateAccelerateAreasResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAccelerateAreasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEndpointGroupRequestParams struct {
	// <p>Global acceleration instance ID.</p>
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// <p>Listener ID.</p>
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// <p>Node group type.</p><p>Enumeration values:</p><ul><li>VIRTUAL: custom endpoint node group</li><li>DEFAULT: default terminal node group</li></ul>
	EndpointGroupType *string `json:"EndpointGroupType,omitnil,omitempty" name:"EndpointGroupType"`

	// <p>Terminal node group configuration.</p>
	EndpointGroupConfiguration *EndpointGroupConfiguration `json:"EndpointGroupConfiguration,omitnil,omitempty" name:"EndpointGroupConfiguration"`
}

type CreateEndpointGroupRequest struct {
	*tchttp.BaseRequest
	
	// <p>Global acceleration instance ID.</p>
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// <p>Listener ID.</p>
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// <p>Node group type.</p><p>Enumeration values:</p><ul><li>VIRTUAL: custom endpoint node group</li><li>DEFAULT: default terminal node group</li></ul>
	EndpointGroupType *string `json:"EndpointGroupType,omitnil,omitempty" name:"EndpointGroupType"`

	// <p>Terminal node group configuration.</p>
	EndpointGroupConfiguration *EndpointGroupConfiguration `json:"EndpointGroupConfiguration,omitnil,omitempty" name:"EndpointGroupConfiguration"`
}

func (r *CreateEndpointGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEndpointGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GlobalAcceleratorId")
	delete(f, "ListenerId")
	delete(f, "EndpointGroupType")
	delete(f, "EndpointGroupConfiguration")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateEndpointGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEndpointGroupResponseParams struct {
	// <p>Task ID.</p>
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// <p>Instance ID of the terminal node group.</p>
	EndpointGroupId *string `json:"EndpointGroupId,omitnil,omitempty" name:"EndpointGroupId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateEndpointGroupResponse struct {
	*tchttp.BaseResponse
	Response *CreateEndpointGroupResponseParams `json:"Response"`
}

func (r *CreateEndpointGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEndpointGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateForwardingPolicyRequestParams struct {
	// <p>Global acceleration instance ID.</p>
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// <p>Listener ID.</p>
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// <p>Domain name.</p><p>Parameter format: format, must meet the regular expression: ^(<a href="?:[a-z0-9-]{0,61}[a-z0-9]">a-z0-9</a>?.)+[a-z]{2,}$</p><p>Input limit: length range is 1-80.</p>
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`
}

type CreateForwardingPolicyRequest struct {
	*tchttp.BaseRequest
	
	// <p>Global acceleration instance ID.</p>
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// <p>Listener ID.</p>
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// <p>Domain name.</p><p>Parameter format: format, must meet the regular expression: ^(<a href="?:[a-z0-9-]{0,61}[a-z0-9]">a-z0-9</a>?.)+[a-z]{2,}$</p><p>Input limit: length range is 1-80.</p>
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`
}

func (r *CreateForwardingPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateForwardingPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GlobalAcceleratorId")
	delete(f, "ListenerId")
	delete(f, "Host")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateForwardingPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateForwardingPolicyResponseParams struct {
	// <p>Asynchronous task ID.</p>
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// <p>Layer-7 forwarding policy ID.</p>
	ForwardingPolicyId *string `json:"ForwardingPolicyId,omitnil,omitempty" name:"ForwardingPolicyId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateForwardingPolicyResponse struct {
	*tchttp.BaseResponse
	Response *CreateForwardingPolicyResponseParams `json:"Response"`
}

func (r *CreateForwardingPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateForwardingPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateForwardingRuleRequestParams struct {
	// <p>Global acceleration instance ID.</p>
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// <p>Listener ID.</p>
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// <p>Policy ID.</p>
	ForwardingPolicyId *string `json:"ForwardingPolicyId,omitnil,omitempty" name:"ForwardingPolicyId"`

	// <p>Layer 7 forwarding rule conditional information.</p><p>Array length cannot exceed 1.</p>
	RuleConditions []*RuleCondition `json:"RuleConditions,omitnil,omitempty" name:"RuleConditions"`

	// <p>Layer 7 forwarding rule behavior information.</p><p>The length of the array cannot exceed 1.</p>
	RuleActions []*RuleAction `json:"RuleActions,omitnil,omitempty" name:"RuleActions"`

	// <p>Origin-pull Header information.</p><p>The maximum length of the array cannot exceed 5. This field is required when RuleActions.RuleActionType is ForwardGroup.</p>
	OriginHeaders []*OriginHeader `json:"OriginHeaders,omitnil,omitempty" name:"OriginHeaders"`

	// <p>Whether origin-pull sni is enabled.</p><p>Default value: False</p><p>This field is required when RuleActions.RuleActionType is ForwardGroup.</p>
	EnableOriginSni *bool `json:"EnableOriginSni,omitnil,omitempty" name:"EnableOriginSni"`

	// <p>Origin sni.</p><p>Input parameter limit: length cannot exceed 80.</p><p>This field is required when EnableOriginSni is True. This field is required when RuleActions.RuleActionType is ForwardGroup.</p>
	OriginSni *string `json:"OriginSni,omitnil,omitempty" name:"OriginSni"`

	// <p>Origin-pull host.</p><p>Input parameter limit: length not exceeding 80.</p><p>This field is required when RuleActions.RuleActionType is ForwardGroup.</p>
	OriginHost *string `json:"OriginHost,omitnil,omitempty" name:"OriginHost"`

	// <p>Origin response headers</p><p>Array length not exceeding 5. An empty array can be passed, representing configuration clearing.</p>
	ResponseHeaders []*ResponseHeaders `json:"ResponseHeaders,omitnil,omitempty" name:"ResponseHeaders"`

	// <p>Delete origin server response headers</p><p>Array length not exceeding 5. An empty array can be passed, representing configuration clearing.</p>
	HideResponseHeaders []*HideResponseHeaders `json:"HideResponseHeaders,omitnil,omitempty" name:"HideResponseHeaders"`
}

type CreateForwardingRuleRequest struct {
	*tchttp.BaseRequest
	
	// <p>Global acceleration instance ID.</p>
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// <p>Listener ID.</p>
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// <p>Policy ID.</p>
	ForwardingPolicyId *string `json:"ForwardingPolicyId,omitnil,omitempty" name:"ForwardingPolicyId"`

	// <p>Layer 7 forwarding rule conditional information.</p><p>Array length cannot exceed 1.</p>
	RuleConditions []*RuleCondition `json:"RuleConditions,omitnil,omitempty" name:"RuleConditions"`

	// <p>Layer 7 forwarding rule behavior information.</p><p>The length of the array cannot exceed 1.</p>
	RuleActions []*RuleAction `json:"RuleActions,omitnil,omitempty" name:"RuleActions"`

	// <p>Origin-pull Header information.</p><p>The maximum length of the array cannot exceed 5. This field is required when RuleActions.RuleActionType is ForwardGroup.</p>
	OriginHeaders []*OriginHeader `json:"OriginHeaders,omitnil,omitempty" name:"OriginHeaders"`

	// <p>Whether origin-pull sni is enabled.</p><p>Default value: False</p><p>This field is required when RuleActions.RuleActionType is ForwardGroup.</p>
	EnableOriginSni *bool `json:"EnableOriginSni,omitnil,omitempty" name:"EnableOriginSni"`

	// <p>Origin sni.</p><p>Input parameter limit: length cannot exceed 80.</p><p>This field is required when EnableOriginSni is True. This field is required when RuleActions.RuleActionType is ForwardGroup.</p>
	OriginSni *string `json:"OriginSni,omitnil,omitempty" name:"OriginSni"`

	// <p>Origin-pull host.</p><p>Input parameter limit: length not exceeding 80.</p><p>This field is required when RuleActions.RuleActionType is ForwardGroup.</p>
	OriginHost *string `json:"OriginHost,omitnil,omitempty" name:"OriginHost"`

	// <p>Origin response headers</p><p>Array length not exceeding 5. An empty array can be passed, representing configuration clearing.</p>
	ResponseHeaders []*ResponseHeaders `json:"ResponseHeaders,omitnil,omitempty" name:"ResponseHeaders"`

	// <p>Delete origin server response headers</p><p>Array length not exceeding 5. An empty array can be passed, representing configuration clearing.</p>
	HideResponseHeaders []*HideResponseHeaders `json:"HideResponseHeaders,omitnil,omitempty" name:"HideResponseHeaders"`
}

func (r *CreateForwardingRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateForwardingRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GlobalAcceleratorId")
	delete(f, "ListenerId")
	delete(f, "ForwardingPolicyId")
	delete(f, "RuleConditions")
	delete(f, "RuleActions")
	delete(f, "OriginHeaders")
	delete(f, "EnableOriginSni")
	delete(f, "OriginSni")
	delete(f, "OriginHost")
	delete(f, "ResponseHeaders")
	delete(f, "HideResponseHeaders")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateForwardingRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateForwardingRuleResponseParams struct {
	// <p>Asynchronous task ID.</p>
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// <p>Layer 7 forwarding rule ID.</p>
	ForwardingRuleId *string `json:"ForwardingRuleId,omitnil,omitempty" name:"ForwardingRuleId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateForwardingRuleResponse struct {
	*tchttp.BaseResponse
	Response *CreateForwardingRuleResponseParams `json:"Response"`
}

func (r *CreateForwardingRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateForwardingRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateGlobalAcceleratorAccessLogRequestParams struct {
	// <p>Unique Id of the sample GA</p>
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// <p>Listener Id</p>
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// <p>Terminal node group Id</p>
	EndpointGroupId *string `json:"EndpointGroupId,omitnil,omitempty" name:"EndpointGroupId"`

	// <p>Logset region</p>
	CloudRegion *string `json:"CloudRegion,omitnil,omitempty" name:"CloudRegion"`

	// <p>Log topic Id</p>
	CloudLogId *string `json:"CloudLogId,omitnil,omitempty" name:"CloudLogId"`

	// <p>Log Set Id</p>
	CloudLogSetId *string `json:"CloudLogSetId,omitnil,omitempty" name:"CloudLogSetId"`

	// <p>Specify data collection fields</p><p>Enumeration values:</p><ul><li>session_time: Layer 4, session duration</li><li>upstream_bytes_received: Layer 4 and Layer 7, number of bytes received from the terminal node</li><li>upstream_bytes_sent: Layer 4 and Layer 7, number of bytes sent to the terminal node</li><li>request_method: Layer 7, GET/POST</li><li>scheme: Layer 7, http/https</li><li>request_uri: Layer 7, uri of the client's raw request</li><li>uri: Layer 7, uri of the current request</li><li>host: Layer 7, domain name accessed by the client (Layer 7)</li><li>remote_user: Layer 7, userName for basic authentication ("-" when unauthenticated)</li><li>http_user_agent: Layer 7, client browser identification</li><li>http_referer: Layer 7, request source URL ("-" when accessed directly from the address bar)</li><li>http_x_forwarded_for: Layer 7, records the client's original IP and the proxy server IP chain it transited</li><li>content_type: Layer 7, content_type</li><li>body_bytes_sent: Layer 7, http body size sent to the client, excluding the header</li><li>request_time: Layer 7, total time from receiving the first byte of the client request to sending the last byte of the response (unit: seconds)</li><li>sent_http_content_type: Layer 7, response content type</li><li>upstream_header_time: Layer 7, arrival time of the terminal node's response header</li><li>upstream_response_length: Layer 7, length of the response body returned by the terminal node</li><li>upstream_response_time: Layer 7, complete response time of the terminal node</li><li>upstream_status: Layer 7, http status code returned by the terminal node</li></ul>
	FieldKeys []*string `json:"FieldKeys,omitnil,omitempty" name:"FieldKeys"`

	// <p>Log description</p>
	FlowLogDescription *string `json:"FlowLogDescription,omitnil,omitempty" name:"FlowLogDescription"`
}

type CreateGlobalAcceleratorAccessLogRequest struct {
	*tchttp.BaseRequest
	
	// <p>Unique Id of the sample GA</p>
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// <p>Listener Id</p>
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// <p>Terminal node group Id</p>
	EndpointGroupId *string `json:"EndpointGroupId,omitnil,omitempty" name:"EndpointGroupId"`

	// <p>Logset region</p>
	CloudRegion *string `json:"CloudRegion,omitnil,omitempty" name:"CloudRegion"`

	// <p>Log topic Id</p>
	CloudLogId *string `json:"CloudLogId,omitnil,omitempty" name:"CloudLogId"`

	// <p>Log Set Id</p>
	CloudLogSetId *string `json:"CloudLogSetId,omitnil,omitempty" name:"CloudLogSetId"`

	// <p>Specify data collection fields</p><p>Enumeration values:</p><ul><li>session_time: Layer 4, session duration</li><li>upstream_bytes_received: Layer 4 and Layer 7, number of bytes received from the terminal node</li><li>upstream_bytes_sent: Layer 4 and Layer 7, number of bytes sent to the terminal node</li><li>request_method: Layer 7, GET/POST</li><li>scheme: Layer 7, http/https</li><li>request_uri: Layer 7, uri of the client's raw request</li><li>uri: Layer 7, uri of the current request</li><li>host: Layer 7, domain name accessed by the client (Layer 7)</li><li>remote_user: Layer 7, userName for basic authentication ("-" when unauthenticated)</li><li>http_user_agent: Layer 7, client browser identification</li><li>http_referer: Layer 7, request source URL ("-" when accessed directly from the address bar)</li><li>http_x_forwarded_for: Layer 7, records the client's original IP and the proxy server IP chain it transited</li><li>content_type: Layer 7, content_type</li><li>body_bytes_sent: Layer 7, http body size sent to the client, excluding the header</li><li>request_time: Layer 7, total time from receiving the first byte of the client request to sending the last byte of the response (unit: seconds)</li><li>sent_http_content_type: Layer 7, response content type</li><li>upstream_header_time: Layer 7, arrival time of the terminal node's response header</li><li>upstream_response_length: Layer 7, length of the response body returned by the terminal node</li><li>upstream_response_time: Layer 7, complete response time of the terminal node</li><li>upstream_status: Layer 7, http status code returned by the terminal node</li></ul>
	FieldKeys []*string `json:"FieldKeys,omitnil,omitempty" name:"FieldKeys"`

	// <p>Log description</p>
	FlowLogDescription *string `json:"FlowLogDescription,omitnil,omitempty" name:"FlowLogDescription"`
}

func (r *CreateGlobalAcceleratorAccessLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateGlobalAcceleratorAccessLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GlobalAcceleratorId")
	delete(f, "ListenerId")
	delete(f, "EndpointGroupId")
	delete(f, "CloudRegion")
	delete(f, "CloudLogId")
	delete(f, "CloudLogSetId")
	delete(f, "FieldKeys")
	delete(f, "FlowLogDescription")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateGlobalAcceleratorAccessLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateGlobalAcceleratorAccessLogResponseParams struct {
	// <p>Log Task Unique Id</p>
	LogPushTaskId *string `json:"LogPushTaskId,omitnil,omitempty" name:"LogPushTaskId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateGlobalAcceleratorAccessLogResponse struct {
	*tchttp.BaseResponse
	Response *CreateGlobalAcceleratorAccessLogResponseParams `json:"Response"`
}

func (r *CreateGlobalAcceleratorAccessLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateGlobalAcceleratorAccessLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateGlobalAcceleratorAclPolicyRequestParams struct {
	// <p>Global acceleration instance ID.</p>
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// <p>Default behavior.</p><p>Enumeration values:</p><ul><li>ACCEPT: Permit all traffic on the access channel by default</li><li>DROP: Deny all traffic on the access channel by default</li></ul>
	DefaultAction *string `json:"DefaultAction,omitnil,omitempty" name:"DefaultAction"`
}

type CreateGlobalAcceleratorAclPolicyRequest struct {
	*tchttp.BaseRequest
	
	// <p>Global acceleration instance ID.</p>
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// <p>Default behavior.</p><p>Enumeration values:</p><ul><li>ACCEPT: Permit all traffic on the access channel by default</li><li>DROP: Deny all traffic on the access channel by default</li></ul>
	DefaultAction *string `json:"DefaultAction,omitnil,omitempty" name:"DefaultAction"`
}

func (r *CreateGlobalAcceleratorAclPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateGlobalAcceleratorAclPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GlobalAcceleratorId")
	delete(f, "DefaultAction")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateGlobalAcceleratorAclPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateGlobalAcceleratorAclPolicyResponseParams struct {
	// <p>Asynchronous task ID.</p>
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// <p>Access control policy ID.</p>
	GlobalAcceleratorAclPolicyId *string `json:"GlobalAcceleratorAclPolicyId,omitnil,omitempty" name:"GlobalAcceleratorAclPolicyId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateGlobalAcceleratorAclPolicyResponse struct {
	*tchttp.BaseResponse
	Response *CreateGlobalAcceleratorAclPolicyResponseParams `json:"Response"`
}

func (r *CreateGlobalAcceleratorAclPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateGlobalAcceleratorAclPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateGlobalAcceleratorAclRuleRequestParams struct {
	// Global acceleration instance ID.
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// Security policy ID
	GlobalAcceleratorAclPolicyId *string `json:"GlobalAcceleratorAclPolicyId,omitnil,omitempty" name:"GlobalAcceleratorAclPolicyId"`

	// Acl information.
	AclEntries []*AclEntries `json:"AclEntries,omitnil,omitempty" name:"AclEntries"`
}

type CreateGlobalAcceleratorAclRuleRequest struct {
	*tchttp.BaseRequest
	
	// Global acceleration instance ID.
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// Security policy ID
	GlobalAcceleratorAclPolicyId *string `json:"GlobalAcceleratorAclPolicyId,omitnil,omitempty" name:"GlobalAcceleratorAclPolicyId"`

	// Acl information.
	AclEntries []*AclEntries `json:"AclEntries,omitnil,omitempty" name:"AclEntries"`
}

func (r *CreateGlobalAcceleratorAclRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateGlobalAcceleratorAclRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GlobalAcceleratorId")
	delete(f, "GlobalAcceleratorAclPolicyId")
	delete(f, "AclEntries")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateGlobalAcceleratorAclRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateGlobalAcceleratorAclRuleResponseParams struct {
	// Asynchronous task ID.
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// ACL rule ID.
	GlobalAcceleratorAclRuleIds []*string `json:"GlobalAcceleratorAclRuleIds,omitnil,omitempty" name:"GlobalAcceleratorAclRuleIds"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateGlobalAcceleratorAclRuleResponse struct {
	*tchttp.BaseResponse
	Response *CreateGlobalAcceleratorAclRuleResponseParams `json:"Response"`
}

func (r *CreateGlobalAcceleratorAclRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateGlobalAcceleratorAclRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateGlobalAcceleratorRequestParams struct {
	// <p>Name.</p><p>Parameter format: starting with a letter or Chinese characters, 2–128 characters in length, supporting letters, digits, Chinese characters, . - _</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Billing mode. PREPAID: prepaid mode, i.e., Monthly Subscription. POSTPAID: postpaid, i.e., pay-as-you-go. Default: POSTPAID. Currently, only pay-as-you-go is supported.</p>
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// <p>Description.</p><p>Parameter format: should not exceed 100 characters.</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>Cross-border type; HighQuality: dedicated BGP-IP cross-border; Unicom: China Unicom Direct Connect cross-border.</p>
	CrossBorderType *string `json:"CrossBorderType,omitnil,omitempty" name:"CrossBorderType"`

	// <p>This Flag represents signing the cross-border service commitment. When using cross-border service, this field is required. True: represents signed.</p>
	CrossBorderPromiseFlag *bool `json:"CrossBorderPromiseFlag,omitnil,omitempty" name:"CrossBorderPromiseFlag"`

	// <p>Tag information.</p>
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type CreateGlobalAcceleratorRequest struct {
	*tchttp.BaseRequest
	
	// <p>Name.</p><p>Parameter format: starting with a letter or Chinese characters, 2–128 characters in length, supporting letters, digits, Chinese characters, . - _</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Billing mode. PREPAID: prepaid mode, i.e., Monthly Subscription. POSTPAID: postpaid, i.e., pay-as-you-go. Default: POSTPAID. Currently, only pay-as-you-go is supported.</p>
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// <p>Description.</p><p>Parameter format: should not exceed 100 characters.</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>Cross-border type; HighQuality: dedicated BGP-IP cross-border; Unicom: China Unicom Direct Connect cross-border.</p>
	CrossBorderType *string `json:"CrossBorderType,omitnil,omitempty" name:"CrossBorderType"`

	// <p>This Flag represents signing the cross-border service commitment. When using cross-border service, this field is required. True: represents signed.</p>
	CrossBorderPromiseFlag *bool `json:"CrossBorderPromiseFlag,omitnil,omitempty" name:"CrossBorderPromiseFlag"`

	// <p>Tag information.</p>
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *CreateGlobalAcceleratorRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateGlobalAcceleratorRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "InstanceChargeType")
	delete(f, "Description")
	delete(f, "CrossBorderType")
	delete(f, "CrossBorderPromiseFlag")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateGlobalAcceleratorRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateGlobalAcceleratorResponseParams struct {
	// <p>Task ID.</p>
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// <p>Global acceleration instance ID.</p>
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateGlobalAcceleratorResponse struct {
	*tchttp.BaseResponse
	Response *CreateGlobalAcceleratorResponseParams `json:"Response"`
}

func (r *CreateGlobalAcceleratorResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateGlobalAcceleratorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateListenerAdditionalCertRequestParams struct {
	// <p>Global acceleration instance ID.</p>
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// <p>Listener ID.</p>
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// <p>Certificate ID.</p><p>Currently, only server certificates can be added.</p>
	AdditionalCertificates []*string `json:"AdditionalCertificates,omitnil,omitempty" name:"AdditionalCertificates"`
}

type CreateListenerAdditionalCertRequest struct {
	*tchttp.BaseRequest
	
	// <p>Global acceleration instance ID.</p>
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// <p>Listener ID.</p>
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// <p>Certificate ID.</p><p>Currently, only server certificates can be added.</p>
	AdditionalCertificates []*string `json:"AdditionalCertificates,omitnil,omitempty" name:"AdditionalCertificates"`
}

func (r *CreateListenerAdditionalCertRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateListenerAdditionalCertRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GlobalAcceleratorId")
	delete(f, "ListenerId")
	delete(f, "AdditionalCertificates")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateListenerAdditionalCertRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateListenerAdditionalCertResponseParams struct {
	// <p>Task ID.</p>
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateListenerAdditionalCertResponse struct {
	*tchttp.BaseResponse
	Response *CreateListenerAdditionalCertResponseParams `json:"Response"`
}

func (r *CreateListenerAdditionalCertResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateListenerAdditionalCertResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateListenerRequestParams struct {
	// <p>Global acceleration instance ID.</p>
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// <p>Name.</p><p>Parameter format: starting with a letter or Chinese characters, 2–128 characters in length, supporting letters, digits, Chinese characters, . - _</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Port range.</p>
	PortRanges *PortRanges `json:"PortRanges,omitnil,omitempty" name:"PortRanges"`

	// <p>Description. Maximum length cannot exceed 100 characters.</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>Listening type, defaults to smart routing.</p><p>Enumeration values:</p><ul><li>Standard: Smart routing.</li></ul>
	ListenerType *string `json:"ListenerType,omitnil,omitempty" name:"ListenerType"`

	// <p>Protocol. Default value: TCP. Supports configuration of 'TCP', 'UDP', 'HTTP', and 'HTTPS'.</p>
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// <p>Connection idle wait time.</p><p>1. For HTTP/HTTPS listener, the default value is 15, with a supported range of 1-60.<br>2. For TCP listener, the default value is 900, with a supported range of 10-900.<br>3. For UDP listener, the default value is 20, with a supported range of 10-20.</p>
	IdleTimeout *uint64 `json:"IdleTimeout,omitnil,omitempty" name:"IdleTimeout"`

	// <p>Layer-4 source IP retrieval mode. Supports 'TOA', 'ProxyProtocol', and 'ProxyProtocolV2'.</p><p>This parameter can be filled in only when the Layer-4 source IP retrieval mode is enabled.</p>
	GetRealIpType *string `json:"GetRealIpType,omitnil,omitempty" name:"GetRealIpType"`

	// <p>Whether to enable session persistence. Supports configuration of 'Open' and 'Close'.</p><p>Enumeration values:</p><ul><li>Open: enable.</li><li>Close: disable.</li></ul><p>Only supported for layer-4 listeners. For layer-7 listeners, modification is not supported.</p>
	ClientAffinity *string `json:"ClientAffinity,omitnil,omitempty" name:"ClientAffinity"`

	// <p>Request timeout.</p><p>Value range: [1, 180]</p><p>Default value: 60</p><p>This parameter is configurable only for HTTPS listeners.</p>
	RequestTimeout *uint64 `json:"RequestTimeout,omitnil,omitempty" name:"RequestTimeout"`

	// <p>Whether to enable layer-7 source IP retrieval mode.</p>
	XForwardedForRealIp *bool `json:"XForwardedForRealIp,omitnil,omitempty" name:"XForwardedForRealIp"`

	// <p>Parsing method.</p><p>Enumeration values:</p><ul><li>UNIDIRECTIONAL: two-way.</li><li>U: one-way.</li></ul><p>For an HTTPS listener, this field is required.</p>
	CertificationType *string `json:"CertificationType,omitnil,omitempty" name:"CertificationType"`

	// <p>Encryption algorithm kit. Supports configuration of 'tls_policy_1.0-2', 'tls_policy_1.1-2', 'tls_policy_1.2', 'tls_policy_1.2_strict', 'tls_policy_1.2_strict-1.3'.</p>
	CipherPolicyId *string `json:"CipherPolicyId,omitnil,omitempty" name:"CipherPolicyId"`

	// <p>Server certificate.</p><p>Input limit: currently only support importing one cert; to use multiple certs, use the cert api CreateListenerAdditionalCert to add other certs.</p><p>This field is required for HTTPS listeners.</p>
	ServerCertificates []*string `json:"ServerCertificates,omitnil,omitempty" name:"ServerCertificates"`

	// <p>Client certificate.</p><p>Input limit: 1. Currently only support importing one cert. To use multiple certs, use the cert api CreateListenerAdditionalCert to add other certs. 2. The cert must be a CA certificate.</p><p>This field is required when HTTPS listener and mutual authentication are enabled.</p>
	ClientCaCertificates []*string `json:"ClientCaCertificates,omitnil,omitempty" name:"ClientCaCertificates"`

	// <p>HTTPS listener supports version selection</p><p>Enumeration values:</p><ul><li>HTTP/1.1: HTTP/1.1</li><li>HTTP/2: HTTP/2</li></ul>
	HttpVersion *string `json:"HttpVersion,omitnil,omitempty" name:"HttpVersion"`
}

type CreateListenerRequest struct {
	*tchttp.BaseRequest
	
	// <p>Global acceleration instance ID.</p>
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// <p>Name.</p><p>Parameter format: starting with a letter or Chinese characters, 2–128 characters in length, supporting letters, digits, Chinese characters, . - _</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Port range.</p>
	PortRanges *PortRanges `json:"PortRanges,omitnil,omitempty" name:"PortRanges"`

	// <p>Description. Maximum length cannot exceed 100 characters.</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>Listening type, defaults to smart routing.</p><p>Enumeration values:</p><ul><li>Standard: Smart routing.</li></ul>
	ListenerType *string `json:"ListenerType,omitnil,omitempty" name:"ListenerType"`

	// <p>Protocol. Default value: TCP. Supports configuration of 'TCP', 'UDP', 'HTTP', and 'HTTPS'.</p>
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// <p>Connection idle wait time.</p><p>1. For HTTP/HTTPS listener, the default value is 15, with a supported range of 1-60.<br>2. For TCP listener, the default value is 900, with a supported range of 10-900.<br>3. For UDP listener, the default value is 20, with a supported range of 10-20.</p>
	IdleTimeout *uint64 `json:"IdleTimeout,omitnil,omitempty" name:"IdleTimeout"`

	// <p>Layer-4 source IP retrieval mode. Supports 'TOA', 'ProxyProtocol', and 'ProxyProtocolV2'.</p><p>This parameter can be filled in only when the Layer-4 source IP retrieval mode is enabled.</p>
	GetRealIpType *string `json:"GetRealIpType,omitnil,omitempty" name:"GetRealIpType"`

	// <p>Whether to enable session persistence. Supports configuration of 'Open' and 'Close'.</p><p>Enumeration values:</p><ul><li>Open: enable.</li><li>Close: disable.</li></ul><p>Only supported for layer-4 listeners. For layer-7 listeners, modification is not supported.</p>
	ClientAffinity *string `json:"ClientAffinity,omitnil,omitempty" name:"ClientAffinity"`

	// <p>Request timeout.</p><p>Value range: [1, 180]</p><p>Default value: 60</p><p>This parameter is configurable only for HTTPS listeners.</p>
	RequestTimeout *uint64 `json:"RequestTimeout,omitnil,omitempty" name:"RequestTimeout"`

	// <p>Whether to enable layer-7 source IP retrieval mode.</p>
	XForwardedForRealIp *bool `json:"XForwardedForRealIp,omitnil,omitempty" name:"XForwardedForRealIp"`

	// <p>Parsing method.</p><p>Enumeration values:</p><ul><li>UNIDIRECTIONAL: two-way.</li><li>U: one-way.</li></ul><p>For an HTTPS listener, this field is required.</p>
	CertificationType *string `json:"CertificationType,omitnil,omitempty" name:"CertificationType"`

	// <p>Encryption algorithm kit. Supports configuration of 'tls_policy_1.0-2', 'tls_policy_1.1-2', 'tls_policy_1.2', 'tls_policy_1.2_strict', 'tls_policy_1.2_strict-1.3'.</p>
	CipherPolicyId *string `json:"CipherPolicyId,omitnil,omitempty" name:"CipherPolicyId"`

	// <p>Server certificate.</p><p>Input limit: currently only support importing one cert; to use multiple certs, use the cert api CreateListenerAdditionalCert to add other certs.</p><p>This field is required for HTTPS listeners.</p>
	ServerCertificates []*string `json:"ServerCertificates,omitnil,omitempty" name:"ServerCertificates"`

	// <p>Client certificate.</p><p>Input limit: 1. Currently only support importing one cert. To use multiple certs, use the cert api CreateListenerAdditionalCert to add other certs. 2. The cert must be a CA certificate.</p><p>This field is required when HTTPS listener and mutual authentication are enabled.</p>
	ClientCaCertificates []*string `json:"ClientCaCertificates,omitnil,omitempty" name:"ClientCaCertificates"`

	// <p>HTTPS listener supports version selection</p><p>Enumeration values:</p><ul><li>HTTP/1.1: HTTP/1.1</li><li>HTTP/2: HTTP/2</li></ul>
	HttpVersion *string `json:"HttpVersion,omitnil,omitempty" name:"HttpVersion"`
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
	delete(f, "GlobalAcceleratorId")
	delete(f, "Name")
	delete(f, "PortRanges")
	delete(f, "Description")
	delete(f, "ListenerType")
	delete(f, "Protocol")
	delete(f, "IdleTimeout")
	delete(f, "GetRealIpType")
	delete(f, "ClientAffinity")
	delete(f, "RequestTimeout")
	delete(f, "XForwardedForRealIp")
	delete(f, "CertificationType")
	delete(f, "CipherPolicyId")
	delete(f, "ServerCertificates")
	delete(f, "ClientCaCertificates")
	delete(f, "HttpVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateListenerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateListenerResponseParams struct {
	// <p>Task ID.</p>
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// <p>Listener ID.</p>
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DeleteAccelerateAreasRequestParams struct {
	// Global acceleration instance ID.
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// Acceleration region ID.
	AcceleratorAreaIds []*string `json:"AcceleratorAreaIds,omitnil,omitempty" name:"AcceleratorAreaIds"`
}

type DeleteAccelerateAreasRequest struct {
	*tchttp.BaseRequest
	
	// Global acceleration instance ID.
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// Acceleration region ID.
	AcceleratorAreaIds []*string `json:"AcceleratorAreaIds,omitnil,omitempty" name:"AcceleratorAreaIds"`
}

func (r *DeleteAccelerateAreasRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAccelerateAreasRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GlobalAcceleratorId")
	delete(f, "AcceleratorAreaIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAccelerateAreasRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAccelerateAreasResponseParams struct {
	// Asynchronous task ID.
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteAccelerateAreasResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAccelerateAreasResponseParams `json:"Response"`
}

func (r *DeleteAccelerateAreasResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAccelerateAreasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteEndpointGroupsRequestParams struct {
	// Global acceleration instance ID.
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// Listener ID.
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// Terminal node group ID.
	EndpointGroupIds []*string `json:"EndpointGroupIds,omitnil,omitempty" name:"EndpointGroupIds"`
}

type DeleteEndpointGroupsRequest struct {
	*tchttp.BaseRequest
	
	// Global acceleration instance ID.
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// Listener ID.
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// Terminal node group ID.
	EndpointGroupIds []*string `json:"EndpointGroupIds,omitnil,omitempty" name:"EndpointGroupIds"`
}

func (r *DeleteEndpointGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteEndpointGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GlobalAcceleratorId")
	delete(f, "ListenerId")
	delete(f, "EndpointGroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteEndpointGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteEndpointGroupsResponseParams struct {
	// Task ID.
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteEndpointGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteEndpointGroupsResponseParams `json:"Response"`
}

func (r *DeleteEndpointGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteEndpointGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteForwardingPolicyRequestParams struct {
	// Global acceleration instance ID.
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// Listener ID.
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// Policy ID.
	ForwardingPolicyId *string `json:"ForwardingPolicyId,omitnil,omitempty" name:"ForwardingPolicyId"`
}

type DeleteForwardingPolicyRequest struct {
	*tchttp.BaseRequest
	
	// Global acceleration instance ID.
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// Listener ID.
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// Policy ID.
	ForwardingPolicyId *string `json:"ForwardingPolicyId,omitnil,omitempty" name:"ForwardingPolicyId"`
}

func (r *DeleteForwardingPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteForwardingPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GlobalAcceleratorId")
	delete(f, "ListenerId")
	delete(f, "ForwardingPolicyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteForwardingPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteForwardingPolicyResponseParams struct {
	// Asynchronous task ID.
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteForwardingPolicyResponse struct {
	*tchttp.BaseResponse
	Response *DeleteForwardingPolicyResponseParams `json:"Response"`
}

func (r *DeleteForwardingPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteForwardingPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteForwardingRuleRequestParams struct {
	// Global acceleration instance ID.
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// Listener ID.
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// Policy ID.
	ForwardingPolicyId *string `json:"ForwardingPolicyId,omitnil,omitempty" name:"ForwardingPolicyId"`

	// Layer 7 forwarding rule ID.
	ForwardingRuleId *string `json:"ForwardingRuleId,omitnil,omitempty" name:"ForwardingRuleId"`
}

type DeleteForwardingRuleRequest struct {
	*tchttp.BaseRequest
	
	// Global acceleration instance ID.
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// Listener ID.
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// Policy ID.
	ForwardingPolicyId *string `json:"ForwardingPolicyId,omitnil,omitempty" name:"ForwardingPolicyId"`

	// Layer 7 forwarding rule ID.
	ForwardingRuleId *string `json:"ForwardingRuleId,omitnil,omitempty" name:"ForwardingRuleId"`
}

func (r *DeleteForwardingRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteForwardingRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GlobalAcceleratorId")
	delete(f, "ListenerId")
	delete(f, "ForwardingPolicyId")
	delete(f, "ForwardingRuleId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteForwardingRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteForwardingRuleResponseParams struct {
	// Asynchronous task ID.
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteForwardingRuleResponse struct {
	*tchttp.BaseResponse
	Response *DeleteForwardingRuleResponseParams `json:"Response"`
}

func (r *DeleteForwardingRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteForwardingRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteGlobalAcceleratorAccessLogRequestParams struct {
	// <p>Log Unique Id</p>
	LogPushTaskId *string `json:"LogPushTaskId,omitnil,omitempty" name:"LogPushTaskId"`

	// <p>Unique Id of the GA instance</p>
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`
}

type DeleteGlobalAcceleratorAccessLogRequest struct {
	*tchttp.BaseRequest
	
	// <p>Log Unique Id</p>
	LogPushTaskId *string `json:"LogPushTaskId,omitnil,omitempty" name:"LogPushTaskId"`

	// <p>Unique Id of the GA instance</p>
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`
}

func (r *DeleteGlobalAcceleratorAccessLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteGlobalAcceleratorAccessLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LogPushTaskId")
	delete(f, "GlobalAcceleratorId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteGlobalAcceleratorAccessLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteGlobalAcceleratorAccessLogResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteGlobalAcceleratorAccessLogResponse struct {
	*tchttp.BaseResponse
	Response *DeleteGlobalAcceleratorAccessLogResponseParams `json:"Response"`
}

func (r *DeleteGlobalAcceleratorAccessLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteGlobalAcceleratorAccessLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteGlobalAcceleratorAclPolicyRequestParams struct {
	// Global acceleration instance ID.
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// Access control policy ID.
	GlobalAcceleratorAclPolicyId *string `json:"GlobalAcceleratorAclPolicyId,omitnil,omitempty" name:"GlobalAcceleratorAclPolicyId"`
}

type DeleteGlobalAcceleratorAclPolicyRequest struct {
	*tchttp.BaseRequest
	
	// Global acceleration instance ID.
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// Access control policy ID.
	GlobalAcceleratorAclPolicyId *string `json:"GlobalAcceleratorAclPolicyId,omitnil,omitempty" name:"GlobalAcceleratorAclPolicyId"`
}

func (r *DeleteGlobalAcceleratorAclPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteGlobalAcceleratorAclPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GlobalAcceleratorId")
	delete(f, "GlobalAcceleratorAclPolicyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteGlobalAcceleratorAclPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteGlobalAcceleratorAclPolicyResponseParams struct {
	// Asynchronous task ID.
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteGlobalAcceleratorAclPolicyResponse struct {
	*tchttp.BaseResponse
	Response *DeleteGlobalAcceleratorAclPolicyResponseParams `json:"Response"`
}

func (r *DeleteGlobalAcceleratorAclPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteGlobalAcceleratorAclPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteGlobalAcceleratorAclRuleRequestParams struct {
	// Global acceleration instance ID.
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// Security policy ID
	GlobalAcceleratorAclPolicyId *string `json:"GlobalAcceleratorAclPolicyId,omitnil,omitempty" name:"GlobalAcceleratorAclPolicyId"`

	// Acl rule ID.
	GlobalAcceleratorAclRuleIds []*string `json:"GlobalAcceleratorAclRuleIds,omitnil,omitempty" name:"GlobalAcceleratorAclRuleIds"`
}

type DeleteGlobalAcceleratorAclRuleRequest struct {
	*tchttp.BaseRequest
	
	// Global acceleration instance ID.
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// Security policy ID
	GlobalAcceleratorAclPolicyId *string `json:"GlobalAcceleratorAclPolicyId,omitnil,omitempty" name:"GlobalAcceleratorAclPolicyId"`

	// Acl rule ID.
	GlobalAcceleratorAclRuleIds []*string `json:"GlobalAcceleratorAclRuleIds,omitnil,omitempty" name:"GlobalAcceleratorAclRuleIds"`
}

func (r *DeleteGlobalAcceleratorAclRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteGlobalAcceleratorAclRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GlobalAcceleratorId")
	delete(f, "GlobalAcceleratorAclPolicyId")
	delete(f, "GlobalAcceleratorAclRuleIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteGlobalAcceleratorAclRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteGlobalAcceleratorAclRuleResponseParams struct {
	// Asynchronous task ID.
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteGlobalAcceleratorAclRuleResponse struct {
	*tchttp.BaseResponse
	Response *DeleteGlobalAcceleratorAclRuleResponseParams `json:"Response"`
}

func (r *DeleteGlobalAcceleratorAclRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteGlobalAcceleratorAclRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteGlobalAcceleratorRequestParams struct {
	// Global acceleration instance ID.
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`
}

type DeleteGlobalAcceleratorRequest struct {
	*tchttp.BaseRequest
	
	// Global acceleration instance ID.
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`
}

func (r *DeleteGlobalAcceleratorRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteGlobalAcceleratorRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GlobalAcceleratorId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteGlobalAcceleratorRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteGlobalAcceleratorResponseParams struct {
	// Task ID.
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteGlobalAcceleratorResponse struct {
	*tchttp.BaseResponse
	Response *DeleteGlobalAcceleratorResponseParams `json:"Response"`
}

func (r *DeleteGlobalAcceleratorResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteGlobalAcceleratorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteListenerAdditionalCertRequestParams struct {
	// Global acceleration instance ID.
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// Listener ID.
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// Certificate ID.
	AdditionalCertificates []*string `json:"AdditionalCertificates,omitnil,omitempty" name:"AdditionalCertificates"`
}

type DeleteListenerAdditionalCertRequest struct {
	*tchttp.BaseRequest
	
	// Global acceleration instance ID.
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// Listener ID.
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// Certificate ID.
	AdditionalCertificates []*string `json:"AdditionalCertificates,omitnil,omitempty" name:"AdditionalCertificates"`
}

func (r *DeleteListenerAdditionalCertRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteListenerAdditionalCertRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GlobalAcceleratorId")
	delete(f, "ListenerId")
	delete(f, "AdditionalCertificates")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteListenerAdditionalCertRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteListenerAdditionalCertResponseParams struct {
	// Task ID.
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteListenerAdditionalCertResponse struct {
	*tchttp.BaseResponse
	Response *DeleteListenerAdditionalCertResponseParams `json:"Response"`
}

func (r *DeleteListenerAdditionalCertResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteListenerAdditionalCertResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteListenerRequestParams struct {
	// Global acceleration instance ID.
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// Listener ID.
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`
}

type DeleteListenerRequest struct {
	*tchttp.BaseRequest
	
	// Global acceleration instance ID.
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// Listener ID.
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`
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
	delete(f, "GlobalAcceleratorId")
	delete(f, "ListenerId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteListenerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteListenerResponseParams struct {
	// Task ID.
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeAccelerateAreasRequestParams struct {
	// <p>Global acceleration instance ID.</p>
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// <p>Offset. The default is 0.</p>
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>Number of instances meeting conditions. Default value: 20. Maximum: 200.</p>
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>Filter criteria. accelerate-region - String - (Filter criterion) Terminal node group region.</p>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeAccelerateAreasRequest struct {
	*tchttp.BaseRequest
	
	// <p>Global acceleration instance ID.</p>
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// <p>Offset. The default is 0.</p>
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>Number of instances meeting conditions. Default value: 20. Maximum: 200.</p>
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>Filter criteria. accelerate-region - String - (Filter criterion) Terminal node group region.</p>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeAccelerateAreasRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccelerateAreasRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GlobalAcceleratorId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccelerateAreasRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccelerateAreasResponseParams struct {
	// <p>Acceleration region information.</p>
	AccelerateAreaSet []*AcceleratorAreas `json:"AccelerateAreaSet,omitnil,omitempty" name:"AccelerateAreaSet"`

	// <p>Number of instances.</p>
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAccelerateAreasResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAccelerateAreasResponseParams `json:"Response"`
}

func (r *DescribeAccelerateAreasResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccelerateAreasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccelerateRegionsRequestParams struct {

}

type DescribeAccelerateRegionsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeAccelerateRegionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccelerateRegionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccelerateRegionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccelerateRegionsResponseParams struct {
	// Acceleration region information.
	AcceleratorRegionSet []*AcceleratorRegionSet `json:"AcceleratorRegionSet,omitnil,omitempty" name:"AcceleratorRegionSet"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAccelerateRegionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAccelerateRegionsResponseParams `json:"Response"`
}

func (r *DescribeAccelerateRegionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccelerateRegionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccessLogParamRequestParams struct {

	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`
}

type DescribeAccessLogParamRequest struct {
	*tchttp.BaseRequest
	
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`
}

func (r *DescribeAccessLogParamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccessLogParamRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GlobalAcceleratorId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccessLogParamRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccessLogParamResponseParams struct {
	// <p>Layer-7 optional parameter.</p>
	L7Param []*string `json:"L7Param,omitnil,omitempty" name:"L7Param"`

	// <p>L4 optional parameter</p>
	L4Param []*string `json:"L4Param,omitnil,omitempty" name:"L4Param"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAccessLogParamResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAccessLogParamResponseParams `json:"Response"`
}

func (r *DescribeAccessLogParamResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccessLogParamResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCrossBorderSettlementRequestParams struct {
	// Global acceleration instance ID.
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// Acceleration region.
	AccelerateRegion *string `json:"AccelerateRegion,omitnil,omitempty" name:"AccelerateRegion"`

	// Region of the terminal node group.
	EndpointGroupRegion *string `json:"EndpointGroupRegion,omitnil,omitempty" name:"EndpointGroupRegion"`

	// Bill year and month time.
	SettlementMonth *uint64 `json:"SettlementMonth,omitnil,omitempty" name:"SettlementMonth"`
}

type DescribeCrossBorderSettlementRequest struct {
	*tchttp.BaseRequest
	
	// Global acceleration instance ID.
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// Acceleration region.
	AccelerateRegion *string `json:"AccelerateRegion,omitnil,omitempty" name:"AccelerateRegion"`

	// Region of the terminal node group.
	EndpointGroupRegion *string `json:"EndpointGroupRegion,omitnil,omitempty" name:"EndpointGroupRegion"`

	// Bill year and month time.
	SettlementMonth *uint64 `json:"SettlementMonth,omitnil,omitempty" name:"SettlementMonth"`
}

func (r *DescribeCrossBorderSettlementRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCrossBorderSettlementRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GlobalAcceleratorId")
	delete(f, "AccelerateRegion")
	delete(f, "EndpointGroupRegion")
	delete(f, "SettlementMonth")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCrossBorderSettlementRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCrossBorderSettlementResponseParams struct {
	// Traffic amount in GB; precision is reserved to 6 decimal places.
	Traffic *float64 `json:"Traffic,omitnil,omitempty" name:"Traffic"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCrossBorderSettlementResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCrossBorderSettlementResponseParams `json:"Response"`
}

func (r *DescribeCrossBorderSettlementResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCrossBorderSettlementResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEndpointGroupsRequestParams struct {
	// <p>Global acceleration instance ID.</p>
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// <p>Offset. Default value: 0.</p>
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>Number of returns. Default value: 10. Maximum value: 10.</p>
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>Filter criteria. endpoint-group-id - String - (Filter criterion) Terminal node group instance ID. endpoint-group-type - String - (Filter criterion) Terminal node group instance type.</p>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeEndpointGroupsRequest struct {
	*tchttp.BaseRequest
	
	// <p>Global acceleration instance ID.</p>
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// <p>Offset. Default value: 0.</p>
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>Number of returns. Default value: 10. Maximum value: 10.</p>
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>Filter criteria. endpoint-group-id - String - (Filter criterion) Terminal node group instance ID. endpoint-group-type - String - (Filter criterion) Terminal node group instance type.</p>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeEndpointGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEndpointGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GlobalAcceleratorId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEndpointGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEndpointGroupsResponseParams struct {
	// <p>Eligible terminal node group.</p>
	EndpointGroupConfigurationSet []*EndpointGroupConfigurationSet `json:"EndpointGroupConfigurationSet,omitnil,omitempty" name:"EndpointGroupConfigurationSet"`

	// <p>Number of instances that meet the criteria.</p>
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeEndpointGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEndpointGroupsResponseParams `json:"Response"`
}

func (r *DescribeEndpointGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEndpointGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeForwardingPolicyRequestParams struct {
	// Global acceleration instance ID.
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// Listener ID.
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returns. Default value: 20. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeForwardingPolicyRequest struct {
	*tchttp.BaseRequest
	
	// Global acceleration instance ID.
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// Listener ID.
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returns. Default value: 20. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeForwardingPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeForwardingPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GlobalAcceleratorId")
	delete(f, "ListenerId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeForwardingPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeForwardingPolicyResponseParams struct {
	// Policy information that meets the conditions.
	ForwardingPolicySet []*ForwardingPolicySet `json:"ForwardingPolicySet,omitnil,omitempty" name:"ForwardingPolicySet"`

	// Number of instances that meet the criteria.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeForwardingPolicyResponse struct {
	*tchttp.BaseResponse
	Response *DescribeForwardingPolicyResponseParams `json:"Response"`
}

func (r *DescribeForwardingPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeForwardingPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeForwardingRuleRequestParams struct {
	// Global acceleration instance ID.
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// Listener ID.
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// Layer 7 forwarding rule ID.
	ForwardingPolicyId *string `json:"ForwardingPolicyId,omitnil,omitempty" name:"ForwardingPolicyId"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returns. Default value: 20. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeForwardingRuleRequest struct {
	*tchttp.BaseRequest
	
	// Global acceleration instance ID.
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// Listener ID.
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// Layer 7 forwarding rule ID.
	ForwardingPolicyId *string `json:"ForwardingPolicyId,omitnil,omitempty" name:"ForwardingPolicyId"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returns. Default value: 20. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeForwardingRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeForwardingRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GlobalAcceleratorId")
	delete(f, "ListenerId")
	delete(f, "ForwardingPolicyId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeForwardingRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeForwardingRuleResponseParams struct {
	// Rule information that meets the conditions.
	ForwardingRuleSet []*ForwardingRuleSet `json:"ForwardingRuleSet,omitnil,omitempty" name:"ForwardingRuleSet"`

	// Number of instances that meet the criteria.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeForwardingRuleResponse struct {
	*tchttp.BaseResponse
	Response *DescribeForwardingRuleResponseParams `json:"Response"`
}

func (r *DescribeForwardingRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeForwardingRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGlobalAcceleratorAccessLogRequestParams struct {
	// <p>Unique Id of the ga instance</p>
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// <p>Query filter parameters. { &quot;Name&quot;: &quot;listener-id&quot;, &quot;Values&quot;: [&quot;listener unique id&quot;] },{ &quot;Name&quot;: &quot;endpoint-group-id&quot;, &quot;Values&quot;: [&quot;Terminal node group unique id&quot;] },{ &quot;Name&quot;: &quot;access_log_id&quot;, &quot;Values&quot;: [&quot;log unique id&quot;] }</p>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>Offset. Default value: 0.</p>
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>Number of returned results.</p><p>Value range: [0, 200]</p>
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeGlobalAcceleratorAccessLogRequest struct {
	*tchttp.BaseRequest
	
	// <p>Unique Id of the ga instance</p>
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// <p>Query filter parameters. { &quot;Name&quot;: &quot;listener-id&quot;, &quot;Values&quot;: [&quot;listener unique id&quot;] },{ &quot;Name&quot;: &quot;endpoint-group-id&quot;, &quot;Values&quot;: [&quot;Terminal node group unique id&quot;] },{ &quot;Name&quot;: &quot;access_log_id&quot;, &quot;Values&quot;: [&quot;log unique id&quot;] }</p>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>Offset. Default value: 0.</p>
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>Number of returned results.</p><p>Value range: [0, 200]</p>
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeGlobalAcceleratorAccessLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGlobalAcceleratorAccessLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GlobalAcceleratorId")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGlobalAcceleratorAccessLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGlobalAcceleratorAccessLogResponseParams struct {
	// <p>Return log task detail</p>
	GlobalAcceleratorAccessLog []*GlobalAcceleratorAccessLog `json:"GlobalAcceleratorAccessLog,omitnil,omitempty" name:"GlobalAcceleratorAccessLog"`

	// <p>Number of log task entries.</p>
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeGlobalAcceleratorAccessLogResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGlobalAcceleratorAccessLogResponseParams `json:"Response"`
}

func (r *DescribeGlobalAcceleratorAccessLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGlobalAcceleratorAccessLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGlobalAcceleratorAclPoliciesRequestParams struct {
	// <p>Global acceleration instance ID.</p>
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// <p>Offset. Default value: 0.</p>
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>Number of returns. Default value: 20. Maximum value: 200.</p>
	Limit *string `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeGlobalAcceleratorAclPoliciesRequest struct {
	*tchttp.BaseRequest
	
	// <p>Global acceleration instance ID.</p>
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// <p>Offset. Default value: 0.</p>
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>Number of returns. Default value: 20. Maximum value: 200.</p>
	Limit *string `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeGlobalAcceleratorAclPoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGlobalAcceleratorAclPoliciesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GlobalAcceleratorId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGlobalAcceleratorAclPoliciesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGlobalAcceleratorAclPoliciesResponseParams struct {
	// <p>Access control policy information.</p>
	GlobalAcceleratorAclPolicySet []*GlobalAcceleratorAclPolicies `json:"GlobalAcceleratorAclPolicySet,omitnil,omitempty" name:"GlobalAcceleratorAclPolicySet"`

	// <p>Total number of instances that meet the criteria.</p>
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeGlobalAcceleratorAclPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGlobalAcceleratorAclPoliciesResponseParams `json:"Response"`
}

func (r *DescribeGlobalAcceleratorAclPoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGlobalAcceleratorAclPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGlobalAcceleratorAclRulesRequestParams struct {
	// <p>Access control policy ID.</p>
	GlobalAcceleratorAclPolicyId *string `json:"GlobalAcceleratorAclPolicyId,omitnil,omitempty" name:"GlobalAcceleratorAclPolicyId"`

	// <p>Offset. Default value: 0.</p>
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>Number of returned results.</p><p>Value range: [1, 200]</p><p>Default value: 20</p>
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeGlobalAcceleratorAclRulesRequest struct {
	*tchttp.BaseRequest
	
	// <p>Access control policy ID.</p>
	GlobalAcceleratorAclPolicyId *string `json:"GlobalAcceleratorAclPolicyId,omitnil,omitempty" name:"GlobalAcceleratorAclPolicyId"`

	// <p>Offset. Default value: 0.</p>
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>Number of returned results.</p><p>Value range: [1, 200]</p><p>Default value: 20</p>
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeGlobalAcceleratorAclRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGlobalAcceleratorAclRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GlobalAcceleratorAclPolicyId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGlobalAcceleratorAclRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGlobalAcceleratorAclRulesResponseParams struct {
	// <p>Eligible Acl rule instance.</p>
	GlobalAcceleratorAclRuleSet []*GlobalAcceleratorAclRuleSet `json:"GlobalAcceleratorAclRuleSet,omitnil,omitempty" name:"GlobalAcceleratorAclRuleSet"`

	// <p>Number of instances that meet the criteria.</p>
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeGlobalAcceleratorAclRulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGlobalAcceleratorAclRulesResponseParams `json:"Response"`
}

func (r *DescribeGlobalAcceleratorAclRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGlobalAcceleratorAclRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGlobalAcceleratorsRequestParams struct {
	// <p>Offset. Default value: 0.</p>
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>Number of returned results.</p><p>Value range: [1, 200]</p><p>Default value: 20</p>
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>Filter criteria. <li>global-accelerator-id - String - (Filter condition) Global acceleration instance ID.</li> <li>global-accelerator-state - String - (Filter condition) Global acceleration instance status.</li></p>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeGlobalAcceleratorsRequest struct {
	*tchttp.BaseRequest
	
	// <p>Offset. Default value: 0.</p>
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>Number of returned results.</p><p>Value range: [1, 200]</p><p>Default value: 20</p>
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>Filter criteria. <li>global-accelerator-id - String - (Filter condition) Global acceleration instance ID.</li> <li>global-accelerator-state - String - (Filter condition) Global acceleration instance status.</li></p>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeGlobalAcceleratorsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGlobalAcceleratorsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGlobalAcceleratorsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGlobalAcceleratorsResponseParams struct {
	// <p>Eligible global acceleration instances.</p>
	GlobalAcceleratorSet []*GlobalAcceleratorSet `json:"GlobalAcceleratorSet,omitnil,omitempty" name:"GlobalAcceleratorSet"`

	// <p>Number of instances that meet the criteria.</p>
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeGlobalAcceleratorsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGlobalAcceleratorsResponseParams `json:"Response"`
}

func (r *DescribeGlobalAcceleratorsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGlobalAcceleratorsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeListenersRequestParams struct {
	// Global acceleration instance ID.
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returns. Default value: 20. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Filter criteria. listener-id - String - (Filter criteria) Listener instance ID.
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeListenersRequest struct {
	*tchttp.BaseRequest
	
	// Global acceleration instance ID.
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returns. Default value: 20. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Filter criteria. listener-id - String - (Filter criteria) Listener instance ID.
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
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
	delete(f, "GlobalAcceleratorId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeListenersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeListenersResponseParams struct {
	// Eligible listener instance.
	ListenerSet []*ListenerSet `json:"ListenerSet,omitnil,omitempty" name:"ListenerSet"`

	// Number of instances that meet the criteria.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeTaskResultRequestParams struct {
	// Asynchronous task ID.
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type DescribeTaskResultRequest struct {
	*tchttp.BaseRequest
	
	// Asynchronous task ID.
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTaskResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskResultResponseParams struct {
	// Task status.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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

type EndpointConfigurations struct {
	// <p>Domain type. Available values: 'Domain', 'PublicIp'.</p>
	EndpointType *string `json:"EndpointType,omitnil,omitempty" name:"EndpointType"`

	// <p>Domain name.</p>
	EndpointService *string `json:"EndpointService,omitnil,omitempty" name:"EndpointService"`

	// <p>Weight.</p>
	Weight *uint64 `json:"Weight,omitnil,omitempty" name:"Weight"`

	// <p>Health check status; HEALTH: healthy; UNHEALTH: unhealthy.</p>
	HealthCheckStatus *string `json:"HealthCheckStatus,omitnil,omitempty" name:"HealthCheckStatus"`
}

type EndpointGroupConfiguration struct {
	// <p>Terminal node group name.</p><p>Parameter format: starting with a letter or Chinese characters, 2–128 characters in length, supporting letters, digits, Chinese characters, . - _</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Region of the terminal node group.</p>
	EndpointGroupRegion *string `json:"EndpointGroupRegion,omitnil,omitempty" name:"EndpointGroupRegion"`

	// <p>Terminal node configuration.</p>
	EndpointConfigurations []*EndpointConfigurations `json:"EndpointConfigurations,omitnil,omitempty" name:"EndpointConfigurations"`

	// <p>Check protocol. Supports configuration of 'TCP', 'HTTP', 'PING', and 'CUSTOM'.</p><p>Enumeration values:</p><ul><li>TCP: When the CLB listener protocol where the terminal node group resides is TCP, choose TCP as the check protocol.</li><li>HTTP: When the CLB listener protocol where the terminal node group resides is HTTP or HTTPS, choose HTTP as the check protocol.</li><li>PING: When the CLB listener protocol where the terminal node group resides is UDP, choose PING as the check protocol.</li><li>CUSTOM: When the CLB listener protocol where the terminal node group resides is UDP or TCP, choose CUSTOM as the check protocol.</li></ul><p>This field is required when health check is enabled.</p>
	CheckType *string `json:"CheckType,omitnil,omitempty" name:"CheckType"`

	// <p>Description.</p><p>Default value: empty by default, representing no configuration description.</p><p>Maximum length: 100 bytes.</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>Check port.</p><p>Input limit: range 1-65535.</p><p>This field is required when CheckType is CUSTOM.</p>
	CheckPort *string `json:"CheckPort,omitnil,omitempty" name:"CheckPort"`

	// <p>Check content. Supports configuration 'TEXT'.</p><p>Enumeration values:</p><ul><li>TEXT: Text content.</li></ul><p>This field is required when CheckType is CUSTOM.</p>
	ContextType *string `json:"ContextType,omitnil,omitempty" name:"ContextType"`

	// <p>Check request.</p><p>Input parameter limit: The byte length must be within 1-500.</p><p>This field is required when CheckType is CUSTOM.</p>
	CheckSendContext *string `json:"CheckSendContext,omitnil,omitempty" name:"CheckSendContext"`

	// <p>Check returned results.</p><p>Input parameter limit: The byte length must be within 1-500.</p><p>When CheckType is CUSTOM, this field is required.</p>
	CheckRecvContext *string `json:"CheckRecvContext,omitnil,omitempty" name:"CheckRecvContext"`

	// <p>Whether to enable health check.</p><p>Default value: False</p>
	EnableHealthCheck *bool `json:"EnableHealthCheck,omitnil,omitempty" name:"EnableHealthCheck"`

	// <p>Response timeout.</p><p>Value range: [1, 100]</p><p>Default value: 2</p><p>This field is required when health check is enabled.</p>
	ConnectTimeout *uint64 `json:"ConnectTimeout,omitnil,omitempty" name:"ConnectTimeout"`

	// <p>Health check interval.</p><p>Value range: [5, 300].</p><p>Default value: 30.</p><p>This field is required when health check is enabled.</p>
	HealthCheckInterval *uint64 `json:"HealthCheckInterval,omitnil,omitempty" name:"HealthCheckInterval"`

	// <p>Unhealthy threshold.</p><p>Value range: [1, 10]</p><p>Default value: 3</p><p>This field is required when health check is enabled.</p>
	UnhealthyThreshold *uint64 `json:"UnhealthyThreshold,omitnil,omitempty" name:"UnhealthyThreshold"`

	// <p>Health threshold.</p><p>Value range: [1, 10]</p><p>Default value: 3</p><p>This field is required when health check is enabled.</p>
	HealthyThreshold *uint64 `json:"HealthyThreshold,omitnil,omitempty" name:"HealthyThreshold"`

	// <p>Origin-pull protocol. HTTP and HTTPS can be configured.</p><p>Enumeration values:</p><ul><li>HTTP: HTTP origin-pull. HTTP can be configured when the listener protocol where the terminal node group resides is HTTP or HTTPS.</li><li>HTTPS: HTTPS origin-pull. HTTPS can be configured when the listener protocol where the terminal node group resides is HTTPS.</li></ul><p>This field is required when the listener protocol where the terminal node group resides is HTTP or HTTPS.</p>
	ForwardProtocol *string `json:"ForwardProtocol,omitnil,omitempty" name:"ForwardProtocol"`

	// <p>Check domain name.</p><p>Input parameter limit: The byte length range is 3-80.</p><p>This field is required when CheckType is HTTP.</p>
	CheckDomain *string `json:"CheckDomain,omitnil,omitempty" name:"CheckDomain"`

	// <p>Check the URL.</p><p>Parameter format: must match the regular expression: ^[a-zA-Z0-9_.\-\/]{1,80}$</p><p>This field is required when CheckType is HTTP.</p>
	CheckPath *string `json:"CheckPath,omitnil,omitempty" name:"CheckPath"`

	// <p>Request method. Supports configuration of 'GET' and 'HEAD'.</p><p>Enumeration values:</p><ul><li>GET: The request method is GET.</li><li>HEAD: The request method is HEAD.</li></ul><p>When CheckType is HTTP, this field is required.</p>
	CheckMethod *string `json:"CheckMethod,omitnil,omitempty" name:"CheckMethod"`

	// <p>Status check code. Supports configuring 'http_2xx', 'http_3xx', 'http_4xx', 'http_5xx'.</p><p>Enumeration values:</p><ul><li>http_2xx: HTTP codes beginning with 2.</li><li>http_3xx: HTTP codes beginning with 3.</li><li>http_4xx: HTTP codes beginning with 4.</li><li>http_5xx: HTTP codes beginning with 5.</li></ul><p>This field is required when CheckType is HTTP.</p>
	StatusMask []*string `json:"StatusMask,omitnil,omitempty" name:"StatusMask"`

	// <p>Port mapping.</p><p>Input limits: Layer 7 supports 1 port mapping, and Layer 4 supports up to 30 port mappings.</p>
	PortOverrides []*PortOverride `json:"PortOverrides,omitnil,omitempty" name:"PortOverrides"`

	// <p>Operator type. Supports configuration 'CMCC', 'CTCC', 'CUCC'.</p><p>Enumeration values:</p><ul><li>CMCC: China Mobile</li><li>CUCC: China Unicom</li><li>CTCC: China Telecom</li></ul><p>This field is required when the terminal node group region is a triple-network region.</p>
	IspType *string `json:"IspType,omitnil,omitempty" name:"IspType"`

	// <p>HPPTS encryption algorithm suite; supports configuration 'tls_policy_1.0-2', 'tls_policy_1.1-2', 'tls_policy_1.2', 'tls_policy_1.2_strict', 'tls_policy_1.2_strict-1.3';</p><p>Enumeration values:</p><ul><li>tls_policy_1.0-2: encryption algorithm suite.</li><li>tls_policy_1.1-2: encryption algorithm suite.</li><li>tls_policy_1.2: encryption algorithm suite.</li><li>tls_policy_1.2_strict: encryption algorithm suite.</li><li>tls_policy_1.2_strict-1.3: encryption algorithm suite.</li></ul><p>This field is required when the origin-pull protocol is HTTPS.</p>
	CipherPolicyId *string `json:"CipherPolicyId,omitnil,omitempty" name:"CipherPolicyId"`

	// <p>Origin-pull protocol. Supports configuration of 'HTTP/1.1' and 'HTTP/2'.</p><p>Enumeration values:</p><ul><li>HTTP/1.1: version HTTP/1.1</li><li>HTTP/2: version HTTP/2</li></ul><p>This field is required when the origin-pull protocol is HTTPS.</p>
	HttpVersion *string `json:"HttpVersion,omitnil,omitempty" name:"HttpVersion"`
}

type EndpointGroupConfigurationSet struct {
	// <p>Global acceleration instance ID.</p>
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// <p>Listener instance ID.</p>
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// <p>Terminal node group ID.</p>
	EndpointGroupId *string `json:"EndpointGroupId,omitnil,omitempty" name:"EndpointGroupId"`

	// <p>Name.</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Region.</p>
	EndpointGroupRegion *string `json:"EndpointGroupRegion,omitnil,omitempty" name:"EndpointGroupRegion"`

	// <p>Description.</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>Endpoint information.</p>
	EndpointConfigurations []*EndpointConfigurations `json:"EndpointConfigurations,omitnil,omitempty" name:"EndpointConfigurations"`

	// <p>Whether to enable health check.</p>
	EnableHealthCheck *bool `json:"EnableHealthCheck,omitnil,omitempty" name:"EnableHealthCheck"`

	// <p>Response timeout.</p>
	ConnectTimeout *uint64 `json:"ConnectTimeout,omitnil,omitempty" name:"ConnectTimeout"`

	// <p>Health check interval.</p>
	HealthCheckInterval *uint64 `json:"HealthCheckInterval,omitnil,omitempty" name:"HealthCheckInterval"`

	// <p>Unhealthy threshold.</p>
	UnhealthyThreshold *uint64 `json:"UnhealthyThreshold,omitnil,omitempty" name:"UnhealthyThreshold"`

	// <p>Health threshold.</p>
	HealthyThreshold *uint64 `json:"HealthyThreshold,omitnil,omitempty" name:"HealthyThreshold"`

	// <p>Select the protocol.</p>
	CheckType *string `json:"CheckType,omitnil,omitempty" name:"CheckType"`

	// <p>Check port.</p>
	CheckPort *uint64 `json:"CheckPort,omitnil,omitempty" name:"CheckPort"`

	// <p>Check content.</p>
	ContextType *string `json:"ContextType,omitnil,omitempty" name:"ContextType"`

	// <p>Check request.</p>
	CheckSendContext *string `json:"CheckSendContext,omitnil,omitempty" name:"CheckSendContext"`

	// <p>Check returned results.</p>
	CheckRecvContext *string `json:"CheckRecvContext,omitnil,omitempty" name:"CheckRecvContext"`

	// <p>Check domain name.</p>
	CheckDomain *string `json:"CheckDomain,omitnil,omitempty" name:"CheckDomain"`

	// <p>Check the URL.</p>
	CheckPath *string `json:"CheckPath,omitnil,omitempty" name:"CheckPath"`

	// <p>Request method.</p>
	CheckMethod *string `json:"CheckMethod,omitnil,omitempty" name:"CheckMethod"`

	// <p>Status check code.</p>
	StatusMask []*string `json:"StatusMask,omitnil,omitempty" name:"StatusMask"`

	// <p>Terminal node group type.</p>
	EndpointGroupType *string `json:"EndpointGroupType,omitnil,omitempty" name:"EndpointGroupType"`

	// <p>Origin-pull protocol.</p>
	ForwardProtocol *string `json:"ForwardProtocol,omitnil,omitempty" name:"ForwardProtocol"`

	// <p>Port mapping info.</p>
	PortOverrides []*PortOverride `json:"PortOverrides,omitnil,omitempty" name:"PortOverrides"`

	// <p>Whether the custom endpoint group is bound to a Layer 7 forwarding rule.</p>
	VirtualExistForwardingRuleFlag *bool `json:"VirtualExistForwardingRuleFlag,omitnil,omitempty" name:"VirtualExistForwardingRuleFlag"`

	// <p>Public IP address of the egress terminal node group.</p>
	OriginPublicIps []*string `json:"OriginPublicIps,omitnil,omitempty" name:"OriginPublicIps"`

	// <p>Operator type. China Mobile (CMCC), China Unicom (CUCC), China Telecom (CTCC).</p>
	IspType *string `json:"IspType,omitnil,omitempty" name:"IspType"`

	// <p>HPPTS encryption algorithm kit</p>
	CipherPolicyId *string `json:"CipherPolicyId,omitnil,omitempty" name:"CipherPolicyId"`

	// <p>Only the HTTPS back-to-source protocol supports selecting ['HTTP/1.1', 'HTTP/2']</p><p>Enumeration values:</p><ul><li>HTTP/1.1: Version HTTP/1.1</li><li>HTTP/2: Version HTTP/2</li></ul>
	HttpVersion *string `json:"HttpVersion,omitnil,omitempty" name:"HttpVersion"`
}

type Filter struct {
	// Attribute name. If more than one Filter exists, the logical relation between these Filters is `AND`.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Attribute value. If a filter has multiple values, the relationship among the values under the same filter is logical OR (OR). When the value type is boolean, it can be directly set to the string "TRUE" or "FALSE".
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type ForwardingPolicySet struct {
	// Global acceleration instance ID.
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// Listener ID.
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// Policy ID.
	ForwardingPolicyId *string `json:"ForwardingPolicyId,omitnil,omitempty" name:"ForwardingPolicyId"`

	// Domain name.
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// Whether it is the default domain name.
	DefaultHostFlag *bool `json:"DefaultHostFlag,omitnil,omitempty" name:"DefaultHostFlag"`
}

type ForwardingRuleSet struct {
	// <p>Conditional information of Layer 7 forwarding rules.</p>
	RuleCondition []*RuleCondition `json:"RuleCondition,omitnil,omitempty" name:"RuleCondition"`

	// <p>Behavior information of the Layer 7 forwarding rule.</p>
	RuleAction []*RuleAction `json:"RuleAction,omitnil,omitempty" name:"RuleAction"`

	// <p>Whether to enable origin-pull Sni.</p>
	EnableOriginSni *bool `json:"EnableOriginSni,omitnil,omitempty" name:"EnableOriginSni"`

	// <p>Origin-pull Sni.</p>
	OriginSni *string `json:"OriginSni,omitnil,omitempty" name:"OriginSni"`

	// <p>Origin-pull Header information.</p>
	OriginHeaders []*OriginHeader `json:"OriginHeaders,omitnil,omitempty" name:"OriginHeaders"`

	// <p>Origin-pull Host.</p>
	OriginHost *string `json:"OriginHost,omitnil,omitempty" name:"OriginHost"`

	// <p>Global acceleration instance ID.</p>
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// <p>Listener ID.</p>
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// <p>Layer-7 forwarding policy ID.</p>
	ForwardingPolicyId *string `json:"ForwardingPolicyId,omitnil,omitempty" name:"ForwardingPolicyId"`

	// <p>Layer 7 forwarding rule ID.</p>
	ForwardingRuleId *string `json:"ForwardingRuleId,omitnil,omitempty" name:"ForwardingRuleId"`

	// <p>Origin server response header</p>
	HideResponseHeaders []*HideResponseHeaders `json:"HideResponseHeaders,omitnil,omitempty" name:"HideResponseHeaders"`

	// <p>Delete origin server response headers</p>
	ResponseHeaders []*ResponseHeaders `json:"ResponseHeaders,omitnil,omitempty" name:"ResponseHeaders"`
}

type GlobalAcceleratorAccessLog struct {
	// <p>Log Unique Id</p>
	LogPushTaskId *string `json:"LogPushTaskId,omitnil,omitempty" name:"LogPushTaskId"`

	// <p>Unique Id of the GA instance.</p>
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// <p>Unique Id of the listener</p>
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// <p>Unique Id of the terminal node group</p>
	EndpointGroupId *string `json:"EndpointGroupId,omitnil,omitempty" name:"EndpointGroupId"`

	// <p>Log task description</p>
	FlowLogDescription *string `json:"FlowLogDescription,omitnil,omitempty" name:"FlowLogDescription"`

	// <p>Region where the logs are located.</p>
	CloudRegion *string `json:"CloudRegion,omitnil,omitempty" name:"CloudRegion"`

	// <p>Log topic Id</p>
	CloudLogId *string `json:"CloudLogId,omitnil,omitempty" name:"CloudLogId"`

	// <p>Log Set Id</p>
	CloudLogSetId *string `json:"CloudLogSetId,omitnil,omitempty" name:"CloudLogSetId"`

	// <p>Select log data collection field</p>
	FieldKeys []*string `json:"FieldKeys,omitnil,omitempty" name:"FieldKeys"`

	// <p>Log task status</p><p>Enumeration values:</p><ul><li>active: Running</li><li>stopped: Suspended</li></ul>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

type GlobalAcceleratorAclPolicies struct {
	// Access control policy ID.
	GlobalAcceleratorAclPolicyId *string `json:"GlobalAcceleratorAclPolicyId,omitnil,omitempty" name:"GlobalAcceleratorAclPolicyId"`

	// Default action.
	DefaultAction *string `json:"DefaultAction,omitnil,omitempty" name:"DefaultAction"`

	// Status.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

type GlobalAcceleratorAclRuleSet struct {
	// Access control policy ID.
	GlobalAcceleratorPolicyId *string `json:"GlobalAcceleratorPolicyId,omitnil,omitempty" name:"GlobalAcceleratorPolicyId"`

	// Acl rule ID.
	GlobalAcceleratorAclRuleId *string `json:"GlobalAcceleratorAclRuleId,omitnil,omitempty" name:"GlobalAcceleratorAclRuleId"`

	// Protocol.
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// Port.
	Port *string `json:"Port,omitnil,omitempty" name:"Port"`

	// IP range.
	SourceCidrBlock *string `json:"SourceCidrBlock,omitnil,omitempty" name:"SourceCidrBlock"`

	// Action.
	Policy *string `json:"Policy,omitnil,omitempty" name:"Policy"`

	// Description.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type GlobalAcceleratorSet struct {
	// <p>Global acceleration instance ID.</p>
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// <p>Global acceleration instance name.</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Global acceleration instance description.</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>Instance creation time of the global acceleration instance.</p>
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// <p>Global acceleration instance status.</p>
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// <p>Billing type of a global acceleration instance.</p>
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// <p>DDoS ID of the global acceleration instance.</p>
	DdosId *string `json:"DdosId,omitnil,omitempty" name:"DdosId"`

	// <p>Number of listeners of the associated acceleration instance.</p>
	ListenerCounts *uint64 `json:"ListenerCounts,omitnil,omitempty" name:"ListenerCounts"`

	// <p>Count of acceleration regions belonging to the acceleration instance.</p>
	AcceleratorAreaCounts *uint64 `json:"AcceleratorAreaCounts,omitnil,omitempty" name:"AcceleratorAreaCounts"`

	// <p>Global acceleration instance status.</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>Domain name.</p>
	Cname *string `json:"Cname,omitnil,omitempty" name:"Cname"`

	// <p>Cross-border type; HighQuality (high-quality cross-border), Unicom (China Unicom cross-border), NotAvailable (not enabled).</p>
	CrossBorderType *string `json:"CrossBorderType,omitnil,omitempty" name:"CrossBorderType"`

	// <p>Tag information.</p>
	TagSet []*Tag `json:"TagSet,omitnil,omitempty" name:"TagSet"`
}

type HideResponseHeaders struct {
	// <p>key</p><p>Parameter format: 1. The string only contain printable ASCII characters. 2. Cannot contain these characters ()&lt;&gt;@,;:\&quot;/[ ]?={ }</p><p>Input limit: Length 1-40.</p>
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// <p>value</p><p>Currently only support inputting an empty string ""</p>
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type IpAddressInfoSet struct {
	// <p>IP address.</p>
	IpAddress *string `json:"IpAddress,omitnil,omitempty" name:"IpAddress"`

	// <p>IP type.</p>
	IspType *string `json:"IspType,omitnil,omitempty" name:"IspType"`

	// <p>Ddos type</p>
	DdosProtectionType *string `json:"DdosProtectionType,omitnil,omitempty" name:"DdosProtectionType"`
}

type ListenerSet struct {
	// Global acceleration instance ID.
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// Listener ID.
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// Listener name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Listener description.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Protocol.
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// Port range.
	PortRanges *PortRanges `json:"PortRanges,omitnil,omitempty" name:"PortRanges"`

	// Whether to enable layer-7 access to source IP mode.
	XForwardedForRealIp *bool `json:"XForwardedForRealIp,omitnil,omitempty" name:"XForwardedForRealIp"`

	// Enable session persistence.
	ClientAffinity *string `json:"ClientAffinity,omitnil,omitempty" name:"ClientAffinity"`

	// Session persistence time.
	ClientAffinityTime *uint64 `json:"ClientAffinityTime,omitnil,omitempty" name:"ClientAffinityTime"`

	// SSL decryption method.
	CertificationType *string `json:"CertificationType,omitnil,omitempty" name:"CertificationType"`

	// Server certificate.
	ServerCertificates []*string `json:"ServerCertificates,omitnil,omitempty" name:"ServerCertificates"`

	// Client certificate.
	ClientCaCertificates []*string `json:"ClientCaCertificates,omitnil,omitempty" name:"ClientCaCertificates"`

	// TLS password suite package.
	CipherPolicyId *string `json:"CipherPolicyId,omitnil,omitempty" name:"CipherPolicyId"`

	// HTTP version.
	HttpVersion *string `json:"HttpVersion,omitnil,omitempty" name:"HttpVersion"`

	// Request timeout.
	RequestTimeout *uint64 `json:"RequestTimeout,omitnil,omitempty" name:"RequestTimeout"`

	// Creation time.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Listener routing type.
	ListenerType *string `json:"ListenerType,omitnil,omitempty" name:"ListenerType"`

	// Listener status.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Number of terminal node groups belonging to the listener.
	EndpointGroupCounts *uint64 `json:"EndpointGroupCounts,omitnil,omitempty" name:"EndpointGroupCounts"`

	// Method for obtaining the source IP at Layer 4.
	GetRealIpType *string `json:"GetRealIpType,omitnil,omitempty" name:"GetRealIpType"`

	// Connection timeout.
	IdleTimeout *uint64 `json:"IdleTimeout,omitnil,omitempty" name:"IdleTimeout"`
}

// Predefined struct for user
type ModifyAccelerateAreasRequestParams struct {
	// <p>Global acceleration instance ID.</p>
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// <p>Acceleration region info.</p><p>Input limit: array length cannot exceed 10.</p>
	AcceleratorAreas []*AcceleratorAreas `json:"AcceleratorAreas,omitnil,omitempty" name:"AcceleratorAreas"`
}

type ModifyAccelerateAreasRequest struct {
	*tchttp.BaseRequest
	
	// <p>Global acceleration instance ID.</p>
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// <p>Acceleration region info.</p><p>Input limit: array length cannot exceed 10.</p>
	AcceleratorAreas []*AcceleratorAreas `json:"AcceleratorAreas,omitnil,omitempty" name:"AcceleratorAreas"`
}

func (r *ModifyAccelerateAreasRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccelerateAreasRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GlobalAcceleratorId")
	delete(f, "AcceleratorAreas")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAccelerateAreasRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAccelerateAreasResponseParams struct {
	// <p>Asynchronous task ID.</p>
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAccelerateAreasResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAccelerateAreasResponseParams `json:"Response"`
}

func (r *ModifyAccelerateAreasResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccelerateAreasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAccessLogStatusRequestParams struct {
	// <p>Log Unique Id</p>
	LogPushTaskId *string `json:"LogPushTaskId,omitnil,omitempty" name:"LogPushTaskId"`

	// <p>Status (Start START, Stop STOP)</p><p>Enumeration values:</p><ul><li>START: Start</li><li>STOP: Stop</li></ul>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>Unique Id of the GA instance.</p>
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`
}

type ModifyAccessLogStatusRequest struct {
	*tchttp.BaseRequest
	
	// <p>Log Unique Id</p>
	LogPushTaskId *string `json:"LogPushTaskId,omitnil,omitempty" name:"LogPushTaskId"`

	// <p>Status (Start START, Stop STOP)</p><p>Enumeration values:</p><ul><li>START: Start</li><li>STOP: Stop</li></ul>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>Unique Id of the GA instance.</p>
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`
}

func (r *ModifyAccessLogStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccessLogStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LogPushTaskId")
	delete(f, "Status")
	delete(f, "GlobalAcceleratorId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAccessLogStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAccessLogStatusResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAccessLogStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAccessLogStatusResponseParams `json:"Response"`
}

func (r *ModifyAccessLogStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccessLogStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyEndpointGroupRequestParams struct {
	// <p>Global acceleration instance ID.</p>
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// <p>Listener ID.</p>
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// <p>Terminal node group ID.</p>
	EndpointGroupId *string `json:"EndpointGroupId,omitnil,omitempty" name:"EndpointGroupId"`

	// <p>Terminal node configuration.</p>
	EndpointConfigurations []*EndpointConfigurations `json:"EndpointConfigurations,omitnil,omitempty" name:"EndpointConfigurations"`

	// <p>Name.</p><p>Parameter format: starting with a letter or Chinese characters, 2–128 characters in length, supporting letters, digits, Chinese characters, . - _</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Description.</p><p>Input limit: maximum length cannot exceed 100 bytes.</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>Whether to enable health check.</p>
	EnableHealthCheck *bool `json:"EnableHealthCheck,omitnil,omitempty" name:"EnableHealthCheck"`

	// <p>Response timeout.</p><p>Value range: [1, 100]</p><p>This parameter is required when health check is enabled.</p>
	ConnectTimeout *uint64 `json:"ConnectTimeout,omitnil,omitempty" name:"ConnectTimeout"`

	// <p>Health check interval.</p><p>Value range: [5, 300].</p><p>This parameter is required when health check is enabled.</p>
	HealthCheckInterval *uint64 `json:"HealthCheckInterval,omitnil,omitempty" name:"HealthCheckInterval"`

	// <p>Unhealthy threshold.</p><p>Value range: [1, 10]</p><p>This field is required when health check is enabled.</p>
	UnhealthyThreshold *uint64 `json:"UnhealthyThreshold,omitnil,omitempty" name:"UnhealthyThreshold"`

	// <p>Health threshold.</p><p>Value range: [1, 10]</p><p>This field is required when health check is enabled.</p>
	HealthyThreshold *uint64 `json:"HealthyThreshold,omitnil,omitempty" name:"HealthyThreshold"`

	// <p>Select the protocol.</p><p>Input parameter limits: support filling in: 'TCP', 'HTTP', 'PING', 'CUSTOM'.</p><p>1. When the listener is TCP, you can choose CUSTOM+TCP.<br>2. When the listener is UDP, you can choose PING+CUSTOM.<br>3. When the listener is HTTP or HTTPS, you can choose HTTP.</p>
	CheckType *string `json:"CheckType,omitnil,omitempty" name:"CheckType"`

	// <p>Check port.</p><p>Value range: [1, 65535]</p><p>This field is required when CheckType is CUSTOM.</p>
	CheckPort *uint64 `json:"CheckPort,omitnil,omitempty" name:"CheckPort"`

	// <p>Check content.</p><p>Input parameter limit: Only TEXT is supported.</p><p>This field is required when CheckType is CUSTOM.</p>
	ContextType *string `json:"ContextType,omitnil,omitempty" name:"ContextType"`

	// <p>Check request.</p><p>Input parameter limit: The length range is 1-500.</p><p>This field is required when CheckType is CUSTOM.</p>
	CheckSendContext *string `json:"CheckSendContext,omitnil,omitempty" name:"CheckSendContext"`

	// <p>Check returned results.</p><p>Input parameter limit: length range is 1-500.</p><p>When CheckType is CUSTOM, this field is required.</p>
	CheckRecvContext *string `json:"CheckRecvContext,omitnil,omitempty" name:"CheckRecvContext"`

	// <p>Check domain name.</p><p>Input parameter limit: The length range is 3-80.</p><p>This field is required when CheckType is HTTP.</p>
	CheckDomain *string `json:"CheckDomain,omitnil,omitempty" name:"CheckDomain"`

	// <p>Check the URL.</p><p>Input parameter limit: length range 3-80.</p><p>This field is required when CheckType is HTTP.</p>
	CheckPath *string `json:"CheckPath,omitnil,omitempty" name:"CheckPath"`

	// <p>Request method.</p><p>Input parameter limit: support filling in 'GET', 'HEAD'.</p><p>This field is required when CheckType is HTTP.</p>
	CheckMethod *string `json:"CheckMethod,omitnil,omitempty" name:"CheckMethod"`

	// <p>Status check code.</p><p>Input parameter limits: support selecting 'http_2xx', 'http_3xx', 'http_4xx', 'http_5xx'.</p><p>This field is required when CheckType is HTTP.</p>
	StatusMask []*string `json:"StatusMask,omitnil,omitempty" name:"StatusMask"`

	// <p>Origin-pull protocol.</p><p>Input parameter limits. Supported values: 'HTTP', 'HTTPS'.</p><p>When the CLB listener protocol is HTTP, only HTTP can be configured. When it is HTTPS, HTTP or HTTPS can be configured.</p>
	ForwardProtocol *string `json:"ForwardProtocol,omitnil,omitempty" name:"ForwardProtocol"`

	// <p>Port mapping.</p><p>When the CLB listener protocol is HTTP or HTTPS, one pair can be configured. When the CLB listener protocol is UDP or TCP, up to 30 pairs can be configured.</p>
	PortOverrides []*PortOverride `json:"PortOverrides,omitnil,omitempty" name:"PortOverrides"`

	// <p>HPPTS encryption algorithm kit</p><p>Input parameter limit: support selecting 'tls_policy_1.0-2', 'tls_policy_1.1-2', 'tls_policy_1.2', 'tls_policy_1.2_strict', 'tls_policy_1.2_strict-1.3'.</p><p>This parameter can be modified only when the CLB listener protocol is HTTPS.</p>
	CipherPolicyId *string `json:"CipherPolicyId,omitnil,omitempty" name:"CipherPolicyId"`

	// <p>Only the HTTPS back-to-source protocol supports selecting ['HTTP/1.1', 'HTTP/2']</p><p>Enumeration values:</p><ul><li>HTTP/1.1: version HTTP/1.1</li><li>HTTP/2: version HTTP/2</li></ul>
	HttpVersion *string `json:"HttpVersion,omitnil,omitempty" name:"HttpVersion"`
}

type ModifyEndpointGroupRequest struct {
	*tchttp.BaseRequest
	
	// <p>Global acceleration instance ID.</p>
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// <p>Listener ID.</p>
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// <p>Terminal node group ID.</p>
	EndpointGroupId *string `json:"EndpointGroupId,omitnil,omitempty" name:"EndpointGroupId"`

	// <p>Terminal node configuration.</p>
	EndpointConfigurations []*EndpointConfigurations `json:"EndpointConfigurations,omitnil,omitempty" name:"EndpointConfigurations"`

	// <p>Name.</p><p>Parameter format: starting with a letter or Chinese characters, 2–128 characters in length, supporting letters, digits, Chinese characters, . - _</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Description.</p><p>Input limit: maximum length cannot exceed 100 bytes.</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>Whether to enable health check.</p>
	EnableHealthCheck *bool `json:"EnableHealthCheck,omitnil,omitempty" name:"EnableHealthCheck"`

	// <p>Response timeout.</p><p>Value range: [1, 100]</p><p>This parameter is required when health check is enabled.</p>
	ConnectTimeout *uint64 `json:"ConnectTimeout,omitnil,omitempty" name:"ConnectTimeout"`

	// <p>Health check interval.</p><p>Value range: [5, 300].</p><p>This parameter is required when health check is enabled.</p>
	HealthCheckInterval *uint64 `json:"HealthCheckInterval,omitnil,omitempty" name:"HealthCheckInterval"`

	// <p>Unhealthy threshold.</p><p>Value range: [1, 10]</p><p>This field is required when health check is enabled.</p>
	UnhealthyThreshold *uint64 `json:"UnhealthyThreshold,omitnil,omitempty" name:"UnhealthyThreshold"`

	// <p>Health threshold.</p><p>Value range: [1, 10]</p><p>This field is required when health check is enabled.</p>
	HealthyThreshold *uint64 `json:"HealthyThreshold,omitnil,omitempty" name:"HealthyThreshold"`

	// <p>Select the protocol.</p><p>Input parameter limits: support filling in: 'TCP', 'HTTP', 'PING', 'CUSTOM'.</p><p>1. When the listener is TCP, you can choose CUSTOM+TCP.<br>2. When the listener is UDP, you can choose PING+CUSTOM.<br>3. When the listener is HTTP or HTTPS, you can choose HTTP.</p>
	CheckType *string `json:"CheckType,omitnil,omitempty" name:"CheckType"`

	// <p>Check port.</p><p>Value range: [1, 65535]</p><p>This field is required when CheckType is CUSTOM.</p>
	CheckPort *uint64 `json:"CheckPort,omitnil,omitempty" name:"CheckPort"`

	// <p>Check content.</p><p>Input parameter limit: Only TEXT is supported.</p><p>This field is required when CheckType is CUSTOM.</p>
	ContextType *string `json:"ContextType,omitnil,omitempty" name:"ContextType"`

	// <p>Check request.</p><p>Input parameter limit: The length range is 1-500.</p><p>This field is required when CheckType is CUSTOM.</p>
	CheckSendContext *string `json:"CheckSendContext,omitnil,omitempty" name:"CheckSendContext"`

	// <p>Check returned results.</p><p>Input parameter limit: length range is 1-500.</p><p>When CheckType is CUSTOM, this field is required.</p>
	CheckRecvContext *string `json:"CheckRecvContext,omitnil,omitempty" name:"CheckRecvContext"`

	// <p>Check domain name.</p><p>Input parameter limit: The length range is 3-80.</p><p>This field is required when CheckType is HTTP.</p>
	CheckDomain *string `json:"CheckDomain,omitnil,omitempty" name:"CheckDomain"`

	// <p>Check the URL.</p><p>Input parameter limit: length range 3-80.</p><p>This field is required when CheckType is HTTP.</p>
	CheckPath *string `json:"CheckPath,omitnil,omitempty" name:"CheckPath"`

	// <p>Request method.</p><p>Input parameter limit: support filling in 'GET', 'HEAD'.</p><p>This field is required when CheckType is HTTP.</p>
	CheckMethod *string `json:"CheckMethod,omitnil,omitempty" name:"CheckMethod"`

	// <p>Status check code.</p><p>Input parameter limits: support selecting 'http_2xx', 'http_3xx', 'http_4xx', 'http_5xx'.</p><p>This field is required when CheckType is HTTP.</p>
	StatusMask []*string `json:"StatusMask,omitnil,omitempty" name:"StatusMask"`

	// <p>Origin-pull protocol.</p><p>Input parameter limits. Supported values: 'HTTP', 'HTTPS'.</p><p>When the CLB listener protocol is HTTP, only HTTP can be configured. When it is HTTPS, HTTP or HTTPS can be configured.</p>
	ForwardProtocol *string `json:"ForwardProtocol,omitnil,omitempty" name:"ForwardProtocol"`

	// <p>Port mapping.</p><p>When the CLB listener protocol is HTTP or HTTPS, one pair can be configured. When the CLB listener protocol is UDP or TCP, up to 30 pairs can be configured.</p>
	PortOverrides []*PortOverride `json:"PortOverrides,omitnil,omitempty" name:"PortOverrides"`

	// <p>HPPTS encryption algorithm kit</p><p>Input parameter limit: support selecting 'tls_policy_1.0-2', 'tls_policy_1.1-2', 'tls_policy_1.2', 'tls_policy_1.2_strict', 'tls_policy_1.2_strict-1.3'.</p><p>This parameter can be modified only when the CLB listener protocol is HTTPS.</p>
	CipherPolicyId *string `json:"CipherPolicyId,omitnil,omitempty" name:"CipherPolicyId"`

	// <p>Only the HTTPS back-to-source protocol supports selecting ['HTTP/1.1', 'HTTP/2']</p><p>Enumeration values:</p><ul><li>HTTP/1.1: version HTTP/1.1</li><li>HTTP/2: version HTTP/2</li></ul>
	HttpVersion *string `json:"HttpVersion,omitnil,omitempty" name:"HttpVersion"`
}

func (r *ModifyEndpointGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyEndpointGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GlobalAcceleratorId")
	delete(f, "ListenerId")
	delete(f, "EndpointGroupId")
	delete(f, "EndpointConfigurations")
	delete(f, "Name")
	delete(f, "Description")
	delete(f, "EnableHealthCheck")
	delete(f, "ConnectTimeout")
	delete(f, "HealthCheckInterval")
	delete(f, "UnhealthyThreshold")
	delete(f, "HealthyThreshold")
	delete(f, "CheckType")
	delete(f, "CheckPort")
	delete(f, "ContextType")
	delete(f, "CheckSendContext")
	delete(f, "CheckRecvContext")
	delete(f, "CheckDomain")
	delete(f, "CheckPath")
	delete(f, "CheckMethod")
	delete(f, "StatusMask")
	delete(f, "ForwardProtocol")
	delete(f, "PortOverrides")
	delete(f, "CipherPolicyId")
	delete(f, "HttpVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyEndpointGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyEndpointGroupResponseParams struct {
	// <p>Task ID.</p>
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyEndpointGroupResponse struct {
	*tchttp.BaseResponse
	Response *ModifyEndpointGroupResponseParams `json:"Response"`
}

func (r *ModifyEndpointGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyEndpointGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyForwardingPolicyRequestParams struct {
	// <p>Global acceleration instance ID.</p>
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// <p>Listener ID.</p>
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// <p>Policy ID.</p>
	ForwardingPolicyId *string `json:"ForwardingPolicyId,omitnil,omitempty" name:"ForwardingPolicyId"`

	// <p>Domain name.</p><p>Input limit: length range is 1-80.</p><p>The format must meet the regular expression: ^(<a href="?:[a-z0-9-]{0,61}[a-z0-9]">a-z0-9</a>?.)+[a-z]{2,}$</p>
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`
}

type ModifyForwardingPolicyRequest struct {
	*tchttp.BaseRequest
	
	// <p>Global acceleration instance ID.</p>
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// <p>Listener ID.</p>
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// <p>Policy ID.</p>
	ForwardingPolicyId *string `json:"ForwardingPolicyId,omitnil,omitempty" name:"ForwardingPolicyId"`

	// <p>Domain name.</p><p>Input limit: length range is 1-80.</p><p>The format must meet the regular expression: ^(<a href="?:[a-z0-9-]{0,61}[a-z0-9]">a-z0-9</a>?.)+[a-z]{2,}$</p>
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`
}

func (r *ModifyForwardingPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyForwardingPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GlobalAcceleratorId")
	delete(f, "ListenerId")
	delete(f, "ForwardingPolicyId")
	delete(f, "Host")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyForwardingPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyForwardingPolicyResponseParams struct {
	// <p>Asynchronous task ID.</p>
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyForwardingPolicyResponse struct {
	*tchttp.BaseResponse
	Response *ModifyForwardingPolicyResponseParams `json:"Response"`
}

func (r *ModifyForwardingPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyForwardingPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyForwardingRuleRequestParams struct {
	// <p>Global acceleration instance ID.</p>
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// <p>Listener ID.</p>
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// <p>Policy ID.</p>
	ForwardingPolicyId *string `json:"ForwardingPolicyId,omitnil,omitempty" name:"ForwardingPolicyId"`

	// <p>Layer 7 forwarding rule ID.</p>
	ForwardingRuleId *string `json:"ForwardingRuleId,omitnil,omitempty" name:"ForwardingRuleId"`

	// <p>Conditional information of Layer 7 forwarding rules.</p><p>Input parameter limit: The array length cannot exceed 1.</p>
	RuleConditions []*RuleCondition `json:"RuleConditions,omitnil,omitempty" name:"RuleConditions"`

	// <p>Layer 7 forwarding rule behavior information.</p><p>Input parameter limit: array length cannot exceed 1.</p>
	RuleActions []*RuleAction `json:"RuleActions,omitnil,omitempty" name:"RuleActions"`

	// <p>Origin-pull Header information.</p><p>Input limitation: The length of the array is between 1 and 5.</p>
	OriginHeaders []*OriginHeader `json:"OriginHeaders,omitnil,omitempty" name:"OriginHeaders"`

	// <p>Whether to enable origin-pull sni.</p>
	EnableOriginSni *bool `json:"EnableOriginSni,omitnil,omitempty" name:"EnableOriginSni"`

	// <p>Origin sni.</p><p>Input parameter limit: length cannot exceed 80.</p><p>This field is required when origin sni is enabled.</p>
	OriginSni *string `json:"OriginSni,omitnil,omitempty" name:"OriginSni"`

	// <p>Origin-pull host.</p><p>Input parameter limit: length cannot exceed 80.</p><p>This field is required when origin-pull sni is enabled.</p>
	OriginHost *string `json:"OriginHost,omitnil,omitempty" name:"OriginHost"`

	// <p>Origin server response headers</p><p>Input limitation: The array length cannot exceed 5.</p>
	ResponseHeaders []*ResponseHeaders `json:"ResponseHeaders,omitnil,omitempty" name:"ResponseHeaders"`

	// <p>Delete origin response headers</p><p>Input parameter limit: array length cannot exceed 5.</p>
	HideResponseHeaders []*HideResponseHeaders `json:"HideResponseHeaders,omitnil,omitempty" name:"HideResponseHeaders"`
}

type ModifyForwardingRuleRequest struct {
	*tchttp.BaseRequest
	
	// <p>Global acceleration instance ID.</p>
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// <p>Listener ID.</p>
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// <p>Policy ID.</p>
	ForwardingPolicyId *string `json:"ForwardingPolicyId,omitnil,omitempty" name:"ForwardingPolicyId"`

	// <p>Layer 7 forwarding rule ID.</p>
	ForwardingRuleId *string `json:"ForwardingRuleId,omitnil,omitempty" name:"ForwardingRuleId"`

	// <p>Conditional information of Layer 7 forwarding rules.</p><p>Input parameter limit: The array length cannot exceed 1.</p>
	RuleConditions []*RuleCondition `json:"RuleConditions,omitnil,omitempty" name:"RuleConditions"`

	// <p>Layer 7 forwarding rule behavior information.</p><p>Input parameter limit: array length cannot exceed 1.</p>
	RuleActions []*RuleAction `json:"RuleActions,omitnil,omitempty" name:"RuleActions"`

	// <p>Origin-pull Header information.</p><p>Input limitation: The length of the array is between 1 and 5.</p>
	OriginHeaders []*OriginHeader `json:"OriginHeaders,omitnil,omitempty" name:"OriginHeaders"`

	// <p>Whether to enable origin-pull sni.</p>
	EnableOriginSni *bool `json:"EnableOriginSni,omitnil,omitempty" name:"EnableOriginSni"`

	// <p>Origin sni.</p><p>Input parameter limit: length cannot exceed 80.</p><p>This field is required when origin sni is enabled.</p>
	OriginSni *string `json:"OriginSni,omitnil,omitempty" name:"OriginSni"`

	// <p>Origin-pull host.</p><p>Input parameter limit: length cannot exceed 80.</p><p>This field is required when origin-pull sni is enabled.</p>
	OriginHost *string `json:"OriginHost,omitnil,omitempty" name:"OriginHost"`

	// <p>Origin server response headers</p><p>Input limitation: The array length cannot exceed 5.</p>
	ResponseHeaders []*ResponseHeaders `json:"ResponseHeaders,omitnil,omitempty" name:"ResponseHeaders"`

	// <p>Delete origin response headers</p><p>Input parameter limit: array length cannot exceed 5.</p>
	HideResponseHeaders []*HideResponseHeaders `json:"HideResponseHeaders,omitnil,omitempty" name:"HideResponseHeaders"`
}

func (r *ModifyForwardingRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyForwardingRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GlobalAcceleratorId")
	delete(f, "ListenerId")
	delete(f, "ForwardingPolicyId")
	delete(f, "ForwardingRuleId")
	delete(f, "RuleConditions")
	delete(f, "RuleActions")
	delete(f, "OriginHeaders")
	delete(f, "EnableOriginSni")
	delete(f, "OriginSni")
	delete(f, "OriginHost")
	delete(f, "ResponseHeaders")
	delete(f, "HideResponseHeaders")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyForwardingRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyForwardingRuleResponseParams struct {
	// <p>Asynchronous task ID.</p>
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyForwardingRuleResponse struct {
	*tchttp.BaseResponse
	Response *ModifyForwardingRuleResponseParams `json:"Response"`
}

func (r *ModifyForwardingRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyForwardingRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyGlobalAcceleratorAccessLogRequestParams struct {
	// <p>Unique Id of the log</p>
	LogPushTaskId *string `json:"LogPushTaskId,omitnil,omitempty" name:"LogPushTaskId"`

	// <p>Unique Id of a GA instance.</p>
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// <p>Log topic Id</p>
	CloudLogId *string `json:"CloudLogId,omitnil,omitempty" name:"CloudLogId"`

	// <p>Log Set Id</p>
	CloudLogSetId *string `json:"CloudLogSetId,omitnil,omitempty" name:"CloudLogSetId"`

	// <p>user-selectable log listening fields</p><p>Enumeration values:</p><ul><li>session_time: Layer 4, session duration</li><li>upstream_bytes_received: Layer 4 and Layer 7, number of bytes received from the terminal node</li><li>upstream_bytes_sent: Layer 4 and Layer 7, number of bytes sent to the terminal node</li><li>request_method: Layer 7, GET/POST</li><li>scheme: Layer 7, http/https</li><li>request_uri: Layer 7, uri of the client's original request</li><li>uri: Layer 7, uri of the current request</li><li>host: Layer 7, domain name accessed by the client (Layer 7)</li><li>remote_user: Layer 7, username for basic authentication ("-" if unauthenticated)</li><li>http_user_agent: Layer 7, client browser identification</li><li>http_referer: Layer 7, request source URL ("-" when accessed directly from the address bar)</li><li>http_x_forwarded_for: Layer 7, records the client's original IP and the chain of proxy server IPs it transited</li><li>content_type: Layer 7, content_type</li><li>body_bytes_sent: Layer 7, http body size sent to the client, excluding the header</li><li>request_time: Layer 7, total time from receiving the first byte of the client request to sending the last byte of the response (unit: seconds)</li><li>sent_http_content_type: Layer 7, response content type</li><li>upstream_header_time: Layer 7, arrival time of the response header from the terminal node</li><li>upstream_response_length: Layer 7, response body length returned by the terminal node</li><li>upstream_response_time: Layer 7, full response time of the terminal node</li><li>upstream_status: Layer 7, http status code returned by the terminal node</li></ul>
	FieldKeys []*string `json:"FieldKeys,omitnil,omitempty" name:"FieldKeys"`

	// <p>Log description</p>
	FlowLogDescription *string `json:"FlowLogDescription,omitnil,omitempty" name:"FlowLogDescription"`
}

type ModifyGlobalAcceleratorAccessLogRequest struct {
	*tchttp.BaseRequest
	
	// <p>Unique Id of the log</p>
	LogPushTaskId *string `json:"LogPushTaskId,omitnil,omitempty" name:"LogPushTaskId"`

	// <p>Unique Id of a GA instance.</p>
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// <p>Log topic Id</p>
	CloudLogId *string `json:"CloudLogId,omitnil,omitempty" name:"CloudLogId"`

	// <p>Log Set Id</p>
	CloudLogSetId *string `json:"CloudLogSetId,omitnil,omitempty" name:"CloudLogSetId"`

	// <p>user-selectable log listening fields</p><p>Enumeration values:</p><ul><li>session_time: Layer 4, session duration</li><li>upstream_bytes_received: Layer 4 and Layer 7, number of bytes received from the terminal node</li><li>upstream_bytes_sent: Layer 4 and Layer 7, number of bytes sent to the terminal node</li><li>request_method: Layer 7, GET/POST</li><li>scheme: Layer 7, http/https</li><li>request_uri: Layer 7, uri of the client's original request</li><li>uri: Layer 7, uri of the current request</li><li>host: Layer 7, domain name accessed by the client (Layer 7)</li><li>remote_user: Layer 7, username for basic authentication ("-" if unauthenticated)</li><li>http_user_agent: Layer 7, client browser identification</li><li>http_referer: Layer 7, request source URL ("-" when accessed directly from the address bar)</li><li>http_x_forwarded_for: Layer 7, records the client's original IP and the chain of proxy server IPs it transited</li><li>content_type: Layer 7, content_type</li><li>body_bytes_sent: Layer 7, http body size sent to the client, excluding the header</li><li>request_time: Layer 7, total time from receiving the first byte of the client request to sending the last byte of the response (unit: seconds)</li><li>sent_http_content_type: Layer 7, response content type</li><li>upstream_header_time: Layer 7, arrival time of the response header from the terminal node</li><li>upstream_response_length: Layer 7, response body length returned by the terminal node</li><li>upstream_response_time: Layer 7, full response time of the terminal node</li><li>upstream_status: Layer 7, http status code returned by the terminal node</li></ul>
	FieldKeys []*string `json:"FieldKeys,omitnil,omitempty" name:"FieldKeys"`

	// <p>Log description</p>
	FlowLogDescription *string `json:"FlowLogDescription,omitnil,omitempty" name:"FlowLogDescription"`
}

func (r *ModifyGlobalAcceleratorAccessLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyGlobalAcceleratorAccessLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LogPushTaskId")
	delete(f, "GlobalAcceleratorId")
	delete(f, "CloudLogId")
	delete(f, "CloudLogSetId")
	delete(f, "FieldKeys")
	delete(f, "FlowLogDescription")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyGlobalAcceleratorAccessLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyGlobalAcceleratorAccessLogResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyGlobalAcceleratorAccessLogResponse struct {
	*tchttp.BaseResponse
	Response *ModifyGlobalAcceleratorAccessLogResponseParams `json:"Response"`
}

func (r *ModifyGlobalAcceleratorAccessLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyGlobalAcceleratorAccessLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyGlobalAcceleratorAclPolicyRequestParams struct {
	// <p>Global acceleration instance ID.</p>
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// <p>Access control policy ID.</p>
	GlobalAcceleratorAclPolicyId *string `json:"GlobalAcceleratorAclPolicyId,omitnil,omitempty" name:"GlobalAcceleratorAclPolicyId"`

	// <p>Access control policy status.</p><p>Enumeration values:</p><ul><li>OPEN: On.</li><li>CLOSE: Off.</li></ul>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

type ModifyGlobalAcceleratorAclPolicyRequest struct {
	*tchttp.BaseRequest
	
	// <p>Global acceleration instance ID.</p>
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// <p>Access control policy ID.</p>
	GlobalAcceleratorAclPolicyId *string `json:"GlobalAcceleratorAclPolicyId,omitnil,omitempty" name:"GlobalAcceleratorAclPolicyId"`

	// <p>Access control policy status.</p><p>Enumeration values:</p><ul><li>OPEN: On.</li><li>CLOSE: Off.</li></ul>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

func (r *ModifyGlobalAcceleratorAclPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyGlobalAcceleratorAclPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GlobalAcceleratorId")
	delete(f, "GlobalAcceleratorAclPolicyId")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyGlobalAcceleratorAclPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyGlobalAcceleratorAclPolicyResponseParams struct {
	// <p>Asynchronous task ID.</p>
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyGlobalAcceleratorAclPolicyResponse struct {
	*tchttp.BaseResponse
	Response *ModifyGlobalAcceleratorAclPolicyResponseParams `json:"Response"`
}

func (r *ModifyGlobalAcceleratorAclPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyGlobalAcceleratorAclPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyGlobalAcceleratorAclRuleRequestParams struct {
	// <p>Global acceleration instance ID.</p>
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// <p>Security policy ID</p>
	GlobalAcceleratorAclPolicyId *string `json:"GlobalAcceleratorAclPolicyId,omitnil,omitempty" name:"GlobalAcceleratorAclPolicyId"`

	// <p>Acl rule ID.</p>
	GlobalAcceleratorAclRuleId *string `json:"GlobalAcceleratorAclRuleId,omitnil,omitempty" name:"GlobalAcceleratorAclRuleId"`

	// <p>Protocol.</p><p>Input parameter limit: support selecting 'TCP', 'UDP'.</p>
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// <p>Port.</p>
	Port *string `json:"Port,omitnil,omitempty" name:"Port"`

	// <p>IP range.</p>
	SourceCidrBlock *string `json:"SourceCidrBlock,omitnil,omitempty" name:"SourceCidrBlock"`

	// <p>Action.</p><p>Input parameter limit: support selecting 'ACCEPT', 'DROP'.</p><p>Enumeration values:</p><ul><li>ACCEPT: permission.</li><li>DROP: deny.</li></ul>
	Policy *string `json:"Policy,omitnil,omitempty" name:"Policy"`

	// <p>Description. Maximum length cannot exceed 100 bytes.</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type ModifyGlobalAcceleratorAclRuleRequest struct {
	*tchttp.BaseRequest
	
	// <p>Global acceleration instance ID.</p>
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// <p>Security policy ID</p>
	GlobalAcceleratorAclPolicyId *string `json:"GlobalAcceleratorAclPolicyId,omitnil,omitempty" name:"GlobalAcceleratorAclPolicyId"`

	// <p>Acl rule ID.</p>
	GlobalAcceleratorAclRuleId *string `json:"GlobalAcceleratorAclRuleId,omitnil,omitempty" name:"GlobalAcceleratorAclRuleId"`

	// <p>Protocol.</p><p>Input parameter limit: support selecting 'TCP', 'UDP'.</p>
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// <p>Port.</p>
	Port *string `json:"Port,omitnil,omitempty" name:"Port"`

	// <p>IP range.</p>
	SourceCidrBlock *string `json:"SourceCidrBlock,omitnil,omitempty" name:"SourceCidrBlock"`

	// <p>Action.</p><p>Input parameter limit: support selecting 'ACCEPT', 'DROP'.</p><p>Enumeration values:</p><ul><li>ACCEPT: permission.</li><li>DROP: deny.</li></ul>
	Policy *string `json:"Policy,omitnil,omitempty" name:"Policy"`

	// <p>Description. Maximum length cannot exceed 100 bytes.</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *ModifyGlobalAcceleratorAclRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyGlobalAcceleratorAclRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GlobalAcceleratorId")
	delete(f, "GlobalAcceleratorAclPolicyId")
	delete(f, "GlobalAcceleratorAclRuleId")
	delete(f, "Protocol")
	delete(f, "Port")
	delete(f, "SourceCidrBlock")
	delete(f, "Policy")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyGlobalAcceleratorAclRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyGlobalAcceleratorAclRuleResponseParams struct {
	// <p>Asynchronous task ID.</p>
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyGlobalAcceleratorAclRuleResponse struct {
	*tchttp.BaseResponse
	Response *ModifyGlobalAcceleratorAclRuleResponseParams `json:"Response"`
}

func (r *ModifyGlobalAcceleratorAclRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyGlobalAcceleratorAclRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyGlobalAcceleratorRequestParams struct {
	// <p>Global acceleration instance ID.</p>
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// <p>Name.</p><p>Parameter format: starting with a letter or Chinese characters, 2–128 characters in length, supporting letters, digits, Chinese characters, . - _</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Description.</p><p>Parameter format: should not exceed 100 characters.</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>Cross-border type.</p><p>Enumeration values:</p><ul><li>HighQuality: high-quality cross-border.</li><li>Unicom: China Unicom cross-border.</li></ul>
	CrossBorderType *string `json:"CrossBorderType,omitnil,omitempty" name:"CrossBorderType"`

	// <p>Indicates whether to complete the cross-border service commitment.</p><p>When CrossBorderType is passed in, this field must be set to true, indicating the cross-border commitment is completed.</p>
	CrossBorderPromiseFlag *bool `json:"CrossBorderPromiseFlag,omitnil,omitempty" name:"CrossBorderPromiseFlag"`
}

type ModifyGlobalAcceleratorRequest struct {
	*tchttp.BaseRequest
	
	// <p>Global acceleration instance ID.</p>
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// <p>Name.</p><p>Parameter format: starting with a letter or Chinese characters, 2–128 characters in length, supporting letters, digits, Chinese characters, . - _</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Description.</p><p>Parameter format: should not exceed 100 characters.</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>Cross-border type.</p><p>Enumeration values:</p><ul><li>HighQuality: high-quality cross-border.</li><li>Unicom: China Unicom cross-border.</li></ul>
	CrossBorderType *string `json:"CrossBorderType,omitnil,omitempty" name:"CrossBorderType"`

	// <p>Indicates whether to complete the cross-border service commitment.</p><p>When CrossBorderType is passed in, this field must be set to true, indicating the cross-border commitment is completed.</p>
	CrossBorderPromiseFlag *bool `json:"CrossBorderPromiseFlag,omitnil,omitempty" name:"CrossBorderPromiseFlag"`
}

func (r *ModifyGlobalAcceleratorRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyGlobalAcceleratorRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GlobalAcceleratorId")
	delete(f, "Name")
	delete(f, "Description")
	delete(f, "CrossBorderType")
	delete(f, "CrossBorderPromiseFlag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyGlobalAcceleratorRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyGlobalAcceleratorResponseParams struct {
	// <p>Asynchronous task ID.</p>
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyGlobalAcceleratorResponse struct {
	*tchttp.BaseResponse
	Response *ModifyGlobalAcceleratorResponseParams `json:"Response"`
}

func (r *ModifyGlobalAcceleratorResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyGlobalAcceleratorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyListenerRequestParams struct {
	// <p>Global acceleration instance ID.</p>
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// <p>Listener ID.</p>
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// <p>Name.</p><p>Parameter format: starting with a letter or Chinese characters, 2–128 characters in length, supporting letters, digits, Chinese characters, . - _</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Description. Maximum length cannot exceed 100 bytes.</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>Connection idle wait time.</p><p>1. For HTTP/HTTPS listener, the supported range is 1-60. 2. For TCP listener, the supported range is 10-900. 3. For UDP listener, the supported range is 10-20.</p>
	IdleTimeout *uint64 `json:"IdleTimeout,omitnil,omitempty" name:"IdleTimeout"`

	// <p>Whether to enable session persistence.</p><p>Enumeration values:</p><ul><li>Open: on.</li><li>Close: off.</li></ul><p>TCP/UDP listeners support modification of this parameter.</p>
	ClientAffinity *string `json:"ClientAffinity,omitnil,omitempty" name:"ClientAffinity"`

	// <p>Session persistence duration.</p><p>Value range: [60, 3600]</p>
	ClientAffinityTime *uint64 `json:"ClientAffinityTime,omitnil,omitempty" name:"ClientAffinityTime"`

	// <p>Request timeout.</p><p>Value range: [1, 180]</p><p>This parameter can be modified only for HTTPS listeners.</p>
	RequestTimeout *uint64 `json:"RequestTimeout,omitnil,omitempty" name:"RequestTimeout"`

	// <p>Whether to enable the layer 7 method of obtaining the client IP.</p><p>This parameter modification is supported only for HTTPS/HTTP listeners.</p>
	XForwardedForRealIp *bool `json:"XForwardedForRealIp,omitnil,omitempty" name:"XForwardedForRealIp"`

	// <p>Parsing method.</p><p>Enumeration values:</p><ul><li>UNIDIRECTIONAL: two-way.</li><li>MUTUAL: one-way.</li></ul><p>Only HTTPS/HTTP listeners support modifying this parameter.</p>
	CertificationType *string `json:"CertificationType,omitnil,omitempty" name:"CertificationType"`

	// <p>Encryption algorithm kit.</p><p>Input limits: support selecting tls_policy_1.0-2', 'tls_policy_1.1-2', 'tls_policy_1.2', 'tls_policy_1.2_strict', 'tls_policy_1.2_strict-1.3'.</p><p>Only HTTPS listeners support modifying this parameter.</p>
	CipherPolicyId *string `json:"CipherPolicyId,omitnil,omitempty" name:"CipherPolicyId"`

	// <p>Server certificate.</p><p>Input limit: currently only support importing one cert; to use multiple certs, use the cert api CreateListenerAdditionalCert to add other certs.</p><p>Only HTTPS listeners support modification of this parameter.</p>
	ServerCertificates []*string `json:"ServerCertificates,omitnil,omitempty" name:"ServerCertificates"`

	// <p>Client certificate.</p><p>Input limitations: 1. Currently only support importing one certificate; to use multiple certificates, use the certificate api CreateListenerAdditionalCert to add other certificates. 2. The certificate must be a CA certificate.</p><p>Only HTTPS listeners support modification of this parameter, and mutual authentication must be enabled.</p>
	ClientCaCertificates []*string `json:"ClientCaCertificates,omitnil,omitempty" name:"ClientCaCertificates"`

	// <p>Method of obtaining the source IP.</p><p>Input parameter limits: support selecting 'ProxyProtocol', 'Close', 'ProxyProtocolV2', 'TOA'.</p><p>Only TCP listeners support modification of this parameter.</p>
	GetRealIpType *string `json:"GetRealIpType,omitnil,omitempty" name:"GetRealIpType"`
}

type ModifyListenerRequest struct {
	*tchttp.BaseRequest
	
	// <p>Global acceleration instance ID.</p>
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// <p>Listener ID.</p>
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// <p>Name.</p><p>Parameter format: starting with a letter or Chinese characters, 2–128 characters in length, supporting letters, digits, Chinese characters, . - _</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Description. Maximum length cannot exceed 100 bytes.</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>Connection idle wait time.</p><p>1. For HTTP/HTTPS listener, the supported range is 1-60. 2. For TCP listener, the supported range is 10-900. 3. For UDP listener, the supported range is 10-20.</p>
	IdleTimeout *uint64 `json:"IdleTimeout,omitnil,omitempty" name:"IdleTimeout"`

	// <p>Whether to enable session persistence.</p><p>Enumeration values:</p><ul><li>Open: on.</li><li>Close: off.</li></ul><p>TCP/UDP listeners support modification of this parameter.</p>
	ClientAffinity *string `json:"ClientAffinity,omitnil,omitempty" name:"ClientAffinity"`

	// <p>Session persistence duration.</p><p>Value range: [60, 3600]</p>
	ClientAffinityTime *uint64 `json:"ClientAffinityTime,omitnil,omitempty" name:"ClientAffinityTime"`

	// <p>Request timeout.</p><p>Value range: [1, 180]</p><p>This parameter can be modified only for HTTPS listeners.</p>
	RequestTimeout *uint64 `json:"RequestTimeout,omitnil,omitempty" name:"RequestTimeout"`

	// <p>Whether to enable the layer 7 method of obtaining the client IP.</p><p>This parameter modification is supported only for HTTPS/HTTP listeners.</p>
	XForwardedForRealIp *bool `json:"XForwardedForRealIp,omitnil,omitempty" name:"XForwardedForRealIp"`

	// <p>Parsing method.</p><p>Enumeration values:</p><ul><li>UNIDIRECTIONAL: two-way.</li><li>MUTUAL: one-way.</li></ul><p>Only HTTPS/HTTP listeners support modifying this parameter.</p>
	CertificationType *string `json:"CertificationType,omitnil,omitempty" name:"CertificationType"`

	// <p>Encryption algorithm kit.</p><p>Input limits: support selecting tls_policy_1.0-2', 'tls_policy_1.1-2', 'tls_policy_1.2', 'tls_policy_1.2_strict', 'tls_policy_1.2_strict-1.3'.</p><p>Only HTTPS listeners support modifying this parameter.</p>
	CipherPolicyId *string `json:"CipherPolicyId,omitnil,omitempty" name:"CipherPolicyId"`

	// <p>Server certificate.</p><p>Input limit: currently only support importing one cert; to use multiple certs, use the cert api CreateListenerAdditionalCert to add other certs.</p><p>Only HTTPS listeners support modification of this parameter.</p>
	ServerCertificates []*string `json:"ServerCertificates,omitnil,omitempty" name:"ServerCertificates"`

	// <p>Client certificate.</p><p>Input limitations: 1. Currently only support importing one certificate; to use multiple certificates, use the certificate api CreateListenerAdditionalCert to add other certificates. 2. The certificate must be a CA certificate.</p><p>Only HTTPS listeners support modification of this parameter, and mutual authentication must be enabled.</p>
	ClientCaCertificates []*string `json:"ClientCaCertificates,omitnil,omitempty" name:"ClientCaCertificates"`

	// <p>Method of obtaining the source IP.</p><p>Input parameter limits: support selecting 'ProxyProtocol', 'Close', 'ProxyProtocolV2', 'TOA'.</p><p>Only TCP listeners support modification of this parameter.</p>
	GetRealIpType *string `json:"GetRealIpType,omitnil,omitempty" name:"GetRealIpType"`
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
	delete(f, "GlobalAcceleratorId")
	delete(f, "ListenerId")
	delete(f, "Name")
	delete(f, "Description")
	delete(f, "IdleTimeout")
	delete(f, "ClientAffinity")
	delete(f, "ClientAffinityTime")
	delete(f, "RequestTimeout")
	delete(f, "XForwardedForRealIp")
	delete(f, "CertificationType")
	delete(f, "CipherPolicyId")
	delete(f, "ServerCertificates")
	delete(f, "ClientCaCertificates")
	delete(f, "GetRealIpType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyListenerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyListenerResponseParams struct {
	// <p>Task ID.</p>
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

type OriginHeader struct {
	// <p>Key.</p><p>Parameter format: 1. The string literal only contains printable ASCII characters. 2. Cannot contain these characters ()&lt;&gt;@,;:\&quot;/[ ]?={ }</p><p>Input parameter limitation: length 1-40.</p>
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// <p>Value.</p><p>Input parameter limit: length cannot exceed 128.</p><p>If the string contains $, you can only configure '$remote_addr', '$remote_port'; otherwise, it is not supported.</p>
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type PortOverride struct {
	// Listening port.
	ListenerPort *uint64 `json:"ListenerPort,omitnil,omitempty" name:"ListenerPort"`

	// Mapping port.
	EndpointPort *uint64 `json:"EndpointPort,omitnil,omitempty" name:"EndpointPort"`
}

type PortRanges struct {
	// Start port.
	FromPort *uint64 `json:"FromPort,omitnil,omitempty" name:"FromPort"`

	// Destination port.
	ToPort *uint64 `json:"ToPort,omitnil,omitempty" name:"ToPort"`
}

// Predefined struct for user
type ReplaceListenerAdditionalCertRequestParams struct {
	// Global acceleration instance ID.
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// Listener ID.
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// Certificate ID.
	AdditionalCertificate *string `json:"AdditionalCertificate,omitnil,omitempty" name:"AdditionalCertificate"`

	// Old certificate ID.
	OldCertificate *string `json:"OldCertificate,omitnil,omitempty" name:"OldCertificate"`
}

type ReplaceListenerAdditionalCertRequest struct {
	*tchttp.BaseRequest
	
	// Global acceleration instance ID.
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// Listener ID.
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// Certificate ID.
	AdditionalCertificate *string `json:"AdditionalCertificate,omitnil,omitempty" name:"AdditionalCertificate"`

	// Old certificate ID.
	OldCertificate *string `json:"OldCertificate,omitnil,omitempty" name:"OldCertificate"`
}

func (r *ReplaceListenerAdditionalCertRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReplaceListenerAdditionalCertRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GlobalAcceleratorId")
	delete(f, "ListenerId")
	delete(f, "AdditionalCertificate")
	delete(f, "OldCertificate")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReplaceListenerAdditionalCertRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReplaceListenerAdditionalCertResponseParams struct {
	// Task ID.
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ReplaceListenerAdditionalCertResponse struct {
	*tchttp.BaseResponse
	Response *ReplaceListenerAdditionalCertResponseParams `json:"Response"`
}

func (r *ReplaceListenerAdditionalCertResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReplaceListenerAdditionalCertResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResponseHeaders struct {
	// <p>key</p><p>Parameter format: 1. The string only contain printable ASCII characters. 2. Cannot contain these characters ()&lt;&gt;@,;:\&quot;/[ ]?={ }</p><p>Input limit: Length 1-40.</p>
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// <p>value</p><p>Input limit: length cannot exceed 128</p><p>If the string contains $, you can only configure '$remote_addr' and '$remote_port'. Otherwise, it is not supported.</p>
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type RuleAction struct {
	// <p>Behavior type of the Layer 7 forwarding rule</p><p>Enumeration values:</p><ul><li>ForwardGroup: The forwarding policy forwards to a terminal node group.</li><li>Drop: The forwarding policy drops the request.</li></ul>
	RuleActionType *string `json:"RuleActionType,omitnil,omitempty" name:"RuleActionType"`

	// <p>Layer 7 forwarding rule action value</p><p>This field is not required to input when RuleActionType is Drop. This field is required when RuleActionType is ForwardGroup, which requires filling in the custom terminal node group ID. The default terminal node group cannot be configured.</p>
	RuleActionValue *string `json:"RuleActionValue,omitnil,omitempty" name:"RuleActionValue"`
}

type RuleCondition struct {
	// <p>Condition type of Layer 7 forwarding rule</p><p>Enumeration values:</p><ul><li>Path: Path</li></ul>
	RuleConditionType *string `json:"RuleConditionType,omitnil,omitempty" name:"RuleConditionType"`

	// <p>Layer 7 forwarding rule condition value</p><p>Parameter format: The format must match the regular expression: ^[a-zA-Z0-9_.-/]{1,80}$</p><p>The array length cannot exceed 1.</p>
	RuleConditionValue []*string `json:"RuleConditionValue,omitnil,omitempty" name:"RuleConditionValue"`
}

type Tag struct {
	// Tag key.
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// Tag value.
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}