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

package v20200210

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

type AIAnalysisResult struct {
	// Summary: describes the session summary.
	// mood: specifies the emotion analysis.
	// intention extraction.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// AI session analysis result.
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`
}

type AICallExtractConfigElement struct {
	// Configuration item type, including.
	// Text.
	// Selector option.
	// Boolean value.
	// Number.
	InfoType *string `json:"InfoType,omitnil,omitempty" name:"InfoType"`

	// Configuration item name, duplicat.
	InfoName *string `json:"InfoName,omitnil,omitempty" name:"InfoName"`

	// Specific content of the configuration item.
	InfoContent *string `json:"InfoContent,omitnil,omitempty" name:"InfoContent"`

	// Example of extracted content from the configuration item.
	Examples []*string `json:"Examples,omitnil,omitempty" name:"Examples"`

	// When infotype is selector, this field needs to be configured.
	Choices []*string `json:"Choices,omitnil,omitempty" name:"Choices"`
}

type AICallExtractResultElement struct {
	// Type of extracted information.
	// Text.
	// Selector options.
	// Boolean value.
	// Number.
	InfoType *string `json:"InfoType,omitnil,omitempty" name:"InfoType"`

	// Name of the extracted information.
	InfoName *string `json:"InfoName,omitnil,omitempty" name:"InfoName"`

	// Specific description of the extracted information.
	InfoContent *string `json:"InfoContent,omitnil,omitempty" name:"InfoContent"`

	// Specific result of the extracted information.
	Result *AICallExtractResultInfo `json:"Result,omitnil,omitempty" name:"Result"`
}

type AICallExtractResultInfo struct {
	// The extracted type is text.
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`

	// The extracted type is option.
	Chosen []*string `json:"Chosen,omitnil,omitempty" name:"Chosen"`

	// The extracted type is a boolean value.
	Boolean *bool `json:"Boolean,omitnil,omitempty" name:"Boolean"`

	// The extracted type is a number.
	Number *float64 `json:"Number,omitnil,omitempty" name:"Number"`
}

type AILatencyDetail struct {
	// Dialog ID.
	RoundId *string `json:"RoundId,omitnil,omitempty" name:"RoundId"`

	// Specifies the asr latency in milliseconds.
	ASRLatency *int64 `json:"ASRLatency,omitnil,omitempty" name:"ASRLatency"`

	// Specifies the tts delay in milliseconds.
	TTSLatency *int64 `json:"TTSLatency,omitnil,omitempty" name:"TTSLatency"`

	// llm latency (ms).
	LLMLatency *int64 `json:"LLMLatency,omitnil,omitempty" name:"LLMLatency"`

	// llm first token latency (ms).
	LLMFirstTokenLatency *int64 `json:"LLMFirstTokenLatency,omitnil,omitempty" name:"LLMFirstTokenLatency"`

	// End-To-End delay (ms).
	ETELatency *int64 `json:"ETELatency,omitnil,omitempty" name:"ETELatency"`
}

type AILatencyStatistics struct {
	// Specifies the asr latency statistics.
	ASRLatency *AILatencyStatisticsInfo `json:"ASRLatency,omitnil,omitempty" name:"ASRLatency"`

	// Specifies the statistics of tts delay.
	TTSLatency *AILatencyStatisticsInfo `json:"TTSLatency,omitnil,omitempty" name:"TTSLatency"`

	// llm latency statistics.
	LLMLatency *AILatencyStatisticsInfo `json:"LLMLatency,omitnil,omitempty" name:"LLMLatency"`

	// Specifies the end-to-end latency statistics.
	ETELatency *AILatencyStatisticsInfo `json:"ETELatency,omitnil,omitempty" name:"ETELatency"`
}

type AILatencyStatisticsInfo struct {
	// Specifies the minimum value.
	MinLatency *int64 `json:"MinLatency,omitnil,omitempty" name:"MinLatency"`

	// Specifies the median.
	MiddleLatency *int64 `json:"MiddleLatency,omitnil,omitempty" name:"MiddleLatency"`

	// p90
	P90Latency *int64 `json:"P90Latency,omitnil,omitempty" name:"P90Latency"`
}

type AITransferItem struct {
	// Name of the function calling for transfer to human.
	TransferFunctionName *string `json:"TransferFunctionName,omitnil,omitempty" name:"TransferFunctionName"`

	// Takes effect when transferfunctionenable is true; the description of transfer_to_human function calling defaults to "transfer to human when the user has to transfer to human (like says transfer to human) or you are instructed to do so.".
	TransferFunctionDesc *string `json:"TransferFunctionDesc,omitnil,omitempty" name:"TransferFunctionDesc"`

	// Skill group id for transferring to human agent.
	TransferSkillGroupId *uint64 `json:"TransferSkillGroupId,omitnil,omitempty" name:"TransferSkillGroupId"`
}

// Predefined struct for user
type AbortAgentCruiseDialingCampaignRequestParams struct {
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// <Task id>.
	CampaignId *uint64 `json:"CampaignId,omitnil,omitempty" name:"CampaignId"`
}

type AbortAgentCruiseDialingCampaignRequest struct {
	*tchttp.BaseRequest
	
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// <Task id>.
	CampaignId *uint64 `json:"CampaignId,omitnil,omitempty" name:"CampaignId"`
}

func (r *AbortAgentCruiseDialingCampaignRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AbortAgentCruiseDialingCampaignRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "CampaignId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AbortAgentCruiseDialingCampaignRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AbortAgentCruiseDialingCampaignResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AbortAgentCruiseDialingCampaignResponse struct {
	*tchttp.BaseResponse
	Response *AbortAgentCruiseDialingCampaignResponseParams `json:"Response"`
}

func (r *AbortAgentCruiseDialingCampaignResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AbortAgentCruiseDialingCampaignResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AbortPredictiveDialingCampaignRequestParams struct {
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Task id.
	CampaignId *int64 `json:"CampaignId,omitnil,omitempty" name:"CampaignId"`
}

type AbortPredictiveDialingCampaignRequest struct {
	*tchttp.BaseRequest
	
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Task id.
	CampaignId *int64 `json:"CampaignId,omitnil,omitempty" name:"CampaignId"`
}

func (r *AbortPredictiveDialingCampaignRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AbortPredictiveDialingCampaignRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "CampaignId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AbortPredictiveDialingCampaignRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AbortPredictiveDialingCampaignResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AbortPredictiveDialingCampaignResponse struct {
	*tchttp.BaseResponse
	Response *AbortPredictiveDialingCampaignResponseParams `json:"Response"`
}

func (r *AbortPredictiveDialingCampaignResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AbortPredictiveDialingCampaignResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AsrData struct {
	// User side.
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// Message content.
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// Timestamp.
	//
	// Deprecated: Timestamp is deprecated.
	Timestamp *int64 `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`

	// Sentence start time, unix millisecond timestamp.
	Start *int64 `json:"Start,omitnil,omitempty" name:"Start"`

	// Sentence end time, unix millisecond timestamp.
	End *int64 `json:"End,omitnil,omitempty" name:"End"`
}

type AudioFileInfo struct {
	// File id.
	FileId *uint64 `json:"FileId,omitnil,omitempty" name:"FileId"`

	// File alias.
	CustomFileName *string `json:"CustomFileName,omitnil,omitempty" name:"CustomFileName"`

	// Filename.
	AudioFileName *string `json:"AudioFileName,omitnil,omitempty" name:"AudioFileName"`

	// Review status: 0 - unreviewed, 1 - approved, 2 - rejected.
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type AutoCalloutTaskCalleeInfo struct {
	// Called number.
	Callee *string `json:"Callee,omitnil,omitempty" name:"Callee"`

	// Call status 0 - initial, 1 - answered, 2 - unanswered, 3 - calling, 4 - pending retry.
	State *uint64 `json:"State,omitnil,omitempty" name:"State"`

	// List of session ids.
	Sessions []*string `json:"Sessions,omitnil,omitempty" name:"Sessions"`
}

type AutoCalloutTaskInfo struct {
	// Task name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Number of called parties.
	CalleeCount *uint64 `json:"CalleeCount,omitnil,omitempty" name:"CalleeCount"`

	// List of calling numbers.
	Callers []*string `json:"Callers,omitnil,omitempty" name:"Callers"`

	// Start timestamp.
	NotBefore *int64 `json:"NotBefore,omitnil,omitempty" name:"NotBefore"`

	// End timestamp
	// .
	// Note: this field may return null, indicating that no valid values can be obtained.
	NotAfter *int64 `json:"NotAfter,omitnil,omitempty" name:"NotAfter"`

	// IvrId used by the task.
	IvrId *uint64 `json:"IvrId,omitnil,omitempty" name:"IvrId"`

	// Task status:.
	// 0 initial: task creation, call not started.
	// 1 running.
	// 2 completed: all calls in the task are completed.
	// 3 ending: the task has expired, but there are still some calls not ended.
	// 4 ended: task terminated due to expiration.
	State *uint64 `json:"State,omitnil,omitempty" name:"State"`

	// <Task id>.
	TaskId *uint64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

// Predefined struct for user
type BindNumberCallInInterfaceRequestParams struct {
	// App ID (required). can be used to view https://console.cloud.tencent.com/ccc.
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Number to be bound.
	Number *string `json:"Number,omitnil,omitempty" name:"Number"`

	// Specifies the callback url to be bound.
	CallInInterface *Interface `json:"CallInInterface,omitnil,omitempty" name:"CallInInterface"`
}

type BindNumberCallInInterfaceRequest struct {
	*tchttp.BaseRequest
	
	// App ID (required). can be used to view https://console.cloud.tencent.com/ccc.
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Number to be bound.
	Number *string `json:"Number,omitnil,omitempty" name:"Number"`

	// Specifies the callback url to be bound.
	CallInInterface *Interface `json:"CallInInterface,omitnil,omitempty" name:"CallInInterface"`
}

func (r *BindNumberCallInInterfaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindNumberCallInInterfaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "Number")
	delete(f, "CallInInterface")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindNumberCallInInterfaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindNumberCallInInterfaceResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type BindNumberCallInInterfaceResponse struct {
	*tchttp.BaseResponse
	Response *BindNumberCallInInterfaceResponseParams `json:"Response"`
}

func (r *BindNumberCallInInterfaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindNumberCallInInterfaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindNumberCallOutSkillGroupRequestParams struct {
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Number to be bound.
	Number *string `json:"Number,omitnil,omitempty" name:"Number"`

	// Skill group id list to be bound.
	SkillGroupIds []*uint64 `json:"SkillGroupIds,omitnil,omitempty" name:"SkillGroupIds"`
}

type BindNumberCallOutSkillGroupRequest struct {
	*tchttp.BaseRequest
	
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Number to be bound.
	Number *string `json:"Number,omitnil,omitempty" name:"Number"`

	// Skill group id list to be bound.
	SkillGroupIds []*uint64 `json:"SkillGroupIds,omitnil,omitempty" name:"SkillGroupIds"`
}

func (r *BindNumberCallOutSkillGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindNumberCallOutSkillGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "Number")
	delete(f, "SkillGroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindNumberCallOutSkillGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindNumberCallOutSkillGroupResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type BindNumberCallOutSkillGroupResponse struct {
	*tchttp.BaseResponse
	Response *BindNumberCallOutSkillGroupResponseParams `json:"Response"`
}

func (r *BindNumberCallOutSkillGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindNumberCallOutSkillGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindStaffSkillGroupListRequestParams struct {
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Agent email.
	StaffEmail *string `json:"StaffEmail,omitnil,omitempty" name:"StaffEmail"`

	// Bound skill group list.
	//
	// Deprecated: SkillGroupList is deprecated.
	SkillGroupList []*int64 `json:"SkillGroupList,omitnil,omitempty" name:"SkillGroupList"`

	// Bound skill group list (required).
	StaffSkillGroupList []*StaffSkillGroupList `json:"StaffSkillGroupList,omitnil,omitempty" name:"StaffSkillGroupList"`
}

type BindStaffSkillGroupListRequest struct {
	*tchttp.BaseRequest
	
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Agent email.
	StaffEmail *string `json:"StaffEmail,omitnil,omitempty" name:"StaffEmail"`

	// Bound skill group list.
	SkillGroupList []*int64 `json:"SkillGroupList,omitnil,omitempty" name:"SkillGroupList"`

	// Bound skill group list (required).
	StaffSkillGroupList []*StaffSkillGroupList `json:"StaffSkillGroupList,omitnil,omitempty" name:"StaffSkillGroupList"`
}

func (r *BindStaffSkillGroupListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindStaffSkillGroupListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "StaffEmail")
	delete(f, "SkillGroupList")
	delete(f, "StaffSkillGroupList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindStaffSkillGroupListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindStaffSkillGroupListResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type BindStaffSkillGroupListResponse struct {
	*tchttp.BaseResponse
	Response *BindStaffSkillGroupListResponseParams `json:"Response"`
}

func (r *BindStaffSkillGroupListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindStaffSkillGroupListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CallInMetrics struct {
	// Number of ivr residency.
	IvrCount *int64 `json:"IvrCount,omitnil,omitempty" name:"IvrCount"`

	// Number in queue.
	QueueCount *int64 `json:"QueueCount,omitnil,omitempty" name:"QueueCount"`

	// Number in ringing.
	RingCount *int64 `json:"RingCount,omitnil,omitempty" name:"RingCount"`

	// Number of connections.
	AcceptCount *int64 `json:"AcceptCount,omitnil,omitempty" name:"AcceptCount"`

	// Number of customer service transferring to the external line.
	TransferOuterCount *int64 `json:"TransferOuterCount,omitnil,omitempty" name:"TransferOuterCount"`

	// Maximum queue duration.
	MaxQueueDuration *int64 `json:"MaxQueueDuration,omitnil,omitempty" name:"MaxQueueDuration"`

	// Average queue duration.
	AvgQueueDuration *int64 `json:"AvgQueueDuration,omitnil,omitempty" name:"AvgQueueDuration"`

	// Maximum ringing duration.
	MaxRingDuration *int64 `json:"MaxRingDuration,omitnil,omitempty" name:"MaxRingDuration"`

	// Average ringing duration.
	AvgRingDuration *int64 `json:"AvgRingDuration,omitnil,omitempty" name:"AvgRingDuration"`

	// Maximum connection duration.
	MaxAcceptDuration *int64 `json:"MaxAcceptDuration,omitnil,omitempty" name:"MaxAcceptDuration"`

	// Average connection duration.
	AvgAcceptDuration *int64 `json:"AvgAcceptDuration,omitnil,omitempty" name:"AvgAcceptDuration"`
}

type CallInNumberMetrics struct {
	// Line number.
	Number *string `json:"Number,omitnil,omitempty" name:"Number"`

	// Line-Related metrics.
	Metrics *CallInMetrics `json:"Metrics,omitnil,omitempty" name:"Metrics"`

	// Bound skill group metrics.
	SkillGroupMetrics []*CallInSkillGroupMetrics `json:"SkillGroupMetrics,omitnil,omitempty" name:"SkillGroupMetrics"`
}

type CallInSkillGroupMetrics struct {
	// Skill group id.
	SkillGroupId *int64 `json:"SkillGroupId,omitnil,omitempty" name:"SkillGroupId"`

	// Data metrics.
	Metrics *CallInMetrics `json:"Metrics,omitnil,omitempty" name:"Metrics"`

	// Skill group name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type CalleeAttribute struct {
	// Called number.
	Callee *string `json:"Callee,omitnil,omitempty" name:"Callee"`

	// Accompanying data.
	UUI *string `json:"UUI,omitnil,omitempty" name:"UUI"`

	// Parameter.
	Variables []*Variable `json:"Variables,omitnil,omitempty" name:"Variables"`
}

type ClientInfo struct {
	// Endpoint type for login. "Web" means Web workbench. "WeChatMiniProgram" refers to wechat mini program.
	ClientType *string `json:"ClientType,omitnil,omitempty" name:"ClientType"`

	// Whether the currently logged-in endpoint is in the foreground. if the endpoint is Web, the value is true. if the endpoint is WeChatMiniProgram, true indicates the wechat mini program is open, and false indicates it is in the background.
	IsConnected *bool `json:"IsConnected,omitnil,omitempty" name:"IsConnected"`
}

// Predefined struct for user
type ControlAIConversationRequestParams struct {
	// Specifies the session ID.
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// App ID (required). can be used to view https://console.cloud.tencent.com/ccc.
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Control command. currently supports the following commands:.
	// 
	// -ServerPushText. specifies the text sent by the server to the AI robot. the AI robot will broadcast the text.
	Command *string `json:"Command,omitnil,omitempty" name:"Command"`

	// Specifies the server-sent broadcast text Command. required when Command is ServerPushText.
	ServerPushText *ServerPushText `json:"ServerPushText,omitnil,omitempty" name:"ServerPushText"`
}

type ControlAIConversationRequest struct {
	*tchttp.BaseRequest
	
	// Specifies the session ID.
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// App ID (required). can be used to view https://console.cloud.tencent.com/ccc.
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Control command. currently supports the following commands:.
	// 
	// -ServerPushText. specifies the text sent by the server to the AI robot. the AI robot will broadcast the text.
	Command *string `json:"Command,omitnil,omitempty" name:"Command"`

	// Specifies the server-sent broadcast text Command. required when Command is ServerPushText.
	ServerPushText *ServerPushText `json:"ServerPushText,omitnil,omitempty" name:"ServerPushText"`
}

func (r *ControlAIConversationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ControlAIConversationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SessionId")
	delete(f, "SdkAppId")
	delete(f, "Command")
	delete(f, "ServerPushText")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ControlAIConversationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ControlAIConversationResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ControlAIConversationResponse struct {
	*tchttp.BaseResponse
	Response *ControlAIConversationResponseParams `json:"Response"`
}

func (r *ControlAIConversationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ControlAIConversationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAIAgentCallRequestParams struct {
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// AI agent id.
	AIAgentId *uint64 `json:"AIAgentId,omitnil,omitempty" name:"AIAgentId"`

	// Callee number.
	Callee *string `json:"Callee,omitnil,omitempty" name:"Callee"`

	// Caller number list
	Callers []*string `json:"Callers,omitnil,omitempty" name:"Callers"`

	// Prompt variable.
	//
	// Deprecated: PromptVariables is deprecated.
	PromptVariables []*Variable `json:"PromptVariables,omitnil,omitempty" name:"PromptVariables"`

	// <P>Prompt variable</p> <p>welcome message variable</p> <p>welcome message delay playback (in seconds): welcome-message-delay</p> <p>dify variable</p>.  
	// 
	// dify-inputs-xxx specifies the inputs variable for dify.
	// 2. the dify-inputs-user specifies the user value for dify.
	// 3. dify-inputs-conversation_id is the conversation_id value of dify.
	Variables []*Variable `json:"Variables,omitnil,omitempty" name:"Variables"`
}

type CreateAIAgentCallRequest struct {
	*tchttp.BaseRequest
	
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// AI agent id.
	AIAgentId *uint64 `json:"AIAgentId,omitnil,omitempty" name:"AIAgentId"`

	// Callee number.
	Callee *string `json:"Callee,omitnil,omitempty" name:"Callee"`

	// Caller number list
	Callers []*string `json:"Callers,omitnil,omitempty" name:"Callers"`

	// Prompt variable.
	PromptVariables []*Variable `json:"PromptVariables,omitnil,omitempty" name:"PromptVariables"`

	// <P>Prompt variable</p> <p>welcome message variable</p> <p>welcome message delay playback (in seconds): welcome-message-delay</p> <p>dify variable</p>.  
	// 
	// dify-inputs-xxx specifies the inputs variable for dify.
	// 2. the dify-inputs-user specifies the user value for dify.
	// 3. dify-inputs-conversation_id is the conversation_id value of dify.
	Variables []*Variable `json:"Variables,omitnil,omitempty" name:"Variables"`
}

func (r *CreateAIAgentCallRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAIAgentCallRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "AIAgentId")
	delete(f, "Callee")
	delete(f, "Callers")
	delete(f, "PromptVariables")
	delete(f, "Variables")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAIAgentCallRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAIAgentCallResponseParams struct {
	// Newly created session id.
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAIAgentCallResponse struct {
	*tchttp.BaseResponse
	Response *CreateAIAgentCallResponseParams `json:"Response"`
}

func (r *CreateAIAgentCallResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAIAgentCallResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAICallRequestParams struct {
	// Application ID (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Called number.
	Callee *string `json:"Callee,omitnil,omitempty" name:"Callee"`

	// Model API protocol type. currently compatible with four protocol types:.
	// 
	// -OpenAI protocol (including GPT, hunyuan, DeepSeek, etc.): "OpenAI".
	// -Azure protocol: "azure".
	// -Specifies the "Minimax" protocol.
	// -Dify protocol: "dify".
	LLMType *string `json:"LLMType,omitnil,omitempty" name:"LLMType"`

	// Model API key, for authentication information, please refer to the respective model's official website
	// 
	// - OpenAI protocol: [GPT](https://help.openai.com/en/articles/4936850-where-do-i-find-my-openai-api-key), [DeepSeek](https://api-docs.deepseek.com/zh-cn/);
	// 
	// - Azure protocol: [Azure GPT](https://learn.microsoft.com/en-us/azure/ai-services/openai/chatgpt-quickstart?tabs=command-line%2Ctypescript%2Cpython-new&pivots=programming-language-studio#key-settings);
	// 
	// - Minimax:[Minimax](https://platform.minimaxi.com/document/Fast%20access?key=66701cf51d57f38758d581b2)
	APIKey *string `json:"APIKey,omitnil,omitempty" name:"APIKey"`

	// Model interface address
	// 
	// - OpenAI protocol
	// GPT:"https://api.openai.com/v1/"
	// Deepseek:"https://api.deepseek.com/v1"
	// 
	// - Azure protocol
	//  "https://{your-resource-name}.openai.azure.com?api-version={api-version}"
	// 
	// - Minimax protocol
	// "https://api.minimax.chat/v1"
	APIUrl *string `json:"APIUrl,omitnil,omitempty" name:"APIUrl"`

	// ## Identity
	// You are Kate from the appointment department at Retell Health calling Cindy over the phone to prepare for the annual checkup coming up. You are a pleasant and friendly receptionist caring deeply for the user. You don't provide medical advice but would use the medical knowledge to understand user responses.
	// 
	// ## Style Guardrails
	// Be Concise: Respond succinctly, addressing one topic at most.
	// Embrace Variety: Use diverse language and rephrasing to enhance clarity without repeating content.
	// Be Conversational: Use everyday language, making the chat feel like talking to a friend.
	// Be Proactive: Lead the conversation, often wrapping up with a question or next-step suggestion.
	// Avoid multiple questions in a single response.
	// Get clarity: If the user only partially answers a question, or if the answer is unclear, keep asking to get clarity.
	// Use a colloquial way of referring to the date (like Friday, January 14th, or Tuesday, January 12th, 2024 at 8am).
	// 
	// ## Response Guideline
	// Adapt and Guess: Try to understand transcripts that may contain transcription errors. Avoid mentioning "transcription error" in the response.
	// Stay in Character: Keep conversations within your role's scope, guiding them back creatively without repeating.
	// Ensure Fluid Dialogue: Respond in a role-appropriate, direct manner to maintain a smooth conversation flow.
	// 
	// ## Task
	// You will follow the steps below, do not skip steps, and only ask up to one question in response.
	// If at any time the user showed anger or wanted a human agent, call transfer_call to transfer to a human representative.
	// 1. Begin with a self-introduction and verify if callee is Cindy.
	//   - if callee is not Cindy, call end_call to hang up, say sorry for the confusion when hanging up.
	//   - if Cindy is not available, call end_call politely to hang up, say you will call back later when hanging up.
	// 2. Inform Cindy she has an annual body check coming up on April 4th, 2024 at 10am PDT. Check if Cindy is available.
	//   - If not, tell Cindy to reschedule online and jump to step 5.
	// 3. Ask Cindy if there's anything that the doctor should know before the annual checkup.
	//   - Ask followup questions as needed to assess the severity of the issue, and understand how it has progressed.
	// 4. Tell Cindy to not eat or drink that day before the checkup. Also tell Cindy to give you a callback if there's any changes in health condition.
	// 5. Ask Cindy if she has any questions, and if so, answer them until there are no questions.
	//   - If user asks something you do not know, let them know you don't have the answer. Ask them if they have any other questions.
	//   - If user do not have any questions, call function end_call to hang up.
	SystemPrompt *string `json:"SystemPrompt,omitnil,omitempty" name:"SystemPrompt"`

	// Model name, such as
	// 
	// - OpenAI protocol
	// "gpt-4o-mini","gpt-4o","deepseek-chat";
	// 
	// - Azure protocol
	// "gpt-4o-mini", "gpt-4o";
	// 
	// - Minimax protocol
	// "deepseek-chat".
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`

	// The following voice parameter values are available by default. If you wish to customize the voice type, please leave VoiceType blank and configure it in the CustomTTSConfig parameter.
	// 
	// Chinese:
	// ZhiMei: Zhimei, customer service female voice
	// ZhiXi: Zhixi, general female voice
	// ZhiQi: Zhiqi, customer service female voice
	// ZhiTian: Zhitian, female child voice
	// AiXiaoJing: Ai Xiaojing, dialogue female voice
	// 
	// English:
	// WeRose:English Female Voice
	// Monika:English Female Voice
	// 
	// Japanese:
	// Nanami
	// 
	// Korean:
	// SunHi
	// 
	// Indonesian (Indonesia):
	// Gadis
	// 
	// Malay (Malaysia):
	// Yasmin
	// 
	//  Tamil (Malaysia):
	// Kani
	// 
	// Thai (Thailand):
	// Achara
	// 
	// Vietnamese (Vietnam):
	// HoaiMy
	// 
	VoiceType *string `json:"VoiceType,omitnil,omitempty" name:"VoiceType"`

	// Caller number list
	Callers []*string `json:"Callers,omitnil,omitempty" name:"Callers"`

	// Used to set the AI Agent Welcome Message.
	WelcomeMessage *string `json:"WelcomeMessage,omitnil,omitempty" name:"WelcomeMessage"`

	// 0: Use welcomeMessage (if empty, the callee speaks first; if not empty, the bot speaks first)
	// 1:   Use AI to automatically generate welcomeMessage and speak first based on the prompt
	WelcomeType *int64 `json:"WelcomeType,omitnil,omitempty" name:"WelcomeType"`

	// 0: interruptible by default, 2: high priority non-interruptible.
	WelcomeMessagePriority *int64 `json:"WelcomeMessagePriority,omitnil,omitempty" name:"WelcomeMessagePriority"`

	// Maximum Waiting Duration (milliseconds), default is 60 seconds, if the user does not speak within this time, the call is automatically terminated
	MaxDuration *int64 `json:"MaxDuration,omitnil,omitempty" name:"MaxDuration"`

	// ASR Supported Languages, default is "zh" Chinese,
	// Fill in the array with up to 4 languages, the first is the primary language for recognition, followed by optional languages,
	// Note: When the primary language is a Chinese dialect, optional languages are invalid
	// Currently, the supported languages are as follows. The English name of the language is on the left side of the equals sign, and the value to be filled in the Language field is on the right side, following ISO639:
	// 1. Chinese = "zh" # Chinese
	// 2. Chinese_TW = "zh-TW" # Taiwan (China)
	// 3. Chinese_DIALECT = "zh-dialect" # Chinese Dialect
	// 4. English = "en" # English
	// 5. Vietnamese = "vi" # Vietnamese
	// 6. Japanese = "ja" # Japanese
	// 7. Korean = "ko" # Korean
	// 8. Indonesia = "id" # Indonesian
	// 9. Thai = "th" # Thai
	// 10. Portuguese = "pt" # Portuguese
	// 11. Turkish = "tr" # Turkish
	// 12. Arabic = "ar" # Arabic
	// 13. Spanish = "es" # Spanish
	// 14. Hindi = "hi" # Hindi
	// 15. French = "fr" # French
	// 16. Malay = "ms" # Malay
	// 17. Filipino = "fil" # Filipino
	// 18. German = "de" # German
	// 19. Italian = "it" # Italian
	// 20. Russian = "ru" # Russian
	Languages []*string `json:"Languages,omitnil,omitempty" name:"Languages"`

	// Interrupt ai speaking mode. default is 0. 0 indicates automatic interruption and 1 indicates no interruption.
	InterruptMode *int64 `json:"InterruptMode,omitnil,omitempty" name:"InterruptMode"`

	// Used when InterruptMode is 0, unit in milliseconds, default is 500ms. It means that the server-side detects ongoing vocal input for the InterruptSpeechDuration milliseconds and then interrupts.
	InterruptSpeechDuration *int64 `json:"InterruptSpeechDuration,omitnil,omitempty" name:"InterruptSpeechDuration"`

	// Whether the model supports (or enables) call_end function calling
	EndFunctionEnable *bool `json:"EndFunctionEnable,omitnil,omitempty" name:"EndFunctionEnable"`

	// Effective when EndFunctionEnable is true; the description of call_end function calling, default is "End the call when user has to leave (like says bye) or you are instructed to do so."
	EndFunctionDesc *string `json:"EndFunctionDesc,omitnil,omitempty" name:"EndFunctionDesc"`

	// Whether the model supports (or enables) transfer_to_human function calling.
	TransferFunctionEnable *bool `json:"TransferFunctionEnable,omitnil,omitempty" name:"TransferFunctionEnable"`

	// Takes effect when transferfunctionenable is true: transfer to human configuration.
	TransferItems []*AITransferItem `json:"TransferItems,omitnil,omitempty" name:"TransferItems"`

	// The duration after which the user hasn't spoken to trigger a notification, minimum 10 seconds, default 10 seconds
	NotifyDuration *int64 `json:"NotifyDuration,omitnil,omitempty" name:"NotifyDuration"`

	// The AI prompt when NotifyDuration has passed without the user speaking, default is "Sorry, I didn't hear you clearly. Can you repeat that?"
	NotifyMessage *string `json:"NotifyMessage,omitnil,omitempty" name:"NotifyMessage"`

	// Maximum number of times to trigger ai prompt sound, unlimited by default.
	NotifyMaxCount *uint64 `json:"NotifyMaxCount,omitnil,omitempty" name:"NotifyMaxCount"`

	// <p>And VoiceType field needs to select one, here is to use your own custom TTS, VoiceType is some built-in sound qualities</p>
	// <ul>
	// <li>Tencent TTS<br>
	// For configuration, please refer to <a href="https://intl.cloud.tencent.com/document/product/1073/92668?from_cn_redirect=1#55924b56-1a73-4663-a7a1-a8dd82d6e823" target="_blank">Tencent Cloud TTS documentation link</a></li>
	// </ul>
	// <div><div class="v-md-pre-wrapper copy-code-mode v-md-pre-wrapper- extra-class"><pre class="v-md-prism-"><code>{
	//     &quot;TTSType&quot;: &quot;tencent&quot;, // String TTS type, currently supports &quot;tencent&quot; and "minixmax", other vendors support in progress
	//     &quot;AppId&quot;: &quot;Your application ID&quot;, // String required
	//     &quot;SecretId&quot;: &quot;Your Secret ID&quot;, // String Required
	//     &quot;SecretKey&quot;:  &quot;Your Secret Key&quot;, // String Required
	//     &quot;VoiceType&quot;: 101001, // Integer Required, Sound quality ID, includes standard and premium sound quality. Premium sound quality is more realistic and differently priced than standard sound quality. See TTS billing overview for details. For the full list of sound quality IDs, see the TTS sound quality list.
	//     "Speed": 1.25, // Integer Optional, speech speed, range: [-2,6], corresponding to different speeds: -2: represents 0.6x -1: represents 0.8x 0: represents 1.0x (default) 1: represents 1.2x 2: represents 1.5x 6: represents 2.5x For more precise speed control, you can retain two decimal places, such as 0.5/1.25/2.81, etc. For parameter value to actual speed conversion, refer to Speed Conversion
	//     &quot;Volume&quot;: 5, // Integer Optional, Volume level, range: [0,10], corresponding to 11 levels of volume, default is 0, which represents normal volume.
	//     &quot;PrimaryLanguage&quot;: 1, // Integer Optional, Primary language 1- Chinese (default) 2- English 3- Japanese
	//     &quot;FastVoiceType&quot;: &quot;xxxx&quot;   // Optional parameter, Fast VRS parameter
	//   }
	// </code></pre>
	// 
	//   </div></div><ul>
	// 
	// </div></div><ul>
	// <li>Azure TTS<br>
	// For configuration, refer to the<a href="https://docs.azure.cn/zh-cn/ai-services/speech-service/speech-synthesis-markup-voice" target="_blank">Azure TTS documentation</a></li>
	// </ul>
	// <div><div class="v-md-pre-wrapper copy-code-mode v-md-pre-wrapper- extra-class"><pre class="v-md-prism-"><code>{
	//     &quot;TTSType&quot;: &quot;azure&quot;, // Required: String TTS type
	//     &quot;SubscriptionKey&quot;: &quot;xxxxxxxx&quot;, // Required: String subscription key
	//     &quot;Region&quot;: &quot;chinanorth3&quot;,  // Required: String subscription region
	//     &quot;VoiceName&quot;: &quot;zh-CN-XiaoxiaoNeural&quot;, // Required: String Timbre Name required
	//     &quot;Language&quot;: &quot;zh-CN&quot;, // Required: String Language for synthesis
	//     &quot;Rate&quot;: 1 // Optional: float Playback Speed 0.5-2 default is 1
	// }
	// </code></pre>
	// 
	// </div></div><ul>
	// <li>Custom</li>
	// </ul>
	// <p>TTS<br>
	// Please refer to the specific protocol standards in the <a href="https://doc.weixin.qq.com/doc/w3_ANQAiAbdAFwHILbJBmtSqSbV1WZ3L?scode=AJEAIQdfAAo5a1xajYANQAiAbdAFw" target="_blank">Tencent documentation</a></p>
	// <div><div class="v-md-pre-wrapper copy-code-mode v-md-pre-wrapper- extra-class"><pre class="v-md-prism-"><code>{
	//   &quot;TTSType&quot;: &quot;custom&quot;, // Required String
	//   &quot;APIKey&quot;: &quot;ApiKey&quot;, // Required String for Authentication
	//   &quot;APIUrl&quot;: &quot;http://0.0.0.0:8080/stream-audio&quot; // Required String, TTS API URL
	//   &quot;AudioFormat&quot;: &quot;wav&quot;, // String, optional, expected audio format, such as mp3, ogg_opus, pcm, wav, default is wav, currently only pcm and wav are supported,
	//   &quot;SampleRate&quot;: 16000,  // Integer, optional, audio sample rate, default is 16000 (16k), recommended value is 16000
	//   &quot;AudioChannel&quot;: 1,    // Integer, optional, number of audio channels, values: 1 or 2, default is 1
	// }
	// </code></pre>
	// 
	// </div></div>
	CustomTTSConfig *string `json:"CustomTTSConfig,omitnil,omitempty" name:"CustomTTSConfig"`

	// Prompt word variable.
	//
	// Deprecated: PromptVariables is deprecated.
	PromptVariables []*Variable `json:"PromptVariables,omitnil,omitempty" name:"PromptVariables"`

	// Automatic speech recognition vad time ranges from 240 to 2000, with a default of 1000, measured in milliseconds. smaller values will make automatic speech recognition segment faster.
	VadSilenceTime *int64 `json:"VadSilenceTime,omitnil,omitempty" name:"VadSilenceTime"`

	// Call content extraction configuration.
	ExtractConfig []*AICallExtractConfigElement `json:"ExtractConfig,omitnil,omitempty" name:"ExtractConfig"`

	// Model temperature control.
	Temperature *float64 `json:"Temperature,omitnil,omitempty" name:"Temperature"`

	// Common variable: <p>prompt content variable</p> <p>welcome message variable</p> <p>welcome message delay playback (in seconds): welcome-message-delay</p> <p>dify variable</p>.  
	// 
	// dify-inputs-xxx specifies the inputs variable for dify.
	// 2. the dify-inputs-user specifies the user value for dify.
	// 3. dify-inputs-conversation_id is the conversation_id value of dify.
	Variables []*Variable `json:"Variables,omitnil,omitempty" name:"Variables"`

	// Specifies the model topP.
	TopP *float64 `json:"TopP,omitnil,omitempty" name:"TopP"`

	// The vad far-field voice suppression capacity (does not impact asr recognition performance). value range: [0, 3]. default is 0. recommended setting: 2 for better far-field voice suppression.
	VadLevel *uint64 `json:"VadLevel,omitnil,omitempty" name:"VadLevel"`

	// Transition.
	ToneWord *ToneWordInfo `json:"ToneWord,omitnil,omitempty" name:"ToneWord"`

	// Compliance prompt sound. 
	// This parameter specifies whether to play morse code during call initiation (default: true), indicating the conversation content is AI-generated.
	// This parameter signifies disabled when set to false. the parameter indicates you understand and agree to the following protocol:.
	// Our side fully acknowledges and understands that according to the laws and regulations including the "cybersecurity law" (https://www.gov.cn/xinwen/2016-11/07/content_5129723.htm), "provision on administration of deep synthesis of internet-based information service" (https://www.gov.cn/zhengce/zhengceku/2022-12/12/content_5731431.htm), "interim measures for the management of generative artificial intelligence services" (https://www.gov.cn/zhengce/zhengceku/202307/content_6891752.htm), and "measures for the identification of artificial intelligence-generated synthetic content" (https://www.gov.cn/zhengce/zhengceku/202503/content_7014286.htm), explicit and implicit identification shall be added to ai-generated synthetic content. based on business needs, we request tencent cloud not to add explicit identification to generated synthetic content. we commit to lawful and compliant use of such content to avoid confusion or misunderstanding. if the ai-generated synthetic content is used to provide services to the public or spread over networks, we will proactively add explicit identification compliant with legal provisions and national standard requirements and bear the legal obligations for ai-generated synthetic content identification. if we fail to properly fulfill the identification obligations for ai-generated content, resulting in adverse consequences or penalties from the competent department, we will fully assume all related responsibilities.
	EnableComplianceAudio *bool `json:"EnableComplianceAudio,omitnil,omitempty" name:"EnableComplianceAudio"`
}

type CreateAICallRequest struct {
	*tchttp.BaseRequest
	
	// Application ID (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Called number.
	Callee *string `json:"Callee,omitnil,omitempty" name:"Callee"`

	// Model API protocol type. currently compatible with four protocol types:.
	// 
	// -OpenAI protocol (including GPT, hunyuan, DeepSeek, etc.): "OpenAI".
	// -Azure protocol: "azure".
	// -Specifies the "Minimax" protocol.
	// -Dify protocol: "dify".
	LLMType *string `json:"LLMType,omitnil,omitempty" name:"LLMType"`

	// Model API key, for authentication information, please refer to the respective model's official website
	// 
	// - OpenAI protocol: [GPT](https://help.openai.com/en/articles/4936850-where-do-i-find-my-openai-api-key), [DeepSeek](https://api-docs.deepseek.com/zh-cn/);
	// 
	// - Azure protocol: [Azure GPT](https://learn.microsoft.com/en-us/azure/ai-services/openai/chatgpt-quickstart?tabs=command-line%2Ctypescript%2Cpython-new&pivots=programming-language-studio#key-settings);
	// 
	// - Minimax:[Minimax](https://platform.minimaxi.com/document/Fast%20access?key=66701cf51d57f38758d581b2)
	APIKey *string `json:"APIKey,omitnil,omitempty" name:"APIKey"`

	// Model interface address
	// 
	// - OpenAI protocol
	// GPT:"https://api.openai.com/v1/"
	// Deepseek:"https://api.deepseek.com/v1"
	// 
	// - Azure protocol
	//  "https://{your-resource-name}.openai.azure.com?api-version={api-version}"
	// 
	// - Minimax protocol
	// "https://api.minimax.chat/v1"
	APIUrl *string `json:"APIUrl,omitnil,omitempty" name:"APIUrl"`

	// ## Identity
	// You are Kate from the appointment department at Retell Health calling Cindy over the phone to prepare for the annual checkup coming up. You are a pleasant and friendly receptionist caring deeply for the user. You don't provide medical advice but would use the medical knowledge to understand user responses.
	// 
	// ## Style Guardrails
	// Be Concise: Respond succinctly, addressing one topic at most.
	// Embrace Variety: Use diverse language and rephrasing to enhance clarity without repeating content.
	// Be Conversational: Use everyday language, making the chat feel like talking to a friend.
	// Be Proactive: Lead the conversation, often wrapping up with a question or next-step suggestion.
	// Avoid multiple questions in a single response.
	// Get clarity: If the user only partially answers a question, or if the answer is unclear, keep asking to get clarity.
	// Use a colloquial way of referring to the date (like Friday, January 14th, or Tuesday, January 12th, 2024 at 8am).
	// 
	// ## Response Guideline
	// Adapt and Guess: Try to understand transcripts that may contain transcription errors. Avoid mentioning "transcription error" in the response.
	// Stay in Character: Keep conversations within your role's scope, guiding them back creatively without repeating.
	// Ensure Fluid Dialogue: Respond in a role-appropriate, direct manner to maintain a smooth conversation flow.
	// 
	// ## Task
	// You will follow the steps below, do not skip steps, and only ask up to one question in response.
	// If at any time the user showed anger or wanted a human agent, call transfer_call to transfer to a human representative.
	// 1. Begin with a self-introduction and verify if callee is Cindy.
	//   - if callee is not Cindy, call end_call to hang up, say sorry for the confusion when hanging up.
	//   - if Cindy is not available, call end_call politely to hang up, say you will call back later when hanging up.
	// 2. Inform Cindy she has an annual body check coming up on April 4th, 2024 at 10am PDT. Check if Cindy is available.
	//   - If not, tell Cindy to reschedule online and jump to step 5.
	// 3. Ask Cindy if there's anything that the doctor should know before the annual checkup.
	//   - Ask followup questions as needed to assess the severity of the issue, and understand how it has progressed.
	// 4. Tell Cindy to not eat or drink that day before the checkup. Also tell Cindy to give you a callback if there's any changes in health condition.
	// 5. Ask Cindy if she has any questions, and if so, answer them until there are no questions.
	//   - If user asks something you do not know, let them know you don't have the answer. Ask them if they have any other questions.
	//   - If user do not have any questions, call function end_call to hang up.
	SystemPrompt *string `json:"SystemPrompt,omitnil,omitempty" name:"SystemPrompt"`

	// Model name, such as
	// 
	// - OpenAI protocol
	// "gpt-4o-mini","gpt-4o","deepseek-chat";
	// 
	// - Azure protocol
	// "gpt-4o-mini", "gpt-4o";
	// 
	// - Minimax protocol
	// "deepseek-chat".
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`

	// The following voice parameter values are available by default. If you wish to customize the voice type, please leave VoiceType blank and configure it in the CustomTTSConfig parameter.
	// 
	// Chinese:
	// ZhiMei: Zhimei, customer service female voice
	// ZhiXi: Zhixi, general female voice
	// ZhiQi: Zhiqi, customer service female voice
	// ZhiTian: Zhitian, female child voice
	// AiXiaoJing: Ai Xiaojing, dialogue female voice
	// 
	// English:
	// WeRose:English Female Voice
	// Monika:English Female Voice
	// 
	// Japanese:
	// Nanami
	// 
	// Korean:
	// SunHi
	// 
	// Indonesian (Indonesia):
	// Gadis
	// 
	// Malay (Malaysia):
	// Yasmin
	// 
	//  Tamil (Malaysia):
	// Kani
	// 
	// Thai (Thailand):
	// Achara
	// 
	// Vietnamese (Vietnam):
	// HoaiMy
	// 
	VoiceType *string `json:"VoiceType,omitnil,omitempty" name:"VoiceType"`

	// Caller number list
	Callers []*string `json:"Callers,omitnil,omitempty" name:"Callers"`

	// Used to set the AI Agent Welcome Message.
	WelcomeMessage *string `json:"WelcomeMessage,omitnil,omitempty" name:"WelcomeMessage"`

	// 0: Use welcomeMessage (if empty, the callee speaks first; if not empty, the bot speaks first)
	// 1:   Use AI to automatically generate welcomeMessage and speak first based on the prompt
	WelcomeType *int64 `json:"WelcomeType,omitnil,omitempty" name:"WelcomeType"`

	// 0: interruptible by default, 2: high priority non-interruptible.
	WelcomeMessagePriority *int64 `json:"WelcomeMessagePriority,omitnil,omitempty" name:"WelcomeMessagePriority"`

	// Maximum Waiting Duration (milliseconds), default is 60 seconds, if the user does not speak within this time, the call is automatically terminated
	MaxDuration *int64 `json:"MaxDuration,omitnil,omitempty" name:"MaxDuration"`

	// ASR Supported Languages, default is "zh" Chinese,
	// Fill in the array with up to 4 languages, the first is the primary language for recognition, followed by optional languages,
	// Note: When the primary language is a Chinese dialect, optional languages are invalid
	// Currently, the supported languages are as follows. The English name of the language is on the left side of the equals sign, and the value to be filled in the Language field is on the right side, following ISO639:
	// 1. Chinese = "zh" # Chinese
	// 2. Chinese_TW = "zh-TW" # Taiwan (China)
	// 3. Chinese_DIALECT = "zh-dialect" # Chinese Dialect
	// 4. English = "en" # English
	// 5. Vietnamese = "vi" # Vietnamese
	// 6. Japanese = "ja" # Japanese
	// 7. Korean = "ko" # Korean
	// 8. Indonesia = "id" # Indonesian
	// 9. Thai = "th" # Thai
	// 10. Portuguese = "pt" # Portuguese
	// 11. Turkish = "tr" # Turkish
	// 12. Arabic = "ar" # Arabic
	// 13. Spanish = "es" # Spanish
	// 14. Hindi = "hi" # Hindi
	// 15. French = "fr" # French
	// 16. Malay = "ms" # Malay
	// 17. Filipino = "fil" # Filipino
	// 18. German = "de" # German
	// 19. Italian = "it" # Italian
	// 20. Russian = "ru" # Russian
	Languages []*string `json:"Languages,omitnil,omitempty" name:"Languages"`

	// Interrupt ai speaking mode. default is 0. 0 indicates automatic interruption and 1 indicates no interruption.
	InterruptMode *int64 `json:"InterruptMode,omitnil,omitempty" name:"InterruptMode"`

	// Used when InterruptMode is 0, unit in milliseconds, default is 500ms. It means that the server-side detects ongoing vocal input for the InterruptSpeechDuration milliseconds and then interrupts.
	InterruptSpeechDuration *int64 `json:"InterruptSpeechDuration,omitnil,omitempty" name:"InterruptSpeechDuration"`

	// Whether the model supports (or enables) call_end function calling
	EndFunctionEnable *bool `json:"EndFunctionEnable,omitnil,omitempty" name:"EndFunctionEnable"`

	// Effective when EndFunctionEnable is true; the description of call_end function calling, default is "End the call when user has to leave (like says bye) or you are instructed to do so."
	EndFunctionDesc *string `json:"EndFunctionDesc,omitnil,omitempty" name:"EndFunctionDesc"`

	// Whether the model supports (or enables) transfer_to_human function calling.
	TransferFunctionEnable *bool `json:"TransferFunctionEnable,omitnil,omitempty" name:"TransferFunctionEnable"`

	// Takes effect when transferfunctionenable is true: transfer to human configuration.
	TransferItems []*AITransferItem `json:"TransferItems,omitnil,omitempty" name:"TransferItems"`

	// The duration after which the user hasn't spoken to trigger a notification, minimum 10 seconds, default 10 seconds
	NotifyDuration *int64 `json:"NotifyDuration,omitnil,omitempty" name:"NotifyDuration"`

	// The AI prompt when NotifyDuration has passed without the user speaking, default is "Sorry, I didn't hear you clearly. Can you repeat that?"
	NotifyMessage *string `json:"NotifyMessage,omitnil,omitempty" name:"NotifyMessage"`

	// Maximum number of times to trigger ai prompt sound, unlimited by default.
	NotifyMaxCount *uint64 `json:"NotifyMaxCount,omitnil,omitempty" name:"NotifyMaxCount"`

	// <p>And VoiceType field needs to select one, here is to use your own custom TTS, VoiceType is some built-in sound qualities</p>
	// <ul>
	// <li>Tencent TTS<br>
	// For configuration, please refer to <a href="https://intl.cloud.tencent.com/document/product/1073/92668?from_cn_redirect=1#55924b56-1a73-4663-a7a1-a8dd82d6e823" target="_blank">Tencent Cloud TTS documentation link</a></li>
	// </ul>
	// <div><div class="v-md-pre-wrapper copy-code-mode v-md-pre-wrapper- extra-class"><pre class="v-md-prism-"><code>{
	//     &quot;TTSType&quot;: &quot;tencent&quot;, // String TTS type, currently supports &quot;tencent&quot; and "minixmax", other vendors support in progress
	//     &quot;AppId&quot;: &quot;Your application ID&quot;, // String required
	//     &quot;SecretId&quot;: &quot;Your Secret ID&quot;, // String Required
	//     &quot;SecretKey&quot;:  &quot;Your Secret Key&quot;, // String Required
	//     &quot;VoiceType&quot;: 101001, // Integer Required, Sound quality ID, includes standard and premium sound quality. Premium sound quality is more realistic and differently priced than standard sound quality. See TTS billing overview for details. For the full list of sound quality IDs, see the TTS sound quality list.
	//     "Speed": 1.25, // Integer Optional, speech speed, range: [-2,6], corresponding to different speeds: -2: represents 0.6x -1: represents 0.8x 0: represents 1.0x (default) 1: represents 1.2x 2: represents 1.5x 6: represents 2.5x For more precise speed control, you can retain two decimal places, such as 0.5/1.25/2.81, etc. For parameter value to actual speed conversion, refer to Speed Conversion
	//     &quot;Volume&quot;: 5, // Integer Optional, Volume level, range: [0,10], corresponding to 11 levels of volume, default is 0, which represents normal volume.
	//     &quot;PrimaryLanguage&quot;: 1, // Integer Optional, Primary language 1- Chinese (default) 2- English 3- Japanese
	//     &quot;FastVoiceType&quot;: &quot;xxxx&quot;   // Optional parameter, Fast VRS parameter
	//   }
	// </code></pre>
	// 
	//   </div></div><ul>
	// 
	// </div></div><ul>
	// <li>Azure TTS<br>
	// For configuration, refer to the<a href="https://docs.azure.cn/zh-cn/ai-services/speech-service/speech-synthesis-markup-voice" target="_blank">Azure TTS documentation</a></li>
	// </ul>
	// <div><div class="v-md-pre-wrapper copy-code-mode v-md-pre-wrapper- extra-class"><pre class="v-md-prism-"><code>{
	//     &quot;TTSType&quot;: &quot;azure&quot;, // Required: String TTS type
	//     &quot;SubscriptionKey&quot;: &quot;xxxxxxxx&quot;, // Required: String subscription key
	//     &quot;Region&quot;: &quot;chinanorth3&quot;,  // Required: String subscription region
	//     &quot;VoiceName&quot;: &quot;zh-CN-XiaoxiaoNeural&quot;, // Required: String Timbre Name required
	//     &quot;Language&quot;: &quot;zh-CN&quot;, // Required: String Language for synthesis
	//     &quot;Rate&quot;: 1 // Optional: float Playback Speed 0.5-2 default is 1
	// }
	// </code></pre>
	// 
	// </div></div><ul>
	// <li>Custom</li>
	// </ul>
	// <p>TTS<br>
	// Please refer to the specific protocol standards in the <a href="https://doc.weixin.qq.com/doc/w3_ANQAiAbdAFwHILbJBmtSqSbV1WZ3L?scode=AJEAIQdfAAo5a1xajYANQAiAbdAFw" target="_blank">Tencent documentation</a></p>
	// <div><div class="v-md-pre-wrapper copy-code-mode v-md-pre-wrapper- extra-class"><pre class="v-md-prism-"><code>{
	//   &quot;TTSType&quot;: &quot;custom&quot;, // Required String
	//   &quot;APIKey&quot;: &quot;ApiKey&quot;, // Required String for Authentication
	//   &quot;APIUrl&quot;: &quot;http://0.0.0.0:8080/stream-audio&quot; // Required String, TTS API URL
	//   &quot;AudioFormat&quot;: &quot;wav&quot;, // String, optional, expected audio format, such as mp3, ogg_opus, pcm, wav, default is wav, currently only pcm and wav are supported,
	//   &quot;SampleRate&quot;: 16000,  // Integer, optional, audio sample rate, default is 16000 (16k), recommended value is 16000
	//   &quot;AudioChannel&quot;: 1,    // Integer, optional, number of audio channels, values: 1 or 2, default is 1
	// }
	// </code></pre>
	// 
	// </div></div>
	CustomTTSConfig *string `json:"CustomTTSConfig,omitnil,omitempty" name:"CustomTTSConfig"`

	// Prompt word variable.
	PromptVariables []*Variable `json:"PromptVariables,omitnil,omitempty" name:"PromptVariables"`

	// Automatic speech recognition vad time ranges from 240 to 2000, with a default of 1000, measured in milliseconds. smaller values will make automatic speech recognition segment faster.
	VadSilenceTime *int64 `json:"VadSilenceTime,omitnil,omitempty" name:"VadSilenceTime"`

	// Call content extraction configuration.
	ExtractConfig []*AICallExtractConfigElement `json:"ExtractConfig,omitnil,omitempty" name:"ExtractConfig"`

	// Model temperature control.
	Temperature *float64 `json:"Temperature,omitnil,omitempty" name:"Temperature"`

	// Common variable: <p>prompt content variable</p> <p>welcome message variable</p> <p>welcome message delay playback (in seconds): welcome-message-delay</p> <p>dify variable</p>.  
	// 
	// dify-inputs-xxx specifies the inputs variable for dify.
	// 2. the dify-inputs-user specifies the user value for dify.
	// 3. dify-inputs-conversation_id is the conversation_id value of dify.
	Variables []*Variable `json:"Variables,omitnil,omitempty" name:"Variables"`

	// Specifies the model topP.
	TopP *float64 `json:"TopP,omitnil,omitempty" name:"TopP"`

	// The vad far-field voice suppression capacity (does not impact asr recognition performance). value range: [0, 3]. default is 0. recommended setting: 2 for better far-field voice suppression.
	VadLevel *uint64 `json:"VadLevel,omitnil,omitempty" name:"VadLevel"`

	// Transition.
	ToneWord *ToneWordInfo `json:"ToneWord,omitnil,omitempty" name:"ToneWord"`

	// Compliance prompt sound. 
	// This parameter specifies whether to play morse code during call initiation (default: true), indicating the conversation content is AI-generated.
	// This parameter signifies disabled when set to false. the parameter indicates you understand and agree to the following protocol:.
	// Our side fully acknowledges and understands that according to the laws and regulations including the "cybersecurity law" (https://www.gov.cn/xinwen/2016-11/07/content_5129723.htm), "provision on administration of deep synthesis of internet-based information service" (https://www.gov.cn/zhengce/zhengceku/2022-12/12/content_5731431.htm), "interim measures for the management of generative artificial intelligence services" (https://www.gov.cn/zhengce/zhengceku/202307/content_6891752.htm), and "measures for the identification of artificial intelligence-generated synthetic content" (https://www.gov.cn/zhengce/zhengceku/202503/content_7014286.htm), explicit and implicit identification shall be added to ai-generated synthetic content. based on business needs, we request tencent cloud not to add explicit identification to generated synthetic content. we commit to lawful and compliant use of such content to avoid confusion or misunderstanding. if the ai-generated synthetic content is used to provide services to the public or spread over networks, we will proactively add explicit identification compliant with legal provisions and national standard requirements and bear the legal obligations for ai-generated synthetic content identification. if we fail to properly fulfill the identification obligations for ai-generated content, resulting in adverse consequences or penalties from the competent department, we will fully assume all related responsibilities.
	EnableComplianceAudio *bool `json:"EnableComplianceAudio,omitnil,omitempty" name:"EnableComplianceAudio"`
}

func (r *CreateAICallRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAICallRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "Callee")
	delete(f, "LLMType")
	delete(f, "APIKey")
	delete(f, "APIUrl")
	delete(f, "SystemPrompt")
	delete(f, "Model")
	delete(f, "VoiceType")
	delete(f, "Callers")
	delete(f, "WelcomeMessage")
	delete(f, "WelcomeType")
	delete(f, "WelcomeMessagePriority")
	delete(f, "MaxDuration")
	delete(f, "Languages")
	delete(f, "InterruptMode")
	delete(f, "InterruptSpeechDuration")
	delete(f, "EndFunctionEnable")
	delete(f, "EndFunctionDesc")
	delete(f, "TransferFunctionEnable")
	delete(f, "TransferItems")
	delete(f, "NotifyDuration")
	delete(f, "NotifyMessage")
	delete(f, "NotifyMaxCount")
	delete(f, "CustomTTSConfig")
	delete(f, "PromptVariables")
	delete(f, "VadSilenceTime")
	delete(f, "ExtractConfig")
	delete(f, "Temperature")
	delete(f, "Variables")
	delete(f, "TopP")
	delete(f, "VadLevel")
	delete(f, "ToneWord")
	delete(f, "EnableComplianceAudio")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAICallRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAICallResponseParams struct {
	// Newly created session ID.
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAICallResponse struct {
	*tchttp.BaseResponse
	Response *CreateAICallResponseParams `json:"Response"`
}

func (r *CreateAICallResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAICallResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAdminURLRequestParams struct {
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Admin account.
	SeatUserId *string `json:"SeatUserId,omitnil,omitempty" name:"SeatUserId"`
}

type CreateAdminURLRequest struct {
	*tchttp.BaseRequest
	
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Admin account.
	SeatUserId *string `json:"SeatUserId,omitnil,omitempty" name:"SeatUserId"`
}

func (r *CreateAdminURLRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAdminURLRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "SeatUserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAdminURLRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAdminURLResponseParams struct {
	// Log-In link.
	URL *string `json:"URL,omitnil,omitempty" name:"URL"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAdminURLResponse struct {
	*tchttp.BaseResponse
	Response *CreateAdminURLResponseParams `json:"Response"`
}

func (r *CreateAdminURLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAdminURLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAgentCruiseDialingCampaignRequestParams struct {
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Task name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Agent account.
	Agent *string `json:"Agent,omitnil,omitempty" name:"Agent"`

	// Single-Round concurrent call volume 1-20.
	ConcurrencyNumber *int64 `json:"ConcurrencyNumber,omitnil,omitempty" name:"ConcurrencyNumber"`

	// Task start time. unix timestamp. the task will automatically start after this time.
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Task termination time. unix timestamp. the task will automatically terminate after this time.
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Called list supporting e.164 or number formats without country code.
	Callees []*string `json:"Callees,omitnil,omitempty" name:"Callees"`

	// Calling list using the number formats displayed on the management side.
	Callers []*string `json:"Callers,omitnil,omitempty" name:"Callers"`

	// Being called sequence: 0 for random 1 for in order.
	CallOrder *int64 `json:"CallOrder,omitnil,omitempty" name:"CallOrder"`

	// Caller custom data, maximum length 1024.
	UUI *string `json:"UUI,omitnil,omitempty" name:"UUI"`
}

type CreateAgentCruiseDialingCampaignRequest struct {
	*tchttp.BaseRequest
	
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Task name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Agent account.
	Agent *string `json:"Agent,omitnil,omitempty" name:"Agent"`

	// Single-Round concurrent call volume 1-20.
	ConcurrencyNumber *int64 `json:"ConcurrencyNumber,omitnil,omitempty" name:"ConcurrencyNumber"`

	// Task start time. unix timestamp. the task will automatically start after this time.
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Task termination time. unix timestamp. the task will automatically terminate after this time.
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Called list supporting e.164 or number formats without country code.
	Callees []*string `json:"Callees,omitnil,omitempty" name:"Callees"`

	// Calling list using the number formats displayed on the management side.
	Callers []*string `json:"Callers,omitnil,omitempty" name:"Callers"`

	// Being called sequence: 0 for random 1 for in order.
	CallOrder *int64 `json:"CallOrder,omitnil,omitempty" name:"CallOrder"`

	// Caller custom data, maximum length 1024.
	UUI *string `json:"UUI,omitnil,omitempty" name:"UUI"`
}

func (r *CreateAgentCruiseDialingCampaignRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAgentCruiseDialingCampaignRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "Name")
	delete(f, "Agent")
	delete(f, "ConcurrencyNumber")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Callees")
	delete(f, "Callers")
	delete(f, "CallOrder")
	delete(f, "UUI")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAgentCruiseDialingCampaignRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAgentCruiseDialingCampaignResponseParams struct {
	// Generated task id.
	CampaignId *int64 `json:"CampaignId,omitnil,omitempty" name:"CampaignId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAgentCruiseDialingCampaignResponse struct {
	*tchttp.BaseResponse
	Response *CreateAgentCruiseDialingCampaignResponseParams `json:"Response"`
}

func (r *CreateAgentCruiseDialingCampaignResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAgentCruiseDialingCampaignResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAutoCalloutTaskRequestParams struct {
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Task starting timestamp. unix second-level timestamp.
	NotBefore *int64 `json:"NotBefore,omitnil,omitempty" name:"NotBefore"`

	// List of called numbers.
	Callees []*string `json:"Callees,omitnil,omitempty" name:"Callees"`

	// List of calling numbers.
	Callers []*string `json:"Callers,omitnil,omitempty" name:"Callers"`

	// IVR Id used for calling. if not filled, AIAgentId needs to be filled.
	IvrId *uint64 `json:"IvrId,omitnil,omitempty" name:"IvrId"`

	// Task name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <Task description>.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Task stop timestamp. unix second-level timestamp.
	NotAfter *int64 `json:"NotAfter,omitnil,omitempty" name:"NotAfter"`

	// Maximum attempts, 1-3 times.
	Tries *uint64 `json:"Tries,omitnil,omitempty" name:"Tries"`

	// Custom variables (supported only in advanced versions).
	Variables []*Variable `json:"Variables,omitnil,omitempty" name:"Variables"`

	// UUI
	UUI *string `json:"UUI,omitnil,omitempty" name:"UUI"`

	// Property of the called.
	CalleeAttributes []*CalleeAttribute `json:"CalleeAttributes,omitnil,omitempty" name:"CalleeAttributes"`

	// IANA time zone name. see https://datatracker.ietf.org/doc/html/draft-ietf-netmod-iana-timezones.
	TimeZone *string `json:"TimeZone,omitnil,omitempty" name:"TimeZone"`

	// Available time period.
	AvailableTime []*TimeRange `json:"AvailableTime,omitnil,omitempty" name:"AvailableTime"`

	// Intelligent agent ID. if not filled, IvrId needs to be filled.
	AIAgentId *int64 `json:"AIAgentId,omitnil,omitempty" name:"AIAgentId"`
}

type CreateAutoCalloutTaskRequest struct {
	*tchttp.BaseRequest
	
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Task starting timestamp. unix second-level timestamp.
	NotBefore *int64 `json:"NotBefore,omitnil,omitempty" name:"NotBefore"`

	// List of called numbers.
	Callees []*string `json:"Callees,omitnil,omitempty" name:"Callees"`

	// List of calling numbers.
	Callers []*string `json:"Callers,omitnil,omitempty" name:"Callers"`

	// IVR Id used for calling. if not filled, AIAgentId needs to be filled.
	IvrId *uint64 `json:"IvrId,omitnil,omitempty" name:"IvrId"`

	// Task name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <Task description>.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Task stop timestamp. unix second-level timestamp.
	NotAfter *int64 `json:"NotAfter,omitnil,omitempty" name:"NotAfter"`

	// Maximum attempts, 1-3 times.
	Tries *uint64 `json:"Tries,omitnil,omitempty" name:"Tries"`

	// Custom variables (supported only in advanced versions).
	Variables []*Variable `json:"Variables,omitnil,omitempty" name:"Variables"`

	// UUI
	UUI *string `json:"UUI,omitnil,omitempty" name:"UUI"`

	// Property of the called.
	CalleeAttributes []*CalleeAttribute `json:"CalleeAttributes,omitnil,omitempty" name:"CalleeAttributes"`

	// IANA time zone name. see https://datatracker.ietf.org/doc/html/draft-ietf-netmod-iana-timezones.
	TimeZone *string `json:"TimeZone,omitnil,omitempty" name:"TimeZone"`

	// Available time period.
	AvailableTime []*TimeRange `json:"AvailableTime,omitnil,omitempty" name:"AvailableTime"`

	// Intelligent agent ID. if not filled, IvrId needs to be filled.
	AIAgentId *int64 `json:"AIAgentId,omitnil,omitempty" name:"AIAgentId"`
}

func (r *CreateAutoCalloutTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAutoCalloutTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "NotBefore")
	delete(f, "Callees")
	delete(f, "Callers")
	delete(f, "IvrId")
	delete(f, "Name")
	delete(f, "Description")
	delete(f, "NotAfter")
	delete(f, "Tries")
	delete(f, "Variables")
	delete(f, "UUI")
	delete(f, "CalleeAttributes")
	delete(f, "TimeZone")
	delete(f, "AvailableTime")
	delete(f, "AIAgentId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAutoCalloutTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAutoCalloutTaskResponseParams struct {
	// Task id.
	TaskId *uint64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAutoCalloutTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateAutoCalloutTaskResponseParams `json:"Response"`
}

func (r *CreateAutoCalloutTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAutoCalloutTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCCCSkillGroupRequestParams struct {
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Skill group name.
	SkillGroupName *string `json:"SkillGroupName,omitnil,omitempty" name:"SkillGroupName"`

	// Skill group type 0-cell phone, 1-online, 3-audio, 4-video.
	SkillGroupType *int64 `json:"SkillGroupType,omitnil,omitempty" name:"SkillGroupType"`

	// The maximum number of people received by the skill group (the maximum number of people that one agent in this skill group can receive) is set to 1 by default. if the skill group type is online, the maximum can be set to one or more.
	// 2. if the skill group type is phone, audio, or video, then the reception limit must be 1.
	MaxConcurrency *uint64 `json:"MaxConcurrency,omitnil,omitempty" name:"MaxConcurrency"`
}

type CreateCCCSkillGroupRequest struct {
	*tchttp.BaseRequest
	
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Skill group name.
	SkillGroupName *string `json:"SkillGroupName,omitnil,omitempty" name:"SkillGroupName"`

	// Skill group type 0-cell phone, 1-online, 3-audio, 4-video.
	SkillGroupType *int64 `json:"SkillGroupType,omitnil,omitempty" name:"SkillGroupType"`

	// The maximum number of people received by the skill group (the maximum number of people that one agent in this skill group can receive) is set to 1 by default. if the skill group type is online, the maximum can be set to one or more.
	// 2. if the skill group type is phone, audio, or video, then the reception limit must be 1.
	MaxConcurrency *uint64 `json:"MaxConcurrency,omitnil,omitempty" name:"MaxConcurrency"`
}

func (r *CreateCCCSkillGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCCCSkillGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "SkillGroupName")
	delete(f, "SkillGroupType")
	delete(f, "MaxConcurrency")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCCCSkillGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCCCSkillGroupResponseParams struct {
	// Skill group id.
	SkillGroupId *int64 `json:"SkillGroupId,omitnil,omitempty" name:"SkillGroupId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCCCSkillGroupResponse struct {
	*tchttp.BaseResponse
	Response *CreateCCCSkillGroupResponseParams `json:"Response"`
}

func (r *CreateCCCSkillGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCCCSkillGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCallOutSessionRequestParams struct {
	// Application id.
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Customer service user ID, generally the customer service email. ensure that the mobile number has been bound. https://intl.cloud.tencent.com/document/product/679/76067?from_cn_redirect=1#.E6.AD.A5.E9.AA.A42.EF.BC.9A.E5.AE.8C.E5.96.84.E8.B4.A6.E5.8F.B7.E4.BF.A1.E6.81.AF.
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// Called number must be preceded by 0086.
	Callee *string `json:"Callee,omitnil,omitempty" name:"Callee"`

	// Caller number (obsolete one and use callers) must be preceded by 0086.
	Caller *string `json:"Caller,omitnil,omitempty" name:"Caller"`

	// Designated caller number list. if the prior number fails, it will automatically switch to the next number that must be preceded by 0086.
	Callers []*string `json:"Callers,omitnil,omitempty" name:"Callers"`

	// Whether to force the use of mobile outbound call. currently only supports true. if set to true, ensure the allowlist is configured. https://intl.cloud.tencent.com/document/product/679/76744?from_cn_redirect=1#.E6.93.8D.E4.BD.9C.E6.AD.A5.E9.AA.A4.
	IsForceUseMobile *bool `json:"IsForceUseMobile,omitnil,omitempty" name:"IsForceUseMobile"`

	// Custom data, length limited to 1024 bytes.
	//
	// Deprecated: Uui is deprecated.
	Uui *string `json:"Uui,omitnil,omitempty" name:"Uui"`

	// Custom data, length limited to 1024 bytes.
	UUI *string `json:"UUI,omitnil,omitempty" name:"UUI"`
}

type CreateCallOutSessionRequest struct {
	*tchttp.BaseRequest
	
	// Application id.
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Customer service user ID, generally the customer service email. ensure that the mobile number has been bound. https://intl.cloud.tencent.com/document/product/679/76067?from_cn_redirect=1#.E6.AD.A5.E9.AA.A42.EF.BC.9A.E5.AE.8C.E5.96.84.E8.B4.A6.E5.8F.B7.E4.BF.A1.E6.81.AF.
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// Called number must be preceded by 0086.
	Callee *string `json:"Callee,omitnil,omitempty" name:"Callee"`

	// Caller number (obsolete one and use callers) must be preceded by 0086.
	Caller *string `json:"Caller,omitnil,omitempty" name:"Caller"`

	// Designated caller number list. if the prior number fails, it will automatically switch to the next number that must be preceded by 0086.
	Callers []*string `json:"Callers,omitnil,omitempty" name:"Callers"`

	// Whether to force the use of mobile outbound call. currently only supports true. if set to true, ensure the allowlist is configured. https://intl.cloud.tencent.com/document/product/679/76744?from_cn_redirect=1#.E6.93.8D.E4.BD.9C.E6.AD.A5.E9.AA.A4.
	IsForceUseMobile *bool `json:"IsForceUseMobile,omitnil,omitempty" name:"IsForceUseMobile"`

	// Custom data, length limited to 1024 bytes.
	Uui *string `json:"Uui,omitnil,omitempty" name:"Uui"`

	// Custom data, length limited to 1024 bytes.
	UUI *string `json:"UUI,omitnil,omitempty" name:"UUI"`
}

func (r *CreateCallOutSessionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCallOutSessionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "UserId")
	delete(f, "Callee")
	delete(f, "Caller")
	delete(f, "Callers")
	delete(f, "IsForceUseMobile")
	delete(f, "Uui")
	delete(f, "UUI")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCallOutSessionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCallOutSessionResponseParams struct {
	// Newly created session id.
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCallOutSessionResponse struct {
	*tchttp.BaseResponse
	Response *CreateCallOutSessionResponseParams `json:"Response"`
}

func (r *CreateCallOutSessionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCallOutSessionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateExtensionRequestParams struct {
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Extension.
	ExtensionId *string `json:"ExtensionId,omitnil,omitempty" name:"ExtensionId"`

	// Extension name.
	ExtensionName *string `json:"ExtensionName,omitnil,omitempty" name:"ExtensionName"`

	// Bound skill group list.
	SkillGroupIds []*uint64 `json:"SkillGroupIds,omitnil,omitempty" name:"SkillGroupIds"`

	// Bound agent email.
	Relation *string `json:"Relation,omitnil,omitempty" name:"Relation"`
}

type CreateExtensionRequest struct {
	*tchttp.BaseRequest
	
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Extension.
	ExtensionId *string `json:"ExtensionId,omitnil,omitempty" name:"ExtensionId"`

	// Extension name.
	ExtensionName *string `json:"ExtensionName,omitnil,omitempty" name:"ExtensionName"`

	// Bound skill group list.
	SkillGroupIds []*uint64 `json:"SkillGroupIds,omitnil,omitempty" name:"SkillGroupIds"`

	// Bound agent email.
	Relation *string `json:"Relation,omitnil,omitempty" name:"Relation"`
}

func (r *CreateExtensionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateExtensionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "ExtensionId")
	delete(f, "ExtensionName")
	delete(f, "SkillGroupIds")
	delete(f, "Relation")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateExtensionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateExtensionResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateExtensionResponse struct {
	*tchttp.BaseResponse
	Response *CreateExtensionResponseParams `json:"Response"`
}

func (r *CreateExtensionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateExtensionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateIVRSessionRequestParams struct {
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Called.
	Callee *string `json:"Callee,omitnil,omitempty" name:"Callee"`

	// Specified ivr id. currently, it supports inbound and automatic outbound types.
	IVRId *int64 `json:"IVRId,omitnil,omitempty" name:"IVRId"`

	// List of calling numbers.
	Callers []*string `json:"Callers,omitnil,omitempty" name:"Callers"`

	// Custom variable.
	Variables []*Variable `json:"Variables,omitnil,omitempty" name:"Variables"`

	// User data.
	UUI *string `json:"UUI,omitnil,omitempty" name:"UUI"`
}

type CreateIVRSessionRequest struct {
	*tchttp.BaseRequest
	
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Called.
	Callee *string `json:"Callee,omitnil,omitempty" name:"Callee"`

	// Specified ivr id. currently, it supports inbound and automatic outbound types.
	IVRId *int64 `json:"IVRId,omitnil,omitempty" name:"IVRId"`

	// List of calling numbers.
	Callers []*string `json:"Callers,omitnil,omitempty" name:"Callers"`

	// Custom variable.
	Variables []*Variable `json:"Variables,omitnil,omitempty" name:"Variables"`

	// User data.
	UUI *string `json:"UUI,omitnil,omitempty" name:"UUI"`
}

func (r *CreateIVRSessionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIVRSessionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "Callee")
	delete(f, "IVRId")
	delete(f, "Callers")
	delete(f, "Variables")
	delete(f, "UUI")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateIVRSessionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateIVRSessionResponseParams struct {
	// Newly created session id.
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateIVRSessionResponse struct {
	*tchttp.BaseResponse
	Response *CreateIVRSessionResponseParams `json:"Response"`
}

func (r *CreateIVRSessionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIVRSessionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOwnNumberApplyRequestParams struct {
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// SIP connection id.
	SipTrunkId *int64 `json:"SipTrunkId,omitnil,omitempty" name:"SipTrunkId"`

	// Circuit-Related parameters.
	DetailList []*OwnNumberApplyDetailItem `json:"DetailList,omitnil,omitempty" name:"DetailList"`

	// Prefix for sending numbers.
	Prefix *string `json:"Prefix,omitnil,omitempty" name:"Prefix"`
}

type CreateOwnNumberApplyRequest struct {
	*tchttp.BaseRequest
	
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// SIP connection id.
	SipTrunkId *int64 `json:"SipTrunkId,omitnil,omitempty" name:"SipTrunkId"`

	// Circuit-Related parameters.
	DetailList []*OwnNumberApplyDetailItem `json:"DetailList,omitnil,omitempty" name:"DetailList"`

	// Prefix for sending numbers.
	Prefix *string `json:"Prefix,omitnil,omitempty" name:"Prefix"`
}

func (r *CreateOwnNumberApplyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOwnNumberApplyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "SipTrunkId")
	delete(f, "DetailList")
	delete(f, "Prefix")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateOwnNumberApplyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOwnNumberApplyResponseParams struct {
	// Approval id.
	ApplyId *uint64 `json:"ApplyId,omitnil,omitempty" name:"ApplyId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateOwnNumberApplyResponse struct {
	*tchttp.BaseResponse
	Response *CreateOwnNumberApplyResponseParams `json:"Response"`
}

func (r *CreateOwnNumberApplyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOwnNumberApplyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePredictiveDialingCampaignRequestParams struct {
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// <Task name>.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Called list supporting e.164 or number formats without country code.
	Callees []*string `json:"Callees,omitnil,omitempty" name:"Callees"`

	// Calling list using the number formats displayed on the management side.
	Callers []*string `json:"Callers,omitnil,omitempty" name:"Callers"`

	// Being called sequence: 0 for random 1 for in order.
	CallOrder *int64 `json:"CallOrder,omitnil,omitempty" name:"CallOrder"`

	// ID of the used skill group of agents.
	SkillGroupId *int64 `json:"SkillGroupId,omitnil,omitempty" name:"SkillGroupId"`

	// Running priority of multiple tasks in the same application, from high to low 1 - 5.
	Priority *int64 `json:"Priority,omitnil,omitempty" name:"Priority"`

	// Expected call drop rate, percentage, 5 - 50.
	ExpectedAbandonRate *int64 `json:"ExpectedAbandonRate,omitnil,omitempty" name:"ExpectedAbandonRate"`

	// Call retry interval, in seconds, [60 - 86,400].
	RetryInterval *int64 `json:"RetryInterval,omitnil,omitempty" name:"RetryInterval"`

	// Task start time. unix timestamp. the task will automatically start after this time.
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Task termination time. unix timestamp. the task will automatically terminate after this time.
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Specified ivr id.
	IVRId *int64 `json:"IVRId,omitnil,omitempty" name:"IVRId"`

	// Number of call retries, 0 - 2.
	RetryTimes *int64 `json:"RetryTimes,omitnil,omitempty" name:"RetryTimes"`

	// Custom variable.
	Variables []*Variable `json:"Variables,omitnil,omitempty" name:"Variables"`

	// UUI
	UUI *string `json:"UUI,omitnil,omitempty" name:"UUI"`

	// Property of the called.
	CalleeAttributes []*CalleeAttribute `json:"CalleeAttributes,omitnil,omitempty" name:"CalleeAttributes"`

	// IANA time zone name. see https://datatracker.ietf.org/doc/html/draft-ietf-netmod-iana-timezones.
	TimeZone *string `json:"TimeZone,omitnil,omitempty" name:"TimeZone"`

	// Value range of available time.
	AvailableTime []*TimeRange `json:"AvailableTime,omitnil,omitempty" name:"AvailableTime"`
}

type CreatePredictiveDialingCampaignRequest struct {
	*tchttp.BaseRequest
	
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// <Task name>.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Called list supporting e.164 or number formats without country code.
	Callees []*string `json:"Callees,omitnil,omitempty" name:"Callees"`

	// Calling list using the number formats displayed on the management side.
	Callers []*string `json:"Callers,omitnil,omitempty" name:"Callers"`

	// Being called sequence: 0 for random 1 for in order.
	CallOrder *int64 `json:"CallOrder,omitnil,omitempty" name:"CallOrder"`

	// ID of the used skill group of agents.
	SkillGroupId *int64 `json:"SkillGroupId,omitnil,omitempty" name:"SkillGroupId"`

	// Running priority of multiple tasks in the same application, from high to low 1 - 5.
	Priority *int64 `json:"Priority,omitnil,omitempty" name:"Priority"`

	// Expected call drop rate, percentage, 5 - 50.
	ExpectedAbandonRate *int64 `json:"ExpectedAbandonRate,omitnil,omitempty" name:"ExpectedAbandonRate"`

	// Call retry interval, in seconds, [60 - 86,400].
	RetryInterval *int64 `json:"RetryInterval,omitnil,omitempty" name:"RetryInterval"`

	// Task start time. unix timestamp. the task will automatically start after this time.
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Task termination time. unix timestamp. the task will automatically terminate after this time.
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Specified ivr id.
	IVRId *int64 `json:"IVRId,omitnil,omitempty" name:"IVRId"`

	// Number of call retries, 0 - 2.
	RetryTimes *int64 `json:"RetryTimes,omitnil,omitempty" name:"RetryTimes"`

	// Custom variable.
	Variables []*Variable `json:"Variables,omitnil,omitempty" name:"Variables"`

	// UUI
	UUI *string `json:"UUI,omitnil,omitempty" name:"UUI"`

	// Property of the called.
	CalleeAttributes []*CalleeAttribute `json:"CalleeAttributes,omitnil,omitempty" name:"CalleeAttributes"`

	// IANA time zone name. see https://datatracker.ietf.org/doc/html/draft-ietf-netmod-iana-timezones.
	TimeZone *string `json:"TimeZone,omitnil,omitempty" name:"TimeZone"`

	// Value range of available time.
	AvailableTime []*TimeRange `json:"AvailableTime,omitnil,omitempty" name:"AvailableTime"`
}

func (r *CreatePredictiveDialingCampaignRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePredictiveDialingCampaignRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "Name")
	delete(f, "Callees")
	delete(f, "Callers")
	delete(f, "CallOrder")
	delete(f, "SkillGroupId")
	delete(f, "Priority")
	delete(f, "ExpectedAbandonRate")
	delete(f, "RetryInterval")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "IVRId")
	delete(f, "RetryTimes")
	delete(f, "Variables")
	delete(f, "UUI")
	delete(f, "CalleeAttributes")
	delete(f, "TimeZone")
	delete(f, "AvailableTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePredictiveDialingCampaignRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePredictiveDialingCampaignResponseParams struct {
	// Generated task id.
	CampaignId *int64 `json:"CampaignId,omitnil,omitempty" name:"CampaignId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreatePredictiveDialingCampaignResponse struct {
	*tchttp.BaseResponse
	Response *CreatePredictiveDialingCampaignResponseParams `json:"Response"`
}

func (r *CreatePredictiveDialingCampaignResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePredictiveDialingCampaignResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSDKLoginTokenRequestParams struct {
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Agent account.
	SeatUserId *string `json:"SeatUserId,omitnil,omitempty" name:"SeatUserId"`

	// Whether the generated token is for one-time verification?.
	OnlyOnce *bool `json:"OnlyOnce,omitnil,omitempty" name:"OnlyOnce"`
}

type CreateSDKLoginTokenRequest struct {
	*tchttp.BaseRequest
	
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Agent account.
	SeatUserId *string `json:"SeatUserId,omitnil,omitempty" name:"SeatUserId"`

	// Whether the generated token is for one-time verification?.
	OnlyOnce *bool `json:"OnlyOnce,omitnil,omitempty" name:"OnlyOnce"`
}

func (r *CreateSDKLoginTokenRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSDKLoginTokenRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "SeatUserId")
	delete(f, "OnlyOnce")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSDKLoginTokenRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSDKLoginTokenResponseParams struct {
	// SDK log-in token.
	Token *string `json:"Token,omitnil,omitempty" name:"Token"`

	// Expiry timestamp. unix timestamp.
	ExpiredTime *int64 `json:"ExpiredTime,omitnil,omitempty" name:"ExpiredTime"`

	// The path in which the sdk is loaded will change with its release.
	SdkURL *string `json:"SdkURL,omitnil,omitempty" name:"SdkURL"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateSDKLoginTokenResponse struct {
	*tchttp.BaseResponse
	Response *CreateSDKLoginTokenResponseParams `json:"Response"`
}

func (r *CreateSDKLoginTokenResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSDKLoginTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateStaffRequestParams struct {
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Customer information, no more than 10.
	Staffs []*SeatUserInfo `json:"Staffs,omitnil,omitempty" name:"Staffs"`

	// Whether to send a password mail or not (the default is true).
	SendPassword *bool `json:"SendPassword,omitnil,omitempty" name:"SendPassword"`
}

type CreateStaffRequest struct {
	*tchttp.BaseRequest
	
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Customer information, no more than 10.
	Staffs []*SeatUserInfo `json:"Staffs,omitnil,omitempty" name:"Staffs"`

	// Whether to send a password mail or not (the default is true).
	SendPassword *bool `json:"SendPassword,omitnil,omitempty" name:"SendPassword"`
}

func (r *CreateStaffRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateStaffRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "Staffs")
	delete(f, "SendPassword")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateStaffRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateStaffResponseParams struct {
	// Error agent list and error information.
	ErrorStaffList []*ErrStaffItem `json:"ErrorStaffList,omitnil,omitempty" name:"ErrorStaffList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateStaffResponse struct {
	*tchttp.BaseResponse
	Response *CreateStaffResponseParams `json:"Response"`
}

func (r *CreateStaffResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateStaffResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUserSigRequestParams struct {
	// App ID (required). can be used to view https://console.cloud.tencent.com/ccc.
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// User ID, must be consistent with the Uid value in the ClientData field.
	Uid *string `json:"Uid,omitnil,omitempty" name:"Uid"`

	// Valid period, in seconds, no more than 1 hr.
	ExpiredTime *int64 `json:"ExpiredTime,omitnil,omitempty" name:"ExpiredTime"`

	// Signature data of the user. required field. standard JSON format.
	ClientData *string `json:"ClientData,omitnil,omitempty" name:"ClientData"`
}

type CreateUserSigRequest struct {
	*tchttp.BaseRequest
	
	// App ID (required). can be used to view https://console.cloud.tencent.com/ccc.
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// User ID, must be consistent with the Uid value in the ClientData field.
	Uid *string `json:"Uid,omitnil,omitempty" name:"Uid"`

	// Valid period, in seconds, no more than 1 hr.
	ExpiredTime *int64 `json:"ExpiredTime,omitnil,omitempty" name:"ExpiredTime"`

	// Signature data of the user. required field. standard JSON format.
	ClientData *string `json:"ClientData,omitnil,omitempty" name:"ClientData"`
}

func (r *CreateUserSigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUserSigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "Uid")
	delete(f, "ExpiredTime")
	delete(f, "ClientData")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateUserSigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUserSigResponseParams struct {
	// Signature result.
	UserSig *string `json:"UserSig,omitnil,omitempty" name:"UserSig"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateUserSigResponse struct {
	*tchttp.BaseResponse
	Response *CreateUserSigResponseParams `json:"Response"`
}

func (r *CreateUserSigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUserSigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCCCSkillGroupRequestParams struct {
	// App ID (required), which can be viewed at https://console.cloud.tencent.com/ccc.
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Skill group ID.
	SkillGroupId *int64 `json:"SkillGroupId,omitnil,omitempty" name:"SkillGroupId"`
}

type DeleteCCCSkillGroupRequest struct {
	*tchttp.BaseRequest
	
	// App ID (required), which can be viewed at https://console.cloud.tencent.com/ccc.
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Skill group ID.
	SkillGroupId *int64 `json:"SkillGroupId,omitnil,omitempty" name:"SkillGroupId"`
}

func (r *DeleteCCCSkillGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCCCSkillGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "SkillGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCCCSkillGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCCCSkillGroupResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteCCCSkillGroupResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCCCSkillGroupResponseParams `json:"Response"`
}

func (r *DeleteCCCSkillGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCCCSkillGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteExtensionRequestParams struct {
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Extension.
	ExtensionId *string `json:"ExtensionId,omitnil,omitempty" name:"ExtensionId"`
}

type DeleteExtensionRequest struct {
	*tchttp.BaseRequest
	
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Extension.
	ExtensionId *string `json:"ExtensionId,omitnil,omitempty" name:"ExtensionId"`
}

func (r *DeleteExtensionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteExtensionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "ExtensionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteExtensionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteExtensionResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteExtensionResponse struct {
	*tchttp.BaseResponse
	Response *DeleteExtensionResponseParams `json:"Response"`
}

func (r *DeleteExtensionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteExtensionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePredictiveDialingCampaignRequestParams struct {
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// <Task id>.
	CampaignId *int64 `json:"CampaignId,omitnil,omitempty" name:"CampaignId"`
}

type DeletePredictiveDialingCampaignRequest struct {
	*tchttp.BaseRequest
	
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// <Task id>.
	CampaignId *int64 `json:"CampaignId,omitnil,omitempty" name:"CampaignId"`
}

func (r *DeletePredictiveDialingCampaignRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePredictiveDialingCampaignRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "CampaignId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeletePredictiveDialingCampaignRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePredictiveDialingCampaignResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeletePredictiveDialingCampaignResponse struct {
	*tchttp.BaseResponse
	Response *DeletePredictiveDialingCampaignResponseParams `json:"Response"`
}

func (r *DeletePredictiveDialingCampaignResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePredictiveDialingCampaignResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteStaffRequestParams struct {
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// List of customer service emails to be deleted, supports up to 200 at a time.
	StaffList []*string `json:"StaffList,omitnil,omitempty" name:"StaffList"`
}

type DeleteStaffRequest struct {
	*tchttp.BaseRequest
	
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// List of customer service emails to be deleted, supports up to 200 at a time.
	StaffList []*string `json:"StaffList,omitnil,omitempty" name:"StaffList"`
}

func (r *DeleteStaffRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteStaffRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "StaffList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteStaffRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteStaffResponseParams struct {
	// List of customer service staff that cannot be deleted when they are online.
	OnlineStaffList []*string `json:"OnlineStaffList,omitnil,omitempty" name:"OnlineStaffList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteStaffResponse struct {
	*tchttp.BaseResponse
	Response *DeleteStaffResponseParams `json:"Response"`
}

func (r *DeleteStaffResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteStaffResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAIAnalysisResultRequestParams struct {
	// App ID (required). can be viewed at https://console.cloud.tencent.com/ccc.
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Specifies the conversation ID.
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// Specifies the search start time.	
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 1737350008
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type DescribeAIAnalysisResultRequest struct {
	*tchttp.BaseRequest
	
	// App ID (required). can be viewed at https://console.cloud.tencent.com/ccc.
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Specifies the conversation ID.
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// Specifies the search start time.	
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 1737350008
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

func (r *DescribeAIAnalysisResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAIAnalysisResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "SessionId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAIAnalysisResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAIAnalysisResultResponseParams struct {
	// AI session analysis result.
	ResultList []*AIAnalysisResult `json:"ResultList,omitnil,omitempty" name:"ResultList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAIAnalysisResultResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAIAnalysisResultResponseParams `json:"Response"`
}

func (r *DescribeAIAnalysisResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAIAnalysisResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAICallExtractResultRequestParams struct {
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Session id.
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// Search for the start time.
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Search for the end time.
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type DescribeAICallExtractResultRequest struct {
	*tchttp.BaseRequest
	
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Session id.
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// Search for the start time.
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Search for the end time.
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

func (r *DescribeAICallExtractResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAICallExtractResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "SessionId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAICallExtractResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAICallExtractResultResponseParams struct {
	// Result list.
	ResultList []*AICallExtractResultElement `json:"ResultList,omitnil,omitempty" name:"ResultList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAICallExtractResultResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAICallExtractResultResponseParams `json:"Response"`
}

func (r *DescribeAICallExtractResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAICallExtractResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAILatencyRequestParams struct {
	// App ID (required), which can be viewed at https://console.cloud.tencent.com/ccc.
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Session ID.
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// Search start time.	
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 1737350008
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type DescribeAILatencyRequest struct {
	*tchttp.BaseRequest
	
	// App ID (required), which can be viewed at https://console.cloud.tencent.com/ccc.
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Session ID.
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// Search start time.	
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 1737350008
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

func (r *DescribeAILatencyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAILatencyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "SessionId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAILatencyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAILatencyResponseParams struct {
	// Detailed data of latency.
	// -1 indicates no corresponding data.
	AILatencyDetail []*AILatencyDetail `json:"AILatencyDetail,omitnil,omitempty" name:"AILatencyDetail"`

	// Latency statistical data.
	// -1 indicates no corresponding data.
	AILatencyStatistics *AILatencyStatistics `json:"AILatencyStatistics,omitnil,omitempty" name:"AILatencyStatistics"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAILatencyResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAILatencyResponseParams `json:"Response"`
}

func (r *DescribeAILatencyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAILatencyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAgentCruiseDialingCampaignRequestParams struct {
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Task id.
	CampaignId *int64 `json:"CampaignId,omitnil,omitempty" name:"CampaignId"`
}

type DescribeAgentCruiseDialingCampaignRequest struct {
	*tchttp.BaseRequest
	
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Task id.
	CampaignId *int64 `json:"CampaignId,omitnil,omitempty" name:"CampaignId"`
}

func (r *DescribeAgentCruiseDialingCampaignRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAgentCruiseDialingCampaignRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "CampaignId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAgentCruiseDialingCampaignRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAgentCruiseDialingCampaignResponseParams struct {
	// Task name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Agent account.
	Agent *string `json:"Agent,omitnil,omitempty" name:"Agent"`

	// Single-Round concurrent call volume 1-20.
	ConcurrencyNumber *int64 `json:"ConcurrencyNumber,omitnil,omitempty" name:"ConcurrencyNumber"`

	// Task start time. unix timestamp. the task will automatically start after this time.
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Task termination time. unix timestamp. the task will automatically terminate after this time.
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Being called sequence: 0 for random 1 for in order.
	CallOrder *int64 `json:"CallOrder,omitnil,omitempty" name:"CallOrder"`

	// Caller custom data, maximum length 1024.
	UUI *string `json:"UUI,omitnil,omitempty" name:"UUI"`

	// Task status 0 not started 1 running 2 completed 3 terminated.
	State *int64 `json:"State,omitnil,omitempty" name:"State"`

	// Total number of called parties.
	TotalCalleeCount *int64 `json:"TotalCalleeCount,omitnil,omitempty" name:"TotalCalleeCount"`

	// Number of calls made and received.
	CalledCalleeCount *int64 `json:"CalledCalleeCount,omitnil,omitempty" name:"CalledCalleeCount"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAgentCruiseDialingCampaignResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAgentCruiseDialingCampaignResponseParams `json:"Response"`
}

func (r *DescribeAgentCruiseDialingCampaignResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAgentCruiseDialingCampaignResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAutoCalloutTaskRequestParams struct {
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Task id.
	TaskId *uint64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type DescribeAutoCalloutTaskRequest struct {
	*tchttp.BaseRequest
	
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Task id.
	TaskId *uint64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *DescribeAutoCalloutTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAutoCalloutTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAutoCalloutTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAutoCalloutTaskResponseParams struct {
	// Task name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <Task description>.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Task start timestamp.
	NotBefore *int64 `json:"NotBefore,omitnil,omitempty" name:"NotBefore"`

	// Task end timestamp.
	// Note: this field may return null, indicating that no valid values can be obtained.
	NotAfter *int64 `json:"NotAfter,omitnil,omitempty" name:"NotAfter"`

	// Calling list.
	Callers []*string `json:"Callers,omitnil,omitempty" name:"Callers"`

	// Called information list.
	Callees []*AutoCalloutTaskCalleeInfo `json:"Callees,omitnil,omitempty" name:"Callees"`

	// IvrId used by the task.
	IvrId *uint64 `json:"IvrId,omitnil,omitempty" name:"IvrId"`

	// Task status: 0 - initial, 1 - running, 2 - completed, 3 - ending, 4 - terminated.
	State *uint64 `json:"State,omitnil,omitempty" name:"State"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAutoCalloutTaskResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAutoCalloutTaskResponseParams `json:"Response"`
}

func (r *DescribeAutoCalloutTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAutoCalloutTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAutoCalloutTasksRequestParams struct {
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// <Page size>.
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Page number.
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`
}

type DescribeAutoCalloutTasksRequest struct {
	*tchttp.BaseRequest
	
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// <Page size>.
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Page number.
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`
}

func (r *DescribeAutoCalloutTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAutoCalloutTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "PageSize")
	delete(f, "PageNumber")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAutoCalloutTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAutoCalloutTasksResponseParams struct {
	// Total quantity.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// <Task list>.
	Tasks []*AutoCalloutTaskInfo `json:"Tasks,omitnil,omitempty" name:"Tasks"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAutoCalloutTasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAutoCalloutTasksResponseParams `json:"Response"`
}

func (r *DescribeAutoCalloutTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAutoCalloutTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCCCBuyInfoListRequestParams struct {
	// Application id list, query all applications when not transmitted.
	SdkAppIds []*int64 `json:"SdkAppIds,omitnil,omitempty" name:"SdkAppIds"`
}

type DescribeCCCBuyInfoListRequest struct {
	*tchttp.BaseRequest
	
	// Application id list, query all applications when not transmitted.
	SdkAppIds []*int64 `json:"SdkAppIds,omitnil,omitempty" name:"SdkAppIds"`
}

func (r *DescribeCCCBuyInfoListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCCCBuyInfoListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCCCBuyInfoListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCCCBuyInfoListResponseParams struct {
	// Total number of applications.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Application purchase information list.
	SdkAppIdBuyList []*SdkAppIdBuyInfo `json:"SdkAppIdBuyList,omitnil,omitempty" name:"SdkAppIdBuyList"`

	// Package purchase information list.
	PackageBuyList []*PackageBuyInfo `json:"PackageBuyList,omitnil,omitempty" name:"PackageBuyList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCCCBuyInfoListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCCCBuyInfoListResponseParams `json:"Response"`
}

func (r *DescribeCCCBuyInfoListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCCCBuyInfoListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCallInMetricsRequestParams struct {
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Whether to return skill group dimension information or not (the default is "yes").
	EnabledSkillGroup *bool `json:"EnabledSkillGroup,omitnil,omitempty" name:"EnabledSkillGroup"`

	// Whether to return line dimension information or not (the default is "no").
	EnabledNumber *bool `json:"EnabledNumber,omitnil,omitempty" name:"EnabledNumber"`

	// Filter skill group list.
	GroupIdList []*int64 `json:"GroupIdList,omitnil,omitempty" name:"GroupIdList"`
}

type DescribeCallInMetricsRequest struct {
	*tchttp.BaseRequest
	
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Whether to return skill group dimension information or not (the default is "yes").
	EnabledSkillGroup *bool `json:"EnabledSkillGroup,omitnil,omitempty" name:"EnabledSkillGroup"`

	// Whether to return line dimension information or not (the default is "no").
	EnabledNumber *bool `json:"EnabledNumber,omitnil,omitempty" name:"EnabledNumber"`

	// Filter skill group list.
	GroupIdList []*int64 `json:"GroupIdList,omitnil,omitempty" name:"GroupIdList"`
}

func (r *DescribeCallInMetricsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCallInMetricsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "EnabledSkillGroup")
	delete(f, "EnabledNumber")
	delete(f, "GroupIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCallInMetricsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCallInMetricsResponseParams struct {
	// Timestamp.
	Timestamp *int64 `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`

	// Overall metrics.
	TotalMetrics *CallInMetrics `json:"TotalMetrics,omitnil,omitempty" name:"TotalMetrics"`

	// Circuit dimension metrics.
	NumberMetrics []*CallInNumberMetrics `json:"NumberMetrics,omitnil,omitempty" name:"NumberMetrics"`

	// Skill group dimension metrics.
	SkillGroupMetrics []*CallInSkillGroupMetrics `json:"SkillGroupMetrics,omitnil,omitempty" name:"SkillGroupMetrics"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCallInMetricsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCallInMetricsResponseParams `json:"Response"`
}

func (r *DescribeCallInMetricsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCallInMetricsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExtensionRequestParams struct {
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Extension.
	ExtensionId *string `json:"ExtensionId,omitnil,omitempty" name:"ExtensionId"`
}

type DescribeExtensionRequest struct {
	*tchttp.BaseRequest
	
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Extension.
	ExtensionId *string `json:"ExtensionId,omitnil,omitempty" name:"ExtensionId"`
}

func (r *DescribeExtensionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExtensionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "ExtensionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeExtensionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExtensionResponseParams struct {
	// Extension.
	ExtensionId *string `json:"ExtensionId,omitnil,omitempty" name:"ExtensionId"`

	// Domain name.
	ExtensionDomain *string `json:"ExtensionDomain,omitnil,omitempty" name:"ExtensionDomain"`

	// Registered password.
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// Proxy server address.
	OutboundProxy *string `json:"OutboundProxy,omitnil,omitempty" name:"OutboundProxy"`

	// Transfer protocol.
	Transport *string `json:"Transport,omitnil,omitempty" name:"Transport"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeExtensionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeExtensionResponseParams `json:"Response"`
}

func (r *DescribeExtensionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExtensionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExtensionsRequestParams struct {
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Page number (starting from 0).
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Filtering extension number list.
	ExtensionIds []*string `json:"ExtensionIds,omitnil,omitempty" name:"ExtensionIds"`

	// Page size.
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Fuzzy query field (fuzzy query for extension number, extension name, agent email, and agent name).
	FuzzingKeyWord *string `json:"FuzzingKeyWord,omitnil,omitempty" name:"FuzzingKeyWord"`

	// Whether to return the current status of the telephone or not.
	IsNeedStatus *bool `json:"IsNeedStatus,omitnil,omitempty" name:"IsNeedStatus"`
}

type DescribeExtensionsRequest struct {
	*tchttp.BaseRequest
	
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Page number (starting from 0).
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Filtering extension number list.
	ExtensionIds []*string `json:"ExtensionIds,omitnil,omitempty" name:"ExtensionIds"`

	// Page size.
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Fuzzy query field (fuzzy query for extension number, extension name, agent email, and agent name).
	FuzzingKeyWord *string `json:"FuzzingKeyWord,omitnil,omitempty" name:"FuzzingKeyWord"`

	// Whether to return the current status of the telephone or not.
	IsNeedStatus *bool `json:"IsNeedStatus,omitnil,omitempty" name:"IsNeedStatus"`
}

func (r *DescribeExtensionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExtensionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "PageNumber")
	delete(f, "ExtensionIds")
	delete(f, "PageSize")
	delete(f, "FuzzingKeyWord")
	delete(f, "IsNeedStatus")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeExtensionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExtensionsResponseParams struct {
	// Total query count.
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// Telephone information list.
	ExtensionList []*ExtensionInfo `json:"ExtensionList,omitnil,omitempty" name:"ExtensionList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeExtensionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeExtensionsResponseParams `json:"Response"`
}

func (r *DescribeExtensionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExtensionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIvrAudioListRequestParams struct {
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Page size, upper limit 50.
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Page number starting from 0.
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// File alias.
	CustomFileName []*string `json:"CustomFileName,omitnil,omitempty" name:"CustomFileName"`

	// Filename.
	AudioFileName []*string `json:"AudioFileName,omitnil,omitempty" name:"AudioFileName"`

	// File id.
	FileId []*uint64 `json:"FileId,omitnil,omitempty" name:"FileId"`
}

type DescribeIvrAudioListRequest struct {
	*tchttp.BaseRequest
	
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Page size, upper limit 50.
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Page number starting from 0.
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// File alias.
	CustomFileName []*string `json:"CustomFileName,omitnil,omitempty" name:"CustomFileName"`

	// Filename.
	AudioFileName []*string `json:"AudioFileName,omitnil,omitempty" name:"AudioFileName"`

	// File id.
	FileId []*uint64 `json:"FileId,omitnil,omitempty" name:"FileId"`
}

func (r *DescribeIvrAudioListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIvrAudioListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "PageSize")
	delete(f, "PageNumber")
	delete(f, "CustomFileName")
	delete(f, "AudioFileName")
	delete(f, "FileId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIvrAudioListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIvrAudioListResponseParams struct {
	// Total quantity.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// File information.
	FileInfo []*AudioFileInfo `json:"FileInfo,omitnil,omitempty" name:"FileInfo"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeIvrAudioListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIvrAudioListResponseParams `json:"Response"`
}

func (r *DescribeIvrAudioListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIvrAudioListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNumbersRequestParams struct {
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Page number, starting from 0.
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Page size, default 20.
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type DescribeNumbersRequest struct {
	*tchttp.BaseRequest
	
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Page number, starting from 0.
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Page size, default 20.
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

func (r *DescribeNumbersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNumbersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNumbersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNumbersResponseParams struct {
	// Total quantity.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Number list.
	Numbers []*NumberInfo `json:"Numbers,omitnil,omitempty" name:"Numbers"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeNumbersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNumbersResponseParams `json:"Response"`
}

func (r *DescribeNumbersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNumbersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePSTNActiveSessionListRequestParams struct {
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Data offset.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned data entries, up to 25.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribePSTNActiveSessionListRequest struct {
	*tchttp.BaseRequest
	
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Data offset.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned data entries, up to 25.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribePSTNActiveSessionListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePSTNActiveSessionListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePSTNActiveSessionListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePSTNActiveSessionListResponseParams struct {
	// Total number of items in the list.
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// List content.
	Sessions []*PSTNSessionInfo `json:"Sessions,omitnil,omitempty" name:"Sessions"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePSTNActiveSessionListResponse struct {
	*tchttp.BaseResponse
	Response *DescribePSTNActiveSessionListResponseParams `json:"Response"`
}

func (r *DescribePSTNActiveSessionListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePSTNActiveSessionListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePredictiveDialingCampaignRequestParams struct {
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// <Task id>.
	CampaignId *int64 `json:"CampaignId,omitnil,omitempty" name:"CampaignId"`
}

type DescribePredictiveDialingCampaignRequest struct {
	*tchttp.BaseRequest
	
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// <Task id>.
	CampaignId *int64 `json:"CampaignId,omitnil,omitempty" name:"CampaignId"`
}

func (r *DescribePredictiveDialingCampaignRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePredictiveDialingCampaignRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "CampaignId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePredictiveDialingCampaignRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePredictiveDialingCampaignResponseParams struct {
	// Task id.
	CampaignId *int64 `json:"CampaignId,omitnil,omitempty" name:"CampaignId"`

	// Task name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Being called sequence: 0 for random 1 for in order.
	CallOrder *int64 `json:"CallOrder,omitnil,omitempty" name:"CallOrder"`

	// ID of the used skill group of agents.
	SkillGroupId *int64 `json:"SkillGroupId,omitnil,omitempty" name:"SkillGroupId"`

	// Specified ivr id.
	IVRId *int64 `json:"IVRId,omitnil,omitempty" name:"IVRId"`

	// Running priority of multiple tasks in the same application, from high to low 1 - 5.
	Priority *int64 `json:"Priority,omitnil,omitempty" name:"Priority"`

	// Expected call drop rate, percentage, 5 - 50.
	ExpectedAbandonRate *int64 `json:"ExpectedAbandonRate,omitnil,omitempty" name:"ExpectedAbandonRate"`

	// Number of call retries, 0 - 2.
	RetryTimes *int64 `json:"RetryTimes,omitnil,omitempty" name:"RetryTimes"`

	// Call retry interval, in seconds, [60 - 86,400].
	RetryInterval *int64 `json:"RetryInterval,omitnil,omitempty" name:"RetryInterval"`

	// Task start time. unix timestamp. the task will automatically start after this time.
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Task termination time. unix timestamp. the task will automatically terminate after this time.
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePredictiveDialingCampaignResponse struct {
	*tchttp.BaseResponse
	Response *DescribePredictiveDialingCampaignResponseParams `json:"Response"`
}

func (r *DescribePredictiveDialingCampaignResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePredictiveDialingCampaignResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePredictiveDialingCampaignsElement struct {
	// <Task id>.
	CampaignId *int64 `json:"CampaignId,omitnil,omitempty" name:"CampaignId"`

	// Task name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Task status 0 - ready to start, 1 - in progress, 2 - paused, 3 - terminated, 4 - completed.
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Task status reasons 0 - normal, 1 - manually ended, 2 - ended due to overtime.
	StatusReason *int64 `json:"StatusReason,omitnil,omitempty" name:"StatusReason"`

	// Number of called numbers.
	CalleeCount *int64 `json:"CalleeCount,omitnil,omitempty" name:"CalleeCount"`

	// Number of completed calls.
	FinishedCalleeCount *int64 `json:"FinishedCalleeCount,omitnil,omitempty" name:"FinishedCalleeCount"`

	// Running priority of multiple tasks in the same application, from high to low 1 - 5.
	Priority *int64 `json:"Priority,omitnil,omitempty" name:"Priority"`

	// ID of the used skill group of agents.
	SkillGroupId *int64 `json:"SkillGroupId,omitnil,omitempty" name:"SkillGroupId"`
}

// Predefined struct for user
type DescribePredictiveDialingCampaignsRequestParams struct {
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Page size, 100 maximum.
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Page number starting from 0.
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Query the task list name keyword.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Query task list skill group id.
	SkillGroupId *int64 `json:"SkillGroupId,omitnil,omitempty" name:"SkillGroupId"`
}

type DescribePredictiveDialingCampaignsRequest struct {
	*tchttp.BaseRequest
	
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Page size, 100 maximum.
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Page number starting from 0.
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Query the task list name keyword.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Query task list skill group id.
	SkillGroupId *int64 `json:"SkillGroupId,omitnil,omitempty" name:"SkillGroupId"`
}

func (r *DescribePredictiveDialingCampaignsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePredictiveDialingCampaignsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "PageSize")
	delete(f, "PageNumber")
	delete(f, "Name")
	delete(f, "SkillGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePredictiveDialingCampaignsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePredictiveDialingCampaignsResponseParams struct {
	// Total data volume.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Data.
	CampaignList []*DescribePredictiveDialingCampaignsElement `json:"CampaignList,omitnil,omitempty" name:"CampaignList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePredictiveDialingCampaignsResponse struct {
	*tchttp.BaseResponse
	Response *DescribePredictiveDialingCampaignsResponseParams `json:"Response"`
}

func (r *DescribePredictiveDialingCampaignsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePredictiveDialingCampaignsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePredictiveDialingSessionsRequestParams struct {
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Generated task id.
	CampaignId *int64 `json:"CampaignId,omitnil,omitempty" name:"CampaignId"`

	// Page size, maximum of 1000.
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Page number starting from 0.
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`
}

type DescribePredictiveDialingSessionsRequest struct {
	*tchttp.BaseRequest
	
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Generated task id.
	CampaignId *int64 `json:"CampaignId,omitnil,omitempty" name:"CampaignId"`

	// Page size, maximum of 1000.
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Page number starting from 0.
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`
}

func (r *DescribePredictiveDialingSessionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePredictiveDialingSessionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "CampaignId")
	delete(f, "PageSize")
	delete(f, "PageNumber")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePredictiveDialingSessionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePredictiveDialingSessionsResponseParams struct {
	// Total data volume.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List of session ids for a call. you can access detailed call bills in batches through https://intl.cloud.tencent.com/document/product/679/47714.?from_cn_redirect=1.
	SessionList []*string `json:"SessionList,omitnil,omitempty" name:"SessionList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePredictiveDialingSessionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribePredictiveDialingSessionsResponseParams `json:"Response"`
}

func (r *DescribePredictiveDialingSessionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePredictiveDialingSessionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProtectedTelCdrRequestParams struct {
	// Start timestamp. unix second-level timestamp.
	StartTimeStamp *int64 `json:"StartTimeStamp,omitnil,omitempty" name:"StartTimeStamp"`

	// End timestamp. unix second-level timestamp.
	EndTimeStamp *int64 `json:"EndTimeStamp,omitnil,omitempty" name:"EndTimeStamp"`

	// For the application id, you can check https://console.cloud.tencent.com/ccc.
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Page size, upper limit 100.
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Page number starting from 0.
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`
}

type DescribeProtectedTelCdrRequest struct {
	*tchttp.BaseRequest
	
	// Start timestamp. unix second-level timestamp.
	StartTimeStamp *int64 `json:"StartTimeStamp,omitnil,omitempty" name:"StartTimeStamp"`

	// End timestamp. unix second-level timestamp.
	EndTimeStamp *int64 `json:"EndTimeStamp,omitnil,omitempty" name:"EndTimeStamp"`

	// For the application id, you can check https://console.cloud.tencent.com/ccc.
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Page size, upper limit 100.
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Page number starting from 0.
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`
}

func (r *DescribeProtectedTelCdrRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProtectedTelCdrRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTimeStamp")
	delete(f, "EndTimeStamp")
	delete(f, "SdkAppId")
	delete(f, "PageSize")
	delete(f, "PageNumber")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProtectedTelCdrRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProtectedTelCdrResponseParams struct {
	// Total number of call records.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Call record.
	//
	// Deprecated: TelCdrs is deprecated.
	TelCdrs []*TelCdrInfo `json:"TelCdrs,omitnil,omitempty" name:"TelCdrs"`

	// Call record.
	TelCdrList []*TelCdrInfo `json:"TelCdrList,omitnil,omitempty" name:"TelCdrList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeProtectedTelCdrResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProtectedTelCdrResponseParams `json:"Response"`
}

func (r *DescribeProtectedTelCdrResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProtectedTelCdrResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSessionDetailRequestParams struct {
	// App ID (required). can be used to view https://console.cloud.tencent.com/ccc.
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Specifies the session id of the call.
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// Start timestamp. Unix second-level timestamp. supports up to nearly 180 days.
	StartTimestamp *int64 `json:"StartTimestamp,omitnil,omitempty" name:"StartTimestamp"`

	// End timestamp, Unix second-level timestamp. the interval range between end time and start time is less than 90 days.
	EndTimestamp *int64 `json:"EndTimestamp,omitnil,omitempty" name:"EndTimestamp"`
}

type DescribeSessionDetailRequest struct {
	*tchttp.BaseRequest
	
	// App ID (required). can be used to view https://console.cloud.tencent.com/ccc.
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Specifies the session id of the call.
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// Start timestamp. Unix second-level timestamp. supports up to nearly 180 days.
	StartTimestamp *int64 `json:"StartTimestamp,omitnil,omitempty" name:"StartTimestamp"`

	// End timestamp, Unix second-level timestamp. the interval range between end time and start time is less than 90 days.
	EndTimestamp *int64 `json:"EndTimestamp,omitnil,omitempty" name:"EndTimestamp"`
}

func (r *DescribeSessionDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSessionDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "SessionId")
	delete(f, "StartTimestamp")
	delete(f, "EndTimestamp")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSessionDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSessionDetailResponseParams struct {
	// Calling number.
	Caller *string `json:"Caller,omitnil,omitempty" name:"Caller"`

	// Called number.
	Callee *string `json:"Callee,omitnil,omitempty" name:"Callee"`

	// Call type. valid values: 1 (outgoing call), 2 (incoming call), 3 (audio dial-in), 5 (predictive outbound call), 6 (internal call).
	CallType *int64 `json:"CallType,omitnil,omitempty" name:"CallType"`

	// Start timestamp. Unix second-level timestamp.
	StartTimeStamp *int64 `json:"StartTimeStamp,omitnil,omitempty" name:"StartTimeStamp"`

	// Ring timestamp. UNIX second-level timestamp.
	RingTimestamp *int64 `json:"RingTimestamp,omitnil,omitempty" name:"RingTimestamp"`

	// Answer timestamp. UNIX second-level timestamp.
	AcceptTimestamp *int64 `json:"AcceptTimestamp,omitnil,omitempty" name:"AcceptTimestamp"`

	// End timestamp, UNIX second-level timestamp.
	EndedTimestamp *int64 `json:"EndedTimestamp,omitnil,omitempty" name:"EndedTimestamp"`

	// Queue entry time. Unix second-level timestamp.
	QueuedTimestamp *int64 `json:"QueuedTimestamp,omitnil,omitempty" name:"QueuedTimestamp"`

	// Agent account.
	StaffUserId *string `json:"StaffUserId,omitnil,omitempty" name:"StaffUserId"`

	// Refers to the EndStatus field in the DescribeTelCdr api.
	EndStatus *int64 `json:"EndStatus,omitnil,omitempty" name:"EndStatus"`

	// Queue skill group ID.
	QueuedSkillGroupId *int64 `json:"QueuedSkillGroupId,omitnil,omitempty" name:"QueuedSkillGroupId"`

	// Queue skill group name.
	QueuedSkillGroupName *string `json:"QueuedSkillGroupName,omitnil,omitempty" name:"QueuedSkillGroupName"`

	// Recording url with authentication and valid period. obtain and pull within a short time frame. do not persist this link.
	RecordURL *string `json:"RecordURL,omitnil,omitempty" name:"RecordURL"`

	// Specifies the COS link for recording transfer to a third party.
	CustomRecordURL *string `json:"CustomRecordURL,omitnil,omitempty" name:"CustomRecordURL"`

	// Recording text information link with authentication and valid period. retrieve it within a short time frame. do not persist this link.
	AsrURL *string `json:"AsrURL,omitnil,omitempty" name:"AsrURL"`

	// Voicemail recording url.
	VoicemailRecordURL []*string `json:"VoicemailRecordURL,omitnil,omitempty" name:"VoicemailRecordURL"`

	// Voicemail recording text information url. purchase the offline speech recognition package through the console and enable the offline speech recognition switch.
	VoicemailAsrURL []*string `json:"VoicemailAsrURL,omitnil,omitempty" name:"VoicemailAsrURL"`

	// IVR key information.
	IVRKeyPressed []*IVRKeyPressedElement `json:"IVRKeyPressed,omitnil,omitempty" name:"IVRKeyPressed"`

	// Satisfaction rate keystroke information.
	PostIVRKeyPressed []*IVRKeyPressedElement `json:"PostIVRKeyPressed,omitnil,omitempty" name:"PostIVRKeyPressed"`

	// Hang-Up side. valid values: seat, user, system.
	HungUpSide *string `json:"HungUpSide,omitnil,omitempty" name:"HungUpSide"`

	// Customer custom data (User-to-User Interface).
	UUI *string `json:"UUI,omitnil,omitempty" name:"UUI"`

	// List of events during a call.
	Events []*SessionEvent `json:"Events,omitnil,omitempty" name:"Events"`

	// List of service participants.
	ServeParticipants []*ServeParticipant `json:"ServeParticipants,omitnil,omitempty" name:"ServeParticipants"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSessionDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSessionDetailResponseParams `json:"Response"`
}

func (r *DescribeSessionDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSessionDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSkillGroupInfoListRequestParams struct {
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Page size, upper limit 100.
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <Page number starting from 0.>.
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Using skill group id when querying a single skill group.
	SkillGroupId *int64 `json:"SkillGroupId,omitnil,omitempty" name:"SkillGroupId"`

	// Used when querying skill groups with a modified time greater or equal to modifiedtime.
	ModifiedTime *int64 `json:"ModifiedTime,omitnil,omitempty" name:"ModifiedTime"`

	// Skill group name.
	SkillGroupName *string `json:"SkillGroupName,omitnil,omitempty" name:"SkillGroupName"`
}

type DescribeSkillGroupInfoListRequest struct {
	*tchttp.BaseRequest
	
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Page size, upper limit 100.
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <Page number starting from 0.>.
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Using skill group id when querying a single skill group.
	SkillGroupId *int64 `json:"SkillGroupId,omitnil,omitempty" name:"SkillGroupId"`

	// Used when querying skill groups with a modified time greater or equal to modifiedtime.
	ModifiedTime *int64 `json:"ModifiedTime,omitnil,omitempty" name:"ModifiedTime"`

	// Skill group name.
	SkillGroupName *string `json:"SkillGroupName,omitnil,omitempty" name:"SkillGroupName"`
}

func (r *DescribeSkillGroupInfoListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSkillGroupInfoListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "PageSize")
	delete(f, "PageNumber")
	delete(f, "SkillGroupId")
	delete(f, "ModifiedTime")
	delete(f, "SkillGroupName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSkillGroupInfoListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSkillGroupInfoListResponseParams struct {
	// Total number of skill groups.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Skill group information list.
	SkillGroupList []*SkillGroupInfoItem `json:"SkillGroupList,omitnil,omitempty" name:"SkillGroupList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSkillGroupInfoListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSkillGroupInfoListResponseParams `json:"Response"`
}

func (r *DescribeSkillGroupInfoListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSkillGroupInfoListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStaffInfoListRequestParams struct {
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Page size, upper limit 9,999.
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Page number starting from 0.
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Agent account used when querying a single agent.
	StaffMail *string `json:"StaffMail,omitnil,omitempty" name:"StaffMail"`

	// Use when querying for agents with a modification time greater or equal to modifiedtime.
	ModifiedTime *int64 `json:"ModifiedTime,omitnil,omitempty" name:"ModifiedTime"`

	// Skill group id.
	SkillGroupId *int64 `json:"SkillGroupId,omitnil,omitempty" name:"SkillGroupId"`
}

type DescribeStaffInfoListRequest struct {
	*tchttp.BaseRequest
	
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Page size, upper limit 9,999.
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Page number starting from 0.
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Agent account used when querying a single agent.
	StaffMail *string `json:"StaffMail,omitnil,omitempty" name:"StaffMail"`

	// Use when querying for agents with a modification time greater or equal to modifiedtime.
	ModifiedTime *int64 `json:"ModifiedTime,omitnil,omitempty" name:"ModifiedTime"`

	// Skill group id.
	SkillGroupId *int64 `json:"SkillGroupId,omitnil,omitempty" name:"SkillGroupId"`
}

func (r *DescribeStaffInfoListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStaffInfoListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "PageSize")
	delete(f, "PageNumber")
	delete(f, "StaffMail")
	delete(f, "ModifiedTime")
	delete(f, "SkillGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStaffInfoListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStaffInfoListResponseParams struct {
	// Total number of agent users.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Agent user information list.
	StaffList []*StaffInfo `json:"StaffList,omitnil,omitempty" name:"StaffList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeStaffInfoListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStaffInfoListResponseParams `json:"Response"`
}

func (r *DescribeStaffInfoListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStaffInfoListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStaffStatusHistoryRequestParams struct {
	// App ID (required). can be used to view https://console.cloud.tencent.com/ccc.
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Agent account.
	StaffUserId *string `json:"StaffUserId,omitnil,omitempty" name:"StaffUserId"`

	// Start timestamp. Unix second-level timestamp. supports up to nearly 180 days.
	StartTimestamp *int64 `json:"StartTimestamp,omitnil,omitempty" name:"StartTimestamp"`

	// End timestamp, Unix second-level timestamp. the interval range between end time and start time is less than 7 days.
	EndTimestamp *int64 `json:"EndTimestamp,omitnil,omitempty" name:"EndTimestamp"`

	// Specifies the cursor used during paginated retrieval.
	Cursor *string `json:"Cursor,omitnil,omitempty" name:"Cursor"`

	// Specifies the pagination size.
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type DescribeStaffStatusHistoryRequest struct {
	*tchttp.BaseRequest
	
	// App ID (required). can be used to view https://console.cloud.tencent.com/ccc.
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Agent account.
	StaffUserId *string `json:"StaffUserId,omitnil,omitempty" name:"StaffUserId"`

	// Start timestamp. Unix second-level timestamp. supports up to nearly 180 days.
	StartTimestamp *int64 `json:"StartTimestamp,omitnil,omitempty" name:"StartTimestamp"`

	// End timestamp, Unix second-level timestamp. the interval range between end time and start time is less than 7 days.
	EndTimestamp *int64 `json:"EndTimestamp,omitnil,omitempty" name:"EndTimestamp"`

	// Specifies the cursor used during paginated retrieval.
	Cursor *string `json:"Cursor,omitnil,omitempty" name:"Cursor"`

	// Specifies the pagination size.
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

func (r *DescribeStaffStatusHistoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStaffStatusHistoryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "StaffUserId")
	delete(f, "StartTimestamp")
	delete(f, "EndTimestamp")
	delete(f, "Cursor")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStaffStatusHistoryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStaffStatusHistoryResponseParams struct {
	// Specifies the agent status data.
	Data []*StaffStatus `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeStaffStatusHistoryResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStaffStatusHistoryResponseParams `json:"Response"`
}

func (r *DescribeStaffStatusHistoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStaffStatusHistoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStaffStatusMetricsRequestParams struct {
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Filter agent list. by default, do not pass all returned agent information.
	StaffList []*string `json:"StaffList,omitnil,omitempty" name:"StaffList"`

	// Filter skill group id list.
	GroupIdList []*int64 `json:"GroupIdList,omitnil,omitempty" name:"GroupIdList"`

	// Filter agent status list agent status free available | busy busy | rest on break | notready not ready | aftercallwork post-call adjustment | offline offline . 
	StatusList []*string `json:"StatusList,omitnil,omitempty" name:"StatusList"`
}

type DescribeStaffStatusMetricsRequest struct {
	*tchttp.BaseRequest
	
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Filter agent list. by default, do not pass all returned agent information.
	StaffList []*string `json:"StaffList,omitnil,omitempty" name:"StaffList"`

	// Filter skill group id list.
	GroupIdList []*int64 `json:"GroupIdList,omitnil,omitempty" name:"GroupIdList"`

	// Filter agent status list agent status free available | busy busy | rest on break | notready not ready | aftercallwork post-call adjustment | offline offline . 
	StatusList []*string `json:"StatusList,omitnil,omitempty" name:"StatusList"`
}

func (r *DescribeStaffStatusMetricsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStaffStatusMetricsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "StaffList")
	delete(f, "GroupIdList")
	delete(f, "StatusList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStaffStatusMetricsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStaffStatusMetricsResponseParams struct {
	// Real-Time information on agent status.
	Metrics []*StaffStatusMetrics `json:"Metrics,omitnil,omitempty" name:"Metrics"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeStaffStatusMetricsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStaffStatusMetricsResponseParams `json:"Response"`
}

func (r *DescribeStaffStatusMetricsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStaffStatusMetricsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTelCallInfoRequestParams struct {
	// Start timestamp, unix timestamp (query dimension supports only daily. for example, to query may 1st, pass starttime:"2023-05-01 00:00:00","endtime":"2023-05-01 23:59:59" timestamp. to query may 1st and may 2nd, pass starttime:"2023-05-01 00:00:00","endtime":"2023-05-02 23:59:59" timestamp).
	StartTimeStamp *int64 `json:"StartTimeStamp,omitnil,omitempty" name:"StartTimeStamp"`

	// End timestamp, unix timestamp, the query time range is up to 90 days (query dimension supports only daily. for example, to query may 1st, pass starttime:"2023-05-01 00:00:00","endtime":"2023-05-01 23:59:59" timestamp. to query may 1st and may 2nd, pass starttime:"2023-05-01 00:00:00","endtime":"2023-05-02 23:59:59" timestamp).
	EndTimeStamp *int64 `json:"EndTimeStamp,omitnil,omitempty" name:"EndTimeStamp"`

	// Application id list, when having multiple ids, the returned value is the sum of all the ids.
	SdkAppIdList []*int64 `json:"SdkAppIdList,omitnil,omitempty" name:"SdkAppIdList"`
}

type DescribeTelCallInfoRequest struct {
	*tchttp.BaseRequest
	
	// Start timestamp, unix timestamp (query dimension supports only daily. for example, to query may 1st, pass starttime:"2023-05-01 00:00:00","endtime":"2023-05-01 23:59:59" timestamp. to query may 1st and may 2nd, pass starttime:"2023-05-01 00:00:00","endtime":"2023-05-02 23:59:59" timestamp).
	StartTimeStamp *int64 `json:"StartTimeStamp,omitnil,omitempty" name:"StartTimeStamp"`

	// End timestamp, unix timestamp, the query time range is up to 90 days (query dimension supports only daily. for example, to query may 1st, pass starttime:"2023-05-01 00:00:00","endtime":"2023-05-01 23:59:59" timestamp. to query may 1st and may 2nd, pass starttime:"2023-05-01 00:00:00","endtime":"2023-05-02 23:59:59" timestamp).
	EndTimeStamp *int64 `json:"EndTimeStamp,omitnil,omitempty" name:"EndTimeStamp"`

	// Application id list, when having multiple ids, the returned value is the sum of all the ids.
	SdkAppIdList []*int64 `json:"SdkAppIdList,omitnil,omitempty" name:"SdkAppIdList"`
}

func (r *DescribeTelCallInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTelCallInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTimeStamp")
	delete(f, "EndTimeStamp")
	delete(f, "SdkAppIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTelCallInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTelCallInfoResponseParams struct {
	// Number of minutes consumed by outbound package.
	TelCallOutCount *int64 `json:"TelCallOutCount,omitnil,omitempty" name:"TelCallOutCount"`

	// Number of minutes consumed by inbound package.
	TelCallInCount *int64 `json:"TelCallInCount,omitnil,omitempty" name:"TelCallInCount"`

	// Number of agent usage statistics.
	SeatUsedCount *int64 `json:"SeatUsedCount,omitnil,omitempty" name:"SeatUsedCount"`

	// Number of minutes consumed by audio package.
	//
	// Deprecated: VoipCallInCount is deprecated.
	VoipCallInCount *int64 `json:"VoipCallInCount,omitnil,omitempty" name:"VoipCallInCount"`

	// Number of minutes consumed by audio package.
	VOIPCallInCount *int64 `json:"VOIPCallInCount,omitnil,omitempty" name:"VOIPCallInCount"`

	// Number of minutes consumed by offline speech-to-text package.
	AsrOfflineCount *int64 `json:"AsrOfflineCount,omitnil,omitempty" name:"AsrOfflineCount"`

	// Number of minutes consumed by real-time speech-to-text package.
	AsrRealtimeCount *int64 `json:"AsrRealtimeCount,omitnil,omitempty" name:"AsrRealtimeCount"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTelCallInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTelCallInfoResponseParams `json:"Response"`
}

func (r *DescribeTelCallInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTelCallInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTelCdrRequestParams struct {
	// Start timestamp, unix timestamp in seconds. supports up to the past 180 days.
	StartTimeStamp *int64 `json:"StartTimeStamp,omitnil,omitempty" name:"StartTimeStamp"`

	// End timestamp, unix timestamp in seconds. the range between the end time and start time is less than 90 days.
	EndTimeStamp *int64 `json:"EndTimeStamp,omitnil,omitempty" name:"EndTimeStamp"`

	// Instance id (deprecated).
	//
	// Deprecated: InstanceId is deprecated.
	InstanceId *int64 `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Maximum number of returned entries (deprecated).
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset (deprecated).
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Page size (required), up to 100.
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <Page number (required), starting from 0.>.
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Filter by phone number.
	Phones []*string `json:"Phones,omitnil,omitempty" name:"Phones"`

	// Filter by sessionid.
	SessionIds []*string `json:"SessionIds,omitnil,omitempty" name:"SessionIds"`
}

type DescribeTelCdrRequest struct {
	*tchttp.BaseRequest
	
	// Start timestamp, unix timestamp in seconds. supports up to the past 180 days.
	StartTimeStamp *int64 `json:"StartTimeStamp,omitnil,omitempty" name:"StartTimeStamp"`

	// End timestamp, unix timestamp in seconds. the range between the end time and start time is less than 90 days.
	EndTimeStamp *int64 `json:"EndTimeStamp,omitnil,omitempty" name:"EndTimeStamp"`

	// Instance id (deprecated).
	InstanceId *int64 `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Maximum number of returned entries (deprecated).
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset (deprecated).
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Page size (required), up to 100.
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <Page number (required), starting from 0.>.
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Filter by phone number.
	Phones []*string `json:"Phones,omitnil,omitempty" name:"Phones"`

	// Filter by sessionid.
	SessionIds []*string `json:"SessionIds,omitnil,omitempty" name:"SessionIds"`
}

func (r *DescribeTelCdrRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTelCdrRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTimeStamp")
	delete(f, "EndTimeStamp")
	delete(f, "InstanceId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "SdkAppId")
	delete(f, "PageSize")
	delete(f, "PageNumber")
	delete(f, "Phones")
	delete(f, "SessionIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTelCdrRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTelCdrResponseParams struct {
	// Total number of call records.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Call record.
	//
	// Deprecated: TelCdrs is deprecated.
	TelCdrs []*TelCdrInfo `json:"TelCdrs,omitnil,omitempty" name:"TelCdrs"`

	// Call record.
	TelCdrList []*TelCdrInfo `json:"TelCdrList,omitnil,omitempty" name:"TelCdrList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTelCdrResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTelCdrResponseParams `json:"Response"`
}

func (r *DescribeTelCdrResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTelCdrResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTelRecordAsrRequestParams struct {
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Session id.
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`
}

type DescribeTelRecordAsrRequest struct {
	*tchttp.BaseRequest
	
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Session id.
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`
}

func (r *DescribeTelRecordAsrRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTelRecordAsrRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "SessionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTelRecordAsrRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTelRecordAsrResponseParams struct {
	// Recording to text information.
	AsrDataList []*AsrData `json:"AsrDataList,omitnil,omitempty" name:"AsrDataList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTelRecordAsrResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTelRecordAsrResponseParams `json:"Response"`
}

func (r *DescribeTelRecordAsrResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTelRecordAsrResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTelSessionRequestParams struct {
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Session id.
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`
}

type DescribeTelSessionRequest struct {
	*tchttp.BaseRequest
	
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Session id.
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`
}

func (r *DescribeTelSessionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTelSessionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "SessionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTelSessionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTelSessionResponseParams struct {
	// Session information.
	Session *PSTNSession `json:"Session,omitnil,omitempty" name:"Session"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTelSessionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTelSessionResponseParams `json:"Response"`
}

func (r *DescribeTelSessionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTelSessionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableCCCPhoneNumberRequestParams struct {
	// Number list starting with 0086.
	PhoneNumbers []*string `json:"PhoneNumbers,omitnil,omitempty" name:"PhoneNumbers"`

	// Disable switch: 0 for enable, 1 for disable.
	Disabled *int64 `json:"Disabled,omitnil,omitempty" name:"Disabled"`

	// TCCC instance application id.
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`
}

type DisableCCCPhoneNumberRequest struct {
	*tchttp.BaseRequest
	
	// Number list starting with 0086.
	PhoneNumbers []*string `json:"PhoneNumbers,omitnil,omitempty" name:"PhoneNumbers"`

	// Disable switch: 0 for enable, 1 for disable.
	Disabled *int64 `json:"Disabled,omitnil,omitempty" name:"Disabled"`

	// TCCC instance application id.
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`
}

func (r *DisableCCCPhoneNumberRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableCCCPhoneNumberRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PhoneNumbers")
	delete(f, "Disabled")
	delete(f, "SdkAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisableCCCPhoneNumberRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableCCCPhoneNumberResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DisableCCCPhoneNumberResponse struct {
	*tchttp.BaseResponse
	Response *DisableCCCPhoneNumberResponseParams `json:"Response"`
}

func (r *DisableCCCPhoneNumberResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableCCCPhoneNumberResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ErrStaffItem struct {
	// Agent email address.
	StaffEmail *string `json:"StaffEmail,omitnil,omitempty" name:"StaffEmail"`

	// Error code.
	Code *string `json:"Code,omitnil,omitempty" name:"Code"`

	// Error description.
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`
}

type EventStaffDetail struct {
	// Specifies the agent data.
	Staffs []*EventStaffElement `json:"Staffs,omitnil,omitempty" name:"Staffs"`
}

type EventStaffElement struct {
	// Agent email address.
	Mail *string `json:"Mail,omitnil,omitempty" name:"Mail"`

	// Agent id.
	StaffNumber *string `json:"StaffNumber,omitnil,omitempty" name:"StaffNumber"`
}

type ExtensionInfo struct {
	// Instance id.
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Extension full name.
	FullExtensionId *string `json:"FullExtensionId,omitnil,omitempty" name:"FullExtensionId"`

	// Extension.
	ExtensionId *string `json:"ExtensionId,omitnil,omitempty" name:"ExtensionId"`

	// Affiliated skill group list.
	SkillGroupId *string `json:"SkillGroupId,omitnil,omitempty" name:"SkillGroupId"`

	// Extension name.
	ExtensionName *string `json:"ExtensionName,omitnil,omitempty" name:"ExtensionName"`

	// Creation time.
	CreateTime *int64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Last modification time.
	ModifyTime *int64 `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`

	// Telephone status (0 offline, 100 free, 200 busy).
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Whether to register.
	Register *bool `json:"Register,omitnil,omitempty" name:"Register"`

	// Bind agent email.
	Relation *string `json:"Relation,omitnil,omitempty" name:"Relation"`

	// Bind agent name.
	RelationName *string `json:"RelationName,omitnil,omitempty" name:"RelationName"`
}

// Predefined struct for user
type ForceMemberOfflineRequestParams struct {
	// App ID (required), which can be viewed at https://console.cloud.tencent.com/ccc.
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Customer service ID.
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

type ForceMemberOfflineRequest struct {
	*tchttp.BaseRequest
	
	// App ID (required), which can be viewed at https://console.cloud.tencent.com/ccc.
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Customer service ID.
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

func (r *ForceMemberOfflineRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ForceMemberOfflineRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "UserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ForceMemberOfflineRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ForceMemberOfflineResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ForceMemberOfflineResponse struct {
	*tchttp.BaseResponse
	Response *ForceMemberOfflineResponseParams `json:"Response"`
}

func (r *ForceMemberOfflineResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ForceMemberOfflineResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ForwardingConfig struct {
	// Whether to enable.
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// 1 unconditional call forwarding 2 conditional call forwarding.
	Condition *int64 `json:"Condition,omitnil,omitempty" name:"Condition"`

	// Call forwarding destination.
	Target *ForwardingTarget `json:"Target,omitnil,omitempty" name:"Target"`
}

type ForwardingTarget struct {
	// Call forwarding target type. valid values: 1 (agent), 2 (skill group), 3 (extension).
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// Specifies the agent account when the call forwarding target is set to agent and Type is 1.
	StaffUserId *string `json:"StaffUserId,omitnil,omitempty" name:"StaffUserId"`

	// Specifies the ID of the skill group as the call forwarding target. fill when Type is 2.
	SkillGroupId *int64 `json:"SkillGroupId,omitnil,omitempty" name:"SkillGroupId"`

	// The account to which the call is forwarded is an extension. fill when Type is 3.
	Extension *string `json:"Extension,omitnil,omitempty" name:"Extension"`
}

// Predefined struct for user
type HangUpCallRequestParams struct {
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Session id.
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`
}

type HangUpCallRequest struct {
	*tchttp.BaseRequest
	
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Session id.
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`
}

func (r *HangUpCallRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *HangUpCallRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "SessionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "HangUpCallRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type HangUpCallResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type HangUpCallResponse struct {
	*tchttp.BaseResponse
	Response *HangUpCallResponseParams `json:"Response"`
}

func (r *HangUpCallResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *HangUpCallResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IVRKeyPressedElement struct {
	// Hit keyword or press.
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// Tag associated with the key.
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`

	// UNIX millisecond timestamp.
	Timestamp *int64 `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`

	// Node tags.
	NodeLabel *string `json:"NodeLabel,omitnil,omitempty" name:"NodeLabel"`

	// User'S original input.
	OriginalContent *string `json:"OriginalContent,omitnil,omitempty" name:"OriginalContent"`

	// TTS prompt content.
	TTSPrompt *string `json:"TTSPrompt,omitnil,omitempty" name:"TTSPrompt"`
}

type Interface struct {
	// API address
	URL *string `json:"URL,omitnil,omitempty" name:"URL"`
}

// Predefined struct for user
type ModifyExtensionRequestParams struct {
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Extension.
	ExtensionId *string `json:"ExtensionId,omitnil,omitempty" name:"ExtensionId"`

	// Extension name.
	ExtensionName *string `json:"ExtensionName,omitnil,omitempty" name:"ExtensionName"`

	// Affiliated skill group list.
	SkillGroupIds []*int64 `json:"SkillGroupIds,omitnil,omitempty" name:"SkillGroupIds"`

	// Bind agent email account.
	Relation *string `json:"Relation,omitnil,omitempty" name:"Relation"`
}

type ModifyExtensionRequest struct {
	*tchttp.BaseRequest
	
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Extension.
	ExtensionId *string `json:"ExtensionId,omitnil,omitempty" name:"ExtensionId"`

	// Extension name.
	ExtensionName *string `json:"ExtensionName,omitnil,omitempty" name:"ExtensionName"`

	// Affiliated skill group list.
	SkillGroupIds []*int64 `json:"SkillGroupIds,omitnil,omitempty" name:"SkillGroupIds"`

	// Bind agent email account.
	Relation *string `json:"Relation,omitnil,omitempty" name:"Relation"`
}

func (r *ModifyExtensionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyExtensionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "ExtensionId")
	delete(f, "ExtensionName")
	delete(f, "SkillGroupIds")
	delete(f, "Relation")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyExtensionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyExtensionResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyExtensionResponse struct {
	*tchttp.BaseResponse
	Response *ModifyExtensionResponseParams `json:"Response"`
}

func (r *ModifyExtensionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyExtensionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyOwnNumberApplyRequestParams struct {
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Circuit-Related parameters.
	DetailList []*OwnNumberApplyDetailItem `json:"DetailList,omitnil,omitempty" name:"DetailList"`

	// Approval id.
	ApplyId *uint64 `json:"ApplyId,omitnil,omitempty" name:"ApplyId"`

	// Prefix for sending numbers.
	Prefix *string `json:"Prefix,omitnil,omitempty" name:"Prefix"`
}

type ModifyOwnNumberApplyRequest struct {
	*tchttp.BaseRequest
	
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Circuit-Related parameters.
	DetailList []*OwnNumberApplyDetailItem `json:"DetailList,omitnil,omitempty" name:"DetailList"`

	// Approval id.
	ApplyId *uint64 `json:"ApplyId,omitnil,omitempty" name:"ApplyId"`

	// Prefix for sending numbers.
	Prefix *string `json:"Prefix,omitnil,omitempty" name:"Prefix"`
}

func (r *ModifyOwnNumberApplyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyOwnNumberApplyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "DetailList")
	delete(f, "ApplyId")
	delete(f, "Prefix")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyOwnNumberApplyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyOwnNumberApplyResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyOwnNumberApplyResponse struct {
	*tchttp.BaseResponse
	Response *ModifyOwnNumberApplyResponseParams `json:"Response"`
}

func (r *ModifyOwnNumberApplyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyOwnNumberApplyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyStaffPasswordRequestParams struct {
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Agent email.
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// The set password.
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`
}

type ModifyStaffPasswordRequest struct {
	*tchttp.BaseRequest
	
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Agent email.
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// The set password.
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`
}

func (r *ModifyStaffPasswordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyStaffPasswordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "Email")
	delete(f, "Password")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyStaffPasswordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyStaffPasswordResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyStaffPasswordResponse struct {
	*tchttp.BaseResponse
	Response *ModifyStaffPasswordResponseParams `json:"Response"`
}

func (r *ModifyStaffPasswordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyStaffPasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyStaffRequestParams struct {
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Agent account.
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// Agent name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Agent phone number (preceded by 0086, example: 008618011111111).
	Phone *string `json:"Phone,omitnil,omitempty" name:"Phone"`

	// Agent nickname.
	Nick *string `json:"Nick,omitnil,omitempty" name:"Nick"`

	// Agent id.
	StaffNo *string `json:"StaffNo,omitnil,omitempty" name:"StaffNo"`

	// Bind skill group id list.
	SkillGroupIds []*int64 `json:"SkillGroupIds,omitnil,omitempty" name:"SkillGroupIds"`

	// Whether the cell phone outbound call switch is enabled or not.
	UseMobileCallOut *bool `json:"UseMobileCallOut,omitnil,omitempty" name:"UseMobileCallOut"`

	// Cell phone answering pattern: 0 - off | 1 - only when offline | 2 - always.
	UseMobileAccept *int64 `json:"UseMobileAccept,omitnil,omitempty" name:"UseMobileAccept"`

	// Agent extension number (starting with 1 to 8, 4 - 6 digits).
	ExtensionNumber *string `json:"ExtensionNumber,omitnil,omitempty" name:"ExtensionNumber"`

	// Call forwarding configuration.
	ForwardingConfig *ForwardingConfig `json:"ForwardingConfig,omitnil,omitempty" name:"ForwardingConfig"`
}

type ModifyStaffRequest struct {
	*tchttp.BaseRequest
	
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Agent account.
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// Agent name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Agent phone number (preceded by 0086, example: 008618011111111).
	Phone *string `json:"Phone,omitnil,omitempty" name:"Phone"`

	// Agent nickname.
	Nick *string `json:"Nick,omitnil,omitempty" name:"Nick"`

	// Agent id.
	StaffNo *string `json:"StaffNo,omitnil,omitempty" name:"StaffNo"`

	// Bind skill group id list.
	SkillGroupIds []*int64 `json:"SkillGroupIds,omitnil,omitempty" name:"SkillGroupIds"`

	// Whether the cell phone outbound call switch is enabled or not.
	UseMobileCallOut *bool `json:"UseMobileCallOut,omitnil,omitempty" name:"UseMobileCallOut"`

	// Cell phone answering pattern: 0 - off | 1 - only when offline | 2 - always.
	UseMobileAccept *int64 `json:"UseMobileAccept,omitnil,omitempty" name:"UseMobileAccept"`

	// Agent extension number (starting with 1 to 8, 4 - 6 digits).
	ExtensionNumber *string `json:"ExtensionNumber,omitnil,omitempty" name:"ExtensionNumber"`

	// Call forwarding configuration.
	ForwardingConfig *ForwardingConfig `json:"ForwardingConfig,omitnil,omitempty" name:"ForwardingConfig"`
}

func (r *ModifyStaffRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyStaffRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "Email")
	delete(f, "Name")
	delete(f, "Phone")
	delete(f, "Nick")
	delete(f, "StaffNo")
	delete(f, "SkillGroupIds")
	delete(f, "UseMobileCallOut")
	delete(f, "UseMobileAccept")
	delete(f, "ExtensionNumber")
	delete(f, "ForwardingConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyStaffRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyStaffResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyStaffResponse struct {
	*tchttp.BaseResponse
	Response *ModifyStaffResponseParams `json:"Response"`
}

func (r *ModifyStaffResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyStaffResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NumberInfo struct {
	// Number.
	Number *string `json:"Number,omitnil,omitempty" name:"Number"`

	// Bound outbound call skill group.
	CallOutSkillGroupIds []*uint64 `json:"CallOutSkillGroupIds,omitnil,omitempty" name:"CallOutSkillGroupIds"`

	// Number status, 1-normal, 2-disabled due to overdue payment, 4-disabled by the administrator, 5-disabled due to violation.
	State *int64 `json:"State,omitnil,omitempty" name:"State"`
}

type OwnNumberApplyDetailItem struct {
	// Number type: 0 - inbound | 1 - outbound | 2 - inbound and outbound.
	CallType *int64 `json:"CallType,omitnil,omitempty" name:"CallType"`

	// Line number.
	PhoneNumber *string `json:"PhoneNumber,omitnil,omitempty" name:"PhoneNumber"`

	// Maximum concurrent call number.
	MaxCallCount *int64 `json:"MaxCallCount,omitnil,omitempty" name:"MaxCallCount"`

	// Maximum number of concurrencies per second.
	MaxCallPSec *int64 `json:"MaxCallPSec,omitnil,omitempty" name:"MaxCallPSec"`

	// Outbound called number format, use {+e.164} or {e.164}. 
	OutboundCalleeFormat *string `json:"OutboundCalleeFormat,omitnil,omitempty" name:"OutboundCalleeFormat"`
}

type PSTNSession struct {
	// Session id.
	SessionID *string `json:"SessionID,omitnil,omitempty" name:"SessionID"`

	// Temporary room id for session.
	RoomID *string `json:"RoomID,omitnil,omitempty" name:"RoomID"`

	// Caller.
	Caller *string `json:"Caller,omitnil,omitempty" name:"Caller"`

	// Called.
	Callee *string `json:"Callee,omitnil,omitempty" name:"Callee"`

	// Start time. unix timestamp.
	StartTimestamp *int64 `json:"StartTimestamp,omitnil,omitempty" name:"StartTimestamp"`

	// Ring time. unix timestamp.
	RingTimestamp *int64 `json:"RingTimestamp,omitnil,omitempty" name:"RingTimestamp"`

	// Answer time. unix timestamp.
	AcceptTimestamp *int64 `json:"AcceptTimestamp,omitnil,omitempty" name:"AcceptTimestamp"`

	// Agent email.
	StaffEmail *string `json:"StaffEmail,omitnil,omitempty" name:"StaffEmail"`

	// Agent id.
	StaffNumber *string `json:"StaffNumber,omitnil,omitempty" name:"StaffNumber"`

	// Session status.
	// Ringing - in progress.
	// SeatJoining - waiting for the agent to answer.
	// InProgress: in progress.
	// Finished - completed.
	SessionStatus *string `json:"SessionStatus,omitnil,omitempty" name:"SessionStatus"`

	// Session call direction, 0 - inbound | 1 - outbound.
	Direction *int64 `json:"Direction,omitnil,omitempty" name:"Direction"`

	// The number used for transferring to the external line (outbound caller).
	OutBoundCaller *string `json:"OutBoundCaller,omitnil,omitempty" name:"OutBoundCaller"`

	// Outbound callee.
	OutBoundCallee *string `json:"OutBoundCallee,omitnil,omitempty" name:"OutBoundCallee"`

	// Caller number protection id. effective when the number protection map feature is activated, and the caller field is empty.
	ProtectedCaller *string `json:"ProtectedCaller,omitnil,omitempty" name:"ProtectedCaller"`

	// Called number protection id. effective when the number protection map feature is activated, and the callee field is empty.
	ProtectedCallee *string `json:"ProtectedCallee,omitnil,omitempty" name:"ProtectedCallee"`
}

type PSTNSessionInfo struct {
	// Session id.
	SessionID *string `json:"SessionID,omitnil,omitempty" name:"SessionID"`

	// Temporary room id for session.
	RoomID *string `json:"RoomID,omitnil,omitempty" name:"RoomID"`

	// Caller.
	Caller *string `json:"Caller,omitnil,omitempty" name:"Caller"`

	// Called.
	Callee *string `json:"Callee,omitnil,omitempty" name:"Callee"`

	// Start time. unix timestamp.
	StartTimestamp *string `json:"StartTimestamp,omitnil,omitempty" name:"StartTimestamp"`

	// Answer time. unix timestamp.
	AcceptTimestamp *string `json:"AcceptTimestamp,omitnil,omitempty" name:"AcceptTimestamp"`

	// Agent email.
	StaffEmail *string `json:"StaffEmail,omitnil,omitempty" name:"StaffEmail"`

	// Agent id.
	StaffNumber *string `json:"StaffNumber,omitnil,omitempty" name:"StaffNumber"`

	// Agent status inprogress ongoing.
	SessionStatus *string `json:"SessionStatus,omitnil,omitempty" name:"SessionStatus"`

	// Session call direction, 0 - inbound | 1 - outbound.
	Direction *int64 `json:"Direction,omitnil,omitempty" name:"Direction"`

	// Ring time. unix timestamp.
	RingTimestamp *int64 `json:"RingTimestamp,omitnil,omitempty" name:"RingTimestamp"`

	// Caller number protection id. effective when the number protection map feature is activated, and the caller field is empty.
	ProtectedCaller *string `json:"ProtectedCaller,omitnil,omitempty" name:"ProtectedCaller"`

	// Called number protection id. effective when the number protection map feature is activated, and the callee field is empty.
	ProtectedCallee *string `json:"ProtectedCallee,omitnil,omitempty" name:"ProtectedCallee"`
}

type PackageBuyInfo struct {
	// Package id.
	PackageId *string `json:"PackageId,omitnil,omitempty" name:"PackageId"`

	// Package type, 0 - outbound call package | 1 - 400 inbound call package.
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// <TOTAL_PACKAGE>.
	CapacitySize *int64 `json:"CapacitySize,omitnil,omitempty" name:"CapacitySize"`

	// Remaining package balance.
	CapacityRemain *int64 `json:"CapacityRemain,omitnil,omitempty" name:"CapacityRemain"`

	// Purchased timestamp.
	BuyTime *int64 `json:"BuyTime,omitnil,omitempty" name:"BuyTime"`

	// Deadline timestamp.
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

// Predefined struct for user
type PausePredictiveDialingCampaignRequestParams struct {
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Task id.
	CampaignId *int64 `json:"CampaignId,omitnil,omitempty" name:"CampaignId"`
}

type PausePredictiveDialingCampaignRequest struct {
	*tchttp.BaseRequest
	
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Task id.
	CampaignId *int64 `json:"CampaignId,omitnil,omitempty" name:"CampaignId"`
}

func (r *PausePredictiveDialingCampaignRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PausePredictiveDialingCampaignRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "CampaignId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PausePredictiveDialingCampaignRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PausePredictiveDialingCampaignResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type PausePredictiveDialingCampaignResponse struct {
	*tchttp.BaseResponse
	Response *PausePredictiveDialingCampaignResponseParams `json:"Response"`
}

func (r *PausePredictiveDialingCampaignResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PausePredictiveDialingCampaignResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PhoneNumBuyInfo struct {
	// Telephone number.
	PhoneNum *string `json:"PhoneNum,omitnil,omitempty" name:"PhoneNum"`

	// Number type, 0 - landline | 1 - virtual business number | 2 - ISP number | 3 - 400 number.
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// Call type of the number, 1 - inbound | 2 - outbound | 3 - inbound and outbound.
	CallType *int64 `json:"CallType,omitnil,omitempty" name:"CallType"`

	// Purchased timestamp.
	BuyTime *int64 `json:"BuyTime,omitnil,omitempty" name:"BuyTime"`

	// Deadline timestamp.
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Number status, 1-normal | 2-suspended due to non-payment | 4-admin suspended | 5-suspended due to violation.
	State *int64 `json:"State,omitnil,omitempty" name:"State"`
}

// Predefined struct for user
type ResetExtensionPasswordRequestParams struct {
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Extension.
	ExtensionId *string `json:"ExtensionId,omitnil,omitempty" name:"ExtensionId"`
}

type ResetExtensionPasswordRequest struct {
	*tchttp.BaseRequest
	
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Extension.
	ExtensionId *string `json:"ExtensionId,omitnil,omitempty" name:"ExtensionId"`
}

func (r *ResetExtensionPasswordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetExtensionPasswordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "ExtensionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetExtensionPasswordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetExtensionPasswordResponseParams struct {
	// Reset password.
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ResetExtensionPasswordResponse struct {
	*tchttp.BaseResponse
	Response *ResetExtensionPasswordResponseParams `json:"Response"`
}

func (r *ResetExtensionPasswordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetExtensionPasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RestoreMemberOnlineRequestParams struct {
	// App ID (required), which can be viewed at https://console.cloud.tencent.com/ccc.
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Customer service ID.
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

type RestoreMemberOnlineRequest struct {
	*tchttp.BaseRequest
	
	// App ID (required), which can be viewed at https://console.cloud.tencent.com/ccc.
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Customer service ID.
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

func (r *RestoreMemberOnlineRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestoreMemberOnlineRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "UserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RestoreMemberOnlineRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RestoreMemberOnlineResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RestoreMemberOnlineResponse struct {
	*tchttp.BaseResponse
	Response *RestoreMemberOnlineResponseParams `json:"Response"`
}

func (r *RestoreMemberOnlineResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestoreMemberOnlineResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResumePredictiveDialingCampaignRequestParams struct {
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// <Task id>.
	CampaignId *int64 `json:"CampaignId,omitnil,omitempty" name:"CampaignId"`
}

type ResumePredictiveDialingCampaignRequest struct {
	*tchttp.BaseRequest
	
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// <Task id>.
	CampaignId *int64 `json:"CampaignId,omitnil,omitempty" name:"CampaignId"`
}

func (r *ResumePredictiveDialingCampaignRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResumePredictiveDialingCampaignRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "CampaignId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResumePredictiveDialingCampaignRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResumePredictiveDialingCampaignResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ResumePredictiveDialingCampaignResponse struct {
	*tchttp.BaseResponse
	Response *ResumePredictiveDialingCampaignResponseParams `json:"Response"`
}

func (r *ResumePredictiveDialingCampaignResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResumePredictiveDialingCampaignResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SdkAppIdBuyInfo struct {
	// Application id.
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Application name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Agent purchase count (still within the validity period).
	StaffBuyNum *int64 `json:"StaffBuyNum,omitnil,omitempty" name:"StaffBuyNum"`

	// Agent purchase list (still within the validity period).
	StaffBuyList []*StaffBuyInfo `json:"StaffBuyList,omitnil,omitempty" name:"StaffBuyList"`

	// List of numbers purchased.
	PhoneNumBuyList []*PhoneNumBuyInfo `json:"PhoneNumBuyList,omitnil,omitempty" name:"PhoneNumBuyList"`

	// Number of office telephones purchased (still within the validity period).
	SipBuyNum *int64 `json:"SipBuyNum,omitnil,omitempty" name:"SipBuyNum"`
}

type SeatUserInfo struct {
	// Agent name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Agent email.
	Mail *string `json:"Mail,omitnil,omitempty" name:"Mail"`

	// Worker number.
	StaffNumber *string `json:"StaffNumber,omitnil,omitempty" name:"StaffNumber"`

	// Agent'S telephone number (with 0086 prefix).
	Phone *string `json:"Phone,omitnil,omitempty" name:"Phone"`

	// Agent nickname.
	Nick *string `json:"Nick,omitnil,omitempty" name:"Nick"`

	// User id.
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// List of skill groups associated with the agent.
	SkillGroupNameList []*string `json:"SkillGroupNameList,omitnil,omitempty" name:"SkillGroupNameList"`

	// 1: admin.
	// 2: quality inspector.
	// 3: ordinary agent.
	// Else: custom role id.
	Role *int64 `json:"Role,omitnil,omitempty" name:"Role"`

	// Agent extension number (starting with 1 to 8, 4 - 6 digits).
	ExtensionNumber *string `json:"ExtensionNumber,omitnil,omitempty" name:"ExtensionNumber"`
}

type ServeParticipant struct {
	// Agent email.
	Mail *string `json:"Mail,omitnil,omitempty" name:"Mail"`

	// Agent phone number.
	Phone *string `json:"Phone,omitnil,omitempty" name:"Phone"`

	// Ringing timestamp, unix second-level timestamp.
	RingTimestamp *int64 `json:"RingTimestamp,omitnil,omitempty" name:"RingTimestamp"`

	// Answer timestamp. unix second-level timestamp.
	AcceptTimestamp *int64 `json:"AcceptTimestamp,omitnil,omitempty" name:"AcceptTimestamp"`

	// End timestamp. unix second-level timestamp.
	EndedTimestamp *int64 `json:"EndedTimestamp,omitnil,omitempty" name:"EndedTimestamp"`

	// Recording id can be indexed to the agent side recording.
	RecordId *string `json:"RecordId,omitnil,omitempty" name:"RecordId"`

	// Participant type: "staffseat", "outboundseat", "staffphoneseat".
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Transfer source agent information.
	TransferFrom *string `json:"TransferFrom,omitnil,omitempty" name:"TransferFrom"`

	// Transfer source participant type is consistent with the type value.
	TransferFromType *string `json:"TransferFromType,omitnil,omitempty" name:"TransferFromType"`

	// Transfer destination agent information.
	TransferTo *string `json:"TransferTo,omitnil,omitempty" name:"TransferTo"`

	// Transfer destination participant type, which is consistent with type values.
	TransferToType *string `json:"TransferToType,omitnil,omitempty" name:"TransferToType"`

	// Skill group id.
	SkillGroupId *int64 `json:"SkillGroupId,omitnil,omitempty" name:"SkillGroupId"`

	// Ending status.
	EndStatusString *string `json:"EndStatusString,omitnil,omitempty" name:"EndStatusString"`

	// Recording url.
	RecordURL *string `json:"RecordURL,omitnil,omitempty" name:"RecordURL"`

	// Participant sequence number, starting from 0.
	Sequence *int64 `json:"Sequence,omitnil,omitempty" name:"Sequence"`

	// Start timestamp. unix second-level timestamp.
	StartTimestamp *int64 `json:"StartTimestamp,omitnil,omitempty" name:"StartTimestamp"`

	// Skill group name.
	SkillGroupName *string `json:"SkillGroupName,omitnil,omitempty" name:"SkillGroupName"`

	// Address of the third-party cos for transferring recording.
	CustomRecordURL *string `json:"CustomRecordURL,omitnil,omitempty" name:"CustomRecordURL"`
}

type ServerPushText struct {
	// Specifies the server push broadcast text.
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`

	// Whether to allow the text to interrupt the robot's speaking.
	Interrupt *bool `json:"Interrupt,omitnil,omitempty" name:"Interrupt"`

	// Specifies whether to automatically close the dialogue task after broadcasting the text.
	StopAfterPlay *bool `json:"StopAfterPlay,omitnil,omitempty" name:"StopAfterPlay"`

	// Specifies the server push broadcast audio.
	// Format description: audio must be mono, sampling rate should be consistent with the corresponding TTS, and coded as a Base64 string.
	// Input rules: when the Audio field is provided, the system will not accept input in the Text field. the system will play the Audio content in the Audio field directly.
	Audio *string `json:"Audio,omitnil,omitempty" name:"Audio"`

	// Defaults to 0. valid only when Interrupt is false.
	// -0 indicates that messages with Interrupt set to false will be dropped when there is an interaction.
	// -Indicates that when there is an interaction in progress, messages with Interrupt set to false will not be dropped but cached and processed after the current interaction is completed.
	// 
	// Note: when DropMode is 1, the cache allows multiple messages. if an interruption occurs subsequently, cached messages will be cleared.
	DropMode *uint64 `json:"DropMode,omitnil,omitempty" name:"DropMode"`

	// Message priority of ServerPushText. 0 means interruptible. 1 means not interruptible.
	// Note: after receiving a message with Priority=1, any other messages will be ignored (including messages with Priority=1) until the message processing of Priority=1 is complete. this field can be used together with the Interrupt and DropMode fields.
	// Example.
	// -Priority=1, Interrupt=true. specifies to Interrupt existing interaction and broadcast immediately. the broadcast will not be interrupted during the process.
	// -Priority=1, Interrupt=false, DropMode=1. waits for the current interaction to complete before broadcasting. the broadcast will not be interrupted during the process.
	Priority *uint64 `json:"Priority,omitnil,omitempty" name:"Priority"`

	// Specifies whether the text is added to the llm history context.
	AddHistory *bool `json:"AddHistory,omitnil,omitempty" name:"AddHistory"`
}

type SessionEvent struct {
	// Event timestamp. Unix second-level timestamp.
	Timestamp *int64 `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`

	// Event type. currently supports StaffHold, StaffUnhold, StaffMute, StaffUnmute.
	EventType *string `json:"EventType,omitnil,omitempty" name:"EventType"`

	// Describes event details related to the agent.
	StaffEventDetail *EventStaffDetail `json:"StaffEventDetail,omitnil,omitempty" name:"StaffEventDetail"`
}

type SkillGroupInfoItem struct {
	// Skill group id.
	SkillGroupId *int64 `json:"SkillGroupId,omitnil,omitempty" name:"SkillGroupId"`

	// Skill group name.
	SkillGroupName *string `json:"SkillGroupName,omitnil,omitempty" name:"SkillGroupName"`

	// (Deprecated) type: im, tel, all (full media).
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Session allocation policy.
	RoutePolicy *string `json:"RoutePolicy,omitnil,omitempty" name:"RoutePolicy"`

	// Whether the session is allocated to the last serviced agent first.
	UsingLastSeat *int64 `json:"UsingLastSeat,omitnil,omitempty" name:"UsingLastSeat"`

	// Maximum concurrency number of single client service (default 1 for telephone type).
	MaxConcurrency *int64 `json:"MaxConcurrency,omitnil,omitempty" name:"MaxConcurrency"`

	// Last modification time.
	LastModifyTimestamp *int64 `json:"LastModifyTimestamp,omitnil,omitempty" name:"LastModifyTimestamp"`

	// Skill group type 0-cell phone, 1-online, 3-audio, 4-video.	.	
	SkillGroupType *int64 `json:"SkillGroupType,omitnil,omitempty" name:"SkillGroupType"`

	// Intra-Skill group line number.
	Alias *string `json:"Alias,omitnil,omitempty" name:"Alias"`

	// Specifies whether to enable simultaneous ring.
	RingAll *bool `json:"RingAll,omitnil,omitempty" name:"RingAll"`
}

type SkillGroupItem struct {
	// Skill group id.
	SkillGroupId *int64 `json:"SkillGroupId,omitnil,omitempty" name:"SkillGroupId"`

	// Skill group name.
	SkillGroupName *string `json:"SkillGroupName,omitnil,omitempty" name:"SkillGroupName"`

	// Priority.
	Priority *int64 `json:"Priority,omitnil,omitempty" name:"Priority"`

	// Type: im, tel, all (full media).
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type StaffBuyInfo struct {
	// Number of agents purchased.
	Num *int64 `json:"Num,omitnil,omitempty" name:"Num"`

	// Purchased timestamp.
	BuyTime *int64 `json:"BuyTime,omitnil,omitempty" name:"BuyTime"`

	// Deadline timestamp.
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Quantity of office telephones purchased.
	SipNum *int64 `json:"SipNum,omitnil,omitempty" name:"SipNum"`
}

type StaffInfo struct {
	// Agent name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Agent email.
	Mail *string `json:"Mail,omitnil,omitempty" name:"Mail"`

	// Agent phone number.
	Phone *string `json:"Phone,omitnil,omitempty" name:"Phone"`

	// Agent nickname.
	Nick *string `json:"Nick,omitnil,omitempty" name:"Nick"`

	// Agent id.
	StaffNumber *string `json:"StaffNumber,omitnil,omitempty" name:"StaffNumber"`

	// User role ID. if a user is bound to multiple roles, RoleIdList takes precedence.
	//
	// Deprecated: RoleId is deprecated.
	RoleId *uint64 `json:"RoleId,omitnil,omitempty" name:"RoleId"`

	// User role id list.
	//
	// Deprecated: RoleIdList is deprecated.
	RoleIdList *uint64 `json:"RoleIdList,omitnil,omitempty" name:"RoleIdList"`

	// Specifies the user role id list.
	RoleList []*uint64 `json:"RoleList,omitnil,omitempty" name:"RoleList"`

	// Affiliated skill group list.
	SkillGroupList []*SkillGroupItem `json:"SkillGroupList,omitnil,omitempty" name:"SkillGroupList"`

	// Last modification time.
	LastModifyTimestamp *int64 `json:"LastModifyTimestamp,omitnil,omitempty" name:"LastModifyTimestamp"`

	// Agent extension number (starting with 1 to 8, 4 - 6 digits).
	ExtensionNumber *string `json:"ExtensionNumber,omitnil,omitempty" name:"ExtensionNumber"`

	// Call forwarding configuration.
	ForwardingConfig *ForwardingConfig `json:"ForwardingConfig,omitnil,omitempty" name:"ForwardingConfig"`
}

type StaffSkillGroupList struct {
	// Skill group id.
	SkillGroupId *int64 `json:"SkillGroupId,omitnil,omitempty" name:"SkillGroupId"`

	// Priority of the agent in the skill group (1 is the highest, 5 is the lowest, 3 by default).
	Priority *int64 `json:"Priority,omitnil,omitempty" name:"Priority"`
}

type StaffStatus struct {
	// Specifies the cursor used for querying in pagination scenarios.
	Cursor *string `json:"Cursor,omitnil,omitempty" name:"Cursor"`

	// Status timestamp. Unix second-level timestamp.
	Timestamp *int64 `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`

	// Agent status free | busy | rest | notReady | afterCallWork | offline.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Specifies the session Id for status association.
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`
}

type StaffStatusExtra struct {
	// IM - text | tel - cell phone | all - full media.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// IN - inbound | out - outbound.
	Direct *string `json:"Direct,omitnil,omitempty" name:"Direct"`
}

type StaffStatusMetrics struct {
	// Agent email.
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// Agent status free available | busy busy | rest on break | notready not ready | aftercallwork post-call adjustment | offline offline.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Supplementary information on agent status.
	StatusExtra *StaffStatusExtra `json:"StatusExtra,omitnil,omitempty" name:"StatusExtra"`

	// Total online duration of the day.
	OnlineDuration *int64 `json:"OnlineDuration,omitnil,omitempty" name:"OnlineDuration"`

	// Total available duration of the day.
	FreeDuration *int64 `json:"FreeDuration,omitnil,omitempty" name:"FreeDuration"`

	// Total busy duration of the day.
	BusyDuration *int64 `json:"BusyDuration,omitnil,omitempty" name:"BusyDuration"`

	// Total not ready status duration of the day.
	NotReadyDuration *int64 `json:"NotReadyDuration,omitnil,omitempty" name:"NotReadyDuration"`

	// Total break duration of the day.
	RestDuration *int64 `json:"RestDuration,omitnil,omitempty" name:"RestDuration"`

	// Adjust the total duration of after-call work for the day.
	AfterCallWorkDuration *int64 `json:"AfterCallWorkDuration,omitnil,omitempty" name:"AfterCallWorkDuration"`

	// Reason for break.
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`

	// Whether to reserve break status.
	ReserveRest *bool `json:"ReserveRest,omitnil,omitempty" name:"ReserveRest"`

	// Whether to reserve not ready status.
	ReserveNotReady *bool `json:"ReserveNotReady,omitnil,omitempty" name:"ReserveNotReady"`

	// Cell phone answering pattern: 0 - off | 1 - only when offline | 2 - always.
	UseMobileAccept *int64 `json:"UseMobileAccept,omitnil,omitempty" name:"UseMobileAccept"`

	// Cell phone outbound call switch.
	UseMobileCallOut *bool `json:"UseMobileCallOut,omitnil,omitempty" name:"UseMobileCallOut"`

	// Last online timestamp.
	// Note: this field may return null, indicating that no valid values can be obtained.
	LastOnlineTimestamp *int64 `json:"LastOnlineTimestamp,omitnil,omitempty" name:"LastOnlineTimestamp"`

	// Last status timestamp.
	// Note: this field may return null, indicating that no valid values can be obtained.
	LastStatusTimestamp *int64 `json:"LastStatusTimestamp,omitnil,omitempty" name:"LastStatusTimestamp"`

	// Specifies the endpoint information for customer service logon.
	ClientInfo []*ClientInfo `json:"ClientInfo,omitnil,omitempty" name:"ClientInfo"`
}

// Predefined struct for user
type StopAutoCalloutTaskRequestParams struct {
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Task id.
	TaskId *uint64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type StopAutoCalloutTaskRequest struct {
	*tchttp.BaseRequest
	
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Task id.
	TaskId *uint64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *StopAutoCalloutTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopAutoCalloutTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopAutoCalloutTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopAutoCalloutTaskResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StopAutoCalloutTaskResponse struct {
	*tchttp.BaseResponse
	Response *StopAutoCalloutTaskResponseParams `json:"Response"`
}

func (r *StopAutoCalloutTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopAutoCalloutTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TelCdrInfo struct {
	// Caller number.
	Caller *string `json:"Caller,omitnil,omitempty" name:"Caller"`

	// Called number.
	Callee *string `json:"Callee,omitnil,omitempty" name:"Callee"`

	// Call initiation timestamp, unix timestamp.
	Time *int64 `json:"Time,omitnil,omitempty" name:"Time"`

	// Call direction: 0 - inbound, 1 - outbound.
	Direction *int64 `json:"Direction,omitnil,omitempty" name:"Direction"`

	// Call Type: 1. Voice outbound call 2. Voice Inbound call 3. Audio Inbound 5 Predictive Dialing Call 6. Internal Call between Staff
	CallType *int64 `json:"CallType,omitnil,omitempty" name:"CallType"`

	// Call duration.
	Duration *int64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// Recording information.
	RecordURL *string `json:"RecordURL,omitnil,omitempty" name:"RecordURL"`

	// Recording id.
	RecordId *string `json:"RecordId,omitnil,omitempty" name:"RecordId"`

	// Agent information.
	SeatUser *SeatUserInfo `json:"SeatUser,omitnil,omitempty" name:"SeatUser"`

	// EndStatus corresponds one-to-one with EndStatusString. the specific enumeration is as follows:.
	// 
	// **Scenario         EndStatus         EndStatusString         status description**.
	// 
	// Incoming and outgoing calls.
	// 
	// Incoming and outgoing calls	0	error	exception termination.
	// 
	// Inbound call 102 ivrGiveUp user hang up during IVR.
	// 
	// Inbound call 103 waitingGiveUp user gives up while queuing.
	// 
	// Inbound call 104 ringingGiveUp user give up during ringing.
	// 
	// Inbound call 105 no agent online.
	// 
	// Inbound call 106 notWorkTime outside working hours.   
	// 
	// Inbound call 107 ivrEnd ends after IVR.
	// 
	// Inbound call 100 blackList. 
	// 
	// Outbound call 2 unconnected unanswered.
	// 
	// Outgoing call        108        restrictedCallee    the callee is restricted due to high risk.
	// 
	// Outbound call 109 too many requests outbound over-frequency limit.
	// 
	// Outgoing call             110	        restrictedArea	    valid values for the area where outgoing calls are restricted.
	// 
	// Outbound call 111 restrictedTime outgoing call time limit.
	//                          
	// Outbound call 201 unknown unknown status.
	// 
	// Outgoing call 202 not answered the callee did not answer.
	// 
	// Outgoing call            203	    userReject	callee rejects and hangs up.
	// 
	// Outbound call 204 powerOff callee is powered off.
	// 
	// Outbound call 205 number does not exist the callee's number is non - existent.
	// 
	// Outbound call 206 busy callee is busy.
	// 
	// Outbound call 207 arrears callee in arrears.
	// 
	// Outbound call 208 operator channel exception.
	// 
	// Outbound call 209 callerCancel call cancellation by the caller.
	// 
	// Outgoing call 210 notInService callee out of service area.
	// 
	// Incoming and outgoing calls 211 clientError client errors.
	// Outbound call 212 carrier blocked.
	EndStatus *int64 `json:"EndStatus,omitnil,omitempty" name:"EndStatus"`

	// Skill group name.
	SkillGroup *string `json:"SkillGroup,omitnil,omitempty" name:"SkillGroup"`

	// Caller'S location.
	CallerLocation *string `json:"CallerLocation,omitnil,omitempty" name:"CallerLocation"`

	// Time spent in ivr stage.
	IVRDuration *int64 `json:"IVRDuration,omitnil,omitempty" name:"IVRDuration"`

	// Ring timestamp. unix second-level timestamp.
	RingTimestamp *int64 `json:"RingTimestamp,omitnil,omitempty" name:"RingTimestamp"`

	// Answer timestamp. unix second-level timestamp.
	AcceptTimestamp *int64 `json:"AcceptTimestamp,omitnil,omitempty" name:"AcceptTimestamp"`

	// End timestamp. unix second-level timestamp.
	EndedTimestamp *int64 `json:"EndedTimestamp,omitnil,omitempty" name:"EndedTimestamp"`

	// IVR key information, e.g. ["1","2","3"].
	IVRKeyPressed []*string `json:"IVRKeyPressed,omitnil,omitempty" name:"IVRKeyPressed"`

	// Hang-Up side, seat, user, system.
	HungUpSide *string `json:"HungUpSide,omitnil,omitempty" name:"HungUpSide"`

	// Service participant list.
	ServeParticipants []*ServeParticipant `json:"ServeParticipants,omitnil,omitempty" name:"ServeParticipants"`

	// Skill group id.
	SkillGroupId *int64 `json:"SkillGroupId,omitnil,omitempty" name:"SkillGroupId"`

	// EndStatus corresponds one-to-one with EndStatusString. the specific enumeration is as follows:.
	// 
	// **Scenario         EndStatus         EndStatusString         status description**.
	// 
	// Incoming and outgoing calls.
	// 
	// Incoming and outgoing calls	0	error	exception termination.
	// 
	// Inbound call 102 ivrGiveUp user hang up during IVR.
	// 
	// Inbound call 103 waitingGiveUp user gives up while queuing.
	// 
	// Inbound call 104 ringingGiveUp user give up during ringing.
	// 
	// Inbound call 105 no agent online.
	// 
	// Inbound call 106 notWorkTime outside working hours.   
	// 
	// Inbound call 107 ivrEnd ends after IVR.
	// 
	// Inbound call 100 blackList. 
	// 
	// Outbound call 2 unconnected unanswered.
	// 
	// Outgoing call        108        restrictedCallee    the callee is restricted due to high risk.
	// 
	// Outbound call 109 too many requests outbound over-frequency limit.
	// 
	// Outgoing call             110	        restrictedArea	    valid values for the area where outgoing calls are restricted.
	// 
	// Outbound call 111 restrictedTime outgoing call time limit.
	//                          
	// Outgoing call 201 unknown unknown status.
	// 
	// Outgoing call 202 notAnswer callee not answered.
	// 
	// Outgoing call            203	    userReject	callee rejects and hangs up.
	// 
	// Outbound call 204 powerOff callee is powered off.
	// 
	// Outbound call 205 number does not exist the callee's number is non - existent.
	// 
	// Outgoing call 206 busy callee is busy.
	// 
	// Outbound call 207 out of credit callee in arrears.
	// 
	// Outbound call 208 operator channel exception.
	// 
	// Outgoing call 209 callerCancel call cancellation by the caller.
	// 
	// Outgoing call 210 notInService callee out of service area.
	// 
	// Incoming and outgoing calls 211 clientError client errors.
	// Outbound call 212 carrier blocked.
	EndStatusString *string `json:"EndStatusString,omitnil,omitempty" name:"EndStatusString"`

	// Session start timestamp. unix second-level timestamp.
	StartTimestamp *int64 `json:"StartTimestamp,omitnil,omitempty" name:"StartTimestamp"`

	// Queue entry time. unix second-level timestamp.
	QueuedTimestamp *int64 `json:"QueuedTimestamp,omitnil,omitempty" name:"QueuedTimestamp"`

	// Post-IVR key information (e.g. [{"key":"1","label":"very satisfied"}]).
	PostIVRKeyPressed []*IVRKeyPressedElement `json:"PostIVRKeyPressed,omitnil,omitempty" name:"PostIVRKeyPressed"`

	// Queue skill group id.
	QueuedSkillGroupId *int64 `json:"QueuedSkillGroupId,omitnil,omitempty" name:"QueuedSkillGroupId"`

	// Session id.
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// Caller number protection id. effective when the number protection map feature is activated, and the caller field is empty.
	ProtectedCaller *string `json:"ProtectedCaller,omitnil,omitempty" name:"ProtectedCaller"`

	// Called number protection id. effective when the number protection map feature is activated, and the callee field is empty.
	ProtectedCallee *string `json:"ProtectedCallee,omitnil,omitempty" name:"ProtectedCallee"`

	// Customer custom data. (user - to - user interface).
	// Note: this field may return null, indicating that no valid values can be obtained.
	//
	// Deprecated: Uui is deprecated.
	Uui *string `json:"Uui,omitnil,omitempty" name:"Uui"`

	// Customer custom data. (user - to - user interface).
	UUI *string `json:"UUI,omitnil,omitempty" name:"UUI"`

	// IVR key information (e.g. [{"key":"1","label":"very satisfied"}]).
	IVRKeyPressedEx []*IVRKeyPressedElement `json:"IVRKeyPressedEx,omitnil,omitempty" name:"IVRKeyPressedEx"`

	// Access to the asr text information address of the recording.
	AsrUrl *string `json:"AsrUrl,omitnil,omitempty" name:"AsrUrl"`

	// ASRUrl status: complete.
	// Completed;.
	// Processing.
	// Generating.
	// NotExists.
	// No record (offline asr generation is not enabled or no package is available).
	AsrStatus *string `json:"AsrStatus,omitnil,omitempty" name:"AsrStatus"`

	// Address of the third-party cos for transferring recording.
	CustomRecordURL *string `json:"CustomRecordURL,omitnil,omitempty" name:"CustomRecordURL"`

	// Remarks.
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// Queue skill group name.
	QueuedSkillGroupName *string `json:"QueuedSkillGroupName,omitnil,omitempty" name:"QueuedSkillGroupName"`

	// Audio message recording url during call.
	VoicemailRecordURL []*string `json:"VoicemailRecordURL,omitnil,omitempty" name:"VoicemailRecordURL"`

	// Text information address of asr audio message during a call.
	VoicemailAsrURL []*string `json:"VoicemailAsrURL,omitnil,omitempty" name:"VoicemailAsrURL"`
}

type TimeRange struct {
	// Start time.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type ToneWordInfo struct {
	// Specifies the timeout period for the first request in seconds.
	FirstSentenceTimeout *float64 `json:"FirstSentenceTimeout,omitnil,omitempty" name:"FirstSentenceTimeout"`

	// Undertakes a modal particle.
	ZHToneWords *ZHToneWordsInfo `json:"ZHToneWords,omitnil,omitempty" name:"ZHToneWords"`
}

// Predefined struct for user
type TransferToManualRequestParams struct {
	// App ID (required), which can be viewed at https://console.cloud.tencent.com/ccc.
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Session ID.
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// Skill group Id.
	SkillGroupId *uint64 `json:"SkillGroupId,omitnil,omitempty" name:"SkillGroupId"`
}

type TransferToManualRequest struct {
	*tchttp.BaseRequest
	
	// App ID (required), which can be viewed at https://console.cloud.tencent.com/ccc.
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Session ID.
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// Skill group Id.
	SkillGroupId *uint64 `json:"SkillGroupId,omitnil,omitempty" name:"SkillGroupId"`
}

func (r *TransferToManualRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TransferToManualRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "SessionId")
	delete(f, "SkillGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TransferToManualRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TransferToManualResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type TransferToManualResponse struct {
	*tchttp.BaseResponse
	Response *TransferToManualResponseParams `json:"Response"`
}

func (r *TransferToManualResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TransferToManualResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnbindNumberCallOutSkillGroupRequestParams struct {
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Number to be unbound.
	Number *string `json:"Number,omitnil,omitempty" name:"Number"`

	// List of skill group ids to be unbound.
	SkillGroupIds []*uint64 `json:"SkillGroupIds,omitnil,omitempty" name:"SkillGroupIds"`
}

type UnbindNumberCallOutSkillGroupRequest struct {
	*tchttp.BaseRequest
	
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Number to be unbound.
	Number *string `json:"Number,omitnil,omitempty" name:"Number"`

	// List of skill group ids to be unbound.
	SkillGroupIds []*uint64 `json:"SkillGroupIds,omitnil,omitempty" name:"SkillGroupIds"`
}

func (r *UnbindNumberCallOutSkillGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindNumberCallOutSkillGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "Number")
	delete(f, "SkillGroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnbindNumberCallOutSkillGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnbindNumberCallOutSkillGroupResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UnbindNumberCallOutSkillGroupResponse struct {
	*tchttp.BaseResponse
	Response *UnbindNumberCallOutSkillGroupResponseParams `json:"Response"`
}

func (r *UnbindNumberCallOutSkillGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindNumberCallOutSkillGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnbindStaffSkillGroupListRequestParams struct {
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Customer service email.
	StaffEmail *string `json:"StaffEmail,omitnil,omitempty" name:"StaffEmail"`

	// Unbound skill group list.
	SkillGroupList []*int64 `json:"SkillGroupList,omitnil,omitempty" name:"SkillGroupList"`
}

type UnbindStaffSkillGroupListRequest struct {
	*tchttp.BaseRequest
	
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Customer service email.
	StaffEmail *string `json:"StaffEmail,omitnil,omitempty" name:"StaffEmail"`

	// Unbound skill group list.
	SkillGroupList []*int64 `json:"SkillGroupList,omitnil,omitempty" name:"SkillGroupList"`
}

func (r *UnbindStaffSkillGroupListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindStaffSkillGroupListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "StaffEmail")
	delete(f, "SkillGroupList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnbindStaffSkillGroupListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnbindStaffSkillGroupListResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UnbindStaffSkillGroupListResponse struct {
	*tchttp.BaseResponse
	Response *UnbindStaffSkillGroupListResponseParams `json:"Response"`
}

func (r *UnbindStaffSkillGroupListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindStaffSkillGroupListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateCCCSkillGroupRequestParams struct {
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Skill group id.
	SkillGroupID *int64 `json:"SkillGroupID,omitnil,omitempty" name:"SkillGroupID"`

	// Modified skill group name.
	SkillGroupName *string `json:"SkillGroupName,omitnil,omitempty" name:"SkillGroupName"`

	// Modified maximum concurrency, with the maximum synchronization being 2.
	MaxConcurrency *int64 `json:"MaxConcurrency,omitnil,omitempty" name:"MaxConcurrency"`

	// True for simultaneous ringing, false for sequential ringing.
	RingAll *bool `json:"RingAll,omitnil,omitempty" name:"RingAll"`
}

type UpdateCCCSkillGroupRequest struct {
	*tchttp.BaseRequest
	
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Skill group id.
	SkillGroupID *int64 `json:"SkillGroupID,omitnil,omitempty" name:"SkillGroupID"`

	// Modified skill group name.
	SkillGroupName *string `json:"SkillGroupName,omitnil,omitempty" name:"SkillGroupName"`

	// Modified maximum concurrency, with the maximum synchronization being 2.
	MaxConcurrency *int64 `json:"MaxConcurrency,omitnil,omitempty" name:"MaxConcurrency"`

	// True for simultaneous ringing, false for sequential ringing.
	RingAll *bool `json:"RingAll,omitnil,omitempty" name:"RingAll"`
}

func (r *UpdateCCCSkillGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateCCCSkillGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "SkillGroupID")
	delete(f, "SkillGroupName")
	delete(f, "MaxConcurrency")
	delete(f, "RingAll")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateCCCSkillGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateCCCSkillGroupResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateCCCSkillGroupResponse struct {
	*tchttp.BaseResponse
	Response *UpdateCCCSkillGroupResponseParams `json:"Response"`
}

func (r *UpdateCCCSkillGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateCCCSkillGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdatePredictiveDialingCampaignRequestParams struct {
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Generated task id.
	CampaignId *int64 `json:"CampaignId,omitnil,omitempty" name:"CampaignId"`

	// Task name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Called list supporting e.164 or number formats without country code.
	Callees []*string `json:"Callees,omitnil,omitempty" name:"Callees"`

	// Calling list using the number formats displayed on the management side.
	Callers []*string `json:"Callers,omitnil,omitempty" name:"Callers"`

	// Being called sequence: 0 for random 1 for in order.
	CallOrder *int64 `json:"CallOrder,omitnil,omitempty" name:"CallOrder"`

	// ID of the used skill group of agents.
	SkillGroupId *int64 `json:"SkillGroupId,omitnil,omitempty" name:"SkillGroupId"`

	// Running priority of multiple tasks in the same application, from high to low 1 - 5.
	Priority *int64 `json:"Priority,omitnil,omitempty" name:"Priority"`

	// Expected call drop rate, percentage, 5 - 50.	.	
	ExpectedAbandonRate *int64 `json:"ExpectedAbandonRate,omitnil,omitempty" name:"ExpectedAbandonRate"`

	// Call retry interval, in seconds, [60 - 86,400].
	RetryInterval *int64 `json:"RetryInterval,omitnil,omitempty" name:"RetryInterval"`

	// Task start time. unix timestamp. the task will automatically start after this time.
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Task termination time. unix timestamp. the task will automatically terminate after this time.
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Specified ivr id.
	IVRId *int64 `json:"IVRId,omitnil,omitempty" name:"IVRId"`

	// Number of call retries, 0 - 2.
	RetryTimes *int64 `json:"RetryTimes,omitnil,omitempty" name:"RetryTimes"`

	// Custom variable.
	Variables []*Variable `json:"Variables,omitnil,omitempty" name:"Variables"`

	// 	UUI
	UUI *string `json:"UUI,omitnil,omitempty" name:"UUI"`

	// Property of the called.
	CalleeAttributes []*CalleeAttribute `json:"CalleeAttributes,omitnil,omitempty" name:"CalleeAttributes"`
}

type UpdatePredictiveDialingCampaignRequest struct {
	*tchttp.BaseRequest
	
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Generated task id.
	CampaignId *int64 `json:"CampaignId,omitnil,omitempty" name:"CampaignId"`

	// Task name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Called list supporting e.164 or number formats without country code.
	Callees []*string `json:"Callees,omitnil,omitempty" name:"Callees"`

	// Calling list using the number formats displayed on the management side.
	Callers []*string `json:"Callers,omitnil,omitempty" name:"Callers"`

	// Being called sequence: 0 for random 1 for in order.
	CallOrder *int64 `json:"CallOrder,omitnil,omitempty" name:"CallOrder"`

	// ID of the used skill group of agents.
	SkillGroupId *int64 `json:"SkillGroupId,omitnil,omitempty" name:"SkillGroupId"`

	// Running priority of multiple tasks in the same application, from high to low 1 - 5.
	Priority *int64 `json:"Priority,omitnil,omitempty" name:"Priority"`

	// Expected call drop rate, percentage, 5 - 50.	.	
	ExpectedAbandonRate *int64 `json:"ExpectedAbandonRate,omitnil,omitempty" name:"ExpectedAbandonRate"`

	// Call retry interval, in seconds, [60 - 86,400].
	RetryInterval *int64 `json:"RetryInterval,omitnil,omitempty" name:"RetryInterval"`

	// Task start time. unix timestamp. the task will automatically start after this time.
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Task termination time. unix timestamp. the task will automatically terminate after this time.
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Specified ivr id.
	IVRId *int64 `json:"IVRId,omitnil,omitempty" name:"IVRId"`

	// Number of call retries, 0 - 2.
	RetryTimes *int64 `json:"RetryTimes,omitnil,omitempty" name:"RetryTimes"`

	// Custom variable.
	Variables []*Variable `json:"Variables,omitnil,omitempty" name:"Variables"`

	// 	UUI
	UUI *string `json:"UUI,omitnil,omitempty" name:"UUI"`

	// Property of the called.
	CalleeAttributes []*CalleeAttribute `json:"CalleeAttributes,omitnil,omitempty" name:"CalleeAttributes"`
}

func (r *UpdatePredictiveDialingCampaignRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdatePredictiveDialingCampaignRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "CampaignId")
	delete(f, "Name")
	delete(f, "Callees")
	delete(f, "Callers")
	delete(f, "CallOrder")
	delete(f, "SkillGroupId")
	delete(f, "Priority")
	delete(f, "ExpectedAbandonRate")
	delete(f, "RetryInterval")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "IVRId")
	delete(f, "RetryTimes")
	delete(f, "Variables")
	delete(f, "UUI")
	delete(f, "CalleeAttributes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdatePredictiveDialingCampaignRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdatePredictiveDialingCampaignResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdatePredictiveDialingCampaignResponse struct {
	*tchttp.BaseResponse
	Response *UpdatePredictiveDialingCampaignResponseParams `json:"Response"`
}

func (r *UpdatePredictiveDialingCampaignResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdatePredictiveDialingCampaignResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UploadAudioInfo struct {
	// File alias (can be duplicated).
	CustomFileName *string `json:"CustomFileName,omitnil,omitempty" name:"CustomFileName"`

	// Audio file link (supports mp3 and wav formats, file size not exceeding 5mb).
	AudioUrl *string `json:"AudioUrl,omitnil,omitempty" name:"AudioUrl"`
}

type UploadIvrAudioFailedInfo struct {
	// Filename.
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// Reason for failure.
	FailedMsg *string `json:"FailedMsg,omitnil,omitempty" name:"FailedMsg"`
}

// Predefined struct for user
type UploadIvrAudioRequestParams struct {
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Audio file list.
	AudioList []*UploadAudioInfo `json:"AudioList,omitnil,omitempty" name:"AudioList"`
}

type UploadIvrAudioRequest struct {
	*tchttp.BaseRequest
	
	// Application id (required) can be found at https://console.cloud.tencent.com/ccc.
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// Audio file list.
	AudioList []*UploadAudioInfo `json:"AudioList,omitnil,omitempty" name:"AudioList"`
}

func (r *UploadIvrAudioRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadIvrAudioRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "AudioList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UploadIvrAudioRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UploadIvrAudioResponseParams struct {
	// List of files that failed to be uploaded.
	FailedFileList []*UploadIvrAudioFailedInfo `json:"FailedFileList,omitnil,omitempty" name:"FailedFileList"`

	// List of successfully uploaded files.
	SuccessFileList []*AudioFileInfo `json:"SuccessFileList,omitnil,omitempty" name:"SuccessFileList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UploadIvrAudioResponse struct {
	*tchttp.BaseResponse
	Response *UploadIvrAudioResponseParams `json:"Response"`
}

func (r *UploadIvrAudioResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadIvrAudioResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Variable struct {
	// Variable name.
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// Variable value.
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type ZHToneWordsInfo struct {
	// Specifies the word list.
	Neutral []*string `json:"Neutral,omitnil,omitempty" name:"Neutral"`

	// Positive word list.
	Positive []*string `json:"Positive,omitnil,omitempty" name:"Positive"`

	// Negative word list.
	Negative []*string `json:"Negative,omitnil,omitempty" name:"Negative"`
}