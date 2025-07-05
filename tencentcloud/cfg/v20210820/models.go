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

package v20210820

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

type ActionFieldConfigDetail struct {
	// Component type
	// The options are as follows:
	// input: text box
	// textarea: multi-line text box
	// number: number input box
	// select: selector
	// cascader: cascade selector
	// radio: single choice
	// time: time selection
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Component label
	Lable *string `json:"Lable,omitnil,omitempty" name:"Lable"`

	// Unique identifier of the component, key when it is sent back to the backend
	Field *string `json:"Field,omitnil,omitempty" name:"Field"`

	// Default value
	// Note: This field may return null, indicating that no valid values can be obtained.
	DefaultValue *string `json:"DefaultValue,omitnil,omitempty" name:"DefaultValue"`

	// The supported configuration items are as follows. You can select the configuration items as needed. If no configuration is required, set the value to {}.
	// 
	// {  placeholder: string (placeholder)
	// 
	//   tooltip: string (prompt message)
	// 
	//   reg: RegExp (regular expression for the input content format)
	// 
	//   max: number (maximum number of input characters for text boxes and upper limit of the input number for number input boxes)
	// 
	//   min: number (lower limit of the input number for number input boxes)
	// 
	//   step: number (step size for number input boxes; default value: 1)
	// 
	//   format: string (format for time selection, for example YYYY-MM-DD and YYYY-MM-DD HH:mm:ss)
	// 
	//   separator: string[] (separator for multi-line input boxes. If it is left blank, no separator is used, and the text string entered by the user is returned directly.)
	// 
	//   multiple: boolean (multiple-choice or not, valid for selectors and cascade selectors)
	// 
	//   options: selector options (support the following two forms)
	// 
	// Directly provide the option array: { value: string; label: string }[]
	// Obtain options by calling the API: { api: string (API URL), params: string[] (interface parameters, corresponding to field of the parameter configuration. The frontend uses the input values of all components corresponding to field as parameters to query data. If no value is input, the frontend directly requests data when components are loaded.) 
	// }
	// }
	Config *string `json:"Config,omitnil,omitempty" name:"Config"`

	// Whether it is required (0: no; 1: yes)
	Required *uint64 `json:"Required,omitnil,omitempty" name:"Required"`

	// The compute configuration passes the verification when other fields that it depends on meet the conditions. (For example, at least one of the three form items must be filled in.)
	// 
	// [fieldName,
	// 
	// { config: This item is retained and will be refined based on subsequent scenes. }
	// 
	// ]
	Validate *string `json:"Validate,omitnil,omitempty" name:"Validate"`

	// Whether it is visible
	Visible *string `json:"Visible,omitnil,omitempty" name:"Visible"`
}

type ActionFieldConfigResult struct {
	// Action ID
	ActionId *uint64 `json:"ActionId,omitnil,omitempty" name:"ActionId"`

	// Action name
	ActionName *string `json:"ActionName,omitnil,omitempty" name:"ActionName"`

	// Filed configuration details corresponding to the action
	ConfigDetail []*ActionFieldConfigDetail `json:"ConfigDetail,omitnil,omitempty" name:"ConfigDetail"`
}

type ActionFilter struct {
	// Keyword
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// Content for search
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type ActionLibraryListResult struct {
	// Action name
	ActionName *string `json:"ActionName,omitnil,omitempty" name:"ActionName"`

	// Action description
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`

	// Action type: ["platform" and "custom"]
	ActionType *string `json:"ActionType,omitnil,omitempty" name:"ActionType"`

	// Creation time
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Creator
	Creator *string `json:"Creator,omitnil,omitempty" name:"Creator"`

	// Update time
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// Action risk description
	RiskDesc *string `json:"RiskDesc,omitnil,omitempty" name:"RiskDesc"`

	// Action ID
	ActionId *uint64 `json:"ActionId,omitnil,omitempty" name:"ActionId"`

	// Action attribute (1: fault; 2: recovery)
	AttributeId *uint64 `json:"AttributeId,omitnil,omitempty" name:"AttributeId"`

	// ID of the associated action
	RelationActionId *uint64 `json:"RelationActionId,omitnil,omitempty" name:"RelationActionId"`

	// Operation command
	ActionCommand *string `json:"ActionCommand,omitnil,omitempty" name:"ActionCommand"`

	// Action type (0: tat; 1: cloud API)
	ActionCommandType *uint64 `json:"ActionCommandType,omitnil,omitempty" name:"ActionCommandType"`

	// Parameters of the custom action, in JSON string format
	ActionContent *string `json:"ActionContent,omitnil,omitempty" name:"ActionContent"`

	// Level-2 type
	// Note: This field may return null, indicating that no valid values can be obtained.
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// Action description
	// Note: This field may return null, indicating that no valid values can be obtained.
	ActionDetail *string `json:"ActionDetail,omitnil,omitempty" name:"ActionDetail"`

	// Whether to allow usage by the current account
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsAllowed *bool `json:"IsAllowed,omitnil,omitempty" name:"IsAllowed"`

	// Link to best practices
	// Note: This field may return null, indicating that no valid values can be obtained.
	ActionBestCase *string `json:"ActionBestCase,omitnil,omitempty" name:"ActionBestCase"`

	// Object type
	// Note: This field may return null, indicating that no valid values can be obtained.
	ObjectType *string `json:"ObjectType,omitnil,omitempty" name:"ObjectType"`

	// Monitoring metric ID list
	// Note: This field may return null, indicating that no valid values can be obtained.
	MetricIdList []*uint64 `json:"MetricIdList,omitnil,omitempty" name:"MetricIdList"`

	// Whether the action is new
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsNewAction *bool `json:"IsNewAction,omitnil,omitempty" name:"IsNewAction"`
}

type ApmServiceInfo struct {
	// Business ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Application name
	// Note: This field may return null, indicating that no valid values can be obtained.
	ServiceNameList []*string `json:"ServiceNameList,omitnil,omitempty" name:"ServiceNameList"`

	// Region ID
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	RegionId *int64 `json:"RegionId,omitnil,omitempty" name:"RegionId"`
}

// Predefined struct for user
type CreateTaskFromActionRequestParams struct {
	// Action ID, which can be obtained by using the action list API DescribeActionLibraryList.
	TaskActionId *uint64 `json:"TaskActionId,omitnil,omitempty" name:"TaskActionId"`

	// ID of the instance participating in the experiment.
	TaskInstances []*string `json:"TaskInstances,omitnil,omitempty" name:"TaskInstances"`

	// Experiment name. If this parameter is left blank, the action name is used by default.
	TaskTitle *string `json:"TaskTitle,omitnil,omitempty" name:"TaskTitle"`

	// Experiment description. If this parameter is left blank, the action description is used by default.
	TaskDescription *string `json:"TaskDescription,omitnil,omitempty" name:"TaskDescription"`

	// General action parameters, which need to be passed in after JSON serialization. The parameters can be obtained by using the action details API DescribeActionFieldConfigList. If this field is left blank, the default action parameters are used by default.
	TaskActionGeneralConfiguration *string `json:"TaskActionGeneralConfiguration,omitnil,omitempty" name:"TaskActionGeneralConfiguration"`

	// Action custom parameters need to be passed in as json serialization. They can be obtained from the action details interface DescribeActionFieldConfigList. If not filled in, the default action parameters will be used. Note: Required parameters have no default values. Be sure to pass in valid values.
	TaskActionCustomConfiguration *string `json:"TaskActionCustomConfiguration,omitnil,omitempty" name:"TaskActionCustomConfiguration"`

	// Automatic experiment pause time, in minutes. If this parameter is left blank, the default value 60 is used.
	TaskPauseDuration *uint64 `json:"TaskPauseDuration,omitnil,omitempty" name:"TaskPauseDuration"`
}

type CreateTaskFromActionRequest struct {
	*tchttp.BaseRequest
	
	// Action ID, which can be obtained by using the action list API DescribeActionLibraryList.
	TaskActionId *uint64 `json:"TaskActionId,omitnil,omitempty" name:"TaskActionId"`

	// ID of the instance participating in the experiment.
	TaskInstances []*string `json:"TaskInstances,omitnil,omitempty" name:"TaskInstances"`

	// Experiment name. If this parameter is left blank, the action name is used by default.
	TaskTitle *string `json:"TaskTitle,omitnil,omitempty" name:"TaskTitle"`

	// Experiment description. If this parameter is left blank, the action description is used by default.
	TaskDescription *string `json:"TaskDescription,omitnil,omitempty" name:"TaskDescription"`

	// General action parameters, which need to be passed in after JSON serialization. The parameters can be obtained by using the action details API DescribeActionFieldConfigList. If this field is left blank, the default action parameters are used by default.
	TaskActionGeneralConfiguration *string `json:"TaskActionGeneralConfiguration,omitnil,omitempty" name:"TaskActionGeneralConfiguration"`

	// Action custom parameters need to be passed in as json serialization. They can be obtained from the action details interface DescribeActionFieldConfigList. If not filled in, the default action parameters will be used. Note: Required parameters have no default values. Be sure to pass in valid values.
	TaskActionCustomConfiguration *string `json:"TaskActionCustomConfiguration,omitnil,omitempty" name:"TaskActionCustomConfiguration"`

	// Automatic experiment pause time, in minutes. If this parameter is left blank, the default value 60 is used.
	TaskPauseDuration *uint64 `json:"TaskPauseDuration,omitnil,omitempty" name:"TaskPauseDuration"`
}

func (r *CreateTaskFromActionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTaskFromActionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskActionId")
	delete(f, "TaskInstances")
	delete(f, "TaskTitle")
	delete(f, "TaskDescription")
	delete(f, "TaskActionGeneralConfiguration")
	delete(f, "TaskActionCustomConfiguration")
	delete(f, "TaskPauseDuration")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTaskFromActionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTaskFromActionResponseParams struct {
	// ID of the successfully created experiment
	TaskId *uint64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateTaskFromActionResponse struct {
	*tchttp.BaseResponse
	Response *CreateTaskFromActionResponseParams `json:"Response"`
}

func (r *CreateTaskFromActionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTaskFromActionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTaskFromTemplateRequestParams struct {
	// Template ID retrieved from the template library
	TemplateId *uint64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// Experiment configuration parameters
	TaskConfig *TaskConfig `json:"TaskConfig,omitnil,omitempty" name:"TaskConfig"`
}

type CreateTaskFromTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Template ID retrieved from the template library
	TemplateId *uint64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// Experiment configuration parameters
	TaskConfig *TaskConfig `json:"TaskConfig,omitnil,omitempty" name:"TaskConfig"`
}

func (r *CreateTaskFromTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTaskFromTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	delete(f, "TaskConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTaskFromTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTaskFromTemplateResponseParams struct {
	// ID of the successfully created experiment
	TaskId *uint64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateTaskFromTemplateResponse struct {
	*tchttp.BaseResponse
	Response *CreateTaskFromTemplateResponseParams `json:"Response"`
}

func (r *CreateTaskFromTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTaskFromTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTaskRequestParams struct {
	// Task ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type DeleteTaskRequest struct {
	*tchttp.BaseRequest
	
	// Task ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *DeleteTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTaskResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteTaskResponse struct {
	*tchttp.BaseResponse
	Response *DeleteTaskResponseParams `json:"Response"`
}

func (r *DeleteTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeActionFieldConfigListRequestParams struct {
	// Action ID list
	ActionIds []*uint64 `json:"ActionIds,omitnil,omitempty" name:"ActionIds"`

	// Object type ID
	ObjectTypeId *uint64 `json:"ObjectTypeId,omitnil,omitempty" name:"ObjectTypeId"`
}

type DescribeActionFieldConfigListRequest struct {
	*tchttp.BaseRequest
	
	// Action ID list
	ActionIds []*uint64 `json:"ActionIds,omitnil,omitempty" name:"ActionIds"`

	// Object type ID
	ObjectTypeId *uint64 `json:"ObjectTypeId,omitnil,omitempty" name:"ObjectTypeId"`
}

func (r *DescribeActionFieldConfigListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeActionFieldConfigListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ActionIds")
	delete(f, "ObjectTypeId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeActionFieldConfigListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeActionFieldConfigListResponseParams struct {
	// List of general filed configuration parameters
	Common []*ActionFieldConfigResult `json:"Common,omitnil,omitempty" name:"Common"`

	// List of action filed configuration parameters
	Results []*ActionFieldConfigResult `json:"Results,omitnil,omitempty" name:"Results"`

	// Information on the decommissioned resource
	// Note: This field may return null, indicating that no valid values can be obtained.
	ResourceOffline []*ResourceOffline `json:"ResourceOffline,omitnil,omitempty" name:"ResourceOffline"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeActionFieldConfigListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeActionFieldConfigListResponseParams `json:"Response"`
}

func (r *DescribeActionFieldConfigListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeActionFieldConfigListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeActionLibraryListRequestParams struct {
	// 0-100
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Default value: 0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Object type ID
	ObjectType *uint64 `json:"ObjectType,omitnil,omitempty" name:"ObjectType"`

	// Keyword value {"action name": "a_title", "description": "a_desc", "action type": "a_type", "creation time": "a_create_time", "level-2 type": "a_resource_type"}
	Filters []*ActionFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Action type. 1: fault action; 2: recovery action.
	Attribute []*int64 `json:"Attribute,omitnil,omitempty" name:"Attribute"`

	// Filter item - action ID
	ActionIds []*uint64 `json:"ActionIds,omitnil,omitempty" name:"ActionIds"`
}

type DescribeActionLibraryListRequest struct {
	*tchttp.BaseRequest
	
	// 0-100
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Default value: 0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Object type ID
	ObjectType *uint64 `json:"ObjectType,omitnil,omitempty" name:"ObjectType"`

	// Keyword value {"action name": "a_title", "description": "a_desc", "action type": "a_type", "creation time": "a_create_time", "level-2 type": "a_resource_type"}
	Filters []*ActionFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Action type. 1: fault action; 2: recovery action.
	Attribute []*int64 `json:"Attribute,omitnil,omitempty" name:"Attribute"`

	// Filter item - action ID
	ActionIds []*uint64 `json:"ActionIds,omitnil,omitempty" name:"ActionIds"`
}

func (r *DescribeActionLibraryListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeActionLibraryListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "ObjectType")
	delete(f, "Filters")
	delete(f, "Attribute")
	delete(f, "ActionIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeActionLibraryListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeActionLibraryListResponseParams struct {
	// Query result list
	Results []*ActionLibraryListResult `json:"Results,omitnil,omitempty" name:"Results"`

	// Number of matching records
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeActionLibraryListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeActionLibraryListResponseParams `json:"Response"`
}

func (r *DescribeActionLibraryListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeActionLibraryListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeObjectTypeListRequestParams struct {
	// Supported object
	// 0: all platform products
	// 1: objects connected to the platform
	// 2: some objects supported by the application
	SupportType *int64 `json:"SupportType,omitnil,omitempty" name:"SupportType"`
}

type DescribeObjectTypeListRequest struct {
	*tchttp.BaseRequest
	
	// Supported object
	// 0: all platform products
	// 1: objects connected to the platform
	// 2: some objects supported by the application
	SupportType *int64 `json:"SupportType,omitnil,omitempty" name:"SupportType"`
}

func (r *DescribeObjectTypeListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeObjectTypeListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SupportType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeObjectTypeListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeObjectTypeListResponseParams struct {
	// Object type list
	ObjectTypeList []*ObjectType `json:"ObjectTypeList,omitnil,omitempty" name:"ObjectTypeList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeObjectTypeListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeObjectTypeListResponseParams `json:"Response"`
}

func (r *DescribeObjectTypeListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeObjectTypeListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePolicy struct {
	// Protection policy ID list
	TaskPolicyIdList []*string `json:"TaskPolicyIdList,omitnil,omitempty" name:"TaskPolicyIdList"`

	// Protection policy status
	TaskPolicyStatus *string `json:"TaskPolicyStatus,omitnil,omitempty" name:"TaskPolicyStatus"`

	// Policy rule
	TaskPolicyRule *string `json:"TaskPolicyRule,omitnil,omitempty" name:"TaskPolicyRule"`

	// Process method when the guardrail policy takes effect. 1: sequential execution, 2: pausing.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TaskPolicyDealType *int64 `json:"TaskPolicyDealType,omitnil,omitempty" name:"TaskPolicyDealType"`
}

// Predefined struct for user
type DescribeTaskExecuteLogsRequestParams struct {
	// Task ID
	TaskId *uint64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Number of returned content lines
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Log start line
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeTaskExecuteLogsRequest struct {
	*tchttp.BaseRequest
	
	// Task ID
	TaskId *uint64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Number of returned content lines
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Log start line
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeTaskExecuteLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskExecuteLogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTaskExecuteLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskExecuteLogsResponseParams struct {
	// Log data
	LogMessage []*string `json:"LogMessage,omitnil,omitempty" name:"LogMessage"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTaskExecuteLogsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTaskExecuteLogsResponseParams `json:"Response"`
}

func (r *DescribeTaskExecuteLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskExecuteLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskListRequestParams struct {
	// Pagination limit
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Pagination offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Experiment name
	TaskTitle *string `json:"TaskTitle,omitnil,omitempty" name:"TaskTitle"`

	// Tag key
	TaskTag []*string `json:"TaskTag,omitnil,omitempty" name:"TaskTag"`

	// Task status (1001: not started; 1002: in progress; 1003: paused; 1004: ended)
	TaskStatus *int64 `json:"TaskStatus,omitnil,omitempty" name:"TaskStatus"`

	// Start time, in fixed format: %Y-%m-%d %H:%M:%S
	TaskStartTime *string `json:"TaskStartTime,omitnil,omitempty" name:"TaskStartTime"`

	// End time, in fixed format: %Y-%m-%d %H:%M:%S
	TaskEndTime *string `json:"TaskEndTime,omitnil,omitempty" name:"TaskEndTime"`

	// Update time, in fixed format: %Y-%m-%d %H:%M:%S
	TaskUpdateTime *string `json:"TaskUpdateTime,omitnil,omitempty" name:"TaskUpdateTime"`

	// Tag pair
	Tags []*TagWithDescribe `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Filtering criteria
	Filters []*ActionFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Experiment ID
	TaskId []*uint64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// ID of the associated application for filtering
	ApplicationId []*string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// Associated application for filtering
	ApplicationName []*string `json:"ApplicationName,omitnil,omitempty" name:"ApplicationName"`

	// Task status for filtering, supporting multiple states (1001: not started; 1002: in progress; 1003: paused; 1004: ended)
	TaskStatusList []*uint64 `json:"TaskStatusList,omitnil,omitempty" name:"TaskStatusList"`
}

type DescribeTaskListRequest struct {
	*tchttp.BaseRequest
	
	// Pagination limit
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Pagination offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Experiment name
	TaskTitle *string `json:"TaskTitle,omitnil,omitempty" name:"TaskTitle"`

	// Tag key
	TaskTag []*string `json:"TaskTag,omitnil,omitempty" name:"TaskTag"`

	// Task status (1001: not started; 1002: in progress; 1003: paused; 1004: ended)
	TaskStatus *int64 `json:"TaskStatus,omitnil,omitempty" name:"TaskStatus"`

	// Start time, in fixed format: %Y-%m-%d %H:%M:%S
	TaskStartTime *string `json:"TaskStartTime,omitnil,omitempty" name:"TaskStartTime"`

	// End time, in fixed format: %Y-%m-%d %H:%M:%S
	TaskEndTime *string `json:"TaskEndTime,omitnil,omitempty" name:"TaskEndTime"`

	// Update time, in fixed format: %Y-%m-%d %H:%M:%S
	TaskUpdateTime *string `json:"TaskUpdateTime,omitnil,omitempty" name:"TaskUpdateTime"`

	// Tag pair
	Tags []*TagWithDescribe `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Filtering criteria
	Filters []*ActionFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Experiment ID
	TaskId []*uint64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// ID of the associated application for filtering
	ApplicationId []*string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// Associated application for filtering
	ApplicationName []*string `json:"ApplicationName,omitnil,omitempty" name:"ApplicationName"`

	// Task status for filtering, supporting multiple states (1001: not started; 1002: in progress; 1003: paused; 1004: ended)
	TaskStatusList []*uint64 `json:"TaskStatusList,omitnil,omitempty" name:"TaskStatusList"`
}

func (r *DescribeTaskListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "TaskTitle")
	delete(f, "TaskTag")
	delete(f, "TaskStatus")
	delete(f, "TaskStartTime")
	delete(f, "TaskEndTime")
	delete(f, "TaskUpdateTime")
	delete(f, "Tags")
	delete(f, "Filters")
	delete(f, "TaskId")
	delete(f, "ApplicationId")
	delete(f, "ApplicationName")
	delete(f, "TaskStatusList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTaskListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskListResponseParams struct {
	// None
	TaskList []*TaskListItem `json:"TaskList,omitnil,omitempty" name:"TaskList"`

	// Number of tables in the list
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTaskListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTaskListResponseParams `json:"Response"`
}

func (r *DescribeTaskListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskPolicyTriggerLogRequestParams struct {
	// Experiment ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Page number
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`

	// Number of entries per page
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type DescribeTaskPolicyTriggerLogRequest struct {
	*tchttp.BaseRequest
	
	// Experiment ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Page number
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`

	// Number of entries per page
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

func (r *DescribeTaskPolicyTriggerLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskPolicyTriggerLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "Page")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTaskPolicyTriggerLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskPolicyTriggerLogResponseParams struct {
	// Triggering log
	// Note: This field may return null, indicating that no valid values can be obtained.
	TriggerLogs []*PolicyTriggerLog `json:"TriggerLogs,omitnil,omitempty" name:"TriggerLogs"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTaskPolicyTriggerLogResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTaskPolicyTriggerLogResponseParams `json:"Response"`
}

func (r *DescribeTaskPolicyTriggerLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskPolicyTriggerLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskRequestParams struct {
	// Task ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type DescribeTaskRequest struct {
	*tchttp.BaseRequest
	
	// Task ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *DescribeTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskResponseParams struct {
	// Task information
	Task *Task `json:"Task,omitnil,omitempty" name:"Task"`

	// Experiment report information corresponding to the task. The value null indicates that no report is exported.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ReportInfo *TaskReportInfo `json:"ReportInfo,omitnil,omitempty" name:"ReportInfo"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTaskResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTaskResponseParams `json:"Response"`
}

func (r *DescribeTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTemplateListRequestParams struct {
	// Pagination limit.Maximum value:100.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Pagination offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Experiment name
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`

	// Tag key
	Tag []*string `json:"Tag,omitnil,omitempty" name:"Tag"`

	// Status. 1: in use; 2: not in use.
	IsUsed *int64 `json:"IsUsed,omitnil,omitempty" name:"IsUsed"`

	// Tag pair
	Tags []*TagWithDescribe `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Template library source. 0: self-built; 1: recommended by experts.
	TemplateSource *int64 `json:"TemplateSource,omitnil,omitempty" name:"TemplateSource"`

	// Template ID
	TemplateIdList []*int64 `json:"TemplateIdList,omitnil,omitempty" name:"TemplateIdList"`

	// Filter parameters
	Filters []*ActionFilter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeTemplateListRequest struct {
	*tchttp.BaseRequest
	
	// Pagination limit.Maximum value:100.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Pagination offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Experiment name
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`

	// Tag key
	Tag []*string `json:"Tag,omitnil,omitempty" name:"Tag"`

	// Status. 1: in use; 2: not in use.
	IsUsed *int64 `json:"IsUsed,omitnil,omitempty" name:"IsUsed"`

	// Tag pair
	Tags []*TagWithDescribe `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Template library source. 0: self-built; 1: recommended by experts.
	TemplateSource *int64 `json:"TemplateSource,omitnil,omitempty" name:"TemplateSource"`

	// Template ID
	TemplateIdList []*int64 `json:"TemplateIdList,omitnil,omitempty" name:"TemplateIdList"`

	// Filter parameters
	Filters []*ActionFilter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeTemplateListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTemplateListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Title")
	delete(f, "Tag")
	delete(f, "IsUsed")
	delete(f, "Tags")
	delete(f, "TemplateSource")
	delete(f, "TemplateIdList")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTemplateListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTemplateListResponseParams struct {
	// Template library list
	TemplateList []*TemplateListItem `json:"TemplateList,omitnil,omitempty" name:"TemplateList"`

	// Number of template libraries in the list
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTemplateListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTemplateListResponseParams `json:"Response"`
}

func (r *DescribeTemplateListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTemplateListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTemplateRequestParams struct {
	// Template library ID
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

type DescribeTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Template library ID
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

func (r *DescribeTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTemplateResponseParams struct {
	// Template library details
	Template *Template `json:"Template,omitnil,omitempty" name:"Template"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTemplateResponseParams `json:"Response"`
}

func (r *DescribeTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExecuteTaskInstanceRequestParams struct {
	// Task ID
	TaskId *uint64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Task action ID
	TaskActionId *uint64 `json:"TaskActionId,omitnil,omitempty" name:"TaskActionId"`

	// Task action instance ID
	TaskInstanceIds []*uint64 `json:"TaskInstanceIds,omitnil,omitempty" name:"TaskInstanceIds"`

	// Whether to operate on the entire task
	IsOperateAll *bool `json:"IsOperateAll,omitnil,omitempty" name:"IsOperateAll"`

	// Operation type (1: start; 2: execute; 3: skip; 5: retry)
	ActionType *uint64 `json:"ActionType,omitnil,omitempty" name:"ActionType"`

	// Action group ID
	TaskGroupId *uint64 `json:"TaskGroupId,omitnil,omitempty" name:"TaskGroupId"`
}

type ExecuteTaskInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Task ID
	TaskId *uint64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Task action ID
	TaskActionId *uint64 `json:"TaskActionId,omitnil,omitempty" name:"TaskActionId"`

	// Task action instance ID
	TaskInstanceIds []*uint64 `json:"TaskInstanceIds,omitnil,omitempty" name:"TaskInstanceIds"`

	// Whether to operate on the entire task
	IsOperateAll *bool `json:"IsOperateAll,omitnil,omitempty" name:"IsOperateAll"`

	// Operation type (1: start; 2: execute; 3: skip; 5: retry)
	ActionType *uint64 `json:"ActionType,omitnil,omitempty" name:"ActionType"`

	// Action group ID
	TaskGroupId *uint64 `json:"TaskGroupId,omitnil,omitempty" name:"TaskGroupId"`
}

func (r *ExecuteTaskInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExecuteTaskInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "TaskActionId")
	delete(f, "TaskInstanceIds")
	delete(f, "IsOperateAll")
	delete(f, "ActionType")
	delete(f, "TaskGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExecuteTaskInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExecuteTaskInstanceResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ExecuteTaskInstanceResponse struct {
	*tchttp.BaseResponse
	Response *ExecuteTaskInstanceResponseParams `json:"Response"`
}

func (r *ExecuteTaskInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExecuteTaskInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExecuteTaskRequestParams struct {
	// ID of the task to be executed
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type ExecuteTaskRequest struct {
	*tchttp.BaseRequest
	
	// ID of the task to be executed
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *ExecuteTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExecuteTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExecuteTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExecuteTaskResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ExecuteTaskResponse struct {
	*tchttp.BaseResponse
	Response *ExecuteTaskResponseParams `json:"Response"`
}

func (r *ExecuteTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExecuteTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTaskRunStatusRequestParams struct {
	// Task ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Task status. 1001: not started; 1002: in progress (executing); 1003: in progress (paused); 1004: execution ended.
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Whether the execution result meets expectations (This field is required when the task status is Execution Ended.)
	IsExpect *bool `json:"IsExpect,omitnil,omitempty" name:"IsExpect"`

	// Experiment result (This field is required when the experiment status changes to Execution Ended.)
	Summary *string `json:"Summary,omitnil,omitempty" name:"Summary"`
}

type ModifyTaskRunStatusRequest struct {
	*tchttp.BaseRequest
	
	// Task ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Task status. 1001: not started; 1002: in progress (executing); 1003: in progress (paused); 1004: execution ended.
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Whether the execution result meets expectations (This field is required when the task status is Execution Ended.)
	IsExpect *bool `json:"IsExpect,omitnil,omitempty" name:"IsExpect"`

	// Experiment result (This field is required when the experiment status changes to Execution Ended.)
	Summary *string `json:"Summary,omitnil,omitempty" name:"Summary"`
}

func (r *ModifyTaskRunStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTaskRunStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "Status")
	delete(f, "IsExpect")
	delete(f, "Summary")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTaskRunStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTaskRunStatusResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyTaskRunStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyTaskRunStatusResponseParams `json:"Response"`
}

func (r *ModifyTaskRunStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTaskRunStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ObjectType struct {
	// Object type ID
	ObjectTypeId *int64 `json:"ObjectTypeId,omitnil,omitempty" name:"ObjectTypeId"`

	// Object type name
	ObjectTypeTitle *string `json:"ObjectTypeTitle,omitnil,omitempty" name:"ObjectTypeTitle"`

	// Level-1 object type
	ObjectTypeLevelOne *string `json:"ObjectTypeLevelOne,omitnil,omitempty" name:"ObjectTypeLevelOne"`

	// Object type parameters
	ObjectTypeParams *ObjectTypeConfig `json:"ObjectTypeParams,omitnil,omitempty" name:"ObjectTypeParams"`

	// JSON parsing rule for TKE APIs. If the value is null, no parsing is needed.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ObjectTypeJsonParse *ObjectTypeJsonParse `json:"ObjectTypeJsonParse,omitnil,omitempty" name:"ObjectTypeJsonParse"`

	// Whether new action is included
	// Note: This field may return null, indicating that no valid values can be obtained.
	ObjectHasNewAction *bool `json:"ObjectHasNewAction,omitnil,omitempty" name:"ObjectHasNewAction"`
}

type ObjectTypeConfig struct {
	// Primary key
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// List of object type configuration fields
	Fields []*ObjectTypeConfigFields `json:"Fields,omitnil,omitempty" name:"Fields"`
}

type ObjectTypeConfigFields struct {
	// instanceId
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// Instance ID
	Header *string `json:"Header,omitnil,omitempty" name:"Header"`

	// Whether the field value needs to be translated. If not, this field returns null.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Transfer *string `json:"Transfer,omitnil,omitempty" name:"Transfer"`

	// Value returned for container Pod resources
	// Note: This field may return null, indicating that no valid values can be obtained.
	JsonParse *string `json:"JsonParse,omitnil,omitempty" name:"JsonParse"`
}

type ObjectTypeJsonParse struct {
	// Namespace
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	NameSpace *string `json:"NameSpace,omitnil,omitempty" name:"NameSpace"`

	// Workload name
	// Note: This field may return null, indicating that no valid values can be obtained.
	WorkloadName *string `json:"WorkloadName,omitnil,omitempty" name:"WorkloadName"`

	// Node IP address
	// Note: This field may return null, indicating that no valid values can be obtained.
	LanIP *string `json:"LanIP,omitnil,omitempty" name:"LanIP"`

	// Node ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type PolicyTriggerLog struct {
	// Experiment ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Name
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Type. 0: trigger; 1: recovery.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TriggerType *int64 `json:"TriggerType,omitnil,omitempty" name:"TriggerType"`

	// Content
	// Note: This field may return null, indicating that no valid values can be obtained.
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// Triggering time
	// Note: This field may return null, indicating that no valid values can be obtained.
	CreatTime *string `json:"CreatTime,omitnil,omitempty" name:"CreatTime"`
}

type ResourceOffline struct {
	// Resource ID
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	ResourceId *int64 `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// Resource decommissioning time
	// Note: This field may return null, indicating that no valid values can be obtained.
	ResourceDeleteTime *string `json:"ResourceDeleteTime,omitnil,omitempty" name:"ResourceDeleteTime"`

	// Resource decommissioning reminder
	// Note: This field may return null, indicating that no valid values can be obtained.
	ResourceDeleteMessage *string `json:"ResourceDeleteMessage,omitnil,omitempty" name:"ResourceDeleteMessage"`
}

type TagWithCreate struct {
	// Tag key
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// Tag value
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

type TagWithDescribe struct {
	// Tag key
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// Tag value
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

type Task struct {
	// Task ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Task name
	TaskTitle *string `json:"TaskTitle,omitnil,omitempty" name:"TaskTitle"`

	// Task description
	TaskDescription *string `json:"TaskDescription,omitnil,omitempty" name:"TaskDescription"`

	// Custom tag
	// Note: This field may return null, indicating that no valid values can be obtained.
	TaskTag *string `json:"TaskTag,omitnil,omitempty" name:"TaskTag"`

	// Task status. 1001: not started; 1002: in progress (executing); 1003: in progress (paused); 1004: execution ended.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TaskStatus *int64 `json:"TaskStatus,omitnil,omitempty" name:"TaskStatus"`

	// Task end status, indicating the status of the task when it ends. 0: not ended; 1: successful; 2: failed; 3: terminated.
	TaskStatusType *int64 `json:"TaskStatusType,omitnil,omitempty" name:"TaskStatusType"`

	// Protection policy
	// Note: This field may return null, indicating that no valid values can be obtained.
	TaskProtectStrategy *string `json:"TaskProtectStrategy,omitnil,omitempty" name:"TaskProtectStrategy"`

	// Task creation time
	TaskCreateTime *string `json:"TaskCreateTime,omitnil,omitempty" name:"TaskCreateTime"`

	// Task update time
	TaskUpdateTime *string `json:"TaskUpdateTime,omitnil,omitempty" name:"TaskUpdateTime"`

	// Task action group
	TaskGroups []*TaskGroup `json:"TaskGroups,omitnil,omitempty" name:"TaskGroups"`

	// Start time
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	TaskStartTime *string `json:"TaskStartTime,omitnil,omitempty" name:"TaskStartTime"`

	// End time
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	TaskEndTime *string `json:"TaskEndTime,omitnil,omitempty" name:"TaskEndTime"`

	// Whether expectations are met. 1: expectations met; 2: expectations not met.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TaskExpect *int64 `json:"TaskExpect,omitnil,omitempty" name:"TaskExpect"`

	// Experiment record
	// Note: This field may return null, indicating that no valid values can be obtained.
	TaskSummary *string `json:"TaskSummary,omitnil,omitempty" name:"TaskSummary"`

	// Task mode. 1: manual execution; 2: automatic execution.
	TaskMode *int64 `json:"TaskMode,omitnil,omitempty" name:"TaskMode"`

	// Automatic pause duration. Unit: minutes.
	TaskPauseDuration *int64 `json:"TaskPauseDuration,omitnil,omitempty" name:"TaskPauseDuration"`

	// Main account that creates the experiment
	TaskOwnerUin *string `json:"TaskOwnerUin,omitnil,omitempty" name:"TaskOwnerUin"`

	// Region ID
	TaskRegionId *int64 `json:"TaskRegionId,omitnil,omitempty" name:"TaskRegionId"`

	// Monitoring metric list
	// Note: This field may return null, indicating that no valid values can be obtained.
	TaskMonitors []*TaskMonitor `json:"TaskMonitors,omitnil,omitempty" name:"TaskMonitors"`

	// Protection policy
	// Note: This field may return null, indicating that no valid values can be obtained.
	TaskPolicy *DescribePolicy `json:"TaskPolicy,omitnil,omitempty" name:"TaskPolicy"`

	// Tag list
	// Note: This field may return null, indicating that no valid values can be obtained.
	Tags []*TagWithDescribe `json:"Tags,omitnil,omitempty" name:"Tags"`

	// ID of the associated experiment plan
	// Note: This field may return null, indicating that no valid values can be obtained.
	TaskPlanId *int64 `json:"TaskPlanId,omitnil,omitempty" name:"TaskPlanId"`

	// Name of the associated experiment plan
	// Note: This field may return null, indicating that no valid values can be obtained.
	TaskPlanTitle *string `json:"TaskPlanTitle,omitnil,omitempty" name:"TaskPlanTitle"`

	// ID of the associated application
	// Note: This field may return null, indicating that no valid values can be obtained.
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// Name of the associated application
	// Note: This field may return null, indicating that no valid values can be obtained.
	ApplicationName *string `json:"ApplicationName,omitnil,omitempty" name:"ApplicationName"`

	// Associated alarm metrics
	// Note: This field may return null, indicating that no valid values can be obtained.
	AlarmPolicy []*string `json:"AlarmPolicy,omitnil,omitempty" name:"AlarmPolicy"`

	// Associated APM services
	// Note: This field may return null, indicating that no valid values can be obtained.
	ApmServiceList []*ApmServiceInfo `json:"ApmServiceList,omitnil,omitempty" name:"ApmServiceList"`

	// ID of the associated threat verification item
	// Note: This field may return null, indicating that no valid values can be obtained.
	VerifyId *uint64 `json:"VerifyId,omitnil,omitempty" name:"VerifyId"`

	// Guardrail processing method. 1: roll back sequentially; 2: pause experiment.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PolicyDealType *int64 `json:"PolicyDealType,omitnil,omitempty" name:"PolicyDealType"`
}

type TaskConfig struct {
	// Action group configurations. The number of configured action groups needs to be consistent with that in the template.
	TaskGroupsConfig []*TaskGroupConfig `json:"TaskGroupsConfig,omitnil,omitempty" name:"TaskGroupsConfig"`

	// Experiment name after change. If this parameter is left blank, the template name is used by default.
	TaskTitle *string `json:"TaskTitle,omitnil,omitempty" name:"TaskTitle"`

	// Experiment description after change. If this parameter is left blank, the template description is used by default.
	TaskDescription *string `json:"TaskDescription,omitnil,omitempty" name:"TaskDescription"`

	// Experiment execution mode. 1: manual execution; 2: automatic execution. If this parameter is left blank, the template execution mode is used by default.
	TaskMode *uint64 `json:"TaskMode,omitnil,omitempty" name:"TaskMode"`

	// Automatic pause time of the experiment, in minutes. If this parameter is left blank, the automatic pause time of the template is used by default.
	TaskPauseDuration *uint64 `json:"TaskPauseDuration,omitnil,omitempty" name:"TaskPauseDuration"`

	// Experiment tag. If this parameter is left blank, the template tag is used by default.
	Tags []*TagWithCreate `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Guardrail processing method. 1: roll back sequentially; 2: pause experiment.
	PolicyDealType *int64 `json:"PolicyDealType,omitnil,omitempty" name:"PolicyDealType"`
}

type TaskGroup struct {
	// Task action ID
	TaskGroupId *int64 `json:"TaskGroupId,omitnil,omitempty" name:"TaskGroupId"`

	// Group name
	TaskGroupTitle *string `json:"TaskGroupTitle,omitnil,omitempty" name:"TaskGroupTitle"`

	// Group description
	// Note: This field may return null, indicating that no valid values can be obtained.
	TaskGroupDescription *string `json:"TaskGroupDescription,omitnil,omitempty" name:"TaskGroupDescription"`

	// Task group order
	TaskGroupOrder *int64 `json:"TaskGroupOrder,omitnil,omitempty" name:"TaskGroupOrder"`

	// Object type ID
	ObjectTypeId *int64 `json:"ObjectTypeId,omitnil,omitempty" name:"ObjectTypeId"`

	// Task group creation time
	TaskGroupCreateTime *string `json:"TaskGroupCreateTime,omitnil,omitempty" name:"TaskGroupCreateTime"`

	// Task group update time
	TaskGroupUpdateTime *string `json:"TaskGroupUpdateTime,omitnil,omitempty" name:"TaskGroupUpdateTime"`

	// List of actions in the group
	TaskGroupActions []*TaskGroupAction `json:"TaskGroupActions,omitnil,omitempty" name:"TaskGroupActions"`

	// Instance list
	TaskGroupInstanceList []*string `json:"TaskGroupInstanceList,omitnil,omitempty" name:"TaskGroupInstanceList"`

	// Execution mode. 1: sequential execution; 2: execution by stage.
	TaskGroupMode *int64 `json:"TaskGroupMode,omitnil,omitempty" name:"TaskGroupMode"`

	// List of instances not involved in the experiment
	TaskGroupDiscardInstanceList []*string `json:"TaskGroupDiscardInstanceList,omitnil,omitempty" name:"TaskGroupDiscardInstanceList"`

	// List of instances involved in the experiment
	// Note: This field may return null, indicating that no valid values can be obtained.
	TaskGroupSelectedInstanceList []*string `json:"TaskGroupSelectedInstanceList,omitnil,omitempty" name:"TaskGroupSelectedInstanceList"`

	// Machine selection rule
	// Note: This field may return null, indicating that no valid values can be obtained.
	TaskGroupInstancesExecuteRule []*TaskGroupInstancesExecuteRules `json:"TaskGroupInstancesExecuteRule,omitnil,omitempty" name:"TaskGroupInstancesExecuteRule"`
}

type TaskGroupAction struct {
	// Task group action ID
	TaskGroupActionId *int64 `json:"TaskGroupActionId,omitnil,omitempty" name:"TaskGroupActionId"`

	// Action instance list of the task group
	TaskGroupInstances []*TaskGroupInstance `json:"TaskGroupInstances,omitnil,omitempty" name:"TaskGroupInstances"`

	// Action ID
	ActionId *int64 `json:"ActionId,omitnil,omitempty" name:"ActionId"`

	// Order of actions in the group
	TaskGroupActionOrder *int64 `json:"TaskGroupActionOrder,omitnil,omitempty" name:"TaskGroupActionOrder"`

	// General configurations of actions in the group
	// Note: This field may return null, indicating that no valid values can be obtained.
	TaskGroupActionGeneralConfiguration *string `json:"TaskGroupActionGeneralConfiguration,omitnil,omitempty" name:"TaskGroupActionGeneralConfiguration"`

	// Custom configurations of actions in the group
	// Note: This field may return null, indicating that no valid values can be obtained.
	TaskGroupActionCustomConfiguration *string `json:"TaskGroupActionCustomConfiguration,omitnil,omitempty" name:"TaskGroupActionCustomConfiguration"`

	// Status of actions in the group
	TaskGroupActionStatus *int64 `json:"TaskGroupActionStatus,omitnil,omitempty" name:"TaskGroupActionStatus"`

	// Action group creation time
	TaskGroupActionCreateTime *string `json:"TaskGroupActionCreateTime,omitnil,omitempty" name:"TaskGroupActionCreateTime"`

	// Action group update time
	TaskGroupActionUpdateTime *string `json:"TaskGroupActionUpdateTime,omitnil,omitempty" name:"TaskGroupActionUpdateTime"`

	// Action name
	ActionTitle *string `json:"ActionTitle,omitnil,omitempty" name:"ActionTitle"`

	// Status. 0: no status; 1: successful; 2; failed; 3: terminated; 4: skipped.
	TaskGroupActionStatusType *int64 `json:"TaskGroupActionStatusType,omitnil,omitempty" name:"TaskGroupActionStatusType"`

	// RandomId
	// Note: This field may return null, indicating that no valid values can be obtained.
	TaskGroupActionRandomId *int64 `json:"TaskGroupActionRandomId,omitnil,omitempty" name:"TaskGroupActionRandomId"`

	// RecoverId
	// Note: This field may return null, indicating that no valid values can be obtained.
	TaskGroupActionRecoverId *int64 `json:"TaskGroupActionRecoverId,omitnil,omitempty" name:"TaskGroupActionRecoverId"`

	// ExecuteId
	// Note: This field may return null, indicating that no valid values can be obtained.
	TaskGroupActionExecuteId *int64 `json:"TaskGroupActionExecuteId,omitnil,omitempty" name:"TaskGroupActionExecuteId"`

	// Called API type. 0: tat; 1: cloud API.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ActionApiType *int64 `json:"ActionApiType,omitnil,omitempty" name:"ActionApiType"`

	// 1: fault; 2: recovery.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ActionAttribute *int64 `json:"ActionAttribute,omitnil,omitempty" name:"ActionAttribute"`

	// Action type: platform and custom
	// Note: This field may return null, indicating that no valid values can be obtained.
	ActionType *string `json:"ActionType,omitnil,omitempty" name:"ActionType"`

	// Whether retry is allowed
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsExecuteRedo *bool `json:"IsExecuteRedo,omitnil,omitempty" name:"IsExecuteRedo"`

	// Action risk level
	// Note: This field may return null, indicating that no valid values can be obtained.
	ActionRisk *string `json:"ActionRisk,omitnil,omitempty" name:"ActionRisk"`

	// Action running time
	// Note: This field may return null, indicating that no valid values can be obtained.
	TaskGroupActionExecuteTime *int64 `json:"TaskGroupActionExecuteTime,omitnil,omitempty" name:"TaskGroupActionExecuteTime"`
}

type TaskGroupActionConfig struct {
	// Order of this action in the action group. The entire order starts from 1. If this parameter is left blank or set to an invalid value, the action whose parameters need to be modified in the template cannot be matched.
	TaskGroupActionOrder *uint64 `json:"TaskGroupActionOrder,omitnil,omitempty" name:"TaskGroupActionOrder"`

	// General action parameters, which need to be passed in after JSON serialization. The parameters can be obtained by using the template details query API. If this field is left blank, action parameters in the template are used by default.
	TaskGroupActionGeneralConfiguration *string `json:"TaskGroupActionGeneralConfiguration,omitnil,omitempty" name:"TaskGroupActionGeneralConfiguration"`

	// Custom action parameters, which need to be passed in after JSON serialization. The parameters can be obtained by using the template details query API. If this field is left blank, action parameters in the template are used by default.
	TaskGroupActionCustomConfiguration *string `json:"TaskGroupActionCustomConfiguration,omitnil,omitempty" name:"TaskGroupActionCustomConfiguration"`
}

type TaskGroupConfig struct {
	// Instance object associated with the action group
	TaskGroupInstances []*string `json:"TaskGroupInstances,omitnil,omitempty" name:"TaskGroupInstances"`

	// Action group name. If this parameter is left blank, the action group name in the template is used by default.
	TaskGroupTitle *string `json:"TaskGroupTitle,omitnil,omitempty" name:"TaskGroupTitle"`

	// Action group description. If this parameter is left blank, the action group description in the template is used by default.
	TaskGroupDescription *string `json:"TaskGroupDescription,omitnil,omitempty" name:"TaskGroupDescription"`

	// Action group execution mode. 1: sequential execution; 2: execution by stage. If this parameter is left blank, the action execution mode in the template is used by default.
	TaskGroupMode *uint64 `json:"TaskGroupMode,omitnil,omitempty" name:"TaskGroupMode"`

	// Action parameters in the action group. If this field is left blank, the action parameters in the template is used by default. You only need to specify the action whose parameters are to be modified during configuration.
	TaskGroupActionsConfig []*TaskGroupActionConfig `json:"TaskGroupActionsConfig,omitnil,omitempty" name:"TaskGroupActionsConfig"`
}

type TaskGroupInstance struct {
	// Instance ID
	TaskGroupInstanceId *int64 `json:"TaskGroupInstanceId,omitnil,omitempty" name:"TaskGroupInstanceId"`

	// Instance ID
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	TaskGroupInstanceObjectId *string `json:"TaskGroupInstanceObjectId,omitnil,omitempty" name:"TaskGroupInstanceObjectId"`

	// Instance action execution status
	TaskGroupInstanceStatus *int64 `json:"TaskGroupInstanceStatus,omitnil,omitempty" name:"TaskGroupInstanceStatus"`

	// Instance creation time
	TaskGroupInstanceCreateTime *string `json:"TaskGroupInstanceCreateTime,omitnil,omitempty" name:"TaskGroupInstanceCreateTime"`

	// Instance update time
	TaskGroupInstanceUpdateTime *string `json:"TaskGroupInstanceUpdateTime,omitnil,omitempty" name:"TaskGroupInstanceUpdateTime"`

	// Status. 0: no status; 1: successful; 2: failed; 3: terminated; 4: skipped.
	TaskGroupInstanceStatusType *int64 `json:"TaskGroupInstanceStatusType,omitnil,omitempty" name:"TaskGroupInstanceStatusType"`

	// Execution start time
	// Note: This field may return null, indicating that no valid values can be obtained.
	TaskGroupInstanceStartTime *string `json:"TaskGroupInstanceStartTime,omitnil,omitempty" name:"TaskGroupInstanceStartTime"`

	// Execution end time
	// Note: This field may return null, indicating that no valid values can be obtained.
	TaskGroupInstanceEndTime *string `json:"TaskGroupInstanceEndTime,omitnil,omitempty" name:"TaskGroupInstanceEndTime"`

	// Instance action execution log
	// Note: This field may return null, indicating that no valid values can be obtained.
	//
	// Deprecated: TaskGroupInstanceExecuteLog is deprecated.
	TaskGroupInstanceExecuteLog *string `json:"TaskGroupInstanceExecuteLog,omitnil,omitempty" name:"TaskGroupInstanceExecuteLog"`

	// Whether the instance can be retried
	// Note: This field may return null, indicating that no valid values can be obtained.
	TaskGroupInstanceIsRedo *bool `json:"TaskGroupInstanceIsRedo,omitnil,omitempty" name:"TaskGroupInstanceIsRedo"`

	// Action instance execution time
	// Note: This field may return null, indicating that no valid values can be obtained.
	TaskGroupInstanceExecuteTime *int64 `json:"TaskGroupInstanceExecuteTime,omitnil,omitempty" name:"TaskGroupInstanceExecuteTime"`
}

type TaskGroupInstancesExecuteRules struct {
	// Instance selection mode
	// Note: This field may return null, indicating that no valid values can be obtained.
	TaskGroupInstancesExecuteMode *int64 `json:"TaskGroupInstancesExecuteMode,omitnil,omitempty" name:"TaskGroupInstancesExecuteMode"`

	// Proportion of selected machines in selection by proportion mode
	// Note: This field may return null, indicating that no valid values can be obtained.
	TaskGroupInstancesExecutePercent *int64 `json:"TaskGroupInstancesExecutePercent,omitnil,omitempty" name:"TaskGroupInstancesExecutePercent"`

	// Number of selected machines in selection by quantity mode
	// Note: This field may return null, indicating that no valid values can be obtained.
	TaskGroupInstancesExecuteNum *int64 `json:"TaskGroupInstancesExecuteNum,omitnil,omitempty" name:"TaskGroupInstancesExecuteNum"`
}

type TaskListItem struct {
	// Task ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Task name
	TaskTitle *string `json:"TaskTitle,omitnil,omitempty" name:"TaskTitle"`

	// Task description
	TaskDescription *string `json:"TaskDescription,omitnil,omitempty" name:"TaskDescription"`

	// Task tag
	// Note: This field may return null, indicating that no valid values can be obtained.
	TaskTag *string `json:"TaskTag,omitnil,omitempty" name:"TaskTag"`

	// Task status (1001: not started; 1002: in progress; 1003: paused; 1004: ended)
	TaskStatus *int64 `json:"TaskStatus,omitnil,omitempty" name:"TaskStatus"`

	// Task creation time
	TaskCreateTime *string `json:"TaskCreateTime,omitnil,omitempty" name:"TaskCreateTime"`

	// Task update time
	TaskUpdateTime *string `json:"TaskUpdateTime,omitnil,omitempty" name:"TaskUpdateTime"`

	// 0: not started; 1: in progress; 2: completed.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TaskPreCheckStatus *int64 `json:"TaskPreCheckStatus,omitnil,omitempty" name:"TaskPreCheckStatus"`

	// Whether the environmental check is passed
	// Note: This field may return null, indicating that no valid values can be obtained.
	TaskPreCheckSuccess *bool `json:"TaskPreCheckSuccess,omitnil,omitempty" name:"TaskPreCheckSuccess"`

	// Whether the experiment result meets expectations. 1: expectations met; 2: expectations not met.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TaskExpect *int64 `json:"TaskExpect,omitnil,omitempty" name:"TaskExpect"`

	// ID of the associated application
	// Note: This field may return null, indicating that no valid values can be obtained.
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// Name of the associated application
	// Note: This field may return null, indicating that no valid values can be obtained.
	ApplicationName *string `json:"ApplicationName,omitnil,omitempty" name:"ApplicationName"`

	// Verification item ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	VerifyId *uint64 `json:"VerifyId,omitnil,omitempty" name:"VerifyId"`

	// Status. 0: no status; 1: successful; 2: failed; 3: terminated.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TaskStatusType *uint64 `json:"TaskStatusType,omitnil,omitempty" name:"TaskStatusType"`
}

type TaskMonitor struct {
	// Experiment monitoring metric ID
	TaskMonitorId *int64 `json:"TaskMonitorId,omitnil,omitempty" name:"TaskMonitorId"`

	// Monitoring metric ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	MetricId *uint64 `json:"MetricId,omitnil,omitempty" name:"MetricId"`

	// Object type ID of the monitoring metric
	TaskMonitorObjectTypeId *int64 `json:"TaskMonitorObjectTypeId,omitnil,omitempty" name:"TaskMonitorObjectTypeId"`

	// Metric name
	MetricName *string `json:"MetricName,omitnil,omitempty" name:"MetricName"`

	// Instance ID list
	InstancesIds []*string `json:"InstancesIds,omitnil,omitempty" name:"InstancesIds"`

	// Metric in Chinese
	// Note: This field may return null, indicating that no valid values can be obtained.
	MetricChineseName *string `json:"MetricChineseName,omitnil,omitempty" name:"MetricChineseName"`

	// Unit
	// Note: This field may return null, indicating that no valid values can be obtained.
	Unit *string `json:"Unit,omitnil,omitempty" name:"Unit"`
}

type TaskReportInfo struct {
	// 0: not started; 1: exporting; 2: export successful; 3: export failed.
	Stage *int64 `json:"Stage,omitnil,omitempty" name:"Stage"`

	// Creation time
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// End time of the validity period
	ExpirationTime *string `json:"ExpirationTime,omitnil,omitempty" name:"ExpirationTime"`

	// Whether it is effective
	Expired *bool `json:"Expired,omitnil,omitempty" name:"Expired"`

	// Address of the COS experiment report file
	// Note: This field may return null, indicating that no valid values can be obtained.
	CosUrl *string `json:"CosUrl,omitnil,omitempty" name:"CosUrl"`

	// Experiment report export log
	// Note: This field may return null, indicating that no valid values can be obtained.
	Log *string `json:"Log,omitnil,omitempty" name:"Log"`
}

type Template struct {
	// Template library ID
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// Template library name
	TemplateTitle *string `json:"TemplateTitle,omitnil,omitempty" name:"TemplateTitle"`

	// Template library description
	TemplateDescription *string `json:"TemplateDescription,omitnil,omitempty" name:"TemplateDescription"`

	// Custom tag
	// Note: This field may return null, indicating that no valid values can be obtained.
	TemplateTag *string `json:"TemplateTag,omitnil,omitempty" name:"TemplateTag"`

	// Usage status. 1: in use; 2: not in use.
	TemplateIsUsed *int64 `json:"TemplateIsUsed,omitnil,omitempty" name:"TemplateIsUsed"`

	// Template library creation time
	TemplateCreateTime *string `json:"TemplateCreateTime,omitnil,omitempty" name:"TemplateCreateTime"`

	// Template library update time
	TemplateUpdateTime *string `json:"TemplateUpdateTime,omitnil,omitempty" name:"TemplateUpdateTime"`

	// Template library mode. 1: manual execution; 2: automatic execution.
	TemplateMode *int64 `json:"TemplateMode,omitnil,omitempty" name:"TemplateMode"`

	// Automatic pause duration. Unit: minutes.
	TemplatePauseDuration *int64 `json:"TemplatePauseDuration,omitnil,omitempty" name:"TemplatePauseDuration"`

	// Main account that creates the experiment
	TemplateOwnerUin *string `json:"TemplateOwnerUin,omitnil,omitempty" name:"TemplateOwnerUin"`

	// Region ID
	TemplateRegionId *int64 `json:"TemplateRegionId,omitnil,omitempty" name:"TemplateRegionId"`

	// Action group
	TemplateGroups []*TemplateGroup `json:"TemplateGroups,omitnil,omitempty" name:"TemplateGroups"`

	// Monitoring metrics
	TemplateMonitors []*TemplateMonitor `json:"TemplateMonitors,omitnil,omitempty" name:"TemplateMonitors"`

	// Guardrail monitoring
	// Note: This field may return null, indicating that no valid values can be obtained.
	TemplatePolicy *TemplatePolicy `json:"TemplatePolicy,omitnil,omitempty" name:"TemplatePolicy"`

	// Tag list
	// Note: This field may return null, indicating that no valid values can be obtained.
	Tags []*TagWithDescribe `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Template library source. 0: self-built; 1: recommended by experts.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TemplateSource *int64 `json:"TemplateSource,omitnil,omitempty" name:"TemplateSource"`

	// Application information on Application Performance Monitoring
	// Note: This field may return null, indicating that no valid values can be obtained.
	ApmServiceList []*ApmServiceInfo `json:"ApmServiceList,omitnil,omitempty" name:"ApmServiceList"`

	// Alarm metrics
	// Note: This field may return null, indicating that no valid values can be obtained.
	AlarmPolicy []*string `json:"AlarmPolicy,omitnil,omitempty" name:"AlarmPolicy"`

	// Guardrail processing method. 1: roll back sequentially; 2: pause experiment.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PolicyDealType *int64 `json:"PolicyDealType,omitnil,omitempty" name:"PolicyDealType"`
}

type TemplateGroup struct {
	// Template library action ID
	TemplateGroupId *int64 `json:"TemplateGroupId,omitnil,omitempty" name:"TemplateGroupId"`

	// List of actions in the template library action group
	TemplateGroupActions []*TemplateGroupAction `json:"TemplateGroupActions,omitnil,omitempty" name:"TemplateGroupActions"`

	// Group name
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`

	// Group description
	// Note: This field may return null, indicating that no valid values can be obtained.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Group order
	Order *int64 `json:"Order,omitnil,omitempty" name:"Order"`

	// Execution mode. 1: sequential execution; 2: execution by stage.
	Mode *int64 `json:"Mode,omitnil,omitempty" name:"Mode"`

	// Object type ID
	ObjectTypeId *int64 `json:"ObjectTypeId,omitnil,omitempty" name:"ObjectTypeId"`

	// Group creation time
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Group update time
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type TemplateGroupAction struct {
	// Template library group action ID
	TemplateGroupActionId *int64 `json:"TemplateGroupActionId,omitnil,omitempty" name:"TemplateGroupActionId"`

	// Action ID
	ActionId *int64 `json:"ActionId,omitnil,omitempty" name:"ActionId"`

	// Order of actions in the group
	Order *int64 `json:"Order,omitnil,omitempty" name:"Order"`

	// General configurations of actions in the group
	// Note: This field may return null, indicating that no valid values can be obtained.
	GeneralConfiguration *string `json:"GeneralConfiguration,omitnil,omitempty" name:"GeneralConfiguration"`

	// Custom configurations of actions in the group
	// Note: This field may return null, indicating that no valid values can be obtained.
	CustomConfiguration *string `json:"CustomConfiguration,omitnil,omitempty" name:"CustomConfiguration"`

	// Action group creation time
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Action group update time
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// Action name
	ActionTitle *string `json:"ActionTitle,omitnil,omitempty" name:"ActionTitle"`

	// Random ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	RandomId *int64 `json:"RandomId,omitnil,omitempty" name:"RandomId"`

	// Recovery action ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	RecoverId *int64 `json:"RecoverId,omitnil,omitempty" name:"RecoverId"`

	// Executed action ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	ExecuteId *int64 `json:"ExecuteId,omitnil,omitempty" name:"ExecuteId"`

	// Called API type. 0: tat; 1: cloud API.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ActionApiType *int64 `json:"ActionApiType,omitnil,omitempty" name:"ActionApiType"`

	// 1: fault; 2: recovery.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ActionAttribute *int64 `json:"ActionAttribute,omitnil,omitempty" name:"ActionAttribute"`

	// Action type: platform and custom
	// Note: This field may return null, indicating that no valid values can be obtained.
	ActionType *string `json:"ActionType,omitnil,omitempty" name:"ActionType"`

	// Action risk level. 1: low-risk; 2: medium-risk; 3: high-risk.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ActionRisk *string `json:"ActionRisk,omitnil,omitempty" name:"ActionRisk"`
}

type TemplateListItem struct {
	// Template library ID
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// Template library name
	TemplateTitle *string `json:"TemplateTitle,omitnil,omitempty" name:"TemplateTitle"`

	// Template library description
	TemplateDescription *string `json:"TemplateDescription,omitnil,omitempty" name:"TemplateDescription"`

	// Template library tag
	// Note: This field may return null, indicating that no valid values can be obtained.
	TemplateTag *string `json:"TemplateTag,omitnil,omitempty" name:"TemplateTag"`

	// Template library status. 1: in use; 2: not in use.
	TemplateIsUsed *int64 `json:"TemplateIsUsed,omitnil,omitempty" name:"TemplateIsUsed"`

	// Template library creation time
	TemplateCreateTime *string `json:"TemplateCreateTime,omitnil,omitempty" name:"TemplateCreateTime"`

	// Template library update time
	TemplateUpdateTime *string `json:"TemplateUpdateTime,omitnil,omitempty" name:"TemplateUpdateTime"`

	// Number of tasks associated with the template library
	TemplateUsedNum *int64 `json:"TemplateUsedNum,omitnil,omitempty" name:"TemplateUsedNum"`

	// Template library source. 0: self-built; 1: recommended by experts.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TemplateSource *int64 `json:"TemplateSource,omitnil,omitempty" name:"TemplateSource"`
}

type TemplateMonitor struct {
	// pk
	MonitorId *int64 `json:"MonitorId,omitnil,omitempty" name:"MonitorId"`

	// Monitoring metric ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	MetricId *int64 `json:"MetricId,omitnil,omitempty" name:"MetricId"`

	// Object type ID of the monitoring metric
	ObjectTypeId *int64 `json:"ObjectTypeId,omitnil,omitempty" name:"ObjectTypeId"`

	// Metric name
	MetricName *string `json:"MetricName,omitnil,omitempty" name:"MetricName"`

	// Metric in Chinese
	// Note: This field may return null, indicating that no valid values can be obtained.
	MetricChineseName *string `json:"MetricChineseName,omitnil,omitempty" name:"MetricChineseName"`
}

type TemplatePolicy struct {
	// Protection policy ID list
	TemplatePolicyIdList []*string `json:"TemplatePolicyIdList,omitnil,omitempty" name:"TemplatePolicyIdList"`

	// Policy rules
	TemplatePolicyRule *string `json:"TemplatePolicyRule,omitnil,omitempty" name:"TemplatePolicyRule"`

	// Process method when the guardrail policy takes effect. 1: sequential execution, 2: pausing.
	TemplatePolicyDealType *int64 `json:"TemplatePolicyDealType,omitnil,omitempty" name:"TemplatePolicyDealType"`
}

// Predefined struct for user
type TriggerPolicyRequestParams struct {
	// Chaos engineering experiment ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Triggering content
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// Triggering type. 0: trigger; 1: recovery.
	TriggerType *int64 `json:"TriggerType,omitnil,omitempty" name:"TriggerType"`
}

type TriggerPolicyRequest struct {
	*tchttp.BaseRequest
	
	// Chaos engineering experiment ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Triggering content
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// Triggering type. 0: trigger; 1: recovery.
	TriggerType *int64 `json:"TriggerType,omitnil,omitempty" name:"TriggerType"`
}

func (r *TriggerPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TriggerPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "Name")
	delete(f, "Content")
	delete(f, "TriggerType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TriggerPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TriggerPolicyResponseParams struct {
	// Experiment ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Whether triggering is successful
	Success *bool `json:"Success,omitnil,omitempty" name:"Success"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type TriggerPolicyResponse struct {
	*tchttp.BaseResponse
	Response *TriggerPolicyResponseParams `json:"Response"`
}

func (r *TriggerPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TriggerPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}