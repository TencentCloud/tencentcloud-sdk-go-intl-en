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
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

type AccountCreateInfo struct {
	// Instance username
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// Instance password
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// List of database permissions
	DBPrivileges []*DBPrivilege `json:"DBPrivileges,omitnil,omitempty" name:"DBPrivileges"`

	// Account remarks
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// Whether it is an admin account. Valid values: true (it is an admin account when the instance is a single-node type and AccountType is L0; when the instance is a two-node type and AccountType is L1), false (it is a standard account when AccountType is L3)
	IsAdmin *bool `json:"IsAdmin,omitnil,omitempty" name:"IsAdmin"`

	// Valid values: `win-windows authentication`, `sql-sqlserver authentication`. Default value: `sql-sqlserver authentication`
	Authentication *string `json:"Authentication,omitnil,omitempty" name:"Authentication"`

	// Account type, which is an extension field of `IsAdmin`. Valid values: `L0` (admin account, only for basic edition), `L1` (privileged account), `L2` (designated account), `L3` (standard account, default)
	AccountType *string `json:"AccountType,omitnil,omitempty" name:"AccountType"`

	// Whether CAM authentication is enabled
	IsCam *bool `json:"IsCam,omitnil,omitempty" name:"IsCam"`
}

type AccountDetail struct {
	// Account name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Account remarks
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// Account creation time
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Account status. 1: creating, 2: normal, 3: modifying, 4: resetting password, -1: deleting
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Account update time
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// Password update time
	PassTime *string `json:"PassTime,omitnil,omitempty" name:"PassTime"`

	// Internal account status, which should be `enable` normally
	InternalStatus *string `json:"InternalStatus,omitnil,omitempty" name:"InternalStatus"`

	// Information of read and write permissions of this account on relevant databases
	Dbs []*DBPrivilege `json:"Dbs,omitnil,omitempty" name:"Dbs"`

	// Whether it is an admin account
	IsAdmin *bool `json:"IsAdmin,omitnil,omitempty" name:"IsAdmin"`

	// Whether it is a CAM managed account
	IsCam *bool `json:"IsCam,omitnil,omitempty" name:"IsCam"`

	// Valid values: `win-windows authentication`, `sql-sqlserver authentication`.
	Authentication *string `json:"Authentication,omitnil,omitempty" name:"Authentication"`

	// The host required for `win-windows authentication` account
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// Account type. Valid values: `L0` (admin account, only for basic edition), `L1` (privileged account), `L2` (designated account), `L3` (standard account).
	AccountType *string `json:"AccountType,omitnil,omitempty" name:"AccountType"`
}

type AccountPassword struct {
	// Username
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// Password
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`
}

type AccountPrivilege struct {
	// Database username
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// Database permission. Valid values: `ReadWrite` (read-write), `ReadOnly` (read-only), `Delete` (delete the database permissions of this account), `DBOwner` (owner).
	Privilege *string `json:"Privilege,omitnil,omitempty" name:"Privilege"`

	// Account name. Valid values: `L0` (admin account, only for basic edition), `L1` (privileged account), `L2` (designated account), `L3` (standard account).
	AccountType *string `json:"AccountType,omitnil,omitempty" name:"AccountType"`
}

type AccountPrivilegeModifyInfo struct {
	// Database username
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// Account permission change information
	DBPrivileges []*DBPrivilegeModifyInfo `json:"DBPrivileges,omitnil,omitempty" name:"DBPrivileges"`

	// Whether it is an instance admin account. Valid values: `true` (Yes. When the instance is single-node and `AccountType` is `L0`, it's an admin account; when the instance is two-node and `AccountType` is `L1`, it's a privileged account), `false` (No. It's a standard account and `AccountType` is `L3`).
	IsAdmin *bool `json:"IsAdmin,omitnil,omitempty" name:"IsAdmin"`

	// Account type, which is an extension field of `IsAdmin`. Valid values: `L0` (admin account, only for basic edition), `L1` (privileged account), `L2` (designated account), `L3` (standard account, default)
	AccountType *string `json:"AccountType,omitnil,omitempty" name:"AccountType"`
}

type AccountRemark struct {
	// Account name
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// New remarks of account
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

type Backup struct {
	// File name. The name of an unarchived backup file is returned by the `DescribeBackupFiles` API instead of this parameter.
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// File size in KB. The size of an unarchived backup file is returned by the `DescribeBackupFiles` API instead of this parameter.
	Size *int64 `json:"Size,omitnil,omitempty" name:"Size"`

	// Backup start time
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Backup end time
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Private network download address. The download address of an unarchived backup file is returned by the `DescribeBackupFiles` API instead of this parameter.
	InternalAddr *string `json:"InternalAddr,omitnil,omitempty" name:"InternalAddr"`

	// Public network download address. The download address of an unarchived backup file is returned by the `DescribeBackupFiles` API instead of this parameter.
	ExternalAddr *string `json:"ExternalAddr,omitnil,omitempty" name:"ExternalAddr"`

	// Unique ID of a backup file, which is used by the `RestoreInstance` API. The unique ID of an unarchived backup file is returned by the `DescribeBackupFiles` API instead of this parameter.
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// Backup file status (0: creating, 1: succeeded, 2: failed)
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// List of databases for multi-database backup
	DBs []*string `json:"DBs,omitnil,omitempty" name:"DBs"`

	// Backup policy (0: instance backup, 1: multi-database backup)
	Strategy *int64 `json:"Strategy,omitnil,omitempty" name:"Strategy"`

	// Backup Mode. Valid values: `0` (scheduled backup); `1` (manual backup); `2` (archive backup).
	BackupWay *int64 `json:"BackupWay,omitnil,omitempty" name:"BackupWay"`

	// Backup task name (customizable)
	BackupName *string `json:"BackupName,omitnil,omitempty" name:"BackupName"`

	// Group ID of unarchived backup files, which can be used as a request parameter in the `DescribeBackupFiles` API to get details of unarchived backup files in the specified group. This parameter is invalid for archived backup files.
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// Backup file format. Valid values:`pkg` (archive file), `single` (unarchived files).
	BackupFormat *string `json:"BackupFormat,omitnil,omitempty" name:"BackupFormat"`

	// The code of current region where the instance resides
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// The download address of cross-region backup in target region
	CrossBackupAddr []*CrossBackupAddr `json:"CrossBackupAddr,omitnil,omitempty" name:"CrossBackupAddr"`

	// The target region and status of cross-region backup
	CrossBackupStatus []*CrossRegionStatus `json:"CrossBackupStatus,omitnil,omitempty" name:"CrossBackupStatus"`
}

type BackupFile struct {
	// Unique ID of a backup file
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// Backup file name
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// File size in KB
	Size *uint64 `json:"Size,omitnil,omitempty" name:"Size"`

	// Name of the database corresponding to the backup file
	DBs []*string `json:"DBs,omitnil,omitempty" name:"DBs"`

	// Download address
	DownloadLink *string `json:"DownloadLink,omitnil,omitempty" name:"DownloadLink"`

	// The code of the region where current instance resides
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// The target region and download address of cross-region backup
	CrossBackupAddr []*CrossBackupAddr `json:"CrossBackupAddr,omitnil,omitempty" name:"CrossBackupAddr"`
}

type BusinessIntelligenceFile struct {
	// File name
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// File type
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// File COS_URL
	FileURL *string `json:"FileURL,omitnil,omitempty" name:"FileURL"`

	// The file path on the server
	FilePath *string `json:"FilePath,omitnil,omitempty" name:"FilePath"`

	// File size in bytes
	FileSize *int64 `json:"FileSize,omitnil,omitempty" name:"FileSize"`

	// File MD5 value
	FileMd5 *string `json:"FileMd5,omitnil,omitempty" name:"FileMd5"`

	// File deployment status. Valid values: `1`(Initialize to be deployed), `2` (Deploying), `3` (Deployment successful), `4` (Deployment failed).
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Remarks
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// File creation time
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Start time of file deployment
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time of file deployment
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Returned error message
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// Business intelligence instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Operation information
	Action *FileAction `json:"Action,omitnil,omitempty" name:"Action"`
}

// Predefined struct for user
type CloneDBRequestParams struct {
	// Instance ID in the format of mssql-j8kv137v
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Clone and rename the databases specified in `ReNameRestoreDatabase`. Please note that the clones must be renamed.
	RenameRestore []*RenameRestoreDatabase `json:"RenameRestore,omitnil,omitempty" name:"RenameRestore"`
}

type CloneDBRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of mssql-j8kv137v
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Clone and rename the databases specified in `ReNameRestoreDatabase`. Please note that the clones must be renamed.
	RenameRestore []*RenameRestoreDatabase `json:"RenameRestore,omitnil,omitempty" name:"RenameRestore"`
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

// Predefined struct for user
type CloneDBResponseParams struct {
	// Async task request ID, which can be used in the `DescribeFlowStatus` API to query the execution result of an async task
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CloneDBResponse struct {
	*tchttp.BaseResponse
	Response *CloneDBResponseParams `json:"Response"`
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

// Predefined struct for user
type CloseInterCommunicationRequestParams struct {
	// IDs of instances with interconnection disabled
	InstanceIdSet []*string `json:"InstanceIdSet,omitnil,omitempty" name:"InstanceIdSet"`
}

type CloseInterCommunicationRequest struct {
	*tchttp.BaseRequest
	
	// IDs of instances with interconnection disabled
	InstanceIdSet []*string `json:"InstanceIdSet,omitnil,omitempty" name:"InstanceIdSet"`
}

func (r *CloseInterCommunicationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseInterCommunicationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIdSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CloseInterCommunicationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloseInterCommunicationResponseParams struct {
	// IDs of instance and async task
	InterInstanceFlowSet []*InterInstanceFlow `json:"InterInstanceFlowSet,omitnil,omitempty" name:"InterInstanceFlowSet"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CloseInterCommunicationResponse struct {
	*tchttp.BaseResponse
	Response *CloseInterCommunicationResponseParams `json:"Response"`
}

func (r *CloseInterCommunicationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseInterCommunicationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CosUploadBackupFile struct {
	// Backup name
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// Backup size
	Size *int64 `json:"Size,omitnil,omitempty" name:"Size"`
}

// Predefined struct for user
type CreateAccountRequestParams struct {
	// Database instance ID in the format of mssql-njj2mtpl
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Database instance account information
	Accounts []*AccountCreateInfo `json:"Accounts,omitnil,omitempty" name:"Accounts"`
}

type CreateAccountRequest struct {
	*tchttp.BaseRequest
	
	// Database instance ID in the format of mssql-njj2mtpl
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Database instance account information
	Accounts []*AccountCreateInfo `json:"Accounts,omitnil,omitempty" name:"Accounts"`
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

// Predefined struct for user
type CreateAccountResponseParams struct {
	// Task flow ID
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAccountResponse struct {
	*tchttp.BaseResponse
	Response *CreateAccountResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateBackupMigrationRequestParams struct {
	// ID of imported target instance
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Migration task restoration type. FULL: full backup restoration, FULL_LOG: full backup and transaction log restoration, FULL_DIFF: full backup and differential backup restoration
	RecoveryType *string `json:"RecoveryType,omitnil,omitempty" name:"RecoveryType"`

	// Backup upload type. COS_URL: the backup is stored in user's Cloud Object Storage, with URL provided. COS_UPLOAD: the backup is stored in the application's Cloud Object Storage and needs to be uploaded by the user.
	UploadType *string `json:"UploadType,omitnil,omitempty" name:"UploadType"`

	// Task name
	MigrationName *string `json:"MigrationName,omitnil,omitempty" name:"MigrationName"`

	// If the UploadType is COS_URL, fill in the URL here. If the UploadType is COS_UPLOAD, fill in the name of the backup file here. Only 1 backup file is supported, but a backup file can involve multiple databases.
	BackupFiles []*string `json:"BackupFiles,omitnil,omitempty" name:"BackupFiles"`
}

type CreateBackupMigrationRequest struct {
	*tchttp.BaseRequest
	
	// ID of imported target instance
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Migration task restoration type. FULL: full backup restoration, FULL_LOG: full backup and transaction log restoration, FULL_DIFF: full backup and differential backup restoration
	RecoveryType *string `json:"RecoveryType,omitnil,omitempty" name:"RecoveryType"`

	// Backup upload type. COS_URL: the backup is stored in user's Cloud Object Storage, with URL provided. COS_UPLOAD: the backup is stored in the application's Cloud Object Storage and needs to be uploaded by the user.
	UploadType *string `json:"UploadType,omitnil,omitempty" name:"UploadType"`

	// Task name
	MigrationName *string `json:"MigrationName,omitnil,omitempty" name:"MigrationName"`

	// If the UploadType is COS_URL, fill in the URL here. If the UploadType is COS_UPLOAD, fill in the name of the backup file here. Only 1 backup file is supported, but a backup file can involve multiple databases.
	BackupFiles []*string `json:"BackupFiles,omitnil,omitempty" name:"BackupFiles"`
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

// Predefined struct for user
type CreateBackupMigrationResponseParams struct {
	// Backup import task ID
	BackupMigrationId *string `json:"BackupMigrationId,omitnil,omitempty" name:"BackupMigrationId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateBackupMigrationResponse struct {
	*tchttp.BaseResponse
	Response *CreateBackupMigrationResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateBackupRequestParams struct {
	// Backup policy (0: instance backup, 1: multi-database backup)
	Strategy *int64 `json:"Strategy,omitnil,omitempty" name:"Strategy"`

	// List of names of databases to be backed up (required only for multi-database backup)
	DBNames []*string `json:"DBNames,omitnil,omitempty" name:"DBNames"`

	// (Required) Instance ID in the format of mssql-i1z41iwd
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Backup name. If this parameter is left empty, a backup name in the format of "[Instance ID]_[Backup start timestamp]" will be automatically generated.
	BackupName *string `json:"BackupName,omitnil,omitempty" name:"BackupName"`


	StorageStrategy *int64 `json:"StorageStrategy,omitnil,omitempty" name:"StorageStrategy"`
}

type CreateBackupRequest struct {
	*tchttp.BaseRequest
	
	// Backup policy (0: instance backup, 1: multi-database backup)
	Strategy *int64 `json:"Strategy,omitnil,omitempty" name:"Strategy"`

	// List of names of databases to be backed up (required only for multi-database backup)
	DBNames []*string `json:"DBNames,omitnil,omitempty" name:"DBNames"`

	// (Required) Instance ID in the format of mssql-i1z41iwd
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Backup name. If this parameter is left empty, a backup name in the format of "[Instance ID]_[Backup start timestamp]" will be automatically generated.
	BackupName *string `json:"BackupName,omitnil,omitempty" name:"BackupName"`

	StorageStrategy *int64 `json:"StorageStrategy,omitnil,omitempty" name:"StorageStrategy"`
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
	delete(f, "StorageStrategy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBackupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBackupResponseParams struct {
	// The async job ID
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateBackupResponse struct {
	*tchttp.BaseResponse
	Response *CreateBackupResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateBasicDBInstancesRequestParams struct {

	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`


	Cpu *uint64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`


	Memory *uint64 `json:"Memory,omitnil,omitempty" name:"Memory"`


	Storage *uint64 `json:"Storage,omitnil,omitempty" name:"Storage"`


	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`


	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`


	MachineType *string `json:"MachineType,omitnil,omitempty" name:"MachineType"`


	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`


	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`


	GoodsNum *uint64 `json:"GoodsNum,omitnil,omitempty" name:"GoodsNum"`


	DBVersion *string `json:"DBVersion,omitnil,omitempty" name:"DBVersion"`


	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`


	SecurityGroupList []*string `json:"SecurityGroupList,omitnil,omitempty" name:"SecurityGroupList"`


	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`


	AutoVoucher *int64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`


	VoucherIds []*string `json:"VoucherIds,omitnil,omitempty" name:"VoucherIds"`


	Weekly []*int64 `json:"Weekly,omitnil,omitempty" name:"Weekly"`


	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`


	Span *int64 `json:"Span,omitnil,omitempty" name:"Span"`


	ResourceTags []*ResourceTag `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`


	Collation *string `json:"Collation,omitnil,omitempty" name:"Collation"`


	TimeZone *string `json:"TimeZone,omitnil,omitempty" name:"TimeZone"`
}

type CreateBasicDBInstancesRequest struct {
	*tchttp.BaseRequest
	
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	Cpu *uint64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	Memory *uint64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	Storage *uint64 `json:"Storage,omitnil,omitempty" name:"Storage"`

	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	MachineType *string `json:"MachineType,omitnil,omitempty" name:"MachineType"`

	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	GoodsNum *uint64 `json:"GoodsNum,omitnil,omitempty" name:"GoodsNum"`

	DBVersion *string `json:"DBVersion,omitnil,omitempty" name:"DBVersion"`

	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	SecurityGroupList []*string `json:"SecurityGroupList,omitnil,omitempty" name:"SecurityGroupList"`

	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	AutoVoucher *int64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	VoucherIds []*string `json:"VoucherIds,omitnil,omitempty" name:"VoucherIds"`

	Weekly []*int64 `json:"Weekly,omitnil,omitempty" name:"Weekly"`

	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	Span *int64 `json:"Span,omitnil,omitempty" name:"Span"`

	ResourceTags []*ResourceTag `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`

	Collation *string `json:"Collation,omitnil,omitempty" name:"Collation"`

	TimeZone *string `json:"TimeZone,omitnil,omitempty" name:"TimeZone"`
}

func (r *CreateBasicDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBasicDBInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Zone")
	delete(f, "Cpu")
	delete(f, "Memory")
	delete(f, "Storage")
	delete(f, "SubnetId")
	delete(f, "VpcId")
	delete(f, "MachineType")
	delete(f, "InstanceChargeType")
	delete(f, "ProjectId")
	delete(f, "GoodsNum")
	delete(f, "DBVersion")
	delete(f, "Period")
	delete(f, "SecurityGroupList")
	delete(f, "AutoRenewFlag")
	delete(f, "AutoVoucher")
	delete(f, "VoucherIds")
	delete(f, "Weekly")
	delete(f, "StartTime")
	delete(f, "Span")
	delete(f, "ResourceTags")
	delete(f, "Collation")
	delete(f, "TimeZone")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBasicDBInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBasicDBInstancesResponseParams struct {

	DealName *string `json:"DealName,omitnil,omitempty" name:"DealName"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateBasicDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *CreateBasicDBInstancesResponseParams `json:"Response"`
}

func (r *CreateBasicDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBasicDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBusinessDBInstancesRequestParams struct {
	// Instance AZ, such as ap-guangzhou-1 (Guangzhou Zone 1). Purchasable AZs for an instance can be obtained through the`DescribeZones` API.
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Instance memory size in GB
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// Instance disk size in GB
	Storage *int64 `json:"Storage,omitnil,omitempty" name:"Storage"`

	// The number of CPU cores of the instance you want to purchase.
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// The host type of purchased instance. Valid values: `CLOUD_PREMIUM` (virtual machine with premium cloud disk), `CLOUD_SSD` (virtual machine with SSD).
	MachineType *string `json:"MachineType,omitnil,omitempty" name:"MachineType"`

	// Project ID
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Number of instances purchased this time. Default value: `1`.
	GoodsNum *int64 `json:"GoodsNum,omitnil,omitempty" name:"GoodsNum"`

	// VPC subnet ID in the format of subnet-bdoe83fa. Both `SubnetId` and `VpcId` need to be set or unset at the same time.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// VPC ID in the format of vpc-dsp338hz. Both `SubnetId` and `VpcId` need to be set or unset at the same time.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// - Supported versions of business intelligence server. Valid values: `201603` (SQL Server 2016 Integration Services), `201703` (SQL Server 2017 Integration Services), `201903` (SQL Server 2019 Integration Services). Default value: `201903`. As the purchasable versions are region-specific, you can use the `DescribeProductConfig` API to query the information of purchasable versions in each region.
	DBVersion *string `json:"DBVersion,omitnil,omitempty" name:"DBVersion"`

	// Security group list, which contains security group IDs in the format of sg-xxx.
	SecurityGroupList []*string `json:"SecurityGroupList,omitnil,omitempty" name:"SecurityGroupList"`

	// Configuration of the maintenance window, which specifies the day of the week when maintenance can be performed. Valid values: `1` (Monday), `2` (Tuesday), `3` (Wednesday), `4` (Thursday), `5` (Friday), `6` (Saturday), `7` (Sunday).
	Weekly []*int64 `json:"Weekly,omitnil,omitempty" name:"Weekly"`

	// Configuration of the maintenance window, which specifies the start time of daily maintenance.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Configuration of the maintenance window, which specifies the maintenance duration in hours.
	Span *int64 `json:"Span,omitnil,omitempty" name:"Span"`

	// Tags associated with the instances to be created
	ResourceTags []*ResourceTag `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`
}

type CreateBusinessDBInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Instance AZ, such as ap-guangzhou-1 (Guangzhou Zone 1). Purchasable AZs for an instance can be obtained through the`DescribeZones` API.
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Instance memory size in GB
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// Instance disk size in GB
	Storage *int64 `json:"Storage,omitnil,omitempty" name:"Storage"`

	// The number of CPU cores of the instance you want to purchase.
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// The host type of purchased instance. Valid values: `CLOUD_PREMIUM` (virtual machine with premium cloud disk), `CLOUD_SSD` (virtual machine with SSD).
	MachineType *string `json:"MachineType,omitnil,omitempty" name:"MachineType"`

	// Project ID
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Number of instances purchased this time. Default value: `1`.
	GoodsNum *int64 `json:"GoodsNum,omitnil,omitempty" name:"GoodsNum"`

	// VPC subnet ID in the format of subnet-bdoe83fa. Both `SubnetId` and `VpcId` need to be set or unset at the same time.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// VPC ID in the format of vpc-dsp338hz. Both `SubnetId` and `VpcId` need to be set or unset at the same time.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// - Supported versions of business intelligence server. Valid values: `201603` (SQL Server 2016 Integration Services), `201703` (SQL Server 2017 Integration Services), `201903` (SQL Server 2019 Integration Services). Default value: `201903`. As the purchasable versions are region-specific, you can use the `DescribeProductConfig` API to query the information of purchasable versions in each region.
	DBVersion *string `json:"DBVersion,omitnil,omitempty" name:"DBVersion"`

	// Security group list, which contains security group IDs in the format of sg-xxx.
	SecurityGroupList []*string `json:"SecurityGroupList,omitnil,omitempty" name:"SecurityGroupList"`

	// Configuration of the maintenance window, which specifies the day of the week when maintenance can be performed. Valid values: `1` (Monday), `2` (Tuesday), `3` (Wednesday), `4` (Thursday), `5` (Friday), `6` (Saturday), `7` (Sunday).
	Weekly []*int64 `json:"Weekly,omitnil,omitempty" name:"Weekly"`

	// Configuration of the maintenance window, which specifies the start time of daily maintenance.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Configuration of the maintenance window, which specifies the maintenance duration in hours.
	Span *int64 `json:"Span,omitnil,omitempty" name:"Span"`

	// Tags associated with the instances to be created
	ResourceTags []*ResourceTag `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`
}

func (r *CreateBusinessDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBusinessDBInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Zone")
	delete(f, "Memory")
	delete(f, "Storage")
	delete(f, "Cpu")
	delete(f, "MachineType")
	delete(f, "ProjectId")
	delete(f, "GoodsNum")
	delete(f, "SubnetId")
	delete(f, "VpcId")
	delete(f, "DBVersion")
	delete(f, "SecurityGroupList")
	delete(f, "Weekly")
	delete(f, "StartTime")
	delete(f, "Span")
	delete(f, "ResourceTags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBusinessDBInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBusinessDBInstancesResponseParams struct {
	// Order name
	DealName *string `json:"DealName,omitnil,omitempty" name:"DealName"`

	// Process ID Note: This field may return null, indicating that no valid values can be obtained.
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// IDs of instances Note: This field may return null, indicating that no valid values can be obtained.
	InstanceIdSet []*string `json:"InstanceIdSet,omitnil,omitempty" name:"InstanceIdSet"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateBusinessDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *CreateBusinessDBInstancesResponseParams `json:"Response"`
}

func (r *CreateBusinessDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBusinessDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBusinessIntelligenceFileRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`


	FileURL *string `json:"FileURL,omitnil,omitempty" name:"FileURL"`

	// File type. Valid values: `FLAT` (flat file as data source), `SSIS` (.ispac SSIS package file)
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// Remarks
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

type CreateBusinessIntelligenceFileRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	FileURL *string `json:"FileURL,omitnil,omitempty" name:"FileURL"`

	// File type. Valid values: `FLAT` (flat file as data source), `SSIS` (.ispac SSIS package file)
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// Remarks
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

func (r *CreateBusinessIntelligenceFileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBusinessIntelligenceFileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "FileURL")
	delete(f, "FileType")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBusinessIntelligenceFileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBusinessIntelligenceFileResponseParams struct {
	// File name
	FileTaskId *string `json:"FileTaskId,omitnil,omitempty" name:"FileTaskId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateBusinessIntelligenceFileResponse struct {
	*tchttp.BaseResponse
	Response *CreateBusinessIntelligenceFileResponseParams `json:"Response"`
}

func (r *CreateBusinessIntelligenceFileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBusinessIntelligenceFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCloudDBInstancesRequestParams struct {
	// Instance AZ, such as `ap-guangzhou-1` (Guangzhou Zone 1). Purchasable AZs for an instance can be obtained through the`DescribeZones` API.
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Instance memory size in GB
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// Instance disk size in GB
	Storage *int64 `json:"Storage,omitnil,omitempty" name:"Storage"`

	// Number of CPU cores
	Cpu *uint64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// The host type of the purchased instance. Valid values: `CLOUD_HSSD` (virtual machine with enhanced SSD), `CLOUD_TSSD` (virtual machine with ulTra SSD), `CLOUD_BSSD` (virtual machine with balanced SSD).
	MachineType *string `json:"MachineType,omitnil,omitempty" name:"MachineType"`

	// Billing mode. Valid values: `PREPAID` (monthly subscription), `POSTPAID` (pay-as-you-go).
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// Project ID
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Number of instances purchased this time. Default value: `1`.  Maximum value: `10`.
	GoodsNum *int64 `json:"GoodsNum,omitnil,omitempty" name:"GoodsNum"`

	// VPC subnet ID in the format of `subnet-bdoe83fa`. Both `SubnetId` and `VpcId` need to be set or unset at the same time.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// VPC ID in the format of `vpc-dsp338hz`. Both `SubnetId` and `VpcId` need to be set or unset at the same time.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// The purchase period of an instance. Default value: `1` (one month).  Maximum value: `48`.
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// Whether to automatically use voucher. Valid values: `0` (no, default), `1` (yes).
	AutoVoucher *int64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// Array of voucher IDs (currently, only one voucher can be used per order)
	VoucherIds []*string `json:"VoucherIds,omitnil,omitempty" name:"VoucherIds"`

	// SQL Server version. Valid values:  `2008R2` (SQL Server 2008 R2 Enterprise), `2012SP3` (SQL Server 2012 Enterprise), `201202` (SQL Server 2012 Standard), `2014SP2` (SQL Server 2014 Enterprise), 201402 (SQL Server 2014 Standard), `2016SP1` (SQL Server 2016 Enterprise), `201602` (SQL Server 2016 Standard), `2017` (SQL Server 2017 Enterprise), `201702` (SQL Server 2017 Standard), `2019` (SQL Server 2019 Enterprise), `201902` (SQL Server 2019 Standard).  Default value: `2008R2`.  The available version varies by region, and you can pull the version information through the `DescribeProductConfig` API.
	DBVersion *string `json:"DBVersion,omitnil,omitempty" name:"DBVersion"`

	// Auto-renewal flag, which takes effect only when purchasing a monthly subscribed instance.  Valid values:  `0` (auto-renewal disabled), `1` (auto-renewal enabled). Default value: `0`.
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// Security group list, which contains security group IDs in the format of `sg-xxx`.
	SecurityGroupList []*string `json:"SecurityGroupList,omitnil,omitempty" name:"SecurityGroupList"`

	// Configuration of the maintenance window, which specifies the day of the week when maintenance can be performed. Valid values: `1` (Monday), `2` (Tuesday), `3` (Wednesday), `4` (Thursday), `5` (Friday), `6` (Saturday), `7` (Sunday).
	Weekly []*int64 `json:"Weekly,omitnil,omitempty" name:"Weekly"`

	// Configuration of the maintenance window, which specifies the start time of daily maintenance.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Configuration of the maintenance window, which specifies the maintenance duration in hours. Hour
	Span *int64 `json:"Span,omitnil,omitempty" name:"Span"`

	// Whether to deploy across AZs. Default value: `false`.
	MultiZones *bool `json:"MultiZones,omitnil,omitempty" name:"MultiZones"`

	// Tags associated with the instances to be created
	ResourceTags []*ResourceTag `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`

	// Collation of system character sets. Default value:  `Chinese_PRC_CI_AS`.
	Collation *string `json:"Collation,omitnil,omitempty" name:"Collation"`

	// System time zone. Default value:  `China Standard Time`.
	TimeZone *string `json:"TimeZone,omitnil,omitempty" name:"TimeZone"`
}

type CreateCloudDBInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Instance AZ, such as `ap-guangzhou-1` (Guangzhou Zone 1). Purchasable AZs for an instance can be obtained through the`DescribeZones` API.
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Instance memory size in GB
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// Instance disk size in GB
	Storage *int64 `json:"Storage,omitnil,omitempty" name:"Storage"`

	// Number of CPU cores
	Cpu *uint64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// The host type of the purchased instance. Valid values: `CLOUD_HSSD` (virtual machine with enhanced SSD), `CLOUD_TSSD` (virtual machine with ulTra SSD), `CLOUD_BSSD` (virtual machine with balanced SSD).
	MachineType *string `json:"MachineType,omitnil,omitempty" name:"MachineType"`

	// Billing mode. Valid values: `PREPAID` (monthly subscription), `POSTPAID` (pay-as-you-go).
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// Project ID
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Number of instances purchased this time. Default value: `1`.  Maximum value: `10`.
	GoodsNum *int64 `json:"GoodsNum,omitnil,omitempty" name:"GoodsNum"`

	// VPC subnet ID in the format of `subnet-bdoe83fa`. Both `SubnetId` and `VpcId` need to be set or unset at the same time.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// VPC ID in the format of `vpc-dsp338hz`. Both `SubnetId` and `VpcId` need to be set or unset at the same time.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// The purchase period of an instance. Default value: `1` (one month).  Maximum value: `48`.
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// Whether to automatically use voucher. Valid values: `0` (no, default), `1` (yes).
	AutoVoucher *int64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// Array of voucher IDs (currently, only one voucher can be used per order)
	VoucherIds []*string `json:"VoucherIds,omitnil,omitempty" name:"VoucherIds"`

	// SQL Server version. Valid values:  `2008R2` (SQL Server 2008 R2 Enterprise), `2012SP3` (SQL Server 2012 Enterprise), `201202` (SQL Server 2012 Standard), `2014SP2` (SQL Server 2014 Enterprise), 201402 (SQL Server 2014 Standard), `2016SP1` (SQL Server 2016 Enterprise), `201602` (SQL Server 2016 Standard), `2017` (SQL Server 2017 Enterprise), `201702` (SQL Server 2017 Standard), `2019` (SQL Server 2019 Enterprise), `201902` (SQL Server 2019 Standard).  Default value: `2008R2`.  The available version varies by region, and you can pull the version information through the `DescribeProductConfig` API.
	DBVersion *string `json:"DBVersion,omitnil,omitempty" name:"DBVersion"`

	// Auto-renewal flag, which takes effect only when purchasing a monthly subscribed instance.  Valid values:  `0` (auto-renewal disabled), `1` (auto-renewal enabled). Default value: `0`.
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// Security group list, which contains security group IDs in the format of `sg-xxx`.
	SecurityGroupList []*string `json:"SecurityGroupList,omitnil,omitempty" name:"SecurityGroupList"`

	// Configuration of the maintenance window, which specifies the day of the week when maintenance can be performed. Valid values: `1` (Monday), `2` (Tuesday), `3` (Wednesday), `4` (Thursday), `5` (Friday), `6` (Saturday), `7` (Sunday).
	Weekly []*int64 `json:"Weekly,omitnil,omitempty" name:"Weekly"`

	// Configuration of the maintenance window, which specifies the start time of daily maintenance.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Configuration of the maintenance window, which specifies the maintenance duration in hours. Hour
	Span *int64 `json:"Span,omitnil,omitempty" name:"Span"`

	// Whether to deploy across AZs. Default value: `false`.
	MultiZones *bool `json:"MultiZones,omitnil,omitempty" name:"MultiZones"`

	// Tags associated with the instances to be created
	ResourceTags []*ResourceTag `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`

	// Collation of system character sets. Default value:  `Chinese_PRC_CI_AS`.
	Collation *string `json:"Collation,omitnil,omitempty" name:"Collation"`

	// System time zone. Default value:  `China Standard Time`.
	TimeZone *string `json:"TimeZone,omitnil,omitempty" name:"TimeZone"`
}

func (r *CreateCloudDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCloudDBInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Zone")
	delete(f, "Memory")
	delete(f, "Storage")
	delete(f, "Cpu")
	delete(f, "MachineType")
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
	delete(f, "MultiZones")
	delete(f, "ResourceTags")
	delete(f, "Collation")
	delete(f, "TimeZone")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCloudDBInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCloudDBInstancesResponseParams struct {
	// Order name
	DealName *string `json:"DealName,omitnil,omitempty" name:"DealName"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCloudDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *CreateCloudDBInstancesResponseParams `json:"Response"`
}

func (r *CreateCloudDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCloudDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCloudReadOnlyDBInstancesRequestParams struct {
	// Instance ID in the format of  `mssql-3l3fgqn7`.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Instance AZ, such as `ap-guangzhou-1` (Guangzhou Zone 1). Purchasable AZs for an instance can be obtained through the`DescribeZones` API.
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Read-only group types. Valid values: `1` (each read-only replica is placed in one auto-created read-only group), `2` (all read-only replicas are placed in one auto-created read-only group), `3` (all read-only replicas are placed in one existing read-only group).
	ReadOnlyGroupType *int64 `json:"ReadOnlyGroupType,omitnil,omitempty" name:"ReadOnlyGroupType"`

	// Instance memory size in GB
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// Instance disk size in GB
	Storage *int64 `json:"Storage,omitnil,omitempty" name:"Storage"`

	// Number of instance cores
	Cpu *uint64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// The host type of purchased instance. Valid values: `CLOUD_HSSD` (virtual machine with enhanced SSD), `CLOUD_TSSD` (virtual machine with ulTra SSD), `CLOUD_BSSD` (virtual machine with balanced SSD).
	MachineType *string `json:"MachineType,omitnil,omitempty" name:"MachineType"`

	// Valid values: `0` (not upgrade the primary instance by default), `1` (upgrade the primary instance to complete the RO deployment).  You need to pass in `1` for this parameter and upgrade the primary instance to cluster edition.
	ReadOnlyGroupForcedUpgrade *int64 `json:"ReadOnlyGroupForcedUpgrade,omitnil,omitempty" name:"ReadOnlyGroupForcedUpgrade"`

	// Existing read-only group ID, which is required when `ReadOnlyGroupType` is `3`.
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitnil,omitempty" name:"ReadOnlyGroupId"`

	// New read-only group ID, which is required when `ReadOnlyGroupType` is `2`.
	ReadOnlyGroupName *string `json:"ReadOnlyGroupName,omitnil,omitempty" name:"ReadOnlyGroupName"`

	// Whether delayed read-only instance removal is enabled in a new read-only group, which is required when `ReadOnlyGroupType` is `2`. Valid values: `1` (enabled), `0` (disabled).  The read-only replica will be automatically removed when the delay between it and the primary instance exceeds the threshold.
	ReadOnlyGroupIsOfflineDelay *int64 `json:"ReadOnlyGroupIsOfflineDelay,omitnil,omitempty" name:"ReadOnlyGroupIsOfflineDelay"`

	// The delay threshold for a new read-only group, which is required when `ReadOnlyGroupType` is `2` and `ReadOnlyGroupIsOfflineDelay` is `1`.
	ReadOnlyGroupMaxDelayTime *int64 `json:"ReadOnlyGroupMaxDelayTime,omitnil,omitempty" name:"ReadOnlyGroupMaxDelayTime"`

	// Minimum number of reserved read-only replicas when the delayed removal is enabled for the new read-only group, which is required when `ReadOnlyGroupType` is `2` and `ReadOnlyGroupIsOfflineDelay` is `1`.
	ReadOnlyGroupMinInGroup *int64 `json:"ReadOnlyGroupMinInGroup,omitnil,omitempty" name:"ReadOnlyGroupMinInGroup"`

	// Billing mode. Valid values: `PREPAID` (monthly subscription), `POSTPAID` (pay-as-you-go).
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// Number of read-only instances to be purchased this time. Default value: `2`.
	GoodsNum *int64 `json:"GoodsNum,omitnil,omitempty" name:"GoodsNum"`

	// VPC subnet ID in the format of `subnet-bdoe83fa`. Both `SubnetId` and `VpcId` need to be set or unset at the same time.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// VPC ID in the format of `vpc-dsp338hz`. Both `SubnetId` and `VpcId` need to be set or unset at the same time.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// The purchase period of an instance. Default value: `1` (one month).  Maximum value: `48`.
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// Security group list, which contains security group IDs in the format of `sg-xxx`.
	SecurityGroupList []*string `json:"SecurityGroupList,omitnil,omitempty" name:"SecurityGroupList"`

	// Whether to automatically use voucher. Valid values: `0` (no, default), `1` (yes).
	AutoVoucher *int64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// Array of voucher IDs (currently, only one voucher can be used per order)
	VoucherIds []*string `json:"VoucherIds,omitnil,omitempty" name:"VoucherIds"`

	// Tags associated with the instances to be created
	ResourceTags []*ResourceTag `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`

	// Collation of system character sets. Default value:  Chinese_PRC_CI_AS
	Collation *string `json:"Collation,omitnil,omitempty" name:"Collation"`

	// System time zone. Default value:  `China Standard Time`
	TimeZone *string `json:"TimeZone,omitnil,omitempty" name:"TimeZone"`
}

type CreateCloudReadOnlyDBInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of  `mssql-3l3fgqn7`.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Instance AZ, such as `ap-guangzhou-1` (Guangzhou Zone 1). Purchasable AZs for an instance can be obtained through the`DescribeZones` API.
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Read-only group types. Valid values: `1` (each read-only replica is placed in one auto-created read-only group), `2` (all read-only replicas are placed in one auto-created read-only group), `3` (all read-only replicas are placed in one existing read-only group).
	ReadOnlyGroupType *int64 `json:"ReadOnlyGroupType,omitnil,omitempty" name:"ReadOnlyGroupType"`

	// Instance memory size in GB
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// Instance disk size in GB
	Storage *int64 `json:"Storage,omitnil,omitempty" name:"Storage"`

	// Number of instance cores
	Cpu *uint64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// The host type of purchased instance. Valid values: `CLOUD_HSSD` (virtual machine with enhanced SSD), `CLOUD_TSSD` (virtual machine with ulTra SSD), `CLOUD_BSSD` (virtual machine with balanced SSD).
	MachineType *string `json:"MachineType,omitnil,omitempty" name:"MachineType"`

	// Valid values: `0` (not upgrade the primary instance by default), `1` (upgrade the primary instance to complete the RO deployment).  You need to pass in `1` for this parameter and upgrade the primary instance to cluster edition.
	ReadOnlyGroupForcedUpgrade *int64 `json:"ReadOnlyGroupForcedUpgrade,omitnil,omitempty" name:"ReadOnlyGroupForcedUpgrade"`

	// Existing read-only group ID, which is required when `ReadOnlyGroupType` is `3`.
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitnil,omitempty" name:"ReadOnlyGroupId"`

	// New read-only group ID, which is required when `ReadOnlyGroupType` is `2`.
	ReadOnlyGroupName *string `json:"ReadOnlyGroupName,omitnil,omitempty" name:"ReadOnlyGroupName"`

	// Whether delayed read-only instance removal is enabled in a new read-only group, which is required when `ReadOnlyGroupType` is `2`. Valid values: `1` (enabled), `0` (disabled).  The read-only replica will be automatically removed when the delay between it and the primary instance exceeds the threshold.
	ReadOnlyGroupIsOfflineDelay *int64 `json:"ReadOnlyGroupIsOfflineDelay,omitnil,omitempty" name:"ReadOnlyGroupIsOfflineDelay"`

	// The delay threshold for a new read-only group, which is required when `ReadOnlyGroupType` is `2` and `ReadOnlyGroupIsOfflineDelay` is `1`.
	ReadOnlyGroupMaxDelayTime *int64 `json:"ReadOnlyGroupMaxDelayTime,omitnil,omitempty" name:"ReadOnlyGroupMaxDelayTime"`

	// Minimum number of reserved read-only replicas when the delayed removal is enabled for the new read-only group, which is required when `ReadOnlyGroupType` is `2` and `ReadOnlyGroupIsOfflineDelay` is `1`.
	ReadOnlyGroupMinInGroup *int64 `json:"ReadOnlyGroupMinInGroup,omitnil,omitempty" name:"ReadOnlyGroupMinInGroup"`

	// Billing mode. Valid values: `PREPAID` (monthly subscription), `POSTPAID` (pay-as-you-go).
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// Number of read-only instances to be purchased this time. Default value: `2`.
	GoodsNum *int64 `json:"GoodsNum,omitnil,omitempty" name:"GoodsNum"`

	// VPC subnet ID in the format of `subnet-bdoe83fa`. Both `SubnetId` and `VpcId` need to be set or unset at the same time.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// VPC ID in the format of `vpc-dsp338hz`. Both `SubnetId` and `VpcId` need to be set or unset at the same time.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// The purchase period of an instance. Default value: `1` (one month).  Maximum value: `48`.
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// Security group list, which contains security group IDs in the format of `sg-xxx`.
	SecurityGroupList []*string `json:"SecurityGroupList,omitnil,omitempty" name:"SecurityGroupList"`

	// Whether to automatically use voucher. Valid values: `0` (no, default), `1` (yes).
	AutoVoucher *int64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// Array of voucher IDs (currently, only one voucher can be used per order)
	VoucherIds []*string `json:"VoucherIds,omitnil,omitempty" name:"VoucherIds"`

	// Tags associated with the instances to be created
	ResourceTags []*ResourceTag `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`

	// Collation of system character sets. Default value:  Chinese_PRC_CI_AS
	Collation *string `json:"Collation,omitnil,omitempty" name:"Collation"`

	// System time zone. Default value:  `China Standard Time`
	TimeZone *string `json:"TimeZone,omitnil,omitempty" name:"TimeZone"`
}

func (r *CreateCloudReadOnlyDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCloudReadOnlyDBInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Zone")
	delete(f, "ReadOnlyGroupType")
	delete(f, "Memory")
	delete(f, "Storage")
	delete(f, "Cpu")
	delete(f, "MachineType")
	delete(f, "ReadOnlyGroupForcedUpgrade")
	delete(f, "ReadOnlyGroupId")
	delete(f, "ReadOnlyGroupName")
	delete(f, "ReadOnlyGroupIsOfflineDelay")
	delete(f, "ReadOnlyGroupMaxDelayTime")
	delete(f, "ReadOnlyGroupMinInGroup")
	delete(f, "InstanceChargeType")
	delete(f, "GoodsNum")
	delete(f, "SubnetId")
	delete(f, "VpcId")
	delete(f, "Period")
	delete(f, "SecurityGroupList")
	delete(f, "AutoVoucher")
	delete(f, "VoucherIds")
	delete(f, "ResourceTags")
	delete(f, "Collation")
	delete(f, "TimeZone")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCloudReadOnlyDBInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCloudReadOnlyDBInstancesResponseParams struct {
	// Order name in array
	DealNames []*string `json:"DealNames,omitnil,omitempty" name:"DealNames"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCloudReadOnlyDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *CreateCloudReadOnlyDBInstancesResponseParams `json:"Response"`
}

func (r *CreateCloudReadOnlyDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCloudReadOnlyDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDBInstancesRequestParams struct {
	// Instance AZ, such as ap-guangzhou-1 (Guangzhou Zone 1). Purchasable AZs for an instance can be obtained through the `DescribeZones` API
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Instance memory size in GB
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// Instance storage capacity in GB
	Storage *int64 `json:"Storage,omitnil,omitempty" name:"Storage"`

	// Billing mode. Valid value: POSTPAID (pay-as-you-go).
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// Project ID
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Number of instances purchased this time. Default value: 1. Maximum value: 10
	GoodsNum *int64 `json:"GoodsNum,omitnil,omitempty" name:"GoodsNum"`

	// VPC subnet ID in the format of subnet-bdoe83fa. `SubnetId` and `VpcId` should be set or ignored simultaneously
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// VPC ID in the format of vpc-dsp338hz. `SubnetId` and `VpcId` should be set or ignored simultaneously
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Length of purchase of instance. The default value is 1, indicating one month. The value cannot exceed 48
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// Whether to automatically use voucher. 0: no, 1: yes. Default value: no
	AutoVoucher *int64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// Array of voucher IDs (currently, only one voucher can be used per order)
	VoucherIds []*string `json:"VoucherIds,omitnil,omitempty" name:"VoucherIds"`

	// SQL Server version. Valid values: `2008R2` (SQL Server 2008 R2 Enterprise), `2012SP3` (SQL Server 2012 Enterprise), `201202` (SQL Server 2012 Standard), `2014SP2` (SQL Server 2014 Enterprise), 201402 (SQL Server 2014 Standard), `2016SP1` (SQL Server 2016 Enterprise), `201602` (SQL Server 2016 Standard), `2017` (SQL Server 2017 Enterprise), `201702` (SQL Server 2017 Standard), `2019` (SQL Server 2019 Enterprise), `201902` (SQL Server 2019 Standard). Default value: `2008R2`. The available version varies by region, and you can pull the version information by calling the `DescribeProductConfig` API.
	DBVersion *string `json:"DBVersion,omitnil,omitempty" name:"DBVersion"`

	// Auto-renewal flag. 0: normal renewal, 1: auto-renewal. Default value: 1.
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// Security group list, which contains security group IDs in the format of sg-xxx.
	SecurityGroupList []*string `json:"SecurityGroupList,omitnil,omitempty" name:"SecurityGroupList"`

	// Configuration of the maintenance window, which specifies the day of the week when maintenance can be performed. Valid values: 1 (Monday), 2 (Tuesday), 3 (Wednesday), 4 (Thursday), 5 (Friday), 6 (Saturday), 7 (Sunday).
	Weekly []*int64 `json:"Weekly,omitnil,omitempty" name:"Weekly"`

	// Configuration of the maintenance window, which specifies the start time of daily maintenance.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Configuration of the maintenance window, which specifies the maintenance duration in hours.
	Span *int64 `json:"Span,omitnil,omitempty" name:"Span"`

	// The type of purchased high-availability instance. Valid values: DUAL (dual-server high availability), CLUSTER (cluster). Default value: DUAL.
	HAType *string `json:"HAType,omitnil,omitempty" name:"HAType"`

	// Whether to deploy across availability zones. Default value: false.
	MultiZones *bool `json:"MultiZones,omitnil,omitempty" name:"MultiZones"`

	// Tags associated with the instances to be created
	ResourceTags []*ResourceTag `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`

	// Collation of system character sets. Default value: `Chinese_PRC_CI_AS`.
	Collation *string `json:"Collation,omitnil,omitempty" name:"Collation"`

	// System time zone. Default value: `China Standard Time`.
	TimeZone *string `json:"TimeZone,omitnil,omitempty" name:"TimeZone"`
}

type CreateDBInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Instance AZ, such as ap-guangzhou-1 (Guangzhou Zone 1). Purchasable AZs for an instance can be obtained through the `DescribeZones` API
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Instance memory size in GB
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// Instance storage capacity in GB
	Storage *int64 `json:"Storage,omitnil,omitempty" name:"Storage"`

	// Billing mode. Valid value: POSTPAID (pay-as-you-go).
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// Project ID
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Number of instances purchased this time. Default value: 1. Maximum value: 10
	GoodsNum *int64 `json:"GoodsNum,omitnil,omitempty" name:"GoodsNum"`

	// VPC subnet ID in the format of subnet-bdoe83fa. `SubnetId` and `VpcId` should be set or ignored simultaneously
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// VPC ID in the format of vpc-dsp338hz. `SubnetId` and `VpcId` should be set or ignored simultaneously
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Length of purchase of instance. The default value is 1, indicating one month. The value cannot exceed 48
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// Whether to automatically use voucher. 0: no, 1: yes. Default value: no
	AutoVoucher *int64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// Array of voucher IDs (currently, only one voucher can be used per order)
	VoucherIds []*string `json:"VoucherIds,omitnil,omitempty" name:"VoucherIds"`

	// SQL Server version. Valid values: `2008R2` (SQL Server 2008 R2 Enterprise), `2012SP3` (SQL Server 2012 Enterprise), `201202` (SQL Server 2012 Standard), `2014SP2` (SQL Server 2014 Enterprise), 201402 (SQL Server 2014 Standard), `2016SP1` (SQL Server 2016 Enterprise), `201602` (SQL Server 2016 Standard), `2017` (SQL Server 2017 Enterprise), `201702` (SQL Server 2017 Standard), `2019` (SQL Server 2019 Enterprise), `201902` (SQL Server 2019 Standard). Default value: `2008R2`. The available version varies by region, and you can pull the version information by calling the `DescribeProductConfig` API.
	DBVersion *string `json:"DBVersion,omitnil,omitempty" name:"DBVersion"`

	// Auto-renewal flag. 0: normal renewal, 1: auto-renewal. Default value: 1.
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// Security group list, which contains security group IDs in the format of sg-xxx.
	SecurityGroupList []*string `json:"SecurityGroupList,omitnil,omitempty" name:"SecurityGroupList"`

	// Configuration of the maintenance window, which specifies the day of the week when maintenance can be performed. Valid values: 1 (Monday), 2 (Tuesday), 3 (Wednesday), 4 (Thursday), 5 (Friday), 6 (Saturday), 7 (Sunday).
	Weekly []*int64 `json:"Weekly,omitnil,omitempty" name:"Weekly"`

	// Configuration of the maintenance window, which specifies the start time of daily maintenance.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Configuration of the maintenance window, which specifies the maintenance duration in hours.
	Span *int64 `json:"Span,omitnil,omitempty" name:"Span"`

	// The type of purchased high-availability instance. Valid values: DUAL (dual-server high availability), CLUSTER (cluster). Default value: DUAL.
	HAType *string `json:"HAType,omitnil,omitempty" name:"HAType"`

	// Whether to deploy across availability zones. Default value: false.
	MultiZones *bool `json:"MultiZones,omitnil,omitempty" name:"MultiZones"`

	// Tags associated with the instances to be created
	ResourceTags []*ResourceTag `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`

	// Collation of system character sets. Default value: `Chinese_PRC_CI_AS`.
	Collation *string `json:"Collation,omitnil,omitempty" name:"Collation"`

	// System time zone. Default value: `China Standard Time`.
	TimeZone *string `json:"TimeZone,omitnil,omitempty" name:"TimeZone"`
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
	delete(f, "Collation")
	delete(f, "TimeZone")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDBInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDBInstancesResponseParams struct {
	// Order name
	DealName *string `json:"DealName,omitnil,omitempty" name:"DealName"`

	// Order name array
	DealNames []*string `json:"DealNames,omitnil,omitempty" name:"DealNames"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *CreateDBInstancesResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateDBRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Database creation information
	DBs []*DBCreateInfo `json:"DBs,omitnil,omitempty" name:"DBs"`
}

type CreateDBRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Database creation information
	DBs []*DBCreateInfo `json:"DBs,omitnil,omitempty" name:"DBs"`
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

// Predefined struct for user
type CreateDBResponseParams struct {
	// Task flow ID
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDBResponse struct {
	*tchttp.BaseResponse
	Response *CreateDBResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateIncrementalMigrationRequestParams struct {
	// ID of imported target instance
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Backup import task ID, which is returned through the API CreateBackupMigration.
	BackupMigrationId *string `json:"BackupMigrationId,omitnil,omitempty" name:"BackupMigrationId"`

	// Incremental backup file. If the UploadType of a full backup file is COS_URL, fill in URL here. If the UploadType is COS_UPLOAD, fill in the name of the backup file here. Only 1 backup file is supported, but a backup file can involve multiple databases.
	BackupFiles []*string `json:"BackupFiles,omitnil,omitempty" name:"BackupFiles"`

	// Whether restoration is required. No: not required. Yes: required. Not required by default.
	IsRecovery *string `json:"IsRecovery,omitnil,omitempty" name:"IsRecovery"`
}

type CreateIncrementalMigrationRequest struct {
	*tchttp.BaseRequest
	
	// ID of imported target instance
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Backup import task ID, which is returned through the API CreateBackupMigration.
	BackupMigrationId *string `json:"BackupMigrationId,omitnil,omitempty" name:"BackupMigrationId"`

	// Incremental backup file. If the UploadType of a full backup file is COS_URL, fill in URL here. If the UploadType is COS_UPLOAD, fill in the name of the backup file here. Only 1 backup file is supported, but a backup file can involve multiple databases.
	BackupFiles []*string `json:"BackupFiles,omitnil,omitempty" name:"BackupFiles"`

	// Whether restoration is required. No: not required. Yes: required. Not required by default.
	IsRecovery *string `json:"IsRecovery,omitnil,omitempty" name:"IsRecovery"`
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

// Predefined struct for user
type CreateIncrementalMigrationResponseParams struct {
	// ID of an incremental backup import task
	IncrementalMigrationId *string `json:"IncrementalMigrationId,omitnil,omitempty" name:"IncrementalMigrationId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateIncrementalMigrationResponse struct {
	*tchttp.BaseResponse
	Response *CreateIncrementalMigrationResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateMigrationRequestParams struct {
	// Migration task name
	MigrateName *string `json:"MigrateName,omitnil,omitempty" name:"MigrateName"`

	// Migration type (1: structure migration, 2: data migration, 3: incremental sync)
	MigrateType *uint64 `json:"MigrateType,omitnil,omitempty" name:"MigrateType"`

	// Migration source type. 1: TencentDB for SQL Server, 2: CVM-based self-created SQL Server database; 3: SQL Server backup restoration, 4: SQL Server backup restoration (in COS mode)
	SourceType *uint64 `json:"SourceType,omitnil,omitempty" name:"SourceType"`

	// Migration source
	Source *MigrateSource `json:"Source,omitnil,omitempty" name:"Source"`

	// Migration target
	Target *MigrateTarget `json:"Target,omitnil,omitempty" name:"Target"`

	// Database objects to be migrated. This parameter is not used for offline migration (SourceType=4 or SourceType=5)
	MigrateDBSet []*MigrateDB `json:"MigrateDBSet,omitnil,omitempty" name:"MigrateDBSet"`

	// Restore and rename the databases listed in `ReNameRestoreDatabase`. If this parameter is left empty, all restored databases will be renamed in the default format. This parameter takes effect only when `SourceType=5`.
	RenameRestore []*RenameRestoreDatabase `json:"RenameRestore,omitnil,omitempty" name:"RenameRestore"`
}

type CreateMigrationRequest struct {
	*tchttp.BaseRequest
	
	// Migration task name
	MigrateName *string `json:"MigrateName,omitnil,omitempty" name:"MigrateName"`

	// Migration type (1: structure migration, 2: data migration, 3: incremental sync)
	MigrateType *uint64 `json:"MigrateType,omitnil,omitempty" name:"MigrateType"`

	// Migration source type. 1: TencentDB for SQL Server, 2: CVM-based self-created SQL Server database; 3: SQL Server backup restoration, 4: SQL Server backup restoration (in COS mode)
	SourceType *uint64 `json:"SourceType,omitnil,omitempty" name:"SourceType"`

	// Migration source
	Source *MigrateSource `json:"Source,omitnil,omitempty" name:"Source"`

	// Migration target
	Target *MigrateTarget `json:"Target,omitnil,omitempty" name:"Target"`

	// Database objects to be migrated. This parameter is not used for offline migration (SourceType=4 or SourceType=5)
	MigrateDBSet []*MigrateDB `json:"MigrateDBSet,omitnil,omitempty" name:"MigrateDBSet"`

	// Restore and rename the databases listed in `ReNameRestoreDatabase`. If this parameter is left empty, all restored databases will be renamed in the default format. This parameter takes effect only when `SourceType=5`.
	RenameRestore []*RenameRestoreDatabase `json:"RenameRestore,omitnil,omitempty" name:"RenameRestore"`
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

// Predefined struct for user
type CreateMigrationResponseParams struct {
	// Migration task ID
	MigrateId *int64 `json:"MigrateId,omitnil,omitempty" name:"MigrateId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateMigrationResponse struct {
	*tchttp.BaseResponse
	Response *CreateMigrationResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateReadOnlyDBInstancesRequestParams struct {

	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`


	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`


	ReadOnlyGroupType *int64 `json:"ReadOnlyGroupType,omitnil,omitempty" name:"ReadOnlyGroupType"`


	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`


	Storage *int64 `json:"Storage,omitnil,omitempty" name:"Storage"`


	ReadOnlyGroupForcedUpgrade *int64 `json:"ReadOnlyGroupForcedUpgrade,omitnil,omitempty" name:"ReadOnlyGroupForcedUpgrade"`


	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitnil,omitempty" name:"ReadOnlyGroupId"`


	ReadOnlyGroupName *string `json:"ReadOnlyGroupName,omitnil,omitempty" name:"ReadOnlyGroupName"`


	ReadOnlyGroupIsOfflineDelay *int64 `json:"ReadOnlyGroupIsOfflineDelay,omitnil,omitempty" name:"ReadOnlyGroupIsOfflineDelay"`


	ReadOnlyGroupMaxDelayTime *int64 `json:"ReadOnlyGroupMaxDelayTime,omitnil,omitempty" name:"ReadOnlyGroupMaxDelayTime"`


	ReadOnlyGroupMinInGroup *int64 `json:"ReadOnlyGroupMinInGroup,omitnil,omitempty" name:"ReadOnlyGroupMinInGroup"`


	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`


	GoodsNum *int64 `json:"GoodsNum,omitnil,omitempty" name:"GoodsNum"`


	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`


	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`


	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`


	SecurityGroupList []*string `json:"SecurityGroupList,omitnil,omitempty" name:"SecurityGroupList"`


	AutoVoucher *int64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`


	VoucherIds []*string `json:"VoucherIds,omitnil,omitempty" name:"VoucherIds"`


	ResourceTags []*ResourceTag `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`


	Collation *string `json:"Collation,omitnil,omitempty" name:"Collation"`


	TimeZone *string `json:"TimeZone,omitnil,omitempty" name:"TimeZone"`
}

type CreateReadOnlyDBInstancesRequest struct {
	*tchttp.BaseRequest
	
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	ReadOnlyGroupType *int64 `json:"ReadOnlyGroupType,omitnil,omitempty" name:"ReadOnlyGroupType"`

	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	Storage *int64 `json:"Storage,omitnil,omitempty" name:"Storage"`

	ReadOnlyGroupForcedUpgrade *int64 `json:"ReadOnlyGroupForcedUpgrade,omitnil,omitempty" name:"ReadOnlyGroupForcedUpgrade"`

	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitnil,omitempty" name:"ReadOnlyGroupId"`

	ReadOnlyGroupName *string `json:"ReadOnlyGroupName,omitnil,omitempty" name:"ReadOnlyGroupName"`

	ReadOnlyGroupIsOfflineDelay *int64 `json:"ReadOnlyGroupIsOfflineDelay,omitnil,omitempty" name:"ReadOnlyGroupIsOfflineDelay"`

	ReadOnlyGroupMaxDelayTime *int64 `json:"ReadOnlyGroupMaxDelayTime,omitnil,omitempty" name:"ReadOnlyGroupMaxDelayTime"`

	ReadOnlyGroupMinInGroup *int64 `json:"ReadOnlyGroupMinInGroup,omitnil,omitempty" name:"ReadOnlyGroupMinInGroup"`

	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	GoodsNum *int64 `json:"GoodsNum,omitnil,omitempty" name:"GoodsNum"`

	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	SecurityGroupList []*string `json:"SecurityGroupList,omitnil,omitempty" name:"SecurityGroupList"`

	AutoVoucher *int64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	VoucherIds []*string `json:"VoucherIds,omitnil,omitempty" name:"VoucherIds"`

	ResourceTags []*ResourceTag `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`

	Collation *string `json:"Collation,omitnil,omitempty" name:"Collation"`

	TimeZone *string `json:"TimeZone,omitnil,omitempty" name:"TimeZone"`
}

func (r *CreateReadOnlyDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateReadOnlyDBInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Zone")
	delete(f, "ReadOnlyGroupType")
	delete(f, "Memory")
	delete(f, "Storage")
	delete(f, "ReadOnlyGroupForcedUpgrade")
	delete(f, "ReadOnlyGroupId")
	delete(f, "ReadOnlyGroupName")
	delete(f, "ReadOnlyGroupIsOfflineDelay")
	delete(f, "ReadOnlyGroupMaxDelayTime")
	delete(f, "ReadOnlyGroupMinInGroup")
	delete(f, "InstanceChargeType")
	delete(f, "GoodsNum")
	delete(f, "SubnetId")
	delete(f, "VpcId")
	delete(f, "Period")
	delete(f, "SecurityGroupList")
	delete(f, "AutoVoucher")
	delete(f, "VoucherIds")
	delete(f, "ResourceTags")
	delete(f, "Collation")
	delete(f, "TimeZone")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateReadOnlyDBInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateReadOnlyDBInstancesResponseParams struct {

	DealNames []*string `json:"DealNames,omitnil,omitempty" name:"DealNames"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateReadOnlyDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *CreateReadOnlyDBInstancesResponseParams `json:"Response"`
}

func (r *CreateReadOnlyDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateReadOnlyDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CrossBackupAddr struct {
	// The target region of cross-region backup
	CrossRegion *string `json:"CrossRegion,omitnil,omitempty" name:"CrossRegion"`

	// The address used to download the cross-region backup over a private network
	CrossInternalAddr *string `json:"CrossInternalAddr,omitnil,omitempty" name:"CrossInternalAddr"`

	// The address used to download the cross-region backup over a public network
	CrossExternalAddr *string `json:"CrossExternalAddr,omitnil,omitempty" name:"CrossExternalAddr"`
}

type CrossRegionStatus struct {
	// The target region of cross-region backup
	CrossRegion *string `json:"CrossRegion,omitnil,omitempty" name:"CrossRegion"`

	// Synchronization status of cross-region backup. Valid values: `0` (creating), `1` (succeeded), `2`: (failed), `4` (syncing)
	CrossStatus *int64 `json:"CrossStatus,omitnil,omitempty" name:"CrossStatus"`
}

type DBCreateInfo struct {
	// Database name
	DBName *string `json:"DBName,omitnil,omitempty" name:"DBName"`

	// Character set, which can be queried by the `DescribeDBCharsets` API. Default value: `Chinese_PRC_CI_AS`.
	Charset *string `json:"Charset,omitnil,omitempty" name:"Charset"`

	// Database account permission information
	Accounts []*AccountPrivilege `json:"Accounts,omitnil,omitempty" name:"Accounts"`

	// Remarks
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

type DBDetail struct {
	// Database name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Character set
	Charset *string `json:"Charset,omitnil,omitempty" name:"Charset"`

	// Remarks
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// Database creation time
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Database status. 1: creating, 2: running, 3: modifying, -1: dropping
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Database account permission information
	Accounts []*AccountPrivilege `json:"Accounts,omitnil,omitempty" name:"Accounts"`

	// Internal status. ONLINE: running
	InternalStatus *string `json:"InternalStatus,omitnil,omitempty" name:"InternalStatus"`

	// TDE status. Valid values: `enable` (enabled), `disable` (disabled).
	Encryption *string `json:"Encryption,omitnil,omitempty" name:"Encryption"`
}

type DBInstance struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Instance name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Project ID of instance
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Instance region ID
	RegionId *int64 `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// Instance AZ ID
	ZoneId *int64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Instance VPC ID, which will be 0 if the basic network is used
	VpcId *int64 `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Instance VPC subnet ID, which will be 0 if the basic network is used
	SubnetId *int64 `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Instance status. Valid values: <li>1: creating</li> <li>2: running</li> <li>3: instance operations restricted (due to the ongoing primary-replica switch)</li> <li>4: isolated</li> <li>5: repossessing</li> <li>6: repossessed</li> <li>7: running tasks (such as backup and rollback tasks)</li> <li>8: eliminated</li> <li>9: expanding capacity</li> <li>10: migrating</li> <li>11: read-only</li> <li>12: restarting</li>  <li>13: modifying configuration and waiting for switch</li> <li>14: implementing pub/sub</li> <li>15: modifying pub/sub configuration</li> <li>16: modifying configuration and switching</li> <li>17: creating read-only instances</li>
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Instance access IP
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// Instance access port
	Vport *int64 `json:"Vport,omitnil,omitempty" name:"Vport"`

	// Instance creation time
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Instance update time
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// Instance billing start time
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Instance billing end time
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Instance isolation time
	IsolateTime *string `json:"IsolateTime,omitnil,omitempty" name:"IsolateTime"`

	// Instance memory size in GB
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// Used storage capacity of instance in GB
	UsedStorage *int64 `json:"UsedStorage,omitnil,omitempty" name:"UsedStorage"`

	// Instance storage capacity in GB
	Storage *int64 `json:"Storage,omitnil,omitempty" name:"Storage"`

	// Instance version
	VersionName *string `json:"VersionName,omitnil,omitempty" name:"VersionName"`

	// Instance renewal flag
	RenewFlag *int64 `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// High-availability instance type. Valid values: 1 (dual-server high-availability), 2 (standalone), 3 (multi-AZ), 4 (multi-AZ cluster), 5 (cluster), 9 (private consumption)
	Model *int64 `json:"Model,omitnil,omitempty" name:"Model"`

	// Instance region name, such as ap-guangzhou
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Instance AZ name, such as ap-guangzhou-1
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Backup time point
	BackupTime *string `json:"BackupTime,omitnil,omitempty" name:"BackupTime"`

	// Instance billing mode. 0: pay-as-you-go
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// Instance UID
	Uid *string `json:"Uid,omitnil,omitempty" name:"Uid"`

	// Number of CPU cores of instance
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// Instance version code
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// Instance type. Valid values: `TS85` (physical machine, local SSD), `Z3` (early version of physical machine, local SSD), `CLOUD_BASIC` (virtual machine, HDD cloud disk), `CLOUD_PREMIUM` (virtual machine, premium cloud disk), `CLOUD_SSD` (virtual machine, SSD), `CLOUD_HSSD` (virtual machine, enhanced SSD), `CLOUD_TSSD` (virtual machine, ulTra SSD), `CLOUD_BSSD` virtual machine, balanced SSD).
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Billing ID
	Pid *int64 `json:"Pid,omitnil,omitempty" name:"Pid"`

	// Unique string-type ID of instance VPC in the format of `vpc-xxx`, which is an empty string if the basic network is used
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// Unique string-type ID of instance subnet in the format of `subnet-xxx`, which is an empty string if the basic network is used
	UniqSubnetId *string `json:"UniqSubnetId,omitnil,omitempty" name:"UniqSubnetId"`

	// Note: This field may return null, indicating that no valid values can be obtained.
	IsolateOperator *string `json:"IsolateOperator,omitnil,omitempty" name:"IsolateOperator"`

	// Note: This field may return null, indicating that no valid values can be obtained.
	SubFlag *string `json:"SubFlag,omitnil,omitempty" name:"SubFlag"`

	// Note: This field may return null, indicating that no valid values can be obtained.
	ROFlag *string `json:"ROFlag,omitnil,omitempty" name:"ROFlag"`

	// Note: This field may return null, indicating that no valid values can be obtained.
	HAFlag *string `json:"HAFlag,omitnil,omitempty" name:"HAFlag"`

	// Note: This field may return null, indicating that no valid values can be obtained.
	ResourceTags []*ResourceTag `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`

	// Note: This field may return null, indicating that no valid values can be obtained.
	BackupModel *string `json:"BackupModel,omitnil,omitempty" name:"BackupModel"`

	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceNote *string `json:"InstanceNote,omitnil,omitempty" name:"InstanceNote"`

	// Backup cycle
	BackupCycle []*int64 `json:"BackupCycle,omitnil,omitempty" name:"BackupCycle"`

	// Backup cycle type. Valid values: `daily`, `weekly`, `monthly`.
	BackupCycleType *string `json:"BackupCycleType,omitnil,omitempty" name:"BackupCycleType"`

	// Data (log) backup retention period
	BackupSaveDays *int64 `json:"BackupSaveDays,omitnil,omitempty" name:"BackupSaveDays"`

	// Instance type. Valid values: `HA` (high-availability), `RO` (read-only), `SI` (basic edition), `BI` (business intelligence service).
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// The target region of cross-region backup. If this parameter left empty, it indicates that cross-region backup is disabled.
	CrossRegions []*string `json:"CrossRegions,omitnil,omitempty" name:"CrossRegions"`

	// Cross-region backup status. Valid values: `enable` (enabled), `disable` (disabed)
	CrossBackupEnabled *string `json:"CrossBackupEnabled,omitnil,omitempty" name:"CrossBackupEnabled"`

	// The retention period of cross-region backup. Default value: 7 days
	CrossBackupSaveDays *uint64 `json:"CrossBackupSaveDays,omitnil,omitempty" name:"CrossBackupSaveDays"`

	// Domain name of the public network address
	DnsPodDomain *string `json:"DnsPodDomain,omitnil,omitempty" name:"DnsPodDomain"`

	// Port number of the public network
	TgwWanVPort *int64 `json:"TgwWanVPort,omitnil,omitempty" name:"TgwWanVPort"`

	// Collation of system character sets. Default value: `Chinese_PRC_CI_AS`.
	Collation *string `json:"Collation,omitnil,omitempty" name:"Collation"`

	// System time zone. Default value: `China Standard Time`.
	TimeZone *string `json:"TimeZone,omitnil,omitempty" name:"TimeZone"`

	// Whether the instance is deployed across AZs
	IsDrZone *bool `json:"IsDrZone,omitnil,omitempty" name:"IsDrZone"`

	// Note: This field may return null, indicating that no valid values can be obtained.
	SlaveZones *SlaveZones `json:"SlaveZones,omitnil,omitempty" name:"SlaveZones"`

	// Note: This field may return null, indicating that no valid values can be obtained.
	Architecture *string `json:"Architecture,omitnil,omitempty" name:"Architecture"`

	// Note: This field may return null, indicating that no valid values can be obtained.
	Style *string `json:"Style,omitnil,omitempty" name:"Style"`
}

type DBPrivilege struct {
	// Database name
	DBName *string `json:"DBName,omitnil,omitempty" name:"DBName"`

	// Database permissions. Valid values: `ReadWrite` (read-write), `ReadOnly` (read-only), `DBOwner` (owner)
	Privilege *string `json:"Privilege,omitnil,omitempty" name:"Privilege"`
}

type DBPrivilegeModifyInfo struct {
	// Database name
	DBName *string `json:"DBName,omitnil,omitempty" name:"DBName"`

	// Permission modification information. Valid values: `ReadWrite` (read-write), `ReadOnly` (read-only), `Delete` (delete the account's permission to this database), `DBOwner` (owner).
	Privilege *string `json:"Privilege,omitnil,omitempty" name:"Privilege"`
}

type DBRemark struct {
	// Database name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Remarks
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

type DBRenameRes struct {
	// Name of the new database
	NewName *string `json:"NewName,omitnil,omitempty" name:"NewName"`

	// Name of the old database
	OldName *string `json:"OldName,omitnil,omitempty" name:"OldName"`
}

type DBTDEEncrypt struct {

	DBName *string `json:"DBName,omitnil,omitempty" name:"DBName"`

	// TDE status. Valid values: `enable` (enabled), `disable` (disabled).
	Encryption *string `json:"Encryption,omitnil,omitempty" name:"Encryption"`
}

type DbNormalDetail struct {
	// Whether it is subscribed. Valid values: `0` (no), `1` (yes)
	IsSubscribed *string `json:"IsSubscribed,omitnil,omitempty" name:"IsSubscribed"`

	// Database collation
	CollationName *string `json:"CollationName,omitnil,omitempty" name:"CollationName"`

	// Whether the cleanup task is enabled to automatically remove old change tracking information when CT is enabled. Valid values: `0` (no), `1` (yes)
	IsAutoCleanupOn *string `json:"IsAutoCleanupOn,omitnil,omitempty" name:"IsAutoCleanupOn"`

	// Whether SQL Server Service Broker is enabled. Valid values: `0` (no), `1` (yes)
	IsBrokerEnabled *string `json:"IsBrokerEnabled,omitnil,omitempty" name:"IsBrokerEnabled"`

	// Whether CDC is enabled. Valid values: `0` (disabled), `1` (enabled)
	IsCdcEnabled *string `json:"IsCdcEnabled,omitnil,omitempty" name:"IsCdcEnabled"`

	// Whether CT is enabled. Valid values: `0` (disabled), `1` (enabled)
	IsDbChainingOn *string `json:"IsDbChainingOn,omitnil,omitempty" name:"IsDbChainingOn"`

	// Whether it is encrypted. Valid values: `0` (no), `1` (yes)
	IsEncrypted *string `json:"IsEncrypted,omitnil,omitempty" name:"IsEncrypted"`

	// Whether full-text indexes are enabled. Valid values: `0` (no), `1` (yes)
	//
	// Deprecated: IsFulltextEnabled is deprecated.
	IsFulltextEnabled *string `json:"IsFulltextEnabled,omitnil,omitempty" name:"IsFulltextEnabled"`

	// Whether it is a mirror database. Valid values: `0` (no), `1` (yes)
	IsMirroring *string `json:"IsMirroring,omitnil,omitempty" name:"IsMirroring"`

	// Whether it is published. Valid values: `0` (no), `1` (yes)
	IsPublished *string `json:"IsPublished,omitnil,omitempty" name:"IsPublished"`

	// Whether snapshots are enabled. Valid values: `0` (no), `1` (yes)
	IsReadCommittedSnapshotOn *string `json:"IsReadCommittedSnapshotOn,omitnil,omitempty" name:"IsReadCommittedSnapshotOn"`

	// Whether it is trustworthy. Valid values: `0` (no), `1` (yes)
	IsTrustworthyOn *string `json:"IsTrustworthyOn,omitnil,omitempty" name:"IsTrustworthyOn"`

	// Mirroring state
	MirroringState *string `json:"MirroringState,omitnil,omitempty" name:"MirroringState"`

	// Database name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Recovery model
	RecoveryModelDesc *string `json:"RecoveryModelDesc,omitnil,omitempty" name:"RecoveryModelDesc"`

	// Retention period (in days) of change tracking information
	RetentionPeriod *string `json:"RetentionPeriod,omitnil,omitempty" name:"RetentionPeriod"`

	// Database status
	StateDesc *string `json:"StateDesc,omitnil,omitempty" name:"StateDesc"`

	// User type
	UserAccessDesc *string `json:"UserAccessDesc,omitnil,omitempty" name:"UserAccessDesc"`

	// Database creation time
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`
}

type DbRollbackTimeInfo struct {
	// Database name
	DBName *string `json:"DBName,omitnil,omitempty" name:"DBName"`

	// Start time of time range available for rollback
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time of time range available for rollback
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type DealInfo struct {
	// Order name
	DealName *string `json:"DealName,omitnil,omitempty" name:"DealName"`

	// Number of items
	Count *uint64 `json:"Count,omitnil,omitempty" name:"Count"`

	// ID of associated flow, which can be used to query the flow execution status
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// This field is required only for an order that creates an instance, indicating the ID of the instance created by the order
	InstanceIdSet []*string `json:"InstanceIdSet,omitnil,omitempty" name:"InstanceIdSet"`

	// Account
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// Instance billing type
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`
}

type DealInstance struct {
	// Instance ID
	InstanceId []*string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Order ID
	DealName *string `json:"DealName,omitnil,omitempty" name:"DealName"`
}

// Predefined struct for user
type DeleteAccountRequestParams struct {
	// Database instance ID in the format of mssql-njj2mtpl
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Array of instance usernames
	UserNames []*string `json:"UserNames,omitnil,omitempty" name:"UserNames"`
}

type DeleteAccountRequest struct {
	*tchttp.BaseRequest
	
	// Database instance ID in the format of mssql-njj2mtpl
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Array of instance usernames
	UserNames []*string `json:"UserNames,omitnil,omitempty" name:"UserNames"`
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

// Predefined struct for user
type DeleteAccountResponseParams struct {
	// Task flow ID
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteAccountResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAccountResponseParams `json:"Response"`
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

// Predefined struct for user
type DeleteBackupMigrationRequestParams struct {
	// Target instance ID, which is returned through the API DescribeBackupMigration.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Backup import task ID, which is returned through the API DescribeBackupMigration.
	BackupMigrationId *string `json:"BackupMigrationId,omitnil,omitempty" name:"BackupMigrationId"`
}

type DeleteBackupMigrationRequest struct {
	*tchttp.BaseRequest
	
	// Target instance ID, which is returned through the API DescribeBackupMigration.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Backup import task ID, which is returned through the API DescribeBackupMigration.
	BackupMigrationId *string `json:"BackupMigrationId,omitnil,omitempty" name:"BackupMigrationId"`
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

// Predefined struct for user
type DeleteBackupMigrationResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteBackupMigrationResponse struct {
	*tchttp.BaseResponse
	Response *DeleteBackupMigrationResponseParams `json:"Response"`
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

// Predefined struct for user
type DeleteBusinessIntelligenceFileRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// File name set
	FileNameSet []*string `json:"FileNameSet,omitnil,omitempty" name:"FileNameSet"`
}

type DeleteBusinessIntelligenceFileRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// File name set
	FileNameSet []*string `json:"FileNameSet,omitnil,omitempty" name:"FileNameSet"`
}

func (r *DeleteBusinessIntelligenceFileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBusinessIntelligenceFileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "FileNameSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteBusinessIntelligenceFileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteBusinessIntelligenceFileResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteBusinessIntelligenceFileResponse struct {
	*tchttp.BaseResponse
	Response *DeleteBusinessIntelligenceFileResponseParams `json:"Response"`
}

func (r *DeleteBusinessIntelligenceFileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBusinessIntelligenceFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDBRequestParams struct {
	// Instance ID in the format of mssql-rljoi3bf
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Array of database names
	Names []*string `json:"Names,omitnil,omitempty" name:"Names"`
}

type DeleteDBRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of mssql-rljoi3bf
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Array of database names
	Names []*string `json:"Names,omitnil,omitempty" name:"Names"`
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

// Predefined struct for user
type DeleteDBResponseParams struct {
	// Task flow ID
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteDBResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDBResponseParams `json:"Response"`
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

// Predefined struct for user
type DeleteIncrementalMigrationRequestParams struct {
	// Target instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Backup import task ID, which is returned through the `CreateBackupMigration` API
	BackupMigrationId *string `json:"BackupMigrationId,omitnil,omitempty" name:"BackupMigrationId"`

	// Incremental backup import task ID, which is returned through the `CreateIncrementalMigration` API
	IncrementalMigrationId *string `json:"IncrementalMigrationId,omitnil,omitempty" name:"IncrementalMigrationId"`
}

type DeleteIncrementalMigrationRequest struct {
	*tchttp.BaseRequest
	
	// Target instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Backup import task ID, which is returned through the `CreateBackupMigration` API
	BackupMigrationId *string `json:"BackupMigrationId,omitnil,omitempty" name:"BackupMigrationId"`

	// Incremental backup import task ID, which is returned through the `CreateIncrementalMigration` API
	IncrementalMigrationId *string `json:"IncrementalMigrationId,omitnil,omitempty" name:"IncrementalMigrationId"`
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

// Predefined struct for user
type DeleteIncrementalMigrationResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteIncrementalMigrationResponse struct {
	*tchttp.BaseResponse
	Response *DeleteIncrementalMigrationResponseParams `json:"Response"`
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

// Predefined struct for user
type DeleteMigrationRequestParams struct {
	// Migration task ID
	MigrateId *uint64 `json:"MigrateId,omitnil,omitempty" name:"MigrateId"`
}

type DeleteMigrationRequest struct {
	*tchttp.BaseRequest
	
	// Migration task ID
	MigrateId *uint64 `json:"MigrateId,omitnil,omitempty" name:"MigrateId"`
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

// Predefined struct for user
type DeleteMigrationResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteMigrationResponse struct {
	*tchttp.BaseResponse
	Response *DeleteMigrationResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeAccountsRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Number of results per page. Value range: 1-100. Default value: 20
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Page number. Default value: 0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Account ID
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Sorting by `createTime`, `updateTime`, or `passTime`. Default value: `createTime` (desc).
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Sorting rule. Valid values: `desc` (descending order), `asc` (ascending order). Default value: `desc`.
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`
}

type DescribeAccountsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Number of results per page. Value range: 1-100. Default value: 20
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Page number. Default value: 0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Account ID
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Sorting by `createTime`, `updateTime`, or `passTime`. Default value: `createTime` (desc).
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Sorting rule. Valid values: `desc` (descending order), `asc` (ascending order). Default value: `desc`.
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`
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
	delete(f, "Name")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccountsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccountsResponseParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Account information list
	Accounts []*AccountDetail `json:"Accounts,omitnil,omitempty" name:"Accounts"`

	// Total number
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAccountsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAccountsResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeBackupCommandRequestParams struct {
	// Backup file type. Full: full backup. FULL_LOG: full backup which needs log increments. FULL_DIFF: full backup which needs differential increments. LOG: log backup. DIFF: differential backup.
	BackupFileType *string `json:"BackupFileType,omitnil,omitempty" name:"BackupFileType"`

	// Database name
	DataBaseName *string `json:"DataBaseName,omitnil,omitempty" name:"DataBaseName"`

	// Whether restoration is required. No: not required. Yes: required.
	IsRecovery *string `json:"IsRecovery,omitnil,omitempty" name:"IsRecovery"`

	// Storage path of backup files. If this parameter is left empty, the default storage path will be D:\\.
	LocalPath *string `json:"LocalPath,omitnil,omitempty" name:"LocalPath"`
}

type DescribeBackupCommandRequest struct {
	*tchttp.BaseRequest
	
	// Backup file type. Full: full backup. FULL_LOG: full backup which needs log increments. FULL_DIFF: full backup which needs differential increments. LOG: log backup. DIFF: differential backup.
	BackupFileType *string `json:"BackupFileType,omitnil,omitempty" name:"BackupFileType"`

	// Database name
	DataBaseName *string `json:"DataBaseName,omitnil,omitempty" name:"DataBaseName"`

	// Whether restoration is required. No: not required. Yes: required.
	IsRecovery *string `json:"IsRecovery,omitnil,omitempty" name:"IsRecovery"`

	// Storage path of backup files. If this parameter is left empty, the default storage path will be D:\\.
	LocalPath *string `json:"LocalPath,omitnil,omitempty" name:"LocalPath"`
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

// Predefined struct for user
type DescribeBackupCommandResponseParams struct {
	// Create a backup command
	Command *string `json:"Command,omitnil,omitempty" name:"Command"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBackupCommandResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBackupCommandResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeBackupFilesRequestParams struct {
	// Instance ID in the format of mssql-njj2mtpl
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Group ID of unarchived backup files, which can be obtained by the `DescribeBackups` API (Querying archived backup record is not supported).
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// Number of entries to be returned per page. Value range: 1-100. Default value: `20`
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Page number. Default value: `0`
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Filter backups by database name. If the parameter is left empty, this filter criterion will not take effect.
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// List items sorting by backup size. Valid values: `desc`(descending order), `asc` (ascending order). Default value: `desc`.
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`
}

type DescribeBackupFilesRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of mssql-njj2mtpl
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Group ID of unarchived backup files, which can be obtained by the `DescribeBackups` API (Querying archived backup record is not supported).
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// Number of entries to be returned per page. Value range: 1-100. Default value: `20`
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Page number. Default value: `0`
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Filter backups by database name. If the parameter is left empty, this filter criterion will not take effect.
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// List items sorting by backup size. Valid values: `desc`(descending order), `asc` (ascending order). Default value: `desc`.
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`
}

func (r *DescribeBackupFilesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupFilesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "GroupId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "DatabaseName")
	delete(f, "OrderBy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackupFilesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupFilesResponseParams struct {
	// Total number of backups
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List of backup file details
	BackupFiles []*BackupFile `json:"BackupFiles,omitnil,omitempty" name:"BackupFiles"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBackupFilesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBackupFilesResponseParams `json:"Response"`
}

func (r *DescribeBackupFilesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupFilesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupMigrationRequestParams struct {
	// ID of imported target instance
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Backup import task ID, which is returned through the API CreateBackupMigration.
	BackupMigrationId *string `json:"BackupMigrationId,omitnil,omitempty" name:"BackupMigrationId"`

	// Import task name
	MigrationName *string `json:"MigrationName,omitnil,omitempty" name:"MigrationName"`

	// Backup file name
	BackupFileName *string `json:"BackupFileName,omitnil,omitempty" name:"BackupFileName"`

	// Status set of import tasks
	StatusSet []*int64 `json:"StatusSet,omitnil,omitempty" name:"StatusSet"`

	// Import task restoration type: FULL,FULL_LOG,FULL_DIFF
	RecoveryType *string `json:"RecoveryType,omitnil,omitempty" name:"RecoveryType"`

	// COS_URL: the backup is stored in user’s Cloud Object Storage, with URL provided. COS_UPLOAD: the backup is stored in the application’s Cloud Object Storage and needs to be uploaded by the user.
	UploadType *string `json:"UploadType,omitnil,omitempty" name:"UploadType"`

	// The maximum number of results returned per page. Default value: `100`.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Page number. Default value: `0`.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Sort by field. Valid values: `name`, `createTime`, `startTime`, `endTime`. By default, the results returned are sorted by `createTime` in the ascending order.
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Sorting order which is valid only when `OrderBy` is specified. Valid values: `asc` (ascending), `desc` (descending). Default value: `asc`.
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`
}

type DescribeBackupMigrationRequest struct {
	*tchttp.BaseRequest
	
	// ID of imported target instance
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Backup import task ID, which is returned through the API CreateBackupMigration.
	BackupMigrationId *string `json:"BackupMigrationId,omitnil,omitempty" name:"BackupMigrationId"`

	// Import task name
	MigrationName *string `json:"MigrationName,omitnil,omitempty" name:"MigrationName"`

	// Backup file name
	BackupFileName *string `json:"BackupFileName,omitnil,omitempty" name:"BackupFileName"`

	// Status set of import tasks
	StatusSet []*int64 `json:"StatusSet,omitnil,omitempty" name:"StatusSet"`

	// Import task restoration type: FULL,FULL_LOG,FULL_DIFF
	RecoveryType *string `json:"RecoveryType,omitnil,omitempty" name:"RecoveryType"`

	// COS_URL: the backup is stored in user’s Cloud Object Storage, with URL provided. COS_UPLOAD: the backup is stored in the application’s Cloud Object Storage and needs to be uploaded by the user.
	UploadType *string `json:"UploadType,omitnil,omitempty" name:"UploadType"`

	// The maximum number of results returned per page. Default value: `100`.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Page number. Default value: `0`.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Sort by field. Valid values: `name`, `createTime`, `startTime`, `endTime`. By default, the results returned are sorted by `createTime` in the ascending order.
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Sorting order which is valid only when `OrderBy` is specified. Valid values: `asc` (ascending), `desc` (descending). Default value: `asc`.
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`
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

// Predefined struct for user
type DescribeBackupMigrationResponseParams struct {
	// Total number of tasks
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Migration task set
	BackupMigrationSet []*Migration `json:"BackupMigrationSet,omitnil,omitempty" name:"BackupMigrationSet"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBackupMigrationResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBackupMigrationResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeBackupUploadSizeRequestParams struct {
	// ID of imported target instance
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Backup import task ID, which is returned through the API CreateBackupMigration
	BackupMigrationId *string `json:"BackupMigrationId,omitnil,omitempty" name:"BackupMigrationId"`

	// Incremental import task ID
	IncrementalMigrationId *string `json:"IncrementalMigrationId,omitnil,omitempty" name:"IncrementalMigrationId"`
}

type DescribeBackupUploadSizeRequest struct {
	*tchttp.BaseRequest
	
	// ID of imported target instance
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Backup import task ID, which is returned through the API CreateBackupMigration
	BackupMigrationId *string `json:"BackupMigrationId,omitnil,omitempty" name:"BackupMigrationId"`

	// Incremental import task ID
	IncrementalMigrationId *string `json:"IncrementalMigrationId,omitnil,omitempty" name:"IncrementalMigrationId"`
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

// Predefined struct for user
type DescribeBackupUploadSizeResponseParams struct {
	// Information of uploaded backups
	CosUploadBackupFileSet []*CosUploadBackupFile `json:"CosUploadBackupFileSet,omitnil,omitempty" name:"CosUploadBackupFileSet"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBackupUploadSizeResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBackupUploadSizeResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeBackupsRequestParams struct {
	// Start name (yyyy-MM-dd HH:mm:ss)
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time (yyyy-MM-dd HH:mm:ss)
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Instance ID in the format of mssql-njj2mtpl
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Number of results per page. Value range: 1-100. Default value: 20
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Page number. Default value: 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Filter by backup name. If this parameter is left empty, backup name will not be used in filtering.
	BackupName *string `json:"BackupName,omitnil,omitempty" name:"BackupName"`

	// Filter by backup policy. Valid values: 0 (instance backup), 1 (multi-database backup). If this parameter is left empty, backup policy will not be used in filtering.
	Strategy *int64 `json:"Strategy,omitnil,omitempty" name:"Strategy"`

	// Filter by backup mode. Valid values: `0` (scheduled backup); `1` (manual backup); `2` (archive backup). Default value: `2`.
	BackupWay *int64 `json:"BackupWay,omitnil,omitempty" name:"BackupWay"`

	// Filter by backup ID. If this parameter is left empty, backup ID will not be used in filtering.
	BackupId *uint64 `json:"BackupId,omitnil,omitempty" name:"BackupId"`

	// Filter backups by the database name. If the parameter is left empty, this filter criteria will not take effect.
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// Whether to group backup files by backup task. Valid value: `0` (no), `1` (yes). Default value: `0`. This parameter is valid only for unarchived backup files.
	Group *int64 `json:"Group,omitnil,omitempty" name:"Group"`

	// Backup type. Valid values: `1` (data backup), `2` (log backup). Default value: `1`.
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// Filter by backup file format. Valid values: `pkg` (archive file), `single` (Unarchived files).
	BackupFormat *string `json:"BackupFormat,omitnil,omitempty" name:"BackupFormat"`
}

type DescribeBackupsRequest struct {
	*tchttp.BaseRequest
	
	// Start name (yyyy-MM-dd HH:mm:ss)
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time (yyyy-MM-dd HH:mm:ss)
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Instance ID in the format of mssql-njj2mtpl
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Number of results per page. Value range: 1-100. Default value: 20
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Page number. Default value: 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Filter by backup name. If this parameter is left empty, backup name will not be used in filtering.
	BackupName *string `json:"BackupName,omitnil,omitempty" name:"BackupName"`

	// Filter by backup policy. Valid values: 0 (instance backup), 1 (multi-database backup). If this parameter is left empty, backup policy will not be used in filtering.
	Strategy *int64 `json:"Strategy,omitnil,omitempty" name:"Strategy"`

	// Filter by backup mode. Valid values: `0` (scheduled backup); `1` (manual backup); `2` (archive backup). Default value: `2`.
	BackupWay *int64 `json:"BackupWay,omitnil,omitempty" name:"BackupWay"`

	// Filter by backup ID. If this parameter is left empty, backup ID will not be used in filtering.
	BackupId *uint64 `json:"BackupId,omitnil,omitempty" name:"BackupId"`

	// Filter backups by the database name. If the parameter is left empty, this filter criteria will not take effect.
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// Whether to group backup files by backup task. Valid value: `0` (no), `1` (yes). Default value: `0`. This parameter is valid only for unarchived backup files.
	Group *int64 `json:"Group,omitnil,omitempty" name:"Group"`

	// Backup type. Valid values: `1` (data backup), `2` (log backup). Default value: `1`.
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// Filter by backup file format. Valid values: `pkg` (archive file), `single` (Unarchived files).
	BackupFormat *string `json:"BackupFormat,omitnil,omitempty" name:"BackupFormat"`
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
	delete(f, "Group")
	delete(f, "Type")
	delete(f, "BackupFormat")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupsResponseParams struct {
	// Total number of backups
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Backup list details
	Backups []*Backup `json:"Backups,omitnil,omitempty" name:"Backups"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBackupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBackupsResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeBusinessIntelligenceFileRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// File name
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// Migration task status set. Valid values: `1` (Initialize to be deployed), `2` (Deploying), `3` (Deployment successful), `4` (Deployment failed)
	StatusSet []*int64 `json:"StatusSet,omitnil,omitempty" name:"StatusSet"`

	// File type. Valid values: `FLAT` (flat files), `SSIS` (project file for business intelligence service).
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// The maximum number of results returned per page. Value range: 1-100.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Page number. Default value: `0`.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Sorting field. Valid values: `file_name`, `create_time`, `start_time`.
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Sorting order: Valid values: `desc`, `asc`.
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`
}

type DescribeBusinessIntelligenceFileRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// File name
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// Migration task status set. Valid values: `1` (Initialize to be deployed), `2` (Deploying), `3` (Deployment successful), `4` (Deployment failed)
	StatusSet []*int64 `json:"StatusSet,omitnil,omitempty" name:"StatusSet"`

	// File type. Valid values: `FLAT` (flat files), `SSIS` (project file for business intelligence service).
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// The maximum number of results returned per page. Value range: 1-100.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Page number. Default value: `0`.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Sorting field. Valid values: `file_name`, `create_time`, `start_time`.
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Sorting order: Valid values: `desc`, `asc`.
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`
}

func (r *DescribeBusinessIntelligenceFileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBusinessIntelligenceFileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "FileName")
	delete(f, "StatusSet")
	delete(f, "FileType")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBusinessIntelligenceFileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBusinessIntelligenceFileResponseParams struct {
	// Total number of file deployment tasks
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// File deployment task set
	BackupMigrationSet []*BusinessIntelligenceFile `json:"BackupMigrationSet,omitnil,omitempty" name:"BackupMigrationSet"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBusinessIntelligenceFileResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBusinessIntelligenceFileResponseParams `json:"Response"`
}

func (r *DescribeBusinessIntelligenceFileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBusinessIntelligenceFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBCharsetsRequestParams struct {
	// Instance ID in the format of mssql-j8kv137v
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeDBCharsetsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of mssql-j8kv137v
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
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

// Predefined struct for user
type DescribeDBCharsetsResponseParams struct {
	// Database character set list
	DatabaseCharsets []*string `json:"DatabaseCharsets,omitnil,omitempty" name:"DatabaseCharsets"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBCharsetsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBCharsetsResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeDBInstanceInterRequestParams struct {
	// The maximum number of results returned per page. Value range: 1-100.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Filter by instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Filter by status. Valid values: `1` (Enabling interworking IP), `2` (Enabled interworking IP), `3` (Adding to interworking group), `4` (Added to interworking group), `5` (Reclaiming interworking IP), `6` (Reclaimed interworking IP), `7` (Removing from interworking group), `8` (Removed from interworking group).
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// The list of instance version numbers
	VersionSet []*string `json:"VersionSet,omitnil,omitempty" name:"VersionSet"`

	// Instance AZ ID in the format of ap-guangzhou-2
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Page number. Default value: `0`.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeDBInstanceInterRequest struct {
	*tchttp.BaseRequest
	
	// The maximum number of results returned per page. Value range: 1-100.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Filter by instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Filter by status. Valid values: `1` (Enabling interworking IP), `2` (Enabled interworking IP), `3` (Adding to interworking group), `4` (Added to interworking group), `5` (Reclaiming interworking IP), `6` (Reclaimed interworking IP), `7` (Removing from interworking group), `8` (Removed from interworking group).
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// The list of instance version numbers
	VersionSet []*string `json:"VersionSet,omitnil,omitempty" name:"VersionSet"`

	// Instance AZ ID in the format of ap-guangzhou-2
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Page number. Default value: `0`.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeDBInstanceInterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstanceInterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "InstanceId")
	delete(f, "Status")
	delete(f, "VersionSet")
	delete(f, "Zone")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBInstanceInterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstanceInterResponseParams struct {
	// Number of records returned
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Details of instance in the interworking group
	InterInstanceSet []*InterInstance `json:"InterInstanceSet,omitnil,omitempty" name:"InterInstanceSet"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBInstanceInterResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBInstanceInterResponseParams `json:"Response"`
}

func (r *DescribeDBInstanceInterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstanceInterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstancesAttributeRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeDBInstancesAttributeRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeDBInstancesAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstancesAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBInstancesAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstancesAttributeResponseParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Archive backup status. Valid values: `enable` (enabled), `disable` (disabled)
	RegularBackupEnable *string `json:"RegularBackupEnable,omitnil,omitempty" name:"RegularBackupEnable"`

	// Archive backup retention period: [90-3650] days
	RegularBackupSaveDays *uint64 `json:"RegularBackupSaveDays,omitnil,omitempty" name:"RegularBackupSaveDays"`

	// Archive backup policy. Valid values: `years` (yearly); `quarters (quarterly); `months` (monthly).
	RegularBackupStrategy *string `json:"RegularBackupStrategy,omitnil,omitempty" name:"RegularBackupStrategy"`

	// The number of retained archive backups
	RegularBackupCounts *uint64 `json:"RegularBackupCounts,omitnil,omitempty" name:"RegularBackupCounts"`

	// Archive backup start date in YYYY-MM-DD format, which is the current time by default.
	RegularBackupStartTime *string `json:"RegularBackupStartTime,omitnil,omitempty" name:"RegularBackupStartTime"`

	// Block process threshold in milliseconds
	BlockedThreshold *int64 `json:"BlockedThreshold,omitnil,omitempty" name:"BlockedThreshold"`

	// Retention period for the files of slow SQL, blocking, deadlock, and extended events.
	EventSaveDays *int64 `json:"EventSaveDays,omitnil,omitempty" name:"EventSaveDays"`

	// TDE configuration
	TDEConfig *TDEConfigAttribute `json:"TDEConfig,omitnil,omitempty" name:"TDEConfig"`


	SSLConfig *SSLConfig `json:"SSLConfig,omitnil,omitempty" name:"SSLConfig"`


	DrReadableInfo *DrReadableInfo `json:"DrReadableInfo,omitnil,omitempty" name:"DrReadableInfo"`


	OldVipList []*OldVip `json:"OldVipList,omitnil,omitempty" name:"OldVipList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBInstancesAttributeResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBInstancesAttributeResponseParams `json:"Response"`
}

func (r *DescribeDBInstancesAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstancesAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstancesRequestParams struct {
	// Project ID
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

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
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Page number. Default value: 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of results per page. Value range: 1-100. Default value: 100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// One or more instance IDs in the format of mssql-si2823jyl
	InstanceIdSet []*string `json:"InstanceIdSet,omitnil,omitempty" name:"InstanceIdSet"`

	// Retrieves billing type. 0: pay-as-you-go
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// Unique string-type ID of instance VPC in the format of `vpc-xxx`. If an empty string ("") is passed in, filtering will be made by basic network.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Unique string-type ID of instance subnet in the format of `subnet-xxx`. If an empty string ("") is passed in, filtering will be made by basic network.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// The list of instance private IPs, such as 172.1.0.12
	VipSet []*string `json:"VipSet,omitnil,omitempty" name:"VipSet"`

	// The list of instance names used for fuzzy match
	InstanceNameSet []*string `json:"InstanceNameSet,omitnil,omitempty" name:"InstanceNameSet"`

	// The list of instance version numbers, such as 2008R2, 2012SP3
	VersionSet []*string `json:"VersionSet,omitnil,omitempty" name:"VersionSet"`

	// Instance availability zone, such as ap-guangzhou-2
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// The list of instance tags
	TagKeys []*string `json:"TagKeys,omitnil,omitempty" name:"TagKeys"`

	// Keyword used for fuzzy match, including instance ID, instance name, and instance private IP
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`

	// Unique Uid of an instance
	UidSet []*string `json:"UidSet,omitnil,omitempty" name:"UidSet"`

	// Instance type. Valid values: `HA` (high-availability), `RO` (read-only), `SI` (basic edition), `BI` (business intelligence service).
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`


	PaginationType *string `json:"PaginationType,omitnil,omitempty" name:"PaginationType"`
}

type DescribeDBInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Project ID
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

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
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Page number. Default value: 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of results per page. Value range: 1-100. Default value: 100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// One or more instance IDs in the format of mssql-si2823jyl
	InstanceIdSet []*string `json:"InstanceIdSet,omitnil,omitempty" name:"InstanceIdSet"`

	// Retrieves billing type. 0: pay-as-you-go
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// Unique string-type ID of instance VPC in the format of `vpc-xxx`. If an empty string ("") is passed in, filtering will be made by basic network.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Unique string-type ID of instance subnet in the format of `subnet-xxx`. If an empty string ("") is passed in, filtering will be made by basic network.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// The list of instance private IPs, such as 172.1.0.12
	VipSet []*string `json:"VipSet,omitnil,omitempty" name:"VipSet"`

	// The list of instance names used for fuzzy match
	InstanceNameSet []*string `json:"InstanceNameSet,omitnil,omitempty" name:"InstanceNameSet"`

	// The list of instance version numbers, such as 2008R2, 2012SP3
	VersionSet []*string `json:"VersionSet,omitnil,omitempty" name:"VersionSet"`

	// Instance availability zone, such as ap-guangzhou-2
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// The list of instance tags
	TagKeys []*string `json:"TagKeys,omitnil,omitempty" name:"TagKeys"`

	// Keyword used for fuzzy match, including instance ID, instance name, and instance private IP
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`

	// Unique Uid of an instance
	UidSet []*string `json:"UidSet,omitnil,omitempty" name:"UidSet"`

	// Instance type. Valid values: `HA` (high-availability), `RO` (read-only), `SI` (basic edition), `BI` (business intelligence service).
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	PaginationType *string `json:"PaginationType,omitnil,omitempty" name:"PaginationType"`
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
	delete(f, "UidSet")
	delete(f, "InstanceType")
	delete(f, "PaginationType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstancesResponseParams struct {
	// Total number of eligible instances. If the results are returned in multiple pages, this value will be the number of all eligible instances but not the number of instances returned according to the current values of `Limit` and `Offset`
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Instance list
	DBInstances []*DBInstance `json:"DBInstances,omitnil,omitempty" name:"DBInstances"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBInstancesResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeDBsNormalRequestParams struct {
	// Instance ID in the format of mssql-7vfv3rk3
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeDBsNormalRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of mssql-7vfv3rk3
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeDBsNormalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBsNormalRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBsNormalRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBsNormalResponseParams struct {
	// Total number of databases of the instance
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Detailed database configurations, such as whether CDC or CT is enabled for the database
	DBList []*DbNormalDetail `json:"DBList,omitnil,omitempty" name:"DBList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBsNormalResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBsNormalResponseParams `json:"Response"`
}

func (r *DescribeDBsNormalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBsNormalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBsRequestParams struct {
	// Instance ID
	InstanceIdSet []*string `json:"InstanceIdSet,omitnil,omitempty" name:"InstanceIdSet"`

	// Number of results per page. Value range: 1-100. Default value: 20
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Page number. Default value: 0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Database name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Sorting rule. Valid values: `desc` (descending order), `asc` (ascending order). Default value: `desc`.
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// TDE status. Valid values: `enable` (enabled), `disable` (disabled).
	Encryption *string `json:"Encryption,omitnil,omitempty" name:"Encryption"`
}

type DescribeDBsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceIdSet []*string `json:"InstanceIdSet,omitnil,omitempty" name:"InstanceIdSet"`

	// Number of results per page. Value range: 1-100. Default value: 20
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Page number. Default value: 0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Database name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Sorting rule. Valid values: `desc` (descending order), `asc` (ascending order). Default value: `desc`.
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// TDE status. Valid values: `enable` (enabled), `disable` (disabled).
	Encryption *string `json:"Encryption,omitnil,omitempty" name:"Encryption"`
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
	delete(f, "Name")
	delete(f, "OrderByType")
	delete(f, "Encryption")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBsResponseParams struct {
	// Number of databases
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List of instance databases
	DBInstances []*InstanceDBDetail `json:"DBInstances,omitnil,omitempty" name:"DBInstances"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBsResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeFlowStatusRequestParams struct {
	// Flow ID
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`
}

type DescribeFlowStatusRequest struct {
	*tchttp.BaseRequest
	
	// Flow ID
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`
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

// Predefined struct for user
type DescribeFlowStatusResponseParams struct {
	// Flow status. 0: succeeded, 1: failed, 2: running
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeFlowStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFlowStatusResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeIncrementalMigrationRequestParams struct {
	// Backup import task ID, which is returned through the API CreateBackupMigration
	BackupMigrationId *string `json:"BackupMigrationId,omitnil,omitempty" name:"BackupMigrationId"`

	// ID of imported target instance
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Backup file name
	BackupFileName *string `json:"BackupFileName,omitnil,omitempty" name:"BackupFileName"`

	// Status set of import tasks
	StatusSet []*int64 `json:"StatusSet,omitnil,omitempty" name:"StatusSet"`

	// The maximum number of results returned per page. Default value: `100`.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Page number. Default value: `0`.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Sort by field. Valid values: `name`, `createTime`, `startTime`, `endTime`. By default, the results returned are sorted by `createTime` in the ascending order.
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Sorting order which is valid only when `OrderBy` is specified. Valid values: `asc` (ascending), `desc` (descending). Default value: `asc`.
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// Incremental backup import task ID, which is returned through the `CreateIncrementalMigration` API.
	IncrementalMigrationId *string `json:"IncrementalMigrationId,omitnil,omitempty" name:"IncrementalMigrationId"`
}

type DescribeIncrementalMigrationRequest struct {
	*tchttp.BaseRequest
	
	// Backup import task ID, which is returned through the API CreateBackupMigration
	BackupMigrationId *string `json:"BackupMigrationId,omitnil,omitempty" name:"BackupMigrationId"`

	// ID of imported target instance
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Backup file name
	BackupFileName *string `json:"BackupFileName,omitnil,omitempty" name:"BackupFileName"`

	// Status set of import tasks
	StatusSet []*int64 `json:"StatusSet,omitnil,omitempty" name:"StatusSet"`

	// The maximum number of results returned per page. Default value: `100`.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Page number. Default value: `0`.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Sort by field. Valid values: `name`, `createTime`, `startTime`, `endTime`. By default, the results returned are sorted by `createTime` in the ascending order.
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Sorting order which is valid only when `OrderBy` is specified. Valid values: `asc` (ascending), `desc` (descending). Default value: `asc`.
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// Incremental backup import task ID, which is returned through the `CreateIncrementalMigration` API.
	IncrementalMigrationId *string `json:"IncrementalMigrationId,omitnil,omitempty" name:"IncrementalMigrationId"`
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

// Predefined struct for user
type DescribeIncrementalMigrationResponseParams struct {
	// Total number of import tasks
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Incremental import task set
	IncrementalMigrationSet []*Migration `json:"IncrementalMigrationSet,omitnil,omitempty" name:"IncrementalMigrationSet"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeIncrementalMigrationResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIncrementalMigrationResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeInstanceByOrdersRequestParams struct {
	// Order ID set
	DealNames []*string `json:"DealNames,omitnil,omitempty" name:"DealNames"`
}

type DescribeInstanceByOrdersRequest struct {
	*tchttp.BaseRequest
	
	// Order ID set
	DealNames []*string `json:"DealNames,omitnil,omitempty" name:"DealNames"`
}

func (r *DescribeInstanceByOrdersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceByOrdersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DealNames")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceByOrdersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceByOrdersResponseParams struct {

	DealInstance []*DealInstance `json:"DealInstance,omitnil,omitempty" name:"DealInstance"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceByOrdersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceByOrdersResponseParams `json:"Response"`
}

func (r *DescribeInstanceByOrdersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceByOrdersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceParamRecordsRequestParams struct {
	// Instance ID in the format of mssql-dj5i29c5n. It is the same as the instance ID displayed in the TencentDB console and the response parameter `InstanceId` of the `DescribeDBInstances` API.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Page number. Default value: `0`.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The maximum number of results returned per page. Maximum value: `100`. Default value: `20`.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeInstanceParamRecordsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of mssql-dj5i29c5n. It is the same as the instance ID displayed in the TencentDB console and the response parameter `InstanceId` of the `DescribeDBInstances` API.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Page number. Default value: `0`.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The maximum number of results returned per page. Maximum value: `100`. Default value: `20`.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
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

// Predefined struct for user
type DescribeInstanceParamRecordsResponseParams struct {
	// Number of eligible records
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Parameter modification records
	Items []*ParamRecord `json:"Items,omitnil,omitempty" name:"Items"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceParamRecordsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceParamRecordsResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeInstanceParamsRequestParams struct {
	// Instance ID in the format of mssql-dj5i29c5n. It is the same as the instance ID displayed in the TencentDB console and the response parameter `InstanceId` of the `DescribeDBInstances` API.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeInstanceParamsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of mssql-dj5i29c5n. It is the same as the instance ID displayed in the TencentDB console and the response parameter `InstanceId` of the `DescribeDBInstances` API.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
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

// Predefined struct for user
type DescribeInstanceParamsResponseParams struct {
	// Total number of instance parameters
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Parameter details
	Items []*ParameterDetail `json:"Items,omitnil,omitempty" name:"Items"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceParamsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceParamsResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeMigrationDetailRequestParams struct {
	// Migration task ID
	MigrateId *uint64 `json:"MigrateId,omitnil,omitempty" name:"MigrateId"`
}

type DescribeMigrationDetailRequest struct {
	*tchttp.BaseRequest
	
	// Migration task ID
	MigrateId *uint64 `json:"MigrateId,omitnil,omitempty" name:"MigrateId"`
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

// Predefined struct for user
type DescribeMigrationDetailResponseParams struct {
	// Migration task ID
	MigrateId *uint64 `json:"MigrateId,omitnil,omitempty" name:"MigrateId"`

	// Migration task name
	MigrateName *string `json:"MigrateName,omitnil,omitempty" name:"MigrateName"`

	// User ID of migration task
	AppId *uint64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// Migration task region
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Migration source type. 1: TencentDB for SQL Server, 2: CVM-based self-created SQL Server database; 3: SQL Server backup restoration, 4: SQL Server backup restoration (in COS mode)
	SourceType *int64 `json:"SourceType,omitnil,omitempty" name:"SourceType"`

	// Migration task creation time
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Migration task start time
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Migration task end time
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Migration task status (1: initializing, 4: migrating, 5: migration failed, 6: migration succeeded)
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Migration task progress
	Progress *int64 `json:"Progress,omitnil,omitempty" name:"Progress"`

	// Migration type (1: structure migration, 2: data migration, 3: incremental sync)
	MigrateType *int64 `json:"MigrateType,omitnil,omitempty" name:"MigrateType"`

	// Migration source
	Source *MigrateSource `json:"Source,omitnil,omitempty" name:"Source"`

	// Migration target
	Target *MigrateTarget `json:"Target,omitnil,omitempty" name:"Target"`

	// Database objects to be migrated. This parameter is not used for offline migration (SourceType=4 or SourceType=5)
	MigrateDBSet []*MigrateDB `json:"MigrateDBSet,omitnil,omitempty" name:"MigrateDBSet"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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
type DescribeMigrationsRequestParams struct {
	// Status set. As long as a migration task is in a status therein, it will be listed
	StatusSet []*int64 `json:"StatusSet,omitnil,omitempty" name:"StatusSet"`

	// Migration task name (fuzzy match)
	MigrateName *string `json:"MigrateName,omitnil,omitempty" name:"MigrateName"`

	// Number of results per page. Value range: 1-100. Default value: 100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Page number. Default value: 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The query results are sorted by keyword. Valid values: name, createTime, startTime, endTime, status
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Sorting order. Valid values: desc, asc
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`
}

type DescribeMigrationsRequest struct {
	*tchttp.BaseRequest
	
	// Status set. As long as a migration task is in a status therein, it will be listed
	StatusSet []*int64 `json:"StatusSet,omitnil,omitempty" name:"StatusSet"`

	// Migration task name (fuzzy match)
	MigrateName *string `json:"MigrateName,omitnil,omitempty" name:"MigrateName"`

	// Number of results per page. Value range: 1-100. Default value: 100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Page number. Default value: 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The query results are sorted by keyword. Valid values: name, createTime, startTime, endTime, status
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Sorting order. Valid values: desc, asc
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`
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

// Predefined struct for user
type DescribeMigrationsResponseParams struct {
	// Total number of query results
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List of query results
	MigrateTaskSet []*MigrateTask `json:"MigrateTaskSet,omitnil,omitempty" name:"MigrateTaskSet"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMigrationsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMigrationsResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeOrdersRequestParams struct {
	// Order array. The order name will be returned upon shipping, which can be used to call the `DescribeOrders` API to query shipment status
	DealNames []*string `json:"DealNames,omitnil,omitempty" name:"DealNames"`
}

type DescribeOrdersRequest struct {
	*tchttp.BaseRequest
	
	// Order array. The order name will be returned upon shipping, which can be used to call the `DescribeOrders` API to query shipment status
	DealNames []*string `json:"DealNames,omitnil,omitempty" name:"DealNames"`
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

// Predefined struct for user
type DescribeOrdersResponseParams struct {
	// Order information array
	Deals []*DealInfo `json:"Deals,omitnil,omitempty" name:"Deals"`

	// Number of orders returned
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeOrdersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOrdersResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeProductConfigRequestParams struct {
	// AZ ID in the format of ap-guangzhou-1
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Type of purchased instance. Valid values: HA - local disk high availability (including dual-machine high availability, AlwaysOn Cluster), RO - local disk read-only replica, SI - cloud disk edition single node, BI - business intelligence service, cvmHA - cloud disk edition high availability, cvmRO - cloud disk edition read-only replica
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`
}

type DescribeProductConfigRequest struct {
	*tchttp.BaseRequest
	
	// AZ ID in the format of ap-guangzhou-1
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Type of purchased instance. Valid values: HA - local disk high availability (including dual-machine high availability, AlwaysOn Cluster), RO - local disk read-only replica, SI - cloud disk edition single node, BI - business intelligence service, cvmHA - cloud disk edition high availability, cvmRO - cloud disk edition read-only replica
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`
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

// Predefined struct for user
type DescribeProductConfigResponseParams struct {
	// Specification information array
	SpecInfoList []*SpecInfo `json:"SpecInfoList,omitnil,omitempty" name:"SpecInfoList"`

	// Number of date entries returned
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeProductConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProductConfigResponseParams `json:"Response"`
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
	// Total number of regions returned
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Region information array
	RegionSet []*RegionInfo `json:"RegionSet,omitnil,omitempty" name:"RegionSet"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeRestoreTimeRangeRequestParams struct {

	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`


	TargetInstanceId *string `json:"TargetInstanceId,omitnil,omitempty" name:"TargetInstanceId"`
}

type DescribeRestoreTimeRangeRequest struct {
	*tchttp.BaseRequest
	
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	TargetInstanceId *string `json:"TargetInstanceId,omitnil,omitempty" name:"TargetInstanceId"`
}

func (r *DescribeRestoreTimeRangeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRestoreTimeRangeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "TargetInstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRestoreTimeRangeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRestoreTimeRangeResponseParams struct {

	MinTime *string `json:"MinTime,omitnil,omitempty" name:"MinTime"`


	MaxTime *string `json:"MaxTime,omitnil,omitempty" name:"MaxTime"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRestoreTimeRangeResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRestoreTimeRangeResponseParams `json:"Response"`
}

func (r *DescribeRestoreTimeRangeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRestoreTimeRangeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRollbackTimeRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// List of databases to be queried
	DBs []*string `json:"DBs,omitnil,omitempty" name:"DBs"`
}

type DescribeRollbackTimeRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// List of databases to be queried
	DBs []*string `json:"DBs,omitnil,omitempty" name:"DBs"`
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

// Predefined struct for user
type DescribeRollbackTimeResponseParams struct {
	// Information of time range available for database rollback
	Details []*DbRollbackTimeInfo `json:"Details,omitnil,omitempty" name:"Details"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRollbackTimeResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRollbackTimeResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeSlowlogsRequestParams struct {
	// Instance ID in the format of mssql-k8voqdlz
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Start time in the format of `yyyy-MM-dd HH:mm:ss`
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time in the format of `yyyy-MM-dd HH:mm:ss`
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Number of results per page. Value range: 1-100. Default value: 20
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Page number. Default value: 0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeSlowlogsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of mssql-k8voqdlz
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Start time in the format of `yyyy-MM-dd HH:mm:ss`
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time in the format of `yyyy-MM-dd HH:mm:ss`
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Number of results per page. Value range: 1-100. Default value: 20
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Page number. Default value: 0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
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

// Predefined struct for user
type DescribeSlowlogsResponseParams struct {
	// Total number of queries
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Information list of slow query logs
	//
	// Deprecated: Slowlogs is deprecated.
	Slowlogs []*SlowlogInfo `json:"Slowlogs,omitnil,omitempty" name:"Slowlogs"`


	SlowLogs []*SlowLog `json:"SlowLogs,omitnil,omitempty" name:"SlowLogs"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSlowlogsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSlowlogsResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeUploadBackupInfoRequestParams struct {
	// ID of imported target instance
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Backup import task ID, which is returned through the API CreateBackupMigration
	BackupMigrationId *string `json:"BackupMigrationId,omitnil,omitempty" name:"BackupMigrationId"`
}

type DescribeUploadBackupInfoRequest struct {
	*tchttp.BaseRequest
	
	// ID of imported target instance
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Backup import task ID, which is returned through the API CreateBackupMigration
	BackupMigrationId *string `json:"BackupMigrationId,omitnil,omitempty" name:"BackupMigrationId"`
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

// Predefined struct for user
type DescribeUploadBackupInfoResponseParams struct {
	// Bucket name
	BucketName *string `json:"BucketName,omitnil,omitempty" name:"BucketName"`

	// Bucket location information
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Storage path
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// Temporary key ID
	TmpSecretId *string `json:"TmpSecretId,omitnil,omitempty" name:"TmpSecretId"`

	// Temporary key (Key)
	TmpSecretKey *string `json:"TmpSecretKey,omitnil,omitempty" name:"TmpSecretKey"`

	// Temporary key (Token)
	//
	// Deprecated: XCosSecurityToken is deprecated.
	XCosSecurityToken *string `json:"XCosSecurityToken,omitnil,omitempty" name:"XCosSecurityToken"`

	// Temporary key start time
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Temporary key expiration time
	ExpiredTime *string `json:"ExpiredTime,omitnil,omitempty" name:"ExpiredTime"`


	CosSecurityToken *string `json:"CosSecurityToken,omitnil,omitempty" name:"CosSecurityToken"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeUploadBackupInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUploadBackupInfoResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeXEventsRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Event type. Valid values: `slow` (Slow SQL event), `blocked` (blocking event),  deadlock` (deadlock event).
	EventType *string `json:"EventType,omitnil,omitempty" name:"EventType"`

	// Generation start time of an extended file in the format of yyyy-MM-dd HH:mm:ss.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Generation end time of an extended file in the format of yyyy-MM-dd HH:mm:ss.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Page number. Default value: `0`
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of entries to be returned per page. Value range: 1-100. Default value: `20`
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeXEventsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Event type. Valid values: `slow` (Slow SQL event), `blocked` (blocking event),  deadlock` (deadlock event).
	EventType *string `json:"EventType,omitnil,omitempty" name:"EventType"`

	// Generation start time of an extended file in the format of yyyy-MM-dd HH:mm:ss.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Generation end time of an extended file in the format of yyyy-MM-dd HH:mm:ss.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Page number. Default value: `0`
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of entries to be returned per page. Value range: 1-100. Default value: `20`
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeXEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeXEventsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "EventType")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeXEventsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeXEventsResponseParams struct {
	// List of extended events
	Events []*Events `json:"Events,omitnil,omitempty" name:"Events"`

	// Total number of extended events
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeXEventsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeXEventsResponseParams `json:"Response"`
}

func (r *DescribeXEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeXEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeZonesRequestParams struct {

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

// Predefined struct for user
type DescribeZonesResponseParams struct {
	// Number of AZs returned
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Array of AZs
	ZoneSet []*ZoneInfo `json:"ZoneSet,omitnil,omitempty" name:"ZoneSet"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeZonesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeZonesResponseParams `json:"Response"`
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

type DrReadableInfo struct {
	// Replica server status. Valid values: enable - running; disable - unavailable
	// Note: This field may return null, indicating that no valid values can be obtained.
	SlaveStatus *string `json:"SlaveStatus,omitnil,omitempty" name:"SlaveStatus"`

	// Replica server readable status. Valid values: enable - enabled; disable - disabled
	// Note: This field may return null, indicating that no valid values can be obtained.
	ReadableStatus *string `json:"ReadableStatus,omitnil,omitempty" name:"ReadableStatus"`

	// Replica server read-only VIP
	// Note: This field may return null, indicating that no valid values can be obtained.
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// Replica server read-only port
	// Note: This field may return null, indicating that no valid values can be obtained.
	VPort *int64 `json:"VPort,omitnil,omitempty" name:"VPort"`

	// Replica server VPC ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// Replica server VPC subnet ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	UniqSubnetId *string `json:"UniqSubnetId,omitnil,omitempty" name:"UniqSubnetId"`
}

type EventConfig struct {
	// Event type. Valid values: `slow` (set threshold for slow SQL ), `blocked` (set threshold for the blocking and deadlock).
	EventType *string `json:"EventType,omitnil,omitempty" name:"EventType"`

	// Threshold in milliseconds. Valid values: `0`(disable), `non-zero` (enable)
	Threshold *int64 `json:"Threshold,omitnil,omitempty" name:"Threshold"`
}

type Events struct {
	// ID
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// File name of an extended event
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// File size of an extended event
	Size *int64 `json:"Size,omitnil,omitempty" name:"Size"`

	// Event type. Valid values: `slow` (Slow SQL event), `blocked` (blocking event),  `deadlock` (deadlock event).
	EventType *string `json:"EventType,omitnil,omitempty" name:"EventType"`

	// Event record status. Valid values: `1` (succeeded), `2` (failed).
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Generation start time of an extended file
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Generation end time of an extended file
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Download address on the private network
	InternalAddr *string `json:"InternalAddr,omitnil,omitempty" name:"InternalAddr"`

	// Download address on the public network
	ExternalAddr *string `json:"ExternalAddr,omitnil,omitempty" name:"ExternalAddr"`
}

type FileAction struct {
	// Allowed operations. Valid values: `view` (view list), `remark` (modify remark), `deploy` (deploy files), `delete` (delete files).
	AllAction []*string `json:"AllAction,omitnil,omitempty" name:"AllAction"`

	// Operation allowed in the current status. If the subset of `AllAction` is empty, no operations will be allowed.
	AllowedAction []*string `json:"AllowedAction,omitnil,omitempty" name:"AllowedAction"`
}

// Predefined struct for user
type InquiryPriceCreateDBInstancesRequestParams struct {
	// AZ ID, which can be obtained through the `Zone` field in the returned value of the `DescribeZones` API
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Memory size in GB
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// Instance capacity in GB
	Storage *int64 `json:"Storage,omitnil,omitempty" name:"Storage"`

	// Billing type. Valid value: POSTPAID.
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// Length of purchase in months. Value range: 1-48. Default value: 1
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// Number of instances purchased at a time. Value range: 1-100. Default value: 1
	GoodsNum *int64 `json:"GoodsNum,omitnil,omitempty" name:"GoodsNum"`

	// SQL version. Valid values: `2008R2` (SQL Server 2008 R2 Enterprise), `2012SP3` (SQL Server 2012 Enterprise), `201202` (SQL Server 2012 Standard), `2014SP2` (SQL Server 2014 Enterprise), `201402` (SQL Server 2014 Standard)`, `2016SP1` (SQL Server 2016 Enterprise), `201602` (SQL Server 2016 Standard), `2017` (SQL Server 2017 Enterprise), `201702` (SQL Server 2017 Standard), `2019` (SQL Server 2019 Enterprise), `201902` (SQL Server 2019 Standard). Default value: `2008R2`. The purchasable version varies by region. It can be queried by the `DescribeProductConfig` API.
	DBVersion *string `json:"DBVersion,omitnil,omitempty" name:"DBVersion"`

	// The number of CPU cores of the instance you want to purchase.
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// The type of instance to be purchased. Valid values: `HA` (high-availability edition, including dual-server high-availability and AlwaysOn cluster edition), `RO` (read-only replica edition), `SI` (single-node edition), `cvmHA` (dual-server high-availability edition for CVM), `cvmRO` (read-only edition for CVM).
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// The host type of the instance to be purchased. Valid values: `PM` (physical machine), `CLOUD_PREMIUM` (virtual machine with premium cloud disk), `CLOUD_SSD` (virtual machine with SSD), 
	// `CLOUD_HSSD` (virtual machine with enhanced SSD), `CLOUD_TSSD` (virtual machine with ulTra SSD), `CLOUD_BSSD` (virtual machine with balanced SSD).
	MachineType *string `json:"MachineType,omitnil,omitempty" name:"MachineType"`
}

type InquiryPriceCreateDBInstancesRequest struct {
	*tchttp.BaseRequest
	
	// AZ ID, which can be obtained through the `Zone` field in the returned value of the `DescribeZones` API
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Memory size in GB
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// Instance capacity in GB
	Storage *int64 `json:"Storage,omitnil,omitempty" name:"Storage"`

	// Billing type. Valid value: POSTPAID.
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// Length of purchase in months. Value range: 1-48. Default value: 1
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// Number of instances purchased at a time. Value range: 1-100. Default value: 1
	GoodsNum *int64 `json:"GoodsNum,omitnil,omitempty" name:"GoodsNum"`

	// SQL version. Valid values: `2008R2` (SQL Server 2008 R2 Enterprise), `2012SP3` (SQL Server 2012 Enterprise), `201202` (SQL Server 2012 Standard), `2014SP2` (SQL Server 2014 Enterprise), `201402` (SQL Server 2014 Standard)`, `2016SP1` (SQL Server 2016 Enterprise), `201602` (SQL Server 2016 Standard), `2017` (SQL Server 2017 Enterprise), `201702` (SQL Server 2017 Standard), `2019` (SQL Server 2019 Enterprise), `201902` (SQL Server 2019 Standard). Default value: `2008R2`. The purchasable version varies by region. It can be queried by the `DescribeProductConfig` API.
	DBVersion *string `json:"DBVersion,omitnil,omitempty" name:"DBVersion"`

	// The number of CPU cores of the instance you want to purchase.
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// The type of instance to be purchased. Valid values: `HA` (high-availability edition, including dual-server high-availability and AlwaysOn cluster edition), `RO` (read-only replica edition), `SI` (single-node edition), `cvmHA` (dual-server high-availability edition for CVM), `cvmRO` (read-only edition for CVM).
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// The host type of the instance to be purchased. Valid values: `PM` (physical machine), `CLOUD_PREMIUM` (virtual machine with premium cloud disk), `CLOUD_SSD` (virtual machine with SSD), 
	// `CLOUD_HSSD` (virtual machine with enhanced SSD), `CLOUD_TSSD` (virtual machine with ulTra SSD), `CLOUD_BSSD` (virtual machine with balanced SSD).
	MachineType *string `json:"MachineType,omitnil,omitempty" name:"MachineType"`
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

// Predefined struct for user
type InquiryPriceCreateDBInstancesResponseParams struct {
	// Price before discount. This value divided by 100 indicates the price; for example, 10010 means 100.10 USD
	OriginalPrice *int64 `json:"OriginalPrice,omitnil,omitempty" name:"OriginalPrice"`

	// The actual price to be paid. This value divided by 100 indicates the price; for example, 10010 means 100.10 USD
	Price *int64 `json:"Price,omitnil,omitempty" name:"Price"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type InquiryPriceCreateDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *InquiryPriceCreateDBInstancesResponseParams `json:"Response"`
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

// Predefined struct for user
type InquiryPriceUpgradeDBInstanceRequestParams struct {
	// Instance ID in the format of mssql-njj2mtpl
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Memory size after instance upgrade in GB, which cannot be smaller than the current instance memory size
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// Storage capacity after instance upgrade in GB, which cannot be smaller than the current instance storage capacity
	Storage *int64 `json:"Storage,omitnil,omitempty" name:"Storage"`

	// The number of CUP cores after the instance is upgraded, which cannot be smaller than that of the current cores.
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`
}

type InquiryPriceUpgradeDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of mssql-njj2mtpl
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Memory size after instance upgrade in GB, which cannot be smaller than the current instance memory size
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// Storage capacity after instance upgrade in GB, which cannot be smaller than the current instance storage capacity
	Storage *int64 `json:"Storage,omitnil,omitempty" name:"Storage"`

	// The number of CUP cores after the instance is upgraded, which cannot be smaller than that of the current cores.
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`
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

// Predefined struct for user
type InquiryPriceUpgradeDBInstanceResponseParams struct {
	// Price before discount. This value divided by 100 indicates the price; for example, 10094 means 100.94 USD
	OriginalPrice *int64 `json:"OriginalPrice,omitnil,omitempty" name:"OriginalPrice"`

	// The actual price to be paid. This value divided by 100 indicates the price; for example, 10094 means 100.94 USD
	Price *int64 `json:"Price,omitnil,omitempty" name:"Price"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type InquiryPriceUpgradeDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *InquiryPriceUpgradeDBInstanceResponseParams `json:"Response"`
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
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Database information list
	DBDetails []*DBDetail `json:"DBDetails,omitnil,omitempty" name:"DBDetails"`
}

type InterInstance struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Instance interworking IP, which can be accessed after the instance is added to the interworking group.
	InterVip *string `json:"InterVip,omitnil,omitempty" name:"InterVip"`

	// Instance interworking port, which can be accessed after the instance is added to the interworking group.
	InterPort *int64 `json:"InterPort,omitnil,omitempty" name:"InterPort"`

	// Instance interworking status. Valid values: `1` (Enabling interworking IP), `2` (Enabled interworking IP), `3` (Adding to interworking group), `4` (Added to interworking group), `5` (Reclaiming interworking IP), `6`(Reclaimed interworking IP), `7` (Removing from interworking group), `8` (Removed from interworking group).
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Instance region, such as ap-guangzhou.
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Instance AZ name, such as ap-guangzhou-1.
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Instance version code
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// Instance version
	VersionName *string `json:"VersionName,omitnil,omitempty" name:"VersionName"`

	// Instance name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Instance access IP
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// Instance access port
	Vport *int64 `json:"Vport,omitnil,omitempty" name:"Vport"`
}

type InterInstanceFlow struct {
	// Instance ID, such as mssql-sdf32n1d.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Instance task ID for enabling or disabling the interworking group. When `FlowId` is less than 0, the interworking group will be enabled or disabled successfully; otherwise, the operation failed.
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`
}

type MigrateDB struct {
	// Name of migrated database
	DBName *string `json:"DBName,omitnil,omitempty" name:"DBName"`
}

type MigrateDetail struct {
	// Name of current step
	StepName *string `json:"StepName,omitnil,omitempty" name:"StepName"`

	// Progress of current step in %
	Progress *int64 `json:"Progress,omitnil,omitempty" name:"Progress"`
}

type MigrateSource struct {
	// Source instance ID in the format of `mssql-si2823jyl`, which is used when `MigrateType` is 1 (TencentDB for SQL Server)
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// ID of source CVM instance, which is used when `MigrateType` is 2 (CVM-based self-created SQL Server database)
	CvmId *string `json:"CvmId,omitnil,omitempty" name:"CvmId"`

	// VPC ID of source CVM instance in the format of vpc-6ys9ont9, which is used when `MigrateType` is 2 (CVM-based self-created SQL Server database)
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// VPC subnet ID of source CVM instance in the format of subnet-h9extioi, which is used when `MigrateType` is 2 (CVM-based self-created SQL Server database)
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Username, which is used when `MigrateType` is 1 or 2
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// Password, which is used when `MigrateType` is 1 or 2
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// Private IP of source CVM database, which is used when `MigrateType` is 2 (CVM-based self-created SQL Server database)
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// Port number of source CVM database, which is used when `MigrateType` is 2 (CVM-based self-created SQL Server database)
	Port *uint64 `json:"Port,omitnil,omitempty" name:"Port"`

	// Source backup address for offline migration, which is used when `MigrateType` is 4 or 5
	Url []*string `json:"Url,omitnil,omitempty" name:"Url"`

	// Source backup password for offline migration, which is used when `MigrateType` is 4 or 5
	UrlPassword *string `json:"UrlPassword,omitnil,omitempty" name:"UrlPassword"`
}

type MigrateTarget struct {
	// ID of target instance in the format of mssql-si2823jyl
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Username of migration target instance
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// Password of migration target instance
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`
}

type MigrateTask struct {
	// Migration task ID
	MigrateId *uint64 `json:"MigrateId,omitnil,omitempty" name:"MigrateId"`

	// Migration task name
	MigrateName *string `json:"MigrateName,omitnil,omitempty" name:"MigrateName"`

	// User ID of migration task
	AppId *uint64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// Migration task region
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Migration source type. 1: TencentDB for SQL Server, 2: CVM-based self-created SQL Server database; 3: SQL Server backup restoration, 4: SQL Server backup restoration (in COS mode)
	SourceType *int64 `json:"SourceType,omitnil,omitempty" name:"SourceType"`

	// Migration task creation time
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Migration task start time
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Migration task end time
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Migration task status (1: initializing, 4: migrating, 5: migration failed, 6: migration succeeded, 7: suspended, 8: deleted, 9: suspending, 10: completing, 11: suspension failed, 12: completion failed)
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Information
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// Whether migration task has been checked (0: not checked, 1: check succeeded, 2: check failed, 3: checking)
	CheckFlag *uint64 `json:"CheckFlag,omitnil,omitempty" name:"CheckFlag"`

	// Migration task progress in %
	Progress *int64 `json:"Progress,omitnil,omitempty" name:"Progress"`

	// Migration task progress details
	MigrateDetail *MigrateDetail `json:"MigrateDetail,omitnil,omitempty" name:"MigrateDetail"`
}

type Migration struct {
	// Backup import task ID or incremental import task ID
	MigrationId *string `json:"MigrationId,omitnil,omitempty" name:"MigrationId"`

	// Backup import task name. For an incremental import task, this field will be left empty.
	// Note: this field may return ‘null’, indicating that no valid values can be obtained.
	MigrationName *string `json:"MigrationName,omitnil,omitempty" name:"MigrationName"`

	// Application ID
	AppId *uint64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// Region
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// ID of migrated target instance
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Migration task restoration type
	RecoveryType *string `json:"RecoveryType,omitnil,omitempty" name:"RecoveryType"`

	// Backup user upload type. COS_URL: the backup is stored in user’s Cloud Object Storage, with URL provided. COS_UPLOAD: the backup is stored in the application’s Cloud Object Storage and needs to be uploaded by the user.
	UploadType *string `json:"UploadType,omitnil,omitempty" name:"UploadType"`

	// Backup file list, which is determined by UploadType. If the upload type is COS_URL, URL will be saved. If the upload type is COS_UPLOAD, the backup name will be saved.
	BackupFiles []*string `json:"BackupFiles,omitnil,omitempty" name:"BackupFiles"`

	// Migration task status. Valid values: `2` (Creation completed), `7` (Importing full backups), `8` (Waiting for incremental backups), `9` (Import success), `10` (Import failure), `12` (Importing incremental backups).
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Migration task creation time
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Migration task start time
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Migration task end time
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// More information
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// Migration detail
	Detail *MigrationDetail `json:"Detail,omitnil,omitempty" name:"Detail"`

	// Operation allowed in the current status
	Action *MigrationAction `json:"Action,omitnil,omitempty" name:"Action"`

	// Whether this is the final restoration. For a full import task, this field will be left empty.
	// Note: this field may return ‘null’, indicating that no valid values can be obtained.
	IsRecovery *string `json:"IsRecovery,omitnil,omitempty" name:"IsRecovery"`

	// Name set of renamed databases
	// Note: This field may return null, indicating that no valid values can be obtained.
	DBRename []*DBRenameRes `json:"DBRename,omitnil,omitempty" name:"DBRename"`
}

type MigrationAction struct {
	// All the allowed operations. Values include: view (viewing a task), modify (modifying a task), start (starting a task), incremental (creating an incremental task), delete (deleting a task), and upload (obtaining the upload permission).
	AllAction []*string `json:"AllAction,omitnil,omitempty" name:"AllAction"`

	// Operation allowed in the current status. If the subset of AllAction is left empty, no operations will be allowed.
	AllowedAction []*string `json:"AllowedAction,omitnil,omitempty" name:"AllowedAction"`
}

type MigrationDetail struct {
	// Total number of steps
	StepAll *int64 `json:"StepAll,omitnil,omitempty" name:"StepAll"`

	// Current step
	StepNow *int64 `json:"StepNow,omitnil,omitempty" name:"StepNow"`

	// Overall progress. For example, “30” means 30%.
	Progress *int64 `json:"Progress,omitnil,omitempty" name:"Progress"`

	// Step information. ‘null’ means the migration has not started
	// Note: this field may return ‘null’, indicating that no valid values can be obtained.
	StepInfo []*MigrationStep `json:"StepInfo,omitnil,omitempty" name:"StepInfo"`
}

type MigrationStep struct {
	// Step sequence
	StepNo *int64 `json:"StepNo,omitnil,omitempty" name:"StepNo"`

	// Step name
	StepName *string `json:"StepName,omitnil,omitempty" name:"StepName"`

	// Step ID in English
	StepId *string `json:"StepId,omitnil,omitempty" name:"StepId"`

	// Step status: 0 (default value), 1 (succeeded), 2 (failed), 3 (in progress), 4 (not started)
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

// Predefined struct for user
type ModifyAccountPrivilegeRequestParams struct {
	// Database instance ID in the format of mssql-njj2mtpl
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Account permission change information
	Accounts []*AccountPrivilegeModifyInfo `json:"Accounts,omitnil,omitempty" name:"Accounts"`
}

type ModifyAccountPrivilegeRequest struct {
	*tchttp.BaseRequest
	
	// Database instance ID in the format of mssql-njj2mtpl
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Account permission change information
	Accounts []*AccountPrivilegeModifyInfo `json:"Accounts,omitnil,omitempty" name:"Accounts"`
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

// Predefined struct for user
type ModifyAccountPrivilegeResponseParams struct {
	// Async task flow ID
	FlowId *uint64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAccountPrivilegeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAccountPrivilegeResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyAccountRemarkRequestParams struct {
	// Instance ID in the format of mssql-j8kv137v
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Information of account for which to modify remarks
	Accounts []*AccountRemark `json:"Accounts,omitnil,omitempty" name:"Accounts"`
}

type ModifyAccountRemarkRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of mssql-j8kv137v
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Information of account for which to modify remarks
	Accounts []*AccountRemark `json:"Accounts,omitnil,omitempty" name:"Accounts"`
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

// Predefined struct for user
type ModifyAccountRemarkResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAccountRemarkResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAccountRemarkResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyBackupMigrationRequestParams struct {
	// ID of imported target instance
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Backup import task ID, which is returned through the API CreateBackupMigration
	BackupMigrationId *string `json:"BackupMigrationId,omitnil,omitempty" name:"BackupMigrationId"`

	// Task name
	MigrationName *string `json:"MigrationName,omitnil,omitempty" name:"MigrationName"`

	// Migration task restoration type: FULL,FULL_LOG,FULL_DIFF
	RecoveryType *string `json:"RecoveryType,omitnil,omitempty" name:"RecoveryType"`

	// COS_URL: the backup is stored in user’s Cloud Object Storage, with URL provided. COS_UPLOAD: the backup is stored in the application’s Cloud Object Storage and needs to be uploaded by the user.
	UploadType *string `json:"UploadType,omitnil,omitempty" name:"UploadType"`

	// If the UploadType is COS_URL, fill in URL here. If the UploadType is COS_UPLOAD, fill in the name of the backup file here. Only 1 backup file is supported, but a backup file can involve multiple databases.
	BackupFiles []*string `json:"BackupFiles,omitnil,omitempty" name:"BackupFiles"`

	// Name set of databases to be renamed
	DBRename []*RenameRestoreDatabase `json:"DBRename,omitnil,omitempty" name:"DBRename"`
}

type ModifyBackupMigrationRequest struct {
	*tchttp.BaseRequest
	
	// ID of imported target instance
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Backup import task ID, which is returned through the API CreateBackupMigration
	BackupMigrationId *string `json:"BackupMigrationId,omitnil,omitempty" name:"BackupMigrationId"`

	// Task name
	MigrationName *string `json:"MigrationName,omitnil,omitempty" name:"MigrationName"`

	// Migration task restoration type: FULL,FULL_LOG,FULL_DIFF
	RecoveryType *string `json:"RecoveryType,omitnil,omitempty" name:"RecoveryType"`

	// COS_URL: the backup is stored in user’s Cloud Object Storage, with URL provided. COS_UPLOAD: the backup is stored in the application’s Cloud Object Storage and needs to be uploaded by the user.
	UploadType *string `json:"UploadType,omitnil,omitempty" name:"UploadType"`

	// If the UploadType is COS_URL, fill in URL here. If the UploadType is COS_UPLOAD, fill in the name of the backup file here. Only 1 backup file is supported, but a backup file can involve multiple databases.
	BackupFiles []*string `json:"BackupFiles,omitnil,omitempty" name:"BackupFiles"`

	// Name set of databases to be renamed
	DBRename []*RenameRestoreDatabase `json:"DBRename,omitnil,omitempty" name:"DBRename"`
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
	delete(f, "DBRename")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyBackupMigrationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBackupMigrationResponseParams struct {
	// Backup import task ID
	BackupMigrationId *string `json:"BackupMigrationId,omitnil,omitempty" name:"BackupMigrationId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyBackupMigrationResponse struct {
	*tchttp.BaseResponse
	Response *ModifyBackupMigrationResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyBackupStrategyRequestParams struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Backup type. Valid values: `weekly` (when length(BackupDay) <=7 && length(BackupDay) >=2), `daily` (when length(BackupDay)=1). Default value: `daily`.
	BackupType *string `json:"BackupType,omitnil,omitempty" name:"BackupType"`

	// Backup time. Value range: an integer from 0 to 23.
	BackupTime *uint64 `json:"BackupTime,omitnil,omitempty" name:"BackupTime"`

	// Backup interval in days when the `BackupType` is `daily`. Valid value: 1.
	BackupDay *uint64 `json:"BackupDay,omitnil,omitempty" name:"BackupDay"`

	// Backup mode. Valid values: `master_pkg` (archive the backup files of the primary node), `master_no_pkg` (do not archive the backup files of the primary node), `slave_pkg` (archive the backup files of the replica node), `slave_no_pkg` (do not archive the backup files of the replica node). Backup files of the replica node are supported only when Always On disaster recovery is enabled.
	BackupModel *string `json:"BackupModel,omitnil,omitempty" name:"BackupModel"`

	// The days of the week on which backup will be performed when “BackupType” is `weekly`. If data backup retention period is less than 7 days, the values will be 1-7, indicating that backup will be performed everyday by default; if data backup retention period is greater than or equal to 7 days, the values will be at least any two days, indicating that backup will be performed at least twice in a week by default.
	BackupCycle []*uint64 `json:"BackupCycle,omitnil,omitempty" name:"BackupCycle"`

	// Data (log) backup retention period. Value range: 3-1830 days, default value: 7 days.
	BackupSaveDays *uint64 `json:"BackupSaveDays,omitnil,omitempty" name:"BackupSaveDays"`

	// Archive backup status. Valid values: `enable` (enabled); `disable` (disabled). Default value: `disable`.
	RegularBackupEnable *string `json:"RegularBackupEnable,omitnil,omitempty" name:"RegularBackupEnable"`

	// Archive backup retention days. Value range: 90–3650 days. Default value: 365 days.
	RegularBackupSaveDays *uint64 `json:"RegularBackupSaveDays,omitnil,omitempty" name:"RegularBackupSaveDays"`

	// Archive backup policy. Valid values: `years` (yearly); `quarters (quarterly); `months` (monthly); Default value: `months`.
	RegularBackupStrategy *string `json:"RegularBackupStrategy,omitnil,omitempty" name:"RegularBackupStrategy"`

	// The number of retained archive backups. Default value: `1`.
	RegularBackupCounts *uint64 `json:"RegularBackupCounts,omitnil,omitempty" name:"RegularBackupCounts"`

	// Archive backup start date in YYYY-MM-DD format, which is the current time by default.
	RegularBackupStartTime *string `json:"RegularBackupStartTime,omitnil,omitempty" name:"RegularBackupStartTime"`
}

type ModifyBackupStrategyRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Backup type. Valid values: `weekly` (when length(BackupDay) <=7 && length(BackupDay) >=2), `daily` (when length(BackupDay)=1). Default value: `daily`.
	BackupType *string `json:"BackupType,omitnil,omitempty" name:"BackupType"`

	// Backup time. Value range: an integer from 0 to 23.
	BackupTime *uint64 `json:"BackupTime,omitnil,omitempty" name:"BackupTime"`

	// Backup interval in days when the `BackupType` is `daily`. Valid value: 1.
	BackupDay *uint64 `json:"BackupDay,omitnil,omitempty" name:"BackupDay"`

	// Backup mode. Valid values: `master_pkg` (archive the backup files of the primary node), `master_no_pkg` (do not archive the backup files of the primary node), `slave_pkg` (archive the backup files of the replica node), `slave_no_pkg` (do not archive the backup files of the replica node). Backup files of the replica node are supported only when Always On disaster recovery is enabled.
	BackupModel *string `json:"BackupModel,omitnil,omitempty" name:"BackupModel"`

	// The days of the week on which backup will be performed when “BackupType” is `weekly`. If data backup retention period is less than 7 days, the values will be 1-7, indicating that backup will be performed everyday by default; if data backup retention period is greater than or equal to 7 days, the values will be at least any two days, indicating that backup will be performed at least twice in a week by default.
	BackupCycle []*uint64 `json:"BackupCycle,omitnil,omitempty" name:"BackupCycle"`

	// Data (log) backup retention period. Value range: 3-1830 days, default value: 7 days.
	BackupSaveDays *uint64 `json:"BackupSaveDays,omitnil,omitempty" name:"BackupSaveDays"`

	// Archive backup status. Valid values: `enable` (enabled); `disable` (disabled). Default value: `disable`.
	RegularBackupEnable *string `json:"RegularBackupEnable,omitnil,omitempty" name:"RegularBackupEnable"`

	// Archive backup retention days. Value range: 90–3650 days. Default value: 365 days.
	RegularBackupSaveDays *uint64 `json:"RegularBackupSaveDays,omitnil,omitempty" name:"RegularBackupSaveDays"`

	// Archive backup policy. Valid values: `years` (yearly); `quarters (quarterly); `months` (monthly); Default value: `months`.
	RegularBackupStrategy *string `json:"RegularBackupStrategy,omitnil,omitempty" name:"RegularBackupStrategy"`

	// The number of retained archive backups. Default value: `1`.
	RegularBackupCounts *uint64 `json:"RegularBackupCounts,omitnil,omitempty" name:"RegularBackupCounts"`

	// Archive backup start date in YYYY-MM-DD format, which is the current time by default.
	RegularBackupStartTime *string `json:"RegularBackupStartTime,omitnil,omitempty" name:"RegularBackupStartTime"`
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
	delete(f, "BackupModel")
	delete(f, "BackupCycle")
	delete(f, "BackupSaveDays")
	delete(f, "RegularBackupEnable")
	delete(f, "RegularBackupSaveDays")
	delete(f, "RegularBackupStrategy")
	delete(f, "RegularBackupCounts")
	delete(f, "RegularBackupStartTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyBackupStrategyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBackupStrategyResponseParams struct {
	// Returned error code.
	//
	// Deprecated: Errno is deprecated.
	Errno *int64 `json:"Errno,omitnil,omitempty" name:"Errno"`

	// Returned error message.
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`


	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyBackupStrategyResponse struct {
	*tchttp.BaseResponse
	Response *ModifyBackupStrategyResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyDBEncryptAttributesRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// A parameter used to enable or disable TDE of the database
	DBTDEEncrypt []*DBTDEEncrypt `json:"DBTDEEncrypt,omitnil,omitempty" name:"DBTDEEncrypt"`
}

type ModifyDBEncryptAttributesRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// A parameter used to enable or disable TDE of the database
	DBTDEEncrypt []*DBTDEEncrypt `json:"DBTDEEncrypt,omitnil,omitempty" name:"DBTDEEncrypt"`
}

func (r *ModifyDBEncryptAttributesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBEncryptAttributesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "DBTDEEncrypt")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDBEncryptAttributesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBEncryptAttributesResponseParams struct {
	// Task flow ID
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDBEncryptAttributesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDBEncryptAttributesResponseParams `json:"Response"`
}

func (r *ModifyDBEncryptAttributesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBEncryptAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceNameRequestParams struct {
	// Database instance ID in the format of mssql-njj2mtpl
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// New name of database instance
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`
}

type ModifyDBInstanceNameRequest struct {
	*tchttp.BaseRequest
	
	// Database instance ID in the format of mssql-njj2mtpl
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// New name of database instance
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`
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

// Predefined struct for user
type ModifyDBInstanceNameResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDBInstanceNameResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDBInstanceNameResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyDBInstanceNetworkRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// ID of the new VPC
	NewVpcId *string `json:"NewVpcId,omitnil,omitempty" name:"NewVpcId"`

	// ID of the new subnet
	NewSubnetId *string `json:"NewSubnetId,omitnil,omitempty" name:"NewSubnetId"`

	// Retention period (in hours) of the original VIP. Value range: `0-168`. Default value: `0`, indicating the original VIP is released immediately.
	OldIpRetainTime *int64 `json:"OldIpRetainTime,omitnil,omitempty" name:"OldIpRetainTime"`

	// New VIP
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`


	DRNetwork *uint64 `json:"DRNetwork,omitnil,omitempty" name:"DRNetwork"`
}

type ModifyDBInstanceNetworkRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// ID of the new VPC
	NewVpcId *string `json:"NewVpcId,omitnil,omitempty" name:"NewVpcId"`

	// ID of the new subnet
	NewSubnetId *string `json:"NewSubnetId,omitnil,omitempty" name:"NewSubnetId"`

	// Retention period (in hours) of the original VIP. Value range: `0-168`. Default value: `0`, indicating the original VIP is released immediately.
	OldIpRetainTime *int64 `json:"OldIpRetainTime,omitnil,omitempty" name:"OldIpRetainTime"`

	// New VIP
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	DRNetwork *uint64 `json:"DRNetwork,omitnil,omitempty" name:"DRNetwork"`
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
	delete(f, "Vip")
	delete(f, "DRNetwork")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDBInstanceNetworkRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceNetworkResponseParams struct {
	// ID of the instance network changing task. You can use the [DescribeFlowStatus](https://intl.cloud.tencent.com/document/product/238/19967?from_cn_redirect=1) API to query the task status.
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDBInstanceNetworkResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDBInstanceNetworkResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyDBInstanceProjectRequestParams struct {
	// Array of instance IDs in the format of mssql-j8kv137v
	InstanceIdSet []*string `json:"InstanceIdSet,omitnil,omitempty" name:"InstanceIdSet"`

	// Project ID. If this parameter is 0, the default project will be used
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type ModifyDBInstanceProjectRequest struct {
	*tchttp.BaseRequest
	
	// Array of instance IDs in the format of mssql-j8kv137v
	InstanceIdSet []*string `json:"InstanceIdSet,omitnil,omitempty" name:"InstanceIdSet"`

	// Project ID. If this parameter is 0, the default project will be used
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
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

// Predefined struct for user
type ModifyDBInstanceProjectResponseParams struct {
	// Number of successfully modified instances
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDBInstanceProjectResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDBInstanceProjectResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyDBNameRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Old database name
	OldDBName *string `json:"OldDBName,omitnil,omitempty" name:"OldDBName"`

	// New name of database
	NewDBName *string `json:"NewDBName,omitnil,omitempty" name:"NewDBName"`
}

type ModifyDBNameRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Old database name
	OldDBName *string `json:"OldDBName,omitnil,omitempty" name:"OldDBName"`

	// New name of database
	NewDBName *string `json:"NewDBName,omitnil,omitempty" name:"NewDBName"`
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

// Predefined struct for user
type ModifyDBNameResponseParams struct {
	// Task flow ID
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDBNameResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDBNameResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyDBRemarkRequestParams struct {
	// Instance ID in the format of mssql-rljoi3bf
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Array of database names and remarks, where each element contains a database name and the corresponding remarks
	DBRemarks []*DBRemark `json:"DBRemarks,omitnil,omitempty" name:"DBRemarks"`
}

type ModifyDBRemarkRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of mssql-rljoi3bf
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Array of database names and remarks, where each element contains a database name and the corresponding remarks
	DBRemarks []*DBRemark `json:"DBRemarks,omitnil,omitempty" name:"DBRemarks"`
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

// Predefined struct for user
type ModifyDBRemarkResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDBRemarkResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDBRemarkResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyDReadableRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Operation type. Valid values: enable - enabling the read-only mode of the replica server; disable - disabling the read-only mode of the replica server
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Replica server network ID, which will be consistent with the primary instance by default if left blank
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Replica server subnet ID, which will be consistent with the primary instance by default if left blank
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Specified replica server read-only VIP, which will be assigned automatically if left blank
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`
}

type ModifyDReadableRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Operation type. Valid values: enable - enabling the read-only mode of the replica server; disable - disabling the read-only mode of the replica server
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Replica server network ID, which will be consistent with the primary instance by default if left blank
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Replica server subnet ID, which will be consistent with the primary instance by default if left blank
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Specified replica server read-only VIP, which will be assigned automatically if left blank
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`
}

func (r *ModifyDReadableRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDReadableRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Type")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "Vip")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDReadableRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDReadableResponseParams struct {
	// Task ID
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDReadableResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDReadableResponseParams `json:"Response"`
}

func (r *ModifyDReadableResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDReadableResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDatabaseCDCRequestParams struct {
	// Array of database names
	DBNames []*string `json:"DBNames,omitnil,omitempty" name:"DBNames"`

	// Enable or disable CDC. Valid values: `enable`, `disable`
	ModifyType *string `json:"ModifyType,omitnil,omitempty" name:"ModifyType"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type ModifyDatabaseCDCRequest struct {
	*tchttp.BaseRequest
	
	// Array of database names
	DBNames []*string `json:"DBNames,omitnil,omitempty" name:"DBNames"`

	// Enable or disable CDC. Valid values: `enable`, `disable`
	ModifyType *string `json:"ModifyType,omitnil,omitempty" name:"ModifyType"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *ModifyDatabaseCDCRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDatabaseCDCRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBNames")
	delete(f, "ModifyType")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDatabaseCDCRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDatabaseCDCResponseParams struct {
	// Task ID
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDatabaseCDCResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDatabaseCDCResponseParams `json:"Response"`
}

func (r *ModifyDatabaseCDCResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDatabaseCDCResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDatabaseCTRequestParams struct {
	// Array of database names
	DBNames []*string `json:"DBNames,omitnil,omitempty" name:"DBNames"`

	// Enable or disable CT. Valid values: `enable`, `disable`
	ModifyType *string `json:"ModifyType,omitnil,omitempty" name:"ModifyType"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Retention period (in days) of change tracking information when CT is enabled. Value range: 3-30. Default value: `3`
	ChangeRetentionDay *int64 `json:"ChangeRetentionDay,omitnil,omitempty" name:"ChangeRetentionDay"`
}

type ModifyDatabaseCTRequest struct {
	*tchttp.BaseRequest
	
	// Array of database names
	DBNames []*string `json:"DBNames,omitnil,omitempty" name:"DBNames"`

	// Enable or disable CT. Valid values: `enable`, `disable`
	ModifyType *string `json:"ModifyType,omitnil,omitempty" name:"ModifyType"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Retention period (in days) of change tracking information when CT is enabled. Value range: 3-30. Default value: `3`
	ChangeRetentionDay *int64 `json:"ChangeRetentionDay,omitnil,omitempty" name:"ChangeRetentionDay"`
}

func (r *ModifyDatabaseCTRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDatabaseCTRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBNames")
	delete(f, "ModifyType")
	delete(f, "InstanceId")
	delete(f, "ChangeRetentionDay")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDatabaseCTRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDatabaseCTResponseParams struct {
	// Task ID
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDatabaseCTResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDatabaseCTResponseParams `json:"Response"`
}

func (r *ModifyDatabaseCTResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDatabaseCTResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDatabaseMdfRequestParams struct {
	// Array of database names
	DBNames []*string `json:"DBNames,omitnil,omitempty" name:"DBNames"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type ModifyDatabaseMdfRequest struct {
	*tchttp.BaseRequest
	
	// Array of database names
	DBNames []*string `json:"DBNames,omitnil,omitempty" name:"DBNames"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *ModifyDatabaseMdfRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDatabaseMdfRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBNames")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDatabaseMdfRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDatabaseMdfResponseParams struct {
	// Task ID
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDatabaseMdfResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDatabaseMdfResponseParams `json:"Response"`
}

func (r *ModifyDatabaseMdfResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDatabaseMdfResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyIncrementalMigrationRequestParams struct {
	// ID of imported target instance
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Backup import task ID, which is returned through the API CreateBackupMigration
	BackupMigrationId *string `json:"BackupMigrationId,omitnil,omitempty" name:"BackupMigrationId"`

	// Incremental backup import task ID, which is returned through the `CreateIncrementalMigration` API.
	IncrementalMigrationId *string `json:"IncrementalMigrationId,omitnil,omitempty" name:"IncrementalMigrationId"`

	// Whether to restore backups. Valid values: `NO`, `YES`. If this parameter is not specified, current settings will be applied.
	IsRecovery *string `json:"IsRecovery,omitnil,omitempty" name:"IsRecovery"`

	// If the UploadType is COS_URL, fill in URL here. If the UploadType is COS_UPLOAD, fill in the name of the backup file here. Only 1 backup file is supported, but a backup file can involve multiple databases.
	BackupFiles []*string `json:"BackupFiles,omitnil,omitempty" name:"BackupFiles"`
}

type ModifyIncrementalMigrationRequest struct {
	*tchttp.BaseRequest
	
	// ID of imported target instance
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Backup import task ID, which is returned through the API CreateBackupMigration
	BackupMigrationId *string `json:"BackupMigrationId,omitnil,omitempty" name:"BackupMigrationId"`

	// Incremental backup import task ID, which is returned through the `CreateIncrementalMigration` API.
	IncrementalMigrationId *string `json:"IncrementalMigrationId,omitnil,omitempty" name:"IncrementalMigrationId"`

	// Whether to restore backups. Valid values: `NO`, `YES`. If this parameter is not specified, current settings will be applied.
	IsRecovery *string `json:"IsRecovery,omitnil,omitempty" name:"IsRecovery"`

	// If the UploadType is COS_URL, fill in URL here. If the UploadType is COS_UPLOAD, fill in the name of the backup file here. Only 1 backup file is supported, but a backup file can involve multiple databases.
	BackupFiles []*string `json:"BackupFiles,omitnil,omitempty" name:"BackupFiles"`
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

// Predefined struct for user
type ModifyIncrementalMigrationResponseParams struct {
	// ID of an incremental backup import task
	IncrementalMigrationId *string `json:"IncrementalMigrationId,omitnil,omitempty" name:"IncrementalMigrationId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyIncrementalMigrationResponse struct {
	*tchttp.BaseResponse
	Response *ModifyIncrementalMigrationResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyInstanceEncryptAttributesRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Certificate ownership. Valid values: `self` (certificate of this account), `others` (certificate of the other account). Default value: `self`.
	CertificateAttribution *string `json:"CertificateAttribution,omitnil,omitempty" name:"CertificateAttribution"`

	// ID of the other referenced root account, which is required when `CertificateAttribution` is `others`.
	QuoteUin *string `json:"QuoteUin,omitnil,omitempty" name:"QuoteUin"`
}

type ModifyInstanceEncryptAttributesRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Certificate ownership. Valid values: `self` (certificate of this account), `others` (certificate of the other account). Default value: `self`.
	CertificateAttribution *string `json:"CertificateAttribution,omitnil,omitempty" name:"CertificateAttribution"`

	// ID of the other referenced root account, which is required when `CertificateAttribution` is `others`.
	QuoteUin *string `json:"QuoteUin,omitnil,omitempty" name:"QuoteUin"`
}

func (r *ModifyInstanceEncryptAttributesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceEncryptAttributesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "CertificateAttribution")
	delete(f, "QuoteUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstanceEncryptAttributesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceEncryptAttributesResponseParams struct {
	// Task flow ID
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyInstanceEncryptAttributesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstanceEncryptAttributesResponseParams `json:"Response"`
}

func (r *ModifyInstanceEncryptAttributesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceEncryptAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceParamRequestParams struct {
	// Instance ID list.
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// List of modified parameters. Each list element has two fields: `Name` and `CurrentValue`. Set `Name` to the parameter name and `CurrentValue` to the new value after modification. <b>Note</b>: if the instance needs to be <b>restarted</b> for the modified parameter to take effect, it will be <b>restarted</b> immediately or during the maintenance time. Before you modify a parameter, you can use the `DescribeInstanceParams` API to query whether the instance needs to be restarted.
	ParamList []*Parameter `json:"ParamList,omitnil,omitempty" name:"ParamList"`

	// When to execute the parameter modification task. Valid values: `0` (execute immediately), `1` (execute during maintenance time). Default value: `0`.
	WaitSwitch *int64 `json:"WaitSwitch,omitnil,omitempty" name:"WaitSwitch"`
}

type ModifyInstanceParamRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID list.
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// List of modified parameters. Each list element has two fields: `Name` and `CurrentValue`. Set `Name` to the parameter name and `CurrentValue` to the new value after modification. <b>Note</b>: if the instance needs to be <b>restarted</b> for the modified parameter to take effect, it will be <b>restarted</b> immediately or during the maintenance time. Before you modify a parameter, you can use the `DescribeInstanceParams` API to query whether the instance needs to be restarted.
	ParamList []*Parameter `json:"ParamList,omitnil,omitempty" name:"ParamList"`

	// When to execute the parameter modification task. Valid values: `0` (execute immediately), `1` (execute during maintenance time). Default value: `0`.
	WaitSwitch *int64 `json:"WaitSwitch,omitnil,omitempty" name:"WaitSwitch"`
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
	delete(f, "WaitSwitch")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstanceParamRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceParamResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyInstanceParamResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstanceParamResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyMigrationRequestParams struct {
	// Migration task ID
	MigrateId *uint64 `json:"MigrateId,omitnil,omitempty" name:"MigrateId"`

	// New name of migration task. If this parameter is left empty, no modification will be made
	MigrateName *string `json:"MigrateName,omitnil,omitempty" name:"MigrateName"`

	// New migration type (1: structure migration, 2: data migration, 3: incremental sync). If this parameter is left empty, no modification will be made
	MigrateType *uint64 `json:"MigrateType,omitnil,omitempty" name:"MigrateType"`

	// Migration source type. 1: TencentDB for SQL Server, 2: CVM-based self-created SQL Server database; 3: SQL Server backup restoration, 4: SQL Server backup restoration (in COS mode). If this parameter is left empty, no modification will be made
	SourceType *uint64 `json:"SourceType,omitnil,omitempty" name:"SourceType"`

	// Migration source. If this parameter is left empty, no modification will be made
	Source *MigrateSource `json:"Source,omitnil,omitempty" name:"Source"`

	// Migration target. If this parameter is left empty, no modification will be made
	Target *MigrateTarget `json:"Target,omitnil,omitempty" name:"Target"`

	// Database objects to be migrated. This parameter is not used for offline migration (SourceType=4 or SourceType=5). If it left empty, no modification will be made
	MigrateDBSet []*MigrateDB `json:"MigrateDBSet,omitnil,omitempty" name:"MigrateDBSet"`
}

type ModifyMigrationRequest struct {
	*tchttp.BaseRequest
	
	// Migration task ID
	MigrateId *uint64 `json:"MigrateId,omitnil,omitempty" name:"MigrateId"`

	// New name of migration task. If this parameter is left empty, no modification will be made
	MigrateName *string `json:"MigrateName,omitnil,omitempty" name:"MigrateName"`

	// New migration type (1: structure migration, 2: data migration, 3: incremental sync). If this parameter is left empty, no modification will be made
	MigrateType *uint64 `json:"MigrateType,omitnil,omitempty" name:"MigrateType"`

	// Migration source type. 1: TencentDB for SQL Server, 2: CVM-based self-created SQL Server database; 3: SQL Server backup restoration, 4: SQL Server backup restoration (in COS mode). If this parameter is left empty, no modification will be made
	SourceType *uint64 `json:"SourceType,omitnil,omitempty" name:"SourceType"`

	// Migration source. If this parameter is left empty, no modification will be made
	Source *MigrateSource `json:"Source,omitnil,omitempty" name:"Source"`

	// Migration target. If this parameter is left empty, no modification will be made
	Target *MigrateTarget `json:"Target,omitnil,omitempty" name:"Target"`

	// Database objects to be migrated. This parameter is not used for offline migration (SourceType=4 or SourceType=5). If it left empty, no modification will be made
	MigrateDBSet []*MigrateDB `json:"MigrateDBSet,omitnil,omitempty" name:"MigrateDBSet"`
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

// Predefined struct for user
type ModifyMigrationResponseParams struct {
	// Migration task ID
	MigrateId *uint64 `json:"MigrateId,omitnil,omitempty" name:"MigrateId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyMigrationResponse struct {
	*tchttp.BaseResponse
	Response *ModifyMigrationResponseParams `json:"Response"`
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

type OldVip struct {
	// Unrecovered old IP addresses
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// IP recovery time
	RecycleTime *string `json:"RecycleTime,omitnil,omitempty" name:"RecycleTime"`

	// Old IP retention time (hours)
	// Note: This field may return null, indicating that no valid values can be obtained.
	OldIpRetainTime *int64 `json:"OldIpRetainTime,omitnil,omitempty" name:"OldIpRetainTime"`
}

// Predefined struct for user
type OpenInterCommunicationRequestParams struct {
	// IDs of instances with interwoking group enabled
	InstanceIdSet []*string `json:"InstanceIdSet,omitnil,omitempty" name:"InstanceIdSet"`
}

type OpenInterCommunicationRequest struct {
	*tchttp.BaseRequest
	
	// IDs of instances with interwoking group enabled
	InstanceIdSet []*string `json:"InstanceIdSet,omitnil,omitempty" name:"InstanceIdSet"`
}

func (r *OpenInterCommunicationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenInterCommunicationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIdSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OpenInterCommunicationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenInterCommunicationResponseParams struct {
	// IDs of instance and async task
	InterInstanceFlowSet []*InterInstanceFlow `json:"InterInstanceFlowSet,omitnil,omitempty" name:"InterInstanceFlowSet"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type OpenInterCommunicationResponse struct {
	*tchttp.BaseResponse
	Response *OpenInterCommunicationResponseParams `json:"Response"`
}

func (r *OpenInterCommunicationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenInterCommunicationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ParamRecord struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Parameter name
	ParamName *string `json:"ParamName,omitnil,omitempty" name:"ParamName"`

	// Parameter value before modification
	OldValue *string `json:"OldValue,omitnil,omitempty" name:"OldValue"`

	// Parameter value after modification
	NewValue *string `json:"NewValue,omitnil,omitempty" name:"NewValue"`

	// Parameter modification status. Valid values: `1` (initializing and waiting for modification), `2` (modification succeed), `3` (modification failed), `4` (modifying)
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Modification time
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`
}

type Parameter struct {
	// Parameter name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Parameter value
	CurrentValue *string `json:"CurrentValue,omitnil,omitempty" name:"CurrentValue"`
}

type ParameterDetail struct {
	// Parameter name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Data type of the parameter. Valid values: `integer`, `enum`
	ParamType *string `json:"ParamType,omitnil,omitempty" name:"ParamType"`

	// Default value of the parameter
	Default *string `json:"Default,omitnil,omitempty" name:"Default"`

	// Parameter description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Current value of the parameter
	CurrentValue *string `json:"CurrentValue,omitnil,omitempty" name:"CurrentValue"`

	// Whether the database needs to be restarted for the modified parameter to take effect. Valid values: `0` (no),`1` (yes)
	NeedReboot *int64 `json:"NeedReboot,omitnil,omitempty" name:"NeedReboot"`

	// Maximum value of the parameter
	Max *int64 `json:"Max,omitnil,omitempty" name:"Max"`

	// Minimum value of the parameter
	Min *int64 `json:"Min,omitnil,omitempty" name:"Min"`

	// Enumerated values of the parameter
	EnumValue []*string `json:"EnumValue,omitnil,omitempty" name:"EnumValue"`

	// Parameter status. Valid values: `0` (normal), `1` (modifying)
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

// Predefined struct for user
type RecycleDBInstanceRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type RecycleDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
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

// Predefined struct for user
type RecycleDBInstanceResponseParams struct {
	// Task ID
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RecycleDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *RecycleDBInstanceResponseParams `json:"Response"`
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
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Region name
	RegionName *string `json:"RegionName,omitnil,omitempty" name:"RegionName"`

	// Numeric ID of region
	RegionId *int64 `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// Current purchasability of this region. UNAVAILABLE: not purchasable, AVAILABLE: purchasable
	RegionState *string `json:"RegionState,omitnil,omitempty" name:"RegionState"`
}

type RenameRestoreDatabase struct {
	// Database name. If the `OldName` database does not exist, a failure will be returned.
	// It can be left empty in offline migration tasks.
	OldName *string `json:"OldName,omitnil,omitempty" name:"OldName"`

	// New database name. In offline migration, `OldName` will be used if `NewName` is left empty (`OldName` and `NewName` cannot be both empty). In database cloning, `OldName` and `NewName` must be both specified and cannot have the same value.
	NewName *string `json:"NewName,omitnil,omitempty" name:"NewName"`
}

// Predefined struct for user
type ResetAccountPasswordRequestParams struct {
	// Database instance ID in the format of mssql-njj2mtpl
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Updated account password information array
	Accounts []*AccountPassword `json:"Accounts,omitnil,omitempty" name:"Accounts"`
}

type ResetAccountPasswordRequest struct {
	*tchttp.BaseRequest
	
	// Database instance ID in the format of mssql-njj2mtpl
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Updated account password information array
	Accounts []*AccountPassword `json:"Accounts,omitnil,omitempty" name:"Accounts"`
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

// Predefined struct for user
type ResetAccountPasswordResponseParams struct {
	// ID of asynchronous task flow for account password modification
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ResetAccountPasswordResponse struct {
	*tchttp.BaseResponse
	Response *ResetAccountPasswordResponseParams `json:"Response"`
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
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// Tag value
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

// Predefined struct for user
type RestartDBInstanceRequestParams struct {
	// Database instance ID in the format of mssql-njj2mtpl
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type RestartDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Database instance ID in the format of mssql-njj2mtpl
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
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

// Predefined struct for user
type RestartDBInstanceResponseParams struct {
	// Async task flow ID
	FlowId *uint64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RestartDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *RestartDBInstanceResponseParams `json:"Response"`
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

// Predefined struct for user
type RestoreInstanceRequestParams struct {
	// Instance ID in the format of mssql-j8kv137v
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Backup file ID, which can be obtained through the `Id` field in the returned value of the `DescribeBackups` API
	BackupId *int64 `json:"BackupId,omitnil,omitempty" name:"BackupId"`

	// ID of the target instance to which the backup is restored. The target instance should be under the same `APPID`. If this parameter is left empty, ID of the source instance will be used.
	TargetInstanceId *string `json:"TargetInstanceId,omitnil,omitempty" name:"TargetInstanceId"`

	// Restore the databases listed in `ReNameRestoreDatabase` and rename them after restoration. If this parameter is left empty, all databases will be restored and renamed in the default format.
	RenameRestore []*RenameRestoreDatabase `json:"RenameRestore,omitnil,omitempty" name:"RenameRestore"`

	// Rollback type. Valid values: `0` (overwriting), `1` (renaming).
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// Database to be overwritten, which is required when overwriting a rollback database.
	DBList []*string `json:"DBList,omitnil,omitempty" name:"DBList"`

	// Group ID of unarchived backup files grouped by backup task
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`
}

type RestoreInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of mssql-j8kv137v
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Backup file ID, which can be obtained through the `Id` field in the returned value of the `DescribeBackups` API
	BackupId *int64 `json:"BackupId,omitnil,omitempty" name:"BackupId"`

	// ID of the target instance to which the backup is restored. The target instance should be under the same `APPID`. If this parameter is left empty, ID of the source instance will be used.
	TargetInstanceId *string `json:"TargetInstanceId,omitnil,omitempty" name:"TargetInstanceId"`

	// Restore the databases listed in `ReNameRestoreDatabase` and rename them after restoration. If this parameter is left empty, all databases will be restored and renamed in the default format.
	RenameRestore []*RenameRestoreDatabase `json:"RenameRestore,omitnil,omitempty" name:"RenameRestore"`

	// Rollback type. Valid values: `0` (overwriting), `1` (renaming).
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// Database to be overwritten, which is required when overwriting a rollback database.
	DBList []*string `json:"DBList,omitnil,omitempty" name:"DBList"`

	// Group ID of unarchived backup files grouped by backup task
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`
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
	delete(f, "Type")
	delete(f, "DBList")
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RestoreInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RestoreInstanceResponseParams struct {
	// Async flow task ID, which can be used to call the `DescribeFlowStatus` API to get the task execution status
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RestoreInstanceResponse struct {
	*tchttp.BaseResponse
	Response *RestoreInstanceResponseParams `json:"Response"`
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

// Predefined struct for user
type RollbackInstanceRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Rollback type. 0: the database rolled back overwrites the original database; 1: the database rolled back is renamed and does not overwrite the original database
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// Target time point for rollback
	Time *string `json:"Time,omitnil,omitempty" name:"Time"`

	// Database to be rolled back
	DBs []*string `json:"DBs,omitnil,omitempty" name:"DBs"`

	// ID of the target instance to which the backup is restored. The target instance should be under the same `APPID`. If this parameter is left empty, ID of the source instance will be used.
	TargetInstanceId *string `json:"TargetInstanceId,omitnil,omitempty" name:"TargetInstanceId"`

	// Rename the databases listed in `ReNameRestoreDatabase`. This parameter takes effect only when `Type = 1` which indicates that backup rollback supports renaming databases. If it is left empty, databases will be renamed in the default format and the `DBs` parameter specifies the databases to be restored.
	RenameRestore []*RenameRestoreDatabase `json:"RenameRestore,omitnil,omitempty" name:"RenameRestore"`
}

type RollbackInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Rollback type. 0: the database rolled back overwrites the original database; 1: the database rolled back is renamed and does not overwrite the original database
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// Target time point for rollback
	Time *string `json:"Time,omitnil,omitempty" name:"Time"`

	// Database to be rolled back
	DBs []*string `json:"DBs,omitnil,omitempty" name:"DBs"`

	// ID of the target instance to which the backup is restored. The target instance should be under the same `APPID`. If this parameter is left empty, ID of the source instance will be used.
	TargetInstanceId *string `json:"TargetInstanceId,omitnil,omitempty" name:"TargetInstanceId"`

	// Rename the databases listed in `ReNameRestoreDatabase`. This parameter takes effect only when `Type = 1` which indicates that backup rollback supports renaming databases. If it is left empty, databases will be renamed in the default format and the `DBs` parameter specifies the databases to be restored.
	RenameRestore []*RenameRestoreDatabase `json:"RenameRestore,omitnil,omitempty" name:"RenameRestore"`
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
	delete(f, "Time")
	delete(f, "DBs")
	delete(f, "TargetInstanceId")
	delete(f, "RenameRestore")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RollbackInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RollbackInstanceResponseParams struct {
	// The async job ID
	FlowId *uint64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RollbackInstanceResponse struct {
	*tchttp.BaseResponse
	Response *RollbackInstanceResponseParams `json:"Response"`
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

// Predefined struct for user
type RunMigrationRequestParams struct {
	// Migration task ID
	MigrateId *uint64 `json:"MigrateId,omitnil,omitempty" name:"MigrateId"`
}

type RunMigrationRequest struct {
	*tchttp.BaseRequest
	
	// Migration task ID
	MigrateId *uint64 `json:"MigrateId,omitnil,omitempty" name:"MigrateId"`
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

// Predefined struct for user
type RunMigrationResponseParams struct {
	// After the migration task starts, the flow ID will be returned
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RunMigrationResponse struct {
	*tchttp.BaseResponse
	Response *RunMigrationResponseParams `json:"Response"`
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

type SSLConfig struct {

	Encryption *string `json:"Encryption,omitnil,omitempty" name:"Encryption"`


	SSLValidityPeriod *string `json:"SSLValidityPeriod,omitnil,omitempty" name:"SSLValidityPeriod"`


	SSLValidity *uint64 `json:"SSLValidity,omitnil,omitempty" name:"SSLValidity"`
}

type SlaveZones struct {
	// Replica AZ region code
	SlaveZone *string `json:"SlaveZone,omitnil,omitempty" name:"SlaveZone"`

	// Replica AZ
	SlaveZoneName *string `json:"SlaveZoneName,omitnil,omitempty" name:"SlaveZoneName"`
}

type SlowLog struct {
	// Unique ID of slow query log file
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// File generation start time
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// File generation end time
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// File size in KB
	Size *int64 `json:"Size,omitnil,omitempty" name:"Size"`

	// Number of logs in file
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// Download address for private network
	InternalAddr *string `json:"InternalAddr,omitnil,omitempty" name:"InternalAddr"`

	// Download address for public network
	ExternalAddr *string `json:"ExternalAddr,omitnil,omitempty" name:"ExternalAddr"`

	// Status (1: success, 2: failure)
	// Note: this field may return null, indicating that no valid values can be obtained.
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type SlowlogInfo struct {
	// Unique ID of slow query log file
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// File generation start time
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// File generation end time
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// File size in KB
	Size *int64 `json:"Size,omitnil,omitempty" name:"Size"`

	// Number of logs in file
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// Download address for private network
	InternalAddr *string `json:"InternalAddr,omitnil,omitempty" name:"InternalAddr"`

	// Download address for public network
	ExternalAddr *string `json:"ExternalAddr,omitnil,omitempty" name:"ExternalAddr"`

	// Status (1: success, 2: failure)
	// Note: this field may return null, indicating that no valid values can be obtained.
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type SpecInfo struct {
	// Instance specification ID. The `SpecId` returned by `DescribeZones` together with the purchasable specification information returned by `DescribeProductConfig` can be used to find out what specifications can be purchased in a specified AZ
	SpecId *int64 `json:"SpecId,omitnil,omitempty" name:"SpecId"`

	// Model ID
	MachineType *string `json:"MachineType,omitnil,omitempty" name:"MachineType"`

	// Model name
	MachineTypeName *string `json:"MachineTypeName,omitnil,omitempty" name:"MachineTypeName"`

	// Database version information. Valid values: 2008R2 (SQL Server 2008 Enterprise), 2012SP3 (SQL Server 2012 Enterprise), 2016SP1 (SQL Server 2016 Enterprise), 201602 (SQL Server 2016 Standard), 2017 (SQL Server 2017 Enterprise)
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// Version name corresponding to the `Version` field
	VersionName *string `json:"VersionName,omitnil,omitempty" name:"VersionName"`

	// Memory size in GB
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// Number of CPU cores
	CPU *int64 `json:"CPU,omitnil,omitempty" name:"CPU"`

	// Minimum disk size under this specification in GB
	MinStorage *int64 `json:"MinStorage,omitnil,omitempty" name:"MinStorage"`

	// Maximum disk size under this specification in GB
	MaxStorage *int64 `json:"MaxStorage,omitnil,omitempty" name:"MaxStorage"`

	// QPS of this specification
	QPS *int64 `json:"QPS,omitnil,omitempty" name:"QPS"`

	// Description of this specification
	SuitInfo *string `json:"SuitInfo,omitnil,omitempty" name:"SuitInfo"`

	// Pid of this specification
	Pid *int64 `json:"Pid,omitnil,omitempty" name:"Pid"`

	// Pay-as-you-go Pid list corresponding to this specification
	// Note: this field may return null, indicating that no valid values can be obtained.
	PostPid []*int64 `json:"PostPid,omitnil,omitempty" name:"PostPid"`

	// Billing mode under this specification. POST: pay-as-you-go
	PayModeStatus *string `json:"PayModeStatus,omitnil,omitempty" name:"PayModeStatus"`

	// Instance type. Valid values: HA (High-Availability Edition, including dual-server high availability and AlwaysOn cluster), RO (read-only replica), SI (Basic Edition)
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// Whether multi-AZ deployment is supported. Valid values: MultiZones (only multi-AZ deployment is supported), SameZones (only single-AZ deployment is supported), ALL (both deployments are supported)
	MultiZonesStatus *string `json:"MultiZonesStatus,omitnil,omitempty" name:"MultiZonesStatus"`
}

// Predefined struct for user
type StartBackupMigrationRequestParams struct {
	// ID of imported target instance
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Backup import task ID, which is returned through the API CreateBackupMigration
	BackupMigrationId *string `json:"BackupMigrationId,omitnil,omitempty" name:"BackupMigrationId"`
}

type StartBackupMigrationRequest struct {
	*tchttp.BaseRequest
	
	// ID of imported target instance
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Backup import task ID, which is returned through the API CreateBackupMigration
	BackupMigrationId *string `json:"BackupMigrationId,omitnil,omitempty" name:"BackupMigrationId"`
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

// Predefined struct for user
type StartBackupMigrationResponseParams struct {
	// Task ID
	FlowId *uint64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StartBackupMigrationResponse struct {
	*tchttp.BaseResponse
	Response *StartBackupMigrationResponseParams `json:"Response"`
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

// Predefined struct for user
type StartIncrementalMigrationRequestParams struct {
	// ID of imported target instance
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Backup import task ID, which is returned through the API CreateBackupMigration
	BackupMigrationId *string `json:"BackupMigrationId,omitnil,omitempty" name:"BackupMigrationId"`

	// ID of an incremental backup import task
	IncrementalMigrationId *string `json:"IncrementalMigrationId,omitnil,omitempty" name:"IncrementalMigrationId"`
}

type StartIncrementalMigrationRequest struct {
	*tchttp.BaseRequest
	
	// ID of imported target instance
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Backup import task ID, which is returned through the API CreateBackupMigration
	BackupMigrationId *string `json:"BackupMigrationId,omitnil,omitempty" name:"BackupMigrationId"`

	// ID of an incremental backup import task
	IncrementalMigrationId *string `json:"IncrementalMigrationId,omitnil,omitempty" name:"IncrementalMigrationId"`
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

// Predefined struct for user
type StartIncrementalMigrationResponseParams struct {
	// Task ID
	FlowId *uint64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StartIncrementalMigrationResponse struct {
	*tchttp.BaseResponse
	Response *StartIncrementalMigrationResponseParams `json:"Response"`
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

// Predefined struct for user
type StartInstanceXEventRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Whether to start or stop an extended event
	EventConfig []*EventConfig `json:"EventConfig,omitnil,omitempty" name:"EventConfig"`
}

type StartInstanceXEventRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Whether to start or stop an extended event
	EventConfig []*EventConfig `json:"EventConfig,omitnil,omitempty" name:"EventConfig"`
}

func (r *StartInstanceXEventRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartInstanceXEventRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "EventConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartInstanceXEventRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartInstanceXEventResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StartInstanceXEventResponse struct {
	*tchttp.BaseResponse
	Response *StartInstanceXEventResponseParams `json:"Response"`
}

func (r *StartInstanceXEventResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartInstanceXEventResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TDEConfigAttribute struct {
	// TDE status. Valid values: `enable` (enabled), `disable` (disabled).
	Encryption *string `json:"Encryption,omitnil,omitempty" name:"Encryption"`

	// Certificate ownership. Valid values: `self` (certificate of the this account), `others` (certificate of the other account), `none` (no certificate).
	CertificateAttribution *string `json:"CertificateAttribution,omitnil,omitempty" name:"CertificateAttribution"`

	// Note: This field may return null, indicating that no valid values can be obtained.
	QuoteUin *string `json:"QuoteUin,omitnil,omitempty" name:"QuoteUin"`
}

// Predefined struct for user
type TerminateDBInstanceRequestParams struct {
	// List of instance IDs manually terminated in the format of [mssql-3l3fgqn7], which are the same as the instance IDs displayed on the TencentDB Console page
	InstanceIdSet []*string `json:"InstanceIdSet,omitnil,omitempty" name:"InstanceIdSet"`
}

type TerminateDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// List of instance IDs manually terminated in the format of [mssql-3l3fgqn7], which are the same as the instance IDs displayed on the TencentDB Console page
	InstanceIdSet []*string `json:"InstanceIdSet,omitnil,omitempty" name:"InstanceIdSet"`
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

// Predefined struct for user
type TerminateDBInstanceResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type TerminateDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *TerminateDBInstanceResponseParams `json:"Response"`
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

// Predefined struct for user
type UpgradeDBInstanceRequestParams struct {
	// Instance ID in the format of mssql-j8kv137v
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Memory size after instance upgrade in GB, which cannot be smaller than the current instance memory size
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// Storage capacity after instance upgrade in GB, which cannot be smaller than the current instance storage capacity
	Storage *int64 `json:"Storage,omitnil,omitempty" name:"Storage"`

	// Whether to automatically use vouchers. 0: no, 1: yes. Default value: 0
	AutoVoucher *int64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// Voucher ID (currently, only one voucher can be used per order)
	VoucherIds []*string `json:"VoucherIds,omitnil,omitempty" name:"VoucherIds"`

	// The number of CUP cores after the instance is upgraded.
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// Upgrade the SQL Server version. Supported versions include SQL Server 2008 Enterprise (`2008R2`), SQL Server 2012 Enterprise (`2012SP3`), etc. As the purchasable versions are region-specific, you can use the `DescribeProductConfig` API to query the information of purchasable versions in each region. Downgrading is unsupported. If this parameter is left empty, the SQL Server version will not be changed.
	DBVersion *string `json:"DBVersion,omitnil,omitempty" name:"DBVersion"`

	// Upgrade the high availability architecture from image-based disaster recovery to Always On cluster disaster recovery. This parameter is valid only for instances which support Always On high availability and run SQL Server 2017 or later. Neither downgrading to image-based disaster recovery nor upgrading from cluster disaster recovery to Always On disaster recovery is supported. If this parameter is left empty, the high availability architecture will not be changed.
	HAType *string `json:"HAType,omitnil,omitempty" name:"HAType"`

	// Change the instance deployment scheme. Valid values: `SameZones` (change to single-AZ deployment, which does not support cross-AZ disaster recovery), `MultiZones` (change to multi-AZ deployment, which supports cross-AZ disaster recovery).
	MultiZones *string `json:"MultiZones,omitnil,omitempty" name:"MultiZones"`

	// The time when configuration adjustment task is performed. Valid values: `0` (execute immediately), `1` (execute during maintenance time). Default value: `1`.
	WaitSwitch *int64 `json:"WaitSwitch,omitnil,omitempty" name:"WaitSwitch"`
}

type UpgradeDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of mssql-j8kv137v
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Memory size after instance upgrade in GB, which cannot be smaller than the current instance memory size
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// Storage capacity after instance upgrade in GB, which cannot be smaller than the current instance storage capacity
	Storage *int64 `json:"Storage,omitnil,omitempty" name:"Storage"`

	// Whether to automatically use vouchers. 0: no, 1: yes. Default value: 0
	AutoVoucher *int64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// Voucher ID (currently, only one voucher can be used per order)
	VoucherIds []*string `json:"VoucherIds,omitnil,omitempty" name:"VoucherIds"`

	// The number of CUP cores after the instance is upgraded.
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// Upgrade the SQL Server version. Supported versions include SQL Server 2008 Enterprise (`2008R2`), SQL Server 2012 Enterprise (`2012SP3`), etc. As the purchasable versions are region-specific, you can use the `DescribeProductConfig` API to query the information of purchasable versions in each region. Downgrading is unsupported. If this parameter is left empty, the SQL Server version will not be changed.
	DBVersion *string `json:"DBVersion,omitnil,omitempty" name:"DBVersion"`

	// Upgrade the high availability architecture from image-based disaster recovery to Always On cluster disaster recovery. This parameter is valid only for instances which support Always On high availability and run SQL Server 2017 or later. Neither downgrading to image-based disaster recovery nor upgrading from cluster disaster recovery to Always On disaster recovery is supported. If this parameter is left empty, the high availability architecture will not be changed.
	HAType *string `json:"HAType,omitnil,omitempty" name:"HAType"`

	// Change the instance deployment scheme. Valid values: `SameZones` (change to single-AZ deployment, which does not support cross-AZ disaster recovery), `MultiZones` (change to multi-AZ deployment, which supports cross-AZ disaster recovery).
	MultiZones *string `json:"MultiZones,omitnil,omitempty" name:"MultiZones"`

	// The time when configuration adjustment task is performed. Valid values: `0` (execute immediately), `1` (execute during maintenance time). Default value: `1`.
	WaitSwitch *int64 `json:"WaitSwitch,omitnil,omitempty" name:"WaitSwitch"`
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
	delete(f, "WaitSwitch")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpgradeDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeDBInstanceResponseParams struct {
	// Order name
	DealName *string `json:"DealName,omitnil,omitempty" name:"DealName"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpgradeDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *UpgradeDBInstanceResponseParams `json:"Response"`
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
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// AZ name
	ZoneName *string `json:"ZoneName,omitnil,omitempty" name:"ZoneName"`

	// Numeric ID of AZ
	ZoneId *int64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// ID of specification purchasable in this AZ, which, together with the returned value of the `DescribeProductConfig` API, can be used to find out the specifications currently purchasable in the AZ
	SpecId *int64 `json:"SpecId,omitnil,omitempty" name:"SpecId"`

	// Information of database versions purchasable under the current AZ and specification. Valid values: 2008R2 (SQL Server 2008 Enterprise), 2012SP3 (SQL Server 2012 Enterprise), 2016SP1 (SQL Server 2016 Enterprise), 201602 (SQL Server 2016 Standard), 2017 (SQL Server 2017 Enterprise)
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`
}