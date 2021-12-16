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

package v20170320

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

type Account struct {

	// New account name
	User *string `json:"User,omitempty" name:"User"`

	// New account domain name
	Host *string `json:"Host,omitempty" name:"Host"`
}

type AccountInfo struct {

	// Account remarks
	Notes *string `json:"Notes,omitempty" name:"Notes"`

	// Account domain name
	Host *string `json:"Host,omitempty" name:"Host"`

	// Account name
	User *string `json:"User,omitempty" name:"User"`

	// Account information modification time
	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`

	// Password modification time
	ModifyPasswordTime *string `json:"ModifyPasswordTime,omitempty" name:"ModifyPasswordTime"`

	// This parameter is no longer supported.
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// The maximum number of instance connections supported by an account
	MaxUserConnections *int64 `json:"MaxUserConnections,omitempty" name:"MaxUserConnections"`
}

type AddTimeWindowRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cdb-c1nl9rpv or cdbro-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Maintenance window on Monday. The format should be 10:00-12:00. You can set multiple time windows on a day. Each time window lasts from half an hour to three hours, and must start and end on the hour or half hour. At least one time window is required in a week. The same rule applies to the following parameters.
	Monday []*string `json:"Monday,omitempty" name:"Monday"`

	// Maintenance window on Tuesday. At least one time window is required in a week.
	Tuesday []*string `json:"Tuesday,omitempty" name:"Tuesday"`

	// Maintenance window on Wednesday. At least one time window is required in a week.
	Wednesday []*string `json:"Wednesday,omitempty" name:"Wednesday"`

	// Maintenance window on Thursday. At least one time window is required in a week.
	Thursday []*string `json:"Thursday,omitempty" name:"Thursday"`

	// Maintenance window on Friday. At least one time window is required in a week.
	Friday []*string `json:"Friday,omitempty" name:"Friday"`

	// Maintenance window on Saturday. At least one time window is required in a week.
	Saturday []*string `json:"Saturday,omitempty" name:"Saturday"`

	// Maintenance window on Sunday. At least one time window is required in a week.
	Sunday []*string `json:"Sunday,omitempty" name:"Sunday"`
}

func (r *AddTimeWindowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddTimeWindowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Monday")
	delete(f, "Tuesday")
	delete(f, "Wednesday")
	delete(f, "Thursday")
	delete(f, "Friday")
	delete(f, "Saturday")
	delete(f, "Sunday")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddTimeWindowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type AddTimeWindowResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddTimeWindowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddTimeWindowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssociateSecurityGroupsRequest struct {
	*tchttp.BaseRequest

	// Security group ID.
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`

	// List of instance IDs, which is an array of one or more instance IDs.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// This parameter takes effect only when the IDs of read-only replicas are passed in. If this parameter is set to `False` or left empty, the security group will be bound to the RO groups of these read-only replicas. If this parameter is set to `True`, the security group will be bound to the read-only replicas themselves.
	ForReadonlyInstance *bool `json:"ForReadonlyInstance,omitempty" name:"ForReadonlyInstance"`
}

func (r *AssociateSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateSecurityGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecurityGroupId")
	delete(f, "InstanceIds")
	delete(f, "ForReadonlyInstance")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AssociateSecurityGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type AssociateSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AssociateSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateSecurityGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BackupConfig struct {

	// Replication mode of secondary database 2. Value range: async, semi-sync
	ReplicationMode *string `json:"ReplicationMode,omitempty" name:"ReplicationMode"`

	// Name of the AZ of secondary database 2, such as ap-shanghai-1
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// Private IP address of secondary database 2
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// Access port of secondary database 2
	Vport *uint64 `json:"Vport,omitempty" name:"Vport"`
}

type BackupInfo struct {

	// Backup filename
	Name *string `json:"Name,omitempty" name:"Name"`

	// Backup file size in bytes
	Size *int64 `json:"Size,omitempty" name:"Size"`

	// Backup snapshot time in the format of yyyy-MM-dd HH:mm:ss, such as 2016-03-17 02:10:37
	Date *string `json:"Date,omitempty" name:"Date"`

	// Download address
	IntranetUrl *string `json:"IntranetUrl,omitempty" name:"IntranetUrl"`

	// Download address
	InternetUrl *string `json:"InternetUrl,omitempty" name:"InternetUrl"`

	// Log type. Valid values: `logical` (logical cold backup), `physical` (physical cold backup).
	Type *string `json:"Type,omitempty" name:"Type"`

	// Backup subtask ID, which is used when backup files are deleted
	BackupId *int64 `json:"BackupId,omitempty" name:"BackupId"`

	// Backup task status. Valid values: `SUCCESS` (backup succeeded), `FAILED` (backup failed), `RUNNING` (backup is in progress).
	Status *string `json:"Status,omitempty" name:"Status"`

	// Backup task completion time
	FinishTime *string `json:"FinishTime,omitempty" name:"FinishTime"`

	// (This field will be disused and is thus not recommended) backup creator. Valid values: `SYSTEM` (created by system), `Uin` (initiator's `Uin` value).
	Creator *string `json:"Creator,omitempty" name:"Creator"`

	// Backup task start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Backup method. Valid values: `full` (full backup), `partial` (partial backup).
	Method *string `json:"Method,omitempty" name:"Method"`

	// Backup mode. Valid values: `manual` (manual backup), `automatic` (automatic backup).
	Way *string `json:"Way,omitempty" name:"Way"`

	// Manual backup alias
	ManualBackupName *string `json:"ManualBackupName,omitempty" name:"ManualBackupName"`
}

type BackupItem struct {

	// Name of the database to be backed up
	Db *string `json:"Db,omitempty" name:"Db"`

	// Name of the table to be backed up. If this parameter is passed in, the specified table in the database will be backed up; otherwise, the database will be backed up.
	Table *string `json:"Table,omitempty" name:"Table"`
}

type BackupLimitVpcItem struct {

	// The region where the backup download restrictions take effect. It must be the same as the common request parameter `Region` of the API.
	Region *string `json:"Region,omitempty" name:"Region"`

	// The list of VPCs used to restrict backup download
	VpcList []*string `json:"VpcList,omitempty" name:"VpcList"`
}

type BackupSummaryItem struct {

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Number of automatic data backups of an instance.
	AutoBackupCount *int64 `json:"AutoBackupCount,omitempty" name:"AutoBackupCount"`

	// Capacity of automatic data backups of an instance.
	AutoBackupVolume *int64 `json:"AutoBackupVolume,omitempty" name:"AutoBackupVolume"`

	// Number of manual data backups of an instance.
	ManualBackupCount *int64 `json:"ManualBackupCount,omitempty" name:"ManualBackupCount"`

	// Capacity of manual data backups of an instance.
	ManualBackupVolume *int64 `json:"ManualBackupVolume,omitempty" name:"ManualBackupVolume"`

	// Total number of data backups of an instance (including automatic backups and manual backups).
	DataBackupCount *int64 `json:"DataBackupCount,omitempty" name:"DataBackupCount"`

	// Total capacity of data backups of an instance.
	DataBackupVolume *int64 `json:"DataBackupVolume,omitempty" name:"DataBackupVolume"`

	// Number of log backups of an instance.
	BinlogBackupCount *int64 `json:"BinlogBackupCount,omitempty" name:"BinlogBackupCount"`

	// Capacity of log backups of an instance.
	BinlogBackupVolume *int64 `json:"BinlogBackupVolume,omitempty" name:"BinlogBackupVolume"`

	// Total capacity of backups of an instance (including data backups and log backups).
	BackupVolume *int64 `json:"BackupVolume,omitempty" name:"BackupVolume"`
}

type BalanceRoGroupLoadRequest struct {
	*tchttp.BaseRequest

	// RO group ID in the format of `cdbrg-c1nl9rpv`.
	RoGroupId *string `json:"RoGroupId,omitempty" name:"RoGroupId"`
}

func (r *BalanceRoGroupLoadRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BalanceRoGroupLoadRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RoGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BalanceRoGroupLoadRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type BalanceRoGroupLoadResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BalanceRoGroupLoadResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BalanceRoGroupLoadResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BinlogInfo struct {

	// Binlog backup filename
	Name *string `json:"Name,omitempty" name:"Name"`

	// Backup file size in bytes
	Size *int64 `json:"Size,omitempty" name:"Size"`

	// File stored time in the format of 2016-03-17 02:10:37
	Date *string `json:"Date,omitempty" name:"Date"`

	// Download address
	IntranetUrl *string `json:"IntranetUrl,omitempty" name:"IntranetUrl"`

	// Download address
	InternetUrl *string `json:"InternetUrl,omitempty" name:"InternetUrl"`

	// Log type. Value range: binlog
	Type *string `json:"Type,omitempty" name:"Type"`

	// Binlog file start file
	BinlogStartTime *string `json:"BinlogStartTime,omitempty" name:"BinlogStartTime"`

	// Binlog file end time
	BinlogFinishTime *string `json:"BinlogFinishTime,omitempty" name:"BinlogFinishTime"`
}

type CloneItem struct {

	// ID of the original instance in a clone task
	SrcInstanceId *string `json:"SrcInstanceId,omitempty" name:"SrcInstanceId"`

	// ID of the cloned instance in a clone task
	DstInstanceId *string `json:"DstInstanceId,omitempty" name:"DstInstanceId"`

	// Clone task ID
	CloneJobId *int64 `json:"CloneJobId,omitempty" name:"CloneJobId"`

	// The policy used in a clone task. Valid values: `timepoint` (roll back to a specific point in time), `backupset` (roll back by using a specific backup file).
	RollbackStrategy *string `json:"RollbackStrategy,omitempty" name:"RollbackStrategy"`

	// The point in time to which the cloned instance will be rolled back
	RollbackTargetTime *string `json:"RollbackTargetTime,omitempty" name:"RollbackTargetTime"`

	// Task start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Task end time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Task status. Valid values: `initial`, `running`, `wait_complete`, `success`, `failed`.
	TaskStatus *string `json:"TaskStatus,omitempty" name:"TaskStatus"`
}

type CloseWanServiceRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page. You can use the [instance list querying API](https://intl.cloud.tencent.com/document/api/236/15872?from_cn_redirect=1) to query the ID, whose value is the `InstanceId` value in output parameters.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *CloseWanServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseWanServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CloseWanServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CloseWanServiceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Async task request ID, which can be used to query the execution result of an async task.
		AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CloseWanServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseWanServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ColumnPrivilege struct {

	// Database name
	Database *string `json:"Database,omitempty" name:"Database"`

	// Table name
	Table *string `json:"Table,omitempty" name:"Table"`

	// Column name
	Column *string `json:"Column,omitempty" name:"Column"`

	// Permission information
	Privileges []*string `json:"Privileges,omitempty" name:"Privileges"`
}

type CommonTimeWindow struct {

	// Time window on Monday in the format of 02:00-06:00
	Monday *string `json:"Monday,omitempty" name:"Monday"`

	// Time window on Tuesday in the format of 02:00-06:00
	Tuesday *string `json:"Tuesday,omitempty" name:"Tuesday"`

	// Time window on Wednesday in the format of 02:00-06:00
	Wednesday *string `json:"Wednesday,omitempty" name:"Wednesday"`

	// Time window on Thursday in the format of 02:00-06:00
	Thursday *string `json:"Thursday,omitempty" name:"Thursday"`

	// Time window on Friday in the format of 02:00-06:00
	Friday *string `json:"Friday,omitempty" name:"Friday"`

	// Time window on Saturday in the format of 02:00-06:00
	Saturday *string `json:"Saturday,omitempty" name:"Saturday"`

	// Time window on Sunday in the format of 02:00-06:00
	Sunday *string `json:"Sunday,omitempty" name:"Sunday"`
}

type CreateAccountsRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// TencentDB account.
	Accounts []*Account `json:"Accounts,omitempty" name:"Accounts"`

	// Password of the new account
	Password *string `json:"Password,omitempty" name:"Password"`

	// Remarks
	Description *string `json:"Description,omitempty" name:"Description"`

	// Maximum connections of the new account. Default value: `10240`. Maximum value: `10240`.
	MaxUserConnections *int64 `json:"MaxUserConnections,omitempty" name:"MaxUserConnections"`
}

func (r *CreateAccountsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAccountsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Accounts")
	delete(f, "Password")
	delete(f, "Description")
	delete(f, "MaxUserConnections")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAccountsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateAccountsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Async task request ID, which can be used to query the execution result of an async task.
		AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAccountsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAccountsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAuditPolicyRequest struct {
	*tchttp.BaseRequest

	// Audit policy name.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Audit rule ID.
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// Instance ID in the format of cdb-c1nl9rpv or cdbro-c1nl9rpv. It is the same as the instance ID displayed in the TencentDB console.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Retention period of audit logs. Valid values:
	// 7: seven days (a week);
	// 30: 30 days (a month);
	// 180: 180 days (six months);
	// 365: 365 days (a year);
	// 1095: 1095 days (three years);
	// 1825: 1825 days (five years).
	// This parameter specifies the retention period (30 days by default) of audit logs, which is valid when you create the first audit policy for an instance. If the instance already has audit policies, this parameter is invalid, but you can use the `ModifyAuditConfig` API to modify the retention period.
	LogExpireDay *int64 `json:"LogExpireDay,omitempty" name:"LogExpireDay"`
}

func (r *CreateAuditPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAuditPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "RuleId")
	delete(f, "InstanceId")
	delete(f, "LogExpireDay")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAuditPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateAuditPolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Audit policy ID.
		PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAuditPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAuditPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateBackupRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Backup method of target instance. Value range: logical (logical cold backup), physical (physical cold backup).
	BackupMethod *string `json:"BackupMethod,omitempty" name:"BackupMethod"`

	// Information of the table to be backed up. If this parameter is not set, the entire instance will be backed up by default. It can be set only in logical backup (i.e., BackupMethod = logical). The specified table must exist; otherwise, backup may fail.
	// For example, if you want to backup tb1 and tb2 in db1 and the entire db2, you should set the parameter as [{"Db": "db1", "Table": "tb1"}, {"Db": "db1", "Table": "tb2"}, {"Db": "db2"} ].
	BackupDBTableList []*BackupItem `json:"BackupDBTableList,omitempty" name:"BackupDBTableList"`

	// Manual backup alias
	ManualBackupName *string `json:"ManualBackupName,omitempty" name:"ManualBackupName"`
}

func (r *CreateBackupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBackupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BackupMethod")
	delete(f, "BackupDBTableList")
	delete(f, "ManualBackupName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBackupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateBackupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Backup task ID
		BackupId *uint64 `json:"BackupId,omitempty" name:"BackupId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateBackupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBackupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCloneInstanceRequest struct {
	*tchttp.BaseRequest

	// ID of the instance to be cloned from
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// To roll back the cloned instance to a specific point in time, set this parameter to a value in the format of "yyyy-mm-dd hh:mm:ss".
	SpecifiedRollbackTime *string `json:"SpecifiedRollbackTime,omitempty" name:"SpecifiedRollbackTime"`

	// To roll back the cloned instance to a specific physical backup file, set this parameter to the ID of the physical backup file. The ID can be obtained by the [DescribeBackups](https://intl.cloud.tencent.com/document/api/236/15842?from_cn_redirect=1) API.
	SpecifiedBackupId *int64 `json:"SpecifiedBackupId,omitempty" name:"SpecifiedBackupId"`

	// VPC ID, which can be obtained by the [DescribeVpcs](https://intl.cloud.tencent.com/document/api/215/15778?from_cn_redirect=1) API. If this parameter is left empty, the classic network will be used by default.
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// VPC subnet ID, which can be obtained by the [DescribeSubnets](https://intl.cloud.tencent.com/document/api/215/15784?from_cn_redirect=1) API. If `UniqVpcId` is set, `UniqSubnetId` will be required.
	UniqSubnetId *string `json:"UniqSubnetId,omitempty" name:"UniqSubnetId"`

	// Memory of the cloned instance in MB, which should be equal to (by default) or larger than that of the original instance
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// Disk capacity of the cloned instance in GB, which should be equal to (by default) or larger than that of the original instance
	Volume *int64 `json:"Volume,omitempty" name:"Volume"`

	// Name of the cloned instance
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// Security group parameter, which can be obtained by the [DescribeProjectSecurityGroups](https://intl.cloud.tencent.com/document/api/236/15850?from_cn_redirect=1) API
	SecurityGroup []*string `json:"SecurityGroup,omitempty" name:"SecurityGroup"`

	// Information of the cloned instance tag
	ResourceTags []*TagInfo `json:"ResourceTags,omitempty" name:"ResourceTags"`

	// The number of CPU cores of the cloned instance. It should be equal to (by default) or larger than that of the original instance.
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

	// Data replication mode. Valid values: 0 (async), 1 (semi-sync), 2 (strong sync). Default value: 0.
	ProtectMode *int64 `json:"ProtectMode,omitempty" name:"ProtectMode"`

	// Multi-AZ or single-AZ. Valid values: 0 (single-AZ), 1 (multi-AZ). Default value: 0.
	DeployMode *int64 `json:"DeployMode,omitempty" name:"DeployMode"`

	// Availability zone information of replica 1 of the cloned instance, which is the same as the value of `Zone` of the original instance by default
	SlaveZone *string `json:"SlaveZone,omitempty" name:"SlaveZone"`

	// Availability zone information of replica 2 of the cloned instance, 
	// which is left empty by default. Specify this parameter when cloning a strong sync source instance.
	BackupZone *string `json:"BackupZone,omitempty" name:"BackupZone"`

	// Resource isolation type of the clone. Valid values: `UNIVERSAL` (general instance), `EXCLUSIVE` (dedicated instance). Default value: `UNIVERSAL`.
	DeviceType *string `json:"DeviceType,omitempty" name:"DeviceType"`

	// The number of nodes of the clone. If this parameter is set to `3` or the `BackupZone` parameter is specified, the clone will have three nodes. If this parameter is set to `2` or left empty, the clone will have two nodes.
	InstanceNodes *int64 `json:"InstanceNodes,omitempty" name:"InstanceNodes"`

	// Placement group ID.
	DeployGroupId *string `json:"DeployGroupId,omitempty" name:"DeployGroupId"`

	// Whether to check the request without creating any instance. Valid values: `true`, `false` (default). After being submitted, the request will be checked to see if it is in correct format and has all required parameters with valid values. An error code is returned if the check failed, and `RequestId` is returned if the check succeeded. After a successful check, no instance will be created if this parameter is set to `true`, whereas an instance will be created and if it is set to `false`.
	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`
}

func (r *CreateCloneInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCloneInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SpecifiedRollbackTime")
	delete(f, "SpecifiedBackupId")
	delete(f, "UniqVpcId")
	delete(f, "UniqSubnetId")
	delete(f, "Memory")
	delete(f, "Volume")
	delete(f, "InstanceName")
	delete(f, "SecurityGroup")
	delete(f, "ResourceTags")
	delete(f, "Cpu")
	delete(f, "ProtectMode")
	delete(f, "DeployMode")
	delete(f, "SlaveZone")
	delete(f, "BackupZone")
	delete(f, "DeviceType")
	delete(f, "InstanceNodes")
	delete(f, "DeployGroupId")
	delete(f, "DryRun")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCloneInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateCloneInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// LimitAsync task request ID, which can be used to query the execution result of an async task
		AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCloneInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCloneInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDBImportJobRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Filename. The file must be a .sql file uploaded to Tencent Cloud.
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// TencentDB username
	User *string `json:"User,omitempty" name:"User"`

	// Password of a TencentDB instance user account
	Password *string `json:"Password,omitempty" name:"Password"`

	// Name of the target database. If this parameter is not passed in, no database is specified.
	DbName *string `json:"DbName,omitempty" name:"DbName"`
}

func (r *CreateDBImportJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDBImportJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "FileName")
	delete(f, "User")
	delete(f, "Password")
	delete(f, "DbName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDBImportJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateDBImportJobResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Async task request ID, which can be used to query the execution result of an async task.
		AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDBImportJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDBImportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDBInstanceHourRequest struct {
	*tchttp.BaseRequest

	// Number of instances. Value range: 1-100. Default value: 1.
	GoodsNum *int64 `json:"GoodsNum,omitempty" name:"GoodsNum"`

	// Instance memory size in MB. Please use the [DescribeDBZoneConfig](https://intl.cloud.tencent.com/document/api/236/17229?from_cn_redirect=1) API to query the supported memory specifications.
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// Instance disk size in GB. Please use the [DescribeDBZoneConfig](https://intl.cloud.tencent.com/document/api/236/17229?from_cn_redirect=1) API to query the supported disk specifications.
	Volume *int64 `json:"Volume,omitempty" name:"Volume"`

	// MySQL version. Valid values: `5.5`, `5.6`, `5.7`, `8.0`. You can use the [DescribeDBZoneConfig](https://intl.cloud.tencent.com/document/api/236/17229?from_cn_redirect=1) API to query the supported versions.
	EngineVersion *string `json:"EngineVersion,omitempty" name:"EngineVersion"`

	// VPC ID. If this parameter is not passed in, the basic network will be selected by default. Please use the [DescribeVpcs](https://intl.cloud.tencent.com/document/api/215/15778?from_cn_redirect=1) API to query the VPCs.
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// VPC subnet ID. If `UniqVpcId` is set, then `UniqSubnetId` will be required. Please use the [DescribeSubnets](https://intl.cloud.tencent.com/document/api/215/15784?from_cn_redirect=1) API to query the subnet lists.
	UniqSubnetId *string `json:"UniqSubnetId,omitempty" name:"UniqSubnetId"`

	// Project ID. If this is left empty, the default project will be used.
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// AZ information. By default, the system will automatically select an AZ. Please use the [DescribeDBZoneConfig](https://intl.cloud.tencent.com/document/api/236/17229?from_cn_redirect=1) API to query the supported AZs.
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// Instance ID, which is required and the same as the primary instance ID when purchasing read-only or disaster recovery instances. Please use the [DescribeDBInstances](https://intl.cloud.tencent.com/document/api/236/15872?from_cn_redirect=1) API to query the instance IDs.
	MasterInstanceId *string `json:"MasterInstanceId,omitempty" name:"MasterInstanceId"`

	// Instance type. Valid values: master (primary instance), dr (disaster recovery instance), ro (read-only instance). Default value: master.
	InstanceRole *string `json:"InstanceRole,omitempty" name:"InstanceRole"`

	// AZ information of the primary instance, which is required for purchasing disaster recovery instances.
	MasterRegion *string `json:"MasterRegion,omitempty" name:"MasterRegion"`

	// Custom port. Value range: [1024-65535].
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// Sets the root account password. Rule: the password can contain 8-64 characters and must contain at least two of the following types of characters: letters, digits, and special symbols (_+-&=!@#$%^*()). This parameter can be specified when purchasing primary instances and is meaningless for read-only or disaster recovery instances.
	Password *string `json:"Password,omitempty" name:"Password"`

	// List of parameters in the format of `ParamList.0.Name=auto_increment&ParamList.0.Value=1`. You can use the [DescribeDefaultParams](https://intl.cloud.tencent.com/document/api/236/32662?from_cn_redirect=1) API to query the configurable parameters.
	ParamList []*ParamInfo `json:"ParamList,omitempty" name:"ParamList"`

	// Data replication mode. Valid values: 0 (async), 1 (semi-sync), 2 (strong sync). Default value: 0. This parameter can be specified when purchasing primary instances and is meaningless for read-only or disaster recovery instances.
	ProtectMode *int64 `json:"ProtectMode,omitempty" name:"ProtectMode"`

	// Multi-AZ. Valid value: 0 (single-AZ), 1 (multi-AZ). Default value: 0. This parameter can be specified when purchasing primary instances and is meaningless for read-only or disaster recovery instances.
	DeployMode *int64 `json:"DeployMode,omitempty" name:"DeployMode"`

	// AZ information of secondary database 1, which is the `Zone` value by default. This parameter can be specified when purchasing primary instances and is meaningless for read-only or disaster recovery instances.
	SlaveZone *string `json:"SlaveZone,omitempty" name:"SlaveZone"`

	// The availability zone information of Replica 2, which is left empty by default. Specify this parameter when purchasing a source instance in the one-source-two-replica architecture.
	BackupZone *string `json:"BackupZone,omitempty" name:"BackupZone"`

	// Security group parameter. You can use the [DescribeProjectSecurityGroups](https://intl.cloud.tencent.com/document/api/236/15850?from_cn_redirect=1) API to query the security group details of a project.
	SecurityGroup []*string `json:"SecurityGroup,omitempty" name:"SecurityGroup"`

	// Read-only instance information. This parameter must be passed in when purchasing read-only instances.
	RoGroup *RoGroup `json:"RoGroup,omitempty" name:"RoGroup"`

	// This field is meaningless when purchasing pay-as-you-go instances.
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`

	// Instance name.
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// Instance tag information.
	ResourceTags []*TagInfo `json:"ResourceTags,omitempty" name:"ResourceTags"`

	// Placement group ID.
	DeployGroupId *string `json:"DeployGroupId,omitempty" name:"DeployGroupId"`

	// A string that is used to guarantee the idempotency of the request, which is generated by the user and must be unique in each request on the same day. The maximum length is 64 ASCII characters. If this parameter is not specified, the idempotency of the request cannot be guaranteed.
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`

	// Instance resource isolation type. Valid values: `UNIVERSAL` (general instance), `EXCLUSIVE` (dedicated instance), `BASIC` (basic instance). Default value: `UNIVERSAL`.
	DeviceType *string `json:"DeviceType,omitempty" name:"DeviceType"`

	// Parameter template ID.
	ParamTemplateId *int64 `json:"ParamTemplateId,omitempty" name:"ParamTemplateId"`

	// The array of alarm policy IDs.
	AlarmPolicyList []*int64 `json:"AlarmPolicyList,omitempty" name:"AlarmPolicyList"`

	// The number of nodes of the instance. To purchase a read-only replica or a basic instance, set this parameter to `1` or leave it empty. To purchase a three-node instance, set this parameter to `3` or specify the `BackupZone` parameter. If the instance to be purchased is a source instance and both `BackupZone` and this parameter are left empty, the value `2` will be used, which indicates the source instance will have two nodes.
	InstanceNodes *int64 `json:"InstanceNodes,omitempty" name:"InstanceNodes"`

	// The number of CPU cores of the instance. If this parameter is left empty, the number of CPU cores depends on the `Memory` value.
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

	// Whether to automatically start disaster recovery synchronization. This parameter takes effect only for disaster recovery instances. Valid values: `0` (no), `1` (yes).
	AutoSyncFlag *int64 `json:"AutoSyncFlag,omitempty" name:"AutoSyncFlag"`

	// Financial cage ID.
	CageId *string `json:"CageId,omitempty" name:"CageId"`

	// The array of alarm policy names, such as ["policy-uyoee9wg"]. If the `AlarmPolicyList` parameter is specified, this parameter is invalid.
	AlarmPolicyIdList []*string `json:"AlarmPolicyIdList,omitempty" name:"AlarmPolicyIdList"`

	// Whether to check the request without creating any instance. Valid values: `true`, `false` (default). After being submitted, the request will be checked to see if it is in correct format and has all required parameters with valid values. An error code is returned if the check failed, and `RequestId` is returned if the check succeeded. After a successful check, no instance will be created if this parameter is set to `true`, whereas an instance will be created and if it is set to `false`.
	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`
}

func (r *CreateDBInstanceHourRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDBInstanceHourRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GoodsNum")
	delete(f, "Memory")
	delete(f, "Volume")
	delete(f, "EngineVersion")
	delete(f, "UniqVpcId")
	delete(f, "UniqSubnetId")
	delete(f, "ProjectId")
	delete(f, "Zone")
	delete(f, "MasterInstanceId")
	delete(f, "InstanceRole")
	delete(f, "MasterRegion")
	delete(f, "Port")
	delete(f, "Password")
	delete(f, "ParamList")
	delete(f, "ProtectMode")
	delete(f, "DeployMode")
	delete(f, "SlaveZone")
	delete(f, "BackupZone")
	delete(f, "SecurityGroup")
	delete(f, "RoGroup")
	delete(f, "AutoRenewFlag")
	delete(f, "InstanceName")
	delete(f, "ResourceTags")
	delete(f, "DeployGroupId")
	delete(f, "ClientToken")
	delete(f, "DeviceType")
	delete(f, "ParamTemplateId")
	delete(f, "AlarmPolicyList")
	delete(f, "InstanceNodes")
	delete(f, "Cpu")
	delete(f, "AutoSyncFlag")
	delete(f, "CageId")
	delete(f, "AlarmPolicyIdList")
	delete(f, "DryRun")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDBInstanceHourRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateDBInstanceHourResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Short order ID.
		DealIds []*string `json:"DealIds,omitempty" name:"DealIds"`

		// Instance ID list
		InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDBInstanceHourResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDBInstanceHourResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDeployGroupRequest struct {
	*tchttp.BaseRequest

	// Name of a placement group, which can contain up to 60 characters.
	DeployGroupName *string `json:"DeployGroupName,omitempty" name:"DeployGroupName"`

	// Description of a placement group, which can contain up to 200 characters.
	Description *string `json:"Description,omitempty" name:"Description"`

	// Affinity policy of placement group. Currently, the value of this parameter can only be 1. Policy 1 indicates the upper limit of instances on one physical machine.
	Affinity []*int64 `json:"Affinity,omitempty" name:"Affinity"`

	// Upper limit of instances on one physical machine as defined in affinity policy 1 of placement group.
	LimitNum *int64 `json:"LimitNum,omitempty" name:"LimitNum"`

	// Model attribute of placement group. Valid values: SH12+SH02, TS85.
	DevClass []*string `json:"DevClass,omitempty" name:"DevClass"`
}

func (r *CreateDeployGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDeployGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeployGroupName")
	delete(f, "Description")
	delete(f, "Affinity")
	delete(f, "LimitNum")
	delete(f, "DevClass")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDeployGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateDeployGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Placement group ID.
		DeployGroupId *string `json:"DeployGroupId,omitempty" name:"DeployGroupId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDeployGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDeployGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateParamTemplateRequest struct {
	*tchttp.BaseRequest

	// Parameter template name.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Parameter template description.
	Description *string `json:"Description,omitempty" name:"Description"`

	// MySQL version number.
	EngineVersion *string `json:"EngineVersion,omitempty" name:"EngineVersion"`

	// Source parameter template ID.
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// List of parameters.
	ParamList []*Parameter `json:"ParamList,omitempty" name:"ParamList"`
}

func (r *CreateParamTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateParamTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Description")
	delete(f, "EngineVersion")
	delete(f, "TemplateId")
	delete(f, "ParamList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateParamTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateParamTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Parameter template ID.
		TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateParamTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateParamTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateRoInstanceIpRequest struct {
	*tchttp.BaseRequest

	// Read-only instance ID in the format of "cdbro-3i70uj0k". Its value is the same as the read-only instance ID in the TencentDB Console.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Subnet descriptor, such as "subnet-1typ0s7d".
	UniqSubnetId *string `json:"UniqSubnetId,omitempty" name:"UniqSubnetId"`

	// VPC descriptor, such as "vpc-a23yt67j". If this field is passed in, `UniqSubnetId` will be required.
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`
}

func (r *CreateRoInstanceIpRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRoInstanceIpRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "UniqSubnetId")
	delete(f, "UniqVpcId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRoInstanceIpRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateRoInstanceIpResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// VPC ID of the read-only instance.
		RoVpcId *int64 `json:"RoVpcId,omitempty" name:"RoVpcId"`

		// Subnet ID of the read-only instance.
		RoSubnetId *int64 `json:"RoSubnetId,omitempty" name:"RoSubnetId"`

		// Private IP address of the read-only instance.
		RoVip *string `json:"RoVip,omitempty" name:"RoVip"`

		// Private port number of the read-only instance.
		RoVport *int64 `json:"RoVport,omitempty" name:"RoVport"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateRoInstanceIpResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRoInstanceIpResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DBSwitchInfo struct {

	// Switch time in the format of yyyy-MM-dd HH:mm:ss, such as 2017-09-03 01:34:31
	SwitchTime *string `json:"SwitchTime,omitempty" name:"SwitchTime"`

	// Switch type. Value range: TRANSFER (data migration), MASTER2SLAVE (primary/secondary switch), RECOVERY (primary/secondary recovery)
	SwitchType *string `json:"SwitchType,omitempty" name:"SwitchType"`
}

type DatabasePrivilege struct {

	// Permission information
	Privileges []*string `json:"Privileges,omitempty" name:"Privileges"`

	// Database name
	Database *string `json:"Database,omitempty" name:"Database"`
}

type DatabasesWithCharacterLists struct {

	// Database name
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`

	// Character set
	CharacterSet *string `json:"CharacterSet,omitempty" name:"CharacterSet"`
}

type DeleteAccountsRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// TencentDB account.
	Accounts []*Account `json:"Accounts,omitempty" name:"Accounts"`
}

func (r *DeleteAccountsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAccountsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Accounts")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAccountsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAccountsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Async task request ID, which can be used to query the execution result of an async task.
		AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAccountsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAccountsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteBackupRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Backup task ID, which is the task ID returned by the [TencentDB instance backup creating API](https://intl.cloud.tencent.com/document/api/236/15844?from_cn_redirect=1).
	BackupId *int64 `json:"BackupId,omitempty" name:"BackupId"`
}

func (r *DeleteBackupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBackupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BackupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteBackupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteBackupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteBackupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBackupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDeployGroupsRequest struct {
	*tchttp.BaseRequest

	// List of IDs of placement groups to be deleted.
	DeployGroupIds []*string `json:"DeployGroupIds,omitempty" name:"DeployGroupIds"`
}

func (r *DeleteDeployGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDeployGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeployGroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDeployGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDeployGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteDeployGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDeployGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteParamTemplateRequest struct {
	*tchttp.BaseRequest

	// Parameter template ID.
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
}

func (r *DeleteParamTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteParamTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteParamTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteParamTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteParamTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteParamTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteTimeWindowRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cdb-c1nl9rpv or cdbro-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DeleteTimeWindowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTimeWindowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTimeWindowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteTimeWindowResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteTimeWindowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTimeWindowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeployGroupInfo struct {

	// ID of a placement group.
	DeployGroupId *string `json:"DeployGroupId,omitempty" name:"DeployGroupId"`

	// Name of a placement group.
	DeployGroupName *string `json:"DeployGroupName,omitempty" name:"DeployGroupName"`

	// Creation time.
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Instance quota of placement group, indicating the maximum number of instances that can be placed in one placement group.
	Quota *int64 `json:"Quota,omitempty" name:"Quota"`

	// Affinity policy of placement group. Currently, only policy 1 is supported, indicating to distribute instances across physical machines.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Affinity *string `json:"Affinity,omitempty" name:"Affinity"`

	// Upper limit of instances in one placement group on one physical machine as defined in affinity policy 1 of placement group.
	// Note: this field may return null, indicating that no valid values can be obtained.
	LimitNum *int64 `json:"LimitNum,omitempty" name:"LimitNum"`

	// Placement group details.
	Description *string `json:"Description,omitempty" name:"Description"`

	// Physical model attribute of placement group.
	// Note: this field may return null, indicating that no valid values can be obtained.
	DevClass *string `json:"DevClass,omitempty" name:"DevClass"`
}

type DescribeAccountPrivilegesRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Database user account.
	User *string `json:"User,omitempty" name:"User"`

	// Database account domain name.
	Host *string `json:"Host,omitempty" name:"Host"`
}

func (r *DescribeAccountPrivilegesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccountPrivilegesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "User")
	delete(f, "Host")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccountPrivilegesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountPrivilegesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Array of global permissions.
		GlobalPrivileges []*string `json:"GlobalPrivileges,omitempty" name:"GlobalPrivileges"`

		// Array of database permissions.
		DatabasePrivileges []*DatabasePrivilege `json:"DatabasePrivileges,omitempty" name:"DatabasePrivileges"`

		// Array of table permissions in the database.
		TablePrivileges []*TablePrivilege `json:"TablePrivileges,omitempty" name:"TablePrivileges"`

		// Array of column permissions in the table.
		ColumnPrivileges []*ColumnPrivilege `json:"ColumnPrivileges,omitempty" name:"ColumnPrivileges"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccountPrivilegesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccountPrivilegesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountsRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Record offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of results to be returned for a single request. Value range: 1-100. Default value: 20.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Regular expression for matching account names, which complies with the rules at MySQL official website.
	AccountRegexp *string `json:"AccountRegexp,omitempty" name:"AccountRegexp"`
}

func (r *DescribeAccountsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccountsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "AccountRegexp")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccountsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of eligible accounts.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Details of eligible accounts.
		Items []*AccountInfo `json:"Items,omitempty" name:"Items"`

		// The maximum number of instance connections (set by the MySQL parameter `max_connections`)
		MaxUserConnections *int64 `json:"MaxUserConnections,omitempty" name:"MaxUserConnections"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccountsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccountsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAsyncRequestInfoRequest struct {
	*tchttp.BaseRequest

	// Async task request ID.
	AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`
}

func (r *DescribeAsyncRequestInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAsyncRequestInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AsyncRequestId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAsyncRequestInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAsyncRequestInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Task execution result. Valid values: INITIAL, RUNNING, SUCCESS, FAILED, KILLED, REMOVED, PAUSED.
	// Note: This field may return null, indicating that no valid values can be obtained.
		Status *string `json:"Status,omitempty" name:"Status"`

		// Task execution information.
	// Note: This field may return null, indicating that no valid values can be obtained.
		Info *string `json:"Info,omitempty" name:"Info"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAsyncRequestInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAsyncRequestInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBackupConfigRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeBackupConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackupConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBackupConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Earliest start time point of automatic backup, such as 2 (for 2:00 AM). (This field has been disused. You are recommended to use the `BackupTimeWindow` field)
		StartTimeMin *int64 `json:"StartTimeMin,omitempty" name:"StartTimeMin"`

		// Latest start time point of automatic backup, such as 6 (for 6:00 AM). (This field has been disused. You are recommended to use the `BackupTimeWindow` field)
		StartTimeMax *int64 `json:"StartTimeMax,omitempty" name:"StartTimeMax"`

		// Backup file retention period in days.
		BackupExpireDays *int64 `json:"BackupExpireDays,omitempty" name:"BackupExpireDays"`

		// Backup mode. Value range: physical, logical
		BackupMethod *string `json:"BackupMethod,omitempty" name:"BackupMethod"`

		// Binlog file retention period in days.
		BinlogExpireDays *int64 `json:"BinlogExpireDays,omitempty" name:"BinlogExpireDays"`

		// Time window for automatic instance backup.
		BackupTimeWindow *CommonTimeWindow `json:"BackupTimeWindow,omitempty" name:"BackupTimeWindow"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBackupConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBackupDownloadRestrictionRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeBackupDownloadRestrictionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupDownloadRestrictionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackupDownloadRestrictionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBackupDownloadRestrictionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Valid values: `NoLimit` (backups can be downloaded over both private and public networks with any IPs), `LimitOnlyIntranet` (backups can be downloaded over the private network with any private IPs), `Customize` (backups can be downloaded over specified VPCs with specified IPs). The `LimitVpc` and `LimitIp` parameters are valid only when this parameter is set to `Customize`.
		LimitType *string `json:"LimitType,omitempty" name:"LimitType"`

		// Valid value: `In` (backups can only be downloaded over the VPCs specified in `LimitVpc`).
		VpcComparisonSymbol *string `json:"VpcComparisonSymbol,omitempty" name:"VpcComparisonSymbol"`

		// Valid values: `In` (backups can only be downloaded with the IPs specified in `LimitIp`), `NotIn` (backups cannot be downloaded with the IPs specified in `LimitIp`).
		IpComparisonSymbol *string `json:"IpComparisonSymbol,omitempty" name:"IpComparisonSymbol"`

		// VPCs used to restrict backup download.
		LimitVpc []*BackupLimitVpcItem `json:"LimitVpc,omitempty" name:"LimitVpc"`

		// IPs used to restrict backup download.
		LimitIp []*string `json:"LimitIp,omitempty" name:"LimitIp"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBackupDownloadRestrictionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupDownloadRestrictionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBackupOverviewRequest struct {
	*tchttp.BaseRequest

	// TencentDB product type to be queried. Currently, only `mysql` is supported.
	Product *string `json:"Product,omitempty" name:"Product"`
}

func (r *DescribeBackupOverviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupOverviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackupOverviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBackupOverviewResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of backups of a user in the current region (including data backups and log backups).
		BackupCount *int64 `json:"BackupCount,omitempty" name:"BackupCount"`

		// Total capacity of backups of a user in the current region.
		BackupVolume *int64 `json:"BackupVolume,omitempty" name:"BackupVolume"`

		// Paid capacity of backups of a user in the current region, i.e., capacity that exceeds the free tier.
		BillingVolume *int64 `json:"BillingVolume,omitempty" name:"BillingVolume"`

		// Backup capacity in the free tier of a user in the current region.
		FreeVolume *int64 `json:"FreeVolume,omitempty" name:"FreeVolume"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBackupOverviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBackupSummariesRequest struct {
	*tchttp.BaseRequest

	// TencentDB product type to be queried. Currently, only `mysql` is supported.
	Product *string `json:"Product,omitempty" name:"Product"`

	// Pagination offset.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Paginated query limit. Default value: 20.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Sorting criterion. Valid values: BackupVolume (backup capacity), DataBackupVolume (data backup capacity), BinlogBackupVolume (log backup capacity), AutoBackupVolume (automatic backup capacity), ManualBackupVolume (manual backup capacity).
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// Sorting order. Valid values: ASC (ascending), DESC (descending).
	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`
}

func (r *DescribeBackupSummariesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupSummariesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Product")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "OrderBy")
	delete(f, "OrderDirection")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackupSummariesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBackupSummariesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Statistical items of instance backup.
		Items []*BackupSummaryItem `json:"Items,omitempty" name:"Items"`

		// Total number of instance backups.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBackupSummariesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupSummariesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBackupsRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Offset. Minimum value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of entries per page. Value range: 1-100. Default value: 20.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeBackupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBackupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of eligible instances.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Details of eligible backups.
		Items []*BackupInfo `json:"Items,omitempty" name:"Items"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBackupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBinlogBackupOverviewRequest struct {
	*tchttp.BaseRequest

	// TencentDB product type to be queried. Currently, only `mysql` is supported.
	Product *string `json:"Product,omitempty" name:"Product"`
}

func (r *DescribeBinlogBackupOverviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBinlogBackupOverviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBinlogBackupOverviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBinlogBackupOverviewResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total capacity of log backups in bytes.
		BinlogBackupVolume *int64 `json:"BinlogBackupVolume,omitempty" name:"BinlogBackupVolume"`

		// Total number of log backups.
		BinlogBackupCount *int64 `json:"BinlogBackupCount,omitempty" name:"BinlogBackupCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBinlogBackupOverviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBinlogBackupOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBinlogsRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Offset. Minimum value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of entries per page. Value range: 1-100. Default value: 20.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeBinlogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBinlogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBinlogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBinlogsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of eligible log files.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Number of eligible binlog files.
		Items []*BinlogInfo `json:"Items,omitempty" name:"Items"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBinlogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBinlogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCloneListRequest struct {
	*tchttp.BaseRequest

	// ID of the original instance. This parameter is used to query the clone task list of a specific original instance.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Paginated query offset. Default value: `0`.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of results per page. Default value: `20`.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeCloneListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloneListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCloneListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCloneListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The number of results which meet the conditions
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Clone task list
		Items []*CloneItem `json:"Items,omitempty" name:"Items"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCloneListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloneListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBImportRecordsRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Start time in the format of yyyy-MM-dd HH:mm:ss, such as 2016-01-01 00:00:01.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time in the format of yyyy-MM-dd HH:mm:ss, such as 2016-01-01 23:59:59.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Pagination parameter indicating the offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Pagination parameter indicating the number of results to be returned for a single request. Value range: 1-100. Default value: 20.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeDBImportRecordsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBImportRecordsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBImportRecordsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBImportRecordsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of eligible import task operation logs.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// List of import operation records.
		Items []*ImportRecord `json:"Items,omitempty" name:"Items"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBImportRecordsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBImportRecordsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstanceCharsetRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page. You can use the [instance list querying API](https://intl.cloud.tencent.com/document/api/236/15872?from_cn_redirect=1) to query the ID, whose value is the `InstanceId` value in output parameters.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeDBInstanceCharsetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstanceCharsetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBInstanceCharsetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstanceCharsetResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Default character set of the instance, such as "latin1" and "utf8".
		Charset *string `json:"Charset,omitempty" name:"Charset"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBInstanceCharsetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstanceCharsetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstanceConfigRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeDBInstanceConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstanceConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBInstanceConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstanceConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Data protection mode of the primary instance. Value range: 0 (async replication), 1 (semi-sync replication), 2 (strong sync replication).
		ProtectMode *int64 `json:"ProtectMode,omitempty" name:"ProtectMode"`

		// Master instance deployment mode. Value range: 0 (single-AZ), 1 (multi-AZ)
		DeployMode *int64 `json:"DeployMode,omitempty" name:"DeployMode"`

		// Instance AZ information in the format of "ap-shanghai-1".
		Zone *string `json:"Zone,omitempty" name:"Zone"`

		// Configurations of the replica node
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
		SlaveConfig *SlaveConfig `json:"SlaveConfig,omitempty" name:"SlaveConfig"`

		// Configurations of the second replica node of a strong-sync instance
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
		BackupConfig *BackupConfig `json:"BackupConfig,omitempty" name:"BackupConfig"`

		// This parameter is only available for multi-AZ instances. It indicates whether the source AZ is the same as the one specified upon purchase. `true`: not the same, `false`: the same.
		Switched *bool `json:"Switched,omitempty" name:"Switched"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBInstanceConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstanceConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstanceGTIDRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page. You can use the [instance list querying API](https://intl.cloud.tencent.com/document/api/236/15872?from_cn_redirect=1) to query the ID, whose value is the `InstanceId` value in output parameters.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeDBInstanceGTIDRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstanceGTIDRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBInstanceGTIDRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstanceGTIDResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// GTID enablement flag. Value range: 0 (not enabled), 1 (enabled).
		IsGTIDOpen *int64 `json:"IsGTIDOpen,omitempty" name:"IsGTIDOpen"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBInstanceGTIDResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstanceGTIDResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstanceInfoRequest struct {
	*tchttp.BaseRequest

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeDBInstanceInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstanceInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBInstanceInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstanceInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Instance ID.
		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

		// Instance name.
		InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

		// Whether encryption is enabled. YES: enabled, NO: not enabled.
		Encryption *string `json:"Encryption,omitempty" name:"Encryption"`

		// Encryption key ID.
	// Note: this field may return null, indicating that no valid values can be obtained.
		KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

		// Key region.
	// Note: this field may return null, indicating that no valid values can be obtained.
		KeyRegion *string `json:"KeyRegion,omitempty" name:"KeyRegion"`

		// The default region of the KMS service currently used by the TencentDB backend service.
	// Note: this field may return `null`, indicating that no valid value can be found.
		DefaultKmsRegion *string `json:"DefaultKmsRegion,omitempty" name:"DefaultKmsRegion"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBInstanceInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstanceInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstanceRebootTimeRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *DescribeDBInstanceRebootTimeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstanceRebootTimeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBInstanceRebootTimeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstanceRebootTimeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of eligible instances.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Returned parameter information.
		Items []*InstanceRebootTime `json:"Items,omitempty" name:"Items"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBInstanceRebootTimeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstanceRebootTimeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstancesRequest struct {
	*tchttp.BaseRequest

	// Project ID.
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Instance type. Value range: 1 (primary), 2 (disaster recovery), 3 (read-only).
	InstanceTypes []*uint64 `json:"InstanceTypes,omitempty" name:"InstanceTypes"`

	// Private IP address of the instance.
	Vips []*string `json:"Vips,omitempty" name:"Vips"`

	// Instance status. Value range: <br>0 - creating <br>1 - running <br>4 - isolating <br>5 - isolated (the instance can be restored and started in the recycle bin)
	Status []*uint64 `json:"Status,omitempty" name:"Status"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of results to be returned for a single request. Default value: 20. Maximum value: 2,000.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Security group ID. When it is used as a filter, the `WithSecurityGroup` parameter should be set to 1.
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`

	// Billing method. Value range: 0 (monthly subscribed), 1 (hourly).
	PayTypes []*uint64 `json:"PayTypes,omitempty" name:"PayTypes"`

	// Instance name.
	InstanceNames []*string `json:"InstanceNames,omitempty" name:"InstanceNames"`

	// Instance task status. Valid values: <br>0 - no task <br>1 - upgrading <br>2 - importing data <br>3 - enabling secondary instance access <br>4 - enabling public network access <br>5 - batch operation in progress <br>6 - rolling back <br>7 - disabling public network access <br>8 - modifying password <br>9 - renaming instance <br>10 - restarting <br>12 - migrating self-built database <br>13 - dropping tables <br>14 - Disaster recovery instance creating sync task <br>15 - waiting for switch <br>16 - switching <br>17 - upgrade and switch completed <br>19 - parameter settings to be executed
	TaskStatus []*uint64 `json:"TaskStatus,omitempty" name:"TaskStatus"`

	// Version of the instance database engine. Value range: 5.1, 5.5, 5.6, 5.7.
	EngineVersions []*string `json:"EngineVersions,omitempty" name:"EngineVersions"`

	// VPC ID.
	VpcIds []*uint64 `json:"VpcIds,omitempty" name:"VpcIds"`

	// AZ ID.
	ZoneIds []*uint64 `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// Subnet ID.
	SubnetIds []*uint64 `json:"SubnetIds,omitempty" name:"SubnetIds"`

	// Lock flag.
	CdbErrors []*int64 `json:"CdbErrors,omitempty" name:"CdbErrors"`

	// Sort by field of the returned result set. Currently, supported values include "InstanceId", "InstanceName", "CreateTime", and "DeadlineTime".
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// Sorting method of the returned result set. Currently, "ASC" or "DESC" is supported.
	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`

	// Whether security group ID is used as a filter
	WithSecurityGroup *int64 `json:"WithSecurityGroup,omitempty" name:"WithSecurityGroup"`

	// Whether dedicated cluster details are included. Value range: 0 (not included), 1 (included)
	WithExCluster *int64 `json:"WithExCluster,omitempty" name:"WithExCluster"`

	// Exclusive cluster ID.
	ExClusterId *string `json:"ExClusterId,omitempty" name:"ExClusterId"`

	// Instance ID.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// Initialization flag. Value range: 0 (not initialized), 1 (initialized).
	InitFlag *int64 `json:"InitFlag,omitempty" name:"InitFlag"`

	// Whether instances corresponding to the disaster recovery relationship are included. Valid values: 0 (not included), 1 (included). Default value: 1. If a primary instance is pulled, the data of the disaster recovery relationship will be in the `DrInfo` field. If a disaster recovery instance is pulled, the data of the disaster recovery relationship will be in the `MasterInfo` field. The disaster recovery relationship contains only partial basic data. To get the detailed data, you need to call an API to pull it.
	WithDr *int64 `json:"WithDr,omitempty" name:"WithDr"`

	// Whether read-only instances are included. Valid values: 0 (not included), 1 (included). Default value: 1.
	WithRo *int64 `json:"WithRo,omitempty" name:"WithRo"`

	// Whether primary instances are included. Valid values: 0 (not included), 1 (included). Default value: 1.
	WithMaster *int64 `json:"WithMaster,omitempty" name:"WithMaster"`

	// Placement group ID list.
	DeployGroupIds []*string `json:"DeployGroupIds,omitempty" name:"DeployGroupIds"`

	// Whether to use the tag key as a filter condition
	TagKeysForSearch []*string `json:"TagKeysForSearch,omitempty" name:"TagKeysForSearch"`

	// Financial cage IDs.
	CageIds []*string `json:"CageIds,omitempty" name:"CageIds"`
}

func (r *DescribeDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "InstanceTypes")
	delete(f, "Vips")
	delete(f, "Status")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "SecurityGroupId")
	delete(f, "PayTypes")
	delete(f, "InstanceNames")
	delete(f, "TaskStatus")
	delete(f, "EngineVersions")
	delete(f, "VpcIds")
	delete(f, "ZoneIds")
	delete(f, "SubnetIds")
	delete(f, "CdbErrors")
	delete(f, "OrderBy")
	delete(f, "OrderDirection")
	delete(f, "WithSecurityGroup")
	delete(f, "WithExCluster")
	delete(f, "ExClusterId")
	delete(f, "InstanceIds")
	delete(f, "InitFlag")
	delete(f, "WithDr")
	delete(f, "WithRo")
	delete(f, "WithMaster")
	delete(f, "DeployGroupIds")
	delete(f, "TagKeysForSearch")
	delete(f, "CageIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of eligible instances.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Instance details.
		Items []*InstanceInfo `json:"Items,omitempty" name:"Items"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBSecurityGroupsRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cdb-c1nl9rpv or cdbro-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// This parameter takes effect only when the ID of read-only replica is passed in. If this parameter is set to `False` or left empty, the security groups bound with the RO group of the read-only replica will be queried. If this parameter is set to `True`, the security groups bound with the read-only replica itself will be queried.
	ForReadonlyInstance *bool `json:"ForReadonlyInstance,omitempty" name:"ForReadonlyInstance"`
}

func (r *DescribeDBSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBSecurityGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ForReadonlyInstance")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBSecurityGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Security group details.
		Groups []*SecurityGroup `json:"Groups,omitempty" name:"Groups"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBSecurityGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBSwitchRecordsRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cdb-c1nl9rpv or cdbro-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Pagination offset.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of entries per page. Value range: 1-2,000. Default value: 50.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeDBSwitchRecordsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBSwitchRecordsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBSwitchRecordsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBSwitchRecordsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of instance switches.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Details of instance switches.
		Items []*DBSwitchInfo `json:"Items,omitempty" name:"Items"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBSwitchRecordsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBSwitchRecordsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBZoneConfigRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeDBZoneConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBZoneConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBZoneConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBZoneConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of configurations in purchasable regions
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Details of configurations in purchasable regions
		Items []*RegionSellConf `json:"Items,omitempty" name:"Items"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBZoneConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBZoneConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDataBackupOverviewRequest struct {
	*tchttp.BaseRequest

	// TencentDB product type to be queried. Currently, only `mysql` is supported.
	Product *string `json:"Product,omitempty" name:"Product"`
}

func (r *DescribeDataBackupOverviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataBackupOverviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataBackupOverviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDataBackupOverviewResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total capacity of data backups in bytes in the current region (including automatic backups and manual backups).
		DataBackupVolume *int64 `json:"DataBackupVolume,omitempty" name:"DataBackupVolume"`

		// Total number of data backups in the current region.
		DataBackupCount *int64 `json:"DataBackupCount,omitempty" name:"DataBackupCount"`

		// Total capacity of automatic backups in the current region.
		AutoBackupVolume *int64 `json:"AutoBackupVolume,omitempty" name:"AutoBackupVolume"`

		// Total number of automatic backups in the current region.
		AutoBackupCount *int64 `json:"AutoBackupCount,omitempty" name:"AutoBackupCount"`

		// Total capacity of manual backups in the current region.
		ManualBackupVolume *int64 `json:"ManualBackupVolume,omitempty" name:"ManualBackupVolume"`

		// Total number of manual backups in the current region.
		ManualBackupCount *int64 `json:"ManualBackupCount,omitempty" name:"ManualBackupCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDataBackupOverviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataBackupOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDatabasesRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Offset. Minimum value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of results to be returned for a single request. Value range: 1-100. Maximum value: 20.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Regular expression for matching database names.
	DatabaseRegexp *string `json:"DatabaseRegexp,omitempty" name:"DatabaseRegexp"`
}

func (r *DescribeDatabasesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatabasesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "DatabaseRegexp")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDatabasesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDatabasesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of eligible instances.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Information of an instance.
		Items []*string `json:"Items,omitempty" name:"Items"`

		// Database name and character set
		DatabaseList []*DatabasesWithCharacterLists `json:"DatabaseList,omitempty" name:"DatabaseList"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDatabasesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatabasesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDefaultParamsRequest struct {
	*tchttp.BaseRequest

	// MySQL version. Currently, the supported versions are ["5.1", "5.5", "5.6", "5.7"].
	EngineVersion *string `json:"EngineVersion,omitempty" name:"EngineVersion"`
}

func (r *DescribeDefaultParamsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDefaultParamsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EngineVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDefaultParamsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDefaultParamsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of parameters
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Parameter details.
		Items []*ParameterDetail `json:"Items,omitempty" name:"Items"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDefaultParamsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDefaultParamsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDeployGroupListRequest struct {
	*tchttp.BaseRequest

	// ID of a placement group.
	DeployGroupId *string `json:"DeployGroupId,omitempty" name:"DeployGroupId"`

	// Name of a placement group.
	DeployGroupName *string `json:"DeployGroupName,omitempty" name:"DeployGroupName"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeDeployGroupListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeployGroupListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeployGroupId")
	delete(f, "DeployGroupName")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDeployGroupListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDeployGroupListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of eligible entries.
		Total *int64 `json:"Total,omitempty" name:"Total"`

		// List of returned results.
	// Note: This field may return null, indicating that no valid values can be obtained.
		Items []*DeployGroupInfo `json:"Items,omitempty" name:"Items"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDeployGroupListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeployGroupListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDeviceMonitorInfoRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// This parameter is used to return the monitoring data of Count 5-minute time periods on the day. Value range: 1-288. If this parameter is not passed in, all monitoring data in a 5-minute granularity on the day will be returned by default.
	Count *uint64 `json:"Count,omitempty" name:"Count"`
}

func (r *DescribeDeviceMonitorInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceMonitorInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Count")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDeviceMonitorInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDeviceMonitorInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// CPU monitoring data of the instance
		Cpu *DeviceCpuInfo `json:"Cpu,omitempty" name:"Cpu"`

		// Memory monitoring data of the instance
		Mem *DeviceMemInfo `json:"Mem,omitempty" name:"Mem"`

		// Network monitoring data of the instance
		Net *DeviceNetInfo `json:"Net,omitempty" name:"Net"`

		// Disk monitoring data of the instance
		Disk *DeviceDiskInfo `json:"Disk,omitempty" name:"Disk"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDeviceMonitorInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceMonitorInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeErrorLogDataRequest struct {
	*tchttp.BaseRequest

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Start timestamp.
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// End timestamp.
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// List of keywords to match. Up to 15 keywords are supported.
	KeyWords []*string `json:"KeyWords,omitempty" name:"KeyWords"`

	// The number of results per page in paginated queries. Default value: 100. Maximum value: 400.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// This parameter is valid only for source or disaster recovery instances. Valid value: `slave`, which indicates pulling logs from the replica.
	InstType *string `json:"InstType,omitempty" name:"InstType"`
}

func (r *DescribeErrorLogDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeErrorLogDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "KeyWords")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "InstType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeErrorLogDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeErrorLogDataResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of eligible entries.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Returned result.
	// Note: this field may return null, indicating that no valid values can be obtained.
		Items []*ErrlogItem `json:"Items,omitempty" name:"Items"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeErrorLogDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeErrorLogDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceParamRecordsRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page. You can use the [instance list querying API](https://intl.cloud.tencent.com/document/api/236/15872?from_cn_redirect=1) to query the ID, whose value is the `InstanceId` value in output parameters.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Pagination offset.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of entries per page.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeInstanceParamRecordsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceParamRecordsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceParamRecordsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceParamRecordsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of eligible records.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Parameter modification records.
		Items []*ParamRecord `json:"Items,omitempty" name:"Items"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceParamRecordsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceParamRecordsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceParamsRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page. You can use the [instance list querying API](https://intl.cloud.tencent.com/document/api/236/15872?from_cn_redirect=1) to query the ID, whose value is the `InstanceId` value in output parameters.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeInstanceParamsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceParamsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceParamsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceParamsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of instance parameters.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Parameter details.
		Items []*ParameterDetail `json:"Items,omitempty" name:"Items"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceParamsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceParamsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeParamTemplateInfoRequest struct {
	*tchttp.BaseRequest

	// Parameter template ID.
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
}

func (r *DescribeParamTemplateInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeParamTemplateInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeParamTemplateInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeParamTemplateInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Parameter template ID.
		TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

		// Parameter template name.
		Name *string `json:"Name,omitempty" name:"Name"`

		// Database engine version specified in the parameter template
		EngineVersion *string `json:"EngineVersion,omitempty" name:"EngineVersion"`

		// Number of parameters in the parameter template
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Parameter details
		Items []*ParameterDetail `json:"Items,omitempty" name:"Items"`

		// Parameter template description
		Description *string `json:"Description,omitempty" name:"Description"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeParamTemplateInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeParamTemplateInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeParamTemplatesRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeParamTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeParamTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeParamTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeParamTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of parameter templates of the user.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Parameter template details.
		Items []*ParamTemplateInfo `json:"Items,omitempty" name:"Items"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeParamTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeParamTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProjectSecurityGroupsRequest struct {
	*tchttp.BaseRequest

	// Project ID.
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *DescribeProjectSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProjectSecurityGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProjectSecurityGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProjectSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Security group details.
		Groups []*SecurityGroup `json:"Groups,omitempty" name:"Groups"`

		// Number of security group rules
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProjectSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProjectSecurityGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRoGroupsRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of `cdb-c1nl9rpv` or `cdb-c1nl9rpv`. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeRoGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRoGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRoGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRoGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// RO group information array. An instance can be associated with multiple RO groups.
		RoGroups []*RoGroup `json:"RoGroups,omitempty" name:"RoGroups"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRoGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRoGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRoMinScaleRequest struct {
	*tchttp.BaseRequest

	// Read-only instance ID in the format of "cdbro-c1nl9rpv". Its value is the same as the instance ID in the TencentDB Console. This parameter and the `MasterInstanceId` parameter cannot both be empty.
	RoInstanceId *string `json:"RoInstanceId,omitempty" name:"RoInstanceId"`

	// Primary instance ID in the format of "cdbro-c1nl9rpv". Its value is the same as the instance ID in the TencentDB Console. This parameter and the `RoInstanceId` parameter cannot both be empty. Note: when the parameters are passed in with `RoInstanceId`, the return value refers to the minimum specification to which a read-only instance can be upgraded; when the parameters are passed in with `MasterInstanceId` but without `RoInstanceId`, the return value refers to the minimum purchasable specification for a read-only instance.
	MasterInstanceId *string `json:"MasterInstanceId,omitempty" name:"MasterInstanceId"`
}

func (r *DescribeRoMinScaleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRoMinScaleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RoInstanceId")
	delete(f, "MasterInstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRoMinScaleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRoMinScaleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Memory size in MB.
		Memory *int64 `json:"Memory,omitempty" name:"Memory"`

		// Disk size in GB.
		Volume *int64 `json:"Volume,omitempty" name:"Volume"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRoMinScaleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRoMinScaleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRollbackRangeTimeRequest struct {
	*tchttp.BaseRequest

	// Instance ID list. An instance ID is in the format of cdb-c1nl9rpv, which is the same as the instance ID displayed on the TencentDB Console page.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *DescribeRollbackRangeTimeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRollbackRangeTimeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRollbackRangeTimeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRollbackRangeTimeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of eligible instances.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Returned parameter information.
		Items []*InstanceRollbackRangeTime `json:"Items,omitempty" name:"Items"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRollbackRangeTimeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRollbackRangeTimeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRollbackTaskDetailRequest struct {
	*tchttp.BaseRequest

	// Instance ID, which is the same as the instance ID displayed in the TencentDB Console. You can use the [DescribeDBInstances API](https://intl.cloud.tencent.com/document/api/236/15872?from_cn_redirect=1) to query the ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Async task ID.
	AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

	// Pagination parameter, i.e., the number of entries to be returned for a single request. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Pagination offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeRollbackTaskDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRollbackTaskDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "AsyncRequestId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRollbackTaskDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRollbackTaskDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of eligible entries.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Rollback task details.
	// Note: this field may return null, indicating that no valid values can be obtained.
		Items []*RollbackTask `json:"Items,omitempty" name:"Items"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRollbackTaskDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRollbackTaskDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSlowLogDataRequest struct {
	*tchttp.BaseRequest

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Start timestamp.
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// End timestamp.
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// Client `Host` list.
	UserHosts []*string `json:"UserHosts,omitempty" name:"UserHosts"`

	// Client username list.
	UserNames []*string `json:"UserNames,omitempty" name:"UserNames"`

	// Accessed database list.
	DataBases []*string `json:"DataBases,omitempty" name:"DataBases"`

	// Sort by field. Valid values: Timestamp, QueryTime, LockTime, RowsExamined, RowsSent.
	SortBy *string `json:"SortBy,omitempty" name:"SortBy"`

	// Sorting order. Valid values: ASC (ascending), DESC (descending).
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// The number of results per page in paginated queries. Default value: 100. Maximum value: 400.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// This parameter is valid only for source or disaster recovery instances. Valid value: `slave`, which indicates pulling logs from the replica.
	InstType *string `json:"InstType,omitempty" name:"InstType"`
}

func (r *DescribeSlowLogDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlowLogDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "UserHosts")
	delete(f, "UserNames")
	delete(f, "DataBases")
	delete(f, "SortBy")
	delete(f, "OrderBy")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "InstType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSlowLogDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSlowLogDataResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of eligible entries.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Queried results.
	// Note: this field may return null, indicating that no valid values can be obtained.
		Items []*SlowLogItem `json:"Items,omitempty" name:"Items"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSlowLogDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlowLogDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSlowLogsRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Offset. Minimum value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of entries per page. Value range: 1-100. Default value: 20.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
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
	delete(f, "InstanceId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSlowLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSlowLogsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of eligible slow logs.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Details of eligible slow logs.
		Items []*SlowLogInfo `json:"Items,omitempty" name:"Items"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeSupportedPrivilegesRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeSupportedPrivilegesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSupportedPrivilegesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSupportedPrivilegesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSupportedPrivilegesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Global permissions supported by the instance
		GlobalSupportedPrivileges []*string `json:"GlobalSupportedPrivileges,omitempty" name:"GlobalSupportedPrivileges"`

		// Database permissions supported by the instance.
		DatabaseSupportedPrivileges []*string `json:"DatabaseSupportedPrivileges,omitempty" name:"DatabaseSupportedPrivileges"`

		// Table permissions supported by the instance.
		TableSupportedPrivileges []*string `json:"TableSupportedPrivileges,omitempty" name:"TableSupportedPrivileges"`

		// Column permissions supported by the instance.
		ColumnSupportedPrivileges []*string `json:"ColumnSupportedPrivileges,omitempty" name:"ColumnSupportedPrivileges"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSupportedPrivilegesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSupportedPrivilegesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTablesRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Database name.
	Database *string `json:"Database,omitempty" name:"Database"`

	// Record offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of results to be returned for a single request. Default value: 20. Maximum value: 2,000.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Regular expression for matching table names, which complies with the rules at MySQL's official website
	TableRegexp *string `json:"TableRegexp,omitempty" name:"TableRegexp"`
}

func (r *DescribeTablesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTablesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Database")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "TableRegexp")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTablesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTablesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of eligible tables.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Information of a table.
		Items []*string `json:"Items,omitempty" name:"Items"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTablesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTablesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTagsOfInstanceIdsRequest struct {
	*tchttp.BaseRequest

	// List of instances.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// Pagination offset.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of entries per page.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeTagsOfInstanceIdsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTagsOfInstanceIdsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTagsOfInstanceIdsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTagsOfInstanceIdsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Pagination offset.
		Offset *int64 `json:"Offset,omitempty" name:"Offset"`

		// Number of entries per page.
		Limit *int64 `json:"Limit,omitempty" name:"Limit"`

		// Instance tag information.
		Rows []*TagsInfoOfInstance `json:"Rows,omitempty" name:"Rows"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTagsOfInstanceIdsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTagsOfInstanceIdsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTasksRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page. You can use the [instance list querying API](https://intl.cloud.tencent.com/document/api/236/15872?from_cn_redirect=1) to query the ID, whose value is the `InstanceId` value in output parameters.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// ID of an async task request, i.e., `AsyncRequestId` returned by relevant TencentDB operations.
	AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

	// Task type. If no value is passed in, all task types will be queried. Valid values:
	// 1 - rolling back a database;
	// 2 - performing an SQL operation;
	// 3 - importing data;
	// 5 - setting a parameter;
	// 6 - initializing a TencentDB instance;
	// 7 - restarting a TencentDB instance;
	// 8 - enabling GTID of a TencentDB instance;
	// 9 - upgrading a read-only instance;
	// 10 - rolling back databases in batches;
	// 11 - upgrading a primary instance;
	// 12 - deleting a TencentDB table;
	// 13 - promoting a disaster recovery instance.
	TaskTypes []*int64 `json:"TaskTypes,omitempty" name:"TaskTypes"`

	// Task status. If no value is passed in, all task statuses will be queried. Valid values:
	// -1 - undefined;
	// 0 - initializing;
	// 1 - running;
	// 2 - succeeded;
	// 3 - failed;
	// 4 - terminated;
	// 5 - deleted;
	// 6 - paused.
	TaskStatus []*int64 `json:"TaskStatus,omitempty" name:"TaskStatus"`

	// Start time of the first task in the format of yyyy-MM-dd HH:mm:ss, such as 2017-12-31 10:40:01. It is used for queries by time range.
	StartTimeBegin *string `json:"StartTimeBegin,omitempty" name:"StartTimeBegin"`

	// End time of the last task in the format of yyyy-MM-dd HH:mm:ss, such as 2017-12-31 10:40:01. It is used for queries by time range.
	StartTimeEnd *string `json:"StartTimeEnd,omitempty" name:"StartTimeEnd"`

	// Record offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of results to be returned for a single request. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "AsyncRequestId")
	delete(f, "TaskTypes")
	delete(f, "TaskStatus")
	delete(f, "StartTimeBegin")
	delete(f, "StartTimeEnd")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTasksResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of eligible instances.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Information of an instance task.
		Items []*TaskDetail `json:"Items,omitempty" name:"Items"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTimeWindowRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cdb-c1nl9rpv or cdbro-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeTimeWindowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTimeWindowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTimeWindowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTimeWindowResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// List of maintenance time windows on Monday.
		Monday []*string `json:"Monday,omitempty" name:"Monday"`

		// List of maintenance time windows on Tuesday.
		Tuesday []*string `json:"Tuesday,omitempty" name:"Tuesday"`

		// List of maintenance time windows on Wednesday.
		Wednesday []*string `json:"Wednesday,omitempty" name:"Wednesday"`

		// List of maintenance time windows on Thursday.
		Thursday []*string `json:"Thursday,omitempty" name:"Thursday"`

		// List of maintenance time windows on Friday.
		Friday []*string `json:"Friday,omitempty" name:"Friday"`

		// List of maintenance time windows on Saturday.
		Saturday []*string `json:"Saturday,omitempty" name:"Saturday"`

		// List of maintenance time windows on Sunday.
		Sunday []*string `json:"Sunday,omitempty" name:"Sunday"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTimeWindowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTimeWindowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUploadedFilesRequest struct {
	*tchttp.BaseRequest

	// File path. `OwnerUin` information of the root account should be entered in this field.
	Path *string `json:"Path,omitempty" name:"Path"`

	// Record offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of results to be returned for a single request. Default value: 20.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeUploadedFilesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUploadedFilesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Path")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUploadedFilesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUploadedFilesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of eligible SQL files.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// List of returned SQL files.
		Items []*SqlFileInfo `json:"Items,omitempty" name:"Items"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUploadedFilesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUploadedFilesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeviceCpuInfo struct {

	// Average instance CPU utilization
	Rate []*DeviceCpuRateInfo `json:"Rate,omitempty" name:"Rate"`

	// CPU monitoring data of the instance
	Load []*int64 `json:"Load,omitempty" name:"Load"`
}

type DeviceCpuRateInfo struct {

	// CPU core number
	CpuCore *int64 `json:"CpuCore,omitempty" name:"CpuCore"`

	// CPU utilization
	Rate []*int64 `json:"Rate,omitempty" name:"Rate"`
}

type DeviceDiskInfo struct {

	// Time percentage of IO operations per second
	IoRatioPerSec []*int64 `json:"IoRatioPerSec,omitempty" name:"IoRatioPerSec"`

	// Average wait time of device I/O operations * 100 in milliseconds. For example, if the value is 201, the average wait time of I/O operations is 201/100 = 2.1 milliseconds.
	IoWaitTime []*int64 `json:"IoWaitTime,omitempty" name:"IoWaitTime"`

	// Average number of read operations completed by the disk per second * 100. For example, if the value is 2,002, the average number of read operations completed by the disk per second is 2,002/100=20.2.
	Read []*int64 `json:"Read,omitempty" name:"Read"`

	// Average number of write operations completed by the disk per second * 100. For example, if the value is 30,001, the average number of write operations completed by the disk per second is 30,001/100=300.01.
	Write []*int64 `json:"Write,omitempty" name:"Write"`

	// Disk capacity. Each value is comprised of two data, with the first data representing the used capacity and the second one representing the total disk capacity.
	CapacityRatio []*int64 `json:"CapacityRatio,omitempty" name:"CapacityRatio"`
}

type DeviceMemInfo struct {

	// Total memory size in KB, which is the value of `total` in the `Mem:` in the `free` command
	Total []*int64 `json:"Total,omitempty" name:"Total"`

	// Used memory size in KB, which is the value of `used` in the `Mem:` row in the `free` command
	Used []*int64 `json:"Used,omitempty" name:"Used"`
}

type DeviceNetInfo struct {

	// Number of TCP connections
	Conn []*int64 `json:"Conn,omitempty" name:"Conn"`

	// ENI inbound packets per second
	PackageIn []*int64 `json:"PackageIn,omitempty" name:"PackageIn"`

	// ENI outbound packets per second
	PackageOut []*int64 `json:"PackageOut,omitempty" name:"PackageOut"`

	// Inbound traffic in Kbps
	FlowIn []*int64 `json:"FlowIn,omitempty" name:"FlowIn"`

	// Outbound traffic in Kbps
	FlowOut []*int64 `json:"FlowOut,omitempty" name:"FlowOut"`
}

type DisassociateSecurityGroupsRequest struct {
	*tchttp.BaseRequest

	// Security group ID.
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`

	// List of instance IDs, which is an array of one or more instance IDs.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// This parameter takes effect only when the IDs of read-only replicas are passed in. If this parameter is set to `False` or left empty, the security group will be unbound from the RO groups of these read-only replicas. If this parameter is set to `True`, the security group will be unbound from the read-only replicas themselves.
	ForReadonlyInstance *bool `json:"ForReadonlyInstance,omitempty" name:"ForReadonlyInstance"`
}

func (r *DisassociateSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisassociateSecurityGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecurityGroupId")
	delete(f, "InstanceIds")
	delete(f, "ForReadonlyInstance")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisassociateSecurityGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DisassociateSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DisassociateSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisassociateSecurityGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DrInfo struct {

	// Disaster recovery instance status
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// AZ information
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Region information
	Region *string `json:"Region,omitempty" name:"Region"`

	// Instance sync status. Possible returned values include:
	// 0 - disaster recovery not synced;
	// 1 - disaster recovery syncing;
	// 2 - disaster recovery synced successfully;
	// 3 - disaster recovery sync failed;
	// 4 - repairing disaster recovery sync;
	SyncStatus *int64 `json:"SyncStatus,omitempty" name:"SyncStatus"`

	// Instance name
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// Instance type
	InstanceType *int64 `json:"InstanceType,omitempty" name:"InstanceType"`
}

type ErrlogItem struct {

	// Error occurrence time.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Timestamp *uint64 `json:"Timestamp,omitempty" name:"Timestamp"`

	// Error details
	// Note: this field may return null, indicating that no valid values can be obtained.
	Content *string `json:"Content,omitempty" name:"Content"`
}

type ImportRecord struct {

	// Status value
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// Status value
	Code *int64 `json:"Code,omitempty" name:"Code"`

	// Execution duration
	CostTime *int64 `json:"CostTime,omitempty" name:"CostTime"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Backend task ID
	WorkId *string `json:"WorkId,omitempty" name:"WorkId"`

	// Name of the file to be imported
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// Execution progress
	Process *int64 `json:"Process,omitempty" name:"Process"`

	// Task creation time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// File size
	FileSize *string `json:"FileSize,omitempty" name:"FileSize"`

	// Task execution information
	Message *string `json:"Message,omitempty" name:"Message"`

	// Task ID
	JobId *int64 `json:"JobId,omitempty" name:"JobId"`

	// Name of the table to be imported
	DbName *string `json:"DbName,omitempty" name:"DbName"`

	// Async task request ID
	AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`
}

type Inbound struct {

	// Policy, which can be ACCEPT or DROP
	Action *string `json:"Action,omitempty" name:"Action"`

	// Source IP or IP range, such as 192.168.0.0/16
	CidrIp *string `json:"CidrIp,omitempty" name:"CidrIp"`

	// Port
	PortRange *string `json:"PortRange,omitempty" name:"PortRange"`

	// Network protocol. UDP and TCP are supported.
	IpProtocol *string `json:"IpProtocol,omitempty" name:"IpProtocol"`

	// The direction of the rule, which is INPUT for inbound rules
	Dir *string `json:"Dir,omitempty" name:"Dir"`

	// Rule description
	Desc *string `json:"Desc,omitempty" name:"Desc"`
}

type InitDBInstancesRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page. You can use the [instance list querying API](https://intl.cloud.tencent.com/document/api/236/15872?from_cn_redirect=1) to query the ID, whose value is the `InstanceId` value in output parameters.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// New password of the instance. Rule: It can only contain 8-64 characters and must contain at least two of the following types of characters: letters, digits, and special characters (!@#$%^*()).
	NewPassword *string `json:"NewPassword,omitempty" name:"NewPassword"`

	// List of instance parameters. Currently, "character_set_server" and "lower_case_table_names" are supported, whose value ranges are ["utf8","latin1","gbk","utf8mb4"] and ["0","1"], respectively.
	Parameters []*ParamInfo `json:"Parameters,omitempty" name:"Parameters"`

	// Instance port. Value range: [1024, 65535].
	Vport *int64 `json:"Vport,omitempty" name:"Vport"`
}

func (r *InitDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InitDBInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "NewPassword")
	delete(f, "Parameters")
	delete(f, "Vport")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InitDBInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type InitDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Array of async task request IDs, which can be used to query the execution results of async tasks.
		AsyncRequestIds []*string `json:"AsyncRequestIds,omitempty" name:"AsyncRequestIds"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InitDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InitDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstanceInfo struct {

	// Public network access status. Value range: 0 (not enabled), 1 (enabled), 2 (disabled)
	WanStatus *int64 `json:"WanStatus,omitempty" name:"WanStatus"`

	// AZ information
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// Initialization flag. Value range: 0 (not initialized), 1 (initialized)
	InitFlag *int64 `json:"InitFlag,omitempty" name:"InitFlag"`

	// VIP information of a read-only instance. This field is exclusive to read-only instances where read-only access is enabled separately
	// Note: This field may return null, indicating that no valid values can be obtained.
	RoVipInfo *RoVipInfo `json:"RoVipInfo,omitempty" name:"RoVipInfo"`

	// Memory capacity in MB
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// Instance status. Value range: 0 (creating), 1 (running), 4 (isolating), 5 (isolated)
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// VPC ID, such as 51102
	VpcId *int64 `json:"VpcId,omitempty" name:"VpcId"`

	// Information of a secondary server
	// Note: This field may return null, indicating that no valid values can be obtained.
	SlaveInfo *SlaveInfo `json:"SlaveInfo,omitempty" name:"SlaveInfo"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Disk capacity in GB
	Volume *int64 `json:"Volume,omitempty" name:"Volume"`

	// Auto-renewal flag. Value range: 0 (auto-renewal not enabled), 1 (auto-renewal enabled), 2 (auto-renewal disabled)
	AutoRenew *int64 `json:"AutoRenew,omitempty" name:"AutoRenew"`

	// Data replication mode. Valid values: 0 (async), 1 (semi-sync), 2 (strong sync)
	ProtectMode *int64 `json:"ProtectMode,omitempty" name:"ProtectMode"`

	// Details of a read-only group
	// Note: This field may return null, indicating that no valid values can be obtained.
	RoGroups []*RoGroup `json:"RoGroups,omitempty" name:"RoGroups"`

	// Subnet ID, such as 2333
	SubnetId *int64 `json:"SubnetId,omitempty" name:"SubnetId"`

	// Instance type. Value range: 1 (primary), 2 (disaster recovery), 3 (read-only)
	InstanceType *int64 `json:"InstanceType,omitempty" name:"InstanceType"`

	// Project ID
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Region information
	Region *string `json:"Region,omitempty" name:"Region"`

	// Instance expiration time
	DeadlineTime *string `json:"DeadlineTime,omitempty" name:"DeadlineTime"`

	// AZ deployment mode. Valid values: 0 (single-AZ), 1 (multi-AZ)
	DeployMode *int64 `json:"DeployMode,omitempty" name:"DeployMode"`

	// Instance task status. 0 - no task; 1 - upgrading; 2 - importing data; 3 - activating secondary; 4 - enabling public network access; 5 - batch operation in progress; 6 - rolling back; 7 - disabling public network access; 8 - changing password; 9 - renaming instance; 10 - restarting; 12 - migrating self-built instance; 13 - dropping table; 14 - creating and syncing disaster recovery instance; 15 - pending upgrade and switch; 16 - upgrade and switch in progress; 17 - upgrade and switch completed
	TaskStatus *int64 `json:"TaskStatus,omitempty" name:"TaskStatus"`

	// Details of a primary instance
	// Note: This field may return null, indicating that no valid values can be obtained.
	MasterInfo *MasterInfo `json:"MasterInfo,omitempty" name:"MasterInfo"`

	// Instance type
	DeviceType *string `json:"DeviceType,omitempty" name:"DeviceType"`

	// Kernel version
	EngineVersion *string `json:"EngineVersion,omitempty" name:"EngineVersion"`

	// Instance name
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// Details of a disaster recovery instance
	// Note: This field may return null, indicating that no valid values can be obtained.
	DrInfo []*DrInfo `json:"DrInfo,omitempty" name:"DrInfo"`

	// Public domain name
	WanDomain *string `json:"WanDomain,omitempty" name:"WanDomain"`

	// Public network port number
	WanPort *int64 `json:"WanPort,omitempty" name:"WanPort"`

	// Billing type
	PayType *int64 `json:"PayType,omitempty" name:"PayType"`

	// Instance creation time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Instance IP
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// Port number
	Vport *int64 `json:"Vport,omitempty" name:"Vport"`

	// Lock flag
	CdbError *int64 `json:"CdbError,omitempty" name:"CdbError"`

	// VPC descriptor, such as "vpc-5v8wn9mg"
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// Subnet descriptor, such as "subnet-1typ0s7d"
	UniqSubnetId *string `json:"UniqSubnetId,omitempty" name:"UniqSubnetId"`

	// Physical ID
	PhysicalId *string `json:"PhysicalId,omitempty" name:"PhysicalId"`

	// Number of cores
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

	// Queries per second
	Qps *int64 `json:"Qps,omitempty" name:"Qps"`

	// AZ name
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// Physical machine model
	// Note: This field may return null, indicating that no valid values can be obtained.
	DeviceClass *string `json:"DeviceClass,omitempty" name:"DeviceClass"`

	// Placement group ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	DeployGroupId *string `json:"DeployGroupId,omitempty" name:"DeployGroupId"`

	// AZ ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// Number of nodes
	InstanceNodes *int64 `json:"InstanceNodes,omitempty" name:"InstanceNodes"`

	// List of tags
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	TagList []*TagInfoItem `json:"TagList,omitempty" name:"TagList"`
}

type InstanceRebootTime struct {

	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Estimated restart time
	TimeInSeconds *int64 `json:"TimeInSeconds,omitempty" name:"TimeInSeconds"`
}

type InstanceRollbackRangeTime struct {

	// Queries database error code
	Code *int64 `json:"Code,omitempty" name:"Code"`

	// Queries database error message
	Message *string `json:"Message,omitempty" name:"Message"`

	// List of instance IDs. An instance ID is in the format of cdb-c1nl9rpv, which is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Time range available for rollback
	Times []*RollbackTimeRange `json:"Times,omitempty" name:"Times"`
}

type IsolateDBInstanceRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page. You can use the [instance list querying API](https://intl.cloud.tencent.com/document/api/236/15872?from_cn_redirect=1) to query the ID, whose value is the `InstanceId` value in output parameters.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *IsolateDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IsolateDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "IsolateDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type IsolateDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Async task request ID, which can be used to query the execution result of an async task. (This returned field has been disused. You can query the isolation status of an instance through the `DescribeDBInstances` API.)
	// Note: this field may return null, indicating that no valid values can be obtained.
		AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *IsolateDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IsolateDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MasterInfo struct {

	// Region information
	Region *string `json:"Region,omitempty" name:"Region"`

	// Region ID
	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`

	// AZ ID
	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// AZ information
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Long instance ID
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// Instance status
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// Instance name
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// Instance type
	InstanceType *int64 `json:"InstanceType,omitempty" name:"InstanceType"`

	// Task status
	TaskStatus *int64 `json:"TaskStatus,omitempty" name:"TaskStatus"`

	// Memory capacity
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// Disk capacity
	Volume *int64 `json:"Volume,omitempty" name:"Volume"`

	// Instance model
	DeviceType *string `json:"DeviceType,omitempty" name:"DeviceType"`

	// Queries per second
	Qps *int64 `json:"Qps,omitempty" name:"Qps"`

	// VPC ID
	VpcId *int64 `json:"VpcId,omitempty" name:"VpcId"`

	// Subnet ID
	SubnetId *int64 `json:"SubnetId,omitempty" name:"SubnetId"`

	// Dedicated cluster ID
	ExClusterId *string `json:"ExClusterId,omitempty" name:"ExClusterId"`

	// Dedicated cluster name
	ExClusterName *string `json:"ExClusterName,omitempty" name:"ExClusterName"`
}

type ModifyAccountDescriptionRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// TencentDB account
	Accounts []*Account `json:"Accounts,omitempty" name:"Accounts"`

	// Database account remarks
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *ModifyAccountDescriptionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccountDescriptionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Accounts")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAccountDescriptionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAccountDescriptionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Async task request ID, which can be used to query the execution result of an async task.
		AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAccountDescriptionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccountDescriptionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAccountMaxUserConnectionsRequest struct {
	*tchttp.BaseRequest

	// List of TencentDB accounts
	Accounts []*Account `json:"Accounts,omitempty" name:"Accounts"`

	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed in the TencentDB console.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Maximum connections of the account. Maximum value: `10240`.
	MaxUserConnections *int64 `json:"MaxUserConnections,omitempty" name:"MaxUserConnections"`
}

func (r *ModifyAccountMaxUserConnectionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccountMaxUserConnectionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Accounts")
	delete(f, "InstanceId")
	delete(f, "MaxUserConnections")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAccountMaxUserConnectionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAccountMaxUserConnectionsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Async task request ID, which can be used to query the execution result of an async task
		AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAccountMaxUserConnectionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccountMaxUserConnectionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAccountPasswordRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// New password of the database account. It can only contain 8-64 characters and must contain at least two of the following types of characters: letters, digits, and special characters (_+-&=!@#$%^*()).
	NewPassword *string `json:"NewPassword,omitempty" name:"NewPassword"`

	// TencentDB account
	Accounts []*Account `json:"Accounts,omitempty" name:"Accounts"`
}

func (r *ModifyAccountPasswordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccountPasswordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "NewPassword")
	delete(f, "Accounts")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAccountPasswordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAccountPasswordResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Async task request ID, which can be used to query the execution result of an async task.
		AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAccountPasswordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccountPasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAccountPrivilegesRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Database account, including username and domain name.
	Accounts []*Account `json:"Accounts,omitempty" name:"Accounts"`

	// Global permission. Valid values: "SELECT", "INSERT", "UPDATE", "DELETE", "CREATE", "PROCESS", "DROP", "REFERENCES", "INDEX", "ALTER", "SHOW DATABASES", "CREATE TEMPORARY TABLES", "LOCK TABLES", "EXECUTE", "CREATE VIEW", "SHOW VIEW", "CREATE ROUTINE", "ALTER ROUTINE", "EVENT", "TRIGGER".
	// Note: if this parameter is not passed in, it means to clear the permission.
	GlobalPrivileges []*string `json:"GlobalPrivileges,omitempty" name:"GlobalPrivileges"`

	// Database permission. Valid values: "SELECT", "INSERT", "UPDATE", "DELETE", "CREATE", 	"DROP", "REFERENCES", "INDEX", "ALTER", "CREATE TEMPORARY TABLES", "LOCK TABLES", "EXECUTE", "CREATE VIEW", "SHOW VIEW", "CREATE ROUTINE", "ALTER ROUTINE", "EVENT", "TRIGGER".
	// Note: if this parameter is not passed in, it means to clear the permission.
	DatabasePrivileges []*DatabasePrivilege `json:"DatabasePrivileges,omitempty" name:"DatabasePrivileges"`

	// Table permission in the database. Valid values: "SELECT", "INSERT", "UPDATE", "DELETE", "CREATE", 	"DROP", "REFERENCES", "INDEX", "ALTER", "CREATE VIEW", "SHOW VIEW", "TRIGGER".
	// Note: if this parameter is not passed in, it means to clear the permission.
	TablePrivileges []*TablePrivilege `json:"TablePrivileges,omitempty" name:"TablePrivileges"`

	// Column permission in table. Valid values: "SELECT", "INSERT", "UPDATE", "REFERENCES".
	// Note: if this parameter is not passed in, it means to clear the permission.
	ColumnPrivileges []*ColumnPrivilege `json:"ColumnPrivileges,omitempty" name:"ColumnPrivileges"`

	// If this parameter is specified, permissions are modified in batches. Valid values: `grant`, `revoke`.
	ModifyAction *string `json:"ModifyAction,omitempty" name:"ModifyAction"`
}

func (r *ModifyAccountPrivilegesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccountPrivilegesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Accounts")
	delete(f, "GlobalPrivileges")
	delete(f, "DatabasePrivileges")
	delete(f, "TablePrivileges")
	delete(f, "ColumnPrivileges")
	delete(f, "ModifyAction")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAccountPrivilegesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAccountPrivilegesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Async task request ID, which can be used to query the execution result of an async task.
		AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAccountPrivilegesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccountPrivilegesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAutoRenewFlagRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// Auto-renewal flag. Value range: 0 (auto-renewal not enabled), 1 (auto-renewal enabled).
	AutoRenew *int64 `json:"AutoRenew,omitempty" name:"AutoRenew"`
}

func (r *ModifyAutoRenewFlagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAutoRenewFlagRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "AutoRenew")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAutoRenewFlagRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAutoRenewFlagResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAutoRenewFlagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAutoRenewFlagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBackupConfigRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Backup file retention period in days. Value range: 7-732.
	ExpireDays *int64 `json:"ExpireDays,omitempty" name:"ExpireDays"`

	// (This parameter will be disused. The `BackupTimeWindow` parameter is recommended.) Backup time range in the format of 02:00-06:00, with the start time and end time on the hour. Valid values: 00:00-12:00, 02:00-06:00, 06:00-10:00, 10:00-14:00, 14:00-18:00, 18:00-22:00, 22:00-02:00.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Automatic backup mode. Only `physical` (physical cold backup) is supported
	BackupMethod *string `json:"BackupMethod,omitempty" name:"BackupMethod"`

	// Binlog retention period in days. Value range: 7-732. It cannot be greater than the retention period of backup files.
	BinlogExpireDays *int64 `json:"BinlogExpireDays,omitempty" name:"BinlogExpireDays"`

	// Backup time window; for example, to set up backup between 10:00 and 14:00 on every Tuesday and Sunday, you should set this parameter as follows: {"Monday": "", "Tuesday": "10:00-14:00", "Wednesday": "", "Thursday": "", "Friday": "", "Saturday": "", "Sunday": "10:00-14:00"} (Note: You can set up backup on different days, but the backup time windows need to be the same. If this field is set, the `StartTime` field will be ignored)
	BackupTimeWindow *CommonTimeWindow `json:"BackupTimeWindow,omitempty" name:"BackupTimeWindow"`
}

func (r *ModifyBackupConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBackupConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ExpireDays")
	delete(f, "StartTime")
	delete(f, "BackupMethod")
	delete(f, "BinlogExpireDays")
	delete(f, "BackupTimeWindow")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyBackupConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBackupConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyBackupConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBackupConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBackupDownloadRestrictionRequest struct {
	*tchttp.BaseRequest

	// Valid values: `NoLimit` (backups can be downloaded over both private and public networks with any IPs), `LimitOnlyIntranet` (backups can be downloaded over the private network with any private IPs), `Customize` (backups can be downloaded over specified VPCs with specified IPs). The `LimitVpc` and `LimitIp` parameters are valid only when this parameter is set to `Customize`.
	LimitType *string `json:"LimitType,omitempty" name:"LimitType"`

	// Valid value: `In` (backups can only be downloaded over the VPCs specified in `LimitVpc`). Default value: `In`.
	VpcComparisonSymbol *string `json:"VpcComparisonSymbol,omitempty" name:"VpcComparisonSymbol"`

	// Valid values: `In` (backups can only be downloaded with the IPs specified in `LimitIp`), `NotIn` (backups cannot be downloaded with the IPs specified in `LimitIp`). Default value: `In`.
	IpComparisonSymbol *string `json:"IpComparisonSymbol,omitempty" name:"IpComparisonSymbol"`

	// VPCs used to restrict backup download.
	LimitVpc []*BackupLimitVpcItem `json:"LimitVpc,omitempty" name:"LimitVpc"`

	// IPs used to restrict backup download.
	LimitIp []*string `json:"LimitIp,omitempty" name:"LimitIp"`
}

func (r *ModifyBackupDownloadRestrictionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBackupDownloadRestrictionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LimitType")
	delete(f, "VpcComparisonSymbol")
	delete(f, "IpComparisonSymbol")
	delete(f, "LimitVpc")
	delete(f, "LimitIp")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyBackupDownloadRestrictionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBackupDownloadRestrictionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyBackupDownloadRestrictionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBackupDownloadRestrictionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDBInstanceNameRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page. You can use the [instance list querying API](https://intl.cloud.tencent.com/document/api/236/15872?from_cn_redirect=1) to query the ID, whose value is the `InstanceId` value in output parameters.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Instance name.
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
}

func (r *ModifyDBInstanceNameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceNameRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "InstanceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDBInstanceNameRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDBInstanceNameResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDBInstanceNameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDBInstanceProjectRequest struct {
	*tchttp.BaseRequest

	// Array of instance IDs in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page. You can use the [instance list querying API](https://intl.cloud.tencent.com/document/api/236/15872?from_cn_redirect=1) to query the ID, whose value is the `InstanceId` value in output parameters.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// Project ID.
	NewProjectId *int64 `json:"NewProjectId,omitempty" name:"NewProjectId"`
}

func (r *ModifyDBInstanceProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "NewProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDBInstanceProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDBInstanceProjectResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDBInstanceProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDBInstanceSecurityGroupsRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cdb-c1nl9rpv or cdbro-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// List of IDs of security groups to be modified, which is an array of one or more security group IDs.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`

	// This parameter takes effect only when the ID of read-only replica is passed in. If this parameter is set to `False` or left empty, the security groups bound with the RO group of the read-only replicas will be modified. If this parameter is set to `True`, the security groups bound with the read-only replica itself will be modified.
	ForReadonlyInstance *bool `json:"ForReadonlyInstance,omitempty" name:"ForReadonlyInstance"`
}

func (r *ModifyDBInstanceSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceSecurityGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SecurityGroupIds")
	delete(f, "ForReadonlyInstance")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDBInstanceSecurityGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDBInstanceSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDBInstanceSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceSecurityGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDBInstanceVipVportRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page. You can use the [instance list querying API](https://intl.cloud.tencent.com/document/api/236/15872?from_cn_redirect=1) to query the ID, whose value is the `InstanceId` value in output parameters.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Destination IP. Either this parameter or `DstPort` must be passed in.
	DstIp *string `json:"DstIp,omitempty" name:"DstIp"`

	// Destination port number. Value range: [1024-65535]. Either this parameter or `DstIp` must be passed in.
	DstPort *int64 `json:"DstPort,omitempty" name:"DstPort"`

	// Unified VPC ID
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// Unified subnet ID.
	UniqSubnetId *string `json:"UniqSubnetId,omitempty" name:"UniqSubnetId"`

	// Repossession duration in hours for old IP in the original network when changing from the basic network to VPC or changing the VPC subnet. Value range: 0-168 hours. Default value: 24 hours.
	ReleaseDuration *int64 `json:"ReleaseDuration,omitempty" name:"ReleaseDuration"`
}

func (r *ModifyDBInstanceVipVportRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceVipVportRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "DstIp")
	delete(f, "DstPort")
	delete(f, "UniqVpcId")
	delete(f, "UniqSubnetId")
	delete(f, "ReleaseDuration")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDBInstanceVipVportRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDBInstanceVipVportResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Async task ID. (This returned field has been disused)
	// Note: this field may return null, indicating that no valid values can be obtained.
		AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDBInstanceVipVportResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceVipVportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceParamRequest struct {
	*tchttp.BaseRequest

	// List of short instance IDs.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// List of parameters to be modified. Every element is a combination of `Name` (parameter name) and `CurrentValue` (new value).
	ParamList []*Parameter `json:"ParamList,omitempty" name:"ParamList"`

	// Template ID. At least one of `ParamList` and `TemplateId` must be passed in.
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// When to perform the parameter adjustment task. Default value: 0. Valid values: 0 - execute immediately, 1 - execute during window. When its value is 1, only one instance ID can be passed in (i.e., only one `InstanceIds` can be passed in).
	WaitSwitch *int64 `json:"WaitSwitch,omitempty" name:"WaitSwitch"`
}

func (r *ModifyInstanceParamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceParamRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "ParamList")
	delete(f, "TemplateId")
	delete(f, "WaitSwitch")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstanceParamRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceParamResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Async task ID, which can be used to query task progress.
		AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyInstanceParamResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceParamResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceTagRequest struct {
	*tchttp.BaseRequest

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Tag to be added or modified.
	ReplaceTags []*TagInfo `json:"ReplaceTags,omitempty" name:"ReplaceTags"`

	// Tag to be deleted.
	DeleteTags []*TagInfo `json:"DeleteTags,omitempty" name:"DeleteTags"`
}

func (r *ModifyInstanceTagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceTagRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ReplaceTags")
	delete(f, "DeleteTags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstanceTagRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceTagResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyInstanceTagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceTagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNameOrDescByDpIdRequest struct {
	*tchttp.BaseRequest

	// ID of a placement group.
	DeployGroupId *string `json:"DeployGroupId,omitempty" name:"DeployGroupId"`

	// Name of a placement group, which can contain up to 60 characters. The placement group name and description cannot both be empty.
	DeployGroupName *string `json:"DeployGroupName,omitempty" name:"DeployGroupName"`

	// Description of a placement group, which can contain up to 200 characters. The placement group name and description cannot both be empty.
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *ModifyNameOrDescByDpIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNameOrDescByDpIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeployGroupId")
	delete(f, "DeployGroupName")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyNameOrDescByDpIdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNameOrDescByDpIdResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyNameOrDescByDpIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNameOrDescByDpIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyParamTemplateRequest struct {
	*tchttp.BaseRequest

	// Template ID.
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// Template name.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Template description.
	Description *string `json:"Description,omitempty" name:"Description"`

	// List of parameters.
	ParamList []*Parameter `json:"ParamList,omitempty" name:"ParamList"`
}

func (r *ModifyParamTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyParamTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	delete(f, "Name")
	delete(f, "Description")
	delete(f, "ParamList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyParamTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyParamTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyParamTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyParamTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRoGroupInfoRequest struct {
	*tchttp.BaseRequest

	// RO group ID.
	RoGroupId *string `json:"RoGroupId,omitempty" name:"RoGroupId"`

	// RO group details.
	RoGroupInfo *RoGroupAttr `json:"RoGroupInfo,omitempty" name:"RoGroupInfo"`

	// Weights of instances in RO group. If the weighting mode of an RO group is changed to custom mode, this parameter must be set, and a weight value needs to be set for each RO instance.
	RoWeightValues []*RoWeightValue `json:"RoWeightValues,omitempty" name:"RoWeightValues"`

	// Whether to rebalance the loads of read-only replicas in the RO group. Valid values: `1` (yes), `0` (no). Default value: `0`. If this parameter is set to `1`, connections to the read-only replicas in the RO group will be interrupted transiently. Please ensure that your application has a reconnection mechanism.
	IsBalanceRoLoad *int64 `json:"IsBalanceRoLoad,omitempty" name:"IsBalanceRoLoad"`
}

func (r *ModifyRoGroupInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRoGroupInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RoGroupId")
	delete(f, "RoGroupInfo")
	delete(f, "RoWeightValues")
	delete(f, "IsBalanceRoLoad")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRoGroupInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRoGroupInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyRoGroupInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRoGroupInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRoReplicationDelayRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Replication delay in seconds. Value range: 1 to 259200.
	ReplicationDelay *int64 `json:"ReplicationDelay,omitempty" name:"ReplicationDelay"`
}

func (r *ModifyRoReplicationDelayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRoReplicationDelayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ReplicationDelay")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRoReplicationDelayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRoReplicationDelayResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyRoReplicationDelayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRoReplicationDelayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTimeWindowRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cdb-c1nl9rpv or cdbro-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Time period available for maintenance after modification in the format of 10:00-12:00. Each period lasts from half an hour to three hours, with the start time and end time aligned by half-hour. Up to two time periods can be set. Start and end time range: [00:00, 24:00].
	TimeRanges []*string `json:"TimeRanges,omitempty" name:"TimeRanges"`

	// Specifies for which day to modify the time period. Value range: Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday. If it is not specified or is left blank, the time period will be modified for every day by default.
	Weekdays []*string `json:"Weekdays,omitempty" name:"Weekdays"`
}

func (r *ModifyTimeWindowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTimeWindowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "TimeRanges")
	delete(f, "Weekdays")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTimeWindowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTimeWindowResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyTimeWindowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTimeWindowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OfflineIsolatedInstancesRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *OfflineIsolatedInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OfflineIsolatedInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OfflineIsolatedInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type OfflineIsolatedInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OfflineIsolatedInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OfflineIsolatedInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OpenDBInstanceGTIDRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *OpenDBInstanceGTIDRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenDBInstanceGTIDRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OpenDBInstanceGTIDRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type OpenDBInstanceGTIDResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Async task request ID, which can be used to query the execution result of an async task.
		AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OpenDBInstanceGTIDResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenDBInstanceGTIDResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OpenWanServiceRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page. You can use the [instance list querying API](https://intl.cloud.tencent.com/document/api/236/15872?from_cn_redirect=1) to query the ID, whose value is the `InstanceId` value in output parameters.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *OpenWanServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenWanServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OpenWanServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type OpenWanServiceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Async task request ID, which can be used to query the execution result of an async task.
		AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OpenWanServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenWanServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Outbound struct {

	// Policy, which can be ACCEPT or DROP
	Action *string `json:"Action,omitempty" name:"Action"`

	// Destination IP or IP range, such as 172.16.0.0/12
	CidrIp *string `json:"CidrIp,omitempty" name:"CidrIp"`

	// Port or port range
	PortRange *string `json:"PortRange,omitempty" name:"PortRange"`

	// Network protocol. UDP and TCP are supported
	IpProtocol *string `json:"IpProtocol,omitempty" name:"IpProtocol"`

	// The direction of the rule, which is OUTPUT for inbound rules
	Dir *string `json:"Dir,omitempty" name:"Dir"`

	// Rule description
	Desc *string `json:"Desc,omitempty" name:"Desc"`
}

type ParamInfo struct {

	// Parameter name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Parameter value
	Value *string `json:"Value,omitempty" name:"Value"`
}

type ParamRecord struct {

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Parameter name
	ParamName *string `json:"ParamName,omitempty" name:"ParamName"`

	// Parameter value before modification
	OldValue *string `json:"OldValue,omitempty" name:"OldValue"`

	// Parameter value after modification
	NewValue *string `json:"NewValue,omitempty" name:"NewValue"`

	// Whether the parameter is modified successfully
	IsSucess *bool `json:"IsSucess,omitempty" name:"IsSucess"`

	// Modification time
	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
}

type ParamTemplateInfo struct {

	// Parameter template ID
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// Parameter template name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Parameter template description
	Description *string `json:"Description,omitempty" name:"Description"`

	// Instance engine version
	EngineVersion *string `json:"EngineVersion,omitempty" name:"EngineVersion"`
}

type Parameter struct {

	// Parameter name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Parameter value
	CurrentValue *string `json:"CurrentValue,omitempty" name:"CurrentValue"`
}

type ParameterDetail struct {

	// Parameter name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Parameter type
	ParamType *string `json:"ParamType,omitempty" name:"ParamType"`

	// Default value of the parameter
	Default *string `json:"Default,omitempty" name:"Default"`

	// Parameter description
	Description *string `json:"Description,omitempty" name:"Description"`

	// Current value of the parameter
	CurrentValue *string `json:"CurrentValue,omitempty" name:"CurrentValue"`

	// Whether the database needs to be restarted for the modified parameter to take effect. Value range: 0 (no); 1 (yes)
	NeedReboot *int64 `json:"NeedReboot,omitempty" name:"NeedReboot"`

	// Maximum value of the parameter
	Max *int64 `json:"Max,omitempty" name:"Max"`

	// Minimum value of the parameter
	Min *int64 `json:"Min,omitempty" name:"Min"`

	// Enumerated values of the parameter. It is null if the parameter is non-enumerated
	EnumValue []*string `json:"EnumValue,omitempty" name:"EnumValue"`
}

type RegionSellConf struct {

	// Region name
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`

	// Area
	Area *string `json:"Area,omitempty" name:"Area"`

	// Whether it is a default region
	IsDefaultRegion *int64 `json:"IsDefaultRegion,omitempty" name:"IsDefaultRegion"`

	// Region name
	Region *string `json:"Region,omitempty" name:"Region"`

	// Sale configuration of the AZ
	ZonesConf []*ZoneSellConf `json:"ZonesConf,omitempty" name:"ZonesConf"`
}

type ReleaseIsolatedDBInstancesRequest struct {
	*tchttp.BaseRequest

	// Array of instance IDs in the format of `cdb-c1nl9rpv`. It is the same as the instance ID displayed on the TencentDB Console page. You can use the [DescribeDBInstances](https://intl.cloud.tencent.com/document/api/236/15872?from_cn_redirect=1) API to query the ID, whose value is the `InstanceId` value in the output parameters.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *ReleaseIsolatedDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReleaseIsolatedDBInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReleaseIsolatedDBInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ReleaseIsolatedDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Deisolation result set.
		Items []*ReleaseResult `json:"Items,omitempty" name:"Items"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReleaseIsolatedDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReleaseIsolatedDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReleaseResult struct {

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Result value of instance deisolation. A returned value of 0 indicates success.
	Code *int64 `json:"Code,omitempty" name:"Code"`

	// Error message for instance deisolation.
	Message *string `json:"Message,omitempty" name:"Message"`
}

type RestartDBInstancesRequest struct {
	*tchttp.BaseRequest

	// Array of instance IDs in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *RestartDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestartDBInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RestartDBInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type RestartDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Async task request ID, which can be used to query the execution result of an async task.
		AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RestartDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestartDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RoGroup struct {

	// Read-only group mode. Valid values: `alone` (the system assigns a read-only group automatically), `allinone` (a new read-only group will be created), `join` (an existing read-only group will be used).
	RoGroupMode *string `json:"RoGroupMode,omitempty" name:"RoGroupMode"`

	// Read-only group ID.
	RoGroupId *string `json:"RoGroupId,omitempty" name:"RoGroupId"`

	// Read-only group name.
	RoGroupName *string `json:"RoGroupName,omitempty" name:"RoGroupName"`

	// Whether to enable the function of isolating an instance that exceeds the latency threshold. If it is enabled, when the latency between the read-only instance and the primary instance exceeds the latency threshold, the read-only instance will be isolated. Valid values: 1 (enabled), 0 (not enabled)
	RoOfflineDelay *int64 `json:"RoOfflineDelay,omitempty" name:"RoOfflineDelay"`

	// Latency threshold
	RoMaxDelayTime *int64 `json:"RoMaxDelayTime,omitempty" name:"RoMaxDelayTime"`

	// Minimum number of instances to be retained. If the number of the purchased read-only instances is smaller than the set value, they will not be removed.
	MinRoInGroup *int64 `json:"MinRoInGroup,omitempty" name:"MinRoInGroup"`

	// Read/write weight distribution mode. Valid values: `system` (weights are assigned by the system automatically), `custom` (weights are customized)
	WeightMode *string `json:"WeightMode,omitempty" name:"WeightMode"`

	// Weight value.
	Weight *int64 `json:"Weight,omitempty" name:"Weight"`

	// Details of read-only instances in read-only group
	RoInstances []*RoInstanceInfo `json:"RoInstances,omitempty" name:"RoInstances"`

	// Private IP of read-only group.
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// Private network port number of read-only group.
	Vport *int64 `json:"Vport,omitempty" name:"Vport"`

	// VPC ID.
	// Note: this field may return null, indicating that no valid values can be obtained.
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// Subnet ID.
	// Note: this field may return null, indicating that no valid values can be obtained.
	UniqSubnetId *string `json:"UniqSubnetId,omitempty" name:"UniqSubnetId"`

	// Read-only group region.
	// Note: this field may return null, indicating that no valid values can be obtained.
	RoGroupRegion *string `json:"RoGroupRegion,omitempty" name:"RoGroupRegion"`

	// Read-only group AZ.
	// Note: this field may return null, indicating that no valid values can be obtained.
	RoGroupZone *string `json:"RoGroupZone,omitempty" name:"RoGroupZone"`
}

type RoGroupAttr struct {

	// RO group name.
	RoGroupName *string `json:"RoGroupName,omitempty" name:"RoGroupName"`

	// Maximum delay threshold for RO instances in seconds. Minimum value: 1. Please note that this value will take effect only if an instance removal policy is enabled in the RO group.
	RoMaxDelayTime *int64 `json:"RoMaxDelayTime,omitempty" name:"RoMaxDelayTime"`

	// Whether to enable instance removal. Valid values: 1 (enabled), 0 (not enabled). Please note that if instance removal is enabled, the delay threshold parameter (`RoMaxDelayTime`) must be set.
	RoOfflineDelay *int64 `json:"RoOfflineDelay,omitempty" name:"RoOfflineDelay"`

	// Minimum number of instances to be retained, which can be set to any value less than or equal to the number of RO instances in the RO group. Please note that if this value is set to be greater than the number of RO instances, no removal will be performed, and if it is set to 0, all instances with an excessive delay will be removed.
	MinRoInGroup *int64 `json:"MinRoInGroup,omitempty" name:"MinRoInGroup"`

	// Weighting mode. Supported values include `system` (automatically assigned by the system) and `custom` (defined by user). Please note that if the `custom` mode is selected, the RO instance weight configuration parameter (RoWeightValues) must be set.
	WeightMode *string `json:"WeightMode,omitempty" name:"WeightMode"`
}

type RoInstanceInfo struct {

	// Master instance ID corresponding to the RO group
	MasterInstanceId *string `json:"MasterInstanceId,omitempty" name:"MasterInstanceId"`

	// RO instance status in the RO group. Value range: online, offline
	RoStatus *string `json:"RoStatus,omitempty" name:"RoStatus"`

	// Last deactivation time of a RO instance in the RO group
	OfflineTime *string `json:"OfflineTime,omitempty" name:"OfflineTime"`

	// RO instance weight in the RO group
	Weight *int64 `json:"Weight,omitempty" name:"Weight"`

	// RO instance region name, such as ap-shanghai
	Region *string `json:"Region,omitempty" name:"Region"`

	// Name of RO AZ, such as ap-shanghai-1
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// RO instance ID in the format of cdbro-c1nl9rpv
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// RO instance status. Value range: 0 (creating), 1 (running), 4 (deleting)
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// Instance type. Value range: 1 (primary), 2 (disaster recovery), 3 (read-only)
	InstanceType *int64 `json:"InstanceType,omitempty" name:"InstanceType"`

	// RO instance name
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// Pay-as-you-go billing status. Value range: 1 (normal), 2 (in arrears)
	HourFeeStatus *int64 `json:"HourFeeStatus,omitempty" name:"HourFeeStatus"`

	// RO instance task status. Value range: <br>0 - no task <br>1 - upgrading <br>2 - importing data <br>3 - activating secondary <br>4 - public network access enabled <br>5 - batch operation in progress <br>6 - rolling back <br>7 - public network access not enabled <br>8 - modifying password <br>9 - renaming instance <br>10 - restarting <br>12 - migrating self-built instance <br>13 - dropping table <br>14 - creating and syncing disaster recovery instance
	TaskStatus *int64 `json:"TaskStatus,omitempty" name:"TaskStatus"`

	// RO instance memory size in MB
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// RO instance disk size in GB
	Volume *int64 `json:"Volume,omitempty" name:"Volume"`

	// Queries per second
	Qps *int64 `json:"Qps,omitempty" name:"Qps"`

	// Private IP address of the RO instance
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// Access port of the RO instance
	Vport *int64 `json:"Vport,omitempty" name:"Vport"`

	// VPC ID of the RO instance
	VpcId *int64 `json:"VpcId,omitempty" name:"VpcId"`

	// VPC subnet ID of the RO instance
	SubnetId *int64 `json:"SubnetId,omitempty" name:"SubnetId"`

	// RO instance specification description. Value range: CUSTOM
	DeviceType *string `json:"DeviceType,omitempty" name:"DeviceType"`

	// Database engine version of the read-only replica. Valid values: `5.1`, `5.5`, `5.6`, `5.7`, `8.0`
	EngineVersion *string `json:"EngineVersion,omitempty" name:"EngineVersion"`

	// RO instance expiration time in the format of yyyy-mm-dd hh:mm:ss. If it is a pay-as-you-go instance, the value of this field is 0000-00-00 00:00:00
	DeadlineTime *string `json:"DeadlineTime,omitempty" name:"DeadlineTime"`

	// RO instance billing method. Value range: 0 (monthly subscribed), 1 (pay-as-you-go), 2 (monthly postpaid)
	PayType *int64 `json:"PayType,omitempty" name:"PayType"`
}

type RoVipInfo struct {

	// VIP status of the read-only instance
	RoVipStatus *int64 `json:"RoVipStatus,omitempty" name:"RoVipStatus"`

	// VPC subnet of the read-only instance
	RoSubnetId *int64 `json:"RoSubnetId,omitempty" name:"RoSubnetId"`

	// VPC of the read-only instance
	RoVpcId *int64 `json:"RoVpcId,omitempty" name:"RoVpcId"`

	// VIP port number of the read-only instance
	RoVport *int64 `json:"RoVport,omitempty" name:"RoVport"`

	// VIP of the read-only instance
	RoVip *string `json:"RoVip,omitempty" name:"RoVip"`
}

type RoWeightValue struct {

	// RO instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Weight value. Value range: [0, 100].
	Weight *int64 `json:"Weight,omitempty" name:"Weight"`
}

type RollbackDBName struct {

	// Original database name before rollback
	// Note: this field may return null, indicating that no valid values can be obtained.
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`

	// New database name after rollback
	// Note: this field may return null, indicating that no valid values can be obtained.
	NewDatabaseName *string `json:"NewDatabaseName,omitempty" name:"NewDatabaseName"`
}

type RollbackInstancesInfo struct {

	// TencentDB instance ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Rollback policy. Valid values: `table` (ultrafast mode), `db` (faster mode), and `full` (fast mode). Default value: `full`. In the ultrafast mode, only backups and binlogs of the tables specified by the `Tables` parameter are imported; if `Tables` does not include all of the tables involved in cross-table operations, the rollback may fail; and the `Database` parameter must be left empty. In the faster mode, only backups and binlogs of the databases specified by the `Databases` parameter are imported, and if `Databases` does not include all of the databases involved in cross-database operations, the rollback may fail. In the fast mode, backups and binlogs of the entire instance will be imported in a speed slower than the other modes.
	Strategy *string `json:"Strategy,omitempty" name:"Strategy"`

	// Database rollback time in the format of yyyy-mm-dd hh:mm:ss
	RollbackTime *string `json:"RollbackTime,omitempty" name:"RollbackTime"`

	// Information of the databases to be rolled back, which means rollback at the database level
	// Note: this field may return null, indicating that no valid values can be obtained.
	Databases []*RollbackDBName `json:"Databases,omitempty" name:"Databases"`

	// Information of the tables to be rolled back, which means rollback at the table level
	// Note: this field may return null, indicating that no valid values can be obtained.
	Tables []*RollbackTables `json:"Tables,omitempty" name:"Tables"`
}

type RollbackTableName struct {

	// Original table name before rollback
	// Note: this field may return null, indicating that no valid values can be obtained.
	TableName *string `json:"TableName,omitempty" name:"TableName"`

	// New table name after rollback
	// Note: this field may return null, indicating that no valid values can be obtained.
	NewTableName *string `json:"NewTableName,omitempty" name:"NewTableName"`
}

type RollbackTables struct {

	// Database name
	// Note: this field may return null, indicating that no valid values can be obtained.
	Database *string `json:"Database,omitempty" name:"Database"`

	// Table details
	// Note: this field may return null, indicating that no valid values can be obtained.
	Table []*RollbackTableName `json:"Table,omitempty" name:"Table"`
}

type RollbackTask struct {

	// Task execution information.
	Info *string `json:"Info,omitempty" name:"Info"`

	// Task execution result. Valid values: INITIAL: initializing, RUNNING: running, SUCCESS: succeeded, FAILED: failed, KILLED: terminated, REMOVED: deleted, PAUSED: paused.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Task execution progress. Value range: [0,100].
	Progress *int64 `json:"Progress,omitempty" name:"Progress"`

	// Task start time.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Task end time.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Rollback task details.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Detail []*RollbackInstancesInfo `json:"Detail,omitempty" name:"Detail"`
}

type RollbackTimeRange struct {

	// Start time available for rollback in the format of yyyy-MM-dd HH:mm:ss, such as 2016-10-29 01:06:04
	Begin *string `json:"Begin,omitempty" name:"Begin"`

	// End time available for rollback in the format of yyyy-MM-dd HH:mm:ss, such as 2016-11-02 11:44:47
	End *string `json:"End,omitempty" name:"End"`
}

type SecurityGroup struct {

	// Project ID
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Creation time in the format of yyyy-mm-dd hh:mm:ss
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Inbound rule
	Inbound []*Inbound `json:"Inbound,omitempty" name:"Inbound"`

	// Outbound rule
	Outbound []*Outbound `json:"Outbound,omitempty" name:"Outbound"`

	// Security group ID
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`

	// Security group name
	SecurityGroupName *string `json:"SecurityGroupName,omitempty" name:"SecurityGroupName"`

	// Security group remarks
	SecurityGroupRemark *string `json:"SecurityGroupRemark,omitempty" name:"SecurityGroupRemark"`
}

type SellConfig struct {

	// (Disused) Device type
	Device *string `json:"Device,omitempty" name:"Device"`

	// (Disused) Purchasable specification description 
	Type *string `json:"Type,omitempty" name:"Type"`

	// (Disused) Instance type 
	CdbType *string `json:"CdbType,omitempty" name:"CdbType"`

	// Memory size in MB
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// CPU core count
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

	// Minimum disk size in GB
	VolumeMin *int64 `json:"VolumeMin,omitempty" name:"VolumeMin"`

	// Maximum disk size in GB
	VolumeMax *int64 `json:"VolumeMax,omitempty" name:"VolumeMax"`

	// Disk increment in GB
	VolumeStep *int64 `json:"VolumeStep,omitempty" name:"VolumeStep"`

	// Number of connections
	Connection *int64 `json:"Connection,omitempty" name:"Connection"`

	// Queries per second
	Qps *int64 `json:"Qps,omitempty" name:"Qps"`

	// IOs per second
	Iops *int64 `json:"Iops,omitempty" name:"Iops"`

	// Application scenario description
	Info *string `json:"Info,omitempty" name:"Info"`

	// Status. Value `0` indicates that this specification is purchasable.
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// (Disused) Tag value
	Tag *int64 `json:"Tag,omitempty" name:"Tag"`

	// Instance resource isolation type. Valid values: `UNIVERSAL` (general instance), `EXCLUSIVE` (dedicated instance), `BASIC` (basic instance).
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	DeviceType *string `json:"DeviceType,omitempty" name:"DeviceType"`

	// Instance resource isolation type. Valid values: `UNIVERSAL` (general instance), `EXCLUSIVE` (dedicated instance), `BASIC` (basic instance).
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	DeviceTypeName *string `json:"DeviceTypeName,omitempty" name:"DeviceTypeName"`
}

type SellType struct {

	// Name of the purchasable instance
	TypeName *string `json:"TypeName,omitempty" name:"TypeName"`

	// Kernel version number
	EngineVersion []*string `json:"EngineVersion,omitempty" name:"EngineVersion"`

	// Configuration details of a purchasable specification
	Configs []*SellConfig `json:"Configs,omitempty" name:"Configs"`
}

type SlaveConfig struct {

	// Replication mode of the secondary database. Value range: async, semi-sync
	ReplicationMode *string `json:"ReplicationMode,omitempty" name:"ReplicationMode"`

	// AZ name of the secondary database, such as ap-shanghai-1
	Zone *string `json:"Zone,omitempty" name:"Zone"`
}

type SlaveInfo struct {

	// Information of secondary server 1
	First *SlaveInstanceInfo `json:"First,omitempty" name:"First"`

	// Information of secondary server 2
	// Note: This field may return null, indicating that no valid values can be obtained.
	Second *SlaveInstanceInfo `json:"Second,omitempty" name:"Second"`
}

type SlaveInstanceInfo struct {

	// Port number
	Vport *int64 `json:"Vport,omitempty" name:"Vport"`

	// Region information
	Region *string `json:"Region,omitempty" name:"Region"`

	// Virtual IP information
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// AZ information
	Zone *string `json:"Zone,omitempty" name:"Zone"`
}

type SlowLogInfo struct {

	// Backup filename
	Name *string `json:"Name,omitempty" name:"Name"`

	// Backup file size in bytes
	Size *int64 `json:"Size,omitempty" name:"Size"`

	// Backup snapshot time in the format of yyyy-MM-dd HH:mm:ss, such as 2016-03-17 02:10:37
	Date *string `json:"Date,omitempty" name:"Date"`

	// Download address on the private network
	IntranetUrl *string `json:"IntranetUrl,omitempty" name:"IntranetUrl"`

	// Download address on the public network
	InternetUrl *string `json:"InternetUrl,omitempty" name:"InternetUrl"`

	// Log type. Value range: slowlog (slow log)
	Type *string `json:"Type,omitempty" name:"Type"`
}

type SlowLogItem struct {

	// SQL execution time.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Timestamp *uint64 `json:"Timestamp,omitempty" name:"Timestamp"`

	// SQL execution duration in seconds.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	QueryTime *float64 `json:"QueryTime,omitempty" name:"QueryTime"`

	// SQL statement.
	// Note: this field may return null, indicating that no valid values can be obtained.
	SqlText *string `json:"SqlText,omitempty" name:"SqlText"`

	// Client address.
	// Note: this field may return null, indicating that no valid values can be obtained.
	UserHost *string `json:"UserHost,omitempty" name:"UserHost"`

	// Username.
	// Note: this field may return null, indicating that no valid values can be obtained.
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// Database name.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Database *string `json:"Database,omitempty" name:"Database"`

	// Lock duration in seconds.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	LockTime *float64 `json:"LockTime,omitempty" name:"LockTime"`

	// Number of scanned rows.
	// Note: this field may return null, indicating that no valid values can be obtained.
	RowsExamined *int64 `json:"RowsExamined,omitempty" name:"RowsExamined"`

	// Number of rows in result set.
	// Note: this field may return null, indicating that no valid values can be obtained.
	RowsSent *int64 `json:"RowsSent,omitempty" name:"RowsSent"`

	// SQL template.
	// Note: this field may return null, indicating that no valid values can be obtained.
	SqlTemplate *string `json:"SqlTemplate,omitempty" name:"SqlTemplate"`

	// SQL statement MD5.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Md5 *string `json:"Md5,omitempty" name:"Md5"`
}

type SqlFileInfo struct {

	// Upload time
	UploadTime *string `json:"UploadTime,omitempty" name:"UploadTime"`

	// Upload progress
	UploadInfo *UploadInfo `json:"UploadInfo,omitempty" name:"UploadInfo"`

	// Filename
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// File size in bytes
	FileSize *int64 `json:"FileSize,omitempty" name:"FileSize"`

	// Whether upload is finished. Valid values: 0 (not completed), 1 (completed)
	IsUploadFinished *int64 `json:"IsUploadFinished,omitempty" name:"IsUploadFinished"`

	// File ID
	FileId *string `json:"FileId,omitempty" name:"FileId"`
}

type StartBatchRollbackRequest struct {
	*tchttp.BaseRequest

	// Details of the instance for rollback
	Instances []*RollbackInstancesInfo `json:"Instances,omitempty" name:"Instances"`
}

func (r *StartBatchRollbackRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartBatchRollbackRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Instances")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartBatchRollbackRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type StartBatchRollbackResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Async task request ID, which can be used to query the execution result of an async task.
		AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StartBatchRollbackResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartBatchRollbackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StartDelayReplicationRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Delayed replication mode. Valid values: `DEFAULT` (replicate according to the specified replication delay), `GTID` (replicate according to the specified GTID), `DUE_TIME` (replicate according to the specified point in time).
	DelayReplicationType *string `json:"DelayReplicationType,omitempty" name:"DelayReplicationType"`

	// Specified point in time. Default value: 0. The maximum value cannot be later than the current time.
	DueTime *int64 `json:"DueTime,omitempty" name:"DueTime"`

	// Specified GITD. This parameter is required when the delayed replication mode is `GTID`.
	Gtid *string `json:"Gtid,omitempty" name:"Gtid"`
}

func (r *StartDelayReplicationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartDelayReplicationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "DelayReplicationType")
	delete(f, "DueTime")
	delete(f, "Gtid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartDelayReplicationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type StartDelayReplicationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Delayed replication task ID. This parameter will be returned if `DelayReplicationType` is not `DEFAULT`. It can be used to view the status of the delayed replication task.
	// Note: this field may return null, indicating that no valid values can be obtained.
		AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StartDelayReplicationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartDelayReplicationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StopDBImportJobRequest struct {
	*tchttp.BaseRequest

	// Async task request ID.
	AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`
}

func (r *StopDBImportJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopDBImportJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AsyncRequestId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopDBImportJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type StopDBImportJobResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StopDBImportJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopDBImportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StopDelayReplicationRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *StopDelayReplicationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopDelayReplicationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopDelayReplicationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type StopDelayReplicationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StopDelayReplicationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopDelayReplicationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StopRollbackRequest struct {
	*tchttp.BaseRequest

	// ID of the instance whose rollback task is canceled
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *StopRollbackRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopRollbackRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopRollbackRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type StopRollbackResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Async task request ID
		AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StopRollbackResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopRollbackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SwitchDBInstanceMasterSlaveRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Specifies the replica server to switched to. Valid values: `first` (the first replica server), `second` (the second replica server). Default value: `first`. `second` is valid only for a multi-AZ instance.
	DstSlave *string `json:"DstSlave,omitempty" name:"DstSlave"`

	// Whether to force the switch. Valid values: `True`, `False` (default). If this parameter is set to `True`, instance data may be lost during the switch.
	ForceSwitch *bool `json:"ForceSwitch,omitempty" name:"ForceSwitch"`

	// Whether to perform the switch during a time window. Valid values: `True`, `False` (default). If `ForceSwitch` is set to `True`, this parameter is invalid.
	WaitSwitch *bool `json:"WaitSwitch,omitempty" name:"WaitSwitch"`
}

func (r *SwitchDBInstanceMasterSlaveRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchDBInstanceMasterSlaveRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "DstSlave")
	delete(f, "ForceSwitch")
	delete(f, "WaitSwitch")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SwitchDBInstanceMasterSlaveRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type SwitchDBInstanceMasterSlaveResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Async task ID
		AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SwitchDBInstanceMasterSlaveResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchDBInstanceMasterSlaveResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SwitchDrInstanceToMasterRequest struct {
	*tchttp.BaseRequest

	// Disaster recovery instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed in the TencentDB console.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *SwitchDrInstanceToMasterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchDrInstanceToMasterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SwitchDrInstanceToMasterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type SwitchDrInstanceToMasterResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Async task request ID, which can be used to query the execution result of an async task
		AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SwitchDrInstanceToMasterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchDrInstanceToMasterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SwitchForUpgradeRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *SwitchForUpgradeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchForUpgradeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SwitchForUpgradeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type SwitchForUpgradeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SwitchForUpgradeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchForUpgradeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TablePrivilege struct {

	// Database name
	Database *string `json:"Database,omitempty" name:"Database"`

	// Table name
	Table *string `json:"Table,omitempty" name:"Table"`

	// Permission information
	Privileges []*string `json:"Privileges,omitempty" name:"Privileges"`
}

type TagInfo struct {

	// Tag key
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// Tag value
	TagValue []*string `json:"TagValue,omitempty" name:"TagValue"`
}

type TagInfoItem struct {

	// Tag key
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// Tag value
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type TagInfoUnit struct {

	// Tag key
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// Tag value
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type TagsInfoOfInstance struct {

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Tag information
	Tags []*TagInfoUnit `json:"Tags,omitempty" name:"Tags"`
}

type TaskDetail struct {

	// Error code.
	Code *int64 `json:"Code,omitempty" name:"Code"`

	// Error message.
	Message *string `json:"Message,omitempty" name:"Message"`

	// ID of an instance task.
	JobId *int64 `json:"JobId,omitempty" name:"JobId"`

	// Instance task progress.
	Progress *int64 `json:"Progress,omitempty" name:"Progress"`

	// Instance task status. Valid values:
	// "UNDEFINED" - undefined;
	// "INITIAL" - initializing;
	// "RUNNING" - running;
	// "SUCCEED" - succeeded;
	// "FAILED" - failed;
	// "KILLED" - terminated;
	// "REMOVED" - deleted;
	// "PAUSED" - paused.
	// "WAITING" - waiting (which can be canceled)
	TaskStatus *string `json:"TaskStatus,omitempty" name:"TaskStatus"`

	// Instance task type. Valid values:
	// "ROLLBACK" - rolling back a database;
	// "SQL OPERATION" - performing an SQL operation;
	// "IMPORT DATA" - importing data;
	// "MODIFY PARAM" - setting a parameter;
	// "INITIAL" - initializing a TencentDB instance;
	// "REBOOT" - restarting a TencentDB instance;
	// "OPEN GTID" - enabling GTID of a TencentDB instance;
	// "UPGRADE RO" - upgrading a read-only instance;
	// "BATCH ROLLBACK" - rolling back databases in batches;
	// "UPGRADE MASTER" - upgrading a primary instance;
	// "DROP TABLES" - dropping a TencentDB table;
	// "SWITCH DR TO MASTER" - promoting a disaster recovery instance.
	TaskType *string `json:"TaskType,omitempty" name:"TaskType"`

	// Instance task start time.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Instance task end time.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// ID of an instance associated with a task.
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// Async task request ID.
	AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`
}

type UpgradeDBInstanceEngineVersionRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cdb-c1nl9rpv or cdbro-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page. You can use the [instance list querying API](https://intl.cloud.tencent.com/document/api/236/15872?from_cn_redirect=1) to query the ID, whose value is the `InstanceId` value in output parameters.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Version of primary instance database engine. Value range: 5.6, 5.7
	EngineVersion *string `json:"EngineVersion,omitempty" name:"EngineVersion"`

	// Mode of switch to a new instance. Value range: 0 (switch immediately), 1 (switch within a time window). Default value: 0. If the value is 1, the switch process will be performed within a time window. Or, you can call the [switching to new instance API](https://intl.cloud.tencent.com/document/product/236/15864?from_cn_redirect=1) to trigger the process.
	WaitSwitch *int64 `json:"WaitSwitch,omitempty" name:"WaitSwitch"`

	// Whether to upgrade kernel minor version. Valid values: 1 (upgrade kernel minor version), 0 (upgrade database engine).
	UpgradeSubversion *int64 `json:"UpgradeSubversion,omitempty" name:"UpgradeSubversion"`

	// Delay threshold. Value range: 1-10
	MaxDelayTime *int64 `json:"MaxDelayTime,omitempty" name:"MaxDelayTime"`
}

func (r *UpgradeDBInstanceEngineVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeDBInstanceEngineVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "EngineVersion")
	delete(f, "WaitSwitch")
	delete(f, "UpgradeSubversion")
	delete(f, "MaxDelayTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpgradeDBInstanceEngineVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type UpgradeDBInstanceEngineVersionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Async task ID. The task execution result can be queried using the [async task execution result querying API](https://intl.cloud.tencent.com/document/api/236/20410?from_cn_redirect=1).
		AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpgradeDBInstanceEngineVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeDBInstanceEngineVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpgradeDBInstanceRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of `cdb-c1nl9rpv` or `cdbro-c1nl9rpv`. It is the same as the instance ID displayed on the TencentDB Console page. You can use the [DescribeDBInstances](https://intl.cloud.tencent.com/document/api/236/15872?from_cn_redirect=1) API to query the ID, whose value is the `InstanceId` value in output parameters.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Memory size in MB after upgrade. To ensure that the `Memory` value to be passed in is valid, please use the [DescribeDBZoneConfig](https://intl.cloud.tencent.com/document/product/236/17229?from_cn_redirect=1) API to query the specifications of the memory that can be upgraded to.
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// Disk size in GB after upgrade. To ensure that the `Volume` value to be passed in is valid, please use the [DescribeDBZoneConfig](https://intl.cloud.tencent.com/document/product/236/17229?from_cn_redirect=1) API to query the specifications of the disk that can be upgraded to.
	Volume *int64 `json:"Volume,omitempty" name:"Volume"`

	// Data replication mode. Valid values: 0 (async), 1 (semi-sync), 2 (strong sync). This parameter can be specified when upgrading primary instances and is meaningless for read-only or disaster recovery instances.
	ProtectMode *int64 `json:"ProtectMode,omitempty" name:"ProtectMode"`

	// Deployment mode. Valid values: 0 (single-AZ), 1 (multi-AZ). Default value: 0. This parameter can be specified when upgrading primary instances and is meaningless for read-only or disaster recovery instances.
	DeployMode *int64 `json:"DeployMode,omitempty" name:"DeployMode"`

	// AZ information of secondary database 1, which is the `Zone` value of the instance by default. This parameter can be specified when upgrading primary instances in multi-AZ mode and is meaningless for read-only or disaster recovery instances. You can use the [DescribeDBZoneConfig](https://intl.cloud.tencent.com/document/product/236/17229?from_cn_redirect=1) API to query the supported AZs.
	SlaveZone *string `json:"SlaveZone,omitempty" name:"SlaveZone"`

	// Version of primary instance database engine. Valid values: 5.5, 5.6, 5.7.
	EngineVersion *string `json:"EngineVersion,omitempty" name:"EngineVersion"`

	// Mode of switch to new instance. Valid values: 0 (switch immediately), 1 (switch within a time window). Default value: 0. If the value is 1, the switch process will be performed within a time window. Or, you can call the [SwitchForUpgrade](https://intl.cloud.tencent.com/document/product/236/15864?from_cn_redirect=1) API to trigger the process.
	WaitSwitch *int64 `json:"WaitSwitch,omitempty" name:"WaitSwitch"`

	// AZ information of secondary database 2, which is empty by default. This parameter can be specified when upgrading primary instances and is meaningless for read-only or disaster recovery instances.
	BackupZone *string `json:"BackupZone,omitempty" name:"BackupZone"`

	// Instance type. Valid values: master (primary instance), dr (disaster recovery instance), ro (read-only instance). Default value: master.
	InstanceRole *string `json:"InstanceRole,omitempty" name:"InstanceRole"`

	// The resource isolation type after the instance is upgraded. Valid values: `UNIVERSAL` (general instance), `EXCLUSIVE` (dedicated instance), `BASIC` (basic instance). If this parameter is left empty, the resource isolation type will be the same as the original one.
	DeviceType *string `json:"DeviceType,omitempty" name:"DeviceType"`

	// The number of CPU cores after the instance is upgraded. If this parameter is left empty, the number of CPU cores will be automatically filled in according to the `Memory` value.
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

	// Whether to enable QuickChange. Valid values: `0` (no), `1` (yes). After QuickChange is enabled, the required resources will be checked. QuickChange is performed only when the required resources support the feature; otherwise, an error message will be returned.
	FastUpgrade *int64 `json:"FastUpgrade,omitempty" name:"FastUpgrade"`

	// Delay threshold. Value range: 1-10. Default value: `10`.
	MaxDelayTime *int64 `json:"MaxDelayTime,omitempty" name:"MaxDelayTime"`
}

func (r *UpgradeDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Memory")
	delete(f, "Volume")
	delete(f, "ProtectMode")
	delete(f, "DeployMode")
	delete(f, "SlaveZone")
	delete(f, "EngineVersion")
	delete(f, "WaitSwitch")
	delete(f, "BackupZone")
	delete(f, "InstanceRole")
	delete(f, "DeviceType")
	delete(f, "Cpu")
	delete(f, "FastUpgrade")
	delete(f, "MaxDelayTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpgradeDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type UpgradeDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Order ID.
		DealIds []*string `json:"DealIds,omitempty" name:"DealIds"`

		// Async task request ID, which can be used to query the execution result of an async task.
		AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpgradeDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UploadInfo struct {

	// Number of parts of file
	AllSliceNum *int64 `json:"AllSliceNum,omitempty" name:"AllSliceNum"`

	// Number of completed parts
	CompleteNum *int64 `json:"CompleteNum,omitempty" name:"CompleteNum"`
}

type ZoneConf struct {

	// AZ deployment mode. Value range: 0 (single-AZ), 1 (multi-AZ)
	DeployMode []*int64 `json:"DeployMode,omitempty" name:"DeployMode"`

	// AZ where the primary instance is located
	MasterZone []*string `json:"MasterZone,omitempty" name:"MasterZone"`

	// AZ where salve database 1 is located when the instance is deployed in multi-AZ mode
	SlaveZone []*string `json:"SlaveZone,omitempty" name:"SlaveZone"`

	// AZ where salve database 2 is located when the instance is deployed in multi-AZ mode
	BackupZone []*string `json:"BackupZone,omitempty" name:"BackupZone"`
}

type ZoneSellConf struct {

	// AZ status. Value range: 0 (not available), 1 (available), 2 (purchasable), 3 (not purchasable), 4 (not displayed)
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// AZ name
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// Whether it is a custom instance type
	IsCustom *bool `json:"IsCustom,omitempty" name:"IsCustom"`

	// Whether disaster recovery is supported
	IsSupportDr *bool `json:"IsSupportDr,omitempty" name:"IsSupportDr"`

	// Whether VPC is supported
	IsSupportVpc *bool `json:"IsSupportVpc,omitempty" name:"IsSupportVpc"`

	// Maximum purchasable quantity of hourly billed instances
	HourInstanceSaleMaxNum *int64 `json:"HourInstanceSaleMaxNum,omitempty" name:"HourInstanceSaleMaxNum"`

	// Whether it is a default AZ
	IsDefaultZone *bool `json:"IsDefaultZone,omitempty" name:"IsDefaultZone"`

	// Whether it is a BM zone
	IsBm *bool `json:"IsBm,omitempty" name:"IsBm"`

	// Supported billing method. Value range: 0 (monthly subscribed), 1 (hourly), 2 (postpaid)
	PayType []*string `json:"PayType,omitempty" name:"PayType"`

	// Data replication type. Value range: 0 (async), 1 (semi-sync), 2 (strong sync)
	ProtectMode []*string `json:"ProtectMode,omitempty" name:"ProtectMode"`

	// AZ name
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// Array of purchasable instance types
	SellType []*SellType `json:"SellType,omitempty" name:"SellType"`

	// Multi-AZ information
	ZoneConf *ZoneConf `json:"ZoneConf,omitempty" name:"ZoneConf"`

	// Information of the supported disaster recovery AZ
	DrZone []*string `json:"DrZone,omitempty" name:"DrZone"`

	// Whether cross-AZ read-only access is supported
	IsSupportRemoteRo *bool `json:"IsSupportRemoteRo,omitempty" name:"IsSupportRemoteRo"`

	// Information of supported cross-AZ read-only zone
	// Note: this field may return null, indicating that no valid values can be obtained.
	RemoteRoZone []*string `json:"RemoteRoZone,omitempty" name:"RemoteRoZone"`
}
