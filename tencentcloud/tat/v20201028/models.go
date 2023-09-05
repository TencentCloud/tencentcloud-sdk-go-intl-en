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

package v20201028

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

type AutomationAgentInfo struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Agent version.
	Version *string `json:"Version,omitnil" name:"Version"`

	// Last heartbeat time
	LastHeartbeatTime *string `json:"LastHeartbeatTime,omitnil" name:"LastHeartbeatTime"`

	// Agent status. Valid values:
	// <li> `Online`
	// <li> `Offline`
	AgentStatus *string `json:"AgentStatus,omitnil" name:"AgentStatus"`

	// Agent runtime environment. Valid values:
	// <li> `Linux`: Linux instance
	// <li> `Windows`: Windows instance
	Environment *string `json:"Environment,omitnil" name:"Environment"`

	// Features supported by the TAT agent.
	SupportFeatures []*string `json:"SupportFeatures,omitnil" name:"SupportFeatures"`
}

// Predefined struct for user
type CancelInvocationRequestParams struct {
	// Execution activity ID
	InvocationId *string `json:"InvocationId,omitnil" name:"InvocationId"`

	// Instance ID list. A maximum of 100 IDs are allowed. Supported instance types:
	// <li> `CVM`
	// <li> `LIGHTHOUSE`
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`
}

type CancelInvocationRequest struct {
	*tchttp.BaseRequest
	
	// Execution activity ID
	InvocationId *string `json:"InvocationId,omitnil" name:"InvocationId"`

	// Instance ID list. A maximum of 100 IDs are allowed. Supported instance types:
	// <li> `CVM`
	// <li> `LIGHTHOUSE`
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`
}

func (r *CancelInvocationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelInvocationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InvocationId")
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CancelInvocationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelInvocationResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CancelInvocationResponse struct {
	*tchttp.BaseResponse
	Response *CancelInvocationResponseParams `json:"Response"`
}

func (r *CancelInvocationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelInvocationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Command struct {
	// Command ID.
	CommandId *string `json:"CommandId,omitnil" name:"CommandId"`

	// Command name.
	CommandName *string `json:"CommandName,omitnil" name:"CommandName"`

	// Command description.
	Description *string `json:"Description,omitnil" name:"Description"`

	// Base64-encoded command.
	Content *string `json:"Content,omitnil" name:"Content"`

	// Command type.
	CommandType *string `json:"CommandType,omitnil" name:"CommandType"`

	// Command execution path.
	WorkingDirectory *string `json:"WorkingDirectory,omitnil" name:"WorkingDirectory"`

	// Command timeout period.
	Timeout *uint64 `json:"Timeout,omitnil" name:"Timeout"`

	// Command creation time.
	CreatedTime *string `json:"CreatedTime,omitnil" name:"CreatedTime"`

	// Command update time.
	UpdatedTime *string `json:"UpdatedTime,omitnil" name:"UpdatedTime"`

	// Whether to enable the custom parameter feature.
	EnableParameter *bool `json:"EnableParameter,omitnil" name:"EnableParameter"`

	// Default custom parameter value.
	DefaultParameters *string `json:"DefaultParameters,omitnil" name:"DefaultParameters"`

	// Formatted description of the command. This parameter is an empty string for user commands and contains values for public commands.
	FormattedDescription *string `json:"FormattedDescription,omitnil" name:"FormattedDescription"`

	// Command creator. `TAT` indicates a public command and `USER` indicates a personal command.
	CreatedBy *string `json:"CreatedBy,omitnil" name:"CreatedBy"`

	// The list of tags bound to the command.
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`

	// The user who executes the command on the instance.
	Username *string `json:"Username,omitnil" name:"Username"`

	// The COS bucket URL for uploading logs.
	OutputCOSBucketUrl *string `json:"OutputCOSBucketUrl,omitnil" name:"OutputCOSBucketUrl"`

	// The COS bucket directory where the logs are saved.
	OutputCOSKeyPrefix *string `json:"OutputCOSKeyPrefix,omitnil" name:"OutputCOSKeyPrefix"`
}

type CommandDocument struct {
	// Base64-encoded command.
	Content *string `json:"Content,omitnil" name:"Content"`

	// Command type.
	CommandType *string `json:"CommandType,omitnil" name:"CommandType"`

	// Timeout period.
	Timeout *uint64 `json:"Timeout,omitnil" name:"Timeout"`

	// Execution path.
	WorkingDirectory *string `json:"WorkingDirectory,omitnil" name:"WorkingDirectory"`

	// The user who executes the command.
	Username *string `json:"Username,omitnil" name:"Username"`

	// URL of the COS bucket to store the output
	OutputCOSBucketUrl *string `json:"OutputCOSBucketUrl,omitnil" name:"OutputCOSBucketUrl"`

	// Prefix of the output file name 
	OutputCOSKeyPrefix *string `json:"OutputCOSKeyPrefix,omitnil" name:"OutputCOSKeyPrefix"`
}

// Predefined struct for user
type CreateCommandRequestParams struct {
	// Command name. The name can be up to 60 bytes, and contain [a-z], [A-Z], [0-9] and [_-.].
	CommandName *string `json:"CommandName,omitnil" name:"CommandName"`

	// Base64-encoded command. The maximum length is 64 KB.
	Content *string `json:"Content,omitnil" name:"Content"`

	// Command description. The maximum length is 120 characters.
	Description *string `json:"Description,omitnil" name:"Description"`

	// Command type. `SHELL` and `POWERSHELL` are supported. The default value is `SHELL`.
	CommandType *string `json:"CommandType,omitnil" name:"CommandType"`

	// Command execution path. The default value is /root for `SHELL` commands and C:\Program Files\qcloud\tat_agent\workdir for `POWERSHELL` commands.
	WorkingDirectory *string `json:"WorkingDirectory,omitnil" name:"WorkingDirectory"`

	// Command timeout period. Default value: 60 seconds. Value range: [1, 86400].
	Timeout *uint64 `json:"Timeout,omitnil" name:"Timeout"`

	// Whether to enable the custom parameter feature.
	// This cannot be modified once created.
	// Default value: `false`.
	EnableParameter *bool `json:"EnableParameter,omitnil" name:"EnableParameter"`

	// The default value of the custom parameter value when it is enabled. The field type is JSON encoded string. For example, {\"varA\": \"222\"}.
	// `key` is the name of the custom parameter and `value` is the default value. Both `key` and `value` are strings.
	// If no parameter value is provided in the `InvokeCommand` API, the default value is used.
	// Up to 20 custom parameters are supported.
	// The name of the custom parameter cannot exceed 64 characters and can contain [a-z], [A-Z], [0-9] and [-_].
	DefaultParameters *string `json:"DefaultParameters,omitnil" name:"DefaultParameters"`

	// Tags bound to the command. At most 10 tags are allowed.
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`

	// The username used to execute the command on the CVM or Lighthouse instance.
	// The principle of least privilege is the best practice for permission management. We recommend you execute TAT commands as a general user. By default, the root user is used to execute commands on Linux and the System user is used on Windows.
	Username *string `json:"Username,omitnil" name:"Username"`

	// The COS bucket URL for uploading logs. The URL must start with `https`, such as `https://BucketName-123454321.cos.ap-beijing.myqcloud.com`.
	OutputCOSBucketUrl *string `json:"OutputCOSBucketUrl,omitnil" name:"OutputCOSBucketUrl"`

	// The COS bucket directory where the logs are saved. Check below for the rules of the directory name. 
	// 1. It must be a combination of number, letters, and visible characters. Up to 60 characters are allowed.
	// 2. Use a slash (/) to create a subdirectory.
	// 3. Consecutive dots (.) and slashes (/) are not allowed. It can not start with a slash (/). 
	OutputCOSKeyPrefix *string `json:"OutputCOSKeyPrefix,omitnil" name:"OutputCOSKeyPrefix"`
}

type CreateCommandRequest struct {
	*tchttp.BaseRequest
	
	// Command name. The name can be up to 60 bytes, and contain [a-z], [A-Z], [0-9] and [_-.].
	CommandName *string `json:"CommandName,omitnil" name:"CommandName"`

	// Base64-encoded command. The maximum length is 64 KB.
	Content *string `json:"Content,omitnil" name:"Content"`

	// Command description. The maximum length is 120 characters.
	Description *string `json:"Description,omitnil" name:"Description"`

	// Command type. `SHELL` and `POWERSHELL` are supported. The default value is `SHELL`.
	CommandType *string `json:"CommandType,omitnil" name:"CommandType"`

	// Command execution path. The default value is /root for `SHELL` commands and C:\Program Files\qcloud\tat_agent\workdir for `POWERSHELL` commands.
	WorkingDirectory *string `json:"WorkingDirectory,omitnil" name:"WorkingDirectory"`

	// Command timeout period. Default value: 60 seconds. Value range: [1, 86400].
	Timeout *uint64 `json:"Timeout,omitnil" name:"Timeout"`

	// Whether to enable the custom parameter feature.
	// This cannot be modified once created.
	// Default value: `false`.
	EnableParameter *bool `json:"EnableParameter,omitnil" name:"EnableParameter"`

	// The default value of the custom parameter value when it is enabled. The field type is JSON encoded string. For example, {\"varA\": \"222\"}.
	// `key` is the name of the custom parameter and `value` is the default value. Both `key` and `value` are strings.
	// If no parameter value is provided in the `InvokeCommand` API, the default value is used.
	// Up to 20 custom parameters are supported.
	// The name of the custom parameter cannot exceed 64 characters and can contain [a-z], [A-Z], [0-9] and [-_].
	DefaultParameters *string `json:"DefaultParameters,omitnil" name:"DefaultParameters"`

	// Tags bound to the command. At most 10 tags are allowed.
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`

	// The username used to execute the command on the CVM or Lighthouse instance.
	// The principle of least privilege is the best practice for permission management. We recommend you execute TAT commands as a general user. By default, the root user is used to execute commands on Linux and the System user is used on Windows.
	Username *string `json:"Username,omitnil" name:"Username"`

	// The COS bucket URL for uploading logs. The URL must start with `https`, such as `https://BucketName-123454321.cos.ap-beijing.myqcloud.com`.
	OutputCOSBucketUrl *string `json:"OutputCOSBucketUrl,omitnil" name:"OutputCOSBucketUrl"`

	// The COS bucket directory where the logs are saved. Check below for the rules of the directory name. 
	// 1. It must be a combination of number, letters, and visible characters. Up to 60 characters are allowed.
	// 2. Use a slash (/) to create a subdirectory.
	// 3. Consecutive dots (.) and slashes (/) are not allowed. It can not start with a slash (/). 
	OutputCOSKeyPrefix *string `json:"OutputCOSKeyPrefix,omitnil" name:"OutputCOSKeyPrefix"`
}

func (r *CreateCommandRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCommandRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CommandName")
	delete(f, "Content")
	delete(f, "Description")
	delete(f, "CommandType")
	delete(f, "WorkingDirectory")
	delete(f, "Timeout")
	delete(f, "EnableParameter")
	delete(f, "DefaultParameters")
	delete(f, "Tags")
	delete(f, "Username")
	delete(f, "OutputCOSBucketUrl")
	delete(f, "OutputCOSKeyPrefix")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCommandRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCommandResponseParams struct {
	// Command ID.
	CommandId *string `json:"CommandId,omitnil" name:"CommandId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateCommandResponse struct {
	*tchttp.BaseResponse
	Response *CreateCommandResponseParams `json:"Response"`
}

func (r *CreateCommandResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCommandResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInvokerRequestParams struct {
	// Invoker name.
	Name *string `json:"Name,omitnil" name:"Name"`

	// Invoker type. It can only be `SCHEDULE` (recurring invokers).
	Type *string `json:"Type,omitnil" name:"Type"`

	// Remote command ID.
	CommandId *string `json:"CommandId,omitnil" name:"CommandId"`

	// ID of the instance bound to the trigger. Up to 100 IDs are allowed.
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`

	// The user who executes the command.
	Username *string `json:"Username,omitnil" name:"Username"`

	// Custom parameters of the command.
	Parameters *string `json:"Parameters,omitnil" name:"Parameters"`

	// Settings required for a recurring invoker.
	ScheduleSettings *ScheduleSettings `json:"ScheduleSettings,omitnil" name:"ScheduleSettings"`
}

type CreateInvokerRequest struct {
	*tchttp.BaseRequest
	
	// Invoker name.
	Name *string `json:"Name,omitnil" name:"Name"`

	// Invoker type. It can only be `SCHEDULE` (recurring invokers).
	Type *string `json:"Type,omitnil" name:"Type"`

	// Remote command ID.
	CommandId *string `json:"CommandId,omitnil" name:"CommandId"`

	// ID of the instance bound to the trigger. Up to 100 IDs are allowed.
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`

	// The user who executes the command.
	Username *string `json:"Username,omitnil" name:"Username"`

	// Custom parameters of the command.
	Parameters *string `json:"Parameters,omitnil" name:"Parameters"`

	// Settings required for a recurring invoker.
	ScheduleSettings *ScheduleSettings `json:"ScheduleSettings,omitnil" name:"ScheduleSettings"`
}

func (r *CreateInvokerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInvokerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Type")
	delete(f, "CommandId")
	delete(f, "InstanceIds")
	delete(f, "Username")
	delete(f, "Parameters")
	delete(f, "ScheduleSettings")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateInvokerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInvokerResponseParams struct {
	// Invoker ID.
	InvokerId *string `json:"InvokerId,omitnil" name:"InvokerId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateInvokerResponse struct {
	*tchttp.BaseResponse
	Response *CreateInvokerResponseParams `json:"Response"`
}

func (r *CreateInvokerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInvokerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCommandRequestParams struct {
	// ID of the command to be deleted.
	CommandId *string `json:"CommandId,omitnil" name:"CommandId"`
}

type DeleteCommandRequest struct {
	*tchttp.BaseRequest
	
	// ID of the command to be deleted.
	CommandId *string `json:"CommandId,omitnil" name:"CommandId"`
}

func (r *DeleteCommandRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCommandRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CommandId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCommandRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCommandResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteCommandResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCommandResponseParams `json:"Response"`
}

func (r *DeleteCommandResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCommandResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteInvokerRequestParams struct {
	// ID of the invoker to be deleted.
	InvokerId *string `json:"InvokerId,omitnil" name:"InvokerId"`
}

type DeleteInvokerRequest struct {
	*tchttp.BaseRequest
	
	// ID of the invoker to be deleted.
	InvokerId *string `json:"InvokerId,omitnil" name:"InvokerId"`
}

func (r *DeleteInvokerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteInvokerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InvokerId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteInvokerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteInvokerResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteInvokerResponse struct {
	*tchttp.BaseResponse
	Response *DeleteInvokerResponseParams `json:"Response"`
}

func (r *DeleteInvokerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteInvokerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAutomationAgentStatusRequestParams struct {
	// List of instance IDs for the query.
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`

	// Filter conditions.<br> <li>`agent-status` - String - Required: No - (Filter condition) Filter by agent status. Valid values: `Online`, `Offline`.<br> <li> `environment` - String - Required: No - (Filter condition) Filter by the agent environment. Valid value: `Linux`.<br> <li> `instance-id` - String - Required: No - (Filter condition) Filter by the instance ID. <br>Up to 10 `Filters` allowed in one request. For each filter, five `Filter.Values` can be specified. `InstanceIds` and `Filters` cannot be specified at the same time.
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// Number of returned results. It defaults to `20`. The maximum is 100. For more information on `Limit`, see the relevant section in the API [Overview](https://intl.cloud.tencent.com/document/api/213/15688?from_cn_redirect=1).
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Offset. The default value is `0`. For more information on `Offset`, see the relevant section in API [Introduction](https://intl.cloud.tencent.com/document/api/213/15688?from_cn_redirect=1).
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`
}

type DescribeAutomationAgentStatusRequest struct {
	*tchttp.BaseRequest
	
	// List of instance IDs for the query.
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`

	// Filter conditions.<br> <li>`agent-status` - String - Required: No - (Filter condition) Filter by agent status. Valid values: `Online`, `Offline`.<br> <li> `environment` - String - Required: No - (Filter condition) Filter by the agent environment. Valid value: `Linux`.<br> <li> `instance-id` - String - Required: No - (Filter condition) Filter by the instance ID. <br>Up to 10 `Filters` allowed in one request. For each filter, five `Filter.Values` can be specified. `InstanceIds` and `Filters` cannot be specified at the same time.
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// Number of returned results. It defaults to `20`. The maximum is 100. For more information on `Limit`, see the relevant section in the API [Overview](https://intl.cloud.tencent.com/document/api/213/15688?from_cn_redirect=1).
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Offset. The default value is `0`. For more information on `Offset`, see the relevant section in API [Introduction](https://intl.cloud.tencent.com/document/api/213/15688?from_cn_redirect=1).
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`
}

func (r *DescribeAutomationAgentStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAutomationAgentStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAutomationAgentStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAutomationAgentStatusResponseParams struct {
	// Agent information list.
	AutomationAgentSet []*AutomationAgentInfo `json:"AutomationAgentSet,omitnil" name:"AutomationAgentSet"`

	// Total number of matching agents.
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeAutomationAgentStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAutomationAgentStatusResponseParams `json:"Response"`
}

func (r *DescribeAutomationAgentStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAutomationAgentStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCommandsRequestParams struct {
	// List of command IDs. Up to 100 IDs are allowed for each request. `CommandIds` and `Filters` cannot be specified at the same time.
	CommandIds []*string `json:"CommandIds,omitnil" name:"CommandIds"`

	// Filter conditions.
	// <li> `command-id` - String - Required: No - (Filter condition) Filter by the command ID.
	// <li> `command-name` - String - Required: No - (Filter condition) Filter by the command name.
	// <li> `command-type` - String - Required: No - (Filter condition) Filter by the command type. Valid values: `SHELL` or `POWERSHELL`.
	// <li> `created-by` - String - Required: No - (Filter condition) Filter by the creator. Valid values: `TAT` (public commands) or `USER` (custom commands).
	// <li> `tag-key` - String - Required: No - (Filter condition) Filter by the tag key.</li>
	// <li> `tag-value` - String - Required: No - (Filter condition) Filter by the tag value.</li>
	// <li> `tag:tag-key` - String - Required: No - (Filter) Filter by the tag key-value pair. The tag-key should be replaced with a specified tag key. For detailed usage, see sample 4.</li>
	// 
	// Up to 10 `Filters` are allowed in one request. Each filter can have up to 5 `Filter.Values`. `CommandIds` and `Filters` cannot be specified at the same time.
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// Number of returned results. It defaults to `20`. The maximum is 100. For more information on `Limit`, see the relevant section in the API [Overview](https://intl.cloud.tencent.com/document/api/213/15688?from_cn_redirect=1).
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Offset. The default value is `0`. For more information on `Offset`, see the relevant section in API [Introduction](https://intl.cloud.tencent.com/document/api/213/15688?from_cn_redirect=1).
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`
}

type DescribeCommandsRequest struct {
	*tchttp.BaseRequest
	
	// List of command IDs. Up to 100 IDs are allowed for each request. `CommandIds` and `Filters` cannot be specified at the same time.
	CommandIds []*string `json:"CommandIds,omitnil" name:"CommandIds"`

	// Filter conditions.
	// <li> `command-id` - String - Required: No - (Filter condition) Filter by the command ID.
	// <li> `command-name` - String - Required: No - (Filter condition) Filter by the command name.
	// <li> `command-type` - String - Required: No - (Filter condition) Filter by the command type. Valid values: `SHELL` or `POWERSHELL`.
	// <li> `created-by` - String - Required: No - (Filter condition) Filter by the creator. Valid values: `TAT` (public commands) or `USER` (custom commands).
	// <li> `tag-key` - String - Required: No - (Filter condition) Filter by the tag key.</li>
	// <li> `tag-value` - String - Required: No - (Filter condition) Filter by the tag value.</li>
	// <li> `tag:tag-key` - String - Required: No - (Filter) Filter by the tag key-value pair. The tag-key should be replaced with a specified tag key. For detailed usage, see sample 4.</li>
	// 
	// Up to 10 `Filters` are allowed in one request. Each filter can have up to 5 `Filter.Values`. `CommandIds` and `Filters` cannot be specified at the same time.
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// Number of returned results. It defaults to `20`. The maximum is 100. For more information on `Limit`, see the relevant section in the API [Overview](https://intl.cloud.tencent.com/document/api/213/15688?from_cn_redirect=1).
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Offset. The default value is `0`. For more information on `Offset`, see the relevant section in API [Introduction](https://intl.cloud.tencent.com/document/api/213/15688?from_cn_redirect=1).
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`
}

func (r *DescribeCommandsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCommandsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CommandIds")
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCommandsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCommandsResponseParams struct {
	// Total number of matching commands.
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// List of command details.
	CommandSet []*Command `json:"CommandSet,omitnil" name:"CommandSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeCommandsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCommandsResponseParams `json:"Response"`
}

func (r *DescribeCommandsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCommandsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInvocationTasksRequestParams struct {
	// List of execution task IDs. Up to 100 IDs are allowed for each request. `InvocationTaskIds` and `Filters` cannot be specified at the same time.
	InvocationTaskIds []*string `json:"InvocationTaskIds,omitnil" name:"InvocationTaskIds"`

	// Filter conditions.<br> <li> `invocation-id` - String - Required: No - (Filter condition) Filter by the execution activity ID.<br> <li> `invocation-task-id` - String - Required: No - (Filter condition) Filter by the execution task ID.<br> <li> `instance-id` - String - Required: No - (Filter condition) Filter by the instance ID. <br> <li> `command-id` - String - Required: No - (Filter condition) Filter by the command ID. <br>Up to 10 `Filters` are allowed for each request. Each filter can have up to five `Filter.Values`. `InvocationTaskIds` and `Filters` cannot be specified at the same time.
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// Number of returned results. It defaults to `20`. The maximum is 100. For more information on `Limit`, see the relevant section in the API [Overview](https://intl.cloud.tencent.com/document/api/213/15688?from_cn_redirect=1).
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Offset. The default value is `0`. For more information on `Offset`, see the relevant section in API [Introduction](https://intl.cloud.tencent.com/document/api/213/15688?from_cn_redirect=1).
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// Whether to hide the output. Valid values:<br><li>`True` (default): Hide the output <br><li>`False`: Show the output <br>
	HideOutput *bool `json:"HideOutput,omitnil" name:"HideOutput"`
}

type DescribeInvocationTasksRequest struct {
	*tchttp.BaseRequest
	
	// List of execution task IDs. Up to 100 IDs are allowed for each request. `InvocationTaskIds` and `Filters` cannot be specified at the same time.
	InvocationTaskIds []*string `json:"InvocationTaskIds,omitnil" name:"InvocationTaskIds"`

	// Filter conditions.<br> <li> `invocation-id` - String - Required: No - (Filter condition) Filter by the execution activity ID.<br> <li> `invocation-task-id` - String - Required: No - (Filter condition) Filter by the execution task ID.<br> <li> `instance-id` - String - Required: No - (Filter condition) Filter by the instance ID. <br> <li> `command-id` - String - Required: No - (Filter condition) Filter by the command ID. <br>Up to 10 `Filters` are allowed for each request. Each filter can have up to five `Filter.Values`. `InvocationTaskIds` and `Filters` cannot be specified at the same time.
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// Number of returned results. It defaults to `20`. The maximum is 100. For more information on `Limit`, see the relevant section in the API [Overview](https://intl.cloud.tencent.com/document/api/213/15688?from_cn_redirect=1).
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Offset. The default value is `0`. For more information on `Offset`, see the relevant section in API [Introduction](https://intl.cloud.tencent.com/document/api/213/15688?from_cn_redirect=1).
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// Whether to hide the output. Valid values:<br><li>`True` (default): Hide the output <br><li>`False`: Show the output <br>
	HideOutput *bool `json:"HideOutput,omitnil" name:"HideOutput"`
}

func (r *DescribeInvocationTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInvocationTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InvocationTaskIds")
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "HideOutput")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInvocationTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInvocationTasksResponseParams struct {
	// Total number of matching execution tasks.
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// List of execution tasks.
	InvocationTaskSet []*InvocationTask `json:"InvocationTaskSet,omitnil" name:"InvocationTaskSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeInvocationTasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInvocationTasksResponseParams `json:"Response"`
}

func (r *DescribeInvocationTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInvocationTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInvocationsRequestParams struct {
	// List of execution activity IDs. Up to 100 IDs are allowed for each request. `InvocationIds` and `Filters` cannot be specified at the same time.
	InvocationIds []*string `json:"InvocationIds,omitnil" name:"InvocationIds"`

	// Filter conditions.<br> <li> `invocation-id` - String - Required: No - (Filter condition) Filter by the execution activity ID.<br> 
	// <li> `command-id` - String - Required: No - (Filter condition) Filter by the command ID. 
	// <li> `command-created-by` - String - Required: No - (Filter condition) Filter by the command type. Valid values: `TAT` (public commands) or `USER` (custom commands).
	// <li> `instance-kind` - String - Required: No - (Filter condition) Filter by the instance type. Valid values: `CVM` or `LIGHTHOUSE`. 
	// <br>Up to 10 `Filters` are allowed for each request. Each filter can have up to five `Filter.Values`. `InvocationIds` and `Filters` cannot be specified at the same time.
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// Number of returned results. It defaults to `20`. The maximum is 100. For more information on `Limit`, see the relevant section in the API [Overview](https://intl.cloud.tencent.com/document/api/213/15688?from_cn_redirect=1).
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Offset. The default value is `0`. For more information on `Offset`, see the relevant section in API [Introduction](https://intl.cloud.tencent.com/document/api/213/15688?from_cn_redirect=1).
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`
}

type DescribeInvocationsRequest struct {
	*tchttp.BaseRequest
	
	// List of execution activity IDs. Up to 100 IDs are allowed for each request. `InvocationIds` and `Filters` cannot be specified at the same time.
	InvocationIds []*string `json:"InvocationIds,omitnil" name:"InvocationIds"`

	// Filter conditions.<br> <li> `invocation-id` - String - Required: No - (Filter condition) Filter by the execution activity ID.<br> 
	// <li> `command-id` - String - Required: No - (Filter condition) Filter by the command ID. 
	// <li> `command-created-by` - String - Required: No - (Filter condition) Filter by the command type. Valid values: `TAT` (public commands) or `USER` (custom commands).
	// <li> `instance-kind` - String - Required: No - (Filter condition) Filter by the instance type. Valid values: `CVM` or `LIGHTHOUSE`. 
	// <br>Up to 10 `Filters` are allowed for each request. Each filter can have up to five `Filter.Values`. `InvocationIds` and `Filters` cannot be specified at the same time.
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// Number of returned results. It defaults to `20`. The maximum is 100. For more information on `Limit`, see the relevant section in the API [Overview](https://intl.cloud.tencent.com/document/api/213/15688?from_cn_redirect=1).
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Offset. The default value is `0`. For more information on `Offset`, see the relevant section in API [Introduction](https://intl.cloud.tencent.com/document/api/213/15688?from_cn_redirect=1).
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`
}

func (r *DescribeInvocationsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInvocationsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InvocationIds")
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInvocationsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInvocationsResponseParams struct {
	// Total number of matching execution activities.
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// List of execution activities.
	InvocationSet []*Invocation `json:"InvocationSet,omitnil" name:"InvocationSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeInvocationsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInvocationsResponseParams `json:"Response"`
}

func (r *DescribeInvocationsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInvocationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInvokerRecordsRequestParams struct {
	// List of invoker IDs. Up to 100 IDs are allowed.
	InvokerIds []*string `json:"InvokerIds,omitnil" name:"InvokerIds"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`
}

type DescribeInvokerRecordsRequest struct {
	*tchttp.BaseRequest
	
	// List of invoker IDs. Up to 100 IDs are allowed.
	InvokerIds []*string `json:"InvokerIds,omitnil" name:"InvokerIds"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`
}

func (r *DescribeInvokerRecordsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInvokerRecordsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InvokerIds")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInvokerRecordsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInvokerRecordsResponseParams struct {
	// Number of matching records.
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// Execution history of an invoker.
	InvokerRecordSet []*InvokerRecord `json:"InvokerRecordSet,omitnil" name:"InvokerRecordSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeInvokerRecordsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInvokerRecordsResponseParams `json:"Response"`
}

func (r *DescribeInvokerRecordsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInvokerRecordsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInvokersRequestParams struct {
	// List of invoker IDs.
	InvokerIds []*string `json:"InvokerIds,omitnil" name:"InvokerIds"`

	// Filter conditions:
	// 
	// <li> `invoker-id` - String - Required: No - (Filter condition) Filter by the invoker ID.
	// <li> `command-id` - String - Required: No - (Filter condition) Filter by the command ID.
	// <li> `type` - String - Required: No - (Filter condition) Filter by the invoker type.
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`
}

type DescribeInvokersRequest struct {
	*tchttp.BaseRequest
	
	// List of invoker IDs.
	InvokerIds []*string `json:"InvokerIds,omitnil" name:"InvokerIds"`

	// Filter conditions:
	// 
	// <li> `invoker-id` - String - Required: No - (Filter condition) Filter by the invoker ID.
	// <li> `command-id` - String - Required: No - (Filter condition) Filter by the command ID.
	// <li> `type` - String - Required: No - (Filter condition) Filter by the invoker type.
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`
}

func (r *DescribeInvokersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInvokersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InvokerIds")
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInvokersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInvokersResponseParams struct {
	// Number of matching invokers.
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// Invoker information.
	InvokerSet []*Invoker `json:"InvokerSet,omitnil" name:"InvokerSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeInvokersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInvokersResponseParams `json:"Response"`
}

func (r *DescribeInvokersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInvokersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRegionsRequestParams struct {

}

type DescribeRegionsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeRegionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRegionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRegionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRegionsResponseParams struct {
	// Number of regions
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// Region information list
	RegionSet []*RegionInfo `json:"RegionSet,omitnil" name:"RegionSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeRegionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRegionsResponseParams `json:"Response"`
}

func (r *DescribeRegionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRegionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableInvokerRequestParams struct {
	// ID of the invoker to be disabled.
	InvokerId *string `json:"InvokerId,omitnil" name:"InvokerId"`
}

type DisableInvokerRequest struct {
	*tchttp.BaseRequest
	
	// ID of the invoker to be disabled.
	InvokerId *string `json:"InvokerId,omitnil" name:"InvokerId"`
}

func (r *DisableInvokerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableInvokerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InvokerId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisableInvokerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableInvokerResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DisableInvokerResponse struct {
	*tchttp.BaseResponse
	Response *DisableInvokerResponseParams `json:"Response"`
}

func (r *DisableInvokerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableInvokerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableInvokerRequestParams struct {
	// ID of the invoker to be enabled.
	InvokerId *string `json:"InvokerId,omitnil" name:"InvokerId"`
}

type EnableInvokerRequest struct {
	*tchttp.BaseRequest
	
	// ID of the invoker to be enabled.
	InvokerId *string `json:"InvokerId,omitnil" name:"InvokerId"`
}

func (r *EnableInvokerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableInvokerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InvokerId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EnableInvokerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableInvokerResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type EnableInvokerResponse struct {
	*tchttp.BaseResponse
	Response *EnableInvokerResponseParams `json:"Response"`
}

func (r *EnableInvokerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableInvokerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filter struct {
	// Field to be filtered.
	Name *string `json:"Name,omitnil" name:"Name"`

	// Filter values of the field.
	Values []*string `json:"Values,omitnil" name:"Values"`
}

type Invocation struct {
	// Execution activity ID.
	InvocationId *string `json:"InvocationId,omitnil" name:"InvocationId"`

	// Command ID.
	CommandId *string `json:"CommandId,omitnil" name:"CommandId"`

	// Execution task status. Valid values:
	// <li> PENDING: Pending 
	// <li> RUNNING: Running
	// <li> SUCCESS: Success
	// <li> FAILED: Failed
	// <li> TIMEOUT: Command timed out
	// <li> PARTIAL_FAILED: Partial failure
	InvocationStatus *string `json:"InvocationStatus,omitnil" name:"InvocationStatus"`

	// Execution task information list.
	InvocationTaskBasicInfoSet []*InvocationTaskBasicInfo `json:"InvocationTaskBasicInfoSet,omitnil" name:"InvocationTaskBasicInfoSet"`

	// Execution activity description.
	Description *string `json:"Description,omitnil" name:"Description"`

	// Start time of the execution activity.
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// End time of the execution activity.
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// Time when the execution activity is created.
	CreatedTime *string `json:"CreatedTime,omitnil" name:"CreatedTime"`

	// Time when the execution activity is updated.
	UpdatedTime *string `json:"UpdatedTime,omitnil" name:"UpdatedTime"`

	// Values of custom parameters.
	Parameters *string `json:"Parameters,omitnil" name:"Parameters"`

	// Default custom parameter value.
	DefaultParameters *string `json:"DefaultParameters,omitnil" name:"DefaultParameters"`

	// Type of the instance executing the command. Valid values: `CVM`, `LIGHTHOUSE`.
	InstanceKind *string `json:"InstanceKind,omitnil" name:"InstanceKind"`

	// The user who executes the command on the instance.
	Username *string `json:"Username,omitnil" name:"Username"`

	// Invocation source.
	InvocationSource *string `json:"InvocationSource,omitnil" name:"InvocationSource"`

	// Base64-encoded command
	CommandContent *string `json:"CommandContent,omitnil" name:"CommandContent"`

	// Command type
	CommandType *string `json:"CommandType,omitnil" name:"CommandType"`

	// Command timeout period, in seconds.
	Timeout *uint64 `json:"Timeout,omitnil" name:"Timeout"`

	// Working directory for executing the command.
	WorkingDirectory *string `json:"WorkingDirectory,omitnil" name:"WorkingDirectory"`

	// The COS bucket URL for uploading logs.
	OutputCOSBucketUrl *string `json:"OutputCOSBucketUrl,omitnil" name:"OutputCOSBucketUrl"`

	// The COS bucket directory where the logs are saved.
	OutputCOSKeyPrefix *string `json:"OutputCOSKeyPrefix,omitnil" name:"OutputCOSKeyPrefix"`
}

type InvocationTask struct {
	// Execution activity ID.
	InvocationId *string `json:"InvocationId,omitnil" name:"InvocationId"`

	// Execution task ID.
	InvocationTaskId *string `json:"InvocationTaskId,omitnil" name:"InvocationTaskId"`

	// Command ID.
	CommandId *string `json:"CommandId,omitnil" name:"CommandId"`

	// Execution task status. Valid values:
	// <li> PENDING: Pending 
	// <li> DELIVERING: Delivering
	// <li> DELIVER_DELAYED: Delivery delayed 
	// <li> DELIVER_FAILED: Delivery failed
	// <li> START_FAILED: Failed to start the command
	// <li> RUNNING: Running
	// <li> SUCCESS: Success
	// <li> FAILED: Failed to execute the command. The exit code is not 0 after execution.
	// <li> TIMEOUT: Command timed out
	// <li> TASK_TIMEOUT: Task timed out
	// <li> CANCELLING: Canceling
	// <li> CANCELLED: Canceled (canceled before execution)
	// <li> TERMINATED: Terminated (canceled during execution)
	TaskStatus *string `json:"TaskStatus,omitnil" name:"TaskStatus"`

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Execution result.
	TaskResult *TaskResult `json:"TaskResult,omitnil" name:"TaskResult"`

	// Start time of the execution task.
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// End time of the execution task.
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// Creation time.
	CreatedTime *string `json:"CreatedTime,omitnil" name:"CreatedTime"`

	// Update time.
	UpdatedTime *string `json:"UpdatedTime,omitnil" name:"UpdatedTime"`

	// Command details of the execution task.
	CommandDocument *CommandDocument `json:"CommandDocument,omitnil" name:"CommandDocument"`

	// Error message displayed when the execution task fails.
	ErrorInfo *string `json:"ErrorInfo,omitnil" name:"ErrorInfo"`

	// Invocation source.
	InvocationSource *string `json:"InvocationSource,omitnil" name:"InvocationSource"`
}

type InvocationTaskBasicInfo struct {
	// Execution task ID.
	InvocationTaskId *string `json:"InvocationTaskId,omitnil" name:"InvocationTaskId"`

	// Execution task status. Valid values:
	// <li> PENDING: Pending 
	// <li> DELIVERING: Delivering
	// <li> DELIVER_DELAYED: Delivery delayed 
	// <li> DELIVER_FAILED: Delivery failed
	// <li> START_FAILED: Failed to start the command
	// <li> RUNNING: Running
	// <li> SUCCESS: Success
	// <li> FAILED: Failed to execute the command. The exit code is not 0 after execution.
	// <li> TIMEOUT: Command timed out
	// <li> TASK_TIMEOUT: Task timed out
	// <li> CANCELLING: Canceling
	// <li> CANCELLED: Canceled (canceled before execution)
	// <li> TERMINATED: Terminated (canceled during execution)
	TaskStatus *string `json:"TaskStatus,omitnil" name:"TaskStatus"`

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

// Predefined struct for user
type InvokeCommandRequestParams struct {
	// ID of the command to be triggered.
	CommandId *string `json:"CommandId,omitnil" name:"CommandId"`

	// IDs of instances about to execute commands. At most 100 IDs are allowed.
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`

	// Custom parameters of the command. The field type is JSON encoded string. For example, {\"varA\": \"222\"}.
	// `key` is the name of the custom parameter and `value` is the default value. Both `key` and `value` are strings.
	// If no parameter value is provided, the DefaultParameters of the command is used.
	// Up to 20 custom parameters are supported.
	// The name of the custom parameter cannot exceed 64 characters and can contain [a-z], [A-Z], [0-9] and [-_].
	Parameters *string `json:"Parameters,omitnil" name:"Parameters"`

	// The username used to execute the command on the CVM or Lighthouse instance.
	// The principle of the least privilege is the best practice for permission management. We recommend you execute TAT commands as a general user. If this is not specified, the Username of the command is used by default.
	Username *string `json:"Username,omitnil" name:"Username"`

	// Execution path of the command. The WorkingDirectory of the command is used by default.
	WorkingDirectory *string `json:"WorkingDirectory,omitnil" name:"WorkingDirectory"`

	// Command timeout period. Value range: [1, 86400]. The Timeout of the command is used by default.
	Timeout *uint64 `json:"Timeout,omitnil" name:"Timeout"`

	// The COS bucket URL for uploading logs. The URL must start with `https`, such as `https://BucketName-123454321.cos.ap-beijing.myqcloud.com`.
	OutputCOSBucketUrl *string `json:"OutputCOSBucketUrl,omitnil" name:"OutputCOSBucketUrl"`

	// The COS bucket directory where the logs are saved. Check below for the rules of the directory name. 
	// 1. It must be a combination of number, letters, and visible characters. Up to 60 characters are allowed.
	// 2. Use a slash (/) to create a subdirectory.
	// 3. ".." can not be used as the folder name. It cannot start with a slash (/), and cannot contain consecutive slashes.
	OutputCOSKeyPrefix *string `json:"OutputCOSKeyPrefix,omitnil" name:"OutputCOSKeyPrefix"`
}

type InvokeCommandRequest struct {
	*tchttp.BaseRequest
	
	// ID of the command to be triggered.
	CommandId *string `json:"CommandId,omitnil" name:"CommandId"`

	// IDs of instances about to execute commands. At most 100 IDs are allowed.
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`

	// Custom parameters of the command. The field type is JSON encoded string. For example, {\"varA\": \"222\"}.
	// `key` is the name of the custom parameter and `value` is the default value. Both `key` and `value` are strings.
	// If no parameter value is provided, the DefaultParameters of the command is used.
	// Up to 20 custom parameters are supported.
	// The name of the custom parameter cannot exceed 64 characters and can contain [a-z], [A-Z], [0-9] and [-_].
	Parameters *string `json:"Parameters,omitnil" name:"Parameters"`

	// The username used to execute the command on the CVM or Lighthouse instance.
	// The principle of the least privilege is the best practice for permission management. We recommend you execute TAT commands as a general user. If this is not specified, the Username of the command is used by default.
	Username *string `json:"Username,omitnil" name:"Username"`

	// Execution path of the command. The WorkingDirectory of the command is used by default.
	WorkingDirectory *string `json:"WorkingDirectory,omitnil" name:"WorkingDirectory"`

	// Command timeout period. Value range: [1, 86400]. The Timeout of the command is used by default.
	Timeout *uint64 `json:"Timeout,omitnil" name:"Timeout"`

	// The COS bucket URL for uploading logs. The URL must start with `https`, such as `https://BucketName-123454321.cos.ap-beijing.myqcloud.com`.
	OutputCOSBucketUrl *string `json:"OutputCOSBucketUrl,omitnil" name:"OutputCOSBucketUrl"`

	// The COS bucket directory where the logs are saved. Check below for the rules of the directory name. 
	// 1. It must be a combination of number, letters, and visible characters. Up to 60 characters are allowed.
	// 2. Use a slash (/) to create a subdirectory.
	// 3. ".." can not be used as the folder name. It cannot start with a slash (/), and cannot contain consecutive slashes.
	OutputCOSKeyPrefix *string `json:"OutputCOSKeyPrefix,omitnil" name:"OutputCOSKeyPrefix"`
}

func (r *InvokeCommandRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InvokeCommandRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CommandId")
	delete(f, "InstanceIds")
	delete(f, "Parameters")
	delete(f, "Username")
	delete(f, "WorkingDirectory")
	delete(f, "Timeout")
	delete(f, "OutputCOSBucketUrl")
	delete(f, "OutputCOSKeyPrefix")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InvokeCommandRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InvokeCommandResponseParams struct {
	// Execution activity ID.
	InvocationId *string `json:"InvocationId,omitnil" name:"InvocationId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type InvokeCommandResponse struct {
	*tchttp.BaseResponse
	Response *InvokeCommandResponseParams `json:"Response"`
}

func (r *InvokeCommandResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InvokeCommandResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Invoker struct {
	// Invoker ID.
	InvokerId *string `json:"InvokerId,omitnil" name:"InvokerId"`

	// Invoker name.
	Name *string `json:"Name,omitnil" name:"Name"`

	// Invoker type.
	Type *string `json:"Type,omitnil" name:"Type"`

	// Command ID.
	CommandId *string `json:"CommandId,omitnil" name:"CommandId"`

	// Username.
	Username *string `json:"Username,omitnil" name:"Username"`

	// Custom parameters.
	Parameters *string `json:"Parameters,omitnil" name:"Parameters"`

	// Instance ID list.
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`

	// Whether to enable the invoker.
	Enable *bool `json:"Enable,omitnil" name:"Enable"`

	// Execution schedule of the invoker. This field is returned for recurring invokers.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	ScheduleSettings *ScheduleSettings `json:"ScheduleSettings,omitnil" name:"ScheduleSettings"`

	// Creation time.
	CreatedTime *string `json:"CreatedTime,omitnil" name:"CreatedTime"`

	// Modification time.
	UpdatedTime *string `json:"UpdatedTime,omitnil" name:"UpdatedTime"`
}

type InvokerRecord struct {
	// Invoker ID.
	InvokerId *string `json:"InvokerId,omitnil" name:"InvokerId"`

	// Execution time.
	InvokeTime *string `json:"InvokeTime,omitnil" name:"InvokeTime"`

	// Execution reason.
	Reason *string `json:"Reason,omitnil" name:"Reason"`

	// Command execution ID.
	InvocationId *string `json:"InvocationId,omitnil" name:"InvocationId"`

	// Trigger result.
	Result *string `json:"Result,omitnil" name:"Result"`
}

// Predefined struct for user
type ModifyCommandRequestParams struct {
	// Command ID.
	CommandId *string `json:"CommandId,omitnil" name:"CommandId"`

	// Command name. The name can be up to 60 bytes, and contain [a-z], [A-Z], [0-9] and [_-.].
	CommandName *string `json:"CommandName,omitnil" name:"CommandName"`

	// Command description. The maximum length is 120 characters.
	Description *string `json:"Description,omitnil" name:"Description"`

	// Base64-encoded command. The maximum length is 64 KB.
	Content *string `json:"Content,omitnil" name:"Content"`

	// Command type. `SHELL` and `POWERSHELL` are supported.
	CommandType *string `json:"CommandType,omitnil" name:"CommandType"`

	// Command execution path.
	WorkingDirectory *string `json:"WorkingDirectory,omitnil" name:"WorkingDirectory"`

	// Command timeout period. Value range: [1, 86400].
	Timeout *uint64 `json:"Timeout,omitnil" name:"Timeout"`

	// The default value of the custom parameter value when it is enabled. The field type is JSON encoded string. For example, {\"varA\": \"222\"}.
	// All parameters are overwritten. All default values are required for modification.
	// Modification is only allowed when `EnableParameter` is `true`.
	// `key` is the name of the custom parameter and `value` is the default value. Both `key` and `value` are strings.
	// Up to 20 custom parameters are supported.
	// The name of the custom parameter cannot exceed 64 characters and can contain [a-z], [A-Z], [0-9] and [-_].
	DefaultParameters *string `json:"DefaultParameters,omitnil" name:"DefaultParameters"`

	// The username used to execute the command on the CVM or Lighthouse instance.
	// The principle of least privilege is the best practice for permission management. We recommend you execute TAT commands as a general user.
	Username *string `json:"Username,omitnil" name:"Username"`

	// The COS bucket URL for uploading logs. The URL must start with `https`, such as `https://BucketName-123454321.cos.ap-beijing.myqcloud.com`.
	OutputCOSBucketUrl *string `json:"OutputCOSBucketUrl,omitnil" name:"OutputCOSBucketUrl"`

	// The COS bucket directory where the logs are saved. Check below for the rules of the directory name. 
	// 1. It must be a combination of number, letters, and visible characters. Up to 60 characters are allowed.
	// 2. Use a slash (/) to create a subdirectory.
	// 3. ".." can not be used as the folder name. It cannot start with a slash (/), and cannot contain consecutive slashes.
	OutputCOSKeyPrefix *string `json:"OutputCOSKeyPrefix,omitnil" name:"OutputCOSKeyPrefix"`
}

type ModifyCommandRequest struct {
	*tchttp.BaseRequest
	
	// Command ID.
	CommandId *string `json:"CommandId,omitnil" name:"CommandId"`

	// Command name. The name can be up to 60 bytes, and contain [a-z], [A-Z], [0-9] and [_-.].
	CommandName *string `json:"CommandName,omitnil" name:"CommandName"`

	// Command description. The maximum length is 120 characters.
	Description *string `json:"Description,omitnil" name:"Description"`

	// Base64-encoded command. The maximum length is 64 KB.
	Content *string `json:"Content,omitnil" name:"Content"`

	// Command type. `SHELL` and `POWERSHELL` are supported.
	CommandType *string `json:"CommandType,omitnil" name:"CommandType"`

	// Command execution path.
	WorkingDirectory *string `json:"WorkingDirectory,omitnil" name:"WorkingDirectory"`

	// Command timeout period. Value range: [1, 86400].
	Timeout *uint64 `json:"Timeout,omitnil" name:"Timeout"`

	// The default value of the custom parameter value when it is enabled. The field type is JSON encoded string. For example, {\"varA\": \"222\"}.
	// All parameters are overwritten. All default values are required for modification.
	// Modification is only allowed when `EnableParameter` is `true`.
	// `key` is the name of the custom parameter and `value` is the default value. Both `key` and `value` are strings.
	// Up to 20 custom parameters are supported.
	// The name of the custom parameter cannot exceed 64 characters and can contain [a-z], [A-Z], [0-9] and [-_].
	DefaultParameters *string `json:"DefaultParameters,omitnil" name:"DefaultParameters"`

	// The username used to execute the command on the CVM or Lighthouse instance.
	// The principle of least privilege is the best practice for permission management. We recommend you execute TAT commands as a general user.
	Username *string `json:"Username,omitnil" name:"Username"`

	// The COS bucket URL for uploading logs. The URL must start with `https`, such as `https://BucketName-123454321.cos.ap-beijing.myqcloud.com`.
	OutputCOSBucketUrl *string `json:"OutputCOSBucketUrl,omitnil" name:"OutputCOSBucketUrl"`

	// The COS bucket directory where the logs are saved. Check below for the rules of the directory name. 
	// 1. It must be a combination of number, letters, and visible characters. Up to 60 characters are allowed.
	// 2. Use a slash (/) to create a subdirectory.
	// 3. ".." can not be used as the folder name. It cannot start with a slash (/), and cannot contain consecutive slashes.
	OutputCOSKeyPrefix *string `json:"OutputCOSKeyPrefix,omitnil" name:"OutputCOSKeyPrefix"`
}

func (r *ModifyCommandRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCommandRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CommandId")
	delete(f, "CommandName")
	delete(f, "Description")
	delete(f, "Content")
	delete(f, "CommandType")
	delete(f, "WorkingDirectory")
	delete(f, "Timeout")
	delete(f, "DefaultParameters")
	delete(f, "Username")
	delete(f, "OutputCOSBucketUrl")
	delete(f, "OutputCOSKeyPrefix")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCommandRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCommandResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyCommandResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCommandResponseParams `json:"Response"`
}

func (r *ModifyCommandResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCommandResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInvokerRequestParams struct {
	// ID of the invoker to be modified.
	InvokerId *string `json:"InvokerId,omitnil" name:"InvokerId"`

	// Name of the invoker to be modified.
	Name *string `json:"Name,omitnil" name:"Name"`

	// Invoker type. It can only be `SCHEDULE` (recurring invokers).
	Type *string `json:"Type,omitnil" name:"Type"`

	// ID of the command to be modified.
	CommandId *string `json:"CommandId,omitnil" name:"CommandId"`

	// The username to be modified.
	Username *string `json:"Username,omitnil" name:"Username"`

	// Custom parameters to be modified.
	Parameters *string `json:"Parameters,omitnil" name:"Parameters"`

	// List of instance IDs to be modified. Up to 100 IDs are allowed.
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`

	// Scheduled invoker settings to be modified.
	ScheduleSettings *ScheduleSettings `json:"ScheduleSettings,omitnil" name:"ScheduleSettings"`
}

type ModifyInvokerRequest struct {
	*tchttp.BaseRequest
	
	// ID of the invoker to be modified.
	InvokerId *string `json:"InvokerId,omitnil" name:"InvokerId"`

	// Name of the invoker to be modified.
	Name *string `json:"Name,omitnil" name:"Name"`

	// Invoker type. It can only be `SCHEDULE` (recurring invokers).
	Type *string `json:"Type,omitnil" name:"Type"`

	// ID of the command to be modified.
	CommandId *string `json:"CommandId,omitnil" name:"CommandId"`

	// The username to be modified.
	Username *string `json:"Username,omitnil" name:"Username"`

	// Custom parameters to be modified.
	Parameters *string `json:"Parameters,omitnil" name:"Parameters"`

	// List of instance IDs to be modified. Up to 100 IDs are allowed.
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`

	// Scheduled invoker settings to be modified.
	ScheduleSettings *ScheduleSettings `json:"ScheduleSettings,omitnil" name:"ScheduleSettings"`
}

func (r *ModifyInvokerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInvokerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InvokerId")
	delete(f, "Name")
	delete(f, "Type")
	delete(f, "CommandId")
	delete(f, "Username")
	delete(f, "Parameters")
	delete(f, "InstanceIds")
	delete(f, "ScheduleSettings")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInvokerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInvokerResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyInvokerResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInvokerResponseParams `json:"Response"`
}

func (r *ModifyInvokerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInvokerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PreviewReplacedCommandContentRequestParams struct {
	// Custom parameters for the preview. The field type is JSON encoded string. For example, {\"varA\": \"222\"}.
	// `key` is the name of the custom parameter and "value" is its specified value. Both "key" and "value" are strings.
	// At most 20 custom parameters are supported.
	// The name of the custom parameter cannot exceed 64 characters and can only contain [a-z], [A-Z], [0-9], [-_].
	// This parameter can be left empty if DefaultParameters is set for the previewed CommandId.
	Parameters *string `json:"Parameters,omitnil" name:"Parameters"`

	// The command to be previewed. If DefaultParameters is set, it is combined with Parameters and Parameters takes priority.
	// `CommandId` or `Content` must be specified.
	CommandId *string `json:"CommandId,omitnil" name:"CommandId"`

	// Base64-encoded command to be previewed. The maximum length is 64 KB.
	// CommandId or Content must be specified.
	Content *string `json:"Content,omitnil" name:"Content"`
}

type PreviewReplacedCommandContentRequest struct {
	*tchttp.BaseRequest
	
	// Custom parameters for the preview. The field type is JSON encoded string. For example, {\"varA\": \"222\"}.
	// `key` is the name of the custom parameter and "value" is its specified value. Both "key" and "value" are strings.
	// At most 20 custom parameters are supported.
	// The name of the custom parameter cannot exceed 64 characters and can only contain [a-z], [A-Z], [0-9], [-_].
	// This parameter can be left empty if DefaultParameters is set for the previewed CommandId.
	Parameters *string `json:"Parameters,omitnil" name:"Parameters"`

	// The command to be previewed. If DefaultParameters is set, it is combined with Parameters and Parameters takes priority.
	// `CommandId` or `Content` must be specified.
	CommandId *string `json:"CommandId,omitnil" name:"CommandId"`

	// Base64-encoded command to be previewed. The maximum length is 64 KB.
	// CommandId or Content must be specified.
	Content *string `json:"Content,omitnil" name:"Content"`
}

func (r *PreviewReplacedCommandContentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PreviewReplacedCommandContentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Parameters")
	delete(f, "CommandId")
	delete(f, "Content")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PreviewReplacedCommandContentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PreviewReplacedCommandContentResponseParams struct {
	// Base64-encoded command with custom parameters.
	ReplacedContent *string `json:"ReplacedContent,omitnil" name:"ReplacedContent"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type PreviewReplacedCommandContentResponse struct {
	*tchttp.BaseResponse
	Response *PreviewReplacedCommandContentResponseParams `json:"Response"`
}

func (r *PreviewReplacedCommandContentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PreviewReplacedCommandContentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RegionInfo struct {
	// Region name, such as `ap-guangzhou`
	Region *string `json:"Region,omitnil" name:"Region"`

	// Region description, such as `Guangzhou`
	RegionName *string `json:"RegionName,omitnil" name:"RegionName"`

	// Region status. `AVAILABLE` indicates the region is available.
	RegionState *string `json:"RegionState,omitnil" name:"RegionState"`
}

// Predefined struct for user
type RunCommandRequestParams struct {
	// Base64-encoded command. The maximum length is 64 KB.
	Content *string `json:"Content,omitnil" name:"Content"`

	// IDs of instances about to execute commands. Up to 100 IDs are allowed. Supported instance types:
	// <li> `CVM`
	// <li> `LIGHTHOUSE`
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`

	// Command name. The name can be up to 60 bytes, and contain [a-z], [A-Z], [0-9] and [_-.].
	CommandName *string `json:"CommandName,omitnil" name:"CommandName"`

	// Command description. The maximum length is 120 characters.
	Description *string `json:"Description,omitnil" name:"Description"`

	// Command type. `SHELL` and `POWERSHELL` are supported. The default value is `SHELL`.
	CommandType *string `json:"CommandType,omitnil" name:"CommandType"`

	// Command execution path. The default value is /root for `SHELL` commands and C:\Program Files\qcloud\tat_agent\workdir for `POWERSHELL` commands.
	WorkingDirectory *string `json:"WorkingDirectory,omitnil" name:"WorkingDirectory"`

	// Command timeout period. Default value: 60 seconds. Value range: [1, 86400].
	Timeout *uint64 `json:"Timeout,omitnil" name:"Timeout"`

	// Whether to save the command. Valid values:
	// <li> `True`: Save
	// <li> `False`: Do not save
	// The default value is `False`.
	SaveCommand *bool `json:"SaveCommand,omitnil" name:"SaveCommand"`

	// Whether to enable the custom parameter feature.
	// This cannot be modified once created.
	// Default value: `false`.
	EnableParameter *bool `json:"EnableParameter,omitnil" name:"EnableParameter"`

	// The default value of the custom parameter value when it is enabled. The field type is JSON encoded string. For example, {\"varA\": \"222\"}.
	// `key` is the name of the custom parameter and `value` is the default value. Both `key` and `value` are strings.
	// If Parameters is not provided, the default values specified here are used.
	// Up to 20 custom parameters are supported.
	// The name of the custom parameter cannot exceed 64 characters and can contain [a-z], [A-Z], [0-9] and [-_].
	DefaultParameters *string `json:"DefaultParameters,omitnil" name:"DefaultParameters"`

	// Custom parameters of `Command`. The field type is JSON encoded string. For example, {\"varA\": \"222\"}.
	// `key` is the name of the custom parameter and `value` is the default value. Both `key` and `value` are strings.
	// If no parameter value is provided, the `DefaultParameters` is used.
	// Up to 20 custom parameters are supported.
	// The name of the custom parameter cannot exceed 64 characters and can contain [a-z], [A-Z], [0-9] and [-_].
	Parameters *string `json:"Parameters,omitnil" name:"Parameters"`

	// The tags of the command. It is available when `SaveCommand` is `True`. A maximum of 10 tags are allowed.
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`

	// The username used to execute the command on the CVM or Lighthouse instance.
	// The principle of least privilege is the best practice for permission management. We recommend you execute TAT commands as a general user. By default, the user `root` is used to execute commands on Linux and the user `System` is used on Windows.
	Username *string `json:"Username,omitnil" name:"Username"`

	// The COS bucket URL for uploading logs. The URL must start with `https`, such as `https://BucketName-123454321.cos.ap-beijing.myqcloud.com`.
	OutputCOSBucketUrl *string `json:"OutputCOSBucketUrl,omitnil" name:"OutputCOSBucketUrl"`

	// The COS bucket directory where the logs are saved. Check below for the rules of the directory name. 
	// 1. It must be a combination of number, letters, and visible characters. Up to 60 characters are allowed.
	// 2. Use a slash (/) to create a subdirectory.
	// 3. ".." can not be used as the folder name. It cannot start with a slash (/), and cannot contain consecutive slashes.
	OutputCOSKeyPrefix *string `json:"OutputCOSKeyPrefix,omitnil" name:"OutputCOSKeyPrefix"`
}

type RunCommandRequest struct {
	*tchttp.BaseRequest
	
	// Base64-encoded command. The maximum length is 64 KB.
	Content *string `json:"Content,omitnil" name:"Content"`

	// IDs of instances about to execute commands. Up to 100 IDs are allowed. Supported instance types:
	// <li> `CVM`
	// <li> `LIGHTHOUSE`
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`

	// Command name. The name can be up to 60 bytes, and contain [a-z], [A-Z], [0-9] and [_-.].
	CommandName *string `json:"CommandName,omitnil" name:"CommandName"`

	// Command description. The maximum length is 120 characters.
	Description *string `json:"Description,omitnil" name:"Description"`

	// Command type. `SHELL` and `POWERSHELL` are supported. The default value is `SHELL`.
	CommandType *string `json:"CommandType,omitnil" name:"CommandType"`

	// Command execution path. The default value is /root for `SHELL` commands and C:\Program Files\qcloud\tat_agent\workdir for `POWERSHELL` commands.
	WorkingDirectory *string `json:"WorkingDirectory,omitnil" name:"WorkingDirectory"`

	// Command timeout period. Default value: 60 seconds. Value range: [1, 86400].
	Timeout *uint64 `json:"Timeout,omitnil" name:"Timeout"`

	// Whether to save the command. Valid values:
	// <li> `True`: Save
	// <li> `False`: Do not save
	// The default value is `False`.
	SaveCommand *bool `json:"SaveCommand,omitnil" name:"SaveCommand"`

	// Whether to enable the custom parameter feature.
	// This cannot be modified once created.
	// Default value: `false`.
	EnableParameter *bool `json:"EnableParameter,omitnil" name:"EnableParameter"`

	// The default value of the custom parameter value when it is enabled. The field type is JSON encoded string. For example, {\"varA\": \"222\"}.
	// `key` is the name of the custom parameter and `value` is the default value. Both `key` and `value` are strings.
	// If Parameters is not provided, the default values specified here are used.
	// Up to 20 custom parameters are supported.
	// The name of the custom parameter cannot exceed 64 characters and can contain [a-z], [A-Z], [0-9] and [-_].
	DefaultParameters *string `json:"DefaultParameters,omitnil" name:"DefaultParameters"`

	// Custom parameters of `Command`. The field type is JSON encoded string. For example, {\"varA\": \"222\"}.
	// `key` is the name of the custom parameter and `value` is the default value. Both `key` and `value` are strings.
	// If no parameter value is provided, the `DefaultParameters` is used.
	// Up to 20 custom parameters are supported.
	// The name of the custom parameter cannot exceed 64 characters and can contain [a-z], [A-Z], [0-9] and [-_].
	Parameters *string `json:"Parameters,omitnil" name:"Parameters"`

	// The tags of the command. It is available when `SaveCommand` is `True`. A maximum of 10 tags are allowed.
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`

	// The username used to execute the command on the CVM or Lighthouse instance.
	// The principle of least privilege is the best practice for permission management. We recommend you execute TAT commands as a general user. By default, the user `root` is used to execute commands on Linux and the user `System` is used on Windows.
	Username *string `json:"Username,omitnil" name:"Username"`

	// The COS bucket URL for uploading logs. The URL must start with `https`, such as `https://BucketName-123454321.cos.ap-beijing.myqcloud.com`.
	OutputCOSBucketUrl *string `json:"OutputCOSBucketUrl,omitnil" name:"OutputCOSBucketUrl"`

	// The COS bucket directory where the logs are saved. Check below for the rules of the directory name. 
	// 1. It must be a combination of number, letters, and visible characters. Up to 60 characters are allowed.
	// 2. Use a slash (/) to create a subdirectory.
	// 3. ".." can not be used as the folder name. It cannot start with a slash (/), and cannot contain consecutive slashes.
	OutputCOSKeyPrefix *string `json:"OutputCOSKeyPrefix,omitnil" name:"OutputCOSKeyPrefix"`
}

func (r *RunCommandRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunCommandRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Content")
	delete(f, "InstanceIds")
	delete(f, "CommandName")
	delete(f, "Description")
	delete(f, "CommandType")
	delete(f, "WorkingDirectory")
	delete(f, "Timeout")
	delete(f, "SaveCommand")
	delete(f, "EnableParameter")
	delete(f, "DefaultParameters")
	delete(f, "Parameters")
	delete(f, "Tags")
	delete(f, "Username")
	delete(f, "OutputCOSBucketUrl")
	delete(f, "OutputCOSKeyPrefix")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RunCommandRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RunCommandResponseParams struct {
	// Command ID.
	CommandId *string `json:"CommandId,omitnil" name:"CommandId"`

	// Execution activity ID.
	InvocationId *string `json:"InvocationId,omitnil" name:"InvocationId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type RunCommandResponse struct {
	*tchttp.BaseResponse
	Response *RunCommandResponseParams `json:"Response"`
}

func (r *RunCommandResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunCommandResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ScheduleSettings struct {
	// Execution policy:
	// <br><li>`ONCE`: Execute once
	// <br><li>`RECURRENCE`: Execute repeatedly
	Policy *string `json:"Policy,omitnil" name:"Policy"`

	// Trigger the crontab expression. This field is required if `Policy` is `RECURRENCE`. The crontab expression is parsed in UTC+8.
	Recurrence *string `json:"Recurrence,omitnil" name:"Recurrence"`

	// The next execution time of the invoker. This field is required if `Policy` is `ONCE`.
	InvokeTime *string `json:"InvokeTime,omitnil" name:"InvokeTime"`
}

type Tag struct {
	// Tag key.
	Key *string `json:"Key,omitnil" name:"Key"`

	// Tag value.
	Value *string `json:"Value,omitnil" name:"Value"`
}

type TaskResult struct {
	// ExitCode of the execution.
	ExitCode *int64 `json:"ExitCode,omitnil" name:"ExitCode"`

	// Base64-encoded command output. The maximum length is 24 KB.
	Output *string `json:"Output,omitnil" name:"Output"`

	// Time when the execution is started.
	ExecStartTime *string `json:"ExecStartTime,omitnil" name:"ExecStartTime"`

	// Time when the execution is ended.
	ExecEndTime *string `json:"ExecEndTime,omitnil" name:"ExecEndTime"`

	// Dropped bytes of the command output.
	Dropped *uint64 `json:"Dropped,omitnil" name:"Dropped"`

	// COS URL of the logs.
	OutputUrl *string `json:"OutputUrl,omitnil" name:"OutputUrl"`

	// Error message for uploading logs to COS.
	OutputUploadCOSErrorInfo *string `json:"OutputUploadCOSErrorInfo,omitnil" name:"OutputUploadCOSErrorInfo"`
}