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

package v20250806

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

type AlarmGroup struct {
	// Specifies the list of Alarm escalation recipient ids.
	// If the Alarm recipient or supervisor does not confirm the Alarm within the Alarm interval, an Alarm will be sent to the next-level superior.
	AlarmEscalationRecipientIds []*string `json:"AlarmEscalationRecipientIds,omitnil,omitempty" name:"AlarmEscalationRecipientIds"`

	// Escalation interval for alarms.
	AlarmEscalationInterval *int64 `json:"AlarmEscalationInterval,omitnil,omitempty" name:"AlarmEscalationInterval"`

	// Alarm notification fatigue configuration.
	NotificationFatigue *NotificationFatigue `json:"NotificationFatigue,omitnil,omitempty" name:"NotificationFatigue"`

	// Alarm channel. valid values: 1. mail, 2. sms, 3. wechat, 4. voice, 5. wecom, 6. Http, 7. wecom group, 8. lark group, 9. dingtalk group, 10. Slack group, 11. Teams group (default: 1. mail). only select one channel.
	AlarmWays []*string `json:"AlarmWays,omitnil,omitempty" name:"AlarmWays"`

	// webhook url list for wecom group/lark group/dingtalk group/Slack group/Teams group.
	WebHooks []*AlarmWayWebHook `json:"WebHooks,omitnil,omitempty" name:"WebHooks"`

	// Alarm recipient type: 1. specified personnel, 2. task owner, 3. duty schedule (default: 1. specified personnel).
	AlarmRecipientType *int64 `json:"AlarmRecipientType,omitnil,omitempty" name:"AlarmRecipientType"`

	// Specifies different business ids based on AlarmRecipientType. valid values: 1 (designated personnel): Alarm recipient id list. 2 (task owner): no configuration required. 3 (duty schedule): schedule id list.
	AlarmRecipientIds []*string `json:"AlarmRecipientIds,omitnil,omitempty" name:"AlarmRecipientIds"`
}

type AlarmMessage struct {
	// Alarm message Id.
	AlarmMessageId *uint64 `json:"AlarmMessageId,omitnil,omitempty" name:"AlarmMessageId"`

	// Alarm time. the same Alarm may be sent multiple times, only the latest Alarm time is displayed.
	AlarmTime *string `json:"AlarmTime,omitnil,omitempty" name:"AlarmTime"`

	// Task name.
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// Task ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Instance data time of the task.
	CurRunDate *string `json:"CurRunDate,omitnil,omitempty" name:"CurRunDate"`

	// Indicates the Alarm cause.
	AlarmReason *string `json:"AlarmReason,omitnil,omitempty" name:"AlarmReason"`

	// Alarm level. 1. ordinary, 2. important, 3. critical.
	AlarmLevel *uint64 `json:"AlarmLevel,omitnil,omitempty" name:"AlarmLevel"`

	// Specifies the Id of the Alarm rule.
	AlarmRuleId *string `json:"AlarmRuleId,omitnil,omitempty" name:"AlarmRuleId"`

	// Alarm channel specifies the notification methods: 1. mail, 2. sms, 3. wechat, 4. voice, 5. wecom, 6. Http, 7. wecom group, 8. lark group, 9. dingtalk group, 10. Slack group, 11. Teams group (default: 1. mail).
	AlarmWays []*string `json:"AlarmWays,omitnil,omitempty" name:"AlarmWays"`

	// Alarm recipient
	AlarmRecipients []*string `json:"AlarmRecipients,omitnil,omitempty" name:"AlarmRecipients"`
}

type AlarmQuietInterval struct {
	// ISO standard. 1 means monday, 7 means sunday.
	DaysOfWeek []*uint64 `json:"DaysOfWeek,omitnil,omitempty" name:"DaysOfWeek"`

	// Start time. precision: hour/minute/second. format: HH:mm:ss.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time. precision: hour/minute/second. format: HH:mm:ss.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type AlarmRuleData struct {
	// Alarm Rule ID
	AlarmRuleId *string `json:"AlarmRuleId,omitnil,omitempty" name:"AlarmRuleId"`

	// Specifies the Alarm rule name.
	AlarmRuleName *string `json:"AlarmRuleName,omitnil,omitempty" name:"AlarmRuleName"`

	// Describes the Alarm rule.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Monitoring Object Type
	// 
	// Task-level Monitoring - Can be configured by Task / Workflow / Project:
	// 1 = Task (default)
	// 2 = Workflow
	// 3 = Project
	// 
	// Project-level Monitoring - Monitors overall task fluctuations within a project:
	// 7 = Project fluctuation monitoring alarm
	MonitorObjectType *int64 `json:"MonitorObjectType,omitnil,omitempty" name:"MonitorObjectType"`

	// Pass different business IDs depending on the value of MonitorType:
	// 
	// 1 (Task): MonitorObjectIds should contain a list of task IDs.
	// 
	// 2 (Workflow): MonitorObjectIds should contain a list of workflow IDs (workflow IDs can be obtained using the ListWorkflows API).
	// 
	// 3 (Project): MonitorObjectIds should contain a list of project IDs.
	MonitorObjectIds []*string `json:"MonitorObjectIds,omitnil,omitempty" name:"MonitorObjectIds"`

	// Alarm Rule Monitoring Types:
	// 
	// failure: Failure alarm
	// 
	// overtime: Timeout alarm
	// 
	// success: Success alarm
	// 
	// backTrackingOrRerunSuccess: Alarm when backfill/rerun succeeds
	// 
	// backTrackingOrRerunFailure: Alarm when backfill/rerun fails
	// 
	// projectFailureInstanceUpwardFluctuationAlarm: Alarm when the upward fluctuation rate of failed instances for the day exceeds the threshold
	// 
	// projectSuccessInstanceDownwardFluctuationAlarm: Alarm when the downward fluctuation rate of successful instances for the day exceeds the threshold
	// 
	// reconciliationFailure: Alarm when an offline reconciliation task fails
	// 
	// reconciliationOvertime: Alarm when an offline reconciliation task runs overtime
	// 
	// reconciliationMismatch: Alarm when the number of mismatched records in reconciliation exceeds the threshold
	AlarmTypes []*string `json:"AlarmTypes,omitnil,omitempty" name:"AlarmTypes"`

	// Whether the Alarm rule is enabled.
	// Valid values: 0 (disabled), 1 (enabled).
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Alarm Rule Configuration Information
	// 
	// Success Alarms - No configuration required;
	// 
	// Failure Alarms - Can be configured to trigger on the first failure or on all retry failures;
	// 
	// Timeout Alarms - Require configuration of the timeout type and timeout threshold;
	// 
	// Project Fluctuation Alarms - Require configuration of the fluctuation rate and the debounce cycle.
	AlarmRuleDetail *AlarmRuleDetail `json:"AlarmRuleDetail,omitnil,omitempty" name:"AlarmRuleDetail"`

	// Alarm level. 1. ordinary, 2. important, 3. critical.
	AlarmLevel *int64 `json:"AlarmLevel,omitnil,omitempty" name:"AlarmLevel"`

	// Specifies the id of the Alarm rule creator.
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// The Alarm rule bound to the bundle client. it is normal if empty, otherwise it corresponds to the rule bound to the bundle client.
	BundleId *string `json:"BundleId,omitnil,omitempty" name:"BundleId"`

	// bundleId is not empty. it indicates the bound client name.
	BundleInfo *string `json:"BundleInfo,omitnil,omitempty" name:"BundleInfo"`

	// Describes the Alarm recipient configuration list.
	AlarmGroups []*AlarmGroup `json:"AlarmGroups,omitnil,omitempty" name:"AlarmGroups"`
}

type AlarmRuleDetail struct {
	// Failure Trigger Condition
	// 
	// 1: Trigger on the first failure
	// 
	// 2: Trigger after all retries are completed (default)
	Trigger *int64 `json:"Trigger,omitnil,omitempty" name:"Trigger"`

	// Backfill/Rerun Trigger Condition
	// 
	// 1: Trigger on the first failure
	// 
	// 2: Trigger after all retries are completed
	DataBackfillOrRerunTrigger *int64 `json:"DataBackfillOrRerunTrigger,omitnil,omitempty" name:"DataBackfillOrRerunTrigger"`

	// Timeout configuration of the periodic instance.
	TimeOutExtInfo []*TimeOutStrategyInfo `json:"TimeOutExtInfo,omitnil,omitempty" name:"TimeOutExtInfo"`

	// Specifies the timeout configuration details for rerunning a backfill instance.
	DataBackfillOrRerunTimeOutExtInfo []*TimeOutStrategyInfo `json:"DataBackfillOrRerunTimeOutExtInfo,omitnil,omitempty" name:"DataBackfillOrRerunTimeOutExtInfo"`

	// Specifies the detail of Alarm configuration for project fluctuation.
	ProjectInstanceStatisticsAlarmInfoList []*ProjectInstanceStatisticsAlarmInfo `json:"ProjectInstanceStatisticsAlarmInfoList,omitnil,omitempty" name:"ProjectInstanceStatisticsAlarmInfoList"`

	// Describes the Alarm configuration information for offline integration reconciliation.
	ReconciliationExtInfo []*ReconciliationStrategyInfo `json:"ReconciliationExtInfo,omitnil,omitempty" name:"ReconciliationExtInfo"`
}

type AlarmWayWebHook struct {
	// Specifies the Alarm channel value.
	// 7. wecom group 8. lark group 9. dingtalk group 10. Slack group 11. Teams group.
	AlarmWay *string `json:"AlarmWay,omitnil,omitempty" name:"AlarmWay"`

	// webhook url list of the Alarm group.
	WebHooks []*string `json:"WebHooks,omitnil,omitempty" name:"WebHooks"`
}

type BackfillInstance struct {
	// Task name.
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// Task ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Specifies the instance data time.
	CurRunDate *string `json:"CurRunDate,omitnil,omitempty" name:"CurRunDate"`

	// Execution status.
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// Start time.
	// 
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time.
	// 
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Execution duration.
	// 
	CostTime *string `json:"CostTime,omitnil,omitempty" name:"CostTime"`
}

type BackfillInstanceCollection struct {
	// Page number
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Pagination size.
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Total pages
	TotalPageNumber *uint64 `json:"TotalPageNumber,omitnil,omitempty" name:"TotalPageNumber"`

	// Total number of records
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Specifies the backfill instance list.
	Items []*BackfillInstance `json:"Items,omitnil,omitempty" name:"Items"`
}

type ChildDependencyConfigPage struct {
	// Total number of results
	// 
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Total pages
	TotalPageNumber *uint64 `json:"TotalPageNumber,omitnil,omitempty" name:"TotalPageNumber"`

	// Page number.
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Pagination size.
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Paging data
	Items []*OpsTaskDepend `json:"Items,omitnil,omitempty" name:"Items"`
}

type CodeFile struct {
	// Script ID
	// 
	CodeFileId *string `json:"CodeFileId,omitnil,omitempty" name:"CodeFileId"`

	// Script name.
	CodeFileName *string `json:"CodeFileName,omitnil,omitempty" name:"CodeFileName"`

	// Specifies the script owner uin.
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// Specifies the script configuration.
	CodeFileConfig *CodeFileConfig `json:"CodeFileConfig,omitnil,omitempty" name:"CodeFileConfig"`

	// Specifies the script content.
	CodeFileContent *string `json:"CodeFileContent,omitnil,omitempty" name:"CodeFileContent"`

	// Latest operator.
	UpdateUserUin *string `json:"UpdateUserUin,omitnil,omitempty" name:"UpdateUserUin"`

	// Project ID.
	// 
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Update time. format: yyyy-MM-dd hh:MM:ss.
	// Note: This field may return null, indicating that no valid 
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// Creation time. format: yyyy-MM-dd hh:MM:ss.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Access permission: SHARED, PRIVATE.
	AccessScope *string `json:"AccessScope,omitnil,omitempty" name:"AccessScope"`

	// Full path of the node, /aaa/bbb/ccc.ipynb, consists of the name of each node.
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`
}

type CodeFileConfig struct {
	// Advanced running parameter variable replacement map-json String,String.
	Params *string `json:"Params,omitnil,omitempty" name:"Params"`

	// notebook kernel session information.
	NotebookSessionInfo *NotebookSessionInfo `json:"NotebookSessionInfo,omitnil,omitempty" name:"NotebookSessionInfo"`
}

type CodeFolderNode struct {
	// Unique identifier
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// Name
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`

	// folder type, script.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Whether it is a leaf node.
	IsLeaf *bool `json:"IsLeaf,omitnil,omitempty" name:"IsLeaf"`

	// Business parameters	
	// 	
	Params *string `json:"Params,omitnil,omitempty" name:"Params"`

	// Permission scope: SHARED, PRIVATE.
	AccessScope *string `json:"AccessScope,omitnil,omitempty" name:"AccessScope"`

	// Node path.
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// Specifies the uin of the responsible person for the directory/file.
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// Creator
	CreateUserUin *string `json:"CreateUserUin,omitnil,omitempty" name:"CreateUserUin"`

	// Specifies the permission of the current user for nodes.	
	NodePermission *string `json:"NodePermission,omitnil,omitempty" name:"NodePermission"`

	// Sub-node list
	// 
	Children []*CodeFolderNode `json:"Children,omitnil,omitempty" name:"Children"`
}

type CodeStudioFileActionResult struct {
	// Indicates whether the operation is successful. valid values: true (successful), false (unsuccessful).
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`

	// Folder ID.
	CodeFileId *string `json:"CodeFileId,omitnil,omitempty" name:"CodeFileId"`
}

type CodeStudioFolderActionResult struct {
	// Indicates whether the operation is successful. valid values: true (successful), false (unsuccessful).
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`

	// Folder ID.
	// 
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`
}

type CodeStudioFolderResult struct {
	// Folder ID.
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`
}

type CreateAlarmRuleData struct {
	// Specifies the unique id of the Alarm rule.
	AlarmRuleId *string `json:"AlarmRuleId,omitnil,omitempty" name:"AlarmRuleId"`
}

// Predefined struct for user
type CreateCodeFileRequestParams struct {
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Filename of the code file.
	CodeFileName *string `json:"CodeFileName,omitnil,omitempty" name:"CodeFileName"`

	// Parent folder path, such as /aaa/bbb/ccc. the path must start with a slash. use / for the root directory.
	ParentFolderPath *string `json:"ParentFolderPath,omitnil,omitempty" name:"ParentFolderPath"`

	// Specifies the code file configuration.
	CodeFileConfig *CodeFileConfig `json:"CodeFileConfig,omitnil,omitempty" name:"CodeFileConfig"`

	// Specifies the code file content.
	CodeFileContent *string `json:"CodeFileContent,omitnil,omitempty" name:"CodeFileContent"`
}

type CreateCodeFileRequest struct {
	*tchttp.BaseRequest
	
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Filename of the code file.
	CodeFileName *string `json:"CodeFileName,omitnil,omitempty" name:"CodeFileName"`

	// Parent folder path, such as /aaa/bbb/ccc. the path must start with a slash. use / for the root directory.
	ParentFolderPath *string `json:"ParentFolderPath,omitnil,omitempty" name:"ParentFolderPath"`

	// Specifies the code file configuration.
	CodeFileConfig *CodeFileConfig `json:"CodeFileConfig,omitnil,omitempty" name:"CodeFileConfig"`

	// Specifies the code file content.
	CodeFileContent *string `json:"CodeFileContent,omitnil,omitempty" name:"CodeFileContent"`
}

func (r *CreateCodeFileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCodeFileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "CodeFileName")
	delete(f, "ParentFolderPath")
	delete(f, "CodeFileConfig")
	delete(f, "CodeFileContent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCodeFileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCodeFileResponseParams struct {
	// Result
	Data *CodeFile `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCodeFileResponse struct {
	*tchttp.BaseResponse
	Response *CreateCodeFileResponseParams `json:"Response"`
}

func (r *CreateCodeFileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCodeFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCodeFolderRequestParams struct {
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Folder name.
	FolderName *string `json:"FolderName,omitnil,omitempty" name:"FolderName"`

	// Parent folder path, such as /aaa/bbb/ccc. the path must start with a slash. use / for the root directory.
	ParentFolderPath *string `json:"ParentFolderPath,omitnil,omitempty" name:"ParentFolderPath"`
}

type CreateCodeFolderRequest struct {
	*tchttp.BaseRequest
	
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Folder name.
	FolderName *string `json:"FolderName,omitnil,omitempty" name:"FolderName"`

	// Parent folder path, such as /aaa/bbb/ccc. the path must start with a slash. use / for the root directory.
	ParentFolderPath *string `json:"ParentFolderPath,omitnil,omitempty" name:"ParentFolderPath"`
}

func (r *CreateCodeFolderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCodeFolderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "FolderName")
	delete(f, "ParentFolderPath")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCodeFolderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCodeFolderResponseParams struct {
	// Indicates whether the operation is successful. valid values: true (successful), false (unsuccessful).
	Data *CodeStudioFolderResult `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCodeFolderResponse struct {
	*tchttp.BaseResponse
	Response *CreateCodeFolderResponseParams `json:"Response"`
}

func (r *CreateCodeFolderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCodeFolderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDataBackfillPlanRequestParams struct {
	// Project ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Backfill task collection.
	TaskIds []*string `json:"TaskIds,omitnil,omitempty" name:"TaskIds"`

	// Specifies the data time configuration for the backfill task.
	DataBackfillRangeList []*DataBackfillRange `json:"DataBackfillRangeList,omitnil,omitempty" name:"DataBackfillRangeList"`

	// Time zone, default UTC+8.
	TimeZone *string `json:"TimeZone,omitnil,omitempty" name:"TimeZone"`

	// Backfill plan name. if left empty, a string of characters is randomly generated by system.
	DataBackfillPlanName *string `json:"DataBackfillPlanName,omitnil,omitempty" name:"DataBackfillPlanName"`

	// Check parent task type. valid values: NONE (do not check ALL), ALL (check ALL upstream parent tasks), MAKE_SCOPE (only check in the currently selected tasks of the backfill plan). default: NONE (do not check).
	CheckParentType *string `json:"CheckParentType,omitnil,omitempty" name:"CheckParentType"`

	// Specifies whether to ignore event dependency for backfill. default true.
	SkipEventListening *bool `json:"SkipEventListening,omitnil,omitempty" name:"SkipEventListening"`

	// Custom workflow self-dependency. valid values: yes or no. if not configured, use the original workflow self-dependency.
	RedefineSelfWorkflowDependency *string `json:"RedefineSelfWorkflowDependency,omitnil,omitempty" name:"RedefineSelfWorkflowDependency"`

	// Customizes the degree of concurrency for instance running. if without configuring, use the existing self-dependent of the task.
	RedefineParallelNum *uint64 `json:"RedefineParallelNum,omitnil,omitempty" name:"RedefineParallelNum"`

	// Scheduling resource group id. if left empty, indicates usage of the original task scheduling execution resource group.
	SchedulerResourceGroupId *string `json:"SchedulerResourceGroupId,omitnil,omitempty" name:"SchedulerResourceGroupId"`

	// Integration task resource group id. indicates usage of the original task scheduling execution resource group if empty.
	IntegrationResourceGroupId *string `json:"IntegrationResourceGroupId,omitnil,omitempty" name:"IntegrationResourceGroupId"`

	// Custom parameter. re-specifies the task's parameters to facilitate the execution of new logic by replenished instances.
	RedefineParamList []*KVPair `json:"RedefineParamList,omitnil,omitempty" name:"RedefineParamList"`

	// Backfill Execution Order - The execution order for backfill instances based on their data time. Effective only when both conditions are met:
	// 
	// 1. Must be the same cycle task.
	// 
	// 2. Priority is given to dependency order. If no dependencies apply, execution follows the configured order.
	// 
	// Valid values:
	// 
	// -NORMAL: No specific order (default)
	// 
	// -ORDER: Execute in chronological order
	// 
	// -REVERSE: Execute in reverse chronological order
	DataTimeOrder *string `json:"DataTimeOrder,omitnil,omitempty" name:"DataTimeOrder"`

	// Backfill Instance Regeneration Cycle - If set, this will redefine the generation cycle of backfill task instances. Currently, only daily instances can be converted into instances generated on the first day of each month.
	// 
	// Valid value:
	// 
	// MONTH_CYCLE: Monthly
	RedefineCycleType *string `json:"RedefineCycleType,omitnil,omitempty" name:"RedefineCycleType"`
}

type CreateDataBackfillPlanRequest struct {
	*tchttp.BaseRequest
	
	// Project ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Backfill task collection.
	TaskIds []*string `json:"TaskIds,omitnil,omitempty" name:"TaskIds"`

	// Specifies the data time configuration for the backfill task.
	DataBackfillRangeList []*DataBackfillRange `json:"DataBackfillRangeList,omitnil,omitempty" name:"DataBackfillRangeList"`

	// Time zone, default UTC+8.
	TimeZone *string `json:"TimeZone,omitnil,omitempty" name:"TimeZone"`

	// Backfill plan name. if left empty, a string of characters is randomly generated by system.
	DataBackfillPlanName *string `json:"DataBackfillPlanName,omitnil,omitempty" name:"DataBackfillPlanName"`

	// Check parent task type. valid values: NONE (do not check ALL), ALL (check ALL upstream parent tasks), MAKE_SCOPE (only check in the currently selected tasks of the backfill plan). default: NONE (do not check).
	CheckParentType *string `json:"CheckParentType,omitnil,omitempty" name:"CheckParentType"`

	// Specifies whether to ignore event dependency for backfill. default true.
	SkipEventListening *bool `json:"SkipEventListening,omitnil,omitempty" name:"SkipEventListening"`

	// Custom workflow self-dependency. valid values: yes or no. if not configured, use the original workflow self-dependency.
	RedefineSelfWorkflowDependency *string `json:"RedefineSelfWorkflowDependency,omitnil,omitempty" name:"RedefineSelfWorkflowDependency"`

	// Customizes the degree of concurrency for instance running. if without configuring, use the existing self-dependent of the task.
	RedefineParallelNum *uint64 `json:"RedefineParallelNum,omitnil,omitempty" name:"RedefineParallelNum"`

	// Scheduling resource group id. if left empty, indicates usage of the original task scheduling execution resource group.
	SchedulerResourceGroupId *string `json:"SchedulerResourceGroupId,omitnil,omitempty" name:"SchedulerResourceGroupId"`

	// Integration task resource group id. indicates usage of the original task scheduling execution resource group if empty.
	IntegrationResourceGroupId *string `json:"IntegrationResourceGroupId,omitnil,omitempty" name:"IntegrationResourceGroupId"`

	// Custom parameter. re-specifies the task's parameters to facilitate the execution of new logic by replenished instances.
	RedefineParamList []*KVPair `json:"RedefineParamList,omitnil,omitempty" name:"RedefineParamList"`

	// Backfill Execution Order - The execution order for backfill instances based on their data time. Effective only when both conditions are met:
	// 
	// 1. Must be the same cycle task.
	// 
	// 2. Priority is given to dependency order. If no dependencies apply, execution follows the configured order.
	// 
	// Valid values:
	// 
	// -NORMAL: No specific order (default)
	// 
	// -ORDER: Execute in chronological order
	// 
	// -REVERSE: Execute in reverse chronological order
	DataTimeOrder *string `json:"DataTimeOrder,omitnil,omitempty" name:"DataTimeOrder"`

	// Backfill Instance Regeneration Cycle - If set, this will redefine the generation cycle of backfill task instances. Currently, only daily instances can be converted into instances generated on the first day of each month.
	// 
	// Valid value:
	// 
	// MONTH_CYCLE: Monthly
	RedefineCycleType *string `json:"RedefineCycleType,omitnil,omitempty" name:"RedefineCycleType"`
}

func (r *CreateDataBackfillPlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDataBackfillPlanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "TaskIds")
	delete(f, "DataBackfillRangeList")
	delete(f, "TimeZone")
	delete(f, "DataBackfillPlanName")
	delete(f, "CheckParentType")
	delete(f, "SkipEventListening")
	delete(f, "RedefineSelfWorkflowDependency")
	delete(f, "RedefineParallelNum")
	delete(f, "SchedulerResourceGroupId")
	delete(f, "IntegrationResourceGroupId")
	delete(f, "RedefineParamList")
	delete(f, "DataTimeOrder")
	delete(f, "RedefineCycleType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDataBackfillPlanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDataBackfillPlanResponseParams struct {
	// Specifies the creation result of the data backfill plan.
	Data *CreateDataReplenishmentPlan `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDataBackfillPlanResponse struct {
	*tchttp.BaseResponse
	Response *CreateDataBackfillPlanResponseParams `json:"Response"`
}

func (r *CreateDataBackfillPlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDataBackfillPlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDataReplenishmentPlan struct {
	// Specifies the backfill plan Id.
	DataBackfillPlanId *string `json:"DataBackfillPlanId,omitnil,omitempty" name:"DataBackfillPlanId"`
}

type CreateFolderResult struct {
	// Folder ID upon successful creation. error will be reported if creation failed.
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`
}

// Predefined struct for user
type CreateOpsAlarmRuleRequestParams struct {
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Specifies the Alarm rule name.
	AlarmRuleName *string `json:"AlarmRuleName,omitnil,omitempty" name:"AlarmRuleName"`

	// Monitoring Object Business ID List - Pass different business IDs depending on the value of MonitorType:
	// 
	// 1 (Task): MonitorObjectIds should contain a list of task IDs.
	// 
	// 2 (Workflow): MonitorObjectIds should contain a list of workflow IDs (workflow IDs can be obtained using the ListWorkflows API).
	// 
	// 3 (Project): MonitorObjectIds should contain a list of project IDs.
	MonitorObjectIds []*string `json:"MonitorObjectIds,omitnil,omitempty" name:"MonitorObjectIds"`

	// Alarm Rule Monitoring Types:
	// 
	// failure: Failure alarm
	// 
	// overtime: Timeout alarm
	// 
	// success: Success alarm
	// 
	// backTrackingOrRerunSuccess: Alarm when backfill/rerun succeeds
	// 
	// backTrackingOrRerunFailure: Alarm when backfill/rerun fails
	// 
	// projectFailureInstanceUpwardFluctuationAlarm: Alarm when the upward fluctuation rate of failed instances for the day exceeds the threshold
	// 
	// projectSuccessInstanceDownwardFluctuationAlarm: Alarm when the downward fluctuation rate of successful instances for the day exceeds the threshold
	// 
	// reconciliationFailure: Alarm when an offline reconciliation task fails
	// 
	// reconciliationOvertime: Alarm when an offline reconciliation task runs overtime
	// 
	// reconciliationMismatch: Alarm when the number of mismatched records in reconciliation exceeds the threshold
	AlarmTypes []*string `json:"AlarmTypes,omitnil,omitempty" name:"AlarmTypes"`

	// Alarm recipient configuration.
	AlarmGroups []*AlarmGroup `json:"AlarmGroups,omitnil,omitempty" name:"AlarmGroups"`

	// Monitoring Object Type
	// 
	// Task-level Monitoring - Can be configured by Task / Workflow / Project:
	// 1 = Task (default)
	// 2 = Workflow
	// 3 = Project
	// 
	// Project-level Monitoring - Monitors overall task fluctuations within a project:
	// 7 = Project fluctuation monitoring alarm
	MonitorObjectType *uint64 `json:"MonitorObjectType,omitnil,omitempty" name:"MonitorObjectType"`

	// Alarm Rule Configuration Information
	// 
	// Success Alarms - No configuration required;
	// 
	// Failure Alarms - Can be configured to trigger on the first failure or on all retry failures;
	// 
	// Timeout Alarms - Require configuration of the timeout type and timeout threshold;
	// 
	// Project Fluctuation Alarms - Require configuration of the fluctuation rate and the debounce cycle.
	AlarmRuleDetail *AlarmRuleDetail `json:"AlarmRuleDetail,omitnil,omitempty" name:"AlarmRuleDetail"`

	// Alarm level. 1. ordinary, 2. important, 3. critical (default: 1. ordinary).
	AlarmLevel *int64 `json:"AlarmLevel,omitnil,omitempty" name:"AlarmLevel"`

	// Describes the Alarm rule.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type CreateOpsAlarmRuleRequest struct {
	*tchttp.BaseRequest
	
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Specifies the Alarm rule name.
	AlarmRuleName *string `json:"AlarmRuleName,omitnil,omitempty" name:"AlarmRuleName"`

	// Monitoring Object Business ID List - Pass different business IDs depending on the value of MonitorType:
	// 
	// 1 (Task): MonitorObjectIds should contain a list of task IDs.
	// 
	// 2 (Workflow): MonitorObjectIds should contain a list of workflow IDs (workflow IDs can be obtained using the ListWorkflows API).
	// 
	// 3 (Project): MonitorObjectIds should contain a list of project IDs.
	MonitorObjectIds []*string `json:"MonitorObjectIds,omitnil,omitempty" name:"MonitorObjectIds"`

	// Alarm Rule Monitoring Types:
	// 
	// failure: Failure alarm
	// 
	// overtime: Timeout alarm
	// 
	// success: Success alarm
	// 
	// backTrackingOrRerunSuccess: Alarm when backfill/rerun succeeds
	// 
	// backTrackingOrRerunFailure: Alarm when backfill/rerun fails
	// 
	// projectFailureInstanceUpwardFluctuationAlarm: Alarm when the upward fluctuation rate of failed instances for the day exceeds the threshold
	// 
	// projectSuccessInstanceDownwardFluctuationAlarm: Alarm when the downward fluctuation rate of successful instances for the day exceeds the threshold
	// 
	// reconciliationFailure: Alarm when an offline reconciliation task fails
	// 
	// reconciliationOvertime: Alarm when an offline reconciliation task runs overtime
	// 
	// reconciliationMismatch: Alarm when the number of mismatched records in reconciliation exceeds the threshold
	AlarmTypes []*string `json:"AlarmTypes,omitnil,omitempty" name:"AlarmTypes"`

	// Alarm recipient configuration.
	AlarmGroups []*AlarmGroup `json:"AlarmGroups,omitnil,omitempty" name:"AlarmGroups"`

	// Monitoring Object Type
	// 
	// Task-level Monitoring - Can be configured by Task / Workflow / Project:
	// 1 = Task (default)
	// 2 = Workflow
	// 3 = Project
	// 
	// Project-level Monitoring - Monitors overall task fluctuations within a project:
	// 7 = Project fluctuation monitoring alarm
	MonitorObjectType *uint64 `json:"MonitorObjectType,omitnil,omitempty" name:"MonitorObjectType"`

	// Alarm Rule Configuration Information
	// 
	// Success Alarms - No configuration required;
	// 
	// Failure Alarms - Can be configured to trigger on the first failure or on all retry failures;
	// 
	// Timeout Alarms - Require configuration of the timeout type and timeout threshold;
	// 
	// Project Fluctuation Alarms - Require configuration of the fluctuation rate and the debounce cycle.
	AlarmRuleDetail *AlarmRuleDetail `json:"AlarmRuleDetail,omitnil,omitempty" name:"AlarmRuleDetail"`

	// Alarm level. 1. ordinary, 2. important, 3. critical (default: 1. ordinary).
	AlarmLevel *int64 `json:"AlarmLevel,omitnil,omitempty" name:"AlarmLevel"`

	// Describes the Alarm rule.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *CreateOpsAlarmRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOpsAlarmRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "AlarmRuleName")
	delete(f, "MonitorObjectIds")
	delete(f, "AlarmTypes")
	delete(f, "AlarmGroups")
	delete(f, "MonitorObjectType")
	delete(f, "AlarmRuleDetail")
	delete(f, "AlarmLevel")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateOpsAlarmRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOpsAlarmRuleResponseParams struct {
	// Specifies the unique id of the Alarm rule.
	Data *CreateAlarmRuleData `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateOpsAlarmRuleResponse struct {
	*tchttp.BaseResponse
	Response *CreateOpsAlarmRuleResponseParams `json:"Response"`
}

func (r *CreateOpsAlarmRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOpsAlarmRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateResourceFileRequestParams struct {
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Resource file name. should be consistent with the uploaded file name.
	ResourceName *string `json:"ResourceName,omitnil,omitempty" name:"ResourceName"`

	// Bucket name. can be obtained from the GetResourceCosPath api.
	BucketName *string `json:"BucketName,omitnil,omitempty" name:"BucketName"`

	// BucketName specifies the cos storage bucket region.
	CosRegion *string `json:"CosRegion,omitnil,omitempty" name:"CosRegion"`

	// Upload path for resource files in the project. value example: /wedata/qxxxm/. root directory, please use /.
	ParentFolderPath *string `json:"ParentFolderPath,omitnil,omitempty" name:"ParentFolderPath"`

	// -Upload file and manual entry are two methods, choose one. if both are provided, the sequence is file > manual entry.
	// -The manually entered value must be an existing cos path. /datastudio/resource/ is the fixed prefix. projectId is the project ID. import a specific value. parentFolderPath is the folder path. name is the file name. value example: /datastudio/resource/projectId/parentFolderPath/name. 
	ResourceFile *string `json:"ResourceFile,omitnil,omitempty" name:"ResourceFile"`

	// Bundle Client ID.
	BundleId *string `json:"BundleId,omitnil,omitempty" name:"BundleId"`

	// bundle client info.
	BundleInfo *string `json:"BundleInfo,omitnil,omitempty" name:"BundleInfo"`
}

type CreateResourceFileRequest struct {
	*tchttp.BaseRequest
	
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Resource file name. should be consistent with the uploaded file name.
	ResourceName *string `json:"ResourceName,omitnil,omitempty" name:"ResourceName"`

	// Bucket name. can be obtained from the GetResourceCosPath api.
	BucketName *string `json:"BucketName,omitnil,omitempty" name:"BucketName"`

	// BucketName specifies the cos storage bucket region.
	CosRegion *string `json:"CosRegion,omitnil,omitempty" name:"CosRegion"`

	// Upload path for resource files in the project. value example: /wedata/qxxxm/. root directory, please use /.
	ParentFolderPath *string `json:"ParentFolderPath,omitnil,omitempty" name:"ParentFolderPath"`

	// -Upload file and manual entry are two methods, choose one. if both are provided, the sequence is file > manual entry.
	// -The manually entered value must be an existing cos path. /datastudio/resource/ is the fixed prefix. projectId is the project ID. import a specific value. parentFolderPath is the folder path. name is the file name. value example: /datastudio/resource/projectId/parentFolderPath/name. 
	ResourceFile *string `json:"ResourceFile,omitnil,omitempty" name:"ResourceFile"`

	// Bundle Client ID.
	BundleId *string `json:"BundleId,omitnil,omitempty" name:"BundleId"`

	// bundle client info.
	BundleInfo *string `json:"BundleInfo,omitnil,omitempty" name:"BundleInfo"`
}

func (r *CreateResourceFileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateResourceFileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "ResourceName")
	delete(f, "BucketName")
	delete(f, "CosRegion")
	delete(f, "ParentFolderPath")
	delete(f, "ResourceFile")
	delete(f, "BundleId")
	delete(f, "BundleInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateResourceFileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateResourceFileResponseParams struct {
	// Create resource file result.
	Data *CreateResourceFileResult `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateResourceFileResponse struct {
	*tchttp.BaseResponse
	Response *CreateResourceFileResponseParams `json:"Response"`
}

func (r *CreateResourceFileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateResourceFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateResourceFileResult struct {
	// Resource file ID.
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`
}

// Predefined struct for user
type CreateResourceFolderRequestParams struct {
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Parent folder's absolute path. Example: /wedata/test. Use / to represent the root directory.
	ParentFolderPath *string `json:"ParentFolderPath,omitnil,omitempty" name:"ParentFolderPath"`

	// Folder name.
	FolderName *string `json:"FolderName,omitnil,omitempty" name:"FolderName"`
}

type CreateResourceFolderRequest struct {
	*tchttp.BaseRequest
	
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Parent folder's absolute path. Example: /wedata/test. Use / to represent the root directory.
	ParentFolderPath *string `json:"ParentFolderPath,omitnil,omitempty" name:"ParentFolderPath"`

	// Folder name.
	FolderName *string `json:"FolderName,omitnil,omitempty" name:"FolderName"`
}

func (r *CreateResourceFolderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateResourceFolderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "ParentFolderPath")
	delete(f, "FolderName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateResourceFolderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateResourceFolderResponseParams struct {
	// Creation result of the folder. Error will be reported if failed to create.
	Data *CreateFolderResult `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateResourceFolderResponse struct {
	*tchttp.BaseResponse
	Response *CreateResourceFolderResponseParams `json:"Response"`
}

func (r *CreateResourceFolderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateResourceFolderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSQLFolderRequestParams struct {
	// Folder name.
	FolderName *string `json:"FolderName,omitnil,omitempty" name:"FolderName"`

	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Parent folder path, /aaa/bbb/ccc. path header must include a slash. pass / to query the root directory.
	ParentFolderPath *string `json:"ParentFolderPath,omitnil,omitempty" name:"ParentFolderPath"`

	// Access permission: SHARED, PRIVATE.
	AccessScope *string `json:"AccessScope,omitnil,omitempty" name:"AccessScope"`
}

type CreateSQLFolderRequest struct {
	*tchttp.BaseRequest
	
	// Folder name.
	FolderName *string `json:"FolderName,omitnil,omitempty" name:"FolderName"`

	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Parent folder path, /aaa/bbb/ccc. path header must include a slash. pass / to query the root directory.
	ParentFolderPath *string `json:"ParentFolderPath,omitnil,omitempty" name:"ParentFolderPath"`

	// Access permission: SHARED, PRIVATE.
	AccessScope *string `json:"AccessScope,omitnil,omitempty" name:"AccessScope"`
}

func (r *CreateSQLFolderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSQLFolderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FolderName")
	delete(f, "ProjectId")
	delete(f, "ParentFolderPath")
	delete(f, "AccessScope")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSQLFolderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSQLFolderResponseParams struct {
	// Indicates whether the operation is successful. valid values: true (successful), false (unsuccessful).
	Data *SqlCreateResult `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateSQLFolderResponse struct {
	*tchttp.BaseResponse
	Response *CreateSQLFolderResponseParams `json:"Response"`
}

func (r *CreateSQLFolderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSQLFolderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSQLScriptRequestParams struct {
	// Script name.
	ScriptName *string `json:"ScriptName,omitnil,omitempty" name:"ScriptName"`

	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Parent folder path, /aaa/bbb/ccc. root directory is empty string or /.
	ParentFolderPath *string `json:"ParentFolderPath,omitnil,omitempty" name:"ParentFolderPath"`

	// Specifies the script configuration for data exploration.
	ScriptConfig *SQLScriptConfig `json:"ScriptConfig,omitnil,omitempty" name:"ScriptConfig"`

	// Specifies the script content, if any, needs to be base64 encoded.
	ScriptContent *string `json:"ScriptContent,omitnil,omitempty" name:"ScriptContent"`

	// Access permission: SHARED, PRIVATE.
	AccessScope *string `json:"AccessScope,omitnil,omitempty" name:"AccessScope"`
}

type CreateSQLScriptRequest struct {
	*tchttp.BaseRequest
	
	// Script name.
	ScriptName *string `json:"ScriptName,omitnil,omitempty" name:"ScriptName"`

	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Parent folder path, /aaa/bbb/ccc. root directory is empty string or /.
	ParentFolderPath *string `json:"ParentFolderPath,omitnil,omitempty" name:"ParentFolderPath"`

	// Specifies the script configuration for data exploration.
	ScriptConfig *SQLScriptConfig `json:"ScriptConfig,omitnil,omitempty" name:"ScriptConfig"`

	// Specifies the script content, if any, needs to be base64 encoded.
	ScriptContent *string `json:"ScriptContent,omitnil,omitempty" name:"ScriptContent"`

	// Access permission: SHARED, PRIVATE.
	AccessScope *string `json:"AccessScope,omitnil,omitempty" name:"AccessScope"`
}

func (r *CreateSQLScriptRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSQLScriptRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ScriptName")
	delete(f, "ProjectId")
	delete(f, "ParentFolderPath")
	delete(f, "ScriptConfig")
	delete(f, "ScriptContent")
	delete(f, "AccessScope")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSQLScriptRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSQLScriptResponseParams struct {
	// Result
	Data *SQLScript `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateSQLScriptResponse struct {
	*tchttp.BaseResponse
	Response *CreateSQLScriptResponseParams `json:"Response"`
}

func (r *CreateSQLScriptResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSQLScriptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateTaskBaseAttribute struct {
	// Task name.
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// Task type ID:
	// 
	// * 21:JDBC SQL
	// * 23:TDSQL-PostgreSQL
	// * 26:OfflineSynchronization
	// * 30:Python
	// * 31:PySpark
	// * 32:DLC SQL
	// * 33:Impala
	// * 34:Hive SQL
	// * 35:Shell
	// * 36:Spark SQL
	// * 38:Shell Form Mode
	// * 39:Spark
	// * 40:TCHouse-P
	// * 41:Kettle
	// * 42:Tchouse-X
	// * 43:TCHouse-X SQL
	// * 46:DLC Spark
	// * 47:TiOne
	// * 48:Trino
	// * 50:DLC PySpark
	// * 92:MapReduce
	// * 130:Branch Node
	// * 131:Merged Node
	// * 132:Notebook
	// * 133:SSH
	// * 134:StarRocks
	// * 137:For-each
	// * 138:Setats SQL
	TaskTypeId *string `json:"TaskTypeId,omitnil,omitempty" name:"TaskTypeId"`

	// Workflow ID.
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// Task owner ID. defaults to the current user.
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// Task description
	TaskDescription *string `json:"TaskDescription,omitnil,omitempty" name:"TaskDescription"`
}

type CreateTaskConfiguration struct {
	// Resource Group ID: Must be obtained via DescribeNormalSchedulerExecutorGroups API to get the ExecutorGroupId.
	ResourceGroup *string `json:"ResourceGroup,omitnil,omitempty" name:"ResourceGroup"`

	// Base64 encoding of the code content.
	CodeContent *string `json:"CodeContent,omitnil,omitempty" name:"CodeContent"`

	// Specifies the configuration list of task extended attributes.
	TaskExtConfigurationList []*TaskExtParameter `json:"TaskExtConfigurationList,omitnil,omitempty" name:"TaskExtConfigurationList"`

	// Cluster ID
	DataCluster *string `json:"DataCluster,omitnil,omitempty" name:"DataCluster"`

	// Specifies the running node.
	BrokerIp *string `json:"BrokerIp,omitnil,omitempty" name:"BrokerIp"`

	// Resource Pool Queue Name: Must be obtained using DescribeProjectClusterQueues API.
	YarnQueue *string `json:"YarnQueue,omitnil,omitempty" name:"YarnQueue"`

	// Source Data Source ID - One or more IDs separated by semicolons (;). Retrieve IDs using the DescribeDataSourceWithoutInfo API.
	SourceServiceId *string `json:"SourceServiceId,omitnil,omitempty" name:"SourceServiceId"`

	// Target data source ID- One or more IDs separated by semicolons (;). Retrieve IDs using the DescribeDataSourceWithoutInfo API.
	TargetServiceId *string `json:"TargetServiceId,omitnil,omitempty" name:"TargetServiceId"`

	// Scheduling parameter.
	TaskSchedulingParameterList []*TaskSchedulingParameter `json:"TaskSchedulingParameterList,omitnil,omitempty" name:"TaskSchedulingParameterList"`

	// ID used by the Bundle.
	BundleId *string `json:"BundleId,omitnil,omitempty" name:"BundleId"`

	// Bundle info.
	BundleInfo *string `json:"BundleInfo,omitnil,omitempty" name:"BundleInfo"`
}

// Predefined struct for user
type CreateTaskRequestParams struct {
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// The basic attributes of the task.
	TaskBaseAttribute *CreateTaskBaseAttribute `json:"TaskBaseAttribute,omitnil,omitempty" name:"TaskBaseAttribute"`

	// Task configurations.
	TaskConfiguration *CreateTaskConfiguration `json:"TaskConfiguration,omitnil,omitempty" name:"TaskConfiguration"`

	// Task scheduling configuration.
	TaskSchedulerConfiguration *CreateTaskSchedulerConfiguration `json:"TaskSchedulerConfiguration,omitnil,omitempty" name:"TaskSchedulerConfiguration"`
}

type CreateTaskRequest struct {
	*tchttp.BaseRequest
	
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// The basic attributes of the task.
	TaskBaseAttribute *CreateTaskBaseAttribute `json:"TaskBaseAttribute,omitnil,omitempty" name:"TaskBaseAttribute"`

	// Task configurations.
	TaskConfiguration *CreateTaskConfiguration `json:"TaskConfiguration,omitnil,omitempty" name:"TaskConfiguration"`

	// Task scheduling configuration.
	TaskSchedulerConfiguration *CreateTaskSchedulerConfiguration `json:"TaskSchedulerConfiguration,omitnil,omitempty" name:"TaskSchedulerConfiguration"`
}

func (r *CreateTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "TaskBaseAttribute")
	delete(f, "TaskConfiguration")
	delete(f, "TaskSchedulerConfiguration")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTaskResponseParams struct {
	// Task ID
	Data *CreateTaskResult `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateTaskResponseParams `json:"Response"`
}

func (r *CreateTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateTaskResult struct {
	// Task ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type CreateTaskSchedulerConfiguration struct {
	// Period type: defaults to DAY_CYCLE.
	// 
	// Supported types. 
	// 
	// ONEOFF_CYCLE: specifies a one-time cycle.
	// YEAR_CYCLE: specifies the year cycle.
	// MONTH_CYCLE: specifies the monthly cycle.
	// WEEK_CYCLE: specifies the week cycle.
	// DAY_CYCLE: specifies the day cycle.
	// HOUR_CYCLE: specifies the hour cycle.
	// MINUTE_CYCLE: specifies the minute cycle.
	// CRONTAB_CYCLE: specifies the crontab expression type.
	CycleType *string `json:"CycleType,omitnil,omitempty" name:"CycleType"`

	// Time zone, defaults to UTC+8.
	ScheduleTimeZone *string `json:"ScheduleTimeZone,omitnil,omitempty" name:"ScheduleTimeZone"`

	// Cron expression, defaults to 0 0 0 * * ? *.
	CrontabExpression *string `json:"CrontabExpression,omitnil,omitempty" name:"CrontabExpression"`

	// Effective date, defaults to 00:00:00 of the current date.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End date, defaults to 2099-12-31 23:59:59.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Execution time: the left-closed interval. Default: 00:00.
	ExecutionStartTime *string `json:"ExecutionStartTime,omitnil,omitempty" name:"ExecutionStartTime"`

	// Execution time: the right closed interval. Default: 23:59.
	ExecutionEndTime *string `json:"ExecutionEndTime,omitnil,omitempty" name:"ExecutionEndTime"`

	// Scheduling type: 0 for normal scheduling, 1 for dry-run scheduling. Default is 0.
	ScheduleRunType *string `json:"ScheduleRunType,omitnil,omitempty" name:"ScheduleRunType"`

	// Calendar scheduling value: 0 or 1, where 1 means ON and 0 means OFF. Default is 0.
	CalendarOpen *string `json:"CalendarOpen,omitnil,omitempty" name:"CalendarOpen"`

	// Calendar scheduling:  the calendar ID.
	CalendarId *string `json:"CalendarId,omitnil,omitempty" name:"CalendarId"`

	// Self-Dependent. Valid values: parallel, serial, orderly. Default value: serial. 
	SelfDepend *string `json:"SelfDepend,omitnil,omitempty" name:"SelfDepend"`

	// Specifies the upstream dependency list.
	UpstreamDependencyConfigList []*DependencyTaskBrief `json:"UpstreamDependencyConfigList,omitnil,omitempty" name:"UpstreamDependencyConfigList"`

	// List of Events
	EventListenerList []*EventListener `json:"EventListenerList,omitnil,omitempty" name:"EventListenerList"`

	// Task scheduling priority. Valid values: 4 (high), 5 (medium), 6 (low). Default: 6.
	RunPriority *string `json:"RunPriority,omitnil,omitempty" name:"RunPriority"`

	// Retry Policy: Retry Wait Time (in minutes): Default 5
	RetryWait *string `json:"RetryWait,omitnil,omitempty" name:"RetryWait"`

	// Retry Policy: maximum attempts. Default: 4.
	MaxRetryAttempts *string `json:"MaxRetryAttempts,omitnil,omitempty" name:"MaxRetryAttempts"`

	// Timeout Handling Policy: Execution Timeout (in minutes), default: -1
	ExecutionTTL *string `json:"ExecutionTTL,omitnil,omitempty" name:"ExecutionTTL"`

	// Timeout Handling Policy: Wait Duration Timeout  (in minutes), default: -1
	WaitExecutionTotalTTL *string `json:"WaitExecutionTotalTTL,omitnil,omitempty" name:"WaitExecutionTotalTTL"`

	// Rerun & Refill Configuration: Default: ALL;
	// 
	// * ALL: Rerun or refill is allowed regardless of whether the task succeeds or fails.
	// 
	// * FAILURE: Rerun or refill is allowed only if the task fails; not allowed if the task succeeds.
	// 
	// * NONE: Rerun or refill is not allowed regardless of success or failure.
	AllowRedoType *string `json:"AllowRedoType,omitnil,omitempty" name:"AllowRedoType"`

	// Output parameter list.
	ParamTaskOutList []*OutTaskParameter `json:"ParamTaskOutList,omitnil,omitempty" name:"ParamTaskOutList"`

	// Input parameter list.
	ParamTaskInList []*InTaskParameter `json:"ParamTaskInList,omitnil,omitempty" name:"ParamTaskInList"`

	// Output registration.
	TaskOutputRegistryList []*TaskDataRegistry `json:"TaskOutputRegistryList,omitnil,omitempty" name:"TaskOutputRegistryList"`

	// **Instance generation policy**.
	// T_PLUS_0: specifies t+0 generation. default policy.
	// T_PLUS_1: specifies t+1 generation.
	InitStrategy *string `json:"InitStrategy,omitnil,omitempty" name:"InitStrategy"`
}

// Predefined struct for user
type CreateWorkflowFolderRequestParams struct {
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Parent folder absolute path, such as /abc/de. if it is the root directory, pass /.
	ParentFolderPath *string `json:"ParentFolderPath,omitnil,omitempty" name:"ParentFolderPath"`

	// Specifies the folder name to be created.
	FolderName *string `json:"FolderName,omitnil,omitempty" name:"FolderName"`
}

type CreateWorkflowFolderRequest struct {
	*tchttp.BaseRequest
	
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Parent folder absolute path, such as /abc/de. if it is the root directory, pass /.
	ParentFolderPath *string `json:"ParentFolderPath,omitnil,omitempty" name:"ParentFolderPath"`

	// Specifies the folder name to be created.
	FolderName *string `json:"FolderName,omitnil,omitempty" name:"FolderName"`
}

func (r *CreateWorkflowFolderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWorkflowFolderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "ParentFolderPath")
	delete(f, "FolderName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateWorkflowFolderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWorkflowFolderResponseParams struct {
	// Folder creation result. error will be reported if creation failed.
	Data *CreateFolderResult `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateWorkflowFolderResponse struct {
	*tchttp.BaseResponse
	Response *CreateWorkflowFolderResponseParams `json:"Response"`
}

func (r *CreateWorkflowFolderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWorkflowFolderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWorkflowRequestParams struct {
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Workflow name.
	WorkflowName *string `json:"WorkflowName,omitnil,omitempty" name:"WorkflowName"`

	// Specifies the folder path.
	ParentFolderPath *string `json:"ParentFolderPath,omitnil,omitempty" name:"ParentFolderPath"`

	// Workflow type. value examples: cycle for periodic workflow; manual for manual workflow. default input: cycle.
	WorkflowType *string `json:"WorkflowType,omitnil,omitempty" name:"WorkflowType"`

	// Workflow description.
	WorkflowDesc *string `json:"WorkflowDesc,omitnil,omitempty" name:"WorkflowDesc"`

	// Workflow owner ID.
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// Workflow parameter.
	WorkflowParams []*ParamInfo `json:"WorkflowParams,omitnil,omitempty" name:"WorkflowParams"`

	// Specifies unified scheduling info.
	WorkflowSchedulerConfiguration *WorkflowSchedulerConfigurationInfo `json:"WorkflowSchedulerConfiguration,omitnil,omitempty" name:"WorkflowSchedulerConfiguration"`

	// BundleId item.
	BundleId *string `json:"BundleId,omitnil,omitempty" name:"BundleId"`

	// Bundle info.
	BundleInfo *string `json:"BundleInfo,omitnil,omitempty" name:"BundleInfo"`
}

type CreateWorkflowRequest struct {
	*tchttp.BaseRequest
	
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Workflow name.
	WorkflowName *string `json:"WorkflowName,omitnil,omitempty" name:"WorkflowName"`

	// Specifies the folder path.
	ParentFolderPath *string `json:"ParentFolderPath,omitnil,omitempty" name:"ParentFolderPath"`

	// Workflow type. value examples: cycle for periodic workflow; manual for manual workflow. default input: cycle.
	WorkflowType *string `json:"WorkflowType,omitnil,omitempty" name:"WorkflowType"`

	// Workflow description.
	WorkflowDesc *string `json:"WorkflowDesc,omitnil,omitempty" name:"WorkflowDesc"`

	// Workflow owner ID.
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// Workflow parameter.
	WorkflowParams []*ParamInfo `json:"WorkflowParams,omitnil,omitempty" name:"WorkflowParams"`

	// Specifies unified scheduling info.
	WorkflowSchedulerConfiguration *WorkflowSchedulerConfigurationInfo `json:"WorkflowSchedulerConfiguration,omitnil,omitempty" name:"WorkflowSchedulerConfiguration"`

	// BundleId item.
	BundleId *string `json:"BundleId,omitnil,omitempty" name:"BundleId"`

	// Bundle info.
	BundleInfo *string `json:"BundleInfo,omitnil,omitempty" name:"BundleInfo"`
}

func (r *CreateWorkflowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWorkflowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "WorkflowName")
	delete(f, "ParentFolderPath")
	delete(f, "WorkflowType")
	delete(f, "WorkflowDesc")
	delete(f, "OwnerUin")
	delete(f, "WorkflowParams")
	delete(f, "WorkflowSchedulerConfiguration")
	delete(f, "BundleId")
	delete(f, "BundleInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateWorkflowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWorkflowResponseParams struct {
	// Returns the workflow ID.
	Data *CreateWorkflowResult `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateWorkflowResponse struct {
	*tchttp.BaseResponse
	Response *CreateWorkflowResponseParams `json:"Response"`
}

func (r *CreateWorkflowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWorkflowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateWorkflowResult struct {
	// Workflow id after successful creation.
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`
}

type DataBackfillRange struct {
	// Start date in yyyy-MM-dd format. indicates the start from 00:00:00 on the specified date.
	StartDate *string `json:"StartDate,omitnil,omitempty" name:"StartDate"`

	// End date in the format yyyy-MM-dd, indicates ending at 23:59:59 of the specified date.
	EndDate *string `json:"EndDate,omitnil,omitempty" name:"EndDate"`

	// Start time of each day between [StartDate, EndDate] in HH:mm format. effective for tasks with a period of hours or less.
	ExecutionStartTime *string `json:"ExecutionStartTime,omitnil,omitempty" name:"ExecutionStartTime"`

	// End time point between [StartDate, EndDate] in HH:mm format. effective for tasks with a period of hours or less.
	ExecutionEndTime *string `json:"ExecutionEndTime,omitnil,omitempty" name:"ExecutionEndTime"`
}

type DeleteAlarmRuleResult struct {
	// Whether deletion succeeded.
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`
}

// Predefined struct for user
type DeleteCodeFileRequestParams struct {
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Code file ID. the parameter value comes from the CreateCodeFile API.
	CodeFileId *string `json:"CodeFileId,omitnil,omitempty" name:"CodeFileId"`
}

type DeleteCodeFileRequest struct {
	*tchttp.BaseRequest
	
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Code file ID. the parameter value comes from the CreateCodeFile API.
	CodeFileId *string `json:"CodeFileId,omitnil,omitempty" name:"CodeFileId"`
}

func (r *DeleteCodeFileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCodeFileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "CodeFileId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCodeFileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCodeFileResponseParams struct {
	// Execution result
	Data *CodeStudioFileActionResult `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteCodeFileResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCodeFileResponseParams `json:"Response"`
}

func (r *DeleteCodeFileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCodeFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCodeFolderRequestParams struct {
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Folder ID. the parameter value comes from the CreateCodeFolder API response.
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`
}

type DeleteCodeFolderRequest struct {
	*tchttp.BaseRequest
	
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Folder ID. the parameter value comes from the CreateCodeFolder API response.
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`
}

func (r *DeleteCodeFolderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCodeFolderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "FolderId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCodeFolderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCodeFolderResponseParams struct {
	// Execution result
	Data *CodeStudioFolderActionResult `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteCodeFolderResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCodeFolderResponseParams `json:"Response"`
}

func (r *DeleteCodeFolderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCodeFolderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteFolderResult struct {
	// Deletion status. true indicates success, false indicates failure.
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`
}

// Predefined struct for user
type DeleteOpsAlarmRuleRequestParams struct {
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Alarm Rule Unique ID. Obtained from the response of the CreateAlarmRule API. Either this field or AlarmRuleName must be provided.
	AlarmRuleId *string `json:"AlarmRuleId,omitnil,omitempty" name:"AlarmRuleId"`
}

type DeleteOpsAlarmRuleRequest struct {
	*tchttp.BaseRequest
	
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Alarm Rule Unique ID. Obtained from the response of the CreateAlarmRule API. Either this field or AlarmRuleName must be provided.
	AlarmRuleId *string `json:"AlarmRuleId,omitnil,omitempty" name:"AlarmRuleId"`
}

func (r *DeleteOpsAlarmRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteOpsAlarmRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "AlarmRuleId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteOpsAlarmRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteOpsAlarmRuleResponseParams struct {
	// Whether deletion succeeded.
	Data *DeleteAlarmRuleResult `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteOpsAlarmRuleResponse struct {
	*tchttp.BaseResponse
	Response *DeleteOpsAlarmRuleResponseParams `json:"Response"`
}

func (r *DeleteOpsAlarmRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteOpsAlarmRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteResourceFileRequestParams struct {
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Resource ID. obtain through the API ListResourceFiles.
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`
}

type DeleteResourceFileRequest struct {
	*tchttp.BaseRequest
	
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Resource ID. obtain through the API ListResourceFiles.
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`
}

func (r *DeleteResourceFileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteResourceFileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "ResourceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteResourceFileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteResourceFileResponseParams struct {
	// Specifies the resource deletion result.
	Data *DeleteResourceFileResult `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteResourceFileResponse struct {
	*tchttp.BaseResponse
	Response *DeleteResourceFileResponseParams `json:"Response"`
}

func (r *DeleteResourceFileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteResourceFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteResourceFileResult struct {
	// true
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`
}

// Predefined struct for user
type DeleteResourceFolderRequestParams struct {
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Folder ID. obtain through the ListResourceFolders API.
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`
}

type DeleteResourceFolderRequest struct {
	*tchttp.BaseRequest
	
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Folder ID. obtain through the ListResourceFolders API.
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`
}

func (r *DeleteResourceFolderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteResourceFolderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "FolderId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteResourceFolderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteResourceFolderResponseParams struct {
	// true represents successful deletion. false represents failure.
	Data *DeleteFolderResult `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteResourceFolderResponse struct {
	*tchttp.BaseResponse
	Response *DeleteResourceFolderResponseParams `json:"Response"`
}

func (r *DeleteResourceFolderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteResourceFolderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSQLFolderRequestParams struct {
	// Folder ID
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`

	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type DeleteSQLFolderRequest struct {
	*tchttp.BaseRequest
	
	// Folder ID
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`

	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

func (r *DeleteSQLFolderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSQLFolderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FolderId")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSQLFolderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSQLFolderResponseParams struct {
	// Operation result.
	Data *SQLContentActionResult `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteSQLFolderResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSQLFolderResponseParams `json:"Response"`
}

func (r *DeleteSQLFolderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSQLFolderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSQLScriptRequestParams struct {
	// Script Id.
	ScriptId *string `json:"ScriptId,omitnil,omitempty" name:"ScriptId"`

	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type DeleteSQLScriptRequest struct {
	*tchttp.BaseRequest
	
	// Script Id.
	ScriptId *string `json:"ScriptId,omitnil,omitempty" name:"ScriptId"`

	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

func (r *DeleteSQLScriptRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSQLScriptRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ScriptId")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSQLScriptRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSQLScriptResponseParams struct {
	// Execution result
	Data *SQLContentActionResult `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteSQLScriptResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSQLScriptResponseParams `json:"Response"`
}

func (r *DeleteSQLScriptResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSQLScriptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTaskRequestParams struct {
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Either Task ID or VirtualTaskId must be provided (optional, choose one).
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Whether to send a notification to downstream task owners when performing task operations.
	// true: Send notification
	// false: Do not send notification
	// default: false.
	OperateInform *bool `json:"OperateInform,omitnil,omitempty" name:"OperateInform"`

	// Task Deletion Mode.
	// true: Do not force downstream task instances to fail
	// false: Force downstream task instances to fail
	// default: false 
	DeleteMode *bool `json:"DeleteMode,omitnil,omitempty" name:"DeleteMode"`
}

type DeleteTaskRequest struct {
	*tchttp.BaseRequest
	
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Either Task ID or VirtualTaskId must be provided (optional, choose one).
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Whether to send a notification to downstream task owners when performing task operations.
	// true: Send notification
	// false: Do not send notification
	// default: false.
	OperateInform *bool `json:"OperateInform,omitnil,omitempty" name:"OperateInform"`

	// Task Deletion Mode.
	// true: Do not force downstream task instances to fail
	// false: Force downstream task instances to fail
	// default: false 
	DeleteMode *bool `json:"DeleteMode,omitnil,omitempty" name:"DeleteMode"`
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
	delete(f, "ProjectId")
	delete(f, "TaskId")
	delete(f, "OperateInform")
	delete(f, "DeleteMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTaskResponseParams struct {
	// Whether deletion succeeded.
	Data *DeleteTaskResult `json:"Data,omitnil,omitempty" name:"Data"`

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

type DeleteTaskResult struct {
	// Deletion status. true indicates success. false indicates failure.
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`
}

// Predefined struct for user
type DeleteWorkflowFolderRequestParams struct {
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Folder ID. Obtain through the ListWorkflowFolders API.
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`
}

type DeleteWorkflowFolderRequest struct {
	*tchttp.BaseRequest
	
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Folder ID. Obtain through the ListWorkflowFolders API.
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`
}

func (r *DeleteWorkflowFolderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteWorkflowFolderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "FolderId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteWorkflowFolderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteWorkflowFolderResponseParams struct {
	// Deletion result
	Data *DeleteFolderResult `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteWorkflowFolderResponse struct {
	*tchttp.BaseResponse
	Response *DeleteWorkflowFolderResponseParams `json:"Response"`
}

func (r *DeleteWorkflowFolderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteWorkflowFolderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteWorkflowRequestParams struct {
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Workflow id.
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`
}

type DeleteWorkflowRequest struct {
	*tchttp.BaseRequest
	
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Workflow id.
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`
}

func (r *DeleteWorkflowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteWorkflowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "WorkflowId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteWorkflowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteWorkflowResponseParams struct {
	// Returns the count of successfully deleted workflow tasks, number of failures, and total number of tasks.
	Data *DeleteWorkflowResult `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteWorkflowResponse struct {
	*tchttp.BaseResponse
	Response *DeleteWorkflowResponseParams `json:"Response"`
}

func (r *DeleteWorkflowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteWorkflowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteWorkflowResult struct {
	// Indicates whether the workflow deletion was successful
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`
}

type DependencyConfigPage struct {
	// Total number of records matching the query conditions.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Total number of pages matching the query conditions.
	TotalPageNumber *uint64 `json:"TotalPageNumber,omitnil,omitempty" name:"TotalPageNumber"`

	// The page number of the current request.
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// The number of entries in the current request data page.
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Paginated Data
	Items []*TaskDependDto `json:"Items,omitnil,omitempty" name:"Items"`
}

type DependencyStrategyTask struct {
	// Wait upstream task instance policy: EXECUTING (execute); WAITING (wait).
	PollingNullStrategy *string `json:"PollingNullStrategy,omitnil,omitempty" name:"PollingNullStrategy"`

	// This field is required only when PollingNullStrategy is set to EXECUTING.
	// Type: List
	// 
	// NOT_EXIST (default) - In cases where minute depends on minute / hour depends on hour, the parent instance does not fall within the scheduling time range of the downstream instance.
	// 
	// PARENT_EXPIRED - The parent instance failed.
	// 
	// PARENT_TIMEOUT - The parent instance timed out.
	// 
	// If any of the above conditions are met, the dependency check for that parent task instance is considered satisfied. In all other cases, the system must wait for the parent instance.
	TaskDependencyExecutingStrategies []*string `json:"TaskDependencyExecutingStrategies,omitnil,omitempty" name:"TaskDependencyExecutingStrategies"`

	// This field is required only when TaskDependencyExecutingStrategies includes PARENT_TIMEOUT.
	// Specifies the timeout duration (in minutes) for the downstream task's dependency on the parent instance execution.
	TaskDependencyExecutingTimeoutValue *int64 `json:"TaskDependencyExecutingTimeoutValue,omitnil,omitempty" name:"TaskDependencyExecutingTimeoutValue"`
}

type DependencyTaskBrief struct {
	// Task ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Main dependency configuration. Valid values:
	// 
	// * CRONTAB
	// * DAY
	// * HOUR
	// * LIST_DAY
	// * LIST_HOUR
	// * LIST_MINUTE
	// * MINUTE
	// * MONTH
	// * RANGE_DAY
	// * RANGE_HOUR
	// * RANGE_MINUTE
	// * WEEK
	// * YEAR
	MainCyclicConfig *string `json:"MainCyclicConfig,omitnil,omitempty" name:"MainCyclicConfig"`

	// Configures secondary dependencies.  Valid values:
	// * ALL_DAY_OF_YEAR
	// * ALL_MONTH_OF_YEAR
	// * CURRENT
	// * CURRENT_DAY
	// * CURRENT_HOUR
	// * CURRENT_MINUTE
	// * CURRENT_MONTH
	// * CURRENT_WEEK
	// * CURRENT_YEAR
	// * PREVIOUS_BEGIN_OF_MONTH
	// * PREVIOUS_DAY
	// * PREVIOUS_DAY_LATER_OFFSET_HOUR
	// * PREVIOUS_DAY_LATER_OFFSET_MINUTE
	// * PREVIOUS_END_OF_MONTH
	// * PREVIOUS_FRIDAY
	// * PREVIOUS_HOUR
	// * PREVIOUS_HOUR_CYCLE
	// * PREVIOUS_HOUR_LATER_OFFSET_MINUTE
	// * PREVIOUS_MINUTE_CYCLE
	// * PREVIOUS_MONTH
	// * PREVIOUS_WEEK
	// * PREVIOUS_WEEKEND
	// * RECENT_DATE
	SubordinateCyclicConfig *string `json:"SubordinateCyclicConfig,omitnil,omitempty" name:"SubordinateCyclicConfig"`

	// Offset in Range/List Mode
	Offset *string `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Dependency Execution Policy
	DependencyStrategy *DependencyStrategyTask `json:"DependencyStrategy,omitnil,omitempty" name:"DependencyStrategy"`
}

type EventListener struct {
	// Event name
	EventName *string `json:"EventName,omitnil,omitempty" name:"EventName"`

	// Event cycle. valid values: SECOND, MIN, HOUR, DAY.
	EventSubType *string `json:"EventSubType,omitnil,omitempty" name:"EventSubType"`

	// Event BROADCAST type: SINGLE, BROADCAST.
	EventBroadcastType *string `json:"EventBroadcastType,omitnil,omitempty" name:"EventBroadcastType"`

	// Extension Information
	// 
	PropertiesList []*ParamInfo `json:"PropertiesList,omitnil,omitempty" name:"PropertiesList"`
}

// Predefined struct for user
type GetAlarmMessageRequestParams struct {
	// Project ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Alarm message Id.
	AlarmMessageId *string `json:"AlarmMessageId,omitnil,omitempty" name:"AlarmMessageId"`

	// Specifies the time zone of the return date. default UTC+8.
	TimeZone *string `json:"TimeZone,omitnil,omitempty" name:"TimeZone"`
}

type GetAlarmMessageRequest struct {
	*tchttp.BaseRequest
	
	// Project ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Alarm message Id.
	AlarmMessageId *string `json:"AlarmMessageId,omitnil,omitempty" name:"AlarmMessageId"`

	// Specifies the time zone of the return date. default UTC+8.
	TimeZone *string `json:"TimeZone,omitnil,omitempty" name:"TimeZone"`
}

func (r *GetAlarmMessageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetAlarmMessageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "AlarmMessageId")
	delete(f, "TimeZone")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetAlarmMessageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetAlarmMessageResponseParams struct {
	// Alarm information.
	Data *AlarmMessage `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetAlarmMessageResponse struct {
	*tchttp.BaseResponse
	Response *GetAlarmMessageResponseParams `json:"Response"`
}

func (r *GetAlarmMessageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetAlarmMessageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetCodeFileRequestParams struct {
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Code file ID. the parameter value comes from the CreateCodeFile API response.
	CodeFileId *string `json:"CodeFileId,omitnil,omitempty" name:"CodeFileId"`

	// true: return file content and configuration. false: only return configuration message. default false.
	IncludeContent *bool `json:"IncludeContent,omitnil,omitempty" name:"IncludeContent"`
}

type GetCodeFileRequest struct {
	*tchttp.BaseRequest
	
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Code file ID. the parameter value comes from the CreateCodeFile API response.
	CodeFileId *string `json:"CodeFileId,omitnil,omitempty" name:"CodeFileId"`

	// true: return file content and configuration. false: only return configuration message. default false.
	IncludeContent *bool `json:"IncludeContent,omitnil,omitempty" name:"IncludeContent"`
}

func (r *GetCodeFileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetCodeFileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "CodeFileId")
	delete(f, "IncludeContent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetCodeFileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetCodeFileResponseParams struct {
	// Specifies the code file detail.
	Data *CodeFile `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetCodeFileResponse struct {
	*tchttp.BaseResponse
	Response *GetCodeFileResponseParams `json:"Response"`
}

func (r *GetCodeFileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetCodeFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetOpsAlarmRuleRequestParams struct {
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Specifies the unique id of the Alarm rule.
	AlarmRuleId *string `json:"AlarmRuleId,omitnil,omitempty" name:"AlarmRuleId"`
}

type GetOpsAlarmRuleRequest struct {
	*tchttp.BaseRequest
	
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Specifies the unique id of the Alarm rule.
	AlarmRuleId *string `json:"AlarmRuleId,omitnil,omitempty" name:"AlarmRuleId"`
}

func (r *GetOpsAlarmRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetOpsAlarmRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "AlarmRuleId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetOpsAlarmRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetOpsAlarmRuleResponseParams struct {
	// Specifies the Alarm rule details.
	Data *AlarmRuleData `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetOpsAlarmRuleResponse struct {
	*tchttp.BaseResponse
	Response *GetOpsAlarmRuleResponseParams `json:"Response"`
}

func (r *GetOpsAlarmRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetOpsAlarmRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetOpsAsyncJobRequestParams struct {
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Asynchronous operation id.
	AsyncId *string `json:"AsyncId,omitnil,omitempty" name:"AsyncId"`
}

type GetOpsAsyncJobRequest struct {
	*tchttp.BaseRequest
	
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Asynchronous operation id.
	AsyncId *string `json:"AsyncId,omitnil,omitempty" name:"AsyncId"`
}

func (r *GetOpsAsyncJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetOpsAsyncJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "AsyncId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetOpsAsyncJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetOpsAsyncJobResponseParams struct {
	// Asynchronous operation result.
	Data *OpsAsyncJobDetail `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetOpsAsyncJobResponse struct {
	*tchttp.BaseResponse
	Response *GetOpsAsyncJobResponseParams `json:"Response"`
}

func (r *GetOpsAsyncJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetOpsAsyncJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetOpsTaskCodeRequestParams struct {
	// Project id.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Task ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type GetOpsTaskCodeRequest struct {
	*tchttp.BaseRequest
	
	// Project id.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Task ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *GetOpsTaskCodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetOpsTaskCodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetOpsTaskCodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetOpsTaskCodeResponseParams struct {
	// Retrieves the task code result.
	Data *TaskCode `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetOpsTaskCodeResponse struct {
	*tchttp.BaseResponse
	Response *GetOpsTaskCodeResponseParams `json:"Response"`
}

func (r *GetOpsTaskCodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetOpsTaskCodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetOpsTaskRequestParams struct {
	// Task ID		
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type GetOpsTaskRequest struct {
	*tchttp.BaseRequest
	
	// Task ID		
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

func (r *GetOpsTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetOpsTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetOpsTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetOpsTaskResponseParams struct {
	// Task details.
	Data *Task `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetOpsTaskResponse struct {
	*tchttp.BaseResponse
	Response *GetOpsTaskResponseParams `json:"Response"`
}

func (r *GetOpsTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetOpsTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetOpsWorkflowRequestParams struct {
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Workflow Id, can be obtained through the ListOpsWorkflows api.
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`
}

type GetOpsWorkflowRequest struct {
	*tchttp.BaseRequest
	
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Workflow Id, can be obtained through the ListOpsWorkflows api.
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`
}

func (r *GetOpsWorkflowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetOpsWorkflowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "WorkflowId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetOpsWorkflowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetOpsWorkflowResponseParams struct {
	// Workflow scheduling details.
	Data *OpsWorkflowDetail `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetOpsWorkflowResponse struct {
	*tchttp.BaseResponse
	Response *GetOpsWorkflowResponseParams `json:"Response"`
}

func (r *GetOpsWorkflowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetOpsWorkflowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetResourceFileRequestParams struct {
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Resource file ID. obtain through the API ListResourceFiles.
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`
}

type GetResourceFileRequest struct {
	*tchttp.BaseRequest
	
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Resource file ID. obtain through the API ListResourceFiles.
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`
}

func (r *GetResourceFileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetResourceFileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "ResourceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetResourceFileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetResourceFileResponseParams struct {
	// Describes the resource file details.
	Data *ResourceFile `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetResourceFileResponse struct {
	*tchttp.BaseResponse
	Response *GetResourceFileResponseParams `json:"Response"`
}

func (r *GetResourceFileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetResourceFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetSQLScriptRequestParams struct {
	// Script Id.
	ScriptId *string `json:"ScriptId,omitnil,omitempty" name:"ScriptId"`

	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type GetSQLScriptRequest struct {
	*tchttp.BaseRequest
	
	// Script Id.
	ScriptId *string `json:"ScriptId,omitnil,omitempty" name:"ScriptId"`

	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

func (r *GetSQLScriptRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetSQLScriptRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ScriptId")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetSQLScriptRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetSQLScriptResponseParams struct {
	// Script details.
	Data *SQLScript `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetSQLScriptResponse struct {
	*tchttp.BaseResponse
	Response *GetSQLScriptResponseParams `json:"Response"`
}

func (r *GetSQLScriptResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetSQLScriptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTaskCodeRequestParams struct {
	// The project id.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Task ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type GetTaskCodeRequest struct {
	*tchttp.BaseRequest
	
	// The project id.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Task ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *GetTaskCodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTaskCodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetTaskCodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTaskCodeResponseParams struct {
	// Retrieves the task code result.
	Data *TaskCodeResult `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetTaskCodeResponse struct {
	*tchttp.BaseResponse
	Response *GetTaskCodeResponseParams `json:"Response"`
}

func (r *GetTaskCodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTaskCodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTaskInstanceLogRequestParams struct {
	// **Project ID**. specifies the project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// **Instance unique id**.
	InstanceKey *string `json:"InstanceKey,omitnil,omitempty" name:"InstanceKey"`

	// **Instance Lifecycle Number** - Identifies a specific execution of an instance.
	// For example: the first run of a cyclic instance has a lifecycle number of 0. If the user reruns that instance later, the second execution will have a lifecycle number of 1;
	// By default, the latest execution is used.
	LifeRoundNum *uint64 `json:"LifeRoundNum,omitnil,omitempty" name:"LifeRoundNum"`

	// **Log level** default All - Info - Debug - Warn - Error - All.
	LogLevel *string `json:"LogLevel,omitnil,omitempty" name:"LogLevel"`

	// **Used when performing paginated log queries; has no specific business meaning.**
	// 
	// For the first query, the value is null.
	// 
	// For subsequent queries, use the value of the NextCursor field returned from the previous query.
	NextCursor *string `json:"NextCursor,omitnil,omitempty" name:"NextCursor"`
}

type GetTaskInstanceLogRequest struct {
	*tchttp.BaseRequest
	
	// **Project ID**. specifies the project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// **Instance unique id**.
	InstanceKey *string `json:"InstanceKey,omitnil,omitempty" name:"InstanceKey"`

	// **Instance Lifecycle Number** - Identifies a specific execution of an instance.
	// For example: the first run of a cyclic instance has a lifecycle number of 0. If the user reruns that instance later, the second execution will have a lifecycle number of 1;
	// By default, the latest execution is used.
	LifeRoundNum *uint64 `json:"LifeRoundNum,omitnil,omitempty" name:"LifeRoundNum"`

	// **Log level** default All - Info - Debug - Warn - Error - All.
	LogLevel *string `json:"LogLevel,omitnil,omitempty" name:"LogLevel"`

	// **Used when performing paginated log queries; has no specific business meaning.**
	// 
	// For the first query, the value is null.
	// 
	// For subsequent queries, use the value of the NextCursor field returned from the previous query.
	NextCursor *string `json:"NextCursor,omitnil,omitempty" name:"NextCursor"`
}

func (r *GetTaskInstanceLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTaskInstanceLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "InstanceKey")
	delete(f, "LifeRoundNum")
	delete(f, "LogLevel")
	delete(f, "NextCursor")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetTaskInstanceLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTaskInstanceLogResponseParams struct {
	// Details of a scheduled instance.
	Data *InstanceLog `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetTaskInstanceLogResponse struct {
	*tchttp.BaseResponse
	Response *GetTaskInstanceLogResponseParams `json:"Response"`
}

func (r *GetTaskInstanceLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTaskInstanceLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTaskInstanceRequestParams struct {
	// Project id.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Instance unique id, can be obtained through ListInstances.
	InstanceKey *string `json:"InstanceKey,omitnil,omitempty" name:"InstanceKey"`

	// **Time zone** timeZone, specifies the time zone of the passed in time string. default UTC+8.
	TimeZone *string `json:"TimeZone,omitnil,omitempty" name:"TimeZone"`
}

type GetTaskInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Project id.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Instance unique id, can be obtained through ListInstances.
	InstanceKey *string `json:"InstanceKey,omitnil,omitempty" name:"InstanceKey"`

	// **Time zone** timeZone, specifies the time zone of the passed in time string. default UTC+8.
	TimeZone *string `json:"TimeZone,omitnil,omitempty" name:"TimeZone"`
}

func (r *GetTaskInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTaskInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "InstanceKey")
	delete(f, "TimeZone")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetTaskInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTaskInstanceResponseParams struct {
	// Details of an instance.
	Data *TaskInstanceDetail `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetTaskInstanceResponse struct {
	*tchttp.BaseResponse
	Response *GetTaskInstanceResponseParams `json:"Response"`
}

func (r *GetTaskInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTaskInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTaskRequestParams struct {
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Task ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type GetTaskRequest struct {
	*tchttp.BaseRequest
	
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Task ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *GetTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTaskResponseParams struct {
	// Task details.
	Data *Task `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetTaskResponse struct {
	*tchttp.BaseResponse
	Response *GetTaskResponseParams `json:"Response"`
}

func (r *GetTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTaskVersionRequestParams struct {
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Task ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Submit version ID. If not specified, the latest submit version will be used by default.
	VersionId *string `json:"VersionId,omitnil,omitempty" name:"VersionId"`
}

type GetTaskVersionRequest struct {
	*tchttp.BaseRequest
	
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Task ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Submit version ID. If not specified, the latest submit version will be used by default.
	VersionId *string `json:"VersionId,omitnil,omitempty" name:"VersionId"`
}

func (r *GetTaskVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTaskVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "TaskId")
	delete(f, "VersionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetTaskVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTaskVersionResponseParams struct {
	// Version detail.
	Data *TaskVersionDetail `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetTaskVersionResponse struct {
	*tchttp.BaseResponse
	Response *GetTaskVersionResponseParams `json:"Response"`
}

func (r *GetTaskVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTaskVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetWorkflowRequestParams struct {
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Workflow ID. Obtained through the  ListWorkflows API.
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`
}

type GetWorkflowRequest struct {
	*tchttp.BaseRequest
	
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Workflow ID. Obtained through the  ListWorkflows API.
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`
}

func (r *GetWorkflowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetWorkflowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "WorkflowId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetWorkflowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetWorkflowResponseParams struct {
	// Describes workflow details.
	Data *WorkflowDetail `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetWorkflowResponse struct {
	*tchttp.BaseResponse
	Response *GetWorkflowResponseParams `json:"Response"`
}

func (r *GetWorkflowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetWorkflowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InTaskParameter struct {
	// Parameter name.
	ParamKey *string `json:"ParamKey,omitnil,omitempty" name:"ParamKey"`

	// Parameter Description: The format is ProjectIdentifier.TaskName.ParameterName
	// Example: project_wedata_1.sh_250820_104107.pp_out
	ParamDesc *string `json:"ParamDesc,omitnil,omitempty" name:"ParamDesc"`

	// Parent Task ID
	FromTaskId *string `json:"FromTaskId,omitnil,omitempty" name:"FromTaskId"`

	// Parent task parameter key.
	FromParamKey *string `json:"FromParamKey,omitnil,omitempty" name:"FromParamKey"`
}

type InstanceExecution struct {
	// Instance unique identifier.
	InstanceKey *string `json:"InstanceKey,omitnil,omitempty" name:"InstanceKey"`

	// **Instance lifecycle number, identifies one-time execution of the instance.**.
	// 
	// For example, the first run of a periodic instance is numbered 0. if the user reruns the instance subsequently, the second execution is numbered 1.
	LifeRoundNum *uint64 `json:"LifeRoundNum,omitnil,omitempty" name:"LifeRoundNum"`

	// **Instance status**.
	// -WAIT_EVENT: specifies the wait for event.
	// -WAIT_UPSTREAM: waiting for upstream.
	// - WAIT_RUN: awaiting execution.
	// - RUNNING: indicates the instance is RUNNING.
	// - SKIP_RUNNING: SKIP RUNNING.
	// - FAILED_RETRY: RETRY on failure.
	// - EXPIRED: failed.
	// -COMPLETED: success.
	InstanceState *string `json:"InstanceState,omitnil,omitempty" name:"InstanceState"`

	// **Trigger type for instance running**.
	// 
	// -RERUN indicates a rerun.
	// -ADDITION indicates supplementary recording.
	// -PERIODIC indicates a period.
	// -APERIODIC indicates non-periodic.
	// -RERUN_SKIP_RUN indicates rerun - empty run.
	// -ADDITION_SKIP_RUN indicates a supplementary empty run.
	// -PERIODIC_SKIP_RUN indicates cycle - empty run.
	// -APERIODIC_SKIP_RUN indicates non-periodic empty run.
	// -MANUAL_TRIGGER indicates manual triggering.
	// -RERUN_MANUAL_TRIGGER indicates manual triggering - rerun.
	RunType *string `json:"RunType,omitnil,omitempty" name:"RunType"`

	// Specifies the number of retry attempts on failure.
	Tries *uint64 `json:"Tries,omitnil,omitempty" name:"Tries"`

	// **Specifies the lifecycle list for instance execution.**.
	ExecutionPhaseList []*InstanceExecutionPhase `json:"ExecutionPhaseList,omitnil,omitempty" name:"ExecutionPhaseList"`

	// Time spent, in milliseconds.
	CostTime *int64 `json:"CostTime,omitnil,omitempty" name:"CostTime"`
}

type InstanceExecutionPhase struct {
	// Start time of the state.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// **Instance lifecycle phase status**.
	// 
	// -WAIT_UPSTREAM indicates the wait event/upstream status.
	// -WAIT_RUN indicates the waiting for running status.
	// -RUNNING indicates the running state.
	// -COMPLETE indicates the final state of completion.
	// - FAILED indicates the final state - retry on failure.
	// -EXPIRED indicates the final state - failure.
	// -SKIP_RUNNING indicates the branch skipped by the upstream branch node in the final state.
	// -HISTORY indicates compatibility with historical instances before 2024-03-30. no need to pay attention to this enum afterward.
	DetailState *string `json:"DetailState,omitnil,omitempty" name:"DetailState"`

	// End time of the state.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type InstanceLog struct {
	// Instance unique id.
	InstanceKey *string `json:"InstanceKey,omitnil,omitempty" name:"InstanceKey"`

	// Project ID.
	// 
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Specifies the code content to run.
	CodeContent *string `json:"CodeContent,omitnil,omitempty" name:"CodeContent"`

	// log information
	LogInfo *string `json:"LogInfo,omitnil,omitempty" name:"LogInfo"`

	// Used for paginated log queries; has no specific business meaning.
	// 
	// For the first query, set the value to null.
	// 
	// For subsequent queries, use the NextCursor value returned from the previous query.
	NextCursor *string `json:"NextCursor,omitnil,omitempty" name:"NextCursor"`
}

type JobDto struct {
	// Task ID of the data exploration.
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// Task name.
	JobName *string `json:"JobName,omitnil,omitempty" name:"JobName"`

	// Task type
	JobType *string `json:"JobType,omitnil,omitempty" name:"JobType"`

	// Script ID
	ScriptId *string `json:"ScriptId,omitnil,omitempty" name:"ScriptId"`

	// Subtask List
	JobExecutionList []*JobExecutionDto `json:"JobExecutionList,omitnil,omitempty" name:"JobExecutionList"`

	// Specifies the script content.
	ScriptContent *string `json:"ScriptContent,omitnil,omitempty" name:"ScriptContent"`

	// State of a task.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Task creation time
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Update time.
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// End time.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Root account UIN.
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// Account UIN.
	UserUin *string `json:"UserUin,omitnil,omitempty" name:"UserUin"`

	// Duration. specifies the time taken.
	TimeCost *int64 `json:"TimeCost,omitnil,omitempty" name:"TimeCost"`

	// Specifies whether the script content is truncated.
	ScriptContentTruncate *bool `json:"ScriptContentTruncate,omitnil,omitempty" name:"ScriptContentTruncate"`
}

type JobExecutionDto struct {
	// Job ID of the data exploration.
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// Query job ID.
	JobExecutionId *string `json:"JobExecutionId,omitnil,omitempty" name:"JobExecutionId"`

	// Specifies the subquery name.
	JobExecutionName *string `json:"JobExecutionName,omitnil,omitempty" name:"JobExecutionName"`

	// Specifies the subquery sql content.
	ScriptContent *string `json:"ScriptContent,omitnil,omitempty" name:"ScriptContent"`

	// Subquery status.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Creation time.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Execution phase.
	ExecuteStageInfo *string `json:"ExecuteStageInfo,omitnil,omitempty" name:"ExecuteStageInfo"`

	// Log path
	LogFilePath *string `json:"LogFilePath,omitnil,omitempty" name:"LogFilePath"`

	// Result path for download.
	ResultFilePath *string `json:"ResultFilePath,omitnil,omitempty" name:"ResultFilePath"`

	// Preview result path.
	ResultPreviewFilePath *string `json:"ResultPreviewFilePath,omitnil,omitempty" name:"ResultPreviewFilePath"`

	// Total number of lines in the task execution result.
	ResultTotalCount *int64 `json:"ResultTotalCount,omitnil,omitempty" name:"ResultTotalCount"`

	// Update time.
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// End time.
	// 
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Duration. specifies the time taken.
	TimeCost *int64 `json:"TimeCost,omitnil,omitempty" name:"TimeCost"`

	// SQL content in the context.
	ContextScriptContent []*string `json:"ContextScriptContent,omitnil,omitempty" name:"ContextScriptContent"`

	// Specifies the preview row count for task execution results.
	ResultPreviewCount *int64 `json:"ResultPreviewCount,omitnil,omitempty" name:"ResultPreviewCount"`

	// Specifies the number of affected rows in task execution.
	ResultEffectCount *int64 `json:"ResultEffectCount,omitnil,omitempty" name:"ResultEffectCount"`

	// Whether the full result is being collected: default false. true indicates the full result is being collected, for the frontend to determine whether to continue to poll.
	CollectingTotalResult *bool `json:"CollectingTotalResult,omitnil,omitempty" name:"CollectingTotalResult"`

	// Specifies whether to truncate the script content.
	ScriptContentTruncate *bool `json:"ScriptContentTruncate,omitnil,omitempty" name:"ScriptContentTruncate"`
}

type KVMap struct {
	// k
	K *string `json:"K,omitnil,omitempty" name:"K"`

	// v
	V *string `json:"V,omitnil,omitempty" name:"V"`
}

type KVPair struct {
	// Key name
	// 
	K *string `json:"K,omitnil,omitempty" name:"K"`

	// The value. do not pass SQL (the request will be deemed as an attack on the api). if needed, transcode the SQL with Base64 and decode it.
	V *string `json:"V,omitnil,omitempty" name:"V"`
}

// Predefined struct for user
type KillTaskInstancesAsyncRequestParams struct {
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Instance id list, can be obtained from ListInstances API.
	InstanceKeyList []*string `json:"InstanceKeyList,omitnil,omitempty" name:"InstanceKeyList"`
}

type KillTaskInstancesAsyncRequest struct {
	*tchttp.BaseRequest
	
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Instance id list, can be obtained from ListInstances API.
	InstanceKeyList []*string `json:"InstanceKeyList,omitnil,omitempty" name:"InstanceKeyList"`
}

func (r *KillTaskInstancesAsyncRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *KillTaskInstancesAsyncRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "InstanceKeyList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "KillTaskInstancesAsyncRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type KillTaskInstancesAsyncResponseParams struct {
	// Async id of the batch termination operation. can be used in the GetAsyncJob API to query execution detail.
	Data *OpsAsyncResponse `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type KillTaskInstancesAsyncResponse struct {
	*tchttp.BaseResponse
	Response *KillTaskInstancesAsyncResponseParams `json:"Response"`
}

func (r *KillTaskInstancesAsyncResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *KillTaskInstancesAsyncResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListAlarmMessages struct {
	// Page number.
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Pagination size.
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Total number of entries
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Total pages
	TotalPageNumber *uint64 `json:"TotalPageNumber,omitnil,omitempty" name:"TotalPageNumber"`

	// Alarm information list.
	Items []*AlarmMessage `json:"Items,omitnil,omitempty" name:"Items"`
}

// Predefined struct for user
type ListAlarmMessagesRequestParams struct {
	// Project id.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Page number for pagination, minimum value is 1.
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Specifies the number of items displayed per page. maximum value: 100.
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Starting Alarm time. format: yyyy-MM-dd HH:MM:ss.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Specifies the Alarm end time in the format yyyy-MM-dd HH:MM:ss.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Alarm level.
	AlarmLevel *uint64 `json:"AlarmLevel,omitnil,omitempty" name:"AlarmLevel"`

	// Alert recipient Id.
	AlarmRecipientId *string `json:"AlarmRecipientId,omitnil,omitempty" name:"AlarmRecipientId"`

	// For incoming and returned filter time zone, default UTC+8.
	TimeZone *string `json:"TimeZone,omitnil,omitempty" name:"TimeZone"`
}

type ListAlarmMessagesRequest struct {
	*tchttp.BaseRequest
	
	// Project id.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Page number for pagination, minimum value is 1.
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Specifies the number of items displayed per page. maximum value: 100.
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Starting Alarm time. format: yyyy-MM-dd HH:MM:ss.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Specifies the Alarm end time in the format yyyy-MM-dd HH:MM:ss.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Alarm level.
	AlarmLevel *uint64 `json:"AlarmLevel,omitnil,omitempty" name:"AlarmLevel"`

	// Alert recipient Id.
	AlarmRecipientId *string `json:"AlarmRecipientId,omitnil,omitempty" name:"AlarmRecipientId"`

	// For incoming and returned filter time zone, default UTC+8.
	TimeZone *string `json:"TimeZone,omitnil,omitempty" name:"TimeZone"`
}

func (r *ListAlarmMessagesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAlarmMessagesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "AlarmLevel")
	delete(f, "AlarmRecipientId")
	delete(f, "TimeZone")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListAlarmMessagesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAlarmMessagesResponseParams struct {
	// Alarm information list.
	Data *ListAlarmMessages `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListAlarmMessagesResponse struct {
	*tchttp.BaseResponse
	Response *ListAlarmMessagesResponseParams `json:"Response"`
}

func (r *ListAlarmMessagesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAlarmMessagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListAlarmRulesResult struct {
	// Page number for pagination. current page number.
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Specifies the number of items to display per page.
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Total number of pages.
	TotalPageNumber *uint64 `json:"TotalPageNumber,omitnil,omitempty" name:"TotalPageNumber"`

	// Count of all Alarm rules.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Alert rule information list.
	Items []*AlarmRuleData `json:"Items,omitnil,omitempty" name:"Items"`
}

// Predefined struct for user
type ListCodeFolderContentsRequestParams struct {
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Parent folder path, such as /aaa/bbb/ccc. path header must include a slash. pass / for the root directory.
	ParentFolderPath *string `json:"ParentFolderPath,omitnil,omitempty" name:"ParentFolderPath"`

	// Folder name or code file name.
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// Queries folder only.
	OnlyFolderNode *bool `json:"OnlyFolderNode,omitnil,omitempty" name:"OnlyFolderNode"`

	// Specifies whether to query only code script created by user themselves.
	OnlyUserSelfScript *bool `json:"OnlyUserSelfScript,omitnil,omitempty" name:"OnlyUserSelfScript"`
}

type ListCodeFolderContentsRequest struct {
	*tchttp.BaseRequest
	
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Parent folder path, such as /aaa/bbb/ccc. path header must include a slash. pass / for the root directory.
	ParentFolderPath *string `json:"ParentFolderPath,omitnil,omitempty" name:"ParentFolderPath"`

	// Folder name or code file name.
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// Queries folder only.
	OnlyFolderNode *bool `json:"OnlyFolderNode,omitnil,omitempty" name:"OnlyFolderNode"`

	// Specifies whether to query only code script created by user themselves.
	OnlyUserSelfScript *bool `json:"OnlyUserSelfScript,omitnil,omitempty" name:"OnlyUserSelfScript"`
}

func (r *ListCodeFolderContentsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListCodeFolderContentsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "ParentFolderPath")
	delete(f, "Keyword")
	delete(f, "OnlyFolderNode")
	delete(f, "OnlyUserSelfScript")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListCodeFolderContentsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListCodeFolderContentsResponseParams struct {
	// Result
	Data []*CodeFolderNode `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListCodeFolderContentsResponse struct {
	*tchttp.BaseResponse
	Response *ListCodeFolderContentsResponseParams `json:"Response"`
}

func (r *ListCodeFolderContentsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListCodeFolderContentsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListDataBackfillInstancesRequestParams struct {
	// Project ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Backfill plan Id.
	DataBackfillPlanId *string `json:"DataBackfillPlanId,omitnil,omitempty" name:"DataBackfillPlanId"`

	// Task ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Page number.
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Pagination size.
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type ListDataBackfillInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Project ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Backfill plan Id.
	DataBackfillPlanId *string `json:"DataBackfillPlanId,omitnil,omitempty" name:"DataBackfillPlanId"`

	// Task ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Page number.
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Pagination size.
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

func (r *ListDataBackfillInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListDataBackfillInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "DataBackfillPlanId")
	delete(f, "TaskId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListDataBackfillInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListDataBackfillInstancesResponseParams struct {
	// All backfill  instances under one backfill  plan.
	Data *BackfillInstanceCollection `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListDataBackfillInstancesResponse struct {
	*tchttp.BaseResponse
	Response *ListDataBackfillInstancesResponseParams `json:"Response"`
}

func (r *ListDataBackfillInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListDataBackfillInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListDownstreamOpsTasksRequestParams struct {
	// Task ID		
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Project ID.
	// 		
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Page number
	PageNumber *string `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Pagination size.
	PageSize *string `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type ListDownstreamOpsTasksRequest struct {
	*tchttp.BaseRequest
	
	// Task ID		
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Project ID.
	// 		
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Page number
	PageNumber *string `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Pagination size.
	PageSize *string `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

func (r *ListDownstreamOpsTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListDownstreamOpsTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "ProjectId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListDownstreamOpsTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListDownstreamOpsTasksResponseParams struct {
	// Downstream dependency description.
	Data *ChildDependencyConfigPage `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListDownstreamOpsTasksResponse struct {
	*tchttp.BaseResponse
	Response *ListDownstreamOpsTasksResponseParams `json:"Response"`
}

func (r *ListDownstreamOpsTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListDownstreamOpsTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListDownstreamTaskInstancesRequestParams struct {
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// **Instance unique id**.
	InstanceKey *string `json:"InstanceKey,omitnil,omitempty" name:"InstanceKey"`

	// **Time zone** timeZone, default UTC+8.
	TimeZone *string `json:"TimeZone,omitnil,omitempty" name:"TimeZone"`

	// **Page number, int** used in conjunction with pageSize and cannot be less than 1. default value: 1.
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Specifies the number of items to display per page. default: 10. value range: 1-100.
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type ListDownstreamTaskInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// **Instance unique id**.
	InstanceKey *string `json:"InstanceKey,omitnil,omitempty" name:"InstanceKey"`

	// **Time zone** timeZone, default UTC+8.
	TimeZone *string `json:"TimeZone,omitnil,omitempty" name:"TimeZone"`

	// **Page number, int** used in conjunction with pageSize and cannot be less than 1. default value: 1.
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Specifies the number of items to display per page. default: 10. value range: 1-100.
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

func (r *ListDownstreamTaskInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListDownstreamTaskInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "InstanceKey")
	delete(f, "TimeZone")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListDownstreamTaskInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListDownstreamTaskInstancesResponseParams struct {
	// Direct downstream instance list.
	Data *TaskInstancePage `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListDownstreamTaskInstancesResponse struct {
	*tchttp.BaseResponse
	Response *ListDownstreamTaskInstancesResponseParams `json:"Response"`
}

func (r *ListDownstreamTaskInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListDownstreamTaskInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListDownstreamTasksRequestParams struct {
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Task ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Page number
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Pagination size.
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type ListDownstreamTasksRequest struct {
	*tchttp.BaseRequest
	
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Task ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Page number
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Pagination size.
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

func (r *ListDownstreamTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListDownstreamTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "TaskId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListDownstreamTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListDownstreamTasksResponseParams struct {
	// Describes the downstream dependency details.
	Data *DependencyConfigPage `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListDownstreamTasksResponse struct {
	*tchttp.BaseResponse
	Response *ListDownstreamTasksResponseParams `json:"Response"`
}

func (r *ListDownstreamTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListDownstreamTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListOpsAlarmRulesRequestParams struct {
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Page number for pagination. defaults to 1.
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Number of items to display per page, defaults to 20, minimum value 1, maximum value 200.
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Monitoring Object Type
	// 
	// Task-level Monitoring - Can be configured by Task / Workflow / Project:
	// 1 = Task (default)
	// 2 = Workflow
	// 3 = Project
	// 
	// Project-level Monitoring - Monitors overall task fluctuations within a project:
	// 7 = Project fluctuation monitoring alarm
	MonitorObjectType *int64 `json:"MonitorObjectType,omitnil,omitempty" name:"MonitorObjectType"`

	// Based on task id, queries Alarm rules.
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Alarm Rule Monitoring Types:
	// 
	// failure: Failure alarm
	// 
	// overtime: Timeout alarm
	// 
	// success: Success alarm
	// 
	// backTrackingOrRerunSuccess: Alarm when backfill/rerun succeeds
	// 
	// backTrackingOrRerunFailure: Alarm when backfill/rerun fails
	// 
	// projectFailureInstanceUpwardFluctuationAlarm: Alarm when the upward fluctuation rate of failed instances for the day exceeds the threshold
	// 
	// projectSuccessInstanceDownwardFluctuationAlarm: Alarm when the downward fluctuation rate of successful instances for the day exceeds the threshold
	// 
	// reconciliationFailure: Alarm when an offline reconciliation task fails
	// 
	// reconciliationOvertime: Alarm when an offline reconciliation task runs overtime
	// 
	// reconciliationMismatch: Alarm when the number of mismatched records in reconciliation exceeds the threshold
	AlarmType *string `json:"AlarmType,omitnil,omitempty" name:"AlarmType"`

	// Queries Alarm rules configured with corresponding Alarm levels.
	// Valid values: 1. ordinary, 2. important, 3. critical.
	AlarmLevel *int64 `json:"AlarmLevel,omitnil,omitempty" name:"AlarmLevel"`

	// Query the alarm rules associated with the configured alarm recipients.
	AlarmRecipientId *string `json:"AlarmRecipientId,omitnil,omitempty" name:"AlarmRecipientId"`

	// Queries the corresponding Alarm rule based on Alarm rule id or rule name.
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// Specifies the creator filter for Alarm rule creation.
	CreateUserUin *string `json:"CreateUserUin,omitnil,omitempty" name:"CreateUserUin"`

	// Start time of the Alarm rule create time range, in the format of 2025-08-17 00:00:00.
	CreateTimeFrom *string `json:"CreateTimeFrom,omitnil,omitempty" name:"CreateTimeFrom"`

	// End time of the Alarm rule creation time range, in the format of "2025-08-26 23:59:59".
	CreateTimeTo *string `json:"CreateTimeTo,omitnil,omitempty" name:"CreateTimeTo"`

	// Filters Alarm rules by last update time, in the format of "2025-08-26 00:00:00".
	UpdateTimeFrom *string `json:"UpdateTimeFrom,omitnil,omitempty" name:"UpdateTimeFrom"`

	// Filters Alarm rules by last update time in the format of "2025-08-26 23:59:59".
	UpdateTimeTo *string `json:"UpdateTimeTo,omitnil,omitempty" name:"UpdateTimeTo"`
}

type ListOpsAlarmRulesRequest struct {
	*tchttp.BaseRequest
	
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Page number for pagination. defaults to 1.
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Number of items to display per page, defaults to 20, minimum value 1, maximum value 200.
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Monitoring Object Type
	// 
	// Task-level Monitoring - Can be configured by Task / Workflow / Project:
	// 1 = Task (default)
	// 2 = Workflow
	// 3 = Project
	// 
	// Project-level Monitoring - Monitors overall task fluctuations within a project:
	// 7 = Project fluctuation monitoring alarm
	MonitorObjectType *int64 `json:"MonitorObjectType,omitnil,omitempty" name:"MonitorObjectType"`

	// Based on task id, queries Alarm rules.
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Alarm Rule Monitoring Types:
	// 
	// failure: Failure alarm
	// 
	// overtime: Timeout alarm
	// 
	// success: Success alarm
	// 
	// backTrackingOrRerunSuccess: Alarm when backfill/rerun succeeds
	// 
	// backTrackingOrRerunFailure: Alarm when backfill/rerun fails
	// 
	// projectFailureInstanceUpwardFluctuationAlarm: Alarm when the upward fluctuation rate of failed instances for the day exceeds the threshold
	// 
	// projectSuccessInstanceDownwardFluctuationAlarm: Alarm when the downward fluctuation rate of successful instances for the day exceeds the threshold
	// 
	// reconciliationFailure: Alarm when an offline reconciliation task fails
	// 
	// reconciliationOvertime: Alarm when an offline reconciliation task runs overtime
	// 
	// reconciliationMismatch: Alarm when the number of mismatched records in reconciliation exceeds the threshold
	AlarmType *string `json:"AlarmType,omitnil,omitempty" name:"AlarmType"`

	// Queries Alarm rules configured with corresponding Alarm levels.
	// Valid values: 1. ordinary, 2. important, 3. critical.
	AlarmLevel *int64 `json:"AlarmLevel,omitnil,omitempty" name:"AlarmLevel"`

	// Query the alarm rules associated with the configured alarm recipients.
	AlarmRecipientId *string `json:"AlarmRecipientId,omitnil,omitempty" name:"AlarmRecipientId"`

	// Queries the corresponding Alarm rule based on Alarm rule id or rule name.
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// Specifies the creator filter for Alarm rule creation.
	CreateUserUin *string `json:"CreateUserUin,omitnil,omitempty" name:"CreateUserUin"`

	// Start time of the Alarm rule create time range, in the format of 2025-08-17 00:00:00.
	CreateTimeFrom *string `json:"CreateTimeFrom,omitnil,omitempty" name:"CreateTimeFrom"`

	// End time of the Alarm rule creation time range, in the format of "2025-08-26 23:59:59".
	CreateTimeTo *string `json:"CreateTimeTo,omitnil,omitempty" name:"CreateTimeTo"`

	// Filters Alarm rules by last update time, in the format of "2025-08-26 00:00:00".
	UpdateTimeFrom *string `json:"UpdateTimeFrom,omitnil,omitempty" name:"UpdateTimeFrom"`

	// Filters Alarm rules by last update time in the format of "2025-08-26 23:59:59".
	UpdateTimeTo *string `json:"UpdateTimeTo,omitnil,omitempty" name:"UpdateTimeTo"`
}

func (r *ListOpsAlarmRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListOpsAlarmRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "MonitorObjectType")
	delete(f, "TaskId")
	delete(f, "AlarmType")
	delete(f, "AlarmLevel")
	delete(f, "AlarmRecipientId")
	delete(f, "Keyword")
	delete(f, "CreateUserUin")
	delete(f, "CreateTimeFrom")
	delete(f, "CreateTimeTo")
	delete(f, "UpdateTimeFrom")
	delete(f, "UpdateTimeTo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListOpsAlarmRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListOpsAlarmRulesResponseParams struct {
	// Alarm information response.
	Data *ListAlarmRulesResult `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListOpsAlarmRulesResponse struct {
	*tchttp.BaseResponse
	Response *ListOpsAlarmRulesResponseParams `json:"Response"`
}

func (r *ListOpsAlarmRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListOpsAlarmRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListOpsTasksPage struct {
	// Total number of results
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Total pages
	TotalPageNumber *uint64 `json:"TotalPageNumber,omitnil,omitempty" name:"TotalPageNumber"`

	// Record list	
	// 	
	Items []*TaskOpsInfo `json:"Items,omitnil,omitempty" name:"Items"`

	// Page number.
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Pagination size.
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

// Predefined struct for user
type ListOpsTasksRequestParams struct {
	// Project ID.
	// 		
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Pagination size.
	PageSize *string `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Page number
	PageNumber *string `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Task type Id. 
	// -20: common data sync.
	//  - 25:ETLTaskType
	//  - 26:ETLTaskType
	//  - 30:python
	//  - 31:pyspark
	//  - 34:HiveSQLTaskType
	//  - 35:shell
	//  - 36:SparkSQLTaskType
	//  - 21:JDBCSQLTaskType
	//  - 32:DLCTaskType
	//  - 33:ImpalaTaskType
	//  - 40:CDWTaskType
	//  - 41:kettle
	//  - 46:DLCSparkTaskType
	// -47: TiOne machine learning.
	//  - 48:TrinoTaskType
	//  - 50:DLCPyspark39:spark
	//  - 92:mr
	// -38: shell script.
	// -70: hivesql script.
	// -1000: common custom business.
	TaskTypeId *string `json:"TaskTypeId,omitnil,omitempty" name:"TaskTypeId"`

	// Workflow ID.
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// Workflow name.
	WorkflowName *string `json:"WorkflowName,omitnil,omitempty" name:"WorkflowName"`

	// Owner id.
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// Folder ID
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`

	// Data source ID.
	SourceServiceId *string `json:"SourceServiceId,omitnil,omitempty" name:"SourceServiceId"`

	// Target data source id.
	TargetServiceId *string `json:"TargetServiceId,omitnil,omitempty" name:"TargetServiceId"`

	// Executor Group ID
	ExecutorGroupId *string `json:"ExecutorGroupId,omitnil,omitempty" name:"ExecutorGroupId"`

	// Task Cycle Type:
	// 
	// * ONEOFF_CYCLE: One-time
	// 
	// * YEAR_CYCLE: Yearly
	// 
	// * MONTH_CYCLE: Monthly
	// 
	// * WEEK_CYCLE: Weekly
	// 
	// * DAY_CYCLE: Daily
	// 
	// * HOUR_CYCLE: Hourly
	// 
	// * MINUTE_CYCLE: Minute-level
	// 
	// * CRONTAB_CYCLE: Crontab expression-based
	CycleType *string `json:"CycleType,omitnil,omitempty" name:"CycleType"`

	// Task Status
	// 
	// -Y: Running
	// 
	// -F: Stopped
	// 
	// -O: Frozen
	// 
	// -T: Stopping
	// 
	// -INVALID: Invalid
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Time zone. defaults to UTC+8.
	TimeZone *string `json:"TimeZone,omitnil,omitempty" name:"TimeZone"`
}

type ListOpsTasksRequest struct {
	*tchttp.BaseRequest
	
	// Project ID.
	// 		
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Pagination size.
	PageSize *string `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Page number
	PageNumber *string `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Task type Id. 
	// -20: common data sync.
	//  - 25:ETLTaskType
	//  - 26:ETLTaskType
	//  - 30:python
	//  - 31:pyspark
	//  - 34:HiveSQLTaskType
	//  - 35:shell
	//  - 36:SparkSQLTaskType
	//  - 21:JDBCSQLTaskType
	//  - 32:DLCTaskType
	//  - 33:ImpalaTaskType
	//  - 40:CDWTaskType
	//  - 41:kettle
	//  - 46:DLCSparkTaskType
	// -47: TiOne machine learning.
	//  - 48:TrinoTaskType
	//  - 50:DLCPyspark39:spark
	//  - 92:mr
	// -38: shell script.
	// -70: hivesql script.
	// -1000: common custom business.
	TaskTypeId *string `json:"TaskTypeId,omitnil,omitempty" name:"TaskTypeId"`

	// Workflow ID.
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// Workflow name.
	WorkflowName *string `json:"WorkflowName,omitnil,omitempty" name:"WorkflowName"`

	// Owner id.
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// Folder ID
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`

	// Data source ID.
	SourceServiceId *string `json:"SourceServiceId,omitnil,omitempty" name:"SourceServiceId"`

	// Target data source id.
	TargetServiceId *string `json:"TargetServiceId,omitnil,omitempty" name:"TargetServiceId"`

	// Executor Group ID
	ExecutorGroupId *string `json:"ExecutorGroupId,omitnil,omitempty" name:"ExecutorGroupId"`

	// Task Cycle Type:
	// 
	// * ONEOFF_CYCLE: One-time
	// 
	// * YEAR_CYCLE: Yearly
	// 
	// * MONTH_CYCLE: Monthly
	// 
	// * WEEK_CYCLE: Weekly
	// 
	// * DAY_CYCLE: Daily
	// 
	// * HOUR_CYCLE: Hourly
	// 
	// * MINUTE_CYCLE: Minute-level
	// 
	// * CRONTAB_CYCLE: Crontab expression-based
	CycleType *string `json:"CycleType,omitnil,omitempty" name:"CycleType"`

	// Task Status
	// 
	// -Y: Running
	// 
	// -F: Stopped
	// 
	// -O: Frozen
	// 
	// -T: Stopping
	// 
	// -INVALID: Invalid
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Time zone. defaults to UTC+8.
	TimeZone *string `json:"TimeZone,omitnil,omitempty" name:"TimeZone"`
}

func (r *ListOpsTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListOpsTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "PageSize")
	delete(f, "PageNumber")
	delete(f, "TaskTypeId")
	delete(f, "WorkflowId")
	delete(f, "WorkflowName")
	delete(f, "OwnerUin")
	delete(f, "FolderId")
	delete(f, "SourceServiceId")
	delete(f, "TargetServiceId")
	delete(f, "ExecutorGroupId")
	delete(f, "CycleType")
	delete(f, "Status")
	delete(f, "TimeZone")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListOpsTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListOpsTasksResponseParams struct {
	// Task list.
	Data *ListOpsTasksPage `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListOpsTasksResponse struct {
	*tchttp.BaseResponse
	Response *ListOpsTasksResponseParams `json:"Response"`
}

func (r *ListOpsTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListOpsTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListOpsWorkflowsRequestParams struct {
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Page number
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Pagination size.
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Folder ID
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`

	// Workflow Status Filter
	// 
	// * ALL_RUNNING: All workflows are running (scheduled)
	// 
	// * ALL_FREEZED: All workflows are paused
	// 
	// * ALL_STOPPED: All workflows are offline
	// 
	// * PART_RUNNING: Some workflows are running (partially scheduled)
	// 
	// * ALL_NO_RUNNING: No workflows are running (unscheduled)
	// 
	// * ALL_INVALID: All workflows are invalid
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Owner ID
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// Workflow type filter criteria. supported values: Cycle or Manual. default: Cycle.
	WorkflowType *string `json:"WorkflowType,omitnil,omitempty" name:"WorkflowType"`

	// Workflow keyword-based filtering supports fuzzy matching of workflow Id/name.
	KeyWord *string `json:"KeyWord,omitnil,omitempty" name:"KeyWord"`

	// Sort item. Options: CreateTime, TaskCount.
	SortItem *string `json:"SortItem,omitnil,omitempty" name:"SortItem"`

	// Sorting method, DESC or ASC, uppercase.
	SortType *string `json:"SortType,omitnil,omitempty" name:"SortType"`

	// CreatorId. specifies the id of the creator.
	CreateUserUin *string `json:"CreateUserUin,omitnil,omitempty" name:"CreateUserUin"`

	// Update time. format: yyyy-MM-dd HH:MM:ss.
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`

	// Creation time. format: yyyy-MM-dd HH:MM:ss.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`
}

type ListOpsWorkflowsRequest struct {
	*tchttp.BaseRequest
	
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Page number
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Pagination size.
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Folder ID
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`

	// Workflow Status Filter
	// 
	// * ALL_RUNNING: All workflows are running (scheduled)
	// 
	// * ALL_FREEZED: All workflows are paused
	// 
	// * ALL_STOPPED: All workflows are offline
	// 
	// * PART_RUNNING: Some workflows are running (partially scheduled)
	// 
	// * ALL_NO_RUNNING: No workflows are running (unscheduled)
	// 
	// * ALL_INVALID: All workflows are invalid
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Owner ID
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// Workflow type filter criteria. supported values: Cycle or Manual. default: Cycle.
	WorkflowType *string `json:"WorkflowType,omitnil,omitempty" name:"WorkflowType"`

	// Workflow keyword-based filtering supports fuzzy matching of workflow Id/name.
	KeyWord *string `json:"KeyWord,omitnil,omitempty" name:"KeyWord"`

	// Sort item. Options: CreateTime, TaskCount.
	SortItem *string `json:"SortItem,omitnil,omitempty" name:"SortItem"`

	// Sorting method, DESC or ASC, uppercase.
	SortType *string `json:"SortType,omitnil,omitempty" name:"SortType"`

	// CreatorId. specifies the id of the creator.
	CreateUserUin *string `json:"CreateUserUin,omitnil,omitempty" name:"CreateUserUin"`

	// Update time. format: yyyy-MM-dd HH:MM:ss.
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`

	// Creation time. format: yyyy-MM-dd HH:MM:ss.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`
}

func (r *ListOpsWorkflowsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListOpsWorkflowsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "FolderId")
	delete(f, "Status")
	delete(f, "OwnerUin")
	delete(f, "WorkflowType")
	delete(f, "KeyWord")
	delete(f, "SortItem")
	delete(f, "SortType")
	delete(f, "CreateUserUin")
	delete(f, "ModifyTime")
	delete(f, "CreateTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListOpsWorkflowsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListOpsWorkflowsResponseParams struct {
	// Workflow list
	Data *OpsWorkflows `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListOpsWorkflowsResponse struct {
	*tchttp.BaseResponse
	Response *ListOpsWorkflowsResponseParams `json:"Response"`
}

func (r *ListOpsWorkflowsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListOpsWorkflowsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListResourceFilesRequestParams struct {
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Data page number, equal to or greater than 1. default 1.
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Specifies the number of data records displayed per page. valid values: 10 to 200. default: 10.
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Resource file name (fuzzy search keyword).
	ResourceName *string `json:"ResourceName,omitnil,omitempty" name:"ResourceName"`

	// Specifies the path of the file's parent folder (for example /a/b/c, querying resource files under the folder c).
	ParentFolderPath *string `json:"ParentFolderPath,omitnil,omitempty" name:"ParentFolderPath"`

	// Creator ID. obtain through the DescribeCurrentUserInfo API.
	CreateUserUin *string `json:"CreateUserUin,omitnil,omitempty" name:"CreateUserUin"`

	// Update time range. specifies the start time in yyyy-MM-dd HH:MM:ss format.
	ModifyTimeStart *string `json:"ModifyTimeStart,omitnil,omitempty" name:"ModifyTimeStart"`

	// Update time range. specifies the end time in yyyy-MM-dd HH:MM:ss format.
	ModifyTimeEnd *string `json:"ModifyTimeEnd,omitnil,omitempty" name:"ModifyTimeEnd"`

	// Create time range. specifies the start time in yyyy-MM-dd HH:MM:ss format.
	CreateTimeStart *string `json:"CreateTimeStart,omitnil,omitempty" name:"CreateTimeStart"`

	// Create time range. specifies the termination time in yyyy-MM-dd HH:MM:ss format.
	CreateTimeEnd *string `json:"CreateTimeEnd,omitnil,omitempty" name:"CreateTimeEnd"`
}

type ListResourceFilesRequest struct {
	*tchttp.BaseRequest
	
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Data page number, equal to or greater than 1. default 1.
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Specifies the number of data records displayed per page. valid values: 10 to 200. default: 10.
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Resource file name (fuzzy search keyword).
	ResourceName *string `json:"ResourceName,omitnil,omitempty" name:"ResourceName"`

	// Specifies the path of the file's parent folder (for example /a/b/c, querying resource files under the folder c).
	ParentFolderPath *string `json:"ParentFolderPath,omitnil,omitempty" name:"ParentFolderPath"`

	// Creator ID. obtain through the DescribeCurrentUserInfo API.
	CreateUserUin *string `json:"CreateUserUin,omitnil,omitempty" name:"CreateUserUin"`

	// Update time range. specifies the start time in yyyy-MM-dd HH:MM:ss format.
	ModifyTimeStart *string `json:"ModifyTimeStart,omitnil,omitempty" name:"ModifyTimeStart"`

	// Update time range. specifies the end time in yyyy-MM-dd HH:MM:ss format.
	ModifyTimeEnd *string `json:"ModifyTimeEnd,omitnil,omitempty" name:"ModifyTimeEnd"`

	// Create time range. specifies the start time in yyyy-MM-dd HH:MM:ss format.
	CreateTimeStart *string `json:"CreateTimeStart,omitnil,omitempty" name:"CreateTimeStart"`

	// Create time range. specifies the termination time in yyyy-MM-dd HH:MM:ss format.
	CreateTimeEnd *string `json:"CreateTimeEnd,omitnil,omitempty" name:"CreateTimeEnd"`
}

func (r *ListResourceFilesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListResourceFilesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "ResourceName")
	delete(f, "ParentFolderPath")
	delete(f, "CreateUserUin")
	delete(f, "ModifyTimeStart")
	delete(f, "ModifyTimeEnd")
	delete(f, "CreateTimeStart")
	delete(f, "CreateTimeEnd")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListResourceFilesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListResourceFilesResponseParams struct {
	// Retrieve the resource file list.
	Data *ResourceFilePage `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListResourceFilesResponse struct {
	*tchttp.BaseResponse
	Response *ListResourceFilesResponseParams `json:"Response"`
}

func (r *ListResourceFilesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListResourceFilesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListResourceFoldersRequestParams struct {
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Resource folder absolute path. value example.
	// /wedata/test
	ParentFolderPath *string `json:"ParentFolderPath,omitnil,omitempty" name:"ParentFolderPath"`

	// Data page number, equal to or greater than 1. default 1.
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Specifies the number of data records displayed per page. valid values: 10 to 200. default: 10.
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type ListResourceFoldersRequest struct {
	*tchttp.BaseRequest
	
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Resource folder absolute path. value example.
	// /wedata/test
	ParentFolderPath *string `json:"ParentFolderPath,omitnil,omitempty" name:"ParentFolderPath"`

	// Data page number, equal to or greater than 1. default 1.
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Specifies the number of data records displayed per page. valid values: 10 to 200. default: 10.
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

func (r *ListResourceFoldersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListResourceFoldersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "ParentFolderPath")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListResourceFoldersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListResourceFoldersResponseParams struct {
	// Paginated resource folder query result.
	Data *ResourceFolderPage `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListResourceFoldersResponse struct {
	*tchttp.BaseResponse
	Response *ListResourceFoldersResponseParams `json:"Response"`
}

func (r *ListResourceFoldersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListResourceFoldersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListSQLFolderContentsRequestParams struct {
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Parent folder path, /aaa/bbb/ccc. path header must include a slash. pass / to query the root directory.
	ParentFolderPath *string `json:"ParentFolderPath,omitnil,omitempty" name:"ParentFolderPath"`

	// Folder name or script name search.
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// Queries only the folder.
	OnlyFolderNode *bool `json:"OnlyFolderNode,omitnil,omitempty" name:"OnlyFolderNode"`

	// Specifies whether to query only scripts created by the user themselves.
	OnlyUserSelfScript *bool `json:"OnlyUserSelfScript,omitnil,omitempty" name:"OnlyUserSelfScript"`

	// Access permission: SHARED, PRIVATE.
	AccessScope *string `json:"AccessScope,omitnil,omitempty" name:"AccessScope"`
}

type ListSQLFolderContentsRequest struct {
	*tchttp.BaseRequest
	
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Parent folder path, /aaa/bbb/ccc. path header must include a slash. pass / to query the root directory.
	ParentFolderPath *string `json:"ParentFolderPath,omitnil,omitempty" name:"ParentFolderPath"`

	// Folder name or script name search.
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// Queries only the folder.
	OnlyFolderNode *bool `json:"OnlyFolderNode,omitnil,omitempty" name:"OnlyFolderNode"`

	// Specifies whether to query only scripts created by the user themselves.
	OnlyUserSelfScript *bool `json:"OnlyUserSelfScript,omitnil,omitempty" name:"OnlyUserSelfScript"`

	// Access permission: SHARED, PRIVATE.
	AccessScope *string `json:"AccessScope,omitnil,omitempty" name:"AccessScope"`
}

func (r *ListSQLFolderContentsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListSQLFolderContentsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "ParentFolderPath")
	delete(f, "Keyword")
	delete(f, "OnlyFolderNode")
	delete(f, "OnlyUserSelfScript")
	delete(f, "AccessScope")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListSQLFolderContentsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListSQLFolderContentsResponseParams struct {
	// Result list
	Data []*SQLFolderNode `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListSQLFolderContentsResponse struct {
	*tchttp.BaseResponse
	Response *ListSQLFolderContentsResponseParams `json:"Response"`
}

func (r *ListSQLFolderContentsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListSQLFolderContentsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListSQLScriptRunsRequestParams struct {
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Script id.
	ScriptId *string `json:"ScriptId,omitnil,omitempty" name:"ScriptId"`

	// Task ID.
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// Search keywords.
	SearchWord *string `json:"SearchWord,omitnil,omitempty" name:"SearchWord"`

	// Specifies the executor ID.
	ExecuteUserUin *string `json:"ExecuteUserUin,omitnil,omitempty" name:"ExecuteUserUin"`

	// Start time.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type ListSQLScriptRunsRequest struct {
	*tchttp.BaseRequest
	
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Script id.
	ScriptId *string `json:"ScriptId,omitnil,omitempty" name:"ScriptId"`

	// Task ID.
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// Search keywords.
	SearchWord *string `json:"SearchWord,omitnil,omitempty" name:"SearchWord"`

	// Specifies the executor ID.
	ExecuteUserUin *string `json:"ExecuteUserUin,omitnil,omitempty" name:"ExecuteUserUin"`

	// Start time.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

func (r *ListSQLScriptRunsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListSQLScriptRunsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "ScriptId")
	delete(f, "JobId")
	delete(f, "SearchWord")
	delete(f, "ExecuteUserUin")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListSQLScriptRunsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListSQLScriptRunsResponseParams struct {
	// Data Exploration runs.
	Data []*JobDto `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListSQLScriptRunsResponse struct {
	*tchttp.BaseResponse
	Response *ListSQLScriptRunsResponseParams `json:"Response"`
}

func (r *ListSQLScriptRunsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListSQLScriptRunsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListTaskInfo struct {
	// Task Array
	Items []*TaskBaseAttribute `json:"Items,omitnil,omitempty" name:"Items"`

	// Current request data page number.
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Number of entries in the current request.
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Total number of data entries that meet the query condition.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Total number of pages that meet the query condition.
	TotalPageNumber *uint64 `json:"TotalPageNumber,omitnil,omitempty" name:"TotalPageNumber"`
}

// Predefined struct for user
type ListTaskInstanceExecutionsRequestParams struct {
	// Project id.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Instance unique id, can be obtained through ListInstances.
	InstanceKey *string `json:"InstanceKey,omitnil,omitempty" name:"InstanceKey"`

	// **Time zone** timeZone, specifies the time zone of the passed in time string. default UTC+8.
	TimeZone *string `json:"TimeZone,omitnil,omitempty" name:"TimeZone"`

	// Size per page. default: 10. maximum: 200.
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Page number, which is 1 by default.
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`
}

type ListTaskInstanceExecutionsRequest struct {
	*tchttp.BaseRequest
	
	// Project id.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Instance unique id, can be obtained through ListInstances.
	InstanceKey *string `json:"InstanceKey,omitnil,omitempty" name:"InstanceKey"`

	// **Time zone** timeZone, specifies the time zone of the passed in time string. default UTC+8.
	TimeZone *string `json:"TimeZone,omitnil,omitempty" name:"TimeZone"`

	// Size per page. default: 10. maximum: 200.
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Page number, which is 1 by default.
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`
}

func (r *ListTaskInstanceExecutionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListTaskInstanceExecutionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "InstanceKey")
	delete(f, "TimeZone")
	delete(f, "PageSize")
	delete(f, "PageNumber")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListTaskInstanceExecutionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListTaskInstanceExecutionsResponseParams struct {
	// Instance details.
	Data *TaskInstanceExecutions `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListTaskInstanceExecutionsResponse struct {
	*tchttp.BaseResponse
	Response *ListTaskInstanceExecutionsResponseParams `json:"Response"`
}

func (r *ListTaskInstanceExecutionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListTaskInstanceExecutionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListTaskInstancesRequestParams struct {
	// **Project ID**. specifies the project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// **Page number. integer.**.
	// Used in conjunction with pageSize and cannot be less than 1. default value: 1.
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Specifies the number of items to display per page. default: 10. value range: 1-100.
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Task name or task ID.
	// Supports fuzzy search filtering. multiple values are separated by commas.
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// **Time zone** timeZone, specifies the time zone of the passed in time string. default UTC+8.
	TimeZone *string `json:"TimeZone,omitnil,omitempty" name:"TimeZone"`

	// **Instance Type**
	// 
	// 0 - Backfill instance
	// 
	// 1 - Cyclic (scheduled) instance
	// 
	// 2 - Non-cyclic (non-scheduled) instance
	InstanceType *uint64 `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// **Instance Status**
	// 
	// WAIT_EVENT: Waiting for event
	// 
	// WAIT_UPSTREAM: Waiting for upstream
	// 
	// WAIT_RUN: Waiting to run
	// 
	// RUNNING: Running
	// 
	// SKIP_RUNNING: Skipped
	// 
	// FAILED_RETRY: Retrying after failure
	// 
	// EXPIRED: Failed
	// 
	// COMPLETED: Succeeded
	InstanceState *string `json:"InstanceState,omitnil,omitempty" name:"InstanceState"`

	// Task type Id.
	TaskTypeId *uint64 `json:"TaskTypeId,omitnil,omitempty" name:"TaskTypeId"`

	// **Task Cycle Type**
	// 
	// ONEOFF_CYCLE: One-time
	// 
	// YEAR_CYCLE: Yearly
	// 
	// MONTH_CYCLE: Monthly
	// 
	// WEEK_CYCLE: Weekly
	// 
	// DAY_CYCLE: Daily
	// 
	// HOUR_CYCLE: Hourly
	// 
	// MINUTE_CYCLE: Minute-level
	// 
	// CRONTAB_CYCLE: Crontab expression-based
	CycleType *string `json:"CycleType,omitnil,omitempty" name:"CycleType"`

	// Task owner id.
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// Folder id
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`

	// Workflow id of the task.
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// **Execution resource group Id**.
	ExecutorGroupId *string `json:"ExecutorGroupId,omitnil,omitempty" name:"ExecutorGroupId"`

	// **Instance Scheduled Time Filter Condition** specifies the filter start time in the time format yyyy-MM-dd HH:MM:ss.
	ScheduleTimeFrom *string `json:"ScheduleTimeFrom,omitnil,omitempty" name:"ScheduleTimeFrom"`

	// **Instance Scheduled Time Filter Condition** specifies the cutoff time in the format of yyyy-MM-dd HH:MM:ss.
	ScheduleTimeTo *string `json:"ScheduleTimeTo,omitnil,omitempty" name:"ScheduleTimeTo"`

	// **Instance Execution Start Time Filter Condition** specifies the start time for filtering. time format: yyyy-MM-dd HH:MM:ss.
	StartTimeFrom *string `json:"StartTimeFrom,omitnil,omitempty" name:"StartTimeFrom"`

	// **Instance Execution Start Time Filter Condition**.
	// Expiration time in yyyy-MM-dd HH:MM:ss format.
	StartTimeTo *string `json:"StartTimeTo,omitnil,omitempty" name:"StartTimeTo"`

	// **Instance Last Update Time Filter Condition**.
	// Expiration time in yyyy-MM-dd HH:MM:ss format.
	LastUpdateTimeFrom *string `json:"LastUpdateTimeFrom,omitnil,omitempty" name:"LastUpdateTimeFrom"`

	// **Instance Last Update Time Filter Condition**.
	// Expiration time in yyyy-MM-dd HH:MM:ss format.
	LastUpdateTimeTo *string `json:"LastUpdateTimeTo,omitnil,omitempty" name:"LastUpdateTimeTo"`

	// **Query Result Sorting Field**
	// 
	// SCHEDULE_DATE: Sort by scheduled execution time
	// 
	// START_TIME: Sort by actual execution start time
	// 
	// END_TIME: Sort by actual execution end time
	// 
	// COST_TIME: Sort by execution duration
	SortColumn *string `json:"SortColumn,omitnil,omitempty" name:"SortColumn"`

	// **Instance Sorting Order**
	// 
	// - ASC 
	// - DESC
	SortType *string `json:"SortType,omitnil,omitempty" name:"SortType"`
}

type ListTaskInstancesRequest struct {
	*tchttp.BaseRequest
	
	// **Project ID**. specifies the project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// **Page number. integer.**.
	// Used in conjunction with pageSize and cannot be less than 1. default value: 1.
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Specifies the number of items to display per page. default: 10. value range: 1-100.
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Task name or task ID.
	// Supports fuzzy search filtering. multiple values are separated by commas.
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// **Time zone** timeZone, specifies the time zone of the passed in time string. default UTC+8.
	TimeZone *string `json:"TimeZone,omitnil,omitempty" name:"TimeZone"`

	// **Instance Type**
	// 
	// 0 - Backfill instance
	// 
	// 1 - Cyclic (scheduled) instance
	// 
	// 2 - Non-cyclic (non-scheduled) instance
	InstanceType *uint64 `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// **Instance Status**
	// 
	// WAIT_EVENT: Waiting for event
	// 
	// WAIT_UPSTREAM: Waiting for upstream
	// 
	// WAIT_RUN: Waiting to run
	// 
	// RUNNING: Running
	// 
	// SKIP_RUNNING: Skipped
	// 
	// FAILED_RETRY: Retrying after failure
	// 
	// EXPIRED: Failed
	// 
	// COMPLETED: Succeeded
	InstanceState *string `json:"InstanceState,omitnil,omitempty" name:"InstanceState"`

	// Task type Id.
	TaskTypeId *uint64 `json:"TaskTypeId,omitnil,omitempty" name:"TaskTypeId"`

	// **Task Cycle Type**
	// 
	// ONEOFF_CYCLE: One-time
	// 
	// YEAR_CYCLE: Yearly
	// 
	// MONTH_CYCLE: Monthly
	// 
	// WEEK_CYCLE: Weekly
	// 
	// DAY_CYCLE: Daily
	// 
	// HOUR_CYCLE: Hourly
	// 
	// MINUTE_CYCLE: Minute-level
	// 
	// CRONTAB_CYCLE: Crontab expression-based
	CycleType *string `json:"CycleType,omitnil,omitempty" name:"CycleType"`

	// Task owner id.
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// Folder id
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`

	// Workflow id of the task.
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// **Execution resource group Id**.
	ExecutorGroupId *string `json:"ExecutorGroupId,omitnil,omitempty" name:"ExecutorGroupId"`

	// **Instance Scheduled Time Filter Condition** specifies the filter start time in the time format yyyy-MM-dd HH:MM:ss.
	ScheduleTimeFrom *string `json:"ScheduleTimeFrom,omitnil,omitempty" name:"ScheduleTimeFrom"`

	// **Instance Scheduled Time Filter Condition** specifies the cutoff time in the format of yyyy-MM-dd HH:MM:ss.
	ScheduleTimeTo *string `json:"ScheduleTimeTo,omitnil,omitempty" name:"ScheduleTimeTo"`

	// **Instance Execution Start Time Filter Condition** specifies the start time for filtering. time format: yyyy-MM-dd HH:MM:ss.
	StartTimeFrom *string `json:"StartTimeFrom,omitnil,omitempty" name:"StartTimeFrom"`

	// **Instance Execution Start Time Filter Condition**.
	// Expiration time in yyyy-MM-dd HH:MM:ss format.
	StartTimeTo *string `json:"StartTimeTo,omitnil,omitempty" name:"StartTimeTo"`

	// **Instance Last Update Time Filter Condition**.
	// Expiration time in yyyy-MM-dd HH:MM:ss format.
	LastUpdateTimeFrom *string `json:"LastUpdateTimeFrom,omitnil,omitempty" name:"LastUpdateTimeFrom"`

	// **Instance Last Update Time Filter Condition**.
	// Expiration time in yyyy-MM-dd HH:MM:ss format.
	LastUpdateTimeTo *string `json:"LastUpdateTimeTo,omitnil,omitempty" name:"LastUpdateTimeTo"`

	// **Query Result Sorting Field**
	// 
	// SCHEDULE_DATE: Sort by scheduled execution time
	// 
	// START_TIME: Sort by actual execution start time
	// 
	// END_TIME: Sort by actual execution end time
	// 
	// COST_TIME: Sort by execution duration
	SortColumn *string `json:"SortColumn,omitnil,omitempty" name:"SortColumn"`

	// **Instance Sorting Order**
	// 
	// - ASC 
	// - DESC
	SortType *string `json:"SortType,omitnil,omitempty" name:"SortType"`
}

func (r *ListTaskInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListTaskInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "Keyword")
	delete(f, "TimeZone")
	delete(f, "InstanceType")
	delete(f, "InstanceState")
	delete(f, "TaskTypeId")
	delete(f, "CycleType")
	delete(f, "OwnerUin")
	delete(f, "FolderId")
	delete(f, "WorkflowId")
	delete(f, "ExecutorGroupId")
	delete(f, "ScheduleTimeFrom")
	delete(f, "ScheduleTimeTo")
	delete(f, "StartTimeFrom")
	delete(f, "StartTimeTo")
	delete(f, "LastUpdateTimeFrom")
	delete(f, "LastUpdateTimeTo")
	delete(f, "SortColumn")
	delete(f, "SortType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListTaskInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListTaskInstancesResponseParams struct {
	// Instance result set.
	Data *TaskInstancePage `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListTaskInstancesResponse struct {
	*tchttp.BaseResponse
	Response *ListTaskInstancesResponseParams `json:"Response"`
}

func (r *ListTaskInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListTaskInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListTaskVersions struct {
	// Record list	
	// 	
	Items []*TaskVersion `json:"Items,omitnil,omitempty" name:"Items"`

	// Total number of records that meet the query condition.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Total number of pages that meet the query condition.
	TotalPageNumber *uint64 `json:"TotalPageNumber,omitnil,omitempty" name:"TotalPageNumber"`

	// Number of records on current page.
	PageCount *uint64 `json:"PageCount,omitnil,omitempty" name:"PageCount"`

	// Specifies the number of entries in the current request data page.
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Specifies the data page number of the current request.
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`
}

// Predefined struct for user
type ListTaskVersionsRequestParams struct {
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Task ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// SAVE version.
	// SUBMIT version.
	// Defaults to SAVE.
	TaskVersionType *string `json:"TaskVersionType,omitnil,omitempty" name:"TaskVersionType"`

	// Specifies the data page number of the request. default value is 1. valid values: equal to or greater than 1.
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Specifies the number of data records displayed per page. default: 10. value range: 10 to 200.
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type ListTaskVersionsRequest struct {
	*tchttp.BaseRequest
	
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Task ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// SAVE version.
	// SUBMIT version.
	// Defaults to SAVE.
	TaskVersionType *string `json:"TaskVersionType,omitnil,omitempty" name:"TaskVersionType"`

	// Specifies the data page number of the request. default value is 1. valid values: equal to or greater than 1.
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Specifies the number of data records displayed per page. default: 10. value range: 10 to 200.
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

func (r *ListTaskVersionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListTaskVersionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "TaskId")
	delete(f, "TaskVersionType")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListTaskVersionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListTaskVersionsResponseParams struct {
	// Task version list.
	Data *ListTaskVersions `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListTaskVersionsResponse struct {
	*tchttp.BaseResponse
	Response *ListTaskVersionsResponseParams `json:"Response"`
}

func (r *ListTaskVersionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListTaskVersionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListTasksRequestParams struct {
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Specifies the data page number of the request. default value is 1. valid values: equal to or greater than 1.
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Specifies the number of data records displayed per page. default: 10. value range: 10 to 200.
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Task name.
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// Workflow ID
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// Owner ID.
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// Task type
	TaskTypeId *int64 `json:"TaskTypeId,omitnil,omitempty" name:"TaskTypeId"`

	// Task Status:
	// * N: New
	// * Y: Scheduling
	// * F: Offline
	// * O: Paused
	// * T: Offlining
	// * INVALID: Invalid
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Submission status.
	Submit *bool `json:"Submit,omitnil,omitempty" name:"Submit"`

	// Bundle id.
	BundleId *string `json:"BundleId,omitnil,omitempty" name:"BundleId"`

	// Creator ID.
	CreateUserUin *string `json:"CreateUserUin,omitnil,omitempty" name:"CreateUserUin"`

	// Modification time range (yyyy-MM-dd HH:mm:ss). Two time values must be provided in the array.
	ModifyTime []*string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`

	// Creation time range (yyyy-MM-dd HH:MM:ss). Two time values must be provided in the array.
	CreateTime []*string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`
}

type ListTasksRequest struct {
	*tchttp.BaseRequest
	
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Specifies the data page number of the request. default value is 1. valid values: equal to or greater than 1.
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Specifies the number of data records displayed per page. default: 10. value range: 10 to 200.
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Task name.
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// Workflow ID
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// Owner ID.
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// Task type
	TaskTypeId *int64 `json:"TaskTypeId,omitnil,omitempty" name:"TaskTypeId"`

	// Task Status:
	// * N: New
	// * Y: Scheduling
	// * F: Offline
	// * O: Paused
	// * T: Offlining
	// * INVALID: Invalid
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Submission status.
	Submit *bool `json:"Submit,omitnil,omitempty" name:"Submit"`

	// Bundle id.
	BundleId *string `json:"BundleId,omitnil,omitempty" name:"BundleId"`

	// Creator ID.
	CreateUserUin *string `json:"CreateUserUin,omitnil,omitempty" name:"CreateUserUin"`

	// Modification time range (yyyy-MM-dd HH:mm:ss). Two time values must be provided in the array.
	ModifyTime []*string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`

	// Creation time range (yyyy-MM-dd HH:MM:ss). Two time values must be provided in the array.
	CreateTime []*string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`
}

func (r *ListTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "TaskName")
	delete(f, "WorkflowId")
	delete(f, "OwnerUin")
	delete(f, "TaskTypeId")
	delete(f, "Status")
	delete(f, "Submit")
	delete(f, "BundleId")
	delete(f, "CreateUserUin")
	delete(f, "ModifyTime")
	delete(f, "CreateTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListTasksResponseParams struct {
	// Describes the task pagination information.
	Data *ListTaskInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListTasksResponse struct {
	*tchttp.BaseResponse
	Response *ListTasksResponseParams `json:"Response"`
}

func (r *ListTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListUpstreamOpsTasksRequestParams struct {
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Task ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Page number
	PageNumber *string `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Pagination size.
	PageSize *string `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type ListUpstreamOpsTasksRequest struct {
	*tchttp.BaseRequest
	
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Task ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Page number
	PageNumber *string `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Pagination size.
	PageSize *string `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

func (r *ListUpstreamOpsTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListUpstreamOpsTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "TaskId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListUpstreamOpsTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListUpstreamOpsTasksResponseParams struct {
	// Upstream task details.
	Data *ParentDependencyConfigPage `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListUpstreamOpsTasksResponse struct {
	*tchttp.BaseResponse
	Response *ListUpstreamOpsTasksResponseParams `json:"Response"`
}

func (r *ListUpstreamOpsTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListUpstreamOpsTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListUpstreamTaskInstancesRequestParams struct {
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// **Instance unique id**.
	InstanceKey *string `json:"InstanceKey,omitnil,omitempty" name:"InstanceKey"`

	// **Time zone** timeZone, default UTC+8.
	TimeZone *string `json:"TimeZone,omitnil,omitempty" name:"TimeZone"`

	// **Page number, int** used in conjunction with pageSize and cannot be less than 1. default value: 1.
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Specifies the number of items to display per page. default: 10. value range: 1-100.
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type ListUpstreamTaskInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// **Instance unique id**.
	InstanceKey *string `json:"InstanceKey,omitnil,omitempty" name:"InstanceKey"`

	// **Time zone** timeZone, default UTC+8.
	TimeZone *string `json:"TimeZone,omitnil,omitempty" name:"TimeZone"`

	// **Page number, int** used in conjunction with pageSize and cannot be less than 1. default value: 1.
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Specifies the number of items to display per page. default: 10. value range: 1-100.
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

func (r *ListUpstreamTaskInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListUpstreamTaskInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "InstanceKey")
	delete(f, "TimeZone")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListUpstreamTaskInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListUpstreamTaskInstancesResponseParams struct {
	// Upstream instance list.
	Data *TaskInstancePage `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListUpstreamTaskInstancesResponse struct {
	*tchttp.BaseResponse
	Response *ListUpstreamTaskInstancesResponseParams `json:"Response"`
}

func (r *ListUpstreamTaskInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListUpstreamTaskInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListUpstreamTasksRequestParams struct {
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Task ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Specifies the data page number of the request. Default value is 1. Valid values: equal to or greater than 1.
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Specifies the data page number of the request. Default value is 1. Valid values: equal to or greater than 1.
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type ListUpstreamTasksRequest struct {
	*tchttp.BaseRequest
	
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Task ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Specifies the data page number of the request. Default value is 1. Valid values: equal to or greater than 1.
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Specifies the data page number of the request. Default value is 1. Valid values: equal to or greater than 1.
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

func (r *ListUpstreamTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListUpstreamTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "TaskId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListUpstreamTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListUpstreamTasksResponseParams struct {
	// Upstream task details.
	Data *DependencyConfigPage `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListUpstreamTasksResponse struct {
	*tchttp.BaseResponse
	Response *ListUpstreamTasksResponseParams `json:"Response"`
}

func (r *ListUpstreamTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListUpstreamTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListWorkflowFoldersRequestParams struct {
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Parent folder absolute path, for example /abc/de, if it is root directory, pass /.
	ParentFolderPath *string `json:"ParentFolderPath,omitnil,omitempty" name:"ParentFolderPath"`

	// Data page number, equal to or greater than 1. default 1.
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Specifies the number of data records displayed per page. valid values: 10 to 200. default: 10.
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type ListWorkflowFoldersRequest struct {
	*tchttp.BaseRequest
	
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Parent folder absolute path, for example /abc/de, if it is root directory, pass /.
	ParentFolderPath *string `json:"ParentFolderPath,omitnil,omitempty" name:"ParentFolderPath"`

	// Data page number, equal to or greater than 1. default 1.
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Specifies the number of data records displayed per page. valid values: 10 to 200. default: 10.
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

func (r *ListWorkflowFoldersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListWorkflowFoldersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "ParentFolderPath")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListWorkflowFoldersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListWorkflowFoldersResponseParams struct {
	// Paginated folder query result.
	Data *WorkflowFolderPage `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListWorkflowFoldersResponse struct {
	*tchttp.BaseResponse
	Response *ListWorkflowFoldersResponseParams `json:"Response"`
}

func (r *ListWorkflowFoldersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListWorkflowFoldersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListWorkflowInfo struct {
	// List items.
	Items []*WorkflowInfo `json:"Items,omitnil,omitempty" name:"Items"`

	// Total number of pages that meet the query condition.
	TotalPageNumber *uint64 `json:"TotalPageNumber,omitnil,omitempty" name:"TotalPageNumber"`

	// Current request data page number.
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Number of entries in the current request.
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Total number of data entries that meet the query condition.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`
}

// Predefined struct for user
type ListWorkflowsRequestParams struct {
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Specifies the data page number of the request. default value is 1. valid values: equal to or greater than 1.
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Specifies the number of data records displayed per page. default: 10. value range: 10 to 200.
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Search keywords.
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// Workflow folder.
	ParentFolderPath *string `json:"ParentFolderPath,omitnil,omitempty" name:"ParentFolderPath"`

	// Workflow type. valid values: cycle and manual.
	WorkflowType *string `json:"WorkflowType,omitnil,omitempty" name:"WorkflowType"`

	// bundleId item.
	BundleId *string `json:"BundleId,omitnil,omitempty" name:"BundleId"`

	// Owner ID
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// Creator ID.
	CreateUserUin *string `json:"CreateUserUin,omitnil,omitempty" name:"CreateUserUin"`

	// Modification time interval yyyy-MM-dd HH:MM:ss. fill in two times in the array.
	ModifyTime []*string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`

	// Creation time range yyyy-MM-dd HH:MM:ss. two times must be filled in the array.
	CreateTime []*string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`
}

type ListWorkflowsRequest struct {
	*tchttp.BaseRequest
	
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Specifies the data page number of the request. default value is 1. valid values: equal to or greater than 1.
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Specifies the number of data records displayed per page. default: 10. value range: 10 to 200.
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Search keywords.
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// Workflow folder.
	ParentFolderPath *string `json:"ParentFolderPath,omitnil,omitempty" name:"ParentFolderPath"`

	// Workflow type. valid values: cycle and manual.
	WorkflowType *string `json:"WorkflowType,omitnil,omitempty" name:"WorkflowType"`

	// bundleId item.
	BundleId *string `json:"BundleId,omitnil,omitempty" name:"BundleId"`

	// Owner ID
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// Creator ID.
	CreateUserUin *string `json:"CreateUserUin,omitnil,omitempty" name:"CreateUserUin"`

	// Modification time interval yyyy-MM-dd HH:MM:ss. fill in two times in the array.
	ModifyTime []*string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`

	// Creation time range yyyy-MM-dd HH:MM:ss. two times must be filled in the array.
	CreateTime []*string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`
}

func (r *ListWorkflowsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListWorkflowsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "Keyword")
	delete(f, "ParentFolderPath")
	delete(f, "WorkflowType")
	delete(f, "BundleId")
	delete(f, "OwnerUin")
	delete(f, "CreateUserUin")
	delete(f, "ModifyTime")
	delete(f, "CreateTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListWorkflowsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListWorkflowsResponseParams struct {
	// Describes workflow pagination information.
	Data *ListWorkflowInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListWorkflowsResponse struct {
	*tchttp.BaseResponse
	Response *ListWorkflowsResponseParams `json:"Response"`
}

func (r *ListWorkflowsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListWorkflowsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAlarmRuleResult struct {
	// Whether update succeeded
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`
}

type NotebookSessionInfo struct {
	// Specifies the session ID.
	NotebookSessionId *string `json:"NotebookSessionId,omitnil,omitempty" name:"NotebookSessionId"`

	// Session Name
	NotebookSessionName *string `json:"NotebookSessionName,omitnil,omitempty" name:"NotebookSessionName"`
}

type NotificationFatigue struct {
	// Number of alarms.
	NotifyCount *uint64 `json:"NotifyCount,omitnil,omitempty" name:"NotifyCount"`

	// Alarm interval, in minutes.
	NotifyInterval *uint64 `json:"NotifyInterval,omitnil,omitempty" name:"NotifyInterval"`

	// Do not disturb period, such as example value.
	// [{DaysOfWeek: [1, 2], StartTime: "00:00:00", EndTime: "09:00:00"}]	
	// Specifies notification muting from 00:00 to 09:00 every monday and tuesday.
	QuietIntervals []*AlarmQuietInterval `json:"QuietIntervals,omitnil,omitempty" name:"QuietIntervals"`
}

type OpsAsyncJobDetail struct {
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Operation ID
	AsyncId *string `json:"AsyncId,omitnil,omitempty" name:"AsyncId"`

	// Asynchronous operation type.
	AsyncType *string `json:"AsyncType,omitnil,omitempty" name:"AsyncType"`

	// Asynchronous operation status: initial status: INIT; Running: RUNNING; Success: SUCCESS; failure: FAIL; partially successful: PART_SUCCESS.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Error message.
	// 
	ErrorDesc *string `json:"ErrorDesc,omitnil,omitempty" name:"ErrorDesc"`

	// Total sub-processes.
	TotalSubProcessCount *uint64 `json:"TotalSubProcessCount,omitnil,omitempty" name:"TotalSubProcessCount"`

	// Number of completed sub-processes.
	FinishedSubProcessCount *uint64 `json:"FinishedSubProcessCount,omitnil,omitempty" name:"FinishedSubProcessCount"`

	// Count of successful sub-processes.
	SuccessSubProcessCount *uint64 `json:"SuccessSubProcessCount,omitnil,omitempty" name:"SuccessSubProcessCount"`

	// Operator id.
	CreateUserUin *string `json:"CreateUserUin,omitnil,omitempty" name:"CreateUserUin"`

	// Creation time.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Update time.
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type OpsAsyncResponse struct {
	// Asynchronous execution record Id.
	AsyncId *string `json:"AsyncId,omitnil,omitempty" name:"AsyncId"`
}

type OpsTaskDepend struct {
	// Task ID
	// 
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Task name.
	// 
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// Workflow id.
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// Workflow name.
	// 
	WorkflowName *string `json:"WorkflowName,omitnil,omitempty" name:"WorkflowName"`

	// Project ID.
	// 
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Project name.
	// 
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// Task Status
	// 
	// -N: New
	// 
	// -Y: Scheduling
	// 
	// -F: Offline
	// 
	// -O: Paused
	// 
	// -T: Offlining
	// 
	// -INVALID: Invalid
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Task type Id.
	// * 21:JDBC SQL
	// * 23:TDSQL-PostgreSQL
	// * 26:OfflineSynchronization
	// * 30:Python
	// * 31:PySpark
	// * 33:Impala
	// * 34:Hive SQL
	// * 35:Shell
	// * 36:Spark SQL
	// * 38:Shell Form Mode
	// * 39:Spark
	// * 40:TCHouse-P
	// * 41:Kettle
	// * 42:Tchouse-X
	// * 43:TCHouse-X SQL
	// * 46:DLC Spark
	// * 47:TiOne
	// * 48:Trino
	// * 50:DLC PySpark
	// * 92:MapReduce
	// * 130:Branch Node
	// * 131:Merged Node
	// * 132:Notebook
	// * 133:SSH
	// * 134:StarRocks
	// * 137:For-each
	// * 138:Setats SQL
	// Note: This field may return null, indicating that no valid values can be obtained.
	TaskTypeId *uint64 `json:"TaskTypeId,omitnil,omitempty" name:"TaskTypeId"`

	// Task type description.
	// -20 : universal data synchronization.
	//  - 25 :  ETLTaskType
	//  - 26 :  ETLTaskType
	//  - 30 :  python
	//  - 31 :  pyspark
	//  - 34 :  hivesql
	//  - 35 :  shell
	//  - 36 :  sparksql
	//  - 21 :  jdbcsql
	//  - 32 :  dlc
	//  - 33 :  ImpalaTaskType
	//  - 40 :  CDWTaskType
	//  - 41 :  kettle
	//  - 42 :  TCHouse-X
	//  - 43 :  TCHouse-X SQL
	//  - 46 :  dlcsparkTaskType
	//  - 47 :  TiOneMachineLearningTaskType
	//  - 48 :  Trino
	//  - 50 :  DLCPyspark
	//  - 23 :  TencentDistributedSQL
	//  - 39 :  spark
	//  - 92 :  MRTaskType
	//  - 38 :  ShellScript
	//  - 70 :  HiveSQLScrip
	// -130: specifies the branch.
	// -131: specifies the merge.
	// -132: specifies the Notebook explore.
	// -133: specifies the SSH node.
	//  - 134 :  StarRocks
	//  - 137 :  For-each
	// -10000: common custom business.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TaskTypeDesc *string `json:"TaskTypeDesc,omitnil,omitempty" name:"TaskTypeDesc"`

	// Folder name.
	FolderName *string `json:"FolderName,omitnil,omitempty" name:"FolderName"`

	// Folder ID
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`

	// Latest submission time.
	FirstSubmitTime *string `json:"FirstSubmitTime,omitnil,omitempty" name:"FirstSubmitTime"`

	// First running time
	// 
	FirstRunTime *string `json:"FirstRunTime,omitnil,omitempty" name:"FirstRunTime"`

	// Describes the scheduling plan display description information.
	ScheduleDesc *string `json:"ScheduleDesc,omitnil,omitempty" name:"ScheduleDesc"`

	// Task Cycle Type
	// 
	// * ONEOFF_CYCLE: One-time
	// 
	// * YEAR_CYCLE: Yearly
	// 
	// * MONTH_CYCLE: Monthly
	// 
	// * WEEK_CYCLE: Weekly
	// 
	// * DAY_CYCLE: Daily
	// 
	// * HOUR_CYCLE: Hourly
	// 
	// * MINUTE_CYCLE: Minute-level
	// 
	// * CRONTAB_CYCLE: Crontab expression-based
	CycleType *string `json:"CycleType,omitnil,omitempty" name:"CycleType"`

	// Specifies the person in charge.
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// Execution start time. format: HH:mm, for example 00:00.
	ExecutionStartTime *string `json:"ExecutionStartTime,omitnil,omitempty" name:"ExecutionStartTime"`

	// Execution end time. format: HH:mm, for example 23:59.
	ExecutionEndTime *string `json:"ExecutionEndTime,omitnil,omitempty" name:"ExecutionEndTime"`
}

type OpsWorkflow struct {
	// Number of Tasks
	TaskCount *uint64 `json:"TaskCount,omitnil,omitempty" name:"TaskCount"`

	// folder name.
	FolderName *string `json:"FolderName,omitnil,omitempty" name:"FolderName"`

	// Workflow folder id.
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`

	// Workflow id.
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// Workflow name.
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	WorkflowName *string `json:"WorkflowName,omitnil,omitempty" name:"WorkflowName"`

	// Specifies the workflow type.
	// -cycle period.
	// -manual.
	WorkflowType *string `json:"WorkflowType,omitnil,omitempty" name:"WorkflowType"`

	// Workflow description.
	WorkflowDesc *string `json:"WorkflowDesc,omitnil,omitempty" name:"WorkflowDesc"`

	// Responsible user id, multiple ';' separators.
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Project name.
	// 
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// Workflow Status
	// 
	// * ALL_RUNNING: All running (all workflows are in scheduling state)
	// 
	// * ALL_FREEZED: All paused
	// 
	// * ALL_STOPPTED: All stopped
	// 
	// * PART_RUNNING: Partially running (some workflows are running, others not)
	// 
	// * ALL_NO_RUNNING: All not scheduled
	// 
	// * ALL_INVALID: All invalid
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Workflow creation time.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Latest update time. specifies development and production updates.
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// Last updated by, including development and production updates.
	UpdateUserUin *string `json:"UpdateUserUin,omitnil,omitempty" name:"UpdateUserUin"`
}

type OpsWorkflowDetail struct {
	// Workflow ID.
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// Workflow description.
	WorkflowDesc *string `json:"WorkflowDesc,omitnil,omitempty" name:"WorkflowDesc"`

	// Specifies the workflow type.
	// -cycle.
	// -manual.
	WorkflowType *string `json:"WorkflowType,omitnil,omitempty" name:"WorkflowType"`

	// Creation time.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Creator
	CreateUserUin *string `json:"CreateUserUin,omitnil,omitempty" name:"CreateUserUin"`

	// Modification time.
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// Delayed execution time. unit: minute.
	StartupTime *uint64 `json:"StartupTime,omitnil,omitempty" name:"StartupTime"`

	// Effective date. specifies the start date.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Configure end date end date.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Task Cycle Type
	// 
	// * ONEOFF_CYCLE: One-time
	// 
	// * YEAR_CYCLE: Yearly
	// 
	// * MONTH_CYCLE: Monthly
	// 
	// * WEEK_CYCLE: Weekly
	// 
	// * DAY_CYCLE: Daily
	// 
	// * HOUR_CYCLE: Hourly
	// 
	// * MINUTE_CYCLE: Minute-level
	// 
	// * CRONTAB_CYCLE: Crontab expression-based
	CycleType *string `json:"CycleType,omitnil,omitempty" name:"CycleType"`

	// Folder ID
	// 
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`

	// Task instance initialization strategy.
	// -T_PLUS_1 (t+1): initializes with a one-day delay.
	// -T_PLUS_0 (t+0): initialize on the day.
	// -T_MINUS_1 (t-1): initialize one day in advance.
	InstanceInitStrategy *string `json:"InstanceInitStrategy,omitnil,omitempty" name:"InstanceInitStrategy"`

	// Specifies the scheduling plan interpretation.
	SchedulerDesc *string `json:"SchedulerDesc,omitnil,omitempty" name:"SchedulerDesc"`

	// First submission time of workflow.
	FirstSubmitTime *string `json:"FirstSubmitTime,omitnil,omitempty" name:"FirstSubmitTime"`

	// Latest submission time of workflow.
	LatestSubmitTime *string `json:"LatestSubmitTime,omitnil,omitempty" name:"LatestSubmitTime"`

	// Workflow Status
	// 
	// * ALL_RUNNING: All running (all workflows are in scheduling state)
	// 
	// * ALL_FREEZED: All paused
	// 
	// * ALL_STOPPTED: All stopped
	// 
	// * PART_RUNNING: Partially running (some workflows are running, others not)
	// 
	// * ALL_NO_RUNNING: All not scheduled
	// 
	// * ALL_INVALID: All invalid
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Person in charge. multiple values are separated by a ";" separator.
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// Workflow name.
	WorkflowName *string `json:"WorkflowName,omitnil,omitempty" name:"WorkflowName"`
}

type OpsWorkflows struct {
	// Record list	
	// 	
	Items []*OpsWorkflow `json:"Items,omitnil,omitempty" name:"Items"`

	// Total number of results
	// 
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Total pages
	// 
	TotalPageNumber *uint64 `json:"TotalPageNumber,omitnil,omitempty" name:"TotalPageNumber"`

	// Pagination size.
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Page number
	// 
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`
}

type OutTaskParameter struct {
	// Parameter name.
	ParamKey *string `json:"ParamKey,omitnil,omitempty" name:"ParamKey"`

	// Parameter definition.
	ParamValue *string `json:"ParamValue,omitnil,omitempty" name:"ParamValue"`
}

type ParamInfo struct {
	// Parameter name.
	ParamKey *string `json:"ParamKey,omitnil,omitempty" name:"ParamKey"`

	// Parameter value.
	ParamValue *string `json:"ParamValue,omitnil,omitempty" name:"ParamValue"`
}

type ParentDependencyConfigPage struct {
	// Total number of results
	// 
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Total pages
	// 
	TotalPageNumber *uint64 `json:"TotalPageNumber,omitnil,omitempty" name:"TotalPageNumber"`

	// Page number.
	// 
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Pagination size.
	// 
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Paging data
	Items []*OpsTaskDepend `json:"Items,omitnil,omitempty" name:"Items"`
}

// Predefined struct for user
type PauseOpsTasksAsyncRequestParams struct {
	// Project ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Task Id list. specifies the list of task ids.
	TaskIds []*string `json:"TaskIds,omitnil,omitempty" name:"TaskIds"`

	// Whether required to terminate the generated instance.
	KillInstance *bool `json:"KillInstance,omitnil,omitempty" name:"KillInstance"`
}

type PauseOpsTasksAsyncRequest struct {
	*tchttp.BaseRequest
	
	// Project ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Task Id list. specifies the list of task ids.
	TaskIds []*string `json:"TaskIds,omitnil,omitempty" name:"TaskIds"`

	// Whether required to terminate the generated instance.
	KillInstance *bool `json:"KillInstance,omitnil,omitempty" name:"KillInstance"`
}

func (r *PauseOpsTasksAsyncRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PauseOpsTasksAsyncRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "TaskIds")
	delete(f, "KillInstance")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PauseOpsTasksAsyncRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PauseOpsTasksAsyncResponseParams struct {
	// Asynchronous operation result.
	Data *OpsAsyncResponse `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type PauseOpsTasksAsyncResponse struct {
	*tchttp.BaseResponse
	Response *PauseOpsTasksAsyncResponseParams `json:"Response"`
}

func (r *PauseOpsTasksAsyncResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PauseOpsTasksAsyncResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProjectInstanceStatisticsAlarmInfo struct {
	// Alarm type
	// 
	// projectFailureInstanceUpwardFluctuationAlarm: specifies the upward fluctuation alert for failed instances.
	// 
	// projectSuccessInstanceDownwardFluctuationAlarm: specifies the downward fluctuation alert for successful instances.
	AlarmType *string `json:"AlarmType,omitnil,omitempty" name:"AlarmType"`

	// Alarm threshold for the downward fluctuation ratio of the number of successful instances.
	// Alarm threshold for the upward fluctuation ratio of the number of failed instances.
	InstanceThresholdCountPercent *uint64 `json:"InstanceThresholdCountPercent,omitnil,omitempty" name:"InstanceThresholdCountPercent"`

	// Cumulative instance number fluctuation threshold.
	InstanceThresholdCount *uint64 `json:"InstanceThresholdCount,omitnil,omitempty" name:"InstanceThresholdCount"`

	// Stability threshold count (debounce configuration statistical cycle count).
	StabilizeThreshold *uint64 `json:"StabilizeThreshold,omitnil,omitempty" name:"StabilizeThreshold"`

	// Stability statistical cycle (anti-shake configuration statistical cycle count).
	StabilizeStatisticsCycle *uint64 `json:"StabilizeStatisticsCycle,omitnil,omitempty" name:"StabilizeStatisticsCycle"`

	// Specifies whether to use cumulative calculation. valid values: false (consecutive), true (cumulative).
	IsCumulant *bool `json:"IsCumulant,omitnil,omitempty" name:"IsCumulant"`

	// Cumulative number of instances for the current day.
	// Specifies the downward fluctuation of failed instance count on the day.
	InstanceCount *uint64 `json:"InstanceCount,omitnil,omitempty" name:"InstanceCount"`
}

type ReconciliationStrategyInfo struct {
	// Offline Integration Task Reconciliation Alarms
	// 
	// reconciliationFailure: Alarm when offline reconciliation task fails
	// 
	// reconciliationOvertime: Alarm when offline reconciliation task runs overtime
	// 
	// reconciliationMismatch: Alarm when the number of mismatched records in reconciliation exceeds the threshold
	RuleType *string `json:"RuleType,omitnil,omitempty" name:"RuleType"`

	// Reconciliation Mismatch Threshold - Required when RuleType = reconciliationMismatch. Specifies the threshold for the number of mismatched records in reconciliation. No default value.
	MismatchCount *uint64 `json:"MismatchCount,omitnil,omitempty" name:"MismatchCount"`

	// Task run timeout threshold for reconciliation: hr, defaults to 0.
	Hour *int64 `json:"Hour,omitnil,omitempty" name:"Hour"`

	// Reconciliation task timeout threshold: minutes, defaults to 1.
	Min *int64 `json:"Min,omitnil,omitempty" name:"Min"`
}

// Predefined struct for user
type RerunTaskInstancesAsyncRequestParams struct {
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Instance id list. obtain from ListInstances.
	InstanceKeyList []*string `json:"InstanceKeyList,omitnil,omitempty" name:"InstanceKeyList"`

	// Rerun type. valid values: 1 (self), 3 (child), 2 (self and child). default: 1.
	RerunType *string `json:"RerunType,omitnil,omitempty" name:"RerunType"`

	// Whether to check upstream tasks: ALL (ALL), MAKE_SCOPE (select), NONE (do not check). default is NONE.
	CheckParentType *string `json:"CheckParentType,omitnil,omitempty" name:"CheckParentType"`

	// Downstream Instance Scope
	// 
	// * WORKFLOW: Within the current workflow (default)
	// 
	// * PROJECT: Within the current project
	// 
	// * ALL: Across all projects with cross-workflow dependencies
	SonRangeType *string `json:"SonRangeType,omitnil,omitempty" name:"SonRangeType"`

	// Specifies whether to ignore event monitoring when rerunning.
	SkipEventListening *bool `json:"SkipEventListening,omitnil,omitempty" name:"SkipEventListening"`

	// Specifies the degree of concurrency for a custom instance run. if not configured, use the existing self-dependent task.
	RedefineParallelNum *int64 `json:"RedefineParallelNum,omitnil,omitempty" name:"RedefineParallelNum"`

	// Custom workflow self-dependency. valid values: yes (enable), no (disable). if not configured, use the existing workflow self-dependency.
	RedefineSelfWorkflowDependency *string `json:"RedefineSelfWorkflowDependency,omitnil,omitempty" name:"RedefineSelfWorkflowDependency"`

	// Rerun instance custom parameter.
	RedefineParamList *KVMap `json:"RedefineParamList,omitnil,omitempty" name:"RedefineParamList"`
}

type RerunTaskInstancesAsyncRequest struct {
	*tchttp.BaseRequest
	
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Instance id list. obtain from ListInstances.
	InstanceKeyList []*string `json:"InstanceKeyList,omitnil,omitempty" name:"InstanceKeyList"`

	// Rerun type. valid values: 1 (self), 3 (child), 2 (self and child). default: 1.
	RerunType *string `json:"RerunType,omitnil,omitempty" name:"RerunType"`

	// Whether to check upstream tasks: ALL (ALL), MAKE_SCOPE (select), NONE (do not check). default is NONE.
	CheckParentType *string `json:"CheckParentType,omitnil,omitempty" name:"CheckParentType"`

	// Downstream Instance Scope
	// 
	// * WORKFLOW: Within the current workflow (default)
	// 
	// * PROJECT: Within the current project
	// 
	// * ALL: Across all projects with cross-workflow dependencies
	SonRangeType *string `json:"SonRangeType,omitnil,omitempty" name:"SonRangeType"`

	// Specifies whether to ignore event monitoring when rerunning.
	SkipEventListening *bool `json:"SkipEventListening,omitnil,omitempty" name:"SkipEventListening"`

	// Specifies the degree of concurrency for a custom instance run. if not configured, use the existing self-dependent task.
	RedefineParallelNum *int64 `json:"RedefineParallelNum,omitnil,omitempty" name:"RedefineParallelNum"`

	// Custom workflow self-dependency. valid values: yes (enable), no (disable). if not configured, use the existing workflow self-dependency.
	RedefineSelfWorkflowDependency *string `json:"RedefineSelfWorkflowDependency,omitnil,omitempty" name:"RedefineSelfWorkflowDependency"`

	// Rerun instance custom parameter.
	RedefineParamList *KVMap `json:"RedefineParamList,omitnil,omitempty" name:"RedefineParamList"`
}

func (r *RerunTaskInstancesAsyncRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RerunTaskInstancesAsyncRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "InstanceKeyList")
	delete(f, "RerunType")
	delete(f, "CheckParentType")
	delete(f, "SonRangeType")
	delete(f, "SkipEventListening")
	delete(f, "RedefineParallelNum")
	delete(f, "RedefineSelfWorkflowDependency")
	delete(f, "RedefineParamList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RerunTaskInstancesAsyncRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RerunTaskInstancesAsyncResponseParams struct {
	// Asynchronous ID returned by the batch rerun operation. You can use the GetAsyncJob API to retrieve detailed execution information.
	Data *OpsAsyncResponse `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RerunTaskInstancesAsyncResponse struct {
	*tchttp.BaseResponse
	Response *RerunTaskInstancesAsyncResponseParams `json:"Response"`
}

func (r *RerunTaskInstancesAsyncResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RerunTaskInstancesAsyncResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResourceFile struct {
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Resource file ID.
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// Resource file name.
	ResourceName *string `json:"ResourceName,omitnil,omitempty" name:"ResourceName"`

	// Source file path.
	LocalPath *string `json:"LocalPath,omitnil,omitempty" name:"LocalPath"`

	// Specifies the COS storage path of the resource object.
	RemotePath *string `json:"RemotePath,omitnil,omitempty" name:"RemotePath"`

	// Specifies the resource file type.
	FileExtensionType *string `json:"FileExtensionType,omitnil,omitempty" name:"FileExtensionType"`

	// Resource size
	Size *string `json:"Size,omitnil,omitempty" name:"Size"`

	// Creator user ID 
	CreatorUserUin *string `json:"CreatorUserUin,omitnil,omitempty" name:"CreatorUserUin"`

	// Creator name
	CreatorUserName *string `json:"CreatorUserName,omitnil,omitempty" name:"CreatorUserName"`

	// Last updated username.
	UpdateUserName *string `json:"UpdateUserName,omitnil,omitempty" name:"UpdateUserName"`

	// Last updated user ID.
	UpdateUserUin *string `json:"UpdateUserUin,omitnil,omitempty" name:"UpdateUserUin"`

	// COS bucket.
	BucketName *string `json:"BucketName,omitnil,omitempty" name:"BucketName"`

	// COS region
	CosRegion *string `json:"CosRegion,omitnil,omitempty" name:"CosRegion"`

	// Specifies the resource source mode.
	ResourceSourceMode *string `json:"ResourceSourceMode,omitnil,omitempty" name:"ResourceSourceMode"`

	// Local project ID.
	BundleId *string `json:"BundleId,omitnil,omitempty" name:"BundleId"`

	// Local project information.
	BundleInfo *string `json:"BundleInfo,omitnil,omitempty" name:"BundleInfo"`
}

type ResourceFileItem struct {
	// Resource file ID.
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// Resource file name.
	ResourceName *string `json:"ResourceName,omitnil,omitempty" name:"ResourceName"`

	// Specifies the resource file type.
	FileExtensionType *string `json:"FileExtensionType,omitnil,omitempty" name:"FileExtensionType"`

	// Resource path
	LocalPath *string `json:"LocalPath,omitnil,omitempty" name:"LocalPath"`
}

type ResourceFilePage struct {
	// Task collection information.
	Items []*ResourceFileItem `json:"Items,omitnil,omitempty" name:"Items"`

	// Total page number
	TotalPageNumber *uint64 `json:"TotalPageNumber,omitnil,omitempty" name:"TotalPageNumber"`

	// Total file count.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Current Page number
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Items per Page
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type ResourceFolder struct {
	// Resource folder ID.
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`

	// Creator ID.
	CreateUserUin *string `json:"CreateUserUin,omitnil,omitempty" name:"CreateUserUin"`

	// Creator's name.
	CreateUserName *string `json:"CreateUserName,omitnil,omitempty" name:"CreateUserName"`

	// Specifies the folder path.
	FolderPath *string `json:"FolderPath,omitnil,omitempty" name:"FolderPath"`

	// Folder name.
	FolderName *string `json:"FolderName,omitnil,omitempty" name:"FolderName"`
}

type ResourceFolderPage struct {
	// Resource folder collection information.
	Items []*ResourceFolder `json:"Items,omitnil,omitempty" name:"Items"`

	// Total page number.
	TotalPageNumber *uint64 `json:"TotalPageNumber,omitnil,omitempty" name:"TotalPageNumber"`

	// Total resource folder count
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Current Page number
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Items per Page
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

// Predefined struct for user
type RunSQLScriptRequestParams struct {
	// Script id.
	ScriptId *string `json:"ScriptId,omitnil,omitempty" name:"ScriptId"`

	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Script content. executed by default if not transmitted. requires Base64 encoding if transmitted.
	ScriptContent *string `json:"ScriptContent,omitnil,omitempty" name:"ScriptContent"`

	// Advanced running parameter, JSON format base64 encode.
	Params *string `json:"Params,omitnil,omitempty" name:"Params"`
}

type RunSQLScriptRequest struct {
	*tchttp.BaseRequest
	
	// Script id.
	ScriptId *string `json:"ScriptId,omitnil,omitempty" name:"ScriptId"`

	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Script content. executed by default if not transmitted. requires Base64 encoding if transmitted.
	ScriptContent *string `json:"ScriptContent,omitnil,omitempty" name:"ScriptContent"`

	// Advanced running parameter, JSON format base64 encode.
	Params *string `json:"Params,omitnil,omitempty" name:"Params"`
}

func (r *RunSQLScriptRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunSQLScriptRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ScriptId")
	delete(f, "ProjectId")
	delete(f, "ScriptContent")
	delete(f, "Params")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RunSQLScriptRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RunSQLScriptResponseParams struct {
	// Explores data tasks.
	Data *JobDto `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RunSQLScriptResponse struct {
	*tchttp.BaseResponse
	Response *RunSQLScriptResponseParams `json:"Response"`
}

func (r *RunSQLScriptResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunSQLScriptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SQLContentActionResult struct {
	// Whether the operation is successful
	// 
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`

	// Folder ID.
	// 
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`
}

type SQLFolderNode struct {
	// Unique identifier
	// 
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// Name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// folder type, script type.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Parent folder path, such as /aaa/bbb/ccc. the path must start with a slash. use / for the root directory.
	ParentFolderPath *string `json:"ParentFolderPath,omitnil,omitempty" name:"ParentFolderPath"`

	// Whether it is a leaf node.
	// 
	IsLeaf *bool `json:"IsLeaf,omitnil,omitempty" name:"IsLeaf"`

	// Business parameters	
	Params *string `json:"Params,omitnil,omitempty" name:"Params"`

	// Permission scope: SHARED, PRIVATE.
	AccessScope *string `json:"AccessScope,omitnil,omitempty" name:"AccessScope"`

	// Node path.
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// Creator
	CreateUserUin *string `json:"CreateUserUin,omitnil,omitempty" name:"CreateUserUin"`

	// Specifies the permission of the current user for nodes.	
	NodePermission *string `json:"NodePermission,omitnil,omitempty" name:"NodePermission"`

	// Sub-node list
	Children []*SQLFolderNode `json:"Children,omitnil,omitempty" name:"Children"`

	// Owner of the file.
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`
}

type SQLScript struct {
	// Script id.
	ScriptId *string `json:"ScriptId,omitnil,omitempty" name:"ScriptId"`

	// Script name.
	ScriptName *string `json:"ScriptName,omitnil,omitempty" name:"ScriptName"`

	// Specifies the script owner uin.
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// Parent folder path, /aaa/bbb/ccc.
	ParentFolderPath *string `json:"ParentFolderPath,omitnil,omitempty" name:"ParentFolderPath"`

	// Specifies the script configuration.
	ScriptConfig *SQLScriptConfig `json:"ScriptConfig,omitnil,omitempty" name:"ScriptConfig"`

	// Specifies the script content.
	ScriptContent *string `json:"ScriptContent,omitnil,omitempty" name:"ScriptContent"`

	// Latest operator.
	UpdateUserUin *string `json:"UpdateUserUin,omitnil,omitempty" name:"UpdateUserUin"`

	// Project ID.
	// 
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Update time. format: yyyy-MM-dd hh:MM:ss.
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// Creation time. format: yyyy-MM-dd hh:MM:ss.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Access permission: SHARED, PRIVATE.
	AccessScope *string `json:"AccessScope,omitnil,omitempty" name:"AccessScope"`

	// Full path of the node, /aaa/bbb/ccc.ipynb, consists of the name of each node.
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`
}

type SQLScriptConfig struct {
	// Data source Id.
	DatasourceId *string `json:"DatasourceId,omitnil,omitempty" name:"DatasourceId"`

	// Specifies the data source environment.
	DatasourceEnv *string `json:"DatasourceEnv,omitnil,omitempty" name:"DatasourceEnv"`

	// Computational resource.
	ComputeResource *string `json:"ComputeResource,omitnil,omitempty" name:"ComputeResource"`

	// Specifies the execution resource group.
	ExecutorGroupId *string `json:"ExecutorGroupId,omitnil,omitempty" name:"ExecutorGroupId"`

	// Advanced running parameter variable replacement map-json String,String.
	Params *string `json:"Params,omitnil,omitempty" name:"Params"`

	// Advanced setting. executes configuration parameters. map-json String,String. use Base64 encode.
	AdvanceConfig *string `json:"AdvanceConfig,omitnil,omitempty" name:"AdvanceConfig"`
}

type SQLStopResult struct {
	// Success status
	// 
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`
}

// Predefined struct for user
type SetSuccessTaskInstancesAsyncRequestParams struct {
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Instance id list, can be obtained from ListInstances API.
	InstanceKeyList []*string `json:"InstanceKeyList,omitnil,omitempty" name:"InstanceKeyList"`
}

type SetSuccessTaskInstancesAsyncRequest struct {
	*tchttp.BaseRequest
	
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Instance id list, can be obtained from ListInstances API.
	InstanceKeyList []*string `json:"InstanceKeyList,omitnil,omitempty" name:"InstanceKeyList"`
}

func (r *SetSuccessTaskInstancesAsyncRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetSuccessTaskInstancesAsyncRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "InstanceKeyList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetSuccessTaskInstancesAsyncRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetSuccessTaskInstancesAsyncResponseParams struct {
	// Async id of the batch successful operation. can be used in the GetAsyncJob API to query execution detail.
	Data *OpsAsyncResponse `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SetSuccessTaskInstancesAsyncResponse struct {
	*tchttp.BaseResponse
	Response *SetSuccessTaskInstancesAsyncResponseParams `json:"Response"`
}

func (r *SetSuccessTaskInstancesAsyncResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetSuccessTaskInstancesAsyncResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SqlCreateResult struct {
	// Folder ID
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`
}

// Predefined struct for user
type StopOpsTasksAsyncRequestParams struct {
	// Project ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Task Id list. specifies the list of task ids.
	TaskIds []*string `json:"TaskIds,omitnil,omitempty" name:"TaskIds"`

	// Specifies whether to terminate the generated instance. default value: false.
	KillInstance *bool `json:"KillInstance,omitnil,omitempty" name:"KillInstance"`
}

type StopOpsTasksAsyncRequest struct {
	*tchttp.BaseRequest
	
	// Project ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Task Id list. specifies the list of task ids.
	TaskIds []*string `json:"TaskIds,omitnil,omitempty" name:"TaskIds"`

	// Specifies whether to terminate the generated instance. default value: false.
	KillInstance *bool `json:"KillInstance,omitnil,omitempty" name:"KillInstance"`
}

func (r *StopOpsTasksAsyncRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopOpsTasksAsyncRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "TaskIds")
	delete(f, "KillInstance")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopOpsTasksAsyncRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopOpsTasksAsyncResponseParams struct {
	// AsyncId
	Data *OpsAsyncResponse `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StopOpsTasksAsyncResponse struct {
	*tchttp.BaseResponse
	Response *StopOpsTasksAsyncResponseParams `json:"Response"`
}

func (r *StopOpsTasksAsyncResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopOpsTasksAsyncResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopSQLScriptRunRequestParams struct {
	// Specifies the query id.
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type StopSQLScriptRunRequest struct {
	*tchttp.BaseRequest
	
	// Specifies the query id.
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

func (r *StopSQLScriptRunRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopSQLScriptRunRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopSQLScriptRunRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopSQLScriptRunResponseParams struct {
	// Execution result
	Data *SQLStopResult `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StopSQLScriptRunResponse struct {
	*tchttp.BaseResponse
	Response *StopSQLScriptRunResponseParams `json:"Response"`
}

func (r *StopSQLScriptRunResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopSQLScriptRunResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitTaskRequestParams struct {
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Task ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Version remarks.
	VersionRemark *string `json:"VersionRemark,omitnil,omitempty" name:"VersionRemark"`
}

type SubmitTaskRequest struct {
	*tchttp.BaseRequest
	
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Task ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Version remarks.
	VersionRemark *string `json:"VersionRemark,omitnil,omitempty" name:"VersionRemark"`
}

func (r *SubmitTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "TaskId")
	delete(f, "VersionRemark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SubmitTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitTaskResponseParams struct {
	// Success or failure.
	Data *SubmitTaskResult `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SubmitTaskResponse struct {
	*tchttp.BaseResponse
	Response *SubmitTaskResponseParams `json:"Response"`
}

func (r *SubmitTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SubmitTaskResult struct {
	// Generated task version ID.
	VersionId *string `json:"VersionId,omitnil,omitempty" name:"VersionId"`

	// Submission status.
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`
}

type Task struct {
	// Describes the basic attributes of the task.
	TaskBaseAttribute *TaskBaseAttribute `json:"TaskBaseAttribute,omitnil,omitempty" name:"TaskBaseAttribute"`

	// Task configurations.
	TaskConfiguration *TaskConfiguration `json:"TaskConfiguration,omitnil,omitempty" name:"TaskConfiguration"`

	// Specifies task scheduling configuration.
	TaskSchedulerConfiguration *TaskSchedulerConfiguration `json:"TaskSchedulerConfiguration,omitnil,omitempty" name:"TaskSchedulerConfiguration"`
}

type TaskBaseAttribute struct {
	// Task ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Specifies the task type ID.
	// 
	// * 21:JDBC SQL
	// * 23:TDSQL-PostgreSQL
	// * 26:OfflineSynchronization
	// * 30:Python
	// * 31:PySpark
	// * 33:Impala
	// * 34:Hive SQL
	// * 35:Shell
	// * 36:Spark SQL
	// * 38:Shell Form Mode
	// * 39:Spark
	// * 40:TCHouse-P
	// * 41:Kettle
	// * 42:Tchouse-X
	// * 43:TCHouse-X SQL
	// * 46:DLC Spark
	// * 47:TiOne
	// * 48:Trino
	// * 50:DLC PySpark
	// * 92:MapReduce
	// * 130:Branch Node
	// * 131:Merged Node
	// * 132:Notebook
	// * 133:SSH
	// * 134:StarRocks
	// * 137:For-each
	// * 138:Setats SQL
	TaskTypeId *uint64 `json:"TaskTypeId,omitnil,omitempty" name:"TaskTypeId"`

	// Workflow ID.
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// Task name.
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// Last save version number.
	TaskLatestVersionNo *string `json:"TaskLatestVersionNo,omitnil,omitempty" name:"TaskLatestVersionNo"`

	// Last submit version number.
	TaskLatestSubmitVersionNo *string `json:"TaskLatestSubmitVersionNo,omitnil,omitempty" name:"TaskLatestSubmitVersionNo"`

	// Workflow name.
	// 
	WorkflowName *string `json:"WorkflowName,omitnil,omitempty" name:"WorkflowName"`

	// Task Status:
	// 
	// * N: New
	// * Y: Scheduling
	// * F: Offline
	// * O: Paused
	// * T: Offlining (in the process of being taken offline)
	// * INVALID: Invalid
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Latest submission status of the task. Specifies whether it has been submitted: true/false.
	Submit *bool `json:"Submit,omitnil,omitempty" name:"Submit"`

	// Task creation time. example: 2022-02-12 11:13:41.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Last update time. example: 2025-08-13 16:34:06.
	LastUpdateTime *string `json:"LastUpdateTime,omitnil,omitempty" name:"LastUpdateTime"`

	// Last Updated By (Name).
	LastUpdateUserName *string `json:"LastUpdateUserName,omitnil,omitempty" name:"LastUpdateUserName"`

	// Last operation time.
	LastOpsTime *string `json:"LastOpsTime,omitnil,omitempty" name:"LastOpsTime"`

	// Last operator name.
	LastOpsUserName *string `json:"LastOpsUserName,omitnil,omitempty" name:"LastOpsUserName"`

	// Task owner ID.
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// Task description
	TaskDescription *string `json:"TaskDescription,omitnil,omitempty" name:"TaskDescription"`

	// Last Updated User ID
	UpdateUserUin *string `json:"UpdateUserUin,omitnil,omitempty" name:"UpdateUserUin"`

	// Created By User ID
	CreateUserUin *string `json:"CreateUserUin,omitnil,omitempty" name:"CreateUserUin"`
}

type TaskCode struct {
	// Specifies the code content.
	CodeContent *string `json:"CodeContent,omitnil,omitempty" name:"CodeContent"`

	// Specifies the size of the code file in bytes.
	CodeFileSize *uint64 `json:"CodeFileSize,omitnil,omitempty" name:"CodeFileSize"`
}

type TaskCodeResult struct {
	// Code content.
	CodeInfo *string `json:"CodeInfo,omitnil,omitempty" name:"CodeInfo"`

	// Code file size. unit: KB.
	CodeFileSize *string `json:"CodeFileSize,omitnil,omitempty" name:"CodeFileSize"`
}

type TaskConfiguration struct {
	// Base64 encoding of the code content.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CodeContent *string `json:"CodeContent,omitnil,omitempty" name:"CodeContent"`

	// Extended attribute configuration list of the task.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TaskExtConfigurationList []*TaskExtParameter `json:"TaskExtConfigurationList,omitnil,omitempty" name:"TaskExtConfigurationList"`

	// Cluster ID
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	DataCluster *string `json:"DataCluster,omitnil,omitempty" name:"DataCluster"`

	// Specifies the specified running node.
	// Note: This field may return null, indicating that no valid values can be obtained.
	BrokerIp *string `json:"BrokerIp,omitnil,omitempty" name:"BrokerIp"`

	// Resource pool queue name. need to pass through DescribeProjectClusterQueues to obtain.
	// Note: This field may return null, indicating that no valid values can be obtained.
	YarnQueue *string `json:"YarnQueue,omitnil,omitempty" name:"YarnQueue"`

	// Source data source ID, separated by;, obtained through DescribeDataSourceWithoutInfo.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SourceServiceId *string `json:"SourceServiceId,omitnil,omitempty" name:"SourceServiceId"`

	// Data source type. use semicolon to separate. need to pass through DescribeDataSourceWithoutInfo to obtain.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SourceServiceType *string `json:"SourceServiceType,omitnil,omitempty" name:"SourceServiceType"`

	// Data source name. use semicolons to separate. need to pass through DescribeDataSourceWithoutInfo to obtain.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SourceServiceName *string `json:"SourceServiceName,omitnil,omitempty" name:"SourceServiceName"`

	// TargetTarget data source ID, separated by semicolons. need to pass through DescribeDataSourceWithoutInfo to obtain.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TargetServiceId *string `json:"TargetServiceId,omitnil,omitempty" name:"TargetServiceId"`

	// Target data source type. uses ; for separation. needs to pass through DescribeDataSourceWithoutInfo for retrieval.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TargetServiceType *string `json:"TargetServiceType,omitnil,omitempty" name:"TargetServiceType"`

	// Target data source name. use semicolon to separate. need to pass through DescribeDataSourceWithoutInfo to obtain.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TargetServiceName *string `json:"TargetServiceName,omitnil,omitempty" name:"TargetServiceName"`

	// Resource group ID: need to pass through DescribeNormalSchedulerExecutorGroups to obtain ExecutorGroupId.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ResourceGroup *string `json:"ResourceGroup,omitnil,omitempty" name:"ResourceGroup"`

	// Resource group name: need to pass through DescribeNormalSchedulerExecutorGroups to obtain ExecutorGroupName.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ResourceGroupName *string `json:"ResourceGroupName,omitnil,omitempty" name:"ResourceGroupName"`

	// Specifies the scheduling parameter.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TaskSchedulingParameterList []*TaskSchedulingParameter `json:"TaskSchedulingParameterList,omitnil,omitempty" name:"TaskSchedulingParameterList"`

	// ID used by the Bundle.
	// Note: This field may return null, indicating that no valid values can be obtained.
	BundleId *string `json:"BundleId,omitnil,omitempty" name:"BundleId"`

	// Bundle info.
	// Note: This field may return null, indicating that no valid values can be obtained.
	BundleInfo *string `json:"BundleInfo,omitnil,omitempty" name:"BundleInfo"`
}

type TaskDataRegistry struct {
	// Data source ID.
	DatasourceId *string `json:"DatasourceId,omitnil,omitempty" name:"DatasourceId"`

	// Database name.
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// Table name
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// Partition name
	PartitionName *string `json:"PartitionName,omitnil,omitempty" name:"PartitionName"`

	// Input output table data type.
	// Input stream:
	//  UPSTREAM,
	// Output stream:
	//   DOWNSTREAM;.
	DataFlowType *string `json:"DataFlowType,omitnil,omitempty" name:"DataFlowType"`

	// Physical unique ID..
	TablePhysicalId *string `json:"TablePhysicalId,omitnil,omitempty" name:"TablePhysicalId"`

	// Database unique id..
	DbGuid *string `json:"DbGuid,omitnil,omitempty" name:"DbGuid"`

	// Unique id of the table.
	TableGuid *string `json:"TableGuid,omitnil,omitempty" name:"TableGuid"`
}

type TaskDependDto struct {
	// Task ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Task name.
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// Workflow id.
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// Workflow name.
	WorkflowName *string `json:"WorkflowName,omitnil,omitempty" name:"WorkflowName"`

	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Task Status:
	// 
	// * N: New
	// 
	// * Y: Scheduling
	// 
	// * F: Offline
	// 
	// * O: Paused
	// 
	// * T: Offlining (in the process of being taken offline)
	// 
	// I* NVALID: Invalid
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Task type id.
	TaskTypeId *uint64 `json:"TaskTypeId,omitnil,omitempty" name:"TaskTypeId"`

	// Task type description.
	// -20 : universal data synchronization.
	//  - 25 :  ETLTaskType
	//  - 26 :  ETLTaskType
	//  - 30 :  python
	//  - 31 :  pyspark
	//  - 34 :  hivesql
	//  - 35 :  shell
	//  - 36 :  sparksql
	//  - 21 :  jdbcsql
	//  - 32 :  dlc
	//  - 33 :  ImpalaTaskType
	//  - 40 :  CDWTaskType
	//  - 41 :  kettle
	//  - 42 :  TCHouse-X
	//  - 43 :  TCHouse-X SQL
	//  - 46 :  dlcsparkTaskType
	//  - 47 :  TiOneMachineLearningTaskType
	//  - 48 :  Trino
	//  - 50 :  DLCPyspark
	//  - 23 :  TencentDistributedSQL
	//  - 39 :  spark
	//  - 92 :  MRTaskType
	//  - 38 :  ShellScript
	//  - 70 :  HiveSQLScrip
	// -130: branch.
	// -131: merge.
	// -132: Notebook 
	// -133: SSH node.
	//  - 134 :  StarRocks
	//  - 137 :  For-each
	// -10000 : custom business common.
	TaskTypeDesc *string `json:"TaskTypeDesc,omitnil,omitempty" name:"TaskTypeDesc"`

	// Specifies scheduling plan display description information.
	ScheduleDesc *string `json:"ScheduleDesc,omitnil,omitempty" name:"ScheduleDesc"`

	// Task start time.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Task end time.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Delay time.
	DelayTime *uint64 `json:"DelayTime,omitnil,omitempty" name:"DelayTime"`

	// Cycle Type, Default: D
	// Supported types:
	// * O: One-time
	// 
	// * Y: Yearly
	// 
	// * M: Monthly
	// 
	// * W: Weekly
	// 
	// * D: Daily
	// 
	// * H: Hourly
	// 
	// * I: Minute
	// 
	// * C: Crontab expression type
	CycleType *string `json:"CycleType,omitnil,omitempty" name:"CycleType"`

	// Owner ID
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// Elastic cycle configuration.
	TaskAction *string `json:"TaskAction,omitnil,omitempty" name:"TaskAction"`

	// Initialization strategy for scheduling.
	InitStrategy *string `json:"InitStrategy,omitnil,omitempty" name:"InitStrategy"`

	// crontab expression.
	CrontabExpression *string `json:"CrontabExpression,omitnil,omitempty" name:"CrontabExpression"`
}

type TaskExtParameter struct {
	// Parameter name.
	// 
	ParamKey *string `json:"ParamKey,omitnil,omitempty" name:"ParamKey"`

	// Parameter value.
	// 
	ParamValue *string `json:"ParamValue,omitnil,omitempty" name:"ParamValue"`
}

type TaskInstance struct {
	// Project id.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// **Instance unique id**.
	InstanceKey *string `json:"InstanceKey,omitnil,omitempty" name:"InstanceKey"`

	// Folder ID.
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`

	// Folder name.
	FolderName *string `json:"FolderName,omitnil,omitempty" name:"FolderName"`

	// Workflow ID.
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// Workflow name.
	WorkflowName *string `json:"WorkflowName,omitnil,omitempty" name:"WorkflowName"`

	// Task ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Task name.
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// Instance data time.
	CurRunDate *string `json:"CurRunDate,omitnil,omitempty" name:"CurRunDate"`

	// **Instance status**.
	// -WAIT_EVENT: specifies the wait for event.
	// -WAIT_UPSTREAM: waiting for upstream.
	// - WAIT_RUN: awaiting execution.
	// - RUNNING: indicates the instance is RUNNING.
	// - SKIP_RUNNING: SKIP RUNNING.
	// - FAILED_RETRY: RETRY on failure.
	// - EXPIRED: failed.
	// -COMPLETED: success.
	InstanceState *string `json:"InstanceState,omitnil,omitempty" name:"InstanceState"`

	// Specifies the instance type.
	// 
	// -0 indicates the supplementary entry type.
	// -Indicates a periodic instance.
	// -2 indicates a non-periodic instance.
	InstanceType *uint64 `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// Owner Uin list.
	OwnerUinList []*string `json:"OwnerUinList,omitnil,omitempty" name:"OwnerUinList"`

	// Cumulative running times.
	TotalRunNum *uint64 `json:"TotalRunNum,omitnil,omitempty" name:"TotalRunNum"`

	// Task type description.
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// Task type id.
	TaskTypeId *int64 `json:"TaskTypeId,omitnil,omitempty" name:"TaskTypeId"`

	// Task Cycle Type
	// 
	// * ONEOFF_CYCLE: One-time
	// 
	// * YEAR_CYCLE: Yearly
	// 
	// * MONTH_CYCLE: Monthly
	// 
	// * WEEK_CYCLE: Weekly
	// 
	// * DAY_CYCLE: Daily
	// 
	// * HOUR_CYCLE: Hourly
	// 
	// * MINUTE_CYCLE: Minute-level
	// 
	// * CRONTAB_CYCLE: Crontab expression-based
	CycleType *string `json:"CycleType,omitnil,omitempty" name:"CycleType"`

	// Retry count limit when execution fails each time.
	TryLimit *uint64 `json:"TryLimit,omitnil,omitempty" name:"TryLimit"`

	// Specifies the number of retry attempts on failure.
	// When triggered by manual rerun, supplementary entry instance, or other methods, the count will be reset to 0 and start again.
	Tries *uint64 `json:"Tries,omitnil,omitempty" name:"Tries"`

	// Operation start time.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Operation completion time.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Time spent, in milliseconds.
	CostTime *int64 `json:"CostTime,omitnil,omitempty" name:"CostTime"`

	// Scheduled dispatch time.
	SchedulerTime *string `json:"SchedulerTime,omitnil,omitempty" name:"SchedulerTime"`

	// Latest update time of the instance. specifies the time format as yyyy-MM-dd HH:MM:ss.
	LastUpdateTime *string `json:"LastUpdateTime,omitnil,omitempty" name:"LastUpdateTime"`

	// Execution resource group ID.
	ExecutorGroupId *string `json:"ExecutorGroupId,omitnil,omitempty" name:"ExecutorGroupId"`

	// Resource group name.
	ExecutorGroupName *string `json:"ExecutorGroupName,omitnil,omitempty" name:"ExecutorGroupName"`
}

type TaskInstanceDetail struct {
	// Specifies the project id.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// **Instance unique id**.
	InstanceKey *string `json:"InstanceKey,omitnil,omitempty" name:"InstanceKey"`

	// Folder ID.
	// 
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`

	// Specifies the folder name.
	FolderName *string `json:"FolderName,omitnil,omitempty" name:"FolderName"`

	// Workflow ID.
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// Workflow name.
	// 
	WorkflowName *string `json:"WorkflowName,omitnil,omitempty" name:"WorkflowName"`

	// Task ID
	// 
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Task name.
	// 
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// Specifies the id corresponding to taskType.
	TaskTypeId *int64 `json:"TaskTypeId,omitnil,omitempty" name:"TaskTypeId"`

	// Task type
	// 
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// Task Cycle Type
	// 
	// * ONEOFF_CYCLE: One-time
	// 
	// * YEAR_CYCLE: Yearly
	// 
	// * MONTH_CYCLE: Monthly
	// 
	// * WEEK_CYCLE: Weekly
	// 
	// * DAY_CYCLE: Daily
	// 
	// * HOUR_CYCLE: Hourly
	// 
	// * MINUTE_CYCLE: Minute-level
	// 
	// * CRONTAB_CYCLE: Crontab expression-based
	CycleType *string `json:"CycleType,omitnil,omitempty" name:"CycleType"`

	// Specifies the instance data time.
	CurRunDate *string `json:"CurRunDate,omitnil,omitempty" name:"CurRunDate"`

	// Instance status.
	// -WAIT_EVENT: wait for event.
	// -WAIT_UPSTREAM: waiting for upstream.
	// - WAIT_RUN: awaiting execution.
	// -RUNNING. specifies the running status.
	// - SKIP_RUNNING: specifies whether to SKIP RUNNING.
	// - FAILED_RETRY: RETRY on failure.
	// -EXPIRED: indicates a failure.
	// -COMPLETED: success.
	InstanceState *string `json:"InstanceState,omitnil,omitempty" name:"InstanceState"`

	// Specifies the instance type.
	// 
	// -0 indicates the replenishment type.
	// -Indicates a periodic instance.
	// -2 indicates a non-periodic instance.
	InstanceType *uint64 `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// owner uin list.
	// 
	OwnerUinList []*string `json:"OwnerUinList,omitnil,omitempty" name:"OwnerUinList"`

	// Cumulative running times.
	TotalRunNum *uint64 `json:"TotalRunNum,omitnil,omitempty" name:"TotalRunNum"`

	// Retry count limit when execution fails each time.
	TryLimit *uint64 `json:"TryLimit,omitnil,omitempty" name:"TryLimit"`

	// **Failure Retry Count** - The number of retry attempts after a failure. When the instance is triggered again through methods such as manual rerun or backfill, this counter is reset to 0 and starts counting again.
	Tries *uint64 `json:"Tries,omitnil,omitempty" name:"Tries"`

	// Time spent, in milliseconds.
	CostTime *int64 `json:"CostTime,omitnil,omitempty" name:"CostTime"`

	// Start time.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Operation completion time.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Scheduled dispatch time.
	SchedulerTime *string `json:"SchedulerTime,omitnil,omitempty" name:"SchedulerTime"`

	// Latest update time of the instance. format: yyyy-MM-dd HH:MM:ss.
	LastUpdateTime *string `json:"LastUpdateTime,omitnil,omitempty" name:"LastUpdateTime"`

	// Execution resource group ID.
	ExecutorGroupId *string `json:"ExecutorGroupId,omitnil,omitempty" name:"ExecutorGroupId"`

	// Resource group name.
	ExecutorGroupName *string `json:"ExecutorGroupName,omitnil,omitempty" name:"ExecutorGroupName"`

	// Brief task failure information.
	JobErrorMsg *string `json:"JobErrorMsg,omitnil,omitempty" name:"JobErrorMsg"`
}

type TaskInstanceExecutions struct {
	// Total number of results
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Total pages
	TotalPageNumber *uint64 `json:"TotalPageNumber,omitnil,omitempty" name:"TotalPageNumber"`

	// Record list
	Items []*InstanceExecution `json:"Items,omitnil,omitempty" name:"Items"`

	// Page number.
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Pagination size.
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type TaskInstancePage struct {
	// **Total number of entries**.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// **Total number of pages.**
	TotalPageNumber *uint64 `json:"TotalPageNumber,omitnil,omitempty" name:"TotalPageNumber"`

	// Page number.
	// 
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Specifies the number of entries per page.
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Data List
	// 
	Items []*TaskInstance `json:"Items,omitnil,omitempty" name:"Items"`
}

type TaskOpsInfo struct {
	// Task ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Task name.
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// Owner id
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// Task Status
	// 
	// -N: New
	// 
	// -Y: Scheduling
	// 
	// -F: Offline
	// 
	// -O: Paused
	// 
	// -T: Offlining
	// 
	// -INVALID: Invalid
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Folder ID
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`

	// Folder name.
	FolderName *string `json:"FolderName,omitnil,omitempty" name:"FolderName"`

	// Workflow id.
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// Workflow name.
	WorkflowName *string `json:"WorkflowName,omitnil,omitempty" name:"WorkflowName"`

	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Project name.
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// Specifies the name of the updater.
	UpdateUserUin *string `json:"UpdateUserUin,omitnil,omitempty" name:"UpdateUserUin"`

	// Task type Id.
	// * 21:JDBC SQL
	// * 23:TDSQL-PostgreSQL
	// * 26:OfflineSynchronization
	// * 30:Python
	// * 31:PySpark
	// * 33:Impala
	// * 34:Hive SQL
	// * 35:Shell
	// * 36:Spark SQL
	// * 38:Shell Form Mode
	// * 39:Spark
	// * 40:TCHouse-P
	// * 41:Kettle
	// * 42:Tchouse-X
	// * 43:TCHouse-X SQL
	// * 46:DLC Spark
	// * 47:TiOne
	// * 48:Trino
	// * 50:DLC PySpark
	// * 92:MapReduce
	// * 130:Branch Node
	// * 131:Merged Node
	// * 132:Notebook
	// * 133:SSH
	// * 134:StarRocks
	// * 137:For-each
	// * 138:Setats SQL
	TaskTypeId *uint64 `json:"TaskTypeId,omitnil,omitempty" name:"TaskTypeId"`

	// Task type description.
	// -Universal data synchronization.
	//  - ETLTaskType
	//  - ETLTaskType
	//  - python
	//  - pyspark
	//  - HiveSQLTaskType
	//  - shell
	//  - SparkSQLTaskType
	//  - JDBCSQLTaskType
	//  - DLCTaskType
	//  - ImpalaTaskType
	//  - CDWTaskType
	//  - kettle
	//  - DLCSparkTaskType
	// -TiOne machine learning.
	//  - TrinoTaskType
	//  - DLCPyspark
	//  - spark
	//  - mr
	// -Specifies the shell script.
	// -hivesql script.
	// -Specifies common custom business.
	TaskTypeDesc *string `json:"TaskTypeDesc,omitnil,omitempty" name:"TaskTypeDesc"`

	// Task Cycle Type
	// 
	// * ONEOFF_CYCLE: One-time
	// 
	// * YEAR_CYCLE: Yearly
	// 
	// * MONTH_CYCLE: Monthly
	// 
	// * WEEK_CYCLE: Weekly
	// 
	// * DAY_CYCLE: Daily
	// 
	// * HOUR_CYCLE: Hourly
	// 
	// * MINUTE_CYCLE: Minute-level
	// 
	// * CRONTAB_CYCLE: Crontab expression-based
	CycleType *string `json:"CycleType,omitnil,omitempty" name:"CycleType"`

	// Resource group ID
	ExecutorGroupId *string `json:"ExecutorGroupId,omitnil,omitempty" name:"ExecutorGroupId"`

	// Scheduling description.
	ScheduleDesc *string `json:"ScheduleDesc,omitnil,omitempty" name:"ScheduleDesc"`

	// Resource group name.
	ExecutorGroupName *string `json:"ExecutorGroupName,omitnil,omitempty" name:"ExecutorGroupName"`

	// Latest scheduling submission time.
	LastSchedulerCommitTime *string `json:"LastSchedulerCommitTime,omitnil,omitempty" name:"LastSchedulerCommitTime"`

	// First execution time.
	FirstRunTime *string `json:"FirstRunTime,omitnil,omitempty" name:"FirstRunTime"`

	// Most recent submission time.
	FirstSubmitTime *string `json:"FirstSubmitTime,omitnil,omitempty" name:"FirstSubmitTime"`

	// Creation time.
	// 
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Latest update time.
	LastUpdateTime *string `json:"LastUpdateTime,omitnil,omitempty" name:"LastUpdateTime"`
}

type TaskSchedulerConfiguration struct {
	// Period type. Supported types:
	// 
	// ONEOFF_CYCLE: specifies a one-time cycle.
	// YEAR_CYCLE: specifies the year cycle.
	// MONTH_CYCLE: specifies the monthly cycle.
	// WEEK_CYCLE: specifies the week cycle.
	// DAY_CYCLE: specifies the day cycle.
	// HOUR_CYCLE: specifies the hour cycle.
	// MINUTE_CYCLE: specifies the minute cycle.
	// CRONTAB_CYCLE: specifies the crontab expression type
	CycleType *string `json:"CycleType,omitnil,omitempty" name:"CycleType"`

	// Time zone.
	ScheduleTimeZone *string `json:"ScheduleTimeZone,omitnil,omitempty" name:"ScheduleTimeZone"`

	// 0 2 3 1,L,2 * ?	
	CrontabExpression *string `json:"CrontabExpression,omitnil,omitempty" name:"CrontabExpression"`

	// Effective date.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End date
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Execution time. the left-closed interval.
	ExecutionStartTime *string `json:"ExecutionStartTime,omitnil,omitempty" name:"ExecutionStartTime"`

	// Execution time. right closed interval.
	ExecutionEndTime *string `json:"ExecutionEndTime,omitnil,omitempty" name:"ExecutionEndTime"`

	// Scheduling type: 0 for normal scheduling, 1 for dry-run scheduling.
	ScheduleRunType *int64 `json:"ScheduleRunType,omitnil,omitempty" name:"ScheduleRunType"`

	// Whether calendar scheduling is enabled. Valid values: 1 (enabled), 0 (disabled).
	CalendarOpen *string `json:"CalendarOpen,omitnil,omitempty" name:"CalendarOpen"`

	// Calendar id.
	CalendarId *string `json:"CalendarId,omitnil,omitempty" name:"CalendarId"`

	// Calendar name, which needs to be obtained from DescribeScheduleCalendarPageList API.
	CalendarName *string `json:"CalendarName,omitnil,omitempty" name:"CalendarName"`

	// Self-Dependent. Valid values: parallel, serial, orderly. Default value: serial. 
	SelfDepend *string `json:"SelfDepend,omitnil,omitempty" name:"SelfDepend"`

	// Specifies the upstream dependency array.
	UpstreamDependencyConfigList []*DependencyTaskBrief `json:"UpstreamDependencyConfigList,omitnil,omitempty" name:"UpstreamDependencyConfigList"`

	// SpecSpecifies the downstream dependency array.
	DownStreamDependencyConfigList []*DependencyTaskBrief `json:"DownStreamDependencyConfigList,omitnil,omitempty" name:"DownStreamDependencyConfigList"`

	// Array of Events
	EventListenerList []*EventListener `json:"EventListenerList,omitnil,omitempty" name:"EventListenerList"`

	// Task scheduling priority. valid values: 4 (high), 5 (medium), 6 (low). default: 6.
	RunPriority *uint64 `json:"RunPriority,omitnil,omitempty" name:"RunPriority"`

	// Retry policy. retry wait time in minutes. default: 5.
	RetryWait *int64 `json:"RetryWait,omitnil,omitempty" name:"RetryWait"`

	// Specifies the maximum attempts of the retry policy. default: 4.
	MaxRetryAttempts *int64 `json:"MaxRetryAttempts,omitnil,omitempty" name:"MaxRetryAttempts"`

	// Timeout Handling Policy: Execution Timeout (in minutes), default: -1
	ExecutionTTL *int64 `json:"ExecutionTTL,omitnil,omitempty" name:"ExecutionTTL"`

	// Timeout Handling Policy: Wait Duration Timeout  (in minutes), default: -1
	WaitExecutionTotalTTL *string `json:"WaitExecutionTotalTTL,omitnil,omitempty" name:"WaitExecutionTotalTTL"`

	// Rerun & Refill Configuration: Default: ALL;
	// 
	// * ALL: Rerun or refill is allowed regardless of whether the task succeeds or fails.
	// 
	// * FAILURE: Rerun or refill is allowed only if the task fails; not allowed if the task succeeds.
	// 
	// * NONE: Rerun or refill is not allowed regardless of success or failure.
	AllowRedoType *string `json:"AllowRedoType,omitnil,omitempty" name:"AllowRedoType"`

	// Output parameter list.
	ParamTaskOutList []*OutTaskParameter `json:"ParamTaskOutList,omitnil,omitempty" name:"ParamTaskOutList"`

	// Input parameter list.
	ParamTaskInList []*InTaskParameter `json:"ParamTaskInList,omitnil,omitempty" name:"ParamTaskInList"`

	// Output registration.
	TaskOutputRegistryList []*TaskDataRegistry `json:"TaskOutputRegistryList,omitnil,omitempty" name:"TaskOutputRegistryList"`

	// **Instance generation policy**.
	// T_PLUS_0: specifies t+0 generation. default policy.
	// T_PLUS_1: specifies t+1 generation.
	InitStrategy *string `json:"InitStrategy,omitnil,omitempty" name:"InitStrategy"`
}

type TaskSchedulingParameter struct {
	// Parameter name.
	ParamKey *string `json:"ParamKey,omitnil,omitempty" name:"ParamKey"`

	// Parameter value.
	ParamValue *string `json:"ParamValue,omitnil,omitempty" name:"ParamValue"`
}

type TaskVersion struct {
	// Creation time
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Version number
	VersionNum *string `json:"VersionNum,omitnil,omitempty" name:"VersionNum"`

	// Creator ID
	CreateUserUin *string `json:"CreateUserUin,omitnil,omitempty" name:"CreateUserUin"`

	// Saved version ID
	VersionId *string `json:"VersionId,omitnil,omitempty" name:"VersionId"`

	// Version description
	VersionRemark *string `json:"VersionRemark,omitnil,omitempty" name:"VersionRemark"`

	// Approval status (only for submit version).
	ApproveStatus *string `json:"ApproveStatus,omitnil,omitempty" name:"ApproveStatus"`

	// Production status (only for submit version).
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Approver (only for submit version).
	ApproveUserUin *string `json:"ApproveUserUin,omitnil,omitempty" name:"ApproveUserUin"`
}

type TaskVersionDetail struct {
	// Creation time.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Version number
	// 
	VersionNum *string `json:"VersionNum,omitnil,omitempty" name:"VersionNum"`

	// Specifies the version creator.
	CreateUserUin *string `json:"CreateUserUin,omitnil,omitempty" name:"CreateUserUin"`

	// Specifies the version Id to save.
	VersionId *string `json:"VersionId,omitnil,omitempty" name:"VersionId"`

	// Version description
	VersionRemark *string `json:"VersionRemark,omitnil,omitempty" name:"VersionRemark"`

	// Approval status (only for submit version).
	ApproveStatus *string `json:"ApproveStatus,omitnil,omitempty" name:"ApproveStatus"`

	// Production status  (only for submit version).
	ApproveTime *string `json:"ApproveTime,omitnil,omitempty" name:"ApproveTime"`

	// Describes the task detail of the version.
	Task *Task `json:"Task,omitnil,omitempty" name:"Task"`

	// Approver Id.
	ApproveUserUin *string `json:"ApproveUserUin,omitnil,omitempty" name:"ApproveUserUin"`
}

type TimeOutStrategyInfo struct {
	// Timeout Alarm Configuration
	// 
	// Expected Execution Duration Timeout - The actual runtime exceeds the estimated execution duration.
	// 
	// Expected Completion Time Timeout - The task has not completed by the estimated completion time.
	// 
	// Expected Scheduling Wait Timeout - The waiting time in the scheduling queue exceeds the estimated wait time.
	// 
	// Cycle-Incomplete Timeout - The task was expected to complete within its scheduled cycle but did not.
	RuleType *int64 `json:"RuleType,omitnil,omitempty" name:"RuleType"`

	// Timeout Value Configuration Type
	// 
	// 1: Fixed value (specified manually)
	// 
	// 2: Average value (calculated automatically)
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// Timeout Specified Value (hours) - The timeout threshold in hours. Default is 1.
	Hour *uint64 `json:"Hour,omitnil,omitempty" name:"Hour"`

	// Timeout Specified Value (minutes) - The timeout threshold in minutes. Default is 1.
	Min *int64 `json:"Min,omitnil,omitempty" name:"Min"`

	// The time zone configuration corresponding to the timeout, such as UTC+7, defaults to UTC+8.
	ScheduleTimeZone *string `json:"ScheduleTimeZone,omitnil,omitempty" name:"ScheduleTimeZone"`
}

// Predefined struct for user
type UpdateCodeFileRequestParams struct {
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Code file ID. the parameter value comes from the CreateCodeFile API response.
	CodeFileId *string `json:"CodeFileId,omitnil,omitempty" name:"CodeFileId"`

	// Specifies the code file configuration.
	CodeFileConfig *CodeFileConfig `json:"CodeFileConfig,omitnil,omitempty" name:"CodeFileConfig"`

	// Specifies the content of the code file.
	CodeFileContent *string `json:"CodeFileContent,omitnil,omitempty" name:"CodeFileContent"`
}

type UpdateCodeFileRequest struct {
	*tchttp.BaseRequest
	
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Code file ID. the parameter value comes from the CreateCodeFile API response.
	CodeFileId *string `json:"CodeFileId,omitnil,omitempty" name:"CodeFileId"`

	// Specifies the code file configuration.
	CodeFileConfig *CodeFileConfig `json:"CodeFileConfig,omitnil,omitempty" name:"CodeFileConfig"`

	// Specifies the content of the code file.
	CodeFileContent *string `json:"CodeFileContent,omitnil,omitempty" name:"CodeFileContent"`
}

func (r *UpdateCodeFileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateCodeFileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "CodeFileId")
	delete(f, "CodeFileConfig")
	delete(f, "CodeFileContent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateCodeFileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateCodeFileResponseParams struct {
	// Result
	Data *CodeFile `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateCodeFileResponse struct {
	*tchttp.BaseResponse
	Response *UpdateCodeFileResponseParams `json:"Response"`
}

func (r *UpdateCodeFileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateCodeFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateCodeFolderRequestParams struct {
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Folder ID. The parameter value can be obtained from the response of the CreateCodeFolder API.
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`

	// Folder name.
	FolderName *string `json:"FolderName,omitnil,omitempty" name:"FolderName"`
}

type UpdateCodeFolderRequest struct {
	*tchttp.BaseRequest
	
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Folder ID. The parameter value can be obtained from the response of the CreateCodeFolder API.
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`

	// Folder name.
	FolderName *string `json:"FolderName,omitnil,omitempty" name:"FolderName"`
}

func (r *UpdateCodeFolderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateCodeFolderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "FolderId")
	delete(f, "FolderName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateCodeFolderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateCodeFolderResponseParams struct {
	// Execution result
	Data *CodeStudioFolderActionResult `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateCodeFolderResponse struct {
	*tchttp.BaseResponse
	Response *UpdateCodeFolderResponseParams `json:"Response"`
}

func (r *UpdateCodeFolderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateCodeFolderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateFolderResult struct {
	// Update status. true indicates update succeeded. false indicates update failed.
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`
}

// Predefined struct for user
type UpdateOpsAlarmRuleRequestParams struct {
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Specifies the unique id of the Alarm rule, obtained through the GetAlarmRule/ListAlarmRule api.
	AlarmRuleId *string `json:"AlarmRuleId,omitnil,omitempty" name:"AlarmRuleId"`

	// Specifies the new rule name of the Alarm rule.
	AlarmRuleName *string `json:"AlarmRuleName,omitnil,omitempty" name:"AlarmRuleName"`

	// Monitoring Object Type
	// 
	// Task-level Monitoring - Can be configured by Task / Workflow / Project:
	// 1 = Task (default)
	// 2 = Workflow
	// 3 = Project
	// 
	// Project-level Monitoring - Monitors overall task fluctuations within a project:
	// 7 = Project fluctuation monitoring alarm
	MonitorObjectType *int64 `json:"MonitorObjectType,omitnil,omitempty" name:"MonitorObjectType"`

	// Pass different business IDs based on the MonitorType setting:
	// 
	// 1 (Task): MonitorObjectIds should be a list of task IDs.
	// 
	// 2 (Workflow): MonitorObjectIds should be a list of workflow IDs (workflow IDs can be obtained via the ListWorkflows API).
	// 
	// 3 (Project): MonitorObjectIds should be a list of project IDs.
	MonitorObjectIds []*string `json:"MonitorObjectIds,omitnil,omitempty" name:"MonitorObjectIds"`

	// Alarm Rule Monitoring Types
	// 
	// failure: Failure alarm
	// 
	// overtime: Timeout alarm
	// 
	// success: Success alarm
	// 
	// backTrackingOrRerunSuccess: Alarm when backfill/rerun succeeds
	// 
	// backTrackingOrRerunFailure: Alarm when backfill/rerun fails
	// 
	// projectFailureInstanceUpwardFluctuationAlarm: Alarm when the upward fluctuation rate of failed instances for the day exceeds the threshold
	// 
	// projectFailureInstanceUpwardVolatilityAlarm: Alarm when the upward fluctuation count of failed instances for the day exceeds the threshold
	// 
	// projectSuccessInstanceDownwardFluctuationAlarm: Alarm when the downward fluctuation rate of successful instances for the day exceeds the threshold
	// 
	// projectSuccessInstanceDownwardVolatilityAlarm: Alarm when the downward fluctuation count of successful instances for the day exceeds the threshold
	// 
	// projectFailureInstanceCountAlarm: Alarm when the number of failed instances for the day exceeds the threshold
	// 
	// projectFailureInstanceProportionAlarm: Alarm when the proportion of failed instances for the day exceeds the threshold
	// 
	// reconciliationFailure: Alarm when offline reconciliation task fails
	// 
	// reconciliationOvertime: Alarm when offline reconciliation task runs overtime
	// 
	// reconciliationMismatch: Alarm when the number of mismatched records in reconciliation exceeds the threshold
	AlarmTypes []*string `json:"AlarmTypes,omitnil,omitempty" name:"AlarmTypes"`

	// Alarm Rule Configuration Information
	// 
	// Success Alarms - No configuration required;
	// 
	// Failure Alarms - Can be configured to trigger on the first failure or on all retry failures;
	// 
	// Timeout Alarms - Require configuration of the timeout type and timeout threshold;
	// 
	// Project Fluctuation Alarms - Require configuration of the fluctuation rate and the debounce cycle.
	AlarmRuleDetail *AlarmRuleDetail `json:"AlarmRuleDetail,omitnil,omitempty" name:"AlarmRuleDetail"`

	// Enable status of the Alarm rule. valid values: 0 (disabled), 1 (enabled).
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Alarm level. 1. ordinary, 2. important, 3. critical.
	AlarmLevel *int64 `json:"AlarmLevel,omitnil,omitempty" name:"AlarmLevel"`

	// Describes the Alarm recipient configuration message.
	AlarmGroups []*AlarmGroup `json:"AlarmGroups,omitnil,omitempty" name:"AlarmGroups"`

	// Alarm description.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type UpdateOpsAlarmRuleRequest struct {
	*tchttp.BaseRequest
	
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Specifies the unique id of the Alarm rule, obtained through the GetAlarmRule/ListAlarmRule api.
	AlarmRuleId *string `json:"AlarmRuleId,omitnil,omitempty" name:"AlarmRuleId"`

	// Specifies the new rule name of the Alarm rule.
	AlarmRuleName *string `json:"AlarmRuleName,omitnil,omitempty" name:"AlarmRuleName"`

	// Monitoring Object Type
	// 
	// Task-level Monitoring - Can be configured by Task / Workflow / Project:
	// 1 = Task (default)
	// 2 = Workflow
	// 3 = Project
	// 
	// Project-level Monitoring - Monitors overall task fluctuations within a project:
	// 7 = Project fluctuation monitoring alarm
	MonitorObjectType *int64 `json:"MonitorObjectType,omitnil,omitempty" name:"MonitorObjectType"`

	// Pass different business IDs based on the MonitorType setting:
	// 
	// 1 (Task): MonitorObjectIds should be a list of task IDs.
	// 
	// 2 (Workflow): MonitorObjectIds should be a list of workflow IDs (workflow IDs can be obtained via the ListWorkflows API).
	// 
	// 3 (Project): MonitorObjectIds should be a list of project IDs.
	MonitorObjectIds []*string `json:"MonitorObjectIds,omitnil,omitempty" name:"MonitorObjectIds"`

	// Alarm Rule Monitoring Types
	// 
	// failure: Failure alarm
	// 
	// overtime: Timeout alarm
	// 
	// success: Success alarm
	// 
	// backTrackingOrRerunSuccess: Alarm when backfill/rerun succeeds
	// 
	// backTrackingOrRerunFailure: Alarm when backfill/rerun fails
	// 
	// projectFailureInstanceUpwardFluctuationAlarm: Alarm when the upward fluctuation rate of failed instances for the day exceeds the threshold
	// 
	// projectFailureInstanceUpwardVolatilityAlarm: Alarm when the upward fluctuation count of failed instances for the day exceeds the threshold
	// 
	// projectSuccessInstanceDownwardFluctuationAlarm: Alarm when the downward fluctuation rate of successful instances for the day exceeds the threshold
	// 
	// projectSuccessInstanceDownwardVolatilityAlarm: Alarm when the downward fluctuation count of successful instances for the day exceeds the threshold
	// 
	// projectFailureInstanceCountAlarm: Alarm when the number of failed instances for the day exceeds the threshold
	// 
	// projectFailureInstanceProportionAlarm: Alarm when the proportion of failed instances for the day exceeds the threshold
	// 
	// reconciliationFailure: Alarm when offline reconciliation task fails
	// 
	// reconciliationOvertime: Alarm when offline reconciliation task runs overtime
	// 
	// reconciliationMismatch: Alarm when the number of mismatched records in reconciliation exceeds the threshold
	AlarmTypes []*string `json:"AlarmTypes,omitnil,omitempty" name:"AlarmTypes"`

	// Alarm Rule Configuration Information
	// 
	// Success Alarms - No configuration required;
	// 
	// Failure Alarms - Can be configured to trigger on the first failure or on all retry failures;
	// 
	// Timeout Alarms - Require configuration of the timeout type and timeout threshold;
	// 
	// Project Fluctuation Alarms - Require configuration of the fluctuation rate and the debounce cycle.
	AlarmRuleDetail *AlarmRuleDetail `json:"AlarmRuleDetail,omitnil,omitempty" name:"AlarmRuleDetail"`

	// Enable status of the Alarm rule. valid values: 0 (disabled), 1 (enabled).
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Alarm level. 1. ordinary, 2. important, 3. critical.
	AlarmLevel *int64 `json:"AlarmLevel,omitnil,omitempty" name:"AlarmLevel"`

	// Describes the Alarm recipient configuration message.
	AlarmGroups []*AlarmGroup `json:"AlarmGroups,omitnil,omitempty" name:"AlarmGroups"`

	// Alarm description.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *UpdateOpsAlarmRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateOpsAlarmRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "AlarmRuleId")
	delete(f, "AlarmRuleName")
	delete(f, "MonitorObjectType")
	delete(f, "MonitorObjectIds")
	delete(f, "AlarmTypes")
	delete(f, "AlarmRuleDetail")
	delete(f, "Status")
	delete(f, "AlarmLevel")
	delete(f, "AlarmGroups")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateOpsAlarmRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateOpsAlarmRuleResponseParams struct {
	// Indicates whether the update is successful.
	// true: update successful. false: failed to update.
	Data *ModifyAlarmRuleResult `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateOpsAlarmRuleResponse struct {
	*tchttp.BaseResponse
	Response *UpdateOpsAlarmRuleResponseParams `json:"Response"`
}

func (r *UpdateOpsAlarmRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateOpsAlarmRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateOpsTasksOwnerRequestParams struct {
	// Project ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Task Id list. specifies the list of task ids.
	TaskIds []*string `json:"TaskIds,omitnil,omitempty" name:"TaskIds"`

	// Task owner Id.
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`
}

type UpdateOpsTasksOwnerRequest struct {
	*tchttp.BaseRequest
	
	// Project ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Task Id list. specifies the list of task ids.
	TaskIds []*string `json:"TaskIds,omitnil,omitempty" name:"TaskIds"`

	// Task owner Id.
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`
}

func (r *UpdateOpsTasksOwnerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateOpsTasksOwnerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "TaskIds")
	delete(f, "OwnerUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateOpsTasksOwnerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateOpsTasksOwnerResponseParams struct {
	// Operation result.
	Data *UpdateTasksOwner `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateOpsTasksOwnerResponse struct {
	*tchttp.BaseResponse
	Response *UpdateOpsTasksOwnerResponseParams `json:"Response"`
}

func (r *UpdateOpsTasksOwnerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateOpsTasksOwnerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateResourceFileRequestParams struct {
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Resource file ID. Can be obtained through the ListResourceFiles API.
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// -Upload file and manual entry are two methods, choose one. if both are provided, the sequence is file > manual entry.
	// -The hand-filled value must be an existing cos path. /datastudio/resource/ is the fixed prefix. projectId is the project ID and requires a specific value. parentFolderPath is the parent folder path. name is the file name. 
	// Hand-filled value example:.
	//    /datastudio/resource/projectId/parentFolderPath/name 
	ResourceFile *string `json:"ResourceFile,omitnil,omitempty" name:"ResourceFile"`

	// Resource name, preferably kept consistent with the file name.
	ResourceName *string `json:"ResourceName,omitnil,omitempty" name:"ResourceName"`

	// Bundle Client ID.
	BundleId *string `json:"BundleId,omitnil,omitempty" name:"BundleId"`

	// Bundle Client Name
	BundleInfo *string `json:"BundleInfo,omitnil,omitempty" name:"BundleInfo"`
}

type UpdateResourceFileRequest struct {
	*tchttp.BaseRequest
	
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Resource file ID. Can be obtained through the ListResourceFiles API.
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// -Upload file and manual entry are two methods, choose one. if both are provided, the sequence is file > manual entry.
	// -The hand-filled value must be an existing cos path. /datastudio/resource/ is the fixed prefix. projectId is the project ID and requires a specific value. parentFolderPath is the parent folder path. name is the file name. 
	// Hand-filled value example:.
	//    /datastudio/resource/projectId/parentFolderPath/name 
	ResourceFile *string `json:"ResourceFile,omitnil,omitempty" name:"ResourceFile"`

	// Resource name, preferably kept consistent with the file name.
	ResourceName *string `json:"ResourceName,omitnil,omitempty" name:"ResourceName"`

	// Bundle Client ID.
	BundleId *string `json:"BundleId,omitnil,omitempty" name:"BundleId"`

	// Bundle Client Name
	BundleInfo *string `json:"BundleInfo,omitnil,omitempty" name:"BundleInfo"`
}

func (r *UpdateResourceFileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateResourceFileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "ResourceId")
	delete(f, "ResourceFile")
	delete(f, "ResourceName")
	delete(f, "BundleId")
	delete(f, "BundleInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateResourceFileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateResourceFileResponseParams struct {
	// Update status.
	Data *UpdateResourceFileResult `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateResourceFileResponse struct {
	*tchttp.BaseResponse
	Response *UpdateResourceFileResponseParams `json:"Response"`
}

func (r *UpdateResourceFileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateResourceFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateResourceFileResult struct {
	// true
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`
}

// Predefined struct for user
type UpdateResourceFolderRequestParams struct {
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Folder ID. obtain through the  ListResourceFolders API.
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`

	// Folder name.
	FolderName *string `json:"FolderName,omitnil,omitempty" name:"FolderName"`
}

type UpdateResourceFolderRequest struct {
	*tchttp.BaseRequest
	
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Folder ID. obtain through the  ListResourceFolders API.
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`

	// Folder name.
	FolderName *string `json:"FolderName,omitnil,omitempty" name:"FolderName"`
}

func (r *UpdateResourceFolderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateResourceFolderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "FolderId")
	delete(f, "FolderName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateResourceFolderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateResourceFolderResponseParams struct {
	// Specifies the update folder result. if the update fails, an error will be reported.
	Data *UpdateFolderResult `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateResourceFolderResponse struct {
	*tchttp.BaseResponse
	Response *UpdateResourceFolderResponseParams `json:"Response"`
}

func (r *UpdateResourceFolderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateResourceFolderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateSQLFolderRequestParams struct {
	// Folder ID
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`

	// Folder name.
	FolderName *string `json:"FolderName,omitnil,omitempty" name:"FolderName"`

	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Access permission: SHARED, PRIVATE.
	AccessScope *string `json:"AccessScope,omitnil,omitempty" name:"AccessScope"`
}

type UpdateSQLFolderRequest struct {
	*tchttp.BaseRequest
	
	// Folder ID
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`

	// Folder name.
	FolderName *string `json:"FolderName,omitnil,omitempty" name:"FolderName"`

	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Access permission: SHARED, PRIVATE.
	AccessScope *string `json:"AccessScope,omitnil,omitempty" name:"AccessScope"`
}

func (r *UpdateSQLFolderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateSQLFolderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FolderId")
	delete(f, "FolderName")
	delete(f, "ProjectId")
	delete(f, "AccessScope")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateSQLFolderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateSQLFolderResponseParams struct {
	// Indicates whether the operation is successful. valid values: true (successful), false (unsuccessful).
	Data *SQLContentActionResult `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateSQLFolderResponse struct {
	*tchttp.BaseResponse
	Response *UpdateSQLFolderResponseParams `json:"Response"`
}

func (r *UpdateSQLFolderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateSQLFolderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateSQLScriptRequestParams struct {
	// Script Id.
	ScriptId *string `json:"ScriptId,omitnil,omitempty" name:"ScriptId"`

	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Specifies the script configuration for data exploration.
	ScriptConfig *SQLScriptConfig `json:"ScriptConfig,omitnil,omitempty" name:"ScriptConfig"`

	// Specifies the script content, needs to be Base64 encoded.
	ScriptContent *string `json:"ScriptContent,omitnil,omitempty" name:"ScriptContent"`
}

type UpdateSQLScriptRequest struct {
	*tchttp.BaseRequest
	
	// Script Id.
	ScriptId *string `json:"ScriptId,omitnil,omitempty" name:"ScriptId"`

	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Specifies the script configuration for data exploration.
	ScriptConfig *SQLScriptConfig `json:"ScriptConfig,omitnil,omitempty" name:"ScriptConfig"`

	// Specifies the script content, needs to be Base64 encoded.
	ScriptContent *string `json:"ScriptContent,omitnil,omitempty" name:"ScriptContent"`
}

func (r *UpdateSQLScriptRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateSQLScriptRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ScriptId")
	delete(f, "ProjectId")
	delete(f, "ScriptConfig")
	delete(f, "ScriptContent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateSQLScriptRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateSQLScriptResponseParams struct {
	// Result
	// 
	Data *SQLScript `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateSQLScriptResponse struct {
	*tchttp.BaseResponse
	Response *UpdateSQLScriptResponseParams `json:"Response"`
}

func (r *UpdateSQLScriptResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateSQLScriptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateTaskBaseAttribute struct {
	// Task name.
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// Task owner ID.
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// Task description
	TaskDescription *string `json:"TaskDescription,omitnil,omitempty" name:"TaskDescription"`
}

type UpdateTaskBrief struct {
	// Describes the basic attributes of the task.
	TaskBaseAttribute *UpdateTaskBaseAttribute `json:"TaskBaseAttribute,omitnil,omitempty" name:"TaskBaseAttribute"`

	// Task configurations.
	TaskConfiguration *TaskConfiguration `json:"TaskConfiguration,omitnil,omitempty" name:"TaskConfiguration"`

	// Task scheduling configuration.
	TaskSchedulerConfiguration *TaskSchedulerConfiguration `json:"TaskSchedulerConfiguration,omitnil,omitempty" name:"TaskSchedulerConfiguration"`
}

// Predefined struct for user
type UpdateTaskRequestParams struct {
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Task ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Describes the basic attributes of the task.
	Task *UpdateTaskBrief `json:"Task,omitnil,omitempty" name:"Task"`
}

type UpdateTaskRequest struct {
	*tchttp.BaseRequest
	
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Task ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Describes the basic attributes of the task.
	Task *UpdateTaskBrief `json:"Task,omitnil,omitempty" name:"Task"`
}

func (r *UpdateTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "TaskId")
	delete(f, "Task")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateTaskResponseParams struct {
	// Task ID
	Data *UpdateTaskResult `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateTaskResponse struct {
	*tchttp.BaseResponse
	Response *UpdateTaskResponseParams `json:"Response"`
}

func (r *UpdateTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateTaskResult struct {
	// Processing result. returns true on success. returns false on failure.
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`
}

type UpdateTasksOwner struct {
	// Describes the result of modifying the task owner.
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`
}

// Predefined struct for user
type UpdateWorkflowFolderRequestParams struct {
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Folder ID. obtain through the ListWorkflowFolders API.
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`

	// Specifies the folder name after update.
	FolderName *string `json:"FolderName,omitnil,omitempty" name:"FolderName"`
}

type UpdateWorkflowFolderRequest struct {
	*tchttp.BaseRequest
	
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Folder ID. obtain through the ListWorkflowFolders API.
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`

	// Specifies the folder name after update.
	FolderName *string `json:"FolderName,omitnil,omitempty" name:"FolderName"`
}

func (r *UpdateWorkflowFolderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateWorkflowFolderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "FolderId")
	delete(f, "FolderName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateWorkflowFolderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateWorkflowFolderResponseParams struct {
	// Specifies the update folder result. if the update fails, an error will be reported.
	Data *UpdateFolderResult `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateWorkflowFolderResponse struct {
	*tchttp.BaseResponse
	Response *UpdateWorkflowFolderResponseParams `json:"Response"`
}

func (r *UpdateWorkflowFolderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateWorkflowFolderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateWorkflowRequestParams struct {
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Workflow ID.
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// Workflow name.
	WorkflowName *string `json:"WorkflowName,omitnil,omitempty" name:"WorkflowName"`

	// Owner ID.
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// Remarks.
	WorkflowDesc *string `json:"WorkflowDesc,omitnil,omitempty" name:"WorkflowDesc"`

	// Workflow parameter list.
	WorkflowParams []*ParamInfo `json:"WorkflowParams,omitnil,omitempty" name:"WorkflowParams"`

	// Specifies unified scheduling parameters.
	WorkflowSchedulerConfiguration *WorkflowSchedulerConfigurationInfo `json:"WorkflowSchedulerConfiguration,omitnil,omitempty" name:"WorkflowSchedulerConfiguration"`

	// BundleId item.
	BundleId *string `json:"BundleId,omitnil,omitempty" name:"BundleId"`

	// Bundle info.
	BundleInfo *string `json:"BundleInfo,omitnil,omitempty" name:"BundleInfo"`
}

type UpdateWorkflowRequest struct {
	*tchttp.BaseRequest
	
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Workflow ID.
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// Workflow name.
	WorkflowName *string `json:"WorkflowName,omitnil,omitempty" name:"WorkflowName"`

	// Owner ID.
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// Remarks.
	WorkflowDesc *string `json:"WorkflowDesc,omitnil,omitempty" name:"WorkflowDesc"`

	// Workflow parameter list.
	WorkflowParams []*ParamInfo `json:"WorkflowParams,omitnil,omitempty" name:"WorkflowParams"`

	// Specifies unified scheduling parameters.
	WorkflowSchedulerConfiguration *WorkflowSchedulerConfigurationInfo `json:"WorkflowSchedulerConfiguration,omitnil,omitempty" name:"WorkflowSchedulerConfiguration"`

	// BundleId item.
	BundleId *string `json:"BundleId,omitnil,omitempty" name:"BundleId"`

	// Bundle info.
	BundleInfo *string `json:"BundleInfo,omitnil,omitempty" name:"BundleInfo"`
}

func (r *UpdateWorkflowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateWorkflowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "WorkflowId")
	delete(f, "WorkflowName")
	delete(f, "OwnerUin")
	delete(f, "WorkflowDesc")
	delete(f, "WorkflowParams")
	delete(f, "WorkflowSchedulerConfiguration")
	delete(f, "BundleId")
	delete(f, "BundleInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateWorkflowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateWorkflowResponseParams struct {
	// true represents success. false represents failure.
	Data *UpdateWorkflowResult `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateWorkflowResponse struct {
	*tchttp.BaseResponse
	Response *UpdateWorkflowResponseParams `json:"Response"`
}

func (r *UpdateWorkflowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateWorkflowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateWorkflowResult struct {
	// Update Workflow Result
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`
}

type WorkflowDetail struct {
	// Workflow name.
	WorkflowName *string `json:"WorkflowName,omitnil,omitempty" name:"WorkflowName"`

	// Owner ID.
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// Creator ID.
	CreateUserUin *string `json:"CreateUserUin,omitnil,omitempty" name:"CreateUserUin"`

	// Workflow type. Valid values: cycle or manual.
	WorkflowType *string `json:"WorkflowType,omitnil,omitempty" name:"WorkflowType"`

	// Workflow parameter array.
	WorkflowParams []*ParamInfo `json:"WorkflowParams,omitnil,omitempty" name:"WorkflowParams"`

	// Unified scheduling parameter.
	// .
	WorkflowSchedulerConfiguration *WorkflowSchedulerConfiguration `json:"WorkflowSchedulerConfiguration,omitnil,omitempty" name:"WorkflowSchedulerConfiguration"`

	// Workflow description.
	WorkflowDesc *string `json:"WorkflowDesc,omitnil,omitempty" name:"WorkflowDesc"`

	// Workflow path.
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// BundleId item.
	BundleId *string `json:"BundleId,omitnil,omitempty" name:"BundleId"`

	// BundleInfo item. specifies the bundle information.
	BundleInfo *string `json:"BundleInfo,omitnil,omitempty" name:"BundleInfo"`
}

type WorkflowFolder struct {
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Folder ID.
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`

	// Specifies the absolute path of the folder.
	FolderPath *string `json:"FolderPath,omitnil,omitempty" name:"FolderPath"`

	// Creator ID.
	CreateUserUin *string `json:"CreateUserUin,omitnil,omitempty" name:"CreateUserUin"`
}

type WorkflowFolderPage struct {
	// Data page number, equal to or greater than 1.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Specifies the number of data records displayed per page. value range: 10 to 200.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Total number of folders.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Total pages
	TotalPageNumber *uint64 `json:"TotalPageNumber,omitnil,omitempty" name:"TotalPageNumber"`

	// Folder list.
	Items []*WorkflowFolder `json:"Items,omitnil,omitempty" name:"Items"`
}

type WorkflowInfo struct {
	// Workflow ID.
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// Workflow name.
	WorkflowName *string `json:"WorkflowName,omitnil,omitempty" name:"WorkflowName"`

	// Workflow type: cycle or manual.
	WorkflowType *string `json:"WorkflowType,omitnil,omitempty" name:"WorkflowType"`

	// Owner ID
	// 
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// Creation time.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Last Modification Time
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`

	// Last updated user ID.
	UpdateUserUin *string `json:"UpdateUserUin,omitnil,omitempty" name:"UpdateUserUin"`

	// Workflow description.
	WorkflowDesc *string `json:"WorkflowDesc,omitnil,omitempty" name:"WorkflowDesc"`

	// Creator ID.
	CreateUserUin *string `json:"CreateUserUin,omitnil,omitempty" name:"CreateUserUin"`
}

type WorkflowSchedulerConfiguration struct {
	// Specifies the time zone.
	ScheduleTimeZone *string `json:"ScheduleTimeZone,omitnil,omitempty" name:"ScheduleTimeZone"`

	// Period type. Supported types:
	// 
	// ONEOFF_CYCLE: specifies a one-time cycle.
	// YEAR_CYCLE: specifies the year cycle.
	// MONTH_CYCLE: specifies the monthly cycle.
	// WEEK_CYCLE: specifies the week cycle.
	// DAY_CYCLE: specifies the day cycle.
	// HOUR_CYCLE: specifies the hour cycle.
	// MINUTE_CYCLE: specifies the minute cycle.
	// CRONTAB_CYCLE: specifies the crontab expression type
	CycleType *string `json:"CycleType,omitnil,omitempty" name:"CycleType"`

	// Self-Dependent. Valid values: parallel, serial, orderly. Default value: serial. 
	SelfDepend *string `json:"SelfDepend,omitnil,omitempty" name:"SelfDepend"`

	// Effective Start Time
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Effective End Time.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Workflow dependency. Valid values: yes or no.
	DependencyWorkflow *string `json:"DependencyWorkflow,omitnil,omitempty" name:"DependencyWorkflow"`

	// Execution time. specifies the left-closed interval. example: 00:00.
	ExecutionStartTime *string `json:"ExecutionStartTime,omitnil,omitempty" name:"ExecutionStartTime"`

	// Execution time right closed interval, for example: 23:59.
	ExecutionEndTime *string `json:"ExecutionEndTime,omitnil,omitempty" name:"ExecutionEndTime"`

	// Cron expression
	// 
	CrontabExpression *string `json:"CrontabExpression,omitnil,omitempty" name:"CrontabExpression"`

	// Whether calendar scheduling is enabled. Valid values: 1 (enabled), 0 (disabled).
	CalendarOpen *string `json:"CalendarOpen,omitnil,omitempty" name:"CalendarOpen"`

	// Calendar name.
	CalendarName *string `json:"CalendarName,omitnil,omitempty" name:"CalendarName"`

	// Calendar id.
	CalendarId *string `json:"CalendarId,omitnil,omitempty" name:"CalendarId"`
}

type WorkflowSchedulerConfigurationInfo struct {
	// Specifies the time zone.
	ScheduleTimeZone *string `json:"ScheduleTimeZone,omitnil,omitempty" name:"ScheduleTimeZone"`

	// Period type. Supported types:
	// 
	// ONEOFF_CYCLE: specifies a one-time cycle.
	// YEAR_CYCLE: specifies the year cycle.
	// MONTH_CYCLE: specifies the monthly cycle.
	// WEEK_CYCLE: specifies the week cycle.
	// DAY_CYCLE: specifies the day cycle.
	// HOUR_CYCLE: specifies the hour cycle.
	// MINUTE_CYCLE: specifies the minute cycle.
	// CRONTAB_CYCLE: specifies the crontab expression type
	CycleType *string `json:"CycleType,omitnil,omitempty" name:"CycleType"`

	// Self-Dependent. Valid values: parallel, serial, orderly. Default value: serial. 
	SelfDepend *string `json:"SelfDepend,omitnil,omitempty" name:"SelfDepend"`

	// Effective Start Time
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Effective End Time
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Cron expression
	CrontabExpression *string `json:"CrontabExpression,omitnil,omitempty" name:"CrontabExpression"`

	// Workflow dependency. Valid values: yes or no.
	DependencyWorkflow *string `json:"DependencyWorkflow,omitnil,omitempty" name:"DependencyWorkflow"`

	// 0: Do not modify
	// 1: Reset the task's upstream dependency configuration to the default value
	ModifyCycleValue *string `json:"ModifyCycleValue,omitnil,omitempty" name:"ModifyCycleValue"`

	// The workflow contains cross-workflow dependencies and uses a cron expression for scheduling. If unified scheduling is saved, unsupported dependencies will be disconnected.
	ClearLink *bool `json:"ClearLink,omitnil,omitempty" name:"ClearLink"`

	// Takes effect when ModifyCycleValue = 1. Indicates the default modification of the upstream dependency time dimension. Valid values are:
	// * CRONTAB
	// * DAY
	// * HOUR
	// * LIST_DAY
	// * LIST_HOUR
	// * LIST_MINUTE
	// * MINUTE
	// * MONTH
	// * RANGE_DAY
	// * RANGE_HOUR
	// * RANGE_MINUTE
	// * WEEK
	// * YEAR
	// 
	// https://capi.woa.com/object/detail?product=wedata&env=api_dev&version=2025-08-06&name=WorkflowSchedulerConfigurationInfo
	MainCyclicConfig *string `json:"MainCyclicConfig,omitnil,omitempty" name:"MainCyclicConfig"`

	// Takes effect when ModifyCycleValue = 1. Indicates the default modification of the upstream dependency - instance scope. Valid values are:
	// * ALL_DAY_OF_YEAR
	// * ALL_MONTH_OF_YEAR
	// * CURRENT
	// * CURRENT_DAY
	// * CURRENT_HOUR
	// * CURRENT_MINUTE
	// * CURRENT_MONTH
	// * CURRENT_WEEK
	// * CURRENT_YEAR
	// * PREVIOUS_BEGIN_OF_MONTH
	// * PREVIOUS_DAY
	// * PREVIOUS_DAY_LATER_OFFSET_HOUR
	// * PREVIOUS_DAY_LATER_OFFSET_MINUTE
	// * PREVIOUS_END_OF_MONTH
	// * PREVIOUS_FRIDAY
	// * PREVIOUS_HOUR
	// * PREVIOUS_HOUR_CYCLE
	// * PREVIOUS_HOUR_LATER_OFFSET_MINUTE
	// * PREVIOUS_MINUTE_CYCLE
	// * PREVIOUS_MONTH
	// * PREVIOUS_WEEK
	// * PREVIOUS_WEEKEND
	// * RECENT_DATE
	// 
	// https://capi.woa.com/object/detail?product=wedata&env=api_dev&version=2025-08-06&name=WorkflowSchedulerConfigurationInfo
	SubordinateCyclicConfig *string `json:"SubordinateCyclicConfig,omitnil,omitempty" name:"SubordinateCyclicConfig"`

	// Execution time left closed interval, for example: 00:00. only required when the CycleType is MINUTE_CYCLE.
	ExecutionStartTime *string `json:"ExecutionStartTime,omitnil,omitempty" name:"ExecutionStartTime"`

	// Execution time right closed interval, for example: 23:59. only required when the CycleType is MINUTE_CYCLE.
	ExecutionEndTime *string `json:"ExecutionEndTime,omitnil,omitempty" name:"ExecutionEndTime"`

	// Whether calendar scheduling is enabled. Valid values: 1 (enabled), 0 (disabled).
	CalendarOpen *string `json:"CalendarOpen,omitnil,omitempty" name:"CalendarOpen"`

	// Calendar id.
	CalendarId *string `json:"CalendarId,omitnil,omitempty" name:"CalendarId"`
}