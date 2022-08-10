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

package v20210416

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type APIGWParams struct {
	// HTTPS
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// POST
	Method *string `json:"Method,omitempty" name:"Method"`
}

// Predefined struct for user
type CheckRuleRequestParams struct {

}

type CheckRuleRequest struct {
	*tchttp.BaseRequest
	
}

func (r *CheckRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckRuleResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CheckRuleResponse struct {
	*tchttp.BaseResponse
	Response *CheckRuleResponseParams `json:"Response"`
}

func (r *CheckRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckTransformationRequestParams struct {
	// JSON string to be processed
	Input *string `json:"Input,omitempty" name:"Input"`

	// Transformation rule list
	Transformations []*Transformation `json:"Transformations,omitempty" name:"Transformations"`
}

type CheckTransformationRequest struct {
	*tchttp.BaseRequest
	
	// JSON string to be processed
	Input *string `json:"Input,omitempty" name:"Input"`

	// Transformation rule list
	Transformations []*Transformation `json:"Transformations,omitempty" name:"Transformations"`
}

func (r *CheckTransformationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckTransformationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Input")
	delete(f, "Transformations")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckTransformationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckTransformationResponseParams struct {
	// Data processed by `Transformations`
	Output *string `json:"Output,omitempty" name:"Output"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CheckTransformationResponse struct {
	*tchttp.BaseResponse
	Response *CheckTransformationResponseParams `json:"Response"`
}

func (r *CheckTransformationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckTransformationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CkafkaDeliveryParams struct {
	// ckafka topic name
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// Six-Segment QCS description of CKafka resource
	ResourceDescription *string `json:"ResourceDescription,omitempty" name:"ResourceDescription"`
}

type CkafkaParams struct {
	// kafka offset
	Offset *string `json:"Offset,omitempty" name:"Offset"`

	// ckafka  topic
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
}

type CkafkaTargetParams struct {
	// CKafka topic to be delivered to
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// Retry policy
	RetryPolicy *RetryPolicy `json:"RetryPolicy,omitempty" name:"RetryPolicy"`
}

type Connection struct {
	// Status
	Status *string `json:"Status,omitempty" name:"Status"`

	// Update time
	ModTime *string `json:"ModTime,omitempty" name:"ModTime"`

	// Switch
	Enable *bool `json:"Enable,omitempty" name:"Enable"`

	// Description
	Description *string `json:"Description,omitempty" name:"Description"`

	// Creation time
	AddTime *string `json:"AddTime,omitempty" name:"AddTime"`

	// Connector ID
	ConnectionId *string `json:"ConnectionId,omitempty" name:"ConnectionId"`

	// Event bus ID
	EventBusId *string `json:"EventBusId,omitempty" name:"EventBusId"`

	// Connector description
	ConnectionDescription *ConnectionDescription `json:"ConnectionDescription,omitempty" name:"ConnectionDescription"`

	// Connector name
	ConnectionName *string `json:"ConnectionName,omitempty" name:"ConnectionName"`

	// Type
	Type *string `json:"Type,omitempty" name:"Type"`
}

type ConnectionDescription struct {
	// Six-Segment QCS resource description. For more information, see [Resource Description Method](https://intl.cloud.tencent.com/document/product/598/10606?from_cn_redirect=1)
	ResourceDescription *string `json:"ResourceDescription,omitempty" name:"ResourceDescription"`

	// API Gateway parameters
	// Note: this field may return null, indicating that no valid values can be obtained.
	APIGWParams *APIGWParams `json:"APIGWParams,omitempty" name:"APIGWParams"`

	// CKafka parameters
	// Note: this field may return null, indicating that no valid values can be obtained.
	CkafkaParams *CkafkaParams `json:"CkafkaParams,omitempty" name:"CkafkaParams"`
}

// Predefined struct for user
type CreateConnectionRequestParams struct {
	// Connector description
	ConnectionDescription *ConnectionDescription `json:"ConnectionDescription,omitempty" name:"ConnectionDescription"`

	// Event bus ID
	EventBusId *string `json:"EventBusId,omitempty" name:"EventBusId"`

	// Connector name
	ConnectionName *string `json:"ConnectionName,omitempty" name:"ConnectionName"`

	// Description
	Description *string `json:"Description,omitempty" name:"Description"`

	// Switch
	Enable *bool `json:"Enable,omitempty" name:"Enable"`

	// Type
	Type *string `json:"Type,omitempty" name:"Type"`
}

type CreateConnectionRequest struct {
	*tchttp.BaseRequest
	
	// Connector description
	ConnectionDescription *ConnectionDescription `json:"ConnectionDescription,omitempty" name:"ConnectionDescription"`

	// Event bus ID
	EventBusId *string `json:"EventBusId,omitempty" name:"EventBusId"`

	// Connector name
	ConnectionName *string `json:"ConnectionName,omitempty" name:"ConnectionName"`

	// Description
	Description *string `json:"Description,omitempty" name:"Description"`

	// Switch
	Enable *bool `json:"Enable,omitempty" name:"Enable"`

	// Type
	Type *string `json:"Type,omitempty" name:"Type"`
}

func (r *CreateConnectionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateConnectionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ConnectionDescription")
	delete(f, "EventBusId")
	delete(f, "ConnectionName")
	delete(f, "Description")
	delete(f, "Enable")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateConnectionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateConnectionResponseParams struct {
	// Connector ID
	ConnectionId *string `json:"ConnectionId,omitempty" name:"ConnectionId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateConnectionResponse struct {
	*tchttp.BaseResponse
	Response *CreateConnectionResponseParams `json:"Response"`
}

func (r *CreateConnectionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateConnectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEventBusRequestParams struct {
	// Event bus name, which can contain 2–60 letters, digits, underscores, and hyphens and must start with a letter and end with a digit or letter
	EventBusName *string `json:"EventBusName,omitempty" name:"EventBusName"`

	// Event bus description, which can contain up to 200 characters of any type
	Description *string `json:"Description,omitempty" name:"Description"`
}

type CreateEventBusRequest struct {
	*tchttp.BaseRequest
	
	// Event bus name, which can contain 2–60 letters, digits, underscores, and hyphens and must start with a letter and end with a digit or letter
	EventBusName *string `json:"EventBusName,omitempty" name:"EventBusName"`

	// Event bus description, which can contain up to 200 characters of any type
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *CreateEventBusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEventBusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EventBusName")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateEventBusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEventBusResponseParams struct {
	// Event bus ID
	EventBusId *string `json:"EventBusId,omitempty" name:"EventBusId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateEventBusResponse struct {
	*tchttp.BaseResponse
	Response *CreateEventBusResponseParams `json:"Response"`
}

func (r *CreateEventBusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEventBusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRuleRequestParams struct {
	// See [Event Pattern](https://intl.cloud.tencent.com/document/product/1359/56084?from_cn_redirect=1)
	EventPattern *string `json:"EventPattern,omitempty" name:"EventPattern"`

	// Event bus ID.
	EventBusId *string `json:"EventBusId,omitempty" name:"EventBusId"`

	// Event bus name, which can contain 2–60 letters, digits, underscores, and hyphens and must start with a letter and end with a digit or letter
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// Switch.
	Enable *bool `json:"Enable,omitempty" name:"Enable"`

	// Event bus description, which can contain up to 200 characters of any type
	Description *string `json:"Description,omitempty" name:"Description"`
}

type CreateRuleRequest struct {
	*tchttp.BaseRequest
	
	// See [Event Pattern](https://intl.cloud.tencent.com/document/product/1359/56084?from_cn_redirect=1)
	EventPattern *string `json:"EventPattern,omitempty" name:"EventPattern"`

	// Event bus ID.
	EventBusId *string `json:"EventBusId,omitempty" name:"EventBusId"`

	// Event bus name, which can contain 2–60 letters, digits, underscores, and hyphens and must start with a letter and end with a digit or letter
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// Switch.
	Enable *bool `json:"Enable,omitempty" name:"Enable"`

	// Event bus description, which can contain up to 200 characters of any type
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *CreateRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EventPattern")
	delete(f, "EventBusId")
	delete(f, "RuleName")
	delete(f, "Enable")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRuleResponseParams struct {
	// Event rule ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateRuleResponse struct {
	*tchttp.BaseResponse
	Response *CreateRuleResponseParams `json:"Response"`
}

func (r *CreateRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTargetRequestParams struct {
	// Event bus ID
	EventBusId *string `json:"EventBusId,omitempty" name:"EventBusId"`

	// Target type
	Type *string `json:"Type,omitempty" name:"Type"`

	// Target description
	TargetDescription *TargetDescription `json:"TargetDescription,omitempty" name:"TargetDescription"`

	// Event rule ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
}

type CreateTargetRequest struct {
	*tchttp.BaseRequest
	
	// Event bus ID
	EventBusId *string `json:"EventBusId,omitempty" name:"EventBusId"`

	// Target type
	Type *string `json:"Type,omitempty" name:"Type"`

	// Target description
	TargetDescription *TargetDescription `json:"TargetDescription,omitempty" name:"TargetDescription"`

	// Event rule ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
}

func (r *CreateTargetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTargetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EventBusId")
	delete(f, "Type")
	delete(f, "TargetDescription")
	delete(f, "RuleId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTargetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTargetResponseParams struct {
	// Target ID
	TargetId *string `json:"TargetId,omitempty" name:"TargetId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateTargetResponse struct {
	*tchttp.BaseResponse
	Response *CreateTargetResponseParams `json:"Response"`
}

func (r *CreateTargetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTargetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTransformationRequestParams struct {
	// Event bus ID
	EventBusId *string `json:"EventBusId,omitempty" name:"EventBusId"`

	// Rule ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// Transformation rule list (currently, only one is supported)
	Transformations []*Transformation `json:"Transformations,omitempty" name:"Transformations"`
}

type CreateTransformationRequest struct {
	*tchttp.BaseRequest
	
	// Event bus ID
	EventBusId *string `json:"EventBusId,omitempty" name:"EventBusId"`

	// Rule ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// Transformation rule list (currently, only one is supported)
	Transformations []*Transformation `json:"Transformations,omitempty" name:"Transformations"`
}

func (r *CreateTransformationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTransformationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EventBusId")
	delete(f, "RuleId")
	delete(f, "Transformations")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTransformationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTransformationResponseParams struct {
	// Generated transformer ID
	TransformationId *string `json:"TransformationId,omitempty" name:"TransformationId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateTransformationResponse struct {
	*tchttp.BaseResponse
	Response *CreateTransformationResponseParams `json:"Response"`
}

func (r *CreateTransformationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTransformationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeadLetterConfig struct {
	// Three modes are supported: DLQ, drop, and ignore error, which correspond to `DLQ`, `DROP`, and `IGNORE_ERROR` respectively
	DisposeMethod *string `json:"DisposeMethod,omitempty" name:"DisposeMethod"`

	// If the DLQ mode is set, this option is required. Error messages will be delivered to the corresponding Kafka topic
	// Note: this field may return null, indicating that no valid values can be obtained.
	CkafkaDeliveryParams *CkafkaDeliveryParams `json:"CkafkaDeliveryParams,omitempty" name:"CkafkaDeliveryParams"`
}

// Predefined struct for user
type DeleteConnectionRequestParams struct {
	// Connector ID
	ConnectionId *string `json:"ConnectionId,omitempty" name:"ConnectionId"`

	// Event bus ID
	EventBusId *string `json:"EventBusId,omitempty" name:"EventBusId"`
}

type DeleteConnectionRequest struct {
	*tchttp.BaseRequest
	
	// Connector ID
	ConnectionId *string `json:"ConnectionId,omitempty" name:"ConnectionId"`

	// Event bus ID
	EventBusId *string `json:"EventBusId,omitempty" name:"EventBusId"`
}

func (r *DeleteConnectionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteConnectionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ConnectionId")
	delete(f, "EventBusId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteConnectionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteConnectionResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteConnectionResponse struct {
	*tchttp.BaseResponse
	Response *DeleteConnectionResponseParams `json:"Response"`
}

func (r *DeleteConnectionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteConnectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteEventBusRequestParams struct {
	// Event bus ID
	EventBusId *string `json:"EventBusId,omitempty" name:"EventBusId"`
}

type DeleteEventBusRequest struct {
	*tchttp.BaseRequest
	
	// Event bus ID
	EventBusId *string `json:"EventBusId,omitempty" name:"EventBusId"`
}

func (r *DeleteEventBusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteEventBusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EventBusId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteEventBusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteEventBusResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteEventBusResponse struct {
	*tchttp.BaseResponse
	Response *DeleteEventBusResponseParams `json:"Response"`
}

func (r *DeleteEventBusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteEventBusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRuleRequestParams struct {
	// Event bus ID
	EventBusId *string `json:"EventBusId,omitempty" name:"EventBusId"`

	// Event rule ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
}

type DeleteRuleRequest struct {
	*tchttp.BaseRequest
	
	// Event bus ID
	EventBusId *string `json:"EventBusId,omitempty" name:"EventBusId"`

	// Event rule ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
}

func (r *DeleteRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EventBusId")
	delete(f, "RuleId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRuleResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteRuleResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRuleResponseParams `json:"Response"`
}

func (r *DeleteRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTargetRequestParams struct {
	// Event bus ID
	EventBusId *string `json:"EventBusId,omitempty" name:"EventBusId"`

	// Delivery target ID
	TargetId *string `json:"TargetId,omitempty" name:"TargetId"`

	// Event rule ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
}

type DeleteTargetRequest struct {
	*tchttp.BaseRequest
	
	// Event bus ID
	EventBusId *string `json:"EventBusId,omitempty" name:"EventBusId"`

	// Delivery target ID
	TargetId *string `json:"TargetId,omitempty" name:"TargetId"`

	// Event rule ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
}

func (r *DeleteTargetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTargetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EventBusId")
	delete(f, "TargetId")
	delete(f, "RuleId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTargetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTargetResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteTargetResponse struct {
	*tchttp.BaseResponse
	Response *DeleteTargetResponseParams `json:"Response"`
}

func (r *DeleteTargetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTargetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTransformationRequestParams struct {
	// Event bus ID
	EventBusId *string `json:"EventBusId,omitempty" name:"EventBusId"`

	// Rule ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// Transformer ID
	TransformationId *string `json:"TransformationId,omitempty" name:"TransformationId"`
}

type DeleteTransformationRequest struct {
	*tchttp.BaseRequest
	
	// Event bus ID
	EventBusId *string `json:"EventBusId,omitempty" name:"EventBusId"`

	// Rule ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// Transformer ID
	TransformationId *string `json:"TransformationId,omitempty" name:"TransformationId"`
}

func (r *DeleteTransformationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTransformationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EventBusId")
	delete(f, "RuleId")
	delete(f, "TransformationId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTransformationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTransformationResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteTransformationResponse struct {
	*tchttp.BaseResponse
	Response *DeleteTransformationResponseParams `json:"Response"`
}

func (r *DeleteTransformationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTransformationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EtlFilter struct {
	// The syntax is the same as that of `Rule`
	Filter *string `json:"Filter,omitempty" name:"Filter"`
}

type EventBus struct {
	// Update time
	ModTime *string `json:"ModTime,omitempty" name:"ModTime"`

	// Event bus description, which can contain up to 200 characters of any type
	Description *string `json:"Description,omitempty" name:"Description"`

	// Creation time
	AddTime *string `json:"AddTime,omitempty" name:"AddTime"`

	// Event bus name, which can contain 2–60 letters, digits, underscores, and hyphens and must start with a letter and end with a digit or letter
	EventBusName *string `json:"EventBusName,omitempty" name:"EventBusName"`

	// Event bus ID
	EventBusId *string `json:"EventBusId,omitempty" name:"EventBusId"`

	// Event bus type
	Type *string `json:"Type,omitempty" name:"Type"`
}

type Extraction struct {
	// JsonPath, which will be `$.` by default if not specified
	ExtractionInputPath *string `json:"ExtractionInputPath,omitempty" name:"ExtractionInputPath"`

	// Valid values: TEXT/JSON
	Format *string `json:"Format,omitempty" name:"Format"`

	// Only required for `Text`
	// Note: this field may return null, indicating that no valid values can be obtained.
	TextParams *TextParams `json:"TextParams,omitempty" name:"TextParams"`
}

type Filter struct {
	// One or more filter values.
	Values []*string `json:"Values,omitempty" name:"Values"`

	// Filter name.
	Name *string `json:"Name,omitempty" name:"Name"`
}

// Predefined struct for user
type GetEventBusRequestParams struct {
	// Event bus ID
	EventBusId *string `json:"EventBusId,omitempty" name:"EventBusId"`
}

type GetEventBusRequest struct {
	*tchttp.BaseRequest
	
	// Event bus ID
	EventBusId *string `json:"EventBusId,omitempty" name:"EventBusId"`
}

func (r *GetEventBusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetEventBusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EventBusId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetEventBusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetEventBusResponseParams struct {
	// Update time
	ModTime *string `json:"ModTime,omitempty" name:"ModTime"`

	// Event bus description
	Description *string `json:"Description,omitempty" name:"Description"`

	// Log topic ID
	ClsTopicId *string `json:"ClsTopicId,omitempty" name:"ClsTopicId"`

	// Creation time
	AddTime *string `json:"AddTime,omitempty" name:"AddTime"`

	// Logset ID
	ClsLogsetId *string `json:"ClsLogsetId,omitempty" name:"ClsLogsetId"`

	// Event bus name
	EventBusName *string `json:"EventBusName,omitempty" name:"EventBusName"`

	// Event bus ID
	EventBusId *string `json:"EventBusId,omitempty" name:"EventBusId"`

	// (Disused) Event bus type
	Type *string `json:"Type,omitempty" name:"Type"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetEventBusResponse struct {
	*tchttp.BaseResponse
	Response *GetEventBusResponseParams `json:"Response"`
}

func (r *GetEventBusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetEventBusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetRuleRequestParams struct {
	// Event bus ID
	EventBusId *string `json:"EventBusId,omitempty" name:"EventBusId"`

	// Event rule ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
}

type GetRuleRequest struct {
	*tchttp.BaseRequest
	
	// Event bus ID
	EventBusId *string `json:"EventBusId,omitempty" name:"EventBusId"`

	// Event rule ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
}

func (r *GetRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EventBusId")
	delete(f, "RuleId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetRuleResponseParams struct {
	// Event bus ID
	EventBusId *string `json:"EventBusId,omitempty" name:"EventBusId"`

	// Event rule ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// Event rule name
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// Event rule status
	Status *string `json:"Status,omitempty" name:"Status"`

	// Switch
	Enable *bool `json:"Enable,omitempty" name:"Enable"`

	// Event rule description
	Description *string `json:"Description,omitempty" name:"Description"`

	// Event pattern
	EventPattern *string `json:"EventPattern,omitempty" name:"EventPattern"`

	// Creation time
	AddTime *string `json:"AddTime,omitempty" name:"AddTime"`

	// Update time
	ModTime *string `json:"ModTime,omitempty" name:"ModTime"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetRuleResponse struct {
	*tchttp.BaseResponse
	Response *GetRuleResponseParams `json:"Response"`
}

func (r *GetRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTransformationRequestParams struct {
	// Event bus ID
	EventBusId *string `json:"EventBusId,omitempty" name:"EventBusId"`

	// Rule ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// Transformer ID
	TransformationId *string `json:"TransformationId,omitempty" name:"TransformationId"`
}

type GetTransformationRequest struct {
	*tchttp.BaseRequest
	
	// Event bus ID
	EventBusId *string `json:"EventBusId,omitempty" name:"EventBusId"`

	// Rule ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// Transformer ID
	TransformationId *string `json:"TransformationId,omitempty" name:"TransformationId"`
}

func (r *GetTransformationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTransformationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EventBusId")
	delete(f, "RuleId")
	delete(f, "TransformationId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetTransformationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTransformationResponseParams struct {
	// Transformation rule list
	Transformations []*Transformation `json:"Transformations,omitempty" name:"Transformations"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetTransformationResponse struct {
	*tchttp.BaseResponse
	Response *GetTransformationResponseParams `json:"Response"`
}

func (r *GetTransformationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTransformationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListConnectionsRequestParams struct {
	// Event bus ID
	EventBusId *string `json:"EventBusId,omitempty" name:"EventBusId"`

	// Indicates by which field to sort the returned results. Valid values: AddTime, ModTime
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Indicates whether the returned results are sorted in ascending or descending order. Valid values: ASC, DESC
	Order *string `json:"Order,omitempty" name:"Order"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

type ListConnectionsRequest struct {
	*tchttp.BaseRequest
	
	// Event bus ID
	EventBusId *string `json:"EventBusId,omitempty" name:"EventBusId"`

	// Indicates by which field to sort the returned results. Valid values: AddTime, ModTime
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Indicates whether the returned results are sorted in ascending or descending order. Valid values: ASC, DESC
	Order *string `json:"Order,omitempty" name:"Order"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *ListConnectionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListConnectionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EventBusId")
	delete(f, "OrderBy")
	delete(f, "Limit")
	delete(f, "Order")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListConnectionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListConnectionsResponseParams struct {
	// Connector information
	Connections []*Connection `json:"Connections,omitempty" name:"Connections"`

	// Total number of connectors
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ListConnectionsResponse struct {
	*tchttp.BaseResponse
	Response *ListConnectionsResponseParams `json:"Response"`
}

func (r *ListConnectionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListConnectionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListEventBusesRequestParams struct {
	// Indicates by which field to sort the returned results. Valid values: AddTime (creation time), ModTime (modification time)
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Indicates whether the returned results are sorted in ascending or descending order. Valid values: ASC (ascending order), DESC (descending order)
	Order *string `json:"Order,omitempty" name:"Order"`

	// Filter. For more information, see the Instance Filter Table below. Each request can contain up to 10 `Filters` and 5 `Filter.Values`.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// Pagination offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

type ListEventBusesRequest struct {
	*tchttp.BaseRequest
	
	// Indicates by which field to sort the returned results. Valid values: AddTime (creation time), ModTime (modification time)
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Indicates whether the returned results are sorted in ascending or descending order. Valid values: ASC (ascending order), DESC (descending order)
	Order *string `json:"Order,omitempty" name:"Order"`

	// Filter. For more information, see the Instance Filter Table below. Each request can contain up to 10 `Filters` and 5 `Filter.Values`.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// Pagination offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *ListEventBusesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListEventBusesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OrderBy")
	delete(f, "Limit")
	delete(f, "Order")
	delete(f, "Filters")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListEventBusesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListEventBusesResponseParams struct {
	// Event bus information
	EventBuses []*EventBus `json:"EventBuses,omitempty" name:"EventBuses"`

	// Total number of event buses
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ListEventBusesResponse struct {
	*tchttp.BaseResponse
	Response *ListEventBusesResponseParams `json:"Response"`
}

func (r *ListEventBusesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListEventBusesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListRulesRequestParams struct {
	// Event bus ID
	EventBusId *string `json:"EventBusId,omitempty" name:"EventBusId"`

	// Indicates by which field to sort the returned results. Valid values: AddTime (creation time), ModTime (modification time)
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Pagination offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Indicates whether the returned results are sorted in ascending or descending order. Valid values: ASC (ascending order), DESC (descending order)
	Order *string `json:"Order,omitempty" name:"Order"`
}

type ListRulesRequest struct {
	*tchttp.BaseRequest
	
	// Event bus ID
	EventBusId *string `json:"EventBusId,omitempty" name:"EventBusId"`

	// Indicates by which field to sort the returned results. Valid values: AddTime (creation time), ModTime (modification time)
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Pagination offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Indicates whether the returned results are sorted in ascending or descending order. Valid values: ASC (ascending order), DESC (descending order)
	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *ListRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EventBusId")
	delete(f, "OrderBy")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Order")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListRulesResponseParams struct {
	// Event rule information
	Rules []*Rule `json:"Rules,omitempty" name:"Rules"`

	// Total number of event rules
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ListRulesResponse struct {
	*tchttp.BaseResponse
	Response *ListRulesResponseParams `json:"Response"`
}

func (r *ListRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListTargetsRequestParams struct {
	// Event rule ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// Event bus ID
	EventBusId *string `json:"EventBusId,omitempty" name:"EventBusId"`

	// Indicates by which field to sort the returned results. Valid values: AddTime (creation time), ModTime (modification time)
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Pagination offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Indicates whether the returned results are sorted in ascending or descending order. Valid values: ASC (ascending order), DESC (descending order)
	Order *string `json:"Order,omitempty" name:"Order"`
}

type ListTargetsRequest struct {
	*tchttp.BaseRequest
	
	// Event rule ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// Event bus ID
	EventBusId *string `json:"EventBusId,omitempty" name:"EventBusId"`

	// Indicates by which field to sort the returned results. Valid values: AddTime (creation time), ModTime (modification time)
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Pagination offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Indicates whether the returned results are sorted in ascending or descending order. Valid values: ASC (ascending order), DESC (descending order)
	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *ListTargetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListTargetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleId")
	delete(f, "EventBusId")
	delete(f, "OrderBy")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Order")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListTargetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListTargetsResponseParams struct {
	// Total number of targets
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Target information
	Targets []*Target `json:"Targets,omitempty" name:"Targets"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ListTargetsResponse struct {
	*tchttp.BaseResponse
	Response *ListTargetsResponseParams `json:"Response"`
}

func (r *ListTargetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListTargetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OutputStructParam struct {
	// Key in the corresponding JSON output
	Key *string `json:"Key,omitempty" name:"Key"`

	// You can enter a JsonPath, constant, or built-in date type
	Value *string `json:"Value,omitempty" name:"Value"`

	// Data type of `Value`. Valid values: STRING, NUMBER, BOOLEAN, NULL, SYS_VARIABLE, JSONPATH
	ValueType *string `json:"ValueType,omitempty" name:"ValueType"`
}

type RetryPolicy struct {
	// Retry interval in seconds
	RetryInterval *uint64 `json:"RetryInterval,omitempty" name:"RetryInterval"`

	// Maximum number of retries
	MaxRetryAttempts *uint64 `json:"MaxRetryAttempts,omitempty" name:"MaxRetryAttempts"`
}

type Rule struct {
	// Status
	Status *string `json:"Status,omitempty" name:"Status"`

	// Modification time
	ModTime *string `json:"ModTime,omitempty" name:"ModTime"`

	// Switch
	Enable *bool `json:"Enable,omitempty" name:"Enable"`

	// Description
	Description *string `json:"Description,omitempty" name:"Description"`

	// Rule ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// Creation time
	AddTime *string `json:"AddTime,omitempty" name:"AddTime"`

	// Event bus ID
	EventBusId *string `json:"EventBusId,omitempty" name:"EventBusId"`

	// Rule name
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// Target overview
	// Note: this field may return null, indicating that no valid values can be obtained.
	Targets []*TargetBrief `json:"Targets,omitempty" name:"Targets"`

	// DLQ rule set by the rule, which may be null
	// Note: this field may return null, indicating that no valid values can be obtained.
	DeadLetterConfig *DeadLetterConfig `json:"DeadLetterConfig,omitempty" name:"DeadLetterConfig"`
}

type SCFParams struct {
	// Maximum waiting time for batch delivery
	BatchTimeout *int64 `json:"BatchTimeout,omitempty" name:"BatchTimeout"`

	// Maximum number of events in batch delivery
	BatchEventCount *int64 `json:"BatchEventCount,omitempty" name:"BatchEventCount"`

	// Enables batch delivery
	EnableBatchDelivery *bool `json:"EnableBatchDelivery,omitempty" name:"EnableBatchDelivery"`
}

type Target struct {
	// Target type
	Type *string `json:"Type,omitempty" name:"Type"`

	// Event bus ID
	EventBusId *string `json:"EventBusId,omitempty" name:"EventBusId"`

	// Target ID
	TargetId *string `json:"TargetId,omitempty" name:"TargetId"`

	// Target description
	TargetDescription *TargetDescription `json:"TargetDescription,omitempty" name:"TargetDescription"`

	// Event rule ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// Enables batch delivery
	// Note: this field may return null, indicating that no valid values can be obtained.
	EnableBatchDelivery *bool `json:"EnableBatchDelivery,omitempty" name:"EnableBatchDelivery"`

	// Maximum waiting time for batch delivery
	// Note: this field may return null, indicating that no valid values can be obtained.
	BatchTimeout *int64 `json:"BatchTimeout,omitempty" name:"BatchTimeout"`

	// Maximum number of events in batch delivery
	// Note: this field may return null, indicating that no valid values can be obtained.
	BatchEventCount *int64 `json:"BatchEventCount,omitempty" name:"BatchEventCount"`
}

type TargetBrief struct {
	// Target ID
	TargetId *string `json:"TargetId,omitempty" name:"TargetId"`

	// Target type
	Type *string `json:"Type,omitempty" name:"Type"`
}

type TargetDescription struct {
	// Six-Segment QCS resource description. For more information, see [Resource Description Method](https://intl.cloud.tencent.com/document/product/598/10606?from_cn_redirect=1)
	ResourceDescription *string `json:"ResourceDescription,omitempty" name:"ResourceDescription"`

	// SCF parameter
	SCFParams *SCFParams `json:"SCFParams,omitempty" name:"SCFParams"`

	// CKafka parameters
	CkafkaTargetParams *CkafkaTargetParams `json:"CkafkaTargetParams,omitempty" name:"CkafkaTargetParams"`
}

type TextParams struct {
	// Comma, | , tab, space, line break, %, or #, which can contain only 1 character.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Separator *string `json:"Separator,omitempty" name:"Separator"`

	// Entered regex (128 characters)
	// Note: this field may return null, indicating that no valid values can be obtained.
	Regex *string `json:"Regex,omitempty" name:"Regex"`
}

type Transform struct {
	// Describes how to transform data
	OutputStructs []*OutputStructParam `json:"OutputStructs,omitempty" name:"OutputStructs"`
}

type Transformation struct {
	// Describes how to extract data
	// Note: this field may return null, indicating that no valid values can be obtained.
	Extraction *Extraction `json:"Extraction,omitempty" name:"Extraction"`

	// Describes how to filter data
	// Note: this field may return null, indicating that no valid values can be obtained.
	EtlFilter *EtlFilter `json:"EtlFilter,omitempty" name:"EtlFilter"`

	// Describes how to transform data
	// Note: this field may return null, indicating that no valid values can be obtained.
	Transform *Transform `json:"Transform,omitempty" name:"Transform"`
}

// Predefined struct for user
type UpdateConnectionRequestParams struct {
	// Connector ID
	ConnectionId *string `json:"ConnectionId,omitempty" name:"ConnectionId"`

	// Event bus ID
	EventBusId *string `json:"EventBusId,omitempty" name:"EventBusId"`

	// Switch
	Enable *bool `json:"Enable,omitempty" name:"Enable"`

	// Description
	Description *string `json:"Description,omitempty" name:"Description"`

	// Connector name
	ConnectionName *string `json:"ConnectionName,omitempty" name:"ConnectionName"`
}

type UpdateConnectionRequest struct {
	*tchttp.BaseRequest
	
	// Connector ID
	ConnectionId *string `json:"ConnectionId,omitempty" name:"ConnectionId"`

	// Event bus ID
	EventBusId *string `json:"EventBusId,omitempty" name:"EventBusId"`

	// Switch
	Enable *bool `json:"Enable,omitempty" name:"Enable"`

	// Description
	Description *string `json:"Description,omitempty" name:"Description"`

	// Connector name
	ConnectionName *string `json:"ConnectionName,omitempty" name:"ConnectionName"`
}

func (r *UpdateConnectionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateConnectionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ConnectionId")
	delete(f, "EventBusId")
	delete(f, "Enable")
	delete(f, "Description")
	delete(f, "ConnectionName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateConnectionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateConnectionResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateConnectionResponse struct {
	*tchttp.BaseResponse
	Response *UpdateConnectionResponseParams `json:"Response"`
}

func (r *UpdateConnectionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateConnectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateEventBusRequestParams struct {
	// Event bus ID
	EventBusId *string `json:"EventBusId,omitempty" name:"EventBusId"`

	// Event bus description, which can contain up to 200 characters of any type
	Description *string `json:"Description,omitempty" name:"Description"`

	// Event bus name, which can contain 2–60 letters, digits, underscores, and hyphens and must start with a letter and end with a digit or letter
	EventBusName *string `json:"EventBusName,omitempty" name:"EventBusName"`
}

type UpdateEventBusRequest struct {
	*tchttp.BaseRequest
	
	// Event bus ID
	EventBusId *string `json:"EventBusId,omitempty" name:"EventBusId"`

	// Event bus description, which can contain up to 200 characters of any type
	Description *string `json:"Description,omitempty" name:"Description"`

	// Event bus name, which can contain 2–60 letters, digits, underscores, and hyphens and must start with a letter and end with a digit or letter
	EventBusName *string `json:"EventBusName,omitempty" name:"EventBusName"`
}

func (r *UpdateEventBusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateEventBusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EventBusId")
	delete(f, "Description")
	delete(f, "EventBusName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateEventBusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateEventBusResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateEventBusResponse struct {
	*tchttp.BaseResponse
	Response *UpdateEventBusResponseParams `json:"Response"`
}

func (r *UpdateEventBusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateEventBusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateRuleRequestParams struct {
	// Event rule ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// Event bus ID
	EventBusId *string `json:"EventBusId,omitempty" name:"EventBusId"`

	// Switch.
	Enable *bool `json:"Enable,omitempty" name:"Enable"`

	// Rule description, which can contain up to 200 characters of any type.
	Description *string `json:"Description,omitempty" name:"Description"`

	// Event rule name, which can contain 2–60 letters, digits, underscores, and hyphens and must start with a letter and end with a digit or letter
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
}

type UpdateRuleRequest struct {
	*tchttp.BaseRequest
	
	// Event rule ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// Event bus ID
	EventBusId *string `json:"EventBusId,omitempty" name:"EventBusId"`

	// Switch.
	Enable *bool `json:"Enable,omitempty" name:"Enable"`

	// Rule description, which can contain up to 200 characters of any type.
	Description *string `json:"Description,omitempty" name:"Description"`

	// Event rule name, which can contain 2–60 letters, digits, underscores, and hyphens and must start with a letter and end with a digit or letter
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
}

func (r *UpdateRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleId")
	delete(f, "EventBusId")
	delete(f, "Enable")
	delete(f, "Description")
	delete(f, "RuleName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateRuleResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateRuleResponse struct {
	*tchttp.BaseResponse
	Response *UpdateRuleResponseParams `json:"Response"`
}

func (r *UpdateRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateTargetRequestParams struct {
	// Event bus ID
	EventBusId *string `json:"EventBusId,omitempty" name:"EventBusId"`

	// Event rule ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// Delivery target ID
	TargetId *string `json:"TargetId,omitempty" name:"TargetId"`

	// Enables batch delivery
	EnableBatchDelivery *bool `json:"EnableBatchDelivery,omitempty" name:"EnableBatchDelivery"`

	// Maximum waiting time for batch delivery
	BatchTimeout *int64 `json:"BatchTimeout,omitempty" name:"BatchTimeout"`

	// Maximum number of events in batch delivery
	BatchEventCount *int64 `json:"BatchEventCount,omitempty" name:"BatchEventCount"`
}

type UpdateTargetRequest struct {
	*tchttp.BaseRequest
	
	// Event bus ID
	EventBusId *string `json:"EventBusId,omitempty" name:"EventBusId"`

	// Event rule ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// Delivery target ID
	TargetId *string `json:"TargetId,omitempty" name:"TargetId"`

	// Enables batch delivery
	EnableBatchDelivery *bool `json:"EnableBatchDelivery,omitempty" name:"EnableBatchDelivery"`

	// Maximum waiting time for batch delivery
	BatchTimeout *int64 `json:"BatchTimeout,omitempty" name:"BatchTimeout"`

	// Maximum number of events in batch delivery
	BatchEventCount *int64 `json:"BatchEventCount,omitempty" name:"BatchEventCount"`
}

func (r *UpdateTargetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateTargetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EventBusId")
	delete(f, "RuleId")
	delete(f, "TargetId")
	delete(f, "EnableBatchDelivery")
	delete(f, "BatchTimeout")
	delete(f, "BatchEventCount")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateTargetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateTargetResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateTargetResponse struct {
	*tchttp.BaseResponse
	Response *UpdateTargetResponseParams `json:"Response"`
}

func (r *UpdateTargetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateTargetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateTransformationRequestParams struct {
	// Event bus ID
	EventBusId *string `json:"EventBusId,omitempty" name:"EventBusId"`

	// Rule ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// Transformer ID
	TransformationId *string `json:"TransformationId,omitempty" name:"TransformationId"`

	// Transformation rule list (currently, only one is supported)
	Transformations []*Transformation `json:"Transformations,omitempty" name:"Transformations"`
}

type UpdateTransformationRequest struct {
	*tchttp.BaseRequest
	
	// Event bus ID
	EventBusId *string `json:"EventBusId,omitempty" name:"EventBusId"`

	// Rule ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// Transformer ID
	TransformationId *string `json:"TransformationId,omitempty" name:"TransformationId"`

	// Transformation rule list (currently, only one is supported)
	Transformations []*Transformation `json:"Transformations,omitempty" name:"Transformations"`
}

func (r *UpdateTransformationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateTransformationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EventBusId")
	delete(f, "RuleId")
	delete(f, "TransformationId")
	delete(f, "Transformations")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateTransformationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateTransformationResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateTransformationResponse struct {
	*tchttp.BaseResponse
	Response *UpdateTransformationResponseParams `json:"Response"`
}

func (r *UpdateTransformationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateTransformationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}