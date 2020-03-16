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

package v20180330

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

type CompleteMigrateJobRequest struct {
	*tchttp.BaseRequest

	// Data migration task ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

func (r *CompleteMigrateJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CompleteMigrateJobRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CompleteMigrateJobResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CompleteMigrateJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CompleteMigrateJobResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ConsistencyParams struct {

	// Data content check parameter, which refers to the proportion of the rows selected for data comparison in all the rows of the table. Value: an integer between 1 and 100.
	SelectRowsPerTable *int64 `json:"SelectRowsPerTable,omitempty" name:"SelectRowsPerTable"`

	// Data content check parameter, which refers to the proportion of the tables selected for data detection in all the tables. Value: an integer between 1 and 100.
	TablesSelectAll *int64 `json:"TablesSelectAll,omitempty" name:"TablesSelectAll"`

	// Data quantity check parameter, which checks whether the numbers of rows are identical. It refers to the proportion of the tables selected for quantity check in all the tables. Value: an integer between 1 and 100.
	TablesSelectCount *int64 `json:"TablesSelectCount,omitempty" name:"TablesSelectCount"`
}

type CreateMigrateCheckJobRequest struct {
	*tchttp.BaseRequest

	// Data migration task ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

func (r *CreateMigrateCheckJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateMigrateCheckJobRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateMigrateCheckJobResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateMigrateCheckJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateMigrateCheckJobResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateMigrateJobRequest struct {
	*tchttp.BaseRequest

	// Data migration task name
	JobName *string `json:"JobName,omitempty" name:"JobName"`

	// Migration task configuration options
	MigrateOption *MigrateOption `json:"MigrateOption,omitempty" name:"MigrateOption"`

	// Source instance database type, which currently supports MySQL, Redis, MongoDB, PostgreSQL, MariaDB, and Percona. For more information on the supported types in a specific region, see the migration task creation page in the console.
	SrcDatabaseType *string `json:"SrcDatabaseType,omitempty" name:"SrcDatabaseType"`

	// Source instance access type. Valid values: extranet (public network), cvm (CVM-based self-created instance), dcg (Direct Connect-enabled instance), vpncloud (Tencent Cloud VPN-enabled instance), cdb (TencentDB instance), ccn (CCN instance)
	SrcAccessType *string `json:"SrcAccessType,omitempty" name:"SrcAccessType"`

	// Source instance information, which is correlated with the migration task type
	SrcInfo *SrcInfo `json:"SrcInfo,omitempty" name:"SrcInfo"`

	// Target instance access type, which currently supports MySQL, Redis, MongoDB, PostgreSQL, MariaDB, and Percona. For more information on the supported types in a specific region, see the migration task creation page in the console.
	DstDatabaseType *string `json:"DstDatabaseType,omitempty" name:"DstDatabaseType"`

	// Target instance access type, which currently only supports cdb (TencentDB instance)
	DstAccessType *string `json:"DstAccessType,omitempty" name:"DstAccessType"`

	// Destination instance information
	DstInfo *DstInfo `json:"DstInfo,omitempty" name:"DstInfo"`

	// Information of the source table to be migrated, which is described in JSON string format. It is required if MigrateOption.MigrateObject is 2 (migrating the specified table).
	// For databases with a database-table structure:
	// [{Database:db1,Table:[table1,table2]},{Database:db2}]
	// For databases with a database-schema-table structure:
	// [{Database:db1,Schema:s1
	// Table:[table1,table2]},{Database:db1,Schema:s2
	// Table:[table1,table2]},{Database:db2,Schema:s1
	// Table:[table1,table2]},{Database:db3},{Database:db4
	// Schema:s1}]
	DatabaseInfo *string `json:"DatabaseInfo,omitempty" name:"DatabaseInfo"`
}

func (r *CreateMigrateJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateMigrateJobRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateMigrateJobResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Data migration task ID
		JobId *string `json:"JobId,omitempty" name:"JobId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateMigrateJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateMigrateJobResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateSyncCheckJobRequest struct {
	*tchttp.BaseRequest

	// Disaster recovery sync task ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

func (r *CreateSyncCheckJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateSyncCheckJobRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateSyncCheckJobResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSyncCheckJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateSyncCheckJobResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateSyncJobRequest struct {
	*tchttp.BaseRequest

	// Disaster recovery sync task name
	JobName *string `json:"JobName,omitempty" name:"JobName"`

	// Configuration options of a disaster recovery sync task
	SyncOption *SyncOption `json:"SyncOption,omitempty" name:"SyncOption"`

	// Source instance database type, which currently only supports mysql
	SrcDatabaseType *string `json:"SrcDatabaseType,omitempty" name:"SrcDatabaseType"`

	// Source instance access type, which currently only supports cdb (TencentDB instances)
	SrcAccessType *string `json:"SrcAccessType,omitempty" name:"SrcAccessType"`

	// Source instance information
	SrcInfo *SyncInstanceInfo `json:"SrcInfo,omitempty" name:"SrcInfo"`

	// Target instance access type, which currently only supports mysql
	DstDatabaseType *string `json:"DstDatabaseType,omitempty" name:"DstDatabaseType"`

	// Target instance access type, which currently only supports cdb (TencentDB instances)
	DstAccessType *string `json:"DstAccessType,omitempty" name:"DstAccessType"`

	// Target instance information
	DstInfo *SyncInstanceInfo `json:"DstInfo,omitempty" name:"DstInfo"`

	// Information of the source table to be synced, which is described in JSON string format.
	// For databases with a database-table structure:
	// [{Database:db1,Table:[table1,table2]},{Database:db2}]
	DatabaseInfo *string `json:"DatabaseInfo,omitempty" name:"DatabaseInfo"`
}

func (r *CreateSyncJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateSyncJobRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateSyncJobResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Disaster recovery sync task ID
		JobId *string `json:"JobId,omitempty" name:"JobId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSyncJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateSyncJobResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteMigrateJobRequest struct {
	*tchttp.BaseRequest

	// Data migration task ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

func (r *DeleteMigrateJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteMigrateJobRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteMigrateJobResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteMigrateJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteMigrateJobResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteSyncJobRequest struct {
	*tchttp.BaseRequest

	// ID of the disaster recovery sync task to be deleted
	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

func (r *DeleteSyncJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteSyncJobRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteSyncJobResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSyncJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteSyncJobResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMigrateCheckJobRequest struct {
	*tchttp.BaseRequest

	// Data migration task ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

func (r *DescribeMigrateCheckJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMigrateCheckJobRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMigrateCheckJobResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Check task status: unavailable, starting, running, finished
		Status *string `json:"Status,omitempty" name:"Status"`

		// Task error code
		ErrorCode *int64 `json:"ErrorCode,omitempty" name:"ErrorCode"`

		// Task error message
		ErrorMessage *string `json:"ErrorMessage,omitempty" name:"ErrorMessage"`

		// Check task progress. For example, "30" means 30% completed
		Progress *string `json:"Progress,omitempty" name:"Progress"`

		// Whether the check succeeds. 0: no; 1: yes; 3: not checked
		CheckFlag *int64 `json:"CheckFlag,omitempty" name:"CheckFlag"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMigrateCheckJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMigrateCheckJobResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMigrateJobsRequest struct {
	*tchttp.BaseRequest

	// Data migration task ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// Data migration task name
	JobName *string `json:"JobName,omitempty" name:"JobName"`

	// Sort by field. Value range: JobId, Status, JobName, MigrateType, RunMode, CreateTime
	Order *string `json:"Order,omitempty" name:"Order"`

	// Sorting order. Value range: ASC (ascending), DESC (descending)
	OrderSeq *string `json:"OrderSeq,omitempty" name:"OrderSeq"`

	// Offset. Default value: 0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of the returned instances. Value range: [1, 100]. Default value: 20
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeMigrateJobsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMigrateJobsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMigrateJobsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of tasks
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Array of task details
		JobList []*MigrateJobInfo `json:"JobList,omitempty" name:"JobList" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMigrateJobsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMigrateJobsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSyncCheckJobRequest struct {
	*tchttp.BaseRequest

	// ID of the disaster recovery sync task to be queried
	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

func (r *DescribeSyncCheckJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSyncCheckJobRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSyncCheckJobResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Task check status: starting, running, finished
		Status *string `json:"Status,omitempty" name:"Status"`

		// Code of the task check result
		ErrorCode *int64 `json:"ErrorCode,omitempty" name:"ErrorCode"`

		// Prompt message
		ErrorMessage *string `json:"ErrorMessage,omitempty" name:"ErrorMessage"`

		// Description of a task execution step
		StepInfo []*SyncCheckStepInfo `json:"StepInfo,omitempty" name:"StepInfo" list`

		// Check flag. 0: checking; 1: successfully checked
		CheckFlag *int64 `json:"CheckFlag,omitempty" name:"CheckFlag"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSyncCheckJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSyncCheckJobResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSyncJobsRequest struct {
	*tchttp.BaseRequest

	// Disaster recovery sync task ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// Disaster recovery sync task name
	JobName *string `json:"JobName,omitempty" name:"JobName"`

	// Sort by field. Value range: JobId, Status, JobName, CreateTime
	Order *string `json:"Order,omitempty" name:"Order"`

	// Sorting order. Value range: ASC (ascending), DESC (descending)
	OrderSeq *string `json:"OrderSeq,omitempty" name:"OrderSeq"`

	// Offset. Default value: 0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of the returned instances. Value range: [1, 100]. Default value: 20
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeSyncJobsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSyncJobsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSyncJobsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of tasks
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Array of task details
		JobList []*SyncJobInfo `json:"JobList,omitempty" name:"JobList" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSyncJobsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSyncJobsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DstInfo struct {

	// Target instance ID, such as cdb-jd92ijd8
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Target instance region, such as ap-guangzhou
	Region *string `json:"Region,omitempty" name:"Region"`

	// Target instance VIP, which has been disused and does not need to be entered
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// Target instance Vport, which has been disused and does not need to be entered
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// Only valid for MySQL currently. For instance-level migration, the value range is: 1 (read-only), 0 (read/write)
	ReadOnly *int64 `json:"ReadOnly,omitempty" name:"ReadOnly"`
}

type MigrateDetailInfo struct {

	// Total number of steps
	StepAll *int64 `json:"StepAll,omitempty" name:"StepAll"`

	// Current step
	StepNow *int64 `json:"StepNow,omitempty" name:"StepNow"`

	// Overall progress, such as:
	Progress *string `json:"Progress,omitempty" name:"Progress"`

	// Progress of the current step, such as:
	CurrentStepProgress *string `json:"CurrentStepProgress,omitempty" name:"CurrentStepProgress"`

	// Master/slave lag in MB, which is valid during incremental sync and currently supported by TencentDB for Redis and MySQL
	MasterSlaveDistance *int64 `json:"MasterSlaveDistance,omitempty" name:"MasterSlaveDistance"`

	// Master/slave lag in seconds, which is valid during incremental sync and currently supported by TencentDB for MySQL
	SecondsBehindMaster *int64 `json:"SecondsBehindMaster,omitempty" name:"SecondsBehindMaster"`

	// Step information
	StepInfo []*MigrateStepDetailInfo `json:"StepInfo,omitempty" name:"StepInfo" list`
}

type MigrateJobInfo struct {

	// Data migration task ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// Data migration task name
	JobName *string `json:"JobName,omitempty" name:"JobName"`

	// Migration task configuration options
	MigrateOption *MigrateOption `json:"MigrateOption,omitempty" name:"MigrateOption"`

	// Source instance database type: MySQL, Redis, MongoDB, PostgreSQL, MariaDB, Percona
	SrcDatabaseType *string `json:"SrcDatabaseType,omitempty" name:"SrcDatabaseType"`

	// Source instance access type. Value range: extranet (public network), cvm (CVM-created instance), dcg (Direct Connect-enabled instance), vpncloud (Tencent Cloud VPN-enabled instance), cdb (TencentDB instance), ccn (CCN instances)
	SrcAccessType *string `json:"SrcAccessType,omitempty" name:"SrcAccessType"`

	// Source instance information, which is correlated with the migration task type
	SrcInfo *SrcInfo `json:"SrcInfo,omitempty" name:"SrcInfo"`

	// Target instance access type: MySQL, Redis, MongoDB, PostgreSQL, MariaDB, Percona
	DstDatabaseType *string `json:"DstDatabaseType,omitempty" name:"DstDatabaseType"`

	// Target instance access type, which currently only supports cdb (TencentDB instance)
	DstAccessType *string `json:"DstAccessType,omitempty" name:"DstAccessType"`

	// Target instance information
	DstInfo *DstInfo `json:"DstInfo,omitempty" name:"DstInfo"`

	// Information of the source table to be migrated. If the entire instance is to be migrated, this field should be []
	DatabaseInfo *string `json:"DatabaseInfo,omitempty" name:"DatabaseInfo"`

	// Task creation/submission time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Task start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Task end time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Task status. Value range: 1 (Creating), 3 (Checking), 4 (CheckPass), 5 (CheckNotPass), 7 (Running), 8 (ReadyComplete), 9 (Success), 10 (Failed), 11 (Stopping), 12 (Completing)
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// Task details
	Detail *MigrateDetailInfo `json:"Detail,omitempty" name:"Detail"`
}

type MigrateOption struct {

	// Task operation mode. Value range: 1 (immediate execution), 2 (scheduled execution)
	RunMode *int64 `json:"RunMode,omitempty" name:"RunMode"`

	// Expected execution time in the format of yyyy-mm-dd hh:mm:ss. If runMode=2, this field is required
	ExpectTime *string `json:"ExpectTime,omitempty" name:"ExpectTime"`

	// Data migration type. Value range: 1 (structural migration), 2 (full migration), 3 (full + incremental migration)
	MigrateType *int64 `json:"MigrateType,omitempty" name:"MigrateType"`

	// Migration subject. 1: entire instance; 2: specified table
	MigrateObject *int64 `json:"MigrateObject,omitempty" name:"MigrateObject"`

	// Parameter of spot data consistency check. 1: not configured; 2: full check; 3: spot check; 4: check inconsistent tables only; 5: no check
	ConsistencyType *int64 `json:"ConsistencyType,omitempty" name:"ConsistencyType"`

	// Whether to overwrite the target database with the root account of the source database. Value range: 0 (no), 1 (yes). This value should be 0 for table or structural migration
	IsOverrideRoot *int64 `json:"IsOverrideRoot,omitempty" name:"IsOverrideRoot"`

	// Additional parameters for different databases, which are described in JSON format. 
	// The following parameters can be defined for Redis: 
	// { 
	// 	"ClientOutputBufferHardLimit":512, 	Hard capacity limit of slave buffer (MB) 
	// 	"ClientOutputBufferSoftLimit":512, 	Soft capacity limit of slave buffer (MB) 
	// 	"ClientOutputBufferPersistTime":60, Soft limit duration of slave buffer (s) 
	// 	"ReplBacklogSize":512, 	Circular buffer capacity limit (MB) 
	// 	"ReplTimeout":120, 		Replication timeout period (s) 
	// }
	// The following parameters can be defined for MongoDB: 
	// {
	// 	'SrcAuthDatabase':'admin', 
	// 	'SrcAuthFlag': "1", 
	// 	'SrcAuthMechanism':"SCRAM-SHA-1"
	// }
	// MySQL currently does not support configuring additional parameters.
	ExternParams *string `json:"ExternParams,omitempty" name:"ExternParams"`

	// Only used for "spot data consistency check". It is required if ConsistencyType is spot check
	ConsistencyParams *ConsistencyParams `json:"ConsistencyParams,omitempty" name:"ConsistencyParams"`
}

type MigrateStepDetailInfo struct {

	// Step number
	StepNo *int64 `json:"StepNo,omitempty" name:"StepNo"`

	// Step name
	StepName *string `json:"StepName,omitempty" name:"StepName"`

	// Step ID
	StepId *string `json:"StepId,omitempty" name:"StepId"`

	// Step status. Value range: 0 (default), 1 (succeeded), 2 (failed), 3 (in progress), 4 (not started)
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// Start time of current step in the format of `yyyy-mm-dd hh:mm:ss`. This field is meaningless if it does not exist or is empty
	// Note: this field may return null, indicating that no valid values can be obtained.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
}

type ModifyMigrateJobRequest struct {
	*tchttp.BaseRequest

	// ID of the data migration task to be modified
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// Data migration task name
	JobName *string `json:"JobName,omitempty" name:"JobName"`

	// Migration task configuration options
	MigrateOption *MigrateOption `json:"MigrateOption,omitempty" name:"MigrateOption"`

	// Source instance access type. Valid values: extranet (public network), cvm (CVM-based self-created instance), dcg (Direct Connect-enabled instance), vpncloud (Tencent Cloud VPN-enabled instance), cdb (TencentDB instance)
	SrcAccessType *string `json:"SrcAccessType,omitempty" name:"SrcAccessType"`

	// Source instance information, which is correlated with the migration task type
	SrcInfo *SrcInfo `json:"SrcInfo,omitempty" name:"SrcInfo"`

	// Target instance access type. Valid values: extranet (public network), cvm (CVM-based self-created instance), dcg (Direct Connect-enabled instance), vpncloud (Tencent Cloud VPN-enabled instance), cdb (TencentDB instance). Currently, only `cdb` is supported
	DstAccessType *string `json:"DstAccessType,omitempty" name:"DstAccessType"`

	// Target instance information. The region where the target instance is located cannot be modified.
	DstInfo *DstInfo `json:"DstInfo,omitempty" name:"DstInfo"`

	// When migrating the specified table, you need to set the information of the source database table to be migrated, which should be described in JSON string format. Below are examples.
	// 
	// For databases with a database-table structure:
	// [{"Database":"db1","Table":["table1","table2"]},{"Database":"db2"}]
	// For databases with a database-schema-table structure:
	// [{"Database":"db1","Schema":"s1","Table":["table1","table2"]},{"Database":"db1","Schema":"s2","Table":["table1","table2"]},{"Database":"db2","Schema":"s1","Table":["table1","table2"]},{"Database":"db3"},{"Database":"db4","Schema":"s1"}]
	// 
	// This field does not need to be set when the entire instance is to be migrated
	DatabaseInfo *string `json:"DatabaseInfo,omitempty" name:"DatabaseInfo"`
}

func (r *ModifyMigrateJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyMigrateJobRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyMigrateJobResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyMigrateJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyMigrateJobResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifySyncJobRequest struct {
	*tchttp.BaseRequest

	// ID of the disaster recovery sync task to be modified
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// Name of the disaster recovery sync task
	JobName *string `json:"JobName,omitempty" name:"JobName"`

	// Configuration options of a disaster recovery sync task
	SyncOption *SyncOption `json:"SyncOption,omitempty" name:"SyncOption"`

	// When syncing the specified table, you need to set the information of the source table to be synced, which should be described in JSON string format. Below are examples.
	// For databases with a database-table structure:
	// [{"Database":"db1","Table":["table1","table2"]},{"Database":"db2"}]
	DatabaseInfo *string `json:"DatabaseInfo,omitempty" name:"DatabaseInfo"`
}

func (r *ModifySyncJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifySyncJobRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifySyncJobResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySyncJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifySyncJobResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SrcInfo struct {

	// Alibaba Cloud AccessKey, which is applicable if the source database is an Alibaba Cloud ApsaraDB for RDS 5.6 instance
	AccessKey *string `json:"AccessKey,omitempty" name:"AccessKey"`

	// Instance IP address
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// Instance port
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// Instance username
	User *string `json:"User,omitempty" name:"User"`

	// Instance password
	Password *string `json:"Password,omitempty" name:"Password"`

	// Alibaba Cloud ApsaraDB for RDS instance ID, which is applicable if the source database is an Alibaba Cloud ApsaraDB for RDS 5.6/5.7 instance
	RdsInstanceId *string `json:"RdsInstanceId,omitempty" name:"RdsInstanceId"`

	// Short CVM instance ID in the format of `ins-olgl39y8`. It is the same as the instance ID displayed on the CVM Console page. For CVM-based self-created instances, this field needs to be passed in
	CvmInstanceId *string `json:"CvmInstanceId,omitempty" name:"CvmInstanceId"`

	// Direct Connect gateway ID in the format of dcg-0rxtqqxb
	UniqDcgId *string `json:"UniqDcgId,omitempty" name:"UniqDcgId"`

	// VPC ID in the format of vpc-92jblxto
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// VPC Subnet ID in the format of subnet-3paxmkdz
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// VPN gateway ID in the format of vpngw-9ghexg7q
	UniqVpnGwId *string `json:"UniqVpnGwId,omitempty" name:"UniqVpnGwId"`

	// Database instance ID in the format of cdb-powiqx8q
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Region name, such as ap-guangzhou
	Region *string `json:"Region,omitempty" name:"Region"`

	// For Alibaba Cloud ApsaraDB for RDS instances, enter "aliyun"; otherwise, enter "others"
	Supplier *string `json:"Supplier,omitempty" name:"Supplier"`

	// CCN instance ID, such as ccn-afp6kltc
	// Note: This field may return null, indicating that no valid values can be obtained.
	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`

	// Database version. This parameter is valid only when the instance is an RDS instance. Value: 5.6 or 5.7. Default value: 5.6
	EngineVersion *string `json:"EngineVersion,omitempty" name:"EngineVersion"`
}

type StartMigrateJobRequest struct {
	*tchttp.BaseRequest

	// Data migration task ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

func (r *StartMigrateJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StartMigrateJobRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StartMigrateJobResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StartMigrateJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StartMigrateJobResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StartSyncJobRequest struct {
	*tchttp.BaseRequest

	// Disaster recovery sync task ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

func (r *StartSyncJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StartSyncJobRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StartSyncJobResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StartSyncJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StartSyncJobResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StopMigrateJobRequest struct {
	*tchttp.BaseRequest

	// Data migration task ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

func (r *StopMigrateJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StopMigrateJobRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StopMigrateJobResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StopMigrateJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StopMigrateJobResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SyncCheckStepInfo struct {

	// Step number
	StepNo *uint64 `json:"StepNo,omitempty" name:"StepNo"`

	// Step name
	StepName *string `json:"StepName,omitempty" name:"StepName"`

	// Code of the step execution result
	StepCode *int64 `json:"StepCode,omitempty" name:"StepCode"`

	// Message about the step execution result
	StepMessage *string `json:"StepMessage,omitempty" name:"StepMessage"`
}

type SyncDetailInfo struct {

	// Total number of steps
	StepAll *int64 `json:"StepAll,omitempty" name:"StepAll"`

	// Current step
	StepNow *int64 `json:"StepNow,omitempty" name:"StepNow"`

	// Overall progress
	Progress *string `json:"Progress,omitempty" name:"Progress"`

	// Progress of the current step
	CurrentStepProgress *string `json:"CurrentStepProgress,omitempty" name:"CurrentStepProgress"`

	// Master/slave delay in MB
	MasterSlaveDistance *int64 `json:"MasterSlaveDistance,omitempty" name:"MasterSlaveDistance"`

	// Master/slave delay in seconds
	SecondsBehindMaster *int64 `json:"SecondsBehindMaster,omitempty" name:"SecondsBehindMaster"`

	// Step information
	StepInfo []*SyncStepDetailInfo `json:"StepInfo,omitempty" name:"StepInfo" list`
}

type SyncInstanceInfo struct {

	// Region name, such as ap-guangzhou
	Region *string `json:"Region,omitempty" name:"Region"`

	// Short instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type SyncJobInfo struct {

	// Disaster recovery task ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// Disaster recovery task name
	JobName *string `json:"JobName,omitempty" name:"JobName"`

	// Task sync
	SyncOption *SyncOption `json:"SyncOption,omitempty" name:"SyncOption"`

	// Source access type
	SrcAccessType *string `json:"SrcAccessType,omitempty" name:"SrcAccessType"`

	// Source data type
	SrcDatabaseType *string `json:"SrcDatabaseType,omitempty" name:"SrcDatabaseType"`

	// Source instance information
	SrcInfo *SyncInstanceInfo `json:"SrcInfo,omitempty" name:"SrcInfo"`

	// Disaster recovery access type
	DstAccessType *string `json:"DstAccessType,omitempty" name:"DstAccessType"`

	// Disaster recovery data type
	DstDatabaseType *string `json:"DstDatabaseType,omitempty" name:"DstDatabaseType"`

	// Disaster recovery instance information
	DstInfo *SyncInstanceInfo `json:"DstInfo,omitempty" name:"DstInfo"`

	// Task information
	Detail *SyncDetailInfo `json:"Detail,omitempty" name:"Detail"`

	// Task status
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// Table to be migrated
	DatabaseInfo *string `json:"DatabaseInfo,omitempty" name:"DatabaseInfo"`

	// Creation time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type SyncOption struct {

	// Sync object. 1: entire instance; 2: specified table
	SyncObject *uint64 `json:"SyncObject,omitempty" name:"SyncObject"`

	// Sync start configuration. 1: start immediately
	RunMode *uint64 `json:"RunMode,omitempty" name:"RunMode"`

	// Sync mode. 3: full + incremental sync
	SyncType *uint64 `json:"SyncType,omitempty" name:"SyncType"`

	// Data consistency check. 1: no configuration required
	ConsistencyType *uint64 `json:"ConsistencyType,omitempty" name:"ConsistencyType"`
}

type SyncStepDetailInfo struct {

	// Step number
	StepNo *uint64 `json:"StepNo,omitempty" name:"StepNo"`

	// Step name
	StepName *string `json:"StepName,omitempty" name:"StepName"`

	// Whether it can be stopped
	CanStop *int64 `json:"CanStop,omitempty" name:"CanStop"`

	// Step ID
	StepId *int64 `json:"StepId,omitempty" name:"StepId"`
}
