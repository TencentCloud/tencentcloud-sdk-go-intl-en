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

package v20180328

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

type AccountCreateInfo struct {

	// Instance username
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// Instance password
	Password *string `json:"Password,omitempty" name:"Password"`

	// List of database permissions
	DBPrivileges []*DBPrivilege `json:"DBPrivileges,omitempty" name:"DBPrivileges"`

	// Account remarks
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// Whether it is an admin account. Default value: no
	IsAdmin *bool `json:"IsAdmin,omitempty" name:"IsAdmin"`
}

type AccountDetail struct {

	// Account name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Account remarks
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// Account creation time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Account status. 1: creating, 2: normal, 3: modifying, 4: resetting password, -1: deleting
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// Account update time
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// Password update time
	PassTime *string `json:"PassTime,omitempty" name:"PassTime"`

	// Internal account status, which should be `enable` normally
	InternalStatus *string `json:"InternalStatus,omitempty" name:"InternalStatus"`

	// Information of read and write permissions of this account on relevant databases
	Dbs []*DBPrivilege `json:"Dbs,omitempty" name:"Dbs"`

	// Whether it is an admin account
	IsAdmin *bool `json:"IsAdmin,omitempty" name:"IsAdmin"`
}

type AccountPassword struct {

	// Username
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// Password
	Password *string `json:"Password,omitempty" name:"Password"`
}

type AccountPrivilege struct {

	// Database username
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// Database permissions. ReadWrite: read/write, ReadOnly: read-only
	Privilege *string `json:"Privilege,omitempty" name:"Privilege"`
}

type AccountPrivilegeModifyInfo struct {

	// Database username
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// Account permission change information
	DBPrivileges []*DBPrivilegeModifyInfo `json:"DBPrivileges,omitempty" name:"DBPrivileges"`
}

type AccountRemark struct {

	// Account name
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// New remarks of account
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

type Backup struct {

	// Filename
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// File size in KB
	Size *int64 `json:"Size,omitempty" name:"Size"`

	// Backup start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Backup end time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Download address for private network
	InternalAddr *string `json:"InternalAddr,omitempty" name:"InternalAddr"`

	// Download address for public network
	ExternalAddr *string `json:"ExternalAddr,omitempty" name:"ExternalAddr"`

	// Unique ID of backup file, which will be used by the `RestoreInstance` API
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// Backup file status (0: creating, 1: succeeded, 2: failed)
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// List of databases for multi-database backup
	DBs []*string `json:"DBs,omitempty" name:"DBs"`

	// Backup policy (0: instance backup, 1: multi-database backup)
	Strategy *int64 `json:"Strategy,omitempty" name:"Strategy"`

	// Backup mode. 0: scheduled, 1: manual
	BackupWay *int64 `json:"BackupWay,omitempty" name:"BackupWay"`

	// Backup name, which can be customized.
	BackupName *string `json:"BackupName,omitempty" name:"BackupName"`
}

type CloneDBRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of mssql-j8kv137v
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Clone and rename the databases specified in `ReNameRestoreDatabase`. Please note that the clones must be renamed.
	RenameRestore []*RenameRestoreDatabase `json:"RenameRestore,omitempty" name:"RenameRestore"`
}

func (r *CloneDBRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloneDBRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "RenameRestore")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CloneDBRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CloneDBResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Async task request ID, which can be used in the `DescribeFlowStatus` API to query the execution result of an async task
		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CloneDBResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloneDBResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CosUploadBackupFile struct {

	// Backup name
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// Backup size
	Size *int64 `json:"Size,omitempty" name:"Size"`
}

type CreateAccountRequest struct {
	*tchttp.BaseRequest

	// Database instance ID in the format of mssql-njj2mtpl
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Database instance account information
	Accounts []*AccountCreateInfo `json:"Accounts,omitempty" name:"Accounts"`
}

func (r *CreateAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Accounts")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateAccountResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Task flow ID
		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateBackupMigrationRequest struct {
	*tchttp.BaseRequest

	// ID of imported target instance
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Migration task restoration type. FULL: full backup restoration, FULL_LOG: full backup and transaction log restoration, FULL_DIFF: full backup and differential backup restoration
	RecoveryType *string `json:"RecoveryType,omitempty" name:"RecoveryType"`

	// Backup upload type. COS_URL: the backup is stored in user’s Cloud Object Storage, with URL provided. COS_UPLOAD: the backup is stored in the application’s Cloud Object Storage and needs to be uploaded by the user.
	UploadType *string `json:"UploadType,omitempty" name:"UploadType"`

	// Task name
	MigrationName *string `json:"MigrationName,omitempty" name:"MigrationName"`

	// If the UploadType is COS_URL, fill in the URL here. If the UploadType is COS_UPLOAD, fill in the name of the backup file here. Only 1 backup file is supported, but a backup file can involve multiple databases.
	BackupFiles []*string `json:"BackupFiles,omitempty" name:"BackupFiles"`
}

func (r *CreateBackupMigrationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBackupMigrationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "RecoveryType")
	delete(f, "UploadType")
	delete(f, "MigrationName")
	delete(f, "BackupFiles")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBackupMigrationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateBackupMigrationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Backup import task ID
		BackupMigrationId *string `json:"BackupMigrationId,omitempty" name:"BackupMigrationId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateBackupMigrationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBackupMigrationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateBackupRequest struct {
	*tchttp.BaseRequest

	// Backup policy (0: instance backup, 1: multi-database backup)
	Strategy *int64 `json:"Strategy,omitempty" name:"Strategy"`

	// List of names of databases to be backed up (required only for multi-database backup)
	DBNames []*string `json:"DBNames,omitempty" name:"DBNames"`

	// Instance ID in the format of mssql-i1z41iwd
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Backup name. If this parameter is left empty, a backup name in the format of "[Instance ID]_[Backup start timestamp]" will be automatically generated.
	BackupName *string `json:"BackupName,omitempty" name:"BackupName"`
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
	delete(f, "Strategy")
	delete(f, "DBNames")
	delete(f, "InstanceId")
	delete(f, "BackupName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBackupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateBackupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The async job ID
		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

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

type CreateDBInstancesRequest struct {
	*tchttp.BaseRequest

	// Instance AZ, such as ap-guangzhou-1 (Guangzhou Zone 1). Purchasable AZs for an instance can be obtained through the `DescribeZones` API
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// Instance memory size in GB
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// Instance storage capacity in GB
	Storage *int64 `json:"Storage,omitempty" name:"Storage"`

	// Billing mode. Valid value: POSTPAID (pay-as-you-go).
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`

	// Project ID
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Number of instances purchased this time. Default value: 1. Maximum value: 10
	GoodsNum *int64 `json:"GoodsNum,omitempty" name:"GoodsNum"`

	// VPC subnet ID in the format of subnet-bdoe83fa. `SubnetId` and `VpcId` should be set or ignored simultaneously
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// VPC ID in the format of vpc-dsp338hz. `SubnetId` and `VpcId` should be set or ignored simultaneously
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Length of purchase of instance. The default value is 1, indicating one month. The value cannot exceed 48
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// Whether to automatically use voucher. 0: no, 1: yes. Default value: no
	AutoVoucher *int64 `json:"AutoVoucher,omitempty" name:"AutoVoucher"`

	// Array of voucher IDs (currently, only one voucher can be used per order)
	VoucherIds []*string `json:"VoucherIds,omitempty" name:"VoucherIds"`

	// SQL Server version. Valid values: 2008R2 (SQL Server 2008 Enterprise), 2012SP3 (SQL Server 2012 Enterprise), 2016SP1 (SQL Server 2016 Enterprise), 201602 (SQL Server 2016 Standard), 2017 (SQL Server 2017 Enterprise). The version purchasable varies by region and can be queried by calling the `DescribeProductConfig` API. If this parameter is left empty, 2008R2 will be used by default.
	DBVersion *string `json:"DBVersion,omitempty" name:"DBVersion"`

	// Auto-renewal flag. 0: normal renewal, 1: auto-renewal. Default value: 1.
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`

	// Security group list, which contains security group IDs in the format of sg-xxx.
	SecurityGroupList []*string `json:"SecurityGroupList,omitempty" name:"SecurityGroupList"`

	// Configuration of the maintenance window, which specifies the day of the week when maintenance can be performed. Valid values: 1 (Monday), 2 (Tuesday), 3 (Wednesday), 4 (Thursday), 5 (Friday), 6 (Saturday), 7 (Sunday).
	Weekly []*int64 `json:"Weekly,omitempty" name:"Weekly"`

	// Configuration of the maintenance window, which specifies the start time of daily maintenance.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Configuration of the maintenance window, which specifies the maintenance duration in hours.
	Span *int64 `json:"Span,omitempty" name:"Span"`

	// The type of purchased high-availability instance. Valid values: DUAL (dual-server high availability), CLUSTER (cluster). Default value: DUAL.
	HAType *string `json:"HAType,omitempty" name:"HAType"`

	// Whether to deploy across availability zones. Default value: false.
	MultiZones *bool `json:"MultiZones,omitempty" name:"MultiZones"`

	// Tags associated with the instances to be created
	ResourceTags []*ResourceTag `json:"ResourceTags,omitempty" name:"ResourceTags"`
}

func (r *CreateDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDBInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Zone")
	delete(f, "Memory")
	delete(f, "Storage")
	delete(f, "InstanceChargeType")
	delete(f, "ProjectId")
	delete(f, "GoodsNum")
	delete(f, "SubnetId")
	delete(f, "VpcId")
	delete(f, "Period")
	delete(f, "AutoVoucher")
	delete(f, "VoucherIds")
	delete(f, "DBVersion")
	delete(f, "AutoRenewFlag")
	delete(f, "SecurityGroupList")
	delete(f, "Weekly")
	delete(f, "StartTime")
	delete(f, "Span")
	delete(f, "HAType")
	delete(f, "MultiZones")
	delete(f, "ResourceTags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDBInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Order name
		DealName *string `json:"DealName,omitempty" name:"DealName"`

		// Order name array
		DealNames []*string `json:"DealNames,omitempty" name:"DealNames"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDBRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Database creation information
	DBs []*DBCreateInfo `json:"DBs,omitempty" name:"DBs"`
}

func (r *CreateDBRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDBRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "DBs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDBRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateDBResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Task flow ID
		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDBResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDBResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateIncrementalMigrationRequest struct {
	*tchttp.BaseRequest

	// ID of imported target instance
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Backup import task ID, which is returned through the API CreateBackupMigration.
	BackupMigrationId *string `json:"BackupMigrationId,omitempty" name:"BackupMigrationId"`

	// Incremental backup file. If the UploadType of a full backup file is COS_URL, fill in URL here. If the UploadType is COS_UPLOAD, fill in the name of the backup file here. Only 1 backup file is supported, but a backup file can involve multiple databases.
	BackupFiles []*string `json:"BackupFiles,omitempty" name:"BackupFiles"`

	// Whether restoration is required. No: not required. Yes: required. Not required by default.
	IsRecovery *string `json:"IsRecovery,omitempty" name:"IsRecovery"`
}

func (r *CreateIncrementalMigrationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIncrementalMigrationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BackupMigrationId")
	delete(f, "BackupFiles")
	delete(f, "IsRecovery")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateIncrementalMigrationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateIncrementalMigrationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// ID of an incremental backup import task
		IncrementalMigrationId *string `json:"IncrementalMigrationId,omitempty" name:"IncrementalMigrationId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateIncrementalMigrationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIncrementalMigrationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateMigrationRequest struct {
	*tchttp.BaseRequest

	// Migration task name
	MigrateName *string `json:"MigrateName,omitempty" name:"MigrateName"`

	// Migration type (1: structure migration, 2: data migration, 3: incremental sync)
	MigrateType *uint64 `json:"MigrateType,omitempty" name:"MigrateType"`

	// Migration source type. 1: TencentDB for SQL Server, 2: CVM-based self-created SQL Server database; 3: SQL Server backup restoration, 4: SQL Server backup restoration (in COS mode)
	SourceType *uint64 `json:"SourceType,omitempty" name:"SourceType"`

	// Migration source
	Source *MigrateSource `json:"Source,omitempty" name:"Source"`

	// Migration target
	Target *MigrateTarget `json:"Target,omitempty" name:"Target"`

	// Database objects to be migrated. This parameter is not used for offline migration (SourceType=4 or SourceType=5)
	MigrateDBSet []*MigrateDB `json:"MigrateDBSet,omitempty" name:"MigrateDBSet"`

	// Restore and rename the databases listed in `ReNameRestoreDatabase`. If this parameter is left empty, all restored databases will be renamed in the default format. This parameter takes effect only when `SourceType=5`.
	RenameRestore []*RenameRestoreDatabase `json:"RenameRestore,omitempty" name:"RenameRestore"`
}

func (r *CreateMigrationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMigrationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MigrateName")
	delete(f, "MigrateType")
	delete(f, "SourceType")
	delete(f, "Source")
	delete(f, "Target")
	delete(f, "MigrateDBSet")
	delete(f, "RenameRestore")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateMigrationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateMigrationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Migration task ID
		MigrateId *int64 `json:"MigrateId,omitempty" name:"MigrateId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateMigrationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMigrationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DBCreateInfo struct {

	// Database name
	DBName *string `json:"DBName,omitempty" name:"DBName"`

	// Character set, which can be queried by the `DescribeDBCharsets` API. Default value: `Chinese_PRC_CI_AS`.
	Charset *string `json:"Charset,omitempty" name:"Charset"`

	// Database account permission information
	Accounts []*AccountPrivilege `json:"Accounts,omitempty" name:"Accounts"`

	// Remarks
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

type DBDetail struct {

	// Database name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Character set
	Charset *string `json:"Charset,omitempty" name:"Charset"`

	// Remarks
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// Database creation time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Database status. 1: creating, 2: running, 3: modifying, -1: dropping
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// Database account permission information
	Accounts []*AccountPrivilege `json:"Accounts,omitempty" name:"Accounts"`

	// Internal status. ONLINE: running
	InternalStatus *string `json:"InternalStatus,omitempty" name:"InternalStatus"`
}

type DBInstance struct {

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Instance name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Project ID of instance
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Instance region ID
	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`

	// Instance AZ ID
	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// Instance VPC ID, which will be 0 if the basic network is used
	VpcId *int64 `json:"VpcId,omitempty" name:"VpcId"`

	// Instance VPC subnet ID, which will be 0 if the basic network is used
	SubnetId *int64 `json:"SubnetId,omitempty" name:"SubnetId"`

	// Instance status. Valid values: <li>1: applying </li> <li>2: running </li> <li>3: restrictedly running (primary/secondary switching) </li> <li>4: isolated </li> <li>5: repossessing </li> <li>6: repossessed </li> <li>7: task running (e.g., backing up or rolling back the instance) </li> <li>8: decommissioned </li> <li>9: scaling </li> <li>10: migrating </li> <li>11: read-only </li> <li>12: restarting </li>
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// Instance access IP
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// Instance access port
	Vport *int64 `json:"Vport,omitempty" name:"Vport"`

	// Instance creation time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Instance update time
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// Instance billing start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Instance billing end time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Instance isolation time
	IsolateTime *string `json:"IsolateTime,omitempty" name:"IsolateTime"`

	// Instance memory size in GB
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// Used storage capacity of instance in GB
	UsedStorage *int64 `json:"UsedStorage,omitempty" name:"UsedStorage"`

	// Instance storage capacity in GB
	Storage *int64 `json:"Storage,omitempty" name:"Storage"`

	// Instance version
	VersionName *string `json:"VersionName,omitempty" name:"VersionName"`

	// Instance renewal flag
	RenewFlag *int64 `json:"RenewFlag,omitempty" name:"RenewFlag"`

	// High-availability instance type. Valid values: 1 (dual-server high-availability), 2 (standalone), 3 (multi-AZ), 4 (multi-AZ cluster), 5 (cluster), 9 (private consumption)
	Model *int64 `json:"Model,omitempty" name:"Model"`

	// Instance region name, such as ap-guangzhou
	Region *string `json:"Region,omitempty" name:"Region"`

	// Instance AZ name, such as ap-guangzhou-1
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// Backup time point
	BackupTime *string `json:"BackupTime,omitempty" name:"BackupTime"`

	// Instance billing mode. 0: pay-as-you-go
	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`

	// Instance UID
	Uid *string `json:"Uid,omitempty" name:"Uid"`

	// Number of CPU cores of instance
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

	// Instance version code
	Version *string `json:"Version,omitempty" name:"Version"`

	// Physical server code
	Type *string `json:"Type,omitempty" name:"Type"`

	// Billing ID
	Pid *int64 `json:"Pid,omitempty" name:"Pid"`

	// Unique string-type ID of instance VPC in the format of `vpc-xxx`, which is an empty string if the basic network is used
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// Unique string-type ID of instance subnet in the format of `subnet-xxx`, which is an empty string if the basic network is used
	UniqSubnetId *string `json:"UniqSubnetId,omitempty" name:"UniqSubnetId"`

	// Instance isolation.
	// Note: this field may return null, indicating that no valid values can be obtained.
	IsolateOperator *string `json:"IsolateOperator,omitempty" name:"IsolateOperator"`

	// Pub/sub flag. Valid values: SUB (subscribe instance), PUB (publish instance). If it is left empty, it refers to a regular instance without a pub/sub design.
	// Note: this field may return null, indicating that no valid values can be obtained.
	SubFlag *string `json:"SubFlag,omitempty" name:"SubFlag"`

	// Read-only flag. Valid values: RO (read-only instance), MASTER (primary instance with read-only instances). If it is left empty, it refers to an instance which is not read-only and has no RO group.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ROFlag *string `json:"ROFlag,omitempty" name:"ROFlag"`

	// Disaster recovery type. Valid values: MIRROR (image), ALWAYSON (AlwaysOn), SINGLE (singleton).
	// Note: this field may return null, indicating that no valid values can be obtained.
	HAFlag *string `json:"HAFlag,omitempty" name:"HAFlag"`

	// The list of tags associated with the instance
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ResourceTags []*ResourceTag `json:"ResourceTags,omitempty" name:"ResourceTags"`
}

type DBPrivilege struct {

	// Database name
	DBName *string `json:"DBName,omitempty" name:"DBName"`

	// Database permissions. ReadWrite: read/write, ReadOnly: read-only
	Privilege *string `json:"Privilege,omitempty" name:"Privilege"`
}

type DBPrivilegeModifyInfo struct {

	// Database name
	DBName *string `json:"DBName,omitempty" name:"DBName"`

	// Permission change information. ReadWrite: read/write, ReadOnly: read-only, Delete: the account has the permission to delete this database
	Privilege *string `json:"Privilege,omitempty" name:"Privilege"`
}

type DBRemark struct {

	// Database name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Remarks
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

type DbRollbackTimeInfo struct {

	// Database name
	DBName *string `json:"DBName,omitempty" name:"DBName"`

	// Start time of time range available for rollback
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time of time range available for rollback
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type DealInfo struct {

	// Order name
	DealName *string `json:"DealName,omitempty" name:"DealName"`

	// Number of items
	Count *uint64 `json:"Count,omitempty" name:"Count"`

	// ID of associated flow, which can be used to query the flow execution status
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

	// This field is required only for an order that creates an instance, indicating the ID of the instance created by the order
	InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet"`

	// Account
	OwnerUin *string `json:"OwnerUin,omitempty" name:"OwnerUin"`

	// Instance billing type
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`
}

type DeleteAccountRequest struct {
	*tchttp.BaseRequest

	// Database instance ID in the format of mssql-njj2mtpl
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Array of instance usernames
	UserNames []*string `json:"UserNames,omitempty" name:"UserNames"`
}

func (r *DeleteAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "UserNames")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAccountResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Task flow ID
		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteBackupMigrationRequest struct {
	*tchttp.BaseRequest

	// Target instance ID, which is returned through the API DescribeBackupMigration.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Backup import task ID, which is returned through the API DescribeBackupMigration.
	BackupMigrationId *string `json:"BackupMigrationId,omitempty" name:"BackupMigrationId"`
}

func (r *DeleteBackupMigrationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBackupMigrationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BackupMigrationId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteBackupMigrationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteBackupMigrationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteBackupMigrationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBackupMigrationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDBRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of mssql-rljoi3bf
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Array of database names
	Names []*string `json:"Names,omitempty" name:"Names"`
}

func (r *DeleteDBRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDBRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Names")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDBRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDBResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Task flow ID
		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteDBResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDBResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteIncrementalMigrationRequest struct {
	*tchttp.BaseRequest

	// Target instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Backup import task ID
	BackupMigrationId *string `json:"BackupMigrationId,omitempty" name:"BackupMigrationId"`

	// ID of an incremental backup import task
	IncrementalMigrationId *string `json:"IncrementalMigrationId,omitempty" name:"IncrementalMigrationId"`
}

func (r *DeleteIncrementalMigrationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteIncrementalMigrationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BackupMigrationId")
	delete(f, "IncrementalMigrationId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteIncrementalMigrationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteIncrementalMigrationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteIncrementalMigrationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteIncrementalMigrationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMigrationRequest struct {
	*tchttp.BaseRequest

	// Migration task ID
	MigrateId *uint64 `json:"MigrateId,omitempty" name:"MigrateId"`
}

func (r *DeleteMigrationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMigrationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MigrateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteMigrationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMigrationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteMigrationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMigrationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountsRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Number of results per page. Value range: 1-100. Default value: 20
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Page number. Default value: 0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
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
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccountsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Instance ID
		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

		// Account information list
		Accounts []*AccountDetail `json:"Accounts,omitempty" name:"Accounts"`

		// Total number
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

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

type DescribeBackupCommandRequest struct {
	*tchttp.BaseRequest

	// Backup file type. Full: full backup. FULL_LOG: full backup which needs log increments. FULL_DIFF: full backup which needs differential increments. LOG: log backup. DIFF: differential backup.
	BackupFileType *string `json:"BackupFileType,omitempty" name:"BackupFileType"`

	// Database name
	DataBaseName *string `json:"DataBaseName,omitempty" name:"DataBaseName"`

	// Whether restoration is required. No: not required. Yes: required.
	IsRecovery *string `json:"IsRecovery,omitempty" name:"IsRecovery"`

	// Storage path of backup files. If this parameter is left empty, the default storage path will be D:\\.
	LocalPath *string `json:"LocalPath,omitempty" name:"LocalPath"`
}

func (r *DescribeBackupCommandRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupCommandRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BackupFileType")
	delete(f, "DataBaseName")
	delete(f, "IsRecovery")
	delete(f, "LocalPath")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackupCommandRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBackupCommandResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Create a backup command
		Command *string `json:"Command,omitempty" name:"Command"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBackupCommandResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupCommandResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBackupMigrationRequest struct {
	*tchttp.BaseRequest

	// ID of imported target instance
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Backup import task ID, which is returned through the API CreateBackupMigration.
	BackupMigrationId *string `json:"BackupMigrationId,omitempty" name:"BackupMigrationId"`

	// Import task name
	MigrationName *string `json:"MigrationName,omitempty" name:"MigrationName"`

	// Backup file name
	BackupFileName *string `json:"BackupFileName,omitempty" name:"BackupFileName"`

	// Status set of import tasks
	StatusSet []*int64 `json:"StatusSet,omitempty" name:"StatusSet"`

	// Import task restoration type: FULL,FULL_LOG,FULL_DIFF
	RecoveryType *string `json:"RecoveryType,omitempty" name:"RecoveryType"`

	// COS_URL: the backup is stored in user’s Cloud Object Storage, with URL provided. COS_UPLOAD: the backup is stored in the application’s Cloud Object Storage and needs to be uploaded by the user.
	UploadType *string `json:"UploadType,omitempty" name:"UploadType"`

	// Paging. Page size
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Paging. Number of pages
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Field for order: name,createTime,startTime,endTime
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// Type of order: desc,asc
	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`
}

func (r *DescribeBackupMigrationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupMigrationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BackupMigrationId")
	delete(f, "MigrationName")
	delete(f, "BackupFileName")
	delete(f, "StatusSet")
	delete(f, "RecoveryType")
	delete(f, "UploadType")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackupMigrationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBackupMigrationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of tasks
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Migration task set
		BackupMigrationSet []*Migration `json:"BackupMigrationSet,omitempty" name:"BackupMigrationSet"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBackupMigrationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupMigrationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBackupUploadSizeRequest struct {
	*tchttp.BaseRequest

	// ID of imported target instance
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Backup import task ID, which is returned through the API CreateBackupMigration
	BackupMigrationId *string `json:"BackupMigrationId,omitempty" name:"BackupMigrationId"`

	// Incremental import task ID
	IncrementalMigrationId *string `json:"IncrementalMigrationId,omitempty" name:"IncrementalMigrationId"`
}

func (r *DescribeBackupUploadSizeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupUploadSizeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BackupMigrationId")
	delete(f, "IncrementalMigrationId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackupUploadSizeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBackupUploadSizeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Information of uploaded backups
		CosUploadBackupFileSet []*CosUploadBackupFile `json:"CosUploadBackupFileSet,omitempty" name:"CosUploadBackupFileSet"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBackupUploadSizeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupUploadSizeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBackupsRequest struct {
	*tchttp.BaseRequest

	// Start name (yyyy-MM-dd HH:mm:ss)
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time (yyyy-MM-dd HH:mm:ss)
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Instance ID in the format of mssql-njj2mtpl
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Number of results per page. Value range: 1-100. Default value: 20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Page number. Default value: 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Filter by backup name. If this parameter is left empty, backup name will not be used in filtering.
	BackupName *string `json:"BackupName,omitempty" name:"BackupName"`

	// Filter by backup policy. Valid values: 0 (instance backup), 1 (multi-database backup). If this parameter is left empty, backup policy will not be used in filtering.
	Strategy *int64 `json:"Strategy,omitempty" name:"Strategy"`

	// Filter by backup mode. Valid values: 0 (automatic backup on a regular basis), 1 (manual backup performed by the user at any time). If this parameter is left empty, backup mode will not be used in filtering.
	BackupWay *int64 `json:"BackupWay,omitempty" name:"BackupWay"`

	// Filter by backup ID. If this parameter is left empty, backup ID will not be used in filtering.
	BackupId *uint64 `json:"BackupId,omitempty" name:"BackupId"`

	// Filter backups by the database name. If the parameter is left empty, this filter criteria will not take effect.
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`
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
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "InstanceId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "BackupName")
	delete(f, "Strategy")
	delete(f, "BackupWay")
	delete(f, "BackupId")
	delete(f, "DatabaseName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBackupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of backups
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Backup list details
		Backups []*Backup `json:"Backups,omitempty" name:"Backups"`

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

type DescribeDBCharsetsRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of mssql-j8kv137v
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeDBCharsetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBCharsetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBCharsetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBCharsetsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Database character set list
		DatabaseCharsets []*string `json:"DatabaseCharsets,omitempty" name:"DatabaseCharsets"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBCharsetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBCharsetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstancesRequest struct {
	*tchttp.BaseRequest

	// Project ID
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Instance status. Valid values:
	// <li>1: applying</li>
	// <li>2: running</li>
	// <li>3: running restrictedly (primary/secondary switching)</li>
	// <li>4: isolated</li>
	// <li>5: repossessing</li>
	// <li>6: repossessed</li>
	// <li>7: executing task (e.g., backing up or rolling back instance)</li>
	// <li>8: deactivated</li>
	// <li>9: scaling out instance</li>
	// <li>10: migrating instance</li>
	// <li>11: read-only</li>
	// <li>12: restarting</li>
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// Page number. Default value: 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of results per page. Value range: 1-100. Default value: 100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// One or more instance IDs in the format of mssql-si2823jyl
	InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet"`

	// Retrieves billing type. 0: pay-as-you-go
	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`

	// Unique string-type ID of instance VPC in the format of `vpc-xxx`. If an empty string ("") is passed in, filtering will be made by basic network.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Unique string-type ID of instance subnet in the format of `subnet-xxx`. If an empty string ("") is passed in, filtering will be made by basic network.
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// The list of instance private IPs, such as 172.1.0.12
	VipSet []*string `json:"VipSet,omitempty" name:"VipSet"`

	// The list of instance names used for fuzzy match
	InstanceNameSet []*string `json:"InstanceNameSet,omitempty" name:"InstanceNameSet"`

	// The list of instance version numbers, such as 2008R2, 2012SP3
	VersionSet []*string `json:"VersionSet,omitempty" name:"VersionSet"`

	// Instance availability zone, such as ap-guangzhou-2
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// The list of instance tags
	TagKeys []*string `json:"TagKeys,omitempty" name:"TagKeys"`

	// Keyword used for fuzzy match, including instance ID, instance name, and instance private IP
	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`
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
	delete(f, "Status")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "InstanceIdSet")
	delete(f, "PayMode")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "VipSet")
	delete(f, "InstanceNameSet")
	delete(f, "VersionSet")
	delete(f, "Zone")
	delete(f, "TagKeys")
	delete(f, "SearchKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of eligible instances. If the results are returned in multiple pages, this value will be the number of all eligible instances but not the number of instances returned according to the current values of `Limit` and `Offset`
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Instance list
		DBInstances []*DBInstance `json:"DBInstances,omitempty" name:"DBInstances"`

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

type DescribeDBsRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet"`

	// Number of results per page. Value range: 1-100. Default value: 20
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Page number. Default value: 0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeDBsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIdSet")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of databases
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// List of instance databases
		DBInstances []*InstanceDBDetail `json:"DBInstances,omitempty" name:"DBInstances"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFlowStatusRequest struct {
	*tchttp.BaseRequest

	// Flow ID
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`
}

func (r *DescribeFlowStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFlowStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FlowId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFlowStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFlowStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Flow status. 0: succeeded, 1: failed, 2: running
		Status *int64 `json:"Status,omitempty" name:"Status"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFlowStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFlowStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIncrementalMigrationRequest struct {
	*tchttp.BaseRequest

	// Backup import task ID, which is returned through the API CreateBackupMigration
	BackupMigrationId *string `json:"BackupMigrationId,omitempty" name:"BackupMigrationId"`

	// ID of imported target instance
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Backup file name
	BackupFileName *string `json:"BackupFileName,omitempty" name:"BackupFileName"`

	// Status set of import tasks
	StatusSet []*int64 `json:"StatusSet,omitempty" name:"StatusSet"`

	// Paging. Page size
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Paging. Number of pages
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Field for order: name,createTime,startTime,endTime
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// Type of order: desc,asc
	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`

	// ID of an incremental backup import task
	IncrementalMigrationId *string `json:"IncrementalMigrationId,omitempty" name:"IncrementalMigrationId"`
}

func (r *DescribeIncrementalMigrationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIncrementalMigrationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BackupMigrationId")
	delete(f, "InstanceId")
	delete(f, "BackupFileName")
	delete(f, "StatusSet")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	delete(f, "IncrementalMigrationId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIncrementalMigrationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIncrementalMigrationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of import tasks
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Incremental import task set
		IncrementalMigrationSet []*Migration `json:"IncrementalMigrationSet,omitempty" name:"IncrementalMigrationSet"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeIncrementalMigrationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIncrementalMigrationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMigrationDetailRequest struct {
	*tchttp.BaseRequest

	// Migration task ID
	MigrateId *uint64 `json:"MigrateId,omitempty" name:"MigrateId"`
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
	delete(f, "MigrateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMigrationDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMigrationDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Migration task ID
		MigrateId *uint64 `json:"MigrateId,omitempty" name:"MigrateId"`

		// Migration task name
		MigrateName *string `json:"MigrateName,omitempty" name:"MigrateName"`

		// User ID of migration task
		AppId *uint64 `json:"AppId,omitempty" name:"AppId"`

		// Migration task region
		Region *string `json:"Region,omitempty" name:"Region"`

		// Migration source type. 1: TencentDB for SQL Server, 2: CVM-based self-created SQL Server database; 3: SQL Server backup restoration, 4: SQL Server backup restoration (in COS mode)
		SourceType *int64 `json:"SourceType,omitempty" name:"SourceType"`

		// Migration task creation time
		CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

		// Migration task start time
		StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

		// Migration task end time
		EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

		// Migration task status (1: initializing, 4: migrating, 5: migration failed, 6: migration succeeded)
		Status *uint64 `json:"Status,omitempty" name:"Status"`

		// Migration task progress
		Progress *int64 `json:"Progress,omitempty" name:"Progress"`

		// Migration type (1: structure migration, 2: data migration, 3: incremental sync)
		MigrateType *int64 `json:"MigrateType,omitempty" name:"MigrateType"`

		// Migration source
		Source *MigrateSource `json:"Source,omitempty" name:"Source"`

		// Migration target
		Target *MigrateTarget `json:"Target,omitempty" name:"Target"`

		// Database objects to be migrated. This parameter is not used for offline migration (SourceType=4 or SourceType=5)
		MigrateDBSet []*MigrateDB `json:"MigrateDBSet,omitempty" name:"MigrateDBSet"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeMigrationsRequest struct {
	*tchttp.BaseRequest

	// Status set. As long as a migration task is in a status therein, it will be listed
	StatusSet []*int64 `json:"StatusSet,omitempty" name:"StatusSet"`

	// Migration task name (fuzzy match)
	MigrateName *string `json:"MigrateName,omitempty" name:"MigrateName"`

	// Number of results per page. Value range: 1-100. Default value: 100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Page number. Default value: 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// The query results are sorted by keyword. Valid values: name, createTime, startTime, endTime, status
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// Sorting order. Valid values: desc, asc
	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`
}

func (r *DescribeMigrationsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMigrationsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StatusSet")
	delete(f, "MigrateName")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMigrationsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMigrationsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of query results
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// List of query results
		MigrateTaskSet []*MigrateTask `json:"MigrateTaskSet,omitempty" name:"MigrateTaskSet"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMigrationsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMigrationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrdersRequest struct {
	*tchttp.BaseRequest

	// Order array. The order name will be returned upon shipping, which can be used to call the `DescribeOrders` API to query shipment status
	DealNames []*string `json:"DealNames,omitempty" name:"DealNames"`
}

func (r *DescribeOrdersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrdersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DealNames")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOrdersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrdersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Order information array
		Deals []*DealInfo `json:"Deals,omitempty" name:"Deals"`

		// Number of orders returned
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOrdersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrdersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProductConfigRequest struct {
	*tchttp.BaseRequest

	// AZ ID in the format of ap-guangzhou-1
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// The type of instances to be purchased. Valid values: HA (High-Availability Edition, including dual-server high availability and AlwaysOn cluster), RO (read-only replica), SI (Basic Edition)
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
}

func (r *DescribeProductConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProductConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Zone")
	delete(f, "InstanceType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProductConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProductConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Specification information array
		SpecInfoList []*SpecInfo `json:"SpecInfoList,omitempty" name:"SpecInfoList"`

		// Number of date entries returned
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProductConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProductConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

type DescribeRegionsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of regions returned
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Region information array
		RegionSet []*RegionInfo `json:"RegionSet,omitempty" name:"RegionSet"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeRollbackTimeRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// List of databases to be queried
	DBs []*string `json:"DBs,omitempty" name:"DBs"`
}

func (r *DescribeRollbackTimeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRollbackTimeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "DBs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRollbackTimeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRollbackTimeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Information of time range available for database rollback
		Details []*DbRollbackTimeInfo `json:"Details,omitempty" name:"Details"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRollbackTimeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRollbackTimeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSlowlogsRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of mssql-k8voqdlz
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Query start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Query end time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Number of results per page. Value range: 1-100. Default value: 20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Page number. Default value: 0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeSlowlogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlowlogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSlowlogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSlowlogsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of queries
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Information list of slow query logs
		Slowlogs []*SlowlogInfo `json:"Slowlogs,omitempty" name:"Slowlogs"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSlowlogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlowlogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUploadBackupInfoRequest struct {
	*tchttp.BaseRequest

	// ID of imported target instance
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Backup import task ID, which is returned through the API CreateBackupMigration
	BackupMigrationId *string `json:"BackupMigrationId,omitempty" name:"BackupMigrationId"`
}

func (r *DescribeUploadBackupInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUploadBackupInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BackupMigrationId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUploadBackupInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUploadBackupInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Bucket name
		BucketName *string `json:"BucketName,omitempty" name:"BucketName"`

		// Bucket location information
		Region *string `json:"Region,omitempty" name:"Region"`

		// Storage path
		Path *string `json:"Path,omitempty" name:"Path"`

		// Temporary key ID
		TmpSecretId *string `json:"TmpSecretId,omitempty" name:"TmpSecretId"`

		// Temporary key (Key)
		TmpSecretKey *string `json:"TmpSecretKey,omitempty" name:"TmpSecretKey"`

		// Temporary key (Token)
		XCosSecurityToken *string `json:"XCosSecurityToken,omitempty" name:"XCosSecurityToken"`

		// Temporary key start time
		StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

		// Temporary key expiration time
		ExpiredTime *string `json:"ExpiredTime,omitempty" name:"ExpiredTime"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUploadBackupInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUploadBackupInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeZonesRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeZonesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeZonesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeZonesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeZonesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of AZs returned
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Array of AZs
		ZoneSet []*ZoneInfo `json:"ZoneSet,omitempty" name:"ZoneSet"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeZonesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeZonesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceCreateDBInstancesRequest struct {
	*tchttp.BaseRequest

	// AZ ID, which can be obtained through the `Zone` field in the returned value of the `DescribeZones` API
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// Memory size in GB
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// Instance capacity in GB
	Storage *int64 `json:"Storage,omitempty" name:"Storage"`

	// Billing type. Valid value: POSTPAID.
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`

	// Length of purchase in months. Value range: 1-48. Default value: 1
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// Number of instances purchased at a time. Value range: 1-100. Default value: 1
	GoodsNum *int64 `json:"GoodsNum,omitempty" name:"GoodsNum"`

	// SQL Server version. Valid values: 2008R2 (SQL Server 2008 Enterprise), 2012SP3 (SQL Server 2012 Enterprise), 2016SP1 (SQL Server 2016 Enterprise), 201602 (SQL Server 2016 Standard), 2017 (SQL Server 2017 Enterprise). Default value: 2008R2.
	DBVersion *string `json:"DBVersion,omitempty" name:"DBVersion"`

	// The number of CPU cores of the instance you want to purchase.
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

	// The type of purchased instance. Valid values: HA (high-availability edition, including dual-server high availability and AlwaysOn cluster), RO (read-only replica), SI (basic edition). Default value: HA.
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// The host type of purchased instance. Valid values: PM (physical machine), CLOUD_PREMIUM (physical machine with premium cloud disk), CLOUD_SSD (physical machine with SSD). Default value: PM.
	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`
}

func (r *InquiryPriceCreateDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceCreateDBInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Zone")
	delete(f, "Memory")
	delete(f, "Storage")
	delete(f, "InstanceChargeType")
	delete(f, "Period")
	delete(f, "GoodsNum")
	delete(f, "DBVersion")
	delete(f, "Cpu")
	delete(f, "InstanceType")
	delete(f, "MachineType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquiryPriceCreateDBInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceCreateDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Price before discount. This value divided by 100 indicates the price; for example, 10010 means 100.10 USD
		OriginalPrice *int64 `json:"OriginalPrice,omitempty" name:"OriginalPrice"`

		// The actual price to be paid. This value divided by 100 indicates the price; for example, 10010 means 100.10 USD
		Price *int64 `json:"Price,omitempty" name:"Price"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceCreateDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceCreateDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceUpgradeDBInstanceRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of mssql-njj2mtpl
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Memory size after instance upgrade in GB, which cannot be smaller than the current instance memory size
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// Storage capacity after instance upgrade in GB, which cannot be smaller than the current instance storage capacity
	Storage *int64 `json:"Storage,omitempty" name:"Storage"`

	// The number of CUP cores after the instance is upgraded, which cannot be smaller than that of the current cores.
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`
}

func (r *InquiryPriceUpgradeDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceUpgradeDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Memory")
	delete(f, "Storage")
	delete(f, "Cpu")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquiryPriceUpgradeDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceUpgradeDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Price before discount. This value divided by 100 indicates the price; for example, 10094 means 100.94 USD
		OriginalPrice *int64 `json:"OriginalPrice,omitempty" name:"OriginalPrice"`

		// The actual price to be paid. This value divided by 100 indicates the price; for example, 10094 means 100.94 USD
		Price *int64 `json:"Price,omitempty" name:"Price"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceUpgradeDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceUpgradeDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstanceDBDetail struct {

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Database information list
	DBDetails []*DBDetail `json:"DBDetails,omitempty" name:"DBDetails"`
}

type MigrateDB struct {

	// Name of migrated database
	DBName *string `json:"DBName,omitempty" name:"DBName"`
}

type MigrateDetail struct {

	// Name of current step
	StepName *string `json:"StepName,omitempty" name:"StepName"`

	// Progress of current step in %
	Progress *int64 `json:"Progress,omitempty" name:"Progress"`
}

type MigrateSource struct {

	// Source instance ID in the format of `mssql-si2823jyl`, which is used when `MigrateType` is 1 (TencentDB for SQL Server)
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// ID of source CVM instance, which is used when `MigrateType` is 2 (CVM-based self-created SQL Server database)
	CvmId *string `json:"CvmId,omitempty" name:"CvmId"`

	// VPC ID of source CVM instance in the format of vpc-6ys9ont9, which is used when `MigrateType` is 2 (CVM-based self-created SQL Server database)
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// VPC subnet ID of source CVM instance in the format of subnet-h9extioi, which is used when `MigrateType` is 2 (CVM-based self-created SQL Server database)
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// Username, which is used when `MigrateType` is 1 or 2
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// Password, which is used when `MigrateType` is 1 or 2
	Password *string `json:"Password,omitempty" name:"Password"`

	// Private IP of source CVM database, which is used when `MigrateType` is 2 (CVM-based self-created SQL Server database)
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// Port number of source CVM database, which is used when `MigrateType` is 2 (CVM-based self-created SQL Server database)
	Port *uint64 `json:"Port,omitempty" name:"Port"`

	// Source backup address for offline migration, which is used when `MigrateType` is 4 or 5
	Url []*string `json:"Url,omitempty" name:"Url"`

	// Source backup password for offline migration, which is used when `MigrateType` is 4 or 5
	UrlPassword *string `json:"UrlPassword,omitempty" name:"UrlPassword"`
}

type MigrateTarget struct {

	// ID of target instance in the format of mssql-si2823jyl
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Username of migration target instance
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// Password of migration target instance
	Password *string `json:"Password,omitempty" name:"Password"`
}

type MigrateTask struct {

	// Migration task ID
	MigrateId *uint64 `json:"MigrateId,omitempty" name:"MigrateId"`

	// Migration task name
	MigrateName *string `json:"MigrateName,omitempty" name:"MigrateName"`

	// User ID of migration task
	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`

	// Migration task region
	Region *string `json:"Region,omitempty" name:"Region"`

	// Migration source type. 1: TencentDB for SQL Server, 2: CVM-based self-created SQL Server database; 3: SQL Server backup restoration, 4: SQL Server backup restoration (in COS mode)
	SourceType *int64 `json:"SourceType,omitempty" name:"SourceType"`

	// Migration task creation time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Migration task start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Migration task end time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Migration task status (1: initializing, 4: migrating, 5: migration failed, 6: migration succeeded, 7: suspended, 8: deleted, 9: suspending, 10: completing, 11: suspension failed, 12: completion failed)
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// Information
	Message *string `json:"Message,omitempty" name:"Message"`

	// Whether migration task has been checked (0: not checked, 1: check succeeded, 2: check failed, 3: checking)
	CheckFlag *uint64 `json:"CheckFlag,omitempty" name:"CheckFlag"`

	// Migration task progress in %
	Progress *int64 `json:"Progress,omitempty" name:"Progress"`

	// Migration task progress details
	MigrateDetail *MigrateDetail `json:"MigrateDetail,omitempty" name:"MigrateDetail"`
}

type Migration struct {

	// Backup import task ID or incremental import task ID
	MigrationId *string `json:"MigrationId,omitempty" name:"MigrationId"`

	// Backup import task name. For an incremental import task, this field will be left empty.
	// Note: this field may return ‘null’, indicating that no valid values can be obtained.
	MigrationName *string `json:"MigrationName,omitempty" name:"MigrationName"`

	// Application ID
	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`

	// Region
	Region *string `json:"Region,omitempty" name:"Region"`

	// ID of migrated target instance
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Migration task restoration type
	RecoveryType *string `json:"RecoveryType,omitempty" name:"RecoveryType"`

	// Backup user upload type. COS_URL: the backup is stored in user’s Cloud Object Storage, with URL provided. COS_UPLOAD: the backup is stored in the application’s Cloud Object Storage and needs to be uploaded by the user.
	UploadType *string `json:"UploadType,omitempty" name:"UploadType"`

	// Backup file list, which is determined by UploadType. If the upload type is COS_URL, URL will be saved. If the upload type is COS_UPLOAD, the backup name will be saved.
	BackupFiles []*string `json:"BackupFiles,omitempty" name:"BackupFiles"`

	// Migration task status. Valid values: `2` (Creation completed), `7` (Importing full backups), `8` (Waiting for incremental backups), `9` (Import success), `10` (Import failure), `12` (Importing incremental backups).
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// Migration task creation time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Migration task start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Migration task end time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// More information
	Message *string `json:"Message,omitempty" name:"Message"`

	// Migration detail
	Detail *MigrationDetail `json:"Detail,omitempty" name:"Detail"`

	// Operation allowed in the current status
	Action *MigrationAction `json:"Action,omitempty" name:"Action"`

	// Whether this is the final restoration. For a full import task, this field will be left empty.
	// Note: this field may return ‘null’, indicating that no valid values can be obtained.
	IsRecovery *string `json:"IsRecovery,omitempty" name:"IsRecovery"`
}

type MigrationAction struct {

	// All the allowed operations. Values include: view (viewing a task), modify (modifying a task), start (starting a task), incremental (creating an incremental task), delete (deleting a task), and upload (obtaining the upload permission).
	AllAction []*string `json:"AllAction,omitempty" name:"AllAction"`

	// Operation allowed in the current status. If the subset of AllAction is left empty, no operations will be allowed.
	AllowedAction []*string `json:"AllowedAction,omitempty" name:"AllowedAction"`
}

type MigrationDetail struct {

	// Total number of steps
	StepAll *int64 `json:"StepAll,omitempty" name:"StepAll"`

	// Current step
	StepNow *int64 `json:"StepNow,omitempty" name:"StepNow"`

	// Overall progress. For example, “30” means 30%.
	Progress *int64 `json:"Progress,omitempty" name:"Progress"`

	// Step information. ‘null’ means the migration has not started
	// Note: this field may return ‘null’, indicating that no valid values can be obtained.
	StepInfo []*MigrationStep `json:"StepInfo,omitempty" name:"StepInfo"`
}

type MigrationStep struct {

	// Step sequence
	StepNo *int64 `json:"StepNo,omitempty" name:"StepNo"`

	// Step name
	StepName *string `json:"StepName,omitempty" name:"StepName"`

	// Step ID in English
	StepId *string `json:"StepId,omitempty" name:"StepId"`

	// Step status: 0 (default value), 1 (succeeded), 2 (failed), 3 (in progress), 4 (not started)
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type ModifyAccountPrivilegeRequest struct {
	*tchttp.BaseRequest

	// Database instance ID in the format of mssql-njj2mtpl
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Account permission change information
	Accounts []*AccountPrivilegeModifyInfo `json:"Accounts,omitempty" name:"Accounts"`
}

func (r *ModifyAccountPrivilegeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccountPrivilegeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Accounts")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAccountPrivilegeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAccountPrivilegeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Async task flow ID
		FlowId *uint64 `json:"FlowId,omitempty" name:"FlowId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAccountPrivilegeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccountPrivilegeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAccountRemarkRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of mssql-j8kv137v
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Information of account for which to modify remarks
	Accounts []*AccountRemark `json:"Accounts,omitempty" name:"Accounts"`
}

func (r *ModifyAccountRemarkRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccountRemarkRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Accounts")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAccountRemarkRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAccountRemarkResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAccountRemarkResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccountRemarkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBackupMigrationRequest struct {
	*tchttp.BaseRequest

	// ID of imported target instance
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Backup import task ID, which is returned through the API CreateBackupMigration
	BackupMigrationId *string `json:"BackupMigrationId,omitempty" name:"BackupMigrationId"`

	// Task name
	MigrationName *string `json:"MigrationName,omitempty" name:"MigrationName"`

	// Migration task restoration type: FULL,FULL_LOG,FULL_DIFF
	RecoveryType *string `json:"RecoveryType,omitempty" name:"RecoveryType"`

	// COS_URL: the backup is stored in user’s Cloud Object Storage, with URL provided. COS_UPLOAD: the backup is stored in the application’s Cloud Object Storage and needs to be uploaded by the user.
	UploadType *string `json:"UploadType,omitempty" name:"UploadType"`

	// If the UploadType is COS_URL, fill in URL here. If the UploadType is COS_UPLOAD, fill in the name of the backup file here. Only 1 backup file is supported, but a backup file can involve multiple databases.
	BackupFiles []*string `json:"BackupFiles,omitempty" name:"BackupFiles"`
}

func (r *ModifyBackupMigrationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBackupMigrationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BackupMigrationId")
	delete(f, "MigrationName")
	delete(f, "RecoveryType")
	delete(f, "UploadType")
	delete(f, "BackupFiles")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyBackupMigrationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBackupMigrationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Backup import task ID
		BackupMigrationId *string `json:"BackupMigrationId,omitempty" name:"BackupMigrationId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyBackupMigrationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBackupMigrationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBackupStrategyRequest struct {
	*tchttp.BaseRequest

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Backup mode, which supports daily backup only. Valid value: daily.
	BackupType *string `json:"BackupType,omitempty" name:"BackupType"`

	// Backup time. Value range: an integer from 0 to 23.
	BackupTime *uint64 `json:"BackupTime,omitempty" name:"BackupTime"`

	// Backup interval in days when the `BackupType` is `daily`. Valid value: 1.
	BackupDay *uint64 `json:"BackupDay,omitempty" name:"BackupDay"`
}

func (r *ModifyBackupStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBackupStrategyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BackupType")
	delete(f, "BackupTime")
	delete(f, "BackupDay")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyBackupStrategyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBackupStrategyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Returned error code.
		Errno *int64 `json:"Errno,omitempty" name:"Errno"`

		// Returned error message.
		Msg *string `json:"Msg,omitempty" name:"Msg"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyBackupStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBackupStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDBInstanceNameRequest struct {
	*tchttp.BaseRequest

	// Database instance ID in the format of mssql-njj2mtpl
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// New name of database instance
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

type ModifyDBInstanceNetworkRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// ID of the new VPC
	NewVpcId *string `json:"NewVpcId,omitempty" name:"NewVpcId"`

	// ID of the new subnet
	NewSubnetId *string `json:"NewSubnetId,omitempty" name:"NewSubnetId"`

	// Retention period (in hours) of the original VIP. Value range: `0-168`. Default value: `0`, indicating the original VIP is released immediately.
	OldIpRetainTime *int64 `json:"OldIpRetainTime,omitempty" name:"OldIpRetainTime"`
}

func (r *ModifyDBInstanceNetworkRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceNetworkRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "NewVpcId")
	delete(f, "NewSubnetId")
	delete(f, "OldIpRetainTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDBInstanceNetworkRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDBInstanceNetworkResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// ID of the instance network changing task. You can use the [DescribeFlowStatus](https://intl.cloud.tencent.com/document/product/238/19967?from_cn_redirect=1) API to query the task status.
		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDBInstanceNetworkResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceNetworkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDBInstanceProjectRequest struct {
	*tchttp.BaseRequest

	// Array of instance IDs in the format of mssql-j8kv137v
	InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet"`

	// Project ID. If this parameter is 0, the default project will be used
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`
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
	delete(f, "InstanceIdSet")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDBInstanceProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDBInstanceProjectResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of successfully modified instances
		Count *int64 `json:"Count,omitempty" name:"Count"`

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

type ModifyDBNameRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Old database name
	OldDBName *string `json:"OldDBName,omitempty" name:"OldDBName"`

	// New name of database
	NewDBName *string `json:"NewDBName,omitempty" name:"NewDBName"`
}

func (r *ModifyDBNameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBNameRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "OldDBName")
	delete(f, "NewDBName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDBNameRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDBNameResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Task flow ID
		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDBNameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDBRemarkRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of mssql-rljoi3bf
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Array of database names and remarks, where each element contains a database name and the corresponding remarks
	DBRemarks []*DBRemark `json:"DBRemarks,omitempty" name:"DBRemarks"`
}

func (r *ModifyDBRemarkRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBRemarkRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "DBRemarks")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDBRemarkRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDBRemarkResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDBRemarkResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBRemarkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyIncrementalMigrationRequest struct {
	*tchttp.BaseRequest

	// ID of imported target instance
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Backup import task ID, which is returned through the API CreateBackupMigration
	BackupMigrationId *string `json:"BackupMigrationId,omitempty" name:"BackupMigrationId"`

	// Incremental import task ID
	IncrementalMigrationId *string `json:"IncrementalMigrationId,omitempty" name:"IncrementalMigrationId"`

	// Whether restoration is required. No: not required. Yes: required.
	IsRecovery *string `json:"IsRecovery,omitempty" name:"IsRecovery"`

	// If the UploadType is COS_URL, fill in URL here. If the UploadType is COS_UPLOAD, fill in the name of the backup file here. Only 1 backup file is supported, but a backup file can involve multiple databases.
	BackupFiles []*string `json:"BackupFiles,omitempty" name:"BackupFiles"`
}

func (r *ModifyIncrementalMigrationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyIncrementalMigrationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BackupMigrationId")
	delete(f, "IncrementalMigrationId")
	delete(f, "IsRecovery")
	delete(f, "BackupFiles")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyIncrementalMigrationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyIncrementalMigrationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// ID of an incremental backup import task
		IncrementalMigrationId *string `json:"IncrementalMigrationId,omitempty" name:"IncrementalMigrationId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyIncrementalMigrationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyIncrementalMigrationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyMigrationRequest struct {
	*tchttp.BaseRequest

	// Migration task ID
	MigrateId *uint64 `json:"MigrateId,omitempty" name:"MigrateId"`

	// New name of migration task. If this parameter is left empty, no modification will be made
	MigrateName *string `json:"MigrateName,omitempty" name:"MigrateName"`

	// New migration type (1: structure migration, 2: data migration, 3: incremental sync). If this parameter is left empty, no modification will be made
	MigrateType *uint64 `json:"MigrateType,omitempty" name:"MigrateType"`

	// Migration source type. 1: TencentDB for SQL Server, 2: CVM-based self-created SQL Server database; 3: SQL Server backup restoration, 4: SQL Server backup restoration (in COS mode). If this parameter is left empty, no modification will be made
	SourceType *uint64 `json:"SourceType,omitempty" name:"SourceType"`

	// Migration source. If this parameter is left empty, no modification will be made
	Source *MigrateSource `json:"Source,omitempty" name:"Source"`

	// Migration target. If this parameter is left empty, no modification will be made
	Target *MigrateTarget `json:"Target,omitempty" name:"Target"`

	// Database objects to be migrated. This parameter is not used for offline migration (SourceType=4 or SourceType=5). If it left empty, no modification will be made
	MigrateDBSet []*MigrateDB `json:"MigrateDBSet,omitempty" name:"MigrateDBSet"`
}

func (r *ModifyMigrationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMigrationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MigrateId")
	delete(f, "MigrateName")
	delete(f, "MigrateType")
	delete(f, "SourceType")
	delete(f, "Source")
	delete(f, "Target")
	delete(f, "MigrateDBSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyMigrationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyMigrationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Migration task ID
		MigrateId *uint64 `json:"MigrateId,omitempty" name:"MigrateId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyMigrationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMigrationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RecycleDBInstanceRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *RecycleDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecycleDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RecycleDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type RecycleDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Task ID
		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RecycleDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecycleDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RegionInfo struct {

	// Region ID in the format of ap-guangzhou
	Region *string `json:"Region,omitempty" name:"Region"`

	// Region name
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`

	// Numeric ID of region
	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`

	// Current purchasability of this region. UNAVAILABLE: not purchasable, AVAILABLE: purchasable
	RegionState *string `json:"RegionState,omitempty" name:"RegionState"`
}

type RenameRestoreDatabase struct {

	// Database name. If the `OldName` database does not exist, a failure will be returned.
	// It can be left empty in offline migration tasks.
	OldName *string `json:"OldName,omitempty" name:"OldName"`

	// New database name. In offline migration, `OldName` will be used if `NewName` is left empty (`OldName` and `NewName` cannot be both empty). In database cloning, `OldName` and `NewName` must be both specified and cannot have the same value.
	NewName *string `json:"NewName,omitempty" name:"NewName"`
}

type ResetAccountPasswordRequest struct {
	*tchttp.BaseRequest

	// Database instance ID in the format of mssql-njj2mtpl
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Updated account password information array
	Accounts []*AccountPassword `json:"Accounts,omitempty" name:"Accounts"`
}

func (r *ResetAccountPasswordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetAccountPasswordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Accounts")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetAccountPasswordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ResetAccountPasswordResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// ID of async task flow for account password change
		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResetAccountPasswordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetAccountPasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResourceTag struct {

	// Tag key
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// Tag value
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type RestartDBInstanceRequest struct {
	*tchttp.BaseRequest

	// Database instance ID in the format of mssql-njj2mtpl
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *RestartDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestartDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RestartDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type RestartDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Async task flow ID
		FlowId *uint64 `json:"FlowId,omitempty" name:"FlowId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RestartDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestartDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RestoreInstanceRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of mssql-j8kv137v
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Backup file ID, which can be obtained through the `Id` field in the returned value of the `DescribeBackups` API
	BackupId *int64 `json:"BackupId,omitempty" name:"BackupId"`

	// ID of the target instance to which the backup is restored. The target instance should be under the same `APPID`. If this parameter is left empty, ID of the source instance will be used.
	TargetInstanceId *string `json:"TargetInstanceId,omitempty" name:"TargetInstanceId"`

	// Restore the databases listed in `ReNameRestoreDatabase` and rename them after restoration. If this parameter is left empty, all databases will be restored and renamed in the default format.
	RenameRestore []*RenameRestoreDatabase `json:"RenameRestore,omitempty" name:"RenameRestore"`
}

func (r *RestoreInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestoreInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BackupId")
	delete(f, "TargetInstanceId")
	delete(f, "RenameRestore")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RestoreInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type RestoreInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Async flow task ID, which can be used to call the `DescribeFlowStatus` API to get the task execution status
		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RestoreInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestoreInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RollbackInstanceRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Rollback type. 0: the database rolled back overwrites the original database; 1: the database rolled back is renamed and does not overwrite the original database
	Type *uint64 `json:"Type,omitempty" name:"Type"`

	// Database to be rolled back
	DBs []*string `json:"DBs,omitempty" name:"DBs"`

	// Target time point for rollback
	Time *string `json:"Time,omitempty" name:"Time"`

	// ID of the target instance to which the backup is restored. The target instance should be under the same `APPID`. If this parameter is left empty, ID of the source instance will be used.
	TargetInstanceId *string `json:"TargetInstanceId,omitempty" name:"TargetInstanceId"`

	// Rename the databases listed in `ReNameRestoreDatabase`. This parameter takes effect only when `Type = 1` which indicates that backup rollback supports renaming databases. If it is left empty, databases will be renamed in the default format and the `DBs` parameter specifies the databases to be restored.
	RenameRestore []*RenameRestoreDatabase `json:"RenameRestore,omitempty" name:"RenameRestore"`
}

func (r *RollbackInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RollbackInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Type")
	delete(f, "DBs")
	delete(f, "Time")
	delete(f, "TargetInstanceId")
	delete(f, "RenameRestore")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RollbackInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type RollbackInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The async job ID
		FlowId *uint64 `json:"FlowId,omitempty" name:"FlowId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RollbackInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RollbackInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RunMigrationRequest struct {
	*tchttp.BaseRequest

	// Migration task ID
	MigrateId *uint64 `json:"MigrateId,omitempty" name:"MigrateId"`
}

func (r *RunMigrationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunMigrationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MigrateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RunMigrationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type RunMigrationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// After the migration task starts, the flow ID will be returned
		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RunMigrationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunMigrationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SlowlogInfo struct {

	// Unique ID of slow query log file
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// File generation start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// File generation end time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// File size in KB
	Size *int64 `json:"Size,omitempty" name:"Size"`

	// Number of logs in file
	Count *int64 `json:"Count,omitempty" name:"Count"`

	// Download address for private network
	InternalAddr *string `json:"InternalAddr,omitempty" name:"InternalAddr"`

	// Download address for public network
	ExternalAddr *string `json:"ExternalAddr,omitempty" name:"ExternalAddr"`

	// Status (1: success, 2: failure)
	// Note: this field may return null, indicating that no valid values can be obtained.
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type SpecInfo struct {

	// Instance specification ID. The `SpecId` returned by `DescribeZones` together with the purchasable specification information returned by `DescribeProductConfig` can be used to find out what specifications can be purchased in a specified AZ
	SpecId *int64 `json:"SpecId,omitempty" name:"SpecId"`

	// Model ID
	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`

	// Model name
	MachineTypeName *string `json:"MachineTypeName,omitempty" name:"MachineTypeName"`

	// Database version information. Valid values: 2008R2 (SQL Server 2008 Enterprise), 2012SP3 (SQL Server 2012 Enterprise), 2016SP1 (SQL Server 2016 Enterprise), 201602 (SQL Server 2016 Standard), 2017 (SQL Server 2017 Enterprise)
	Version *string `json:"Version,omitempty" name:"Version"`

	// Version name corresponding to the `Version` field
	VersionName *string `json:"VersionName,omitempty" name:"VersionName"`

	// Memory size in GB
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// Number of CPU cores
	CPU *int64 `json:"CPU,omitempty" name:"CPU"`

	// Minimum disk size under this specification in GB
	MinStorage *int64 `json:"MinStorage,omitempty" name:"MinStorage"`

	// Maximum disk size under this specification in GB
	MaxStorage *int64 `json:"MaxStorage,omitempty" name:"MaxStorage"`

	// QPS of this specification
	QPS *int64 `json:"QPS,omitempty" name:"QPS"`

	// Description of this specification
	SuitInfo *string `json:"SuitInfo,omitempty" name:"SuitInfo"`

	// Pid of this specification
	Pid *int64 `json:"Pid,omitempty" name:"Pid"`

	// Pay-as-you-go Pid list corresponding to this specification
	// Note: this field may return null, indicating that no valid values can be obtained.
	PostPid []*int64 `json:"PostPid,omitempty" name:"PostPid"`

	// Billing mode under this specification. POST: pay-as-you-go
	PayModeStatus *string `json:"PayModeStatus,omitempty" name:"PayModeStatus"`

	// Instance type. Valid values: HA (High-Availability Edition, including dual-server high availability and AlwaysOn cluster), RO (read-only replica), SI (Basic Edition)
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// Whether multi-AZ deployment is supported. Valid values: MultiZones (only multi-AZ deployment is supported), SameZones (only single-AZ deployment is supported), ALL (both deployments are supported)
	MultiZonesStatus *string `json:"MultiZonesStatus,omitempty" name:"MultiZonesStatus"`
}

type StartBackupMigrationRequest struct {
	*tchttp.BaseRequest

	// ID of imported target instance
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Backup import task ID, which is returned through the API CreateBackupMigration
	BackupMigrationId *string `json:"BackupMigrationId,omitempty" name:"BackupMigrationId"`
}

func (r *StartBackupMigrationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartBackupMigrationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BackupMigrationId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartBackupMigrationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type StartBackupMigrationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Task ID
		FlowId *uint64 `json:"FlowId,omitempty" name:"FlowId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StartBackupMigrationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartBackupMigrationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StartIncrementalMigrationRequest struct {
	*tchttp.BaseRequest

	// ID of imported target instance
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Backup import task ID, which is returned through the API CreateBackupMigration
	BackupMigrationId *string `json:"BackupMigrationId,omitempty" name:"BackupMigrationId"`

	// ID of an incremental backup import task
	IncrementalMigrationId *string `json:"IncrementalMigrationId,omitempty" name:"IncrementalMigrationId"`
}

func (r *StartIncrementalMigrationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartIncrementalMigrationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BackupMigrationId")
	delete(f, "IncrementalMigrationId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartIncrementalMigrationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type StartIncrementalMigrationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Task ID
		FlowId *uint64 `json:"FlowId,omitempty" name:"FlowId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StartIncrementalMigrationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartIncrementalMigrationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TerminateDBInstanceRequest struct {
	*tchttp.BaseRequest

	// List of instance IDs manually terminated in the format of [mssql-3l3fgqn7], which are the same as the instance IDs displayed on the TencentDB Console page
	InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet"`
}

func (r *TerminateDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TerminateDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIdSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TerminateDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type TerminateDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TerminateDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TerminateDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpgradeDBInstanceRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of mssql-j8kv137v
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Memory size after instance upgrade in GB, which cannot be smaller than the current instance memory size
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// Storage capacity after instance upgrade in GB, which cannot be smaller than the current instance storage capacity
	Storage *int64 `json:"Storage,omitempty" name:"Storage"`

	// Whether to automatically use vouchers. 0: no, 1: yes. Default value: 0
	AutoVoucher *int64 `json:"AutoVoucher,omitempty" name:"AutoVoucher"`

	// Voucher ID (currently, only one voucher can be used per order)
	VoucherIds []*string `json:"VoucherIds,omitempty" name:"VoucherIds"`

	// The number of CUP cores after the instance is upgraded.
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

	// Upgrade the SQL Server version. Supported versions include SQL Server 2008 Enterprise (`2008R2`), SQL Server 2012 Enterprise (`2012SP3`), etc. As the purchasable versions are region-specific, you can use the `DescribeProductConfig` API to query the information of purchasable versions in each region. Downgrading is unsupported. If this parameter is left empty, the SQL Server version will not be changed.
	DBVersion *string `json:"DBVersion,omitempty" name:"DBVersion"`

	// Upgrade the high availability architecture from image-based disaster recovery to Always On cluster disaster recovery. This parameter is valid only for instances which support Always On high availability and run SQL Server 2017 or later. Neither downgrading to image-based disaster recovery nor upgrading from cluster disaster recovery to Always On disaster recovery is supported. If this parameter is left empty, the high availability architecture will not be changed.
	HAType *string `json:"HAType,omitempty" name:"HAType"`

	// Change the instance deployment scheme. Valid values: `SameZones` (change to single-AZ deployment, which does not support cross-AZ disaster recovery), `MultiZones` (change to multi-AZ deployment, which supports cross-AZ disaster recovery).
	MultiZones *string `json:"MultiZones,omitempty" name:"MultiZones"`
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
	delete(f, "Storage")
	delete(f, "AutoVoucher")
	delete(f, "VoucherIds")
	delete(f, "Cpu")
	delete(f, "DBVersion")
	delete(f, "HAType")
	delete(f, "MultiZones")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpgradeDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type UpgradeDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Order name
		DealName *string `json:"DealName,omitempty" name:"DealName"`

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

type ZoneInfo struct {

	// AZ ID in the format of ap-guangzhou-1 (i.e., Guangzhou Zone 1)
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// AZ name
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// Numeric ID of AZ
	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// ID of specification purchasable in this AZ, which, together with the returned value of the `DescribeProductConfig` API, can be used to find out the specifications currently purchasable in the AZ
	SpecId *int64 `json:"SpecId,omitempty" name:"SpecId"`

	// Information of database versions purchasable under the current AZ and specification. Valid values: 2008R2 (SQL Server 2008 Enterprise), 2012SP3 (SQL Server 2012 Enterprise), 2016SP1 (SQL Server 2016 Enterprise), 201602 (SQL Server 2016 Standard), 2017 (SQL Server 2017 Enterprise)
	Version *string `json:"Version,omitempty" name:"Version"`
}
