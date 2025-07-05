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

package v20201221

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

type CosToken struct {
	// Unique request ID
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`

	// Bucket name
	Bucket *string `json:"Bucket,omitnil,omitempty" name:"Bucket"`

	// Bucket region
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// `SecretId` of temporary key
	TmpSecretId *string `json:"TmpSecretId,omitnil,omitempty" name:"TmpSecretId"`

	// `SecretKey` of temporary key
	TmpSecretKey *string `json:"TmpSecretKey,omitnil,omitempty" name:"TmpSecretKey"`

	// `sessionToken` of temporary key
	SessionToken *string `json:"SessionToken,omitnil,omitempty" name:"SessionToken"`

	// `StartTime` of temporary key acquisition
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// `ExpiredTime` of temporary key
	ExpiredTime *string `json:"ExpiredTime,omitnil,omitempty" name:"ExpiredTime"`

	// Full package path
	FullPath *string `json:"FullPath,omitnil,omitempty" name:"FullPath"`
}

// Predefined struct for user
type CreateCosTokenV2RequestParams struct {
	// Service ID
	ServiceId *string `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`

	// Package name
	PkgName *string `json:"PkgName,omitnil,omitempty" name:"PkgName"`

	// optType. 1: upload; 2: query
	OptType *int64 `json:"OptType,omitnil,omitempty" name:"OptType"`

	// Source channel
	SourceChannel *int64 `json:"SourceChannel,omitnil,omitempty" name:"SourceChannel"`

	// Input parameter of `deployVersion`
	TimeVersion *string `json:"TimeVersion,omitnil,omitempty" name:"TimeVersion"`
}

type CreateCosTokenV2Request struct {
	*tchttp.BaseRequest
	
	// Service ID
	ServiceId *string `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`

	// Package name
	PkgName *string `json:"PkgName,omitnil,omitempty" name:"PkgName"`

	// optType. 1: upload; 2: query
	OptType *int64 `json:"OptType,omitnil,omitempty" name:"OptType"`

	// Source channel
	SourceChannel *int64 `json:"SourceChannel,omitnil,omitempty" name:"SourceChannel"`

	// Input parameter of `deployVersion`
	TimeVersion *string `json:"TimeVersion,omitnil,omitempty" name:"TimeVersion"`
}

func (r *CreateCosTokenV2Request) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCosTokenV2Request) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "PkgName")
	delete(f, "OptType")
	delete(f, "SourceChannel")
	delete(f, "TimeVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCosTokenV2Request has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCosTokenV2ResponseParams struct {
	// `CosToken` object in case of success and `null` in case of failure
	// Note: this field may return null, indicating that no valid values can be obtained.
	Result *CosToken `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCosTokenV2Response struct {
	*tchttp.BaseResponse
	Response *CreateCosTokenV2ResponseParams `json:"Response"`
}

func (r *CreateCosTokenV2Response) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCosTokenV2Response) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNamespaceRequestParams struct {
	// Namespace name
	NamespaceName *string `json:"NamespaceName,omitnil,omitempty" name:"NamespaceName"`

	// VPC name
	Vpc *string `json:"Vpc,omitnil,omitempty" name:"Vpc"`

	// Subnet list
	SubnetIds []*string `json:"SubnetIds,omitnil,omitempty" name:"SubnetIds"`

	// Namespace description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// K8s version
	K8sVersion *string `json:"K8sVersion,omitnil,omitempty" name:"K8sVersion"`

	// Source channel
	SourceChannel *int64 `json:"SourceChannel,omitnil,omitempty" name:"SourceChannel"`

	// Whether to enable the TSW service.
	EnableTswTraceService *bool `json:"EnableTswTraceService,omitnil,omitempty" name:"EnableTswTraceService"`
}

type CreateNamespaceRequest struct {
	*tchttp.BaseRequest
	
	// Namespace name
	NamespaceName *string `json:"NamespaceName,omitnil,omitempty" name:"NamespaceName"`

	// VPC name
	Vpc *string `json:"Vpc,omitnil,omitempty" name:"Vpc"`

	// Subnet list
	SubnetIds []*string `json:"SubnetIds,omitnil,omitempty" name:"SubnetIds"`

	// Namespace description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// K8s version
	K8sVersion *string `json:"K8sVersion,omitnil,omitempty" name:"K8sVersion"`

	// Source channel
	SourceChannel *int64 `json:"SourceChannel,omitnil,omitempty" name:"SourceChannel"`

	// Whether to enable the TSW service.
	EnableTswTraceService *bool `json:"EnableTswTraceService,omitnil,omitempty" name:"EnableTswTraceService"`
}

func (r *CreateNamespaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNamespaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NamespaceName")
	delete(f, "Vpc")
	delete(f, "SubnetIds")
	delete(f, "Description")
	delete(f, "K8sVersion")
	delete(f, "SourceChannel")
	delete(f, "EnableTswTraceService")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateNamespaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNamespaceResponseParams struct {
	// Namespace ID in case of success and `null` in case of failure
	// Note: this field may return null, indicating that no valid values can be obtained.
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateNamespaceResponse struct {
	*tchttp.BaseResponse
	Response *CreateNamespaceResponseParams `json:"Response"`
}

func (r *CreateNamespaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNamespaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateResourceRequestParams struct {
	// Namespace ID
	NamespaceId *string `json:"NamespaceId,omitnil,omitempty" name:"NamespaceId"`

	// Resource type. Valid values: CFS (file system), CLS (log service), TSE_SRE (registry)
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// Resource ID
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// Source channel
	SourceChannel *int64 `json:"SourceChannel,omitnil,omitempty" name:"SourceChannel"`
}

type CreateResourceRequest struct {
	*tchttp.BaseRequest
	
	// Namespace ID
	NamespaceId *string `json:"NamespaceId,omitnil,omitempty" name:"NamespaceId"`

	// Resource type. Valid values: CFS (file system), CLS (log service), TSE_SRE (registry)
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// Resource ID
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// Source channel
	SourceChannel *int64 `json:"SourceChannel,omitnil,omitempty" name:"SourceChannel"`
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
	delete(f, "NamespaceId")
	delete(f, "ResourceType")
	delete(f, "ResourceId")
	delete(f, "SourceChannel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateResourceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateResourceResponseParams struct {
	// Success or failure
	// Note: this field may return null, indicating that no valid values can be obtained.
	Result *bool `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

// Predefined struct for user
type CreateServiceV2RequestParams struct {
	// Service name
	ServiceName *string `json:"ServiceName,omitnil,omitempty" name:"ServiceName"`

	// Description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Whether to use the default image service. 1: yes; 0: no
	UseDefaultImageService *int64 `json:"UseDefaultImageService,omitnil,omitempty" name:"UseDefaultImageService"`

	// Type of the bound repository. 0: Personal Edition; 1: Enterprise Edition
	RepoType *int64 `json:"RepoType,omitnil,omitempty" name:"RepoType"`

	// Instance ID of Enterprise Edition image service
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Address of the bound image server
	RepoServer *string `json:"RepoServer,omitnil,omitempty" name:"RepoServer"`

	// Name of the bound image repository
	RepoName *string `json:"RepoName,omitnil,omitempty" name:"RepoName"`

	// Source channel
	SourceChannel *int64 `json:"SourceChannel,omitnil,omitempty" name:"SourceChannel"`

	// Service subnet
	SubnetList []*string `json:"SubnetList,omitnil,omitempty" name:"SubnetList"`

	// Programming language 
	// - JAVA
	// - OTHER
	CodingLanguage *string `json:"CodingLanguage,omitnil,omitempty" name:"CodingLanguage"`

	// Deployment mode 
	// - IMAGE
	// - JAR
	// - WAR
	DeployMode *string `json:"DeployMode,omitnil,omitempty" name:"DeployMode"`
}

type CreateServiceV2Request struct {
	*tchttp.BaseRequest
	
	// Service name
	ServiceName *string `json:"ServiceName,omitnil,omitempty" name:"ServiceName"`

	// Description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Whether to use the default image service. 1: yes; 0: no
	UseDefaultImageService *int64 `json:"UseDefaultImageService,omitnil,omitempty" name:"UseDefaultImageService"`

	// Type of the bound repository. 0: Personal Edition; 1: Enterprise Edition
	RepoType *int64 `json:"RepoType,omitnil,omitempty" name:"RepoType"`

	// Instance ID of Enterprise Edition image service
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Address of the bound image server
	RepoServer *string `json:"RepoServer,omitnil,omitempty" name:"RepoServer"`

	// Name of the bound image repository
	RepoName *string `json:"RepoName,omitnil,omitempty" name:"RepoName"`

	// Source channel
	SourceChannel *int64 `json:"SourceChannel,omitnil,omitempty" name:"SourceChannel"`

	// Service subnet
	SubnetList []*string `json:"SubnetList,omitnil,omitempty" name:"SubnetList"`

	// Programming language 
	// - JAVA
	// - OTHER
	CodingLanguage *string `json:"CodingLanguage,omitnil,omitempty" name:"CodingLanguage"`

	// Deployment mode 
	// - IMAGE
	// - JAR
	// - WAR
	DeployMode *string `json:"DeployMode,omitnil,omitempty" name:"DeployMode"`
}

func (r *CreateServiceV2Request) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateServiceV2Request) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceName")
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateServiceV2Request has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateServiceV2ResponseParams struct {
	// Service code
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateServiceV2Response struct {
	*tchttp.BaseResponse
	Response *CreateServiceV2ResponseParams `json:"Response"`
}

func (r *CreateServiceV2Response) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateServiceV2Response) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteIngressRequestParams struct {
	// tem NamespaceId
	NamespaceId *string `json:"NamespaceId,omitnil,omitempty" name:"NamespaceId"`

	// EKS namespace name
	EksNamespace *string `json:"EksNamespace,omitnil,omitempty" name:"EksNamespace"`

	// Ingress rule name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Source channel
	SourceChannel *int64 `json:"SourceChannel,omitnil,omitempty" name:"SourceChannel"`
}

type DeleteIngressRequest struct {
	*tchttp.BaseRequest
	
	// tem NamespaceId
	NamespaceId *string `json:"NamespaceId,omitnil,omitempty" name:"NamespaceId"`

	// EKS namespace name
	EksNamespace *string `json:"EksNamespace,omitnil,omitempty" name:"EksNamespace"`

	// Ingress rule name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Source channel
	SourceChannel *int64 `json:"SourceChannel,omitnil,omitempty" name:"SourceChannel"`
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
	delete(f, "NamespaceId")
	delete(f, "EksNamespace")
	delete(f, "Name")
	delete(f, "SourceChannel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteIngressRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteIngressResponseParams struct {
	// Whether deletion succeeded
	Result *bool `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DeployServiceV2RequestParams struct {
	// Service ID
	ServiceId *string `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`

	// Container port
	ContainerPort *uint64 `json:"ContainerPort,omitnil,omitempty" name:"ContainerPort"`

	// Number of initialized pods
	InitPodNum *uint64 `json:"InitPodNum,omitnil,omitempty" name:"InitPodNum"`

	// CPU specification
	CpuSpec *float64 `json:"CpuSpec,omitnil,omitempty" name:"CpuSpec"`

	// Memory specification
	MemorySpec *float64 `json:"MemorySpec,omitnil,omitempty" name:"MemorySpec"`

	// Environment ID
	NamespaceId *string `json:"NamespaceId,omitnil,omitempty" name:"NamespaceId"`

	// Image repository
	ImgRepo *string `json:"ImgRepo,omitnil,omitempty" name:"ImgRepo"`

	// Version description
	VersionDesc *string `json:"VersionDesc,omitnil,omitempty" name:"VersionDesc"`

	// Launch parameters
	JvmOpts *string `json:"JvmOpts,omitnil,omitempty" name:"JvmOpts"`

	// Auto scaling configuration. If this parameter is left empty, auto scaling will not be enabled
	EsInfo *EsInfo `json:"EsInfo,omitnil,omitempty" name:"EsInfo"`

	// Environment variable configuration
	EnvConf []*Pair `json:"EnvConf,omitnil,omitempty" name:"EnvConf"`

	// Log configuration
	LogConfs []*string `json:"LogConfs,omitnil,omitempty" name:"LogConfs"`

	// Data volume configuration
	StorageConfs []*StorageConf `json:"StorageConfs,omitnil,omitempty" name:"StorageConfs"`

	// Data volume mount configuration
	StorageMountConfs []*StorageMountConf `json:"StorageMountConfs,omitnil,omitempty" name:"StorageMountConfs"`

	// Deployment type.
	// - JAR: deployment through JAR package
	// - WAR: deployment through WAR package
	// - IMAGE: deployment through image
	DeployMode *string `json:"DeployMode,omitnil,omitempty" name:"DeployMode"`

	// When the deployment type is `IMAGE`, this parameter indicates the image tag.
	// When the deployment type is `JAR` or `WAR`, this parameter indicates the package version number.
	DeployVersion *string `json:"DeployVersion,omitnil,omitempty" name:"DeployVersion"`

	// Package name, which is required when using JAR or WAR packages for deployment.
	PkgName *string `json:"PkgName,omitnil,omitempty" name:"PkgName"`

	// JDK version.
	// - KONA: use KONA JDK.
	// - OPEN: use open JDK.
	JdkVersion *string `json:"JdkVersion,omitnil,omitempty" name:"JdkVersion"`

	// Security group IDs
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// Log output configuration
	LogOutputConf *LogOutputConf `json:"LogOutputConf,omitnil,omitempty" name:"LogOutputConf"`

	// Source channel
	SourceChannel *int64 `json:"SourceChannel,omitnil,omitempty" name:"SourceChannel"`

	// Version description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Image command
	ImageCommand *string `json:"ImageCommand,omitnil,omitempty" name:"ImageCommand"`

	// Image command parameters
	ImageArgs []*string `json:"ImageArgs,omitnil,omitempty" name:"ImageArgs"`

	// Service port mapping.
	PortMappings []*PortMapping `json:"PortMappings,omitnil,omitempty" name:"PortMappings"`

	// Whether to add the registry’s default configurations.
	UseRegistryDefaultConfig *bool `json:"UseRegistryDefaultConfig,omitnil,omitempty" name:"UseRegistryDefaultConfig"`

	// Mounting configurations
	SettingConfs []*MountedSettingConf `json:"SettingConfs,omitnil,omitempty" name:"SettingConfs"`

	// EKS access configuration
	EksService *EksService `json:"EksService,omitnil,omitempty" name:"EksService"`

	// ID of the version that you want to roll back to
	VersionId *string `json:"VersionId,omitnil,omitempty" name:"VersionId"`

	// The script to run after startup
	PostStart *string `json:"PostStart,omitnil,omitempty" name:"PostStart"`

	// The script to run before stop
	PreStop *string `json:"PreStop,omitnil,omitempty" name:"PreStop"`

	// Configuration of 
	DeployStrategyConf *DeployStrategyConf `json:"DeployStrategyConf,omitnil,omitempty" name:"DeployStrategyConf"`

	// Configuration of aliveness probe
	Liveness *HealthCheckConfig `json:"Liveness,omitnil,omitempty" name:"Liveness"`

	// Configuration of readiness probe
	Readiness *HealthCheckConfig `json:"Readiness,omitnil,omitempty" name:"Readiness"`
}

type DeployServiceV2Request struct {
	*tchttp.BaseRequest
	
	// Service ID
	ServiceId *string `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`

	// Container port
	ContainerPort *uint64 `json:"ContainerPort,omitnil,omitempty" name:"ContainerPort"`

	// Number of initialized pods
	InitPodNum *uint64 `json:"InitPodNum,omitnil,omitempty" name:"InitPodNum"`

	// CPU specification
	CpuSpec *float64 `json:"CpuSpec,omitnil,omitempty" name:"CpuSpec"`

	// Memory specification
	MemorySpec *float64 `json:"MemorySpec,omitnil,omitempty" name:"MemorySpec"`

	// Environment ID
	NamespaceId *string `json:"NamespaceId,omitnil,omitempty" name:"NamespaceId"`

	// Image repository
	ImgRepo *string `json:"ImgRepo,omitnil,omitempty" name:"ImgRepo"`

	// Version description
	VersionDesc *string `json:"VersionDesc,omitnil,omitempty" name:"VersionDesc"`

	// Launch parameters
	JvmOpts *string `json:"JvmOpts,omitnil,omitempty" name:"JvmOpts"`

	// Auto scaling configuration. If this parameter is left empty, auto scaling will not be enabled
	EsInfo *EsInfo `json:"EsInfo,omitnil,omitempty" name:"EsInfo"`

	// Environment variable configuration
	EnvConf []*Pair `json:"EnvConf,omitnil,omitempty" name:"EnvConf"`

	// Log configuration
	LogConfs []*string `json:"LogConfs,omitnil,omitempty" name:"LogConfs"`

	// Data volume configuration
	StorageConfs []*StorageConf `json:"StorageConfs,omitnil,omitempty" name:"StorageConfs"`

	// Data volume mount configuration
	StorageMountConfs []*StorageMountConf `json:"StorageMountConfs,omitnil,omitempty" name:"StorageMountConfs"`

	// Deployment type.
	// - JAR: deployment through JAR package
	// - WAR: deployment through WAR package
	// - IMAGE: deployment through image
	DeployMode *string `json:"DeployMode,omitnil,omitempty" name:"DeployMode"`

	// When the deployment type is `IMAGE`, this parameter indicates the image tag.
	// When the deployment type is `JAR` or `WAR`, this parameter indicates the package version number.
	DeployVersion *string `json:"DeployVersion,omitnil,omitempty" name:"DeployVersion"`

	// Package name, which is required when using JAR or WAR packages for deployment.
	PkgName *string `json:"PkgName,omitnil,omitempty" name:"PkgName"`

	// JDK version.
	// - KONA: use KONA JDK.
	// - OPEN: use open JDK.
	JdkVersion *string `json:"JdkVersion,omitnil,omitempty" name:"JdkVersion"`

	// Security group IDs
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// Log output configuration
	LogOutputConf *LogOutputConf `json:"LogOutputConf,omitnil,omitempty" name:"LogOutputConf"`

	// Source channel
	SourceChannel *int64 `json:"SourceChannel,omitnil,omitempty" name:"SourceChannel"`

	// Version description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Image command
	ImageCommand *string `json:"ImageCommand,omitnil,omitempty" name:"ImageCommand"`

	// Image command parameters
	ImageArgs []*string `json:"ImageArgs,omitnil,omitempty" name:"ImageArgs"`

	// Service port mapping.
	PortMappings []*PortMapping `json:"PortMappings,omitnil,omitempty" name:"PortMappings"`

	// Whether to add the registry’s default configurations.
	UseRegistryDefaultConfig *bool `json:"UseRegistryDefaultConfig,omitnil,omitempty" name:"UseRegistryDefaultConfig"`

	// Mounting configurations
	SettingConfs []*MountedSettingConf `json:"SettingConfs,omitnil,omitempty" name:"SettingConfs"`

	// EKS access configuration
	EksService *EksService `json:"EksService,omitnil,omitempty" name:"EksService"`

	// ID of the version that you want to roll back to
	VersionId *string `json:"VersionId,omitnil,omitempty" name:"VersionId"`

	// The script to run after startup
	PostStart *string `json:"PostStart,omitnil,omitempty" name:"PostStart"`

	// The script to run before stop
	PreStop *string `json:"PreStop,omitnil,omitempty" name:"PreStop"`

	// Configuration of 
	DeployStrategyConf *DeployStrategyConf `json:"DeployStrategyConf,omitnil,omitempty" name:"DeployStrategyConf"`

	// Configuration of aliveness probe
	Liveness *HealthCheckConfig `json:"Liveness,omitnil,omitempty" name:"Liveness"`

	// Configuration of readiness probe
	Readiness *HealthCheckConfig `json:"Readiness,omitnil,omitempty" name:"Readiness"`
}

func (r *DeployServiceV2Request) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeployServiceV2Request) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "ContainerPort")
	delete(f, "InitPodNum")
	delete(f, "CpuSpec")
	delete(f, "MemorySpec")
	delete(f, "NamespaceId")
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
	delete(f, "PortMappings")
	delete(f, "UseRegistryDefaultConfig")
	delete(f, "SettingConfs")
	delete(f, "EksService")
	delete(f, "VersionId")
	delete(f, "PostStart")
	delete(f, "PreStop")
	delete(f, "DeployStrategyConf")
	delete(f, "Liveness")
	delete(f, "Readiness")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeployServiceV2Request has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeployServiceV2ResponseParams struct {
	// Version ID (which can be ignored for the frontend)
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeployServiceV2Response struct {
	*tchttp.BaseResponse
	Response *DeployServiceV2ResponseParams `json:"Response"`
}

func (r *DeployServiceV2Response) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeployServiceV2Response) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeployStrategyConf struct {
	// Total batches
	TotalBatchCount *int64 `json:"TotalBatchCount,omitnil,omitempty" name:"TotalBatchCount"`

	// Number of instances for the beta batch
	BetaBatchNum *int64 `json:"BetaBatchNum,omitnil,omitempty" name:"BetaBatchNum"`

	// Batch deploy policy. `0`: automatically; `1`: manually. If you use beta batch, the policy for beta batch must be `0`. The policy specified here only applies to the rest batches.
	DeployStrategyType *int64 `json:"DeployStrategyType,omitnil,omitempty" name:"DeployStrategyType"`

	// Interval between batches
	BatchInterval *int64 `json:"BatchInterval,omitnil,omitempty" name:"BatchInterval"`
}

// Predefined struct for user
type DescribeIngressRequestParams struct {
	// tem namespaceId
	NamespaceId *string `json:"NamespaceId,omitnil,omitempty" name:"NamespaceId"`

	// EKS namespace name
	EksNamespace *string `json:"EksNamespace,omitnil,omitempty" name:"EksNamespace"`

	// Ingress rule name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Source channel
	SourceChannel *int64 `json:"SourceChannel,omitnil,omitempty" name:"SourceChannel"`
}

type DescribeIngressRequest struct {
	*tchttp.BaseRequest
	
	// tem namespaceId
	NamespaceId *string `json:"NamespaceId,omitnil,omitempty" name:"NamespaceId"`

	// EKS namespace name
	EksNamespace *string `json:"EksNamespace,omitnil,omitempty" name:"EksNamespace"`

	// Ingress rule name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Source channel
	SourceChannel *int64 `json:"SourceChannel,omitnil,omitempty" name:"SourceChannel"`
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
	delete(f, "NamespaceId")
	delete(f, "EksNamespace")
	delete(f, "Name")
	delete(f, "SourceChannel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIngressRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIngressResponseParams struct {
	// Ingress rule configuration
	Result *IngressInfo `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// namespace id
	NamespaceId *string `json:"NamespaceId,omitnil,omitempty" name:"NamespaceId"`

	// namespace
	EksNamespace *string `json:"EksNamespace,omitnil,omitempty" name:"EksNamespace"`

	// Source channel
	SourceChannel *int64 `json:"SourceChannel,omitnil,omitempty" name:"SourceChannel"`

	// Ingress rule name list.
	Names []*string `json:"Names,omitnil,omitempty" name:"Names"`
}

type DescribeIngressesRequest struct {
	*tchttp.BaseRequest
	
	// namespace id
	NamespaceId *string `json:"NamespaceId,omitnil,omitempty" name:"NamespaceId"`

	// namespace
	EksNamespace *string `json:"EksNamespace,omitnil,omitempty" name:"EksNamespace"`

	// Source channel
	SourceChannel *int64 `json:"SourceChannel,omitnil,omitempty" name:"SourceChannel"`

	// Ingress rule name list.
	Names []*string `json:"Names,omitnil,omitempty" name:"Names"`
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
	delete(f, "NamespaceId")
	delete(f, "EksNamespace")
	delete(f, "SourceChannel")
	delete(f, "Names")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIngressesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIngressesResponseParams struct {
	// Ingress array
	// Note: this field may return null, indicating that no valid values can be obtained.
	Result []*IngressInfo `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeNamespacesRequestParams struct {
	// Number of items per page
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Page number
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Source
	SourceChannel *int64 `json:"SourceChannel,omitnil,omitempty" name:"SourceChannel"`
}

type DescribeNamespacesRequest struct {
	*tchttp.BaseRequest
	
	// Number of items per page
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Page number
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Source
	SourceChannel *int64 `json:"SourceChannel,omitnil,omitempty" name:"SourceChannel"`
}

func (r *DescribeNamespacesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNamespacesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "SourceChannel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNamespacesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNamespacesResponseParams struct {
	// Returned result
	Result *NamespacePage `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeNamespacesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNamespacesResponseParams `json:"Response"`
}

func (r *DescribeNamespacesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNamespacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRelatedIngressesRequestParams struct {
	// Environment ID.
	NamespaceId *string `json:"NamespaceId,omitnil,omitempty" name:"NamespaceId"`

	// EKS namespace.
	EksNamespace *string `json:"EksNamespace,omitnil,omitempty" name:"EksNamespace"`

	// Source channel.
	SourceChannel *int64 `json:"SourceChannel,omitnil,omitempty" name:"SourceChannel"`

	// Service ID.
	ServiceId *string `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`
}

type DescribeRelatedIngressesRequest struct {
	*tchttp.BaseRequest
	
	// Environment ID.
	NamespaceId *string `json:"NamespaceId,omitnil,omitempty" name:"NamespaceId"`

	// EKS namespace.
	EksNamespace *string `json:"EksNamespace,omitnil,omitempty" name:"EksNamespace"`

	// Source channel.
	SourceChannel *int64 `json:"SourceChannel,omitnil,omitempty" name:"SourceChannel"`

	// Service ID.
	ServiceId *string `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`
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
	delete(f, "NamespaceId")
	delete(f, "EksNamespace")
	delete(f, "SourceChannel")
	delete(f, "ServiceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRelatedIngressesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRelatedIngressesResponseParams struct {
	// Ingress array.
	// Note: this field may return `null`, indicating that no valid value can be obtained.
	Result []*IngressInfo `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// Page number
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of items per page
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Total number
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Request ID
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`

	// Number of items
	PodList []*RunVersionPod `json:"PodList,omitnil,omitempty" name:"PodList"`
}

// Predefined struct for user
type DescribeServiceRunPodListV2RequestParams struct {
	// Environment ID
	NamespaceId *string `json:"NamespaceId,omitnil,omitempty" name:"NamespaceId"`

	// Service name ID
	ServiceId *string `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`

	// Number of items per page. Default value: 20
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Page number. Default value: 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Pod status 
	// - Running 
	// - Pending 
	// - Error
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Pod name
	PodName *string `json:"PodName,omitnil,omitempty" name:"PodName"`

	// Source channel
	SourceChannel *int64 `json:"SourceChannel,omitnil,omitempty" name:"SourceChannel"`
}

type DescribeServiceRunPodListV2Request struct {
	*tchttp.BaseRequest
	
	// Environment ID
	NamespaceId *string `json:"NamespaceId,omitnil,omitempty" name:"NamespaceId"`

	// Service name ID
	ServiceId *string `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`

	// Number of items per page. Default value: 20
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Page number. Default value: 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Pod status 
	// - Running 
	// - Pending 
	// - Error
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Pod name
	PodName *string `json:"PodName,omitnil,omitempty" name:"PodName"`

	// Source channel
	SourceChannel *int64 `json:"SourceChannel,omitnil,omitempty" name:"SourceChannel"`
}

func (r *DescribeServiceRunPodListV2Request) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServiceRunPodListV2Request) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NamespaceId")
	delete(f, "ServiceId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Status")
	delete(f, "PodName")
	delete(f, "SourceChannel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeServiceRunPodListV2Request has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeServiceRunPodListV2ResponseParams struct {
	// Returned result
	Result *DescribeRunPodPage `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeServiceRunPodListV2Response struct {
	*tchttp.BaseResponse
	Response *DescribeServiceRunPodListV2ResponseParams `json:"Response"`
}

func (r *DescribeServiceRunPodListV2Response) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServiceRunPodListV2Response) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EksService struct {
	// Service name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Available ports
	Ports []*int64 `json:"Ports,omitnil,omitempty" name:"Ports"`

	// Yaml contents
	Yaml *string `json:"Yaml,omitnil,omitempty" name:"Yaml"`

	// Service name
	// Note: this field may return `null`, indicating that no valid value can be obtained.
	ServiceName *string `json:"ServiceName,omitnil,omitempty" name:"ServiceName"`

	// Version name
	// Note: this field may return `null`, indicating that no valid value can be obtained.
	VersionName *string `json:"VersionName,omitnil,omitempty" name:"VersionName"`

	// Private IP
	// Note: this field may return `null`, indicating that no valid value can be obtained.
	ClusterIp []*string `json:"ClusterIp,omitnil,omitempty" name:"ClusterIp"`

	// Public IP
	// Note: this field may return `null`, indicating that no valid value can be obtained.
	ExternalIp *string `json:"ExternalIp,omitnil,omitempty" name:"ExternalIp"`

	// The access type. Valid values:
	// - EXTERNAL (internet access)
	// - VPC（Intra-VPC access)
	// - CLUSTER (Intra-cluster access)
	// Note: this field may return `null`, indicating that no valid value can be obtained.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Subnet ID. It is filled when the access type is `VPC`.
	// Note: this field may return `null`, indicating that no valid value is obtained.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Load balancer ID. It is filled when the access type is `EXTERNAL` or `CLUSTER`. It’s created automatically by default.
	// Note: this field may return `null`, indicating that no valid value is obtained.
	LoadBalanceId *string `json:"LoadBalanceId,omitnil,omitempty" name:"LoadBalanceId"`

	// Port Mapping
	// Note: this field may return `null`, indicating that no valid value can be obtained.
	PortMappings []*PortMapping `json:"PortMappings,omitnil,omitempty" name:"PortMappings"`
}

type EsInfo struct {
	// Minimum number of instances
	MinAliveInstances *int64 `json:"MinAliveInstances,omitnil,omitempty" name:"MinAliveInstances"`

	// Maximum number of instances
	MaxAliveInstances *int64 `json:"MaxAliveInstances,omitnil,omitempty" name:"MaxAliveInstances"`

	// Auto scaling policy. 1: CPU; 2: memory
	EsStrategy *int64 `json:"EsStrategy,omitnil,omitempty" name:"EsStrategy"`

	// Auto scaling condition value
	Threshold *uint64 `json:"Threshold,omitnil,omitempty" name:"Threshold"`

	// Version ID
	VersionId *string `json:"VersionId,omitnil,omitempty" name:"VersionId"`
}

// Predefined struct for user
type GenerateDownloadUrlRequestParams struct {
	// Service ID
	ServiceId *string `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`

	// Package Name
	PkgName *string `json:"PkgName,omitnil,omitempty" name:"PkgName"`

	// Version of the package to download
	DeployVersion *string `json:"DeployVersion,omitnil,omitempty" name:"DeployVersion"`

	// Source channel
	SourceChannel *int64 `json:"SourceChannel,omitnil,omitempty" name:"SourceChannel"`
}

type GenerateDownloadUrlRequest struct {
	*tchttp.BaseRequest
	
	// Service ID
	ServiceId *string `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`

	// Package Name
	PkgName *string `json:"PkgName,omitnil,omitempty" name:"PkgName"`

	// Version of the package to download
	DeployVersion *string `json:"DeployVersion,omitnil,omitempty" name:"DeployVersion"`

	// Source channel
	SourceChannel *int64 `json:"SourceChannel,omitnil,omitempty" name:"SourceChannel"`
}

func (r *GenerateDownloadUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GenerateDownloadUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "PkgName")
	delete(f, "DeployVersion")
	delete(f, "SourceChannel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GenerateDownloadUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GenerateDownloadUrlResponseParams struct {
	// Temp download URL for the package
	// Note: this field may return `null`, indicating that no valid value can be obtained.
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GenerateDownloadUrlResponse struct {
	*tchttp.BaseResponse
	Response *GenerateDownloadUrlResponseParams `json:"Response"`
}

func (r *GenerateDownloadUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GenerateDownloadUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type HealthCheckConfig struct {
	// Health check type. Valid values: `HttpGet`，`TcpSocket`，`Exec`
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// The protocol type. It’s only valid when the health check type is `HttpGet`.
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// The request path. It’s only valid when the health check type is `HttpGet`.
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// The script to be executed. It’s only valid when the health check type is `Exec`.
	Exec *string `json:"Exec,omitnil,omitempty" name:"Exec"`

	// The request port. It’s only valid when the health check type is `HttpGet` or `TcpSocket `.
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// The initial delay for health check in seconds. Default: `0`
	InitialDelaySeconds *int64 `json:"InitialDelaySeconds,omitnil,omitempty" name:"InitialDelaySeconds"`

	// Timeout period in seconds. Default: `1`
	TimeoutSeconds *int64 `json:"TimeoutSeconds,omitnil,omitempty" name:"TimeoutSeconds"`

	// Interval period in seconds. Default: `10`
	PeriodSeconds *int64 `json:"PeriodSeconds,omitnil,omitempty" name:"PeriodSeconds"`
}

type IngressInfo struct {
	// tem namespaceId
	// Note: this field may return null, indicating that no valid values can be obtained.
	NamespaceId *string `json:"NamespaceId,omitnil,omitempty" name:"NamespaceId"`

	// eks namespace
	EksNamespace *string `json:"EksNamespace,omitnil,omitempty" name:"EksNamespace"`

	// ip version
	AddressIPVersion *string `json:"AddressIPVersion,omitnil,omitempty" name:"AddressIPVersion"`

	// ingress name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Rules configuration
	Rules []*IngressRule `json:"Rules,omitnil,omitempty" name:"Rules"`

	// clb ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	ClbId *string `json:"ClbId,omitnil,omitempty" name:"ClbId"`

	// TLS configuration
	// Note: this field may return null, indicating that no valid values can be obtained.
	Tls []*IngressTls `json:"Tls,omitnil,omitempty" name:"Tls"`

	// eks clusterId
	// Note: this field may return null, indicating that no valid values can be obtained.
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// clb ip
	// Note: this field may return null, indicating that no valid values can be obtained.
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// Creation time.
	// Note: this field may return `null`, indicating that no valid value can be obtained.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Whether to listen on both the HTTP Port 80 and HTTPS Port 443. The default value is `false`. The optional value `true` means listening on both the HTTP Port 80 and HTTPS Port 443.
	Mixed *bool `json:"Mixed,omitnil,omitempty" name:"Mixed"`
}

type IngressRule struct {
	// ingress rule value
	Http *IngressRuleValue `json:"Http,omitnil,omitempty" name:"Http"`

	// Host address
	// Note: this field may return null, indicating that no valid values can be obtained.
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// Protocol. Options include HTTP and HTTPS. The default option is HTTP.
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`
}

type IngressRuleBackend struct {
	// EKS service name
	ServiceName *string `json:"ServiceName,omitnil,omitempty" name:"ServiceName"`

	// EKS service port
	ServicePort *int64 `json:"ServicePort,omitnil,omitempty" name:"ServicePort"`
}

type IngressRulePath struct {
	// Path information
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// Backend configuration
	Backend *IngressRuleBackend `json:"Backend,omitnil,omitempty" name:"Backend"`
}

type IngressRuleValue struct {
	// Overall rule configuration
	Paths []*IngressRulePath `json:"Paths,omitnil,omitempty" name:"Paths"`
}

type IngressTls struct {
	// Host array. An empty array indicates the default certificate for all domain names.
	Hosts []*string `json:"Hosts,omitnil,omitempty" name:"Hosts"`

	// Secret name. If a certificate is used, this field is left empty.
	SecretName *string `json:"SecretName,omitnil,omitempty" name:"SecretName"`

	// SSL Certificate Id
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`
}

type LogOutputConf struct {
	// Log consumer type
	OutputType *string `json:"OutputType,omitnil,omitempty" name:"OutputType"`

	// CLS logset
	ClsLogsetName *string `json:"ClsLogsetName,omitnil,omitempty" name:"ClsLogsetName"`

	// CLS log topic
	ClsLogTopicId *string `json:"ClsLogTopicId,omitnil,omitempty" name:"ClsLogTopicId"`

	// CLS logset ID
	ClsLogsetId *string `json:"ClsLogsetId,omitnil,omitempty" name:"ClsLogsetId"`

	// CLS log topic name
	ClsLogTopicName *string `json:"ClsLogTopicName,omitnil,omitempty" name:"ClsLogTopicName"`
}

// Predefined struct for user
type ModifyIngressRequestParams struct {
	// Ingress rule configuration
	Ingress *IngressInfo `json:"Ingress,omitnil,omitempty" name:"Ingress"`

	// Source channel
	SourceChannel *int64 `json:"SourceChannel,omitnil,omitempty" name:"SourceChannel"`
}

type ModifyIngressRequest struct {
	*tchttp.BaseRequest
	
	// Ingress rule configuration
	Ingress *IngressInfo `json:"Ingress,omitnil,omitempty" name:"Ingress"`

	// Source channel
	SourceChannel *int64 `json:"SourceChannel,omitnil,omitempty" name:"SourceChannel"`
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
	// Created successfully
	// Note: this field may return null, indicating that no valid values can be obtained.
	Result *bool `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type ModifyNamespaceRequestParams struct {
	// Environment ID
	NamespaceId *string `json:"NamespaceId,omitnil,omitempty" name:"NamespaceId"`

	// Namespace name
	NamespaceName *string `json:"NamespaceName,omitnil,omitempty" name:"NamespaceName"`

	// Namespace description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// VPC name
	Vpc *string `json:"Vpc,omitnil,omitempty" name:"Vpc"`

	// Subnet
	SubnetIds []*string `json:"SubnetIds,omitnil,omitempty" name:"SubnetIds"`

	// Source channel
	SourceChannel *int64 `json:"SourceChannel,omitnil,omitempty" name:"SourceChannel"`
}

type ModifyNamespaceRequest struct {
	*tchttp.BaseRequest
	
	// Environment ID
	NamespaceId *string `json:"NamespaceId,omitnil,omitempty" name:"NamespaceId"`

	// Namespace name
	NamespaceName *string `json:"NamespaceName,omitnil,omitempty" name:"NamespaceName"`

	// Namespace description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// VPC name
	Vpc *string `json:"Vpc,omitnil,omitempty" name:"Vpc"`

	// Subnet
	SubnetIds []*string `json:"SubnetIds,omitnil,omitempty" name:"SubnetIds"`

	// Source channel
	SourceChannel *int64 `json:"SourceChannel,omitnil,omitempty" name:"SourceChannel"`
}

func (r *ModifyNamespaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNamespaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NamespaceId")
	delete(f, "NamespaceName")
	delete(f, "Description")
	delete(f, "Vpc")
	delete(f, "SubnetIds")
	delete(f, "SourceChannel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyNamespaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNamespaceResponseParams struct {
	// Namespace ID in case of success and `null` in case of failure
	// Note: this field may return null, indicating that no valid values can be obtained.
	Result *bool `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyNamespaceResponse struct {
	*tchttp.BaseResponse
	Response *ModifyNamespaceResponseParams `json:"Response"`
}

func (r *ModifyNamespaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNamespaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyServiceInfoRequestParams struct {
	// Service ID.
	ServiceId *string `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`

	// Description.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Source channel.
	SourceChannel *int64 `json:"SourceChannel,omitnil,omitempty" name:"SourceChannel"`
}

type ModifyServiceInfoRequest struct {
	*tchttp.BaseRequest
	
	// Service ID.
	ServiceId *string `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`

	// Description.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Source channel.
	SourceChannel *int64 `json:"SourceChannel,omitnil,omitempty" name:"SourceChannel"`
}

func (r *ModifyServiceInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyServiceInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "Description")
	delete(f, "SourceChannel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyServiceInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyServiceInfoResponseParams struct {
	// Results.
	// Note: this field may return `null`, indicating that no valid value can be obtained.
	Result *bool `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyServiceInfoResponse struct {
	*tchttp.BaseResponse
	Response *ModifyServiceInfoResponseParams `json:"Response"`
}

func (r *ModifyServiceInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyServiceInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MountedSettingConf struct {
	// Configuration Name
	ConfigDataName *string `json:"ConfigDataName,omitnil,omitempty" name:"ConfigDataName"`

	// Mount point path
	MountedPath *string `json:"MountedPath,omitnil,omitempty" name:"MountedPath"`

	// Configuration Content
	Data []*Pair `json:"Data,omitnil,omitempty" name:"Data"`
}

type NamespacePage struct {
	// Records
	Records []*TemNamespaceInfo `json:"Records,omitnil,omitempty" name:"Records"`

	// Total number
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// Number of items
	Size *int64 `json:"Size,omitnil,omitempty" name:"Size"`

	// Number of pages
	Pages *int64 `json:"Pages,omitnil,omitempty" name:"Pages"`
}

type Pair struct {
	// Key
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// Value
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type PortMapping struct {
	// Port.
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// Mapped port.
	TargetPort *int64 `json:"TargetPort,omitnil,omitempty" name:"TargetPort"`

	// TCP/UDP protocol stack.
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`
}

// Predefined struct for user
type RestartServiceRunPodRequestParams struct {
	// Environment ID.
	NamespaceId *string `json:"NamespaceId,omitnil,omitempty" name:"NamespaceId"`

	// Service ID.
	ServiceId *string `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`

	// Pod name.
	PodName *string `json:"PodName,omitnil,omitempty" name:"PodName"`

	// Number of items per page.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Page number.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Pod status.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Source channel.
	SourceChannel *int64 `json:"SourceChannel,omitnil,omitempty" name:"SourceChannel"`
}

type RestartServiceRunPodRequest struct {
	*tchttp.BaseRequest
	
	// Environment ID.
	NamespaceId *string `json:"NamespaceId,omitnil,omitempty" name:"NamespaceId"`

	// Service ID.
	ServiceId *string `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`

	// Pod name.
	PodName *string `json:"PodName,omitnil,omitempty" name:"PodName"`

	// Number of items per page.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Page number.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Pod status.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Source channel.
	SourceChannel *int64 `json:"SourceChannel,omitnil,omitempty" name:"SourceChannel"`
}

func (r *RestartServiceRunPodRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestartServiceRunPodRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NamespaceId")
	delete(f, "ServiceId")
	delete(f, "PodName")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Status")
	delete(f, "SourceChannel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RestartServiceRunPodRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RestartServiceRunPodResponseParams struct {
	// Returned results.
	// Note: this field may return `null`, indicating that no valid value can be obtained.
	Result *bool `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RestartServiceRunPodResponse struct {
	*tchttp.BaseResponse
	Response *RestartServiceRunPodResponseParams `json:"Response"`
}

func (r *RestartServiceRunPodResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestartServiceRunPodResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RunVersionPod struct {
	// Shell address
	Webshell *string `json:"Webshell,omitnil,omitempty" name:"Webshell"`

	// Pod ID
	PodId *string `json:"PodId,omitnil,omitempty" name:"PodId"`

	// Status
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Creation time
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Pod IP.
	PodIp *string `json:"PodIp,omitnil,omitempty" name:"PodIp"`

	// Availability zone.
	// Note: this field may return `null`, indicating that no valid value can be obtained.
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Deployed version.
	// Note: this field may return `null`, indicating that no valid value can be obtained.
	DeployVersion *string `json:"DeployVersion,omitnil,omitempty" name:"DeployVersion"`

	// Number of Restarts
	// Note: This is field may return `null`, indicating that no valid value can be obtained.
	RestartCount *int64 `json:"RestartCount,omitnil,omitempty" name:"RestartCount"`
}

type StorageConf struct {
	// Storage volume name
	StorageVolName *string `json:"StorageVolName,omitnil,omitempty" name:"StorageVolName"`

	// Storage volume path
	StorageVolPath *string `json:"StorageVolPath,omitnil,omitempty" name:"StorageVolPath"`

	// Storage volume IP
	// Note: this field may return null, indicating that no valid values can be obtained.
	StorageVolIp *string `json:"StorageVolIp,omitnil,omitempty" name:"StorageVolIp"`
}

type StorageMountConf struct {
	// Data volume name
	VolumeName *string `json:"VolumeName,omitnil,omitempty" name:"VolumeName"`

	// Data volume binding path
	MountPath *string `json:"MountPath,omitnil,omitempty" name:"MountPath"`
}

type TemNamespaceInfo struct {
	// Namespace ID
	NamespaceId *string `json:"NamespaceId,omitnil,omitempty" name:"NamespaceId"`

	// Channel
	Channel *string `json:"Channel,omitnil,omitempty" name:"Channel"`

	// Namespace name
	NamespaceName *string `json:"NamespaceName,omitnil,omitempty" name:"NamespaceName"`

	// Region name
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Namespace description
	// Note: this field may return null, indicating that no valid values can be obtained.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Status. 1: terminated; 0: normal
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// VPC
	Vpc *string `json:"Vpc,omitnil,omitempty" name:"Vpc"`

	// Creation time
	CreateDate *string `json:"CreateDate,omitnil,omitempty" name:"CreateDate"`

	// Modification time
	ModifyDate *string `json:"ModifyDate,omitnil,omitempty" name:"ModifyDate"`

	// Modifier
	Modifier *string `json:"Modifier,omitnil,omitempty" name:"Modifier"`

	// Creator
	Creator *string `json:"Creator,omitnil,omitempty" name:"Creator"`

	// Number of services
	ServiceNum *int64 `json:"ServiceNum,omitnil,omitempty" name:"ServiceNum"`

	// Number of running instances
	RunInstancesNum *int64 `json:"RunInstancesNum,omitnil,omitempty" name:"RunInstancesNum"`

	// Subnet
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// TCB environment status
	TcbEnvStatus *string `json:"TcbEnvStatus,omitnil,omitempty" name:"TcbEnvStatus"`

	// eks cluster status
	ClusterStatus *string `json:"ClusterStatus,omitnil,omitempty" name:"ClusterStatus"`

	// Whether to enable TSW
	EnableTswTraceService *bool `json:"EnableTswTraceService,omitnil,omitempty" name:"EnableTswTraceService"`
}