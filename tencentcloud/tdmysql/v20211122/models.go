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

package v20211122

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

type AnalysisInstanceInfo struct {
	// <p>Number of replicas</p>
	ReplicasNum *uint64 `json:"ReplicasNum,omitnil,omitempty" name:"ReplicasNum"`
}

type AnalysisRelationInfo struct {
	// <p>Source instance Id</p>
	PrimaryInstanceId *string `json:"PrimaryInstanceId,omitnil,omitempty" name:"PrimaryInstanceId"`

	// <p>Analysis engine instance Id</p>
	AnalysisInstanceId *string `json:"AnalysisInstanceId,omitnil,omitempty" name:"AnalysisInstanceId"`

	// <p>Analysis engine relationship status</p><p>Enumeration values:</p><ul><li>creating: Creating</li><li>running: Running</li><li>destroyed: Terminated</li></ul>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>Creation time.</p>
	CreateAt *string `json:"CreateAt,omitnil,omitempty" name:"CreateAt"`

	// <p>Update time.</p>
	UpdateAt *string `json:"UpdateAt,omitnil,omitempty" name:"UpdateAt"`
}

type ArchiveLogInterval struct {
	// End time.
	// Note: This field may return null, indicating that no valid values can be obtained.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Major version.
	// Note: This field may return null, indicating that no valid values can be obtained.
	MajorVersion *string `json:"MajorVersion,omitnil,omitempty" name:"MajorVersion"`

	// Minor version.
	// Note: This field may return null, indicating that no valid values can be obtained.
	MinorVersion *string `json:"MinorVersion,omitnil,omitempty" name:"MinorVersion"`

	// Start time.
	// Note: This field may return null, indicating that no valid values can be obtained.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`
}

type ArchiveLogModel struct {
	// Archivelog ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	ArchiveLogId *int64 `json:"ArchiveLogId,omitnil,omitempty" name:"ArchiveLogId"`

	// Backup duration
	// Note: This field may return null, indicating that no valid values can be obtained.
	BackupDuration *int64 `json:"BackupDuration,omitnil,omitempty" name:"BackupDuration"`

	// Backup status
	// Note: This field may return null, indicating that no valid values can be obtained.
	BackupStatus *string `json:"BackupStatus,omitnil,omitempty" name:"BackupStatus"`

	// Backup end time.
	// Note: This field may return null, indicating that no valid values can be obtained.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Error message
	// Note: This field may return null, indicating that no valid values can be obtained.
	ErrorMessage *string `json:"ErrorMessage,omitnil,omitempty" name:"ErrorMessage"`

	// Expiration time
	// Note: This field may return null, indicating that no valid values can be obtained.
	ExpiredTime *string `json:"ExpiredTime,omitnil,omitempty" name:"ExpiredTime"`

	// Backup file name
	// Note: This field may return null, indicating that no valid values can be obtained.
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// Backup set file size in Byte
	// Note: This field may return null, indicating that no valid values can be obtained.
	FileSize *int64 `json:"FileSize,omitnil,omitempty" name:"FileSize"`

	// Instance ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Backup start time.
	// Note: This field may return null, indicating that no valid values can be obtained.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`
}

type AutoScalingConfig struct {
	// <p>Minimum value of ccu</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	RangeMin *float64 `json:"RangeMin,omitnil,omitempty" name:"RangeMin"`

	// <p>Maximum value of ccu</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	RangeMax *float64 `json:"RangeMax,omitnil,omitempty" name:"RangeMax"`
}

type BackupMethodStatisticsModel struct {
	// <p>Auto-backup size in Byte</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	AutoBackupSize *int64 `json:"AutoBackupSize,omitnil,omitempty" name:"AutoBackupSize"`

	// <p>Average size of total backup per day in Byte</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	AverageSizePerDay *int64 `json:"AverageSizePerDay,omitnil,omitempty" name:"AverageSizePerDay"`

	// <p>Manual backup size, unit Byte</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	ManualBackupSize *int64 `json:"ManualBackupSize,omitnil,omitempty" name:"ManualBackupSize"`

	// <p>Total backup size in Byte</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	TotalSize *int64 `json:"TotalSize,omitnil,omitempty" name:"TotalSize"`
}

type BackupMethodStatisticsOutPut struct {
	// <p>Auto-backup size in Byte</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	AutoBackupSize []*int64 `json:"AutoBackupSize,omitnil,omitempty" name:"AutoBackupSize"`

	// <p>Manual backup size, unit Byte</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	ManualBackupSize []*int64 `json:"ManualBackupSize,omitnil,omitempty" name:"ManualBackupSize"`
}

type BackupPolicyModelInput struct {
	// <P>Backup end time.</p>
	BackupEndTime *string `json:"BackupEndTime,omitnil,omitempty" name:"BackupEndTime"`

	// <P>Backup method: physical physical backup, snapshot snapshot backup</p>
	BackupMethod *string `json:"BackupMethod,omitnil,omitempty" name:"BackupMethod"`

	// <P>Backup start time</p>
	BackupStartTime *string `json:"BackupStartTime,omitnil,omitempty" name:"BackupStartTime"`

	// <P>Whether full backup is enabled</p>
	EnableFull *int64 `json:"EnableFull,omitnil,omitempty" name:"EnableFull"`

	// <P>Whether to enable log backup</p>
	EnableLog *int64 `json:"EnableLog,omitnil,omitempty" name:"EnableLog"`

	// <P>Full backup retention time can currently only be set to 7 days.</p>
	FullRetentionPeriod *int64 `json:"FullRetentionPeriod,omitnil,omitempty" name:"FullRetentionPeriod"`

	// <p>Instance ID.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <P>Log retention days. currently, can only set retention to 7 days.</p>
	LogRetentionPeriod *int64 `json:"LogRetentionPeriod,omitnil,omitempty" name:"LogRetentionPeriod"`

	// <P>Days of the week to perform backup.</p>
	PeriodTime *string `json:"PeriodTime,omitnil,omitempty" name:"PeriodTime"`

	// <p>Storage type: COS, SNAPSHOT</p>valid values: <ul><li> COS: COS storage</li><li> SNAPSHOT: cloud disk SNAPSHOT</li></ul>
	StorageType *string `json:"StorageType,omitnil,omitempty" name:"StorageType"`
}

type BackupPolicyModelOutPut struct {
	// <p>Backup end time</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	BackupEndTime *string `json:"BackupEndTime,omitnil,omitempty" name:"BackupEndTime"`

	// <p>Backup method</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	BackupMethod *string `json:"BackupMethod,omitnil,omitempty" name:"BackupMethod"`

	// <p>Backup start time</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	BackupStartTime *string `json:"BackupStartTime,omitnil,omitempty" name:"BackupStartTime"`

	// <p>Engine type</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	DBType *string `json:"DBType,omitnil,omitempty" name:"DBType"`

	// <p>Engine version</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	DBVersion *string `json:"DBVersion,omitnil,omitempty" name:"DBVersion"`

	// <p>Whether full backup is enabled</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	EnableFull *int64 `json:"EnableFull,omitnil,omitempty" name:"EnableFull"`

	// <p>Whether to enable log backup</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	EnableLog *int64 `json:"EnableLog,omitnil,omitempty" name:"EnableLog"`

	// <p>Expected next backup time</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	ExpectedNextBackupPeriod *string `json:"ExpectedNextBackupPeriod,omitnil,omitempty" name:"ExpectedNextBackupPeriod"`

	// <p>Full backup retention time</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	FullRetentionPeriod *int64 `json:"FullRetentionPeriod,omitnil,omitempty" name:"FullRetentionPeriod"`

	// <p>Policy ID</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// <p>Instance ID.</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Log retention days</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	LogRetentionPeriod *int64 `json:"LogRetentionPeriod,omitnil,omitempty" name:"LogRetentionPeriod"`

	// <p>Days of the week to perform backup</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	PeriodTime *string `json:"PeriodTime,omitnil,omitempty" name:"PeriodTime"`

	// <p>Region</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// <p>Backup cycle type 0:Weekly</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	PeriodType *int64 `json:"PeriodType,omitnil,omitempty" name:"PeriodType"`
}

type BackupSetModel struct {
	// Backup set duration, unit sec
	// Note: This field may return null, indicating that no valid values can be obtained.
	BackupDuration *int64 `json:"BackupDuration,omitnil,omitempty" name:"BackupDuration"`

	// Backup method
	// Note: This field may return null, indicating that no valid values can be obtained.
	BackupMethod *string `json:"BackupMethod,omitnil,omitempty" name:"BackupMethod"`

	// Backup note name
	// Note: This field may return null, indicating that no valid values can be obtained.
	BackupName *string `json:"BackupName,omitnil,omitempty" name:"BackupName"`

	// Backup set progress percentage
	// Note: This field may return null, indicating that no valid values can be obtained.
	BackupProgress *float64 `json:"BackupProgress,omitnil,omitempty" name:"BackupProgress"`

	// Backup set ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	BackupSetId *int64 `json:"BackupSetId,omitnil,omitempty" name:"BackupSetId"`

	// Backup status
	// Note: This field may return null, indicating that no valid values can be obtained.
	BackupStatus *string `json:"BackupStatus,omitnil,omitempty" name:"BackupStatus"`

	// Backup set type
	// Note: This field may return null, indicating that no valid values can be obtained.
	BackupType *string `json:"BackupType,omitnil,omitempty" name:"BackupType"`

	// Instance version number
	// Note: This field may return null, indicating that no valid values can be obtained.
	DBVersion *string `json:"DBVersion,omitnil,omitempty" name:"DBVersion"`

	// Backup end time.
	// Note: This field may return null, indicating that no valid values can be obtained.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Transaction commit end time
	// Note: This field may return null, indicating that no valid values can be obtained.
	EndTrxTime *string `json:"EndTrxTime,omitnil,omitempty" name:"EndTrxTime"`

	// Error message
	// Note: This field may return null, indicating that no valid values can be obtained.
	ErrorMessage *string `json:"ErrorMessage,omitnil,omitempty" name:"ErrorMessage"`

	// Backup expiration time
	// Note: This field may return null, indicating that no valid values can be obtained.
	ExpiredTime *string `json:"ExpiredTime,omitnil,omitempty" name:"ExpiredTime"`

	// Backup file name
	// Note: This field may return null, indicating that no valid values can be obtained.
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// Backup set file size in Byte
	// Note: This field may return null, indicating that no valid values can be obtained.
	FileSize *int64 `json:"FileSize,omitnil,omitempty" name:"FileSize"`

	// Instance ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Backup trigger. 0: autobackup; 1: manual.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ManualBackup *int64 `json:"ManualBackup,omitnil,omitempty" name:"ManualBackup"`

	// Backup start time.
	// Note: This field may return null, indicating that no valid values can be obtained.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`
}

type BackupSetsReqFilter struct {
	// <p>Backup method, multiple values separated by commas</p><p>Enumeration value:</p><ul><li>physical: Physical backup</li><li>snapshot: Snapshot backup</li></ul>
	// Note: This field may return null, indicating that no valid values can be obtained.
	BackupMethod *string `json:"BackupMethod,omitnil,omitempty" name:"BackupMethod"`

	// <p>Backup status, multiple values separated by commas. Description of meaning: pending scheduling pending, running running, success success, failed failed</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	BackupStatus *string `json:"BackupStatus,omitnil,omitempty" name:"BackupStatus"`

	// <p>Backup type, multiple values separated by commas. Description of meaning: full backup full</p><p>Enumeration value:</p><ul><li>full: Full backup</li></ul>
	// Note: This field may return null, indicating that no valid values can be obtained.
	BackupType *string `json:"BackupType,omitnil,omitempty" name:"BackupType"`

	// <p>Backup trigger mode</p><p>Enumeration value:</p><ul><li>0: Auto-backup</li><li>1: Manual backup</li></ul>
	// Note: This field may return null, indicating that no valid values can be obtained.
	ManualBackup *string `json:"ManualBackup,omitnil,omitempty" name:"ManualBackup"`
}

type BackupStatisticsModel struct {
	// <p>Average size of total backup per day in Byte</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	AverageSizePerDay *int64 `json:"AverageSizePerDay,omitnil,omitempty" name:"AverageSizePerDay"`

	// <p>Backup size of data, in Byte</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	DataBackupSize *int64 `json:"DataBackupSize,omitnil,omitempty" name:"DataBackupSize"`

	// <p>Log backup size in Byte</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	LogBackupSize *int64 `json:"LogBackupSize,omitnil,omitempty" name:"LogBackupSize"`

	// <p>Total backup count</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// <p>Total backup size in Byte</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	TotalSize *int64 `json:"TotalSize,omitnil,omitempty" name:"TotalSize"`
}

type BackupTypeStatisticsModel struct {
	// <p>Backup size of data, in Byte</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	DataBackupSize []*int64 `json:"DataBackupSize,omitnil,omitempty" name:"DataBackupSize"`

	// <p>Log backup size in Byte</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	LogBackupSize []*int64 `json:"LogBackupSize,omitnil,omitempty" name:"LogBackupSize"`
}

type BinlogInfo struct {
	// Unique ID of the log service
	Sid *string `json:"Sid,omitnil,omitempty" name:"Sid"`

	// Log service type
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Unique ID of the instance
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

// Predefined struct for user
type CancelIsolateDBInstancesRequestParams struct {
	// List of isolated instance ids required.
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

type CancelIsolateDBInstancesRequest struct {
	*tchttp.BaseRequest
	
	// List of isolated instance ids required.
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

func (r *CancelIsolateDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelIsolateDBInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CancelIsolateDBInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelIsolateDBInstancesResponseParams struct {
	// List of successfully isolated instance ids.
	SuccessInstanceIds []*string `json:"SuccessInstanceIds,omitnil,omitempty" name:"SuccessInstanceIds"`

	// Isolation removal failed instance Id list.
	FailedInstanceIds []*string `json:"FailedInstanceIds,omitnil,omitempty" name:"FailedInstanceIds"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CancelIsolateDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *CancelIsolateDBInstancesResponseParams `json:"Response"`
}

func (r *CancelIsolateDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelIsolateDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CloneInstanceModel struct {
	// Clone task end time.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CloneEndTime *string `json:"CloneEndTime,omitnil,omitempty" name:"CloneEndTime"`

	// Clone record ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CloneId *int64 `json:"CloneId,omitnil,omitempty" name:"CloneId"`

	// Clone instance type.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CloneInsType *string `json:"CloneInsType,omitnil,omitempty" name:"CloneInsType"`

	// Clone instance ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CloneInstanceId *string `json:"CloneInstanceId,omitnil,omitempty" name:"CloneInstanceId"`

	// Whether the clone instance is deleted.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CloneInstanceIsDeleted *bool `json:"CloneInstanceIsDeleted,omitnil,omitempty" name:"CloneInstanceIsDeleted"`

	// Task progress of clone.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CloneProgress *float64 `json:"CloneProgress,omitnil,omitempty" name:"CloneProgress"`

	// Task start time of the clone.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CloneStartTime *string `json:"CloneStartTime,omitnil,omitempty" name:"CloneStartTime"`

	// Task status.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CloneStatus *string `json:"CloneStatus,omitnil,omitempty" name:"CloneStatus"`

	// Clone time.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CloneTime *string `json:"CloneTime,omitnil,omitempty" name:"CloneTime"`

	// Clone type.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CloneType *string `json:"CloneType,omitnil,omitempty" name:"CloneType"`

	// Creation time.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Engine type.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DBType *string `json:"DBType,omitnil,omitempty" name:"DBType"`

	// Region.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Source instance ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SourceInstanceId *string `json:"SourceInstanceId,omitnil,omitempty" name:"SourceInstanceId"`
}

type ConstraintRange struct {
	// Minimum value when the constraint type is section.
	Min *string `json:"Min,omitnil,omitempty" name:"Min"`

	// Specifies the maximum value when the constraint type is section.
	Max *string `json:"Max,omitnil,omitempty" name:"Max"`
}

// Predefined struct for user
type CreateCloneInstanceRequestParams struct {
	// <p>Creating an Instance Region</p>
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// <p>Character type vpcid</p>
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// <p>Character type subnetid</p>
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// <p>Purchase specification</p>
	SpecCode *string `json:"SpecCode,omitnil,omitempty" name:"SpecCode"`

	// <p>Node disk capacity (unit: GB)</p>
	Disk *int64 `json:"Disk,omitnil,omitempty" name:"Disk"`

	// <p>Number of storage nodes</p>
	StorageNodeNum *int64 `json:"StorageNodeNum,omitnil,omitempty" name:"StorageNodeNum"`

	// <p>Source instance id</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Instance name. The required length is 1-60. It can contain Chinese characters, English case, digits, hyphens (-), and underscores (_).</p>
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// <p>Tag key-value pair array</p>
	ResourceTags []*ResourceTag `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`

	// <p>Backup and rollback name</p>
	BackupName *string `json:"BackupName,omitnil,omitempty" name:"BackupName"`

	// <p>CPU cores of the storage node</p>
	StorageNodeCpu *int64 `json:"StorageNodeCpu,omitnil,omitempty" name:"StorageNodeCpu"`

	// <p>Storage node memory size</p>
	StorageNodeMem *int64 `json:"StorageNodeMem,omitnil,omitempty" name:"StorageNodeMem"`

	// <p>Create version</p>
	CreateVersion *string `json:"CreateVersion,omitnil,omitempty" name:"CreateVersion"`

	// <p>Create port number</p>
	Vport *int64 `json:"Vport,omitnil,omitempty" name:"Vport"`

	// <p>Recovery time point</p>
	RecoverTime *string `json:"RecoverTime,omitnil,omitempty" name:"RecoverTime"`

	// <p>Instance Architecture Type, separate: decoupled architecture; hybrid: peer-to-peer architecture</p>
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// <p>Disk Type, CLOUD_HSSD enhanced SSD, CLOUD_TCS local SSD disk</p>
	StorageType *string `json:"StorageType,omitnil,omitempty" name:"StorageType"`

	// <p>Multi-AZ list, Zones[0] refers to the primary AZ</p>
	Zones []*string `json:"Zones,omitnil,omitempty" name:"Zones"`

	// <p>Number of replicas</p>
	FullReplications *int64 `json:"FullReplications,omitnil,omitempty" name:"FullReplications"`

	// <p>Instance mode, normal: standard type; enhanced: enhanced</p>
	InstanceMode *string `json:"InstanceMode,omitnil,omitempty" name:"InstanceMode"`

	// <p>Security group id list</p>
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`
}

type CreateCloneInstanceRequest struct {
	*tchttp.BaseRequest
	
	// <p>Creating an Instance Region</p>
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// <p>Character type vpcid</p>
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// <p>Character type subnetid</p>
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// <p>Purchase specification</p>
	SpecCode *string `json:"SpecCode,omitnil,omitempty" name:"SpecCode"`

	// <p>Node disk capacity (unit: GB)</p>
	Disk *int64 `json:"Disk,omitnil,omitempty" name:"Disk"`

	// <p>Number of storage nodes</p>
	StorageNodeNum *int64 `json:"StorageNodeNum,omitnil,omitempty" name:"StorageNodeNum"`

	// <p>Source instance id</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Instance name. The required length is 1-60. It can contain Chinese characters, English case, digits, hyphens (-), and underscores (_).</p>
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// <p>Tag key-value pair array</p>
	ResourceTags []*ResourceTag `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`

	// <p>Backup and rollback name</p>
	BackupName *string `json:"BackupName,omitnil,omitempty" name:"BackupName"`

	// <p>CPU cores of the storage node</p>
	StorageNodeCpu *int64 `json:"StorageNodeCpu,omitnil,omitempty" name:"StorageNodeCpu"`

	// <p>Storage node memory size</p>
	StorageNodeMem *int64 `json:"StorageNodeMem,omitnil,omitempty" name:"StorageNodeMem"`

	// <p>Create version</p>
	CreateVersion *string `json:"CreateVersion,omitnil,omitempty" name:"CreateVersion"`

	// <p>Create port number</p>
	Vport *int64 `json:"Vport,omitnil,omitempty" name:"Vport"`

	// <p>Recovery time point</p>
	RecoverTime *string `json:"RecoverTime,omitnil,omitempty" name:"RecoverTime"`

	// <p>Instance Architecture Type, separate: decoupled architecture; hybrid: peer-to-peer architecture</p>
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// <p>Disk Type, CLOUD_HSSD enhanced SSD, CLOUD_TCS local SSD disk</p>
	StorageType *string `json:"StorageType,omitnil,omitempty" name:"StorageType"`

	// <p>Multi-AZ list, Zones[0] refers to the primary AZ</p>
	Zones []*string `json:"Zones,omitnil,omitempty" name:"Zones"`

	// <p>Number of replicas</p>
	FullReplications *int64 `json:"FullReplications,omitnil,omitempty" name:"FullReplications"`

	// <p>Instance mode, normal: standard type; enhanced: enhanced</p>
	InstanceMode *string `json:"InstanceMode,omitnil,omitempty" name:"InstanceMode"`

	// <p>Security group id list</p>
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`
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
	delete(f, "Zone")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "SpecCode")
	delete(f, "Disk")
	delete(f, "StorageNodeNum")
	delete(f, "InstanceId")
	delete(f, "InstanceName")
	delete(f, "ResourceTags")
	delete(f, "BackupName")
	delete(f, "StorageNodeCpu")
	delete(f, "StorageNodeMem")
	delete(f, "CreateVersion")
	delete(f, "Vport")
	delete(f, "RecoverTime")
	delete(f, "InstanceType")
	delete(f, "StorageType")
	delete(f, "Zones")
	delete(f, "FullReplications")
	delete(f, "InstanceMode")
	delete(f, "SecurityGroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCloneInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCloneInstanceResponseParams struct {
	// <p>Clone instance ID</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Task ID.</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCloneInstanceResponse struct {
	*tchttp.BaseResponse
	Response *CreateCloneInstanceResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateDBInstancesRequestParams struct {
	// <p>Creating an Instance Region</p>
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// <p>Character type vpcid</p>
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// <p>Character type subnetid</p>
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// <p>Purchase specification</p>
	SpecCode *string `json:"SpecCode,omitnil,omitempty" name:"SpecCode"`

	// <p>Node disk capacity (unit: GB)</p>
	Disk *int64 `json:"Disk,omitnil,omitempty" name:"Disk"`

	// <p>Number of storage nodes</p>
	StorageNodeNum *int64 `json:"StorageNodeNum,omitnil,omitempty" name:"StorageNodeNum"`

	// <p>Number of node replicas for storage, up to 5, must be an odd number</p>
	Replications *int64 `json:"Replications,omitnil,omitempty" name:"Replications"`

	// <p>Instance count. Maximum is 10.</p>
	InstanceCount *int64 `json:"InstanceCount,omitnil,omitempty" name:"InstanceCount"`

	// <p>Number of replicas</p>
	FullReplications *int64 `json:"FullReplications,omitnil,omitempty" name:"FullReplications"`

	// <p>Create an instance version, using the current latest version by default</p>
	CreateVersion *string `json:"CreateVersion,omitnil,omitempty" name:"CreateVersion"`

	// <p>Instance name. The required length is 1-60. It can contain Chinese characters, English case, digits, hyphens (-), and underscores (_).</p>
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// <p>Tag key-value pair array</p>
	ResourceTags []*ResourceTag `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`

	// <p>Initialize instance parameters. For example:<br>character_set_server (character set, defaults to utf8),<br>lower_case_table_names (table name case sensitivity, 0 - sensitive; 1 - insensitive, default is 0)</p>
	InitParams []*InstanceParam `json:"InitParams,omitnil,omitempty" name:"InitParams"`

	// <p>Time unit, m: month</p>
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// <p>Commodity duration size</p>
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// <p>CPU cores of the storage node</p>
	StorageNodeCpu *int64 `json:"StorageNodeCpu,omitnil,omitempty" name:"StorageNodeCpu"`

	// <p>Storage node memory size</p>
	StorageNodeMem *int64 `json:"StorageNodeMem,omitnil,omitempty" name:"StorageNodeMem"`

	// <p>Payment mode. 0 means pay-as-you-go/postpaid, 1 means prepaid.</p>
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// <p>Number of control nodes</p>
	MCNum *int64 `json:"MCNum,omitnil,omitempty" name:"MCNum"`

	// <p>Custom port</p>
	Vport *int64 `json:"Vport,omitnil,omitempty" name:"Vport"`

	// <p>Multi-AZ availability zone list</p>
	Zones []*string `json:"Zones,omitnil,omitempty" name:"Zones"`

	// <p>Whether to use a coupon.</p>
	AutoVoucher *bool `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// <p>Coupon list</p>
	VoucherIds []*string `json:"VoucherIds,omitnil,omitempty" name:"VoucherIds"`

	// <p>Instance Architecture Type, separate: decoupled architecture; hybrid: peer-to-peer architecture</p>
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// <p>Disk Type, CLOUD_HSSD enhanced SSD, CLOUD_TCS local SSD disk</p>
	StorageType *string `json:"StorageType,omitnil,omitempty" name:"StorageType"`

	// <p>AZ mode. 1: Single AZ, 2: Multi-AZ non-primary AZ, 3: Multi-AZ primary AZ</p>
	AZMode *int64 `json:"AZMode,omitnil,omitempty" name:"AZMode"`

	// <p>Instance mode</p>
	InstanceMode *string `json:"InstanceMode,omitnil,omitempty" name:"InstanceMode"`

	// <p>Parameter template id</p>
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// <p>Compatible mode, enum:MySQL,HBase</p>
	SQLMode *string `json:"SQLMode,omitnil,omitempty" name:"SQLMode"`

	// <p>ccu configuration of the svls instance</p>
	AutoScaleConfig *AutoScalingConfig `json:"AutoScaleConfig,omitnil,omitempty" name:"AutoScaleConfig"`

	// <p>Bind to security group list</p>
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// <p>root userName. The default is dbaadmin in the current version. It will reset to dbaadmin even if a value is passed.</p>
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// <p>dbaadmin password</p>
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// <p>Whether transparent data encryption is enabled. 0: not enabled; 1: enabled</p>
	EncryptionEnable *int64 `json:"EncryptionEnable,omitnil,omitempty" name:"EncryptionEnable"`
}

type CreateDBInstancesRequest struct {
	*tchttp.BaseRequest
	
	// <p>Creating an Instance Region</p>
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// <p>Character type vpcid</p>
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// <p>Character type subnetid</p>
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// <p>Purchase specification</p>
	SpecCode *string `json:"SpecCode,omitnil,omitempty" name:"SpecCode"`

	// <p>Node disk capacity (unit: GB)</p>
	Disk *int64 `json:"Disk,omitnil,omitempty" name:"Disk"`

	// <p>Number of storage nodes</p>
	StorageNodeNum *int64 `json:"StorageNodeNum,omitnil,omitempty" name:"StorageNodeNum"`

	// <p>Number of node replicas for storage, up to 5, must be an odd number</p>
	Replications *int64 `json:"Replications,omitnil,omitempty" name:"Replications"`

	// <p>Instance count. Maximum is 10.</p>
	InstanceCount *int64 `json:"InstanceCount,omitnil,omitempty" name:"InstanceCount"`

	// <p>Number of replicas</p>
	FullReplications *int64 `json:"FullReplications,omitnil,omitempty" name:"FullReplications"`

	// <p>Create an instance version, using the current latest version by default</p>
	CreateVersion *string `json:"CreateVersion,omitnil,omitempty" name:"CreateVersion"`

	// <p>Instance name. The required length is 1-60. It can contain Chinese characters, English case, digits, hyphens (-), and underscores (_).</p>
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// <p>Tag key-value pair array</p>
	ResourceTags []*ResourceTag `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`

	// <p>Initialize instance parameters. For example:<br>character_set_server (character set, defaults to utf8),<br>lower_case_table_names (table name case sensitivity, 0 - sensitive; 1 - insensitive, default is 0)</p>
	InitParams []*InstanceParam `json:"InitParams,omitnil,omitempty" name:"InitParams"`

	// <p>Time unit, m: month</p>
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// <p>Commodity duration size</p>
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// <p>CPU cores of the storage node</p>
	StorageNodeCpu *int64 `json:"StorageNodeCpu,omitnil,omitempty" name:"StorageNodeCpu"`

	// <p>Storage node memory size</p>
	StorageNodeMem *int64 `json:"StorageNodeMem,omitnil,omitempty" name:"StorageNodeMem"`

	// <p>Payment mode. 0 means pay-as-you-go/postpaid, 1 means prepaid.</p>
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// <p>Number of control nodes</p>
	MCNum *int64 `json:"MCNum,omitnil,omitempty" name:"MCNum"`

	// <p>Custom port</p>
	Vport *int64 `json:"Vport,omitnil,omitempty" name:"Vport"`

	// <p>Multi-AZ availability zone list</p>
	Zones []*string `json:"Zones,omitnil,omitempty" name:"Zones"`

	// <p>Whether to use a coupon.</p>
	AutoVoucher *bool `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// <p>Coupon list</p>
	VoucherIds []*string `json:"VoucherIds,omitnil,omitempty" name:"VoucherIds"`

	// <p>Instance Architecture Type, separate: decoupled architecture; hybrid: peer-to-peer architecture</p>
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// <p>Disk Type, CLOUD_HSSD enhanced SSD, CLOUD_TCS local SSD disk</p>
	StorageType *string `json:"StorageType,omitnil,omitempty" name:"StorageType"`

	// <p>AZ mode. 1: Single AZ, 2: Multi-AZ non-primary AZ, 3: Multi-AZ primary AZ</p>
	AZMode *int64 `json:"AZMode,omitnil,omitempty" name:"AZMode"`

	// <p>Instance mode</p>
	InstanceMode *string `json:"InstanceMode,omitnil,omitempty" name:"InstanceMode"`

	// <p>Parameter template id</p>
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// <p>Compatible mode, enum:MySQL,HBase</p>
	SQLMode *string `json:"SQLMode,omitnil,omitempty" name:"SQLMode"`

	// <p>ccu configuration of the svls instance</p>
	AutoScaleConfig *AutoScalingConfig `json:"AutoScaleConfig,omitnil,omitempty" name:"AutoScaleConfig"`

	// <p>Bind to security group list</p>
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// <p>root userName. The default is dbaadmin in the current version. It will reset to dbaadmin even if a value is passed.</p>
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// <p>dbaadmin password</p>
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// <p>Whether transparent data encryption is enabled. 0: not enabled; 1: enabled</p>
	EncryptionEnable *int64 `json:"EncryptionEnable,omitnil,omitempty" name:"EncryptionEnable"`
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
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "SpecCode")
	delete(f, "Disk")
	delete(f, "StorageNodeNum")
	delete(f, "Replications")
	delete(f, "InstanceCount")
	delete(f, "FullReplications")
	delete(f, "CreateVersion")
	delete(f, "InstanceName")
	delete(f, "ResourceTags")
	delete(f, "InitParams")
	delete(f, "TimeUnit")
	delete(f, "TimeSpan")
	delete(f, "StorageNodeCpu")
	delete(f, "StorageNodeMem")
	delete(f, "PayMode")
	delete(f, "MCNum")
	delete(f, "Vport")
	delete(f, "Zones")
	delete(f, "AutoVoucher")
	delete(f, "VoucherIds")
	delete(f, "InstanceType")
	delete(f, "StorageType")
	delete(f, "AZMode")
	delete(f, "InstanceMode")
	delete(f, "TemplateId")
	delete(f, "SQLMode")
	delete(f, "AutoScaleConfig")
	delete(f, "SecurityGroupIds")
	delete(f, "UserName")
	delete(f, "Password")
	delete(f, "EncryptionEnable")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDBInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDBInstancesResponseParams struct {
	// <p>Instance ID.</p>
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// <p>Task ID.</p>
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

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
type CreateDBSBackupRequestParams struct {
	// <p>Backup method: physical, snapshot. This value should be consistent with the backupMethod in the API response of DescribeDBSBackupPolicy.</p><p>Enumeration value:</p><ul><li>physical: Physical backup</li><li>snapshot: Snapshot backup</li></ul>
	BackupMethod *string `json:"BackupMethod,omitnil,omitempty" name:"BackupMethod"`

	// <P>Backup type: currently, only supports full.</p>
	BackupType *string `json:"BackupType,omitnil,omitempty" name:"BackupType"`

	// <p>Instance ID.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <P>Backup notes.</p>
	BackupName *string `json:"BackupName,omitnil,omitempty" name:"BackupName"`
}

type CreateDBSBackupRequest struct {
	*tchttp.BaseRequest
	
	// <p>Backup method: physical, snapshot. This value should be consistent with the backupMethod in the API response of DescribeDBSBackupPolicy.</p><p>Enumeration value:</p><ul><li>physical: Physical backup</li><li>snapshot: Snapshot backup</li></ul>
	BackupMethod *string `json:"BackupMethod,omitnil,omitempty" name:"BackupMethod"`

	// <P>Backup type: currently, only supports full.</p>
	BackupType *string `json:"BackupType,omitnil,omitempty" name:"BackupType"`

	// <p>Instance ID.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <P>Backup notes.</p>
	BackupName *string `json:"BackupName,omitnil,omitempty" name:"BackupName"`
}

func (r *CreateDBSBackupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDBSBackupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BackupMethod")
	delete(f, "BackupType")
	delete(f, "InstanceId")
	delete(f, "BackupName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDBSBackupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDBSBackupResponseParams struct {
	// <p>Backup set ID.</p>.
	BackupSetId *int64 `json:"BackupSetId,omitnil,omitempty" name:"BackupSetId"`

	// <P>Whether it is successful</p>.
	IsSuccess *bool `json:"IsSuccess,omitnil,omitempty" name:"IsSuccess"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDBSBackupResponse struct {
	*tchttp.BaseResponse
	Response *CreateDBSBackupResponseParams `json:"Response"`
}

func (r *CreateDBSBackupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDBSBackupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUsersRequestParams struct {
	// <p>Instance id</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Add user list</p>
	Users []*User `json:"Users,omitnil,omitempty" name:"Users"`

	// <p>Unencrypted password</p>
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// <p>Encryption password</p>
	EncryptedPassword *string `json:"EncryptedPassword,omitnil,omitempty" name:"EncryptedPassword"`

	// <p>User description</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type CreateUsersRequest struct {
	*tchttp.BaseRequest
	
	// <p>Instance id</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Add user list</p>
	Users []*User `json:"Users,omitnil,omitempty" name:"Users"`

	// <p>Unencrypted password</p>
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// <p>Encryption password</p>
	EncryptedPassword *string `json:"EncryptedPassword,omitnil,omitempty" name:"EncryptedPassword"`

	// <p>User description</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *CreateUsersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUsersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Users")
	delete(f, "Password")
	delete(f, "EncryptedPassword")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateUsersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUsersResponseParams struct {
	// <p>Task ID.</p>
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateUsersResponse struct {
	*tchttp.BaseResponse
	Response *CreateUsersResponseParams `json:"Response"`
}

func (r *CreateUsersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUsersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DBParamValue struct {
	// Parameter name.
	Param *string `json:"Param,omitnil,omitempty" name:"Param"`

	// Parameter value.
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type DataBackupStatisticsModel struct {
	// Auto-backup count
	// Note: This field may return null, indicating that no valid values can be obtained.
	AutoBackupCount *int64 `json:"AutoBackupCount,omitnil,omitempty" name:"AutoBackupCount"`

	// Auto-backup size, in Byte
	// Note: This field may return null, indicating that no valid values can be obtained.
	AutoBackupSize *int64 `json:"AutoBackupSize,omitnil,omitempty" name:"AutoBackupSize"`

	// Avg backup size per each, in Byte
	// Note: This field may return null, indicating that no valid values can be obtained.
	AverageSizePerBackup *int64 `json:"AverageSizePerBackup,omitnil,omitempty" name:"AverageSizePerBackup"`

	// Avg daily backup size, Byte
	// Note: This field may return null, indicating that no valid values can be obtained.
	AverageSizePerDay *int64 `json:"AverageSizePerDay,omitnil,omitempty" name:"AverageSizePerDay"`

	// Manual backup count
	// Note: This field may return null, indicating that no valid values can be obtained.
	ManualBackupCount *int64 `json:"ManualBackupCount,omitnil,omitempty" name:"ManualBackupCount"`

	// Manual backup size, in Byte
	// Note: This field may return null, indicating that no valid values can be obtained.
	ManualBackupSize *int64 `json:"ManualBackupSize,omitnil,omitempty" name:"ManualBackupSize"`

	// Number of data backups
	// Note: This field may return null, indicating that no valid values can be obtained.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Data backup size, in Byte
	// Note: This field may return null, indicating that no valid values can be obtained.
	TotalSize *int64 `json:"TotalSize,omitnil,omitempty" name:"TotalSize"`
}

type Database struct {
	// Database name.
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`
}

type DatabaseFunction struct {
	// Function name.
	Func *string `json:"Func,omitnil,omitempty" name:"Func"`
}

type DatabasePrivileges struct {
	// database name
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// Permission list
	Privileges []*string `json:"Privileges,omitnil,omitempty" name:"Privileges"`
}

type DatabaseProcedure struct {
	// Stored procedure name.
	Proc *string `json:"Proc,omitnil,omitempty" name:"Proc"`
}

type DatabaseTable struct {
	// Table name
	Table *string `json:"Table,omitnil,omitempty" name:"Table"`
}

type DatabaseView struct {
	// View name.
	View *string `json:"View,omitnil,omitempty" name:"View"`
}

// Predefined struct for user
type DeleteDBSBackupSetsRequestParams struct {
	// <p>Instance ID.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Backup set list. the value comes from the DescribeDBSBackupSets api response.</p>
	BackupSetIdList []*int64 `json:"BackupSetIdList,omitnil,omitempty" name:"BackupSetIdList"`
}

type DeleteDBSBackupSetsRequest struct {
	*tchttp.BaseRequest
	
	// <p>Instance ID.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Backup set list. the value comes from the DescribeDBSBackupSets api response.</p>
	BackupSetIdList []*int64 `json:"BackupSetIdList,omitnil,omitempty" name:"BackupSetIdList"`
}

func (r *DeleteDBSBackupSetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDBSBackupSetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BackupSetIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDBSBackupSetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDBSBackupSetsResponseParams struct {
	// <P>Number of backups deleted this time.</p>
	Deleted *int64 `json:"Deleted,omitnil,omitempty" name:"Deleted"`

	// <P>Whether all are deleted successfully.</p>
	IsSuccess *bool `json:"IsSuccess,omitnil,omitempty" name:"IsSuccess"`

	// <P>Total number of backups to be deleted.</p>
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteDBSBackupSetsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDBSBackupSetsResponseParams `json:"Response"`
}

func (r *DeleteDBSBackupSetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDBSBackupSetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteUsersRequestParams struct {
	// <p>Instance id</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Batch delete user list</p>
	Users []*User `json:"Users,omitnil,omitempty" name:"Users"`
}

type DeleteUsersRequest struct {
	*tchttp.BaseRequest
	
	// <p>Instance id</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Batch delete user list</p>
	Users []*User `json:"Users,omitnil,omitempty" name:"Users"`
}

func (r *DeleteUsersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUsersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Users")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteUsersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteUsersResponseParams struct {
	// <p>Task ID.</p>
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteUsersResponse struct {
	*tchttp.BaseResponse
	Response *DeleteUsersResponseParams `json:"Response"`
}

func (r *DeleteUsersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUsersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstanceDetailRequestParams struct {
	// <p>Instance ID.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeDBInstanceDetailRequest struct {
	*tchttp.BaseRequest
	
	// <p>Instance ID.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeDBInstanceDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstanceDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBInstanceDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstanceDetailResponseParams struct {
	// <p>Instance name</p>
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// <p>Region</p>
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// <p>Character type vpcid</p>
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// <p>Character type subnetid</p>
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// <p>Create an instance version</p>
	CreateVersion *string `json:"CreateVersion,omitnil,omitempty" name:"CreateVersion"`

	// <p>Subnet IP</p>
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// <p>Subnet port</p>
	Vport *int64 `json:"Vport,omitnil,omitempty" name:"Vport"`

	// <p>Instance status</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>Node disk capacity (unit: GB)</p>
	Disk *int64 `json:"Disk,omitnil,omitempty" name:"Disk"`

	// <p>Number of storage nodes</p>
	StorageNodeNum *int64 `json:"StorageNodeNum,omitnil,omitempty" name:"StorageNodeNum"`

	// <p>Initialize instance parameters</p>
	InitParams []*InstanceParam `json:"InitParams,omitnil,omitempty" name:"InitParams"`

	// <p>Instance tag information.</p>
	ResourceTags []*ResourceTag `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`

	// <p>Creation time.</p>
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// <p>Update time.</p>
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// <p>Number of storage node replicas</p>
	Replications *int64 `json:"Replications,omitnil,omitempty" name:"Replications"`

	// <p>Number of replicas</p>
	FullReplications *int64 `json:"FullReplications,omitnil,omitempty" name:"FullReplications"`

	// <p>Character set</p>
	CharSet *string `json:"CharSet,omitnil,omitempty" name:"CharSet"`

	// <p>Node information</p>
	Node []*NodeInfo `json:"Node,omitnil,omitempty" name:"Node"`

	// <p>Region of the instance</p>
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// <p>Instance specification</p>
	SpecCode *string `json:"SpecCode,omitnil,omitempty" name:"SpecCode"`

	// <p>Instance ID.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Status description in Chinese of the instance</p>
	StatusDesc *string `json:"StatusDesc,omitnil,omitempty" name:"StatusDesc"`

	// <p>CPU cores of the storage node</p>
	StorageNodeCpu *int64 `json:"StorageNodeCpu,omitnil,omitempty" name:"StorageNodeCpu"`

	// <p>Storage node memory size</p>
	StorageNodeMem *int64 `json:"StorageNodeMem,omitnil,omitempty" name:"StorageNodeMem"`

	// <p>Renewal flag</p>
	RenewFlag *int64 `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// <p>Payment mode, 0 pay-as-you-go, 1 prepaid</p>
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// <p>Expiration time</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	ExpireAt *string `json:"ExpireAt,omitnil,omitempty" name:"ExpireAt"`

	// <p>Isolation time</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsolatedAt *string `json:"IsolatedAt,omitnil,omitempty" name:"IsolatedAt"`

	// <p>Instance Architecture Type, separate: decoupled architecture; hybrid: peer-to-peer architecture</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// <p>Disk Type, CLOUD_HSSD enhanced SSD, CLOUD_TCS local SSD disk</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	StorageType *string `json:"StorageType,omitnil,omitempty" name:"StorageType"`

	// <p>List of instance node availability zones. Zones[0] refers to the primary AZ</p>
	Zones []*string `json:"Zones,omitnil,omitempty" name:"Zones"`

	// <p>Disk usage of the largest node</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	DiskUsage *int64 `json:"DiskUsage,omitnil,omitempty" name:"DiskUsage"`

	// <p>Whether binlog is enabled</p>
	BinlogStatus *int64 `json:"BinlogStatus,omitnil,omitempty" name:"BinlogStatus"`

	// <p>az mode. 1: Single az 2: Multi-az non-primary az mode 3: Multi-az primary az mode</p>
	AZMode *int64 `json:"AZMode,omitnil,omitempty" name:"AZMode"`

	// <p>Disaster recovery flag. 1: No disaster recovery relationship; 2: Primary instance for disaster recovery; 3: Disaster Recovery Standby Instance</p>
	StandbyFlag *int64 `json:"StandbyFlag,omitnil,omitempty" name:"StandbyFlag"`

	// <p>cdc node type</p>
	BinlogType *string `json:"BinlogType,omitnil,omitempty" name:"BinlogType"`

	// <p>1 means supported, 0 means unsupported</p>
	TimingModifyInstanceFlag *int64 `json:"TimingModifyInstanceFlag,omitnil,omitempty" name:"TimingModifyInstanceFlag"`

	// <p>cpu cores of the columnar node</p>
	ColumnarNodeCpu *int64 `json:"ColumnarNodeCpu,omitnil,omitempty" name:"ColumnarNodeCpu"`

	// <p>Columnar node memory size</p>
	ColumnarNodeMem *int64 `json:"ColumnarNodeMem,omitnil,omitempty" name:"ColumnarNodeMem"`

	// <p>Number of columnar nodes</p>
	ColumnarNodeNum *int64 `json:"ColumnarNodeNum,omitnil,omitempty" name:"ColumnarNodeNum"`

	// <p>Columnar node disk size</p>
	ColumnarNodeDisk *int64 `json:"ColumnarNodeDisk,omitnil,omitempty" name:"ColumnarNodeDisk"`

	// <p>Columnar node disk type</p>
	ColumnarNodeStorageType *string `json:"ColumnarNodeStorageType,omitnil,omitempty" name:"ColumnarNodeStorageType"`

	// <p>Columnar node specification</p>
	ColumnarNodeSpecCode *string `json:"ColumnarNodeSpecCode,omitnil,omitempty" name:"ColumnarNodeSpecCode"`

	// <p>Columnar storage vip</p>
	ColumnarVip *string `json:"ColumnarVip,omitnil,omitempty" name:"ColumnarVip"`

	// <p>Columnar vport</p>
	ColumnarVport *int64 `json:"ColumnarVport,omitnil,omitempty" name:"ColumnarVport"`

	// <p>Whether the instance supports columnar storage</p>
	IsSupportColumnar *bool `json:"IsSupportColumnar,omitnil,omitempty" name:"IsSupportColumnar"`

	// <p>Instance type</p>
	InstanceCategory *int64 `json:"InstanceCategory,omitnil,omitempty" name:"InstanceCategory"`

	// <p>Compatible mode</p>
	SQLMode *string `json:"SQLMode,omitnil,omitempty" name:"SQLMode"`

	// <p>Whether modification of the number of replicas is supported</p>
	IsSwitchFullReplicationsEnable *bool `json:"IsSwitchFullReplicationsEnable,omitnil,omitempty" name:"IsSwitchFullReplicationsEnable"`

	// <p>Instance type</p>
	InstanceMode *string `json:"InstanceMode,omitnil,omitempty" name:"InstanceMode"`

	// <p>dumper vip</p>
	DumperVip *string `json:"DumperVip,omitnil,omitempty" name:"DumperVip"`

	// <p>dumper vport</p>
	DumperVport *int64 `json:"DumperVport,omitnil,omitempty" name:"DumperVport"`

	// <p>ccu adjustment range of the svls instance</p>
	AutoScaleConfig *AutoScalingConfig `json:"AutoScaleConfig,omitnil,omitempty" name:"AutoScaleConfig"`

	// <p>Parameter template id</p>
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// <p>Parameter template name</p>
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// <p>Instance analysis engine mode</p><p>Enumeration value:</p><ul><li>libra: LibraDB analysis engine instance</li></ul>
	AnalysisMode *string `json:"AnalysisMode,omitnil,omitempty" name:"AnalysisMode"`

	// <p>Analysis engine instance relationship</p>
	AnalysisRelationInfos []*AnalysisRelationInfo `json:"AnalysisRelationInfos,omitnil,omitempty" name:"AnalysisRelationInfos"`

	// <p>Analysis engine instance info</p><p>Cpu, Memory, and Disk reuse top-level fields</p>
	AnalysisInstanceInfo *AnalysisInstanceInfo `json:"AnalysisInstanceInfo,omitnil,omitempty" name:"AnalysisInstanceInfo"`

	// <p>Maintenance window configuration</p>
	MaintenanceWindow *MaintenanceWindowInfo `json:"MaintenanceWindow,omitnil,omitempty" name:"MaintenanceWindow"`

	// <p>Whether transparent data encryption is enabled. 0: not enabled; 1: enabled</p>
	EncryptionEnable *int64 `json:"EncryptionEnable,omitnil,omitempty" name:"EncryptionEnable"`

	// <p>Real-use kms region for subsequent call to kms service</p>
	EncryptionKmsRegion *string `json:"EncryptionKmsRegion,omitnil,omitempty" name:"EncryptionKmsRegion"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBInstanceDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBInstanceDetailResponseParams `json:"Response"`
}

func (r *DescribeDBInstanceDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstanceDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstancesRequestParams struct {
	// <p>Filter parameters</p>
	Filters []*InstanceFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>Maximum return count, defaults to 20, maximum 100</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>Offset, which is an integer multiple of Limit.</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>Specified query engine type</p><p>Enumeration value:</p><ul><li>libra: Column storage engine</li></ul>
	EngineType *string `json:"EngineType,omitnil,omitempty" name:"EngineType"`
}

type DescribeDBInstancesRequest struct {
	*tchttp.BaseRequest
	
	// <p>Filter parameters</p>
	Filters []*InstanceFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>Maximum return count, defaults to 20, maximum 100</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>Offset, which is an integer multiple of Limit.</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>Specified query engine type</p><p>Enumeration value:</p><ul><li>libra: Column storage engine</li></ul>
	EngineType *string `json:"EngineType,omitnil,omitempty" name:"EngineType"`
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
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "EngineType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstancesResponseParams struct {
	// <p>Return to instance list information</p>
	Instances []*InstanceInfo `json:"Instances,omitnil,omitempty" name:"Instances"`

	// <p>Total number of conditions met</p>
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

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
type DescribeDBParametersRequestParams struct {
	// Instance ID, for example: tdsql3-ow728lmc.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeDBParametersRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID, for example: tdsql3-ow728lmc.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeDBParametersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBParametersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBParametersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBParametersResponseParams struct {
	// Instance ID, for example: tdsql3-ow728lmc.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Request the current parameter value of the instance.
	Params []*ParamDesc `json:"Params,omitnil,omitempty" name:"Params"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBParametersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBParametersResponseParams `json:"Response"`
}

func (r *DescribeDBParametersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBParametersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBSArchiveLogsRequestParams struct {
	// <p>Instance ID.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Logging ID</p>
	ArchiveLogId *int64 `json:"ArchiveLogId,omitnil,omitempty" name:"ArchiveLogId"`

	// <p>End time.</p>
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>Backup status: pending, running, success, failed</p>
	FilterStatus *string `json:"FilterStatus,omitnil,omitempty" name:"FilterStatus"`

	// <p>Limit on number</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>Offset.</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>Sorting field, enumerate: StartTime, EndTime, ExpiredTime, FileSize, BackupDuration</p>
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// <p>Sorting method: ASC: sequential, DESC: reverse</p>
	OrderType *string `json:"OrderType,omitnil,omitempty" name:"OrderType"`

	// <p>Start time.</p>
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`
}

type DescribeDBSArchiveLogsRequest struct {
	*tchttp.BaseRequest
	
	// <p>Instance ID.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Logging ID</p>
	ArchiveLogId *int64 `json:"ArchiveLogId,omitnil,omitempty" name:"ArchiveLogId"`

	// <p>End time.</p>
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>Backup status: pending, running, success, failed</p>
	FilterStatus *string `json:"FilterStatus,omitnil,omitempty" name:"FilterStatus"`

	// <p>Limit on number</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>Offset.</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>Sorting field, enumerate: StartTime, EndTime, ExpiredTime, FileSize, BackupDuration</p>
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// <p>Sorting method: ASC: sequential, DESC: reverse</p>
	OrderType *string `json:"OrderType,omitnil,omitempty" name:"OrderType"`

	// <p>Start time.</p>
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`
}

func (r *DescribeDBSArchiveLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBSArchiveLogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ArchiveLogId")
	delete(f, "EndTime")
	delete(f, "FilterStatus")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "OrderBy")
	delete(f, "OrderType")
	delete(f, "StartTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBSArchiveLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBSArchiveLogsResponseParams struct {
	// <p>Archivelog list</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	Items []*ArchiveLogModel `json:"Items,omitnil,omitempty" name:"Items"`

	// <p>Total.</p>
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBSArchiveLogsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBSArchiveLogsResponseParams `json:"Response"`
}

func (r *DescribeDBSArchiveLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBSArchiveLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBSAvailableRecoveryTimeRequestParams struct {
	// <p>Instance ID.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeDBSAvailableRecoveryTimeRequest struct {
	*tchttp.BaseRequest
	
	// <p>Instance ID.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeDBSAvailableRecoveryTimeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBSAvailableRecoveryTimeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBSAvailableRecoveryTimeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBSAvailableRecoveryTimeResponseParams struct {
	// <P>End time</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <P>Recoverable time interval.</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	Intervals []*ArchiveLogInterval `json:"Intervals,omitnil,omitempty" name:"Intervals"`

	// <P>Start time</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBSAvailableRecoveryTimeResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBSAvailableRecoveryTimeResponseParams `json:"Response"`
}

func (r *DescribeDBSAvailableRecoveryTimeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBSAvailableRecoveryTimeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBSBackupPolicyRequestParams struct {
	// <p>Instance ID.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeDBSBackupPolicyRequest struct {
	*tchttp.BaseRequest
	
	// <p>Instance ID.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeDBSBackupPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBSBackupPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBSBackupPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBSBackupPolicyResponseParams struct {
	// <p>Backup policy list</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	Items []*BackupPolicyModelOutPut `json:"Items,omitnil,omitempty" name:"Items"`

	// <p>Total.</p>
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBSBackupPolicyResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBSBackupPolicyResponseParams `json:"Response"`
}

func (r *DescribeDBSBackupPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBSBackupPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBSBackupSetsRequestParams struct {
	// <p>Instance ID.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Instance Backup set ID</p>
	BackupSetId *int64 `json:"BackupSetId,omitnil,omitempty" name:"BackupSetId"`

	// <p>End time.</p>
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>Filtering Conditions</p>
	FilterBy *BackupSetsReqFilter `json:"FilterBy,omitnil,omitempty" name:"FilterBy"`

	// <p>Query count [0,200]</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>Query offset [0,INF] this time.</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>StartTime,EndTime,ExpiredTime,BackupSetId,BackupDuration</p>
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// <p>ASC,DESC</p>
	OrderType *string `json:"OrderType,omitnil,omitempty" name:"OrderType"`

	// <p>Start time.</p>
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`
}

type DescribeDBSBackupSetsRequest struct {
	*tchttp.BaseRequest
	
	// <p>Instance ID.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Instance Backup set ID</p>
	BackupSetId *int64 `json:"BackupSetId,omitnil,omitempty" name:"BackupSetId"`

	// <p>End time.</p>
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>Filtering Conditions</p>
	FilterBy *BackupSetsReqFilter `json:"FilterBy,omitnil,omitempty" name:"FilterBy"`

	// <p>Query count [0,200]</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>Query offset [0,INF] this time.</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>StartTime,EndTime,ExpiredTime,BackupSetId,BackupDuration</p>
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// <p>ASC,DESC</p>
	OrderType *string `json:"OrderType,omitnil,omitempty" name:"OrderType"`

	// <p>Start time.</p>
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`
}

func (r *DescribeDBSBackupSetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBSBackupSetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BackupSetId")
	delete(f, "EndTime")
	delete(f, "FilterBy")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "OrderBy")
	delete(f, "OrderType")
	delete(f, "StartTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBSBackupSetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBSBackupSetsResponseParams struct {
	// <p>Backup set list</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	Items []*BackupSetModel `json:"Items,omitnil,omitempty" name:"Items"`

	// <p>Total.</p>
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBSBackupSetsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBSBackupSetsResponseParams `json:"Response"`
}

func (r *DescribeDBSBackupSetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBSBackupSetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBSBackupStatisticsDetailRequestParams struct {
	// <p>Instance ID.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>End time.</p>
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>Start time.</p>
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`
}

type DescribeDBSBackupStatisticsDetailRequest struct {
	*tchttp.BaseRequest
	
	// <p>Instance ID.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>End time.</p>
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>Start time.</p>
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`
}

func (r *DescribeDBSBackupStatisticsDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBSBackupStatisticsDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "EndTime")
	delete(f, "StartTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBSBackupStatisticsDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBSBackupStatisticsDetailResponseParams struct {
	// <p>Statistics by backup method</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	BackupMethodStatistics *BackupMethodStatisticsOutPut `json:"BackupMethodStatistics,omitnil,omitempty" name:"BackupMethodStatistics"`

	// <p>Backup time</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	BackupTime []*string `json:"BackupTime,omitnil,omitempty" name:"BackupTime"`

	// <p>Data type statistics by backup</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	BackupTypeStatistics *BackupTypeStatisticsModel `json:"BackupTypeStatistics,omitnil,omitempty" name:"BackupTypeStatistics"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBSBackupStatisticsDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBSBackupStatisticsDetailResponseParams `json:"Response"`
}

func (r *DescribeDBSBackupStatisticsDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBSBackupStatisticsDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBSBackupStatisticsRequestParams struct {
	// <p>Instance ID.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>End time.</p>
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>Start time.</p>
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`
}

type DescribeDBSBackupStatisticsRequest struct {
	*tchttp.BaseRequest
	
	// <p>Instance ID.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>End time.</p>
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>Start time.</p>
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`
}

func (r *DescribeDBSBackupStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBSBackupStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "EndTime")
	delete(f, "StartTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBSBackupStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBSBackupStatisticsResponseParams struct {
	// <p>Backup method statistics</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	BackupMethodStatistics *BackupMethodStatisticsModel `json:"BackupMethodStatistics,omitnil,omitempty" name:"BackupMethodStatistics"`

	// <p>Backup set statistics</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	BackupStatistics *BackupStatisticsModel `json:"BackupStatistics,omitnil,omitempty" name:"BackupStatistics"`

	// <p>Backup statistics.</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	DataBackupStatistics *DataBackupStatisticsModel `json:"DataBackupStatistics,omitnil,omitempty" name:"DataBackupStatistics"`

	// <p>Log backup statistics</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	LogBackupStatistics *LogBackupStatisticsModel `json:"LogBackupStatistics,omitnil,omitempty" name:"LogBackupStatistics"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBSBackupStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBSBackupStatisticsResponseParams `json:"Response"`
}

func (r *DescribeDBSBackupStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBSBackupStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBSCloneInstancesRequestParams struct {
	// <p>Source instance ID.</p>
	SourceInstanceId *string `json:"SourceInstanceId,omitnil,omitempty" name:"SourceInstanceId"`

	// <P>Engine type</p>
	DBType *string `json:"DBType,omitnil,omitempty" name:"DBType"`

	// <P>Creation end time</p>
	EndCreateTime *string `json:"EndCreateTime,omitnil,omitempty" name:"EndCreateTime"`

	// <p>Clone kind: PITR time point BackupSet backup set</p>
	FilterCloneType *string `json:"FilterCloneType,omitnil,omitempty" name:"FilterCloneType"`

	// <P>Query count [0,200]</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>Query offset [0,INF]</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>Sorting field: CloneTime, CreateTime</p>
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// <p>Sorting type: ASC, DESC</p>
	OrderType *string `json:"OrderType,omitnil,omitempty" name:"OrderType"`

	// <P>Creation start time</p>
	StartCreateTime *string `json:"StartCreateTime,omitnil,omitempty" name:"StartCreateTime"`
}

type DescribeDBSCloneInstancesRequest struct {
	*tchttp.BaseRequest
	
	// <p>Source instance ID.</p>
	SourceInstanceId *string `json:"SourceInstanceId,omitnil,omitempty" name:"SourceInstanceId"`

	// <P>Engine type</p>
	DBType *string `json:"DBType,omitnil,omitempty" name:"DBType"`

	// <P>Creation end time</p>
	EndCreateTime *string `json:"EndCreateTime,omitnil,omitempty" name:"EndCreateTime"`

	// <p>Clone kind: PITR time point BackupSet backup set</p>
	FilterCloneType *string `json:"FilterCloneType,omitnil,omitempty" name:"FilterCloneType"`

	// <P>Query count [0,200]</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>Query offset [0,INF]</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>Sorting field: CloneTime, CreateTime</p>
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// <p>Sorting type: ASC, DESC</p>
	OrderType *string `json:"OrderType,omitnil,omitempty" name:"OrderType"`

	// <P>Creation start time</p>
	StartCreateTime *string `json:"StartCreateTime,omitnil,omitempty" name:"StartCreateTime"`
}

func (r *DescribeDBSCloneInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBSCloneInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SourceInstanceId")
	delete(f, "DBType")
	delete(f, "EndCreateTime")
	delete(f, "FilterCloneType")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "OrderBy")
	delete(f, "OrderType")
	delete(f, "StartCreateTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBSCloneInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBSCloneInstancesResponseParams struct {
	// <P>Clone list</p>.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Items []*CloneInstanceModel `json:"Items,omitnil,omitempty" name:"Items"`

	// <p>Total</p>
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBSCloneInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBSCloneInstancesResponseParams `json:"Response"`
}

func (r *DescribeDBSCloneInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBSCloneInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBSecurityGroupsRequestParams struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeDBSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBSecurityGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBSecurityGroupsResponseParams struct {
	// Security group description.
	Groups []*SecurityGroup `json:"Groups,omitnil,omitempty" name:"Groups"`

	// Instance VIP.
	// Note: This field may return null, indicating that no valid values can be obtained.
	VIP *string `json:"VIP,omitnil,omitempty" name:"VIP"`

	// Instance port.
	// Note: This field may return null, indicating that no valid values can be obtained.
	VPort *string `json:"VPort,omitnil,omitempty" name:"VPort"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBSecurityGroupsResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeDatabaseObjectsRequestParams struct {
	// <p>Instance ID, such as tdsql3-42f40429.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Database name, obtained via the DescribeDatabases API.</p>
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// <p>Pagination index</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>Number of items per page</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>Table name matching expression</p>
	TableRegexp *string `json:"TableRegexp,omitnil,omitempty" name:"TableRegexp"`
}

type DescribeDatabaseObjectsRequest struct {
	*tchttp.BaseRequest
	
	// <p>Instance ID, such as tdsql3-42f40429.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Database name, obtained via the DescribeDatabases API.</p>
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// <p>Pagination index</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>Number of items per page</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>Table name matching expression</p>
	TableRegexp *string `json:"TableRegexp,omitnil,omitempty" name:"TableRegexp"`
}

func (r *DescribeDatabaseObjectsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatabaseObjectsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "DbName")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "TableRegexp")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDatabaseObjectsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatabaseObjectsResponseParams struct {
	// <p>Passthrough input parameter.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Database name.</p>
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// <p>Table list.</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	Tables []*DatabaseTable `json:"Tables,omitnil,omitempty" name:"Tables"`

	// <p>View list.</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	Views []*DatabaseView `json:"Views,omitnil,omitempty" name:"Views"`

	// <p>Stored procedure list.</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	Procs []*DatabaseProcedure `json:"Procs,omitnil,omitempty" name:"Procs"`

	// <p>Function list.</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	Funcs []*DatabaseFunction `json:"Funcs,omitnil,omitempty" name:"Funcs"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDatabaseObjectsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDatabaseObjectsResponseParams `json:"Response"`
}

func (r *DescribeDatabaseObjectsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatabaseObjectsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatabasesRequestParams struct {
	// <p>Instance ID, such as tdsql3-ow7t8lmc.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Pagination index</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>Number of items per page</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>Database name matching expression</p>
	DatabaseRegexp *string `json:"DatabaseRegexp,omitnil,omitempty" name:"DatabaseRegexp"`
}

type DescribeDatabasesRequest struct {
	*tchttp.BaseRequest
	
	// <p>Instance ID, such as tdsql3-ow7t8lmc.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Pagination index</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>Number of items per page</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>Database name matching expression</p>
	DatabaseRegexp *string `json:"DatabaseRegexp,omitnil,omitempty" name:"DatabaseRegexp"`
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

// Predefined struct for user
type DescribeDatabasesResponseParams struct {
	// <p>Passthrough instance ID</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>The database list on the instance.</p>
	Databases []*Database `json:"Databases,omitnil,omitempty" name:"Databases"`

	// <p>Total quantity.</p>
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDatabasesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDatabasesResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeFlowRequestParams struct {

}

type DescribeFlowRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeFlowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFlowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFlowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFlowResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeFlowResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFlowResponseParams `json:"Response"`
}

func (r *DescribeFlowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFlowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceSSLStatusRequestParams struct {
	// <p>Instance ID.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeInstanceSSLStatusRequest struct {
	*tchttp.BaseRequest
	
	// <p>Instance ID.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeInstanceSSLStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceSSLStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceSSLStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceSSLStatusResponseParams struct {
	// <p>SSL enable status</p><p>Enumeration values:</p><ul><li>Enabled: SSL is on</li><li>Disabled: SSL is closed</li><li>Enabling: SSL is enabling</li><li>Disabling: SSL is disabling</li></ul>
	SSLStatus *string `json:"SSLStatus,omitnil,omitempty" name:"SSLStatus"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceSSLStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceSSLStatusResponseParams `json:"Response"`
}

func (r *DescribeInstanceSSLStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceSSLStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMaintenanceWindowRequestParams struct {
	// <p>Instance ID.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeMaintenanceWindowRequest struct {
	*tchttp.BaseRequest
	
	// <p>Instance ID.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeMaintenanceWindowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMaintenanceWindowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMaintenanceWindowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMaintenanceWindowResponseParams struct {
	// <p>Instance ID.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Ops window time range</p>
	MaintenanceWindow *string `json:"MaintenanceWindow,omitnil,omitempty" name:"MaintenanceWindow"`

	// <p>Ops window number of days range</p>
	WeekDays []*string `json:"WeekDays,omitnil,omitempty" name:"WeekDays"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMaintenanceWindowResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMaintenanceWindowResponseParams `json:"Response"`
}

func (r *DescribeMaintenanceWindowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMaintenanceWindowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSaleInfoRequestParams struct {
	// <p>Region of the disaster recovery primary instance</p>
	SrcRegion *string `json:"SrcRegion,omitnil,omitempty" name:"SrcRegion"`

	// <p>Instance id</p><p>Input this parameter to return the sales information of the availability zone in the region where this instance is located.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Specify the sales information that supports a certain type instance.</p><p>Enumeration value:</p><ul><li>serverless: Returns all regions that support serverless instance type.</li></ul><p>Default value: None</p><p>Currently only support specifying serverless. Import other info or leave it blank to return all sales region information by default.</p>
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`
}

type DescribeSaleInfoRequest struct {
	*tchttp.BaseRequest
	
	// <p>Region of the disaster recovery primary instance</p>
	SrcRegion *string `json:"SrcRegion,omitnil,omitempty" name:"SrcRegion"`

	// <p>Instance id</p><p>Input this parameter to return the sales information of the availability zone in the region where this instance is located.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Specify the sales information that supports a certain type instance.</p><p>Enumeration value:</p><ul><li>serverless: Returns all regions that support serverless instance type.</li></ul><p>Default value: None</p><p>Currently only support specifying serverless. Import other info or leave it blank to return all sales region information by default.</p>
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`
}

func (r *DescribeSaleInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSaleInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SrcRegion")
	delete(f, "InstanceId")
	delete(f, "InstanceType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSaleInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSaleInfoResponseParams struct {
	// <p>Return marketable region information</p>
	RegionList []*DescribeSaleRegionInfo `json:"RegionList,omitnil,omitempty" name:"RegionList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSaleInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSaleInfoResponseParams `json:"Response"`
}

func (r *DescribeSaleInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSaleInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSaleRegionInfo struct {
	// <p>English string of Region</p>
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// <p>Purchasable Zone list</p>
	ZoneList []*DescribeSaleZonesInfo `json:"ZoneList,omitnil,omitempty" name:"ZoneList"`

	// <p>Region Chinese character string</p>
	RegionName *string `json:"RegionName,omitnil,omitempty" name:"RegionName"`

	// <p>Whether to sell. 1: sell, 0: do not sell</p>
	Enable *int64 `json:"Enable,omitnil,omitempty" name:"Enable"`

	// <p>Optional quantity of multiple availability</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	AvailableZoneNum *int64 `json:"AvailableZoneNum,omitnil,omitempty" name:"AvailableZoneNum"`

	// <p>Whether to allow usage log replica</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsSupportedLogReplica *bool `json:"IsSupportedLogReplica,omitnil,omitempty" name:"IsSupportedLogReplica"`

	// <p>Availability zone combination</p>
	ZoneGroup []*DescribeSaleZonesGroup `json:"ZoneGroup,omitnil,omitempty" name:"ZoneGroup"`

	// <p>Whether serverless is supported</p>
	IsSupportServerless *bool `json:"IsSupportServerless,omitnil,omitempty" name:"IsSupportServerless"`
}

type DescribeSaleZonesGroup struct {
	// <p>Number of availability zones</p>
	ZoneNum *int64 `json:"ZoneNum,omitnil,omitempty" name:"ZoneNum"`

	// <p>Availability zone combination</p>
	Zones []*string `json:"Zones,omitnil,omitempty" name:"Zones"`

	// <p>Supported types</p>
	SupportTypes []*string `json:"SupportTypes,omitnil,omitempty" name:"SupportTypes"`

	// <p>Supported disk</p><p>Enumeration value:</p><ul><li>CLOUD_TCS: Local disk</li><li>CLOUD_HSSD: Enhanced cloud disk</li></ul>
	// Note: This field may return null, indicating that no valid values can be obtained.
	AvailableDiskTypes []*string `json:"AvailableDiskTypes,omitnil,omitempty" name:"AvailableDiskTypes"`

	// <p>Whether serverless is supported</p>
	IsSupportServerless *bool `json:"IsSupportServerless,omitnil,omitempty" name:"IsSupportServerless"`
}

type DescribeSaleZonesInfo struct {
	// <p>English string of Zone</p>
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// <p>Zone Chinese character string</p>
	ZoneName *string `json:"ZoneName,omitnil,omitempty" name:"ZoneName"`

	// <p>Whether to sell, 1: sell; 0: do not sell</p>
	Enable *int64 `json:"Enable,omitnil,omitempty" name:"Enable"`

	// <p>Whether the default primary AZ is used. 0 means no, 1 means yes.</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsDefaultMaster *int64 `json:"IsDefaultMaster,omitnil,omitempty" name:"IsDefaultMaster"`

	// <p>Selectable disk types in availability zones</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	AvailableDiskTypes []*string `json:"AvailableDiskTypes,omitnil,omitempty" name:"AvailableDiskTypes"`

	// <p>Whether it is an exclusive availability zone</p>
	SupportTypes []*string `json:"SupportTypes,omitnil,omitempty" name:"SupportTypes"`

	// <p>Whether serverless is supported</p>
	IsSupportServerless *bool `json:"IsSupportServerless,omitnil,omitempty" name:"IsSupportServerless"`
}

// Predefined struct for user
type DescribeSlowLogsRequestParams struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Search log start time
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time to retrieve logs
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Filter criteria.
	LogFilter []*LogFilter `json:"LogFilter,omitnil,omitempty" name:"LogFilter"`

	// Items per page limit
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Sort, selectable: ASC, DESC
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// Sorting criteria may not be the same as selectable fields used to sort according to business.
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`
}

type DescribeSlowLogsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Search log start time
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time to retrieve logs
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Filter criteria.
	LogFilter []*LogFilter `json:"LogFilter,omitnil,omitempty" name:"LogFilter"`

	// Items per page limit
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Sort, selectable: ASC, DESC
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// Sorting criteria may not be the same as selectable fields used to sort according to business.
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`
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
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "LogFilter")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Order")
	delete(f, "OrderBy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSlowLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSlowLogsResponseParams struct {
	// Total number of logs
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Log details.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Items []*SlowLogData `json:"Items,omitnil,omitempty" name:"Items"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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
type DescribeSpecsRequestParams struct {
	// <p>Number of replicas. Currently supported range: 1-3. Default is 3.</p>
	FullReplications *int64 `json:"FullReplications,omitnil,omitempty" name:"FullReplications"`

	// <p>Exclusive instance</p>
	IsExclusiveInstance *int64 `json:"IsExclusiveInstance,omitnil,omitempty" name:"IsExclusiveInstance"`
}

type DescribeSpecsRequest struct {
	*tchttp.BaseRequest
	
	// <p>Number of replicas. Currently supported range: 1-3. Default is 3.</p>
	FullReplications *int64 `json:"FullReplications,omitnil,omitempty" name:"FullReplications"`

	// <p>Exclusive instance</p>
	IsExclusiveInstance *int64 `json:"IsExclusiveInstance,omitnil,omitempty" name:"IsExclusiveInstance"`
}

func (r *DescribeSpecsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSpecsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FullReplications")
	delete(f, "IsExclusiveInstance")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSpecsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSpecsResponseParams struct {
	// <p>Purchasable specification list of peer nodes</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	HybridNodeSpecs []*StorageNodeSpec `json:"HybridNodeSpecs,omitnil,omitempty" name:"HybridNodeSpecs"`

	// <p>Purchasable specification list of svls nodes</p>
	ServerlessCcuSpec []*ServerlessCcu `json:"ServerlessCcuSpec,omitnil,omitempty" name:"ServerlessCcuSpec"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSpecsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSpecsResponseParams `json:"Response"`
}

func (r *DescribeSpecsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSpecsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserPrivilegesRequestParams struct {
	// Instance ID, such as tdsql3-5baee8df.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Accessing host allowed for the user. Username+host uniquely determines an account.
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// Login username.
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// Database name. If it is \*, query global permission (\*.\*) and ignore the Type and Object parameter.
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// The name of the specific Type, for example, when Type is table, it is exactly the table name. If both DbName and Type are function names, Object represents the specific object name and cannot be \* or empty.
	Object *string `json:"Object,omitnil,omitempty" name:"Object"`

	// Type, can be set to table, view, proc, func, and \*. When DbName is a specific database name and Type is \*, it queries the database permission (i.e., db.\*), ignoring the Object parameter.
	ObjectType *string `json:"ObjectType,omitnil,omitempty" name:"ObjectType"`

	// When Type=table, ColName as \* indicates the permission to query the table. If it is a specific field name, it indicates the permission to query the corresponding field.
	ColName *string `json:"ColName,omitnil,omitempty" name:"ColName"`
}

type DescribeUserPrivilegesRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID, such as tdsql3-5baee8df.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Accessing host allowed for the user. Username+host uniquely determines an account.
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// Login username.
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// Database name. If it is \*, query global permission (\*.\*) and ignore the Type and Object parameter.
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// The name of the specific Type, for example, when Type is table, it is exactly the table name. If both DbName and Type are function names, Object represents the specific object name and cannot be \* or empty.
	Object *string `json:"Object,omitnil,omitempty" name:"Object"`

	// Type, can be set to table, view, proc, func, and \*. When DbName is a specific database name and Type is \*, it queries the database permission (i.e., db.\*), ignoring the Object parameter.
	ObjectType *string `json:"ObjectType,omitnil,omitempty" name:"ObjectType"`

	// When Type=table, ColName as \* indicates the permission to query the table. If it is a specific field name, it indicates the permission to query the corresponding field.
	ColName *string `json:"ColName,omitnil,omitempty" name:"ColName"`
}

func (r *DescribeUserPrivilegesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserPrivilegesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Host")
	delete(f, "UserName")
	delete(f, "DbName")
	delete(f, "Object")
	delete(f, "ObjectType")
	delete(f, "ColName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserPrivilegesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserPrivilegesResponseParams struct {
	// Permission list
	// Note: This field may return null, indicating that no valid values can be obtained.
	Privileges []*string `json:"Privileges,omitnil,omitempty" name:"Privileges"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeUserPrivilegesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserPrivilegesResponseParams `json:"Response"`
}

func (r *DescribeUserPrivilegesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserPrivilegesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUsersRequestParams struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeUsersRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeUsersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUsersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUsersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUsersResponseParams struct {
	// User list.
	Users []*UserInfo `json:"Users,omitnil,omitempty" name:"Users"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeUsersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUsersResponseParams `json:"Response"`
}

func (r *DescribeUsersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUsersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroyInstancesRequestParams struct {
	// List of isolated instance ids required.
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

type DestroyInstancesRequest struct {
	*tchttp.BaseRequest
	
	// List of isolated instance ids required.
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

func (r *DestroyInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DestroyInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroyInstancesResponseParams struct {
	// List of successfully isolated instance ids.
	SuccessInstanceIds []*string `json:"SuccessInstanceIds,omitnil,omitempty" name:"SuccessInstanceIds"`

	// Isolation removal failed instance Id list.
	FailedInstanceIds []*string `json:"FailedInstanceIds,omitnil,omitempty" name:"FailedInstanceIds"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DestroyInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DestroyInstancesResponseParams `json:"Response"`
}

func (r *DestroyInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExpandInstanceRequestParams struct {
	// <p>Instance ID.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Expand storage nodes to how many nodes. If no change, pass the current number of nodes</p>
	StorageNodeNum *int64 `json:"StorageNodeNum,omitnil,omitempty" name:"StorageNodeNum"`

	// <p>Availability zone list</p>
	Zones []*string `json:"Zones,omitnil,omitempty" name:"Zones"`

	// <p>az mode. 1: Single az 2: Multi-az non-primary az mode 3: Multi-az primary az mode</p>
	AZMode *int64 `json:"AZMode,omitnil,omitempty" name:"AZMode"`

	// <p>AZMode 3 means the primary AZ</p>
	PrimaryAZ *string `json:"PrimaryAZ,omitnil,omitempty" name:"PrimaryAZ"`

	// <p>Number of replicas, value ranges from 1 to 3</p>
	FullReplications *int64 `json:"FullReplications,omitnil,omitempty" name:"FullReplications"`
}

type ExpandInstanceRequest struct {
	*tchttp.BaseRequest
	
	// <p>Instance ID.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Expand storage nodes to how many nodes. If no change, pass the current number of nodes</p>
	StorageNodeNum *int64 `json:"StorageNodeNum,omitnil,omitempty" name:"StorageNodeNum"`

	// <p>Availability zone list</p>
	Zones []*string `json:"Zones,omitnil,omitempty" name:"Zones"`

	// <p>az mode. 1: Single az 2: Multi-az non-primary az mode 3: Multi-az primary az mode</p>
	AZMode *int64 `json:"AZMode,omitnil,omitempty" name:"AZMode"`

	// <p>AZMode 3 means the primary AZ</p>
	PrimaryAZ *string `json:"PrimaryAZ,omitnil,omitempty" name:"PrimaryAZ"`

	// <p>Number of replicas, value ranges from 1 to 3</p>
	FullReplications *int64 `json:"FullReplications,omitnil,omitempty" name:"FullReplications"`
}

func (r *ExpandInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExpandInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StorageNodeNum")
	delete(f, "Zones")
	delete(f, "AZMode")
	delete(f, "PrimaryAZ")
	delete(f, "FullReplications")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExpandInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExpandInstanceResponseParams struct {
	// <p>Asynchronous task ID. You can call the Query Task API to get task status with this ID.</p>
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ExpandInstanceResponse struct {
	*tchttp.BaseResponse
	Response *ExpandInstanceResponseParams `json:"Response"`
}

func (r *ExpandInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExpandInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Explain struct {
	// <p>Identifier</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`

	// <p>Query type</p><p>Enumeration value:</p><ul><li>SIMPLE: A regular query with no subquery and UNION. Single table or ordinary JOIN is SIMPLE.</li></ul>
	// Note: This field may return null, indicating that no valid values can be obtained.
	SelectType *string `json:"SelectType,omitnil,omitempty" name:"SelectType"`

	// <p>Table name</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	Table *string `json:"Table,omitnil,omitempty" name:"Table"`

	// <p>Partition</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	Partitions *string `json:"Partitions,omitnil,omitempty" name:"Partitions"`

	// <p>Access type</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>Possibly used indexes</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	PossibleKeys *string `json:"PossibleKeys,omitnil,omitempty" name:"PossibleKeys"`

	// <p>Used Indexes</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// <p>Used Indexes length</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	KeyLen *string `json:"KeyLen,omitnil,omitempty" name:"KeyLen"`

	// <p>Column or constant to compare with the index</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	Ref *string `json:"Ref,omitnil,omitempty" name:"Ref"`

	// <p>Estimate the number of scanned rows</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	Rows *string `json:"Rows,omitnil,omitempty" name:"Rows"`

	// <p>Estimated percentage of remaining rows after conditional filtering</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	Filtered *string `json:"Filtered,omitnil,omitempty" name:"Filtered"`

	// <p>Additional information, such as Using index, Using filesort</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`
}

type InstanceFilter struct {
	// Filter key, support InstanceId, VpcId, SubnetId, Vip, Vport, Status, InstanceName, TagKey.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Filter value.
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type InstanceInfo struct {
	// <p>Number of compute nodes</p>
	//
	// Deprecated: ComputeNodeNum is deprecated.
	ComputeNodeNum *int64 `json:"ComputeNodeNum,omitnil,omitempty" name:"ComputeNodeNum"`

	// <p>Region</p>
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// <p>Creating an Instance Version</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	CreateVersion *string `json:"CreateVersion,omitnil,omitempty" name:"CreateVersion"`

	// <p>Initialize instance parameter</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	InitParams []*InstanceParam `json:"InitParams,omitnil,omitempty" name:"InitParams"`

	// <p>Instance status: creating, created, initializing, running, modifying, isolating, isolated, destroying, destroyed</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>Instance id</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Number of storage nodes</p>
	StorageNodeNum *int64 `json:"StorageNodeNum,omitnil,omitempty" name:"StorageNodeNum"`

	// <p>Instance tag information</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	ResourceTags []*ResourceTag `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`

	// <p>Instance name</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// <p>cpu cores of the computing node</p>
	//
	// Deprecated: Cpu is deprecated.
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// <p>Character type vpcid</p>
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// <p>Computing node mem, in GB</p>
	//
	// Deprecated: Mem is deprecated.
	Mem *int64 `json:"Mem,omitnil,omitempty" name:"Mem"`

	// <p>Subnet IP</p>
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// <p>Character type subnetid</p>
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// <p>Subnet port</p>
	Vport *int64 `json:"Vport,omitnil,omitempty" name:"Vport"`

	// <p>Node disk capacity (unit: GB)</p>
	Disk *int64 `json:"Disk,omitnil,omitempty" name:"Disk"`

	// <p>Instance Creation Time</p>
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// <p>Region of the instance</p>
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// <p>Status description in Chinese of the instance</p>
	StatusDesc *string `json:"StatusDesc,omitnil,omitempty" name:"StatusDesc"`

	// <p>CPU cores of the control node</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	//
	// Deprecated: MCCpu is deprecated.
	MCCpu *int64 `json:"MCCpu,omitnil,omitempty" name:"MCCpu"`

	// <p>CPU size of the control node</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	//
	// Deprecated: MCMem is deprecated.
	MCMem *int64 `json:"MCMem,omitnil,omitempty" name:"MCMem"`

	// <p>CPU cores of the computing node</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	//
	// Deprecated: ComputerNodeCpu is deprecated.
	ComputerNodeCpu *int64 `json:"ComputerNodeCpu,omitnil,omitempty" name:"ComputerNodeCpu"`

	// <p>Compute node memory size</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	//
	// Deprecated: ComputerNodeMem is deprecated.
	ComputerNodeMem *int64 `json:"ComputerNodeMem,omitnil,omitempty" name:"ComputerNodeMem"`

	// <p>CPU cores of the storage node</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	StorageNodeCpu *int64 `json:"StorageNodeCpu,omitnil,omitempty" name:"StorageNodeCpu"`

	// <p>Storage node memory size</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	StorageNodeMem *int64 `json:"StorageNodeMem,omitnil,omitempty" name:"StorageNodeMem"`

	// <p>Number of control nodes</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	//
	// Deprecated: MCNum is deprecated.
	MCNum *int64 `json:"MCNum,omitnil,omitempty" name:"MCNum"`

	// <p>Renewal flag</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	RenewFlag *int64 `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// <p>Payment mode, 0 pay-as-you-go; 1 annual/monthly subscription</p>
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// <p>User tag, inner: internal user; external: external user</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	AccountTag *string `json:"AccountTag,omitnil,omitempty" name:"AccountTag"`

	// <p>Instance Architecture Type, separate: decoupled architecture; hyper: peer-to-peer architecture</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// <p>Disk Type, CLOUD_HSSD enhanced SSD, CLOUD_TCS local SSD disk</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	StorageType *string `json:"StorageType,omitnil,omitempty" name:"StorageType"`

	// <p>&quot;0000-00-00 00:00:00&quot;</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	DestroyedAt *string `json:"DestroyedAt,omitnil,omitempty" name:"DestroyedAt"`

	// <p>&quot;0000-00-00 00:00:00&quot;</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	ExpireAt *string `json:"ExpireAt,omitnil,omitempty" name:"ExpireAt"`

	// <p>&quot;0000-00-00 00:00:00&quot;</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsolatedAt *string `json:"IsolatedAt,omitnil,omitempty" name:"IsolatedAt"`

	// <p>&quot;0000-00-00 00:00:00&quot;</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsolatedFrom *string `json:"IsolatedFrom,omitnil,omitempty" name:"IsolatedFrom"`

	// <p>1</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	Replications *int64 `json:"Replications,omitnil,omitempty" name:"Replications"`

	// <p>Number of replicas</p>
	FullReplications *int64 `json:"FullReplications,omitnil,omitempty" name:"FullReplications"`

	// <p>Account information</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Account information</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	SubAccountUin *string `json:"SubAccountUin,omitnil,omitempty" name:"SubAccountUin"`

	// <p>Account information</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// <p>AZ information</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	Zones []*string `json:"Zones,omitnil,omitempty" name:"Zones"`

	// <p>Instance node</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	Nodes []*InstanceNode `json:"Nodes,omitnil,omitempty" name:"Nodes"`

	// <p>Whether binlog is on</p>
	BinlogStatus *int64 `json:"BinlogStatus,omitnil,omitempty" name:"BinlogStatus"`

	// <p>Number of cdc node cores</p>
	//
	// Deprecated: CdcNodeCpu is deprecated.
	CdcNodeCpu *int64 `json:"CdcNodeCpu,omitnil,omitempty" name:"CdcNodeCpu"`

	// <p>cdc node memory size</p>
	//
	// Deprecated: CdcNodeMem is deprecated.
	CdcNodeMem *int64 `json:"CdcNodeMem,omitnil,omitempty" name:"CdcNodeMem"`

	// <p>Number of cdc nodes</p>
	//
	// Deprecated: CdcNodeNum is deprecated.
	CdcNodeNum *int64 `json:"CdcNodeNum,omitnil,omitempty" name:"CdcNodeNum"`

	// <p>az mode. 1: Single az, 2: Multi-az non-primary az mode, 3: Multi-az primary az mode</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	AZMode *int64 `json:"AZMode,omitnil,omitempty" name:"AZMode"`

	// <p>Disaster recovery flag. 1: No disaster recovery relationship; 2: Primary instance for disaster recovery; 3: Disaster Recovery Standby Instance</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	StandbyFlag *int64 `json:"StandbyFlag,omitnil,omitempty" name:"StandbyFlag"`

	// <p>Number of connected standby instances (Valid only when StandbyFlag == 2)</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	StandbySecondaryNum *int64 `json:"StandbySecondaryNum,omitnil,omitempty" name:"StandbySecondaryNum"`

	// <p>cpu cores of the columnar node</p>
	ColumnarNodeCpu *int64 `json:"ColumnarNodeCpu,omitnil,omitempty" name:"ColumnarNodeCpu"`

	// <p>Columnar node memory size</p>
	ColumnarNodeMem *int64 `json:"ColumnarNodeMem,omitnil,omitempty" name:"ColumnarNodeMem"`

	// <p>Number of columnar nodes</p>
	ColumnarNodeNum *int64 `json:"ColumnarNodeNum,omitnil,omitempty" name:"ColumnarNodeNum"`

	// <p>Columnar node disk capacity (unit: GB)</p>
	ColumnarNodeDisk *int64 `json:"ColumnarNodeDisk,omitnil,omitempty" name:"ColumnarNodeDisk"`

	// <p>Columnar node disk type</p>
	ColumnarNodeStorageType *string `json:"ColumnarNodeStorageType,omitnil,omitempty" name:"ColumnarNodeStorageType"`

	// <p>Exclusive flags, 1: Primary instance (dedicated), 2: Primary instance, 3: Disaster recovery instance, 4: Disaster recovery instance (dedicated)</p>
	InstanceCategory *int64 `json:"InstanceCategory,omitnil,omitempty" name:"InstanceCategory"`

	// <p>dbdc-xxxxx</p>
	ExclusiveClusterId *string `json:"ExclusiveClusterId,omitnil,omitempty" name:"ExclusiveClusterId"`

	// <p>Compatible mode</p>
	SQLMode *string `json:"SQLMode,omitnil,omitempty" name:"SQLMode"`

	// <p>Instance mode</p>
	InstanceMode *string `json:"InstanceMode,omitnil,omitempty" name:"InstanceMode"`

	// <p>Instance delivery platform</p>
	//
	// Deprecated: ClusterId is deprecated.
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// <p>Auto-scaling configuration</p>
	AutoScaleConfig *AutoScalingConfig `json:"AutoScaleConfig,omitnil,omitempty" name:"AutoScaleConfig"`

	// <p>Analytical engine mode</p><p>Enumeration value:</p><ul><li>libra: LibraDB analytical engine mode</li></ul>
	AnalysisMode *string `json:"AnalysisMode,omitnil,omitempty" name:"AnalysisMode"`

	// <p>Analysis engine relationship information</p>
	AnalysisRelationInfos []*AnalysisRelationInfo `json:"AnalysisRelationInfos,omitnil,omitempty" name:"AnalysisRelationInfos"`

	// <p>Analysis engine instance info</p>
	AnalysisInstanceInfo *AnalysisInstanceInfo `json:"AnalysisInstanceInfo,omitnil,omitempty" name:"AnalysisInstanceInfo"`
}

type InstanceNode struct {
	// Primary key
	// Note: This field may return null, indicating that no valid values can be obtained.
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// Instance ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Node Id
	// Note: This field may return null, indicating that no valid values can be obtained.
	NodeId *string `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// Instance Ip
	// Note: This field may return null, indicating that no valid values can be obtained.
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// Eni IP of the instance
	// Note: This field may return null, indicating that no valid values can be obtained.
	EniIp *string `json:"EniIp,omitnil,omitempty" name:"EniIp"`

	// Instance Port
	// Note: This field may return null, indicating that no valid values can be obtained.
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// Instance SpecCode
	// Note: This field may return null, indicating that no valid values can be obtained.
	SpecCode *string `json:"SpecCode,omitnil,omitempty" name:"SpecCode"`

	// Instance NodeName
	// Note: This field may return null, indicating that no valid values can be obtained.
	NodeName *string `json:"NodeName,omitnil,omitempty" name:"NodeName"`

	// Instance Cpu
	// Note: This field may return null, indicating that no valid values can be obtained.
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// Instance memory
	// Note: This field may return null, indicating that no valid values can be obtained.
	Mem *int64 `json:"Mem,omitnil,omitempty" name:"Mem"`

	// Instance Disk
	// Note: This field may return null, indicating that no valid values can be obtained.
	Disk *int64 `json:"Disk,omitnil,omitempty" name:"Disk"`

	// Instance type
	// Note: This field may return null, indicating that no valid values can be obtained.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Instance status
	// Note: This field may return null, indicating that no valid values can be obtained.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// instance version
	// Note: This field may return null, indicating that no valid values can be obtained.
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// Region
	// Note: This field may return null, indicating that no valid values can be obtained.
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Instance LocalDNS
	// Note: This field may return null, indicating that no valid values can be obtained.
	LocalDNS *string `json:"LocalDNS,omitnil,omitempty" name:"LocalDNS"`

	// Instance Region
	// Note: This field may return null, indicating that no valid values can be obtained.
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Instance log disk
	// Note: This field may return null, indicating that no valid values can be obtained.
	LogDisk *int64 `json:"LogDisk,omitnil,omitempty" name:"LogDisk"`

	// Instance data disk
	// Note: This field may return null, indicating that no valid values can be obtained.
	DataDisk *int64 `json:"DataDisk,omitnil,omitempty" name:"DataDisk"`

	// Zone ID of the instance
	// Note: This field may return null, indicating that no valid values can be obtained.
	ZoneID *string `json:"ZoneID,omitnil,omitempty" name:"ZoneID"`

	// Instance SpecName
	// Note: This field may return null, indicating that no valid values can be obtained.
	SpecName *string `json:"SpecName,omitnil,omitempty" name:"SpecName"`

	// Instance Replicas
	// Note: This field may return null, indicating that no valid values can be obtained.
	Replicas *int64 `json:"Replicas,omitnil,omitempty" name:"Replicas"`

	// Instance Shards
	// Note: This field may return null, indicating that no valid values can be obtained.
	Shards *int64 `json:"Shards,omitnil,omitempty" name:"Shards"`

	// Instance data replica
	// Note: This field may return null, indicating that no valid values can be obtained.
	DataReplicas *int64 `json:"DataReplicas,omitnil,omitempty" name:"DataReplicas"`

	// Initialize parameter
	// Note: This field may return null, indicating that no valid values can be obtained.
	Params *string `json:"Params,omitnil,omitempty" name:"Params"`

	// Storage medium, CLOUD_PREMIUM: Premium Cloud Disk, CLOUD_SSD: SSD cloud disk, CLOUD_HSSD: HSSD cloud disk
	// Note: This field may return null, indicating that no valid values can be obtained.
	StorageType *string `json:"StorageType,omitnil,omitempty" name:"StorageType"`
}

type InstanceParam struct {
	// config key
	Param *string `json:"Param,omitnil,omitempty" name:"Param"`

	// Config value
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

// Predefined struct for user
type IsolateDBInstanceRequestParams struct {
	// List of isolated instance ids required.
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

type IsolateDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// List of isolated instance ids required.
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
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
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "IsolateDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type IsolateDBInstanceResponseParams struct {
	// Isolation successful instance Id list.
	SuccessInstanceIds []*string `json:"SuccessInstanceIds,omitnil,omitempty" name:"SuccessInstanceIds"`

	// Isolation failed instance Id list.
	FailedInstanceIds []*string `json:"FailedInstanceIds,omitnil,omitempty" name:"FailedInstanceIds"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type IsolateDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *IsolateDBInstanceResponseParams `json:"Response"`
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

type LogBackupStatisticsModel struct {
	// <p>Avg size of each log backup in Byte</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	AverageSizePerBackup *int64 `json:"AverageSizePerBackup,omitnil,omitempty" name:"AverageSizePerBackup"`

	// <p>Avg daily log backup size in Byte</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	AverageSizePerDay *int64 `json:"AverageSizePerDay,omitnil,omitempty" name:"AverageSizePerDay"`

	// <p>Number of log backups</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// <p>Log backup size in Byte</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	TotalSize *int64 `json:"TotalSize,omitnil,omitempty" name:"TotalSize"`
}

type LogFilter struct {
	// Filter criterion name.
	// 
	// For example: sql - SQL command details
	// 
	// host - client IP
	// user - database account;
	// dbName – Database name;
	// sqlType - SQL type;
	// Error code
	// 
	// execTime - Execution time
	// lockWaitTime - Lock waiting time
	// ioWaitTime - IO wait time
	// trxLivingTime - Transaction execution time
	// Cpu time
	// 
	// threadId - Thread ID
	// trxId - Transaction ID
	// checkRows - Number of scanned rows
	// affectRows - Number of rows affected
	// sentRows - Number of rows returned
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Filter condition match type. Supported:
	// INC – Includes (multiple values are in a || relationship before).
	// EXC - excluding (multiple values are in an && relationship)
	// EQS – equal to (multiple values before are in a || relationship).
	// NEQ – not equal to (multiple values are in && relationship)
	// 
	// RG - Range;
	Compare *string `json:"Compare,omitnil,omitempty" name:"Compare"`

	// Filter condition matching value. When Compare=RG, for example ["1-100","200-300"].
	Value []*string `json:"Value,omitnil,omitempty" name:"Value"`
}

type MaintenanceWindowInfo struct {

	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`


	Duration *int64 `json:"Duration,omitnil,omitempty" name:"Duration"`


	WeekDays []*string `json:"WeekDays,omitnil,omitempty" name:"WeekDays"`
}

// Predefined struct for user
type ModifyAutoRenewFlagRequestParams struct {
	// <P>Instance list that needs to be modified</p>.
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// <P>1 enables automatic renewal, 0 disables automatic renewal.</p>.
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`
}

type ModifyAutoRenewFlagRequest struct {
	*tchttp.BaseRequest
	
	// <P>Instance list that needs to be modified</p>.
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// <P>1 enables automatic renewal, 0 disables automatic renewal.</p>.
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`
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
	delete(f, "AutoRenewFlag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAutoRenewFlagRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAutoRenewFlagResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAutoRenewFlagResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAutoRenewFlagResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyDBInstanceSecurityGroupsRequestParams struct {
	// Instance id.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Security group ID list to modify. an array of one or more security group ids.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`
}

type ModifyDBInstanceSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// Instance id.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Security group ID list to modify. an array of one or more security group ids.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDBInstanceSecurityGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceSecurityGroupsResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDBInstanceSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDBInstanceSecurityGroupsResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyDBInstanceVPortRequestParams struct {
	// Instance ID, such as tdsql3-5baee8df.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// New VPC port 3308
	Vport *int64 `json:"Vport,omitnil,omitempty" name:"Vport"`
}

type ModifyDBInstanceVPortRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID, such as tdsql3-5baee8df.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// New VPC port 3308
	Vport *int64 `json:"Vport,omitnil,omitempty" name:"Vport"`
}

func (r *ModifyDBInstanceVPortRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceVPortRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Vport")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDBInstanceVPortRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceVPortResponseParams struct {
	// Return the async task FlowId
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDBInstanceVPortResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDBInstanceVPortResponseParams `json:"Response"`
}

func (r *ModifyDBInstanceVPortResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceVPortResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBParametersRequestParams struct {
	// Instance ID, for example: tdsql3-ow728lmc.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Parameter list, each element is a combination of Param and Value.
	Params []*DBParamValue `json:"Params,omitnil,omitempty" name:"Params"`
}

type ModifyDBParametersRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID, for example: tdsql3-ow728lmc.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Parameter list, each element is a combination of Param and Value.
	Params []*DBParamValue `json:"Params,omitnil,omitempty" name:"Params"`
}

func (r *ModifyDBParametersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBParametersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Params")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDBParametersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBParametersResponseParams struct {
	// Task id.
	TaskID *int64 `json:"TaskID,omitnil,omitempty" name:"TaskID"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDBParametersResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDBParametersResponseParams `json:"Response"`
}

func (r *ModifyDBParametersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBParametersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBSBackupPolicyRequestParams struct {
	// <p>Backup policy</p>
	BackupPolicy *BackupPolicyModelInput `json:"BackupPolicy,omitnil,omitempty" name:"BackupPolicy"`

	// <p>Instance ID.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type ModifyDBSBackupPolicyRequest struct {
	*tchttp.BaseRequest
	
	// <p>Backup policy</p>
	BackupPolicy *BackupPolicyModelInput `json:"BackupPolicy,omitnil,omitempty" name:"BackupPolicy"`

	// <p>Instance ID.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *ModifyDBSBackupPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBSBackupPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BackupPolicy")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDBSBackupPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBSBackupPolicyResponseParams struct {
	// <p>Whether it is successful</p>
	IsSuccess *bool `json:"IsSuccess,omitnil,omitempty" name:"IsSuccess"`

	// <p>Message</p>
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDBSBackupPolicyResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDBSBackupPolicyResponseParams `json:"Response"`
}

func (r *ModifyDBSBackupPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBSBackupPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBSBackupSetCommentRequestParams struct {
	// <p>Instance ID.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Backup set ID. the value comes from the DescribeDBSBackupSets api response.</p>
	BackupSetId *int64 `json:"BackupSetId,omitnil,omitempty" name:"BackupSetId"`

	// <P>Backup notes.</p>
	NewBackupName *string `json:"NewBackupName,omitnil,omitempty" name:"NewBackupName"`
}

type ModifyDBSBackupSetCommentRequest struct {
	*tchttp.BaseRequest
	
	// <p>Instance ID.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Backup set ID. the value comes from the DescribeDBSBackupSets api response.</p>
	BackupSetId *int64 `json:"BackupSetId,omitnil,omitempty" name:"BackupSetId"`

	// <P>Backup notes.</p>
	NewBackupName *string `json:"NewBackupName,omitnil,omitempty" name:"NewBackupName"`
}

func (r *ModifyDBSBackupSetCommentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBSBackupSetCommentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BackupSetId")
	delete(f, "NewBackupName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDBSBackupSetCommentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBSBackupSetCommentResponseParams struct {
	// <P>Whether it is successful</p>.
	IsSuccess *bool `json:"IsSuccess,omitnil,omitempty" name:"IsSuccess"`

	// <P>Modify information</p>.
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDBSBackupSetCommentResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDBSBackupSetCommentResponseParams `json:"Response"`
}

func (r *ModifyDBSBackupSetCommentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBSBackupSetCommentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceNameRequestParams struct {
	// Instance id that needs to be modified.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Modified instance name. required length: 1-60. can contain chinese, english uppercase and lowercase letters, digits, -, _.
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`
}

type ModifyInstanceNameRequest struct {
	*tchttp.BaseRequest
	
	// Instance id that needs to be modified.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Modified instance name. required length: 1-60. can contain chinese, english uppercase and lowercase letters, digits, -, _.
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`
}

func (r *ModifyInstanceNameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceNameRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "InstanceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstanceNameRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceNameResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyInstanceNameResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstanceNameResponseParams `json:"Response"`
}

func (r *ModifyInstanceNameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceNetworkRequestParams struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// VpcId of the target VPC network
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Subnet ID of the target VPC network
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// VIP retention duration, in hours, value ranges from 0 to 168. 0 means immediate release with a one-minute delay. Not specified, default is 24 hr for VIP release.
	VipReleaseDelay *uint64 `json:"VipReleaseDelay,omitnil,omitempty" name:"VipReleaseDelay"`

	// Assign vip modification. Leave blank for a random vip.
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`
}

type ModifyInstanceNetworkRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// VpcId of the target VPC network
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Subnet ID of the target VPC network
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// VIP retention duration, in hours, value ranges from 0 to 168. 0 means immediate release with a one-minute delay. Not specified, default is 24 hr for VIP release.
	VipReleaseDelay *uint64 `json:"VipReleaseDelay,omitnil,omitempty" name:"VipReleaseDelay"`

	// Assign vip modification. Leave blank for a random vip.
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`
}

func (r *ModifyInstanceNetworkRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceNetworkRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "VipReleaseDelay")
	delete(f, "Vip")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstanceNetworkRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceNetworkResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyInstanceNetworkResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstanceNetworkResponseParams `json:"Response"`
}

func (r *ModifyInstanceNetworkResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceNetworkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceSSLStatusRequestParams struct {
	// <p>Instance ID.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Whether to enable SSL.</p>
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`
}

type ModifyInstanceSSLStatusRequest struct {
	*tchttp.BaseRequest
	
	// <p>Instance ID.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Whether to enable SSL.</p>
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`
}

func (r *ModifyInstanceSSLStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceSSLStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Enabled")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstanceSSLStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceSSLStatusResponseParams struct {
	// <p>Async process ID.</p>
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyInstanceSSLStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstanceSSLStatusResponseParams `json:"Response"`
}

func (r *ModifyInstanceSSLStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceSSLStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMaintenanceWindowRequestParams struct {
	// <p>Instance ID.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Ops window start time</p><p>Parameter format: hh:mm:ss</p>
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>Ops window duration</p><p>Value ranges from 1 to 3</p><p>Unit: hour</p>
	Duration *int64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// <p>Ops window date</p><p>Enumeration value:</p><ul><li>Monday: Monday</li><li>Tuesday: Tuesday</li><li>Wednesday: Wednesday</li><li>Thursday: Thursday</li><li>Friday: Friday</li><li>Saturday: Saturday</li><li>Sunday: Sunday</li></ul>
	WeekDays []*string `json:"WeekDays,omitnil,omitempty" name:"WeekDays"`
}

type ModifyMaintenanceWindowRequest struct {
	*tchttp.BaseRequest
	
	// <p>Instance ID.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Ops window start time</p><p>Parameter format: hh:mm:ss</p>
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>Ops window duration</p><p>Value ranges from 1 to 3</p><p>Unit: hour</p>
	Duration *int64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// <p>Ops window date</p><p>Enumeration value:</p><ul><li>Monday: Monday</li><li>Tuesday: Tuesday</li><li>Wednesday: Wednesday</li><li>Thursday: Thursday</li><li>Friday: Friday</li><li>Saturday: Saturday</li><li>Sunday: Sunday</li></ul>
	WeekDays []*string `json:"WeekDays,omitnil,omitempty" name:"WeekDays"`
}

func (r *ModifyMaintenanceWindowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMaintenanceWindowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "Duration")
	delete(f, "WeekDays")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyMaintenanceWindowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMaintenanceWindowResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyMaintenanceWindowResponse struct {
	*tchttp.BaseResponse
	Response *ModifyMaintenanceWindowResponseParams `json:"Response"`
}

func (r *ModifyMaintenanceWindowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMaintenanceWindowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUserPrivilegesRequestParams struct {
	// Instance ID, such as tdsql3-5baee8df.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Login username and host information
	Users []*User `json:"Users,omitnil,omitempty" name:"Users"`

	// Global permission
	GlobalPrivileges []*string `json:"GlobalPrivileges,omitnil,omitempty" name:"GlobalPrivileges"`

	// Database-level permission
	DatabasePrivileges []*DatabasePrivileges `json:"DatabasePrivileges,omitnil,omitempty" name:"DatabasePrivileges"`

	// Table-level permission
	TablePrivileges []*TablePrivileges `json:"TablePrivileges,omitnil,omitempty" name:"TablePrivileges"`
}

type ModifyUserPrivilegesRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID, such as tdsql3-5baee8df.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Login username and host information
	Users []*User `json:"Users,omitnil,omitempty" name:"Users"`

	// Global permission
	GlobalPrivileges []*string `json:"GlobalPrivileges,omitnil,omitempty" name:"GlobalPrivileges"`

	// Database-level permission
	DatabasePrivileges []*DatabasePrivileges `json:"DatabasePrivileges,omitnil,omitempty" name:"DatabasePrivileges"`

	// Table-level permission
	TablePrivileges []*TablePrivileges `json:"TablePrivileges,omitnil,omitempty" name:"TablePrivileges"`
}

func (r *ModifyUserPrivilegesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUserPrivilegesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Users")
	delete(f, "GlobalPrivileges")
	delete(f, "DatabasePrivileges")
	delete(f, "TablePrivileges")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyUserPrivilegesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUserPrivilegesResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyUserPrivilegesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyUserPrivilegesResponseParams `json:"Response"`
}

func (r *ModifyUserPrivilegesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUserPrivilegesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NodeInfo struct {
	// <p>Node IP information</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	IP *string `json:"IP,omitnil,omitempty" name:"IP"`

	// <p>Node types, such as sqlengine, tdstore, mc</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>Unique identifier of the node</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	NodeId *string `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// <p>Node port information</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// <p>Availability zone of the node</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// <p>Machine ip of the node</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// <p>Node log service information</p>
	BinlogInfo []*BinlogInfo `json:"BinlogInfo,omitnil,omitempty" name:"BinlogInfo"`

	// <p>Number of CPUs of the node</p>
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// <p>Node mem size</p>
	Mem *int64 `json:"Mem,omitnil,omitempty" name:"Mem"`

	// <p>Node disk size</p>
	DataDisk *int64 `json:"DataDisk,omitnil,omitempty" name:"DataDisk"`
}

type ParamConstraint struct {
	// Constraint type, for example enumeration, interval.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Lists the available options when the constraint type is enum.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Enum *string `json:"Enum,omitnil,omitempty" name:"Enum"`

	// Value range when the constraint type is section.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Range *ConstraintRange `json:"Range,omitnil,omitempty" name:"Range"`

	// Valid values when the constraint type is string.
	// Note: This field may return null, indicating that no valid values can be obtained.
	String *string `json:"String,omitnil,omitempty" name:"String"`
}

type ParamDesc struct {
	// Parameter name.
	Param *string `json:"Param,omitnil,omitempty" name:"Param"`

	// Current parameter value.
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// The configured value is the same as the value once the parameter takes effect.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SetValue *string `json:"SetValue,omitnil,omitempty" name:"SetValue"`

	// System default value.
	Default *string `json:"Default,omitnil,omitempty" name:"Default"`

	// Parameter limits.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Constraint *ParamConstraint `json:"Constraint,omitnil,omitempty" name:"Constraint"`

	// Whether a value has been set. false: not set. true: has set.
	HaveSetValue *bool `json:"HaveSetValue,omitnil,omitempty" name:"HaveSetValue"`

	// true: restart required.
	NeedRestart *bool `json:"NeedRestart,omitnil,omitempty" name:"NeedRestart"`

	// Parameter description.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type ResetUserPasswordInfo struct {
	// <p>Username.</p>
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// <p>host</p>
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// <p>plaintext password</p>
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// <p>Encryption password</p>
	EncryptedPassword *string `json:"EncryptedPassword,omitnil,omitempty" name:"EncryptedPassword"`
}

// Predefined struct for user
type ResetUserPasswordRequestParams struct {
	// Username.
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Host IP, IP range ending with % to denote permission for all IPs in the range
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// New password, required length 8-32, include at least two of English, digits and symbols.
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// Encryption password
	EncryptedPassword *string `json:"EncryptedPassword,omitnil,omitempty" name:"EncryptedPassword"`
}

type ResetUserPasswordRequest struct {
	*tchttp.BaseRequest
	
	// Username.
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Host IP, IP range ending with % to denote permission for all IPs in the range
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// New password, required length 8-32, include at least two of English, digits and symbols.
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// Encryption password
	EncryptedPassword *string `json:"EncryptedPassword,omitnil,omitempty" name:"EncryptedPassword"`
}

func (r *ResetUserPasswordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetUserPasswordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserName")
	delete(f, "InstanceId")
	delete(f, "Host")
	delete(f, "Password")
	delete(f, "EncryptedPassword")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetUserPasswordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetUserPasswordResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ResetUserPasswordResponse struct {
	*tchttp.BaseResponse
	Response *ResetUserPasswordResponseParams `json:"Response"`
}

func (r *ResetUserPasswordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetUserPasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetUsersPasswordRequestParams struct {
	// <p>Instance id</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Reset user password list</p>
	Users []*ResetUserPasswordInfo `json:"Users,omitnil,omitempty" name:"Users"`
}

type ResetUsersPasswordRequest struct {
	*tchttp.BaseRequest
	
	// <p>Instance id</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Reset user password list</p>
	Users []*ResetUserPasswordInfo `json:"Users,omitnil,omitempty" name:"Users"`
}

func (r *ResetUsersPasswordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetUsersPasswordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Users")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetUsersPasswordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetUsersPasswordResponseParams struct {
	// <p>Task ID.</p>
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ResetUsersPasswordResponse struct {
	*tchttp.BaseResponse
	Response *ResetUsersPasswordResponseParams `json:"Response"`
}

func (r *ResetUsersPasswordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetUsersPasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResourceTag struct {
	// Tag key
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// Tag value
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

// Predefined struct for user
type RestartDBInstancesRequestParams struct {
	// <p>Array of instance IDs that must be restarted</p>
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// <p>Restart time. Leave it blank to restart now.</p>
	RestartTime *string `json:"RestartTime,omitnil,omitempty" name:"RestartTime"`
}

type RestartDBInstancesRequest struct {
	*tchttp.BaseRequest
	
	// <p>Array of instance IDs that must be restarted</p>
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// <p>Restart time. Leave it blank to restart now.</p>
	RestartTime *string `json:"RestartTime,omitnil,omitempty" name:"RestartTime"`
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
	delete(f, "RestartTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RestartDBInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RestartDBInstancesResponseParams struct {
	// <p>Async task id.</p>
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RestartDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *RestartDBInstancesResponseParams `json:"Response"`
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

type SecurityGroup struct {
	// Project ID.
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// CreationTime, time format: yyyy-mm-dd hh:mm:ss.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Security group ID.
	SecurityGroupId *string `json:"SecurityGroupId,omitnil,omitempty" name:"SecurityGroupId"`

	// Security group name.
	SecurityGroupName *string `json:"SecurityGroupName,omitnil,omitempty" name:"SecurityGroupName"`

	// Security group remarks
	SecurityGroupRemark *string `json:"SecurityGroupRemark,omitnil,omitempty" name:"SecurityGroupRemark"`

	// Inbound rule.
	Inbound []*SecurityGroupBound `json:"Inbound,omitnil,omitempty" name:"Inbound"`

	// Outbound rule.
	Outbound []*SecurityGroupBound `json:"Outbound,omitnil,omitempty" name:"Outbound"`
}

type SecurityGroupBound struct {
	// Source IP or IP range, such as 192.168.0.0/16.
	CidrIp *string `json:"CidrIp,omitnil,omitempty" name:"CidrIp"`

	// Policy, ACCEPT or DROP.
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`

	// Port.
	PortRange *string `json:"PortRange,omitnil,omitempty" name:"PortRange"`

	// Network protocol, supports UDP, TCP.
	IpProtocol *string `json:"IpProtocol,omitnil,omitempty" name:"IpProtocol"`
}

type ServerlessCcu struct {
	// <p>ccu minimum value</p>
	MinCcu *int64 `json:"MinCcu,omitnil,omitempty" name:"MinCcu"`

	// <p>Maximum value of ccu</p>
	MaxCcu []*int64 `json:"MaxCcu,omitnil,omitempty" name:"MaxCcu"`
}

type SlowLogData struct {
	// <p>Sql execution time</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	Timestamp *string `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`

	// <p>Sql execution duration (seconds)</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	QueryTime *float64 `json:"QueryTime,omitnil,omitempty" name:"QueryTime"`

	// <p>Sql statement</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	SqlText *string `json:"SqlText,omitnil,omitempty" name:"SqlText"`

	// <p>Client IP address</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	UserHost *string `json:"UserHost,omitnil,omitempty" name:"UserHost"`

	// <p>Username.</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// <p>Database name.</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// <p>Lock duration (seconds)</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	LockTime *float64 `json:"LockTime,omitnil,omitempty" name:"LockTime"`

	// <p>Number of scanned rows</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	RowsExamined *uint64 `json:"RowsExamined,omitnil,omitempty" name:"RowsExamined"`

	// <p>Result set row count</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	RowsSent *uint64 `json:"RowsSent,omitnil,omitempty" name:"RowsSent"`

	// <p>Transaction ID</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	TransactionId *string `json:"TransactionId,omitnil,omitempty" name:"TransactionId"`

	// <p>rpc duration</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	RpcTime *float64 `json:"RpcTime,omitnil,omitempty" name:"RpcTime"`

	// <p>rpc duration for node interaction with storage</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	StorageRpcTime *float64 `json:"StorageRpcTime,omitnil,omitempty" name:"StorageRpcTime"`

	// <p>rpc retry latency</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	RpcRetryDelayTime *float64 `json:"RpcRetryDelayTime,omitnil,omitempty" name:"RpcRetryDelayTime"`

	// <p>node Name</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	NodeId *string `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// <p>rpc link tracing</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	RpcTrace *string `json:"RpcTrace,omitnil,omitempty" name:"RpcTrace"`

	// <p>TDStore lock duration</p><p>Unit: seconds</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	TDStoreLockTime *float64 `json:"TDStoreLockTime,omitnil,omitempty" name:"TDStoreLockTime"`

	// <p>Global ID</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	TraceId *string `json:"TraceId,omitnil,omitempty" name:"TraceId"`

	// <p>Execution plan</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	Explain []*Explain `json:"Explain,omitnil,omitempty" name:"Explain"`
}

type StorageNodeSpec struct {
	// <p>Specification code</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	SpecCode *string `json:"SpecCode,omitnil,omitempty" name:"SpecCode"`

	// <p>CPU cores of the storage node</p>
	StorageNodeCpu *int64 `json:"StorageNodeCpu,omitnil,omitempty" name:"StorageNodeCpu"`

	// <p>Storage node memory size</p>
	StorageNodeMem *int64 `json:"StorageNodeMem,omitnil,omitempty" name:"StorageNodeMem"`

	// <p>Maximum quantity of storage nodes</p>
	StorageNodeMaxNum *int64 `json:"StorageNodeMaxNum,omitnil,omitempty" name:"StorageNodeMaxNum"`

	// <p>Node disk size capacity limit</p>
	StorageNodeMaxDisk *int64 `json:"StorageNodeMaxDisk,omitnil,omitempty" name:"StorageNodeMaxDisk"`

	// <p>Minimum number of storage nodes</p>
	StorageNodeMinNum *int64 `json:"StorageNodeMinNum,omitnil,omitempty" name:"StorageNodeMinNum"`

	// <p>Node disk size minimum</p>
	StorageNodeMinDisk *int64 `json:"StorageNodeMinDisk,omitnil,omitempty" name:"StorageNodeMinDisk"`

	// <p>Disk Type, CLOUD_HSSD enhanced SSD, CLOUD_TCS local SSD disk</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	StorageType *string `json:"StorageType,omitnil,omitempty" name:"StorageType"`

	// <p>Default disk size of storage node for frontend display</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	StorageNodeDefaultDisk *int64 `json:"StorageNodeDefaultDisk,omitnil,omitempty" name:"StorageNodeDefaultDisk"`

	// <p>Specification support billing mode list</p>
	InstanceMode []*string `json:"InstanceMode,omitnil,omitempty" name:"InstanceMode"`

	// <p>Disk Type CLOUD_DISK: cloud disk LOCAL_DISK: local disk</p>
	DiskTypeCategory *string `json:"DiskTypeCategory,omitnil,omitempty" name:"DiskTypeCategory"`
}

type TablePrivileges struct {
	// DATABASE name
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// Table name
	Table *string `json:"Table,omitnil,omitempty" name:"Table"`

	// Permission list
	Privileges []*string `json:"Privileges,omitnil,omitempty" name:"Privileges"`
}

// Predefined struct for user
type UpgradeInstanceRequestParams struct {
	// <p>Instance ID.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Instance specification code</p>
	SpecCode *string `json:"SpecCode,omitnil,omitempty" name:"SpecCode"`

	// <p>Node disk capacity (unit: GB)</p>
	Disk *int64 `json:"Disk,omitnil,omitempty" name:"Disk"`

	// <p>CPU cores of the storage node</p>
	StorageNodeCpu *int64 `json:"StorageNodeCpu,omitnil,omitempty" name:"StorageNodeCpu"`

	// <p>Storage node memory size</p>
	StorageNodeMem *int64 `json:"StorageNodeMem,omitnil,omitempty" name:"StorageNodeMem"`

	// <p>Disk Type. Pass this parameter when the disk type needs to be modified.</p>
	StorageType *string `json:"StorageType,omitnil,omitempty" name:"StorageType"`
}

type UpgradeInstanceRequest struct {
	*tchttp.BaseRequest
	
	// <p>Instance ID.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Instance specification code</p>
	SpecCode *string `json:"SpecCode,omitnil,omitempty" name:"SpecCode"`

	// <p>Node disk capacity (unit: GB)</p>
	Disk *int64 `json:"Disk,omitnil,omitempty" name:"Disk"`

	// <p>CPU cores of the storage node</p>
	StorageNodeCpu *int64 `json:"StorageNodeCpu,omitnil,omitempty" name:"StorageNodeCpu"`

	// <p>Storage node memory size</p>
	StorageNodeMem *int64 `json:"StorageNodeMem,omitnil,omitempty" name:"StorageNodeMem"`

	// <p>Disk Type. Pass this parameter when the disk type needs to be modified.</p>
	StorageType *string `json:"StorageType,omitnil,omitempty" name:"StorageType"`
}

func (r *UpgradeInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SpecCode")
	delete(f, "Disk")
	delete(f, "StorageNodeCpu")
	delete(f, "StorageNodeMem")
	delete(f, "StorageType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpgradeInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeInstanceResponseParams struct {
	// <p>Task ID.</p>
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpgradeInstanceResponse struct {
	*tchttp.BaseResponse
	Response *UpgradeInstanceResponseParams `json:"Response"`
}

func (r *UpgradeInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type User struct {
	// Username.
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// host
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`
}

type UserInfo struct {
	// Username.
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// User description
	// Note: This field may return null, indicating that no valid values can be obtained.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Host IP, IP range ending with % to denote permission for all IPs in the range
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// Creation time.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Update time
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}