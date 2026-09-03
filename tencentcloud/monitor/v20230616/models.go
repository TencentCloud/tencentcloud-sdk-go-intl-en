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

package v20230616

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

type AgentInfo struct {
	// <p>Agent ID</p>
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// <p>Agent name</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Agent description</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>Agent Category.</p>
	Category *string `json:"Category,omitnil,omitempty" name:"Category"`

	// <p>Status: draft/configured/running/standby/disabled</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>List of associated skill IDs.</p>
	SkillIds []*string `json:"SkillIds,omitnil,omitempty" name:"SkillIds"`

	// <p>Associated resource map ID.</p>
	ResourceMapId *string `json:"ResourceMapId,omitnil,omitempty" name:"ResourceMapId"`

	// <p>Associated mcp id.</p>
	MCPIds []*string `json:"MCPIds,omitnil,omitempty" name:"MCPIds"`

	// <p>Resource Tag.</p>
	CamTags []*Tag `json:"CamTags,omitnil,omitempty" name:"CamTags"`

	// <p>Environment variables required by the agent at runtime</p>
	EnvVars []*EnvVar `json:"EnvVars,omitnil,omitempty" name:"EnvVars"`
}

type AlarmLable struct {
	// label name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// label value
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type AlarmNotifyHistory struct {
	// Unique notification ID.
	NotifyId *string `json:"NotifyId,omitnil,omitempty" name:"NotifyId"`

	// Alert policy ID
	PolicyId *string `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// Alarm cycle iD
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// Notification time in Unix timestamp (in seconds).
	NotifyTime *int64 `json:"NotifyTime,omitnil,omitempty" name:"NotifyTime"`

	// Trigger time in Unix timestamp (in seconds).
	TriggerTime *int64 `json:"TriggerTime,omitnil,omitempty" name:"TriggerTime"`

	// Alarm severity level. Valid values: None, Note, Warn, and Serious.
	TriggerLevel *string `json:"TriggerLevel,omitnil,omitempty" name:"TriggerLevel"`

	// alert content
	AlarmContent *string `json:"AlarmContent,omitnil,omitempty" name:"AlarmContent"`

	// Alarm object
	AlarmObject *string `json:"AlarmObject,omitnil,omitempty" name:"AlarmObject"`

	// Alarm notification channel collection involved this time
	// Note: This field may return null, indicating that no valid values can be obtained.
	ChannelSet []*string `json:"ChannelSet,omitnil,omitempty" name:"ChannelSet"`

	// Recipient information of the channel
	ChannelsReceivers []*ChannelsReceivers `json:"ChannelsReceivers,omitnil,omitempty" name:"ChannelsReceivers"`

	// Alarm policy name
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`

	// Prometheus Instance ID, valid only when MT_PROME
	PromeInstanceID *string `json:"PromeInstanceID,omitnil,omitempty" name:"PromeInstanceID"`

	// Region of the Prometheus Instance. Valid at that time only for MT_PROME.
	PromeInstanceRegion *string `json:"PromeInstanceRegion,omitnil,omitempty" name:"PromeInstanceRegion"`

	// Notification template related configuration information
	Notices []*NotifyRelatedNotice `json:"Notices,omitnil,omitempty" name:"Notices"`

	// Alarm trigger status. Valid values: Trigger and Recovery.
	TriggerStatus *string `json:"TriggerStatus,omitnil,omitempty" name:"TriggerStatus"`

	// Console page address related to the present Prometheus notification history, valid only when MR_PROME
	PromeConsoleURL *string `json:"PromeConsoleURL,omitnil,omitempty" name:"PromeConsoleURL"`

	// Alarm label
	Labels []*AlarmLable `json:"Labels,omitnil,omitempty" name:"Labels"`
}

type ArtifactInfo struct {
	// <p>Product ID</p>
	ArtifactId *string `json:"ArtifactId,omitnil,omitempty" name:"ArtifactId"`

	// <p>Product name</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Physical type</p>
	MimeType *string `json:"MimeType,omitnil,omitempty" name:"MimeType"`

	// <p>File size (byte)</p>
	SizeBytes *int64 `json:"SizeBytes,omitnil,omitempty" name:"SizeBytes"`

	// <p>Whether it is public</p>
	IsGlobal *bool `json:"IsGlobal,omitnil,omitempty" name:"IsGlobal"`

	// <p>Creation time (Unix timestamp in seconds).</p>
	CreatedAt *int64 `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// <p>Modification time.</p>
	UpdatedAt *int64 `json:"UpdatedAt,omitnil,omitempty" name:"UpdatedAt"`

	// <p>Agent ID that generated the artifact</p>
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// <p>Skill ID that generates the artifact</p>
	SkillId *string `json:"SkillId,omitnil,omitempty" name:"SkillId"`

	// <p>For parsing calls to the download API</p>
	StoragePath *string `json:"StoragePath,omitnil,omitempty" name:"StoragePath"`
}

// Predefined struct for user
type CancelAIWorkbenchChatRequestParams struct {
	// <p>Session id.</p>
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`
}

type CancelAIWorkbenchChatRequest struct {
	*tchttp.BaseRequest
	
	// <p>Session id.</p>
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`
}

func (r *CancelAIWorkbenchChatRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelAIWorkbenchChatRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SessionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CancelAIWorkbenchChatRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelAIWorkbenchChatResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CancelAIWorkbenchChatResponse struct {
	*tchttp.BaseResponse
	Response *CancelAIWorkbenchChatResponseParams `json:"Response"`
}

func (r *CancelAIWorkbenchChatResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelAIWorkbenchChatResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ChannelsReceivers struct {
	// Notification channel name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ChannelName *string `json:"ChannelName,omitnil,omitempty" name:"ChannelName"`

	// Recipient.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Receivers []*string `json:"Receivers,omitnil,omitempty" name:"Receivers"`

	// Sending result. Valid values: 0, (invalid), 1 (successful), 2 (failed), and 3 (no sending required).
	// Note: This field may return null, indicating that no valid values can be obtained.
	SendStatus *string `json:"SendStatus,omitnil,omitempty" name:"SendStatus"`
}

type ContentBlockInfo struct {
	// <p>Type.</p>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>Data content.</p>
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`
}

// Predefined struct for user
type CreateAIWorkbenchAgentRequestParams struct {
	// <p>Agent Name</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Agent description</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>Agent Category</p>
	Category *string `json:"Category,omitnil,omitempty" name:"Category"`

	// <p>Agent tag</p>
	Tags []*string `json:"Tags,omitnil,omitempty" name:"Tags"`

	// <p>Agent prompt content</p>
	Instruction *InstructionConfig `json:"Instruction,omitnil,omitempty" name:"Instruction"`

	// <p>List of associated skill IDs.</p>
	SkillIds []*string `json:"SkillIds,omitnil,omitempty" name:"SkillIds"`

	// <p>Source: builtin / custom</p>
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// <p>Map ID of the associated resource</p>
	ResourceMapId *string `json:"ResourceMapId,omitnil,omitempty" name:"ResourceMapId"`

	// <p>Associated mcp tool</p>
	MCPIds []*string `json:"MCPIds,omitnil,omitempty" name:"MCPIds"`

	// <p>Resource tag</p>
	CamTags []*Tag `json:"CamTags,omitnil,omitempty" name:"CamTags"`

	// <p>agent runtime environment variable</p>
	EnvVars []*EnvVar `json:"EnvVars,omitnil,omitempty" name:"EnvVars"`
}

type CreateAIWorkbenchAgentRequest struct {
	*tchttp.BaseRequest
	
	// <p>Agent Name</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Agent description</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>Agent Category</p>
	Category *string `json:"Category,omitnil,omitempty" name:"Category"`

	// <p>Agent tag</p>
	Tags []*string `json:"Tags,omitnil,omitempty" name:"Tags"`

	// <p>Agent prompt content</p>
	Instruction *InstructionConfig `json:"Instruction,omitnil,omitempty" name:"Instruction"`

	// <p>List of associated skill IDs.</p>
	SkillIds []*string `json:"SkillIds,omitnil,omitempty" name:"SkillIds"`

	// <p>Source: builtin / custom</p>
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// <p>Map ID of the associated resource</p>
	ResourceMapId *string `json:"ResourceMapId,omitnil,omitempty" name:"ResourceMapId"`

	// <p>Associated mcp tool</p>
	MCPIds []*string `json:"MCPIds,omitnil,omitempty" name:"MCPIds"`

	// <p>Resource tag</p>
	CamTags []*Tag `json:"CamTags,omitnil,omitempty" name:"CamTags"`

	// <p>agent runtime environment variable</p>
	EnvVars []*EnvVar `json:"EnvVars,omitnil,omitempty" name:"EnvVars"`
}

func (r *CreateAIWorkbenchAgentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAIWorkbenchAgentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Description")
	delete(f, "Category")
	delete(f, "Tags")
	delete(f, "Instruction")
	delete(f, "SkillIds")
	delete(f, "Source")
	delete(f, "ResourceMapId")
	delete(f, "MCPIds")
	delete(f, "CamTags")
	delete(f, "EnvVars")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAIWorkbenchAgentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAIWorkbenchAgentResponseParams struct {
	// <p>Agent ID</p>
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAIWorkbenchAgentResponse struct {
	*tchttp.BaseResponse
	Response *CreateAIWorkbenchAgentResponseParams `json:"Response"`
}

func (r *CreateAIWorkbenchAgentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAIWorkbenchAgentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAIWorkbenchTaskRequestParams struct {
	// <p>Task Name</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Task description</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>Associated Agent ID</p>
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// <p>Prompt Template</p>
	PromptTemplate *string `json:"PromptTemplate,omitnil,omitempty" name:"PromptTemplate"`

	// <p>Output format: markdown / json</p>
	OutputFormat *string `json:"OutputFormat,omitnil,omitempty" name:"OutputFormat"`

	// <p>Trigger type: manual / cron / webhook</p>
	TriggerType *string `json:"TriggerType,omitnil,omitempty" name:"TriggerType"`

	// <p>Cron expression</p>
	CronExpr *string `json:"CronExpr,omitnil,omitempty" name:"CronExpr"`

	// <p>Cron time zone</p>
	CronTimezone *string `json:"CronTimezone,omitnil,omitempty" name:"CronTimezone"`

	// <p>Associated resource map ID</p>
	ResourceMapId *string `json:"ResourceMapId,omitnil,omitempty" name:"ResourceMapId"`

	// <p>Skill ID list</p>
	SkillIds []*string `json:"SkillIds,omitnil,omitempty" name:"SkillIds"`

	// <p>MCP endpoint ID list</p>
	McpEndpointIds []*string `json:"McpEndpointIds,omitnil,omitempty" name:"McpEndpointIds"`

	// <p>Timeout (seconds)</p>
	TimeoutSec *int64 `json:"TimeoutSec,omitnil,omitempty" name:"TimeoutSec"`

	// <p>Retry count</p>
	RetryCount *int64 `json:"RetryCount,omitnil,omitempty" name:"RetryCount"`

	// <p>Whether to enable</p>
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`
}

type CreateAIWorkbenchTaskRequest struct {
	*tchttp.BaseRequest
	
	// <p>Task Name</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Task description</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>Associated Agent ID</p>
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// <p>Prompt Template</p>
	PromptTemplate *string `json:"PromptTemplate,omitnil,omitempty" name:"PromptTemplate"`

	// <p>Output format: markdown / json</p>
	OutputFormat *string `json:"OutputFormat,omitnil,omitempty" name:"OutputFormat"`

	// <p>Trigger type: manual / cron / webhook</p>
	TriggerType *string `json:"TriggerType,omitnil,omitempty" name:"TriggerType"`

	// <p>Cron expression</p>
	CronExpr *string `json:"CronExpr,omitnil,omitempty" name:"CronExpr"`

	// <p>Cron time zone</p>
	CronTimezone *string `json:"CronTimezone,omitnil,omitempty" name:"CronTimezone"`

	// <p>Associated resource map ID</p>
	ResourceMapId *string `json:"ResourceMapId,omitnil,omitempty" name:"ResourceMapId"`

	// <p>Skill ID list</p>
	SkillIds []*string `json:"SkillIds,omitnil,omitempty" name:"SkillIds"`

	// <p>MCP endpoint ID list</p>
	McpEndpointIds []*string `json:"McpEndpointIds,omitnil,omitempty" name:"McpEndpointIds"`

	// <p>Timeout (seconds)</p>
	TimeoutSec *int64 `json:"TimeoutSec,omitnil,omitempty" name:"TimeoutSec"`

	// <p>Retry count</p>
	RetryCount *int64 `json:"RetryCount,omitnil,omitempty" name:"RetryCount"`

	// <p>Whether to enable</p>
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`
}

func (r *CreateAIWorkbenchTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAIWorkbenchTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Description")
	delete(f, "AgentId")
	delete(f, "PromptTemplate")
	delete(f, "OutputFormat")
	delete(f, "TriggerType")
	delete(f, "CronExpr")
	delete(f, "CronTimezone")
	delete(f, "ResourceMapId")
	delete(f, "SkillIds")
	delete(f, "McpEndpointIds")
	delete(f, "TimeoutSec")
	delete(f, "RetryCount")
	delete(f, "Enabled")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAIWorkbenchTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAIWorkbenchTaskResponseParams struct {
	// <p>Task ID.</p>
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAIWorkbenchTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateAIWorkbenchTaskResponseParams `json:"Response"`
}

func (r *CreateAIWorkbenchTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAIWorkbenchTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAIWorkbenchAgentRequestParams struct {
	// <p>Agent ID</p>
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`
}

type DeleteAIWorkbenchAgentRequest struct {
	*tchttp.BaseRequest
	
	// <p>Agent ID</p>
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`
}

func (r *DeleteAIWorkbenchAgentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAIWorkbenchAgentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AgentId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAIWorkbenchAgentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAIWorkbenchAgentResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteAIWorkbenchAgentResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAIWorkbenchAgentResponseParams `json:"Response"`
}

func (r *DeleteAIWorkbenchAgentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAIWorkbenchAgentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAIWorkbenchTaskRequestParams struct {
	// <p>Task ID.</p>
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type DeleteAIWorkbenchTaskRequest struct {
	*tchttp.BaseRequest
	
	// <p>Task ID.</p>
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *DeleteAIWorkbenchTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAIWorkbenchTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAIWorkbenchTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAIWorkbenchTaskResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteAIWorkbenchTaskResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAIWorkbenchTaskResponseParams `json:"Response"`
}

func (r *DeleteAIWorkbenchTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAIWorkbenchTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAIWorkbenchAgentRequestParams struct {
	// <p>Agent ID</p>
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`
}

type DescribeAIWorkbenchAgentRequest struct {
	*tchttp.BaseRequest
	
	// <p>Agent ID</p>
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`
}

func (r *DescribeAIWorkbenchAgentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAIWorkbenchAgentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AgentId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAIWorkbenchAgentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAIWorkbenchAgentResponseParams struct {
	// <p>Agent Information</p>
	Agent *AgentInfo `json:"Agent,omitnil,omitempty" name:"Agent"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAIWorkbenchAgentResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAIWorkbenchAgentResponseParams `json:"Response"`
}

func (r *DescribeAIWorkbenchAgentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAIWorkbenchAgentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAIWorkbenchArtifactRequestParams struct {
	// <p>Product ID</p>
	ArtifactId *string `json:"ArtifactId,omitnil,omitempty" name:"ArtifactId"`

	// <p>Whether to download the URL</p><p><code>1</code> = required, <code>0</code> or not passed = not required</p>
	NeedDownloadURL *int64 `json:"NeedDownloadURL,omitnil,omitempty" name:"NeedDownloadURL"`
}

type DescribeAIWorkbenchArtifactRequest struct {
	*tchttp.BaseRequest
	
	// <p>Product ID</p>
	ArtifactId *string `json:"ArtifactId,omitnil,omitempty" name:"ArtifactId"`

	// <p>Whether to download the URL</p><p><code>1</code> = required, <code>0</code> or not passed = not required</p>
	NeedDownloadURL *int64 `json:"NeedDownloadURL,omitnil,omitempty" name:"NeedDownloadURL"`
}

func (r *DescribeAIWorkbenchArtifactRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAIWorkbenchArtifactRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ArtifactId")
	delete(f, "NeedDownloadURL")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAIWorkbenchArtifactRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAIWorkbenchArtifactResponseParams struct {
	// <p>Product information</p>
	Artifact *ArtifactInfo `json:"Artifact,omitnil,omitempty" name:"Artifact"`

	// <p>COS pre-signed download URL</p>
	DownloadURL *string `json:"DownloadURL,omitnil,omitempty" name:"DownloadURL"`

	// <p>Download URL expiration time (in RFC3339 format)</p>
	DownloadURLExpiredAt *string `json:"DownloadURLExpiredAt,omitnil,omitempty" name:"DownloadURLExpiredAt"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAIWorkbenchArtifactResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAIWorkbenchArtifactResponseParams `json:"Response"`
}

func (r *DescribeAIWorkbenchArtifactResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAIWorkbenchArtifactResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAIWorkbenchExecutionRequestParams struct {
	// <p>Execution ID</p>
	ExecutionId *string `json:"ExecutionId,omitnil,omitempty" name:"ExecutionId"`
}

type DescribeAIWorkbenchExecutionRequest struct {
	*tchttp.BaseRequest
	
	// <p>Execution ID</p>
	ExecutionId *string `json:"ExecutionId,omitnil,omitempty" name:"ExecutionId"`
}

func (r *DescribeAIWorkbenchExecutionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAIWorkbenchExecutionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ExecutionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAIWorkbenchExecutionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAIWorkbenchExecutionResponseParams struct {
	// <p>Execution Record</p>
	Execution *ExecutionInfo `json:"Execution,omitnil,omitempty" name:"Execution"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAIWorkbenchExecutionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAIWorkbenchExecutionResponseParams `json:"Response"`
}

func (r *DescribeAIWorkbenchExecutionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAIWorkbenchExecutionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAIWorkbenchSessionRequestParams struct {
	// <p>Session ID</p>
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`
}

type DescribeAIWorkbenchSessionRequest struct {
	*tchttp.BaseRequest
	
	// <p>Session ID</p>
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`
}

func (r *DescribeAIWorkbenchSessionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAIWorkbenchSessionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SessionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAIWorkbenchSessionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAIWorkbenchSessionResponseParams struct {
	// <p>Session information</p>
	Session *SessionInfo `json:"Session,omitnil,omitempty" name:"Session"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAIWorkbenchSessionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAIWorkbenchSessionResponseParams `json:"Response"`
}

func (r *DescribeAIWorkbenchSessionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAIWorkbenchSessionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAIWorkbenchSkillRequestParams struct {
	// <p>Skill ID</p>
	SkillId *string `json:"SkillId,omitnil,omitempty" name:"SkillId"`
}

type DescribeAIWorkbenchSkillRequest struct {
	*tchttp.BaseRequest
	
	// <p>Skill ID</p>
	SkillId *string `json:"SkillId,omitnil,omitempty" name:"SkillId"`
}

func (r *DescribeAIWorkbenchSkillRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAIWorkbenchSkillRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SkillId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAIWorkbenchSkillRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAIWorkbenchSkillResponseParams struct {
	// <p>Skill information.</p>
	Skill *SkillInfo `json:"Skill,omitnil,omitempty" name:"Skill"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAIWorkbenchSkillResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAIWorkbenchSkillResponseParams `json:"Response"`
}

func (r *DescribeAIWorkbenchSkillResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAIWorkbenchSkillResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAlarmNotifyHistoriesRequestParams struct {
	// Monitoring type
	MonitorType *string `json:"MonitorType,omitnil,omitempty" name:"MonitorType"`

	// Start time, used as a Unix timestamp in seconds.
	QueryBaseTime *int64 `json:"QueryBaseTime,omitnil,omitempty" name:"QueryBaseTime"`

	// Period to query before QueryBaseTime, in seconds.
	QueryBeforeSeconds *int64 `json:"QueryBeforeSeconds,omitnil,omitempty" name:"QueryBeforeSeconds"`

	// Pagination parameter.
	PageParams *PageByNoParams `json:"PageParams,omitnil,omitempty" name:"PageParams"`

	// Fill in when the monitoring type is MT_QCE. Namespace of the affiliation.
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// Fill in when the monitoring type is MT_QCE. Alarm policy type
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// Query the notification history of a policy
	PolicyId *string `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`
}

type DescribeAlarmNotifyHistoriesRequest struct {
	*tchttp.BaseRequest
	
	// Monitoring type
	MonitorType *string `json:"MonitorType,omitnil,omitempty" name:"MonitorType"`

	// Start time, used as a Unix timestamp in seconds.
	QueryBaseTime *int64 `json:"QueryBaseTime,omitnil,omitempty" name:"QueryBaseTime"`

	// Period to query before QueryBaseTime, in seconds.
	QueryBeforeSeconds *int64 `json:"QueryBeforeSeconds,omitnil,omitempty" name:"QueryBeforeSeconds"`

	// Pagination parameter.
	PageParams *PageByNoParams `json:"PageParams,omitnil,omitempty" name:"PageParams"`

	// Fill in when the monitoring type is MT_QCE. Namespace of the affiliation.
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// Fill in when the monitoring type is MT_QCE. Alarm policy type
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// Query the notification history of a policy
	PolicyId *string `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`
}

func (r *DescribeAlarmNotifyHistoriesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlarmNotifyHistoriesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MonitorType")
	delete(f, "QueryBaseTime")
	delete(f, "QueryBeforeSeconds")
	delete(f, "PageParams")
	delete(f, "Namespace")
	delete(f, "ModelName")
	delete(f, "PolicyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAlarmNotifyHistoriesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAlarmNotifyHistoriesResponseParams struct {
	// Alarm history
	AlarmNotifyHistoryList []*AlarmNotifyHistory `json:"AlarmNotifyHistoryList,omitnil,omitempty" name:"AlarmNotifyHistoryList"`

	// Pagination condition
	PageResult *PageByNoResult `json:"PageResult,omitnil,omitempty" name:"PageResult"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAlarmNotifyHistoriesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAlarmNotifyHistoriesResponseParams `json:"Response"`
}

func (r *DescribeAlarmNotifyHistoriesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlarmNotifyHistoriesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnvEntry struct {
	// <p>Environment variable value</p>
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// <p>Whether to mask</p>
	Sensitive *bool `json:"Sensitive,omitnil,omitempty" name:"Sensitive"`
}

type EnvVar struct {
	// <p>Environment variable key</p>
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// <p>Environment variable value</p>
	Value *EnvEntry `json:"Value,omitnil,omitempty" name:"Value"`
}

type ExecutionInfo struct {
	// <p>Task name</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Task ID.</p>
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// <p>Execution ID</p>
	ExecutionId *string `json:"ExecutionId,omitnil,omitempty" name:"ExecutionId"`

	// <p>Agent ID</p>
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// <p>Session ID</p>
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// <p>Trigger type: manual / cron / webhook</p>
	TriggerType *string `json:"TriggerType,omitnil,omitempty" name:"TriggerType"`

	// <p>Status: pending/running/completed/failed/timeout/cancelled</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>Execution Abstract</p>
	Summary *string `json:"Summary,omitnil,omitempty" name:"Summary"`

	// <p>Execution time (ms)</p>
	DurationMs *int64 `json:"DurationMs,omitnil,omitempty" name:"DurationMs"`
}

// Predefined struct for user
type GetAIWorkbenchArtifactDownloadURLRequestParams struct {
	// <p>Session ID.</p>
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// <p>Artifact ID</p>
	ArtifactId *string `json:"ArtifactId,omitnil,omitempty" name:"ArtifactId"`
}

type GetAIWorkbenchArtifactDownloadURLRequest struct {
	*tchttp.BaseRequest
	
	// <p>Session ID.</p>
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// <p>Artifact ID</p>
	ArtifactId *string `json:"ArtifactId,omitnil,omitempty" name:"ArtifactId"`
}

func (r *GetAIWorkbenchArtifactDownloadURLRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetAIWorkbenchArtifactDownloadURLRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SessionId")
	delete(f, "ArtifactId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetAIWorkbenchArtifactDownloadURLRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetAIWorkbenchArtifactDownloadURLResponseParams struct {
	// <p>COS pre-signed HTTPS download URL</p>
	DownloadURL *string `json:"DownloadURL,omitnil,omitempty" name:"DownloadURL"`

	// <p>URL expiration time (RFC3339 format)</p>
	ExpiredAt *string `json:"ExpiredAt,omitnil,omitempty" name:"ExpiredAt"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetAIWorkbenchArtifactDownloadURLResponse struct {
	*tchttp.BaseResponse
	Response *GetAIWorkbenchArtifactDownloadURLResponseParams `json:"Response"`
}

func (r *GetAIWorkbenchArtifactDownloadURLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetAIWorkbenchArtifactDownloadURLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstructionConfig struct {
	// <p>Role definition</p>
	RolePosition *string `json:"RolePosition,omitnil,omitempty" name:"RolePosition"`

	// <p>Core responsibility</p>
	CoreDuty *string `json:"CoreDuty,omitnil,omitempty" name:"CoreDuty"`

	// <p>Core principle</p>
	CoreTruths *string `json:"CoreTruths,omitnil,omitempty" name:"CoreTruths"`

	// <p>Style constraints</p>
	Vibe *string `json:"Vibe,omitnil,omitempty" name:"Vibe"`

	// <p>Notes</p>
	Boundaries *string `json:"Boundaries,omitnil,omitempty" name:"Boundaries"`
}

// Predefined struct for user
type ListAIWorkbenchAgentsRequestParams struct {
	// <p>Number of items per page</p>
	PerPage *int64 `json:"PerPage,omitnil,omitempty" name:"PerPage"`

	// <p>Page number.</p>
	PageNo *int64 `json:"PageNo,omitnil,omitempty" name:"PageNo"`

	// <p>Status filtering</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>Category filtering</p>
	Category *string `json:"Category,omitnil,omitempty" name:"Category"`

	// <p>Search keyword</p>
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// <p>Filter by source</p>
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// <p>Agent ID list filtering</p>
	AgentIds []*string `json:"AgentIds,omitnil,omitempty" name:"AgentIds"`
}

type ListAIWorkbenchAgentsRequest struct {
	*tchttp.BaseRequest
	
	// <p>Number of items per page</p>
	PerPage *int64 `json:"PerPage,omitnil,omitempty" name:"PerPage"`

	// <p>Page number.</p>
	PageNo *int64 `json:"PageNo,omitnil,omitempty" name:"PageNo"`

	// <p>Status filtering</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>Category filtering</p>
	Category *string `json:"Category,omitnil,omitempty" name:"Category"`

	// <p>Search keyword</p>
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// <p>Filter by source</p>
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// <p>Agent ID list filtering</p>
	AgentIds []*string `json:"AgentIds,omitnil,omitempty" name:"AgentIds"`
}

func (r *ListAIWorkbenchAgentsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAIWorkbenchAgentsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PerPage")
	delete(f, "PageNo")
	delete(f, "Status")
	delete(f, "Category")
	delete(f, "Keyword")
	delete(f, "Source")
	delete(f, "AgentIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListAIWorkbenchAgentsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAIWorkbenchAgentsResponseParams struct {
	// <p>Agent list</p>
	Agents []*AgentInfo `json:"Agents,omitnil,omitempty" name:"Agents"`

	// <p>Pagination result</p>
	PageResult *PageByNumResult `json:"PageResult,omitnil,omitempty" name:"PageResult"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListAIWorkbenchAgentsResponse struct {
	*tchttp.BaseResponse
	Response *ListAIWorkbenchAgentsResponseParams `json:"Response"`
}

func (r *ListAIWorkbenchAgentsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAIWorkbenchAgentsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAIWorkbenchArtifactsRequestParams struct {
	// <p>Number of items per page</p>
	PerPage *int64 `json:"PerPage,omitnil,omitempty" name:"PerPage"`

	// <p>Page number.</p>
	PageNo *int64 `json:"PageNo,omitnil,omitempty" name:"PageNo"`

	// <p>Session ID.</p>
	SessionIds []*string `json:"SessionIds,omitnil,omitempty" name:"SessionIds"`

	// <p>Message content type</p>
	MimeTypes []*string `json:"MimeTypes,omitnil,omitempty" name:"MimeTypes"`

	// <p>Sorting order</p><p>Enumeration values:</p><ul><li>ASC: ascending order</li><li>DESC: descending order</li></ul>
	OrderDirection *string `json:"OrderDirection,omitnil,omitempty" name:"OrderDirection"`
}

type ListAIWorkbenchArtifactsRequest struct {
	*tchttp.BaseRequest
	
	// <p>Number of items per page</p>
	PerPage *int64 `json:"PerPage,omitnil,omitempty" name:"PerPage"`

	// <p>Page number.</p>
	PageNo *int64 `json:"PageNo,omitnil,omitempty" name:"PageNo"`

	// <p>Session ID.</p>
	SessionIds []*string `json:"SessionIds,omitnil,omitempty" name:"SessionIds"`

	// <p>Message content type</p>
	MimeTypes []*string `json:"MimeTypes,omitnil,omitempty" name:"MimeTypes"`

	// <p>Sorting order</p><p>Enumeration values:</p><ul><li>ASC: ascending order</li><li>DESC: descending order</li></ul>
	OrderDirection *string `json:"OrderDirection,omitnil,omitempty" name:"OrderDirection"`
}

func (r *ListAIWorkbenchArtifactsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAIWorkbenchArtifactsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PerPage")
	delete(f, "PageNo")
	delete(f, "SessionIds")
	delete(f, "MimeTypes")
	delete(f, "OrderDirection")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListAIWorkbenchArtifactsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAIWorkbenchArtifactsResponseParams struct {
	// <p>Product list</p>
	Artifacts []*ArtifactInfo `json:"Artifacts,omitnil,omitempty" name:"Artifacts"`

	// <p>Pagination result.</p>
	PageResult *PageByNumResult `json:"PageResult,omitnil,omitempty" name:"PageResult"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListAIWorkbenchArtifactsResponse struct {
	*tchttp.BaseResponse
	Response *ListAIWorkbenchArtifactsResponseParams `json:"Response"`
}

func (r *ListAIWorkbenchArtifactsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAIWorkbenchArtifactsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAIWorkbenchExecutionsRequestParams struct {
	// <p>Number of items per page</p>
	PerPage *int64 `json:"PerPage,omitnil,omitempty" name:"PerPage"`

	// <p>Page number.</p>
	PageNo *int64 `json:"PageNo,omitnil,omitempty" name:"PageNo"`

	// <p>Filter by Agent</p>
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// <p>Filter by status</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>Execution ID list filter</p>
	ExecutionIds []*string `json:"ExecutionIds,omitnil,omitempty" name:"ExecutionIds"`

	// <p>Task ID.</p>
	TaskIds []*string `json:"TaskIds,omitnil,omitempty" name:"TaskIds"`

	// <p>Trigger mode</p>
	TriggerType *string `json:"TriggerType,omitnil,omitempty" name:"TriggerType"`

	// <p>Key value</p>
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// <p>Whether to enable</p>
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`
}

type ListAIWorkbenchExecutionsRequest struct {
	*tchttp.BaseRequest
	
	// <p>Number of items per page</p>
	PerPage *int64 `json:"PerPage,omitnil,omitempty" name:"PerPage"`

	// <p>Page number.</p>
	PageNo *int64 `json:"PageNo,omitnil,omitempty" name:"PageNo"`

	// <p>Filter by Agent</p>
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// <p>Filter by status</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>Execution ID list filter</p>
	ExecutionIds []*string `json:"ExecutionIds,omitnil,omitempty" name:"ExecutionIds"`

	// <p>Task ID.</p>
	TaskIds []*string `json:"TaskIds,omitnil,omitempty" name:"TaskIds"`

	// <p>Trigger mode</p>
	TriggerType *string `json:"TriggerType,omitnil,omitempty" name:"TriggerType"`

	// <p>Key value</p>
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// <p>Whether to enable</p>
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`
}

func (r *ListAIWorkbenchExecutionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAIWorkbenchExecutionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PerPage")
	delete(f, "PageNo")
	delete(f, "AgentId")
	delete(f, "Status")
	delete(f, "ExecutionIds")
	delete(f, "TaskIds")
	delete(f, "TriggerType")
	delete(f, "Keyword")
	delete(f, "Enabled")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListAIWorkbenchExecutionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAIWorkbenchExecutionsResponseParams struct {
	// <p>Execution list.</p>
	Executions []*ExecutionInfo `json:"Executions,omitnil,omitempty" name:"Executions"`

	// <p>Pagination result.</p>
	PageResult *PageByNumResult `json:"PageResult,omitnil,omitempty" name:"PageResult"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListAIWorkbenchExecutionsResponse struct {
	*tchttp.BaseResponse
	Response *ListAIWorkbenchExecutionsResponseParams `json:"Response"`
}

func (r *ListAIWorkbenchExecutionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAIWorkbenchExecutionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAIWorkbenchMCPsRequestParams struct {
	// <p>Number of items per page</p>
	PerPage *int64 `json:"PerPage,omitnil,omitempty" name:"PerPage"`

	// <p>Page number.</p>
	PageNo *int64 `json:"PageNo,omitnil,omitempty" name:"PageNo"`

	// <p>Filter by transmission protocol</p>
	Transport *string `json:"Transport,omitnil,omitempty" name:"Transport"`

	// <p>Search keyword</p>
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// <p>Whether to enable filter</p>
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// <p>Associated mcp</p>
	MCPIds []*string `json:"MCPIds,omitnil,omitempty" name:"MCPIds"`

	// <p>MCP type (built-in/private)</p><p>Enumeration values:</p><ul><li>builtin: platform built-in</li><li>private: user-customized</li></ul>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type ListAIWorkbenchMCPsRequest struct {
	*tchttp.BaseRequest
	
	// <p>Number of items per page</p>
	PerPage *int64 `json:"PerPage,omitnil,omitempty" name:"PerPage"`

	// <p>Page number.</p>
	PageNo *int64 `json:"PageNo,omitnil,omitempty" name:"PageNo"`

	// <p>Filter by transmission protocol</p>
	Transport *string `json:"Transport,omitnil,omitempty" name:"Transport"`

	// <p>Search keyword</p>
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// <p>Whether to enable filter</p>
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// <p>Associated mcp</p>
	MCPIds []*string `json:"MCPIds,omitnil,omitempty" name:"MCPIds"`

	// <p>MCP type (built-in/private)</p><p>Enumeration values:</p><ul><li>builtin: platform built-in</li><li>private: user-customized</li></ul>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

func (r *ListAIWorkbenchMCPsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAIWorkbenchMCPsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PerPage")
	delete(f, "PageNo")
	delete(f, "Transport")
	delete(f, "Keyword")
	delete(f, "Enabled")
	delete(f, "MCPIds")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListAIWorkbenchMCPsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAIWorkbenchMCPsResponseParams struct {
	// <p>MCP list</p>
	MCPs []*MCPInfo `json:"MCPs,omitnil,omitempty" name:"MCPs"`

	// <p>Pagination result.</p>
	PageResult *PageByNumResult `json:"PageResult,omitnil,omitempty" name:"PageResult"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListAIWorkbenchMCPsResponse struct {
	*tchttp.BaseResponse
	Response *ListAIWorkbenchMCPsResponseParams `json:"Response"`
}

func (r *ListAIWorkbenchMCPsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAIWorkbenchMCPsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAIWorkbenchMessagesRequestParams struct {
	// <p>Conversation ID</p>
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// <p>Tag for cursor pagination</p>
	Cursor *string `json:"Cursor,omitnil,omitempty" name:"Cursor"`

	// <p>Window size</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>Pull sequence</p>
	Direction *string `json:"Direction,omitnil,omitempty" name:"Direction"`
}

type ListAIWorkbenchMessagesRequest struct {
	*tchttp.BaseRequest
	
	// <p>Conversation ID</p>
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// <p>Tag for cursor pagination</p>
	Cursor *string `json:"Cursor,omitnil,omitempty" name:"Cursor"`

	// <p>Window size</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>Pull sequence</p>
	Direction *string `json:"Direction,omitnil,omitempty" name:"Direction"`
}

func (r *ListAIWorkbenchMessagesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAIWorkbenchMessagesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SessionId")
	delete(f, "Cursor")
	delete(f, "Limit")
	delete(f, "Direction")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListAIWorkbenchMessagesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAIWorkbenchMessagesResponseParams struct {
	// <p>Message list.</p>
	Messages []*MessageInfo `json:"Messages,omitnil,omitempty" name:"Messages"`

	// <p>Next cursor</p>
	NextCursor *string `json:"NextCursor,omitnil,omitempty" name:"NextCursor"`

	// <p>Is there a follow-up?</p>
	HasMore *bool `json:"HasMore,omitnil,omitempty" name:"HasMore"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListAIWorkbenchMessagesResponse struct {
	*tchttp.BaseResponse
	Response *ListAIWorkbenchMessagesResponseParams `json:"Response"`
}

func (r *ListAIWorkbenchMessagesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAIWorkbenchMessagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAIWorkbenchResourceInstancesRequestParams struct {
	// <p>Resource map ID</p>
	ResourceMapId *string `json:"ResourceMapId,omitnil,omitempty" name:"ResourceMapId"`

	// <p>Pagination parameters</p>
	PageParams *PageByNumParams `json:"PageParams,omitnil,omitempty" name:"PageParams"`
}

type ListAIWorkbenchResourceInstancesRequest struct {
	*tchttp.BaseRequest
	
	// <p>Resource map ID</p>
	ResourceMapId *string `json:"ResourceMapId,omitnil,omitempty" name:"ResourceMapId"`

	// <p>Pagination parameters</p>
	PageParams *PageByNumParams `json:"PageParams,omitnil,omitempty" name:"PageParams"`
}

func (r *ListAIWorkbenchResourceInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAIWorkbenchResourceInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ResourceMapId")
	delete(f, "PageParams")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListAIWorkbenchResourceInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAIWorkbenchResourceInstancesResponseParams struct {
	// <p>Resource instance list</p>
	Instances []*ResourceInstance `json:"Instances,omitnil,omitempty" name:"Instances"`

	// <p>Pagination result</p>
	PageResult *PageByNumResult `json:"PageResult,omitnil,omitempty" name:"PageResult"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListAIWorkbenchResourceInstancesResponse struct {
	*tchttp.BaseResponse
	Response *ListAIWorkbenchResourceInstancesResponseParams `json:"Response"`
}

func (r *ListAIWorkbenchResourceInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAIWorkbenchResourceInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAIWorkbenchResourceMapsRequestParams struct {
	// <p>Number of items per page</p>
	PerPage *int64 `json:"PerPage,omitnil,omitempty" name:"PerPage"`

	// <p>Page number.</p>
	PageNo *int64 `json:"PageNo,omitnil,omitempty" name:"PageNo"`

	// <p>Search by name</p>
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`
}

type ListAIWorkbenchResourceMapsRequest struct {
	*tchttp.BaseRequest
	
	// <p>Number of items per page</p>
	PerPage *int64 `json:"PerPage,omitnil,omitempty" name:"PerPage"`

	// <p>Page number.</p>
	PageNo *int64 `json:"PageNo,omitnil,omitempty" name:"PageNo"`

	// <p>Search by name</p>
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`
}

func (r *ListAIWorkbenchResourceMapsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAIWorkbenchResourceMapsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PerPage")
	delete(f, "PageNo")
	delete(f, "Keyword")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListAIWorkbenchResourceMapsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAIWorkbenchResourceMapsResponseParams struct {
	// <p>Resource map list</p>
	ResourceMaps []*ResourceMapInfo `json:"ResourceMaps,omitnil,omitempty" name:"ResourceMaps"`

	// <p>Pagination result.</p>
	PageResult *PageByNumResult `json:"PageResult,omitnil,omitempty" name:"PageResult"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListAIWorkbenchResourceMapsResponse struct {
	*tchttp.BaseResponse
	Response *ListAIWorkbenchResourceMapsResponseParams `json:"Response"`
}

func (r *ListAIWorkbenchResourceMapsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAIWorkbenchResourceMapsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAIWorkbenchSessionsRequestParams struct {
	// <p>Number of items per page</p>
	PerPage *int64 `json:"PerPage,omitnil,omitempty" name:"PerPage"`

	// <p>Page number.</p>
	PageNo *int64 `json:"PageNo,omitnil,omitempty" name:"PageNo"`

	// <p>Filter by Agent</p>
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// <p>Search keyword</p>
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// <p>Session ID list filtering</p>
	SessionIds []*string `json:"SessionIds,omitnil,omitempty" name:"SessionIds"`
}

type ListAIWorkbenchSessionsRequest struct {
	*tchttp.BaseRequest
	
	// <p>Number of items per page</p>
	PerPage *int64 `json:"PerPage,omitnil,omitempty" name:"PerPage"`

	// <p>Page number.</p>
	PageNo *int64 `json:"PageNo,omitnil,omitempty" name:"PageNo"`

	// <p>Filter by Agent</p>
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// <p>Search keyword</p>
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// <p>Session ID list filtering</p>
	SessionIds []*string `json:"SessionIds,omitnil,omitempty" name:"SessionIds"`
}

func (r *ListAIWorkbenchSessionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAIWorkbenchSessionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PerPage")
	delete(f, "PageNo")
	delete(f, "AgentId")
	delete(f, "Keyword")
	delete(f, "SessionIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListAIWorkbenchSessionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAIWorkbenchSessionsResponseParams struct {
	// <p>Session list</p>
	Sessions []*SessionInfo `json:"Sessions,omitnil,omitempty" name:"Sessions"`

	// <p>Pagination result</p>
	PageResult *PageByNumResult `json:"PageResult,omitnil,omitempty" name:"PageResult"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListAIWorkbenchSessionsResponse struct {
	*tchttp.BaseResponse
	Response *ListAIWorkbenchSessionsResponseParams `json:"Response"`
}

func (r *ListAIWorkbenchSessionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAIWorkbenchSessionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAIWorkbenchSkillsRequestParams struct {
	// <p>Number of items per page</p>
	PerPage *int64 `json:"PerPage,omitnil,omitempty" name:"PerPage"`

	// <p>Page number.</p>
	PageNo *int64 `json:"PageNo,omitnil,omitempty" name:"PageNo"`

	// <p>Filter by type</p>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>Search keyword</p>
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// <p>Whether to enable filter</p>
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// <p>Skill ID list filter</p>
	SkillIds []*string `json:"SkillIds,omitnil,omitempty" name:"SkillIds"`
}

type ListAIWorkbenchSkillsRequest struct {
	*tchttp.BaseRequest
	
	// <p>Number of items per page</p>
	PerPage *int64 `json:"PerPage,omitnil,omitempty" name:"PerPage"`

	// <p>Page number.</p>
	PageNo *int64 `json:"PageNo,omitnil,omitempty" name:"PageNo"`

	// <p>Filter by type</p>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>Search keyword</p>
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// <p>Whether to enable filter</p>
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// <p>Skill ID list filter</p>
	SkillIds []*string `json:"SkillIds,omitnil,omitempty" name:"SkillIds"`
}

func (r *ListAIWorkbenchSkillsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAIWorkbenchSkillsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PerPage")
	delete(f, "PageNo")
	delete(f, "Type")
	delete(f, "Keyword")
	delete(f, "Enabled")
	delete(f, "SkillIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListAIWorkbenchSkillsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAIWorkbenchSkillsResponseParams struct {
	// <p>List of skills</p>
	Skills []*SkillInfo `json:"Skills,omitnil,omitempty" name:"Skills"`

	// <p>Pagination result</p>
	PageResult *PageByNumResult `json:"PageResult,omitnil,omitempty" name:"PageResult"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListAIWorkbenchSkillsResponse struct {
	*tchttp.BaseResponse
	Response *ListAIWorkbenchSkillsResponseParams `json:"Response"`
}

func (r *ListAIWorkbenchSkillsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAIWorkbenchSkillsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAIWorkbenchTasksRequestParams struct {
	// <p>Number of items per page</p>
	PerPage *int64 `json:"PerPage,omitnil,omitempty" name:"PerPage"`

	// <p>Page number.</p>
	PageNo *int64 `json:"PageNo,omitnil,omitempty" name:"PageNo"`

	// <p>Filter by Agent</p>
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// <p>Filter by trigger type</p>
	TriggerType *string `json:"TriggerType,omitnil,omitempty" name:"TriggerType"`

	// <p>Search keyword</p>
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// <p>Task ID list filter</p>
	TaskIds []*string `json:"TaskIds,omitnil,omitempty" name:"TaskIds"`

	// <p>Whether to enable filter criteria</p>
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`
}

type ListAIWorkbenchTasksRequest struct {
	*tchttp.BaseRequest
	
	// <p>Number of items per page</p>
	PerPage *int64 `json:"PerPage,omitnil,omitempty" name:"PerPage"`

	// <p>Page number.</p>
	PageNo *int64 `json:"PageNo,omitnil,omitempty" name:"PageNo"`

	// <p>Filter by Agent</p>
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// <p>Filter by trigger type</p>
	TriggerType *string `json:"TriggerType,omitnil,omitempty" name:"TriggerType"`

	// <p>Search keyword</p>
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// <p>Task ID list filter</p>
	TaskIds []*string `json:"TaskIds,omitnil,omitempty" name:"TaskIds"`

	// <p>Whether to enable filter criteria</p>
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`
}

func (r *ListAIWorkbenchTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAIWorkbenchTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PerPage")
	delete(f, "PageNo")
	delete(f, "AgentId")
	delete(f, "TriggerType")
	delete(f, "Keyword")
	delete(f, "TaskIds")
	delete(f, "Enabled")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListAIWorkbenchTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAIWorkbenchTasksResponseParams struct {
	// <p>Task List</p>
	Tasks []*TaskInfo `json:"Tasks,omitnil,omitempty" name:"Tasks"`

	// <p>Pagination result</p>
	PageResult *PageByNumResult `json:"PageResult,omitnil,omitempty" name:"PageResult"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListAIWorkbenchTasksResponse struct {
	*tchttp.BaseResponse
	Response *ListAIWorkbenchTasksResponseParams `json:"Response"`
}

func (r *ListAIWorkbenchTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAIWorkbenchTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MCPInfo struct {
	// <p>mcp ID</p>
	MCPId *string `json:"MCPId,omitnil,omitempty" name:"MCPId"`

	// <p>MCP name</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>MCP description</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>MCP URL</p>
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// <p>Transport protocol: sse / streamable_http / stdio</p>
	Transport *string `json:"Transport,omitnil,omitempty" name:"Transport"`

	// <p>Authentication type: none / bearer / basic / api_key</p>
	AuthType *string `json:"AuthType,omitnil,omitempty" name:"AuthType"`

	// <p>Authentication key (masked in the response)</p>
	AuthSecret *string `json:"AuthSecret,omitnil,omitempty" name:"AuthSecret"`

	// <p>Timeout (s)</p>
	Timeout *int64 `json:"Timeout,omitnil,omitempty" name:"Timeout"`

	// <p>Retry count</p>
	RetryCount *int64 `json:"RetryCount,omitnil,omitempty" name:"RetryCount"`

	// <p>Request header JSON</p>
	Headers *string `json:"Headers,omitnil,omitempty" name:"Headers"`

	// <p>Whether to enable</p>
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`
}

type MessageInfo struct {
	// <p>Entity id</p>
	EntryId *string `json:"EntryId,omitnil,omitempty" name:"EntryId"`

	// <p>Conversation ID</p>
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// <p>Role: user / assistant</p>
	Role *string `json:"Role,omitnil,omitempty" name:"Role"`

	// <p>Message content</p>
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// <p>Status.</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>Block content.</p>
	ContentBlocks []*ContentBlockInfo `json:"ContentBlocks,omitnil,omitempty" name:"ContentBlocks"`
}

type NotifyRelatedNotice struct {
	// Notification template ID
	NoticeId *string `json:"NoticeId,omitnil,omitempty" name:"NoticeId"`

	// Name of the notification template
	NoticeName *string `json:"NoticeName,omitnil,omitempty" name:"NoticeName"`
}

type PageByNoParams struct {
	// Number of items per page.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PerPage *int64 `json:"PerPage,omitnil,omitempty" name:"PerPage"`

	// Page number, starting from 1.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PageNo *string `json:"PageNo,omitnil,omitempty" name:"PageNo"`
}

type PageByNoResult struct {
	// Total data.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Total number of pages.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TotalPage *int64 `json:"TotalPage,omitnil,omitempty" name:"TotalPage"`

	// Current page number.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CurrentPageNo *int64 `json:"CurrentPageNo,omitnil,omitempty" name:"CurrentPageNo"`

	// [Deprecated] Whether it has reached the end.
	// Note: This field may return null, indicating that no valid values can be obtained.
	//
	// Deprecated: IsEnd is deprecated.
	IsEnd *bool `json:"IsEnd,omitnil,omitempty" name:"IsEnd"`

	// Whether it has traversed to the end.
	End *bool `json:"End,omitnil,omitempty" name:"End"`
}

type PageByNumParams struct {
	// <p>Number of items per page</p>
	PerPage *int64 `json:"PerPage,omitnil,omitempty" name:"PerPage"`

	// <p>Page number, starting from 1</p>
	PageNo *int64 `json:"PageNo,omitnil,omitempty" name:"PageNo"`
}

type PageByNumResult struct {
	// <p>Total number of data</p>
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// <p>Total number of pages</p>
	TotalPage *int64 `json:"TotalPage,omitnil,omitempty" name:"TotalPage"`

	// <p>Current page number</p>
	CurrentPageNo *int64 `json:"CurrentPageNo,omitnil,omitempty" name:"CurrentPageNo"`
}

type ResourceInstance struct {
	// <p>Instance ID</p>
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// <p>Service name</p>
	Service *string `json:"Service,omitnil,omitempty" name:"Service"`

	// <p>Region.</p>
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// <p>Ready?</p>
	IsReady *bool `json:"IsReady,omitnil,omitempty" name:"IsReady"`
}

type ResourceMapInfo struct {
	// <p>Resource map ID</p>
	ResourceMapId *string `json:"ResourceMapId,omitnil,omitempty" name:"ResourceMapId"`

	// <p>Resource map name</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Resource map description</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>Total number of instances</p>
	InstanceCount *int64 `json:"InstanceCount,omitnil,omitempty" name:"InstanceCount"`
}

type SessionInfo struct {
	// <p>Session ID</p>
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// <p>Agent ID</p>
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// <p>Session title</p>
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`

	// <p>Status: active / archived / deleted</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>If the session is triggered by a task, carry the task ID that triggers the session.</p>
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type SkillInfo struct {
	// <p>Skill ID</p>
	SkillId *string `json:"SkillId,omitnil,omitempty" name:"SkillId"`

	// <p>Skill name</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Skill description.</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>Whether to enable</p>
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`
}

type Tag struct {
	// Tag key
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// Tag value
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type TaskInfo struct {
	// <p>Task ID.</p>
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// <p>Task name</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Task description</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>Associated Agent ID</p>
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// <p>Prompt Template</p>
	PromptTemplate *string `json:"PromptTemplate,omitnil,omitempty" name:"PromptTemplate"`

	// <p>Output format: markdown / json</p>
	OutputFormat *string `json:"OutputFormat,omitnil,omitempty" name:"OutputFormat"`

	// <p>Trigger type: manual / cron / webhook</p>
	TriggerType *string `json:"TriggerType,omitnil,omitempty" name:"TriggerType"`

	// <p>Cron expression</p>
	CronExpr *string `json:"CronExpr,omitnil,omitempty" name:"CronExpr"`

	// <p>Cron time zone</p>
	CronTimezone *string `json:"CronTimezone,omitnil,omitempty" name:"CronTimezone"`

	// <p>List of associated skill IDs.</p>
	SkillIds []*string `json:"SkillIds,omitnil,omitempty" name:"SkillIds"`

	// <p>Associated MCP endpoint ID list</p>
	McpEndpointIds []*string `json:"McpEndpointIds,omitnil,omitempty" name:"McpEndpointIds"`

	// <p>Timeout (seconds)</p>
	TimeoutSec *int64 `json:"TimeoutSec,omitnil,omitempty" name:"TimeoutSec"`

	// <p>Retry count</p>
	RetryCount *int64 `json:"RetryCount,omitnil,omitempty" name:"RetryCount"`

	// <p>Notification id</p>
	NotifyIds []*string `json:"NotifyIds,omitnil,omitempty" name:"NotifyIds"`

	// <p>Whether to enable</p>
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`
}

// Predefined struct for user
type TriggerAIWorkbenchTaskRequestParams struct {
	// <p>Task ID.</p>
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type TriggerAIWorkbenchTaskRequest struct {
	*tchttp.BaseRequest
	
	// <p>Task ID.</p>
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *TriggerAIWorkbenchTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TriggerAIWorkbenchTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TriggerAIWorkbenchTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TriggerAIWorkbenchTaskResponseParams struct {
	// <p>Execution ID.</p>
	ExecutionId *string `json:"ExecutionId,omitnil,omitempty" name:"ExecutionId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type TriggerAIWorkbenchTaskResponse struct {
	*tchttp.BaseResponse
	Response *TriggerAIWorkbenchTaskResponseParams `json:"Response"`
}

func (r *TriggerAIWorkbenchTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TriggerAIWorkbenchTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateAIWorkbenchAgentRequestParams struct {
	// <p>Agent ID</p>
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// <p>Agent name</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Agent description</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>Agent Category.</p>
	Category *string `json:"Category,omitnil,omitempty" name:"Category"`

	// <p>Agent Tag.</p>
	Tags []*string `json:"Tags,omitnil,omitempty" name:"Tags"`

	// <p>Agent prompt</p>
	Instruction *InstructionConfig `json:"Instruction,omitnil,omitempty" name:"Instruction"`

	// <p>List of associated skill IDs.</p>
	SkillIds []*string `json:"SkillIds,omitnil,omitempty" name:"SkillIds"`

	// <p>Source</p>
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// <p>Status.</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>ID of the associated resource map</p>
	ResourceMapId *string `json:"ResourceMapId,omitnil,omitempty" name:"ResourceMapId"`

	// <p>Associated mcp</p>
	MCPIds []*string `json:"MCPIds,omitnil,omitempty" name:"MCPIds"`

	// <p>Environment variables required by the agent at runtime</p>
	EnvVars []*EnvVar `json:"EnvVars,omitnil,omitempty" name:"EnvVars"`
}

type UpdateAIWorkbenchAgentRequest struct {
	*tchttp.BaseRequest
	
	// <p>Agent ID</p>
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// <p>Agent name</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Agent description</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>Agent Category.</p>
	Category *string `json:"Category,omitnil,omitempty" name:"Category"`

	// <p>Agent Tag.</p>
	Tags []*string `json:"Tags,omitnil,omitempty" name:"Tags"`

	// <p>Agent prompt</p>
	Instruction *InstructionConfig `json:"Instruction,omitnil,omitempty" name:"Instruction"`

	// <p>List of associated skill IDs.</p>
	SkillIds []*string `json:"SkillIds,omitnil,omitempty" name:"SkillIds"`

	// <p>Source</p>
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// <p>Status.</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>ID of the associated resource map</p>
	ResourceMapId *string `json:"ResourceMapId,omitnil,omitempty" name:"ResourceMapId"`

	// <p>Associated mcp</p>
	MCPIds []*string `json:"MCPIds,omitnil,omitempty" name:"MCPIds"`

	// <p>Environment variables required by the agent at runtime</p>
	EnvVars []*EnvVar `json:"EnvVars,omitnil,omitempty" name:"EnvVars"`
}

func (r *UpdateAIWorkbenchAgentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateAIWorkbenchAgentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AgentId")
	delete(f, "Name")
	delete(f, "Description")
	delete(f, "Category")
	delete(f, "Tags")
	delete(f, "Instruction")
	delete(f, "SkillIds")
	delete(f, "Source")
	delete(f, "Status")
	delete(f, "ResourceMapId")
	delete(f, "MCPIds")
	delete(f, "EnvVars")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateAIWorkbenchAgentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateAIWorkbenchAgentResponseParams struct {
	// <p>Agent information after the update</p>
	Agent *AgentInfo `json:"Agent,omitnil,omitempty" name:"Agent"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateAIWorkbenchAgentResponse struct {
	*tchttp.BaseResponse
	Response *UpdateAIWorkbenchAgentResponseParams `json:"Response"`
}

func (r *UpdateAIWorkbenchAgentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateAIWorkbenchAgentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}