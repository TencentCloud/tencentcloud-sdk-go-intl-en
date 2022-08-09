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

package v20190725

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

// Predefined struct for user
type AssignProjectRequestParams struct {
	// List of instance IDs in the format of cmgo-p8vnipr5. It is the same as the instance ID displayed on the TencentDB Console page
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// Unique ID of an existing project (instead of a new project).
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`
}

type AssignProjectRequest struct {
	*tchttp.BaseRequest
	
	// List of instance IDs in the format of cmgo-p8vnipr5. It is the same as the instance ID displayed on the TencentDB Console page
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// Unique ID of an existing project (instead of a new project).
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *AssignProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssignProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AssignProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssignProjectResponseParams struct {
	// List of the returned async task IDs
	FlowIds []*uint64 `json:"FlowIds,omitempty" name:"FlowIds"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AssignProjectResponse struct {
	*tchttp.BaseResponse
	Response *AssignProjectResponseParams `json:"Response"`
}

func (r *AssignProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssignProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BackupDownloadTask struct {
	// Task creation time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Backup name
	BackupName *string `json:"BackupName,omitempty" name:"BackupName"`

	// Shard name
	ReplicaSetId *string `json:"ReplicaSetId,omitempty" name:"ReplicaSetId"`

	// Backup size in bytes
	BackupSize *int64 `json:"BackupSize,omitempty" name:"BackupSize"`

	// Task status. Valid values: `0` (waiting for execution), `1` (downloading), `2` (downloaded), `3` (download failed), `4` (waiting for retry)
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// Task progress in percentage
	Percent *int64 `json:"Percent,omitempty" name:"Percent"`

	// Task duration in seconds
	TimeSpend *int64 `json:"TimeSpend,omitempty" name:"TimeSpend"`

	// Backup download address
	Url *string `json:"Url,omitempty" name:"Url"`

	// Backup type of the backup file. Valid values: `0` (logical backup), `1` (physical backup)
	BackupMethod *int64 `json:"BackupMethod,omitempty" name:"BackupMethod"`

	// Backup description you set when starting a backup task
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	BackupDesc *string `json:"BackupDesc,omitempty" name:"BackupDesc"`
}

type BackupDownloadTaskStatus struct {
	// Shard name
	ReplicaSetId *string `json:"ReplicaSetId,omitempty" name:"ReplicaSetId"`

	// Task status. Valid values: `0` (waiting for execution), `1` (downloading), `2` (downloaded), `3` (download failed), `4` (waiting for retry)
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type BackupFile struct {
	// ID of the replica set/shard to which a backup file belongs
	ReplicateSetId *string `json:"ReplicateSetId,omitempty" name:"ReplicateSetId"`

	// Path to a backup file
	File *string `json:"File,omitempty" name:"File"`
}

type BackupInfo struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Backup mode. 0: automatic backup; 1: manual backup
	BackupType *uint64 `json:"BackupType,omitempty" name:"BackupType"`

	// Backup name
	BackupName *string `json:"BackupName,omitempty" name:"BackupName"`

	// Backup remarks
	// Note: This field may return null, indicating that no valid values can be obtained.
	BackupDesc *string `json:"BackupDesc,omitempty" name:"BackupDesc"`

	// Backup file size in KB
	// Note: This field may return null, indicating that no valid values can be obtained.
	BackupSize *uint64 `json:"BackupSize,omitempty" name:"BackupSize"`

	// Backup start time
	// Note: This field may return null, indicating that no valid values can be obtained.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Backup end time
	// Note: This field may return null, indicating that no valid values can be obtained.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Backup status. 1: backing up; 2: backed up successful
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// Backup method. 0: logical backup; 1: physical backup
	BackupMethod *uint64 `json:"BackupMethod,omitempty" name:"BackupMethod"`
}

type ClientConnection struct {
	// Client IP of a connection
	IP *string `json:"IP,omitempty" name:"IP"`

	// Number of connections corresponding to a client IP
	Count *uint64 `json:"Count,omitempty" name:"Count"`

	// Whether it is the Tencent Cloud IP for automated testing
	InternalService *bool `json:"InternalService,omitempty" name:"InternalService"`
}

// Predefined struct for user
type CreateBackupDBInstanceRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Valid values: 0 (logical backup), 1 (physical backup)
	BackupMethod *int64 `json:"BackupMethod,omitempty" name:"BackupMethod"`

	// Backup remarks
	BackupRemark *string `json:"BackupRemark,omitempty" name:"BackupRemark"`
}

type CreateBackupDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Valid values: 0 (logical backup), 1 (physical backup)
	BackupMethod *int64 `json:"BackupMethod,omitempty" name:"BackupMethod"`

	// Backup remarks
	BackupRemark *string `json:"BackupRemark,omitempty" name:"BackupRemark"`
}

func (r *CreateBackupDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBackupDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BackupMethod")
	delete(f, "BackupRemark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBackupDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBackupDBInstanceResponseParams struct {
	// The status of the queried backup process.
	AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateBackupDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *CreateBackupDBInstanceResponseParams `json:"Response"`
}

func (r *CreateBackupDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBackupDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBackupDownloadTaskRequestParams struct {
	// Instance ID in the format of "cmgo-p8vnipr5", which is the same as the instance ID displayed in the TencentDB console.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// The name of the backup file to be downloaded, which can be obtained by the `DescribeDBBackups` API.
	BackupName *string `json:"BackupName,omitempty" name:"BackupName"`

	// Specify the node name of a replica set instance or the shard name list of a sharded cluster instance. Only backups of the specified node or shards will be downloaded.
	// Suppose you have a replica set instance (ID: cmgo-p8vnipr5), you can use the sample code `BackupSets.0=cmgo-p8vnipr5_0` to download the full backup. For a replica set instance, the parameter value must be in the format of "instance ID_0".
	// Suppose you have a sharded cluster instance (ID: cmgo-p8vnipr5), you can use the sample code `BackupSets.0=cmgo-p8vnipr5_0&BackupSets.1=cmgo-p8vnipr5_1` to download the backup data of shard 0 and shard 1. To download the full backup, please specify all shard names.
	BackupSets []*ReplicaSetInfo `json:"BackupSets,omitempty" name:"BackupSets"`
}

type CreateBackupDownloadTaskRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of "cmgo-p8vnipr5", which is the same as the instance ID displayed in the TencentDB console.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// The name of the backup file to be downloaded, which can be obtained by the `DescribeDBBackups` API.
	BackupName *string `json:"BackupName,omitempty" name:"BackupName"`

	// Specify the node name of a replica set instance or the shard name list of a sharded cluster instance. Only backups of the specified node or shards will be downloaded.
	// Suppose you have a replica set instance (ID: cmgo-p8vnipr5), you can use the sample code `BackupSets.0=cmgo-p8vnipr5_0` to download the full backup. For a replica set instance, the parameter value must be in the format of "instance ID_0".
	// Suppose you have a sharded cluster instance (ID: cmgo-p8vnipr5), you can use the sample code `BackupSets.0=cmgo-p8vnipr5_0&BackupSets.1=cmgo-p8vnipr5_1` to download the backup data of shard 0 and shard 1. To download the full backup, please specify all shard names.
	BackupSets []*ReplicaSetInfo `json:"BackupSets,omitempty" name:"BackupSets"`
}

func (r *CreateBackupDownloadTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBackupDownloadTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BackupName")
	delete(f, "BackupSets")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBackupDownloadTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBackupDownloadTaskResponseParams struct {
	// Download task status
	Tasks []*BackupDownloadTaskStatus `json:"Tasks,omitempty" name:"Tasks"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateBackupDownloadTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateBackupDownloadTaskResponseParams `json:"Response"`
}

func (r *CreateBackupDownloadTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBackupDownloadTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDBInstanceHourRequestParams struct {
	// Instance memory size in GB
	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`

	// Instance disk size in GB
	Volume *uint64 `json:"Volume,omitempty" name:"Volume"`

	// Number of replica sets. When a replica set instance is created, this parameter must be set to 1. When a sharding instance is created, please see the parameters returned by the DescribeSpecInfo API
	ReplicateSetNum *uint64 `json:"ReplicateSetNum,omitempty" name:"ReplicateSetNum"`

	// The number of nodes in each replica set. The value range is subject to the response parameter of the `DescribeSpecInfo` API.
	NodeNum *uint64 `json:"NodeNum,omitempty" name:"NodeNum"`

	// Version number. For the specific purchasable versions supported, please see the return result of the `DescribeSpecInfo` API. The correspondences between parameters and versions are as follows: MONGO_3_WT: MongoDB 3.2 WiredTiger Edition; MONGO_3_ROCKS: MongoDB 3.2 RocksDB Edition; MONGO_36_WT: MongoDB 3.6 WiredTiger Edition; MONGO_40_WT: MongoDB 4.0 WiredTiger Edition; MONGO_42_WT: MongoDB 4.2 WiredTiger Edition.
	MongoVersion *string `json:"MongoVersion,omitempty" name:"MongoVersion"`

	// Server type. HIO: high IO; HIO10G: 10-Gigabit high IO
	MachineCode *string `json:"MachineCode,omitempty" name:"MachineCode"`

	// Number of instances. Minimum value: 1. Maximum value: 10
	GoodsNum *uint64 `json:"GoodsNum,omitempty" name:"GoodsNum"`

	// AZ in the format of ap-guangzhou-2. If multi-AZ deployment is enabled, this parameter refers to the primary AZ and must be one of the values of `AvailabilityZoneList`.
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// Instance type. REPLSET: replica set; SHARD: sharding cluster
	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`

	// VPC ID. If this parameter is not set, the basic network will be selected by default
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// VPC subnet ID. If VpcId is set, then SubnetId will be required
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// Instance password, which must contain 8 to 16 characters and comprise at least two of the following types: letters, digits, and symbols (!@#%^*()). If it is left empty, the password is in the format of "instance ID+@+root account UIN". For example, if the instance ID is "cmgo-higv73ed" and the root account UIN "100000001", the instance password will be "cmgo-higv73ed@100000001".
	Password *string `json:"Password,omitempty" name:"Password"`

	// Project ID. If this parameter is not set, the default project will be used
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Instance tag information
	Tags []*TagInfo `json:"Tags,omitempty" name:"Tags"`

	// Instance type. Valid values: `1` (primary instance), `2` (temp instance), `3` (read-only instance), `4` (disaster recovery instance), `5` (cloned instance).
	Clone *int64 `json:"Clone,omitempty" name:"Clone"`

	// Parent instance ID. It is required if the `Clone` is 3 or 4.
	Father *string `json:"Father,omitempty" name:"Father"`

	// Security group.
	SecurityGroup []*string `json:"SecurityGroup,omitempty" name:"SecurityGroup"`

	// The point in time to which the cloned instance will be rolled back. This parameter is required for a cloned instance. The point in time in the format of 2021-08-13 16:30:00 must be within the last seven days.
	RestoreTime *string `json:"RestoreTime,omitempty" name:"RestoreTime"`

	// Instance name, which can contain up to 60 letters, digits, or symbols (_-).
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// AZ list when multi-AZ deployment is enabled. For the specific purchasable versions which support multi-AZ deployment, please see the return result of the `DescribeSpecInfo` API. Notes: 1. Nodes of a multi-AZ instance must be deployed across three AZs. 2. To ensure a successful cross-AZ switch, you should not deploy most of the nodes to the same AZ. (For example, a three-node sharded cluster instance does not support deploying two or more nodes in the same AZ.) 3. MongoDB 4.2 and later versions do not support multi-AZ deployment. 4. Read-Only and disaster recovery instances do not support multi-AZ deployment. 5. Instances in the classic network do not support multi-AZ deployment.
	AvailabilityZoneList []*string `json:"AvailabilityZoneList,omitempty" name:"AvailabilityZoneList"`

	// The number of mongos CPUs, which is required for a sharded cluster instance of MongoDB 4.2 WiredTiger. For the specific purchasable versions supported, please see the return result of the `DescribeSpecInfo` API.
	MongosCpu *uint64 `json:"MongosCpu,omitempty" name:"MongosCpu"`

	// The size of mongos memory, which is required for a sharded cluster instance of MongoDB 4.2 WiredTiger. For the specific purchasable versions supported, please see the return result of the `DescribeSpecInfo` API.
	MongosMemory *uint64 `json:"MongosMemory,omitempty" name:"MongosMemory"`

	// The number of mongos routers, which is required for a sharded cluster instance of MongoDB 4.2 WiredTiger. For the specific purchasable versions supported, please see the return result of the `DescribeSpecInfo` API. Note: please purchase 3-32 mongos routers for high availability.
	MongosNodeNum *uint64 `json:"MongosNodeNum,omitempty" name:"MongosNodeNum"`
}

type CreateDBInstanceHourRequest struct {
	*tchttp.BaseRequest
	
	// Instance memory size in GB
	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`

	// Instance disk size in GB
	Volume *uint64 `json:"Volume,omitempty" name:"Volume"`

	// Number of replica sets. When a replica set instance is created, this parameter must be set to 1. When a sharding instance is created, please see the parameters returned by the DescribeSpecInfo API
	ReplicateSetNum *uint64 `json:"ReplicateSetNum,omitempty" name:"ReplicateSetNum"`

	// The number of nodes in each replica set. The value range is subject to the response parameter of the `DescribeSpecInfo` API.
	NodeNum *uint64 `json:"NodeNum,omitempty" name:"NodeNum"`

	// Version number. For the specific purchasable versions supported, please see the return result of the `DescribeSpecInfo` API. The correspondences between parameters and versions are as follows: MONGO_3_WT: MongoDB 3.2 WiredTiger Edition; MONGO_3_ROCKS: MongoDB 3.2 RocksDB Edition; MONGO_36_WT: MongoDB 3.6 WiredTiger Edition; MONGO_40_WT: MongoDB 4.0 WiredTiger Edition; MONGO_42_WT: MongoDB 4.2 WiredTiger Edition.
	MongoVersion *string `json:"MongoVersion,omitempty" name:"MongoVersion"`

	// Server type. HIO: high IO; HIO10G: 10-Gigabit high IO
	MachineCode *string `json:"MachineCode,omitempty" name:"MachineCode"`

	// Number of instances. Minimum value: 1. Maximum value: 10
	GoodsNum *uint64 `json:"GoodsNum,omitempty" name:"GoodsNum"`

	// AZ in the format of ap-guangzhou-2. If multi-AZ deployment is enabled, this parameter refers to the primary AZ and must be one of the values of `AvailabilityZoneList`.
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// Instance type. REPLSET: replica set; SHARD: sharding cluster
	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`

	// VPC ID. If this parameter is not set, the basic network will be selected by default
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// VPC subnet ID. If VpcId is set, then SubnetId will be required
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// Instance password, which must contain 8 to 16 characters and comprise at least two of the following types: letters, digits, and symbols (!@#%^*()). If it is left empty, the password is in the format of "instance ID+@+root account UIN". For example, if the instance ID is "cmgo-higv73ed" and the root account UIN "100000001", the instance password will be "cmgo-higv73ed@100000001".
	Password *string `json:"Password,omitempty" name:"Password"`

	// Project ID. If this parameter is not set, the default project will be used
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Instance tag information
	Tags []*TagInfo `json:"Tags,omitempty" name:"Tags"`

	// Instance type. Valid values: `1` (primary instance), `2` (temp instance), `3` (read-only instance), `4` (disaster recovery instance), `5` (cloned instance).
	Clone *int64 `json:"Clone,omitempty" name:"Clone"`

	// Parent instance ID. It is required if the `Clone` is 3 or 4.
	Father *string `json:"Father,omitempty" name:"Father"`

	// Security group.
	SecurityGroup []*string `json:"SecurityGroup,omitempty" name:"SecurityGroup"`

	// The point in time to which the cloned instance will be rolled back. This parameter is required for a cloned instance. The point in time in the format of 2021-08-13 16:30:00 must be within the last seven days.
	RestoreTime *string `json:"RestoreTime,omitempty" name:"RestoreTime"`

	// Instance name, which can contain up to 60 letters, digits, or symbols (_-).
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// AZ list when multi-AZ deployment is enabled. For the specific purchasable versions which support multi-AZ deployment, please see the return result of the `DescribeSpecInfo` API. Notes: 1. Nodes of a multi-AZ instance must be deployed across three AZs. 2. To ensure a successful cross-AZ switch, you should not deploy most of the nodes to the same AZ. (For example, a three-node sharded cluster instance does not support deploying two or more nodes in the same AZ.) 3. MongoDB 4.2 and later versions do not support multi-AZ deployment. 4. Read-Only and disaster recovery instances do not support multi-AZ deployment. 5. Instances in the classic network do not support multi-AZ deployment.
	AvailabilityZoneList []*string `json:"AvailabilityZoneList,omitempty" name:"AvailabilityZoneList"`

	// The number of mongos CPUs, which is required for a sharded cluster instance of MongoDB 4.2 WiredTiger. For the specific purchasable versions supported, please see the return result of the `DescribeSpecInfo` API.
	MongosCpu *uint64 `json:"MongosCpu,omitempty" name:"MongosCpu"`

	// The size of mongos memory, which is required for a sharded cluster instance of MongoDB 4.2 WiredTiger. For the specific purchasable versions supported, please see the return result of the `DescribeSpecInfo` API.
	MongosMemory *uint64 `json:"MongosMemory,omitempty" name:"MongosMemory"`

	// The number of mongos routers, which is required for a sharded cluster instance of MongoDB 4.2 WiredTiger. For the specific purchasable versions supported, please see the return result of the `DescribeSpecInfo` API. Note: please purchase 3-32 mongos routers for high availability.
	MongosNodeNum *uint64 `json:"MongosNodeNum,omitempty" name:"MongosNodeNum"`
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
	delete(f, "Memory")
	delete(f, "Volume")
	delete(f, "ReplicateSetNum")
	delete(f, "NodeNum")
	delete(f, "MongoVersion")
	delete(f, "MachineCode")
	delete(f, "GoodsNum")
	delete(f, "Zone")
	delete(f, "ClusterType")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "Password")
	delete(f, "ProjectId")
	delete(f, "Tags")
	delete(f, "Clone")
	delete(f, "Father")
	delete(f, "SecurityGroup")
	delete(f, "RestoreTime")
	delete(f, "InstanceName")
	delete(f, "AvailabilityZoneList")
	delete(f, "MongosCpu")
	delete(f, "MongosMemory")
	delete(f, "MongosNodeNum")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDBInstanceHourRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDBInstanceHourResponseParams struct {
	// Order ID
	DealId *string `json:"DealId,omitempty" name:"DealId"`

	// List of IDs of created instances
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateDBInstanceHourResponse struct {
	*tchttp.BaseResponse
	Response *CreateDBInstanceHourResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateDBInstanceRequestParams struct {
	// The number of nodes in each replica set. The value range is subject to the response parameter of the `DescribeSpecInfo` API.
	NodeNum *uint64 `json:"NodeNum,omitempty" name:"NodeNum"`

	// Instance memory size in GB.
	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`

	// Instance disk size in GB.
	Volume *uint64 `json:"Volume,omitempty" name:"Volume"`

	// Version number. For the specific purchasable versions supported, please see the return result of the `DescribeSpecInfo` API. The correspondences between parameters and versions are as follows: MONGO_3_WT: MongoDB 3.2 WiredTiger Edition; MONGO_3_ROCKS: MongoDB 3.2 RocksDB Edition; MONGO_36_WT: MongoDB 3.6 WiredTiger Edition; MONGO_40_WT: MongoDB 4.0 WiredTiger Edition; MONGO_42_WT: MongoDB 4.2 WiredTiger Edition.
	MongoVersion *string `json:"MongoVersion,omitempty" name:"MongoVersion"`

	// Number of instances. Minimum value: 1. Maximum value: 10.
	GoodsNum *uint64 `json:"GoodsNum,omitempty" name:"GoodsNum"`

	// AZ in the format of ap-guangzhou-2. If multi-AZ deployment is enabled, this parameter refers to the primary AZ and must be one of the values of `AvailabilityZoneList`.
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// Instance validity period in months. Valid values: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36.
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// Server type. Valid values: HIO (high IO), HIO10G (10-gigabit high IO), STDS5 (standard).
	MachineCode *string `json:"MachineCode,omitempty" name:"MachineCode"`

	// Instance type. Valid values: REPLSET (replica set), SHARD (sharded cluster), STANDALONE (single-node).
	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`

	// Number of replica sets. To create a replica set instance, set this parameter to 1; to create a shard instance, see the parameters returned by the `DescribeSpecInfo` API; to create a single-node instance, set this parameter to 0.
	ReplicateSetNum *uint64 `json:"ReplicateSetNum,omitempty" name:"ReplicateSetNum"`

	// Project ID. If this parameter is not set, the default project will be used.
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// VPC ID. If this parameter is not set, the classic network will be used. Please use the `DescribeVpcs` API to query the VPC list.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// VPC subnet ID. If `UniqVpcId` is set, then `UniqSubnetId` will be required. Please use the `DescribeSubnets` API to query the subnet list.
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// Instance password, which must contain 8 to 16 characters and comprise at least two of the following types: letters, digits, and symbols (!@#%^*()). If it is left empty, the password is in the format of "instance ID+@+root account UIN". For example, if the instance ID is "cmgo-higv73ed" and the root account UIN "100000001", the instance password will be "cmgo-higv73ed@100000001".
	Password *string `json:"Password,omitempty" name:"Password"`

	// Instance tag information.
	Tags []*TagInfo `json:"Tags,omitempty" name:"Tags"`

	// Auto-renewal flag. Valid values: 0 (auto-renewal not enabled), 1 (auto-renewal enabled). Default value: 0.
	AutoRenewFlag *uint64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`

	// Whether to automatically use a voucher. Valid values: 1 (yes), 0 (no). Default value: 0.
	AutoVoucher *uint64 `json:"AutoVoucher,omitempty" name:"AutoVoucher"`

	// Instance type. Valid values: `1` (primary instance), `2` (temp instance), `3` (read-only instance), `4` (disaster recovery instance), `5` (cloned instance).
	Clone *int64 `json:"Clone,omitempty" name:"Clone"`

	// Primary instance ID. It is required for read-only, disaster recovery, and cloned instances.
	Father *string `json:"Father,omitempty" name:"Father"`

	// Security group.
	SecurityGroup []*string `json:"SecurityGroup,omitempty" name:"SecurityGroup"`

	// The point in time to which the cloned instance will be rolled back. This parameter is required for a cloned instance. The point in time in the format of 2021-08-13 16:30:00 must be within the last seven days.
	RestoreTime *string `json:"RestoreTime,omitempty" name:"RestoreTime"`

	// Instance name, which can contain up to 60 letters, digits, or symbols (_-).
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// AZ list when multi-AZ deployment is enabled. For the specific purchasable versions which support multi-AZ deployment, please see the return result of the `DescribeSpecInfo` API. Notes: 1. Nodes of a multi-AZ instance must be deployed across three AZs. 2. To ensure a successful cross-AZ switch, you should not deploy most of the nodes to the same AZ. (For example, a three-node sharded cluster instance does not support deploying two or more nodes in the same AZ.) 3. MongoDB 4.2 and later versions do not support multi-AZ deployment. 4. Read-Only and disaster recovery instances do not support multi-AZ deployment. 5. Instances in the classic network do not support multi-AZ deployment.
	AvailabilityZoneList []*string `json:"AvailabilityZoneList,omitempty" name:"AvailabilityZoneList"`

	// The number of mongos CPUs, which is required for a sharded cluster instance of MongoDB 4.2 WiredTiger. For the specific purchasable versions supported, please see the return result of the `DescribeSpecInfo` API.
	MongosCpu *uint64 `json:"MongosCpu,omitempty" name:"MongosCpu"`

	// The size of mongos memory, which is required for a sharded cluster instance of MongoDB 4.2 WiredTiger. For the specific purchasable versions supported, please see the return result of the `DescribeSpecInfo` API.
	MongosMemory *uint64 `json:"MongosMemory,omitempty" name:"MongosMemory"`

	// The number of mongos routers, which is required for a sharded cluster instance of MongoDB 4.2 WiredTiger. For the specific purchasable versions supported, please see the return result of the `DescribeSpecInfo` API. Note: please purchase 3-32 mongos routers for high availability.
	MongosNodeNum *uint64 `json:"MongosNodeNum,omitempty" name:"MongosNodeNum"`
}

type CreateDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// The number of nodes in each replica set. The value range is subject to the response parameter of the `DescribeSpecInfo` API.
	NodeNum *uint64 `json:"NodeNum,omitempty" name:"NodeNum"`

	// Instance memory size in GB.
	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`

	// Instance disk size in GB.
	Volume *uint64 `json:"Volume,omitempty" name:"Volume"`

	// Version number. For the specific purchasable versions supported, please see the return result of the `DescribeSpecInfo` API. The correspondences between parameters and versions are as follows: MONGO_3_WT: MongoDB 3.2 WiredTiger Edition; MONGO_3_ROCKS: MongoDB 3.2 RocksDB Edition; MONGO_36_WT: MongoDB 3.6 WiredTiger Edition; MONGO_40_WT: MongoDB 4.0 WiredTiger Edition; MONGO_42_WT: MongoDB 4.2 WiredTiger Edition.
	MongoVersion *string `json:"MongoVersion,omitempty" name:"MongoVersion"`

	// Number of instances. Minimum value: 1. Maximum value: 10.
	GoodsNum *uint64 `json:"GoodsNum,omitempty" name:"GoodsNum"`

	// AZ in the format of ap-guangzhou-2. If multi-AZ deployment is enabled, this parameter refers to the primary AZ and must be one of the values of `AvailabilityZoneList`.
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// Instance validity period in months. Valid values: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36.
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// Server type. Valid values: HIO (high IO), HIO10G (10-gigabit high IO), STDS5 (standard).
	MachineCode *string `json:"MachineCode,omitempty" name:"MachineCode"`

	// Instance type. Valid values: REPLSET (replica set), SHARD (sharded cluster), STANDALONE (single-node).
	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`

	// Number of replica sets. To create a replica set instance, set this parameter to 1; to create a shard instance, see the parameters returned by the `DescribeSpecInfo` API; to create a single-node instance, set this parameter to 0.
	ReplicateSetNum *uint64 `json:"ReplicateSetNum,omitempty" name:"ReplicateSetNum"`

	// Project ID. If this parameter is not set, the default project will be used.
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// VPC ID. If this parameter is not set, the classic network will be used. Please use the `DescribeVpcs` API to query the VPC list.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// VPC subnet ID. If `UniqVpcId` is set, then `UniqSubnetId` will be required. Please use the `DescribeSubnets` API to query the subnet list.
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// Instance password, which must contain 8 to 16 characters and comprise at least two of the following types: letters, digits, and symbols (!@#%^*()). If it is left empty, the password is in the format of "instance ID+@+root account UIN". For example, if the instance ID is "cmgo-higv73ed" and the root account UIN "100000001", the instance password will be "cmgo-higv73ed@100000001".
	Password *string `json:"Password,omitempty" name:"Password"`

	// Instance tag information.
	Tags []*TagInfo `json:"Tags,omitempty" name:"Tags"`

	// Auto-renewal flag. Valid values: 0 (auto-renewal not enabled), 1 (auto-renewal enabled). Default value: 0.
	AutoRenewFlag *uint64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`

	// Whether to automatically use a voucher. Valid values: 1 (yes), 0 (no). Default value: 0.
	AutoVoucher *uint64 `json:"AutoVoucher,omitempty" name:"AutoVoucher"`

	// Instance type. Valid values: `1` (primary instance), `2` (temp instance), `3` (read-only instance), `4` (disaster recovery instance), `5` (cloned instance).
	Clone *int64 `json:"Clone,omitempty" name:"Clone"`

	// Primary instance ID. It is required for read-only, disaster recovery, and cloned instances.
	Father *string `json:"Father,omitempty" name:"Father"`

	// Security group.
	SecurityGroup []*string `json:"SecurityGroup,omitempty" name:"SecurityGroup"`

	// The point in time to which the cloned instance will be rolled back. This parameter is required for a cloned instance. The point in time in the format of 2021-08-13 16:30:00 must be within the last seven days.
	RestoreTime *string `json:"RestoreTime,omitempty" name:"RestoreTime"`

	// Instance name, which can contain up to 60 letters, digits, or symbols (_-).
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// AZ list when multi-AZ deployment is enabled. For the specific purchasable versions which support multi-AZ deployment, please see the return result of the `DescribeSpecInfo` API. Notes: 1. Nodes of a multi-AZ instance must be deployed across three AZs. 2. To ensure a successful cross-AZ switch, you should not deploy most of the nodes to the same AZ. (For example, a three-node sharded cluster instance does not support deploying two or more nodes in the same AZ.) 3. MongoDB 4.2 and later versions do not support multi-AZ deployment. 4. Read-Only and disaster recovery instances do not support multi-AZ deployment. 5. Instances in the classic network do not support multi-AZ deployment.
	AvailabilityZoneList []*string `json:"AvailabilityZoneList,omitempty" name:"AvailabilityZoneList"`

	// The number of mongos CPUs, which is required for a sharded cluster instance of MongoDB 4.2 WiredTiger. For the specific purchasable versions supported, please see the return result of the `DescribeSpecInfo` API.
	MongosCpu *uint64 `json:"MongosCpu,omitempty" name:"MongosCpu"`

	// The size of mongos memory, which is required for a sharded cluster instance of MongoDB 4.2 WiredTiger. For the specific purchasable versions supported, please see the return result of the `DescribeSpecInfo` API.
	MongosMemory *uint64 `json:"MongosMemory,omitempty" name:"MongosMemory"`

	// The number of mongos routers, which is required for a sharded cluster instance of MongoDB 4.2 WiredTiger. For the specific purchasable versions supported, please see the return result of the `DescribeSpecInfo` API. Note: please purchase 3-32 mongos routers for high availability.
	MongosNodeNum *uint64 `json:"MongosNodeNum,omitempty" name:"MongosNodeNum"`
}

func (r *CreateDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NodeNum")
	delete(f, "Memory")
	delete(f, "Volume")
	delete(f, "MongoVersion")
	delete(f, "GoodsNum")
	delete(f, "Zone")
	delete(f, "Period")
	delete(f, "MachineCode")
	delete(f, "ClusterType")
	delete(f, "ReplicateSetNum")
	delete(f, "ProjectId")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "Password")
	delete(f, "Tags")
	delete(f, "AutoRenewFlag")
	delete(f, "AutoVoucher")
	delete(f, "Clone")
	delete(f, "Father")
	delete(f, "SecurityGroup")
	delete(f, "RestoreTime")
	delete(f, "InstanceName")
	delete(f, "AvailabilityZoneList")
	delete(f, "MongosCpu")
	delete(f, "MongosMemory")
	delete(f, "MongosNodeNum")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDBInstanceResponseParams struct {
	// Order ID.
	DealId *string `json:"DealId,omitempty" name:"DealId"`

	// List of IDs of created instances.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *CreateDBInstanceResponseParams `json:"Response"`
}

func (r *CreateDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DBInstanceInfo struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Region information
	Region *string `json:"Region,omitempty" name:"Region"`
}

type DBInstancePrice struct {
	// Unit price.
	// Note: this field may return null, indicating that no valid values can be obtained.
	UnitPrice *float64 `json:"UnitPrice,omitempty" name:"UnitPrice"`

	// Original price.
	OriginalPrice *float64 `json:"OriginalPrice,omitempty" name:"OriginalPrice"`

	// Discounted price.
	DiscountPrice *float64 `json:"DiscountPrice,omitempty" name:"DiscountPrice"`
}

// Predefined struct for user
type DescribeAsyncRequestInfoRequestParams struct {
	// Async task ID, which is returned by APIs related to async tasks, such as `CreateBackupDBInstance`.
	AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`
}

type DescribeAsyncRequestInfoRequest struct {
	*tchttp.BaseRequest
	
	// Async task ID, which is returned by APIs related to async tasks, such as `CreateBackupDBInstance`.
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

// Predefined struct for user
type DescribeAsyncRequestInfoResponseParams struct {
	// Status. Valid values: `initial` (initializing), `running`, `paused` (paused due to failure), `undoed` (rolled back due to failure), `failed` (ended due to failure), `success`
	Status *string `json:"Status,omitempty" name:"Status"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAsyncRequestInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAsyncRequestInfoResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeBackupAccessRequestParams struct {
	// Instance ID in the format of cmgo-p8vnipr5. It is the same as the instance ID displayed on the TencentDB Console page
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Name of the backup file for which to get the download permission
	BackupName *string `json:"BackupName,omitempty" name:"BackupName"`
}

type DescribeBackupAccessRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of cmgo-p8vnipr5. It is the same as the instance ID displayed on the TencentDB Console page
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Name of the backup file for which to get the download permission
	BackupName *string `json:"BackupName,omitempty" name:"BackupName"`
}

func (r *DescribeBackupAccessRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupAccessRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BackupName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackupAccessRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupAccessResponseParams struct {
	// Instance region
	Region *string `json:"Region,omitempty" name:"Region"`

	// The bucket where a backup file is located
	Bucket *string `json:"Bucket,omitempty" name:"Bucket"`

	// Storage information of a backup file
	Files []*BackupFile `json:"Files,omitempty" name:"Files"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBackupAccessResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBackupAccessResponseParams `json:"Response"`
}

func (r *DescribeBackupAccessResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupAccessResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupDownloadTaskRequestParams struct {
	// Instance ID in the format of "cmgo-p8vnipr5", which is the same as the instance ID displayed in the TencentDB console
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// The name of a backup file with download tasks to be queried
	BackupName *string `json:"BackupName,omitempty" name:"BackupName"`

	// The start time of the query period. Tasks whose start time and end time fall within the query period will be queried. If it is left empty, the start time can be any time earlier than the end time.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// The end time of the query period. Tasks will be queried if their start and end times fall within the query period. If it is left empty, the end time can be any time later than the start time.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// The maximum number of results returned per page. Value range: 1-100. Default value: `20`.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Offset for pagination. Default value: `0`.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// The field used to sort the results. Valid values: `createTime` (default), `finishTime`.
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// Sort order. Valid values: `asc`, `desc` (default).
	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`

	// The status of the tasks to be queried. Valid values: `0` (waiting for execution), `1` (downloading), `2` (downloaded), `3` (download failed), `4` (waiting for retry). If it is left empty, tasks in any status will be returned.
	Status []*int64 `json:"Status,omitempty" name:"Status"`
}

type DescribeBackupDownloadTaskRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of "cmgo-p8vnipr5", which is the same as the instance ID displayed in the TencentDB console
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// The name of a backup file with download tasks to be queried
	BackupName *string `json:"BackupName,omitempty" name:"BackupName"`

	// The start time of the query period. Tasks whose start time and end time fall within the query period will be queried. If it is left empty, the start time can be any time earlier than the end time.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// The end time of the query period. Tasks will be queried if their start and end times fall within the query period. If it is left empty, the end time can be any time later than the start time.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// The maximum number of results returned per page. Value range: 1-100. Default value: `20`.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Offset for pagination. Default value: `0`.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// The field used to sort the results. Valid values: `createTime` (default), `finishTime`.
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// Sort order. Valid values: `asc`, `desc` (default).
	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`

	// The status of the tasks to be queried. Valid values: `0` (waiting for execution), `1` (downloading), `2` (downloaded), `3` (download failed), `4` (waiting for retry). If it is left empty, tasks in any status will be returned.
	Status []*int64 `json:"Status,omitempty" name:"Status"`
}

func (r *DescribeBackupDownloadTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupDownloadTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BackupName")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackupDownloadTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupDownloadTaskResponseParams struct {
	// Total number of results
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// The list of download tasks
	Tasks []*BackupDownloadTask `json:"Tasks,omitempty" name:"Tasks"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBackupDownloadTaskResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBackupDownloadTaskResponseParams `json:"Response"`
}

func (r *DescribeBackupDownloadTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupDownloadTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClientConnectionsRequestParams struct {
	// Instance ID in the format of cmgo-p8vnipr5. It is the same as the instance ID displayed on the TencentDB Console page
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Number of results to be returned for a single request. Value range: 1-1,000. Default value: 1,000
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeClientConnectionsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of cmgo-p8vnipr5. It is the same as the instance ID displayed on the TencentDB Console page
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Number of results to be returned for a single request. Value range: 1-1,000. Default value: 1,000
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeClientConnectionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClientConnectionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClientConnectionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClientConnectionsResponseParams struct {
	// Client connection information, including client IP and number of connections
	Clients []*ClientConnection `json:"Clients,omitempty" name:"Clients"`

	// The total number of records that meet the query condition, which can be used for paginated queries.
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeClientConnectionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClientConnectionsResponseParams `json:"Response"`
}

func (r *DescribeClientConnectionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClientConnectionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBBackupsRequestParams struct {
	// Instance ID in the format of cmgo-p8vnipr5. It is the same as the instance ID displayed on the TencentDB Console page
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Backup mode. Valid values: `0` (logical backup), `1` (physical backup), `2` (both modes). Default value: `0`.
	BackupMethod *int64 `json:"BackupMethod,omitempty" name:"BackupMethod"`

	// Number of entries per page. Maximum value: `100`. If this parameter is left empty, all entries will be returned.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Pagination offset, starting from `0`. Default value: `0`.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeDBBackupsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of cmgo-p8vnipr5. It is the same as the instance ID displayed on the TencentDB Console page
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Backup mode. Valid values: `0` (logical backup), `1` (physical backup), `2` (both modes). Default value: `0`.
	BackupMethod *int64 `json:"BackupMethod,omitempty" name:"BackupMethod"`

	// Number of entries per page. Maximum value: `100`. If this parameter is left empty, all entries will be returned.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Pagination offset, starting from `0`. Default value: `0`.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeDBBackupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBBackupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BackupMethod")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBBackupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBBackupsResponseParams struct {
	// Backup list
	BackupList []*BackupInfo `json:"BackupList,omitempty" name:"BackupList"`

	// Total number of backups
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDBBackupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBBackupsResponseParams `json:"Response"`
}

func (r *DescribeDBBackupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBBackupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstanceDealRequestParams struct {
	// Order ID. It is returned by the `CreateDBInstance` and other APIs.
	DealId *string `json:"DealId,omitempty" name:"DealId"`
}

type DescribeDBInstanceDealRequest struct {
	*tchttp.BaseRequest
	
	// Order ID. It is returned by the `CreateDBInstance` and other APIs.
	DealId *string `json:"DealId,omitempty" name:"DealId"`
}

func (r *DescribeDBInstanceDealRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstanceDealRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DealId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBInstanceDealRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstanceDealResponseParams struct {
	// Order status. Valid values: 1 (unpaid), 2 (paid), 3 (delivering), 4 (delivered), 5 (delivery failed), 6 (refunded), 7 (order closed), 8 (order closed because it failed to be paid within timeout period).
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// Original price of the order.
	OriginalPrice *float64 `json:"OriginalPrice,omitempty" name:"OriginalPrice"`

	// Discounted price of the order.
	DiscountPrice *float64 `json:"DiscountPrice,omitempty" name:"DiscountPrice"`

	// Operation performed by the order. Valid values: purchase, renew, upgrade, downgrade, refund.
	Action *string `json:"Action,omitempty" name:"Action"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDBInstanceDealResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBInstanceDealResponseParams `json:"Response"`
}

func (r *DescribeDBInstanceDealResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstanceDealResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstancesRequestParams struct {
	// List of instance IDs in the format of cmgo-p8vnipr5. It is the same as the instance ID displayed on the TencentDB Console page
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// Instance type. Valid values: 0 (all instances), 1 (promoted), 2 (temp), 3 (read-only), -1 (promoted + read-only + disaster recovery)
	InstanceType *int64 `json:"InstanceType,omitempty" name:"InstanceType"`

	// Cluster type. Valid values: 0 (replica set instance), 1 (sharding instance), -1 (all instances)
	ClusterType *int64 `json:"ClusterType,omitempty" name:"ClusterType"`

	// Instance status. Valid values: `0` (to be initialized), `1` (executing task), `2` (running), `-2` (isolated monthly-subscribed instance), `-3` (isolated pay-as-you-go instance)
	Status []*int64 `json:"Status,omitempty" name:"Status"`

	// VPC ID. This parameter can be left empty for the basic network
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Subnet ID of VPC. This parameter can be left empty for the basic network. If it is passed in as an input parameter, the corresponding VpcId must be set
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// Billing type. Valid value: 0 (pay-as-you-go)
	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`

	// Number of results to be returned for a single request. Valid values: 1-100. Default value: 20
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Sort by field of the returned result set. Currently, supported values include "ProjectId", "InstanceName", and "CreateTime". The return results are sorted in ascending order by default.
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// Sorting method of the return result set. Currently, "ASC" or "DESC" is supported
	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`

	// Project ID
	ProjectIds []*uint64 `json:"ProjectIds,omitempty" name:"ProjectIds"`

	// Search keyword, which can be instance ID, instance name, or complete IP
	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`

	// Tag information
	Tags *TagInfo `json:"Tags,omitempty" name:"Tags"`
}

type DescribeDBInstancesRequest struct {
	*tchttp.BaseRequest
	
	// List of instance IDs in the format of cmgo-p8vnipr5. It is the same as the instance ID displayed on the TencentDB Console page
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// Instance type. Valid values: 0 (all instances), 1 (promoted), 2 (temp), 3 (read-only), -1 (promoted + read-only + disaster recovery)
	InstanceType *int64 `json:"InstanceType,omitempty" name:"InstanceType"`

	// Cluster type. Valid values: 0 (replica set instance), 1 (sharding instance), -1 (all instances)
	ClusterType *int64 `json:"ClusterType,omitempty" name:"ClusterType"`

	// Instance status. Valid values: `0` (to be initialized), `1` (executing task), `2` (running), `-2` (isolated monthly-subscribed instance), `-3` (isolated pay-as-you-go instance)
	Status []*int64 `json:"Status,omitempty" name:"Status"`

	// VPC ID. This parameter can be left empty for the basic network
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Subnet ID of VPC. This parameter can be left empty for the basic network. If it is passed in as an input parameter, the corresponding VpcId must be set
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// Billing type. Valid value: 0 (pay-as-you-go)
	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`

	// Number of results to be returned for a single request. Valid values: 1-100. Default value: 20
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Sort by field of the returned result set. Currently, supported values include "ProjectId", "InstanceName", and "CreateTime". The return results are sorted in ascending order by default.
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// Sorting method of the return result set. Currently, "ASC" or "DESC" is supported
	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`

	// Project ID
	ProjectIds []*uint64 `json:"ProjectIds,omitempty" name:"ProjectIds"`

	// Search keyword, which can be instance ID, instance name, or complete IP
	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`

	// Tag information
	Tags *TagInfo `json:"Tags,omitempty" name:"Tags"`
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
	delete(f, "InstanceIds")
	delete(f, "InstanceType")
	delete(f, "ClusterType")
	delete(f, "Status")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "PayMode")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	delete(f, "ProjectIds")
	delete(f, "SearchKey")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstancesResponseParams struct {
	// Number of eligible instances.
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// List of instance details
	InstanceDetails []*InstanceDetail `json:"InstanceDetails,omitempty" name:"InstanceDetails"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeInstanceParamsRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type DescribeInstanceParamsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
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

// Predefined struct for user
type DescribeInstanceParamsResponseParams struct {
	// The collection of enum parameters
	InstanceEnumParam []*InstanceEnumParam `json:"InstanceEnumParam,omitempty" name:"InstanceEnumParam"`

	// The collection of integer parameters
	InstanceIntegerParam []*InstanceIntegerParam `json:"InstanceIntegerParam,omitempty" name:"InstanceIntegerParam"`

	// The collection of text parameters
	InstanceTextParam []*InstanceTextParam `json:"InstanceTextParam,omitempty" name:"InstanceTextParam"`

	// The collection of string parameters used to represent time ranges
	InstanceMultiParam []*InstanceMultiParam `json:"InstanceMultiParam,omitempty" name:"InstanceMultiParam"`

	// The total number of modifiable parameters of the instance, such as 0
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeSecurityGroupRequestParams struct {
	// Instance ID in the format of "cmgo-p8vnipr5"
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type DescribeSecurityGroupRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of "cmgo-p8vnipr5"
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeSecurityGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecurityGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityGroupResponseParams struct {
	// Security groups associated with the instance
	Groups []*SecurityGroup `json:"Groups,omitempty" name:"Groups"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSecurityGroupResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSecurityGroupResponseParams `json:"Response"`
}

func (r *DescribeSecurityGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSlowLogPatternsRequestParams struct {
	// Instance ID in the format of `cmgo-p8vnipr5`, which is the same as the instance ID displayed on the TencentDB Console page
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Start time of slow log in the format of `yyyy-mm-dd hh:mm:ss`, such as 2019-06-01 10:00:00. The query time range cannot exceed 24 hours. Only slow logs for the last 7 days can be queried.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time of slow log in the format of `yyyy-mm-dd hh:mm:ss`, such as 2019-06-02 12:00:00. The query time range cannot exceed 24 hours. Only slow logs for the last 7 days can be queried.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Threshold of slow log execution time in milliseconds. Minimum value: 100. Slow logs whose execution time exceeds the threshold will be returned.
	SlowMS *uint64 `json:"SlowMS,omitempty" name:"SlowMS"`

	// Offset. Minimum value: 0. Maximum value: 10000. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of entries per page. Minimum value: 1. Maximum value: 100. Default value: 20.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Slow log format, which can be JSON. If this parameter is left empty, the slow log will be returned in its native format.
	Format *string `json:"Format,omitempty" name:"Format"`
}

type DescribeSlowLogPatternsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of `cmgo-p8vnipr5`, which is the same as the instance ID displayed on the TencentDB Console page
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Start time of slow log in the format of `yyyy-mm-dd hh:mm:ss`, such as 2019-06-01 10:00:00. The query time range cannot exceed 24 hours. Only slow logs for the last 7 days can be queried.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time of slow log in the format of `yyyy-mm-dd hh:mm:ss`, such as 2019-06-02 12:00:00. The query time range cannot exceed 24 hours. Only slow logs for the last 7 days can be queried.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Threshold of slow log execution time in milliseconds. Minimum value: 100. Slow logs whose execution time exceeds the threshold will be returned.
	SlowMS *uint64 `json:"SlowMS,omitempty" name:"SlowMS"`

	// Offset. Minimum value: 0. Maximum value: 10000. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of entries per page. Minimum value: 1. Maximum value: 100. Default value: 20.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Slow log format, which can be JSON. If this parameter is left empty, the slow log will be returned in its native format.
	Format *string `json:"Format,omitempty" name:"Format"`
}

func (r *DescribeSlowLogPatternsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlowLogPatternsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "SlowMS")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Format")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSlowLogPatternsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSlowLogPatternsResponseParams struct {
	// Total number of slow logs
	Count *uint64 `json:"Count,omitempty" name:"Count"`

	// Slow log statistics
	SlowLogPatterns []*SlowLogPattern `json:"SlowLogPatterns,omitempty" name:"SlowLogPatterns"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSlowLogPatternsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSlowLogPatternsResponseParams `json:"Response"`
}

func (r *DescribeSlowLogPatternsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlowLogPatternsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSlowLogsRequestParams struct {
	// Instance ID in the format of `cmgo-p8vnipr5`, which is the same as the instance ID displayed on the TencentDB Console page
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Start time of slow log in the format of `yyyy-mm-dd hh:mm:ss`, such as 2019-06-01 10:00:00. The query time range cannot exceed 24 hours. Only slow logs for the last 7 days can be queried.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time of slow log in the format of `yyyy-mm-dd hh:mm:ss`, such as 2019-06-02 12:00:00. The query time range cannot exceed 24 hours. Only slow logs for the last 7 days can be queried.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Threshold of slow log execution time in milliseconds. Minimum value: 100. Slow logs whose execution time exceeds the threshold will be returned.
	SlowMS *uint64 `json:"SlowMS,omitempty" name:"SlowMS"`

	// Offset. Minimum value: 0. Maximum value: 10000. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of entries per page. Minimum value: 1. Maximum value: 100. Default value: 20.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Slow log format, which can be JSON. If this parameter is left empty, the slow log will be returned in its native format.
	Format *string `json:"Format,omitempty" name:"Format"`
}

type DescribeSlowLogsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of `cmgo-p8vnipr5`, which is the same as the instance ID displayed on the TencentDB Console page
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Start time of slow log in the format of `yyyy-mm-dd hh:mm:ss`, such as 2019-06-01 10:00:00. The query time range cannot exceed 24 hours. Only slow logs for the last 7 days can be queried.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time of slow log in the format of `yyyy-mm-dd hh:mm:ss`, such as 2019-06-02 12:00:00. The query time range cannot exceed 24 hours. Only slow logs for the last 7 days can be queried.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Threshold of slow log execution time in milliseconds. Minimum value: 100. Slow logs whose execution time exceeds the threshold will be returned.
	SlowMS *uint64 `json:"SlowMS,omitempty" name:"SlowMS"`

	// Offset. Minimum value: 0. Maximum value: 10000. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of entries per page. Minimum value: 1. Maximum value: 100. Default value: 20.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Slow log format, which can be JSON. If this parameter is left empty, the slow log will be returned in its native format.
	Format *string `json:"Format,omitempty" name:"Format"`
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
	delete(f, "SlowMS")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Format")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSlowLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSlowLogsResponseParams struct {
	// Total number of slow logs
	Count *uint64 `json:"Count,omitempty" name:"Count"`

	// Slow log details
	SlowLogs []*string `json:"SlowLogs,omitempty" name:"SlowLogs"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeSpecInfoRequestParams struct {
	// AZ to be queried
	Zone *string `json:"Zone,omitempty" name:"Zone"`
}

type DescribeSpecInfoRequest struct {
	*tchttp.BaseRequest
	
	// AZ to be queried
	Zone *string `json:"Zone,omitempty" name:"Zone"`
}

func (r *DescribeSpecInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSpecInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Zone")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSpecInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSpecInfoResponseParams struct {
	// List of purchasable instance specifications
	SpecInfoList []*SpecificationInfo `json:"SpecInfoList,omitempty" name:"SpecInfoList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSpecInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSpecInfoResponseParams `json:"Response"`
}

func (r *DescribeSpecInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSpecInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type FlushInstanceRouterConfigRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type FlushInstanceRouterConfigRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *FlushInstanceRouterConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FlushInstanceRouterConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "FlushInstanceRouterConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type FlushInstanceRouterConfigResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type FlushInstanceRouterConfigResponse struct {
	*tchttp.BaseResponse
	Response *FlushInstanceRouterConfigResponseParams `json:"Response"`
}

func (r *FlushInstanceRouterConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FlushInstanceRouterConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquirePriceCreateDBInstancesRequestParams struct {
	// Instance region name in the format of ap-guangzhou-2.
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// The number of nodes in each replica set. The value range is subject to the response parameter of the `DescribeSpecInfo` API.
	NodeNum *int64 `json:"NodeNum,omitempty" name:"NodeNum"`

	// Instance memory size in GB.
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// Instance disk size in GB.
	Volume *int64 `json:"Volume,omitempty" name:"Volume"`

	// Version number. For the specific purchasable versions supported, please see the return result of the `DescribeSpecInfo` API. The correspondences between parameters and versions are as follows: MONGO_3_WT: MongoDB 3.2 WiredTiger Edition; MONGO_3_ROCKS: MongoDB 3.2 RocksDB Edition; MONGO_36_WT: MongoDB 3.6 WiredTiger Edition; MONGO_40_WT: MongoDB 4.0 WiredTiger Edition.
	MongoVersion *string `json:"MongoVersion,omitempty" name:"MongoVersion"`

	// Server type. Valid values: `HIO` (high IO), `HIO10G` (ten-gigabit high IO)
	MachineCode *string `json:"MachineCode,omitempty" name:"MachineCode"`

	// Number of instances. Minimum value: 1. Maximum value: 10.
	GoodsNum *int64 `json:"GoodsNum,omitempty" name:"GoodsNum"`

	// Instance validity period in months. Valid values: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36.
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// Instance type. Valid values: REPLSET (replica set), SHARD (sharded cluster), STANDALONE (single-node).
	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`

	// Number of replica sets. To create a replica set instance, set this parameter to 1; to create a shard instance, see the parameters returned by the `DescribeSpecInfo` API; to create a single-node instance, set this parameter to 0.
	ReplicateSetNum *int64 `json:"ReplicateSetNum,omitempty" name:"ReplicateSetNum"`
}

type InquirePriceCreateDBInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Instance region name in the format of ap-guangzhou-2.
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// The number of nodes in each replica set. The value range is subject to the response parameter of the `DescribeSpecInfo` API.
	NodeNum *int64 `json:"NodeNum,omitempty" name:"NodeNum"`

	// Instance memory size in GB.
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// Instance disk size in GB.
	Volume *int64 `json:"Volume,omitempty" name:"Volume"`

	// Version number. For the specific purchasable versions supported, please see the return result of the `DescribeSpecInfo` API. The correspondences between parameters and versions are as follows: MONGO_3_WT: MongoDB 3.2 WiredTiger Edition; MONGO_3_ROCKS: MongoDB 3.2 RocksDB Edition; MONGO_36_WT: MongoDB 3.6 WiredTiger Edition; MONGO_40_WT: MongoDB 4.0 WiredTiger Edition.
	MongoVersion *string `json:"MongoVersion,omitempty" name:"MongoVersion"`

	// Server type. Valid values: `HIO` (high IO), `HIO10G` (ten-gigabit high IO)
	MachineCode *string `json:"MachineCode,omitempty" name:"MachineCode"`

	// Number of instances. Minimum value: 1. Maximum value: 10.
	GoodsNum *int64 `json:"GoodsNum,omitempty" name:"GoodsNum"`

	// Instance validity period in months. Valid values: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36.
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// Instance type. Valid values: REPLSET (replica set), SHARD (sharded cluster), STANDALONE (single-node).
	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`

	// Number of replica sets. To create a replica set instance, set this parameter to 1; to create a shard instance, see the parameters returned by the `DescribeSpecInfo` API; to create a single-node instance, set this parameter to 0.
	ReplicateSetNum *int64 `json:"ReplicateSetNum,omitempty" name:"ReplicateSetNum"`
}

func (r *InquirePriceCreateDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquirePriceCreateDBInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Zone")
	delete(f, "NodeNum")
	delete(f, "Memory")
	delete(f, "Volume")
	delete(f, "MongoVersion")
	delete(f, "MachineCode")
	delete(f, "GoodsNum")
	delete(f, "Period")
	delete(f, "ClusterType")
	delete(f, "ReplicateSetNum")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquirePriceCreateDBInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquirePriceCreateDBInstancesResponseParams struct {
	// Price.
	Price *DBInstancePrice `json:"Price,omitempty" name:"Price"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type InquirePriceCreateDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *InquirePriceCreateDBInstancesResponseParams `json:"Response"`
}

func (r *InquirePriceCreateDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquirePriceCreateDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquirePriceModifyDBInstanceSpecRequestParams struct {
	// Instance ID in the format of cmgo-p8vnipr5. It is the same as the instance ID displayed in the TencentDB Console.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Instance memory size in GB after specification adjustment.
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// Instance disk size in GB after specification adjustment.
	Volume *int64 `json:"Volume,omitempty" name:"Volume"`

	// Node quantity after configuration modification. The value range is subject to the response parameter of the `DescribeSpecInfo` API. If this parameter is left empty, the node quantity remains unchanged.
	NodeNum *int64 `json:"NodeNum,omitempty" name:"NodeNum"`

	// Shard quantity after configuration modification, which can only be increased rather than decreased. The value range is subject to the response parameter of the `DescribeSpecInfo` API. If this parameter is left empty, the shard quantity remains unchanged.
	ReplicateSetNum *int64 `json:"ReplicateSetNum,omitempty" name:"ReplicateSetNum"`
}

type InquirePriceModifyDBInstanceSpecRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of cmgo-p8vnipr5. It is the same as the instance ID displayed in the TencentDB Console.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Instance memory size in GB after specification adjustment.
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// Instance disk size in GB after specification adjustment.
	Volume *int64 `json:"Volume,omitempty" name:"Volume"`

	// Node quantity after configuration modification. The value range is subject to the response parameter of the `DescribeSpecInfo` API. If this parameter is left empty, the node quantity remains unchanged.
	NodeNum *int64 `json:"NodeNum,omitempty" name:"NodeNum"`

	// Shard quantity after configuration modification, which can only be increased rather than decreased. The value range is subject to the response parameter of the `DescribeSpecInfo` API. If this parameter is left empty, the shard quantity remains unchanged.
	ReplicateSetNum *int64 `json:"ReplicateSetNum,omitempty" name:"ReplicateSetNum"`
}

func (r *InquirePriceModifyDBInstanceSpecRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquirePriceModifyDBInstanceSpecRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Memory")
	delete(f, "Volume")
	delete(f, "NodeNum")
	delete(f, "ReplicateSetNum")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquirePriceModifyDBInstanceSpecRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquirePriceModifyDBInstanceSpecResponseParams struct {
	// Price.
	Price *DBInstancePrice `json:"Price,omitempty" name:"Price"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type InquirePriceModifyDBInstanceSpecResponse struct {
	*tchttp.BaseResponse
	Response *InquirePriceModifyDBInstanceSpecResponseParams `json:"Response"`
}

func (r *InquirePriceModifyDBInstanceSpecResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquirePriceModifyDBInstanceSpecResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquirePriceRenewDBInstancesRequestParams struct {
	// Instance ID in the format of cmgo-p8vnipr5. It is the same as the instance ID displayed in the TencentDB Console. This API supports operations on up to 5 instances at a time.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// The parameter setting for the prepaid mode (monthly subscription mode). This parameter can specify the renewal period, whether to set automatic renewal, and other attributes of the monthly subscription instance.
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitempty" name:"InstanceChargePrepaid"`
}

type InquirePriceRenewDBInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of cmgo-p8vnipr5. It is the same as the instance ID displayed in the TencentDB Console. This API supports operations on up to 5 instances at a time.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// The parameter setting for the prepaid mode (monthly subscription mode). This parameter can specify the renewal period, whether to set automatic renewal, and other attributes of the monthly subscription instance.
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitempty" name:"InstanceChargePrepaid"`
}

func (r *InquirePriceRenewDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquirePriceRenewDBInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "InstanceChargePrepaid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquirePriceRenewDBInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquirePriceRenewDBInstancesResponseParams struct {
	// Price.
	Price *DBInstancePrice `json:"Price,omitempty" name:"Price"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type InquirePriceRenewDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *InquirePriceRenewDBInstancesResponseParams `json:"Response"`
}

func (r *InquirePriceRenewDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquirePriceRenewDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstanceChargePrepaid struct {
	// Purchased usage period (in month). Valid values: `1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36`. Default value: `1`.
	// (This parameter is required in `InquirePriceRenewDBInstances` and `RenewDBInstances` APIs.)
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// Auto-renewal flag. Valid values:
	// `NOTIFY_AND_AUTO_RENEW`: notify expiration and renew automatically
	// `NOTIFY_AND_MANUAL_RENEW`: notify expiration but not renew automatically
	// `DISABLE_NOTIFY_AND_MANUAL_RENEW`: neither notify expiration nor renew automatically
	// 
	// Default value: `NOTIFY_AND_MANUAL_RENEW`. If this parameter is specified as `NOTIFY_AND_AUTO_RENEW`, the instance will be automatically renewed on a monthly basis when the account balance is sufficient.
	// (This parameter is required in `InquirePriceRenewDBInstances` and `RenewDBInstances` APIs.)
	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`
}

type InstanceDetail struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Instance name
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// Billing type. Valid value: 0 (pay-as-you-go)
	PayMode *uint64 `json:"PayMode,omitempty" name:"PayMode"`

	// Project ID
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Cluster type. Valid values: 0 (replica set instance), 1 (sharding instance),
	ClusterType *uint64 `json:"ClusterType,omitempty" name:"ClusterType"`

	// Region information
	Region *string `json:"Region,omitempty" name:"Region"`

	// AZ information
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// Network type. Valid values: 0 (basic network), 1 (VPC)
	NetType *uint64 `json:"NetType,omitempty" name:"NetType"`

	// VPC ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Subnet ID of VPC
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// Instance status. Valid values: 0 (to be initialized), 1 (in process), 2 (running), -2 (expired)
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// Instance IP
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// Port number
	Vport *uint64 `json:"Vport,omitempty" name:"Vport"`

	// Instance creation time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Instance expiration time
	DeadLine *string `json:"DeadLine,omitempty" name:"DeadLine"`

	// Instance version information
	MongoVersion *string `json:"MongoVersion,omitempty" name:"MongoVersion"`

	// Instance memory size in MB
	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`

	// Instance disk size in MB
	Volume *uint64 `json:"Volume,omitempty" name:"Volume"`

	// Number of CPU cores of an instance
	CpuNum *uint64 `json:"CpuNum,omitempty" name:"CpuNum"`

	// Instance machine type
	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`

	// Number of secondary nodes of an instance
	SecondaryNum *uint64 `json:"SecondaryNum,omitempty" name:"SecondaryNum"`

	// Number of instance shards
	ReplicationSetNum *uint64 `json:"ReplicationSetNum,omitempty" name:"ReplicationSetNum"`

	// Instance auto-renewal flag. Valid values: 0 (manual renewal), 1 (auto-renewal), 2 (no renewal upon expiration)
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`

	// Used capacity in MB
	UsedVolume *uint64 `json:"UsedVolume,omitempty" name:"UsedVolume"`

	// Start time of the maintenance time window
	MaintenanceStart *string `json:"MaintenanceStart,omitempty" name:"MaintenanceStart"`

	// End time of the maintenance time window
	MaintenanceEnd *string `json:"MaintenanceEnd,omitempty" name:"MaintenanceEnd"`

	// Shard information
	ReplicaSets []*ShardInfo `json:"ReplicaSets,omitempty" name:"ReplicaSets"`

	// Information of read-only instances
	ReadonlyInstances []*DBInstanceInfo `json:"ReadonlyInstances,omitempty" name:"ReadonlyInstances"`

	// Information of disaster recovery instances
	StandbyInstances []*DBInstanceInfo `json:"StandbyInstances,omitempty" name:"StandbyInstances"`

	// Information of temp instances
	CloneInstances []*DBInstanceInfo `json:"CloneInstances,omitempty" name:"CloneInstances"`

	// Information of associated instances. For a promoted instance, this field represents information of its temp instance; for a temp instance, this field represents information of its promoted instance; and for a read-only/disaster recovery instance, this field represents information of its primary instance
	RelatedInstance *DBInstanceInfo `json:"RelatedInstance,omitempty" name:"RelatedInstance"`

	// Instance tag information set
	Tags []*TagInfo `json:"Tags,omitempty" name:"Tags"`

	// Instance version tag
	InstanceVer *uint64 `json:"InstanceVer,omitempty" name:"InstanceVer"`

	// Instance version tag
	ClusterVer *uint64 `json:"ClusterVer,omitempty" name:"ClusterVer"`

	// Protocol information. Valid values: 1 (mongodb), 2 (dynamodb)
	Protocol *uint64 `json:"Protocol,omitempty" name:"Protocol"`

	// Instance type. Valid values: 1 (promoted instance), 2 (temp instance), 3 (read-only instance), 4 (disaster recovery instance)
	InstanceType *uint64 `json:"InstanceType,omitempty" name:"InstanceType"`

	// Instance status description
	InstanceStatusDesc *string `json:"InstanceStatusDesc,omitempty" name:"InstanceStatusDesc"`

	// Physical instance ID. For an instance that has been rolled back and replaced, its InstanceId and RealInstanceId are different. The physical instance ID is needed in such scenarios as getting monitoring data from Barad
	RealInstanceId *string `json:"RealInstanceId,omitempty" name:"RealInstanceId"`
}

type InstanceEnumParam struct {
	// Current value
	CurrentValue *string `json:"CurrentValue,omitempty" name:"CurrentValue"`

	// Default value
	DefaultValue *string `json:"DefaultValue,omitempty" name:"DefaultValue"`

	// Acceptable values
	EnumValue []*string `json:"EnumValue,omitempty" name:"EnumValue"`

	// Whether to restart the instance for the parameter to take effect. Valid values: `1` (yes), `0` (no, which means the parameter setting takes effect immediately)
	NeedRestart *string `json:"NeedRestart,omitempty" name:"NeedRestart"`

	// Parameter name
	ParamName *string `json:"ParamName,omitempty" name:"ParamName"`

	// Parameter description
	Tips []*string `json:"Tips,omitempty" name:"Tips"`

	// Data type of the parameter
	ValueType *string `json:"ValueType,omitempty" name:"ValueType"`

	// Whether `CurrentValue` is the parameter value actually in use. Valid values: `1` (yes), `0` (no)
	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

type InstanceIntegerParam struct {
	// Current value
	CurrentValue *string `json:"CurrentValue,omitempty" name:"CurrentValue"`

	// Default value
	DefaultValue *string `json:"DefaultValue,omitempty" name:"DefaultValue"`

	// Maximum value
	Max *string `json:"Max,omitempty" name:"Max"`

	// Minimum value
	Min *string `json:"Min,omitempty" name:"Min"`

	// Whether to restart the instance for the parameter to take effect. Valid values: `1` (yes), `0` (no, which means the parameter setting takes effect immediately)
	NeedRestart *string `json:"NeedRestart,omitempty" name:"NeedRestart"`

	// Parameter name
	ParamName *string `json:"ParamName,omitempty" name:"ParamName"`

	// Parameter description
	Tips []*string `json:"Tips,omitempty" name:"Tips"`

	// Data type of the parameter
	ValueType *string `json:"ValueType,omitempty" name:"ValueType"`

	// Whether `CurrentValue` is the parameter value actually in use. Valid values: `1` (yes), `0` (no)
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// Redundant field which can be ignored
	Unit *string `json:"Unit,omitempty" name:"Unit"`
}

type InstanceMultiParam struct {
	// Current value
	CurrentValue *string `json:"CurrentValue,omitempty" name:"CurrentValue"`

	// Default value
	DefaultValue *string `json:"DefaultValue,omitempty" name:"DefaultValue"`

	// Acceptable values
	EnumValue []*string `json:"EnumValue,omitempty" name:"EnumValue"`

	// Whether to restart the instance for the parameter to take effect. Valid values: `1` (yes), `0` (no, which means the parameter setting takes effect immediately)
	NeedRestart *string `json:"NeedRestart,omitempty" name:"NeedRestart"`

	// Parameter name
	ParamName *string `json:"ParamName,omitempty" name:"ParamName"`

	// Whether `CurrentValue` is the parameter value actually in use. Valid values: `1` (yes), `0` (no)
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// Parameter description
	Tips []*string `json:"Tips,omitempty" name:"Tips"`

	// Data type of the current value. Default value: `multi`
	ValueType *string `json:"ValueType,omitempty" name:"ValueType"`
}

type InstanceTextParam struct {
	// Current value
	CurrentValue *string `json:"CurrentValue,omitempty" name:"CurrentValue"`

	// Default value
	DefaultValue *string `json:"DefaultValue,omitempty" name:"DefaultValue"`

	// Whether to restart the instance for the parameter to take effect
	NeedRestart *string `json:"NeedRestart,omitempty" name:"NeedRestart"`

	// Parameter name
	ParamName *string `json:"ParamName,omitempty" name:"ParamName"`

	// Value of a text parameter
	TextValue *string `json:"TextValue,omitempty" name:"TextValue"`

	// Parameter description
	Tips []*string `json:"Tips,omitempty" name:"Tips"`

	// Value type
	ValueType *string `json:"ValueType,omitempty" name:"ValueType"`

	// Whether `CurrentValue` is the parameter value actually in use. Valid values: `1` (yes), `0` (no)
	Status *string `json:"Status,omitempty" name:"Status"`
}

// Predefined struct for user
type IsolateDBInstanceRequestParams struct {
	// Instance ID in the format of cmgo-p8vnipr5. It is the same as the instance ID displayed on the TencentDB Console page
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type IsolateDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of cmgo-p8vnipr5. It is the same as the instance ID displayed on the TencentDB Console page
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

// Predefined struct for user
type IsolateDBInstanceResponseParams struct {
	// Async task request ID, which can be used to query the execution result of an async task.
	AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

// Predefined struct for user
type ModifyDBInstanceNetworkAddressRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Old IP retention period in minutes. The old IP will be released after the specified time, and both the old and new IPs can be accessed before the release. The value `0` indicates that the old IP will be reclaimed immediately.
	OldIpExpiredTime *uint64 `json:"OldIpExpiredTime,omitempty" name:"OldIpExpiredTime"`

	// ID of the VPC to which the new IP belongs after the switch. When it is classic network, this field will be empty.
	NewUniqVpcId *string `json:"NewUniqVpcId,omitempty" name:"NewUniqVpcId"`

	// ID of the subnet to which the new IP belongs after the switch. When it is classic network, this field will be empty.
	NewUniqSubnetId *string `json:"NewUniqSubnetId,omitempty" name:"NewUniqSubnetId"`

	// IP information to be modified
	NetworkAddresses []*ModifyNetworkAddress `json:"NetworkAddresses,omitempty" name:"NetworkAddresses"`
}

type ModifyDBInstanceNetworkAddressRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Old IP retention period in minutes. The old IP will be released after the specified time, and both the old and new IPs can be accessed before the release. The value `0` indicates that the old IP will be reclaimed immediately.
	OldIpExpiredTime *uint64 `json:"OldIpExpiredTime,omitempty" name:"OldIpExpiredTime"`

	// ID of the VPC to which the new IP belongs after the switch. When it is classic network, this field will be empty.
	NewUniqVpcId *string `json:"NewUniqVpcId,omitempty" name:"NewUniqVpcId"`

	// ID of the subnet to which the new IP belongs after the switch. When it is classic network, this field will be empty.
	NewUniqSubnetId *string `json:"NewUniqSubnetId,omitempty" name:"NewUniqSubnetId"`

	// IP information to be modified
	NetworkAddresses []*ModifyNetworkAddress `json:"NetworkAddresses,omitempty" name:"NetworkAddresses"`
}

func (r *ModifyDBInstanceNetworkAddressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceNetworkAddressRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "OldIpExpiredTime")
	delete(f, "NewUniqVpcId")
	delete(f, "NewUniqSubnetId")
	delete(f, "NetworkAddresses")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDBInstanceNetworkAddressRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceNetworkAddressResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyDBInstanceNetworkAddressResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDBInstanceNetworkAddressResponseParams `json:"Response"`
}

func (r *ModifyDBInstanceNetworkAddressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceNetworkAddressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceSecurityGroupRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Target security group IDs
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`
}

type ModifyDBInstanceSecurityGroupRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Target security group IDs
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`
}

func (r *ModifyDBInstanceSecurityGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceSecurityGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SecurityGroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDBInstanceSecurityGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceSecurityGroupResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyDBInstanceSecurityGroupResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDBInstanceSecurityGroupResponseParams `json:"Response"`
}

func (r *ModifyDBInstanceSecurityGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceSecurityGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceSpecRequestParams struct {
	// Instance ID in the format of cmgo-p8vnipr5. It is the same as the instance ID displayed on the TencentDB Console page
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Memory size after instance configuration change in GB. Memory and disk must be upgraded or degraded simultaneously
	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`

	// Disk size after instance configuration change in GB. Memory and disk must be upgraded or degraded simultaneously. For degradation, the new disk capacity must be greater than 1.2 times the used disk capacity
	Volume *uint64 `json:"Volume,omitempty" name:"Volume"`

	// Oplog size after instance configuration change in GB, which ranges from 10% to 90% of the disk capacity and is 10% of the disk capacity by default
	OplogSize *uint64 `json:"OplogSize,omitempty" name:"OplogSize"`

	// Node quantity after configuration modification. The value range is subject to the response parameter of the `DescribeSpecInfo` API. If this parameter is left empty, the node quantity remains unchanged.
	NodeNum *uint64 `json:"NodeNum,omitempty" name:"NodeNum"`

	// Shard quantity after configuration modification, which can only be increased rather than decreased. The value range is subject to the response parameter of the `DescribeSpecInfo` API. If this parameter is left empty, the shard quantity remains unchanged.
	ReplicateSetNum *uint64 `json:"ReplicateSetNum,omitempty" name:"ReplicateSetNum"`

	// Switch time. Valid values: `0` (upon modification completion), `1` (during maintenance time). Default value: `0`. If the quantity of nodes or shards is modified, the value will be `0`.
	InMaintenance *uint64 `json:"InMaintenance,omitempty" name:"InMaintenance"`
}

type ModifyDBInstanceSpecRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of cmgo-p8vnipr5. It is the same as the instance ID displayed on the TencentDB Console page
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Memory size after instance configuration change in GB. Memory and disk must be upgraded or degraded simultaneously
	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`

	// Disk size after instance configuration change in GB. Memory and disk must be upgraded or degraded simultaneously. For degradation, the new disk capacity must be greater than 1.2 times the used disk capacity
	Volume *uint64 `json:"Volume,omitempty" name:"Volume"`

	// Oplog size after instance configuration change in GB, which ranges from 10% to 90% of the disk capacity and is 10% of the disk capacity by default
	OplogSize *uint64 `json:"OplogSize,omitempty" name:"OplogSize"`

	// Node quantity after configuration modification. The value range is subject to the response parameter of the `DescribeSpecInfo` API. If this parameter is left empty, the node quantity remains unchanged.
	NodeNum *uint64 `json:"NodeNum,omitempty" name:"NodeNum"`

	// Shard quantity after configuration modification, which can only be increased rather than decreased. The value range is subject to the response parameter of the `DescribeSpecInfo` API. If this parameter is left empty, the shard quantity remains unchanged.
	ReplicateSetNum *uint64 `json:"ReplicateSetNum,omitempty" name:"ReplicateSetNum"`

	// Switch time. Valid values: `0` (upon modification completion), `1` (during maintenance time). Default value: `0`. If the quantity of nodes or shards is modified, the value will be `0`.
	InMaintenance *uint64 `json:"InMaintenance,omitempty" name:"InMaintenance"`
}

func (r *ModifyDBInstanceSpecRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceSpecRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Memory")
	delete(f, "Volume")
	delete(f, "OplogSize")
	delete(f, "NodeNum")
	delete(f, "ReplicateSetNum")
	delete(f, "InMaintenance")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDBInstanceSpecRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceSpecResponseParams struct {
	// Order ID
	DealId *string `json:"DealId,omitempty" name:"DealId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyDBInstanceSpecResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDBInstanceSpecResponseParams `json:"Response"`
}

func (r *ModifyDBInstanceSpecResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceSpecResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNetworkAddress struct {
	// New IP
	NewIPAddress *string `json:"NewIPAddress,omitempty" name:"NewIPAddress"`

	// Old IP
	OldIpAddress *string `json:"OldIpAddress,omitempty" name:"OldIpAddress"`
}

// Predefined struct for user
type OfflineIsolatedDBInstanceRequestParams struct {
	// Instance ID in the format of cmgo-p8vnipr5. It is the same as the instance ID displayed on the TencentDB Console page
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type OfflineIsolatedDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of cmgo-p8vnipr5. It is the same as the instance ID displayed on the TencentDB Console page
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *OfflineIsolatedDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OfflineIsolatedDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OfflineIsolatedDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OfflineIsolatedDBInstanceResponseParams struct {
	// Async task request ID, which can be used to query the execution result of an async task.
	AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type OfflineIsolatedDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *OfflineIsolatedDBInstanceResponseParams `json:"Response"`
}

func (r *OfflineIsolatedDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OfflineIsolatedDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RenameInstanceRequestParams struct {
	// Instance ID in the format of cmgo-p8vnipr5. It is the same as the instance ID displayed on the TencentDB Console page
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Custom name of the instance, which can contain up to 60 letters, digits, or symbols (_-)
	NewName *string `json:"NewName,omitempty" name:"NewName"`
}

type RenameInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of cmgo-p8vnipr5. It is the same as the instance ID displayed on the TencentDB Console page
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Custom name of the instance, which can contain up to 60 letters, digits, or symbols (_-)
	NewName *string `json:"NewName,omitempty" name:"NewName"`
}

func (r *RenameInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenameInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "NewName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RenameInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RenameInstanceResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RenameInstanceResponse struct {
	*tchttp.BaseResponse
	Response *RenameInstanceResponseParams `json:"Response"`
}

func (r *RenameInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenameInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RenewDBInstancesRequestParams struct {
	// IDs of one or more instances to be operated. The value can be obtained from the `InstanceId` parameter returned by the `DescribeInstances` API. Up to 100 instances can be requested at a time.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// The parameter setting for the prepaid mode (monthly subscription mode). This parameter can specify the renewal period, whether to set automatic renewal, and other attributes of the monthly subscription instance. This parameter is mandatory in monthly subscription.
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitempty" name:"InstanceChargePrepaid"`
}

type RenewDBInstancesRequest struct {
	*tchttp.BaseRequest
	
	// IDs of one or more instances to be operated. The value can be obtained from the `InstanceId` parameter returned by the `DescribeInstances` API. Up to 100 instances can be requested at a time.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// The parameter setting for the prepaid mode (monthly subscription mode). This parameter can specify the renewal period, whether to set automatic renewal, and other attributes of the monthly subscription instance. This parameter is mandatory in monthly subscription.
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitempty" name:"InstanceChargePrepaid"`
}

func (r *RenewDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewDBInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "InstanceChargePrepaid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RenewDBInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RenewDBInstancesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RenewDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *RenewDBInstancesResponseParams `json:"Response"`
}

func (r *RenewDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReplicaSetInfo struct {
	// Replica set ID
	ReplicaSetId *string `json:"ReplicaSetId,omitempty" name:"ReplicaSetId"`
}

// Predefined struct for user
type ResetDBInstancePasswordRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Instance account name
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// New password, which must contain at least eight characters
	Password *string `json:"Password,omitempty" name:"Password"`
}

type ResetDBInstancePasswordRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Instance account name
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// New password, which must contain at least eight characters
	Password *string `json:"Password,omitempty" name:"Password"`
}

func (r *ResetDBInstancePasswordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetDBInstancePasswordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "UserName")
	delete(f, "Password")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetDBInstancePasswordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetDBInstancePasswordResponseParams struct {
	// Async request ID, which is used to query the running status of the process.
	AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ResetDBInstancePasswordResponse struct {
	*tchttp.BaseResponse
	Response *ResetDBInstancePasswordResponseParams `json:"Response"`
}

func (r *ResetDBInstancePasswordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetDBInstancePasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SecurityGroup struct {
	// Project ID
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Creation time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Inbound rule
	Inbound []*SecurityGroupBound `json:"Inbound,omitempty" name:"Inbound"`

	// Outbound rule
	Outbound []*SecurityGroupBound `json:"Outbound,omitempty" name:"Outbound"`

	// Security group ID
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`

	// Security group name
	SecurityGroupName *string `json:"SecurityGroupName,omitempty" name:"SecurityGroupName"`

	// Security group remarks
	SecurityGroupRemark *string `json:"SecurityGroupRemark,omitempty" name:"SecurityGroupRemark"`
}

type SecurityGroupBound struct {
	// Execution rule. Valid values: `ACCEPT`, `DROP`
	Action *string `json:"Action,omitempty" name:"Action"`

	// IP range
	CidrIp *string `json:"CidrIp,omitempty" name:"CidrIp"`

	// Port range
	PortRange *string `json:"PortRange,omitempty" name:"PortRange"`

	// Transport layer protocol. Valid values: `tcp`, `udp`, `ALL`
	IpProtocol *string `json:"IpProtocol,omitempty" name:"IpProtocol"`

	// All the addresses that the security group ID represents
	Id *string `json:"Id,omitempty" name:"Id"`

	// All the addresses that the address group ID represents
	AddressModule *string `json:"AddressModule,omitempty" name:"AddressModule"`

	// All the protocols and ports that the service group ID represents
	ServiceModule *string `json:"ServiceModule,omitempty" name:"ServiceModule"`

	// Description
	Desc *string `json:"Desc,omitempty" name:"Desc"`
}

type ShardInfo struct {
	// Used shard capacity
	UsedVolume *float64 `json:"UsedVolume,omitempty" name:"UsedVolume"`

	// Shard ID
	ReplicaSetId *string `json:"ReplicaSetId,omitempty" name:"ReplicaSetId"`

	// Shard name
	ReplicaSetName *string `json:"ReplicaSetName,omitempty" name:"ReplicaSetName"`

	// Shard memory size in MB
	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`

	// Shard disk size in MB
	Volume *uint64 `json:"Volume,omitempty" name:"Volume"`

	// Shard oplog size in MB
	OplogSize *uint64 `json:"OplogSize,omitempty" name:"OplogSize"`

	// Number of secondary nodes of a shard
	SecondaryNum *uint64 `json:"SecondaryNum,omitempty" name:"SecondaryNum"`

	// Shard physical ID
	RealReplicaSetId *string `json:"RealReplicaSetId,omitempty" name:"RealReplicaSetId"`
}

type SlowLogPattern struct {
	// Slow log pattern
	Pattern *string `json:"Pattern,omitempty" name:"Pattern"`

	// Maximum execution time
	MaxTime *uint64 `json:"MaxTime,omitempty" name:"MaxTime"`

	// Average execution time
	AverageTime *uint64 `json:"AverageTime,omitempty" name:"AverageTime"`

	// Number of slow logs in this pattern
	Total *uint64 `json:"Total,omitempty" name:"Total"`
}

type SpecItem struct {
	// Specification information identifier
	SpecCode *string `json:"SpecCode,omitempty" name:"SpecCode"`

	// Specification purchasable flag. Valid values: 0 (not purchasable), 1 (purchasable)
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// Computing resource specification in terms of CPU core
	Cpu *uint64 `json:"Cpu,omitempty" name:"Cpu"`

	// Memory size in MB
	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`

	// Default disk size in MB
	DefaultStorage *uint64 `json:"DefaultStorage,omitempty" name:"DefaultStorage"`

	// Maximum disk size in MB
	MaxStorage *uint64 `json:"MaxStorage,omitempty" name:"MaxStorage"`

	// Minimum disk size in MB
	MinStorage *uint64 `json:"MinStorage,omitempty" name:"MinStorage"`

	// Maximum QPS
	Qps *uint64 `json:"Qps,omitempty" name:"Qps"`

	// Maximum number of connections
	Conns *uint64 `json:"Conns,omitempty" name:"Conns"`

	// MongoDB version information of an instance
	MongoVersionCode *string `json:"MongoVersionCode,omitempty" name:"MongoVersionCode"`

	// MongoDB version number of an instance
	MongoVersionValue *uint64 `json:"MongoVersionValue,omitempty" name:"MongoVersionValue"`

	// MongoDB version number of an instance (short)
	Version *string `json:"Version,omitempty" name:"Version"`

	// Storage engine
	EngineName *string `json:"EngineName,omitempty" name:"EngineName"`

	// Cluster type. Valid values: 1 (sharding cluster), 0 (replica set cluster)
	ClusterType *uint64 `json:"ClusterType,omitempty" name:"ClusterType"`

	// Minimum number of secondary nodes in a replica set
	MinNodeNum *uint64 `json:"MinNodeNum,omitempty" name:"MinNodeNum"`

	// Maximum number of secondary nodes in a replica set
	MaxNodeNum *uint64 `json:"MaxNodeNum,omitempty" name:"MaxNodeNum"`

	// Minimum number of shards
	MinReplicateSetNum *uint64 `json:"MinReplicateSetNum,omitempty" name:"MinReplicateSetNum"`

	// Maximum number of shards
	MaxReplicateSetNum *uint64 `json:"MaxReplicateSetNum,omitempty" name:"MaxReplicateSetNum"`

	// Minimum number of secondary nodes in a shard
	MinReplicateSetNodeNum *uint64 `json:"MinReplicateSetNodeNum,omitempty" name:"MinReplicateSetNodeNum"`

	// Maximum number of secondary nodes in a shard
	MaxReplicateSetNodeNum *uint64 `json:"MaxReplicateSetNodeNum,omitempty" name:"MaxReplicateSetNodeNum"`

	// Server type. Valid values: 0 (HIO), 4 (HIO10G)
	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`
}

type SpecificationInfo struct {
	// Region information
	Region *string `json:"Region,omitempty" name:"Region"`

	// AZ information
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// Purchasable specification information
	SpecItems []*SpecItem `json:"SpecItems,omitempty" name:"SpecItems"`

	// Whether cross-AZ deployment is supported. Valid values: `1` (yes), `0` (no).
	SupportMultiAZ *int64 `json:"SupportMultiAZ,omitempty" name:"SupportMultiAZ"`
}

type TagInfo struct {
	// Tag key
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// Tag value
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}