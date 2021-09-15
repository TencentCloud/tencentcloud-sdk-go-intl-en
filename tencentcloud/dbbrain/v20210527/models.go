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

package v20210527

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

type AddUserContactRequest struct {
	*tchttp.BaseRequest

	// Recipient name, which can contain up to 20 letters, digits, spaces, and special symbols `!@#$%^&*()_+-=()` and cannot begin with an underscore.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Email address, which can contain letters, digits, underscores, and the @ symbol, cannot begin with an underscore, and must be unique.
	ContactInfo *string `json:"ContactInfo,omitempty" name:"ContactInfo"`

	// Service type, which is fixed to `mysql`.
	Product *string `json:"Product,omitempty" name:"Product"`
}

func (r *AddUserContactRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddUserContactRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "ContactInfo")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddUserContactRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type AddUserContactResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// ID of successfully added contact.
		Id *int64 `json:"Id,omitempty" name:"Id"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddUserContactResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddUserContactResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ContactItem struct {

	// Contact ID.
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// Contact name.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Contact email.
	Mail *string `json:"Mail,omitempty" name:"Mail"`
}

type CreateDBDiagReportTaskRequest struct {
	*tchttp.BaseRequest

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Start time, such as "2020-11-08T14:00:00+08:00".
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time, such as "2020-11-09T14:00:00+08:00".
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Whether to send an email. Valid values: 0 (yes), 1 (no).
	SendMailFlag *int64 `json:"SendMailFlag,omitempty" name:"SendMailFlag"`

	// Array of contact IDs to receive email.
	ContactPerson []*int64 `json:"ContactPerson,omitempty" name:"ContactPerson"`

	// Array of contact group IDs to receive email.
	ContactGroup []*int64 `json:"ContactGroup,omitempty" name:"ContactGroup"`

	// Service type. Valid values: mysql (TencentDB for MySQL), cynosdb (TDSQL-C for MySQL). Default value: mysql.
	Product *string `json:"Product,omitempty" name:"Product"`
}

func (r *CreateDBDiagReportTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDBDiagReportTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "SendMailFlag")
	delete(f, "ContactPerson")
	delete(f, "ContactGroup")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDBDiagReportTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateDBDiagReportTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Async task request ID, which can be used to query the execution result of an async task.
	// Note: this field may return null, indicating that no valid values can be obtained.
		AsyncRequestId *int64 `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDBDiagReportTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDBDiagReportTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDBDiagReportUrlRequest struct {
	*tchttp.BaseRequest

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Health report task ID, which can be queried through `DescribeDBDiagReportTasks`.
	AsyncRequestId *int64 `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

	// Service type. Valid values: mysql (TencentDB for MySQL), cynosdb (TDSQL-C for MySQL). Default value: mysql.
	Product *string `json:"Product,omitempty" name:"Product"`
}

func (r *CreateDBDiagReportUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDBDiagReportUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "AsyncRequestId")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDBDiagReportUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateDBDiagReportUrlResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Health report URL.
		ReportUrl *string `json:"ReportUrl,omitempty" name:"ReportUrl"`

		// Expiration timestamp of health report URL (in seconds).
		ExpireTime *int64 `json:"ExpireTime,omitempty" name:"ExpireTime"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDBDiagReportUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDBDiagReportUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateMailProfileRequest struct {
	*tchttp.BaseRequest

	// Email configuration.
	ProfileInfo *ProfileInfo `json:"ProfileInfo,omitempty" name:"ProfileInfo"`

	// Configuration level. Valid values: User (user-level), Instance (instance-level). For database inspection emails, it should be `User`. For scheduled task emails, it should be `Instance`.
	ProfileLevel *string `json:"ProfileLevel,omitempty" name:"ProfileLevel"`

	// Configuration name, which needs to be unique. For database inspection emails, this name can be customized as needed. For scheduled task emails, the name should be in the format of "scheduler_" + {instanceId}, such as "schduler_cdb-test".
	ProfileName *string `json:"ProfileName,omitempty" name:"ProfileName"`

	// Configuration type. Valid values: "dbScan_mail_configuration" (email configuration of database inspection report), "scheduler_mail_configuration" (email configuration of scheduled task report).
	ProfileType *string `json:"ProfileType,omitempty" name:"ProfileType"`

	// Service type. Valid values: mysql (TencentDB for MySQL), cynosdb (TDSQL-C for MySQL).
	Product *string `json:"Product,omitempty" name:"Product"`

	// Instance ID bound with the configuration, which is set when the configuration level is `Instance`. Only one instance can be bound at a time. When the configuration level is `User`, leave this parameter empty.
	BindInstanceIds []*string `json:"BindInstanceIds,omitempty" name:"BindInstanceIds"`
}

func (r *CreateMailProfileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMailProfileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProfileInfo")
	delete(f, "ProfileLevel")
	delete(f, "ProfileName")
	delete(f, "ProfileType")
	delete(f, "Product")
	delete(f, "BindInstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateMailProfileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateMailProfileResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateMailProfileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMailProfileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSchedulerMailProfileRequest struct {
	*tchttp.BaseRequest

	// Value range: 1–7, representing Monday to Sunday respectively.
	WeekConfiguration []*int64 `json:"WeekConfiguration,omitempty" name:"WeekConfiguration"`

	// Email configuration.
	ProfileInfo *ProfileInfo `json:"ProfileInfo,omitempty" name:"ProfileInfo"`

	// Configuration name, which needs to be unique. For scheduled task emails, the name should be in the format of "scheduler_" + {instanceId}, such as "schduler_cdb-test".
	ProfileName *string `json:"ProfileName,omitempty" name:"ProfileName"`

	// ID of the instance for which to configure subscription.
	BindInstanceId *string `json:"BindInstanceId,omitempty" name:"BindInstanceId"`

	// Service type. Valid values: mysql (TencentDB for MySQL), cynosdb (TDSQL-C for MySQL). Default value: mysql.
	Product *string `json:"Product,omitempty" name:"Product"`
}

func (r *CreateSchedulerMailProfileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSchedulerMailProfileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WeekConfiguration")
	delete(f, "ProfileInfo")
	delete(f, "ProfileName")
	delete(f, "BindInstanceId")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSchedulerMailProfileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateSchedulerMailProfileResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSchedulerMailProfileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSchedulerMailProfileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSecurityAuditLogExportTaskRequest struct {
	*tchttp.BaseRequest

	// Security audit group ID.
	SecAuditGroupId *string `json:"SecAuditGroupId,omitempty" name:"SecAuditGroupId"`

	// Exported log start time, such as 2020-12-28 00:00:00.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Exported log end time, such as 2020-12-28 01:00:00.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Service type. Valid values: mysql (TencentDB for MySQL).
	Product *string `json:"Product,omitempty" name:"Product"`

	// Log risk level list. Valid values: 0 (no risk), 1 (low risk), 2 (medium risk), 3 (high risk).
	DangerLevels []*int64 `json:"DangerLevels,omitempty" name:"DangerLevels"`
}

func (r *CreateSecurityAuditLogExportTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSecurityAuditLogExportTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecAuditGroupId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Product")
	delete(f, "DangerLevels")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSecurityAuditLogExportTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateSecurityAuditLogExportTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Log export task Id.
		AsyncRequestId *uint64 `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSecurityAuditLogExportTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSecurityAuditLogExportTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSecurityAuditLogExportTasksRequest struct {
	*tchttp.BaseRequest

	// Security audit group ID.
	SecAuditGroupId *string `json:"SecAuditGroupId,omitempty" name:"SecAuditGroupId"`

	// Log export task ID list. This API will ignore task IDs that do not exist or have been deleted.
	AsyncRequestIds []*uint64 `json:"AsyncRequestIds,omitempty" name:"AsyncRequestIds"`

	// Service type. Valid values: mysql (TencentDB for MySQL).
	Product *string `json:"Product,omitempty" name:"Product"`
}

func (r *DeleteSecurityAuditLogExportTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSecurityAuditLogExportTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecAuditGroupId")
	delete(f, "AsyncRequestIds")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSecurityAuditLogExportTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSecurityAuditLogExportTasksResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSecurityAuditLogExportTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSecurityAuditLogExportTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAllUserContactRequest struct {
	*tchttp.BaseRequest

	// Service type, which is fixed to `mysql`.
	Product *string `json:"Product,omitempty" name:"Product"`

	// Array of contact names. Fuzzy search is supported.
	Names []*string `json:"Names,omitempty" name:"Names"`
}

func (r *DescribeAllUserContactRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAllUserContactRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Product")
	delete(f, "Names")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAllUserContactRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAllUserContactResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of contacts.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Contact information.
	// Note: this field may return null, indicating that no valid values can be obtained.
		Contacts []*ContactItem `json:"Contacts,omitempty" name:"Contacts"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAllUserContactResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAllUserContactResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAllUserGroupRequest struct {
	*tchttp.BaseRequest

	// Service type, which is fixed to `mysql`.
	Product *string `json:"Product,omitempty" name:"Product"`

	// Array of contact group names. Fuzzy search is supported.
	Names []*string `json:"Names,omitempty" name:"Names"`
}

func (r *DescribeAllUserGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAllUserGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Product")
	delete(f, "Names")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAllUserGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAllUserGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of groups.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Group information.
	// Note: this field may return null, indicating that no valid values can be obtained.
		Groups []*GroupItem `json:"Groups,omitempty" name:"Groups"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAllUserGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAllUserGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBDiagEventRequest struct {
	*tchttp.BaseRequest

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Event ID, which can be obtained through the `DescribeDBDiagHistory` API.
	EventId *int64 `json:"EventId,omitempty" name:"EventId"`

	// Service type. Valid values: mysql (TencentDB for MySQL), cynosdb (TDSQL-C for MySQL). Default value: mysql.
	Product *string `json:"Product,omitempty" name:"Product"`
}

func (r *DescribeDBDiagEventRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBDiagEventRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "EventId")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBDiagEventRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBDiagEventResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Diagnosis item.
		DiagItem *string `json:"DiagItem,omitempty" name:"DiagItem"`

		// Diagnosis type.
		DiagType *string `json:"DiagType,omitempty" name:"DiagType"`

		// Event ID.
		EventId *int64 `json:"EventId,omitempty" name:"EventId"`

		// Diagnosis event details. If there is no additional explanation information, the output will be empty.
		Explanation *string `json:"Explanation,omitempty" name:"Explanation"`

		// Diagnosis summary.
		Outline *string `json:"Outline,omitempty" name:"Outline"`

		// Found problem.
		Problem *string `json:"Problem,omitempty" name:"Problem"`

		// Severity, which can be divided into 5 levels: 1: fatal, 2: severe, 3: warning, 4: notice, 5: healthy.
		Severity *int64 `json:"Severity,omitempty" name:"Severity"`

		// Start time
		StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

		// Suggestions. If there are no suggestions, the output will be empty.
		Suggestions *string `json:"Suggestions,omitempty" name:"Suggestions"`

		// Reserved field.
	// Note: this field may return null, indicating that no valid values can be obtained.
		Metric *string `json:"Metric,omitempty" name:"Metric"`

		// End time.
		EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBDiagEventResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBDiagEventResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBDiagHistoryRequest struct {
	*tchttp.BaseRequest

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Start time, such as "2019-09-10 12:13:14".
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time, such as "2019-09-11 12:13:14". The interval between the end time and the start time can be up to 2 days.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Service type. Valid values: mysql (TencentDB for MySQL), cynosdb (TDSQL-C for MySQL). Default value: mysql.
	Product *string `json:"Product,omitempty" name:"Product"`
}

func (r *DescribeDBDiagHistoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBDiagHistoryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBDiagHistoryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBDiagHistoryResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Event description.
		Events []*DiagHistoryEventItem `json:"Events,omitempty" name:"Events"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBDiagHistoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBDiagHistoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBDiagReportTasksRequest struct {
	*tchttp.BaseRequest

	// Start time of the first task in the format of yyyy-MM-dd HH:mm:ss, such as 2019-09-10 12:13:14. It is used for queries by time range.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time of the last task in the format of yyyy-MM-dd HH:mm:ss, such as 2019-09-10 12:13:14. It is used for queries by time range.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Array of instance IDs, which is used to filter the task list of a specified instance.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// Source that triggers the task. Valid values: `DAILY_INSPECTION` (instance inspection), `SCHEDULED` (scheduled task), and `MANUAL` (manual trigger).
	Sources []*string `json:"Sources,omitempty" name:"Sources"`

	// Health level. Valid values: `HEALTH` (healthy), `SUB_HEALTH` (suboptimal), `RISK` (risky), and `HIGH_RISK` (critical).
	HealthLevels *string `json:"HealthLevels,omitempty" name:"HealthLevels"`

	// Task status. Valid values: `created` (created), `chosen` (to be executed), `running` (being executed), `failed` (failed), and `finished` (completed).
	TaskStatuses *string `json:"TaskStatuses,omitempty" name:"TaskStatuses"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Service type. Valid values: mysql (TencentDB for MySQL), cynosdb (TDSQL-C for MySQL). Default value: mysql.
	Product *string `json:"Product,omitempty" name:"Product"`
}

func (r *DescribeDBDiagReportTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBDiagReportTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "InstanceIds")
	delete(f, "Sources")
	delete(f, "HealthLevels")
	delete(f, "TaskStatuses")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBDiagReportTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBDiagReportTasksResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of tasks.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Task list.
		Tasks []*HealthReportTask `json:"Tasks,omitempty" name:"Tasks"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBDiagReportTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBDiagReportTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBSpaceStatusRequest struct {
	*tchttp.BaseRequest

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Query period in days. The end date is the current date, and the query period is 7 days by default.
	RangeDays *int64 `json:"RangeDays,omitempty" name:"RangeDays"`

	// Service type. Valid values: mysql (TencentDB for MySQL), cynosdb (TDSQL-C for MySQL). Default value: mysql.
	Product *string `json:"Product,omitempty" name:"Product"`
}

func (r *DescribeDBSpaceStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBSpaceStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "RangeDays")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBSpaceStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBSpaceStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Disk usage growth in MB.
		Growth *int64 `json:"Growth,omitempty" name:"Growth"`

		// Available disk space in MB.
		Remain *int64 `json:"Remain,omitempty" name:"Remain"`

		// Total disk space in MB.
		Total *int64 `json:"Total,omitempty" name:"Total"`

		// Estimated number of available days.
		AvailableDays *int64 `json:"AvailableDays,omitempty" name:"AvailableDays"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBSpaceStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBSpaceStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDiagDBInstancesRequest struct {
	*tchttp.BaseRequest

	// Whether it is an instance supported by DBbrain. It is fixed to `true`.
	IsSupported *bool `json:"IsSupported,omitempty" name:"IsSupported"`

	// Service type. Valid values: mysql (TencentDB for MySQL), cynosdb (TDSQL-C for MySQL). Default value: mysql.
	Product *string `json:"Product,omitempty" name:"Product"`

	// Pagination parameter indicating the offset.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Pagination parameter. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Query by instance name.
	InstanceNames []*string `json:"InstanceNames,omitempty" name:"InstanceNames"`

	// Query by instance ID.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// Query by region.
	Regions []*string `json:"Regions,omitempty" name:"Regions"`
}

func (r *DescribeDiagDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDiagDBInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IsSupported")
	delete(f, "Product")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "InstanceNames")
	delete(f, "InstanceIds")
	delete(f, "Regions")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDiagDBInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDiagDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of instances.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Status of all instance inspection. 0: all instance inspection enabled, 1: all instance inspection disabled.
		DbScanStatus *int64 `json:"DbScanStatus,omitempty" name:"DbScanStatus"`

		// Instance information.
		Items []*InstanceInfo `json:"Items,omitempty" name:"Items"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDiagDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDiagDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeHealthScoreRequest struct {
	*tchttp.BaseRequest

	// Instance ID for which to get the health score.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Time to get the health score in the format of `2019-09-10 12:13:14`.
	Time *string `json:"Time,omitempty" name:"Time"`

	// Service type. Valid values: mysql (TencentDB for MySQL), cynosdb (TDSQL-C for MySQL). Default value: mysql.
	Product *string `json:"Product,omitempty" name:"Product"`
}

func (r *DescribeHealthScoreRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHealthScoreRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Time")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeHealthScoreRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeHealthScoreResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Health score and deduction for exceptions.
		Data *HealthScoreInfo `json:"Data,omitempty" name:"Data"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeHealthScoreResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHealthScoreResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMailProfileRequest struct {
	*tchttp.BaseRequest

	// Configuration type. Valid values: "dbScan_mail_configuration" (email configuration of database inspection report), "scheduler_mail_configuration" (email configuration of scheduled task report).
	ProfileType *string `json:"ProfileType,omitempty" name:"ProfileType"`

	// Service type. Valid values: mysql (TencentDB for MySQL), cynosdb (TDSQL-C for MySQL). Default value: mysql.
	Product *string `json:"Product,omitempty" name:"Product"`

	// Pagination offset.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of results per page in paginated queries. Maximum value: 50.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Query by email configuration name. The name of the scheduled task email configuration should be in the format of "scheduler_"+{instanceId}.
	ProfileName *string `json:"ProfileName,omitempty" name:"ProfileName"`
}

func (r *DescribeMailProfileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMailProfileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProfileType")
	delete(f, "Product")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "ProfileName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMailProfileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMailProfileResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Email configuration details.
	// Note: this field may return null, indicating that no valid values can be obtained.
		ProfileList []*UserProfile `json:"ProfileList,omitempty" name:"ProfileList"`

		// Total number of email templates.
	// Note: this field may return null, indicating that no valid values can be obtained.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMailProfileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMailProfileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMySqlProcessListRequest struct {
	*tchttp.BaseRequest

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Thread ID, which is used to filter the thread list.
	ID *uint64 `json:"ID,omitempty" name:"ID"`

	// Thread operation account name, which is used to filter the thread list.
	User *string `json:"User,omitempty" name:"User"`

	// Thread operation host address, which is used to filter the thread list.
	Host *string `json:"Host,omitempty" name:"Host"`

	// Thread operation database, which is used to filter the thread list.
	DB *string `json:"DB,omitempty" name:"DB"`

	// Thread operation status, which is used to filter the thread list.
	State *string `json:"State,omitempty" name:"State"`

	// Thread execution type, which is used to filter the thread list.
	Command *string `json:"Command,omitempty" name:"Command"`

	// Minimum operation duration of the thread in seconds, which is used to filter the list of threads whose operation duration is greater than this value.
	Time *uint64 `json:"Time,omitempty" name:"Time"`

	// Thread operation statement, which is used to filter the thread list.
	Info *string `json:"Info,omitempty" name:"Info"`

	// Number of returned results. Default value: 20.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Service type. Valid values: mysql (TencentDB for MySQL), cynosdb (TDSQL-C for MySQL). Default value: mysql.
	Product *string `json:"Product,omitempty" name:"Product"`
}

func (r *DescribeMySqlProcessListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMySqlProcessListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ID")
	delete(f, "User")
	delete(f, "Host")
	delete(f, "DB")
	delete(f, "State")
	delete(f, "Command")
	delete(f, "Time")
	delete(f, "Info")
	delete(f, "Limit")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMySqlProcessListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMySqlProcessListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// List of real-time threads.
		ProcessList []*MySqlProcess `json:"ProcessList,omitempty" name:"ProcessList"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMySqlProcessListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMySqlProcessListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecurityAuditLogDownloadUrlsRequest struct {
	*tchttp.BaseRequest

	// Security audit group ID.
	SecAuditGroupId *string `json:"SecAuditGroupId,omitempty" name:"SecAuditGroupId"`

	// Async task Id.
	AsyncRequestId *uint64 `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

	// Service type. Valid values: mysql (TencentDB for MySQL).
	Product *string `json:"Product,omitempty" name:"Product"`
}

func (r *DescribeSecurityAuditLogDownloadUrlsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityAuditLogDownloadUrlsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecAuditGroupId")
	delete(f, "AsyncRequestId")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecurityAuditLogDownloadUrlsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecurityAuditLogDownloadUrlsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// List of COS URLs of the export results. If the result set is large, it may be divided into multiple URLs for download.
		Urls []*string `json:"Urls,omitempty" name:"Urls"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSecurityAuditLogDownloadUrlsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityAuditLogDownloadUrlsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecurityAuditLogExportTasksRequest struct {
	*tchttp.BaseRequest

	// Security audit group ID.
	SecAuditGroupId *string `json:"SecAuditGroupId,omitempty" name:"SecAuditGroupId"`

	// Service type. Valid values: mysql (TencentDB for MySQL).
	Product *string `json:"Product,omitempty" name:"Product"`

	// List of log export task IDs.
	AsyncRequestIds []*uint64 `json:"AsyncRequestIds,omitempty" name:"AsyncRequestIds"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeSecurityAuditLogExportTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityAuditLogExportTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecAuditGroupId")
	delete(f, "Product")
	delete(f, "AsyncRequestIds")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecurityAuditLogExportTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecurityAuditLogExportTasksResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// List of security audit log export tasks.
		Tasks []*SecLogExportTaskInfo `json:"Tasks,omitempty" name:"Tasks"`

		// Total numbers of security audit log export tasks.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSecurityAuditLogExportTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityAuditLogExportTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSlowLogTimeSeriesStatsRequest struct {
	*tchttp.BaseRequest

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Start time, such as "2019-09-10 12:13:14".
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time, such as "2019-09-10 12:13:14". The interval between the end time and the start time can be up to 7 days.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Service type. Valid values: mysql (TencentDB for MySQL), cynosdb (TDSQL-C for MySQL). Default value: mysql.
	Product *string `json:"Product,omitempty" name:"Product"`
}

func (r *DescribeSlowLogTimeSeriesStatsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlowLogTimeSeriesStatsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSlowLogTimeSeriesStatsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSlowLogTimeSeriesStatsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Time range in seconds in histogram.
		Period *int64 `json:"Period,omitempty" name:"Period"`

		// Number of slow logs in specified time range.
		TimeSeries []*TimeSlice `json:"TimeSeries,omitempty" name:"TimeSeries"`

		// Instance CPU utilization monitoring data in specified time range.
		SeriesData *MonitorMetricSeriesData `json:"SeriesData,omitempty" name:"SeriesData"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSlowLogTimeSeriesStatsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlowLogTimeSeriesStatsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSlowLogTopSqlsRequest struct {
	*tchttp.BaseRequest

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Start time, such as "2019-09-10 12:13:14".
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time, such as "2019-09-10 12:13:14". The interval between the end time and the start time can be up to 7 days.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Sorting key. Valid values: QueryTime, ExecTimes, RowsSent, LockTime, RowsExamined. Default value: QueryTime.
	SortBy *string `json:"SortBy,omitempty" name:"SortBy"`

	// Sorting order. Valid values: ASC (ascending), DESC (descending). Default value: DESC.
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Database name array.
	SchemaList []*SchemaItem `json:"SchemaList,omitempty" name:"SchemaList"`

	// Service type. Valid values: mysql (TencentDB for MySQL), cynosdb (TDSQL-C for MySQL). Default value: mysql.
	Product *string `json:"Product,omitempty" name:"Product"`
}

func (r *DescribeSlowLogTopSqlsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlowLogTopSqlsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "SortBy")
	delete(f, "OrderBy")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "SchemaList")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSlowLogTopSqlsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSlowLogTopSqlsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of eligible entries.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// List of top slow SQL statements
		Rows []*SlowLogTopSqlItem `json:"Rows,omitempty" name:"Rows"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSlowLogTopSqlsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlowLogTopSqlsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSlowLogUserHostStatsRequest struct {
	*tchttp.BaseRequest

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Start time of the time range in the format of yyyy-MM-dd HH:mm:ss, such as 2019-09-10 12:13:14.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time of the time range in the format of yyyy-MM-dd HH:mm:ss, such as 2019-09-10 12:13:14.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Service type. Valid values: mysql (TencentDB for MySQL), cynosdb (TDSQL-C for MySQL). Default value: mysql.
	Product *string `json:"Product,omitempty" name:"Product"`

	// MD5 value of SOL template
	Md5 *string `json:"Md5,omitempty" name:"Md5"`
}

func (r *DescribeSlowLogUserHostStatsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlowLogUserHostStatsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Product")
	delete(f, "Md5")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSlowLogUserHostStatsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSlowLogUserHostStatsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of source addresses.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Detailed list of the proportion of slow logs from each source address.
		Items []*SlowLogHost `json:"Items,omitempty" name:"Items"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSlowLogUserHostStatsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlowLogUserHostStatsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTopSpaceSchemaTimeSeriesRequest struct {
	*tchttp.BaseRequest

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Number of returned top databases. Maximum value: 100. Default value: 20.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Field used to sort top databases. Valid values: DataLength, IndexLength, TotalLength, DataFree, FragRatio, TableRows, PhysicalFileSize (supported only by TencentDB for MySQL instances). For TencentDB for MySQL instances, the default value is `PhysicalFileSize`. For other database instances, the default value is `TotalLength`.
	SortBy *string `json:"SortBy,omitempty" name:"SortBy"`

	// Start date, such as "2021-01-01". It can be as early as 29 days before the current date and is 6 days before the end date by default.
	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

	// End date, such as "2021-01-01". It can be as early as 29 days before the current date and is the current date by default.
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`

	// Service type. Valid values: mysql (TencentDB for MySQL), cynosdb (TDSQL-C for MySQL). Default value: mysql.
	Product *string `json:"Product,omitempty" name:"Product"`
}

func (r *DescribeTopSpaceSchemaTimeSeriesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopSpaceSchemaTimeSeriesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Limit")
	delete(f, "SortBy")
	delete(f, "StartDate")
	delete(f, "EndDate")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTopSpaceSchemaTimeSeriesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTopSpaceSchemaTimeSeriesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Time series list of the returned space statistics of top databases.
		TopSpaceSchemaTimeSeries []*SchemaSpaceTimeSeries `json:"TopSpaceSchemaTimeSeries,omitempty" name:"TopSpaceSchemaTimeSeries"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTopSpaceSchemaTimeSeriesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopSpaceSchemaTimeSeriesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTopSpaceSchemasRequest struct {
	*tchttp.BaseRequest

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Number of returned top databases. Maximum value: 100. Default value: 20.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Field used to sort top databases. Valid values: DataLength, IndexLength, TotalLength, DataFree, FragRatio, TableRows, PhysicalFileSize (supported only by TencentDB for MySQL instances). For TencentDB for MySQL instances, the default value is `PhysicalFileSize`. For other database instances, the default value is `TotalLength`.
	SortBy *string `json:"SortBy,omitempty" name:"SortBy"`

	// Service type. Valid values: mysql (TencentDB for MySQL), cynosdb (TDSQL-C for MySQL). Default value: mysql.
	Product *string `json:"Product,omitempty" name:"Product"`
}

func (r *DescribeTopSpaceSchemasRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopSpaceSchemasRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Limit")
	delete(f, "SortBy")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTopSpaceSchemasRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTopSpaceSchemasResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// List of the returned space statistics of top databases.
		TopSpaceSchemas []*SchemaSpaceData `json:"TopSpaceSchemas,omitempty" name:"TopSpaceSchemas"`

		// Timestamp (in seconds) of database space data collection points
		Timestamp *int64 `json:"Timestamp,omitempty" name:"Timestamp"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTopSpaceSchemasResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopSpaceSchemasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTopSpaceTableTimeSeriesRequest struct {
	*tchttp.BaseRequest

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Number of returned top tables. Maximum value: 100. Default value: 20.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Field used to sort top tables. Valid values: DataLength, IndexLength, TotalLength, DataFree, FragRatio, TableRows, PhysicalFileSize. Default value: PhysicalFileSize.
	SortBy *string `json:"SortBy,omitempty" name:"SortBy"`

	// Start date, such as "2021-01-01". It can be as early as 29 days before the current date and is 6 days before the end date by default.
	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

	// End date, such as "2021-01-01". It can be as early as 29 days before the current date and is the current date by default.
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`

	// Service type. Valid values: mysql (TencentDB for MySQL), cynosdb (TDSQL-C for MySQL). Default value: mysql.
	Product *string `json:"Product,omitempty" name:"Product"`
}

func (r *DescribeTopSpaceTableTimeSeriesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopSpaceTableTimeSeriesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Limit")
	delete(f, "SortBy")
	delete(f, "StartDate")
	delete(f, "EndDate")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTopSpaceTableTimeSeriesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTopSpaceTableTimeSeriesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Time series list of the returned space statistics of top tables.
		TopSpaceTableTimeSeries []*TableSpaceTimeSeries `json:"TopSpaceTableTimeSeries,omitempty" name:"TopSpaceTableTimeSeries"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTopSpaceTableTimeSeriesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopSpaceTableTimeSeriesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTopSpaceTablesRequest struct {
	*tchttp.BaseRequest

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Number of returned top tables. Maximum value: 100. Default value: 20.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Field used to sort top tables. Valid values: DataLength, IndexLength, TotalLength, DataFree, FragRatio, TableRows, PhysicalFileSize (only supported for TencentDB for MySQL instances). For TencentDB for MySQL instances, the default value is `PhysicalFileSize`. For other database instances, the default value is `TotalLength`.
	SortBy *string `json:"SortBy,omitempty" name:"SortBy"`

	// Service type. Valid values: mysql (TencentDB for MySQL), cynosdb (TDSQL-C for MySQL). Default value: mysql.
	Product *string `json:"Product,omitempty" name:"Product"`
}

func (r *DescribeTopSpaceTablesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopSpaceTablesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Limit")
	delete(f, "SortBy")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTopSpaceTablesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTopSpaceTablesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// List of the returned space statistics of top tables.
		TopSpaceTables []*TableSpaceData `json:"TopSpaceTables,omitempty" name:"TopSpaceTables"`

		// Timestamp (in seconds) of tablespace data collection points
		Timestamp *int64 `json:"Timestamp,omitempty" name:"Timestamp"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTopSpaceTablesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopSpaceTablesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserSqlAdviceRequest struct {
	*tchttp.BaseRequest

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// SQL statement.
	SqlText *string `json:"SqlText,omitempty" name:"SqlText"`

	// Database name.
	Schema *string `json:"Schema,omitempty" name:"Schema"`
}

func (r *DescribeUserSqlAdviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserSqlAdviceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SqlText")
	delete(f, "Schema")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserSqlAdviceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserSqlAdviceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// SQL statement optimization suggestions, which can be parsed into JSON arrays. If there is no need for optimization, the output will be empty.
		Advices *string `json:"Advices,omitempty" name:"Advices"`

		// Notes of SQL statement optimization suggestions, which can be parsed into String arrays. If there is no need for optimization, the output will be empty.
		Comments *string `json:"Comments,omitempty" name:"Comments"`

		// SQL statement.
		SqlText *string `json:"SqlText,omitempty" name:"SqlText"`

		// Database name.
		Schema *string `json:"Schema,omitempty" name:"Schema"`

		// DDL information of related tables, which can be parsed into JSON arrays.
		Tables *string `json:"Tables,omitempty" name:"Tables"`

		// SQL execution plan, which can be parsed into JSON arrays. If there is no need for optimization, the output will be empty.
		SqlPlan *string `json:"SqlPlan,omitempty" name:"SqlPlan"`

		// Cost saving details after SQL statement optimization, which can be parsed into JSON arrays. If there is no need for optimization, the output will be empty.
		Cost *string `json:"Cost,omitempty" name:"Cost"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserSqlAdviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserSqlAdviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DiagHistoryEventItem struct {

	// Diagnosis type.
	DiagType *string `json:"DiagType,omitempty" name:"DiagType"`

	// End time.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Start time.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Unique event ID.
	EventId *int64 `json:"EventId,omitempty" name:"EventId"`

	// Severity, which can be divided into 5 levels: 1: fatal, 2: severe, 3: warning, 4: notice, 5: healthy.
	Severity *int64 `json:"Severity,omitempty" name:"Severity"`

	// Diagnosis summary.
	Outline *string `json:"Outline,omitempty" name:"Outline"`

	// Diagnosis item description.
	DiagItem *string `json:"DiagItem,omitempty" name:"DiagItem"`

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Reserved field.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Metric *string `json:"Metric,omitempty" name:"Metric"`

	// Region.
	Region *string `json:"Region,omitempty" name:"Region"`
}

type EventInfo struct {

	// Event ID.
	EventId *int64 `json:"EventId,omitempty" name:"EventId"`

	// Diagnosis type.
	DiagType *string `json:"DiagType,omitempty" name:"DiagType"`

	// Start time.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Summary.
	Outline *string `json:"Outline,omitempty" name:"Outline"`

	// Severity, which can be divided into 5 levels: 1: fatal, 2: severe, 3: warning, 4: notice, 5: healthy.
	Severity *int64 `json:"Severity,omitempty" name:"Severity"`

	// Deduction.
	ScoreLost *int64 `json:"ScoreLost,omitempty" name:"ScoreLost"`

	// Reserved field.
	Metric *string `json:"Metric,omitempty" name:"Metric"`

	// Number of alarms.
	Count *int64 `json:"Count,omitempty" name:"Count"`
}

type GroupItem struct {

	// Group ID.
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// Group name.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Number of group members.
	MemberCount *int64 `json:"MemberCount,omitempty" name:"MemberCount"`
}

type HealthReportTask struct {

	// Async task request ID.
	AsyncRequestId *int64 `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

	// Source that triggers the task. Valid values: `DAILY_INSPECTION` (instance inspection), `SCHEDULED` (scheduled task), and `MANUAL` (manual trigger).
	Source *string `json:"Source,omitempty" name:"Source"`

	// Task progress in %.
	Progress *int64 `json:"Progress,omitempty" name:"Progress"`

	// Task creation time.
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Task start time.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Task end time.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Basic information of the instance to which the task belongs.
	InstanceInfo *InstanceBasicInfo `json:"InstanceInfo,omitempty" name:"InstanceInfo"`

	// Health information in health report.
	HealthStatus *HealthStatus `json:"HealthStatus,omitempty" name:"HealthStatus"`
}

type HealthScoreInfo struct {

	// Exception details.
	IssueTypes []*IssueTypeInfo `json:"IssueTypes,omitempty" name:"IssueTypes"`

	// Total number of exceptions.
	EventsTotalCount *int64 `json:"EventsTotalCount,omitempty" name:"EventsTotalCount"`

	// Health score.
	HealthScore *int64 `json:"HealthScore,omitempty" name:"HealthScore"`

	// Health level, such as "HEALTH", "SUB_HEALTH", "RISK", and "HIGH_RISK".
	HealthLevel *string `json:"HealthLevel,omitempty" name:"HealthLevel"`
}

type HealthStatus struct {

	// Health score out of 100 points.
	HealthScore *int64 `json:"HealthScore,omitempty" name:"HealthScore"`

	// Health level. Valid values: `HEALTH` (healthy), `SUB_HEALTH` (suboptimal), `RISK` (risky), and `HIGH_RISK` (critical).
	HealthLevel *string `json:"HealthLevel,omitempty" name:"HealthLevel"`

	// Total deducted scores.
	ScoreLost *int64 `json:"ScoreLost,omitempty" name:"ScoreLost"`

	// Deduction details.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ScoreDetails []*ScoreDetail `json:"ScoreDetails,omitempty" name:"ScoreDetails"`
}

type InstanceBasicInfo struct {

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Instance name.
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// Private IP of instance.
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// Private port of instance.
	Vport *int64 `json:"Vport,omitempty" name:"Vport"`

	// Instance service.
	Product *string `json:"Product,omitempty" name:"Product"`

	// Instance engine version.
	EngineVersion *string `json:"EngineVersion,omitempty" name:"EngineVersion"`
}

type InstanceConfs struct {

	// Whether to enable database inspection. Valid values: Yes, No.
	DailyInspection *string `json:"DailyInspection,omitempty" name:"DailyInspection"`

	// Whether to enable instance overview. Valid values: Yes, No.
	OverviewDisplay *string `json:"OverviewDisplay,omitempty" name:"OverviewDisplay"`
}

type InstanceInfo struct {

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Instance name.
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// Instance region.
	Region *string `json:"Region,omitempty" name:"Region"`

	// Health score.
	HealthScore *int64 `json:"HealthScore,omitempty" name:"HealthScore"`

	// Service.
	Product *string `json:"Product,omitempty" name:"Product"`

	// Number of exceptions.
	EventCount *int64 `json:"EventCount,omitempty" name:"EventCount"`

	// Instance type. Valid values: 1 (MASTER), 2 (DR), 3 (RO), 4 (SDR)
	InstanceType *int64 `json:"InstanceType,omitempty" name:"InstanceType"`

	// Number of cores.
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

	// Memory in MB.
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// Disk storage in GB.
	Volume *int64 `json:"Volume,omitempty" name:"Volume"`

	// Database version.
	EngineVersion *string `json:"EngineVersion,omitempty" name:"EngineVersion"`

	// Private network address.
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// Private network port.
	Vport *int64 `json:"Vport,omitempty" name:"Vport"`

	// Access source.
	Source *string `json:"Source,omitempty" name:"Source"`

	// Group ID.
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// Group name.
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// Instance status. Valid values: 0 (delivering), 1 (running), 4 (terminating), 5 (isolated)
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// Unified subnet ID.
	UniqSubnetId *string `json:"UniqSubnetId,omitempty" name:"UniqSubnetId"`

	// TencentDB instance type.
	DeployMode *string `json:"DeployMode,omitempty" name:"DeployMode"`

	// TencentDB instance initialization flag. Valid values: 0 (not initialized), 1 (initialized).
	InitFlag *int64 `json:"InitFlag,omitempty" name:"InitFlag"`

	// Task status.
	TaskStatus *int64 `json:"TaskStatus,omitempty" name:"TaskStatus"`

	// Unified VPC ID.
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// Instance inspection/overview status.
	InstanceConf *InstanceConfs `json:"InstanceConf,omitempty" name:"InstanceConf"`

	// Resource expiration time.
	DeadlineTime *string `json:"DeadlineTime,omitempty" name:"DeadlineTime"`

	// Whether it is an instance supported by DBbrain.
	IsSupported *bool `json:"IsSupported,omitempty" name:"IsSupported"`

	// Status of instance security audit log. Valid values: ON (enabled), OFF (disabled).
	SecAuditStatus *string `json:"SecAuditStatus,omitempty" name:"SecAuditStatus"`

	// Status of instance audit log. Valid values: ALL_AUDIT (full audit is enabled), RULE_AUDIT (rule audit is enabled), UNBOUND (audit is disabled).
	AuditPolicyStatus *string `json:"AuditPolicyStatus,omitempty" name:"AuditPolicyStatus"`

	// Running status of instance audit log. Valid values: normal (running), paused (suspension due to overdue payment).
	AuditRunningStatus *string `json:"AuditRunningStatus,omitempty" name:"AuditRunningStatus"`
}

type IssueTypeInfo struct {

	// Metric categories: AVAILABILITY, MAINTAINABILITY, PERFORMANCE, and RELIABILITY
	IssueType *string `json:"IssueType,omitempty" name:"IssueType"`

	// Exception.
	Events []*EventInfo `json:"Events,omitempty" name:"Events"`

	// Total number of exceptions.
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

type KillMySqlThreadsRequest struct {
	*tchttp.BaseRequest

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// The stage of a session killing task. Valid values: Prepare (preparation stage), Commit (commit stage).
	Stage *string `json:"Stage,omitempty" name:"Stage"`

	// The ID list of MySQL sessions to be killed. This parameter is used in the “Prepare” stage.
	Threads []*int64 `json:"Threads,omitempty" name:"Threads"`

	// Execution ID. This parameter is used in the “Commit” stage.
	SqlExecId *string `json:"SqlExecId,omitempty" name:"SqlExecId"`

	// Service type. Valid values: mysql (TencentDB for MySQL), cynosdb (TDSQL-C for MySQL). Default value: mysql.
	Product *string `json:"Product,omitempty" name:"Product"`
}

func (r *KillMySqlThreadsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *KillMySqlThreadsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Stage")
	delete(f, "Threads")
	delete(f, "SqlExecId")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "KillMySqlThreadsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type KillMySqlThreadsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The ID list of MySQL sessions that have been killed.
		Threads []*int64 `json:"Threads,omitempty" name:"Threads"`

		// Execution ID, which is output in the “Prepare” stage and used to specify the ID of the session to be killed in the “Commit” stage.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
		SqlExecId *string `json:"SqlExecId,omitempty" name:"SqlExecId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *KillMySqlThreadsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *KillMySqlThreadsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MailConfiguration struct {

	// Whether to enable email sending. Valid values: 0 (no), 1 (yes).
	SendMail *int64 `json:"SendMail,omitempty" name:"SendMail"`

	// Region configuration, such as "ap-guangzhou" and "ap-shanghai". For the inspection email sending template, configure the region where you need to send the inspection email. For the subscription email sending template, configure the region where the current subscribed instance resides.
	Region []*string `json:"Region,omitempty" name:"Region"`

	// Sends a report with the specified health level, such as "HEALTH", "SUB_HEALTH", "RISK", and "HIGH_RISK".
	HealthStatus []*string `json:"HealthStatus,omitempty" name:"HealthStatus"`

	// Contact ID. Either `ContactPerson` or `ContactGroup` should be passed in.
	ContactPerson []*int64 `json:"ContactPerson,omitempty" name:"ContactPerson"`

	// Contact group ID. Either `ContactPerson` or `ContactGroup` should be passed in.
	ContactGroup []*int64 `json:"ContactGroup,omitempty" name:"ContactGroup"`
}

type ModifyDiagDBInstanceConfRequest struct {
	*tchttp.BaseRequest

	// Instance configuration, including inspection and overview switch.
	InstanceConfs *InstanceConfs `json:"InstanceConfs,omitempty" name:"InstanceConfs"`

	// Target regions of the request. If the value is `All`, it is applied to all regions.
	Regions *string `json:"Regions,omitempty" name:"Regions"`

	// Service type. Valid values: mysql (TencentDB for MySQL), cynosdb (TDSQL-C for MySQL).
	Product *string `json:"Product,omitempty" name:"Product"`

	// ID of the instance to modify.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *ModifyDiagDBInstanceConfRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDiagDBInstanceConfRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceConfs")
	delete(f, "Regions")
	delete(f, "Product")
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDiagDBInstanceConfRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDiagDBInstanceConfResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDiagDBInstanceConfResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDiagDBInstanceConfResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MonitorFloatMetric struct {

	// Metric name.
	Metric *string `json:"Metric,omitempty" name:"Metric"`

	// Metric unit.
	Unit *string `json:"Unit,omitempty" name:"Unit"`

	// Metric value.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Values []*float64 `json:"Values,omitempty" name:"Values"`
}

type MonitorFloatMetricSeriesData struct {

	// Monitoring metric.
	Series []*MonitorFloatMetric `json:"Series,omitempty" name:"Series"`

	// Timestamp corresponding to monitoring metric.
	Timestamp []*int64 `json:"Timestamp,omitempty" name:"Timestamp"`
}

type MonitorMetric struct {

	// Metric name.
	Metric *string `json:"Metric,omitempty" name:"Metric"`

	// Metric unit.
	Unit *string `json:"Unit,omitempty" name:"Unit"`

	// Metric value.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Values []*float64 `json:"Values,omitempty" name:"Values"`
}

type MonitorMetricSeriesData struct {

	// Monitoring metric.
	Series []*MonitorMetric `json:"Series,omitempty" name:"Series"`

	// Timestamp corresponding to monitoring metric.
	Timestamp []*int64 `json:"Timestamp,omitempty" name:"Timestamp"`
}

type MySqlProcess struct {

	// Thread ID.
	ID *string `json:"ID,omitempty" name:"ID"`

	// Thread operation account name.
	User *string `json:"User,omitempty" name:"User"`

	// Thread operation host address.
	Host *string `json:"Host,omitempty" name:"Host"`

	// Thread operation database.
	DB *string `json:"DB,omitempty" name:"DB"`

	// Thread operation status.
	State *string `json:"State,omitempty" name:"State"`

	// Thread execution type.
	Command *string `json:"Command,omitempty" name:"Command"`

	// Thread operation duration in seconds.
	Time *string `json:"Time,omitempty" name:"Time"`

	// Thread operation statement.
	Info *string `json:"Info,omitempty" name:"Info"`
}

type ProfileInfo struct {

	// Email language, such as `en`.
	Language *string `json:"Language,omitempty" name:"Language"`

	// Email template content.
	MailConfiguration *MailConfiguration `json:"MailConfiguration,omitempty" name:"MailConfiguration"`
}

type SchemaItem struct {

	// Database name
	Schema *string `json:"Schema,omitempty" name:"Schema"`
}

type SchemaSpaceData struct {

	// Database name.
	TableSchema *string `json:"TableSchema,omitempty" name:"TableSchema"`

	// Data space in MB.
	DataLength *float64 `json:"DataLength,omitempty" name:"DataLength"`

	// Index space in MB.
	IndexLength *float64 `json:"IndexLength,omitempty" name:"IndexLength"`

	// Fragmented space in MB.
	DataFree *float64 `json:"DataFree,omitempty" name:"DataFree"`

	// Total space usage in MB.
	TotalLength *float64 `json:"TotalLength,omitempty" name:"TotalLength"`

	// Fragmentation rate in %.
	FragRatio *float64 `json:"FragRatio,omitempty" name:"FragRatio"`

	// Number of rows.
	TableRows *int64 `json:"TableRows,omitempty" name:"TableRows"`

	// Total size in MB of physical files exclusive to all tables in the database.
	// Note: this field may return null, indicating that no valid values can be obtained.
	PhysicalFileSize *float64 `json:"PhysicalFileSize,omitempty" name:"PhysicalFileSize"`
}

type SchemaSpaceTimeSeries struct {

	// Database name
	TableSchema *string `json:"TableSchema,omitempty" name:"TableSchema"`

	// Space metric value in a unit of time interval
	SeriesData *MonitorMetricSeriesData `json:"SeriesData,omitempty" name:"SeriesData"`
}

type ScoreDetail struct {

	// Deduction item type. Valid values: availability, maintainability, performance, and reliability.
	IssueType *string `json:"IssueType,omitempty" name:"IssueType"`

	// Total deducted scores.
	ScoreLost *int64 `json:"ScoreLost,omitempty" name:"ScoreLost"`

	// Upper limit of the deducted scores.
	ScoreLostMax *int64 `json:"ScoreLostMax,omitempty" name:"ScoreLostMax"`

	// Deduction item list.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Items []*ScoreItem `json:"Items,omitempty" name:"Items"`
}

type ScoreItem struct {

	// Exception diagnosis item name.
	DiagItem *string `json:"DiagItem,omitempty" name:"DiagItem"`

	// Diagnosis item type. Valid values: availability, maintainability, performance, and reliability.
	IssueType *string `json:"IssueType,omitempty" name:"IssueType"`

	// Health level. Valid values: information, reminder, alarm, serious, fatal.
	TopSeverity *string `json:"TopSeverity,omitempty" name:"TopSeverity"`

	// Number of occurrences of this exception diagnosis item.
	Count *int64 `json:"Count,omitempty" name:"Count"`

	// Deducted scores.
	ScoreLost *int64 `json:"ScoreLost,omitempty" name:"ScoreLost"`
}

type SecLogExportTaskInfo struct {

	// Async task Id.
	AsyncRequestId *uint64 `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

	// Task start time.
	// Note: this field may return null, indicating that no valid values can be obtained.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Task end time.
	// Note: this field may return null, indicating that no valid values can be obtained.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Task creation time.
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Task status.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Task progress.
	Progress *uint64 `json:"Progress,omitempty" name:"Progress"`

	// Exported log start time.
	// Note: this field may return null, indicating that no valid values can be obtained.
	LogStartTime *string `json:"LogStartTime,omitempty" name:"LogStartTime"`

	// Exported log end time.
	// Note: this field may return null, indicating that no valid values can be obtained.
	LogEndTime *string `json:"LogEndTime,omitempty" name:"LogEndTime"`

	// Total size of log files in KB.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TotalSize *uint64 `json:"TotalSize,omitempty" name:"TotalSize"`

	// List of risk levels. 0: no risk; 1: low risk; 2: medium risk; 3 high risk.
	// Note: this field may return null, indicating that no valid values can be obtained.
	DangerLevels []*uint64 `json:"DangerLevels,omitempty" name:"DangerLevels"`
}

type SlowLogHost struct {

	// Source addresses.
	UserHost *string `json:"UserHost,omitempty" name:"UserHost"`

	// Proportion (in %) of slow logs from this source address to the total number of slow logs.
	Ratio *float64 `json:"Ratio,omitempty" name:"Ratio"`

	// Number of slow logs from this source address.
	Count *int64 `json:"Count,omitempty" name:"Count"`
}

type SlowLogTopSqlItem struct {

	// Total SQL lock wait time in seconds.
	LockTime *float64 `json:"LockTime,omitempty" name:"LockTime"`

	// Maximum lock wait time in seconds
	LockTimeMax *float64 `json:"LockTimeMax,omitempty" name:"LockTimeMax"`

	// Minimum lock wait time in seconds
	LockTimeMin *float64 `json:"LockTimeMin,omitempty" name:"LockTimeMin"`

	// Total number of scanned rows
	RowsExamined *int64 `json:"RowsExamined,omitempty" name:"RowsExamined"`

	// Maximum number of scanned rows
	RowsExaminedMax *int64 `json:"RowsExaminedMax,omitempty" name:"RowsExaminedMax"`

	// Minimum number of scanned rows
	RowsExaminedMin *int64 `json:"RowsExaminedMin,omitempty" name:"RowsExaminedMin"`

	// Total duration in seconds
	QueryTime *float64 `json:"QueryTime,omitempty" name:"QueryTime"`

	// Maximum execution time in seconds
	QueryTimeMax *float64 `json:"QueryTimeMax,omitempty" name:"QueryTimeMax"`

	// Minimum execution time in seconds
	QueryTimeMin *float64 `json:"QueryTimeMin,omitempty" name:"QueryTimeMin"`

	// Total number of returned rows
	RowsSent *int64 `json:"RowsSent,omitempty" name:"RowsSent"`

	// Maximum number of returned rows
	RowsSentMax *int64 `json:"RowsSentMax,omitempty" name:"RowsSentMax"`

	// Minimum number of returned rows
	RowsSentMin *int64 `json:"RowsSentMin,omitempty" name:"RowsSentMin"`

	// Number of executions
	ExecTimes *int64 `json:"ExecTimes,omitempty" name:"ExecTimes"`

	// SQL template
	SqlTemplate *string `json:"SqlTemplate,omitempty" name:"SqlTemplate"`

	// SQL statements with parameter (random)
	SqlText *string `json:"SqlText,omitempty" name:"SqlText"`

	// Database name
	Schema *string `json:"Schema,omitempty" name:"Schema"`

	// Ratio of total duration in %
	QueryTimeRatio *float64 `json:"QueryTimeRatio,omitempty" name:"QueryTimeRatio"`

	// Ratio of total SQL lock wait time in %
	LockTimeRatio *float64 `json:"LockTimeRatio,omitempty" name:"LockTimeRatio"`

	// Ratio of total number of scanned rows in %
	RowsExaminedRatio *float64 `json:"RowsExaminedRatio,omitempty" name:"RowsExaminedRatio"`

	// Ratio of total number of returned rows in %
	RowsSentRatio *float64 `json:"RowsSentRatio,omitempty" name:"RowsSentRatio"`

	// Average execution time in seconds
	QueryTimeAvg *float64 `json:"QueryTimeAvg,omitempty" name:"QueryTimeAvg"`

	// Average number of returned rows
	RowsSentAvg *float64 `json:"RowsSentAvg,omitempty" name:"RowsSentAvg"`

	// Average lock wait time in seconds
	LockTimeAvg *float64 `json:"LockTimeAvg,omitempty" name:"LockTimeAvg"`

	// Average number of scanned rows
	RowsExaminedAvg *float64 `json:"RowsExaminedAvg,omitempty" name:"RowsExaminedAvg"`

	// MD5 value of SOL template
	Md5 *string `json:"Md5,omitempty" name:"Md5"`
}

type TableSpaceData struct {

	// Table name.
	TableName *string `json:"TableName,omitempty" name:"TableName"`

	// Database name.
	TableSchema *string `json:"TableSchema,omitempty" name:"TableSchema"`

	// Database table storage engine.
	Engine *string `json:"Engine,omitempty" name:"Engine"`

	// Data space in MB.
	DataLength *float64 `json:"DataLength,omitempty" name:"DataLength"`

	// Index space in MB.
	IndexLength *float64 `json:"IndexLength,omitempty" name:"IndexLength"`

	// Fragmented space in MB.
	DataFree *float64 `json:"DataFree,omitempty" name:"DataFree"`

	// Total space usage in MB.
	TotalLength *float64 `json:"TotalLength,omitempty" name:"TotalLength"`

	// Fragmentation rate in %.
	FragRatio *float64 `json:"FragRatio,omitempty" name:"FragRatio"`

	// Number of rows.
	TableRows *int64 `json:"TableRows,omitempty" name:"TableRows"`

	// Size in MB of the physical file exclusive to a table.
	PhysicalFileSize *float64 `json:"PhysicalFileSize,omitempty" name:"PhysicalFileSize"`
}

type TableSpaceTimeSeries struct {

	// Table name.
	TableName *string `json:"TableName,omitempty" name:"TableName"`

	// Database name.
	TableSchema *string `json:"TableSchema,omitempty" name:"TableSchema"`

	// Database table storage engine.
	Engine *string `json:"Engine,omitempty" name:"Engine"`

	// Space metric value in a unit of time interval
	SeriesData *MonitorFloatMetricSeriesData `json:"SeriesData,omitempty" name:"SeriesData"`
}

type TimeSlice struct {

	// Total number
	Count *int64 `json:"Count,omitempty" name:"Count"`

	// Statistics start time
	Timestamp *int64 `json:"Timestamp,omitempty" name:"Timestamp"`
}

type UserProfile struct {

	// Configured ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	ProfileId *string `json:"ProfileId,omitempty" name:"ProfileId"`

	// Configuration type. Valid values: "dbScan_mail_configuration" (email configuration of database inspection report), "scheduler_mail_configuration" (email configuration of scheduled task report).
	// Note: this field may return null, indicating that no valid values can be obtained.
	ProfileType *string `json:"ProfileType,omitempty" name:"ProfileType"`

	// Configuration level. Valid values: User (user-level), Instance (instance-level). For database inspection emails, it should be `User`. For scheduled task emails, it should be `Instance`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ProfileLevel *string `json:"ProfileLevel,omitempty" name:"ProfileLevel"`

	// Configuration name.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ProfileName *string `json:"ProfileName,omitempty" name:"ProfileName"`

	// Configuration details.
	ProfileInfo *ProfileInfo `json:"ProfileInfo,omitempty" name:"ProfileInfo"`
}
