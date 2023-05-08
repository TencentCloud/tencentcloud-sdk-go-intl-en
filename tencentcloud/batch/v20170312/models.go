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

package v20170312

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

type Activity struct {
	// Activity ID
	ActivityId *string `json:"ActivityId,omitempty" name:"ActivityId"`

	// Compute node ID
	ComputeNodeId *string `json:"ComputeNodeId,omitempty" name:"ComputeNodeId"`

	// Activity type. Values: `CREATE_COMPUTE_NODE`, `TERMINATE_COMPUTE_NODE`
	ComputeNodeActivityType *string `json:"ComputeNodeActivityType,omitempty" name:"ComputeNodeActivityType"`

	// Compute environment ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// Cause of the activity
	Cause *string `json:"Cause,omitempty" name:"Cause"`

	// Activity state
	ActivityState *string `json:"ActivityState,omitempty" name:"ActivityState"`

	// Reason of going to this state
	StateReason *string `json:"StateReason,omitempty" name:"StateReason"`

	// Activity start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Activity end time
	// Note: This field may return `null`, indicating that no valid value was found.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// CVM instance ID
	// Note: This field may return `null`, indicating that no valid value was found.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type AgentRunningMode struct {
	// Scenario type. Windows is supported
	Scene *string `json:"Scene,omitempty" name:"Scene"`

	// The user that runs the Agent
	User *string `json:"User,omitempty" name:"User"`

	// The session that runs the Agent
	Session *string `json:"Session,omitempty" name:"Session"`
}

type AnonymousComputeEnv struct {
	// Compute environment management type
	EnvType *string `json:"EnvType,omitempty" name:"EnvType"`

	// Compute environment's parameters
	EnvData *EnvData `json:"EnvData,omitempty" name:"EnvData"`

	// Data disk mounting option
	MountDataDisks []*MountDataDisk `json:"MountDataDisks,omitempty" name:"MountDataDisks"`

	// Agent running mode; applicable for Windows
	AgentRunningMode *AgentRunningMode `json:"AgentRunningMode,omitempty" name:"AgentRunningMode"`
}

type Application struct {
	// Task execution command
	Command *string `json:"Command,omitempty" name:"Command"`

	// Delivery form of the application. Values: `PACKAGE` (package stored in a remote location), `LOCAL` (local computer).
	DeliveryForm *string `json:"DeliveryForm,omitempty" name:"DeliveryForm"`

	// Remote storage path of the application package
	PackagePath *string `json:"PackagePath,omitempty" name:"PackagePath"`

	// Relevant configuration of the Docker used by the application. In case that the Docker configuration is used, "LOCAL" DeliveryForm means that the application software inside the Docker image is used directly and run in Docker mode; "PACKAGE" DeliveryForm means that the remote application package is run in Docker mode after being injected into the Docker image. To avoid compatibility issues with different versions of Docker, the Docker installation package and relevant dependencies are taken care of by BatchCompute. For custom images where Docker has already been installed, uninstall Docker first and then use the Docker feature.
	Docker *Docker `json:"Docker,omitempty" name:"Docker"`
}

// Predefined struct for user
type AttachInstancesRequestParams struct {
	// Compute environment ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// List of instances that added to the compute environment
	Instances []*Instance `json:"Instances,omitempty" name:"Instances"`
}

type AttachInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Compute environment ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// List of instances that added to the compute environment
	Instances []*Instance `json:"Instances,omitempty" name:"Instances"`
}

func (r *AttachInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AttachInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "Instances")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AttachInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AttachInstancesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AttachInstancesResponse struct {
	*tchttp.BaseResponse
	Response *AttachInstancesResponseParams `json:"Response"`
}

func (r *AttachInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AttachInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Authentication struct {
	// Authentication scenario such as COS
	Scene *string `json:"Scene,omitempty" name:"Scene"`

	// SecretId
	SecretId *string `json:"SecretId,omitempty" name:"SecretId"`

	// SecretKey
	SecretKey *string `json:"SecretKey,omitempty" name:"SecretKey"`
}

type ComputeEnvCreateInfo struct {
	// Compute environment ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// Compute environment name
	// Note: This field may return `null`, indicating that no valid value was found.
	EnvName *string `json:"EnvName,omitempty" name:"EnvName"`

	// Compute environment description
	// Note: This field may return `null`, indicating that no valid value was found.
	EnvDescription *string `json:"EnvDescription,omitempty" name:"EnvDescription"`

	// Compute environment type. Only "MANAGED" type is supported
	EnvType *string `json:"EnvType,omitempty" name:"EnvType"`

	// Compute environment parameter
	EnvData *EnvData `json:"EnvData,omitempty" name:"EnvData"`

	// Data disk mounting option
	// Note: This field may return `null`, indicating that no valid value was found.
	MountDataDisks []*MountDataDisk `json:"MountDataDisks,omitempty" name:"MountDataDisks"`

	// Input mapping
	// Note: This field may return `null`, indicating that no valid value was found.
	InputMappings []*InputMapping `json:"InputMappings,omitempty" name:"InputMappings"`

	// Authorization information
	// Note: This field may return `null`, indicating that no valid value was found.
	Authentications []*Authentication `json:"Authentications,omitempty" name:"Authentications"`

	// Notification information
	// Note: This field may return `null`, indicating that no valid value was found.
	Notifications []*Notification `json:"Notifications,omitempty" name:"Notifications"`

	// Number of desired compute nodes
	DesiredComputeNodeCount *uint64 `json:"DesiredComputeNodeCount,omitempty" name:"DesiredComputeNodeCount"`

	// Tag list of the compute environment.
	// Note: This field may return `null`, indicating that no valid value was found.
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

type ComputeEnvData struct {
	// List of CVM instance types
	InstanceTypes []*string `json:"InstanceTypes,omitempty" name:"InstanceTypes"`
}

type ComputeEnvView struct {
	// Compute environment ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// Compute environment name
	EnvName *string `json:"EnvName,omitempty" name:"EnvName"`

	// Location information
	Placement *Placement `json:"Placement,omitempty" name:"Placement"`

	// Creation time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Compute node statistical metrics
	ComputeNodeMetrics *ComputeNodeMetrics `json:"ComputeNodeMetrics,omitempty" name:"ComputeNodeMetrics"`

	// Compute environment type
	EnvType *string `json:"EnvType,omitempty" name:"EnvType"`

	// Number of desired compute nodes
	DesiredComputeNodeCount *uint64 `json:"DesiredComputeNodeCount,omitempty" name:"DesiredComputeNodeCount"`

	// Compute environment resource type. Values: `CVM`, `CPM` (Bare Metal)
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`

	// Next action
	NextAction *string `json:"NextAction,omitempty" name:"NextAction"`

	// Number of compute nodes added to the compute environment
	AttachedComputeNodeCount *uint64 `json:"AttachedComputeNodeCount,omitempty" name:"AttachedComputeNodeCount"`

	// Tag list bound to the compute environment.
	// Note: This field may return `null`, indicating that no valid value was found.
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

type ComputeNode struct {
	// Compute node ID
	ComputeNodeId *string `json:"ComputeNodeId,omitempty" name:"ComputeNodeId"`

	// Compute node instance ID. In a CVM scenario, this parameter is the CVM InstanceId
	ComputeNodeInstanceId *string `json:"ComputeNodeInstanceId,omitempty" name:"ComputeNodeInstanceId"`

	// Compute node state
	ComputeNodeState *string `json:"ComputeNodeState,omitempty" name:"ComputeNodeState"`

	// Number of CPU cores
	Cpu *uint64 `json:"Cpu,omitempty" name:"Cpu"`

	// Memory size in GiB
	Mem *uint64 `json:"Mem,omitempty" name:"Mem"`

	// Resource creation time
	ResourceCreatedTime *string `json:"ResourceCreatedTime,omitempty" name:"ResourceCreatedTime"`

	// Available capacity of the compute node when running TaskInstance. 0 means that the compute node is busy.
	TaskInstanceNumAvailable *uint64 `json:"TaskInstanceNumAvailable,omitempty" name:"TaskInstanceNumAvailable"`

	// BatchCompute Agent version
	AgentVersion *string `json:"AgentVersion,omitempty" name:"AgentVersion"`

	// Private IP of the instance
	PrivateIpAddresses []*string `json:"PrivateIpAddresses,omitempty" name:"PrivateIpAddresses"`

	// Public IP of the instance
	PublicIpAddresses []*string `json:"PublicIpAddresses,omitempty" name:"PublicIpAddresses"`

	// Compute environment resource type. Values: `CVM`, `CPM` (Bare Metal)
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`

	// Source of compute environment resources. <br>`BATCH_CREATED`: Instances created by BatchCompute.<br>
	// `USER_ATTACHED`: Instances added to the compute environment by the user.
	ResourceOrigin *string `json:"ResourceOrigin,omitempty" name:"ResourceOrigin"`
}

type ComputeNodeMetrics struct {
	// Number of compute nodes that have been submitted
	SubmittedCount *uint64 `json:"SubmittedCount,omitempty" name:"SubmittedCount"`

	// Number of compute nodes that are being created
	CreatingCount *uint64 `json:"CreatingCount,omitempty" name:"CreatingCount"`

	// Number of compute nodes that failed to be created
	CreationFailedCount *uint64 `json:"CreationFailedCount,omitempty" name:"CreationFailedCount"`

	// Number of compute nodes that have been created
	CreatedCount *uint64 `json:"CreatedCount,omitempty" name:"CreatedCount"`

	// Number of running compute nodes
	RunningCount *uint64 `json:"RunningCount,omitempty" name:"RunningCount"`

	// Number of compute nodes that are being terminated
	DeletingCount *uint64 `json:"DeletingCount,omitempty" name:"DeletingCount"`

	// Number of exceptional compute nodes
	AbnormalCount *uint64 `json:"AbnormalCount,omitempty" name:"AbnormalCount"`
}

// Predefined struct for user
type CreateComputeEnvRequestParams struct {
	// Compute environment information
	ComputeEnv *NamedComputeEnv `json:"ComputeEnv,omitempty" name:"ComputeEnv"`

	// Location information
	Placement *Placement `json:"Placement,omitempty" name:"Placement"`

	// The string used to guarantee the idempotency of the request, which is generated by the user and must be unique for different requests. The maximum length is 64 ASCII characters. If this parameter is not specified, the idempotency of the request cannot be guaranteed.
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`
}

type CreateComputeEnvRequest struct {
	*tchttp.BaseRequest
	
	// Compute environment information
	ComputeEnv *NamedComputeEnv `json:"ComputeEnv,omitempty" name:"ComputeEnv"`

	// Location information
	Placement *Placement `json:"Placement,omitempty" name:"Placement"`

	// The string used to guarantee the idempotency of the request, which is generated by the user and must be unique for different requests. The maximum length is 64 ASCII characters. If this parameter is not specified, the idempotency of the request cannot be guaranteed.
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`
}

func (r *CreateComputeEnvRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateComputeEnvRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ComputeEnv")
	delete(f, "Placement")
	delete(f, "ClientToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateComputeEnvRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateComputeEnvResponseParams struct {
	// Compute environment ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateComputeEnvResponse struct {
	*tchttp.BaseResponse
	Response *CreateComputeEnvResponseParams `json:"Response"`
}

func (r *CreateComputeEnvResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateComputeEnvResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTaskTemplateRequestParams struct {
	// Task template name
	TaskTemplateName *string `json:"TaskTemplateName,omitempty" name:"TaskTemplateName"`

	// Task template content with the same parameter requirements as the task
	TaskTemplateInfo *Task `json:"TaskTemplateInfo,omitempty" name:"TaskTemplateInfo"`

	// Task template description
	TaskTemplateDescription *string `json:"TaskTemplateDescription,omitempty" name:"TaskTemplateDescription"`

	// Specifies the tags you want to bind to a task template. Each task template supports up to 10 tags.
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

type CreateTaskTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Task template name
	TaskTemplateName *string `json:"TaskTemplateName,omitempty" name:"TaskTemplateName"`

	// Task template content with the same parameter requirements as the task
	TaskTemplateInfo *Task `json:"TaskTemplateInfo,omitempty" name:"TaskTemplateInfo"`

	// Task template description
	TaskTemplateDescription *string `json:"TaskTemplateDescription,omitempty" name:"TaskTemplateDescription"`

	// Specifies the tags you want to bind to a task template. Each task template supports up to 10 tags.
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

func (r *CreateTaskTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTaskTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskTemplateName")
	delete(f, "TaskTemplateInfo")
	delete(f, "TaskTemplateDescription")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTaskTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTaskTemplateResponseParams struct {
	// Task template ID
	TaskTemplateId *string `json:"TaskTemplateId,omitempty" name:"TaskTemplateId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateTaskTemplateResponse struct {
	*tchttp.BaseResponse
	Response *CreateTaskTemplateResponseParams `json:"Response"`
}

func (r *CreateTaskTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTaskTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DataDisk struct {
	// Data disk size (in GB). The minimum adjustment increment is 10 GB. The value range varies by data disk type. For more information on limits, see [Storage Overview](https://intl.cloud.tencent.com/document/product/213/4952?from_cn_redirect=1). The default value is 0, indicating that no data disk is purchased. For more information, see the product documentation.
	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`

	// Data disk type. For more information about limits on different data disk types, see [Storage Overview](https://intl.cloud.tencent.com/document/product/213/4952?from_cn_redirect=1). Valid values: <br><li>LOCAL_BASIC: local disk<br><li>LOCAL_SSD: local SSD disk<br><li>LOCAL_NVME: local NVME disk, specified in the `InstanceType`<br><li>LOCAL_PRO: local HDD disk, specified in the `InstanceType`<br><li>CLOUD_BASIC: HDD cloud disk<br><li>CLOUD_PREMIUM: Premium Cloud Storage<br><li>CLOUD_SSD: SSD<br><li>CLOUD_HSSD: Enhanced SSD<br><li>CLOUD_TSSD: Tremendous SSD<br><li>CLOUD_BSSD: Balanced SSD<br><br>Default value: LOCAL_BASIC.<br><br>This parameter is invalid for the `ResizeInstanceDisk` API.
	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`

	// Data disk ID. Note that it’s not available for `LOCAL_BASIC` and `LOCAL_SSD` disks.
	// It is only used as a response parameter for APIs such as `DescribeInstances`, and cannot be used as a request parameter for APIs such as `RunInstances`.
	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`

	// Whether to terminate the data disk when its CVM is terminated. Valid values:
	// <li>TRUE: terminate the data disk when its CVM is terminated. This value only supports pay-as-you-go cloud disks billed on an hourly basis.
	// <li>FALSE: retain the data disk when its CVM is terminated.<br>
	// Default value: TRUE<br>
	// Currently this parameter is only used in the `RunInstances` API.
	// Note: This field may return null, indicating that no valid value is found.
	DeleteWithInstance *bool `json:"DeleteWithInstance,omitempty" name:"DeleteWithInstance"`

	// Data disk snapshot ID. The size of the selected data disk snapshot must be smaller than that of the data disk.
	// Note: This field may return null, indicating that no valid value is found.
	SnapshotId *string `json:"SnapshotId,omitempty" name:"SnapshotId"`

	// Specifies whether the data disk is encrypted. Valid values: 
	// <li>TRUE: encrypted
	// <li>FALSE: not encrypted<br>
	// Default value: FALSE<br>
	// This parameter is only used with `RunInstances`.
	// Note: this field may return `null`, indicating that no valid value is obtained.
	Encrypt *bool `json:"Encrypt,omitempty" name:"Encrypt"`

	// ID of the custom CMK in the format of UUID or “kms-abcd1234”. This parameter is used to encrypt cloud disks.
	// 
	// Currently, this parameter is only used in the `RunInstances` API.
	// Note: this field may return null, indicating that no valid values can be obtained.
	KmsKeyId *string `json:"KmsKeyId,omitempty" name:"KmsKeyId"`

	// Cloud disk performance, in MB/s
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ThroughputPerformance *int64 `json:"ThroughputPerformance,omitempty" name:"ThroughputPerformance"`

	// ID of the dedicated cluster to which the instance belongs.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	CdcId *string `json:"CdcId,omitempty" name:"CdcId"`
}

// Predefined struct for user
type DeleteComputeEnvRequestParams struct {
	// Compute environment ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`
}

type DeleteComputeEnvRequest struct {
	*tchttp.BaseRequest
	
	// Compute environment ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`
}

func (r *DeleteComputeEnvRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteComputeEnvRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteComputeEnvRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteComputeEnvResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteComputeEnvResponse struct {
	*tchttp.BaseResponse
	Response *DeleteComputeEnvResponseParams `json:"Response"`
}

func (r *DeleteComputeEnvResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteComputeEnvResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteJobRequestParams struct {
	// Job ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

type DeleteJobRequest struct {
	*tchttp.BaseRequest
	
	// Job ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

func (r *DeleteJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteJobResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteJobResponse struct {
	*tchttp.BaseResponse
	Response *DeleteJobResponseParams `json:"Response"`
}

func (r *DeleteJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTaskTemplatesRequestParams struct {
	// This API is used to delete task template information.
	TaskTemplateIds []*string `json:"TaskTemplateIds,omitempty" name:"TaskTemplateIds"`
}

type DeleteTaskTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// This API is used to delete task template information.
	TaskTemplateIds []*string `json:"TaskTemplateIds,omitempty" name:"TaskTemplateIds"`
}

func (r *DeleteTaskTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTaskTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskTemplateIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTaskTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTaskTemplatesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteTaskTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteTaskTemplatesResponseParams `json:"Response"`
}

func (r *DeleteTaskTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTaskTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Dependence struct {
	// Dependency start task name 
	StartTask *string `json:"StartTask,omitempty" name:"StartTask"`

	// Dependency end task name 
	EndTask *string `json:"EndTask,omitempty" name:"EndTask"`
}

// Predefined struct for user
type DescribeAvailableCvmInstanceTypesRequestParams struct {
	// Filter.
	// <li> zone - String - Required: No - (Filter) Filter by availability zone.</li>
	// <li> instance-family - String - Required: No - (Filter) Filter by model family such as S1, I1, and M1.</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeAvailableCvmInstanceTypesRequest struct {
	*tchttp.BaseRequest
	
	// Filter.
	// <li> zone - String - Required: No - (Filter) Filter by availability zone.</li>
	// <li> instance-family - String - Required: No - (Filter) Filter by model family such as S1, I1, and M1.</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeAvailableCvmInstanceTypesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAvailableCvmInstanceTypesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAvailableCvmInstanceTypesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAvailableCvmInstanceTypesResponseParams struct {
	// Array of model configurations
	InstanceTypeConfigSet []*InstanceTypeConfig `json:"InstanceTypeConfigSet,omitempty" name:"InstanceTypeConfigSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAvailableCvmInstanceTypesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAvailableCvmInstanceTypesResponseParams `json:"Response"`
}

func (r *DescribeAvailableCvmInstanceTypesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAvailableCvmInstanceTypesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeComputeEnvActivitiesRequestParams struct {
	// Compute environment ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// Offset
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Maximum number of returned items
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Filter
	// <li> `compute-node-id` - String - Optional - Filter by the compute node ID.</li>
	Filters *Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeComputeEnvActivitiesRequest struct {
	*tchttp.BaseRequest
	
	// Compute environment ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// Offset
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Maximum number of returned items
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Filter
	// <li> `compute-node-id` - String - Optional - Filter by the compute node ID.</li>
	Filters *Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeComputeEnvActivitiesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeComputeEnvActivitiesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeComputeEnvActivitiesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeComputeEnvActivitiesResponseParams struct {
	// List of activities in the compute environment
	ActivitySet []*Activity `json:"ActivitySet,omitempty" name:"ActivitySet"`

	// Number of activities
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeComputeEnvActivitiesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeComputeEnvActivitiesResponseParams `json:"Response"`
}

func (r *DescribeComputeEnvActivitiesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeComputeEnvActivitiesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeComputeEnvCreateInfoRequestParams struct {
	// Compute environment ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`
}

type DescribeComputeEnvCreateInfoRequest struct {
	*tchttp.BaseRequest
	
	// Compute environment ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`
}

func (r *DescribeComputeEnvCreateInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeComputeEnvCreateInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeComputeEnvCreateInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeComputeEnvCreateInfoResponseParams struct {
	// Compute environment ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// Compute environment name
	EnvName *string `json:"EnvName,omitempty" name:"EnvName"`

	// Compute environment description
	// Note: This field may return `null`, indicating that no valid value was found.
	EnvDescription *string `json:"EnvDescription,omitempty" name:"EnvDescription"`

	// Compute environment type. Only `MANAGED` is supported
	EnvType *string `json:"EnvType,omitempty" name:"EnvType"`

	// Compute environment parameter
	EnvData *EnvData `json:"EnvData,omitempty" name:"EnvData"`

	// Data disk mounting option
	MountDataDisks []*MountDataDisk `json:"MountDataDisks,omitempty" name:"MountDataDisks"`

	// Input mapping
	InputMappings []*InputMapping `json:"InputMappings,omitempty" name:"InputMappings"`

	// Authorization information
	Authentications []*Authentication `json:"Authentications,omitempty" name:"Authentications"`

	// Notification information
	Notifications []*Notification `json:"Notifications,omitempty" name:"Notifications"`

	// Number of desired compute nodes
	DesiredComputeNodeCount *int64 `json:"DesiredComputeNodeCount,omitempty" name:"DesiredComputeNodeCount"`

	// Tag list bound to the compute environment.
	// Note: This field may return `null`, indicating that no valid value was found.
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeComputeEnvCreateInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeComputeEnvCreateInfoResponseParams `json:"Response"`
}

func (r *DescribeComputeEnvCreateInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeComputeEnvCreateInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeComputeEnvCreateInfosRequestParams struct {
	// List of compute environment IDs, which cannot be specified together with the `Filters` parameter.
	EnvIds []*string `json:"EnvIds,omitempty" name:"EnvIds"`

	// Filter conditions
	// <li> `zone` - String - Optional - Filter by the availability zone.</li>
	// <li> `env-id` - String - Optional - Filter by the compute environment ID.</li>
	// <li> `env-name` - String - Optional - Filter by the compute environment name.</li>
	// It cannot be specified together with `EnvIds`.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// Offset
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Maximum number of returned items
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeComputeEnvCreateInfosRequest struct {
	*tchttp.BaseRequest
	
	// List of compute environment IDs, which cannot be specified together with the `Filters` parameter.
	EnvIds []*string `json:"EnvIds,omitempty" name:"EnvIds"`

	// Filter conditions
	// <li> `zone` - String - Optional - Filter by the availability zone.</li>
	// <li> `env-id` - String - Optional - Filter by the compute environment ID.</li>
	// <li> `env-name` - String - Optional - Filter by the compute environment name.</li>
	// It cannot be specified together with `EnvIds`.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// Offset
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Maximum number of returned items
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeComputeEnvCreateInfosRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeComputeEnvCreateInfosRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeComputeEnvCreateInfosRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeComputeEnvCreateInfosResponseParams struct {
	// Number of compute environments
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// List of compute environment creation information
	ComputeEnvCreateInfoSet []*ComputeEnvCreateInfo `json:"ComputeEnvCreateInfoSet,omitempty" name:"ComputeEnvCreateInfoSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeComputeEnvCreateInfosResponse struct {
	*tchttp.BaseResponse
	Response *DescribeComputeEnvCreateInfosResponseParams `json:"Response"`
}

func (r *DescribeComputeEnvCreateInfosResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeComputeEnvCreateInfosResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeComputeEnvRequestParams struct {
	// Compute environment ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`
}

type DescribeComputeEnvRequest struct {
	*tchttp.BaseRequest
	
	// Compute environment ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`
}

func (r *DescribeComputeEnvRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeComputeEnvRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeComputeEnvRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeComputeEnvResponseParams struct {
	// Compute environment ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// Compute environment name
	EnvName *string `json:"EnvName,omitempty" name:"EnvName"`

	// Location information
	Placement *Placement `json:"Placement,omitempty" name:"Placement"`

	// Compute environment creation time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// List of compute nodes
	ComputeNodeSet []*ComputeNode `json:"ComputeNodeSet,omitempty" name:"ComputeNodeSet"`

	// Compute node statistical metrics
	ComputeNodeMetrics *ComputeNodeMetrics `json:"ComputeNodeMetrics,omitempty" name:"ComputeNodeMetrics"`

	// Number of desired compute nodes
	DesiredComputeNodeCount *uint64 `json:"DesiredComputeNodeCount,omitempty" name:"DesiredComputeNodeCount"`

	// Compute environment type
	EnvType *string `json:"EnvType,omitempty" name:"EnvType"`

	// Compute environment resource type. Values: `CVM`, `CPM` (Bare Metal)
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`

	// Next action
	NextAction *string `json:"NextAction,omitempty" name:"NextAction"`

	// Number of compute nodes added to the compute environment
	AttachedComputeNodeCount *uint64 `json:"AttachedComputeNodeCount,omitempty" name:"AttachedComputeNodeCount"`

	// Tag list bound to the compute environment.
	// Note: This field may return `null`, indicating that no valid value was found.
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeComputeEnvResponse struct {
	*tchttp.BaseResponse
	Response *DescribeComputeEnvResponseParams `json:"Response"`
}

func (r *DescribeComputeEnvResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeComputeEnvResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeComputeEnvsRequestParams struct {
	// List of compute environment IDs, which cannot be specified together with the `Filters` parameter.
	EnvIds []*string `json:"EnvIds,omitempty" name:"EnvIds"`

	// Filters
	// <li> `zone` - String - Optional - Availability zone.</li>
	// <li> `env-id` - String - Optional - Compute environment ID.</li>
	// <li> `env-name` - String - Optional - Compute environment name.</li>
	// <li> `resource-type` - String - Optional - Compute resource type (`CVM` or `CPM`).</li>
	// <li> `tag-key` - String - Optional - Tag key.</li>
	// </li>`tag-value` - String - Optional - Tag value.</li>
	// <li> `tag:tag-key` - String - Optional - Tag key-value pair. Replace `tag-key` with the actual tag key.</li>
	// It cannot be specified together with `EnvIds`.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// Offset
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Maximum number of returned items
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeComputeEnvsRequest struct {
	*tchttp.BaseRequest
	
	// List of compute environment IDs, which cannot be specified together with the `Filters` parameter.
	EnvIds []*string `json:"EnvIds,omitempty" name:"EnvIds"`

	// Filters
	// <li> `zone` - String - Optional - Availability zone.</li>
	// <li> `env-id` - String - Optional - Compute environment ID.</li>
	// <li> `env-name` - String - Optional - Compute environment name.</li>
	// <li> `resource-type` - String - Optional - Compute resource type (`CVM` or `CPM`).</li>
	// <li> `tag-key` - String - Optional - Tag key.</li>
	// </li>`tag-value` - String - Optional - Tag value.</li>
	// <li> `tag:tag-key` - String - Optional - Tag key-value pair. Replace `tag-key` with the actual tag key.</li>
	// It cannot be specified together with `EnvIds`.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// Offset
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Maximum number of returned items
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeComputeEnvsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeComputeEnvsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeComputeEnvsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeComputeEnvsResponseParams struct {
	// List of compute environments
	ComputeEnvSet []*ComputeEnvView `json:"ComputeEnvSet,omitempty" name:"ComputeEnvSet"`

	// Number of compute environments
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeComputeEnvsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeComputeEnvsResponseParams `json:"Response"`
}

func (r *DescribeComputeEnvsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeComputeEnvsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCvmZoneInstanceConfigInfosRequestParams struct {
	// Filter.
	// <li> zone - String - Required: No - (Filter) Filter by availability zone.</li>
	// <li> instance-family - String - Required: No - (Filter) Filter by model family such as S1, I1, and M1.</li>
	// <li> instance-type - String - Required: No - (Filter) Filter by model.</li>
	// <li> instance-charge-type - String - Required: No - (Filter) Filter by instance billing method. ( POSTPAID_BY_HOUR: pay-as-you-go | SPOTPAID: bidding.)  </li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeCvmZoneInstanceConfigInfosRequest struct {
	*tchttp.BaseRequest
	
	// Filter.
	// <li> zone - String - Required: No - (Filter) Filter by availability zone.</li>
	// <li> instance-family - String - Required: No - (Filter) Filter by model family such as S1, I1, and M1.</li>
	// <li> instance-type - String - Required: No - (Filter) Filter by model.</li>
	// <li> instance-charge-type - String - Required: No - (Filter) Filter by instance billing method. ( POSTPAID_BY_HOUR: pay-as-you-go | SPOTPAID: bidding.)  </li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeCvmZoneInstanceConfigInfosRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCvmZoneInstanceConfigInfosRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCvmZoneInstanceConfigInfosRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCvmZoneInstanceConfigInfosResponseParams struct {
	// List of model configurations in the availability zone.
	InstanceTypeQuotaSet []*InstanceTypeQuotaItem `json:"InstanceTypeQuotaSet,omitempty" name:"InstanceTypeQuotaSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCvmZoneInstanceConfigInfosResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCvmZoneInstanceConfigInfosResponseParams `json:"Response"`
}

func (r *DescribeCvmZoneInstanceConfigInfosResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCvmZoneInstanceConfigInfosResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceCategoriesRequestParams struct {

}

type DescribeInstanceCategoriesRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeInstanceCategoriesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceCategoriesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceCategoriesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceCategoriesResponseParams struct {
	// List of CVM instance categories
	InstanceCategorySet []*InstanceCategoryItem `json:"InstanceCategorySet,omitempty" name:"InstanceCategorySet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeInstanceCategoriesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceCategoriesResponseParams `json:"Response"`
}

func (r *DescribeInstanceCategoriesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceCategoriesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeJobRequestParams struct {
	// Job ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

type DescribeJobRequest struct {
	*tchttp.BaseRequest
	
	// Job ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

func (r *DescribeJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeJobResponseParams struct {
	// Job ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// Job name
	JobName *string `json:"JobName,omitempty" name:"JobName"`

	// Availability zone information
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// Job priority
	Priority *int64 `json:"Priority,omitempty" name:"Priority"`

	// Job state
	JobState *string `json:"JobState,omitempty" name:"JobState"`

	// Creation Date
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// End time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Task view information
	TaskSet []*TaskView `json:"TaskSet,omitempty" name:"TaskSet"`

	// Information of the dependency among tasks
	DependenceSet []*Dependence `json:"DependenceSet,omitempty" name:"DependenceSet"`

	// Task statistical metrics
	TaskMetrics *TaskMetrics `json:"TaskMetrics,omitempty" name:"TaskMetrics"`

	// Task instance statistical metrics
	TaskInstanceMetrics *TaskInstanceMetrics `json:"TaskInstanceMetrics,omitempty" name:"TaskInstanceMetrics"`

	// Job failure reason
	StateReason *string `json:"StateReason,omitempty" name:"StateReason"`

	// List of tags bound with the job.
	// Note: This field may return `null`, indicating that no valid value was found.
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// Next action
	// Note: This field may return `null`, indicating that no valid value was found.
	NextAction *string `json:"NextAction,omitempty" name:"NextAction"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeJobResponse struct {
	*tchttp.BaseResponse
	Response *DescribeJobResponseParams `json:"Response"`
}

func (r *DescribeJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeJobSubmitInfoRequestParams struct {
	// Job ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

type DescribeJobSubmitInfoRequest struct {
	*tchttp.BaseRequest
	
	// Job ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

func (r *DescribeJobSubmitInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeJobSubmitInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeJobSubmitInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeJobSubmitInfoResponseParams struct {
	// Job ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// Job name
	JobName *string `json:"JobName,omitempty" name:"JobName"`

	// Job description
	JobDescription *string `json:"JobDescription,omitempty" name:"JobDescription"`

	// Job priority. Tasks (Task) and task instances (TaskInstance) inherit the job priority
	Priority *int64 `json:"Priority,omitempty" name:"Priority"`

	// Information of tasks in the job
	Tasks []*Task `json:"Tasks,omitempty" name:"Tasks"`

	// Dependency information
	Dependences []*Dependence `json:"Dependences,omitempty" name:"Dependences"`

	// List of tags bound with the job.
	// Note: This field may return `null`, indicating that no valid value was found.
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeJobSubmitInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeJobSubmitInfoResponseParams `json:"Response"`
}

func (r *DescribeJobSubmitInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeJobSubmitInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeJobsRequestParams struct {
	// List of job IDs. It cannot be specified together with `Filters`.
	JobIds []*string `json:"JobIds,omitempty" name:"JobIds"`

	// Filter
	// <li> `job-id` - String - Optional - Filter by the job ID.</li>
	// <li> `job-name` - String - Optional - Filter by the job name.</li>
	// <li> `job-state` - String - Optional - Filter by the job state.</li>
	// <li> `zone` - String - Optional - Filter by the availability zone.</li>
	// <li> `tag-key` - String - Optional - Tag key.</li>
	// <li> `tag-value` - String - Optional - Tag value.</li>
	// <li> `tag:tag-key` - String - Optional - Filter by the tag key-value pair. The tag-key should be replaced by a specified tag key.</li>
	// It cannot be specified together with `JobIds`.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// Offset
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Maximum number of returned items
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeJobsRequest struct {
	*tchttp.BaseRequest
	
	// List of job IDs. It cannot be specified together with `Filters`.
	JobIds []*string `json:"JobIds,omitempty" name:"JobIds"`

	// Filter
	// <li> `job-id` - String - Optional - Filter by the job ID.</li>
	// <li> `job-name` - String - Optional - Filter by the job name.</li>
	// <li> `job-state` - String - Optional - Filter by the job state.</li>
	// <li> `zone` - String - Optional - Filter by the availability zone.</li>
	// <li> `tag-key` - String - Optional - Tag key.</li>
	// <li> `tag-value` - String - Optional - Tag value.</li>
	// <li> `tag:tag-key` - String - Optional - Filter by the tag key-value pair. The tag-key should be replaced by a specified tag key.</li>
	// It cannot be specified together with `JobIds`.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// Offset
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Maximum number of returned items
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeJobsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeJobsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeJobsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeJobsResponseParams struct {
	// List of jobs
	JobSet []*JobView `json:"JobSet,omitempty" name:"JobSet"`

	// Number of matched jobs
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeJobsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeJobsResponseParams `json:"Response"`
}

func (r *DescribeJobsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeJobsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskLogsRequestParams struct {
	// Job ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// Task name
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// Set of task instances
	TaskInstanceIndexes []*uint64 `json:"TaskInstanceIndexes,omitempty" name:"TaskInstanceIndexes"`

	// The start point of query
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Maximum number of task instances returned
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeTaskLogsRequest struct {
	*tchttp.BaseRequest
	
	// Job ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// Task name
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// Set of task instances
	TaskInstanceIndexes []*uint64 `json:"TaskInstanceIndexes,omitempty" name:"TaskInstanceIndexes"`

	// The start point of query
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Maximum number of task instances returned
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeTaskLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskLogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	delete(f, "TaskName")
	delete(f, "TaskInstanceIndexes")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTaskLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskLogsResponseParams struct {
	// Total number of task instances
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Set of task instance log details
	TaskInstanceLogSet []*TaskInstanceLog `json:"TaskInstanceLogSet,omitempty" name:"TaskInstanceLogSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTaskLogsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTaskLogsResponseParams `json:"Response"`
}

func (r *DescribeTaskLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskRequestParams struct {
	// Job ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// Task name
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// Offset
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned results. Default value: 100. Maximum value: 1,000.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Filter as detailed below:
	// <li> `task-instance-type` - String - Optional - Filter by the task instance state. (`SUBMITTED`, `PENDING`, `RUNNABLE`, `STARTING`, `RUNNING`, `SUCCEED`, `FAILED`, `FAILED_INTERRUPTED`).</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeTaskRequest struct {
	*tchttp.BaseRequest
	
	// Job ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// Task name
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// Offset
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned results. Default value: 100. Maximum value: 1,000.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Filter as detailed below:
	// <li> `task-instance-type` - String - Optional - Filter by the task instance state. (`SUBMITTED`, `PENDING`, `RUNNABLE`, `STARTING`, `RUNNING`, `SUCCEED`, `FAILED`, `FAILED_INTERRUPTED`).</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	delete(f, "TaskName")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskResponseParams struct {
	// Job ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// Task name
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// Task status
	TaskState *string `json:"TaskState,omitempty" name:"TaskState"`

	// Creation time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// End time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Total number of task instances
	TaskInstanceTotalCount *int64 `json:"TaskInstanceTotalCount,omitempty" name:"TaskInstanceTotalCount"`

	// Task instance information
	TaskInstanceSet []*TaskInstanceView `json:"TaskInstanceSet,omitempty" name:"TaskInstanceSet"`

	// Task instance statistical metrics
	TaskInstanceMetrics *TaskInstanceMetrics `json:"TaskInstanceMetrics,omitempty" name:"TaskInstanceMetrics"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTaskResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTaskResponseParams `json:"Response"`
}

func (r *DescribeTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskTemplatesRequestParams struct {
	// List of task template IDs. It cannot be specified together with `Filters`.
	TaskTemplateIds []*string `json:"TaskTemplateIds,omitempty" name:"TaskTemplateIds"`

	// Filter
	// <li> `task-template-name` - String - Optional - Task template name.</li>
	// <li> `tag-key` - String - Optional - Tag key.</li>
	// <li> `tag-value` - String - Optional - Tag value.</li>
	// <li> `tag:tag-key` - String - Optional - Tag key-value pair. The tag-key should be replaced by a specified tag key.</li>
	// It cannot be specified together with `TaskTemplateIds`.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// Offset
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Maximum number of returned items
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeTaskTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// List of task template IDs. It cannot be specified together with `Filters`.
	TaskTemplateIds []*string `json:"TaskTemplateIds,omitempty" name:"TaskTemplateIds"`

	// Filter
	// <li> `task-template-name` - String - Optional - Task template name.</li>
	// <li> `tag-key` - String - Optional - Tag key.</li>
	// <li> `tag-value` - String - Optional - Tag value.</li>
	// <li> `tag:tag-key` - String - Optional - Tag key-value pair. The tag-key should be replaced by a specified tag key.</li>
	// It cannot be specified together with `TaskTemplateIds`.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// Offset
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Maximum number of returned items
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeTaskTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskTemplateIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTaskTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskTemplatesResponseParams struct {
	// List of task templates
	TaskTemplateSet []*TaskTemplateView `json:"TaskTemplateSet,omitempty" name:"TaskTemplateSet"`

	// Number of task templates
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTaskTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTaskTemplatesResponseParams `json:"Response"`
}

func (r *DescribeTaskTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetachInstancesRequestParams struct {
	// Compute environment ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// List of instance IDs
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

type DetachInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Compute environment ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// List of instance IDs
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *DetachInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetachInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DetachInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetachInstancesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DetachInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DetachInstancesResponseParams `json:"Response"`
}

func (r *DetachInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetachInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Docker struct {
	// Docker Hub username or Tencent Registry username
	User *string `json:"User,omitempty" name:"User"`

	// Docker Hub password or Tencent Registry password
	Password *string `json:"Password,omitempty" name:"Password"`

	// For Docker Hub, enter "[user/repo]:[tag]"; for Tencent Registry, enter "ccr.ccs.tencentyun.com/[namespace/repo]:[tag]"
	Image *string `json:"Image,omitempty" name:"Image"`

	// For Docker Hub, this can be left blank, but please ensure public network access is present. For Tencent Registry, the server address is "ccr.ccs.tencentyun.com"
	Server *string `json:"Server,omitempty" name:"Server"`

	// Maximum retry attempts to load docket images. Range: 0 - 10. Default: `0`
	MaxRetryCount *uint64 `json:"MaxRetryCount,omitempty" name:"MaxRetryCount"`

	// Docker image loading timeout period (in seconds). Range: 1 - 360
	DelayOnRetry *uint64 `json:"DelayOnRetry,omitempty" name:"DelayOnRetry"`

	// Docker command execution parameter
	// Note: This field may return `null`, indicating that no valid value was found.
	DockerRunOption *string `json:"DockerRunOption,omitempty" name:"DockerRunOption"`
}

type EnhancedService struct {
	// Enables cloud security service. If this parameter is not specified, the cloud security service will be enabled by default.
	SecurityService *RunSecurityServiceEnabled `json:"SecurityService,omitempty" name:"SecurityService"`

	// Enables cloud monitor service. If this parameter is not specified, the cloud monitor service will be enabled by default.
	MonitorService *RunMonitorServiceEnabled `json:"MonitorService,omitempty" name:"MonitorService"`

	// Whether to enable the TAT service. If this parameter is not specified, the TAT service is enabled for public images and disabled for other images by default.
	AutomationService *RunAutomationServiceEnabled `json:"AutomationService,omitempty" name:"AutomationService"`
}

type EnvData struct {
	// CVM instance type. It cannot be specified together with `InstanceTypes` or `InstanceTypeOptions`.
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// CVM image ID
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// System disk configuration of the instance
	SystemDisk *SystemDisk `json:"SystemDisk,omitempty" name:"SystemDisk"`

	// Data disk configuration of the instance
	DataDisks []*DataDisk `json:"DataDisks,omitempty" name:"DataDisks"`

	// Information of the VPC configuration. It cannot be specified together with `Zones` and `VirtualPrivateClouds`.
	VirtualPrivateCloud *VirtualPrivateCloud `json:"VirtualPrivateCloud,omitempty" name:"VirtualPrivateCloud"`

	// Public network bandwidth configuration
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitempty" name:"InternetAccessible"`

	// CVM instance display name
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// Instance login settings
	LoginSettings *LoginSettings `json:"LoginSettings,omitempty" name:"LoginSettings"`

	// Security groups associated with the instance
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`

	// Whether to activate CWPP agent and Cloud Monitor. CWPP agent and Cloud Monitor are activated by default.
	EnhancedService *EnhancedService `json:"EnhancedService,omitempty" name:"EnhancedService"`

	// CVM instance billing method <br><li>`POSTPAID_BY_HOUR` (default): Hourly-billed pay-as-you-go <br><li>`SPOTPAID`: Spot instance <br>
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`

	// Market-related options for instances, such as parameters related to spot instances.
	InstanceMarketOptions *InstanceMarketOptionsRequest `json:"InstanceMarketOptions,omitempty" name:"InstanceMarketOptions"`

	// Types of CVM instances to create (up to 10). The system creates compute nodes of types specified in this list from top to down till the creation is successful. It cannot be specified together with `InstanceType` or `InstanceTypeOptions`. 
	InstanceTypes []*string `json:"InstanceTypes,omitempty" name:"InstanceTypes"`

	// CVM instance model configuration. It cannot be specified together with `InstanceType` or `InstanceTypes`.
	InstanceTypeOptions *InstanceTypeOptions `json:"InstanceTypeOptions,omitempty" name:"InstanceTypeOptions"`

	// List of availability zones. You can create CVM cross AZs in the same region. It cannot be specified together with `VirtualPrivateCloud` or `VirtualPrivateClouds`.
	Zones []*string `json:"Zones,omitempty" name:"Zones"`

	// List of VPCs (creation of CVM instances across VPCs is supported). It cannot be specified together with `VirtualPrivateCloud` or `Zones`.
	VirtualPrivateClouds []*VirtualPrivateCloud `json:"VirtualPrivateClouds,omitempty" name:"VirtualPrivateClouds"`
}

type EnvVar struct {
	// Environment variable name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Environment variable value
	Value *string `json:"Value,omitempty" name:"Value"`
}

type EventConfig struct {
	// Event type. Value range: <br/><li>`JOB_RUNNING`: The job is running, applicable to `SubmitJob`. </li><li>`JOB_SUCCEED`: The job succeeded, applicable to `SubmitJob`. </li><li>`JOB_FAILED`: The job failed, applicable to `SubmitJob`. </li><li>`JOB_FAILED_INTERRUPTED`: The job failed and the instance is retained, applicable to `SubmitJob`. </li><li>`TASK_RUNNING`: The task is running, applicable to `SubmitJob`. </li><li>`TASK_SUCCEED`: The task succeeded, applicable to `SubmitJob`. </li><li>`TASK_FAILED`: The task failed, applicable to `SubmitJob`. </li><li>`TASK_FAILED_INTERRUPTED`: The task failed and the instance is retained, applicable to `SubmitJob`. </li><li>`TASK_INSTANCE_RUNNING`: The task instance is running, applicable to `SubmitJob`. </li><li>`TASK_INSTANCE_SUCCEED`: The task instance succeeded, applicable to `SubmitJob`. </li><li>`TASK_INSTANCE_FAILED`: The task instance failed, applicable to `SubmitJob`. </li><li>`TASK_INSTANCE_FAILED_INTERRUPTED`: The task instance failed and the instance is retained, applicable to `SubmitJob`. </li><li>`COMPUTE_ENV_CREATED`: the compute environment has been created, applicable to "CreateComputeEnv". </li><li>`COMPUTE_ENV_DELETED`: The compute environment has been deleted, applicable to `CreateComputeEnv`. </li><li>`COMPUTE_NODE_CREATED`: The compute node has been created, applicable to `CreateComputeEnv` and `SubmitJob`. </li><li>`COMPUTE_NODE_CREATION_FAILED`: The compute node creation failed, applicable to `CreateComputeEnv` and `SubmitJob`. </li><li>`COMPUTE_NODE_RUNNING`: The compute node is running, applicable to `CreateComputeEnv` and `SubmitJob`. </li><li>`COMPUTE_NODE_ABNORMAL`: The compute node is exceptional, applicable to "CreateComputeEnv" and "SubmitJob". </li><li>`COMPUTE_NODE_DELETING`: The compute node has been deleted, applicable to `CreateComputeEnv` and `SubmitJob`. </li>
	EventName *string `json:"EventName,omitempty" name:"EventName"`

	// Custom key-value pair
	EventVars []*EventVar `json:"EventVars,omitempty" name:"EventVars"`
}

type EventVar struct {
	// Custom key
	Name *string `json:"Name,omitempty" name:"Name"`

	// Custom value
	Value *string `json:"Value,omitempty" name:"Value"`
}

type Externals struct {
	// Release address
	// Note: This field may return null, indicating that no valid value is found.
	ReleaseAddress *bool `json:"ReleaseAddress,omitempty" name:"ReleaseAddress"`

	// Not supported network. Value: <br><li>BASIC: classic network<br><li>VPC1.0: VPC1.0
	// Note: This field may return null, indicating that no valid value was found.
	UnsupportNetworks []*string `json:"UnsupportNetworks,omitempty" name:"UnsupportNetworks"`

	// Attributes of local HDD storage
	// Note: This field may return null, indicating that no valid value is found.
	StorageBlockAttr *StorageBlock `json:"StorageBlockAttr,omitempty" name:"StorageBlockAttr"`
}

type Filter struct {
	// Filters.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Filter values.
	Values []*string `json:"Values,omitempty" name:"Values"`
}

type InputMapping struct {
	// Source path
	SourcePath *string `json:"SourcePath,omitempty" name:"SourcePath"`

	// Destination path
	DestinationPath *string `json:"DestinationPath,omitempty" name:"DestinationPath"`

	// Mounting configuration item parameter
	MountOptionParameter *string `json:"MountOptionParameter,omitempty" name:"MountOptionParameter"`
}

type Instance struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Image ID
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// Instance login settings.
	LoginSettings *LoginSettings `json:"LoginSettings,omitempty" name:"LoginSettings"`
}

type InstanceCategoryItem struct {
	// Instance type name
	InstanceCategory *string `json:"InstanceCategory,omitempty" name:"InstanceCategory"`

	// List of instance families
	InstanceFamilySet []*string `json:"InstanceFamilySet,omitempty" name:"InstanceFamilySet"`
}

type InstanceMarketOptionsRequest struct {
	// Spot-related options
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	SpotOptions *SpotMarketOptions `json:"SpotOptions,omitempty" name:"SpotOptions"`

	// Market type. Valid value: `spot`.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	MarketType *string `json:"MarketType,omitempty" name:"MarketType"`
}

type InstanceTypeConfig struct {
	// Memory size in GB.
	Mem *int64 `json:"Mem,omitempty" name:"Mem"`

	// Number of CPU cores.
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

	// Instance model.
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// Availability zone.
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// Instance model family.
	InstanceFamily *string `json:"InstanceFamily,omitempty" name:"InstanceFamily"`
}

type InstanceTypeOptions struct {
	// Number of CPU cores
	CPU *uint64 `json:"CPU,omitempty" name:"CPU"`

	// Memory size in GB.
	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`

	// Instance model category. Values: `ALL` (default), `GENERAL`, `GENERAL_2`, `GENERAL_3`, `COMPUTE`, `COMPUTE_2`, and `COMPUTE_3`. 
	InstanceCategories []*string `json:"InstanceCategories,omitempty" name:"InstanceCategories"`
}

type InstanceTypeQuotaItem struct {
	// Availability zone.
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// Instance model.
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// Instance billing plan. Valid values: <br><li>POSTPAID_BY_HOUR: pay after use. You are billed for your traffic by the hour.<br><li>`CDHPAID`: [`CDH`](https://intl.cloud.tencent.com/document/product/416?from_cn_redirect=1) billing plan. Applicable to `CDH` only, not the instances on the host.
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`

	// ENI type. For example, 25 represents an ENI of 25 GB.
	NetworkCard *int64 `json:"NetworkCard,omitempty" name:"NetworkCard"`

	// Additional data.
	// Note: This field may return null, indicating that no valid value is found.
	Externals *Externals `json:"Externals,omitempty" name:"Externals"`

	// Number of CPU cores of an instance model.
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

	// Instance memory capacity; unit: `GB`.
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// Instance model family.
	InstanceFamily *string `json:"InstanceFamily,omitempty" name:"InstanceFamily"`

	// Model name.
	TypeName *string `json:"TypeName,omitempty" name:"TypeName"`

	// List of local disk specifications. If the parameter returns null, it means that local disks cannot be created.
	LocalDiskTypeList []*LocalDiskType `json:"LocalDiskTypeList,omitempty" name:"LocalDiskTypeList"`

	// Whether an instance model is available. Valid values: <br><li>SELL: available <br><li>SOLD_OUT: sold out
	Status *string `json:"Status,omitempty" name:"Status"`

	// Price of an instance model.
	Price *ItemPrice `json:"Price,omitempty" name:"Price"`

	// Details of out-of-stock items
	// Note: this field may return null, indicating that no valid value is obtained.
	SoldOutReason *string `json:"SoldOutReason,omitempty" name:"SoldOutReason"`

	// Private network bandwidth, in Gbps.
	InstanceBandwidth *float64 `json:"InstanceBandwidth,omitempty" name:"InstanceBandwidth"`

	// The max packet sending and receiving capability (in 10k PPS).
	InstancePps *int64 `json:"InstancePps,omitempty" name:"InstancePps"`

	// Number of local storage blocks.
	StorageBlockAmount *int64 `json:"StorageBlockAmount,omitempty" name:"StorageBlockAmount"`

	// CPU type.
	CpuType *string `json:"CpuType,omitempty" name:"CpuType"`

	// Number of GPUs of the instance.
	Gpu *int64 `json:"Gpu,omitempty" name:"Gpu"`

	// Number of FPGAs of the instance.
	Fpga *int64 `json:"Fpga,omitempty" name:"Fpga"`

	// Descriptive information of the instance.
	Remark *string `json:"Remark,omitempty" name:"Remark"`


	GpuCount *float64 `json:"GpuCount,omitempty" name:"GpuCount"`

	// CPU clock rate of the instance
	Frequency *string `json:"Frequency,omitempty" name:"Frequency"`
}

type InternetAccessible struct {
	// Network connection billing plan. Valid value: <br><li>TRAFFIC_POSTPAID_BY_HOUR: pay after use. You are billed for your traffic, by the hour.
	InternetChargeType *string `json:"InternetChargeType,omitempty" name:"InternetChargeType"`

	// The maximum outbound bandwidth of the public network, in Mbps. The default value is 0 Mbps. The upper limit of bandwidth varies for different models. For more information, see [Purchase Network Bandwidth](https://intl.cloud.tencent.com/document/product/213/12523?from_cn_redirect=1).
	InternetMaxBandwidthOut *int64 `json:"InternetMaxBandwidthOut,omitempty" name:"InternetMaxBandwidthOut"`

	// Whether to assign a public IP. Valid values: <br><li>TRUE: Assign a public IP <br><li>FALSE: Do not assign a public IP <br><br>If the public network bandwidth is greater than 0 Mbps, you can choose whether to assign a public IP; by default a public IP will be assigned. If the public network bandwidth is 0 Mbps, you will not be able to assign a public IP.
	PublicIpAssigned *bool `json:"PublicIpAssigned,omitempty" name:"PublicIpAssigned"`

	// Bandwidth package ID. To obatin the IDs, you can call [`DescribeBandwidthPackages`](https://intl.cloud.tencent.com/document/api/215/19209?from_cn_redirect=1) and look for the `BandwidthPackageId` fields in the response.
	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" name:"BandwidthPackageId"`
}

type ItemPrice struct {
	// The original unit price for pay-as-you-go mode in USD. <br><li>When a billing tier is returned, it indicates the price fo the returned billing tier. For example, if `UnitPriceSecondStep` is returned, it refers to the unit price for the usage between 0 to 96 hours. Otherwise, it refers to that the unit price for unlimited usage.
	// Note: this field may return null, indicating that no valid value is obtained.
	UnitPrice *float64 `json:"UnitPrice,omitempty" name:"UnitPrice"`

	// Billing unit for pay-as-you-go mode. Valid values: <br><li>HOUR: billed on an hourly basis. It's used for hourly postpaid instances (`POSTPAID_BY_HOUR`). <br><li>GB: bill by traffic in GB. It's used for postpaid products that are billed by the hourly traffic (`TRAFFIC_POSTPAID_BY_HOUR`).
	// Note: this field may return null, indicating that no valid value is obtained.
	ChargeUnit *string `json:"ChargeUnit,omitempty" name:"ChargeUnit"`

	// The original price of a pay-in-advance instance, in USD.
	// Note: this field may return null, indicating that no valid value is obtained.
	OriginalPrice *float64 `json:"OriginalPrice,omitempty" name:"OriginalPrice"`

	// Discount price of a prepaid instance, in USD.
	// Note: this field may return null, indicating that no valid value is obtained.
	DiscountPrice *float64 `json:"DiscountPrice,omitempty" name:"DiscountPrice"`

	// Percentage of the original price. For example, if you enter "20.0", the discounted price will be 20% of the original price.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Discount *float64 `json:"Discount,omitempty" name:"Discount"`

	// The discounted unit price for pay-as-you-go mode in USD. <br><li>When a billing tier is returned, it indicates the price fo the returned billing tier. For example, if `UnitPriceSecondStep` is returned, it refers to the unit price for the usage between 0 to 96 hours. Otherwise, it refers to that the unit price for unlimited usage.
	// Note: this field may return null, indicating that no valid value is obtained.
	UnitPriceDiscount *float64 `json:"UnitPriceDiscount,omitempty" name:"UnitPriceDiscount"`

	// Original unit price for the usage between 96 to 360 hours in USD. It's applicable to pay-as-you-go mode.
	// Note: this field may return null, indicating that no valid value is obtained.
	UnitPriceSecondStep *float64 `json:"UnitPriceSecondStep,omitempty" name:"UnitPriceSecondStep"`

	// Discounted unit price for the usage between 96 to 360 hours in USD. It's applicable to pay-as-you-go mode.
	// Note: this field may return null, indicating that no valid value is obtained.
	UnitPriceDiscountSecondStep *float64 `json:"UnitPriceDiscountSecondStep,omitempty" name:"UnitPriceDiscountSecondStep"`

	// Original unit price for the usage after 360 hours in USD. It's applicable to pay-as-you-go mode.
	// Note: this field may return null, indicating that no valid value is obtained.
	UnitPriceThirdStep *float64 `json:"UnitPriceThirdStep,omitempty" name:"UnitPriceThirdStep"`

	// Discounted unit price for the usage after 360 hours in USD. It's applicable to pay-as-you-go mode.
	// Note: this field may return null, indicating that no valid value is obtained.
	UnitPriceDiscountThirdStep *float64 `json:"UnitPriceDiscountThirdStep,omitempty" name:"UnitPriceDiscountThirdStep"`

	// Original 3-year payment, in USD. This parameter is only available to upfront payment mode.
	// Note: this field may return `null`, indicating that no valid value was found.
	// Note: this field may return `null`, indicating that no valid value was found.
	OriginalPriceThreeYear *float64 `json:"OriginalPriceThreeYear,omitempty" name:"OriginalPriceThreeYear"`

	// Discounted 3-year upfront payment, in USD. This parameter is only available to upfront payment mode.
	// Note: this field may return `null`, indicating that no valid value was found.
	// Note: this field may return `null`, indicating that no valid value was found.
	DiscountPriceThreeYear *float64 `json:"DiscountPriceThreeYear,omitempty" name:"DiscountPriceThreeYear"`

	// Discount for 3-year upfront payment. For example, 20.0 indicates 80% off.
	// Note: this field may return `null`, indicating that no valid value was found.
	// Note: this field may return `null`, indicating that no valid value was found.
	DiscountThreeYear *float64 `json:"DiscountThreeYear,omitempty" name:"DiscountThreeYear"`

	// Original 5-year payment, in USD. This parameter is only available to upfront payment mode.
	// Note: this field may return `null`, indicating that no valid value was found.
	// Note: this field may return `null`, indicating that no valid value was found.
	OriginalPriceFiveYear *float64 `json:"OriginalPriceFiveYear,omitempty" name:"OriginalPriceFiveYear"`

	// Discounted 5-year upfront payment, in USD. This parameter is only available to upfront payment mode.
	// Note: this field may return `null`, indicating that no valid value was found.
	// Note: this field may return `null`, indicating that no valid value was found.
	DiscountPriceFiveYear *float64 `json:"DiscountPriceFiveYear,omitempty" name:"DiscountPriceFiveYear"`

	// Discount for 5-year upfront payment. For example, 20.0 indicates 80% off.
	// Note: this field may return `null`, indicating that no valid value was found.
	// Note: this field may return `null`, indicating that no valid value was found.
	DiscountFiveYear *float64 `json:"DiscountFiveYear,omitempty" name:"DiscountFiveYear"`

	// Original 1-year payment, in USD. This parameter is only available to upfront payment mode.
	// Note: this field may return `null`, indicating that no valid value was found.
	// Note: this field may return `null`, indicating that no valid value was found.
	OriginalPriceOneYear *float64 `json:"OriginalPriceOneYear,omitempty" name:"OriginalPriceOneYear"`

	// Discounted 1-year payment, in USD. This parameter is only available to upfront payment mode.
	// Note: this field may return `null`, indicating that no valid value was found.
	// Note: this field may return `null`, indicating that no valid value was found.
	DiscountPriceOneYear *float64 `json:"DiscountPriceOneYear,omitempty" name:"DiscountPriceOneYear"`

	// Discount for 1-year upfront payment. For example, 20.0 indicates 80% off.
	// Note: this field may return `null`, indicating that no valid value was found.
	// Note: this field may return `null`, indicating that no valid value was found.
	DiscountOneYear *float64 `json:"DiscountOneYear,omitempty" name:"DiscountOneYear"`
}

type JobView struct {
	// Job ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// Job name
	JobName *string `json:"JobName,omitempty" name:"JobName"`

	// Job state
	JobState *string `json:"JobState,omitempty" name:"JobState"`

	// Job priority
	Priority *int64 `json:"Priority,omitempty" name:"Priority"`

	// Location information
	// Note: This field may return `null`, indicating that no valid value was found.
	Placement *Placement `json:"Placement,omitempty" name:"Placement"`

	// Creation time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// End time
	// Note: This field may return `null`, indicating that no valid value was found.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Task statistical metrics
	TaskMetrics *TaskMetrics `json:"TaskMetrics,omitempty" name:"TaskMetrics"`

	// List of tags bound with the job.
	// Note: This field may return `null`, indicating that no valid value was found.
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

type LocalDiskType struct {
	// Type of a local disk.
	Type *string `json:"Type,omitempty" name:"Type"`

	// Attributes of a local disk.
	PartitionType *string `json:"PartitionType,omitempty" name:"PartitionType"`

	// Minimum size of a local disk.
	MinSize *int64 `json:"MinSize,omitempty" name:"MinSize"`

	// Maximum size of a local disk.
	MaxSize *int64 `json:"MaxSize,omitempty" name:"MaxSize"`

	// Whether a local disk is required during purchase. Valid values:<br><li>REQUIRED: required<br><li>OPTIONAL: optional
	Required *string `json:"Required,omitempty" name:"Required"`
}

type LoginSettings struct {
	// Login password of the instance. <br><li>Linux instances: 8-16 characters, containing at least two of the following categories: [a-z, A-Z], [0-9] and [()`~!@#$%^&*-+=|{}[]:;',.?/]. <br><li>Windows instances: 12-16 characters, containing at least three of the following categories: [a-z], [A-Z], [0-9] and [()`~!@#$%^&*-+={}[]:;',.?/]. <br><br>If this parameter is not specified, a random password will be generated and sent to you via the Message Center.
	Password *string `json:"Password,omitempty" name:"Password"`

	// List of key IDs. After an instance is associated with a key, you can access the instance with the private key in the key pair. You can call `DescribeKeyPairs` to obtain `KeyId`. Key and password cannot be specified at the same time. Windows instances do not support keys. Currently, you can only specify one key when purchasing an instance.
	KeyIds []*string `json:"KeyIds,omitempty" name:"KeyIds"`

	// Whether to keep the original settings of an image. Values: `TRUE` (default), `FALSE`. It cannot be specified together with `Password` or `KeyIds.N`. You can specify this parameter as `TRUE` only when you create an instance using a custom image, a shared image, or an imported image. 
	KeepImageLogin *string `json:"KeepImageLogin,omitempty" name:"KeepImageLogin"`
}

// Predefined struct for user
type ModifyComputeEnvRequestParams struct {
	// Compute environment ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// Number of desired compute nodes
	DesiredComputeNodeCount *int64 `json:"DesiredComputeNodeCount,omitempty" name:"DesiredComputeNodeCount"`

	// Compute environment name
	EnvName *string `json:"EnvName,omitempty" name:"EnvName"`

	// Compute environment description
	EnvDescription *string `json:"EnvDescription,omitempty" name:"EnvDescription"`

	// Compute environment attributes
	EnvData *ComputeEnvData `json:"EnvData,omitempty" name:"EnvData"`
}

type ModifyComputeEnvRequest struct {
	*tchttp.BaseRequest
	
	// Compute environment ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// Number of desired compute nodes
	DesiredComputeNodeCount *int64 `json:"DesiredComputeNodeCount,omitempty" name:"DesiredComputeNodeCount"`

	// Compute environment name
	EnvName *string `json:"EnvName,omitempty" name:"EnvName"`

	// Compute environment description
	EnvDescription *string `json:"EnvDescription,omitempty" name:"EnvDescription"`

	// Compute environment attributes
	EnvData *ComputeEnvData `json:"EnvData,omitempty" name:"EnvData"`
}

func (r *ModifyComputeEnvRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyComputeEnvRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "DesiredComputeNodeCount")
	delete(f, "EnvName")
	delete(f, "EnvDescription")
	delete(f, "EnvData")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyComputeEnvRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyComputeEnvResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyComputeEnvResponse struct {
	*tchttp.BaseResponse
	Response *ModifyComputeEnvResponseParams `json:"Response"`
}

func (r *ModifyComputeEnvResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyComputeEnvResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTaskTemplateRequestParams struct {
	// Task template ID
	TaskTemplateId *string `json:"TaskTemplateId,omitempty" name:"TaskTemplateId"`

	// Task template name
	TaskTemplateName *string `json:"TaskTemplateName,omitempty" name:"TaskTemplateName"`

	// Task template description
	TaskTemplateDescription *string `json:"TaskTemplateDescription,omitempty" name:"TaskTemplateDescription"`

	// Task template information
	TaskTemplateInfo *Task `json:"TaskTemplateInfo,omitempty" name:"TaskTemplateInfo"`
}

type ModifyTaskTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Task template ID
	TaskTemplateId *string `json:"TaskTemplateId,omitempty" name:"TaskTemplateId"`

	// Task template name
	TaskTemplateName *string `json:"TaskTemplateName,omitempty" name:"TaskTemplateName"`

	// Task template description
	TaskTemplateDescription *string `json:"TaskTemplateDescription,omitempty" name:"TaskTemplateDescription"`

	// Task template information
	TaskTemplateInfo *Task `json:"TaskTemplateInfo,omitempty" name:"TaskTemplateInfo"`
}

func (r *ModifyTaskTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTaskTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskTemplateId")
	delete(f, "TaskTemplateName")
	delete(f, "TaskTemplateDescription")
	delete(f, "TaskTemplateInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTaskTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTaskTemplateResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyTaskTemplateResponse struct {
	*tchttp.BaseResponse
	Response *ModifyTaskTemplateResponseParams `json:"Response"`
}

func (r *ModifyTaskTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTaskTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MountDataDisk struct {
	// Mounting point. A valid path (for Linux) for a drive (for Windows, such as "H:\\")
	LocalPath *string `json:"LocalPath,omitempty" name:"LocalPath"`

	// File system type. Linux: `EXT3` (default) and `EXT4`. Windows: `NTFS`
	FileSystemType *string `json:"FileSystemType,omitempty" name:"FileSystemType"`
}

type NamedComputeEnv struct {
	// Compute environment name
	EnvName *string `json:"EnvName,omitempty" name:"EnvName"`

	// Number of desired compute nodes
	DesiredComputeNodeCount *int64 `json:"DesiredComputeNodeCount,omitempty" name:"DesiredComputeNodeCount"`

	// Compute environment description
	EnvDescription *string `json:"EnvDescription,omitempty" name:"EnvDescription"`

	// Compute environment management type
	EnvType *string `json:"EnvType,omitempty" name:"EnvType"`

	// Compute environment's specific parameters
	EnvData *EnvData `json:"EnvData,omitempty" name:"EnvData"`

	// Data disk mounting option
	MountDataDisks []*MountDataDisk `json:"MountDataDisks,omitempty" name:"MountDataDisks"`

	// Authorization information
	Authentications []*Authentication `json:"Authentications,omitempty" name:"Authentications"`

	// Input mapping information
	InputMappings []*InputMapping `json:"InputMappings,omitempty" name:"InputMappings"`

	// Agent running mode; applicable for Windows
	AgentRunningMode *AgentRunningMode `json:"AgentRunningMode,omitempty" name:"AgentRunningMode"`

	// Notification information
	Notifications []*Notification `json:"Notifications,omitempty" name:"Notifications"`

	// Policy for inactive nodes. Default: `RECREATE`, which means that instance resources will be re-created periodically for compute nodes where instance creation fails or is abnormally returned.
	ActionIfComputeNodeInactive *string `json:"ActionIfComputeNodeInactive,omitempty" name:"ActionIfComputeNodeInactive"`

	// When the instances are failed to be created or returned because of exceptions, the related compute node will retry to create instances periodically. This parameter specifies the maximum retry attempts. The max value is 100 and the default value is `7`.
	ResourceMaxRetryCount *int64 `json:"ResourceMaxRetryCount,omitempty" name:"ResourceMaxRetryCount"`

	// List of tags to bind with a compute environment. Each compute environment can have up to 10 tags.
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// Target of the notification
	// Values: `CMQ`, `TDMQ_CMQ`
	// `CMQ`: Tencent Cloud CMQ (default)
	// `TDMQ_CMQ`: Tencent Cloud TDMQ_CMQ.<br/>Note: CMQ has been discontinued. Please use `TDMQ_CMQ`. See [CMQ Queue Migration to TDMQ for CMQ](https://intl.cloud.tencent.com/document/product/406/60860?from_cn_redirect=1)
	NotificationTarget *string `json:"NotificationTarget,omitempty" name:"NotificationTarget"`
}

type Notification struct {
	// CMQ topic name which should be valid and associated with a subscription
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// Event configuration
	EventConfigs []*EventConfig `json:"EventConfigs,omitempty" name:"EventConfigs"`
}

type OutputMapping struct {
	// Source path
	SourcePath *string `json:"SourcePath,omitempty" name:"SourcePath"`

	// Destination path
	DestinationPath *string `json:"DestinationPath,omitempty" name:"DestinationPath"`

	// Output mapping options
	// Note: This field may return `null`, indicating that no valid value was found.
	OutputMappingOption *OutputMappingOption `json:"OutputMappingOption,omitempty" name:"OutputMappingOption"`
}

type OutputMappingConfig struct {
	// Storage type. Only `COS` is supported.
	Scene *string `json:"Scene,omitempty" name:"Scene"`

	// Number of concurrent workers
	WorkerNum *int64 `json:"WorkerNum,omitempty" name:"WorkerNum"`

	// Size of a worker part, in MB.
	WorkerPartSize *int64 `json:"WorkerPartSize,omitempty" name:"WorkerPartSize"`
}

type OutputMappingOption struct {
	// The mapped output workspace on the container side for the instance.
	// `BATCH_WORKSPACE` (Default): The workspace is the workspace defined for the usage of Batch Compute. Batch Compute ensures the isolation between jobs.
	// `GLOBAL_WORKSPACE`: The workspace is the instance OS space..
	// Note: This field may return `null`, indicating that no valid value was found.
	Workspace *string `json:"Workspace,omitempty" name:"Workspace"`
}

type Placement struct {
	// ID of the availability zone where the instance resides. You can call the [DescribeZones](https://intl.cloud.tencent.com/document/product/213/35071) API and obtain the ID in the returned `Zone` field.
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// ID of the project to which the instance belongs. This parameter can be obtained from the `projectId` returned by [DescribeProject](https://intl.cloud.tencent.com/document/api/651/78725?from_cn_redirect=1). If this is left empty, the default project is used.
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// ID list of CDHs from which the instance can be created. If you have purchased CDHs and specify this parameter, the instances you purchase will be randomly deployed on the CDHs.
	HostIds []*string `json:"HostIds,omitempty" name:"HostIds"`

	// IPs of the hosts to create CVMs
	HostIps []*string `json:"HostIps,omitempty" name:"HostIps"`

	// The ID of the CDH to which the instance belongs, only used as an output parameter.
	HostId *string `json:"HostId,omitempty" name:"HostId"`
}

type RedirectInfo struct {
	// Standard output redirection path
	StdoutRedirectPath *string `json:"StdoutRedirectPath,omitempty" name:"StdoutRedirectPath"`

	// Standard error redirection path
	StderrRedirectPath *string `json:"StderrRedirectPath,omitempty" name:"StderrRedirectPath"`

	// Standard output redirection file name, which supports three placeholders: `${BATCH_JOB_ID}`, `${BATCH_TASK_NAME}`, and `${BATCH_TASK_INSTANCE_INDEX}`
	StdoutRedirectFileName *string `json:"StdoutRedirectFileName,omitempty" name:"StdoutRedirectFileName"`

	// Standard error redirection file name, which supports three placeholders: `${BATCH_JOB_ID}`, `${BATCH_TASK_NAME}`, and `${BATCH_TASK_INSTANCE_INDEX}`
	StderrRedirectFileName *string `json:"StderrRedirectFileName,omitempty" name:"StderrRedirectFileName"`
}

type RedirectLocalInfo struct {
	// Standard output redirection local path
	StdoutLocalPath *string `json:"StdoutLocalPath,omitempty" name:"StdoutLocalPath"`

	// Standard error redirection local path
	StderrLocalPath *string `json:"StderrLocalPath,omitempty" name:"StderrLocalPath"`

	// Standard output redirection local file name, which supports three placeholders: `${BATCH_JOB_ID}`, `${BATCH_TASK_NAME}`, and `${BATCH_TASK_INSTANCE_INDEX}`
	StdoutLocalFileName *string `json:"StdoutLocalFileName,omitempty" name:"StdoutLocalFileName"`

	// Standard error redirection local file name, which supports three placeholders: `${BATCH_JOB_ID}`, `${BATCH_TASK_NAME}`, and `${BATCH_TASK_INSTANCE_INDEX}`
	StderrLocalFileName *string `json:"StderrLocalFileName,omitempty" name:"StderrLocalFileName"`
}

// Predefined struct for user
type RetryJobsRequestParams struct {
	// List of job IDs.
	JobIds []*string `json:"JobIds,omitempty" name:"JobIds"`
}

type RetryJobsRequest struct {
	*tchttp.BaseRequest
	
	// List of job IDs.
	JobIds []*string `json:"JobIds,omitempty" name:"JobIds"`
}

func (r *RetryJobsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RetryJobsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RetryJobsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RetryJobsResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RetryJobsResponse struct {
	*tchttp.BaseResponse
	Response *RetryJobsResponseParams `json:"Response"`
}

func (r *RetryJobsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RetryJobsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RunAutomationServiceEnabled struct {
	// Whether to enable the TAT service. Valid values: <br><li>`TRUE`: yes;<br><li>`FALSE`: no<br><br>Default: `FALSE`.
	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`
}

type RunMonitorServiceEnabled struct {
	// Whether to enable [Cloud Monitor](https://intl.cloud.tencent.com/document/product/248?from_cn_redirect=1). Valid values: <br><li>TRUE: enable Cloud Monitor <br><li>FALSE: do not enable Cloud Monitor <br><br>Default value: TRUE.
	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`
}

type RunSecurityServiceEnabled struct {
	// Whether to enable [Cloud Security](https://intl.cloud.tencent.com/document/product/296?from_cn_redirect=1). Valid values: <br><li>TRUE: enable Cloud Security <br><li>FALSE: do not enable Cloud Security <br><br>Default value: TRUE.
	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`
}

type SpotMarketOptions struct {
	// Bidding price
	MaxPrice *string `json:"MaxPrice,omitempty" name:"MaxPrice"`

	// Bidding request type. Currently only "one-time" is supported.
	SpotInstanceType *string `json:"SpotInstanceType,omitempty" name:"SpotInstanceType"`
}

type StorageBlock struct {
	// Local HDD storage type. Value: LOCAL_PRO.
	// Note: This field may return null, indicating that no valid value is found.
	Type *string `json:"Type,omitempty" name:"Type"`

	// Minimum capacity of local HDD storage
	// Note: This field may return null, indicating that no valid value is found.
	MinSize *int64 `json:"MinSize,omitempty" name:"MinSize"`

	// Maximum capacity of local HDD storage
	// Note: This field may return null, indicating that no valid value is found.
	MaxSize *int64 `json:"MaxSize,omitempty" name:"MaxSize"`
}

type SystemDisk struct {
	// System disk type. For more information about the limits of system disk types, please see [Storage Overview](https://intl.cloud.tencent.com/document/product/213/4952?from_cn_redirect=1). Valid values:<br><li>LOCAL_BASIC: local disk<br><li>LOCAL_SSD: local SSD disk<br><li>CLOUD_BASIC: HDD cloud disk<br><li>CLOUD_SSD: SSD cloud disk<br><li>CLOUD_PREMIUM: Premium cloud storage<br><li>CLOUD_BSSD: Balanced SSD<br><br>The disk currently in stock will be used by default.
	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`

	// System disk ID. System disks whose type is `LOCAL_BASIC` or `LOCAL_SSD` do not have an ID and do not support this parameter.
	// It is only used as a response parameter for APIs such as `DescribeInstances`, and cannot be used as a request parameter for APIs such as `RunInstances`.
	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`

	// System disk size; unit: GB; default value: 50 GB.
	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`

	// ID of the dedicated cluster to which the instance belongs.
	CdcId *string `json:"CdcId,omitempty" name:"CdcId"`
}

type Tag struct {
	// Tag key
	// Note: This field may return `null`, indicating that no valid value was found.
	Key *string `json:"Key,omitempty" name:"Key"`

	// Tag value
	// Note: This field may return `null`, indicating that no valid value was found.
	Value *string `json:"Value,omitempty" name:"Value"`
}

type Task struct {
	// Application information
	Application *Application `json:"Application,omitempty" name:"Application"`

	// Task name, which should be unique within a job
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// Number of running task instances
	TaskInstanceNum *uint64 `json:"TaskInstanceNum,omitempty" name:"TaskInstanceNum"`

	// Compute environment information. One (and only one) parameter must be specified for ComputeEnv and EnvId.
	ComputeEnv *AnonymousComputeEnv `json:"ComputeEnv,omitempty" name:"ComputeEnv"`

	// Compute environment ID. One (and only one) parameter must be specified for ComputeEnv and EnvId.
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// Redirection information
	RedirectInfo *RedirectInfo `json:"RedirectInfo,omitempty" name:"RedirectInfo"`

	// Local redirection information
	RedirectLocalInfo *RedirectLocalInfo `json:"RedirectLocalInfo,omitempty" name:"RedirectLocalInfo"`

	// Input mapping
	InputMappings []*InputMapping `json:"InputMappings,omitempty" name:"InputMappings"`

	// Output mapping
	OutputMappings []*OutputMapping `json:"OutputMappings,omitempty" name:"OutputMappings"`

	// Output mapping configuration
	OutputMappingConfigs []*OutputMappingConfig `json:"OutputMappingConfigs,omitempty" name:"OutputMappingConfigs"`

	// Custom environment variable
	EnvVars []*EnvVar `json:"EnvVars,omitempty" name:"EnvVars"`

	// Authorization information
	Authentications []*Authentication `json:"Authentications,omitempty" name:"Authentications"`

	// The processing method after the TaskInstance fails; Value range: `TERMINATE` (default), `INTERRUPT`, `FAST_INTERRUPT`.
	FailedAction *string `json:"FailedAction,omitempty" name:"FailedAction"`

	// The maximum number of retries after the task fails. Range: 0 - 5. Default value: 0
	MaxRetryCount *uint64 `json:"MaxRetryCount,omitempty" name:"MaxRetryCount"`

	// Timeout period of the task in seconds. Defaults value: 86400
	Timeout *uint64 `json:"Timeout,omitempty" name:"Timeout"`

	// The maximum number of concurrent tasks. Range: 0 - 200000. There is no limit by default.
	MaxConcurrentNum *uint64 `json:"MaxConcurrentNum,omitempty" name:"MaxConcurrentNum"`

	// Restarts the compute node after the task is completed. This is suitable for specifying the compute environment for task execution.
	RestartComputeNode *bool `json:"RestartComputeNode,omitempty" name:"RestartComputeNode"`

	// Maximum number of retry attempts after failing to create computing resources such as the CVM in the task launch process. Default: `0`; Maximum: `100`.
	ResourceMaxRetryCount *uint64 `json:"ResourceMaxRetryCount,omitempty" name:"ResourceMaxRetryCount"`
}

type TaskInstanceLog struct {
	// Task instance
	TaskInstanceIndex *uint64 `json:"TaskInstanceIndex,omitempty" name:"TaskInstanceIndex"`

	// Standard output log (Base64-encoded, up to 2048 bytes after decompression)
	// Note: This field may return `null`, indicating that no valid value was found.
	StdoutLog *string `json:"StdoutLog,omitempty" name:"StdoutLog"`

	// Standard error log (Base64-encoded, up to 2048 bytes after decompression)
	// Note: This field may return `null`, indicating that no valid value was found.
	StderrLog *string `json:"StderrLog,omitempty" name:"StderrLog"`

	// Standard output redirection path
	// Note: This field may return `null`, indicating that no valid value was found.
	StdoutRedirectPath *string `json:"StdoutRedirectPath,omitempty" name:"StdoutRedirectPath"`

	// Standard error redirection path
	// Note: This field may return `null`, indicating that no valid value was found.
	StderrRedirectPath *string `json:"StderrRedirectPath,omitempty" name:"StderrRedirectPath"`

	// Standard output redirection file name
	// Note: This field may return `null`, indicating that no valid value was found.
	StdoutRedirectFileName *string `json:"StdoutRedirectFileName,omitempty" name:"StdoutRedirectFileName"`

	// Standard error redirection file name
	// Note: This field may return `null`, indicating that no valid value was found.
	StderrRedirectFileName *string `json:"StderrRedirectFileName,omitempty" name:"StderrRedirectFileName"`
}

type TaskInstanceMetrics struct {
	// Number of submitted tasks
	SubmittedCount *int64 `json:"SubmittedCount,omitempty" name:"SubmittedCount"`

	// Number of pending tasks
	PendingCount *int64 `json:"PendingCount,omitempty" name:"PendingCount"`

	// Number of Runnable tasks
	RunnableCount *int64 `json:"RunnableCount,omitempty" name:"RunnableCount"`

	// Number of starting tasks
	StartingCount *int64 `json:"StartingCount,omitempty" name:"StartingCount"`

	// Number of running tasks
	RunningCount *int64 `json:"RunningCount,omitempty" name:"RunningCount"`

	// Number of successful tasks
	SucceedCount *int64 `json:"SucceedCount,omitempty" name:"SucceedCount"`

	// Number of failed and interrupted tasks
	FailedInterruptedCount *int64 `json:"FailedInterruptedCount,omitempty" name:"FailedInterruptedCount"`

	// Number of failed tasks
	FailedCount *int64 `json:"FailedCount,omitempty" name:"FailedCount"`
}

type TaskInstanceView struct {
	// Task instance index
	TaskInstanceIndex *int64 `json:"TaskInstanceIndex,omitempty" name:"TaskInstanceIndex"`

	// Task instance state
	TaskInstanceState *string `json:"TaskInstanceState,omitempty" name:"TaskInstanceState"`

	// Exit code after application execution is completed
	// Note: This field may return `null`, indicating that no valid value was found.
	ExitCode *int64 `json:"ExitCode,omitempty" name:"ExitCode"`

	// Task instance state reason. If the task instance fails, the reason for the failure will be logged.
	StateReason *string `json:"StateReason,omitempty" name:"StateReason"`

	// The `InstanceId` of the compute node (e.g., CVM instance) where the task instance is running. This field is empty if the task instance is not running or has already been completed and will change when the task instance is retried.
	// Note: This field may return `null`, indicating that no valid value was found.
	ComputeNodeInstanceId *string `json:"ComputeNodeInstanceId,omitempty" name:"ComputeNodeInstanceId"`

	// Creation time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Start time
	// Note: This field may return `null`, indicating that no valid value was found.
	LaunchTime *string `json:"LaunchTime,omitempty" name:"LaunchTime"`

	// Running start time
	// Note: This field may return `null`, indicating that no valid value was found.
	RunningTime *string `json:"RunningTime,omitempty" name:"RunningTime"`

	// Task end time
	// Note: This field may return `null`, indicating that no valid value was found.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Redirection information
	RedirectInfo *RedirectInfo `json:"RedirectInfo,omitempty" name:"RedirectInfo"`

	// Task instance state reason details. If the task instance fails, the reason for the failure will be logged
	StateDetailedReason *string `json:"StateDetailedReason,omitempty" name:"StateDetailedReason"`
}

type TaskMetrics struct {
	// Number of submitted tasks
	SubmittedCount *int64 `json:"SubmittedCount,omitempty" name:"SubmittedCount"`

	// Number of pending tasks
	PendingCount *int64 `json:"PendingCount,omitempty" name:"PendingCount"`

	// Number of Runnable tasks
	RunnableCount *int64 `json:"RunnableCount,omitempty" name:"RunnableCount"`

	// Number of starting tasks
	StartingCount *int64 `json:"StartingCount,omitempty" name:"StartingCount"`

	// Number of running tasks
	RunningCount *int64 `json:"RunningCount,omitempty" name:"RunningCount"`

	// Number of successful tasks
	SucceedCount *int64 `json:"SucceedCount,omitempty" name:"SucceedCount"`

	// Number of failed and interrupted tasks
	FailedInterruptedCount *int64 `json:"FailedInterruptedCount,omitempty" name:"FailedInterruptedCount"`

	// Failed count
	FailedCount *int64 `json:"FailedCount,omitempty" name:"FailedCount"`
}

type TaskTemplateView struct {
	// Task template ID
	TaskTemplateId *string `json:"TaskTemplateId,omitempty" name:"TaskTemplateId"`

	// Task template name
	TaskTemplateName *string `json:"TaskTemplateName,omitempty" name:"TaskTemplateName"`

	// Task template description
	TaskTemplateDescription *string `json:"TaskTemplateDescription,omitempty" name:"TaskTemplateDescription"`

	// Task template information
	TaskTemplateInfo *Task `json:"TaskTemplateInfo,omitempty" name:"TaskTemplateInfo"`

	// Creation time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Tag list bound to the task template.
	// Note: This field may return `null`, indicating that no valid value was found.
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

type TaskView struct {
	// Task name
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// Task status
	TaskState *string `json:"TaskState,omitempty" name:"TaskState"`

	// Start time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// End time
	// Note: This field may return `null`, indicating that no valid value was found.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

// Predefined struct for user
type TerminateComputeNodeRequestParams struct {
	// Compute environment ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// Compute node ID
	ComputeNodeId *string `json:"ComputeNodeId,omitempty" name:"ComputeNodeId"`
}

type TerminateComputeNodeRequest struct {
	*tchttp.BaseRequest
	
	// Compute environment ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// Compute node ID
	ComputeNodeId *string `json:"ComputeNodeId,omitempty" name:"ComputeNodeId"`
}

func (r *TerminateComputeNodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TerminateComputeNodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "ComputeNodeId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TerminateComputeNodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TerminateComputeNodeResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type TerminateComputeNodeResponse struct {
	*tchttp.BaseResponse
	Response *TerminateComputeNodeResponseParams `json:"Response"`
}

func (r *TerminateComputeNodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TerminateComputeNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TerminateComputeNodesRequestParams struct {
	// Compute environment ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// List of compute node IDs
	ComputeNodeIds []*string `json:"ComputeNodeIds,omitempty" name:"ComputeNodeIds"`
}

type TerminateComputeNodesRequest struct {
	*tchttp.BaseRequest
	
	// Compute environment ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// List of compute node IDs
	ComputeNodeIds []*string `json:"ComputeNodeIds,omitempty" name:"ComputeNodeIds"`
}

func (r *TerminateComputeNodesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TerminateComputeNodesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "ComputeNodeIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TerminateComputeNodesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TerminateComputeNodesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type TerminateComputeNodesResponse struct {
	*tchttp.BaseResponse
	Response *TerminateComputeNodesResponseParams `json:"Response"`
}

func (r *TerminateComputeNodesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TerminateComputeNodesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TerminateJobRequestParams struct {
	// Job ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

type TerminateJobRequest struct {
	*tchttp.BaseRequest
	
	// Job ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

func (r *TerminateJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TerminateJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TerminateJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TerminateJobResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type TerminateJobResponse struct {
	*tchttp.BaseResponse
	Response *TerminateJobResponseParams `json:"Response"`
}

func (r *TerminateJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TerminateJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TerminateTaskInstanceRequestParams struct {
	// Job ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// Task name
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// Task instance index
	TaskInstanceIndex *int64 `json:"TaskInstanceIndex,omitempty" name:"TaskInstanceIndex"`
}

type TerminateTaskInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Job ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// Task name
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// Task instance index
	TaskInstanceIndex *int64 `json:"TaskInstanceIndex,omitempty" name:"TaskInstanceIndex"`
}

func (r *TerminateTaskInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TerminateTaskInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	delete(f, "TaskName")
	delete(f, "TaskInstanceIndex")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TerminateTaskInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TerminateTaskInstanceResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type TerminateTaskInstanceResponse struct {
	*tchttp.BaseResponse
	Response *TerminateTaskInstanceResponseParams `json:"Response"`
}

func (r *TerminateTaskInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TerminateTaskInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VirtualPrivateCloud struct {
	// VPC ID in the format of `vpc-xxx`. To obtain valid VPC IDs, you can log in to the [console](https://console.cloud.tencent.com/vpc/vpc?rid=1) or call the [DescribeVpcEx](https://intl.cloud.tencent.com/document/api/215/1372?from_cn_redirect=1) API and look for the `unVpcId` fields in the response. If you specify `DEFAULT` for both `VpcId` and `SubnetId` when creating an instance, the default VPC will be used.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// VPC subnet ID in the format `subnet-xxx`. To obtain valid subnet IDs, you can log in to the [console](https://console.cloud.tencent.com/vpc/subnet?rid=1) or call [DescribeSubnets](https://intl.cloud.tencent.com/document/api/215/15784?from_cn_redirect=1) and look for the `unSubnetId` fields in the response. If you specify `DEFAULT` for both `SubnetId` and `VpcId` when creating an instance, the default VPC will be used.
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// Whether to use a CVM instance as a public gateway. The public gateway is only available when the instance has a public IP and resides in a VPC. Valid values: <br><li>`TRUE`: yes;<br><li>`FALSE`: no<br><br>Default: `FALSE`.
	AsVpcGateway *bool `json:"AsVpcGateway,omitempty" name:"AsVpcGateway"`

	// Array of VPC subnet IPs. You can use this parameter when creating instances or modifying VPC attributes of instances. Currently you can specify multiple IPs in one subnet only when creating multiple instances at the same time.
	PrivateIpAddresses []*string `json:"PrivateIpAddresses,omitempty" name:"PrivateIpAddresses"`

	// Number of IPv6 addresses randomly generated for the ENI.
	Ipv6AddressCount *uint64 `json:"Ipv6AddressCount,omitempty" name:"Ipv6AddressCount"`
}