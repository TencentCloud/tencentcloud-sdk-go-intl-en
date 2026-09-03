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

package v20211111

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

type AuthToken struct {
	// AuthToken basic information.
	Base *AuthTokenBase `json:"Base,omitnil,omitempty" name:"Base"`

	// AuthToken throttling array.
	Limits []*AuthTokenLimit `json:"Limits,omitnil,omitempty" name:"Limits"`
}

type AuthTokenBase struct {
	// Token value.
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// Token alias.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Token description.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Token creation time.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Token status.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

type AuthTokenLimit struct {
	// Frequency limit policy. Valid values: PerMinute (frequency limit per minute) and PerDay (daily frequency limit).
	Strategy *string `json:"Strategy,omitnil,omitempty" name:"Strategy"`

	// Upper limit.
	Max *uint64 `json:"Max,omitnil,omitempty" name:"Max"`
}

type CBSConfig struct {
	// Storage size.
	// Note: This field may return null, indicating that no valid values can be obtained.
	VolumeSizeInGB *int64 `json:"VolumeSizeInGB,omitnil,omitempty" name:"VolumeSizeInGB"`
}

type CFSConfig struct {
	// CFS instance ID.
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// Storage path.
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// Mounting type of CFS. Valid values: STORAGE and SOURCE, which respectively indicate the storage expansion mode and the data source mode. The default value is STORAGE.Note: This field may return null, indicating that no valid values can be obtained.
	MountType *string `json:"MountType,omitnil,omitempty" name:"MountType"`

	// Protocol. Valid values: NFS and TURBO.Note: This field may return null, indicating that no valid values can be obtained.
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`
}

type CFSTurbo struct {
	// CFSTurbo instance ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// CFSTurbo path.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`
}

type CodeRepoConfig struct {
	// Code repository ID.
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// Target address for the code repository download.
	TargetPath *string `json:"TargetPath,omitnil,omitempty" name:"TargetPath"`
}

type Container struct {
	// Name.Note: This field may return null, indicating that no valid values can be obtained.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// id
	// Note: This field may return null, indicating that no valid values can be obtained.
	ContainerId *string `json:"ContainerId,omitnil,omitempty" name:"ContainerId"`

	// Image address.Note: This field may return null, indicating that no valid values can be obtained.
	Image *string `json:"Image,omitnil,omitempty" name:"Image"`

	// Container status.Note: This field may return null, indicating that no valid values can be obtained.
	Status *ContainerStatus `json:"Status,omitnil,omitempty" name:"Status"`
}

type ContainerStatus struct {
	// Number of restarts.Note: This field may return null, indicating that no valid values can be obtained.
	RestartCount *int64 `json:"RestartCount,omitnil,omitempty" name:"RestartCount"`

	// Status.Note: This field may return null, indicating that no valid values can be obtained.
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// Whether it is ready.Note: This field may return null, indicating that no valid values can be obtained.
	Ready *bool `json:"Ready,omitnil,omitempty" name:"Ready"`

	// Status reason.Note: This field may return null, indicating that no valid values can be obtained.
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`

	// Container error message.Note: This field may return null, indicating that no valid values can be obtained.
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`
}

type CosPathInfo struct {
	// Bucket.Note: This field may return null, indicating that no valid values can be obtained.
	Bucket *string `json:"Bucket,omitnil,omitempty" name:"Bucket"`

	// Region.Note: This field may return null, indicating that no valid values can be obtained.
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Path list. Only one path is supported.Note: This field may return null, indicating that no valid values can be obtained.
	Paths []*string `json:"Paths,omitnil,omitempty" name:"Paths"`
}

// Predefined struct for user
type CreateTrainingTaskRequestParams struct {
	// Training task name. The name cannot exceed 60 characters in length, and can contain only Chinese characters, letters, digits, underscores (_), and hyphens (-). It must start with a Chinese character, letter, or digit.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Billing mode. For example, PREPAID indicates yearly/monthly subscription (resource group).
	// POSTPAID_BY_HOUR indicates pay-as-you-go mode.
	ChargeType *string `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`

	// Resource configuration. Specify the CVM instance specification ID and number of nodes. The API for querying the CVM instance specification ID is DescribeBillingSpecsPrice. For example, [{"Role":"WORKER", "InstanceType": "TI.S.MEDIUM.POST", "InstanceNum": 1}].
	ResourceConfigInfos []*ResourceConfigInfo `json:"ResourceConfigInfos,omitnil,omitempty" name:"ResourceConfigInfos"`

	// TI Workspace ID. Used solely for the "Workspace" allowlist feature. To use this feature, please contact a TI administrator to enable allowlisting.
	TiProjectId *string `json:"TiProjectId,omitnil,omitempty" name:"TiProjectId"`

	// Training framework name, which can be queried via the DescribeTrainingFrameworks API. For example, SPARK, PYSPARK, TENSORFLOW, and PYTORCH.
	FrameworkName *string `json:"FrameworkName,omitnil,omitempty" name:"FrameworkName"`

	// Training framework version, which can be queried via the DescribeTrainingFrameworks API. For example, 1.15 and 1.9.
	FrameworkVersion *string `json:"FrameworkVersion,omitnil,omitempty" name:"FrameworkVersion"`

	// Training framework environment, which can be queried via the DescribeTrainingFrameworks API. For example, tf1.15-py3.7-cpu and torch1.9-py3.8-cuda11.1-gpu.
	FrameworkEnvironment *string `json:"FrameworkEnvironment,omitnil,omitempty" name:"FrameworkEnvironment"`

	// ID of the prepaid dedicated resource group, which can be queried via the DescribeBillingResourceGroups API.
	ResourceGroupId *string `json:"ResourceGroupId,omitnil,omitempty" name:"ResourceGroupId"`

	// Tag configuration.
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Custom image information.
	ImageInfo *ImageInfo `json:"ImageInfo,omitnil,omitempty" name:"ImageInfo"`

	// COS code package path.
	CodePackagePath *CosPathInfo `json:"CodePackagePath,omitnil,omitempty" name:"CodePackagePath"`

	// Task startup command. Specify this parameter based on the task training mode. If the configuration fails due to special characters, use the EncodedStartCmdInfo parameter instead.
	StartCmdInfo *StartCmdInfo `json:"StartCmdInfo,omitnil,omitempty" name:"StartCmdInfo"`

	// Training mode, which can be queried via the DescribeTrainingFrameworks API. For example, PS_WORKER, DDP, MPI, and HOROVOD.
	TrainingMode *string `json:"TrainingMode,omitnil,omitempty" name:"TrainingMode"`

	// Data configurations. This parameter depends on the DataSource field. The maximum number of configurations is 10.
	DataConfigs []*DataConfig `json:"DataConfigs,omitnil,omitempty" name:"DataConfigs"`

	// VPC Id
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Subnet ID.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// COS training output path.
	Output *CosPathInfo `json:"Output,omitnil,omitempty" name:"Output"`

	// CLS logging configuration.
	LogConfig *LogConfig `json:"LogConfig,omitnil,omitempty" name:"LogConfig"`

	// Tuning parameters. The value of this parameter cannot exceed 2048 characters in length.
	TuningParameters *string `json:"TuningParameters,omitnil,omitempty" name:"TuningParameters"`

	// Indicates whether to report logs.
	LogEnable *bool `json:"LogEnable,omitnil,omitempty" name:"LogEnable"`

	// Remarks. The value of this parameter cannot exceed 1024 characters.
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// Data source. For example, DATASET, COS, CFS, CFSTurbo, HDFS, and GooseFSx.
	DataSource *string `json:"DataSource,omitnil,omitempty" name:"DataSource"`

	// Callback URL. This parameter is used for the asynchronous callback to create, start, or stop training tasks. For the callback format and content, see [[TI-ONE API Callback Description]](https://www.tencentcloud.com/document/product/851/84292?from_cn_redirect=1).
	CallbackUrl *string `json:"CallbackUrl,omitnil,omitempty" name:"CallbackUrl"`

	// Encoded task startup command. If StartCmdInfo is also configured, only this parameter takes effect.
	EncodedStartCmdInfo *EncodedStartCmdInfo `json:"EncodedStartCmdInfo,omitnil,omitempty" name:"EncodedStartCmdInfo"`

	// Code repository configuration.
	CodeRepos []*CodeRepoConfig `json:"CodeRepos,omitnil,omitempty" name:"CodeRepos"`

	// Network exposure configuration.
	ExposeNetworkConfig *ExposeNetworkConfig `json:"ExposeNetworkConfig,omitnil,omitempty" name:"ExposeNetworkConfig"`

	// Environment Variables.
	Envs []*EnvVar `json:"Envs,omitnil,omitempty" name:"Envs"`

	// Train tool configuration.
	TrainToolConfig *TrainToolConfig `json:"TrainToolConfig,omitnil,omitempty" name:"TrainToolConfig"`

	// Training Diagnostic Tool Configuration.
	ResourceSupplyAttribute *ResourceSupplyAttribute `json:"ResourceSupplyAttribute,omitnil,omitempty" name:"ResourceSupplyAttribute"`

	// Queue ID.
	Queues []*string `json:"Queues,omitnil,omitempty" name:"Queues"`
}

type CreateTrainingTaskRequest struct {
	*tchttp.BaseRequest
	
	// Training task name. The name cannot exceed 60 characters in length, and can contain only Chinese characters, letters, digits, underscores (_), and hyphens (-). It must start with a Chinese character, letter, or digit.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Billing mode. For example, PREPAID indicates yearly/monthly subscription (resource group).
	// POSTPAID_BY_HOUR indicates pay-as-you-go mode.
	ChargeType *string `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`

	// Resource configuration. Specify the CVM instance specification ID and number of nodes. The API for querying the CVM instance specification ID is DescribeBillingSpecsPrice. For example, [{"Role":"WORKER", "InstanceType": "TI.S.MEDIUM.POST", "InstanceNum": 1}].
	ResourceConfigInfos []*ResourceConfigInfo `json:"ResourceConfigInfos,omitnil,omitempty" name:"ResourceConfigInfos"`

	// TI Workspace ID. Used solely for the "Workspace" allowlist feature. To use this feature, please contact a TI administrator to enable allowlisting.
	TiProjectId *string `json:"TiProjectId,omitnil,omitempty" name:"TiProjectId"`

	// Training framework name, which can be queried via the DescribeTrainingFrameworks API. For example, SPARK, PYSPARK, TENSORFLOW, and PYTORCH.
	FrameworkName *string `json:"FrameworkName,omitnil,omitempty" name:"FrameworkName"`

	// Training framework version, which can be queried via the DescribeTrainingFrameworks API. For example, 1.15 and 1.9.
	FrameworkVersion *string `json:"FrameworkVersion,omitnil,omitempty" name:"FrameworkVersion"`

	// Training framework environment, which can be queried via the DescribeTrainingFrameworks API. For example, tf1.15-py3.7-cpu and torch1.9-py3.8-cuda11.1-gpu.
	FrameworkEnvironment *string `json:"FrameworkEnvironment,omitnil,omitempty" name:"FrameworkEnvironment"`

	// ID of the prepaid dedicated resource group, which can be queried via the DescribeBillingResourceGroups API.
	ResourceGroupId *string `json:"ResourceGroupId,omitnil,omitempty" name:"ResourceGroupId"`

	// Tag configuration.
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Custom image information.
	ImageInfo *ImageInfo `json:"ImageInfo,omitnil,omitempty" name:"ImageInfo"`

	// COS code package path.
	CodePackagePath *CosPathInfo `json:"CodePackagePath,omitnil,omitempty" name:"CodePackagePath"`

	// Task startup command. Specify this parameter based on the task training mode. If the configuration fails due to special characters, use the EncodedStartCmdInfo parameter instead.
	StartCmdInfo *StartCmdInfo `json:"StartCmdInfo,omitnil,omitempty" name:"StartCmdInfo"`

	// Training mode, which can be queried via the DescribeTrainingFrameworks API. For example, PS_WORKER, DDP, MPI, and HOROVOD.
	TrainingMode *string `json:"TrainingMode,omitnil,omitempty" name:"TrainingMode"`

	// Data configurations. This parameter depends on the DataSource field. The maximum number of configurations is 10.
	DataConfigs []*DataConfig `json:"DataConfigs,omitnil,omitempty" name:"DataConfigs"`

	// VPC Id
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Subnet ID.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// COS training output path.
	Output *CosPathInfo `json:"Output,omitnil,omitempty" name:"Output"`

	// CLS logging configuration.
	LogConfig *LogConfig `json:"LogConfig,omitnil,omitempty" name:"LogConfig"`

	// Tuning parameters. The value of this parameter cannot exceed 2048 characters in length.
	TuningParameters *string `json:"TuningParameters,omitnil,omitempty" name:"TuningParameters"`

	// Indicates whether to report logs.
	LogEnable *bool `json:"LogEnable,omitnil,omitempty" name:"LogEnable"`

	// Remarks. The value of this parameter cannot exceed 1024 characters.
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// Data source. For example, DATASET, COS, CFS, CFSTurbo, HDFS, and GooseFSx.
	DataSource *string `json:"DataSource,omitnil,omitempty" name:"DataSource"`

	// Callback URL. This parameter is used for the asynchronous callback to create, start, or stop training tasks. For the callback format and content, see [[TI-ONE API Callback Description]](https://www.tencentcloud.com/document/product/851/84292?from_cn_redirect=1).
	CallbackUrl *string `json:"CallbackUrl,omitnil,omitempty" name:"CallbackUrl"`

	// Encoded task startup command. If StartCmdInfo is also configured, only this parameter takes effect.
	EncodedStartCmdInfo *EncodedStartCmdInfo `json:"EncodedStartCmdInfo,omitnil,omitempty" name:"EncodedStartCmdInfo"`

	// Code repository configuration.
	CodeRepos []*CodeRepoConfig `json:"CodeRepos,omitnil,omitempty" name:"CodeRepos"`

	// Network exposure configuration.
	ExposeNetworkConfig *ExposeNetworkConfig `json:"ExposeNetworkConfig,omitnil,omitempty" name:"ExposeNetworkConfig"`

	// Environment Variables.
	Envs []*EnvVar `json:"Envs,omitnil,omitempty" name:"Envs"`

	// Train tool configuration.
	TrainToolConfig *TrainToolConfig `json:"TrainToolConfig,omitnil,omitempty" name:"TrainToolConfig"`

	// Training Diagnostic Tool Configuration.
	ResourceSupplyAttribute *ResourceSupplyAttribute `json:"ResourceSupplyAttribute,omitnil,omitempty" name:"ResourceSupplyAttribute"`

	// Queue ID.
	Queues []*string `json:"Queues,omitnil,omitempty" name:"Queues"`
}

func (r *CreateTrainingTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTrainingTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "ChargeType")
	delete(f, "ResourceConfigInfos")
	delete(f, "TiProjectId")
	delete(f, "FrameworkName")
	delete(f, "FrameworkVersion")
	delete(f, "FrameworkEnvironment")
	delete(f, "ResourceGroupId")
	delete(f, "Tags")
	delete(f, "ImageInfo")
	delete(f, "CodePackagePath")
	delete(f, "StartCmdInfo")
	delete(f, "TrainingMode")
	delete(f, "DataConfigs")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "Output")
	delete(f, "LogConfig")
	delete(f, "TuningParameters")
	delete(f, "LogEnable")
	delete(f, "Remark")
	delete(f, "DataSource")
	delete(f, "CallbackUrl")
	delete(f, "EncodedStartCmdInfo")
	delete(f, "CodeRepos")
	delete(f, "ExposeNetworkConfig")
	delete(f, "Envs")
	delete(f, "TrainToolConfig")
	delete(f, "ResourceSupplyAttribute")
	delete(f, "Queues")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTrainingTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTrainingTaskResponseParams struct {
	// Training task ID.
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateTrainingTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateTrainingTaskResponseParams `json:"Response"`
}

func (r *CreateTrainingTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTrainingTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CronScaleJob struct {
	// Cron expression, which identifies the task execution time, and is accurate to minutes.
	Schedule *string `json:"Schedule,omitnil,omitempty" name:"Schedule"`

	// Scheduled task name.Note: This field may return null, indicating that no valid values can be obtained.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Number of target instances.Note: This field may return null, indicating that no valid values can be obtained.
	TargetReplicas *int64 `json:"TargetReplicas,omitnil,omitempty" name:"TargetReplicas"`

	// Minimum target.Note: This field may return null, indicating that no valid values can be obtained.
	MinReplicas *int64 `json:"MinReplicas,omitnil,omitempty" name:"MinReplicas"`

	// Maximum target.Note: This field may return null, indicating that no valid values can be obtained.
	MaxReplicas *int64 `json:"MaxReplicas,omitnil,omitempty" name:"MaxReplicas"`

	// Exception periods, defined by Cron expressions, during which tasks are not executed. Up to 3 Cron expressions are supported.Note: This field may return null, indicating that no valid values can be obtained.
	ExcludeDates []*string `json:"ExcludeDates,omitnil,omitempty" name:"ExcludeDates"`
}

type CrossTenantENIInfo struct {
	// Pod IP address.Note: This field may return null, indicating that no valid values can be obtained.
	PrimaryIP *string `json:"PrimaryIP,omitnil,omitempty" name:"PrimaryIP"`

	// Pod port.Note: This field may return null, indicating that no valid values can be obtained.
	Port *string `json:"Port,omitnil,omitempty" name:"Port"`
}

type DataConfig struct {
	// Mapping path.
	MappingPath *string `json:"MappingPath,omitnil,omitempty" name:"MappingPath"`

	// Storage purpose.
	// Valid values: BUILTIN_CODE, BUILTIN_DATA, BUILTIN_MODEL, USER_DATA, USER_CODE, USER_MODEL, OUTPUT, and OTHER.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DataSourceUsage *string `json:"DataSourceUsage,omitnil,omitempty" name:"DataSourceUsage"`

	// DATASET, COS, CFS, CFSTurbo, GooseFSx, HDFS, and WEDATA_HDFS
	// Note: This field may return null, indicating that no valid values can be obtained.
	DataSourceType *string `json:"DataSourceType,omitnil,omitempty" name:"DataSourceType"`

	// Data from the data set.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DataSetSource *DataSetConfig `json:"DataSetSource,omitnil,omitempty" name:"DataSetSource"`

	// Data from COS.
	// Note: This field may return null, indicating that no valid values can be obtained.
	COSSource *CosPathInfo `json:"COSSource,omitnil,omitempty" name:"COSSource"`

	// Data from CFS.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CFSSource *CFSConfig `json:"CFSSource,omitnil,omitempty" name:"CFSSource"`

	// Data from HDFS.
	// Note: This field may return null, indicating that no valid values can be obtained.
	HDFSSource *HDFSConfig `json:"HDFSSource,omitnil,omitempty" name:"HDFSSource"`

	// GooseFS data.
	// Note: This field may return null, indicating that no valid values can be obtained.
	GooseFSSource *GooseFS `json:"GooseFSSource,omitnil,omitempty" name:"GooseFSSource"`

	// TurboFS data.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CFSTurboSource *CFSTurbo `json:"CFSTurboSource,omitnil,omitempty" name:"CFSTurboSource"`

	// Information from local disks.
	// Note: This field may return null, indicating that no valid values can be obtained.
	LocalDiskSource *LocalDisk `json:"LocalDiskSource,omitnil,omitempty" name:"LocalDiskSource"`

	// CBS configuration information.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CBSSource *CBSConfig `json:"CBSSource,omitnil,omitempty" name:"CBSSource"`

	// Host path information.
	HostPathSource *HostPath `json:"HostPathSource,omitnil,omitempty" name:"HostPathSource"`


	PublicDataSource *PublicDataSourceFS `json:"PublicDataSource,omitnil,omitempty" name:"PublicDataSource"`
}

type DataSetConfig struct {
	// Data set ID.
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`
}

// Predefined struct for user
type DescribeModelServiceGroupsRequestParams struct {
	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned results. The default value is 20, and the maximum value is 100.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// The sorting order of the output list. Valid values: ASC (ascending order) and DESC (descending order).
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// Field to sort by. Valid values: CreateTime and UpdateTime.
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`

	// Pagination parameters. Supported filterable field names include:["ClusterId", "ServiceId", "ServiceGroupName", "ServiceGroupId","Status","CreatedBy","ModelVersionId"]
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Tag filtering parameters.
	TagFilters []*TagFilter `json:"TagFilters,omitnil,omitempty" name:"TagFilters"`

	// Service classification.
	//
	// Deprecated: ServiceCategory is deprecated.
	ServiceCategory *string `json:"ServiceCategory,omitnil,omitempty" name:"ServiceCategory"`
}

type DescribeModelServiceGroupsRequest struct {
	*tchttp.BaseRequest
	
	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned results. The default value is 20, and the maximum value is 100.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// The sorting order of the output list. Valid values: ASC (ascending order) and DESC (descending order).
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// Field to sort by. Valid values: CreateTime and UpdateTime.
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`

	// Pagination parameters. Supported filterable field names include:["ClusterId", "ServiceId", "ServiceGroupName", "ServiceGroupId","Status","CreatedBy","ModelVersionId"]
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Tag filtering parameters.
	TagFilters []*TagFilter `json:"TagFilters,omitnil,omitempty" name:"TagFilters"`

	// Service classification.
	ServiceCategory *string `json:"ServiceCategory,omitnil,omitempty" name:"ServiceCategory"`
}

func (r *DescribeModelServiceGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModelServiceGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Order")
	delete(f, "OrderField")
	delete(f, "Filters")
	delete(f, "TagFilters")
	delete(f, "ServiceCategory")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeModelServiceGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModelServiceGroupsResponseParams struct {
	// Number of inference service groups.Note: This field may return null, indicating that no valid values can be obtained.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Service group information.Note: This field may return null, indicating that no valid values can be obtained.
	ServiceGroups []*ServiceGroup `json:"ServiceGroups,omitnil,omitempty" name:"ServiceGroups"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeModelServiceGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeModelServiceGroupsResponseParams `json:"Response"`
}

func (r *DescribeModelServiceGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModelServiceGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EncodedStartCmdInfo struct {
	// Startup command of the task, which is input in base64 format. Note that the complete input of {"StartCmd":"","PsStartCmd":"","WorkerStartCmd":""} is required for conversion.
	StartCmdInfo *string `json:"StartCmdInfo,omitnil,omitempty" name:"StartCmdInfo"`
}

type EnvVar struct {
	// Environment variable key.Note: This field may return null, indicating that no valid values can be obtained.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Environment variable value.Note: This field may return null, indicating that no valid values can be obtained.
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type ExecAction struct {
	// Execution command list.
	Command []*string `json:"Command,omitnil,omitempty" name:"Command"`
}

type ExposeNetworkConfig struct {

	SSHConfig *SSHConfig `json:"SSHConfig,omitnil,omitempty" name:"SSHConfig"`


	ExposePortConfig *ExposePortConfig `json:"ExposePortConfig,omitnil,omitempty" name:"ExposePortConfig"`
}

type ExposePortConfig struct {

	Enable *bool `json:"Enable,omitnil,omitempty" name:"Enable"`


	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`


	ClbId *string `json:"ClbId,omitnil,omitempty" name:"ClbId"`


	ClbHost *string `json:"ClbHost,omitnil,omitempty" name:"ClbHost"`
}

type Filter struct {
	// Filter field name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Filter field values.
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`

	// Whether to enable reverse query.
	Negative *bool `json:"Negative,omitnil,omitempty" name:"Negative"`

	// Whether to enable fuzzy matching.
	Fuzzy *bool `json:"Fuzzy,omitnil,omitempty" name:"Fuzzy"`
}

type GooseFS struct {
	// GooseFS instance ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// GooseFS type, including GooseFS and GooseFSx.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Path to mount the GooseFSx instance.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// GooseFS namespace.
	// Note: This field may return null, indicating that no valid values can be obtained.
	NameSpace *string `json:"NameSpace,omitnil,omitempty" name:"NameSpace"`
}

type GooseFSx struct {
	// GooseFSx instance ID.
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// Path to mount the GooseFSx instance.
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`
}

type GpuDetail struct {
	// GPU card type. Enumeration values: V100, A100, T4.Note: This field may return null, indicating that no valid values can be obtained.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// GPU card quantity, in 1/100 cards. For example, 100 represents 1 card.Note: This field may return null, indicating that no valid values can be obtained.
	Value *uint64 `json:"Value,omitnil,omitempty" name:"Value"`
}

type HDFSConfig struct {
	// Cluster instance ID, such as emr-xxxxxxxx.
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// Path.
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`
}

type HTTPGetAction struct {
	// HTTP path.
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// Called port.
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`
}

type HealthProbe struct {
	// Liveness probe.
	LivenessProbe *Probe `json:"LivenessProbe,omitnil,omitempty" name:"LivenessProbe"`

	// Readiness probe.
	ReadinessProbe *Probe `json:"ReadinessProbe,omitnil,omitempty" name:"ReadinessProbe"`

	// Startup probe.
	StartupProbe *Probe `json:"StartupProbe,omitnil,omitempty" name:"StartupProbe"`
}

type HorizontalPodAutoscaler struct {
	// Minimum number of instances.Note: This field may return null, indicating that no valid values can be obtained.
	MinReplicas *int64 `json:"MinReplicas,omitnil,omitempty" name:"MinReplicas"`

	// Maximum number of instances.Note: This field may return null, indicating that no valid values can be obtained.
	MaxReplicas *int64 `json:"MaxReplicas,omitnil,omitempty" name:"MaxReplicas"`

	// Supported."gpu-util": GPU utilization; value range: 10-100. "cpu-util": CPU utilization; value range: 10-100. "memory-util": memory utilization; value range: 10-100. "service-qps": the QPS value of single instances; value range: 1-5000."concurrency-util": the number of concurrent requests of single instances. Value range: 1-100000.Note: This field may return null, indicating that no valid values can be obtained.
	HpaMetrics []*Option `json:"HpaMetrics,omitnil,omitempty" name:"HpaMetrics"`

	// Scale-out cooldown period, in seconds.
	ScaleUpStabilizationWindowSeconds *int64 `json:"ScaleUpStabilizationWindowSeconds,omitnil,omitempty" name:"ScaleUpStabilizationWindowSeconds"`

	// Scale-in cooldown period, in seconds.
	ScaleDownStabilizationWindowSeconds *int64 `json:"ScaleDownStabilizationWindowSeconds,omitnil,omitempty" name:"ScaleDownStabilizationWindowSeconds"`
}

type HostPath struct {
	// Host path to be mounted.
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`
}

type ImageInfo struct {
	// Image type. Valid values: TCR ( which indicates a Tencent Container Registry (TCR) image), CCR (which indicates a TCR Personal Edition image), PreSet (which indicates a platform preset image), and CUSTOM (which indicates a third-party custom image).
	ImageType *string `json:"ImageType,omitnil,omitempty" name:"ImageType"`

	// Image address.
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// Region corresponding to the TCR image.Note: This field may return null, indicating that no valid values can be obtained.
	RegistryRegion *string `json:"RegistryRegion,omitnil,omitempty" name:"RegistryRegion"`

	// Instance ID corresponding to the TCR image.Note: This field may return null, indicating that no valid values can be obtained.
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Whether to allow exporting all content.Note: This field may return null, indicating that no valid values can be obtained.
	AllowSaveAllContent *bool `json:"AllowSaveAllContent,omitnil,omitempty" name:"AllowSaveAllContent"`

	// Image name.Note: This field may return null, indicating that no valid values can be obtained.
	ImageName *string `json:"ImageName,omitnil,omitempty" name:"ImageName"`

	// Whether to support data generation.Note: This field may return null, indicating that no valid values can be obtained.
	SupportDataPipeline *bool `json:"SupportDataPipeline,omitnil,omitempty" name:"SupportDataPipeline"`


	ImageSecret *ImageSecret `json:"ImageSecret,omitnil,omitempty" name:"ImageSecret"`
}

type ImageSecret struct {

	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`


	Username *string `json:"Username,omitnil,omitempty" name:"Username"`


	Password *string `json:"Password,omitnil,omitempty" name:"Password"`


	SecretId *string `json:"SecretId,omitnil,omitempty" name:"SecretId"`
}

type InferCodeInfo struct {
	// Details of Cloud Object Storage (COS) where the inference code is located.Note: This field may return null, indicating that no valid values can be obtained.
	CosPathInfo *CosPathInfo `json:"CosPathInfo,omitnil,omitempty" name:"CosPathInfo"`
}

type LocalDisk struct {
	// Node ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Local path.
	// Note: This field may return null, indicating that no valid values can be obtained.
	LocalPath *string `json:"LocalPath,omitnil,omitempty" name:"LocalPath"`
}

type LogConfig struct {
	// Logs should be shipped to a Cloud Log Service (CLS) log set.Note: This field may return null, indicating that no valid values can be obtained.
	LogsetId *string `json:"LogsetId,omitnil,omitempty" name:"LogsetId"`

	// Logs should be shipped to a CLS topic.Note: This field may return null, indicating that no valid values can be obtained.
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`
}

type ModelInfo struct {
	// The model version ID is returned by the DescribeTrainingModelVersion API when querying the model.Enter the task ID of the Automated Machine Learning (AutoML) model.
	ModelVersionId *string `json:"ModelVersionId,omitnil,omitempty" name:"ModelVersionId"`

	// Model ID.
	ModelId *string `json:"ModelId,omitnil,omitempty" name:"ModelId"`

	// Model name.
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// Model version.
	ModelVersion *string `json:"ModelVersion,omitnil,omitempty" name:"ModelVersion"`

	// Model source.
	ModelSource *string `json:"ModelSource,omitnil,omitempty" name:"ModelSource"`

	// COS path information.
	CosPathInfo *CosPathInfo `json:"CosPathInfo,omitnil,omitempty" name:"CosPathInfo"`

	// GooseFSx configuration. This parameter takes effect if ModelSource is GooseFSx.
	// Note: This field may return null, indicating that no valid values can be obtained.
	GooseFSx *GooseFSx `json:"GooseFSx,omitnil,omitempty" name:"GooseFSx"`

	// Algorithm framework corresponding to the model (reserved field).Note: This field may return null, indicating that no valid values can be obtained.
	AlgorithmFramework *string `json:"AlgorithmFramework,omitnil,omitempty" name:"AlgorithmFramework"`

	// Default: NORMAL; accelerated model: ACCELERATE; automatic learning model: AUTO_ML.Note: This field may return null, indicating that no valid values can be obtained.
	ModelType *string `json:"ModelType,omitnil,omitempty" name:"ModelType"`

	// Model format.Note: This field may return null, indicating that no valid values can be obtained.
	ModelFormat *string `json:"ModelFormat,omitnil,omitempty" name:"ModelFormat"`

	// Whether it is a private LLM.Note: This field may return null, indicating that no valid values can be obtained.
	IsPrivateModel *bool `json:"IsPrivateModel,omitnil,omitempty" name:"IsPrivateModel"`

	// Model category. Valid values: MultiModal (multi-modal) and LLM (text LLM).
	ModelCategory *string `json:"ModelCategory,omitnil,omitempty" name:"ModelCategory"`

	// Data source configurations.
	PublicDataSource *PublicDataSourceFS `json:"PublicDataSource,omitnil,omitempty" name:"PublicDataSource"`
}

type NumOrPercent struct {
	// Valid values: Num and Percent, which indicate quantity and percentage respectively. The default value is Num.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Numeric value.
	Value *int64 `json:"Value,omitnil,omitempty" name:"Value"`
}

type Option struct {
	// Metric name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Metric value.
	Value *int64 `json:"Value,omitnil,omitempty" name:"Value"`
}

type Pod struct {
	// Pod name.Note: This field may return null, indicating that no valid values can be obtained.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Unique ID of the Pod.Note: This field may return null, indicating that no valid values can be obtained.
	Uid *string `json:"Uid,omitnil,omitempty" name:"Uid"`

	// Service payment mode.Note: This field may return null, indicating that no valid values can be obtained.
	ChargeType *string `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`

	// Pod status.Note: This field may return null, indicating that no valid values can be obtained.
	Phase *string `json:"Phase,omitnil,omitempty" name:"Phase"`

	// Pod IP address.Note: This field may return null, indicating that no valid values can be obtained.
	IP *string `json:"IP,omitnil,omitempty" name:"IP"`

	// Pod creation time.Note: This field may return null, indicating that no valid values can be obtained.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Container list.Note: This field may return null, indicating that no valid values can be obtained.
	//
	// Deprecated: Containers is deprecated.
	Containers *Container `json:"Containers,omitnil,omitempty" name:"Containers"`

	// Container list.Note: This field may return null, indicating that no valid values can be obtained.
	ContainerInfos []*Container `json:"ContainerInfos,omitnil,omitempty" name:"ContainerInfos"`

	// Container calling information.Note: This field may return null, indicating that no valid values can be obtained.
	CrossTenantENIInfo *CrossTenantENIInfo `json:"CrossTenantENIInfo,omitnil,omitempty" name:"CrossTenantENIInfo"`

	// Instance status information.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Instance scheduling start time.
	StartScheduleTime *string `json:"StartScheduleTime,omitnil,omitempty" name:"StartScheduleTime"`

	// Supplemental instance status information.
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// Node IP address of the current instance.
	NodeIP *string `json:"NodeIP,omitnil,omitempty" name:"NodeIP"`

	// Node ID of the current instance.
	NodeId *string `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// Resource group ID to which the instance belonged.
	ResourceGroupId *string `json:"ResourceGroupId,omitnil,omitempty" name:"ResourceGroupId"`

	// Resource group name.
	ResourceGroupName *string `json:"ResourceGroupName,omitnil,omitempty" name:"ResourceGroupName"`

	// Resource usage information of the instance.
	ResourceInfo *ResourceInfo `json:"ResourceInfo,omitnil,omitempty" name:"ResourceInfo"`
}

type PodSSHInfo struct {
	// IP address of the Pod.
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// SSH port of the Pod.
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// SSH access command.
	LoginCommand *string `json:"LoginCommand,omitnil,omitempty" name:"LoginCommand"`
}

type Probe struct {
	// Probe action.
	ProbeAction *ProbeAction `json:"ProbeAction,omitnil,omitempty" name:"ProbeAction"`

	// Delay in waiting for a service startup.
	InitialDelaySeconds *int64 `json:"InitialDelaySeconds,omitnil,omitempty" name:"InitialDelaySeconds"`

	// Polling check interval.
	PeriodSeconds *int64 `json:"PeriodSeconds,omitnil,omitempty" name:"PeriodSeconds"`

	// Check timeout duration.
	TimeoutSeconds *int64 `json:"TimeoutSeconds,omitnil,omitempty" name:"TimeoutSeconds"`

	// Number of acknowledged failed detections.
	FailureThreshold *int64 `json:"FailureThreshold,omitnil,omitempty" name:"FailureThreshold"`

	// Number of acknowledged successful detections. The default values for readiness, liveness, and startup statuses are 3, 1, and 1.
	SuccessThreshold *int64 `json:"SuccessThreshold,omitnil,omitempty" name:"SuccessThreshold"`
}

type ProbeAction struct {
	// HTTP GET action.
	HTTPGet *HTTPGetAction `json:"HTTPGet,omitnil,omitempty" name:"HTTPGet"`

	// Executes check command action.
	Exec *ExecAction `json:"Exec,omitnil,omitempty" name:"Exec"`

	// TCP Socket check action.
	TCPSocket *TCPSocketAction `json:"TCPSocket,omitnil,omitempty" name:"TCPSocket"`

	// Probe type. The default value is HTTPGet. Valid values: HTTPGet, Exec, and TCPSocket.
	ActionType *string `json:"ActionType,omitnil,omitempty" name:"ActionType"`
}

type PublicDataSourceFS struct {
	// Data source ID.
	DataSourceId *string `json:"DataSourceId,omitnil,omitempty" name:"DataSourceId"`

	// Relative subpath to the data source.
	SubPath *string `json:"SubPath,omitnil,omitempty" name:"SubPath"`
}

type RDMAConfig struct {
	// Whether to enable RDMA.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Enable *bool `json:"Enable,omitnil,omitempty" name:"Enable"`
}

type ResourceConfigInfo struct {
	// Role. For example, PS, WORKER, DRIVER, and EXECUTOR.
	Role *string `json:"Role,omitnil,omitempty" name:"Role"`

	// Number of CPU cores, which is required to be configured when resource groups are used. Unit: 1/1000, where 1000 represents 1 core.
	Cpu *uint64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// Memory size, in MB. This parameter needs to be configured when resource groups are used.
	Memory *uint64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// GPU card type, which is required to be configured when resource groups are used.
	GpuType *string `json:"GpuType,omitnil,omitempty" name:"GpuType"`

	// Number of GPU cards, which is required to be configured when resource groups are used. Unit: 1/100, where 100 represents 1 card.
	Gpu *uint64 `json:"Gpu,omitnil,omitempty" name:"Gpu"`

	// CVM instance specification ID.
	// CVM instance specification (for postpaid billing). Valid values:
	// TI.S.LARGE.POST: 4C8G 
	// TI.S.2XLARGE16.POST:  8C16G 
	// TI.S.2XLARGE32.POST:  8C32G 
	// TI.S.4XLARGE32.POST:  16C32G
	// TI.S.4XLARGE64.POST:  16C64G
	// TI.S.6XLARGE48.POST:  24C48G
	// TI.S.6XLARGE96.POST:  24C96G
	// TI.S.8XLARGE64.POST:  32C64G
	// TI.S.8XLARGE128.POST : 32C128G
	// TI.GN10.2XLARGE40.POST: 8C40G V100*1 
	// TI.GN10.5XLARGE80.POST:  18C80G V100*2 
	// TI.GN10.10XLARGE160.POST :  32C160G V100*4
	// TI.GN10.20XLARGE320.POST :  72C320G V100*8
	// TI.GN7.8XLARGE128.POST: 32C128G T4*1 
	// TI.GN7.10XLARGE160.POST: 40C160G T4*2 
	// TI.GN7.20XLARGE320.POST: 80C32
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// Number of compute nodes.
	InstanceNum *uint64 `json:"InstanceNum,omitnil,omitempty" name:"InstanceNum"`

	// CVM instance specification name.
	// CVM instance specification (for postpaid billing). Valid values:
	// 4C8G 
	// 8C16G 
	// 8C32G 
	// 16C32G
	// 6C64G
	// 24C48G
	// 24C96G
	// 32C64G
	// 32C128G
	// 8C40G V100*1 
	// 8C80G V100*2 
	// 32C160G V100*4
	// 72C320G V100*8
	// 32C128G T4*1 
	// 40C160G T4*2 
	// 80C32
	InstanceTypeAlias *string `json:"InstanceTypeAlias,omitnil,omitempty" name:"InstanceTypeAlias"`

	// RDMA configuration.
	// Note: This field may return null, indicating that no valid values can be obtained.
	RDMAConfig *RDMAConfig `json:"RDMAConfig,omitnil,omitempty" name:"RDMAConfig"`
}

type ResourceGroupInfo struct {
	// Resource group ID.
	ResourceGroupId *string `json:"ResourceGroupId,omitnil,omitempty" name:"ResourceGroupId"`

	// Resource group name.
	ResourceGroupName *string `json:"ResourceGroupName,omitnil,omitempty" name:"ResourceGroupName"`
}

type ResourceInfo struct {
	// Processor resource, in 1/1000 cores.Note: This field may return null, indicating that no valid values can be obtained.
	Cpu *uint64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// Memory resource, in MB.Note: This field may return null, indicating that no valid values can be obtained.
	Memory *uint64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// Number of GPU card resources, in 0.01 units of GpuType.Gpu=100 indicates the use of "1" GPU card. However, this "1" card could refer to a virtualized 1/4 card or a full physical card, depending on the instance type.Example 1: If the instance type includes 1 virtual GPU card, and each virtual GPU card corresponds to 1/4 of a physical T4 card, then GpuType=T4, Gpu=100, and RealGpu=25.Example 2: If the instance type includes 4 full GPU cards, and each card corresponds to 1 physical T4 card, then GpuType=T4, Gpu=400, and RealGpu=400.Note: This field may return null, indicating that no valid values can be obtained.
	Gpu *uint64 `json:"Gpu,omitnil,omitempty" name:"Gpu"`

	// GPU card model. Valid values: T4 and V100. It only displays the current GPU card model. If multiple types of cards are used simultaneously, see the value of RealGpuDetailSet.Note: This field may return null, indicating that no valid values can be obtained.
	GpuType *string `json:"GpuType,omitnil,omitempty" name:"GpuType"`

	// It is not required for creation or update operations. This field is used for display only.The actual GPU card resources for postpaid instances using fractional GPU cards. This value represents the total number of actual physical GPU cards consumed.RealGpu=100 indicates the consumption of 1 GPU card. Depending on the actual instance type, this may represent: 4 instances each using a 1/4 card, 2 instances each using a 1/2 card, or 1 instance using a full card.
	RealGpu *uint64 `json:"RealGpu,omitnil,omitempty" name:"RealGpu"`

	// It is not required for creation or update operations. This field is used for display only. It involves detailed GPU usage information.
	RealGpuDetailSet []*GpuDetail `json:"RealGpuDetailSet,omitnil,omitempty" name:"RealGpuDetailSet"`

	// Indicates whether to enable RDMA.
	// Note: This field may return null, indicating that no valid values can be obtained.
	EnableRDMA *bool `json:"EnableRDMA,omitnil,omitempty" name:"EnableRDMA"`


	RootDisk *uint64 `json:"RootDisk,omitnil,omitempty" name:"RootDisk"`


	DataDisk *uint64 `json:"DataDisk,omitnil,omitempty" name:"DataDisk"`
}

type ResourceSupplyAttribute struct {

	SupplyType *string `json:"SupplyType,omitnil,omitempty" name:"SupplyType"`


	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`
}

type RollingUpdate struct {
	// Maximum unavailability for rolling updates.
	MaxUnavailable *NumOrPercent `json:"MaxUnavailable,omitnil,omitempty" name:"MaxUnavailable"`

	// Maximum number of new instances during rolling updates. 
	MaxSurge *NumOrPercent `json:"MaxSurge,omitnil,omitempty" name:"MaxSurge"`
}

type SSHConfig struct {
	// Whether to enable SSH.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Enable *bool `json:"Enable,omitnil,omitempty" name:"Enable"`

	// Public key information.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PublicKey *string `json:"PublicKey,omitnil,omitempty" name:"PublicKey"`

	// Port number.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// Login command.
	// Note: This field may return null, indicating that no valid values can be obtained.
	LoginCommand *string `json:"LoginCommand,omitnil,omitempty" name:"LoginCommand"`

	// Whether to change the login address.
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsAddressChanged *bool `json:"IsAddressChanged,omitnil,omitempty" name:"IsAddressChanged"`

	// Pod access information.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PodSSHInfo *PodSSHInfo `json:"PodSSHInfo,omitnil,omitempty" name:"PodSSHInfo"`
}

type ScheduledAction struct {
	// Whether to stop the service on schedule. Valid values: true and false. If the value is true, ScheduleStopTime is required. If the value is false, ScheduleStopTime does not take effect.
	ScheduleStop *bool `json:"ScheduleStop,omitnil,omitempty" name:"ScheduleStop"`

	// Time to execute the scheduled stop. Format: "2022-01-26T19:46:22+08:00".
	ScheduleStopTime *string `json:"ScheduleStopTime,omitnil,omitempty" name:"ScheduleStopTime"`
}

type SchedulingPolicy struct {
	// Whether to enable cross-resource-group scheduling.
	CrossResourceGroupScheduling *bool `json:"CrossResourceGroupScheduling,omitnil,omitempty" name:"CrossResourceGroupScheduling"`
}

type Service struct {
	// Service group ID.
	ServiceGroupId *string `json:"ServiceGroupId,omitnil,omitempty" name:"ServiceGroupId"`

	// Service ID.
	ServiceId *string `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`

	// Service group name.
	ServiceGroupName *string `json:"ServiceGroupName,omitnil,omitempty" name:"ServiceGroupName"`

	// Service description.Note: This field may return null, indicating that no valid values can be obtained.
	ServiceDescription *string `json:"ServiceDescription,omitnil,omitempty" name:"ServiceDescription"`

	// Service details.Note: This field may return null, indicating that no valid values can be obtained.
	ServiceInfo *ServiceInfo `json:"ServiceInfo,omitnil,omitempty" name:"ServiceInfo"`

	// Cluster ID.Note: This field may return null, indicating that no valid values can be obtained.
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Region.Note: This field may return null, indicating that no valid values can be obtained.
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Namespace.Note: This field may return null, indicating that no valid values can be obtained.
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// Billing type.Note: This field may return null, indicating that no valid values can be obtained.
	ChargeType *string `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`

	// Resource group ID for yearly/monthly subscription services. The value is null for pay-as-you-go services.Note: This field may return null, indicating that no valid values can be obtained.
	ResourceGroupId *string `json:"ResourceGroupId,omitnil,omitempty" name:"ResourceGroupId"`

	// Resource group name corresponding to the yearly/monthly subscription service.Note: This field may return null, indicating that no valid values can be obtained.
	ResourceGroupName *string `json:"ResourceGroupName,omitnil,omitempty" name:"ResourceGroupName"`

	// Tag of the service.Note: This field may return null, indicating that no valid values can be obtained.
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Name of the ingress where the service is located.Note: This field may return null, indicating that no valid values can be obtained.
	IngressName *string `json:"IngressName,omitnil,omitempty" name:"IngressName"`

	// Creator.Note: This field may return null, indicating that no valid values can be obtained.
	CreatedBy *string `json:"CreatedBy,omitnil,omitempty" name:"CreatedBy"`

	// Creation time.Note: This field may return null, indicating that no valid values can be obtained.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Update time.Note: This field may return null, indicating that no valid values can be obtained.
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// Root account.Note: This field may return null, indicating that no valid values can be obtained.
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// Sub-account.Note: This field may return null, indicating that no valid values can be obtained.
	SubUin *string `json:"SubUin,omitnil,omitempty" name:"SubUin"`

	// app_id
	// Note: This field may return null, indicating that no valid values can be obtained.
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// Operational status of the service.Note: This field may return null, indicating that no valid values can be obtained.
	BusinessStatus *string `json:"BusinessStatus,omitnil,omitempty" name:"BusinessStatus"`

	// Deprecated. See the corresponding field in ServiceInfo.Note: This field may return null, indicating that no valid values can be obtained.
	//
	// Deprecated: ServiceLimit is deprecated.
	ServiceLimit *ServiceLimit `json:"ServiceLimit,omitnil,omitempty" name:"ServiceLimit"`

	// Deprecated. See the corresponding field in ServiceInfo.Note: This field may return null, indicating that no valid values can be obtained.
	//
	// Deprecated: ScheduledAction is deprecated.
	ScheduledAction *ScheduledAction `json:"ScheduledAction,omitnil,omitempty" name:"ScheduledAction"`

	// Cause for service creation failure. The default value of this field is CREATE_SUCCEED upon successful creation.Note: This field may return null, indicating that no valid values can be obtained.
	CreateFailedReason *string `json:"CreateFailedReason,omitnil,omitempty" name:"CreateFailedReason"`

	// Service status.CREATING: creating.CREATE_FAILED: creation failed.Normal: running.Stopped: stopped.Stopping: stopping.Abnormal: error.Pending: starting.Waiting: getting ready.Note: This field may return null, indicating that no valid values can be obtained.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Billing information.Note: This field may return null, indicating that no valid values can be obtained.
	BillingInfo *string `json:"BillingInfo,omitnil,omitempty" name:"BillingInfo"`

	// Model weight.Note: This field may return null, indicating that no valid values can be obtained.
	Weight *int64 `json:"Weight,omitnil,omitempty" name:"Weight"`

	// Creation source of the service.AUTO_ML: comes from one-click release of automatic learning.DEFAULT: other sources.Note: This field may return null, indicating that no valid values can be obtained.
	CreateSource *string `json:"CreateSource,omitnil,omitempty" name:"CreateSource"`

	// Version number.Note: This field may return null, indicating that no valid values can be obtained.
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// The latest version number of services under a service group.Note: This field may return null, indicating that no valid values can be obtained.
	LatestVersion *string `json:"LatestVersion,omitnil,omitempty" name:"LatestVersion"`

	// Resource group category. Valid values: NORMAL (hosting) and SW (half-hosting).Note: This field may return null, indicating that no valid values can be obtained.
	ResourceGroupSWType *string `json:"ResourceGroupSWType,omitnil,omitempty" name:"ResourceGroupSWType"`

	// Archiving status of the service. Valid values: Waiting (pending archiving) and Archived (archived).Note: This field may return null, indicating that no valid values can be obtained.
	ArchiveStatus *string `json:"ArchiveStatus,omitnil,omitempty" name:"ArchiveStatus"`

	// Deployment type of the service. Valid values: STANDARD (standard deployment) and DIST (multi-machine distributed deployment). The default value is STANDARD.Note: This field may return null, indicating that no valid values can be obtained.
	DeployType *string `json:"DeployType,omitnil,omitempty" name:"DeployType"`

	// Number of instances per replica. This parameter is valid only when the deployment type is DIST. Default value: 1.Note: This field may return null, indicating that no valid values can be obtained.
	InstancePerReplicas *string `json:"InstancePerReplicas,omitnil,omitempty" name:"InstancePerReplicas"`

	// Source for monitoring queries.Enumeration value. May differ from CreateSource in certain scenarios. This field is designed to be compatible.
	MonitorSource *string `json:"MonitorSource,omitnil,omitempty" name:"MonitorSource"`

	// Sub-account name of the service creator.
	SubUinName *string `json:"SubUinName,omitnil,omitempty" name:"SubUinName"`

	// Scheduling policy of the service.
	SchedulingPolicy *SchedulingPolicy `json:"SchedulingPolicy,omitnil,omitempty" name:"SchedulingPolicy"`

	// External resource group information, indicating which resources are borrowed from resource groups.
	ExternalResourceGroups []*ResourceGroupInfo `json:"ExternalResourceGroups,omitnil,omitempty" name:"ExternalResourceGroups"`
}

type ServiceEIP struct {
	// Whether to enable access from the TI-ONE private network to external resources.Note: This field may return null, indicating that no valid values can be obtained.
	EnableEIP *bool `json:"EnableEIP,omitnil,omitempty" name:"EnableEIP"`

	// User VPC ID.Note: This field may return null, indicating that no valid values can be obtained.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// User subnet ID.Note: This field may return null, indicating that no valid values can be obtained.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`
}

type ServiceGroup struct {
	// Service group ID.
	ServiceGroupId *string `json:"ServiceGroupId,omitnil,omitempty" name:"ServiceGroupId"`

	// Service group name.
	ServiceGroupName *string `json:"ServiceGroupName,omitnil,omitempty" name:"ServiceGroupName"`

	// Creator.
	CreatedBy *string `json:"CreatedBy,omitnil,omitempty" name:"CreatedBy"`

	// Creation time.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Update time.
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// Root account.
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// Total number of services in the service group.Note: This field may return null, indicating that no valid values can be obtained.
	ServiceCount *uint64 `json:"ServiceCount,omitnil,omitempty" name:"ServiceCount"`

	// Number of running services in the service group.Note: This field may return null, indicating that no valid values can be obtained.
	RunningServiceCount *uint64 `json:"RunningServiceCount,omitnil,omitempty" name:"RunningServiceCount"`

	// Service description.Note: This field may return null, indicating that no valid values can be obtained.
	Services []*Service `json:"Services,omitnil,omitempty" name:"Services"`

	// Service group status, which aligns with service status.CREATING: creating.CREATE_FAILED: creation failed.Normal: running.Stopped: stopped.Stopping: stopping.Abnormal: error.Pending: starting.Waiting: getting ready.Note: This field may return null, indicating that no valid values can be obtained.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Service group tags.Note: This field may return null, indicating that no valid values can be obtained.
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// The latest version in the service group.Note: This field may return null, indicating that no valid values can be obtained.
	LatestVersion *string `json:"LatestVersion,omitnil,omitempty" name:"LatestVersion"`

	// Operational status of the service.CREATING: creating.CREATE_FAILED: creation failed.ARREARS_STOP: service suspended due to overdue payments.BILLING: billing.WHITELIST_USING: allowlist feature is in trial.WHITELIST_STOP: insufficient allowlist quota.Note: This field may return null, indicating that no valid values can be obtained.
	BusinessStatus *string `json:"BusinessStatus,omitnil,omitempty" name:"BusinessStatus"`

	// Billing information of the service.Note: This field may return null, indicating that no valid values can be obtained.
	BillingInfo *string `json:"BillingInfo,omitnil,omitempty" name:"BillingInfo"`

	// Creation source of the service.Note: This field may return null, indicating that no valid values can be obtained.
	CreateSource *string `json:"CreateSource,omitnil,omitempty" name:"CreateSource"`

	// Weight update status of the service group.UPDATING: updating.UPDATED: updated successfully.UPDATE FAILED: failed to update.Note: This field may return null, indicating that no valid values can be obtained.
	WeightUpdateStatus *string `json:"WeightUpdateStatus,omitnil,omitempty" name:"WeightUpdateStatus"`

	// Number of running Pods in the service group.Note: This field may return null, indicating that no valid values can be obtained.
	ReplicasCount *uint64 `json:"ReplicasCount,omitnil,omitempty" name:"ReplicasCount"`

	// Expected number of Pods under the service group.Note: This field may return null, indicating that no valid values can be obtained.
	AvailableReplicasCount *uint64 `json:"AvailableReplicasCount,omitnil,omitempty" name:"AvailableReplicasCount"`

	// Service group's subuin.
	SubUin *string `json:"SubUin,omitnil,omitempty" name:"SubUin"`

	// Service group's app_id.
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// Whether to enable authentication.
	AuthorizationEnable *bool `json:"AuthorizationEnable,omitnil,omitempty" name:"AuthorizationEnable"`

	// List of throttling authentication tokens.
	AuthTokens []*AuthToken `json:"AuthTokens,omitnil,omitempty" name:"AuthTokens"`

	// Field for monitoring creation source.
	MonitorSource *string `json:"MonitorSource,omitnil,omitempty" name:"MonitorSource"`

	// Nickname of the sub-user.
	SubUinName *string `json:"SubUinName,omitnil,omitempty" name:"SubUinName"`
}

type ServiceInfo struct {
	// Expected number of running Pods. The instance is 0 when the stop status is reached.Corresponding relationships under different billing and scaling modes are as follows.PREPAID and POSTPAID_BY_HOUR:Corresponding number of instances in the manual scaling mode.Corresponding number of instances based on the default time-based policy in the auto-scaling mode.HYBRID_PAID:
	// Corresponding number of instances for postpaid instances in the manual scaling mode.Corresponding number of instances under the default time-based policy for postpaid instances in the auto-scaling mode.Note: This field may return null, indicating that no valid values can be obtained.
	Replicas *int64 `json:"Replicas,omitnil,omitempty" name:"Replicas"`

	// Image information.Note: This field may return null, indicating that no valid values can be obtained.
	ImageInfo *ImageInfo `json:"ImageInfo,omitnil,omitempty" name:"ImageInfo"`

	// Environment variables.Note: This field may return null, indicating that no valid values can be obtained.
	Env []*EnvVar `json:"Env,omitnil,omitempty" name:"Env"`

	// Resource information.Note: This field may return null, indicating that no valid values can be obtained.
	Resources *ResourceInfo `json:"Resources,omitnil,omitempty" name:"Resources"`

	// Type specifications corresponding to the postpaid instance.Note: This field may return null, indicating that no valid values can be obtained.
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// Model information.Note: This field may return null, indicating that no valid values can be obtained.
	ModelInfo *ModelInfo `json:"ModelInfo,omitnil,omitempty" name:"ModelInfo"`

	// Whether to enable logs.Note: This field may return null, indicating that no valid values can be obtained.
	LogEnable *bool `json:"LogEnable,omitnil,omitempty" name:"LogEnable"`

	// Log configurations.Note: This field may return null, indicating that no valid values can be obtained.
	LogConfig *LogConfig `json:"LogConfig,omitnil,omitempty" name:"LogConfig"`

	// Whether to enable authentication.Note: This field may return null, indicating that no valid values can be obtained.
	AuthorizationEnable *bool `json:"AuthorizationEnable,omitnil,omitempty" name:"AuthorizationEnable"`

	// HPA configurations.Note: This field may return null, indicating that no valid values can be obtained.
	HorizontalPodAutoscaler *HorizontalPodAutoscaler `json:"HorizontalPodAutoscaler,omitnil,omitempty" name:"HorizontalPodAutoscaler"`

	// Description of the service status.Note: This field may return null, indicating that no valid values can be obtained.
	Status *WorkloadStatus `json:"Status,omitnil,omitempty" name:"Status"`

	// Weight.Note: This field may return null, indicating that no valid values can be obtained.
	Weight *uint64 `json:"Weight,omitnil,omitempty" name:"Weight"`

	// Total resources.Note: This field may return null, indicating that no valid values can be obtained.
	ResourceTotal *ResourceInfo `json:"ResourceTotal,omitnil,omitempty" name:"ResourceTotal"`

	// Number of historical instances.Note: This field may return null, indicating that no valid values can be obtained.
	OldReplicas *int64 `json:"OldReplicas,omitnil,omitempty" name:"OldReplicas"`

	// This parameter is valid when the billing mode is HYBRID_PAID, and is used to identify the number of prepaid instances in the hybrid billing mode. The default value is 1 if this parameter is left unspecified.Note: This field may return null, indicating that no valid values can be obtained.
	HybridBillingPrepaidReplicas *int64 `json:"HybridBillingPrepaidReplicas,omitnil,omitempty" name:"HybridBillingPrepaidReplicas"`

	// Number of instances during the historical HYBRID_PAID period. The user restores services.Note: This field may return null, indicating that no valid values can be obtained.
	OldHybridBillingPrepaidReplicas *int64 `json:"OldHybridBillingPrepaidReplicas,omitnil,omitempty" name:"OldHybridBillingPrepaidReplicas"`

	// Whether to enable hot update for the model. By default, hot update is disabled.Note: This field may return null, indicating that no valid values can be obtained.
	ModelHotUpdateEnable *bool `json:"ModelHotUpdateEnable,omitnil,omitempty" name:"ModelHotUpdateEnable"`

	// Service specification alias.
	InstanceAlias *string `json:"InstanceAlias,omitnil,omitempty" name:"InstanceAlias"`

	// Instance quantity adjusting mode. Defaults to manual.Supported valid values: AUTO (automatic), MANUAL (manual).Note: This field may return null, indicating that no valid values can be obtained.
	ScaleMode *string `json:"ScaleMode,omitnil,omitempty" name:"ScaleMode"`

	// Scheduled scaling task.Note: This field may return null, indicating that no valid values can be obtained.
	CronScaleJobs []*CronScaleJob `json:"CronScaleJobs,omitnil,omitempty" name:"CronScaleJobs"`

	// Scheduled scaling policy.Note: This field may return null, indicating that no valid values can be obtained.
	ScaleStrategy *string `json:"ScaleStrategy,omitnil,omitempty" name:"ScaleStrategy"`

	// Configurations of the scheduled stop.Note: This field may return null, indicating that no valid values can be obtained.
	ScheduledAction *ScheduledAction `json:"ScheduledAction,omitnil,omitempty" name:"ScheduledAction"`

	// Instance list.Note: This field may return null, indicating that no valid values can be obtained.
	//
	// Deprecated: PodList is deprecated.
	PodList []*string `json:"PodList,omitnil,omitempty" name:"PodList"`

	// Pod list information.Note: This field may return null, indicating that no valid values can be obtained.
	//
	// Deprecated: Pods is deprecated.
	Pods *Pod `json:"Pods,omitnil,omitempty" name:"Pods"`

	// Pod list information.Note: This field may return null, indicating that no valid values can be obtained.
	PodInfos []*Pod `json:"PodInfos,omitnil,omitempty" name:"PodInfos"`

	// Configurations related to speed limit and throttling of services.Note: This field may return null, indicating that no valid values can be obtained.
	ServiceLimit *ServiceLimit `json:"ServiceLimit,omitnil,omitempty" name:"ServiceLimit"`

	// Whether to enable model acceleration, which is only valid for models in the StableDiffusion (dynamic acceleration) format.Note: This field may return null, indicating that no valid values can be obtained.
	ModelTurboEnable *bool `json:"ModelTurboEnable,omitnil,omitempty" name:"ModelTurboEnable"`

	// Mounting.Note: This field may return null, indicating that no valid values can be obtained.
	VolumeMount *VolumeMount `json:"VolumeMount,omitnil,omitempty" name:"VolumeMount"`

	// Inference code information.Note: This field may return null, indicating that no valid values can be obtained.
	InferCodeInfo *InferCodeInfo `json:"InferCodeInfo,omitnil,omitempty" name:"InferCodeInfo"`

	// Service startup command.Note: This field may return null, indicating that no valid values can be obtained.
	Command *string `json:"Command,omitnil,omitempty" name:"Command"`

	// Settings of enabling the TI-ONE private network to access external resources.Note: This field may return null, indicating that no valid values can be obtained.
	ServiceEIP *ServiceEIP `json:"ServiceEIP,omitnil,omitempty" name:"ServiceEIP"`

	// Service port, with the default value of 8501.Note: This field may return null, indicating that no valid values can be obtained.
	ServicePort *int64 `json:"ServicePort,omitnil,omitempty" name:"ServicePort"`

	// Graceful exit time limit of the service, in seconds. Default value: 30. Minimum value: 1.
	TerminationGracePeriodSeconds *int64 `json:"TerminationGracePeriodSeconds,omitnil,omitempty" name:"TerminationGracePeriodSeconds"`

	// Command executed before the service instance stops. The instance ends after the command execution is completed or after the execution time exceeds the graceful exit time limit.
	PreStopCommand []*string `json:"PreStopCommand,omitnil,omitempty" name:"PreStopCommand"`

	// Whether to enable the gRPC port.
	GrpcEnable *bool `json:"GrpcEnable,omitnil,omitempty" name:"GrpcEnable"`

	// Health probe.
	HealthProbe *HealthProbe `json:"HealthProbe,omitnil,omitempty" name:"HealthProbe"`

	// Rolling update configurations.
	RollingUpdate *RollingUpdate `json:"RollingUpdate,omitnil,omitempty" name:"RollingUpdate"`

	// Number of instances per replica. This parameter is valid only when the deployment type is DIST or ROLE. Default value: 1.
	InstancePerReplicas *int64 `json:"InstancePerReplicas,omitnil,omitempty" name:"InstancePerReplicas"`

	// Batch data disk mounting configurations.
	VolumeMounts []*VolumeMount `json:"VolumeMounts,omitnil,omitempty" name:"VolumeMounts"`
}

type ServiceLimit struct {
	// Whether to enable throttling and speed limit at the instance level. Valid values: true and false. If the value is true, InstanceRpsLimit is required. If the value is false, InstanceRpsLimit does not take effect.
	EnableInstanceRpsLimit *bool `json:"EnableInstanceRpsLimit,omitnil,omitempty" name:"EnableInstanceRpsLimit"`

	// Speed limit for the requests per second (RPS) of each service instance. 0 indicates no throttling.
	InstanceRpsLimit *int64 `json:"InstanceRpsLimit,omitnil,omitempty" name:"InstanceRpsLimit"`

	// Whether to enable the maximum concurrency quantity limit for a single instance. Valid values: true and false. If the value is true, InstanceReqLimit is required. If the value is false, InstanceReqLimit does not take effect.
	EnableInstanceReqLimit *bool `json:"EnableInstanceReqLimit,omitnil,omitempty" name:"EnableInstanceReqLimit"`

	// Maximum concurrency for each service instance.
	InstanceReqLimit *int64 `json:"InstanceReqLimit,omitnil,omitempty" name:"InstanceReqLimit"`
}

type StartCmdInfo struct {
	// Startup command.
	StartCmd *string `json:"StartCmd,omitnil,omitempty" name:"StartCmd"`

	// Startup command for ps nodes.
	PsStartCmd *string `json:"PsStartCmd,omitnil,omitempty" name:"PsStartCmd"`

	// Startup command for Worker nodes.
	WorkerStartCmd *string `json:"WorkerStartCmd,omitnil,omitempty" name:"WorkerStartCmd"`
}

type StatefulSetCondition struct {
	// Information.Note: This field may return null, indicating that no valid values can be obtained.
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// Reason.Note: This field may return null, indicating that no valid values can be obtained.
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`

	// Status of the condition, True, False or Unknown.Note: This field may return null, indicating that no valid values can be obtained.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Type.Note: This field may return null, indicating that no valid values can be obtained.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Last update time.Note: This field may return null, indicating that no valid values can be obtained.
	LastTransitionTime *string `json:"LastTransitionTime,omitnil,omitempty" name:"LastTransitionTime"`

	// Last update time.Note: This field may return null, indicating that no valid values can be obtained.
	LastUpdateTime *string `json:"LastUpdateTime,omitnil,omitempty" name:"LastUpdateTime"`
}

type TCPSocketAction struct {
	// Called port.
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`
}

type Tag struct {
	// Tag key.Note: This field may return null, indicating that no valid values can be obtained.
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// Tag value.Note: This field may return null, indicating that no valid values can be obtained.
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

type TagFilter struct {
	// Tag key.
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// Multiple tag values.
	TagValues []*string `json:"TagValues,omitnil,omitempty" name:"TagValues"`
}

type TrainToolConfig struct {

	EnableHangMonitor *bool `json:"EnableHangMonitor,omitnil,omitempty" name:"EnableHangMonitor"`


	HangMonitorNodes []*string `json:"HangMonitorNodes,omitnil,omitempty" name:"HangMonitorNodes"`


	LogHangTimeoutInMinute *uint64 `json:"LogHangTimeoutInMinute,omitnil,omitempty" name:"LogHangTimeoutInMinute"`
}

type VolumeMount struct {
	// Cloud File Storage (CFS) configuration information.
	CFSConfig *CFSConfig `json:"CFSConfig,omitnil,omitempty" name:"CFSConfig"`

	// Mounting source type. Valid values: CFS and COS. The default value is CFS.
	VolumeSourceType *string `json:"VolumeSourceType,omitnil,omitempty" name:"VolumeSourceType"`

	// Mounting path in the custom container.Note: This field may return null, indicating that no valid values can be obtained.
	MountPath *string `json:"MountPath,omitnil,omitempty" name:"MountPath"`
}

type WorkloadStatus struct {
	// Number of current instances.
	Replicas *int64 `json:"Replicas,omitnil,omitempty" name:"Replicas"`

	// Number of updated instances.
	UpdatedReplicas *int64 `json:"UpdatedReplicas,omitnil,omitempty" name:"UpdatedReplicas"`

	// Number of ready instances.
	ReadyReplicas *int64 `json:"ReadyReplicas,omitnil,omitempty" name:"ReadyReplicas"`

	// Number of available instances.
	AvailableReplicas *int64 `json:"AvailableReplicas,omitnil,omitempty" name:"AvailableReplicas"`

	// Number of unavailable instances.
	UnavailableReplicas *int64 `json:"UnavailableReplicas,omitnil,omitempty" name:"UnavailableReplicas"`

	// Normal: running.Abnormal: service abnormalities, such as container startup failure.Waiting: service waiting, such as container image pulling.Stopped: stopped.Pending: starting.Stopping: stopping.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Workload status information.
	//
	// Deprecated: StatefulSetCondition is deprecated.
	StatefulSetCondition []*StatefulSetCondition `json:"StatefulSetCondition,omitnil,omitempty" name:"StatefulSetCondition"`

	// Status information of workload history.
	Conditions []*StatefulSetCondition `json:"Conditions,omitnil,omitempty" name:"Conditions"`

	// Display the reason when the status is abnormal.Note: This field may return null, indicating that no valid values can be obtained.
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`
}