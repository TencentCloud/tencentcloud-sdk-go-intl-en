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

type Autoscaler struct {
	// Minimum number of instances in a scaling group
	MinReplicas *int64 `json:"MinReplicas,omitempty" name:"MinReplicas"`

	// Maximum number of instances in a scaling group
	MaxReplicas *int64 `json:"MaxReplicas,omitempty" name:"MaxReplicas"`

	// Policy of the scaling rule
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	HorizontalAutoscaler []*HorizontalAutoscaler `json:"HorizontalAutoscaler,omitempty" name:"HorizontalAutoscaler"`

	// Scheduled auto-scaler policy
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	CronHorizontalAutoscaler []*CronHorizontalAutoscaler `json:"CronHorizontalAutoscaler,omitempty" name:"CronHorizontalAutoscaler"`

	// Scaling rule ID
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	AutoscalerId *string `json:"AutoscalerId,omitempty" name:"AutoscalerId"`

	// Scaling rule name
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	AutoscalerName *string `json:"AutoscalerName,omitempty" name:"AutoscalerName"`

	// Description of the scaling rule
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Description *string `json:"Description,omitempty" name:"Description"`

	// Creation time
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	CreateDate *string `json:"CreateDate,omitempty" name:"CreateDate"`

	// Modification time
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	ModifyDate *string `json:"ModifyDate,omitempty" name:"ModifyDate"`

	// Start Time
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	EnableDate *string `json:"EnableDate,omitempty" name:"EnableDate"`

	// Whether it is enabled
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`
}

type ConfigData struct {
	// Configuration name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Creation time.
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// List of associated applications
	RelatedApplications []*TemService `json:"RelatedApplications,omitempty" name:"RelatedApplications"`

	// Configuration item
	Data []*Pair `json:"Data,omitempty" name:"Data"`
}

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
type CreateApplicationAutoscalerRequestParams struct {
	// Application ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// Environment ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// Source channel. Please keep the default value.
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// Auto scaling rule
	Autoscaler *Autoscaler `json:"Autoscaler,omitempty" name:"Autoscaler"`
}

type CreateApplicationAutoscalerRequest struct {
	*tchttp.BaseRequest
	
	// Application ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// Environment ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// Source channel. Please keep the default value.
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// Auto scaling rule
	Autoscaler *Autoscaler `json:"Autoscaler,omitempty" name:"Autoscaler"`
}

func (r *CreateApplicationAutoscalerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateApplicationAutoscalerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	delete(f, "EnvironmentId")
	delete(f, "SourceChannel")
	delete(f, "Autoscaler")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateApplicationAutoscalerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateApplicationAutoscalerResponseParams struct {
	// Scaling rule ID
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Result *string `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateApplicationAutoscalerResponse struct {
	*tchttp.BaseResponse
	Response *CreateApplicationAutoscalerResponseParams `json:"Response"`
}

func (r *CreateApplicationAutoscalerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateApplicationAutoscalerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateApplicationRequestParams struct {
	// Application name
	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`

	// Description
	Description *string `json:"Description,omitempty" name:"Description"`

	// Whether to use the default image service. `1`: yes; `0`: no
	UseDefaultImageService *int64 `json:"UseDefaultImageService,omitempty" name:"UseDefaultImageService"`

	// Type of the bound repository. `0`: TCR Personal; `1`: TCR Enterprise
	RepoType *int64 `json:"RepoType,omitempty" name:"RepoType"`

	// TCR Enterprise instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Address of the bound image server
	RepoServer *string `json:"RepoServer,omitempty" name:"RepoServer"`

	// Name of the bound image repository
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`

	// Source channel
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// Application subnet
	SubnetList []*string `json:"SubnetList,omitempty" name:"SubnetList"`

	// Programming language 
	// - JAVA
	// - OTHER
	CodingLanguage *string `json:"CodingLanguage,omitempty" name:"CodingLanguage"`

	// Deployment mode 
	// - IMAGE
	// - JAR
	// - WAR
	DeployMode *string `json:"DeployMode,omitempty" name:"DeployMode"`

	// Whether to enable APM tracing for the Java application. `1`: Enable, `0`: Disable
	EnableTracing *int64 `json:"EnableTracing,omitempty" name:"EnableTracing"`

	// Parameters of the default image service
	UseDefaultImageServiceParameters *UseDefaultRepoParameters `json:"UseDefaultImageServiceParameters,omitempty" name:"UseDefaultImageServiceParameters"`

	// Tag
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

type CreateApplicationRequest struct {
	*tchttp.BaseRequest
	
	// Application name
	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`

	// Description
	Description *string `json:"Description,omitempty" name:"Description"`

	// Whether to use the default image service. `1`: yes; `0`: no
	UseDefaultImageService *int64 `json:"UseDefaultImageService,omitempty" name:"UseDefaultImageService"`

	// Type of the bound repository. `0`: TCR Personal; `1`: TCR Enterprise
	RepoType *int64 `json:"RepoType,omitempty" name:"RepoType"`

	// TCR Enterprise instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Address of the bound image server
	RepoServer *string `json:"RepoServer,omitempty" name:"RepoServer"`

	// Name of the bound image repository
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`

	// Source channel
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// Application subnet
	SubnetList []*string `json:"SubnetList,omitempty" name:"SubnetList"`

	// Programming language 
	// - JAVA
	// - OTHER
	CodingLanguage *string `json:"CodingLanguage,omitempty" name:"CodingLanguage"`

	// Deployment mode 
	// - IMAGE
	// - JAR
	// - WAR
	DeployMode *string `json:"DeployMode,omitempty" name:"DeployMode"`

	// Whether to enable APM tracing for the Java application. `1`: Enable, `0`: Disable
	EnableTracing *int64 `json:"EnableTracing,omitempty" name:"EnableTracing"`

	// Parameters of the default image service
	UseDefaultImageServiceParameters *UseDefaultRepoParameters `json:"UseDefaultImageServiceParameters,omitempty" name:"UseDefaultImageServiceParameters"`

	// Tag
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

func (r *CreateApplicationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateApplicationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationName")
	delete(f, "Description")
	delete(f, "UseDefaultImageService")
	delete(f, "RepoType")
	delete(f, "InstanceId")
	delete(f, "RepoServer")
	delete(f, "RepoName")
	delete(f, "SourceChannel")
	delete(f, "SubnetList")
	delete(f, "CodingLanguage")
	delete(f, "DeployMode")
	delete(f, "EnableTracing")
	delete(f, "UseDefaultImageServiceParameters")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateApplicationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateApplicationResponseParams struct {
	// ID of the created application
	Result *string `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateApplicationResponse struct {
	*tchttp.BaseResponse
	Response *CreateApplicationResponseParams `json:"Response"`
}

func (r *CreateApplicationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateApplicationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateApplicationServiceRequestParams struct {
	// Application ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// Environment ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// Source channel. Please keep the default value.
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// Details of the access policy
	Service *ServicePortMapping `json:"Service,omitempty" name:"Service"`
}

type CreateApplicationServiceRequest struct {
	*tchttp.BaseRequest
	
	// Application ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// Environment ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// Source channel. Please keep the default value.
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// Details of the access policy
	Service *ServicePortMapping `json:"Service,omitempty" name:"Service"`
}

func (r *CreateApplicationServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateApplicationServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	delete(f, "EnvironmentId")
	delete(f, "SourceChannel")
	delete(f, "Service")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateApplicationServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateApplicationServiceResponseParams struct {
	// Whether the action succeeded 
	// Note: This field may return null, indicating that no valid values can be obtained.
	Result *bool `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateApplicationServiceResponse struct {
	*tchttp.BaseResponse
	Response *CreateApplicationServiceResponseParams `json:"Response"`
}

func (r *CreateApplicationServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateApplicationServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateConfigDataRequestParams struct {
	// Environment ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// Configuration name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Source channel. Please keep the default value.
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// Configuration information
	Data []*Pair `json:"Data,omitempty" name:"Data"`
}

type CreateConfigDataRequest struct {
	*tchttp.BaseRequest
	
	// Environment ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// Configuration name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Source channel. Please keep the default value.
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// Configuration information
	Data []*Pair `json:"Data,omitempty" name:"Data"`
}

func (r *CreateConfigDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateConfigDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvironmentId")
	delete(f, "Name")
	delete(f, "SourceChannel")
	delete(f, "Data")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateConfigDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateConfigDataResponseParams struct {
	// Whether the creation is successful
	Result *bool `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateConfigDataResponse struct {
	*tchttp.BaseResponse
	Response *CreateConfigDataResponseParams `json:"Response"`
}

func (r *CreateConfigDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateConfigDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

	// Environment description
	Description *string `json:"Description,omitempty" name:"Description"`

	// VPC name
	Vpc *string `json:"Vpc,omitempty" name:"Vpc"`

	// List of subnets
	SubnetIds []*string `json:"SubnetIds,omitempty" name:"SubnetIds"`

	// Kubernetes version
	K8sVersion *string `json:"K8sVersion,omitempty" name:"K8sVersion"`

	// Source channel
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// Whether to enable the TSW service
	EnableTswTraceService *bool `json:"EnableTswTraceService,omitempty" name:"EnableTswTraceService"`

	// Tag
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// Environment type. Values: `test`, `pre`, `prod`
	EnvType *string `json:"EnvType,omitempty" name:"EnvType"`

	// The region to create the environment
	CreateRegion *string `json:"CreateRegion,omitempty" name:"CreateRegion"`

	// Whether to create a VPC
	SetupVpc *bool `json:"SetupVpc,omitempty" name:"SetupVpc"`

	// Whether to create a TMP instance
	SetupPrometheus *bool `json:"SetupPrometheus,omitempty" name:"SetupPrometheus"`

	// TMP instance ID
	PrometheusId *string `json:"PrometheusId,omitempty" name:"PrometheusId"`

	// APM ID
	ApmId *string `json:"ApmId,omitempty" name:"ApmId"`
}

type CreateEnvironmentRequest struct {
	*tchttp.BaseRequest
	
	// Environment name
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// Environment description
	Description *string `json:"Description,omitempty" name:"Description"`

	// VPC name
	Vpc *string `json:"Vpc,omitempty" name:"Vpc"`

	// List of subnets
	SubnetIds []*string `json:"SubnetIds,omitempty" name:"SubnetIds"`

	// Kubernetes version
	K8sVersion *string `json:"K8sVersion,omitempty" name:"K8sVersion"`

	// Source channel
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// Whether to enable the TSW service
	EnableTswTraceService *bool `json:"EnableTswTraceService,omitempty" name:"EnableTswTraceService"`

	// Tag
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// Environment type. Values: `test`, `pre`, `prod`
	EnvType *string `json:"EnvType,omitempty" name:"EnvType"`

	// The region to create the environment
	CreateRegion *string `json:"CreateRegion,omitempty" name:"CreateRegion"`

	// Whether to create a VPC
	SetupVpc *bool `json:"SetupVpc,omitempty" name:"SetupVpc"`

	// Whether to create a TMP instance
	SetupPrometheus *bool `json:"SetupPrometheus,omitempty" name:"SetupPrometheus"`

	// TMP instance ID
	PrometheusId *string `json:"PrometheusId,omitempty" name:"PrometheusId"`

	// APM ID
	ApmId *string `json:"ApmId,omitempty" name:"ApmId"`
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
	delete(f, "Description")
	delete(f, "Vpc")
	delete(f, "SubnetIds")
	delete(f, "K8sVersion")
	delete(f, "SourceChannel")
	delete(f, "EnableTswTraceService")
	delete(f, "Tags")
	delete(f, "EnvType")
	delete(f, "CreateRegion")
	delete(f, "SetupVpc")
	delete(f, "SetupPrometheus")
	delete(f, "PrometheusId")
	delete(f, "ApmId")
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
type CreateLogConfigRequestParams struct {
	// Environment ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// Configuration name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Collection type. Values: `container_stdout` (standard); `container_file` (file)
	InputType *string `json:"InputType,omitempty" name:"InputType"`

	// Application ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// Logset ID
	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`

	// Log topic ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// Log withdrawal mode. Values: `minimalist_log` (full text in a single line); `multiline_log` (full text in multiple lines); `json_log` (JSON); `fullregex_log` (regex in a single line); `multiline_fullregex_log` (regex in multiple lines)
	LogType *string `json:"LogType,omitempty" name:"LogType"`

	// The first line regex. It’s valid when `LogType` is `multiline_log`.
	BeginningRegex *string `json:"BeginningRegex,omitempty" name:"BeginningRegex"`

	// Directory of files to collect. It’s valid when `InputType` is `container_file`.
	LogPath *string `json:"LogPath,omitempty" name:"LogPath"`

	// Name pattern of files to collect. It’s valid when `InputType` is `container_file`.
	FilePattern *string `json:"FilePattern,omitempty" name:"FilePattern"`

	// Export
	ExtractRule *LogConfigExtractRule `json:"ExtractRule,omitempty" name:"ExtractRule"`
}

type CreateLogConfigRequest struct {
	*tchttp.BaseRequest
	
	// Environment ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// Configuration name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Collection type. Values: `container_stdout` (standard); `container_file` (file)
	InputType *string `json:"InputType,omitempty" name:"InputType"`

	// Application ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// Logset ID
	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`

	// Log topic ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// Log withdrawal mode. Values: `minimalist_log` (full text in a single line); `multiline_log` (full text in multiple lines); `json_log` (JSON); `fullregex_log` (regex in a single line); `multiline_fullregex_log` (regex in multiple lines)
	LogType *string `json:"LogType,omitempty" name:"LogType"`

	// The first line regex. It’s valid when `LogType` is `multiline_log`.
	BeginningRegex *string `json:"BeginningRegex,omitempty" name:"BeginningRegex"`

	// Directory of files to collect. It’s valid when `InputType` is `container_file`.
	LogPath *string `json:"LogPath,omitempty" name:"LogPath"`

	// Name pattern of files to collect. It’s valid when `InputType` is `container_file`.
	FilePattern *string `json:"FilePattern,omitempty" name:"FilePattern"`

	// Export
	ExtractRule *LogConfigExtractRule `json:"ExtractRule,omitempty" name:"ExtractRule"`
}

func (r *CreateLogConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLogConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvironmentId")
	delete(f, "Name")
	delete(f, "InputType")
	delete(f, "ApplicationId")
	delete(f, "LogsetId")
	delete(f, "TopicId")
	delete(f, "LogType")
	delete(f, "BeginningRegex")
	delete(f, "LogPath")
	delete(f, "FilePattern")
	delete(f, "ExtractRule")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateLogConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLogConfigResponseParams struct {
	// Whether the creation is successful
	Result *bool `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateLogConfigResponse struct {
	*tchttp.BaseResponse
	Response *CreateLogConfigResponseParams `json:"Response"`
}

func (r *CreateLogConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLogConfigResponse) FromJsonString(s string) error {
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
type DeleteApplicationAutoscalerRequestParams struct {
	// Application ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// Environment ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// Source channel. Please keep the default value.
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// Scaling rule ID
	AutoscalerId *string `json:"AutoscalerId,omitempty" name:"AutoscalerId"`
}

type DeleteApplicationAutoscalerRequest struct {
	*tchttp.BaseRequest
	
	// Application ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// Environment ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// Source channel. Please keep the default value.
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// Scaling rule ID
	AutoscalerId *string `json:"AutoscalerId,omitempty" name:"AutoscalerId"`
}

func (r *DeleteApplicationAutoscalerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteApplicationAutoscalerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	delete(f, "EnvironmentId")
	delete(f, "SourceChannel")
	delete(f, "AutoscalerId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteApplicationAutoscalerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteApplicationAutoscalerResponseParams struct {
	// Whether the action is successful
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Result *bool `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteApplicationAutoscalerResponse struct {
	*tchttp.BaseResponse
	Response *DeleteApplicationAutoscalerResponseParams `json:"Response"`
}

func (r *DeleteApplicationAutoscalerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteApplicationAutoscalerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteApplicationRequestParams struct {
	// Application ID.
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// Environment ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// Source channel. Please keep the default value.
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// Whether to delete this application automatically when there is no running version.
	DeleteApplicationIfNoRunningVersion *bool `json:"DeleteApplicationIfNoRunningVersion,omitempty" name:"DeleteApplicationIfNoRunningVersion"`
}

type DeleteApplicationRequest struct {
	*tchttp.BaseRequest
	
	// Application ID.
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// Environment ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// Source channel. Please keep the default value.
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// Whether to delete this application automatically when there is no running version.
	DeleteApplicationIfNoRunningVersion *bool `json:"DeleteApplicationIfNoRunningVersion,omitempty" name:"DeleteApplicationIfNoRunningVersion"`
}

func (r *DeleteApplicationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteApplicationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	delete(f, "EnvironmentId")
	delete(f, "SourceChannel")
	delete(f, "DeleteApplicationIfNoRunningVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteApplicationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteApplicationResponseParams struct {
	// Returned result.
	Result *bool `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteApplicationResponse struct {
	*tchttp.BaseResponse
	Response *DeleteApplicationResponseParams `json:"Response"`
}

func (r *DeleteApplicationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteApplicationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteApplicationServiceRequestParams struct {
	// Application ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// Source channel. Please keep the default value.
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// Environment ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// Service name
	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
}

type DeleteApplicationServiceRequest struct {
	*tchttp.BaseRequest
	
	// Application ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// Source channel. Please keep the default value.
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// Environment ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// Service name
	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
}

func (r *DeleteApplicationServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteApplicationServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	delete(f, "SourceChannel")
	delete(f, "EnvironmentId")
	delete(f, "ServiceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteApplicationServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteApplicationServiceResponseParams struct {
	// Whether the action succeeded 
	// Note: This field may return null, indicating that no valid values can be obtained.
	Result *bool `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteApplicationServiceResponse struct {
	*tchttp.BaseResponse
	Response *DeleteApplicationServiceResponseParams `json:"Response"`
}

func (r *DeleteApplicationServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteApplicationServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

	// Auto scaling policy. (Disused. Please use APIs for auto scaling policy combinations)
	HorizontalAutoscaler []*HorizontalAutoscaler `json:"HorizontalAutoscaler,omitempty" name:"HorizontalAutoscaler"`

	// Scheduled scaling policy (Disused. Please use APIs for auto scaling policy combinations)
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

	// Configuration of metrics of this application
	EnablePrometheusConf *EnablePrometheusConf `json:"EnablePrometheusConf,omitempty" name:"EnablePrometheusConf"`

	// `1`: Automatically enable APM tracing (Skywalking)
	// `0`: Disable APM tracing
	EnableTracing *int64 `json:"EnableTracing,omitempty" name:"EnableTracing"`

	// `1`: Automatically enable metrics collection (open-telemetry)
	// `0`: Disable metrics collection
	EnableMetrics *int64 `json:"EnableMetrics,omitempty" name:"EnableMetrics"`

	// ID of the TCR instance used for image deployment
	TcrInstanceId *string `json:"TcrInstanceId,omitempty" name:"TcrInstanceId"`

	// Image server address for image deployment
	RepoServer *string `json:"RepoServer,omitempty" name:"RepoServer"`

	// Type of the repository. `0`: TCR Personal; `1`: TCR Enterprise; `2`: Public repository; `3`: TEM hosted repository; `4`: Demo repository
	RepoType *int64 `json:"RepoType,omitempty" name:"RepoType"`
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

	// Auto scaling policy. (Disused. Please use APIs for auto scaling policy combinations)
	HorizontalAutoscaler []*HorizontalAutoscaler `json:"HorizontalAutoscaler,omitempty" name:"HorizontalAutoscaler"`

	// Scheduled scaling policy (Disused. Please use APIs for auto scaling policy combinations)
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

	// Configuration of metrics of this application
	EnablePrometheusConf *EnablePrometheusConf `json:"EnablePrometheusConf,omitempty" name:"EnablePrometheusConf"`

	// `1`: Automatically enable APM tracing (Skywalking)
	// `0`: Disable APM tracing
	EnableTracing *int64 `json:"EnableTracing,omitempty" name:"EnableTracing"`

	// `1`: Automatically enable metrics collection (open-telemetry)
	// `0`: Disable metrics collection
	EnableMetrics *int64 `json:"EnableMetrics,omitempty" name:"EnableMetrics"`

	// ID of the TCR instance used for image deployment
	TcrInstanceId *string `json:"TcrInstanceId,omitempty" name:"TcrInstanceId"`

	// Image server address for image deployment
	RepoServer *string `json:"RepoServer,omitempty" name:"RepoServer"`

	// Type of the repository. `0`: TCR Personal; `1`: TCR Enterprise; `2`: Public repository; `3`: TEM hosted repository; `4`: Demo repository
	RepoType *int64 `json:"RepoType,omitempty" name:"RepoType"`
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
	delete(f, "EnableTracing")
	delete(f, "EnableMetrics")
	delete(f, "TcrInstanceId")
	delete(f, "RepoServer")
	delete(f, "RepoType")
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
type DescribeApplicationAutoscalerListRequestParams struct {
	// Application ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// Environment ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// Source channel. Please keep the default value.
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`
}

type DescribeApplicationAutoscalerListRequest struct {
	*tchttp.BaseRequest
	
	// Application ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// Environment ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// Source channel. Please keep the default value.
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`
}

func (r *DescribeApplicationAutoscalerListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationAutoscalerListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	delete(f, "EnvironmentId")
	delete(f, "SourceChannel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApplicationAutoscalerListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApplicationAutoscalerListResponseParams struct {
	// Scaling rule
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Result []*Autoscaler `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeApplicationAutoscalerListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApplicationAutoscalerListResponseParams `json:"Response"`
}

func (r *DescribeApplicationAutoscalerListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationAutoscalerListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApplicationInfoRequestParams struct {
	// Application ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// Source channel. Please keep the default value.
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// Environment ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`
}

type DescribeApplicationInfoRequest struct {
	*tchttp.BaseRequest
	
	// Application ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// Source channel. Please keep the default value.
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// Environment ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`
}

func (r *DescribeApplicationInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	delete(f, "SourceChannel")
	delete(f, "EnvironmentId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApplicationInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApplicationInfoResponseParams struct {
	// Returned result.
	Result *TemServiceVersionInfo `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeApplicationInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApplicationInfoResponseParams `json:"Response"`
}

func (r *DescribeApplicationInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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
type DescribeApplicationServiceListRequestParams struct {
	// ID of the environment
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// ID of the application
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// xx
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`
}

type DescribeApplicationServiceListRequest struct {
	*tchttp.BaseRequest
	
	// ID of the environment
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// ID of the application
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// xx
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`
}

func (r *DescribeApplicationServiceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationServiceListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvironmentId")
	delete(f, "ApplicationId")
	delete(f, "SourceChannel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApplicationServiceListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApplicationServiceListResponseParams struct {
	// Application EKS service list
	Result *EksService `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeApplicationServiceListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApplicationServiceListResponseParams `json:"Response"`
}

func (r *DescribeApplicationServiceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationServiceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApplicationsRequestParams struct {
	// ID of the environment
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// Pagination limit
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Pagination offset
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Source channel. Please keep the default value.
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// Application ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// Keyword for searching.
	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`

	// Filters for query 
	Filters []*QueryFilter `json:"Filters,omitempty" name:"Filters"`

	// Sorting field
	SortInfo *SortType `json:"SortInfo,omitempty" name:"SortInfo"`
}

type DescribeApplicationsRequest struct {
	*tchttp.BaseRequest
	
	// ID of the environment
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// Pagination limit
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Pagination offset
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Source channel. Please keep the default value.
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// Application ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// Keyword for searching.
	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`

	// Filters for query 
	Filters []*QueryFilter `json:"Filters,omitempty" name:"Filters"`

	// Sorting field
	SortInfo *SortType `json:"SortInfo,omitempty" name:"SortInfo"`
}

func (r *DescribeApplicationsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvironmentId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "SourceChannel")
	delete(f, "ApplicationId")
	delete(f, "Keyword")
	delete(f, "Filters")
	delete(f, "SortInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApplicationsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApplicationsResponseParams struct {
	// Returned result.
	Result *ServicePage `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeApplicationsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApplicationsResponseParams `json:"Response"`
}

func (r *DescribeApplicationsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApplicationsStatusRequestParams struct {
	// Source channel. Please keep the default value.
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// Environment ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`
}

type DescribeApplicationsStatusRequest struct {
	*tchttp.BaseRequest
	
	// Source channel. Please keep the default value.
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// Environment ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`
}

func (r *DescribeApplicationsStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationsStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SourceChannel")
	delete(f, "EnvironmentId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApplicationsStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApplicationsStatusResponseParams struct {
	// Returned result.
	Result []*ServiceVersionBrief `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeApplicationsStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApplicationsStatusResponseParams `json:"Response"`
}

func (r *DescribeApplicationsStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationsStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeConfigDataListPage struct {
	// Record
	Records []*ConfigData `json:"Records,omitempty" name:"Records"`

	// Paging cursor
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	ContinueToken *string `json:"ContinueToken,omitempty" name:"ContinueToken"`

	// Remaining number
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	RemainingCount *int64 `json:"RemainingCount,omitempty" name:"RemainingCount"`
}

// Predefined struct for user
type DescribeConfigDataListRequestParams struct {
	// Environment ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// Source channel. Please keep the default value.
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// Paging cursor
	ContinueToken *string `json:"ContinueToken,omitempty" name:"ContinueToken"`

	// Pagination limit
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeConfigDataListRequest struct {
	*tchttp.BaseRequest
	
	// Environment ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// Source channel. Please keep the default value.
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// Paging cursor
	ContinueToken *string `json:"ContinueToken,omitempty" name:"ContinueToken"`

	// Pagination limit
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeConfigDataListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConfigDataListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvironmentId")
	delete(f, "SourceChannel")
	delete(f, "ContinueToken")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeConfigDataListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConfigDataListResponseParams struct {
	// Configuration list.
	Result *DescribeConfigDataListPage `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeConfigDataListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeConfigDataListResponseParams `json:"Response"`
}

func (r *DescribeConfigDataListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConfigDataListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConfigDataRequestParams struct {
	// Environment ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// Configuration name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Source channel. Please keep the default value.
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`
}

type DescribeConfigDataRequest struct {
	*tchttp.BaseRequest
	
	// Environment ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// Configuration name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Source channel. Please keep the default value.
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`
}

func (r *DescribeConfigDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConfigDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvironmentId")
	delete(f, "Name")
	delete(f, "SourceChannel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeConfigDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConfigDataResponseParams struct {
	// Configuration
	Result *ConfigData `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeConfigDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeConfigDataResponseParams `json:"Response"`
}

func (r *DescribeConfigDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConfigDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEnvironmentRequestParams struct {
	// Namespace ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// Source channel
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`
}

type DescribeEnvironmentRequest struct {
	*tchttp.BaseRequest
	
	// Namespace ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// Source channel
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`
}

func (r *DescribeEnvironmentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEnvironmentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvironmentId")
	delete(f, "SourceChannel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEnvironmentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEnvironmentResponseParams struct {
	// Environment information
	Result *NamespaceInfo `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeEnvironmentResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEnvironmentResponseParams `json:"Response"`
}

func (r *DescribeEnvironmentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEnvironmentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEnvironmentStatusRequestParams struct {
	// ID of the environment
	EnvironmentIds []*string `json:"EnvironmentIds,omitempty" name:"EnvironmentIds"`

	// Source channel
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`
}

type DescribeEnvironmentStatusRequest struct {
	*tchttp.BaseRequest
	
	// ID of the environment
	EnvironmentIds []*string `json:"EnvironmentIds,omitempty" name:"EnvironmentIds"`

	// Source channel
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`
}

func (r *DescribeEnvironmentStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEnvironmentStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvironmentIds")
	delete(f, "SourceChannel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEnvironmentStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEnvironmentStatusResponseParams struct {
	// List of environment status
	Result []*NamespaceStatusInfo `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeEnvironmentStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEnvironmentStatusResponseParams `json:"Response"`
}

func (r *DescribeEnvironmentStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEnvironmentStatusResponse) FromJsonString(s string) error {
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

	// Filters for query 
	Filters []*QueryFilter `json:"Filters,omitempty" name:"Filters"`

	// Sorting field
	SortInfo *SortType `json:"SortInfo,omitempty" name:"SortInfo"`

	// Environment ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`
}

type DescribeEnvironmentsRequest struct {
	*tchttp.BaseRequest
	
	// Pagination limit
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Page offset
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Source
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// Filters for query 
	Filters []*QueryFilter `json:"Filters,omitempty" name:"Filters"`

	// Sorting field
	SortInfo *SortType `json:"SortInfo,omitempty" name:"SortInfo"`

	// Environment ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`
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
	delete(f, "Filters")
	delete(f, "SortInfo")
	delete(f, "EnvironmentId")
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
type DescribeLogConfigRequestParams struct {
	// Environment ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// Configuration name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Application ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`
}

type DescribeLogConfigRequest struct {
	*tchttp.BaseRequest
	
	// Environment ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// Configuration name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Application ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`
}

func (r *DescribeLogConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvironmentId")
	delete(f, "Name")
	delete(f, "ApplicationId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLogConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLogConfigResponseParams struct {
	// Configuration
	Result *LogConfig `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeLogConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLogConfigResponseParams `json:"Response"`
}

func (r *DescribeLogConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePagedLogConfigListRequestParams struct {
	// Environment ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// Application ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// Application name
	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`

	// Name of the rule
	Name *string `json:"Name,omitempty" name:"Name"`

	// Number of entries per page. Default value: 20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Paging cursor
	ContinueToken *string `json:"ContinueToken,omitempty" name:"ContinueToken"`
}

type DescribePagedLogConfigListRequest struct {
	*tchttp.BaseRequest
	
	// Environment ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// Application ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// Application name
	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`

	// Name of the rule
	Name *string `json:"Name,omitempty" name:"Name"`

	// Number of entries per page. Default value: 20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Paging cursor
	ContinueToken *string `json:"ContinueToken,omitempty" name:"ContinueToken"`
}

func (r *DescribePagedLogConfigListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePagedLogConfigListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvironmentId")
	delete(f, "ApplicationId")
	delete(f, "ApplicationName")
	delete(f, "Name")
	delete(f, "Limit")
	delete(f, "ContinueToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePagedLogConfigListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePagedLogConfigListResponseParams struct {
	// List of log collecting configurations
	Result *LogConfigListPage `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePagedLogConfigListResponse struct {
	*tchttp.BaseResponse
	Response *DescribePagedLogConfigListResponseParams `json:"Response"`
}

func (r *DescribePagedLogConfigListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePagedLogConfigListResponse) FromJsonString(s string) error {
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

// Predefined struct for user
type DestroyConfigDataRequestParams struct {
	// Environment ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// Configuration name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Source channel. Please keep the default value.
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`
}

type DestroyConfigDataRequest struct {
	*tchttp.BaseRequest
	
	// Environment ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// Configuration name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Source channel. Please keep the default value.
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`
}

func (r *DestroyConfigDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyConfigDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvironmentId")
	delete(f, "Name")
	delete(f, "SourceChannel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DestroyConfigDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroyConfigDataResponseParams struct {
	// Returned result.
	Result *bool `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DestroyConfigDataResponse struct {
	*tchttp.BaseResponse
	Response *DestroyConfigDataResponseParams `json:"Response"`
}

func (r *DestroyConfigDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyConfigDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroyEnvironmentRequestParams struct {
	// Namespace ID.
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// Namespace
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`
}

type DestroyEnvironmentRequest struct {
	*tchttp.BaseRequest
	
	// Namespace ID.
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// Namespace
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`
}

func (r *DestroyEnvironmentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyEnvironmentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvironmentId")
	delete(f, "SourceChannel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DestroyEnvironmentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroyEnvironmentResponseParams struct {
	// Returned result.
	Result *bool `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DestroyEnvironmentResponse struct {
	*tchttp.BaseResponse
	Response *DestroyEnvironmentResponseParams `json:"Response"`
}

func (r *DestroyEnvironmentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyEnvironmentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroyLogConfigRequestParams struct {
	// Environment ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// Configuration name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Application ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`
}

type DestroyLogConfigRequest struct {
	*tchttp.BaseRequest
	
	// Environment ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// Configuration name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Application ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`
}

func (r *DestroyLogConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyLogConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvironmentId")
	delete(f, "Name")
	delete(f, "ApplicationId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DestroyLogConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroyLogConfigResponseParams struct {
	// Returned result.
	Result *bool `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DestroyLogConfigResponse struct {
	*tchttp.BaseResponse
	Response *DestroyLogConfigResponseParams `json:"Response"`
}

func (r *DestroyLogConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyLogConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableApplicationAutoscalerRequestParams struct {
	// Application ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// Environment ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// Source channel. Please keep the default value.
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// Scaling rule ID
	AutoscalerId *string `json:"AutoscalerId,omitempty" name:"AutoscalerId"`
}

type DisableApplicationAutoscalerRequest struct {
	*tchttp.BaseRequest
	
	// Application ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// Environment ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// Source channel. Please keep the default value.
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// Scaling rule ID
	AutoscalerId *string `json:"AutoscalerId,omitempty" name:"AutoscalerId"`
}

func (r *DisableApplicationAutoscalerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableApplicationAutoscalerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	delete(f, "EnvironmentId")
	delete(f, "SourceChannel")
	delete(f, "AutoscalerId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisableApplicationAutoscalerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableApplicationAutoscalerResponseParams struct {
	// Whether the action succeeded 
	// Note: This field may return null, indicating that no valid values can be obtained.
	Result *bool `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DisableApplicationAutoscalerResponse struct {
	*tchttp.BaseResponse
	Response *DisableApplicationAutoscalerResponseParams `json:"Response"`
}

func (r *DisableApplicationAutoscalerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableApplicationAutoscalerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

	// Details of each type of access configuration
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	ServicePortMappingList []*ServicePortMapping `json:"ServicePortMappingList,omitempty" name:"ServicePortMappingList"`

	// Flush all types
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	FlushAll *bool `json:"FlushAll,omitempty" name:"FlushAll"`

	// `0`: Do not inject. `1`: Inject registry information automatically for the next deployment
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	EnableRegistryNextDeploy *int64 `json:"EnableRegistryNextDeploy,omitempty" name:"EnableRegistryNextDeploy"`

	// The application ID returned.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// Whether all the application IPs are ready
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	AllIpDone *bool `json:"AllIpDone,omitempty" name:"AllIpDone"`

	// CLB domain name
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	ExternalDomain *string `json:"ExternalDomain,omitempty" name:"ExternalDomain"`
}

// Predefined struct for user
type EnableApplicationAutoscalerRequestParams struct {
	// Service ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// Environment ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// Source channel. Please keep the default value.
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// Scaling rule ID
	AutoscalerId *string `json:"AutoscalerId,omitempty" name:"AutoscalerId"`
}

type EnableApplicationAutoscalerRequest struct {
	*tchttp.BaseRequest
	
	// Service ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// Environment ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// Source channel. Please keep the default value.
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// Scaling rule ID
	AutoscalerId *string `json:"AutoscalerId,omitempty" name:"AutoscalerId"`
}

func (r *EnableApplicationAutoscalerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableApplicationAutoscalerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	delete(f, "EnvironmentId")
	delete(f, "SourceChannel")
	delete(f, "AutoscalerId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EnableApplicationAutoscalerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableApplicationAutoscalerResponseParams struct {
	// Whether the action succeeded 
	// Note: This field may return null, indicating that no valid values can be obtained.
	Result *bool `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type EnableApplicationAutoscalerResponse struct {
	*tchttp.BaseResponse
	Response *EnableApplicationAutoscalerResponseParams `json:"Response"`
}

func (r *EnableApplicationAutoscalerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableApplicationAutoscalerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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
	// (Optional) Minimum number of instances
	MinReplicas *int64 `json:"MinReplicas,omitempty" name:"MinReplicas"`

	// (Optional) Maximum number of instances
	MaxReplicas *int64 `json:"MaxReplicas,omitempty" name:"MaxReplicas"`

	// Metric measurement
	// `CPU`: CPU utilization (%)
	// `MEMORY`: Memory utilization (%)
	// `CPU_CORE_USED`: CPU usage (core)
	// `MEMORY_SIZE_USED`: Memory usage (MiB)
	// `NETWORK_BANDWIDTH_RECEIVE`: Network bandwidth in (Mbps)
	// `NETWORK_BANDWIDTH_TRANSMIT`: Network bandwidth out (Mbps)
	// `NETWORK_TRAFFIC_RECEIVE`: Network traffic in (MiB/s)
	// `NETWORK_TRAFFIC_TRANSMIT`: Network traffic  out (MiB/s)
	// `NETWORK_PACKETS_RECEIVE`: Network packets in (packets/sec)
	// `NETWORK_PACKETS_TRANSMIT`: Network packets out (packets/sec)
	// `FS_IOPS_WRITE`: Disk writes (count/sec)
	// `FS_IOPS_READ`: Disk reads (count/sec)
	// `FS_SIZE_WRITE`: Disk write size (MiB/s)
	// `FS_SIZE_READ`: Disk read size (MiB/s)
	Metrics *string `json:"Metrics,omitempty" name:"Metrics"`

	// The value of threshold (integer)
	Threshold *int64 `json:"Threshold,omitempty" name:"Threshold"`

	// Whether it is enabled
	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`

	// The value of threshold (demical)
	// Note: This field may return null, indicating that no valid values can be obtained.
	DoubleThreshold *float64 `json:"DoubleThreshold,omitempty" name:"DoubleThreshold"`
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

	// CLB domain name
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Domain *string `json:"Domain,omitempty" name:"Domain"`
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

type LogConfig struct {
	// Name.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Collection type. Values: `container_stdout` (standard); `container_file` (file)
	InputType *string `json:"InputType,omitempty" name:"InputType"`

	// Logset ID
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`

	// Log topic ID
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// Log withdrawal mode. Values: `minimalist_log` (full text in a single line); `multiline_log` (full text in multiple lines); `fullregex_log` (regex in a single line); `multiline_fullregex_log` (regex in multiple lines), `json_log` (JSON); 
	LogType *string `json:"LogType,omitempty" name:"LogType"`

	// First line regex. It’s valid when `LogType` is `multiline_log` or `multiline_fullregex_log`.
	// Note: This field may return `null`, indicating that no valid value was found.
	BeginningRegex *string `json:"BeginningRegex,omitempty" name:"BeginningRegex"`

	// Directory of files to collect. It’s valid when `InputType` is `container_file`.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	LogPath *string `json:"LogPath,omitempty" name:"LogPath"`

	// Name pattern of files to collect. It’s valid when `InputType` is `container_file`.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	FilePattern *string `json:"FilePattern,omitempty" name:"FilePattern"`

	// Creation time.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	CreateDate *string `json:"CreateDate,omitempty" name:"CreateDate"`

	// Update time
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	ModifyDate *string `json:"ModifyDate,omitempty" name:"ModifyDate"`

	// Application ID
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// Application name
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`

	// Export rules
	// Note: This field may return `null`, indicating that no valid value was found.
	ExtractRule *LogConfigExtractRule `json:"ExtractRule,omitempty" name:"ExtractRule"`
}

type LogConfigExtractRule struct {
	// First line regex
	// Note: This field may return `null`, indicating that no valid value was found.
	BeginningRegex *string `json:"BeginningRegex,omitempty" name:"BeginningRegex"`

	// Withdrawl result
	// Note: This field may return `null`, indicating that no valid value was found.
	Keys []*string `json:"Keys,omitempty" name:"Keys"`

	// Filter keys
	// Note: This field may return `null`, indicating that no valid value was found.
	FilterKeys []*string `json:"FilterKeys,omitempty" name:"FilterKeys"`

	// Filter values
	// Note: This field may return `null`, indicating that no valid value was found.
	FilterRegex []*string `json:"FilterRegex,omitempty" name:"FilterRegex"`

	// Log regex
	// Note: This field may return `null`, indicating that no valid value was found.
	LogRegex *string `json:"LogRegex,omitempty" name:"LogRegex"`

	// Time field
	// Note: This field may return `null`, indicating that no valid value was found.
	TimeKey *string `json:"TimeKey,omitempty" name:"TimeKey"`

	// Time Format
	// Note: This field may return `null`, indicating that no valid value was found.
	TimeFormat *string `json:"TimeFormat,omitempty" name:"TimeFormat"`

	// - Enable the upload of the log that failed to parse
	// Note: This field may return `null`, indicating that no valid value was found.
	UnMatchUpload *string `json:"UnMatchUpload,omitempty" name:"UnMatchUpload"`

	// Key of log failed to be parsed
	// Note: This field may return `null`, indicating that no valid value was found.
	UnMatchedKey *string `json:"UnMatchedKey,omitempty" name:"UnMatchedKey"`

	// tracking
	// Note: This field may return null, indicating that no valid values can be obtained.
	Backtracking *string `json:"Backtracking,omitempty" name:"Backtracking"`

	// Separator
	// Note: This field may return null, indicating that no valid values can be obtained.
	Delimiter *string `json:"Delimiter,omitempty" name:"Delimiter"`
}

type LogConfigListPage struct {
	// Record
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Records []*LogConfig `json:"Records,omitempty" name:"Records"`

	// Paging cursor
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	ContinueToken *string `json:"ContinueToken,omitempty" name:"ContinueToken"`
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
type ModifyApplicationAutoscalerRequestParams struct {
	// Application ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// Environment ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// Source channel. Please keep the default value.
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// Scaling rule ID
	AutoscalerId *string `json:"AutoscalerId,omitempty" name:"AutoscalerId"`

	// Auto scaling policy
	Autoscaler *Autoscaler `json:"Autoscaler,omitempty" name:"Autoscaler"`
}

type ModifyApplicationAutoscalerRequest struct {
	*tchttp.BaseRequest
	
	// Application ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// Environment ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// Source channel. Please keep the default value.
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// Scaling rule ID
	AutoscalerId *string `json:"AutoscalerId,omitempty" name:"AutoscalerId"`

	// Auto scaling policy
	Autoscaler *Autoscaler `json:"Autoscaler,omitempty" name:"Autoscaler"`
}

func (r *ModifyApplicationAutoscalerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApplicationAutoscalerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	delete(f, "EnvironmentId")
	delete(f, "SourceChannel")
	delete(f, "AutoscalerId")
	delete(f, "Autoscaler")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyApplicationAutoscalerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyApplicationAutoscalerResponseParams struct {
	// Whether the action is successful
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Result *bool `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyApplicationAutoscalerResponse struct {
	*tchttp.BaseResponse
	Response *ModifyApplicationAutoscalerResponseParams `json:"Response"`
}

func (r *ModifyApplicationAutoscalerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApplicationAutoscalerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyApplicationInfoRequestParams struct {
	// Application ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// Description
	Description *string `json:"Description,omitempty" name:"Description"`

	// Source channel
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// (Disused) Whether to enable the call chain. 
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

	// (Disused) Whether to enable the call chain. 
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
type ModifyApplicationServiceRequestParams struct {
	// Application ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// Environment ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// Source channel. Please keep the default value.
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// Full access mode settings
	Service *EksService `json:"Service,omitempty" name:"Service"`

	// Single entry access mode settings
	Data *ServicePortMapping `json:"Data,omitempty" name:"Data"`
}

type ModifyApplicationServiceRequest struct {
	*tchttp.BaseRequest
	
	// Application ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// Environment ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// Source channel. Please keep the default value.
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// Full access mode settings
	Service *EksService `json:"Service,omitempty" name:"Service"`

	// Single entry access mode settings
	Data *ServicePortMapping `json:"Data,omitempty" name:"Data"`
}

func (r *ModifyApplicationServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApplicationServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	delete(f, "EnvironmentId")
	delete(f, "SourceChannel")
	delete(f, "Service")
	delete(f, "Data")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyApplicationServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyApplicationServiceResponseParams struct {
	// Whether the action succeeded 
	// Note: This field may return null, indicating that no valid values can be obtained.
	Result *bool `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyApplicationServiceResponse struct {
	*tchttp.BaseResponse
	Response *ModifyApplicationServiceResponseParams `json:"Response"`
}

func (r *ModifyApplicationServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApplicationServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyConfigDataRequestParams struct {
	// Environment ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// Configuration name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Source channel. Please keep the default value.
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// Configuration information
	Data []*Pair `json:"Data,omitempty" name:"Data"`
}

type ModifyConfigDataRequest struct {
	*tchttp.BaseRequest
	
	// Environment ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// Configuration name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Source channel. Please keep the default value.
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// Configuration information
	Data []*Pair `json:"Data,omitempty" name:"Data"`
}

func (r *ModifyConfigDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyConfigDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvironmentId")
	delete(f, "Name")
	delete(f, "SourceChannel")
	delete(f, "Data")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyConfigDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyConfigDataResponseParams struct {
	// Result of the modification
	Result *bool `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyConfigDataResponse struct {
	*tchttp.BaseResponse
	Response *ModifyConfigDataResponseParams `json:"Response"`
}

func (r *ModifyConfigDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyConfigDataResponse) FromJsonString(s string) error {
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

	// Environment type. Values: `test`, `pre`, `prod`
	EnvType *string `json:"EnvType,omitempty" name:"EnvType"`
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

	// Environment type. Values: `test`, `pre`, `prod`
	EnvType *string `json:"EnvType,omitempty" name:"EnvType"`
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
	delete(f, "EnvType")
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

// Predefined struct for user
type ModifyLogConfigRequestParams struct {
	// Environment ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// Configuration name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Log collector configuration
	Data *LogConfig `json:"Data,omitempty" name:"Data"`

	// Application ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`
}

type ModifyLogConfigRequest struct {
	*tchttp.BaseRequest
	
	// Environment ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// Configuration name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Log collector configuration
	Data *LogConfig `json:"Data,omitempty" name:"Data"`

	// Application ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`
}

func (r *ModifyLogConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLogConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvironmentId")
	delete(f, "Name")
	delete(f, "Data")
	delete(f, "ApplicationId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLogConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLogConfigResponseParams struct {
	// Result of the modification
	Result *bool `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyLogConfigResponse struct {
	*tchttp.BaseResponse
	Response *ModifyLogConfigResponseParams `json:"Response"`
}

func (r *ModifyLogConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLogConfigResponse) FromJsonString(s string) error {
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

type NamespaceInfo struct {
	// ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// (Disused) Name
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// Region
	Region *string `json:"Region,omitempty" name:"Region"`

	// vpc id
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Array of subnet IDs
	SubnetIds []*string `json:"SubnetIds,omitempty" name:"SubnetIds"`

	// Description
	Description *string `json:"Description,omitempty" name:"Description"`

	// Creation time
	CreatedDate *string `json:"CreatedDate,omitempty" name:"CreatedDate"`

	// Environment name
	// Note: This field may return null, indicating that no valid values can be obtained.
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// APM instance ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	ApmInstanceId *string `json:"ApmInstanceId,omitempty" name:"ApmInstanceId"`

	// Whether the environment is locked. `1`: Locked, `0`: Not locked
	// Note: This field may return null, indicating that no valid values can be obtained.
	Locked *int64 `json:"Locked,omitempty" name:"Locked"`

	// Tag
	// Note: This field may return `null`, indicating that no valid value was found.
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// Environment type. Values: `test`, `pre`, `prod`
	// Note: This field may return `null`, indicating that no valid value was found.
	EnvType *string `json:"EnvType,omitempty" name:"EnvType"`
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

	// Current entry
	// Note: This field may return null, indicating that no valid values can be obtained.
	Current *int64 `json:"Current,omitempty" name:"Current"`
}

type NamespaceStatusInfo struct {
	// ID of the environment
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// Environment name
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// TCB envId | EKS clusterId
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Environment status
	ClusterStatus *string `json:"ClusterStatus,omitempty" name:"ClusterStatus"`

	// Whether the environment is being started. `null` is returned if it’s not being started.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	EnvironmentStartingStatus *TemEnvironmentStartingStatus `json:"EnvironmentStartingStatus,omitempty" name:"EnvironmentStartingStatus"`

	// Whether the environment is being stopped. `null` is returned if it’s not being stopped.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	EnvironmentStoppingStatus *TemEnvironmentStoppingStatus `json:"EnvironmentStoppingStatus,omitempty" name:"EnvironmentStoppingStatus"`
}

type NodeInfo struct {
	// Node name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Availability zone of the node
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// Node subnet ID
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// Number of available IPs
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	AvailableIpCount *string `json:"AvailableIpCount,omitempty" name:"AvailableIpCount"`

	// CIDR block
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Cidr *string `json:"Cidr,omitempty" name:"Cidr"`
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

	// K8s service name
	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
}

type QueryFilter struct {
	// Name of the field to query
	Name *string `json:"Name,omitempty" name:"Name"`

	// Value of the field to query
	Value []*string `json:"Value,omitempty" name:"Value"`
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

	// Information of the node whether the instance locates
	// Note: This field may return `null`, indicating that no valid value was found.
	NodeInfo *NodeInfo `json:"NodeInfo,omitempty" name:"NodeInfo"`

	// Start time
	// Note: this field may return `null`, indicating that no valid value can be obtained.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Whether the status is unhealthy or healthy
	// Note: this field may return `null`, indicating that no valid value can be obtained.
	Unhealthy *bool `json:"Unhealthy,omitempty" name:"Unhealthy"`

	// Warning message when the result is unhealthy
	// Note: This field may return `null`, indicating that no valid value was found.
	UnhealthyWarningMsg *string `json:"UnhealthyWarningMsg,omitempty" name:"UnhealthyWarningMsg"`

	// Version ID
	// Note: This field may return `null`, indicating that no valid value was found.
	VersionId *string `json:"VersionId,omitempty" name:"VersionId"`

	// Application name
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`
}

type ServicePage struct {
	// List of applications
	Records []*TemService `json:"Records,omitempty" name:"Records"`

	// Total number of applications
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// Number of applications per page
	Size *int64 `json:"Size,omitempty" name:"Size"`

	// Total number of pages
	Pages *int64 `json:"Pages,omitempty" name:"Pages"`

	// Number of current entries
	// Note: This field may return `null`, indicating that no valid value was found.
	Current *int64 `json:"Current,omitempty" name:"Current"`
}

type ServicePortMapping struct {
	// Specifies how a layer-4 proxy is created.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Type *string `json:"Type,omitempty" name:"Type"`

	// Application name
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`

	// VIP for access within the environment
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	ClusterIp *string `json:"ClusterIp,omitempty" name:"ClusterIp"`

	// Cluster external IP
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	ExternalIp *string `json:"ExternalIp,omitempty" name:"ExternalIp"`

	// Subnet ID
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// VPC ID
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Load balancer ID
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	LoadBalanceId *string `json:"LoadBalanceId,omitempty" name:"LoadBalanceId"`

	// YAML contents
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Yaml *string `json:"Yaml,omitempty" name:"Yaml"`

	// List of exposed ports
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Ports []*int64 `json:"Ports,omitempty" name:"Ports"`

	// Port mapping array 
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	PortMappingItemList []*ServicePortMappingItem `json:"PortMappingItemList,omitempty" name:"PortMappingItemList"`

	// CLB domain name
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	ExternalDomain *string `json:"ExternalDomain,omitempty" name:"ExternalDomain"`
}

type ServicePortMappingItem struct {
	// Application access port
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// Application listening port
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	TargetPort *int64 `json:"TargetPort,omitempty" name:"TargetPort"`

	// Protocol type
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
}

type ServiceVersionBrief struct {
	// Version name
	VersionName *string `json:"VersionName,omitempty" name:"VersionName"`

	// Status of version
	Status *string `json:"Status,omitempty" name:"Status"`

	// (Disused) Whether to enable elastic scaling
	EnableEs *int64 `json:"EnableEs,omitempty" name:"EnableEs"`

	// Number of current instances
	CurrentInstances *int64 `json:"CurrentInstances,omitempty" name:"CurrentInstances"`

	// Version ID
	VersionId *string `json:"VersionId,omitempty" name:"VersionId"`

	// (Disused) Log output configuration
	// Note: This field may return null, indicating that no valid values can be obtained.
	LogOutputConf *LogOutputConf `json:"LogOutputConf,omitempty" name:"LogOutputConf"`

	// Expected number of instances
	// Note: This field may return null, indicating that no valid values can be obtained.
	ExpectedInstances *int64 `json:"ExpectedInstances,omitempty" name:"ExpectedInstances"`

	// Deployment mode
	// Note: This field may return null, indicating that no valid values can be obtained.
	DeployMode *string `json:"DeployMode,omitempty" name:"DeployMode"`

	// Task ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	BuildTaskId *string `json:"BuildTaskId,omitempty" name:"BuildTaskId"`

	// Environment ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// Environment name
	// Note: This field may return null, indicating that no valid values can be obtained.
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// Application ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// Application name
	// Note: This field may return null, indicating that no valid values can be obtained.
	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`

	// Whether the application is being deployed
	// Note: This field may return null, indicating that no valid values can be obtained.
	UnderDeploying *bool `json:"UnderDeploying,omitempty" name:"UnderDeploying"`

	// Status of batch deployment
	// Note: This field may return `null`, indicating that no valid value was found.
	BatchDeployStatus *string `json:"BatchDeployStatus,omitempty" name:"BatchDeployStatus"`

	// Availability zones
	// Note: This field may return `null`, indicating that no valid value was found.
	Zones []*string `json:"Zones,omitempty" name:"Zones"`

	// Node information
	// Note: This field may return `null`, indicating that no valid value was found.
	NodeInfos []*NodeInfo `json:"NodeInfos,omitempty" name:"NodeInfos"`

	// Pod information
	// Note: This field may return `null`, indicating that no valid value was found.
	PodList *DescribeRunPodPage `json:"PodList,omitempty" name:"PodList"`

	// Workload information
	// Note: This field may return `null`, indicating that no valid value was found.
	WorkloadInfo *WorkloadInfo `json:"WorkloadInfo,omitempty" name:"WorkloadInfo"`

	// Creation time
	// Note: This field may return `null`, indicating that no valid value was found.
	CreateDate *string `json:"CreateDate,omitempty" name:"CreateDate"`

	// Region ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
}

type SortType struct {
	// Name of the sorting field
	Key *string `json:"Key,omitempty" name:"Key"`

	// `0`: Ascending; `1`: Descending 
	Type *int64 `json:"Type,omitempty" name:"Type"`
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

type Tag struct {
	// The tag key.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// The tag value.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type TemEnvironmentStartingStatus struct {
	// Number of applications to start
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	ApplicationNumNeedToStart *int64 `json:"ApplicationNumNeedToStart,omitempty" name:"ApplicationNumNeedToStart"`

	// Number of started applictions
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	StartedApplicationNum *int64 `json:"StartedApplicationNum,omitempty" name:"StartedApplicationNum"`

	// Number of applications failed to be started
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	StartFailedApplicationNum *int64 `json:"StartFailedApplicationNum,omitempty" name:"StartFailedApplicationNum"`
}

type TemEnvironmentStoppingStatus struct {
	// Number of applications to stop
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	ApplicationNumNeedToStop *int64 `json:"ApplicationNumNeedToStop,omitempty" name:"ApplicationNumNeedToStop"`

	// Number of stopped applications
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	StoppedApplicationNum *int64 `json:"StoppedApplicationNum,omitempty" name:"StoppedApplicationNum"`

	// Number of applications failed to be stopped
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	StopFailedApplicationNum *int64 `json:"StopFailedApplicationNum,omitempty" name:"StopFailedApplicationNum"`
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

	// User AppId
	// Note: This field may return null, indicating that no valid values can be obtained.
	AppId *string `json:"AppId,omitempty" name:"AppId"`

	// User UIN
	// Note: This field may return null, indicating that no valid values can be obtained.
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// The UIN of sub-account
	// Note: This field may return null, indicating that no valid values can be obtained.
	SubAccountUin *string `json:"SubAccountUin,omitempty" name:"SubAccountUin"`

	// Application ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Tag.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// Whether it’s authorized to access the resource
	// Note: This field may return null, indicating that no valid values can be obtained.
	HasAuthority *bool `json:"HasAuthority,omitempty" name:"HasAuthority"`

	// Environment type. Values: `test`, `pre`, `prod`
	// Note: This field may return `null`, indicating that no valid value was found.
	EnvType *string `json:"EnvType,omitempty" name:"EnvType"`

	// Region code
	// Note: This field may return `null`, indicating that no valid value was found.
	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
}

type TemService struct {
	// Version ID
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// Application name
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`

	// Description
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Description *string `json:"Description,omitempty" name:"Description"`

	// ID of the environment
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// Creation time.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	CreateDate *string `json:"CreateDate,omitempty" name:"CreateDate"`

	// Modification time
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	ModifyDate *string `json:"ModifyDate,omitempty" name:"ModifyDate"`

	// Modifier
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Modifier *string `json:"Modifier,omitempty" name:"Modifier"`

	// Creator account
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Creator *string `json:"Creator,omitempty" name:"Creator"`

	// TCR Individual or TCR Enterprise
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	RepoType *int64 `json:"RepoType,omitempty" name:"RepoType"`

	// ID of the TCR Enterprise instance
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Name of the TCR instance
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`

	// Programming Language
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	CodingLanguage *string `json:"CodingLanguage,omitempty" name:"CodingLanguage"`

	// Deployment mode
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	DeployMode *string `json:"DeployMode,omitempty" name:"DeployMode"`

	// Environment name
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// The instance information where the application is running
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	ActiveVersions []*ServiceVersionBrief `json:"ActiveVersions,omitempty" name:"ActiveVersions"`

	// Whether to enable link tracing
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	EnableTracing *uint64 `json:"EnableTracing,omitempty" name:"EnableTracing"`

	// Tag
	// Note: This field may return `null`, indicating that no valid value was found.
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// Whether it’s authorized to access the resource
	// Note: This field may return null, indicating that no valid values can be obtained.
	HasAuthority *bool `json:"HasAuthority,omitempty" name:"HasAuthority"`
}

type TemServiceVersionInfo struct {
	// Version ID
	VersionId *string `json:"VersionId,omitempty" name:"VersionId"`

	// Application ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// Deployment mode
	DeployMode *string `json:"DeployMode,omitempty" name:"DeployMode"`

	// JDK version
	JdkVersion *string `json:"JdkVersion,omitempty" name:"JdkVersion"`

	// Description
	Description *string `json:"Description,omitempty" name:"Description"`

	// Deployed version
	DeployVersion *string `json:"DeployVersion,omitempty" name:"DeployVersion"`

	// Publish mode
	PublishMode *string `json:"PublishMode,omitempty" name:"PublishMode"`

	// Launch parameter
	JvmOpts *string `json:"JvmOpts,omitempty" name:"JvmOpts"`

	// Number of initial pods
	InitPodNum *int64 `json:"InitPodNum,omitempty" name:"InitPodNum"`

	// CPU specification
	CpuSpec *float64 `json:"CpuSpec,omitempty" name:"CpuSpec"`

	// Memory specification
	MemorySpec *float64 `json:"MemorySpec,omitempty" name:"MemorySpec"`

	// Image path
	ImgRepo *string `json:"ImgRepo,omitempty" name:"ImgRepo"`

	// Image name
	ImgName *string `json:"ImgName,omitempty" name:"ImgName"`

	// Image version
	ImgVersion *string `json:"ImgVersion,omitempty" name:"ImgVersion"`

	// Scaling configuration
	// Note: This field may return null, indicating that no valid values can be obtained.
	EsInfo *EsInfo `json:"EsInfo,omitempty" name:"EsInfo"`

	// Environment configuration
	EnvConf []*Pair `json:"EnvConf,omitempty" name:"EnvConf"`

	// Storage configuration
	StorageConfs []*StorageConf `json:"StorageConfs,omitempty" name:"StorageConfs"`

	// Running status
	Status *string `json:"Status,omitempty" name:"Status"`

	// VPC
	Vpc *string `json:"Vpc,omitempty" name:"Vpc"`

	// Subnets
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// Creation time
	CreateDate *string `json:"CreateDate,omitempty" name:"CreateDate"`

	// Modification time
	ModifyDate *string `json:"ModifyDate,omitempty" name:"ModifyDate"`

	// Mounting configuration
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	StorageMountConfs []*StorageMountConf `json:"StorageMountConfs,omitempty" name:"StorageMountConfs"`

	// Version name
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	VersionName *string `json:"VersionName,omitempty" name:"VersionName"`

	// Log output configuration
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	LogOutputConf *LogOutputConf `json:"LogOutputConf,omitempty" name:"LogOutputConf"`

	// Application name
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`

	// Application description
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	ApplicationDescription *string `json:"ApplicationDescription,omitempty" name:"ApplicationDescription"`

	// Environment name
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// Environment ID
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// Public network address
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	PublicDomain *string `json:"PublicDomain,omitempty" name:"PublicDomain"`

	// Whether to enable public network access
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	EnablePublicAccess *bool `json:"EnablePublicAccess,omitempty" name:"EnablePublicAccess"`

	// Number of current instances
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	CurrentInstances *int64 `json:"CurrentInstances,omitempty" name:"CurrentInstances"`

	// Number of expected instances
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	ExpectedInstances *int64 `json:"ExpectedInstances,omitempty" name:"ExpectedInstances"`

	// Programming Language
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	CodingLanguage *string `json:"CodingLanguage,omitempty" name:"CodingLanguage"`

	// Program package name
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	PkgName *string `json:"PkgName,omitempty" name:"PkgName"`

	// Whether to enable auto scaling
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	EsEnable *int64 `json:"EsEnable,omitempty" name:"EsEnable"`

	// Auto scaling policy
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	EsStrategy *int64 `json:"EsStrategy,omitempty" name:"EsStrategy"`

	// Image tag
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	ImageTag *string `json:"ImageTag,omitempty" name:"ImageTag"`

	// Whether to enable logging
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	LogEnable *int64 `json:"LogEnable,omitempty" name:"LogEnable"`

	// Minimum number of instances
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	MinAliveInstances *string `json:"MinAliveInstances,omitempty" name:"MinAliveInstances"`

	// Security group IDs
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`

	// Image command
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	ImageCommand *string `json:"ImageCommand,omitempty" name:"ImageCommand"`

	// Image command parameters
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	ImageArgs []*string `json:"ImageArgs,omitempty" name:"ImageArgs"`

	// Whether to use the default registry configurations
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	UseRegistryDefaultConfig *bool `json:"UseRegistryDefaultConfig,omitempty" name:"UseRegistryDefaultConfig"`

	// EKS access configuration
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Service *EksService `json:"Service,omitempty" name:"Service"`

	// Mounting configurations
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	SettingConfs []*MountedSettingConf `json:"SettingConfs,omitempty" name:"SettingConfs"`

	// Log path information
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	LogConfs []*string `json:"LogConfs,omitempty" name:"LogConfs"`

	// The script to execute right after the startup
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	PostStart *string `json:"PostStart,omitempty" name:"PostStart"`

	// The script to run before stop
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	PreStop *string `json:"PreStop,omitempty" name:"PreStop"`

	// Configuration of aliveness probe
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Liveness *HealthCheckConfig `json:"Liveness,omitempty" name:"Liveness"`

	// Configuration of readiness probe
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Readiness *HealthCheckConfig `json:"Readiness,omitempty" name:"Readiness"`

	// Auto scaling policy
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	HorizontalAutoscaler []*HorizontalAutoscaler `json:"HorizontalAutoscaler,omitempty" name:"HorizontalAutoscaler"`

	// Scheduled auto scaling policy
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	CronHorizontalAutoscaler []*CronHorizontalAutoscaler `json:"CronHorizontalAutoscaler,omitempty" name:"CronHorizontalAutoscaler"`

	// Availability zone of the application
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Zones []*string `json:"Zones,omitempty" name:"Zones"`

	// The latest deployment time
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	LastDeployDate *string `json:"LastDeployDate,omitempty" name:"LastDeployDate"`

	// The latest successful deployment time
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	LastDeploySuccessDate *string `json:"LastDeploySuccessDate,omitempty" name:"LastDeploySuccessDate"`

	// Information of the node whether the application is deployed
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	NodeInfos []*NodeInfo `json:"NodeInfos,omitempty" name:"NodeInfos"`

	// Image type. Values: `0` (Demo image), `1` (Normal image)
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	ImageType *int64 `json:"ImageType,omitempty" name:"ImageType"`

	// Whether to 
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	EnableTracing *uint64 `json:"EnableTracing,omitempty" name:"EnableTracing"`

	// (Disused) Whether to enable linkage tracing and report. It only takes effect when EnableTracing = `1`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	EnableTracingReport *uint64 `json:"EnableTracingReport,omitempty" name:"EnableTracingReport"`

	// Image type. `0`: Individual image; `1`: Enterprise image; `2`: Public image
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	RepoType *uint64 `json:"RepoType,omitempty" name:"RepoType"`

	// Status of batch deployment: `batch_updating`, `batch_updating_waiting_confirm`
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	BatchDeployStatus *string `json:"BatchDeployStatus,omitempty" name:"BatchDeployStatus"`

	// APM instance ID
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	ApmInstanceId *string `json:"ApmInstanceId,omitempty" name:"ApmInstanceId"`

	// Workload information 
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	WorkloadInfo *WorkloadInfo `json:"WorkloadInfo,omitempty" name:"WorkloadInfo"`

	// Whether to enable application acceleration
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	SpeedUp *bool `json:"SpeedUp,omitempty" name:"SpeedUp"`

	// Configuration of the startup probe
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	StartupProbe *HealthCheckConfig `json:"StartupProbe,omitempty" name:"StartupProbe"`

	// OS version. Values:
	// - ALPINE
	// - CENTOS
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	OsFlavour *string `json:"OsFlavour,omitempty" name:"OsFlavour"`

	// Image repository server
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	RepoServer *string `json:"RepoServer,omitempty" name:"RepoServer"`

	// Whether the application is being deployed
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	UnderDeploying *bool `json:"UnderDeploying,omitempty" name:"UnderDeploying"`

	// Whether to enable application metric monitoring 
	// Note: This field may return null, indicating that no valid values can be obtained.
	EnablePrometheusConf *EnablePrometheusConf `json:"EnablePrometheusConf,omitempty" name:"EnablePrometheusConf"`

	// Whether it’s stopped manually
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	StoppedManually *bool `json:"StoppedManually,omitempty" name:"StoppedManually"`

	// TCR instance ID
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	TcrInstanceId *string `json:"TcrInstanceId,omitempty" name:"TcrInstanceId"`

	// `1`: Automatically enable metrics collection (open-telemetry)
	// `0`: Disable metrics collection
	// Note: This field may return null, indicating that no valid values can be obtained.
	EnableMetrics *int64 `json:"EnableMetrics,omitempty" name:"EnableMetrics"`

	// User AppId
	// Note: This field may return `null`, indicating that no valid value was found.
	AppId *string `json:"AppId,omitempty" name:"AppId"`

	// Sub Account UIN
	// Note: This field may return `null`, indicating that no valid value was found.
	SubAccountUin *string `json:"SubAccountUin,omitempty" name:"SubAccountUin"`

	// User UIN
	// Note: This field may return `null`, indicating that no valid value was found.
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// Region
	// Note: This field may return `null`, indicating that no valid value was found.
	Region *string `json:"Region,omitempty" name:"Region"`

	// Application group ID
	// Note: This field may return `null`, indicating that no valid value was found.
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// Whether to enable registry
	// Note: This field may return `null`, indicating that no valid value was found.
	EnableRegistry *int64 `json:"EnableRegistry,omitempty" name:"EnableRegistry"`

	// Array of scaling rules
	// Note: This field may return `null`, indicating that no valid value was found.
	AutoscalerList []*Autoscaler `json:"AutoscalerList,omitempty" name:"AutoscalerList"`

	// Modifier
	// Note: This field may return `null`, indicating that no valid value was found.
	Modifier *string `json:"Modifier,omitempty" name:"Modifier"`

	// Creator
	// Note: This field may return `null`, indicating that no valid value was found.
	Creator *string `json:"Creator,omitempty" name:"Creator"`

	// Deployment strategy
	// Note: This field may return `null`, indicating that no valid value was found.
	DeployStrategyConf *DeployStrategyConf `json:"DeployStrategyConf,omitempty" name:"DeployStrategyConf"`

	// List of pods
	// Note: This field may return `null`, indicating that no valid value was found.
	PodList *DescribeRunPodPage `json:"PodList,omitempty" name:"PodList"`

	// Whether the configuration has been changed during deployment
	// Note: This field may return `null`, indicating that no valid value was found.
	ConfEdited *bool `json:"ConfEdited,omitempty" name:"ConfEdited"`

	// Tag
	// Note: This field may return `null`, indicating that no valid value was found.
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// Whether to encode
	// Note: This field may return null, indicating that no valid values can be obtained.
	PreStopEncoded *string `json:"PreStopEncoded,omitempty" name:"PreStopEncoded"`

	// Whether to encode
	// Note: This field may return null, indicating that no valid values can be obtained.
	PostStartEncoded *string `json:"PostStartEncoded,omitempty" name:"PostStartEncoded"`
}

type UseDefaultRepoParameters struct {
	// TCR Enterprise instance name
	// Note: This field may return null, indicating that no valid values can be obtained.
	EnterpriseInstanceName *string `json:"EnterpriseInstanceName,omitempty" name:"EnterpriseInstanceName"`

	// TCR Enterprise billing mode. `0`: Pay-as-you-go 
	// Note: This field may return null, indicating that no valid values can be obtained.
	EnterpriseInstanceChargeType *int64 `json:"EnterpriseInstanceChargeType,omitempty" name:"EnterpriseInstanceChargeType"`

	// Edition of the TCR Enterprise. Values: `basic`, `standard`, `premium` (Advanced edition)
	// Note: This field may return null, indicating that no valid values can be obtained.
	EnterpriseInstanceType *string `json:"EnterpriseInstanceType,omitempty" name:"EnterpriseInstanceType"`
}

type WorkloadInfo struct {
	// The resource ID.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Application name
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`

	// Version name
	// Note: This field may return `null`, indicating that no valid value was found.
	VersionName *string `json:"VersionName,omitempty" name:"VersionName"`

	// Number of ready replicas
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	ReadyReplicas *int64 `json:"ReadyReplicas,omitempty" name:"ReadyReplicas"`

	// Number of replicas
	// Note: This field may return `null`, indicating that no valid value was found.
	Replicas *int64 `json:"Replicas,omitempty" name:"Replicas"`

	// Number of updated replicas
	// Note: This field may return `null`, indicating that no valid value was found.
	UpdatedReplicas *int64 `json:"UpdatedReplicas,omitempty" name:"UpdatedReplicas"`

	// Number of replicas ready for update
	// Note: This field may return `null`, indicating that no valid value was found.
	UpdatedReadyReplicas *int64 `json:"UpdatedReadyReplicas,omitempty" name:"UpdatedReadyReplicas"`

	// ## Version Updates
	// Note: This field may return `null`, indicating that no valid value was found.
	UpdateRevision *string `json:"UpdateRevision,omitempty" name:"UpdateRevision"`

	// Current Version
	// Note: This field may return `null`, indicating that no valid value was found.
	CurrentRevision *string `json:"CurrentRevision,omitempty" name:"CurrentRevision"`
}