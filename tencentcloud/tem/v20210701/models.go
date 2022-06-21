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

package v20210701

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

type CosToken struct {
	// Unique request ID
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`

	// Bucket name
	Bucket *string `json:"Bucket,omitempty" name:"Bucket"`

	// Bucket region
	Region *string `json:"Region,omitempty" name:"Region"`

	// Temporary key SecretId
	TmpSecretId *string `json:"TmpSecretId,omitempty" name:"TmpSecretId"`

	// Temporary key SecretKey
	TmpSecretKey *string `json:"TmpSecretKey,omitempty" name:"TmpSecretKey"`

	// `sessionToken` of temporary key
	SessionToken *string `json:"SessionToken,omitempty" name:"SessionToken"`

	// Start time of temporary key acquisition
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// `ExpiredTime` of temporary key
	ExpiredTime *string `json:"ExpiredTime,omitempty" name:"ExpiredTime"`

	// Full package path
	FullPath *string `json:"FullPath,omitempty" name:"FullPath"`
}

// Predefined struct for user
type CreateCosTokenRequestParams struct {
	// Application ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// Package name
	PkgName *string `json:"PkgName,omitempty" name:"PkgName"`

	// Operation type. 1: upload; 2: query
	OptType *int64 `json:"OptType,omitempty" name:"OptType"`

	// Source channel
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// Input parameter of `deployVersion`
	TimeVersion *string `json:"TimeVersion,omitempty" name:"TimeVersion"`
}

type CreateCosTokenRequest struct {
	*tchttp.BaseRequest
	
	// Application ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// Package name
	PkgName *string `json:"PkgName,omitempty" name:"PkgName"`

	// Operation type. 1: upload; 2: query
	OptType *int64 `json:"OptType,omitempty" name:"OptType"`

	// Source channel
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// Input parameter of `deployVersion`
	TimeVersion *string `json:"TimeVersion,omitempty" name:"TimeVersion"`
}

func (r *CreateCosTokenRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCosTokenRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	delete(f, "PkgName")
	delete(f, "OptType")
	delete(f, "SourceChannel")
	delete(f, "TimeVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCosTokenRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCosTokenResponseParams struct {
	// `CosToken` object in case of success and `null` in case of failure
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Result *CosToken `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateCosTokenResponse struct {
	*tchttp.BaseResponse
	Response *CreateCosTokenResponseParams `json:"Response"`
}

func (r *CreateCosTokenResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCosTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEnvironmentRequestParams struct {
	// Environment name
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// VPC name
	Vpc *string `json:"Vpc,omitempty" name:"Vpc"`

	// List of subnets
	SubnetIds []*string `json:"SubnetIds,omitempty" name:"SubnetIds"`

	// Environment description
	Description *string `json:"Description,omitempty" name:"Description"`

	// Kubernetes version
	K8sVersion *string `json:"K8sVersion,omitempty" name:"K8sVersion"`

	// Source channel
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// Whether to enable the TSW service
	EnableTswTraceService *bool `json:"EnableTswTraceService,omitempty" name:"EnableTswTraceService"`
}

type CreateEnvironmentRequest struct {
	*tchttp.BaseRequest
	
	// Environment name
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// VPC name
	Vpc *string `json:"Vpc,omitempty" name:"Vpc"`

	// List of subnets
	SubnetIds []*string `json:"SubnetIds,omitempty" name:"SubnetIds"`

	// Environment description
	Description *string `json:"Description,omitempty" name:"Description"`

	// Kubernetes version
	K8sVersion *string `json:"K8sVersion,omitempty" name:"K8sVersion"`

	// Source channel
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// Whether to enable the TSW service
	EnableTswTraceService *bool `json:"EnableTswTraceService,omitempty" name:"EnableTswTraceService"`
}

func (r *CreateEnvironmentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEnvironmentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvironmentName")
	delete(f, "Vpc")
	delete(f, "SubnetIds")
	delete(f, "Description")
	delete(f, "K8sVersion")
	delete(f, "SourceChannel")
	delete(f, "EnableTswTraceService")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateEnvironmentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEnvironmentResponseParams struct {
	// Environment ID in case of success and `null` in case of failure
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Result *string `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateEnvironmentResponse struct {
	*tchttp.BaseResponse
	Response *CreateEnvironmentResponseParams `json:"Response"`
}

func (r *CreateEnvironmentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEnvironmentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateResourceRequestParams struct {
	// Environment ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// Resource type. Valid values: `CFS` (file system), `CLS` (log service), `TSE_SRE` (registry)
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`

	// Resource ID
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// Source channel
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// Source of the resource. Values: `existing` (choose an existing resource), `creating` (create a new resource)
	ResourceFrom *string `json:"ResourceFrom,omitempty" name:"ResourceFrom"`

	// Resource extra configuration
	ResourceConfig *string `json:"ResourceConfig,omitempty" name:"ResourceConfig"`
}

type CreateResourceRequest struct {
	*tchttp.BaseRequest
	
	// Environment ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// Resource type. Valid values: `CFS` (file system), `CLS` (log service), `TSE_SRE` (registry)
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`

	// Resource ID
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// Source channel
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// Source of the resource. Values: `existing` (choose an existing resource), `creating` (create a new resource)
	ResourceFrom *string `json:"ResourceFrom,omitempty" name:"ResourceFrom"`

	// Resource extra configuration
	ResourceConfig *string `json:"ResourceConfig,omitempty" name:"ResourceConfig"`
}

func (r *CreateResourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateResourceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvironmentId")
	delete(f, "ResourceType")
	delete(f, "ResourceId")
	delete(f, "SourceChannel")
	delete(f, "ResourceFrom")
	delete(f, "ResourceConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateResourceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateResourceResponseParams struct {
	// Result
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Result *bool `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateResourceResponse struct {
	*tchttp.BaseResponse
	Response *CreateResourceResponseParams `json:"Response"`
}

func (r *CreateResourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CronHorizontalAutoscaler struct {
	// Name of a scheduled scaling policy
	Name *string `json:"Name,omitempty" name:"Name"`

	// Policy period
	// "* * *" indicates three ranges. The first is day, the second month, and the third week. The three parts are separated by spaces.
	// Examples:
	// * * * (every day)
	// * * 0-3 (every Sunday through Wednesday)
	// 1,11,21 * * (1st, 11th, and 21st of every month)
	Period *string `json:"Period,omitempty" name:"Period"`

	// Details of a scheduled scaling policy
	Schedules []*CronHorizontalAutoscalerSchedule `json:"Schedules,omitempty" name:"Schedules"`

	// Enabled or not
	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`

	// Policy priority. The higher the value, the higher the priority. The minimum value is 0.
	Priority *int64 `json:"Priority,omitempty" name:"Priority"`
}

type CronHorizontalAutoscalerSchedule struct {
	// Triggering time, in the format of HH:MM
	// Example:
	// 00:00 (Trigger at midnight)
	StartAt *string `json:"StartAt,omitempty" name:"StartAt"`

	// Number of target pods (less than 50)
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	TargetReplicas *int64 `json:"TargetReplicas,omitempty" name:"TargetReplicas"`
}

// Predefined struct for user
type DeleteIngressRequestParams struct {
	// Environment ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// Environment namespace
	ClusterNamespace *string `json:"ClusterNamespace,omitempty" name:"ClusterNamespace"`

	// Ingress rule name
	IngressName *string `json:"IngressName,omitempty" name:"IngressName"`

	// Source channel
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`
}

type DeleteIngressRequest struct {
	*tchttp.BaseRequest
	
	// Environment ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// Environment namespace
	ClusterNamespace *string `json:"ClusterNamespace,omitempty" name:"ClusterNamespace"`

	// Ingress rule name
	IngressName *string `json:"IngressName,omitempty" name:"IngressName"`

	// Source channel
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`
}

func (r *DeleteIngressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteIngressRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvironmentId")
	delete(f, "ClusterNamespace")
	delete(f, "IngressName")
	delete(f, "SourceChannel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteIngressRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteIngressResponseParams struct {
	// Whether deletion is successful
	Result *bool `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteIngressResponse struct {
	*tchttp.BaseResponse
	Response *DeleteIngressResponseParams `json:"Response"`
}

func (r *DeleteIngressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteIngressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeployApplicationRequestParams struct {
	// Application ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// Number of initialized pods
	InitPodNum *uint64 `json:"InitPodNum,omitempty" name:"InitPodNum"`

	// CPU specification
	CpuSpec *float64 `json:"CpuSpec,omitempty" name:"CpuSpec"`

	// Memory specification
	MemorySpec *float64 `json:"MemorySpec,omitempty" name:"MemorySpec"`

	// Environment ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// Image repository
	ImgRepo *string `json:"ImgRepo,omitempty" name:"ImgRepo"`

	// Version description
	VersionDesc *string `json:"VersionDesc,omitempty" name:"VersionDesc"`

	// Launch parameters
	JvmOpts *string `json:"JvmOpts,omitempty" name:"JvmOpts"`

	// Auto scaling configuration (This field is disused. Please use `HorizontalAutoscaler` to set the auto scaling policy.)
	EsInfo *EsInfo `json:"EsInfo,omitempty" name:"EsInfo"`

	// Environment variable configuration
	EnvConf []*Pair `json:"EnvConf,omitempty" name:"EnvConf"`

	// Log configuration
	LogConfs []*string `json:"LogConfs,omitempty" name:"LogConfs"`

	// Data volume configuration
	StorageConfs []*StorageConf `json:"StorageConfs,omitempty" name:"StorageConfs"`

	// Data volume mount configuration
	StorageMountConfs []*StorageMountConf `json:"StorageMountConfs,omitempty" name:"StorageMountConfs"`

	// Deployment type
	// - JAR: deployment through JAR package
	// - WAR: deployment through WAR package
	// - IMAGE: deployment through image
	DeployMode *string `json:"DeployMode,omitempty" name:"DeployMode"`

	// When the deployment type is `IMAGE`, this parameter indicates the image tag
	// When the deployment type is `JAR` or `WAR`, this parameter indicates the package version number
	DeployVersion *string `json:"DeployVersion,omitempty" name:"DeployVersion"`

	// Package name, which is required when using JAR or WAR packages for deployment
	PkgName *string `json:"PkgName,omitempty" name:"PkgName"`

	// JDK version
	// - KONA: use KONA JDK
	// - OPEN: use open JDK
	// - KONA: use KONA JDK
	// - OPEN: use open JDK
	JdkVersion *string `json:"JdkVersion,omitempty" name:"JdkVersion"`

	// Security group IDs
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`

	// Log output configuration
	LogOutputConf *LogOutputConf `json:"LogOutputConf,omitempty" name:"LogOutputConf"`

	// Source channel
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// Version description
	Description *string `json:"Description,omitempty" name:"Description"`

	// Image command
	ImageCommand *string `json:"ImageCommand,omitempty" name:"ImageCommand"`

	// Image command parameters
	ImageArgs []*string `json:"ImageArgs,omitempty" name:"ImageArgs"`

	// Whether to add the registry's default configurations
	UseRegistryDefaultConfig *bool `json:"UseRegistryDefaultConfig,omitempty" name:"UseRegistryDefaultConfig"`

	// Mounting configurations
	SettingConfs []*MountedSettingConf `json:"SettingConfs,omitempty" name:"SettingConfs"`

	// Application access configuration
	Service *EksService `json:"Service,omitempty" name:"Service"`

	// ID of the version that you want to roll back to
	VersionId *string `json:"VersionId,omitempty" name:"VersionId"`

	// The script to run after startup
	PostStart *string `json:"PostStart,omitempty" name:"PostStart"`

	// The script to run before stop
	PreStop *string `json:"PreStop,omitempty" name:"PreStop"`

	// Configuration of aliveness probe
	Liveness *HealthCheckConfig `json:"Liveness,omitempty" name:"Liveness"`

	// Configuration of readiness probe
	Readiness *HealthCheckConfig `json:"Readiness,omitempty" name:"Readiness"`

	// Configuration of batch release policies
	DeployStrategyConf *DeployStrategyConf `json:"DeployStrategyConf,omitempty" name:"DeployStrategyConf"`

	// Auto scaling policy
	HorizontalAutoscaler []*HorizontalAutoscaler `json:"HorizontalAutoscaler,omitempty" name:"HorizontalAutoscaler"`

	// Scheduled auto scaling policy
	CronHorizontalAutoscaler []*CronHorizontalAutoscaler `json:"CronHorizontalAutoscaler,omitempty" name:"CronHorizontalAutoscaler"`

	// Specifies whether to enable logging. `1`: enable; `0`: do not enable
	LogEnable *int64 `json:"LogEnable,omitempty" name:"LogEnable"`

	// Whether the configuration is modified (except for the image configuration)
	ConfEdited *bool `json:"ConfEdited,omitempty" name:"ConfEdited"`

	// Whether the application acceleration is enabled 
	SpeedUp *bool `json:"SpeedUp,omitempty" name:"SpeedUp"`

	// Whether to enable probing
	StartupProbe *HealthCheckConfig `json:"StartupProbe,omitempty" name:"StartupProbe"`

	// The version of the operating system
	// If `openjdk` is selected, the value can be: 
	// - ALPINE
	// - CENTOS
	// If `konajdk` is selected, the value can be: 
	// - ALPINE
	// - TENCENTOS
	OsFlavour *string `json:"OsFlavour,omitempty" name:"OsFlavour"`

	// Specifies whether to enable Prometheus metric
	EnablePrometheusConf *EnablePrometheusConf `json:"EnablePrometheusConf,omitempty" name:"EnablePrometheusConf"`
}

type DeployApplicationRequest struct {
	*tchttp.BaseRequest
	
	// Application ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// Number of initialized pods
	InitPodNum *uint64 `json:"InitPodNum,omitempty" name:"InitPodNum"`

	// CPU specification
	CpuSpec *float64 `json:"CpuSpec,omitempty" name:"CpuSpec"`

	// Memory specification
	MemorySpec *float64 `json:"MemorySpec,omitempty" name:"MemorySpec"`

	// Environment ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// Image repository
	ImgRepo *string `json:"ImgRepo,omitempty" name:"ImgRepo"`

	// Version description
	VersionDesc *string `json:"VersionDesc,omitempty" name:"VersionDesc"`

	// Launch parameters
	JvmOpts *string `json:"JvmOpts,omitempty" name:"JvmOpts"`

	// Auto scaling configuration (This field is disused. Please use `HorizontalAutoscaler` to set the auto scaling policy.)
	EsInfo *EsInfo `json:"EsInfo,omitempty" name:"EsInfo"`

	// Environment variable configuration
	EnvConf []*Pair `json:"EnvConf,omitempty" name:"EnvConf"`

	// Log configuration
	LogConfs []*string `json:"LogConfs,omitempty" name:"LogConfs"`

	// Data volume configuration
	StorageConfs []*StorageConf `json:"StorageConfs,omitempty" name:"StorageConfs"`

	// Data volume mount configuration
	StorageMountConfs []*StorageMountConf `json:"StorageMountConfs,omitempty" name:"StorageMountConfs"`

	// Deployment type
	// - JAR: deployment through JAR package
	// - WAR: deployment through WAR package
	// - IMAGE: deployment through image
	DeployMode *string `json:"DeployMode,omitempty" name:"DeployMode"`

	// When the deployment type is `IMAGE`, this parameter indicates the image tag
	// When the deployment type is `JAR` or `WAR`, this parameter indicates the package version number
	DeployVersion *string `json:"DeployVersion,omitempty" name:"DeployVersion"`

	// Package name, which is required when using JAR or WAR packages for deployment
	PkgName *string `json:"PkgName,omitempty" name:"PkgName"`

	// JDK version
	// - KONA: use KONA JDK
	// - OPEN: use open JDK
	// - KONA: use KONA JDK
	// - OPEN: use open JDK
	JdkVersion *string `json:"JdkVersion,omitempty" name:"JdkVersion"`

	// Security group IDs
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`

	// Log output configuration
	LogOutputConf *LogOutputConf `json:"LogOutputConf,omitempty" name:"LogOutputConf"`

	// Source channel
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// Version description
	Description *string `json:"Description,omitempty" name:"Description"`

	// Image command
	ImageCommand *string `json:"ImageCommand,omitempty" name:"ImageCommand"`

	// Image command parameters
	ImageArgs []*string `json:"ImageArgs,omitempty" name:"ImageArgs"`

	// Whether to add the registry's default configurations
	UseRegistryDefaultConfig *bool `json:"UseRegistryDefaultConfig,omitempty" name:"UseRegistryDefaultConfig"`

	// Mounting configurations
	SettingConfs []*MountedSettingConf `json:"SettingConfs,omitempty" name:"SettingConfs"`

	// Application access configuration
	Service *EksService `json:"Service,omitempty" name:"Service"`

	// ID of the version that you want to roll back to
	VersionId *string `json:"VersionId,omitempty" name:"VersionId"`

	// The script to run after startup
	PostStart *string `json:"PostStart,omitempty" name:"PostStart"`

	// The script to run before stop
	PreStop *string `json:"PreStop,omitempty" name:"PreStop"`

	// Configuration of aliveness probe
	Liveness *HealthCheckConfig `json:"Liveness,omitempty" name:"Liveness"`

	// Configuration of readiness probe
	Readiness *HealthCheckConfig `json:"Readiness,omitempty" name:"Readiness"`

	// Configuration of batch release policies
	DeployStrategyConf *DeployStrategyConf `json:"DeployStrategyConf,omitempty" name:"DeployStrategyConf"`

	// Auto scaling policy
	HorizontalAutoscaler []*HorizontalAutoscaler `json:"HorizontalAutoscaler,omitempty" name:"HorizontalAutoscaler"`

	// Scheduled auto scaling policy
	CronHorizontalAutoscaler []*CronHorizontalAutoscaler `json:"CronHorizontalAutoscaler,omitempty" name:"CronHorizontalAutoscaler"`

	// Specifies whether to enable logging. `1`: enable; `0`: do not enable
	LogEnable *int64 `json:"LogEnable,omitempty" name:"LogEnable"`

	// Whether the configuration is modified (except for the image configuration)
	ConfEdited *bool `json:"ConfEdited,omitempty" name:"ConfEdited"`

	// Whether the application acceleration is enabled 
	SpeedUp *bool `json:"SpeedUp,omitempty" name:"SpeedUp"`

	// Whether to enable probing
	StartupProbe *HealthCheckConfig `json:"StartupProbe,omitempty" name:"StartupProbe"`

	// The version of the operating system
	// If `openjdk` is selected, the value can be: 
	// - ALPINE
	// - CENTOS
	// If `konajdk` is selected, the value can be: 
	// - ALPINE
	// - TENCENTOS
	OsFlavour *string `json:"OsFlavour,omitempty" name:"OsFlavour"`

	// Specifies whether to enable Prometheus metric
	EnablePrometheusConf *EnablePrometheusConf `json:"EnablePrometheusConf,omitempty" name:"EnablePrometheusConf"`
}

func (r *DeployApplicationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeployApplicationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	delete(f, "InitPodNum")
	delete(f, "CpuSpec")
	delete(f, "MemorySpec")
	delete(f, "EnvironmentId")
	delete(f, "ImgRepo")
	delete(f, "VersionDesc")
	delete(f, "JvmOpts")
	delete(f, "EsInfo")
	delete(f, "EnvConf")
	delete(f, "LogConfs")
	delete(f, "StorageConfs")
	delete(f, "StorageMountConfs")
	delete(f, "DeployMode")
	delete(f, "DeployVersion")
	delete(f, "PkgName")
	delete(f, "JdkVersion")
	delete(f, "SecurityGroupIds")
	delete(f, "LogOutputConf")
	delete(f, "SourceChannel")
	delete(f, "Description")
	delete(f, "ImageCommand")
	delete(f, "ImageArgs")
	delete(f, "UseRegistryDefaultConfig")
	delete(f, "SettingConfs")
	delete(f, "Service")
	delete(f, "VersionId")
	delete(f, "PostStart")
	delete(f, "PreStop")
	delete(f, "Liveness")
	delete(f, "Readiness")
	delete(f, "DeployStrategyConf")
	delete(f, "HorizontalAutoscaler")
	delete(f, "CronHorizontalAutoscaler")
	delete(f, "LogEnable")
	delete(f, "ConfEdited")
	delete(f, "SpeedUp")
	delete(f, "StartupProbe")
	delete(f, "OsFlavour")
	delete(f, "EnablePrometheusConf")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeployApplicationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeployApplicationResponseParams struct {
	// Version ID (which can be ignored for the frontend)
	Result *string `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeployApplicationResponse struct {
	*tchttp.BaseResponse
	Response *DeployApplicationResponseParams `json:"Response"`
}

func (r *DeployApplicationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeployApplicationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeployStrategyConf struct {
	// Total batches
	TotalBatchCount *int64 `json:"TotalBatchCount,omitempty" name:"TotalBatchCount"`

	// Number of pods for the beta batch
	BetaBatchNum *int64 `json:"BetaBatchNum,omitempty" name:"BetaBatchNum"`

	// Batch deployment policy. `0`: automatically; `1`: manually; `2`: beta batch (manual), `3`: initial release
	DeployStrategyType *int64 `json:"DeployStrategyType,omitempty" name:"DeployStrategyType"`

	// Interval between batches
	BatchInterval *int64 `json:"BatchInterval,omitempty" name:"BatchInterval"`

	// The minimum number of available pods
	MinAvailable *int64 `json:"MinAvailable,omitempty" name:"MinAvailable"`

	// Whether to enable force release
	Force *bool `json:"Force,omitempty" name:"Force"`
}

// Predefined struct for user
type DescribeApplicationPodsRequestParams struct {
	// Environment ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// Application ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// Number of items per page. Default value: 20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Page number. Default value: 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Pod status 
	// - Running 
	// - Pending 
	// - Error
	Status *string `json:"Status,omitempty" name:"Status"`

	// Pod name
	PodName *string `json:"PodName,omitempty" name:"PodName"`

	// Source channel
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`
}

type DescribeApplicationPodsRequest struct {
	*tchttp.BaseRequest
	
	// Environment ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// Application ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// Number of items per page. Default value: 20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Page number. Default value: 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Pod status 
	// - Running 
	// - Pending 
	// - Error
	Status *string `json:"Status,omitempty" name:"Status"`

	// Pod name
	PodName *string `json:"PodName,omitempty" name:"PodName"`

	// Source channel
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`
}

func (r *DescribeApplicationPodsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationPodsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvironmentId")
	delete(f, "ApplicationId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Status")
	delete(f, "PodName")
	delete(f, "SourceChannel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApplicationPodsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApplicationPodsResponseParams struct {
	// Returned result
	Result *DescribeRunPodPage `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeApplicationPodsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApplicationPodsResponseParams `json:"Response"`
}

func (r *DescribeApplicationPodsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationPodsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEnvironmentsRequestParams struct {
	// Pagination limit
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Page offset
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Source
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`
}

type DescribeEnvironmentsRequest struct {
	*tchttp.BaseRequest
	
	// Pagination limit
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Page offset
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Source
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`
}

func (r *DescribeEnvironmentsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEnvironmentsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "SourceChannel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEnvironmentsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEnvironmentsResponseParams struct {
	// Returned result
	Result *NamespacePage `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeEnvironmentsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEnvironmentsResponseParams `json:"Response"`
}

func (r *DescribeEnvironmentsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEnvironmentsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIngressRequestParams struct {
	// Environment ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// Environment namespace
	ClusterNamespace *string `json:"ClusterNamespace,omitempty" name:"ClusterNamespace"`

	// Ingress rule name
	IngressName *string `json:"IngressName,omitempty" name:"IngressName"`

	// Source channel
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`
}

type DescribeIngressRequest struct {
	*tchttp.BaseRequest
	
	// Environment ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// Environment namespace
	ClusterNamespace *string `json:"ClusterNamespace,omitempty" name:"ClusterNamespace"`

	// Ingress rule name
	IngressName *string `json:"IngressName,omitempty" name:"IngressName"`

	// Source channel
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`
}

func (r *DescribeIngressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIngressRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvironmentId")
	delete(f, "ClusterNamespace")
	delete(f, "IngressName")
	delete(f, "SourceChannel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIngressRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIngressResponseParams struct {
	// Ingress rule configuration
	Result *IngressInfo `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeIngressResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIngressResponseParams `json:"Response"`
}

func (r *DescribeIngressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIngressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIngressesRequestParams struct {
	// Environment ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// Environment namespace
	ClusterNamespace *string `json:"ClusterNamespace,omitempty" name:"ClusterNamespace"`

	// Source channel
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// Ingress rule name list
	IngressNames []*string `json:"IngressNames,omitempty" name:"IngressNames"`
}

type DescribeIngressesRequest struct {
	*tchttp.BaseRequest
	
	// Environment ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// Environment namespace
	ClusterNamespace *string `json:"ClusterNamespace,omitempty" name:"ClusterNamespace"`

	// Source channel
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// Ingress rule name list
	IngressNames []*string `json:"IngressNames,omitempty" name:"IngressNames"`
}

func (r *DescribeIngressesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIngressesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvironmentId")
	delete(f, "ClusterNamespace")
	delete(f, "SourceChannel")
	delete(f, "IngressNames")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIngressesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIngressesResponseParams struct {
	// Ingress array
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Result []*IngressInfo `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeIngressesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIngressesResponseParams `json:"Response"`
}

func (r *DescribeIngressesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIngressesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRelatedIngressesRequestParams struct {
	// Environment ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// Environment namespace
	ClusterNamespace *string `json:"ClusterNamespace,omitempty" name:"ClusterNamespace"`

	// Source channel
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// Application ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`
}

type DescribeRelatedIngressesRequest struct {
	*tchttp.BaseRequest
	
	// Environment ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// Environment namespace
	ClusterNamespace *string `json:"ClusterNamespace,omitempty" name:"ClusterNamespace"`

	// Source channel
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// Application ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`
}

func (r *DescribeRelatedIngressesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRelatedIngressesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvironmentId")
	delete(f, "ClusterNamespace")
	delete(f, "SourceChannel")
	delete(f, "ApplicationId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRelatedIngressesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRelatedIngressesResponseParams struct {
	// Ingress array
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Result []*IngressInfo `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeRelatedIngressesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRelatedIngressesResponseParams `json:"Response"`
}

func (r *DescribeRelatedIngressesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRelatedIngressesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRunPodPage struct {
	// Page offset
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of records per page
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Total number of returned records
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Request ID
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`

	// List of pods
	PodList []*RunVersionPod `json:"PodList,omitempty" name:"PodList"`
}

type EksService struct {
	// Service name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Available ports
	Ports []*int64 `json:"Ports,omitempty" name:"Ports"`

	// Yaml contents
	Yaml *string `json:"Yaml,omitempty" name:"Yaml"`

	// Service name
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`

	// Version name
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	VersionName *string `json:"VersionName,omitempty" name:"VersionName"`

	// Private IP
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ClusterIp []*string `json:"ClusterIp,omitempty" name:"ClusterIp"`

	// Public IP
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ExternalIp *string `json:"ExternalIp,omitempty" name:"ExternalIp"`

	// The access type. Valid values:
	// - EXTERNAL (internet access)
	// - VPC (Intra-VPC access)
	// - CLUSTER (Intra-cluster access)
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Type *string `json:"Type,omitempty" name:"Type"`

	// Subnet ID. It is filled when the access type is `VPC`.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// Load balancer ID. It is filled when the access type is `EXTERNAL` or `CLUSTER`. It’s created automatically by default.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	LoadBalanceId *string `json:"LoadBalanceId,omitempty" name:"LoadBalanceId"`

	// Port mapping
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	PortMappings []*PortMapping `json:"PortMappings,omitempty" name:"PortMappings"`
}

type EnablePrometheusConf struct {
	// The listening port of the applicaiton
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// URL path for monitoring
	Path *string `json:"Path,omitempty" name:"Path"`
}

type EsInfo struct {
	// Minimum number of instances
	MinAliveInstances *int64 `json:"MinAliveInstances,omitempty" name:"MinAliveInstances"`

	// Maximum number of instances
	MaxAliveInstances *int64 `json:"MaxAliveInstances,omitempty" name:"MaxAliveInstances"`

	// Auto scaling policy. 1: CPU; 2: memory
	EsStrategy *int64 `json:"EsStrategy,omitempty" name:"EsStrategy"`

	// Auto scaling condition value
	Threshold *uint64 `json:"Threshold,omitempty" name:"Threshold"`

	// Version ID
	VersionId *string `json:"VersionId,omitempty" name:"VersionId"`
}

// Predefined struct for user
type GenerateApplicationPackageDownloadUrlRequestParams struct {
	// Application ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// Package name
	PkgName *string `json:"PkgName,omitempty" name:"PkgName"`

	// Version of the package to download
	DeployVersion *string `json:"DeployVersion,omitempty" name:"DeployVersion"`

	// Source channel
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`
}

type GenerateApplicationPackageDownloadUrlRequest struct {
	*tchttp.BaseRequest
	
	// Application ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// Package name
	PkgName *string `json:"PkgName,omitempty" name:"PkgName"`

	// Version of the package to download
	DeployVersion *string `json:"DeployVersion,omitempty" name:"DeployVersion"`

	// Source channel
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`
}

func (r *GenerateApplicationPackageDownloadUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GenerateApplicationPackageDownloadUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	delete(f, "PkgName")
	delete(f, "DeployVersion")
	delete(f, "SourceChannel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GenerateApplicationPackageDownloadUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GenerateApplicationPackageDownloadUrlResponseParams struct {
	// Temp download URL for the package
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Result *string `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GenerateApplicationPackageDownloadUrlResponse struct {
	*tchttp.BaseResponse
	Response *GenerateApplicationPackageDownloadUrlResponseParams `json:"Response"`
}

func (r *GenerateApplicationPackageDownloadUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GenerateApplicationPackageDownloadUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type HealthCheckConfig struct {
	// Health check type. Valid values: `HttpGet`, `TcpSocket`, `Exec`
	Type *string `json:"Type,omitempty" name:"Type"`

	// The protocol type. It’s only valid when the health check type is `HttpGet`.
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// The request path. It’s only valid when the health check type is `HttpGet`.
	Path *string `json:"Path,omitempty" name:"Path"`

	// The script to be executed. It’s only valid when the health check type is `Exec`.
	Exec *string `json:"Exec,omitempty" name:"Exec"`

	// The request port. It’s only valid when the health check type is `HttpGet` or `TcpSocket `.
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// The initial delay for health check in seconds. Default: `0`
	InitialDelaySeconds *int64 `json:"InitialDelaySeconds,omitempty" name:"InitialDelaySeconds"`

	// Timeout period in seconds. Default: `1`
	TimeoutSeconds *int64 `json:"TimeoutSeconds,omitempty" name:"TimeoutSeconds"`

	// Interval period in seconds. Default: `10`
	PeriodSeconds *int64 `json:"PeriodSeconds,omitempty" name:"PeriodSeconds"`
}

type HorizontalAutoscaler struct {
	// Minimum number of instances
	MinReplicas *int64 `json:"MinReplicas,omitempty" name:"MinReplicas"`

	// Maximum number of instances
	MaxReplicas *int64 `json:"MaxReplicas,omitempty" name:"MaxReplicas"`

	// Metrics (CPU or memory)
	Metrics *string `json:"Metrics,omitempty" name:"Metrics"`

	// Threshold (percentage)
	Threshold *int64 `json:"Threshold,omitempty" name:"Threshold"`
}

type IngressInfo struct {
	// Environment ID
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// Environment namespace
	ClusterNamespace *string `json:"ClusterNamespace,omitempty" name:"ClusterNamespace"`

	// ip version
	AddressIPVersion *string `json:"AddressIPVersion,omitempty" name:"AddressIPVersion"`

	// ingress name
	IngressName *string `json:"IngressName,omitempty" name:"IngressName"`

	// Rules configuration
	Rules []*IngressRule `json:"Rules,omitempty" name:"Rules"`

	// clb ID
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ClbId *string `json:"ClbId,omitempty" name:"ClbId"`

	// TLS configuration
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Tls []*IngressTls `json:"Tls,omitempty" name:"Tls"`

	// Environment cluster ID
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// clb ip
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// Creation time
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Whether to listen on both the HTTP 80 port and HTTPS 443 port. The default value is `false`. The optional value `true` means listening on both the HTTP 80 port and HTTPS 443 port.
	Mixed *bool `json:"Mixed,omitempty" name:"Mixed"`

	// Redirection mode. Values:
	// - `AUTO` (automatically redirect HTTP to HTTPS)
	// - `NONE` (no redirection)
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	RewriteType *string `json:"RewriteType,omitempty" name:"RewriteType"`
}

type IngressRule struct {
	// ingress rule value
	Http *IngressRuleValue `json:"Http,omitempty" name:"Http"`

	// Host address
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Host *string `json:"Host,omitempty" name:"Host"`

	// Protocol. Options include HTTP and HTTPS. The default option is HTTP.
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
}

type IngressRuleBackend struct {
	// EKS service name
	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`

	// EKS service port
	ServicePort *int64 `json:"ServicePort,omitempty" name:"ServicePort"`
}

type IngressRulePath struct {
	// Path information
	Path *string `json:"Path,omitempty" name:"Path"`

	// Backend configuration
	Backend *IngressRuleBackend `json:"Backend,omitempty" name:"Backend"`
}

type IngressRuleValue struct {
	// Overall rule configuration
	Paths []*IngressRulePath `json:"Paths,omitempty" name:"Paths"`
}

type IngressTls struct {
	// Host array. An empty array indicates the default certificate for all domain names.
	Hosts []*string `json:"Hosts,omitempty" name:"Hosts"`

	// Secret name. If a certificate is used, this field is left empty.
	SecretName *string `json:"SecretName,omitempty" name:"SecretName"`

	// SSL Certificate Id
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`
}

type LogOutputConf struct {
	// Log consumer type
	OutputType *string `json:"OutputType,omitempty" name:"OutputType"`

	// CLS logset
	ClsLogsetName *string `json:"ClsLogsetName,omitempty" name:"ClsLogsetName"`

	// CLS log topic
	ClsLogTopicId *string `json:"ClsLogTopicId,omitempty" name:"ClsLogTopicId"`

	// CLS logset ID
	ClsLogsetId *string `json:"ClsLogsetId,omitempty" name:"ClsLogsetId"`

	// CLS log topic name
	ClsLogTopicName *string `json:"ClsLogTopicName,omitempty" name:"ClsLogTopicName"`
}

// Predefined struct for user
type ModifyApplicationInfoRequestParams struct {
	// Application ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// Description
	Description *string `json:"Description,omitempty" name:"Description"`

	// Source channel
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// Whether to enable the call chain. Valid values: `0`: disable; `1`: enable
	EnableTracing *uint64 `json:"EnableTracing,omitempty" name:"EnableTracing"`
}

type ModifyApplicationInfoRequest struct {
	*tchttp.BaseRequest
	
	// Application ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// Description
	Description *string `json:"Description,omitempty" name:"Description"`

	// Source channel
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// Whether to enable the call chain. Valid values: `0`: disable; `1`: enable
	EnableTracing *uint64 `json:"EnableTracing,omitempty" name:"EnableTracing"`
}

func (r *ModifyApplicationInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApplicationInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	delete(f, "Description")
	delete(f, "SourceChannel")
	delete(f, "EnableTracing")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyApplicationInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyApplicationInfoResponseParams struct {
	// Success or not
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Result *bool `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyApplicationInfoResponse struct {
	*tchttp.BaseResponse
	Response *ModifyApplicationInfoResponseParams `json:"Response"`
}

func (r *ModifyApplicationInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApplicationInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyEnvironmentRequestParams struct {
	// Environment ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// Environment name
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// Environment description
	Description *string `json:"Description,omitempty" name:"Description"`

	// VPC name
	Vpc *string `json:"Vpc,omitempty" name:"Vpc"`

	// Subnets
	SubnetIds []*string `json:"SubnetIds,omitempty" name:"SubnetIds"`

	// Source channel
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`
}

type ModifyEnvironmentRequest struct {
	*tchttp.BaseRequest
	
	// Environment ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// Environment name
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// Environment description
	Description *string `json:"Description,omitempty" name:"Description"`

	// VPC name
	Vpc *string `json:"Vpc,omitempty" name:"Vpc"`

	// Subnets
	SubnetIds []*string `json:"SubnetIds,omitempty" name:"SubnetIds"`

	// Source channel
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`
}

func (r *ModifyEnvironmentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyEnvironmentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvironmentId")
	delete(f, "EnvironmentName")
	delete(f, "Description")
	delete(f, "Vpc")
	delete(f, "SubnetIds")
	delete(f, "SourceChannel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyEnvironmentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyEnvironmentResponseParams struct {
	// Environment ID in case of success and `null` in case of failure
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Result *bool `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyEnvironmentResponse struct {
	*tchttp.BaseResponse
	Response *ModifyEnvironmentResponseParams `json:"Response"`
}

func (r *ModifyEnvironmentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyEnvironmentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyIngressRequestParams struct {
	// Ingress rule configuration
	Ingress *IngressInfo `json:"Ingress,omitempty" name:"Ingress"`

	// Source channel
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`
}

type ModifyIngressRequest struct {
	*tchttp.BaseRequest
	
	// Ingress rule configuration
	Ingress *IngressInfo `json:"Ingress,omitempty" name:"Ingress"`

	// Source channel
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`
}

func (r *ModifyIngressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyIngressRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ingress")
	delete(f, "SourceChannel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyIngressRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyIngressResponseParams struct {
	// Created successfully.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Result *bool `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyIngressResponse struct {
	*tchttp.BaseResponse
	Response *ModifyIngressResponseParams `json:"Response"`
}

func (r *ModifyIngressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyIngressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MountedSettingConf struct {
	// Configuration name
	ConfigDataName *string `json:"ConfigDataName,omitempty" name:"ConfigDataName"`

	// Mount point path
	MountedPath *string `json:"MountedPath,omitempty" name:"MountedPath"`

	// Configuration content
	Data []*Pair `json:"Data,omitempty" name:"Data"`

	// Encrypt configuration name
	SecretDataName *string `json:"SecretDataName,omitempty" name:"SecretDataName"`
}

type NamespacePage struct {
	// Details of the returned records
	Records []*TemNamespaceInfo `json:"Records,omitempty" name:"Records"`

	// Total number of returned records
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// Number of records per page
	Size *int64 `json:"Size,omitempty" name:"Size"`

	// Total number of pages
	Pages *int64 `json:"Pages,omitempty" name:"Pages"`
}

type Pair struct {
	// Key
	Key *string `json:"Key,omitempty" name:"Key"`

	// Value
	Value *string `json:"Value,omitempty" name:"Value"`

	// `default``: Custom. `reserved`: System variable. `referenced`: Referenced configuration item.
	// Note: This field may return `null`, indicating that no valid value can be found.
	Type *string `json:"Type,omitempty" name:"Type"`

	// Configuration name
	// Note: This field may return `null`, indicating that no valid value can be found.
	Config *string `json:"Config,omitempty" name:"Config"`

	// Encrypt configuration name
	// Note: This field may return `null`, indicating that no valid value was found.
	Secret *string `json:"Secret,omitempty" name:"Secret"`
}

type PortMapping struct {
	// Port.
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// Mapped port
	TargetPort *int64 `json:"TargetPort,omitempty" name:"TargetPort"`

	// TCP/UDP protocol stack.
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
}

// Predefined struct for user
type RestartApplicationPodRequestParams struct {
	// Environment ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// Application ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// Name
	PodName *string `json:"PodName,omitempty" name:"PodName"`

	// Number of items per page
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Page offset
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Pod status
	Status *string `json:"Status,omitempty" name:"Status"`

	// Source channel
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`
}

type RestartApplicationPodRequest struct {
	*tchttp.BaseRequest
	
	// Environment ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// Application ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// Name
	PodName *string `json:"PodName,omitempty" name:"PodName"`

	// Number of items per page
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Page offset
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Pod status
	Status *string `json:"Status,omitempty" name:"Status"`

	// Source channel
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`
}

func (r *RestartApplicationPodRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestartApplicationPodRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvironmentId")
	delete(f, "ApplicationId")
	delete(f, "PodName")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Status")
	delete(f, "SourceChannel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RestartApplicationPodRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RestartApplicationPodResponseParams struct {
	// Returned result
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Result *bool `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RestartApplicationPodResponse struct {
	*tchttp.BaseResponse
	Response *RestartApplicationPodResponseParams `json:"Response"`
}

func (r *RestartApplicationPodResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestartApplicationPodResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RestartApplicationRequestParams struct {
	// Application ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// Retain as default
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// Environment ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`
}

type RestartApplicationRequest struct {
	*tchttp.BaseRequest
	
	// Application ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// Retain as default
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// Environment ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`
}

func (r *RestartApplicationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestartApplicationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	delete(f, "SourceChannel")
	delete(f, "EnvironmentId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RestartApplicationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RestartApplicationResponseParams struct {
	// Returned result
	Result *bool `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RestartApplicationResponse struct {
	*tchttp.BaseResponse
	Response *RestartApplicationResponseParams `json:"Response"`
}

func (r *RestartApplicationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestartApplicationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RollingUpdateApplicationByVersionRequestParams struct {
	// Application ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// Environment ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// Update version. For image-based deployment, it is the value. For deployment with JAR/WAR files, it is `Version`.
	DeployVersion *string `json:"DeployVersion,omitempty" name:"DeployVersion"`

	// JAR/WAR package name. It’s only required for deployment with JAR/WAR files.
	PackageName *string `json:"PackageName,omitempty" name:"PackageName"`

	// Request source. Options: `IntelliJ`, `Coding`
	From *string `json:"From,omitempty" name:"From"`

	// The deployment policy. Values: `AUTO` (automatically deploy), `BETA` (deploy a small batch first to test the result, and deploy the rest automatically) and `MANUAL` (manually deploy)
	DeployStrategyType *string `json:"DeployStrategyType,omitempty" name:"DeployStrategyType"`

	// Total number of batches
	TotalBatchCount *int64 `json:"TotalBatchCount,omitempty" name:"TotalBatchCount"`

	// Interval between the batches
	BatchInterval *int64 `json:"BatchInterval,omitempty" name:"BatchInterval"`

	// Number of instances in a beta batch
	BetaBatchNum *int64 `json:"BetaBatchNum,omitempty" name:"BetaBatchNum"`

	// Minimum number of available instances during the deployment
	MinAvailable *int64 `json:"MinAvailable,omitempty" name:"MinAvailable"`
}

type RollingUpdateApplicationByVersionRequest struct {
	*tchttp.BaseRequest
	
	// Application ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// Environment ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// Update version. For image-based deployment, it is the value. For deployment with JAR/WAR files, it is `Version`.
	DeployVersion *string `json:"DeployVersion,omitempty" name:"DeployVersion"`

	// JAR/WAR package name. It’s only required for deployment with JAR/WAR files.
	PackageName *string `json:"PackageName,omitempty" name:"PackageName"`

	// Request source. Options: `IntelliJ`, `Coding`
	From *string `json:"From,omitempty" name:"From"`

	// The deployment policy. Values: `AUTO` (automatically deploy), `BETA` (deploy a small batch first to test the result, and deploy the rest automatically) and `MANUAL` (manually deploy)
	DeployStrategyType *string `json:"DeployStrategyType,omitempty" name:"DeployStrategyType"`

	// Total number of batches
	TotalBatchCount *int64 `json:"TotalBatchCount,omitempty" name:"TotalBatchCount"`

	// Interval between the batches
	BatchInterval *int64 `json:"BatchInterval,omitempty" name:"BatchInterval"`

	// Number of instances in a beta batch
	BetaBatchNum *int64 `json:"BetaBatchNum,omitempty" name:"BetaBatchNum"`

	// Minimum number of available instances during the deployment
	MinAvailable *int64 `json:"MinAvailable,omitempty" name:"MinAvailable"`
}

func (r *RollingUpdateApplicationByVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RollingUpdateApplicationByVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	delete(f, "EnvironmentId")
	delete(f, "DeployVersion")
	delete(f, "PackageName")
	delete(f, "From")
	delete(f, "DeployStrategyType")
	delete(f, "TotalBatchCount")
	delete(f, "BatchInterval")
	delete(f, "BetaBatchNum")
	delete(f, "MinAvailable")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RollingUpdateApplicationByVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RollingUpdateApplicationByVersionResponseParams struct {
	// Version ID
	Result *string `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RollingUpdateApplicationByVersionResponse struct {
	*tchttp.BaseResponse
	Response *RollingUpdateApplicationByVersionResponseParams `json:"Response"`
}

func (r *RollingUpdateApplicationByVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RollingUpdateApplicationByVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RunVersionPod struct {
	// Shell address
	Webshell *string `json:"Webshell,omitempty" name:"Webshell"`

	// Pod ID
	PodId *string `json:"PodId,omitempty" name:"PodId"`

	// Status
	Status *string `json:"Status,omitempty" name:"Status"`

	// Creation time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Pod IP
	PodIp *string `json:"PodIp,omitempty" name:"PodIp"`

	// Availability zone
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// Deployed version
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	DeployVersion *string `json:"DeployVersion,omitempty" name:"DeployVersion"`

	// Number of restarts
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	RestartCount *int64 `json:"RestartCount,omitempty" name:"RestartCount"`

	// Whether the pod is ready
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Ready *bool `json:"Ready,omitempty" name:"Ready"`

	// Container status
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ContainerState *string `json:"ContainerState,omitempty" name:"ContainerState"`
}

// Predefined struct for user
type StopApplicationRequestParams struct {
	// Application ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// Retain as default
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// Environment ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`
}

type StopApplicationRequest struct {
	*tchttp.BaseRequest
	
	// Application ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// Retain as default
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// Environment ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`
}

func (r *StopApplicationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopApplicationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	delete(f, "SourceChannel")
	delete(f, "EnvironmentId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopApplicationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopApplicationResponseParams struct {
	// Returned result
	Result *bool `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type StopApplicationResponse struct {
	*tchttp.BaseResponse
	Response *StopApplicationResponseParams `json:"Response"`
}

func (r *StopApplicationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopApplicationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StorageConf struct {
	// Storage volume name
	StorageVolName *string `json:"StorageVolName,omitempty" name:"StorageVolName"`

	// Storage volume path
	StorageVolPath *string `json:"StorageVolPath,omitempty" name:"StorageVolPath"`

	// Storage volume IP
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	StorageVolIp *string `json:"StorageVolIp,omitempty" name:"StorageVolIp"`
}

type StorageMountConf struct {
	// Data volume name
	VolumeName *string `json:"VolumeName,omitempty" name:"VolumeName"`

	// Data volume binding path
	MountPath *string `json:"MountPath,omitempty" name:"MountPath"`
}

type TemNamespaceInfo struct {
	// Environment ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// Channel
	Channel *string `json:"Channel,omitempty" name:"Channel"`

	// Environment name
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// Region name
	Region *string `json:"Region,omitempty" name:"Region"`

	// Environment description
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Description *string `json:"Description,omitempty" name:"Description"`

	// Status. `1`: terminated; `0`: normal
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// VPC network
	Vpc *string `json:"Vpc,omitempty" name:"Vpc"`

	// Creation time
	CreateDate *string `json:"CreateDate,omitempty" name:"CreateDate"`

	// Modification time
	ModifyDate *string `json:"ModifyDate,omitempty" name:"ModifyDate"`

	// Modifier
	Modifier *string `json:"Modifier,omitempty" name:"Modifier"`

	// Creator
	Creator *string `json:"Creator,omitempty" name:"Creator"`

	// Number of applications
	ApplicationNum *int64 `json:"ApplicationNum,omitempty" name:"ApplicationNum"`

	// Number of running instances
	RunInstancesNum *int64 `json:"RunInstancesNum,omitempty" name:"RunInstancesNum"`

	// Subnet
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// Environment cluster status
	ClusterStatus *string `json:"ClusterStatus,omitempty" name:"ClusterStatus"`

	// Whether to enable TSW
	EnableTswTraceService *bool `json:"EnableTswTraceService,omitempty" name:"EnableTswTraceService"`

	// Whether the environment is locked. `1`: locked; `0`: not locked
	Locked *int64 `json:"Locked,omitempty" name:"Locked"`
}