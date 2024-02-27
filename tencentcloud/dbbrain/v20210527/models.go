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
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

// Predefined struct for user
type AddUserContactRequestParams struct {
	// Recipient name, which can contain up to 20 letters, digits, spaces, and symbols `!@#$%^&*()_+-=()` and cannot begin with an underscore.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Email address, which can contain letters, digits, underscores, and the @ symbol, cannot begin with an underscore, and must be unique.
	ContactInfo *string `json:"ContactInfo,omitnil,omitempty" name:"ContactInfo"`

	// Service type, which is fixed to `mysql`.
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type AddUserContactRequest struct {
	*tchttp.BaseRequest
	
	// Recipient name, which can contain up to 20 letters, digits, spaces, and symbols `!@#$%^&*()_+-=()` and cannot begin with an underscore.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Email address, which can contain letters, digits, underscores, and the @ symbol, cannot begin with an underscore, and must be unique.
	ContactInfo *string `json:"ContactInfo,omitnil,omitempty" name:"ContactInfo"`

	// Service type, which is fixed to `mysql`.
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
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

// Predefined struct for user
type AddUserContactResponseParams struct {
	// ID of the successfully added recipient.
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AddUserContactResponse struct {
	*tchttp.BaseResponse
	Response *AddUserContactResponseParams `json:"Response"`
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

type AuditInstance struct {
	// Audit status. Valid values: `ON` (Enabled), `OFF` (Not enabled).
	AuditStatus *string `json:"AuditStatus,omitnil,omitempty" name:"AuditStatus"`

	// Audit log size. This parameter is only used for the free trial edition of Database Audit.
	BillingAmount *int64 `json:"BillingAmount,omitnil,omitempty" name:"BillingAmount"`

	// Billing confirmation status. Valid values: `0` (Unconfirmed), `1` (Confirmed).
	BillingConfirmed *int64 `json:"BillingConfirmed,omitnil,omitempty" name:"BillingConfirmed"`

	// Infrequent access storage period
	ColdLogExpireDay *int64 `json:"ColdLogExpireDay,omitnil,omitempty" name:"ColdLogExpireDay"`

	// Storage size of infrequently accessed logs in MB
	ColdLogSize *int64 `json:"ColdLogSize,omitnil,omitempty" name:"ColdLogSize"`

	// Storage period of frequently accessed logs in days
	HotLogExpireDay *int64 `json:"HotLogExpireDay,omitnil,omitempty" name:"HotLogExpireDay"`

	// Storage size of frequently accessed logs in MB
	HotLogSize *int64 `json:"HotLogSize,omitnil,omitempty" name:"HotLogSize"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Log retention period in days, which is the sum of the frequent and infrequent access storage periods.
	LogExpireDay *int64 `json:"LogExpireDay,omitnil,omitempty" name:"LogExpireDay"`

	// Instance creation time
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Instance details
	InstanceInfo *AuditInstanceInfo `json:"InstanceInfo,omitnil,omitempty" name:"InstanceInfo"`
}

type AuditInstanceFilter struct {
	// Filter name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Filter value
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type AuditInstanceInfo struct {
	// appId
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// Audit status. Valid values: `0` (Not enabled), `1` (Enabled).
	AuditStatus *int64 `json:"AuditStatus,omitnil,omitempty" name:"AuditStatus"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Instance name
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Project ID
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// The region where the instance resides
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Resource tags
	// Note: u200dThis field may returnu200d·nullu200d, indicating that no valid values can be obtained.
	ResourceTags []*string `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`
}

// Predefined struct for user
type CloseAuditServiceRequestParams struct {
	// Service type. Valid values: `dcdb` (TDSQL for MySQL), `mariadb` (TencentDB for MariaDB).
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// Use the value of `u200cProduct` for this parameter, such as `dcdb` and `mariadb`.
	NodeRequestType *string `json:"NodeRequestType,omitnil,omitempty" name:"NodeRequestType"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type CloseAuditServiceRequest struct {
	*tchttp.BaseRequest
	
	// Service type. Valid values: `dcdb` (TDSQL for MySQL), `mariadb` (TencentDB for MariaDB).
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// Use the value of `u200cProduct` for this parameter, such as `dcdb` and `mariadb`.
	NodeRequestType *string `json:"NodeRequestType,omitnil,omitempty" name:"NodeRequestType"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *CloseAuditServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseAuditServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Product")
	delete(f, "NodeRequestType")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CloseAuditServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloseAuditServiceResponseParams struct {
	// If `0` is returned, audit is successfully disabled; otherwise, an exception will be returned, indicating that audit has failed to be disabled.
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CloseAuditServiceResponse struct {
	*tchttp.BaseResponse
	Response *CloseAuditServiceResponseParams `json:"Response"`
}

func (r *CloseAuditServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseAuditServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ContactItem struct {
	// Recipient ID.
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// Recipient name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Recipient email.
	Mail *string `json:"Mail,omitnil,omitempty" name:"Mail"`
}

// Predefined struct for user
type CreateDBDiagReportTaskRequestParams struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Start time, such as "2020-11-08T14:00:00+08:00".
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time, such as "2020-11-09T14:00:00+08:00".
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Whether to send an email. Valid values: `0` (yes), `1` (no).
	SendMailFlag *int64 `json:"SendMailFlag,omitnil,omitempty" name:"SendMailFlag"`

	// Array of the IDs of recipients to receive email.
	ContactPerson []*int64 `json:"ContactPerson,omitnil,omitempty" name:"ContactPerson"`

	// Array of IDs of recipient groups to receive email.
	ContactGroup []*int64 `json:"ContactGroup,omitnil,omitempty" name:"ContactGroup"`

	// Service type. Valid values: `mysql` (TencentDB for MySQL), `cynosdb` (TDSQL-C for MySQL). Default value: `mysql`.
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type CreateDBDiagReportTaskRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Start time, such as "2020-11-08T14:00:00+08:00".
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time, such as "2020-11-09T14:00:00+08:00".
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Whether to send an email. Valid values: `0` (yes), `1` (no).
	SendMailFlag *int64 `json:"SendMailFlag,omitnil,omitempty" name:"SendMailFlag"`

	// Array of the IDs of recipients to receive email.
	ContactPerson []*int64 `json:"ContactPerson,omitnil,omitempty" name:"ContactPerson"`

	// Array of IDs of recipient groups to receive email.
	ContactGroup []*int64 `json:"ContactGroup,omitnil,omitempty" name:"ContactGroup"`

	// Service type. Valid values: `mysql` (TencentDB for MySQL), `cynosdb` (TDSQL-C for MySQL). Default value: `mysql`.
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
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

// Predefined struct for user
type CreateDBDiagReportTaskResponseParams struct {
	// Async task request ID, which can be used to query the execution result of an async task.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AsyncRequestId *int64 `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDBDiagReportTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateDBDiagReportTaskResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateDBDiagReportUrlRequestParams struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Health report task ID, which can be queried through `DescribeDBDiagReportTasks`.
	AsyncRequestId *int64 `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// Service type. Valid values: `mysql` (TencentDB for MySQL), `cynosdb` (TDSQL-C for MySQL). Default value: `mysql`.
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type CreateDBDiagReportUrlRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Health report task ID, which can be queried through `DescribeDBDiagReportTasks`.
	AsyncRequestId *int64 `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// Service type. Valid values: `mysql` (TencentDB for MySQL), `cynosdb` (TDSQL-C for MySQL). Default value: `mysql`.
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
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

// Predefined struct for user
type CreateDBDiagReportUrlResponseParams struct {
	// Health report URL.
	ReportUrl *string `json:"ReportUrl,omitnil,omitempty" name:"ReportUrl"`

	// Expiration timestamp of the health report URL (in seconds).
	ExpireTime *int64 `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDBDiagReportUrlResponse struct {
	*tchttp.BaseResponse
	Response *CreateDBDiagReportUrlResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateKillTaskRequestParams struct {
	// ID of the instance associated with the session killing task.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Task duration in seconds. Pass in `-1` to stop the task manually.
	Duration *int64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// Client IP, which is a task filter.
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// Database name, which is a task filter. Multiple database names are separated by comma.
	DB *string `json:"DB,omitnil,omitempty" name:"DB"`

	// Related command, which is a task filter. Multiple commands are separated by comma.
	Command *string `json:"Command,omitnil,omitempty" name:"Command"`

	// Task filter. Filtering by single filter prefix is supported.
	Info *string `json:"Info,omitnil,omitempty" name:"Info"`

	// User type, which is a task filter.
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// Session duration in seconds, which is a task filter.
	Time *int64 `json:"Time,omitnil,omitempty" name:"Time"`

	// Service type. Valid values: `mysql` (TencentDB for MySQL), `cynosdb` (TDSQL-C for MySQL). Default value: `mysql`.
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type CreateKillTaskRequest struct {
	*tchttp.BaseRequest
	
	// ID of the instance associated with the session killing task.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Task duration in seconds. Pass in `-1` to stop the task manually.
	Duration *int64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// Client IP, which is a task filter.
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// Database name, which is a task filter. Multiple database names are separated by comma.
	DB *string `json:"DB,omitnil,omitempty" name:"DB"`

	// Related command, which is a task filter. Multiple commands are separated by comma.
	Command *string `json:"Command,omitnil,omitempty" name:"Command"`

	// Task filter. Filtering by single filter prefix is supported.
	Info *string `json:"Info,omitnil,omitempty" name:"Info"`

	// User type, which is a task filter.
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// Session duration in seconds, which is a task filter.
	Time *int64 `json:"Time,omitnil,omitempty" name:"Time"`

	// Service type. Valid values: `mysql` (TencentDB for MySQL), `cynosdb` (TDSQL-C for MySQL). Default value: `mysql`.
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

func (r *CreateKillTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateKillTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Duration")
	delete(f, "Host")
	delete(f, "DB")
	delete(f, "Command")
	delete(f, "Info")
	delete(f, "User")
	delete(f, "Time")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateKillTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateKillTaskResponseParams struct {
	// Task status. `1` is returned if the session killing task is successfully created.
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateKillTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateKillTaskResponseParams `json:"Response"`
}

func (r *CreateKillTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateKillTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMailProfileRequestParams struct {
	// Email configuration.
	ProfileInfo *ProfileInfo `json:"ProfileInfo,omitnil,omitempty" name:"ProfileInfo"`

	// Configuration level. Valid values: `User` (user-level), `Instance` (instance-level). For database inspection emails, it should be `User`. For scheduled task emails, it should be `Instance`.
	ProfileLevel *string `json:"ProfileLevel,omitnil,omitempty" name:"ProfileLevel"`

	// Configuration name, which needs to be unique. For database inspection emails, this name can be customized as needed. For scheduled task emails, the name should be in the format of "scheduler_" + {instanceId}, such as "schduler_cdb-test".
	ProfileName *string `json:"ProfileName,omitnil,omitempty" name:"ProfileName"`

	// Configuration type. Valid values: `dbScan_mail_configuration` (email configuration of the database inspection report), `scheduler_mail_configuration` (email configuration of the scheduled task report).
	ProfileType *string `json:"ProfileType,omitnil,omitempty" name:"ProfileType"`

	// Service type. Valid values: `mysql` (TencentDB for MySQL), `cynosdb` (TDSQL-C for MySQL).
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// Instance ID bound with the configuration, which is set when the configuration level is `Instance`. Only one instance can be bound at a time. When the configuration level is `User`, leave this parameter empty.
	BindInstanceIds []*string `json:"BindInstanceIds,omitnil,omitempty" name:"BindInstanceIds"`
}

type CreateMailProfileRequest struct {
	*tchttp.BaseRequest
	
	// Email configuration.
	ProfileInfo *ProfileInfo `json:"ProfileInfo,omitnil,omitempty" name:"ProfileInfo"`

	// Configuration level. Valid values: `User` (user-level), `Instance` (instance-level). For database inspection emails, it should be `User`. For scheduled task emails, it should be `Instance`.
	ProfileLevel *string `json:"ProfileLevel,omitnil,omitempty" name:"ProfileLevel"`

	// Configuration name, which needs to be unique. For database inspection emails, this name can be customized as needed. For scheduled task emails, the name should be in the format of "scheduler_" + {instanceId}, such as "schduler_cdb-test".
	ProfileName *string `json:"ProfileName,omitnil,omitempty" name:"ProfileName"`

	// Configuration type. Valid values: `dbScan_mail_configuration` (email configuration of the database inspection report), `scheduler_mail_configuration` (email configuration of the scheduled task report).
	ProfileType *string `json:"ProfileType,omitnil,omitempty" name:"ProfileType"`

	// Service type. Valid values: `mysql` (TencentDB for MySQL), `cynosdb` (TDSQL-C for MySQL).
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// Instance ID bound with the configuration, which is set when the configuration level is `Instance`. Only one instance can be bound at a time. When the configuration level is `User`, leave this parameter empty.
	BindInstanceIds []*string `json:"BindInstanceIds,omitnil,omitempty" name:"BindInstanceIds"`
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

// Predefined struct for user
type CreateMailProfileResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateMailProfileResponse struct {
	*tchttp.BaseResponse
	Response *CreateMailProfileResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateProxySessionKillTaskRequestParams struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Service type. Valid value: `redis` (TencentDB for Redis).
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type CreateProxySessionKillTaskRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Service type. Valid value: `redis` (TencentDB for Redis).
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

func (r *CreateProxySessionKillTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProxySessionKillTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateProxySessionKillTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateProxySessionKillTaskResponseParams struct {
	// Async task ID that is returned after the session killing task is created.
	AsyncRequestId *int64 `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateProxySessionKillTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateProxySessionKillTaskResponseParams `json:"Response"`
}

func (r *CreateProxySessionKillTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProxySessionKillTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRedisBigKeyAnalysisTaskRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Service type. Valid value: `redis` (TencentDB for Redis).
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// The list of the serial numbers of shard nodes. When the list is empty, all shard nodes will be selected.
	ShardIds []*int64 `json:"ShardIds,omitnil,omitempty" name:"ShardIds"`

	// The list of separators of top key prefixes.
	// Currently, the following separators are supported: ",", ";", ":", "_", "-", "+", "@", "=", "|", "#", ".". When the list is empty, all separators will be selected by default.
	KeyDelimiterList []*string `json:"KeyDelimiterList,omitnil,omitempty" name:"KeyDelimiterList"`
}

type CreateRedisBigKeyAnalysisTaskRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Service type. Valid value: `redis` (TencentDB for Redis).
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// The list of the serial numbers of shard nodes. When the list is empty, all shard nodes will be selected.
	ShardIds []*int64 `json:"ShardIds,omitnil,omitempty" name:"ShardIds"`

	// The list of separators of top key prefixes.
	// Currently, the following separators are supported: ",", ";", ":", "_", "-", "+", "@", "=", "|", "#", ".". When the list is empty, all separators will be selected by default.
	KeyDelimiterList []*string `json:"KeyDelimiterList,omitnil,omitempty" name:"KeyDelimiterList"`
}

func (r *CreateRedisBigKeyAnalysisTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRedisBigKeyAnalysisTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Product")
	delete(f, "ShardIds")
	delete(f, "KeyDelimiterList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRedisBigKeyAnalysisTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRedisBigKeyAnalysisTaskResponseParams struct {
	// Async task ID
	AsyncRequestId *int64 `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateRedisBigKeyAnalysisTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateRedisBigKeyAnalysisTaskResponseParams `json:"Response"`
}

func (r *CreateRedisBigKeyAnalysisTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRedisBigKeyAnalysisTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSchedulerMailProfileRequestParams struct {
	// Value range: 1-7, representing Monday to Sunday respectively.
	WeekConfiguration []*int64 `json:"WeekConfiguration,omitnil,omitempty" name:"WeekConfiguration"`

	// Email configuration.
	ProfileInfo *ProfileInfo `json:"ProfileInfo,omitnil,omitempty" name:"ProfileInfo"`

	// Configuration name, which needs to be unique. For scheduled task emails, the name should be in the format of "scheduler_" + {instanceId}, such as "schduler_cdb-test".
	ProfileName *string `json:"ProfileName,omitnil,omitempty" name:"ProfileName"`

	// ID of the instance for which to configure subscription.
	BindInstanceId *string `json:"BindInstanceId,omitnil,omitempty" name:"BindInstanceId"`

	// Service type. Valid values: `mysql` (TencentDB for MySQL), `cynosdb` (TDSQL-C for MySQL). Default value: `mysql`.
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type CreateSchedulerMailProfileRequest struct {
	*tchttp.BaseRequest
	
	// Value range: 1-7, representing Monday to Sunday respectively.
	WeekConfiguration []*int64 `json:"WeekConfiguration,omitnil,omitempty" name:"WeekConfiguration"`

	// Email configuration.
	ProfileInfo *ProfileInfo `json:"ProfileInfo,omitnil,omitempty" name:"ProfileInfo"`

	// Configuration name, which needs to be unique. For scheduled task emails, the name should be in the format of "scheduler_" + {instanceId}, such as "schduler_cdb-test".
	ProfileName *string `json:"ProfileName,omitnil,omitempty" name:"ProfileName"`

	// ID of the instance for which to configure subscription.
	BindInstanceId *string `json:"BindInstanceId,omitnil,omitempty" name:"BindInstanceId"`

	// Service type. Valid values: `mysql` (TencentDB for MySQL), `cynosdb` (TDSQL-C for MySQL). Default value: `mysql`.
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
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

// Predefined struct for user
type CreateSchedulerMailProfileResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateSchedulerMailProfileResponse struct {
	*tchttp.BaseResponse
	Response *CreateSchedulerMailProfileResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateSecurityAuditLogExportTaskRequestParams struct {
	// Security audit group ID.
	SecAuditGroupId *string `json:"SecAuditGroupId,omitnil,omitempty" name:"SecAuditGroupId"`

	// Exported log start time, such as 2020-12-28 00:00:00.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Exported log end time, such as 2020-12-28 01:00:00.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Service type. Valid value: `mysql` (TencentDB for MySQL).
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// List of log risk levels. Valid values: `0` (no risk), `1` (low risk), `2` (medium risk), `3` (high risk).
	DangerLevels []*int64 `json:"DangerLevels,omitnil,omitempty" name:"DangerLevels"`
}

type CreateSecurityAuditLogExportTaskRequest struct {
	*tchttp.BaseRequest
	
	// Security audit group ID.
	SecAuditGroupId *string `json:"SecAuditGroupId,omitnil,omitempty" name:"SecAuditGroupId"`

	// Exported log start time, such as 2020-12-28 00:00:00.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Exported log end time, such as 2020-12-28 01:00:00.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Service type. Valid value: `mysql` (TencentDB for MySQL).
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// List of log risk levels. Valid values: `0` (no risk), `1` (low risk), `2` (medium risk), `3` (high risk).
	DangerLevels []*int64 `json:"DangerLevels,omitnil,omitempty" name:"DangerLevels"`
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

// Predefined struct for user
type CreateSecurityAuditLogExportTaskResponseParams struct {
	// Log export task Id.
	AsyncRequestId *uint64 `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateSecurityAuditLogExportTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateSecurityAuditLogExportTaskResponseParams `json:"Response"`
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

// Predefined struct for user
type DeleteDBDiagReportTasksRequestParams struct {
	// List of IDs of tasks to be deleted
	AsyncRequestIds []*int64 `json:"AsyncRequestIds,omitnil,omitempty" name:"AsyncRequestIds"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Service type. Valid values: `mysql` (TencentDB for MySQL), `cynosdb` (TDSQL-C for MySQL). Default value: `mysql`.
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type DeleteDBDiagReportTasksRequest struct {
	*tchttp.BaseRequest
	
	// List of IDs of tasks to be deleted
	AsyncRequestIds []*int64 `json:"AsyncRequestIds,omitnil,omitempty" name:"AsyncRequestIds"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Service type. Valid values: `mysql` (TencentDB for MySQL), `cynosdb` (TDSQL-C for MySQL). Default value: `mysql`.
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

func (r *DeleteDBDiagReportTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDBDiagReportTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AsyncRequestIds")
	delete(f, "InstanceId")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDBDiagReportTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDBDiagReportTasksResponseParams struct {
	// Task deletion status (`0`: Successful)
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteDBDiagReportTasksResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDBDiagReportTasksResponseParams `json:"Response"`
}

func (r *DeleteDBDiagReportTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDBDiagReportTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSecurityAuditLogExportTasksRequestParams struct {
	// Security audit group ID.
	SecAuditGroupId *string `json:"SecAuditGroupId,omitnil,omitempty" name:"SecAuditGroupId"`

	// List of log export task IDs. This API will ignore task IDs that do not exist or have been deleted.
	AsyncRequestIds []*uint64 `json:"AsyncRequestIds,omitnil,omitempty" name:"AsyncRequestIds"`

	// Service type. Valid value: `mysql` (TencentDB for MySQL).
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type DeleteSecurityAuditLogExportTasksRequest struct {
	*tchttp.BaseRequest
	
	// Security audit group ID.
	SecAuditGroupId *string `json:"SecAuditGroupId,omitnil,omitempty" name:"SecAuditGroupId"`

	// List of log export task IDs. This API will ignore task IDs that do not exist or have been deleted.
	AsyncRequestIds []*uint64 `json:"AsyncRequestIds,omitnil,omitempty" name:"AsyncRequestIds"`

	// Service type. Valid value: `mysql` (TencentDB for MySQL).
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
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

// Predefined struct for user
type DeleteSecurityAuditLogExportTasksResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteSecurityAuditLogExportTasksResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSecurityAuditLogExportTasksResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeAllUserContactRequestParams struct {
	// Service type, which is fixed to `mysql`.
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// Array of recipient names. Fuzzy search is supported.
	Names []*string `json:"Names,omitnil,omitempty" name:"Names"`
}

type DescribeAllUserContactRequest struct {
	*tchttp.BaseRequest
	
	// Service type, which is fixed to `mysql`.
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// Array of recipient names. Fuzzy search is supported.
	Names []*string `json:"Names,omitnil,omitempty" name:"Names"`
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

// Predefined struct for user
type DescribeAllUserContactResponseParams struct {
	// Total number of recipients.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Recipient information.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Contacts []*ContactItem `json:"Contacts,omitnil,omitempty" name:"Contacts"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAllUserContactResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAllUserContactResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeAllUserGroupRequestParams struct {
	// Service type, which is fixed to `mysql`.
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// Array of recipient group names. Fuzzy search is supported.
	Names []*string `json:"Names,omitnil,omitempty" name:"Names"`
}

type DescribeAllUserGroupRequest struct {
	*tchttp.BaseRequest
	
	// Service type, which is fixed to `mysql`.
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// Array of recipient group names. Fuzzy search is supported.
	Names []*string `json:"Names,omitnil,omitempty" name:"Names"`
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

// Predefined struct for user
type DescribeAllUserGroupResponseParams struct {
	// Total number of groups.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Group information.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Groups []*GroupItem `json:"Groups,omitnil,omitempty" name:"Groups"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAllUserGroupResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAllUserGroupResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeAuditInstanceListRequestParams struct {
	// Service type. Valid values: `dcdb` (TDSQL for MySQL), `mariadb` (TencentDB for MariaDB).
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// Use the value of `u200cProduct` for this parameter, such as `dcdb` and `mariadb`.
	NodeRequestType *string `json:"NodeRequestType,omitnil,omitempty" name:"NodeRequestType"`

	// Audit status. Valid values: `0` (Not enabled), `1` (Enabled). Default value: `0`.
	AuditSwitch *int64 `json:"AuditSwitch,omitnil,omitempty" name:"AuditSwitch"`

	// The offset. Default value: `0`.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The number of queried items. Default value: `20`. Max value: `100`.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Filters for querying instances
	Filters []*AuditInstanceFilter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeAuditInstanceListRequest struct {
	*tchttp.BaseRequest
	
	// Service type. Valid values: `dcdb` (TDSQL for MySQL), `mariadb` (TencentDB for MariaDB).
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// Use the value of `u200cProduct` for this parameter, such as `dcdb` and `mariadb`.
	NodeRequestType *string `json:"NodeRequestType,omitnil,omitempty" name:"NodeRequestType"`

	// Audit status. Valid values: `0` (Not enabled), `1` (Enabled). Default value: `0`.
	AuditSwitch *int64 `json:"AuditSwitch,omitnil,omitempty" name:"AuditSwitch"`

	// The offset. Default value: `0`.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The number of queried items. Default value: `20`. Max value: `100`.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Filters for querying instances
	Filters []*AuditInstanceFilter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeAuditInstanceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuditInstanceListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Product")
	delete(f, "NodeRequestType")
	delete(f, "AuditSwitch")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAuditInstanceListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAuditInstanceListResponseParams struct {
	// The number of eligible instances.
	// Note: u200dThis field may return `null`, indicating that no valid values can be obtained.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Instance details
	Items []*AuditInstance `json:"Items,omitnil,omitempty" name:"Items"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAuditInstanceListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAuditInstanceListResponseParams `json:"Response"`
}

func (r *DescribeAuditInstanceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuditInstanceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBDiagEventRequestParams struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Event ID, which can be obtained through the `DescribeDBDiagHistory` API.
	EventId *int64 `json:"EventId,omitnil,omitempty" name:"EventId"`

	// Service type. Valid values: `mysql` (TencentDB for MySQL), `cynosdb` (TDSQL-C for MySQL). Default value: `mysql`.
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type DescribeDBDiagEventRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Event ID, which can be obtained through the `DescribeDBDiagHistory` API.
	EventId *int64 `json:"EventId,omitnil,omitempty" name:"EventId"`

	// Service type. Valid values: `mysql` (TencentDB for MySQL), `cynosdb` (TDSQL-C for MySQL). Default value: `mysql`.
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
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

// Predefined struct for user
type DescribeDBDiagEventResponseParams struct {
	// Diagnosis item.
	DiagItem *string `json:"DiagItem,omitnil,omitempty" name:"DiagItem"`

	// Diagnosis type.
	DiagType *string `json:"DiagType,omitnil,omitempty" name:"DiagType"`

	// Event ID.
	EventId *int64 `json:"EventId,omitnil,omitempty" name:"EventId"`

	// Diagnosis event details. If there is no additional explanation information, the output will be empty.
	Explanation *string `json:"Explanation,omitnil,omitempty" name:"Explanation"`

	// Diagnosis summary.
	Outline *string `json:"Outline,omitnil,omitempty" name:"Outline"`

	// Found problem.
	Problem *string `json:"Problem,omitnil,omitempty" name:"Problem"`

	// Severity, which can be divided into 5 levels: `1` (Critical), `2` (Severe), `3` (Alarm), `4` (Reminder), `5` (healthy).
	Severity *int64 `json:"Severity,omitnil,omitempty" name:"Severity"`

	// Start time
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Suggestions. If there are no suggestions, the output will be empty.
	Suggestions *string `json:"Suggestions,omitnil,omitempty" name:"Suggestions"`

	// Reserved field.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Metric *string `json:"Metric,omitnil,omitempty" name:"Metric"`

	// End time.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBDiagEventResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBDiagEventResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeDBDiagEventsRequestParams struct {
	// Start time in the format of “2021-05-27 00:00:00”. The earliest time that can be queried is 30 days before the current time.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time in the format of "2021-05-27 01:00:00". The interval between the end time and the start time can be up to 7 days.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Risk level list. Valid values in descending order of severity: `1` (critical), `2` (serious), `3` (alarm), `4` (warning), `5` (healthy).
	Severities []*int64 `json:"Severities,omitnil,omitempty" name:"Severities"`

	// Instance ID list.
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 50.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeDBDiagEventsRequest struct {
	*tchttp.BaseRequest
	
	// Start time in the format of “2021-05-27 00:00:00”. The earliest time that can be queried is 30 days before the current time.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time in the format of "2021-05-27 01:00:00". The interval between the end time and the start time can be up to 7 days.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Risk level list. Valid values in descending order of severity: `1` (critical), `2` (serious), `3` (alarm), `4` (warning), `5` (healthy).
	Severities []*int64 `json:"Severities,omitnil,omitempty" name:"Severities"`

	// Instance ID list.
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 50.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeDBDiagEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBDiagEventsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Severities")
	delete(f, "InstanceIds")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBDiagEventsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBDiagEventsResponseParams struct {
	// Total number of diagnosis events.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Diagnosis event list.
	Items []*DiagHistoryEventItem `json:"Items,omitnil,omitempty" name:"Items"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBDiagEventsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBDiagEventsResponseParams `json:"Response"`
}

func (r *DescribeDBDiagEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBDiagEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBDiagHistoryRequestParams struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Start time, such as "2019-09-10 12:13:14".
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time, such as "2019-09-11 12:13:14". The interval between the end time and the start time can be up to 2 days.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Service type. Valid values: `mysql` (TencentDB for MySQL), `cynosdb` (TDSQL-C for MySQL). Default value: `mysql`.
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type DescribeDBDiagHistoryRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Start time, such as "2019-09-10 12:13:14".
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time, such as "2019-09-11 12:13:14". The interval between the end time and the start time can be up to 2 days.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Service type. Valid values: `mysql` (TencentDB for MySQL), `cynosdb` (TDSQL-C for MySQL). Default value: `mysql`.
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
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

// Predefined struct for user
type DescribeDBDiagHistoryResponseParams struct {
	// Event description.
	Events []*DiagHistoryEventItem `json:"Events,omitnil,omitempty" name:"Events"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBDiagHistoryResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBDiagHistoryResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeDBDiagReportTasksRequestParams struct {
	// Start time of the first task in the format of yyyy-MM-dd HH:mm:ss, such as 2019-09-10 12:13:14. It is used for queries by time range.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time of the last task in the format of yyyy-MM-dd HH:mm:ss, such as 2019-09-10 12:13:14. It is used for queries by time range.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Array of instance IDs, which is used to filter the task list of the specified instance.
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// Source that triggers the task. Valid values: `DAILY_INSPECTION` (instance inspection), `SCHEDULED` (scheduled task), and `MANUAL` (manual trigger).
	Sources []*string `json:"Sources,omitnil,omitempty" name:"Sources"`

	// Health level. Valid values: `HEALTH` (healthy), `SUB_HEALTH` (suboptimal), `RISK` (risky), and `HIGH_RISK` (critical).
	HealthLevels *string `json:"HealthLevels,omitnil,omitempty" name:"HealthLevels"`

	// Task status. Valid values: `created` (created), `chosen` (to be executed), `running` (being executed), `failed` (failed), and `finished` (completed).
	TaskStatuses *string `json:"TaskStatuses,omitnil,omitempty" name:"TaskStatuses"`

	// Offset. Default value: `0`.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned results. Default value: `20`. Maximum value: `100`.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Service type. Valid values: `mysql` (TencentDB for MySQL), `cynosdb` (TDSQL-C for MySQL). Default value: `mysql`.
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type DescribeDBDiagReportTasksRequest struct {
	*tchttp.BaseRequest
	
	// Start time of the first task in the format of yyyy-MM-dd HH:mm:ss, such as 2019-09-10 12:13:14. It is used for queries by time range.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time of the last task in the format of yyyy-MM-dd HH:mm:ss, such as 2019-09-10 12:13:14. It is used for queries by time range.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Array of instance IDs, which is used to filter the task list of the specified instance.
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// Source that triggers the task. Valid values: `DAILY_INSPECTION` (instance inspection), `SCHEDULED` (scheduled task), and `MANUAL` (manual trigger).
	Sources []*string `json:"Sources,omitnil,omitempty" name:"Sources"`

	// Health level. Valid values: `HEALTH` (healthy), `SUB_HEALTH` (suboptimal), `RISK` (risky), and `HIGH_RISK` (critical).
	HealthLevels *string `json:"HealthLevels,omitnil,omitempty" name:"HealthLevels"`

	// Task status. Valid values: `created` (created), `chosen` (to be executed), `running` (being executed), `failed` (failed), and `finished` (completed).
	TaskStatuses *string `json:"TaskStatuses,omitnil,omitempty" name:"TaskStatuses"`

	// Offset. Default value: `0`.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned results. Default value: `20`. Maximum value: `100`.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Service type. Valid values: `mysql` (TencentDB for MySQL), `cynosdb` (TDSQL-C for MySQL). Default value: `mysql`.
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
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

// Predefined struct for user
type DescribeDBDiagReportTasksResponseParams struct {
	// Total number of tasks.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List of tasks.
	Tasks []*HealthReportTask `json:"Tasks,omitnil,omitempty" name:"Tasks"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBDiagReportTasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBDiagReportTasksResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeDBSpaceStatusRequestParams struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Query period in days. The end date is the current date, and the query period is 7 days by default.
	RangeDays *int64 `json:"RangeDays,omitnil,omitempty" name:"RangeDays"`

	// Service type. Valid values: `mysql` (TencentDB for MySQL), `cynosdb` (TDSQL-C for MySQL). Default value: `mysql`.
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type DescribeDBSpaceStatusRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Query period in days. The end date is the current date, and the query period is 7 days by default.
	RangeDays *int64 `json:"RangeDays,omitnil,omitempty" name:"RangeDays"`

	// Service type. Valid values: `mysql` (TencentDB for MySQL), `cynosdb` (TDSQL-C for MySQL). Default value: `mysql`.
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
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

// Predefined struct for user
type DescribeDBSpaceStatusResponseParams struct {
	// Disk usage growth in MB.
	Growth *int64 `json:"Growth,omitnil,omitempty" name:"Growth"`

	// Available disk space in MB.
	Remain *int64 `json:"Remain,omitnil,omitempty" name:"Remain"`

	// Total disk space in MB.
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// Estimated number of available days.
	AvailableDays *int64 `json:"AvailableDays,omitnil,omitempty" name:"AvailableDays"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBSpaceStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBSpaceStatusResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeDiagDBInstancesRequestParams struct {
	// Whether it is an instance supported by DBbrain. It is fixed to `true`.
	IsSupported *bool `json:"IsSupported,omitnil,omitempty" name:"IsSupported"`

	// Service type. Valid values: mysql (TencentDB for MySQL), cynosdb (TDSQL-C for MySQL). Default value: mysql.
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// Pagination parameter indicating the offset.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Pagination parameter. Maximum value: 100.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Query by instance name.
	InstanceNames []*string `json:"InstanceNames,omitnil,omitempty" name:"InstanceNames"`

	// Query by instance ID.
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// Query by region.
	Regions []*string `json:"Regions,omitnil,omitempty" name:"Regions"`
}

type DescribeDiagDBInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Whether it is an instance supported by DBbrain. It is fixed to `true`.
	IsSupported *bool `json:"IsSupported,omitnil,omitempty" name:"IsSupported"`

	// Service type. Valid values: mysql (TencentDB for MySQL), cynosdb (TDSQL-C for MySQL). Default value: mysql.
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// Pagination parameter indicating the offset.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Pagination parameter. Maximum value: 100.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Query by instance name.
	InstanceNames []*string `json:"InstanceNames,omitnil,omitempty" name:"InstanceNames"`

	// Query by instance ID.
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// Query by region.
	Regions []*string `json:"Regions,omitnil,omitempty" name:"Regions"`
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

// Predefined struct for user
type DescribeDiagDBInstancesResponseParams struct {
	// Total number of instances.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Status of all instance inspection. 0: all instance inspection enabled, 1: all instance inspection disabled.
	DbScanStatus *int64 `json:"DbScanStatus,omitnil,omitempty" name:"DbScanStatus"`

	// Instance information.
	Items []*InstanceInfo `json:"Items,omitnil,omitempty" name:"Items"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDiagDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDiagDBInstancesResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeHealthScoreRequestParams struct {
	// Instance ID for which to get the health score.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Time to get the health score in the format of `2019-09-10 12:13:14`.
	Time *string `json:"Time,omitnil,omitempty" name:"Time"`

	// Service type. Valid values: `mysql` (TencentDB for MySQL), `cynosdb` (TDSQL-C for MySQL). Default value: `mysql`.
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type DescribeHealthScoreRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID for which to get the health score.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Time to get the health score in the format of `2019-09-10 12:13:14`.
	Time *string `json:"Time,omitnil,omitempty" name:"Time"`

	// Service type. Valid values: `mysql` (TencentDB for MySQL), `cynosdb` (TDSQL-C for MySQL). Default value: `mysql`.
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
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

// Predefined struct for user
type DescribeHealthScoreResponseParams struct {
	// Health score and deduction for exceptions.
	Data *HealthScoreInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeHealthScoreResponse struct {
	*tchttp.BaseResponse
	Response *DescribeHealthScoreResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeMailProfileRequestParams struct {
	// Configuration type. Valid values: `dbScan_mail_configuration` (email configuration of the database inspection report), `scheduler_mail_configuration` (email configuration of the scheduled task report).
	ProfileType *string `json:"ProfileType,omitnil,omitempty" name:"ProfileType"`

	// Service type. Valid values: `mysql` (TencentDB for MySQL), `cynosdb` (TDSQL-C for MySQL). Default value: `mysql`.
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// Pagination offset.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of results per page in paginated queries. Maximum value: `50`.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Query by email configuration name. The name of the scheduled task email configuration should be in the format of "scheduler_"+{instanceId}.
	ProfileName *string `json:"ProfileName,omitnil,omitempty" name:"ProfileName"`
}

type DescribeMailProfileRequest struct {
	*tchttp.BaseRequest
	
	// Configuration type. Valid values: `dbScan_mail_configuration` (email configuration of the database inspection report), `scheduler_mail_configuration` (email configuration of the scheduled task report).
	ProfileType *string `json:"ProfileType,omitnil,omitempty" name:"ProfileType"`

	// Service type. Valid values: `mysql` (TencentDB for MySQL), `cynosdb` (TDSQL-C for MySQL). Default value: `mysql`.
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// Pagination offset.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of results per page in paginated queries. Maximum value: `50`.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Query by email configuration name. The name of the scheduled task email configuration should be in the format of "scheduler_"+{instanceId}.
	ProfileName *string `json:"ProfileName,omitnil,omitempty" name:"ProfileName"`
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

// Predefined struct for user
type DescribeMailProfileResponseParams struct {
	// Email configuration details.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ProfileList []*UserProfile `json:"ProfileList,omitnil,omitempty" name:"ProfileList"`

	// Total number of the configured emails.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMailProfileResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMailProfileResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeMySqlProcessListRequestParams struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Thread ID, which is used to filter the thread list.
	ID *uint64 `json:"ID,omitnil,omitempty" name:"ID"`

	// Thread operation account name, which is used to filter the thread list.
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// Thread operation host address, which is used to filter the thread list.
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// Thread operation database, which is used to filter the thread list.
	DB *string `json:"DB,omitnil,omitempty" name:"DB"`

	// Thread operation status, which is used to filter the thread list.
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// Thread execution type, which is used to filter the thread list.
	Command *string `json:"Command,omitnil,omitempty" name:"Command"`

	// Minimum operation duration of the thread in seconds, which is used to filter the list of threads whose operation duration is greater than this value.
	Time *uint64 `json:"Time,omitnil,omitempty" name:"Time"`

	// Thread operation statement, which is used to filter the thread list.
	Info *string `json:"Info,omitnil,omitempty" name:"Info"`

	// Number of returned results. Default value: 20.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Service type. Valid values: mysql (TencentDB for MySQL), cynosdb (TDSQL-C for MySQL). Default value: mysql.
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type DescribeMySqlProcessListRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Thread ID, which is used to filter the thread list.
	ID *uint64 `json:"ID,omitnil,omitempty" name:"ID"`

	// Thread operation account name, which is used to filter the thread list.
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// Thread operation host address, which is used to filter the thread list.
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// Thread operation database, which is used to filter the thread list.
	DB *string `json:"DB,omitnil,omitempty" name:"DB"`

	// Thread operation status, which is used to filter the thread list.
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// Thread execution type, which is used to filter the thread list.
	Command *string `json:"Command,omitnil,omitempty" name:"Command"`

	// Minimum operation duration of the thread in seconds, which is used to filter the list of threads whose operation duration is greater than this value.
	Time *uint64 `json:"Time,omitnil,omitempty" name:"Time"`

	// Thread operation statement, which is used to filter the thread list.
	Info *string `json:"Info,omitnil,omitempty" name:"Info"`

	// Number of returned results. Default value: 20.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Service type. Valid values: mysql (TencentDB for MySQL), cynosdb (TDSQL-C for MySQL). Default value: mysql.
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
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

// Predefined struct for user
type DescribeMySqlProcessListResponseParams struct {
	// List of real-time threads.
	ProcessList []*MySqlProcess `json:"ProcessList,omitnil,omitempty" name:"ProcessList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMySqlProcessListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMySqlProcessListResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeProxyProcessStatisticsRequestParams struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// The proxy ID you want to query under the instance
	InstanceProxyId *string `json:"InstanceProxyId,omitnil,omitempty" name:"InstanceProxyId"`

	// Number of returned results.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Service type. Valid value: `redis` (TencentDB for Redis).
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// Offset. Default value: `0`.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Sort by field. Valid values: `AllConn`, `ActiveConn`, `Ip`.
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// Sorting order. Valid values: `DESC`, `ASC`.
	OrderDirection *string `json:"OrderDirection,omitnil,omitempty" name:"OrderDirection"`
}

type DescribeProxyProcessStatisticsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// The proxy ID you want to query under the instance
	InstanceProxyId *string `json:"InstanceProxyId,omitnil,omitempty" name:"InstanceProxyId"`

	// Number of returned results.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Service type. Valid value: `redis` (TencentDB for Redis).
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// Offset. Default value: `0`.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Sort by field. Valid values: `AllConn`, `ActiveConn`, `Ip`.
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// Sorting order. Valid values: `DESC`, `ASC`.
	OrderDirection *string `json:"OrderDirection,omitnil,omitempty" name:"OrderDirection"`
}

func (r *DescribeProxyProcessStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProxyProcessStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "InstanceProxyId")
	delete(f, "Limit")
	delete(f, "Product")
	delete(f, "Offset")
	delete(f, "SortBy")
	delete(f, "OrderDirection")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProxyProcessStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProxyProcessStatisticsResponseParams struct {
	// Real-time session statistics.
	ProcessStatistics *ProcessStatistic `json:"ProcessStatistics,omitnil,omitempty" name:"ProcessStatistics"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeProxyProcessStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProxyProcessStatisticsResponseParams `json:"Response"`
}

func (r *DescribeProxyProcessStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProxyProcessStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProxySessionKillTasksRequestParams struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// The async session killing task ID, which is obtained after the API `CreateProxySessionKillTask` is successfully called.
	AsyncRequestIds []*int64 `json:"AsyncRequestIds,omitnil,omitempty" name:"AsyncRequestIds"`

	// Service type. Valid value: `redis` (TencentDB for Redis).
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type DescribeProxySessionKillTasksRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// The async session killing task ID, which is obtained after the API `CreateProxySessionKillTask` is successfully called.
	AsyncRequestIds []*int64 `json:"AsyncRequestIds,omitnil,omitempty" name:"AsyncRequestIds"`

	// Service type. Valid value: `redis` (TencentDB for Redis).
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

func (r *DescribeProxySessionKillTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProxySessionKillTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "AsyncRequestIds")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProxySessionKillTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProxySessionKillTasksResponseParams struct {
	// Session killing task details.
	Tasks []*TaskInfo `json:"Tasks,omitnil,omitempty" name:"Tasks"`

	// Total number of tasks.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeProxySessionKillTasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProxySessionKillTasksResponseParams `json:"Response"`
}

func (r *DescribeProxySessionKillTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProxySessionKillTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRedisTopKeyPrefixListRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Date for query, such as `2021-05-27`. You can select a date as early as in the last 30 days for query.
	Date *string `json:"Date,omitnil,omitempty" name:"Date"`

	// Service type. Valid value: `redis` (TencentDB for Redis).
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// The number of queried items. Default value: `20`. Max value: `100`.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeRedisTopKeyPrefixListRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Date for query, such as `2021-05-27`. You can select a date as early as in the last 30 days for query.
	Date *string `json:"Date,omitnil,omitempty" name:"Date"`

	// Service type. Valid value: `redis` (TencentDB for Redis).
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// The number of queried items. Default value: `20`. Max value: `100`.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeRedisTopKeyPrefixListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRedisTopKeyPrefixListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Date")
	delete(f, "Product")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRedisTopKeyPrefixListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRedisTopKeyPrefixListResponseParams struct {
	// List of top key prefixes
	Items []*RedisPreKeySpaceData `json:"Items,omitnil,omitempty" name:"Items"`

	// Data collection timestamp in seconds
	Timestamp *int64 `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRedisTopKeyPrefixListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRedisTopKeyPrefixListResponseParams `json:"Response"`
}

func (r *DescribeRedisTopKeyPrefixListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRedisTopKeyPrefixListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityAuditLogDownloadUrlsRequestParams struct {
	// Security audit group ID.
	SecAuditGroupId *string `json:"SecAuditGroupId,omitnil,omitempty" name:"SecAuditGroupId"`

	// Async task Id.
	AsyncRequestId *uint64 `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// Service type. Valid value: `mysql` (TencentDB for MySQL).
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type DescribeSecurityAuditLogDownloadUrlsRequest struct {
	*tchttp.BaseRequest
	
	// Security audit group ID.
	SecAuditGroupId *string `json:"SecAuditGroupId,omitnil,omitempty" name:"SecAuditGroupId"`

	// Async task Id.
	AsyncRequestId *uint64 `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// Service type. Valid value: `mysql` (TencentDB for MySQL).
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
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

// Predefined struct for user
type DescribeSecurityAuditLogDownloadUrlsResponseParams struct {
	// List of COS URLs of the export results. If the result set is large, it may be divided into multiple URLs for download.
	Urls []*string `json:"Urls,omitnil,omitempty" name:"Urls"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSecurityAuditLogDownloadUrlsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSecurityAuditLogDownloadUrlsResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeSecurityAuditLogExportTasksRequestParams struct {
	// Security audit group ID.
	SecAuditGroupId *string `json:"SecAuditGroupId,omitnil,omitempty" name:"SecAuditGroupId"`

	// Service type. Valid value: `mysql` (TencentDB for MySQL).
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// List of log export task IDs.
	AsyncRequestIds []*uint64 `json:"AsyncRequestIds,omitnil,omitempty" name:"AsyncRequestIds"`

	// Offset. Default value: `0`.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned results. Default value: `20`. Maximum value: `100`.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeSecurityAuditLogExportTasksRequest struct {
	*tchttp.BaseRequest
	
	// Security audit group ID.
	SecAuditGroupId *string `json:"SecAuditGroupId,omitnil,omitempty" name:"SecAuditGroupId"`

	// Service type. Valid value: `mysql` (TencentDB for MySQL).
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// List of log export task IDs.
	AsyncRequestIds []*uint64 `json:"AsyncRequestIds,omitnil,omitempty" name:"AsyncRequestIds"`

	// Offset. Default value: `0`.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned results. Default value: `20`. Maximum value: `100`.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
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

// Predefined struct for user
type DescribeSecurityAuditLogExportTasksResponseParams struct {
	// List of security audit log export tasks.
	Tasks []*SecLogExportTaskInfo `json:"Tasks,omitnil,omitempty" name:"Tasks"`

	// Total numbers of security audit log export tasks.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSecurityAuditLogExportTasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSecurityAuditLogExportTasksResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeSlowLogTimeSeriesStatsRequestParams struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Start time, such as "2019-09-10 12:13:14".
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time, such as "2019-09-10 12:13:14". The interval between the end time and the start time can be up to 7 days.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Service type. Valid values: `mysql` (TencentDB for MySQL), `cynosdb` (TDSQL-C for MySQL). Default value: `mysql`.
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type DescribeSlowLogTimeSeriesStatsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Start time, such as "2019-09-10 12:13:14".
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time, such as "2019-09-10 12:13:14". The interval between the end time and the start time can be up to 7 days.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Service type. Valid values: `mysql` (TencentDB for MySQL), `cynosdb` (TDSQL-C for MySQL). Default value: `mysql`.
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
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

// Predefined struct for user
type DescribeSlowLogTimeSeriesStatsResponseParams struct {
	// Time range in seconds in histogram.
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// Number of slow logs in the specified time range.
	TimeSeries []*TimeSlice `json:"TimeSeries,omitnil,omitempty" name:"TimeSeries"`

	// Instance CPU utilization monitoring data in the specified time range.
	SeriesData *MonitorMetricSeriesData `json:"SeriesData,omitnil,omitempty" name:"SeriesData"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSlowLogTimeSeriesStatsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSlowLogTimeSeriesStatsResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeSlowLogTopSqlsRequestParams struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Start time, such as "2019-09-10 12:13:14".
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time in the format of "2019-09-11 10:13:14". The interval between the end time and the start time can be up to 7 days.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Sorting key. Valid values: `QueryTime`, `ExecTimes`, `RowsSent`, `LockTime`, `RowsExamined`. Default value: `QueryTime`.
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// Sorting order. Valid values: `ASC` (ascending), `DESC` (descending). Default value: `DESC`.
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Number of returned results. Default value: `20`. Maximum value: `100`.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset. Default value: `0`.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Database name array.
	SchemaList []*SchemaItem `json:"SchemaList,omitnil,omitempty" name:"SchemaList"`

	// Service type. Valid values: `mysql` (TencentDB for MySQL), `cynosdb` (TDSQL-C for MySQL). Default value: `mysql`.
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type DescribeSlowLogTopSqlsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Start time, such as "2019-09-10 12:13:14".
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time in the format of "2019-09-11 10:13:14". The interval between the end time and the start time can be up to 7 days.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Sorting key. Valid values: `QueryTime`, `ExecTimes`, `RowsSent`, `LockTime`, `RowsExamined`. Default value: `QueryTime`.
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// Sorting order. Valid values: `ASC` (ascending), `DESC` (descending). Default value: `DESC`.
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Number of returned results. Default value: `20`. Maximum value: `100`.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset. Default value: `0`.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Database name array.
	SchemaList []*SchemaItem `json:"SchemaList,omitnil,omitempty" name:"SchemaList"`

	// Service type. Valid values: `mysql` (TencentDB for MySQL), `cynosdb` (TDSQL-C for MySQL). Default value: `mysql`.
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
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

// Predefined struct for user
type DescribeSlowLogTopSqlsResponseParams struct {
	// Number of eligible entries.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List of top slow SQL statements
	Rows []*SlowLogTopSqlItem `json:"Rows,omitnil,omitempty" name:"Rows"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSlowLogTopSqlsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSlowLogTopSqlsResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeSlowLogUserHostStatsRequestParams struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Start time of the time range in the format of yyyy-MM-dd HH:mm:ss, such as 2019-09-10 12:13:14.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time of the time range in the format of yyyy-MM-dd HH:mm:ss, such as 2019-09-10 12:13:14.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Service type. Valid values: mysql (TencentDB for MySQL), cynosdb (TDSQL-C for MySQL). Default value: mysql.
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// MD5 value of the SQL template
	Md5 *string `json:"Md5,omitnil,omitempty" name:"Md5"`
}

type DescribeSlowLogUserHostStatsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Start time of the time range in the format of yyyy-MM-dd HH:mm:ss, such as 2019-09-10 12:13:14.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time of the time range in the format of yyyy-MM-dd HH:mm:ss, such as 2019-09-10 12:13:14.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Service type. Valid values: mysql (TencentDB for MySQL), cynosdb (TDSQL-C for MySQL). Default value: mysql.
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// MD5 value of the SQL template
	Md5 *string `json:"Md5,omitnil,omitempty" name:"Md5"`
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

// Predefined struct for user
type DescribeSlowLogUserHostStatsResponseParams struct {
	// Total number of source addresses.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Detailed list of the proportion of slow logs from each source address.
	Items []*SlowLogHost `json:"Items,omitnil,omitempty" name:"Items"`

	// Detailed list of the percentages of slow logs from different source usernames
	UserNameItems []*SlowLogUser `json:"UserNameItems,omitnil,omitempty" name:"UserNameItems"`

	// The number of source users
	UserTotalCount *int64 `json:"UserTotalCount,omitnil,omitempty" name:"UserTotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSlowLogUserHostStatsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSlowLogUserHostStatsResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeSlowLogsRequestParams struct {
	// Service type. Valid values: `mysql` (TencentDB for MySQL), `cynosdb` (TDSQL-C for MySQL). Default value: `mysql`.
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// MD5 value of a SQL template
	Md5 *string `json:"Md5,omitnil,omitempty" name:"Md5"`

	// Start time in the format of "2019-09-10 12:13:14".
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time in the format of "2019-09-11 10:13:14". The interval between the end time and the start time can be up to 7 days.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// The offset. Default value: `0`.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The number of queried items. Default value: `20`. Max value: `100`.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Database list
	DB []*string `json:"DB,omitnil,omitempty" name:"DB"`

	// Keyword
	Key []*string `json:"Key,omitnil,omitempty" name:"Key"`

	// User
	User []*string `json:"User,omitnil,omitempty" name:"User"`

	// IP
	Ip []*string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// Duration range. The left and right borders of the range are the zeroth and first element of the array, respectively.
	Time []*int64 `json:"Time,omitnil,omitempty" name:"Time"`
}

type DescribeSlowLogsRequest struct {
	*tchttp.BaseRequest
	
	// Service type. Valid values: `mysql` (TencentDB for MySQL), `cynosdb` (TDSQL-C for MySQL). Default value: `mysql`.
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// MD5 value of a SQL template
	Md5 *string `json:"Md5,omitnil,omitempty" name:"Md5"`

	// Start time in the format of "2019-09-10 12:13:14".
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time in the format of "2019-09-11 10:13:14". The interval between the end time and the start time can be up to 7 days.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// The offset. Default value: `0`.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The number of queried items. Default value: `20`. Max value: `100`.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Database list
	DB []*string `json:"DB,omitnil,omitempty" name:"DB"`

	// Keyword
	Key []*string `json:"Key,omitnil,omitempty" name:"Key"`

	// User
	User []*string `json:"User,omitnil,omitempty" name:"User"`

	// IP
	Ip []*string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// Duration range. The left and right borders of the range are the zeroth and first element of the array, respectively.
	Time []*int64 `json:"Time,omitnil,omitempty" name:"Time"`
}

func (r *DescribeSlowLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlowLogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Product")
	delete(f, "InstanceId")
	delete(f, "Md5")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "DB")
	delete(f, "Key")
	delete(f, "User")
	delete(f, "Ip")
	delete(f, "Time")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSlowLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSlowLogsResponseParams struct {
	// Number of eligible entries.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Slow log details
	Rows []*SlowLogInfoItem `json:"Rows,omitnil,omitempty" name:"Rows"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSlowLogsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSlowLogsResponseParams `json:"Response"`
}

func (r *DescribeSlowLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlowLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopSpaceSchemaTimeSeriesRequestParams struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Number of returned top databases. Maximum value: `100`. Default value: `20`.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Field used to sort top databases. Valid values: `DataLength`, `IndexLength`, `TotalLength`, `DataFree`, `FragRatio`, `TableRows`, `PhysicalFileSize` (supported only by TencentDB for MySQL instances). For TencentDB for MySQL instances, the default value is `PhysicalFileSize`. For other database instances, the default value is `TotalLength`.
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// Start date, such as "2021-01-01". It can be as early as 29 days before the current date and is 6 days before the end date by default.
	StartDate *string `json:"StartDate,omitnil,omitempty" name:"StartDate"`

	// End date, such as "2021-01-01". It can be as early as 29 days before the current date and is the current date by default.
	EndDate *string `json:"EndDate,omitnil,omitempty" name:"EndDate"`

	// Service type. Valid values: `mysql` (TencentDB for MySQL), `cynosdb` (TDSQL-C for MySQL). Default value: `mysql`.
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type DescribeTopSpaceSchemaTimeSeriesRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Number of returned top databases. Maximum value: `100`. Default value: `20`.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Field used to sort top databases. Valid values: `DataLength`, `IndexLength`, `TotalLength`, `DataFree`, `FragRatio`, `TableRows`, `PhysicalFileSize` (supported only by TencentDB for MySQL instances). For TencentDB for MySQL instances, the default value is `PhysicalFileSize`. For other database instances, the default value is `TotalLength`.
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// Start date, such as "2021-01-01". It can be as early as 29 days before the current date and is 6 days before the end date by default.
	StartDate *string `json:"StartDate,omitnil,omitempty" name:"StartDate"`

	// End date, such as "2021-01-01". It can be as early as 29 days before the current date and is the current date by default.
	EndDate *string `json:"EndDate,omitnil,omitempty" name:"EndDate"`

	// Service type. Valid values: `mysql` (TencentDB for MySQL), `cynosdb` (TDSQL-C for MySQL). Default value: `mysql`.
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
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

// Predefined struct for user
type DescribeTopSpaceSchemaTimeSeriesResponseParams struct {
	// Time series list of the returned space statistics of top databases.
	TopSpaceSchemaTimeSeries []*SchemaSpaceTimeSeries `json:"TopSpaceSchemaTimeSeries,omitnil,omitempty" name:"TopSpaceSchemaTimeSeries"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTopSpaceSchemaTimeSeriesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTopSpaceSchemaTimeSeriesResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeTopSpaceSchemasRequestParams struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Number of returned top databases. Maximum value: 100. Default value: 20.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Field used to sort top databases. Valid values: DataLength, IndexLength, TotalLength, DataFree, FragRatio, TableRows, PhysicalFileSize (supported only by TencentDB for MySQL instances). For TencentDB for MySQL instances, the default value is `PhysicalFileSize`. For other database instances, the default value is `TotalLength`.
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// Service type. Valid values: mysql (TencentDB for MySQL), cynosdb (TDSQL-C for MySQL). Default value: mysql.
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type DescribeTopSpaceSchemasRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Number of returned top databases. Maximum value: 100. Default value: 20.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Field used to sort top databases. Valid values: DataLength, IndexLength, TotalLength, DataFree, FragRatio, TableRows, PhysicalFileSize (supported only by TencentDB for MySQL instances). For TencentDB for MySQL instances, the default value is `PhysicalFileSize`. For other database instances, the default value is `TotalLength`.
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// Service type. Valid values: mysql (TencentDB for MySQL), cynosdb (TDSQL-C for MySQL). Default value: mysql.
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
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

// Predefined struct for user
type DescribeTopSpaceSchemasResponseParams struct {
	// List of the returned space statistics of top databases.
	TopSpaceSchemas []*SchemaSpaceData `json:"TopSpaceSchemas,omitnil,omitempty" name:"TopSpaceSchemas"`

	// Timestamp (in seconds) of database space data collection points
	Timestamp *int64 `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTopSpaceSchemasResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTopSpaceSchemasResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeTopSpaceTableTimeSeriesRequestParams struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Number of returned top tables. Maximum value: `100`. Default value: `20`.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Field used to sort top tables. Valid values: `DataLength`, `IndexLength`, `TotalLength`, `DataFree`, `FragRatio`, `TableRows`, `PhysicalFileSize`. Default value: `PhysicalFileSize`.
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// Start date, such as "2021-01-01". It can be as early as 29 days before the current date and is 6 days before the end date by default.
	StartDate *string `json:"StartDate,omitnil,omitempty" name:"StartDate"`

	// End date, such as "2021-01-01". It can be as early as 29 days before the current date and is the current date by default.
	EndDate *string `json:"EndDate,omitnil,omitempty" name:"EndDate"`

	// Service type. Valid values: `mysql` (TencentDB for MySQL), `cynosdb` (TDSQL-C for MySQL). Default value: `mysql`.
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type DescribeTopSpaceTableTimeSeriesRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Number of returned top tables. Maximum value: `100`. Default value: `20`.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Field used to sort top tables. Valid values: `DataLength`, `IndexLength`, `TotalLength`, `DataFree`, `FragRatio`, `TableRows`, `PhysicalFileSize`. Default value: `PhysicalFileSize`.
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// Start date, such as "2021-01-01". It can be as early as 29 days before the current date and is 6 days before the end date by default.
	StartDate *string `json:"StartDate,omitnil,omitempty" name:"StartDate"`

	// End date, such as "2021-01-01". It can be as early as 29 days before the current date and is the current date by default.
	EndDate *string `json:"EndDate,omitnil,omitempty" name:"EndDate"`

	// Service type. Valid values: `mysql` (TencentDB for MySQL), `cynosdb` (TDSQL-C for MySQL). Default value: `mysql`.
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
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

// Predefined struct for user
type DescribeTopSpaceTableTimeSeriesResponseParams struct {
	// Time series list of the returned space statistics of top tables.
	TopSpaceTableTimeSeries []*TableSpaceTimeSeries `json:"TopSpaceTableTimeSeries,omitnil,omitempty" name:"TopSpaceTableTimeSeries"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTopSpaceTableTimeSeriesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTopSpaceTableTimeSeriesResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeTopSpaceTablesRequestParams struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Number of returned top tables. Maximum value: `100`. Default value: `20`.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Field used to sort top tables. Valid values: `DataLength`, `IndexLength`, `TotalLength`, `DataFree`, `FragRatio`, `TableRows`, `PhysicalFileSize` (only supported for TencentDB for MySQL instances). For TencentDB for MySQL instances, the default value is `PhysicalFileSize`. For other database instances, the default value is `TotalLength`.
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// Service type. Valid values: `mysql` (TencentDB for MySQL), `cynosdb` (TDSQL-C for MySQL). Default value: `mysql`.
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type DescribeTopSpaceTablesRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Number of returned top tables. Maximum value: `100`. Default value: `20`.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Field used to sort top tables. Valid values: `DataLength`, `IndexLength`, `TotalLength`, `DataFree`, `FragRatio`, `TableRows`, `PhysicalFileSize` (only supported for TencentDB for MySQL instances). For TencentDB for MySQL instances, the default value is `PhysicalFileSize`. For other database instances, the default value is `TotalLength`.
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// Service type. Valid values: `mysql` (TencentDB for MySQL), `cynosdb` (TDSQL-C for MySQL). Default value: `mysql`.
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
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

// Predefined struct for user
type DescribeTopSpaceTablesResponseParams struct {
	// List of the returned space statistics of top tables.
	TopSpaceTables []*TableSpaceData `json:"TopSpaceTables,omitnil,omitempty" name:"TopSpaceTables"`

	// Timestamp (in seconds) of tablespace data collection points
	Timestamp *int64 `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTopSpaceTablesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTopSpaceTablesResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeUserSqlAdviceRequestParams struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// SQL statement.
	SqlText *string `json:"SqlText,omitnil,omitempty" name:"SqlText"`

	// Database name.
	Schema *string `json:"Schema,omitnil,omitempty" name:"Schema"`

	// Service type. Valid values: `mysql` (TencentDB for MySQL), `cynosdb` (TDSQL-C for MySQL), `dbbrain-mysql` (self-built MySQL). Default value: `mysql`.
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type DescribeUserSqlAdviceRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// SQL statement.
	SqlText *string `json:"SqlText,omitnil,omitempty" name:"SqlText"`

	// Database name.
	Schema *string `json:"Schema,omitnil,omitempty" name:"Schema"`

	// Service type. Valid values: `mysql` (TencentDB for MySQL), `cynosdb` (TDSQL-C for MySQL), `dbbrain-mysql` (self-built MySQL). Default value: `mysql`.
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
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
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserSqlAdviceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserSqlAdviceResponseParams struct {
	// SQL statement optimization suggestions, which can be parsed into JSON arrays. If there is no need for optimization, the output will be empty.
	Advices *string `json:"Advices,omitnil,omitempty" name:"Advices"`

	// Notes of SQL statement optimization suggestions, which can be parsed into String arrays. If there is no need for optimization, the output will be empty.
	Comments *string `json:"Comments,omitnil,omitempty" name:"Comments"`

	// SQL statement.
	SqlText *string `json:"SqlText,omitnil,omitempty" name:"SqlText"`

	// Database name.
	Schema *string `json:"Schema,omitnil,omitempty" name:"Schema"`

	// DDL information of related tables, which can be parsed into JSON arrays.
	Tables *string `json:"Tables,omitnil,omitempty" name:"Tables"`

	// SQL execution plan, which can be parsed into JSON arrays. If there is no need for optimization, the output will be empty.
	SqlPlan *string `json:"SqlPlan,omitnil,omitempty" name:"SqlPlan"`

	// Cost saving details after SQL statement optimization, which can be parsed into JSON arrays. If there is no need for optimization, the output will be empty.
	Cost *string `json:"Cost,omitnil,omitempty" name:"Cost"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeUserSqlAdviceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserSqlAdviceResponseParams `json:"Response"`
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
	DiagType *string `json:"DiagType,omitnil,omitempty" name:"DiagType"`

	// End time.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Start time.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Unique event ID.
	EventId *int64 `json:"EventId,omitnil,omitempty" name:"EventId"`

	// Severity, which can be divided into 5 levels: 1: fatal, 2: severe, 3: warning, 4: notice, 5: healthy.
	Severity *int64 `json:"Severity,omitnil,omitempty" name:"Severity"`

	// Diagnosis summary.
	Outline *string `json:"Outline,omitnil,omitempty" name:"Outline"`

	// Diagnosis item description.
	DiagItem *string `json:"DiagItem,omitnil,omitempty" name:"DiagItem"`

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Reserved field.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Metric *string `json:"Metric,omitnil,omitempty" name:"Metric"`

	// Region.
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`
}

type EventInfo struct {
	// Event ID.
	EventId *int64 `json:"EventId,omitnil,omitempty" name:"EventId"`

	// Diagnosis type.
	DiagType *string `json:"DiagType,omitnil,omitempty" name:"DiagType"`

	// Start time.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Summary.
	Outline *string `json:"Outline,omitnil,omitempty" name:"Outline"`

	// Severity, which can be divided into 5 levels: `1` (Critical), `2` (Severe), `3` (Alarm), `4` (Reminder), `5` (Healthy).
	Severity *int64 `json:"Severity,omitnil,omitempty" name:"Severity"`

	// Deduction.
	ScoreLost *int64 `json:"ScoreLost,omitnil,omitempty" name:"ScoreLost"`

	// Reserved field.
	Metric *string `json:"Metric,omitnil,omitempty" name:"Metric"`

	// Number of alarms.
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`
}

type GroupItem struct {
	// Group ID.
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// Group name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Number of group members.
	MemberCount *int64 `json:"MemberCount,omitnil,omitempty" name:"MemberCount"`
}

type HealthReportTask struct {
	// Async task request ID.
	AsyncRequestId *int64 `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// Source that triggers the task. Valid values: `DAILY_INSPECTION` (instance inspection), `SCHEDULED` (scheduled task), and `MANUAL` (manual trigger).
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// Task progress in %.
	Progress *int64 `json:"Progress,omitnil,omitempty" name:"Progress"`

	// Task creation time.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Task start time.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Task end time.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Basic information of the instance to which the task belongs.
	InstanceInfo *InstanceBasicInfo `json:"InstanceInfo,omitnil,omitempty" name:"InstanceInfo"`

	// Health information in health report.
	HealthStatus *HealthStatus `json:"HealthStatus,omitnil,omitempty" name:"HealthStatus"`
}

type HealthScoreInfo struct {
	// Exception details.
	IssueTypes []*IssueTypeInfo `json:"IssueTypes,omitnil,omitempty" name:"IssueTypes"`

	// Total number of exceptions.
	EventsTotalCount *int64 `json:"EventsTotalCount,omitnil,omitempty" name:"EventsTotalCount"`

	// Health score.
	HealthScore *int64 `json:"HealthScore,omitnil,omitempty" name:"HealthScore"`

	// Health level, such as `HEALTH`, `SUB_HEALTH`, `RISK`, and `HIGH_RISK`.
	HealthLevel *string `json:"HealthLevel,omitnil,omitempty" name:"HealthLevel"`
}

type HealthStatus struct {
	// Health score out of 100 points.
	HealthScore *int64 `json:"HealthScore,omitnil,omitempty" name:"HealthScore"`

	// Health level. Valid values: `HEALTH` (healthy), `SUB_HEALTH` (sub-healthy), `RISK` (dangerous), and `HIGH_RISK` (high-risk).
	HealthLevel *string `json:"HealthLevel,omitnil,omitempty" name:"HealthLevel"`

	// Total deducted scores.
	ScoreLost *int64 `json:"ScoreLost,omitnil,omitempty" name:"ScoreLost"`

	// Deduction details.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ScoreDetails []*ScoreDetail `json:"ScoreDetails,omitnil,omitempty" name:"ScoreDetails"`
}

type InstanceBasicInfo struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Instance name.
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Private IP of the instance.
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// Private port of the instance.
	Vport *int64 `json:"Vport,omitnil,omitempty" name:"Vport"`

	// Instance service.
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// Instance engine version.
	EngineVersion *string `json:"EngineVersion,omitnil,omitempty" name:"EngineVersion"`
}

type InstanceConfs struct {
	// Whether to enable database inspection. Valid values: Yes, No.
	DailyInspection *string `json:"DailyInspection,omitnil,omitempty" name:"DailyInspection"`

	// Whether to enable instance overview. Valid values: Yes, No.
	OverviewDisplay *string `json:"OverviewDisplay,omitnil,omitempty" name:"OverviewDisplay"`

	// Custom big key analysis separator for Redis only
	// Note: This field may return null, indicating that no valid values can be obtained.
	KeyDelimiters []*string `json:"KeyDelimiters,omitnil,omitempty" name:"KeyDelimiters"`
}

type InstanceInfo struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Instance name.
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Instance region.
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Health score.
	HealthScore *int64 `json:"HealthScore,omitnil,omitempty" name:"HealthScore"`

	// Service.
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// Number of exceptions.
	EventCount *int64 `json:"EventCount,omitnil,omitempty" name:"EventCount"`

	// Instance type. Valid values: 1 (MASTER), 2 (DR), 3 (RO), 4 (SDR)
	InstanceType *int64 `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// Number of cores.
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// Memory in MB.
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// Disk storage in GB.
	Volume *int64 `json:"Volume,omitnil,omitempty" name:"Volume"`

	// Database version.
	EngineVersion *string `json:"EngineVersion,omitnil,omitempty" name:"EngineVersion"`

	// Private network address.
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// Private network port.
	Vport *int64 `json:"Vport,omitnil,omitempty" name:"Vport"`

	// Access source.
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// Group ID.
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// Group name.
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// Instance status. Valid values: 0 (delivering), 1 (running), 4 (terminating), 5 (isolated)
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Unified subnet ID.
	UniqSubnetId *string `json:"UniqSubnetId,omitnil,omitempty" name:"UniqSubnetId"`

	// TencentDB instance type.
	DeployMode *string `json:"DeployMode,omitnil,omitempty" name:"DeployMode"`

	// TencentDB instance initialization flag. Valid values: 0 (not initialized), 1 (initialized).
	InitFlag *int64 `json:"InitFlag,omitnil,omitempty" name:"InitFlag"`

	// Task status.
	TaskStatus *int64 `json:"TaskStatus,omitnil,omitempty" name:"TaskStatus"`

	// Unified VPC ID.
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// Instance inspection/overview status.
	InstanceConf *InstanceConfs `json:"InstanceConf,omitnil,omitempty" name:"InstanceConf"`

	// Resource expiration time.
	DeadlineTime *string `json:"DeadlineTime,omitnil,omitempty" name:"DeadlineTime"`

	// Whether it is an instance supported by DBbrain.
	IsSupported *bool `json:"IsSupported,omitnil,omitempty" name:"IsSupported"`

	// Status of instance security audit log. Valid values: ON (enabled), OFF (disabled).
	SecAuditStatus *string `json:"SecAuditStatus,omitnil,omitempty" name:"SecAuditStatus"`

	// Status of instance audit log. Valid values: ALL_AUDIT (full audit is enabled), RULE_AUDIT (rule audit is enabled), UNBOUND (audit is disabled).
	AuditPolicyStatus *string `json:"AuditPolicyStatus,omitnil,omitempty" name:"AuditPolicyStatus"`

	// Running status of instance audit log. Valid values: normal (running), paused (suspension due to overdue payment).
	AuditRunningStatus *string `json:"AuditRunningStatus,omitnil,omitempty" name:"AuditRunningStatus"`

	// Private VIP 
	// Note: This field may return null, indicating that no valid values can be obtained.
	InternalVip *string `json:"InternalVip,omitnil,omitempty" name:"InternalVip"`

	// Private network port 
	// Note: This field may return null, indicating that no valid values can be obtained.
	InternalVport *int64 `json:"InternalVport,omitnil,omitempty" name:"InternalVport"`

	// Creation time
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Cluster ID. This field is only required for cluster database products like TDSQL-C. 
	// Note: This field may return null, indicating that no valid values can be obtained.
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Cluster name. This field is only required for cluster database products like TDSQL-C. 
	// Note: This field may return null, indicating that no valid values can be obtained.
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`
}

type IssueTypeInfo struct {
	// Metric categories. Valid values: `AVAILABILITY`, `MAINTAINABILITY`, `PERFORMANCE`, and `RELIABILITY`.
	IssueType *string `json:"IssueType,omitnil,omitempty" name:"IssueType"`

	// Exception.
	Events []*EventInfo `json:"Events,omitnil,omitempty" name:"Events"`

	// Total number of exceptions.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`
}

// Predefined struct for user
type KillMySqlThreadsRequestParams struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// The stage of a session killing task. Valid values: `Prepare` (preparation stage), `Commit` (commit stage).
	Stage *string `json:"Stage,omitnil,omitempty" name:"Stage"`

	// List of IDs of the MySQL sessions to be killed. This parameter is used in the `Prepare` stage.
	Threads []*int64 `json:"Threads,omitnil,omitempty" name:"Threads"`

	// Execution ID. This parameter is used in the `Commit` stage.
	SqlExecId *string `json:"SqlExecId,omitnil,omitempty" name:"SqlExecId"`

	// Service type. Valid values: `mysql` (TencentDB for MySQL), `cynosdb` (TDSQL-C for MySQL). Default value: `mysql`.
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// Whether to record the thread killing history. The default value is `true`, indicating “yes”. You can set it to `false` (“no”) to speed up the killing process.
	RecordHistory *bool `json:"RecordHistory,omitnil,omitempty" name:"RecordHistory"`
}

type KillMySqlThreadsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// The stage of a session killing task. Valid values: `Prepare` (preparation stage), `Commit` (commit stage).
	Stage *string `json:"Stage,omitnil,omitempty" name:"Stage"`

	// List of IDs of the MySQL sessions to be killed. This parameter is used in the `Prepare` stage.
	Threads []*int64 `json:"Threads,omitnil,omitempty" name:"Threads"`

	// Execution ID. This parameter is used in the `Commit` stage.
	SqlExecId *string `json:"SqlExecId,omitnil,omitempty" name:"SqlExecId"`

	// Service type. Valid values: `mysql` (TencentDB for MySQL), `cynosdb` (TDSQL-C for MySQL). Default value: `mysql`.
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// Whether to record the thread killing history. The default value is `true`, indicating “yes”. You can set it to `false` (“no”) to speed up the killing process.
	RecordHistory *bool `json:"RecordHistory,omitnil,omitempty" name:"RecordHistory"`
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
	delete(f, "RecordHistory")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "KillMySqlThreadsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type KillMySqlThreadsResponseParams struct {
	// List of IDs of the MySQL sessions that have been killed.
	Threads []*int64 `json:"Threads,omitnil,omitempty" name:"Threads"`

	// Execution ID, which is output in the `Prepare` stage and used to specify the ID of the session to be killed in the `Commit` stage.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SqlExecId *string `json:"SqlExecId,omitnil,omitempty" name:"SqlExecId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type KillMySqlThreadsResponse struct {
	*tchttp.BaseResponse
	Response *KillMySqlThreadsResponseParams `json:"Response"`
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
	// Whether to enable email sending. Valid values: `0` (no), `1` (yes).
	SendMail *int64 `json:"SendMail,omitnil,omitempty" name:"SendMail"`

	// Region configuration, such as "ap-guangzhou" and "ap-shanghai". For the inspection email sending template, configure the region where you need to send the inspection email. For the subscription email sending template, configure the region where the current subscribed instance resides.
	Region []*string `json:"Region,omitnil,omitempty" name:"Region"`

	// Sends a report with the specified health level, such as `HEALTH`, `SUB_HEALTH`, `RISK`, and `HIGH_RISK`.
	HealthStatus []*string `json:"HealthStatus,omitnil,omitempty" name:"HealthStatus"`

	// Recipient ID. Either `ContactPerson` or `ContactGroup` should be passed in.
	ContactPerson []*int64 `json:"ContactPerson,omitnil,omitempty" name:"ContactPerson"`

	// Recipient group ID. Either `ContactPerson` or `ContactGroup` should be passed in.
	ContactGroup []*int64 `json:"ContactGroup,omitnil,omitempty" name:"ContactGroup"`
}

// Predefined struct for user
type ModifyAuditServiceRequestParams struct {
	// Service type. Valid values: `dcdb` (TDSQL for MySQL), `mariadb` (TencentDB for MariaDB).
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// Use the value of `u200cProduct` for this parameter, such as `dcdb` and `mariadb`.
	NodeRequestType *string `json:"NodeRequestType,omitnil,omitempty" name:"NodeRequestType"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Total log retention period in days. Valid values: `7`, `30`, `90`, `180`, `365`, `1095`, `1825`.
	LogExpireDay *int64 `json:"LogExpireDay,omitnil,omitempty" name:"LogExpireDay"`

	// Storage period of frequently accessed logs in days. Valid values: `7`, `30`, `90`, `180`, `365`, `1095`, `1825`.
	HotLogExpireDay *int64 `json:"HotLogExpireDay,omitnil,omitempty" name:"HotLogExpireDay"`
}

type ModifyAuditServiceRequest struct {
	*tchttp.BaseRequest
	
	// Service type. Valid values: `dcdb` (TDSQL for MySQL), `mariadb` (TencentDB for MariaDB).
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// Use the value of `u200cProduct` for this parameter, such as `dcdb` and `mariadb`.
	NodeRequestType *string `json:"NodeRequestType,omitnil,omitempty" name:"NodeRequestType"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Total log retention period in days. Valid values: `7`, `30`, `90`, `180`, `365`, `1095`, `1825`.
	LogExpireDay *int64 `json:"LogExpireDay,omitnil,omitempty" name:"LogExpireDay"`

	// Storage period of frequently accessed logs in days. Valid values: `7`, `30`, `90`, `180`, `365`, `1095`, `1825`.
	HotLogExpireDay *int64 `json:"HotLogExpireDay,omitnil,omitempty" name:"HotLogExpireDay"`
}

func (r *ModifyAuditServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAuditServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Product")
	delete(f, "NodeRequestType")
	delete(f, "InstanceId")
	delete(f, "LogExpireDay")
	delete(f, "HotLogExpireDay")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAuditServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAuditServiceResponseParams struct {
	// Audit configuration modification result. If `0` is returned, the modification is successful; otherwise, an exception will be returned, indicating that the modification failed.
	Success *int64 `json:"Success,omitnil,omitempty" name:"Success"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAuditServiceResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAuditServiceResponseParams `json:"Response"`
}

func (r *ModifyAuditServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAuditServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDiagDBInstanceConfRequestParams struct {
	// Instance configuration, including inspection and overview switch.
	InstanceConfs *InstanceConfs `json:"InstanceConfs,omitnil,omitempty" name:"InstanceConfs"`

	// Target regions of the request. If the value is `All`, it is applied to all regions.
	Regions *string `json:"Regions,omitnil,omitempty" name:"Regions"`

	// Service type. Valid values: `mysql` (TencentDB for MySQL), `cynosdb` (TDSQL-C for MySQL).
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// ID of the instance to modify.
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

type ModifyDiagDBInstanceConfRequest struct {
	*tchttp.BaseRequest
	
	// Instance configuration, including inspection and overview switch.
	InstanceConfs *InstanceConfs `json:"InstanceConfs,omitnil,omitempty" name:"InstanceConfs"`

	// Target regions of the request. If the value is `All`, it is applied to all regions.
	Regions *string `json:"Regions,omitnil,omitempty" name:"Regions"`

	// Service type. Valid values: `mysql` (TencentDB for MySQL), `cynosdb` (TDSQL-C for MySQL).
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// ID of the instance to modify.
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
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

// Predefined struct for user
type ModifyDiagDBInstanceConfResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDiagDBInstanceConfResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDiagDBInstanceConfResponseParams `json:"Response"`
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
	Metric *string `json:"Metric,omitnil,omitempty" name:"Metric"`

	// Metric unit.
	Unit *string `json:"Unit,omitnil,omitempty" name:"Unit"`

	// Metric value.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Values []*float64 `json:"Values,omitnil,omitempty" name:"Values"`
}

type MonitorFloatMetricSeriesData struct {
	// Monitoring metric.
	Series []*MonitorFloatMetric `json:"Series,omitnil,omitempty" name:"Series"`

	// Timestamp corresponding to monitoring metric.
	Timestamp []*int64 `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`
}

type MonitorMetric struct {
	// Metric name.
	Metric *string `json:"Metric,omitnil,omitempty" name:"Metric"`

	// Metric unit.
	Unit *string `json:"Unit,omitnil,omitempty" name:"Unit"`

	// Metric value.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Values []*float64 `json:"Values,omitnil,omitempty" name:"Values"`
}

type MonitorMetricSeriesData struct {
	// Monitoring metric.
	Series []*MonitorMetric `json:"Series,omitnil,omitempty" name:"Series"`

	// Timestamp corresponding to monitoring metric.
	Timestamp []*int64 `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`
}

type MySqlProcess struct {
	// Thread ID.
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`

	// Thread operation account name.
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// Thread operation host address.
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// Thread operation database.
	DB *string `json:"DB,omitnil,omitempty" name:"DB"`

	// Thread operation status.
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// Thread execution type.
	Command *string `json:"Command,omitnil,omitempty" name:"Command"`

	// Thread operation duration in seconds.
	Time *string `json:"Time,omitnil,omitempty" name:"Time"`

	// Thread operation statement.
	Info *string `json:"Info,omitnil,omitempty" name:"Info"`
}

// Predefined struct for user
type OpenAuditServiceRequestParams struct {
	// Service type. Valid values: `dcdb` (TDSQL for MySQL), `mariadb` (TencentDB for MariaDB).
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// Use the value of `u200cProduct` for this parameter, such as `dcdb` and `mariadb`.
	NodeRequestType *string `json:"NodeRequestType,omitnil,omitempty" name:"NodeRequestType"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Total log retention period in days. Valid values: `7`, `30`, `90`, `180`, `365`, `1095`, `1825`.
	LogExpireDay *int64 `json:"LogExpireDay,omitnil,omitempty" name:"LogExpireDay"`

	// Storage period of frequently accessed logs in days. Valid values: `7`, `30`, `90`, `180`, `365`, `1095`, `1825`.
	HotLogExpireDay *int64 `json:"HotLogExpireDay,omitnil,omitempty" name:"HotLogExpireDay"`
}

type OpenAuditServiceRequest struct {
	*tchttp.BaseRequest
	
	// Service type. Valid values: `dcdb` (TDSQL for MySQL), `mariadb` (TencentDB for MariaDB).
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// Use the value of `u200cProduct` for this parameter, such as `dcdb` and `mariadb`.
	NodeRequestType *string `json:"NodeRequestType,omitnil,omitempty" name:"NodeRequestType"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Total log retention period in days. Valid values: `7`, `30`, `90`, `180`, `365`, `1095`, `1825`.
	LogExpireDay *int64 `json:"LogExpireDay,omitnil,omitempty" name:"LogExpireDay"`

	// Storage period of frequently accessed logs in days. Valid values: `7`, `30`, `90`, `180`, `365`, `1095`, `1825`.
	HotLogExpireDay *int64 `json:"HotLogExpireDay,omitnil,omitempty" name:"HotLogExpireDay"`
}

func (r *OpenAuditServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenAuditServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Product")
	delete(f, "NodeRequestType")
	delete(f, "InstanceId")
	delete(f, "LogExpireDay")
	delete(f, "HotLogExpireDay")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OpenAuditServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenAuditServiceResponseParams struct {
	// Audit is successfully enabled only when the value of this parameter is `0`.
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type OpenAuditServiceResponse struct {
	*tchttp.BaseResponse
	Response *OpenAuditServiceResponseParams `json:"Response"`
}

func (r *OpenAuditServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenAuditServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProcessStatistic struct {
	// Array of session details
	Items []*SessionItem `json:"Items,omitnil,omitempty" name:"Items"`

	// The total number of connections
	AllConnSum *int64 `json:"AllConnSum,omitnil,omitempty" name:"AllConnSum"`

	// The total number of active connections
	ActiveConnSum *int64 `json:"ActiveConnSum,omitnil,omitempty" name:"ActiveConnSum"`
}

type ProfileInfo struct {
	// Email language, such as `en`.
	Language *string `json:"Language,omitnil,omitempty" name:"Language"`

	// Email template content.
	MailConfiguration *MailConfiguration `json:"MailConfiguration,omitnil,omitempty" name:"MailConfiguration"`
}

type RedisPreKeySpaceData struct {
	// Average element length
	AveElementSize *int64 `json:"AveElementSize,omitnil,omitempty" name:"AveElementSize"`

	// Total memory usage in bytes
	Length *int64 `json:"Length,omitnil,omitempty" name:"Length"`

	// Key prefix
	KeyPreIndex *string `json:"KeyPreIndex,omitnil,omitempty" name:"KeyPreIndex"`

	// The number of elements
	ItemCount *int64 `json:"ItemCount,omitnil,omitempty" name:"ItemCount"`

	// The number of keys
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// The max element length
	MaxElementSize *int64 `json:"MaxElementSize,omitnil,omitempty" name:"MaxElementSize"`
}

type SchemaItem struct {
	// Database name
	Schema *string `json:"Schema,omitnil,omitempty" name:"Schema"`
}

type SchemaSpaceData struct {
	// Database name.
	TableSchema *string `json:"TableSchema,omitnil,omitempty" name:"TableSchema"`

	// Data space in MB.
	DataLength *float64 `json:"DataLength,omitnil,omitempty" name:"DataLength"`

	// Index space in MB.
	IndexLength *float64 `json:"IndexLength,omitnil,omitempty" name:"IndexLength"`

	// Fragmented space in MB.
	DataFree *float64 `json:"DataFree,omitnil,omitempty" name:"DataFree"`

	// Total space usage in MB.
	TotalLength *float64 `json:"TotalLength,omitnil,omitempty" name:"TotalLength"`

	// Fragmentation rate in %.
	FragRatio *float64 `json:"FragRatio,omitnil,omitempty" name:"FragRatio"`

	// Number of rows.
	TableRows *int64 `json:"TableRows,omitnil,omitempty" name:"TableRows"`

	// Total size in MB of physical files exclusive to all tables in the database.
	// Note: this field may return null, indicating that no valid values can be obtained.
	PhysicalFileSize *float64 `json:"PhysicalFileSize,omitnil,omitempty" name:"PhysicalFileSize"`
}

type SchemaSpaceTimeSeries struct {
	// Database name
	TableSchema *string `json:"TableSchema,omitnil,omitempty" name:"TableSchema"`

	// Space metric value in a unit of time interval
	SeriesData *MonitorMetricSeriesData `json:"SeriesData,omitnil,omitempty" name:"SeriesData"`
}

type ScoreDetail struct {
	// Deduction item type. Valid values: `Availability`, `Maintainability`, `Performance`, `Reliability`.
	IssueType *string `json:"IssueType,omitnil,omitempty" name:"IssueType"`

	// Total deducted scores.
	ScoreLost *int64 `json:"ScoreLost,omitnil,omitempty" name:"ScoreLost"`

	// Upper limit of the deducted scores.
	ScoreLostMax *int64 `json:"ScoreLostMax,omitnil,omitempty" name:"ScoreLostMax"`

	// List of deduction items.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Items []*ScoreItem `json:"Items,omitnil,omitempty" name:"Items"`
}

type ScoreItem struct {
	// Exception diagnosis item name.
	DiagItem *string `json:"DiagItem,omitnil,omitempty" name:"DiagItem"`

	// Diagnosis item type. Valid values: `Availability`, `Maintainability`, `Performance`, `Reliability`.
	IssueType *string `json:"IssueType,omitnil,omitempty" name:"IssueType"`

	// Health level. Valid values: `Healthy`, `Reminder`, `Alarm`, `Severe`, `Critical`.
	TopSeverity *string `json:"TopSeverity,omitnil,omitempty" name:"TopSeverity"`

	// Number of occurrences of this exception diagnosis item.
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// Deducted scores.
	ScoreLost *int64 `json:"ScoreLost,omitnil,omitempty" name:"ScoreLost"`
}

type SecLogExportTaskInfo struct {
	// Async task Id.
	AsyncRequestId *uint64 `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// Task start time.
	// Note: This field may return null, indicating that no valid values can be obtained.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Task end time.
	// Note: This field may return null, indicating that no valid values can be obtained.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Task creation time.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Task status.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Task progress.
	Progress *uint64 `json:"Progress,omitnil,omitempty" name:"Progress"`

	// Exported log start time.
	// Note: This field may return null, indicating that no valid values can be obtained.
	LogStartTime *string `json:"LogStartTime,omitnil,omitempty" name:"LogStartTime"`

	// Exported log end time.
	// Note: This field may return null, indicating that no valid values can be obtained.
	LogEndTime *string `json:"LogEndTime,omitnil,omitempty" name:"LogEndTime"`

	// Total size of log files in KB.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TotalSize *uint64 `json:"TotalSize,omitnil,omitempty" name:"TotalSize"`

	// List of risk levels. Valid values: `0` (no risk), `1` (low risk), `2` (medium risk), `3` (high risk).
	// Note: This field may return null, indicating that no valid values can be obtained.
	DangerLevels []*uint64 `json:"DangerLevels,omitnil,omitempty" name:"DangerLevels"`
}

type SessionItem struct {
	// Access source
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// The number of active connections from the current access source
	ActiveConn *string `json:"ActiveConn,omitnil,omitempty" name:"ActiveConn"`

	// The total number of connections from the current access source
	AllConn *int64 `json:"AllConn,omitnil,omitempty" name:"AllConn"`
}

type SlowLogHost struct {
	// Source addresses.
	UserHost *string `json:"UserHost,omitnil,omitempty" name:"UserHost"`

	// Proportion (in %) of slow logs from this source address to the total number of slow logs.
	Ratio *float64 `json:"Ratio,omitnil,omitempty" name:"Ratio"`

	// Number of slow logs from this source address.
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`
}

type SlowLogInfoItem struct {
	// Slow log start time
	Timestamp *string `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`

	// SQL statement
	SqlText *string `json:"SqlText,omitnil,omitempty" name:"SqlText"`

	// Database
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// User source
	// Note: This field may return null, indicating that no valid values can be obtained.
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// IP source
	// Note: This field may return null, indicating that no valid values can be obtained.
	UserHost *string `json:"UserHost,omitnil,omitempty" name:"UserHost"`

	// Execution time in seconds
	QueryTime *int64 `json:"QueryTime,omitnil,omitempty" name:"QueryTime"`

	// Lock time in seconds
	// Note: This field may return null, indicating that no valid values can be obtained.
	LockTime *int64 `json:"LockTime,omitnil,omitempty" name:"LockTime"`

	// Number of scanned rows
	// Note: This field may return null, indicating that no valid values can be obtained.
	RowsExamined *int64 `json:"RowsExamined,omitnil,omitempty" name:"RowsExamined"`

	// Number of returned rows
	// Note: This field may return null, indicating that no valid values can be obtained.
	RowsSent *int64 `json:"RowsSent,omitnil,omitempty" name:"RowsSent"`
}

type SlowLogTopSqlItem struct {
	// Total SQL lock wait time in seconds.
	LockTime *float64 `json:"LockTime,omitnil,omitempty" name:"LockTime"`

	// Maximum lock wait time in seconds
	LockTimeMax *float64 `json:"LockTimeMax,omitnil,omitempty" name:"LockTimeMax"`

	// Minimum lock wait time in seconds
	LockTimeMin *float64 `json:"LockTimeMin,omitnil,omitempty" name:"LockTimeMin"`

	// Total number of scanned rows
	RowsExamined *int64 `json:"RowsExamined,omitnil,omitempty" name:"RowsExamined"`

	// Maximum number of scanned rows
	RowsExaminedMax *int64 `json:"RowsExaminedMax,omitnil,omitempty" name:"RowsExaminedMax"`

	// Minimum number of scanned rows
	RowsExaminedMin *int64 `json:"RowsExaminedMin,omitnil,omitempty" name:"RowsExaminedMin"`

	// Total duration in seconds
	QueryTime *float64 `json:"QueryTime,omitnil,omitempty" name:"QueryTime"`

	// Maximum execution time in seconds
	QueryTimeMax *float64 `json:"QueryTimeMax,omitnil,omitempty" name:"QueryTimeMax"`

	// Minimum execution time in seconds
	QueryTimeMin *float64 `json:"QueryTimeMin,omitnil,omitempty" name:"QueryTimeMin"`

	// Total number of returned rows
	RowsSent *int64 `json:"RowsSent,omitnil,omitempty" name:"RowsSent"`

	// Maximum number of returned rows
	RowsSentMax *int64 `json:"RowsSentMax,omitnil,omitempty" name:"RowsSentMax"`

	// Minimum number of returned rows
	RowsSentMin *int64 `json:"RowsSentMin,omitnil,omitempty" name:"RowsSentMin"`

	// Number of executions
	ExecTimes *int64 `json:"ExecTimes,omitnil,omitempty" name:"ExecTimes"`

	// SQL template
	SqlTemplate *string `json:"SqlTemplate,omitnil,omitempty" name:"SqlTemplate"`

	// SQL statements with parameter (random)
	SqlText *string `json:"SqlText,omitnil,omitempty" name:"SqlText"`

	// Database name
	Schema *string `json:"Schema,omitnil,omitempty" name:"Schema"`

	// Ratio of the total duration in %
	QueryTimeRatio *float64 `json:"QueryTimeRatio,omitnil,omitempty" name:"QueryTimeRatio"`

	// Ratio of the total SQL lock wait time in %
	LockTimeRatio *float64 `json:"LockTimeRatio,omitnil,omitempty" name:"LockTimeRatio"`

	// Ratio of total number of scanned rows in %
	RowsExaminedRatio *float64 `json:"RowsExaminedRatio,omitnil,omitempty" name:"RowsExaminedRatio"`

	// Ratio of total number of returned rows in %
	RowsSentRatio *float64 `json:"RowsSentRatio,omitnil,omitempty" name:"RowsSentRatio"`

	// Average execution time in seconds
	QueryTimeAvg *float64 `json:"QueryTimeAvg,omitnil,omitempty" name:"QueryTimeAvg"`

	// Average number of returned rows
	RowsSentAvg *float64 `json:"RowsSentAvg,omitnil,omitempty" name:"RowsSentAvg"`

	// Average lock wait time in seconds
	LockTimeAvg *float64 `json:"LockTimeAvg,omitnil,omitempty" name:"LockTimeAvg"`

	// Average number of scanned rows
	RowsExaminedAvg *float64 `json:"RowsExaminedAvg,omitnil,omitempty" name:"RowsExaminedAvg"`

	// MD5 value of the SQL template
	Md5 *string `json:"Md5,omitnil,omitempty" name:"Md5"`
}

type SlowLogUser struct {
	// Source username
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// Percentage of the number of slow logs from this source username to the total number of slow logs
	Ratio *float64 `json:"Ratio,omitnil,omitempty" name:"Ratio"`

	// Number of slow logs from this source username
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`
}

type TableSpaceData struct {
	// Table name.
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// Database name.
	TableSchema *string `json:"TableSchema,omitnil,omitempty" name:"TableSchema"`

	// Database table storage engine.
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// Data space in MB.
	DataLength *float64 `json:"DataLength,omitnil,omitempty" name:"DataLength"`

	// Index space in MB.
	IndexLength *float64 `json:"IndexLength,omitnil,omitempty" name:"IndexLength"`

	// Fragmented space in MB.
	DataFree *float64 `json:"DataFree,omitnil,omitempty" name:"DataFree"`

	// Total space usage in MB.
	TotalLength *float64 `json:"TotalLength,omitnil,omitempty" name:"TotalLength"`

	// Fragmentation rate in %.
	FragRatio *float64 `json:"FragRatio,omitnil,omitempty" name:"FragRatio"`

	// Number of rows.
	TableRows *int64 `json:"TableRows,omitnil,omitempty" name:"TableRows"`

	// Size in MB of the physical file exclusive to a table.
	PhysicalFileSize *float64 `json:"PhysicalFileSize,omitnil,omitempty" name:"PhysicalFileSize"`
}

type TableSpaceTimeSeries struct {
	// Table name.
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// Database name.
	TableSchema *string `json:"TableSchema,omitnil,omitempty" name:"TableSchema"`

	// Database table storage engine.
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// Space metric value in a unit of time interval
	SeriesData *MonitorFloatMetricSeriesData `json:"SeriesData,omitnil,omitempty" name:"SeriesData"`
}

type TaskInfo struct {
	// Async task ID.
	AsyncRequestId *int64 `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// List of all proxies of the current instance.
	InstProxyList []*string `json:"InstProxyList,omitnil,omitempty" name:"InstProxyList"`

	// Total number of proxies of the current instance.
	InstProxyCount *int64 `json:"InstProxyCount,omitnil,omitempty" name:"InstProxyCount"`

	// Task creation time.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Task start time.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Task status. Valid values: `created` (create), `chosen` (to be executed), `running` (being executed), `failed` (failed), and `finished` (completed).
	TaskStatus *string `json:"TaskStatus,omitnil,omitempty" name:"TaskStatus"`

	// IDs of the proxies that have completed the session killing tasks.
	FinishedProxyList []*string `json:"FinishedProxyList,omitnil,omitempty" name:"FinishedProxyList"`

	// IDs of the proxies that failed to execute the session killing tasks.
	FailedProxyList []*string `json:"FailedProxyList,omitnil,omitempty" name:"FailedProxyList"`

	// Task end time.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Task progress.
	Progress *int64 `json:"Progress,omitnil,omitempty" name:"Progress"`

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type TimeSlice struct {
	// Total number
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// Statistics start time
	Timestamp *int64 `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`
}

type UserProfile struct {
	// Configured ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	ProfileId *string `json:"ProfileId,omitnil,omitempty" name:"ProfileId"`

	// Configuration type. Valid values: `dbScan_mail_configuration` (email configuration of the database inspection report), `scheduler_mail_configuration` (email configuration of the scheduled task report).
	// Note: This field may return null, indicating that no valid values can be obtained.
	ProfileType *string `json:"ProfileType,omitnil,omitempty" name:"ProfileType"`

	// Configuration level. Valid values: `User` (user-level), `Instance` (instance-level). For database inspection emails, it should be `User`. For scheduled task emails, it should be `Instance`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ProfileLevel *string `json:"ProfileLevel,omitnil,omitempty" name:"ProfileLevel"`

	// Configuration name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ProfileName *string `json:"ProfileName,omitnil,omitempty" name:"ProfileName"`

	// Configuration details.
	ProfileInfo *ProfileInfo `json:"ProfileInfo,omitnil,omitempty" name:"ProfileInfo"`
}