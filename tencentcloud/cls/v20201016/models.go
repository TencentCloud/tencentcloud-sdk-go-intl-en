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

package v20201016

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

type AlarmInfo struct {

	// Alarm policy name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Monitoring object list
	AlarmTargets []*AlarmTargetInfo `json:"AlarmTargets,omitempty" name:"AlarmTargets"`

	// Monitoring task running time point
	MonitorTime *MonitorTime `json:"MonitorTime,omitempty" name:"MonitorTime"`

	// Trigger condition
	Condition *string `json:"Condition,omitempty" name:"Condition"`

	// Alarm persistence cycle. An alarm will be triggered only after the corresponding trigger condition is met for the number of times specified by `TriggerCount`. Value range: 1–10.
	TriggerCount *int64 `json:"TriggerCount,omitempty" name:"TriggerCount"`

	// Repeated alarm interval in minutes. Value range: 0–1440.
	AlarmPeriod *int64 `json:"AlarmPeriod,omitempty" name:"AlarmPeriod"`

	// List of associated alarm notification templates
	AlarmNoticeIds []*string `json:"AlarmNoticeIds,omitempty" name:"AlarmNoticeIds"`

	// Enablement status
	Status *bool `json:"Status,omitempty" name:"Status"`

	// Alarm policy ID
	AlarmId *string `json:"AlarmId,omitempty" name:"AlarmId"`

	// Creation time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Last update time
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// Custom notification template
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	MessageTemplate *string `json:"MessageTemplate,omitempty" name:"MessageTemplate"`

	// Custom callback template
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	CallBack *CallBackInfo `json:"CallBack,omitempty" name:"CallBack"`

	// Multi-Dimensional analysis settings
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Analysis []*AnalysisDimensional `json:"Analysis,omitempty" name:"Analysis"`
}

type AlarmNotice struct {

	// Alarm notification template name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Alarm template type. Valid values:
	// <br><li> `Trigger`: alarm triggered
	// <br><li> `Recovery`: alarm cleared
	// <br><li> `All`: alarm triggered and alarm cleared
	Type *string `json:"Type,omitempty" name:"Type"`

	// Information of the recipient in alarm notification template
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	NoticeReceivers []*NoticeReceiver `json:"NoticeReceivers,omitempty" name:"NoticeReceivers"`

	// Callback information of alarm notification template
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	WebCallbacks []*WebCallback `json:"WebCallbacks,omitempty" name:"WebCallbacks"`

	// Alarm notification template ID
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	AlarmNoticeId *string `json:"AlarmNoticeId,omitempty" name:"AlarmNoticeId"`

	// Creation time
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Last update time
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type AlarmTarget struct {

	// Log topic ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// Query statement
	Query *string `json:"Query,omitempty" name:"Query"`

	// Monitoring object number, which is incremental from 1.
	Number *int64 `json:"Number,omitempty" name:"Number"`

	// Offset of the query start time from the current time in minutes. The value cannot be positive. Value range: -1440–0.
	StartTimeOffset *int64 `json:"StartTimeOffset,omitempty" name:"StartTimeOffset"`

	// Offset of the query end time from the current time in minutes. The value cannot be positive and must be greater than `StartTimeOffset`. Value range: -1440–0.
	EndTimeOffset *int64 `json:"EndTimeOffset,omitempty" name:"EndTimeOffset"`

	// Logset ID
	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`
}

type AlarmTargetInfo struct {

	// Logset ID
	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`

	// Logset name
	LogsetName *string `json:"LogsetName,omitempty" name:"LogsetName"`

	// Log topic ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// Log topic name
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// Query statement
	Query *string `json:"Query,omitempty" name:"Query"`

	// Monitoring object number
	Number *int64 `json:"Number,omitempty" name:"Number"`

	// Offset of the query start time from the current time. The value cannot be positive. Value range: -1440–0.
	StartTimeOffset *int64 `json:"StartTimeOffset,omitempty" name:"StartTimeOffset"`

	// Offset of the query end time from the current time. The value cannot be positive and must be greater than `StartTimeOffset`. Value range: -1440–0.
	EndTimeOffset *int64 `json:"EndTimeOffset,omitempty" name:"EndTimeOffset"`
}

type AnalysisDimensional struct {

	// Analysis name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Analysis type. Valid values: `query`, `field`
	Type *string `json:"Type,omitempty" name:"Type"`

	// Analysis content
	Content *string `json:"Content,omitempty" name:"Content"`
}

type ApplyConfigToMachineGroupRequest struct {
	*tchttp.BaseRequest

	// Collection configuration ID
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`

	// Machine group ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *ApplyConfigToMachineGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyConfigToMachineGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ConfigId")
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ApplyConfigToMachineGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ApplyConfigToMachineGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ApplyConfigToMachineGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyConfigToMachineGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AsyncContextTask struct {

	// Logset ID
	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`

	// Log topic ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// Creation time, which is a timestamp accurate down to the millisecond
	CreateTime *int64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// Status. Valid values: `0`: to be started; `1`: running; `2`: completed; `-1`: failed
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// Offline context search task ID
	AsyncContextTaskId *string `json:"AsyncContextTaskId,omitempty" name:"AsyncContextTaskId"`

	// Error message of task failure
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ErrorMessage *string `json:"ErrorMessage,omitempty" name:"ErrorMessage"`

	// Log package number
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	PkgId *string `json:"PkgId,omitempty" name:"PkgId"`

	// Log number in log package
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	PkgLogId *string `json:"PkgLogId,omitempty" name:"PkgLogId"`

	// Log time
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Time *int64 `json:"Time,omitempty" name:"Time"`

	// Task completion time, which is a timestamp accurate down to the millisecond
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	FinishTime *int64 `json:"FinishTime,omitempty" name:"FinishTime"`

	// Associated offline search ID
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	AsyncSearchTaskId *string `json:"AsyncSearchTaskId,omitempty" name:"AsyncSearchTaskId"`
}

type AsyncSearchTask struct {

	// Logset ID
	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`

	// Log topic ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// Creation time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Status. Valid values: `0`: to be started; `1`: running; `2`: completed; `-1`: failed
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// Offline search task ID
	AsyncSearchTaskId *string `json:"AsyncSearchTaskId,omitempty" name:"AsyncSearchTaskId"`

	// Query statement
	Query *string `json:"Query,omitempty" name:"Query"`

	// Start time of the log to be queried, which is a Unix timestamp in milliseconds
	From *int64 `json:"From,omitempty" name:"From"`

	// End time of the log to be queried, which is a Unix timestamp in milliseconds
	To *int64 `json:"To,omitempty" name:"To"`

	// Log scan order. Valid values: `asc`: ascending; `desc`: descending
	Sort *string `json:"Sort,omitempty" name:"Sort"`

	// Error message of task failure
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ErrorMessage *string `json:"ErrorMessage,omitempty" name:"ErrorMessage"`

	// Total number of logs matched in offline search task
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	LogCount *int64 `json:"LogCount,omitempty" name:"LogCount"`

	// Task completion time
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	FinishTime *string `json:"FinishTime,omitempty" name:"FinishTime"`
}

type CallBackInfo struct {

	// `Body` during callback
	Body *string `json:"Body,omitempty" name:"Body"`

	// `Headers` during callback
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Headers []*string `json:"Headers,omitempty" name:"Headers"`
}

type Column struct {

	// Column name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Column attribute
	Type *string `json:"Type,omitempty" name:"Type"`
}

type CompressInfo struct {

	// Compression format. Valid values: `gzip`, `lzop`, `none` (no compression)
	Format *string `json:"Format,omitempty" name:"Format"`
}

type ConfigInfo struct {

	// Collection rule configuration ID
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`

	// Log formatting method
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	LogFormat *string `json:"LogFormat,omitempty" name:"LogFormat"`

	// Log collection path
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Path *string `json:"Path,omitempty" name:"Path"`

	// Type of the log to be collected. Valid values: `json_log`: log in JSON format; `delimiter_log`: log in delimited format; `minimalist_log`: minimalist log; `multiline_log`: log in multi-line format; `fullregex_log`: log in full regex format. Default value: `minimalist_log`
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	LogType *string `json:"LogType,omitempty" name:"LogType"`

	// Extraction rule. If `ExtractRule` is set, `LogType` must be set
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ExtractRule *ExtractRuleInfo `json:"ExtractRule,omitempty" name:"ExtractRule"`

	// Collection path blocklist
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ExcludePaths []*ExcludePathInfo `json:"ExcludePaths,omitempty" name:"ExcludePaths"`

	// Log topic ID (TopicId) of collection configuration
	Output *string `json:"Output,omitempty" name:"Output"`

	// Update time
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// Creation time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Custom parsing string
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	UserDefineRule *string `json:"UserDefineRule,omitempty" name:"UserDefineRule"`
}

type ContentInfo struct {

	// Content format. Valid values: `json`, `csv`
	Format *string `json:"Format,omitempty" name:"Format"`

	// CSV format content description
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Csv *CsvInfo `json:"Csv,omitempty" name:"Csv"`

	// JSON format content description
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Json *JsonInfo `json:"Json,omitempty" name:"Json"`
}

type CreateAlarmNoticeRequest struct {
	*tchttp.BaseRequest

	// Alarm template name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Alarm template type. Valid values:
	// <br><li> `Trigger`: alarm triggered
	// <br><li> `Recovery`: alarm cleared
	// <br><li> `All`: alarm triggered and alarm cleared
	Type *string `json:"Type,omitempty" name:"Type"`

	// Information of the recipient in alarm template
	NoticeReceivers []*NoticeReceiver `json:"NoticeReceivers,omitempty" name:"NoticeReceivers"`

	// Alarm template callback information
	WebCallbacks []*WebCallback `json:"WebCallbacks,omitempty" name:"WebCallbacks"`
}

func (r *CreateAlarmNoticeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAlarmNoticeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Type")
	delete(f, "NoticeReceivers")
	delete(f, "WebCallbacks")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAlarmNoticeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateAlarmNoticeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Alarm template ID
		AlarmNoticeId *string `json:"AlarmNoticeId,omitempty" name:"AlarmNoticeId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAlarmNoticeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAlarmNoticeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAlarmRequest struct {
	*tchttp.BaseRequest

	// Alarm policy name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Monitoring object list
	AlarmTargets []*AlarmTarget `json:"AlarmTargets,omitempty" name:"AlarmTargets"`

	// Monitoring task running time point
	MonitorTime *MonitorTime `json:"MonitorTime,omitempty" name:"MonitorTime"`

	// Trigger condition
	Condition *string `json:"Condition,omitempty" name:"Condition"`

	// Alarm persistence cycle. An alarm will be triggered only after the corresponding trigger condition is met for the number of times specified by `TriggerCount`. Value range: 1–10.
	TriggerCount *int64 `json:"TriggerCount,omitempty" name:"TriggerCount"`

	// Repeated alarm interval in minutes. Value range: 0–1440.
	AlarmPeriod *int64 `json:"AlarmPeriod,omitempty" name:"AlarmPeriod"`

	// List of associated alarm notification templates
	AlarmNoticeIds []*string `json:"AlarmNoticeIds,omitempty" name:"AlarmNoticeIds"`

	// Whether to enable the alarm policy. Default value: true
	Status *bool `json:"Status,omitempty" name:"Status"`

	// Custom alarm content
	MessageTemplate *string `json:"MessageTemplate,omitempty" name:"MessageTemplate"`

	// Custom callback
	CallBack *CallBackInfo `json:"CallBack,omitempty" name:"CallBack"`

	// Multi-Dimensional analysis
	Analysis []*AnalysisDimensional `json:"Analysis,omitempty" name:"Analysis"`
}

func (r *CreateAlarmRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAlarmRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "AlarmTargets")
	delete(f, "MonitorTime")
	delete(f, "Condition")
	delete(f, "TriggerCount")
	delete(f, "AlarmPeriod")
	delete(f, "AlarmNoticeIds")
	delete(f, "Status")
	delete(f, "MessageTemplate")
	delete(f, "CallBack")
	delete(f, "Analysis")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAlarmRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateAlarmResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Alarm policy ID.
		AlarmId *string `json:"AlarmId,omitempty" name:"AlarmId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAlarmResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAlarmResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAsyncContextTaskRequest struct {
	*tchttp.BaseRequest

	// Log topic ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// Log time in milliseconds
	Time *int64 `json:"Time,omitempty" name:"Time"`

	// Log package number
	PkgId *string `json:"PkgId,omitempty" name:"PkgId"`

	// Log number in log package
	PkgLogId *string `json:"PkgLogId,omitempty" name:"PkgLogId"`

	// Logset ID
	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`

	// Offline search task ID
	AsyncSearchTaskId *string `json:"AsyncSearchTaskId,omitempty" name:"AsyncSearchTaskId"`
}

func (r *CreateAsyncContextTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAsyncContextTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	delete(f, "Time")
	delete(f, "PkgId")
	delete(f, "PkgLogId")
	delete(f, "LogsetId")
	delete(f, "AsyncSearchTaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAsyncContextTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateAsyncContextTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Offline context search task ID
		AsyncContextTaskId *string `json:"AsyncContextTaskId,omitempty" name:"AsyncContextTaskId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAsyncContextTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAsyncContextTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAsyncSearchTaskRequest struct {
	*tchttp.BaseRequest

	// Logset ID
	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`

	// Log topic ID. Currently, only log topics whose `StorageType` is `cold` are supported.
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// Query statement. Maximum length: 1024
	Query *string `json:"Query,omitempty" name:"Query"`

	// Start time of the log to be queried, which is a Unix timestamp in milliseconds
	From *int64 `json:"From,omitempty" name:"From"`

	// End time of the log to be queried, which is a Unix timestamp in milliseconds
	To *int64 `json:"To,omitempty" name:"To"`

	// Log scan order. Valid values: `asc`: ascending; `desc`: descending. Default value: desc
	Sort *string `json:"Sort,omitempty" name:"Sort"`
}

func (r *CreateAsyncSearchTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAsyncSearchTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LogsetId")
	delete(f, "TopicId")
	delete(f, "Query")
	delete(f, "From")
	delete(f, "To")
	delete(f, "Sort")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAsyncSearchTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateAsyncSearchTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAsyncSearchTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAsyncSearchTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateConfigRequest struct {
	*tchttp.BaseRequest

	// Collection configuration name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Log topic ID (TopicId) of collection configuration
	Output *string `json:"Output,omitempty" name:"Output"`

	// Log collection path containing the filename
	Path *string `json:"Path,omitempty" name:"Path"`

	// Type of the log to be collected. Valid values: `json_log`: log in JSON format; `delimiter_log`: log in delimited format; `minimalist_log`: minimalist log; `multiline_log`: log in multi-line format; `fullregex_log`: log in full regex format. Default value: `minimalist_log`
	LogType *string `json:"LogType,omitempty" name:"LogType"`

	// Extraction rule. If `ExtractRule` is set, `LogType` must be set.
	ExtractRule *ExtractRuleInfo `json:"ExtractRule,omitempty" name:"ExtractRule"`

	// Collection path blocklist
	ExcludePaths []*ExcludePathInfo `json:"ExcludePaths,omitempty" name:"ExcludePaths"`

	// Custom collection rule, which is a serialized JSON string
	UserDefineRule *string `json:"UserDefineRule,omitempty" name:"UserDefineRule"`
}

func (r *CreateConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Output")
	delete(f, "Path")
	delete(f, "LogType")
	delete(f, "ExtractRule")
	delete(f, "ExcludePaths")
	delete(f, "UserDefineRule")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Collection configuration ID
		ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateExportRequest struct {
	*tchttp.BaseRequest

	// Log topic
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// Log export search statement
	Query *string `json:"Query,omitempty" name:"Query"`

	// Number of logs to be exported. Maximum value: 10 million
	Count *uint64 `json:"Count,omitempty" name:"Count"`

	// Start time of the log to be exported, which is a timestamp in milliseconds
	From *int64 `json:"From,omitempty" name:"From"`

	// End time of the log to be exported, which is a timestamp in milliseconds
	To *int64 `json:"To,omitempty" name:"To"`

	// Exported log sorting order by time. Valid values: `asc`: ascending; `desc`: descending. Default value: `desc`
	Order *string `json:"Order,omitempty" name:"Order"`

	// Exported log data format. Valid values: `json`, `csv`. Default value: `json`
	Format *string `json:"Format,omitempty" name:"Format"`
}

func (r *CreateExportRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateExportRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	delete(f, "Query")
	delete(f, "Count")
	delete(f, "From")
	delete(f, "To")
	delete(f, "Order")
	delete(f, "Format")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateExportRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateExportResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Log export ID.
		ExportId *string `json:"ExportId,omitempty" name:"ExportId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateExportResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateExportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateIndexRequest struct {
	*tchttp.BaseRequest

	// Log topic ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// Index rule
	Rule *RuleInfo `json:"Rule,omitempty" name:"Rule"`

	// Whether to take effect. Default value: true
	Status *bool `json:"Status,omitempty" name:"Status"`
}

func (r *CreateIndexRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIndexRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	delete(f, "Rule")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateIndexRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateIndexResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateIndexResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIndexResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateLogsetRequest struct {
	*tchttp.BaseRequest

	// Logset name, which must be unique
	LogsetName *string `json:"LogsetName,omitempty" name:"LogsetName"`

	// Tag description list. Up to 10 tag key-value pairs are supported and must be unique.
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

func (r *CreateLogsetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLogsetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LogsetName")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateLogsetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateLogsetResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Logset ID
		LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateLogsetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLogsetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateMachineGroupRequest struct {
	*tchttp.BaseRequest

	// Machine group name, which must be unique
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// Type of the machine group to be created. Valid values: `ip`: use the IP string list in `Values` to create a machine group; `label`: use the tag string list in `Values` to create a machine group
	MachineGroupType *MachineGroupTypeInfo `json:"MachineGroupType,omitempty" name:"MachineGroupType"`

	// Tag description list. This parameter is used to bind a tag to a machine group. Up to 10 tag key-value pairs are supported, and a resource can be bound to only one tag key.
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// Whether to enable automatic update for the machine group
	AutoUpdate *bool `json:"AutoUpdate,omitempty" name:"AutoUpdate"`

	// Update start time. We recommend you update LogListener during off-peak hours.
	UpdateStartTime *string `json:"UpdateStartTime,omitempty" name:"UpdateStartTime"`

	// Update end time. We recommend you update LogListener during off-peak hours.
	UpdateEndTime *string `json:"UpdateEndTime,omitempty" name:"UpdateEndTime"`

	// Whether to enable the service log to record the logs generated by the LogListener service itself. After it is enabled, the internal logset `cls_service_logging` and the `loglistener_status`, `loglistener_alarm`, and `loglistener_business` log topics will be created, which will not incur fees
	ServiceLogging *bool `json:"ServiceLogging,omitempty" name:"ServiceLogging"`
}

func (r *CreateMachineGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMachineGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupName")
	delete(f, "MachineGroupType")
	delete(f, "Tags")
	delete(f, "AutoUpdate")
	delete(f, "UpdateStartTime")
	delete(f, "UpdateEndTime")
	delete(f, "ServiceLogging")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateMachineGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateMachineGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Machine group ID
		GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateMachineGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMachineGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateShipperRequest struct {
	*tchttp.BaseRequest

	// ID of the log topic to which the shipping rule to be created belongs
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// Destination bucket in the shipping rule to be created
	Bucket *string `json:"Bucket,omitempty" name:"Bucket"`

	// Prefix of the shipping directory in the shipping rule to be created
	Prefix *string `json:"Prefix,omitempty" name:"Prefix"`

	// Shipping rule name
	ShipperName *string `json:"ShipperName,omitempty" name:"ShipperName"`

	// Shipping time interval in seconds. Default value: 300. Value range: 300–900
	Interval *uint64 `json:"Interval,omitempty" name:"Interval"`

	// Maximum size of a file to be shipped, in MB. Default value: 256. Value range: 100–256
	MaxSize *uint64 `json:"MaxSize,omitempty" name:"MaxSize"`

	// Filter rules for shipped logs. Only logs matching the rules can be shipped. All rules are in the AND relationship, and up to five rules can be added. If the array is empty, no filtering will be performed, and all logs will be shipped
	FilterRules []*FilterRuleInfo `json:"FilterRules,omitempty" name:"FilterRules"`

	// Partition rule of shipped log, which can be represented in `strftime` time format
	Partition *string `json:"Partition,omitempty" name:"Partition"`

	// Compression configuration of shipped log
	Compress *CompressInfo `json:"Compress,omitempty" name:"Compress"`

	// Format configuration of shipped log content
	Content *ContentInfo `json:"Content,omitempty" name:"Content"`
}

func (r *CreateShipperRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateShipperRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	delete(f, "Bucket")
	delete(f, "Prefix")
	delete(f, "ShipperName")
	delete(f, "Interval")
	delete(f, "MaxSize")
	delete(f, "FilterRules")
	delete(f, "Partition")
	delete(f, "Compress")
	delete(f, "Content")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateShipperRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateShipperResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Shipping rule ID
		ShipperId *string `json:"ShipperId,omitempty" name:"ShipperId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateShipperResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateShipperResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateTopicRequest struct {
	*tchttp.BaseRequest

	// Logset ID
	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`

	// Log topic name
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// Number of log topic partitions. Default value: 1. Maximum value: 10
	PartitionCount *int64 `json:"PartitionCount,omitempty" name:"PartitionCount"`

	// Tag description list. This parameter is used to bind a tag to a log topic. Up to 10 tag key-value pairs are supported, and a resource can be bound to only one tag key.
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// Whether to enable automatic split. Default value: true
	AutoSplit *bool `json:"AutoSplit,omitempty" name:"AutoSplit"`

	// Maximum number of partitions to split into for this topic if automatic split is enabled. Default value: 50
	MaxSplitPartitions *int64 `json:"MaxSplitPartitions,omitempty" name:"MaxSplitPartitions"`

	// Log topic storage class. Valid values: `hot`: real-time storage; `cold`: offline storage. Default value: `hot`. If `cold` is passed in, please contact the customer service to add the log topic to the allowlist first.
	StorageType *string `json:"StorageType,omitempty" name:"StorageType"`

	// Lifecycle in days. Value range: 1–366. Default value: 30
	Period *int64 `json:"Period,omitempty" name:"Period"`
}

func (r *CreateTopicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTopicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LogsetId")
	delete(f, "TopicName")
	delete(f, "PartitionCount")
	delete(f, "Tags")
	delete(f, "AutoSplit")
	delete(f, "MaxSplitPartitions")
	delete(f, "StorageType")
	delete(f, "Period")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTopicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateTopicResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Log topic ID
		TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTopicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTopicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CsvInfo struct {

	// Whether to print `key` on the first row of the CSV file
	PrintKey *bool `json:"PrintKey,omitempty" name:"PrintKey"`

	// Names of keys
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Keys []*string `json:"Keys,omitempty" name:"Keys"`

	// Field delimiter
	Delimiter *string `json:"Delimiter,omitempty" name:"Delimiter"`

	// Escape character used to enclose any field delimiter in field content. You can enter only a single quotation mark, double quotation mark, or an empty string.
	EscapeChar *string `json:"EscapeChar,omitempty" name:"EscapeChar"`

	// Content used to populate non-existing fields
	NonExistingField *string `json:"NonExistingField,omitempty" name:"NonExistingField"`
}

type DeleteAlarmNoticeRequest struct {
	*tchttp.BaseRequest

	// Alarm notification template
	AlarmNoticeId *string `json:"AlarmNoticeId,omitempty" name:"AlarmNoticeId"`
}

func (r *DeleteAlarmNoticeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAlarmNoticeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AlarmNoticeId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAlarmNoticeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAlarmNoticeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAlarmNoticeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAlarmNoticeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAlarmRequest struct {
	*tchttp.BaseRequest

	// Alarm policy ID.
	AlarmId *string `json:"AlarmId,omitempty" name:"AlarmId"`
}

func (r *DeleteAlarmRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAlarmRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AlarmId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAlarmRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAlarmResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAlarmResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAlarmResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAsyncContextTaskRequest struct {
	*tchttp.BaseRequest

	// Log topic ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// Offline context search task ID
	AsyncContextTaskId *string `json:"AsyncContextTaskId,omitempty" name:"AsyncContextTaskId"`
}

func (r *DeleteAsyncContextTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAsyncContextTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	delete(f, "AsyncContextTaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAsyncContextTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAsyncContextTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAsyncContextTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAsyncContextTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAsyncSearchTaskRequest struct {
	*tchttp.BaseRequest

	// Offline search task ID
	AsyncSearchTaskId *string `json:"AsyncSearchTaskId,omitempty" name:"AsyncSearchTaskId"`

	// Log topic ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
}

func (r *DeleteAsyncSearchTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAsyncSearchTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AsyncSearchTaskId")
	delete(f, "TopicId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAsyncSearchTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAsyncSearchTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAsyncSearchTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAsyncSearchTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteConfigFromMachineGroupRequest struct {
	*tchttp.BaseRequest

	// Machine group ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// Collection configuration ID
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`
}

func (r *DeleteConfigFromMachineGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteConfigFromMachineGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "ConfigId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteConfigFromMachineGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteConfigFromMachineGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteConfigFromMachineGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteConfigFromMachineGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteConfigRequest struct {
	*tchttp.BaseRequest

	// Collection rule configuration ID
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`
}

func (r *DeleteConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ConfigId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteExportRequest struct {
	*tchttp.BaseRequest

	// Log export ID
	ExportId *string `json:"ExportId,omitempty" name:"ExportId"`
}

func (r *DeleteExportRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteExportRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ExportId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteExportRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteExportResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteExportResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteExportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteIndexRequest struct {
	*tchttp.BaseRequest

	// Log topic ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
}

func (r *DeleteIndexRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteIndexRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteIndexRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteIndexResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteIndexResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteIndexResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteLogsetRequest struct {
	*tchttp.BaseRequest

	// Logset ID
	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`
}

func (r *DeleteLogsetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLogsetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LogsetId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteLogsetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteLogsetResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteLogsetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLogsetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMachineGroupRequest struct {
	*tchttp.BaseRequest

	// Machine group ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DeleteMachineGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMachineGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteMachineGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMachineGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteMachineGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMachineGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteShipperRequest struct {
	*tchttp.BaseRequest

	// Shipping rule ID
	ShipperId *string `json:"ShipperId,omitempty" name:"ShipperId"`
}

func (r *DeleteShipperRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteShipperRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ShipperId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteShipperRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteShipperResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteShipperResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteShipperResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteTopicRequest struct {
	*tchttp.BaseRequest

	// Log topic ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
}

func (r *DeleteTopicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTopicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTopicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteTopicResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteTopicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTopicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmNoticesRequest struct {
	*tchttp.BaseRequest

	// <br><li> name
	// 
	// Filter by **alarm notification template name**.
	// Type: String
	// 
	// Required: no
	// 
	// <br><li> alarmNoticeId
	// 
	// Filter by **alarm notification template ID**.
	// Type: String
	// 
	// Required: no
	// 
	// <br><li> uid
	// 
	// Filter by **recipient ID**.
	// 
	// Type: String
	// 
	// Required: no
	// 
	// <br><li> groupId
	// 
	// Filter by **user group ID**.
	// 
	// Type: String
	// 
	// Required: no
	// 
	// Each request can contain up to 10 `Filters` and 5 `Filter.Values`.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// Page offset. Default value: 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Maximum number of entries per page. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeAlarmNoticesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlarmNoticesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAlarmNoticesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmNoticesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Alarm notification template list
	// Note: this field may return `null`, indicating that no valid values can be obtained.
		AlarmNotices []*AlarmNotice `json:"AlarmNotices,omitempty" name:"AlarmNotices"`

		// Total number of eligible alarm notification templates
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAlarmNoticesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlarmNoticesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmsRequest struct {
	*tchttp.BaseRequest

	// <br><li> name
	// 
	// Filter by **alarm policy name**.
	// Type: String
	// 
	// Required: no
	// 
	// <br><li> alarmId
	// 
	// Filter by **alarm policy ID**.
	// Type: String
	// 
	// Required: no
	// 
	// <br><li> topicId
	// 
	// Filter by **log topic ID of monitoring object**.
	// 
	// Type: String
	// 
	// Required: no
	// 
	// <br><li> enable
	// 
	// Filter by **enablement status**.
	// 
	// Type: String
	// 
	// Required: no
	// 
	// Each request can contain up to 10 `Filters` and 5 `Filter.Values`.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// Page offset. Default value: 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Maximum number of entries per page. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeAlarmsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlarmsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAlarmsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Alarm policy list
		Alarms []*AlarmInfo `json:"Alarms,omitempty" name:"Alarms"`

		// Number of eligible alarm policies
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAlarmsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlarmsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAsyncContextResultRequest struct {
	*tchttp.BaseRequest

	// Offline search task ID
	AsyncContextTaskId *string `json:"AsyncContextTaskId,omitempty" name:"AsyncContextTaskId"`

	// Log package number
	PkgId *string `json:"PkgId,omitempty" name:"PkgId"`

	// Log number in log package
	PkgLogId *string `json:"PkgLogId,omitempty" name:"PkgLogId"`

	// Log topic ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// Number of previous logs. Default value: 10
	PrevLogs *int64 `json:"PrevLogs,omitempty" name:"PrevLogs"`

	// Number of next logs. Default value: 10
	NextLogs *int64 `json:"NextLogs,omitempty" name:"NextLogs"`
}

func (r *DescribeAsyncContextResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAsyncContextResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AsyncContextTaskId")
	delete(f, "PkgId")
	delete(f, "PkgLogId")
	delete(f, "TopicId")
	delete(f, "PrevLogs")
	delete(f, "NextLogs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAsyncContextResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAsyncContextResultResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Whether the previous logs have been returned
		PrevOver *bool `json:"PrevOver,omitempty" name:"PrevOver"`

		// Whether the next logs have been returned
		NextOver *bool `json:"NextOver,omitempty" name:"NextOver"`

		// Log content
		Results []*LogInfo `json:"Results,omitempty" name:"Results"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAsyncContextResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAsyncContextResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAsyncContextTasksRequest struct {
	*tchttp.BaseRequest

	// Page offset. Default value: 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Maximum number of entries per page. Default value: 20. Maximum value: 100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// <br><li> topicId
	// 
	// Filter by **log topic ID**.
	// Type: String
	// 
	// Required: no
	// 
	// <br><li> logsetId
	// 
	// Filter by **logset ID**. You can call `DescribeLogsets` to query the list of created logsets or log in to the console to view them. You can also call `CreateLogset` to create a logset.
	// 
	// Type: String
	// 
	// Required: no
	// 
	// Each request can contain up to 10 `Filters` and 5 `Filter.Values`.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeAsyncContextTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAsyncContextTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAsyncContextTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAsyncContextTasksResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Offline context search task list
	// Note: this field may return `null`, indicating that no valid values can be obtained.
		AsyncContextTasks []*AsyncContextTask `json:"AsyncContextTasks,omitempty" name:"AsyncContextTasks"`

		// Total number of offline context search tasks
	// Note: this field may return `null`, indicating that no valid values can be obtained.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAsyncContextTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAsyncContextTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAsyncSearchResultRequest struct {
	*tchttp.BaseRequest

	// Offline search task ID
	AsyncSearchTaskId *string `json:"AsyncSearchTaskId,omitempty" name:"AsyncSearchTaskId"`

	// Logset ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// This field is used to load more logs. Pass through the last `Context` value returned to get more log content.
	Context *string `json:"Context,omitempty" name:"Context"`

	// Number of logs returned in a single call. Default value: 20. Maximum value: 500
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeAsyncSearchResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAsyncSearchResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AsyncSearchTaskId")
	delete(f, "TopicId")
	delete(f, "Context")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAsyncSearchResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAsyncSearchResultResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// `Context` for loading subsequent content
		Context *string `json:"Context,omitempty" name:"Context"`

		// Whether all log query results are returned
		ListOver *bool `json:"ListOver,omitempty" name:"ListOver"`

		// Log content
		Results []*LogInfo `json:"Results,omitempty" name:"Results"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAsyncSearchResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAsyncSearchResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAsyncSearchTasksRequest struct {
	*tchttp.BaseRequest

	// Page offset. Default value: 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Maximum number of entries per page. Default value: 20. Maximum value: 100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// <br><li> topicId
	// 
	// Filter by **log topic ID**.
	// Type: String
	// 
	// Required: no
	// 
	// <br><li> logsetId
	// 
	// Filter by **logset ID**. You can call `DescribeLogsets` to query the list of created logsets or log in to the console to view them. You can also call `CreateLogset` to create a logset.
	// 
	// Type: String
	// 
	// Required: no
	// 
	// Each request can contain up to 10 `Filters` and 5 `Filter.Values`.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeAsyncSearchTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAsyncSearchTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAsyncSearchTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAsyncSearchTasksResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Offline search task list
	// Note: this field may return `null`, indicating that no valid values can be obtained.
		AsyncSearchTasks []*AsyncSearchTask `json:"AsyncSearchTasks,omitempty" name:"AsyncSearchTasks"`

		// Total number of offline search tasks
	// Note: this field may return `null`, indicating that no valid values can be obtained.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAsyncSearchTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAsyncSearchTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeConfigMachineGroupsRequest struct {
	*tchttp.BaseRequest

	// Collection configuration ID
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`
}

func (r *DescribeConfigMachineGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConfigMachineGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ConfigId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeConfigMachineGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeConfigMachineGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// List of machine groups bound to the collection rule configuration
	// Note: this field may return `null`, indicating that no valid values can be obtained.
		MachineGroups []*MachineGroupInfo `json:"MachineGroups,omitempty" name:"MachineGroups"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeConfigMachineGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConfigMachineGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeConfigsRequest struct {
	*tchttp.BaseRequest

	// <br><li> configName
	// 
	// Filter by fuzzy match of **collection configuration name**
	// Type: String
	// 
	// Required: no
	// 
	// <br><li> configId
	// 
	// Filter by **collection configuration ID**.
	// Type: String
	// 
	// Required: no
	// 
	// <br><li> topicId
	// 
	// Filter by **log topic**.
	// 
	// Type: String
	// 
	// Required: no
	// 
	// Each request can contain up to 10 `Filters` and 5 `Filter.Values`.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// Page offset. Default value: 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Maximum number of entries per page. Default value: 20. Maximum value: 100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeConfigsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConfigsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeConfigsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeConfigsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Collection configuration list
	// Note: this field may return `null`, indicating that no valid values can be obtained.
		Configs []*ConfigInfo `json:"Configs,omitempty" name:"Configs"`

		// Total number of filtered items
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeConfigsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConfigsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeExportsRequest struct {
	*tchttp.BaseRequest

	// Log topic ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// Page offset. Default value: 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Maximum number of entries per page. Default value: 20. Maximum value: 100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeExportsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExportsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeExportsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeExportsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// List of exported logs
		Exports []*ExportInfo `json:"Exports,omitempty" name:"Exports"`

		// Total number
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeExportsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExportsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIndexRequest struct {
	*tchttp.BaseRequest

	// Log topic ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
}

func (r *DescribeIndexRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIndexRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIndexRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIndexResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Log topic ID
		TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

		// Whether it takes effect
		Status *bool `json:"Status,omitempty" name:"Status"`

		// Index configuration information
	// Note: this field may return `null`, indicating that no valid values can be obtained.
		Rule *RuleInfo `json:"Rule,omitempty" name:"Rule"`

		// Index modification time. The default value is the index creation time.
		ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeIndexResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIndexResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLogContextRequest struct {
	*tchttp.BaseRequest

	// Log topic ID to be queried
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// Log time in the format of `YYYY-mm-dd HH:MM:SS`
	BTime *string `json:"BTime,omitempty" name:"BTime"`

	// Log package number
	PkgId *string `json:"PkgId,omitempty" name:"PkgId"`

	// Log number in log package
	PkgLogId *int64 `json:"PkgLogId,omitempty" name:"PkgLogId"`

	// Number of previous logs. Default value: 10
	PrevLogs *int64 `json:"PrevLogs,omitempty" name:"PrevLogs"`

	// Number of next logs. Default value: 10
	NextLogs *int64 `json:"NextLogs,omitempty" name:"NextLogs"`
}

func (r *DescribeLogContextRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogContextRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	delete(f, "BTime")
	delete(f, "PkgId")
	delete(f, "PkgLogId")
	delete(f, "PrevLogs")
	delete(f, "NextLogs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLogContextRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLogContextResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Log context information set
		LogContextInfos []*LogContextInfo `json:"LogContextInfos,omitempty" name:"LogContextInfos"`

		// Whether the previous logs have been returned
		PrevOver *bool `json:"PrevOver,omitempty" name:"PrevOver"`

		// Whether the next logs have been returned
		NextOver *bool `json:"NextOver,omitempty" name:"NextOver"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLogContextResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogContextResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLogsetsRequest struct {
	*tchttp.BaseRequest

	// <br><li> logsetName
	// 
	// Filter by **logset name**.
	// Type: String
	// 
	// Required: no
	// 
	// <br><li> logsetId
	// 
	// Filter by **logset ID**.
	// Type: String
	// 
	// Required: no
	// 
	// <br><li> tagKey
	// 
	// Filter by **tag key**.
	// 
	// Type: String
	// 
	// Required: no
	// 
	// <br><li> tag:tagKey
	// 
	// Filter by **tag key-value pair**. The `tagKey` should be replaced with a specified tag key.
	// Type: String
	// 
	// Required: no
	// 
	// 
	// Each request can contain up to 10 `Filters` and 5 `Filter.Values`.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// Page offset. Default value: 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Maximum number of entries per page. Default value: 20. Maximum value: 100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeLogsetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogsetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLogsetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLogsetsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of pages
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Logset list
		Logsets []*LogsetInfo `json:"Logsets,omitempty" name:"Logsets"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLogsetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogsetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMachineGroupConfigsRequest struct {
	*tchttp.BaseRequest

	// Machine group ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DescribeMachineGroupConfigsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMachineGroupConfigsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMachineGroupConfigsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMachineGroupConfigsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Collection rule configuration list
	// Note: this field may return `null`, indicating that no valid values can be obtained.
		Configs []*ConfigInfo `json:"Configs,omitempty" name:"Configs"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMachineGroupConfigsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMachineGroupConfigsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMachineGroupsRequest struct {
	*tchttp.BaseRequest

	// <br><li> machineGroupName
	// 
	// Filter by **machine group name**.
	// Type: String
	// 
	// Required: no
	// 
	// <br><li> machineGroupId
	// 
	// Filter by **machine group ID**.
	// Type: String
	// 
	// Required: no
	// 
	// <br><li> tagKey
	// 
	// Filter by **tag key**.
	// 
	// Type: String
	// 
	// Required: no
	// 
	// <br><li> tag:tagKey
	// 
	// Filter by **tag key-value pair**. The `tagKey` should be replaced with a specified tag key.
	// Type: String
	// 
	// Required: no
	// 
	// 
	// Each request can contain up to 10 `Filters` and 5 `Filter.Values`.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// Page offset. Default value: 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Maximum number of entries per page. Default value: 20. Maximum value: 100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeMachineGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMachineGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMachineGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMachineGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Machine group information list
	// Note: this field may return `null`, indicating that no valid values can be obtained.
		MachineGroups []*MachineGroupInfo `json:"MachineGroups,omitempty" name:"MachineGroups"`

		// Total number of pages
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMachineGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMachineGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMachinesRequest struct {
	*tchttp.BaseRequest

	// ID of the machine group to be queried
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DescribeMachinesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMachinesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMachinesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMachinesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Group of machine status information
		Machines []*MachineInfo `json:"Machines,omitempty" name:"Machines"`

		// Whether to enable the automatic update feature for the machine group
		AutoUpdate *int64 `json:"AutoUpdate,omitempty" name:"AutoUpdate"`

		// Preset start time of automatic update of machine group
		UpdateStartTime *string `json:"UpdateStartTime,omitempty" name:"UpdateStartTime"`

		// Preset end time of automatic update of machine group
		UpdateEndTime *string `json:"UpdateEndTime,omitempty" name:"UpdateEndTime"`

		// Latest LogListener version available to the current user
		LatestAgentVersion *string `json:"LatestAgentVersion,omitempty" name:"LatestAgentVersion"`

		// Whether to enable the service log
		ServiceLogging *bool `json:"ServiceLogging,omitempty" name:"ServiceLogging"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMachinesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMachinesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePartitionsRequest struct {
	*tchttp.BaseRequest

	// Log topic ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
}

func (r *DescribePartitionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePartitionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePartitionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribePartitionsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Partition list
		Partitions []*PartitionInfo `json:"Partitions,omitempty" name:"Partitions"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePartitionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePartitionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeShipperTasksRequest struct {
	*tchttp.BaseRequest

	// Shipping rule ID
	ShipperId *string `json:"ShipperId,omitempty" name:"ShipperId"`

	// Query start timestamp in milliseconds, which can be within the last three days
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// Query end timestamp in milliseconds
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeShipperTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeShipperTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ShipperId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeShipperTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeShipperTasksResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Shipping task list
	// Note: this field may return `null`, indicating that no valid values can be obtained.
		Tasks []*ShipperTaskInfo `json:"Tasks,omitempty" name:"Tasks"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeShipperTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeShipperTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeShippersRequest struct {
	*tchttp.BaseRequest

	// <br><li> shipperName
	// 
	// Filter by **shipping rule name**.
	// Type: String
	// 
	// Required: no
	// 
	// <br><li> shipperId
	// 
	// Filter by **shipping rule ID**.
	// Type: String
	// 
	// Required: no
	// 
	// <br><li> topicId
	// 
	// Filter by **log topic**.
	// 
	// Type: String
	// 
	// Required: no
	// 
	// Each request can contain up to 10 `Filters` and 5 `Filter.Values`.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// Page offset. Default value: 0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Maximum number of entries per page. Default value: 20. Maximum value: 100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeShippersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeShippersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeShippersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeShippersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Shipping rule list
	// Note: this field may return `null`, indicating that no valid values can be obtained.
		Shippers []*ShipperInfo `json:"Shippers,omitempty" name:"Shippers"`

		// Total number of results obtained in this query
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeShippersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeShippersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTopicsRequest struct {
	*tchttp.BaseRequest

	// <br><li> topicName
	// 
	// Filter by **log topic name**.
	// Type: String
	// 
	// Required: no
	// 
	// <br><li> topicId
	// 
	// Filter by **log topic ID**.
	// Type: String
	// 
	// Required: no
	// 
	// <br><li> logsetId
	// 
	// Filter by **logset ID**. You can call `DescribeLogsets` to query the list of created logsets or log in to the console to view them. You can also call `CreateLogset` to create a logset.
	// 
	// Type: String
	// 
	// Required: no
	// 
	// <br><li> tagKey
	// 
	// Filter by **tag key**.
	// 
	// Type: String
	// 
	// Required: no
	// 
	// <br><li> tag:tagKey
	// 
	// Filter by **tag key-value pair**. The `tag-key` should be replaced with a specified tag key. For more information on how to use it, please see sample 2.
	// 
	// Type: String
	// 
	// Required: no
	// 
	// <br><li> storageType
	// 
	// Filter by **log topic storage class**. Valid values: `hot`: real-time storage; `cold`: offline storage.
	// Type: String
	// 
	// Required: no
	// 
	// 
	// Each request can contain up to 10 `Filters` and 100 `Filter.Values`.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// Page offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Maximum number of entries per page. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeTopicsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopicsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTopicsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTopicsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Log topic list
		Topics []*TopicInfo `json:"Topics,omitempty" name:"Topics"`

		// Total number
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTopicsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopicsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExcludePathInfo struct {

	// Type. Valid values: `File`, `Path`
	Type *string `json:"Type,omitempty" name:"Type"`

	// Specific content corresponding to `Type`
	Value *string `json:"Value,omitempty" name:"Value"`
}

type ExportInfo struct {

	// Log topic ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// Log export task ID
	ExportId *string `json:"ExportId,omitempty" name:"ExportId"`

	// Log export query statement
	Query *string `json:"Query,omitempty" name:"Query"`

	// Log export filename
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// Log file size
	FileSize *uint64 `json:"FileSize,omitempty" name:"FileSize"`

	// Log export time sorting
	Order *string `json:"Order,omitempty" name:"Order"`

	// Log export format
	Format *string `json:"Format,omitempty" name:"Format"`

	// Number of logs to be exported
	Count *uint64 `json:"Count,omitempty" name:"Count"`

	// Log download status. Valid values: `Processing`: exporting; `Complete`: completed; `Failed`: failed; `Expired`: expired (3-day validity period).
	Status *string `json:"Status,omitempty" name:"Status"`

	// Log export start time
	From *int64 `json:"From,omitempty" name:"From"`

	// Log export end time
	To *int64 `json:"To,omitempty" name:"To"`

	// Log export path
	CosPath *string `json:"CosPath,omitempty" name:"CosPath"`

	// Log export creation time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type ExtractRuleInfo struct {

	// Time field key name. `time_key` and `time_format` must appear in pairs
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	TimeKey *string `json:"TimeKey,omitempty" name:"TimeKey"`

	// Time field format. For more information, please see the output parameters of the time format description of the `strftime` function in C language
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	TimeFormat *string `json:"TimeFormat,omitempty" name:"TimeFormat"`

	// Delimiter for delimited log, which is valid only if `log_type` is `delimiter_log`
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Delimiter *string `json:"Delimiter,omitempty" name:"Delimiter"`

	// Full log matching rule, which is valid only if `log_type` is `fullregex_log`
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	LogRegex *string `json:"LogRegex,omitempty" name:"LogRegex"`

	// First-Line matching rule, which is valid only if `log_type` is `multiline_log` or `fullregex_log`
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	BeginRegex *string `json:"BeginRegex,omitempty" name:"BeginRegex"`

	// Key name of each extracted field. An empty key indicates to discard the field. This parameter is valid only if `log_type` is `delimiter_log`. `json_log` logs use the key of JSON itself
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Keys []*string `json:"Keys,omitempty" name:"Keys"`

	// Log keys to be filtered and the corresponding regex
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	FilterKeyRegex []*KeyRegexInfo `json:"FilterKeyRegex,omitempty" name:"FilterKeyRegex"`

	// Whether to upload the logs that failed to be parsed. Valid values: `true`: yes; `false`: no
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	UnMatchUpLoadSwitch *bool `json:"UnMatchUpLoadSwitch,omitempty" name:"UnMatchUpLoadSwitch"`

	// Unmatched log key
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	UnMatchLogKey *string `json:"UnMatchLogKey,omitempty" name:"UnMatchLogKey"`

	// Size of the data to be rewound in incremental collection mode. Default value: -1 (full collection)
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Backtracking *int64 `json:"Backtracking,omitempty" name:"Backtracking"`
}

type Filter struct {

	// Field to be filtered
	Key *string `json:"Key,omitempty" name:"Key"`

	// Value to be filtered
	Values []*string `json:"Values,omitempty" name:"Values"`
}

type FilterRuleInfo struct {

	// Filter rule key
	Key *string `json:"Key,omitempty" name:"Key"`

	// Filter rule
	Regex *string `json:"Regex,omitempty" name:"Regex"`

	// Filter rule value
	Value *string `json:"Value,omitempty" name:"Value"`
}

type FullTextInfo struct {

	// Case sensitivity
	CaseSensitive *bool `json:"CaseSensitive,omitempty" name:"CaseSensitive"`

	// Full-Text index delimiter. Each character in the string represents a delimiter.
	Tokenizer *string `json:"Tokenizer,omitempty" name:"Tokenizer"`

	// Whether Chinese characters are contained
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ContainZH *bool `json:"ContainZH,omitempty" name:"ContainZH"`
}

type GetAlarmLogRequest struct {
	*tchttp.BaseRequest

	// Start time of the log to be queried, which is a Unix timestamp in milliseconds
	From *int64 `json:"From,omitempty" name:"From"`

	// End time of the log to be queried, which is a Unix timestamp in milliseconds
	To *int64 `json:"To,omitempty" name:"To"`

	// Query statement. Maximum length: 1024
	Query *string `json:"Query,omitempty" name:"Query"`

	// Number of logs returned in a single query. Maximum value: 100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// This field is used to load more logs. Pass through the last `Context` value returned to get more log content.
	Context *string `json:"Context,omitempty" name:"Context"`

	// Order of the logs sorted by time returned by the log API. Valid values: `asc`: ascending; `desc`: descending. Default value: `desc`
	Sort *string `json:"Sort,omitempty" name:"Sort"`

	// 
	UseNewAnalysis *bool `json:"UseNewAnalysis,omitempty" name:"UseNewAnalysis"`
}

func (r *GetAlarmLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetAlarmLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "From")
	delete(f, "To")
	delete(f, "Query")
	delete(f, "Limit")
	delete(f, "Context")
	delete(f, "Sort")
	delete(f, "UseNewAnalysis")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetAlarmLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetAlarmLogResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// `Context` for loading subsequent content
		Context *string `json:"Context,omitempty" name:"Context"`

		// Whether all log query results are returned
		ListOver *bool `json:"ListOver,omitempty" name:"ListOver"`

		// Whether the return is the analysis result
		Analysis *bool `json:"Analysis,omitempty" name:"Analysis"`

		// If `Analysis` is `true`, column name of the analysis result will be returned; otherwise, empty content will be returned.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
		ColNames []*string `json:"ColNames,omitempty" name:"ColNames"`

		// Log query result. If `Analysis` is `True`, `null` may be returned
	// Note: this field may return `null`, indicating that no valid values can be obtained.
		Results []*LogInfo `json:"Results,omitempty" name:"Results"`

		// Log analysis result. If `Analysis` is `False`, `null` may be returned
	// Note: this field may return `null`, indicating that no valid values can be obtained.
		AnalysisResults []*LogItems `json:"AnalysisResults,omitempty" name:"AnalysisResults"`

		// 
		AnalysisRecords []*string `json:"AnalysisRecords,omitempty" name:"AnalysisRecords"`

		// 
		Columns []*Column `json:"Columns,omitempty" name:"Columns"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetAlarmLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetAlarmLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type JsonInfo struct {

	// Enablement flag
	EnableTag *bool `json:"EnableTag,omitempty" name:"EnableTag"`

	// Metadata information list
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	MetaFields []*string `json:"MetaFields,omitempty" name:"MetaFields"`
}

type KeyRegexInfo struct {

	// Log key to be filtered
	Key *string `json:"Key,omitempty" name:"Key"`

	// Filter rule regex corresponding to key
	Regex *string `json:"Regex,omitempty" name:"Regex"`
}

type KeyValueInfo struct {

	// The field that needs to be configured with a key value or metafield index
	Key *string `json:"Key,omitempty" name:"Key"`

	// Field index description information
	Value *ValueInfo `json:"Value,omitempty" name:"Value"`
}

type LogContextInfo struct {

	// Log source device
	Source *string `json:"Source,omitempty" name:"Source"`

	// Collection path
	Filename *string `json:"Filename,omitempty" name:"Filename"`

	// Log content
	Content *string `json:"Content,omitempty" name:"Content"`

	// Log package number
	PkgId *string `json:"PkgId,omitempty" name:"PkgId"`

	// Log number in log package
	PkgLogId *int64 `json:"PkgLogId,omitempty" name:"PkgLogId"`

	// Log timestamp
	BTime *int64 `json:"BTime,omitempty" name:"BTime"`
}

type LogInfo struct {

	// Log time in milliseconds
	Time *int64 `json:"Time,omitempty" name:"Time"`

	// Log topic ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// Log topic name
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// Log source IP
	Source *string `json:"Source,omitempty" name:"Source"`

	// Log filename
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// ID of the request package for log reporting
	PkgId *string `json:"PkgId,omitempty" name:"PkgId"`

	// Log ID in request package
	PkgLogId *string `json:"PkgLogId,omitempty" name:"PkgLogId"`

	// Serialized JSON string of log content
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	LogJson *string `json:"LogJson,omitempty" name:"LogJson"`
}

type LogItem struct {

	// Log key
	Key *string `json:"Key,omitempty" name:"Key"`

	// Log value
	Value *string `json:"Value,omitempty" name:"Value"`
}

type LogItems struct {

	// Key-Value pair returned in analysis result
	Data []*LogItem `json:"Data,omitempty" name:"Data"`
}

type LogsetInfo struct {

	// Logset ID
	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`

	// Logset name
	LogsetName *string `json:"LogsetName,omitempty" name:"LogsetName"`

	// Creation time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Tag bound to logset
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// Number of log topics in logset
	TopicCount *int64 `json:"TopicCount,omitempty" name:"TopicCount"`

	// If `AssumerUin` is not empty, it indicates the service provider who creates the logset
	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
}

type MachineGroupInfo struct {

	// Machine group ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// Machine group name
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// Machine group type
	MachineGroupType *MachineGroupTypeInfo `json:"MachineGroupType,omitempty" name:"MachineGroupType"`

	// Creation time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// List of tags bound to machine group
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// Whether to enable automatic update for the machine group
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	AutoUpdate *string `json:"AutoUpdate,omitempty" name:"AutoUpdate"`

	// Update start time. We recommend you update LogListener during off-peak hours.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	UpdateStartTime *string `json:"UpdateStartTime,omitempty" name:"UpdateStartTime"`

	// Update end time. We recommend you update LogListener during off-peak hours.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	UpdateEndTime *string `json:"UpdateEndTime,omitempty" name:"UpdateEndTime"`

	// Whether to enable the service log to record the logs generated by the LogListener service itself. After it is enabled, the internal logset `cls_service_logging` and the `loglistener_status`, `loglistener_alarm`, and `loglistener_business` log topics will be created, which will not incur fees.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ServiceLogging *bool `json:"ServiceLogging,omitempty" name:"ServiceLogging"`
}

type MachineGroupTypeInfo struct {

	// Machine group type. Valid values: `ip`: the IP addresses of collection machines are stored in `Values` of the machine group; `label`: the tags of the machines are stored in `Values` of the machine group
	Type *string `json:"Type,omitempty" name:"Type"`

	// Machine description list
	Values []*string `json:"Values,omitempty" name:"Values"`
}

type MachineInfo struct {

	// Machine IP
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// Machine status. Valid values: `0`: exceptional; `1`: normal
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// Machine disconnection time. If the value is empty, the machine is normal. If the machine is exceptional, a specific value will be returned.
	OfflineTime *string `json:"OfflineTime,omitempty" name:"OfflineTime"`

	// Whether to enable automatic update for the machine. Valid values: `0`: no; `1`: yes
	AutoUpdate *int64 `json:"AutoUpdate,omitempty" name:"AutoUpdate"`

	// Current machine version number
	Version *string `json:"Version,omitempty" name:"Version"`

	// Machine update feature status
	UpdateStatus *int64 `json:"UpdateStatus,omitempty" name:"UpdateStatus"`

	// Machine update result flag
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// Machine update result information
	ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`
}

type MergePartitionRequest struct {
	*tchttp.BaseRequest

	// Log topic ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// Merged `PartitionId`
	PartitionId *int64 `json:"PartitionId,omitempty" name:"PartitionId"`
}

func (r *MergePartitionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *MergePartitionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	delete(f, "PartitionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "MergePartitionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type MergePartitionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Merge result set
		Partitions []*PartitionInfo `json:"Partitions,omitempty" name:"Partitions"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *MergePartitionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *MergePartitionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAlarmNoticeRequest struct {
	*tchttp.BaseRequest

	// Alarm notification template ID
	AlarmNoticeId *string `json:"AlarmNoticeId,omitempty" name:"AlarmNoticeId"`

	// Alarm template name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Alarm template type. Valid values:
	// <br><li> `Trigger`: alarm triggered
	// <br><li> `Recovery`: alarm cleared
	// <br><li> `All`: alarm triggered and alarm cleared
	Type *string `json:"Type,omitempty" name:"Type"`

	// Information of the recipient in alarm template
	NoticeReceivers []*NoticeReceiver `json:"NoticeReceivers,omitempty" name:"NoticeReceivers"`

	// Alarm template callback information
	WebCallbacks []*WebCallback `json:"WebCallbacks,omitempty" name:"WebCallbacks"`
}

func (r *ModifyAlarmNoticeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAlarmNoticeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AlarmNoticeId")
	delete(f, "Name")
	delete(f, "Type")
	delete(f, "NoticeReceivers")
	delete(f, "WebCallbacks")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAlarmNoticeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAlarmNoticeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAlarmNoticeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAlarmNoticeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAlarmRequest struct {
	*tchttp.BaseRequest

	// Alarm policy ID
	AlarmId *string `json:"AlarmId,omitempty" name:"AlarmId"`

	// Alarm policy name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Monitoring task running time point
	MonitorTime *MonitorTime `json:"MonitorTime,omitempty" name:"MonitorTime"`

	// Trigger condition
	Condition *string `json:"Condition,omitempty" name:"Condition"`

	// Alarm persistence cycle. An alarm will be triggered only after the corresponding trigger condition is met for the number of times specified by `TriggerCount`. Value range: 1–10.
	TriggerCount *int64 `json:"TriggerCount,omitempty" name:"TriggerCount"`

	// Repeated alarm interval in minutes. Value range: 0–1440.
	AlarmPeriod *int64 `json:"AlarmPeriod,omitempty" name:"AlarmPeriod"`

	// List of associated alarm notification templates
	AlarmNoticeIds []*string `json:"AlarmNoticeIds,omitempty" name:"AlarmNoticeIds"`

	// Monitoring object list
	AlarmTargets []*AlarmTarget `json:"AlarmTargets,omitempty" name:"AlarmTargets"`

	// Whether to enable the alarm policy
	Status *bool `json:"Status,omitempty" name:"Status"`

	// Custom alarm content
	MessageTemplate *string `json:"MessageTemplate,omitempty" name:"MessageTemplate"`

	// Custom callback
	CallBack *CallBackInfo `json:"CallBack,omitempty" name:"CallBack"`

	// Multi-Dimensional analysis
	Analysis []*AnalysisDimensional `json:"Analysis,omitempty" name:"Analysis"`
}

func (r *ModifyAlarmRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAlarmRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AlarmId")
	delete(f, "Name")
	delete(f, "MonitorTime")
	delete(f, "Condition")
	delete(f, "TriggerCount")
	delete(f, "AlarmPeriod")
	delete(f, "AlarmNoticeIds")
	delete(f, "AlarmTargets")
	delete(f, "Status")
	delete(f, "MessageTemplate")
	delete(f, "CallBack")
	delete(f, "Analysis")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAlarmRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAlarmResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAlarmResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAlarmResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyConfigRequest struct {
	*tchttp.BaseRequest

	// Collection rule configuration ID
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`

	// Collection rule configuration name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Log collection path containing the filename
	Path *string `json:"Path,omitempty" name:"Path"`

	// Type of the log to be collected. Valid values: `json_log`: log in JSON format; `delimiter_log`: log in delimited format; `minimalist_log`: minimalist log; `multiline_log`: log in multi-line format; `fullregex_log`: log in full regex format. Default value: `minimalist_log`
	LogType *string `json:"LogType,omitempty" name:"LogType"`

	// Extraction rule. If `ExtractRule` is set, `LogType` must be set.
	ExtractRule *ExtractRuleInfo `json:"ExtractRule,omitempty" name:"ExtractRule"`

	// Collection path blocklist
	ExcludePaths []*ExcludePathInfo `json:"ExcludePaths,omitempty" name:"ExcludePaths"`

	// Log topic (TopicId) associated with collection configuration
	Output *string `json:"Output,omitempty" name:"Output"`

	// Custom parsing string, which is a serialized JSON string
	UserDefineRule *string `json:"UserDefineRule,omitempty" name:"UserDefineRule"`
}

func (r *ModifyConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ConfigId")
	delete(f, "Name")
	delete(f, "Path")
	delete(f, "LogType")
	delete(f, "ExtractRule")
	delete(f, "ExcludePaths")
	delete(f, "Output")
	delete(f, "UserDefineRule")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyIndexRequest struct {
	*tchttp.BaseRequest

	// Log topic ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// It does not take effect by default
	Status *bool `json:"Status,omitempty" name:"Status"`

	// Index rule. Either `Rule` or `Effective` must exist.
	Rule *RuleInfo `json:"Rule,omitempty" name:"Rule"`
}

func (r *ModifyIndexRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyIndexRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	delete(f, "Status")
	delete(f, "Rule")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyIndexRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyIndexResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyIndexResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyIndexResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLogsetRequest struct {
	*tchttp.BaseRequest

	// Logset ID
	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`

	// Logset name
	LogsetName *string `json:"LogsetName,omitempty" name:"LogsetName"`

	// Tag key-value pair bound to logset. Up to 10 tag key-value pairs are supported, and a resource can be bound to only one tag key at any time.
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

func (r *ModifyLogsetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLogsetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LogsetId")
	delete(f, "LogsetName")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLogsetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLogsetResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyLogsetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLogsetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyMachineGroupRequest struct {
	*tchttp.BaseRequest

	// Machine group ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// Machine group name
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// Machine group type
	MachineGroupType *MachineGroupTypeInfo `json:"MachineGroupType,omitempty" name:"MachineGroupType"`

	// Tag list
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// Whether to enable automatic update for the machine group
	AutoUpdate *bool `json:"AutoUpdate,omitempty" name:"AutoUpdate"`

	// Update start time. We recommend you update LogListener during off-peak hours.
	UpdateStartTime *string `json:"UpdateStartTime,omitempty" name:"UpdateStartTime"`

	// Update end time. We recommend you update LogListener during off-peak hours.
	UpdateEndTime *string `json:"UpdateEndTime,omitempty" name:"UpdateEndTime"`

	// Whether to enable the service log to record the logs generated by the LogListener service itself. After it is enabled, the internal logset `cls_service_logging` and the `loglistener_status`, `loglistener_alarm`, and `loglistener_business` log topics will be created, which will not incur fees.
	ServiceLogging *bool `json:"ServiceLogging,omitempty" name:"ServiceLogging"`
}

func (r *ModifyMachineGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMachineGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "GroupName")
	delete(f, "MachineGroupType")
	delete(f, "Tags")
	delete(f, "AutoUpdate")
	delete(f, "UpdateStartTime")
	delete(f, "UpdateEndTime")
	delete(f, "ServiceLogging")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyMachineGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyMachineGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyMachineGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMachineGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyShipperRequest struct {
	*tchttp.BaseRequest

	// Shipping rule ID
	ShipperId *string `json:"ShipperId,omitempty" name:"ShipperId"`

	// New destination bucket in shipping rule
	Bucket *string `json:"Bucket,omitempty" name:"Bucket"`

	// New destination directory prefix in shipping rule
	Prefix *string `json:"Prefix,omitempty" name:"Prefix"`

	// Shipping rule status
	Status *bool `json:"Status,omitempty" name:"Status"`

	// Shipping rule name
	ShipperName *string `json:"ShipperName,omitempty" name:"ShipperName"`

	// Shipping time interval in seconds. Default value: 300. Value range: 300–900
	Interval *uint64 `json:"Interval,omitempty" name:"Interval"`

	// Maximum size of a file to be shipped, in MB. Default value: 256. Value range: 100–256
	MaxSize *uint64 `json:"MaxSize,omitempty" name:"MaxSize"`

	// Filter rules for shipped logs. Only logs matching the rules can be shipped. All rules are in the AND relationship, and up to five rules can be added. If the array is empty, no filtering will be performed, and all logs will be shipped.
	FilterRules []*FilterRuleInfo `json:"FilterRules,omitempty" name:"FilterRules"`

	// Partition rule of shipped log, which can be represented in `strftime` time format
	Partition *string `json:"Partition,omitempty" name:"Partition"`

	// Compression configuration of shipped log
	Compress *CompressInfo `json:"Compress,omitempty" name:"Compress"`

	// Format configuration of shipped log content
	Content *ContentInfo `json:"Content,omitempty" name:"Content"`
}

func (r *ModifyShipperRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyShipperRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ShipperId")
	delete(f, "Bucket")
	delete(f, "Prefix")
	delete(f, "Status")
	delete(f, "ShipperName")
	delete(f, "Interval")
	delete(f, "MaxSize")
	delete(f, "FilterRules")
	delete(f, "Partition")
	delete(f, "Compress")
	delete(f, "Content")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyShipperRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyShipperResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyShipperResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyShipperResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTopicRequest struct {
	*tchttp.BaseRequest

	// Log topic ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// Log topic name
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// Tag description list. This parameter is used to bind a tag to a log topic. Up to 10 tag key-value pairs are supported, and they must be unique.
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// Whether to start collection for this log topic
	Status *bool `json:"Status,omitempty" name:"Status"`

	// Whether to enable automatic split
	AutoSplit *bool `json:"AutoSplit,omitempty" name:"AutoSplit"`

	// Maximum number of partitions to split into for this topic if automatic split is enabled
	MaxSplitPartitions *int64 `json:"MaxSplitPartitions,omitempty" name:"MaxSplitPartitions"`

	// Lifecycle in days. Value range: 1–366
	Period *int64 `json:"Period,omitempty" name:"Period"`
}

func (r *ModifyTopicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTopicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	delete(f, "TopicName")
	delete(f, "Tags")
	delete(f, "Status")
	delete(f, "AutoSplit")
	delete(f, "MaxSplitPartitions")
	delete(f, "Period")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTopicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTopicResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyTopicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTopicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MonitorTime struct {

	// Valid values:
	// <br><li> `Period`: periodic execution
	// <br><li> `Fixed`: scheduled execution
	Type *string `json:"Type,omitempty" name:"Type"`

	// Execution interval or scheduled time point in minutes. Value range: 1–1440.
	Time *int64 `json:"Time,omitempty" name:"Time"`
}

type NoticeReceiver struct {

	// Recipient type. Valid values:
	// <br><li> `Uin`: user ID
	// <br><li> `Group`: user group ID
	// Currently, other recipient types are not supported.
	ReceiverType *string `json:"ReceiverType,omitempty" name:"ReceiverType"`

	// Recipient
	ReceiverIds []*int64 `json:"ReceiverIds,omitempty" name:"ReceiverIds"`

	// Notification method
	// <br><li> `Email`: email
	// <br><li> `Sms`: SMS
	// <br><li> `WeChat`: WeChat
	// <br><li> `Phone`: phone
	ReceiverChannels []*string `json:"ReceiverChannels,omitempty" name:"ReceiverChannels"`

	// Start time for allowed message receipt
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time for allowed message receipt
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Index
	Index *int64 `json:"Index,omitempty" name:"Index"`
}

type PartitionInfo struct {

	// Partition ID
	PartitionId *int64 `json:"PartitionId,omitempty" name:"PartitionId"`

	// Partition status. Valid values: `readwrite`, `readonly`
	Status *string `json:"Status,omitempty" name:"Status"`

	// Partition hash start key
	InclusiveBeginKey *string `json:"InclusiveBeginKey,omitempty" name:"InclusiveBeginKey"`

	// Partition hash end key
	ExclusiveEndKey *string `json:"ExclusiveEndKey,omitempty" name:"ExclusiveEndKey"`

	// Partition creation time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Last modified of read-only partition
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	LastWriteTime *string `json:"LastWriteTime,omitempty" name:"LastWriteTime"`
}

type RetryShipperTaskRequest struct {
	*tchttp.BaseRequest

	// Shipping rule ID
	ShipperId *string `json:"ShipperId,omitempty" name:"ShipperId"`

	// Shipping task ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *RetryShipperTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RetryShipperTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ShipperId")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RetryShipperTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type RetryShipperTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RetryShipperTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RetryShipperTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RuleInfo struct {

	// Full-Text index configuration
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	FullText *FullTextInfo `json:"FullText,omitempty" name:"FullText"`

	// Key-Value index configuration
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	KeyValue *RuleKeyValueInfo `json:"KeyValue,omitempty" name:"KeyValue"`

	// Metafield index configuration
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Tag *RuleTagInfo `json:"Tag,omitempty" name:"Tag"`
}

type RuleKeyValueInfo struct {

	// Case sensitivity
	CaseSensitive *bool `json:"CaseSensitive,omitempty" name:"CaseSensitive"`

	// Key-Value pair information of the index to be created. Up to 100 key-value pairs can be configured.
	KeyValues []*KeyValueInfo `json:"KeyValues,omitempty" name:"KeyValues"`
}

type RuleTagInfo struct {

	// Case sensitivity
	CaseSensitive *bool `json:"CaseSensitive,omitempty" name:"CaseSensitive"`

	// Field information in tag index configuration
	KeyValues []*KeyValueInfo `json:"KeyValues,omitempty" name:"KeyValues"`
}

type SearchLogRequest struct {
	*tchttp.BaseRequest

	// Log topic ID to be queried
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// Start time of the log to be queried, which is a Unix timestamp in milliseconds
	From *int64 `json:"From,omitempty" name:"From"`

	// End time of the log to be queried, which is a Unix timestamp in milliseconds
	To *int64 `json:"To,omitempty" name:"To"`

	// Query statement. Maximum length: 4096
	Query *string `json:"Query,omitempty" name:"Query"`

	// Number of raw logs returned in a single query. Maximum value: 100. If the query statement (Query) contains an SQL query, you need to specify the number of SQL query results in `Query`. For more information, please visit https://intl.cloud.tencent.com/document/product/614/58977?from_cn_redirect=1
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// This field is used to load more logs. Pass through the last `Context` value returned to get more log content. It will expire after 1 hour.
	Context *string `json:"Context,omitempty" name:"Context"`

	// Order of the logs sorted by time returned by the log API. Valid values: `asc`: ascending; `desc`: descending. Default value: `desc`
	Sort *string `json:"Sort,omitempty" name:"Sort"`

	// If the value is `true`, the new search method will be used, and the response parameters `AnalysisRecords` and `Columns` will be valid. If the value is `false`, the old search method will be used, and `AnalysisResults` and `ColNames` will be valid.
	UseNewAnalysis *bool `json:"UseNewAnalysis,omitempty" name:"UseNewAnalysis"`
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
	delete(f, "TopicId")
	delete(f, "From")
	delete(f, "To")
	delete(f, "Query")
	delete(f, "Limit")
	delete(f, "Context")
	delete(f, "Sort")
	delete(f, "UseNewAnalysis")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SearchLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type SearchLogResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// `Context` for loading subsequent content. It will expire after 1 hour.
		Context *string `json:"Context,omitempty" name:"Context"`

		// Whether to return all raw log query results. This parameter is meaningless if the query statement (Query) contains an SQL query.
		ListOver *bool `json:"ListOver,omitempty" name:"ListOver"`

		// Whether the return is the analysis result
		Analysis *bool `json:"Analysis,omitempty" name:"Analysis"`

		// If `Analysis` is `true`, column name of the analysis result will be returned; otherwise, empty content will be returned.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
		ColNames []*string `json:"ColNames,omitempty" name:"ColNames"`

		// Log query result. If `Analysis` is `True`, `null` may be returned
	// Note: this field may return `null`, indicating that no valid values can be obtained.
		Results []*LogInfo `json:"Results,omitempty" name:"Results"`

		// Log analysis result. If `Analysis` is `False`, `null` may be returned
	// Note: this field may return `null`, indicating that no valid values can be obtained.
		AnalysisResults []*LogItems `json:"AnalysisResults,omitempty" name:"AnalysisResults"`

		// New log analysis result, which will be valid if `UseNewAnalysis` is `true`
	// Note: this field may return `null`, indicating that no valid values can be obtained.
		AnalysisRecords []*string `json:"AnalysisRecords,omitempty" name:"AnalysisRecords"`

		// Column attribute of log analysis, which will be valid if `UseNewAnalysis` is `true`
	// Note: this field may return `null`, indicating that no valid values can be obtained.
		Columns []*Column `json:"Columns,omitempty" name:"Columns"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type ShipperInfo struct {

	// Shipping rule ID
	ShipperId *string `json:"ShipperId,omitempty" name:"ShipperId"`

	// Log topic ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// Bucket address shipped to
	Bucket *string `json:"Bucket,omitempty" name:"Bucket"`

	// Shipping prefix directory
	Prefix *string `json:"Prefix,omitempty" name:"Prefix"`

	// Shipping rule name
	ShipperName *string `json:"ShipperName,omitempty" name:"ShipperName"`

	// Shipping time interval in seconds
	Interval *uint64 `json:"Interval,omitempty" name:"Interval"`

	// Maximum size of shipped file in MB
	MaxSize *uint64 `json:"MaxSize,omitempty" name:"MaxSize"`

	// Whether it takes effect
	Status *bool `json:"Status,omitempty" name:"Status"`

	// Filter rule for shipped log
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	FilterRules []*FilterRuleInfo `json:"FilterRules,omitempty" name:"FilterRules"`

	// Partition rule of shipped log, which can be represented in `strftime` time format
	Partition *string `json:"Partition,omitempty" name:"Partition"`

	// Compression configuration of shipped log
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Compress *CompressInfo `json:"Compress,omitempty" name:"Compress"`

	// Format configuration of shipped log content
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Content *ContentInfo `json:"Content,omitempty" name:"Content"`

	// Creation time of shipped log
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type ShipperTaskInfo struct {

	// Shipping task ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// Shipping information ID
	ShipperId *string `json:"ShipperId,omitempty" name:"ShipperId"`

	// Log topic ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// Start timestamp of the current batch of shipped logs in milliseconds
	RangeStart *int64 `json:"RangeStart,omitempty" name:"RangeStart"`

	// End timestamp of the current batch of shipped logs in milliseconds
	RangeEnd *int64 `json:"RangeEnd,omitempty" name:"RangeEnd"`

	// Start timestamp of the current shipping task in milliseconds
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// End timestamp of the current shipping task in milliseconds
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// Result of the current shipping task. Valid values: `success`, `running`, `failed`
	Status *string `json:"Status,omitempty" name:"Status"`

	// Result details
	Message *string `json:"Message,omitempty" name:"Message"`
}

type SplitPartitionRequest struct {
	*tchttp.BaseRequest

	// Log topic ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// ID of the partition to be split
	PartitionId *int64 `json:"PartitionId,omitempty" name:"PartitionId"`

	// Partition split hash key position, which is meaningful only if `Number=2` is set
	SplitKey *string `json:"SplitKey,omitempty" name:"SplitKey"`

	// Number of partitions to split into, which is optional. Default value: 2
	Number *int64 `json:"Number,omitempty" name:"Number"`
}

func (r *SplitPartitionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SplitPartitionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	delete(f, "PartitionId")
	delete(f, "SplitKey")
	delete(f, "Number")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SplitPartitionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type SplitPartitionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Split result set
		Partitions []*PartitionInfo `json:"Partitions,omitempty" name:"Partitions"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SplitPartitionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SplitPartitionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Tag struct {

	// Tag key
	Key *string `json:"Key,omitempty" name:"Key"`

	// Tag value
	Value *string `json:"Value,omitempty" name:"Value"`
}

type TopicInfo struct {

	// Logset ID
	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`

	// Log topic ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// Log topic name
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// Number of topic partitions
	PartitionCount *int64 `json:"PartitionCount,omitempty" name:"PartitionCount"`

	// Whether index is enabled
	Index *bool `json:"Index,omitempty" name:"Index"`

	// Creation time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Whether collection is enabled in the log topic
	Status *bool `json:"Status,omitempty" name:"Status"`

	// Information of tags bound to log topic
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// Whether automatic split is enabled for this topic
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	AutoSplit *bool `json:"AutoSplit,omitempty" name:"AutoSplit"`

	// Maximum number of partitions to split into for this topic if automatic split is enabled
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	MaxSplitPartitions *int64 `json:"MaxSplitPartitions,omitempty" name:"MaxSplitPartitions"`

	// Log topic storage class
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	StorageType *string `json:"StorageType,omitempty" name:"StorageType"`

	// Lifecycle in days
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Period *int64 `json:"Period,omitempty" name:"Period"`
}

type UploadLogRequest struct {
	*tchttp.BaseRequest
}

func (r *UploadLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UploadLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type UploadLogResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UploadLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ValueInfo struct {

	// Field type. Valid values: `long`, `text`, `double`
	Type *string `json:"Type,omitempty" name:"Type"`

	// Field delimiter, which is meaningful only if the field type is `text`. Each character in the entered string represents a delimiter.
	Tokenizer *string `json:"Tokenizer,omitempty" name:"Tokenizer"`

	// Whether the analysis feature is enabled for the field
	SqlFlag *bool `json:"SqlFlag,omitempty" name:"SqlFlag"`

	// Whether Chinese characters are contained
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ContainZH *bool `json:"ContainZH,omitempty" name:"ContainZH"`
}

type WebCallback struct {

	// Callback address
	Url *string `json:"Url,omitempty" name:"Url"`

	// Callback type. Valid values:
	// <br><li> WeCom
	// <br><li> Http
	CallbackType *string `json:"CallbackType,omitempty" name:"CallbackType"`

	// Callback method. Valid values:
	// <br><li> POST
	// <br><li> PUT
	// Default value: `POST`. This parameter is required if `CallbackType` is `Http`.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Method *string `json:"Method,omitempty" name:"Method"`

	// Request header
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Headers []*string `json:"Headers,omitempty" name:"Headers"`

	// Request content, which is required when `CallbackType` is `Http`.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Body *string `json:"Body,omitempty" name:"Body"`

	// Number
	Index *int64 `json:"Index,omitempty" name:"Index"`
}
