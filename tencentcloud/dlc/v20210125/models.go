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

package v20210125

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

// Predefined struct for user
type CancelTaskRequestParams struct {
	// Globally unique task ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

type CancelTaskRequest struct {
	*tchttp.BaseRequest
	
	// Globally unique task ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *CancelTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CancelTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelTaskResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CancelTaskResponse struct {
	*tchttp.BaseResponse
	Response *CancelTaskResponseParams `json:"Response"`
}

func (r *CancelTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Column struct {
	// Column name, which is case-insensitive and can contain up to 25 characters.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Column type. Valid values:
	// string|tinyint|smallint|int|bigint|boolean|float|double|decimal|timestamp|date|binary|array<data_type>|map<primitive_type, data_type>|struct<col_name : data_type [COMMENT col_comment], ...>|uniontype<data_type, data_type, ...>.
	Type *string `json:"Type,omitempty" name:"Type"`

	// Class comment.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// Length of the entire numeric value
	// Note: This field may return null, indicating that no valid values can be obtained.
	Precision *int64 `json:"Precision,omitempty" name:"Precision"`

	// Length of the decimal part
	// Note: This field may return null, indicating that no valid values can be obtained.
	Scale *int64 `json:"Scale,omitempty" name:"Scale"`

	// Whether the column is null.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Nullable *string `json:"Nullable,omitempty" name:"Nullable"`

	// Field position
	// Note: This field may return null, indicating that no valid values can be obtained.
	Position *int64 `json:"Position,omitempty" name:"Position"`

	// Field creation time
	// Note: This field may return null, indicating that no valid values can be obtained.
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Field modification time
	// Note: This field may return null, indicating that no valid values can be obtained.
	ModifiedTime *string `json:"ModifiedTime,omitempty" name:"ModifiedTime"`

	// Whether the column is the partition field.
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsPartition *bool `json:"IsPartition,omitempty" name:"IsPartition"`
}

// Predefined struct for user
type CreateDataEngineRequestParams struct {
	// The engine type. Valid values: `spark` and `presto`.
	EngineType *string `json:"EngineType,omitempty" name:"EngineType"`

	// The name of the virtual cluster.
	DataEngineName *string `json:"DataEngineName,omitempty" name:"DataEngineName"`

	// The cluster type. Valid values: `spark_private`, `presto_private`, `presto_cu`, and `spark_cu`.
	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`

	// The billing mode. Valid values: `0` (shared engine), `1` (pay-as-you-go), and `2` (monthly subscription).
	Mode *int64 `json:"Mode,omitempty" name:"Mode"`

	// Whether to automatically start the clusters.
	AutoResume *bool `json:"AutoResume,omitempty" name:"AutoResume"`

	// The minimum number of clusters.
	MinClusters *int64 `json:"MinClusters,omitempty" name:"MinClusters"`

	// The maximum number of clusters.
	MaxClusters *int64 `json:"MaxClusters,omitempty" name:"MaxClusters"`

	// Whether the cluster is the default one.
	DefaultDataEngine *bool `json:"DefaultDataEngine,omitempty" name:"DefaultDataEngine"`

	// The VPC CIDR block.
	CidrBlock *string `json:"CidrBlock,omitempty" name:"CidrBlock"`

	// The description.
	Message *string `json:"Message,omitempty" name:"Message"`

	// The cluster size.
	Size *int64 `json:"Size,omitempty" name:"Size"`

	// The pay mode. Valid value: `0` (postpaid, default) and `1` (prepaid) (currently not available).
	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`

	// The resource period. For the postpaid mode, the value is 3600 (default); for the prepaid mode, the value must be in the range of 1–120, representing purchasing the resource for 1–120 months.
	TimeSpan *int64 `json:"TimeSpan,omitempty" name:"TimeSpan"`

	// The unit of the resource period. Valid values: `s` (default) for the postpaid mode and `m` for the prepaid mode.
	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`

	// The auto-renewal status of the resource. For the postpaid mode, no renewal is required, and the value is fixed to `0`. For the prepaid mode, valid values are `0` (manual), `1` (auto), and `2` (no renewal). If this parameter is set to `0` for a key account in the prepaid mode, auto-renewal applies. It defaults to `0`.
	AutoRenew *int64 `json:"AutoRenew,omitempty" name:"AutoRenew"`

	// The tags to be set for the resource being created.
	Tags []*TagInfo `json:"Tags,omitempty" name:"Tags"`

	// Whether to automatically suspend clusters. Valid values: `false` (default, no) and `true` (yes).
	AutoSuspend *bool `json:"AutoSuspend,omitempty" name:"AutoSuspend"`

	// Whether to enable scheduled start and suspension of clusters. Valid values: `0` (disable) and `1` (enable). Note: This policy and the auto-suspension policy are mutually exclusive.
	CrontabResumeSuspend *int64 `json:"CrontabResumeSuspend,omitempty" name:"CrontabResumeSuspend"`

	// The complex policy for scheduled start and suspension, including the start/suspension time and suspension policy.
	CrontabResumeSuspendStrategy *CrontabResumeSuspendStrategy `json:"CrontabResumeSuspendStrategy,omitempty" name:"CrontabResumeSuspendStrategy"`

	// The type of tasks to be executed by the engine, which defaults to SQL.
	EngineExecType *string `json:"EngineExecType,omitempty" name:"EngineExecType"`

	// The max task concurrency of a cluster, which defaults to 5.
	MaxConcurrency *int64 `json:"MaxConcurrency,omitempty" name:"MaxConcurrency"`

	// The task queue time limit, which defaults to 0. When the actual queue time exceeds the value set here, scale-out may be triggered. Setting this parameter to 0 represents that scale-out may be triggered immediately after a task queues up.
	TolerableQueueTime *int64 `json:"TolerableQueueTime,omitempty" name:"TolerableQueueTime"`

	// The cluster auto-suspension time, which defaults to 10 min.
	AutoSuspendTime *int64 `json:"AutoSuspendTime,omitempty" name:"AutoSuspendTime"`

	// The resource type. Valid values: `Standard_CU` (standard) and `Memory_CU` (memory).
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`

	// The advanced configurations of clusters.
	DataEngineConfigPairs []*DataEngineConfigPair `json:"DataEngineConfigPairs,omitempty" name:"DataEngineConfigPairs"`

	// The version name of cluster image, such as SuperSQL-P 1.1 and SuperSQL-S 3.2. If no value is passed in, a cluster is created using the latest image version.
	ImageVersionName *string `json:"ImageVersionName,omitempty" name:"ImageVersionName"`

	// The name of the primary cluster.
	MainClusterName *string `json:"MainClusterName,omitempty" name:"MainClusterName"`


	ElasticSwitch *bool `json:"ElasticSwitch,omitempty" name:"ElasticSwitch"`


	ElasticLimit *int64 `json:"ElasticLimit,omitempty" name:"ElasticLimit"`
}

type CreateDataEngineRequest struct {
	*tchttp.BaseRequest
	
	// The engine type. Valid values: `spark` and `presto`.
	EngineType *string `json:"EngineType,omitempty" name:"EngineType"`

	// The name of the virtual cluster.
	DataEngineName *string `json:"DataEngineName,omitempty" name:"DataEngineName"`

	// The cluster type. Valid values: `spark_private`, `presto_private`, `presto_cu`, and `spark_cu`.
	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`

	// The billing mode. Valid values: `0` (shared engine), `1` (pay-as-you-go), and `2` (monthly subscription).
	Mode *int64 `json:"Mode,omitempty" name:"Mode"`

	// Whether to automatically start the clusters.
	AutoResume *bool `json:"AutoResume,omitempty" name:"AutoResume"`

	// The minimum number of clusters.
	MinClusters *int64 `json:"MinClusters,omitempty" name:"MinClusters"`

	// The maximum number of clusters.
	MaxClusters *int64 `json:"MaxClusters,omitempty" name:"MaxClusters"`

	// Whether the cluster is the default one.
	DefaultDataEngine *bool `json:"DefaultDataEngine,omitempty" name:"DefaultDataEngine"`

	// The VPC CIDR block.
	CidrBlock *string `json:"CidrBlock,omitempty" name:"CidrBlock"`

	// The description.
	Message *string `json:"Message,omitempty" name:"Message"`

	// The cluster size.
	Size *int64 `json:"Size,omitempty" name:"Size"`

	// The pay mode. Valid value: `0` (postpaid, default) and `1` (prepaid) (currently not available).
	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`

	// The resource period. For the postpaid mode, the value is 3600 (default); for the prepaid mode, the value must be in the range of 1–120, representing purchasing the resource for 1–120 months.
	TimeSpan *int64 `json:"TimeSpan,omitempty" name:"TimeSpan"`

	// The unit of the resource period. Valid values: `s` (default) for the postpaid mode and `m` for the prepaid mode.
	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`

	// The auto-renewal status of the resource. For the postpaid mode, no renewal is required, and the value is fixed to `0`. For the prepaid mode, valid values are `0` (manual), `1` (auto), and `2` (no renewal). If this parameter is set to `0` for a key account in the prepaid mode, auto-renewal applies. It defaults to `0`.
	AutoRenew *int64 `json:"AutoRenew,omitempty" name:"AutoRenew"`

	// The tags to be set for the resource being created.
	Tags []*TagInfo `json:"Tags,omitempty" name:"Tags"`

	// Whether to automatically suspend clusters. Valid values: `false` (default, no) and `true` (yes).
	AutoSuspend *bool `json:"AutoSuspend,omitempty" name:"AutoSuspend"`

	// Whether to enable scheduled start and suspension of clusters. Valid values: `0` (disable) and `1` (enable). Note: This policy and the auto-suspension policy are mutually exclusive.
	CrontabResumeSuspend *int64 `json:"CrontabResumeSuspend,omitempty" name:"CrontabResumeSuspend"`

	// The complex policy for scheduled start and suspension, including the start/suspension time and suspension policy.
	CrontabResumeSuspendStrategy *CrontabResumeSuspendStrategy `json:"CrontabResumeSuspendStrategy,omitempty" name:"CrontabResumeSuspendStrategy"`

	// The type of tasks to be executed by the engine, which defaults to SQL.
	EngineExecType *string `json:"EngineExecType,omitempty" name:"EngineExecType"`

	// The max task concurrency of a cluster, which defaults to 5.
	MaxConcurrency *int64 `json:"MaxConcurrency,omitempty" name:"MaxConcurrency"`

	// The task queue time limit, which defaults to 0. When the actual queue time exceeds the value set here, scale-out may be triggered. Setting this parameter to 0 represents that scale-out may be triggered immediately after a task queues up.
	TolerableQueueTime *int64 `json:"TolerableQueueTime,omitempty" name:"TolerableQueueTime"`

	// The cluster auto-suspension time, which defaults to 10 min.
	AutoSuspendTime *int64 `json:"AutoSuspendTime,omitempty" name:"AutoSuspendTime"`

	// The resource type. Valid values: `Standard_CU` (standard) and `Memory_CU` (memory).
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`

	// The advanced configurations of clusters.
	DataEngineConfigPairs []*DataEngineConfigPair `json:"DataEngineConfigPairs,omitempty" name:"DataEngineConfigPairs"`

	// The version name of cluster image, such as SuperSQL-P 1.1 and SuperSQL-S 3.2. If no value is passed in, a cluster is created using the latest image version.
	ImageVersionName *string `json:"ImageVersionName,omitempty" name:"ImageVersionName"`

	// The name of the primary cluster.
	MainClusterName *string `json:"MainClusterName,omitempty" name:"MainClusterName"`

	ElasticSwitch *bool `json:"ElasticSwitch,omitempty" name:"ElasticSwitch"`

	ElasticLimit *int64 `json:"ElasticLimit,omitempty" name:"ElasticLimit"`
}

func (r *CreateDataEngineRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDataEngineRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EngineType")
	delete(f, "DataEngineName")
	delete(f, "ClusterType")
	delete(f, "Mode")
	delete(f, "AutoResume")
	delete(f, "MinClusters")
	delete(f, "MaxClusters")
	delete(f, "DefaultDataEngine")
	delete(f, "CidrBlock")
	delete(f, "Message")
	delete(f, "Size")
	delete(f, "PayMode")
	delete(f, "TimeSpan")
	delete(f, "TimeUnit")
	delete(f, "AutoRenew")
	delete(f, "Tags")
	delete(f, "AutoSuspend")
	delete(f, "CrontabResumeSuspend")
	delete(f, "CrontabResumeSuspendStrategy")
	delete(f, "EngineExecType")
	delete(f, "MaxConcurrency")
	delete(f, "TolerableQueueTime")
	delete(f, "AutoSuspendTime")
	delete(f, "ResourceType")
	delete(f, "DataEngineConfigPairs")
	delete(f, "ImageVersionName")
	delete(f, "MainClusterName")
	delete(f, "ElasticSwitch")
	delete(f, "ElasticLimit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDataEngineRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDataEngineResponseParams struct {
	// The ID of the virtual engine.
	DataEngineId *string `json:"DataEngineId,omitempty" name:"DataEngineId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateDataEngineResponse struct {
	*tchttp.BaseResponse
	Response *CreateDataEngineResponseParams `json:"Response"`
}

func (r *CreateDataEngineResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDataEngineResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInternalTableRequestParams struct {
	// The basic table information.
	TableBaseInfo *TableBaseInfo `json:"TableBaseInfo,omitempty" name:"TableBaseInfo"`

	// The table fields.
	Columns []*TColumn `json:"Columns,omitempty" name:"Columns"`

	// The table partitions.
	Partitions []*TPartition `json:"Partitions,omitempty" name:"Partitions"`

	// The table properties.
	Properties []*Property `json:"Properties,omitempty" name:"Properties"`
}

type CreateInternalTableRequest struct {
	*tchttp.BaseRequest
	
	// The basic table information.
	TableBaseInfo *TableBaseInfo `json:"TableBaseInfo,omitempty" name:"TableBaseInfo"`

	// The table fields.
	Columns []*TColumn `json:"Columns,omitempty" name:"Columns"`

	// The table partitions.
	Partitions []*TPartition `json:"Partitions,omitempty" name:"Partitions"`

	// The table properties.
	Properties []*Property `json:"Properties,omitempty" name:"Properties"`
}

func (r *CreateInternalTableRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInternalTableRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TableBaseInfo")
	delete(f, "Columns")
	delete(f, "Partitions")
	delete(f, "Properties")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateInternalTableRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInternalTableResponseParams struct {
	// The SQL statements for creating the managed internal table.
	Execution *string `json:"Execution,omitempty" name:"Execution"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateInternalTableResponse struct {
	*tchttp.BaseResponse
	Response *CreateInternalTableResponseParams `json:"Response"`
}

func (r *CreateInternalTableResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInternalTableResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateResultDownloadRequestParams struct {
	// The result query task ID.
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// The result format.
	Format *string `json:"Format,omitempty" name:"Format"`

	// Whether to re-generate a file to download. This parameter applies only when the last task is `timeout` or `error`.
	Force *bool `json:"Force,omitempty" name:"Force"`
}

type CreateResultDownloadRequest struct {
	*tchttp.BaseRequest
	
	// The result query task ID.
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// The result format.
	Format *string `json:"Format,omitempty" name:"Format"`

	// Whether to re-generate a file to download. This parameter applies only when the last task is `timeout` or `error`.
	Force *bool `json:"Force,omitempty" name:"Force"`
}

func (r *CreateResultDownloadRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateResultDownloadRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "Format")
	delete(f, "Force")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateResultDownloadRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateResultDownloadResponseParams struct {
	// The download task ID.
	DownloadId *string `json:"DownloadId,omitempty" name:"DownloadId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateResultDownloadResponse struct {
	*tchttp.BaseResponse
	Response *CreateResultDownloadResponseParams `json:"Response"`
}

func (r *CreateResultDownloadResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateResultDownloadResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSparkAppRequestParams struct {
	// Spark application name
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 1: Spark JAR application; 2: Spark streaming application
	AppType *int64 `json:"AppType,omitempty" name:"AppType"`

	// The data engine executing the Spark job
	DataEngine *string `json:"DataEngine,omitempty" name:"DataEngine"`

	// Execution entry of the Spark application
	AppFile *string `json:"AppFile,omitempty" name:"AppFile"`

	// Execution role ID of the Spark job
	RoleArn *int64 `json:"RoleArn,omitempty" name:"RoleArn"`

	// Driver resource specification of the Spark job. Valid values: `small`, `medium`, `large`, `xlarge`.
	AppDriverSize *string `json:"AppDriverSize,omitempty" name:"AppDriverSize"`

	// Executor resource specification of the Spark job. Valid values: `small`, `medium`, `large`, `xlarge`.
	AppExecutorSize *string `json:"AppExecutorSize,omitempty" name:"AppExecutorSize"`

	// Number of Spark job executors
	AppExecutorNums *int64 `json:"AppExecutorNums,omitempty" name:"AppExecutorNums"`

	// This field has been disused. Use the `Datasource` field instead.
	Eni *string `json:"Eni,omitempty" name:"Eni"`

	// Whether it is upload locally. Valid values: `cos`, `lakefs`.
	IsLocal *string `json:"IsLocal,omitempty" name:"IsLocal"`

	// Main class of the Spark JAR job during execution
	MainClass *string `json:"MainClass,omitempty" name:"MainClass"`

	// Spark configurations separated by line break
	AppConf *string `json:"AppConf,omitempty" name:"AppConf"`

	// Whether it is upload locally. Valid values: `cos`, `lakefs`.
	IsLocalJars *string `json:"IsLocalJars,omitempty" name:"IsLocalJars"`

	// Dependency JAR packages of the Spark JAR job separated by comma
	AppJars *string `json:"AppJars,omitempty" name:"AppJars"`

	// Whether it is upload locally. Valid values: `cos`, `lakefs`.
	IsLocalFiles *string `json:"IsLocalFiles,omitempty" name:"IsLocalFiles"`

	// Dependency resources of the Spark job separated by comma
	AppFiles *string `json:"AppFiles,omitempty" name:"AppFiles"`

	// Command line parameters of the Spark job
	CmdArgs *string `json:"CmdArgs,omitempty" name:"CmdArgs"`

	// This parameter takes effect only for Spark flow tasks.
	MaxRetries *int64 `json:"MaxRetries,omitempty" name:"MaxRetries"`

	// Data source name
	DataSource *string `json:"DataSource,omitempty" name:"DataSource"`

	// PySpark: Dependency upload method. 1: cos; 2: lakefs (this method needs to be used in the console but cannot be called through APIs).
	IsLocalPythonFiles *string `json:"IsLocalPythonFiles,omitempty" name:"IsLocalPythonFiles"`

	// PySpark: Python dependency, which can be in .py, .zip, or .egg format. Multiple files should be separated by comma.
	AppPythonFiles *string `json:"AppPythonFiles,omitempty" name:"AppPythonFiles"`

	// Archives: Dependency upload method. 1: cos; 2: lakefs (this method needs to be used in the console but cannot be called through APIs).
	IsLocalArchives *string `json:"IsLocalArchives,omitempty" name:"IsLocalArchives"`

	// Archives: Dependency resources
	AppArchives *string `json:"AppArchives,omitempty" name:"AppArchives"`

	// The Spark image version.
	SparkImage *string `json:"SparkImage,omitempty" name:"SparkImage"`

	// The Spark image version name.
	SparkImageVersion *string `json:"SparkImageVersion,omitempty" name:"SparkImageVersion"`

	// The specified executor count (max), which defaults to 1. This parameter applies if the "Dynamic" mode is selected. If the "Dynamic" mode is not selected, the executor count is equal to `AppExecutorNums`.
	AppExecutorMaxNumbers *int64 `json:"AppExecutorMaxNumbers,omitempty" name:"AppExecutorMaxNumbers"`
}

type CreateSparkAppRequest struct {
	*tchttp.BaseRequest
	
	// Spark application name
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 1: Spark JAR application; 2: Spark streaming application
	AppType *int64 `json:"AppType,omitempty" name:"AppType"`

	// The data engine executing the Spark job
	DataEngine *string `json:"DataEngine,omitempty" name:"DataEngine"`

	// Execution entry of the Spark application
	AppFile *string `json:"AppFile,omitempty" name:"AppFile"`

	// Execution role ID of the Spark job
	RoleArn *int64 `json:"RoleArn,omitempty" name:"RoleArn"`

	// Driver resource specification of the Spark job. Valid values: `small`, `medium`, `large`, `xlarge`.
	AppDriverSize *string `json:"AppDriverSize,omitempty" name:"AppDriverSize"`

	// Executor resource specification of the Spark job. Valid values: `small`, `medium`, `large`, `xlarge`.
	AppExecutorSize *string `json:"AppExecutorSize,omitempty" name:"AppExecutorSize"`

	// Number of Spark job executors
	AppExecutorNums *int64 `json:"AppExecutorNums,omitempty" name:"AppExecutorNums"`

	// This field has been disused. Use the `Datasource` field instead.
	Eni *string `json:"Eni,omitempty" name:"Eni"`

	// Whether it is upload locally. Valid values: `cos`, `lakefs`.
	IsLocal *string `json:"IsLocal,omitempty" name:"IsLocal"`

	// Main class of the Spark JAR job during execution
	MainClass *string `json:"MainClass,omitempty" name:"MainClass"`

	// Spark configurations separated by line break
	AppConf *string `json:"AppConf,omitempty" name:"AppConf"`

	// Whether it is upload locally. Valid values: `cos`, `lakefs`.
	IsLocalJars *string `json:"IsLocalJars,omitempty" name:"IsLocalJars"`

	// Dependency JAR packages of the Spark JAR job separated by comma
	AppJars *string `json:"AppJars,omitempty" name:"AppJars"`

	// Whether it is upload locally. Valid values: `cos`, `lakefs`.
	IsLocalFiles *string `json:"IsLocalFiles,omitempty" name:"IsLocalFiles"`

	// Dependency resources of the Spark job separated by comma
	AppFiles *string `json:"AppFiles,omitempty" name:"AppFiles"`

	// Command line parameters of the Spark job
	CmdArgs *string `json:"CmdArgs,omitempty" name:"CmdArgs"`

	// This parameter takes effect only for Spark flow tasks.
	MaxRetries *int64 `json:"MaxRetries,omitempty" name:"MaxRetries"`

	// Data source name
	DataSource *string `json:"DataSource,omitempty" name:"DataSource"`

	// PySpark: Dependency upload method. 1: cos; 2: lakefs (this method needs to be used in the console but cannot be called through APIs).
	IsLocalPythonFiles *string `json:"IsLocalPythonFiles,omitempty" name:"IsLocalPythonFiles"`

	// PySpark: Python dependency, which can be in .py, .zip, or .egg format. Multiple files should be separated by comma.
	AppPythonFiles *string `json:"AppPythonFiles,omitempty" name:"AppPythonFiles"`

	// Archives: Dependency upload method. 1: cos; 2: lakefs (this method needs to be used in the console but cannot be called through APIs).
	IsLocalArchives *string `json:"IsLocalArchives,omitempty" name:"IsLocalArchives"`

	// Archives: Dependency resources
	AppArchives *string `json:"AppArchives,omitempty" name:"AppArchives"`

	// The Spark image version.
	SparkImage *string `json:"SparkImage,omitempty" name:"SparkImage"`

	// The Spark image version name.
	SparkImageVersion *string `json:"SparkImageVersion,omitempty" name:"SparkImageVersion"`

	// The specified executor count (max), which defaults to 1. This parameter applies if the "Dynamic" mode is selected. If the "Dynamic" mode is not selected, the executor count is equal to `AppExecutorNums`.
	AppExecutorMaxNumbers *int64 `json:"AppExecutorMaxNumbers,omitempty" name:"AppExecutorMaxNumbers"`
}

func (r *CreateSparkAppRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSparkAppRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppName")
	delete(f, "AppType")
	delete(f, "DataEngine")
	delete(f, "AppFile")
	delete(f, "RoleArn")
	delete(f, "AppDriverSize")
	delete(f, "AppExecutorSize")
	delete(f, "AppExecutorNums")
	delete(f, "Eni")
	delete(f, "IsLocal")
	delete(f, "MainClass")
	delete(f, "AppConf")
	delete(f, "IsLocalJars")
	delete(f, "AppJars")
	delete(f, "IsLocalFiles")
	delete(f, "AppFiles")
	delete(f, "CmdArgs")
	delete(f, "MaxRetries")
	delete(f, "DataSource")
	delete(f, "IsLocalPythonFiles")
	delete(f, "AppPythonFiles")
	delete(f, "IsLocalArchives")
	delete(f, "AppArchives")
	delete(f, "SparkImage")
	delete(f, "SparkImageVersion")
	delete(f, "AppExecutorMaxNumbers")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSparkAppRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSparkAppResponseParams struct {
	// The unique ID of the application.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SparkAppId *string `json:"SparkAppId,omitempty" name:"SparkAppId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateSparkAppResponse struct {
	*tchttp.BaseResponse
	Response *CreateSparkAppResponseParams `json:"Response"`
}

func (r *CreateSparkAppResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSparkAppResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSparkAppTaskRequestParams struct {
	// Spark job name
	JobName *string `json:"JobName,omitempty" name:"JobName"`

	// Command line parameters of the Spark job separated by space. They are generally used for periodic calls.
	CmdArgs *string `json:"CmdArgs,omitempty" name:"CmdArgs"`
}

type CreateSparkAppTaskRequest struct {
	*tchttp.BaseRequest
	
	// Spark job name
	JobName *string `json:"JobName,omitempty" name:"JobName"`

	// Command line parameters of the Spark job separated by space. They are generally used for periodic calls.
	CmdArgs *string `json:"CmdArgs,omitempty" name:"CmdArgs"`
}

func (r *CreateSparkAppTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSparkAppTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobName")
	delete(f, "CmdArgs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSparkAppTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSparkAppTaskResponseParams struct {
	// Batch ID
	BatchId *string `json:"BatchId,omitempty" name:"BatchId"`

	// Task ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateSparkAppTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateSparkAppTaskResponseParams `json:"Response"`
}

func (r *CreateSparkAppTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSparkAppTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTaskRequestParams struct {
	// Computing task. This parameter contains the task type and related configuration information.
	Task *Task `json:"Task,omitempty" name:"Task"`

	// Database name. If there is a database name in the SQL statement, the database in the SQL statement will be used first; otherwise, the database specified by this parameter will be used (note: when submitting the database creation SQL statement, passed in an empty string for this field).
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`

	// Name of the default data source
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitempty" name:"DatasourceConnectionName"`

	// Data engine name. If this parameter is not specified, the task will be submitted to the default engine.
	DataEngineName *string `json:"DataEngineName,omitempty" name:"DataEngineName"`
}

type CreateTaskRequest struct {
	*tchttp.BaseRequest
	
	// Computing task. This parameter contains the task type and related configuration information.
	Task *Task `json:"Task,omitempty" name:"Task"`

	// Database name. If there is a database name in the SQL statement, the database in the SQL statement will be used first; otherwise, the database specified by this parameter will be used (note: when submitting the database creation SQL statement, passed in an empty string for this field).
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`

	// Name of the default data source
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitempty" name:"DatasourceConnectionName"`

	// Data engine name. If this parameter is not specified, the task will be submitted to the default engine.
	DataEngineName *string `json:"DataEngineName,omitempty" name:"DataEngineName"`
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
	delete(f, "Task")
	delete(f, "DatabaseName")
	delete(f, "DatasourceConnectionName")
	delete(f, "DataEngineName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTaskResponseParams struct {
	// Task ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

// Predefined struct for user
type CreateTasksRequestParams struct {
	// Database name. If there is a database name in the SQL statement, the database in the SQL statement will be used first; otherwise, the database specified by this parameter will be used (note: when submitting the database creation SQL statement, passed in an empty string for this field).
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`

	// SQL task information
	Tasks *TasksInfo `json:"Tasks,omitempty" name:"Tasks"`

	// Data source name. Default value: DataLakeCatalog.
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitempty" name:"DatasourceConnectionName"`

	// Compute engine name. If this parameter is not specified, the task will be submitted to the default engine.
	DataEngineName *string `json:"DataEngineName,omitempty" name:"DataEngineName"`
}

type CreateTasksRequest struct {
	*tchttp.BaseRequest
	
	// Database name. If there is a database name in the SQL statement, the database in the SQL statement will be used first; otherwise, the database specified by this parameter will be used (note: when submitting the database creation SQL statement, passed in an empty string for this field).
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`

	// SQL task information
	Tasks *TasksInfo `json:"Tasks,omitempty" name:"Tasks"`

	// Data source name. Default value: DataLakeCatalog.
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitempty" name:"DatasourceConnectionName"`

	// Compute engine name. If this parameter is not specified, the task will be submitted to the default engine.
	DataEngineName *string `json:"DataEngineName,omitempty" name:"DataEngineName"`
}

func (r *CreateTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DatabaseName")
	delete(f, "Tasks")
	delete(f, "DatasourceConnectionName")
	delete(f, "DataEngineName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTasksResponseParams struct {
	// ID of the current batch of submitted tasks
	BatchId *string `json:"BatchId,omitempty" name:"BatchId"`

	// Collection of task IDs arranged in order of execution
	TaskIdSet []*string `json:"TaskIdSet,omitempty" name:"TaskIdSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateTasksResponse struct {
	*tchttp.BaseResponse
	Response *CreateTasksResponseParams `json:"Response"`
}

func (r *CreateTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CrontabResumeSuspendStrategy struct {
	// The scheduled start time, such as 8:00 AM every Monday.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ResumeTime *string `json:"ResumeTime,omitempty" name:"ResumeTime"`

	// The scheduled suspension time, such as 8:00 PM every Monday.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SuspendTime *string `json:"SuspendTime,omitempty" name:"SuspendTime"`

	// The suspension setting. Valid values: `0` (suspension after task end, default) and `1` (force suspension).
	// Note: This field may return null, indicating that no valid values can be obtained.
	SuspendStrategy *int64 `json:"SuspendStrategy,omitempty" name:"SuspendStrategy"`
}

type DataEngineConfigPair struct {

}

type DataGovernPolicy struct {

}

// Predefined struct for user
type DeleteSparkAppRequestParams struct {
	// Spark application name
	AppName *string `json:"AppName,omitempty" name:"AppName"`
}

type DeleteSparkAppRequest struct {
	*tchttp.BaseRequest
	
	// Spark application name
	AppName *string `json:"AppName,omitempty" name:"AppName"`
}

func (r *DeleteSparkAppRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSparkAppRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSparkAppRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSparkAppResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteSparkAppResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSparkAppResponseParams `json:"Response"`
}

func (r *DeleteSparkAppResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSparkAppResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEngineUsageInfoRequestParams struct {
	// The house ID.
	DataEngineId *string `json:"DataEngineId,omitempty" name:"DataEngineId"`
}

type DescribeEngineUsageInfoRequest struct {
	*tchttp.BaseRequest
	
	// The house ID.
	DataEngineId *string `json:"DataEngineId,omitempty" name:"DataEngineId"`
}

func (r *DescribeEngineUsageInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEngineUsageInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DataEngineId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEngineUsageInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEngineUsageInfoResponseParams struct {
	// The total cluster spec.
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// The used cluster spec.
	Used *int64 `json:"Used,omitempty" name:"Used"`

	// The available cluster spec.
	Available *int64 `json:"Available,omitempty" name:"Available"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeEngineUsageInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEngineUsageInfoResponseParams `json:"Response"`
}

func (r *DescribeEngineUsageInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEngineUsageInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeForbiddenTableProRequestParams struct {

}

type DescribeForbiddenTableProRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeForbiddenTableProRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeForbiddenTableProRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeForbiddenTableProRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeForbiddenTableProResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeForbiddenTableProResponse struct {
	*tchttp.BaseResponse
	Response *DescribeForbiddenTableProResponseParams `json:"Response"`
}

func (r *DescribeForbiddenTableProResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeForbiddenTableProResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLakeFsDirSummaryRequestParams struct {

}

type DescribeLakeFsDirSummaryRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeLakeFsDirSummaryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLakeFsDirSummaryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLakeFsDirSummaryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLakeFsDirSummaryResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeLakeFsDirSummaryResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLakeFsDirSummaryResponseParams `json:"Response"`
}

func (r *DescribeLakeFsDirSummaryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLakeFsDirSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLakeFsInfoRequestParams struct {

}

type DescribeLakeFsInfoRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeLakeFsInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLakeFsInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLakeFsInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLakeFsInfoResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeLakeFsInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLakeFsInfoResponseParams `json:"Response"`
}

func (r *DescribeLakeFsInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLakeFsInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResultDownloadRequestParams struct {
	// The query task ID.
	DownloadId *string `json:"DownloadId,omitempty" name:"DownloadId"`
}

type DescribeResultDownloadRequest struct {
	*tchttp.BaseRequest
	
	// The query task ID.
	DownloadId *string `json:"DownloadId,omitempty" name:"DownloadId"`
}

func (r *DescribeResultDownloadRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResultDownloadRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DownloadId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeResultDownloadRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResultDownloadResponseParams struct {
	// The file save path.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Path *string `json:"Path,omitempty" name:"Path"`

	// The task status. Valid values: `init`, `queue`, `format`, `compress`, `success`, `timeout`, and `error`.
	Status *string `json:"Status,omitempty" name:"Status"`

	// The task exception cause.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Reason *string `json:"Reason,omitempty" name:"Reason"`

	// The temporary secret ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SecretId *string `json:"SecretId,omitempty" name:"SecretId"`

	// The temporary secret key.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SecretKey *string `json:"SecretKey,omitempty" name:"SecretKey"`

	// The temporary token.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Token *string `json:"Token,omitempty" name:"Token"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeResultDownloadResponse struct {
	*tchttp.BaseResponse
	Response *DescribeResultDownloadResponseParams `json:"Response"`
}

func (r *DescribeResultDownloadResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResultDownloadResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSparkAppJobRequestParams struct {
	// Spark job ID. If it co-exists with `JobName`, `JobName` will become invalid.
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// Spark job name
	JobName *string `json:"JobName,omitempty" name:"JobName"`
}

type DescribeSparkAppJobRequest struct {
	*tchttp.BaseRequest
	
	// Spark job ID. If it co-exists with `JobName`, `JobName` will become invalid.
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// Spark job name
	JobName *string `json:"JobName,omitempty" name:"JobName"`
}

func (r *DescribeSparkAppJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSparkAppJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	delete(f, "JobName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSparkAppJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSparkAppJobResponseParams struct {
	// Spark job details
	// Note: This field may return null, indicating that no valid values can be obtained.
	Job *SparkJobInfo `json:"Job,omitempty" name:"Job"`

	// Whether the queried Spark job exists
	IsExists *bool `json:"IsExists,omitempty" name:"IsExists"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSparkAppJobResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSparkAppJobResponseParams `json:"Response"`
}

func (r *DescribeSparkAppJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSparkAppJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSparkAppJobsRequestParams struct {
	// The returned results are sorted by this field.
	SortBy *string `json:"SortBy,omitempty" name:"SortBy"`

	// Descending or ascending order, such as `desc`.
	Sorting *string `json:"Sorting,omitempty" name:"Sorting"`

	// Filter by this parameter, which can be `spark-job-name`.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// Update start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Update end time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Query list offset
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Query list limit
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeSparkAppJobsRequest struct {
	*tchttp.BaseRequest
	
	// The returned results are sorted by this field.
	SortBy *string `json:"SortBy,omitempty" name:"SortBy"`

	// Descending or ascending order, such as `desc`.
	Sorting *string `json:"Sorting,omitempty" name:"Sorting"`

	// Filter by this parameter, which can be `spark-job-name`.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// Update start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Update end time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Query list offset
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Query list limit
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeSparkAppJobsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSparkAppJobsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SortBy")
	delete(f, "Sorting")
	delete(f, "Filters")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSparkAppJobsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSparkAppJobsResponseParams struct {
	// Detailed list of Spark jobs
	SparkAppJobs []*SparkJobInfo `json:"SparkAppJobs,omitempty" name:"SparkAppJobs"`

	// Total number of Spark jobs
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSparkAppJobsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSparkAppJobsResponseParams `json:"Response"`
}

func (r *DescribeSparkAppJobsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSparkAppJobsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSparkAppTasksRequestParams struct {
	// Spark job ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// Paginated query offset
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Paginated query limit
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Execution instance ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// Update start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Update end time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Filter by this parameter, which can be `task-state`.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeSparkAppTasksRequest struct {
	*tchttp.BaseRequest
	
	// Spark job ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// Paginated query offset
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Paginated query limit
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Execution instance ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// Update start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Update end time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Filter by this parameter, which can be `task-state`.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeSparkAppTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSparkAppTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "TaskId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSparkAppTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSparkAppTasksResponseParams struct {
	// Task result (this field has been disused)
	// Note: This field may return null, indicating that no valid values can be obtained.
	Tasks *TaskResponseInfo `json:"Tasks,omitempty" name:"Tasks"`

	// Total number of tasks
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// List of task results
	// Note: This field may return null, indicating that no valid values can be obtained.
	SparkAppTasks []*TaskResponseInfo `json:"SparkAppTasks,omitempty" name:"SparkAppTasks"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSparkAppTasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSparkAppTasksResponseParams `json:"Response"`
}

func (r *DescribeSparkAppTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSparkAppTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskResultRequestParams struct {
	// Unique task ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// The pagination information returned by the last response. This parameter can be omitted for the first response, where the data will be returned from the beginning. The data with a volume set by the `MaxResults` field is returned each time.
	NextToken *string `json:"NextToken,omitempty" name:"NextToken"`

	// Maximum number of returned rows. Value range: 0–1,000. Default value: 1,000.
	MaxResults *int64 `json:"MaxResults,omitempty" name:"MaxResults"`
}

type DescribeTaskResultRequest struct {
	*tchttp.BaseRequest
	
	// Unique task ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// The pagination information returned by the last response. This parameter can be omitted for the first response, where the data will be returned from the beginning. The data with a volume set by the `MaxResults` field is returned each time.
	NextToken *string `json:"NextToken,omitempty" name:"NextToken"`

	// Maximum number of returned rows. Value range: 0–1,000. Default value: 1,000.
	MaxResults *int64 `json:"MaxResults,omitempty" name:"MaxResults"`
}

func (r *DescribeTaskResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "NextToken")
	delete(f, "MaxResults")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTaskResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskResultResponseParams struct {
	// The queried task information. If the returned value is empty, the task with the entered task ID does not exist. The task result will be returned only if the task status is `2` (succeeded).
	// Note: This field may return null, indicating that no valid values can be obtained.
	TaskInfo *TaskResultInfo `json:"TaskInfo,omitempty" name:"TaskInfo"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTaskResultResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTaskResultResponseParams `json:"Response"`
}

func (r *DescribeTaskResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTasksRequestParams struct {
	// Number of returned results. Default value: 10. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Filter. The following filters are supported, and the `Name` input parameter must be one of them. Up to 50 `task-id` values can be filtered, while up to 5 other parameters can be filtered in total.
	// task-id - String - (filter by task ID). `task-id` format: e386471f-139a-4e59-877f-50ece8135b99.
	// task-state - String - (filter exactly by task status). Valid values: `0` (initial), `1` (running), `2` (succeeded), `-1` (failed).
	// task-sql-keyword - String - (filter fuzzily by SQL statement keyword, such as `DROP TABLE`).
	// task-operator- string (filter by sub-UIN)
	// task-kind - string (filter by task type)
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// Sorting field. Valid values: `create-time` (default value), `update-time`.
	SortBy *string `json:"SortBy,omitempty" name:"SortBy"`

	// Sorting order. Valid values: `asc` (ascending order), `desc` (descending order). Default value: `asc`.
	Sorting *string `json:"Sorting,omitempty" name:"Sorting"`

	// Start time in the format of `yyyy-mm-dd HH:MM:SS`, which is the current time seven days ago by default.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time in the format of `yyyy-mm-dd HH:MM:SS`, which is the current time by default. The time span is (0, 30] days. Data in the last 45 days can be queried.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Filter by compute resource name
	DataEngineName *string `json:"DataEngineName,omitempty" name:"DataEngineName"`
}

type DescribeTasksRequest struct {
	*tchttp.BaseRequest
	
	// Number of returned results. Default value: 10. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Filter. The following filters are supported, and the `Name` input parameter must be one of them. Up to 50 `task-id` values can be filtered, while up to 5 other parameters can be filtered in total.
	// task-id - String - (filter by task ID). `task-id` format: e386471f-139a-4e59-877f-50ece8135b99.
	// task-state - String - (filter exactly by task status). Valid values: `0` (initial), `1` (running), `2` (succeeded), `-1` (failed).
	// task-sql-keyword - String - (filter fuzzily by SQL statement keyword, such as `DROP TABLE`).
	// task-operator- string (filter by sub-UIN)
	// task-kind - string (filter by task type)
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// Sorting field. Valid values: `create-time` (default value), `update-time`.
	SortBy *string `json:"SortBy,omitempty" name:"SortBy"`

	// Sorting order. Valid values: `asc` (ascending order), `desc` (descending order). Default value: `asc`.
	Sorting *string `json:"Sorting,omitempty" name:"Sorting"`

	// Start time in the format of `yyyy-mm-dd HH:MM:SS`, which is the current time seven days ago by default.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time in the format of `yyyy-mm-dd HH:MM:SS`, which is the current time by default. The time span is (0, 30] days. Data in the last 45 days can be queried.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Filter by compute resource name
	DataEngineName *string `json:"DataEngineName,omitempty" name:"DataEngineName"`
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
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "SortBy")
	delete(f, "Sorting")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "DataEngineName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTasksResponseParams struct {
	// List of task objects.
	TaskList []*TaskResponseInfo `json:"TaskList,omitempty" name:"TaskList"`

	// Total number of instances
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// The task overview.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TasksOverview *TasksOverview `json:"TasksOverview,omitempty" name:"TasksOverview"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTasksResponseParams `json:"Response"`
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

type Execution struct {
	// The automatically generated SQL statements.
	SQL *string `json:"SQL,omitempty" name:"SQL"`
}

type Filter struct {
	// Attribute name. If more than one filter exists, the logical relationship between these filters is `OR`.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Attribute value. If multiple values exist in one filter, the logical relationship between these values is `OR`.
	Values []*string `json:"Values,omitempty" name:"Values"`
}

// Predefined struct for user
type GenerateCreateMangedTableSqlRequestParams struct {
	// The basic table information.
	TableBaseInfo *TableBaseInfo `json:"TableBaseInfo,omitempty" name:"TableBaseInfo"`

	// The table fields.
	Columns []*TColumn `json:"Columns,omitempty" name:"Columns"`

	// The table partitions.
	Partitions []*TPartition `json:"Partitions,omitempty" name:"Partitions"`

	// The table properties.
	Properties []*Property `json:"Properties,omitempty" name:"Properties"`
}

type GenerateCreateMangedTableSqlRequest struct {
	*tchttp.BaseRequest
	
	// The basic table information.
	TableBaseInfo *TableBaseInfo `json:"TableBaseInfo,omitempty" name:"TableBaseInfo"`

	// The table fields.
	Columns []*TColumn `json:"Columns,omitempty" name:"Columns"`

	// The table partitions.
	Partitions []*TPartition `json:"Partitions,omitempty" name:"Partitions"`

	// The table properties.
	Properties []*Property `json:"Properties,omitempty" name:"Properties"`
}

func (r *GenerateCreateMangedTableSqlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GenerateCreateMangedTableSqlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TableBaseInfo")
	delete(f, "Columns")
	delete(f, "Partitions")
	delete(f, "Properties")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GenerateCreateMangedTableSqlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GenerateCreateMangedTableSqlResponseParams struct {
	// The SQL statements for creating the managed internal table.
	Execution *Execution `json:"Execution,omitempty" name:"Execution"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GenerateCreateMangedTableSqlResponse struct {
	*tchttp.BaseResponse
	Response *GenerateCreateMangedTableSqlResponseParams `json:"Response"`
}

func (r *GenerateCreateMangedTableSqlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GenerateCreateMangedTableSqlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type KVPair struct {
	// Configured key
	// Note: This field may return null, indicating that no valid values can be obtained.
	Key *string `json:"Key,omitempty" name:"Key"`

	// Configured value
	// Note: This field may return null, indicating that no valid values can be obtained.
	Value *string `json:"Value,omitempty" name:"Value"`
}

// Predefined struct for user
type ModifyGovernEventRuleRequestParams struct {

}

type ModifyGovernEventRuleRequest struct {
	*tchttp.BaseRequest
	
}

func (r *ModifyGovernEventRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyGovernEventRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyGovernEventRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyGovernEventRuleResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyGovernEventRuleResponse struct {
	*tchttp.BaseResponse
	Response *ModifyGovernEventRuleResponseParams `json:"Response"`
}

func (r *ModifyGovernEventRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyGovernEventRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySparkAppRequestParams struct {
	// Spark application name
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 1: Spark JAR application; 2: Spark streaming application
	AppType *int64 `json:"AppType,omitempty" name:"AppType"`

	// The data engine executing the Spark job
	DataEngine *string `json:"DataEngine,omitempty" name:"DataEngine"`

	// Execution entry of the Spark application
	AppFile *string `json:"AppFile,omitempty" name:"AppFile"`

	// Execution role ID of the Spark job
	RoleArn *int64 `json:"RoleArn,omitempty" name:"RoleArn"`

	// Driver resource specification of the Spark job. Valid values: `small`, `medium`, `large`, `xlarge`.
	AppDriverSize *string `json:"AppDriverSize,omitempty" name:"AppDriverSize"`

	// Executor resource specification of the Spark job. Valid values: `small`, `medium`, `large`, `xlarge`.
	AppExecutorSize *string `json:"AppExecutorSize,omitempty" name:"AppExecutorSize"`

	// Number of Spark job executors
	AppExecutorNums *int64 `json:"AppExecutorNums,omitempty" name:"AppExecutorNums"`

	// Spark application ID
	SparkAppId *string `json:"SparkAppId,omitempty" name:"SparkAppId"`

	// This field has been disused. Use the `Datasource` field instead.
	Eni *string `json:"Eni,omitempty" name:"Eni"`

	// Whether it is uploaded locally. Valid values: `cos`, `lakefs`.
	IsLocal *string `json:"IsLocal,omitempty" name:"IsLocal"`

	// Main class of the Spark JAR job during execution
	MainClass *string `json:"MainClass,omitempty" name:"MainClass"`

	// Spark configurations separated by line break
	AppConf *string `json:"AppConf,omitempty" name:"AppConf"`

	// JAR resource dependency upload method. 1: cos; 2: lakefs (this method needs to be used in the console but cannot be called through APIs).
	IsLocalJars *string `json:"IsLocalJars,omitempty" name:"IsLocalJars"`

	// Dependency JAR packages of the Spark JAR job separated by comma
	AppJars *string `json:"AppJars,omitempty" name:"AppJars"`

	// File resource dependency upload method. 1: cos; 2: lakefs (this method needs to be used in the console but cannot be called through APIs).
	IsLocalFiles *string `json:"IsLocalFiles,omitempty" name:"IsLocalFiles"`

	// Dependency resources of the Spark job separated by comma
	AppFiles *string `json:"AppFiles,omitempty" name:"AppFiles"`

	// PySpark: Dependency upload method. 1: cos; 2: lakefs (this method needs to be used in the console but cannot be called through APIs).
	IsLocalPythonFiles *string `json:"IsLocalPythonFiles,omitempty" name:"IsLocalPythonFiles"`

	// PySpark: Python dependency, which can be in .py, .zip, or .egg format. Multiple files should be separated by comma.
	AppPythonFiles *string `json:"AppPythonFiles,omitempty" name:"AppPythonFiles"`

	// Command line parameters of the Spark job
	CmdArgs *string `json:"CmdArgs,omitempty" name:"CmdArgs"`

	// This parameter takes effect only for Spark flow tasks.
	MaxRetries *int64 `json:"MaxRetries,omitempty" name:"MaxRetries"`

	// Data source name
	DataSource *string `json:"DataSource,omitempty" name:"DataSource"`

	// Archives: Dependency upload method. 1: cos; 2: lakefs (this method needs to be used in the console but cannot be called through APIs).
	IsLocalArchives *string `json:"IsLocalArchives,omitempty" name:"IsLocalArchives"`

	// Archives: Dependency resources
	AppArchives *string `json:"AppArchives,omitempty" name:"AppArchives"`

	// The Spark image version.
	SparkImage *string `json:"SparkImage,omitempty" name:"SparkImage"`

	// The Spark image version name.
	SparkImageVersion *string `json:"SparkImageVersion,omitempty" name:"SparkImageVersion"`

	// The specified executor count (max), which defaults to 1. This parameter applies if the "Dynamic" mode is selected. If the "Dynamic" mode is not selected, the executor count is equal to `AppExecutorNums`.
	AppExecutorMaxNumbers *int64 `json:"AppExecutorMaxNumbers,omitempty" name:"AppExecutorMaxNumbers"`
}

type ModifySparkAppRequest struct {
	*tchttp.BaseRequest
	
	// Spark application name
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 1: Spark JAR application; 2: Spark streaming application
	AppType *int64 `json:"AppType,omitempty" name:"AppType"`

	// The data engine executing the Spark job
	DataEngine *string `json:"DataEngine,omitempty" name:"DataEngine"`

	// Execution entry of the Spark application
	AppFile *string `json:"AppFile,omitempty" name:"AppFile"`

	// Execution role ID of the Spark job
	RoleArn *int64 `json:"RoleArn,omitempty" name:"RoleArn"`

	// Driver resource specification of the Spark job. Valid values: `small`, `medium`, `large`, `xlarge`.
	AppDriverSize *string `json:"AppDriverSize,omitempty" name:"AppDriverSize"`

	// Executor resource specification of the Spark job. Valid values: `small`, `medium`, `large`, `xlarge`.
	AppExecutorSize *string `json:"AppExecutorSize,omitempty" name:"AppExecutorSize"`

	// Number of Spark job executors
	AppExecutorNums *int64 `json:"AppExecutorNums,omitempty" name:"AppExecutorNums"`

	// Spark application ID
	SparkAppId *string `json:"SparkAppId,omitempty" name:"SparkAppId"`

	// This field has been disused. Use the `Datasource` field instead.
	Eni *string `json:"Eni,omitempty" name:"Eni"`

	// Whether it is uploaded locally. Valid values: `cos`, `lakefs`.
	IsLocal *string `json:"IsLocal,omitempty" name:"IsLocal"`

	// Main class of the Spark JAR job during execution
	MainClass *string `json:"MainClass,omitempty" name:"MainClass"`

	// Spark configurations separated by line break
	AppConf *string `json:"AppConf,omitempty" name:"AppConf"`

	// JAR resource dependency upload method. 1: cos; 2: lakefs (this method needs to be used in the console but cannot be called through APIs).
	IsLocalJars *string `json:"IsLocalJars,omitempty" name:"IsLocalJars"`

	// Dependency JAR packages of the Spark JAR job separated by comma
	AppJars *string `json:"AppJars,omitempty" name:"AppJars"`

	// File resource dependency upload method. 1: cos; 2: lakefs (this method needs to be used in the console but cannot be called through APIs).
	IsLocalFiles *string `json:"IsLocalFiles,omitempty" name:"IsLocalFiles"`

	// Dependency resources of the Spark job separated by comma
	AppFiles *string `json:"AppFiles,omitempty" name:"AppFiles"`

	// PySpark: Dependency upload method. 1: cos; 2: lakefs (this method needs to be used in the console but cannot be called through APIs).
	IsLocalPythonFiles *string `json:"IsLocalPythonFiles,omitempty" name:"IsLocalPythonFiles"`

	// PySpark: Python dependency, which can be in .py, .zip, or .egg format. Multiple files should be separated by comma.
	AppPythonFiles *string `json:"AppPythonFiles,omitempty" name:"AppPythonFiles"`

	// Command line parameters of the Spark job
	CmdArgs *string `json:"CmdArgs,omitempty" name:"CmdArgs"`

	// This parameter takes effect only for Spark flow tasks.
	MaxRetries *int64 `json:"MaxRetries,omitempty" name:"MaxRetries"`

	// Data source name
	DataSource *string `json:"DataSource,omitempty" name:"DataSource"`

	// Archives: Dependency upload method. 1: cos; 2: lakefs (this method needs to be used in the console but cannot be called through APIs).
	IsLocalArchives *string `json:"IsLocalArchives,omitempty" name:"IsLocalArchives"`

	// Archives: Dependency resources
	AppArchives *string `json:"AppArchives,omitempty" name:"AppArchives"`

	// The Spark image version.
	SparkImage *string `json:"SparkImage,omitempty" name:"SparkImage"`

	// The Spark image version name.
	SparkImageVersion *string `json:"SparkImageVersion,omitempty" name:"SparkImageVersion"`

	// The specified executor count (max), which defaults to 1. This parameter applies if the "Dynamic" mode is selected. If the "Dynamic" mode is not selected, the executor count is equal to `AppExecutorNums`.
	AppExecutorMaxNumbers *int64 `json:"AppExecutorMaxNumbers,omitempty" name:"AppExecutorMaxNumbers"`
}

func (r *ModifySparkAppRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySparkAppRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppName")
	delete(f, "AppType")
	delete(f, "DataEngine")
	delete(f, "AppFile")
	delete(f, "RoleArn")
	delete(f, "AppDriverSize")
	delete(f, "AppExecutorSize")
	delete(f, "AppExecutorNums")
	delete(f, "SparkAppId")
	delete(f, "Eni")
	delete(f, "IsLocal")
	delete(f, "MainClass")
	delete(f, "AppConf")
	delete(f, "IsLocalJars")
	delete(f, "AppJars")
	delete(f, "IsLocalFiles")
	delete(f, "AppFiles")
	delete(f, "IsLocalPythonFiles")
	delete(f, "AppPythonFiles")
	delete(f, "CmdArgs")
	delete(f, "MaxRetries")
	delete(f, "DataSource")
	delete(f, "IsLocalArchives")
	delete(f, "AppArchives")
	delete(f, "SparkImage")
	delete(f, "SparkImageVersion")
	delete(f, "AppExecutorMaxNumbers")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySparkAppRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySparkAppResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifySparkAppResponse struct {
	*tchttp.BaseResponse
	Response *ModifySparkAppResponseParams `json:"Response"`
}

func (r *ModifySparkAppResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySparkAppResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Policy struct {
	// The name of the target database. `*` represents all databases in the current catalog. To grant admin permissions, it must be `*`; to grant data connection permissions, it must be null; to grant other permissions, it can be any database.
	Database *string `json:"Database,omitempty" name:"Database"`

	// The name of the target data source. To grant admin permission, it must be `*` (all resources at this level); to grant data source and database permissions, it must be `COSDataCatalog` or `*`; to grant table permissions, it can be a custom data source; if it is left empty, `DataLakeCatalog` is used. Note: To grant permissions on a custom data source, the permissions that can be managed in the Data Lake Compute console are subsets of the account permissions granted when you connect the data source to the console.
	Catalog *string `json:"Catalog,omitempty" name:"Catalog"`

	// The name of the target table. `*` represents all tables in the current database. To grant admin permissions, it must be `*`; to grant data connection and database permissions, it must be null; to grant other permissions, it can be any table.
	Table *string `json:"Table,omitempty" name:"Table"`

	// The target permissions, which vary by permission level. Admin: `ALL` (default); data connection: `CREATE`; database: `ALL`, `CREATE`, `ALTER`, and `DROP`; table: `ALL`, `SELECT`, `INSERT`, `ALTER`, `DELETE`, `DROP`, and `UPDATE`. Note: For table permissions, if a data source other than `COSDataCatalog` is specified, only the `SELECT` permission can be granted here.
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// The permission type. Valid values: `ADMIN`, `DATASOURCE`, `DATABASE`, `TABLE`, `VIEW`, `FUNCTION`, `COLUMN`, and `ENGINE`. Note: If it is left empty, `ADMIN` is used.
	PolicyType *string `json:"PolicyType,omitempty" name:"PolicyType"`

	// The name of the target function. `*` represents all functions in the current catalog. To grant admin permissions, it must be `*`; to grant data connection permissions, it must be null; to grant other permissions, it can be any function.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Function *string `json:"Function,omitempty" name:"Function"`

	// The name of the target view. `*` represents all views in the current database. To grant admin permissions, it must be `*`; to grant data connection and database permissions, it must be null; to grant other permissions, it can be any view.
	// Note: This field may return null, indicating that no valid values can be obtained.
	View *string `json:"View,omitempty" name:"View"`

	// The name of the target column. `*` represents all columns. To grant admin permissions, it must be `*`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Column *string `json:"Column,omitempty" name:"Column"`

	// The name of the target data engine. `*` represents all engines. To grant admin permissions, it must be `*`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DataEngine *string `json:"DataEngine,omitempty" name:"DataEngine"`

	// Whether the grantee is allowed to further grant the permissions. Valid values: `false` (default) and `true` (the grantee can grant permissions gained here to other sub-users).
	// Note: This field may return null, indicating that no valid values can be obtained.
	ReAuth *bool `json:"ReAuth,omitempty" name:"ReAuth"`

	// The permission source, which is not required when input parameters are passed in. Valid values: `USER` (from the user) and `WORKGROUP` (from one or more associated work groups).
	// Note: This field may return null, indicating that no valid values can be obtained.
	Source *string `json:"Source,omitempty" name:"Source"`

	// The grant mode, which is not required as an input parameter. Valid values: `COMMON` and `SENIOR`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Mode *string `json:"Mode,omitempty" name:"Mode"`

	// The operator, which is not required as an input parameter.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Operator *string `json:"Operator,omitempty" name:"Operator"`

	// The permission policy creation time, which is not required as an input parameter.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// The ID of the work group, which applies only when the value of the `Source` field is `WORKGROUP`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SourceId *int64 `json:"SourceId,omitempty" name:"SourceId"`

	// The name of the work group, which applies only when the value of the `Source` field is `WORKGROUP`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SourceName *string `json:"SourceName,omitempty" name:"SourceName"`

	// The policy ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Id *int64 `json:"Id,omitempty" name:"Id"`
}

type Property struct {
	// The property key name.
	Key *string `json:"Key,omitempty" name:"Key"`

	// The property value.
	Value *string `json:"Value,omitempty" name:"Value"`
}

type SQLTask struct {
	// Base64-encrypted SQL statement
	SQL *string `json:"SQL,omitempty" name:"SQL"`

	// Task configuration information
	Config []*KVPair `json:"Config,omitempty" name:"Config"`
}

type SparkJobInfo struct {
	// Spark job ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// Spark job name
	JobName *string `json:"JobName,omitempty" name:"JobName"`

	// Spark job type. Valid values: `1` (batch job), `2` (streaming job).
	JobType *int64 `json:"JobType,omitempty" name:"JobType"`

	// Engine name
	DataEngine *string `json:"DataEngine,omitempty" name:"DataEngine"`

	// This field has been disused. Use the `Datasource` field instead.
	Eni *string `json:"Eni,omitempty" name:"Eni"`

	// Whether the program package is uploaded locally. Valid values: `cos`, `lakefs`.
	IsLocal *string `json:"IsLocal,omitempty" name:"IsLocal"`

	// Program package path
	JobFile *string `json:"JobFile,omitempty" name:"JobFile"`

	// Role ID
	RoleArn *int64 `json:"RoleArn,omitempty" name:"RoleArn"`

	// Main class of Spark job execution
	MainClass *string `json:"MainClass,omitempty" name:"MainClass"`

	// Command line parameters of the Spark job separated by space
	CmdArgs *string `json:"CmdArgs,omitempty" name:"CmdArgs"`

	// Native Spark configurations separated by line break
	JobConf *string `json:"JobConf,omitempty" name:"JobConf"`

	// Whether the dependency JAR packages are uploaded locally. Valid values: `cos`, `lakefs`.
	IsLocalJars *string `json:"IsLocalJars,omitempty" name:"IsLocalJars"`

	// Dependency JAR packages of the Spark job separated by comma
	JobJars *string `json:"JobJars,omitempty" name:"JobJars"`

	// Whether the dependency file is uploaded locally. Valid values: `cos`, `lakefs`.
	IsLocalFiles *string `json:"IsLocalFiles,omitempty" name:"IsLocalFiles"`

	// Dependency files of the Spark job separated by comma
	JobFiles *string `json:"JobFiles,omitempty" name:"JobFiles"`

	// Driver resource size of the Spark job
	JobDriverSize *string `json:"JobDriverSize,omitempty" name:"JobDriverSize"`

	// Executor resource size of the Spark job
	JobExecutorSize *string `json:"JobExecutorSize,omitempty" name:"JobExecutorSize"`

	// Number of Spark job executors
	JobExecutorNums *int64 `json:"JobExecutorNums,omitempty" name:"JobExecutorNums"`

	// Maximum number of retries of the Spark flow task
	JobMaxAttempts *int64 `json:"JobMaxAttempts,omitempty" name:"JobMaxAttempts"`

	// Spark job creator
	JobCreator *string `json:"JobCreator,omitempty" name:"JobCreator"`

	// Spark job creation time
	JobCreateTime *int64 `json:"JobCreateTime,omitempty" name:"JobCreateTime"`

	// Spark job update time
	JobUpdateTime *uint64 `json:"JobUpdateTime,omitempty" name:"JobUpdateTime"`

	// Last task ID of the Spark job
	CurrentTaskId *string `json:"CurrentTaskId,omitempty" name:"CurrentTaskId"`

	// Last status of the Spark job
	JobStatus *int64 `json:"JobStatus,omitempty" name:"JobStatus"`

	// Spark streaming job statistics
	// Note: This field may return null, indicating that no valid values can be obtained.
	StreamingStat *StreamingStatistics `json:"StreamingStat,omitempty" name:"StreamingStat"`

	// Data source name
	// Note: This field may return null, indicating that no valid values can be obtained.
	DataSource *string `json:"DataSource,omitempty" name:"DataSource"`

	// PySpark: Dependency upload method. 1: cos; 2: lakefs (this method needs to be used in the console but cannot be called through APIs).
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsLocalPythonFiles *string `json:"IsLocalPythonFiles,omitempty" name:"IsLocalPythonFiles"`

	// Note: This returned value has been disused.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AppPythonFiles *string `json:"AppPythonFiles,omitempty" name:"AppPythonFiles"`

	// Archives: Dependency upload method. 1: cos; 2: lakefs (this method needs to be used in the console but cannot be called through APIs).
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsLocalArchives *string `json:"IsLocalArchives,omitempty" name:"IsLocalArchives"`

	// Archives: Dependency resources
	// Note: This field may return null, indicating that no valid values can be obtained.
	JobArchives *string `json:"JobArchives,omitempty" name:"JobArchives"`

	// The Spark image version.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SparkImage *string `json:"SparkImage,omitempty" name:"SparkImage"`

	// PySpark: Python dependency, which can be in .py, .zip, or .egg format. Multiple files should be separated by comma.
	// Note: This field may return null, indicating that no valid values can be obtained.
	JobPythonFiles *string `json:"JobPythonFiles,omitempty" name:"JobPythonFiles"`

	// Number of tasks running or ready to run under the current job
	// Note: This field may return null, indicating that no valid values can be obtained.
	TaskNum *int64 `json:"TaskNum,omitempty" name:"TaskNum"`

	// Engine status. -100 (default value): unknown; -2–11: normal.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DataEngineStatus *int64 `json:"DataEngineStatus,omitempty" name:"DataEngineStatus"`

	// The specified executor count (max), which defaults to 1. This parameter applies if the "Dynamic" mode is selected. If the "Dynamic" mode is not selected, the executor count is equal to `JobExecutorNums`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	JobExecutorMaxNumbers *int64 `json:"JobExecutorMaxNumbers,omitempty" name:"JobExecutorMaxNumbers"`

	// The image version.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SparkImageVersion *string `json:"SparkImageVersion,omitempty" name:"SparkImageVersion"`
}

type StreamingStatistics struct {
	// Task start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Number of data receivers
	Receivers *int64 `json:"Receivers,omitempty" name:"Receivers"`

	// Number of receivers in service
	NumActiveReceivers *int64 `json:"NumActiveReceivers,omitempty" name:"NumActiveReceivers"`

	// Number of inactive receivers
	NumInactiveReceivers *int64 `json:"NumInactiveReceivers,omitempty" name:"NumInactiveReceivers"`

	// Number of running batches
	NumActiveBatches *int64 `json:"NumActiveBatches,omitempty" name:"NumActiveBatches"`

	// Number of batches to be processed
	NumRetainedCompletedBatches *int64 `json:"NumRetainedCompletedBatches,omitempty" name:"NumRetainedCompletedBatches"`

	// Number of completed batches
	NumTotalCompletedBatches *int64 `json:"NumTotalCompletedBatches,omitempty" name:"NumTotalCompletedBatches"`

	// Average input speed
	AverageInputRate *float64 `json:"AverageInputRate,omitempty" name:"AverageInputRate"`

	// Average queue time
	AverageSchedulingDelay *float64 `json:"AverageSchedulingDelay,omitempty" name:"AverageSchedulingDelay"`

	// Average processing time
	AverageProcessingTime *float64 `json:"AverageProcessingTime,omitempty" name:"AverageProcessingTime"`

	// Average latency
	AverageTotalDelay *float64 `json:"AverageTotalDelay,omitempty" name:"AverageTotalDelay"`
}

// Predefined struct for user
type SuspendResumeDataEngineRequestParams struct {
	// The name of a virtual cluster.
	DataEngineName *string `json:"DataEngineName,omitempty" name:"DataEngineName"`

	// The operation type: `suspend` or `resume`.
	Operate *string `json:"Operate,omitempty" name:"Operate"`
}

type SuspendResumeDataEngineRequest struct {
	*tchttp.BaseRequest
	
	// The name of a virtual cluster.
	DataEngineName *string `json:"DataEngineName,omitempty" name:"DataEngineName"`

	// The operation type: `suspend` or `resume`.
	Operate *string `json:"Operate,omitempty" name:"Operate"`
}

func (r *SuspendResumeDataEngineRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SuspendResumeDataEngineRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DataEngineName")
	delete(f, "Operate")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SuspendResumeDataEngineRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SuspendResumeDataEngineResponseParams struct {
	// The details of the virtual cluster.
	DataEngineName *string `json:"DataEngineName,omitempty" name:"DataEngineName"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SuspendResumeDataEngineResponse struct {
	*tchttp.BaseResponse
	Response *SuspendResumeDataEngineResponseParams `json:"Response"`
}

func (r *SuspendResumeDataEngineResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SuspendResumeDataEngineResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SwitchDataEngineRequestParams struct {
	// The name of the primary cluster.
	DataEngineName *string `json:"DataEngineName,omitempty" name:"DataEngineName"`

	// Whether to start the standby cluster.
	StartStandbyCluster *bool `json:"StartStandbyCluster,omitempty" name:"StartStandbyCluster"`
}

type SwitchDataEngineRequest struct {
	*tchttp.BaseRequest
	
	// The name of the primary cluster.
	DataEngineName *string `json:"DataEngineName,omitempty" name:"DataEngineName"`

	// Whether to start the standby cluster.
	StartStandbyCluster *bool `json:"StartStandbyCluster,omitempty" name:"StartStandbyCluster"`
}

func (r *SwitchDataEngineRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchDataEngineRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DataEngineName")
	delete(f, "StartStandbyCluster")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SwitchDataEngineRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SwitchDataEngineResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SwitchDataEngineResponse struct {
	*tchttp.BaseResponse
	Response *SwitchDataEngineResponseParams `json:"Response"`
}

func (r *SwitchDataEngineResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchDataEngineResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TColumn struct {
	// The field name.
	Name *string `json:"Name,omitempty" name:"Name"`

	// The field type.
	Type *string `json:"Type,omitempty" name:"Type"`

	// The field description.
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// The default field value.
	Default *string `json:"Default,omitempty" name:"Default"`

	// Whether the field is not null.
	NotNull *bool `json:"NotNull,omitempty" name:"NotNull"`
}

type TPartition struct {
	// The field name.
	Name *string `json:"Name,omitempty" name:"Name"`

	// The field type.
	Type *string `json:"Type,omitempty" name:"Type"`

	// The field description.
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// The partition type.
	PartitionType *string `json:"PartitionType,omitempty" name:"PartitionType"`

	// The partition format.
	PartitionFormat *string `json:"PartitionFormat,omitempty" name:"PartitionFormat"`

	// The separator count of the partition conversion policy.
	PartitionDot *int64 `json:"PartitionDot,omitempty" name:"PartitionDot"`

	// The partition conversion policy.
	Transform *string `json:"Transform,omitempty" name:"Transform"`

	// The policy parameters.
	TransformArgs []*string `json:"TransformArgs,omitempty" name:"TransformArgs"`
}

type TableBaseInfo struct {
	// The database name.
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`

	// The table name.
	TableName *string `json:"TableName,omitempty" name:"TableName"`

	// The data source name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitempty" name:"DatasourceConnectionName"`

	// The table remarks.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TableComment *string `json:"TableComment,omitempty" name:"TableComment"`

	// The specific type: `table` or `view`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Type *string `json:"Type,omitempty" name:"Type"`

	// The data format type, such as `hive` and `iceberg`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TableFormat *string `json:"TableFormat,omitempty" name:"TableFormat"`

	// The table creator name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	UserAlias *string `json:"UserAlias,omitempty" name:"UserAlias"`

	// The table creator ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	UserSubUin *string `json:"UserSubUin,omitempty" name:"UserSubUin"`

	// The data governance configuration.
	// Note: This field may return null, indicating that no valid values can be obtained.
	GovernPolicy *DataGovernPolicy `json:"GovernPolicy,omitempty" name:"GovernPolicy"`
}

type TagInfo struct {
	// The tag key.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// The tag value.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type Task struct {
	// SQL query task
	SQLTask *SQLTask `json:"SQLTask,omitempty" name:"SQLTask"`

	// Spark SQL query task
	SparkSQLTask *SQLTask `json:"SparkSQLTask,omitempty" name:"SparkSQLTask"`
}

type TaskResponseInfo struct {
	// Database name of the task
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`

	// Data volume of the task
	DataAmount *int64 `json:"DataAmount,omitempty" name:"DataAmount"`

	// Task ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// The compute time in ms.
	UsedTime *int64 `json:"UsedTime,omitempty" name:"UsedTime"`

	// Task output path
	OutputPath *string `json:"OutputPath,omitempty" name:"OutputPath"`

	// Task creation time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Task status. Valid values: `0` (initial), `1` (executing), `2` (executed successfully), `-1` (failed to execute), `-3` (canceled).
	State *int64 `json:"State,omitempty" name:"State"`

	// SQL statement type of the task, such as DDL and DML.
	SQLType *string `json:"SQLType,omitempty" name:"SQLType"`

	// SQL statement of the task
	SQL *string `json:"SQL,omitempty" name:"SQL"`

	// Whether the result has expired
	ResultExpired *bool `json:"ResultExpired,omitempty" name:"ResultExpired"`

	// Number of affected data rows
	RowAffectInfo *string `json:"RowAffectInfo,omitempty" name:"RowAffectInfo"`

	// Dataset of task results
	// Note: This field may return null, indicating that no valid values can be obtained.
	DataSet *string `json:"DataSet,omitempty" name:"DataSet"`

	// Failure information, such as `errorMessage`. This field has been disused.
	Error *string `json:"Error,omitempty" name:"Error"`

	// Task progress (%)
	Percentage *int64 `json:"Percentage,omitempty" name:"Percentage"`

	// Output information of task execution
	OutputMessage *string `json:"OutputMessage,omitempty" name:"OutputMessage"`

	// Type of the engine executing the SQL statement
	TaskType *string `json:"TaskType,omitempty" name:"TaskType"`

	// Task progress details
	// Note: This field may return null, indicating that no valid values can be obtained.
	ProgressDetail *string `json:"ProgressDetail,omitempty" name:"ProgressDetail"`

	// Task end time
	// Note: This field may return null, indicating that no valid values can be obtained.
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// Compute resource ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	DataEngineId *string `json:"DataEngineId,omitempty" name:"DataEngineId"`

	// Sub-UIN that executes the SQL statement
	// Note: This field may return null, indicating that no valid values can be obtained.
	OperateUin *string `json:"OperateUin,omitempty" name:"OperateUin"`

	// Compute resource name
	// Note: This field may return null, indicating that no valid values can be obtained.
	DataEngineName *string `json:"DataEngineName,omitempty" name:"DataEngineName"`

	// Whether the import type is local import or COS
	// Note: This field may return null, indicating that no valid values can be obtained.
	InputType *string `json:"InputType,omitempty" name:"InputType"`

	// Import configuration
	// Note: This field may return null, indicating that no valid values can be obtained.
	InputConf *string `json:"InputConf,omitempty" name:"InputConf"`

	// Number of data entries
	// Note: This field may return null, indicating that no valid values can be obtained.
	DataNumber *int64 `json:"DataNumber,omitempty" name:"DataNumber"`

	// Whether the data can be downloaded
	// Note: This field may return null, indicating that no valid values can be obtained.
	CanDownload *bool `json:"CanDownload,omitempty" name:"CanDownload"`

	// User alias
	// Note: This field may return null, indicating that no valid values can be obtained.
	UserAlias *string `json:"UserAlias,omitempty" name:"UserAlias"`

	// Spark application job name
	// Note: This field may return null, indicating that no valid values can be obtained.
	SparkJobName *string `json:"SparkJobName,omitempty" name:"SparkJobName"`

	// Spark application job ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	SparkJobId *string `json:"SparkJobId,omitempty" name:"SparkJobId"`

	// JAR file of the Spark application entry
	// Note: This field may return null, indicating that no valid values can be obtained.
	SparkJobFile *string `json:"SparkJobFile,omitempty" name:"SparkJobFile"`

	// Spark UI URL
	// Note: This field may return null, indicating that no valid values can be obtained.
	UiUrl *string `json:"UiUrl,omitempty" name:"UiUrl"`

	// The task time in ms.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TotalTime *int64 `json:"TotalTime,omitempty" name:"TotalTime"`

	// The program entry parameter for running a task under a Spark job.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CmdArgs *string `json:"CmdArgs,omitempty" name:"CmdArgs"`
}

type TaskResultInfo struct {
	// Unique task ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// Name of the default selected data source when the current job is executed
	// Note: This field may return null, indicating that no valid values can be obtained.
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitempty" name:"DatasourceConnectionName"`

	// Name of the default selected database when the current job is executed
	// Note: This field may return null, indicating that no valid values can be obtained.
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`

	// The currently executed SQL statement. Each task contains one SQL statement.
	SQL *string `json:"SQL,omitempty" name:"SQL"`

	// Type of the executed task. Valid values: `DDL`, `DML`, `DQL`.
	SQLType *string `json:"SQLType,omitempty" name:"SQLType"`

	// Current status of the task. `0`: initial; `1`: task running; `2`: task execution succeeded; `-1`: task execution failed; `-3`: task terminated manually by the user. The task execution result will be returned only if task execution succeeds.
	State *int64 `json:"State,omitempty" name:"State"`

	// Amount of the data scanned in bytes
	DataAmount *int64 `json:"DataAmount,omitempty" name:"DataAmount"`

	// The compute time in ms.
	UsedTime *int64 `json:"UsedTime,omitempty" name:"UsedTime"`

	// Address of the COS bucket for storing the task result
	OutputPath *string `json:"OutputPath,omitempty" name:"OutputPath"`

	// Task creation timestamp
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Task execution information. `success` will be returned if the task succeeds; otherwise, the failure cause will be returned.
	OutputMessage *string `json:"OutputMessage,omitempty" name:"OutputMessage"`

	// Number of affected rows
	RowAffectInfo *string `json:"RowAffectInfo,omitempty" name:"RowAffectInfo"`

	// Schema information of the result
	// Note: This field may return null, indicating that no valid values can be obtained.
	ResultSchema []*Column `json:"ResultSchema,omitempty" name:"ResultSchema"`

	// Result information. After it is unescaped, each element of the outer array is a data row.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ResultSet *string `json:"ResultSet,omitempty" name:"ResultSet"`

	// Pagination information. If there is no more result data, `nextToken` will be empty.
	NextToken *string `json:"NextToken,omitempty" name:"NextToken"`

	// Task progress (%)
	Percentage *int64 `json:"Percentage,omitempty" name:"Percentage"`

	// Task progress details
	ProgressDetail *string `json:"ProgressDetail,omitempty" name:"ProgressDetail"`

	// Console display format. Valid values: `table`, `text`.
	DisplayFormat *string `json:"DisplayFormat,omitempty" name:"DisplayFormat"`

	// The task time in ms.
	TotalTime *int64 `json:"TotalTime,omitempty" name:"TotalTime"`
}

type TasksInfo struct {
	// Task type. Valid values: `SQLTask` (SQL query task), `SparkSQLTask` (Spark SQL query task).
	TaskType *string `json:"TaskType,omitempty" name:"TaskType"`

	// Fault tolerance policy. `Proceed`: continues to execute subsequent tasks after the current task fails or is canceled. `Terminate`: terminates the execution of subsequent tasks after the current task fails or is canceled, and marks all subsequent tasks as canceled.
	FailureTolerance *string `json:"FailureTolerance,omitempty" name:"FailureTolerance"`

	// Base64-encrypted SQL statements separated by ";". Up to 50 tasks can be submitted at a time, and they will be executed strictly in sequence.
	SQL *string `json:"SQL,omitempty" name:"SQL"`

	// Configuration information of the task. Currently, only `SparkSQLTask` tasks are supported.
	Config []*KVPair `json:"Config,omitempty" name:"Config"`

	// User-defined parameters of the task
	Params []*KVPair `json:"Params,omitempty" name:"Params"`
}

type TasksOverview struct {
	// The number of tasks in queue.
	TaskQueuedCount *int64 `json:"TaskQueuedCount,omitempty" name:"TaskQueuedCount"`

	// The number of initialized tasks.
	TaskInitCount *int64 `json:"TaskInitCount,omitempty" name:"TaskInitCount"`

	// The number of tasks in progress.
	TaskRunningCount *int64 `json:"TaskRunningCount,omitempty" name:"TaskRunningCount"`

	// The total number of tasks in this time range.
	TotalTaskCount *int64 `json:"TotalTaskCount,omitempty" name:"TotalTaskCount"`
}

// Predefined struct for user
type UpdateRowFilterRequestParams struct {
	// The ID of the row filter policy, which can be obtained using the `DescribeUserInfo` or `DescribeWorkGroupInfo` API.
	PolicyId *int64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// The new filter policy.
	Policy *Policy `json:"Policy,omitempty" name:"Policy"`
}

type UpdateRowFilterRequest struct {
	*tchttp.BaseRequest
	
	// The ID of the row filter policy, which can be obtained using the `DescribeUserInfo` or `DescribeWorkGroupInfo` API.
	PolicyId *int64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// The new filter policy.
	Policy *Policy `json:"Policy,omitempty" name:"Policy"`
}

func (r *UpdateRowFilterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateRowFilterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PolicyId")
	delete(f, "Policy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateRowFilterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateRowFilterResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateRowFilterResponse struct {
	*tchttp.BaseResponse
	Response *UpdateRowFilterResponseParams `json:"Response"`
}

func (r *UpdateRowFilterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateRowFilterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}