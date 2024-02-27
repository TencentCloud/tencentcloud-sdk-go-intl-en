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

package v20211206

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

type CheckStep struct {
	// Step number
	// Note: This field may return null, indicating that no valid values can be obtained.
	StepNo *uint64 `json:"StepNo,omitnil,omitempty" name:"StepNo"`

	// Step ID such as `ConnectDBCheck`, `VersionCheck`, and `SrcPrivilegeCheck`. The specific check items are subject to source and target instances.
	// Note: This field may return null, indicating that no valid values can be obtained.
	StepId *string `json:"StepId,omitnil,omitempty" name:"StepId"`

	// Step name
	// Note: This field may return null, indicating that no valid values can be obtained.
	StepName *string `json:"StepName,omitnil,omitempty" name:"StepName"`

	// Result of this check step. Valid values: `pass`, `failed`, `notStarted`, `blocked`, `warning`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	StepStatus *string `json:"StepStatus,omitnil,omitempty" name:"StepStatus"`

	// Error message in this check step
	// Note: This field may return null, indicating that no valid values can be obtained.
	StepMessage *string `json:"StepMessage,omitnil,omitempty" name:"StepMessage"`

	// Specific check item in this check step
	// Note: This field may return null, indicating that no valid values can be obtained.
	DetailCheckItems []*DetailCheckItem `json:"DetailCheckItems,omitnil,omitempty" name:"DetailCheckItems"`

	// Whether this step was skipped
	// Note: This field may return null, indicating that no valid values can be obtained.
	HasSkipped *bool `json:"HasSkipped,omitnil,omitempty" name:"HasSkipped"`
}

type CheckStepInfo struct {
	// Task start time
	// Note: This field may return null, indicating that no valid values can be obtained.
	StartAt *string `json:"StartAt,omitnil,omitempty" name:"StartAt"`

	// Task end time
	// Note: This field may return null, indicating that no valid values can be obtained.
	EndAt *string `json:"EndAt,omitnil,omitempty" name:"EndAt"`

	// Task step information
	// Note: This field may return null, indicating that no valid values can be obtained.
	Progress *ProcessProgress `json:"Progress,omitnil,omitempty" name:"Progress"`
}

type CompareAbstractInfo struct {
	// Configuration parameters of the check task
	// Note: This field may return null, indicating that no valid values can be obtained.
	Options *CompareOptions `json:"Options,omitnil,omitempty" name:"Options"`

	// Consistency check objects
	// Note: This field may return null, indicating that no valid values can be obtained.
	Objects *CompareObject `json:"Objects,omitnil,omitempty" name:"Objects"`

	// Comparison conclusion. Valid values: `same`, `different`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Conclusion *string `json:"Conclusion,omitnil,omitempty" name:"Conclusion"`

	// Task status. Valid values: `success`, `failed`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Total number of tables
	// Note: This field may return null, indicating that no valid values can be obtained.
	TotalTables *uint64 `json:"TotalTables,omitnil,omitempty" name:"TotalTables"`

	// Number of checked tables
	// Note: This field may return null, indicating that no valid values can be obtained.
	CheckedTables *uint64 `json:"CheckedTables,omitnil,omitempty" name:"CheckedTables"`

	// Number of inconsistent tables
	// Note: This field may return null, indicating that no valid values can be obtained.
	DifferentTables *uint64 `json:"DifferentTables,omitnil,omitempty" name:"DifferentTables"`

	// Number of skipped tables
	// Note: This field may return null, indicating that no valid values can be obtained.
	SkippedTables *uint64 `json:"SkippedTables,omitnil,omitempty" name:"SkippedTables"`

	// The estimated number of tables
	// Note: This field may return null, indicating that no valid values can be obtained.
	NearlyTableCount *uint64 `json:"NearlyTableCount,omitnil,omitempty" name:"NearlyTableCount"`

	// Number of inconsistent data rows
	// Note: This field may return null, indicating that no valid values can be obtained.
	DifferentRows *uint64 `json:"DifferentRows,omitnil,omitempty" name:"DifferentRows"`

	// Source database row count, which takes effect only when the comparison type is **Row count comparison**.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SrcSampleRows *uint64 `json:"SrcSampleRows,omitnil,omitempty" name:"SrcSampleRows"`

	// Target database row count, which takes effect only when the comparison type is **Row count comparison**.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DstSampleRows *uint64 `json:"DstSampleRows,omitnil,omitempty" name:"DstSampleRows"`

	// Start time
	// Note: This field may return null, indicating that no valid values can be obtained.
	StartedAt *string `json:"StartedAt,omitnil,omitempty" name:"StartedAt"`

	// End time
	// Note: This field may return null, indicating that no valid values can be obtained.
	FinishedAt *string `json:"FinishedAt,omitnil,omitempty" name:"FinishedAt"`
}

type CompareDetailInfo struct {
	// Details of inconsistent tables
	// Note: This field may return null, indicating that no valid values can be obtained.
	Difference *DifferenceDetail `json:"Difference,omitnil,omitempty" name:"Difference"`

	// Details of skipped tables
	// Note: This field may return null, indicating that no valid values can be obtained.
	Skipped *SkippedDetail `json:"Skipped,omitnil,omitempty" name:"Skipped"`
}

type CompareObject struct {
	// Data comparison object mode (`all`: Entire instance; `partial`: Some objects)
	// Note: This field may return null, indicating that no valid values can be obtained.
	ObjectMode *string `json:"ObjectMode,omitnil,omitempty" name:"ObjectMode"`

	// Object list
	// Note: This field may return null, indicating that no valid values can be obtained.
	ObjectItems []*CompareObjectItem `json:"ObjectItems,omitnil,omitempty" name:"ObjectItems"`

	// Advanced object type (`account`: Account; `index`: Index; `shardkey`: Shard key, which may be adjusted later; `schema`: Database/table structure)
	// Note: This field may return null, indicating that no valid values can be obtained.
	AdvancedObjects []*string `json:"AdvancedObjects,omitnil,omitempty" name:"AdvancedObjects"`
}

type CompareObjectItem struct {
	// Database name
	// Note: This field may return null, indicating that no valid values can be obtained.
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// Database selection mode. Valid values: `all`, `partial`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DbMode *string `json:"DbMode,omitnil,omitempty" name:"DbMode"`

	// Schema name
	// Note: This field may return null, indicating that no valid values can be obtained.
	SchemaName *string `json:"SchemaName,omitnil,omitempty" name:"SchemaName"`

	// Schema selection mode. Valid values: `all`, `partial`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TableMode *string `json:"TableMode,omitnil,omitempty" name:"TableMode"`

	// Table configuration for data consistency check, which is required if `TableMode` is `partial`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Tables []*CompareTableItem `json:"Tables,omitnil,omitempty" name:"Tables"`

	// View selection mode. Valid values: `all`, `partial`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ViewMode *string `json:"ViewMode,omitnil,omitempty" name:"ViewMode"`

	// View configuration for data consistency check, which is required if `ViewMode` is `partial`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Views []*CompareViewItem `json:"Views,omitnil,omitempty" name:"Views"`
}

type CompareOptions struct {
	// Comparison type: (`dataCheck`: Full data comparison; `sampleDataCheck`: Sampling data comparison; `rowsCount`: Row count comparison)
	// Note: This field may return null, indicating that no valid values can be obtained.
	Method *string `json:"Method,omitnil,omitempty" name:"Method"`

	// Sampling rate. Value range: 0-100%.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SampleRate *int64 `json:"SampleRate,omitnil,omitempty" name:"SampleRate"`

	// The number of threads, which defaults to 1. Value range: 1-5.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ThreadCount *int64 `json:"ThreadCount,omitnil,omitempty" name:"ThreadCount"`
}

type CompareTableItem struct {
	// Table name
	// Note: This field may return null, indicating that no valid values can be obtained.
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`
}

type CompareTaskInfo struct {
	// Data consistency check task ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	CompareTaskId *string `json:"CompareTaskId,omitnil,omitempty" name:"CompareTaskId"`

	// Data consistency check result. Valid values: `unstart` (the task is not started); `running` (the task is running); `canceled` (the task is stopped); `failed` (the task failed); `inconsistent` (the data is inconsistent); `consistent` (the data is consistent); `notexist` (the task does not exist).
	// Note: This field may return null, indicating that no valid values can be obtained.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

type CompareTaskItem struct {
	// Task ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// Data consistency check task ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	CompareTaskId *string `json:"CompareTaskId,omitnil,omitempty" name:"CompareTaskId"`

	// Data consistency check task name
	// Note: This field may return null, indicating that no valid values can be obtained.
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// Data consistency check task status. Valid values: `created`, `readyRun`, `running`, `success`, `stopping`, `failed`, `canceled`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Data consistency check task configuration
	// Note: This field may return null, indicating that no valid values can be obtained.
	Config *CompareObject `json:"Config,omitnil,omitempty" name:"Config"`

	// Check details of the data consistency check task
	// Note: This field may return null, indicating that no valid values can be obtained.
	CheckProcess *ProcessProgress `json:"CheckProcess,omitnil,omitempty" name:"CheckProcess"`

	// Running details of the data consistency check task
	// Note: This field may return null, indicating that no valid values can be obtained.
	CompareProcess *ProcessProgress `json:"CompareProcess,omitnil,omitempty" name:"CompareProcess"`

	// Comparison result. Valid values: `same`, `different`, `skipAll`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Conclusion *string `json:"Conclusion,omitnil,omitempty" name:"Conclusion"`

	// Task creation time
	// Note: This field may return null, indicating that no valid values can be obtained.
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// Task start time
	// Note: This field may return null, indicating that no valid values can be obtained.
	StartedAt *string `json:"StartedAt,omitnil,omitempty" name:"StartedAt"`

	// Comparison end time
	// Note: This field may return null, indicating that no valid values can be obtained.
	FinishedAt *string `json:"FinishedAt,omitnil,omitempty" name:"FinishedAt"`

	// Comparison type: (`dataCheck`: Full data comparison; `sampleDataCheck`: Sampling data comparison; `rowsCount`: Row count comparison)
	// Note: This field may return null, indicating that no valid values can be obtained.
	Method *string `json:"Method,omitnil,omitempty" name:"Method"`

	// Configuration information of the comparison task
	// Note: This field may return null, indicating that no valid values can be obtained.
	Options *CompareOptions `json:"Options,omitnil,omitempty" name:"Options"`

	// Consistency check prompt message
	// Note: This field may return null, indicating that no valid values can be obtained.
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`
}

type CompareViewItem struct {
	// View name
	// Note: This field may return null, indicating that no valid values can be obtained.
	ViewName *string `json:"ViewName,omitnil,omitempty" name:"ViewName"`
}

// Predefined struct for user
type CompleteMigrateJobRequestParams struct {
	// Data migration task ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// The way to complete the task, which is supported only for legacy MySQL migration tasks. Valid values: `waitForSync` (wait for the source-replica lag to become 0 before stopping); `immediately` (complete immediately without waiting for source-replica sync). Default value: `waitForSync`.
	CompleteMode *string `json:"CompleteMode,omitnil,omitempty" name:"CompleteMode"`
}

type CompleteMigrateJobRequest struct {
	*tchttp.BaseRequest
	
	// Data migration task ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// The way to complete the task, which is supported only for legacy MySQL migration tasks. Valid values: `waitForSync` (wait for the source-replica lag to become 0 before stopping); `immediately` (complete immediately without waiting for source-replica sync). Default value: `waitForSync`.
	CompleteMode *string `json:"CompleteMode,omitnil,omitempty" name:"CompleteMode"`
}

func (r *CompleteMigrateJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CompleteMigrateJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	delete(f, "CompleteMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CompleteMigrateJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CompleteMigrateJobResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CompleteMigrateJobResponse struct {
	*tchttp.BaseResponse
	Response *CompleteMigrateJobResponseParams `json:"Response"`
}

func (r *CompleteMigrateJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CompleteMigrateJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ConfigureSyncJobRequestParams struct {
	// Sync task instance ID in the format of `sync-werwfs23`, which is used to identify a sync task.
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// Source database access type. Valid values: `cdb` (database); `cvm` (self-build on CVM); `vpc` (VPC); `extranet` (public network); `vpncloud` (VPN access); `dcg` (Direct Connect); `ccn` (CCN); `intranet` (intranet). Note that the valid values are subject to the current link.
	SrcAccessType *string `json:"SrcAccessType,omitnil,omitempty" name:"SrcAccessType"`

	// Target database access type. Valid values: `cdb` (database); `cvm` (self-build on CVM); `vpc` (VPC); `extranet` (public network); `vpncloud` (VPN access); `dcg` (Direct Connect); `ccn` (CCN); `intranet` (intranet); `ckafka` (CKafka instance). Note that the valid values are subject to the current link.
	DstAccessType *string `json:"DstAccessType,omitnil,omitempty" name:"DstAccessType"`

	// Information of synced database/table objects
	Objects *Objects `json:"Objects,omitnil,omitempty" name:"Objects"`

	// Sync task name
	JobName *string `json:"JobName,omitnil,omitempty" name:"JobName"`

	// Enumerated values: `liteMode`: Lite mode; `fullMode`: Standard mode
	JobMode *string `json:"JobMode,omitnil,omitempty" name:"JobMode"`

	// Running mode. Valid values: `Immediate`, `Timed`. Default value: `Immediate`.
	RunMode *string `json:"RunMode,omitnil,omitempty" name:"RunMode"`

	// Expected start time in the format of "2006-01-02 15:04:05", which is required if `RunMode` is `Timed`.
	ExpectRunTime *string `json:"ExpectRunTime,omitnil,omitempty" name:"ExpectRunTime"`

	// Source database information. This parameter only applies to single-node databases, and `SrcNodeType` must be `single`.
	SrcInfo *Endpoint `json:"SrcInfo,omitnil,omitempty" name:"SrcInfo"`

	// Source database information. This parameter is valid for multi-node databases, and the value of `SrcNodeType` must be `cluster`.
	SrcInfos *SyncDBEndpointInfos `json:"SrcInfos,omitnil,omitempty" name:"SrcInfos"`

	// Enumerated values: `single` (for single-node source database), `cluster` (for multi-node source database).
	SrcNodeType *string `json:"SrcNodeType,omitnil,omitempty" name:"SrcNodeType"`

	// Target database information. This parameter is used by single-node databases.
	DstInfo *Endpoint `json:"DstInfo,omitnil,omitempty" name:"DstInfo"`

	// Target database information. This parameter is valid for multi-node databases, and the value of `DstNodeType` must be `cluster`.
	DstInfos *SyncDBEndpointInfos `json:"DstInfos,omitnil,omitempty" name:"DstInfos"`

	// Enumerated values: `single` (for single-node target database), `cluster` (for multi-node target database).
	DstNodeType *string `json:"DstNodeType,omitnil,omitempty" name:"DstNodeType"`

	// Sync task options. The `RateLimitOption` option cannot take effect currently. To modify the speed limit settings, use the `ModifySyncRateLimit` API.
	Options *Options `json:"Options,omitnil,omitempty" name:"Options"`

	// Automatic retry time, which can be set to 5-720 minutes. 0 indicates that retry is disabled.
	AutoRetryTimeRangeMinutes *int64 `json:"AutoRetryTimeRangeMinutes,omitnil,omitempty" name:"AutoRetryTimeRangeMinutes"`
}

type ConfigureSyncJobRequest struct {
	*tchttp.BaseRequest
	
	// Sync task instance ID in the format of `sync-werwfs23`, which is used to identify a sync task.
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// Source database access type. Valid values: `cdb` (database); `cvm` (self-build on CVM); `vpc` (VPC); `extranet` (public network); `vpncloud` (VPN access); `dcg` (Direct Connect); `ccn` (CCN); `intranet` (intranet). Note that the valid values are subject to the current link.
	SrcAccessType *string `json:"SrcAccessType,omitnil,omitempty" name:"SrcAccessType"`

	// Target database access type. Valid values: `cdb` (database); `cvm` (self-build on CVM); `vpc` (VPC); `extranet` (public network); `vpncloud` (VPN access); `dcg` (Direct Connect); `ccn` (CCN); `intranet` (intranet); `ckafka` (CKafka instance). Note that the valid values are subject to the current link.
	DstAccessType *string `json:"DstAccessType,omitnil,omitempty" name:"DstAccessType"`

	// Information of synced database/table objects
	Objects *Objects `json:"Objects,omitnil,omitempty" name:"Objects"`

	// Sync task name
	JobName *string `json:"JobName,omitnil,omitempty" name:"JobName"`

	// Enumerated values: `liteMode`: Lite mode; `fullMode`: Standard mode
	JobMode *string `json:"JobMode,omitnil,omitempty" name:"JobMode"`

	// Running mode. Valid values: `Immediate`, `Timed`. Default value: `Immediate`.
	RunMode *string `json:"RunMode,omitnil,omitempty" name:"RunMode"`

	// Expected start time in the format of "2006-01-02 15:04:05", which is required if `RunMode` is `Timed`.
	ExpectRunTime *string `json:"ExpectRunTime,omitnil,omitempty" name:"ExpectRunTime"`

	// Source database information. This parameter only applies to single-node databases, and `SrcNodeType` must be `single`.
	SrcInfo *Endpoint `json:"SrcInfo,omitnil,omitempty" name:"SrcInfo"`

	// Source database information. This parameter is valid for multi-node databases, and the value of `SrcNodeType` must be `cluster`.
	SrcInfos *SyncDBEndpointInfos `json:"SrcInfos,omitnil,omitempty" name:"SrcInfos"`

	// Enumerated values: `single` (for single-node source database), `cluster` (for multi-node source database).
	SrcNodeType *string `json:"SrcNodeType,omitnil,omitempty" name:"SrcNodeType"`

	// Target database information. This parameter is used by single-node databases.
	DstInfo *Endpoint `json:"DstInfo,omitnil,omitempty" name:"DstInfo"`

	// Target database information. This parameter is valid for multi-node databases, and the value of `DstNodeType` must be `cluster`.
	DstInfos *SyncDBEndpointInfos `json:"DstInfos,omitnil,omitempty" name:"DstInfos"`

	// Enumerated values: `single` (for single-node target database), `cluster` (for multi-node target database).
	DstNodeType *string `json:"DstNodeType,omitnil,omitempty" name:"DstNodeType"`

	// Sync task options. The `RateLimitOption` option cannot take effect currently. To modify the speed limit settings, use the `ModifySyncRateLimit` API.
	Options *Options `json:"Options,omitnil,omitempty" name:"Options"`

	// Automatic retry time, which can be set to 5-720 minutes. 0 indicates that retry is disabled.
	AutoRetryTimeRangeMinutes *int64 `json:"AutoRetryTimeRangeMinutes,omitnil,omitempty" name:"AutoRetryTimeRangeMinutes"`
}

func (r *ConfigureSyncJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ConfigureSyncJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	delete(f, "SrcAccessType")
	delete(f, "DstAccessType")
	delete(f, "Objects")
	delete(f, "JobName")
	delete(f, "JobMode")
	delete(f, "RunMode")
	delete(f, "ExpectRunTime")
	delete(f, "SrcInfo")
	delete(f, "SrcInfos")
	delete(f, "SrcNodeType")
	delete(f, "DstInfo")
	delete(f, "DstInfos")
	delete(f, "DstNodeType")
	delete(f, "Options")
	delete(f, "AutoRetryTimeRangeMinutes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ConfigureSyncJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ConfigureSyncJobResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ConfigureSyncJobResponse struct {
	*tchttp.BaseResponse
	Response *ConfigureSyncJobResponseParams `json:"Response"`
}

func (r *ConfigureSyncJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ConfigureSyncJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConflictHandleOption struct {
	// Conditionally overwritten column
	// Note: This field may return null, indicating that no valid values can be obtained.
	ConditionColumn *string `json:"ConditionColumn,omitnil,omitempty" name:"ConditionColumn"`

	// Conditional overwrite operation
	// Note: This field may return null, indicating that no valid values can be obtained.
	ConditionOperator *string `json:"ConditionOperator,omitnil,omitempty" name:"ConditionOperator"`

	// Conditional overwrite priority configuration
	// Note: This field may return null, indicating that no valid values can be obtained.
	ConditionOrderInSrcAndDst *string `json:"ConditionOrderInSrcAndDst,omitnil,omitempty" name:"ConditionOrderInSrcAndDst"`
}

type ConsistencyOption struct {
	// Data consistency check type. Valid values: `full`, `noCheck`, `notConfigured`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`
}

// Predefined struct for user
type ContinueMigrateJobRequestParams struct {
	// Data migration task ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

type ContinueMigrateJobRequest struct {
	*tchttp.BaseRequest
	
	// Data migration task ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

func (r *ContinueMigrateJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ContinueMigrateJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ContinueMigrateJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ContinueMigrateJobResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ContinueMigrateJobResponse struct {
	*tchttp.BaseResponse
	Response *ContinueMigrateJobResponseParams `json:"Response"`
}

func (r *ContinueMigrateJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ContinueMigrateJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ContinueSyncJobRequestParams struct {
	// Sync task ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

type ContinueSyncJobRequest struct {
	*tchttp.BaseRequest
	
	// Sync task ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

func (r *ContinueSyncJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ContinueSyncJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ContinueSyncJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ContinueSyncJobResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ContinueSyncJobResponse struct {
	*tchttp.BaseResponse
	Response *ContinueSyncJobResponseParams `json:"Response"`
}

func (r *ContinueSyncJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ContinueSyncJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCheckSyncJobRequestParams struct {
	// Sync task ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

type CreateCheckSyncJobRequest struct {
	*tchttp.BaseRequest
	
	// Sync task ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

func (r *CreateCheckSyncJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCheckSyncJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCheckSyncJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCheckSyncJobResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCheckSyncJobResponse struct {
	*tchttp.BaseResponse
	Response *CreateCheckSyncJobResponseParams `json:"Response"`
}

func (r *CreateCheckSyncJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCheckSyncJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCompareTaskRequestParams struct {
	// Task ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// Data consistency check task name. If this parameter is left empty, the value of `CompareTaskId` will be assigned to it.
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// Data comparison object mode. Valid values: `sameAsMigrate` (all migration objects); `custom` (custom mode). Default value: `sameAsMigrate`.
	ObjectMode *string `json:"ObjectMode,omitnil,omitempty" name:"ObjectMode"`

	// Configuration of the data consistency check object
	Objects *CompareObject `json:"Objects,omitnil,omitempty" name:"Objects"`

	// Consistency check options
	Options *CompareOptions `json:"Options,omitnil,omitempty" name:"Options"`
}

type CreateCompareTaskRequest struct {
	*tchttp.BaseRequest
	
	// Task ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// Data consistency check task name. If this parameter is left empty, the value of `CompareTaskId` will be assigned to it.
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// Data comparison object mode. Valid values: `sameAsMigrate` (all migration objects); `custom` (custom mode). Default value: `sameAsMigrate`.
	ObjectMode *string `json:"ObjectMode,omitnil,omitempty" name:"ObjectMode"`

	// Configuration of the data consistency check object
	Objects *CompareObject `json:"Objects,omitnil,omitempty" name:"Objects"`

	// Consistency check options
	Options *CompareOptions `json:"Options,omitnil,omitempty" name:"Options"`
}

func (r *CreateCompareTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCompareTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	delete(f, "TaskName")
	delete(f, "ObjectMode")
	delete(f, "Objects")
	delete(f, "Options")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCompareTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCompareTaskResponseParams struct {
	// Data consistency check task ID in the format of `dts-8yv4w2i1-cmp-37skmii9`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CompareTaskId *string `json:"CompareTaskId,omitnil,omitempty" name:"CompareTaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCompareTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateCompareTaskResponseParams `json:"Response"`
}

func (r *CreateCompareTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCompareTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMigrateCheckJobRequestParams struct {
	// Data migration task ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

type CreateMigrateCheckJobRequest struct {
	*tchttp.BaseRequest
	
	// Data migration task ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

func (r *CreateMigrateCheckJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMigrateCheckJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateMigrateCheckJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMigrateCheckJobResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateMigrateCheckJobResponse struct {
	*tchttp.BaseResponse
	Response *CreateMigrateCheckJobResponseParams `json:"Response"`
}

func (r *CreateMigrateCheckJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMigrateCheckJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMigrationServiceRequestParams struct {
	// Source database type. Valid values: `mysql`, `redis`, `percona`, `mongodb`, `postgresql`, `sqlserver`, `mariadb`, and `cynosdbmysql`.
	SrcDatabaseType *string `json:"SrcDatabaseType,omitnil,omitempty" name:"SrcDatabaseType"`

	// Target database type. Valid values: `mysql`, `redis`, `percona`, `mongodb` ,`postgresql`, `sqlserver`, `mariadb`, and `cynosdbmysql`.
	DstDatabaseType *string `json:"DstDatabaseType,omitnil,omitempty" name:"DstDatabaseType"`

	// Source instance region, such as `ap-guangzhou`.
	SrcRegion *string `json:"SrcRegion,omitnil,omitempty" name:"SrcRegion"`

	// Target instance region, such as `ap-guangzhou`. Note that it must be the same as the API request region.
	DstRegion *string `json:"DstRegion,omitnil,omitempty" name:"DstRegion"`

	// Instance specification. Valid values: `small`, `medium`, `large`, `xlarge`, `2xlarge`.
	InstanceClass *string `json:"InstanceClass,omitnil,omitempty" name:"InstanceClass"`

	// Quantity. Value range: [1,15]. Default value: `1`.
	Count *uint64 `json:"Count,omitnil,omitempty" name:"Count"`

	// Migration task name, which can contain up to 128 characters.
	JobName *string `json:"JobName,omitnil,omitempty" name:"JobName"`

	// Tag information
	Tags []*TagItem `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type CreateMigrationServiceRequest struct {
	*tchttp.BaseRequest
	
	// Source database type. Valid values: `mysql`, `redis`, `percona`, `mongodb`, `postgresql`, `sqlserver`, `mariadb`, and `cynosdbmysql`.
	SrcDatabaseType *string `json:"SrcDatabaseType,omitnil,omitempty" name:"SrcDatabaseType"`

	// Target database type. Valid values: `mysql`, `redis`, `percona`, `mongodb` ,`postgresql`, `sqlserver`, `mariadb`, and `cynosdbmysql`.
	DstDatabaseType *string `json:"DstDatabaseType,omitnil,omitempty" name:"DstDatabaseType"`

	// Source instance region, such as `ap-guangzhou`.
	SrcRegion *string `json:"SrcRegion,omitnil,omitempty" name:"SrcRegion"`

	// Target instance region, such as `ap-guangzhou`. Note that it must be the same as the API request region.
	DstRegion *string `json:"DstRegion,omitnil,omitempty" name:"DstRegion"`

	// Instance specification. Valid values: `small`, `medium`, `large`, `xlarge`, `2xlarge`.
	InstanceClass *string `json:"InstanceClass,omitnil,omitempty" name:"InstanceClass"`

	// Quantity. Value range: [1,15]. Default value: `1`.
	Count *uint64 `json:"Count,omitnil,omitempty" name:"Count"`

	// Migration task name, which can contain up to 128 characters.
	JobName *string `json:"JobName,omitnil,omitempty" name:"JobName"`

	// Tag information
	Tags []*TagItem `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *CreateMigrationServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMigrationServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SrcDatabaseType")
	delete(f, "DstDatabaseType")
	delete(f, "SrcRegion")
	delete(f, "DstRegion")
	delete(f, "InstanceClass")
	delete(f, "Count")
	delete(f, "JobName")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateMigrationServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMigrationServiceResponseParams struct {
	// The list of migration task IDs randomly generated in the format of `dts-c1f6rs21` after a successful order placement
	// Note: This field may return null, indicating that no valid values can be obtained.
	JobIds []*string `json:"JobIds,omitnil,omitempty" name:"JobIds"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateMigrationServiceResponse struct {
	*tchttp.BaseResponse
	Response *CreateMigrationServiceResponseParams `json:"Response"`
}

func (r *CreateMigrationServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMigrationServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateModifyCheckSyncJobRequestParams struct {
	// Sync task ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

type CreateModifyCheckSyncJobRequest struct {
	*tchttp.BaseRequest
	
	// Sync task ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

func (r *CreateModifyCheckSyncJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateModifyCheckSyncJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateModifyCheckSyncJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateModifyCheckSyncJobResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateModifyCheckSyncJobResponse struct {
	*tchttp.BaseResponse
	Response *CreateModifyCheckSyncJobResponseParams `json:"Response"`
}

func (r *CreateModifyCheckSyncJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateModifyCheckSyncJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSyncJobRequestParams struct {
	// Billing mode. Valid values: `PrePay` (monthly subscription); `PostPay` (pay-as-you-go). Currently, DTS at Tencent Cloud International is free of charge.
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// Source database type, such as `mysql`, `cynosdbmysql`, `tdapg`, `tdpg`, and `tdsqlmysql`.
	SrcDatabaseType *string `json:"SrcDatabaseType,omitnil,omitempty" name:"SrcDatabaseType"`

	// Source database region, such as `ap-guangzhou`.
	SrcRegion *string `json:"SrcRegion,omitnil,omitempty" name:"SrcRegion"`

	// Target database type, such as `mysql`, `cynosdbmysql`, `tdapg`, `tdpg`, `tdsqlmysql`, and `kafka`.
	DstDatabaseType *string `json:"DstDatabaseType,omitnil,omitempty" name:"DstDatabaseType"`

	// Target database region, such as `ap-guangzhou`.
	DstRegion *string `json:"DstRegion,omitnil,omitempty" name:"DstRegion"`

	// Sync task specification, such as `Standard`.
	Specification *string `json:"Specification,omitnil,omitempty" name:"Specification"`

	// Tag information
	Tags []*TagItem `json:"Tags,omitnil,omitempty" name:"Tags"`

	// The number of sync tasks purchased at a time. Value range: [1, 10]. Default value: `1`.
	Count *uint64 `json:"Count,omitnil,omitempty" name:"Count"`

	// Auto-renewal flag, which takes effect if `PayMode` is `PrePay`. Valid values: `1` (auto-renewal enabled); `0` (auto-renewal disabled). Default value: `0`.
	AutoRenew *uint64 `json:"AutoRenew,omitnil,omitempty" name:"AutoRenew"`

	// Sync link specification, such as `micro`, `small`, `medium`, and `large`. Default value: `medium`.
	InstanceClass *string `json:"InstanceClass,omitnil,omitempty" name:"InstanceClass"`

	// Sync task name
	JobName *string `json:"JobName,omitnil,omitempty" name:"JobName"`

	// ID of the existing task used to create a similar task
	ExistedJobId *string `json:"ExistedJobId,omitnil,omitempty" name:"ExistedJobId"`
}

type CreateSyncJobRequest struct {
	*tchttp.BaseRequest
	
	// Billing mode. Valid values: `PrePay` (monthly subscription); `PostPay` (pay-as-you-go). Currently, DTS at Tencent Cloud International is free of charge.
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// Source database type, such as `mysql`, `cynosdbmysql`, `tdapg`, `tdpg`, and `tdsqlmysql`.
	SrcDatabaseType *string `json:"SrcDatabaseType,omitnil,omitempty" name:"SrcDatabaseType"`

	// Source database region, such as `ap-guangzhou`.
	SrcRegion *string `json:"SrcRegion,omitnil,omitempty" name:"SrcRegion"`

	// Target database type, such as `mysql`, `cynosdbmysql`, `tdapg`, `tdpg`, `tdsqlmysql`, and `kafka`.
	DstDatabaseType *string `json:"DstDatabaseType,omitnil,omitempty" name:"DstDatabaseType"`

	// Target database region, such as `ap-guangzhou`.
	DstRegion *string `json:"DstRegion,omitnil,omitempty" name:"DstRegion"`

	// Sync task specification, such as `Standard`.
	Specification *string `json:"Specification,omitnil,omitempty" name:"Specification"`

	// Tag information
	Tags []*TagItem `json:"Tags,omitnil,omitempty" name:"Tags"`

	// The number of sync tasks purchased at a time. Value range: [1, 10]. Default value: `1`.
	Count *uint64 `json:"Count,omitnil,omitempty" name:"Count"`

	// Auto-renewal flag, which takes effect if `PayMode` is `PrePay`. Valid values: `1` (auto-renewal enabled); `0` (auto-renewal disabled). Default value: `0`.
	AutoRenew *uint64 `json:"AutoRenew,omitnil,omitempty" name:"AutoRenew"`

	// Sync link specification, such as `micro`, `small`, `medium`, and `large`. Default value: `medium`.
	InstanceClass *string `json:"InstanceClass,omitnil,omitempty" name:"InstanceClass"`

	// Sync task name
	JobName *string `json:"JobName,omitnil,omitempty" name:"JobName"`

	// ID of the existing task used to create a similar task
	ExistedJobId *string `json:"ExistedJobId,omitnil,omitempty" name:"ExistedJobId"`
}

func (r *CreateSyncJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSyncJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PayMode")
	delete(f, "SrcDatabaseType")
	delete(f, "SrcRegion")
	delete(f, "DstDatabaseType")
	delete(f, "DstRegion")
	delete(f, "Specification")
	delete(f, "Tags")
	delete(f, "Count")
	delete(f, "AutoRenew")
	delete(f, "InstanceClass")
	delete(f, "JobName")
	delete(f, "ExistedJobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSyncJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSyncJobResponseParams struct {
	// Sync task IDs
	JobIds []*string `json:"JobIds,omitnil,omitempty" name:"JobIds"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateSyncJobResponse struct {
	*tchttp.BaseResponse
	Response *CreateSyncJobResponseParams `json:"Response"`
}

func (r *CreateSyncJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSyncJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DBEndpointInfo struct {
	// Instance region
	// Note: This field may return null, indicating that no valid values can be obtained.
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Instances network access type. Valid values: `extranet` (public network); `ipv6` (public IPv6); `cvm` (self-build on CVM); `dcg` (Direct Connect); `vpncloud` (VPN access); `cdb` (database); `ccn` (CCN); `intranet` (intranet); `vpc` (VPC). Note that the valid values are subject to the current link.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AccessType *string `json:"AccessType,omitnil,omitempty" name:"AccessType"`

	// Database type, such as `mysql`, `redis`, `mongodb`, `postgresql`, `mariadb`, and `percona`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DatabaseType *string `json:"DatabaseType,omitnil,omitempty" name:"DatabaseType"`

	// Node type. Valid values: empty or `simple` (general node); `cluster` (cluster node).
	// Note: This field may return null, indicating that no valid values can be obtained.
	NodeType *string `json:"NodeType,omitnil,omitempty" name:"NodeType"`

	// Database information
	// Note: This field may return null, indicating that no valid values can be obtained.
	Info []*DBInfo `json:"Info,omitnil,omitempty" name:"Info"`

	// Instance service provider, such as "aliyun" and "others".
	// Note: This field may return null, indicating that no valid values can be obtained.
	Supplier *string `json:"Supplier,omitnil,omitempty" name:"Supplier"`

	// For MongoDB, you can define the following parameters: 	['AuthDatabase':'admin', 
	// 'AuthFlag': "1",	'AuthMechanism':"SCRAM-SHA-1"]
	// Note: This field may return null, indicating that no valid values can be obtained.
	ExtraAttr []*KeyValuePairOption `json:"ExtraAttr,omitnil,omitempty" name:"ExtraAttr"`

	// Network environment of the database. This parameter is required when `AccessType` is `ccn`. Valid values: `UserIDC` (user IDC), `TencentVPC` (Tencent Cloud VPC).
	// Note: This field may return null, indicating that no valid values can be obtained.
	DatabaseNetEnv *string `json:"DatabaseNetEnv,omitnil,omitempty" name:"DatabaseNetEnv"`
}

type DBInfo struct {
	// Node role in a distributed database, such as the mongos node in MongoDB.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Role *string `json:"Role,omitnil,omitempty" name:"Role"`

	// Kernel version, such as the different kernel versions of MariaDB.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DbKernel *string `json:"DbKernel,omitnil,omitempty" name:"DbKernel"`

	// Instance IP address, which is required for the following access types: public network, Direct Connect, VPN, CCN, intranet, and VPC.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// Instance port, which is required for the following access types: public network, self-build on CVM, Direct Connect, VPN, CCN, intranet, and VPC.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Port *uint64 `json:"Port,omitnil,omitempty" name:"Port"`

	// Instance username
	// Note: This field may return null, indicating that no valid values can be obtained.
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// Instance password
	// Note: This field may return null, indicating that no valid values can be obtained.
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// Short CVM instance ID in the format of `ins-olgl39y8`, which is required if the access type is `cvm`. It is the same as the instance ID displayed in the CVM console.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CvmInstanceId *string `json:"CvmInstanceId,omitnil,omitempty" name:"CvmInstanceId"`

	// VPN gateway ID in the format of `vpngw-9ghexg7q`, which is required if the access type is `vpncloud`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	UniqVpnGwId *string `json:"UniqVpnGwId,omitnil,omitempty" name:"UniqVpnGwId"`

	// Direct Connect gateway ID in the format of `dcg-0rxtqqxb`, which is required if the access type is `dcg`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	UniqDcgId *string `json:"UniqDcgId,omitnil,omitempty" name:"UniqDcgId"`

	// Database instance ID in the format of `cdb-powiqx8q`, which is required if the access type is `cdb`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// CCN instance ID such as `ccn-afp6kltc`
	// Note: This field may return null, indicating that no valid values can be obtained.
	CcnGwId *string `json:"CcnGwId,omitnil,omitempty" name:"CcnGwId"`

	// VPC ID in the format of `vpc-92jblxto`, which is required if the access type is `vpc`, `vpncloud`, `ccn`, or `dcg`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// ID of the subnet in the VPC in the format of `subnet-3paxmkdz`, which is required if the access type is `vpc`, `vpncloud`, `ccn`, or `dcg`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Database version in the format of `5.6` or `5.7`, which takes effect only if the instance is an RDS instance. Default value: `5.6`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	EngineVersion *string `json:"EngineVersion,omitnil,omitempty" name:"EngineVersion"`

	// Instance account
	// Note: This field may return null, indicating that no valid values can be obtained.
	Account *string `json:"Account,omitnil,omitempty" name:"Account"`

	// The role used for cross-account migration, which can contain [a-zA-Z0-9\-\_]+.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AccountRole *string `json:"AccountRole,omitnil,omitempty" name:"AccountRole"`

	// The account to which the resource belongs. Valid values: empty or `self` (the current account); `other` (another account).
	// Note: This field may return null, indicating that no valid values can be obtained.
	AccountMode *string `json:"AccountMode,omitnil,omitempty" name:"AccountMode"`

	// ID of the temporary key
	// Note: This field may return null, indicating that no valid values can be obtained.
	TmpSecretId *string `json:"TmpSecretId,omitnil,omitempty" name:"TmpSecretId"`

	// Key of the temporary key
	// Note: This field may return null, indicating that no valid values can be obtained.
	TmpSecretKey *string `json:"TmpSecretKey,omitnil,omitempty" name:"TmpSecretKey"`

	// Temporary token
	// Note: This field may return null, indicating that no valid values can be obtained.
	TmpToken *string `json:"TmpToken,omitnil,omitempty" name:"TmpToken"`
}

type DBItem struct {
	// Name of the database to be migrated or synced, which is required if `ObjectMode` is `partial`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// Name of the database after migration or sync, which is the same as the source database name by default.
	// Note: This field may return null, indicating that no valid values can be obtained.
	NewDbName *string `json:"NewDbName,omitnil,omitempty" name:"NewDbName"`

	// The schema to be migrated or synced
	// Note: This field may return null, indicating that no valid values can be obtained.
	SchemaName *string `json:"SchemaName,omitnil,omitempty" name:"SchemaName"`

	// Name of the schema after migration or sync
	// Note: This field may return null, indicating that no valid values can be obtained.
	NewSchemaName *string `json:"NewSchemaName,omitnil,omitempty" name:"NewSchemaName"`

	// Database selection mode, which is required if `ObjectMode` is `partial`. Valid values: `all`, `partial`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DBMode *string `json:"DBMode,omitnil,omitempty" name:"DBMode"`

	// Schema selection mode. Valid values: `all`, `partial`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SchemaMode *string `json:"SchemaMode,omitnil,omitempty" name:"SchemaMode"`

	// Table selection mode, which is required if `DBMode` is `partial`. Valid values: `all`, `partial`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TableMode *string `json:"TableMode,omitnil,omitempty" name:"TableMode"`

	// The set of table objects, which is required if `TableMode` is `partial`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Tables []*TableItem `json:"Tables,omitnil,omitempty" name:"Tables"`

	// View selection mode. Valid values: `all`, `partial`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ViewMode *string `json:"ViewMode,omitnil,omitempty" name:"ViewMode"`

	// The set of view objects, which is required if `ViewMode` is `partial`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Views []*ViewItem `json:"Views,omitnil,omitempty" name:"Views"`

	// Role selection mode, which is exclusive to PostgreSQL. Valid values: `all`, `partial`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	RoleMode *string `json:"RoleMode,omitnil,omitempty" name:"RoleMode"`

	// Role, which is exclusive to PostgreSQL and required if `RoleMode` is `partial`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Roles []*RoleItem `json:"Roles,omitnil,omitempty" name:"Roles"`

	// Sync mode. Valid values: `partial`, `all`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	FunctionMode *string `json:"FunctionMode,omitnil,omitempty" name:"FunctionMode"`

	// Sync mode. Valid values: `partial`, `all`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TriggerMode *string `json:"TriggerMode,omitnil,omitempty" name:"TriggerMode"`

	// Sync mode. Valid values: `partial`, `all`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	EventMode *string `json:"EventMode,omitnil,omitempty" name:"EventMode"`

	// Sync mode. Valid values: `partial`, `all`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ProcedureMode *string `json:"ProcedureMode,omitnil,omitempty" name:"ProcedureMode"`

	// This parameter is required if `FunctionMode` is `partial`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Functions []*string `json:"Functions,omitnil,omitempty" name:"Functions"`

	// This parameter is required if `ProcedureMode` is `partial`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Procedures []*string `json:"Procedures,omitnil,omitempty" name:"Procedures"`

	// This parameter is required if `EventMode` is `partial`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Events []*string `json:"Events,omitnil,omitempty" name:"Events"`

	// This parameter is required if `TriggerMode` is `partial`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Triggers []*string `json:"Triggers,omitnil,omitempty" name:"Triggers"`
}

type Database struct {
	// Name of the database to be migrated or synced, which is required if `ObjectMode` is `Partial`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// Name of the database after migration or sync, which is the same as the source database name by default.
	// Note: This field may return null, indicating that no valid values can be obtained.
	NewDbName *string `json:"NewDbName,omitnil,omitempty" name:"NewDbName"`

	// Database selection mode, which is required if `Mode` is `Partial`. Valid values: `All`, `Partial`. Note that the sync of advanced objects does not depend on this parameter. To sync an entire database, set this parameter to `All`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DbMode *string `json:"DbMode,omitnil,omitempty" name:"DbMode"`

	// The schema to be migrated or synced
	// Note: This field may return null, indicating that no valid values can be obtained.
	SchemaName *string `json:"SchemaName,omitnil,omitempty" name:"SchemaName"`

	// Name of the schema after migration or sync
	// Note: This field may return null, indicating that no valid values can be obtained.
	NewSchemaName *string `json:"NewSchemaName,omitnil,omitempty" name:"NewSchemaName"`

	// Table selection mode, which is required if `DBMode` is `Partial`. Valid values: `All`, `Partial`. To sync an entire database, set this parameter to `All`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TableMode *string `json:"TableMode,omitnil,omitempty" name:"TableMode"`

	// The set of table objects, which is required if `TableMode` is `Partial`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Tables []*Table `json:"Tables,omitnil,omitempty" name:"Tables"`

	// View selection mode. Valid values: `All`, `Partial`. To sync an entire database, set this parameter to `All`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ViewMode *string `json:"ViewMode,omitnil,omitempty" name:"ViewMode"`

	// The set of view objects, which is required if `ViewMode` is `Partial`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Views []*View `json:"Views,omitnil,omitempty" name:"Views"`

	// Sync mode. Valid values: `All`, `Partial`. To sync an entire database, set this parameter to `All`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	FunctionMode *string `json:"FunctionMode,omitnil,omitempty" name:"FunctionMode"`

	// This parameter is required if `FunctionMode` is `Partial`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Functions []*string `json:"Functions,omitnil,omitempty" name:"Functions"`

	// Sync mode. Valid values: `All`, `Partial`. To sync an entire database, set this parameter to `All`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ProcedureMode *string `json:"ProcedureMode,omitnil,omitempty" name:"ProcedureMode"`

	// This parameter is required if `ProcedureMode` is `Partial`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Procedures []*string `json:"Procedures,omitnil,omitempty" name:"Procedures"`

	// Trigger sync mode. Valid values: `All`, `Partial`. To sync an entire database, set this parameter to `All`. Currently, the advanced object “trigger” is not supported for data sync.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TriggerMode *string `json:"TriggerMode,omitnil,omitempty" name:"TriggerMode"`

	// This parameter is used to specify the names of the triggers to be migrated when the value of `TriggerMode` is `partial`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Triggers []*string `json:"Triggers,omitnil,omitempty" name:"Triggers"`

	// Event sync mode. Valid values: `All`, `Partial`. To sync an entire database, set this parameter to `All`. Currently, the advanced object “event” is not supported for data sync.
	// Note: This field may return null, indicating that no valid values can be obtained.
	EventMode *string `json:"EventMode,omitnil,omitempty" name:"EventMode"`

	// This parameter is used to specify the names of the events to be migrated when the value of `EventMode` is `partial`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Events []*string `json:"Events,omitnil,omitempty" name:"Events"`
}

type DatabaseTableObject struct {
	// Migration object type. Valid values: `all`, `partial`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ObjectMode *string `json:"ObjectMode,omitnil,omitempty" name:"ObjectMode"`

	// Migration object, which is required if `ObjectMode` is `partial`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Databases []*DBItem `json:"Databases,omitnil,omitempty" name:"Databases"`

	// Advanced object type, such as trigger, function, procedure, and event.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AdvancedObjects []*string `json:"AdvancedObjects,omitnil,omitempty" name:"AdvancedObjects"`
}

type DdlOption struct {
	// DDL type, such as database, table, view, and index.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DdlObject *string `json:"DdlObject,omitnil,omitempty" name:"DdlObject"`

	// DDL value. Valid values: [Create,Drop,Alter] for database <br>[Create,Drop,Alter,Truncate,Rename] for table <br/>[Create,Drop] for view <br/>[Create,Drop] for index
	// Note: This field may return null, indicating that no valid values can be obtained.
	DdlValue []*string `json:"DdlValue,omitnil,omitempty" name:"DdlValue"`
}

// Predefined struct for user
type DeleteCompareTaskRequestParams struct {
	// Migration task ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// Data consistency check task ID in the format of `dts-8yv4w2i1-cmp-37skmii9`
	CompareTaskId *string `json:"CompareTaskId,omitnil,omitempty" name:"CompareTaskId"`
}

type DeleteCompareTaskRequest struct {
	*tchttp.BaseRequest
	
	// Migration task ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// Data consistency check task ID in the format of `dts-8yv4w2i1-cmp-37skmii9`
	CompareTaskId *string `json:"CompareTaskId,omitnil,omitempty" name:"CompareTaskId"`
}

func (r *DeleteCompareTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCompareTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	delete(f, "CompareTaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCompareTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCompareTaskResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteCompareTaskResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCompareTaskResponseParams `json:"Response"`
}

func (r *DeleteCompareTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCompareTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCheckSyncJobResultRequestParams struct {
	// Sync task instance ID in the format of `sync-werwfs23`, which is used to identify a sync task. This parameter is required.
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

type DescribeCheckSyncJobResultRequest struct {
	*tchttp.BaseRequest
	
	// Sync task instance ID in the format of `sync-werwfs23`, which is used to identify a sync task. This parameter is required.
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

func (r *DescribeCheckSyncJobResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCheckSyncJobResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCheckSyncJobResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCheckSyncJobResultResponseParams struct {
	// Execution status of the check task. Valid values: `notStarted`, `running`, `failed`, `success`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Total number of steps
	// Note: This field may return null, indicating that no valid values can be obtained.
	StepCount *uint64 `json:"StepCount,omitnil,omitempty" name:"StepCount"`

	// The current step
	// Note: This field may return null, indicating that no valid values can be obtained.
	StepCur *uint64 `json:"StepCur,omitnil,omitempty" name:"StepCur"`

	// Overall progress. Value range: 0-100.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Progress *uint64 `json:"Progress,omitnil,omitempty" name:"Progress"`

	// Step information
	// Note: This field may return null, indicating that no valid values can be obtained.
	StepInfos []*StepInfo `json:"StepInfos,omitnil,omitempty" name:"StepInfos"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCheckSyncJobResultResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCheckSyncJobResultResponseParams `json:"Response"`
}

func (r *DescribeCheckSyncJobResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCheckSyncJobResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCompareReportRequestParams struct {
	// Migration task ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// Check task ID
	CompareTaskId *string `json:"CompareTaskId,omitnil,omitempty" name:"CompareTaskId"`

	// Number of inconsistent objects to be returned
	DifferenceLimit *uint64 `json:"DifferenceLimit,omitnil,omitempty" name:"DifferenceLimit"`

	// Offset of inconsistent objects
	DifferenceOffset *uint64 `json:"DifferenceOffset,omitnil,omitempty" name:"DifferenceOffset"`

	// Search criterion: Inconsistent database name
	DifferenceDB *string `json:"DifferenceDB,omitnil,omitempty" name:"DifferenceDB"`

	// Search criterion: Inconsistent table name
	DifferenceTable *string `json:"DifferenceTable,omitnil,omitempty" name:"DifferenceTable"`

	// Number of unchecked objects to be returned
	SkippedLimit *uint64 `json:"SkippedLimit,omitnil,omitempty" name:"SkippedLimit"`

	// Offset of unchecked objects
	SkippedOffset *uint64 `json:"SkippedOffset,omitnil,omitempty" name:"SkippedOffset"`

	// Search criterion: Unchecked database name
	SkippedDB *string `json:"SkippedDB,omitnil,omitempty" name:"SkippedDB"`

	// Search criterion: Unchecked table name
	SkippedTable *string `json:"SkippedTable,omitnil,omitempty" name:"SkippedTable"`
}

type DescribeCompareReportRequest struct {
	*tchttp.BaseRequest
	
	// Migration task ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// Check task ID
	CompareTaskId *string `json:"CompareTaskId,omitnil,omitempty" name:"CompareTaskId"`

	// Number of inconsistent objects to be returned
	DifferenceLimit *uint64 `json:"DifferenceLimit,omitnil,omitempty" name:"DifferenceLimit"`

	// Offset of inconsistent objects
	DifferenceOffset *uint64 `json:"DifferenceOffset,omitnil,omitempty" name:"DifferenceOffset"`

	// Search criterion: Inconsistent database name
	DifferenceDB *string `json:"DifferenceDB,omitnil,omitempty" name:"DifferenceDB"`

	// Search criterion: Inconsistent table name
	DifferenceTable *string `json:"DifferenceTable,omitnil,omitempty" name:"DifferenceTable"`

	// Number of unchecked objects to be returned
	SkippedLimit *uint64 `json:"SkippedLimit,omitnil,omitempty" name:"SkippedLimit"`

	// Offset of unchecked objects
	SkippedOffset *uint64 `json:"SkippedOffset,omitnil,omitempty" name:"SkippedOffset"`

	// Search criterion: Unchecked database name
	SkippedDB *string `json:"SkippedDB,omitnil,omitempty" name:"SkippedDB"`

	// Search criterion: Unchecked table name
	SkippedTable *string `json:"SkippedTable,omitnil,omitempty" name:"SkippedTable"`
}

func (r *DescribeCompareReportRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCompareReportRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	delete(f, "CompareTaskId")
	delete(f, "DifferenceLimit")
	delete(f, "DifferenceOffset")
	delete(f, "DifferenceDB")
	delete(f, "DifferenceTable")
	delete(f, "SkippedLimit")
	delete(f, "SkippedOffset")
	delete(f, "SkippedDB")
	delete(f, "SkippedTable")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCompareReportRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCompareReportResponseParams struct {
	// Summary information of data consistency check
	// Note: This field may return null, indicating that no valid values can be obtained.
	Abstract *CompareAbstractInfo `json:"Abstract,omitnil,omitempty" name:"Abstract"`

	// Data consistency check details
	// Note: This field may return null, indicating that no valid values can be obtained.
	Detail *CompareDetailInfo `json:"Detail,omitnil,omitempty" name:"Detail"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCompareReportResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCompareReportResponseParams `json:"Response"`
}

func (r *DescribeCompareReportResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCompareReportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCompareTasksRequestParams struct {
	// Migration task ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// Number of tasks to be displayed per page. Default value: `20`.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Pagination offset
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Check task ID
	CompareTaskId *string `json:"CompareTaskId,omitnil,omitempty" name:"CompareTaskId"`

	// Data consistency check task status. Valid values: `created`, `readyRun`, `running`, `success`, `stopping`, `failed`, `canceled`.
	Status []*string `json:"Status,omitnil,omitempty" name:"Status"`
}

type DescribeCompareTasksRequest struct {
	*tchttp.BaseRequest
	
	// Migration task ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// Number of tasks to be displayed per page. Default value: `20`.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Pagination offset
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Check task ID
	CompareTaskId *string `json:"CompareTaskId,omitnil,omitempty" name:"CompareTaskId"`

	// Data consistency check task status. Valid values: `created`, `readyRun`, `running`, `success`, `stopping`, `failed`, `canceled`.
	Status []*string `json:"Status,omitnil,omitempty" name:"Status"`
}

func (r *DescribeCompareTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCompareTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "CompareTaskId")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCompareTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCompareTasksResponseParams struct {
	// Quantity
	// Note: This field may return null, indicating that no valid values can be obtained.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List of data consistency check tasks
	// Note: This field may return null, indicating that no valid values can be obtained.
	Items []*CompareTaskItem `json:"Items,omitnil,omitempty" name:"Items"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCompareTasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCompareTasksResponseParams `json:"Response"`
}

func (r *DescribeCompareTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCompareTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMigrateDBInstancesRequestParams struct {
	// Database type, such as `mysql`.
	DatabaseType *string `json:"DatabaseType,omitnil,omitempty" name:"DatabaseType"`

	// Specifies whether the instance is the migration source or target. Valid values: `src` (source); `dts` (target).
	MigrateRole *string `json:"MigrateRole,omitnil,omitempty" name:"MigrateRole"`

	// Database instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Database instance name
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Number of results to be returned
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The account to which the resource belongs. Valid values: empty or `self` (the current account); `other` (another account).
	AccountMode *string `json:"AccountMode,omitnil,omitempty" name:"AccountMode"`

	// ID of the temporary key, which is required if the operation is performed across accounts.
	TmpSecretId *string `json:"TmpSecretId,omitnil,omitempty" name:"TmpSecretId"`

	// Key of the temporary key, which is required if the operation is performed across accounts.
	TmpSecretKey *string `json:"TmpSecretKey,omitnil,omitempty" name:"TmpSecretKey"`

	// Temporary token, which is required if the operation is performed across accounts.
	TmpToken *string `json:"TmpToken,omitnil,omitempty" name:"TmpToken"`
}

type DescribeMigrateDBInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Database type, such as `mysql`.
	DatabaseType *string `json:"DatabaseType,omitnil,omitempty" name:"DatabaseType"`

	// Specifies whether the instance is the migration source or target. Valid values: `src` (source); `dts` (target).
	MigrateRole *string `json:"MigrateRole,omitnil,omitempty" name:"MigrateRole"`

	// Database instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Database instance name
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Number of results to be returned
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The account to which the resource belongs. Valid values: empty or `self` (the current account); `other` (another account).
	AccountMode *string `json:"AccountMode,omitnil,omitempty" name:"AccountMode"`

	// ID of the temporary key, which is required if the operation is performed across accounts.
	TmpSecretId *string `json:"TmpSecretId,omitnil,omitempty" name:"TmpSecretId"`

	// Key of the temporary key, which is required if the operation is performed across accounts.
	TmpSecretKey *string `json:"TmpSecretKey,omitnil,omitempty" name:"TmpSecretKey"`

	// Temporary token, which is required if the operation is performed across accounts.
	TmpToken *string `json:"TmpToken,omitnil,omitempty" name:"TmpToken"`
}

func (r *DescribeMigrateDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMigrateDBInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DatabaseType")
	delete(f, "MigrateRole")
	delete(f, "InstanceId")
	delete(f, "InstanceName")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "AccountMode")
	delete(f, "TmpSecretId")
	delete(f, "TmpSecretKey")
	delete(f, "TmpToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMigrateDBInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMigrateDBInstancesResponseParams struct {
	// Number of eligible items
	// Note: This field may return null, indicating that no valid values can be obtained.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Instance list
	// Note: This field may return null, indicating that no valid values can be obtained.
	Instances []*MigrateDBItem `json:"Instances,omitnil,omitempty" name:"Instances"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMigrateDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMigrateDBInstancesResponseParams `json:"Response"`
}

func (r *DescribeMigrateDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMigrateDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMigrationCheckJobRequestParams struct {
	// Task ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

type DescribeMigrationCheckJobRequest struct {
	*tchttp.BaseRequest
	
	// Task ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

func (r *DescribeMigrationCheckJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMigrationCheckJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMigrationCheckJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMigrationCheckJobResponseParams struct {
	// Check task execution status. Valid values: `notStarted`, `running`, `failed`, `success`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Check task result message
	// Note: This field may return null, indicating that no valid values can be obtained.
	BriefMsg *string `json:"BriefMsg,omitnil,omitempty" name:"BriefMsg"`

	// Check step
	// Note: This field may return null, indicating that no valid values can be obtained.
	StepInfo []*CheckStep `json:"StepInfo,omitnil,omitempty" name:"StepInfo"`

	// Check result. Valid values: `checkPass`, `checkNotPass`.
	CheckFlag *string `json:"CheckFlag,omitnil,omitempty" name:"CheckFlag"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMigrationCheckJobResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMigrationCheckJobResponseParams `json:"Response"`
}

func (r *DescribeMigrationCheckJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMigrationCheckJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMigrationDetailRequestParams struct {
	// Data migration task ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

type DescribeMigrationDetailRequest struct {
	*tchttp.BaseRequest
	
	// Data migration task ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

func (r *DescribeMigrationDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMigrationDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMigrationDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMigrationDetailResponseParams struct {
	// Data migration task ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// Data migration task name
	// Note: This field may return null, indicating that no valid values can be obtained.
	JobName *string `json:"JobName,omitnil,omitempty" name:"JobName"`

	// Task creation (submission) time in the format of `yyyy-mm-dd hh:mm:ss`
	// Note: This field may return null, indicating that no valid values can be obtained.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Task update time in the format of `yyyy-mm-dd hh:mm:ss`
	// Note: This field may return null, indicating that no valid values can be obtained.
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// Task start time in the format of `yyyy-mm-dd hh:mm:ss`
	// Note: This field may return null, indicating that no valid values can be obtained.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Task end time in the format of `yyyy-mm-dd hh:mm:ss`
	// Note: This field may return null, indicating that no valid values can be obtained.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Migration task error message
	// Note: This field may return null, indicating that no valid values can be obtained.
	BriefMsg *string `json:"BriefMsg,omitnil,omitempty" name:"BriefMsg"`

	// Task status. Valid values: `created`(Created), `checking` (Checking), `checkPass` (Check passed), `checkNotPass` (Check not passed), `readyRun` (Ready for running), `running` (Running), `readyComplete` (Preparation completed), `success` (Successful), `failed` (Failed), `stopping` (Stopping), `completing` (Completing), `pausing` (Pausing), `manualPaused` (Paused). Note: This field may return null, indicating that no valid values can be obtained.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Task operation information
	// Note: This field may return null, indicating that no valid values can be obtained.
	Action *MigrateAction `json:"Action,omitnil,omitempty" name:"Action"`

	// Information of the migration task execution process. The check and migration step information will be displayed in the check and migration stages respectively.
	// Note: This field may return null, indicating that no valid values can be obtained.
	StepInfo *MigrateDetailInfo `json:"StepInfo,omitnil,omitempty" name:"StepInfo"`

	// Source instance information
	// Note: This field may return null, indicating that no valid values can be obtained.
	SrcInfo *DBEndpointInfo `json:"SrcInfo,omitnil,omitempty" name:"SrcInfo"`

	// Target database information
	// Note: This field may return null, indicating that no valid values can be obtained.
	DstInfo *DBEndpointInfo `json:"DstInfo,omitnil,omitempty" name:"DstInfo"`

	// Data consistency check result
	// Note: This field may return null, indicating that no valid values can be obtained.
	CompareTask *CompareTaskInfo `json:"CompareTask,omitnil,omitempty" name:"CompareTask"`

	// Tag information
	// Note: This field may return null, indicating that no valid values can be obtained.
	Tags []*TagItem `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Running mode. Valid values: `immediate`, `timed`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	RunMode *string `json:"RunMode,omitnil,omitempty" name:"RunMode"`

	// Expected start time in the format of "2006-01-02 15:04:05", which is required if `RunMode` is `timed`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ExpectRunTime *string `json:"ExpectRunTime,omitnil,omitempty" name:"ExpectRunTime"`

	// Migration options, which describe how the task performs migration.
	// Note: This field may return null, indicating that no valid values can be obtained.
	MigrateOption *MigrateOption `json:"MigrateOption,omitnil,omitempty" name:"MigrateOption"`

	// Check task running details
	// Note: This field may return null, indicating that no valid values can be obtained.
	CheckStepInfo *CheckStepInfo `json:"CheckStepInfo,omitnil,omitempty" name:"CheckStepInfo"`

	// Billing information
	// Note: This field may return null, indicating that no valid values can be obtained.
	TradeInfo *TradeInfo `json:"TradeInfo,omitnil,omitempty" name:"TradeInfo"`

	// Task error information
	// Note: This field may return null, indicating that no valid values can be obtained.
	ErrorInfo []*ErrorInfoItem `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`

	// Whether the task can be reentered in the full export stage. Valid values: `yes`, `no`. `yes`: The current task can be reentered. `no`: The current task is in the full export stage which cannot be reentered. If the value of this parameter is `no`, the checkpoint restart is not supported when the task is restarted in the export stage.
	DumperResumeCtrl *string `json:"DumperResumeCtrl,omitnil,omitempty" name:"DumperResumeCtrl"`

	// Task throttling information
	// Note: u200dThis field may returnu200d·nullu200d, indicating that no valid values can be obtained.
	RateLimitOption *RateLimitOption `json:"RateLimitOption,omitnil,omitempty" name:"RateLimitOption"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMigrationDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMigrationDetailResponseParams `json:"Response"`
}

func (r *DescribeMigrationDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMigrationDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMigrationJobsRequestParams struct {
	// Data migration task ID such as `dts-amm1jw5q`
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// Data migration task name
	JobName *string `json:"JobName,omitnil,omitempty" name:"JobName"`

	// Data migration task status. Valid values: `created`, `checking`, `checkPass`, `checkNotPass`, `readyRun`, `running`, `readyComplete`, `success`, `failed`, `stopping`, `completing`.
	Status []*string `json:"Status,omitnil,omitempty" name:"Status"`

	// Source instance ID in the format of `cdb-c1nl9rpv`
	SrcInstanceId *string `json:"SrcInstanceId,omitnil,omitempty" name:"SrcInstanceId"`

	// Source instance region, such as `ap-guangzhou`.
	SrcRegion *string `json:"SrcRegion,omitnil,omitempty" name:"SrcRegion"`

	// Source database type, such as `sqlserver`, `mysql`, `mongodb`, `redis`, `tendis`, `keewidb`, `clickhouse`, `cynosdbmysql`, `percona`, `tdsqlpercona`, `mariadb`, `tdsqlmysql`, `postgresql.
	SrcDatabaseType []*string `json:"SrcDatabaseType,omitnil,omitempty" name:"SrcDatabaseType"`

	// Source instance access type. Valid values: `extranet` (public network); `vpncloud` (VPN access); `dcg` (Direct Connect); `ccn` (CCN); `cdb` (Database); `cvm` (self-build on CVM).
	SrcAccessType []*string `json:"SrcAccessType,omitnil,omitempty" name:"SrcAccessType"`

	// Target instance ID in the format of `cdb-c1nl9rpv`
	DstInstanceId *string `json:"DstInstanceId,omitnil,omitempty" name:"DstInstanceId"`

	// Target instance region, such as `ap-guangzhou`.
	DstRegion *string `json:"DstRegion,omitnil,omitempty" name:"DstRegion"`

	// Target database type, such as `sqlserver`, `mysql`, `mongodb`, `redis`, `tendis`, `keewidb`, `clickhouse`, `cynosdbmysql`, `percona`, `tdsqlpercona`, `mariadb`, `tdsqlmysql`, `postgresql.
	DstDatabaseType []*string `json:"DstDatabaseType,omitnil,omitempty" name:"DstDatabaseType"`

	// Target instance access type. Valid values: `extranet` (public network); `vpncloud` (VPN access); `dcg` (Direct Connect); `ccn` (CCN); `cdb` (Database); `cvm` (self-build on CVM).
	DstAccessType []*string `json:"DstAccessType,omitnil,omitempty" name:"DstAccessType"`

	// Task running mode. Valid values: `immediate`, `timed`.
	RunMode *string `json:"RunMode,omitnil,omitempty" name:"RunMode"`

	// Sorting order. Valid values: `asc`, `desc`. Default value: `desc` by creation time.
	OrderSeq *string `json:"OrderSeq,omitnil,omitempty" name:"OrderSeq"`

	// Number of instances to be returned. Value range: [1,100]. Default value: `20`.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset. Default value: `0`.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Tag filter
	TagFilters []*TagFilter `json:"TagFilters,omitnil,omitempty" name:"TagFilters"`
}

type DescribeMigrationJobsRequest struct {
	*tchttp.BaseRequest
	
	// Data migration task ID such as `dts-amm1jw5q`
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// Data migration task name
	JobName *string `json:"JobName,omitnil,omitempty" name:"JobName"`

	// Data migration task status. Valid values: `created`, `checking`, `checkPass`, `checkNotPass`, `readyRun`, `running`, `readyComplete`, `success`, `failed`, `stopping`, `completing`.
	Status []*string `json:"Status,omitnil,omitempty" name:"Status"`

	// Source instance ID in the format of `cdb-c1nl9rpv`
	SrcInstanceId *string `json:"SrcInstanceId,omitnil,omitempty" name:"SrcInstanceId"`

	// Source instance region, such as `ap-guangzhou`.
	SrcRegion *string `json:"SrcRegion,omitnil,omitempty" name:"SrcRegion"`

	// Source database type, such as `sqlserver`, `mysql`, `mongodb`, `redis`, `tendis`, `keewidb`, `clickhouse`, `cynosdbmysql`, `percona`, `tdsqlpercona`, `mariadb`, `tdsqlmysql`, `postgresql.
	SrcDatabaseType []*string `json:"SrcDatabaseType,omitnil,omitempty" name:"SrcDatabaseType"`

	// Source instance access type. Valid values: `extranet` (public network); `vpncloud` (VPN access); `dcg` (Direct Connect); `ccn` (CCN); `cdb` (Database); `cvm` (self-build on CVM).
	SrcAccessType []*string `json:"SrcAccessType,omitnil,omitempty" name:"SrcAccessType"`

	// Target instance ID in the format of `cdb-c1nl9rpv`
	DstInstanceId *string `json:"DstInstanceId,omitnil,omitempty" name:"DstInstanceId"`

	// Target instance region, such as `ap-guangzhou`.
	DstRegion *string `json:"DstRegion,omitnil,omitempty" name:"DstRegion"`

	// Target database type, such as `sqlserver`, `mysql`, `mongodb`, `redis`, `tendis`, `keewidb`, `clickhouse`, `cynosdbmysql`, `percona`, `tdsqlpercona`, `mariadb`, `tdsqlmysql`, `postgresql.
	DstDatabaseType []*string `json:"DstDatabaseType,omitnil,omitempty" name:"DstDatabaseType"`

	// Target instance access type. Valid values: `extranet` (public network); `vpncloud` (VPN access); `dcg` (Direct Connect); `ccn` (CCN); `cdb` (Database); `cvm` (self-build on CVM).
	DstAccessType []*string `json:"DstAccessType,omitnil,omitempty" name:"DstAccessType"`

	// Task running mode. Valid values: `immediate`, `timed`.
	RunMode *string `json:"RunMode,omitnil,omitempty" name:"RunMode"`

	// Sorting order. Valid values: `asc`, `desc`. Default value: `desc` by creation time.
	OrderSeq *string `json:"OrderSeq,omitnil,omitempty" name:"OrderSeq"`

	// Number of instances to be returned. Value range: [1,100]. Default value: `20`.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset. Default value: `0`.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Tag filter
	TagFilters []*TagFilter `json:"TagFilters,omitnil,omitempty" name:"TagFilters"`
}

func (r *DescribeMigrationJobsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMigrationJobsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	delete(f, "JobName")
	delete(f, "Status")
	delete(f, "SrcInstanceId")
	delete(f, "SrcRegion")
	delete(f, "SrcDatabaseType")
	delete(f, "SrcAccessType")
	delete(f, "DstInstanceId")
	delete(f, "DstRegion")
	delete(f, "DstDatabaseType")
	delete(f, "DstAccessType")
	delete(f, "RunMode")
	delete(f, "OrderSeq")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "TagFilters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMigrationJobsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMigrationJobsResponseParams struct {
	// Number of migration tasks
	// Note: This field may return null, indicating that no valid values can be obtained.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Migration task list
	// Note: This field may return null, indicating that no valid values can be obtained.
	JobList []*JobItem `json:"JobList,omitnil,omitempty" name:"JobList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMigrationJobsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMigrationJobsResponseParams `json:"Response"`
}

func (r *DescribeMigrationJobsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMigrationJobsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModifyCheckSyncJobResultRequestParams struct {
	// Sync task ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

type DescribeModifyCheckSyncJobResultRequest struct {
	*tchttp.BaseRequest
	
	// Sync task ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

func (r *DescribeModifyCheckSyncJobResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModifyCheckSyncJobResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeModifyCheckSyncJobResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModifyCheckSyncJobResultResponseParams struct {
	// Execution status of the check task Valid values: `notStarted` (Not started), `running` (Running), `failed` (Failed), `success` (Successful).
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Number of check steps Note: This field may return null, indicating that no valid values can be obtained.
	StepCount *uint64 `json:"StepCount,omitnil,omitempty" name:"StepCount"`

	// Current step Note: This field may return null, indicating that no valid values can be obtained.
	StepCur *uint64 `json:"StepCur,omitnil,omitempty" name:"StepCur"`

	// Overall progress. Value range: 0-100. Note: This field may return null, indicating that no valid values can be obtained.
	Progress *uint64 `json:"Progress,omitnil,omitempty" name:"Progress"`

	// Step details Note: This field may return null, indicating that no valid values can be obtained.
	StepInfos []*StepInfo `json:"StepInfos,omitnil,omitempty" name:"StepInfos"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeModifyCheckSyncJobResultResponse struct {
	*tchttp.BaseResponse
	Response *DescribeModifyCheckSyncJobResultResponseParams `json:"Response"`
}

func (r *DescribeModifyCheckSyncJobResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModifyCheckSyncJobResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSyncJobsRequestParams struct {
	// Sync task ID, such as `sync-werwfs23`.
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// Sync task name
	JobName *string `json:"JobName,omitnil,omitempty" name:"JobName"`

	// Sort by field, such as `CreateTime`.
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// Sorting order. Valid values: `ASC`, `DESC`. Default value: `DESC` by `CreateTime`.
	OrderSeq *string `json:"OrderSeq,omitnil,omitempty" name:"OrderSeq"`

	// Offset. Default value: `0`.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of sync task instances to be returned. Value range: [1,100]. Default value: `20`.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// The set of status values, such as `Initialized,CheckPass,Running,ResumableErr,Stopped`.
	Status []*string `json:"Status,omitnil,omitempty" name:"Status"`

	// Running mode. Valid values: `Immediate`, `Timed`.
	RunMode *string `json:"RunMode,omitnil,omitempty" name:"RunMode"`

	// Task type, such as `mysql2mysql` (sync from MySQL to MySQL).
	JobType *string `json:"JobType,omitnil,omitempty" name:"JobType"`

	// Billing mode. Valid values: `PrePay` (prepaid); `PostPay` (postpaid).
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// tag
	TagFilters []*TagFilter `json:"TagFilters,omitnil,omitempty" name:"TagFilters"`
}

type DescribeSyncJobsRequest struct {
	*tchttp.BaseRequest
	
	// Sync task ID, such as `sync-werwfs23`.
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// Sync task name
	JobName *string `json:"JobName,omitnil,omitempty" name:"JobName"`

	// Sort by field, such as `CreateTime`.
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// Sorting order. Valid values: `ASC`, `DESC`. Default value: `DESC` by `CreateTime`.
	OrderSeq *string `json:"OrderSeq,omitnil,omitempty" name:"OrderSeq"`

	// Offset. Default value: `0`.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of sync task instances to be returned. Value range: [1,100]. Default value: `20`.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// The set of status values, such as `Initialized,CheckPass,Running,ResumableErr,Stopped`.
	Status []*string `json:"Status,omitnil,omitempty" name:"Status"`

	// Running mode. Valid values: `Immediate`, `Timed`.
	RunMode *string `json:"RunMode,omitnil,omitempty" name:"RunMode"`

	// Task type, such as `mysql2mysql` (sync from MySQL to MySQL).
	JobType *string `json:"JobType,omitnil,omitempty" name:"JobType"`

	// Billing mode. Valid values: `PrePay` (prepaid); `PostPay` (postpaid).
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// tag
	TagFilters []*TagFilter `json:"TagFilters,omitnil,omitempty" name:"TagFilters"`
}

func (r *DescribeSyncJobsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSyncJobsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	delete(f, "JobName")
	delete(f, "Order")
	delete(f, "OrderSeq")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Status")
	delete(f, "RunMode")
	delete(f, "JobType")
	delete(f, "PayMode")
	delete(f, "TagFilters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSyncJobsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSyncJobsResponseParams struct {
	// Number of tasks
	// Note: This field may return null, indicating that no valid values can be obtained.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Array of task details
	// Note: This field may return null, indicating that no valid values can be obtained.
	JobList []*SyncJobInfo `json:"JobList,omitnil,omitempty" name:"JobList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSyncJobsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSyncJobsResponseParams `json:"Response"`
}

func (r *DescribeSyncJobsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSyncJobsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroyMigrateJobRequestParams struct {
	// Task ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

type DestroyMigrateJobRequest struct {
	*tchttp.BaseRequest
	
	// Task ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

func (r *DestroyMigrateJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyMigrateJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DestroyMigrateJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroyMigrateJobResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DestroyMigrateJobResponse struct {
	*tchttp.BaseResponse
	Response *DestroyMigrateJobResponseParams `json:"Response"`
}

func (r *DestroyMigrateJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyMigrateJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroySyncJobRequestParams struct {
	// Sync task ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

type DestroySyncJobRequest struct {
	*tchttp.BaseRequest
	
	// Sync task ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

func (r *DestroySyncJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroySyncJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DestroySyncJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroySyncJobResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DestroySyncJobResponse struct {
	*tchttp.BaseResponse
	Response *DestroySyncJobResponseParams `json:"Response"`
}

func (r *DestroySyncJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroySyncJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DetailCheckItem struct {
	// Check item name, such as source database permission check.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CheckItemName *string `json:"CheckItemName,omitnil,omitempty" name:"CheckItemName"`

	// Check item details
	// Note: This field may return null, indicating that no valid values can be obtained.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Check item result. Valid values: `pass` (pass); `failed` (failure); `warning` (pass with warning).
	// Note: This field may return null, indicating that no valid values can be obtained.
	CheckResult *string `json:"CheckResult,omitnil,omitempty" name:"CheckResult"`

	// The cause of the check item failure
	// Note: This field may return null, indicating that no valid values can be obtained.
	FailureReason *string `json:"FailureReason,omitnil,omitempty" name:"FailureReason"`

	// Solution
	// Note: This field may return null, indicating that no valid values can be obtained.
	Solution *string `json:"Solution,omitnil,omitempty" name:"Solution"`

	// Execution error log
	// Note: This field may return null, indicating that no valid values can be obtained.
	ErrorLog []*string `json:"ErrorLog,omitnil,omitempty" name:"ErrorLog"`

	// URL of the detailed help document
	// Note: This field may return null, indicating that no valid values can be obtained.
	HelpDoc []*string `json:"HelpDoc,omitnil,omitempty" name:"HelpDoc"`

	// Prompt text for ignoring a risk
	// Note: This field may return null, indicating that no valid values can be obtained.
	SkipInfo *string `json:"SkipInfo,omitnil,omitempty" name:"SkipInfo"`
}

type DifferenceDetail struct {
	// Number of inconsistent tables
	// Note: This field may return null, indicating that no valid values can be obtained.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Details of inconsistent tables
	// Note: This field may return null, indicating that no valid values can be obtained.
	Items []*DifferenceItem `json:"Items,omitnil,omitempty" name:"Items"`
}

type DifferenceItem struct {
	// Database name
	// Note: This field may return null, indicating that no valid values can be obtained.
	Db *string `json:"Db,omitnil,omitempty" name:"Db"`

	// Table name
	// Note: This field may return null, indicating that no valid values can be obtained.
	Table *string `json:"Table,omitnil,omitempty" name:"Table"`

	// Chunk ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	Chunk *int64 `json:"Chunk,omitnil,omitempty" name:"Chunk"`

	// Source database value
	// Note: This field may return null, indicating that no valid values can be obtained.
	SrcItem *string `json:"SrcItem,omitnil,omitempty" name:"SrcItem"`

	// Target database value
	// Note: This field may return null, indicating that no valid values can be obtained.
	DstItem *string `json:"DstItem,omitnil,omitempty" name:"DstItem"`

	// Index name
	// Note: This field may return null, indicating that no valid values can be obtained.
	IndexName *string `json:"IndexName,omitnil,omitempty" name:"IndexName"`

	// First index key
	// Note: This field may return null, indicating that no valid values can be obtained.
	LowerBoundary *string `json:"LowerBoundary,omitnil,omitempty" name:"LowerBoundary"`

	// Last index key
	// Note: This field may return null, indicating that no valid values can be obtained.
	UpperBoundary *string `json:"UpperBoundary,omitnil,omitempty" name:"UpperBoundary"`

	// Comparison time in ms
	// Note: This field may return null, indicating that no valid values can be obtained.
	CostTime *float64 `json:"CostTime,omitnil,omitempty" name:"CostTime"`

	// Completion time
	// Note: This field may return null, indicating that no valid values can be obtained.
	FinishedAt *string `json:"FinishedAt,omitnil,omitempty" name:"FinishedAt"`
}

type DynamicOptions struct {
	// DML and DDL options to be synced. Valid values: `Insert` (INSERT), `Update` (UPDATE), `Delete` (DELETE), `DDL` (structure sync), `PartialDDL` (custom option, which is used together with `DdlOptions`). This parameter is required, and its value will overwrite the previous value. Note: This field may return null, indicating that no valid values can be obtained.
	OpTypes []*string `json:"OpTypes,omitnil,omitempty" name:"OpTypes"`

	// DDL options to be synced. This parameter is required when `OpTypes` is `PartialDDL`, and its value will overwrite the previous value. Note: This field may return null, indicating that no valid values can be obtained.
	DdlOptions []*DdlOption `json:"DdlOptions,omitnil,omitempty" name:"DdlOptions"`

	// Conflict resolution method. Valid values: `ReportError` (Report error), `Ignore` (Ignore), `Cover` (Overwrite), `ConditionCover` (Conditionally overwrite). Currently, this parameter cannot be modified if the target of the link is Kafka. Note: This field may return null, indicating that no valid values can be obtained.
	ConflictHandleType *string `json:"ConflictHandleType,omitnil,omitempty" name:"ConflictHandleType"`

	// Detailed options of the conflict resolution method, such as the conditionally overwritten rows and condition operations for the “conditionally overwrite” method. The internal field of this parameter cannot be modified separately. If this parameter needs to be updated, update it fully. Note: This field may return null, indicating that no valid values can be obtained.
	ConflictHandleOption *ConflictHandleOption `json:"ConflictHandleOption,omitnil,omitempty" name:"ConflictHandleOption"`
}

type Endpoint struct {
	// Region name, such as `ap-guangzhou`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Node type of TDSQL for MySQL. Enumerated values: `proxy`, `set`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Role *string `json:"Role,omitnil,omitempty" name:"Role"`

	// Database kernel type, which is used to distinguish between different kernels in TDSQL. Valid values: `percona`, `mariadb`, `mysql`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DbKernel *string `json:"DbKernel,omitnil,omitempty" name:"DbKernel"`

	// Database instance ID in the format of `cdb-powiqx8q`
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Instance IP address, which is required if the access type is not `cdb`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// Instance port, which is required if the access type is not `cdb`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Port *uint64 `json:"Port,omitnil,omitempty" name:"Port"`

	// Username, which is required for an instance authenticated by username and password.
	// Note: This field may return null, indicating that no valid values can be obtained.
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// Password, which is required for an instance authenticated by username and password.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// Database name, which is required if the database type is `cdwpg`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// VPC ID in the format of `vpc-92jblxto`, which is required if the access type is `vpc`, `dcg`, or `vpncloud`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// ID of the subnet in the VPC in the format of `subnet-3paxmkdz`, which is required if the access type is `vpc`, `dcg`, or `vpncloud`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Short CVM instance ID in the format of `ins-olgl39y8`, which is required if the access type is `cvm`. It is the same as the instance ID displayed in the CVM console.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CvmInstanceId *string `json:"CvmInstanceId,omitnil,omitempty" name:"CvmInstanceId"`

	// Direct Connect gateway ID in the format of `dcg-0rxtqqxb`, which is required if the access type is `dcg`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	UniqDcgId *string `json:"UniqDcgId,omitnil,omitempty" name:"UniqDcgId"`

	// VPN gateway ID in the format of `vpngw-9ghexg7q`, which is required if the access type is `vpncloud`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	UniqVpnGwId *string `json:"UniqVpnGwId,omitnil,omitempty" name:"UniqVpnGwId"`

	// CCN instance ID in the format of `ccn-afp6kltc`, which is required if the access type is `ccn`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CcnId *string `json:"CcnId,omitnil,omitempty" name:"CcnId"`

	// Cloud vendor type. For Alibaba Cloud ApsaraDB for RDS instances, enter `aliyun`; otherwise, enter `others`. Default value: `others`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Supplier *string `json:"Supplier,omitnil,omitempty" name:"Supplier"`

	// Database version in the format of `5.6` or `5.7`, which takes effect only if the instance is an RDS instance. Default value: `5.6`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	EngineVersion *string `json:"EngineVersion,omitnil,omitempty" name:"EngineVersion"`

	// Instance account, which is required if the operation is performed across accounts.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Account *string `json:"Account,omitnil,omitempty" name:"Account"`

	// The account to which the resource belongs. Valid values: empty or `self` (the current account); `other` (another account).
	// Note: This field may return null, indicating that no valid values can be obtained.
	AccountMode *string `json:"AccountMode,omitnil,omitempty" name:"AccountMode"`

	// The role used for cross-account sync, which can contain [a-zA-Z0-9\-\_]+ and is required if the operation is performed across accounts.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AccountRole *string `json:"AccountRole,omitnil,omitempty" name:"AccountRole"`

	// External role ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	RoleExternalId *string `json:"RoleExternalId,omitnil,omitempty" name:"RoleExternalId"`

	// ID of the temporary key, which is required if the operation is performed across accounts.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TmpSecretId *string `json:"TmpSecretId,omitnil,omitempty" name:"TmpSecretId"`

	// Key of the temporary key, which is required if the operation is performed across accounts.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TmpSecretKey *string `json:"TmpSecretKey,omitnil,omitempty" name:"TmpSecretKey"`

	// Temporary token, which is required if the operation is performed across accounts.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TmpToken *string `json:"TmpToken,omitnil,omitempty" name:"TmpToken"`

	// Whether to enable encrypted transfer (`UnEncrypted`: No; `Encrypted`: Yes). Default value: `UnEncrypted`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	EncryptConn *string `json:"EncryptConn,omitnil,omitempty" name:"EncryptConn"`

	// Network environment of the database. This parameter is required when `AccessType` is `ccn`. Valid values: `UserIDC` (user IDC), `TencentVPC` (Tencent Cloud VPC).
	// Note: This field may return null, indicating that no valid values can be obtained.
	DatabaseNetEnv *string `json:"DatabaseNetEnv,omitnil,omitempty" name:"DatabaseNetEnv"`

	// The root account of CCN in the scenario where the database is connected to CCN under another Tencent Cloud account
	// Note: u200dThis field may return `null`, indicating that no valid values can be obtained.
	CcnOwnerUin *string `json:"CcnOwnerUin,omitnil,omitempty" name:"CcnOwnerUin"`
}

type ErrInfo struct {
	// Cause of the error
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`

	// Error message
	// Note: u200dThis field may returnu200d·nullu200d, indicating that no valid values can be obtained.
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// Solution
	// Note: u200dThis field may returnu200d·nullu200d, indicating that no valid values can be obtained.
	Solution *string `json:"Solution,omitnil,omitempty" name:"Solution"`
}

type ErrorInfoItem struct {
	// Error code
	// Note: This field may return null, indicating that no valid values can be obtained.
	Code *string `json:"Code,omitnil,omitempty" name:"Code"`

	// Solution
	// Note: This field may return null, indicating that no valid values can be obtained.
	Solution *string `json:"Solution,omitnil,omitempty" name:"Solution"`

	// Error log
	// Note: This field may return null, indicating that no valid values can be obtained.
	ErrorLog *string `json:"ErrorLog,omitnil,omitempty" name:"ErrorLog"`

	// Help document
	// Note: This field may return null, indicating that no valid values can be obtained.
	HelpDoc *string `json:"HelpDoc,omitnil,omitempty" name:"HelpDoc"`
}

// Predefined struct for user
type IsolateMigrateJobRequestParams struct {
	// Task ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

type IsolateMigrateJobRequest struct {
	*tchttp.BaseRequest
	
	// Task ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

func (r *IsolateMigrateJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IsolateMigrateJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "IsolateMigrateJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type IsolateMigrateJobResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type IsolateMigrateJobResponse struct {
	*tchttp.BaseResponse
	Response *IsolateMigrateJobResponseParams `json:"Response"`
}

func (r *IsolateMigrateJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IsolateMigrateJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type IsolateSyncJobRequestParams struct {
	// Sync task ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

type IsolateSyncJobRequest struct {
	*tchttp.BaseRequest
	
	// Sync task ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

func (r *IsolateSyncJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IsolateSyncJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "IsolateSyncJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type IsolateSyncJobResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type IsolateSyncJobResponse struct {
	*tchttp.BaseResponse
	Response *IsolateSyncJobResponseParams `json:"Response"`
}

func (r *IsolateSyncJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IsolateSyncJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type JobItem struct {
	// Data migration task ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// Data migration task name
	// Note: This field may return null, indicating that no valid values can be obtained.
	JobName *string `json:"JobName,omitnil,omitempty" name:"JobName"`

	// Task creation (submission) time in the format of `yyyy-mm-dd hh:mm:ss`
	// Note: This field may return null, indicating that no valid values can be obtained.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Task update time in the format of `yyyy-mm-dd hh:mm:ss`
	// Note: This field may return null, indicating that no valid values can be obtained.
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// Task start time in the format of `yyyy-mm-dd hh:mm:ss`
	// Note: This field may return null, indicating that no valid values can be obtained.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Task end time in the format of `yyyy-mm-dd hh:mm:ss`
	// Note: This field may return null, indicating that no valid values can be obtained.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Migration task error message
	// Note: This field may return null, indicating that no valid values can be obtained.
	BriefMsg *string `json:"BriefMsg,omitnil,omitempty" name:"BriefMsg"`

	// Task status. Valid values: `creating` (Creating), `created`(Created), `checking` (Checking), `checkPass` (Check passed), `checkNotPass` (Check not passed), `readyRun` (Ready for running), `running` (Running), `readyComplete` (Preparation completed), `success` (Successful), `failed` (Failed), `stopping` (Stopping), `completing` (Completing), `pausing` (Pausing), `manualPaused` (Paused). Note: This field may return null, indicating that no valid values can be obtained.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Task running mode. Valid values: `immediate`, `timed`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	RunMode *string `json:"RunMode,omitnil,omitempty" name:"RunMode"`

	// Expected start time in the format of "2022-07-11 16:20:49", which is required if `RunMode` is `timed`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ExpectRunTime *string `json:"ExpectRunTime,omitnil,omitempty" name:"ExpectRunTime"`

	// Task operation information
	// Note: This field may return null, indicating that no valid values can be obtained.
	Action *MigrateAction `json:"Action,omitnil,omitempty" name:"Action"`

	// Information of the migration task execution process
	// Note: This field may return null, indicating that no valid values can be obtained.
	StepInfo *MigrateDetailInfo `json:"StepInfo,omitnil,omitempty" name:"StepInfo"`

	// Source instance information
	// Note: This field may return null, indicating that no valid values can be obtained.
	SrcInfo *DBEndpointInfo `json:"SrcInfo,omitnil,omitempty" name:"SrcInfo"`

	// Target database information
	// Note: This field may return null, indicating that no valid values can be obtained.
	DstInfo *DBEndpointInfo `json:"DstInfo,omitnil,omitempty" name:"DstInfo"`

	// Data consistency check result
	// Note: This field may return null, indicating that no valid values can be obtained.
	CompareTask *CompareTaskInfo `json:"CompareTask,omitnil,omitempty" name:"CompareTask"`

	// Billing status information
	// Note: This field may return null, indicating that no valid values can be obtained.
	TradeInfo *TradeInfo `json:"TradeInfo,omitnil,omitempty" name:"TradeInfo"`

	// Tag information
	// Note: This field may return null, indicating that no valid values can be obtained.
	Tags []*TagItem `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Information of automatic retry time
	// Note: This field may return null, indicating that no valid values can be obtained.
	AutoRetryTimeRangeMinutes *int64 `json:"AutoRetryTimeRangeMinutes,omitnil,omitempty" name:"AutoRetryTimeRangeMinutes"`

	// Whether the task can be reentered in the full export stage. Valid values: `yes`, `no`. `yes`: The current task can be reentered. `no`: The current task is in the full export stage which cannot be reentered. If the value of this parameter is `no`, the checkpoint restart is not supported when the task is restarted in the export stage.
	// Note: u200dThis field may returnu200d·nullu200d, indicating that no valid values can be obtained.
	DumperResumeCtrl *string `json:"DumperResumeCtrl,omitnil,omitempty" name:"DumperResumeCtrl"`
}

type KafkaOption struct {
	// Type of data that is delivered to Kafka, such as `Avro` and `Json`.
	DataType *string `json:"DataType,omitnil,omitempty" name:"DataType"`

	// Topic sync policy, such as `Single` (deliver all data to a single topic), `Multi` (deliver data to multiple custom topics).
	TopicType *string `json:"TopicType,omitnil,omitempty" name:"TopicType"`

	// Topic for DDL storage
	DDLTopicName *string `json:"DDLTopicName,omitnil,omitempty" name:"DDLTopicName"`

	// Topic description
	TopicRules []*TopicRule `json:"TopicRules,omitnil,omitempty" name:"TopicRules"`
}

type KeyValuePairOption struct {
	// Option key
	// Note: This field may return null, indicating that no valid values can be obtained.
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// Option value
	// Note: This field may return null, indicating that no valid values can be obtained.
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type MigrateAction struct {
	// List of all operations in the task
	// Note: This field may return null, indicating that no valid values can be obtained.
	AllAction []*string `json:"AllAction,omitnil,omitempty" name:"AllAction"`

	// List of allowed operations in the task under the current status
	// Note: This field may return null, indicating that no valid values can be obtained.
	AllowedAction []*string `json:"AllowedAction,omitnil,omitempty" name:"AllowedAction"`
}

type MigrateDBItem struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Instance name
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Instance VIP
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// Instance Vport
	Vport *int64 `json:"Vport,omitnil,omitempty" name:"Vport"`

	// Whether the instance can be migrated. Valid values: `1 (yes); `0` (no).
	Usable *int64 `json:"Usable,omitnil,omitempty" name:"Usable"`

	// The cause why the instance cannot be migrated
	Hint *string `json:"Hint,omitnil,omitempty" name:"Hint"`
}

type MigrateDetailInfo struct {
	// Total number of steps
	// Note: This field may return null, indicating that no valid values can be obtained.
	StepAll *uint64 `json:"StepAll,omitnil,omitempty" name:"StepAll"`

	// Current step
	// Note: This field may return null, indicating that no valid values can be obtained.
	StepNow *uint64 `json:"StepNow,omitnil,omitempty" name:"StepNow"`

	// Source-replica lag in MB. This parameter takes effect only when the task is normal and is in the last step of migration or sync (binlog sync). If it is invalid, `-1` will be returned.
	// Note: This field may return null, indicating that no valid values can be obtained.
	MasterSlaveDistance *int64 `json:"MasterSlaveDistance,omitnil,omitempty" name:"MasterSlaveDistance"`

	// Source-replica lag in seconds. This parameter takes effect only when the task is normal and is in the last step of migration or sync (binlog sync). If it is invalid, `-1` will be returned.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SecondsBehindMaster *int64 `json:"SecondsBehindMaster,omitnil,omitempty" name:"SecondsBehindMaster"`

	// Step information
	// Note: This field may return null, indicating that no valid values can be obtained.
	StepInfo []*StepDetailInfo `json:"StepInfo,omitnil,omitempty" name:"StepInfo"`
}

type MigrateOption struct {
	// Migration object options, which tell DTS which database/table objects should be migrated.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DatabaseTable *DatabaseTableObject `json:"DatabaseTable,omitnil,omitempty" name:"DatabaseTable"`

	// Migration type. Valid values: `full`, `structure`, `fullAndIncrement`. Default value: `fullAndIncrement`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	MigrateType *string `json:"MigrateType,omitnil,omitempty" name:"MigrateType"`

	// Data consistency check option. Data consistency check is disabled by default.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Consistency *ConsistencyOption `json:"Consistency,omitnil,omitempty" name:"Consistency"`

	// Whether to migrate accounts. Valid values: `yes`, `no`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsMigrateAccount *bool `json:"IsMigrateAccount,omitnil,omitempty" name:"IsMigrateAccount"`

	// Whether to use the `Root` account in the source database to overwrite that in the target database. Valid values: `false`, `true`. For database/table or structural migration, you should specify `false`. Note that this parameter takes effect only for OldDTS.
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsOverrideRoot *bool `json:"IsOverrideRoot,omitnil,omitempty" name:"IsOverrideRoot"`

	// Whether to set the target database to read-only during migration, which takes effect only for MySQL databases. Valid values: `true`, `false`. Default value: `false`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsDstReadOnly *bool `json:"IsDstReadOnly,omitnil,omitempty" name:"IsDstReadOnly"`

	// Additional information. You can set additional parameters for certain database types. For Redis, you can define the following parameters: 
	// ["DstWriteMode": `normal`. 	Target database write mode. Valid values: `clearData` (Clear the target instance data), overwrite` (Execute the task in overwriting mode), `normal` (Follow the normal steps) 	"IsDstReadOnly": `true`. 	Whether to set the target database to read-only for a migration task. Valid values: `true` (Yes), `false` (No) 	"ClientOutputBufferHardLimit": 512. 	Hard limit of the replica buffer zone capacity in MB. 	"ClientOutputBufferSoftLimit": 512. 	Soft limit of the replica buffer zone capacity in MB. 	"ClientOutputBufferPersistTime": 60. Soft limit duration of the replica buffer zone in seconds. 	"ReplBacklogSize": 512, 	Limit of the circular buffer zone capacity in MB. 	"ReplTimeout":120，		Replication timeout period in seconds]
	// Note: u200dThis field may returnu200d·nullu200d, indicating that no valid values can be obtained.
	ExtraAttr []*KeyValuePairOption `json:"ExtraAttr,omitnil,omitempty" name:"ExtraAttr"`
}

// Predefined struct for user
type ModifyCompareTaskNameRequestParams struct {
	// Migration task ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// Data consistency check task ID in the format of `dts-8yv4w2i1-cmp-37skmii9`
	CompareTaskId *string `json:"CompareTaskId,omitnil,omitempty" name:"CompareTaskId"`

	// Data consistency check task name
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`
}

type ModifyCompareTaskNameRequest struct {
	*tchttp.BaseRequest
	
	// Migration task ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// Data consistency check task ID in the format of `dts-8yv4w2i1-cmp-37skmii9`
	CompareTaskId *string `json:"CompareTaskId,omitnil,omitempty" name:"CompareTaskId"`

	// Data consistency check task name
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`
}

func (r *ModifyCompareTaskNameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCompareTaskNameRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	delete(f, "CompareTaskId")
	delete(f, "TaskName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCompareTaskNameRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCompareTaskNameResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyCompareTaskNameResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCompareTaskNameResponseParams `json:"Response"`
}

func (r *ModifyCompareTaskNameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCompareTaskNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCompareTaskRequestParams struct {
	// Task ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// Data consistency check task ID in the format of `dts-8yv4w2i1-cmp-37skmii9`
	CompareTaskId *string `json:"CompareTaskId,omitnil,omitempty" name:"CompareTaskId"`

	// Task name
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// Data comparison object mode. Valid values: `sameAsMigrate` (All migration objects), `custom` (Custom mode. The custom comparison objects must be a subset of the migration objects). Default value: `sameAsMigrate`.
	ObjectMode *string `json:"ObjectMode,omitnil,omitempty" name:"ObjectMode"`

	// Compared object, which is required if `CompareObjectMode` is `custom`.
	Objects *CompareObject `json:"Objects,omitnil,omitempty" name:"Objects"`

	// Consistency check options
	Options *CompareOptions `json:"Options,omitnil,omitempty" name:"Options"`
}

type ModifyCompareTaskRequest struct {
	*tchttp.BaseRequest
	
	// Task ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// Data consistency check task ID in the format of `dts-8yv4w2i1-cmp-37skmii9`
	CompareTaskId *string `json:"CompareTaskId,omitnil,omitempty" name:"CompareTaskId"`

	// Task name
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// Data comparison object mode. Valid values: `sameAsMigrate` (All migration objects), `custom` (Custom mode. The custom comparison objects must be a subset of the migration objects). Default value: `sameAsMigrate`.
	ObjectMode *string `json:"ObjectMode,omitnil,omitempty" name:"ObjectMode"`

	// Compared object, which is required if `CompareObjectMode` is `custom`.
	Objects *CompareObject `json:"Objects,omitnil,omitempty" name:"Objects"`

	// Consistency check options
	Options *CompareOptions `json:"Options,omitnil,omitempty" name:"Options"`
}

func (r *ModifyCompareTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCompareTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	delete(f, "CompareTaskId")
	delete(f, "TaskName")
	delete(f, "ObjectMode")
	delete(f, "Objects")
	delete(f, "Options")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCompareTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCompareTaskResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyCompareTaskResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCompareTaskResponseParams `json:"Response"`
}

func (r *ModifyCompareTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCompareTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMigrateJobSpecRequestParams struct {
	// Task ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// New instance specification. Valid values: `micro`, `small`, `medium`, `large`, `xlarge`, `2xlarge`.
	NewInstanceClass *string `json:"NewInstanceClass,omitnil,omitempty" name:"NewInstanceClass"`
}

type ModifyMigrateJobSpecRequest struct {
	*tchttp.BaseRequest
	
	// Task ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// New instance specification. Valid values: `micro`, `small`, `medium`, `large`, `xlarge`, `2xlarge`.
	NewInstanceClass *string `json:"NewInstanceClass,omitnil,omitempty" name:"NewInstanceClass"`
}

func (r *ModifyMigrateJobSpecRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMigrateJobSpecRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	delete(f, "NewInstanceClass")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyMigrateJobSpecRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMigrateJobSpecResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyMigrateJobSpecResponse struct {
	*tchttp.BaseResponse
	Response *ModifyMigrateJobSpecResponseParams `json:"Response"`
}

func (r *ModifyMigrateJobSpecResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMigrateJobSpecResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMigrateNameRequestParams struct {
	// Migration task ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// New migration task name
	JobName *string `json:"JobName,omitnil,omitempty" name:"JobName"`
}

type ModifyMigrateNameRequest struct {
	*tchttp.BaseRequest
	
	// Migration task ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// New migration task name
	JobName *string `json:"JobName,omitnil,omitempty" name:"JobName"`
}

func (r *ModifyMigrateNameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMigrateNameRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	delete(f, "JobName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyMigrateNameRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMigrateNameResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyMigrateNameResponse struct {
	*tchttp.BaseResponse
	Response *ModifyMigrateNameResponseParams `json:"Response"`
}

func (r *ModifyMigrateNameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMigrateNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMigrationJobRequestParams struct {
	// Task ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// Running mode. Valid values: `immediate`, `timed`.
	RunMode *string `json:"RunMode,omitnil,omitempty" name:"RunMode"`

	// Migration task configuration options, which describe how the task performs migration. The `RateLimitOption` option cannot be configured. To modify the speed limit settings of the task, use the `ModifyMigrateRateLimit` API after the task starts running.
	MigrateOption *MigrateOption `json:"MigrateOption,omitnil,omitempty" name:"MigrateOption"`

	// Source instance information
	SrcInfo *DBEndpointInfo `json:"SrcInfo,omitnil,omitempty" name:"SrcInfo"`

	// Target instance information
	DstInfo *DBEndpointInfo `json:"DstInfo,omitnil,omitempty" name:"DstInfo"`

	// Migration task name, which can contain up to 128 characters.
	JobName *string `json:"JobName,omitnil,omitempty" name:"JobName"`

	// Expected start time in the format of "2006-01-02 15:04:05", which is required if `RunMode` is `timed`.
	ExpectRunTime *string `json:"ExpectRunTime,omitnil,omitempty" name:"ExpectRunTime"`

	// Tag information
	Tags []*TagItem `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Automatic retry time, which can be set to 5-720 minutes. 0 indicates that retry is disabled.
	AutoRetryTimeRangeMinutes *int64 `json:"AutoRetryTimeRangeMinutes,omitnil,omitempty" name:"AutoRetryTimeRangeMinutes"`
}

type ModifyMigrationJobRequest struct {
	*tchttp.BaseRequest
	
	// Task ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// Running mode. Valid values: `immediate`, `timed`.
	RunMode *string `json:"RunMode,omitnil,omitempty" name:"RunMode"`

	// Migration task configuration options, which describe how the task performs migration. The `RateLimitOption` option cannot be configured. To modify the speed limit settings of the task, use the `ModifyMigrateRateLimit` API after the task starts running.
	MigrateOption *MigrateOption `json:"MigrateOption,omitnil,omitempty" name:"MigrateOption"`

	// Source instance information
	SrcInfo *DBEndpointInfo `json:"SrcInfo,omitnil,omitempty" name:"SrcInfo"`

	// Target instance information
	DstInfo *DBEndpointInfo `json:"DstInfo,omitnil,omitempty" name:"DstInfo"`

	// Migration task name, which can contain up to 128 characters.
	JobName *string `json:"JobName,omitnil,omitempty" name:"JobName"`

	// Expected start time in the format of "2006-01-02 15:04:05", which is required if `RunMode` is `timed`.
	ExpectRunTime *string `json:"ExpectRunTime,omitnil,omitempty" name:"ExpectRunTime"`

	// Tag information
	Tags []*TagItem `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Automatic retry time, which can be set to 5-720 minutes. 0 indicates that retry is disabled.
	AutoRetryTimeRangeMinutes *int64 `json:"AutoRetryTimeRangeMinutes,omitnil,omitempty" name:"AutoRetryTimeRangeMinutes"`
}

func (r *ModifyMigrationJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMigrationJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	delete(f, "RunMode")
	delete(f, "MigrateOption")
	delete(f, "SrcInfo")
	delete(f, "DstInfo")
	delete(f, "JobName")
	delete(f, "ExpectRunTime")
	delete(f, "Tags")
	delete(f, "AutoRetryTimeRangeMinutes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyMigrationJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMigrationJobResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyMigrationJobResponse struct {
	*tchttp.BaseResponse
	Response *ModifyMigrationJobResponseParams `json:"Response"`
}

func (r *ModifyMigrationJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMigrationJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySyncJobConfigRequestParams struct {
	// Sync task ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// The modified sync objects
	DynamicObjects *Objects `json:"DynamicObjects,omitnil,omitempty" name:"DynamicObjects"`

	// The modified sync task options
	DynamicOptions *DynamicOptions `json:"DynamicOptions,omitnil,omitempty" name:"DynamicOptions"`
}

type ModifySyncJobConfigRequest struct {
	*tchttp.BaseRequest
	
	// Sync task ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// The modified sync objects
	DynamicObjects *Objects `json:"DynamicObjects,omitnil,omitempty" name:"DynamicObjects"`

	// The modified sync task options
	DynamicOptions *DynamicOptions `json:"DynamicOptions,omitnil,omitempty" name:"DynamicOptions"`
}

func (r *ModifySyncJobConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySyncJobConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	delete(f, "DynamicObjects")
	delete(f, "DynamicOptions")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySyncJobConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySyncJobConfigResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifySyncJobConfigResponse struct {
	*tchttp.BaseResponse
	Response *ModifySyncJobConfigResponseParams `json:"Response"`
}

func (r *ModifySyncJobConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySyncJobConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Objects struct {
	// Sync object type. Valid value: `Partial` (Partial objects). Note: This field may return null, indicating that no valid values can be obtained.
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`

	// Sync object, which is required if `Mode` is `Partial`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Databases []*Database `json:"Databases,omitnil,omitempty" name:"Databases"`

	// Advanced object type, such as functions and procedures. If you need to sync advanced objects, the initialization type must include structure initialization, that is, `Options.InitType` must be `Structure` or `Full`. Note: This field may return null, indicating that no valid values can be obtained.
	AdvancedObjects []*string `json:"AdvancedObjects,omitnil,omitempty" name:"AdvancedObjects"`

	// A redundant field that specifies the online DDL type
	// Note: This field may return null, indicating that no valid values can be obtained.
	OnlineDDL *OnlineDDL `json:"OnlineDDL,omitnil,omitempty" name:"OnlineDDL"`
}

type OnlineDDL struct {
	// Status
	// Note: This field may return null, indicating that no valid values can be obtained.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

type Options struct {
	// Sync initialization option. Valid values: `data` (full data initialization); `Structure` (structure initialization); `Full` (full data and structure initialization); `None` (incremental data only). Default value: `Full`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	InitType *string `json:"InitType,omitnil,omitempty" name:"InitType"`

	// Processing method for duplicate tables. Valid values: `ReportErrorAfterCheck`, `InitializeAfterDelete`, `ExecuteAfterIgnore`. Default value: `ReportErrorAfterCheck`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DealOfExistSameTable *string `json:"DealOfExistSameTable,omitnil,omitempty" name:"DealOfExistSameTable"`

	// Conflict processing option. Valid values: `ReportError` (report an error); `Ignore` (ignore); `Cover` (overwrite); `ConditionCover` (conditionally overwrite). Default value: `ReportError`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ConflictHandleType *string `json:"ConflictHandleType,omitnil,omitempty" name:"ConflictHandleType"`

	// Whether to add the additional column
	// Note: This field may return null, indicating that no valid values can be obtained.
	AddAdditionalColumn *bool `json:"AddAdditionalColumn,omitnil,omitempty" name:"AddAdditionalColumn"`

	// DML and DDL options to be synced. Valid values: `Insert` (INSERT operations); `Update` (UPDATE operations); `Delete` (DELETE operations); `DDL` (structure sync); `PartialDDL` (custom option, which is used together with `DdlOptions`). You can also leave this parameter empty.
	// Note: This field may return null, indicating that no valid values can be obtained.
	OpTypes []*string `json:"OpTypes,omitnil,omitempty" name:"OpTypes"`

	// Detailed option for conflict processing, such as condition rows and operations in conditional overwrite.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ConflictHandleOption *ConflictHandleOption `json:"ConflictHandleOption,omitnil,omitempty" name:"ConflictHandleOption"`

	// DDL statements to be synced
	// Note: This field may return null, indicating that no valid values can be obtained.
	DdlOptions []*DdlOption `json:"DdlOptions,omitnil,omitempty" name:"DdlOptions"`

	// Kafka sync options
	// Note: This field may return null, indicating that no valid values can be obtained.
	KafkaOption *KafkaOption `json:"KafkaOption,omitnil,omitempty" name:"KafkaOption"`

	// Task speed limit information. This parameter can only be used as an output parameter.
	// Note: u200dThis field may returnu200d·nullu200d, indicating that no valid values can be obtained.
	RateLimitOption *RateLimitOption `json:"RateLimitOption,omitnil,omitempty" name:"RateLimitOption"`

	// Settings of the automatic retry time range
	// Note: u200dThis field may returnu200d·nullu200d, indicating that no valid values can be obtained.
	AutoRetryTimeRangeMinutes *int64 `json:"AutoRetryTimeRangeMinutes,omitnil,omitempty" name:"AutoRetryTimeRangeMinutes"`
}

// Predefined struct for user
type PauseMigrateJobRequestParams struct {
	// Data migration task ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

type PauseMigrateJobRequest struct {
	*tchttp.BaseRequest
	
	// Data migration task ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

func (r *PauseMigrateJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PauseMigrateJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PauseMigrateJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PauseMigrateJobResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type PauseMigrateJobResponse struct {
	*tchttp.BaseResponse
	Response *PauseMigrateJobResponseParams `json:"Response"`
}

func (r *PauseMigrateJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PauseMigrateJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PauseSyncJobRequestParams struct {
	// Sync task ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

type PauseSyncJobRequest struct {
	*tchttp.BaseRequest
	
	// Sync task ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

func (r *PauseSyncJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PauseSyncJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PauseSyncJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PauseSyncJobResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type PauseSyncJobResponse struct {
	*tchttp.BaseResponse
	Response *PauseSyncJobResponseParams `json:"Response"`
}

func (r *PauseSyncJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PauseSyncJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProcessProgress struct {
	// Step status. Valid values: `notStarted`, `running`, `success`, `failed`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Progress information
	// Note: This field may return null, indicating that no valid values can be obtained.
	Percent *uint64 `json:"Percent,omitnil,omitempty" name:"Percent"`

	// Total number of steps
	// Note: This field may return null, indicating that no valid values can be obtained.
	StepAll *uint64 `json:"StepAll,omitnil,omitempty" name:"StepAll"`

	// Current step
	// Note: This field may return null, indicating that no valid values can be obtained.
	StepNow *uint64 `json:"StepNow,omitnil,omitempty" name:"StepNow"`

	// The prompt output in the current step
	// Note: This field may return null, indicating that no valid values can be obtained.
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// Step information
	// Note: This field may return null, indicating that no valid values can be obtained.
	Steps []*StepDetailInfo `json:"Steps,omitnil,omitempty" name:"Steps"`
}

type ProcessStepTip struct {
	// Prompt message
	// Note: This field may return null, indicating that no valid values can be obtained.
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// Solution
	// Note: This field may return null, indicating that no valid values can be obtained.
	Solution *string `json:"Solution,omitnil,omitempty" name:"Solution"`

	// Help document
	// Note: This field may return null, indicating that no valid values can be obtained.
	HelpDoc *string `json:"HelpDoc,omitnil,omitempty" name:"HelpDoc"`
}

type RateLimitOption struct {
	// The number of full export threads that have taken effect.
	// Note: u200dThis field may returnu200d·nullu200d, indicating that no valid values can be obtained.
	CurrentDumpThread *int64 `json:"CurrentDumpThread,omitnil,omitempty" name:"CurrentDumpThread"`

	// The default number of full export threads.
	// Note: u200dThis field may returnu200d·nullu200d, indicating that no valid values can be obtained.
	DefaultDumpThread *int64 `json:"DefaultDumpThread,omitnil,omitempty" name:"DefaultDumpThread"`

	// The full export RPS that has taken effect.	
	// Note: u200dThis field may returnu200d·nullu200d, indicating that no valid values can be obtained.
	CurrentDumpRps *int64 `json:"CurrentDumpRps,omitnil,omitempty" name:"CurrentDumpRps"`

	// The default full export RPS.	
	// Note: u200dThis field may return null, indicating that no valid values can be obtained.
	DefaultDumpRps *int64 `json:"DefaultDumpRps,omitnil,omitempty" name:"DefaultDumpRps"`

	// The number of full import threads that have taken effect.
	// Note: u200dThis field may returnu200d·nullu200d, indicating that no valid values can be obtained.
	CurrentLoadThread *int64 `json:"CurrentLoadThread,omitnil,omitempty" name:"CurrentLoadThread"`

	// The default number of full import threads.
	// Note: u200dThis field may returnu200d·nullu200d, indicating that no valid values can be obtained.
	DefaultLoadThread *int64 `json:"DefaultLoadThread,omitnil,omitempty" name:"DefaultLoadThread"`

	// The full import RPS that has taken effect.	
	// Note: u200dThis field may returnu200d·nullu200d, indicating that no valid values can be obtained.
	CurrentLoadRps *int64 `json:"CurrentLoadRps,omitnil,omitempty" name:"CurrentLoadRps"`

	// The default full import RPS.	
	// Note: u200dThis field may returnu200d·nullu200d, indicating that no valid values can be obtained.
	DefaultLoadRps *int64 `json:"DefaultLoadRps,omitnil,omitempty" name:"DefaultLoadRps"`

	// The number of incremental import threads that have taken effect.
	// Note: u200dThis field may returnu200d·nullu200d, indicating that no valid values can be obtained.
	CurrentSinkerThread *int64 `json:"CurrentSinkerThread,omitnil,omitempty" name:"CurrentSinkerThread"`

	// The default number of incremental import threads.
	// Note: u200dThis field may returnu200d·nullu200d, indicating that no valid values can be obtained.
	DefaultSinkerThread *int64 `json:"DefaultSinkerThread,omitnil,omitempty" name:"DefaultSinkerThread"`

	// Whether the speed limit has been set. Valid values: `no` (No), `yes` (Yes).
	// Note: u200dThis field may returnu200d·nullu200d, indicating that no valid values can be obtained.
	HasUserSetRateLimit *string `json:"HasUserSetRateLimit,omitnil,omitempty" name:"HasUserSetRateLimit"`
}

// Predefined struct for user
type RecoverMigrateJobRequestParams struct {
	// Task ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

type RecoverMigrateJobRequest struct {
	*tchttp.BaseRequest
	
	// Task ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

func (r *RecoverMigrateJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecoverMigrateJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RecoverMigrateJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecoverMigrateJobResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RecoverMigrateJobResponse struct {
	*tchttp.BaseResponse
	Response *RecoverMigrateJobResponseParams `json:"Response"`
}

func (r *RecoverMigrateJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecoverMigrateJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecoverSyncJobRequestParams struct {
	// Sync task instance ID in the format of `sync-werwfs23`, which is used to identify a sync task.
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

type RecoverSyncJobRequest struct {
	*tchttp.BaseRequest
	
	// Sync task instance ID in the format of `sync-werwfs23`, which is used to identify a sync task.
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

func (r *RecoverSyncJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecoverSyncJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RecoverSyncJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecoverSyncJobResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RecoverSyncJobResponse struct {
	*tchttp.BaseResponse
	Response *RecoverSyncJobResponseParams `json:"Response"`
}

func (r *RecoverSyncJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecoverSyncJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResizeSyncJobRequestParams struct {
	// Sync task ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// Task specification
	NewInstanceClass *string `json:"NewInstanceClass,omitnil,omitempty" name:"NewInstanceClass"`
}

type ResizeSyncJobRequest struct {
	*tchttp.BaseRequest
	
	// Sync task ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// Task specification
	NewInstanceClass *string `json:"NewInstanceClass,omitnil,omitempty" name:"NewInstanceClass"`
}

func (r *ResizeSyncJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResizeSyncJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	delete(f, "NewInstanceClass")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResizeSyncJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResizeSyncJobResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ResizeSyncJobResponse struct {
	*tchttp.BaseResponse
	Response *ResizeSyncJobResponseParams `json:"Response"`
}

func (r *ResizeSyncJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResizeSyncJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResumeMigrateJobRequestParams struct {
	// Data migration task ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// Task resumption mode. Valid values: `clearData` (Clearing the target instance data); `overwrite` (Executing the task in overwrite mode); `normal` (Following the normal process without additional operations). `clearData` and `overwrite` are only valid for Redis links and `normal` for non-Redis links.
	ResumeOption *string `json:"ResumeOption,omitnil,omitempty" name:"ResumeOption"`
}

type ResumeMigrateJobRequest struct {
	*tchttp.BaseRequest
	
	// Data migration task ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// Task resumption mode. Valid values: `clearData` (Clearing the target instance data); `overwrite` (Executing the task in overwrite mode); `normal` (Following the normal process without additional operations). `clearData` and `overwrite` are only valid for Redis links and `normal` for non-Redis links.
	ResumeOption *string `json:"ResumeOption,omitnil,omitempty" name:"ResumeOption"`
}

func (r *ResumeMigrateJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResumeMigrateJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	delete(f, "ResumeOption")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResumeMigrateJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResumeMigrateJobResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ResumeMigrateJobResponse struct {
	*tchttp.BaseResponse
	Response *ResumeMigrateJobResponseParams `json:"Response"`
}

func (r *ResumeMigrateJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResumeMigrateJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResumeSyncJobRequestParams struct {
	// Sync task ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

type ResumeSyncJobRequest struct {
	*tchttp.BaseRequest
	
	// Sync task ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

func (r *ResumeSyncJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResumeSyncJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResumeSyncJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResumeSyncJobResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ResumeSyncJobResponse struct {
	*tchttp.BaseResponse
	Response *ResumeSyncJobResponseParams `json:"Response"`
}

func (r *ResumeSyncJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResumeSyncJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RoleItem struct {
	// Role name
	// Note: This field may return null, indicating that no valid values can be obtained.
	RoleName *string `json:"RoleName,omitnil,omitempty" name:"RoleName"`

	// Role name after migration
	// Note: This field may return null, indicating that no valid values can be obtained.
	NewRoleName *string `json:"NewRoleName,omitnil,omitempty" name:"NewRoleName"`
}

// Predefined struct for user
type SkipCheckItemRequestParams struct {
	// Data migration task ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// ID of the check step to be skipped, which is obtained in the `StepInfo[i].StepId` field returned by the `DescribeMigrationCheckJob` API, such as "OptimizeCheck".
	StepIds []*string `json:"StepIds,omitnil,omitempty" name:"StepIds"`

	// When the check fails due to foreign key dependency, you can use this field to specify whether to migrate the foreign key dependency. The foreign key dependency won’t be migrated when `StepIds` contains `ConstraintCheck` and the value of this field is `shield`, and will be migrated when `StepIds` contains `ConstraintCheck` and the value of this field is `migrate`.
	ForeignKeyFlag *string `json:"ForeignKeyFlag,omitnil,omitempty" name:"ForeignKeyFlag"`
}

type SkipCheckItemRequest struct {
	*tchttp.BaseRequest
	
	// Data migration task ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// ID of the check step to be skipped, which is obtained in the `StepInfo[i].StepId` field returned by the `DescribeMigrationCheckJob` API, such as "OptimizeCheck".
	StepIds []*string `json:"StepIds,omitnil,omitempty" name:"StepIds"`

	// When the check fails due to foreign key dependency, you can use this field to specify whether to migrate the foreign key dependency. The foreign key dependency won’t be migrated when `StepIds` contains `ConstraintCheck` and the value of this field is `shield`, and will be migrated when `StepIds` contains `ConstraintCheck` and the value of this field is `migrate`.
	ForeignKeyFlag *string `json:"ForeignKeyFlag,omitnil,omitempty" name:"ForeignKeyFlag"`
}

func (r *SkipCheckItemRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SkipCheckItemRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	delete(f, "StepIds")
	delete(f, "ForeignKeyFlag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SkipCheckItemRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SkipCheckItemResponseParams struct {
	// Message prompted for skipping the check item
	// Note: This field may return null, indicating that no valid values can be obtained.
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SkipCheckItemResponse struct {
	*tchttp.BaseResponse
	Response *SkipCheckItemResponseParams `json:"Response"`
}

func (r *SkipCheckItemResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SkipCheckItemResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SkipSyncCheckItemRequestParams struct {
	// Task ID, such as "sync-4ddgid2".
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// ID of the check step to be skipped, which is obtained in the `StepInfos[i].StepId` field returned by the `DescribeCheckSyncJobResult` API, such as "OptimizeCheck".
	StepIds []*string `json:"StepIds,omitnil,omitempty" name:"StepIds"`
}

type SkipSyncCheckItemRequest struct {
	*tchttp.BaseRequest
	
	// Task ID, such as "sync-4ddgid2".
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// ID of the check step to be skipped, which is obtained in the `StepInfos[i].StepId` field returned by the `DescribeCheckSyncJobResult` API, such as "OptimizeCheck".
	StepIds []*string `json:"StepIds,omitnil,omitempty" name:"StepIds"`
}

func (r *SkipSyncCheckItemRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SkipSyncCheckItemRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	delete(f, "StepIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SkipSyncCheckItemRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SkipSyncCheckItemResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SkipSyncCheckItemResponse struct {
	*tchttp.BaseResponse
	Response *SkipSyncCheckItemResponseParams `json:"Response"`
}

func (r *SkipSyncCheckItemResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SkipSyncCheckItemResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SkippedDetail struct {
	// Number of skipped tables
	// Note: This field may return null, indicating that no valid values can be obtained.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Details of skipped tables
	// Note: This field may return null, indicating that no valid values can be obtained.
	Items []*SkippedItem `json:"Items,omitnil,omitempty" name:"Items"`
}

type SkippedItem struct {
	// Database name
	// Note: This field may return null, indicating that no valid values can be obtained.
	Db *string `json:"Db,omitnil,omitempty" name:"Db"`

	// Table name
	// Note: This field may return null, indicating that no valid values can be obtained.
	Table *string `json:"Table,omitnil,omitempty" name:"Table"`

	// The cause why check is not initiated
	// Note: This field may return null, indicating that no valid values can be obtained.
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`
}

// Predefined struct for user
type StartCompareRequestParams struct {
	// Migration task ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// Data consistency check task ID in the format of `dts-8yv4w2i1-cmp-37skmii9`
	CompareTaskId *string `json:"CompareTaskId,omitnil,omitempty" name:"CompareTaskId"`
}

type StartCompareRequest struct {
	*tchttp.BaseRequest
	
	// Migration task ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// Data consistency check task ID in the format of `dts-8yv4w2i1-cmp-37skmii9`
	CompareTaskId *string `json:"CompareTaskId,omitnil,omitempty" name:"CompareTaskId"`
}

func (r *StartCompareRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartCompareRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	delete(f, "CompareTaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartCompareRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartCompareResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StartCompareResponse struct {
	*tchttp.BaseResponse
	Response *StartCompareResponseParams `json:"Response"`
}

func (r *StartCompareResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartCompareResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartMigrateJobRequestParams struct {
	// Data migration task ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

type StartMigrateJobRequest struct {
	*tchttp.BaseRequest
	
	// Data migration task ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

func (r *StartMigrateJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartMigrateJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartMigrateJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartMigrateJobResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StartMigrateJobResponse struct {
	*tchttp.BaseResponse
	Response *StartMigrateJobResponseParams `json:"Response"`
}

func (r *StartMigrateJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartMigrateJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartModifySyncJobRequestParams struct {
	// Sync task ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

type StartModifySyncJobRequest struct {
	*tchttp.BaseRequest
	
	// Sync task ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

func (r *StartModifySyncJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartModifySyncJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartModifySyncJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartModifySyncJobResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StartModifySyncJobResponse struct {
	*tchttp.BaseResponse
	Response *StartModifySyncJobResponseParams `json:"Response"`
}

func (r *StartModifySyncJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartModifySyncJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartSyncJobRequestParams struct {
	// Sync task ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

type StartSyncJobRequest struct {
	*tchttp.BaseRequest
	
	// Sync task ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

func (r *StartSyncJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartSyncJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartSyncJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartSyncJobResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StartSyncJobResponse struct {
	*tchttp.BaseResponse
	Response *StartSyncJobResponseParams `json:"Response"`
}

func (r *StartSyncJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartSyncJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StepDetailInfo struct {
	// Step number
	// Note: This field may return null, indicating that no valid values can be obtained.
	StepNo *uint64 `json:"StepNo,omitnil,omitempty" name:"StepNo"`

	// Step name
	// Note: This field may return null, indicating that no valid values can be obtained.
	StepName *string `json:"StepName,omitnil,omitempty" name:"StepName"`

	// Step ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	StepId *string `json:"StepId,omitnil,omitempty" name:"StepId"`

	// Step status. Valid values: `success`, `failed`, `running`, `notStarted`. Default value: `notStarted`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Start time of the current step in the format of "yyyy-mm-dd hh:mm:ss". If this field does not exist or is empty, it is meaningless.
	// Note: This field may return null, indicating that no valid values can be obtained.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Step error message
	// Note: This field may return null, indicating that no valid values can be obtained.
	StepMessage *string `json:"StepMessage,omitnil,omitempty" name:"StepMessage"`

	// Execution progress
	// Note: This field may return null, indicating that no valid values can be obtained.
	Percent *uint64 `json:"Percent,omitnil,omitempty" name:"Percent"`

	// Error message
	// Note: This field may return null, indicating that no valid values can be obtained.
	Errors []*ProcessStepTip `json:"Errors,omitnil,omitempty" name:"Errors"`

	// Warning
	// Note: This field may return null, indicating that no valid values can be obtained.
	Warnings []*ProcessStepTip `json:"Warnings,omitnil,omitempty" name:"Warnings"`
}

type StepInfo struct {
	// Step number
	// Note: This field may return null, indicating that no valid values can be obtained.
	StepNo *uint64 `json:"StepNo,omitnil,omitempty" name:"StepNo"`

	// Step name
	// Note: This field may return null, indicating that no valid values can be obtained.
	StepName *string `json:"StepName,omitnil,omitempty" name:"StepName"`

	// Step ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	StepId *string `json:"StepId,omitnil,omitempty" name:"StepId"`

	// Status of the current step. Valid values: `notStarted`, `running`, `failed`, `finished, `skipped`, `paused`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Step start time, which may be null.
	// Note: This field may return null, indicating that no valid values can be obtained.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Error message
	// Note: This field may return null, indicating that no valid values can be obtained.
	Errors []*StepTip `json:"Errors,omitnil,omitempty" name:"Errors"`

	// Warning message
	// Note: This field may return null, indicating that no valid values can be obtained.
	Warnings []*StepTip `json:"Warnings,omitnil,omitempty" name:"Warnings"`

	// Current step progress. Value range: 0-100. The value `-1` indicates that the progress of the current step is unavailable. Note: This field may return null, indicating that no valid values can be obtained.
	Progress *int64 `json:"Progress,omitnil,omitempty" name:"Progress"`
}

type StepTip struct {
	// Error code
	// Note: This field may return null, indicating that no valid values can be obtained.
	Code *string `json:"Code,omitnil,omitempty" name:"Code"`

	// Error message
	// Note: This field may return null, indicating that no valid values can be obtained.
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// Solution
	// Note: This field may return null, indicating that no valid values can be obtained.
	Solution *string `json:"Solution,omitnil,omitempty" name:"Solution"`

	// Help document
	// Note: This field may return null, indicating that no valid values can be obtained.
	HelpDoc *string `json:"HelpDoc,omitnil,omitempty" name:"HelpDoc"`

	// Whether the current step is skipped
	// Note: This field may return null, indicating that no valid values can be obtained.
	SkipInfo *string `json:"SkipInfo,omitnil,omitempty" name:"SkipInfo"`
}

// Predefined struct for user
type StopCompareRequestParams struct {
	// Migration task ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// Data consistency check task ID in the format of `dts-8yv4w2i1-cmp-37skmii9`
	CompareTaskId *string `json:"CompareTaskId,omitnil,omitempty" name:"CompareTaskId"`
}

type StopCompareRequest struct {
	*tchttp.BaseRequest
	
	// Migration task ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// Data consistency check task ID in the format of `dts-8yv4w2i1-cmp-37skmii9`
	CompareTaskId *string `json:"CompareTaskId,omitnil,omitempty" name:"CompareTaskId"`
}

func (r *StopCompareRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopCompareRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	delete(f, "CompareTaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopCompareRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopCompareResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StopCompareResponse struct {
	*tchttp.BaseResponse
	Response *StopCompareResponseParams `json:"Response"`
}

func (r *StopCompareResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopCompareResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopMigrateJobRequestParams struct {
	// Data migration task ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

type StopMigrateJobRequest struct {
	*tchttp.BaseRequest
	
	// Data migration task ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

func (r *StopMigrateJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopMigrateJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopMigrateJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopMigrateJobResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StopMigrateJobResponse struct {
	*tchttp.BaseResponse
	Response *StopMigrateJobResponseParams `json:"Response"`
}

func (r *StopMigrateJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopMigrateJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopSyncJobRequestParams struct {
	// Sync task ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

type StopSyncJobRequest struct {
	*tchttp.BaseRequest
	
	// Sync task ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

func (r *StopSyncJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopSyncJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopSyncJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopSyncJobResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StopSyncJobResponse struct {
	*tchttp.BaseResponse
	Response *StopSyncJobResponseParams `json:"Response"`
}

func (r *StopSyncJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopSyncJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SyncDBEndpointInfos struct {
	// Region of the database
	// Note: This field may return null, indicating that no valid values can be obtained.
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Instance network access type. Valid values: `extranet` (public network); `ipv6` (public IPv6); `cvm` (self-build on CVM); `dcg` (Direct Connect); `vpncloud` (VPN access); `cdb` (database); `ccn` (CCN); `intranet` (intranet); `vpc` (VPC). Note that the valid values are subject to the current link.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AccessType *string `json:"AccessType,omitnil,omitempty" name:"AccessType"`

	// Database type, such as `mysql`, `redis`, `mongodb`, `postgresql`, `mariadb`, and `percona`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DatabaseType *string `json:"DatabaseType,omitnil,omitempty" name:"DatabaseType"`

	// Database information
	// Note: This field may return null, indicating that no valid values can be obtained.
	Info []*Endpoint `json:"Info,omitnil,omitempty" name:"Info"`
}

type SyncDetailInfo struct {
	// Total number of steps
	// Note: This field may return null, indicating that no valid values can be obtained.
	StepAll *uint64 `json:"StepAll,omitnil,omitempty" name:"StepAll"`

	// Current step
	// Note: This field may return null, indicating that no valid values can be obtained.
	StepNow *uint64 `json:"StepNow,omitnil,omitempty" name:"StepNow"`

	// Overall progress
	// Note: This field may return null, indicating that no valid values can be obtained.
	Progress *int64 `json:"Progress,omitnil,omitempty" name:"Progress"`

	// Progress of the current step. Value range: 0-100. The value of `-1` indicates that you can't check the progress of the current step.
	// Note: u200dThis field may returnu200d·nullu200d, indicating that no valid values can be obtained.
	CurrentStepProgress *int64 `json:"CurrentStepProgress,omitnil,omitempty" name:"CurrentStepProgress"`

	// Data volume difference between the sync source and target
	// Note: This field may return null, indicating that no valid values can be obtained.
	MasterSlaveDistance *int64 `json:"MasterSlaveDistance,omitnil,omitempty" name:"MasterSlaveDistance"`

	// Time difference between the sync source and target
	// Note: This field may return null, indicating that no valid values can be obtained.
	SecondsBehindMaster *int64 `json:"SecondsBehindMaster,omitnil,omitempty" name:"SecondsBehindMaster"`

	// Overall description
	// Note: This field may return null, indicating that no valid values can be obtained.
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// Step details
	// Note: This field may return null, indicating that no valid values can be obtained.
	StepInfos []*StepInfo `json:"StepInfos,omitnil,omitempty" name:"StepInfos"`

	// Cause of the failure of initiating data consistency check
	// Note: This field may return null, indicating that no valid values can be obtained.
	CauseOfCompareDisable *string `json:"CauseOfCompareDisable,omitnil,omitempty" name:"CauseOfCompareDisable"`

	// Task error and the corresponding solution
	// Note: u200dThis field may returnu200d·nullu200d, indicating that no valid values can be obtained.
	ErrInfo *ErrInfo `json:"ErrInfo,omitnil,omitempty" name:"ErrInfo"`
}

type SyncJobInfo struct {
	// Sync task ID, such as `sync-btso140`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// Sync task name
	// Note: This field may return null, indicating that no valid values can be obtained.
	JobName *string `json:"JobName,omitnil,omitempty" name:"JobName"`

	// Billing mode. Valid values: `PostPay` (pay-as-you-go); `PrePay` (monthly subscription).
	// Note: This field may return null, indicating that no valid values can be obtained.
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// Running mode. Valid values: `Immediate`, `Timed`. Default value: `Immediate`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	RunMode *string `json:"RunMode,omitnil,omitempty" name:"RunMode"`

	// Expected execution time in the format of `yyyy-mm-dd hh:mm:ss`
	// Note: This field may return null, indicating that no valid values can be obtained.
	ExpectRunTime *string `json:"ExpectRunTime,omitnil,omitempty" name:"ExpectRunTime"`

	// All supported operations
	// Note: This field may return null, indicating that no valid values can be obtained.
	AllActions []*string `json:"AllActions,omitnil,omitempty" name:"AllActions"`

	// Operations that can be performed under the current status
	// Note: This field may return null, indicating that no valid values can be obtained.
	Actions []*string `json:"Actions,omitnil,omitempty" name:"Actions"`

	// Sync options
	// Note: This field may return null, indicating that no valid values can be obtained.
	Options *Options `json:"Options,omitnil,omitempty" name:"Options"`

	// Sync database/table objects
	// Note: This field may return null, indicating that no valid values can be obtained.
	Objects *Objects `json:"Objects,omitnil,omitempty" name:"Objects"`

	// Task specification
	// Note: This field may return null, indicating that no valid values can be obtained.
	Specification *string `json:"Specification,omitnil,omitempty" name:"Specification"`

	// Expiration time in the format of `yyyy-mm-dd hh:mm:ss`
	// Note: This field may return null, indicating that no valid values can be obtained.
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// Source database region, such as `ap-guangzhou`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SrcRegion *string `json:"SrcRegion,omitnil,omitempty" name:"SrcRegion"`

	// Source database type, such as `mysql`, `cynosdbmysql`, `tdapg`, `tdpg`, and `tdsqlmysql`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SrcDatabaseType *string `json:"SrcDatabaseType,omitnil,omitempty" name:"SrcDatabaseType"`

	// Source database access type. Valid values: `cdb` (database); `cvm` (self-build on CVM); `vpc` (VPC); `extranet` (public network); `vpncloud` (VPN access); `dcg` (Direct Connect); `ccn` (CCN); `intranet` (intranet).
	// Note: This field may return null, indicating that no valid values can be obtained.
	SrcAccessType *string `json:"SrcAccessType,omitnil,omitempty" name:"SrcAccessType"`

	// Source database information. This parameter is used by single-node databases.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SrcInfo *Endpoint `json:"SrcInfo,omitnil,omitempty" name:"SrcInfo"`

	// Valid values: `cluster`, `single`. `single`: For single-node source databases; `cluster`: For multi-node source databases.
	// Note: u200dThis field may returnu200d·nullu200d, indicating that no valid values can be obtained.
	SrcNodeType *string `json:"SrcNodeType,omitnil,omitempty" name:"SrcNodeType"`

	// Source database information. This parameter is used for multi-node databases.
	// Note: u200dThis field may returnu200d·nullu200d, indicating that no valid values can be obtained.
	SrcInfos *SyncDBEndpointInfos `json:"SrcInfos,omitnil,omitempty" name:"SrcInfos"`

	// Target database region, such as `ap-guangzhou`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DstRegion *string `json:"DstRegion,omitnil,omitempty" name:"DstRegion"`

	// Target database type, such as `mysql`, `cynosdbmysql`, `tdapg`, `tdpg`, and `tdsqlmysql`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DstDatabaseType *string `json:"DstDatabaseType,omitnil,omitempty" name:"DstDatabaseType"`

	// Target database access type. Valid values: `cdb` (database); `cvm` (self-build on CVM); `vpc` (VPC); `extranet` (public network); `vpncloud` (VPN access); `dcg` (Direct Connect); `ccn` (CCN); `intranet` (intranet).
	// Note: This field may return null, indicating that no valid values can be obtained.
	DstAccessType *string `json:"DstAccessType,omitnil,omitempty" name:"DstAccessType"`

	// Target database information. This parameter is used by single-node databases.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DstInfo *Endpoint `json:"DstInfo,omitnil,omitempty" name:"DstInfo"`

	// Valid values: `cluster`, `single`. `single`: For single-node target databases; `cluster`: For multi-node target databases.
	// Note: u200dThis field may returnu200d·nullu200d, indicating that no valid values can be obtained.
	DstNodeType *string `json:"DstNodeType,omitnil,omitempty" name:"DstNodeType"`

	// Target database information. This parameter is used for multi-node databases.
	// Note: u200dThis field may returnu200d·nullu200d, indicating that no valid values can be obtained.
	DstInfos *SyncDBEndpointInfos `json:"DstInfos,omitnil,omitempty" name:"DstInfos"`

	// Creation time in the format of `yyyy-mm-dd hh:mm:ss`
	// Note: This field may return null, indicating that no valid values can be obtained.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Start time in the format of `yyyy-mm-dd hh:mm:ss`
	// Note: This field may return null, indicating that no valid values can be obtained.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Task status. Valid values: `UnInitialized`, `Initialized`, `Checking`, `CheckPass`, `CheckNotPass`, `ReadyRunning`, `Running`, `Pausing`, `Paused`, `Stopping`, `Stopped`, `ResumableErr`, `Resuming`, `Failed`, `Released`, `Resetting`, `Unknown`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// End time in the format of `yyyy-mm-dd hh:mm:ss`
	// Note: This field may return null, indicating that no valid values can be obtained.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Tag information
	// Note: This field may return null, indicating that no valid values can be obtained.
	Tags []*TagItem `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Step information of the sync task
	// Note: This field may return null, indicating that no valid values can be obtained.
	Detail *SyncDetailInfo `json:"Detail,omitnil,omitempty" name:"Detail"`

	// Billing status. Valid values: `Normal`, `Resizing`, `Renewing`, `Isolating`, `Isolated`, `Offlining`, `Offlined`, `NotBilled`, `Recovering`, `PostPay2Prepaying`, `PrePay2Postpaying`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TradeStatus *string `json:"TradeStatus,omitnil,omitempty" name:"TradeStatus"`

	// Sync link specification, such as `micro`, `small`, `medium`, and `large`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceClass *string `json:"InstanceClass,omitnil,omitempty" name:"InstanceClass"`

	// Auto-renewal flag, which takes effect if `PayMode` is `PrePay`. Valid values: `1` (auto-renewal enabled); `0` (auto-renewal disabled).
	// Note: This field may return null, indicating that no valid values can be obtained.
	AutoRenew *uint64 `json:"AutoRenew,omitnil,omitempty" name:"AutoRenew"`

	// Deletion time in the format of `yyyy-mm-dd hh:mm:ss`
	// Note: This field may return null, indicating that no valid values can be obtained.
	OfflineTime *string `json:"OfflineTime,omitnil,omitempty" name:"OfflineTime"`

	// Settings of automatic retry time
	// Note: This field may return null, indicating that no valid values can be obtained.
	AutoRetryTimeRangeMinutes *int64 `json:"AutoRetryTimeRangeMinutes,omitnil,omitempty" name:"AutoRetryTimeRangeMinutes"`

	// Whether the task can be reentered in the full export stage. Valid values: `yes`, `no`. `yes`: The current task can be reentered. `no`: The current task is in the full export stage which cannot be reentered. If the value of this parameter is `no`, the checkpoint restart is not supported when the task is restarted in the export stage.
	// Note: u200dThis field may returnu200d·nullu200d, indicating that no valid values can be obtained.
	DumperResumeCtrl *string `json:"DumperResumeCtrl,omitnil,omitempty" name:"DumperResumeCtrl"`
}

type Table struct {
	// Table name
	// Note: This field may return null, indicating that no valid values can be obtained.
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// New table name
	// Note: This field may return null, indicating that no valid values can be obtained.
	NewTableName *string `json:"NewTableName,omitnil,omitempty" name:"NewTableName"`

	// Filter condition
	// Note: This field may return null, indicating that no valid values can be obtained.
	FilterCondition *string `json:"FilterCondition,omitnil,omitempty" name:"FilterCondition"`

	// The temp tables to be synced. This parameter is mutually exclusive with `NewTableName`. It is valid only when the configured sync objects are table-level ones and `TableEditMode` is `pt`. To sync temp tables generated when pt-osc or other tools are used during the sync process, you must configure this parameter first. For example, if you want to perform the pt-osc operation on a table named "t1", configure this parameter as ["\_t1\_new","\_t1\_old"]; to perform the gh-ost operation on t1, configure it as ["\_t1\_ghc","\_t1\_gho","\_t1\_del"]. Temp tables generated by pt-osc and gh-ost operations can be configured at the same time.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TmpTables []*string `json:"TmpTables,omitnil,omitempty" name:"TmpTables"`

	// Table editing type. Valid values: `rename` (table mapping); `pt` (additional table sync).
	// Note: This field may return null, indicating that no valid values can be obtained.
	TableEditMode *string `json:"TableEditMode,omitnil,omitempty" name:"TableEditMode"`
}

type TableItem struct {
	// Name of the migrated table, which is case-sensitive
	// Note: This field may return null, indicating that no valid values can be obtained.
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// New name of the migrated table. This parameter is required when `TableEditMode` is `rename`. It is mutually exclusive with `TmpTables`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	NewTableName *string `json:"NewTableName,omitnil,omitempty" name:"NewTableName"`

	// The temp tables to be migrated. This parameter is mutually exclusive with `NewTableName`. It is valid only when the configured migration objects are table-level ones and `TableEditMode` is `pt`. To migrate temp tables generated when pt-osc or other tools are used during the migration process, you must configure this parameter first. For example, if you want to perform the pt-osc operation on a table named "t1", configure this parameter as ["_t1_new","_t1_old"]; to perform the gh-ost operation on t1, configure it as ["_t1_ghc","_t1_gho","_t1_del"]. Temp tables generated by pt-osc and gh-ost operations can be configured at the same time.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TmpTables []*string `json:"TmpTables,omitnil,omitempty" name:"TmpTables"`

	// Table editing type. Valid values: `rename` (table mapping); `pt` (additional table sync).
	// Note: This field may return null, indicating that no valid values can be obtained.
	TableEditMode *string `json:"TableEditMode,omitnil,omitempty" name:"TableEditMode"`
}

type TagFilter struct {
	// Tag key
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// Tag value
	TagValue []*string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

type TagItem struct {
	// Tag key
	// Note: This field may return null, indicating that no valid values can be obtained.
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// Tag value
	// Note: This field may return null, indicating that no valid values can be obtained.
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

type TopicRule struct {
	// Topic name
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// Topic partitioning policy. If the topic sync policy is delivering data to multiple custom topics (`TopicType` = `Multi`), the value of this parameter is `Random` (deliver to a random partition). If the topic sync policy is delivering all data to a single topic (`TopicType` = `Single`), this parameter has three valid values: `AllInPartitionZero` (deliver all data to partition0), `PartitionByTable` (partition by table name), `PartitionByTableAndKey` (partition by table name and primary key).
	PartitionType *string `json:"PartitionType,omitnil,omitempty" name:"PartitionType"`

	// Database name matching rule. This parameter takes effect only when `TopicType` is `Multi`. Valid values: `Regular` (match by regex), `Default` (default rule for the remaining databases that cannot be matched by regex). The default rule must be included in the array of matching rules.
	DbMatchMode *string `json:"DbMatchMode,omitnil,omitempty" name:"DbMatchMode"`

	// Database name, which can only be matched by regex when `TopicType` is `Multi` and `DbMatchMode` is `Regular`.
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// Table name matching rule. This parameter takes effect only when `TopicType` is `Multi`. Valid values: `Regular` (match by regex), `Default` (default rule for the remaining databases that cannot be matched by regex). The default rule must be included in the array of matching rules.
	TableMatchMode *string `json:"TableMatchMode,omitnil,omitempty" name:"TableMatchMode"`

	// Table name, which can only be matched by regex when `TopicType` is `Multi` and `DbMatchMode` is `Regular`.
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`
}

type TradeInfo struct {
	// Order number
	// Note: This field may return null, indicating that no valid values can be obtained.
	DealName *string `json:"DealName,omitnil,omitempty" name:"DealName"`

	// Last order number
	// Note: This field may return null, indicating that no valid values can be obtained.
	LastDealName *string `json:"LastDealName,omitnil,omitempty" name:"LastDealName"`

	// Instance specification. Valid values: `micro`, `small`, `medium`, `large`, `xlarge`, `2xlarge`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceClass *string `json:"InstanceClass,omitnil,omitempty" name:"InstanceClass"`

	// Task billing status. Valid values: `normal` (billed or to be billed); `resizing` (adjusting configuration); `reversing` (topping up, which is a short status); `isolating` (isolating, which is a short status); `isolated` (isolated); `offlining` (deleting); `offlined` (deleted); `notBilled` (not billed).
	// Note: This field may return null, indicating that no valid values can be obtained.
	TradeStatus *string `json:"TradeStatus,omitnil,omitempty" name:"TradeStatus"`

	// Expiration time in the format of "yyyy-mm-dd hh:mm:ss"
	// Note: This field may return null, indicating that no valid values can be obtained.
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// Deletion time in the format of "yyyy-mm-dd hh:mm:ss"
	// Note: This field may return null, indicating that no valid values can be obtained.
	OfflineTime *string `json:"OfflineTime,omitnil,omitempty" name:"OfflineTime"`

	// Isolation time in the format of "yyyy-mm-dd hh:mm:ss"
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsolateTime *string `json:"IsolateTime,omitnil,omitempty" name:"IsolateTime"`

	// The cause of deletion
	// Note: This field may return null, indicating that no valid values can be obtained.
	OfflineReason *string `json:"OfflineReason,omitnil,omitempty" name:"OfflineReason"`

	// The cause of isolation
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsolateReason *string `json:"IsolateReason,omitnil,omitempty" name:"IsolateReason"`

	// Billing mode. Valid values: `postpay` (postpaid); `prepay` (prepaid).
	// Note: This field may return null, indicating that no valid values can be obtained.
	PayType *string `json:"PayType,omitnil,omitempty" name:"PayType"`

	// Task billing type. Valid values: `billing` (billed); `notBilling` (free); `promotions` (in promotion).
	// Note: This field may return null, indicating that no valid values can be obtained.
	BillingType *string `json:"BillingType,omitnil,omitempty" name:"BillingType"`
}

type View struct {
	// View name
	// Note: This field may return null, indicating that no valid values can be obtained.
	ViewName *string `json:"ViewName,omitnil,omitempty" name:"ViewName"`

	// Reserved field. Currently, a view cannot be renamed. Note: This field may return null, indicating that no valid values can be obtained.
	NewViewName *string `json:"NewViewName,omitnil,omitempty" name:"NewViewName"`
}

type ViewItem struct {
	// View name
	// Note: This field may return null, indicating that no valid values can be obtained.
	ViewName *string `json:"ViewName,omitnil,omitempty" name:"ViewName"`

	// View name after migration
	// Note: This field may return null, indicating that no valid values can be obtained.
	NewViewName *string `json:"NewViewName,omitnil,omitempty" name:"NewViewName"`
}